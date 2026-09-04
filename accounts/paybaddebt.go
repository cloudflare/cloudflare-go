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

// PayBadDebtService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPayBadDebtService] method instead.
type PayBadDebtService struct {
	Options []option.RequestOption
}

// NewPayBadDebtService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPayBadDebtService(opts ...option.RequestOption) (r *PayBadDebtService) {
	r = &PayBadDebtService{}
	r.Options = opts
	return
}

// Pays outstanding bad debt for an account. Discovers all debt automatically and
// handles invoice deduplication.
func (r *PayBadDebtService) New(ctx context.Context, params PayBadDebtNewParams, opts ...option.RequestOption) (res *PayBadDebtNewResponse, err error) {
	var env PayBadDebtNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/pay-bad-debt", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type PayBadDebtNewResponse struct {
	// The Stripe client secret for frontend payment confirmation.
	ClientSecret string                    `json:"client_secret"`
	JSON         payBadDebtNewResponseJSON `json:"-"`
}

// payBadDebtNewResponseJSON contains the JSON metadata for the struct
// [PayBadDebtNewResponse]
type payBadDebtNewResponseJSON struct {
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PayBadDebtNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payBadDebtNewResponseJSON) RawJSON() string {
	return r.raw
}

type PayBadDebtNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The payment method to use. If omitted, the default payment method is used.
	PaymentMethodID param.Field[string] `json:"payment_method_id"`
}

func (r PayBadDebtNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PayBadDebtNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	Result   PayBadDebtNewResponse `json:"result" api:"required"`
	// Whether the API call was successful
	Success PayBadDebtNewResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    payBadDebtNewResponseEnvelopeJSON    `json:"-"`
}

// payBadDebtNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [PayBadDebtNewResponseEnvelope]
type payBadDebtNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayBadDebtNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payBadDebtNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PayBadDebtNewResponseEnvelopeSuccess bool

const (
	PayBadDebtNewResponseEnvelopeSuccessTrue PayBadDebtNewResponseEnvelopeSuccess = true
)

func (r PayBadDebtNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PayBadDebtNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
