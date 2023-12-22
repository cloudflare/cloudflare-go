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

func TestRadarHTTPTopAseTlsVersionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.HTTPs.Tops.Ases.TlsVersions.Get(
		context.TODO(),
		cloudflare.RadarHTTPTopAseTlsVersionGetParamsTlsVersionTlSv1_0,
		cloudflare.RadarHTTPTopAseTlsVersionGetParams{
			ASN:          cloudflare.F([]string{"15169", "15169", "15169"}),
			BotClass:     cloudflare.F([]cloudflare.RadarHTTPTopAseTlsVersionGetParamsBotClass{cloudflare.RadarHTTPTopAseTlsVersionGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTopAseTlsVersionGetParamsBotClassLikelyHuman}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPTopAseTlsVersionGetParamsDateRange{cloudflare.RadarHTTPTopAseTlsVersionGetParamsDateRange1d, cloudflare.RadarHTTPTopAseTlsVersionGetParamsDateRange7d, cloudflare.RadarHTTPTopAseTlsVersionGetParamsDateRange14d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTopAseTlsVersionGetParamsDeviceType{cloudflare.RadarHTTPTopAseTlsVersionGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPTopAseTlsVersionGetParamsDeviceTypeMobile, cloudflare.RadarHTTPTopAseTlsVersionGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(cloudflare.RadarHTTPTopAseTlsVersionGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTopAseTlsVersionGetParamsHTTPProtocol{cloudflare.RadarHTTPTopAseTlsVersionGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTopAseTlsVersionGetParamsHTTPProtocolHTTPs}),
			HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTopAseTlsVersionGetParamsHTTPVersion{cloudflare.RadarHTTPTopAseTlsVersionGetParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTopAseTlsVersionGetParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTopAseTlsVersionGetParamsHTTPVersionHttPv3}),
			IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTopAseTlsVersionGetParamsIPVersion{cloudflare.RadarHTTPTopAseTlsVersionGetParamsIPVersionIPv4, cloudflare.RadarHTTPTopAseTlsVersionGetParamsIPVersionIPv6}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
			Name:         cloudflare.F([]string{"main_series", "main_series", "main_series"}),
			Os:           cloudflare.F([]cloudflare.RadarHTTPTopAseTlsVersionGetParamsO{cloudflare.RadarHTTPTopAseTlsVersionGetParamsOWindows, cloudflare.RadarHTTPTopAseTlsVersionGetParamsOMacosx, cloudflare.RadarHTTPTopAseTlsVersionGetParamsOAndroid}),
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
