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

func TestCookieListWithOptionalParams(t *testing.T) {
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
	_, err := client.PageShield.Cookies.List(context.TODO(), page_shield.CookieListParams{
		ZoneID:    cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Direction: cloudflare.F(page_shield.CookieListParamsDirectionAsc),
		Domain:    cloudflare.F("example.com"),
		Export:    cloudflare.F(page_shield.CookieListParamsExportCsv),
		Hosts:     cloudflare.F("blog.cloudflare.com,www.example*,*cloudflare.com"),
		HTTPOnly:  cloudflare.F(true),
		Name:      cloudflare.F("session_id"),
		OrderBy:   cloudflare.F(page_shield.CookieListParamsOrderByFirstSeenAt),
		Page:      cloudflare.F("2"),
		PageURL:   cloudflare.F("example.com/page,*/checkout,example.com/*,*checkout*"),
		Path:      cloudflare.F("/"),
		PerPage:   cloudflare.F(100.000000),
		SameSite:  cloudflare.F(page_shield.CookieListParamsSameSiteLax),
		Secure:    cloudflare.F(true),
		Type:      cloudflare.F(page_shield.CookieListParamsTypeFirstParty),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCookieGet(t *testing.T) {
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
	_, err := client.PageShield.Cookies.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		page_shield.CookieGetParams{
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
