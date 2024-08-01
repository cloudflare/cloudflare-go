// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/filters"
	"github.com/cloudflare/cloudflare-go/v2/firewall"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

func TestRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Firewall.Rules.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		firewall.RuleNewParams{
			Action: cloudflare.F(firewall.RuleNewParamsAction{
				Mode: cloudflare.F(firewall.RuleNewParamsActionModeChallenge),
				Response: cloudflare.F(firewall.RuleNewParamsActionResponse{
					Body:        cloudflare.F("<error>This request has been rate-limited.</error>"),
					ContentType: cloudflare.F("text/xml"),
				}),
				Timeout: cloudflare.F(86400.000000),
			}),
			Filter: cloudflare.F(filters.FirewallFilterParam{
				Description: cloudflare.F("Restrict access from these browsers on this address range."),
				Expression:  cloudflare.F("(http.request.uri.path ~ \".*wp-login.php\" or http.request.uri.path ~ \".*xmlrpc.php\") and ip.addr ne 172.16.22.155"),
				Paused:      cloudflare.F(false),
				Ref:         cloudflare.F("FIL-100"),
			}),
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

func TestRuleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Firewall.Rules.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"372e67954025e0ba6aaa6d586b9e0b60",
		firewall.RuleUpdateParams{
			Action: cloudflare.F(firewall.RuleUpdateParamsAction{
				Mode: cloudflare.F(firewall.RuleUpdateParamsActionModeChallenge),
				Response: cloudflare.F(firewall.RuleUpdateParamsActionResponse{
					Body:        cloudflare.F("<error>This request has been rate-limited.</error>"),
					ContentType: cloudflare.F("text/xml"),
				}),
				Timeout: cloudflare.F(86400.000000),
			}),
			Filter: cloudflare.F(filters.FirewallFilterParam{
				Description: cloudflare.F("Restrict access from these browsers on this address range."),
				Expression:  cloudflare.F("(http.request.uri.path ~ \".*wp-login.php\" or http.request.uri.path ~ \".*xmlrpc.php\") and ip.addr ne 172.16.22.155"),
				Paused:      cloudflare.F(false),
				Ref:         cloudflare.F("FIL-100"),
			}),
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

func TestRuleListWithOptionalParams(t *testing.T) {
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
	_, err := client.Firewall.Rules.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		firewall.RuleListParams{
			ID:          cloudflare.F("372e67954025e0ba6aaa6d586b9e0b60"),
			Action:      cloudflare.F("block"),
			Description: cloudflare.F("mir"),
			Page:        cloudflare.F(1.000000),
			Paused:      cloudflare.F(false),
			PerPage:     cloudflare.F(5.000000),
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

func TestRuleDelete(t *testing.T) {
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
	_, err := client.Firewall.Rules.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"372e67954025e0ba6aaa6d586b9e0b60",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRuleEdit(t *testing.T) {
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
	_, err := client.Firewall.Rules.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"372e67954025e0ba6aaa6d586b9e0b60",
		firewall.RuleEditParams{},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRuleGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Firewall.Rules.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		firewall.RuleGetParams{
			PathID:  cloudflare.F("372e67954025e0ba6aaa6d586b9e0b60"),
			QueryID: cloudflare.F("372e67954025e0ba6aaa6d586b9e0b60"),
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
