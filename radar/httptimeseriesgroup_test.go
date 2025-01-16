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
		AggInterval:  cloudflare.F(radar.HTTPTimeseriesGroupBotClassParamsAggInterval15m),
		ASN:          cloudflare.F([]string{"string"}),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		DeviceType:   cloudflare.F([]radar.HTTPTimeseriesGroupBotClassParamsDeviceType{radar.HTTPTimeseriesGroupBotClassParamsDeviceTypeDesktop}),
		Format:       cloudflare.F(radar.HTTPTimeseriesGroupBotClassParamsFormatJson),
		HTTPProtocol: cloudflare.F([]radar.HTTPTimeseriesGroupBotClassParamsHTTPProtocol{radar.HTTPTimeseriesGroupBotClassParamsHTTPProtocolHTTP}),
		HTTPVersion:  cloudflare.F([]radar.HTTPTimeseriesGroupBotClassParamsHTTPVersion{radar.HTTPTimeseriesGroupBotClassParamsHTTPVersionHttPv1}),
		IPVersion:    cloudflare.F([]radar.HTTPTimeseriesGroupBotClassParamsIPVersion{radar.HTTPTimeseriesGroupBotClassParamsIPVersionIPv4}),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"string"}),
		OS:           cloudflare.F([]radar.HTTPTimeseriesGroupBotClassParamsOS{radar.HTTPTimeseriesGroupBotClassParamsOSWindows}),
		TLSVersion:   cloudflare.F([]radar.HTTPTimeseriesGroupBotClassParamsTLSVersion{radar.HTTPTimeseriesGroupBotClassParamsTLSVersionTlSv1_0}),
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
		AggInterval:   cloudflare.F(radar.HTTPTimeseriesGroupBrowserParamsAggInterval15m),
		ASN:           cloudflare.F([]string{"string"}),
		BotClass:      cloudflare.F([]radar.HTTPTimeseriesGroupBrowserParamsBotClass{radar.HTTPTimeseriesGroupBrowserParamsBotClassLikelyAutomated}),
		Continent:     cloudflare.F([]string{"string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now()}),
		DateRange:     cloudflare.F([]string{"7d"}),
		DateStart:     cloudflare.F([]time.Time{time.Now()}),
		DeviceType:    cloudflare.F([]radar.HTTPTimeseriesGroupBrowserParamsDeviceType{radar.HTTPTimeseriesGroupBrowserParamsDeviceTypeDesktop}),
		Format:        cloudflare.F(radar.HTTPTimeseriesGroupBrowserParamsFormatJson),
		HTTPProtocol:  cloudflare.F([]radar.HTTPTimeseriesGroupBrowserParamsHTTPProtocol{radar.HTTPTimeseriesGroupBrowserParamsHTTPProtocolHTTP}),
		HTTPVersion:   cloudflare.F([]radar.HTTPTimeseriesGroupBrowserParamsHTTPVersion{radar.HTTPTimeseriesGroupBrowserParamsHTTPVersionHttPv1}),
		IPVersion:     cloudflare.F([]radar.HTTPTimeseriesGroupBrowserParamsIPVersion{radar.HTTPTimeseriesGroupBrowserParamsIPVersionIPv4}),
		LimitPerGroup: cloudflare.F(int64(4)),
		Location:      cloudflare.F([]string{"string"}),
		Name:          cloudflare.F([]string{"string"}),
		OS:            cloudflare.F([]radar.HTTPTimeseriesGroupBrowserParamsOS{radar.HTTPTimeseriesGroupBrowserParamsOSWindows}),
		TLSVersion:    cloudflare.F([]radar.HTTPTimeseriesGroupBrowserParamsTLSVersion{radar.HTTPTimeseriesGroupBrowserParamsTLSVersionTlSv1_0}),
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
		AggInterval:  cloudflare.F(radar.HTTPTimeseriesGroupBrowserFamilyParamsAggInterval15m),
		ASN:          cloudflare.F([]string{"string"}),
		BotClass:     cloudflare.F([]radar.HTTPTimeseriesGroupBrowserFamilyParamsBotClass{radar.HTTPTimeseriesGroupBrowserFamilyParamsBotClassLikelyAutomated}),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		DeviceType:   cloudflare.F([]radar.HTTPTimeseriesGroupBrowserFamilyParamsDeviceType{radar.HTTPTimeseriesGroupBrowserFamilyParamsDeviceTypeDesktop}),
		Format:       cloudflare.F(radar.HTTPTimeseriesGroupBrowserFamilyParamsFormatJson),
		HTTPProtocol: cloudflare.F([]radar.HTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocol{radar.HTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocolHTTP}),
		HTTPVersion:  cloudflare.F([]radar.HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersion{radar.HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersionHttPv1}),
		IPVersion:    cloudflare.F([]radar.HTTPTimeseriesGroupBrowserFamilyParamsIPVersion{radar.HTTPTimeseriesGroupBrowserFamilyParamsIPVersionIPv4}),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"string"}),
		OS:           cloudflare.F([]radar.HTTPTimeseriesGroupBrowserFamilyParamsOS{radar.HTTPTimeseriesGroupBrowserFamilyParamsOSWindows}),
		TLSVersion:   cloudflare.F([]radar.HTTPTimeseriesGroupBrowserFamilyParamsTLSVersion{radar.HTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSv1_0}),
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
		AggInterval:  cloudflare.F(radar.HTTPTimeseriesGroupDeviceTypeParamsAggInterval15m),
		ASN:          cloudflare.F([]string{"string"}),
		BotClass:     cloudflare.F([]radar.HTTPTimeseriesGroupDeviceTypeParamsBotClass{radar.HTTPTimeseriesGroupDeviceTypeParamsBotClassLikelyAutomated}),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		Format:       cloudflare.F(radar.HTTPTimeseriesGroupDeviceTypeParamsFormatJson),
		HTTPProtocol: cloudflare.F([]radar.HTTPTimeseriesGroupDeviceTypeParamsHTTPProtocol{radar.HTTPTimeseriesGroupDeviceTypeParamsHTTPProtocolHTTP}),
		HTTPVersion:  cloudflare.F([]radar.HTTPTimeseriesGroupDeviceTypeParamsHTTPVersion{radar.HTTPTimeseriesGroupDeviceTypeParamsHTTPVersionHttPv1}),
		IPVersion:    cloudflare.F([]radar.HTTPTimeseriesGroupDeviceTypeParamsIPVersion{radar.HTTPTimeseriesGroupDeviceTypeParamsIPVersionIPv4}),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"string"}),
		OS:           cloudflare.F([]radar.HTTPTimeseriesGroupDeviceTypeParamsOS{radar.HTTPTimeseriesGroupDeviceTypeParamsOSWindows}),
		TLSVersion:   cloudflare.F([]radar.HTTPTimeseriesGroupDeviceTypeParamsTLSVersion{radar.HTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSv1_0}),
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
		AggInterval: cloudflare.F(radar.HTTPTimeseriesGroupHTTPProtocolParamsAggInterval15m),
		ASN:         cloudflare.F([]string{"string"}),
		BotClass:    cloudflare.F([]radar.HTTPTimeseriesGroupHTTPProtocolParamsBotClass{radar.HTTPTimeseriesGroupHTTPProtocolParamsBotClassLikelyAutomated}),
		Continent:   cloudflare.F([]string{"string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now()}),
		DateRange:   cloudflare.F([]string{"7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now()}),
		DeviceType:  cloudflare.F([]radar.HTTPTimeseriesGroupHTTPProtocolParamsDeviceType{radar.HTTPTimeseriesGroupHTTPProtocolParamsDeviceTypeDesktop}),
		Format:      cloudflare.F(radar.HTTPTimeseriesGroupHTTPProtocolParamsFormatJson),
		HTTPVersion: cloudflare.F([]radar.HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersion{radar.HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersionHttPv1}),
		IPVersion:   cloudflare.F([]radar.HTTPTimeseriesGroupHTTPProtocolParamsIPVersion{radar.HTTPTimeseriesGroupHTTPProtocolParamsIPVersionIPv4}),
		Location:    cloudflare.F([]string{"string"}),
		Name:        cloudflare.F([]string{"string"}),
		OS:          cloudflare.F([]radar.HTTPTimeseriesGroupHTTPProtocolParamsOS{radar.HTTPTimeseriesGroupHTTPProtocolParamsOSWindows}),
		TLSVersion:  cloudflare.F([]radar.HTTPTimeseriesGroupHTTPProtocolParamsTLSVersion{radar.HTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSv1_0}),
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
		AggInterval:  cloudflare.F(radar.HTTPTimeseriesGroupHTTPVersionParamsAggInterval15m),
		ASN:          cloudflare.F([]string{"string"}),
		BotClass:     cloudflare.F([]radar.HTTPTimeseriesGroupHTTPVersionParamsBotClass{radar.HTTPTimeseriesGroupHTTPVersionParamsBotClassLikelyAutomated}),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		DeviceType:   cloudflare.F([]radar.HTTPTimeseriesGroupHTTPVersionParamsDeviceType{radar.HTTPTimeseriesGroupHTTPVersionParamsDeviceTypeDesktop}),
		Format:       cloudflare.F(radar.HTTPTimeseriesGroupHTTPVersionParamsFormatJson),
		HTTPProtocol: cloudflare.F([]radar.HTTPTimeseriesGroupHTTPVersionParamsHTTPProtocol{radar.HTTPTimeseriesGroupHTTPVersionParamsHTTPProtocolHTTP}),
		IPVersion:    cloudflare.F([]radar.HTTPTimeseriesGroupHTTPVersionParamsIPVersion{radar.HTTPTimeseriesGroupHTTPVersionParamsIPVersionIPv4}),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"string"}),
		OS:           cloudflare.F([]radar.HTTPTimeseriesGroupHTTPVersionParamsOS{radar.HTTPTimeseriesGroupHTTPVersionParamsOSWindows}),
		TLSVersion:   cloudflare.F([]radar.HTTPTimeseriesGroupHTTPVersionParamsTLSVersion{radar.HTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSv1_0}),
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
		AggInterval:  cloudflare.F(radar.HTTPTimeseriesGroupIPVersionParamsAggInterval15m),
		ASN:          cloudflare.F([]string{"string"}),
		BotClass:     cloudflare.F([]radar.HTTPTimeseriesGroupIPVersionParamsBotClass{radar.HTTPTimeseriesGroupIPVersionParamsBotClassLikelyAutomated}),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		DeviceType:   cloudflare.F([]radar.HTTPTimeseriesGroupIPVersionParamsDeviceType{radar.HTTPTimeseriesGroupIPVersionParamsDeviceTypeDesktop}),
		Format:       cloudflare.F(radar.HTTPTimeseriesGroupIPVersionParamsFormatJson),
		HTTPProtocol: cloudflare.F([]radar.HTTPTimeseriesGroupIPVersionParamsHTTPProtocol{radar.HTTPTimeseriesGroupIPVersionParamsHTTPProtocolHTTP}),
		HTTPVersion:  cloudflare.F([]radar.HTTPTimeseriesGroupIPVersionParamsHTTPVersion{radar.HTTPTimeseriesGroupIPVersionParamsHTTPVersionHttPv1}),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"string"}),
		OS:           cloudflare.F([]radar.HTTPTimeseriesGroupIPVersionParamsOS{radar.HTTPTimeseriesGroupIPVersionParamsOSWindows}),
		TLSVersion:   cloudflare.F([]radar.HTTPTimeseriesGroupIPVersionParamsTLSVersion{radar.HTTPTimeseriesGroupIPVersionParamsTLSVersionTlSv1_0}),
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
		AggInterval:  cloudflare.F(radar.HTTPTimeseriesGroupOSParamsAggInterval15m),
		ASN:          cloudflare.F([]string{"string"}),
		BotClass:     cloudflare.F([]radar.HTTPTimeseriesGroupOSParamsBotClass{radar.HTTPTimeseriesGroupOSParamsBotClassLikelyAutomated}),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		DeviceType:   cloudflare.F([]radar.HTTPTimeseriesGroupOSParamsDeviceType{radar.HTTPTimeseriesGroupOSParamsDeviceTypeDesktop}),
		Format:       cloudflare.F(radar.HTTPTimeseriesGroupOSParamsFormatJson),
		HTTPProtocol: cloudflare.F([]radar.HTTPTimeseriesGroupOSParamsHTTPProtocol{radar.HTTPTimeseriesGroupOSParamsHTTPProtocolHTTP}),
		HTTPVersion:  cloudflare.F([]radar.HTTPTimeseriesGroupOSParamsHTTPVersion{radar.HTTPTimeseriesGroupOSParamsHTTPVersionHttPv1}),
		IPVersion:    cloudflare.F([]radar.HTTPTimeseriesGroupOSParamsIPVersion{radar.HTTPTimeseriesGroupOSParamsIPVersionIPv4}),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"string"}),
		TLSVersion:   cloudflare.F([]radar.HTTPTimeseriesGroupOSParamsTLSVersion{radar.HTTPTimeseriesGroupOSParamsTLSVersionTlSv1_0}),
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
		AggInterval:  cloudflare.F(radar.HTTPTimeseriesGroupPostQuantumParamsAggInterval15m),
		ASN:          cloudflare.F([]string{"string"}),
		BotClass:     cloudflare.F([]radar.HTTPTimeseriesGroupPostQuantumParamsBotClass{radar.HTTPTimeseriesGroupPostQuantumParamsBotClassLikelyAutomated}),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		DeviceType:   cloudflare.F([]radar.HTTPTimeseriesGroupPostQuantumParamsDeviceType{radar.HTTPTimeseriesGroupPostQuantumParamsDeviceTypeDesktop}),
		Format:       cloudflare.F(radar.HTTPTimeseriesGroupPostQuantumParamsFormatJson),
		HTTPProtocol: cloudflare.F([]radar.HTTPTimeseriesGroupPostQuantumParamsHTTPProtocol{radar.HTTPTimeseriesGroupPostQuantumParamsHTTPProtocolHTTP}),
		HTTPVersion:  cloudflare.F([]radar.HTTPTimeseriesGroupPostQuantumParamsHTTPVersion{radar.HTTPTimeseriesGroupPostQuantumParamsHTTPVersionHttPv1}),
		IPVersion:    cloudflare.F([]radar.HTTPTimeseriesGroupPostQuantumParamsIPVersion{radar.HTTPTimeseriesGroupPostQuantumParamsIPVersionIPv4}),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"string"}),
		OS:           cloudflare.F([]radar.HTTPTimeseriesGroupPostQuantumParamsOS{radar.HTTPTimeseriesGroupPostQuantumParamsOSWindows}),
		TLSVersion:   cloudflare.F([]radar.HTTPTimeseriesGroupPostQuantumParamsTLSVersion{radar.HTTPTimeseriesGroupPostQuantumParamsTLSVersionTlSv1_0}),
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
		AggInterval:  cloudflare.F(radar.HTTPTimeseriesGroupTLSVersionParamsAggInterval15m),
		ASN:          cloudflare.F([]string{"string"}),
		BotClass:     cloudflare.F([]radar.HTTPTimeseriesGroupTLSVersionParamsBotClass{radar.HTTPTimeseriesGroupTLSVersionParamsBotClassLikelyAutomated}),
		Continent:    cloudflare.F([]string{"string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now()}),
		DateRange:    cloudflare.F([]string{"7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now()}),
		DeviceType:   cloudflare.F([]radar.HTTPTimeseriesGroupTLSVersionParamsDeviceType{radar.HTTPTimeseriesGroupTLSVersionParamsDeviceTypeDesktop}),
		Format:       cloudflare.F(radar.HTTPTimeseriesGroupTLSVersionParamsFormatJson),
		HTTPProtocol: cloudflare.F([]radar.HTTPTimeseriesGroupTLSVersionParamsHTTPProtocol{radar.HTTPTimeseriesGroupTLSVersionParamsHTTPProtocolHTTP}),
		HTTPVersion:  cloudflare.F([]radar.HTTPTimeseriesGroupTLSVersionParamsHTTPVersion{radar.HTTPTimeseriesGroupTLSVersionParamsHTTPVersionHttPv1}),
		IPVersion:    cloudflare.F([]radar.HTTPTimeseriesGroupTLSVersionParamsIPVersion{radar.HTTPTimeseriesGroupTLSVersionParamsIPVersionIPv4}),
		Location:     cloudflare.F([]string{"string"}),
		Name:         cloudflare.F([]string{"string"}),
		OS:           cloudflare.F([]radar.HTTPTimeseriesGroupTLSVersionParamsOS{radar.HTTPTimeseriesGroupTLSVersionParamsOSWindows}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
