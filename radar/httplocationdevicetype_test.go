// File generated from our OpenAPI spec by Stainless.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/cloudflare/cloudflare-go/radar"
)

func TestHTTPLocationDeviceTypeGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Locations.DeviceType.Get(
		context.TODO(),
		radar.HTTPLocationDeviceTypeGetParamsDeviceTypeDesktop,
		radar.HTTPLocationDeviceTypeGetParams{
			ASN:          cloudflare.F([]string{"string", "string", "string"}),
			BotClass:     cloudflare.F([]radar.HTTPLocationDeviceTypeGetParamsBotClass{radar.HTTPLocationDeviceTypeGetParamsBotClassLikelyAutomated, radar.HTTPLocationDeviceTypeGetParamsBotClassLikelyHuman}),
			Continent:    cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]radar.HTTPLocationDeviceTypeGetParamsDateRange{radar.HTTPLocationDeviceTypeGetParamsDateRange1d, radar.HTTPLocationDeviceTypeGetParamsDateRange2d, radar.HTTPLocationDeviceTypeGetParamsDateRange7d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			Format:       cloudflare.F(radar.HTTPLocationDeviceTypeGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]radar.HTTPLocationDeviceTypeGetParamsHTTPProtocol{radar.HTTPLocationDeviceTypeGetParamsHTTPProtocolHTTP, radar.HTTPLocationDeviceTypeGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:  cloudflare.F([]radar.HTTPLocationDeviceTypeGetParamsHTTPVersion{radar.HTTPLocationDeviceTypeGetParamsHTTPVersionHttPv1, radar.HTTPLocationDeviceTypeGetParamsHTTPVersionHttPv2, radar.HTTPLocationDeviceTypeGetParamsHTTPVersionHttPv3}),
			IPVersion:    cloudflare.F([]radar.HTTPLocationDeviceTypeGetParamsIPVersion{radar.HTTPLocationDeviceTypeGetParamsIPVersionIPv4, radar.HTTPLocationDeviceTypeGetParamsIPVersionIPv6}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"string", "string", "string"}),
			Name:         cloudflare.F([]string{"string", "string", "string"}),
			OS:           cloudflare.F([]radar.HTTPLocationDeviceTypeGetParamsOS{radar.HTTPLocationDeviceTypeGetParamsOSWindows, radar.HTTPLocationDeviceTypeGetParamsOSMacosx, radar.HTTPLocationDeviceTypeGetParamsOSIos}),
			TLSVersion:   cloudflare.F([]radar.HTTPLocationDeviceTypeGetParamsTLSVersion{radar.HTTPLocationDeviceTypeGetParamsTLSVersionTlSv1_0, radar.HTTPLocationDeviceTypeGetParamsTLSVersionTlSv1_1, radar.HTTPLocationDeviceTypeGetParamsTLSVersionTlSv1_2}),
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
