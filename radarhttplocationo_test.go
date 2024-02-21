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

func TestRadarHTTPLocationOGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Locations.Os.Get(
		context.TODO(),
		cloudflare.RadarHTTPLocationOGetParamsOsWindows,
		cloudflare.RadarHTTPLocationOGetParams{
			Asn:          cloudflare.F([]string{"string", "string", "string"}),
			BotClass:     cloudflare.F([]cloudflare.RadarHTTPLocationOGetParamsBotClass{cloudflare.RadarHTTPLocationOGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPLocationOGetParamsBotClassLikelyHuman}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPLocationOGetParamsDateRange{cloudflare.RadarHTTPLocationOGetParamsDateRange1d, cloudflare.RadarHTTPLocationOGetParamsDateRange2d, cloudflare.RadarHTTPLocationOGetParamsDateRange7d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]cloudflare.RadarHTTPLocationOGetParamsDeviceType{cloudflare.RadarHTTPLocationOGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPLocationOGetParamsDeviceTypeMobile, cloudflare.RadarHTTPLocationOGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(cloudflare.RadarHTTPLocationOGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPLocationOGetParamsHTTPProtocol{cloudflare.RadarHTTPLocationOGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPLocationOGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPLocationOGetParamsHTTPVersion{cloudflare.RadarHTTPLocationOGetParamsHTTPVersionHttPv1, cloudflare.RadarHTTPLocationOGetParamsHTTPVersionHttPv2, cloudflare.RadarHTTPLocationOGetParamsHTTPVersionHttPv3}),
			IPVersion:    cloudflare.F([]cloudflare.RadarHTTPLocationOGetParamsIPVersion{cloudflare.RadarHTTPLocationOGetParamsIPVersionIPv4, cloudflare.RadarHTTPLocationOGetParamsIPVersionIPv6}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"string", "string", "string"}),
			Name:         cloudflare.F([]string{"string", "string", "string"}),
			TLSVersion:   cloudflare.F([]cloudflare.RadarHTTPLocationOGetParamsTLSVersion{cloudflare.RadarHTTPLocationOGetParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPLocationOGetParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPLocationOGetParamsTLSVersionTlSv1_2}),
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
