// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancers_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/load_balancers"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func TestMonitorPreviewNewWithOptionalParams(t *testing.T) {
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
	_, err := client.LoadBalancers.Monitors.Previews.New(
		context.TODO(),
		"f1aba936b94213e5b8dca0c0dbf1f9cc",
		load_balancers.MonitorPreviewNewParams{
			AccountID:       cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			AllowInsecure:   cloudflare.F(true),
			ConsecutiveDown: cloudflare.F(int64(0)),
			ConsecutiveUp:   cloudflare.F(int64(0)),
			Description:     cloudflare.F("Login page monitor"),
			ExpectedBody:    cloudflare.F("alive"),
			ExpectedCodes:   cloudflare.F("2xx"),
			FollowRedirects: cloudflare.F(true),
			Header: cloudflare.F(map[string][]string{
				"Host":     {"example.com"},
				"X-App-ID": {"abc123"},
			}),
			Interval:  cloudflare.F(int64(0)),
			Method:    cloudflare.F("GET"),
			Path:      cloudflare.F("/health"),
			Port:      cloudflare.F(int64(0)),
			ProbeZone: cloudflare.F("example.com"),
			Retries:   cloudflare.F(int64(0)),
			Timeout:   cloudflare.F(int64(0)),
			Type:      cloudflare.F(load_balancers.MonitorPreviewNewParamsTypeHTTP),
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
