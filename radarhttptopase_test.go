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

func TestRadarHTTPTopAseListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Tops.Ases.List(context.TODO(), cloudflare.RadarHTTPTopAseListParams{
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPTopAseListParamsBotClass{cloudflare.RadarHTTPTopAseListParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTopAseListParamsBotClassLikelyHuman}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPTopAseListParamsDateRange{cloudflare.RadarHTTPTopAseListParamsDateRange1d, cloudflare.RadarHTTPTopAseListParamsDateRange2d, cloudflare.RadarHTTPTopAseListParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTopAseListParamsDeviceType{cloudflare.RadarHTTPTopAseListParamsDeviceTypeDesktop, cloudflare.RadarHTTPTopAseListParamsDeviceTypeMobile, cloudflare.RadarHTTPTopAseListParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPTopAseListParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTopAseListParamsHTTPProtocol{cloudflare.RadarHTTPTopAseListParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTopAseListParamsHTTPProtocolHTTPs}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTopAseListParamsHTTPVersion{cloudflare.RadarHTTPTopAseListParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTopAseListParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTopAseListParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTopAseListParamsIPVersion{cloudflare.RadarHTTPTopAseListParamsIPVersionIPv4, cloudflare.RadarHTTPTopAseListParamsIPVersionIPv6}),
		Limit:        cloudflare.F(int64(5)),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		Os:           cloudflare.F([]cloudflare.RadarHTTPTopAseListParamsO{cloudflare.RadarHTTPTopAseListParamsOWindows, cloudflare.RadarHTTPTopAseListParamsOMacosx, cloudflare.RadarHTTPTopAseListParamsOIos}),
		TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPTopAseListParamsTlsVersion{cloudflare.RadarHTTPTopAseListParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTopAseListParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTopAseListParamsTlsVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
