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

func TestHTTPLocationTLSVersionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Locations.TLSVersion.Get(
		context.TODO(),
		radar.HTTPLocationTLSVersionGetParamsTLSVersionTlSv1_0,
		radar.HTTPLocationTLSVersionGetParams{
			ASN:           cloudflare.F([]string{"string", "string", "string"}),
			BotClass:      cloudflare.F([]radar.HTTPLocationTLSVersionGetParamsBotClass{radar.HTTPLocationTLSVersionGetParamsBotClassLikelyAutomated, radar.HTTPLocationTLSVersionGetParamsBotClassLikelyHuman}),
			BrowserFamily: cloudflare.F([]radar.HTTPLocationTLSVersionGetParamsBrowserFamily{radar.HTTPLocationTLSVersionGetParamsBrowserFamilyChrome, radar.HTTPLocationTLSVersionGetParamsBrowserFamilyEdge, radar.HTTPLocationTLSVersionGetParamsBrowserFamilyFirefox}),
			Continent:     cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:     cloudflare.F([]radar.HTTPLocationTLSVersionGetParamsDateRange{radar.HTTPLocationTLSVersionGetParamsDateRange1d, radar.HTTPLocationTLSVersionGetParamsDateRange2d, radar.HTTPLocationTLSVersionGetParamsDateRange7d}),
			DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:    cloudflare.F([]radar.HTTPLocationTLSVersionGetParamsDeviceType{radar.HTTPLocationTLSVersionGetParamsDeviceTypeDesktop, radar.HTTPLocationTLSVersionGetParamsDeviceTypeMobile, radar.HTTPLocationTLSVersionGetParamsDeviceTypeOther}),
			Format:        cloudflare.F(radar.HTTPLocationTLSVersionGetParamsFormatJson),
			HTTPProtocol:  cloudflare.F([]radar.HTTPLocationTLSVersionGetParamsHTTPProtocol{radar.HTTPLocationTLSVersionGetParamsHTTPProtocolHTTP, radar.HTTPLocationTLSVersionGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:   cloudflare.F([]radar.HTTPLocationTLSVersionGetParamsHTTPVersion{radar.HTTPLocationTLSVersionGetParamsHTTPVersionHttPv1, radar.HTTPLocationTLSVersionGetParamsHTTPVersionHttPv2, radar.HTTPLocationTLSVersionGetParamsHTTPVersionHttPv3}),
			IPVersion:     cloudflare.F([]radar.HTTPLocationTLSVersionGetParamsIPVersion{radar.HTTPLocationTLSVersionGetParamsIPVersionIPv4, radar.HTTPLocationTLSVersionGetParamsIPVersionIPv6}),
			Limit:         cloudflare.F(int64(5)),
			Location:      cloudflare.F([]string{"string", "string", "string"}),
			Name:          cloudflare.F([]string{"string", "string", "string"}),
			OS:            cloudflare.F([]radar.HTTPLocationTLSVersionGetParamsOS{radar.HTTPLocationTLSVersionGetParamsOSWindows, radar.HTTPLocationTLSVersionGetParamsOSMacosx, radar.HTTPLocationTLSVersionGetParamsOSIos}),
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
