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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Firewall.AccessRules.New(context.TODO(), cloudflare.FirewallAccessRuleNewParams{
		Configuration: cloudflare.F[cloudflare.FirewallAccessRuleNewParamsConfiguration](cloudflare.FirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfiguration(cloudflare.FirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfiguration{
			Target: cloudflare.F(cloudflare.FirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfigurationTargetIP),
			Value:  cloudflare.F("198.51.100.4"),
		})),
		Mode:      cloudflare.F(cloudflare.FirewallAccessRuleNewParamsModeChallenge),
		AccountID: cloudflare.F("string"),
		ZoneID:    cloudflare.F("string"),
		Notes:     cloudflare.F("This rule is enabled because of an event that occurred on date X."),
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
	_, err := client.Firewall.AccessRules.List(context.TODO(), cloudflare.FirewallAccessRuleListParams{
		AccountID: cloudflare.F("string"),
		ZoneID:    cloudflare.F("string"),
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
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFirewallAccessRuleDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Firewall.AccessRules.Delete(
		context.TODO(),
		map[string]interface{}{},
		cloudflare.FirewallAccessRuleDeleteParams{
			AccountID: cloudflare.F("string"),
			ZoneID:    cloudflare.F("string"),
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
	_, err := client.Firewall.AccessRules.Edit(
		context.TODO(),
		map[string]interface{}{},
		cloudflare.FirewallAccessRuleEditParams{
			Configuration: cloudflare.F[cloudflare.FirewallAccessRuleEditParamsConfiguration](cloudflare.FirewallAccessRuleEditParamsConfigurationLegacyJhsIPConfiguration(cloudflare.FirewallAccessRuleEditParamsConfigurationLegacyJhsIPConfiguration{
				Target: cloudflare.F(cloudflare.FirewallAccessRuleEditParamsConfigurationLegacyJhsIPConfigurationTargetIP),
				Value:  cloudflare.F("198.51.100.4"),
			})),
			Mode:      cloudflare.F(cloudflare.FirewallAccessRuleEditParamsModeChallenge),
			AccountID: cloudflare.F("string"),
			ZoneID:    cloudflare.F("string"),
			Notes:     cloudflare.F("This rule is enabled because of an event that occurred on date X."),
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

func TestFirewallAccessRuleGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Firewall.AccessRules.Get(
		context.TODO(),
		map[string]interface{}{},
		cloudflare.FirewallAccessRuleGetParams{
			AccountID: cloudflare.F("string"),
			ZoneID:    cloudflare.F("string"),
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
