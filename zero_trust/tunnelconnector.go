// File generated from our OpenAPI spec by Stainless.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// TunnelConnectorService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewTunnelConnectorService] method
// instead.
type TunnelConnectorService struct {
	Options []option.RequestOption
}

// NewTunnelConnectorService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTunnelConnectorService(opts ...option.RequestOption) (r *TunnelConnectorService) {
	r = &TunnelConnectorService{}
	r.Options = opts
	return
}

// Fetches connector and connection details for a Cloudflare Tunnel.
func (r *TunnelConnectorService) Get(ctx context.Context, tunnelID string, connectorID string, query TunnelConnectorGetParams, opts ...option.RequestOption) (res *TunnelTunnelClient, err error) {
	opts = append(r.Options[:], opts...)
	var env TunnelConnectorGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/connectors/%s", query.AccountID, tunnelID, connectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TunnelConnectorGetParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type TunnelConnectorGetResponseEnvelope struct {
	Errors   []TunnelConnectorGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TunnelConnectorGetResponseEnvelopeMessages `json:"messages,required"`
	// A client (typically cloudflared) that maintains connections to a Cloudflare data
	// center.
	Result TunnelTunnelClient `json:"result,required"`
	// Whether the API call was successful
	Success TunnelConnectorGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelConnectorGetResponseEnvelopeJSON    `json:"-"`
}

// tunnelConnectorGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [TunnelConnectorGetResponseEnvelope]
type tunnelConnectorGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConnectorGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelConnectorGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TunnelConnectorGetResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    tunnelConnectorGetResponseEnvelopeErrorsJSON `json:"-"`
}

// tunnelConnectorGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [TunnelConnectorGetResponseEnvelopeErrors]
type tunnelConnectorGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConnectorGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelConnectorGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TunnelConnectorGetResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    tunnelConnectorGetResponseEnvelopeMessagesJSON `json:"-"`
}

// tunnelConnectorGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [TunnelConnectorGetResponseEnvelopeMessages]
type tunnelConnectorGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConnectorGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelConnectorGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelConnectorGetResponseEnvelopeSuccess bool

const (
	TunnelConnectorGetResponseEnvelopeSuccessTrue TunnelConnectorGetResponseEnvelopeSuccess = true
)
