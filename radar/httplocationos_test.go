// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/radar"
)

func TestHTTPLocationOSGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Locations.OS.Get(
		context.TODO(),
		radar.HTTPLocationOSGetParamsOSWindows,
		radar.HTTPLocationOSGetParams{
			ASN:           cloudflare.F([]string{"string"}),
			BotClass:      cloudflare.F([]radar.HTTPLocationOSGetParamsBotClass{radar.HTTPLocationOSGetParamsBotClassLikelyAutomated}),
			BrowserFamily: cloudflare.F([]radar.HTTPLocationOSGetParamsBrowserFamily{radar.HTTPLocationOSGetParamsBrowserFamilyChrome}),
			Continent:     cloudflare.F([]string{"string"}),
			DateEnd:       cloudflare.F([]time.Time{time.Now()}),
			DateRange:     cloudflare.F([]string{"7d"}),
			DateStart:     cloudflare.F([]time.Time{time.Now()}),
			DeviceType:    cloudflare.F([]radar.HTTPLocationOSGetParamsDeviceType{radar.HTTPLocationOSGetParamsDeviceTypeDesktop}),
			Format:        cloudflare.F(radar.HTTPLocationOSGetParamsFormatJson),
			HTTPProtocol:  cloudflare.F([]radar.HTTPLocationOSGetParamsHTTPProtocol{radar.HTTPLocationOSGetParamsHTTPProtocolHTTP}),
			HTTPVersion:   cloudflare.F([]radar.HTTPLocationOSGetParamsHTTPVersion{radar.HTTPLocationOSGetParamsHTTPVersionHttPv1}),
			IPVersion:     cloudflare.F([]radar.HTTPLocationOSGetParamsIPVersion{radar.HTTPLocationOSGetParamsIPVersionIPv4}),
			Limit:         cloudflare.F(int64(5)),
			Location:      cloudflare.F([]string{"string"}),
			Name:          cloudflare.F([]string{"string"}),
			TLSVersion:    cloudflare.F([]radar.HTTPLocationOSGetParamsTLSVersion{radar.HTTPLocationOSGetParamsTLSVersionTlSv1_0}),
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
