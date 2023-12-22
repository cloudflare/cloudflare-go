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
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Zones.Accesses.Groups.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZoneAccessGroupNewParams{
			Include: cloudflare.F([]cloudflare.ZoneAccessGroupNewParamsInclude{cloudflare.ZoneAccessGroupNewParamsIncludeEmailRule(cloudflare.ZoneAccessGroupNewParamsIncludeEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupNewParamsIncludeEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZoneAccessGroupNewParamsIncludeEmailRule(cloudflare.ZoneAccessGroupNewParamsIncludeEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupNewParamsIncludeEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZoneAccessGroupNewParamsIncludeEmailRule(cloudflare.ZoneAccessGroupNewParamsIncludeEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupNewParamsIncludeEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			Name: cloudflare.F("Allow devs"),
			Exclude: cloudflare.F([]cloudflare.ZoneAccessGroupNewParamsExclude{cloudflare.ZoneAccessGroupNewParamsExcludeEmailRule(cloudflare.ZoneAccessGroupNewParamsExcludeEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupNewParamsExcludeEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZoneAccessGroupNewParamsExcludeEmailRule(cloudflare.ZoneAccessGroupNewParamsExcludeEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupNewParamsExcludeEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZoneAccessGroupNewParamsExcludeEmailRule(cloudflare.ZoneAccessGroupNewParamsExcludeEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupNewParamsExcludeEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			Require: cloudflare.F([]cloudflare.ZoneAccessGroupNewParamsRequire{cloudflare.ZoneAccessGroupNewParamsRequireEmailRule(cloudflare.ZoneAccessGroupNewParamsRequireEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupNewParamsRequireEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZoneAccessGroupNewParamsRequireEmailRule(cloudflare.ZoneAccessGroupNewParamsRequireEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupNewParamsRequireEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZoneAccessGroupNewParamsRequireEmailRule(cloudflare.ZoneAccessGroupNewParamsRequireEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupNewParamsRequireEmailRuleEmail{
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
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Zones.Accesses.Groups.Get(
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
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Zones.Accesses.Groups.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		map[string]interface{}{},
		cloudflare.ZoneAccessGroupUpdateParams{
			Include: cloudflare.F([]cloudflare.ZoneAccessGroupUpdateParamsInclude{cloudflare.ZoneAccessGroupUpdateParamsIncludeEmailRule(cloudflare.ZoneAccessGroupUpdateParamsIncludeEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupUpdateParamsIncludeEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZoneAccessGroupUpdateParamsIncludeEmailRule(cloudflare.ZoneAccessGroupUpdateParamsIncludeEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupUpdateParamsIncludeEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZoneAccessGroupUpdateParamsIncludeEmailRule(cloudflare.ZoneAccessGroupUpdateParamsIncludeEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupUpdateParamsIncludeEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			Name: cloudflare.F("Allow devs"),
			Exclude: cloudflare.F([]cloudflare.ZoneAccessGroupUpdateParamsExclude{cloudflare.ZoneAccessGroupUpdateParamsExcludeEmailRule(cloudflare.ZoneAccessGroupUpdateParamsExcludeEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupUpdateParamsExcludeEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZoneAccessGroupUpdateParamsExcludeEmailRule(cloudflare.ZoneAccessGroupUpdateParamsExcludeEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupUpdateParamsExcludeEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZoneAccessGroupUpdateParamsExcludeEmailRule(cloudflare.ZoneAccessGroupUpdateParamsExcludeEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupUpdateParamsExcludeEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			Require: cloudflare.F([]cloudflare.ZoneAccessGroupUpdateParamsRequire{cloudflare.ZoneAccessGroupUpdateParamsRequireEmailRule(cloudflare.ZoneAccessGroupUpdateParamsRequireEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupUpdateParamsRequireEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZoneAccessGroupUpdateParamsRequireEmailRule(cloudflare.ZoneAccessGroupUpdateParamsRequireEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupUpdateParamsRequireEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZoneAccessGroupUpdateParamsRequireEmailRule(cloudflare.ZoneAccessGroupUpdateParamsRequireEmailRule{
				Email: cloudflare.F(cloudflare.ZoneAccessGroupUpdateParamsRequireEmailRuleEmail{
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
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
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
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Zones.Accesses.Groups.Delete(
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
