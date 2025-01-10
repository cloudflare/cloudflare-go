// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/firewall"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func TestWAFPackageRuleListWithOptionalParams(t *testing.T) {
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
	_, err := client.Firewall.WAF.Packages.Rules.List(
		context.TODO(),
		"a25a9a7e9c00afc1fb2e0245519d725b",
		firewall.WAFPackageRuleListParams{
			ZoneID:      cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Description: cloudflare.F("SQL injection prevention for SELECT statements"),
			Direction:   cloudflare.F(firewall.WAFPackageRuleListParamsDirectionAsc),
			GroupID:     cloudflare.F("de677e5818985db1285d0e80225f06e5"),
			Match:       cloudflare.F(firewall.WAFPackageRuleListParamsMatchAny),
			Mode:        cloudflare.F(firewall.WAFPackageRuleListParamsModeDis),
			Order:       cloudflare.F(firewall.WAFPackageRuleListParamsOrderPriority),
			Page:        cloudflare.F(1.000000),
			PerPage:     cloudflare.F(5.000000),
			Priority:    cloudflare.F("priority"),
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

func TestWAFPackageRuleEditWithOptionalParams(t *testing.T) {
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
	_, err := client.Firewall.WAF.Packages.Rules.Edit(
		context.TODO(),
		"a25a9a7e9c00afc1fb2e0245519d725b",
		"a25a9a7e9c00afc1fb2e0245519d725b",
		firewall.WAFPackageRuleEditParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Mode:   cloudflare.F(firewall.WAFPackageRuleEditParamsModeDefault),
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

func TestWAFPackageRuleGet(t *testing.T) {
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
	_, err := client.Firewall.WAF.Packages.Rules.Get(
		context.TODO(),
		"a25a9a7e9c00afc1fb2e0245519d725b",
		"a25a9a7e9c00afc1fb2e0245519d725b",
		firewall.WAFPackageRuleGetParams{
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
