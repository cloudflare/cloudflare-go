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

func TestRadarHTTPTopAseIPVersionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.HTTPs.Tops.Ases.IPVersions.Get(
		context.TODO(),
		cloudflare.RadarHTTPTopAseIPVersionGetParamsIPVersionIPv4,
		cloudflare.RadarHTTPTopAseIPVersionGetParams{
			ASN:          cloudflare.F([]string{"15169", "15169", "15169"}),
			BotClass:     cloudflare.F([]cloudflare.RadarHTTPTopAseIPVersionGetParamsBotClass{cloudflare.RadarHTTPTopAseIPVersionGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTopAseIPVersionGetParamsBotClassLikelyHuman}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPTopAseIPVersionGetParamsDateRange{cloudflare.RadarHTTPTopAseIPVersionGetParamsDateRange1d, cloudflare.RadarHTTPTopAseIPVersionGetParamsDateRange7d, cloudflare.RadarHTTPTopAseIPVersionGetParamsDateRange14d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTopAseIPVersionGetParamsDeviceType{cloudflare.RadarHTTPTopAseIPVersionGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPTopAseIPVersionGetParamsDeviceTypeMobile, cloudflare.RadarHTTPTopAseIPVersionGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(cloudflare.RadarHTTPTopAseIPVersionGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTopAseIPVersionGetParamsHTTPProtocol{cloudflare.RadarHTTPTopAseIPVersionGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTopAseIPVersionGetParamsHTTPProtocolHTTPs}),
			HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTopAseIPVersionGetParamsHTTPVersion{cloudflare.RadarHTTPTopAseIPVersionGetParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTopAseIPVersionGetParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTopAseIPVersionGetParamsHTTPVersionHttPv3}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
			Name:         cloudflare.F([]string{"main_series", "main_series", "main_series"}),
			Os:           cloudflare.F([]cloudflare.RadarHTTPTopAseIPVersionGetParamsO{cloudflare.RadarHTTPTopAseIPVersionGetParamsOWindows, cloudflare.RadarHTTPTopAseIPVersionGetParamsOMacosx, cloudflare.RadarHTTPTopAseIPVersionGetParamsOAndroid}),
			TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPTopAseIPVersionGetParamsTlsVersion{cloudflare.RadarHTTPTopAseIPVersionGetParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTopAseIPVersionGetParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTopAseIPVersionGetParamsTlsVersionTlSv1_2}),
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
