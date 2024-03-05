// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestRadarHTTPAseHTTPMethodGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Ases.HTTPMethod.Get(
		context.TODO(),
		cloudflare.RadarHTTPAseHTTPMethodGetParamsHTTPVersionHttPv1,
		cloudflare.RadarHTTPAseHTTPMethodGetParams{
			ASN:          cloudflare.F([]string{"string", "string", "string"}),
			BotClass:     cloudflare.F([]cloudflare.RadarHTTPAseHTTPMethodGetParamsBotClass{cloudflare.RadarHTTPAseHTTPMethodGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPAseHTTPMethodGetParamsBotClassLikelyHuman}),
			Continent:    cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPAseHTTPMethodGetParamsDateRange{cloudflare.RadarHTTPAseHTTPMethodGetParamsDateRange1d, cloudflare.RadarHTTPAseHTTPMethodGetParamsDateRange2d, cloudflare.RadarHTTPAseHTTPMethodGetParamsDateRange7d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]cloudflare.RadarHTTPAseHTTPMethodGetParamsDeviceType{cloudflare.RadarHTTPAseHTTPMethodGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPAseHTTPMethodGetParamsDeviceTypeMobile, cloudflare.RadarHTTPAseHTTPMethodGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(cloudflare.RadarHTTPAseHTTPMethodGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPAseHTTPMethodGetParamsHTTPProtocol{cloudflare.RadarHTTPAseHTTPMethodGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPAseHTTPMethodGetParamsHTTPProtocolHTTPS}),
			IPVersion:    cloudflare.F([]cloudflare.RadarHTTPAseHTTPMethodGetParamsIPVersion{cloudflare.RadarHTTPAseHTTPMethodGetParamsIPVersionIPv4, cloudflare.RadarHTTPAseHTTPMethodGetParamsIPVersionIPv6}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"string", "string", "string"}),
			Name:         cloudflare.F([]string{"string", "string", "string"}),
			OS:           cloudflare.F([]cloudflare.RadarHTTPAseHTTPMethodGetParamsOS{cloudflare.RadarHTTPAseHTTPMethodGetParamsOSWindows, cloudflare.RadarHTTPAseHTTPMethodGetParamsOSMacosx, cloudflare.RadarHTTPAseHTTPMethodGetParamsOSIos}),
			TLSVersion:   cloudflare.F([]cloudflare.RadarHTTPAseHTTPMethodGetParamsTLSVersion{cloudflare.RadarHTTPAseHTTPMethodGetParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPAseHTTPMethodGetParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPAseHTTPMethodGetParamsTLSVersionTlSv1_2}),
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
