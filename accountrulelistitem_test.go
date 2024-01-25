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

func TestAccountRuleListItemGet(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Rules.Lists.Items.Get(
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

func TestAccountRuleListItemDeleteWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Rules.Lists.Items.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"2c0fc9fa937b11eaa1b71c4d701ab86e",
		cloudflare.AccountRuleListItemDeleteParams{
			Items: cloudflare.F([]cloudflare.AccountRuleListItemDeleteParamsItem{{
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

func TestAccountRuleListItemListsNewListItems(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Rules.Lists.Items.ListsNewListItems(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"2c0fc9fa937b11eaa1b71c4d701ab86e",
		cloudflare.AccountRuleListItemListsNewListItemsParams{
			Body: cloudflare.F([]cloudflare.AccountRuleListItemListsNewListItemsParamsBody{{
				ASN:     cloudflare.F(int64(5567)),
				Comment: cloudflare.F("Private IP address"),
				Hostname: cloudflare.F(cloudflare.AccountRuleListItemListsNewListItemsParamsBodyHostname{
					URLHostname: cloudflare.F("example.com"),
				}),
				IP: cloudflare.F("10.0.0.1"),
				Redirect: cloudflare.F(cloudflare.AccountRuleListItemListsNewListItemsParamsBodyRedirect{
					IncludeSubdomains:   cloudflare.F(true),
					PreservePathSuffix:  cloudflare.F(true),
					PreserveQueryString: cloudflare.F(true),
					SourceURL:           cloudflare.F("example.com/arch"),
					StatusCode:          cloudflare.F(cloudflare.AccountRuleListItemListsNewListItemsParamsBodyRedirectStatusCode301),
					SubpathMatching:     cloudflare.F(true),
					TargetURL:           cloudflare.F("https://archlinux.org/"),
				}),
			}, {
				ASN:     cloudflare.F(int64(5567)),
				Comment: cloudflare.F("Private IP address"),
				Hostname: cloudflare.F(cloudflare.AccountRuleListItemListsNewListItemsParamsBodyHostname{
					URLHostname: cloudflare.F("example.com"),
				}),
				IP: cloudflare.F("10.0.0.1"),
				Redirect: cloudflare.F(cloudflare.AccountRuleListItemListsNewListItemsParamsBodyRedirect{
					IncludeSubdomains:   cloudflare.F(true),
					PreservePathSuffix:  cloudflare.F(true),
					PreserveQueryString: cloudflare.F(true),
					SourceURL:           cloudflare.F("example.com/arch"),
					StatusCode:          cloudflare.F(cloudflare.AccountRuleListItemListsNewListItemsParamsBodyRedirectStatusCode301),
					SubpathMatching:     cloudflare.F(true),
					TargetURL:           cloudflare.F("https://archlinux.org/"),
				}),
			}, {
				ASN:     cloudflare.F(int64(5567)),
				Comment: cloudflare.F("Private IP address"),
				Hostname: cloudflare.F(cloudflare.AccountRuleListItemListsNewListItemsParamsBodyHostname{
					URLHostname: cloudflare.F("example.com"),
				}),
				IP: cloudflare.F("10.0.0.1"),
				Redirect: cloudflare.F(cloudflare.AccountRuleListItemListsNewListItemsParamsBodyRedirect{
					IncludeSubdomains:   cloudflare.F(true),
					PreservePathSuffix:  cloudflare.F(true),
					PreserveQueryString: cloudflare.F(true),
					SourceURL:           cloudflare.F("example.com/arch"),
					StatusCode:          cloudflare.F(cloudflare.AccountRuleListItemListsNewListItemsParamsBodyRedirectStatusCode301),
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

func TestAccountRuleListItemListsGetListItemsWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Rules.Lists.Items.ListsGetListItems(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"2c0fc9fa937b11eaa1b71c4d701ab86e",
		cloudflare.AccountRuleListItemListsGetListItemsParams{
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

func TestAccountRuleListItemListsUpdateAllListItems(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Rules.Lists.Items.ListsUpdateAllListItems(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"2c0fc9fa937b11eaa1b71c4d701ab86e",
		cloudflare.AccountRuleListItemListsUpdateAllListItemsParams{
			Body: cloudflare.F([]cloudflare.AccountRuleListItemListsUpdateAllListItemsParamsBody{{
				ASN:     cloudflare.F(int64(5567)),
				Comment: cloudflare.F("Private IP address"),
				Hostname: cloudflare.F(cloudflare.AccountRuleListItemListsUpdateAllListItemsParamsBodyHostname{
					URLHostname: cloudflare.F("example.com"),
				}),
				IP: cloudflare.F("10.0.0.1"),
				Redirect: cloudflare.F(cloudflare.AccountRuleListItemListsUpdateAllListItemsParamsBodyRedirect{
					IncludeSubdomains:   cloudflare.F(true),
					PreservePathSuffix:  cloudflare.F(true),
					PreserveQueryString: cloudflare.F(true),
					SourceURL:           cloudflare.F("example.com/arch"),
					StatusCode:          cloudflare.F(cloudflare.AccountRuleListItemListsUpdateAllListItemsParamsBodyRedirectStatusCode301),
					SubpathMatching:     cloudflare.F(true),
					TargetURL:           cloudflare.F("https://archlinux.org/"),
				}),
			}, {
				ASN:     cloudflare.F(int64(5567)),
				Comment: cloudflare.F("Private IP address"),
				Hostname: cloudflare.F(cloudflare.AccountRuleListItemListsUpdateAllListItemsParamsBodyHostname{
					URLHostname: cloudflare.F("example.com"),
				}),
				IP: cloudflare.F("10.0.0.1"),
				Redirect: cloudflare.F(cloudflare.AccountRuleListItemListsUpdateAllListItemsParamsBodyRedirect{
					IncludeSubdomains:   cloudflare.F(true),
					PreservePathSuffix:  cloudflare.F(true),
					PreserveQueryString: cloudflare.F(true),
					SourceURL:           cloudflare.F("example.com/arch"),
					StatusCode:          cloudflare.F(cloudflare.AccountRuleListItemListsUpdateAllListItemsParamsBodyRedirectStatusCode301),
					SubpathMatching:     cloudflare.F(true),
					TargetURL:           cloudflare.F("https://archlinux.org/"),
				}),
			}, {
				ASN:     cloudflare.F(int64(5567)),
				Comment: cloudflare.F("Private IP address"),
				Hostname: cloudflare.F(cloudflare.AccountRuleListItemListsUpdateAllListItemsParamsBodyHostname{
					URLHostname: cloudflare.F("example.com"),
				}),
				IP: cloudflare.F("10.0.0.1"),
				Redirect: cloudflare.F(cloudflare.AccountRuleListItemListsUpdateAllListItemsParamsBodyRedirect{
					IncludeSubdomains:   cloudflare.F(true),
					PreservePathSuffix:  cloudflare.F(true),
					PreserveQueryString: cloudflare.F(true),
					SourceURL:           cloudflare.F("example.com/arch"),
					StatusCode:          cloudflare.F(cloudflare.AccountRuleListItemListsUpdateAllListItemsParamsBodyRedirectStatusCode301),
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
