// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package billing_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/billing"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

func TestProfileNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Billing.Profiles.New(context.TODO(), billing.ProfileNewParams{
		AccountID:             cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Address:               cloudflare.F("123 Main Street"),
		Address2:              cloudflare.F("Apt 1"),
		BillingEmail:          cloudflare.F("billing@example.com"),
		BuyingRatePlan:        cloudflare.F("buying_rate_plan"),
		CaptchaChallengeJWT:   cloudflare.F("captcha_challenge_jwt"),
		CfTurnstileResponse:   cloudflare.F("cf_turnstile_response"),
		City:                  cloudflare.F("Anytown"),
		Company:               cloudflare.F("Example Inc"),
		Country:               cloudflare.F("US"),
		FirstName:             cloudflare.F("John"),
		HCaptchaResponse:      cloudflare.F("h_captcha_response"),
		LastName:              cloudflare.F("Doe"),
		PreferredLocale:       cloudflare.F("en-US"),
		SecondaryBillingEmail: cloudflare.F("secondary@example.com"),
		State:                 cloudflare.F("CA"),
		TaxIDType:             cloudflare.F("tax_id_type"),
		Telephone:             cloudflare.F("+1-555-555-5555"),
		Vat:                   cloudflare.F("vat"),
		Zipcode:               cloudflare.F("94103"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProfileUpdateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Billing.Profiles.Update(context.TODO(), billing.ProfileUpdateParams{
		AccountID:             cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Address:               cloudflare.F("123 Main Street"),
		Address2:              cloudflare.F("Apt 1"),
		BillingEmail:          cloudflare.F("billing@example.com"),
		BuyingRatePlan:        cloudflare.F("buying_rate_plan"),
		CaptchaChallengeJWT:   cloudflare.F("captcha_challenge_jwt"),
		CfTurnstileResponse:   cloudflare.F("cf_turnstile_response"),
		City:                  cloudflare.F("Anytown"),
		Company:               cloudflare.F("Example Inc"),
		Country:               cloudflare.F("US"),
		FirstName:             cloudflare.F("John"),
		HCaptchaResponse:      cloudflare.F("h_captcha_response"),
		LastName:              cloudflare.F("Doe"),
		PreferredLocale:       cloudflare.F("en-US"),
		SecondaryBillingEmail: cloudflare.F("secondary@example.com"),
		State:                 cloudflare.F("CA"),
		TaxIDType:             cloudflare.F("tax_id_type"),
		Telephone:             cloudflare.F("+1-555-555-5555"),
		Vat:                   cloudflare.F("vat"),
		Zipcode:               cloudflare.F("94103"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProfileDelete(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	err := client.Billing.Profiles.Delete(context.TODO(), billing.ProfileDeleteParams{
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

func TestProfileGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Billing.Profiles.Get(context.TODO(), billing.ProfileGetParams{
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

func TestProfileUpdateBillingEmailWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Billing.Profiles.UpdateBillingEmail(context.TODO(), billing.ProfileUpdateBillingEmailParams{
		AccountID:             cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		BillingEmail:          cloudflare.F("billing@example.com"),
		PreferredLocale:       cloudflare.F("en-US"),
		SecondaryBillingEmail: cloudflare.F("secondary@example.com"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
