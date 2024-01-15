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

func TestAccountUrlscannerScanNewWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Urlscanner.Scan.New(
		context.TODO(),
		"string",
		cloudflare.AccountUrlscannerScanNewParams{
			URL:                    cloudflare.F("https://www.example.com"),
			CustomHeaders:          cloudflare.F[any](map[string]interface{}{}),
			ScreenshotsResolutions: cloudflare.F([]cloudflare.AccountUrlscannerScanNewParamsScreenshotsResolution{cloudflare.AccountUrlscannerScanNewParamsScreenshotsResolutionDesktop, cloudflare.AccountUrlscannerScanNewParamsScreenshotsResolutionMobile, cloudflare.AccountUrlscannerScanNewParamsScreenshotsResolutionTablet}),
			Visibility:             cloudflare.F(cloudflare.AccountUrlscannerScanNewParamsVisibilityPublic),
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

func TestAccountUrlscannerScanGet(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Urlscanner.Scan.Get(
		context.TODO(),
		"string",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountUrlscannerScanListWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Urlscanner.Scan.List(
		context.TODO(),
		"string",
		cloudflare.AccountUrlscannerScanListParams{
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
