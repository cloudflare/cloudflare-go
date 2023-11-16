// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestContextCancel(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	cancelCtx, cancel := context.WithCancel(context.Background())
	cancel()
	res, err := client.Zones.New(cancelCtx, cloudflare.ZoneNewParams{
		Account: cloudflare.F(cloudflare.ZoneNewParamsAccount{
			ID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		}),
		Name: cloudflare.F("example.com"),
		Type: cloudflare.F(cloudflare.ZoneNewParamsTypeFull),
	})
	if err == nil || res != nil {
		t.Error("Expected there to be a cancel error and for the response to be nil")
	}
}

// neverTransport never completes a request and waits for the Context to be done.
type neverTransport struct{}

func (t *neverTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	<-req.Context().Done()
	return nil, fmt.Errorf("cancelled")
}

func TestContextCancelDelay(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
		option.WithHTTPClient(&http.Client{Transport: &neverTransport{}}),
	)
	cancelCtx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Millisecond * time.Duration(2))
		cancel()
	}()
	res, err := client.Zones.New(cancelCtx, cloudflare.ZoneNewParams{
		Account: cloudflare.F(cloudflare.ZoneNewParamsAccount{
			ID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		}),
		Name: cloudflare.F("example.com"),
		Type: cloudflare.F(cloudflare.ZoneNewParamsTypeFull),
	})
	if err == nil || res != nil {
		t.Error("expected there to be a cancel error and for the response to be nil")
	}
}

func TestContextDeadline(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}

	testTimeout := time.After(3 * time.Second)
	testDone := make(chan bool)

	deadline := time.Now().Add(100 * time.Millisecond)
	deadlineCtx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	go func() {
		client := cloudflare.NewClient(
			option.WithBaseURL(baseURL),
			option.WithAPIKey("my-cloudflare-api-key"),
			option.WithEmail("dev@cloudflare.com"),
			option.WithHTTPClient(&http.Client{Transport: &neverTransport{}}),
		)
		res, err := client.Zones.New(deadlineCtx, cloudflare.ZoneNewParams{
			Account: cloudflare.F(cloudflare.ZoneNewParamsAccount{
				ID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			}),
			Name: cloudflare.F("example.com"),
			Type: cloudflare.F(cloudflare.ZoneNewParamsTypeFull),
		})
		if err == nil || res != nil {
			t.Error("expected there to be a deadline error and for the response to be nil")
		}
		testDone <- true
	}()

	select {
	case <-testTimeout:
		t.Fatal("client didn't finish in time")
	case <-testDone:
		diff := time.Now().Sub(deadline)
		if diff < -20*time.Millisecond || 20*time.Millisecond < diff {
			t.Logf("error difference: %v", diff)
			t.Fatal("client did not return within 20ms of context deadline")
		}
	}
}
