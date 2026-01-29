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

func TestEmailRoutingSummaryV2WithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.SummaryV2(
		context.TODO(),
		radar.EmailRoutingSummaryV2ParamsDimensionIPVersion,
		radar.EmailRoutingSummaryV2Params{
			ARC:           cloudflare.F([]radar.EmailRoutingSummaryV2ParamsARC{radar.EmailRoutingSummaryV2ParamsARCPass}),
			DateEnd:       cloudflare.F([]time.Time{time.Now()}),
			DateRange:     cloudflare.F([]string{"7d"}),
			DateStart:     cloudflare.F([]time.Time{time.Now()}),
			DKIM:          cloudflare.F([]radar.EmailRoutingSummaryV2ParamsDKIM{radar.EmailRoutingSummaryV2ParamsDKIMPass}),
			DMARC:         cloudflare.F([]radar.EmailRoutingSummaryV2ParamsDMARC{radar.EmailRoutingSummaryV2ParamsDMARCPass}),
			Encrypted:     cloudflare.F([]radar.EmailRoutingSummaryV2ParamsEncrypted{radar.EmailRoutingSummaryV2ParamsEncryptedEncrypted}),
			Format:        cloudflare.F(radar.EmailRoutingSummaryV2ParamsFormatJson),
			IPVersion:     cloudflare.F([]radar.EmailRoutingSummaryV2ParamsIPVersion{radar.EmailRoutingSummaryV2ParamsIPVersionIPv4}),
			LimitPerGroup: cloudflare.F(int64(10)),
			Name:          cloudflare.F([]string{"main_series"}),
			SPF:           cloudflare.F([]radar.EmailRoutingSummaryV2ParamsSPF{radar.EmailRoutingSummaryV2ParamsSPFPass}),
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

func TestEmailRoutingTimeseriesGroupsV2WithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.TimeseriesGroupsV2(
		context.TODO(),
		radar.EmailRoutingTimeseriesGroupsV2ParamsDimensionIPVersion,
		radar.EmailRoutingTimeseriesGroupsV2Params{
			AggInterval:   cloudflare.F(radar.EmailRoutingTimeseriesGroupsV2ParamsAggInterval1h),
			ARC:           cloudflare.F([]radar.EmailRoutingTimeseriesGroupsV2ParamsARC{radar.EmailRoutingTimeseriesGroupsV2ParamsARCPass}),
			DateEnd:       cloudflare.F([]time.Time{time.Now()}),
			DateRange:     cloudflare.F([]string{"7d"}),
			DateStart:     cloudflare.F([]time.Time{time.Now()}),
			DKIM:          cloudflare.F([]radar.EmailRoutingTimeseriesGroupsV2ParamsDKIM{radar.EmailRoutingTimeseriesGroupsV2ParamsDKIMPass}),
			DMARC:         cloudflare.F([]radar.EmailRoutingTimeseriesGroupsV2ParamsDMARC{radar.EmailRoutingTimeseriesGroupsV2ParamsDMARCPass}),
			Encrypted:     cloudflare.F([]radar.EmailRoutingTimeseriesGroupsV2ParamsEncrypted{radar.EmailRoutingTimeseriesGroupsV2ParamsEncryptedEncrypted}),
			Format:        cloudflare.F(radar.EmailRoutingTimeseriesGroupsV2ParamsFormatJson),
			IPVersion:     cloudflare.F([]radar.EmailRoutingTimeseriesGroupsV2ParamsIPVersion{radar.EmailRoutingTimeseriesGroupsV2ParamsIPVersionIPv4}),
			LimitPerGroup: cloudflare.F(int64(10)),
			Name:          cloudflare.F([]string{"main_series"}),
			SPF:           cloudflare.F([]radar.EmailRoutingTimeseriesGroupsV2ParamsSPF{radar.EmailRoutingTimeseriesGroupsV2ParamsSPFPass}),
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
