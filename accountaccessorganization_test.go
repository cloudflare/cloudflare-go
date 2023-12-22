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

func TestAccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Accounts.Accesses.Organizations.ZeroTrustOrganizationNewYourZeroTrustOrganization(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cloudflare.AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationParams{
			AuthDomain:             cloudflare.F("test.cloudflareaccess.com"),
			Name:                   cloudflare.F("Widget Corps Internal Applications"),
			AutoRedirectToIdentity: cloudflare.F(true),
			IsUiReadOnly:           cloudflare.F(true),
			LoginDesign: cloudflare.F(cloudflare.AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationParamsLoginDesign{
				BackgroundColor: cloudflare.F("#c5ed1b"),
				FooterText:      cloudflare.F("This is an example description."),
				HeaderText:      cloudflare.F("This is an example description."),
				LogoPath:        cloudflare.F("https://example.com/logo.png"),
				TextColor:       cloudflare.F("#c5ed1b"),
			}),
			UiReadOnlyToggleReason:         cloudflare.F("Temporarily turn off the UI read only lock to make a change via the UI"),
			UserSeatExpirationInactiveTime: cloudflare.F("720h"),
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

func TestAccountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganization(t *testing.T) {
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
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Accounts.Accesses.Organizations.ZeroTrustOrganizationGetYourZeroTrustOrganization(context.TODO(), "699d98642c564d2e855e9661899b7252")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Accounts.Accesses.Organizations.ZeroTrustOrganizationUpdateYourZeroTrustOrganization(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cloudflare.AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParams{
			AuthDomain:             cloudflare.F("test.cloudflareaccess.com"),
			AutoRedirectToIdentity: cloudflare.F(true),
			IsUiReadOnly:           cloudflare.F(true),
			LoginDesign: cloudflare.F(cloudflare.AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParamsLoginDesign{
				BackgroundColor: cloudflare.F("#c5ed1b"),
				FooterText:      cloudflare.F("This is an example description."),
				HeaderText:      cloudflare.F("This is an example description."),
				LogoPath:        cloudflare.F("https://example.com/logo.png"),
				TextColor:       cloudflare.F("#c5ed1b"),
			}),
			Name:                           cloudflare.F("Widget Corps Internal Applications"),
			UiReadOnlyToggleReason:         cloudflare.F("Temporarily turn off the UI read only lock to make a change via the UI"),
			UserSeatExpirationInactiveTime: cloudflare.F("720h"),
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
