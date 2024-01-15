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

func TestRadarAttackLayer7HTTPMethodListTimeseriesGroupsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.HTTPMethods.ListTimeseriesGroups(context.TODO(), cloudflare.RadarAttackLayer7HTTPMethodListTimeseriesGroupsParams{
		AggInterval:       cloudflare.F(cloudflare.RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsAggInterval1h),
		ASN:               cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         cloudflare.F([]cloudflare.RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange{cloudflare.RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange1d, cloudflare.RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange2d, cloudflare.RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange7d}),
		DateStart:         cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            cloudflare.F(cloudflare.RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsFormatJson),
		HTTPVersion:       cloudflare.F([]cloudflare.RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsHTTPVersion{cloudflare.RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsHTTPVersionHttPv1, cloudflare.RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsHTTPVersionHttPv2, cloudflare.RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsHTTPVersionHttPv3}),
		IPVersion:         cloudflare.F([]cloudflare.RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsIPVersion{cloudflare.RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsIPVersionIPv4, cloudflare.RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsIPVersionIPv6}),
		Location:          cloudflare.F([]string{"string", "string", "string"}),
		MitigationProduct: cloudflare.F([]cloudflare.RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsMitigationProduct{cloudflare.RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsMitigationProductDdos, cloudflare.RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsMitigationProductWaf, cloudflare.RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsMitigationProductBotManagement}),
		Name:              cloudflare.F([]string{"string", "string", "string"}),
		Normalization:     cloudflare.F(cloudflare.RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
