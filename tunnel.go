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
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// TunnelService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewTunnelService] method instead.
type TunnelService struct {
	Options        []option.RequestOption
	Configurations *TunnelConfigurationService
	Connections    *TunnelConnectionService
	Tokens         *TunnelTokenService
	Connectors     *TunnelConnectorService
	Management     *TunnelManagementService
}

// NewTunnelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTunnelService(opts ...option.RequestOption) (r *TunnelService) {
	r = &TunnelService{}
	r.Options = opts
	r.Configurations = NewTunnelConfigurationService(opts...)
	r.Connections = NewTunnelConnectionService(opts...)
	r.Tokens = NewTunnelTokenService(opts...)
	r.Connectors = NewTunnelConnectorService(opts...)
	r.Management = NewTunnelManagementService(opts...)
	return
}

// Creates a new Argo Tunnel in an account.
func (r *TunnelService) New(ctx context.Context, accountID string, body TunnelNewParams, opts ...option.RequestOption) (res *TunnelNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TunnelNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/tunnels", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists and filters all types of Tunnels in an account.
func (r *TunnelService) List(ctx context.Context, accountID string, query TunnelListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[TunnelListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/tunnels", accountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists and filters all types of Tunnels in an account.
func (r *TunnelService) ListAutoPaging(ctx context.Context, accountID string, query TunnelListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[TunnelListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, accountID, query, opts...))
}

// Deletes an Argo Tunnel from an account.
func (r *TunnelService) Delete(ctx context.Context, accountID string, tunnelID string, body TunnelDeleteParams, opts ...option.RequestOption) (res *TunnelDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TunnelDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/tunnels/%s", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing Cloudflare Tunnel.
func (r *TunnelService) Edit(ctx context.Context, accountID string, tunnelID string, body TunnelEditParams, opts ...option.RequestOption) (res *TunnelEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TunnelEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Argo Tunnel.
func (r *TunnelService) Get(ctx context.Context, accountID string, tunnelID string, opts ...option.RequestOption) (res *TunnelGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TunnelGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/tunnels/%s", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TunnelNewResponse struct {
	// UUID of the tunnel.
	ID string `json:"id,required"`
	// The tunnel connections between your origin and Cloudflare's edge.
	Connections []TunnelNewResponseConnection `json:"connections,required"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A user-friendly name for the tunnel.
	Name string `json:"name,required"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time             `json:"deleted_at,nullable" format:"date-time"`
	JSON      tunnelNewResponseJSON `json:"-"`
}

// tunnelNewResponseJSON contains the JSON metadata for the struct
// [TunnelNewResponse]
type tunnelNewResponseJSON struct {
	ID          apijson.Field
	Connections apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	DeletedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelNewResponseConnection struct {
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                          `json:"uuid"`
	JSON tunnelNewResponseConnectionJSON `json:"-"`
}

// tunnelNewResponseConnectionJSON contains the JSON metadata for the struct
// [TunnelNewResponseConnection]
type tunnelNewResponseConnectionJSON struct {
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TunnelNewResponseConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [TunnelListResponseTunnelCfdTunnel] or
// [TunnelListResponseTunnelWarpConnectorTunnel].
type TunnelListResponse interface {
	implementsTunnelListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*TunnelListResponse)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type TunnelListResponseTunnelCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []TunnelListResponseTunnelCfdTunnelConnection `json:"connections"`
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
	TunType TunnelListResponseTunnelCfdTunnelTunType `json:"tun_type"`
	JSON    tunnelListResponseTunnelCfdTunnelJSON    `json:"-"`
}

// tunnelListResponseTunnelCfdTunnelJSON contains the JSON metadata for the struct
// [TunnelListResponseTunnelCfdTunnel]
type tunnelListResponseTunnelCfdTunnelJSON struct {
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

func (r *TunnelListResponseTunnelCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r TunnelListResponseTunnelCfdTunnel) implementsTunnelListResponse() {}

type TunnelListResponseTunnelCfdTunnelConnection struct {
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
	Uuid string                                          `json:"uuid"`
	JSON tunnelListResponseTunnelCfdTunnelConnectionJSON `json:"-"`
}

// tunnelListResponseTunnelCfdTunnelConnectionJSON contains the JSON metadata for
// the struct [TunnelListResponseTunnelCfdTunnelConnection]
type tunnelListResponseTunnelCfdTunnelConnectionJSON struct {
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

func (r *TunnelListResponseTunnelCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type TunnelListResponseTunnelCfdTunnelTunType string

const (
	TunnelListResponseTunnelCfdTunnelTunTypeCfdTunnel     TunnelListResponseTunnelCfdTunnelTunType = "cfd_tunnel"
	TunnelListResponseTunnelCfdTunnelTunTypeWarpConnector TunnelListResponseTunnelCfdTunnelTunType = "warp_connector"
	TunnelListResponseTunnelCfdTunnelTunTypeIPSec         TunnelListResponseTunnelCfdTunnelTunType = "ip_sec"
	TunnelListResponseTunnelCfdTunnelTunTypeGre           TunnelListResponseTunnelCfdTunnelTunType = "gre"
	TunnelListResponseTunnelCfdTunnelTunTypeCni           TunnelListResponseTunnelCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type TunnelListResponseTunnelWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []TunnelListResponseTunnelWarpConnectorTunnelConnection `json:"connections"`
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
	TunType TunnelListResponseTunnelWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    tunnelListResponseTunnelWarpConnectorTunnelJSON    `json:"-"`
}

// tunnelListResponseTunnelWarpConnectorTunnelJSON contains the JSON metadata for
// the struct [TunnelListResponseTunnelWarpConnectorTunnel]
type tunnelListResponseTunnelWarpConnectorTunnelJSON struct {
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

func (r *TunnelListResponseTunnelWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r TunnelListResponseTunnelWarpConnectorTunnel) implementsTunnelListResponse() {}

type TunnelListResponseTunnelWarpConnectorTunnelConnection struct {
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
	Uuid string                                                    `json:"uuid"`
	JSON tunnelListResponseTunnelWarpConnectorTunnelConnectionJSON `json:"-"`
}

// tunnelListResponseTunnelWarpConnectorTunnelConnectionJSON contains the JSON
// metadata for the struct [TunnelListResponseTunnelWarpConnectorTunnelConnection]
type tunnelListResponseTunnelWarpConnectorTunnelConnectionJSON struct {
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

func (r *TunnelListResponseTunnelWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type TunnelListResponseTunnelWarpConnectorTunnelTunType string

const (
	TunnelListResponseTunnelWarpConnectorTunnelTunTypeCfdTunnel     TunnelListResponseTunnelWarpConnectorTunnelTunType = "cfd_tunnel"
	TunnelListResponseTunnelWarpConnectorTunnelTunTypeWarpConnector TunnelListResponseTunnelWarpConnectorTunnelTunType = "warp_connector"
	TunnelListResponseTunnelWarpConnectorTunnelTunTypeIPSec         TunnelListResponseTunnelWarpConnectorTunnelTunType = "ip_sec"
	TunnelListResponseTunnelWarpConnectorTunnelTunTypeGre           TunnelListResponseTunnelWarpConnectorTunnelTunType = "gre"
	TunnelListResponseTunnelWarpConnectorTunnelTunTypeCni           TunnelListResponseTunnelWarpConnectorTunnelTunType = "cni"
)

type TunnelDeleteResponse struct {
	// UUID of the tunnel.
	ID string `json:"id,required"`
	// The tunnel connections between your origin and Cloudflare's edge.
	Connections []TunnelDeleteResponseConnection `json:"connections,required"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A user-friendly name for the tunnel.
	Name string `json:"name,required"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time                `json:"deleted_at,nullable" format:"date-time"`
	JSON      tunnelDeleteResponseJSON `json:"-"`
}

// tunnelDeleteResponseJSON contains the JSON metadata for the struct
// [TunnelDeleteResponse]
type tunnelDeleteResponseJSON struct {
	ID          apijson.Field
	Connections apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	DeletedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelDeleteResponseConnection struct {
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                             `json:"uuid"`
	JSON tunnelDeleteResponseConnectionJSON `json:"-"`
}

// tunnelDeleteResponseConnectionJSON contains the JSON metadata for the struct
// [TunnelDeleteResponseConnection]
type tunnelDeleteResponseConnectionJSON struct {
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TunnelDeleteResponseConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [TunnelEditResponseTunnelCfdTunnel] or
// [TunnelEditResponseTunnelWarpConnectorTunnel].
type TunnelEditResponse interface {
	implementsTunnelEditResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*TunnelEditResponse)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type TunnelEditResponseTunnelCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []TunnelEditResponseTunnelCfdTunnelConnection `json:"connections"`
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
	TunType TunnelEditResponseTunnelCfdTunnelTunType `json:"tun_type"`
	JSON    tunnelEditResponseTunnelCfdTunnelJSON    `json:"-"`
}

// tunnelEditResponseTunnelCfdTunnelJSON contains the JSON metadata for the struct
// [TunnelEditResponseTunnelCfdTunnel]
type tunnelEditResponseTunnelCfdTunnelJSON struct {
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

func (r *TunnelEditResponseTunnelCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r TunnelEditResponseTunnelCfdTunnel) implementsTunnelEditResponse() {}

type TunnelEditResponseTunnelCfdTunnelConnection struct {
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
	Uuid string                                          `json:"uuid"`
	JSON tunnelEditResponseTunnelCfdTunnelConnectionJSON `json:"-"`
}

// tunnelEditResponseTunnelCfdTunnelConnectionJSON contains the JSON metadata for
// the struct [TunnelEditResponseTunnelCfdTunnelConnection]
type tunnelEditResponseTunnelCfdTunnelConnectionJSON struct {
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

func (r *TunnelEditResponseTunnelCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type TunnelEditResponseTunnelCfdTunnelTunType string

const (
	TunnelEditResponseTunnelCfdTunnelTunTypeCfdTunnel     TunnelEditResponseTunnelCfdTunnelTunType = "cfd_tunnel"
	TunnelEditResponseTunnelCfdTunnelTunTypeWarpConnector TunnelEditResponseTunnelCfdTunnelTunType = "warp_connector"
	TunnelEditResponseTunnelCfdTunnelTunTypeIPSec         TunnelEditResponseTunnelCfdTunnelTunType = "ip_sec"
	TunnelEditResponseTunnelCfdTunnelTunTypeGre           TunnelEditResponseTunnelCfdTunnelTunType = "gre"
	TunnelEditResponseTunnelCfdTunnelTunTypeCni           TunnelEditResponseTunnelCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type TunnelEditResponseTunnelWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []TunnelEditResponseTunnelWarpConnectorTunnelConnection `json:"connections"`
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
	TunType TunnelEditResponseTunnelWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    tunnelEditResponseTunnelWarpConnectorTunnelJSON    `json:"-"`
}

// tunnelEditResponseTunnelWarpConnectorTunnelJSON contains the JSON metadata for
// the struct [TunnelEditResponseTunnelWarpConnectorTunnel]
type tunnelEditResponseTunnelWarpConnectorTunnelJSON struct {
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

func (r *TunnelEditResponseTunnelWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r TunnelEditResponseTunnelWarpConnectorTunnel) implementsTunnelEditResponse() {}

type TunnelEditResponseTunnelWarpConnectorTunnelConnection struct {
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
	Uuid string                                                    `json:"uuid"`
	JSON tunnelEditResponseTunnelWarpConnectorTunnelConnectionJSON `json:"-"`
}

// tunnelEditResponseTunnelWarpConnectorTunnelConnectionJSON contains the JSON
// metadata for the struct [TunnelEditResponseTunnelWarpConnectorTunnelConnection]
type tunnelEditResponseTunnelWarpConnectorTunnelConnectionJSON struct {
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

func (r *TunnelEditResponseTunnelWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type TunnelEditResponseTunnelWarpConnectorTunnelTunType string

const (
	TunnelEditResponseTunnelWarpConnectorTunnelTunTypeCfdTunnel     TunnelEditResponseTunnelWarpConnectorTunnelTunType = "cfd_tunnel"
	TunnelEditResponseTunnelWarpConnectorTunnelTunTypeWarpConnector TunnelEditResponseTunnelWarpConnectorTunnelTunType = "warp_connector"
	TunnelEditResponseTunnelWarpConnectorTunnelTunTypeIPSec         TunnelEditResponseTunnelWarpConnectorTunnelTunType = "ip_sec"
	TunnelEditResponseTunnelWarpConnectorTunnelTunTypeGre           TunnelEditResponseTunnelWarpConnectorTunnelTunType = "gre"
	TunnelEditResponseTunnelWarpConnectorTunnelTunTypeCni           TunnelEditResponseTunnelWarpConnectorTunnelTunType = "cni"
)

type TunnelGetResponse struct {
	// UUID of the tunnel.
	ID string `json:"id,required"`
	// The tunnel connections between your origin and Cloudflare's edge.
	Connections []TunnelGetResponseConnection `json:"connections,required"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A user-friendly name for the tunnel.
	Name string `json:"name,required"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time             `json:"deleted_at,nullable" format:"date-time"`
	JSON      tunnelGetResponseJSON `json:"-"`
}

// tunnelGetResponseJSON contains the JSON metadata for the struct
// [TunnelGetResponse]
type tunnelGetResponseJSON struct {
	ID          apijson.Field
	Connections apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	DeletedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelGetResponseConnection struct {
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                          `json:"uuid"`
	JSON tunnelGetResponseConnectionJSON `json:"-"`
}

// tunnelGetResponseConnectionJSON contains the JSON metadata for the struct
// [TunnelGetResponseConnection]
type tunnelGetResponseConnectionJSON struct {
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TunnelGetResponseConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelNewParams struct {
	// A user-friendly name for the tunnel.
	Name param.Field[string] `json:"name,required"`
	// Sets the password required to run the tunnel. Must be at least 32 bytes and
	// encoded as a base64 string.
	TunnelSecret param.Field[interface{}] `json:"tunnel_secret,required"`
}

func (r TunnelNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TunnelNewResponseEnvelope struct {
	Errors   []TunnelNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TunnelNewResponseEnvelopeMessages `json:"messages,required"`
	Result   TunnelNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success TunnelNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelNewResponseEnvelopeJSON    `json:"-"`
}

// tunnelNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [TunnelNewResponseEnvelope]
type tunnelNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelNewResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    tunnelNewResponseEnvelopeErrorsJSON `json:"-"`
}

// tunnelNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [TunnelNewResponseEnvelopeErrors]
type tunnelNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelNewResponseEnvelopeMessages struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    tunnelNewResponseEnvelopeMessagesJSON `json:"-"`
}

// tunnelNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [TunnelNewResponseEnvelopeMessages]
type tunnelNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TunnelNewResponseEnvelopeSuccess bool

const (
	TunnelNewResponseEnvelopeSuccessTrue TunnelNewResponseEnvelopeSuccess = true
)

type TunnelListParams struct {
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
	PerPage param.Field[float64] `query:"per_page"`
	// The types of tunnels to filter separated by a comma.
	TunTypes      param.Field[string]    `query:"tun_types"`
	WasActiveAt   param.Field[time.Time] `query:"was_active_at" format:"date-time"`
	WasInactiveAt param.Field[time.Time] `query:"was_inactive_at" format:"date-time"`
}

// URLQuery serializes [TunnelListParams]'s query parameters as `url.Values`.
func (r TunnelListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TunnelDeleteParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r TunnelDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type TunnelDeleteResponseEnvelope struct {
	Errors   []TunnelDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TunnelDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   TunnelDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success TunnelDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelDeleteResponseEnvelopeJSON    `json:"-"`
}

// tunnelDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [TunnelDeleteResponseEnvelope]
type tunnelDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelDeleteResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    tunnelDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// tunnelDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [TunnelDeleteResponseEnvelopeErrors]
type tunnelDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelDeleteResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    tunnelDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// tunnelDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [TunnelDeleteResponseEnvelopeMessages]
type tunnelDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TunnelDeleteResponseEnvelopeSuccess bool

const (
	TunnelDeleteResponseEnvelopeSuccessTrue TunnelDeleteResponseEnvelopeSuccess = true
)

type TunnelEditParams struct {
	// A user-friendly name for the tunnel.
	Name param.Field[string] `json:"name"`
	// Sets the password required to run a locally-managed tunnel. Must be at least 32
	// bytes and encoded as a base64 string.
	TunnelSecret param.Field[string] `json:"tunnel_secret"`
}

func (r TunnelEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TunnelEditResponseEnvelope struct {
	Errors   []TunnelEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TunnelEditResponseEnvelopeMessages `json:"messages,required"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result TunnelEditResponse `json:"result,required"`
	// Whether the API call was successful
	Success TunnelEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelEditResponseEnvelopeJSON    `json:"-"`
}

// tunnelEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [TunnelEditResponseEnvelope]
type tunnelEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelEditResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    tunnelEditResponseEnvelopeErrorsJSON `json:"-"`
}

// tunnelEditResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [TunnelEditResponseEnvelopeErrors]
type tunnelEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelEditResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    tunnelEditResponseEnvelopeMessagesJSON `json:"-"`
}

// tunnelEditResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [TunnelEditResponseEnvelopeMessages]
type tunnelEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TunnelEditResponseEnvelopeSuccess bool

const (
	TunnelEditResponseEnvelopeSuccessTrue TunnelEditResponseEnvelopeSuccess = true
)

type TunnelGetResponseEnvelope struct {
	Errors   []TunnelGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TunnelGetResponseEnvelopeMessages `json:"messages,required"`
	Result   TunnelGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success TunnelGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelGetResponseEnvelopeJSON    `json:"-"`
}

// tunnelGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [TunnelGetResponseEnvelope]
type tunnelGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelGetResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    tunnelGetResponseEnvelopeErrorsJSON `json:"-"`
}

// tunnelGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [TunnelGetResponseEnvelopeErrors]
type tunnelGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelGetResponseEnvelopeMessages struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    tunnelGetResponseEnvelopeMessagesJSON `json:"-"`
}

// tunnelGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [TunnelGetResponseEnvelopeMessages]
type tunnelGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TunnelGetResponseEnvelopeSuccess bool

const (
	TunnelGetResponseEnvelopeSuccessTrue TunnelGetResponseEnvelopeSuccess = true
)
