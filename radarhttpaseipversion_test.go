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

func TestRadarHTTPAseIPVersionGetWithOptionalParams(t *testing.T) {
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
		cloudflare.RadarHTTPAseIPVersionGetParamsIPVersionIPv4,
		cloudflare.RadarHTTPAseIPVersionGetParams{
			ASN:          cloudflare.F([]string{"string", "string", "string"}),
			BotClass:     cloudflare.F([]cloudflare.RadarHTTPAseIPVersionGetParamsBotClass{cloudflare.RadarHTTPAseIPVersionGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPAseIPVersionGetParamsBotClassLikelyHuman}),
			Continent:    cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPAseIPVersionGetParamsDateRange{cloudflare.RadarHTTPAseIPVersionGetParamsDateRange1d, cloudflare.RadarHTTPAseIPVersionGetParamsDateRange2d, cloudflare.RadarHTTPAseIPVersionGetParamsDateRange7d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]cloudflare.RadarHTTPAseIPVersionGetParamsDeviceType{cloudflare.RadarHTTPAseIPVersionGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPAseIPVersionGetParamsDeviceTypeMobile, cloudflare.RadarHTTPAseIPVersionGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(cloudflare.RadarHTTPAseIPVersionGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPAseIPVersionGetParamsHTTPProtocol{cloudflare.RadarHTTPAseIPVersionGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPAseIPVersionGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPAseIPVersionGetParamsHTTPVersion{cloudflare.RadarHTTPAseIPVersionGetParamsHTTPVersionHttPv1, cloudflare.RadarHTTPAseIPVersionGetParamsHTTPVersionHttPv2, cloudflare.RadarHTTPAseIPVersionGetParamsHTTPVersionHttPv3}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"string", "string", "string"}),
			Name:         cloudflare.F([]string{"string", "string", "string"}),
			OS:           cloudflare.F([]cloudflare.RadarHTTPAseIPVersionGetParamsOS{cloudflare.RadarHTTPAseIPVersionGetParamsOSWindows, cloudflare.RadarHTTPAseIPVersionGetParamsOSMacosx, cloudflare.RadarHTTPAseIPVersionGetParamsOSIos}),
			TLSVersion:   cloudflare.F([]cloudflare.RadarHTTPAseIPVersionGetParamsTLSVersion{cloudflare.RadarHTTPAseIPVersionGetParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPAseIPVersionGetParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPAseIPVersionGetParamsTLSVersionTlSv1_2}),
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
