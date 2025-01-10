// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/radar"
)

func TestHTTPAseIPVersionGetWithOptionalParams(t *testing.T) {
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
			ASN:           cloudflare.F([]string{"string"}),
			BotClass:      cloudflare.F([]radar.HTTPAseIPVersionGetParamsBotClass{radar.HTTPAseIPVersionGetParamsBotClassLikelyAutomated}),
			BrowserFamily: cloudflare.F([]radar.HTTPAseIPVersionGetParamsBrowserFamily{radar.HTTPAseIPVersionGetParamsBrowserFamilyChrome}),
			Continent:     cloudflare.F([]string{"string"}),
			DateEnd:       cloudflare.F([]time.Time{time.Now()}),
			DateRange:     cloudflare.F([]string{"7d"}),
			DateStart:     cloudflare.F([]time.Time{time.Now()}),
			DeviceType:    cloudflare.F([]radar.HTTPAseIPVersionGetParamsDeviceType{radar.HTTPAseIPVersionGetParamsDeviceTypeDesktop}),
			Format:        cloudflare.F(radar.HTTPAseIPVersionGetParamsFormatJson),
			HTTPProtocol:  cloudflare.F([]radar.HTTPAseIPVersionGetParamsHTTPProtocol{radar.HTTPAseIPVersionGetParamsHTTPProtocolHTTP}),
			HTTPVersion:   cloudflare.F([]radar.HTTPAseIPVersionGetParamsHTTPVersion{radar.HTTPAseIPVersionGetParamsHTTPVersionHttPv1}),
			Limit:         cloudflare.F(int64(5)),
			Location:      cloudflare.F([]string{"string"}),
			Name:          cloudflare.F([]string{"string"}),
			OS:            cloudflare.F([]radar.HTTPAseIPVersionGetParamsOS{radar.HTTPAseIPVersionGetParamsOSWindows}),
			TLSVersion:    cloudflare.F([]radar.HTTPAseIPVersionGetParamsTLSVersion{radar.HTTPAseIPVersionGetParamsTLSVersionTlSv1_0}),
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
