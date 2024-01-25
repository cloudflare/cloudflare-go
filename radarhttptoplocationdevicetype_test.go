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

func TestRadarHTTPTopLocationDeviceTypeGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Tops.Locations.DeviceTypes.Get(
		context.TODO(),
		cloudflare.RadarHTTPTopLocationDeviceTypeGetParamsDeviceTypeDesktop,
		cloudflare.RadarHTTPTopLocationDeviceTypeGetParams{
			ASN:          cloudflare.F([]string{"string", "string", "string"}),
			BotClass:     cloudflare.F([]cloudflare.RadarHTTPTopLocationDeviceTypeGetParamsBotClass{cloudflare.RadarHTTPTopLocationDeviceTypeGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTopLocationDeviceTypeGetParamsBotClassLikelyHuman}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPTopLocationDeviceTypeGetParamsDateRange{cloudflare.RadarHTTPTopLocationDeviceTypeGetParamsDateRange1d, cloudflare.RadarHTTPTopLocationDeviceTypeGetParamsDateRange2d, cloudflare.RadarHTTPTopLocationDeviceTypeGetParamsDateRange7d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			Format:       cloudflare.F(cloudflare.RadarHTTPTopLocationDeviceTypeGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTopLocationDeviceTypeGetParamsHTTPProtocol{cloudflare.RadarHTTPTopLocationDeviceTypeGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTopLocationDeviceTypeGetParamsHTTPProtocolHTTPs}),
			HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTopLocationDeviceTypeGetParamsHTTPVersion{cloudflare.RadarHTTPTopLocationDeviceTypeGetParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTopLocationDeviceTypeGetParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTopLocationDeviceTypeGetParamsHTTPVersionHttPv3}),
			IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTopLocationDeviceTypeGetParamsIPVersion{cloudflare.RadarHTTPTopLocationDeviceTypeGetParamsIPVersionIPv4, cloudflare.RadarHTTPTopLocationDeviceTypeGetParamsIPVersionIPv6}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"string", "string", "string"}),
			Name:         cloudflare.F([]string{"string", "string", "string"}),
			Os:           cloudflare.F([]cloudflare.RadarHTTPTopLocationDeviceTypeGetParamsO{cloudflare.RadarHTTPTopLocationDeviceTypeGetParamsOWindows, cloudflare.RadarHTTPTopLocationDeviceTypeGetParamsOMacosx, cloudflare.RadarHTTPTopLocationDeviceTypeGetParamsOIos}),
			TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPTopLocationDeviceTypeGetParamsTlsVersion{cloudflare.RadarHTTPTopLocationDeviceTypeGetParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTopLocationDeviceTypeGetParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTopLocationDeviceTypeGetParamsTlsVersionTlSv1_2}),
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
