// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/zero_trust"
)

func TestDevicePolicyDefaultEditWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Policies.Default.Edit(context.TODO(), zero_trust.DevicePolicyDefaultEditParams{
		AccountID:           cloudflare.F("699d98642c564d2e855e9661899b7252"),
		AllowModeSwitch:     cloudflare.F(true),
		AllowUpdates:        cloudflare.F(true),
		AllowedToLeave:      cloudflare.F(true),
		AutoConnect:         cloudflare.F(0.000000),
		CaptivePortal:       cloudflare.F(180.000000),
		DisableAutoFallback: cloudflare.F(true),
		Exclude: cloudflare.F([]zero_trust.SplitTunnelExcludeUnionParam{zero_trust.SplitTunnelExcludeTeamsDevicesExcludeSplitTunnelWithAddressParam{
			Address:     cloudflare.F("192.0.2.0/24"),
			Description: cloudflare.F("Exclude testing domains from the tunnel"),
		}}),
		ExcludeOfficeIPs: cloudflare.F(true),
		Include: cloudflare.F([]zero_trust.SplitTunnelIncludeUnionParam{zero_trust.SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddressParam{
			Address:     cloudflare.F("192.0.2.0/24"),
			Description: cloudflare.F("Include testing domains in the tunnel"),
		}}),
		LANAllowMinutes:            cloudflare.F(30.000000),
		LANAllowSubnetSize:         cloudflare.F(24.000000),
		RegisterInterfaceIPWithDNS: cloudflare.F(true),
		SccmVpnBoundarySupport:     cloudflare.F(false),
		ServiceModeV2: cloudflare.F(zero_trust.DevicePolicyDefaultEditParamsServiceModeV2{
			Mode: cloudflare.F("proxy"),
			Port: cloudflare.F(3000.000000),
		}),
		SupportURL:     cloudflare.F("https://1.1.1.1/help"),
		SwitchLocked:   cloudflare.F(true),
		TunnelProtocol: cloudflare.F("wireguard"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDevicePolicyDefaultGet(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Policies.Default.Get(context.TODO(), zero_trust.DevicePolicyDefaultGetParams{
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
