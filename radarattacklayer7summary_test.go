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

func TestRadarAttackLayer7SummaryGetWithOptionalParams(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Radar.Attacks.Layer7.Summary.Get(context.TODO(), cloudflare.RadarAttackLayer7SummaryGetParams{
		Asn:       cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarAttackLayer7SummaryGetParamsDateRange{cloudflare.RadarAttackLayer7SummaryGetParamsDateRange1d, cloudflare.RadarAttackLayer7SummaryGetParamsDateRange2d, cloudflare.RadarAttackLayer7SummaryGetParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:    cloudflare.F(cloudflare.RadarAttackLayer7SummaryGetParamsFormatJson),
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

func TestRadarAttackLayer7SummaryHTTPMethodWithOptionalParams(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Radar.Attacks.Layer7.Summary.HTTPMethod(context.TODO(), cloudflare.RadarAttackLayer7SummaryHTTPMethodParams{
		Asn:               cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         cloudflare.F([]cloudflare.RadarAttackLayer7SummaryHTTPMethodParamsDateRange{cloudflare.RadarAttackLayer7SummaryHTTPMethodParamsDateRange1d, cloudflare.RadarAttackLayer7SummaryHTTPMethodParamsDateRange2d, cloudflare.RadarAttackLayer7SummaryHTTPMethodParamsDateRange7d}),
		DateStart:         cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            cloudflare.F(cloudflare.RadarAttackLayer7SummaryHTTPMethodParamsFormatJson),
		HTTPVersion:       cloudflare.F([]cloudflare.RadarAttackLayer7SummaryHTTPMethodParamsHTTPVersion{cloudflare.RadarAttackLayer7SummaryHTTPMethodParamsHTTPVersionHttPv1, cloudflare.RadarAttackLayer7SummaryHTTPMethodParamsHTTPVersionHttPv2, cloudflare.RadarAttackLayer7SummaryHTTPMethodParamsHTTPVersionHttPv3}),
		IPVersion:         cloudflare.F([]cloudflare.RadarAttackLayer7SummaryHTTPMethodParamsIPVersion{cloudflare.RadarAttackLayer7SummaryHTTPMethodParamsIPVersionIPv4, cloudflare.RadarAttackLayer7SummaryHTTPMethodParamsIPVersionIPv6}),
		Location:          cloudflare.F([]string{"string", "string", "string"}),
		MitigationProduct: cloudflare.F([]cloudflare.RadarAttackLayer7SummaryHTTPMethodParamsMitigationProduct{cloudflare.RadarAttackLayer7SummaryHTTPMethodParamsMitigationProductDDOS, cloudflare.RadarAttackLayer7SummaryHTTPMethodParamsMitigationProductWAF, cloudflare.RadarAttackLayer7SummaryHTTPMethodParamsMitigationProductBotManagement}),
		Name:              cloudflare.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7SummaryHTTPVersionWithOptionalParams(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Radar.Attacks.Layer7.Summary.HTTPVersion(context.TODO(), cloudflare.RadarAttackLayer7SummaryHTTPVersionParams{
		Asn:               cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         cloudflare.F([]cloudflare.RadarAttackLayer7SummaryHTTPVersionParamsDateRange{cloudflare.RadarAttackLayer7SummaryHTTPVersionParamsDateRange1d, cloudflare.RadarAttackLayer7SummaryHTTPVersionParamsDateRange2d, cloudflare.RadarAttackLayer7SummaryHTTPVersionParamsDateRange7d}),
		DateStart:         cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            cloudflare.F(cloudflare.RadarAttackLayer7SummaryHTTPVersionParamsFormatJson),
		HTTPMethod:        cloudflare.F([]cloudflare.RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod{cloudflare.RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodGet, cloudflare.RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodPost, cloudflare.RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodDelete}),
		IPVersion:         cloudflare.F([]cloudflare.RadarAttackLayer7SummaryHTTPVersionParamsIPVersion{cloudflare.RadarAttackLayer7SummaryHTTPVersionParamsIPVersionIPv4, cloudflare.RadarAttackLayer7SummaryHTTPVersionParamsIPVersionIPv6}),
		Location:          cloudflare.F([]string{"string", "string", "string"}),
		MitigationProduct: cloudflare.F([]cloudflare.RadarAttackLayer7SummaryHTTPVersionParamsMitigationProduct{cloudflare.RadarAttackLayer7SummaryHTTPVersionParamsMitigationProductDDOS, cloudflare.RadarAttackLayer7SummaryHTTPVersionParamsMitigationProductWAF, cloudflare.RadarAttackLayer7SummaryHTTPVersionParamsMitigationProductBotManagement}),
		Name:              cloudflare.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7SummaryIPVersionWithOptionalParams(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Radar.Attacks.Layer7.Summary.IPVersion(context.TODO(), cloudflare.RadarAttackLayer7SummaryIPVersionParams{
		Asn:               cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         cloudflare.F([]cloudflare.RadarAttackLayer7SummaryIPVersionParamsDateRange{cloudflare.RadarAttackLayer7SummaryIPVersionParamsDateRange1d, cloudflare.RadarAttackLayer7SummaryIPVersionParamsDateRange2d, cloudflare.RadarAttackLayer7SummaryIPVersionParamsDateRange7d}),
		DateStart:         cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            cloudflare.F(cloudflare.RadarAttackLayer7SummaryIPVersionParamsFormatJson),
		HTTPMethod:        cloudflare.F([]cloudflare.RadarAttackLayer7SummaryIPVersionParamsHTTPMethod{cloudflare.RadarAttackLayer7SummaryIPVersionParamsHTTPMethodGet, cloudflare.RadarAttackLayer7SummaryIPVersionParamsHTTPMethodPost, cloudflare.RadarAttackLayer7SummaryIPVersionParamsHTTPMethodDelete}),
		HTTPVersion:       cloudflare.F([]cloudflare.RadarAttackLayer7SummaryIPVersionParamsHTTPVersion{cloudflare.RadarAttackLayer7SummaryIPVersionParamsHTTPVersionHttPv1, cloudflare.RadarAttackLayer7SummaryIPVersionParamsHTTPVersionHttPv2, cloudflare.RadarAttackLayer7SummaryIPVersionParamsHTTPVersionHttPv3}),
		Location:          cloudflare.F([]string{"string", "string", "string"}),
		MitigationProduct: cloudflare.F([]cloudflare.RadarAttackLayer7SummaryIPVersionParamsMitigationProduct{cloudflare.RadarAttackLayer7SummaryIPVersionParamsMitigationProductDDOS, cloudflare.RadarAttackLayer7SummaryIPVersionParamsMitigationProductWAF, cloudflare.RadarAttackLayer7SummaryIPVersionParamsMitigationProductBotManagement}),
		Name:              cloudflare.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7SummaryManagedRulesWithOptionalParams(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Radar.Attacks.Layer7.Summary.ManagedRules(context.TODO(), cloudflare.RadarAttackLayer7SummaryManagedRulesParams{
		Asn:               cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         cloudflare.F([]cloudflare.RadarAttackLayer7SummaryManagedRulesParamsDateRange{cloudflare.RadarAttackLayer7SummaryManagedRulesParamsDateRange1d, cloudflare.RadarAttackLayer7SummaryManagedRulesParamsDateRange2d, cloudflare.RadarAttackLayer7SummaryManagedRulesParamsDateRange7d}),
		DateStart:         cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            cloudflare.F(cloudflare.RadarAttackLayer7SummaryManagedRulesParamsFormatJson),
		HTTPMethod:        cloudflare.F([]cloudflare.RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod{cloudflare.RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodGet, cloudflare.RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodPost, cloudflare.RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodDelete}),
		HTTPVersion:       cloudflare.F([]cloudflare.RadarAttackLayer7SummaryManagedRulesParamsHTTPVersion{cloudflare.RadarAttackLayer7SummaryManagedRulesParamsHTTPVersionHttPv1, cloudflare.RadarAttackLayer7SummaryManagedRulesParamsHTTPVersionHttPv2, cloudflare.RadarAttackLayer7SummaryManagedRulesParamsHTTPVersionHttPv3}),
		IPVersion:         cloudflare.F([]cloudflare.RadarAttackLayer7SummaryManagedRulesParamsIPVersion{cloudflare.RadarAttackLayer7SummaryManagedRulesParamsIPVersionIPv4, cloudflare.RadarAttackLayer7SummaryManagedRulesParamsIPVersionIPv6}),
		Location:          cloudflare.F([]string{"string", "string", "string"}),
		MitigationProduct: cloudflare.F([]cloudflare.RadarAttackLayer7SummaryManagedRulesParamsMitigationProduct{cloudflare.RadarAttackLayer7SummaryManagedRulesParamsMitigationProductDDOS, cloudflare.RadarAttackLayer7SummaryManagedRulesParamsMitigationProductWAF, cloudflare.RadarAttackLayer7SummaryManagedRulesParamsMitigationProductBotManagement}),
		Name:              cloudflare.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7SummaryMitigationProductWithOptionalParams(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Radar.Attacks.Layer7.Summary.MitigationProduct(context.TODO(), cloudflare.RadarAttackLayer7SummaryMitigationProductParams{
		Asn:         cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarAttackLayer7SummaryMitigationProductParamsDateRange{cloudflare.RadarAttackLayer7SummaryMitigationProductParamsDateRange1d, cloudflare.RadarAttackLayer7SummaryMitigationProductParamsDateRange2d, cloudflare.RadarAttackLayer7SummaryMitigationProductParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:      cloudflare.F(cloudflare.RadarAttackLayer7SummaryMitigationProductParamsFormatJson),
		HTTPMethod:  cloudflare.F([]cloudflare.RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod{cloudflare.RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodGet, cloudflare.RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodPost, cloudflare.RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodDelete}),
		HTTPVersion: cloudflare.F([]cloudflare.RadarAttackLayer7SummaryMitigationProductParamsHTTPVersion{cloudflare.RadarAttackLayer7SummaryMitigationProductParamsHTTPVersionHttPv1, cloudflare.RadarAttackLayer7SummaryMitigationProductParamsHTTPVersionHttPv2, cloudflare.RadarAttackLayer7SummaryMitigationProductParamsHTTPVersionHttPv3}),
		IPVersion:   cloudflare.F([]cloudflare.RadarAttackLayer7SummaryMitigationProductParamsIPVersion{cloudflare.RadarAttackLayer7SummaryMitigationProductParamsIPVersionIPv4, cloudflare.RadarAttackLayer7SummaryMitigationProductParamsIPVersionIPv6}),
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
