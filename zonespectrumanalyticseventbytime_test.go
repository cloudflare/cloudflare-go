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

func TestZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Spectrums.Analytics.Events.Bytimes.SpectrumAnalyticsByTimeGetAnalyticsByTime(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParams{
			Dimensions: cloudflare.F([]cloudflare.ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsDimension{cloudflare.ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsDimensionEvent, cloudflare.ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsDimensionAppID}),
			Filters:    cloudflare.F("event==disconnect%20AND%20coloName!=SFO"),
			Metrics:    cloudflare.F([]cloudflare.ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetric{cloudflare.ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetricCount, cloudflare.ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetricBytesIngress}),
			Since:      cloudflare.F(time.Now()),
			Sort:       cloudflare.F([]interface{}{"+count", "-bytesIngress"}),
			TimeDelta:  cloudflare.F(cloudflare.ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDeltaMinute),
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
