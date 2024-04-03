// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/user"
)

func TestFirewallAccessRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.User.Firewall.AccessRules.New(context.TODO(), user.FirewallAccessRuleNewParams{
		Configuration: cloudflare.F[user.FirewallAccessRuleNewParamsConfiguration](user.FirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfiguration(user.FirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfiguration{
			Target: cloudflare.F(user.FirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfigurationTargetIP),
			Value:  cloudflare.F("198.51.100.4"),
		})),
		Mode:  cloudflare.F(user.FirewallAccessRuleNewParamsModeChallenge),
		Notes: cloudflare.F("This rule is enabled because of an event that occurred on date X."),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFirewallAccessRuleListWithOptionalParams(t *testing.T) {
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
	_, err := client.User.Firewall.AccessRules.List(context.TODO(), user.FirewallAccessRuleListParams{
		Direction: cloudflare.F(user.FirewallAccessRuleListParamsDirectionDesc),
		EgsPagination: cloudflare.F(user.FirewallAccessRuleListParamsEgsPagination{
			Json: cloudflare.F(user.FirewallAccessRuleListParamsEgsPaginationJson{
				Page:    cloudflare.F(1.000000),
				PerPage: cloudflare.F(1.000000),
			}),
		}),
		Filters: cloudflare.F(user.FirewallAccessRuleListParamsFilters{
			ConfigurationTarget: cloudflare.F(user.FirewallAccessRuleListParamsFiltersConfigurationTargetIP),
			ConfigurationValue:  cloudflare.F("198.51.100.4"),
			Match:               cloudflare.F(user.FirewallAccessRuleListParamsFiltersMatchAny),
			Mode:                cloudflare.F(user.FirewallAccessRuleListParamsFiltersModeChallenge),
			Notes:               cloudflare.F("my note"),
		}),
		Order:   cloudflare.F(user.FirewallAccessRuleListParamsOrderMode),
		Page:    cloudflare.F(1.000000),
		PerPage: cloudflare.F(20.000000),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFirewallAccessRuleDelete(t *testing.T) {
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
	_, err := client.User.Firewall.AccessRules.Delete(
		context.TODO(),
		"92f17202ed8bd63d69a66b86a49a8f6b",
		user.FirewallAccessRuleDeleteParams{
			Body: cloudflare.F[any](map[string]interface{}{}),
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

func TestFirewallAccessRuleEditWithOptionalParams(t *testing.T) {
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
	_, err := client.User.Firewall.AccessRules.Edit(
		context.TODO(),
		"92f17202ed8bd63d69a66b86a49a8f6b",
		user.FirewallAccessRuleEditParams{
			Mode:  cloudflare.F(user.FirewallAccessRuleEditParamsModeChallenge),
			Notes: cloudflare.F("This rule is enabled because of an event that occurred on date X."),
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
