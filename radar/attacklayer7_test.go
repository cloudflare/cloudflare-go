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

func TestAttackLayer7SummaryV2WithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.SummaryV2(
		context.TODO(),
		radar.AttackLayer7SummaryV2ParamsDimensionHTTPMethod,
		radar.AttackLayer7SummaryV2Params{
			ASN:               cloudflare.F([]string{"string"}),
			Continent:         cloudflare.F([]string{"string"}),
			DateEnd:           cloudflare.F([]time.Time{time.Now()}),
			DateRange:         cloudflare.F([]string{"7d"}),
			DateStart:         cloudflare.F([]time.Time{time.Now()}),
			Format:            cloudflare.F(radar.AttackLayer7SummaryV2ParamsFormatJson),
			HTTPMethod:        cloudflare.F([]radar.AttackLayer7SummaryV2ParamsHTTPMethod{radar.AttackLayer7SummaryV2ParamsHTTPMethodGet}),
			HTTPVersion:       cloudflare.F([]radar.AttackLayer7SummaryV2ParamsHTTPVersion{radar.AttackLayer7SummaryV2ParamsHTTPVersionHttPv1}),
			IPVersion:         cloudflare.F([]radar.AttackLayer7SummaryV2ParamsIPVersion{radar.AttackLayer7SummaryV2ParamsIPVersionIPv4}),
			LimitPerGroup:     cloudflare.F(int64(10)),
			Location:          cloudflare.F([]string{"string"}),
			MitigationProduct: cloudflare.F([]radar.AttackLayer7SummaryV2ParamsMitigationProduct{radar.AttackLayer7SummaryV2ParamsMitigationProductDDoS}),
			Name:              cloudflare.F([]string{"main_series"}),
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

func TestAttackLayer7TimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Timeseries(context.TODO(), radar.AttackLayer7TimeseriesParams{
		AggInterval:       cloudflare.F(radar.AttackLayer7TimeseriesParamsAggInterval1h),
		ASN:               cloudflare.F([]string{"string"}),
		Continent:         cloudflare.F([]string{"string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now()}),
		DateRange:         cloudflare.F([]string{"7d"}),
		DateStart:         cloudflare.F([]time.Time{time.Now()}),
		Format:            cloudflare.F(radar.AttackLayer7TimeseriesParamsFormatJson),
		HTTPMethod:        cloudflare.F([]radar.AttackLayer7TimeseriesParamsHTTPMethod{radar.AttackLayer7TimeseriesParamsHTTPMethodGet}),
		HTTPVersion:       cloudflare.F([]radar.AttackLayer7TimeseriesParamsHTTPVersion{radar.AttackLayer7TimeseriesParamsHTTPVersionHttPv1}),
		IPVersion:         cloudflare.F([]radar.AttackLayer7TimeseriesParamsIPVersion{radar.AttackLayer7TimeseriesParamsIPVersionIPv4}),
		Location:          cloudflare.F([]string{"string"}),
		MitigationProduct: cloudflare.F([]radar.AttackLayer7TimeseriesParamsMitigationProduct{radar.AttackLayer7TimeseriesParamsMitigationProductDDoS}),
		Name:              cloudflare.F([]string{"main_series"}),
		Normalization:     cloudflare.F(radar.AttackLayer7TimeseriesParamsNormalizationMin0Max),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer7TimeseriesGroupsV2WithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroupsV2(
		context.TODO(),
		radar.AttackLayer7TimeseriesGroupsV2ParamsDimensionHTTPMethod,
		radar.AttackLayer7TimeseriesGroupsV2Params{
			AggInterval:       cloudflare.F(radar.AttackLayer7TimeseriesGroupsV2ParamsAggInterval1h),
			ASN:               cloudflare.F([]string{"string"}),
			Continent:         cloudflare.F([]string{"string"}),
			DateEnd:           cloudflare.F([]time.Time{time.Now()}),
			DateRange:         cloudflare.F([]string{"7d"}),
			DateStart:         cloudflare.F([]time.Time{time.Now()}),
			Format:            cloudflare.F(radar.AttackLayer7TimeseriesGroupsV2ParamsFormatJson),
			HTTPMethod:        cloudflare.F([]radar.AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod{radar.AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodGet}),
			HTTPVersion:       cloudflare.F([]radar.AttackLayer7TimeseriesGroupsV2ParamsHTTPVersion{radar.AttackLayer7TimeseriesGroupsV2ParamsHTTPVersionHttPv1}),
			IPVersion:         cloudflare.F([]radar.AttackLayer7TimeseriesGroupsV2ParamsIPVersion{radar.AttackLayer7TimeseriesGroupsV2ParamsIPVersionIPv4}),
			LimitPerGroup:     cloudflare.F(int64(10)),
			Location:          cloudflare.F([]string{"string"}),
			MitigationProduct: cloudflare.F([]radar.AttackLayer7TimeseriesGroupsV2ParamsMitigationProduct{radar.AttackLayer7TimeseriesGroupsV2ParamsMitigationProductDDoS}),
			Name:              cloudflare.F([]string{"main_series"}),
			Normalization:     cloudflare.F(radar.AttackLayer7TimeseriesGroupsV2ParamsNormalizationPercentage),
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
