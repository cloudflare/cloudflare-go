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

func TestRadarHTTPTopLocationOGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.HTTPs.Tops.Locations.Os.Get(
		context.TODO(),
		cloudflare.RadarHTTPTopLocationOGetParamsOsWindows,
		cloudflare.RadarHTTPTopLocationOGetParams{
			ASN:          cloudflare.F([]string{"15169", "15169", "15169"}),
			BotClass:     cloudflare.F([]cloudflare.RadarHTTPTopLocationOGetParamsBotClass{cloudflare.RadarHTTPTopLocationOGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTopLocationOGetParamsBotClassLikelyHuman}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPTopLocationOGetParamsDateRange{cloudflare.RadarHTTPTopLocationOGetParamsDateRange1d, cloudflare.RadarHTTPTopLocationOGetParamsDateRange7d, cloudflare.RadarHTTPTopLocationOGetParamsDateRange14d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTopLocationOGetParamsDeviceType{cloudflare.RadarHTTPTopLocationOGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPTopLocationOGetParamsDeviceTypeMobile, cloudflare.RadarHTTPTopLocationOGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(cloudflare.RadarHTTPTopLocationOGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTopLocationOGetParamsHTTPProtocol{cloudflare.RadarHTTPTopLocationOGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTopLocationOGetParamsHTTPProtocolHTTPs}),
			HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTopLocationOGetParamsHTTPVersion{cloudflare.RadarHTTPTopLocationOGetParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTopLocationOGetParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTopLocationOGetParamsHTTPVersionHttPv3}),
			IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTopLocationOGetParamsIPVersion{cloudflare.RadarHTTPTopLocationOGetParamsIPVersionIPv4, cloudflare.RadarHTTPTopLocationOGetParamsIPVersionIPv6}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
			Name:         cloudflare.F([]string{"main_series", "main_series", "main_series"}),
			TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPTopLocationOGetParamsTlsVersion{cloudflare.RadarHTTPTopLocationOGetParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTopLocationOGetParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTopLocationOGetParamsTlsVersionTlSv1_2}),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
