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

func TestAccountChallengeWidgetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Challenges.Widgets.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.AccountChallengeWidgetNewParams{
			Domains:        cloudflare.F([]string{"203.0.113.1", "cloudflare.com", "blog.example.com"}),
			Mode:           cloudflare.F(cloudflare.AccountChallengeWidgetNewParamsModeInvisible),
			Name:           cloudflare.F("blog.cloudflare.com login form"),
			Direction:      cloudflare.F(cloudflare.AccountChallengeWidgetNewParamsDirectionAsc),
			Order:          cloudflare.F(cloudflare.AccountChallengeWidgetNewParamsOrderID),
			Page:           cloudflare.F(1.000000),
			PerPage:        cloudflare.F(5.000000),
			BotFightMode:   cloudflare.F(true),
			ClearanceLevel: cloudflare.F(cloudflare.AccountChallengeWidgetNewParamsClearanceLevelInteractive),
			Offlabel:       cloudflare.F(true),
			Region:         cloudflare.F(cloudflare.AccountChallengeWidgetNewParamsRegionWorld),
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

func TestAccountChallengeWidgetGet(t *testing.T) {
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
	_, err := client.Accounts.Challenges.Widgets.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"0x4AAF00AAAABn0R22HWm-YUc",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountChallengeWidgetUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Challenges.Widgets.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"0x4AAF00AAAABn0R22HWm-YUc",
		cloudflare.AccountChallengeWidgetUpdateParams{
			Domains:        cloudflare.F([]string{"203.0.113.1", "cloudflare.com", "blog.example.com"}),
			Mode:           cloudflare.F(cloudflare.AccountChallengeWidgetUpdateParamsModeInvisible),
			Name:           cloudflare.F("blog.cloudflare.com login form"),
			BotFightMode:   cloudflare.F(true),
			ClearanceLevel: cloudflare.F(cloudflare.AccountChallengeWidgetUpdateParamsClearanceLevelInteractive),
			Offlabel:       cloudflare.F(true),
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

func TestAccountChallengeWidgetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Challenges.Widgets.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.AccountChallengeWidgetListParams{
			Direction: cloudflare.F(cloudflare.AccountChallengeWidgetListParamsDirectionAsc),
			Order:     cloudflare.F(cloudflare.AccountChallengeWidgetListParamsOrderID),
			Page:      cloudflare.F(1.000000),
			PerPage:   cloudflare.F(5.000000),
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

func TestAccountChallengeWidgetDelete(t *testing.T) {
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
	_, err := client.Accounts.Challenges.Widgets.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"0x4AAF00AAAABn0R22HWm-YUc",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountChallengeWidgetRotateSecretWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Challenges.Widgets.RotateSecret(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"0x4AAF00AAAABn0R22HWm-YUc",
		cloudflare.AccountChallengeWidgetRotateSecretParams{
			InvalidateImmediately: cloudflare.F(true),
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
