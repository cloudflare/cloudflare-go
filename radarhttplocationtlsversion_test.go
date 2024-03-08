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

func TestRadarHTTPLocationTLSVersionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Locations.TLSVersion.Get(
		context.TODO(),
		cloudflare.RadarHTTPLocationTLSVersionGetParamsTLSVersionTlSv1_0,
		cloudflare.RadarHTTPLocationTLSVersionGetParams{
			ASN:          cloudflare.F([]string{"string", "string", "string"}),
			BotClass:     cloudflare.F([]cloudflare.RadarHTTPLocationTLSVersionGetParamsBotClass{cloudflare.RadarHTTPLocationTLSVersionGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPLocationTLSVersionGetParamsBotClassLikelyHuman}),
			Continent:    cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPLocationTLSVersionGetParamsDateRange{cloudflare.RadarHTTPLocationTLSVersionGetParamsDateRange1d, cloudflare.RadarHTTPLocationTLSVersionGetParamsDateRange2d, cloudflare.RadarHTTPLocationTLSVersionGetParamsDateRange7d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]cloudflare.RadarHTTPLocationTLSVersionGetParamsDeviceType{cloudflare.RadarHTTPLocationTLSVersionGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPLocationTLSVersionGetParamsDeviceTypeMobile, cloudflare.RadarHTTPLocationTLSVersionGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(cloudflare.RadarHTTPLocationTLSVersionGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPLocationTLSVersionGetParamsHTTPProtocol{cloudflare.RadarHTTPLocationTLSVersionGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPLocationTLSVersionGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPLocationTLSVersionGetParamsHTTPVersion{cloudflare.RadarHTTPLocationTLSVersionGetParamsHTTPVersionHttPv1, cloudflare.RadarHTTPLocationTLSVersionGetParamsHTTPVersionHttPv2, cloudflare.RadarHTTPLocationTLSVersionGetParamsHTTPVersionHttPv3}),
			IPVersion:    cloudflare.F([]cloudflare.RadarHTTPLocationTLSVersionGetParamsIPVersion{cloudflare.RadarHTTPLocationTLSVersionGetParamsIPVersionIPv4, cloudflare.RadarHTTPLocationTLSVersionGetParamsIPVersionIPv6}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"string", "string", "string"}),
			Name:         cloudflare.F([]string{"string", "string", "string"}),
			OS:           cloudflare.F([]cloudflare.RadarHTTPLocationTLSVersionGetParamsOS{cloudflare.RadarHTTPLocationTLSVersionGetParamsOSWindows, cloudflare.RadarHTTPLocationTLSVersionGetParamsOSMacosx, cloudflare.RadarHTTPLocationTLSVersionGetParamsOSIos}),
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
