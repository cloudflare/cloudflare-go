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

// BadDebtService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBadDebtService] method instead.
type BadDebtService struct {
	Options []option.RequestOption
}

// NewBadDebtService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBadDebtService(opts ...option.RequestOption) (r *BadDebtService) {
	r = &BadDebtService{}
	r.Options = opts
	return
}

// Gets bad debt information for an account, including outstanding invoices and
// total debt amount.
func (r *BadDebtService) Get(ctx context.Context, query BadDebtGetParams, opts ...option.RequestOption) (res *BadDebtGetResponse, err error) {
	var env BadDebtGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/billing/bad-debt", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type BadDebtGetResponse struct {
	// Amount already paid towards the debt.
	AlreadyPaid float64 `json:"already_paid"`
	// The current bad debt status of the account.
	BadDebtStatus string `json:"bad_debt_status"`
	// List of outstanding invoices contributing to bad debt.
	Invoices []BadDebtGetResponseInvoice `json:"invoices"`
	// Total outstanding debt amount.
	TotalDebtAmount float64                `json:"total_debt_amount"`
	JSON            badDebtGetResponseJSON `json:"-"`
}

// badDebtGetResponseJSON contains the JSON metadata for the struct
// [BadDebtGetResponse]
type badDebtGetResponseJSON struct {
	AlreadyPaid     apijson.Field
	BadDebtStatus   apijson.Field
	Invoices        apijson.Field
	TotalDebtAmount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BadDebtGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r badDebtGetResponseJSON) RawJSON() string {
	return r.raw
}

type BadDebtGetResponseInvoice struct {
	// Billing history item identifier.
	ID string `json:"id"`
	// The billing item action.
	Action string `json:"action"`
	// The amount associated with this billing item.
	Amount float64 `json:"amount"`
	// The amount remaining to pay.
	AmountToPay float64 `json:"amount_to_pay"`
	// The currency of the billing item.
	Currency string `json:"currency"`
	// The billing item description.
	Description string `json:"description"`
	// The external invoice identifier.
	ExternalInvoiceID string `json:"external_invoice_id"`
	// URL to the hosted invoice.
	HostedInvoiceURL string `json:"hosted_invoice_url"`
	// The associated invoice identifier.
	InvoiceID string `json:"invoice_id"`
	// When the billing event occurred.
	OccurredAt time.Time `json:"occurred_at" format:"date-time"`
	// The associated receipt identifier.
	ReceiptID string `json:"receipt_id"`
	// The source of the billing item.
	Source string `json:"source"`
	// The source invoice identifier.
	SourceInvoiceID string `json:"source_invoice_id"`
	// The status of the billing item.
	Status string `json:"status"`
	// The billing item type.
	Type string                        `json:"type"`
	JSON badDebtGetResponseInvoiceJSON `json:"-"`
}

// badDebtGetResponseInvoiceJSON contains the JSON metadata for the struct
// [BadDebtGetResponseInvoice]
type badDebtGetResponseInvoiceJSON struct {
	ID                apijson.Field
	Action            apijson.Field
	Amount            apijson.Field
	AmountToPay       apijson.Field
	Currency          apijson.Field
	Description       apijson.Field
	ExternalInvoiceID apijson.Field
	HostedInvoiceURL  apijson.Field
	InvoiceID         apijson.Field
	OccurredAt        apijson.Field
	ReceiptID         apijson.Field
	Source            apijson.Field
	SourceInvoiceID   apijson.Field
	Status            apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *BadDebtGetResponseInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r badDebtGetResponseInvoiceJSON) RawJSON() string {
	return r.raw
}

type BadDebtGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type BadDebtGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	Result   BadDebtGetResponse    `json:"result" api:"required"`
	// Whether the API call was successful
	Success BadDebtGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    badDebtGetResponseEnvelopeJSON    `json:"-"`
}

// badDebtGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [BadDebtGetResponseEnvelope]
type badDebtGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BadDebtGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r badDebtGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type BadDebtGetResponseEnvelopeSuccess bool

const (
	BadDebtGetResponseEnvelopeSuccessTrue BadDebtGetResponseEnvelopeSuccess = true
)

func (r BadDebtGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BadDebtGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
