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
		ASN:       cloudflare.F([]string{"string"}),
		Continent: cloudflare.F([]string{"string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now()}),
		DateRange: cloudflare.F([]string{"7d"}),
		DateStart: cloudflare.F([]time.Time{time.Now()}),
		Format:    cloudflare.F(radar.AttackLayer7SummaryGetParamsFormatJson),
		Location:  cloudflare.F([]string{"string"}),
		Name:      cloudflare.F([]string{"string"}),
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
		ASN:               cloudflare.F([]string{"string"}),
		Continent:         cloudflare.F([]string{"string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now()}),
		DateRange:         cloudflare.F([]string{"7d"}),
		DateStart:         cloudflare.F([]time.Time{time.Now()}),
		Format:            cloudflare.F(radar.AttackLayer7SummaryHTTPMethodParamsFormatJson),
		HTTPVersion:       cloudflare.F([]radar.AttackLayer7SummaryHTTPMethodParamsHTTPVersion{radar.AttackLayer7SummaryHTTPMethodParamsHTTPVersionHttPv1}),
		IPVersion:         cloudflare.F([]radar.AttackLayer7SummaryHTTPMethodParamsIPVersion{radar.AttackLayer7SummaryHTTPMethodParamsIPVersionIPv4}),
		Location:          cloudflare.F([]string{"string"}),
		MitigationProduct: cloudflare.F([]radar.AttackLayer7SummaryHTTPMethodParamsMitigationProduct{radar.AttackLayer7SummaryHTTPMethodParamsMitigationProductDDoS}),
		Name:              cloudflare.F([]string{"string"}),
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
		Name:              cloudflare.F([]string{"string"}),
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
		Name:              cloudflare.F([]string{"string"}),
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
		Location:          cloudflare.F([]string{"string"}),
		MitigationProduct: cloudflare.F([]radar.AttackLayer7SummaryManagedRulesParamsMitigationProduct{radar.AttackLayer7SummaryManagedRulesParamsMitigationProductDDoS}),
		Name:              cloudflare.F([]string{"string"}),
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
		ASN:         cloudflare.F([]string{"string"}),
		Continent:   cloudflare.F([]string{"string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now()}),
		DateRange:   cloudflare.F([]string{"7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now()}),
		Format:      cloudflare.F(radar.AttackLayer7SummaryMitigationProductParamsFormatJson),
		HTTPMethod:  cloudflare.F([]radar.AttackLayer7SummaryMitigationProductParamsHTTPMethod{radar.AttackLayer7SummaryMitigationProductParamsHTTPMethodGet}),
		HTTPVersion: cloudflare.F([]radar.AttackLayer7SummaryMitigationProductParamsHTTPVersion{radar.AttackLayer7SummaryMitigationProductParamsHTTPVersionHttPv1}),
		IPVersion:   cloudflare.F([]radar.AttackLayer7SummaryMitigationProductParamsIPVersion{radar.AttackLayer7SummaryMitigationProductParamsIPVersionIPv4}),
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
