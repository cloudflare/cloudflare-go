// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/radar"
)

func TestNetflowSummaryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Netflows.Summary(context.TODO(), radar.NetflowSummaryParams{
		ASN:       cloudflare.F([]string{"string"}),
		Continent: cloudflare.F([]string{"string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now()}),
		DateRange: cloudflare.F([]string{"7d"}),
		DateStart: cloudflare.F([]time.Time{time.Now()}),
		Format:    cloudflare.F(radar.NetflowSummaryParamsFormatJson),
		Location:  cloudflare.F([]string{"string"}),
		Name:      cloudflare.F([]string{"main_series"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNetflowTimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Netflows.Timeseries(context.TODO(), radar.NetflowTimeseriesParams{
		AggInterval:   cloudflare.F(radar.NetflowTimeseriesParamsAggInterval1h),
		ASN:           cloudflare.F([]string{"string"}),
		Continent:     cloudflare.F([]string{"string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now()}),
		DateRange:     cloudflare.F([]string{"7d"}),
		DateStart:     cloudflare.F([]time.Time{time.Now()}),
		Format:        cloudflare.F(radar.NetflowTimeseriesParamsFormatJson),
		Location:      cloudflare.F([]string{"string"}),
		Name:          cloudflare.F([]string{"main_series"}),
		Normalization: cloudflare.F(radar.NetflowTimeseriesParamsNormalizationMin0Max),
		Product:       cloudflare.F([]radar.NetflowTimeseriesParamsProduct{radar.NetflowTimeseriesParamsProductHTTP}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
