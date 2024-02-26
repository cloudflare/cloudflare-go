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

func TestAccessGroupNewWithOptionalParams(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Access.Groups.New(context.TODO(), cloudflare.AccessGroupNewParams{
		AccountID: cloudflare.F("string"),
		ZoneID:    cloudflare.F("string"),
		Include: cloudflare.F([]cloudflare.AccessGroupNewParamsInclude{cloudflare.AccessGroupNewParamsIncludeAccessEmailRule(cloudflare.AccessGroupNewParamsIncludeAccessEmailRule{
			Email: cloudflare.F(cloudflare.AccessGroupNewParamsIncludeAccessEmailRuleEmail{
				Email: cloudflare.F("test@example.com"),
			}),
		}), cloudflare.AccessGroupNewParamsIncludeAccessEmailRule(cloudflare.AccessGroupNewParamsIncludeAccessEmailRule{
			Email: cloudflare.F(cloudflare.AccessGroupNewParamsIncludeAccessEmailRuleEmail{
				Email: cloudflare.F("test@example.com"),
			}),
		}), cloudflare.AccessGroupNewParamsIncludeAccessEmailRule(cloudflare.AccessGroupNewParamsIncludeAccessEmailRule{
			Email: cloudflare.F(cloudflare.AccessGroupNewParamsIncludeAccessEmailRuleEmail{
				Email: cloudflare.F("test@example.com"),
			}),
		})}),
		Name: cloudflare.F("Allow devs"),
		Exclude: cloudflare.F([]cloudflare.AccessGroupNewParamsExclude{cloudflare.AccessGroupNewParamsExcludeAccessEmailRule(cloudflare.AccessGroupNewParamsExcludeAccessEmailRule{
			Email: cloudflare.F(cloudflare.AccessGroupNewParamsExcludeAccessEmailRuleEmail{
				Email: cloudflare.F("test@example.com"),
			}),
		}), cloudflare.AccessGroupNewParamsExcludeAccessEmailRule(cloudflare.AccessGroupNewParamsExcludeAccessEmailRule{
			Email: cloudflare.F(cloudflare.AccessGroupNewParamsExcludeAccessEmailRuleEmail{
				Email: cloudflare.F("test@example.com"),
			}),
		}), cloudflare.AccessGroupNewParamsExcludeAccessEmailRule(cloudflare.AccessGroupNewParamsExcludeAccessEmailRule{
			Email: cloudflare.F(cloudflare.AccessGroupNewParamsExcludeAccessEmailRuleEmail{
				Email: cloudflare.F("test@example.com"),
			}),
		})}),
		IsDefault: cloudflare.F(true),
		Require: cloudflare.F([]cloudflare.AccessGroupNewParamsRequire{cloudflare.AccessGroupNewParamsRequireAccessEmailRule(cloudflare.AccessGroupNewParamsRequireAccessEmailRule{
			Email: cloudflare.F(cloudflare.AccessGroupNewParamsRequireAccessEmailRuleEmail{
				Email: cloudflare.F("test@example.com"),
			}),
		}), cloudflare.AccessGroupNewParamsRequireAccessEmailRule(cloudflare.AccessGroupNewParamsRequireAccessEmailRule{
			Email: cloudflare.F(cloudflare.AccessGroupNewParamsRequireAccessEmailRuleEmail{
				Email: cloudflare.F("test@example.com"),
			}),
		}), cloudflare.AccessGroupNewParamsRequireAccessEmailRule(cloudflare.AccessGroupNewParamsRequireAccessEmailRule{
			Email: cloudflare.F(cloudflare.AccessGroupNewParamsRequireAccessEmailRuleEmail{
				Email: cloudflare.F("test@example.com"),
			}),
		})}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessGroupUpdateWithOptionalParams(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Access.Groups.Update(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.AccessGroupUpdateParams{
			AccountID: cloudflare.F("string"),
			ZoneID:    cloudflare.F("string"),
			Include: cloudflare.F([]cloudflare.AccessGroupUpdateParamsInclude{cloudflare.AccessGroupUpdateParamsIncludeAccessEmailRule(cloudflare.AccessGroupUpdateParamsIncludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessGroupUpdateParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessGroupUpdateParamsIncludeAccessEmailRule(cloudflare.AccessGroupUpdateParamsIncludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessGroupUpdateParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessGroupUpdateParamsIncludeAccessEmailRule(cloudflare.AccessGroupUpdateParamsIncludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessGroupUpdateParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			Name: cloudflare.F("Allow devs"),
			Exclude: cloudflare.F([]cloudflare.AccessGroupUpdateParamsExclude{cloudflare.AccessGroupUpdateParamsExcludeAccessEmailRule(cloudflare.AccessGroupUpdateParamsExcludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessGroupUpdateParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessGroupUpdateParamsExcludeAccessEmailRule(cloudflare.AccessGroupUpdateParamsExcludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessGroupUpdateParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessGroupUpdateParamsExcludeAccessEmailRule(cloudflare.AccessGroupUpdateParamsExcludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessGroupUpdateParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			IsDefault: cloudflare.F(true),
			Require: cloudflare.F([]cloudflare.AccessGroupUpdateParamsRequire{cloudflare.AccessGroupUpdateParamsRequireAccessEmailRule(cloudflare.AccessGroupUpdateParamsRequireAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessGroupUpdateParamsRequireAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessGroupUpdateParamsRequireAccessEmailRule(cloudflare.AccessGroupUpdateParamsRequireAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessGroupUpdateParamsRequireAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessGroupUpdateParamsRequireAccessEmailRule(cloudflare.AccessGroupUpdateParamsRequireAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessGroupUpdateParamsRequireAccessEmailRuleEmail{
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

func TestAccessGroupList(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Access.Groups.List(context.TODO(), cloudflare.AccessGroupListParams{
		AccountID: cloudflare.F("string"),
		ZoneID:    cloudflare.F("string"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessGroupDelete(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Access.Groups.Delete(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.AccessGroupDeleteParams{
			AccountID: cloudflare.F("string"),
			ZoneID:    cloudflare.F("string"),
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

func TestAccessGroupGet(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Access.Groups.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.AccessGroupGetParams{
			AccountID: cloudflare.F("string"),
			ZoneID:    cloudflare.F("string"),
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
