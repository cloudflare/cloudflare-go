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

func TestRadarAttackLayer3TimeseriesGroupBitrateListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.TimeseriesGroups.Bitrates.List(context.TODO(), cloudflare.RadarAttackLayer3TimeseriesGroupBitrateListParams{
		AggInterval:   cloudflare.F(cloudflare.RadarAttackLayer3TimeseriesGroupBitrateListParamsAggInterval1h),
		DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     cloudflare.F([]cloudflare.RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange{cloudflare.RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange1d, cloudflare.RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange2d, cloudflare.RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange7d}),
		DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction:     cloudflare.F(cloudflare.RadarAttackLayer3TimeseriesGroupBitrateListParamsDirectionOrigin),
		Format:        cloudflare.F(cloudflare.RadarAttackLayer3TimeseriesGroupBitrateListParamsFormatJson),
		IPVersion:     cloudflare.F([]cloudflare.RadarAttackLayer3TimeseriesGroupBitrateListParamsIPVersion{cloudflare.RadarAttackLayer3TimeseriesGroupBitrateListParamsIPVersionIPv4, cloudflare.RadarAttackLayer3TimeseriesGroupBitrateListParamsIPVersionIPv6}),
		Location:      cloudflare.F([]string{"string", "string", "string"}),
		Name:          cloudflare.F([]string{"string", "string", "string"}),
		Normalization: cloudflare.F(cloudflare.RadarAttackLayer3TimeseriesGroupBitrateListParamsNormalizationPercentage),
		Protocol:      cloudflare.F([]cloudflare.RadarAttackLayer3TimeseriesGroupBitrateListParamsProtocol{cloudflare.RadarAttackLayer3TimeseriesGroupBitrateListParamsProtocolUdp, cloudflare.RadarAttackLayer3TimeseriesGroupBitrateListParamsProtocolTcp, cloudflare.RadarAttackLayer3TimeseriesGroupBitrateListParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
