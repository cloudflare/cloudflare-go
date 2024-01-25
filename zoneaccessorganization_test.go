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

func TestZoneAccessOrganizationRevokeUser(t *testing.T) {
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
	_, err := client.Zones.Accesses.Organizations.RevokeUser(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cloudflare.ZoneAccessOrganizationRevokeUserParams{
			Email: cloudflare.F("test@example.com"),
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

func TestZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Accesses.Organizations.ZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganization(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cloudflare.ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationParams{
			AuthDomain:   cloudflare.F("test.cloudflareaccess.com"),
			Name:         cloudflare.F("Widget Corps Internal Applications"),
			IsUiReadOnly: cloudflare.F(true),
			LoginDesign: cloudflare.F(cloudflare.ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationParamsLoginDesign{
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

func TestZoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganization(t *testing.T) {
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
	_, err := client.Zones.Accesses.Organizations.ZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganization(context.TODO(), "699d98642c564d2e855e9661899b7252")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Accesses.Organizations.ZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganization(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cloudflare.ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationParams{
			AuthDomain:   cloudflare.F("test.cloudflareaccess.com"),
			IsUiReadOnly: cloudflare.F(true),
			LoginDesign: cloudflare.F(cloudflare.ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationParamsLoginDesign{
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
