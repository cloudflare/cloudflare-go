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

func TestRadarHTTPTopAseTlsVersionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Tops.Ases.TlsVersions.Get(
		context.TODO(),
		cloudflare.RadarHTTPTopAseTlsVersionGetParamsTlsVersionTlSv1_0,
		cloudflare.RadarHTTPTopAseTlsVersionGetParams{
			ASN:          cloudflare.F([]string{"string", "string", "string"}),
			BotClass:     cloudflare.F([]cloudflare.RadarHTTPTopAseTlsVersionGetParamsBotClass{cloudflare.RadarHTTPTopAseTlsVersionGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTopAseTlsVersionGetParamsBotClassLikelyHuman}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPTopAseTlsVersionGetParamsDateRange{cloudflare.RadarHTTPTopAseTlsVersionGetParamsDateRange1d, cloudflare.RadarHTTPTopAseTlsVersionGetParamsDateRange2d, cloudflare.RadarHTTPTopAseTlsVersionGetParamsDateRange7d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTopAseTlsVersionGetParamsDeviceType{cloudflare.RadarHTTPTopAseTlsVersionGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPTopAseTlsVersionGetParamsDeviceTypeMobile, cloudflare.RadarHTTPTopAseTlsVersionGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(cloudflare.RadarHTTPTopAseTlsVersionGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTopAseTlsVersionGetParamsHTTPProtocol{cloudflare.RadarHTTPTopAseTlsVersionGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTopAseTlsVersionGetParamsHTTPProtocolHTTPs}),
			HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTopAseTlsVersionGetParamsHTTPVersion{cloudflare.RadarHTTPTopAseTlsVersionGetParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTopAseTlsVersionGetParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTopAseTlsVersionGetParamsHTTPVersionHttPv3}),
			IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTopAseTlsVersionGetParamsIPVersion{cloudflare.RadarHTTPTopAseTlsVersionGetParamsIPVersionIPv4, cloudflare.RadarHTTPTopAseTlsVersionGetParamsIPVersionIPv6}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"string", "string", "string"}),
			Name:         cloudflare.F([]string{"string", "string", "string"}),
			Os:           cloudflare.F([]cloudflare.RadarHTTPTopAseTlsVersionGetParamsO{cloudflare.RadarHTTPTopAseTlsVersionGetParamsOWindows, cloudflare.RadarHTTPTopAseTlsVersionGetParamsOMacosx, cloudflare.RadarHTTPTopAseTlsVersionGetParamsOIos}),
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
