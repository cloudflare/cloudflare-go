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

func TestPageShieldScriptListWithOptionalParams(t *testing.T) {
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
	_, err := client.PageShield.Scripts.List(context.TODO(), cloudflare.PageShieldScriptListParams{
		ZoneID:              cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Direction:           cloudflare.F(cloudflare.PageShieldScriptListParamsDirectionAsc),
		ExcludeCdnCgi:       cloudflare.F(true),
		ExcludeDuplicates:   cloudflare.F(true),
		ExcludeURLs:         cloudflare.F("blog.cloudflare.com,www.example"),
		Export:              cloudflare.F(cloudflare.PageShieldScriptListParamsExportCsv),
		Hosts:               cloudflare.F("blog.cloudflare.com,www.example*,*cloudflare.com"),
		OrderBy:             cloudflare.F(cloudflare.PageShieldScriptListParamsOrderByFirstSeenAt),
		Page:                cloudflare.F("string"),
		PageURL:             cloudflare.F("example.com/page,*/checkout,example.com/*,*checkout*"),
		PerPage:             cloudflare.F(100.000000),
		PrioritizeMalicious: cloudflare.F(true),
		Status:              cloudflare.F("active,inactive"),
		URLs:                cloudflare.F("blog.cloudflare.com,www.example"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPageShieldScriptGet(t *testing.T) {
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
	_, err := client.PageShield.Scripts.Get(
		context.TODO(),
		"c9ef84a6bf5e47138c75d95e2f933e8f",
		cloudflare.PageShieldScriptGetParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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
