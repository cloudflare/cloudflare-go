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

func TestHTTPLocationDeviceTypeGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Locations.DeviceType.Get(
		context.TODO(),
		radar.HTTPLocationDeviceTypeGetParamsDeviceTypeDesktop,
		radar.HTTPLocationDeviceTypeGetParams{
			ASN:           cloudflare.F([]string{"string", "string", "string"}),
			BotClass:      cloudflare.F([]radar.HTTPLocationDeviceTypeGetParamsBotClass{radar.HTTPLocationDeviceTypeGetParamsBotClassLikelyAutomated, radar.HTTPLocationDeviceTypeGetParamsBotClassLikelyHuman}),
			BrowserFamily: cloudflare.F([]radar.HTTPLocationDeviceTypeGetParamsBrowserFamily{radar.HTTPLocationDeviceTypeGetParamsBrowserFamilyChrome, radar.HTTPLocationDeviceTypeGetParamsBrowserFamilyEdge, radar.HTTPLocationDeviceTypeGetParamsBrowserFamilyFirefox}),
			Continent:     cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:     cloudflare.F([]string{"7d", "7d", "7d"}),
			DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			Format:        cloudflare.F(radar.HTTPLocationDeviceTypeGetParamsFormatJson),
			HTTPProtocol:  cloudflare.F([]radar.HTTPLocationDeviceTypeGetParamsHTTPProtocol{radar.HTTPLocationDeviceTypeGetParamsHTTPProtocolHTTP, radar.HTTPLocationDeviceTypeGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:   cloudflare.F([]radar.HTTPLocationDeviceTypeGetParamsHTTPVersion{radar.HTTPLocationDeviceTypeGetParamsHTTPVersionHttPv1, radar.HTTPLocationDeviceTypeGetParamsHTTPVersionHttPv2, radar.HTTPLocationDeviceTypeGetParamsHTTPVersionHttPv3}),
			IPVersion:     cloudflare.F([]radar.HTTPLocationDeviceTypeGetParamsIPVersion{radar.HTTPLocationDeviceTypeGetParamsIPVersionIPv4, radar.HTTPLocationDeviceTypeGetParamsIPVersionIPv6}),
			Limit:         cloudflare.F(int64(5)),
			Location:      cloudflare.F([]string{"string", "string", "string"}),
			Name:          cloudflare.F([]string{"string", "string", "string"}),
			OS:            cloudflare.F([]radar.HTTPLocationDeviceTypeGetParamsOS{radar.HTTPLocationDeviceTypeGetParamsOSWindows, radar.HTTPLocationDeviceTypeGetParamsOSMacosx, radar.HTTPLocationDeviceTypeGetParamsOSIos}),
			TLSVersion:    cloudflare.F([]radar.HTTPLocationDeviceTypeGetParamsTLSVersion{radar.HTTPLocationDeviceTypeGetParamsTLSVersionTlSv1_0, radar.HTTPLocationDeviceTypeGetParamsTLSVersionTlSv1_1, radar.HTTPLocationDeviceTypeGetParamsTLSVersionTlSv1_2}),
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
