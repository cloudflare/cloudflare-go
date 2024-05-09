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

func TestHTTPAseOSGetWithOptionalParams(t *testing.T) {
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
		radar.HTTPAseOSGetParamsOSWindows,
		radar.HTTPAseOSGetParams{
			ASN:           cloudflare.F([]string{"string", "string", "string"}),
			BotClass:      cloudflare.F([]radar.HTTPAseOSGetParamsBotClass{radar.HTTPAseOSGetParamsBotClassLikelyAutomated, radar.HTTPAseOSGetParamsBotClassLikelyHuman}),
			BrowserFamily: cloudflare.F([]radar.HTTPAseOSGetParamsBrowserFamily{radar.HTTPAseOSGetParamsBrowserFamilyChrome, radar.HTTPAseOSGetParamsBrowserFamilyEdge, radar.HTTPAseOSGetParamsBrowserFamilyFirefox}),
			Continent:     cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:     cloudflare.F([]radar.HTTPAseOSGetParamsDateRange{radar.HTTPAseOSGetParamsDateRange1d, radar.HTTPAseOSGetParamsDateRange2d, radar.HTTPAseOSGetParamsDateRange7d}),
			DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:    cloudflare.F([]radar.HTTPAseOSGetParamsDeviceType{radar.HTTPAseOSGetParamsDeviceTypeDesktop, radar.HTTPAseOSGetParamsDeviceTypeMobile, radar.HTTPAseOSGetParamsDeviceTypeOther}),
			Format:        cloudflare.F(radar.HTTPAseOSGetParamsFormatJson),
			HTTPProtocol:  cloudflare.F([]radar.HTTPAseOSGetParamsHTTPProtocol{radar.HTTPAseOSGetParamsHTTPProtocolHTTP, radar.HTTPAseOSGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:   cloudflare.F([]radar.HTTPAseOSGetParamsHTTPVersion{radar.HTTPAseOSGetParamsHTTPVersionHttPv1, radar.HTTPAseOSGetParamsHTTPVersionHttPv2, radar.HTTPAseOSGetParamsHTTPVersionHttPv3}),
			IPVersion:     cloudflare.F([]radar.HTTPAseOSGetParamsIPVersion{radar.HTTPAseOSGetParamsIPVersionIPv4, radar.HTTPAseOSGetParamsIPVersionIPv6}),
			Limit:         cloudflare.F(int64(5)),
			Location:      cloudflare.F([]string{"string", "string", "string"}),
			Name:          cloudflare.F([]string{"string", "string", "string"}),
			TLSVersion:    cloudflare.F([]radar.HTTPAseOSGetParamsTLSVersion{radar.HTTPAseOSGetParamsTLSVersionTlSv1_0, radar.HTTPAseOSGetParamsTLSVersionTlSv1_1, radar.HTTPAseOSGetParamsTLSVersionTlSv1_2}),
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
