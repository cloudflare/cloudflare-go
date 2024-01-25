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

func TestRadarHTTPTimeseriesGroupByBrowserFamilyListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.ByBrowserFamily.List(context.TODO(), cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParams{
		AggInterval:  cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParamsAggInterval1h),
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParamsBotClass{cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParamsBotClassLikelyHuman}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange{cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange1d, cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange2d, cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDeviceType{cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDeviceTypeDesktop, cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDeviceTypeMobile, cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParamsHTTPProtocol{cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParamsHTTPProtocolHTTPs}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParamsHTTPVersion{cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParamsIPVersion{cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParamsIPVersionIPv4, cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		Os:           cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParamsO{cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParamsOWindows, cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParamsOMacosx, cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParamsOIos}),
		TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParamsTlsVersion{cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTimeseriesGroupByBrowserFamilyListParamsTlsVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
