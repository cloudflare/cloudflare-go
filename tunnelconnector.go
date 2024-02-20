// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
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
func (r *TunnelConnectorService) Get(ctx context.Context, accountID string, tunnelID string, connectorID string, opts ...option.RequestOption) (res *TunnelConnectorGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TunnelConnectorGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/connectors/%s", accountID, tunnelID, connectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A client (typically cloudflared) that maintains connections to a Cloudflare data
// center.
type TunnelConnectorGetResponse struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id"`
	// The cloudflared OS architecture used to establish this connection.
	Arch string `json:"arch"`
	// The version of the remote tunnel configuration. Used internally to sync
	// cloudflared with the Zero Trust dashboard.
	ConfigVersion int64 `json:"config_version"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Conns []TunnelConnectorGetResponseConn `json:"conns"`
	// Features enabled for the Cloudflare Tunnel.
	Features []string `json:"features"`
	// Timestamp of when the tunnel connection was started.
	RunAt time.Time `json:"run_at" format:"date-time"`
	// The cloudflared version used to establish this connection.
	Version string                         `json:"version"`
	JSON    tunnelConnectorGetResponseJSON `json:"-"`
}

// tunnelConnectorGetResponseJSON contains the JSON metadata for the struct
// [TunnelConnectorGetResponse]
type tunnelConnectorGetResponseJSON struct {
	ID            apijson.Field
	Arch          apijson.Field
	ConfigVersion apijson.Field
	Conns         apijson.Field
	Features      apijson.Field
	RunAt         apijson.Field
	Version       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TunnelConnectorGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelConnectorGetResponseConn struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id"`
	// UUID of the cloudflared instance.
	ClientID interface{} `json:"client_id"`
	// The cloudflared version used to establish this connection.
	ClientVersion string `json:"client_version"`
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// Timestamp of when the connection was established.
	OpenedAt time.Time `json:"opened_at" format:"date-time"`
	// The public IP address of the host running cloudflared.
	OriginIP string `json:"origin_ip"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                             `json:"uuid"`
	JSON tunnelConnectorGetResponseConnJSON `json:"-"`
}

// tunnelConnectorGetResponseConnJSON contains the JSON metadata for the struct
// [TunnelConnectorGetResponseConn]
type tunnelConnectorGetResponseConnJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TunnelConnectorGetResponseConn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelConnectorGetResponseEnvelope struct {
	Errors   []TunnelConnectorGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TunnelConnectorGetResponseEnvelopeMessages `json:"messages,required"`
	// A client (typically cloudflared) that maintains connections to a Cloudflare data
	// center.
	Result TunnelConnectorGetResponse `json:"result,required"`
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

// Whether the API call was successful
type TunnelConnectorGetResponseEnvelopeSuccess bool

const (
	TunnelConnectorGetResponseEnvelopeSuccessTrue TunnelConnectorGetResponseEnvelopeSuccess = true
)
