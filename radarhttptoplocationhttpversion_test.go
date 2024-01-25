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

func TestRadarHTTPTopLocationHTTPVersionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Tops.Locations.HTTPVersions.Get(
		context.TODO(),
		cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsHTTPVersionHttPv1,
		cloudflare.RadarHTTPTopLocationHTTPVersionGetParams{
			ASN:          cloudflare.F([]string{"string", "string", "string"}),
			BotClass:     cloudflare.F([]cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsBotClass{cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsBotClassLikelyHuman}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsDateRange{cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsDateRange1d, cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsDateRange2d, cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsDateRange7d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsDeviceType{cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsDeviceTypeMobile, cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsHTTPProtocol{cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsHTTPProtocolHTTPs}),
			IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsIPVersion{cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsIPVersionIPv4, cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsIPVersionIPv6}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"string", "string", "string"}),
			Name:         cloudflare.F([]string{"string", "string", "string"}),
			Os:           cloudflare.F([]cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsO{cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsOWindows, cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsOMacosx, cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsOIos}),
			TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsTlsVersion{cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTopLocationHTTPVersionGetParamsTlsVersionTlSv1_2}),
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
