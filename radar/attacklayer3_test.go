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

func TestAttackLayer3SummaryV2WithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.SummaryV2(
		context.TODO(),
		radar.AttackLayer3SummaryV2ParamsDimensionProtocol,
		radar.AttackLayer3SummaryV2Params{
			Continent:     cloudflare.F([]string{"string"}),
			DateEnd:       cloudflare.F([]time.Time{time.Now()}),
			DateRange:     cloudflare.F([]string{"7d"}),
			DateStart:     cloudflare.F([]time.Time{time.Now()}),
			Direction:     cloudflare.F(radar.AttackLayer3SummaryV2ParamsDirectionOrigin),
			Format:        cloudflare.F(radar.AttackLayer3SummaryV2ParamsFormatJson),
			IPVersion:     cloudflare.F([]radar.AttackLayer3SummaryV2ParamsIPVersion{radar.AttackLayer3SummaryV2ParamsIPVersionIPv4}),
			LimitPerGroup: cloudflare.F(int64(10)),
			Location:      cloudflare.F([]string{"string"}),
			Name:          cloudflare.F([]string{"main_series"}),
			Protocol:      cloudflare.F([]radar.AttackLayer3SummaryV2ParamsProtocol{radar.AttackLayer3SummaryV2ParamsProtocolUdp}),
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

func TestAttackLayer3TimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Timeseries(context.TODO(), radar.AttackLayer3TimeseriesParams{
		AggInterval:   cloudflare.F(radar.AttackLayer3TimeseriesParamsAggInterval1h),
		ASN:           cloudflare.F([]string{"string"}),
		Continent:     cloudflare.F([]string{"string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now()}),
		DateRange:     cloudflare.F([]string{"7d"}),
		DateStart:     cloudflare.F([]time.Time{time.Now()}),
		Direction:     cloudflare.F(radar.AttackLayer3TimeseriesParamsDirectionOrigin),
		Format:        cloudflare.F(radar.AttackLayer3TimeseriesParamsFormatJson),
		IPVersion:     cloudflare.F([]radar.AttackLayer3TimeseriesParamsIPVersion{radar.AttackLayer3TimeseriesParamsIPVersionIPv4}),
		Location:      cloudflare.F([]string{"string"}),
		Metric:        cloudflare.F(radar.AttackLayer3TimeseriesParamsMetricBytes),
		Name:          cloudflare.F([]string{"main_series"}),
		Normalization: cloudflare.F(radar.AttackLayer3TimeseriesParamsNormalizationMin0Max),
		Protocol:      cloudflare.F([]radar.AttackLayer3TimeseriesParamsProtocol{radar.AttackLayer3TimeseriesParamsProtocolUdp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3TimeseriesGroupsV2WithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.TimeseriesGroupsV2(
		context.TODO(),
		radar.AttackLayer3TimeseriesGroupsV2ParamsDimensionProtocol,
		radar.AttackLayer3TimeseriesGroupsV2Params{
			AggInterval:   cloudflare.F(radar.AttackLayer3TimeseriesGroupsV2ParamsAggInterval1h),
			Continent:     cloudflare.F([]string{"string"}),
			DateEnd:       cloudflare.F([]time.Time{time.Now()}),
			DateRange:     cloudflare.F([]string{"7d"}),
			DateStart:     cloudflare.F([]time.Time{time.Now()}),
			Direction:     cloudflare.F(radar.AttackLayer3TimeseriesGroupsV2ParamsDirectionOrigin),
			Format:        cloudflare.F(radar.AttackLayer3TimeseriesGroupsV2ParamsFormatJson),
			IPVersion:     cloudflare.F([]radar.AttackLayer3TimeseriesGroupsV2ParamsIPVersion{radar.AttackLayer3TimeseriesGroupsV2ParamsIPVersionIPv4}),
			LimitPerGroup: cloudflare.F(int64(10)),
			Location:      cloudflare.F([]string{"string"}),
			Name:          cloudflare.F([]string{"main_series"}),
			Normalization: cloudflare.F(radar.AttackLayer3TimeseriesGroupsV2ParamsNormalizationPercentage),
			Protocol:      cloudflare.F([]radar.AttackLayer3TimeseriesGroupsV2ParamsProtocol{radar.AttackLayer3TimeseriesGroupsV2ParamsProtocolUdp}),
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
