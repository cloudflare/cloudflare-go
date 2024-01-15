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

func TestRadarHTTPTimeseriesGroupByHTTPVersionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.ByHTTPVersion.List(context.TODO(), cloudflare.RadarHTTPTimeseriesGroupByHTTPVersionListParams{
		AggInterval:  cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupByHTTPVersionListParamsAggInterval1h),
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByHTTPVersionListParamsBotClass{cloudflare.RadarHTTPTimeseriesGroupByHTTPVersionListParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTimeseriesGroupByHTTPVersionListParamsBotClassLikelyHuman}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange{cloudflare.RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange1d, cloudflare.RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange2d, cloudflare.RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByHTTPVersionListParamsDeviceType{cloudflare.RadarHTTPTimeseriesGroupByHTTPVersionListParamsDeviceTypeDesktop, cloudflare.RadarHTTPTimeseriesGroupByHTTPVersionListParamsDeviceTypeMobile, cloudflare.RadarHTTPTimeseriesGroupByHTTPVersionListParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupByHTTPVersionListParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByHTTPVersionListParamsHTTPProtocol{cloudflare.RadarHTTPTimeseriesGroupByHTTPVersionListParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTimeseriesGroupByHTTPVersionListParamsHTTPProtocolHTTPs}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByHTTPVersionListParamsIPVersion{cloudflare.RadarHTTPTimeseriesGroupByHTTPVersionListParamsIPVersionIPv4, cloudflare.RadarHTTPTimeseriesGroupByHTTPVersionListParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		Os:           cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByHTTPVersionListParamsO{cloudflare.RadarHTTPTimeseriesGroupByHTTPVersionListParamsOWindows, cloudflare.RadarHTTPTimeseriesGroupByHTTPVersionListParamsOMacosx, cloudflare.RadarHTTPTimeseriesGroupByHTTPVersionListParamsOIos}),
		TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByHTTPVersionListParamsTlsVersion{cloudflare.RadarHTTPTimeseriesGroupByHTTPVersionListParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTimeseriesGroupByHTTPVersionListParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTimeseriesGroupByHTTPVersionListParamsTlsVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
