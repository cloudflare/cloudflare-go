// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v7/shared"
)

// PaymentMethodService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentMethodService] method instead.
type PaymentMethodService struct {
	Options []option.RequestOption
}

// NewPaymentMethodService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPaymentMethodService(opts ...option.RequestOption) (r *PaymentMethodService) {
	r = &PaymentMethodService{}
	r.Options = opts
	return
}

// Creates a new payment method for an account.
func (r *PaymentMethodService) New(ctx context.Context, params PaymentMethodNewParams, opts ...option.RequestOption) (res *PaymentMethodNewResponse, err error) {
	var env PaymentMethodNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/payment-methods", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates a payment method for an account.
func (r *PaymentMethodService) Update(ctx context.Context, paymentMethodID string, params PaymentMethodUpdateParams, opts ...option.RequestOption) (res *PaymentMethodUpdateResponse, err error) {
	var env PaymentMethodUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if paymentMethodID == "" {
		err = errors.New("missing required payment_method_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/payment-methods/%s", params.AccountID, paymentMethodID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists all payment methods for an account.
func (r *PaymentMethodService) List(ctx context.Context, params PaymentMethodListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[PaymentMethodListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/payment-methods", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists all payment methods for an account.
func (r *PaymentMethodService) ListAutoPaging(ctx context.Context, params PaymentMethodListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[PaymentMethodListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes a payment method from an account.
func (r *PaymentMethodService) Delete(ctx context.Context, paymentMethodID string, body PaymentMethodDeleteParams, opts ...option.RequestOption) (res *interface{}, err error) {
	var env PaymentMethodDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if paymentMethodID == "" {
		err = errors.New("missing required payment_method_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/payment-methods/%s", body.AccountID, paymentMethodID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Gets a specific payment method for an account.
func (r *PaymentMethodService) Get(ctx context.Context, paymentMethodID string, query PaymentMethodGetParams, opts ...option.RequestOption) (res *PaymentMethodGetResponse, err error) {
	var env PaymentMethodGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if paymentMethodID == "" {
		err = errors.New("missing required payment_method_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/payment-methods/%s", query.AccountID, paymentMethodID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Sets a payment method as the default for an account.
func (r *PaymentMethodService) SetAsDefault(ctx context.Context, paymentMethodID string, body PaymentMethodSetAsDefaultParams, opts ...option.RequestOption) (res *interface{}, err error) {
	var env PaymentMethodSetAsDefaultResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if paymentMethodID == "" {
		err = errors.New("missing required payment_method_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/payment-methods/%s/set-as-default", body.AccountID, paymentMethodID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type PaymentMethodNewResponse struct {
	// Payment method identifier.
	ID string `json:"id"`
	// Billing address line 1.
	Address string `json:"address"`
	// Billing address line 2.
	Address2 string `json:"address2"`
	// Bank account type.
	BankAccountType string `json:"bank_account_type"`
	// Bank code.
	BankCode string `json:"bank_code"`
	// Bank country.
	BankCountry string `json:"bank_country"`
	// Bank name for bank-based payment methods.
	BankName string `json:"bank_name"`
	// Bank routing number.
	BankRoutingNumber string `json:"bank_routing_number"`
	// Cash App cash tag.
	CashappCashTag string `json:"cashapp_cash_tag"`
	// Billing city.
	City string `json:"city"`
	// Billing country.
	Country string `json:"country"`
	// Whether this is the default payment method.
	Default bool `json:"default"`
	// Card expiration date.
	ExpirationDate string `json:"expiration_date"`
	// Billing first name.
	FirstName string `json:"first_name"`
	// Last four digits of the card number.
	LastFour string `json:"last_four"`
	// Billing last name.
	LastName string `json:"last_name"`
	// A nickname for the payment method.
	NickName string `json:"nick_name"`
	// Email associated with the payment account.
	PaymentAccountEmail string `json:"payment_account_email"`
	// Payment email address.
	PaymentEmail string `json:"payment_email"`
	// Billing state.
	State string `json:"state"`
	// The payment method type.
	Type PaymentMethodNewResponseType `json:"type"`
	// Billing zip code.
	Zipcode string                       `json:"zipcode"`
	JSON    paymentMethodNewResponseJSON `json:"-"`
}

// paymentMethodNewResponseJSON contains the JSON metadata for the struct
// [PaymentMethodNewResponse]
type paymentMethodNewResponseJSON struct {
	ID                  apijson.Field
	Address             apijson.Field
	Address2            apijson.Field
	BankAccountType     apijson.Field
	BankCode            apijson.Field
	BankCountry         apijson.Field
	BankName            apijson.Field
	BankRoutingNumber   apijson.Field
	CashappCashTag      apijson.Field
	City                apijson.Field
	Country             apijson.Field
	Default             apijson.Field
	ExpirationDate      apijson.Field
	FirstName           apijson.Field
	LastFour            apijson.Field
	LastName            apijson.Field
	NickName            apijson.Field
	PaymentAccountEmail apijson.Field
	PaymentEmail        apijson.Field
	State               apijson.Field
	Type                apijson.Field
	Zipcode             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *PaymentMethodNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentMethodNewResponseJSON) RawJSON() string {
	return r.raw
}

// The payment method type.
type PaymentMethodNewResponseType string

const (
	PaymentMethodNewResponseTypeCreditCard     PaymentMethodNewResponseType = "CREDIT_CARD"
	PaymentMethodNewResponseTypePaypal         PaymentMethodNewResponseType = "PAYPAL"
	PaymentMethodNewResponseTypeCashapp        PaymentMethodNewResponseType = "CASHAPP"
	PaymentMethodNewResponseTypeSepaDebit      PaymentMethodNewResponseType = "SEPA_DEBIT"
	PaymentMethodNewResponseTypeLink           PaymentMethodNewResponseType = "LINK"
	PaymentMethodNewResponseTypeACHDirectDebit PaymentMethodNewResponseType = "ACH_DIRECT_DEBIT"
)

func (r PaymentMethodNewResponseType) IsKnown() bool {
	switch r {
	case PaymentMethodNewResponseTypeCreditCard, PaymentMethodNewResponseTypePaypal, PaymentMethodNewResponseTypeCashapp, PaymentMethodNewResponseTypeSepaDebit, PaymentMethodNewResponseTypeLink, PaymentMethodNewResponseTypeACHDirectDebit:
		return true
	}
	return false
}

type PaymentMethodUpdateResponse struct {
	// Payment method identifier.
	ID string `json:"id"`
	// Billing address line 1.
	Address string `json:"address"`
	// Billing address line 2.
	Address2 string `json:"address2"`
	// Bank account type.
	BankAccountType string `json:"bank_account_type"`
	// Bank code.
	BankCode string `json:"bank_code"`
	// Bank country.
	BankCountry string `json:"bank_country"`
	// Bank name for bank-based payment methods.
	BankName string `json:"bank_name"`
	// Bank routing number.
	BankRoutingNumber string `json:"bank_routing_number"`
	// Cash App cash tag.
	CashappCashTag string `json:"cashapp_cash_tag"`
	// Billing city.
	City string `json:"city"`
	// Billing country.
	Country string `json:"country"`
	// Whether this is the default payment method.
	Default bool `json:"default"`
	// Card expiration date.
	ExpirationDate string `json:"expiration_date"`
	// Billing first name.
	FirstName string `json:"first_name"`
	// Last four digits of the card number.
	LastFour string `json:"last_four"`
	// Billing last name.
	LastName string `json:"last_name"`
	// A nickname for the payment method.
	NickName string `json:"nick_name"`
	// Email associated with the payment account.
	PaymentAccountEmail string `json:"payment_account_email"`
	// Payment email address.
	PaymentEmail string `json:"payment_email"`
	// Billing state.
	State string `json:"state"`
	// The payment method type.
	Type PaymentMethodUpdateResponseType `json:"type"`
	// Billing zip code.
	Zipcode string                          `json:"zipcode"`
	JSON    paymentMethodUpdateResponseJSON `json:"-"`
}

// paymentMethodUpdateResponseJSON contains the JSON metadata for the struct
// [PaymentMethodUpdateResponse]
type paymentMethodUpdateResponseJSON struct {
	ID                  apijson.Field
	Address             apijson.Field
	Address2            apijson.Field
	BankAccountType     apijson.Field
	BankCode            apijson.Field
	BankCountry         apijson.Field
	BankName            apijson.Field
	BankRoutingNumber   apijson.Field
	CashappCashTag      apijson.Field
	City                apijson.Field
	Country             apijson.Field
	Default             apijson.Field
	ExpirationDate      apijson.Field
	FirstName           apijson.Field
	LastFour            apijson.Field
	LastName            apijson.Field
	NickName            apijson.Field
	PaymentAccountEmail apijson.Field
	PaymentEmail        apijson.Field
	State               apijson.Field
	Type                apijson.Field
	Zipcode             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *PaymentMethodUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentMethodUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The payment method type.
type PaymentMethodUpdateResponseType string

const (
	PaymentMethodUpdateResponseTypeCreditCard     PaymentMethodUpdateResponseType = "CREDIT_CARD"
	PaymentMethodUpdateResponseTypePaypal         PaymentMethodUpdateResponseType = "PAYPAL"
	PaymentMethodUpdateResponseTypeCashapp        PaymentMethodUpdateResponseType = "CASHAPP"
	PaymentMethodUpdateResponseTypeSepaDebit      PaymentMethodUpdateResponseType = "SEPA_DEBIT"
	PaymentMethodUpdateResponseTypeLink           PaymentMethodUpdateResponseType = "LINK"
	PaymentMethodUpdateResponseTypeACHDirectDebit PaymentMethodUpdateResponseType = "ACH_DIRECT_DEBIT"
)

func (r PaymentMethodUpdateResponseType) IsKnown() bool {
	switch r {
	case PaymentMethodUpdateResponseTypeCreditCard, PaymentMethodUpdateResponseTypePaypal, PaymentMethodUpdateResponseTypeCashapp, PaymentMethodUpdateResponseTypeSepaDebit, PaymentMethodUpdateResponseTypeLink, PaymentMethodUpdateResponseTypeACHDirectDebit:
		return true
	}
	return false
}

type PaymentMethodListResponse struct {
	// Payment method identifier.
	ID string `json:"id"`
	// Billing address line 1.
	Address string `json:"address"`
	// Billing address line 2.
	Address2 string `json:"address2"`
	// Bank account type.
	BankAccountType string `json:"bank_account_type"`
	// Bank code.
	BankCode string `json:"bank_code"`
	// Bank country.
	BankCountry string `json:"bank_country"`
	// Bank name for bank-based payment methods.
	BankName string `json:"bank_name"`
	// Bank routing number.
	BankRoutingNumber string `json:"bank_routing_number"`
	// Cash App cash tag.
	CashappCashTag string `json:"cashapp_cash_tag"`
	// Billing city.
	City string `json:"city"`
	// Billing country.
	Country string `json:"country"`
	// Whether this is the default payment method.
	Default bool `json:"default"`
	// Card expiration date.
	ExpirationDate string `json:"expiration_date"`
	// Billing first name.
	FirstName string `json:"first_name"`
	// Last four digits of the card number.
	LastFour string `json:"last_four"`
	// Billing last name.
	LastName string `json:"last_name"`
	// A nickname for the payment method.
	NickName string `json:"nick_name"`
	// Email associated with the payment account.
	PaymentAccountEmail string `json:"payment_account_email"`
	// Payment email address.
	PaymentEmail string `json:"payment_email"`
	// Billing state.
	State string `json:"state"`
	// The payment method type.
	Type PaymentMethodListResponseType `json:"type"`
	// Billing zip code.
	Zipcode string                        `json:"zipcode"`
	JSON    paymentMethodListResponseJSON `json:"-"`
}

// paymentMethodListResponseJSON contains the JSON metadata for the struct
// [PaymentMethodListResponse]
type paymentMethodListResponseJSON struct {
	ID                  apijson.Field
	Address             apijson.Field
	Address2            apijson.Field
	BankAccountType     apijson.Field
	BankCode            apijson.Field
	BankCountry         apijson.Field
	BankName            apijson.Field
	BankRoutingNumber   apijson.Field
	CashappCashTag      apijson.Field
	City                apijson.Field
	Country             apijson.Field
	Default             apijson.Field
	ExpirationDate      apijson.Field
	FirstName           apijson.Field
	LastFour            apijson.Field
	LastName            apijson.Field
	NickName            apijson.Field
	PaymentAccountEmail apijson.Field
	PaymentEmail        apijson.Field
	State               apijson.Field
	Type                apijson.Field
	Zipcode             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *PaymentMethodListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentMethodListResponseJSON) RawJSON() string {
	return r.raw
}

// The payment method type.
type PaymentMethodListResponseType string

const (
	PaymentMethodListResponseTypeCreditCard     PaymentMethodListResponseType = "CREDIT_CARD"
	PaymentMethodListResponseTypePaypal         PaymentMethodListResponseType = "PAYPAL"
	PaymentMethodListResponseTypeCashapp        PaymentMethodListResponseType = "CASHAPP"
	PaymentMethodListResponseTypeSepaDebit      PaymentMethodListResponseType = "SEPA_DEBIT"
	PaymentMethodListResponseTypeLink           PaymentMethodListResponseType = "LINK"
	PaymentMethodListResponseTypeACHDirectDebit PaymentMethodListResponseType = "ACH_DIRECT_DEBIT"
)

func (r PaymentMethodListResponseType) IsKnown() bool {
	switch r {
	case PaymentMethodListResponseTypeCreditCard, PaymentMethodListResponseTypePaypal, PaymentMethodListResponseTypeCashapp, PaymentMethodListResponseTypeSepaDebit, PaymentMethodListResponseTypeLink, PaymentMethodListResponseTypeACHDirectDebit:
		return true
	}
	return false
}

type PaymentMethodGetResponse struct {
	// Payment method identifier.
	ID string `json:"id"`
	// Billing address line 1.
	Address string `json:"address"`
	// Billing address line 2.
	Address2 string `json:"address2"`
	// Bank account type.
	BankAccountType string `json:"bank_account_type"`
	// Bank code.
	BankCode string `json:"bank_code"`
	// Bank country.
	BankCountry string `json:"bank_country"`
	// Bank name for bank-based payment methods.
	BankName string `json:"bank_name"`
	// Bank routing number.
	BankRoutingNumber string `json:"bank_routing_number"`
	// Cash App cash tag.
	CashappCashTag string `json:"cashapp_cash_tag"`
	// Billing city.
	City string `json:"city"`
	// Billing country.
	Country string `json:"country"`
	// Whether this is the default payment method.
	Default bool `json:"default"`
	// Card expiration date.
	ExpirationDate string `json:"expiration_date"`
	// Billing first name.
	FirstName string `json:"first_name"`
	// Last four digits of the card number.
	LastFour string `json:"last_four"`
	// Billing last name.
	LastName string `json:"last_name"`
	// A nickname for the payment method.
	NickName string `json:"nick_name"`
	// Email associated with the payment account.
	PaymentAccountEmail string `json:"payment_account_email"`
	// Payment email address.
	PaymentEmail string `json:"payment_email"`
	// Billing state.
	State string `json:"state"`
	// The payment method type.
	Type PaymentMethodGetResponseType `json:"type"`
	// Billing zip code.
	Zipcode string                       `json:"zipcode"`
	JSON    paymentMethodGetResponseJSON `json:"-"`
}

// paymentMethodGetResponseJSON contains the JSON metadata for the struct
// [PaymentMethodGetResponse]
type paymentMethodGetResponseJSON struct {
	ID                  apijson.Field
	Address             apijson.Field
	Address2            apijson.Field
	BankAccountType     apijson.Field
	BankCode            apijson.Field
	BankCountry         apijson.Field
	BankName            apijson.Field
	BankRoutingNumber   apijson.Field
	CashappCashTag      apijson.Field
	City                apijson.Field
	Country             apijson.Field
	Default             apijson.Field
	ExpirationDate      apijson.Field
	FirstName           apijson.Field
	LastFour            apijson.Field
	LastName            apijson.Field
	NickName            apijson.Field
	PaymentAccountEmail apijson.Field
	PaymentEmail        apijson.Field
	State               apijson.Field
	Type                apijson.Field
	Zipcode             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *PaymentMethodGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentMethodGetResponseJSON) RawJSON() string {
	return r.raw
}

// The payment method type.
type PaymentMethodGetResponseType string

const (
	PaymentMethodGetResponseTypeCreditCard     PaymentMethodGetResponseType = "CREDIT_CARD"
	PaymentMethodGetResponseTypePaypal         PaymentMethodGetResponseType = "PAYPAL"
	PaymentMethodGetResponseTypeCashapp        PaymentMethodGetResponseType = "CASHAPP"
	PaymentMethodGetResponseTypeSepaDebit      PaymentMethodGetResponseType = "SEPA_DEBIT"
	PaymentMethodGetResponseTypeLink           PaymentMethodGetResponseType = "LINK"
	PaymentMethodGetResponseTypeACHDirectDebit PaymentMethodGetResponseType = "ACH_DIRECT_DEBIT"
)

func (r PaymentMethodGetResponseType) IsKnown() bool {
	switch r {
	case PaymentMethodGetResponseTypeCreditCard, PaymentMethodGetResponseTypePaypal, PaymentMethodGetResponseTypeCashapp, PaymentMethodGetResponseTypeSepaDebit, PaymentMethodGetResponseTypeLink, PaymentMethodGetResponseTypeACHDirectDebit:
		return true
	}
	return false
}

type PaymentMethodNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Billing address line 1.
	Address param.Field[string] `json:"address"`
	// Billing address line 2.
	Address2 param.Field[string] `json:"address2"`
	// Bank account type.
	BankAccountType param.Field[string] `json:"bank_account_type"`
	// Bank code.
	BankCode param.Field[string] `json:"bank_code"`
	// Bank country.
	BankCountry param.Field[string] `json:"bank_country"`
	// Bank name for bank-based payment methods.
	BankName param.Field[string] `json:"bank_name"`
	// Bank routing number.
	BankRoutingNumber param.Field[string] `json:"bank_routing_number"`
	// Cash App cash tag.
	CashappCashTag param.Field[string] `json:"cashapp_cash_tag"`
	// Billing city.
	City param.Field[string] `json:"city"`
	// Billing country.
	Country param.Field[string] `json:"country"`
	// Whether this is the default payment method.
	Default param.Field[bool] `json:"default"`
	// Device data for fraud prevention.
	DeviceData param.Field[string] `json:"device_data"`
	// Billing first name.
	FirstName param.Field[string] `json:"first_name"`
	// Billing last name.
	LastName param.Field[string] `json:"last_name"`
	// A nickname for the payment method.
	NickName param.Field[string] `json:"nick_name"`
	// Email associated with the payment account.
	PaymentAccountEmail param.Field[string] `json:"payment_account_email"`
	// Payment email address.
	PaymentEmail param.Field[string] `json:"payment_email"`
	// The payment gateway used.
	PaymentGateway param.Field[string] `json:"payment_gateway"`
	// Payment nonce for tokenized payments.
	PaymentNonce param.Field[string] `json:"payment_nonce"`
	// Billing state.
	State param.Field[string] `json:"state"`
	// The payment method type.
	Type param.Field[PaymentMethodNewParamsType] `json:"type"`
	// Billing zip code.
	Zipcode param.Field[string] `json:"zipcode"`
}

func (r PaymentMethodNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The payment method type.
type PaymentMethodNewParamsType string

const (
	PaymentMethodNewParamsTypeCreditCard     PaymentMethodNewParamsType = "CREDIT_CARD"
	PaymentMethodNewParamsTypePaypal         PaymentMethodNewParamsType = "PAYPAL"
	PaymentMethodNewParamsTypeCashapp        PaymentMethodNewParamsType = "CASHAPP"
	PaymentMethodNewParamsTypeSepaDebit      PaymentMethodNewParamsType = "SEPA_DEBIT"
	PaymentMethodNewParamsTypeLink           PaymentMethodNewParamsType = "LINK"
	PaymentMethodNewParamsTypeACHDirectDebit PaymentMethodNewParamsType = "ACH_DIRECT_DEBIT"
)

func (r PaymentMethodNewParamsType) IsKnown() bool {
	switch r {
	case PaymentMethodNewParamsTypeCreditCard, PaymentMethodNewParamsTypePaypal, PaymentMethodNewParamsTypeCashapp, PaymentMethodNewParamsTypeSepaDebit, PaymentMethodNewParamsTypeLink, PaymentMethodNewParamsTypeACHDirectDebit:
		return true
	}
	return false
}

type PaymentMethodNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo    `json:"errors" api:"required"`
	Messages []shared.ResponseInfo    `json:"messages" api:"required"`
	Result   PaymentMethodNewResponse `json:"result" api:"required"`
	// Whether the API call was successful
	Success PaymentMethodNewResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    paymentMethodNewResponseEnvelopeJSON    `json:"-"`
}

// paymentMethodNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [PaymentMethodNewResponseEnvelope]
type paymentMethodNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentMethodNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentMethodNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PaymentMethodNewResponseEnvelopeSuccess bool

const (
	PaymentMethodNewResponseEnvelopeSuccessTrue PaymentMethodNewResponseEnvelopeSuccess = true
)

func (r PaymentMethodNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PaymentMethodNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PaymentMethodUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Billing address line 1.
	Address param.Field[string] `json:"address"`
	// Billing address line 2.
	Address2 param.Field[string] `json:"address2"`
	// Bank account type.
	BankAccountType param.Field[string] `json:"bank_account_type"`
	// Bank code.
	BankCode param.Field[string] `json:"bank_code"`
	// Bank country.
	BankCountry param.Field[string] `json:"bank_country"`
	// Bank name for bank-based payment methods.
	BankName param.Field[string] `json:"bank_name"`
	// Bank routing number.
	BankRoutingNumber param.Field[string] `json:"bank_routing_number"`
	// Cash App cash tag.
	CashappCashTag param.Field[string] `json:"cashapp_cash_tag"`
	// Billing city.
	City param.Field[string] `json:"city"`
	// Billing country.
	Country param.Field[string] `json:"country"`
	// Whether this is the default payment method.
	Default param.Field[bool] `json:"default"`
	// Device data for fraud prevention.
	DeviceData param.Field[string] `json:"device_data"`
	// Billing first name.
	FirstName param.Field[string] `json:"first_name"`
	// Billing last name.
	LastName param.Field[string] `json:"last_name"`
	// A nickname for the payment method.
	NickName param.Field[string] `json:"nick_name"`
	// Email associated with the payment account.
	PaymentAccountEmail param.Field[string] `json:"payment_account_email"`
	// Payment email address.
	PaymentEmail param.Field[string] `json:"payment_email"`
	// The payment gateway used.
	PaymentGateway param.Field[string] `json:"payment_gateway"`
	// Payment nonce for tokenized payments.
	PaymentNonce param.Field[string] `json:"payment_nonce"`
	// Billing state.
	State param.Field[string] `json:"state"`
	// The payment method type.
	Type param.Field[PaymentMethodUpdateParamsType] `json:"type"`
	// Billing zip code.
	Zipcode param.Field[string] `json:"zipcode"`
}

func (r PaymentMethodUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The payment method type.
type PaymentMethodUpdateParamsType string

const (
	PaymentMethodUpdateParamsTypeCreditCard     PaymentMethodUpdateParamsType = "CREDIT_CARD"
	PaymentMethodUpdateParamsTypePaypal         PaymentMethodUpdateParamsType = "PAYPAL"
	PaymentMethodUpdateParamsTypeCashapp        PaymentMethodUpdateParamsType = "CASHAPP"
	PaymentMethodUpdateParamsTypeSepaDebit      PaymentMethodUpdateParamsType = "SEPA_DEBIT"
	PaymentMethodUpdateParamsTypeLink           PaymentMethodUpdateParamsType = "LINK"
	PaymentMethodUpdateParamsTypeACHDirectDebit PaymentMethodUpdateParamsType = "ACH_DIRECT_DEBIT"
)

func (r PaymentMethodUpdateParamsType) IsKnown() bool {
	switch r {
	case PaymentMethodUpdateParamsTypeCreditCard, PaymentMethodUpdateParamsTypePaypal, PaymentMethodUpdateParamsTypeCashapp, PaymentMethodUpdateParamsTypeSepaDebit, PaymentMethodUpdateParamsTypeLink, PaymentMethodUpdateParamsTypeACHDirectDebit:
		return true
	}
	return false
}

type PaymentMethodUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo       `json:"errors" api:"required"`
	Messages []shared.ResponseInfo       `json:"messages" api:"required"`
	Result   PaymentMethodUpdateResponse `json:"result" api:"required"`
	// Whether the API call was successful
	Success PaymentMethodUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    paymentMethodUpdateResponseEnvelopeJSON    `json:"-"`
}

// paymentMethodUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [PaymentMethodUpdateResponseEnvelope]
type paymentMethodUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentMethodUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentMethodUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PaymentMethodUpdateResponseEnvelopeSuccess bool

const (
	PaymentMethodUpdateResponseEnvelopeSuccessTrue PaymentMethodUpdateResponseEnvelopeSuccess = true
)

func (r PaymentMethodUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PaymentMethodUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PaymentMethodListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [PaymentMethodListParams]'s query parameters as
// `url.Values`.
func (r PaymentMethodListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type PaymentMethodDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type PaymentMethodDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	Result   interface{}           `json:"result" api:"required"`
	// Whether the API call was successful
	Success PaymentMethodDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    paymentMethodDeleteResponseEnvelopeJSON    `json:"-"`
}

// paymentMethodDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [PaymentMethodDeleteResponseEnvelope]
type paymentMethodDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentMethodDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentMethodDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PaymentMethodDeleteResponseEnvelopeSuccess bool

const (
	PaymentMethodDeleteResponseEnvelopeSuccessTrue PaymentMethodDeleteResponseEnvelopeSuccess = true
)

func (r PaymentMethodDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PaymentMethodDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PaymentMethodGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type PaymentMethodGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo    `json:"errors" api:"required"`
	Messages []shared.ResponseInfo    `json:"messages" api:"required"`
	Result   PaymentMethodGetResponse `json:"result" api:"required"`
	// Whether the API call was successful
	Success PaymentMethodGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    paymentMethodGetResponseEnvelopeJSON    `json:"-"`
}

// paymentMethodGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PaymentMethodGetResponseEnvelope]
type paymentMethodGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentMethodGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentMethodGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PaymentMethodGetResponseEnvelopeSuccess bool

const (
	PaymentMethodGetResponseEnvelopeSuccessTrue PaymentMethodGetResponseEnvelopeSuccess = true
)

func (r PaymentMethodGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PaymentMethodGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PaymentMethodSetAsDefaultParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type PaymentMethodSetAsDefaultResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	Result   interface{}           `json:"result" api:"required"`
	// Whether the API call was successful
	Success PaymentMethodSetAsDefaultResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    paymentMethodSetAsDefaultResponseEnvelopeJSON    `json:"-"`
}

// paymentMethodSetAsDefaultResponseEnvelopeJSON contains the JSON metadata for the
// struct [PaymentMethodSetAsDefaultResponseEnvelope]
type paymentMethodSetAsDefaultResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentMethodSetAsDefaultResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentMethodSetAsDefaultResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PaymentMethodSetAsDefaultResponseEnvelopeSuccess bool

const (
	PaymentMethodSetAsDefaultResponseEnvelopeSuccessTrue PaymentMethodSetAsDefaultResponseEnvelopeSuccess = true
)

func (r PaymentMethodSetAsDefaultResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PaymentMethodSetAsDefaultResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
