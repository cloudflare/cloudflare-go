// File generated from our OpenAPI spec by Stainless.

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

func TestAttackLayer3SummaryBitrateWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Radar.Attacks.Layer3.Summary.Bitrate(context.TODO(), radar.AttackLayer3SummaryBitrateParams{
		Continent: cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]radar.AttackLayer3SummaryBitrateParamsDateRange{radar.AttackLayer3SummaryBitrateParamsDateRange1d, radar.AttackLayer3SummaryBitrateParamsDateRange2d, radar.AttackLayer3SummaryBitrateParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction: cloudflare.F(radar.AttackLayer3SummaryBitrateParamsDirectionOrigin),
		Format:    cloudflare.F(radar.AttackLayer3SummaryBitrateParamsFormatJson),
		IPVersion: cloudflare.F([]radar.AttackLayer3SummaryBitrateParamsIPVersion{radar.AttackLayer3SummaryBitrateParamsIPVersionIPv4, radar.AttackLayer3SummaryBitrateParamsIPVersionIPv6}),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		Protocol:  cloudflare.F([]radar.AttackLayer3SummaryBitrateParamsProtocol{radar.AttackLayer3SummaryBitrateParamsProtocolUdp, radar.AttackLayer3SummaryBitrateParamsProtocolTcp, radar.AttackLayer3SummaryBitrateParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3SummaryDurationWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Radar.Attacks.Layer3.Summary.Duration(context.TODO(), radar.AttackLayer3SummaryDurationParams{
		Continent: cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]radar.AttackLayer3SummaryDurationParamsDateRange{radar.AttackLayer3SummaryDurationParamsDateRange1d, radar.AttackLayer3SummaryDurationParamsDateRange2d, radar.AttackLayer3SummaryDurationParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction: cloudflare.F(radar.AttackLayer3SummaryDurationParamsDirectionOrigin),
		Format:    cloudflare.F(radar.AttackLayer3SummaryDurationParamsFormatJson),
		IPVersion: cloudflare.F([]radar.AttackLayer3SummaryDurationParamsIPVersion{radar.AttackLayer3SummaryDurationParamsIPVersionIPv4, radar.AttackLayer3SummaryDurationParamsIPVersionIPv6}),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		Protocol:  cloudflare.F([]radar.AttackLayer3SummaryDurationParamsProtocol{radar.AttackLayer3SummaryDurationParamsProtocolUdp, radar.AttackLayer3SummaryDurationParamsProtocolTcp, radar.AttackLayer3SummaryDurationParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3SummaryGetWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Radar.Attacks.Layer3.Summary.Get(context.TODO(), radar.AttackLayer3SummaryGetParams{
		ASN:       cloudflare.F([]string{"string", "string", "string"}),
		Continent: cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]radar.AttackLayer3SummaryGetParamsDateRange{radar.AttackLayer3SummaryGetParamsDateRange1d, radar.AttackLayer3SummaryGetParamsDateRange2d, radar.AttackLayer3SummaryGetParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:    cloudflare.F(radar.AttackLayer3SummaryGetParamsFormatJson),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3SummaryIPVersionWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Radar.Attacks.Layer3.Summary.IPVersion(context.TODO(), radar.AttackLayer3SummaryIPVersionParams{
		Continent: cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]radar.AttackLayer3SummaryIPVersionParamsDateRange{radar.AttackLayer3SummaryIPVersionParamsDateRange1d, radar.AttackLayer3SummaryIPVersionParamsDateRange2d, radar.AttackLayer3SummaryIPVersionParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction: cloudflare.F(radar.AttackLayer3SummaryIPVersionParamsDirectionOrigin),
		Format:    cloudflare.F(radar.AttackLayer3SummaryIPVersionParamsFormatJson),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		Protocol:  cloudflare.F([]radar.AttackLayer3SummaryIPVersionParamsProtocol{radar.AttackLayer3SummaryIPVersionParamsProtocolUdp, radar.AttackLayer3SummaryIPVersionParamsProtocolTcp, radar.AttackLayer3SummaryIPVersionParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3SummaryProtocolWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Radar.Attacks.Layer3.Summary.Protocol(context.TODO(), radar.AttackLayer3SummaryProtocolParams{
		Continent: cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]radar.AttackLayer3SummaryProtocolParamsDateRange{radar.AttackLayer3SummaryProtocolParamsDateRange1d, radar.AttackLayer3SummaryProtocolParamsDateRange2d, radar.AttackLayer3SummaryProtocolParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction: cloudflare.F(radar.AttackLayer3SummaryProtocolParamsDirectionOrigin),
		Format:    cloudflare.F(radar.AttackLayer3SummaryProtocolParamsFormatJson),
		IPVersion: cloudflare.F([]radar.AttackLayer3SummaryProtocolParamsIPVersion{radar.AttackLayer3SummaryProtocolParamsIPVersionIPv4, radar.AttackLayer3SummaryProtocolParamsIPVersionIPv6}),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3SummaryVectorWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Radar.Attacks.Layer3.Summary.Vector(context.TODO(), radar.AttackLayer3SummaryVectorParams{
		Continent: cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]radar.AttackLayer3SummaryVectorParamsDateRange{radar.AttackLayer3SummaryVectorParamsDateRange1d, radar.AttackLayer3SummaryVectorParamsDateRange2d, radar.AttackLayer3SummaryVectorParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction: cloudflare.F(radar.AttackLayer3SummaryVectorParamsDirectionOrigin),
		Format:    cloudflare.F(radar.AttackLayer3SummaryVectorParamsFormatJson),
		IPVersion: cloudflare.F([]radar.AttackLayer3SummaryVectorParamsIPVersion{radar.AttackLayer3SummaryVectorParamsIPVersionIPv4, radar.AttackLayer3SummaryVectorParamsIPVersionIPv6}),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		Protocol:  cloudflare.F([]radar.AttackLayer3SummaryVectorParamsProtocol{radar.AttackLayer3SummaryVectorParamsProtocolUdp, radar.AttackLayer3SummaryVectorParamsProtocolTcp, radar.AttackLayer3SummaryVectorParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
