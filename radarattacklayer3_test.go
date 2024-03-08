// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
)

func TestRadarAttackLayer3TimeseriesWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Radar.Attacks.Layer3.Timeseries(context.TODO(), cloudflare.RadarAttackLayer3TimeseriesParams{
		AggInterval:   cloudflare.F(cloudflare.RadarAttackLayer3TimeseriesParamsAggInterval1h),
		ASN:           cloudflare.F([]string{"string", "string", "string"}),
		Continent:     cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     cloudflare.F([]cloudflare.RadarAttackLayer3TimeseriesParamsDateRange{cloudflare.RadarAttackLayer3TimeseriesParamsDateRange1d, cloudflare.RadarAttackLayer3TimeseriesParamsDateRange2d, cloudflare.RadarAttackLayer3TimeseriesParamsDateRange7d}),
		DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction:     cloudflare.F(cloudflare.RadarAttackLayer3TimeseriesParamsDirectionOrigin),
		Format:        cloudflare.F(cloudflare.RadarAttackLayer3TimeseriesParamsFormatJson),
		IPVersion:     cloudflare.F([]cloudflare.RadarAttackLayer3TimeseriesParamsIPVersion{cloudflare.RadarAttackLayer3TimeseriesParamsIPVersionIPv4, cloudflare.RadarAttackLayer3TimeseriesParamsIPVersionIPv6}),
		Location:      cloudflare.F([]string{"string", "string", "string"}),
		Metric:        cloudflare.F(cloudflare.RadarAttackLayer3TimeseriesParamsMetricBytes),
		Name:          cloudflare.F([]string{"string", "string", "string"}),
		Normalization: cloudflare.F(cloudflare.RadarAttackLayer3TimeseriesParamsNormalizationMin0Max),
		Protocol:      cloudflare.F([]cloudflare.RadarAttackLayer3TimeseriesParamsProtocol{cloudflare.RadarAttackLayer3TimeseriesParamsProtocolUdp, cloudflare.RadarAttackLayer3TimeseriesParamsProtocolTcp, cloudflare.RadarAttackLayer3TimeseriesParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
