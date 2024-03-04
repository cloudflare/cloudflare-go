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

func TestRadarHTTPLocationIPVersionGetWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Radar.HTTP.Locations.IPVersion.Get(
		context.TODO(),
		cloudflare.RadarHTTPLocationIPVersionGetParamsIPVersionIPv4,
		cloudflare.RadarHTTPLocationIPVersionGetParams{
			ASN:          cloudflare.F([]string{"string", "string", "string"}),
			BotClass:     cloudflare.F([]cloudflare.RadarHTTPLocationIPVersionGetParamsBotClass{cloudflare.RadarHTTPLocationIPVersionGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPLocationIPVersionGetParamsBotClassLikelyHuman}),
			Continent:    cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPLocationIPVersionGetParamsDateRange{cloudflare.RadarHTTPLocationIPVersionGetParamsDateRange1d, cloudflare.RadarHTTPLocationIPVersionGetParamsDateRange2d, cloudflare.RadarHTTPLocationIPVersionGetParamsDateRange7d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]cloudflare.RadarHTTPLocationIPVersionGetParamsDeviceType{cloudflare.RadarHTTPLocationIPVersionGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPLocationIPVersionGetParamsDeviceTypeMobile, cloudflare.RadarHTTPLocationIPVersionGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(cloudflare.RadarHTTPLocationIPVersionGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPLocationIPVersionGetParamsHTTPProtocol{cloudflare.RadarHTTPLocationIPVersionGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPLocationIPVersionGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPLocationIPVersionGetParamsHTTPVersion{cloudflare.RadarHTTPLocationIPVersionGetParamsHTTPVersionHttPv1, cloudflare.RadarHTTPLocationIPVersionGetParamsHTTPVersionHttPv2, cloudflare.RadarHTTPLocationIPVersionGetParamsHTTPVersionHttPv3}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"string", "string", "string"}),
			Name:         cloudflare.F([]string{"string", "string", "string"}),
			OS:           cloudflare.F([]cloudflare.RadarHTTPLocationIPVersionGetParamsOS{cloudflare.RadarHTTPLocationIPVersionGetParamsOSWindows, cloudflare.RadarHTTPLocationIPVersionGetParamsOSMacosx, cloudflare.RadarHTTPLocationIPVersionGetParamsOSIos}),
			TLSVersion:   cloudflare.F([]cloudflare.RadarHTTPLocationIPVersionGetParamsTLSVersion{cloudflare.RadarHTTPLocationIPVersionGetParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPLocationIPVersionGetParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPLocationIPVersionGetParamsTLSVersionTlSv1_2}),
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
