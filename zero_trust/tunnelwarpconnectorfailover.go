// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/shared"
)

// TunnelWARPConnectorFailoverService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTunnelWARPConnectorFailoverService] method instead.
type TunnelWARPConnectorFailoverService struct {
	Options []option.RequestOption
}

// NewTunnelWARPConnectorFailoverService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewTunnelWARPConnectorFailoverService(opts ...option.RequestOption) (r *TunnelWARPConnectorFailoverService) {
	r = &TunnelWARPConnectorFailoverService{}
	r.Options = opts
	return
}

// Triggers a manual failover for a specific WARP Connector Tunnel, setting the
// specified client as the active connector. The tunnel must be configured for high
// availability (HA) and the client must be linked to the tunnel.
func (r *TunnelWARPConnectorFailoverService) Update(ctx context.Context, tunnelID string, params TunnelWARPConnectorFailoverUpdateParams, opts ...option.RequestOption) (res *TunnelWARPConnectorFailoverUpdateResponse, err error) {
	var env TunnelWARPConnectorFailoverUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/warp_connector/%s/failover", params.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type TunnelWARPConnectorFailoverUpdateResponse = interface{}

type TunnelWARPConnectorFailoverUpdateParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// UUID of the Cloudflare Tunnel connector.
	ClientID param.Field[string] `json:"client_id" api:"required" format:"uuid"`
}

func (r TunnelWARPConnectorFailoverUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TunnelWARPConnectorFailoverUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo                     `json:"errors" api:"required"`
	Messages []shared.ResponseInfo                     `json:"messages" api:"required"`
	Result   TunnelWARPConnectorFailoverUpdateResponse `json:"result" api:"required,nullable"`
	// Whether the API call was successful
	Success TunnelWARPConnectorFailoverUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    tunnelWARPConnectorFailoverUpdateResponseEnvelopeJSON    `json:"-"`
}

// tunnelWARPConnectorFailoverUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [TunnelWARPConnectorFailoverUpdateResponseEnvelope]
type tunnelWARPConnectorFailoverUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelWARPConnectorFailoverUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorFailoverUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelWARPConnectorFailoverUpdateResponseEnvelopeSuccess bool

const (
	TunnelWARPConnectorFailoverUpdateResponseEnvelopeSuccessTrue TunnelWARPConnectorFailoverUpdateResponseEnvelopeSuccess = true
)

func (r TunnelWARPConnectorFailoverUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelWARPConnectorFailoverUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
