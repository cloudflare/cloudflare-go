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

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
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

// Returns cost and usage data for a single Cloudflare account, aligned with the
// [FinOps FOCUS v1.3](https://focus.finops.org/focus-specification/v1-3/) Cost and
// Usage dataset specification.
//
// Each record represents one billable metric for one account on one day. This
// includes all metered usage, including usage that falls within free-tier
// allowances and may result in zero cost.
//
// **Note:** Cost and pricing fields are not yet populated and will be absent from
// responses until billing integration is complete.
//
// When `from` and `to` are omitted, defaults to the start of the current month
// through today. The maximum date range is 31 days.
func (r *UsageService) Get(ctx context.Context, params UsageGetParams, opts ...option.RequestOption) (res *[]UsageGetResponse, err error) {
	var env UsageGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/billable/usage", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns billable usage data for PayGo (self-serve) accounts. When no query
// parameters are provided, returns usage for the current billing period. This
// endpoint is currently in alpha and access is restricted to select accounts.
// While in alpha, the endpoint may get breaking changes.
func (r *UsageService) Paygo(ctx context.Context, params UsagePaygoParams, opts ...option.RequestOption) (res *[]UsagePaygoResponse, err error) {
	var env UsagePaygoResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/paygo-usage", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// A single cost and usage record for a metered product within a specific charge
// period, aligned with the FinOps FOCUS v1.3 specification.
type UsageGetResponse struct {
	// Public identifier of the Cloudflare account (account tag).
	BillingAccountID string `json:"BillingAccountId" api:"required"`
	// Display name of the Cloudflare account.
	BillingAccountName string `json:"BillingAccountName" api:"required"`
	// Highest-level classification of a charge based on the nature of how it gets
	// billed. Currently only "Usage" is supported.
	ChargeCategory UsageGetResponseChargeCategory `json:"ChargeCategory" api:"required"`
	// Self-contained summary of the charge's purpose and price.
	ChargeDescription string `json:"ChargeDescription" api:"required"`
	// Indicates how often a charge occurs. Currently only "Usage-Based" is supported.
	ChargeFrequency UsageGetResponseChargeFrequency `json:"ChargeFrequency" api:"required"`
	// Exclusive end of the time interval during which the usage was consumed.
	ChargePeriodEnd time.Time `json:"ChargePeriodEnd" api:"required" format:"date-time"`
	// Inclusive start of the time interval during which the usage was consumed.
	ChargePeriodStart time.Time `json:"ChargePeriodStart" api:"required" format:"date-time"`
	// Measured usage amount within the charge period. Reflects raw metered consumption
	// before pricing transformations.
	ConsumedQuantity float64 `json:"ConsumedQuantity" api:"required"`
	// Unit of measure for the consumed quantity (e.g., "GB", "Requests",
	// "vCPU-Hours").
	ConsumedUnit string `json:"ConsumedUnit" api:"required"`
	// Name of the entity providing the underlying infrastructure or platform.
	HostProviderName string `json:"HostProviderName" api:"required"`
	// Name of the entity responsible for invoicing for the services consumed.
	InvoiceIssuerName string `json:"InvoiceIssuerName" api:"required"`
	// Name of the entity that made the services available for purchase.
	ServiceProviderName string `json:"ServiceProviderName" api:"required"`
	// The display name of the billable metric. Cloudflare extension; replaces FOCUS
	// SkuMeter.
	XBillableMetricName string `json:"x_BillableMetricName" api:"required"`
	// A charge serving as the basis for invoicing, inclusive of all reduced rates and
	// discounts while excluding the amortization of upfront charges (one-time or
	// recurring).
	BilledCost float64 `json:"BilledCost" api:"nullable"`
	// Currency that a charge was billed in (ISO 4217).
	BillingCurrency string `json:"BillingCurrency" api:"nullable"`
	// Exclusive end of the billing cycle that contains this usage record.
	BillingPeriodEnd time.Time `json:"BillingPeriodEnd" api:"nullable" format:"date-time"`
	// Inclusive start of the billing cycle that contains this usage record.
	BillingPeriodStart time.Time `json:"BillingPeriodStart" api:"nullable" format:"date-time"`
	// Indicates whether the row represents a correction to one or more charges
	// invoiced in a previous billing period.
	ChargeClass UsageGetResponseChargeClass `json:"ChargeClass" api:"nullable"`
	// Cost calculated by multiplying ContractedUnitPrice and the corresponding
	// PricingQuantity.
	ContractedCost float64 `json:"ContractedCost" api:"nullable"`
	// The agreed-upon unit price for a single PricingUnit of the associated billable
	// metric, inclusive of negotiated discounts, if present, while excluding any other
	// discounts.
	ContractedUnitPrice float64 `json:"ContractedUnitPrice" api:"nullable"`
	// The amortized cost of the charge after applying all reduced rates, discounts,
	// and the applicable portion of relevant, prepaid purchases (one-time or
	// recurring) that covered the charge.
	EffectiveCost float64 `json:"EffectiveCost" api:"nullable"`
	// Cost calculated by multiplying ListUnitPrice and the corresponding
	// PricingQuantity.
	ListCost float64 `json:"ListCost" api:"nullable"`
	// Suggested provider-published unit price for a single PricingUnit of the
	// associated billable metric, exclusive of any discounts.
	ListUnitPrice float64 `json:"ListUnitPrice" api:"nullable"`
	// Volume of a given service used or purchased, based on the PricingUnit.
	PricingQuantity float64 `json:"PricingQuantity" api:"nullable"`
	// Provider-specified measurement unit for determining unit prices, indicating how
	// the provider rates measured usage after applying pricing rules like block
	// pricing.
	PricingUnit string `json:"PricingUnit" api:"nullable"`
	// Provider-assigned identifier for an isolated geographic area where a service is
	// provided.
	RegionID string `json:"RegionId" api:"nullable"`
	// Name of an isolated geographic area where a service is provided.
	RegionName string `json:"RegionName" api:"nullable"`
	// Unique identifier assigned to a grouping of services. For Cloudflare, this is
	// the subscription or contract ID.
	SubAccountID string `json:"SubAccountId"`
	// Name assigned to a grouping of services. For Cloudflare, this is the
	// subscription or contract display name.
	SubAccountName string `json:"SubAccountName"`
	// The unique identifier for the billable metric in the Cloudflare catalog.
	// Cloudflare extension; replaces FOCUS SkuId.
	XBillableMetricID string `json:"x_BillableMetricId"`
	// The product family the charge belongs to (e.g., "R2", "Workers"). Cloudflare
	// extension; replaces FOCUS ServiceName.
	XProductFamilyName string `json:"x_ProductFamilyName"`
	// The identifier for the Cloudflare zone (zone tag). Cloudflare extension.
	XZoneID string `json:"x_ZoneId" api:"nullable"`
	// The display name of the Cloudflare zone. Cloudflare extension.
	XZoneName string               `json:"x_ZoneName" api:"nullable"`
	JSON      usageGetResponseJSON `json:"-"`
}

// usageGetResponseJSON contains the JSON metadata for the struct
// [UsageGetResponse]
type usageGetResponseJSON struct {
	BillingAccountID    apijson.Field
	BillingAccountName  apijson.Field
	ChargeCategory      apijson.Field
	ChargeDescription   apijson.Field
	ChargeFrequency     apijson.Field
	ChargePeriodEnd     apijson.Field
	ChargePeriodStart   apijson.Field
	ConsumedQuantity    apijson.Field
	ConsumedUnit        apijson.Field
	HostProviderName    apijson.Field
	InvoiceIssuerName   apijson.Field
	ServiceProviderName apijson.Field
	XBillableMetricName apijson.Field
	BilledCost          apijson.Field
	BillingCurrency     apijson.Field
	BillingPeriodEnd    apijson.Field
	BillingPeriodStart  apijson.Field
	ChargeClass         apijson.Field
	ContractedCost      apijson.Field
	ContractedUnitPrice apijson.Field
	EffectiveCost       apijson.Field
	ListCost            apijson.Field
	ListUnitPrice       apijson.Field
	PricingQuantity     apijson.Field
	PricingUnit         apijson.Field
	RegionID            apijson.Field
	RegionName          apijson.Field
	SubAccountID        apijson.Field
	SubAccountName      apijson.Field
	XBillableMetricID   apijson.Field
	XProductFamilyName  apijson.Field
	XZoneID             apijson.Field
	XZoneName           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *UsageGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetResponseJSON) RawJSON() string {
	return r.raw
}

// Highest-level classification of a charge based on the nature of how it gets
// billed. Currently only "Usage" is supported.
type UsageGetResponseChargeCategory string

const (
	UsageGetResponseChargeCategoryUsage UsageGetResponseChargeCategory = "Usage"
)

func (r UsageGetResponseChargeCategory) IsKnown() bool {
	switch r {
	case UsageGetResponseChargeCategoryUsage:
		return true
	}
	return false
}

// Indicates how often a charge occurs. Currently only "Usage-Based" is supported.
type UsageGetResponseChargeFrequency string

const (
	UsageGetResponseChargeFrequencyUsageBased UsageGetResponseChargeFrequency = "Usage-Based"
)

func (r UsageGetResponseChargeFrequency) IsKnown() bool {
	switch r {
	case UsageGetResponseChargeFrequencyUsageBased:
		return true
	}
	return false
}

// Indicates whether the row represents a correction to one or more charges
// invoiced in a previous billing period.
type UsageGetResponseChargeClass string

const (
	UsageGetResponseChargeClassCorrection UsageGetResponseChargeClass = "Correction"
)

func (r UsageGetResponseChargeClass) IsKnown() bool {
	switch r {
	case UsageGetResponseChargeClassCorrection:
		return true
	}
	return false
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
	// A display name for the unit of measurement used for the product (for example,
	// "GB-months", "GB-seconds"). May be empty when the unit is implicit in the
	// service name.
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
	ServiceName string `json:"ServiceName" api:"required"`
	// Identifies the product family for the Cloudflare service.
	ServiceFamilyName string                 `json:"ServiceFamilyName"`
	JSON              usagePaygoResponseJSON `json:"-"`
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
	ServiceFamilyName        apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *UsagePaygoResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usagePaygoResponseJSON) RawJSON() string {
	return r.raw
}

type UsageGetParams struct {
	// Represents a Cloudflare resource identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Start date for the usage query (ISO 8601). Required if `to` is set. When omitted
	// along with `to`, defaults to the start of the current month. Filters by charge
	// period (when consumption happened), not billing period. The maximum date range
	// is 31 days.
	From param.Field[time.Time] `query:"from" format:"date"`
	// Filter results by billable metric id (e.g., workers_standard_requests).
	Metric param.Field[string] `query:"metric"`
	// End date for the usage query (ISO 8601). Required if `from` is set. When omitted
	// along with `from`, defaults to today. Filters by charge period (when consumption
	// happened), not billing period. The maximum date range is 31 days.
	To param.Field[time.Time] `query:"to" format:"date"`
}

// URLQuery serializes [UsageGetParams]'s query parameters as `url.Values`.
func (r UsageGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Successful response containing an array of FOCUS-aligned cost and usage records.
type UsageGetResponseEnvelope struct {
	// Contains error details if the request failed.
	Errors []UsageGetResponseEnvelopeErrors `json:"errors" api:"required,nullable"`
	// Contains informational notices about the response.
	Messages []UsageGetResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	// Contains the array of cost and usage records.
	Result []UsageGetResponse `json:"result" api:"required"`
	// Indicates whether the API call was successful.
	Success UsageGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    usageGetResponseEnvelopeJSON    `json:"-"`
}

// usageGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [UsageGetResponseEnvelope]
type usageGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Represents an API notice or error detail.
type UsageGetResponseEnvelopeErrors struct {
	// Describes the error or notice.
	Message string `json:"message" api:"required"`
	// Identifies the error or notice type.
	Code int64                              `json:"code"`
	JSON usageGetResponseEnvelopeErrorsJSON `json:"-"`
}

// usageGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [UsageGetResponseEnvelopeErrors]
type usageGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Represents an API notice or error detail.
type UsageGetResponseEnvelopeMessages struct {
	// Describes the error or notice.
	Message string `json:"message" api:"required"`
	// Identifies the error or notice type.
	Code int64                                `json:"code"`
	JSON usageGetResponseEnvelopeMessagesJSON `json:"-"`
}

// usageGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [UsageGetResponseEnvelopeMessages]
type usageGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Indicates whether the API call was successful.
type UsageGetResponseEnvelopeSuccess bool

const (
	UsageGetResponseEnvelopeSuccessTrue UsageGetResponseEnvelopeSuccess = true
)

func (r UsageGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UsageGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type UsagePaygoParams struct {
	// Represents a Cloudflare resource identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Start date for the usage query (ISO 8601).
	From param.Field[time.Time] `query:"from" format:"date"`
	// End date for the usage query (ISO 8601).
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
