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

func TestRadarHTTPTopLocationHTTPVersionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.HTTPs.Tops.Locations.HTTPVersions.Get(
		context.TODO(),
		cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsHTTPVersionHttPv1,
		cloudflare.RadarHTTPTopLocationHTTPVersionGetParams{
			ASN:          cloudflare.F([]string{"15169", "15169", "15169"}),
			BotClass:     cloudflare.F([]cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsBotClass{cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsBotClassLikelyHuman}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsDateRange{cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsDateRange1d, cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsDateRange7d, cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsDateRange14d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsDeviceType{cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsDeviceTypeMobile, cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsHTTPProtocol{cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsHTTPProtocolHTTPs}),
			IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsIPVersion{cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsIPVersionIPv4, cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsIPVersionIPv6}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
			Name:         cloudflare.F([]string{"main_series", "main_series", "main_series"}),
			Os:           cloudflare.F([]cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsO{cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsOWindows, cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsOMacosx, cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsOAndroid}),
			TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsTlsVersion{cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsTlsVersionTlSv1_2}),
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
