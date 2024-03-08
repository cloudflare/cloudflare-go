// File generated from our OpenAPI spec by Stainless.

package url_scanner_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/cloudflare/cloudflare-go/url_scanner"
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.URLScanner.Scan(
		context.TODO(),
		"string",
		url_scanner.URLScannerScanParams{
			AccountScans: cloudflare.F(true),
			DateEnd:      cloudflare.F(time.Now()),
			DateStart:    cloudflare.F(time.Now()),
			Hostname:     cloudflare.F("example.com"),
			IP:           cloudflare.F("1.1.1.1"),
			Limit:        cloudflare.F(int64(100)),
			NextCursor:   cloudflare.F("string"),
			PageHostname: cloudflare.F("string"),
			PageIP:       cloudflare.F("string"),
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
