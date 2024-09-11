// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package managed_transforms_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/managed_transforms"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

func TestManagedTransformList(t *testing.T) {
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
	_, err := client.ManagedTransforms.List(context.TODO(), managed_transforms.ManagedTransformListParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestManagedTransformEdit(t *testing.T) {
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
	_, err := client.ManagedTransforms.Edit(context.TODO(), managed_transforms.ManagedTransformEditParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		ManagedRequestHeaders: cloudflare.F([]managed_transforms.RequestModelParam{{
			ID:      cloudflare.F("add_cf-bot-score_header"),
			Enabled: cloudflare.F(true),
		}, {
			ID:      cloudflare.F("add_cf-bot-score_header"),
			Enabled: cloudflare.F(true),
		}, {
			ID:      cloudflare.F("add_cf-bot-score_header"),
			Enabled: cloudflare.F(true),
		}}),
		ManagedResponseHeaders: cloudflare.F([]managed_transforms.RequestModelParam{{
			ID:      cloudflare.F("add_cf-bot-score_header"),
			Enabled: cloudflare.F(true),
		}, {
			ID:      cloudflare.F("add_cf-bot-score_header"),
			Enabled: cloudflare.F(true),
		}, {
			ID:      cloudflare.F("add_cf-bot-score_header"),
			Enabled: cloudflare.F(true),
		}}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
