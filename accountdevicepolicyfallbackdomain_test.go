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

func TestAccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackList(t *testing.T) {
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
	_, err := client.Accounts.Devices.Policies.FallbackDomains.DevicesGetLocalDomainFallbackList(context.TODO(), "699d98642c564d2e855e9661899b7252")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicy(t *testing.T) {
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
	_, err := client.Accounts.Devices.Policies.FallbackDomains.DevicesGetLocalDomainFallbackListForADeviceSettingsPolicy(
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

func TestAccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackList(t *testing.T) {
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
	_, err := client.Accounts.Devices.Policies.FallbackDomains.DevicesSetLocalDomainFallbackList(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cloudflare.AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListParams{
			Body: cloudflare.F([]cloudflare.AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListParamsBody{{
				Description: cloudflare.F("Domain bypass for local development"),
				DNSServer:   cloudflare.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
				Suffix:      cloudflare.F("example.com"),
			}, {
				Description: cloudflare.F("Domain bypass for local development"),
				DNSServer:   cloudflare.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
				Suffix:      cloudflare.F("example.com"),
			}, {
				Description: cloudflare.F("Domain bypass for local development"),
				DNSServer:   cloudflare.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
				Suffix:      cloudflare.F("example.com"),
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

func TestAccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicy(t *testing.T) {
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
	_, err := client.Accounts.Devices.Policies.FallbackDomains.DevicesSetLocalDomainFallbackListForADeviceSettingsPolicy(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyParams{
			Body: cloudflare.F([]cloudflare.AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyParamsBody{{
				Description: cloudflare.F("Domain bypass for local development"),
				DNSServer:   cloudflare.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
				Suffix:      cloudflare.F("example.com"),
			}, {
				Description: cloudflare.F("Domain bypass for local development"),
				DNSServer:   cloudflare.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
				Suffix:      cloudflare.F("example.com"),
			}, {
				Description: cloudflare.F("Domain bypass for local development"),
				DNSServer:   cloudflare.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
				Suffix:      cloudflare.F("example.com"),
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
