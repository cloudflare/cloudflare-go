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

func TestFirewallWAFPackageGroupListWithOptionalParams(t *testing.T) {
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
	_, err := client.Firewall.WAF.Packages.Groups.List(
		context.TODO(),
		"a25a9a7e9c00afc1fb2e0245519d725b",
		cloudflare.FirewallWAFPackageGroupListParams{
			ZoneID:    cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Direction: cloudflare.F(cloudflare.FirewallWAFPackageGroupListParamsDirectionDesc),
			Match:     cloudflare.F(cloudflare.FirewallWAFPackageGroupListParamsMatchAny),
			Mode:      cloudflare.F(cloudflare.FirewallWAFPackageGroupListParamsModeOn),
			Order:     cloudflare.F(cloudflare.FirewallWAFPackageGroupListParamsOrderMode),
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

func TestFirewallWAFPackageGroupEditWithOptionalParams(t *testing.T) {
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
	_, err := client.Firewall.WAF.Packages.Groups.Edit(
		context.TODO(),
		"a25a9a7e9c00afc1fb2e0245519d725b",
		"a25a9a7e9c00afc1fb2e0245519d725b",
		cloudflare.FirewallWAFPackageGroupEditParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Mode:   cloudflare.F(cloudflare.FirewallWAFPackageGroupEditParamsModeOn),
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

func TestFirewallWAFPackageGroupGet(t *testing.T) {
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
	_, err := client.Firewall.WAF.Packages.Groups.Get(
		context.TODO(),
		"a25a9a7e9c00afc1fb2e0245519d725b",
		"a25a9a7e9c00afc1fb2e0245519d725b",
		cloudflare.FirewallWAFPackageGroupGetParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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
