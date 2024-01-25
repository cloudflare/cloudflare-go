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

func TestZoneAccessGroupNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Accesses.Groups.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZoneAccessGroupNewParams{
			Include: cloudflare.F([]cloudflare.ZoneAccessGroupNewParamsInclude{cloudflare.ZoneAccessGroupNewParamsIncludePajwohLqEmailRule(cloudflare.ZoneAccessGroupNewParamsIncludePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupNewParamsIncludePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZoneAccessGroupNewParamsIncludePajwohLqEmailRule(cloudflare.ZoneAccessGroupNewParamsIncludePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupNewParamsIncludePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZoneAccessGroupNewParamsIncludePajwohLqEmailRule(cloudflare.ZoneAccessGroupNewParamsIncludePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupNewParamsIncludePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			Name: cloudflare.F("Allow devs"),
			Exclude: cloudflare.F([]cloudflare.ZoneAccessGroupNewParamsExclude{cloudflare.ZoneAccessGroupNewParamsExcludePajwohLqEmailRule(cloudflare.ZoneAccessGroupNewParamsExcludePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupNewParamsExcludePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZoneAccessGroupNewParamsExcludePajwohLqEmailRule(cloudflare.ZoneAccessGroupNewParamsExcludePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupNewParamsExcludePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZoneAccessGroupNewParamsExcludePajwohLqEmailRule(cloudflare.ZoneAccessGroupNewParamsExcludePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupNewParamsExcludePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			Require: cloudflare.F([]cloudflare.ZoneAccessGroupNewParamsRequire{cloudflare.ZoneAccessGroupNewParamsRequirePajwohLqEmailRule(cloudflare.ZoneAccessGroupNewParamsRequirePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupNewParamsRequirePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZoneAccessGroupNewParamsRequirePajwohLqEmailRule(cloudflare.ZoneAccessGroupNewParamsRequirePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupNewParamsRequirePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZoneAccessGroupNewParamsRequirePajwohLqEmailRule(cloudflare.ZoneAccessGroupNewParamsRequirePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupNewParamsRequirePajwohLqEmailRuleEmail{
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

func TestZoneAccessGroupGet(t *testing.T) {
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
	_, err := client.Zones.Accesses.Groups.Get(
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

func TestZoneAccessGroupUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Accesses.Groups.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.ZoneAccessGroupUpdateParams{
			Include: cloudflare.F([]cloudflare.ZoneAccessGroupUpdateParamsInclude{cloudflare.ZoneAccessGroupUpdateParamsIncludePajwohLqEmailRule(cloudflare.ZoneAccessGroupUpdateParamsIncludePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupUpdateParamsIncludePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZoneAccessGroupUpdateParamsIncludePajwohLqEmailRule(cloudflare.ZoneAccessGroupUpdateParamsIncludePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupUpdateParamsIncludePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZoneAccessGroupUpdateParamsIncludePajwohLqEmailRule(cloudflare.ZoneAccessGroupUpdateParamsIncludePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupUpdateParamsIncludePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			Name: cloudflare.F("Allow devs"),
			Exclude: cloudflare.F([]cloudflare.ZoneAccessGroupUpdateParamsExclude{cloudflare.ZoneAccessGroupUpdateParamsExcludePajwohLqEmailRule(cloudflare.ZoneAccessGroupUpdateParamsExcludePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupUpdateParamsExcludePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZoneAccessGroupUpdateParamsExcludePajwohLqEmailRule(cloudflare.ZoneAccessGroupUpdateParamsExcludePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupUpdateParamsExcludePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZoneAccessGroupUpdateParamsExcludePajwohLqEmailRule(cloudflare.ZoneAccessGroupUpdateParamsExcludePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupUpdateParamsExcludePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			Require: cloudflare.F([]cloudflare.ZoneAccessGroupUpdateParamsRequire{cloudflare.ZoneAccessGroupUpdateParamsRequirePajwohLqEmailRule(cloudflare.ZoneAccessGroupUpdateParamsRequirePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupUpdateParamsRequirePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZoneAccessGroupUpdateParamsRequirePajwohLqEmailRule(cloudflare.ZoneAccessGroupUpdateParamsRequirePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupUpdateParamsRequirePajwohLqEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZoneAccessGroupUpdateParamsRequirePajwohLqEmailRule(cloudflare.ZoneAccessGroupUpdateParamsRequirePajwohLqEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupUpdateParamsRequirePajwohLqEmailRuleEmail{
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

func TestZoneAccessGroupList(t *testing.T) {
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
	_, err := client.Zones.Accesses.Groups.List(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneAccessGroupDelete(t *testing.T) {
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
	_, err := client.Zones.Accesses.Groups.Delete(
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
