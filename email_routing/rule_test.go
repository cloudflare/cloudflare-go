// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/email_routing"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func TestRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailRouting.Rules.New(context.TODO(), email_routing.RuleNewParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Actions: cloudflare.F([]email_routing.ActionParam{{
			Type:  cloudflare.F(email_routing.ActionTypeDrop),
			Value: cloudflare.F([]string{"destinationaddress@example.net"}),
		}}),
		Matchers: cloudflare.F([]email_routing.MatcherParam{{
			Field: cloudflare.F(email_routing.MatcherFieldTo),
			Type:  cloudflare.F(email_routing.MatcherTypeLiteral),
			Value: cloudflare.F("test@example.com"),
		}}),
		Enabled:  cloudflare.F(email_routing.RuleNewParamsEnabledTrue),
		Name:     cloudflare.F("Send to user@example.net rule."),
		Priority: cloudflare.F(0.000000),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRuleUpdateWithOptionalParams(t *testing.T) {
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
		"a7e6fb77503c41d8a7f3113c6918f10c",
		email_routing.RuleUpdateParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Actions: cloudflare.F([]email_routing.ActionParam{{
				Type:  cloudflare.F(email_routing.ActionTypeDrop),
				Value: cloudflare.F([]string{"destinationaddress@example.net"}),
			}}),
			Matchers: cloudflare.F([]email_routing.MatcherParam{{
				Field: cloudflare.F(email_routing.MatcherFieldTo),
				Type:  cloudflare.F(email_routing.MatcherTypeLiteral),
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
	_, err := client.EmailRouting.Rules.List(context.TODO(), email_routing.RuleListParams{
		ZoneID:  cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Enabled: cloudflare.F(email_routing.RuleListParamsEnabledTrue),
		Page:    cloudflare.F(1.000000),
		PerPage: cloudflare.F(5.000000),
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.EmailRouting.Rules.Delete(
		context.TODO(),
		"a7e6fb77503c41d8a7f3113c6918f10c",
		email_routing.RuleDeleteParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.EmailRouting.Rules.Get(
		context.TODO(),
		"a7e6fb77503c41d8a7f3113c6918f10c",
		email_routing.RuleGetParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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
