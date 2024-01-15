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

func TestRadarAttackLayer7MitigationProductListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.MitigationProducts.List(context.TODO(), cloudflare.RadarAttackLayer7MitigationProductListParams{
		ASN:         cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarAttackLayer7MitigationProductListParamsDateRange{cloudflare.RadarAttackLayer7MitigationProductListParamsDateRange1d, cloudflare.RadarAttackLayer7MitigationProductListParamsDateRange2d, cloudflare.RadarAttackLayer7MitigationProductListParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:      cloudflare.F(cloudflare.RadarAttackLayer7MitigationProductListParamsFormatJson),
		HTTPMethod:  cloudflare.F([]cloudflare.RadarAttackLayer7MitigationProductListParamsHTTPMethod{cloudflare.RadarAttackLayer7MitigationProductListParamsHTTPMethodGet, cloudflare.RadarAttackLayer7MitigationProductListParamsHTTPMethodPost, cloudflare.RadarAttackLayer7MitigationProductListParamsHTTPMethodDelete}),
		HTTPVersion: cloudflare.F([]cloudflare.RadarAttackLayer7MitigationProductListParamsHTTPVersion{cloudflare.RadarAttackLayer7MitigationProductListParamsHTTPVersionHttPv1, cloudflare.RadarAttackLayer7MitigationProductListParamsHTTPVersionHttPv2, cloudflare.RadarAttackLayer7MitigationProductListParamsHTTPVersionHttPv3}),
		IPVersion:   cloudflare.F([]cloudflare.RadarAttackLayer7MitigationProductListParamsIPVersion{cloudflare.RadarAttackLayer7MitigationProductListParamsIPVersionIPv4, cloudflare.RadarAttackLayer7MitigationProductListParamsIPVersionIPv6}),
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

func TestRadarAttackLayer7MitigationProductListTimeseriesGroupsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.MitigationProducts.ListTimeseriesGroups(context.TODO(), cloudflare.RadarAttackLayer7MitigationProductListTimeseriesGroupsParams{
		AggInterval:   cloudflare.F(cloudflare.RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsAggInterval1h),
		ASN:           cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     cloudflare.F([]cloudflare.RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange{cloudflare.RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange1d, cloudflare.RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange2d, cloudflare.RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange7d}),
		DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:        cloudflare.F(cloudflare.RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsFormatJson),
		HTTPMethod:    cloudflare.F([]cloudflare.RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod{cloudflare.RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodGet, cloudflare.RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodPost, cloudflare.RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodDelete}),
		HTTPVersion:   cloudflare.F([]cloudflare.RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPVersion{cloudflare.RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPVersionHttPv1, cloudflare.RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPVersionHttPv2, cloudflare.RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPVersionHttPv3}),
		IPVersion:     cloudflare.F([]cloudflare.RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsIPVersion{cloudflare.RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsIPVersionIPv4, cloudflare.RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsIPVersionIPv6}),
		Location:      cloudflare.F([]string{"string", "string", "string"}),
		Name:          cloudflare.F([]string{"string", "string", "string"}),
		Normalization: cloudflare.F(cloudflare.RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
