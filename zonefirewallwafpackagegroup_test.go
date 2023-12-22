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

func TestZoneFirewallWafPackageGroupGet(t *testing.T) {
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
	_, err := client.Zones.Firewalls.Wafs.Packages.Groups.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"a25a9a7e9c00afc1fb2e0245519d725b",
		"de677e5818985db1285d0e80225f06e5",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneFirewallWafPackageGroupUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Firewalls.Wafs.Packages.Groups.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"a25a9a7e9c00afc1fb2e0245519d725b",
		"de677e5818985db1285d0e80225f06e5",
		cloudflare.ZoneFirewallWafPackageGroupUpdateParams{
			Mode: cloudflare.F(cloudflare.ZoneFirewallWafPackageGroupUpdateParamsModeOn),
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

func TestZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Firewalls.Wafs.Packages.Groups.WafRuleGroupsListWafRuleGroups(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"a25a9a7e9c00afc1fb2e0245519d725b",
		cloudflare.ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParams{
			Direction: cloudflare.F(cloudflare.ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParamsDirectionDesc),
			Match:     cloudflare.F(cloudflare.ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParamsMatchAny),
			Mode:      cloudflare.F(cloudflare.ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParamsModeOn),
			Order:     cloudflare.F(cloudflare.ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParamsOrderMode),
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
