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

func TestUserFirewallAccessRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.User.Firewall.AccessRules.New(context.TODO(), cloudflare.UserFirewallAccessRuleNewParams{
		Configuration: cloudflare.F[cloudflare.UserFirewallAccessRuleNewParamsConfiguration](cloudflare.UserFirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfiguration(cloudflare.UserFirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfiguration{
			Target: cloudflare.F(cloudflare.UserFirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfigurationTargetIP),
			Value:  cloudflare.F("198.51.100.4"),
		})),
		Mode:  cloudflare.F(cloudflare.UserFirewallAccessRuleNewParamsModeChallenge),
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

func TestUserFirewallAccessRuleListWithOptionalParams(t *testing.T) {
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
	_, err := client.User.Firewall.AccessRules.List(context.TODO(), cloudflare.UserFirewallAccessRuleListParams{
		Direction: cloudflare.F(cloudflare.UserFirewallAccessRuleListParamsDirectionDesc),
		EgsPagination: cloudflare.F(cloudflare.UserFirewallAccessRuleListParamsEgsPagination{
			Json: cloudflare.F(cloudflare.UserFirewallAccessRuleListParamsEgsPaginationJson{
				Page:    cloudflare.F(1.000000),
				PerPage: cloudflare.F(1.000000),
			}),
		}),
		Filters: cloudflare.F(cloudflare.UserFirewallAccessRuleListParamsFilters{
			ConfigurationTarget: cloudflare.F(cloudflare.UserFirewallAccessRuleListParamsFiltersConfigurationTargetIP),
			ConfigurationValue:  cloudflare.F("198.51.100.4"),
			Match:               cloudflare.F(cloudflare.UserFirewallAccessRuleListParamsFiltersMatchAny),
			Mode:                cloudflare.F(cloudflare.UserFirewallAccessRuleListParamsFiltersModeChallenge),
			Notes:               cloudflare.F("my note"),
		}),
		Order:   cloudflare.F(cloudflare.UserFirewallAccessRuleListParamsOrderMode),
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

func TestUserFirewallAccessRuleDelete(t *testing.T) {
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
	_, err := client.User.Firewall.AccessRules.Delete(context.TODO(), "92f17202ed8bd63d69a66b86a49a8f6b")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserFirewallAccessRuleEditWithOptionalParams(t *testing.T) {
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
		cloudflare.UserFirewallAccessRuleEditParams{
			Mode:  cloudflare.F(cloudflare.UserFirewallAccessRuleEditParamsModeChallenge),
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
