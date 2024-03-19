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

func TestHTTPAseIPVersionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Ases.IPVersion.Get(
		context.TODO(),
		radar.HTTPAseIPVersionGetParamsIPVersionIPv4,
		radar.HTTPAseIPVersionGetParams{
			ASN:          cloudflare.F([]string{"string", "string", "string"}),
			BotClass:     cloudflare.F([]radar.HTTPAseIPVersionGetParamsBotClass{radar.HTTPAseIPVersionGetParamsBotClassLikelyAutomated, radar.HTTPAseIPVersionGetParamsBotClassLikelyHuman}),
			Continent:    cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]radar.HTTPAseIPVersionGetParamsDateRange{radar.HTTPAseIPVersionGetParamsDateRange1d, radar.HTTPAseIPVersionGetParamsDateRange2d, radar.HTTPAseIPVersionGetParamsDateRange7d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]radar.HTTPAseIPVersionGetParamsDeviceType{radar.HTTPAseIPVersionGetParamsDeviceTypeDesktop, radar.HTTPAseIPVersionGetParamsDeviceTypeMobile, radar.HTTPAseIPVersionGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(radar.HTTPAseIPVersionGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]radar.HTTPAseIPVersionGetParamsHTTPProtocol{radar.HTTPAseIPVersionGetParamsHTTPProtocolHTTP, radar.HTTPAseIPVersionGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:  cloudflare.F([]radar.HTTPAseIPVersionGetParamsHTTPVersion{radar.HTTPAseIPVersionGetParamsHTTPVersionHttPv1, radar.HTTPAseIPVersionGetParamsHTTPVersionHttPv2, radar.HTTPAseIPVersionGetParamsHTTPVersionHttPv3}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"string", "string", "string"}),
			Name:         cloudflare.F([]string{"string", "string", "string"}),
			OS:           cloudflare.F([]radar.HTTPAseIPVersionGetParamsOS{radar.HTTPAseIPVersionGetParamsOSWindows, radar.HTTPAseIPVersionGetParamsOSMacosx, radar.HTTPAseIPVersionGetParamsOSIos}),
			TLSVersion:   cloudflare.F([]radar.HTTPAseIPVersionGetParamsTLSVersion{radar.HTTPAseIPVersionGetParamsTLSVersionTlSv1_0, radar.HTTPAseIPVersionGetParamsTLSVersionTlSv1_1, radar.HTTPAseIPVersionGetParamsTLSVersionTlSv1_2}),
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
