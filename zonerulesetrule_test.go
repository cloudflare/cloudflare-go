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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Zones.Rulesets.Rules.Update(
		context.TODO(),
		"9f1839b6152d298aca64c4e906b6d074",
		"2f2feab2026849078ba485f918791bdc",
		"3a03d665bac047339bb530ecb439a90d",
		cloudflare.ZoneRulesetRuleUpdateParams{
			Position: cloudflare.F[cloudflare.ZoneRulesetRuleUpdateParamsPosition](cloudflare.ZoneRulesetRuleUpdateParamsPositionObject(cloudflare.ZoneRulesetRuleUpdateParamsPositionObject{
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Zones.Rulesets.Rules.Delete(
		context.TODO(),
		"9f1839b6152d298aca64c4e906b6d074",
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Zones.Rulesets.Rules.ZoneRulesetsNewAZoneRulesetRule(
		context.TODO(),
		"9f1839b6152d298aca64c4e906b6d074",
		"2f2feab2026849078ba485f918791bdc",
		cloudflare.ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleParams{
			Position: cloudflare.F[cloudflare.ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleParamsPosition](cloudflare.ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleParamsPositionObject(cloudflare.ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleParamsPositionObject{
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
