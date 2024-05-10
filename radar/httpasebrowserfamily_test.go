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

func TestHTTPAseBrowserFamilyGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Ases.BrowserFamily.Get(
		context.TODO(),
		radar.HTTPAseBrowserFamilyGetParamsBrowserFamilyChrome,
		radar.HTTPAseBrowserFamilyGetParams{
			ASN:          cloudflare.F([]string{"string", "string", "string"}),
			BotClass:     cloudflare.F([]radar.HTTPAseBrowserFamilyGetParamsBotClass{radar.HTTPAseBrowserFamilyGetParamsBotClassLikelyAutomated, radar.HTTPAseBrowserFamilyGetParamsBotClassLikelyHuman}),
			Continent:    cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]radar.HTTPAseBrowserFamilyGetParamsDateRange{radar.HTTPAseBrowserFamilyGetParamsDateRange1d, radar.HTTPAseBrowserFamilyGetParamsDateRange2d, radar.HTTPAseBrowserFamilyGetParamsDateRange7d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]radar.HTTPAseBrowserFamilyGetParamsDeviceType{radar.HTTPAseBrowserFamilyGetParamsDeviceTypeDesktop, radar.HTTPAseBrowserFamilyGetParamsDeviceTypeMobile, radar.HTTPAseBrowserFamilyGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(radar.HTTPAseBrowserFamilyGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]radar.HTTPAseBrowserFamilyGetParamsHTTPProtocol{radar.HTTPAseBrowserFamilyGetParamsHTTPProtocolHTTP, radar.HTTPAseBrowserFamilyGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:  cloudflare.F([]radar.HTTPAseBrowserFamilyGetParamsHTTPVersion{radar.HTTPAseBrowserFamilyGetParamsHTTPVersionHttPv1, radar.HTTPAseBrowserFamilyGetParamsHTTPVersionHttPv2, radar.HTTPAseBrowserFamilyGetParamsHTTPVersionHttPv3}),
			IPVersion:    cloudflare.F([]radar.HTTPAseBrowserFamilyGetParamsIPVersion{radar.HTTPAseBrowserFamilyGetParamsIPVersionIPv4, radar.HTTPAseBrowserFamilyGetParamsIPVersionIPv6}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"string", "string", "string"}),
			Name:         cloudflare.F([]string{"string", "string", "string"}),
			OS:           cloudflare.F([]radar.HTTPAseBrowserFamilyGetParamsOS{radar.HTTPAseBrowserFamilyGetParamsOSWindows, radar.HTTPAseBrowserFamilyGetParamsOSMacosx, radar.HTTPAseBrowserFamilyGetParamsOSIos}),
			TLSVersion:   cloudflare.F([]radar.HTTPAseBrowserFamilyGetParamsTLSVersion{radar.HTTPAseBrowserFamilyGetParamsTLSVersionTlSv1_0, radar.HTTPAseBrowserFamilyGetParamsTLSVersionTlSv1_1, radar.HTTPAseBrowserFamilyGetParamsTLSVersionTlSv1_2}),
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
