// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZeroTrustTunnelConnectorService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZeroTrustTunnelConnectorService] method instead.
type ZeroTrustTunnelConnectorService struct {
	Options []option.RequestOption
}

// NewZeroTrustTunnelConnectorService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustTunnelConnectorService(opts ...option.RequestOption) (r *ZeroTrustTunnelConnectorService) {
	r = &ZeroTrustTunnelConnectorService{}
	r.Options = opts
	return
}

// Fetches connector and connection details for a Cloudflare Tunnel.
func (r *ZeroTrustTunnelConnectorService) Get(ctx context.Context, tunnelID string, connectorID string, query ZeroTrustTunnelConnectorGetParams, opts ...option.RequestOption) (res *TunnelTunnelClient, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustTunnelConnectorGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/connectors/%s", query.AccountID, tunnelID, connectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustTunnelConnectorGetParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZeroTrustTunnelConnectorGetResponseEnvelope struct {
	Errors   []ZeroTrustTunnelConnectorGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustTunnelConnectorGetResponseEnvelopeMessages `json:"messages,required"`
	// A client (typically cloudflared) that maintains connections to a Cloudflare data
	// center.
	Result TunnelTunnelClient `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustTunnelConnectorGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustTunnelConnectorGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustTunnelConnectorGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustTunnelConnectorGetResponseEnvelope]
type zeroTrustTunnelConnectorGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelConnectorGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustTunnelConnectorGetResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zeroTrustTunnelConnectorGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustTunnelConnectorGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustTunnelConnectorGetResponseEnvelopeErrors]
type zeroTrustTunnelConnectorGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelConnectorGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustTunnelConnectorGetResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zeroTrustTunnelConnectorGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustTunnelConnectorGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustTunnelConnectorGetResponseEnvelopeMessages]
type zeroTrustTunnelConnectorGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelConnectorGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustTunnelConnectorGetResponseEnvelopeSuccess bool

const (
	ZeroTrustTunnelConnectorGetResponseEnvelopeSuccessTrue ZeroTrustTunnelConnectorGetResponseEnvelopeSuccess = true
)
