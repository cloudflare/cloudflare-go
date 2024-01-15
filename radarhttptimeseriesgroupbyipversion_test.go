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

func TestRadarHTTPTimeseriesGroupByIPVersionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.ByIPVersion.List(context.TODO(), cloudflare.RadarHTTPTimeseriesGroupByIPVersionListParams{
		AggInterval:  cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupByIPVersionListParamsAggInterval1h),
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByIPVersionListParamsBotClass{cloudflare.RadarHTTPTimeseriesGroupByIPVersionListParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTimeseriesGroupByIPVersionListParamsBotClassLikelyHuman}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange{cloudflare.RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange1d, cloudflare.RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange2d, cloudflare.RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByIPVersionListParamsDeviceType{cloudflare.RadarHTTPTimeseriesGroupByIPVersionListParamsDeviceTypeDesktop, cloudflare.RadarHTTPTimeseriesGroupByIPVersionListParamsDeviceTypeMobile, cloudflare.RadarHTTPTimeseriesGroupByIPVersionListParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupByIPVersionListParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByIPVersionListParamsHTTPProtocol{cloudflare.RadarHTTPTimeseriesGroupByIPVersionListParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTimeseriesGroupByIPVersionListParamsHTTPProtocolHTTPs}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByIPVersionListParamsHTTPVersion{cloudflare.RadarHTTPTimeseriesGroupByIPVersionListParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTimeseriesGroupByIPVersionListParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTimeseriesGroupByIPVersionListParamsHTTPVersionHttPv3}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		Os:           cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByIPVersionListParamsO{cloudflare.RadarHTTPTimeseriesGroupByIPVersionListParamsOWindows, cloudflare.RadarHTTPTimeseriesGroupByIPVersionListParamsOMacosx, cloudflare.RadarHTTPTimeseriesGroupByIPVersionListParamsOIos}),
		TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByIPVersionListParamsTlsVersion{cloudflare.RadarHTTPTimeseriesGroupByIPVersionListParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTimeseriesGroupByIPVersionListParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTimeseriesGroupByIPVersionListParamsTlsVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
