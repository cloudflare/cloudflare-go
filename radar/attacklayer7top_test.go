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
	_, err := client.Radar.Attacks.Layer7.Top.Attacks(context.TODO(), radar.AttackLayer7TopAttacksParams{
		ASN:              cloudflare.F([]string{"string", "string", "string"}),
		Continent:        cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:          cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:        cloudflare.F([]radar.AttackLayer7TopAttacksParamsDateRange{radar.AttackLayer7TopAttacksParamsDateRange1d, radar.AttackLayer7TopAttacksParamsDateRange2d, radar.AttackLayer7TopAttacksParamsDateRange7d}),
		DateStart:        cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:           cloudflare.F(radar.AttackLayer7TopAttacksParamsFormatJson),
		Limit:            cloudflare.F(int64(5)),
		LimitDirection:   cloudflare.F(radar.AttackLayer7TopAttacksParamsLimitDirectionOrigin),
		LimitPerLocation: cloudflare.F(int64(10)),
		Location:         cloudflare.F([]string{"string", "string", "string"}),
		Magnitude:        cloudflare.F(radar.AttackLayer7TopAttacksParamsMagnitudeMitigatedRequests),
		Name:             cloudflare.F([]string{"string", "string", "string"}),
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
	_, err := client.Radar.Attacks.Layer7.Top.Industry(context.TODO(), radar.AttackLayer7TopIndustryParams{
		ASN:       cloudflare.F([]string{"string", "string", "string"}),
		Continent: cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]radar.AttackLayer7TopIndustryParamsDateRange{radar.AttackLayer7TopIndustryParamsDateRange1d, radar.AttackLayer7TopIndustryParamsDateRange2d, radar.AttackLayer7TopIndustryParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:    cloudflare.F(radar.AttackLayer7TopIndustryParamsFormatJson),
		Limit:     cloudflare.F(int64(5)),
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

func TestAttackLayer7TopVerticalWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Top.Vertical(context.TODO(), radar.AttackLayer7TopVerticalParams{
		ASN:       cloudflare.F([]string{"string", "string", "string"}),
		Continent: cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]radar.AttackLayer7TopVerticalParamsDateRange{radar.AttackLayer7TopVerticalParamsDateRange1d, radar.AttackLayer7TopVerticalParamsDateRange2d, radar.AttackLayer7TopVerticalParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:    cloudflare.F(radar.AttackLayer7TopVerticalParamsFormatJson),
		Limit:     cloudflare.F(int64(5)),
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
