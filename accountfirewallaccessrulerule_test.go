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

func TestAccountFirewallAccessRuleRuleGet(t *testing.T) {
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
	_, err := client.Accounts.Firewalls.AccessRules.Rules.Get(
		context.TODO(),
		map[string]interface{}{},
		map[string]interface{}{},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountFirewallAccessRuleRuleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Firewalls.AccessRules.Rules.Update(
		context.TODO(),
		map[string]interface{}{},
		map[string]interface{}{},
		cloudflare.AccountFirewallAccessRuleRuleUpdateParams{
			Configuration: cloudflare.F[cloudflare.AccountFirewallAccessRuleRuleUpdateParamsConfiguration](cloudflare.AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTIPConfiguration(cloudflare.AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTIPConfiguration{
				Target: cloudflare.F(cloudflare.AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTIPConfigurationTargetIP),
				Value:  cloudflare.F("198.51.100.4"),
			})),
			Mode:  cloudflare.F(cloudflare.AccountFirewallAccessRuleRuleUpdateParamsModeChallenge),
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

func TestAccountFirewallAccessRuleRuleDelete(t *testing.T) {
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
	_, err := client.Accounts.Firewalls.AccessRules.Rules.Delete(
		context.TODO(),
		map[string]interface{}{},
		map[string]interface{}{},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Firewalls.AccessRules.Rules.IPAccessRulesForAnAccountNewAnIPAccessRule(
		context.TODO(),
		map[string]interface{}{},
		cloudflare.AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParams{
			Configuration: cloudflare.F[cloudflare.AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfiguration](cloudflare.AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTIPConfiguration(cloudflare.AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTIPConfiguration{
				Target: cloudflare.F(cloudflare.AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTIPConfigurationTargetIP),
				Value:  cloudflare.F("198.51.100.4"),
			})),
			Mode:  cloudflare.F(cloudflare.AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsModeChallenge),
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

func TestAccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Firewalls.AccessRules.Rules.IPAccessRulesForAnAccountListIPAccessRules(
		context.TODO(),
		map[string]interface{}{},
		cloudflare.AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParams{
			Direction: cloudflare.F(cloudflare.AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsDirectionDesc),
			EgsPagination: cloudflare.F(cloudflare.AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsEgsPagination{
				Json: cloudflare.F(cloudflare.AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsEgsPaginationJson{
					Page:    cloudflare.F(1.000000),
					PerPage: cloudflare.F(1.000000),
				}),
			}),
			Filters: cloudflare.F(cloudflare.AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFilters{
				ConfigurationTarget: cloudflare.F(cloudflare.AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFiltersConfigurationTargetIP),
				ConfigurationValue:  cloudflare.F("198.51.100.4"),
				Match:               cloudflare.F(cloudflare.AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFiltersMatchAny),
				Mode:                cloudflare.F(cloudflare.AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFiltersModeChallenge),
				Notes:               cloudflare.F("my note"),
			}),
			Order:   cloudflare.F(cloudflare.AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsOrderMode),
			Page:    cloudflare.F(1.000000),
			PerPage: cloudflare.F(20.000000),
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
