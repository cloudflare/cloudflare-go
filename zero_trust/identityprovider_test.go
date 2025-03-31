// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
)

func TestIdentityProviderNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.IdentityProviders.New(context.TODO(), zero_trust.IdentityProviderNewParams{
		IdentityProvider: zero_trust.AzureADParam{
			Config: cloudflare.F(zero_trust.AzureADConfigParam{
				Claims:                   cloudflare.F([]string{"email_verified", "preferred_username", "custom_claim_name"}),
				ClientID:                 cloudflare.F("<your client id>"),
				ClientSecret:             cloudflare.F("<your client secret>"),
				ConditionalAccessEnabled: cloudflare.F(true),
				DirectoryID:              cloudflare.F("<your azure directory uuid>"),
				EmailClaimName:           cloudflare.F("custom_claim_name"),
				Prompt:                   cloudflare.F(zero_trust.AzureADConfigPromptLogin),
				SupportGroups:            cloudflare.F(true),
			}),
			Name: cloudflare.F("Widget Corps IDP"),
			Type: cloudflare.F(zero_trust.IdentityProviderTypeOnetimepin),
			SCIMConfig: cloudflare.F(zero_trust.IdentityProviderSCIMConfigParam{
				Enabled:                cloudflare.F(true),
				IdentityUpdateBehavior: cloudflare.F(zero_trust.IdentityProviderSCIMConfigIdentityUpdateBehaviorAutomatic),
				SeatDeprovision:        cloudflare.F(true),
				UserDeprovision:        cloudflare.F(true),
			}),
		},
		AccountID: cloudflare.F("account_id"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIdentityProviderUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.IdentityProviders.Update(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.IdentityProviderUpdateParams{
			IdentityProvider: zero_trust.AzureADParam{
				Config: cloudflare.F(zero_trust.AzureADConfigParam{
					Claims:                   cloudflare.F([]string{"email_verified", "preferred_username", "custom_claim_name"}),
					ClientID:                 cloudflare.F("<your client id>"),
					ClientSecret:             cloudflare.F("<your client secret>"),
					ConditionalAccessEnabled: cloudflare.F(true),
					DirectoryID:              cloudflare.F("<your azure directory uuid>"),
					EmailClaimName:           cloudflare.F("custom_claim_name"),
					Prompt:                   cloudflare.F(zero_trust.AzureADConfigPromptLogin),
					SupportGroups:            cloudflare.F(true),
				}),
				Name: cloudflare.F("Widget Corps IDP"),
				Type: cloudflare.F(zero_trust.IdentityProviderTypeOnetimepin),
				SCIMConfig: cloudflare.F(zero_trust.IdentityProviderSCIMConfigParam{
					Enabled:                cloudflare.F(true),
					IdentityUpdateBehavior: cloudflare.F(zero_trust.IdentityProviderSCIMConfigIdentityUpdateBehaviorAutomatic),
					SeatDeprovision:        cloudflare.F(true),
					UserDeprovision:        cloudflare.F(true),
				}),
			},
			AccountID: cloudflare.F("account_id"),
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

func TestIdentityProviderListWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.IdentityProviders.List(context.TODO(), zero_trust.IdentityProviderListParams{
		AccountID:   cloudflare.F("account_id"),
		SCIMEnabled: cloudflare.F("scim_enabled"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIdentityProviderDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.IdentityProviders.Delete(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.IdentityProviderDeleteParams{
			AccountID: cloudflare.F("account_id"),
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

func TestIdentityProviderGetWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.IdentityProviders.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.IdentityProviderGetParams{
			AccountID: cloudflare.F("account_id"),
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
