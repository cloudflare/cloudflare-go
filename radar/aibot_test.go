// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/radar"
)

func TestAIBotSummaryV2WithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.AI.Bots.SummaryV2(
		context.TODO(),
		radar.AIBotSummaryV2ParamsDimensionUserAgent,
		radar.AIBotSummaryV2Params{
			ASN:           cloudflare.F([]string{"string"}),
			ContentType:   cloudflare.F([]radar.AIBotSummaryV2ParamsContentType{radar.AIBotSummaryV2ParamsContentTypeHTML}),
			Continent:     cloudflare.F([]string{"string"}),
			CrawlPurpose:  cloudflare.F([]string{"string"}),
			DateEnd:       cloudflare.F([]time.Time{time.Now()}),
			DateRange:     cloudflare.F([]string{"7d"}),
			DateStart:     cloudflare.F([]time.Time{time.Now()}),
			Format:        cloudflare.F(radar.AIBotSummaryV2ParamsFormatJson),
			Industry:      cloudflare.F([]string{"string"}),
			LimitPerGroup: cloudflare.F(int64(10)),
			Location:      cloudflare.F([]string{"string"}),
			Name:          cloudflare.F([]string{"main_series"}),
			UserAgent:     cloudflare.F([]string{"string"}),
			Vertical:      cloudflare.F([]string{"string"}),
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

func TestAIBotTimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.AI.Bots.Timeseries(context.TODO(), radar.AIBotTimeseriesParams{
		AggInterval:   cloudflare.F(radar.AIBotTimeseriesParamsAggInterval1h),
		ASN:           cloudflare.F([]string{"string"}),
		ContentType:   cloudflare.F([]radar.AIBotTimeseriesParamsContentType{radar.AIBotTimeseriesParamsContentTypeHTML}),
		Continent:     cloudflare.F([]string{"string"}),
		CrawlPurpose:  cloudflare.F([]string{"string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now()}),
		DateRange:     cloudflare.F([]string{"7d"}),
		DateStart:     cloudflare.F([]time.Time{time.Now()}),
		Format:        cloudflare.F(radar.AIBotTimeseriesParamsFormatJson),
		Industry:      cloudflare.F([]string{"string"}),
		LimitPerGroup: cloudflare.F(int64(10)),
		Location:      cloudflare.F([]string{"string"}),
		Name:          cloudflare.F([]string{"main_series"}),
		UserAgent:     cloudflare.F([]string{"string"}),
		Vertical:      cloudflare.F([]string{"string"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIBotTimeseriesGroupsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.AI.Bots.TimeseriesGroups(
		context.TODO(),
		radar.AIBotTimeseriesGroupsParamsDimensionUserAgent,
		radar.AIBotTimeseriesGroupsParams{
			AggInterval:   cloudflare.F(radar.AIBotTimeseriesGroupsParamsAggInterval1h),
			ASN:           cloudflare.F([]string{"string"}),
			ContentType:   cloudflare.F([]radar.AIBotTimeseriesGroupsParamsContentType{radar.AIBotTimeseriesGroupsParamsContentTypeHTML}),
			Continent:     cloudflare.F([]string{"string"}),
			CrawlPurpose:  cloudflare.F([]string{"string"}),
			DateEnd:       cloudflare.F([]time.Time{time.Now()}),
			DateRange:     cloudflare.F([]string{"7d"}),
			DateStart:     cloudflare.F([]time.Time{time.Now()}),
			Format:        cloudflare.F(radar.AIBotTimeseriesGroupsParamsFormatJson),
			Industry:      cloudflare.F([]string{"string"}),
			LimitPerGroup: cloudflare.F(int64(10)),
			Location:      cloudflare.F([]string{"string"}),
			Name:          cloudflare.F([]string{"main_series"}),
			Normalization: cloudflare.F(radar.AIBotTimeseriesGroupsParamsNormalizationPercentage),
			UserAgent:     cloudflare.F([]string{"string"}),
			Vertical:      cloudflare.F([]string{"string"}),
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
