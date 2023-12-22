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

func TestRadarHTTPTimeseryTlsVersionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.HTTPs.Timeseries.TlsVersions.List(context.TODO(), cloudflare.RadarHTTPTimeseryTlsVersionListParams{
		AggInterval:  cloudflare.F(cloudflare.RadarHTTPTimeseryTlsVersionListParamsAggInterval15m),
		ASN:          cloudflare.F([]string{"15169", "15169", "15169"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPTimeseryTlsVersionListParamsBotClass{cloudflare.RadarHTTPTimeseryTlsVersionListParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTimeseryTlsVersionListParamsBotClassLikelyHuman}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPTimeseryTlsVersionListParamsDateRange{cloudflare.RadarHTTPTimeseryTlsVersionListParamsDateRange1d, cloudflare.RadarHTTPTimeseryTlsVersionListParamsDateRange7d, cloudflare.RadarHTTPTimeseryTlsVersionListParamsDateRange14d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTimeseryTlsVersionListParamsDeviceType{cloudflare.RadarHTTPTimeseryTlsVersionListParamsDeviceTypeDesktop, cloudflare.RadarHTTPTimeseryTlsVersionListParamsDeviceTypeMobile, cloudflare.RadarHTTPTimeseryTlsVersionListParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPTimeseryTlsVersionListParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTimeseryTlsVersionListParamsHTTPProtocol{cloudflare.RadarHTTPTimeseryTlsVersionListParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTimeseryTlsVersionListParamsHTTPProtocolHTTPs}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTimeseryTlsVersionListParamsHTTPVersion{cloudflare.RadarHTTPTimeseryTlsVersionListParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTimeseryTlsVersionListParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTimeseryTlsVersionListParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTimeseryTlsVersionListParamsIPVersion{cloudflare.RadarHTTPTimeseryTlsVersionListParamsIPVersionIPv4, cloudflare.RadarHTTPTimeseryTlsVersionListParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
		Name:         cloudflare.F([]string{"main_series", "main_series", "main_series"}),
		Os:           cloudflare.F([]cloudflare.RadarHTTPTimeseryTlsVersionListParamsO{cloudflare.RadarHTTPTimeseryTlsVersionListParamsOWindows, cloudflare.RadarHTTPTimeseryTlsVersionListParamsOMacosx, cloudflare.RadarHTTPTimeseryTlsVersionListParamsOAndroid}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
