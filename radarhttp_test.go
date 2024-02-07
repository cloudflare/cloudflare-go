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

func TestRadarHTTPBotClassesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.BotClasses(context.TODO(), cloudflare.RadarHTTPBotClassesParams{
		AggInterval:  cloudflare.F(cloudflare.RadarHTTPBotClassesParamsAggInterval1h),
		Asn:          cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPBotClassesParamsDateRange{cloudflare.RadarHTTPBotClassesParamsDateRange1d, cloudflare.RadarHTTPBotClassesParamsDateRange2d, cloudflare.RadarHTTPBotClassesParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPBotClassesParamsDeviceType{cloudflare.RadarHTTPBotClassesParamsDeviceTypeDesktop, cloudflare.RadarHTTPBotClassesParamsDeviceTypeMobile, cloudflare.RadarHTTPBotClassesParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPBotClassesParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPBotClassesParamsHTTPProtocol{cloudflare.RadarHTTPBotClassesParamsHTTPProtocolHTTP, cloudflare.RadarHTTPBotClassesParamsHTTPProtocolHTTPs}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPBotClassesParamsHTTPVersion{cloudflare.RadarHTTPBotClassesParamsHTTPVersionHttPv1, cloudflare.RadarHTTPBotClassesParamsHTTPVersionHttPv2, cloudflare.RadarHTTPBotClassesParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPBotClassesParamsIPVersion{cloudflare.RadarHTTPBotClassesParamsIPVersionIPv4, cloudflare.RadarHTTPBotClassesParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		Os:           cloudflare.F([]cloudflare.RadarHTTPBotClassesParamsO{cloudflare.RadarHTTPBotClassesParamsOWindows, cloudflare.RadarHTTPBotClassesParamsOMacosx, cloudflare.RadarHTTPBotClassesParamsOIos}),
		TLSVersion:   cloudflare.F([]cloudflare.RadarHTTPBotClassesParamsTLSVersion{cloudflare.RadarHTTPBotClassesParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPBotClassesParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPBotClassesParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPBrowserFamiliesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.BrowserFamilies(context.TODO(), cloudflare.RadarHTTPBrowserFamiliesParams{
		AggInterval:  cloudflare.F(cloudflare.RadarHTTPBrowserFamiliesParamsAggInterval1h),
		Asn:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPBrowserFamiliesParamsBotClass{cloudflare.RadarHTTPBrowserFamiliesParamsBotClassLikelyAutomated, cloudflare.RadarHTTPBrowserFamiliesParamsBotClassLikelyHuman}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPBrowserFamiliesParamsDateRange{cloudflare.RadarHTTPBrowserFamiliesParamsDateRange1d, cloudflare.RadarHTTPBrowserFamiliesParamsDateRange2d, cloudflare.RadarHTTPBrowserFamiliesParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPBrowserFamiliesParamsDeviceType{cloudflare.RadarHTTPBrowserFamiliesParamsDeviceTypeDesktop, cloudflare.RadarHTTPBrowserFamiliesParamsDeviceTypeMobile, cloudflare.RadarHTTPBrowserFamiliesParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPBrowserFamiliesParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPBrowserFamiliesParamsHTTPProtocol{cloudflare.RadarHTTPBrowserFamiliesParamsHTTPProtocolHTTP, cloudflare.RadarHTTPBrowserFamiliesParamsHTTPProtocolHTTPs}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPBrowserFamiliesParamsHTTPVersion{cloudflare.RadarHTTPBrowserFamiliesParamsHTTPVersionHttPv1, cloudflare.RadarHTTPBrowserFamiliesParamsHTTPVersionHttPv2, cloudflare.RadarHTTPBrowserFamiliesParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPBrowserFamiliesParamsIPVersion{cloudflare.RadarHTTPBrowserFamiliesParamsIPVersionIPv4, cloudflare.RadarHTTPBrowserFamiliesParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		Os:           cloudflare.F([]cloudflare.RadarHTTPBrowserFamiliesParamsO{cloudflare.RadarHTTPBrowserFamiliesParamsOWindows, cloudflare.RadarHTTPBrowserFamiliesParamsOMacosx, cloudflare.RadarHTTPBrowserFamiliesParamsOIos}),
		TLSVersion:   cloudflare.F([]cloudflare.RadarHTTPBrowserFamiliesParamsTLSVersion{cloudflare.RadarHTTPBrowserFamiliesParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPBrowserFamiliesParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPBrowserFamiliesParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPBrowsersWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Browsers(context.TODO(), cloudflare.RadarHTTPBrowsersParams{
		AggInterval:   cloudflare.F(cloudflare.RadarHTTPBrowsersParamsAggInterval1h),
		Asn:           cloudflare.F([]string{"string", "string", "string"}),
		BotClass:      cloudflare.F([]cloudflare.RadarHTTPBrowsersParamsBotClass{cloudflare.RadarHTTPBrowsersParamsBotClassLikelyAutomated, cloudflare.RadarHTTPBrowsersParamsBotClassLikelyHuman}),
		DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     cloudflare.F([]cloudflare.RadarHTTPBrowsersParamsDateRange{cloudflare.RadarHTTPBrowsersParamsDateRange1d, cloudflare.RadarHTTPBrowsersParamsDateRange2d, cloudflare.RadarHTTPBrowsersParamsDateRange7d}),
		DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:    cloudflare.F([]cloudflare.RadarHTTPBrowsersParamsDeviceType{cloudflare.RadarHTTPBrowsersParamsDeviceTypeDesktop, cloudflare.RadarHTTPBrowsersParamsDeviceTypeMobile, cloudflare.RadarHTTPBrowsersParamsDeviceTypeOther}),
		Format:        cloudflare.F(cloudflare.RadarHTTPBrowsersParamsFormatJson),
		HTTPProtocol:  cloudflare.F([]cloudflare.RadarHTTPBrowsersParamsHTTPProtocol{cloudflare.RadarHTTPBrowsersParamsHTTPProtocolHTTP, cloudflare.RadarHTTPBrowsersParamsHTTPProtocolHTTPs}),
		HTTPVersion:   cloudflare.F([]cloudflare.RadarHTTPBrowsersParamsHTTPVersion{cloudflare.RadarHTTPBrowsersParamsHTTPVersionHttPv1, cloudflare.RadarHTTPBrowsersParamsHTTPVersionHttPv2, cloudflare.RadarHTTPBrowsersParamsHTTPVersionHttPv3}),
		IPVersion:     cloudflare.F([]cloudflare.RadarHTTPBrowsersParamsIPVersion{cloudflare.RadarHTTPBrowsersParamsIPVersionIPv4, cloudflare.RadarHTTPBrowsersParamsIPVersionIPv6}),
		LimitPerGroup: cloudflare.F(int64(4)),
		Location:      cloudflare.F([]string{"string", "string", "string"}),
		Name:          cloudflare.F([]string{"string", "string", "string"}),
		Os:            cloudflare.F([]cloudflare.RadarHTTPBrowsersParamsO{cloudflare.RadarHTTPBrowsersParamsOWindows, cloudflare.RadarHTTPBrowsersParamsOMacosx, cloudflare.RadarHTTPBrowsersParamsOIos}),
		TLSVersion:    cloudflare.F([]cloudflare.RadarHTTPBrowsersParamsTLSVersion{cloudflare.RadarHTTPBrowsersParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPBrowsersParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPBrowsersParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPDeviceTypesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.DeviceTypes(context.TODO(), cloudflare.RadarHTTPDeviceTypesParams{
		AggInterval:  cloudflare.F(cloudflare.RadarHTTPDeviceTypesParamsAggInterval1h),
		Asn:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPDeviceTypesParamsBotClass{cloudflare.RadarHTTPDeviceTypesParamsBotClassLikelyAutomated, cloudflare.RadarHTTPDeviceTypesParamsBotClassLikelyHuman}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPDeviceTypesParamsDateRange{cloudflare.RadarHTTPDeviceTypesParamsDateRange1d, cloudflare.RadarHTTPDeviceTypesParamsDateRange2d, cloudflare.RadarHTTPDeviceTypesParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:       cloudflare.F(cloudflare.RadarHTTPDeviceTypesParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPDeviceTypesParamsHTTPProtocol{cloudflare.RadarHTTPDeviceTypesParamsHTTPProtocolHTTP, cloudflare.RadarHTTPDeviceTypesParamsHTTPProtocolHTTPs}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPDeviceTypesParamsHTTPVersion{cloudflare.RadarHTTPDeviceTypesParamsHTTPVersionHttPv1, cloudflare.RadarHTTPDeviceTypesParamsHTTPVersionHttPv2, cloudflare.RadarHTTPDeviceTypesParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPDeviceTypesParamsIPVersion{cloudflare.RadarHTTPDeviceTypesParamsIPVersionIPv4, cloudflare.RadarHTTPDeviceTypesParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		Os:           cloudflare.F([]cloudflare.RadarHTTPDeviceTypesParamsO{cloudflare.RadarHTTPDeviceTypesParamsOWindows, cloudflare.RadarHTTPDeviceTypesParamsOMacosx, cloudflare.RadarHTTPDeviceTypesParamsOIos}),
		TLSVersion:   cloudflare.F([]cloudflare.RadarHTTPDeviceTypesParamsTLSVersion{cloudflare.RadarHTTPDeviceTypesParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPDeviceTypesParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPDeviceTypesParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPHTTPProtocolsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.HTTPProtocols(context.TODO(), cloudflare.RadarHTTPHTTPProtocolsParams{
		AggInterval: cloudflare.F(cloudflare.RadarHttphttpProtocolsParamsAggInterval1h),
		Asn:         cloudflare.F([]string{"string", "string", "string"}),
		BotClass:    cloudflare.F([]cloudflare.RadarHttphttpProtocolsParamsBotClass{cloudflare.RadarHttphttpProtocolsParamsBotClassLikelyAutomated, cloudflare.RadarHttphttpProtocolsParamsBotClassLikelyHuman}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarHttphttpProtocolsParamsDateRange{cloudflare.RadarHttphttpProtocolsParamsDateRange1d, cloudflare.RadarHttphttpProtocolsParamsDateRange2d, cloudflare.RadarHttphttpProtocolsParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:  cloudflare.F([]cloudflare.RadarHttphttpProtocolsParamsDeviceType{cloudflare.RadarHttphttpProtocolsParamsDeviceTypeDesktop, cloudflare.RadarHttphttpProtocolsParamsDeviceTypeMobile, cloudflare.RadarHttphttpProtocolsParamsDeviceTypeOther}),
		Format:      cloudflare.F(cloudflare.RadarHttphttpProtocolsParamsFormatJson),
		HTTPVersion: cloudflare.F([]cloudflare.RadarHttphttpProtocolsParamsHTTPVersion{cloudflare.RadarHttphttpProtocolsParamsHTTPVersionHttPv1, cloudflare.RadarHttphttpProtocolsParamsHTTPVersionHttPv2, cloudflare.RadarHttphttpProtocolsParamsHTTPVersionHttPv3}),
		IPVersion:   cloudflare.F([]cloudflare.RadarHttphttpProtocolsParamsIPVersion{cloudflare.RadarHttphttpProtocolsParamsIPVersionIPv4, cloudflare.RadarHttphttpProtocolsParamsIPVersionIPv6}),
		Location:    cloudflare.F([]string{"string", "string", "string"}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		Os:          cloudflare.F([]cloudflare.RadarHttphttpProtocolsParamsO{cloudflare.RadarHttphttpProtocolsParamsOWindows, cloudflare.RadarHttphttpProtocolsParamsOMacosx, cloudflare.RadarHttphttpProtocolsParamsOIos}),
		TLSVersion:  cloudflare.F([]cloudflare.RadarHttphttpProtocolsParamsTLSVersion{cloudflare.RadarHttphttpProtocolsParamsTLSVersionTlSv1_0, cloudflare.RadarHttphttpProtocolsParamsTLSVersionTlSv1_1, cloudflare.RadarHttphttpProtocolsParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPHTTPVersionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.HTTPVersions(context.TODO(), cloudflare.RadarHTTPHTTPVersionsParams{
		AggInterval:  cloudflare.F(cloudflare.RadarHttphttpVersionsParamsAggInterval1h),
		Asn:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHttphttpVersionsParamsBotClass{cloudflare.RadarHttphttpVersionsParamsBotClassLikelyAutomated, cloudflare.RadarHttphttpVersionsParamsBotClassLikelyHuman}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHttphttpVersionsParamsDateRange{cloudflare.RadarHttphttpVersionsParamsDateRange1d, cloudflare.RadarHttphttpVersionsParamsDateRange2d, cloudflare.RadarHttphttpVersionsParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHttphttpVersionsParamsDeviceType{cloudflare.RadarHttphttpVersionsParamsDeviceTypeDesktop, cloudflare.RadarHttphttpVersionsParamsDeviceTypeMobile, cloudflare.RadarHttphttpVersionsParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHttphttpVersionsParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHttphttpVersionsParamsHTTPProtocol{cloudflare.RadarHttphttpVersionsParamsHTTPProtocolHTTP, cloudflare.RadarHttphttpVersionsParamsHTTPProtocolHTTPs}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHttphttpVersionsParamsIPVersion{cloudflare.RadarHttphttpVersionsParamsIPVersionIPv4, cloudflare.RadarHttphttpVersionsParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		Os:           cloudflare.F([]cloudflare.RadarHttphttpVersionsParamsO{cloudflare.RadarHttphttpVersionsParamsOWindows, cloudflare.RadarHttphttpVersionsParamsOMacosx, cloudflare.RadarHttphttpVersionsParamsOIos}),
		TLSVersion:   cloudflare.F([]cloudflare.RadarHttphttpVersionsParamsTLSVersion{cloudflare.RadarHttphttpVersionsParamsTLSVersionTlSv1_0, cloudflare.RadarHttphttpVersionsParamsTLSVersionTlSv1_1, cloudflare.RadarHttphttpVersionsParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPIPVersionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.IPVersions(context.TODO(), cloudflare.RadarHTTPIPVersionsParams{
		AggInterval:  cloudflare.F(cloudflare.RadarHttpipVersionsParamsAggInterval1h),
		Asn:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHttpipVersionsParamsBotClass{cloudflare.RadarHttpipVersionsParamsBotClassLikelyAutomated, cloudflare.RadarHttpipVersionsParamsBotClassLikelyHuman}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHttpipVersionsParamsDateRange{cloudflare.RadarHttpipVersionsParamsDateRange1d, cloudflare.RadarHttpipVersionsParamsDateRange2d, cloudflare.RadarHttpipVersionsParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHttpipVersionsParamsDeviceType{cloudflare.RadarHttpipVersionsParamsDeviceTypeDesktop, cloudflare.RadarHttpipVersionsParamsDeviceTypeMobile, cloudflare.RadarHttpipVersionsParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHttpipVersionsParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHttpipVersionsParamsHTTPProtocol{cloudflare.RadarHttpipVersionsParamsHTTPProtocolHTTP, cloudflare.RadarHttpipVersionsParamsHTTPProtocolHTTPs}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHttpipVersionsParamsHTTPVersion{cloudflare.RadarHttpipVersionsParamsHTTPVersionHttPv1, cloudflare.RadarHttpipVersionsParamsHTTPVersionHttPv2, cloudflare.RadarHttpipVersionsParamsHTTPVersionHttPv3}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		Os:           cloudflare.F([]cloudflare.RadarHttpipVersionsParamsO{cloudflare.RadarHttpipVersionsParamsOWindows, cloudflare.RadarHttpipVersionsParamsOMacosx, cloudflare.RadarHttpipVersionsParamsOIos}),
		TLSVersion:   cloudflare.F([]cloudflare.RadarHttpipVersionsParamsTLSVersion{cloudflare.RadarHttpipVersionsParamsTLSVersionTlSv1_0, cloudflare.RadarHttpipVersionsParamsTLSVersionTlSv1_1, cloudflare.RadarHttpipVersionsParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPOssWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Oss(context.TODO(), cloudflare.RadarHTTPOssParams{
		AggInterval:  cloudflare.F(cloudflare.RadarHTTPOssParamsAggInterval1h),
		Asn:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPOssParamsBotClass{cloudflare.RadarHTTPOssParamsBotClassLikelyAutomated, cloudflare.RadarHTTPOssParamsBotClassLikelyHuman}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPOssParamsDateRange{cloudflare.RadarHTTPOssParamsDateRange1d, cloudflare.RadarHTTPOssParamsDateRange2d, cloudflare.RadarHTTPOssParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPOssParamsDeviceType{cloudflare.RadarHTTPOssParamsDeviceTypeDesktop, cloudflare.RadarHTTPOssParamsDeviceTypeMobile, cloudflare.RadarHTTPOssParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPOssParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPOssParamsHTTPProtocol{cloudflare.RadarHTTPOssParamsHTTPProtocolHTTP, cloudflare.RadarHTTPOssParamsHTTPProtocolHTTPs}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPOssParamsHTTPVersion{cloudflare.RadarHTTPOssParamsHTTPVersionHttPv1, cloudflare.RadarHTTPOssParamsHTTPVersionHttPv2, cloudflare.RadarHTTPOssParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPOssParamsIPVersion{cloudflare.RadarHTTPOssParamsIPVersionIPv4, cloudflare.RadarHTTPOssParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		TLSVersion:   cloudflare.F([]cloudflare.RadarHTTPOssParamsTLSVersion{cloudflare.RadarHTTPOssParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPOssParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPOssParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
