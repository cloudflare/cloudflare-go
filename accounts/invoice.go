// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/shared"
)

// InvoiceService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvoiceService] method instead.
type InvoiceService struct {
	Options []option.RequestOption
}

// NewInvoiceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInvoiceService(opts ...option.RequestOption) (r *InvoiceService) {
	r = &InvoiceService{}
	r.Options = opts
	return
}

// Toggles PDF invoice generation for an account.
func (r *InvoiceService) Edit(ctx context.Context, params InvoiceEditParams, opts ...option.RequestOption) (res *interface{}, err error) {
	var env InvoiceEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/invoices", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type InvoiceEditParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Whether to enable or disable PDF invoice generation.
	Toggle param.Field[bool] `json:"toggle"`
}

func (r InvoiceEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	Result   interface{}           `json:"result" api:"required"`
	// Whether the API call was successful
	Success InvoiceEditResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    invoiceEditResponseEnvelopeJSON    `json:"-"`
}

// invoiceEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [InvoiceEditResponseEnvelope]
type invoiceEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type InvoiceEditResponseEnvelopeSuccess bool

const (
	InvoiceEditResponseEnvelopeSuccessTrue InvoiceEditResponseEnvelopeSuccess = true
)

func (r InvoiceEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case InvoiceEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
