// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestOriginTLSClientAuthHostnameUpdate(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.OriginTLSClientAuth.Hostnames.Update(context.TODO(), cloudflare.OriginTLSClientAuthHostnameUpdateParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Config: cloudflare.F([]cloudflare.OriginTLSClientAuthHostnameUpdateParamsConfig{{
			CertID:   cloudflare.F("2458ce5a-0c35-4c7f-82c7-8e9487d3ff60"),
			Enabled:  cloudflare.F(true),
			Hostname: cloudflare.F("app.example.com"),
		}, {
			CertID:   cloudflare.F("2458ce5a-0c35-4c7f-82c7-8e9487d3ff60"),
			Enabled:  cloudflare.F(true),
			Hostname: cloudflare.F("app.example.com"),
		}, {
			CertID:   cloudflare.F("2458ce5a-0c35-4c7f-82c7-8e9487d3ff60"),
			Enabled:  cloudflare.F(true),
			Hostname: cloudflare.F("app.example.com"),
		}}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOriginTLSClientAuthHostnameGet(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.OriginTLSClientAuth.Hostnames.Get(
		context.TODO(),
		"app.example.com",
		cloudflare.OriginTLSClientAuthHostnameGetParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
