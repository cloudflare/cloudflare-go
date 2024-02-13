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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Firewalls.AccessRules.New(
		context.TODO(),
		"string",
		map[string]interface{}{},
		cloudflare.FirewallAccessRuleNewParams{
			Configuration: cloudflare.F[cloudflare.FirewallAccessRuleNewParamsConfiguration](cloudflare.FirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfiguration(cloudflare.FirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfiguration{
				Target: cloudflare.F(cloudflare.FirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfigurationTargetIP),
				Value:  cloudflare.F("198.51.100.4"),
			})),
			Mode:  cloudflare.F(cloudflare.FirewallAccessRuleNewParamsModeChallenge),
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

func TestFirewallAccessRuleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Firewalls.AccessRules.Update(
		context.TODO(),
		map[string]interface{}{},
		cloudflare.FirewallAccessRuleUpdateParams{
			AccountIdentifier: cloudflare.F[any](map[string]interface{}{}),
			Configuration: cloudflare.F[cloudflare.FirewallAccessRuleUpdateParamsConfiguration](cloudflare.FirewallAccessRuleUpdateParamsConfigurationLegacyJhsIPConfiguration(cloudflare.FirewallAccessRuleUpdateParamsConfigurationLegacyJhsIPConfiguration{
				Target: cloudflare.F(cloudflare.FirewallAccessRuleUpdateParamsConfigurationLegacyJhsIPConfigurationTargetIP),
				Value:  cloudflare.F("198.51.100.4"),
			})),
			Mode:  cloudflare.F(cloudflare.FirewallAccessRuleUpdateParamsModeChallenge),
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Firewalls.AccessRules.List(
		context.TODO(),
		"string",
		map[string]interface{}{},
		cloudflare.FirewallAccessRuleListParams{
			Direction: cloudflare.F(cloudflare.FirewallAccessRuleListParamsDirectionDesc),
			EgsPagination: cloudflare.F(cloudflare.FirewallAccessRuleListParamsEgsPagination{
				Json: cloudflare.F(cloudflare.FirewallAccessRuleListParamsEgsPaginationJson{
					Page:    cloudflare.F(1.000000),
					PerPage: cloudflare.F(1.000000),
				}),
			}),
			Filters: cloudflare.F(cloudflare.FirewallAccessRuleListParamsFilters{
				ConfigurationTarget: cloudflare.F(cloudflare.FirewallAccessRuleListParamsFiltersConfigurationTargetIP),
				ConfigurationValue:  cloudflare.F("198.51.100.4"),
				Match:               cloudflare.F(cloudflare.FirewallAccessRuleListParamsFiltersMatchAny),
				Mode:                cloudflare.F(cloudflare.FirewallAccessRuleListParamsFiltersModeChallenge),
				Notes:               cloudflare.F("my note"),
			}),
			Order:   cloudflare.F(cloudflare.FirewallAccessRuleListParamsOrderMode),
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Firewalls.AccessRules.Delete(
		context.TODO(),
		"string",
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

func TestFirewallAccessRuleGet(t *testing.T) {
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
	_, err := client.Firewalls.AccessRules.Get(
		context.TODO(),
		"string",
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
