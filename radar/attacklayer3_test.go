// File generated from our OpenAPI spec by Stainless.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/cloudflare/cloudflare-go/radar"
)

func TestAttackLayer3TimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Timeseries(context.TODO(), radar.AttackLayer3TimeseriesParams{
		AggInterval:   cloudflare.F(radar.AttackLayer3TimeseriesParamsAggInterval1h),
		ASN:           cloudflare.F([]string{"string", "string", "string"}),
		Continent:     cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     cloudflare.F([]radar.AttackLayer3TimeseriesParamsDateRange{radar.AttackLayer3TimeseriesParamsDateRange1d, radar.AttackLayer3TimeseriesParamsDateRange2d, radar.AttackLayer3TimeseriesParamsDateRange7d}),
		DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction:     cloudflare.F(radar.AttackLayer3TimeseriesParamsDirectionOrigin),
		Format:        cloudflare.F(radar.AttackLayer3TimeseriesParamsFormatJson),
		IPVersion:     cloudflare.F([]radar.AttackLayer3TimeseriesParamsIPVersion{radar.AttackLayer3TimeseriesParamsIPVersionIPv4, radar.AttackLayer3TimeseriesParamsIPVersionIPv6}),
		Location:      cloudflare.F([]string{"string", "string", "string"}),
		Metric:        cloudflare.F(radar.AttackLayer3TimeseriesParamsMetricBytes),
		Name:          cloudflare.F([]string{"string", "string", "string"}),
		Normalization: cloudflare.F(radar.AttackLayer3TimeseriesParamsNormalizationMin0Max),
		Protocol:      cloudflare.F([]radar.AttackLayer3TimeseriesParamsProtocol{radar.AttackLayer3TimeseriesParamsProtocolUdp, radar.AttackLayer3TimeseriesParamsProtocolTcp, radar.AttackLayer3TimeseriesParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
