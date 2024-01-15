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

func TestZoneFirewallWafPackageGet(t *testing.T) {
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
	_, err := client.Zones.Firewalls.Wafs.Packages.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"a25a9a7e9c00afc1fb2e0245519d725b",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneFirewallWafPackageUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Firewalls.Wafs.Packages.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"a25a9a7e9c00afc1fb2e0245519d725b",
		cloudflare.ZoneFirewallWafPackageUpdateParams{
			ActionMode:  cloudflare.F(cloudflare.ZoneFirewallWafPackageUpdateParamsActionModeSimulate),
			Sensitivity: cloudflare.F(cloudflare.ZoneFirewallWafPackageUpdateParamsSensitivityHigh),
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

func TestZoneFirewallWafPackageWafPackagesListWafPackagesWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Firewalls.Wafs.Packages.WafPackagesListWafPackages(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZoneFirewallWafPackageWafPackagesListWafPackagesParams{
			Direction: cloudflare.F(cloudflare.ZoneFirewallWafPackageWafPackagesListWafPackagesParamsDirectionDesc),
			Match:     cloudflare.F(cloudflare.ZoneFirewallWafPackageWafPackagesListWafPackagesParamsMatchAny),
			Order:     cloudflare.F(cloudflare.ZoneFirewallWafPackageWafPackagesListWafPackagesParamsOrderName),
			Page:      cloudflare.F(1.000000),
			PerPage:   cloudflare.F(5.000000),
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
