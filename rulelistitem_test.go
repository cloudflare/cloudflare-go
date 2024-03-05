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

func TestRuleListItemNew(t *testing.T) {
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
		cloudflare.RuleListItemNewParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Body: cloudflare.F([]cloudflare.RuleListItemNewParamsBody{{
				ASN:     cloudflare.F(int64(5567)),
				Comment: cloudflare.F("Private IP address"),
				Hostname: cloudflare.F(cloudflare.RuleListItemNewParamsBodyHostname{
					URLHostname: cloudflare.F("example.com"),
				}),
				IP: cloudflare.F("10.0.0.1"),
				Redirect: cloudflare.F(cloudflare.RuleListItemNewParamsBodyRedirect{
					IncludeSubdomains:   cloudflare.F(true),
					PreservePathSuffix:  cloudflare.F(true),
					PreserveQueryString: cloudflare.F(true),
					SourceURL:           cloudflare.F("example.com/arch"),
					StatusCode:          cloudflare.F(cloudflare.RuleListItemNewParamsBodyRedirectStatusCode301),
					SubpathMatching:     cloudflare.F(true),
					TargetURL:           cloudflare.F("https://archlinux.org/"),
				}),
			}, {
				ASN:     cloudflare.F(int64(5567)),
				Comment: cloudflare.F("Private IP address"),
				Hostname: cloudflare.F(cloudflare.RuleListItemNewParamsBodyHostname{
					URLHostname: cloudflare.F("example.com"),
				}),
				IP: cloudflare.F("10.0.0.1"),
				Redirect: cloudflare.F(cloudflare.RuleListItemNewParamsBodyRedirect{
					IncludeSubdomains:   cloudflare.F(true),
					PreservePathSuffix:  cloudflare.F(true),
					PreserveQueryString: cloudflare.F(true),
					SourceURL:           cloudflare.F("example.com/arch"),
					StatusCode:          cloudflare.F(cloudflare.RuleListItemNewParamsBodyRedirectStatusCode301),
					SubpathMatching:     cloudflare.F(true),
					TargetURL:           cloudflare.F("https://archlinux.org/"),
				}),
			}, {
				ASN:     cloudflare.F(int64(5567)),
				Comment: cloudflare.F("Private IP address"),
				Hostname: cloudflare.F(cloudflare.RuleListItemNewParamsBodyHostname{
					URLHostname: cloudflare.F("example.com"),
				}),
				IP: cloudflare.F("10.0.0.1"),
				Redirect: cloudflare.F(cloudflare.RuleListItemNewParamsBodyRedirect{
					IncludeSubdomains:   cloudflare.F(true),
					PreservePathSuffix:  cloudflare.F(true),
					PreserveQueryString: cloudflare.F(true),
					SourceURL:           cloudflare.F("example.com/arch"),
					StatusCode:          cloudflare.F(cloudflare.RuleListItemNewParamsBodyRedirectStatusCode301),
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

func TestRuleListItemUpdate(t *testing.T) {
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
		cloudflare.RuleListItemUpdateParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Body: cloudflare.F([]cloudflare.RuleListItemUpdateParamsBody{{
				ASN:     cloudflare.F(int64(5567)),
				Comment: cloudflare.F("Private IP address"),
				Hostname: cloudflare.F(cloudflare.RuleListItemUpdateParamsBodyHostname{
					URLHostname: cloudflare.F("example.com"),
				}),
				IP: cloudflare.F("10.0.0.1"),
				Redirect: cloudflare.F(cloudflare.RuleListItemUpdateParamsBodyRedirect{
					IncludeSubdomains:   cloudflare.F(true),
					PreservePathSuffix:  cloudflare.F(true),
					PreserveQueryString: cloudflare.F(true),
					SourceURL:           cloudflare.F("example.com/arch"),
					StatusCode:          cloudflare.F(cloudflare.RuleListItemUpdateParamsBodyRedirectStatusCode301),
					SubpathMatching:     cloudflare.F(true),
					TargetURL:           cloudflare.F("https://archlinux.org/"),
				}),
			}, {
				ASN:     cloudflare.F(int64(5567)),
				Comment: cloudflare.F("Private IP address"),
				Hostname: cloudflare.F(cloudflare.RuleListItemUpdateParamsBodyHostname{
					URLHostname: cloudflare.F("example.com"),
				}),
				IP: cloudflare.F("10.0.0.1"),
				Redirect: cloudflare.F(cloudflare.RuleListItemUpdateParamsBodyRedirect{
					IncludeSubdomains:   cloudflare.F(true),
					PreservePathSuffix:  cloudflare.F(true),
					PreserveQueryString: cloudflare.F(true),
					SourceURL:           cloudflare.F("example.com/arch"),
					StatusCode:          cloudflare.F(cloudflare.RuleListItemUpdateParamsBodyRedirectStatusCode301),
					SubpathMatching:     cloudflare.F(true),
					TargetURL:           cloudflare.F("https://archlinux.org/"),
				}),
			}, {
				ASN:     cloudflare.F(int64(5567)),
				Comment: cloudflare.F("Private IP address"),
				Hostname: cloudflare.F(cloudflare.RuleListItemUpdateParamsBodyHostname{
					URLHostname: cloudflare.F("example.com"),
				}),
				IP: cloudflare.F("10.0.0.1"),
				Redirect: cloudflare.F(cloudflare.RuleListItemUpdateParamsBodyRedirect{
					IncludeSubdomains:   cloudflare.F(true),
					PreservePathSuffix:  cloudflare.F(true),
					PreserveQueryString: cloudflare.F(true),
					SourceURL:           cloudflare.F("example.com/arch"),
					StatusCode:          cloudflare.F(cloudflare.RuleListItemUpdateParamsBodyRedirectStatusCode301),
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

func TestRuleListItemListWithOptionalParams(t *testing.T) {
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
		cloudflare.RuleListItemListParams{
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

func TestRuleListItemDeleteWithOptionalParams(t *testing.T) {
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
		cloudflare.RuleListItemDeleteParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Items: cloudflare.F([]cloudflare.RuleListItemDeleteParamsItem{{
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

func TestRuleListItemGet(t *testing.T) {
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
