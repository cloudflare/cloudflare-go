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

func TestEmailSecuritySummaryV2WithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.SummaryV2(
		context.TODO(),
		radar.EmailSecuritySummaryV2ParamsDimensionSpam,
		radar.EmailSecuritySummaryV2Params{
			ARC:           cloudflare.F([]radar.EmailSecuritySummaryV2ParamsARC{radar.EmailSecuritySummaryV2ParamsARCPass}),
			DateEnd:       cloudflare.F([]time.Time{time.Now()}),
			DateRange:     cloudflare.F([]string{"7d"}),
			DateStart:     cloudflare.F([]time.Time{time.Now()}),
			DKIM:          cloudflare.F([]radar.EmailSecuritySummaryV2ParamsDKIM{radar.EmailSecuritySummaryV2ParamsDKIMPass}),
			DMARC:         cloudflare.F([]radar.EmailSecuritySummaryV2ParamsDMARC{radar.EmailSecuritySummaryV2ParamsDMARCPass}),
			Format:        cloudflare.F(radar.EmailSecuritySummaryV2ParamsFormatJson),
			LimitPerGroup: cloudflare.F(int64(10)),
			Name:          cloudflare.F([]string{"main_series"}),
			SPF:           cloudflare.F([]radar.EmailSecuritySummaryV2ParamsSPF{radar.EmailSecuritySummaryV2ParamsSPFPass}),
			TLSVersion:    cloudflare.F([]radar.EmailSecuritySummaryV2ParamsTLSVersion{radar.EmailSecuritySummaryV2ParamsTLSVersionTlSv1_0}),
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

func TestEmailSecurityTimeseriesGroupsV2WithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroupsV2(
		context.TODO(),
		radar.EmailSecurityTimeseriesGroupsV2ParamsDimensionSpam,
		radar.EmailSecurityTimeseriesGroupsV2Params{
			AggInterval:   cloudflare.F(radar.EmailSecurityTimeseriesGroupsV2ParamsAggInterval1h),
			ARC:           cloudflare.F([]radar.EmailSecurityTimeseriesGroupsV2ParamsARC{radar.EmailSecurityTimeseriesGroupsV2ParamsARCPass}),
			DateEnd:       cloudflare.F([]time.Time{time.Now()}),
			DateRange:     cloudflare.F([]string{"7d"}),
			DateStart:     cloudflare.F([]time.Time{time.Now()}),
			DKIM:          cloudflare.F([]radar.EmailSecurityTimeseriesGroupsV2ParamsDKIM{radar.EmailSecurityTimeseriesGroupsV2ParamsDKIMPass}),
			DMARC:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupsV2ParamsDMARC{radar.EmailSecurityTimeseriesGroupsV2ParamsDMARCPass}),
			Format:        cloudflare.F(radar.EmailSecurityTimeseriesGroupsV2ParamsFormatJson),
			LimitPerGroup: cloudflare.F(int64(10)),
			Name:          cloudflare.F([]string{"main_series"}),
			SPF:           cloudflare.F([]radar.EmailSecurityTimeseriesGroupsV2ParamsSPF{radar.EmailSecurityTimeseriesGroupsV2ParamsSPFPass}),
			TLSVersion:    cloudflare.F([]radar.EmailSecurityTimeseriesGroupsV2ParamsTLSVersion{radar.EmailSecurityTimeseriesGroupsV2ParamsTLSVersionTlSv1_0}),
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
