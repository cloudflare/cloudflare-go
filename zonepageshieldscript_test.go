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

func TestZonePageShieldScriptGet(t *testing.T) {
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
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Zones.PageShields.Scripts.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"c9ef84a6bf5e47138c75d95e2f933e8f",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZonePageShieldScriptPageShieldListPageShieldScriptsWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Zones.PageShields.Scripts.PageShieldListPageShieldScripts(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZonePageShieldScriptPageShieldListPageShieldScriptsParams{
			Direction:           cloudflare.F(cloudflare.ZonePageShieldScriptPageShieldListPageShieldScriptsParamsDirectionAsc),
			ExcludeCdnCgi:       cloudflare.F(true),
			ExcludeURLs:         cloudflare.F("blog.cloudflare.com,www.example"),
			Export:              cloudflare.F(cloudflare.ZonePageShieldScriptPageShieldListPageShieldScriptsParamsExportCsv),
			Hosts:               cloudflare.F("blog.cloudflare.com,www.example"),
			OrderBy:             cloudflare.F(cloudflare.ZonePageShieldScriptPageShieldListPageShieldScriptsParamsOrderByFirstSeenAt),
			Page:                cloudflare.F(2.000000),
			PageURL:             cloudflare.F("example.com/page"),
			PerPage:             cloudflare.F(100.000000),
			PrioritizeMalicious: cloudflare.F(true),
			Status:              cloudflare.F("active,inactive"),
			URLs:                cloudflare.F("blog.cloudflare.com,www.example"),
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
