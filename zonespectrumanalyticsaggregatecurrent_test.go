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

func TestZoneSpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Spectrums.Analytics.Aggregates.Currents.SpectrumAggregateAnalyticsGetCurrentAggregatedAnalytics(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZoneSpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsParams{
			AppIDParam: cloudflare.F("ea95132c15732412d22c1476fa83f27a,d122c5f4bb71e25cc9e86ab43b142e2f"),
			AppID:      cloudflare.F("ea95132c15732412d22c1476fa83f27a,d122c5f4bb71e25cc9e86ab43b142e2f"),
			ColoName:   cloudflare.F("PDX"),
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
