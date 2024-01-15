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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.User.Tokens.Get(context.TODO(), map[string]interface{}{})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserTokenUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.User.Tokens.Update(
		context.TODO(),
		map[string]interface{}{},
		cloudflare.UserTokenUpdateParams{
			Name: cloudflare.F("readonly token"),
			Policies: cloudflare.F([]cloudflare.UserTokenUpdateParamsPolicy{{
				Effect:           cloudflare.F(cloudflare.UserTokenUpdateParamsPoliciesEffectAllow),
				PermissionGroups: cloudflare.F([]cloudflare.UserTokenUpdateParamsPoliciesPermissionGroup{{}, {}}),
				Resources: cloudflare.F[any](map[string]interface{}{
					"com.cloudflare.api.account.zone.22b1de5f1c0e4b3ea97bb1e963b06a43": "*",
					"com.cloudflare.api.account.zone.eb78d65290b24279ba6f44721b3ea3c4": "*",
				}),
			}, {
				Effect:           cloudflare.F(cloudflare.UserTokenUpdateParamsPoliciesEffectAllow),
				PermissionGroups: cloudflare.F([]cloudflare.UserTokenUpdateParamsPoliciesPermissionGroup{{}, {}}),
				Resources: cloudflare.F[any](map[string]interface{}{
					"com.cloudflare.api.account.zone.22b1de5f1c0e4b3ea97bb1e963b06a43": "*",
					"com.cloudflare.api.account.zone.eb78d65290b24279ba6f44721b3ea3c4": "*",
				}),
			}, {
				Effect:           cloudflare.F(cloudflare.UserTokenUpdateParamsPoliciesEffectAllow),
				PermissionGroups: cloudflare.F([]cloudflare.UserTokenUpdateParamsPoliciesPermissionGroup{{}, {}}),
				Resources: cloudflare.F[any](map[string]interface{}{
					"com.cloudflare.api.account.zone.22b1de5f1c0e4b3ea97bb1e963b06a43": "*",
					"com.cloudflare.api.account.zone.eb78d65290b24279ba6f44721b3ea3c4": "*",
				}),
			}}),
			Status: cloudflare.F(cloudflare.UserTokenUpdateParamsStatusActive),
			Condition: cloudflare.F(cloudflare.UserTokenUpdateParamsCondition{
				RequestIP: cloudflare.F(cloudflare.UserTokenUpdateParamsConditionRequestIP{
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.User.Tokens.Delete(context.TODO(), map[string]interface{}{})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserTokenUserAPITokensNewTokenWithOptionalParams(t *testing.T) {
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
	_, err := client.User.Tokens.UserAPITokensNewToken(context.TODO(), cloudflare.UserTokenUserAPITokensNewTokenParams{
		Name: cloudflare.F("readonly token"),
		Policies: cloudflare.F([]cloudflare.UserTokenUserAPITokensNewTokenParamsPolicy{{
			Effect:           cloudflare.F(cloudflare.UserTokenUserAPITokensNewTokenParamsPoliciesEffectAllow),
			PermissionGroups: cloudflare.F([]cloudflare.UserTokenUserAPITokensNewTokenParamsPoliciesPermissionGroup{{}, {}}),
			Resources: cloudflare.F[any](map[string]interface{}{
				"com.cloudflare.api.account.zone.22b1de5f1c0e4b3ea97bb1e963b06a43": "*",
				"com.cloudflare.api.account.zone.eb78d65290b24279ba6f44721b3ea3c4": "*",
			}),
		}, {
			Effect:           cloudflare.F(cloudflare.UserTokenUserAPITokensNewTokenParamsPoliciesEffectAllow),
			PermissionGroups: cloudflare.F([]cloudflare.UserTokenUserAPITokensNewTokenParamsPoliciesPermissionGroup{{}, {}}),
			Resources: cloudflare.F[any](map[string]interface{}{
				"com.cloudflare.api.account.zone.22b1de5f1c0e4b3ea97bb1e963b06a43": "*",
				"com.cloudflare.api.account.zone.eb78d65290b24279ba6f44721b3ea3c4": "*",
			}),
		}, {
			Effect:           cloudflare.F(cloudflare.UserTokenUserAPITokensNewTokenParamsPoliciesEffectAllow),
			PermissionGroups: cloudflare.F([]cloudflare.UserTokenUserAPITokensNewTokenParamsPoliciesPermissionGroup{{}, {}}),
			Resources: cloudflare.F[any](map[string]interface{}{
				"com.cloudflare.api.account.zone.22b1de5f1c0e4b3ea97bb1e963b06a43": "*",
				"com.cloudflare.api.account.zone.eb78d65290b24279ba6f44721b3ea3c4": "*",
			}),
		}}),
		Condition: cloudflare.F(cloudflare.UserTokenUserAPITokensNewTokenParamsCondition{
			RequestIP: cloudflare.F(cloudflare.UserTokenUserAPITokensNewTokenParamsConditionRequestIP{
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

func TestUserTokenUserAPITokensListTokensWithOptionalParams(t *testing.T) {
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
	_, err := client.User.Tokens.UserAPITokensListTokens(context.TODO(), cloudflare.UserTokenUserAPITokensListTokensParams{
		Direction: cloudflare.F(cloudflare.UserTokenUserAPITokensListTokensParamsDirectionDesc),
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
