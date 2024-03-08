// File generated from our OpenAPI spec by Stainless.

package rulesets_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/cloudflare/cloudflare-go/rulesets"
)

func TestRulesetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Rulesets.New(context.TODO(), rulesets.RulesetNewParams{
		Kind:  cloudflare.F(rulesets.RulesetNewParamsKindRoot),
		Name:  cloudflare.F("My ruleset"),
		Phase: cloudflare.F(rulesets.RulesetNewParamsPhaseHTTPRequestFirewallCustom),
		Rules: cloudflare.F([]rulesets.RulesetNewParamsRule{rulesets.RulesetNewParamsRulesRulesetsBlockRule(rulesets.RulesetNewParamsRulesRulesetsBlockRule{
			Action: cloudflare.F(rulesets.RulesetNewParamsRulesRulesetsBlockRuleActionBlock),
			ActionParameters: cloudflare.F(rulesets.RulesetNewParamsRulesRulesetsBlockRuleActionParameters{
				Response: cloudflare.F(rulesets.RulesetNewParamsRulesRulesetsBlockRuleActionParametersResponse{
					Content:     cloudflare.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
					ContentType: cloudflare.F("application/json"),
					StatusCode:  cloudflare.F(int64(400)),
				}),
			}),
			Description: cloudflare.F("Block when the IP address is not 1.1.1.1"),
			Enabled:     cloudflare.F(true),
			Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
			ID:          cloudflare.F("3a03d665bac047339bb530ecb439a90d"),
			Logging: cloudflare.F(rulesets.RulesetNewParamsRulesRulesetsBlockRuleLogging{
				Enabled: cloudflare.F(true),
			}),
			Ref: cloudflare.F("my_ref"),
		}), rulesets.RulesetNewParamsRulesRulesetsBlockRule(rulesets.RulesetNewParamsRulesRulesetsBlockRule{
			Action: cloudflare.F(rulesets.RulesetNewParamsRulesRulesetsBlockRuleActionBlock),
			ActionParameters: cloudflare.F(rulesets.RulesetNewParamsRulesRulesetsBlockRuleActionParameters{
				Response: cloudflare.F(rulesets.RulesetNewParamsRulesRulesetsBlockRuleActionParametersResponse{
					Content:     cloudflare.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
					ContentType: cloudflare.F("application/json"),
					StatusCode:  cloudflare.F(int64(400)),
				}),
			}),
			Description: cloudflare.F("Block when the IP address is not 1.1.1.1"),
			Enabled:     cloudflare.F(true),
			Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
			ID:          cloudflare.F("3a03d665bac047339bb530ecb439a90d"),
			Logging: cloudflare.F(rulesets.RulesetNewParamsRulesRulesetsBlockRuleLogging{
				Enabled: cloudflare.F(true),
			}),
			Ref: cloudflare.F("my_ref"),
		}), rulesets.RulesetNewParamsRulesRulesetsBlockRule(rulesets.RulesetNewParamsRulesRulesetsBlockRule{
			Action: cloudflare.F(rulesets.RulesetNewParamsRulesRulesetsBlockRuleActionBlock),
			ActionParameters: cloudflare.F(rulesets.RulesetNewParamsRulesRulesetsBlockRuleActionParameters{
				Response: cloudflare.F(rulesets.RulesetNewParamsRulesRulesetsBlockRuleActionParametersResponse{
					Content:     cloudflare.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
					ContentType: cloudflare.F("application/json"),
					StatusCode:  cloudflare.F(int64(400)),
				}),
			}),
			Description: cloudflare.F("Block when the IP address is not 1.1.1.1"),
			Enabled:     cloudflare.F(true),
			Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
			ID:          cloudflare.F("3a03d665bac047339bb530ecb439a90d"),
			Logging: cloudflare.F(rulesets.RulesetNewParamsRulesRulesetsBlockRuleLogging{
				Enabled: cloudflare.F(true),
			}),
			Ref: cloudflare.F("my_ref"),
		})}),
		AccountID:   cloudflare.F("string"),
		ZoneID:      cloudflare.F("string"),
		Description: cloudflare.F("My ruleset to execute managed rulesets"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRulesetUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Rulesets.Update(
		context.TODO(),
		"2f2feab2026849078ba485f918791bdc",
		rulesets.RulesetUpdateParams{
			ID: cloudflare.F("2f2feab2026849078ba485f918791bdc"),
			Rules: cloudflare.F([]rulesets.RulesetUpdateParamsRule{rulesets.RulesetUpdateParamsRulesRulesetsBlockRule(rulesets.RulesetUpdateParamsRulesRulesetsBlockRule{
				Action: cloudflare.F(rulesets.RulesetUpdateParamsRulesRulesetsBlockRuleActionBlock),
				ActionParameters: cloudflare.F(rulesets.RulesetUpdateParamsRulesRulesetsBlockRuleActionParameters{
					Response: cloudflare.F(rulesets.RulesetUpdateParamsRulesRulesetsBlockRuleActionParametersResponse{
						Content:     cloudflare.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
						ContentType: cloudflare.F("application/json"),
						StatusCode:  cloudflare.F(int64(400)),
					}),
				}),
				Description: cloudflare.F("Block when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				ID:          cloudflare.F("3a03d665bac047339bb530ecb439a90d"),
				Logging: cloudflare.F(rulesets.RulesetUpdateParamsRulesRulesetsBlockRuleLogging{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			}), rulesets.RulesetUpdateParamsRulesRulesetsBlockRule(rulesets.RulesetUpdateParamsRulesRulesetsBlockRule{
				Action: cloudflare.F(rulesets.RulesetUpdateParamsRulesRulesetsBlockRuleActionBlock),
				ActionParameters: cloudflare.F(rulesets.RulesetUpdateParamsRulesRulesetsBlockRuleActionParameters{
					Response: cloudflare.F(rulesets.RulesetUpdateParamsRulesRulesetsBlockRuleActionParametersResponse{
						Content:     cloudflare.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
						ContentType: cloudflare.F("application/json"),
						StatusCode:  cloudflare.F(int64(400)),
					}),
				}),
				Description: cloudflare.F("Block when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				ID:          cloudflare.F("3a03d665bac047339bb530ecb439a90d"),
				Logging: cloudflare.F(rulesets.RulesetUpdateParamsRulesRulesetsBlockRuleLogging{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			}), rulesets.RulesetUpdateParamsRulesRulesetsBlockRule(rulesets.RulesetUpdateParamsRulesRulesetsBlockRule{
				Action: cloudflare.F(rulesets.RulesetUpdateParamsRulesRulesetsBlockRuleActionBlock),
				ActionParameters: cloudflare.F(rulesets.RulesetUpdateParamsRulesRulesetsBlockRuleActionParameters{
					Response: cloudflare.F(rulesets.RulesetUpdateParamsRulesRulesetsBlockRuleActionParametersResponse{
						Content:     cloudflare.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
						ContentType: cloudflare.F("application/json"),
						StatusCode:  cloudflare.F(int64(400)),
					}),
				}),
				Description: cloudflare.F("Block when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				ID:          cloudflare.F("3a03d665bac047339bb530ecb439a90d"),
				Logging: cloudflare.F(rulesets.RulesetUpdateParamsRulesRulesetsBlockRuleLogging{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			})}),
			AccountID:   cloudflare.F("string"),
			ZoneID:      cloudflare.F("string"),
			Description: cloudflare.F("My ruleset to execute managed rulesets"),
			Kind:        cloudflare.F(rulesets.RulesetUpdateParamsKindRoot),
			Name:        cloudflare.F("My ruleset"),
			Phase:       cloudflare.F(rulesets.RulesetUpdateParamsPhaseHTTPRequestFirewallCustom),
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

func TestRulesetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Rulesets.List(context.TODO(), rulesets.RulesetListParams{
		AccountID: cloudflare.F("string"),
		ZoneID:    cloudflare.F("string"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRulesetDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Rulesets.Delete(
		context.TODO(),
		"2f2feab2026849078ba485f918791bdc",
		rulesets.RulesetDeleteParams{
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

func TestRulesetGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Rulesets.Get(
		context.TODO(),
		"2f2feab2026849078ba485f918791bdc",
		rulesets.RulesetGetParams{
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
