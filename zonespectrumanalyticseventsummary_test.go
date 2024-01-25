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

func TestZoneSpectrumAnalyticsEventSummarySpectrumAnalyticsSummaryGetAnalyticsSummaryWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Spectrums.Analytics.Events.Summaries.SpectrumAnalyticsSummaryGetAnalyticsSummary(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZoneSpectrumAnalyticsEventSummarySpectrumAnalyticsSummaryGetAnalyticsSummaryParams{
			Dimensions: cloudflare.F([]cloudflare.ZoneSpectrumAnalyticsEventSummarySpectrumAnalyticsSummaryGetAnalyticsSummaryParamsDimension{cloudflare.ZoneSpectrumAnalyticsEventSummarySpectrumAnalyticsSummaryGetAnalyticsSummaryParamsDimensionEvent, cloudflare.ZoneSpectrumAnalyticsEventSummarySpectrumAnalyticsSummaryGetAnalyticsSummaryParamsDimensionAppID}),
			Filters:    cloudflare.F("event==disconnect%20AND%20coloName!=SFO"),
			Metrics:    cloudflare.F([]cloudflare.ZoneSpectrumAnalyticsEventSummarySpectrumAnalyticsSummaryGetAnalyticsSummaryParamsMetric{cloudflare.ZoneSpectrumAnalyticsEventSummarySpectrumAnalyticsSummaryGetAnalyticsSummaryParamsMetricCount, cloudflare.ZoneSpectrumAnalyticsEventSummarySpectrumAnalyticsSummaryGetAnalyticsSummaryParamsMetricBytesIngress}),
			Since:      cloudflare.F(time.Now()),
			Sort:       cloudflare.F([]interface{}{"+count", "-bytesIngress"}),
			Until:      cloudflare.F(time.Now()),
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
