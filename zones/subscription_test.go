// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
	"github.com/cloudflare/cloudflare-go/v5/zones"
)

func TestSubscriptionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Subscriptions.New(context.TODO(), zones.SubscriptionNewParams{
		ZoneID: cloudflare.F("506e3185e9c882d175a2d0cb0093d9f2"),
		Subscription: shared.SubscriptionParam{
			Frequency: cloudflare.F(shared.SubscriptionFrequencyMonthly),
			RatePlan: cloudflare.F(shared.RatePlanParam{
				ID:                cloudflare.F(shared.RatePlanIDFree),
				Currency:          cloudflare.F("USD"),
				ExternallyManaged: cloudflare.F(false),
				IsContract:        cloudflare.F(false),
				PublicName:        cloudflare.F("Business Plan"),
				Scope:             cloudflare.F("zone"),
				Sets:              cloudflare.F([]string{"string"}),
			}),
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

func TestSubscriptionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Subscriptions.Update(context.TODO(), zones.SubscriptionUpdateParams{
		ZoneID: cloudflare.F("506e3185e9c882d175a2d0cb0093d9f2"),
		Subscription: shared.SubscriptionParam{
			Frequency: cloudflare.F(shared.SubscriptionFrequencyMonthly),
			RatePlan: cloudflare.F(shared.RatePlanParam{
				ID:                cloudflare.F(shared.RatePlanIDFree),
				Currency:          cloudflare.F("USD"),
				ExternallyManaged: cloudflare.F(false),
				IsContract:        cloudflare.F(false),
				PublicName:        cloudflare.F("Business Plan"),
				Scope:             cloudflare.F("zone"),
				Sets:              cloudflare.F([]string{"string"}),
			}),
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

func TestSubscriptionGet(t *testing.T) {
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
	_, err := client.Zones.Subscriptions.Get(context.TODO(), zones.SubscriptionGetParams{
		ZoneID: cloudflare.F("506e3185e9c882d175a2d0cb0093d9f2"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
