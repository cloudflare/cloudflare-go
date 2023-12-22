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

func TestZoneEmailRoutingRuleGet(t *testing.T) {
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
	_, err := client.Zones.Emails.Routings.Rules.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"a7e6fb77503c41d8a7f3113c6918f10c",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneEmailRoutingRuleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Emails.Routings.Rules.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"a7e6fb77503c41d8a7f3113c6918f10c",
		cloudflare.ZoneEmailRoutingRuleUpdateParams{
			Actions: cloudflare.F([]cloudflare.ZoneEmailRoutingRuleUpdateParamsAction{{
				Type:  cloudflare.F(cloudflare.ZoneEmailRoutingRuleUpdateParamsActionsTypeForward),
				Value: cloudflare.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
			}, {
				Type:  cloudflare.F(cloudflare.ZoneEmailRoutingRuleUpdateParamsActionsTypeForward),
				Value: cloudflare.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
			}, {
				Type:  cloudflare.F(cloudflare.ZoneEmailRoutingRuleUpdateParamsActionsTypeForward),
				Value: cloudflare.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
			}}),
			Matchers: cloudflare.F([]cloudflare.ZoneEmailRoutingRuleUpdateParamsMatcher{{
				Field: cloudflare.F(cloudflare.ZoneEmailRoutingRuleUpdateParamsMatchersFieldTo),
				Type:  cloudflare.F(cloudflare.ZoneEmailRoutingRuleUpdateParamsMatchersTypeLiteral),
				Value: cloudflare.F("test@example.com"),
			}, {
				Field: cloudflare.F(cloudflare.ZoneEmailRoutingRuleUpdateParamsMatchersFieldTo),
				Type:  cloudflare.F(cloudflare.ZoneEmailRoutingRuleUpdateParamsMatchersTypeLiteral),
				Value: cloudflare.F("test@example.com"),
			}, {
				Field: cloudflare.F(cloudflare.ZoneEmailRoutingRuleUpdateParamsMatchersFieldTo),
				Type:  cloudflare.F(cloudflare.ZoneEmailRoutingRuleUpdateParamsMatchersTypeLiteral),
				Value: cloudflare.F("test@example.com"),
			}}),
			Enabled:  cloudflare.F(cloudflare.ZoneEmailRoutingRuleUpdateParamsEnabledTrue),
			Name:     cloudflare.F("Send to user@example.net rule."),
			Priority: cloudflare.F(0.000000),
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

func TestZoneEmailRoutingRuleDelete(t *testing.T) {
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
	_, err := client.Zones.Emails.Routings.Rules.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"a7e6fb77503c41d8a7f3113c6918f10c",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Emails.Routings.Rules.EmailRoutingRoutingRulesNewRoutingRule(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParams{
			Actions: cloudflare.F([]cloudflare.ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsAction{{
				Type:  cloudflare.F(cloudflare.ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsActionsTypeForward),
				Value: cloudflare.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
			}, {
				Type:  cloudflare.F(cloudflare.ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsActionsTypeForward),
				Value: cloudflare.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
			}, {
				Type:  cloudflare.F(cloudflare.ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsActionsTypeForward),
				Value: cloudflare.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
			}}),
			Matchers: cloudflare.F([]cloudflare.ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatcher{{
				Field: cloudflare.F(cloudflare.ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatchersFieldTo),
				Type:  cloudflare.F(cloudflare.ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatchersTypeLiteral),
				Value: cloudflare.F("test@example.com"),
			}, {
				Field: cloudflare.F(cloudflare.ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatchersFieldTo),
				Type:  cloudflare.F(cloudflare.ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatchersTypeLiteral),
				Value: cloudflare.F("test@example.com"),
			}, {
				Field: cloudflare.F(cloudflare.ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatchersFieldTo),
				Type:  cloudflare.F(cloudflare.ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatchersTypeLiteral),
				Value: cloudflare.F("test@example.com"),
			}}),
			Enabled:  cloudflare.F(cloudflare.ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsEnabledTrue),
			Name:     cloudflare.F("Send to user@example.net rule."),
			Priority: cloudflare.F(0.000000),
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

func TestZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Emails.Routings.Rules.EmailRoutingRoutingRulesListRoutingRules(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParams{
			Enabled: cloudflare.F(cloudflare.ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParamsEnabledTrue),
			Page:    cloudflare.F(1.000000),
			PerPage: cloudflare.F(5.000000),
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
