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

func TestRadarHTTPTimeseriesGroupByBotClassListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.ByBotClass.List(context.TODO(), cloudflare.RadarHTTPTimeseriesGroupByBotClassListParams{
		AggInterval:  cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupByBotClassListParamsAggInterval1h),
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByBotClassListParamsDateRange{cloudflare.RadarHTTPTimeseriesGroupByBotClassListParamsDateRange1d, cloudflare.RadarHTTPTimeseriesGroupByBotClassListParamsDateRange2d, cloudflare.RadarHTTPTimeseriesGroupByBotClassListParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByBotClassListParamsDeviceType{cloudflare.RadarHTTPTimeseriesGroupByBotClassListParamsDeviceTypeDesktop, cloudflare.RadarHTTPTimeseriesGroupByBotClassListParamsDeviceTypeMobile, cloudflare.RadarHTTPTimeseriesGroupByBotClassListParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupByBotClassListParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByBotClassListParamsHTTPProtocol{cloudflare.RadarHTTPTimeseriesGroupByBotClassListParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTimeseriesGroupByBotClassListParamsHTTPProtocolHTTPs}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByBotClassListParamsHTTPVersion{cloudflare.RadarHTTPTimeseriesGroupByBotClassListParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTimeseriesGroupByBotClassListParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTimeseriesGroupByBotClassListParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByBotClassListParamsIPVersion{cloudflare.RadarHTTPTimeseriesGroupByBotClassListParamsIPVersionIPv4, cloudflare.RadarHTTPTimeseriesGroupByBotClassListParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		Os:           cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByBotClassListParamsO{cloudflare.RadarHTTPTimeseriesGroupByBotClassListParamsOWindows, cloudflare.RadarHTTPTimeseriesGroupByBotClassListParamsOMacosx, cloudflare.RadarHTTPTimeseriesGroupByBotClassListParamsOIos}),
		TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByBotClassListParamsTlsVersion{cloudflare.RadarHTTPTimeseriesGroupByBotClassListParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTimeseriesGroupByBotClassListParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTimeseriesGroupByBotClassListParamsTlsVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
