// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/accounts"
	"github.com/cloudflare/cloudflare-go/v3/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/shared"
)

func TestTokenNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Tokens.New(context.TODO(), accounts.TokenNewParams{
		AccountID: cloudflare.F("eb78d65290b24279ba6f44721b3ea3c4"),
		Name:      cloudflare.F("readonly token"),
		Policies: cloudflare.F([]shared.TokenPolicyParam{{
			Effect: cloudflare.F(shared.TokenPolicyEffectAllow),
			PermissionGroups: cloudflare.F([]shared.TokenPolicyPermissionGroupParam{{
				Meta: cloudflare.F(shared.TokenPolicyPermissionGroupsMetaParam{
					Key:   cloudflare.F("key"),
					Value: cloudflare.F("value"),
				}),
			}, {
				Meta: cloudflare.F(shared.TokenPolicyPermissionGroupsMetaParam{
					Key:   cloudflare.F("key"),
					Value: cloudflare.F("value"),
				}),
			}}),
			Resources: cloudflare.F(shared.TokenPolicyResourcesParam{
				Resource: cloudflare.F("resource"),
				Scope:    cloudflare.F("scope"),
			}),
		}, {
			Effect: cloudflare.F(shared.TokenPolicyEffectAllow),
			PermissionGroups: cloudflare.F([]shared.TokenPolicyPermissionGroupParam{{
				Meta: cloudflare.F(shared.TokenPolicyPermissionGroupsMetaParam{
					Key:   cloudflare.F("key"),
					Value: cloudflare.F("value"),
				}),
			}, {
				Meta: cloudflare.F(shared.TokenPolicyPermissionGroupsMetaParam{
					Key:   cloudflare.F("key"),
					Value: cloudflare.F("value"),
				}),
			}}),
			Resources: cloudflare.F(shared.TokenPolicyResourcesParam{
				Resource: cloudflare.F("resource"),
				Scope:    cloudflare.F("scope"),
			}),
		}, {
			Effect: cloudflare.F(shared.TokenPolicyEffectAllow),
			PermissionGroups: cloudflare.F([]shared.TokenPolicyPermissionGroupParam{{
				Meta: cloudflare.F(shared.TokenPolicyPermissionGroupsMetaParam{
					Key:   cloudflare.F("key"),
					Value: cloudflare.F("value"),
				}),
			}, {
				Meta: cloudflare.F(shared.TokenPolicyPermissionGroupsMetaParam{
					Key:   cloudflare.F("key"),
					Value: cloudflare.F("value"),
				}),
			}}),
			Resources: cloudflare.F(shared.TokenPolicyResourcesParam{
				Resource: cloudflare.F("resource"),
				Scope:    cloudflare.F("scope"),
			}),
		}}),
		Condition: cloudflare.F(accounts.TokenNewParamsCondition{
			RequestIP: cloudflare.F(accounts.TokenNewParamsConditionRequestIP{
				In:    cloudflare.F([]shared.TokenConditionCIDRListParam{"123.123.123.0/24", "2606:4700::/32"}),
				NotIn: cloudflare.F([]shared.TokenConditionCIDRListParam{"123.123.123.100/24", "2606:4700:4700::/48"}),
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

func TestTokenUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Tokens.Update(
		context.TODO(),
		"ed17574386854bf78a67040be0a770b0",
		accounts.TokenUpdateParams{
			AccountID: cloudflare.F("eb78d65290b24279ba6f44721b3ea3c4"),
			Token: shared.TokenParam{
				Name: cloudflare.F("readonly token"),
				Policies: cloudflare.F([]shared.TokenPolicyParam{{
					Effect: cloudflare.F(shared.TokenPolicyEffectAllow),
					PermissionGroups: cloudflare.F([]shared.TokenPolicyPermissionGroupParam{{
						Meta: cloudflare.F(shared.TokenPolicyPermissionGroupsMetaParam{
							Key:   cloudflare.F("key"),
							Value: cloudflare.F("value"),
						}),
					}, {
						Meta: cloudflare.F(shared.TokenPolicyPermissionGroupsMetaParam{
							Key:   cloudflare.F("key"),
							Value: cloudflare.F("value"),
						}),
					}}),
					Resources: cloudflare.F(shared.TokenPolicyResourcesParam{
						Resource: cloudflare.F("resource"),
						Scope:    cloudflare.F("scope"),
					}),
				}, {
					Effect: cloudflare.F(shared.TokenPolicyEffectAllow),
					PermissionGroups: cloudflare.F([]shared.TokenPolicyPermissionGroupParam{{
						Meta: cloudflare.F(shared.TokenPolicyPermissionGroupsMetaParam{
							Key:   cloudflare.F("key"),
							Value: cloudflare.F("value"),
						}),
					}, {
						Meta: cloudflare.F(shared.TokenPolicyPermissionGroupsMetaParam{
							Key:   cloudflare.F("key"),
							Value: cloudflare.F("value"),
						}),
					}}),
					Resources: cloudflare.F(shared.TokenPolicyResourcesParam{
						Resource: cloudflare.F("resource"),
						Scope:    cloudflare.F("scope"),
					}),
				}, {
					Effect: cloudflare.F(shared.TokenPolicyEffectAllow),
					PermissionGroups: cloudflare.F([]shared.TokenPolicyPermissionGroupParam{{
						Meta: cloudflare.F(shared.TokenPolicyPermissionGroupsMetaParam{
							Key:   cloudflare.F("key"),
							Value: cloudflare.F("value"),
						}),
					}, {
						Meta: cloudflare.F(shared.TokenPolicyPermissionGroupsMetaParam{
							Key:   cloudflare.F("key"),
							Value: cloudflare.F("value"),
						}),
					}}),
					Resources: cloudflare.F(shared.TokenPolicyResourcesParam{
						Resource: cloudflare.F("resource"),
						Scope:    cloudflare.F("scope"),
					}),
				}}),
				Status: cloudflare.F(shared.TokenStatusActive),
				Condition: cloudflare.F(shared.TokenConditionParam{
					RequestIP: cloudflare.F(shared.TokenConditionRequestIPParam{
						In:    cloudflare.F([]shared.TokenConditionCIDRListParam{"123.123.123.0/24", "2606:4700::/32"}),
						NotIn: cloudflare.F([]shared.TokenConditionCIDRListParam{"123.123.123.100/24", "2606:4700:4700::/48"}),
					}),
				}),
				ExpiresOn: cloudflare.F(time.Now()),
				NotBefore: cloudflare.F(time.Now()),
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
	_, err := client.Accounts.Tokens.List(context.TODO(), accounts.TokenListParams{
		AccountID: cloudflare.F("eb78d65290b24279ba6f44721b3ea3c4"),
		Direction: cloudflare.F(accounts.TokenListParamsDirectionAsc),
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
	_, err := client.Accounts.Tokens.Delete(
		context.TODO(),
		"ed17574386854bf78a67040be0a770b0",
		accounts.TokenDeleteParams{
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
	_, err := client.Accounts.Tokens.Get(
		context.TODO(),
		"ed17574386854bf78a67040be0a770b0",
		accounts.TokenGetParams{
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
	_, err := client.Accounts.Tokens.Verify(context.TODO(), accounts.TokenVerifyParams{
		AccountID: cloudflare.F("eb78d65290b24279ba6f44721b3ea3c4"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
