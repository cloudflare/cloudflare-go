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

func TestRadarHTTPTopAseDeviceTypeGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.HTTPs.Tops.Ases.DeviceTypes.Get(
		context.TODO(),
		cloudflare.RadarHTTPTopAseDeviceTypeGetParamsDeviceTypeDesktop,
		cloudflare.RadarHTTPTopAseDeviceTypeGetParams{
			ASN:          cloudflare.F([]string{"15169", "15169", "15169"}),
			BotClass:     cloudflare.F([]cloudflare.RadarHTTPTopAseDeviceTypeGetParamsBotClass{cloudflare.RadarHTTPTopAseDeviceTypeGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTopAseDeviceTypeGetParamsBotClassLikelyHuman}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPTopAseDeviceTypeGetParamsDateRange{cloudflare.RadarHTTPTopAseDeviceTypeGetParamsDateRange1d, cloudflare.RadarHTTPTopAseDeviceTypeGetParamsDateRange7d, cloudflare.RadarHTTPTopAseDeviceTypeGetParamsDateRange14d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			Format:       cloudflare.F(cloudflare.RadarHTTPTopAseDeviceTypeGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTopAseDeviceTypeGetParamsHTTPProtocol{cloudflare.RadarHTTPTopAseDeviceTypeGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTopAseDeviceTypeGetParamsHTTPProtocolHTTPs}),
			HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTopAseDeviceTypeGetParamsHTTPVersion{cloudflare.RadarHTTPTopAseDeviceTypeGetParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTopAseDeviceTypeGetParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTopAseDeviceTypeGetParamsHTTPVersionHttPv3}),
			IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTopAseDeviceTypeGetParamsIPVersion{cloudflare.RadarHTTPTopAseDeviceTypeGetParamsIPVersionIPv4, cloudflare.RadarHTTPTopAseDeviceTypeGetParamsIPVersionIPv6}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
			Name:         cloudflare.F([]string{"main_series", "main_series", "main_series"}),
			Os:           cloudflare.F([]cloudflare.RadarHTTPTopAseDeviceTypeGetParamsO{cloudflare.RadarHTTPTopAseDeviceTypeGetParamsOWindows, cloudflare.RadarHTTPTopAseDeviceTypeGetParamsOMacosx, cloudflare.RadarHTTPTopAseDeviceTypeGetParamsOAndroid}),
			TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPTopAseDeviceTypeGetParamsTlsVersion{cloudflare.RadarHTTPTopAseDeviceTypeGetParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTopAseDeviceTypeGetParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTopAseDeviceTypeGetParamsTlsVersionTlSv1_2}),
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
