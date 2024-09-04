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

func TestHTTPAseHTTPMethodGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Ases.HTTPMethod.Get(
		context.TODO(),
		radar.HTTPAseHTTPMethodGetParamsHTTPVersionHttPv1,
		radar.HTTPAseHTTPMethodGetParams{
			ASN:           cloudflare.F([]string{"string", "string", "string"}),
			BotClass:      cloudflare.F([]radar.HTTPAseHTTPMethodGetParamsBotClass{radar.HTTPAseHTTPMethodGetParamsBotClassLikelyAutomated, radar.HTTPAseHTTPMethodGetParamsBotClassLikelyHuman}),
			BrowserFamily: cloudflare.F([]radar.HTTPAseHTTPMethodGetParamsBrowserFamily{radar.HTTPAseHTTPMethodGetParamsBrowserFamilyChrome, radar.HTTPAseHTTPMethodGetParamsBrowserFamilyEdge, radar.HTTPAseHTTPMethodGetParamsBrowserFamilyFirefox}),
			Continent:     cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:     cloudflare.F([]string{"7d", "7d", "7d"}),
			DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:    cloudflare.F([]radar.HTTPAseHTTPMethodGetParamsDeviceType{radar.HTTPAseHTTPMethodGetParamsDeviceTypeDesktop, radar.HTTPAseHTTPMethodGetParamsDeviceTypeMobile, radar.HTTPAseHTTPMethodGetParamsDeviceTypeOther}),
			Format:        cloudflare.F(radar.HTTPAseHTTPMethodGetParamsFormatJson),
			HTTPProtocol:  cloudflare.F([]radar.HTTPAseHTTPMethodGetParamsHTTPProtocol{radar.HTTPAseHTTPMethodGetParamsHTTPProtocolHTTP, radar.HTTPAseHTTPMethodGetParamsHTTPProtocolHTTPS}),
			IPVersion:     cloudflare.F([]radar.HTTPAseHTTPMethodGetParamsIPVersion{radar.HTTPAseHTTPMethodGetParamsIPVersionIPv4, radar.HTTPAseHTTPMethodGetParamsIPVersionIPv6}),
			Limit:         cloudflare.F(int64(5)),
			Location:      cloudflare.F([]string{"string", "string", "string"}),
			Name:          cloudflare.F([]string{"string", "string", "string"}),
			OS:            cloudflare.F([]radar.HTTPAseHTTPMethodGetParamsOS{radar.HTTPAseHTTPMethodGetParamsOSWindows, radar.HTTPAseHTTPMethodGetParamsOSMacosx, radar.HTTPAseHTTPMethodGetParamsOSIos}),
			TLSVersion:    cloudflare.F([]radar.HTTPAseHTTPMethodGetParamsTLSVersion{radar.HTTPAseHTTPMethodGetParamsTLSVersionTlSv1_0, radar.HTTPAseHTTPMethodGetParamsTLSVersionTlSv1_1, radar.HTTPAseHTTPMethodGetParamsTLSVersionTlSv1_2}),
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
