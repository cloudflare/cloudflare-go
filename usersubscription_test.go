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

func TestUserSubscriptionUpdateWithOptionalParams(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.User.Subscriptions.Update(
		context.TODO(),
		"506e3185e9c882d175a2d0cb0093d9f2",
		cloudflare.UserSubscriptionUpdateParams{
			App: cloudflare.F(cloudflare.UserSubscriptionUpdateParamsApp{
				InstallID: cloudflare.F("string"),
			}),
			ComponentValues: cloudflare.F([]cloudflare.UserSubscriptionUpdateParamsComponentValue{{
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
			Frequency: cloudflare.F(cloudflare.UserSubscriptionUpdateParamsFrequencyMonthly),
			RatePlan: cloudflare.F(cloudflare.UserSubscriptionUpdateParamsRatePlan{
				Currency:          cloudflare.F("USD"),
				ExternallyManaged: cloudflare.F(false),
				ID:                cloudflare.F[any]("free"),
				IsContract:        cloudflare.F(false),
				PublicName:        cloudflare.F("Business Plan"),
				Scope:             cloudflare.F("zone"),
				Sets:              cloudflare.F([]string{"string", "string", "string"}),
			}),
			Zone: cloudflare.F(cloudflare.UserSubscriptionUpdateParamsZone{}),
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

func TestUserSubscriptionList(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.User.Subscriptions.List(context.TODO())
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserSubscriptionDelete(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.User.Subscriptions.Delete(context.TODO(), "506e3185e9c882d175a2d0cb0093d9f2")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
