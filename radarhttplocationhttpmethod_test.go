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

func TestRadarHTTPLocationHTTPMethodGetWithOptionalParams(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Radar.HTTP.Locations.HTTPMethod.Get(
		context.TODO(),
		cloudflare.RadarHTTPLocationHTTPMethodGetParamsHTTPVersionHttPv1,
		cloudflare.RadarHTTPLocationHTTPMethodGetParams{
			ASN:          cloudflare.F([]string{"string", "string", "string"}),
			BotClass:     cloudflare.F([]cloudflare.RadarHTTPLocationHTTPMethodGetParamsBotClass{cloudflare.RadarHTTPLocationHTTPMethodGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPLocationHTTPMethodGetParamsBotClassLikelyHuman}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPLocationHTTPMethodGetParamsDateRange{cloudflare.RadarHTTPLocationHTTPMethodGetParamsDateRange1d, cloudflare.RadarHTTPLocationHTTPMethodGetParamsDateRange2d, cloudflare.RadarHTTPLocationHTTPMethodGetParamsDateRange7d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]cloudflare.RadarHTTPLocationHTTPMethodGetParamsDeviceType{cloudflare.RadarHTTPLocationHTTPMethodGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPLocationHTTPMethodGetParamsDeviceTypeMobile, cloudflare.RadarHTTPLocationHTTPMethodGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(cloudflare.RadarHTTPLocationHTTPMethodGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPLocationHTTPMethodGetParamsHTTPProtocol{cloudflare.RadarHTTPLocationHTTPMethodGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPLocationHTTPMethodGetParamsHTTPProtocolHTTPS}),
			IPVersion:    cloudflare.F([]cloudflare.RadarHTTPLocationHTTPMethodGetParamsIPVersion{cloudflare.RadarHTTPLocationHTTPMethodGetParamsIPVersionIPv4, cloudflare.RadarHTTPLocationHTTPMethodGetParamsIPVersionIPv6}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"string", "string", "string"}),
			Name:         cloudflare.F([]string{"string", "string", "string"}),
			OS:           cloudflare.F([]cloudflare.RadarHTTPLocationHTTPMethodGetParamsOS{cloudflare.RadarHTTPLocationHTTPMethodGetParamsOSWindows, cloudflare.RadarHTTPLocationHTTPMethodGetParamsOSMacosx, cloudflare.RadarHTTPLocationHTTPMethodGetParamsOSIos}),
			TLSVersion:   cloudflare.F([]cloudflare.RadarHTTPLocationHTTPMethodGetParamsTLSVersion{cloudflare.RadarHTTPLocationHTTPMethodGetParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPLocationHTTPMethodGetParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPLocationHTTPMethodGetParamsTLSVersionTlSv1_2}),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
