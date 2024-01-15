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

func TestRadarAttackLayer7VerticalListTimeseriesGroupsWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Radar.Attacks.Layer7.Verticals.ListTimeseriesGroups(context.TODO(), cloudflare.RadarAttackLayer7VerticalListTimeseriesGroupsParams{
		AggInterval:       cloudflare.F(cloudflare.RadarAttackLayer7VerticalListTimeseriesGroupsParamsAggInterval1h),
		ASN:               cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         cloudflare.F([]cloudflare.RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange{cloudflare.RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange1d, cloudflare.RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange2d, cloudflare.RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange7d}),
		DateStart:         cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            cloudflare.F(cloudflare.RadarAttackLayer7VerticalListTimeseriesGroupsParamsFormatJson),
		HTTPMethod:        cloudflare.F([]cloudflare.RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod{cloudflare.RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodGet, cloudflare.RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodPost, cloudflare.RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodDelete}),
		HTTPVersion:       cloudflare.F([]cloudflare.RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPVersion{cloudflare.RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPVersionHttPv1, cloudflare.RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPVersionHttPv2, cloudflare.RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPVersionHttPv3}),
		IPVersion:         cloudflare.F([]cloudflare.RadarAttackLayer7VerticalListTimeseriesGroupsParamsIPVersion{cloudflare.RadarAttackLayer7VerticalListTimeseriesGroupsParamsIPVersionIPv4, cloudflare.RadarAttackLayer7VerticalListTimeseriesGroupsParamsIPVersionIPv6}),
		LimitPerGroup:     cloudflare.F(int64(4)),
		Location:          cloudflare.F([]string{"string", "string", "string"}),
		MitigationProduct: cloudflare.F([]cloudflare.RadarAttackLayer7VerticalListTimeseriesGroupsParamsMitigationProduct{cloudflare.RadarAttackLayer7VerticalListTimeseriesGroupsParamsMitigationProductDdos, cloudflare.RadarAttackLayer7VerticalListTimeseriesGroupsParamsMitigationProductWaf, cloudflare.RadarAttackLayer7VerticalListTimeseriesGroupsParamsMitigationProductBotManagement}),
		Name:              cloudflare.F([]string{"string", "string", "string"}),
		Normalization:     cloudflare.F(cloudflare.RadarAttackLayer7VerticalListTimeseriesGroupsParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7VerticalListTopsWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Radar.Attacks.Layer7.Verticals.ListTops(context.TODO(), cloudflare.RadarAttackLayer7VerticalListTopsParams{
		ASN:       cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarAttackLayer7VerticalListTopsParamsDateRange{cloudflare.RadarAttackLayer7VerticalListTopsParamsDateRange1d, cloudflare.RadarAttackLayer7VerticalListTopsParamsDateRange2d, cloudflare.RadarAttackLayer7VerticalListTopsParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:    cloudflare.F(cloudflare.RadarAttackLayer7VerticalListTopsParamsFormatJson),
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
