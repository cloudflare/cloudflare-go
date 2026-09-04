// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package billing

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/shared"
)

// ProfileService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProfileService] method instead.
type ProfileService struct {
	Options       []option.RequestOption
	PaymentMethod *ProfilePaymentMethodService
}

// NewProfileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProfileService(opts ...option.RequestOption) (r *ProfileService) {
	r = &ProfileService{}
	r.Options = opts
	r.PaymentMethod = NewProfilePaymentMethodService(opts...)
	return
}

// Creates a billing profile for an account.
func (r *ProfileService) New(ctx context.Context, params ProfileNewParams, opts ...option.RequestOption) (res *ProfileNewResponse, err error) {
	var env ProfileNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/billing/profile", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates the billing profile for an account.
func (r *ProfileService) Update(ctx context.Context, params ProfileUpdateParams, opts ...option.RequestOption) (res *ProfileUpdateResponse, err error) {
	var env ProfileUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/billing/profile", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Deletes the billing profile for an account.
func (r *ProfileService) Delete(ctx context.Context, body ProfileDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return err
	}
	path := fmt.Sprintf("accounts/%s/billing/profile", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Gets the current billing profile for the account.
func (r *ProfileService) Get(ctx context.Context, query ProfileGetParams, opts ...option.RequestOption) (res *ProfileGetResponse, err error) {
	var env ProfileGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/billing/profile", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates the billing email addresses and preferred locale for an account.
func (r *ProfileService) UpdateBillingEmail(ctx context.Context, params ProfileUpdateBillingEmailParams, opts ...option.RequestOption) (res *ProfileUpdateBillingEmailResponse, err error) {
	var env ProfileUpdateBillingEmailResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/billing/profile", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ProfileNewResponse struct {
	// Billing item identifier tag.
	ID                     string                 `json:"id"`
	AccountType            string                 `json:"account_type"`
	Address                string                 `json:"address"`
	Address2               string                 `json:"address2"`
	Balance                string                 `json:"balance"`
	CardExpiryMonth        int64                  `json:"card_expiry_month"`
	CardExpiryYear         int64                  `json:"card_expiry_year"`
	CardNumber             string                 `json:"card_number"`
	City                   string                 `json:"city"`
	Company                string                 `json:"company"`
	Country                string                 `json:"country"`
	CreatedOn              time.Time              `json:"created_on" format:"date-time"`
	EditedOn               time.Time              `json:"edited_on" format:"date-time"`
	EnterpriseBillingEmail string                 `json:"enterprise_billing_email"`
	EnterprisePrimaryEmail string                 `json:"enterprise_primary_email"`
	FirstName              string                 `json:"first_name"`
	IsPartner              bool                   `json:"is_partner"`
	LastName               string                 `json:"last_name"`
	NextBillDate           time.Time              `json:"next_bill_date" format:"date-time"`
	PaymentAddress         string                 `json:"payment_address"`
	PaymentAddress2        string                 `json:"payment_address2"`
	PaymentCity            string                 `json:"payment_city"`
	PaymentCountry         string                 `json:"payment_country"`
	PaymentEmail           string                 `json:"payment_email"`
	PaymentFirstName       string                 `json:"payment_first_name"`
	PaymentLastName        string                 `json:"payment_last_name"`
	PaymentState           string                 `json:"payment_state"`
	PaymentZipcode         string                 `json:"payment_zipcode"`
	PrimaryEmail           string                 `json:"primary_email"`
	State                  string                 `json:"state"`
	TaxIDType              string                 `json:"tax_id_type"`
	Telephone              string                 `json:"telephone"`
	ValidationCode         string                 `json:"validation_code"`
	Vat                    string                 `json:"vat"`
	Zipcode                string                 `json:"zipcode"`
	JSON                   profileNewResponseJSON `json:"-"`
}

// profileNewResponseJSON contains the JSON metadata for the struct
// [ProfileNewResponse]
type profileNewResponseJSON struct {
	ID                     apijson.Field
	AccountType            apijson.Field
	Address                apijson.Field
	Address2               apijson.Field
	Balance                apijson.Field
	CardExpiryMonth        apijson.Field
	CardExpiryYear         apijson.Field
	CardNumber             apijson.Field
	City                   apijson.Field
	Company                apijson.Field
	Country                apijson.Field
	CreatedOn              apijson.Field
	EditedOn               apijson.Field
	EnterpriseBillingEmail apijson.Field
	EnterprisePrimaryEmail apijson.Field
	FirstName              apijson.Field
	IsPartner              apijson.Field
	LastName               apijson.Field
	NextBillDate           apijson.Field
	PaymentAddress         apijson.Field
	PaymentAddress2        apijson.Field
	PaymentCity            apijson.Field
	PaymentCountry         apijson.Field
	PaymentEmail           apijson.Field
	PaymentFirstName       apijson.Field
	PaymentLastName        apijson.Field
	PaymentState           apijson.Field
	PaymentZipcode         apijson.Field
	PrimaryEmail           apijson.Field
	State                  apijson.Field
	TaxIDType              apijson.Field
	Telephone              apijson.Field
	ValidationCode         apijson.Field
	Vat                    apijson.Field
	Zipcode                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ProfileNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileNewResponseJSON) RawJSON() string {
	return r.raw
}

type ProfileUpdateResponse struct {
	// Billing item identifier tag.
	ID                     string                    `json:"id"`
	AccountType            string                    `json:"account_type"`
	Address                string                    `json:"address"`
	Address2               string                    `json:"address2"`
	Balance                string                    `json:"balance"`
	CardExpiryMonth        int64                     `json:"card_expiry_month"`
	CardExpiryYear         int64                     `json:"card_expiry_year"`
	CardNumber             string                    `json:"card_number"`
	City                   string                    `json:"city"`
	Company                string                    `json:"company"`
	Country                string                    `json:"country"`
	CreatedOn              time.Time                 `json:"created_on" format:"date-time"`
	EditedOn               time.Time                 `json:"edited_on" format:"date-time"`
	EnterpriseBillingEmail string                    `json:"enterprise_billing_email"`
	EnterprisePrimaryEmail string                    `json:"enterprise_primary_email"`
	FirstName              string                    `json:"first_name"`
	IsPartner              bool                      `json:"is_partner"`
	LastName               string                    `json:"last_name"`
	NextBillDate           time.Time                 `json:"next_bill_date" format:"date-time"`
	PaymentAddress         string                    `json:"payment_address"`
	PaymentAddress2        string                    `json:"payment_address2"`
	PaymentCity            string                    `json:"payment_city"`
	PaymentCountry         string                    `json:"payment_country"`
	PaymentEmail           string                    `json:"payment_email"`
	PaymentFirstName       string                    `json:"payment_first_name"`
	PaymentLastName        string                    `json:"payment_last_name"`
	PaymentState           string                    `json:"payment_state"`
	PaymentZipcode         string                    `json:"payment_zipcode"`
	PrimaryEmail           string                    `json:"primary_email"`
	State                  string                    `json:"state"`
	TaxIDType              string                    `json:"tax_id_type"`
	Telephone              string                    `json:"telephone"`
	ValidationCode         string                    `json:"validation_code"`
	Vat                    string                    `json:"vat"`
	Zipcode                string                    `json:"zipcode"`
	JSON                   profileUpdateResponseJSON `json:"-"`
}

// profileUpdateResponseJSON contains the JSON metadata for the struct
// [ProfileUpdateResponse]
type profileUpdateResponseJSON struct {
	ID                     apijson.Field
	AccountType            apijson.Field
	Address                apijson.Field
	Address2               apijson.Field
	Balance                apijson.Field
	CardExpiryMonth        apijson.Field
	CardExpiryYear         apijson.Field
	CardNumber             apijson.Field
	City                   apijson.Field
	Company                apijson.Field
	Country                apijson.Field
	CreatedOn              apijson.Field
	EditedOn               apijson.Field
	EnterpriseBillingEmail apijson.Field
	EnterprisePrimaryEmail apijson.Field
	FirstName              apijson.Field
	IsPartner              apijson.Field
	LastName               apijson.Field
	NextBillDate           apijson.Field
	PaymentAddress         apijson.Field
	PaymentAddress2        apijson.Field
	PaymentCity            apijson.Field
	PaymentCountry         apijson.Field
	PaymentEmail           apijson.Field
	PaymentFirstName       apijson.Field
	PaymentLastName        apijson.Field
	PaymentState           apijson.Field
	PaymentZipcode         apijson.Field
	PrimaryEmail           apijson.Field
	State                  apijson.Field
	TaxIDType              apijson.Field
	Telephone              apijson.Field
	ValidationCode         apijson.Field
	Vat                    apijson.Field
	Zipcode                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ProfileUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ProfileGetResponse struct {
	// Billing item identifier tag.
	ID                     string                 `json:"id"`
	AccountType            string                 `json:"account_type"`
	Address                string                 `json:"address"`
	Address2               string                 `json:"address2"`
	Balance                string                 `json:"balance"`
	CardExpiryMonth        int64                  `json:"card_expiry_month"`
	CardExpiryYear         int64                  `json:"card_expiry_year"`
	CardNumber             string                 `json:"card_number"`
	City                   string                 `json:"city"`
	Company                string                 `json:"company"`
	Country                string                 `json:"country"`
	CreatedOn              time.Time              `json:"created_on" format:"date-time"`
	EditedOn               time.Time              `json:"edited_on" format:"date-time"`
	EnterpriseBillingEmail string                 `json:"enterprise_billing_email"`
	EnterprisePrimaryEmail string                 `json:"enterprise_primary_email"`
	FirstName              string                 `json:"first_name"`
	IsPartner              bool                   `json:"is_partner"`
	LastName               string                 `json:"last_name"`
	NextBillDate           time.Time              `json:"next_bill_date" format:"date-time"`
	PaymentAddress         string                 `json:"payment_address"`
	PaymentAddress2        string                 `json:"payment_address2"`
	PaymentCity            string                 `json:"payment_city"`
	PaymentCountry         string                 `json:"payment_country"`
	PaymentEmail           string                 `json:"payment_email"`
	PaymentFirstName       string                 `json:"payment_first_name"`
	PaymentLastName        string                 `json:"payment_last_name"`
	PaymentState           string                 `json:"payment_state"`
	PaymentZipcode         string                 `json:"payment_zipcode"`
	PrimaryEmail           string                 `json:"primary_email"`
	State                  string                 `json:"state"`
	TaxIDType              string                 `json:"tax_id_type"`
	Telephone              string                 `json:"telephone"`
	ValidationCode         string                 `json:"validation_code"`
	Vat                    string                 `json:"vat"`
	Zipcode                string                 `json:"zipcode"`
	JSON                   profileGetResponseJSON `json:"-"`
}

// profileGetResponseJSON contains the JSON metadata for the struct
// [ProfileGetResponse]
type profileGetResponseJSON struct {
	ID                     apijson.Field
	AccountType            apijson.Field
	Address                apijson.Field
	Address2               apijson.Field
	Balance                apijson.Field
	CardExpiryMonth        apijson.Field
	CardExpiryYear         apijson.Field
	CardNumber             apijson.Field
	City                   apijson.Field
	Company                apijson.Field
	Country                apijson.Field
	CreatedOn              apijson.Field
	EditedOn               apijson.Field
	EnterpriseBillingEmail apijson.Field
	EnterprisePrimaryEmail apijson.Field
	FirstName              apijson.Field
	IsPartner              apijson.Field
	LastName               apijson.Field
	NextBillDate           apijson.Field
	PaymentAddress         apijson.Field
	PaymentAddress2        apijson.Field
	PaymentCity            apijson.Field
	PaymentCountry         apijson.Field
	PaymentEmail           apijson.Field
	PaymentFirstName       apijson.Field
	PaymentLastName        apijson.Field
	PaymentState           apijson.Field
	PaymentZipcode         apijson.Field
	PrimaryEmail           apijson.Field
	State                  apijson.Field
	TaxIDType              apijson.Field
	Telephone              apijson.Field
	ValidationCode         apijson.Field
	Vat                    apijson.Field
	Zipcode                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ProfileGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileGetResponseJSON) RawJSON() string {
	return r.raw
}

type ProfileUpdateBillingEmailResponse struct {
	// Billing item identifier tag.
	ID                     string                                `json:"id"`
	AccountType            string                                `json:"account_type"`
	Address                string                                `json:"address"`
	Address2               string                                `json:"address2"`
	Balance                string                                `json:"balance"`
	CardExpiryMonth        int64                                 `json:"card_expiry_month"`
	CardExpiryYear         int64                                 `json:"card_expiry_year"`
	CardNumber             string                                `json:"card_number"`
	City                   string                                `json:"city"`
	Company                string                                `json:"company"`
	Country                string                                `json:"country"`
	CreatedOn              time.Time                             `json:"created_on" format:"date-time"`
	EditedOn               time.Time                             `json:"edited_on" format:"date-time"`
	EnterpriseBillingEmail string                                `json:"enterprise_billing_email"`
	EnterprisePrimaryEmail string                                `json:"enterprise_primary_email"`
	FirstName              string                                `json:"first_name"`
	IsPartner              bool                                  `json:"is_partner"`
	LastName               string                                `json:"last_name"`
	NextBillDate           time.Time                             `json:"next_bill_date" format:"date-time"`
	PaymentAddress         string                                `json:"payment_address"`
	PaymentAddress2        string                                `json:"payment_address2"`
	PaymentCity            string                                `json:"payment_city"`
	PaymentCountry         string                                `json:"payment_country"`
	PaymentEmail           string                                `json:"payment_email"`
	PaymentFirstName       string                                `json:"payment_first_name"`
	PaymentLastName        string                                `json:"payment_last_name"`
	PaymentState           string                                `json:"payment_state"`
	PaymentZipcode         string                                `json:"payment_zipcode"`
	PrimaryEmail           string                                `json:"primary_email"`
	State                  string                                `json:"state"`
	TaxIDType              string                                `json:"tax_id_type"`
	Telephone              string                                `json:"telephone"`
	ValidationCode         string                                `json:"validation_code"`
	Vat                    string                                `json:"vat"`
	Zipcode                string                                `json:"zipcode"`
	JSON                   profileUpdateBillingEmailResponseJSON `json:"-"`
}

// profileUpdateBillingEmailResponseJSON contains the JSON metadata for the struct
// [ProfileUpdateBillingEmailResponse]
type profileUpdateBillingEmailResponseJSON struct {
	ID                     apijson.Field
	AccountType            apijson.Field
	Address                apijson.Field
	Address2               apijson.Field
	Balance                apijson.Field
	CardExpiryMonth        apijson.Field
	CardExpiryYear         apijson.Field
	CardNumber             apijson.Field
	City                   apijson.Field
	Company                apijson.Field
	Country                apijson.Field
	CreatedOn              apijson.Field
	EditedOn               apijson.Field
	EnterpriseBillingEmail apijson.Field
	EnterprisePrimaryEmail apijson.Field
	FirstName              apijson.Field
	IsPartner              apijson.Field
	LastName               apijson.Field
	NextBillDate           apijson.Field
	PaymentAddress         apijson.Field
	PaymentAddress2        apijson.Field
	PaymentCity            apijson.Field
	PaymentCountry         apijson.Field
	PaymentEmail           apijson.Field
	PaymentFirstName       apijson.Field
	PaymentLastName        apijson.Field
	PaymentState           apijson.Field
	PaymentZipcode         apijson.Field
	PrimaryEmail           apijson.Field
	State                  apijson.Field
	TaxIDType              apijson.Field
	Telephone              apijson.Field
	ValidationCode         apijson.Field
	Vat                    apijson.Field
	Zipcode                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ProfileUpdateBillingEmailResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileUpdateBillingEmailResponseJSON) RawJSON() string {
	return r.raw
}

type ProfileNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Street address line 1.
	Address param.Field[string] `json:"address"`
	// Street address line 2 (apt, suite, etc.).
	Address2 param.Field[string] `json:"address2"`
	// Primary billing email address.
	BillingEmail param.Field[string] `json:"billing_email"`
	// Rate plan being purchased right after profile setup.
	BuyingRatePlan param.Field[string] `json:"buying_rate_plan"`
	// Captcha challenge JWT issued during onboarding.
	CaptchaChallengeJWT param.Field[string] `json:"captcha_challenge_jwt"`
	// Cloudflare Turnstile response.
	CfTurnstileResponse param.Field[string] `json:"cf_turnstile_response"`
	// City on the billing profile.
	City param.Field[string] `json:"city"`
	// Company name on the billing profile.
	Company param.Field[string] `json:"company"`
	// ISO 3166-1 alpha-2 country code.
	Country param.Field[string] `json:"country"`
	// First name on the billing profile.
	FirstName param.Field[string] `json:"first_name"`
	// hCaptcha response.
	HCaptchaResponse param.Field[string] `json:"h_captcha_response"`
	// Last name on the billing profile.
	LastName param.Field[string] `json:"last_name"`
	// Preferred locale for invoice rendering (BCP 47).
	PreferredLocale param.Field[string] `json:"preferred_locale"`
	// Secondary billing email address for CC on invoices.
	SecondaryBillingEmail param.Field[string] `json:"secondary_billing_email"`
	// State or region on the billing profile.
	State param.Field[string] `json:"state"`
	// Type of tax ID provided.
	TaxIDType param.Field[string] `json:"tax_id_type"`
	// Contact phone number.
	Telephone param.Field[string] `json:"telephone"`
	// VAT identifier.
	Vat param.Field[string] `json:"vat"`
	// ZIP or postal code.
	Zipcode param.Field[string] `json:"zipcode"`
}

func (r ProfileNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProfileNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	Result   ProfileNewResponse    `json:"result" api:"required"`
	// Whether the API call was successful
	Success ProfileNewResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    profileNewResponseEnvelopeJSON    `json:"-"`
}

// profileNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ProfileNewResponseEnvelope]
type profileNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProfileNewResponseEnvelopeSuccess bool

const (
	ProfileNewResponseEnvelopeSuccessTrue ProfileNewResponseEnvelopeSuccess = true
)

func (r ProfileNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProfileNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProfileUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Street address line 1.
	Address param.Field[string] `json:"address"`
	// Street address line 2 (apt, suite, etc.).
	Address2 param.Field[string] `json:"address2"`
	// Primary billing email address.
	BillingEmail param.Field[string] `json:"billing_email"`
	// Rate plan being purchased right after profile setup.
	BuyingRatePlan param.Field[string] `json:"buying_rate_plan"`
	// Captcha challenge JWT issued during onboarding.
	CaptchaChallengeJWT param.Field[string] `json:"captcha_challenge_jwt"`
	// Cloudflare Turnstile response.
	CfTurnstileResponse param.Field[string] `json:"cf_turnstile_response"`
	// City on the billing profile.
	City param.Field[string] `json:"city"`
	// Company name on the billing profile.
	Company param.Field[string] `json:"company"`
	// ISO 3166-1 alpha-2 country code.
	Country param.Field[string] `json:"country"`
	// First name on the billing profile.
	FirstName param.Field[string] `json:"first_name"`
	// hCaptcha response.
	HCaptchaResponse param.Field[string] `json:"h_captcha_response"`
	// Last name on the billing profile.
	LastName param.Field[string] `json:"last_name"`
	// Preferred locale for invoice rendering (BCP 47).
	PreferredLocale param.Field[string] `json:"preferred_locale"`
	// Secondary billing email address for CC on invoices.
	SecondaryBillingEmail param.Field[string] `json:"secondary_billing_email"`
	// State or region on the billing profile.
	State param.Field[string] `json:"state"`
	// Type of tax ID provided.
	TaxIDType param.Field[string] `json:"tax_id_type"`
	// Contact phone number.
	Telephone param.Field[string] `json:"telephone"`
	// VAT identifier.
	Vat param.Field[string] `json:"vat"`
	// ZIP or postal code.
	Zipcode param.Field[string] `json:"zipcode"`
}

func (r ProfileUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProfileUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	Result   ProfileUpdateResponse `json:"result" api:"required"`
	// Whether the API call was successful
	Success ProfileUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    profileUpdateResponseEnvelopeJSON    `json:"-"`
}

// profileUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [ProfileUpdateResponseEnvelope]
type profileUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProfileUpdateResponseEnvelopeSuccess bool

const (
	ProfileUpdateResponseEnvelopeSuccessTrue ProfileUpdateResponseEnvelopeSuccess = true
)

func (r ProfileUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProfileUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProfileDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ProfileGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ProfileGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	Result   ProfileGetResponse    `json:"result" api:"required"`
	// Whether the API call was successful
	Success ProfileGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    profileGetResponseEnvelopeJSON    `json:"-"`
}

// profileGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ProfileGetResponseEnvelope]
type profileGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProfileGetResponseEnvelopeSuccess bool

const (
	ProfileGetResponseEnvelopeSuccessTrue ProfileGetResponseEnvelopeSuccess = true
)

func (r ProfileGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProfileGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProfileUpdateBillingEmailParams struct {
	// Identifier
	AccountID             param.Field[string] `path:"account_id" api:"required"`
	BillingEmail          param.Field[string] `json:"billing_email"`
	PreferredLocale       param.Field[string] `json:"preferred_locale"`
	SecondaryBillingEmail param.Field[string] `json:"secondary_billing_email"`
}

func (r ProfileUpdateBillingEmailParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProfileUpdateBillingEmailResponseEnvelope struct {
	Errors   []shared.ResponseInfo             `json:"errors" api:"required"`
	Messages []shared.ResponseInfo             `json:"messages" api:"required"`
	Result   ProfileUpdateBillingEmailResponse `json:"result" api:"required"`
	// Whether the API call was successful
	Success ProfileUpdateBillingEmailResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    profileUpdateBillingEmailResponseEnvelopeJSON    `json:"-"`
}

// profileUpdateBillingEmailResponseEnvelopeJSON contains the JSON metadata for the
// struct [ProfileUpdateBillingEmailResponseEnvelope]
type profileUpdateBillingEmailResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileUpdateBillingEmailResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileUpdateBillingEmailResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProfileUpdateBillingEmailResponseEnvelopeSuccess bool

const (
	ProfileUpdateBillingEmailResponseEnvelopeSuccessTrue ProfileUpdateBillingEmailResponseEnvelopeSuccess = true
)

func (r ProfileUpdateBillingEmailResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProfileUpdateBillingEmailResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
