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

func TestRadarAttackLayer7HTTPVersionListTimeseriesGroupsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.HTTPVersions.ListTimeseriesGroups(context.TODO(), cloudflare.RadarAttackLayer7HTTPVersionListTimeseriesGroupsParams{
		AggInterval:       cloudflare.F(cloudflare.RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsAggInterval1h),
		ASN:               cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         cloudflare.F([]cloudflare.RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange{cloudflare.RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange1d, cloudflare.RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange2d, cloudflare.RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange7d}),
		DateStart:         cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            cloudflare.F(cloudflare.RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsFormatJson),
		HTTPMethod:        cloudflare.F([]cloudflare.RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod{cloudflare.RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodGet, cloudflare.RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodPost, cloudflare.RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodDelete}),
		IPVersion:         cloudflare.F([]cloudflare.RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsIPVersion{cloudflare.RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsIPVersionIPv4, cloudflare.RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsIPVersionIPv6}),
		Location:          cloudflare.F([]string{"string", "string", "string"}),
		MitigationProduct: cloudflare.F([]cloudflare.RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsMitigationProduct{cloudflare.RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsMitigationProductDdos, cloudflare.RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsMitigationProductWaf, cloudflare.RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsMitigationProductBotManagement}),
		Name:              cloudflare.F([]string{"string", "string", "string"}),
		Normalization:     cloudflare.F(cloudflare.RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
