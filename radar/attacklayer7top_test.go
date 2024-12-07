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

func TestAttackLayer7TopAttacksWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Top.Attacks(context.TODO(), radar.AttackLayer7TopAttacksParams{
		ASN:               cloudflare.F([]string{"string"}),
		Continent:         cloudflare.F([]string{"string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now()}),
		DateRange:         cloudflare.F([]string{"7d"}),
		DateStart:         cloudflare.F([]time.Time{time.Now()}),
		Format:            cloudflare.F(radar.AttackLayer7TopAttacksParamsFormatJson),
		HTTPMethod:        cloudflare.F([]radar.AttackLayer7TopAttacksParamsHTTPMethod{radar.AttackLayer7TopAttacksParamsHTTPMethodGet}),
		HTTPVersion:       cloudflare.F([]radar.AttackLayer7TopAttacksParamsHTTPVersion{radar.AttackLayer7TopAttacksParamsHTTPVersionHttPv1}),
		IPVersion:         cloudflare.F([]radar.AttackLayer7TopAttacksParamsIPVersion{radar.AttackLayer7TopAttacksParamsIPVersionIPv4}),
		Limit:             cloudflare.F(int64(5)),
		LimitDirection:    cloudflare.F(radar.AttackLayer7TopAttacksParamsLimitDirectionOrigin),
		LimitPerLocation:  cloudflare.F(int64(10)),
		Location:          cloudflare.F([]string{"string"}),
		Magnitude:         cloudflare.F(radar.AttackLayer7TopAttacksParamsMagnitudeAffectedZones),
		MitigationProduct: cloudflare.F([]radar.AttackLayer7TopAttacksParamsMitigationProduct{radar.AttackLayer7TopAttacksParamsMitigationProductDDoS}),
		Name:              cloudflare.F([]string{"string"}),
		Normalization:     cloudflare.F(radar.AttackLayer7TopAttacksParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer7TopIndustryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Top.Industry(context.TODO(), radar.AttackLayer7TopIndustryParams{
		ASN:               cloudflare.F([]string{"string"}),
		Continent:         cloudflare.F([]string{"string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now()}),
		DateRange:         cloudflare.F([]string{"7d"}),
		DateStart:         cloudflare.F([]time.Time{time.Now()}),
		Format:            cloudflare.F(radar.AttackLayer7TopIndustryParamsFormatJson),
		HTTPMethod:        cloudflare.F([]radar.AttackLayer7TopIndustryParamsHTTPMethod{radar.AttackLayer7TopIndustryParamsHTTPMethodGet}),
		HTTPVersion:       cloudflare.F([]radar.AttackLayer7TopIndustryParamsHTTPVersion{radar.AttackLayer7TopIndustryParamsHTTPVersionHttPv1}),
		IPVersion:         cloudflare.F([]radar.AttackLayer7TopIndustryParamsIPVersion{radar.AttackLayer7TopIndustryParamsIPVersionIPv4}),
		Limit:             cloudflare.F(int64(5)),
		Location:          cloudflare.F([]string{"string"}),
		MitigationProduct: cloudflare.F([]radar.AttackLayer7TopIndustryParamsMitigationProduct{radar.AttackLayer7TopIndustryParamsMitigationProductDDoS}),
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

func TestAttackLayer7TopVerticalWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Top.Vertical(context.TODO(), radar.AttackLayer7TopVerticalParams{
		ASN:               cloudflare.F([]string{"string"}),
		Continent:         cloudflare.F([]string{"string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now()}),
		DateRange:         cloudflare.F([]string{"7d"}),
		DateStart:         cloudflare.F([]time.Time{time.Now()}),
		Format:            cloudflare.F(radar.AttackLayer7TopVerticalParamsFormatJson),
		HTTPMethod:        cloudflare.F([]radar.AttackLayer7TopVerticalParamsHTTPMethod{radar.AttackLayer7TopVerticalParamsHTTPMethodGet}),
		HTTPVersion:       cloudflare.F([]radar.AttackLayer7TopVerticalParamsHTTPVersion{radar.AttackLayer7TopVerticalParamsHTTPVersionHttPv1}),
		IPVersion:         cloudflare.F([]radar.AttackLayer7TopVerticalParamsIPVersion{radar.AttackLayer7TopVerticalParamsIPVersionIPv4}),
		Limit:             cloudflare.F(int64(5)),
		Location:          cloudflare.F([]string{"string"}),
		MitigationProduct: cloudflare.F([]radar.AttackLayer7TopVerticalParamsMitigationProduct{radar.AttackLayer7TopVerticalParamsMitigationProductDDoS}),
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
