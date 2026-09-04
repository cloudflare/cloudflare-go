// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/accounts"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

func TestPaymentMethodNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.PaymentMethods.New(context.TODO(), accounts.PaymentMethodNewParams{
		AccountID:           cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Address:             cloudflare.F("address"),
		Address2:            cloudflare.F("address2"),
		BankAccountType:     cloudflare.F("bank_account_type"),
		BankCode:            cloudflare.F("bank_code"),
		BankCountry:         cloudflare.F("bank_country"),
		BankName:            cloudflare.F("bank_name"),
		BankRoutingNumber:   cloudflare.F("bank_routing_number"),
		CashappCashTag:      cloudflare.F("cashapp_cash_tag"),
		City:                cloudflare.F("city"),
		Country:             cloudflare.F("country"),
		Default:             cloudflare.F(true),
		DeviceData:          cloudflare.F("device_data"),
		FirstName:           cloudflare.F("first_name"),
		LastName:            cloudflare.F("last_name"),
		NickName:            cloudflare.F("nick_name"),
		PaymentAccountEmail: cloudflare.F("payment_account_email"),
		PaymentEmail:        cloudflare.F("payment_email"),
		PaymentGateway:      cloudflare.F("payment_gateway"),
		PaymentNonce:        cloudflare.F("payment_nonce"),
		State:               cloudflare.F("state"),
		Type:                cloudflare.F(accounts.PaymentMethodNewParamsTypeCreditCard),
		Zipcode:             cloudflare.F("zipcode"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentMethodUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.PaymentMethods.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		accounts.PaymentMethodUpdateParams{
			AccountID:           cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Address:             cloudflare.F("address"),
			Address2:            cloudflare.F("address2"),
			BankAccountType:     cloudflare.F("bank_account_type"),
			BankCode:            cloudflare.F("bank_code"),
			BankCountry:         cloudflare.F("bank_country"),
			BankName:            cloudflare.F("bank_name"),
			BankRoutingNumber:   cloudflare.F("bank_routing_number"),
			CashappCashTag:      cloudflare.F("cashapp_cash_tag"),
			City:                cloudflare.F("city"),
			Country:             cloudflare.F("country"),
			Default:             cloudflare.F(true),
			DeviceData:          cloudflare.F("device_data"),
			FirstName:           cloudflare.F("first_name"),
			LastName:            cloudflare.F("last_name"),
			NickName:            cloudflare.F("nick_name"),
			PaymentAccountEmail: cloudflare.F("payment_account_email"),
			PaymentEmail:        cloudflare.F("payment_email"),
			PaymentGateway:      cloudflare.F("payment_gateway"),
			PaymentNonce:        cloudflare.F("payment_nonce"),
			State:               cloudflare.F("state"),
			Type:                cloudflare.F(accounts.PaymentMethodUpdateParamsTypeCreditCard),
			Zipcode:             cloudflare.F("zipcode"),
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

func TestPaymentMethodListWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.PaymentMethods.List(context.TODO(), accounts.PaymentMethodListParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Page:      cloudflare.F(int64(1)),
		PerPage:   cloudflare.F(int64(1)),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentMethodDelete(t *testing.T) {
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
	_, err := client.Accounts.PaymentMethods.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		accounts.PaymentMethodDeleteParams{
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

func TestPaymentMethodGet(t *testing.T) {
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
	_, err := client.Accounts.PaymentMethods.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		accounts.PaymentMethodGetParams{
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

func TestPaymentMethodSetAsDefault(t *testing.T) {
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
	_, err := client.Accounts.PaymentMethods.SetAsDefault(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		accounts.PaymentMethodSetAsDefaultParams{
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
