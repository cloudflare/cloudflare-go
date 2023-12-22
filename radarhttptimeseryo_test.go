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

func TestRadarHTTPTimeseryOListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.HTTPs.Timeseries.Os.List(context.TODO(), cloudflare.RadarHTTPTimeseryOListParams{
		AggInterval:   cloudflare.F(cloudflare.RadarHTTPTimeseryOListParamsAggInterval15m),
		ASN:           cloudflare.F([]string{"15169", "15169", "15169"}),
		BotClass:      cloudflare.F([]cloudflare.RadarHTTPTimeseryOListParamsBotClass{cloudflare.RadarHTTPTimeseryOListParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTimeseryOListParamsBotClassLikelyHuman}),
		DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     cloudflare.F([]cloudflare.RadarHTTPTimeseryOListParamsDateRange{cloudflare.RadarHTTPTimeseryOListParamsDateRange1d, cloudflare.RadarHTTPTimeseryOListParamsDateRange7d, cloudflare.RadarHTTPTimeseryOListParamsDateRange14d}),
		DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:    cloudflare.F([]cloudflare.RadarHTTPTimeseryOListParamsDeviceType{cloudflare.RadarHTTPTimeseryOListParamsDeviceTypeDesktop, cloudflare.RadarHTTPTimeseryOListParamsDeviceTypeMobile, cloudflare.RadarHTTPTimeseryOListParamsDeviceTypeOther}),
		Format:        cloudflare.F(cloudflare.RadarHTTPTimeseryOListParamsFormatJson),
		HTTPProtocol:  cloudflare.F([]cloudflare.RadarHTTPTimeseryOListParamsHTTPProtocol{cloudflare.RadarHTTPTimeseryOListParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTimeseryOListParamsHTTPProtocolHTTPs}),
		HTTPVersion:   cloudflare.F([]cloudflare.RadarHTTPTimeseryOListParamsHTTPVersion{cloudflare.RadarHTTPTimeseryOListParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTimeseryOListParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTimeseryOListParamsHTTPVersionHttPv3}),
		IPVersion:     cloudflare.F([]cloudflare.RadarHTTPTimeseryOListParamsIPVersion{cloudflare.RadarHTTPTimeseryOListParamsIPVersionIPv4, cloudflare.RadarHTTPTimeseryOListParamsIPVersionIPv6}),
		LimitPerGroup: cloudflare.F(int64(2)),
		Location:      cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
		Name:          cloudflare.F([]string{"main_series", "main_series", "main_series"}),
		TlsVersion:    cloudflare.F([]cloudflare.RadarHTTPTimeseryOListParamsTlsVersion{cloudflare.RadarHTTPTimeseryOListParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTimeseryOListParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTimeseryOListParamsTlsVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
