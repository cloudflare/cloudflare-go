// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// TunnelTokenService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTunnelTokenService] method instead.
type TunnelTokenService struct {
	Options []option.RequestOption
}

// NewTunnelTokenService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTunnelTokenService(opts ...option.RequestOption) (r *TunnelTokenService) {
	r = &TunnelTokenService{}
	r.Options = opts
	return
}

// Gets the token used to associate cloudflared with a specific tunnel.
func (r *TunnelTokenService) Get(ctx context.Context, tunnelID string, query TunnelTokenGetParams, opts ...option.RequestOption) (res *string, err error) {
	var env TunnelTokenGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/token", query.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TunnelTokenGetParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type TunnelTokenGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   string                `json:"result,required"`
	// Whether the API call was successful
	Success TunnelTokenGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelTokenGetResponseEnvelopeJSON    `json:"-"`
}

// tunnelTokenGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [TunnelTokenGetResponseEnvelope]
type tunnelTokenGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelTokenGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelTokenGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelTokenGetResponseEnvelopeSuccess bool

const (
	TunnelTokenGetResponseEnvelopeSuccessTrue TunnelTokenGetResponseEnvelopeSuccess = true
)

func (r TunnelTokenGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelTokenGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
