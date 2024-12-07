// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_network_monitoring_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/magic_network_monitoring"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func TestConfigNewWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicNetworkMonitoring.Configs.New(context.TODO(), magic_network_monitoring.ConfigNewParams{
		AccountID:       cloudflare.F("6f91088a406011ed95aed352566e8d4c"),
		DefaultSampling: cloudflare.F(1.000000),
		Name:            cloudflare.F("cloudflare user's account"),
		RouterIPs:       cloudflare.F([]string{"203.0.113.1"}),
		WARPDevices: cloudflare.F([]magic_network_monitoring.ConfigNewParamsWARPDevice{{
			ID:       cloudflare.F("5360368d-b351-4791-abe1-93550dabd351"),
			Name:     cloudflare.F("My warp device"),
			RouterIP: cloudflare.F("203.0.113.1"),
		}}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConfigUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicNetworkMonitoring.Configs.Update(context.TODO(), magic_network_monitoring.ConfigUpdateParams{
		AccountID:       cloudflare.F("6f91088a406011ed95aed352566e8d4c"),
		DefaultSampling: cloudflare.F(1.000000),
		Name:            cloudflare.F("cloudflare user's account"),
		RouterIPs:       cloudflare.F([]string{"203.0.113.1"}),
		WARPDevices: cloudflare.F([]magic_network_monitoring.ConfigUpdateParamsWARPDevice{{
			ID:       cloudflare.F("5360368d-b351-4791-abe1-93550dabd351"),
			Name:     cloudflare.F("My warp device"),
			RouterIP: cloudflare.F("203.0.113.1"),
		}}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConfigDelete(t *testing.T) {
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
	_, err := client.MagicNetworkMonitoring.Configs.Delete(context.TODO(), magic_network_monitoring.ConfigDeleteParams{
		AccountID: cloudflare.F("6f91088a406011ed95aed352566e8d4c"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConfigEditWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicNetworkMonitoring.Configs.Edit(context.TODO(), magic_network_monitoring.ConfigEditParams{
		AccountID:       cloudflare.F("6f91088a406011ed95aed352566e8d4c"),
		DefaultSampling: cloudflare.F(1.000000),
		Name:            cloudflare.F("cloudflare user's account"),
		RouterIPs:       cloudflare.F([]string{"203.0.113.1"}),
		WARPDevices: cloudflare.F([]magic_network_monitoring.ConfigEditParamsWARPDevice{{
			ID:       cloudflare.F("5360368d-b351-4791-abe1-93550dabd351"),
			Name:     cloudflare.F("My warp device"),
			RouterIP: cloudflare.F("203.0.113.1"),
		}}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConfigGet(t *testing.T) {
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
	_, err := client.MagicNetworkMonitoring.Configs.Get(context.TODO(), magic_network_monitoring.ConfigGetParams{
		AccountID: cloudflare.F("6f91088a406011ed95aed352566e8d4c"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
