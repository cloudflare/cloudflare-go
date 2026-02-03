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

func TestDNSSummaryV2WithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.SummaryV2(
		context.TODO(),
		radar.DNSSummaryV2ParamsDimensionIPVersion,
		radar.DNSSummaryV2Params{
			ASN:            cloudflare.F([]string{"string"}),
			CacheHit:       cloudflare.F([]bool{true}),
			Continent:      cloudflare.F([]string{"string"}),
			DateEnd:        cloudflare.F([]time.Time{time.Now()}),
			DateRange:      cloudflare.F([]string{"7d"}),
			DateStart:      cloudflare.F([]time.Time{time.Now()}),
			DNSSEC:         cloudflare.F([]radar.DNSSummaryV2ParamsDNSSEC{radar.DNSSummaryV2ParamsDNSSECInvalid}),
			DNSSECAware:    cloudflare.F([]radar.DNSSummaryV2ParamsDNSSECAware{radar.DNSSummaryV2ParamsDNSSECAwareSupported}),
			DNSSECE2E:      cloudflare.F([]bool{true}),
			Format:         cloudflare.F(radar.DNSSummaryV2ParamsFormatJson),
			IPVersion:      cloudflare.F([]radar.DNSSummaryV2ParamsIPVersion{radar.DNSSummaryV2ParamsIPVersionIPv4}),
			LimitPerGroup:  cloudflare.F(int64(10)),
			Location:       cloudflare.F([]string{"string"}),
			MatchingAnswer: cloudflare.F([]bool{true}),
			Name:           cloudflare.F([]string{"main_series"}),
			Nodata:         cloudflare.F([]bool{true}),
			Protocol:       cloudflare.F([]radar.DNSSummaryV2ParamsProtocol{radar.DNSSummaryV2ParamsProtocolUdp}),
			QueryType:      cloudflare.F([]radar.DNSSummaryV2ParamsQueryType{radar.DNSSummaryV2ParamsQueryTypeA}),
			ResponseCode:   cloudflare.F([]radar.DNSSummaryV2ParamsResponseCode{radar.DNSSummaryV2ParamsResponseCodeNoerror}),
			ResponseTTL:    cloudflare.F([]radar.DNSSummaryV2ParamsResponseTTL{radar.DNSSummaryV2ParamsResponseTTLLte1M}),
			TLD:            cloudflare.F([]string{"com"}),
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

func TestDNSTimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.Timeseries(context.TODO(), radar.DNSTimeseriesParams{
		AggInterval:    cloudflare.F(radar.DNSTimeseriesParamsAggInterval1h),
		ASN:            cloudflare.F([]string{"string"}),
		CacheHit:       cloudflare.F([]bool{true}),
		Continent:      cloudflare.F([]string{"string"}),
		DateEnd:        cloudflare.F([]time.Time{time.Now()}),
		DateRange:      cloudflare.F([]string{"7d"}),
		DateStart:      cloudflare.F([]time.Time{time.Now()}),
		DNSSEC:         cloudflare.F([]radar.DNSTimeseriesParamsDNSSEC{radar.DNSTimeseriesParamsDNSSECInvalid}),
		DNSSECAware:    cloudflare.F([]radar.DNSTimeseriesParamsDNSSECAware{radar.DNSTimeseriesParamsDNSSECAwareSupported}),
		DNSSECE2E:      cloudflare.F([]bool{true}),
		Format:         cloudflare.F(radar.DNSTimeseriesParamsFormatJson),
		IPVersion:      cloudflare.F([]radar.DNSTimeseriesParamsIPVersion{radar.DNSTimeseriesParamsIPVersionIPv4}),
		Location:       cloudflare.F([]string{"string"}),
		MatchingAnswer: cloudflare.F([]bool{true}),
		Name:           cloudflare.F([]string{"main_series"}),
		Nodata:         cloudflare.F([]bool{true}),
		Protocol:       cloudflare.F([]radar.DNSTimeseriesParamsProtocol{radar.DNSTimeseriesParamsProtocolUdp}),
		QueryType:      cloudflare.F([]radar.DNSTimeseriesParamsQueryType{radar.DNSTimeseriesParamsQueryTypeA}),
		ResponseCode:   cloudflare.F([]radar.DNSTimeseriesParamsResponseCode{radar.DNSTimeseriesParamsResponseCodeNoerror}),
		ResponseTTL:    cloudflare.F([]radar.DNSTimeseriesParamsResponseTTL{radar.DNSTimeseriesParamsResponseTTLLte1M}),
		TLD:            cloudflare.F([]string{"com"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDNSTimeseriesGroupsV2WithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.TimeseriesGroupsV2(
		context.TODO(),
		radar.DNSTimeseriesGroupsV2ParamsDimensionIPVersion,
		radar.DNSTimeseriesGroupsV2Params{
			AggInterval:    cloudflare.F(radar.DNSTimeseriesGroupsV2ParamsAggInterval1h),
			ASN:            cloudflare.F([]string{"string"}),
			CacheHit:       cloudflare.F([]bool{true}),
			Continent:      cloudflare.F([]string{"string"}),
			DateEnd:        cloudflare.F([]time.Time{time.Now()}),
			DateRange:      cloudflare.F([]string{"7d"}),
			DateStart:      cloudflare.F([]time.Time{time.Now()}),
			DNSSEC:         cloudflare.F([]radar.DNSTimeseriesGroupsV2ParamsDNSSEC{radar.DNSTimeseriesGroupsV2ParamsDNSSECInvalid}),
			DNSSECAware:    cloudflare.F([]radar.DNSTimeseriesGroupsV2ParamsDNSSECAware{radar.DNSTimeseriesGroupsV2ParamsDNSSECAwareSupported}),
			DNSSECE2E:      cloudflare.F([]bool{true}),
			Format:         cloudflare.F(radar.DNSTimeseriesGroupsV2ParamsFormatJson),
			IPVersion:      cloudflare.F([]radar.DNSTimeseriesGroupsV2ParamsIPVersion{radar.DNSTimeseriesGroupsV2ParamsIPVersionIPv4}),
			LimitPerGroup:  cloudflare.F(int64(10)),
			Location:       cloudflare.F([]string{"string"}),
			MatchingAnswer: cloudflare.F([]bool{true}),
			Name:           cloudflare.F([]string{"main_series"}),
			Nodata:         cloudflare.F([]bool{true}),
			Normalization:  cloudflare.F(radar.DNSTimeseriesGroupsV2ParamsNormalizationPercentage),
			Protocol:       cloudflare.F([]radar.DNSTimeseriesGroupsV2ParamsProtocol{radar.DNSTimeseriesGroupsV2ParamsProtocolUdp}),
			QueryType:      cloudflare.F([]radar.DNSTimeseriesGroupsV2ParamsQueryType{radar.DNSTimeseriesGroupsV2ParamsQueryTypeA}),
			ResponseCode:   cloudflare.F([]radar.DNSTimeseriesGroupsV2ParamsResponseCode{radar.DNSTimeseriesGroupsV2ParamsResponseCodeNoerror}),
			ResponseTTL:    cloudflare.F([]radar.DNSTimeseriesGroupsV2ParamsResponseTTL{radar.DNSTimeseriesGroupsV2ParamsResponseTTLLte1M}),
			TLD:            cloudflare.F([]string{"com"}),
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
