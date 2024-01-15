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

func TestRadarAttackLayer3TimeseryListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Timeseries.List(context.TODO(), cloudflare.RadarAttackLayer3TimeseryListParams{
		AggInterval:   cloudflare.F(cloudflare.RadarAttackLayer3TimeseryListParamsAggInterval1h),
		ASN:           cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     cloudflare.F([]cloudflare.RadarAttackLayer3TimeseryListParamsDateRange{cloudflare.RadarAttackLayer3TimeseryListParamsDateRange1d, cloudflare.RadarAttackLayer3TimeseryListParamsDateRange2d, cloudflare.RadarAttackLayer3TimeseryListParamsDateRange7d}),
		DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction:     cloudflare.F(cloudflare.RadarAttackLayer3TimeseryListParamsDirectionOrigin),
		Format:        cloudflare.F(cloudflare.RadarAttackLayer3TimeseryListParamsFormatJson),
		IPVersion:     cloudflare.F([]cloudflare.RadarAttackLayer3TimeseryListParamsIPVersion{cloudflare.RadarAttackLayer3TimeseryListParamsIPVersionIPv4, cloudflare.RadarAttackLayer3TimeseryListParamsIPVersionIPv6}),
		Location:      cloudflare.F([]string{"string", "string", "string"}),
		Metric:        cloudflare.F(cloudflare.RadarAttackLayer3TimeseryListParamsMetricBytes),
		Name:          cloudflare.F([]string{"string", "string", "string"}),
		Normalization: cloudflare.F(cloudflare.RadarAttackLayer3TimeseryListParamsNormalizationMin0Max),
		Protocol:      cloudflare.F([]cloudflare.RadarAttackLayer3TimeseryListParamsProtocol{cloudflare.RadarAttackLayer3TimeseryListParamsProtocolUdp, cloudflare.RadarAttackLayer3TimeseryListParamsProtocolTcp, cloudflare.RadarAttackLayer3TimeseryListParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
