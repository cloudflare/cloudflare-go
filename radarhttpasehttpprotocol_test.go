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

func TestRadarHTTPAseHTTPProtocolGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Ases.HTTPProtocol.Get(
		context.TODO(),
		cloudflare.RadarHTTPAseHTTPProtocolGetParamsHTTPProtocolHTTP,
		cloudflare.RadarHTTPAseHTTPProtocolGetParams{
			ASN:        cloudflare.F([]string{"string", "string", "string"}),
			BotClass:   cloudflare.F([]cloudflare.RadarHTTPAseHTTPProtocolGetParamsBotClass{cloudflare.RadarHTTPAseHTTPProtocolGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPAseHTTPProtocolGetParamsBotClassLikelyHuman}),
			Continent:  cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:  cloudflare.F([]cloudflare.RadarHTTPAseHTTPProtocolGetParamsDateRange{cloudflare.RadarHTTPAseHTTPProtocolGetParamsDateRange1d, cloudflare.RadarHTTPAseHTTPProtocolGetParamsDateRange2d, cloudflare.RadarHTTPAseHTTPProtocolGetParamsDateRange7d}),
			DateStart:  cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType: cloudflare.F([]cloudflare.RadarHTTPAseHTTPProtocolGetParamsDeviceType{cloudflare.RadarHTTPAseHTTPProtocolGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPAseHTTPProtocolGetParamsDeviceTypeMobile, cloudflare.RadarHTTPAseHTTPProtocolGetParamsDeviceTypeOther}),
			Format:     cloudflare.F(cloudflare.RadarHTTPAseHTTPProtocolGetParamsFormatJson),
			IPVersion:  cloudflare.F([]cloudflare.RadarHTTPAseHTTPProtocolGetParamsIPVersion{cloudflare.RadarHTTPAseHTTPProtocolGetParamsIPVersionIPv4, cloudflare.RadarHTTPAseHTTPProtocolGetParamsIPVersionIPv6}),
			Limit:      cloudflare.F(int64(5)),
			Location:   cloudflare.F([]string{"string", "string", "string"}),
			Name:       cloudflare.F([]string{"string", "string", "string"}),
			OS:         cloudflare.F([]cloudflare.RadarHTTPAseHTTPProtocolGetParamsOS{cloudflare.RadarHTTPAseHTTPProtocolGetParamsOSWindows, cloudflare.RadarHTTPAseHTTPProtocolGetParamsOSMacosx, cloudflare.RadarHTTPAseHTTPProtocolGetParamsOSIos}),
			TLSVersion: cloudflare.F([]cloudflare.RadarHTTPAseHTTPProtocolGetParamsTLSVersion{cloudflare.RadarHTTPAseHTTPProtocolGetParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPAseHTTPProtocolGetParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPAseHTTPProtocolGetParamsTLSVersionTlSv1_2}),
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
