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

func TestRadarAttackLayer3TimeseriesGroupVectorListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.TimeseriesGroups.Vectors.List(context.TODO(), cloudflare.RadarAttackLayer3TimeseriesGroupVectorListParams{
		AggInterval:   cloudflare.F(cloudflare.RadarAttackLayer3TimeseriesGroupVectorListParamsAggInterval1h),
		DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     cloudflare.F([]cloudflare.RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange{cloudflare.RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange1d, cloudflare.RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange2d, cloudflare.RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange7d}),
		DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction:     cloudflare.F(cloudflare.RadarAttackLayer3TimeseriesGroupVectorListParamsDirectionOrigin),
		Format:        cloudflare.F(cloudflare.RadarAttackLayer3TimeseriesGroupVectorListParamsFormatJson),
		IPVersion:     cloudflare.F([]cloudflare.RadarAttackLayer3TimeseriesGroupVectorListParamsIPVersion{cloudflare.RadarAttackLayer3TimeseriesGroupVectorListParamsIPVersionIPv4, cloudflare.RadarAttackLayer3TimeseriesGroupVectorListParamsIPVersionIPv6}),
		LimitPerGroup: cloudflare.F(int64(4)),
		Location:      cloudflare.F([]string{"string", "string", "string"}),
		Name:          cloudflare.F([]string{"string", "string", "string"}),
		Normalization: cloudflare.F(cloudflare.RadarAttackLayer3TimeseriesGroupVectorListParamsNormalizationPercentage),
		Protocol:      cloudflare.F([]cloudflare.RadarAttackLayer3TimeseriesGroupVectorListParamsProtocol{cloudflare.RadarAttackLayer3TimeseriesGroupVectorListParamsProtocolUdp, cloudflare.RadarAttackLayer3TimeseriesGroupVectorListParamsProtocolTcp, cloudflare.RadarAttackLayer3TimeseriesGroupVectorListParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
