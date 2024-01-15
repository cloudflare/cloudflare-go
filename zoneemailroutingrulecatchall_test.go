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

func TestZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRule(t *testing.T) {
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
	_, err := client.Zones.Emails.Routings.Rules.CatchAlls.EmailRoutingRoutingRulesGetCatchAllRule(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Emails.Routings.Rules.CatchAlls.EmailRoutingRoutingRulesUpdateCatchAllRule(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParams{
			Actions: cloudflare.F([]cloudflare.ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsAction{{
				Type:  cloudflare.F(cloudflare.ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsActionsTypeForward),
				Value: cloudflare.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
			}, {
				Type:  cloudflare.F(cloudflare.ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsActionsTypeForward),
				Value: cloudflare.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
			}, {
				Type:  cloudflare.F(cloudflare.ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsActionsTypeForward),
				Value: cloudflare.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
			}}),
			Matchers: cloudflare.F([]cloudflare.ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsMatcher{{
				Type: cloudflare.F(cloudflare.ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsMatchersTypeAll),
			}, {
				Type: cloudflare.F(cloudflare.ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsMatchersTypeAll),
			}, {
				Type: cloudflare.F(cloudflare.ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsMatchersTypeAll),
			}}),
			Enabled: cloudflare.F(cloudflare.ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsEnabledTrue),
			Name:    cloudflare.F("Send to user@example.net rule."),
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
