// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// WarpConnectorService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWarpConnectorService] method
// instead.
type WarpConnectorService struct {
	Options []option.RequestOption
}

// NewWarpConnectorService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWarpConnectorService(opts ...option.RequestOption) (r *WarpConnectorService) {
	r = &WarpConnectorService{}
	r.Options = opts
	return
}

// Creates a new Warp Connector Tunnel in an account.
func (r *WarpConnectorService) New(ctx context.Context, accountID string, body WarpConnectorNewParams, opts ...option.RequestOption) (res *WarpConnectorNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/warp_connector", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a single Warp Connector Tunnel.
func (r *WarpConnectorService) Get(ctx context.Context, accountID string, tunnelID string, opts ...option.RequestOption) (res *WarpConnectorGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/warp_connector/%s", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing Warp Connector Tunnel.
func (r *WarpConnectorService) Update(ctx context.Context, accountID string, tunnelID string, body WarpConnectorUpdateParams, opts ...option.RequestOption) (res *WarpConnectorUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/warp_connector/%s", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Lists and filters Warp Connector Tunnels in an account.
func (r *WarpConnectorService) List(ctx context.Context, accountID string, query WarpConnectorListParams, opts ...option.RequestOption) (res *WarpConnectorListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/warp_connector", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a Warp Connector Tunnel from an account.
func (r *WarpConnectorService) Delete(ctx context.Context, accountID string, tunnelID string, body WarpConnectorDeleteParams, opts ...option.RequestOption) (res *WarpConnectorDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/warp_connector/%s", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type WarpConnectorNewResponse struct {
	Errors   []WarpConnectorNewResponseError   `json:"errors"`
	Messages []WarpConnectorNewResponseMessage `json:"messages"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result WarpConnectorNewResponseResult `json:"result"`
	// Whether the API call was successful
	Success WarpConnectorNewResponseSuccess `json:"success"`
	JSON    warpConnectorNewResponseJSON    `json:"-"`
}

// warpConnectorNewResponseJSON contains the JSON metadata for the struct
// [WarpConnectorNewResponse]
type warpConnectorNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WarpConnectorNewResponseError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    warpConnectorNewResponseErrorJSON `json:"-"`
}

// warpConnectorNewResponseErrorJSON contains the JSON metadata for the struct
// [WarpConnectorNewResponseError]
type warpConnectorNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WarpConnectorNewResponseMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    warpConnectorNewResponseMessageJSON `json:"-"`
}

// warpConnectorNewResponseMessageJSON contains the JSON metadata for the struct
// [WarpConnectorNewResponseMessage]
type warpConnectorNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [WarpConnectorNewResponseResultXlFXheKkCfdTunnel] or
// [WarpConnectorNewResponseResultXlFXheKkWarpConnectorTunnel].
type WarpConnectorNewResponseResult interface {
	implementsWarpConnectorNewResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WarpConnectorNewResponseResult)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type WarpConnectorNewResponseResultXlFXheKkCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WarpConnectorNewResponseResultXlFXheKkCfdTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for the tunnel.
	Name string `json:"name"`
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	RemoteConfig bool `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status string `json:"status"`
	// The type of tunnel.
	TunType WarpConnectorNewResponseResultXlFXheKkCfdTunnelTunType `json:"tun_type"`
	JSON    warpConnectorNewResponseResultXlFXheKkCfdTunnelJSON    `json:"-"`
}

// warpConnectorNewResponseResultXlFXheKkCfdTunnelJSON contains the JSON metadata
// for the struct [WarpConnectorNewResponseResultXlFXheKkCfdTunnel]
type warpConnectorNewResponseResultXlFXheKkCfdTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	RemoteConfig    apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WarpConnectorNewResponseResultXlFXheKkCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WarpConnectorNewResponseResultXlFXheKkCfdTunnel) implementsWarpConnectorNewResponseResult() {}

type WarpConnectorNewResponseResultXlFXheKkCfdTunnelConnection struct {
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
	Uuid string                                                        `json:"uuid"`
	JSON warpConnectorNewResponseResultXlFXheKkCfdTunnelConnectionJSON `json:"-"`
}

// warpConnectorNewResponseResultXlFXheKkCfdTunnelConnectionJSON contains the JSON
// metadata for the struct
// [WarpConnectorNewResponseResultXlFXheKkCfdTunnelConnection]
type warpConnectorNewResponseResultXlFXheKkCfdTunnelConnectionJSON struct {
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

func (r *WarpConnectorNewResponseResultXlFXheKkCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type WarpConnectorNewResponseResultXlFXheKkCfdTunnelTunType string

const (
	WarpConnectorNewResponseResultXlFXheKkCfdTunnelTunTypeCfdTunnel     WarpConnectorNewResponseResultXlFXheKkCfdTunnelTunType = "cfd_tunnel"
	WarpConnectorNewResponseResultXlFXheKkCfdTunnelTunTypeWarpConnector WarpConnectorNewResponseResultXlFXheKkCfdTunnelTunType = "warp_connector"
	WarpConnectorNewResponseResultXlFXheKkCfdTunnelTunTypeIPSec         WarpConnectorNewResponseResultXlFXheKkCfdTunnelTunType = "ip_sec"
	WarpConnectorNewResponseResultXlFXheKkCfdTunnelTunTypeGre           WarpConnectorNewResponseResultXlFXheKkCfdTunnelTunType = "gre"
	WarpConnectorNewResponseResultXlFXheKkCfdTunnelTunTypeCni           WarpConnectorNewResponseResultXlFXheKkCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type WarpConnectorNewResponseResultXlFXheKkWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WarpConnectorNewResponseResultXlFXheKkWarpConnectorTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for the tunnel.
	Name string `json:"name"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status string `json:"status"`
	// The type of tunnel.
	TunType WarpConnectorNewResponseResultXlFXheKkWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    warpConnectorNewResponseResultXlFXheKkWarpConnectorTunnelJSON    `json:"-"`
}

// warpConnectorNewResponseResultXlFXheKkWarpConnectorTunnelJSON contains the JSON
// metadata for the struct
// [WarpConnectorNewResponseResultXlFXheKkWarpConnectorTunnel]
type warpConnectorNewResponseResultXlFXheKkWarpConnectorTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WarpConnectorNewResponseResultXlFXheKkWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WarpConnectorNewResponseResultXlFXheKkWarpConnectorTunnel) implementsWarpConnectorNewResponseResult() {
}

type WarpConnectorNewResponseResultXlFXheKkWarpConnectorTunnelConnection struct {
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
	Uuid string                                                                  `json:"uuid"`
	JSON warpConnectorNewResponseResultXlFXheKkWarpConnectorTunnelConnectionJSON `json:"-"`
}

// warpConnectorNewResponseResultXlFXheKkWarpConnectorTunnelConnectionJSON contains
// the JSON metadata for the struct
// [WarpConnectorNewResponseResultXlFXheKkWarpConnectorTunnelConnection]
type warpConnectorNewResponseResultXlFXheKkWarpConnectorTunnelConnectionJSON struct {
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

func (r *WarpConnectorNewResponseResultXlFXheKkWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type WarpConnectorNewResponseResultXlFXheKkWarpConnectorTunnelTunType string

const (
	WarpConnectorNewResponseResultXlFXheKkWarpConnectorTunnelTunTypeCfdTunnel     WarpConnectorNewResponseResultXlFXheKkWarpConnectorTunnelTunType = "cfd_tunnel"
	WarpConnectorNewResponseResultXlFXheKkWarpConnectorTunnelTunTypeWarpConnector WarpConnectorNewResponseResultXlFXheKkWarpConnectorTunnelTunType = "warp_connector"
	WarpConnectorNewResponseResultXlFXheKkWarpConnectorTunnelTunTypeIPSec         WarpConnectorNewResponseResultXlFXheKkWarpConnectorTunnelTunType = "ip_sec"
	WarpConnectorNewResponseResultXlFXheKkWarpConnectorTunnelTunTypeGre           WarpConnectorNewResponseResultXlFXheKkWarpConnectorTunnelTunType = "gre"
	WarpConnectorNewResponseResultXlFXheKkWarpConnectorTunnelTunTypeCni           WarpConnectorNewResponseResultXlFXheKkWarpConnectorTunnelTunType = "cni"
)

// Whether the API call was successful
type WarpConnectorNewResponseSuccess bool

const (
	WarpConnectorNewResponseSuccessTrue WarpConnectorNewResponseSuccess = true
)

type WarpConnectorGetResponse struct {
	Errors   []WarpConnectorGetResponseError   `json:"errors"`
	Messages []WarpConnectorGetResponseMessage `json:"messages"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result WarpConnectorGetResponseResult `json:"result"`
	// Whether the API call was successful
	Success WarpConnectorGetResponseSuccess `json:"success"`
	JSON    warpConnectorGetResponseJSON    `json:"-"`
}

// warpConnectorGetResponseJSON contains the JSON metadata for the struct
// [WarpConnectorGetResponse]
type warpConnectorGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WarpConnectorGetResponseError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    warpConnectorGetResponseErrorJSON `json:"-"`
}

// warpConnectorGetResponseErrorJSON contains the JSON metadata for the struct
// [WarpConnectorGetResponseError]
type warpConnectorGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WarpConnectorGetResponseMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    warpConnectorGetResponseMessageJSON `json:"-"`
}

// warpConnectorGetResponseMessageJSON contains the JSON metadata for the struct
// [WarpConnectorGetResponseMessage]
type warpConnectorGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [WarpConnectorGetResponseResultXlFXheKkCfdTunnel] or
// [WarpConnectorGetResponseResultXlFXheKkWarpConnectorTunnel].
type WarpConnectorGetResponseResult interface {
	implementsWarpConnectorGetResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WarpConnectorGetResponseResult)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type WarpConnectorGetResponseResultXlFXheKkCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WarpConnectorGetResponseResultXlFXheKkCfdTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for the tunnel.
	Name string `json:"name"`
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	RemoteConfig bool `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status string `json:"status"`
	// The type of tunnel.
	TunType WarpConnectorGetResponseResultXlFXheKkCfdTunnelTunType `json:"tun_type"`
	JSON    warpConnectorGetResponseResultXlFXheKkCfdTunnelJSON    `json:"-"`
}

// warpConnectorGetResponseResultXlFXheKkCfdTunnelJSON contains the JSON metadata
// for the struct [WarpConnectorGetResponseResultXlFXheKkCfdTunnel]
type warpConnectorGetResponseResultXlFXheKkCfdTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	RemoteConfig    apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WarpConnectorGetResponseResultXlFXheKkCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WarpConnectorGetResponseResultXlFXheKkCfdTunnel) implementsWarpConnectorGetResponseResult() {}

type WarpConnectorGetResponseResultXlFXheKkCfdTunnelConnection struct {
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
	Uuid string                                                        `json:"uuid"`
	JSON warpConnectorGetResponseResultXlFXheKkCfdTunnelConnectionJSON `json:"-"`
}

// warpConnectorGetResponseResultXlFXheKkCfdTunnelConnectionJSON contains the JSON
// metadata for the struct
// [WarpConnectorGetResponseResultXlFXheKkCfdTunnelConnection]
type warpConnectorGetResponseResultXlFXheKkCfdTunnelConnectionJSON struct {
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

func (r *WarpConnectorGetResponseResultXlFXheKkCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type WarpConnectorGetResponseResultXlFXheKkCfdTunnelTunType string

const (
	WarpConnectorGetResponseResultXlFXheKkCfdTunnelTunTypeCfdTunnel     WarpConnectorGetResponseResultXlFXheKkCfdTunnelTunType = "cfd_tunnel"
	WarpConnectorGetResponseResultXlFXheKkCfdTunnelTunTypeWarpConnector WarpConnectorGetResponseResultXlFXheKkCfdTunnelTunType = "warp_connector"
	WarpConnectorGetResponseResultXlFXheKkCfdTunnelTunTypeIPSec         WarpConnectorGetResponseResultXlFXheKkCfdTunnelTunType = "ip_sec"
	WarpConnectorGetResponseResultXlFXheKkCfdTunnelTunTypeGre           WarpConnectorGetResponseResultXlFXheKkCfdTunnelTunType = "gre"
	WarpConnectorGetResponseResultXlFXheKkCfdTunnelTunTypeCni           WarpConnectorGetResponseResultXlFXheKkCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type WarpConnectorGetResponseResultXlFXheKkWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WarpConnectorGetResponseResultXlFXheKkWarpConnectorTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for the tunnel.
	Name string `json:"name"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status string `json:"status"`
	// The type of tunnel.
	TunType WarpConnectorGetResponseResultXlFXheKkWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    warpConnectorGetResponseResultXlFXheKkWarpConnectorTunnelJSON    `json:"-"`
}

// warpConnectorGetResponseResultXlFXheKkWarpConnectorTunnelJSON contains the JSON
// metadata for the struct
// [WarpConnectorGetResponseResultXlFXheKkWarpConnectorTunnel]
type warpConnectorGetResponseResultXlFXheKkWarpConnectorTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WarpConnectorGetResponseResultXlFXheKkWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WarpConnectorGetResponseResultXlFXheKkWarpConnectorTunnel) implementsWarpConnectorGetResponseResult() {
}

type WarpConnectorGetResponseResultXlFXheKkWarpConnectorTunnelConnection struct {
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
	Uuid string                                                                  `json:"uuid"`
	JSON warpConnectorGetResponseResultXlFXheKkWarpConnectorTunnelConnectionJSON `json:"-"`
}

// warpConnectorGetResponseResultXlFXheKkWarpConnectorTunnelConnectionJSON contains
// the JSON metadata for the struct
// [WarpConnectorGetResponseResultXlFXheKkWarpConnectorTunnelConnection]
type warpConnectorGetResponseResultXlFXheKkWarpConnectorTunnelConnectionJSON struct {
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

func (r *WarpConnectorGetResponseResultXlFXheKkWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type WarpConnectorGetResponseResultXlFXheKkWarpConnectorTunnelTunType string

const (
	WarpConnectorGetResponseResultXlFXheKkWarpConnectorTunnelTunTypeCfdTunnel     WarpConnectorGetResponseResultXlFXheKkWarpConnectorTunnelTunType = "cfd_tunnel"
	WarpConnectorGetResponseResultXlFXheKkWarpConnectorTunnelTunTypeWarpConnector WarpConnectorGetResponseResultXlFXheKkWarpConnectorTunnelTunType = "warp_connector"
	WarpConnectorGetResponseResultXlFXheKkWarpConnectorTunnelTunTypeIPSec         WarpConnectorGetResponseResultXlFXheKkWarpConnectorTunnelTunType = "ip_sec"
	WarpConnectorGetResponseResultXlFXheKkWarpConnectorTunnelTunTypeGre           WarpConnectorGetResponseResultXlFXheKkWarpConnectorTunnelTunType = "gre"
	WarpConnectorGetResponseResultXlFXheKkWarpConnectorTunnelTunTypeCni           WarpConnectorGetResponseResultXlFXheKkWarpConnectorTunnelTunType = "cni"
)

// Whether the API call was successful
type WarpConnectorGetResponseSuccess bool

const (
	WarpConnectorGetResponseSuccessTrue WarpConnectorGetResponseSuccess = true
)

type WarpConnectorUpdateResponse struct {
	Errors   []WarpConnectorUpdateResponseError   `json:"errors"`
	Messages []WarpConnectorUpdateResponseMessage `json:"messages"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result WarpConnectorUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success WarpConnectorUpdateResponseSuccess `json:"success"`
	JSON    warpConnectorUpdateResponseJSON    `json:"-"`
}

// warpConnectorUpdateResponseJSON contains the JSON metadata for the struct
// [WarpConnectorUpdateResponse]
type warpConnectorUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WarpConnectorUpdateResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    warpConnectorUpdateResponseErrorJSON `json:"-"`
}

// warpConnectorUpdateResponseErrorJSON contains the JSON metadata for the struct
// [WarpConnectorUpdateResponseError]
type warpConnectorUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WarpConnectorUpdateResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    warpConnectorUpdateResponseMessageJSON `json:"-"`
}

// warpConnectorUpdateResponseMessageJSON contains the JSON metadata for the struct
// [WarpConnectorUpdateResponseMessage]
type warpConnectorUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [WarpConnectorUpdateResponseResultXlFXheKkCfdTunnel] or
// [WarpConnectorUpdateResponseResultXlFXheKkWarpConnectorTunnel].
type WarpConnectorUpdateResponseResult interface {
	implementsWarpConnectorUpdateResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WarpConnectorUpdateResponseResult)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type WarpConnectorUpdateResponseResultXlFXheKkCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WarpConnectorUpdateResponseResultXlFXheKkCfdTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for the tunnel.
	Name string `json:"name"`
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	RemoteConfig bool `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status string `json:"status"`
	// The type of tunnel.
	TunType WarpConnectorUpdateResponseResultXlFXheKkCfdTunnelTunType `json:"tun_type"`
	JSON    warpConnectorUpdateResponseResultXlFXheKkCfdTunnelJSON    `json:"-"`
}

// warpConnectorUpdateResponseResultXlFXheKkCfdTunnelJSON contains the JSON
// metadata for the struct [WarpConnectorUpdateResponseResultXlFXheKkCfdTunnel]
type warpConnectorUpdateResponseResultXlFXheKkCfdTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	RemoteConfig    apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WarpConnectorUpdateResponseResultXlFXheKkCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WarpConnectorUpdateResponseResultXlFXheKkCfdTunnel) implementsWarpConnectorUpdateResponseResult() {
}

type WarpConnectorUpdateResponseResultXlFXheKkCfdTunnelConnection struct {
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
	Uuid string                                                           `json:"uuid"`
	JSON warpConnectorUpdateResponseResultXlFXheKkCfdTunnelConnectionJSON `json:"-"`
}

// warpConnectorUpdateResponseResultXlFXheKkCfdTunnelConnectionJSON contains the
// JSON metadata for the struct
// [WarpConnectorUpdateResponseResultXlFXheKkCfdTunnelConnection]
type warpConnectorUpdateResponseResultXlFXheKkCfdTunnelConnectionJSON struct {
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

func (r *WarpConnectorUpdateResponseResultXlFXheKkCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type WarpConnectorUpdateResponseResultXlFXheKkCfdTunnelTunType string

const (
	WarpConnectorUpdateResponseResultXlFXheKkCfdTunnelTunTypeCfdTunnel     WarpConnectorUpdateResponseResultXlFXheKkCfdTunnelTunType = "cfd_tunnel"
	WarpConnectorUpdateResponseResultXlFXheKkCfdTunnelTunTypeWarpConnector WarpConnectorUpdateResponseResultXlFXheKkCfdTunnelTunType = "warp_connector"
	WarpConnectorUpdateResponseResultXlFXheKkCfdTunnelTunTypeIPSec         WarpConnectorUpdateResponseResultXlFXheKkCfdTunnelTunType = "ip_sec"
	WarpConnectorUpdateResponseResultXlFXheKkCfdTunnelTunTypeGre           WarpConnectorUpdateResponseResultXlFXheKkCfdTunnelTunType = "gre"
	WarpConnectorUpdateResponseResultXlFXheKkCfdTunnelTunTypeCni           WarpConnectorUpdateResponseResultXlFXheKkCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type WarpConnectorUpdateResponseResultXlFXheKkWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WarpConnectorUpdateResponseResultXlFXheKkWarpConnectorTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for the tunnel.
	Name string `json:"name"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status string `json:"status"`
	// The type of tunnel.
	TunType WarpConnectorUpdateResponseResultXlFXheKkWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    warpConnectorUpdateResponseResultXlFXheKkWarpConnectorTunnelJSON    `json:"-"`
}

// warpConnectorUpdateResponseResultXlFXheKkWarpConnectorTunnelJSON contains the
// JSON metadata for the struct
// [WarpConnectorUpdateResponseResultXlFXheKkWarpConnectorTunnel]
type warpConnectorUpdateResponseResultXlFXheKkWarpConnectorTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WarpConnectorUpdateResponseResultXlFXheKkWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WarpConnectorUpdateResponseResultXlFXheKkWarpConnectorTunnel) implementsWarpConnectorUpdateResponseResult() {
}

type WarpConnectorUpdateResponseResultXlFXheKkWarpConnectorTunnelConnection struct {
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
	Uuid string                                                                     `json:"uuid"`
	JSON warpConnectorUpdateResponseResultXlFXheKkWarpConnectorTunnelConnectionJSON `json:"-"`
}

// warpConnectorUpdateResponseResultXlFXheKkWarpConnectorTunnelConnectionJSON
// contains the JSON metadata for the struct
// [WarpConnectorUpdateResponseResultXlFXheKkWarpConnectorTunnelConnection]
type warpConnectorUpdateResponseResultXlFXheKkWarpConnectorTunnelConnectionJSON struct {
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

func (r *WarpConnectorUpdateResponseResultXlFXheKkWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type WarpConnectorUpdateResponseResultXlFXheKkWarpConnectorTunnelTunType string

const (
	WarpConnectorUpdateResponseResultXlFXheKkWarpConnectorTunnelTunTypeCfdTunnel     WarpConnectorUpdateResponseResultXlFXheKkWarpConnectorTunnelTunType = "cfd_tunnel"
	WarpConnectorUpdateResponseResultXlFXheKkWarpConnectorTunnelTunTypeWarpConnector WarpConnectorUpdateResponseResultXlFXheKkWarpConnectorTunnelTunType = "warp_connector"
	WarpConnectorUpdateResponseResultXlFXheKkWarpConnectorTunnelTunTypeIPSec         WarpConnectorUpdateResponseResultXlFXheKkWarpConnectorTunnelTunType = "ip_sec"
	WarpConnectorUpdateResponseResultXlFXheKkWarpConnectorTunnelTunTypeGre           WarpConnectorUpdateResponseResultXlFXheKkWarpConnectorTunnelTunType = "gre"
	WarpConnectorUpdateResponseResultXlFXheKkWarpConnectorTunnelTunTypeCni           WarpConnectorUpdateResponseResultXlFXheKkWarpConnectorTunnelTunType = "cni"
)

// Whether the API call was successful
type WarpConnectorUpdateResponseSuccess bool

const (
	WarpConnectorUpdateResponseSuccessTrue WarpConnectorUpdateResponseSuccess = true
)

type WarpConnectorListResponse struct {
	Errors     []WarpConnectorListResponseError    `json:"errors"`
	Messages   []WarpConnectorListResponseMessage  `json:"messages"`
	Result     []WarpConnectorListResponseResult   `json:"result"`
	ResultInfo WarpConnectorListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success WarpConnectorListResponseSuccess `json:"success"`
	JSON    warpConnectorListResponseJSON    `json:"-"`
}

// warpConnectorListResponseJSON contains the JSON metadata for the struct
// [WarpConnectorListResponse]
type warpConnectorListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WarpConnectorListResponseError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    warpConnectorListResponseErrorJSON `json:"-"`
}

// warpConnectorListResponseErrorJSON contains the JSON metadata for the struct
// [WarpConnectorListResponseError]
type warpConnectorListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WarpConnectorListResponseMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    warpConnectorListResponseMessageJSON `json:"-"`
}

// warpConnectorListResponseMessageJSON contains the JSON metadata for the struct
// [WarpConnectorListResponseMessage]
type warpConnectorListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [WarpConnectorListResponseResultXlFXheKkCfdTunnel] or
// [WarpConnectorListResponseResultXlFXheKkWarpConnectorTunnel].
type WarpConnectorListResponseResult interface {
	implementsWarpConnectorListResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WarpConnectorListResponseResult)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type WarpConnectorListResponseResultXlFXheKkCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WarpConnectorListResponseResultXlFXheKkCfdTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for the tunnel.
	Name string `json:"name"`
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	RemoteConfig bool `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status string `json:"status"`
	// The type of tunnel.
	TunType WarpConnectorListResponseResultXlFXheKkCfdTunnelTunType `json:"tun_type"`
	JSON    warpConnectorListResponseResultXlFXheKkCfdTunnelJSON    `json:"-"`
}

// warpConnectorListResponseResultXlFXheKkCfdTunnelJSON contains the JSON metadata
// for the struct [WarpConnectorListResponseResultXlFXheKkCfdTunnel]
type warpConnectorListResponseResultXlFXheKkCfdTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	RemoteConfig    apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WarpConnectorListResponseResultXlFXheKkCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WarpConnectorListResponseResultXlFXheKkCfdTunnel) implementsWarpConnectorListResponseResult() {
}

type WarpConnectorListResponseResultXlFXheKkCfdTunnelConnection struct {
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
	Uuid string                                                         `json:"uuid"`
	JSON warpConnectorListResponseResultXlFXheKkCfdTunnelConnectionJSON `json:"-"`
}

// warpConnectorListResponseResultXlFXheKkCfdTunnelConnectionJSON contains the JSON
// metadata for the struct
// [WarpConnectorListResponseResultXlFXheKkCfdTunnelConnection]
type warpConnectorListResponseResultXlFXheKkCfdTunnelConnectionJSON struct {
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

func (r *WarpConnectorListResponseResultXlFXheKkCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type WarpConnectorListResponseResultXlFXheKkCfdTunnelTunType string

const (
	WarpConnectorListResponseResultXlFXheKkCfdTunnelTunTypeCfdTunnel     WarpConnectorListResponseResultXlFXheKkCfdTunnelTunType = "cfd_tunnel"
	WarpConnectorListResponseResultXlFXheKkCfdTunnelTunTypeWarpConnector WarpConnectorListResponseResultXlFXheKkCfdTunnelTunType = "warp_connector"
	WarpConnectorListResponseResultXlFXheKkCfdTunnelTunTypeIPSec         WarpConnectorListResponseResultXlFXheKkCfdTunnelTunType = "ip_sec"
	WarpConnectorListResponseResultXlFXheKkCfdTunnelTunTypeGre           WarpConnectorListResponseResultXlFXheKkCfdTunnelTunType = "gre"
	WarpConnectorListResponseResultXlFXheKkCfdTunnelTunTypeCni           WarpConnectorListResponseResultXlFXheKkCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type WarpConnectorListResponseResultXlFXheKkWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WarpConnectorListResponseResultXlFXheKkWarpConnectorTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for the tunnel.
	Name string `json:"name"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status string `json:"status"`
	// The type of tunnel.
	TunType WarpConnectorListResponseResultXlFXheKkWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    warpConnectorListResponseResultXlFXheKkWarpConnectorTunnelJSON    `json:"-"`
}

// warpConnectorListResponseResultXlFXheKkWarpConnectorTunnelJSON contains the JSON
// metadata for the struct
// [WarpConnectorListResponseResultXlFXheKkWarpConnectorTunnel]
type warpConnectorListResponseResultXlFXheKkWarpConnectorTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WarpConnectorListResponseResultXlFXheKkWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WarpConnectorListResponseResultXlFXheKkWarpConnectorTunnel) implementsWarpConnectorListResponseResult() {
}

type WarpConnectorListResponseResultXlFXheKkWarpConnectorTunnelConnection struct {
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
	Uuid string                                                                   `json:"uuid"`
	JSON warpConnectorListResponseResultXlFXheKkWarpConnectorTunnelConnectionJSON `json:"-"`
}

// warpConnectorListResponseResultXlFXheKkWarpConnectorTunnelConnectionJSON
// contains the JSON metadata for the struct
// [WarpConnectorListResponseResultXlFXheKkWarpConnectorTunnelConnection]
type warpConnectorListResponseResultXlFXheKkWarpConnectorTunnelConnectionJSON struct {
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

func (r *WarpConnectorListResponseResultXlFXheKkWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type WarpConnectorListResponseResultXlFXheKkWarpConnectorTunnelTunType string

const (
	WarpConnectorListResponseResultXlFXheKkWarpConnectorTunnelTunTypeCfdTunnel     WarpConnectorListResponseResultXlFXheKkWarpConnectorTunnelTunType = "cfd_tunnel"
	WarpConnectorListResponseResultXlFXheKkWarpConnectorTunnelTunTypeWarpConnector WarpConnectorListResponseResultXlFXheKkWarpConnectorTunnelTunType = "warp_connector"
	WarpConnectorListResponseResultXlFXheKkWarpConnectorTunnelTunTypeIPSec         WarpConnectorListResponseResultXlFXheKkWarpConnectorTunnelTunType = "ip_sec"
	WarpConnectorListResponseResultXlFXheKkWarpConnectorTunnelTunTypeGre           WarpConnectorListResponseResultXlFXheKkWarpConnectorTunnelTunType = "gre"
	WarpConnectorListResponseResultXlFXheKkWarpConnectorTunnelTunTypeCni           WarpConnectorListResponseResultXlFXheKkWarpConnectorTunnelTunType = "cni"
)

type WarpConnectorListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                 `json:"total_count"`
	JSON       warpConnectorListResponseResultInfoJSON `json:"-"`
}

// warpConnectorListResponseResultInfoJSON contains the JSON metadata for the
// struct [WarpConnectorListResponseResultInfo]
type warpConnectorListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WarpConnectorListResponseSuccess bool

const (
	WarpConnectorListResponseSuccessTrue WarpConnectorListResponseSuccess = true
)

type WarpConnectorDeleteResponse struct {
	Errors   []WarpConnectorDeleteResponseError   `json:"errors"`
	Messages []WarpConnectorDeleteResponseMessage `json:"messages"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result WarpConnectorDeleteResponseResult `json:"result"`
	// Whether the API call was successful
	Success WarpConnectorDeleteResponseSuccess `json:"success"`
	JSON    warpConnectorDeleteResponseJSON    `json:"-"`
}

// warpConnectorDeleteResponseJSON contains the JSON metadata for the struct
// [WarpConnectorDeleteResponse]
type warpConnectorDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WarpConnectorDeleteResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    warpConnectorDeleteResponseErrorJSON `json:"-"`
}

// warpConnectorDeleteResponseErrorJSON contains the JSON metadata for the struct
// [WarpConnectorDeleteResponseError]
type warpConnectorDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WarpConnectorDeleteResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    warpConnectorDeleteResponseMessageJSON `json:"-"`
}

// warpConnectorDeleteResponseMessageJSON contains the JSON metadata for the struct
// [WarpConnectorDeleteResponseMessage]
type warpConnectorDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [WarpConnectorDeleteResponseResultXlFXheKkCfdTunnel] or
// [WarpConnectorDeleteResponseResultXlFXheKkWarpConnectorTunnel].
type WarpConnectorDeleteResponseResult interface {
	implementsWarpConnectorDeleteResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WarpConnectorDeleteResponseResult)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type WarpConnectorDeleteResponseResultXlFXheKkCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WarpConnectorDeleteResponseResultXlFXheKkCfdTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for the tunnel.
	Name string `json:"name"`
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	RemoteConfig bool `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status string `json:"status"`
	// The type of tunnel.
	TunType WarpConnectorDeleteResponseResultXlFXheKkCfdTunnelTunType `json:"tun_type"`
	JSON    warpConnectorDeleteResponseResultXlFXheKkCfdTunnelJSON    `json:"-"`
}

// warpConnectorDeleteResponseResultXlFXheKkCfdTunnelJSON contains the JSON
// metadata for the struct [WarpConnectorDeleteResponseResultXlFXheKkCfdTunnel]
type warpConnectorDeleteResponseResultXlFXheKkCfdTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	RemoteConfig    apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WarpConnectorDeleteResponseResultXlFXheKkCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WarpConnectorDeleteResponseResultXlFXheKkCfdTunnel) implementsWarpConnectorDeleteResponseResult() {
}

type WarpConnectorDeleteResponseResultXlFXheKkCfdTunnelConnection struct {
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
	Uuid string                                                           `json:"uuid"`
	JSON warpConnectorDeleteResponseResultXlFXheKkCfdTunnelConnectionJSON `json:"-"`
}

// warpConnectorDeleteResponseResultXlFXheKkCfdTunnelConnectionJSON contains the
// JSON metadata for the struct
// [WarpConnectorDeleteResponseResultXlFXheKkCfdTunnelConnection]
type warpConnectorDeleteResponseResultXlFXheKkCfdTunnelConnectionJSON struct {
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

func (r *WarpConnectorDeleteResponseResultXlFXheKkCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type WarpConnectorDeleteResponseResultXlFXheKkCfdTunnelTunType string

const (
	WarpConnectorDeleteResponseResultXlFXheKkCfdTunnelTunTypeCfdTunnel     WarpConnectorDeleteResponseResultXlFXheKkCfdTunnelTunType = "cfd_tunnel"
	WarpConnectorDeleteResponseResultXlFXheKkCfdTunnelTunTypeWarpConnector WarpConnectorDeleteResponseResultXlFXheKkCfdTunnelTunType = "warp_connector"
	WarpConnectorDeleteResponseResultXlFXheKkCfdTunnelTunTypeIPSec         WarpConnectorDeleteResponseResultXlFXheKkCfdTunnelTunType = "ip_sec"
	WarpConnectorDeleteResponseResultXlFXheKkCfdTunnelTunTypeGre           WarpConnectorDeleteResponseResultXlFXheKkCfdTunnelTunType = "gre"
	WarpConnectorDeleteResponseResultXlFXheKkCfdTunnelTunTypeCni           WarpConnectorDeleteResponseResultXlFXheKkCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type WarpConnectorDeleteResponseResultXlFXheKkWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WarpConnectorDeleteResponseResultXlFXheKkWarpConnectorTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for the tunnel.
	Name string `json:"name"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status string `json:"status"`
	// The type of tunnel.
	TunType WarpConnectorDeleteResponseResultXlFXheKkWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    warpConnectorDeleteResponseResultXlFXheKkWarpConnectorTunnelJSON    `json:"-"`
}

// warpConnectorDeleteResponseResultXlFXheKkWarpConnectorTunnelJSON contains the
// JSON metadata for the struct
// [WarpConnectorDeleteResponseResultXlFXheKkWarpConnectorTunnel]
type warpConnectorDeleteResponseResultXlFXheKkWarpConnectorTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WarpConnectorDeleteResponseResultXlFXheKkWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WarpConnectorDeleteResponseResultXlFXheKkWarpConnectorTunnel) implementsWarpConnectorDeleteResponseResult() {
}

type WarpConnectorDeleteResponseResultXlFXheKkWarpConnectorTunnelConnection struct {
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
	Uuid string                                                                     `json:"uuid"`
	JSON warpConnectorDeleteResponseResultXlFXheKkWarpConnectorTunnelConnectionJSON `json:"-"`
}

// warpConnectorDeleteResponseResultXlFXheKkWarpConnectorTunnelConnectionJSON
// contains the JSON metadata for the struct
// [WarpConnectorDeleteResponseResultXlFXheKkWarpConnectorTunnelConnection]
type warpConnectorDeleteResponseResultXlFXheKkWarpConnectorTunnelConnectionJSON struct {
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

func (r *WarpConnectorDeleteResponseResultXlFXheKkWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type WarpConnectorDeleteResponseResultXlFXheKkWarpConnectorTunnelTunType string

const (
	WarpConnectorDeleteResponseResultXlFXheKkWarpConnectorTunnelTunTypeCfdTunnel     WarpConnectorDeleteResponseResultXlFXheKkWarpConnectorTunnelTunType = "cfd_tunnel"
	WarpConnectorDeleteResponseResultXlFXheKkWarpConnectorTunnelTunTypeWarpConnector WarpConnectorDeleteResponseResultXlFXheKkWarpConnectorTunnelTunType = "warp_connector"
	WarpConnectorDeleteResponseResultXlFXheKkWarpConnectorTunnelTunTypeIPSec         WarpConnectorDeleteResponseResultXlFXheKkWarpConnectorTunnelTunType = "ip_sec"
	WarpConnectorDeleteResponseResultXlFXheKkWarpConnectorTunnelTunTypeGre           WarpConnectorDeleteResponseResultXlFXheKkWarpConnectorTunnelTunType = "gre"
	WarpConnectorDeleteResponseResultXlFXheKkWarpConnectorTunnelTunTypeCni           WarpConnectorDeleteResponseResultXlFXheKkWarpConnectorTunnelTunType = "cni"
)

// Whether the API call was successful
type WarpConnectorDeleteResponseSuccess bool

const (
	WarpConnectorDeleteResponseSuccessTrue WarpConnectorDeleteResponseSuccess = true
)

type WarpConnectorNewParams struct {
	// A user-friendly name for the tunnel.
	Name param.Field[string] `json:"name,required"`
}

func (r WarpConnectorNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WarpConnectorUpdateParams struct {
	// A user-friendly name for the tunnel.
	Name param.Field[string] `json:"name"`
	// Sets the password required to run a locally-managed tunnel. Must be at least 32
	// bytes and encoded as a base64 string.
	TunnelSecret param.Field[string] `json:"tunnel_secret"`
}

func (r WarpConnectorUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WarpConnectorListParams struct {
	ExcludePrefix param.Field[string] `query:"exclude_prefix"`
	// If provided, include only tunnels that were created (and not deleted) before
	// this time.
	ExistedAt     param.Field[time.Time] `query:"existed_at" format:"date-time"`
	IncludePrefix param.Field[string]    `query:"include_prefix"`
	// If `true`, only include deleted tunnels. If `false`, exclude deleted tunnels. If
	// empty, all tunnels will be included.
	IsDeleted param.Field[bool] `query:"is_deleted"`
	// A user-friendly name for the tunnel.
	Name param.Field[string] `query:"name"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of results to display.
	PerPage       param.Field[float64]   `query:"per_page"`
	WasActiveAt   param.Field[time.Time] `query:"was_active_at" format:"date-time"`
	WasInactiveAt param.Field[time.Time] `query:"was_inactive_at" format:"date-time"`
}

// URLQuery serializes [WarpConnectorListParams]'s query parameters as
// `url.Values`.
func (r WarpConnectorListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WarpConnectorDeleteParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r WarpConnectorDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
