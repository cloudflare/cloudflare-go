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

func TestRadarHTTPTopLocationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Tops.Locations.List(context.TODO(), cloudflare.RadarHTTPTopLocationListParams{
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPTopLocationListParamsBotClass{cloudflare.RadarHTTPTopLocationListParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTopLocationListParamsBotClassLikelyHuman}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPTopLocationListParamsDateRange{cloudflare.RadarHTTPTopLocationListParamsDateRange1d, cloudflare.RadarHTTPTopLocationListParamsDateRange2d, cloudflare.RadarHTTPTopLocationListParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTopLocationListParamsDeviceType{cloudflare.RadarHTTPTopLocationListParamsDeviceTypeDesktop, cloudflare.RadarHTTPTopLocationListParamsDeviceTypeMobile, cloudflare.RadarHTTPTopLocationListParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPTopLocationListParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTopLocationListParamsHTTPProtocol{cloudflare.RadarHTTPTopLocationListParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTopLocationListParamsHTTPProtocolHTTPs}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTopLocationListParamsHTTPVersion{cloudflare.RadarHTTPTopLocationListParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTopLocationListParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTopLocationListParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTopLocationListParamsIPVersion{cloudflare.RadarHTTPTopLocationListParamsIPVersionIPv4, cloudflare.RadarHTTPTopLocationListParamsIPVersionIPv6}),
		Limit:        cloudflare.F(int64(5)),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		Os:           cloudflare.F([]cloudflare.RadarHTTPTopLocationListParamsO{cloudflare.RadarHTTPTopLocationListParamsOWindows, cloudflare.RadarHTTPTopLocationListParamsOMacosx, cloudflare.RadarHTTPTopLocationListParamsOIos}),
		TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPTopLocationListParamsTlsVersion{cloudflare.RadarHTTPTopLocationListParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTopLocationListParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTopLocationListParamsTlsVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
