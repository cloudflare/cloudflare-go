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

func TestRadarHTTPTopAseHTTPProtocolGetWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Radars.HTTPs.Tops.Ases.HTTPProtocols.Get(
		context.TODO(),
		cloudflare.RadarHTTPTopAseHTTPProtocolGetParamsHTTPProtocolHTTP,
		cloudflare.RadarHTTPTopAseHTTPProtocolGetParams{
			ASN:          cloudflare.F([]string{"15169", "15169", "15169"}),
			BotClass:     cloudflare.F([]cloudflare.RadarHTTPTopAseHTTPProtocolGetParamsBotClass{cloudflare.RadarHTTPTopAseHTTPProtocolGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTopAseHTTPProtocolGetParamsBotClassLikelyHuman}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPTopAseHTTPProtocolGetParamsDateRange{cloudflare.RadarHTTPTopAseHTTPProtocolGetParamsDateRange1d, cloudflare.RadarHTTPTopAseHTTPProtocolGetParamsDateRange7d, cloudflare.RadarHTTPTopAseHTTPProtocolGetParamsDateRange14d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTopAseHTTPProtocolGetParamsDeviceType{cloudflare.RadarHTTPTopAseHTTPProtocolGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPTopAseHTTPProtocolGetParamsDeviceTypeMobile, cloudflare.RadarHTTPTopAseHTTPProtocolGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(cloudflare.RadarHTTPTopAseHTTPProtocolGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTopAseHTTPProtocolGetParamsHTTPProtocol{cloudflare.RadarHTTPTopAseHTTPProtocolGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTopAseHTTPProtocolGetParamsHTTPProtocolHTTPs}),
			IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTopAseHTTPProtocolGetParamsIPVersion{cloudflare.RadarHTTPTopAseHTTPProtocolGetParamsIPVersionIPv4, cloudflare.RadarHTTPTopAseHTTPProtocolGetParamsIPVersionIPv6}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
			Name:         cloudflare.F([]string{"main_series", "main_series", "main_series"}),
			Os:           cloudflare.F([]cloudflare.RadarHTTPTopAseHTTPProtocolGetParamsO{cloudflare.RadarHTTPTopAseHTTPProtocolGetParamsOWindows, cloudflare.RadarHTTPTopAseHTTPProtocolGetParamsOMacosx, cloudflare.RadarHTTPTopAseHTTPProtocolGetParamsOAndroid}),
			TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPTopAseHTTPProtocolGetParamsTlsVersion{cloudflare.RadarHTTPTopAseHTTPProtocolGetParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTopAseHTTPProtocolGetParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTopAseHTTPProtocolGetParamsTlsVersionTlSv1_2}),
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
