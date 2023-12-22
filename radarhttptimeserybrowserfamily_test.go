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

func TestRadarHTTPTimeseryBrowserFamilyListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.HTTPs.Timeseries.BrowserFamilies.List(context.TODO(), cloudflare.RadarHTTPTimeseryBrowserFamilyListParams{
		AggInterval:  cloudflare.F(cloudflare.RadarHTTPTimeseryBrowserFamilyListParamsAggInterval15m),
		ASN:          cloudflare.F([]string{"15169", "15169", "15169"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPTimeseryBrowserFamilyListParamsDateRange{cloudflare.RadarHTTPTimeseryBrowserFamilyListParamsDateRange1d, cloudflare.RadarHTTPTimeseryBrowserFamilyListParamsDateRange7d, cloudflare.RadarHTTPTimeseryBrowserFamilyListParamsDateRange14d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTimeseryBrowserFamilyListParamsDeviceType{cloudflare.RadarHTTPTimeseryBrowserFamilyListParamsDeviceTypeDesktop, cloudflare.RadarHTTPTimeseryBrowserFamilyListParamsDeviceTypeMobile, cloudflare.RadarHTTPTimeseryBrowserFamilyListParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPTimeseryBrowserFamilyListParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTimeseryBrowserFamilyListParamsHTTPProtocol{cloudflare.RadarHTTPTimeseryBrowserFamilyListParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTimeseryBrowserFamilyListParamsHTTPProtocolHTTPs}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTimeseryBrowserFamilyListParamsHTTPVersion{cloudflare.RadarHTTPTimeseryBrowserFamilyListParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTimeseryBrowserFamilyListParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTimeseryBrowserFamilyListParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTimeseryBrowserFamilyListParamsIPVersion{cloudflare.RadarHTTPTimeseryBrowserFamilyListParamsIPVersionIPv4, cloudflare.RadarHTTPTimeseryBrowserFamilyListParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
		Name:         cloudflare.F([]string{"main_series", "main_series", "main_series"}),
		Os:           cloudflare.F([]cloudflare.RadarHTTPTimeseryBrowserFamilyListParamsO{cloudflare.RadarHTTPTimeseryBrowserFamilyListParamsOWindows, cloudflare.RadarHTTPTimeseryBrowserFamilyListParamsOMacosx, cloudflare.RadarHTTPTimeseryBrowserFamilyListParamsOAndroid}),
		TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPTimeseryBrowserFamilyListParamsTlsVersion{cloudflare.RadarHTTPTimeseryBrowserFamilyListParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTimeseryBrowserFamilyListParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTimeseryBrowserFamilyListParamsTlsVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
