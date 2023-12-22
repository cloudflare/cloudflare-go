// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountCfdTunnelConnectionService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountCfdTunnelConnectionService] method instead.
type AccountCfdTunnelConnectionService struct {
	Options []option.RequestOption
}

// NewAccountCfdTunnelConnectionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountCfdTunnelConnectionService(opts ...option.RequestOption) (r *AccountCfdTunnelConnectionService) {
	r = &AccountCfdTunnelConnectionService{}
	r.Options = opts
	return
}

// Removes connections that are in a disconnected or pending reconnect state. We
// recommend running this command after shutting down a tunnel.
func (r *AccountCfdTunnelConnectionService) Delete(ctx context.Context, accountIdentifier string, tunnelID string, body AccountCfdTunnelConnectionDeleteParams, opts ...option.RequestOption) (res *AccountCfdTunnelConnectionDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/connections", accountIdentifier, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Fetches connection details for a Cloudflare Tunnel.
func (r *AccountCfdTunnelConnectionService) CloudflareTunnelListCloudflareTunnelConnections(ctx context.Context, accountIdentifier string, tunnelID string, opts ...option.RequestOption) (res *TunnelConnectionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/connections", accountIdentifier, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TunnelConnectionsResponse struct {
	Errors     []TunnelConnectionsResponseError    `json:"errors"`
	Messages   []TunnelConnectionsResponseMessage  `json:"messages"`
	Result     []TunnelConnectionsResponseResult   `json:"result"`
	ResultInfo TunnelConnectionsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success TunnelConnectionsResponseSuccess `json:"success"`
	JSON    tunnelConnectionsResponseJSON    `json:"-"`
}

// tunnelConnectionsResponseJSON contains the JSON metadata for the struct
// [TunnelConnectionsResponse]
type tunnelConnectionsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConnectionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelConnectionsResponseError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    tunnelConnectionsResponseErrorJSON `json:"-"`
}

// tunnelConnectionsResponseErrorJSON contains the JSON metadata for the struct
// [TunnelConnectionsResponseError]
type tunnelConnectionsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConnectionsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelConnectionsResponseMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    tunnelConnectionsResponseMessageJSON `json:"-"`
}

// tunnelConnectionsResponseMessageJSON contains the JSON metadata for the struct
// [TunnelConnectionsResponseMessage]
type tunnelConnectionsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConnectionsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A connection between cloudflared and a Cloudflare data center.
type TunnelConnectionsResponseResult struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id"`
	// The cloudflared OS architecture used to establish this connection.
	Arch string `json:"arch"`
	// The version of the remote tunnel configuration. Used internally to sync
	// cloudflared with the Zero Trust dashboard.
	ConfigVersion int64 `json:"config_version"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Conns []TunnelConnectionsResponseResultConn `json:"conns"`
	// Features enabled for the Cloudflare Tunnel.
	Features []string `json:"features"`
	// Timestamp of when the tunnel connection was started.
	RunAt time.Time `json:"run_at" format:"date-time"`
	// The cloudflared version used to establish this connection.
	Version string                              `json:"version"`
	JSON    tunnelConnectionsResponseResultJSON `json:"-"`
}

// tunnelConnectionsResponseResultJSON contains the JSON metadata for the struct
// [TunnelConnectionsResponseResult]
type tunnelConnectionsResponseResultJSON struct {
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

func (r *TunnelConnectionsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelConnectionsResponseResultConn struct {
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
	Uuid string                                  `json:"uuid"`
	JSON tunnelConnectionsResponseResultConnJSON `json:"-"`
}

// tunnelConnectionsResponseResultConnJSON contains the JSON metadata for the
// struct [TunnelConnectionsResponseResultConn]
type tunnelConnectionsResponseResultConnJSON struct {
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

func (r *TunnelConnectionsResponseResultConn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelConnectionsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                 `json:"total_count"`
	JSON       tunnelConnectionsResponseResultInfoJSON `json:"-"`
}

// tunnelConnectionsResponseResultInfoJSON contains the JSON metadata for the
// struct [TunnelConnectionsResponseResultInfo]
type tunnelConnectionsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConnectionsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TunnelConnectionsResponseSuccess bool

const (
	TunnelConnectionsResponseSuccessTrue TunnelConnectionsResponseSuccess = true
)

type AccountCfdTunnelConnectionDeleteResponse struct {
	Errors   []AccountCfdTunnelConnectionDeleteResponseError   `json:"errors"`
	Messages []AccountCfdTunnelConnectionDeleteResponseMessage `json:"messages"`
	Result   interface{}                                       `json:"result"`
	// Whether the API call was successful
	Success AccountCfdTunnelConnectionDeleteResponseSuccess `json:"success"`
	JSON    accountCfdTunnelConnectionDeleteResponseJSON    `json:"-"`
}

// accountCfdTunnelConnectionDeleteResponseJSON contains the JSON metadata for the
// struct [AccountCfdTunnelConnectionDeleteResponse]
type accountCfdTunnelConnectionDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelConnectionDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCfdTunnelConnectionDeleteResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountCfdTunnelConnectionDeleteResponseErrorJSON `json:"-"`
}

// accountCfdTunnelConnectionDeleteResponseErrorJSON contains the JSON metadata for
// the struct [AccountCfdTunnelConnectionDeleteResponseError]
type accountCfdTunnelConnectionDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelConnectionDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCfdTunnelConnectionDeleteResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountCfdTunnelConnectionDeleteResponseMessageJSON `json:"-"`
}

// accountCfdTunnelConnectionDeleteResponseMessageJSON contains the JSON metadata
// for the struct [AccountCfdTunnelConnectionDeleteResponseMessage]
type accountCfdTunnelConnectionDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelConnectionDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountCfdTunnelConnectionDeleteResponseSuccess bool

const (
	AccountCfdTunnelConnectionDeleteResponseSuccessTrue AccountCfdTunnelConnectionDeleteResponseSuccess = true
)

type AccountCfdTunnelConnectionDeleteParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountCfdTunnelConnectionDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
