// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// BillingService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBillingService] method instead.
type BillingService struct {
	Options       []option.RequestOption
	Topup         *BillingTopupService
	SpendingLimit *BillingSpendingLimitService
}

// NewBillingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBillingService(opts ...option.RequestOption) (r *BillingService) {
	r = &BillingService{}
	r.Options = opts
	r.Topup = NewBillingTopupService(opts...)
	r.SpendingLimit = NewBillingSpendingLimitService(opts...)
	return
}

// Retrieve the current credit balance, payment method info, and top-up
// configuration.
func (r *BillingService) CreditBalance(ctx context.Context, query BillingCreditBalanceParams, opts ...option.RequestOption) (res *BillingCreditBalanceResponse, err error) {
	var env BillingCreditBalanceResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/billing/credit-balance", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve a list of past invoices with pagination, optionally filtered by type.
func (r *BillingService) InvoiceHistory(ctx context.Context, params BillingInvoiceHistoryParams, opts ...option.RequestOption) (res *BillingInvoiceHistoryResponse, err error) {
	var env BillingInvoiceHistoryResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/billing/invoice-history", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve a preview of the upcoming invoice including line items and tax.
func (r *BillingService) InvoicePreview(ctx context.Context, query BillingInvoicePreviewParams, opts ...option.RequestOption) (res *BillingInvoicePreviewResponse, err error) {
	var env BillingInvoicePreviewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/billing/invoice-preview", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve aggregated usage meter event summaries for the given time range.
func (r *BillingService) UsageHistory(ctx context.Context, params BillingUsageHistoryParams, opts ...option.RequestOption) (res *BillingUsageHistoryResponse, err error) {
	var env BillingUsageHistoryResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/billing/usage-history", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type BillingCreditBalanceResponse struct {
	Balance                 float64                                   `json:"balance" api:"required"`
	HasDefaultPaymentMethod bool                                      `json:"has_default_payment_method" api:"required"`
	PaymentMethod           BillingCreditBalanceResponsePaymentMethod `json:"payment_method" api:"required,nullable"`
	TopupConfig             BillingCreditBalanceResponseTopupConfig   `json:"topup_config" api:"required"`
	FirstTopupSuccess       bool                                      `json:"first_topup_success"`
	JSON                    billingCreditBalanceResponseJSON          `json:"-"`
}

// billingCreditBalanceResponseJSON contains the JSON metadata for the struct
// [BillingCreditBalanceResponse]
type billingCreditBalanceResponseJSON struct {
	Balance                 apijson.Field
	HasDefaultPaymentMethod apijson.Field
	PaymentMethod           apijson.Field
	TopupConfig             apijson.Field
	FirstTopupSuccess       apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *BillingCreditBalanceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingCreditBalanceResponseJSON) RawJSON() string {
	return r.raw
}

type BillingCreditBalanceResponsePaymentMethod struct {
	Brand string                                        `json:"brand"`
	Last4 string                                        `json:"last4"`
	JSON  billingCreditBalanceResponsePaymentMethodJSON `json:"-"`
}

// billingCreditBalanceResponsePaymentMethodJSON contains the JSON metadata for the
// struct [BillingCreditBalanceResponsePaymentMethod]
type billingCreditBalanceResponsePaymentMethodJSON struct {
	Brand       apijson.Field
	Last4       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingCreditBalanceResponsePaymentMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingCreditBalanceResponsePaymentMethodJSON) RawJSON() string {
	return r.raw
}

type BillingCreditBalanceResponseTopupConfig struct {
	Amount         float64                                     `json:"amount" api:"required,nullable"`
	DisabledReason string                                      `json:"disabledReason" api:"required,nullable"`
	Error          string                                      `json:"error" api:"required,nullable"`
	LastFailedAt   float64                                     `json:"lastFailedAt" api:"required,nullable"`
	Threshold      float64                                     `json:"threshold" api:"required,nullable"`
	JSON           billingCreditBalanceResponseTopupConfigJSON `json:"-"`
}

// billingCreditBalanceResponseTopupConfigJSON contains the JSON metadata for the
// struct [BillingCreditBalanceResponseTopupConfig]
type billingCreditBalanceResponseTopupConfigJSON struct {
	Amount         apijson.Field
	DisabledReason apijson.Field
	Error          apijson.Field
	LastFailedAt   apijson.Field
	Threshold      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BillingCreditBalanceResponseTopupConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingCreditBalanceResponseTopupConfigJSON) RawJSON() string {
	return r.raw
}

type BillingInvoiceHistoryResponse struct {
	Invoices   []BillingInvoiceHistoryResponseInvoice  `json:"invoices" api:"required"`
	Pagination BillingInvoiceHistoryResponsePagination `json:"pagination" api:"required"`
	JSON       billingInvoiceHistoryResponseJSON       `json:"-"`
}

// billingInvoiceHistoryResponseJSON contains the JSON metadata for the struct
// [BillingInvoiceHistoryResponse]
type billingInvoiceHistoryResponseJSON struct {
	Invoices    apijson.Field
	Pagination  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingInvoiceHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingInvoiceHistoryResponseJSON) RawJSON() string {
	return r.raw
}

type BillingInvoiceHistoryResponseInvoice struct {
	AmountDue       float64                                  `json:"amount_due" api:"required"`
	AmountPaid      float64                                  `json:"amount_paid" api:"required"`
	AmountRemaining float64                                  `json:"amount_remaining" api:"required"`
	Currency        string                                   `json:"currency" api:"required"`
	ID              string                                   `json:"id" api:"nullable"`
	AttemptCount    float64                                  `json:"attempt_count"`
	Attempted       bool                                     `json:"attempted"`
	AutoAdvance     bool                                     `json:"auto_advance" api:"nullable"`
	Created         float64                                  `json:"created"`
	CreatedBy       string                                   `json:"created_by"`
	Description     string                                   `json:"description" api:"nullable"`
	InvoiceOrigin   string                                   `json:"invoice_origin"`
	InvoicePDF      string                                   `json:"invoice_pdf" api:"nullable"`
	Status          string                                   `json:"status" api:"nullable"`
	JSON            billingInvoiceHistoryResponseInvoiceJSON `json:"-"`
}

// billingInvoiceHistoryResponseInvoiceJSON contains the JSON metadata for the
// struct [BillingInvoiceHistoryResponseInvoice]
type billingInvoiceHistoryResponseInvoiceJSON struct {
	AmountDue       apijson.Field
	AmountPaid      apijson.Field
	AmountRemaining apijson.Field
	Currency        apijson.Field
	ID              apijson.Field
	AttemptCount    apijson.Field
	Attempted       apijson.Field
	AutoAdvance     apijson.Field
	Created         apijson.Field
	CreatedBy       apijson.Field
	Description     apijson.Field
	InvoiceOrigin   apijson.Field
	InvoicePDF      apijson.Field
	Status          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BillingInvoiceHistoryResponseInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingInvoiceHistoryResponseInvoiceJSON) RawJSON() string {
	return r.raw
}

type BillingInvoiceHistoryResponsePagination struct {
	HasMore    bool                                        `json:"has_more" api:"required"`
	Page       float64                                     `json:"page" api:"required"`
	PerPage    float64                                     `json:"per_page" api:"required"`
	TotalCount float64                                     `json:"total_count" api:"required"`
	JSON       billingInvoiceHistoryResponsePaginationJSON `json:"-"`
}

// billingInvoiceHistoryResponsePaginationJSON contains the JSON metadata for the
// struct [BillingInvoiceHistoryResponsePagination]
type billingInvoiceHistoryResponsePaginationJSON struct {
	HasMore     apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingInvoiceHistoryResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingInvoiceHistoryResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type BillingInvoicePreviewResponse struct {
	ID              string                                     `json:"id" api:"required"`
	AmountDue       float64                                    `json:"amount_due" api:"required"`
	AmountPaid      float64                                    `json:"amount_paid" api:"required"`
	AmountRemaining float64                                    `json:"amount_remaining" api:"required"`
	Currency        string                                     `json:"currency" api:"required"`
	InvoiceLines    []BillingInvoicePreviewResponseInvoiceLine `json:"invoice_lines" api:"required"`
	PeriodEnd       float64                                    `json:"period_end" api:"required"`
	PeriodStart     float64                                    `json:"period_start" api:"required"`
	Status          BillingInvoicePreviewResponseStatus        `json:"status" api:"required"`
	JSON            billingInvoicePreviewResponseJSON          `json:"-"`
}

// billingInvoicePreviewResponseJSON contains the JSON metadata for the struct
// [BillingInvoicePreviewResponse]
type billingInvoicePreviewResponseJSON struct {
	ID              apijson.Field
	AmountDue       apijson.Field
	AmountPaid      apijson.Field
	AmountRemaining apijson.Field
	Currency        apijson.Field
	InvoiceLines    apijson.Field
	PeriodEnd       apijson.Field
	PeriodStart     apijson.Field
	Status          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BillingInvoicePreviewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingInvoicePreviewResponseJSON) RawJSON() string {
	return r.raw
}

type BillingInvoicePreviewResponseInvoiceLine struct {
	Amount              float64                                                       `json:"amount" api:"required"`
	Currency            string                                                        `json:"currency" api:"required"`
	Description         string                                                        `json:"description" api:"required,nullable"`
	Period              BillingInvoicePreviewResponseInvoiceLinesPeriod               `json:"period" api:"required"`
	Pricing             BillingInvoicePreviewResponseInvoiceLinesPricing              `json:"pricing" api:"required"`
	Quantity            float64                                                       `json:"quantity" api:"required"`
	PretaxCreditAmounts []BillingInvoicePreviewResponseInvoiceLinesPretaxCreditAmount `json:"pretax_credit_amounts"`
	JSON                billingInvoicePreviewResponseInvoiceLineJSON                  `json:"-"`
}

// billingInvoicePreviewResponseInvoiceLineJSON contains the JSON metadata for the
// struct [BillingInvoicePreviewResponseInvoiceLine]
type billingInvoicePreviewResponseInvoiceLineJSON struct {
	Amount              apijson.Field
	Currency            apijson.Field
	Description         apijson.Field
	Period              apijson.Field
	Pricing             apijson.Field
	Quantity            apijson.Field
	PretaxCreditAmounts apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *BillingInvoicePreviewResponseInvoiceLine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingInvoicePreviewResponseInvoiceLineJSON) RawJSON() string {
	return r.raw
}

type BillingInvoicePreviewResponseInvoiceLinesPeriod struct {
	End   float64                                             `json:"end" api:"required"`
	Start float64                                             `json:"start" api:"required"`
	JSON  billingInvoicePreviewResponseInvoiceLinesPeriodJSON `json:"-"`
}

// billingInvoicePreviewResponseInvoiceLinesPeriodJSON contains the JSON metadata
// for the struct [BillingInvoicePreviewResponseInvoiceLinesPeriod]
type billingInvoicePreviewResponseInvoiceLinesPeriodJSON struct {
	End         apijson.Field
	Start       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingInvoicePreviewResponseInvoiceLinesPeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingInvoicePreviewResponseInvoiceLinesPeriodJSON) RawJSON() string {
	return r.raw
}

type BillingInvoicePreviewResponseInvoiceLinesPricing struct {
	UnitAmountDecimal string                                               `json:"unit_amount_decimal" api:"required,nullable"`
	JSON              billingInvoicePreviewResponseInvoiceLinesPricingJSON `json:"-"`
}

// billingInvoicePreviewResponseInvoiceLinesPricingJSON contains the JSON metadata
// for the struct [BillingInvoicePreviewResponseInvoiceLinesPricing]
type billingInvoicePreviewResponseInvoiceLinesPricingJSON struct {
	UnitAmountDecimal apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *BillingInvoicePreviewResponseInvoiceLinesPricing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingInvoicePreviewResponseInvoiceLinesPricingJSON) RawJSON() string {
	return r.raw
}

type BillingInvoicePreviewResponseInvoiceLinesPretaxCreditAmount struct {
	Amount                   float64                                                         `json:"amount" api:"required"`
	Type                     string                                                          `json:"type" api:"required"`
	CreditBalanceTransaction string                                                          `json:"credit_balance_transaction" api:"nullable"`
	Discount                 string                                                          `json:"discount" api:"nullable"`
	JSON                     billingInvoicePreviewResponseInvoiceLinesPretaxCreditAmountJSON `json:"-"`
}

// billingInvoicePreviewResponseInvoiceLinesPretaxCreditAmountJSON contains the
// JSON metadata for the struct
// [BillingInvoicePreviewResponseInvoiceLinesPretaxCreditAmount]
type billingInvoicePreviewResponseInvoiceLinesPretaxCreditAmountJSON struct {
	Amount                   apijson.Field
	Type                     apijson.Field
	CreditBalanceTransaction apijson.Field
	Discount                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *BillingInvoicePreviewResponseInvoiceLinesPretaxCreditAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingInvoicePreviewResponseInvoiceLinesPretaxCreditAmountJSON) RawJSON() string {
	return r.raw
}

type BillingInvoicePreviewResponseStatus string

const (
	BillingInvoicePreviewResponseStatusDraft         BillingInvoicePreviewResponseStatus = "draft"
	BillingInvoicePreviewResponseStatusOpen          BillingInvoicePreviewResponseStatus = "open"
	BillingInvoicePreviewResponseStatusPaid          BillingInvoicePreviewResponseStatus = "paid"
	BillingInvoicePreviewResponseStatusUncollectible BillingInvoicePreviewResponseStatus = "uncollectible"
	BillingInvoicePreviewResponseStatusVoid          BillingInvoicePreviewResponseStatus = "void"
)

func (r BillingInvoicePreviewResponseStatus) IsKnown() bool {
	switch r {
	case BillingInvoicePreviewResponseStatusDraft, BillingInvoicePreviewResponseStatusOpen, BillingInvoicePreviewResponseStatusPaid, BillingInvoicePreviewResponseStatusUncollectible, BillingInvoicePreviewResponseStatusVoid:
		return true
	}
	return false
}

type BillingUsageHistoryResponse struct {
	History []BillingUsageHistoryResponseHistory `json:"history" api:"required"`
	JSON    billingUsageHistoryResponseJSON      `json:"-"`
}

// billingUsageHistoryResponseJSON contains the JSON metadata for the struct
// [BillingUsageHistoryResponse]
type billingUsageHistoryResponseJSON struct {
	History     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingUsageHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingUsageHistoryResponseJSON) RawJSON() string {
	return r.raw
}

type BillingUsageHistoryResponseHistory struct {
	ID              string                                 `json:"id" api:"required"`
	AggregatedValue float64                                `json:"aggregated_value" api:"required"`
	EndTime         float64                                `json:"end_time" api:"required"`
	StartTime       float64                                `json:"start_time" api:"required"`
	JSON            billingUsageHistoryResponseHistoryJSON `json:"-"`
}

// billingUsageHistoryResponseHistoryJSON contains the JSON metadata for the struct
// [BillingUsageHistoryResponseHistory]
type billingUsageHistoryResponseHistoryJSON struct {
	ID              apijson.Field
	AggregatedValue apijson.Field
	EndTime         apijson.Field
	StartTime       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BillingUsageHistoryResponseHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingUsageHistoryResponseHistoryJSON) RawJSON() string {
	return r.raw
}

type BillingCreditBalanceParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type BillingCreditBalanceResponseEnvelope struct {
	Errors     []BillingCreditBalanceResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages   []BillingCreditBalanceResponseEnvelopeMessages `json:"messages" api:"required"`
	Result     BillingCreditBalanceResponse                   `json:"result" api:"required"`
	Success    BillingCreditBalanceResponseEnvelopeSuccess    `json:"success" api:"required"`
	ResultInfo BillingCreditBalanceResponseEnvelopeResultInfo `json:"result_info"`
	JSON       billingCreditBalanceResponseEnvelopeJSON       `json:"-"`
}

// billingCreditBalanceResponseEnvelopeJSON contains the JSON metadata for the
// struct [BillingCreditBalanceResponseEnvelope]
type billingCreditBalanceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingCreditBalanceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingCreditBalanceResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BillingCreditBalanceResponseEnvelopeErrors struct {
	Code    float64                                        `json:"code" api:"required"`
	Message string                                         `json:"message" api:"required"`
	JSON    billingCreditBalanceResponseEnvelopeErrorsJSON `json:"-"`
}

// billingCreditBalanceResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [BillingCreditBalanceResponseEnvelopeErrors]
type billingCreditBalanceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingCreditBalanceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingCreditBalanceResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type BillingCreditBalanceResponseEnvelopeMessages struct {
	Code    float64                                          `json:"code" api:"required"`
	Message string                                           `json:"message" api:"required"`
	JSON    billingCreditBalanceResponseEnvelopeMessagesJSON `json:"-"`
}

// billingCreditBalanceResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [BillingCreditBalanceResponseEnvelopeMessages]
type billingCreditBalanceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingCreditBalanceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingCreditBalanceResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type BillingCreditBalanceResponseEnvelopeSuccess bool

const (
	BillingCreditBalanceResponseEnvelopeSuccessTrue BillingCreditBalanceResponseEnvelopeSuccess = true
)

func (r BillingCreditBalanceResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BillingCreditBalanceResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BillingCreditBalanceResponseEnvelopeResultInfo struct {
	HasMore    bool                                               `json:"has_more" api:"required"`
	Page       float64                                            `json:"page" api:"required"`
	PerPage    float64                                            `json:"per_page" api:"required"`
	TotalCount float64                                            `json:"total_count" api:"required"`
	JSON       billingCreditBalanceResponseEnvelopeResultInfoJSON `json:"-"`
}

// billingCreditBalanceResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [BillingCreditBalanceResponseEnvelopeResultInfo]
type billingCreditBalanceResponseEnvelopeResultInfoJSON struct {
	HasMore     apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingCreditBalanceResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingCreditBalanceResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type BillingInvoiceHistoryParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filter invoice type: auto, manual, or all.
	Type param.Field[BillingInvoiceHistoryParamsType] `query:"type"`
}

// URLQuery serializes [BillingInvoiceHistoryParams]'s query parameters as
// `url.Values`.
func (r BillingInvoiceHistoryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter invoice type: auto, manual, or all.
type BillingInvoiceHistoryParamsType string

const (
	BillingInvoiceHistoryParamsTypeAuto   BillingInvoiceHistoryParamsType = "auto"
	BillingInvoiceHistoryParamsTypeAll    BillingInvoiceHistoryParamsType = "all"
	BillingInvoiceHistoryParamsTypeManual BillingInvoiceHistoryParamsType = "manual"
)

func (r BillingInvoiceHistoryParamsType) IsKnown() bool {
	switch r {
	case BillingInvoiceHistoryParamsTypeAuto, BillingInvoiceHistoryParamsTypeAll, BillingInvoiceHistoryParamsTypeManual:
		return true
	}
	return false
}

type BillingInvoiceHistoryResponseEnvelope struct {
	Errors     []BillingInvoiceHistoryResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages   []BillingInvoiceHistoryResponseEnvelopeMessages `json:"messages" api:"required"`
	Result     BillingInvoiceHistoryResponse                   `json:"result" api:"required"`
	Success    BillingInvoiceHistoryResponseEnvelopeSuccess    `json:"success" api:"required"`
	ResultInfo BillingInvoiceHistoryResponseEnvelopeResultInfo `json:"result_info"`
	JSON       billingInvoiceHistoryResponseEnvelopeJSON       `json:"-"`
}

// billingInvoiceHistoryResponseEnvelopeJSON contains the JSON metadata for the
// struct [BillingInvoiceHistoryResponseEnvelope]
type billingInvoiceHistoryResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingInvoiceHistoryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingInvoiceHistoryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BillingInvoiceHistoryResponseEnvelopeErrors struct {
	Code    float64                                         `json:"code" api:"required"`
	Message string                                          `json:"message" api:"required"`
	JSON    billingInvoiceHistoryResponseEnvelopeErrorsJSON `json:"-"`
}

// billingInvoiceHistoryResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [BillingInvoiceHistoryResponseEnvelopeErrors]
type billingInvoiceHistoryResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingInvoiceHistoryResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingInvoiceHistoryResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type BillingInvoiceHistoryResponseEnvelopeMessages struct {
	Code    float64                                           `json:"code" api:"required"`
	Message string                                            `json:"message" api:"required"`
	JSON    billingInvoiceHistoryResponseEnvelopeMessagesJSON `json:"-"`
}

// billingInvoiceHistoryResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [BillingInvoiceHistoryResponseEnvelopeMessages]
type billingInvoiceHistoryResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingInvoiceHistoryResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingInvoiceHistoryResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type BillingInvoiceHistoryResponseEnvelopeSuccess bool

const (
	BillingInvoiceHistoryResponseEnvelopeSuccessTrue BillingInvoiceHistoryResponseEnvelopeSuccess = true
)

func (r BillingInvoiceHistoryResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BillingInvoiceHistoryResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BillingInvoiceHistoryResponseEnvelopeResultInfo struct {
	HasMore    bool                                                `json:"has_more" api:"required"`
	Page       float64                                             `json:"page" api:"required"`
	PerPage    float64                                             `json:"per_page" api:"required"`
	TotalCount float64                                             `json:"total_count" api:"required"`
	JSON       billingInvoiceHistoryResponseEnvelopeResultInfoJSON `json:"-"`
}

// billingInvoiceHistoryResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [BillingInvoiceHistoryResponseEnvelopeResultInfo]
type billingInvoiceHistoryResponseEnvelopeResultInfoJSON struct {
	HasMore     apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingInvoiceHistoryResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingInvoiceHistoryResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type BillingInvoicePreviewParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type BillingInvoicePreviewResponseEnvelope struct {
	Errors     []BillingInvoicePreviewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages   []BillingInvoicePreviewResponseEnvelopeMessages `json:"messages" api:"required"`
	Result     BillingInvoicePreviewResponse                   `json:"result" api:"required"`
	Success    BillingInvoicePreviewResponseEnvelopeSuccess    `json:"success" api:"required"`
	ResultInfo BillingInvoicePreviewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       billingInvoicePreviewResponseEnvelopeJSON       `json:"-"`
}

// billingInvoicePreviewResponseEnvelopeJSON contains the JSON metadata for the
// struct [BillingInvoicePreviewResponseEnvelope]
type billingInvoicePreviewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingInvoicePreviewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingInvoicePreviewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BillingInvoicePreviewResponseEnvelopeErrors struct {
	Code    float64                                         `json:"code" api:"required"`
	Message string                                          `json:"message" api:"required"`
	JSON    billingInvoicePreviewResponseEnvelopeErrorsJSON `json:"-"`
}

// billingInvoicePreviewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [BillingInvoicePreviewResponseEnvelopeErrors]
type billingInvoicePreviewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingInvoicePreviewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingInvoicePreviewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type BillingInvoicePreviewResponseEnvelopeMessages struct {
	Code    float64                                           `json:"code" api:"required"`
	Message string                                            `json:"message" api:"required"`
	JSON    billingInvoicePreviewResponseEnvelopeMessagesJSON `json:"-"`
}

// billingInvoicePreviewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [BillingInvoicePreviewResponseEnvelopeMessages]
type billingInvoicePreviewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingInvoicePreviewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingInvoicePreviewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type BillingInvoicePreviewResponseEnvelopeSuccess bool

const (
	BillingInvoicePreviewResponseEnvelopeSuccessTrue BillingInvoicePreviewResponseEnvelopeSuccess = true
)

func (r BillingInvoicePreviewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BillingInvoicePreviewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BillingInvoicePreviewResponseEnvelopeResultInfo struct {
	HasMore    bool                                                `json:"has_more" api:"required"`
	Page       float64                                             `json:"page" api:"required"`
	PerPage    float64                                             `json:"per_page" api:"required"`
	TotalCount float64                                             `json:"total_count" api:"required"`
	JSON       billingInvoicePreviewResponseEnvelopeResultInfoJSON `json:"-"`
}

// billingInvoicePreviewResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [BillingInvoicePreviewResponseEnvelopeResultInfo]
type billingInvoicePreviewResponseEnvelopeResultInfoJSON struct {
	HasMore     apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingInvoicePreviewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingInvoicePreviewResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type BillingUsageHistoryParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Grouping window for usage data.
	ValueGroupingWindow param.Field[BillingUsageHistoryParamsValueGroupingWindow] `query:"value_grouping_window" api:"required"`
	// End time as Unix timestamp in milliseconds.
	EndTime param.Field[float64] `query:"end_time"`
	// Start time as Unix timestamp in milliseconds.
	StartTime param.Field[float64] `query:"start_time"`
}

// URLQuery serializes [BillingUsageHistoryParams]'s query parameters as
// `url.Values`.
func (r BillingUsageHistoryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Grouping window for usage data.
type BillingUsageHistoryParamsValueGroupingWindow string

const (
	BillingUsageHistoryParamsValueGroupingWindowDay  BillingUsageHistoryParamsValueGroupingWindow = "day"
	BillingUsageHistoryParamsValueGroupingWindowHour BillingUsageHistoryParamsValueGroupingWindow = "hour"
)

func (r BillingUsageHistoryParamsValueGroupingWindow) IsKnown() bool {
	switch r {
	case BillingUsageHistoryParamsValueGroupingWindowDay, BillingUsageHistoryParamsValueGroupingWindowHour:
		return true
	}
	return false
}

type BillingUsageHistoryResponseEnvelope struct {
	Errors     []BillingUsageHistoryResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages   []BillingUsageHistoryResponseEnvelopeMessages `json:"messages" api:"required"`
	Result     BillingUsageHistoryResponse                   `json:"result" api:"required"`
	Success    BillingUsageHistoryResponseEnvelopeSuccess    `json:"success" api:"required"`
	ResultInfo BillingUsageHistoryResponseEnvelopeResultInfo `json:"result_info"`
	JSON       billingUsageHistoryResponseEnvelopeJSON       `json:"-"`
}

// billingUsageHistoryResponseEnvelopeJSON contains the JSON metadata for the
// struct [BillingUsageHistoryResponseEnvelope]
type billingUsageHistoryResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingUsageHistoryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingUsageHistoryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BillingUsageHistoryResponseEnvelopeErrors struct {
	Code    float64                                       `json:"code" api:"required"`
	Message string                                        `json:"message" api:"required"`
	JSON    billingUsageHistoryResponseEnvelopeErrorsJSON `json:"-"`
}

// billingUsageHistoryResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [BillingUsageHistoryResponseEnvelopeErrors]
type billingUsageHistoryResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingUsageHistoryResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingUsageHistoryResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type BillingUsageHistoryResponseEnvelopeMessages struct {
	Code    float64                                         `json:"code" api:"required"`
	Message string                                          `json:"message" api:"required"`
	JSON    billingUsageHistoryResponseEnvelopeMessagesJSON `json:"-"`
}

// billingUsageHistoryResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [BillingUsageHistoryResponseEnvelopeMessages]
type billingUsageHistoryResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingUsageHistoryResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingUsageHistoryResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type BillingUsageHistoryResponseEnvelopeSuccess bool

const (
	BillingUsageHistoryResponseEnvelopeSuccessTrue BillingUsageHistoryResponseEnvelopeSuccess = true
)

func (r BillingUsageHistoryResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BillingUsageHistoryResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BillingUsageHistoryResponseEnvelopeResultInfo struct {
	HasMore    bool                                              `json:"has_more" api:"required"`
	Page       float64                                           `json:"page" api:"required"`
	PerPage    float64                                           `json:"per_page" api:"required"`
	TotalCount float64                                           `json:"total_count" api:"required"`
	JSON       billingUsageHistoryResponseEnvelopeResultInfoJSON `json:"-"`
}

// billingUsageHistoryResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [BillingUsageHistoryResponseEnvelopeResultInfo]
type billingUsageHistoryResponseEnvelopeResultInfoJSON struct {
	HasMore     apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingUsageHistoryResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingUsageHistoryResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
