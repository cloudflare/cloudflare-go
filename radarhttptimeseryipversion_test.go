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

func TestRadarHTTPTimeseryIPVersionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.HTTPs.Timeseries.IPVersions.List(context.TODO(), cloudflare.RadarHTTPTimeseryIPVersionListParams{
		AggInterval:  cloudflare.F(cloudflare.RadarHTTPTimeseryIPVersionListParamsAggInterval15m),
		ASN:          cloudflare.F([]string{"15169", "15169", "15169"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPTimeseryIPVersionListParamsBotClass{cloudflare.RadarHTTPTimeseryIPVersionListParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTimeseryIPVersionListParamsBotClassLikelyHuman}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPTimeseryIPVersionListParamsDateRange{cloudflare.RadarHTTPTimeseryIPVersionListParamsDateRange1d, cloudflare.RadarHTTPTimeseryIPVersionListParamsDateRange7d, cloudflare.RadarHTTPTimeseryIPVersionListParamsDateRange14d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTimeseryIPVersionListParamsDeviceType{cloudflare.RadarHTTPTimeseryIPVersionListParamsDeviceTypeDesktop, cloudflare.RadarHTTPTimeseryIPVersionListParamsDeviceTypeMobile, cloudflare.RadarHTTPTimeseryIPVersionListParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPTimeseryIPVersionListParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTimeseryIPVersionListParamsHTTPProtocol{cloudflare.RadarHTTPTimeseryIPVersionListParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTimeseryIPVersionListParamsHTTPProtocolHTTPs}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTimeseryIPVersionListParamsHTTPVersion{cloudflare.RadarHTTPTimeseryIPVersionListParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTimeseryIPVersionListParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTimeseryIPVersionListParamsHTTPVersionHttPv3}),
		Location:     cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
		Name:         cloudflare.F([]string{"main_series", "main_series", "main_series"}),
		Os:           cloudflare.F([]cloudflare.RadarHTTPTimeseryIPVersionListParamsO{cloudflare.RadarHTTPTimeseryIPVersionListParamsOWindows, cloudflare.RadarHTTPTimeseryIPVersionListParamsOMacosx, cloudflare.RadarHTTPTimeseryIPVersionListParamsOAndroid}),
		TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPTimeseryIPVersionListParamsTlsVersion{cloudflare.RadarHTTPTimeseryIPVersionListParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTimeseryIPVersionListParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTimeseryIPVersionListParamsTlsVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
