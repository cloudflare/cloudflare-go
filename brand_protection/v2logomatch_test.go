// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brand_protection_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/brand_protection"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

func TestV2LogoMatchGetWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.BrandProtection.V2.LogoMatches.Get(context.TODO(), brand_protection.V2LogoMatchGetParams{
		AccountID: cloudflare.F("x"),
		QueryID:   cloudflare.F("x"),
		Download:  cloudflare.F("download"),
		Limit:     cloudflare.F("limit"),
		Offset:    cloudflare.F("offset"),
		Order:     cloudflare.F(brand_protection.V2LogoMatchGetParamsOrderAsc),
		OrderBy:   cloudflare.F(brand_protection.V2LogoMatchGetParamsOrderByMatchedAt),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
