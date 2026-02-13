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

func TestAITimeseriesGroupSummaryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.AI.TimeseriesGroups.Summary(
		context.TODO(),
		radar.AITimeseriesGroupSummaryParamsDimensionUserAgent,
		radar.AITimeseriesGroupSummaryParams{
			ASN:           cloudflare.F([]string{"string"}),
			ContentType:   cloudflare.F([]radar.AITimeseriesGroupSummaryParamsContentType{radar.AITimeseriesGroupSummaryParamsContentTypeHTML}),
			Continent:     cloudflare.F([]string{"string"}),
			CrawlPurpose:  cloudflare.F([]string{"string"}),
			DateEnd:       cloudflare.F([]time.Time{time.Now()}),
			DateRange:     cloudflare.F([]string{"7d"}),
			DateStart:     cloudflare.F([]time.Time{time.Now()}),
			Format:        cloudflare.F(radar.AITimeseriesGroupSummaryParamsFormatJson),
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

func TestAITimeseriesGroupTimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.AI.TimeseriesGroups.Timeseries(context.TODO(), radar.AITimeseriesGroupTimeseriesParams{
		AggInterval:   cloudflare.F(radar.AITimeseriesGroupTimeseriesParamsAggInterval1h),
		ASN:           cloudflare.F([]string{"string"}),
		ContentType:   cloudflare.F([]radar.AITimeseriesGroupTimeseriesParamsContentType{radar.AITimeseriesGroupTimeseriesParamsContentTypeHTML}),
		Continent:     cloudflare.F([]string{"string"}),
		CrawlPurpose:  cloudflare.F([]string{"string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now()}),
		DateRange:     cloudflare.F([]string{"7d"}),
		DateStart:     cloudflare.F([]time.Time{time.Now()}),
		Format:        cloudflare.F(radar.AITimeseriesGroupTimeseriesParamsFormatJson),
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

func TestAITimeseriesGroupTimeseriesGroupsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.AI.TimeseriesGroups.TimeseriesGroups(
		context.TODO(),
		radar.AITimeseriesGroupTimeseriesGroupsParamsDimensionUserAgent,
		radar.AITimeseriesGroupTimeseriesGroupsParams{
			AggInterval:   cloudflare.F(radar.AITimeseriesGroupTimeseriesGroupsParamsAggInterval1h),
			ASN:           cloudflare.F([]string{"string"}),
			ContentType:   cloudflare.F([]radar.AITimeseriesGroupTimeseriesGroupsParamsContentType{radar.AITimeseriesGroupTimeseriesGroupsParamsContentTypeHTML}),
			Continent:     cloudflare.F([]string{"string"}),
			CrawlPurpose:  cloudflare.F([]string{"string"}),
			DateEnd:       cloudflare.F([]time.Time{time.Now()}),
			DateRange:     cloudflare.F([]string{"7d"}),
			DateStart:     cloudflare.F([]time.Time{time.Now()}),
			Format:        cloudflare.F(radar.AITimeseriesGroupTimeseriesGroupsParamsFormatJson),
			Industry:      cloudflare.F([]string{"string"}),
			LimitPerGroup: cloudflare.F(int64(10)),
			Location:      cloudflare.F([]string{"string"}),
			Name:          cloudflare.F([]string{"main_series"}),
			Normalization: cloudflare.F(radar.AITimeseriesGroupTimeseriesGroupsParamsNormalizationPercentage),
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

func TestAITimeseriesGroupUserAgentWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.AI.TimeseriesGroups.UserAgent(context.TODO(), radar.AITimeseriesGroupUserAgentParams{
		AggInterval:   cloudflare.F(radar.AITimeseriesGroupUserAgentParamsAggInterval1h),
		ASN:           cloudflare.F([]string{"string"}),
		Continent:     cloudflare.F([]string{"string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now()}),
		DateRange:     cloudflare.F([]string{"7d"}),
		DateStart:     cloudflare.F([]time.Time{time.Now()}),
		Format:        cloudflare.F(radar.AITimeseriesGroupUserAgentParamsFormatJson),
		LimitPerGroup: cloudflare.F(int64(10)),
		Location:      cloudflare.F([]string{"string"}),
		Name:          cloudflare.F([]string{"main_series"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
