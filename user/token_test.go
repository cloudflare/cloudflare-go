// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/user"
)

func TestTokenNewWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.User.Tokens.New(context.TODO(), user.TokenNewParams{
		Name: cloudflare.F("readonly token"),
		Policies: cloudflare.F([]user.PolicyParam{{
			Effect: cloudflare.F(user.PolicyEffectAllow),
			PermissionGroups: cloudflare.F([]user.PolicyPermissionGroupParam{{
				Meta: cloudflare.F[any](map[string]interface{}{
					"label":  "load_balancer_admin",
					"scopes": "com.cloudflare.api.account",
				}),
			}, {
				Meta: cloudflare.F[any](map[string]interface{}{
					"label":  "fbm_user",
					"scopes": "com.cloudflare.api.account",
				}),
			}}),
			Resources: cloudflare.F[any](map[string]interface{}{
				"com.cloudflare.api.account.zone.22b1de5f1c0e4b3ea97bb1e963b06a43": "*",
				"com.cloudflare.api.account.zone.eb78d65290b24279ba6f44721b3ea3c4": "*",
			}),
		}, {
			Effect: cloudflare.F(user.PolicyEffectAllow),
			PermissionGroups: cloudflare.F([]user.PolicyPermissionGroupParam{{
				Meta: cloudflare.F[any](map[string]interface{}{
					"label":  "load_balancer_admin",
					"scopes": "com.cloudflare.api.account",
				}),
			}, {
				Meta: cloudflare.F[any](map[string]interface{}{
					"label":  "fbm_user",
					"scopes": "com.cloudflare.api.account",
				}),
			}}),
			Resources: cloudflare.F[any](map[string]interface{}{
				"com.cloudflare.api.account.zone.22b1de5f1c0e4b3ea97bb1e963b06a43": "*",
				"com.cloudflare.api.account.zone.eb78d65290b24279ba6f44721b3ea3c4": "*",
			}),
		}, {
			Effect: cloudflare.F(user.PolicyEffectAllow),
			PermissionGroups: cloudflare.F([]user.PolicyPermissionGroupParam{{
				Meta: cloudflare.F[any](map[string]interface{}{
					"label":  "load_balancer_admin",
					"scopes": "com.cloudflare.api.account",
				}),
			}, {
				Meta: cloudflare.F[any](map[string]interface{}{
					"label":  "fbm_user",
					"scopes": "com.cloudflare.api.account",
				}),
			}}),
			Resources: cloudflare.F[any](map[string]interface{}{
				"com.cloudflare.api.account.zone.22b1de5f1c0e4b3ea97bb1e963b06a43": "*",
				"com.cloudflare.api.account.zone.eb78d65290b24279ba6f44721b3ea3c4": "*",
			}),
		}}),
		Condition: cloudflare.F(user.TokenNewParamsCondition{
			RequestIP: cloudflare.F(user.TokenNewParamsConditionRequestIP{
				In:    cloudflare.F([]user.CIDRListParam{"123.123.123.0/24", "2606:4700::/32"}),
				NotIn: cloudflare.F([]user.CIDRListParam{"123.123.123.100/24", "2606:4700:4700::/48"}),
			}),
		}),
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

func TestTokenUpdateWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.User.Tokens.Update(
		context.TODO(),
		map[string]interface{}{},
		user.TokenUpdateParams{
			Token: user.TokenParam{
				Condition: cloudflare.F(user.TokenConditionParam{
					RequestIP: cloudflare.F(user.TokenConditionRequestIPParam{
						In:    cloudflare.F([]user.CIDRListParam{"123.123.123.0/24", "2606:4700::/32"}),
						NotIn: cloudflare.F([]user.CIDRListParam{"123.123.123.100/24", "2606:4700:4700::/48"}),
					}),
				}),
				Name:      cloudflare.F("readonly token"),
				NotBefore: cloudflare.F(time.Now()),
				Policies: cloudflare.F([]user.PolicyParam{{
					Effect: cloudflare.F(user.PolicyEffectAllow),
					PermissionGroups: cloudflare.F([]user.PolicyPermissionGroupParam{{
						Meta: cloudflare.F[any](map[string]interface{}{
							"label":  "load_balancer_admin",
							"scopes": "com.cloudflare.api.account",
						}),
					}, {
						Meta: cloudflare.F[any](map[string]interface{}{
							"label":  "fbm_user",
							"scopes": "com.cloudflare.api.account",
						}),
					}}),
					Resources: cloudflare.F[any](map[string]interface{}{
						"com.cloudflare.api.account.zone.22b1de5f1c0e4b3ea97bb1e963b06a43": "*",
						"com.cloudflare.api.account.zone.eb78d65290b24279ba6f44721b3ea3c4": "*",
					}),
				}, {
					Effect: cloudflare.F(user.PolicyEffectAllow),
					PermissionGroups: cloudflare.F([]user.PolicyPermissionGroupParam{{
						Meta: cloudflare.F[any](map[string]interface{}{
							"label":  "load_balancer_admin",
							"scopes": "com.cloudflare.api.account",
						}),
					}, {
						Meta: cloudflare.F[any](map[string]interface{}{
							"label":  "fbm_user",
							"scopes": "com.cloudflare.api.account",
						}),
					}}),
					Resources: cloudflare.F[any](map[string]interface{}{
						"com.cloudflare.api.account.zone.22b1de5f1c0e4b3ea97bb1e963b06a43": "*",
						"com.cloudflare.api.account.zone.eb78d65290b24279ba6f44721b3ea3c4": "*",
					}),
				}, {
					Effect: cloudflare.F(user.PolicyEffectAllow),
					PermissionGroups: cloudflare.F([]user.PolicyPermissionGroupParam{{
						Meta: cloudflare.F[any](map[string]interface{}{
							"label":  "load_balancer_admin",
							"scopes": "com.cloudflare.api.account",
						}),
					}, {
						Meta: cloudflare.F[any](map[string]interface{}{
							"label":  "fbm_user",
							"scopes": "com.cloudflare.api.account",
						}),
					}}),
					Resources: cloudflare.F[any](map[string]interface{}{
						"com.cloudflare.api.account.zone.22b1de5f1c0e4b3ea97bb1e963b06a43": "*",
						"com.cloudflare.api.account.zone.eb78d65290b24279ba6f44721b3ea3c4": "*",
					}),
				}}),
				Status: cloudflare.F(user.TokenStatusActive),
			},
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

func TestTokenListWithOptionalParams(t *testing.T) {
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
	_, err := client.User.Tokens.List(context.TODO(), user.TokenListParams{
		Direction: cloudflare.F(user.TokenListParamsDirectionDesc),
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

func TestTokenDelete(t *testing.T) {
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
	_, err := client.User.Tokens.Delete(context.TODO(), map[string]interface{}{})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTokenGet(t *testing.T) {
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
	_, err := client.User.Tokens.Get(context.TODO(), map[string]interface{}{})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTokenVerify(t *testing.T) {
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
	_, err := client.User.Tokens.Verify(context.TODO())
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
