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

func TestUserFirewallAccessRuleRuleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Firewalls.AccessRules.Rules.Update(
		context.TODO(),
		"92f17202ed8bd63d69a66b86a49a8f6b",
		cloudflare.UserFirewallAccessRuleRuleUpdateParams{
			Mode:  cloudflare.F(cloudflare.UserFirewallAccessRuleRuleUpdateParamsModeChallenge),
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

func TestUserFirewallAccessRuleRuleDelete(t *testing.T) {
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
	_, err := client.Users.Firewalls.AccessRules.Rules.Delete(context.TODO(), "92f17202ed8bd63d69a66b86a49a8f6b")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Firewalls.AccessRules.Rules.IPAccessRulesForAUserNewAnIPAccessRule(context.TODO(), cloudflare.UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParams{
		Configuration: cloudflare.F[cloudflare.UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfiguration](cloudflare.UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationIPConfiguration(cloudflare.UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationIPConfiguration{
			Target: cloudflare.F(cloudflare.UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationIPConfigurationTargetIP),
			Value:  cloudflare.F("198.51.100.4"),
		})),
		Mode:  cloudflare.F(cloudflare.UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsModeChallenge),
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

func TestUserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Firewalls.AccessRules.Rules.IPAccessRulesForAUserListIPAccessRules(context.TODO(), cloudflare.UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParams{
		Direction: cloudflare.F(cloudflare.UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsDirectionDesc),
		EgsPagination: cloudflare.F(cloudflare.UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsEgsPagination{
			Json: cloudflare.F(cloudflare.UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsEgsPaginationJson{
				Page:    cloudflare.F(1.000000),
				PerPage: cloudflare.F(1.000000),
			}),
		}),
		Filters: cloudflare.F(cloudflare.UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFilters{
			ConfigurationTarget: cloudflare.F(cloudflare.UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersConfigurationTargetIP),
			ConfigurationValue:  cloudflare.F("198.51.100.4"),
			Match:               cloudflare.F(cloudflare.UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersMatchAny),
			Mode:                cloudflare.F(cloudflare.UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersModeChallenge),
			Notes:               cloudflare.F("my note"),
		}),
		Order:   cloudflare.F(cloudflare.UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsOrderMode),
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
