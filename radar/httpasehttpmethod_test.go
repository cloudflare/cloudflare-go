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

func TestHTTPAseHTTPMethodGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Ases.HTTPMethod.Get(
		context.TODO(),
		radar.HTTPAseHTTPMethodGetParamsHTTPVersionHttPv1,
		radar.HTTPAseHTTPMethodGetParams{
			ASN:           cloudflare.F([]string{"string"}),
			BotClass:      cloudflare.F([]radar.HTTPAseHTTPMethodGetParamsBotClass{radar.HTTPAseHTTPMethodGetParamsBotClassLikelyAutomated}),
			BrowserFamily: cloudflare.F([]radar.HTTPAseHTTPMethodGetParamsBrowserFamily{radar.HTTPAseHTTPMethodGetParamsBrowserFamilyChrome}),
			Continent:     cloudflare.F([]string{"string"}),
			DateEnd:       cloudflare.F([]time.Time{time.Now()}),
			DateRange:     cloudflare.F([]string{"7d"}),
			DateStart:     cloudflare.F([]time.Time{time.Now()}),
			DeviceType:    cloudflare.F([]radar.HTTPAseHTTPMethodGetParamsDeviceType{radar.HTTPAseHTTPMethodGetParamsDeviceTypeDesktop}),
			Format:        cloudflare.F(radar.HTTPAseHTTPMethodGetParamsFormatJson),
			HTTPProtocol:  cloudflare.F([]radar.HTTPAseHTTPMethodGetParamsHTTPProtocol{radar.HTTPAseHTTPMethodGetParamsHTTPProtocolHTTP}),
			IPVersion:     cloudflare.F([]radar.HTTPAseHTTPMethodGetParamsIPVersion{radar.HTTPAseHTTPMethodGetParamsIPVersionIPv4}),
			Limit:         cloudflare.F(int64(5)),
			Location:      cloudflare.F([]string{"string"}),
			Name:          cloudflare.F([]string{"string"}),
			OS:            cloudflare.F([]radar.HTTPAseHTTPMethodGetParamsOS{radar.HTTPAseHTTPMethodGetParamsOSWindows}),
			TLSVersion:    cloudflare.F([]radar.HTTPAseHTTPMethodGetParamsTLSVersion{radar.HTTPAseHTTPMethodGetParamsTLSVersionTlSv1_0}),
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
