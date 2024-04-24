// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
)

func TestDevicePolicyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Policies.New(context.TODO(), zero_trust.DevicePolicyNewParams{
		AccountID:           cloudflare.F("699d98642c564d2e855e9661899b7252"),
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
		LANAllowMinutes:     cloudflare.F(30.000000),
		LANAllowSubnetSize:  cloudflare.F(24.000000),
		ServiceModeV2: cloudflare.F(zero_trust.DevicePolicyNewParamsServiceModeV2{
			Mode: cloudflare.F("proxy"),
			Port: cloudflare.F(3000.000000),
		}),
		SupportURL:   cloudflare.F("https://1.1.1.1/help"),
		SwitchLocked: cloudflare.F(true),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDevicePolicyList(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Policies.List(context.TODO(), zero_trust.DevicePolicyListParams{
		AccountID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDevicePolicyDelete(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Policies.Delete(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.DevicePolicyDeleteParams{
			AccountID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
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

func TestDevicePolicyEditWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Policies.Edit(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.DevicePolicyEditParams{
			AccountID:           cloudflare.F("699d98642c564d2e855e9661899b7252"),
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
			ServiceModeV2: cloudflare.F(zero_trust.DevicePolicyEditParamsServiceModeV2{
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

func TestDevicePolicyGet(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Policies.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.DevicePolicyGetParams{
			AccountID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
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
