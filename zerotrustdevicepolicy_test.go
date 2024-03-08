// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
)

func TestZeroTrustDevicePolicyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Policies.New(context.TODO(), cloudflare.ZeroTrustDevicePolicyNewParams{
		AccountID:           cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
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
		ServiceModeV2: cloudflare.F(cloudflare.ZeroTrustDevicePolicyNewParamsServiceModeV2{
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

func TestZeroTrustDevicePolicyList(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Policies.List(context.TODO(), cloudflare.ZeroTrustDevicePolicyListParams{
		AccountID: cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZeroTrustDevicePolicyDelete(t *testing.T) {
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
		cloudflare.ZeroTrustDevicePolicyDeleteParams{
			AccountID: cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
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

func TestZeroTrustDevicePolicyEditWithOptionalParams(t *testing.T) {
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
		cloudflare.ZeroTrustDevicePolicyEditParams{
			AccountID:           cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
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
			ServiceModeV2: cloudflare.F(cloudflare.ZeroTrustDevicePolicyEditParamsServiceModeV2{
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

func TestZeroTrustDevicePolicyGet(t *testing.T) {
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
		cloudflare.ZeroTrustDevicePolicyGetParams{
			AccountID: cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
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
