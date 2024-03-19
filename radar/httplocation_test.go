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

func TestHTTPLocationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Locations.Get(context.TODO(), radar.HTTPLocationGetParams{
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]radar.HTTPLocationGetParamsBotClass{radar.HTTPLocationGetParamsBotClassLikelyAutomated, radar.HTTPLocationGetParamsBotClassLikelyHuman}),
		Continent:    cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]radar.HTTPLocationGetParamsDateRange{radar.HTTPLocationGetParamsDateRange1d, radar.HTTPLocationGetParamsDateRange2d, radar.HTTPLocationGetParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]radar.HTTPLocationGetParamsDeviceType{radar.HTTPLocationGetParamsDeviceTypeDesktop, radar.HTTPLocationGetParamsDeviceTypeMobile, radar.HTTPLocationGetParamsDeviceTypeOther}),
		Format:       cloudflare.F(radar.HTTPLocationGetParamsFormatJson),
		HTTPProtocol: cloudflare.F([]radar.HTTPLocationGetParamsHTTPProtocol{radar.HTTPLocationGetParamsHTTPProtocolHTTP, radar.HTTPLocationGetParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]radar.HTTPLocationGetParamsHTTPVersion{radar.HTTPLocationGetParamsHTTPVersionHttPv1, radar.HTTPLocationGetParamsHTTPVersionHttPv2, radar.HTTPLocationGetParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]radar.HTTPLocationGetParamsIPVersion{radar.HTTPLocationGetParamsIPVersionIPv4, radar.HTTPLocationGetParamsIPVersionIPv6}),
		Limit:        cloudflare.F(int64(5)),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		OS:           cloudflare.F([]radar.HTTPLocationGetParamsOS{radar.HTTPLocationGetParamsOSWindows, radar.HTTPLocationGetParamsOSMacosx, radar.HTTPLocationGetParamsOSIos}),
		TLSVersion:   cloudflare.F([]radar.HTTPLocationGetParamsTLSVersion{radar.HTTPLocationGetParamsTLSVersionTlSv1_0, radar.HTTPLocationGetParamsTLSVersionTlSv1_1, radar.HTTPLocationGetParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
