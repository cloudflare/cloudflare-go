// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/radar"
)

func TestAttackLayer3TimeseriesGroupBitrateWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.TimeseriesGroups.Bitrate(context.TODO(), radar.AttackLayer3TimeseriesGroupBitrateParams{
		AggInterval:   cloudflare.F(radar.AttackLayer3TimeseriesGroupBitrateParamsAggInterval15m),
		Continent:     cloudflare.F([]string{"string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now()}),
		DateRange:     cloudflare.F([]string{"7d"}),
		DateStart:     cloudflare.F([]time.Time{time.Now()}),
		Direction:     cloudflare.F(radar.AttackLayer3TimeseriesGroupBitrateParamsDirectionOrigin),
		Format:        cloudflare.F(radar.AttackLayer3TimeseriesGroupBitrateParamsFormatJson),
		IPVersion:     cloudflare.F([]radar.AttackLayer3TimeseriesGroupBitrateParamsIPVersion{radar.AttackLayer3TimeseriesGroupBitrateParamsIPVersionIPv4}),
		Location:      cloudflare.F([]string{"string"}),
		Name:          cloudflare.F([]string{"string"}),
		Normalization: cloudflare.F(radar.AttackLayer3TimeseriesGroupBitrateParamsNormalizationPercentage),
		Protocol:      cloudflare.F([]radar.AttackLayer3TimeseriesGroupBitrateParamsProtocol{radar.AttackLayer3TimeseriesGroupBitrateParamsProtocolUdp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3TimeseriesGroupDurationWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.TimeseriesGroups.Duration(context.TODO(), radar.AttackLayer3TimeseriesGroupDurationParams{
		AggInterval:   cloudflare.F(radar.AttackLayer3TimeseriesGroupDurationParamsAggInterval15m),
		Continent:     cloudflare.F([]string{"string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now()}),
		DateRange:     cloudflare.F([]string{"7d"}),
		DateStart:     cloudflare.F([]time.Time{time.Now()}),
		Direction:     cloudflare.F(radar.AttackLayer3TimeseriesGroupDurationParamsDirectionOrigin),
		Format:        cloudflare.F(radar.AttackLayer3TimeseriesGroupDurationParamsFormatJson),
		IPVersion:     cloudflare.F([]radar.AttackLayer3TimeseriesGroupDurationParamsIPVersion{radar.AttackLayer3TimeseriesGroupDurationParamsIPVersionIPv4}),
		Location:      cloudflare.F([]string{"string"}),
		Name:          cloudflare.F([]string{"string"}),
		Normalization: cloudflare.F(radar.AttackLayer3TimeseriesGroupDurationParamsNormalizationPercentage),
		Protocol:      cloudflare.F([]radar.AttackLayer3TimeseriesGroupDurationParamsProtocol{radar.AttackLayer3TimeseriesGroupDurationParamsProtocolUdp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3TimeseriesGroupGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.TimeseriesGroups.Get(context.TODO(), radar.AttackLayer3TimeseriesGroupGetParams{
		AggInterval: cloudflare.F(radar.AttackLayer3TimeseriesGroupGetParamsAggInterval15m),
		ASN:         cloudflare.F([]string{"string"}),
		Continent:   cloudflare.F([]string{"string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now()}),
		DateRange:   cloudflare.F([]string{"7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now()}),
		Format:      cloudflare.F(radar.AttackLayer3TimeseriesGroupGetParamsFormatJson),
		Location:    cloudflare.F([]string{"string"}),
		Name:        cloudflare.F([]string{"string"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3TimeseriesGroupIndustryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.TimeseriesGroups.Industry(context.TODO(), radar.AttackLayer3TimeseriesGroupIndustryParams{
		AggInterval:   cloudflare.F(radar.AttackLayer3TimeseriesGroupIndustryParamsAggInterval15m),
		Continent:     cloudflare.F([]string{"string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now()}),
		DateRange:     cloudflare.F([]string{"7d"}),
		DateStart:     cloudflare.F([]time.Time{time.Now()}),
		Direction:     cloudflare.F(radar.AttackLayer3TimeseriesGroupIndustryParamsDirectionOrigin),
		Format:        cloudflare.F(radar.AttackLayer3TimeseriesGroupIndustryParamsFormatJson),
		IPVersion:     cloudflare.F([]radar.AttackLayer3TimeseriesGroupIndustryParamsIPVersion{radar.AttackLayer3TimeseriesGroupIndustryParamsIPVersionIPv4}),
		LimitPerGroup: cloudflare.F(int64(4)),
		Location:      cloudflare.F([]string{"string"}),
		Name:          cloudflare.F([]string{"string"}),
		Normalization: cloudflare.F(radar.AttackLayer3TimeseriesGroupIndustryParamsNormalizationPercentage),
		Protocol:      cloudflare.F([]radar.AttackLayer3TimeseriesGroupIndustryParamsProtocol{radar.AttackLayer3TimeseriesGroupIndustryParamsProtocolUdp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3TimeseriesGroupIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.TimeseriesGroups.IPVersion(context.TODO(), radar.AttackLayer3TimeseriesGroupIPVersionParams{
		AggInterval:   cloudflare.F(radar.AttackLayer3TimeseriesGroupIPVersionParamsAggInterval15m),
		Continent:     cloudflare.F([]string{"string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now()}),
		DateRange:     cloudflare.F([]string{"7d"}),
		DateStart:     cloudflare.F([]time.Time{time.Now()}),
		Direction:     cloudflare.F(radar.AttackLayer3TimeseriesGroupIPVersionParamsDirectionOrigin),
		Format:        cloudflare.F(radar.AttackLayer3TimeseriesGroupIPVersionParamsFormatJson),
		Location:      cloudflare.F([]string{"string"}),
		Name:          cloudflare.F([]string{"string"}),
		Normalization: cloudflare.F(radar.AttackLayer3TimeseriesGroupIPVersionParamsNormalizationPercentage),
		Protocol:      cloudflare.F([]radar.AttackLayer3TimeseriesGroupIPVersionParamsProtocol{radar.AttackLayer3TimeseriesGroupIPVersionParamsProtocolUdp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3TimeseriesGroupProtocolWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.TimeseriesGroups.Protocol(context.TODO(), radar.AttackLayer3TimeseriesGroupProtocolParams{
		AggInterval:   cloudflare.F(radar.AttackLayer3TimeseriesGroupProtocolParamsAggInterval15m),
		Continent:     cloudflare.F([]string{"string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now()}),
		DateRange:     cloudflare.F([]string{"7d"}),
		DateStart:     cloudflare.F([]time.Time{time.Now()}),
		Direction:     cloudflare.F(radar.AttackLayer3TimeseriesGroupProtocolParamsDirectionOrigin),
		Format:        cloudflare.F(radar.AttackLayer3TimeseriesGroupProtocolParamsFormatJson),
		IPVersion:     cloudflare.F([]radar.AttackLayer3TimeseriesGroupProtocolParamsIPVersion{radar.AttackLayer3TimeseriesGroupProtocolParamsIPVersionIPv4}),
		Location:      cloudflare.F([]string{"string"}),
		Name:          cloudflare.F([]string{"string"}),
		Normalization: cloudflare.F(radar.AttackLayer3TimeseriesGroupProtocolParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3TimeseriesGroupVectorWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.TimeseriesGroups.Vector(context.TODO(), radar.AttackLayer3TimeseriesGroupVectorParams{
		AggInterval:   cloudflare.F(radar.AttackLayer3TimeseriesGroupVectorParamsAggInterval15m),
		Continent:     cloudflare.F([]string{"string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now()}),
		DateRange:     cloudflare.F([]string{"7d"}),
		DateStart:     cloudflare.F([]time.Time{time.Now()}),
		Direction:     cloudflare.F(radar.AttackLayer3TimeseriesGroupVectorParamsDirectionOrigin),
		Format:        cloudflare.F(radar.AttackLayer3TimeseriesGroupVectorParamsFormatJson),
		IPVersion:     cloudflare.F([]radar.AttackLayer3TimeseriesGroupVectorParamsIPVersion{radar.AttackLayer3TimeseriesGroupVectorParamsIPVersionIPv4}),
		LimitPerGroup: cloudflare.F(int64(4)),
		Location:      cloudflare.F([]string{"string"}),
		Name:          cloudflare.F([]string{"string"}),
		Normalization: cloudflare.F(radar.AttackLayer3TimeseriesGroupVectorParamsNormalizationPercentage),
		Protocol:      cloudflare.F([]radar.AttackLayer3TimeseriesGroupVectorParamsProtocol{radar.AttackLayer3TimeseriesGroupVectorParamsProtocolUdp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3TimeseriesGroupVerticalWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.TimeseriesGroups.Vertical(context.TODO(), radar.AttackLayer3TimeseriesGroupVerticalParams{
		AggInterval:   cloudflare.F(radar.AttackLayer3TimeseriesGroupVerticalParamsAggInterval15m),
		Continent:     cloudflare.F([]string{"string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now()}),
		DateRange:     cloudflare.F([]string{"7d"}),
		DateStart:     cloudflare.F([]time.Time{time.Now()}),
		Direction:     cloudflare.F(radar.AttackLayer3TimeseriesGroupVerticalParamsDirectionOrigin),
		Format:        cloudflare.F(radar.AttackLayer3TimeseriesGroupVerticalParamsFormatJson),
		IPVersion:     cloudflare.F([]radar.AttackLayer3TimeseriesGroupVerticalParamsIPVersion{radar.AttackLayer3TimeseriesGroupVerticalParamsIPVersionIPv4}),
		LimitPerGroup: cloudflare.F(int64(4)),
		Location:      cloudflare.F([]string{"string"}),
		Name:          cloudflare.F([]string{"string"}),
		Normalization: cloudflare.F(radar.AttackLayer3TimeseriesGroupVerticalParamsNormalizationPercentage),
		Protocol:      cloudflare.F([]radar.AttackLayer3TimeseriesGroupVerticalParamsProtocol{radar.AttackLayer3TimeseriesGroupVerticalParamsProtocolUdp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
