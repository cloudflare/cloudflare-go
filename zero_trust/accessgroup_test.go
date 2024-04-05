// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
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
	)
	_, err := client.ZeroTrust.Access.Groups.New(context.TODO(), zero_trust.AccessGroupNewParams{
		Include: cloudflare.F([]zero_trust.IncludeItemUnionParam{zero_trust.EmailRuleParam{
			Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
				Email: cloudflare.F("test@example.com"),
			}),
		}, zero_trust.EmailRuleParam{
			Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
				Email: cloudflare.F("test@example.com"),
			}),
		}, zero_trust.EmailRuleParam{
			Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
				Email: cloudflare.F("test@example.com"),
			}),
		}}),
		Name:      cloudflare.F("Allow devs"),
		AccountID: cloudflare.F("string"),
		ZoneID:    cloudflare.F("string"),
		Exclude: cloudflare.F([]zero_trust.ExcludeItemUnionParam{zero_trust.EmailRuleParam{
			Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
				Email: cloudflare.F("test@example.com"),
			}),
		}, zero_trust.EmailRuleParam{
			Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
				Email: cloudflare.F("test@example.com"),
			}),
		}, zero_trust.EmailRuleParam{
			Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
				Email: cloudflare.F("test@example.com"),
			}),
		}}),
		IsDefault: cloudflare.F(true),
		Require: cloudflare.F([]zero_trust.RequireItemUnionParam{zero_trust.EmailRuleParam{
			Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
				Email: cloudflare.F("test@example.com"),
			}),
		}, zero_trust.EmailRuleParam{
			Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
				Email: cloudflare.F("test@example.com"),
			}),
		}, zero_trust.EmailRuleParam{
			Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
				Email: cloudflare.F("test@example.com"),
			}),
		}}),
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
	)
	_, err := client.ZeroTrust.Access.Groups.Update(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.AccessGroupUpdateParams{
			Include: cloudflare.F([]zero_trust.IncludeItemUnionParam{zero_trust.EmailRuleParam{
				Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
					Email: cloudflare.F("test@example.com"),
				}),
			}, zero_trust.EmailRuleParam{
				Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
					Email: cloudflare.F("test@example.com"),
				}),
			}, zero_trust.EmailRuleParam{
				Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
					Email: cloudflare.F("test@example.com"),
				}),
			}}),
			Name:      cloudflare.F("Allow devs"),
			AccountID: cloudflare.F("string"),
			ZoneID:    cloudflare.F("string"),
			Exclude: cloudflare.F([]zero_trust.ExcludeItemUnionParam{zero_trust.EmailRuleParam{
				Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
					Email: cloudflare.F("test@example.com"),
				}),
			}, zero_trust.EmailRuleParam{
				Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
					Email: cloudflare.F("test@example.com"),
				}),
			}, zero_trust.EmailRuleParam{
				Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
					Email: cloudflare.F("test@example.com"),
				}),
			}}),
			IsDefault: cloudflare.F(true),
			Require: cloudflare.F([]zero_trust.RequireItemUnionParam{zero_trust.EmailRuleParam{
				Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
					Email: cloudflare.F("test@example.com"),
				}),
			}, zero_trust.EmailRuleParam{
				Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
					Email: cloudflare.F("test@example.com"),
				}),
			}, zero_trust.EmailRuleParam{
				Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
					Email: cloudflare.F("test@example.com"),
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

func TestAccessGroupListWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Groups.List(context.TODO(), zero_trust.AccessGroupListParams{
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

func TestAccessGroupDeleteWithOptionalParams(t *testing.T) {
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
		zero_trust.AccessGroupDeleteParams{
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

func TestAccessGroupGetWithOptionalParams(t *testing.T) {
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
		zero_trust.AccessGroupGetParams{
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
