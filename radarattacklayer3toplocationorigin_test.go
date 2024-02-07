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

func TestRadarAttackLayer3TopLocationOriginListWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Radar.Attacks.Layer3.Top.Locations.Origin.List(context.TODO(), cloudflare.RadarAttackLayer3TopLocationOriginListParams{
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarAttackLayer3TopLocationOriginListParamsDateRange{cloudflare.RadarAttackLayer3TopLocationOriginListParamsDateRange1d, cloudflare.RadarAttackLayer3TopLocationOriginListParamsDateRange2d, cloudflare.RadarAttackLayer3TopLocationOriginListParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:    cloudflare.F(cloudflare.RadarAttackLayer3TopLocationOriginListParamsFormatJson),
		IPVersion: cloudflare.F([]cloudflare.RadarAttackLayer3TopLocationOriginListParamsIPVersion{cloudflare.RadarAttackLayer3TopLocationOriginListParamsIPVersionIPv4, cloudflare.RadarAttackLayer3TopLocationOriginListParamsIPVersionIPv6}),
		Limit:     cloudflare.F(int64(5)),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		Protocol:  cloudflare.F([]cloudflare.RadarAttackLayer3TopLocationOriginListParamsProtocol{cloudflare.RadarAttackLayer3TopLocationOriginListParamsProtocolUdp, cloudflare.RadarAttackLayer3TopLocationOriginListParamsProtocolTcp, cloudflare.RadarAttackLayer3TopLocationOriginListParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
