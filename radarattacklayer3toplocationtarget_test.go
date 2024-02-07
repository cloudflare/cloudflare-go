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

func TestRadarAttackLayer3TopLocationTargetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Top.Locations.Target.List(context.TODO(), cloudflare.RadarAttackLayer3TopLocationTargetListParams{
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarAttackLayer3TopLocationTargetListParamsDateRange{cloudflare.RadarAttackLayer3TopLocationTargetListParamsDateRange1d, cloudflare.RadarAttackLayer3TopLocationTargetListParamsDateRange2d, cloudflare.RadarAttackLayer3TopLocationTargetListParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:    cloudflare.F(cloudflare.RadarAttackLayer3TopLocationTargetListParamsFormatJson),
		IPVersion: cloudflare.F([]cloudflare.RadarAttackLayer3TopLocationTargetListParamsIPVersion{cloudflare.RadarAttackLayer3TopLocationTargetListParamsIPVersionIPv4, cloudflare.RadarAttackLayer3TopLocationTargetListParamsIPVersionIPv6}),
		Limit:     cloudflare.F(int64(5)),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		Protocol:  cloudflare.F([]cloudflare.RadarAttackLayer3TopLocationTargetListParamsProtocol{cloudflare.RadarAttackLayer3TopLocationTargetListParamsProtocolUdp, cloudflare.RadarAttackLayer3TopLocationTargetListParamsProtocolTcp, cloudflare.RadarAttackLayer3TopLocationTargetListParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
