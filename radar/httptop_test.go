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

func TestHTTPTopBrowserWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Top.Browser(context.TODO(), radar.HTTPTopBrowserParams{
		ASN:           cloudflare.F([]string{"string"}),
		BotClass:      cloudflare.F([]radar.HTTPTopBrowserParamsBotClass{radar.HTTPTopBrowserParamsBotClassLikelyAutomated}),
		BrowserFamily: cloudflare.F([]radar.HTTPTopBrowserParamsBrowserFamily{radar.HTTPTopBrowserParamsBrowserFamilyChrome}),
		Continent:     cloudflare.F([]string{"string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now()}),
		DateRange:     cloudflare.F([]string{"7d"}),
		DateStart:     cloudflare.F([]time.Time{time.Now()}),
		DeviceType:    cloudflare.F([]radar.HTTPTopBrowserParamsDeviceType{radar.HTTPTopBrowserParamsDeviceTypeDesktop}),
		Format:        cloudflare.F(radar.HTTPTopBrowserParamsFormatJson),
		HTTPProtocol:  cloudflare.F([]radar.HTTPTopBrowserParamsHTTPProtocol{radar.HTTPTopBrowserParamsHTTPProtocolHTTP}),
		HTTPVersion:   cloudflare.F([]radar.HTTPTopBrowserParamsHTTPVersion{radar.HTTPTopBrowserParamsHTTPVersionHttPv1}),
		IPVersion:     cloudflare.F([]radar.HTTPTopBrowserParamsIPVersion{radar.HTTPTopBrowserParamsIPVersionIPv4}),
		Limit:         cloudflare.F(int64(5)),
		Location:      cloudflare.F([]string{"string"}),
		Name:          cloudflare.F([]string{"string"}),
		OS:            cloudflare.F([]radar.HTTPTopBrowserParamsOS{radar.HTTPTopBrowserParamsOSWindows}),
		TLSVersion:    cloudflare.F([]radar.HTTPTopBrowserParamsTLSVersion{radar.HTTPTopBrowserParamsTLSVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPTopBrowserFamilyWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Top.BrowserFamily(context.TODO(), radar.HTTPTopBrowserFamilyParams{
		ASN:           cloudflare.F([]string{"string"}),
		BotClass:      cloudflare.F([]radar.HTTPTopBrowserFamilyParamsBotClass{radar.HTTPTopBrowserFamilyParamsBotClassLikelyAutomated}),
		BrowserFamily: cloudflare.F([]radar.HTTPTopBrowserFamilyParamsBrowserFamily{radar.HTTPTopBrowserFamilyParamsBrowserFamilyChrome}),
		Continent:     cloudflare.F([]string{"string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now()}),
		DateRange:     cloudflare.F([]string{"7d"}),
		DateStart:     cloudflare.F([]time.Time{time.Now()}),
		DeviceType:    cloudflare.F([]radar.HTTPTopBrowserFamilyParamsDeviceType{radar.HTTPTopBrowserFamilyParamsDeviceTypeDesktop}),
		Format:        cloudflare.F(radar.HTTPTopBrowserFamilyParamsFormatJson),
		HTTPProtocol:  cloudflare.F([]radar.HTTPTopBrowserFamilyParamsHTTPProtocol{radar.HTTPTopBrowserFamilyParamsHTTPProtocolHTTP}),
		HTTPVersion:   cloudflare.F([]radar.HTTPTopBrowserFamilyParamsHTTPVersion{radar.HTTPTopBrowserFamilyParamsHTTPVersionHttPv1}),
		IPVersion:     cloudflare.F([]radar.HTTPTopBrowserFamilyParamsIPVersion{radar.HTTPTopBrowserFamilyParamsIPVersionIPv4}),
		Limit:         cloudflare.F(int64(5)),
		Location:      cloudflare.F([]string{"string"}),
		Name:          cloudflare.F([]string{"string"}),
		OS:            cloudflare.F([]radar.HTTPTopBrowserFamilyParamsOS{radar.HTTPTopBrowserFamilyParamsOSWindows}),
		TLSVersion:    cloudflare.F([]radar.HTTPTopBrowserFamilyParamsTLSVersion{radar.HTTPTopBrowserFamilyParamsTLSVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
