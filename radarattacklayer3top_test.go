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

func TestRadarAttackLayer3TopAttacksWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Top.Attacks(context.TODO(), cloudflare.RadarAttackLayer3TopAttacksParams{
		DateEnd:          cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:        cloudflare.F([]cloudflare.RadarAttackLayer3TopAttacksParamsDateRange{cloudflare.RadarAttackLayer3TopAttacksParamsDateRange1d, cloudflare.RadarAttackLayer3TopAttacksParamsDateRange2d, cloudflare.RadarAttackLayer3TopAttacksParamsDateRange7d}),
		DateStart:        cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:           cloudflare.F(cloudflare.RadarAttackLayer3TopAttacksParamsFormatJson),
		IPVersion:        cloudflare.F([]cloudflare.RadarAttackLayer3TopAttacksParamsIPVersion{cloudflare.RadarAttackLayer3TopAttacksParamsIPVersionIPv4, cloudflare.RadarAttackLayer3TopAttacksParamsIPVersionIPv6}),
		Limit:            cloudflare.F(int64(5)),
		LimitDirection:   cloudflare.F(cloudflare.RadarAttackLayer3TopAttacksParamsLimitDirectionOrigin),
		LimitPerLocation: cloudflare.F(int64(10)),
		Location:         cloudflare.F([]string{"string", "string", "string"}),
		Name:             cloudflare.F([]string{"string", "string", "string"}),
		Protocol:         cloudflare.F([]cloudflare.RadarAttackLayer3TopAttacksParamsProtocol{cloudflare.RadarAttackLayer3TopAttacksParamsProtocolUdp, cloudflare.RadarAttackLayer3TopAttacksParamsProtocolTcp, cloudflare.RadarAttackLayer3TopAttacksParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer3TopIndustryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Top.Industry(context.TODO(), cloudflare.RadarAttackLayer3TopIndustryParams{
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarAttackLayer3TopIndustryParamsDateRange{cloudflare.RadarAttackLayer3TopIndustryParamsDateRange1d, cloudflare.RadarAttackLayer3TopIndustryParamsDateRange2d, cloudflare.RadarAttackLayer3TopIndustryParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:    cloudflare.F(cloudflare.RadarAttackLayer3TopIndustryParamsFormatJson),
		IPVersion: cloudflare.F([]cloudflare.RadarAttackLayer3TopIndustryParamsIPVersion{cloudflare.RadarAttackLayer3TopIndustryParamsIPVersionIPv4, cloudflare.RadarAttackLayer3TopIndustryParamsIPVersionIPv6}),
		Limit:     cloudflare.F(int64(5)),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		Protocol:  cloudflare.F([]cloudflare.RadarAttackLayer3TopIndustryParamsProtocol{cloudflare.RadarAttackLayer3TopIndustryParamsProtocolUdp, cloudflare.RadarAttackLayer3TopIndustryParamsProtocolTcp, cloudflare.RadarAttackLayer3TopIndustryParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer3TopVerticalWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Top.Vertical(context.TODO(), cloudflare.RadarAttackLayer3TopVerticalParams{
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarAttackLayer3TopVerticalParamsDateRange{cloudflare.RadarAttackLayer3TopVerticalParamsDateRange1d, cloudflare.RadarAttackLayer3TopVerticalParamsDateRange2d, cloudflare.RadarAttackLayer3TopVerticalParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:    cloudflare.F(cloudflare.RadarAttackLayer3TopVerticalParamsFormatJson),
		IPVersion: cloudflare.F([]cloudflare.RadarAttackLayer3TopVerticalParamsIPVersion{cloudflare.RadarAttackLayer3TopVerticalParamsIPVersionIPv4, cloudflare.RadarAttackLayer3TopVerticalParamsIPVersionIPv6}),
		Limit:     cloudflare.F(int64(5)),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		Protocol:  cloudflare.F([]cloudflare.RadarAttackLayer3TopVerticalParamsProtocol{cloudflare.RadarAttackLayer3TopVerticalParamsProtocolUdp, cloudflare.RadarAttackLayer3TopVerticalParamsProtocolTcp, cloudflare.RadarAttackLayer3TopVerticalParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
