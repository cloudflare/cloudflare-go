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

func TestAccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeList(t *testing.T) {
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
	_, err := client.Accounts.Devices.Policies.Includes.DevicesGetSplitTunnelIncludeList(context.TODO(), "699d98642c564d2e855e9661899b7252")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicy(t *testing.T) {
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
	_, err := client.Accounts.Devices.Policies.Includes.DevicesGetSplitTunnelIncludeListForADeviceSettingsPolicy(
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

func TestAccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeList(t *testing.T) {
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
	_, err := client.Accounts.Devices.Policies.Includes.DevicesSetSplitTunnelIncludeList(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cloudflare.AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListParams{
			Body: cloudflare.F([]cloudflare.AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListParamsBody{{
				Address:     cloudflare.F("192.0.2.0/24"),
				Description: cloudflare.F("Include testing domains from the tunnel"),
				Host:        cloudflare.F("*.example.com"),
			}, {
				Address:     cloudflare.F("192.0.2.0/24"),
				Description: cloudflare.F("Include testing domains from the tunnel"),
				Host:        cloudflare.F("*.example.com"),
			}, {
				Address:     cloudflare.F("192.0.2.0/24"),
				Description: cloudflare.F("Include testing domains from the tunnel"),
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

func TestAccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicy(t *testing.T) {
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
	_, err := client.Accounts.Devices.Policies.Includes.DevicesSetSplitTunnelIncludeListForADeviceSettingsPolicy(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyParams{
			Body: cloudflare.F([]cloudflare.AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyParamsBody{{
				Address:     cloudflare.F("192.0.2.0/24"),
				Description: cloudflare.F("Include testing domains from the tunnel"),
				Host:        cloudflare.F("*.example.com"),
			}, {
				Address:     cloudflare.F("192.0.2.0/24"),
				Description: cloudflare.F("Include testing domains from the tunnel"),
				Host:        cloudflare.F("*.example.com"),
			}, {
				Address:     cloudflare.F("192.0.2.0/24"),
				Description: cloudflare.F("Include testing domains from the tunnel"),
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
