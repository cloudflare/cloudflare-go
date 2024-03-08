// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
)

func TestRadarAttackLayer7TimeseriesGroupGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.Get(context.TODO(), cloudflare.RadarAttackLayer7TimeseriesGroupGetParams{
		AggInterval: cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupGetParamsAggInterval1h),
		ASN:         cloudflare.F([]string{"string", "string", "string"}),
		Continent:   cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupGetParamsDateRange{cloudflare.RadarAttackLayer7TimeseriesGroupGetParamsDateRange1d, cloudflare.RadarAttackLayer7TimeseriesGroupGetParamsDateRange2d, cloudflare.RadarAttackLayer7TimeseriesGroupGetParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:      cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupGetParamsFormatJson),
		Location:    cloudflare.F([]string{"string", "string", "string"}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7TimeseriesGroupHTTPMethodWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.HTTPMethod(context.TODO(), cloudflare.RadarAttackLayer7TimeseriesGroupHTTPMethodParams{
		AggInterval:       cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval1h),
		ASN:               cloudflare.F([]string{"string", "string", "string"}),
		Continent:         cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange{cloudflare.RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange1d, cloudflare.RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange2d, cloudflare.RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange7d}),
		DateStart:         cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupHTTPMethodParamsFormatJson),
		HTTPVersion:       cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersion{cloudflare.RadarAttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersionHttPv1, cloudflare.RadarAttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersionHttPv2, cloudflare.RadarAttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersionHttPv3}),
		IPVersion:         cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupHTTPMethodParamsIPVersion{cloudflare.RadarAttackLayer7TimeseriesGroupHTTPMethodParamsIPVersionIPv4, cloudflare.RadarAttackLayer7TimeseriesGroupHTTPMethodParamsIPVersionIPv6}),
		Location:          cloudflare.F([]string{"string", "string", "string"}),
		MitigationProduct: cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct{cloudflare.RadarAttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductDDOS, cloudflare.RadarAttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductWAF, cloudflare.RadarAttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductBotManagement}),
		Name:              cloudflare.F([]string{"string", "string", "string"}),
		Normalization:     cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupHTTPMethodParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7TimeseriesGroupHTTPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.HTTPVersion(context.TODO(), cloudflare.RadarAttackLayer7TimeseriesGroupHTTPVersionParams{
		AggInterval:       cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval1h),
		ASN:               cloudflare.F([]string{"string", "string", "string"}),
		Continent:         cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange{cloudflare.RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange1d, cloudflare.RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange2d, cloudflare.RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange7d}),
		DateStart:         cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupHTTPVersionParamsFormatJson),
		HTTPMethod:        cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod{cloudflare.RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodGet, cloudflare.RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPost, cloudflare.RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodDelete}),
		IPVersion:         cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupHTTPVersionParamsIPVersion{cloudflare.RadarAttackLayer7TimeseriesGroupHTTPVersionParamsIPVersionIPv4, cloudflare.RadarAttackLayer7TimeseriesGroupHTTPVersionParamsIPVersionIPv6}),
		Location:          cloudflare.F([]string{"string", "string", "string"}),
		MitigationProduct: cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct{cloudflare.RadarAttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductDDOS, cloudflare.RadarAttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductWAF, cloudflare.RadarAttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductBotManagement}),
		Name:              cloudflare.F([]string{"string", "string", "string"}),
		Normalization:     cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupHTTPVersionParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7TimeseriesGroupIndustryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.Industry(context.TODO(), cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParams{
		AggInterval:       cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsAggInterval1h),
		ASN:               cloudflare.F([]string{"string", "string", "string"}),
		Continent:         cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange{cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange1d, cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange2d, cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange7d}),
		DateStart:         cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsFormatJson),
		HTTPMethod:        cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod{cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodGet, cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPost, cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodDelete}),
		HTTPVersion:       cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPVersion{cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPVersionHttPv1, cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPVersionHttPv2, cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPVersionHttPv3}),
		IPVersion:         cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsIPVersion{cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsIPVersionIPv4, cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsIPVersionIPv6}),
		LimitPerGroup:     cloudflare.F(int64(4)),
		Location:          cloudflare.F([]string{"string", "string", "string"}),
		MitigationProduct: cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsMitigationProduct{cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsMitigationProductDDOS, cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsMitigationProductWAF, cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsMitigationProductBotManagement}),
		Name:              cloudflare.F([]string{"string", "string", "string"}),
		Normalization:     cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7TimeseriesGroupIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.IPVersion(context.TODO(), cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParams{
		AggInterval:       cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParamsAggInterval1h),
		ASN:               cloudflare.F([]string{"string", "string", "string"}),
		Continent:         cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange{cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange1d, cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange2d, cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange7d}),
		DateStart:         cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParamsFormatJson),
		HTTPMethod:        cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod{cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodGet, cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPost, cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodDelete}),
		HTTPVersion:       cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPVersion{cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPVersionHttPv1, cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPVersionHttPv2, cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPVersionHttPv3}),
		Location:          cloudflare.F([]string{"string", "string", "string"}),
		MitigationProduct: cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct{cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParamsMitigationProductDDOS, cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParamsMitigationProductWAF, cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParamsMitigationProductBotManagement}),
		Name:              cloudflare.F([]string{"string", "string", "string"}),
		Normalization:     cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7TimeseriesGroupManagedRulesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.ManagedRules(context.TODO(), cloudflare.RadarAttackLayer7TimeseriesGroupManagedRulesParams{
		AggInterval:       cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupManagedRulesParamsAggInterval1h),
		ASN:               cloudflare.F([]string{"string", "string", "string"}),
		Continent:         cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange{cloudflare.RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange1d, cloudflare.RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange2d, cloudflare.RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange7d}),
		DateStart:         cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupManagedRulesParamsFormatJson),
		HTTPMethod:        cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod{cloudflare.RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodGet, cloudflare.RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPost, cloudflare.RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodDelete}),
		HTTPVersion:       cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersion{cloudflare.RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersionHttPv1, cloudflare.RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersionHttPv2, cloudflare.RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersionHttPv3}),
		IPVersion:         cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupManagedRulesParamsIPVersion{cloudflare.RadarAttackLayer7TimeseriesGroupManagedRulesParamsIPVersionIPv4, cloudflare.RadarAttackLayer7TimeseriesGroupManagedRulesParamsIPVersionIPv6}),
		Location:          cloudflare.F([]string{"string", "string", "string"}),
		MitigationProduct: cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct{cloudflare.RadarAttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductDDOS, cloudflare.RadarAttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductWAF, cloudflare.RadarAttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductBotManagement}),
		Name:              cloudflare.F([]string{"string", "string", "string"}),
		Normalization:     cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupManagedRulesParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7TimeseriesGroupMitigationProductWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.MitigationProduct(context.TODO(), cloudflare.RadarAttackLayer7TimeseriesGroupMitigationProductParams{
		AggInterval:   cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupMitigationProductParamsAggInterval1h),
		ASN:           cloudflare.F([]string{"string", "string", "string"}),
		Continent:     cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange{cloudflare.RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange1d, cloudflare.RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange2d, cloudflare.RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange7d}),
		DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:        cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupMitigationProductParamsFormatJson),
		HTTPMethod:    cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod{cloudflare.RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodGet, cloudflare.RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPost, cloudflare.RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodDelete}),
		HTTPVersion:   cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersion{cloudflare.RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersionHttPv1, cloudflare.RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersionHttPv2, cloudflare.RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersionHttPv3}),
		IPVersion:     cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupMitigationProductParamsIPVersion{cloudflare.RadarAttackLayer7TimeseriesGroupMitigationProductParamsIPVersionIPv4, cloudflare.RadarAttackLayer7TimeseriesGroupMitigationProductParamsIPVersionIPv6}),
		Location:      cloudflare.F([]string{"string", "string", "string"}),
		Name:          cloudflare.F([]string{"string", "string", "string"}),
		Normalization: cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupMitigationProductParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7TimeseriesGroupVerticalWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.Vertical(context.TODO(), cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParams{
		AggInterval:       cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsAggInterval1h),
		ASN:               cloudflare.F([]string{"string", "string", "string"}),
		Continent:         cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange{cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange1d, cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange2d, cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange7d}),
		DateStart:         cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsFormatJson),
		HTTPMethod:        cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod{cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodGet, cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPost, cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodDelete}),
		HTTPVersion:       cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPVersion{cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPVersionHttPv1, cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPVersionHttPv2, cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPVersionHttPv3}),
		IPVersion:         cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsIPVersion{cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsIPVersionIPv4, cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsIPVersionIPv6}),
		LimitPerGroup:     cloudflare.F(int64(4)),
		Location:          cloudflare.F([]string{"string", "string", "string"}),
		MitigationProduct: cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsMitigationProduct{cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsMitigationProductDDOS, cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsMitigationProductWAF, cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsMitigationProductBotManagement}),
		Name:              cloudflare.F([]string{"string", "string", "string"}),
		Normalization:     cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
