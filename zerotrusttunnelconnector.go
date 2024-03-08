// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *ZeroTrustTunnelConnectorService) Get(ctx context.Context, tunnelID string, connectorID string, query ZeroTrustTunnelConnectorGetParams, opts ...option.RequestOption) (res *ZeroTrustTunnelConnectorGetResponse, err error) {
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

// A client (typically cloudflared) that maintains connections to a Cloudflare data
// center.
type ZeroTrustTunnelConnectorGetResponse struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id"`
	// The cloudflared OS architecture used to establish this connection.
	Arch string `json:"arch"`
	// The version of the remote tunnel configuration. Used internally to sync
	// cloudflared with the Zero Trust dashboard.
	ConfigVersion int64 `json:"config_version"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Conns []ZeroTrustTunnelConnectorGetResponseConn `json:"conns"`
	// Features enabled for the Cloudflare Tunnel.
	Features []string `json:"features"`
	// Timestamp of when the tunnel connection was started.
	RunAt time.Time `json:"run_at" format:"date-time"`
	// The cloudflared version used to establish this connection.
	Version string                                  `json:"version"`
	JSON    zeroTrustTunnelConnectorGetResponseJSON `json:"-"`
}

// zeroTrustTunnelConnectorGetResponseJSON contains the JSON metadata for the
// struct [ZeroTrustTunnelConnectorGetResponse]
type zeroTrustTunnelConnectorGetResponseJSON struct {
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

func (r *ZeroTrustTunnelConnectorGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelConnectorGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustTunnelConnectorGetResponseConn struct {
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
	UUID string                                      `json:"uuid"`
	JSON zeroTrustTunnelConnectorGetResponseConnJSON `json:"-"`
}

// zeroTrustTunnelConnectorGetResponseConnJSON contains the JSON metadata for the
// struct [ZeroTrustTunnelConnectorGetResponseConn]
type zeroTrustTunnelConnectorGetResponseConnJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	UUID               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustTunnelConnectorGetResponseConn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelConnectorGetResponseConnJSON) RawJSON() string {
	return r.raw
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
	Result ZeroTrustTunnelConnectorGetResponse `json:"result,required"`
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

func (r zeroTrustTunnelConnectorGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustTunnelConnectorGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustTunnelConnectorGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustTunnelConnectorGetResponseEnvelopeSuccess bool

const (
	ZeroTrustTunnelConnectorGetResponseEnvelopeSuccessTrue ZeroTrustTunnelConnectorGetResponseEnvelopeSuccess = true
)
