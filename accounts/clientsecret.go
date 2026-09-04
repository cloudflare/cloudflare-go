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

// ClientSecretService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientSecretService] method instead.
type ClientSecretService struct {
	Options []option.RequestOption
}

// NewClientSecretService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewClientSecretService(opts ...option.RequestOption) (r *ClientSecretService) {
	r = &ClientSecretService{}
	r.Options = opts
	return
}

// Creates a Stripe setup intent for adding a payment method to an account. Returns
// a client secret for frontend payment method collection.
func (r *ClientSecretService) New(ctx context.Context, body ClientSecretNewParams, opts ...option.RequestOption) (res *ClientSecretNewResponse, err error) {
	var env ClientSecretNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/client-secret", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ClientSecretNewResponse struct {
	// The Stripe client secret for frontend payment confirmation.
	ClientSecret string                      `json:"client_secret"`
	JSON         clientSecretNewResponseJSON `json:"-"`
}

// clientSecretNewResponseJSON contains the JSON metadata for the struct
// [ClientSecretNewResponse]
type clientSecretNewResponseJSON struct {
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ClientSecretNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientSecretNewResponseJSON) RawJSON() string {
	return r.raw
}

type ClientSecretNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ClientSecretNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo   `json:"errors" api:"required"`
	Messages []shared.ResponseInfo   `json:"messages" api:"required"`
	Result   ClientSecretNewResponse `json:"result" api:"required"`
	// Whether the API call was successful
	Success ClientSecretNewResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    clientSecretNewResponseEnvelopeJSON    `json:"-"`
}

// clientSecretNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ClientSecretNewResponseEnvelope]
type clientSecretNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientSecretNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientSecretNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ClientSecretNewResponseEnvelopeSuccess bool

const (
	ClientSecretNewResponseEnvelopeSuccessTrue ClientSecretNewResponseEnvelopeSuccess = true
)

func (r ClientSecretNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ClientSecretNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
