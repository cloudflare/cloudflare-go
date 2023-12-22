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

func TestRadarHTTPTopAseBotClassGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.HTTPs.Tops.Ases.BotClasses.Get(
		context.TODO(),
		cloudflare.RadarHTTPTopAseBotClassGetParamsBotClassLikelyAutomated,
		cloudflare.RadarHTTPTopAseBotClassGetParams{
			ASN:          cloudflare.F([]string{"15169", "15169", "15169"}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPTopAseBotClassGetParamsDateRange{cloudflare.RadarHTTPTopAseBotClassGetParamsDateRange1d, cloudflare.RadarHTTPTopAseBotClassGetParamsDateRange7d, cloudflare.RadarHTTPTopAseBotClassGetParamsDateRange14d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTopAseBotClassGetParamsDeviceType{cloudflare.RadarHTTPTopAseBotClassGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPTopAseBotClassGetParamsDeviceTypeMobile, cloudflare.RadarHTTPTopAseBotClassGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(cloudflare.RadarHTTPTopAseBotClassGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTopAseBotClassGetParamsHTTPProtocol{cloudflare.RadarHTTPTopAseBotClassGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTopAseBotClassGetParamsHTTPProtocolHTTPs}),
			HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTopAseBotClassGetParamsHTTPVersion{cloudflare.RadarHTTPTopAseBotClassGetParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTopAseBotClassGetParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTopAseBotClassGetParamsHTTPVersionHttPv3}),
			IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTopAseBotClassGetParamsIPVersion{cloudflare.RadarHTTPTopAseBotClassGetParamsIPVersionIPv4, cloudflare.RadarHTTPTopAseBotClassGetParamsIPVersionIPv6}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
			Name:         cloudflare.F([]string{"main_series", "main_series", "main_series"}),
			Os:           cloudflare.F([]cloudflare.RadarHTTPTopAseBotClassGetParamsO{cloudflare.RadarHTTPTopAseBotClassGetParamsOWindows, cloudflare.RadarHTTPTopAseBotClassGetParamsOMacosx, cloudflare.RadarHTTPTopAseBotClassGetParamsOAndroid}),
			TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPTopAseBotClassGetParamsTlsVersion{cloudflare.RadarHTTPTopAseBotClassGetParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTopAseBotClassGetParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTopAseBotClassGetParamsTlsVersionTlSv1_2}),
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
