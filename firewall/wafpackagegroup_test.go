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

func TestWAFPackageGroupListWithOptionalParams(t *testing.T) {
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
		firewall.WAFPackageGroupListParams{
			ZoneID:     cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Direction:  cloudflare.F(firewall.WAFPackageGroupListParamsDirectionAsc),
			Match:      cloudflare.F(firewall.WAFPackageGroupListParamsMatchAny),
			Mode:       cloudflare.F(firewall.WAFPackageGroupListParamsModeOn),
			Name:       cloudflare.F("Project Honey Pot"),
			Order:      cloudflare.F(firewall.WAFPackageGroupListParamsOrderMode),
			Page:       cloudflare.F(1.000000),
			PerPage:    cloudflare.F(5.000000),
			RulesCount: cloudflare.F(10.000000),
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

func TestWAFPackageGroupEditWithOptionalParams(t *testing.T) {
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
		firewall.WAFPackageGroupEditParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Mode:   cloudflare.F(firewall.WAFPackageGroupEditParamsModeOn),
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

func TestWAFPackageGroupGet(t *testing.T) {
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
		firewall.WAFPackageGroupGetParams{
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
