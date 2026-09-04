// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// ReceiptService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReceiptService] method instead.
type ReceiptService struct {
	Options []option.RequestOption
}

// NewReceiptService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewReceiptService(opts ...option.RequestOption) (r *ReceiptService) {
	r = &ReceiptService{}
	r.Options = opts
	return
}

// Downloads a receipt as a PDF document.
func (r *ReceiptService) PDF(ctx context.Context, receiptID string, params ReceiptPDFParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/pdf")}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if receiptID == "" {
		err = errors.New("missing required receipt_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/receipts/%s/pdf", params.AccountID, receiptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type ReceiptPDFParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The document type to generate.
	Doctype param.Field[string] `query:"doctype"`
}

// URLQuery serializes [ReceiptPDFParams]'s query parameters as `url.Values`.
func (r ReceiptPDFParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
