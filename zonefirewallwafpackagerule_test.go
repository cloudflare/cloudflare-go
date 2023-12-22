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

func TestZoneFirewallWafPackageRuleGet(t *testing.T) {
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
	_, err := client.Zones.Firewalls.Wafs.Packages.Rules.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"a25a9a7e9c00afc1fb2e0245519d725b",
		"f939de3be84e66e757adcdcb87908023",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneFirewallWafPackageRuleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Firewalls.Wafs.Packages.Rules.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"a25a9a7e9c00afc1fb2e0245519d725b",
		"f939de3be84e66e757adcdcb87908023",
		cloudflare.ZoneFirewallWafPackageRuleUpdateParams{
			Mode: cloudflare.F(cloudflare.ZoneFirewallWafPackageRuleUpdateParamsModeOn),
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

func TestZoneFirewallWafPackageRuleWafRulesListWafRulesWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Firewalls.Wafs.Packages.Rules.WafRulesListWafRules(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"a25a9a7e9c00afc1fb2e0245519d725b",
		cloudflare.ZoneFirewallWafPackageRuleWafRulesListWafRulesParams{
			Direction: cloudflare.F(cloudflare.ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsDirectionDesc),
			Match:     cloudflare.F(cloudflare.ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsMatchAny),
			Mode:      cloudflare.F(cloudflare.ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsModeChl),
			Order:     cloudflare.F(cloudflare.ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsOrderPriority),
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
