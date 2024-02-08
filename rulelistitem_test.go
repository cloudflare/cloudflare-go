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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Rules.Lists.Items.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"2c0fc9fa937b11eaa1b71c4d701ab86e",
		cloudflare.RuleListItemDeleteParams{
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

func TestRuleListItemListsNewListItems(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Rules.Lists.Items.ListsNewListItems(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"2c0fc9fa937b11eaa1b71c4d701ab86e",
		cloudflare.RuleListItemListsNewListItemsParams{
			Body: cloudflare.F([]cloudflare.RuleListItemListsNewListItemsParamsBody{{
				Asn:     cloudflare.F(int64(5567)),
				Comment: cloudflare.F("Private IP address"),
				Hostname: cloudflare.F(cloudflare.RuleListItemListsNewListItemsParamsBodyHostname{
					URLHostname: cloudflare.F("example.com"),
				}),
				IP: cloudflare.F("10.0.0.1"),
				Redirect: cloudflare.F(cloudflare.RuleListItemListsNewListItemsParamsBodyRedirect{
					IncludeSubdomains:   cloudflare.F(true),
					PreservePathSuffix:  cloudflare.F(true),
					PreserveQueryString: cloudflare.F(true),
					SourceURL:           cloudflare.F("example.com/arch"),
					StatusCode:          cloudflare.F(cloudflare.RuleListItemListsNewListItemsParamsBodyRedirectStatusCode301),
					SubpathMatching:     cloudflare.F(true),
					TargetURL:           cloudflare.F("https://archlinux.org/"),
				}),
			}, {
				Asn:     cloudflare.F(int64(5567)),
				Comment: cloudflare.F("Private IP address"),
				Hostname: cloudflare.F(cloudflare.RuleListItemListsNewListItemsParamsBodyHostname{
					URLHostname: cloudflare.F("example.com"),
				}),
				IP: cloudflare.F("10.0.0.1"),
				Redirect: cloudflare.F(cloudflare.RuleListItemListsNewListItemsParamsBodyRedirect{
					IncludeSubdomains:   cloudflare.F(true),
					PreservePathSuffix:  cloudflare.F(true),
					PreserveQueryString: cloudflare.F(true),
					SourceURL:           cloudflare.F("example.com/arch"),
					StatusCode:          cloudflare.F(cloudflare.RuleListItemListsNewListItemsParamsBodyRedirectStatusCode301),
					SubpathMatching:     cloudflare.F(true),
					TargetURL:           cloudflare.F("https://archlinux.org/"),
				}),
			}, {
				Asn:     cloudflare.F(int64(5567)),
				Comment: cloudflare.F("Private IP address"),
				Hostname: cloudflare.F(cloudflare.RuleListItemListsNewListItemsParamsBodyHostname{
					URLHostname: cloudflare.F("example.com"),
				}),
				IP: cloudflare.F("10.0.0.1"),
				Redirect: cloudflare.F(cloudflare.RuleListItemListsNewListItemsParamsBodyRedirect{
					IncludeSubdomains:   cloudflare.F(true),
					PreservePathSuffix:  cloudflare.F(true),
					PreserveQueryString: cloudflare.F(true),
					SourceURL:           cloudflare.F("example.com/arch"),
					StatusCode:          cloudflare.F(cloudflare.RuleListItemListsNewListItemsParamsBodyRedirectStatusCode301),
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

func TestRuleListItemListsGetListItemsWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Rules.Lists.Items.ListsGetListItems(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"2c0fc9fa937b11eaa1b71c4d701ab86e",
		cloudflare.RuleListItemListsGetListItemsParams{
			Cursor:  cloudflare.F("zzz"),
			PerPage: cloudflare.F(int64(1)),
			Search:  cloudflare.F("1.1.1."),
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

func TestRuleListItemListsUpdateAllListItems(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Rules.Lists.Items.ListsUpdateAllListItems(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"2c0fc9fa937b11eaa1b71c4d701ab86e",
		cloudflare.RuleListItemListsUpdateAllListItemsParams{
			Body: cloudflare.F([]cloudflare.RuleListItemListsUpdateAllListItemsParamsBody{{
				Asn:     cloudflare.F(int64(5567)),
				Comment: cloudflare.F("Private IP address"),
				Hostname: cloudflare.F(cloudflare.RuleListItemListsUpdateAllListItemsParamsBodyHostname{
					URLHostname: cloudflare.F("example.com"),
				}),
				IP: cloudflare.F("10.0.0.1"),
				Redirect: cloudflare.F(cloudflare.RuleListItemListsUpdateAllListItemsParamsBodyRedirect{
					IncludeSubdomains:   cloudflare.F(true),
					PreservePathSuffix:  cloudflare.F(true),
					PreserveQueryString: cloudflare.F(true),
					SourceURL:           cloudflare.F("example.com/arch"),
					StatusCode:          cloudflare.F(cloudflare.RuleListItemListsUpdateAllListItemsParamsBodyRedirectStatusCode301),
					SubpathMatching:     cloudflare.F(true),
					TargetURL:           cloudflare.F("https://archlinux.org/"),
				}),
			}, {
				Asn:     cloudflare.F(int64(5567)),
				Comment: cloudflare.F("Private IP address"),
				Hostname: cloudflare.F(cloudflare.RuleListItemListsUpdateAllListItemsParamsBodyHostname{
					URLHostname: cloudflare.F("example.com"),
				}),
				IP: cloudflare.F("10.0.0.1"),
				Redirect: cloudflare.F(cloudflare.RuleListItemListsUpdateAllListItemsParamsBodyRedirect{
					IncludeSubdomains:   cloudflare.F(true),
					PreservePathSuffix:  cloudflare.F(true),
					PreserveQueryString: cloudflare.F(true),
					SourceURL:           cloudflare.F("example.com/arch"),
					StatusCode:          cloudflare.F(cloudflare.RuleListItemListsUpdateAllListItemsParamsBodyRedirectStatusCode301),
					SubpathMatching:     cloudflare.F(true),
					TargetURL:           cloudflare.F("https://archlinux.org/"),
				}),
			}, {
				Asn:     cloudflare.F(int64(5567)),
				Comment: cloudflare.F("Private IP address"),
				Hostname: cloudflare.F(cloudflare.RuleListItemListsUpdateAllListItemsParamsBodyHostname{
					URLHostname: cloudflare.F("example.com"),
				}),
				IP: cloudflare.F("10.0.0.1"),
				Redirect: cloudflare.F(cloudflare.RuleListItemListsUpdateAllListItemsParamsBodyRedirect{
					IncludeSubdomains:   cloudflare.F(true),
					PreservePathSuffix:  cloudflare.F(true),
					PreserveQueryString: cloudflare.F(true),
					SourceURL:           cloudflare.F("example.com/arch"),
					StatusCode:          cloudflare.F(cloudflare.RuleListItemListsUpdateAllListItemsParamsBodyRedirectStatusCode301),
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
