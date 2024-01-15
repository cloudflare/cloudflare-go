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

func TestRadarHTTPTopAseIPVersionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Tops.Ases.IPVersions.Get(
		context.TODO(),
		cloudflare.RadarHTTPTopAseIPVersionGetParamsIPVersionIPv4,
		cloudflare.RadarHTTPTopAseIPVersionGetParams{
			ASN:          cloudflare.F([]string{"string", "string", "string"}),
			BotClass:     cloudflare.F([]cloudflare.RadarHTTPTopAseIPVersionGetParamsBotClass{cloudflare.RadarHTTPTopAseIPVersionGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTopAseIPVersionGetParamsBotClassLikelyHuman}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPTopAseIPVersionGetParamsDateRange{cloudflare.RadarHTTPTopAseIPVersionGetParamsDateRange1d, cloudflare.RadarHTTPTopAseIPVersionGetParamsDateRange2d, cloudflare.RadarHTTPTopAseIPVersionGetParamsDateRange7d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTopAseIPVersionGetParamsDeviceType{cloudflare.RadarHTTPTopAseIPVersionGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPTopAseIPVersionGetParamsDeviceTypeMobile, cloudflare.RadarHTTPTopAseIPVersionGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(cloudflare.RadarHTTPTopAseIPVersionGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTopAseIPVersionGetParamsHTTPProtocol{cloudflare.RadarHTTPTopAseIPVersionGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTopAseIPVersionGetParamsHTTPProtocolHTTPs}),
			HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTopAseIPVersionGetParamsHTTPVersion{cloudflare.RadarHTTPTopAseIPVersionGetParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTopAseIPVersionGetParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTopAseIPVersionGetParamsHTTPVersionHttPv3}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"string", "string", "string"}),
			Name:         cloudflare.F([]string{"string", "string", "string"}),
			Os:           cloudflare.F([]cloudflare.RadarHTTPTopAseIPVersionGetParamsO{cloudflare.RadarHTTPTopAseIPVersionGetParamsOWindows, cloudflare.RadarHTTPTopAseIPVersionGetParamsOMacosx, cloudflare.RadarHTTPTopAseIPVersionGetParamsOIos}),
			TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPTopAseIPVersionGetParamsTlsVersion{cloudflare.RadarHTTPTopAseIPVersionGetParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTopAseIPVersionGetParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTopAseIPVersionGetParamsTlsVersionTlSv1_2}),
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
