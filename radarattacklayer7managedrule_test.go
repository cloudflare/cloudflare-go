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

func TestRadarAttackLayer7ManagedRuleListTimeseriesGroupsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.ManagedRules.ListTimeseriesGroups(context.TODO(), cloudflare.RadarAttackLayer7ManagedRuleListTimeseriesGroupsParams{
		AggInterval:       cloudflare.F(cloudflare.RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsAggInterval1h),
		ASN:               cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         cloudflare.F([]cloudflare.RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange{cloudflare.RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange1d, cloudflare.RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange2d, cloudflare.RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange7d}),
		DateStart:         cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            cloudflare.F(cloudflare.RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsFormatJson),
		HTTPMethod:        cloudflare.F([]cloudflare.RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod{cloudflare.RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodGet, cloudflare.RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodPost, cloudflare.RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodDelete}),
		HTTPVersion:       cloudflare.F([]cloudflare.RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPVersion{cloudflare.RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPVersionHttPv1, cloudflare.RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPVersionHttPv2, cloudflare.RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPVersionHttPv3}),
		IPVersion:         cloudflare.F([]cloudflare.RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsIPVersion{cloudflare.RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsIPVersionIPv4, cloudflare.RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsIPVersionIPv6}),
		Location:          cloudflare.F([]string{"string", "string", "string"}),
		MitigationProduct: cloudflare.F([]cloudflare.RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsMitigationProduct{cloudflare.RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsMitigationProductDdos, cloudflare.RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsMitigationProductWaf, cloudflare.RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsMitigationProductBotManagement}),
		Name:              cloudflare.F([]string{"string", "string", "string"}),
		Normalization:     cloudflare.F(cloudflare.RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
