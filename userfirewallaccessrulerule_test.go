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

func TestUserFirewallAccessRuleRuleNewWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Users.Firewall.AccessRules.Rules.New(context.TODO(), cloudflare.UserFirewallAccessRuleRuleNewParams{
		Configuration: cloudflare.F[cloudflare.UserFirewallAccessRuleRuleNewParamsConfiguration](cloudflare.UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsIPConfiguration(cloudflare.UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsIPConfiguration{
			Target: cloudflare.F(cloudflare.UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsIPConfigurationTargetIP),
			Value:  cloudflare.F("198.51.100.4"),
		})),
		Mode:  cloudflare.F(cloudflare.UserFirewallAccessRuleRuleNewParamsModeChallenge),
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Users.Firewall.AccessRules.Rules.Update(
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

func TestUserFirewallAccessRuleRuleListWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Users.Firewall.AccessRules.Rules.List(context.TODO(), cloudflare.UserFirewallAccessRuleRuleListParams{
		Direction: cloudflare.F(cloudflare.UserFirewallAccessRuleRuleListParamsDirectionDesc),
		EgsPagination: cloudflare.F(cloudflare.UserFirewallAccessRuleRuleListParamsEgsPagination{
			Json: cloudflare.F(cloudflare.UserFirewallAccessRuleRuleListParamsEgsPaginationJson{
				Page:    cloudflare.F(1.000000),
				PerPage: cloudflare.F(1.000000),
			}),
		}),
		Filters: cloudflare.F(cloudflare.UserFirewallAccessRuleRuleListParamsFilters{
			ConfigurationTarget: cloudflare.F(cloudflare.UserFirewallAccessRuleRuleListParamsFiltersConfigurationTargetIP),
			ConfigurationValue:  cloudflare.F("198.51.100.4"),
			Match:               cloudflare.F(cloudflare.UserFirewallAccessRuleRuleListParamsFiltersMatchAny),
			Mode:                cloudflare.F(cloudflare.UserFirewallAccessRuleRuleListParamsFiltersModeChallenge),
			Notes:               cloudflare.F("my note"),
		}),
		Order:   cloudflare.F(cloudflare.UserFirewallAccessRuleRuleListParamsOrderMode),
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Users.Firewall.AccessRules.Rules.Delete(context.TODO(), "92f17202ed8bd63d69a66b86a49a8f6b")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
