// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package billing

import (
	"context"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/shared"
)

// BillingService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBillingService] method instead.
type BillingService struct {
	Options       []option.RequestOption
	Profiles      *ProfileService
	Usage         *UsageService
	Credits       *CreditService
	History       *HistoryService
	BadDebt       *BadDebtService
	UnpaidInvoice *UnpaidInvoiceService
	RatePlans     *RatePlanService
}

// NewBillingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBillingService(opts ...option.RequestOption) (r *BillingService) {
	r = &BillingService{}
	r.Options = opts
	r.Profiles = NewProfileService(opts...)
	r.Usage = NewUsageService(opts...)
	r.Credits = NewCreditService(opts...)
	r.History = NewHistoryService(opts...)
	r.BadDebt = NewBadDebtService(opts...)
	r.UnpaidInvoice = NewUnpaidInvoiceService(opts...)
	r.RatePlans = NewRatePlanService(opts...)
	return
}

// Validates a billing address and returns validated address suggestions.
// Authentication is not enforced to support pre-signup address validation flows,
// so credentials are accepted but not required.
func (r *BillingService) AddressValidation(ctx context.Context, body BillingAddressValidationParams, opts ...option.RequestOption) (res *BillingAddressValidationResponse, err error) {
	var env BillingAddressValidationResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "billing/address-validation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type BillingAddressValidationResponse struct {
	// List of validated address suggestions.
	ValidatedAddresses []BillingAddressValidationResponseValidatedAddress `json:"validated_addresses"`
	JSON               billingAddressValidationResponseJSON               `json:"-"`
}

// billingAddressValidationResponseJSON contains the JSON metadata for the struct
// [BillingAddressValidationResponse]
type billingAddressValidationResponseJSON struct {
	ValidatedAddresses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *BillingAddressValidationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingAddressValidationResponseJSON) RawJSON() string {
	return r.raw
}

type BillingAddressValidationResponseValidatedAddress struct {
	// Validated address line 1.
	Address string `json:"address"`
	// Validated address line 2.
	Address2 string `json:"address2"`
	// Validated city.
	City string `json:"city"`
	// Validated country code.
	Country string `json:"country"`
	// Validated state or province.
	State string `json:"state"`
	// The validation result code.
	ValidationCode string `json:"validation_code"`
	// Validated postal or zip code.
	Zipcode string                                               `json:"zipcode"`
	JSON    billingAddressValidationResponseValidatedAddressJSON `json:"-"`
}

// billingAddressValidationResponseValidatedAddressJSON contains the JSON metadata
// for the struct [BillingAddressValidationResponseValidatedAddress]
type billingAddressValidationResponseValidatedAddressJSON struct {
	Address        apijson.Field
	Address2       apijson.Field
	City           apijson.Field
	Country        apijson.Field
	State          apijson.Field
	ValidationCode apijson.Field
	Zipcode        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BillingAddressValidationResponseValidatedAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingAddressValidationResponseValidatedAddressJSON) RawJSON() string {
	return r.raw
}

type BillingAddressValidationParams struct {
	// Address line 1.
	Address param.Field[string] `json:"address"`
	// Address line 2.
	Address2 param.Field[string] `json:"address2"`
	// City.
	City param.Field[string] `json:"city"`
	// Country code.
	Country param.Field[string] `json:"country"`
	// State or province.
	State param.Field[string] `json:"state"`
	// Postal or zip code.
	Zipcode param.Field[string] `json:"zipcode"`
}

func (r BillingAddressValidationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BillingAddressValidationResponseEnvelope struct {
	Errors   []shared.ResponseInfo            `json:"errors" api:"required"`
	Messages []shared.ResponseInfo            `json:"messages" api:"required"`
	Result   BillingAddressValidationResponse `json:"result" api:"required"`
	// Whether the API call was successful
	Success BillingAddressValidationResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    billingAddressValidationResponseEnvelopeJSON    `json:"-"`
}

// billingAddressValidationResponseEnvelopeJSON contains the JSON metadata for the
// struct [BillingAddressValidationResponseEnvelope]
type billingAddressValidationResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingAddressValidationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingAddressValidationResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type BillingAddressValidationResponseEnvelopeSuccess bool

const (
	BillingAddressValidationResponseEnvelopeSuccessTrue BillingAddressValidationResponseEnvelopeSuccess = true
)

func (r BillingAddressValidationResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BillingAddressValidationResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
