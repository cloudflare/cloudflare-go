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

func TestRadarHTTPTimeseriesGroupBotClassWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Radar.HTTP.TimeseriesGroups.BotClass(context.TODO(), cloudflare.RadarHTTPTimeseriesGroupBotClassParams{
		AggInterval:  cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupBotClassParamsAggInterval1h),
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		Continent:    cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupBotClassParamsDateRange{cloudflare.RadarHTTPTimeseriesGroupBotClassParamsDateRange1d, cloudflare.RadarHTTPTimeseriesGroupBotClassParamsDateRange2d, cloudflare.RadarHTTPTimeseriesGroupBotClassParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupBotClassParamsDeviceType{cloudflare.RadarHTTPTimeseriesGroupBotClassParamsDeviceTypeDesktop, cloudflare.RadarHTTPTimeseriesGroupBotClassParamsDeviceTypeMobile, cloudflare.RadarHTTPTimeseriesGroupBotClassParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupBotClassParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupBotClassParamsHTTPProtocol{cloudflare.RadarHTTPTimeseriesGroupBotClassParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTimeseriesGroupBotClassParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupBotClassParamsHTTPVersion{cloudflare.RadarHTTPTimeseriesGroupBotClassParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTimeseriesGroupBotClassParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTimeseriesGroupBotClassParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupBotClassParamsIPVersion{cloudflare.RadarHTTPTimeseriesGroupBotClassParamsIPVersionIPv4, cloudflare.RadarHTTPTimeseriesGroupBotClassParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		OS:           cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupBotClassParamsOS{cloudflare.RadarHTTPTimeseriesGroupBotClassParamsOSWindows, cloudflare.RadarHTTPTimeseriesGroupBotClassParamsOSMacosx, cloudflare.RadarHTTPTimeseriesGroupBotClassParamsOSIos}),
		TLSVersion:   cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupBotClassParamsTLSVersion{cloudflare.RadarHTTPTimeseriesGroupBotClassParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPTimeseriesGroupBotClassParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPTimeseriesGroupBotClassParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPTimeseriesGroupBrowserWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Radar.HTTP.TimeseriesGroups.Browser(context.TODO(), cloudflare.RadarHTTPTimeseriesGroupBrowserParams{
		AggInterval:   cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupBrowserParamsAggInterval1h),
		ASN:           cloudflare.F([]string{"string", "string", "string"}),
		BotClass:      cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupBrowserParamsBotClass{cloudflare.RadarHTTPTimeseriesGroupBrowserParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTimeseriesGroupBrowserParamsBotClassLikelyHuman}),
		Continent:     cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupBrowserParamsDateRange{cloudflare.RadarHTTPTimeseriesGroupBrowserParamsDateRange1d, cloudflare.RadarHTTPTimeseriesGroupBrowserParamsDateRange2d, cloudflare.RadarHTTPTimeseriesGroupBrowserParamsDateRange7d}),
		DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:    cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupBrowserParamsDeviceType{cloudflare.RadarHTTPTimeseriesGroupBrowserParamsDeviceTypeDesktop, cloudflare.RadarHTTPTimeseriesGroupBrowserParamsDeviceTypeMobile, cloudflare.RadarHTTPTimeseriesGroupBrowserParamsDeviceTypeOther}),
		Format:        cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupBrowserParamsFormatJson),
		HTTPProtocol:  cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupBrowserParamsHTTPProtocol{cloudflare.RadarHTTPTimeseriesGroupBrowserParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTimeseriesGroupBrowserParamsHTTPProtocolHTTPS}),
		HTTPVersion:   cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupBrowserParamsHTTPVersion{cloudflare.RadarHTTPTimeseriesGroupBrowserParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTimeseriesGroupBrowserParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTimeseriesGroupBrowserParamsHTTPVersionHttPv3}),
		IPVersion:     cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupBrowserParamsIPVersion{cloudflare.RadarHTTPTimeseriesGroupBrowserParamsIPVersionIPv4, cloudflare.RadarHTTPTimeseriesGroupBrowserParamsIPVersionIPv6}),
		LimitPerGroup: cloudflare.F(int64(4)),
		Location:      cloudflare.F([]string{"string", "string", "string"}),
		Name:          cloudflare.F([]string{"string", "string", "string"}),
		OS:            cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupBrowserParamsOS{cloudflare.RadarHTTPTimeseriesGroupBrowserParamsOSWindows, cloudflare.RadarHTTPTimeseriesGroupBrowserParamsOSMacosx, cloudflare.RadarHTTPTimeseriesGroupBrowserParamsOSIos}),
		TLSVersion:    cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupBrowserParamsTLSVersion{cloudflare.RadarHTTPTimeseriesGroupBrowserParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPTimeseriesGroupBrowserParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPTimeseriesGroupBrowserParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPTimeseriesGroupBrowserFamilyWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Radar.HTTP.TimeseriesGroups.BrowserFamily(context.TODO(), cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParams{
		AggInterval:  cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParamsAggInterval1h),
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParamsBotClass{cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParamsBotClassLikelyHuman}),
		Continent:    cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange{cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange1d, cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange2d, cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParamsDeviceType{cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParamsDeviceTypeDesktop, cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParamsDeviceTypeMobile, cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocol{cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParamsHTTPVersion{cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParamsIPVersion{cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParamsIPVersionIPv4, cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		OS:           cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParamsOS{cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParamsOSWindows, cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParamsOSMacosx, cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParamsOSIos}),
		TLSVersion:   cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParamsTLSVersion{cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPTimeseriesGroupDeviceTypeWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Radar.HTTP.TimeseriesGroups.DeviceType(context.TODO(), cloudflare.RadarHTTPTimeseriesGroupDeviceTypeParams{
		AggInterval:  cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupDeviceTypeParamsAggInterval1h),
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupDeviceTypeParamsBotClass{cloudflare.RadarHTTPTimeseriesGroupDeviceTypeParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTimeseriesGroupDeviceTypeParamsBotClassLikelyHuman}),
		Continent:    cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange{cloudflare.RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange1d, cloudflare.RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange2d, cloudflare.RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:       cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupDeviceTypeParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupDeviceTypeParamsHTTPProtocol{cloudflare.RadarHTTPTimeseriesGroupDeviceTypeParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTimeseriesGroupDeviceTypeParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupDeviceTypeParamsHTTPVersion{cloudflare.RadarHTTPTimeseriesGroupDeviceTypeParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTimeseriesGroupDeviceTypeParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTimeseriesGroupDeviceTypeParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupDeviceTypeParamsIPVersion{cloudflare.RadarHTTPTimeseriesGroupDeviceTypeParamsIPVersionIPv4, cloudflare.RadarHTTPTimeseriesGroupDeviceTypeParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		OS:           cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupDeviceTypeParamsOS{cloudflare.RadarHTTPTimeseriesGroupDeviceTypeParamsOSWindows, cloudflare.RadarHTTPTimeseriesGroupDeviceTypeParamsOSMacosx, cloudflare.RadarHTTPTimeseriesGroupDeviceTypeParamsOSIos}),
		TLSVersion:   cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupDeviceTypeParamsTLSVersion{cloudflare.RadarHTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPTimeseriesGroupHTTPProtocolWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Radar.HTTP.TimeseriesGroups.HTTPProtocol(context.TODO(), cloudflare.RadarHTTPTimeseriesGroupHTTPProtocolParams{
		AggInterval: cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupHTTPProtocolParamsAggInterval1h),
		ASN:         cloudflare.F([]string{"string", "string", "string"}),
		BotClass:    cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupHTTPProtocolParamsBotClass{cloudflare.RadarHTTPTimeseriesGroupHTTPProtocolParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTimeseriesGroupHTTPProtocolParamsBotClassLikelyHuman}),
		Continent:   cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange{cloudflare.RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange1d, cloudflare.RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange2d, cloudflare.RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:  cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupHTTPProtocolParamsDeviceType{cloudflare.RadarHTTPTimeseriesGroupHTTPProtocolParamsDeviceTypeDesktop, cloudflare.RadarHTTPTimeseriesGroupHTTPProtocolParamsDeviceTypeMobile, cloudflare.RadarHTTPTimeseriesGroupHTTPProtocolParamsDeviceTypeOther}),
		Format:      cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupHTTPProtocolParamsFormatJson),
		HTTPVersion: cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupHTTPProtocolParamsHTTPVersion{cloudflare.RadarHTTPTimeseriesGroupHTTPProtocolParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTimeseriesGroupHTTPProtocolParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTimeseriesGroupHTTPProtocolParamsHTTPVersionHttPv3}),
		IPVersion:   cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupHTTPProtocolParamsIPVersion{cloudflare.RadarHTTPTimeseriesGroupHTTPProtocolParamsIPVersionIPv4, cloudflare.RadarHTTPTimeseriesGroupHTTPProtocolParamsIPVersionIPv6}),
		Location:    cloudflare.F([]string{"string", "string", "string"}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		OS:          cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupHTTPProtocolParamsOS{cloudflare.RadarHTTPTimeseriesGroupHTTPProtocolParamsOSWindows, cloudflare.RadarHTTPTimeseriesGroupHTTPProtocolParamsOSMacosx, cloudflare.RadarHTTPTimeseriesGroupHTTPProtocolParamsOSIos}),
		TLSVersion:  cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupHTTPProtocolParamsTLSVersion{cloudflare.RadarHTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPTimeseriesGroupHTTPVersionWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Radar.HTTP.TimeseriesGroups.HTTPVersion(context.TODO(), cloudflare.RadarHTTPTimeseriesGroupHTTPVersionParams{
		AggInterval:  cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupHTTPVersionParamsAggInterval1h),
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupHTTPVersionParamsBotClass{cloudflare.RadarHTTPTimeseriesGroupHTTPVersionParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTimeseriesGroupHTTPVersionParamsBotClassLikelyHuman}),
		Continent:    cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange{cloudflare.RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange1d, cloudflare.RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange2d, cloudflare.RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupHTTPVersionParamsDeviceType{cloudflare.RadarHTTPTimeseriesGroupHTTPVersionParamsDeviceTypeDesktop, cloudflare.RadarHTTPTimeseriesGroupHTTPVersionParamsDeviceTypeMobile, cloudflare.RadarHTTPTimeseriesGroupHTTPVersionParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupHTTPVersionParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupHTTPVersionParamsHTTPProtocol{cloudflare.RadarHTTPTimeseriesGroupHTTPVersionParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTimeseriesGroupHTTPVersionParamsHTTPProtocolHTTPS}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupHTTPVersionParamsIPVersion{cloudflare.RadarHTTPTimeseriesGroupHTTPVersionParamsIPVersionIPv4, cloudflare.RadarHTTPTimeseriesGroupHTTPVersionParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		OS:           cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupHTTPVersionParamsOS{cloudflare.RadarHTTPTimeseriesGroupHTTPVersionParamsOSWindows, cloudflare.RadarHTTPTimeseriesGroupHTTPVersionParamsOSMacosx, cloudflare.RadarHTTPTimeseriesGroupHTTPVersionParamsOSIos}),
		TLSVersion:   cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupHTTPVersionParamsTLSVersion{cloudflare.RadarHTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPTimeseriesGroupIPVersionWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Radar.HTTP.TimeseriesGroups.IPVersion(context.TODO(), cloudflare.RadarHTTPTimeseriesGroupIPVersionParams{
		AggInterval:  cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupIPVersionParamsAggInterval1h),
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupIPVersionParamsBotClass{cloudflare.RadarHTTPTimeseriesGroupIPVersionParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTimeseriesGroupIPVersionParamsBotClassLikelyHuman}),
		Continent:    cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupIPVersionParamsDateRange{cloudflare.RadarHTTPTimeseriesGroupIPVersionParamsDateRange1d, cloudflare.RadarHTTPTimeseriesGroupIPVersionParamsDateRange2d, cloudflare.RadarHTTPTimeseriesGroupIPVersionParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupIPVersionParamsDeviceType{cloudflare.RadarHTTPTimeseriesGroupIPVersionParamsDeviceTypeDesktop, cloudflare.RadarHTTPTimeseriesGroupIPVersionParamsDeviceTypeMobile, cloudflare.RadarHTTPTimeseriesGroupIPVersionParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupIPVersionParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupIPVersionParamsHTTPProtocol{cloudflare.RadarHTTPTimeseriesGroupIPVersionParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTimeseriesGroupIPVersionParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupIPVersionParamsHTTPVersion{cloudflare.RadarHTTPTimeseriesGroupIPVersionParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTimeseriesGroupIPVersionParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTimeseriesGroupIPVersionParamsHTTPVersionHttPv3}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		OS:           cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupIPVersionParamsOS{cloudflare.RadarHTTPTimeseriesGroupIPVersionParamsOSWindows, cloudflare.RadarHTTPTimeseriesGroupIPVersionParamsOSMacosx, cloudflare.RadarHTTPTimeseriesGroupIPVersionParamsOSIos}),
		TLSVersion:   cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupIPVersionParamsTLSVersion{cloudflare.RadarHTTPTimeseriesGroupIPVersionParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPTimeseriesGroupIPVersionParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPTimeseriesGroupIPVersionParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPTimeseriesGroupOSWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Radar.HTTP.TimeseriesGroups.OS(context.TODO(), cloudflare.RadarHTTPTimeseriesGroupOSParams{
		AggInterval:  cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupOSParamsAggInterval1h),
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupOSParamsBotClass{cloudflare.RadarHTTPTimeseriesGroupOSParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTimeseriesGroupOSParamsBotClassLikelyHuman}),
		Continent:    cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupOSParamsDateRange{cloudflare.RadarHTTPTimeseriesGroupOSParamsDateRange1d, cloudflare.RadarHTTPTimeseriesGroupOSParamsDateRange2d, cloudflare.RadarHTTPTimeseriesGroupOSParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupOSParamsDeviceType{cloudflare.RadarHTTPTimeseriesGroupOSParamsDeviceTypeDesktop, cloudflare.RadarHTTPTimeseriesGroupOSParamsDeviceTypeMobile, cloudflare.RadarHTTPTimeseriesGroupOSParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupOSParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupOSParamsHTTPProtocol{cloudflare.RadarHTTPTimeseriesGroupOSParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTimeseriesGroupOSParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupOSParamsHTTPVersion{cloudflare.RadarHTTPTimeseriesGroupOSParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTimeseriesGroupOSParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTimeseriesGroupOSParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupOSParamsIPVersion{cloudflare.RadarHTTPTimeseriesGroupOSParamsIPVersionIPv4, cloudflare.RadarHTTPTimeseriesGroupOSParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		TLSVersion:   cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupOSParamsTLSVersion{cloudflare.RadarHTTPTimeseriesGroupOSParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPTimeseriesGroupOSParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPTimeseriesGroupOSParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPTimeseriesGroupTLSVersionWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Radar.HTTP.TimeseriesGroups.TLSVersion(context.TODO(), cloudflare.RadarHTTPTimeseriesGroupTLSVersionParams{
		AggInterval:  cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupTLSVersionParamsAggInterval1h),
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupTLSVersionParamsBotClass{cloudflare.RadarHTTPTimeseriesGroupTLSVersionParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTimeseriesGroupTLSVersionParamsBotClassLikelyHuman}),
		Continent:    cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupTLSVersionParamsDateRange{cloudflare.RadarHTTPTimeseriesGroupTLSVersionParamsDateRange1d, cloudflare.RadarHTTPTimeseriesGroupTLSVersionParamsDateRange2d, cloudflare.RadarHTTPTimeseriesGroupTLSVersionParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupTLSVersionParamsDeviceType{cloudflare.RadarHTTPTimeseriesGroupTLSVersionParamsDeviceTypeDesktop, cloudflare.RadarHTTPTimeseriesGroupTLSVersionParamsDeviceTypeMobile, cloudflare.RadarHTTPTimeseriesGroupTLSVersionParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupTLSVersionParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupTLSVersionParamsHTTPProtocol{cloudflare.RadarHTTPTimeseriesGroupTLSVersionParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTimeseriesGroupTLSVersionParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupTLSVersionParamsHTTPVersion{cloudflare.RadarHTTPTimeseriesGroupTLSVersionParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTimeseriesGroupTLSVersionParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTimeseriesGroupTLSVersionParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupTLSVersionParamsIPVersion{cloudflare.RadarHTTPTimeseriesGroupTLSVersionParamsIPVersionIPv4, cloudflare.RadarHTTPTimeseriesGroupTLSVersionParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		OS:           cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupTLSVersionParamsOS{cloudflare.RadarHTTPTimeseriesGroupTLSVersionParamsOSWindows, cloudflare.RadarHTTPTimeseriesGroupTLSVersionParamsOSMacosx, cloudflare.RadarHTTPTimeseriesGroupTLSVersionParamsOSIos}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
