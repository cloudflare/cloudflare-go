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

func TestRadarAttackLayer3TimeseriesGroupIPVersionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.TimeseriesGroups.IPVersions.List(context.TODO(), cloudflare.RadarAttackLayer3TimeseriesGroupIPVersionListParams{
		AggInterval:   cloudflare.F(cloudflare.RadarAttackLayer3TimeseriesGroupIPVersionListParamsAggInterval1h),
		DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     cloudflare.F([]cloudflare.RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange{cloudflare.RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange1d, cloudflare.RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange2d, cloudflare.RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange7d}),
		DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction:     cloudflare.F(cloudflare.RadarAttackLayer3TimeseriesGroupIPVersionListParamsDirectionOrigin),
		Format:        cloudflare.F(cloudflare.RadarAttackLayer3TimeseriesGroupIPVersionListParamsFormatJson),
		Location:      cloudflare.F([]string{"string", "string", "string"}),
		Name:          cloudflare.F([]string{"string", "string", "string"}),
		Normalization: cloudflare.F(cloudflare.RadarAttackLayer3TimeseriesGroupIPVersionListParamsNormalizationPercentage),
		Protocol:      cloudflare.F([]cloudflare.RadarAttackLayer3TimeseriesGroupIPVersionListParamsProtocol{cloudflare.RadarAttackLayer3TimeseriesGroupIPVersionListParamsProtocolUdp, cloudflare.RadarAttackLayer3TimeseriesGroupIPVersionListParamsProtocolTcp, cloudflare.RadarAttackLayer3TimeseriesGroupIPVersionListParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
