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

func TestAccountRulesetGet(t *testing.T) {
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
	_, err := client.Accounts.Rulesets.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"2f2feab2026849078ba485f918791bdc",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountRulesetUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Rulesets.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"2f2feab2026849078ba485f918791bdc",
		cloudflare.AccountRulesetUpdateParams{
			Rules: cloudflare.F([]cloudflare.AccountRulesetUpdateParamsRule{cloudflare.AccountRulesetUpdateParamsRulesCreateUpdateRule(cloudflare.AccountRulesetUpdateParamsRulesCreateUpdateRule{
				Action: cloudflare.F("execute"),
				ActionParameters: cloudflare.F[any](map[string]interface{}{
					"id": "4814384a9e5d4991b9815dcfc25d2f1f",
				}),
				Description: cloudflare.F("Execute the OWASP ruleset when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				Logging: cloudflare.F(cloudflare.AccountRulesetUpdateParamsRulesCreateUpdateRuleLogging{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			}), cloudflare.AccountRulesetUpdateParamsRulesCreateUpdateRule(cloudflare.AccountRulesetUpdateParamsRulesCreateUpdateRule{
				Action: cloudflare.F("execute"),
				ActionParameters: cloudflare.F[any](map[string]interface{}{
					"id": "4814384a9e5d4991b9815dcfc25d2f1f",
				}),
				Description: cloudflare.F("Execute the OWASP ruleset when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				Logging: cloudflare.F(cloudflare.AccountRulesetUpdateParamsRulesCreateUpdateRuleLogging{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			}), cloudflare.AccountRulesetUpdateParamsRulesCreateUpdateRule(cloudflare.AccountRulesetUpdateParamsRulesCreateUpdateRule{
				Action: cloudflare.F("execute"),
				ActionParameters: cloudflare.F[any](map[string]interface{}{
					"id": "4814384a9e5d4991b9815dcfc25d2f1f",
				}),
				Description: cloudflare.F("Execute the OWASP ruleset when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				Logging: cloudflare.F(cloudflare.AccountRulesetUpdateParamsRulesCreateUpdateRuleLogging{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			})}),
			Description: cloudflare.F("My ruleset to execute managed rulesets"),
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

func TestAccountRulesetDelete(t *testing.T) {
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
	err := client.Accounts.Rulesets.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"2f2feab2026849078ba485f918791bdc",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountRulesetAccountRulesetsNewAnAccountRulesetWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Rulesets.AccountRulesetsNewAnAccountRuleset(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParams{
			Kind:  cloudflare.F(cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsKindRoot),
			Name:  cloudflare.F("My ruleset"),
			Phase: cloudflare.F("http_request_firewall_managed"),
			Rules: cloudflare.F([]cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRule{cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesCreateUpdateRule(cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesCreateUpdateRule{
				Action: cloudflare.F("execute"),
				ActionParameters: cloudflare.F[any](map[string]interface{}{
					"id": "4814384a9e5d4991b9815dcfc25d2f1f",
				}),
				Description: cloudflare.F("Execute the OWASP ruleset when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				Logging: cloudflare.F(cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesCreateUpdateRuleLogging{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			}), cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesCreateUpdateRule(cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesCreateUpdateRule{
				Action: cloudflare.F("execute"),
				ActionParameters: cloudflare.F[any](map[string]interface{}{
					"id": "4814384a9e5d4991b9815dcfc25d2f1f",
				}),
				Description: cloudflare.F("Execute the OWASP ruleset when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				Logging: cloudflare.F(cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesCreateUpdateRuleLogging{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			}), cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesCreateUpdateRule(cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesCreateUpdateRule{
				Action: cloudflare.F("execute"),
				ActionParameters: cloudflare.F[any](map[string]interface{}{
					"id": "4814384a9e5d4991b9815dcfc25d2f1f",
				}),
				Description: cloudflare.F("Execute the OWASP ruleset when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				Logging: cloudflare.F(cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesCreateUpdateRuleLogging{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			})}),
			Description: cloudflare.F("My ruleset to execute managed rulesets"),
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

func TestAccountRulesetAccountRulesetsListAccountRulesets(t *testing.T) {
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
	_, err := client.Accounts.Rulesets.AccountRulesetsListAccountRulesets(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
