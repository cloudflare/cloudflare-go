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

func TestRadarHTTPTimeseryBrowserListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.HTTPs.Timeseries.Browsers.List(context.TODO(), cloudflare.RadarHTTPTimeseryBrowserListParams{
		AggInterval:   cloudflare.F(cloudflare.RadarHTTPTimeseryBrowserListParamsAggInterval15m),
		ASN:           cloudflare.F([]string{"15169", "15169", "15169"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     cloudflare.F([]cloudflare.RadarHTTPTimeseryBrowserListParamsDateRange{cloudflare.RadarHTTPTimeseryBrowserListParamsDateRange1d, cloudflare.RadarHTTPTimeseryBrowserListParamsDateRange7d, cloudflare.RadarHTTPTimeseryBrowserListParamsDateRange14d}),
		DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:    cloudflare.F([]cloudflare.RadarHTTPTimeseryBrowserListParamsDeviceType{cloudflare.RadarHTTPTimeseryBrowserListParamsDeviceTypeDesktop, cloudflare.RadarHTTPTimeseryBrowserListParamsDeviceTypeMobile, cloudflare.RadarHTTPTimeseryBrowserListParamsDeviceTypeOther}),
		Format:        cloudflare.F(cloudflare.RadarHTTPTimeseryBrowserListParamsFormatJson),
		HTTPProtocol:  cloudflare.F([]cloudflare.RadarHTTPTimeseryBrowserListParamsHTTPProtocol{cloudflare.RadarHTTPTimeseryBrowserListParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTimeseryBrowserListParamsHTTPProtocolHTTPs}),
		HTTPVersion:   cloudflare.F([]cloudflare.RadarHTTPTimeseryBrowserListParamsHTTPVersion{cloudflare.RadarHTTPTimeseryBrowserListParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTimeseryBrowserListParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTimeseryBrowserListParamsHTTPVersionHttPv3}),
		IPVersion:     cloudflare.F([]cloudflare.RadarHTTPTimeseryBrowserListParamsIPVersion{cloudflare.RadarHTTPTimeseryBrowserListParamsIPVersionIPv4, cloudflare.RadarHTTPTimeseryBrowserListParamsIPVersionIPv6}),
		LimitPerGroup: cloudflare.F(int64(2)),
		Location:      cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
		Name:          cloudflare.F([]string{"main_series", "main_series", "main_series"}),
		Os:            cloudflare.F([]cloudflare.RadarHTTPTimeseryBrowserListParamsO{cloudflare.RadarHTTPTimeseryBrowserListParamsOWindows, cloudflare.RadarHTTPTimeseryBrowserListParamsOMacosx, cloudflare.RadarHTTPTimeseryBrowserListParamsOAndroid}),
		TlsVersion:    cloudflare.F([]cloudflare.RadarHTTPTimeseryBrowserListParamsTlsVersion{cloudflare.RadarHTTPTimeseryBrowserListParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTimeseryBrowserListParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTimeseryBrowserListParamsTlsVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
