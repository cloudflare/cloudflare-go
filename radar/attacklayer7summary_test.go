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

func TestAttackLayer7SummaryGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Summary.Get(context.TODO(), radar.AttackLayer7SummaryGetParams{
		ASN:       cloudflare.F([]string{"string", "string", "string"}),
		Continent: cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:    cloudflare.F(radar.AttackLayer7SummaryGetParamsFormatJson),
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

func TestAttackLayer7SummaryHTTPMethodWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Summary.HTTPMethod(context.TODO(), radar.AttackLayer7SummaryHTTPMethodParams{
		ASN:               cloudflare.F([]string{"string", "string", "string"}),
		Continent:         cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:         cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            cloudflare.F(radar.AttackLayer7SummaryHTTPMethodParamsFormatJson),
		HTTPVersion:       cloudflare.F([]radar.AttackLayer7SummaryHTTPMethodParamsHTTPVersion{radar.AttackLayer7SummaryHTTPMethodParamsHTTPVersionHttPv1, radar.AttackLayer7SummaryHTTPMethodParamsHTTPVersionHttPv2, radar.AttackLayer7SummaryHTTPMethodParamsHTTPVersionHttPv3}),
		IPVersion:         cloudflare.F([]radar.AttackLayer7SummaryHTTPMethodParamsIPVersion{radar.AttackLayer7SummaryHTTPMethodParamsIPVersionIPv4, radar.AttackLayer7SummaryHTTPMethodParamsIPVersionIPv6}),
		Location:          cloudflare.F([]string{"string", "string", "string"}),
		MitigationProduct: cloudflare.F([]radar.AttackLayer7SummaryHTTPMethodParamsMitigationProduct{radar.AttackLayer7SummaryHTTPMethodParamsMitigationProductDDoS, radar.AttackLayer7SummaryHTTPMethodParamsMitigationProductWAF, radar.AttackLayer7SummaryHTTPMethodParamsMitigationProductBotManagement}),
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

func TestAttackLayer7SummaryHTTPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Summary.HTTPVersion(context.TODO(), radar.AttackLayer7SummaryHTTPVersionParams{
		ASN:               cloudflare.F([]string{"string", "string", "string"}),
		Continent:         cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:         cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            cloudflare.F(radar.AttackLayer7SummaryHTTPVersionParamsFormatJson),
		HTTPMethod:        cloudflare.F([]radar.AttackLayer7SummaryHTTPVersionParamsHTTPMethod{radar.AttackLayer7SummaryHTTPVersionParamsHTTPMethodGet, radar.AttackLayer7SummaryHTTPVersionParamsHTTPMethodPost, radar.AttackLayer7SummaryHTTPVersionParamsHTTPMethodDelete}),
		IPVersion:         cloudflare.F([]radar.AttackLayer7SummaryHTTPVersionParamsIPVersion{radar.AttackLayer7SummaryHTTPVersionParamsIPVersionIPv4, radar.AttackLayer7SummaryHTTPVersionParamsIPVersionIPv6}),
		Location:          cloudflare.F([]string{"string", "string", "string"}),
		MitigationProduct: cloudflare.F([]radar.AttackLayer7SummaryHTTPVersionParamsMitigationProduct{radar.AttackLayer7SummaryHTTPVersionParamsMitigationProductDDoS, radar.AttackLayer7SummaryHTTPVersionParamsMitigationProductWAF, radar.AttackLayer7SummaryHTTPVersionParamsMitigationProductBotManagement}),
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

func TestAttackLayer7SummaryIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Summary.IPVersion(context.TODO(), radar.AttackLayer7SummaryIPVersionParams{
		ASN:               cloudflare.F([]string{"string", "string", "string"}),
		Continent:         cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:         cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            cloudflare.F(radar.AttackLayer7SummaryIPVersionParamsFormatJson),
		HTTPMethod:        cloudflare.F([]radar.AttackLayer7SummaryIPVersionParamsHTTPMethod{radar.AttackLayer7SummaryIPVersionParamsHTTPMethodGet, radar.AttackLayer7SummaryIPVersionParamsHTTPMethodPost, radar.AttackLayer7SummaryIPVersionParamsHTTPMethodDelete}),
		HTTPVersion:       cloudflare.F([]radar.AttackLayer7SummaryIPVersionParamsHTTPVersion{radar.AttackLayer7SummaryIPVersionParamsHTTPVersionHttPv1, radar.AttackLayer7SummaryIPVersionParamsHTTPVersionHttPv2, radar.AttackLayer7SummaryIPVersionParamsHTTPVersionHttPv3}),
		Location:          cloudflare.F([]string{"string", "string", "string"}),
		MitigationProduct: cloudflare.F([]radar.AttackLayer7SummaryIPVersionParamsMitigationProduct{radar.AttackLayer7SummaryIPVersionParamsMitigationProductDDoS, radar.AttackLayer7SummaryIPVersionParamsMitigationProductWAF, radar.AttackLayer7SummaryIPVersionParamsMitigationProductBotManagement}),
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

func TestAttackLayer7SummaryManagedRulesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Summary.ManagedRules(context.TODO(), radar.AttackLayer7SummaryManagedRulesParams{
		ASN:               cloudflare.F([]string{"string", "string", "string"}),
		Continent:         cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:         cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            cloudflare.F(radar.AttackLayer7SummaryManagedRulesParamsFormatJson),
		HTTPMethod:        cloudflare.F([]radar.AttackLayer7SummaryManagedRulesParamsHTTPMethod{radar.AttackLayer7SummaryManagedRulesParamsHTTPMethodGet, radar.AttackLayer7SummaryManagedRulesParamsHTTPMethodPost, radar.AttackLayer7SummaryManagedRulesParamsHTTPMethodDelete}),
		HTTPVersion:       cloudflare.F([]radar.AttackLayer7SummaryManagedRulesParamsHTTPVersion{radar.AttackLayer7SummaryManagedRulesParamsHTTPVersionHttPv1, radar.AttackLayer7SummaryManagedRulesParamsHTTPVersionHttPv2, radar.AttackLayer7SummaryManagedRulesParamsHTTPVersionHttPv3}),
		IPVersion:         cloudflare.F([]radar.AttackLayer7SummaryManagedRulesParamsIPVersion{radar.AttackLayer7SummaryManagedRulesParamsIPVersionIPv4, radar.AttackLayer7SummaryManagedRulesParamsIPVersionIPv6}),
		Location:          cloudflare.F([]string{"string", "string", "string"}),
		MitigationProduct: cloudflare.F([]radar.AttackLayer7SummaryManagedRulesParamsMitigationProduct{radar.AttackLayer7SummaryManagedRulesParamsMitigationProductDDoS, radar.AttackLayer7SummaryManagedRulesParamsMitigationProductWAF, radar.AttackLayer7SummaryManagedRulesParamsMitigationProductBotManagement}),
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

func TestAttackLayer7SummaryMitigationProductWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Summary.MitigationProduct(context.TODO(), radar.AttackLayer7SummaryMitigationProductParams{
		ASN:         cloudflare.F([]string{"string", "string", "string"}),
		Continent:   cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:      cloudflare.F(radar.AttackLayer7SummaryMitigationProductParamsFormatJson),
		HTTPMethod:  cloudflare.F([]radar.AttackLayer7SummaryMitigationProductParamsHTTPMethod{radar.AttackLayer7SummaryMitigationProductParamsHTTPMethodGet, radar.AttackLayer7SummaryMitigationProductParamsHTTPMethodPost, radar.AttackLayer7SummaryMitigationProductParamsHTTPMethodDelete}),
		HTTPVersion: cloudflare.F([]radar.AttackLayer7SummaryMitigationProductParamsHTTPVersion{radar.AttackLayer7SummaryMitigationProductParamsHTTPVersionHttPv1, radar.AttackLayer7SummaryMitigationProductParamsHTTPVersionHttPv2, radar.AttackLayer7SummaryMitigationProductParamsHTTPVersionHttPv3}),
		IPVersion:   cloudflare.F([]radar.AttackLayer7SummaryMitigationProductParamsIPVersion{radar.AttackLayer7SummaryMitigationProductParamsIPVersionIPv4, radar.AttackLayer7SummaryMitigationProductParamsIPVersionIPv6}),
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
