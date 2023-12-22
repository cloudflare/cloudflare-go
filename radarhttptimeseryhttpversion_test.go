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

func TestRadarHTTPTimeseryHTTPVersionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.HTTPs.Timeseries.HTTPVersions.List(context.TODO(), cloudflare.RadarHTTPTimeseryHTTPVersionListParams{
		AggInterval:  cloudflare.F(cloudflare.RadarHTTPTimeseryHTTPVersionListParamsAggInterval15m),
		ASN:          cloudflare.F([]string{"15169", "15169", "15169"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPTimeseryHTTPVersionListParamsBotClass{cloudflare.RadarHTTPTimeseryHTTPVersionListParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTimeseryHTTPVersionListParamsBotClassLikelyHuman}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPTimeseryHTTPVersionListParamsDateRange{cloudflare.RadarHTTPTimeseryHTTPVersionListParamsDateRange1d, cloudflare.RadarHTTPTimeseryHTTPVersionListParamsDateRange7d, cloudflare.RadarHTTPTimeseryHTTPVersionListParamsDateRange14d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTimeseryHTTPVersionListParamsDeviceType{cloudflare.RadarHTTPTimeseryHTTPVersionListParamsDeviceTypeDesktop, cloudflare.RadarHTTPTimeseryHTTPVersionListParamsDeviceTypeMobile, cloudflare.RadarHTTPTimeseryHTTPVersionListParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPTimeseryHTTPVersionListParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTimeseryHTTPVersionListParamsHTTPProtocol{cloudflare.RadarHTTPTimeseryHTTPVersionListParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTimeseryHTTPVersionListParamsHTTPProtocolHTTPs}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTimeseryHTTPVersionListParamsIPVersion{cloudflare.RadarHTTPTimeseryHTTPVersionListParamsIPVersionIPv4, cloudflare.RadarHTTPTimeseryHTTPVersionListParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
		Name:         cloudflare.F([]string{"main_series", "main_series", "main_series"}),
		Os:           cloudflare.F([]cloudflare.RadarHTTPTimeseryHTTPVersionListParamsO{cloudflare.RadarHTTPTimeseryHTTPVersionListParamsOWindows, cloudflare.RadarHTTPTimeseryHTTPVersionListParamsOMacosx, cloudflare.RadarHTTPTimeseryHTTPVersionListParamsOAndroid}),
		TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPTimeseryHTTPVersionListParamsTlsVersion{cloudflare.RadarHTTPTimeseryHTTPVersionListParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTimeseryHTTPVersionListParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTimeseryHTTPVersionListParamsTlsVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
