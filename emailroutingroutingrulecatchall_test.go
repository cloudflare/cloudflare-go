// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
)

func TestEmailRoutingRoutingRuleCatchAllUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailRouting.Routing.Rules.CatchAlls.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.EmailRoutingRoutingRuleCatchAllUpdateParams{
			Actions: cloudflare.F([]cloudflare.EmailRoutingRoutingRuleCatchAllUpdateParamsAction{{
				Type:  cloudflare.F(cloudflare.EmailRoutingRoutingRuleCatchAllUpdateParamsActionsTypeForward),
				Value: cloudflare.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
			}, {
				Type:  cloudflare.F(cloudflare.EmailRoutingRoutingRuleCatchAllUpdateParamsActionsTypeForward),
				Value: cloudflare.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
			}, {
				Type:  cloudflare.F(cloudflare.EmailRoutingRoutingRuleCatchAllUpdateParamsActionsTypeForward),
				Value: cloudflare.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
			}}),
			Matchers: cloudflare.F([]cloudflare.EmailRoutingRoutingRuleCatchAllUpdateParamsMatcher{{
				Type: cloudflare.F(cloudflare.EmailRoutingRoutingRuleCatchAllUpdateParamsMatchersTypeAll),
			}, {
				Type: cloudflare.F(cloudflare.EmailRoutingRoutingRuleCatchAllUpdateParamsMatchersTypeAll),
			}, {
				Type: cloudflare.F(cloudflare.EmailRoutingRoutingRuleCatchAllUpdateParamsMatchersTypeAll),
			}}),
			Enabled: cloudflare.F(cloudflare.EmailRoutingRoutingRuleCatchAllUpdateParamsEnabledTrue),
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

func TestEmailRoutingRoutingRuleCatchAllGet(t *testing.T) {
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
	_, err := client.EmailRouting.Routing.Rules.CatchAlls.Get(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
