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

func TestPhaseUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Rulesets.Phases.Update(
		context.TODO(),
		rulesets.PhaseUpdateParamsRulesetPhaseHTTPRequestFirewallCustom,
		rulesets.PhaseUpdateParams{
			ID: cloudflare.F("2f2feab2026849078ba485f918791bdc"),
			Rules: cloudflare.F([]rulesets.PhaseUpdateParamsRule{rulesets.PhaseUpdateParamsRulesRulesetsBlockRule(rulesets.PhaseUpdateParamsRulesRulesetsBlockRule{
				Action: cloudflare.F(rulesets.PhaseUpdateParamsRulesRulesetsBlockRuleActionBlock),
				ActionParameters: cloudflare.F(rulesets.PhaseUpdateParamsRulesRulesetsBlockRuleActionParameters{
					Response: cloudflare.F(rulesets.PhaseUpdateParamsRulesRulesetsBlockRuleActionParametersResponse{
						Content:     cloudflare.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
						ContentType: cloudflare.F("application/json"),
						StatusCode:  cloudflare.F(int64(400)),
					}),
				}),
				Description: cloudflare.F("Block when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				ID:          cloudflare.F("3a03d665bac047339bb530ecb439a90d"),
				Logging: cloudflare.F(rulesets.PhaseUpdateParamsRulesRulesetsBlockRuleLogging{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			}), rulesets.PhaseUpdateParamsRulesRulesetsBlockRule(rulesets.PhaseUpdateParamsRulesRulesetsBlockRule{
				Action: cloudflare.F(rulesets.PhaseUpdateParamsRulesRulesetsBlockRuleActionBlock),
				ActionParameters: cloudflare.F(rulesets.PhaseUpdateParamsRulesRulesetsBlockRuleActionParameters{
					Response: cloudflare.F(rulesets.PhaseUpdateParamsRulesRulesetsBlockRuleActionParametersResponse{
						Content:     cloudflare.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
						ContentType: cloudflare.F("application/json"),
						StatusCode:  cloudflare.F(int64(400)),
					}),
				}),
				Description: cloudflare.F("Block when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				ID:          cloudflare.F("3a03d665bac047339bb530ecb439a90d"),
				Logging: cloudflare.F(rulesets.PhaseUpdateParamsRulesRulesetsBlockRuleLogging{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			}), rulesets.PhaseUpdateParamsRulesRulesetsBlockRule(rulesets.PhaseUpdateParamsRulesRulesetsBlockRule{
				Action: cloudflare.F(rulesets.PhaseUpdateParamsRulesRulesetsBlockRuleActionBlock),
				ActionParameters: cloudflare.F(rulesets.PhaseUpdateParamsRulesRulesetsBlockRuleActionParameters{
					Response: cloudflare.F(rulesets.PhaseUpdateParamsRulesRulesetsBlockRuleActionParametersResponse{
						Content:     cloudflare.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
						ContentType: cloudflare.F("application/json"),
						StatusCode:  cloudflare.F(int64(400)),
					}),
				}),
				Description: cloudflare.F("Block when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				ID:          cloudflare.F("3a03d665bac047339bb530ecb439a90d"),
				Logging: cloudflare.F(rulesets.PhaseUpdateParamsRulesRulesetsBlockRuleLogging{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			})}),
			AccountID:   cloudflare.F("string"),
			ZoneID:      cloudflare.F("string"),
			Description: cloudflare.F("My ruleset to execute managed rulesets"),
			Kind:        cloudflare.F(rulesets.PhaseUpdateParamsKindRoot),
			Name:        cloudflare.F("My ruleset"),
			Phase:       cloudflare.F(rulesets.PhaseUpdateParamsPhaseHTTPRequestFirewallCustom),
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

func TestPhaseGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Rulesets.Phases.Get(
		context.TODO(),
		rulesets.PhaseGetParamsRulesetPhaseHTTPRequestFirewallCustom,
		rulesets.PhaseGetParams{
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
