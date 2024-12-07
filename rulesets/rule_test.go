// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rulesets_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/rulesets"
)

func TestRuleNewWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.Rulesets.Rules.New(
		context.TODO(),
		"2f2feab2026849078ba485f918791bdc",
		rulesets.RuleNewParams{
			Body: rulesets.BlockRuleParam{
				ID:     cloudflare.F("3a03d665bac047339bb530ecb439a90d"),
				Action: cloudflare.F(rulesets.BlockRuleActionBlock),
				ActionParameters: cloudflare.F(rulesets.BlockRuleActionParametersParam{
					Response: cloudflare.F(rulesets.BlockRuleActionParametersResponseParam{
						Content:     cloudflare.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
						ContentType: cloudflare.F("application/json"),
						StatusCode:  cloudflare.F(int64(400)),
					}),
				}),
				Description: cloudflare.F("Block when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F(true),
				ExposedCredentialCheck: cloudflare.F(rulesets.BlockRuleExposedCredentialCheckParam{
					PasswordExpression: cloudflare.F("url_decode(http.request.body.form[\\\"password\\\"][0])"),
					UsernameExpression: cloudflare.F("url_decode(http.request.body.form[\\\"username\\\"][0])"),
				}),
				Expression: cloudflare.F("ip.src ne 1.1.1.1"),
				Logging: cloudflare.F(rulesets.LoggingParam{
					Enabled: cloudflare.F(true),
				}),
				Ratelimit: cloudflare.F(rulesets.BlockRuleRatelimitParam{
					Characteristics:         cloudflare.F([]string{"ip.src"}),
					Period:                  cloudflare.F(rulesets.BlockRuleRatelimitPeriod10),
					CountingExpression:      cloudflare.F("http.request.body.raw eq \"abcd\""),
					MitigationTimeout:       cloudflare.F(int64(600)),
					RequestsPerPeriod:       cloudflare.F(int64(1000)),
					RequestsToOrigin:        cloudflare.F(true),
					ScorePerPeriod:          cloudflare.F(int64(400)),
					ScoreResponseHeaderName: cloudflare.F("my-score"),
				}),
				Ref: cloudflare.F("my_ref"),
			},
			AccountID: cloudflare.F("account_id"),
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

func TestRuleDeleteWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.Rulesets.Rules.Delete(
		context.TODO(),
		"2f2feab2026849078ba485f918791bdc",
		"3a03d665bac047339bb530ecb439a90d",
		rulesets.RuleDeleteParams{
			AccountID: cloudflare.F("account_id"),
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

func TestRuleEditWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.Rulesets.Rules.Edit(
		context.TODO(),
		"2f2feab2026849078ba485f918791bdc",
		"3a03d665bac047339bb530ecb439a90d",
		rulesets.RuleEditParams{
			Body: rulesets.BlockRuleParam{
				ID:     cloudflare.F("3a03d665bac047339bb530ecb439a90d"),
				Action: cloudflare.F(rulesets.BlockRuleActionBlock),
				ActionParameters: cloudflare.F(rulesets.BlockRuleActionParametersParam{
					Response: cloudflare.F(rulesets.BlockRuleActionParametersResponseParam{
						Content:     cloudflare.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
						ContentType: cloudflare.F("application/json"),
						StatusCode:  cloudflare.F(int64(400)),
					}),
				}),
				Description: cloudflare.F("Block when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F(true),
				ExposedCredentialCheck: cloudflare.F(rulesets.BlockRuleExposedCredentialCheckParam{
					PasswordExpression: cloudflare.F("url_decode(http.request.body.form[\\\"password\\\"][0])"),
					UsernameExpression: cloudflare.F("url_decode(http.request.body.form[\\\"username\\\"][0])"),
				}),
				Expression: cloudflare.F("ip.src ne 1.1.1.1"),
				Logging: cloudflare.F(rulesets.LoggingParam{
					Enabled: cloudflare.F(true),
				}),
				Ratelimit: cloudflare.F(rulesets.BlockRuleRatelimitParam{
					Characteristics:         cloudflare.F([]string{"ip.src"}),
					Period:                  cloudflare.F(rulesets.BlockRuleRatelimitPeriod10),
					CountingExpression:      cloudflare.F("http.request.body.raw eq \"abcd\""),
					MitigationTimeout:       cloudflare.F(int64(600)),
					RequestsPerPeriod:       cloudflare.F(int64(1000)),
					RequestsToOrigin:        cloudflare.F(true),
					ScorePerPeriod:          cloudflare.F(int64(400)),
					ScoreResponseHeaderName: cloudflare.F("my-score"),
				}),
				Ref: cloudflare.F("my_ref"),
			},
			AccountID: cloudflare.F("account_id"),
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
