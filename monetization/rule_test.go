// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package monetization_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/monetization"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

func TestRuleUpdate(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Monetization.Rules.Update(context.TODO(), monetization.RuleUpdateParams{
		ZoneID: cloudflare.F("9f1839b6152d298aca64c4e906b6d074"),
		MonetizationRulesetInput: monetization.MonetizationRulesetInputParam{
			Rules: cloudflare.F([]monetization.MonetizationRuleInputUnionParam{monetization.MonetizationRuleInputFixedPriceParam{
				Address:     cloudflare.F("0x1234567890abcdef1234567890abcdef12345678==="),
				Expression:  cloudflare.F(`(http.request.uri.path eq "/premium" and http.request.method in {"GET" "POST"})`),
				Price:       cloudflare.F("250000"),
				Scheme:      cloudflare.F(monetization.MonetizationRuleInputFixedPriceSchemeExact),
				ID:          cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
				Description: cloudflare.F("Premium API endpoint"),
				Enabled:     cloudflare.F(true),
			}}),
		},
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRuleDelete(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Monetization.Rules.Delete(context.TODO(), monetization.RuleDeleteParams{
		ZoneID: cloudflare.F("9f1839b6152d298aca64c4e906b6d074"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRuleDeleteRule(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Monetization.Rules.DeleteRule(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		monetization.RuleDeleteRuleParams{
			ZoneID: cloudflare.F("9f1839b6152d298aca64c4e906b6d074"),
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

func TestRuleEditRuleWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Monetization.Rules.EditRule(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		monetization.RuleEditRuleParams{
			ZoneID: cloudflare.F("9f1839b6152d298aca64c4e906b6d074"),
			MonetizationRulePatch: monetization.MonetizationRulePatchParam{
				Address:     cloudflare.F("0x1234567890abcdef1234567890abcdef12345678==="),
				Description: cloudflare.F("Premium API endpoint"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F(`(http.request.uri.path eq "/premium" and http.request.method in {"GET" "POST"})`),
				Price:       cloudflare.F("250000"),
				Scheme:      cloudflare.F(monetization.MonetizationRulePatchSchemeExact),
			},
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

func TestRuleGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Monetization.Rules.Get(context.TODO(), monetization.RuleGetParams{
		ZoneID: cloudflare.F("9f1839b6152d298aca64c4e906b6d074"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRuleGetRule(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Monetization.Rules.GetRule(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		monetization.RuleGetRuleParams{
			ZoneID: cloudflare.F("9f1839b6152d298aca64c4e906b6d074"),
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
