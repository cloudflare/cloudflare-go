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

func TestRadarHTTPAseTLSVersionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Ases.TLSVersion.Get(
		context.TODO(),
		cloudflare.RadarHTTPAseTLSVersionGetParamsTLSVersionTlSv1_0,
		cloudflare.RadarHTTPAseTLSVersionGetParams{
			ASN:          cloudflare.F([]string{"string", "string", "string"}),
			BotClass:     cloudflare.F([]cloudflare.RadarHTTPAseTLSVersionGetParamsBotClass{cloudflare.RadarHTTPAseTLSVersionGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPAseTLSVersionGetParamsBotClassLikelyHuman}),
			Continent:    cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPAseTLSVersionGetParamsDateRange{cloudflare.RadarHTTPAseTLSVersionGetParamsDateRange1d, cloudflare.RadarHTTPAseTLSVersionGetParamsDateRange2d, cloudflare.RadarHTTPAseTLSVersionGetParamsDateRange7d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]cloudflare.RadarHTTPAseTLSVersionGetParamsDeviceType{cloudflare.RadarHTTPAseTLSVersionGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPAseTLSVersionGetParamsDeviceTypeMobile, cloudflare.RadarHTTPAseTLSVersionGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(cloudflare.RadarHTTPAseTLSVersionGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPAseTLSVersionGetParamsHTTPProtocol{cloudflare.RadarHTTPAseTLSVersionGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPAseTLSVersionGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPAseTLSVersionGetParamsHTTPVersion{cloudflare.RadarHTTPAseTLSVersionGetParamsHTTPVersionHttPv1, cloudflare.RadarHTTPAseTLSVersionGetParamsHTTPVersionHttPv2, cloudflare.RadarHTTPAseTLSVersionGetParamsHTTPVersionHttPv3}),
			IPVersion:    cloudflare.F([]cloudflare.RadarHTTPAseTLSVersionGetParamsIPVersion{cloudflare.RadarHTTPAseTLSVersionGetParamsIPVersionIPv4, cloudflare.RadarHTTPAseTLSVersionGetParamsIPVersionIPv6}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"string", "string", "string"}),
			Name:         cloudflare.F([]string{"string", "string", "string"}),
			OS:           cloudflare.F([]cloudflare.RadarHTTPAseTLSVersionGetParamsOS{cloudflare.RadarHTTPAseTLSVersionGetParamsOSWindows, cloudflare.RadarHTTPAseTLSVersionGetParamsOSMacosx, cloudflare.RadarHTTPAseTLSVersionGetParamsOSIos}),
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
