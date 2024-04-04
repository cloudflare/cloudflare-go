// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rulesets_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/rulesets"
)

func TestRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Rulesets.Rules.New(
		context.TODO(),
		"2f2feab2026849078ba485f918791bdc",
		rulesets.RuleNewParamsBlockRule{
			AccountID: cloudflare.F("string"),
			ZoneID:    cloudflare.F("string"),
			ID:        cloudflare.F("3a03d665bac047339bb530ecb439a90d"),
			Action:    cloudflare.F(rulesets.RuleNewParamsBlockRuleActionBlock),
			ActionParameters: cloudflare.F(rulesets.RuleNewParamsBlockRuleActionParameters{
				Response: cloudflare.F(rulesets.RuleNewParamsBlockRuleActionParametersResponse{
					Content:     cloudflare.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
					ContentType: cloudflare.F("application/json"),
					StatusCode:  cloudflare.F(int64(400)),
				}),
			}),
			Description: cloudflare.F("Block when the IP address is not 1.1.1.1"),
			Enabled:     cloudflare.F(true),
			Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
			Logging: cloudflare.F(shared.UnnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751cParam{
				Enabled: cloudflare.F(true),
			}),
			Ref: cloudflare.F("my_ref"),
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
	_, err := client.Rulesets.Rules.Delete(
		context.TODO(),
		"2f2feab2026849078ba485f918791bdc",
		"3a03d665bac047339bb530ecb439a90d",
		rulesets.RuleDeleteParams{
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

func TestRuleEditWithOptionalParams(t *testing.T) {
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
	_, err := client.Rulesets.Rules.Edit(
		context.TODO(),
		"2f2feab2026849078ba485f918791bdc",
		"3a03d665bac047339bb530ecb439a90d",
		rulesets.RuleEditParamsBlockRule{
			AccountID: cloudflare.F("string"),
			ZoneID:    cloudflare.F("string"),
			ID:        cloudflare.F("3a03d665bac047339bb530ecb439a90d"),
			Action:    cloudflare.F(rulesets.RuleEditParamsBlockRuleActionBlock),
			ActionParameters: cloudflare.F(rulesets.RuleEditParamsBlockRuleActionParameters{
				Response: cloudflare.F(rulesets.RuleEditParamsBlockRuleActionParametersResponse{
					Content:     cloudflare.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
					ContentType: cloudflare.F("application/json"),
					StatusCode:  cloudflare.F(int64(400)),
				}),
			}),
			Description: cloudflare.F("Block when the IP address is not 1.1.1.1"),
			Enabled:     cloudflare.F(true),
			Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
			Logging: cloudflare.F(shared.UnnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751cParam{
				Enabled: cloudflare.F(true),
			}),
			Ref: cloudflare.F("my_ref"),
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
