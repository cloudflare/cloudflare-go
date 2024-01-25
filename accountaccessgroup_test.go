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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Access.Groups.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Access.Groups.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.AccountAccessGroupUpdateParams{
			Include: cloudflare.F([]cloudflare.AccountAccessGroupUpdateParamsInclude{cloudflare.AccountAccessGroupUpdateParamsIncludePajwohLqEmailRule(cloudflare.AccountAccessGroupUpdateParamsIncludePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupUpdateParamsIncludePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccountAccessGroupUpdateParamsIncludePajwohLqEmailRule(cloudflare.AccountAccessGroupUpdateParamsIncludePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupUpdateParamsIncludePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccountAccessGroupUpdateParamsIncludePajwohLqEmailRule(cloudflare.AccountAccessGroupUpdateParamsIncludePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupUpdateParamsIncludePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			Name: cloudflare.F("Allow devs"),
			Exclude: cloudflare.F([]cloudflare.AccountAccessGroupUpdateParamsExclude{cloudflare.AccountAccessGroupUpdateParamsExcludePajwohLqEmailRule(cloudflare.AccountAccessGroupUpdateParamsExcludePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupUpdateParamsExcludePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccountAccessGroupUpdateParamsExcludePajwohLqEmailRule(cloudflare.AccountAccessGroupUpdateParamsExcludePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupUpdateParamsExcludePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccountAccessGroupUpdateParamsExcludePajwohLqEmailRule(cloudflare.AccountAccessGroupUpdateParamsExcludePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupUpdateParamsExcludePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			IsDefault: cloudflare.F(true),
			Require: cloudflare.F([]cloudflare.AccountAccessGroupUpdateParamsRequire{cloudflare.AccountAccessGroupUpdateParamsRequirePajwohLqEmailRule(cloudflare.AccountAccessGroupUpdateParamsRequirePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupUpdateParamsRequirePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccountAccessGroupUpdateParamsRequirePajwohLqEmailRule(cloudflare.AccountAccessGroupUpdateParamsRequirePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupUpdateParamsRequirePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccountAccessGroupUpdateParamsRequirePajwohLqEmailRule(cloudflare.AccountAccessGroupUpdateParamsRequirePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupUpdateParamsRequirePajwohLqEmailRuleEmail{
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Access.Groups.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Access.Groups.AccessGroupsNewAnAccessGroup(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParams{
			Include: cloudflare.F([]cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude{cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqEmailRule(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqEmailRule(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqEmailRule(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			Name: cloudflare.F("Allow devs"),
			Exclude: cloudflare.F([]cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude{cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqEmailRule(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqEmailRule(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqEmailRule(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			IsDefault: cloudflare.F(true),
			Require: cloudflare.F([]cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire{cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqEmailRule(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqEmailRule(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqEmailRule(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqEmailRuleEmail{
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Access.Groups.AccessGroupsListAccessGroups(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
