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

func TestOrganizationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Organizations.New(context.TODO(), zero_trust.OrganizationNewParams{
		AuthDomain:               cloudflare.F("test.cloudflareaccess.com"),
		Name:                     cloudflare.F("Widget Corps Internal Applications"),
		AccountID:                cloudflare.F("account_id"),
		AllowAuthenticateViaWARP: cloudflare.F(true),
		AutoRedirectToIdentity:   cloudflare.F(true),
		IsUIReadOnly:             cloudflare.F(true),
		LoginDesign: cloudflare.F(zero_trust.LoginDesignParam{
			BackgroundColor: cloudflare.F("#c5ed1b"),
			FooterText:      cloudflare.F("This is an example description."),
			HeaderText:      cloudflare.F("This is an example description."),
			LogoPath:        cloudflare.F("https://example.com/logo.png"),
			TextColor:       cloudflare.F("#c5ed1b"),
		}),
		SessionDuration:                cloudflare.F("24h"),
		UIReadOnlyToggleReason:         cloudflare.F("Temporarily turn off the UI read only lock to make a change via the UI"),
		UserSeatExpirationInactiveTime: cloudflare.F("730h"),
		WARPAuthSessionDuration:        cloudflare.F("24h"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Organizations.Update(context.TODO(), zero_trust.OrganizationUpdateParams{
		AccountID:                cloudflare.F("account_id"),
		AllowAuthenticateViaWARP: cloudflare.F(true),
		AuthDomain:               cloudflare.F("test.cloudflareaccess.com"),
		AutoRedirectToIdentity:   cloudflare.F(true),
		CustomPages: cloudflare.F(zero_trust.OrganizationUpdateParamsCustomPages{
			Forbidden:      cloudflare.F("699d98642c564d2e855e9661899b7252"),
			IdentityDenied: cloudflare.F("699d98642c564d2e855e9661899b7252"),
		}),
		IsUIReadOnly: cloudflare.F(true),
		LoginDesign: cloudflare.F(zero_trust.LoginDesignParam{
			BackgroundColor: cloudflare.F("#c5ed1b"),
			FooterText:      cloudflare.F("This is an example description."),
			HeaderText:      cloudflare.F("This is an example description."),
			LogoPath:        cloudflare.F("https://example.com/logo.png"),
			TextColor:       cloudflare.F("#c5ed1b"),
		}),
		Name:                           cloudflare.F("Widget Corps Internal Applications"),
		SessionDuration:                cloudflare.F("24h"),
		UIReadOnlyToggleReason:         cloudflare.F("Temporarily turn off the UI read only lock to make a change via the UI"),
		UserSeatExpirationInactiveTime: cloudflare.F("730h"),
		WARPAuthSessionDuration:        cloudflare.F("24h"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationListWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Organizations.List(context.TODO(), zero_trust.OrganizationListParams{
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

func TestOrganizationRevokeUsersWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Organizations.RevokeUsers(context.TODO(), zero_trust.OrganizationRevokeUsersParams{
		Email:             cloudflare.F("test@example.com"),
		AccountID:         cloudflare.F("account_id"),
		QueryDevices:      cloudflare.F(true),
		BodyDevices:       cloudflare.F(true),
		UserUID:           cloudflare.F("699d98642c564d2e855e9661899b7252"),
		WARPSessionReauth: cloudflare.F(true),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
