// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/radar"
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
			ASN:          cloudflare.F([]string{"string"}),
			BotClass:     cloudflare.F([]radar.HTTPAseBrowserFamilyGetParamsBotClass{radar.HTTPAseBrowserFamilyGetParamsBotClassLikelyAutomated}),
			Continent:    cloudflare.F([]string{"string"}),
			DateEnd:      cloudflare.F([]time.Time{time.Now()}),
			DateRange:    cloudflare.F([]string{"7d"}),
			DateStart:    cloudflare.F([]time.Time{time.Now()}),
			DeviceType:   cloudflare.F([]radar.HTTPAseBrowserFamilyGetParamsDeviceType{radar.HTTPAseBrowserFamilyGetParamsDeviceTypeDesktop}),
			Format:       cloudflare.F(radar.HTTPAseBrowserFamilyGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]radar.HTTPAseBrowserFamilyGetParamsHTTPProtocol{radar.HTTPAseBrowserFamilyGetParamsHTTPProtocolHTTP}),
			HTTPVersion:  cloudflare.F([]radar.HTTPAseBrowserFamilyGetParamsHTTPVersion{radar.HTTPAseBrowserFamilyGetParamsHTTPVersionHttPv1}),
			IPVersion:    cloudflare.F([]radar.HTTPAseBrowserFamilyGetParamsIPVersion{radar.HTTPAseBrowserFamilyGetParamsIPVersionIPv4}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"string"}),
			Name:         cloudflare.F([]string{"string"}),
			OS:           cloudflare.F([]radar.HTTPAseBrowserFamilyGetParamsOS{radar.HTTPAseBrowserFamilyGetParamsOSWindows}),
			TLSVersion:   cloudflare.F([]radar.HTTPAseBrowserFamilyGetParamsTLSVersion{radar.HTTPAseBrowserFamilyGetParamsTLSVersionTlSv1_0}),
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
