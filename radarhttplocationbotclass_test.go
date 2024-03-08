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

func TestRadarHTTPLocationBotClassGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Locations.BotClass.Get(
		context.TODO(),
		cloudflare.RadarHTTPLocationBotClassGetParamsBotClassLikelyAutomated,
		cloudflare.RadarHTTPLocationBotClassGetParams{
			ASN:          cloudflare.F([]string{"string", "string", "string"}),
			Continent:    cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPLocationBotClassGetParamsDateRange{cloudflare.RadarHTTPLocationBotClassGetParamsDateRange1d, cloudflare.RadarHTTPLocationBotClassGetParamsDateRange2d, cloudflare.RadarHTTPLocationBotClassGetParamsDateRange7d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]cloudflare.RadarHTTPLocationBotClassGetParamsDeviceType{cloudflare.RadarHTTPLocationBotClassGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPLocationBotClassGetParamsDeviceTypeMobile, cloudflare.RadarHTTPLocationBotClassGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(cloudflare.RadarHTTPLocationBotClassGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPLocationBotClassGetParamsHTTPProtocol{cloudflare.RadarHTTPLocationBotClassGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPLocationBotClassGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPLocationBotClassGetParamsHTTPVersion{cloudflare.RadarHTTPLocationBotClassGetParamsHTTPVersionHttPv1, cloudflare.RadarHTTPLocationBotClassGetParamsHTTPVersionHttPv2, cloudflare.RadarHTTPLocationBotClassGetParamsHTTPVersionHttPv3}),
			IPVersion:    cloudflare.F([]cloudflare.RadarHTTPLocationBotClassGetParamsIPVersion{cloudflare.RadarHTTPLocationBotClassGetParamsIPVersionIPv4, cloudflare.RadarHTTPLocationBotClassGetParamsIPVersionIPv6}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"string", "string", "string"}),
			Name:         cloudflare.F([]string{"string", "string", "string"}),
			OS:           cloudflare.F([]cloudflare.RadarHTTPLocationBotClassGetParamsOS{cloudflare.RadarHTTPLocationBotClassGetParamsOSWindows, cloudflare.RadarHTTPLocationBotClassGetParamsOSMacosx, cloudflare.RadarHTTPLocationBotClassGetParamsOSIos}),
			TLSVersion:   cloudflare.F([]cloudflare.RadarHTTPLocationBotClassGetParamsTLSVersion{cloudflare.RadarHTTPLocationBotClassGetParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPLocationBotClassGetParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPLocationBotClassGetParamsTLSVersionTlSv1_2}),
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
