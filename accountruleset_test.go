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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Rulesets.Get(
		context.TODO(),
		"abf9b32d38c5f572afde3336ec0ce302",
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Rulesets.Update(
		context.TODO(),
		"abf9b32d38c5f572afde3336ec0ce302",
		"2f2feab2026849078ba485f918791bdc",
		cloudflare.AccountRulesetUpdateParams{
			ID:          cloudflare.F[any](map[string]interface{}{}),
			Description: cloudflare.F("My ruleset to execute managed rulesets"),
			Kind:        cloudflare.F(cloudflare.AccountRulesetUpdateParamsKindRoot),
			Name:        cloudflare.F("My ruleset"),
			Phase:       cloudflare.F(cloudflare.AccountRulesetUpdateParamsPhaseHTTPRequestFirewallCustom),
			Rules: cloudflare.F([]cloudflare.AccountRulesetUpdateParamsRule{cloudflare.AccountRulesetUpdateParamsRulesOexZd8xKBlockRule(cloudflare.AccountRulesetUpdateParamsRulesOexZd8xKBlockRule{
				Action: cloudflare.F(cloudflare.AccountRulesetUpdateParamsRulesOexZd8xKBlockRuleActionBlock),
				ActionParameters: cloudflare.F(cloudflare.AccountRulesetUpdateParamsRulesOexZd8xKBlockRuleActionParameters{
					Response: cloudflare.F(cloudflare.AccountRulesetUpdateParamsRulesOexZd8xKBlockRuleActionParametersResponse{
						Content:     cloudflare.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
						ContentType: cloudflare.F("application/json"),
						StatusCode:  cloudflare.F(int64(400)),
					}),
				}),
				Description: cloudflare.F[any]("Block when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F[any](map[string]interface{}{}),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				ID:          cloudflare.F("3a03d665bac047339bb530ecb439a90d"),
				Logging: cloudflare.F(cloudflare.AccountRulesetUpdateParamsRulesOexZd8xKBlockRuleLogging{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			}), cloudflare.AccountRulesetUpdateParamsRulesOexZd8xKBlockRule(cloudflare.AccountRulesetUpdateParamsRulesOexZd8xKBlockRule{
				Action: cloudflare.F(cloudflare.AccountRulesetUpdateParamsRulesOexZd8xKBlockRuleActionBlock),
				ActionParameters: cloudflare.F(cloudflare.AccountRulesetUpdateParamsRulesOexZd8xKBlockRuleActionParameters{
					Response: cloudflare.F(cloudflare.AccountRulesetUpdateParamsRulesOexZd8xKBlockRuleActionParametersResponse{
						Content:     cloudflare.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
						ContentType: cloudflare.F("application/json"),
						StatusCode:  cloudflare.F(int64(400)),
					}),
				}),
				Description: cloudflare.F[any]("Block when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F[any](map[string]interface{}{}),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				ID:          cloudflare.F("3a03d665bac047339bb530ecb439a90d"),
				Logging: cloudflare.F(cloudflare.AccountRulesetUpdateParamsRulesOexZd8xKBlockRuleLogging{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			}), cloudflare.AccountRulesetUpdateParamsRulesOexZd8xKBlockRule(cloudflare.AccountRulesetUpdateParamsRulesOexZd8xKBlockRule{
				Action: cloudflare.F(cloudflare.AccountRulesetUpdateParamsRulesOexZd8xKBlockRuleActionBlock),
				ActionParameters: cloudflare.F(cloudflare.AccountRulesetUpdateParamsRulesOexZd8xKBlockRuleActionParameters{
					Response: cloudflare.F(cloudflare.AccountRulesetUpdateParamsRulesOexZd8xKBlockRuleActionParametersResponse{
						Content:     cloudflare.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
						ContentType: cloudflare.F("application/json"),
						StatusCode:  cloudflare.F(int64(400)),
					}),
				}),
				Description: cloudflare.F[any]("Block when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F[any](map[string]interface{}{}),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				ID:          cloudflare.F("3a03d665bac047339bb530ecb439a90d"),
				Logging: cloudflare.F(cloudflare.AccountRulesetUpdateParamsRulesOexZd8xKBlockRuleLogging{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			})}),
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	err := client.Accounts.Rulesets.Delete(
		context.TODO(),
		"abf9b32d38c5f572afde3336ec0ce302",
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Rulesets.AccountRulesetsNewAnAccountRuleset(
		context.TODO(),
		"abf9b32d38c5f572afde3336ec0ce302",
		cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParams{
			ID:          cloudflare.F[any](map[string]interface{}{}),
			Description: cloudflare.F("My ruleset to execute managed rulesets"),
			Kind:        cloudflare.F(cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsKindRoot),
			Name:        cloudflare.F("My ruleset"),
			Phase:       cloudflare.F(cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhaseHTTPRequestFirewallCustom),
			Rules: cloudflare.F([]cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRule{cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRule(cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRule{
				Action: cloudflare.F(cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRuleActionBlock),
				ActionParameters: cloudflare.F(cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRuleActionParameters{
					Response: cloudflare.F(cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRuleActionParametersResponse{
						Content:     cloudflare.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
						ContentType: cloudflare.F("application/json"),
						StatusCode:  cloudflare.F(int64(400)),
					}),
				}),
				Description: cloudflare.F[any]("Block when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F[any](map[string]interface{}{}),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				ID:          cloudflare.F("3a03d665bac047339bb530ecb439a90d"),
				Logging: cloudflare.F(cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRuleLogging{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			}), cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRule(cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRule{
				Action: cloudflare.F(cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRuleActionBlock),
				ActionParameters: cloudflare.F(cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRuleActionParameters{
					Response: cloudflare.F(cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRuleActionParametersResponse{
						Content:     cloudflare.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
						ContentType: cloudflare.F("application/json"),
						StatusCode:  cloudflare.F(int64(400)),
					}),
				}),
				Description: cloudflare.F[any]("Block when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F[any](map[string]interface{}{}),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				ID:          cloudflare.F("3a03d665bac047339bb530ecb439a90d"),
				Logging: cloudflare.F(cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRuleLogging{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			}), cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRule(cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRule{
				Action: cloudflare.F(cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRuleActionBlock),
				ActionParameters: cloudflare.F(cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRuleActionParameters{
					Response: cloudflare.F(cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRuleActionParametersResponse{
						Content:     cloudflare.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
						ContentType: cloudflare.F("application/json"),
						StatusCode:  cloudflare.F(int64(400)),
					}),
				}),
				Description: cloudflare.F[any]("Block when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F[any](map[string]interface{}{}),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				ID:          cloudflare.F("3a03d665bac047339bb530ecb439a90d"),
				Logging: cloudflare.F(cloudflare.AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRuleLogging{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			})}),
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Rulesets.AccountRulesetsListAccountRulesets(context.TODO(), "abf9b32d38c5f572afde3336ec0ce302")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
