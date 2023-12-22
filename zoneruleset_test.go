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

func TestZoneRulesetGet(t *testing.T) {
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
	_, err := client.Zones.Rulesets.Get(
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

func TestZoneRulesetUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Rulesets.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"2f2feab2026849078ba485f918791bdc",
		cloudflare.ZoneRulesetUpdateParams{
			Rules: cloudflare.F([]cloudflare.ZoneRulesetUpdateParamsRule{cloudflare.ZoneRulesetUpdateParamsRulesCreateUpdateRule(cloudflare.ZoneRulesetUpdateParamsRulesCreateUpdateRule{
				Action: cloudflare.F("execute"),
				ActionParameters: cloudflare.F[any](map[string]interface{}{
					"id": "4814384a9e5d4991b9815dcfc25d2f1f",
				}),
				Description: cloudflare.F("Execute the OWASP ruleset when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				Logging: cloudflare.F(cloudflare.ZoneRulesetUpdateParamsRulesCreateUpdateRuleLogging{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			}), cloudflare.ZoneRulesetUpdateParamsRulesCreateUpdateRule(cloudflare.ZoneRulesetUpdateParamsRulesCreateUpdateRule{
				Action: cloudflare.F("execute"),
				ActionParameters: cloudflare.F[any](map[string]interface{}{
					"id": "4814384a9e5d4991b9815dcfc25d2f1f",
				}),
				Description: cloudflare.F("Execute the OWASP ruleset when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				Logging: cloudflare.F(cloudflare.ZoneRulesetUpdateParamsRulesCreateUpdateRuleLogging{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			}), cloudflare.ZoneRulesetUpdateParamsRulesCreateUpdateRule(cloudflare.ZoneRulesetUpdateParamsRulesCreateUpdateRule{
				Action: cloudflare.F("execute"),
				ActionParameters: cloudflare.F[any](map[string]interface{}{
					"id": "4814384a9e5d4991b9815dcfc25d2f1f",
				}),
				Description: cloudflare.F("Execute the OWASP ruleset when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				Logging: cloudflare.F(cloudflare.ZoneRulesetUpdateParamsRulesCreateUpdateRuleLogging{
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

func TestZoneRulesetDelete(t *testing.T) {
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
	err := client.Zones.Rulesets.Delete(
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

func TestZoneRulesetZoneRulesetsNewAZoneRulesetWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Rulesets.ZoneRulesetsNewAZoneRuleset(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZoneRulesetZoneRulesetsNewAZoneRulesetParams{
			Kind:  cloudflare.F(cloudflare.ZoneRulesetZoneRulesetsNewAZoneRulesetParamsKindRoot),
			Name:  cloudflare.F("My ruleset"),
			Phase: cloudflare.F("http_request_firewall_managed"),
			Rules: cloudflare.F([]cloudflare.ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRule{cloudflare.ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesCreateUpdateRule(cloudflare.ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesCreateUpdateRule{
				Action: cloudflare.F("execute"),
				ActionParameters: cloudflare.F[any](map[string]interface{}{
					"id": "4814384a9e5d4991b9815dcfc25d2f1f",
				}),
				Description: cloudflare.F("Execute the OWASP ruleset when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				Logging: cloudflare.F(cloudflare.ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesCreateUpdateRuleLogging{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			}), cloudflare.ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesCreateUpdateRule(cloudflare.ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesCreateUpdateRule{
				Action: cloudflare.F("execute"),
				ActionParameters: cloudflare.F[any](map[string]interface{}{
					"id": "4814384a9e5d4991b9815dcfc25d2f1f",
				}),
				Description: cloudflare.F("Execute the OWASP ruleset when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				Logging: cloudflare.F(cloudflare.ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesCreateUpdateRuleLogging{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			}), cloudflare.ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesCreateUpdateRule(cloudflare.ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesCreateUpdateRule{
				Action: cloudflare.F("execute"),
				ActionParameters: cloudflare.F[any](map[string]interface{}{
					"id": "4814384a9e5d4991b9815dcfc25d2f1f",
				}),
				Description: cloudflare.F("Execute the OWASP ruleset when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				Logging: cloudflare.F(cloudflare.ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesCreateUpdateRuleLogging{
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

func TestZoneRulesetZoneRulesetsListZoneRulesets(t *testing.T) {
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
	_, err := client.Zones.Rulesets.ZoneRulesetsListZoneRulesets(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
