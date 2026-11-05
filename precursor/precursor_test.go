// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package precursor_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/precursor"
)

func TestPrecursorUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Precursor.Update(context.TODO(), precursor.PrecursorUpdateParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		PrecursorConfig: precursor.PrecursorConfigParam{
			Mode: cloudflare.F(precursor.PrecursorConfigModeMonitor),
			Rules: cloudflare.F([]precursor.PrecursorConfigRuleParam{{
				Expression:       cloudflare.F(`http.request.uri.path eq "/shop"`),
				ShowInterstitial: cloudflare.F(true),
				Description:      cloudflare.F("Show an interstitial on the shop page"),
				Enabled:          cloudflare.F(true),
			}}),
			ShowInterstitial: cloudflare.F(false),
		},
		CloudflareVersion: cloudflare.F(precursor.PrecursorUpdateParamsCloudflareVersion1_0),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPrecursorGet(t *testing.T) {
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
	_, err := client.Precursor.Get(context.TODO(), precursor.PrecursorGetParams{
		ZoneID:            cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		CloudflareVersion: cloudflare.F(precursor.PrecursorGetParamsCloudflareVersion1_0),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
