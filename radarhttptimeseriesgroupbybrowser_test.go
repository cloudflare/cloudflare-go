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

func TestRadarHTTPTimeseriesGroupByBrowserListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.ByBrowser.List(context.TODO(), cloudflare.RadarHTTPTimeseriesGroupByBrowserListParams{
		AggInterval:   cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupByBrowserListParamsAggInterval1h),
		ASN:           cloudflare.F([]string{"string", "string", "string"}),
		BotClass:      cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByBrowserListParamsBotClass{cloudflare.RadarHTTPTimeseriesGroupByBrowserListParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTimeseriesGroupByBrowserListParamsBotClassLikelyHuman}),
		DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByBrowserListParamsDateRange{cloudflare.RadarHTTPTimeseriesGroupByBrowserListParamsDateRange1d, cloudflare.RadarHTTPTimeseriesGroupByBrowserListParamsDateRange2d, cloudflare.RadarHTTPTimeseriesGroupByBrowserListParamsDateRange7d}),
		DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:    cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByBrowserListParamsDeviceType{cloudflare.RadarHTTPTimeseriesGroupByBrowserListParamsDeviceTypeDesktop, cloudflare.RadarHTTPTimeseriesGroupByBrowserListParamsDeviceTypeMobile, cloudflare.RadarHTTPTimeseriesGroupByBrowserListParamsDeviceTypeOther}),
		Format:        cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupByBrowserListParamsFormatJson),
		HTTPProtocol:  cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByBrowserListParamsHTTPProtocol{cloudflare.RadarHTTPTimeseriesGroupByBrowserListParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTimeseriesGroupByBrowserListParamsHTTPProtocolHTTPs}),
		HTTPVersion:   cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByBrowserListParamsHTTPVersion{cloudflare.RadarHTTPTimeseriesGroupByBrowserListParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTimeseriesGroupByBrowserListParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTimeseriesGroupByBrowserListParamsHTTPVersionHttPv3}),
		IPVersion:     cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByBrowserListParamsIPVersion{cloudflare.RadarHTTPTimeseriesGroupByBrowserListParamsIPVersionIPv4, cloudflare.RadarHTTPTimeseriesGroupByBrowserListParamsIPVersionIPv6}),
		LimitPerGroup: cloudflare.F(int64(4)),
		Location:      cloudflare.F([]string{"string", "string", "string"}),
		Name:          cloudflare.F([]string{"string", "string", "string"}),
		Os:            cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByBrowserListParamsO{cloudflare.RadarHTTPTimeseriesGroupByBrowserListParamsOWindows, cloudflare.RadarHTTPTimeseriesGroupByBrowserListParamsOMacosx, cloudflare.RadarHTTPTimeseriesGroupByBrowserListParamsOIos}),
		TlsVersion:    cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByBrowserListParamsTlsVersion{cloudflare.RadarHTTPTimeseriesGroupByBrowserListParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTimeseriesGroupByBrowserListParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTimeseriesGroupByBrowserListParamsTlsVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
