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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Radar.HTTP.Summary.BotClass(context.TODO(), cloudflare.RadarHTTPSummaryBotClassParams{
		Asn:          cloudflare.F([]string{"string", "string", "string"}),
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
		Os:           cloudflare.F([]cloudflare.RadarHTTPSummaryBotClassParamsO{cloudflare.RadarHTTPSummaryBotClassParamsOWindows, cloudflare.RadarHTTPSummaryBotClassParamsOMacosx, cloudflare.RadarHTTPSummaryBotClassParamsOIos}),
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Radar.HTTP.Summary.DeviceType(context.TODO(), cloudflare.RadarHTTPSummaryDeviceTypeParams{
		Asn:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPSummaryDeviceTypeParamsBotClass{cloudflare.RadarHTTPSummaryDeviceTypeParamsBotClassLikelyAutomated, cloudflare.RadarHTTPSummaryDeviceTypeParamsBotClassLikelyHuman}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPSummaryDeviceTypeParamsDateRange{cloudflare.RadarHTTPSummaryDeviceTypeParamsDateRange1d, cloudflare.RadarHTTPSummaryDeviceTypeParamsDateRange2d, cloudflare.RadarHTTPSummaryDeviceTypeParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:       cloudflare.F(cloudflare.RadarHTTPSummaryDeviceTypeParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPSummaryDeviceTypeParamsHTTPProtocol{cloudflare.RadarHTTPSummaryDeviceTypeParamsHTTPProtocolHTTP, cloudflare.RadarHTTPSummaryDeviceTypeParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPSummaryDeviceTypeParamsHTTPVersion{cloudflare.RadarHTTPSummaryDeviceTypeParamsHTTPVersionHttPv1, cloudflare.RadarHTTPSummaryDeviceTypeParamsHTTPVersionHttPv2, cloudflare.RadarHTTPSummaryDeviceTypeParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPSummaryDeviceTypeParamsIPVersion{cloudflare.RadarHTTPSummaryDeviceTypeParamsIPVersionIPv4, cloudflare.RadarHTTPSummaryDeviceTypeParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		Os:           cloudflare.F([]cloudflare.RadarHTTPSummaryDeviceTypeParamsO{cloudflare.RadarHTTPSummaryDeviceTypeParamsOWindows, cloudflare.RadarHTTPSummaryDeviceTypeParamsOMacosx, cloudflare.RadarHTTPSummaryDeviceTypeParamsOIos}),
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Radar.HTTP.Summary.HTTPProtocol(context.TODO(), cloudflare.RadarHTTPSummaryHTTPProtocolParams{
		Asn:         cloudflare.F([]string{"string", "string", "string"}),
		BotClass:    cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPProtocolParamsBotClass{cloudflare.RadarHTTPSummaryHTTPProtocolParamsBotClassLikelyAutomated, cloudflare.RadarHTTPSummaryHTTPProtocolParamsBotClassLikelyHuman}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPProtocolParamsDateRange{cloudflare.RadarHTTPSummaryHTTPProtocolParamsDateRange1d, cloudflare.RadarHTTPSummaryHTTPProtocolParamsDateRange2d, cloudflare.RadarHTTPSummaryHTTPProtocolParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:  cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPProtocolParamsDeviceType{cloudflare.RadarHTTPSummaryHTTPProtocolParamsDeviceTypeDesktop, cloudflare.RadarHTTPSummaryHTTPProtocolParamsDeviceTypeMobile, cloudflare.RadarHTTPSummaryHTTPProtocolParamsDeviceTypeOther}),
		Format:      cloudflare.F(cloudflare.RadarHTTPSummaryHTTPProtocolParamsFormatJson),
		HTTPVersion: cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPProtocolParamsHTTPVersion{cloudflare.RadarHTTPSummaryHTTPProtocolParamsHTTPVersionHttPv1, cloudflare.RadarHTTPSummaryHTTPProtocolParamsHTTPVersionHttPv2, cloudflare.RadarHTTPSummaryHTTPProtocolParamsHTTPVersionHttPv3}),
		IPVersion:   cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPProtocolParamsIPVersion{cloudflare.RadarHTTPSummaryHTTPProtocolParamsIPVersionIPv4, cloudflare.RadarHTTPSummaryHTTPProtocolParamsIPVersionIPv6}),
		Location:    cloudflare.F([]string{"string", "string", "string"}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		Os:          cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPProtocolParamsO{cloudflare.RadarHTTPSummaryHTTPProtocolParamsOWindows, cloudflare.RadarHTTPSummaryHTTPProtocolParamsOMacosx, cloudflare.RadarHTTPSummaryHTTPProtocolParamsOIos}),
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Radar.HTTP.Summary.HTTPVersion(context.TODO(), cloudflare.RadarHTTPSummaryHTTPVersionParams{
		Asn:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPVersionParamsBotClass{cloudflare.RadarHTTPSummaryHTTPVersionParamsBotClassLikelyAutomated, cloudflare.RadarHTTPSummaryHTTPVersionParamsBotClassLikelyHuman}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPVersionParamsDateRange{cloudflare.RadarHTTPSummaryHTTPVersionParamsDateRange1d, cloudflare.RadarHTTPSummaryHTTPVersionParamsDateRange2d, cloudflare.RadarHTTPSummaryHTTPVersionParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPVersionParamsDeviceType{cloudflare.RadarHTTPSummaryHTTPVersionParamsDeviceTypeDesktop, cloudflare.RadarHTTPSummaryHTTPVersionParamsDeviceTypeMobile, cloudflare.RadarHTTPSummaryHTTPVersionParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPSummaryHTTPVersionParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPVersionParamsHTTPProtocol{cloudflare.RadarHTTPSummaryHTTPVersionParamsHTTPProtocolHTTP, cloudflare.RadarHTTPSummaryHTTPVersionParamsHTTPProtocolHTTPS}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPVersionParamsIPVersion{cloudflare.RadarHTTPSummaryHTTPVersionParamsIPVersionIPv4, cloudflare.RadarHTTPSummaryHTTPVersionParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		Os:           cloudflare.F([]cloudflare.RadarHTTPSummaryHTTPVersionParamsO{cloudflare.RadarHTTPSummaryHTTPVersionParamsOWindows, cloudflare.RadarHTTPSummaryHTTPVersionParamsOMacosx, cloudflare.RadarHTTPSummaryHTTPVersionParamsOIos}),
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Radar.HTTP.Summary.IPVersion(context.TODO(), cloudflare.RadarHTTPSummaryIPVersionParams{
		Asn:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPSummaryIPVersionParamsBotClass{cloudflare.RadarHTTPSummaryIPVersionParamsBotClassLikelyAutomated, cloudflare.RadarHTTPSummaryIPVersionParamsBotClassLikelyHuman}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPSummaryIPVersionParamsDateRange{cloudflare.RadarHTTPSummaryIPVersionParamsDateRange1d, cloudflare.RadarHTTPSummaryIPVersionParamsDateRange2d, cloudflare.RadarHTTPSummaryIPVersionParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPSummaryIPVersionParamsDeviceType{cloudflare.RadarHTTPSummaryIPVersionParamsDeviceTypeDesktop, cloudflare.RadarHTTPSummaryIPVersionParamsDeviceTypeMobile, cloudflare.RadarHTTPSummaryIPVersionParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPSummaryIPVersionParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPSummaryIPVersionParamsHTTPProtocol{cloudflare.RadarHTTPSummaryIPVersionParamsHTTPProtocolHTTP, cloudflare.RadarHTTPSummaryIPVersionParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPSummaryIPVersionParamsHTTPVersion{cloudflare.RadarHTTPSummaryIPVersionParamsHTTPVersionHttPv1, cloudflare.RadarHTTPSummaryIPVersionParamsHTTPVersionHttPv2, cloudflare.RadarHTTPSummaryIPVersionParamsHTTPVersionHttPv3}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		Os:           cloudflare.F([]cloudflare.RadarHTTPSummaryIPVersionParamsO{cloudflare.RadarHTTPSummaryIPVersionParamsOWindows, cloudflare.RadarHTTPSummaryIPVersionParamsOMacosx, cloudflare.RadarHTTPSummaryIPVersionParamsOIos}),
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

func TestRadarHTTPSummaryOsWithOptionalParams(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Radar.HTTP.Summary.Os(context.TODO(), cloudflare.RadarHTTPSummaryOsParams{
		Asn:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPSummaryOsParamsBotClass{cloudflare.RadarHTTPSummaryOsParamsBotClassLikelyAutomated, cloudflare.RadarHTTPSummaryOsParamsBotClassLikelyHuman}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPSummaryOsParamsDateRange{cloudflare.RadarHTTPSummaryOsParamsDateRange1d, cloudflare.RadarHTTPSummaryOsParamsDateRange2d, cloudflare.RadarHTTPSummaryOsParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPSummaryOsParamsDeviceType{cloudflare.RadarHTTPSummaryOsParamsDeviceTypeDesktop, cloudflare.RadarHTTPSummaryOsParamsDeviceTypeMobile, cloudflare.RadarHTTPSummaryOsParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPSummaryOsParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPSummaryOsParamsHTTPProtocol{cloudflare.RadarHTTPSummaryOsParamsHTTPProtocolHTTP, cloudflare.RadarHTTPSummaryOsParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPSummaryOsParamsHTTPVersion{cloudflare.RadarHTTPSummaryOsParamsHTTPVersionHttPv1, cloudflare.RadarHTTPSummaryOsParamsHTTPVersionHttPv2, cloudflare.RadarHTTPSummaryOsParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPSummaryOsParamsIPVersion{cloudflare.RadarHTTPSummaryOsParamsIPVersionIPv4, cloudflare.RadarHTTPSummaryOsParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		TLSVersion:   cloudflare.F([]cloudflare.RadarHTTPSummaryOsParamsTLSVersion{cloudflare.RadarHTTPSummaryOsParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPSummaryOsParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPSummaryOsParamsTLSVersionTlSv1_2}),
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Radar.HTTP.Summary.TLSVersion(context.TODO(), cloudflare.RadarHTTPSummaryTLSVersionParams{
		Asn:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPSummaryTLSVersionParamsBotClass{cloudflare.RadarHTTPSummaryTLSVersionParamsBotClassLikelyAutomated, cloudflare.RadarHTTPSummaryTLSVersionParamsBotClassLikelyHuman}),
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
		Os:           cloudflare.F([]cloudflare.RadarHTTPSummaryTLSVersionParamsO{cloudflare.RadarHTTPSummaryTLSVersionParamsOWindows, cloudflare.RadarHTTPSummaryTLSVersionParamsOMacosx, cloudflare.RadarHTTPSummaryTLSVersionParamsOIos}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
