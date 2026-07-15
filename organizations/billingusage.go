// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organizations

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

// BillingUsageService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBillingUsageService] method instead.
type BillingUsageService struct {
	Options []option.RequestOption
}

// NewBillingUsageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBillingUsageService(opts ...option.RequestOption) (r *BillingUsageService) {
	r = &BillingUsageService{}
	r.Options = opts
	return
}

// Returns cost and usage data for all accounts within an organization, aligned
// with the [FinOps FOCUS v1.3](https://focus.finops.org/focus-specification/v1-3/)
// Cost and Usage dataset specification.
//
// Each record represents one billable metric for one account on one day. This
// includes all metered usage, including usage that falls within free-tier
// allowances and may result in zero cost. The response includes usage for every
// account belonging to the specified organization.
//
// **Note:** Cost and pricing fields are not yet populated and will be absent from
// responses until billing integration is complete.
//
// When `from` and `to` are omitted, defaults to the start of the current month
// through today. The maximum date range is 31 days.
func (r *BillingUsageService) Get(ctx context.Context, organizationID string, query BillingUsageGetParams, opts ...option.RequestOption) (res *[]BillingUsageGetResponse, err error) {
	var env BillingUsageGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("organizations/%s/billable/usage", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// A single cost and usage record for a metered product within a specific charge
// period, aligned with the FinOps FOCUS v1.3 specification.
type BillingUsageGetResponse struct {
	// Public identifier of the Cloudflare account (account tag).
	BillingAccountID string `json:"BillingAccountId" api:"required"`
	// Display name of the Cloudflare account.
	BillingAccountName string `json:"BillingAccountName" api:"required"`
	// Highest-level classification of a charge based on the nature of how it gets
	// billed. Currently only "Usage" is supported.
	ChargeCategory BillingUsageGetResponseChargeCategory `json:"ChargeCategory" api:"required"`
	// Self-contained summary of the charge's purpose and price.
	ChargeDescription string `json:"ChargeDescription" api:"required"`
	// Indicates how often a charge occurs. Currently only "Usage-Based" is supported.
	ChargeFrequency BillingUsageGetResponseChargeFrequency `json:"ChargeFrequency" api:"required"`
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
	ChargeClass BillingUsageGetResponseChargeClass `json:"ChargeClass" api:"nullable"`
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
	XZoneName string                      `json:"x_ZoneName" api:"nullable"`
	JSON      billingUsageGetResponseJSON `json:"-"`
}

// billingUsageGetResponseJSON contains the JSON metadata for the struct
// [BillingUsageGetResponse]
type billingUsageGetResponseJSON struct {
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

func (r *BillingUsageGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingUsageGetResponseJSON) RawJSON() string {
	return r.raw
}

// Highest-level classification of a charge based on the nature of how it gets
// billed. Currently only "Usage" is supported.
type BillingUsageGetResponseChargeCategory string

const (
	BillingUsageGetResponseChargeCategoryUsage BillingUsageGetResponseChargeCategory = "Usage"
)

func (r BillingUsageGetResponseChargeCategory) IsKnown() bool {
	switch r {
	case BillingUsageGetResponseChargeCategoryUsage:
		return true
	}
	return false
}

// Indicates how often a charge occurs. Currently only "Usage-Based" is supported.
type BillingUsageGetResponseChargeFrequency string

const (
	BillingUsageGetResponseChargeFrequencyUsageBased BillingUsageGetResponseChargeFrequency = "Usage-Based"
)

func (r BillingUsageGetResponseChargeFrequency) IsKnown() bool {
	switch r {
	case BillingUsageGetResponseChargeFrequencyUsageBased:
		return true
	}
	return false
}

// Indicates whether the row represents a correction to one or more charges
// invoiced in a previous billing period.
type BillingUsageGetResponseChargeClass string

const (
	BillingUsageGetResponseChargeClassCorrection BillingUsageGetResponseChargeClass = "Correction"
)

func (r BillingUsageGetResponseChargeClass) IsKnown() bool {
	switch r {
	case BillingUsageGetResponseChargeClassCorrection:
		return true
	}
	return false
}

type BillingUsageGetParams struct {
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

// URLQuery serializes [BillingUsageGetParams]'s query parameters as `url.Values`.
func (r BillingUsageGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Successful response containing an array of FOCUS-aligned cost and usage records.
type BillingUsageGetResponseEnvelope struct {
	// Contains error details if the request failed.
	Errors []BillingUsageGetResponseEnvelopeErrors `json:"errors" api:"required,nullable"`
	// Contains informational notices about the response.
	Messages []BillingUsageGetResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	// Contains the array of cost and usage records.
	Result []BillingUsageGetResponse `json:"result" api:"required"`
	// Indicates whether the API call was successful.
	Success BillingUsageGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    billingUsageGetResponseEnvelopeJSON    `json:"-"`
}

// billingUsageGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [BillingUsageGetResponseEnvelope]
type billingUsageGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingUsageGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingUsageGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Represents an API notice or error detail.
type BillingUsageGetResponseEnvelopeErrors struct {
	// Describes the error or notice.
	Message string `json:"message" api:"required"`
	// Identifies the error or notice type.
	Code int64                                     `json:"code"`
	JSON billingUsageGetResponseEnvelopeErrorsJSON `json:"-"`
}

// billingUsageGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [BillingUsageGetResponseEnvelopeErrors]
type billingUsageGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingUsageGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingUsageGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Represents an API notice or error detail.
type BillingUsageGetResponseEnvelopeMessages struct {
	// Describes the error or notice.
	Message string `json:"message" api:"required"`
	// Identifies the error or notice type.
	Code int64                                       `json:"code"`
	JSON billingUsageGetResponseEnvelopeMessagesJSON `json:"-"`
}

// billingUsageGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [BillingUsageGetResponseEnvelopeMessages]
type billingUsageGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingUsageGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingUsageGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Indicates whether the API call was successful.
type BillingUsageGetResponseEnvelopeSuccess bool

const (
	BillingUsageGetResponseEnvelopeSuccessTrue BillingUsageGetResponseEnvelopeSuccess = true
)

func (r BillingUsageGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BillingUsageGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
