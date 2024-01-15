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

func TestRadarHTTPSummaryHTTPProtocolListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summaries.HTTPProtocols.List(context.TODO(), cloudflare.RadarHTTPSummaryHTTPProtocolListParams{
		ASN:         cloudflare.F([]string{"string", "string", "string"}),
		BotClass:    cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPProtocolListParamsBotClass{cloudflare.RadarHTTPSummaryHTTPProtocolListParamsBotClassLikelyAutomated, cloudflare.RadarHTTPSummaryHTTPProtocolListParamsBotClassLikelyHuman}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPProtocolListParamsDateRange{cloudflare.RadarHTTPSummaryHTTPProtocolListParamsDateRange1d, cloudflare.RadarHTTPSummaryHTTPProtocolListParamsDateRange2d, cloudflare.RadarHTTPSummaryHTTPProtocolListParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:  cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPProtocolListParamsDeviceType{cloudflare.RadarHTTPSummaryHTTPProtocolListParamsDeviceTypeDesktop, cloudflare.RadarHTTPSummaryHTTPProtocolListParamsDeviceTypeMobile, cloudflare.RadarHTTPSummaryHTTPProtocolListParamsDeviceTypeOther}),
		Format:      cloudflare.F(cloudflare.RadarHTTPSummaryHTTPProtocolListParamsFormatJson),
		HTTPVersion: cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPProtocolListParamsHTTPVersion{cloudflare.RadarHTTPSummaryHTTPProtocolListParamsHTTPVersionHttPv1, cloudflare.RadarHTTPSummaryHTTPProtocolListParamsHTTPVersionHttPv2, cloudflare.RadarHTTPSummaryHTTPProtocolListParamsHTTPVersionHttPv3}),
		IPVersion:   cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPProtocolListParamsIPVersion{cloudflare.RadarHTTPSummaryHTTPProtocolListParamsIPVersionIPv4, cloudflare.RadarHTTPSummaryHTTPProtocolListParamsIPVersionIPv6}),
		Location:    cloudflare.F([]string{"string", "string", "string"}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		Os:          cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPProtocolListParamsO{cloudflare.RadarHTTPSummaryHTTPProtocolListParamsOWindows, cloudflare.RadarHTTPSummaryHTTPProtocolListParamsOMacosx, cloudflare.RadarHTTPSummaryHTTPProtocolListParamsOIos}),
		TlsVersion:  cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPProtocolListParamsTlsVersion{cloudflare.RadarHTTPSummaryHTTPProtocolListParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPSummaryHTTPProtocolListParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPSummaryHTTPProtocolListParamsTlsVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
