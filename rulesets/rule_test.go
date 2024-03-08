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
		rulesets.RuleNewParams{
			AccountID: cloudflare.F("string"),
			ZoneID:    cloudflare.F("string"),
			Position: cloudflare.F[rulesets.RuleNewParamsPosition](rulesets.RuleNewParamsPositionBeforePosition(rulesets.RuleNewParamsPositionBeforePosition{
				Before: cloudflare.F("da5e8e506c8e7877fe06cdf4c41add54"),
			})),
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
		rulesets.RuleEditParams{
			AccountID: cloudflare.F("string"),
			ZoneID:    cloudflare.F("string"),
			Position: cloudflare.F[rulesets.RuleEditParamsPosition](rulesets.RuleEditParamsPositionBeforePosition(rulesets.RuleEditParamsPositionBeforePosition{
				Before: cloudflare.F("da5e8e506c8e7877fe06cdf4c41add54"),
			})),
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
