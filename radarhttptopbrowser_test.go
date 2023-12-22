// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestRadarHTTPTopBrowserListWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Radars.HTTPs.Tops.Browsers.List(context.TODO(), cloudflare.RadarHTTPTopBrowserListParams{
		ASN:          cloudflare.F([]string{"15169", "15169", "15169"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPTopBrowserListParamsBotClass{cloudflare.RadarHTTPTopBrowserListParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTopBrowserListParamsBotClassLikelyHuman}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPTopBrowserListParamsDateRange{cloudflare.RadarHTTPTopBrowserListParamsDateRange1d, cloudflare.RadarHTTPTopBrowserListParamsDateRange7d, cloudflare.RadarHTTPTopBrowserListParamsDateRange14d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTopBrowserListParamsDeviceType{cloudflare.RadarHTTPTopBrowserListParamsDeviceTypeDesktop, cloudflare.RadarHTTPTopBrowserListParamsDeviceTypeMobile, cloudflare.RadarHTTPTopBrowserListParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPTopBrowserListParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTopBrowserListParamsHTTPProtocol{cloudflare.RadarHTTPTopBrowserListParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTopBrowserListParamsHTTPProtocolHTTPs}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTopBrowserListParamsHTTPVersion{cloudflare.RadarHTTPTopBrowserListParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTopBrowserListParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTopBrowserListParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTopBrowserListParamsIPVersion{cloudflare.RadarHTTPTopBrowserListParamsIPVersionIPv4, cloudflare.RadarHTTPTopBrowserListParamsIPVersionIPv6}),
		Limit:        cloudflare.F(int64(5)),
		Location:     cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
		Name:         cloudflare.F([]string{"main_series", "main_series", "main_series"}),
		Os:           cloudflare.F([]cloudflare.RadarHTTPTopBrowserListParamsO{cloudflare.RadarHTTPTopBrowserListParamsOWindows, cloudflare.RadarHTTPTopBrowserListParamsOMacosx, cloudflare.RadarHTTPTopBrowserListParamsOAndroid}),
		TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPTopBrowserListParamsTlsVersion{cloudflare.RadarHTTPTopBrowserListParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTopBrowserListParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTopBrowserListParamsTlsVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
