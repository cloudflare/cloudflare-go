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

func TestRadarHTTPSummaryDeviceTypeListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.HTTPs.Summaries.DeviceTypes.List(context.TODO(), cloudflare.RadarHTTPSummaryDeviceTypeListParams{
		ASN:          cloudflare.F([]string{"15169", "15169", "15169"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPSummaryDeviceTypeListParamsBotClass{cloudflare.RadarHTTPSummaryDeviceTypeListParamsBotClassLikelyAutomated, cloudflare.RadarHTTPSummaryDeviceTypeListParamsBotClassLikelyHuman}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPSummaryDeviceTypeListParamsDateRange{cloudflare.RadarHTTPSummaryDeviceTypeListParamsDateRange1d, cloudflare.RadarHTTPSummaryDeviceTypeListParamsDateRange7d, cloudflare.RadarHTTPSummaryDeviceTypeListParamsDateRange14d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:       cloudflare.F(cloudflare.RadarHTTPSummaryDeviceTypeListParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPSummaryDeviceTypeListParamsHTTPProtocol{cloudflare.RadarHTTPSummaryDeviceTypeListParamsHTTPProtocolHTTP, cloudflare.RadarHTTPSummaryDeviceTypeListParamsHTTPProtocolHTTPs}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPSummaryDeviceTypeListParamsHTTPVersion{cloudflare.RadarHTTPSummaryDeviceTypeListParamsHTTPVersionHttPv1, cloudflare.RadarHTTPSummaryDeviceTypeListParamsHTTPVersionHttPv2, cloudflare.RadarHTTPSummaryDeviceTypeListParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPSummaryDeviceTypeListParamsIPVersion{cloudflare.RadarHTTPSummaryDeviceTypeListParamsIPVersionIPv4, cloudflare.RadarHTTPSummaryDeviceTypeListParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
		Name:         cloudflare.F([]string{"main_series", "main_series", "main_series"}),
		Os:           cloudflare.F([]cloudflare.RadarHTTPSummaryDeviceTypeListParamsO{cloudflare.RadarHTTPSummaryDeviceTypeListParamsOWindows, cloudflare.RadarHTTPSummaryDeviceTypeListParamsOMacosx, cloudflare.RadarHTTPSummaryDeviceTypeListParamsOAndroid}),
		TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPSummaryDeviceTypeListParamsTlsVersion{cloudflare.RadarHTTPSummaryDeviceTypeListParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPSummaryDeviceTypeListParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPSummaryDeviceTypeListParamsTlsVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
