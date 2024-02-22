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

func TestRadarHTTPTopBrowserFamiliesWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Radar.HTTP.Top.BrowserFamilies(context.TODO(), cloudflare.RadarHTTPTopBrowserFamiliesParams{
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPTopBrowserFamiliesParamsBotClass{cloudflare.RadarHTTPTopBrowserFamiliesParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTopBrowserFamiliesParamsBotClassLikelyHuman}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPTopBrowserFamiliesParamsDateRange{cloudflare.RadarHTTPTopBrowserFamiliesParamsDateRange1d, cloudflare.RadarHTTPTopBrowserFamiliesParamsDateRange2d, cloudflare.RadarHTTPTopBrowserFamiliesParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTopBrowserFamiliesParamsDeviceType{cloudflare.RadarHTTPTopBrowserFamiliesParamsDeviceTypeDesktop, cloudflare.RadarHTTPTopBrowserFamiliesParamsDeviceTypeMobile, cloudflare.RadarHTTPTopBrowserFamiliesParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPTopBrowserFamiliesParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTopBrowserFamiliesParamsHTTPProtocol{cloudflare.RadarHTTPTopBrowserFamiliesParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTopBrowserFamiliesParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTopBrowserFamiliesParamsHTTPVersion{cloudflare.RadarHTTPTopBrowserFamiliesParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTopBrowserFamiliesParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTopBrowserFamiliesParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTopBrowserFamiliesParamsIPVersion{cloudflare.RadarHTTPTopBrowserFamiliesParamsIPVersionIPv4, cloudflare.RadarHTTPTopBrowserFamiliesParamsIPVersionIPv6}),
		Limit:        cloudflare.F(int64(5)),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		OS:           cloudflare.F([]cloudflare.RadarHTTPTopBrowserFamiliesParamsOS{cloudflare.RadarHTTPTopBrowserFamiliesParamsOSWindows, cloudflare.RadarHTTPTopBrowserFamiliesParamsOSMacosx, cloudflare.RadarHTTPTopBrowserFamiliesParamsOSIos}),
		TLSVersion:   cloudflare.F([]cloudflare.RadarHTTPTopBrowserFamiliesParamsTLSVersion{cloudflare.RadarHTTPTopBrowserFamiliesParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPTopBrowserFamiliesParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPTopBrowserFamiliesParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPTopBrowsersWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Radar.HTTP.Top.Browsers(context.TODO(), cloudflare.RadarHTTPTopBrowsersParams{
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPTopBrowsersParamsBotClass{cloudflare.RadarHTTPTopBrowsersParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTopBrowsersParamsBotClassLikelyHuman}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPTopBrowsersParamsDateRange{cloudflare.RadarHTTPTopBrowsersParamsDateRange1d, cloudflare.RadarHTTPTopBrowsersParamsDateRange2d, cloudflare.RadarHTTPTopBrowsersParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTopBrowsersParamsDeviceType{cloudflare.RadarHTTPTopBrowsersParamsDeviceTypeDesktop, cloudflare.RadarHTTPTopBrowsersParamsDeviceTypeMobile, cloudflare.RadarHTTPTopBrowsersParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPTopBrowsersParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTopBrowsersParamsHTTPProtocol{cloudflare.RadarHTTPTopBrowsersParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTopBrowsersParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTopBrowsersParamsHTTPVersion{cloudflare.RadarHTTPTopBrowsersParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTopBrowsersParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTopBrowsersParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTopBrowsersParamsIPVersion{cloudflare.RadarHTTPTopBrowsersParamsIPVersionIPv4, cloudflare.RadarHTTPTopBrowsersParamsIPVersionIPv6}),
		Limit:        cloudflare.F(int64(5)),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		OS:           cloudflare.F([]cloudflare.RadarHTTPTopBrowsersParamsOS{cloudflare.RadarHTTPTopBrowsersParamsOSWindows, cloudflare.RadarHTTPTopBrowsersParamsOSMacosx, cloudflare.RadarHTTPTopBrowsersParamsOSIos}),
		TLSVersion:   cloudflare.F([]cloudflare.RadarHTTPTopBrowsersParamsTLSVersion{cloudflare.RadarHTTPTopBrowsersParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPTopBrowsersParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPTopBrowsersParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
