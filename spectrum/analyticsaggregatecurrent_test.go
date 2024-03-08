// File generated from our OpenAPI spec by Stainless.

package spectrum_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/cloudflare/cloudflare-go/spectrum"
)

func TestAnalyticsAggregateCurrentGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Spectrum.Analytics.Aggregates.Currents.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		spectrum.AnalyticsAggregateCurrentGetParams{
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
