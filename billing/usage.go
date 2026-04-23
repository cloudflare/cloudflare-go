// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package billing

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// UsageService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUsageService] method instead.
type UsageService struct {
	Options []option.RequestOption
}

// NewUsageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUsageService(opts ...option.RequestOption) (r *UsageService) {
	r = &UsageService{}
	r.Options = opts
	return
}

// Returns billable usage data for PayGo (self-serve) accounts. When no query
// parameters are provided, returns usage for the current billing period. This
// endpoint is currently in beta and access is restricted to select accounts.
func (r *UsageService) Paygo(ctx context.Context, params UsagePaygoParams, opts ...option.RequestOption) (res *[]UsagePaygoResponse, err error) {
	var env UsagePaygoResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/billing/usage/paygo", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Represents a single billable usage record.
type UsagePaygoResponse struct {
	// Specifies the billing currency code (ISO 4217).
	BillingCurrency string `json:"BillingCurrency" api:"required"`
	// Indicates the start of the billing period.
	BillingPeriodStart time.Time `json:"BillingPeriodStart" api:"required" format:"date-time"`
	// Indicates the end of the charge period.
	ChargePeriodEnd time.Time `json:"ChargePeriodEnd" api:"required" format:"date-time"`
	// Indicates the start of the charge period.
	ChargePeriodStart time.Time `json:"ChargePeriodStart" api:"required" format:"date-time"`
	// Specifies the quantity consumed during this charge period.
	ConsumedQuantity float64 `json:"ConsumedQuantity" api:"required"`
	// Specifies the unit of measurement for consumed quantity.
	ConsumedUnit string `json:"ConsumedUnit" api:"required"`
	// Specifies the cost for this charge period in the billing currency.
	ContractedCost float64 `json:"ContractedCost" api:"required"`
	// Specifies the cumulated cost for the billing period in the billing currency.
	CumulatedContractedCost float64 `json:"CumulatedContractedCost" api:"required"`
	// Specifies the cumulated pricing quantity for the billing period.
	CumulatedPricingQuantity int64 `json:"CumulatedPricingQuantity" api:"required"`
	// Specifies the pricing quantity for this charge period.
	PricingQuantity int64 `json:"PricingQuantity" api:"required"`
	// Identifies the Cloudflare service.
	ServiceName string                 `json:"ServiceName" api:"required"`
	JSON        usagePaygoResponseJSON `json:"-"`
}

// usagePaygoResponseJSON contains the JSON metadata for the struct
// [UsagePaygoResponse]
type usagePaygoResponseJSON struct {
	BillingCurrency          apijson.Field
	BillingPeriodStart       apijson.Field
	ChargePeriodEnd          apijson.Field
	ChargePeriodStart        apijson.Field
	ConsumedQuantity         apijson.Field
	ConsumedUnit             apijson.Field
	ContractedCost           apijson.Field
	CumulatedContractedCost  apijson.Field
	CumulatedPricingQuantity apijson.Field
	PricingQuantity          apijson.Field
	ServiceName              apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *UsagePaygoResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usagePaygoResponseJSON) RawJSON() string {
	return r.raw
}

type UsagePaygoParams struct {
	// Represents a Cloudflare resource identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Defines the start date for the usage query (e.g., 2025-02-01).
	From param.Field[time.Time] `query:"from" format:"date"`
	// Defines the end date for the usage query (e.g., 2025-03-01).
	To param.Field[time.Time] `query:"to" format:"date"`
}

// URLQuery serializes [UsagePaygoParams]'s query parameters as `url.Values`.
func (r UsagePaygoParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Represents a successful response containing billable usage records.
type UsagePaygoResponseEnvelope struct {
	// Contains error details if the request failed.
	Errors []UsagePaygoResponseEnvelopeErrors `json:"errors" api:"required,nullable"`
	// Contains informational notices about the response.
	Messages []UsagePaygoResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	// Contains the array of billable usage records.
	Result []UsagePaygoResponse `json:"result" api:"required"`
	// Indicates whether the API call was successful.
	Success UsagePaygoResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    usagePaygoResponseEnvelopeJSON    `json:"-"`
}

// usagePaygoResponseEnvelopeJSON contains the JSON metadata for the struct
// [UsagePaygoResponseEnvelope]
type usagePaygoResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsagePaygoResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usagePaygoResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Represents an API notice or error detail.
type UsagePaygoResponseEnvelopeErrors struct {
	// Describes the error or notice.
	Message string `json:"message" api:"required"`
	// Identifies the error or notice type.
	Code int64                                `json:"code"`
	JSON usagePaygoResponseEnvelopeErrorsJSON `json:"-"`
}

// usagePaygoResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [UsagePaygoResponseEnvelopeErrors]
type usagePaygoResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsagePaygoResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usagePaygoResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Represents an API notice or error detail.
type UsagePaygoResponseEnvelopeMessages struct {
	// Describes the error or notice.
	Message string `json:"message" api:"required"`
	// Identifies the error or notice type.
	Code int64                                  `json:"code"`
	JSON usagePaygoResponseEnvelopeMessagesJSON `json:"-"`
}

// usagePaygoResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [UsagePaygoResponseEnvelopeMessages]
type usagePaygoResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsagePaygoResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usagePaygoResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Indicates whether the API call was successful.
type UsagePaygoResponseEnvelopeSuccess bool

const (
	UsagePaygoResponseEnvelopeSuccessTrue UsagePaygoResponseEnvelopeSuccess = true
)

func (r UsagePaygoResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UsagePaygoResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
