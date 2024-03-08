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

func TestRadarHTTPAseGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Ases.Get(context.TODO(), cloudflare.RadarHTTPAseGetParams{
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPAseGetParamsBotClass{cloudflare.RadarHTTPAseGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPAseGetParamsBotClassLikelyHuman}),
		Continent:    cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPAseGetParamsDateRange{cloudflare.RadarHTTPAseGetParamsDateRange1d, cloudflare.RadarHTTPAseGetParamsDateRange2d, cloudflare.RadarHTTPAseGetParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPAseGetParamsDeviceType{cloudflare.RadarHTTPAseGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPAseGetParamsDeviceTypeMobile, cloudflare.RadarHTTPAseGetParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPAseGetParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPAseGetParamsHTTPProtocol{cloudflare.RadarHTTPAseGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPAseGetParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPAseGetParamsHTTPVersion{cloudflare.RadarHTTPAseGetParamsHTTPVersionHttPv1, cloudflare.RadarHTTPAseGetParamsHTTPVersionHttPv2, cloudflare.RadarHTTPAseGetParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPAseGetParamsIPVersion{cloudflare.RadarHTTPAseGetParamsIPVersionIPv4, cloudflare.RadarHTTPAseGetParamsIPVersionIPv6}),
		Limit:        cloudflare.F(int64(5)),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		OS:           cloudflare.F([]cloudflare.RadarHTTPAseGetParamsOS{cloudflare.RadarHTTPAseGetParamsOSWindows, cloudflare.RadarHTTPAseGetParamsOSMacosx, cloudflare.RadarHTTPAseGetParamsOSIos}),
		TLSVersion:   cloudflare.F([]cloudflare.RadarHTTPAseGetParamsTLSVersion{cloudflare.RadarHTTPAseGetParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPAseGetParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPAseGetParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
