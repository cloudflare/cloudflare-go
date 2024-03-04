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

func TestRadarQualityIQISummaryWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Radar.Quality.IQI.Summary(context.TODO(), cloudflare.RadarQualityIQISummaryParams{
		Metric:    cloudflare.F(cloudflare.RadarQualityIQISummaryParamsMetricBandwidth),
		ASN:       cloudflare.F([]string{"string", "string", "string"}),
		Continent: cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarQualityIQISummaryParamsDateRange{cloudflare.RadarQualityIQISummaryParamsDateRange1d, cloudflare.RadarQualityIQISummaryParamsDateRange2d, cloudflare.RadarQualityIQISummaryParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:    cloudflare.F(cloudflare.RadarQualityIQISummaryParamsFormatJson),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarQualityIQITimeseriesGroupsWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Radar.Quality.IQI.TimeseriesGroups(context.TODO(), cloudflare.RadarQualityIQITimeseriesGroupsParams{
		Metric:        cloudflare.F(cloudflare.RadarQualityIQITimeseriesGroupsParamsMetricBandwidth),
		AggInterval:   cloudflare.F(cloudflare.RadarQualityIQITimeseriesGroupsParamsAggInterval1h),
		ASN:           cloudflare.F([]string{"string", "string", "string"}),
		Continent:     cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     cloudflare.F([]cloudflare.RadarQualityIQITimeseriesGroupsParamsDateRange{cloudflare.RadarQualityIQITimeseriesGroupsParamsDateRange1d, cloudflare.RadarQualityIQITimeseriesGroupsParamsDateRange2d, cloudflare.RadarQualityIQITimeseriesGroupsParamsDateRange7d}),
		DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:        cloudflare.F(cloudflare.RadarQualityIQITimeseriesGroupsParamsFormatJson),
		Interpolation: cloudflare.F(true),
		Location:      cloudflare.F([]string{"string", "string", "string"}),
		Name:          cloudflare.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
