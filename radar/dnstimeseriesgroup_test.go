// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/radar"
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
		AggInterval: cloudflare.F(radar.DNSTimeseriesGroupCacheHitParamsAggInterval1h),
		ASN:         cloudflare.F([]string{"string", "string", "string"}),
		Continent:   cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:      cloudflare.F(radar.DNSTimeseriesGroupCacheHitParamsFormatJson),
		Location:    cloudflare.F([]string{"string", "string", "string"}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		Tld:         cloudflare.F([]string{"string", "string", "string"}),
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
		AggInterval: cloudflare.F(radar.DNSTimeseriesGroupDNSSECParamsAggInterval1h),
		ASN:         cloudflare.F([]string{"string", "string", "string"}),
		Continent:   cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:      cloudflare.F(radar.DNSTimeseriesGroupDNSSECParamsFormatJson),
		Location:    cloudflare.F([]string{"string", "string", "string"}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		Tld:         cloudflare.F([]string{"string", "string", "string"}),
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
		AggInterval: cloudflare.F(radar.DNSTimeseriesGroupDNSSECAwareParamsAggInterval1h),
		ASN:         cloudflare.F([]string{"string", "string", "string"}),
		Continent:   cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:      cloudflare.F(radar.DNSTimeseriesGroupDNSSECAwareParamsFormatJson),
		Location:    cloudflare.F([]string{"string", "string", "string"}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		Tld:         cloudflare.F([]string{"string", "string", "string"}),
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
		AggInterval: cloudflare.F(radar.DNSTimeseriesGroupDnssece2EParamsAggInterval1h),
		ASN:         cloudflare.F([]string{"string", "string", "string"}),
		Continent:   cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:      cloudflare.F(radar.DNSTimeseriesGroupDnssece2EParamsFormatJson),
		Location:    cloudflare.F([]string{"string", "string", "string"}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		Tld:         cloudflare.F([]string{"string", "string", "string"}),
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
		AggInterval: cloudflare.F(radar.DNSTimeseriesGroupIPVersionParamsAggInterval1h),
		ASN:         cloudflare.F([]string{"string", "string", "string"}),
		Continent:   cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:      cloudflare.F(radar.DNSTimeseriesGroupIPVersionParamsFormatJson),
		Location:    cloudflare.F([]string{"string", "string", "string"}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		Tld:         cloudflare.F([]string{"string", "string", "string"}),
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
		AggInterval: cloudflare.F(radar.DNSTimeseriesGroupMatchingAnswerParamsAggInterval1h),
		ASN:         cloudflare.F([]string{"string", "string", "string"}),
		Continent:   cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:      cloudflare.F(radar.DNSTimeseriesGroupMatchingAnswerParamsFormatJson),
		Location:    cloudflare.F([]string{"string", "string", "string"}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		Tld:         cloudflare.F([]string{"string", "string", "string"}),
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
		AggInterval: cloudflare.F(radar.DNSTimeseriesGroupProtocolParamsAggInterval1h),
		ASN:         cloudflare.F([]string{"string", "string", "string"}),
		Continent:   cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:      cloudflare.F(radar.DNSTimeseriesGroupProtocolParamsFormatJson),
		Location:    cloudflare.F([]string{"string", "string", "string"}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		Tld:         cloudflare.F([]string{"string", "string", "string"}),
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
		AggInterval: cloudflare.F(radar.DNSTimeseriesGroupQueryTypeParamsAggInterval1h),
		ASN:         cloudflare.F([]string{"string", "string", "string"}),
		Continent:   cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:      cloudflare.F(radar.DNSTimeseriesGroupQueryTypeParamsFormatJson),
		Location:    cloudflare.F([]string{"string", "string", "string"}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		Tld:         cloudflare.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDNSTimeseriesGroupResponseCodesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.TimeseriesGroups.ResponseCodes(context.TODO(), radar.DNSTimeseriesGroupResponseCodesParams{
		AggInterval: cloudflare.F(radar.DNSTimeseriesGroupResponseCodesParamsAggInterval1h),
		ASN:         cloudflare.F([]string{"string", "string", "string"}),
		Continent:   cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:      cloudflare.F(radar.DNSTimeseriesGroupResponseCodesParamsFormatJson),
		Location:    cloudflare.F([]string{"string", "string", "string"}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		Tld:         cloudflare.F([]string{"string", "string", "string"}),
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
		AggInterval: cloudflare.F(radar.DNSTimeseriesGroupResponseTTLParamsAggInterval1h),
		ASN:         cloudflare.F([]string{"string", "string", "string"}),
		Continent:   cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:      cloudflare.F(radar.DNSTimeseriesGroupResponseTTLParamsFormatJson),
		Location:    cloudflare.F([]string{"string", "string", "string"}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		Tld:         cloudflare.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
