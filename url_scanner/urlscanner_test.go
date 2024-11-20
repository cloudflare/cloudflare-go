// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package url_scanner_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/url_scanner"
)

func TestURLScannerBulk(t *testing.T) {
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
	_, err := client.URLScanner.Bulk(
		context.TODO(),
		"accountId",
		url_scanner.URLScannerBulkParams{
			Body: []url_scanner.URLScannerBulkParamsBody{{
				URL:         cloudflare.F("https://www.example.com"),
				Customagent: cloudflare.F("customagent"),
				CustomHeaders: cloudflare.F(map[string]string{
					"foo": "string",
				}),
				Referer:                cloudflare.F("referer"),
				ScreenshotsResolutions: cloudflare.F([]url_scanner.URLScannerBulkParamsBodyScreenshotsResolution{url_scanner.URLScannerBulkParamsBodyScreenshotsResolutionDesktop}),
				Visibility:             cloudflare.F(url_scanner.URLScannerBulkParamsBodyVisibilityPublic),
			}},
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
