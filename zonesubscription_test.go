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

func TestZoneSubscriptionNewWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Zones.Subscriptions.New(
		context.TODO(),
		"506e3185e9c882d175a2d0cb0093d9f2",
		cloudflare.ZoneSubscriptionNewParams{
			App: cloudflare.F(cloudflare.ZoneSubscriptionNewParamsApp{
				InstallID: cloudflare.F("string"),
			}),
			ComponentValues: cloudflare.F([]cloudflare.ZoneSubscriptionNewParamsComponentValue{{
				Default: cloudflare.F(5.000000),
				Name:    cloudflare.F("page_rules"),
				Price:   cloudflare.F(5.000000),
				Value:   cloudflare.F(20.000000),
			}, {
				Default: cloudflare.F(5.000000),
				Name:    cloudflare.F("page_rules"),
				Price:   cloudflare.F(5.000000),
				Value:   cloudflare.F(20.000000),
			}, {
				Default: cloudflare.F(5.000000),
				Name:    cloudflare.F("page_rules"),
				Price:   cloudflare.F(5.000000),
				Value:   cloudflare.F(20.000000),
			}}),
			Frequency: cloudflare.F(cloudflare.ZoneSubscriptionNewParamsFrequencyMonthly),
			RatePlan: cloudflare.F(cloudflare.ZoneSubscriptionNewParamsRatePlan{
				Currency:          cloudflare.F("USD"),
				ExternallyManaged: cloudflare.F(false),
				ID:                cloudflare.F[any]("free"),
				IsContract:        cloudflare.F(false),
				PublicName:        cloudflare.F("Business Plan"),
				Scope:             cloudflare.F("zone"),
				Sets:              cloudflare.F([]string{"string", "string", "string"}),
			}),
			Zone: cloudflare.F(cloudflare.ZoneSubscriptionNewParamsZone{}),
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

func TestZoneSubscriptionList(t *testing.T) {
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
	)
	_, err := client.Zones.Subscriptions.List(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneSubscriptionGet(t *testing.T) {
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
	)
	_, err := client.Zones.Subscriptions.Get(context.TODO(), "506e3185e9c882d175a2d0cb0093d9f2")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
