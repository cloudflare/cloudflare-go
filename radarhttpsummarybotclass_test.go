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

func TestRadarHTTPSummaryBotClassListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.HTTPs.Summaries.BotClasses.List(context.TODO(), cloudflare.RadarHTTPSummaryBotClassListParams{
		ASN:          cloudflare.F([]string{"15169", "15169", "15169"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPSummaryBotClassListParamsDateRange{cloudflare.RadarHTTPSummaryBotClassListParamsDateRange1d, cloudflare.RadarHTTPSummaryBotClassListParamsDateRange7d, cloudflare.RadarHTTPSummaryBotClassListParamsDateRange14d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPSummaryBotClassListParamsDeviceType{cloudflare.RadarHTTPSummaryBotClassListParamsDeviceTypeDesktop, cloudflare.RadarHTTPSummaryBotClassListParamsDeviceTypeMobile, cloudflare.RadarHTTPSummaryBotClassListParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPSummaryBotClassListParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPSummaryBotClassListParamsHTTPProtocol{cloudflare.RadarHTTPSummaryBotClassListParamsHTTPProtocolHTTP, cloudflare.RadarHTTPSummaryBotClassListParamsHTTPProtocolHTTPs}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPSummaryBotClassListParamsHTTPVersion{cloudflare.RadarHTTPSummaryBotClassListParamsHTTPVersionHttPv1, cloudflare.RadarHTTPSummaryBotClassListParamsHTTPVersionHttPv2, cloudflare.RadarHTTPSummaryBotClassListParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPSummaryBotClassListParamsIPVersion{cloudflare.RadarHTTPSummaryBotClassListParamsIPVersionIPv4, cloudflare.RadarHTTPSummaryBotClassListParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
		Name:         cloudflare.F([]string{"main_series", "main_series", "main_series"}),
		Os:           cloudflare.F([]cloudflare.RadarHTTPSummaryBotClassListParamsO{cloudflare.RadarHTTPSummaryBotClassListParamsOWindows, cloudflare.RadarHTTPSummaryBotClassListParamsOMacosx, cloudflare.RadarHTTPSummaryBotClassListParamsOAndroid}),
		TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPSummaryBotClassListParamsTlsVersion{cloudflare.RadarHTTPSummaryBotClassListParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPSummaryBotClassListParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPSummaryBotClassListParamsTlsVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
