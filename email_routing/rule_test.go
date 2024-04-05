// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/email_routing"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
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
	_, err := client.EmailRouting.Rules.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		email_routing.RuleNewParams{
			Actions: cloudflare.F([]email_routing.ActionParam{{
				Type:  cloudflare.F(email_routing.ActionTypeForward),
				Value: cloudflare.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
			}, {
				Type:  cloudflare.F(email_routing.ActionTypeForward),
				Value: cloudflare.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
			}, {
				Type:  cloudflare.F(email_routing.ActionTypeForward),
				Value: cloudflare.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
			}}),
			Matchers: cloudflare.F([]email_routing.MatcherParam{{
				Field: cloudflare.F(email_routing.MatcherFieldTo),
				Type:  cloudflare.F(email_routing.MatcherTypeLiteral),
				Value: cloudflare.F("test@example.com"),
			}, {
				Field: cloudflare.F(email_routing.MatcherFieldTo),
				Type:  cloudflare.F(email_routing.MatcherTypeLiteral),
				Value: cloudflare.F("test@example.com"),
			}, {
				Field: cloudflare.F(email_routing.MatcherFieldTo),
				Type:  cloudflare.F(email_routing.MatcherTypeLiteral),
				Value: cloudflare.F("test@example.com"),
			}}),
			Enabled:  cloudflare.F(email_routing.RuleNewParamsEnabledTrue),
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

func TestRuleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailRouting.Rules.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"a7e6fb77503c41d8a7f3113c6918f10c",
		email_routing.RuleUpdateParams{
			Actions: cloudflare.F([]email_routing.ActionItemParam{{
				Type:  cloudflare.F(email_routing.ActionItemTypeForward),
				Value: cloudflare.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
			}, {
				Type:  cloudflare.F(email_routing.ActionItemTypeForward),
				Value: cloudflare.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
			}, {
				Type:  cloudflare.F(email_routing.ActionItemTypeForward),
				Value: cloudflare.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
			}}),
			Matchers: cloudflare.F([]email_routing.MatcherItemParam{{
				Field: cloudflare.F(email_routing.MatcherItemFieldTo),
				Type:  cloudflare.F(email_routing.MatcherItemTypeLiteral),
				Value: cloudflare.F("test@example.com"),
			}, {
				Field: cloudflare.F(email_routing.MatcherItemFieldTo),
				Type:  cloudflare.F(email_routing.MatcherItemTypeLiteral),
				Value: cloudflare.F("test@example.com"),
			}, {
				Field: cloudflare.F(email_routing.MatcherItemFieldTo),
				Type:  cloudflare.F(email_routing.MatcherItemTypeLiteral),
				Value: cloudflare.F("test@example.com"),
			}}),
			Enabled:  cloudflare.F(email_routing.RuleUpdateParamsEnabledTrue),
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

func TestRuleListWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailRouting.Rules.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		email_routing.RuleListParams{
			Enabled: cloudflare.F(email_routing.RuleListParamsEnabledTrue),
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

func TestRuleDelete(t *testing.T) {
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
	_, err := client.EmailRouting.Rules.Delete(
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

func TestRuleGet(t *testing.T) {
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
	_, err := client.EmailRouting.Rules.Get(
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
