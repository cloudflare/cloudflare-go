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
		ASN:               cloudflare.F([]string{"string", "string", "string"}),
		Continent:         cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:         cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            cloudflare.F(radar.AttackLayer7TopAttacksParamsFormatJson),
		HTTPMethod:        cloudflare.F([]radar.AttackLayer7TopAttacksParamsHTTPMethod{radar.AttackLayer7TopAttacksParamsHTTPMethodGet, radar.AttackLayer7TopAttacksParamsHTTPMethodPost, radar.AttackLayer7TopAttacksParamsHTTPMethodDelete}),
		HTTPVersion:       cloudflare.F([]radar.AttackLayer7TopAttacksParamsHTTPVersion{radar.AttackLayer7TopAttacksParamsHTTPVersionHttPv1, radar.AttackLayer7TopAttacksParamsHTTPVersionHttPv2, radar.AttackLayer7TopAttacksParamsHTTPVersionHttPv3}),
		IPVersion:         cloudflare.F([]radar.AttackLayer7TopAttacksParamsIPVersion{radar.AttackLayer7TopAttacksParamsIPVersionIPv4, radar.AttackLayer7TopAttacksParamsIPVersionIPv6}),
		Limit:             cloudflare.F(int64(5)),
		LimitDirection:    cloudflare.F(radar.AttackLayer7TopAttacksParamsLimitDirectionOrigin),
		LimitPerLocation:  cloudflare.F(int64(10)),
		Location:          cloudflare.F([]string{"string", "string", "string"}),
		Magnitude:         cloudflare.F(radar.AttackLayer7TopAttacksParamsMagnitudeMitigatedRequests),
		MitigationProduct: cloudflare.F([]radar.AttackLayer7TopAttacksParamsMitigationProduct{radar.AttackLayer7TopAttacksParamsMitigationProductDDoS, radar.AttackLayer7TopAttacksParamsMitigationProductWAF, radar.AttackLayer7TopAttacksParamsMitigationProductBotManagement}),
		Name:              cloudflare.F([]string{"string", "string", "string"}),
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
		ASN:               cloudflare.F([]string{"string", "string", "string"}),
		Continent:         cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:         cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            cloudflare.F(radar.AttackLayer7TopIndustryParamsFormatJson),
		HTTPMethod:        cloudflare.F([]radar.AttackLayer7TopIndustryParamsHTTPMethod{radar.AttackLayer7TopIndustryParamsHTTPMethodGet, radar.AttackLayer7TopIndustryParamsHTTPMethodPost, radar.AttackLayer7TopIndustryParamsHTTPMethodDelete}),
		HTTPVersion:       cloudflare.F([]radar.AttackLayer7TopIndustryParamsHTTPVersion{radar.AttackLayer7TopIndustryParamsHTTPVersionHttPv1, radar.AttackLayer7TopIndustryParamsHTTPVersionHttPv2, radar.AttackLayer7TopIndustryParamsHTTPVersionHttPv3}),
		IPVersion:         cloudflare.F([]radar.AttackLayer7TopIndustryParamsIPVersion{radar.AttackLayer7TopIndustryParamsIPVersionIPv4, radar.AttackLayer7TopIndustryParamsIPVersionIPv6}),
		Limit:             cloudflare.F(int64(5)),
		Location:          cloudflare.F([]string{"string", "string", "string"}),
		MitigationProduct: cloudflare.F([]radar.AttackLayer7TopIndustryParamsMitigationProduct{radar.AttackLayer7TopIndustryParamsMitigationProductDDoS, radar.AttackLayer7TopIndustryParamsMitigationProductWAF, radar.AttackLayer7TopIndustryParamsMitigationProductBotManagement}),
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
		ASN:               cloudflare.F([]string{"string", "string", "string"}),
		Continent:         cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:         cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            cloudflare.F(radar.AttackLayer7TopVerticalParamsFormatJson),
		HTTPMethod:        cloudflare.F([]radar.AttackLayer7TopVerticalParamsHTTPMethod{radar.AttackLayer7TopVerticalParamsHTTPMethodGet, radar.AttackLayer7TopVerticalParamsHTTPMethodPost, radar.AttackLayer7TopVerticalParamsHTTPMethodDelete}),
		HTTPVersion:       cloudflare.F([]radar.AttackLayer7TopVerticalParamsHTTPVersion{radar.AttackLayer7TopVerticalParamsHTTPVersionHttPv1, radar.AttackLayer7TopVerticalParamsHTTPVersionHttPv2, radar.AttackLayer7TopVerticalParamsHTTPVersionHttPv3}),
		IPVersion:         cloudflare.F([]radar.AttackLayer7TopVerticalParamsIPVersion{radar.AttackLayer7TopVerticalParamsIPVersionIPv4, radar.AttackLayer7TopVerticalParamsIPVersionIPv6}),
		Limit:             cloudflare.F(int64(5)),
		Location:          cloudflare.F([]string{"string", "string", "string"}),
		MitigationProduct: cloudflare.F([]radar.AttackLayer7TopVerticalParamsMitigationProduct{radar.AttackLayer7TopVerticalParamsMitigationProductDDoS, radar.AttackLayer7TopVerticalParamsMitigationProductWAF, radar.AttackLayer7TopVerticalParamsMitigationProductBotManagement}),
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
