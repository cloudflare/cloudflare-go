// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package billing

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/shared"
	"github.com/tidwall/gjson"
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
//
// Deprecated: Use `get_account_usage_v2` instead.
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

// Returns high-level usage information for the account, including coverage, and
// subscription metadata.
func (r *UsageService) GetAccountUsageInfoV1(ctx context.Context, query UsageGetAccountUsageInfoV1Params, opts ...option.RequestOption) (res *UsageGetAccountUsageInfoV1Response, err error) {
	var env UsageGetAccountUsageInfoV1ResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/billable-usage/info", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns billable usage data for the account. When no query parameters are
// provided, returns usage for the current billing period.
func (r *UsageService) GetAccountUsageV1(ctx context.Context, params UsageGetAccountUsageV1Params, opts ...option.RequestOption) (res *[]UsageGetAccountUsageV1Response, err error) {
	var env UsageGetAccountUsageV1ResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/billable-usage", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
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
func (r *UsageService) GetAccountUsageV2(ctx context.Context, params UsageGetAccountUsageV2Params, opts ...option.RequestOption) (res *[]UsageGetAccountUsageV2Response, err error) {
	var env UsageGetAccountUsageV2ResponseEnvelope
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

// Returns billable usage data for the account. When no query parameters are
// provided, returns usage for the current billing period.
//
// Deprecated: Use `get_account_usage_v1` instead.
func (r *UsageService) Paygo(ctx context.Context, params UsagePaygoParams, opts ...option.RequestOption) (res *[]UsagePaygoResponse, err error) {
	var env UsagePaygoResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/billable-usage", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns high-level usage information for the account, including coverage, and
// subscription metadata.
//
// Deprecated: Use `get_account_usage_info_v1` instead.
func (r *UsageService) PaygoInfo(ctx context.Context, query UsagePaygoInfoParams, opts ...option.RequestOption) (res *UsagePaygoInfoResponse, err error) {
	var env UsagePaygoInfoResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/billable-usage/info", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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
	// The unique identifier for the billable metric in the Cloudflare catalog.
	// Cloudflare extension; replaces FOCUS SkuId.
	XBillableMetricID string `json:"x_BillableMetricId" api:"required"`
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
	// Tag values for the requested `GroupBy` keys. Omitted when `GroupBy` is not
	// provided. Missing keys are omitted, and key-only tags are returned as boolean
	// `true`. All other tag values are strings.
	Tags map[string]UsageGetResponseTagsUnion `json:"Tags"`
	// The product category the charge belongs to (e.g., "Developer", "Cloudflare
	// One"). Cloudflare extension; replaces FOCUS ServiceCategory.
	XProductCategoryName string `json:"x_ProductCategoryName"`
	// The unique identifier for the product family in the Cloudflare catalog.
	// Cloudflare extension; replaces FOCUS ServiceId.
	XProductFamilyID string `json:"x_ProductFamilyId"`
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
	BillingAccountID     apijson.Field
	BillingAccountName   apijson.Field
	ChargeCategory       apijson.Field
	ChargeDescription    apijson.Field
	ChargeFrequency      apijson.Field
	ChargePeriodEnd      apijson.Field
	ChargePeriodStart    apijson.Field
	ConsumedQuantity     apijson.Field
	ConsumedUnit         apijson.Field
	HostProviderName     apijson.Field
	InvoiceIssuerName    apijson.Field
	ServiceProviderName  apijson.Field
	XBillableMetricID    apijson.Field
	XBillableMetricName  apijson.Field
	BilledCost           apijson.Field
	BillingCurrency      apijson.Field
	BillingPeriodEnd     apijson.Field
	BillingPeriodStart   apijson.Field
	ChargeClass          apijson.Field
	ContractedCost       apijson.Field
	ContractedUnitPrice  apijson.Field
	EffectiveCost        apijson.Field
	ListCost             apijson.Field
	ListUnitPrice        apijson.Field
	PricingQuantity      apijson.Field
	PricingUnit          apijson.Field
	RegionID             apijson.Field
	RegionName           apijson.Field
	SubAccountID         apijson.Field
	SubAccountName       apijson.Field
	Tags                 apijson.Field
	XProductCategoryName apijson.Field
	XProductFamilyID     apijson.Field
	XProductFamilyName   apijson.Field
	XZoneID              apijson.Field
	XZoneName            apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
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

// Union satisfied by [shared.UnionString] or [UsageGetResponseTagsBoolean].
type UsageGetResponseTagsUnion interface {
	ImplementsUsageGetResponseTagsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UsageGetResponseTagsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(UsageGetResponseTagsBoolean(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(UsageGetResponseTagsBoolean(false)),
		},
	)
}

type UsageGetResponseTagsBoolean bool

const (
	UsageGetResponseTagsBooleanTrue UsageGetResponseTagsBoolean = true
)

func (r UsageGetResponseTagsBoolean) IsKnown() bool {
	switch r {
	case UsageGetResponseTagsBooleanTrue:
		return true
	}
	return false
}

func (r UsageGetResponseTagsBoolean) ImplementsUsageGetResponseTagsUnion() {}

// Contains the usage info.
type UsageGetAccountUsageInfoV1Response struct {
	// Indicates whether the account is covered.
	Covered bool `json:"covered" api:"required"`
	// List of subscriptions for the account.
	Subscriptions []UsageGetAccountUsageInfoV1ResponseSubscription `json:"subscriptions" api:"required"`
	JSON          usageGetAccountUsageInfoV1ResponseJSON           `json:"-"`
}

// usageGetAccountUsageInfoV1ResponseJSON contains the JSON metadata for the struct
// [UsageGetAccountUsageInfoV1Response]
type usageGetAccountUsageInfoV1ResponseJSON struct {
	Covered       apijson.Field
	Subscriptions apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UsageGetAccountUsageInfoV1Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetAccountUsageInfoV1ResponseJSON) RawJSON() string {
	return r.raw
}

type UsageGetAccountUsageInfoV1ResponseSubscription struct {
	// The identifier for the Cloudflare subscription.
	ID string `json:"id" api:"required"`
	// The subscription billing cycle anchor timestamp.
	BillingCycleAnchorTimestamp time.Time `json:"billing_cycle_anchor_timestamp" api:"required" format:"date-time"`
	// The subscription start timestamp.
	StartTimestamp time.Time `json:"start_timestamp" api:"required" format:"date-time"`
	// The subscription end timestamp. Omitted for active subscriptions; present only
	// when the subscription has been cancelled.
	EndTimestamp time.Time                                          `json:"end_timestamp" format:"date-time"`
	JSON         usageGetAccountUsageInfoV1ResponseSubscriptionJSON `json:"-"`
}

// usageGetAccountUsageInfoV1ResponseSubscriptionJSON contains the JSON metadata
// for the struct [UsageGetAccountUsageInfoV1ResponseSubscription]
type usageGetAccountUsageInfoV1ResponseSubscriptionJSON struct {
	ID                          apijson.Field
	BillingCycleAnchorTimestamp apijson.Field
	StartTimestamp              apijson.Field
	EndTimestamp                apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *UsageGetAccountUsageInfoV1ResponseSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetAccountUsageInfoV1ResponseSubscriptionJSON) RawJSON() string {
	return r.raw
}

// Represents a single billable usage record.
//
// This schema carries 19 of the 21 columns FOCUS 1.3 marks as mandatory. Mandatory
// columns are always present, using an explicit null when the value is unknown, so
// consumers can distinguish "unknown" from "not provided".
//
// Known gap 1: `ServiceCategory` (FOCUS 1.3 section 3.1.55) is not yet implemented
// and does not appear in this schema. Cloudflare's product catalog does not yet
// have a stakeholder-approved mapping from product family to a FOCUS
// ServiceCategory value, so the column is omitted entirely rather than shipping
// unapproved values. It will be added once that mapping exists.
//
// Known gap 2: `BillingPeriodEnd` (FOCUS 1.3 section 3.1.4) is not yet implemented
// and does not appear in this schema, because no authoritative source for it
// exists today. Deriving it by calendar arithmetic drifts for billing cycle
// anchors on day 29-31, and the billing provider's current period end describes
// only the current period, so it is wrong for backdated records.
// `BillingPeriodStart` is correctly sourced and is still reported, so records
// carry a billing period start with no corresponding end until an authoritative
// end date is available.
//
// Per FOCUS 1.3 section 4.1.4.1, the columns that are not part of FOCUS
// (`ServiceFamilyName`, `CumulatedPricingQuantity`, `CumulatedContractedCost`,
// `ZoneId`, `ZoneName` and `SubscriptionId`) would normally carry an `x_` prefix.
// They are kept unprefixed here to avoid a breaking change for existing consumers.
type UsageGetAccountUsageV1Response struct {
	// The amount invoiced for this charge. PayGo is billed directly by Cloudflare, so
	// this equals ContractedCost.
	BilledCost float64 `json:"BilledCost" api:"required"`
	// The identifier of the account the charge is billed to (account tag).
	BillingAccountID string `json:"BillingAccountId" api:"required"`
	// The display name of the billing account. Null when the name could not be
	// resolved.
	BillingAccountName string `json:"BillingAccountName" api:"required,nullable"`
	// Specifies the billing currency code (ISO 4217).
	BillingCurrency string `json:"BillingCurrency" api:"required"`
	// Indicates the start of the billing period. There is no `BillingPeriodEnd`
	// counterpart; see the known gaps described on this schema.
	BillingPeriodStart time.Time `json:"BillingPeriodStart" api:"required" format:"date-time"`
	// Describes the nature of the charge. Always "Usage" for this endpoint, which only
	// returns metered usage.
	ChargeCategory UsageGetAccountUsageV1ResponseChargeCategory `json:"ChargeCategory" api:"required"`
	// Indicates whether the row corrects a previously invoiced billing period. Always
	// null for this endpoint, which does not return corrections.
	ChargeClass string `json:"ChargeClass" api:"required,nullable"`
	// A human-readable summary of the charge.
	ChargeDescription string `json:"ChargeDescription" api:"required,nullable"`
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
	// Specifies the portion of usage that is actually subject to a unit price.
	CumulatedPricingQuantity int64 `json:"CumulatedPricingQuantity" api:"required"`
	// The amortized cost of the charge. PayGo has no upfront commitments, so this
	// equals ContractedCost.
	EffectiveCost float64 `json:"EffectiveCost" api:"required"`
	// The provider that hosts the infrastructure or platform the service runs on.
	HostProviderName string `json:"HostProviderName" api:"required"`
	// The entity that issues the invoice for this charge.
	InvoiceIssuerName string `json:"InvoiceIssuerName" api:"required"`
	// The cost at published list prices, before any discount. PayGo has no commitment
	// discounts, so this equals ContractedCost.
	ListCost float64 `json:"ListCost" api:"required"`
	// Specifies the pricing quantity for this charge period.
	PricingQuantity int64 `json:"PricingQuantity" api:"required"`
	// The unit that PricingQuantity is expressed in. Unlike ConsumedUnit this is never
	// empty; it falls back to "Count" when the service has no explicit unit.
	PricingUnit string `json:"PricingUnit" api:"required"`
	// Identifies the Cloudflare service.
	ServiceName string `json:"ServiceName" api:"required"`
	// The provider of the purchased service.
	ServiceProviderName string `json:"ServiceProviderName" api:"required"`
	// Identifies the product family for the Cloudflare service.
	ServiceFamilyName string `json:"ServiceFamilyName"`
	// The identifier for the Cloudflare subscription.
	SubscriptionID string `json:"SubscriptionId" api:"nullable"`
	// The identifier for the Cloudflare zone (zone tag).
	ZoneID string `json:"ZoneId" api:"nullable"`
	// The display name of the Cloudflare zone.
	ZoneName string                             `json:"ZoneName" api:"nullable"`
	JSON     usageGetAccountUsageV1ResponseJSON `json:"-"`
}

// usageGetAccountUsageV1ResponseJSON contains the JSON metadata for the struct
// [UsageGetAccountUsageV1Response]
type usageGetAccountUsageV1ResponseJSON struct {
	BilledCost               apijson.Field
	BillingAccountID         apijson.Field
	BillingAccountName       apijson.Field
	BillingCurrency          apijson.Field
	BillingPeriodStart       apijson.Field
	ChargeCategory           apijson.Field
	ChargeClass              apijson.Field
	ChargeDescription        apijson.Field
	ChargePeriodEnd          apijson.Field
	ChargePeriodStart        apijson.Field
	ConsumedQuantity         apijson.Field
	ConsumedUnit             apijson.Field
	ContractedCost           apijson.Field
	CumulatedContractedCost  apijson.Field
	CumulatedPricingQuantity apijson.Field
	EffectiveCost            apijson.Field
	HostProviderName         apijson.Field
	InvoiceIssuerName        apijson.Field
	ListCost                 apijson.Field
	PricingQuantity          apijson.Field
	PricingUnit              apijson.Field
	ServiceName              apijson.Field
	ServiceProviderName      apijson.Field
	ServiceFamilyName        apijson.Field
	SubscriptionID           apijson.Field
	ZoneID                   apijson.Field
	ZoneName                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *UsageGetAccountUsageV1Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetAccountUsageV1ResponseJSON) RawJSON() string {
	return r.raw
}

// Describes the nature of the charge. Always "Usage" for this endpoint, which only
// returns metered usage.
type UsageGetAccountUsageV1ResponseChargeCategory string

const (
	UsageGetAccountUsageV1ResponseChargeCategoryUsage UsageGetAccountUsageV1ResponseChargeCategory = "Usage"
)

func (r UsageGetAccountUsageV1ResponseChargeCategory) IsKnown() bool {
	switch r {
	case UsageGetAccountUsageV1ResponseChargeCategoryUsage:
		return true
	}
	return false
}

// A single cost and usage record for a metered product within a specific charge
// period, aligned with the FinOps FOCUS v1.3 specification.
type UsageGetAccountUsageV2Response struct {
	// Public identifier of the Cloudflare account (account tag).
	BillingAccountID string `json:"BillingAccountId" api:"required"`
	// Display name of the Cloudflare account.
	BillingAccountName string `json:"BillingAccountName" api:"required"`
	// Highest-level classification of a charge based on the nature of how it gets
	// billed. Currently only "Usage" is supported.
	ChargeCategory UsageGetAccountUsageV2ResponseChargeCategory `json:"ChargeCategory" api:"required"`
	// Self-contained summary of the charge's purpose and price.
	ChargeDescription string `json:"ChargeDescription" api:"required"`
	// Indicates how often a charge occurs. Currently only "Usage-Based" is supported.
	ChargeFrequency UsageGetAccountUsageV2ResponseChargeFrequency `json:"ChargeFrequency" api:"required"`
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
	// The unique identifier for the billable metric in the Cloudflare catalog.
	// Cloudflare extension; replaces FOCUS SkuId.
	XBillableMetricID string `json:"x_BillableMetricId" api:"required"`
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
	ChargeClass UsageGetAccountUsageV2ResponseChargeClass `json:"ChargeClass" api:"nullable"`
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
	// Tag values for the requested `GroupBy` keys. Omitted when `GroupBy` is not
	// provided. Missing keys are omitted, and key-only tags are returned as boolean
	// `true`. All other tag values are strings.
	Tags map[string]UsageGetAccountUsageV2ResponseTagsUnion `json:"Tags"`
	// The product category the charge belongs to (e.g., "Developer", "Cloudflare
	// One"). Cloudflare extension; replaces FOCUS ServiceCategory.
	XProductCategoryName string `json:"x_ProductCategoryName"`
	// The unique identifier for the product family in the Cloudflare catalog.
	// Cloudflare extension; replaces FOCUS ServiceId.
	XProductFamilyID string `json:"x_ProductFamilyId"`
	// The product family the charge belongs to (e.g., "R2", "Workers"). Cloudflare
	// extension; replaces FOCUS ServiceName.
	XProductFamilyName string `json:"x_ProductFamilyName"`
	// The identifier for the Cloudflare zone (zone tag). Cloudflare extension.
	XZoneID string `json:"x_ZoneId" api:"nullable"`
	// The display name of the Cloudflare zone. Cloudflare extension.
	XZoneName string                             `json:"x_ZoneName" api:"nullable"`
	JSON      usageGetAccountUsageV2ResponseJSON `json:"-"`
}

// usageGetAccountUsageV2ResponseJSON contains the JSON metadata for the struct
// [UsageGetAccountUsageV2Response]
type usageGetAccountUsageV2ResponseJSON struct {
	BillingAccountID     apijson.Field
	BillingAccountName   apijson.Field
	ChargeCategory       apijson.Field
	ChargeDescription    apijson.Field
	ChargeFrequency      apijson.Field
	ChargePeriodEnd      apijson.Field
	ChargePeriodStart    apijson.Field
	ConsumedQuantity     apijson.Field
	ConsumedUnit         apijson.Field
	HostProviderName     apijson.Field
	InvoiceIssuerName    apijson.Field
	ServiceProviderName  apijson.Field
	XBillableMetricID    apijson.Field
	XBillableMetricName  apijson.Field
	BilledCost           apijson.Field
	BillingCurrency      apijson.Field
	BillingPeriodEnd     apijson.Field
	BillingPeriodStart   apijson.Field
	ChargeClass          apijson.Field
	ContractedCost       apijson.Field
	ContractedUnitPrice  apijson.Field
	EffectiveCost        apijson.Field
	ListCost             apijson.Field
	ListUnitPrice        apijson.Field
	PricingQuantity      apijson.Field
	PricingUnit          apijson.Field
	RegionID             apijson.Field
	RegionName           apijson.Field
	SubAccountID         apijson.Field
	SubAccountName       apijson.Field
	Tags                 apijson.Field
	XProductCategoryName apijson.Field
	XProductFamilyID     apijson.Field
	XProductFamilyName   apijson.Field
	XZoneID              apijson.Field
	XZoneName            apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *UsageGetAccountUsageV2Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetAccountUsageV2ResponseJSON) RawJSON() string {
	return r.raw
}

// Highest-level classification of a charge based on the nature of how it gets
// billed. Currently only "Usage" is supported.
type UsageGetAccountUsageV2ResponseChargeCategory string

const (
	UsageGetAccountUsageV2ResponseChargeCategoryUsage UsageGetAccountUsageV2ResponseChargeCategory = "Usage"
)

func (r UsageGetAccountUsageV2ResponseChargeCategory) IsKnown() bool {
	switch r {
	case UsageGetAccountUsageV2ResponseChargeCategoryUsage:
		return true
	}
	return false
}

// Indicates how often a charge occurs. Currently only "Usage-Based" is supported.
type UsageGetAccountUsageV2ResponseChargeFrequency string

const (
	UsageGetAccountUsageV2ResponseChargeFrequencyUsageBased UsageGetAccountUsageV2ResponseChargeFrequency = "Usage-Based"
)

func (r UsageGetAccountUsageV2ResponseChargeFrequency) IsKnown() bool {
	switch r {
	case UsageGetAccountUsageV2ResponseChargeFrequencyUsageBased:
		return true
	}
	return false
}

// Indicates whether the row represents a correction to one or more charges
// invoiced in a previous billing period.
type UsageGetAccountUsageV2ResponseChargeClass string

const (
	UsageGetAccountUsageV2ResponseChargeClassCorrection UsageGetAccountUsageV2ResponseChargeClass = "Correction"
)

func (r UsageGetAccountUsageV2ResponseChargeClass) IsKnown() bool {
	switch r {
	case UsageGetAccountUsageV2ResponseChargeClassCorrection:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString] or
// [UsageGetAccountUsageV2ResponseTagsBoolean].
type UsageGetAccountUsageV2ResponseTagsUnion interface {
	ImplementsUsageGetAccountUsageV2ResponseTagsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UsageGetAccountUsageV2ResponseTagsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(UsageGetAccountUsageV2ResponseTagsBoolean(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(UsageGetAccountUsageV2ResponseTagsBoolean(false)),
		},
	)
}

type UsageGetAccountUsageV2ResponseTagsBoolean bool

const (
	UsageGetAccountUsageV2ResponseTagsBooleanTrue UsageGetAccountUsageV2ResponseTagsBoolean = true
)

func (r UsageGetAccountUsageV2ResponseTagsBoolean) IsKnown() bool {
	switch r {
	case UsageGetAccountUsageV2ResponseTagsBooleanTrue:
		return true
	}
	return false
}

func (r UsageGetAccountUsageV2ResponseTagsBoolean) ImplementsUsageGetAccountUsageV2ResponseTagsUnion() {
}

// Represents a single billable usage record.
//
// This schema carries 19 of the 21 columns FOCUS 1.3 marks as mandatory. Mandatory
// columns are always present, using an explicit null when the value is unknown, so
// consumers can distinguish "unknown" from "not provided".
//
// Known gap 1: `ServiceCategory` (FOCUS 1.3 section 3.1.55) is not yet implemented
// and does not appear in this schema. Cloudflare's product catalog does not yet
// have a stakeholder-approved mapping from product family to a FOCUS
// ServiceCategory value, so the column is omitted entirely rather than shipping
// unapproved values. It will be added once that mapping exists.
//
// Known gap 2: `BillingPeriodEnd` (FOCUS 1.3 section 3.1.4) is not yet implemented
// and does not appear in this schema, because no authoritative source for it
// exists today. Deriving it by calendar arithmetic drifts for billing cycle
// anchors on day 29-31, and the billing provider's current period end describes
// only the current period, so it is wrong for backdated records.
// `BillingPeriodStart` is correctly sourced and is still reported, so records
// carry a billing period start with no corresponding end until an authoritative
// end date is available.
//
// Per FOCUS 1.3 section 4.1.4.1, the columns that are not part of FOCUS
// (`ServiceFamilyName`, `CumulatedPricingQuantity`, `CumulatedContractedCost`,
// `ZoneId`, `ZoneName` and `SubscriptionId`) would normally carry an `x_` prefix.
// They are kept unprefixed here to avoid a breaking change for existing consumers.
type UsagePaygoResponse struct {
	// The amount invoiced for this charge. PayGo is billed directly by Cloudflare, so
	// this equals ContractedCost.
	BilledCost float64 `json:"BilledCost" api:"required"`
	// The identifier of the account the charge is billed to (account tag).
	BillingAccountID string `json:"BillingAccountId" api:"required"`
	// The display name of the billing account. Null when the name could not be
	// resolved.
	BillingAccountName string `json:"BillingAccountName" api:"required,nullable"`
	// Specifies the billing currency code (ISO 4217).
	BillingCurrency string `json:"BillingCurrency" api:"required"`
	// Indicates the start of the billing period. There is no `BillingPeriodEnd`
	// counterpart; see the known gaps described on this schema.
	BillingPeriodStart time.Time `json:"BillingPeriodStart" api:"required" format:"date-time"`
	// Describes the nature of the charge. Always "Usage" for this endpoint, which only
	// returns metered usage.
	ChargeCategory UsagePaygoResponseChargeCategory `json:"ChargeCategory" api:"required"`
	// Indicates whether the row corrects a previously invoiced billing period. Always
	// null for this endpoint, which does not return corrections.
	ChargeClass string `json:"ChargeClass" api:"required,nullable"`
	// A human-readable summary of the charge.
	ChargeDescription string `json:"ChargeDescription" api:"required,nullable"`
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
	// Specifies the portion of usage that is actually subject to a unit price.
	CumulatedPricingQuantity int64 `json:"CumulatedPricingQuantity" api:"required"`
	// The amortized cost of the charge. PayGo has no upfront commitments, so this
	// equals ContractedCost.
	EffectiveCost float64 `json:"EffectiveCost" api:"required"`
	// The provider that hosts the infrastructure or platform the service runs on.
	HostProviderName string `json:"HostProviderName" api:"required"`
	// The entity that issues the invoice for this charge.
	InvoiceIssuerName string `json:"InvoiceIssuerName" api:"required"`
	// The cost at published list prices, before any discount. PayGo has no commitment
	// discounts, so this equals ContractedCost.
	ListCost float64 `json:"ListCost" api:"required"`
	// Specifies the pricing quantity for this charge period.
	PricingQuantity int64 `json:"PricingQuantity" api:"required"`
	// The unit that PricingQuantity is expressed in. Unlike ConsumedUnit this is never
	// empty; it falls back to "Count" when the service has no explicit unit.
	PricingUnit string `json:"PricingUnit" api:"required"`
	// Identifies the Cloudflare service.
	ServiceName string `json:"ServiceName" api:"required"`
	// The provider of the purchased service.
	ServiceProviderName string `json:"ServiceProviderName" api:"required"`
	// Identifies the product family for the Cloudflare service.
	ServiceFamilyName string `json:"ServiceFamilyName"`
	// The identifier for the Cloudflare subscription.
	SubscriptionID string `json:"SubscriptionId" api:"nullable"`
	// The identifier for the Cloudflare zone (zone tag).
	ZoneID string `json:"ZoneId" api:"nullable"`
	// The display name of the Cloudflare zone.
	ZoneName string                 `json:"ZoneName" api:"nullable"`
	JSON     usagePaygoResponseJSON `json:"-"`
}

// usagePaygoResponseJSON contains the JSON metadata for the struct
// [UsagePaygoResponse]
type usagePaygoResponseJSON struct {
	BilledCost               apijson.Field
	BillingAccountID         apijson.Field
	BillingAccountName       apijson.Field
	BillingCurrency          apijson.Field
	BillingPeriodStart       apijson.Field
	ChargeCategory           apijson.Field
	ChargeClass              apijson.Field
	ChargeDescription        apijson.Field
	ChargePeriodEnd          apijson.Field
	ChargePeriodStart        apijson.Field
	ConsumedQuantity         apijson.Field
	ConsumedUnit             apijson.Field
	ContractedCost           apijson.Field
	CumulatedContractedCost  apijson.Field
	CumulatedPricingQuantity apijson.Field
	EffectiveCost            apijson.Field
	HostProviderName         apijson.Field
	InvoiceIssuerName        apijson.Field
	ListCost                 apijson.Field
	PricingQuantity          apijson.Field
	PricingUnit              apijson.Field
	ServiceName              apijson.Field
	ServiceProviderName      apijson.Field
	ServiceFamilyName        apijson.Field
	SubscriptionID           apijson.Field
	ZoneID                   apijson.Field
	ZoneName                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *UsagePaygoResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usagePaygoResponseJSON) RawJSON() string {
	return r.raw
}

// Describes the nature of the charge. Always "Usage" for this endpoint, which only
// returns metered usage.
type UsagePaygoResponseChargeCategory string

const (
	UsagePaygoResponseChargeCategoryUsage UsagePaygoResponseChargeCategory = "Usage"
)

func (r UsagePaygoResponseChargeCategory) IsKnown() bool {
	switch r {
	case UsagePaygoResponseChargeCategoryUsage:
		return true
	}
	return false
}

// Contains the usage info.
type UsagePaygoInfoResponse struct {
	// Indicates whether the account is covered.
	Covered bool `json:"covered" api:"required"`
	// List of subscriptions for the account.
	Subscriptions []UsagePaygoInfoResponseSubscription `json:"subscriptions" api:"required"`
	JSON          usagePaygoInfoResponseJSON           `json:"-"`
}

// usagePaygoInfoResponseJSON contains the JSON metadata for the struct
// [UsagePaygoInfoResponse]
type usagePaygoInfoResponseJSON struct {
	Covered       apijson.Field
	Subscriptions apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UsagePaygoInfoResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usagePaygoInfoResponseJSON) RawJSON() string {
	return r.raw
}

type UsagePaygoInfoResponseSubscription struct {
	// The identifier for the Cloudflare subscription.
	ID string `json:"id" api:"required"`
	// The subscription billing cycle anchor timestamp.
	BillingCycleAnchorTimestamp time.Time `json:"billing_cycle_anchor_timestamp" api:"required" format:"date-time"`
	// The subscription start timestamp.
	StartTimestamp time.Time `json:"start_timestamp" api:"required" format:"date-time"`
	// The subscription end timestamp. Omitted for active subscriptions; present only
	// when the subscription has been cancelled.
	EndTimestamp time.Time                              `json:"end_timestamp" format:"date-time"`
	JSON         usagePaygoInfoResponseSubscriptionJSON `json:"-"`
}

// usagePaygoInfoResponseSubscriptionJSON contains the JSON metadata for the struct
// [UsagePaygoInfoResponseSubscription]
type usagePaygoInfoResponseSubscriptionJSON struct {
	ID                          apijson.Field
	BillingCycleAnchorTimestamp apijson.Field
	StartTimestamp              apijson.Field
	EndTimestamp                apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *UsagePaygoInfoResponseSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usagePaygoInfoResponseSubscriptionJSON) RawJSON() string {
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

type UsageGetAccountUsageInfoV1Params struct {
	// Represents a Cloudflare resource identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

// Represents a successful response containing subscription info.
type UsageGetAccountUsageInfoV1ResponseEnvelope struct {
	// Contains error details if the request failed.
	Errors []UsageGetAccountUsageInfoV1ResponseEnvelopeErrors `json:"errors" api:"required,nullable"`
	// Contains any informational messages from the API.
	Messages []UsageGetAccountUsageInfoV1ResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	// Contains the usage info.
	Result UsageGetAccountUsageInfoV1Response `json:"result" api:"required"`
	// Indicates whether the API call was successful.
	Success UsageGetAccountUsageInfoV1ResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    usageGetAccountUsageInfoV1ResponseEnvelopeJSON    `json:"-"`
}

// usageGetAccountUsageInfoV1ResponseEnvelopeJSON contains the JSON metadata for
// the struct [UsageGetAccountUsageInfoV1ResponseEnvelope]
type usageGetAccountUsageInfoV1ResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageGetAccountUsageInfoV1ResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetAccountUsageInfoV1ResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Represents an API notice or error detail.
type UsageGetAccountUsageInfoV1ResponseEnvelopeErrors struct {
	// Describes the error or notice.
	Message string `json:"message" api:"required"`
	// Identifies the error or notice type.
	Code int64                                                `json:"code"`
	JSON usageGetAccountUsageInfoV1ResponseEnvelopeErrorsJSON `json:"-"`
}

// usageGetAccountUsageInfoV1ResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [UsageGetAccountUsageInfoV1ResponseEnvelopeErrors]
type usageGetAccountUsageInfoV1ResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageGetAccountUsageInfoV1ResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetAccountUsageInfoV1ResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Represents an API notice or error detail.
type UsageGetAccountUsageInfoV1ResponseEnvelopeMessages struct {
	// Describes the error or notice.
	Message string `json:"message" api:"required"`
	// Identifies the error or notice type.
	Code int64                                                  `json:"code"`
	JSON usageGetAccountUsageInfoV1ResponseEnvelopeMessagesJSON `json:"-"`
}

// usageGetAccountUsageInfoV1ResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [UsageGetAccountUsageInfoV1ResponseEnvelopeMessages]
type usageGetAccountUsageInfoV1ResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageGetAccountUsageInfoV1ResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetAccountUsageInfoV1ResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Indicates whether the API call was successful.
type UsageGetAccountUsageInfoV1ResponseEnvelopeSuccess bool

const (
	UsageGetAccountUsageInfoV1ResponseEnvelopeSuccessTrue UsageGetAccountUsageInfoV1ResponseEnvelopeSuccess = true
)

func (r UsageGetAccountUsageInfoV1ResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UsageGetAccountUsageInfoV1ResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type UsageGetAccountUsageV1Params struct {
	// Represents a Cloudflare resource identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Start date for the usage query (ISO 8601). The provided time range must include
	// the subscription billing cycle anchor day, otherwise no usage data is returned.
	// Use the info endpoint to retrieve the subscription anchor day.
	From param.Field[time.Time] `query:"from" format:"date"`
	// End date for the usage query (ISO 8601).
	To param.Field[time.Time] `query:"to" format:"date"`
}

// URLQuery serializes [UsageGetAccountUsageV1Params]'s query parameters as
// `url.Values`.
func (r UsageGetAccountUsageV1Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Represents a successful response containing billable usage records.
type UsageGetAccountUsageV1ResponseEnvelope struct {
	// Contains error details if the request failed.
	Errors []UsageGetAccountUsageV1ResponseEnvelopeErrors `json:"errors" api:"required,nullable"`
	// Contains informational notices about the response.
	Messages []UsageGetAccountUsageV1ResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	// Contains the array of billable usage records.
	Result []UsageGetAccountUsageV1Response `json:"result" api:"required"`
	// Indicates whether the API call was successful.
	Success UsageGetAccountUsageV1ResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    usageGetAccountUsageV1ResponseEnvelopeJSON    `json:"-"`
}

// usageGetAccountUsageV1ResponseEnvelopeJSON contains the JSON metadata for the
// struct [UsageGetAccountUsageV1ResponseEnvelope]
type usageGetAccountUsageV1ResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageGetAccountUsageV1ResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetAccountUsageV1ResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Represents an API notice or error detail.
type UsageGetAccountUsageV1ResponseEnvelopeErrors struct {
	// Describes the error or notice.
	Message string `json:"message" api:"required"`
	// Identifies the error or notice type.
	Code int64                                            `json:"code"`
	JSON usageGetAccountUsageV1ResponseEnvelopeErrorsJSON `json:"-"`
}

// usageGetAccountUsageV1ResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [UsageGetAccountUsageV1ResponseEnvelopeErrors]
type usageGetAccountUsageV1ResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageGetAccountUsageV1ResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetAccountUsageV1ResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Represents an API notice or error detail.
type UsageGetAccountUsageV1ResponseEnvelopeMessages struct {
	// Describes the error or notice.
	Message string `json:"message" api:"required"`
	// Identifies the error or notice type.
	Code int64                                              `json:"code"`
	JSON usageGetAccountUsageV1ResponseEnvelopeMessagesJSON `json:"-"`
}

// usageGetAccountUsageV1ResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [UsageGetAccountUsageV1ResponseEnvelopeMessages]
type usageGetAccountUsageV1ResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageGetAccountUsageV1ResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetAccountUsageV1ResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Indicates whether the API call was successful.
type UsageGetAccountUsageV1ResponseEnvelopeSuccess bool

const (
	UsageGetAccountUsageV1ResponseEnvelopeSuccessTrue UsageGetAccountUsageV1ResponseEnvelopeSuccess = true
)

func (r UsageGetAccountUsageV1ResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UsageGetAccountUsageV1ResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type UsageGetAccountUsageV2Params struct {
	// Represents a Cloudflare resource identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Start date for the usage query (ISO 8601). Required if `to` is set. When omitted
	// along with `to`, defaults to the start of the current month. Filters by charge
	// period (when consumption happened), not billing period. The maximum date range
	// is 31 days.
	From param.Field[time.Time] `query:"from" format:"date"`
	// End date for the usage query (ISO 8601). Required if `from` is set. When omitted
	// along with `from`, defaults to today. Filters by charge period (when consumption
	// happened), not billing period. The maximum date range is 31 days.
	To param.Field[time.Time] `query:"to" format:"date"`
}

// URLQuery serializes [UsageGetAccountUsageV2Params]'s query parameters as
// `url.Values`.
func (r UsageGetAccountUsageV2Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Successful response containing an array of FOCUS-aligned cost and usage records.
type UsageGetAccountUsageV2ResponseEnvelope struct {
	// Contains error details if the request failed.
	Errors []UsageGetAccountUsageV2ResponseEnvelopeErrors `json:"errors" api:"required,nullable"`
	// Contains informational notices about the response.
	Messages []UsageGetAccountUsageV2ResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	// Contains the array of cost and usage records.
	Result []UsageGetAccountUsageV2Response `json:"result" api:"required"`
	// Indicates whether the API call was successful.
	Success UsageGetAccountUsageV2ResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    usageGetAccountUsageV2ResponseEnvelopeJSON    `json:"-"`
}

// usageGetAccountUsageV2ResponseEnvelopeJSON contains the JSON metadata for the
// struct [UsageGetAccountUsageV2ResponseEnvelope]
type usageGetAccountUsageV2ResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageGetAccountUsageV2ResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetAccountUsageV2ResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Represents an API notice or error detail.
type UsageGetAccountUsageV2ResponseEnvelopeErrors struct {
	// Describes the error or notice.
	Message string `json:"message" api:"required"`
	// Identifies the error or notice type.
	Code int64                                            `json:"code"`
	JSON usageGetAccountUsageV2ResponseEnvelopeErrorsJSON `json:"-"`
}

// usageGetAccountUsageV2ResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [UsageGetAccountUsageV2ResponseEnvelopeErrors]
type usageGetAccountUsageV2ResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageGetAccountUsageV2ResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetAccountUsageV2ResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Represents an API notice or error detail.
type UsageGetAccountUsageV2ResponseEnvelopeMessages struct {
	// Describes the error or notice.
	Message string `json:"message" api:"required"`
	// Identifies the error or notice type.
	Code int64                                              `json:"code"`
	JSON usageGetAccountUsageV2ResponseEnvelopeMessagesJSON `json:"-"`
}

// usageGetAccountUsageV2ResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [UsageGetAccountUsageV2ResponseEnvelopeMessages]
type usageGetAccountUsageV2ResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageGetAccountUsageV2ResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetAccountUsageV2ResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Indicates whether the API call was successful.
type UsageGetAccountUsageV2ResponseEnvelopeSuccess bool

const (
	UsageGetAccountUsageV2ResponseEnvelopeSuccessTrue UsageGetAccountUsageV2ResponseEnvelopeSuccess = true
)

func (r UsageGetAccountUsageV2ResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UsageGetAccountUsageV2ResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type UsagePaygoParams struct {
	// Represents a Cloudflare resource identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Start date for the usage query (ISO 8601). The provided time range must include
	// the subscription billing cycle anchor day, otherwise no usage data is returned.
	// Use the info endpoint to retrieve the subscription anchor day.
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

type UsagePaygoInfoParams struct {
	// Represents a Cloudflare resource identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

// Represents a successful response containing subscription info.
type UsagePaygoInfoResponseEnvelope struct {
	// Contains error details if the request failed.
	Errors []UsagePaygoInfoResponseEnvelopeErrors `json:"errors" api:"required,nullable"`
	// Contains any informational messages from the API.
	Messages []UsagePaygoInfoResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	// Contains the usage info.
	Result UsagePaygoInfoResponse `json:"result" api:"required"`
	// Indicates whether the API call was successful.
	Success UsagePaygoInfoResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    usagePaygoInfoResponseEnvelopeJSON    `json:"-"`
}

// usagePaygoInfoResponseEnvelopeJSON contains the JSON metadata for the struct
// [UsagePaygoInfoResponseEnvelope]
type usagePaygoInfoResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsagePaygoInfoResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usagePaygoInfoResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Represents an API notice or error detail.
type UsagePaygoInfoResponseEnvelopeErrors struct {
	// Describes the error or notice.
	Message string `json:"message" api:"required"`
	// Identifies the error or notice type.
	Code int64                                    `json:"code"`
	JSON usagePaygoInfoResponseEnvelopeErrorsJSON `json:"-"`
}

// usagePaygoInfoResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [UsagePaygoInfoResponseEnvelopeErrors]
type usagePaygoInfoResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsagePaygoInfoResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usagePaygoInfoResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Represents an API notice or error detail.
type UsagePaygoInfoResponseEnvelopeMessages struct {
	// Describes the error or notice.
	Message string `json:"message" api:"required"`
	// Identifies the error or notice type.
	Code int64                                      `json:"code"`
	JSON usagePaygoInfoResponseEnvelopeMessagesJSON `json:"-"`
}

// usagePaygoInfoResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [UsagePaygoInfoResponseEnvelopeMessages]
type usagePaygoInfoResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsagePaygoInfoResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usagePaygoInfoResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Indicates whether the API call was successful.
type UsagePaygoInfoResponseEnvelopeSuccess bool

const (
	UsagePaygoInfoResponseEnvelopeSuccessTrue UsagePaygoInfoResponseEnvelopeSuccess = true
)

func (r UsagePaygoInfoResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UsagePaygoInfoResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
