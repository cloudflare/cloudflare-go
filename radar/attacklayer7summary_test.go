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
		ASN:               cloudflare.F([]string{"string"}),
		Continent:         cloudflare.F([]string{"string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now()}),
		DateRange:         cloudflare.F([]string{"7d"}),
		DateStart:         cloudflare.F([]time.Time{time.Now()}),
		Format:            cloudflare.F(radar.AttackLayer7SummaryHTTPMethodParamsFormatJson),
		HTTPVersion:       cloudflare.F([]radar.AttackLayer7SummaryHTTPMethodParamsHTTPVersion{radar.AttackLayer7SummaryHTTPMethodParamsHTTPVersionHttPv1}),
		IPVersion:         cloudflare.F([]radar.AttackLayer7SummaryHTTPMethodParamsIPVersion{radar.AttackLayer7SummaryHTTPMethodParamsIPVersionIPv4}),
		LimitPerGroup:     cloudflare.F(int64(10)),
		Location:          cloudflare.F([]string{"string"}),
		MitigationProduct: cloudflare.F([]radar.AttackLayer7SummaryHTTPMethodParamsMitigationProduct{radar.AttackLayer7SummaryHTTPMethodParamsMitigationProductDDoS}),
		Name:              cloudflare.F([]string{"main_series"}),
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
		ASN:               cloudflare.F([]string{"string"}),
		Continent:         cloudflare.F([]string{"string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now()}),
		DateRange:         cloudflare.F([]string{"7d"}),
		DateStart:         cloudflare.F([]time.Time{time.Now()}),
		Format:            cloudflare.F(radar.AttackLayer7SummaryHTTPVersionParamsFormatJson),
		HTTPMethod:        cloudflare.F([]radar.AttackLayer7SummaryHTTPVersionParamsHTTPMethod{radar.AttackLayer7SummaryHTTPVersionParamsHTTPMethodGet}),
		IPVersion:         cloudflare.F([]radar.AttackLayer7SummaryHTTPVersionParamsIPVersion{radar.AttackLayer7SummaryHTTPVersionParamsIPVersionIPv4}),
		Location:          cloudflare.F([]string{"string"}),
		MitigationProduct: cloudflare.F([]radar.AttackLayer7SummaryHTTPVersionParamsMitigationProduct{radar.AttackLayer7SummaryHTTPVersionParamsMitigationProductDDoS}),
		Name:              cloudflare.F([]string{"main_series"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer7SummaryIndustryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Summary.Industry(context.TODO(), radar.AttackLayer7SummaryIndustryParams{
		ASN:               cloudflare.F([]string{"string"}),
		Continent:         cloudflare.F([]string{"string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now()}),
		DateRange:         cloudflare.F([]string{"7d"}),
		DateStart:         cloudflare.F([]time.Time{time.Now()}),
		Format:            cloudflare.F(radar.AttackLayer7SummaryIndustryParamsFormatJson),
		HTTPMethod:        cloudflare.F([]radar.AttackLayer7SummaryIndustryParamsHTTPMethod{radar.AttackLayer7SummaryIndustryParamsHTTPMethodGet}),
		HTTPVersion:       cloudflare.F([]radar.AttackLayer7SummaryIndustryParamsHTTPVersion{radar.AttackLayer7SummaryIndustryParamsHTTPVersionHttPv1}),
		IPVersion:         cloudflare.F([]radar.AttackLayer7SummaryIndustryParamsIPVersion{radar.AttackLayer7SummaryIndustryParamsIPVersionIPv4}),
		LimitPerGroup:     cloudflare.F(int64(10)),
		Location:          cloudflare.F([]string{"string"}),
		MitigationProduct: cloudflare.F([]radar.AttackLayer7SummaryIndustryParamsMitigationProduct{radar.AttackLayer7SummaryIndustryParamsMitigationProductDDoS}),
		Name:              cloudflare.F([]string{"main_series"}),
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
		ASN:               cloudflare.F([]string{"string"}),
		Continent:         cloudflare.F([]string{"string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now()}),
		DateRange:         cloudflare.F([]string{"7d"}),
		DateStart:         cloudflare.F([]time.Time{time.Now()}),
		Format:            cloudflare.F(radar.AttackLayer7SummaryIPVersionParamsFormatJson),
		HTTPMethod:        cloudflare.F([]radar.AttackLayer7SummaryIPVersionParamsHTTPMethod{radar.AttackLayer7SummaryIPVersionParamsHTTPMethodGet}),
		HTTPVersion:       cloudflare.F([]radar.AttackLayer7SummaryIPVersionParamsHTTPVersion{radar.AttackLayer7SummaryIPVersionParamsHTTPVersionHttPv1}),
		Location:          cloudflare.F([]string{"string"}),
		MitigationProduct: cloudflare.F([]radar.AttackLayer7SummaryIPVersionParamsMitigationProduct{radar.AttackLayer7SummaryIPVersionParamsMitigationProductDDoS}),
		Name:              cloudflare.F([]string{"main_series"}),
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
		ASN:               cloudflare.F([]string{"string"}),
		Continent:         cloudflare.F([]string{"string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now()}),
		DateRange:         cloudflare.F([]string{"7d"}),
		DateStart:         cloudflare.F([]time.Time{time.Now()}),
		Format:            cloudflare.F(radar.AttackLayer7SummaryManagedRulesParamsFormatJson),
		HTTPMethod:        cloudflare.F([]radar.AttackLayer7SummaryManagedRulesParamsHTTPMethod{radar.AttackLayer7SummaryManagedRulesParamsHTTPMethodGet}),
		HTTPVersion:       cloudflare.F([]radar.AttackLayer7SummaryManagedRulesParamsHTTPVersion{radar.AttackLayer7SummaryManagedRulesParamsHTTPVersionHttPv1}),
		IPVersion:         cloudflare.F([]radar.AttackLayer7SummaryManagedRulesParamsIPVersion{radar.AttackLayer7SummaryManagedRulesParamsIPVersionIPv4}),
		LimitPerGroup:     cloudflare.F(int64(10)),
		Location:          cloudflare.F([]string{"string"}),
		MitigationProduct: cloudflare.F([]radar.AttackLayer7SummaryManagedRulesParamsMitigationProduct{radar.AttackLayer7SummaryManagedRulesParamsMitigationProductDDoS}),
		Name:              cloudflare.F([]string{"main_series"}),
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
		ASN:           cloudflare.F([]string{"string"}),
		Continent:     cloudflare.F([]string{"string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now()}),
		DateRange:     cloudflare.F([]string{"7d"}),
		DateStart:     cloudflare.F([]time.Time{time.Now()}),
		Format:        cloudflare.F(radar.AttackLayer7SummaryMitigationProductParamsFormatJson),
		HTTPMethod:    cloudflare.F([]radar.AttackLayer7SummaryMitigationProductParamsHTTPMethod{radar.AttackLayer7SummaryMitigationProductParamsHTTPMethodGet}),
		HTTPVersion:   cloudflare.F([]radar.AttackLayer7SummaryMitigationProductParamsHTTPVersion{radar.AttackLayer7SummaryMitigationProductParamsHTTPVersionHttPv1}),
		IPVersion:     cloudflare.F([]radar.AttackLayer7SummaryMitigationProductParamsIPVersion{radar.AttackLayer7SummaryMitigationProductParamsIPVersionIPv4}),
		LimitPerGroup: cloudflare.F(int64(10)),
		Location:      cloudflare.F([]string{"string"}),
		Name:          cloudflare.F([]string{"main_series"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer7SummaryVerticalWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Summary.Vertical(context.TODO(), radar.AttackLayer7SummaryVerticalParams{
		ASN:               cloudflare.F([]string{"string"}),
		Continent:         cloudflare.F([]string{"string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now()}),
		DateRange:         cloudflare.F([]string{"7d"}),
		DateStart:         cloudflare.F([]time.Time{time.Now()}),
		Format:            cloudflare.F(radar.AttackLayer7SummaryVerticalParamsFormatJson),
		HTTPMethod:        cloudflare.F([]radar.AttackLayer7SummaryVerticalParamsHTTPMethod{radar.AttackLayer7SummaryVerticalParamsHTTPMethodGet}),
		HTTPVersion:       cloudflare.F([]radar.AttackLayer7SummaryVerticalParamsHTTPVersion{radar.AttackLayer7SummaryVerticalParamsHTTPVersionHttPv1}),
		IPVersion:         cloudflare.F([]radar.AttackLayer7SummaryVerticalParamsIPVersion{radar.AttackLayer7SummaryVerticalParamsIPVersionIPv4}),
		LimitPerGroup:     cloudflare.F(int64(10)),
		Location:          cloudflare.F([]string{"string"}),
		MitigationProduct: cloudflare.F([]radar.AttackLayer7SummaryVerticalParamsMitigationProduct{radar.AttackLayer7SummaryVerticalParamsMitigationProductDDoS}),
		Name:              cloudflare.F([]string{"main_series"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
