// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/accounts"
	"github.com/cloudflare/cloudflare-go/v5/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
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
	_, err := client.Accounts.Tokens.New(context.TODO(), accounts.TokenNewParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Name:      cloudflare.F("readonly token"),
		Policies: cloudflare.F([]shared.TokenPolicyParam{{
			Effect: cloudflare.F(shared.TokenPolicyEffectAllow),
			PermissionGroups: cloudflare.F([]shared.TokenPolicyPermissionGroupParam{{
				ID: cloudflare.F("c8fed203ed3043cba015a93ad1616f1f"),
				Meta: cloudflare.F(shared.TokenPolicyPermissionGroupsMetaParam{
					Key:   cloudflare.F("key"),
					Value: cloudflare.F("value"),
				}),
			}, {
				ID: cloudflare.F("82e64a83756745bbbb1c9c2701bf816b"),
				Meta: cloudflare.F(shared.TokenPolicyPermissionGroupsMetaParam{
					Key:   cloudflare.F("key"),
					Value: cloudflare.F("value"),
				}),
			}}),
			Resources: cloudflare.F[shared.TokenPolicyResourcesUnionParam](shared.TokenPolicyResourcesIAMResourcesTypeObjectStringParam(map[string]string{
				"foo": "string",
			})),
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
	_, err := client.Accounts.Tokens.Update(
		context.TODO(),
		"ed17574386854bf78a67040be0a770b0",
		accounts.TokenUpdateParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Token: shared.TokenParam{
				Condition: cloudflare.F(shared.TokenConditionParam{
					RequestIP: cloudflare.F(shared.TokenConditionRequestIPParam{
						In:    cloudflare.F([]shared.TokenConditionCIDRListParam{"123.123.123.0/24", "2606:4700::/32"}),
						NotIn: cloudflare.F([]shared.TokenConditionCIDRListParam{"123.123.123.100/24", "2606:4700:4700::/48"}),
					}),
				}),
				ExpiresOn: cloudflare.F(time.Now()),
				Name:      cloudflare.F("readonly token"),
				NotBefore: cloudflare.F(time.Now()),
				Policies: cloudflare.F([]shared.TokenPolicyParam{{
					Effect: cloudflare.F(shared.TokenPolicyEffectAllow),
					PermissionGroups: cloudflare.F([]shared.TokenPolicyPermissionGroupParam{{
						ID: cloudflare.F("c8fed203ed3043cba015a93ad1616f1f"),
						Meta: cloudflare.F(shared.TokenPolicyPermissionGroupsMetaParam{
							Key:   cloudflare.F("key"),
							Value: cloudflare.F("value"),
						}),
					}, {
						ID: cloudflare.F("82e64a83756745bbbb1c9c2701bf816b"),
						Meta: cloudflare.F(shared.TokenPolicyPermissionGroupsMetaParam{
							Key:   cloudflare.F("key"),
							Value: cloudflare.F("value"),
						}),
					}}),
					Resources: cloudflare.F[shared.TokenPolicyResourcesUnionParam](shared.TokenPolicyResourcesIAMResourcesTypeObjectStringParam(map[string]string{
						"foo": "string",
					})),
				}}),
				Status: cloudflare.F(shared.TokenStatusActive),
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
	_, err := client.Accounts.Tokens.List(context.TODO(), accounts.TokenListParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Direction: cloudflare.F(accounts.TokenListParamsDirectionDesc),
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
	_, err := client.Accounts.Tokens.Delete(
		context.TODO(),
		"ed17574386854bf78a67040be0a770b0",
		accounts.TokenDeleteParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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
	_, err := client.Accounts.Tokens.Get(
		context.TODO(),
		"ed17574386854bf78a67040be0a770b0",
		accounts.TokenGetParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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
	_, err := client.Accounts.Tokens.Verify(context.TODO(), accounts.TokenVerifyParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
