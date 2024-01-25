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

func TestRadarAttackLayer7IndustryListTimeseriesGroupsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Industries.ListTimeseriesGroups(context.TODO(), cloudflare.RadarAttackLayer7IndustryListTimeseriesGroupsParams{
		AggInterval:       cloudflare.F(cloudflare.RadarAttackLayer7IndustryListTimeseriesGroupsParamsAggInterval1h),
		ASN:               cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         cloudflare.F([]cloudflare.RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange{cloudflare.RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange1d, cloudflare.RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange2d, cloudflare.RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange7d}),
		DateStart:         cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            cloudflare.F(cloudflare.RadarAttackLayer7IndustryListTimeseriesGroupsParamsFormatJson),
		HTTPMethod:        cloudflare.F([]cloudflare.RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod{cloudflare.RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodGet, cloudflare.RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodPost, cloudflare.RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodDelete}),
		HTTPVersion:       cloudflare.F([]cloudflare.RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPVersion{cloudflare.RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPVersionHttPv1, cloudflare.RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPVersionHttPv2, cloudflare.RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPVersionHttPv3}),
		IPVersion:         cloudflare.F([]cloudflare.RadarAttackLayer7IndustryListTimeseriesGroupsParamsIPVersion{cloudflare.RadarAttackLayer7IndustryListTimeseriesGroupsParamsIPVersionIPv4, cloudflare.RadarAttackLayer7IndustryListTimeseriesGroupsParamsIPVersionIPv6}),
		LimitPerGroup:     cloudflare.F(int64(4)),
		Location:          cloudflare.F([]string{"string", "string", "string"}),
		MitigationProduct: cloudflare.F([]cloudflare.RadarAttackLayer7IndustryListTimeseriesGroupsParamsMitigationProduct{cloudflare.RadarAttackLayer7IndustryListTimeseriesGroupsParamsMitigationProductDdos, cloudflare.RadarAttackLayer7IndustryListTimeseriesGroupsParamsMitigationProductWaf, cloudflare.RadarAttackLayer7IndustryListTimeseriesGroupsParamsMitigationProductBotManagement}),
		Name:              cloudflare.F([]string{"string", "string", "string"}),
		Normalization:     cloudflare.F(cloudflare.RadarAttackLayer7IndustryListTimeseriesGroupsParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7IndustryListTopsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Industries.ListTops(context.TODO(), cloudflare.RadarAttackLayer7IndustryListTopsParams{
		ASN:       cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarAttackLayer7IndustryListTopsParamsDateRange{cloudflare.RadarAttackLayer7IndustryListTopsParamsDateRange1d, cloudflare.RadarAttackLayer7IndustryListTopsParamsDateRange2d, cloudflare.RadarAttackLayer7IndustryListTopsParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:    cloudflare.F(cloudflare.RadarAttackLayer7IndustryListTopsParamsFormatJson),
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
