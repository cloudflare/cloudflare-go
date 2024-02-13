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

// CfdTunnelService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewCfdTunnelService] method instead.
type CfdTunnelService struct {
	Options        []option.RequestOption
	Configurations *CfdTunnelConfigurationService
	Connections    *CfdTunnelConnectionService
	Tokens         *CfdTunnelTokenService
	Connectors     *CfdTunnelConnectorService
	Management     *CfdTunnelManagementService
}

// NewCfdTunnelService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCfdTunnelService(opts ...option.RequestOption) (r *CfdTunnelService) {
	r = &CfdTunnelService{}
	r.Options = opts
	r.Configurations = NewCfdTunnelConfigurationService(opts...)
	r.Connections = NewCfdTunnelConnectionService(opts...)
	r.Tokens = NewCfdTunnelTokenService(opts...)
	r.Connectors = NewCfdTunnelConnectorService(opts...)
	r.Management = NewCfdTunnelManagementService(opts...)
	return
}

// Updates an existing Cloudflare Tunnel.
func (r *CfdTunnelService) Update(ctx context.Context, accountID string, tunnelID string, body CfdTunnelUpdateParams, opts ...option.RequestOption) (res *CfdTunnelUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CfdTunnelUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a Cloudflare Tunnel from an account.
func (r *CfdTunnelService) Delete(ctx context.Context, accountID string, tunnelID string, body CfdTunnelDeleteParams, opts ...option.RequestOption) (res *CfdTunnelDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CfdTunnelDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Creates a new Cloudflare Tunnel in an account.
func (r *CfdTunnelService) CloudflareTunnelNewACloudflareTunnel(ctx context.Context, accountID string, body CfdTunnelCloudflareTunnelNewACloudflareTunnelParams, opts ...option.RequestOption) (res *CfdTunnelCloudflareTunnelNewACloudflareTunnelResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cfd_tunnel", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists and filters Cloudflare Tunnels in an account.
func (r *CfdTunnelService) CloudflareTunnelListCloudflareTunnels(ctx context.Context, accountID string, query CfdTunnelCloudflareTunnelListCloudflareTunnelsParams, opts ...option.RequestOption) (res *[]CfdTunnelCloudflareTunnelListCloudflareTunnelsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cfd_tunnel", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Cloudflare Tunnel.
func (r *CfdTunnelService) Get(ctx context.Context, accountID string, tunnelID string, opts ...option.RequestOption) (res *CfdTunnelGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CfdTunnelGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [CfdTunnelUpdateResponseTunnelCfdTunnel] or
// [CfdTunnelUpdateResponseTunnelWarpConnectorTunnel].
type CfdTunnelUpdateResponse interface {
	implementsCfdTunnelUpdateResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*CfdTunnelUpdateResponse)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type CfdTunnelUpdateResponseTunnelCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []CfdTunnelUpdateResponseTunnelCfdTunnelConnection `json:"connections"`
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
	TunType CfdTunnelUpdateResponseTunnelCfdTunnelTunType `json:"tun_type"`
	JSON    cfdTunnelUpdateResponseTunnelCfdTunnelJSON    `json:"-"`
}

// cfdTunnelUpdateResponseTunnelCfdTunnelJSON contains the JSON metadata for the
// struct [CfdTunnelUpdateResponseTunnelCfdTunnel]
type cfdTunnelUpdateResponseTunnelCfdTunnelJSON struct {
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

func (r *CfdTunnelUpdateResponseTunnelCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CfdTunnelUpdateResponseTunnelCfdTunnel) implementsCfdTunnelUpdateResponse() {}

type CfdTunnelUpdateResponseTunnelCfdTunnelConnection struct {
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
	Uuid string                                               `json:"uuid"`
	JSON cfdTunnelUpdateResponseTunnelCfdTunnelConnectionJSON `json:"-"`
}

// cfdTunnelUpdateResponseTunnelCfdTunnelConnectionJSON contains the JSON metadata
// for the struct [CfdTunnelUpdateResponseTunnelCfdTunnelConnection]
type cfdTunnelUpdateResponseTunnelCfdTunnelConnectionJSON struct {
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

func (r *CfdTunnelUpdateResponseTunnelCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type CfdTunnelUpdateResponseTunnelCfdTunnelTunType string

const (
	CfdTunnelUpdateResponseTunnelCfdTunnelTunTypeCfdTunnel     CfdTunnelUpdateResponseTunnelCfdTunnelTunType = "cfd_tunnel"
	CfdTunnelUpdateResponseTunnelCfdTunnelTunTypeWarpConnector CfdTunnelUpdateResponseTunnelCfdTunnelTunType = "warp_connector"
	CfdTunnelUpdateResponseTunnelCfdTunnelTunTypeIPSec         CfdTunnelUpdateResponseTunnelCfdTunnelTunType = "ip_sec"
	CfdTunnelUpdateResponseTunnelCfdTunnelTunTypeGre           CfdTunnelUpdateResponseTunnelCfdTunnelTunType = "gre"
	CfdTunnelUpdateResponseTunnelCfdTunnelTunTypeCni           CfdTunnelUpdateResponseTunnelCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type CfdTunnelUpdateResponseTunnelWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []CfdTunnelUpdateResponseTunnelWarpConnectorTunnelConnection `json:"connections"`
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
	TunType CfdTunnelUpdateResponseTunnelWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    cfdTunnelUpdateResponseTunnelWarpConnectorTunnelJSON    `json:"-"`
}

// cfdTunnelUpdateResponseTunnelWarpConnectorTunnelJSON contains the JSON metadata
// for the struct [CfdTunnelUpdateResponseTunnelWarpConnectorTunnel]
type cfdTunnelUpdateResponseTunnelWarpConnectorTunnelJSON struct {
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

func (r *CfdTunnelUpdateResponseTunnelWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CfdTunnelUpdateResponseTunnelWarpConnectorTunnel) implementsCfdTunnelUpdateResponse() {}

type CfdTunnelUpdateResponseTunnelWarpConnectorTunnelConnection struct {
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
	JSON cfdTunnelUpdateResponseTunnelWarpConnectorTunnelConnectionJSON `json:"-"`
}

// cfdTunnelUpdateResponseTunnelWarpConnectorTunnelConnectionJSON contains the JSON
// metadata for the struct
// [CfdTunnelUpdateResponseTunnelWarpConnectorTunnelConnection]
type cfdTunnelUpdateResponseTunnelWarpConnectorTunnelConnectionJSON struct {
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

func (r *CfdTunnelUpdateResponseTunnelWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type CfdTunnelUpdateResponseTunnelWarpConnectorTunnelTunType string

const (
	CfdTunnelUpdateResponseTunnelWarpConnectorTunnelTunTypeCfdTunnel     CfdTunnelUpdateResponseTunnelWarpConnectorTunnelTunType = "cfd_tunnel"
	CfdTunnelUpdateResponseTunnelWarpConnectorTunnelTunTypeWarpConnector CfdTunnelUpdateResponseTunnelWarpConnectorTunnelTunType = "warp_connector"
	CfdTunnelUpdateResponseTunnelWarpConnectorTunnelTunTypeIPSec         CfdTunnelUpdateResponseTunnelWarpConnectorTunnelTunType = "ip_sec"
	CfdTunnelUpdateResponseTunnelWarpConnectorTunnelTunTypeGre           CfdTunnelUpdateResponseTunnelWarpConnectorTunnelTunType = "gre"
	CfdTunnelUpdateResponseTunnelWarpConnectorTunnelTunTypeCni           CfdTunnelUpdateResponseTunnelWarpConnectorTunnelTunType = "cni"
)

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [CfdTunnelDeleteResponseTunnelCfdTunnel] or
// [CfdTunnelDeleteResponseTunnelWarpConnectorTunnel].
type CfdTunnelDeleteResponse interface {
	implementsCfdTunnelDeleteResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*CfdTunnelDeleteResponse)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type CfdTunnelDeleteResponseTunnelCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []CfdTunnelDeleteResponseTunnelCfdTunnelConnection `json:"connections"`
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
	TunType CfdTunnelDeleteResponseTunnelCfdTunnelTunType `json:"tun_type"`
	JSON    cfdTunnelDeleteResponseTunnelCfdTunnelJSON    `json:"-"`
}

// cfdTunnelDeleteResponseTunnelCfdTunnelJSON contains the JSON metadata for the
// struct [CfdTunnelDeleteResponseTunnelCfdTunnel]
type cfdTunnelDeleteResponseTunnelCfdTunnelJSON struct {
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

func (r *CfdTunnelDeleteResponseTunnelCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CfdTunnelDeleteResponseTunnelCfdTunnel) implementsCfdTunnelDeleteResponse() {}

type CfdTunnelDeleteResponseTunnelCfdTunnelConnection struct {
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
	Uuid string                                               `json:"uuid"`
	JSON cfdTunnelDeleteResponseTunnelCfdTunnelConnectionJSON `json:"-"`
}

// cfdTunnelDeleteResponseTunnelCfdTunnelConnectionJSON contains the JSON metadata
// for the struct [CfdTunnelDeleteResponseTunnelCfdTunnelConnection]
type cfdTunnelDeleteResponseTunnelCfdTunnelConnectionJSON struct {
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

func (r *CfdTunnelDeleteResponseTunnelCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type CfdTunnelDeleteResponseTunnelCfdTunnelTunType string

const (
	CfdTunnelDeleteResponseTunnelCfdTunnelTunTypeCfdTunnel     CfdTunnelDeleteResponseTunnelCfdTunnelTunType = "cfd_tunnel"
	CfdTunnelDeleteResponseTunnelCfdTunnelTunTypeWarpConnector CfdTunnelDeleteResponseTunnelCfdTunnelTunType = "warp_connector"
	CfdTunnelDeleteResponseTunnelCfdTunnelTunTypeIPSec         CfdTunnelDeleteResponseTunnelCfdTunnelTunType = "ip_sec"
	CfdTunnelDeleteResponseTunnelCfdTunnelTunTypeGre           CfdTunnelDeleteResponseTunnelCfdTunnelTunType = "gre"
	CfdTunnelDeleteResponseTunnelCfdTunnelTunTypeCni           CfdTunnelDeleteResponseTunnelCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type CfdTunnelDeleteResponseTunnelWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []CfdTunnelDeleteResponseTunnelWarpConnectorTunnelConnection `json:"connections"`
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
	TunType CfdTunnelDeleteResponseTunnelWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    cfdTunnelDeleteResponseTunnelWarpConnectorTunnelJSON    `json:"-"`
}

// cfdTunnelDeleteResponseTunnelWarpConnectorTunnelJSON contains the JSON metadata
// for the struct [CfdTunnelDeleteResponseTunnelWarpConnectorTunnel]
type cfdTunnelDeleteResponseTunnelWarpConnectorTunnelJSON struct {
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

func (r *CfdTunnelDeleteResponseTunnelWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CfdTunnelDeleteResponseTunnelWarpConnectorTunnel) implementsCfdTunnelDeleteResponse() {}

type CfdTunnelDeleteResponseTunnelWarpConnectorTunnelConnection struct {
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
	JSON cfdTunnelDeleteResponseTunnelWarpConnectorTunnelConnectionJSON `json:"-"`
}

// cfdTunnelDeleteResponseTunnelWarpConnectorTunnelConnectionJSON contains the JSON
// metadata for the struct
// [CfdTunnelDeleteResponseTunnelWarpConnectorTunnelConnection]
type cfdTunnelDeleteResponseTunnelWarpConnectorTunnelConnectionJSON struct {
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

func (r *CfdTunnelDeleteResponseTunnelWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type CfdTunnelDeleteResponseTunnelWarpConnectorTunnelTunType string

const (
	CfdTunnelDeleteResponseTunnelWarpConnectorTunnelTunTypeCfdTunnel     CfdTunnelDeleteResponseTunnelWarpConnectorTunnelTunType = "cfd_tunnel"
	CfdTunnelDeleteResponseTunnelWarpConnectorTunnelTunTypeWarpConnector CfdTunnelDeleteResponseTunnelWarpConnectorTunnelTunType = "warp_connector"
	CfdTunnelDeleteResponseTunnelWarpConnectorTunnelTunTypeIPSec         CfdTunnelDeleteResponseTunnelWarpConnectorTunnelTunType = "ip_sec"
	CfdTunnelDeleteResponseTunnelWarpConnectorTunnelTunTypeGre           CfdTunnelDeleteResponseTunnelWarpConnectorTunnelTunType = "gre"
	CfdTunnelDeleteResponseTunnelWarpConnectorTunnelTunTypeCni           CfdTunnelDeleteResponseTunnelWarpConnectorTunnelTunType = "cni"
)

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by
// [CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelCfdTunnel] or
// [CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelWarpConnectorTunnel].
type CfdTunnelCloudflareTunnelNewACloudflareTunnelResponse interface {
	implementsCfdTunnelCloudflareTunnelNewACloudflareTunnelResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*CfdTunnelCloudflareTunnelNewACloudflareTunnelResponse)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelCfdTunnelConnection `json:"connections"`
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
	TunType CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelCfdTunnelTunType `json:"tun_type"`
	JSON    cfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelCfdTunnelJSON    `json:"-"`
}

// cfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelCfdTunnelJSON
// contains the JSON metadata for the struct
// [CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelCfdTunnel]
type cfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelCfdTunnelJSON struct {
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

func (r *CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelCfdTunnel) implementsCfdTunnelCloudflareTunnelNewACloudflareTunnelResponse() {
}

type CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelCfdTunnelConnection struct {
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
	Uuid string                                                                             `json:"uuid"`
	JSON cfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelCfdTunnelConnectionJSON `json:"-"`
}

// cfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelCfdTunnelConnectionJSON
// contains the JSON metadata for the struct
// [CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelCfdTunnelConnection]
type cfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelCfdTunnelConnectionJSON struct {
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

func (r *CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelCfdTunnelTunType string

const (
	CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelCfdTunnelTunTypeCfdTunnel     CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelCfdTunnelTunType = "cfd_tunnel"
	CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelCfdTunnelTunTypeWarpConnector CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelCfdTunnelTunType = "warp_connector"
	CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelCfdTunnelTunTypeIPSec         CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelCfdTunnelTunType = "ip_sec"
	CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelCfdTunnelTunTypeGre           CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelCfdTunnelTunType = "gre"
	CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelCfdTunnelTunTypeCni           CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelWarpConnectorTunnelConnection `json:"connections"`
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
	TunType CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    cfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelWarpConnectorTunnelJSON    `json:"-"`
}

// cfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelWarpConnectorTunnelJSON
// contains the JSON metadata for the struct
// [CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelWarpConnectorTunnel]
type cfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelWarpConnectorTunnelJSON struct {
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

func (r *CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelWarpConnectorTunnel) implementsCfdTunnelCloudflareTunnelNewACloudflareTunnelResponse() {
}

type CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelWarpConnectorTunnelConnection struct {
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
	Uuid string                                                                                       `json:"uuid"`
	JSON cfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelWarpConnectorTunnelConnectionJSON `json:"-"`
}

// cfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelWarpConnectorTunnelConnectionJSON
// contains the JSON metadata for the struct
// [CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelWarpConnectorTunnelConnection]
type cfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelWarpConnectorTunnelConnectionJSON struct {
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

func (r *CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelWarpConnectorTunnelTunType string

const (
	CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelWarpConnectorTunnelTunTypeCfdTunnel     CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelWarpConnectorTunnelTunType = "cfd_tunnel"
	CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelWarpConnectorTunnelTunTypeWarpConnector CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelWarpConnectorTunnelTunType = "warp_connector"
	CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelWarpConnectorTunnelTunTypeIPSec         CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelWarpConnectorTunnelTunType = "ip_sec"
	CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelWarpConnectorTunnelTunTypeGre           CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelWarpConnectorTunnelTunType = "gre"
	CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelWarpConnectorTunnelTunTypeCni           CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseTunnelWarpConnectorTunnelTunType = "cni"
)

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by
// [CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelCfdTunnel] or
// [CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelWarpConnectorTunnel].
type CfdTunnelCloudflareTunnelListCloudflareTunnelsResponse interface {
	implementsCfdTunnelCloudflareTunnelListCloudflareTunnelsResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*CfdTunnelCloudflareTunnelListCloudflareTunnelsResponse)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelCfdTunnelConnection `json:"connections"`
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
	TunType CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelCfdTunnelTunType `json:"tun_type"`
	JSON    cfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelCfdTunnelJSON    `json:"-"`
}

// cfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelCfdTunnelJSON
// contains the JSON metadata for the struct
// [CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelCfdTunnel]
type cfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelCfdTunnelJSON struct {
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

func (r *CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelCfdTunnel) implementsCfdTunnelCloudflareTunnelListCloudflareTunnelsResponse() {
}

type CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelCfdTunnelConnection struct {
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
	Uuid string                                                                              `json:"uuid"`
	JSON cfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelCfdTunnelConnectionJSON `json:"-"`
}

// cfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelCfdTunnelConnectionJSON
// contains the JSON metadata for the struct
// [CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelCfdTunnelConnection]
type cfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelCfdTunnelConnectionJSON struct {
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

func (r *CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelCfdTunnelTunType string

const (
	CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelCfdTunnelTunTypeCfdTunnel     CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelCfdTunnelTunType = "cfd_tunnel"
	CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelCfdTunnelTunTypeWarpConnector CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelCfdTunnelTunType = "warp_connector"
	CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelCfdTunnelTunTypeIPSec         CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelCfdTunnelTunType = "ip_sec"
	CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelCfdTunnelTunTypeGre           CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelCfdTunnelTunType = "gre"
	CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelCfdTunnelTunTypeCni           CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelWarpConnectorTunnelConnection `json:"connections"`
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
	TunType CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    cfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelWarpConnectorTunnelJSON    `json:"-"`
}

// cfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelWarpConnectorTunnelJSON
// contains the JSON metadata for the struct
// [CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelWarpConnectorTunnel]
type cfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelWarpConnectorTunnelJSON struct {
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

func (r *CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelWarpConnectorTunnel) implementsCfdTunnelCloudflareTunnelListCloudflareTunnelsResponse() {
}

type CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelWarpConnectorTunnelConnection struct {
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
	Uuid string                                                                                        `json:"uuid"`
	JSON cfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelWarpConnectorTunnelConnectionJSON `json:"-"`
}

// cfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelWarpConnectorTunnelConnectionJSON
// contains the JSON metadata for the struct
// [CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelWarpConnectorTunnelConnection]
type cfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelWarpConnectorTunnelConnectionJSON struct {
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

func (r *CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelWarpConnectorTunnelTunType string

const (
	CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelWarpConnectorTunnelTunTypeCfdTunnel     CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelWarpConnectorTunnelTunType = "cfd_tunnel"
	CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelWarpConnectorTunnelTunTypeWarpConnector CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelWarpConnectorTunnelTunType = "warp_connector"
	CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelWarpConnectorTunnelTunTypeIPSec         CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelWarpConnectorTunnelTunType = "ip_sec"
	CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelWarpConnectorTunnelTunTypeGre           CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelWarpConnectorTunnelTunType = "gre"
	CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelWarpConnectorTunnelTunTypeCni           CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseTunnelWarpConnectorTunnelTunType = "cni"
)

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [CfdTunnelGetResponseTunnelCfdTunnel] or
// [CfdTunnelGetResponseTunnelWarpConnectorTunnel].
type CfdTunnelGetResponse interface {
	implementsCfdTunnelGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*CfdTunnelGetResponse)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type CfdTunnelGetResponseTunnelCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []CfdTunnelGetResponseTunnelCfdTunnelConnection `json:"connections"`
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
	TunType CfdTunnelGetResponseTunnelCfdTunnelTunType `json:"tun_type"`
	JSON    cfdTunnelGetResponseTunnelCfdTunnelJSON    `json:"-"`
}

// cfdTunnelGetResponseTunnelCfdTunnelJSON contains the JSON metadata for the
// struct [CfdTunnelGetResponseTunnelCfdTunnel]
type cfdTunnelGetResponseTunnelCfdTunnelJSON struct {
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

func (r *CfdTunnelGetResponseTunnelCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CfdTunnelGetResponseTunnelCfdTunnel) implementsCfdTunnelGetResponse() {}

type CfdTunnelGetResponseTunnelCfdTunnelConnection struct {
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
	Uuid string                                            `json:"uuid"`
	JSON cfdTunnelGetResponseTunnelCfdTunnelConnectionJSON `json:"-"`
}

// cfdTunnelGetResponseTunnelCfdTunnelConnectionJSON contains the JSON metadata for
// the struct [CfdTunnelGetResponseTunnelCfdTunnelConnection]
type cfdTunnelGetResponseTunnelCfdTunnelConnectionJSON struct {
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

func (r *CfdTunnelGetResponseTunnelCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type CfdTunnelGetResponseTunnelCfdTunnelTunType string

const (
	CfdTunnelGetResponseTunnelCfdTunnelTunTypeCfdTunnel     CfdTunnelGetResponseTunnelCfdTunnelTunType = "cfd_tunnel"
	CfdTunnelGetResponseTunnelCfdTunnelTunTypeWarpConnector CfdTunnelGetResponseTunnelCfdTunnelTunType = "warp_connector"
	CfdTunnelGetResponseTunnelCfdTunnelTunTypeIPSec         CfdTunnelGetResponseTunnelCfdTunnelTunType = "ip_sec"
	CfdTunnelGetResponseTunnelCfdTunnelTunTypeGre           CfdTunnelGetResponseTunnelCfdTunnelTunType = "gre"
	CfdTunnelGetResponseTunnelCfdTunnelTunTypeCni           CfdTunnelGetResponseTunnelCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type CfdTunnelGetResponseTunnelWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []CfdTunnelGetResponseTunnelWarpConnectorTunnelConnection `json:"connections"`
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
	TunType CfdTunnelGetResponseTunnelWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    cfdTunnelGetResponseTunnelWarpConnectorTunnelJSON    `json:"-"`
}

// cfdTunnelGetResponseTunnelWarpConnectorTunnelJSON contains the JSON metadata for
// the struct [CfdTunnelGetResponseTunnelWarpConnectorTunnel]
type cfdTunnelGetResponseTunnelWarpConnectorTunnelJSON struct {
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

func (r *CfdTunnelGetResponseTunnelWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CfdTunnelGetResponseTunnelWarpConnectorTunnel) implementsCfdTunnelGetResponse() {}

type CfdTunnelGetResponseTunnelWarpConnectorTunnelConnection struct {
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
	Uuid string                                                      `json:"uuid"`
	JSON cfdTunnelGetResponseTunnelWarpConnectorTunnelConnectionJSON `json:"-"`
}

// cfdTunnelGetResponseTunnelWarpConnectorTunnelConnectionJSON contains the JSON
// metadata for the struct
// [CfdTunnelGetResponseTunnelWarpConnectorTunnelConnection]
type cfdTunnelGetResponseTunnelWarpConnectorTunnelConnectionJSON struct {
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

func (r *CfdTunnelGetResponseTunnelWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type CfdTunnelGetResponseTunnelWarpConnectorTunnelTunType string

const (
	CfdTunnelGetResponseTunnelWarpConnectorTunnelTunTypeCfdTunnel     CfdTunnelGetResponseTunnelWarpConnectorTunnelTunType = "cfd_tunnel"
	CfdTunnelGetResponseTunnelWarpConnectorTunnelTunTypeWarpConnector CfdTunnelGetResponseTunnelWarpConnectorTunnelTunType = "warp_connector"
	CfdTunnelGetResponseTunnelWarpConnectorTunnelTunTypeIPSec         CfdTunnelGetResponseTunnelWarpConnectorTunnelTunType = "ip_sec"
	CfdTunnelGetResponseTunnelWarpConnectorTunnelTunTypeGre           CfdTunnelGetResponseTunnelWarpConnectorTunnelTunType = "gre"
	CfdTunnelGetResponseTunnelWarpConnectorTunnelTunTypeCni           CfdTunnelGetResponseTunnelWarpConnectorTunnelTunType = "cni"
)

type CfdTunnelUpdateParams struct {
	// A user-friendly name for the tunnel.
	Name param.Field[string] `json:"name"`
	// Sets the password required to run a locally-managed tunnel. Must be at least 32
	// bytes and encoded as a base64 string.
	TunnelSecret param.Field[string] `json:"tunnel_secret"`
}

func (r CfdTunnelUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CfdTunnelUpdateResponseEnvelope struct {
	Errors   []CfdTunnelUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CfdTunnelUpdateResponseEnvelopeMessages `json:"messages,required"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result CfdTunnelUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success CfdTunnelUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    cfdTunnelUpdateResponseEnvelopeJSON    `json:"-"`
}

// cfdTunnelUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [CfdTunnelUpdateResponseEnvelope]
type cfdTunnelUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CfdTunnelUpdateResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    cfdTunnelUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// cfdTunnelUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CfdTunnelUpdateResponseEnvelopeErrors]
type cfdTunnelUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CfdTunnelUpdateResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    cfdTunnelUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// cfdTunnelUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CfdTunnelUpdateResponseEnvelopeMessages]
type cfdTunnelUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CfdTunnelUpdateResponseEnvelopeSuccess bool

const (
	CfdTunnelUpdateResponseEnvelopeSuccessTrue CfdTunnelUpdateResponseEnvelopeSuccess = true
)

type CfdTunnelDeleteParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r CfdTunnelDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type CfdTunnelDeleteResponseEnvelope struct {
	Errors   []CfdTunnelDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CfdTunnelDeleteResponseEnvelopeMessages `json:"messages,required"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result CfdTunnelDeleteResponse `json:"result,required"`
	// Whether the API call was successful
	Success CfdTunnelDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    cfdTunnelDeleteResponseEnvelopeJSON    `json:"-"`
}

// cfdTunnelDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [CfdTunnelDeleteResponseEnvelope]
type cfdTunnelDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CfdTunnelDeleteResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    cfdTunnelDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// cfdTunnelDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CfdTunnelDeleteResponseEnvelopeErrors]
type cfdTunnelDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CfdTunnelDeleteResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    cfdTunnelDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// cfdTunnelDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CfdTunnelDeleteResponseEnvelopeMessages]
type cfdTunnelDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CfdTunnelDeleteResponseEnvelopeSuccess bool

const (
	CfdTunnelDeleteResponseEnvelopeSuccessTrue CfdTunnelDeleteResponseEnvelopeSuccess = true
)

type CfdTunnelCloudflareTunnelNewACloudflareTunnelParams struct {
	// A user-friendly name for the tunnel.
	Name param.Field[string] `json:"name,required"`
	// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
	// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
	// tunnel on the Zero Trust dashboard or using the
	// [Cloudflare Tunnel configuration](https://api.cloudflare.com/#cloudflare-tunnel-configuration-properties)
	// endpoint.
	ConfigSrc param.Field[CfdTunnelCloudflareTunnelNewACloudflareTunnelParamsConfigSrc] `json:"config_src"`
	// Sets the password required to run a locally-managed tunnel. Must be at least 32
	// bytes and encoded as a base64 string.
	TunnelSecret param.Field[string] `json:"tunnel_secret"`
}

func (r CfdTunnelCloudflareTunnelNewACloudflareTunnelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
// tunnel on the Zero Trust dashboard or using the
// [Cloudflare Tunnel configuration](https://api.cloudflare.com/#cloudflare-tunnel-configuration-properties)
// endpoint.
type CfdTunnelCloudflareTunnelNewACloudflareTunnelParamsConfigSrc string

const (
	CfdTunnelCloudflareTunnelNewACloudflareTunnelParamsConfigSrcLocal      CfdTunnelCloudflareTunnelNewACloudflareTunnelParamsConfigSrc = "local"
	CfdTunnelCloudflareTunnelNewACloudflareTunnelParamsConfigSrcCloudflare CfdTunnelCloudflareTunnelNewACloudflareTunnelParamsConfigSrc = "cloudflare"
)

type CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseEnvelope struct {
	Errors   []CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseEnvelopeMessages `json:"messages,required"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result CfdTunnelCloudflareTunnelNewACloudflareTunnelResponse `json:"result,required"`
	// Whether the API call was successful
	Success CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseEnvelopeSuccess `json:"success,required"`
	JSON    cfdTunnelCloudflareTunnelNewACloudflareTunnelResponseEnvelopeJSON    `json:"-"`
}

// cfdTunnelCloudflareTunnelNewACloudflareTunnelResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseEnvelope]
type cfdTunnelCloudflareTunnelNewACloudflareTunnelResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseEnvelopeErrors struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    cfdTunnelCloudflareTunnelNewACloudflareTunnelResponseEnvelopeErrorsJSON `json:"-"`
}

// cfdTunnelCloudflareTunnelNewACloudflareTunnelResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseEnvelopeErrors]
type cfdTunnelCloudflareTunnelNewACloudflareTunnelResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseEnvelopeMessages struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    cfdTunnelCloudflareTunnelNewACloudflareTunnelResponseEnvelopeMessagesJSON `json:"-"`
}

// cfdTunnelCloudflareTunnelNewACloudflareTunnelResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseEnvelopeMessages]
type cfdTunnelCloudflareTunnelNewACloudflareTunnelResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseEnvelopeSuccess bool

const (
	CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseEnvelopeSuccessTrue CfdTunnelCloudflareTunnelNewACloudflareTunnelResponseEnvelopeSuccess = true
)

type CfdTunnelCloudflareTunnelListCloudflareTunnelsParams struct {
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

// URLQuery serializes [CfdTunnelCloudflareTunnelListCloudflareTunnelsParams]'s
// query parameters as `url.Values`.
func (r CfdTunnelCloudflareTunnelListCloudflareTunnelsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelope struct {
	Errors   []CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelopeMessages `json:"messages,required"`
	Result   []CfdTunnelCloudflareTunnelListCloudflareTunnelsResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelopeResultInfo `json:"result_info"`
	JSON       cfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelopeJSON       `json:"-"`
}

// cfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelope]
type cfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelopeErrors struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    cfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelopeErrorsJSON `json:"-"`
}

// cfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelopeErrors]
type cfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelopeMessages struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    cfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelopeMessagesJSON `json:"-"`
}

// cfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelopeMessages]
type cfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelopeSuccess bool

const (
	CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelopeSuccessTrue CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelopeSuccess = true
)

type CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                      `json:"total_count"`
	JSON       cfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelopeResultInfoJSON `json:"-"`
}

// cfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelopeResultInfo]
type cfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelCloudflareTunnelListCloudflareTunnelsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CfdTunnelGetResponseEnvelope struct {
	Errors   []CfdTunnelGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CfdTunnelGetResponseEnvelopeMessages `json:"messages,required"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result CfdTunnelGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success CfdTunnelGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    cfdTunnelGetResponseEnvelopeJSON    `json:"-"`
}

// cfdTunnelGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CfdTunnelGetResponseEnvelope]
type cfdTunnelGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CfdTunnelGetResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    cfdTunnelGetResponseEnvelopeErrorsJSON `json:"-"`
}

// cfdTunnelGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [CfdTunnelGetResponseEnvelopeErrors]
type cfdTunnelGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CfdTunnelGetResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    cfdTunnelGetResponseEnvelopeMessagesJSON `json:"-"`
}

// cfdTunnelGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CfdTunnelGetResponseEnvelopeMessages]
type cfdTunnelGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CfdTunnelGetResponseEnvelopeSuccess bool

const (
	CfdTunnelGetResponseEnvelopeSuccessTrue CfdTunnelGetResponseEnvelopeSuccess = true
)
