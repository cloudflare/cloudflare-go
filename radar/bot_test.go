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

func TestBotListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Bots.List(context.TODO(), radar.BotListParams{
		BotCategory:           cloudflare.F(radar.BotListParamsBotCategorySearchEngineCrawler),
		BotOperator:           cloudflare.F("botOperator"),
		BotVerificationStatus: cloudflare.F(radar.BotListParamsBotVerificationStatusVerified),
		Format:                cloudflare.F(radar.BotListParamsFormatJson),
		Kind:                  cloudflare.F(radar.BotListParamsKindAgent),
		Limit:                 cloudflare.F(int64(1)),
		Offset:                cloudflare.F(int64(0)),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBotGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Bots.Get(
		context.TODO(),
		"gptbot",
		radar.BotGetParams{
			Format: cloudflare.F(radar.BotGetParamsFormatJson),
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

func TestBotSummaryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Bots.Summary(
		context.TODO(),
		radar.BotSummaryParamsDimensionBot,
		radar.BotSummaryParams{
			ASN:                   cloudflare.F([]string{"string"}),
			Bot:                   cloudflare.F([]string{"string"}),
			BotCategory:           cloudflare.F([]radar.BotSummaryParamsBotCategory{radar.BotSummaryParamsBotCategorySearchEngineCrawler}),
			BotKind:               cloudflare.F([]radar.BotSummaryParamsBotKind{radar.BotSummaryParamsBotKindAgent}),
			BotOperator:           cloudflare.F([]string{"string"}),
			BotVerificationStatus: cloudflare.F([]radar.BotSummaryParamsBotVerificationStatus{radar.BotSummaryParamsBotVerificationStatusVerified}),
			Continent:             cloudflare.F([]string{"string"}),
			DateEnd:               cloudflare.F([]time.Time{time.Now()}),
			DateRange:             cloudflare.F([]string{"7d"}),
			DateStart:             cloudflare.F([]time.Time{time.Now()}),
			Format:                cloudflare.F(radar.BotSummaryParamsFormatJson),
			LimitPerGroup:         cloudflare.F(int64(10)),
			Location:              cloudflare.F([]string{"string"}),
			Name:                  cloudflare.F([]string{"main_series"}),
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

func TestBotTimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Bots.Timeseries(context.TODO(), radar.BotTimeseriesParams{
		AggInterval:           cloudflare.F(radar.BotTimeseriesParamsAggInterval1h),
		ASN:                   cloudflare.F([]string{"string"}),
		Bot:                   cloudflare.F([]string{"string"}),
		BotCategory:           cloudflare.F([]radar.BotTimeseriesParamsBotCategory{radar.BotTimeseriesParamsBotCategorySearchEngineCrawler}),
		BotKind:               cloudflare.F([]radar.BotTimeseriesParamsBotKind{radar.BotTimeseriesParamsBotKindAgent}),
		BotOperator:           cloudflare.F([]string{"string"}),
		BotVerificationStatus: cloudflare.F([]radar.BotTimeseriesParamsBotVerificationStatus{radar.BotTimeseriesParamsBotVerificationStatusVerified}),
		Continent:             cloudflare.F([]string{"string"}),
		DateEnd:               cloudflare.F([]time.Time{time.Now()}),
		DateRange:             cloudflare.F([]string{"7d"}),
		DateStart:             cloudflare.F([]time.Time{time.Now()}),
		Format:                cloudflare.F(radar.BotTimeseriesParamsFormatJson),
		Location:              cloudflare.F([]string{"string"}),
		Name:                  cloudflare.F([]string{"main_series"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBotTimeseriesGroupsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Bots.TimeseriesGroups(
		context.TODO(),
		radar.BotTimeseriesGroupsParamsDimensionBot,
		radar.BotTimeseriesGroupsParams{
			AggInterval:           cloudflare.F(radar.BotTimeseriesGroupsParamsAggInterval1h),
			ASN:                   cloudflare.F([]string{"string"}),
			Bot:                   cloudflare.F([]string{"string"}),
			BotCategory:           cloudflare.F([]radar.BotTimeseriesGroupsParamsBotCategory{radar.BotTimeseriesGroupsParamsBotCategorySearchEngineCrawler}),
			BotKind:               cloudflare.F([]radar.BotTimeseriesGroupsParamsBotKind{radar.BotTimeseriesGroupsParamsBotKindAgent}),
			BotOperator:           cloudflare.F([]string{"string"}),
			BotVerificationStatus: cloudflare.F([]radar.BotTimeseriesGroupsParamsBotVerificationStatus{radar.BotTimeseriesGroupsParamsBotVerificationStatusVerified}),
			Continent:             cloudflare.F([]string{"string"}),
			DateEnd:               cloudflare.F([]time.Time{time.Now()}),
			DateRange:             cloudflare.F([]string{"7d"}),
			DateStart:             cloudflare.F([]time.Time{time.Now()}),
			Format:                cloudflare.F(radar.BotTimeseriesGroupsParamsFormatJson),
			LimitPerGroup:         cloudflare.F(int64(10)),
			Location:              cloudflare.F([]string{"string"}),
			Name:                  cloudflare.F([]string{"main_series"}),
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
