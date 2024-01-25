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

func TestRadarHTTPTimeseriesGroupByDeviceTypeListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.ByDeviceType.List(context.TODO(), cloudflare.RadarHTTPTimeseriesGroupByDeviceTypeListParams{
		AggInterval:  cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupByDeviceTypeListParamsAggInterval1h),
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByDeviceTypeListParamsBotClass{cloudflare.RadarHTTPTimeseriesGroupByDeviceTypeListParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTimeseriesGroupByDeviceTypeListParamsBotClassLikelyHuman}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange{cloudflare.RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange1d, cloudflare.RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange2d, cloudflare.RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:       cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupByDeviceTypeListParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByDeviceTypeListParamsHTTPProtocol{cloudflare.RadarHTTPTimeseriesGroupByDeviceTypeListParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTimeseriesGroupByDeviceTypeListParamsHTTPProtocolHTTPs}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByDeviceTypeListParamsHTTPVersion{cloudflare.RadarHTTPTimeseriesGroupByDeviceTypeListParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTimeseriesGroupByDeviceTypeListParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTimeseriesGroupByDeviceTypeListParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByDeviceTypeListParamsIPVersion{cloudflare.RadarHTTPTimeseriesGroupByDeviceTypeListParamsIPVersionIPv4, cloudflare.RadarHTTPTimeseriesGroupByDeviceTypeListParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		Os:           cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByDeviceTypeListParamsO{cloudflare.RadarHTTPTimeseriesGroupByDeviceTypeListParamsOWindows, cloudflare.RadarHTTPTimeseriesGroupByDeviceTypeListParamsOMacosx, cloudflare.RadarHTTPTimeseriesGroupByDeviceTypeListParamsOIos}),
		TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByDeviceTypeListParamsTlsVersion{cloudflare.RadarHTTPTimeseriesGroupByDeviceTypeListParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTimeseriesGroupByDeviceTypeListParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTimeseriesGroupByDeviceTypeListParamsTlsVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
