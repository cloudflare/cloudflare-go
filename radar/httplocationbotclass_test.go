// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/radar"
)

func TestHTTPLocationBotClassGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Locations.BotClass.Get(
		context.TODO(),
		radar.HTTPLocationBotClassGetParamsBotClassLikelyAutomated,
		radar.HTTPLocationBotClassGetParams{
			ASN:          cloudflare.F([]string{"string", "string", "string"}),
			Continent:    cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]radar.HTTPLocationBotClassGetParamsDateRange{radar.HTTPLocationBotClassGetParamsDateRange1d, radar.HTTPLocationBotClassGetParamsDateRange2d, radar.HTTPLocationBotClassGetParamsDateRange7d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]radar.HTTPLocationBotClassGetParamsDeviceType{radar.HTTPLocationBotClassGetParamsDeviceTypeDesktop, radar.HTTPLocationBotClassGetParamsDeviceTypeMobile, radar.HTTPLocationBotClassGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(radar.HTTPLocationBotClassGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]radar.HTTPLocationBotClassGetParamsHTTPProtocol{radar.HTTPLocationBotClassGetParamsHTTPProtocolHTTP, radar.HTTPLocationBotClassGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:  cloudflare.F([]radar.HTTPLocationBotClassGetParamsHTTPVersion{radar.HTTPLocationBotClassGetParamsHTTPVersionHttPv1, radar.HTTPLocationBotClassGetParamsHTTPVersionHttPv2, radar.HTTPLocationBotClassGetParamsHTTPVersionHttPv3}),
			IPVersion:    cloudflare.F([]radar.HTTPLocationBotClassGetParamsIPVersion{radar.HTTPLocationBotClassGetParamsIPVersionIPv4, radar.HTTPLocationBotClassGetParamsIPVersionIPv6}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"string", "string", "string"}),
			Name:         cloudflare.F([]string{"string", "string", "string"}),
			OS:           cloudflare.F([]radar.HTTPLocationBotClassGetParamsOS{radar.HTTPLocationBotClassGetParamsOSWindows, radar.HTTPLocationBotClassGetParamsOSMacosx, radar.HTTPLocationBotClassGetParamsOSIos}),
			TLSVersion:   cloudflare.F([]radar.HTTPLocationBotClassGetParamsTLSVersion{radar.HTTPLocationBotClassGetParamsTLSVersionTlSv1_0, radar.HTTPLocationBotClassGetParamsTLSVersionTlSv1_1, radar.HTTPLocationBotClassGetParamsTLSVersionTlSv1_2}),
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
