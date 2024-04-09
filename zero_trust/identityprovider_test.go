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

func TestIdentityProviderNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.IdentityProviders.New(context.TODO(), zero_trust.IdentityProviderNewParamsAzureAD{
		Config: cloudflare.F(zero_trust.IdentityProviderNewParamsAzureADConfig{
			ClientID:                 cloudflare.F("<your client id>"),
			ClientSecret:             cloudflare.F("<your client secret>"),
			Claims:                   cloudflare.F([]string{"email_verified", "preferred_username", "custom_claim_name"}),
			EmailClaimName:           cloudflare.F("custom_claim_name"),
			ConditionalAccessEnabled: cloudflare.F(true),
			DirectoryID:              cloudflare.F("<your azure directory uuid>"),
			Prompt:                   cloudflare.F(zero_trust.IdentityProviderNewParamsAzureADConfigPromptLogin),
			SupportGroups:            cloudflare.F(true),
		}),
		Name:      cloudflare.F("Widget Corps IDP"),
		Type:      cloudflare.F(zero_trust.IdentityProviderTypeOnetimepin),
		AccountID: cloudflare.F("string"),
		ZoneID:    cloudflare.F("string"),
		ScimConfig: cloudflare.F(zero_trust.ScimConfigParam{
			Enabled:                cloudflare.F(true),
			GroupMemberDeprovision: cloudflare.F(true),
			SeatDeprovision:        cloudflare.F(true),
			Secret:                 cloudflare.F("string"),
			UserDeprovision:        cloudflare.F(true),
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

func TestIdentityProviderUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.IdentityProviders.Update(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.IdentityProviderUpdateParamsAzureAD{
			Config: cloudflare.F(zero_trust.IdentityProviderUpdateParamsAzureADConfig{
				ClientID:                 cloudflare.F("<your client id>"),
				ClientSecret:             cloudflare.F("<your client secret>"),
				Claims:                   cloudflare.F([]string{"email_verified", "preferred_username", "custom_claim_name"}),
				EmailClaimName:           cloudflare.F("custom_claim_name"),
				ConditionalAccessEnabled: cloudflare.F(true),
				DirectoryID:              cloudflare.F("<your azure directory uuid>"),
				Prompt:                   cloudflare.F(zero_trust.IdentityProviderUpdateParamsAzureADConfigPromptLogin),
				SupportGroups:            cloudflare.F(true),
			}),
			Name:      cloudflare.F("Widget Corps IDP"),
			Type:      cloudflare.F(zero_trust.IdentityProviderTypeOnetimepin),
			AccountID: cloudflare.F("string"),
			ZoneID:    cloudflare.F("string"),
			ScimConfig: cloudflare.F(zero_trust.ScimConfigParam{
				Enabled:                cloudflare.F(true),
				GroupMemberDeprovision: cloudflare.F(true),
				SeatDeprovision:        cloudflare.F(true),
				Secret:                 cloudflare.F("string"),
				UserDeprovision:        cloudflare.F(true),
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

func TestIdentityProviderListWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.IdentityProviders.List(context.TODO(), zero_trust.IdentityProviderListParams{
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

func TestIdentityProviderDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.IdentityProviders.Delete(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.IdentityProviderDeleteParams{
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

func TestIdentityProviderGetWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.IdentityProviders.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.IdentityProviderGetParams{
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
