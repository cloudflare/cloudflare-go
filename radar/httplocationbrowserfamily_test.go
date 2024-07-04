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

func TestHTTPLocationBrowserFamilyGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Locations.BrowserFamily.Get(
		context.TODO(),
		radar.HTTPLocationBrowserFamilyGetParamsBrowserFamilyChrome,
		radar.HTTPLocationBrowserFamilyGetParams{
			ASN:          cloudflare.F([]string{"string", "string", "string"}),
			BotClass:     cloudflare.F([]radar.HTTPLocationBrowserFamilyGetParamsBotClass{radar.HTTPLocationBrowserFamilyGetParamsBotClassLikelyAutomated, radar.HTTPLocationBrowserFamilyGetParamsBotClassLikelyHuman}),
			Continent:    cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]string{"7d", "7d", "7d"}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]radar.HTTPLocationBrowserFamilyGetParamsDeviceType{radar.HTTPLocationBrowserFamilyGetParamsDeviceTypeDesktop, radar.HTTPLocationBrowserFamilyGetParamsDeviceTypeMobile, radar.HTTPLocationBrowserFamilyGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(radar.HTTPLocationBrowserFamilyGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]radar.HTTPLocationBrowserFamilyGetParamsHTTPProtocol{radar.HTTPLocationBrowserFamilyGetParamsHTTPProtocolHTTP, radar.HTTPLocationBrowserFamilyGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:  cloudflare.F([]radar.HTTPLocationBrowserFamilyGetParamsHTTPVersion{radar.HTTPLocationBrowserFamilyGetParamsHTTPVersionHttPv1, radar.HTTPLocationBrowserFamilyGetParamsHTTPVersionHttPv2, radar.HTTPLocationBrowserFamilyGetParamsHTTPVersionHttPv3}),
			IPVersion:    cloudflare.F([]radar.HTTPLocationBrowserFamilyGetParamsIPVersion{radar.HTTPLocationBrowserFamilyGetParamsIPVersionIPv4, radar.HTTPLocationBrowserFamilyGetParamsIPVersionIPv6}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"string", "string", "string"}),
			Name:         cloudflare.F([]string{"string", "string", "string"}),
			OS:           cloudflare.F([]radar.HTTPLocationBrowserFamilyGetParamsOS{radar.HTTPLocationBrowserFamilyGetParamsOSWindows, radar.HTTPLocationBrowserFamilyGetParamsOSMacosx, radar.HTTPLocationBrowserFamilyGetParamsOSIos}),
			TLSVersion:   cloudflare.F([]radar.HTTPLocationBrowserFamilyGetParamsTLSVersion{radar.HTTPLocationBrowserFamilyGetParamsTLSVersionTlSv1_0, radar.HTTPLocationBrowserFamilyGetParamsTLSVersionTlSv1_1, radar.HTTPLocationBrowserFamilyGetParamsTLSVersionTlSv1_2}),
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
