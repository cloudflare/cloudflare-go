// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
)

func TestCardNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Cards.New(context.TODO(), cloudflare.CardNewParams{
		Type:             cloudflare.F(cloudflare.CardNewParamsTypeVirtual),
		AccountToken:     cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		CardProgramToken: cloudflare.F("00000000-0000-0000-1000-000000000000"),
		Carrier: cloudflare.F(cloudflare.CardNewParamsCarrier{
			QrCodeURL: cloudflare.F("string"),
		}),
		DigitalCardArtToken: cloudflare.F("00000000-0000-0000-1000-000000000000"),
		ExpMonth:            cloudflare.F("06"),
		ExpYear:             cloudflare.F("2027"),
		Memo:                cloudflare.F("New Card"),
		Pin:                 cloudflare.F("string"),
		ProductID:           cloudflare.F("1"),
		ShippingAddress: cloudflare.F(cloudflare.CardNewParamsShippingAddress{
			FirstName:   cloudflare.F("Michael"),
			LastName:    cloudflare.F("Bluth"),
			Line2Text:   cloudflare.F("The Bluth Company"),
			Address1:    cloudflare.F("5 Broad Street"),
			Address2:    cloudflare.F("Unit 25A"),
			City:        cloudflare.F("NEW YORK"),
			State:       cloudflare.F("NY"),
			PostalCode:  cloudflare.F("10001-1809"),
			Country:     cloudflare.F("USA"),
			Email:       cloudflare.F("johnny@appleseed.com"),
			PhoneNumber: cloudflare.F("+12124007676"),
		}),
		ShippingMethod:     cloudflare.F(cloudflare.CardNewParamsShippingMethodStandard),
		SpendLimit:         cloudflare.F(int64(1000)),
		SpendLimitDuration: cloudflare.F(cloudflare.CardNewParamsSpendLimitDurationTransaction),
		State:              cloudflare.F(cloudflare.CardNewParamsStateOpen),
		IdempotencyKey:     cloudflare.F("string"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Cards.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardUpdateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Cards.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		cloudflare.CardUpdateParams{
			AuthRuleToken:       cloudflare.F("string"),
			DigitalCardArtToken: cloudflare.F("00000000-0000-0000-1000-000000000000"),
			Memo:                cloudflare.F("Updated Name"),
			Pin:                 cloudflare.F("string"),
			SpendLimit:          cloudflare.F(int64(100)),
			SpendLimitDuration:  cloudflare.F(cloudflare.CardUpdateParamsSpendLimitDurationForever),
			State:               cloudflare.F(cloudflare.CardUpdateParamsStateOpen),
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

func TestCardProvisionWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Cards.Provision(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		cloudflare.CardProvisionParams{
			Certificate:    cloudflare.F("U3RhaW5sZXNzIHJvY2tz"),
			DigitalWallet:  cloudflare.F(cloudflare.CardProvisionParamsDigitalWalletGooglePay),
			Nonce:          cloudflare.F("U3RhaW5sZXNzIHJvY2tz"),
			NonceSignature: cloudflare.F("U3RhaW5sZXNzIHJvY2tz"),
			IdempotencyKey: cloudflare.F("string"),
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
