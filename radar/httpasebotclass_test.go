// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/radar"
)

func TestHTTPAseBotClassGetWithOptionalParams(t *testing.T) {
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
		radar.HTTPAseBotClassGetParamsBotClassLikelyAutomated,
		radar.HTTPAseBotClassGetParams{
			ASN:           cloudflare.F([]string{"string", "string", "string"}),
			BrowserFamily: cloudflare.F([]radar.HTTPAseBotClassGetParamsBrowserFamily{radar.HTTPAseBotClassGetParamsBrowserFamilyChrome, radar.HTTPAseBotClassGetParamsBrowserFamilyEdge, radar.HTTPAseBotClassGetParamsBrowserFamilyFirefox}),
			Continent:     cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:     cloudflare.F([]string{"7d", "7d", "7d"}),
			DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:    cloudflare.F([]radar.HTTPAseBotClassGetParamsDeviceType{radar.HTTPAseBotClassGetParamsDeviceTypeDesktop, radar.HTTPAseBotClassGetParamsDeviceTypeMobile, radar.HTTPAseBotClassGetParamsDeviceTypeOther}),
			Format:        cloudflare.F(radar.HTTPAseBotClassGetParamsFormatJson),
			HTTPProtocol:  cloudflare.F([]radar.HTTPAseBotClassGetParamsHTTPProtocol{radar.HTTPAseBotClassGetParamsHTTPProtocolHTTP, radar.HTTPAseBotClassGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:   cloudflare.F([]radar.HTTPAseBotClassGetParamsHTTPVersion{radar.HTTPAseBotClassGetParamsHTTPVersionHttPv1, radar.HTTPAseBotClassGetParamsHTTPVersionHttPv2, radar.HTTPAseBotClassGetParamsHTTPVersionHttPv3}),
			IPVersion:     cloudflare.F([]radar.HTTPAseBotClassGetParamsIPVersion{radar.HTTPAseBotClassGetParamsIPVersionIPv4, radar.HTTPAseBotClassGetParamsIPVersionIPv6}),
			Limit:         cloudflare.F(int64(5)),
			Location:      cloudflare.F([]string{"string", "string", "string"}),
			Name:          cloudflare.F([]string{"string", "string", "string"}),
			OS:            cloudflare.F([]radar.HTTPAseBotClassGetParamsOS{radar.HTTPAseBotClassGetParamsOSWindows, radar.HTTPAseBotClassGetParamsOSMacosx, radar.HTTPAseBotClassGetParamsOSIos}),
			TLSVersion:    cloudflare.F([]radar.HTTPAseBotClassGetParamsTLSVersion{radar.HTTPAseBotClassGetParamsTLSVersionTlSv1_0, radar.HTTPAseBotClassGetParamsTLSVersionTlSv1_1, radar.HTTPAseBotClassGetParamsTLSVersionTlSv1_2}),
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
