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

func TestRadarHTTPSummaryOListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.HTTPs.Summaries.Os.List(context.TODO(), cloudflare.RadarHTTPSummaryOListParams{
		ASN:          cloudflare.F([]string{"15169", "15169", "15169"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPSummaryOListParamsBotClass{cloudflare.RadarHTTPSummaryOListParamsBotClassLikelyAutomated, cloudflare.RadarHTTPSummaryOListParamsBotClassLikelyHuman}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPSummaryOListParamsDateRange{cloudflare.RadarHTTPSummaryOListParamsDateRange1d, cloudflare.RadarHTTPSummaryOListParamsDateRange7d, cloudflare.RadarHTTPSummaryOListParamsDateRange14d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPSummaryOListParamsDeviceType{cloudflare.RadarHTTPSummaryOListParamsDeviceTypeDesktop, cloudflare.RadarHTTPSummaryOListParamsDeviceTypeMobile, cloudflare.RadarHTTPSummaryOListParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPSummaryOListParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPSummaryOListParamsHTTPProtocol{cloudflare.RadarHTTPSummaryOListParamsHTTPProtocolHTTP, cloudflare.RadarHTTPSummaryOListParamsHTTPProtocolHTTPs}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPSummaryOListParamsHTTPVersion{cloudflare.RadarHTTPSummaryOListParamsHTTPVersionHttPv1, cloudflare.RadarHTTPSummaryOListParamsHTTPVersionHttPv2, cloudflare.RadarHTTPSummaryOListParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPSummaryOListParamsIPVersion{cloudflare.RadarHTTPSummaryOListParamsIPVersionIPv4, cloudflare.RadarHTTPSummaryOListParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
		Name:         cloudflare.F([]string{"main_series", "main_series", "main_series"}),
		TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPSummaryOListParamsTlsVersion{cloudflare.RadarHTTPSummaryOListParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPSummaryOListParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPSummaryOListParamsTlsVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
