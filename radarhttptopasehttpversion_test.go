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

func TestRadarHTTPTopAseHTTPVersionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Tops.Ases.HTTPVersions.Get(
		context.TODO(),
		cloudflare.RadarHTTPTopAseHTTPVersionGetParamsHTTPVersionHttPv1,
		cloudflare.RadarHTTPTopAseHTTPVersionGetParams{
			ASN:          cloudflare.F([]string{"string", "string", "string"}),
			BotClass:     cloudflare.F([]cloudflare.RadarHTTPTopAseHTTPVersionGetParamsBotClass{cloudflare.RadarHTTPTopAseHTTPVersionGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTopAseHTTPVersionGetParamsBotClassLikelyHuman}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPTopAseHTTPVersionGetParamsDateRange{cloudflare.RadarHTTPTopAseHTTPVersionGetParamsDateRange1d, cloudflare.RadarHTTPTopAseHTTPVersionGetParamsDateRange2d, cloudflare.RadarHTTPTopAseHTTPVersionGetParamsDateRange7d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTopAseHTTPVersionGetParamsDeviceType{cloudflare.RadarHTTPTopAseHTTPVersionGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPTopAseHTTPVersionGetParamsDeviceTypeMobile, cloudflare.RadarHTTPTopAseHTTPVersionGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(cloudflare.RadarHTTPTopAseHTTPVersionGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTopAseHTTPVersionGetParamsHTTPProtocol{cloudflare.RadarHTTPTopAseHTTPVersionGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTopAseHTTPVersionGetParamsHTTPProtocolHTTPs}),
			IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTopAseHTTPVersionGetParamsIPVersion{cloudflare.RadarHTTPTopAseHTTPVersionGetParamsIPVersionIPv4, cloudflare.RadarHTTPTopAseHTTPVersionGetParamsIPVersionIPv6}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"string", "string", "string"}),
			Name:         cloudflare.F([]string{"string", "string", "string"}),
			Os:           cloudflare.F([]cloudflare.RadarHTTPTopAseHTTPVersionGetParamsO{cloudflare.RadarHTTPTopAseHTTPVersionGetParamsOWindows, cloudflare.RadarHTTPTopAseHTTPVersionGetParamsOMacosx, cloudflare.RadarHTTPTopAseHTTPVersionGetParamsOIos}),
			TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPTopAseHTTPVersionGetParamsTlsVersion{cloudflare.RadarHTTPTopAseHTTPVersionGetParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTopAseHTTPVersionGetParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTopAseHTTPVersionGetParamsTlsVersionTlSv1_2}),
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
