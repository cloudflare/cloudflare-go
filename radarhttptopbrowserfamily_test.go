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

func TestRadarHTTPTopBrowserFamilyListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.HTTPs.Tops.BrowserFamilies.List(context.TODO(), cloudflare.RadarHTTPTopBrowserFamilyListParams{
		ASN:          cloudflare.F([]string{"15169", "15169", "15169"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPTopBrowserFamilyListParamsBotClass{cloudflare.RadarHTTPTopBrowserFamilyListParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTopBrowserFamilyListParamsBotClassLikelyHuman}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPTopBrowserFamilyListParamsDateRange{cloudflare.RadarHTTPTopBrowserFamilyListParamsDateRange1d, cloudflare.RadarHTTPTopBrowserFamilyListParamsDateRange7d, cloudflare.RadarHTTPTopBrowserFamilyListParamsDateRange14d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTopBrowserFamilyListParamsDeviceType{cloudflare.RadarHTTPTopBrowserFamilyListParamsDeviceTypeDesktop, cloudflare.RadarHTTPTopBrowserFamilyListParamsDeviceTypeMobile, cloudflare.RadarHTTPTopBrowserFamilyListParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPTopBrowserFamilyListParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTopBrowserFamilyListParamsHTTPProtocol{cloudflare.RadarHTTPTopBrowserFamilyListParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTopBrowserFamilyListParamsHTTPProtocolHTTPs}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTopBrowserFamilyListParamsHTTPVersion{cloudflare.RadarHTTPTopBrowserFamilyListParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTopBrowserFamilyListParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTopBrowserFamilyListParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTopBrowserFamilyListParamsIPVersion{cloudflare.RadarHTTPTopBrowserFamilyListParamsIPVersionIPv4, cloudflare.RadarHTTPTopBrowserFamilyListParamsIPVersionIPv6}),
		Limit:        cloudflare.F(int64(5)),
		Location:     cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
		Name:         cloudflare.F([]string{"main_series", "main_series", "main_series"}),
		Os:           cloudflare.F([]cloudflare.RadarHTTPTopBrowserFamilyListParamsO{cloudflare.RadarHTTPTopBrowserFamilyListParamsOWindows, cloudflare.RadarHTTPTopBrowserFamilyListParamsOMacosx, cloudflare.RadarHTTPTopBrowserFamilyListParamsOAndroid}),
		TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPTopBrowserFamilyListParamsTlsVersion{cloudflare.RadarHTTPTopBrowserFamilyListParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTopBrowserFamilyListParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTopBrowserFamilyListParamsTlsVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
