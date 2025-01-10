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

// TunnelManagementService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTunnelManagementService] method instead.
type TunnelManagementService struct {
	Options []option.RequestOption
}

// NewTunnelManagementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTunnelManagementService(opts ...option.RequestOption) (r *TunnelManagementService) {
	r = &TunnelManagementService{}
	r.Options = opts
	return
}

// Gets a management token used to access the management resources (i.e. Streaming
// Logs) of a tunnel.
func (r *TunnelManagementService) New(ctx context.Context, tunnelID string, params TunnelManagementNewParams, opts ...option.RequestOption) (res *string, err error) {
	var env TunnelManagementNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/management", params.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TunnelManagementNewParams struct {
	// Cloudflare account ID
	AccountID param.Field[string]                              `path:"account_id,required"`
	Resources param.Field[[]TunnelManagementNewParamsResource] `json:"resources,required"`
}

func (r TunnelManagementNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Management resources the token will have access to.
type TunnelManagementNewParamsResource string

const (
	TunnelManagementNewParamsResourceLogs TunnelManagementNewParamsResource = "logs"
)

func (r TunnelManagementNewParamsResource) IsKnown() bool {
	switch r {
	case TunnelManagementNewParamsResourceLogs:
		return true
	}
	return false
}

type TunnelManagementNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   string                `json:"result,required"`
	// Whether the API call was successful
	Success TunnelManagementNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelManagementNewResponseEnvelopeJSON    `json:"-"`
}

// tunnelManagementNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [TunnelManagementNewResponseEnvelope]
type tunnelManagementNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelManagementNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelManagementNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelManagementNewResponseEnvelopeSuccess bool

const (
	TunnelManagementNewResponseEnvelopeSuccessTrue TunnelManagementNewResponseEnvelopeSuccess = true
)

func (r TunnelManagementNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelManagementNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
