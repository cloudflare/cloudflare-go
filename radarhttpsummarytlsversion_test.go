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

func TestRadarHTTPSummaryTlsVersionListWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Radar.HTTP.Summaries.TlsVersions.List(context.TODO(), cloudflare.RadarHTTPSummaryTlsVersionListParams{
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPSummaryTlsVersionListParamsBotClass{cloudflare.RadarHTTPSummaryTlsVersionListParamsBotClassLikelyAutomated, cloudflare.RadarHTTPSummaryTlsVersionListParamsBotClassLikelyHuman}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPSummaryTlsVersionListParamsDateRange{cloudflare.RadarHTTPSummaryTlsVersionListParamsDateRange1d, cloudflare.RadarHTTPSummaryTlsVersionListParamsDateRange2d, cloudflare.RadarHTTPSummaryTlsVersionListParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPSummaryTlsVersionListParamsDeviceType{cloudflare.RadarHTTPSummaryTlsVersionListParamsDeviceTypeDesktop, cloudflare.RadarHTTPSummaryTlsVersionListParamsDeviceTypeMobile, cloudflare.RadarHTTPSummaryTlsVersionListParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPSummaryTlsVersionListParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPSummaryTlsVersionListParamsHTTPProtocol{cloudflare.RadarHTTPSummaryTlsVersionListParamsHTTPProtocolHTTP, cloudflare.RadarHTTPSummaryTlsVersionListParamsHTTPProtocolHTTPs}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPSummaryTlsVersionListParamsHTTPVersion{cloudflare.RadarHTTPSummaryTlsVersionListParamsHTTPVersionHttPv1, cloudflare.RadarHTTPSummaryTlsVersionListParamsHTTPVersionHttPv2, cloudflare.RadarHTTPSummaryTlsVersionListParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPSummaryTlsVersionListParamsIPVersion{cloudflare.RadarHTTPSummaryTlsVersionListParamsIPVersionIPv4, cloudflare.RadarHTTPSummaryTlsVersionListParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		Os:           cloudflare.F([]cloudflare.RadarHTTPSummaryTlsVersionListParamsO{cloudflare.RadarHTTPSummaryTlsVersionListParamsOWindows, cloudflare.RadarHTTPSummaryTlsVersionListParamsOMacosx, cloudflare.RadarHTTPSummaryTlsVersionListParamsOIos}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
