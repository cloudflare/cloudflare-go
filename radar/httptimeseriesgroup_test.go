// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/radar"
)

func TestHTTPTimeseriesGroupBotClassWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.BotClass(context.TODO(), radar.HTTPTimeseriesGroupBotClassParams{
		AggInterval:  cloudflare.F(radar.HTTPTimeseriesGroupBotClassParamsAggInterval1h),
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		Continent:    cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]radar.HTTPTimeseriesGroupBotClassParamsDeviceType{radar.HTTPTimeseriesGroupBotClassParamsDeviceTypeDesktop, radar.HTTPTimeseriesGroupBotClassParamsDeviceTypeMobile, radar.HTTPTimeseriesGroupBotClassParamsDeviceTypeOther}),
		Format:       cloudflare.F(radar.HTTPTimeseriesGroupBotClassParamsFormatJson),
		HTTPProtocol: cloudflare.F([]radar.HTTPTimeseriesGroupBotClassParamsHTTPProtocol{radar.HTTPTimeseriesGroupBotClassParamsHTTPProtocolHTTP, radar.HTTPTimeseriesGroupBotClassParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]radar.HTTPTimeseriesGroupBotClassParamsHTTPVersion{radar.HTTPTimeseriesGroupBotClassParamsHTTPVersionHttPv1, radar.HTTPTimeseriesGroupBotClassParamsHTTPVersionHttPv2, radar.HTTPTimeseriesGroupBotClassParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]radar.HTTPTimeseriesGroupBotClassParamsIPVersion{radar.HTTPTimeseriesGroupBotClassParamsIPVersionIPv4, radar.HTTPTimeseriesGroupBotClassParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		OS:           cloudflare.F([]radar.HTTPTimeseriesGroupBotClassParamsOS{radar.HTTPTimeseriesGroupBotClassParamsOSWindows, radar.HTTPTimeseriesGroupBotClassParamsOSMacosx, radar.HTTPTimeseriesGroupBotClassParamsOSIos}),
		TLSVersion:   cloudflare.F([]radar.HTTPTimeseriesGroupBotClassParamsTLSVersion{radar.HTTPTimeseriesGroupBotClassParamsTLSVersionTlSv1_0, radar.HTTPTimeseriesGroupBotClassParamsTLSVersionTlSv1_1, radar.HTTPTimeseriesGroupBotClassParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPTimeseriesGroupBrowserWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.Browser(context.TODO(), radar.HTTPTimeseriesGroupBrowserParams{
		AggInterval:   cloudflare.F(radar.HTTPTimeseriesGroupBrowserParamsAggInterval1h),
		ASN:           cloudflare.F([]string{"string", "string", "string"}),
		BotClass:      cloudflare.F([]radar.HTTPTimeseriesGroupBrowserParamsBotClass{radar.HTTPTimeseriesGroupBrowserParamsBotClassLikelyAutomated, radar.HTTPTimeseriesGroupBrowserParamsBotClassLikelyHuman}),
		Continent:     cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:    cloudflare.F([]radar.HTTPTimeseriesGroupBrowserParamsDeviceType{radar.HTTPTimeseriesGroupBrowserParamsDeviceTypeDesktop, radar.HTTPTimeseriesGroupBrowserParamsDeviceTypeMobile, radar.HTTPTimeseriesGroupBrowserParamsDeviceTypeOther}),
		Format:        cloudflare.F(radar.HTTPTimeseriesGroupBrowserParamsFormatJson),
		HTTPProtocol:  cloudflare.F([]radar.HTTPTimeseriesGroupBrowserParamsHTTPProtocol{radar.HTTPTimeseriesGroupBrowserParamsHTTPProtocolHTTP, radar.HTTPTimeseriesGroupBrowserParamsHTTPProtocolHTTPS}),
		HTTPVersion:   cloudflare.F([]radar.HTTPTimeseriesGroupBrowserParamsHTTPVersion{radar.HTTPTimeseriesGroupBrowserParamsHTTPVersionHttPv1, radar.HTTPTimeseriesGroupBrowserParamsHTTPVersionHttPv2, radar.HTTPTimeseriesGroupBrowserParamsHTTPVersionHttPv3}),
		IPVersion:     cloudflare.F([]radar.HTTPTimeseriesGroupBrowserParamsIPVersion{radar.HTTPTimeseriesGroupBrowserParamsIPVersionIPv4, radar.HTTPTimeseriesGroupBrowserParamsIPVersionIPv6}),
		LimitPerGroup: cloudflare.F(int64(4)),
		Location:      cloudflare.F([]string{"string", "string", "string"}),
		Name:          cloudflare.F([]string{"string", "string", "string"}),
		OS:            cloudflare.F([]radar.HTTPTimeseriesGroupBrowserParamsOS{radar.HTTPTimeseriesGroupBrowserParamsOSWindows, radar.HTTPTimeseriesGroupBrowserParamsOSMacosx, radar.HTTPTimeseriesGroupBrowserParamsOSIos}),
		TLSVersion:    cloudflare.F([]radar.HTTPTimeseriesGroupBrowserParamsTLSVersion{radar.HTTPTimeseriesGroupBrowserParamsTLSVersionTlSv1_0, radar.HTTPTimeseriesGroupBrowserParamsTLSVersionTlSv1_1, radar.HTTPTimeseriesGroupBrowserParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPTimeseriesGroupBrowserFamilyWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.BrowserFamily(context.TODO(), radar.HTTPTimeseriesGroupBrowserFamilyParams{
		AggInterval:  cloudflare.F(radar.HTTPTimeseriesGroupBrowserFamilyParamsAggInterval1h),
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]radar.HTTPTimeseriesGroupBrowserFamilyParamsBotClass{radar.HTTPTimeseriesGroupBrowserFamilyParamsBotClassLikelyAutomated, radar.HTTPTimeseriesGroupBrowserFamilyParamsBotClassLikelyHuman}),
		Continent:    cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]radar.HTTPTimeseriesGroupBrowserFamilyParamsDeviceType{radar.HTTPTimeseriesGroupBrowserFamilyParamsDeviceTypeDesktop, radar.HTTPTimeseriesGroupBrowserFamilyParamsDeviceTypeMobile, radar.HTTPTimeseriesGroupBrowserFamilyParamsDeviceTypeOther}),
		Format:       cloudflare.F(radar.HTTPTimeseriesGroupBrowserFamilyParamsFormatJson),
		HTTPProtocol: cloudflare.F([]radar.HTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocol{radar.HTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocolHTTP, radar.HTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]radar.HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersion{radar.HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersionHttPv1, radar.HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersionHttPv2, radar.HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]radar.HTTPTimeseriesGroupBrowserFamilyParamsIPVersion{radar.HTTPTimeseriesGroupBrowserFamilyParamsIPVersionIPv4, radar.HTTPTimeseriesGroupBrowserFamilyParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		OS:           cloudflare.F([]radar.HTTPTimeseriesGroupBrowserFamilyParamsOS{radar.HTTPTimeseriesGroupBrowserFamilyParamsOSWindows, radar.HTTPTimeseriesGroupBrowserFamilyParamsOSMacosx, radar.HTTPTimeseriesGroupBrowserFamilyParamsOSIos}),
		TLSVersion:   cloudflare.F([]radar.HTTPTimeseriesGroupBrowserFamilyParamsTLSVersion{radar.HTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSv1_0, radar.HTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSv1_1, radar.HTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPTimeseriesGroupDeviceTypeWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.DeviceType(context.TODO(), radar.HTTPTimeseriesGroupDeviceTypeParams{
		AggInterval:  cloudflare.F(radar.HTTPTimeseriesGroupDeviceTypeParamsAggInterval1h),
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]radar.HTTPTimeseriesGroupDeviceTypeParamsBotClass{radar.HTTPTimeseriesGroupDeviceTypeParamsBotClassLikelyAutomated, radar.HTTPTimeseriesGroupDeviceTypeParamsBotClassLikelyHuman}),
		Continent:    cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:       cloudflare.F(radar.HTTPTimeseriesGroupDeviceTypeParamsFormatJson),
		HTTPProtocol: cloudflare.F([]radar.HTTPTimeseriesGroupDeviceTypeParamsHTTPProtocol{radar.HTTPTimeseriesGroupDeviceTypeParamsHTTPProtocolHTTP, radar.HTTPTimeseriesGroupDeviceTypeParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]radar.HTTPTimeseriesGroupDeviceTypeParamsHTTPVersion{radar.HTTPTimeseriesGroupDeviceTypeParamsHTTPVersionHttPv1, radar.HTTPTimeseriesGroupDeviceTypeParamsHTTPVersionHttPv2, radar.HTTPTimeseriesGroupDeviceTypeParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]radar.HTTPTimeseriesGroupDeviceTypeParamsIPVersion{radar.HTTPTimeseriesGroupDeviceTypeParamsIPVersionIPv4, radar.HTTPTimeseriesGroupDeviceTypeParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		OS:           cloudflare.F([]radar.HTTPTimeseriesGroupDeviceTypeParamsOS{radar.HTTPTimeseriesGroupDeviceTypeParamsOSWindows, radar.HTTPTimeseriesGroupDeviceTypeParamsOSMacosx, radar.HTTPTimeseriesGroupDeviceTypeParamsOSIos}),
		TLSVersion:   cloudflare.F([]radar.HTTPTimeseriesGroupDeviceTypeParamsTLSVersion{radar.HTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSv1_0, radar.HTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSv1_1, radar.HTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPTimeseriesGroupHTTPProtocolWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.HTTPProtocol(context.TODO(), radar.HTTPTimeseriesGroupHTTPProtocolParams{
		AggInterval: cloudflare.F(radar.HTTPTimeseriesGroupHTTPProtocolParamsAggInterval1h),
		ASN:         cloudflare.F([]string{"string", "string", "string"}),
		BotClass:    cloudflare.F([]radar.HTTPTimeseriesGroupHTTPProtocolParamsBotClass{radar.HTTPTimeseriesGroupHTTPProtocolParamsBotClassLikelyAutomated, radar.HTTPTimeseriesGroupHTTPProtocolParamsBotClassLikelyHuman}),
		Continent:   cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:  cloudflare.F([]radar.HTTPTimeseriesGroupHTTPProtocolParamsDeviceType{radar.HTTPTimeseriesGroupHTTPProtocolParamsDeviceTypeDesktop, radar.HTTPTimeseriesGroupHTTPProtocolParamsDeviceTypeMobile, radar.HTTPTimeseriesGroupHTTPProtocolParamsDeviceTypeOther}),
		Format:      cloudflare.F(radar.HTTPTimeseriesGroupHTTPProtocolParamsFormatJson),
		HTTPVersion: cloudflare.F([]radar.HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersion{radar.HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersionHttPv1, radar.HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersionHttPv2, radar.HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersionHttPv3}),
		IPVersion:   cloudflare.F([]radar.HTTPTimeseriesGroupHTTPProtocolParamsIPVersion{radar.HTTPTimeseriesGroupHTTPProtocolParamsIPVersionIPv4, radar.HTTPTimeseriesGroupHTTPProtocolParamsIPVersionIPv6}),
		Location:    cloudflare.F([]string{"string", "string", "string"}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		OS:          cloudflare.F([]radar.HTTPTimeseriesGroupHTTPProtocolParamsOS{radar.HTTPTimeseriesGroupHTTPProtocolParamsOSWindows, radar.HTTPTimeseriesGroupHTTPProtocolParamsOSMacosx, radar.HTTPTimeseriesGroupHTTPProtocolParamsOSIos}),
		TLSVersion:  cloudflare.F([]radar.HTTPTimeseriesGroupHTTPProtocolParamsTLSVersion{radar.HTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSv1_0, radar.HTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSv1_1, radar.HTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPTimeseriesGroupHTTPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.HTTPVersion(context.TODO(), radar.HTTPTimeseriesGroupHTTPVersionParams{
		AggInterval:  cloudflare.F(radar.HTTPTimeseriesGroupHTTPVersionParamsAggInterval1h),
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]radar.HTTPTimeseriesGroupHTTPVersionParamsBotClass{radar.HTTPTimeseriesGroupHTTPVersionParamsBotClassLikelyAutomated, radar.HTTPTimeseriesGroupHTTPVersionParamsBotClassLikelyHuman}),
		Continent:    cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]radar.HTTPTimeseriesGroupHTTPVersionParamsDeviceType{radar.HTTPTimeseriesGroupHTTPVersionParamsDeviceTypeDesktop, radar.HTTPTimeseriesGroupHTTPVersionParamsDeviceTypeMobile, radar.HTTPTimeseriesGroupHTTPVersionParamsDeviceTypeOther}),
		Format:       cloudflare.F(radar.HTTPTimeseriesGroupHTTPVersionParamsFormatJson),
		HTTPProtocol: cloudflare.F([]radar.HTTPTimeseriesGroupHTTPVersionParamsHTTPProtocol{radar.HTTPTimeseriesGroupHTTPVersionParamsHTTPProtocolHTTP, radar.HTTPTimeseriesGroupHTTPVersionParamsHTTPProtocolHTTPS}),
		IPVersion:    cloudflare.F([]radar.HTTPTimeseriesGroupHTTPVersionParamsIPVersion{radar.HTTPTimeseriesGroupHTTPVersionParamsIPVersionIPv4, radar.HTTPTimeseriesGroupHTTPVersionParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		OS:           cloudflare.F([]radar.HTTPTimeseriesGroupHTTPVersionParamsOS{radar.HTTPTimeseriesGroupHTTPVersionParamsOSWindows, radar.HTTPTimeseriesGroupHTTPVersionParamsOSMacosx, radar.HTTPTimeseriesGroupHTTPVersionParamsOSIos}),
		TLSVersion:   cloudflare.F([]radar.HTTPTimeseriesGroupHTTPVersionParamsTLSVersion{radar.HTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSv1_0, radar.HTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSv1_1, radar.HTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPTimeseriesGroupIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.IPVersion(context.TODO(), radar.HTTPTimeseriesGroupIPVersionParams{
		AggInterval:  cloudflare.F(radar.HTTPTimeseriesGroupIPVersionParamsAggInterval1h),
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]radar.HTTPTimeseriesGroupIPVersionParamsBotClass{radar.HTTPTimeseriesGroupIPVersionParamsBotClassLikelyAutomated, radar.HTTPTimeseriesGroupIPVersionParamsBotClassLikelyHuman}),
		Continent:    cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]radar.HTTPTimeseriesGroupIPVersionParamsDeviceType{radar.HTTPTimeseriesGroupIPVersionParamsDeviceTypeDesktop, radar.HTTPTimeseriesGroupIPVersionParamsDeviceTypeMobile, radar.HTTPTimeseriesGroupIPVersionParamsDeviceTypeOther}),
		Format:       cloudflare.F(radar.HTTPTimeseriesGroupIPVersionParamsFormatJson),
		HTTPProtocol: cloudflare.F([]radar.HTTPTimeseriesGroupIPVersionParamsHTTPProtocol{radar.HTTPTimeseriesGroupIPVersionParamsHTTPProtocolHTTP, radar.HTTPTimeseriesGroupIPVersionParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]radar.HTTPTimeseriesGroupIPVersionParamsHTTPVersion{radar.HTTPTimeseriesGroupIPVersionParamsHTTPVersionHttPv1, radar.HTTPTimeseriesGroupIPVersionParamsHTTPVersionHttPv2, radar.HTTPTimeseriesGroupIPVersionParamsHTTPVersionHttPv3}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		OS:           cloudflare.F([]radar.HTTPTimeseriesGroupIPVersionParamsOS{radar.HTTPTimeseriesGroupIPVersionParamsOSWindows, radar.HTTPTimeseriesGroupIPVersionParamsOSMacosx, radar.HTTPTimeseriesGroupIPVersionParamsOSIos}),
		TLSVersion:   cloudflare.F([]radar.HTTPTimeseriesGroupIPVersionParamsTLSVersion{radar.HTTPTimeseriesGroupIPVersionParamsTLSVersionTlSv1_0, radar.HTTPTimeseriesGroupIPVersionParamsTLSVersionTlSv1_1, radar.HTTPTimeseriesGroupIPVersionParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPTimeseriesGroupOSWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.OS(context.TODO(), radar.HTTPTimeseriesGroupOSParams{
		AggInterval:  cloudflare.F(radar.HTTPTimeseriesGroupOSParamsAggInterval1h),
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]radar.HTTPTimeseriesGroupOSParamsBotClass{radar.HTTPTimeseriesGroupOSParamsBotClassLikelyAutomated, radar.HTTPTimeseriesGroupOSParamsBotClassLikelyHuman}),
		Continent:    cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]radar.HTTPTimeseriesGroupOSParamsDeviceType{radar.HTTPTimeseriesGroupOSParamsDeviceTypeDesktop, radar.HTTPTimeseriesGroupOSParamsDeviceTypeMobile, radar.HTTPTimeseriesGroupOSParamsDeviceTypeOther}),
		Format:       cloudflare.F(radar.HTTPTimeseriesGroupOSParamsFormatJson),
		HTTPProtocol: cloudflare.F([]radar.HTTPTimeseriesGroupOSParamsHTTPProtocol{radar.HTTPTimeseriesGroupOSParamsHTTPProtocolHTTP, radar.HTTPTimeseriesGroupOSParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]radar.HTTPTimeseriesGroupOSParamsHTTPVersion{radar.HTTPTimeseriesGroupOSParamsHTTPVersionHttPv1, radar.HTTPTimeseriesGroupOSParamsHTTPVersionHttPv2, radar.HTTPTimeseriesGroupOSParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]radar.HTTPTimeseriesGroupOSParamsIPVersion{radar.HTTPTimeseriesGroupOSParamsIPVersionIPv4, radar.HTTPTimeseriesGroupOSParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		TLSVersion:   cloudflare.F([]radar.HTTPTimeseriesGroupOSParamsTLSVersion{radar.HTTPTimeseriesGroupOSParamsTLSVersionTlSv1_0, radar.HTTPTimeseriesGroupOSParamsTLSVersionTlSv1_1, radar.HTTPTimeseriesGroupOSParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPTimeseriesGroupPostQuantumWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.PostQuantum(context.TODO(), radar.HTTPTimeseriesGroupPostQuantumParams{
		AggInterval:  cloudflare.F(radar.HTTPTimeseriesGroupPostQuantumParamsAggInterval1h),
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]radar.HTTPTimeseriesGroupPostQuantumParamsBotClass{radar.HTTPTimeseriesGroupPostQuantumParamsBotClassLikelyAutomated, radar.HTTPTimeseriesGroupPostQuantumParamsBotClassLikelyHuman}),
		Continent:    cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]radar.HTTPTimeseriesGroupPostQuantumParamsDeviceType{radar.HTTPTimeseriesGroupPostQuantumParamsDeviceTypeDesktop, radar.HTTPTimeseriesGroupPostQuantumParamsDeviceTypeMobile, radar.HTTPTimeseriesGroupPostQuantumParamsDeviceTypeOther}),
		Format:       cloudflare.F(radar.HTTPTimeseriesGroupPostQuantumParamsFormatJson),
		HTTPProtocol: cloudflare.F([]radar.HTTPTimeseriesGroupPostQuantumParamsHTTPProtocol{radar.HTTPTimeseriesGroupPostQuantumParamsHTTPProtocolHTTP, radar.HTTPTimeseriesGroupPostQuantumParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]radar.HTTPTimeseriesGroupPostQuantumParamsHTTPVersion{radar.HTTPTimeseriesGroupPostQuantumParamsHTTPVersionHttPv1, radar.HTTPTimeseriesGroupPostQuantumParamsHTTPVersionHttPv2, radar.HTTPTimeseriesGroupPostQuantumParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]radar.HTTPTimeseriesGroupPostQuantumParamsIPVersion{radar.HTTPTimeseriesGroupPostQuantumParamsIPVersionIPv4, radar.HTTPTimeseriesGroupPostQuantumParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		OS:           cloudflare.F([]radar.HTTPTimeseriesGroupPostQuantumParamsOS{radar.HTTPTimeseriesGroupPostQuantumParamsOSWindows, radar.HTTPTimeseriesGroupPostQuantumParamsOSMacosx, radar.HTTPTimeseriesGroupPostQuantumParamsOSIos}),
		TLSVersion:   cloudflare.F([]radar.HTTPTimeseriesGroupPostQuantumParamsTLSVersion{radar.HTTPTimeseriesGroupPostQuantumParamsTLSVersionTlSv1_0, radar.HTTPTimeseriesGroupPostQuantumParamsTLSVersionTlSv1_1, radar.HTTPTimeseriesGroupPostQuantumParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPTimeseriesGroupTLSVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.TLSVersion(context.TODO(), radar.HTTPTimeseriesGroupTLSVersionParams{
		AggInterval:  cloudflare.F(radar.HTTPTimeseriesGroupTLSVersionParamsAggInterval1h),
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]radar.HTTPTimeseriesGroupTLSVersionParamsBotClass{radar.HTTPTimeseriesGroupTLSVersionParamsBotClassLikelyAutomated, radar.HTTPTimeseriesGroupTLSVersionParamsBotClassLikelyHuman}),
		Continent:    cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]radar.HTTPTimeseriesGroupTLSVersionParamsDeviceType{radar.HTTPTimeseriesGroupTLSVersionParamsDeviceTypeDesktop, radar.HTTPTimeseriesGroupTLSVersionParamsDeviceTypeMobile, radar.HTTPTimeseriesGroupTLSVersionParamsDeviceTypeOther}),
		Format:       cloudflare.F(radar.HTTPTimeseriesGroupTLSVersionParamsFormatJson),
		HTTPProtocol: cloudflare.F([]radar.HTTPTimeseriesGroupTLSVersionParamsHTTPProtocol{radar.HTTPTimeseriesGroupTLSVersionParamsHTTPProtocolHTTP, radar.HTTPTimeseriesGroupTLSVersionParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]radar.HTTPTimeseriesGroupTLSVersionParamsHTTPVersion{radar.HTTPTimeseriesGroupTLSVersionParamsHTTPVersionHttPv1, radar.HTTPTimeseriesGroupTLSVersionParamsHTTPVersionHttPv2, radar.HTTPTimeseriesGroupTLSVersionParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]radar.HTTPTimeseriesGroupTLSVersionParamsIPVersion{radar.HTTPTimeseriesGroupTLSVersionParamsIPVersionIPv4, radar.HTTPTimeseriesGroupTLSVersionParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		OS:           cloudflare.F([]radar.HTTPTimeseriesGroupTLSVersionParamsOS{radar.HTTPTimeseriesGroupTLSVersionParamsOSWindows, radar.HTTPTimeseriesGroupTLSVersionParamsOSMacosx, radar.HTTPTimeseriesGroupTLSVersionParamsOSIos}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
