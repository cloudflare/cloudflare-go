// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/radar"
)

func TestAttackLayer3SummaryBitrateWithOptionalParams(t *testing.T) {
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
		Continent: cloudflare.F([]string{"string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now()}),
		DateRange: cloudflare.F([]string{"7d"}),
		DateStart: cloudflare.F([]time.Time{time.Now()}),
		Direction: cloudflare.F(radar.AttackLayer3SummaryBitrateParamsDirectionOrigin),
		Format:    cloudflare.F(radar.AttackLayer3SummaryBitrateParamsFormatJson),
		IPVersion: cloudflare.F([]radar.AttackLayer3SummaryBitrateParamsIPVersion{radar.AttackLayer3SummaryBitrateParamsIPVersionIPv4}),
		Location:  cloudflare.F([]string{"string"}),
		Name:      cloudflare.F([]string{"main_series"}),
		Protocol:  cloudflare.F([]radar.AttackLayer3SummaryBitrateParamsProtocol{radar.AttackLayer3SummaryBitrateParamsProtocolUdp}),
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
		Continent: cloudflare.F([]string{"string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now()}),
		DateRange: cloudflare.F([]string{"7d"}),
		DateStart: cloudflare.F([]time.Time{time.Now()}),
		Direction: cloudflare.F(radar.AttackLayer3SummaryDurationParamsDirectionOrigin),
		Format:    cloudflare.F(radar.AttackLayer3SummaryDurationParamsFormatJson),
		IPVersion: cloudflare.F([]radar.AttackLayer3SummaryDurationParamsIPVersion{radar.AttackLayer3SummaryDurationParamsIPVersionIPv4}),
		Location:  cloudflare.F([]string{"string"}),
		Name:      cloudflare.F([]string{"main_series"}),
		Protocol:  cloudflare.F([]radar.AttackLayer3SummaryDurationParamsProtocol{radar.AttackLayer3SummaryDurationParamsProtocolUdp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3SummaryIndustryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Summary.Industry(context.TODO(), radar.AttackLayer3SummaryIndustryParams{
		Continent:     cloudflare.F([]string{"string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now()}),
		DateRange:     cloudflare.F([]string{"7d"}),
		DateStart:     cloudflare.F([]time.Time{time.Now()}),
		Direction:     cloudflare.F(radar.AttackLayer3SummaryIndustryParamsDirectionOrigin),
		Format:        cloudflare.F(radar.AttackLayer3SummaryIndustryParamsFormatJson),
		IPVersion:     cloudflare.F([]radar.AttackLayer3SummaryIndustryParamsIPVersion{radar.AttackLayer3SummaryIndustryParamsIPVersionIPv4}),
		LimitPerGroup: cloudflare.F(int64(10)),
		Location:      cloudflare.F([]string{"string"}),
		Name:          cloudflare.F([]string{"main_series"}),
		Protocol:      cloudflare.F([]radar.AttackLayer3SummaryIndustryParamsProtocol{radar.AttackLayer3SummaryIndustryParamsProtocolUdp}),
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
		Continent: cloudflare.F([]string{"string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now()}),
		DateRange: cloudflare.F([]string{"7d"}),
		DateStart: cloudflare.F([]time.Time{time.Now()}),
		Direction: cloudflare.F(radar.AttackLayer3SummaryIPVersionParamsDirectionOrigin),
		Format:    cloudflare.F(radar.AttackLayer3SummaryIPVersionParamsFormatJson),
		Location:  cloudflare.F([]string{"string"}),
		Name:      cloudflare.F([]string{"main_series"}),
		Protocol:  cloudflare.F([]radar.AttackLayer3SummaryIPVersionParamsProtocol{radar.AttackLayer3SummaryIPVersionParamsProtocolUdp}),
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
		Continent: cloudflare.F([]string{"string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now()}),
		DateRange: cloudflare.F([]string{"7d"}),
		DateStart: cloudflare.F([]time.Time{time.Now()}),
		Direction: cloudflare.F(radar.AttackLayer3SummaryProtocolParamsDirectionOrigin),
		Format:    cloudflare.F(radar.AttackLayer3SummaryProtocolParamsFormatJson),
		IPVersion: cloudflare.F([]radar.AttackLayer3SummaryProtocolParamsIPVersion{radar.AttackLayer3SummaryProtocolParamsIPVersionIPv4}),
		Location:  cloudflare.F([]string{"string"}),
		Name:      cloudflare.F([]string{"main_series"}),
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
		Continent:     cloudflare.F([]string{"string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now()}),
		DateRange:     cloudflare.F([]string{"7d"}),
		DateStart:     cloudflare.F([]time.Time{time.Now()}),
		Direction:     cloudflare.F(radar.AttackLayer3SummaryVectorParamsDirectionOrigin),
		Format:        cloudflare.F(radar.AttackLayer3SummaryVectorParamsFormatJson),
		IPVersion:     cloudflare.F([]radar.AttackLayer3SummaryVectorParamsIPVersion{radar.AttackLayer3SummaryVectorParamsIPVersionIPv4}),
		LimitPerGroup: cloudflare.F(int64(10)),
		Location:      cloudflare.F([]string{"string"}),
		Name:          cloudflare.F([]string{"main_series"}),
		Protocol:      cloudflare.F([]radar.AttackLayer3SummaryVectorParamsProtocol{radar.AttackLayer3SummaryVectorParamsProtocolUdp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3SummaryVerticalWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Summary.Vertical(context.TODO(), radar.AttackLayer3SummaryVerticalParams{
		Continent:     cloudflare.F([]string{"string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now()}),
		DateRange:     cloudflare.F([]string{"7d"}),
		DateStart:     cloudflare.F([]time.Time{time.Now()}),
		Direction:     cloudflare.F(radar.AttackLayer3SummaryVerticalParamsDirectionOrigin),
		Format:        cloudflare.F(radar.AttackLayer3SummaryVerticalParamsFormatJson),
		IPVersion:     cloudflare.F([]radar.AttackLayer3SummaryVerticalParamsIPVersion{radar.AttackLayer3SummaryVerticalParamsIPVersionIPv4}),
		LimitPerGroup: cloudflare.F(int64(10)),
		Location:      cloudflare.F([]string{"string"}),
		Name:          cloudflare.F([]string{"main_series"}),
		Protocol:      cloudflare.F([]radar.AttackLayer3SummaryVerticalParamsProtocol{radar.AttackLayer3SummaryVerticalParamsProtocolUdp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
