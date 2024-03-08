// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
)

func TestRadarHTTPAseOSGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Ases.OS.Get(
		context.TODO(),
		cloudflare.RadarHTTPAseOSGetParamsOSWindows,
		cloudflare.RadarHTTPAseOSGetParams{
			ASN:          cloudflare.F([]string{"string", "string", "string"}),
			BotClass:     cloudflare.F([]cloudflare.RadarHTTPAseOSGetParamsBotClass{cloudflare.RadarHTTPAseOSGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPAseOSGetParamsBotClassLikelyHuman}),
			Continent:    cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPAseOSGetParamsDateRange{cloudflare.RadarHTTPAseOSGetParamsDateRange1d, cloudflare.RadarHTTPAseOSGetParamsDateRange2d, cloudflare.RadarHTTPAseOSGetParamsDateRange7d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]cloudflare.RadarHTTPAseOSGetParamsDeviceType{cloudflare.RadarHTTPAseOSGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPAseOSGetParamsDeviceTypeMobile, cloudflare.RadarHTTPAseOSGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(cloudflare.RadarHTTPAseOSGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPAseOSGetParamsHTTPProtocol{cloudflare.RadarHTTPAseOSGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPAseOSGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPAseOSGetParamsHTTPVersion{cloudflare.RadarHTTPAseOSGetParamsHTTPVersionHttPv1, cloudflare.RadarHTTPAseOSGetParamsHTTPVersionHttPv2, cloudflare.RadarHTTPAseOSGetParamsHTTPVersionHttPv3}),
			IPVersion:    cloudflare.F([]cloudflare.RadarHTTPAseOSGetParamsIPVersion{cloudflare.RadarHTTPAseOSGetParamsIPVersionIPv4, cloudflare.RadarHTTPAseOSGetParamsIPVersionIPv6}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"string", "string", "string"}),
			Name:         cloudflare.F([]string{"string", "string", "string"}),
			TLSVersion:   cloudflare.F([]cloudflare.RadarHTTPAseOSGetParamsTLSVersion{cloudflare.RadarHTTPAseOSGetParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPAseOSGetParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPAseOSGetParamsTLSVersionTlSv1_2}),
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
