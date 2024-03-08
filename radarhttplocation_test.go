// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
)

func TestRadarHTTPLocationGetWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Radar.HTTP.Locations.Get(context.TODO(), cloudflare.RadarHTTPLocationGetParams{
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPLocationGetParamsBotClass{cloudflare.RadarHTTPLocationGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPLocationGetParamsBotClassLikelyHuman}),
		Continent:    cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPLocationGetParamsDateRange{cloudflare.RadarHTTPLocationGetParamsDateRange1d, cloudflare.RadarHTTPLocationGetParamsDateRange2d, cloudflare.RadarHTTPLocationGetParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPLocationGetParamsDeviceType{cloudflare.RadarHTTPLocationGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPLocationGetParamsDeviceTypeMobile, cloudflare.RadarHTTPLocationGetParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPLocationGetParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPLocationGetParamsHTTPProtocol{cloudflare.RadarHTTPLocationGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPLocationGetParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPLocationGetParamsHTTPVersion{cloudflare.RadarHTTPLocationGetParamsHTTPVersionHttPv1, cloudflare.RadarHTTPLocationGetParamsHTTPVersionHttPv2, cloudflare.RadarHTTPLocationGetParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPLocationGetParamsIPVersion{cloudflare.RadarHTTPLocationGetParamsIPVersionIPv4, cloudflare.RadarHTTPLocationGetParamsIPVersionIPv6}),
		Limit:        cloudflare.F(int64(5)),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		OS:           cloudflare.F([]cloudflare.RadarHTTPLocationGetParamsOS{cloudflare.RadarHTTPLocationGetParamsOSWindows, cloudflare.RadarHTTPLocationGetParamsOSMacosx, cloudflare.RadarHTTPLocationGetParamsOSIos}),
		TLSVersion:   cloudflare.F([]cloudflare.RadarHTTPLocationGetParamsTLSVersion{cloudflare.RadarHTTPLocationGetParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPLocationGetParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPLocationGetParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
