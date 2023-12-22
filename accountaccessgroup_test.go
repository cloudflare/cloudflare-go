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

func TestAccountAccessGroupGet(t *testing.T) {
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
	_, err := client.Accounts.Accesses.Groups.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		map[string]interface{}{},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountAccessGroupUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Accesses.Groups.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		map[string]interface{}{},
		cloudflare.AccountAccessGroupUpdateParams{
			Include: cloudflare.F([]cloudflare.AccountAccessGroupUpdateParamsInclude{cloudflare.AccountAccessGroupUpdateParamsIncludeEmailRule(cloudflare.AccountAccessGroupUpdateParamsIncludeEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupUpdateParamsIncludeEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccountAccessGroupUpdateParamsIncludeEmailRule(cloudflare.AccountAccessGroupUpdateParamsIncludeEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupUpdateParamsIncludeEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccountAccessGroupUpdateParamsIncludeEmailRule(cloudflare.AccountAccessGroupUpdateParamsIncludeEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupUpdateParamsIncludeEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			Name: cloudflare.F("Allow devs"),
			Exclude: cloudflare.F([]cloudflare.AccountAccessGroupUpdateParamsExclude{cloudflare.AccountAccessGroupUpdateParamsExcludeEmailRule(cloudflare.AccountAccessGroupUpdateParamsExcludeEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupUpdateParamsExcludeEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccountAccessGroupUpdateParamsExcludeEmailRule(cloudflare.AccountAccessGroupUpdateParamsExcludeEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupUpdateParamsExcludeEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccountAccessGroupUpdateParamsExcludeEmailRule(cloudflare.AccountAccessGroupUpdateParamsExcludeEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupUpdateParamsExcludeEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			Require: cloudflare.F([]cloudflare.AccountAccessGroupUpdateParamsRequire{cloudflare.AccountAccessGroupUpdateParamsRequireEmailRule(cloudflare.AccountAccessGroupUpdateParamsRequireEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupUpdateParamsRequireEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccountAccessGroupUpdateParamsRequireEmailRule(cloudflare.AccountAccessGroupUpdateParamsRequireEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupUpdateParamsRequireEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccountAccessGroupUpdateParamsRequireEmailRule(cloudflare.AccountAccessGroupUpdateParamsRequireEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupUpdateParamsRequireEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
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

func TestAccountAccessGroupDelete(t *testing.T) {
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
	_, err := client.Accounts.Accesses.Groups.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		map[string]interface{}{},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountAccessGroupAccessGroupsNewAnAccessGroupWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Accesses.Groups.AccessGroupsNewAnAccessGroup(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParams{
			Include: cloudflare.F([]cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude{cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeEmailRule(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeEmailRule(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeEmailRule(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			Name: cloudflare.F("Allow devs"),
			Exclude: cloudflare.F([]cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude{cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeEmailRule(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeEmailRule(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeEmailRule(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			Require: cloudflare.F([]cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire{cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireEmailRule(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireEmailRule(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireEmailRule(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
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

func TestAccountAccessGroupAccessGroupsListAccessGroups(t *testing.T) {
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
	_, err := client.Accounts.Accesses.Groups.AccessGroupsListAccessGroups(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
