// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
)

func TestRadarHTTPSummaryBotClassWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.BotClass(context.TODO(), cloudflare.RadarHTTPSummaryBotClassParams{
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		Continent:    cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPSummaryBotClassParamsDateRange{cloudflare.RadarHTTPSummaryBotClassParamsDateRange1d, cloudflare.RadarHTTPSummaryBotClassParamsDateRange2d, cloudflare.RadarHTTPSummaryBotClassParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPSummaryBotClassParamsDeviceType{cloudflare.RadarHTTPSummaryBotClassParamsDeviceTypeDesktop, cloudflare.RadarHTTPSummaryBotClassParamsDeviceTypeMobile, cloudflare.RadarHTTPSummaryBotClassParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPSummaryBotClassParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPSummaryBotClassParamsHTTPProtocol{cloudflare.RadarHTTPSummaryBotClassParamsHTTPProtocolHTTP, cloudflare.RadarHTTPSummaryBotClassParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPSummaryBotClassParamsHTTPVersion{cloudflare.RadarHTTPSummaryBotClassParamsHTTPVersionHttPv1, cloudflare.RadarHTTPSummaryBotClassParamsHTTPVersionHttPv2, cloudflare.RadarHTTPSummaryBotClassParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPSummaryBotClassParamsIPVersion{cloudflare.RadarHTTPSummaryBotClassParamsIPVersionIPv4, cloudflare.RadarHTTPSummaryBotClassParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		OS:           cloudflare.F([]cloudflare.RadarHTTPSummaryBotClassParamsOS{cloudflare.RadarHTTPSummaryBotClassParamsOSWindows, cloudflare.RadarHTTPSummaryBotClassParamsOSMacosx, cloudflare.RadarHTTPSummaryBotClassParamsOSIos}),
		TLSVersion:   cloudflare.F([]cloudflare.RadarHTTPSummaryBotClassParamsTLSVersion{cloudflare.RadarHTTPSummaryBotClassParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPSummaryBotClassParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPSummaryBotClassParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPSummaryDeviceTypeWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.DeviceType(context.TODO(), cloudflare.RadarHTTPSummaryDeviceTypeParams{
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPSummaryDeviceTypeParamsBotClass{cloudflare.RadarHTTPSummaryDeviceTypeParamsBotClassLikelyAutomated, cloudflare.RadarHTTPSummaryDeviceTypeParamsBotClassLikelyHuman}),
		Continent:    cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPSummaryDeviceTypeParamsDateRange{cloudflare.RadarHTTPSummaryDeviceTypeParamsDateRange1d, cloudflare.RadarHTTPSummaryDeviceTypeParamsDateRange2d, cloudflare.RadarHTTPSummaryDeviceTypeParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:       cloudflare.F(cloudflare.RadarHTTPSummaryDeviceTypeParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPSummaryDeviceTypeParamsHTTPProtocol{cloudflare.RadarHTTPSummaryDeviceTypeParamsHTTPProtocolHTTP, cloudflare.RadarHTTPSummaryDeviceTypeParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPSummaryDeviceTypeParamsHTTPVersion{cloudflare.RadarHTTPSummaryDeviceTypeParamsHTTPVersionHttPv1, cloudflare.RadarHTTPSummaryDeviceTypeParamsHTTPVersionHttPv2, cloudflare.RadarHTTPSummaryDeviceTypeParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPSummaryDeviceTypeParamsIPVersion{cloudflare.RadarHTTPSummaryDeviceTypeParamsIPVersionIPv4, cloudflare.RadarHTTPSummaryDeviceTypeParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		OS:           cloudflare.F([]cloudflare.RadarHTTPSummaryDeviceTypeParamsOS{cloudflare.RadarHTTPSummaryDeviceTypeParamsOSWindows, cloudflare.RadarHTTPSummaryDeviceTypeParamsOSMacosx, cloudflare.RadarHTTPSummaryDeviceTypeParamsOSIos}),
		TLSVersion:   cloudflare.F([]cloudflare.RadarHTTPSummaryDeviceTypeParamsTLSVersion{cloudflare.RadarHTTPSummaryDeviceTypeParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPSummaryDeviceTypeParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPSummaryDeviceTypeParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPSummaryHTTPProtocolWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.HTTPProtocol(context.TODO(), cloudflare.RadarHTTPSummaryHTTPProtocolParams{
		ASN:         cloudflare.F([]string{"string", "string", "string"}),
		BotClass:    cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPProtocolParamsBotClass{cloudflare.RadarHTTPSummaryHTTPProtocolParamsBotClassLikelyAutomated, cloudflare.RadarHTTPSummaryHTTPProtocolParamsBotClassLikelyHuman}),
		Continent:   cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPProtocolParamsDateRange{cloudflare.RadarHTTPSummaryHTTPProtocolParamsDateRange1d, cloudflare.RadarHTTPSummaryHTTPProtocolParamsDateRange2d, cloudflare.RadarHTTPSummaryHTTPProtocolParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:  cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPProtocolParamsDeviceType{cloudflare.RadarHTTPSummaryHTTPProtocolParamsDeviceTypeDesktop, cloudflare.RadarHTTPSummaryHTTPProtocolParamsDeviceTypeMobile, cloudflare.RadarHTTPSummaryHTTPProtocolParamsDeviceTypeOther}),
		Format:      cloudflare.F(cloudflare.RadarHTTPSummaryHTTPProtocolParamsFormatJson),
		HTTPVersion: cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPProtocolParamsHTTPVersion{cloudflare.RadarHTTPSummaryHTTPProtocolParamsHTTPVersionHttPv1, cloudflare.RadarHTTPSummaryHTTPProtocolParamsHTTPVersionHttPv2, cloudflare.RadarHTTPSummaryHTTPProtocolParamsHTTPVersionHttPv3}),
		IPVersion:   cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPProtocolParamsIPVersion{cloudflare.RadarHTTPSummaryHTTPProtocolParamsIPVersionIPv4, cloudflare.RadarHTTPSummaryHTTPProtocolParamsIPVersionIPv6}),
		Location:    cloudflare.F([]string{"string", "string", "string"}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		OS:          cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPProtocolParamsOS{cloudflare.RadarHTTPSummaryHTTPProtocolParamsOSWindows, cloudflare.RadarHTTPSummaryHTTPProtocolParamsOSMacosx, cloudflare.RadarHTTPSummaryHTTPProtocolParamsOSIos}),
		TLSVersion:  cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPProtocolParamsTLSVersion{cloudflare.RadarHTTPSummaryHTTPProtocolParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPSummaryHTTPProtocolParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPSummaryHTTPProtocolParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPSummaryHTTPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.HTTPVersion(context.TODO(), cloudflare.RadarHTTPSummaryHTTPVersionParams{
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPVersionParamsBotClass{cloudflare.RadarHTTPSummaryHTTPVersionParamsBotClassLikelyAutomated, cloudflare.RadarHTTPSummaryHTTPVersionParamsBotClassLikelyHuman}),
		Continent:    cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPVersionParamsDateRange{cloudflare.RadarHTTPSummaryHTTPVersionParamsDateRange1d, cloudflare.RadarHTTPSummaryHTTPVersionParamsDateRange2d, cloudflare.RadarHTTPSummaryHTTPVersionParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPVersionParamsDeviceType{cloudflare.RadarHTTPSummaryHTTPVersionParamsDeviceTypeDesktop, cloudflare.RadarHTTPSummaryHTTPVersionParamsDeviceTypeMobile, cloudflare.RadarHTTPSummaryHTTPVersionParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPSummaryHTTPVersionParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPVersionParamsHTTPProtocol{cloudflare.RadarHTTPSummaryHTTPVersionParamsHTTPProtocolHTTP, cloudflare.RadarHTTPSummaryHTTPVersionParamsHTTPProtocolHTTPS}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPVersionParamsIPVersion{cloudflare.RadarHTTPSummaryHTTPVersionParamsIPVersionIPv4, cloudflare.RadarHTTPSummaryHTTPVersionParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		OS:           cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPVersionParamsOS{cloudflare.RadarHTTPSummaryHTTPVersionParamsOSWindows, cloudflare.RadarHTTPSummaryHTTPVersionParamsOSMacosx, cloudflare.RadarHTTPSummaryHTTPVersionParamsOSIos}),
		TLSVersion:   cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPVersionParamsTLSVersion{cloudflare.RadarHTTPSummaryHTTPVersionParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPSummaryHTTPVersionParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPSummaryHTTPVersionParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPSummaryIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.IPVersion(context.TODO(), cloudflare.RadarHTTPSummaryIPVersionParams{
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPSummaryIPVersionParamsBotClass{cloudflare.RadarHTTPSummaryIPVersionParamsBotClassLikelyAutomated, cloudflare.RadarHTTPSummaryIPVersionParamsBotClassLikelyHuman}),
		Continent:    cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPSummaryIPVersionParamsDateRange{cloudflare.RadarHTTPSummaryIPVersionParamsDateRange1d, cloudflare.RadarHTTPSummaryIPVersionParamsDateRange2d, cloudflare.RadarHTTPSummaryIPVersionParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPSummaryIPVersionParamsDeviceType{cloudflare.RadarHTTPSummaryIPVersionParamsDeviceTypeDesktop, cloudflare.RadarHTTPSummaryIPVersionParamsDeviceTypeMobile, cloudflare.RadarHTTPSummaryIPVersionParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPSummaryIPVersionParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPSummaryIPVersionParamsHTTPProtocol{cloudflare.RadarHTTPSummaryIPVersionParamsHTTPProtocolHTTP, cloudflare.RadarHTTPSummaryIPVersionParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPSummaryIPVersionParamsHTTPVersion{cloudflare.RadarHTTPSummaryIPVersionParamsHTTPVersionHttPv1, cloudflare.RadarHTTPSummaryIPVersionParamsHTTPVersionHttPv2, cloudflare.RadarHTTPSummaryIPVersionParamsHTTPVersionHttPv3}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		OS:           cloudflare.F([]cloudflare.RadarHTTPSummaryIPVersionParamsOS{cloudflare.RadarHTTPSummaryIPVersionParamsOSWindows, cloudflare.RadarHTTPSummaryIPVersionParamsOSMacosx, cloudflare.RadarHTTPSummaryIPVersionParamsOSIos}),
		TLSVersion:   cloudflare.F([]cloudflare.RadarHTTPSummaryIPVersionParamsTLSVersion{cloudflare.RadarHTTPSummaryIPVersionParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPSummaryIPVersionParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPSummaryIPVersionParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPSummaryOSWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.OS(context.TODO(), cloudflare.RadarHTTPSummaryOSParams{
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPSummaryOSParamsBotClass{cloudflare.RadarHTTPSummaryOSParamsBotClassLikelyAutomated, cloudflare.RadarHTTPSummaryOSParamsBotClassLikelyHuman}),
		Continent:    cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPSummaryOSParamsDateRange{cloudflare.RadarHTTPSummaryOSParamsDateRange1d, cloudflare.RadarHTTPSummaryOSParamsDateRange2d, cloudflare.RadarHTTPSummaryOSParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPSummaryOSParamsDeviceType{cloudflare.RadarHTTPSummaryOSParamsDeviceTypeDesktop, cloudflare.RadarHTTPSummaryOSParamsDeviceTypeMobile, cloudflare.RadarHTTPSummaryOSParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPSummaryOSParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPSummaryOSParamsHTTPProtocol{cloudflare.RadarHTTPSummaryOSParamsHTTPProtocolHTTP, cloudflare.RadarHTTPSummaryOSParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPSummaryOSParamsHTTPVersion{cloudflare.RadarHTTPSummaryOSParamsHTTPVersionHttPv1, cloudflare.RadarHTTPSummaryOSParamsHTTPVersionHttPv2, cloudflare.RadarHTTPSummaryOSParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPSummaryOSParamsIPVersion{cloudflare.RadarHTTPSummaryOSParamsIPVersionIPv4, cloudflare.RadarHTTPSummaryOSParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		TLSVersion:   cloudflare.F([]cloudflare.RadarHTTPSummaryOSParamsTLSVersion{cloudflare.RadarHTTPSummaryOSParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPSummaryOSParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPSummaryOSParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPSummaryTLSVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.TLSVersion(context.TODO(), cloudflare.RadarHTTPSummaryTLSVersionParams{
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPSummaryTLSVersionParamsBotClass{cloudflare.RadarHTTPSummaryTLSVersionParamsBotClassLikelyAutomated, cloudflare.RadarHTTPSummaryTLSVersionParamsBotClassLikelyHuman}),
		Continent:    cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPSummaryTLSVersionParamsDateRange{cloudflare.RadarHTTPSummaryTLSVersionParamsDateRange1d, cloudflare.RadarHTTPSummaryTLSVersionParamsDateRange2d, cloudflare.RadarHTTPSummaryTLSVersionParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPSummaryTLSVersionParamsDeviceType{cloudflare.RadarHTTPSummaryTLSVersionParamsDeviceTypeDesktop, cloudflare.RadarHTTPSummaryTLSVersionParamsDeviceTypeMobile, cloudflare.RadarHTTPSummaryTLSVersionParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPSummaryTLSVersionParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPSummaryTLSVersionParamsHTTPProtocol{cloudflare.RadarHTTPSummaryTLSVersionParamsHTTPProtocolHTTP, cloudflare.RadarHTTPSummaryTLSVersionParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPSummaryTLSVersionParamsHTTPVersion{cloudflare.RadarHTTPSummaryTLSVersionParamsHTTPVersionHttPv1, cloudflare.RadarHTTPSummaryTLSVersionParamsHTTPVersionHttPv2, cloudflare.RadarHTTPSummaryTLSVersionParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPSummaryTLSVersionParamsIPVersion{cloudflare.RadarHTTPSummaryTLSVersionParamsIPVersionIPv4, cloudflare.RadarHTTPSummaryTLSVersionParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		OS:           cloudflare.F([]cloudflare.RadarHTTPSummaryTLSVersionParamsOS{cloudflare.RadarHTTPSummaryTLSVersionParamsOSWindows, cloudflare.RadarHTTPSummaryTLSVersionParamsOSMacosx, cloudflare.RadarHTTPSummaryTLSVersionParamsOSIos}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
