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

func TestRadarHTTPTopLocationIPVersionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.HTTPs.Tops.Locations.IPVersions.Get(
		context.TODO(),
		cloudflare.RadarHTTPTopLocationIPVersionGetParamsIPVersionIPv4,
		cloudflare.RadarHTTPTopLocationIPVersionGetParams{
			ASN:          cloudflare.F([]string{"15169", "15169", "15169"}),
			BotClass:     cloudflare.F([]cloudflare.RadarHTTPTopLocationIPVersionGetParamsBotClass{cloudflare.RadarHTTPTopLocationIPVersionGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTopLocationIPVersionGetParamsBotClassLikelyHuman}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPTopLocationIPVersionGetParamsDateRange{cloudflare.RadarHTTPTopLocationIPVersionGetParamsDateRange1d, cloudflare.RadarHTTPTopLocationIPVersionGetParamsDateRange7d, cloudflare.RadarHTTPTopLocationIPVersionGetParamsDateRange14d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTopLocationIPVersionGetParamsDeviceType{cloudflare.RadarHTTPTopLocationIPVersionGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPTopLocationIPVersionGetParamsDeviceTypeMobile, cloudflare.RadarHTTPTopLocationIPVersionGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(cloudflare.RadarHTTPTopLocationIPVersionGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTopLocationIPVersionGetParamsHTTPProtocol{cloudflare.RadarHTTPTopLocationIPVersionGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTopLocationIPVersionGetParamsHTTPProtocolHTTPs}),
			HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTopLocationIPVersionGetParamsHTTPVersion{cloudflare.RadarHTTPTopLocationIPVersionGetParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTopLocationIPVersionGetParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTopLocationIPVersionGetParamsHTTPVersionHttPv3}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
			Name:         cloudflare.F([]string{"main_series", "main_series", "main_series"}),
			Os:           cloudflare.F([]cloudflare.RadarHTTPTopLocationIPVersionGetParamsO{cloudflare.RadarHTTPTopLocationIPVersionGetParamsOWindows, cloudflare.RadarHTTPTopLocationIPVersionGetParamsOMacosx, cloudflare.RadarHTTPTopLocationIPVersionGetParamsOAndroid}),
			TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPTopLocationIPVersionGetParamsTlsVersion{cloudflare.RadarHTTPTopLocationIPVersionGetParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTopLocationIPVersionGetParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTopLocationIPVersionGetParamsTlsVersionTlSv1_2}),
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
