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

func TestAttackLayer3TopAttacksWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Top.Attacks(context.TODO(), radar.AttackLayer3TopAttacksParams{
		Continent:        cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:          cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:        cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:        cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:           cloudflare.F(radar.AttackLayer3TopAttacksParamsFormatJson),
		IPVersion:        cloudflare.F([]radar.AttackLayer3TopAttacksParamsIPVersion{radar.AttackLayer3TopAttacksParamsIPVersionIPv4, radar.AttackLayer3TopAttacksParamsIPVersionIPv6}),
		Limit:            cloudflare.F(int64(5)),
		LimitDirection:   cloudflare.F(radar.AttackLayer3TopAttacksParamsLimitDirectionOrigin),
		LimitPerLocation: cloudflare.F(int64(10)),
		Location:         cloudflare.F([]string{"string", "string", "string"}),
		Name:             cloudflare.F([]string{"string", "string", "string"}),
		Normalization:    cloudflare.F(radar.AttackLayer3TopAttacksParamsNormalizationPercentage),
		Protocol:         cloudflare.F([]radar.AttackLayer3TopAttacksParamsProtocol{radar.AttackLayer3TopAttacksParamsProtocolUdp, radar.AttackLayer3TopAttacksParamsProtocolTCP, radar.AttackLayer3TopAttacksParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3TopIndustryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Top.Industry(context.TODO(), radar.AttackLayer3TopIndustryParams{
		Continent: cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:    cloudflare.F(radar.AttackLayer3TopIndustryParamsFormatJson),
		IPVersion: cloudflare.F([]radar.AttackLayer3TopIndustryParamsIPVersion{radar.AttackLayer3TopIndustryParamsIPVersionIPv4, radar.AttackLayer3TopIndustryParamsIPVersionIPv6}),
		Limit:     cloudflare.F(int64(5)),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		Protocol:  cloudflare.F([]radar.AttackLayer3TopIndustryParamsProtocol{radar.AttackLayer3TopIndustryParamsProtocolUdp, radar.AttackLayer3TopIndustryParamsProtocolTCP, radar.AttackLayer3TopIndustryParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3TopVerticalWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Top.Vertical(context.TODO(), radar.AttackLayer3TopVerticalParams{
		Continent: cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:    cloudflare.F(radar.AttackLayer3TopVerticalParamsFormatJson),
		IPVersion: cloudflare.F([]radar.AttackLayer3TopVerticalParamsIPVersion{radar.AttackLayer3TopVerticalParamsIPVersionIPv4, radar.AttackLayer3TopVerticalParamsIPVersionIPv6}),
		Limit:     cloudflare.F(int64(5)),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		Protocol:  cloudflare.F([]radar.AttackLayer3TopVerticalParamsProtocol{radar.AttackLayer3TopVerticalParamsProtocolUdp, radar.AttackLayer3TopVerticalParamsProtocolTCP, radar.AttackLayer3TopVerticalParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
