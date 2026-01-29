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

func TestAS112SummaryV2WithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.AS112.SummaryV2(
		context.TODO(),
		radar.AS112SummaryV2ParamsDimensionDNSSEC,
		radar.AS112SummaryV2Params{
			Continent:     cloudflare.F([]string{"string"}),
			DateEnd:       cloudflare.F([]time.Time{time.Now()}),
			DateRange:     cloudflare.F([]string{"7d"}),
			DateStart:     cloudflare.F([]time.Time{time.Now()}),
			Format:        cloudflare.F(radar.AS112SummaryV2ParamsFormatJson),
			LimitPerGroup: cloudflare.F(int64(10)),
			Location:      cloudflare.F([]string{"string"}),
			Name:          cloudflare.F([]string{"main_series"}),
			Protocol:      cloudflare.F([]radar.AS112SummaryV2ParamsProtocol{radar.AS112SummaryV2ParamsProtocolUdp}),
			QueryType:     cloudflare.F([]radar.AS112SummaryV2ParamsQueryType{radar.AS112SummaryV2ParamsQueryTypeA}),
			ResponseCode:  cloudflare.F([]radar.AS112SummaryV2ParamsResponseCode{radar.AS112SummaryV2ParamsResponseCodeNoerror}),
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

func TestAS112TimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.AS112.Timeseries(context.TODO(), radar.AS112TimeseriesParams{
		AggInterval:  cloudflare.F(radar.AS112TimeseriesParamsAggInterval1h),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		Format:       cloudflare.F(radar.AS112TimeseriesParamsFormatJson),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"main_series"}),
		Protocol:     cloudflare.F([]radar.AS112TimeseriesParamsProtocol{radar.AS112TimeseriesParamsProtocolUdp}),
		QueryType:    cloudflare.F([]radar.AS112TimeseriesParamsQueryType{radar.AS112TimeseriesParamsQueryTypeA}),
		ResponseCode: cloudflare.F([]radar.AS112TimeseriesParamsResponseCode{radar.AS112TimeseriesParamsResponseCodeNoerror}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAS112TimeseriesGroupsV2WithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.AS112.TimeseriesGroupsV2(
		context.TODO(),
		radar.AS112TimeseriesGroupsV2ParamsDimensionDNSSEC,
		radar.AS112TimeseriesGroupsV2Params{
			AggInterval:   cloudflare.F(radar.AS112TimeseriesGroupsV2ParamsAggInterval1h),
			Continent:     cloudflare.F([]string{"string"}),
			DateEnd:       cloudflare.F([]time.Time{time.Now()}),
			DateRange:     cloudflare.F([]string{"7d"}),
			DateStart:     cloudflare.F([]time.Time{time.Now()}),
			Format:        cloudflare.F(radar.AS112TimeseriesGroupsV2ParamsFormatJson),
			LimitPerGroup: cloudflare.F(int64(10)),
			Location:      cloudflare.F([]string{"string"}),
			Name:          cloudflare.F([]string{"main_series"}),
			Protocol:      cloudflare.F([]radar.AS112TimeseriesGroupsV2ParamsProtocol{radar.AS112TimeseriesGroupsV2ParamsProtocolUdp}),
			QueryType:     cloudflare.F([]radar.AS112TimeseriesGroupsV2ParamsQueryType{radar.AS112TimeseriesGroupsV2ParamsQueryTypeA}),
			ResponseCode:  cloudflare.F([]radar.AS112TimeseriesGroupsV2ParamsResponseCode{radar.AS112TimeseriesGroupsV2ParamsResponseCodeNoerror}),
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
