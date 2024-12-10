// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/spectrum"
)

func TestAnalyticsEventSummaryGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Spectrum.Analytics.Events.Summaries.Get(context.TODO(), spectrum.AnalyticsEventSummaryGetParams{
		ZoneID:     cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Dimensions: cloudflare.F([]spectrum.Dimension{spectrum.DimensionEvent, spectrum.DimensionAppID}),
		Filters:    cloudflare.F("event==disconnect%20AND%20coloName!=SFO"),
		Metrics:    cloudflare.F([]spectrum.AnalyticsEventSummaryGetParamsMetric{spectrum.AnalyticsEventSummaryGetParamsMetricCount, spectrum.AnalyticsEventSummaryGetParamsMetricBytesIngress}),
		Since:      cloudflare.F(time.Now()),
		Sort:       cloudflare.F([]string{"+count", "-bytesIngress"}),
		Until:      cloudflare.F(time.Now()),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
