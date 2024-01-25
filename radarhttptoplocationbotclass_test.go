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

func TestRadarHTTPTopLocationBotClassGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Tops.Locations.BotClasses.Get(
		context.TODO(),
		cloudflare.RadarHTTPTopLocationBotClassGetParamsBotClassLikelyAutomated,
		cloudflare.RadarHTTPTopLocationBotClassGetParams{
			ASN:          cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPTopLocationBotClassGetParamsDateRange{cloudflare.RadarHTTPTopLocationBotClassGetParamsDateRange1d, cloudflare.RadarHTTPTopLocationBotClassGetParamsDateRange2d, cloudflare.RadarHTTPTopLocationBotClassGetParamsDateRange7d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTopLocationBotClassGetParamsDeviceType{cloudflare.RadarHTTPTopLocationBotClassGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPTopLocationBotClassGetParamsDeviceTypeMobile, cloudflare.RadarHTTPTopLocationBotClassGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(cloudflare.RadarHTTPTopLocationBotClassGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTopLocationBotClassGetParamsHTTPProtocol{cloudflare.RadarHTTPTopLocationBotClassGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTopLocationBotClassGetParamsHTTPProtocolHTTPs}),
			HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTopLocationBotClassGetParamsHTTPVersion{cloudflare.RadarHTTPTopLocationBotClassGetParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTopLocationBotClassGetParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTopLocationBotClassGetParamsHTTPVersionHttPv3}),
			IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTopLocationBotClassGetParamsIPVersion{cloudflare.RadarHTTPTopLocationBotClassGetParamsIPVersionIPv4, cloudflare.RadarHTTPTopLocationBotClassGetParamsIPVersionIPv6}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"string", "string", "string"}),
			Name:         cloudflare.F([]string{"string", "string", "string"}),
			Os:           cloudflare.F([]cloudflare.RadarHTTPTopLocationBotClassGetParamsO{cloudflare.RadarHTTPTopLocationBotClassGetParamsOWindows, cloudflare.RadarHTTPTopLocationBotClassGetParamsOMacosx, cloudflare.RadarHTTPTopLocationBotClassGetParamsOIos}),
			TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPTopLocationBotClassGetParamsTlsVersion{cloudflare.RadarHTTPTopLocationBotClassGetParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTopLocationBotClassGetParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTopLocationBotClassGetParamsTlsVersionTlSv1_2}),
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
