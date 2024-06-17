// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

func TestIAMResourceGroupNewWithOptionalParams(t *testing.T) {
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
	_, err := client.IAM.ResourceGroups.New(context.TODO(), cloudflare.IAMResourceGroupNewParams{
		AccountID: cloudflare.F("eb78d65290b24279ba6f44721b3ea3c4"),
		Scope: cloudflare.F(cloudflare.IAMResourceGroupNewParamsScope{
			Key: cloudflare.F("com.cloudflare.api.account.eb78d65290b24279ba6f44721b3ea3c4"),
			Objects: cloudflare.F([]cloudflare.IAMResourceGroupNewParamsScopeObject{{
				Key: cloudflare.F("com.cloudflare.api.account.zone.23f8d65290b24279ba6f44721b3eaad5"),
			}, {
				Key: cloudflare.F("com.cloudflare.api.account.zone.23f8d65290b24279ba6f44721b3eaad5"),
			}, {
				Key: cloudflare.F("com.cloudflare.api.account.zone.23f8d65290b24279ba6f44721b3eaad5"),
			}}),
		}),
		Meta: cloudflare.F[any](map[string]interface{}{
			"editable": "false",
		}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIAMResourceGroupUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.IAM.ResourceGroups.Update(
		context.TODO(),
		"6d7f2f5f5b1d4a0e9081fdc98d432fd1",
		cloudflare.IAMResourceGroupUpdateParams{
			AccountID: cloudflare.F("eb78d65290b24279ba6f44721b3ea3c4"),
			Scope: cloudflare.F(cloudflare.IAMResourceGroupUpdateParamsScope{
				Key: cloudflare.F("com.cloudflare.api.account.eb78d65290b24279ba6f44721b3ea3c4"),
				Objects: cloudflare.F([]cloudflare.IAMResourceGroupUpdateParamsScopeObject{{
					Key: cloudflare.F("com.cloudflare.api.account.zone.23f8d65290b24279ba6f44721b3eaad5"),
				}, {
					Key: cloudflare.F("com.cloudflare.api.account.zone.23f8d65290b24279ba6f44721b3eaad5"),
				}, {
					Key: cloudflare.F("com.cloudflare.api.account.zone.23f8d65290b24279ba6f44721b3eaad5"),
				}}),
			}),
			Meta: cloudflare.F[any](map[string]interface{}{
				"editable": "false",
			}),
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

func TestIAMResourceGroupListWithOptionalParams(t *testing.T) {
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
	_, err := client.IAM.ResourceGroups.List(context.TODO(), cloudflare.IAMResourceGroupListParams{
		AccountID: cloudflare.F("eb78d65290b24279ba6f44721b3ea3c4"),
		ID:        cloudflare.F("6d7f2f5f5b1d4a0e9081fdc98d432fd1"),
		Name:      cloudflare.F("NameOfTheResourceGroup"),
		Page:      cloudflare.F(1.000000),
		PerPage:   cloudflare.F(5.000000),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIAMResourceGroupDelete(t *testing.T) {
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
	_, err := client.IAM.ResourceGroups.Delete(
		context.TODO(),
		"6d7f2f5f5b1d4a0e9081fdc98d432fd1",
		cloudflare.IAMResourceGroupDeleteParams{
			AccountID: cloudflare.F("eb78d65290b24279ba6f44721b3ea3c4"),
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

func TestIAMResourceGroupGet(t *testing.T) {
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
	_, err := client.IAM.ResourceGroups.Get(
		context.TODO(),
		"6d7f2f5f5b1d4a0e9081fdc98d432fd1",
		cloudflare.IAMResourceGroupGetParams{
			AccountID: cloudflare.F("eb78d65290b24279ba6f44721b3ea3c4"),
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
