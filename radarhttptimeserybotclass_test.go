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

func TestRadarHTTPTimeseryBotClassListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.HTTPs.Timeseries.BotClasses.List(context.TODO(), cloudflare.RadarHTTPTimeseryBotClassListParams{
		AggInterval:  cloudflare.F(cloudflare.RadarHTTPTimeseryBotClassListParamsAggInterval15m),
		ASN:          cloudflare.F([]string{"15169", "15169", "15169"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPTimeseryBotClassListParamsDateRange{cloudflare.RadarHTTPTimeseryBotClassListParamsDateRange1d, cloudflare.RadarHTTPTimeseryBotClassListParamsDateRange7d, cloudflare.RadarHTTPTimeseryBotClassListParamsDateRange14d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTimeseryBotClassListParamsDeviceType{cloudflare.RadarHTTPTimeseryBotClassListParamsDeviceTypeDesktop, cloudflare.RadarHTTPTimeseryBotClassListParamsDeviceTypeMobile, cloudflare.RadarHTTPTimeseryBotClassListParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPTimeseryBotClassListParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTimeseryBotClassListParamsHTTPProtocol{cloudflare.RadarHTTPTimeseryBotClassListParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTimeseryBotClassListParamsHTTPProtocolHTTPs}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTimeseryBotClassListParamsHTTPVersion{cloudflare.RadarHTTPTimeseryBotClassListParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTimeseryBotClassListParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTimeseryBotClassListParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTimeseryBotClassListParamsIPVersion{cloudflare.RadarHTTPTimeseryBotClassListParamsIPVersionIPv4, cloudflare.RadarHTTPTimeseryBotClassListParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
		Name:         cloudflare.F([]string{"main_series", "main_series", "main_series"}),
		Os:           cloudflare.F([]cloudflare.RadarHTTPTimeseryBotClassListParamsO{cloudflare.RadarHTTPTimeseryBotClassListParamsOWindows, cloudflare.RadarHTTPTimeseryBotClassListParamsOMacosx, cloudflare.RadarHTTPTimeseryBotClassListParamsOAndroid}),
		TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPTimeseryBotClassListParamsTlsVersion{cloudflare.RadarHTTPTimeseryBotClassListParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTimeseryBotClassListParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTimeseryBotClassListParamsTlsVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
