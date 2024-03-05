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

func TestZeroTrustAccessGroupNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Groups.New(context.TODO(), cloudflare.ZeroTrustAccessGroupNewParams{
		Include: cloudflare.F([]cloudflare.ZeroTrustAccessGroupNewParamsInclude{cloudflare.ZeroTrustAccessGroupNewParamsIncludeAccessEmailRule(cloudflare.ZeroTrustAccessGroupNewParamsIncludeAccessEmailRule{
			Email: cloudflare.F(cloudflare.ZeroTrustAccessGroupNewParamsIncludeAccessEmailRuleEmail{
				Email: cloudflare.F("test@example.com"),
			}),
		}), cloudflare.ZeroTrustAccessGroupNewParamsIncludeAccessEmailRule(cloudflare.ZeroTrustAccessGroupNewParamsIncludeAccessEmailRule{
			Email: cloudflare.F(cloudflare.ZeroTrustAccessGroupNewParamsIncludeAccessEmailRuleEmail{
				Email: cloudflare.F("test@example.com"),
			}),
		}), cloudflare.ZeroTrustAccessGroupNewParamsIncludeAccessEmailRule(cloudflare.ZeroTrustAccessGroupNewParamsIncludeAccessEmailRule{
			Email: cloudflare.F(cloudflare.ZeroTrustAccessGroupNewParamsIncludeAccessEmailRuleEmail{
				Email: cloudflare.F("test@example.com"),
			}),
		})}),
		Name:      cloudflare.F("Allow devs"),
		AccountID: cloudflare.F("string"),
		ZoneID:    cloudflare.F("string"),
		Exclude: cloudflare.F([]cloudflare.ZeroTrustAccessGroupNewParamsExclude{cloudflare.ZeroTrustAccessGroupNewParamsExcludeAccessEmailRule(cloudflare.ZeroTrustAccessGroupNewParamsExcludeAccessEmailRule{
			Email: cloudflare.F(cloudflare.ZeroTrustAccessGroupNewParamsExcludeAccessEmailRuleEmail{
				Email: cloudflare.F("test@example.com"),
			}),
		}), cloudflare.ZeroTrustAccessGroupNewParamsExcludeAccessEmailRule(cloudflare.ZeroTrustAccessGroupNewParamsExcludeAccessEmailRule{
			Email: cloudflare.F(cloudflare.ZeroTrustAccessGroupNewParamsExcludeAccessEmailRuleEmail{
				Email: cloudflare.F("test@example.com"),
			}),
		}), cloudflare.ZeroTrustAccessGroupNewParamsExcludeAccessEmailRule(cloudflare.ZeroTrustAccessGroupNewParamsExcludeAccessEmailRule{
			Email: cloudflare.F(cloudflare.ZeroTrustAccessGroupNewParamsExcludeAccessEmailRuleEmail{
				Email: cloudflare.F("test@example.com"),
			}),
		})}),
		IsDefault: cloudflare.F(true),
		Require: cloudflare.F([]cloudflare.ZeroTrustAccessGroupNewParamsRequire{cloudflare.ZeroTrustAccessGroupNewParamsRequireAccessEmailRule(cloudflare.ZeroTrustAccessGroupNewParamsRequireAccessEmailRule{
			Email: cloudflare.F(cloudflare.ZeroTrustAccessGroupNewParamsRequireAccessEmailRuleEmail{
				Email: cloudflare.F("test@example.com"),
			}),
		}), cloudflare.ZeroTrustAccessGroupNewParamsRequireAccessEmailRule(cloudflare.ZeroTrustAccessGroupNewParamsRequireAccessEmailRule{
			Email: cloudflare.F(cloudflare.ZeroTrustAccessGroupNewParamsRequireAccessEmailRuleEmail{
				Email: cloudflare.F("test@example.com"),
			}),
		}), cloudflare.ZeroTrustAccessGroupNewParamsRequireAccessEmailRule(cloudflare.ZeroTrustAccessGroupNewParamsRequireAccessEmailRule{
			Email: cloudflare.F(cloudflare.ZeroTrustAccessGroupNewParamsRequireAccessEmailRuleEmail{
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

func TestZeroTrustAccessGroupUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Groups.Update(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.ZeroTrustAccessGroupUpdateParams{
			Include: cloudflare.F([]cloudflare.ZeroTrustAccessGroupUpdateParamsInclude{cloudflare.ZeroTrustAccessGroupUpdateParamsIncludeAccessEmailRule(cloudflare.ZeroTrustAccessGroupUpdateParamsIncludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.ZeroTrustAccessGroupUpdateParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZeroTrustAccessGroupUpdateParamsIncludeAccessEmailRule(cloudflare.ZeroTrustAccessGroupUpdateParamsIncludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.ZeroTrustAccessGroupUpdateParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZeroTrustAccessGroupUpdateParamsIncludeAccessEmailRule(cloudflare.ZeroTrustAccessGroupUpdateParamsIncludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.ZeroTrustAccessGroupUpdateParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			Name:      cloudflare.F("Allow devs"),
			AccountID: cloudflare.F("string"),
			ZoneID:    cloudflare.F("string"),
			Exclude: cloudflare.F([]cloudflare.ZeroTrustAccessGroupUpdateParamsExclude{cloudflare.ZeroTrustAccessGroupUpdateParamsExcludeAccessEmailRule(cloudflare.ZeroTrustAccessGroupUpdateParamsExcludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.ZeroTrustAccessGroupUpdateParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZeroTrustAccessGroupUpdateParamsExcludeAccessEmailRule(cloudflare.ZeroTrustAccessGroupUpdateParamsExcludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.ZeroTrustAccessGroupUpdateParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZeroTrustAccessGroupUpdateParamsExcludeAccessEmailRule(cloudflare.ZeroTrustAccessGroupUpdateParamsExcludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.ZeroTrustAccessGroupUpdateParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			IsDefault: cloudflare.F(true),
			Require: cloudflare.F([]cloudflare.ZeroTrustAccessGroupUpdateParamsRequire{cloudflare.ZeroTrustAccessGroupUpdateParamsRequireAccessEmailRule(cloudflare.ZeroTrustAccessGroupUpdateParamsRequireAccessEmailRule{
				Email: cloudflare.F(cloudflare.ZeroTrustAccessGroupUpdateParamsRequireAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZeroTrustAccessGroupUpdateParamsRequireAccessEmailRule(cloudflare.ZeroTrustAccessGroupUpdateParamsRequireAccessEmailRule{
				Email: cloudflare.F(cloudflare.ZeroTrustAccessGroupUpdateParamsRequireAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZeroTrustAccessGroupUpdateParamsRequireAccessEmailRule(cloudflare.ZeroTrustAccessGroupUpdateParamsRequireAccessEmailRule{
				Email: cloudflare.F(cloudflare.ZeroTrustAccessGroupUpdateParamsRequireAccessEmailRuleEmail{
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

func TestZeroTrustAccessGroupListWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Groups.List(context.TODO(), cloudflare.ZeroTrustAccessGroupListParams{
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

func TestZeroTrustAccessGroupDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Groups.Delete(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.ZeroTrustAccessGroupDeleteParams{
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

func TestZeroTrustAccessGroupGetWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Groups.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.ZeroTrustAccessGroupGetParams{
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
