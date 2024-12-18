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

func TestHTTPAseGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Ases.Get(context.TODO(), radar.HTTPAseGetParams{
		ASN:           cloudflare.F([]string{"string"}),
		BotClass:      cloudflare.F([]radar.HTTPAseGetParamsBotClass{radar.HTTPAseGetParamsBotClassLikelyAutomated}),
		BrowserFamily: cloudflare.F([]radar.HTTPAseGetParamsBrowserFamily{radar.HTTPAseGetParamsBrowserFamilyChrome}),
		Continent:     cloudflare.F([]string{"string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now()}),
		DateRange:     cloudflare.F([]string{"7d"}),
		DateStart:     cloudflare.F([]time.Time{time.Now()}),
		DeviceType:    cloudflare.F([]radar.HTTPAseGetParamsDeviceType{radar.HTTPAseGetParamsDeviceTypeDesktop}),
		Format:        cloudflare.F(radar.HTTPAseGetParamsFormatJson),
		HTTPProtocol:  cloudflare.F([]radar.HTTPAseGetParamsHTTPProtocol{radar.HTTPAseGetParamsHTTPProtocolHTTP}),
		HTTPVersion:   cloudflare.F([]radar.HTTPAseGetParamsHTTPVersion{radar.HTTPAseGetParamsHTTPVersionHttPv1}),
		IPVersion:     cloudflare.F([]radar.HTTPAseGetParamsIPVersion{radar.HTTPAseGetParamsIPVersionIPv4}),
		Limit:         cloudflare.F(int64(5)),
		Location:      cloudflare.F([]string{"string"}),
		Name:          cloudflare.F([]string{"string"}),
		OS:            cloudflare.F([]radar.HTTPAseGetParamsOS{radar.HTTPAseGetParamsOSWindows}),
		TLSVersion:    cloudflare.F([]radar.HTTPAseGetParamsTLSVersion{radar.HTTPAseGetParamsTLSVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
