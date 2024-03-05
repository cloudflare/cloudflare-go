// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestRadarAttackLayer3SummaryBitrateWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Summary.Bitrate(context.TODO(), cloudflare.RadarAttackLayer3SummaryBitrateParams{
		Continent: cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarAttackLayer3SummaryBitrateParamsDateRange{cloudflare.RadarAttackLayer3SummaryBitrateParamsDateRange1d, cloudflare.RadarAttackLayer3SummaryBitrateParamsDateRange2d, cloudflare.RadarAttackLayer3SummaryBitrateParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction: cloudflare.F(cloudflare.RadarAttackLayer3SummaryBitrateParamsDirectionOrigin),
		Format:    cloudflare.F(cloudflare.RadarAttackLayer3SummaryBitrateParamsFormatJson),
		IPVersion: cloudflare.F([]cloudflare.RadarAttackLayer3SummaryBitrateParamsIPVersion{cloudflare.RadarAttackLayer3SummaryBitrateParamsIPVersionIPv4, cloudflare.RadarAttackLayer3SummaryBitrateParamsIPVersionIPv6}),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		Protocol:  cloudflare.F([]cloudflare.RadarAttackLayer3SummaryBitrateParamsProtocol{cloudflare.RadarAttackLayer3SummaryBitrateParamsProtocolUdp, cloudflare.RadarAttackLayer3SummaryBitrateParamsProtocolTcp, cloudflare.RadarAttackLayer3SummaryBitrateParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer3SummaryDurationWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Summary.Duration(context.TODO(), cloudflare.RadarAttackLayer3SummaryDurationParams{
		Continent: cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarAttackLayer3SummaryDurationParamsDateRange{cloudflare.RadarAttackLayer3SummaryDurationParamsDateRange1d, cloudflare.RadarAttackLayer3SummaryDurationParamsDateRange2d, cloudflare.RadarAttackLayer3SummaryDurationParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction: cloudflare.F(cloudflare.RadarAttackLayer3SummaryDurationParamsDirectionOrigin),
		Format:    cloudflare.F(cloudflare.RadarAttackLayer3SummaryDurationParamsFormatJson),
		IPVersion: cloudflare.F([]cloudflare.RadarAttackLayer3SummaryDurationParamsIPVersion{cloudflare.RadarAttackLayer3SummaryDurationParamsIPVersionIPv4, cloudflare.RadarAttackLayer3SummaryDurationParamsIPVersionIPv6}),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		Protocol:  cloudflare.F([]cloudflare.RadarAttackLayer3SummaryDurationParamsProtocol{cloudflare.RadarAttackLayer3SummaryDurationParamsProtocolUdp, cloudflare.RadarAttackLayer3SummaryDurationParamsProtocolTcp, cloudflare.RadarAttackLayer3SummaryDurationParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer3SummaryGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Summary.Get(context.TODO(), cloudflare.RadarAttackLayer3SummaryGetParams{
		ASN:       cloudflare.F([]string{"string", "string", "string"}),
		Continent: cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarAttackLayer3SummaryGetParamsDateRange{cloudflare.RadarAttackLayer3SummaryGetParamsDateRange1d, cloudflare.RadarAttackLayer3SummaryGetParamsDateRange2d, cloudflare.RadarAttackLayer3SummaryGetParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:    cloudflare.F(cloudflare.RadarAttackLayer3SummaryGetParamsFormatJson),
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

func TestRadarAttackLayer3SummaryIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Summary.IPVersion(context.TODO(), cloudflare.RadarAttackLayer3SummaryIPVersionParams{
		Continent: cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarAttackLayer3SummaryIPVersionParamsDateRange{cloudflare.RadarAttackLayer3SummaryIPVersionParamsDateRange1d, cloudflare.RadarAttackLayer3SummaryIPVersionParamsDateRange2d, cloudflare.RadarAttackLayer3SummaryIPVersionParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction: cloudflare.F(cloudflare.RadarAttackLayer3SummaryIPVersionParamsDirectionOrigin),
		Format:    cloudflare.F(cloudflare.RadarAttackLayer3SummaryIPVersionParamsFormatJson),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		Protocol:  cloudflare.F([]cloudflare.RadarAttackLayer3SummaryIPVersionParamsProtocol{cloudflare.RadarAttackLayer3SummaryIPVersionParamsProtocolUdp, cloudflare.RadarAttackLayer3SummaryIPVersionParamsProtocolTcp, cloudflare.RadarAttackLayer3SummaryIPVersionParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer3SummaryProtocolWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Summary.Protocol(context.TODO(), cloudflare.RadarAttackLayer3SummaryProtocolParams{
		Continent: cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarAttackLayer3SummaryProtocolParamsDateRange{cloudflare.RadarAttackLayer3SummaryProtocolParamsDateRange1d, cloudflare.RadarAttackLayer3SummaryProtocolParamsDateRange2d, cloudflare.RadarAttackLayer3SummaryProtocolParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction: cloudflare.F(cloudflare.RadarAttackLayer3SummaryProtocolParamsDirectionOrigin),
		Format:    cloudflare.F(cloudflare.RadarAttackLayer3SummaryProtocolParamsFormatJson),
		IPVersion: cloudflare.F([]cloudflare.RadarAttackLayer3SummaryProtocolParamsIPVersion{cloudflare.RadarAttackLayer3SummaryProtocolParamsIPVersionIPv4, cloudflare.RadarAttackLayer3SummaryProtocolParamsIPVersionIPv6}),
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

func TestRadarAttackLayer3SummaryVectorWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Summary.Vector(context.TODO(), cloudflare.RadarAttackLayer3SummaryVectorParams{
		Continent: cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarAttackLayer3SummaryVectorParamsDateRange{cloudflare.RadarAttackLayer3SummaryVectorParamsDateRange1d, cloudflare.RadarAttackLayer3SummaryVectorParamsDateRange2d, cloudflare.RadarAttackLayer3SummaryVectorParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction: cloudflare.F(cloudflare.RadarAttackLayer3SummaryVectorParamsDirectionOrigin),
		Format:    cloudflare.F(cloudflare.RadarAttackLayer3SummaryVectorParamsFormatJson),
		IPVersion: cloudflare.F([]cloudflare.RadarAttackLayer3SummaryVectorParamsIPVersion{cloudflare.RadarAttackLayer3SummaryVectorParamsIPVersionIPv4, cloudflare.RadarAttackLayer3SummaryVectorParamsIPVersionIPv6}),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		Protocol:  cloudflare.F([]cloudflare.RadarAttackLayer3SummaryVectorParamsProtocol{cloudflare.RadarAttackLayer3SummaryVectorParamsProtocolUdp, cloudflare.RadarAttackLayer3SummaryVectorParamsProtocolTcp, cloudflare.RadarAttackLayer3SummaryVectorParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
