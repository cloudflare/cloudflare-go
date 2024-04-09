// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package challenges_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/challenges"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

func TestWidgetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Challenges.Widgets.New(context.TODO(), challenges.WidgetNewParams{
		AccountID:      cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Domains:        cloudflare.F([]challenges.WidgetDomainParam{"203.0.113.1", "cloudflare.com", "blog.example.com"}),
		Mode:           cloudflare.F(challenges.WidgetNewParamsModeInvisible),
		Name:           cloudflare.F("blog.cloudflare.com login form"),
		Direction:      cloudflare.F(challenges.WidgetNewParamsDirectionAsc),
		Order:          cloudflare.F(challenges.WidgetNewParamsOrderID),
		Page:           cloudflare.F(1.000000),
		PerPage:        cloudflare.F(5.000000),
		BotFightMode:   cloudflare.F(false),
		ClearanceLevel: cloudflare.F(challenges.WidgetNewParamsClearanceLevelInteractive),
		Offlabel:       cloudflare.F(false),
		Region:         cloudflare.F(challenges.WidgetNewParamsRegionWorld),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWidgetUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Challenges.Widgets.Update(
		context.TODO(),
		"0x4AAF00AAAABn0R22HWm-YUc",
		challenges.WidgetUpdateParams{
			AccountID:      cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Domains:        cloudflare.F([]challenges.WidgetDomainParam{"203.0.113.1", "cloudflare.com", "blog.example.com"}),
			Mode:           cloudflare.F(challenges.WidgetUpdateParamsModeInvisible),
			Name:           cloudflare.F("blog.cloudflare.com login form"),
			BotFightMode:   cloudflare.F(false),
			ClearanceLevel: cloudflare.F(challenges.WidgetUpdateParamsClearanceLevelInteractive),
			Offlabel:       cloudflare.F(false),
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

func TestWidgetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Challenges.Widgets.List(context.TODO(), challenges.WidgetListParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Direction: cloudflare.F(challenges.WidgetListParamsDirectionAsc),
		Order:     cloudflare.F(challenges.WidgetListParamsOrderID),
		Page:      cloudflare.F(1.000000),
		PerPage:   cloudflare.F(5.000000),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWidgetDelete(t *testing.T) {
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
	_, err := client.Challenges.Widgets.Delete(
		context.TODO(),
		"0x4AAF00AAAABn0R22HWm-YUc",
		challenges.WidgetDeleteParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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

func TestWidgetGet(t *testing.T) {
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
	_, err := client.Challenges.Widgets.Get(
		context.TODO(),
		"0x4AAF00AAAABn0R22HWm-YUc",
		challenges.WidgetGetParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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

func TestWidgetRotateSecretWithOptionalParams(t *testing.T) {
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
	_, err := client.Challenges.Widgets.RotateSecret(
		context.TODO(),
		"0x4AAF00AAAABn0R22HWm-YUc",
		challenges.WidgetRotateSecretParams{
			AccountID:             cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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
