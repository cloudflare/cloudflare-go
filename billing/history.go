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
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// HistoryService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHistoryService] method instead.
type HistoryService struct {
	Options []option.RequestOption
}

// NewHistoryService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHistoryService(opts ...option.RequestOption) (r *HistoryService) {
	r = &HistoryService{}
	r.Options = opts
	return
}

// Gets the billing history for an account.
func (r *HistoryService) List(ctx context.Context, params HistoryListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[HistoryListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/billing/history", params.AccountID)
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

// Gets the billing history for an account.
func (r *HistoryService) ListAutoPaging(ctx context.Context, params HistoryListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[HistoryListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

type HistoryListResponse struct {
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
	Type string                  `json:"type"`
	JSON historyListResponseJSON `json:"-"`
}

// historyListResponseJSON contains the JSON metadata for the struct
// [HistoryListResponse]
type historyListResponseJSON struct {
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

func (r *HistoryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r historyListResponseJSON) RawJSON() string {
	return r.raw
}

type HistoryListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Filter billing history by status.
	Status param.Field[string] `query:"status"`
}

// URLQuery serializes [HistoryListParams]'s query parameters as `url.Values`.
func (r HistoryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
