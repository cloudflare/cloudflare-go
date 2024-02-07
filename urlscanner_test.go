// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestURLScannerScanWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.URLScanner.Scan(
		context.TODO(),
		"string",
		cloudflare.URLScannerScanParams{
			AccountScans: cloudflare.F(true),
			DateEnd:      cloudflare.F(time.Now()),
			DateStart:    cloudflare.F(time.Now()),
			Hostname:     cloudflare.F("example.com"),
			Limit:        cloudflare.F(int64(100)),
			NextCursor:   cloudflare.F("string"),
			PageHostname: cloudflare.F("string"),
			PagePath:     cloudflare.F("string"),
			PageURL:      cloudflare.F("string"),
			Path:         cloudflare.F("/samples/subresource-integrity/"),
			ScanID:       cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			URL:          cloudflare.F("https://example.com/?hello"),
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
