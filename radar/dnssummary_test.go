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

func TestDNSSummaryCacheHitWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.Summary.CacheHit(context.TODO(), radar.DNSSummaryCacheHitParams{
		ASN:          cloudflare.F([]string{"string"}),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		Format:       cloudflare.F(radar.DNSSummaryCacheHitParamsFormatJson),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"main_series"}),
		Nodata:       cloudflare.F([]bool{true}),
		Protocol:     cloudflare.F([]radar.DNSSummaryCacheHitParamsProtocol{radar.DNSSummaryCacheHitParamsProtocolUdp}),
		QueryType:    cloudflare.F([]radar.DNSSummaryCacheHitParamsQueryType{radar.DNSSummaryCacheHitParamsQueryTypeA}),
		ResponseCode: cloudflare.F([]radar.DNSSummaryCacheHitParamsResponseCode{radar.DNSSummaryCacheHitParamsResponseCodeNoerror}),
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

func TestDNSSummaryDNSSECWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.Summary.DNSSEC(context.TODO(), radar.DNSSummaryDNSSECParams{
		ASN:          cloudflare.F([]string{"string"}),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		Format:       cloudflare.F(radar.DNSSummaryDNSSECParamsFormatJson),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"main_series"}),
		Nodata:       cloudflare.F([]bool{true}),
		Protocol:     cloudflare.F([]radar.DNSSummaryDNSSECParamsProtocol{radar.DNSSummaryDNSSECParamsProtocolUdp}),
		QueryType:    cloudflare.F([]radar.DNSSummaryDNSSECParamsQueryType{radar.DNSSummaryDNSSECParamsQueryTypeA}),
		ResponseCode: cloudflare.F([]radar.DNSSummaryDNSSECParamsResponseCode{radar.DNSSummaryDNSSECParamsResponseCodeNoerror}),
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

func TestDNSSummaryDNSSECAwareWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.Summary.DNSSECAware(context.TODO(), radar.DNSSummaryDNSSECAwareParams{
		ASN:          cloudflare.F([]string{"string"}),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		Format:       cloudflare.F(radar.DNSSummaryDNSSECAwareParamsFormatJson),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"main_series"}),
		Nodata:       cloudflare.F([]bool{true}),
		Protocol:     cloudflare.F([]radar.DNSSummaryDNSSECAwareParamsProtocol{radar.DNSSummaryDNSSECAwareParamsProtocolUdp}),
		QueryType:    cloudflare.F([]radar.DNSSummaryDNSSECAwareParamsQueryType{radar.DNSSummaryDNSSECAwareParamsQueryTypeA}),
		ResponseCode: cloudflare.F([]radar.DNSSummaryDNSSECAwareParamsResponseCode{radar.DNSSummaryDNSSECAwareParamsResponseCodeNoerror}),
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

func TestDNSSummaryDNSSECE2EWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.Summary.DNSSECE2E(context.TODO(), radar.DNSSummaryDNSSECE2EParams{
		ASN:          cloudflare.F([]string{"string"}),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		Format:       cloudflare.F(radar.DNSSummaryDnssece2EParamsFormatJson),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"main_series"}),
		Nodata:       cloudflare.F([]bool{true}),
		Protocol:     cloudflare.F([]radar.DNSSummaryDnssece2EParamsProtocol{radar.DNSSummaryDnssece2EParamsProtocolUdp}),
		QueryType:    cloudflare.F([]radar.DNSSummaryDnssece2EParamsQueryType{radar.DNSSummaryDnssece2EParamsQueryTypeA}),
		ResponseCode: cloudflare.F([]radar.DNSSummaryDnssece2EParamsResponseCode{radar.DNSSummaryDnssece2EParamsResponseCodeNoerror}),
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

func TestDNSSummaryIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.Summary.IPVersion(context.TODO(), radar.DNSSummaryIPVersionParams{
		ASN:          cloudflare.F([]string{"string"}),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		Format:       cloudflare.F(radar.DNSSummaryIPVersionParamsFormatJson),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"main_series"}),
		Nodata:       cloudflare.F([]bool{true}),
		Protocol:     cloudflare.F([]radar.DNSSummaryIPVersionParamsProtocol{radar.DNSSummaryIPVersionParamsProtocolUdp}),
		QueryType:    cloudflare.F([]radar.DNSSummaryIPVersionParamsQueryType{radar.DNSSummaryIPVersionParamsQueryTypeA}),
		ResponseCode: cloudflare.F([]radar.DNSSummaryIPVersionParamsResponseCode{radar.DNSSummaryIPVersionParamsResponseCodeNoerror}),
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

func TestDNSSummaryMatchingAnswerWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.Summary.MatchingAnswer(context.TODO(), radar.DNSSummaryMatchingAnswerParams{
		ASN:          cloudflare.F([]string{"string"}),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		Format:       cloudflare.F(radar.DNSSummaryMatchingAnswerParamsFormatJson),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"main_series"}),
		Nodata:       cloudflare.F([]bool{true}),
		Protocol:     cloudflare.F([]radar.DNSSummaryMatchingAnswerParamsProtocol{radar.DNSSummaryMatchingAnswerParamsProtocolUdp}),
		QueryType:    cloudflare.F([]radar.DNSSummaryMatchingAnswerParamsQueryType{radar.DNSSummaryMatchingAnswerParamsQueryTypeA}),
		ResponseCode: cloudflare.F([]radar.DNSSummaryMatchingAnswerParamsResponseCode{radar.DNSSummaryMatchingAnswerParamsResponseCodeNoerror}),
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

func TestDNSSummaryProtocolWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.Summary.Protocol(context.TODO(), radar.DNSSummaryProtocolParams{
		ASN:          cloudflare.F([]string{"string"}),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		Format:       cloudflare.F(radar.DNSSummaryProtocolParamsFormatJson),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"main_series"}),
		Nodata:       cloudflare.F([]bool{true}),
		QueryType:    cloudflare.F([]radar.DNSSummaryProtocolParamsQueryType{radar.DNSSummaryProtocolParamsQueryTypeA}),
		ResponseCode: cloudflare.F([]radar.DNSSummaryProtocolParamsResponseCode{radar.DNSSummaryProtocolParamsResponseCodeNoerror}),
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

func TestDNSSummaryQueryTypeWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.Summary.QueryType(context.TODO(), radar.DNSSummaryQueryTypeParams{
		ASN:           cloudflare.F([]string{"string"}),
		Continent:     cloudflare.F([]string{"string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now()}),
		DateRange:     cloudflare.F([]string{"7d"}),
		DateStart:     cloudflare.F([]time.Time{time.Now()}),
		Format:        cloudflare.F(radar.DNSSummaryQueryTypeParamsFormatJson),
		LimitPerGroup: cloudflare.F(int64(10)),
		Location:      cloudflare.F([]string{"string"}),
		Name:          cloudflare.F([]string{"main_series"}),
		Nodata:        cloudflare.F([]bool{true}),
		Protocol:      cloudflare.F([]radar.DNSSummaryQueryTypeParamsProtocol{radar.DNSSummaryQueryTypeParamsProtocolUdp}),
		ResponseCode:  cloudflare.F([]radar.DNSSummaryQueryTypeParamsResponseCode{radar.DNSSummaryQueryTypeParamsResponseCodeNoerror}),
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

func TestDNSSummaryResponseCodeWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.Summary.ResponseCode(context.TODO(), radar.DNSSummaryResponseCodeParams{
		ASN:           cloudflare.F([]string{"string"}),
		Continent:     cloudflare.F([]string{"string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now()}),
		DateRange:     cloudflare.F([]string{"7d"}),
		DateStart:     cloudflare.F([]time.Time{time.Now()}),
		Format:        cloudflare.F(radar.DNSSummaryResponseCodeParamsFormatJson),
		LimitPerGroup: cloudflare.F(int64(10)),
		Location:      cloudflare.F([]string{"string"}),
		Name:          cloudflare.F([]string{"main_series"}),
		Nodata:        cloudflare.F([]bool{true}),
		Protocol:      cloudflare.F([]radar.DNSSummaryResponseCodeParamsProtocol{radar.DNSSummaryResponseCodeParamsProtocolUdp}),
		QueryType:     cloudflare.F([]radar.DNSSummaryResponseCodeParamsQueryType{radar.DNSSummaryResponseCodeParamsQueryTypeA}),
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

func TestDNSSummaryResponseTTLWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.Summary.ResponseTTL(context.TODO(), radar.DNSSummaryResponseTTLParams{
		ASN:          cloudflare.F([]string{"string"}),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		Format:       cloudflare.F(radar.DNSSummaryResponseTTLParamsFormatJson),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"main_series"}),
		Nodata:       cloudflare.F([]bool{true}),
		Protocol:     cloudflare.F([]radar.DNSSummaryResponseTTLParamsProtocol{radar.DNSSummaryResponseTTLParamsProtocolUdp}),
		QueryType:    cloudflare.F([]radar.DNSSummaryResponseTTLParamsQueryType{radar.DNSSummaryResponseTTLParamsQueryTypeA}),
		ResponseCode: cloudflare.F([]radar.DNSSummaryResponseTTLParamsResponseCode{radar.DNSSummaryResponseTTLParamsResponseCodeNoerror}),
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
