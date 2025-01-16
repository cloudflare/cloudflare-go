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

func TestAttackLayer7TimeseriesGroupGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.Get(context.TODO(), radar.AttackLayer7TimeseriesGroupGetParams{
		AggInterval: cloudflare.F(radar.AttackLayer7TimeseriesGroupGetParamsAggInterval15m),
		ASN:         cloudflare.F([]string{"string"}),
		Continent:   cloudflare.F([]string{"string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now()}),
		DateRange:   cloudflare.F([]string{"7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now()}),
		Format:      cloudflare.F(radar.AttackLayer7TimeseriesGroupGetParamsFormatJson),
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

func TestAttackLayer7TimeseriesGroupHTTPMethodWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.HTTPMethod(context.TODO(), radar.AttackLayer7TimeseriesGroupHTTPMethodParams{
		AggInterval:       cloudflare.F(radar.AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval15m),
		ASN:               cloudflare.F([]string{"string"}),
		Continent:         cloudflare.F([]string{"string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now()}),
		DateRange:         cloudflare.F([]string{"7d"}),
		DateStart:         cloudflare.F([]time.Time{time.Now()}),
		Format:            cloudflare.F(radar.AttackLayer7TimeseriesGroupHTTPMethodParamsFormatJson),
		HTTPVersion:       cloudflare.F([]radar.AttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersion{radar.AttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersionHttPv1}),
		IPVersion:         cloudflare.F([]radar.AttackLayer7TimeseriesGroupHTTPMethodParamsIPVersion{radar.AttackLayer7TimeseriesGroupHTTPMethodParamsIPVersionIPv4}),
		Location:          cloudflare.F([]string{"string"}),
		MitigationProduct: cloudflare.F([]radar.AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct{radar.AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductDDoS}),
		Name:              cloudflare.F([]string{"string"}),
		Normalization:     cloudflare.F(radar.AttackLayer7TimeseriesGroupHTTPMethodParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer7TimeseriesGroupHTTPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.HTTPVersion(context.TODO(), radar.AttackLayer7TimeseriesGroupHTTPVersionParams{
		AggInterval:       cloudflare.F(radar.AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval15m),
		ASN:               cloudflare.F([]string{"string"}),
		Continent:         cloudflare.F([]string{"string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now()}),
		DateRange:         cloudflare.F([]string{"7d"}),
		DateStart:         cloudflare.F([]time.Time{time.Now()}),
		Format:            cloudflare.F(radar.AttackLayer7TimeseriesGroupHTTPVersionParamsFormatJson),
		HTTPMethod:        cloudflare.F([]radar.AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod{radar.AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodGet}),
		IPVersion:         cloudflare.F([]radar.AttackLayer7TimeseriesGroupHTTPVersionParamsIPVersion{radar.AttackLayer7TimeseriesGroupHTTPVersionParamsIPVersionIPv4}),
		Location:          cloudflare.F([]string{"string"}),
		MitigationProduct: cloudflare.F([]radar.AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct{radar.AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductDDoS}),
		Name:              cloudflare.F([]string{"string"}),
		Normalization:     cloudflare.F(radar.AttackLayer7TimeseriesGroupHTTPVersionParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer7TimeseriesGroupIndustryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.Industry(context.TODO(), radar.AttackLayer7TimeseriesGroupIndustryParams{
		AggInterval:       cloudflare.F(radar.AttackLayer7TimeseriesGroupIndustryParamsAggInterval15m),
		ASN:               cloudflare.F([]string{"string"}),
		Continent:         cloudflare.F([]string{"string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now()}),
		DateRange:         cloudflare.F([]string{"7d"}),
		DateStart:         cloudflare.F([]time.Time{time.Now()}),
		Format:            cloudflare.F(radar.AttackLayer7TimeseriesGroupIndustryParamsFormatJson),
		HTTPMethod:        cloudflare.F([]radar.AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod{radar.AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodGet}),
		HTTPVersion:       cloudflare.F([]radar.AttackLayer7TimeseriesGroupIndustryParamsHTTPVersion{radar.AttackLayer7TimeseriesGroupIndustryParamsHTTPVersionHttPv1}),
		IPVersion:         cloudflare.F([]radar.AttackLayer7TimeseriesGroupIndustryParamsIPVersion{radar.AttackLayer7TimeseriesGroupIndustryParamsIPVersionIPv4}),
		LimitPerGroup:     cloudflare.F(int64(4)),
		Location:          cloudflare.F([]string{"string"}),
		MitigationProduct: cloudflare.F([]radar.AttackLayer7TimeseriesGroupIndustryParamsMitigationProduct{radar.AttackLayer7TimeseriesGroupIndustryParamsMitigationProductDDoS}),
		Name:              cloudflare.F([]string{"string"}),
		Normalization:     cloudflare.F(radar.AttackLayer7TimeseriesGroupIndustryParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer7TimeseriesGroupIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.IPVersion(context.TODO(), radar.AttackLayer7TimeseriesGroupIPVersionParams{
		AggInterval:       cloudflare.F(radar.AttackLayer7TimeseriesGroupIPVersionParamsAggInterval15m),
		ASN:               cloudflare.F([]string{"string"}),
		Continent:         cloudflare.F([]string{"string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now()}),
		DateRange:         cloudflare.F([]string{"7d"}),
		DateStart:         cloudflare.F([]time.Time{time.Now()}),
		Format:            cloudflare.F(radar.AttackLayer7TimeseriesGroupIPVersionParamsFormatJson),
		HTTPMethod:        cloudflare.F([]radar.AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod{radar.AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodGet}),
		HTTPVersion:       cloudflare.F([]radar.AttackLayer7TimeseriesGroupIPVersionParamsHTTPVersion{radar.AttackLayer7TimeseriesGroupIPVersionParamsHTTPVersionHttPv1}),
		Location:          cloudflare.F([]string{"string"}),
		MitigationProduct: cloudflare.F([]radar.AttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct{radar.AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductDDoS}),
		Name:              cloudflare.F([]string{"string"}),
		Normalization:     cloudflare.F(radar.AttackLayer7TimeseriesGroupIPVersionParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer7TimeseriesGroupManagedRulesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.ManagedRules(context.TODO(), radar.AttackLayer7TimeseriesGroupManagedRulesParams{
		AggInterval:       cloudflare.F(radar.AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval15m),
		ASN:               cloudflare.F([]string{"string"}),
		Continent:         cloudflare.F([]string{"string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now()}),
		DateRange:         cloudflare.F([]string{"7d"}),
		DateStart:         cloudflare.F([]time.Time{time.Now()}),
		Format:            cloudflare.F(radar.AttackLayer7TimeseriesGroupManagedRulesParamsFormatJson),
		HTTPMethod:        cloudflare.F([]radar.AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod{radar.AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodGet}),
		HTTPVersion:       cloudflare.F([]radar.AttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersion{radar.AttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersionHttPv1}),
		IPVersion:         cloudflare.F([]radar.AttackLayer7TimeseriesGroupManagedRulesParamsIPVersion{radar.AttackLayer7TimeseriesGroupManagedRulesParamsIPVersionIPv4}),
		Location:          cloudflare.F([]string{"string"}),
		MitigationProduct: cloudflare.F([]radar.AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct{radar.AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductDDoS}),
		Name:              cloudflare.F([]string{"string"}),
		Normalization:     cloudflare.F(radar.AttackLayer7TimeseriesGroupManagedRulesParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer7TimeseriesGroupMitigationProductWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.MitigationProduct(context.TODO(), radar.AttackLayer7TimeseriesGroupMitigationProductParams{
		AggInterval:   cloudflare.F(radar.AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval15m),
		ASN:           cloudflare.F([]string{"string"}),
		Continent:     cloudflare.F([]string{"string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now()}),
		DateRange:     cloudflare.F([]string{"7d"}),
		DateStart:     cloudflare.F([]time.Time{time.Now()}),
		Format:        cloudflare.F(radar.AttackLayer7TimeseriesGroupMitigationProductParamsFormatJson),
		HTTPMethod:    cloudflare.F([]radar.AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod{radar.AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodGet}),
		HTTPVersion:   cloudflare.F([]radar.AttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersion{radar.AttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersionHttPv1}),
		IPVersion:     cloudflare.F([]radar.AttackLayer7TimeseriesGroupMitigationProductParamsIPVersion{radar.AttackLayer7TimeseriesGroupMitigationProductParamsIPVersionIPv4}),
		Location:      cloudflare.F([]string{"string"}),
		Name:          cloudflare.F([]string{"string"}),
		Normalization: cloudflare.F(radar.AttackLayer7TimeseriesGroupMitigationProductParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer7TimeseriesGroupVerticalWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.Vertical(context.TODO(), radar.AttackLayer7TimeseriesGroupVerticalParams{
		AggInterval:       cloudflare.F(radar.AttackLayer7TimeseriesGroupVerticalParamsAggInterval15m),
		ASN:               cloudflare.F([]string{"string"}),
		Continent:         cloudflare.F([]string{"string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now()}),
		DateRange:         cloudflare.F([]string{"7d"}),
		DateStart:         cloudflare.F([]time.Time{time.Now()}),
		Format:            cloudflare.F(radar.AttackLayer7TimeseriesGroupVerticalParamsFormatJson),
		HTTPMethod:        cloudflare.F([]radar.AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod{radar.AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodGet}),
		HTTPVersion:       cloudflare.F([]radar.AttackLayer7TimeseriesGroupVerticalParamsHTTPVersion{radar.AttackLayer7TimeseriesGroupVerticalParamsHTTPVersionHttPv1}),
		IPVersion:         cloudflare.F([]radar.AttackLayer7TimeseriesGroupVerticalParamsIPVersion{radar.AttackLayer7TimeseriesGroupVerticalParamsIPVersionIPv4}),
		LimitPerGroup:     cloudflare.F(int64(4)),
		Location:          cloudflare.F([]string{"string"}),
		MitigationProduct: cloudflare.F([]radar.AttackLayer7TimeseriesGroupVerticalParamsMitigationProduct{radar.AttackLayer7TimeseriesGroupVerticalParamsMitigationProductDDoS}),
		Name:              cloudflare.F([]string{"string"}),
		Normalization:     cloudflare.F(radar.AttackLayer7TimeseriesGroupVerticalParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
