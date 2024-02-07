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
	var env WarpConnectorNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/warp_connector", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Warp Connector Tunnel.
func (r *WarpConnectorService) Get(ctx context.Context, accountID string, tunnelID string, opts ...option.RequestOption) (res *WarpConnectorGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WarpConnectorGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/warp_connector/%s", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing Warp Connector Tunnel.
func (r *WarpConnectorService) Update(ctx context.Context, accountID string, tunnelID string, body WarpConnectorUpdateParams, opts ...option.RequestOption) (res *WarpConnectorUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WarpConnectorUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/warp_connector/%s", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists and filters Warp Connector Tunnels in an account.
func (r *WarpConnectorService) List(ctx context.Context, accountID string, query WarpConnectorListParams, opts ...option.RequestOption) (res *[]WarpConnectorListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WarpConnectorListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/warp_connector", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a Warp Connector Tunnel from an account.
func (r *WarpConnectorService) Delete(ctx context.Context, accountID string, tunnelID string, body WarpConnectorDeleteParams, opts ...option.RequestOption) (res *WarpConnectorDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WarpConnectorDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/warp_connector/%s", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [WarpConnectorNewResponseXlFXheKkCfdTunnel] or
// [WarpConnectorNewResponseXlFXheKkWarpConnectorTunnel].
type WarpConnectorNewResponse interface {
	implementsWarpConnectorNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WarpConnectorNewResponse)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type WarpConnectorNewResponseXlFXheKkCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WarpConnectorNewResponseXlFXheKkCfdTunnelConnection `json:"connections"`
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
	TunType WarpConnectorNewResponseXlFXheKkCfdTunnelTunType `json:"tun_type"`
	JSON    warpConnectorNewResponseXlFXheKkCfdTunnelJSON    `json:"-"`
}

// warpConnectorNewResponseXlFXheKkCfdTunnelJSON contains the JSON metadata for the
// struct [WarpConnectorNewResponseXlFXheKkCfdTunnel]
type warpConnectorNewResponseXlFXheKkCfdTunnelJSON struct {
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

func (r *WarpConnectorNewResponseXlFXheKkCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WarpConnectorNewResponseXlFXheKkCfdTunnel) implementsWarpConnectorNewResponse() {}

type WarpConnectorNewResponseXlFXheKkCfdTunnelConnection struct {
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
	Uuid string                                                  `json:"uuid"`
	JSON warpConnectorNewResponseXlFXheKkCfdTunnelConnectionJSON `json:"-"`
}

// warpConnectorNewResponseXlFXheKkCfdTunnelConnectionJSON contains the JSON
// metadata for the struct [WarpConnectorNewResponseXlFXheKkCfdTunnelConnection]
type warpConnectorNewResponseXlFXheKkCfdTunnelConnectionJSON struct {
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

func (r *WarpConnectorNewResponseXlFXheKkCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type WarpConnectorNewResponseXlFXheKkCfdTunnelTunType string

const (
	WarpConnectorNewResponseXlFXheKkCfdTunnelTunTypeCfdTunnel     WarpConnectorNewResponseXlFXheKkCfdTunnelTunType = "cfd_tunnel"
	WarpConnectorNewResponseXlFXheKkCfdTunnelTunTypeWarpConnector WarpConnectorNewResponseXlFXheKkCfdTunnelTunType = "warp_connector"
	WarpConnectorNewResponseXlFXheKkCfdTunnelTunTypeIPSec         WarpConnectorNewResponseXlFXheKkCfdTunnelTunType = "ip_sec"
	WarpConnectorNewResponseXlFXheKkCfdTunnelTunTypeGre           WarpConnectorNewResponseXlFXheKkCfdTunnelTunType = "gre"
	WarpConnectorNewResponseXlFXheKkCfdTunnelTunTypeCni           WarpConnectorNewResponseXlFXheKkCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type WarpConnectorNewResponseXlFXheKkWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WarpConnectorNewResponseXlFXheKkWarpConnectorTunnelConnection `json:"connections"`
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
	TunType WarpConnectorNewResponseXlFXheKkWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    warpConnectorNewResponseXlFXheKkWarpConnectorTunnelJSON    `json:"-"`
}

// warpConnectorNewResponseXlFXheKkWarpConnectorTunnelJSON contains the JSON
// metadata for the struct [WarpConnectorNewResponseXlFXheKkWarpConnectorTunnel]
type warpConnectorNewResponseXlFXheKkWarpConnectorTunnelJSON struct {
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

func (r *WarpConnectorNewResponseXlFXheKkWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WarpConnectorNewResponseXlFXheKkWarpConnectorTunnel) implementsWarpConnectorNewResponse() {}

type WarpConnectorNewResponseXlFXheKkWarpConnectorTunnelConnection struct {
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
	Uuid string                                                            `json:"uuid"`
	JSON warpConnectorNewResponseXlFXheKkWarpConnectorTunnelConnectionJSON `json:"-"`
}

// warpConnectorNewResponseXlFXheKkWarpConnectorTunnelConnectionJSON contains the
// JSON metadata for the struct
// [WarpConnectorNewResponseXlFXheKkWarpConnectorTunnelConnection]
type warpConnectorNewResponseXlFXheKkWarpConnectorTunnelConnectionJSON struct {
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

func (r *WarpConnectorNewResponseXlFXheKkWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type WarpConnectorNewResponseXlFXheKkWarpConnectorTunnelTunType string

const (
	WarpConnectorNewResponseXlFXheKkWarpConnectorTunnelTunTypeCfdTunnel     WarpConnectorNewResponseXlFXheKkWarpConnectorTunnelTunType = "cfd_tunnel"
	WarpConnectorNewResponseXlFXheKkWarpConnectorTunnelTunTypeWarpConnector WarpConnectorNewResponseXlFXheKkWarpConnectorTunnelTunType = "warp_connector"
	WarpConnectorNewResponseXlFXheKkWarpConnectorTunnelTunTypeIPSec         WarpConnectorNewResponseXlFXheKkWarpConnectorTunnelTunType = "ip_sec"
	WarpConnectorNewResponseXlFXheKkWarpConnectorTunnelTunTypeGre           WarpConnectorNewResponseXlFXheKkWarpConnectorTunnelTunType = "gre"
	WarpConnectorNewResponseXlFXheKkWarpConnectorTunnelTunTypeCni           WarpConnectorNewResponseXlFXheKkWarpConnectorTunnelTunType = "cni"
)

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [WarpConnectorGetResponseXlFXheKkCfdTunnel] or
// [WarpConnectorGetResponseXlFXheKkWarpConnectorTunnel].
type WarpConnectorGetResponse interface {
	implementsWarpConnectorGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WarpConnectorGetResponse)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type WarpConnectorGetResponseXlFXheKkCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WarpConnectorGetResponseXlFXheKkCfdTunnelConnection `json:"connections"`
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
	TunType WarpConnectorGetResponseXlFXheKkCfdTunnelTunType `json:"tun_type"`
	JSON    warpConnectorGetResponseXlFXheKkCfdTunnelJSON    `json:"-"`
}

// warpConnectorGetResponseXlFXheKkCfdTunnelJSON contains the JSON metadata for the
// struct [WarpConnectorGetResponseXlFXheKkCfdTunnel]
type warpConnectorGetResponseXlFXheKkCfdTunnelJSON struct {
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

func (r *WarpConnectorGetResponseXlFXheKkCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WarpConnectorGetResponseXlFXheKkCfdTunnel) implementsWarpConnectorGetResponse() {}

type WarpConnectorGetResponseXlFXheKkCfdTunnelConnection struct {
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
	Uuid string                                                  `json:"uuid"`
	JSON warpConnectorGetResponseXlFXheKkCfdTunnelConnectionJSON `json:"-"`
}

// warpConnectorGetResponseXlFXheKkCfdTunnelConnectionJSON contains the JSON
// metadata for the struct [WarpConnectorGetResponseXlFXheKkCfdTunnelConnection]
type warpConnectorGetResponseXlFXheKkCfdTunnelConnectionJSON struct {
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

func (r *WarpConnectorGetResponseXlFXheKkCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type WarpConnectorGetResponseXlFXheKkCfdTunnelTunType string

const (
	WarpConnectorGetResponseXlFXheKkCfdTunnelTunTypeCfdTunnel     WarpConnectorGetResponseXlFXheKkCfdTunnelTunType = "cfd_tunnel"
	WarpConnectorGetResponseXlFXheKkCfdTunnelTunTypeWarpConnector WarpConnectorGetResponseXlFXheKkCfdTunnelTunType = "warp_connector"
	WarpConnectorGetResponseXlFXheKkCfdTunnelTunTypeIPSec         WarpConnectorGetResponseXlFXheKkCfdTunnelTunType = "ip_sec"
	WarpConnectorGetResponseXlFXheKkCfdTunnelTunTypeGre           WarpConnectorGetResponseXlFXheKkCfdTunnelTunType = "gre"
	WarpConnectorGetResponseXlFXheKkCfdTunnelTunTypeCni           WarpConnectorGetResponseXlFXheKkCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type WarpConnectorGetResponseXlFXheKkWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WarpConnectorGetResponseXlFXheKkWarpConnectorTunnelConnection `json:"connections"`
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
	TunType WarpConnectorGetResponseXlFXheKkWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    warpConnectorGetResponseXlFXheKkWarpConnectorTunnelJSON    `json:"-"`
}

// warpConnectorGetResponseXlFXheKkWarpConnectorTunnelJSON contains the JSON
// metadata for the struct [WarpConnectorGetResponseXlFXheKkWarpConnectorTunnel]
type warpConnectorGetResponseXlFXheKkWarpConnectorTunnelJSON struct {
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

func (r *WarpConnectorGetResponseXlFXheKkWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WarpConnectorGetResponseXlFXheKkWarpConnectorTunnel) implementsWarpConnectorGetResponse() {}

type WarpConnectorGetResponseXlFXheKkWarpConnectorTunnelConnection struct {
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
	Uuid string                                                            `json:"uuid"`
	JSON warpConnectorGetResponseXlFXheKkWarpConnectorTunnelConnectionJSON `json:"-"`
}

// warpConnectorGetResponseXlFXheKkWarpConnectorTunnelConnectionJSON contains the
// JSON metadata for the struct
// [WarpConnectorGetResponseXlFXheKkWarpConnectorTunnelConnection]
type warpConnectorGetResponseXlFXheKkWarpConnectorTunnelConnectionJSON struct {
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

func (r *WarpConnectorGetResponseXlFXheKkWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type WarpConnectorGetResponseXlFXheKkWarpConnectorTunnelTunType string

const (
	WarpConnectorGetResponseXlFXheKkWarpConnectorTunnelTunTypeCfdTunnel     WarpConnectorGetResponseXlFXheKkWarpConnectorTunnelTunType = "cfd_tunnel"
	WarpConnectorGetResponseXlFXheKkWarpConnectorTunnelTunTypeWarpConnector WarpConnectorGetResponseXlFXheKkWarpConnectorTunnelTunType = "warp_connector"
	WarpConnectorGetResponseXlFXheKkWarpConnectorTunnelTunTypeIPSec         WarpConnectorGetResponseXlFXheKkWarpConnectorTunnelTunType = "ip_sec"
	WarpConnectorGetResponseXlFXheKkWarpConnectorTunnelTunTypeGre           WarpConnectorGetResponseXlFXheKkWarpConnectorTunnelTunType = "gre"
	WarpConnectorGetResponseXlFXheKkWarpConnectorTunnelTunTypeCni           WarpConnectorGetResponseXlFXheKkWarpConnectorTunnelTunType = "cni"
)

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [WarpConnectorUpdateResponseXlFXheKkCfdTunnel] or
// [WarpConnectorUpdateResponseXlFXheKkWarpConnectorTunnel].
type WarpConnectorUpdateResponse interface {
	implementsWarpConnectorUpdateResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WarpConnectorUpdateResponse)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type WarpConnectorUpdateResponseXlFXheKkCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WarpConnectorUpdateResponseXlFXheKkCfdTunnelConnection `json:"connections"`
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
	TunType WarpConnectorUpdateResponseXlFXheKkCfdTunnelTunType `json:"tun_type"`
	JSON    warpConnectorUpdateResponseXlFXheKkCfdTunnelJSON    `json:"-"`
}

// warpConnectorUpdateResponseXlFXheKkCfdTunnelJSON contains the JSON metadata for
// the struct [WarpConnectorUpdateResponseXlFXheKkCfdTunnel]
type warpConnectorUpdateResponseXlFXheKkCfdTunnelJSON struct {
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

func (r *WarpConnectorUpdateResponseXlFXheKkCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WarpConnectorUpdateResponseXlFXheKkCfdTunnel) implementsWarpConnectorUpdateResponse() {}

type WarpConnectorUpdateResponseXlFXheKkCfdTunnelConnection struct {
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
	Uuid string                                                     `json:"uuid"`
	JSON warpConnectorUpdateResponseXlFXheKkCfdTunnelConnectionJSON `json:"-"`
}

// warpConnectorUpdateResponseXlFXheKkCfdTunnelConnectionJSON contains the JSON
// metadata for the struct [WarpConnectorUpdateResponseXlFXheKkCfdTunnelConnection]
type warpConnectorUpdateResponseXlFXheKkCfdTunnelConnectionJSON struct {
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

func (r *WarpConnectorUpdateResponseXlFXheKkCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type WarpConnectorUpdateResponseXlFXheKkCfdTunnelTunType string

const (
	WarpConnectorUpdateResponseXlFXheKkCfdTunnelTunTypeCfdTunnel     WarpConnectorUpdateResponseXlFXheKkCfdTunnelTunType = "cfd_tunnel"
	WarpConnectorUpdateResponseXlFXheKkCfdTunnelTunTypeWarpConnector WarpConnectorUpdateResponseXlFXheKkCfdTunnelTunType = "warp_connector"
	WarpConnectorUpdateResponseXlFXheKkCfdTunnelTunTypeIPSec         WarpConnectorUpdateResponseXlFXheKkCfdTunnelTunType = "ip_sec"
	WarpConnectorUpdateResponseXlFXheKkCfdTunnelTunTypeGre           WarpConnectorUpdateResponseXlFXheKkCfdTunnelTunType = "gre"
	WarpConnectorUpdateResponseXlFXheKkCfdTunnelTunTypeCni           WarpConnectorUpdateResponseXlFXheKkCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type WarpConnectorUpdateResponseXlFXheKkWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WarpConnectorUpdateResponseXlFXheKkWarpConnectorTunnelConnection `json:"connections"`
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
	TunType WarpConnectorUpdateResponseXlFXheKkWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    warpConnectorUpdateResponseXlFXheKkWarpConnectorTunnelJSON    `json:"-"`
}

// warpConnectorUpdateResponseXlFXheKkWarpConnectorTunnelJSON contains the JSON
// metadata for the struct [WarpConnectorUpdateResponseXlFXheKkWarpConnectorTunnel]
type warpConnectorUpdateResponseXlFXheKkWarpConnectorTunnelJSON struct {
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

func (r *WarpConnectorUpdateResponseXlFXheKkWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WarpConnectorUpdateResponseXlFXheKkWarpConnectorTunnel) implementsWarpConnectorUpdateResponse() {
}

type WarpConnectorUpdateResponseXlFXheKkWarpConnectorTunnelConnection struct {
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
	Uuid string                                                               `json:"uuid"`
	JSON warpConnectorUpdateResponseXlFXheKkWarpConnectorTunnelConnectionJSON `json:"-"`
}

// warpConnectorUpdateResponseXlFXheKkWarpConnectorTunnelConnectionJSON contains
// the JSON metadata for the struct
// [WarpConnectorUpdateResponseXlFXheKkWarpConnectorTunnelConnection]
type warpConnectorUpdateResponseXlFXheKkWarpConnectorTunnelConnectionJSON struct {
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

func (r *WarpConnectorUpdateResponseXlFXheKkWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type WarpConnectorUpdateResponseXlFXheKkWarpConnectorTunnelTunType string

const (
	WarpConnectorUpdateResponseXlFXheKkWarpConnectorTunnelTunTypeCfdTunnel     WarpConnectorUpdateResponseXlFXheKkWarpConnectorTunnelTunType = "cfd_tunnel"
	WarpConnectorUpdateResponseXlFXheKkWarpConnectorTunnelTunTypeWarpConnector WarpConnectorUpdateResponseXlFXheKkWarpConnectorTunnelTunType = "warp_connector"
	WarpConnectorUpdateResponseXlFXheKkWarpConnectorTunnelTunTypeIPSec         WarpConnectorUpdateResponseXlFXheKkWarpConnectorTunnelTunType = "ip_sec"
	WarpConnectorUpdateResponseXlFXheKkWarpConnectorTunnelTunTypeGre           WarpConnectorUpdateResponseXlFXheKkWarpConnectorTunnelTunType = "gre"
	WarpConnectorUpdateResponseXlFXheKkWarpConnectorTunnelTunTypeCni           WarpConnectorUpdateResponseXlFXheKkWarpConnectorTunnelTunType = "cni"
)

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [WarpConnectorListResponseXlFXheKkCfdTunnel] or
// [WarpConnectorListResponseXlFXheKkWarpConnectorTunnel].
type WarpConnectorListResponse interface {
	implementsWarpConnectorListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WarpConnectorListResponse)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type WarpConnectorListResponseXlFXheKkCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WarpConnectorListResponseXlFXheKkCfdTunnelConnection `json:"connections"`
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
	TunType WarpConnectorListResponseXlFXheKkCfdTunnelTunType `json:"tun_type"`
	JSON    warpConnectorListResponseXlFXheKkCfdTunnelJSON    `json:"-"`
}

// warpConnectorListResponseXlFXheKkCfdTunnelJSON contains the JSON metadata for
// the struct [WarpConnectorListResponseXlFXheKkCfdTunnel]
type warpConnectorListResponseXlFXheKkCfdTunnelJSON struct {
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

func (r *WarpConnectorListResponseXlFXheKkCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WarpConnectorListResponseXlFXheKkCfdTunnel) implementsWarpConnectorListResponse() {}

type WarpConnectorListResponseXlFXheKkCfdTunnelConnection struct {
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
	Uuid string                                                   `json:"uuid"`
	JSON warpConnectorListResponseXlFXheKkCfdTunnelConnectionJSON `json:"-"`
}

// warpConnectorListResponseXlFXheKkCfdTunnelConnectionJSON contains the JSON
// metadata for the struct [WarpConnectorListResponseXlFXheKkCfdTunnelConnection]
type warpConnectorListResponseXlFXheKkCfdTunnelConnectionJSON struct {
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

func (r *WarpConnectorListResponseXlFXheKkCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type WarpConnectorListResponseXlFXheKkCfdTunnelTunType string

const (
	WarpConnectorListResponseXlFXheKkCfdTunnelTunTypeCfdTunnel     WarpConnectorListResponseXlFXheKkCfdTunnelTunType = "cfd_tunnel"
	WarpConnectorListResponseXlFXheKkCfdTunnelTunTypeWarpConnector WarpConnectorListResponseXlFXheKkCfdTunnelTunType = "warp_connector"
	WarpConnectorListResponseXlFXheKkCfdTunnelTunTypeIPSec         WarpConnectorListResponseXlFXheKkCfdTunnelTunType = "ip_sec"
	WarpConnectorListResponseXlFXheKkCfdTunnelTunTypeGre           WarpConnectorListResponseXlFXheKkCfdTunnelTunType = "gre"
	WarpConnectorListResponseXlFXheKkCfdTunnelTunTypeCni           WarpConnectorListResponseXlFXheKkCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type WarpConnectorListResponseXlFXheKkWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WarpConnectorListResponseXlFXheKkWarpConnectorTunnelConnection `json:"connections"`
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
	TunType WarpConnectorListResponseXlFXheKkWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    warpConnectorListResponseXlFXheKkWarpConnectorTunnelJSON    `json:"-"`
}

// warpConnectorListResponseXlFXheKkWarpConnectorTunnelJSON contains the JSON
// metadata for the struct [WarpConnectorListResponseXlFXheKkWarpConnectorTunnel]
type warpConnectorListResponseXlFXheKkWarpConnectorTunnelJSON struct {
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

func (r *WarpConnectorListResponseXlFXheKkWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WarpConnectorListResponseXlFXheKkWarpConnectorTunnel) implementsWarpConnectorListResponse() {}

type WarpConnectorListResponseXlFXheKkWarpConnectorTunnelConnection struct {
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
	Uuid string                                                             `json:"uuid"`
	JSON warpConnectorListResponseXlFXheKkWarpConnectorTunnelConnectionJSON `json:"-"`
}

// warpConnectorListResponseXlFXheKkWarpConnectorTunnelConnectionJSON contains the
// JSON metadata for the struct
// [WarpConnectorListResponseXlFXheKkWarpConnectorTunnelConnection]
type warpConnectorListResponseXlFXheKkWarpConnectorTunnelConnectionJSON struct {
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

func (r *WarpConnectorListResponseXlFXheKkWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type WarpConnectorListResponseXlFXheKkWarpConnectorTunnelTunType string

const (
	WarpConnectorListResponseXlFXheKkWarpConnectorTunnelTunTypeCfdTunnel     WarpConnectorListResponseXlFXheKkWarpConnectorTunnelTunType = "cfd_tunnel"
	WarpConnectorListResponseXlFXheKkWarpConnectorTunnelTunTypeWarpConnector WarpConnectorListResponseXlFXheKkWarpConnectorTunnelTunType = "warp_connector"
	WarpConnectorListResponseXlFXheKkWarpConnectorTunnelTunTypeIPSec         WarpConnectorListResponseXlFXheKkWarpConnectorTunnelTunType = "ip_sec"
	WarpConnectorListResponseXlFXheKkWarpConnectorTunnelTunTypeGre           WarpConnectorListResponseXlFXheKkWarpConnectorTunnelTunType = "gre"
	WarpConnectorListResponseXlFXheKkWarpConnectorTunnelTunTypeCni           WarpConnectorListResponseXlFXheKkWarpConnectorTunnelTunType = "cni"
)

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [WarpConnectorDeleteResponseXlFXheKkCfdTunnel] or
// [WarpConnectorDeleteResponseXlFXheKkWarpConnectorTunnel].
type WarpConnectorDeleteResponse interface {
	implementsWarpConnectorDeleteResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WarpConnectorDeleteResponse)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type WarpConnectorDeleteResponseXlFXheKkCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WarpConnectorDeleteResponseXlFXheKkCfdTunnelConnection `json:"connections"`
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
	TunType WarpConnectorDeleteResponseXlFXheKkCfdTunnelTunType `json:"tun_type"`
	JSON    warpConnectorDeleteResponseXlFXheKkCfdTunnelJSON    `json:"-"`
}

// warpConnectorDeleteResponseXlFXheKkCfdTunnelJSON contains the JSON metadata for
// the struct [WarpConnectorDeleteResponseXlFXheKkCfdTunnel]
type warpConnectorDeleteResponseXlFXheKkCfdTunnelJSON struct {
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

func (r *WarpConnectorDeleteResponseXlFXheKkCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WarpConnectorDeleteResponseXlFXheKkCfdTunnel) implementsWarpConnectorDeleteResponse() {}

type WarpConnectorDeleteResponseXlFXheKkCfdTunnelConnection struct {
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
	Uuid string                                                     `json:"uuid"`
	JSON warpConnectorDeleteResponseXlFXheKkCfdTunnelConnectionJSON `json:"-"`
}

// warpConnectorDeleteResponseXlFXheKkCfdTunnelConnectionJSON contains the JSON
// metadata for the struct [WarpConnectorDeleteResponseXlFXheKkCfdTunnelConnection]
type warpConnectorDeleteResponseXlFXheKkCfdTunnelConnectionJSON struct {
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

func (r *WarpConnectorDeleteResponseXlFXheKkCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type WarpConnectorDeleteResponseXlFXheKkCfdTunnelTunType string

const (
	WarpConnectorDeleteResponseXlFXheKkCfdTunnelTunTypeCfdTunnel     WarpConnectorDeleteResponseXlFXheKkCfdTunnelTunType = "cfd_tunnel"
	WarpConnectorDeleteResponseXlFXheKkCfdTunnelTunTypeWarpConnector WarpConnectorDeleteResponseXlFXheKkCfdTunnelTunType = "warp_connector"
	WarpConnectorDeleteResponseXlFXheKkCfdTunnelTunTypeIPSec         WarpConnectorDeleteResponseXlFXheKkCfdTunnelTunType = "ip_sec"
	WarpConnectorDeleteResponseXlFXheKkCfdTunnelTunTypeGre           WarpConnectorDeleteResponseXlFXheKkCfdTunnelTunType = "gre"
	WarpConnectorDeleteResponseXlFXheKkCfdTunnelTunTypeCni           WarpConnectorDeleteResponseXlFXheKkCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type WarpConnectorDeleteResponseXlFXheKkWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WarpConnectorDeleteResponseXlFXheKkWarpConnectorTunnelConnection `json:"connections"`
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
	TunType WarpConnectorDeleteResponseXlFXheKkWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    warpConnectorDeleteResponseXlFXheKkWarpConnectorTunnelJSON    `json:"-"`
}

// warpConnectorDeleteResponseXlFXheKkWarpConnectorTunnelJSON contains the JSON
// metadata for the struct [WarpConnectorDeleteResponseXlFXheKkWarpConnectorTunnel]
type warpConnectorDeleteResponseXlFXheKkWarpConnectorTunnelJSON struct {
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

func (r *WarpConnectorDeleteResponseXlFXheKkWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WarpConnectorDeleteResponseXlFXheKkWarpConnectorTunnel) implementsWarpConnectorDeleteResponse() {
}

type WarpConnectorDeleteResponseXlFXheKkWarpConnectorTunnelConnection struct {
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
	Uuid string                                                               `json:"uuid"`
	JSON warpConnectorDeleteResponseXlFXheKkWarpConnectorTunnelConnectionJSON `json:"-"`
}

// warpConnectorDeleteResponseXlFXheKkWarpConnectorTunnelConnectionJSON contains
// the JSON metadata for the struct
// [WarpConnectorDeleteResponseXlFXheKkWarpConnectorTunnelConnection]
type warpConnectorDeleteResponseXlFXheKkWarpConnectorTunnelConnectionJSON struct {
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

func (r *WarpConnectorDeleteResponseXlFXheKkWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type WarpConnectorDeleteResponseXlFXheKkWarpConnectorTunnelTunType string

const (
	WarpConnectorDeleteResponseXlFXheKkWarpConnectorTunnelTunTypeCfdTunnel     WarpConnectorDeleteResponseXlFXheKkWarpConnectorTunnelTunType = "cfd_tunnel"
	WarpConnectorDeleteResponseXlFXheKkWarpConnectorTunnelTunTypeWarpConnector WarpConnectorDeleteResponseXlFXheKkWarpConnectorTunnelTunType = "warp_connector"
	WarpConnectorDeleteResponseXlFXheKkWarpConnectorTunnelTunTypeIPSec         WarpConnectorDeleteResponseXlFXheKkWarpConnectorTunnelTunType = "ip_sec"
	WarpConnectorDeleteResponseXlFXheKkWarpConnectorTunnelTunTypeGre           WarpConnectorDeleteResponseXlFXheKkWarpConnectorTunnelTunType = "gre"
	WarpConnectorDeleteResponseXlFXheKkWarpConnectorTunnelTunTypeCni           WarpConnectorDeleteResponseXlFXheKkWarpConnectorTunnelTunType = "cni"
)

type WarpConnectorNewParams struct {
	// A user-friendly name for the tunnel.
	Name param.Field[string] `json:"name,required"`
}

func (r WarpConnectorNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WarpConnectorNewResponseEnvelope struct {
	Errors   []WarpConnectorNewResponseEnvelopeErrors   `json:"errors"`
	Messages []WarpConnectorNewResponseEnvelopeMessages `json:"messages"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result WarpConnectorNewResponse `json:"result"`
	// Whether the API call was successful
	Success WarpConnectorNewResponseEnvelopeSuccess `json:"success"`
	JSON    warpConnectorNewResponseEnvelopeJSON    `json:"-"`
}

// warpConnectorNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [WarpConnectorNewResponseEnvelope]
type warpConnectorNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WarpConnectorNewResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    warpConnectorNewResponseEnvelopeErrorsJSON `json:"-"`
}

// warpConnectorNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WarpConnectorNewResponseEnvelopeErrors]
type warpConnectorNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WarpConnectorNewResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    warpConnectorNewResponseEnvelopeMessagesJSON `json:"-"`
}

// warpConnectorNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WarpConnectorNewResponseEnvelopeMessages]
type warpConnectorNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WarpConnectorNewResponseEnvelopeSuccess bool

const (
	WarpConnectorNewResponseEnvelopeSuccessTrue WarpConnectorNewResponseEnvelopeSuccess = true
)

type WarpConnectorGetResponseEnvelope struct {
	Errors   []WarpConnectorGetResponseEnvelopeErrors   `json:"errors"`
	Messages []WarpConnectorGetResponseEnvelopeMessages `json:"messages"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result WarpConnectorGetResponse `json:"result"`
	// Whether the API call was successful
	Success WarpConnectorGetResponseEnvelopeSuccess `json:"success"`
	JSON    warpConnectorGetResponseEnvelopeJSON    `json:"-"`
}

// warpConnectorGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [WarpConnectorGetResponseEnvelope]
type warpConnectorGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WarpConnectorGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    warpConnectorGetResponseEnvelopeErrorsJSON `json:"-"`
}

// warpConnectorGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WarpConnectorGetResponseEnvelopeErrors]
type warpConnectorGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WarpConnectorGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    warpConnectorGetResponseEnvelopeMessagesJSON `json:"-"`
}

// warpConnectorGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WarpConnectorGetResponseEnvelopeMessages]
type warpConnectorGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WarpConnectorGetResponseEnvelopeSuccess bool

const (
	WarpConnectorGetResponseEnvelopeSuccessTrue WarpConnectorGetResponseEnvelopeSuccess = true
)

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

type WarpConnectorUpdateResponseEnvelope struct {
	Errors   []WarpConnectorUpdateResponseEnvelopeErrors   `json:"errors"`
	Messages []WarpConnectorUpdateResponseEnvelopeMessages `json:"messages"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result WarpConnectorUpdateResponse `json:"result"`
	// Whether the API call was successful
	Success WarpConnectorUpdateResponseEnvelopeSuccess `json:"success"`
	JSON    warpConnectorUpdateResponseEnvelopeJSON    `json:"-"`
}

// warpConnectorUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [WarpConnectorUpdateResponseEnvelope]
type warpConnectorUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WarpConnectorUpdateResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    warpConnectorUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// warpConnectorUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WarpConnectorUpdateResponseEnvelopeErrors]
type warpConnectorUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WarpConnectorUpdateResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    warpConnectorUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// warpConnectorUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [WarpConnectorUpdateResponseEnvelopeMessages]
type warpConnectorUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WarpConnectorUpdateResponseEnvelopeSuccess bool

const (
	WarpConnectorUpdateResponseEnvelopeSuccessTrue WarpConnectorUpdateResponseEnvelopeSuccess = true
)

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

type WarpConnectorListResponseEnvelope struct {
	Errors     []WarpConnectorListResponseEnvelopeErrors   `json:"errors"`
	Messages   []WarpConnectorListResponseEnvelopeMessages `json:"messages"`
	Result     []WarpConnectorListResponse                 `json:"result"`
	ResultInfo WarpConnectorListResponseEnvelopeResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success WarpConnectorListResponseEnvelopeSuccess `json:"success"`
	JSON    warpConnectorListResponseEnvelopeJSON    `json:"-"`
}

// warpConnectorListResponseEnvelopeJSON contains the JSON metadata for the struct
// [WarpConnectorListResponseEnvelope]
type warpConnectorListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WarpConnectorListResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    warpConnectorListResponseEnvelopeErrorsJSON `json:"-"`
}

// warpConnectorListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WarpConnectorListResponseEnvelopeErrors]
type warpConnectorListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WarpConnectorListResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    warpConnectorListResponseEnvelopeMessagesJSON `json:"-"`
}

// warpConnectorListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WarpConnectorListResponseEnvelopeMessages]
type warpConnectorListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WarpConnectorListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                         `json:"total_count"`
	JSON       warpConnectorListResponseEnvelopeResultInfoJSON `json:"-"`
}

// warpConnectorListResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [WarpConnectorListResponseEnvelopeResultInfo]
type warpConnectorListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WarpConnectorListResponseEnvelopeSuccess bool

const (
	WarpConnectorListResponseEnvelopeSuccessTrue WarpConnectorListResponseEnvelopeSuccess = true
)

type WarpConnectorDeleteParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r WarpConnectorDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type WarpConnectorDeleteResponseEnvelope struct {
	Errors   []WarpConnectorDeleteResponseEnvelopeErrors   `json:"errors"`
	Messages []WarpConnectorDeleteResponseEnvelopeMessages `json:"messages"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result WarpConnectorDeleteResponse `json:"result"`
	// Whether the API call was successful
	Success WarpConnectorDeleteResponseEnvelopeSuccess `json:"success"`
	JSON    warpConnectorDeleteResponseEnvelopeJSON    `json:"-"`
}

// warpConnectorDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [WarpConnectorDeleteResponseEnvelope]
type warpConnectorDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WarpConnectorDeleteResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    warpConnectorDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// warpConnectorDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WarpConnectorDeleteResponseEnvelopeErrors]
type warpConnectorDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WarpConnectorDeleteResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    warpConnectorDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// warpConnectorDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [WarpConnectorDeleteResponseEnvelopeMessages]
type warpConnectorDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WarpConnectorDeleteResponseEnvelopeSuccess bool

const (
	WarpConnectorDeleteResponseEnvelopeSuccessTrue WarpConnectorDeleteResponseEnvelopeSuccess = true
)
