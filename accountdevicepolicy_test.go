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

func TestAccountDevicePolicyGet(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Devices.Policies.Get(
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

func TestAccountDevicePolicyUpdateWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Devices.Policies.Update(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.AccountDevicePolicyUpdateParams{
			AllowModeSwitch:     cloudflare.F(true),
			AllowUpdates:        cloudflare.F(true),
			AllowedToLeave:      cloudflare.F(true),
			AutoConnect:         cloudflare.F(0.000000),
			CaptivePortal:       cloudflare.F(180.000000),
			Description:         cloudflare.F("Policy for test teams."),
			DisableAutoFallback: cloudflare.F(true),
			Enabled:             cloudflare.F(true),
			ExcludeOfficeIPs:    cloudflare.F(true),
			Match:               cloudflare.F("user.identity == \"test@cloudflare.com\""),
			Name:                cloudflare.F("Allow Developers"),
			Precedence:          cloudflare.F(100.000000),
			ServiceModeV2: cloudflare.F(cloudflare.AccountDevicePolicyUpdateParamsServiceModeV2{
				Mode: cloudflare.F("proxy"),
				Port: cloudflare.F(3000.000000),
			}),
			SupportURL:   cloudflare.F("https://1.1.1.1/help"),
			SwitchLocked: cloudflare.F(true),
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

func TestAccountDevicePolicyDelete(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Devices.Policies.Delete(
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

func TestAccountDevicePolicyDevicesNewDeviceSettingsPolicyWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Devices.Policies.DevicesNewDeviceSettingsPolicy(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cloudflare.AccountDevicePolicyDevicesNewDeviceSettingsPolicyParams{
			Match:               cloudflare.F("user.identity == \"test@cloudflare.com\""),
			Name:                cloudflare.F("Allow Developers"),
			Precedence:          cloudflare.F(100.000000),
			AllowModeSwitch:     cloudflare.F(true),
			AllowUpdates:        cloudflare.F(true),
			AllowedToLeave:      cloudflare.F(true),
			AutoConnect:         cloudflare.F(0.000000),
			CaptivePortal:       cloudflare.F(180.000000),
			Description:         cloudflare.F("Policy for test teams."),
			DisableAutoFallback: cloudflare.F(true),
			Enabled:             cloudflare.F(true),
			ExcludeOfficeIPs:    cloudflare.F(true),
			LanAllowMinutes:     cloudflare.F(30.000000),
			LanAllowSubnetSize:  cloudflare.F(24.000000),
			ServiceModeV2: cloudflare.F(cloudflare.AccountDevicePolicyDevicesNewDeviceSettingsPolicyParamsServiceModeV2{
				Mode: cloudflare.F("proxy"),
				Port: cloudflare.F(3000.000000),
			}),
			SupportURL:   cloudflare.F("https://1.1.1.1/help"),
			SwitchLocked: cloudflare.F(true),
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

func TestAccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicy(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Devices.Policies.DevicesGetDefaultDeviceSettingsPolicy(context.TODO(), "699d98642c564d2e855e9661899b7252")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountDevicePolicyDevicesListDeviceSettingsPolicies(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Devices.Policies.DevicesListDeviceSettingsPolicies(context.TODO(), "699d98642c564d2e855e9661899b7252")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Devices.Policies.DevicesUpdateDefaultDeviceSettingsPolicy(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cloudflare.AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyParams{
			AllowModeSwitch:     cloudflare.F(true),
			AllowUpdates:        cloudflare.F(true),
			AllowedToLeave:      cloudflare.F(true),
			AutoConnect:         cloudflare.F(0.000000),
			CaptivePortal:       cloudflare.F(180.000000),
			DisableAutoFallback: cloudflare.F(true),
			ExcludeOfficeIPs:    cloudflare.F(true),
			ServiceModeV2: cloudflare.F(cloudflare.AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyParamsServiceModeV2{
				Mode: cloudflare.F("proxy"),
				Port: cloudflare.F(3000.000000),
			}),
			SupportURL:   cloudflare.F("https://1.1.1.1/help"),
			SwitchLocked: cloudflare.F(true),
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
