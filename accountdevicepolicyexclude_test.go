// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestAccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeList(t *testing.T) {
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
	_, err := client.Accounts.Devices.Policies.Excludes.DevicesGetSplitTunnelExcludeList(context.TODO(), "699d98642c564d2e855e9661899b7252")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicy(t *testing.T) {
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
	_, err := client.Accounts.Devices.Policies.Excludes.DevicesGetSplitTunnelExcludeListForADeviceSettingsPolicy(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeList(t *testing.T) {
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
	_, err := client.Accounts.Devices.Policies.Excludes.DevicesSetSplitTunnelExcludeList(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cloudflare.AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListParams{
			Body: cloudflare.F([]cloudflare.AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListParamsBody{{
				Address:     cloudflare.F("192.0.2.0/24"),
				Description: cloudflare.F("Exclude testing domains from the tunnel"),
				Host:        cloudflare.F("*.example.com"),
			}, {
				Address:     cloudflare.F("192.0.2.0/24"),
				Description: cloudflare.F("Exclude testing domains from the tunnel"),
				Host:        cloudflare.F("*.example.com"),
			}, {
				Address:     cloudflare.F("192.0.2.0/24"),
				Description: cloudflare.F("Exclude testing domains from the tunnel"),
				Host:        cloudflare.F("*.example.com"),
			}}),
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

func TestAccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicy(t *testing.T) {
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
	_, err := client.Accounts.Devices.Policies.Excludes.DevicesSetSplitTunnelExcludeListForADeviceSettingsPolicy(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyParams{
			Body: cloudflare.F([]cloudflare.AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyParamsBody{{
				Address:     cloudflare.F("192.0.2.0/24"),
				Description: cloudflare.F("Exclude testing domains from the tunnel"),
				Host:        cloudflare.F("*.example.com"),
			}, {
				Address:     cloudflare.F("192.0.2.0/24"),
				Description: cloudflare.F("Exclude testing domains from the tunnel"),
				Host:        cloudflare.F("*.example.com"),
			}, {
				Address:     cloudflare.F("192.0.2.0/24"),
				Description: cloudflare.F("Exclude testing domains from the tunnel"),
				Host:        cloudflare.F("*.example.com"),
			}}),
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
