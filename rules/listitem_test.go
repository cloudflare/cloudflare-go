// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rules_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/rules"
)

func TestListItemNew(t *testing.T) {
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
	_, err := client.Rules.Lists.Items.New(
		context.TODO(),
		"2c0fc9fa937b11eaa1b71c4d701ab86e",
		rules.ListItemNewParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Body: cloudflare.F([]rules.ListItemNewParamsBody{{
				ASN:     cloudflare.F(int64(5567)),
				Comment: cloudflare.F("Private IP address"),
				Hostname: cloudflare.F(rules.HostnameParam{
					URLHostname: cloudflare.F("example.com"),
				}),
				IP: cloudflare.F("10.0.0.1"),
				Redirect: cloudflare.F(rules.RedirectParam{
					IncludeSubdomains:   cloudflare.F(true),
					PreservePathSuffix:  cloudflare.F(true),
					PreserveQueryString: cloudflare.F(true),
					SourceURL:           cloudflare.F("example.com/arch"),
					StatusCode:          cloudflare.F(rules.RedirectStatusCode301),
					SubpathMatching:     cloudflare.F(true),
					TargetURL:           cloudflare.F("https://archlinux.org/"),
				}),
			}, {
				ASN:     cloudflare.F(int64(5567)),
				Comment: cloudflare.F("Private IP address"),
				Hostname: cloudflare.F(rules.HostnameParam{
					URLHostname: cloudflare.F("example.com"),
				}),
				IP: cloudflare.F("10.0.0.1"),
				Redirect: cloudflare.F(rules.RedirectParam{
					IncludeSubdomains:   cloudflare.F(true),
					PreservePathSuffix:  cloudflare.F(true),
					PreserveQueryString: cloudflare.F(true),
					SourceURL:           cloudflare.F("example.com/arch"),
					StatusCode:          cloudflare.F(rules.RedirectStatusCode301),
					SubpathMatching:     cloudflare.F(true),
					TargetURL:           cloudflare.F("https://archlinux.org/"),
				}),
			}, {
				ASN:     cloudflare.F(int64(5567)),
				Comment: cloudflare.F("Private IP address"),
				Hostname: cloudflare.F(rules.HostnameParam{
					URLHostname: cloudflare.F("example.com"),
				}),
				IP: cloudflare.F("10.0.0.1"),
				Redirect: cloudflare.F(rules.RedirectParam{
					IncludeSubdomains:   cloudflare.F(true),
					PreservePathSuffix:  cloudflare.F(true),
					PreserveQueryString: cloudflare.F(true),
					SourceURL:           cloudflare.F("example.com/arch"),
					StatusCode:          cloudflare.F(rules.RedirectStatusCode301),
					SubpathMatching:     cloudflare.F(true),
					TargetURL:           cloudflare.F("https://archlinux.org/"),
				}),
			}}),
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

func TestListItemUpdate(t *testing.T) {
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
	_, err := client.Rules.Lists.Items.Update(
		context.TODO(),
		"2c0fc9fa937b11eaa1b71c4d701ab86e",
		rules.ListItemUpdateParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Body: cloudflare.F([]rules.ListItemUpdateParamsBody{{
				ASN:     cloudflare.F(int64(5567)),
				Comment: cloudflare.F("Private IP address"),
				Hostname: cloudflare.F(rules.HostnameParam{
					URLHostname: cloudflare.F("example.com"),
				}),
				IP: cloudflare.F("10.0.0.1"),
				Redirect: cloudflare.F(rules.RedirectParam{
					IncludeSubdomains:   cloudflare.F(true),
					PreservePathSuffix:  cloudflare.F(true),
					PreserveQueryString: cloudflare.F(true),
					SourceURL:           cloudflare.F("example.com/arch"),
					StatusCode:          cloudflare.F(rules.RedirectStatusCode301),
					SubpathMatching:     cloudflare.F(true),
					TargetURL:           cloudflare.F("https://archlinux.org/"),
				}),
			}, {
				ASN:     cloudflare.F(int64(5567)),
				Comment: cloudflare.F("Private IP address"),
				Hostname: cloudflare.F(rules.HostnameParam{
					URLHostname: cloudflare.F("example.com"),
				}),
				IP: cloudflare.F("10.0.0.1"),
				Redirect: cloudflare.F(rules.RedirectParam{
					IncludeSubdomains:   cloudflare.F(true),
					PreservePathSuffix:  cloudflare.F(true),
					PreserveQueryString: cloudflare.F(true),
					SourceURL:           cloudflare.F("example.com/arch"),
					StatusCode:          cloudflare.F(rules.RedirectStatusCode301),
					SubpathMatching:     cloudflare.F(true),
					TargetURL:           cloudflare.F("https://archlinux.org/"),
				}),
			}, {
				ASN:     cloudflare.F(int64(5567)),
				Comment: cloudflare.F("Private IP address"),
				Hostname: cloudflare.F(rules.HostnameParam{
					URLHostname: cloudflare.F("example.com"),
				}),
				IP: cloudflare.F("10.0.0.1"),
				Redirect: cloudflare.F(rules.RedirectParam{
					IncludeSubdomains:   cloudflare.F(true),
					PreservePathSuffix:  cloudflare.F(true),
					PreserveQueryString: cloudflare.F(true),
					SourceURL:           cloudflare.F("example.com/arch"),
					StatusCode:          cloudflare.F(rules.RedirectStatusCode301),
					SubpathMatching:     cloudflare.F(true),
					TargetURL:           cloudflare.F("https://archlinux.org/"),
				}),
			}}),
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

func TestListItemListWithOptionalParams(t *testing.T) {
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
	_, err := client.Rules.Lists.Items.List(
		context.TODO(),
		"2c0fc9fa937b11eaa1b71c4d701ab86e",
		rules.ListItemListParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Cursor:    cloudflare.F("zzz"),
			PerPage:   cloudflare.F(int64(1)),
			Search:    cloudflare.F("1.1.1."),
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

func TestListItemDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Rules.Lists.Items.Delete(
		context.TODO(),
		"2c0fc9fa937b11eaa1b71c4d701ab86e",
		rules.ListItemDeleteParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Items: cloudflare.F([]rules.ListItemDeleteParamsItem{{
				ID: cloudflare.F("34b12448945f11eaa1b71c4d701ab86e"),
			}}),
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

func TestListItemGet(t *testing.T) {
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
	_, err := client.Rules.Lists.Items.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"2c0fc9fa937b11eaa1b71c4d701ab86e",
		"34b12448945f11eaa1b71c4d701ab86e",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
