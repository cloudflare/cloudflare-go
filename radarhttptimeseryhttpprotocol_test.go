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

func TestRadarHTTPTimeseryHTTPProtocolListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.HTTPs.Timeseries.HTTPProtocols.List(context.TODO(), cloudflare.RadarHTTPTimeseryHTTPProtocolListParams{
		AggInterval: cloudflare.F(cloudflare.RadarHTTPTimeseryHTTPProtocolListParamsAggInterval15m),
		ASN:         cloudflare.F([]string{"15169", "15169", "15169"}),
		BotClass:    cloudflare.F([]cloudflare.RadarHTTPTimeseryHTTPProtocolListParamsBotClass{cloudflare.RadarHTTPTimeseryHTTPProtocolListParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTimeseryHTTPProtocolListParamsBotClassLikelyHuman}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarHTTPTimeseryHTTPProtocolListParamsDateRange{cloudflare.RadarHTTPTimeseryHTTPProtocolListParamsDateRange1d, cloudflare.RadarHTTPTimeseryHTTPProtocolListParamsDateRange7d, cloudflare.RadarHTTPTimeseryHTTPProtocolListParamsDateRange14d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:  cloudflare.F([]cloudflare.RadarHTTPTimeseryHTTPProtocolListParamsDeviceType{cloudflare.RadarHTTPTimeseryHTTPProtocolListParamsDeviceTypeDesktop, cloudflare.RadarHTTPTimeseryHTTPProtocolListParamsDeviceTypeMobile, cloudflare.RadarHTTPTimeseryHTTPProtocolListParamsDeviceTypeOther}),
		Format:      cloudflare.F(cloudflare.RadarHTTPTimeseryHTTPProtocolListParamsFormatJson),
		HTTPVersion: cloudflare.F([]cloudflare.RadarHTTPTimeseryHTTPProtocolListParamsHTTPVersion{cloudflare.RadarHTTPTimeseryHTTPProtocolListParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTimeseryHTTPProtocolListParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTimeseryHTTPProtocolListParamsHTTPVersionHttPv3}),
		IPVersion:   cloudflare.F([]cloudflare.RadarHTTPTimeseryHTTPProtocolListParamsIPVersion{cloudflare.RadarHTTPTimeseryHTTPProtocolListParamsIPVersionIPv4, cloudflare.RadarHTTPTimeseryHTTPProtocolListParamsIPVersionIPv6}),
		Location:    cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
		Name:        cloudflare.F([]string{"main_series", "main_series", "main_series"}),
		Os:          cloudflare.F([]cloudflare.RadarHTTPTimeseryHTTPProtocolListParamsO{cloudflare.RadarHTTPTimeseryHTTPProtocolListParamsOWindows, cloudflare.RadarHTTPTimeseryHTTPProtocolListParamsOMacosx, cloudflare.RadarHTTPTimeseryHTTPProtocolListParamsOAndroid}),
		TlsVersion:  cloudflare.F([]cloudflare.RadarHTTPTimeseryHTTPProtocolListParamsTlsVersion{cloudflare.RadarHTTPTimeseryHTTPProtocolListParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTimeseryHTTPProtocolListParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTimeseryHTTPProtocolListParamsTlsVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
