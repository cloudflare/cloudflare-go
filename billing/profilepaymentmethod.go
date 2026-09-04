// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package billing

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

// ProfilePaymentMethodService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProfilePaymentMethodService] method instead.
type ProfilePaymentMethodService struct {
	Options []option.RequestOption
}

// NewProfilePaymentMethodService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewProfilePaymentMethodService(opts ...option.RequestOption) (r *ProfilePaymentMethodService) {
	r = &ProfilePaymentMethodService{}
	r.Options = opts
	return
}

// Creates a Stripe payment intent for adding or updating a payment method on the
// account's billing profile. Returns a client secret for frontend payment method
// collection.
func (r *ProfilePaymentMethodService) New(ctx context.Context, body ProfilePaymentMethodNewParams, opts ...option.RequestOption) (res *ProfilePaymentMethodNewResponse, err error) {
	var env ProfilePaymentMethodNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/billing/profile/payment-method", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ProfilePaymentMethodNewResponse struct {
	// The Stripe client secret for frontend payment method collection.
	ClientSecret string `json:"client_secret"`
	// The type of Stripe intent created.
	IntentType string                              `json:"intent_type"`
	JSON       profilePaymentMethodNewResponseJSON `json:"-"`
}

// profilePaymentMethodNewResponseJSON contains the JSON metadata for the struct
// [ProfilePaymentMethodNewResponse]
type profilePaymentMethodNewResponseJSON struct {
	ClientSecret apijson.Field
	IntentType   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProfilePaymentMethodNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profilePaymentMethodNewResponseJSON) RawJSON() string {
	return r.raw
}

type ProfilePaymentMethodNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ProfilePaymentMethodNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo           `json:"errors" api:"required"`
	Messages []shared.ResponseInfo           `json:"messages" api:"required"`
	Result   ProfilePaymentMethodNewResponse `json:"result" api:"required"`
	// Whether the API call was successful
	Success ProfilePaymentMethodNewResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    profilePaymentMethodNewResponseEnvelopeJSON    `json:"-"`
}

// profilePaymentMethodNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ProfilePaymentMethodNewResponseEnvelope]
type profilePaymentMethodNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfilePaymentMethodNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profilePaymentMethodNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProfilePaymentMethodNewResponseEnvelopeSuccess bool

const (
	ProfilePaymentMethodNewResponseEnvelopeSuccessTrue ProfilePaymentMethodNewResponseEnvelopeSuccess = true
)

func (r ProfilePaymentMethodNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProfilePaymentMethodNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
