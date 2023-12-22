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

func TestRadarHTTPTopAseOGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.HTTPs.Tops.Ases.Os.Get(
		context.TODO(),
		cloudflare.RadarHTTPTopAseOGetParamsOsWindows,
		cloudflare.RadarHTTPTopAseOGetParams{
			ASN:          cloudflare.F([]string{"15169", "15169", "15169"}),
			BotClass:     cloudflare.F([]cloudflare.RadarHTTPTopAseOGetParamsBotClass{cloudflare.RadarHTTPTopAseOGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTopAseOGetParamsBotClassLikelyHuman}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPTopAseOGetParamsDateRange{cloudflare.RadarHTTPTopAseOGetParamsDateRange1d, cloudflare.RadarHTTPTopAseOGetParamsDateRange7d, cloudflare.RadarHTTPTopAseOGetParamsDateRange14d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTopAseOGetParamsDeviceType{cloudflare.RadarHTTPTopAseOGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPTopAseOGetParamsDeviceTypeMobile, cloudflare.RadarHTTPTopAseOGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(cloudflare.RadarHTTPTopAseOGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTopAseOGetParamsHTTPProtocol{cloudflare.RadarHTTPTopAseOGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTopAseOGetParamsHTTPProtocolHTTPs}),
			HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTopAseOGetParamsHTTPVersion{cloudflare.RadarHTTPTopAseOGetParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTopAseOGetParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTopAseOGetParamsHTTPVersionHttPv3}),
			IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTopAseOGetParamsIPVersion{cloudflare.RadarHTTPTopAseOGetParamsIPVersionIPv4, cloudflare.RadarHTTPTopAseOGetParamsIPVersionIPv6}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
			Name:         cloudflare.F([]string{"main_series", "main_series", "main_series"}),
			TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPTopAseOGetParamsTlsVersion{cloudflare.RadarHTTPTopAseOGetParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTopAseOGetParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTopAseOGetParamsTlsVersionTlSv1_2}),
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
