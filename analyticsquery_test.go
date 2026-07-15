// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/shared"
)

func TestAnalyticsQuerySummary(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.AnalyticsQuery.Summary(
		context.TODO(),
		"access-logins",
		cloudflare.AnalyticsQuerySummaryParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Filters: cloudflare.F([]cloudflare.AnalyticsQuerySummaryParamsFilter{{
				Name:   cloudflare.F("country"),
				Op:     cloudflare.F("in"),
				Values: cloudflare.F([]cloudflare.AnalyticsQuerySummaryParamsFiltersValueUnion{shared.UnionString("US"), shared.UnionString("CA"), shared.UnionString("GB")}),
			}}),
			From:    cloudflare.F(time.Now()),
			GroupBy: cloudflare.F([]string{"string"}),
			Stats:   cloudflare.F([]string{"attemptsTotal"}),
			To:      cloudflare.F(time.Now()),
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

func TestAnalyticsQueryTimeseries(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.AnalyticsQuery.Timeseries(
		context.TODO(),
		"shadow_it",
		cloudflare.AnalyticsQueryTimeseriesParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Filters: cloudflare.F([]cloudflare.AnalyticsQueryTimeseriesParamsFilter{{
				Name:   cloudflare.F("allowed"),
				Op:     cloudflare.F("eq"),
				Values: cloudflare.F([]cloudflare.AnalyticsQueryTimeseriesParamsFiltersValueUnion{shared.UnionBool(true)}),
			}}),
			From:       cloudflare.F(time.Now()),
			GroupBy:    cloudflare.F([]string{"country", "allowed"}),
			Resolution: cloudflare.F("day"),
			Stats:      cloudflare.F([]string{"attemptsTotal"}),
			To:         cloudflare.F(time.Now()),
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

func TestAnalyticsQueryTopN(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.AnalyticsQuery.TopN(
		context.TODO(),
		"gateway-http",
		cloudflare.AnalyticsQueryTopNParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Filters: cloudflare.F([]cloudflare.AnalyticsQueryTopNParamsFilter{{
				Name:   cloudflare.F("country"),
				Op:     cloudflare.F("in"),
				Values: cloudflare.F([]cloudflare.AnalyticsQueryTopNParamsFiltersValueUnion{shared.UnionString("US"), shared.UnionString("CA"), shared.UnionString("GB")}),
			}}),
			From:    cloudflare.F(time.Now()),
			GroupBy: cloudflare.F([]string{"appName", "appCategory"}),
			N:       cloudflare.F(int64(10)),
			OrderBy: cloudflare.F("bytesTotal"),
			Stats:   cloudflare.F([]string{"bytesTotal", "requestsTotal"}),
			To:      cloudflare.F(time.Now()),
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
