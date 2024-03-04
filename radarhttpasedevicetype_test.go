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

func TestRadarHTTPAseDeviceTypeGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Ases.DeviceType.Get(
		context.TODO(),
		cloudflare.RadarHTTPAseDeviceTypeGetParamsDeviceTypeDesktop,
		cloudflare.RadarHTTPAseDeviceTypeGetParams{
			ASN:          cloudflare.F([]string{"string", "string", "string"}),
			BotClass:     cloudflare.F([]cloudflare.RadarHTTPAseDeviceTypeGetParamsBotClass{cloudflare.RadarHTTPAseDeviceTypeGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPAseDeviceTypeGetParamsBotClassLikelyHuman}),
			Continent:    cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPAseDeviceTypeGetParamsDateRange{cloudflare.RadarHTTPAseDeviceTypeGetParamsDateRange1d, cloudflare.RadarHTTPAseDeviceTypeGetParamsDateRange2d, cloudflare.RadarHTTPAseDeviceTypeGetParamsDateRange7d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			Format:       cloudflare.F(cloudflare.RadarHTTPAseDeviceTypeGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPAseDeviceTypeGetParamsHTTPProtocol{cloudflare.RadarHTTPAseDeviceTypeGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPAseDeviceTypeGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPAseDeviceTypeGetParamsHTTPVersion{cloudflare.RadarHTTPAseDeviceTypeGetParamsHTTPVersionHttPv1, cloudflare.RadarHTTPAseDeviceTypeGetParamsHTTPVersionHttPv2, cloudflare.RadarHTTPAseDeviceTypeGetParamsHTTPVersionHttPv3}),
			IPVersion:    cloudflare.F([]cloudflare.RadarHTTPAseDeviceTypeGetParamsIPVersion{cloudflare.RadarHTTPAseDeviceTypeGetParamsIPVersionIPv4, cloudflare.RadarHTTPAseDeviceTypeGetParamsIPVersionIPv6}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"string", "string", "string"}),
			Name:         cloudflare.F([]string{"string", "string", "string"}),
			OS:           cloudflare.F([]cloudflare.RadarHTTPAseDeviceTypeGetParamsOS{cloudflare.RadarHTTPAseDeviceTypeGetParamsOSWindows, cloudflare.RadarHTTPAseDeviceTypeGetParamsOSMacosx, cloudflare.RadarHTTPAseDeviceTypeGetParamsOSIos}),
			TLSVersion:   cloudflare.F([]cloudflare.RadarHTTPAseDeviceTypeGetParamsTLSVersion{cloudflare.RadarHTTPAseDeviceTypeGetParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPAseDeviceTypeGetParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPAseDeviceTypeGetParamsTLSVersionTlSv1_2}),
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
