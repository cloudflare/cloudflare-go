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

// UnpaidInvoiceService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUnpaidInvoiceService] method instead.
type UnpaidInvoiceService struct {
	Options []option.RequestOption
}

// NewUnpaidInvoiceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUnpaidInvoiceService(opts ...option.RequestOption) (r *UnpaidInvoiceService) {
	r = &UnpaidInvoiceService{}
	r.Options = opts
	return
}

// Gets unpaid invoice information for an account.
func (r *UnpaidInvoiceService) Get(ctx context.Context, query UnpaidInvoiceGetParams, opts ...option.RequestOption) (res *UnpaidInvoiceGetResponse, err error) {
	var env UnpaidInvoiceGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/billing/unpaid-invoice", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type UnpaidInvoiceGetResponse struct {
	// List of unpaid invoices.
	Invoices []UnpaidInvoiceGetResponseInvoice `json:"invoices"`
	JSON     unpaidInvoiceGetResponseJSON      `json:"-"`
}

// unpaidInvoiceGetResponseJSON contains the JSON metadata for the struct
// [UnpaidInvoiceGetResponse]
type unpaidInvoiceGetResponseJSON struct {
	Invoices    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnpaidInvoiceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unpaidInvoiceGetResponseJSON) RawJSON() string {
	return r.raw
}

type UnpaidInvoiceGetResponseInvoice struct {
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
	Type string                              `json:"type"`
	JSON unpaidInvoiceGetResponseInvoiceJSON `json:"-"`
}

// unpaidInvoiceGetResponseInvoiceJSON contains the JSON metadata for the struct
// [UnpaidInvoiceGetResponseInvoice]
type unpaidInvoiceGetResponseInvoiceJSON struct {
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

func (r *UnpaidInvoiceGetResponseInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unpaidInvoiceGetResponseInvoiceJSON) RawJSON() string {
	return r.raw
}

type UnpaidInvoiceGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type UnpaidInvoiceGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo    `json:"errors" api:"required"`
	Messages []shared.ResponseInfo    `json:"messages" api:"required"`
	Result   UnpaidInvoiceGetResponse `json:"result" api:"required"`
	// Whether the API call was successful
	Success UnpaidInvoiceGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    unpaidInvoiceGetResponseEnvelopeJSON    `json:"-"`
}

// unpaidInvoiceGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [UnpaidInvoiceGetResponseEnvelope]
type unpaidInvoiceGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnpaidInvoiceGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unpaidInvoiceGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type UnpaidInvoiceGetResponseEnvelopeSuccess bool

const (
	UnpaidInvoiceGetResponseEnvelopeSuccessTrue UnpaidInvoiceGetResponseEnvelopeSuccess = true
)

func (r UnpaidInvoiceGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UnpaidInvoiceGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
