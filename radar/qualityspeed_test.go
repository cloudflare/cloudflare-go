// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/radar"
)

func TestQualitySpeedHistogramWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Quality.Speed.Histogram(context.TODO(), radar.QualitySpeedHistogramParams{
		ASN:         cloudflare.F([]string{"string"}),
		BucketSize:  cloudflare.F(int64(0)),
		Continent:   cloudflare.F([]string{"string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now()}),
		Format:      cloudflare.F(radar.QualitySpeedHistogramParamsFormatJson),
		Location:    cloudflare.F([]string{"string"}),
		MetricGroup: cloudflare.F(radar.QualitySpeedHistogramParamsMetricGroupBandwidth),
		Name:        cloudflare.F([]string{"string"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestQualitySpeedSummaryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Quality.Speed.Summary(context.TODO(), radar.QualitySpeedSummaryParams{
		ASN:       cloudflare.F([]string{"string"}),
		Continent: cloudflare.F([]string{"string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now()}),
		Format:    cloudflare.F(radar.QualitySpeedSummaryParamsFormatJson),
		Location:  cloudflare.F([]string{"string"}),
		Name:      cloudflare.F([]string{"string"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
