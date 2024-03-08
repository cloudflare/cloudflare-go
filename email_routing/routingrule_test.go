// File generated from our OpenAPI spec by Stainless.

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

func TestRoutingRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailRouting.Routing.Rules.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		email_routing.RoutingRuleNewParams{
			Actions: cloudflare.F([]email_routing.RoutingRuleNewParamsAction{{
				Type:  cloudflare.F(email_routing.RoutingRuleNewParamsActionsTypeForward),
				Value: cloudflare.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
			}, {
				Type:  cloudflare.F(email_routing.RoutingRuleNewParamsActionsTypeForward),
				Value: cloudflare.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
			}, {
				Type:  cloudflare.F(email_routing.RoutingRuleNewParamsActionsTypeForward),
				Value: cloudflare.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
			}}),
			Matchers: cloudflare.F([]email_routing.RoutingRuleNewParamsMatcher{{
				Field: cloudflare.F(email_routing.RoutingRuleNewParamsMatchersFieldTo),
				Type:  cloudflare.F(email_routing.RoutingRuleNewParamsMatchersTypeLiteral),
				Value: cloudflare.F("test@example.com"),
			}, {
				Field: cloudflare.F(email_routing.RoutingRuleNewParamsMatchersFieldTo),
				Type:  cloudflare.F(email_routing.RoutingRuleNewParamsMatchersTypeLiteral),
				Value: cloudflare.F("test@example.com"),
			}, {
				Field: cloudflare.F(email_routing.RoutingRuleNewParamsMatchersFieldTo),
				Type:  cloudflare.F(email_routing.RoutingRuleNewParamsMatchersTypeLiteral),
				Value: cloudflare.F("test@example.com"),
			}}),
			Enabled:  cloudflare.F(email_routing.RoutingRuleNewParamsEnabledTrue),
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

func TestRoutingRuleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailRouting.Routing.Rules.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"a7e6fb77503c41d8a7f3113c6918f10c",
		email_routing.RoutingRuleUpdateParams{
			Actions: cloudflare.F([]email_routing.RoutingRuleUpdateParamsAction{{
				Type:  cloudflare.F(email_routing.RoutingRuleUpdateParamsActionsTypeForward),
				Value: cloudflare.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
			}, {
				Type:  cloudflare.F(email_routing.RoutingRuleUpdateParamsActionsTypeForward),
				Value: cloudflare.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
			}, {
				Type:  cloudflare.F(email_routing.RoutingRuleUpdateParamsActionsTypeForward),
				Value: cloudflare.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
			}}),
			Matchers: cloudflare.F([]email_routing.RoutingRuleUpdateParamsMatcher{{
				Field: cloudflare.F(email_routing.RoutingRuleUpdateParamsMatchersFieldTo),
				Type:  cloudflare.F(email_routing.RoutingRuleUpdateParamsMatchersTypeLiteral),
				Value: cloudflare.F("test@example.com"),
			}, {
				Field: cloudflare.F(email_routing.RoutingRuleUpdateParamsMatchersFieldTo),
				Type:  cloudflare.F(email_routing.RoutingRuleUpdateParamsMatchersTypeLiteral),
				Value: cloudflare.F("test@example.com"),
			}, {
				Field: cloudflare.F(email_routing.RoutingRuleUpdateParamsMatchersFieldTo),
				Type:  cloudflare.F(email_routing.RoutingRuleUpdateParamsMatchersTypeLiteral),
				Value: cloudflare.F("test@example.com"),
			}}),
			Enabled:  cloudflare.F(email_routing.RoutingRuleUpdateParamsEnabledTrue),
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

func TestRoutingRuleListWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailRouting.Routing.Rules.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		email_routing.RoutingRuleListParams{
			Enabled: cloudflare.F(email_routing.RoutingRuleListParamsEnabledTrue),
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

func TestRoutingRuleDelete(t *testing.T) {
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
	_, err := client.EmailRouting.Routing.Rules.Delete(
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

func TestRoutingRuleGet(t *testing.T) {
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
	_, err := client.EmailRouting.Routing.Rules.Get(
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
