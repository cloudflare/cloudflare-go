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

func TestRadarHTTPTimeseryDeviceTypeListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.HTTPs.Timeseries.DeviceTypes.List(context.TODO(), cloudflare.RadarHTTPTimeseryDeviceTypeListParams{
		AggInterval:  cloudflare.F(cloudflare.RadarHTTPTimeseryDeviceTypeListParamsAggInterval15m),
		ASN:          cloudflare.F([]string{"15169", "15169", "15169"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPTimeseryDeviceTypeListParamsBotClass{cloudflare.RadarHTTPTimeseryDeviceTypeListParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTimeseryDeviceTypeListParamsBotClassLikelyHuman}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPTimeseryDeviceTypeListParamsDateRange{cloudflare.RadarHTTPTimeseryDeviceTypeListParamsDateRange1d, cloudflare.RadarHTTPTimeseryDeviceTypeListParamsDateRange7d, cloudflare.RadarHTTPTimeseryDeviceTypeListParamsDateRange14d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:       cloudflare.F(cloudflare.RadarHTTPTimeseryDeviceTypeListParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTimeseryDeviceTypeListParamsHTTPProtocol{cloudflare.RadarHTTPTimeseryDeviceTypeListParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTimeseryDeviceTypeListParamsHTTPProtocolHTTPs}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTimeseryDeviceTypeListParamsHTTPVersion{cloudflare.RadarHTTPTimeseryDeviceTypeListParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTimeseryDeviceTypeListParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTimeseryDeviceTypeListParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTimeseryDeviceTypeListParamsIPVersion{cloudflare.RadarHTTPTimeseryDeviceTypeListParamsIPVersionIPv4, cloudflare.RadarHTTPTimeseryDeviceTypeListParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
		Name:         cloudflare.F([]string{"main_series", "main_series", "main_series"}),
		Os:           cloudflare.F([]cloudflare.RadarHTTPTimeseryDeviceTypeListParamsO{cloudflare.RadarHTTPTimeseryDeviceTypeListParamsOWindows, cloudflare.RadarHTTPTimeseryDeviceTypeListParamsOMacosx, cloudflare.RadarHTTPTimeseryDeviceTypeListParamsOAndroid}),
		TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPTimeseryDeviceTypeListParamsTlsVersion{cloudflare.RadarHTTPTimeseryDeviceTypeListParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTimeseryDeviceTypeListParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTimeseryDeviceTypeListParamsTlsVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
