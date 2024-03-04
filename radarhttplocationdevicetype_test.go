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

func TestRadarHTTPLocationDeviceTypeGetWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Radar.HTTP.Locations.DeviceType.Get(
		context.TODO(),
		cloudflare.RadarHTTPLocationDeviceTypeGetParamsDeviceTypeDesktop,
		cloudflare.RadarHTTPLocationDeviceTypeGetParams{
			ASN:          cloudflare.F([]string{"string", "string", "string"}),
			BotClass:     cloudflare.F([]cloudflare.RadarHTTPLocationDeviceTypeGetParamsBotClass{cloudflare.RadarHTTPLocationDeviceTypeGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPLocationDeviceTypeGetParamsBotClassLikelyHuman}),
			Continent:    cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPLocationDeviceTypeGetParamsDateRange{cloudflare.RadarHTTPLocationDeviceTypeGetParamsDateRange1d, cloudflare.RadarHTTPLocationDeviceTypeGetParamsDateRange2d, cloudflare.RadarHTTPLocationDeviceTypeGetParamsDateRange7d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			Format:       cloudflare.F(cloudflare.RadarHTTPLocationDeviceTypeGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPLocationDeviceTypeGetParamsHTTPProtocol{cloudflare.RadarHTTPLocationDeviceTypeGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPLocationDeviceTypeGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPLocationDeviceTypeGetParamsHTTPVersion{cloudflare.RadarHTTPLocationDeviceTypeGetParamsHTTPVersionHttPv1, cloudflare.RadarHTTPLocationDeviceTypeGetParamsHTTPVersionHttPv2, cloudflare.RadarHTTPLocationDeviceTypeGetParamsHTTPVersionHttPv3}),
			IPVersion:    cloudflare.F([]cloudflare.RadarHTTPLocationDeviceTypeGetParamsIPVersion{cloudflare.RadarHTTPLocationDeviceTypeGetParamsIPVersionIPv4, cloudflare.RadarHTTPLocationDeviceTypeGetParamsIPVersionIPv6}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"string", "string", "string"}),
			Name:         cloudflare.F([]string{"string", "string", "string"}),
			OS:           cloudflare.F([]cloudflare.RadarHTTPLocationDeviceTypeGetParamsOS{cloudflare.RadarHTTPLocationDeviceTypeGetParamsOSWindows, cloudflare.RadarHTTPLocationDeviceTypeGetParamsOSMacosx, cloudflare.RadarHTTPLocationDeviceTypeGetParamsOSIos}),
			TLSVersion:   cloudflare.F([]cloudflare.RadarHTTPLocationDeviceTypeGetParamsTLSVersion{cloudflare.RadarHTTPLocationDeviceTypeGetParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPLocationDeviceTypeGetParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPLocationDeviceTypeGetParamsTLSVersionTlSv1_2}),
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
