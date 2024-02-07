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

func TestRadarHTTPTLSVersionListWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Radar.HTTP.TLSVersion.List(context.TODO(), cloudflare.RadarHTTPTLSVersionListParams{
		AggInterval:  cloudflare.F(cloudflare.RadarHttptlsVersionListParamsAggInterval1h),
		Asn:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHttptlsVersionListParamsBotClass{cloudflare.RadarHttptlsVersionListParamsBotClassLikelyAutomated, cloudflare.RadarHttptlsVersionListParamsBotClassLikelyHuman}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHttptlsVersionListParamsDateRange{cloudflare.RadarHttptlsVersionListParamsDateRange1d, cloudflare.RadarHttptlsVersionListParamsDateRange2d, cloudflare.RadarHttptlsVersionListParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHttptlsVersionListParamsDeviceType{cloudflare.RadarHttptlsVersionListParamsDeviceTypeDesktop, cloudflare.RadarHttptlsVersionListParamsDeviceTypeMobile, cloudflare.RadarHttptlsVersionListParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHttptlsVersionListParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHttptlsVersionListParamsHTTPProtocol{cloudflare.RadarHttptlsVersionListParamsHTTPProtocolHTTP, cloudflare.RadarHttptlsVersionListParamsHTTPProtocolHTTPs}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHttptlsVersionListParamsHTTPVersion{cloudflare.RadarHttptlsVersionListParamsHTTPVersionHttPv1, cloudflare.RadarHttptlsVersionListParamsHTTPVersionHttPv2, cloudflare.RadarHttptlsVersionListParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHttptlsVersionListParamsIPVersion{cloudflare.RadarHttptlsVersionListParamsIPVersionIPv4, cloudflare.RadarHttptlsVersionListParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		Os:           cloudflare.F([]cloudflare.RadarHttptlsVersionListParamsO{cloudflare.RadarHttptlsVersionListParamsOWindows, cloudflare.RadarHttptlsVersionListParamsOMacosx, cloudflare.RadarHttptlsVersionListParamsOIos}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
