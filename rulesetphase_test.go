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

func TestRulesetPhaseUpdateWithOptionalParams(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Rulesets.Phases.Update(
		context.TODO(),
		cloudflare.RulesetPhaseUpdateParamsRulesetPhaseHTTPRequestFirewallCustom,
		cloudflare.RulesetPhaseUpdateParams{
			AccountID: cloudflare.F("string"),
			ZoneID:    cloudflare.F("abf9b32d38c5f572afde3336ec0ce302"),
			ID:        cloudflare.F("2f2feab2026849078ba485f918791bdc"),
			Rules: cloudflare.F([]cloudflare.RulesetPhaseUpdateParamsRule{cloudflare.RulesetPhaseUpdateParamsRulesRulesetsBlockRule(cloudflare.RulesetPhaseUpdateParamsRulesRulesetsBlockRule{
				Action: cloudflare.F(cloudflare.RulesetPhaseUpdateParamsRulesRulesetsBlockRuleActionBlock),
				ActionParameters: cloudflare.F(cloudflare.RulesetPhaseUpdateParamsRulesRulesetsBlockRuleActionParameters{
					Response: cloudflare.F(cloudflare.RulesetPhaseUpdateParamsRulesRulesetsBlockRuleActionParametersResponse{
						Content:     cloudflare.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
						ContentType: cloudflare.F("application/json"),
						StatusCode:  cloudflare.F(int64(400)),
					}),
				}),
				Description: cloudflare.F("Block when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				ID:          cloudflare.F("3a03d665bac047339bb530ecb439a90d"),
				Logging: cloudflare.F(cloudflare.RulesetPhaseUpdateParamsRulesRulesetsBlockRuleLogging{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			}), cloudflare.RulesetPhaseUpdateParamsRulesRulesetsBlockRule(cloudflare.RulesetPhaseUpdateParamsRulesRulesetsBlockRule{
				Action: cloudflare.F(cloudflare.RulesetPhaseUpdateParamsRulesRulesetsBlockRuleActionBlock),
				ActionParameters: cloudflare.F(cloudflare.RulesetPhaseUpdateParamsRulesRulesetsBlockRuleActionParameters{
					Response: cloudflare.F(cloudflare.RulesetPhaseUpdateParamsRulesRulesetsBlockRuleActionParametersResponse{
						Content:     cloudflare.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
						ContentType: cloudflare.F("application/json"),
						StatusCode:  cloudflare.F(int64(400)),
					}),
				}),
				Description: cloudflare.F("Block when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				ID:          cloudflare.F("3a03d665bac047339bb530ecb439a90d"),
				Logging: cloudflare.F(cloudflare.RulesetPhaseUpdateParamsRulesRulesetsBlockRuleLogging{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			}), cloudflare.RulesetPhaseUpdateParamsRulesRulesetsBlockRule(cloudflare.RulesetPhaseUpdateParamsRulesRulesetsBlockRule{
				Action: cloudflare.F(cloudflare.RulesetPhaseUpdateParamsRulesRulesetsBlockRuleActionBlock),
				ActionParameters: cloudflare.F(cloudflare.RulesetPhaseUpdateParamsRulesRulesetsBlockRuleActionParameters{
					Response: cloudflare.F(cloudflare.RulesetPhaseUpdateParamsRulesRulesetsBlockRuleActionParametersResponse{
						Content:     cloudflare.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
						ContentType: cloudflare.F("application/json"),
						StatusCode:  cloudflare.F(int64(400)),
					}),
				}),
				Description: cloudflare.F("Block when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				ID:          cloudflare.F("3a03d665bac047339bb530ecb439a90d"),
				Logging: cloudflare.F(cloudflare.RulesetPhaseUpdateParamsRulesRulesetsBlockRuleLogging{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			})}),
			Description: cloudflare.F("My ruleset to execute managed rulesets"),
			Kind:        cloudflare.F(cloudflare.RulesetPhaseUpdateParamsKindRoot),
			Name:        cloudflare.F("My ruleset"),
			Phase:       cloudflare.F(cloudflare.RulesetPhaseUpdateParamsPhaseHTTPRequestFirewallCustom),
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

func TestRulesetPhaseGet(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Rulesets.Phases.Get(
		context.TODO(),
		cloudflare.RulesetPhaseGetParamsRulesetPhaseHTTPRequestFirewallCustom,
		cloudflare.RulesetPhaseGetParams{
			AccountID: cloudflare.F("string"),
			ZoneID:    cloudflare.F("abf9b32d38c5f572afde3336ec0ce302"),
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
