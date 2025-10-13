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

func TestHTTPSummaryV2WithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.SummaryV2(
		context.TODO(),
		radar.HTTPSummaryV2ParamsDimensionAdm1,
		radar.HTTPSummaryV2Params{
			ASN:           cloudflare.F([]string{"string"}),
			BotClass:      cloudflare.F([]radar.HTTPSummaryV2ParamsBotClass{radar.HTTPSummaryV2ParamsBotClassLikelyAutomated}),
			Continent:     cloudflare.F([]string{"string"}),
			DateEnd:       cloudflare.F([]time.Time{time.Now()}),
			DateRange:     cloudflare.F([]string{"7d"}),
			DateStart:     cloudflare.F([]time.Time{time.Now()}),
			DeviceType:    cloudflare.F([]radar.HTTPSummaryV2ParamsDeviceType{radar.HTTPSummaryV2ParamsDeviceTypeDesktop}),
			Format:        cloudflare.F(radar.HTTPSummaryV2ParamsFormatJson),
			GeoID:         cloudflare.F([]string{"string"}),
			HTTPProtocol:  cloudflare.F([]radar.HTTPSummaryV2ParamsHTTPProtocol{radar.HTTPSummaryV2ParamsHTTPProtocolHTTP}),
			HTTPVersion:   cloudflare.F([]radar.HTTPSummaryV2ParamsHTTPVersion{radar.HTTPSummaryV2ParamsHTTPVersionHttPv1}),
			IPVersion:     cloudflare.F([]radar.HTTPSummaryV2ParamsIPVersion{radar.HTTPSummaryV2ParamsIPVersionIPv4}),
			LimitPerGroup: cloudflare.F(int64(10)),
			Location:      cloudflare.F([]string{"string"}),
			Name:          cloudflare.F([]string{"main_series"}),
			OS:            cloudflare.F([]radar.HTTPSummaryV2ParamsOS{radar.HTTPSummaryV2ParamsOSWindows}),
			TLSVersion:    cloudflare.F([]radar.HTTPSummaryV2ParamsTLSVersion{radar.HTTPSummaryV2ParamsTLSVersionTlSv1_0}),
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

func TestHTTPTimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Timeseries(context.TODO(), radar.HTTPTimeseriesParams{
		AggInterval:   cloudflare.F(radar.HTTPTimeseriesParamsAggInterval1h),
		ASN:           cloudflare.F([]string{"string"}),
		BotClass:      cloudflare.F([]radar.HTTPTimeseriesParamsBotClass{radar.HTTPTimeseriesParamsBotClassLikelyAutomated}),
		BrowserFamily: cloudflare.F([]radar.HTTPTimeseriesParamsBrowserFamily{radar.HTTPTimeseriesParamsBrowserFamilyChrome}),
		Continent:     cloudflare.F([]string{"string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now()}),
		DateRange:     cloudflare.F([]string{"7d"}),
		DateStart:     cloudflare.F([]time.Time{time.Now()}),
		DeviceType:    cloudflare.F([]radar.HTTPTimeseriesParamsDeviceType{radar.HTTPTimeseriesParamsDeviceTypeDesktop}),
		Format:        cloudflare.F(radar.HTTPTimeseriesParamsFormatJson),
		GeoID:         cloudflare.F([]string{"string"}),
		HTTPProtocol:  cloudflare.F([]radar.HTTPTimeseriesParamsHTTPProtocol{radar.HTTPTimeseriesParamsHTTPProtocolHTTP}),
		HTTPVersion:   cloudflare.F([]radar.HTTPTimeseriesParamsHTTPVersion{radar.HTTPTimeseriesParamsHTTPVersionHttPv1}),
		IPVersion:     cloudflare.F([]radar.HTTPTimeseriesParamsIPVersion{radar.HTTPTimeseriesParamsIPVersionIPv4}),
		Location:      cloudflare.F([]string{"string"}),
		Name:          cloudflare.F([]string{"main_series"}),
		Normalization: cloudflare.F(radar.HTTPTimeseriesParamsNormalizationMin0Max),
		OS:            cloudflare.F([]radar.HTTPTimeseriesParamsOS{radar.HTTPTimeseriesParamsOSWindows}),
		TLSVersion:    cloudflare.F([]radar.HTTPTimeseriesParamsTLSVersion{radar.HTTPTimeseriesParamsTLSVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPTimeseriesGroupsV2WithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroupsV2(
		context.TODO(),
		radar.HTTPTimeseriesGroupsV2ParamsDimensionAdm1,
		radar.HTTPTimeseriesGroupsV2Params{
			AggInterval:   cloudflare.F(radar.HTTPTimeseriesGroupsV2ParamsAggInterval1h),
			ASN:           cloudflare.F([]string{"string"}),
			BotClass:      cloudflare.F([]radar.HTTPTimeseriesGroupsV2ParamsBotClass{radar.HTTPTimeseriesGroupsV2ParamsBotClassLikelyAutomated}),
			Continent:     cloudflare.F([]string{"string"}),
			DateEnd:       cloudflare.F([]time.Time{time.Now()}),
			DateRange:     cloudflare.F([]string{"7d"}),
			DateStart:     cloudflare.F([]time.Time{time.Now()}),
			DeviceType:    cloudflare.F([]radar.HTTPTimeseriesGroupsV2ParamsDeviceType{radar.HTTPTimeseriesGroupsV2ParamsDeviceTypeDesktop}),
			Format:        cloudflare.F(radar.HTTPTimeseriesGroupsV2ParamsFormatJson),
			GeoID:         cloudflare.F([]string{"string"}),
			HTTPProtocol:  cloudflare.F([]radar.HTTPTimeseriesGroupsV2ParamsHTTPProtocol{radar.HTTPTimeseriesGroupsV2ParamsHTTPProtocolHTTP}),
			HTTPVersion:   cloudflare.F([]radar.HTTPTimeseriesGroupsV2ParamsHTTPVersion{radar.HTTPTimeseriesGroupsV2ParamsHTTPVersionHttPv1}),
			IPVersion:     cloudflare.F([]radar.HTTPTimeseriesGroupsV2ParamsIPVersion{radar.HTTPTimeseriesGroupsV2ParamsIPVersionIPv4}),
			LimitPerGroup: cloudflare.F(int64(10)),
			Location:      cloudflare.F([]string{"string"}),
			Name:          cloudflare.F([]string{"main_series"}),
			Normalization: cloudflare.F(radar.HTTPTimeseriesGroupsV2ParamsNormalizationPercentage),
			OS:            cloudflare.F([]radar.HTTPTimeseriesGroupsV2ParamsOS{radar.HTTPTimeseriesGroupsV2ParamsOSWindows}),
			TLSVersion:    cloudflare.F([]radar.HTTPTimeseriesGroupsV2ParamsTLSVersion{radar.HTTPTimeseriesGroupsV2ParamsTLSVersionTlSv1_0}),
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
