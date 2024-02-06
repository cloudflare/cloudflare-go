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

func TestEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRule(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Emails.Routings.Rules.CatchAlls.EmailRoutingRoutingRulesGetCatchAllRule(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Emails.Routings.Rules.CatchAlls.EmailRoutingRoutingRulesUpdateCatchAllRule(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParams{
			Actions: cloudflare.F([]cloudflare.EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsAction{{
				Type:  cloudflare.F(cloudflare.EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsActionsTypeForward),
				Value: cloudflare.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
			}, {
				Type:  cloudflare.F(cloudflare.EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsActionsTypeForward),
				Value: cloudflare.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
			}, {
				Type:  cloudflare.F(cloudflare.EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsActionsTypeForward),
				Value: cloudflare.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
			}}),
			Matchers: cloudflare.F([]cloudflare.EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsMatcher{{
				Type: cloudflare.F(cloudflare.EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsMatchersTypeAll),
			}, {
				Type: cloudflare.F(cloudflare.EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsMatchersTypeAll),
			}, {
				Type: cloudflare.F(cloudflare.EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsMatchersTypeAll),
			}}),
			Enabled: cloudflare.F(cloudflare.EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsEnabledTrue),
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
