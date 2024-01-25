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

func TestRadarHTTPTopLocationTlsVersionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Tops.Locations.TlsVersions.Get(
		context.TODO(),
		cloudflare.RadarHTTPTopLocationTlsVersionGetParamsTlsVersionTlSv1_0,
		cloudflare.RadarHTTPTopLocationTlsVersionGetParams{
			ASN:          cloudflare.F([]string{"string", "string", "string"}),
			BotClass:     cloudflare.F([]cloudflare.RadarHTTPTopLocationTlsVersionGetParamsBotClass{cloudflare.RadarHTTPTopLocationTlsVersionGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTopLocationTlsVersionGetParamsBotClassLikelyHuman}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPTopLocationTlsVersionGetParamsDateRange{cloudflare.RadarHTTPTopLocationTlsVersionGetParamsDateRange1d, cloudflare.RadarHTTPTopLocationTlsVersionGetParamsDateRange2d, cloudflare.RadarHTTPTopLocationTlsVersionGetParamsDateRange7d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTopLocationTlsVersionGetParamsDeviceType{cloudflare.RadarHTTPTopLocationTlsVersionGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPTopLocationTlsVersionGetParamsDeviceTypeMobile, cloudflare.RadarHTTPTopLocationTlsVersionGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(cloudflare.RadarHTTPTopLocationTlsVersionGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTopLocationTlsVersionGetParamsHTTPProtocol{cloudflare.RadarHTTPTopLocationTlsVersionGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTopLocationTlsVersionGetParamsHTTPProtocolHTTPs}),
			HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTopLocationTlsVersionGetParamsHTTPVersion{cloudflare.RadarHTTPTopLocationTlsVersionGetParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTopLocationTlsVersionGetParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTopLocationTlsVersionGetParamsHTTPVersionHttPv3}),
			IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTopLocationTlsVersionGetParamsIPVersion{cloudflare.RadarHTTPTopLocationTlsVersionGetParamsIPVersionIPv4, cloudflare.RadarHTTPTopLocationTlsVersionGetParamsIPVersionIPv6}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"string", "string", "string"}),
			Name:         cloudflare.F([]string{"string", "string", "string"}),
			Os:           cloudflare.F([]cloudflare.RadarHTTPTopLocationTlsVersionGetParamsO{cloudflare.RadarHTTPTopLocationTlsVersionGetParamsOWindows, cloudflare.RadarHTTPTopLocationTlsVersionGetParamsOMacosx, cloudflare.RadarHTTPTopLocationTlsVersionGetParamsOIos}),
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
