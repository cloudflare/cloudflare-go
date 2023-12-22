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

func TestZoneRulesetRuleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Rulesets.Rules.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"2f2feab2026849078ba485f918791bdc",
		"3a03d665bac047339bb530ecb439a90d",
		cloudflare.ZoneRulesetRuleUpdateParams{
			Action:     cloudflare.F("execute"),
			Expression: cloudflare.F("ip.src ne 1.1.1.1"),
			ActionParameters: cloudflare.F[any](map[string]interface{}{
				"id": "4814384a9e5d4991b9815dcfc25d2f1f",
			}),
			Description: cloudflare.F("Execute the OWASP ruleset when the IP address is not 1.1.1.1"),
			Enabled:     cloudflare.F(true),
			Logging: cloudflare.F(cloudflare.ZoneRulesetRuleUpdateParamsLogging{
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

func TestZoneRulesetRuleDelete(t *testing.T) {
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
	_, err := client.Zones.Rulesets.Rules.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"2f2feab2026849078ba485f918791bdc",
		"3a03d665bac047339bb530ecb439a90d",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Rulesets.Rules.ZoneRulesetsNewAZoneRulesetRule(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"2f2feab2026849078ba485f918791bdc",
		cloudflare.ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleParams{
			Action:     cloudflare.F("execute"),
			Expression: cloudflare.F("ip.src ne 1.1.1.1"),
			ActionParameters: cloudflare.F[any](map[string]interface{}{
				"id": "4814384a9e5d4991b9815dcfc25d2f1f",
			}),
			Description: cloudflare.F("Execute the OWASP ruleset when the IP address is not 1.1.1.1"),
			Enabled:     cloudflare.F(true),
			Logging: cloudflare.F(cloudflare.ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleParamsLogging{
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
