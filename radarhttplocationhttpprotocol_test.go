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

func TestRadarHTTPLocationHTTPProtocolGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Locations.HTTPProtocol.Get(
		context.TODO(),
		cloudflare.RadarHTTPLocationHTTPProtocolGetParamsHTTPProtocolHTTP,
		cloudflare.RadarHTTPLocationHTTPProtocolGetParams{
			ASN:        cloudflare.F([]string{"string", "string", "string"}),
			BotClass:   cloudflare.F([]cloudflare.RadarHTTPLocationHTTPProtocolGetParamsBotClass{cloudflare.RadarHTTPLocationHTTPProtocolGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPLocationHTTPProtocolGetParamsBotClassLikelyHuman}),
			Continent:  cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:  cloudflare.F([]cloudflare.RadarHTTPLocationHTTPProtocolGetParamsDateRange{cloudflare.RadarHTTPLocationHTTPProtocolGetParamsDateRange1d, cloudflare.RadarHTTPLocationHTTPProtocolGetParamsDateRange2d, cloudflare.RadarHTTPLocationHTTPProtocolGetParamsDateRange7d}),
			DateStart:  cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType: cloudflare.F([]cloudflare.RadarHTTPLocationHTTPProtocolGetParamsDeviceType{cloudflare.RadarHTTPLocationHTTPProtocolGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPLocationHTTPProtocolGetParamsDeviceTypeMobile, cloudflare.RadarHTTPLocationHTTPProtocolGetParamsDeviceTypeOther}),
			Format:     cloudflare.F(cloudflare.RadarHTTPLocationHTTPProtocolGetParamsFormatJson),
			IPVersion:  cloudflare.F([]cloudflare.RadarHTTPLocationHTTPProtocolGetParamsIPVersion{cloudflare.RadarHTTPLocationHTTPProtocolGetParamsIPVersionIPv4, cloudflare.RadarHTTPLocationHTTPProtocolGetParamsIPVersionIPv6}),
			Limit:      cloudflare.F(int64(5)),
			Location:   cloudflare.F([]string{"string", "string", "string"}),
			Name:       cloudflare.F([]string{"string", "string", "string"}),
			OS:         cloudflare.F([]cloudflare.RadarHTTPLocationHTTPProtocolGetParamsOS{cloudflare.RadarHTTPLocationHTTPProtocolGetParamsOSWindows, cloudflare.RadarHTTPLocationHTTPProtocolGetParamsOSMacosx, cloudflare.RadarHTTPLocationHTTPProtocolGetParamsOSIos}),
			TLSVersion: cloudflare.F([]cloudflare.RadarHTTPLocationHTTPProtocolGetParamsTLSVersion{cloudflare.RadarHTTPLocationHTTPProtocolGetParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPLocationHTTPProtocolGetParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPLocationHTTPProtocolGetParamsTLSVersionTlSv1_2}),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
