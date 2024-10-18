// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/page_shield"
)

func TestScriptListWithOptionalParams(t *testing.T) {
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
	_, err := client.PageShield.Scripts.List(context.TODO(), page_shield.ScriptListParams{
		ZoneID:              cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Direction:           cloudflare.F(page_shield.ScriptListParamsDirectionAsc),
		ExcludeCDNCGI:       cloudflare.F(true),
		ExcludeDuplicates:   cloudflare.F(true),
		ExcludeURLs:         cloudflare.F("blog.cloudflare.com,www.example"),
		Export:              cloudflare.F(page_shield.ScriptListParamsExportCsv),
		Hosts:               cloudflare.F("blog.cloudflare.com,www.example*,*cloudflare.com"),
		OrderBy:             cloudflare.F(page_shield.ScriptListParamsOrderByFirstSeenAt),
		Page:                cloudflare.F("2"),
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

func TestScriptGet(t *testing.T) {
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
		"023e105f4ecef8ad9ca31a8372d0c353",
		page_shield.ScriptGetParams{
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
