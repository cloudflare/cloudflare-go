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

func TestRadarHTTPTopLocationHTTPProtocolGetWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Radar.HTTP.Tops.Locations.HTTPProtocols.Get(
		context.TODO(),
		cloudflare.RadarHTTPTopLocationHTTPProtocolGetParamsHTTPProtocolHTTP,
		cloudflare.RadarHTTPTopLocationHTTPProtocolGetParams{
			ASN:          cloudflare.F([]string{"string", "string", "string"}),
			BotClass:     cloudflare.F([]cloudflare.RadarHTTPTopLocationHTTPProtocolGetParamsBotClass{cloudflare.RadarHTTPTopLocationHTTPProtocolGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTopLocationHTTPProtocolGetParamsBotClassLikelyHuman}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPTopLocationHTTPProtocolGetParamsDateRange{cloudflare.RadarHTTPTopLocationHTTPProtocolGetParamsDateRange1d, cloudflare.RadarHTTPTopLocationHTTPProtocolGetParamsDateRange2d, cloudflare.RadarHTTPTopLocationHTTPProtocolGetParamsDateRange7d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTopLocationHTTPProtocolGetParamsDeviceType{cloudflare.RadarHTTPTopLocationHTTPProtocolGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPTopLocationHTTPProtocolGetParamsDeviceTypeMobile, cloudflare.RadarHTTPTopLocationHTTPProtocolGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(cloudflare.RadarHTTPTopLocationHTTPProtocolGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTopLocationHTTPProtocolGetParamsHTTPProtocol{cloudflare.RadarHTTPTopLocationHTTPProtocolGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTopLocationHTTPProtocolGetParamsHTTPProtocolHTTPs}),
			IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTopLocationHTTPProtocolGetParamsIPVersion{cloudflare.RadarHTTPTopLocationHTTPProtocolGetParamsIPVersionIPv4, cloudflare.RadarHTTPTopLocationHTTPProtocolGetParamsIPVersionIPv6}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"string", "string", "string"}),
			Name:         cloudflare.F([]string{"string", "string", "string"}),
			Os:           cloudflare.F([]cloudflare.RadarHTTPTopLocationHTTPProtocolGetParamsO{cloudflare.RadarHTTPTopLocationHTTPProtocolGetParamsOWindows, cloudflare.RadarHTTPTopLocationHTTPProtocolGetParamsOMacosx, cloudflare.RadarHTTPTopLocationHTTPProtocolGetParamsOIos}),
			TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPTopLocationHTTPProtocolGetParamsTlsVersion{cloudflare.RadarHTTPTopLocationHTTPProtocolGetParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTopLocationHTTPProtocolGetParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTopLocationHTTPProtocolGetParamsTlsVersionTlSv1_2}),
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
