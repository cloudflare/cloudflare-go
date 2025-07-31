// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brand_protection_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/brand_protection"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func TestLogoMatchDownloadWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test, 401 Unauthorized error")
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
	_, err := client.BrandProtection.LogoMatches.Download(context.TODO(), brand_protection.LogoMatchDownloadParams{
		AccountID: cloudflare.F("x"),
		Limit:     cloudflare.F("limit"),
		LogoID:    cloudflare.F([]string{"string"}),
		Offset:    cloudflare.F("offset"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLogoMatchGetWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test, 401 Unauthorized error")
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
	_, err := client.BrandProtection.LogoMatches.Get(context.TODO(), brand_protection.LogoMatchGetParams{
		AccountID: cloudflare.F("x"),
		Limit:     cloudflare.F("limit"),
		LogoID:    cloudflare.F([]string{"string"}),
		Offset:    cloudflare.F("offset"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
