// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package managed_headers_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/managed_headers"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

func TestManagedHeaderList(t *testing.T) {
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
	_, err := client.ManagedHeaders.List(context.TODO(), managed_headers.ManagedHeaderListParams{
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

func TestManagedHeaderEdit(t *testing.T) {
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
	_, err := client.ManagedHeaders.Edit(context.TODO(), managed_headers.ManagedHeaderEditParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		ManagedRequestHeaders: cloudflare.F([]managed_headers.RequestListItemParam{{
			Enabled: cloudflare.F(true),
			ID:      cloudflare.F("add_cf-bot-score_header"),
		}, {
			Enabled: cloudflare.F(true),
			ID:      cloudflare.F("add_cf-bot-score_header"),
		}, {
			Enabled: cloudflare.F(true),
			ID:      cloudflare.F("add_cf-bot-score_header"),
		}}),
		ManagedResponseHeaders: cloudflare.F([]managed_headers.RequestListItemParam{{
			Enabled: cloudflare.F(true),
			ID:      cloudflare.F("add_cf-bot-score_header"),
		}, {
			Enabled: cloudflare.F(true),
			ID:      cloudflare.F("add_cf-bot-score_header"),
		}, {
			Enabled: cloudflare.F(true),
			ID:      cloudflare.F("add_cf-bot-score_header"),
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
