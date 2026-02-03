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

func TestDNSTimeseriesGroupCacheHitWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.TimeseriesGroups.CacheHit(context.TODO(), radar.DNSTimeseriesGroupCacheHitParams{
		AggInterval:  cloudflare.F(radar.DNSTimeseriesGroupCacheHitParamsAggInterval1h),
		ASN:          cloudflare.F([]string{"string"}),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		Format:       cloudflare.F(radar.DNSTimeseriesGroupCacheHitParamsFormatJson),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"main_series"}),
		Nodata:       cloudflare.F([]bool{true}),
		Protocol:     cloudflare.F([]radar.DNSTimeseriesGroupCacheHitParamsProtocol{radar.DNSTimeseriesGroupCacheHitParamsProtocolUdp}),
		QueryType:    cloudflare.F([]radar.DNSTimeseriesGroupCacheHitParamsQueryType{radar.DNSTimeseriesGroupCacheHitParamsQueryTypeA}),
		ResponseCode: cloudflare.F([]radar.DNSTimeseriesGroupCacheHitParamsResponseCode{radar.DNSTimeseriesGroupCacheHitParamsResponseCodeNoerror}),
		TLD:          cloudflare.F([]string{"com"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDNSTimeseriesGroupDNSSECWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.TimeseriesGroups.DNSSEC(context.TODO(), radar.DNSTimeseriesGroupDNSSECParams{
		AggInterval:  cloudflare.F(radar.DNSTimeseriesGroupDNSSECParamsAggInterval1h),
		ASN:          cloudflare.F([]string{"string"}),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		Format:       cloudflare.F(radar.DNSTimeseriesGroupDNSSECParamsFormatJson),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"main_series"}),
		Nodata:       cloudflare.F([]bool{true}),
		Protocol:     cloudflare.F([]radar.DNSTimeseriesGroupDNSSECParamsProtocol{radar.DNSTimeseriesGroupDNSSECParamsProtocolUdp}),
		QueryType:    cloudflare.F([]radar.DNSTimeseriesGroupDNSSECParamsQueryType{radar.DNSTimeseriesGroupDNSSECParamsQueryTypeA}),
		ResponseCode: cloudflare.F([]radar.DNSTimeseriesGroupDNSSECParamsResponseCode{radar.DNSTimeseriesGroupDNSSECParamsResponseCodeNoerror}),
		TLD:          cloudflare.F([]string{"com"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDNSTimeseriesGroupDNSSECAwareWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.TimeseriesGroups.DNSSECAware(context.TODO(), radar.DNSTimeseriesGroupDNSSECAwareParams{
		AggInterval:  cloudflare.F(radar.DNSTimeseriesGroupDNSSECAwareParamsAggInterval1h),
		ASN:          cloudflare.F([]string{"string"}),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		Format:       cloudflare.F(radar.DNSTimeseriesGroupDNSSECAwareParamsFormatJson),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"main_series"}),
		Nodata:       cloudflare.F([]bool{true}),
		Protocol:     cloudflare.F([]radar.DNSTimeseriesGroupDNSSECAwareParamsProtocol{radar.DNSTimeseriesGroupDNSSECAwareParamsProtocolUdp}),
		QueryType:    cloudflare.F([]radar.DNSTimeseriesGroupDNSSECAwareParamsQueryType{radar.DNSTimeseriesGroupDNSSECAwareParamsQueryTypeA}),
		ResponseCode: cloudflare.F([]radar.DNSTimeseriesGroupDNSSECAwareParamsResponseCode{radar.DNSTimeseriesGroupDNSSECAwareParamsResponseCodeNoerror}),
		TLD:          cloudflare.F([]string{"com"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDNSTimeseriesGroupDNSSECE2EWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.TimeseriesGroups.DNSSECE2E(context.TODO(), radar.DNSTimeseriesGroupDNSSECE2EParams{
		AggInterval:  cloudflare.F(radar.DNSTimeseriesGroupDnssece2EParamsAggInterval1h),
		ASN:          cloudflare.F([]string{"string"}),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		Format:       cloudflare.F(radar.DNSTimeseriesGroupDnssece2EParamsFormatJson),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"main_series"}),
		Nodata:       cloudflare.F([]bool{true}),
		Protocol:     cloudflare.F([]radar.DNSTimeseriesGroupDnssece2EParamsProtocol{radar.DNSTimeseriesGroupDnssece2EParamsProtocolUdp}),
		QueryType:    cloudflare.F([]radar.DNSTimeseriesGroupDnssece2EParamsQueryType{radar.DNSTimeseriesGroupDnssece2EParamsQueryTypeA}),
		ResponseCode: cloudflare.F([]radar.DNSTimeseriesGroupDnssece2EParamsResponseCode{radar.DNSTimeseriesGroupDnssece2EParamsResponseCodeNoerror}),
		TLD:          cloudflare.F([]string{"com"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDNSTimeseriesGroupIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.TimeseriesGroups.IPVersion(context.TODO(), radar.DNSTimeseriesGroupIPVersionParams{
		AggInterval:  cloudflare.F(radar.DNSTimeseriesGroupIPVersionParamsAggInterval1h),
		ASN:          cloudflare.F([]string{"string"}),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		Format:       cloudflare.F(radar.DNSTimeseriesGroupIPVersionParamsFormatJson),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"main_series"}),
		Nodata:       cloudflare.F([]bool{true}),
		Protocol:     cloudflare.F([]radar.DNSTimeseriesGroupIPVersionParamsProtocol{radar.DNSTimeseriesGroupIPVersionParamsProtocolUdp}),
		QueryType:    cloudflare.F([]radar.DNSTimeseriesGroupIPVersionParamsQueryType{radar.DNSTimeseriesGroupIPVersionParamsQueryTypeA}),
		ResponseCode: cloudflare.F([]radar.DNSTimeseriesGroupIPVersionParamsResponseCode{radar.DNSTimeseriesGroupIPVersionParamsResponseCodeNoerror}),
		TLD:          cloudflare.F([]string{"com"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDNSTimeseriesGroupMatchingAnswerWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.TimeseriesGroups.MatchingAnswer(context.TODO(), radar.DNSTimeseriesGroupMatchingAnswerParams{
		AggInterval:  cloudflare.F(radar.DNSTimeseriesGroupMatchingAnswerParamsAggInterval1h),
		ASN:          cloudflare.F([]string{"string"}),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		Format:       cloudflare.F(radar.DNSTimeseriesGroupMatchingAnswerParamsFormatJson),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"main_series"}),
		Nodata:       cloudflare.F([]bool{true}),
		Protocol:     cloudflare.F([]radar.DNSTimeseriesGroupMatchingAnswerParamsProtocol{radar.DNSTimeseriesGroupMatchingAnswerParamsProtocolUdp}),
		QueryType:    cloudflare.F([]radar.DNSTimeseriesGroupMatchingAnswerParamsQueryType{radar.DNSTimeseriesGroupMatchingAnswerParamsQueryTypeA}),
		ResponseCode: cloudflare.F([]radar.DNSTimeseriesGroupMatchingAnswerParamsResponseCode{radar.DNSTimeseriesGroupMatchingAnswerParamsResponseCodeNoerror}),
		TLD:          cloudflare.F([]string{"com"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDNSTimeseriesGroupProtocolWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.TimeseriesGroups.Protocol(context.TODO(), radar.DNSTimeseriesGroupProtocolParams{
		AggInterval:  cloudflare.F(radar.DNSTimeseriesGroupProtocolParamsAggInterval1h),
		ASN:          cloudflare.F([]string{"string"}),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		Format:       cloudflare.F(radar.DNSTimeseriesGroupProtocolParamsFormatJson),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"main_series"}),
		Nodata:       cloudflare.F([]bool{true}),
		QueryType:    cloudflare.F([]radar.DNSTimeseriesGroupProtocolParamsQueryType{radar.DNSTimeseriesGroupProtocolParamsQueryTypeA}),
		ResponseCode: cloudflare.F([]radar.DNSTimeseriesGroupProtocolParamsResponseCode{radar.DNSTimeseriesGroupProtocolParamsResponseCodeNoerror}),
		TLD:          cloudflare.F([]string{"com"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDNSTimeseriesGroupQueryTypeWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.TimeseriesGroups.QueryType(context.TODO(), radar.DNSTimeseriesGroupQueryTypeParams{
		AggInterval:   cloudflare.F(radar.DNSTimeseriesGroupQueryTypeParamsAggInterval1h),
		ASN:           cloudflare.F([]string{"string"}),
		Continent:     cloudflare.F([]string{"string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now()}),
		DateRange:     cloudflare.F([]string{"7d"}),
		DateStart:     cloudflare.F([]time.Time{time.Now()}),
		Format:        cloudflare.F(radar.DNSTimeseriesGroupQueryTypeParamsFormatJson),
		LimitPerGroup: cloudflare.F(int64(10)),
		Location:      cloudflare.F([]string{"string"}),
		Name:          cloudflare.F([]string{"main_series"}),
		Nodata:        cloudflare.F([]bool{true}),
		Protocol:      cloudflare.F([]radar.DNSTimeseriesGroupQueryTypeParamsProtocol{radar.DNSTimeseriesGroupQueryTypeParamsProtocolUdp}),
		ResponseCode:  cloudflare.F([]radar.DNSTimeseriesGroupQueryTypeParamsResponseCode{radar.DNSTimeseriesGroupQueryTypeParamsResponseCodeNoerror}),
		TLD:           cloudflare.F([]string{"com"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDNSTimeseriesGroupResponseCodeWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.TimeseriesGroups.ResponseCode(context.TODO(), radar.DNSTimeseriesGroupResponseCodeParams{
		AggInterval:   cloudflare.F(radar.DNSTimeseriesGroupResponseCodeParamsAggInterval1h),
		ASN:           cloudflare.F([]string{"string"}),
		Continent:     cloudflare.F([]string{"string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now()}),
		DateRange:     cloudflare.F([]string{"7d"}),
		DateStart:     cloudflare.F([]time.Time{time.Now()}),
		Format:        cloudflare.F(radar.DNSTimeseriesGroupResponseCodeParamsFormatJson),
		LimitPerGroup: cloudflare.F(int64(10)),
		Location:      cloudflare.F([]string{"string"}),
		Name:          cloudflare.F([]string{"main_series"}),
		Nodata:        cloudflare.F([]bool{true}),
		Protocol:      cloudflare.F([]radar.DNSTimeseriesGroupResponseCodeParamsProtocol{radar.DNSTimeseriesGroupResponseCodeParamsProtocolUdp}),
		QueryType:     cloudflare.F([]radar.DNSTimeseriesGroupResponseCodeParamsQueryType{radar.DNSTimeseriesGroupResponseCodeParamsQueryTypeA}),
		TLD:           cloudflare.F([]string{"com"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDNSTimeseriesGroupResponseTTLWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.TimeseriesGroups.ResponseTTL(context.TODO(), radar.DNSTimeseriesGroupResponseTTLParams{
		AggInterval:  cloudflare.F(radar.DNSTimeseriesGroupResponseTTLParamsAggInterval1h),
		ASN:          cloudflare.F([]string{"string"}),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		Format:       cloudflare.F(radar.DNSTimeseriesGroupResponseTTLParamsFormatJson),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"main_series"}),
		Nodata:       cloudflare.F([]bool{true}),
		Protocol:     cloudflare.F([]radar.DNSTimeseriesGroupResponseTTLParamsProtocol{radar.DNSTimeseriesGroupResponseTTLParamsProtocolUdp}),
		QueryType:    cloudflare.F([]radar.DNSTimeseriesGroupResponseTTLParamsQueryType{radar.DNSTimeseriesGroupResponseTTLParamsQueryTypeA}),
		ResponseCode: cloudflare.F([]radar.DNSTimeseriesGroupResponseTTLParamsResponseCode{radar.DNSTimeseriesGroupResponseTTLParamsResponseCodeNoerror}),
		TLD:          cloudflare.F([]string{"com"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
