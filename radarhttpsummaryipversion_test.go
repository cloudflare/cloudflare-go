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

func TestRadarHTTPSummaryIPVersionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.HTTPs.Summaries.IPVersions.List(context.TODO(), cloudflare.RadarHTTPSummaryIPVersionListParams{
		ASN:          cloudflare.F([]string{"15169", "15169", "15169"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPSummaryIPVersionListParamsBotClass{cloudflare.RadarHTTPSummaryIPVersionListParamsBotClassLikelyAutomated, cloudflare.RadarHTTPSummaryIPVersionListParamsBotClassLikelyHuman}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPSummaryIPVersionListParamsDateRange{cloudflare.RadarHTTPSummaryIPVersionListParamsDateRange1d, cloudflare.RadarHTTPSummaryIPVersionListParamsDateRange7d, cloudflare.RadarHTTPSummaryIPVersionListParamsDateRange14d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPSummaryIPVersionListParamsDeviceType{cloudflare.RadarHTTPSummaryIPVersionListParamsDeviceTypeDesktop, cloudflare.RadarHTTPSummaryIPVersionListParamsDeviceTypeMobile, cloudflare.RadarHTTPSummaryIPVersionListParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPSummaryIPVersionListParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPSummaryIPVersionListParamsHTTPProtocol{cloudflare.RadarHTTPSummaryIPVersionListParamsHTTPProtocolHTTP, cloudflare.RadarHTTPSummaryIPVersionListParamsHTTPProtocolHTTPs}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPSummaryIPVersionListParamsHTTPVersion{cloudflare.RadarHTTPSummaryIPVersionListParamsHTTPVersionHttPv1, cloudflare.RadarHTTPSummaryIPVersionListParamsHTTPVersionHttPv2, cloudflare.RadarHTTPSummaryIPVersionListParamsHTTPVersionHttPv3}),
		Location:     cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
		Name:         cloudflare.F([]string{"main_series", "main_series", "main_series"}),
		Os:           cloudflare.F([]cloudflare.RadarHTTPSummaryIPVersionListParamsO{cloudflare.RadarHTTPSummaryIPVersionListParamsOWindows, cloudflare.RadarHTTPSummaryIPVersionListParamsOMacosx, cloudflare.RadarHTTPSummaryIPVersionListParamsOAndroid}),
		TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPSummaryIPVersionListParamsTlsVersion{cloudflare.RadarHTTPSummaryIPVersionListParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPSummaryIPVersionListParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPSummaryIPVersionListParamsTlsVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
