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

func TestHTTPLocationIPVersionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Locations.IPVersion.Get(
		context.TODO(),
		radar.HTTPLocationIPVersionGetParamsIPVersionIPv4,
		radar.HTTPLocationIPVersionGetParams{
			ASN:           cloudflare.F([]string{"string", "string", "string"}),
			BotClass:      cloudflare.F([]radar.HTTPLocationIPVersionGetParamsBotClass{radar.HTTPLocationIPVersionGetParamsBotClassLikelyAutomated, radar.HTTPLocationIPVersionGetParamsBotClassLikelyHuman}),
			BrowserFamily: cloudflare.F([]radar.HTTPLocationIPVersionGetParamsBrowserFamily{radar.HTTPLocationIPVersionGetParamsBrowserFamilyChrome, radar.HTTPLocationIPVersionGetParamsBrowserFamilyEdge, radar.HTTPLocationIPVersionGetParamsBrowserFamilyFirefox}),
			Continent:     cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:     cloudflare.F([]string{"7d", "7d", "7d"}),
			DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:    cloudflare.F([]radar.HTTPLocationIPVersionGetParamsDeviceType{radar.HTTPLocationIPVersionGetParamsDeviceTypeDesktop, radar.HTTPLocationIPVersionGetParamsDeviceTypeMobile, radar.HTTPLocationIPVersionGetParamsDeviceTypeOther}),
			Format:        cloudflare.F(radar.HTTPLocationIPVersionGetParamsFormatJson),
			HTTPProtocol:  cloudflare.F([]radar.HTTPLocationIPVersionGetParamsHTTPProtocol{radar.HTTPLocationIPVersionGetParamsHTTPProtocolHTTP, radar.HTTPLocationIPVersionGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:   cloudflare.F([]radar.HTTPLocationIPVersionGetParamsHTTPVersion{radar.HTTPLocationIPVersionGetParamsHTTPVersionHttPv1, radar.HTTPLocationIPVersionGetParamsHTTPVersionHttPv2, radar.HTTPLocationIPVersionGetParamsHTTPVersionHttPv3}),
			Limit:         cloudflare.F(int64(5)),
			Location:      cloudflare.F([]string{"string", "string", "string"}),
			Name:          cloudflare.F([]string{"string", "string", "string"}),
			OS:            cloudflare.F([]radar.HTTPLocationIPVersionGetParamsOS{radar.HTTPLocationIPVersionGetParamsOSWindows, radar.HTTPLocationIPVersionGetParamsOSMacosx, radar.HTTPLocationIPVersionGetParamsOSIos}),
			TLSVersion:    cloudflare.F([]radar.HTTPLocationIPVersionGetParamsTLSVersion{radar.HTTPLocationIPVersionGetParamsTLSVersionTlSv1_0, radar.HTTPLocationIPVersionGetParamsTLSVersionTlSv1_1, radar.HTTPLocationIPVersionGetParamsTLSVersionTlSv1_2}),
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
