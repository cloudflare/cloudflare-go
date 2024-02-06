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

func TestAccessIdentityProviderGet(t *testing.T) {
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
	_, err := client.Access.IdentityProviders.Get(
		context.TODO(),
		"string",
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessIdentityProviderUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Access.IdentityProviders.Update(
		context.TODO(),
		"string",
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.AccessIdentityProviderUpdateParamsAccessAzureAd{
			Config: cloudflare.F(cloudflare.AccessIdentityProviderUpdateParamsAccessAzureAdConfig{
				ClientID:                 cloudflare.F("<your client id>"),
				ClientSecret:             cloudflare.F("<your client secret>"),
				Claims:                   cloudflare.F([]string{"email_verified", "preferred_username", "custom_claim_name"}),
				EmailClaimName:           cloudflare.F("custom_claim_name"),
				ConditionalAccessEnabled: cloudflare.F(true),
				DirectoryID:              cloudflare.F("<your azure directory uuid>"),
				SupportGroups:            cloudflare.F(true),
			}),
			Name: cloudflare.F("Widget Corps IDP"),
			ScimConfig: cloudflare.F(cloudflare.AccessIdentityProviderUpdateParamsAccessAzureAdScimConfig{
				Enabled:                cloudflare.F(true),
				GroupMemberDeprovision: cloudflare.F(true),
				SeatDeprovision:        cloudflare.F(true),
				Secret:                 cloudflare.F("string"),
				UserDeprovision:        cloudflare.F(true),
			}),
			Type: cloudflare.F(cloudflare.AccessIdentityProviderUpdateParamsAccessAzureAdTypeOnetimepin),
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

func TestAccessIdentityProviderDelete(t *testing.T) {
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
	_, err := client.Access.IdentityProviders.Delete(
		context.TODO(),
		"string",
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderWithOptionalParams(t *testing.T) {
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
	_, err := client.Access.IdentityProviders.AccessIdentityProvidersAddAnAccessIdentityProvider(
		context.TODO(),
		"string",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAd{
			Config: cloudflare.F(cloudflare.AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdConfig{
				ClientID:                 cloudflare.F("<your client id>"),
				ClientSecret:             cloudflare.F("<your client secret>"),
				Claims:                   cloudflare.F([]string{"email_verified", "preferred_username", "custom_claim_name"}),
				EmailClaimName:           cloudflare.F("custom_claim_name"),
				ConditionalAccessEnabled: cloudflare.F(true),
				DirectoryID:              cloudflare.F("<your azure directory uuid>"),
				SupportGroups:            cloudflare.F(true),
			}),
			Name: cloudflare.F("Widget Corps IDP"),
			ScimConfig: cloudflare.F(cloudflare.AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdScimConfig{
				Enabled:                cloudflare.F(true),
				GroupMemberDeprovision: cloudflare.F(true),
				SeatDeprovision:        cloudflare.F(true),
				Secret:                 cloudflare.F("string"),
				UserDeprovision:        cloudflare.F(true),
			}),
			Type: cloudflare.F(cloudflare.AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdTypeOnetimepin),
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

func TestAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProviders(t *testing.T) {
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
	_, err := client.Access.IdentityProviders.AccessIdentityProvidersListAccessIdentityProviders(
		context.TODO(),
		"string",
		"023e105f4ecef8ad9ca31a8372d0c353",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
