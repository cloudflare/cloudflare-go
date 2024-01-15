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

func TestRadarHTTPTopLocationOGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Tops.Locations.Os.Get(
		context.TODO(),
		cloudflare.RadarHTTPTopLocationOGetParamsOsWindows,
		cloudflare.RadarHTTPTopLocationOGetParams{
			ASN:          cloudflare.F([]string{"string", "string", "string"}),
			BotClass:     cloudflare.F([]cloudflare.RadarHTTPTopLocationOGetParamsBotClass{cloudflare.RadarHTTPTopLocationOGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTopLocationOGetParamsBotClassLikelyHuman}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPTopLocationOGetParamsDateRange{cloudflare.RadarHTTPTopLocationOGetParamsDateRange1d, cloudflare.RadarHTTPTopLocationOGetParamsDateRange2d, cloudflare.RadarHTTPTopLocationOGetParamsDateRange7d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTopLocationOGetParamsDeviceType{cloudflare.RadarHTTPTopLocationOGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPTopLocationOGetParamsDeviceTypeMobile, cloudflare.RadarHTTPTopLocationOGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(cloudflare.RadarHTTPTopLocationOGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTopLocationOGetParamsHTTPProtocol{cloudflare.RadarHTTPTopLocationOGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTopLocationOGetParamsHTTPProtocolHTTPs}),
			HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTopLocationOGetParamsHTTPVersion{cloudflare.RadarHTTPTopLocationOGetParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTopLocationOGetParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTopLocationOGetParamsHTTPVersionHttPv3}),
			IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTopLocationOGetParamsIPVersion{cloudflare.RadarHTTPTopLocationOGetParamsIPVersionIPv4, cloudflare.RadarHTTPTopLocationOGetParamsIPVersionIPv6}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"string", "string", "string"}),
			Name:         cloudflare.F([]string{"string", "string", "string"}),
			TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPTopLocationOGetParamsTlsVersion{cloudflare.RadarHTTPTopLocationOGetParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTopLocationOGetParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTopLocationOGetParamsTlsVersionTlSv1_2}),
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
