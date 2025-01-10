// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/radar"
)

func TestHTTPSummaryBotClassWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.BotClass(context.TODO(), radar.HTTPSummaryBotClassParams{
		ASN:          cloudflare.F([]string{"string"}),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		DeviceType:   cloudflare.F([]radar.HTTPSummaryBotClassParamsDeviceType{radar.HTTPSummaryBotClassParamsDeviceTypeDesktop}),
		Format:       cloudflare.F(radar.HTTPSummaryBotClassParamsFormatJson),
		HTTPProtocol: cloudflare.F([]radar.HTTPSummaryBotClassParamsHTTPProtocol{radar.HTTPSummaryBotClassParamsHTTPProtocolHTTP}),
		HTTPVersion:  cloudflare.F([]radar.HTTPSummaryBotClassParamsHTTPVersion{radar.HTTPSummaryBotClassParamsHTTPVersionHttPv1}),
		IPVersion:    cloudflare.F([]radar.HTTPSummaryBotClassParamsIPVersion{radar.HTTPSummaryBotClassParamsIPVersionIPv4}),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"string"}),
		OS:           cloudflare.F([]radar.HTTPSummaryBotClassParamsOS{radar.HTTPSummaryBotClassParamsOSWindows}),
		TLSVersion:   cloudflare.F([]radar.HTTPSummaryBotClassParamsTLSVersion{radar.HTTPSummaryBotClassParamsTLSVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPSummaryDeviceTypeWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.DeviceType(context.TODO(), radar.HTTPSummaryDeviceTypeParams{
		ASN:          cloudflare.F([]string{"string"}),
		BotClass:     cloudflare.F([]radar.HTTPSummaryDeviceTypeParamsBotClass{radar.HTTPSummaryDeviceTypeParamsBotClassLikelyAutomated}),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		Format:       cloudflare.F(radar.HTTPSummaryDeviceTypeParamsFormatJson),
		HTTPProtocol: cloudflare.F([]radar.HTTPSummaryDeviceTypeParamsHTTPProtocol{radar.HTTPSummaryDeviceTypeParamsHTTPProtocolHTTP}),
		HTTPVersion:  cloudflare.F([]radar.HTTPSummaryDeviceTypeParamsHTTPVersion{radar.HTTPSummaryDeviceTypeParamsHTTPVersionHttPv1}),
		IPVersion:    cloudflare.F([]radar.HTTPSummaryDeviceTypeParamsIPVersion{radar.HTTPSummaryDeviceTypeParamsIPVersionIPv4}),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"string"}),
		OS:           cloudflare.F([]radar.HTTPSummaryDeviceTypeParamsOS{radar.HTTPSummaryDeviceTypeParamsOSWindows}),
		TLSVersion:   cloudflare.F([]radar.HTTPSummaryDeviceTypeParamsTLSVersion{radar.HTTPSummaryDeviceTypeParamsTLSVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPSummaryHTTPProtocolWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.HTTPProtocol(context.TODO(), radar.HTTPSummaryHTTPProtocolParams{
		ASN:         cloudflare.F([]string{"string"}),
		BotClass:    cloudflare.F([]radar.HTTPSummaryHTTPProtocolParamsBotClass{radar.HTTPSummaryHTTPProtocolParamsBotClassLikelyAutomated}),
		Continent:   cloudflare.F([]string{"string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now()}),
		DateRange:   cloudflare.F([]string{"7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now()}),
		DeviceType:  cloudflare.F([]radar.HTTPSummaryHTTPProtocolParamsDeviceType{radar.HTTPSummaryHTTPProtocolParamsDeviceTypeDesktop}),
		Format:      cloudflare.F(radar.HTTPSummaryHTTPProtocolParamsFormatJson),
		HTTPVersion: cloudflare.F([]radar.HTTPSummaryHTTPProtocolParamsHTTPVersion{radar.HTTPSummaryHTTPProtocolParamsHTTPVersionHttPv1}),
		IPVersion:   cloudflare.F([]radar.HTTPSummaryHTTPProtocolParamsIPVersion{radar.HTTPSummaryHTTPProtocolParamsIPVersionIPv4}),
		Location:    cloudflare.F([]string{"string"}),
		Name:        cloudflare.F([]string{"string"}),
		OS:          cloudflare.F([]radar.HTTPSummaryHTTPProtocolParamsOS{radar.HTTPSummaryHTTPProtocolParamsOSWindows}),
		TLSVersion:  cloudflare.F([]radar.HTTPSummaryHTTPProtocolParamsTLSVersion{radar.HTTPSummaryHTTPProtocolParamsTLSVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPSummaryHTTPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.HTTPVersion(context.TODO(), radar.HTTPSummaryHTTPVersionParams{
		ASN:          cloudflare.F([]string{"string"}),
		BotClass:     cloudflare.F([]radar.HTTPSummaryHTTPVersionParamsBotClass{radar.HTTPSummaryHTTPVersionParamsBotClassLikelyAutomated}),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		DeviceType:   cloudflare.F([]radar.HTTPSummaryHTTPVersionParamsDeviceType{radar.HTTPSummaryHTTPVersionParamsDeviceTypeDesktop}),
		Format:       cloudflare.F(radar.HTTPSummaryHTTPVersionParamsFormatJson),
		HTTPProtocol: cloudflare.F([]radar.HTTPSummaryHTTPVersionParamsHTTPProtocol{radar.HTTPSummaryHTTPVersionParamsHTTPProtocolHTTP}),
		IPVersion:    cloudflare.F([]radar.HTTPSummaryHTTPVersionParamsIPVersion{radar.HTTPSummaryHTTPVersionParamsIPVersionIPv4}),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"string"}),
		OS:           cloudflare.F([]radar.HTTPSummaryHTTPVersionParamsOS{radar.HTTPSummaryHTTPVersionParamsOSWindows}),
		TLSVersion:   cloudflare.F([]radar.HTTPSummaryHTTPVersionParamsTLSVersion{radar.HTTPSummaryHTTPVersionParamsTLSVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPSummaryIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.IPVersion(context.TODO(), radar.HTTPSummaryIPVersionParams{
		ASN:          cloudflare.F([]string{"string"}),
		BotClass:     cloudflare.F([]radar.HTTPSummaryIPVersionParamsBotClass{radar.HTTPSummaryIPVersionParamsBotClassLikelyAutomated}),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		DeviceType:   cloudflare.F([]radar.HTTPSummaryIPVersionParamsDeviceType{radar.HTTPSummaryIPVersionParamsDeviceTypeDesktop}),
		Format:       cloudflare.F(radar.HTTPSummaryIPVersionParamsFormatJson),
		HTTPProtocol: cloudflare.F([]radar.HTTPSummaryIPVersionParamsHTTPProtocol{radar.HTTPSummaryIPVersionParamsHTTPProtocolHTTP}),
		HTTPVersion:  cloudflare.F([]radar.HTTPSummaryIPVersionParamsHTTPVersion{radar.HTTPSummaryIPVersionParamsHTTPVersionHttPv1}),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"string"}),
		OS:           cloudflare.F([]radar.HTTPSummaryIPVersionParamsOS{radar.HTTPSummaryIPVersionParamsOSWindows}),
		TLSVersion:   cloudflare.F([]radar.HTTPSummaryIPVersionParamsTLSVersion{radar.HTTPSummaryIPVersionParamsTLSVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPSummaryOSWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.OS(context.TODO(), radar.HTTPSummaryOSParams{
		ASN:          cloudflare.F([]string{"string"}),
		BotClass:     cloudflare.F([]radar.HTTPSummaryOSParamsBotClass{radar.HTTPSummaryOSParamsBotClassLikelyAutomated}),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		DeviceType:   cloudflare.F([]radar.HTTPSummaryOSParamsDeviceType{radar.HTTPSummaryOSParamsDeviceTypeDesktop}),
		Format:       cloudflare.F(radar.HTTPSummaryOSParamsFormatJson),
		HTTPProtocol: cloudflare.F([]radar.HTTPSummaryOSParamsHTTPProtocol{radar.HTTPSummaryOSParamsHTTPProtocolHTTP}),
		HTTPVersion:  cloudflare.F([]radar.HTTPSummaryOSParamsHTTPVersion{radar.HTTPSummaryOSParamsHTTPVersionHttPv1}),
		IPVersion:    cloudflare.F([]radar.HTTPSummaryOSParamsIPVersion{radar.HTTPSummaryOSParamsIPVersionIPv4}),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"string"}),
		TLSVersion:   cloudflare.F([]radar.HTTPSummaryOSParamsTLSVersion{radar.HTTPSummaryOSParamsTLSVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPSummaryPostQuantumWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.PostQuantum(context.TODO(), radar.HTTPSummaryPostQuantumParams{
		ASN:          cloudflare.F([]string{"string"}),
		BotClass:     cloudflare.F([]radar.HTTPSummaryPostQuantumParamsBotClass{radar.HTTPSummaryPostQuantumParamsBotClassLikelyAutomated}),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		DeviceType:   cloudflare.F([]radar.HTTPSummaryPostQuantumParamsDeviceType{radar.HTTPSummaryPostQuantumParamsDeviceTypeDesktop}),
		Format:       cloudflare.F(radar.HTTPSummaryPostQuantumParamsFormatJson),
		HTTPProtocol: cloudflare.F([]radar.HTTPSummaryPostQuantumParamsHTTPProtocol{radar.HTTPSummaryPostQuantumParamsHTTPProtocolHTTP}),
		HTTPVersion:  cloudflare.F([]radar.HTTPSummaryPostQuantumParamsHTTPVersion{radar.HTTPSummaryPostQuantumParamsHTTPVersionHttPv1}),
		IPVersion:    cloudflare.F([]radar.HTTPSummaryPostQuantumParamsIPVersion{radar.HTTPSummaryPostQuantumParamsIPVersionIPv4}),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"string"}),
		OS:           cloudflare.F([]radar.HTTPSummaryPostQuantumParamsOS{radar.HTTPSummaryPostQuantumParamsOSWindows}),
		TLSVersion:   cloudflare.F([]radar.HTTPSummaryPostQuantumParamsTLSVersion{radar.HTTPSummaryPostQuantumParamsTLSVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPSummaryTLSVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.TLSVersion(context.TODO(), radar.HTTPSummaryTLSVersionParams{
		ASN:          cloudflare.F([]string{"string"}),
		BotClass:     cloudflare.F([]radar.HTTPSummaryTLSVersionParamsBotClass{radar.HTTPSummaryTLSVersionParamsBotClassLikelyAutomated}),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		DeviceType:   cloudflare.F([]radar.HTTPSummaryTLSVersionParamsDeviceType{radar.HTTPSummaryTLSVersionParamsDeviceTypeDesktop}),
		Format:       cloudflare.F(radar.HTTPSummaryTLSVersionParamsFormatJson),
		HTTPProtocol: cloudflare.F([]radar.HTTPSummaryTLSVersionParamsHTTPProtocol{radar.HTTPSummaryTLSVersionParamsHTTPProtocolHTTP}),
		HTTPVersion:  cloudflare.F([]radar.HTTPSummaryTLSVersionParamsHTTPVersion{radar.HTTPSummaryTLSVersionParamsHTTPVersionHttPv1}),
		IPVersion:    cloudflare.F([]radar.HTTPSummaryTLSVersionParamsIPVersion{radar.HTTPSummaryTLSVersionParamsIPVersionIPv4}),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"string"}),
		OS:           cloudflare.F([]radar.HTTPSummaryTLSVersionParamsOS{radar.HTTPSummaryTLSVersionParamsOSWindows}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
