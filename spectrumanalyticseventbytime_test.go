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

func TestSpectrumAnalyticsEventBytimeGetWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Spectrum.Analytics.Events.Bytimes.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.SpectrumAnalyticsEventBytimeGetParams{
			Dimensions: cloudflare.F([]cloudflare.SpectrumAnalyticsEventBytimeGetParamsDimension{cloudflare.SpectrumAnalyticsEventBytimeGetParamsDimensionEvent, cloudflare.SpectrumAnalyticsEventBytimeGetParamsDimensionAppID}),
			Filters:    cloudflare.F("event==disconnect%20AND%20coloName!=SFO"),
			Metrics:    cloudflare.F([]cloudflare.SpectrumAnalyticsEventBytimeGetParamsMetric{cloudflare.SpectrumAnalyticsEventBytimeGetParamsMetricCount, cloudflare.SpectrumAnalyticsEventBytimeGetParamsMetricBytesIngress}),
			Since:      cloudflare.F(time.Now()),
			Sort:       cloudflare.F([]interface{}{"+count", "-bytesIngress"}),
			TimeDelta:  cloudflare.F(cloudflare.SpectrumAnalyticsEventBytimeGetParamsTimeDeltaMinute),
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
