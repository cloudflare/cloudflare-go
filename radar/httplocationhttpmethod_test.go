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

func TestHTTPLocationHTTPMethodGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Locations.HTTPMethod.Get(
		context.TODO(),
		radar.HTTPLocationHTTPMethodGetParamsHTTPVersionHttPv1,
		radar.HTTPLocationHTTPMethodGetParams{
			ASN:           cloudflare.F([]string{"string", "string", "string"}),
			BotClass:      cloudflare.F([]radar.HTTPLocationHTTPMethodGetParamsBotClass{radar.HTTPLocationHTTPMethodGetParamsBotClassLikelyAutomated, radar.HTTPLocationHTTPMethodGetParamsBotClassLikelyHuman}),
			BrowserFamily: cloudflare.F([]radar.HTTPLocationHTTPMethodGetParamsBrowserFamily{radar.HTTPLocationHTTPMethodGetParamsBrowserFamilyChrome, radar.HTTPLocationHTTPMethodGetParamsBrowserFamilyEdge, radar.HTTPLocationHTTPMethodGetParamsBrowserFamilyFirefox}),
			Continent:     cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:     cloudflare.F([]radar.HTTPLocationHTTPMethodGetParamsDateRange{radar.HTTPLocationHTTPMethodGetParamsDateRange1d, radar.HTTPLocationHTTPMethodGetParamsDateRange2d, radar.HTTPLocationHTTPMethodGetParamsDateRange7d}),
			DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:    cloudflare.F([]radar.HTTPLocationHTTPMethodGetParamsDeviceType{radar.HTTPLocationHTTPMethodGetParamsDeviceTypeDesktop, radar.HTTPLocationHTTPMethodGetParamsDeviceTypeMobile, radar.HTTPLocationHTTPMethodGetParamsDeviceTypeOther}),
			Format:        cloudflare.F(radar.HTTPLocationHTTPMethodGetParamsFormatJson),
			HTTPProtocol:  cloudflare.F([]radar.HTTPLocationHTTPMethodGetParamsHTTPProtocol{radar.HTTPLocationHTTPMethodGetParamsHTTPProtocolHTTP, radar.HTTPLocationHTTPMethodGetParamsHTTPProtocolHTTPS}),
			IPVersion:     cloudflare.F([]radar.HTTPLocationHTTPMethodGetParamsIPVersion{radar.HTTPLocationHTTPMethodGetParamsIPVersionIPv4, radar.HTTPLocationHTTPMethodGetParamsIPVersionIPv6}),
			Limit:         cloudflare.F(int64(5)),
			Location:      cloudflare.F([]string{"string", "string", "string"}),
			Name:          cloudflare.F([]string{"string", "string", "string"}),
			OS:            cloudflare.F([]radar.HTTPLocationHTTPMethodGetParamsOS{radar.HTTPLocationHTTPMethodGetParamsOSWindows, radar.HTTPLocationHTTPMethodGetParamsOSMacosx, radar.HTTPLocationHTTPMethodGetParamsOSIos}),
			TLSVersion:    cloudflare.F([]radar.HTTPLocationHTTPMethodGetParamsTLSVersion{radar.HTTPLocationHTTPMethodGetParamsTLSVersionTlSv1_0, radar.HTTPLocationHTTPMethodGetParamsTLSVersionTlSv1_1, radar.HTTPLocationHTTPMethodGetParamsTLSVersionTlSv1_2}),
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
