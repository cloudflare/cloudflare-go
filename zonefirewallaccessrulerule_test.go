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

func TestZoneFirewallAccessRuleRuleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Firewalls.AccessRules.Rules.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"92f17202ed8bd63d69a66b86a49a8f6b",
		cloudflare.ZoneFirewallAccessRuleRuleUpdateParams{
			Mode:  cloudflare.F(cloudflare.ZoneFirewallAccessRuleRuleUpdateParamsModeChallenge),
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

func TestZoneFirewallAccessRuleRuleDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Firewalls.AccessRules.Rules.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"92f17202ed8bd63d69a66b86a49a8f6b",
		cloudflare.ZoneFirewallAccessRuleRuleDeleteParams{
			Cascade: cloudflare.F(cloudflare.ZoneFirewallAccessRuleRuleDeleteParamsCascadeNone),
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

func TestZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Firewalls.AccessRules.Rules.IPAccessRulesForAZoneNewAnIPAccessRule(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParams{
			Configuration: cloudflare.F[cloudflare.ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfiguration](cloudflare.ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationIPConfiguration(cloudflare.ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationIPConfiguration{
				Target: cloudflare.F(cloudflare.ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationIPConfigurationTargetIP),
				Value:  cloudflare.F("198.51.100.4"),
			})),
			Mode:  cloudflare.F(cloudflare.ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsModeChallenge),
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

func TestZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Firewalls.AccessRules.Rules.IPAccessRulesForAZoneListIPAccessRules(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParams{
			Direction: cloudflare.F(cloudflare.ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsDirectionDesc),
			EgsPagination: cloudflare.F(cloudflare.ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsEgsPagination{
				Json: cloudflare.F(cloudflare.ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsEgsPaginationJson{
					Page:    cloudflare.F(1.000000),
					PerPage: cloudflare.F(1.000000),
				}),
			}),
			Filters: cloudflare.F(cloudflare.ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFilters{
				ConfigurationTarget: cloudflare.F(cloudflare.ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFiltersConfigurationTargetIP),
				ConfigurationValue:  cloudflare.F("198.51.100.4"),
				Match:               cloudflare.F(cloudflare.ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFiltersMatchAny),
				Mode:                cloudflare.F(cloudflare.ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFiltersModeChallenge),
				Notes:               cloudflare.F("my note"),
			}),
			Order:   cloudflare.F(cloudflare.ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsOrderMode),
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
