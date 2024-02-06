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

func TestAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationWithOptionalParams(t *testing.T) {
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
	_, err := client.Access.Organizations.ZeroTrustOrganizationNewYourZeroTrustOrganization(
		context.TODO(),
		"string",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationParams{
			AuthDomain:               cloudflare.F("test.cloudflareaccess.com"),
			Name:                     cloudflare.F("Widget Corps Internal Applications"),
			AllowAuthenticateViaWarp: cloudflare.F(true),
			AutoRedirectToIdentity:   cloudflare.F(true),
			IsUiReadOnly:             cloudflare.F(true),
			LoginDesign: cloudflare.F(cloudflare.AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationParamsLoginDesign{
				BackgroundColor: cloudflare.F("#c5ed1b"),
				FooterText:      cloudflare.F("This is an example description."),
				HeaderText:      cloudflare.F("This is an example description."),
				LogoPath:        cloudflare.F("https://example.com/logo.png"),
				TextColor:       cloudflare.F("#c5ed1b"),
			}),
			SessionDuration:                cloudflare.F("24h"),
			UiReadOnlyToggleReason:         cloudflare.F("Temporarily turn off the UI read only lock to make a change via the UI"),
			UserSeatExpirationInactiveTime: cloudflare.F("720h"),
			WarpAuthSessionDuration:        cloudflare.F("24h"),
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

func TestAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganization(t *testing.T) {
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
	_, err := client.Access.Organizations.ZeroTrustOrganizationGetYourZeroTrustOrganization(
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

func TestAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationWithOptionalParams(t *testing.T) {
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
	_, err := client.Access.Organizations.ZeroTrustOrganizationUpdateYourZeroTrustOrganization(
		context.TODO(),
		"string",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParams{
			AllowAuthenticateViaWarp: cloudflare.F(true),
			AuthDomain:               cloudflare.F("test.cloudflareaccess.com"),
			AutoRedirectToIdentity:   cloudflare.F(true),
			CustomPages: cloudflare.F(cloudflare.AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParamsCustomPages{
				Forbidden:      cloudflare.F("699d98642c564d2e855e9661899b7252"),
				IdentityDenied: cloudflare.F("699d98642c564d2e855e9661899b7252"),
			}),
			IsUiReadOnly: cloudflare.F(true),
			LoginDesign: cloudflare.F(cloudflare.AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParamsLoginDesign{
				BackgroundColor: cloudflare.F("#c5ed1b"),
				FooterText:      cloudflare.F("This is an example description."),
				HeaderText:      cloudflare.F("This is an example description."),
				LogoPath:        cloudflare.F("https://example.com/logo.png"),
				TextColor:       cloudflare.F("#c5ed1b"),
			}),
			Name:                           cloudflare.F("Widget Corps Internal Applications"),
			SessionDuration:                cloudflare.F("24h"),
			UiReadOnlyToggleReason:         cloudflare.F("Temporarily turn off the UI read only lock to make a change via the UI"),
			UserSeatExpirationInactiveTime: cloudflare.F("720h"),
			WarpAuthSessionDuration:        cloudflare.F("24h"),
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
