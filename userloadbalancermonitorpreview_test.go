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

func TestUserLoadBalancerMonitorPreviewLoadBalancerMonitorsPreviewMonitorWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Users.LoadBalancers.Monitors.Previews.LoadBalancerMonitorsPreviewMonitor(
		context.TODO(),
		"f1aba936b94213e5b8dca0c0dbf1f9cc",
		cloudflare.UserLoadBalancerMonitorPreviewLoadBalancerMonitorsPreviewMonitorParams{
			ExpectedCodes:   cloudflare.F("2xx"),
			AllowInsecure:   cloudflare.F(true),
			ExpectedBody:    cloudflare.F("alive"),
			FollowRedirects: cloudflare.F(true),
			Header: cloudflare.F[any](map[string]interface{}{
				"Host": map[string]interface{}{
					"0": "example.com",
				},
				"X-App-ID": map[string]interface{}{
					"0": "abc123",
				},
			}),
			Method:  cloudflare.F("GET"),
			Path:    cloudflare.F("/health"),
			Port:    cloudflare.F(int64(0)),
			Retries: cloudflare.F(int64(0)),
			Timeout: cloudflare.F(int64(0)),
			Type:    cloudflare.F(cloudflare.UserLoadBalancerMonitorPreviewLoadBalancerMonitorsPreviewMonitorParamsTypeHTTPs),
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
