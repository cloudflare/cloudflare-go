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

func TestAnalyticsQueryDataSecurityFindingSummary(t *testing.T) {
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
	_, err := client.AnalyticsQuery.DataSecurity.Findings.Summary(context.TODO(), cloudflare.AnalyticsQueryDataSecurityFindingSummaryParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Filters: cloudflare.F([]cloudflare.AnalyticsQueryDataSecurityFindingSummaryParamsFilter{{
			Name:   cloudflare.F("country"),
			Op:     cloudflare.F("in"),
			Values: cloudflare.F([]cloudflare.AnalyticsQueryDataSecurityFindingSummaryParamsFiltersValueUnion{shared.UnionString("US"), shared.UnionString("CA"), shared.UnionString("GB")}),
		}}),
		From: cloudflare.F(time.Now()),
		To:   cloudflare.F(time.Now()),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAnalyticsQueryDataSecurityFindingTimeseries(t *testing.T) {
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
	_, err := client.AnalyticsQuery.DataSecurity.Findings.Timeseries(context.TODO(), cloudflare.AnalyticsQueryDataSecurityFindingTimeseriesParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Filters: cloudflare.F([]cloudflare.AnalyticsQueryDataSecurityFindingTimeseriesParamsFilter{{
			Name:   cloudflare.F("country"),
			Op:     cloudflare.F("in"),
			Values: cloudflare.F([]cloudflare.AnalyticsQueryDataSecurityFindingTimeseriesParamsFiltersValueUnion{shared.UnionString("US"), shared.UnionString("CA"), shared.UnionString("GB")}),
		}}),
		From: cloudflare.F(time.Now()),
		To:   cloudflare.F(time.Now()),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
