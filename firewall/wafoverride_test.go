// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/firewall"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

func TestWAFOverrideNew(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.Firewall.WAF.Overrides.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		firewall.WAFOverrideNewParams{
			URLs: cloudflare.F([]firewall.OverrideURLParam{"shop.example.com/*", "shop.example.com/*", "shop.example.com/*"}),
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

func TestWAFOverrideUpdateWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.Firewall.WAF.Overrides.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		firewall.WAFOverrideUpdateParams{
			PathID: cloudflare.F("de677e5818985db1285d0e80225f06e5"),
			BodyID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			RewriteAction: cloudflare.F(firewall.RewriteActionParam{
				Block:     cloudflare.F(firewall.RewriteActionBlockChallenge),
				Challenge: cloudflare.F(firewall.RewriteActionChallengeChallenge),
				Default:   cloudflare.F(firewall.RewriteActionDefaultChallenge),
				Disable:   cloudflare.F(firewall.RewriteActionDisableChallenge),
				Simulate:  cloudflare.F(firewall.RewriteActionSimulateChallenge),
			}),
			Rules: cloudflare.F(firewall.WAFRuleParam{
				"100015": firewall.WAFRuleItemDisable,
			}),
			URLs: cloudflare.F([]firewall.OverrideURLParam{"shop.example.com/*", "shop.example.com/*", "shop.example.com/*"}),
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

func TestWAFOverrideListWithOptionalParams(t *testing.T) {
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
	_, err := client.Firewall.WAF.Overrides.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		firewall.WAFOverrideListParams{
			Page:    cloudflare.F(1.000000),
			PerPage: cloudflare.F(5.000000),
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

func TestWAFOverrideDelete(t *testing.T) {
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
	_, err := client.Firewall.WAF.Overrides.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
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

func TestWAFOverrideGet(t *testing.T) {
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
	_, err := client.Firewall.WAF.Overrides.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
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
