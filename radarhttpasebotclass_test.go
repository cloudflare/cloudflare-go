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

func TestRadarHTTPAseBotClassGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Ases.BotClass.Get(
		context.TODO(),
		cloudflare.RadarHTTPAseBotClassGetParamsBotClassLikelyAutomated,
		cloudflare.RadarHTTPAseBotClassGetParams{
			ASN:          cloudflare.F([]string{"string", "string", "string"}),
			Continent:    cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPAseBotClassGetParamsDateRange{cloudflare.RadarHTTPAseBotClassGetParamsDateRange1d, cloudflare.RadarHTTPAseBotClassGetParamsDateRange2d, cloudflare.RadarHTTPAseBotClassGetParamsDateRange7d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]cloudflare.RadarHTTPAseBotClassGetParamsDeviceType{cloudflare.RadarHTTPAseBotClassGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPAseBotClassGetParamsDeviceTypeMobile, cloudflare.RadarHTTPAseBotClassGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(cloudflare.RadarHTTPAseBotClassGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPAseBotClassGetParamsHTTPProtocol{cloudflare.RadarHTTPAseBotClassGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPAseBotClassGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPAseBotClassGetParamsHTTPVersion{cloudflare.RadarHTTPAseBotClassGetParamsHTTPVersionHttPv1, cloudflare.RadarHTTPAseBotClassGetParamsHTTPVersionHttPv2, cloudflare.RadarHTTPAseBotClassGetParamsHTTPVersionHttPv3}),
			IPVersion:    cloudflare.F([]cloudflare.RadarHTTPAseBotClassGetParamsIPVersion{cloudflare.RadarHTTPAseBotClassGetParamsIPVersionIPv4, cloudflare.RadarHTTPAseBotClassGetParamsIPVersionIPv6}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"string", "string", "string"}),
			Name:         cloudflare.F([]string{"string", "string", "string"}),
			OS:           cloudflare.F([]cloudflare.RadarHTTPAseBotClassGetParamsOS{cloudflare.RadarHTTPAseBotClassGetParamsOSWindows, cloudflare.RadarHTTPAseBotClassGetParamsOSMacosx, cloudflare.RadarHTTPAseBotClassGetParamsOSIos}),
			TLSVersion:   cloudflare.F([]cloudflare.RadarHTTPAseBotClassGetParamsTLSVersion{cloudflare.RadarHTTPAseBotClassGetParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPAseBotClassGetParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPAseBotClassGetParamsTLSVersionTlSv1_2}),
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
