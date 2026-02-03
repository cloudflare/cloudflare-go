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

func TestDNSTopAsesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.Top.Ases(context.TODO(), radar.DNSTopAsesParams{
		ASN:            cloudflare.F([]string{"string"}),
		CacheHit:       cloudflare.F([]bool{true}),
		Continent:      cloudflare.F([]string{"string"}),
		DateEnd:        cloudflare.F([]time.Time{time.Now()}),
		DateRange:      cloudflare.F([]string{"7d"}),
		DateStart:      cloudflare.F([]time.Time{time.Now()}),
		DNSSEC:         cloudflare.F([]radar.DNSTopAsesParamsDNSSEC{radar.DNSTopAsesParamsDNSSECInvalid}),
		DNSSECAware:    cloudflare.F([]radar.DNSTopAsesParamsDNSSECAware{radar.DNSTopAsesParamsDNSSECAwareSupported}),
		DNSSECE2E:      cloudflare.F([]bool{true}),
		Domain:         cloudflare.F([]string{"google.com"}),
		Format:         cloudflare.F(radar.DNSTopAsesParamsFormatJson),
		IPVersion:      cloudflare.F([]radar.DNSTopAsesParamsIPVersion{radar.DNSTopAsesParamsIPVersionIPv4}),
		Limit:          cloudflare.F(int64(1)),
		Location:       cloudflare.F([]string{"string"}),
		MatchingAnswer: cloudflare.F([]bool{true}),
		Name:           cloudflare.F([]string{"main_series"}),
		Nodata:         cloudflare.F([]bool{true}),
		Protocol:       cloudflare.F([]radar.DNSTopAsesParamsProtocol{radar.DNSTopAsesParamsProtocolUdp}),
		QueryType:      cloudflare.F([]radar.DNSTopAsesParamsQueryType{radar.DNSTopAsesParamsQueryTypeA}),
		ResponseCode:   cloudflare.F([]radar.DNSTopAsesParamsResponseCode{radar.DNSTopAsesParamsResponseCodeNoerror}),
		ResponseTTL:    cloudflare.F([]radar.DNSTopAsesParamsResponseTTL{radar.DNSTopAsesParamsResponseTTLLte1M}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDNSTopLocationsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.Top.Locations(context.TODO(), radar.DNSTopLocationsParams{
		ASN:            cloudflare.F([]string{"string"}),
		CacheHit:       cloudflare.F([]bool{true}),
		Continent:      cloudflare.F([]string{"string"}),
		DateEnd:        cloudflare.F([]time.Time{time.Now()}),
		DateRange:      cloudflare.F([]string{"7d"}),
		DateStart:      cloudflare.F([]time.Time{time.Now()}),
		DNSSEC:         cloudflare.F([]radar.DNSTopLocationsParamsDNSSEC{radar.DNSTopLocationsParamsDNSSECInvalid}),
		DNSSECAware:    cloudflare.F([]radar.DNSTopLocationsParamsDNSSECAware{radar.DNSTopLocationsParamsDNSSECAwareSupported}),
		DNSSECE2E:      cloudflare.F([]bool{true}),
		Domain:         cloudflare.F([]string{"google.com"}),
		Format:         cloudflare.F(radar.DNSTopLocationsParamsFormatJson),
		IPVersion:      cloudflare.F([]radar.DNSTopLocationsParamsIPVersion{radar.DNSTopLocationsParamsIPVersionIPv4}),
		Limit:          cloudflare.F(int64(1)),
		Location:       cloudflare.F([]string{"string"}),
		MatchingAnswer: cloudflare.F([]bool{true}),
		Name:           cloudflare.F([]string{"main_series"}),
		Nodata:         cloudflare.F([]bool{true}),
		Protocol:       cloudflare.F([]radar.DNSTopLocationsParamsProtocol{radar.DNSTopLocationsParamsProtocolUdp}),
		QueryType:      cloudflare.F([]radar.DNSTopLocationsParamsQueryType{radar.DNSTopLocationsParamsQueryTypeA}),
		ResponseCode:   cloudflare.F([]radar.DNSTopLocationsParamsResponseCode{radar.DNSTopLocationsParamsResponseCodeNoerror}),
		ResponseTTL:    cloudflare.F([]radar.DNSTopLocationsParamsResponseTTL{radar.DNSTopLocationsParamsResponseTTLLte1M}),
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
