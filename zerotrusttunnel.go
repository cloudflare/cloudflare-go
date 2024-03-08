// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
)

// ZeroTrustTunnelService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustTunnelService] method
// instead.
type ZeroTrustTunnelService struct {
	Options        []option.RequestOption
	Configurations *ZeroTrustTunnelConfigurationService
	Connections    *ZeroTrustTunnelConnectionService
	Token          *ZeroTrustTunnelTokenService
	Connectors     *ZeroTrustTunnelConnectorService
	Management     *ZeroTrustTunnelManagementService
}

// NewZeroTrustTunnelService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZeroTrustTunnelService(opts ...option.RequestOption) (r *ZeroTrustTunnelService) {
	r = &ZeroTrustTunnelService{}
	r.Options = opts
	r.Configurations = NewZeroTrustTunnelConfigurationService(opts...)
	r.Connections = NewZeroTrustTunnelConnectionService(opts...)
	r.Token = NewZeroTrustTunnelTokenService(opts...)
	r.Connectors = NewZeroTrustTunnelConnectorService(opts...)
	r.Management = NewZeroTrustTunnelManagementService(opts...)
	return
}

// Creates a new Argo Tunnel in an account.
func (r *ZeroTrustTunnelService) New(ctx context.Context, params ZeroTrustTunnelNewParams, opts ...option.RequestOption) (res *ZeroTrustTunnelNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustTunnelNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/tunnels", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists and filters all types of Tunnels in an account.
func (r *ZeroTrustTunnelService) List(ctx context.Context, params ZeroTrustTunnelListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[ZeroTrustTunnelListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/tunnels", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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
func (r *ZeroTrustTunnelService) ListAutoPaging(ctx context.Context, params ZeroTrustTunnelListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[ZeroTrustTunnelListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes an Argo Tunnel from an account.
func (r *ZeroTrustTunnelService) Delete(ctx context.Context, tunnelID string, params ZeroTrustTunnelDeleteParams, opts ...option.RequestOption) (res *ZeroTrustTunnelDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustTunnelDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/tunnels/%s", params.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing Cloudflare Tunnel.
func (r *ZeroTrustTunnelService) Edit(ctx context.Context, tunnelID string, params ZeroTrustTunnelEditParams, opts ...option.RequestOption) (res *ZeroTrustTunnelEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustTunnelEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s", params.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Argo Tunnel.
func (r *ZeroTrustTunnelService) Get(ctx context.Context, tunnelID string, query ZeroTrustTunnelGetParams, opts ...option.RequestOption) (res *ZeroTrustTunnelGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustTunnelGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/tunnels/%s", query.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustTunnelNewResponse struct {
	// UUID of the tunnel.
	ID string `json:"id,required"`
	// The tunnel connections between your origin and Cloudflare's edge.
	Connections []ZeroTrustTunnelNewResponseConnection `json:"connections,required"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A user-friendly name for the tunnel.
	Name string `json:"name,required"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time                      `json:"deleted_at,nullable" format:"date-time"`
	JSON      zeroTrustTunnelNewResponseJSON `json:"-"`
}

// zeroTrustTunnelNewResponseJSON contains the JSON metadata for the struct
// [ZeroTrustTunnelNewResponse]
type zeroTrustTunnelNewResponseJSON struct {
	ID          apijson.Field
	Connections apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	DeletedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelNewResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustTunnelNewResponseConnection struct {
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// UUID of the Cloudflare Tunnel connection.
	UUID string                                   `json:"uuid"`
	JSON zeroTrustTunnelNewResponseConnectionJSON `json:"-"`
}

// zeroTrustTunnelNewResponseConnectionJSON contains the JSON metadata for the
// struct [ZeroTrustTunnelNewResponseConnection]
type zeroTrustTunnelNewResponseConnectionJSON struct {
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	UUID               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustTunnelNewResponseConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelNewResponseConnectionJSON) RawJSON() string {
	return r.raw
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [ZeroTrustTunnelListResponseTunnelCfdTunnel] or
// [ZeroTrustTunnelListResponseTunnelWARPConnectorTunnel].
type ZeroTrustTunnelListResponse interface {
	implementsZeroTrustTunnelListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustTunnelListResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustTunnelListResponseTunnelCfdTunnel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustTunnelListResponseTunnelWARPConnectorTunnel{}),
		},
	)
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type ZeroTrustTunnelListResponseTunnelCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []ZeroTrustTunnelListResponseTunnelCfdTunnelConnection `json:"connections"`
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
	TunType ZeroTrustTunnelListResponseTunnelCfdTunnelTunType `json:"tun_type"`
	JSON    zeroTrustTunnelListResponseTunnelCfdTunnelJSON    `json:"-"`
}

// zeroTrustTunnelListResponseTunnelCfdTunnelJSON contains the JSON metadata for
// the struct [ZeroTrustTunnelListResponseTunnelCfdTunnel]
type zeroTrustTunnelListResponseTunnelCfdTunnelJSON struct {
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

func (r *ZeroTrustTunnelListResponseTunnelCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelListResponseTunnelCfdTunnelJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustTunnelListResponseTunnelCfdTunnel) implementsZeroTrustTunnelListResponse() {}

type ZeroTrustTunnelListResponseTunnelCfdTunnelConnection struct {
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
	UUID string                                                   `json:"uuid"`
	JSON zeroTrustTunnelListResponseTunnelCfdTunnelConnectionJSON `json:"-"`
}

// zeroTrustTunnelListResponseTunnelCfdTunnelConnectionJSON contains the JSON
// metadata for the struct [ZeroTrustTunnelListResponseTunnelCfdTunnelConnection]
type zeroTrustTunnelListResponseTunnelCfdTunnelConnectionJSON struct {
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

func (r *ZeroTrustTunnelListResponseTunnelCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelListResponseTunnelCfdTunnelConnectionJSON) RawJSON() string {
	return r.raw
}

// The type of tunnel.
type ZeroTrustTunnelListResponseTunnelCfdTunnelTunType string

const (
	ZeroTrustTunnelListResponseTunnelCfdTunnelTunTypeCfdTunnel     ZeroTrustTunnelListResponseTunnelCfdTunnelTunType = "cfd_tunnel"
	ZeroTrustTunnelListResponseTunnelCfdTunnelTunTypeWARPConnector ZeroTrustTunnelListResponseTunnelCfdTunnelTunType = "warp_connector"
	ZeroTrustTunnelListResponseTunnelCfdTunnelTunTypeIPSec         ZeroTrustTunnelListResponseTunnelCfdTunnelTunType = "ip_sec"
	ZeroTrustTunnelListResponseTunnelCfdTunnelTunTypeGRE           ZeroTrustTunnelListResponseTunnelCfdTunnelTunType = "gre"
	ZeroTrustTunnelListResponseTunnelCfdTunnelTunTypeCni           ZeroTrustTunnelListResponseTunnelCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type ZeroTrustTunnelListResponseTunnelWARPConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []ZeroTrustTunnelListResponseTunnelWARPConnectorTunnelConnection `json:"connections"`
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
	TunType ZeroTrustTunnelListResponseTunnelWARPConnectorTunnelTunType `json:"tun_type"`
	JSON    zeroTrustTunnelListResponseTunnelWARPConnectorTunnelJSON    `json:"-"`
}

// zeroTrustTunnelListResponseTunnelWARPConnectorTunnelJSON contains the JSON
// metadata for the struct [ZeroTrustTunnelListResponseTunnelWARPConnectorTunnel]
type zeroTrustTunnelListResponseTunnelWARPConnectorTunnelJSON struct {
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

func (r *ZeroTrustTunnelListResponseTunnelWARPConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelListResponseTunnelWARPConnectorTunnelJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustTunnelListResponseTunnelWARPConnectorTunnel) implementsZeroTrustTunnelListResponse() {
}

type ZeroTrustTunnelListResponseTunnelWARPConnectorTunnelConnection struct {
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
	UUID string                                                             `json:"uuid"`
	JSON zeroTrustTunnelListResponseTunnelWARPConnectorTunnelConnectionJSON `json:"-"`
}

// zeroTrustTunnelListResponseTunnelWARPConnectorTunnelConnectionJSON contains the
// JSON metadata for the struct
// [ZeroTrustTunnelListResponseTunnelWARPConnectorTunnelConnection]
type zeroTrustTunnelListResponseTunnelWARPConnectorTunnelConnectionJSON struct {
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

func (r *ZeroTrustTunnelListResponseTunnelWARPConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelListResponseTunnelWARPConnectorTunnelConnectionJSON) RawJSON() string {
	return r.raw
}

// The type of tunnel.
type ZeroTrustTunnelListResponseTunnelWARPConnectorTunnelTunType string

const (
	ZeroTrustTunnelListResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel     ZeroTrustTunnelListResponseTunnelWARPConnectorTunnelTunType = "cfd_tunnel"
	ZeroTrustTunnelListResponseTunnelWARPConnectorTunnelTunTypeWARPConnector ZeroTrustTunnelListResponseTunnelWARPConnectorTunnelTunType = "warp_connector"
	ZeroTrustTunnelListResponseTunnelWARPConnectorTunnelTunTypeIPSec         ZeroTrustTunnelListResponseTunnelWARPConnectorTunnelTunType = "ip_sec"
	ZeroTrustTunnelListResponseTunnelWARPConnectorTunnelTunTypeGRE           ZeroTrustTunnelListResponseTunnelWARPConnectorTunnelTunType = "gre"
	ZeroTrustTunnelListResponseTunnelWARPConnectorTunnelTunTypeCni           ZeroTrustTunnelListResponseTunnelWARPConnectorTunnelTunType = "cni"
)

type ZeroTrustTunnelDeleteResponse struct {
	// UUID of the tunnel.
	ID string `json:"id,required"`
	// The tunnel connections between your origin and Cloudflare's edge.
	Connections []ZeroTrustTunnelDeleteResponseConnection `json:"connections,required"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A user-friendly name for the tunnel.
	Name string `json:"name,required"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time                         `json:"deleted_at,nullable" format:"date-time"`
	JSON      zeroTrustTunnelDeleteResponseJSON `json:"-"`
}

// zeroTrustTunnelDeleteResponseJSON contains the JSON metadata for the struct
// [ZeroTrustTunnelDeleteResponse]
type zeroTrustTunnelDeleteResponseJSON struct {
	ID          apijson.Field
	Connections apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	DeletedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustTunnelDeleteResponseConnection struct {
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// UUID of the Cloudflare Tunnel connection.
	UUID string                                      `json:"uuid"`
	JSON zeroTrustTunnelDeleteResponseConnectionJSON `json:"-"`
}

// zeroTrustTunnelDeleteResponseConnectionJSON contains the JSON metadata for the
// struct [ZeroTrustTunnelDeleteResponseConnection]
type zeroTrustTunnelDeleteResponseConnectionJSON struct {
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	UUID               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustTunnelDeleteResponseConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelDeleteResponseConnectionJSON) RawJSON() string {
	return r.raw
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [ZeroTrustTunnelEditResponseTunnelCfdTunnel] or
// [ZeroTrustTunnelEditResponseTunnelWARPConnectorTunnel].
type ZeroTrustTunnelEditResponse interface {
	implementsZeroTrustTunnelEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustTunnelEditResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustTunnelEditResponseTunnelCfdTunnel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustTunnelEditResponseTunnelWARPConnectorTunnel{}),
		},
	)
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type ZeroTrustTunnelEditResponseTunnelCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []ZeroTrustTunnelEditResponseTunnelCfdTunnelConnection `json:"connections"`
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
	TunType ZeroTrustTunnelEditResponseTunnelCfdTunnelTunType `json:"tun_type"`
	JSON    zeroTrustTunnelEditResponseTunnelCfdTunnelJSON    `json:"-"`
}

// zeroTrustTunnelEditResponseTunnelCfdTunnelJSON contains the JSON metadata for
// the struct [ZeroTrustTunnelEditResponseTunnelCfdTunnel]
type zeroTrustTunnelEditResponseTunnelCfdTunnelJSON struct {
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

func (r *ZeroTrustTunnelEditResponseTunnelCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelEditResponseTunnelCfdTunnelJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustTunnelEditResponseTunnelCfdTunnel) implementsZeroTrustTunnelEditResponse() {}

type ZeroTrustTunnelEditResponseTunnelCfdTunnelConnection struct {
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
	UUID string                                                   `json:"uuid"`
	JSON zeroTrustTunnelEditResponseTunnelCfdTunnelConnectionJSON `json:"-"`
}

// zeroTrustTunnelEditResponseTunnelCfdTunnelConnectionJSON contains the JSON
// metadata for the struct [ZeroTrustTunnelEditResponseTunnelCfdTunnelConnection]
type zeroTrustTunnelEditResponseTunnelCfdTunnelConnectionJSON struct {
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

func (r *ZeroTrustTunnelEditResponseTunnelCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelEditResponseTunnelCfdTunnelConnectionJSON) RawJSON() string {
	return r.raw
}

// The type of tunnel.
type ZeroTrustTunnelEditResponseTunnelCfdTunnelTunType string

const (
	ZeroTrustTunnelEditResponseTunnelCfdTunnelTunTypeCfdTunnel     ZeroTrustTunnelEditResponseTunnelCfdTunnelTunType = "cfd_tunnel"
	ZeroTrustTunnelEditResponseTunnelCfdTunnelTunTypeWARPConnector ZeroTrustTunnelEditResponseTunnelCfdTunnelTunType = "warp_connector"
	ZeroTrustTunnelEditResponseTunnelCfdTunnelTunTypeIPSec         ZeroTrustTunnelEditResponseTunnelCfdTunnelTunType = "ip_sec"
	ZeroTrustTunnelEditResponseTunnelCfdTunnelTunTypeGRE           ZeroTrustTunnelEditResponseTunnelCfdTunnelTunType = "gre"
	ZeroTrustTunnelEditResponseTunnelCfdTunnelTunTypeCni           ZeroTrustTunnelEditResponseTunnelCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type ZeroTrustTunnelEditResponseTunnelWARPConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []ZeroTrustTunnelEditResponseTunnelWARPConnectorTunnelConnection `json:"connections"`
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
	TunType ZeroTrustTunnelEditResponseTunnelWARPConnectorTunnelTunType `json:"tun_type"`
	JSON    zeroTrustTunnelEditResponseTunnelWARPConnectorTunnelJSON    `json:"-"`
}

// zeroTrustTunnelEditResponseTunnelWARPConnectorTunnelJSON contains the JSON
// metadata for the struct [ZeroTrustTunnelEditResponseTunnelWARPConnectorTunnel]
type zeroTrustTunnelEditResponseTunnelWARPConnectorTunnelJSON struct {
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

func (r *ZeroTrustTunnelEditResponseTunnelWARPConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelEditResponseTunnelWARPConnectorTunnelJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustTunnelEditResponseTunnelWARPConnectorTunnel) implementsZeroTrustTunnelEditResponse() {
}

type ZeroTrustTunnelEditResponseTunnelWARPConnectorTunnelConnection struct {
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
	UUID string                                                             `json:"uuid"`
	JSON zeroTrustTunnelEditResponseTunnelWARPConnectorTunnelConnectionJSON `json:"-"`
}

// zeroTrustTunnelEditResponseTunnelWARPConnectorTunnelConnectionJSON contains the
// JSON metadata for the struct
// [ZeroTrustTunnelEditResponseTunnelWARPConnectorTunnelConnection]
type zeroTrustTunnelEditResponseTunnelWARPConnectorTunnelConnectionJSON struct {
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

func (r *ZeroTrustTunnelEditResponseTunnelWARPConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelEditResponseTunnelWARPConnectorTunnelConnectionJSON) RawJSON() string {
	return r.raw
}

// The type of tunnel.
type ZeroTrustTunnelEditResponseTunnelWARPConnectorTunnelTunType string

const (
	ZeroTrustTunnelEditResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel     ZeroTrustTunnelEditResponseTunnelWARPConnectorTunnelTunType = "cfd_tunnel"
	ZeroTrustTunnelEditResponseTunnelWARPConnectorTunnelTunTypeWARPConnector ZeroTrustTunnelEditResponseTunnelWARPConnectorTunnelTunType = "warp_connector"
	ZeroTrustTunnelEditResponseTunnelWARPConnectorTunnelTunTypeIPSec         ZeroTrustTunnelEditResponseTunnelWARPConnectorTunnelTunType = "ip_sec"
	ZeroTrustTunnelEditResponseTunnelWARPConnectorTunnelTunTypeGRE           ZeroTrustTunnelEditResponseTunnelWARPConnectorTunnelTunType = "gre"
	ZeroTrustTunnelEditResponseTunnelWARPConnectorTunnelTunTypeCni           ZeroTrustTunnelEditResponseTunnelWARPConnectorTunnelTunType = "cni"
)

type ZeroTrustTunnelGetResponse struct {
	// UUID of the tunnel.
	ID string `json:"id,required"`
	// The tunnel connections between your origin and Cloudflare's edge.
	Connections []ZeroTrustTunnelGetResponseConnection `json:"connections,required"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A user-friendly name for the tunnel.
	Name string `json:"name,required"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time                      `json:"deleted_at,nullable" format:"date-time"`
	JSON      zeroTrustTunnelGetResponseJSON `json:"-"`
}

// zeroTrustTunnelGetResponseJSON contains the JSON metadata for the struct
// [ZeroTrustTunnelGetResponse]
type zeroTrustTunnelGetResponseJSON struct {
	ID          apijson.Field
	Connections apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	DeletedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustTunnelGetResponseConnection struct {
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// UUID of the Cloudflare Tunnel connection.
	UUID string                                   `json:"uuid"`
	JSON zeroTrustTunnelGetResponseConnectionJSON `json:"-"`
}

// zeroTrustTunnelGetResponseConnectionJSON contains the JSON metadata for the
// struct [ZeroTrustTunnelGetResponseConnection]
type zeroTrustTunnelGetResponseConnectionJSON struct {
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	UUID               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustTunnelGetResponseConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelGetResponseConnectionJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustTunnelNewParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// A user-friendly name for the tunnel.
	Name param.Field[string] `json:"name,required"`
	// Sets the password required to run the tunnel. Must be at least 32 bytes and
	// encoded as a base64 string.
	TunnelSecret param.Field[interface{}] `json:"tunnel_secret,required"`
}

func (r ZeroTrustTunnelNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustTunnelNewResponseEnvelope struct {
	Errors   []ZeroTrustTunnelNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustTunnelNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustTunnelNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustTunnelNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustTunnelNewResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustTunnelNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZeroTrustTunnelNewResponseEnvelope]
type zeroTrustTunnelNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustTunnelNewResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zeroTrustTunnelNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustTunnelNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZeroTrustTunnelNewResponseEnvelopeErrors]
type zeroTrustTunnelNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustTunnelNewResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zeroTrustTunnelNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustTunnelNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZeroTrustTunnelNewResponseEnvelopeMessages]
type zeroTrustTunnelNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustTunnelNewResponseEnvelopeSuccess bool

const (
	ZeroTrustTunnelNewResponseEnvelopeSuccessTrue ZeroTrustTunnelNewResponseEnvelopeSuccess = true
)

type ZeroTrustTunnelListParams struct {
	// Cloudflare account ID
	AccountID     param.Field[string] `path:"account_id,required"`
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

// URLQuery serializes [ZeroTrustTunnelListParams]'s query parameters as
// `url.Values`.
func (r ZeroTrustTunnelListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZeroTrustTunnelDeleteParams struct {
	// Cloudflare account ID
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r ZeroTrustTunnelDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZeroTrustTunnelDeleteResponseEnvelope struct {
	Errors   []ZeroTrustTunnelDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustTunnelDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustTunnelDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustTunnelDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustTunnelDeleteResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustTunnelDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustTunnelDeleteResponseEnvelope]
type zeroTrustTunnelDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustTunnelDeleteResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zeroTrustTunnelDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustTunnelDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZeroTrustTunnelDeleteResponseEnvelopeErrors]
type zeroTrustTunnelDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustTunnelDeleteResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zeroTrustTunnelDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustTunnelDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZeroTrustTunnelDeleteResponseEnvelopeMessages]
type zeroTrustTunnelDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustTunnelDeleteResponseEnvelopeSuccess bool

const (
	ZeroTrustTunnelDeleteResponseEnvelopeSuccessTrue ZeroTrustTunnelDeleteResponseEnvelopeSuccess = true
)

type ZeroTrustTunnelEditParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// A user-friendly name for the tunnel.
	Name param.Field[string] `json:"name"`
	// Sets the password required to run a locally-managed tunnel. Must be at least 32
	// bytes and encoded as a base64 string.
	TunnelSecret param.Field[string] `json:"tunnel_secret"`
}

func (r ZeroTrustTunnelEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustTunnelEditResponseEnvelope struct {
	Errors   []ZeroTrustTunnelEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustTunnelEditResponseEnvelopeMessages `json:"messages,required"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result ZeroTrustTunnelEditResponse `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustTunnelEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustTunnelEditResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustTunnelEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustTunnelEditResponseEnvelope]
type zeroTrustTunnelEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustTunnelEditResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zeroTrustTunnelEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustTunnelEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZeroTrustTunnelEditResponseEnvelopeErrors]
type zeroTrustTunnelEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustTunnelEditResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zeroTrustTunnelEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustTunnelEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZeroTrustTunnelEditResponseEnvelopeMessages]
type zeroTrustTunnelEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustTunnelEditResponseEnvelopeSuccess bool

const (
	ZeroTrustTunnelEditResponseEnvelopeSuccessTrue ZeroTrustTunnelEditResponseEnvelopeSuccess = true
)

type ZeroTrustTunnelGetParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZeroTrustTunnelGetResponseEnvelope struct {
	Errors   []ZeroTrustTunnelGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustTunnelGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustTunnelGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustTunnelGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustTunnelGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustTunnelGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZeroTrustTunnelGetResponseEnvelope]
type zeroTrustTunnelGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustTunnelGetResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zeroTrustTunnelGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustTunnelGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZeroTrustTunnelGetResponseEnvelopeErrors]
type zeroTrustTunnelGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustTunnelGetResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zeroTrustTunnelGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustTunnelGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZeroTrustTunnelGetResponseEnvelopeMessages]
type zeroTrustTunnelGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustTunnelGetResponseEnvelopeSuccess bool

const (
	ZeroTrustTunnelGetResponseEnvelopeSuccessTrue ZeroTrustTunnelGetResponseEnvelopeSuccess = true
)
