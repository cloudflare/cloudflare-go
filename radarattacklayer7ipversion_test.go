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

func TestRadarAttackLayer7IPVersionListTimeseriesGroupsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.IPVersions.ListTimeseriesGroups(context.TODO(), cloudflare.RadarAttackLayer7IPVersionListTimeseriesGroupsParams{
		AggInterval:       cloudflare.F(cloudflare.RadarAttackLayer7IPVersionListTimeseriesGroupsParamsAggInterval1h),
		ASN:               cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         cloudflare.F([]cloudflare.RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange{cloudflare.RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange1d, cloudflare.RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange2d, cloudflare.RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange7d}),
		DateStart:         cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            cloudflare.F(cloudflare.RadarAttackLayer7IPVersionListTimeseriesGroupsParamsFormatJson),
		HTTPMethod:        cloudflare.F([]cloudflare.RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod{cloudflare.RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodGet, cloudflare.RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodPost, cloudflare.RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodDelete}),
		HTTPVersion:       cloudflare.F([]cloudflare.RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPVersion{cloudflare.RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPVersionHttPv1, cloudflare.RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPVersionHttPv2, cloudflare.RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPVersionHttPv3}),
		Location:          cloudflare.F([]string{"string", "string", "string"}),
		MitigationProduct: cloudflare.F([]cloudflare.RadarAttackLayer7IPVersionListTimeseriesGroupsParamsMitigationProduct{cloudflare.RadarAttackLayer7IPVersionListTimeseriesGroupsParamsMitigationProductDdos, cloudflare.RadarAttackLayer7IPVersionListTimeseriesGroupsParamsMitigationProductWaf, cloudflare.RadarAttackLayer7IPVersionListTimeseriesGroupsParamsMitigationProductBotManagement}),
		Name:              cloudflare.F([]string{"string", "string", "string"}),
		Normalization:     cloudflare.F(cloudflare.RadarAttackLayer7IPVersionListTimeseriesGroupsParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
