// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestUserTokenNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Tokens.New(context.TODO(), cloudflare.UserTokenNewParams{
		Name: cloudflare.F("readonly token"),
		Policies: cloudflare.F([]cloudflare.UserTokenNewParamsPolicy{{
			Effect:           cloudflare.F(cloudflare.UserTokenNewParamsPoliciesEffectAllow),
			PermissionGroups: cloudflare.F([]cloudflare.UserTokenNewParamsPoliciesPermissionGroup{{}, {}}),
			Resources: cloudflare.F[any](map[string]interface{}{
				"com.cloudflare.api.account.zone.22b1de5f1c0e4b3ea97bb1e963b06a43": "*",
				"com.cloudflare.api.account.zone.eb78d65290b24279ba6f44721b3ea3c4": "*",
			}),
		}, {
			Effect:           cloudflare.F(cloudflare.UserTokenNewParamsPoliciesEffectAllow),
			PermissionGroups: cloudflare.F([]cloudflare.UserTokenNewParamsPoliciesPermissionGroup{{}, {}}),
			Resources: cloudflare.F[any](map[string]interface{}{
				"com.cloudflare.api.account.zone.22b1de5f1c0e4b3ea97bb1e963b06a43": "*",
				"com.cloudflare.api.account.zone.eb78d65290b24279ba6f44721b3ea3c4": "*",
			}),
		}, {
			Effect:           cloudflare.F(cloudflare.UserTokenNewParamsPoliciesEffectAllow),
			PermissionGroups: cloudflare.F([]cloudflare.UserTokenNewParamsPoliciesPermissionGroup{{}, {}}),
			Resources: cloudflare.F[any](map[string]interface{}{
				"com.cloudflare.api.account.zone.22b1de5f1c0e4b3ea97bb1e963b06a43": "*",
				"com.cloudflare.api.account.zone.eb78d65290b24279ba6f44721b3ea3c4": "*",
			}),
		}}),
		Condition: cloudflare.F(cloudflare.UserTokenNewParamsCondition{
			RequestIP: cloudflare.F(cloudflare.UserTokenNewParamsConditionRequestIP{
				In:    cloudflare.F([]string{"123.123.123.0/24", "2606:4700::/32"}),
				NotIn: cloudflare.F([]string{"123.123.123.100/24", "2606:4700:4700::/48"}),
			}),
		}),
		ExpiresOn: cloudflare.F(time.Now()),
		NotBefore: cloudflare.F(time.Now()),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserTokenListWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Tokens.List(context.TODO(), cloudflare.UserTokenListParams{
		Direction: cloudflare.F(cloudflare.UserTokenListParamsDirectionDesc),
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

func TestUserTokenDelete(t *testing.T) {
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
	_, err := client.Users.Tokens.Delete(context.TODO(), map[string]interface{}{})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserTokenGet(t *testing.T) {
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
	_, err := client.Users.Tokens.Get(context.TODO(), map[string]interface{}{})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserTokenReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Tokens.Replace(
		context.TODO(),
		map[string]interface{}{},
		cloudflare.UserTokenReplaceParams{
			Name: cloudflare.F("readonly token"),
			Policies: cloudflare.F([]cloudflare.UserTokenReplaceParamsPolicy{{
				Effect:           cloudflare.F(cloudflare.UserTokenReplaceParamsPoliciesEffectAllow),
				PermissionGroups: cloudflare.F([]cloudflare.UserTokenReplaceParamsPoliciesPermissionGroup{{}, {}}),
				Resources: cloudflare.F[any](map[string]interface{}{
					"com.cloudflare.api.account.zone.22b1de5f1c0e4b3ea97bb1e963b06a43": "*",
					"com.cloudflare.api.account.zone.eb78d65290b24279ba6f44721b3ea3c4": "*",
				}),
			}, {
				Effect:           cloudflare.F(cloudflare.UserTokenReplaceParamsPoliciesEffectAllow),
				PermissionGroups: cloudflare.F([]cloudflare.UserTokenReplaceParamsPoliciesPermissionGroup{{}, {}}),
				Resources: cloudflare.F[any](map[string]interface{}{
					"com.cloudflare.api.account.zone.22b1de5f1c0e4b3ea97bb1e963b06a43": "*",
					"com.cloudflare.api.account.zone.eb78d65290b24279ba6f44721b3ea3c4": "*",
				}),
			}, {
				Effect:           cloudflare.F(cloudflare.UserTokenReplaceParamsPoliciesEffectAllow),
				PermissionGroups: cloudflare.F([]cloudflare.UserTokenReplaceParamsPoliciesPermissionGroup{{}, {}}),
				Resources: cloudflare.F[any](map[string]interface{}{
					"com.cloudflare.api.account.zone.22b1de5f1c0e4b3ea97bb1e963b06a43": "*",
					"com.cloudflare.api.account.zone.eb78d65290b24279ba6f44721b3ea3c4": "*",
				}),
			}}),
			Status: cloudflare.F(cloudflare.UserTokenReplaceParamsStatusActive),
			Condition: cloudflare.F(cloudflare.UserTokenReplaceParamsCondition{
				RequestIP: cloudflare.F(cloudflare.UserTokenReplaceParamsConditionRequestIP{
					In:    cloudflare.F([]string{"123.123.123.0/24", "2606:4700::/32"}),
					NotIn: cloudflare.F([]string{"123.123.123.100/24", "2606:4700:4700::/48"}),
				}),
			}),
			ExpiresOn: cloudflare.F(time.Now()),
			NotBefore: cloudflare.F(time.Now()),
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

func TestUserTokenVerify(t *testing.T) {
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
	_, err := client.Users.Tokens.Verify(context.TODO())
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
