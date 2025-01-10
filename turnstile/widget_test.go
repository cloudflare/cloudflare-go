// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package turnstile_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/turnstile"
)

func TestWidgetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Turnstile.Widgets.New(context.TODO(), turnstile.WidgetNewParams{
		AccountID:      cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Domains:        cloudflare.F([]turnstile.WidgetDomainParam{"203.0.113.1", "cloudflare.com", "blog.example.com"}),
		Mode:           cloudflare.F(turnstile.WidgetNewParamsModeNonInteractive),
		Name:           cloudflare.F("blog.cloudflare.com login form"),
		Direction:      cloudflare.F(turnstile.WidgetNewParamsDirectionAsc),
		Order:          cloudflare.F(turnstile.WidgetNewParamsOrderID),
		Page:           cloudflare.F(1.000000),
		PerPage:        cloudflare.F(5.000000),
		BotFightMode:   cloudflare.F(false),
		ClearanceLevel: cloudflare.F(turnstile.WidgetNewParamsClearanceLevelNoClearance),
		EphemeralID:    cloudflare.F(false),
		Offlabel:       cloudflare.F(false),
		Region:         cloudflare.F(turnstile.WidgetNewParamsRegionWorld),
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
	_, err := client.Turnstile.Widgets.Update(
		context.TODO(),
		"0x4AAF00AAAABn0R22HWm-YUc",
		turnstile.WidgetUpdateParams{
			AccountID:      cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Domains:        cloudflare.F([]turnstile.WidgetDomainParam{"203.0.113.1", "cloudflare.com", "blog.example.com"}),
			Mode:           cloudflare.F(turnstile.WidgetUpdateParamsModeNonInteractive),
			Name:           cloudflare.F("blog.cloudflare.com login form"),
			BotFightMode:   cloudflare.F(false),
			ClearanceLevel: cloudflare.F(turnstile.WidgetUpdateParamsClearanceLevelNoClearance),
			EphemeralID:    cloudflare.F(false),
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
	_, err := client.Turnstile.Widgets.List(context.TODO(), turnstile.WidgetListParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Direction: cloudflare.F(turnstile.WidgetListParamsDirectionAsc),
		Order:     cloudflare.F(turnstile.WidgetListParamsOrderID),
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
	_, err := client.Turnstile.Widgets.Delete(
		context.TODO(),
		"0x4AAF00AAAABn0R22HWm-YUc",
		turnstile.WidgetDeleteParams{
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
	_, err := client.Turnstile.Widgets.Get(
		context.TODO(),
		"0x4AAF00AAAABn0R22HWm-YUc",
		turnstile.WidgetGetParams{
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
	_, err := client.Turnstile.Widgets.RotateSecret(
		context.TODO(),
		"0x4AAF00AAAABn0R22HWm-YUc",
		turnstile.WidgetRotateSecretParams{
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
