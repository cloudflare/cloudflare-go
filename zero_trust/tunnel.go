// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// TunnelService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewTunnelService] method instead.
type TunnelService struct {
	Options        []option.RequestOption
	Configurations *TunnelConfigurationService
	Connections    *TunnelConnectionService
	Token          *TunnelTokenService
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
	r.Token = NewTunnelTokenService(opts...)
	r.Connectors = NewTunnelConnectorService(opts...)
	r.Management = NewTunnelManagementService(opts...)
	return
}

// Creates a new Argo Tunnel in an account.
func (r *TunnelService) New(ctx context.Context, params TunnelNewParams, opts ...option.RequestOption) (res *TunnelNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TunnelNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/tunnels", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists and filters all types of Tunnels in an account.
func (r *TunnelService) List(ctx context.Context, params TunnelListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[TunnelListResponse], err error) {
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
func (r *TunnelService) ListAutoPaging(ctx context.Context, params TunnelListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[TunnelListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes an Argo Tunnel from an account.
func (r *TunnelService) Delete(ctx context.Context, tunnelID string, params TunnelDeleteParams, opts ...option.RequestOption) (res *TunnelDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TunnelDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/tunnels/%s", params.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing Cloudflare Tunnel.
func (r *TunnelService) Edit(ctx context.Context, tunnelID string, params TunnelEditParams, opts ...option.RequestOption) (res *TunnelEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TunnelEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s", params.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Argo Tunnel.
func (r *TunnelService) Get(ctx context.Context, tunnelID string, query TunnelGetParams, opts ...option.RequestOption) (res *TunnelGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TunnelGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/tunnels/%s", query.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Connection []ConnectionItem

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type Tunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections Connection `json:"connections"`
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
	TunType TunnelTunType `json:"tun_type"`
	JSON    tunnelJSON    `json:"-"`
}

// tunnelJSON contains the JSON metadata for the struct [Tunnel]
type tunnelJSON struct {
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

func (r *Tunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelJSON) RawJSON() string {
	return r.raw
}

func (r Tunnel) implementsWARPConnectorWARPConnectorNewResponse() {}

func (r Tunnel) implementsWARPConnectorWARPConnectorListResponse() {}

func (r Tunnel) implementsWARPConnectorWARPConnectorDeleteResponse() {}

func (r Tunnel) implementsWARPConnectorWARPConnectorEditResponse() {}

func (r Tunnel) implementsWARPConnectorWARPConnectorGetResponse() {}

func (r Tunnel) implementsZeroTrustTunnelListResponse() {}

func (r Tunnel) implementsZeroTrustTunnelEditResponse() {}

// The type of tunnel.
type TunnelTunType string

const (
	TunnelTunTypeCfdTunnel     TunnelTunType = "cfd_tunnel"
	TunnelTunTypeWARPConnector TunnelTunType = "warp_connector"
	TunnelTunTypeIPSec         TunnelTunType = "ip_sec"
	TunnelTunTypeGRE           TunnelTunType = "gre"
	TunnelTunTypeCni           TunnelTunType = "cni"
)

func (r TunnelTunType) IsKnown() bool {
	switch r {
	case TunnelTunTypeCfdTunnel, TunnelTunTypeWARPConnector, TunnelTunTypeIPSec, TunnelTunTypeGRE, TunnelTunTypeCni:
		return true
	}
	return false
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type TunnelParam struct {
	// Cloudflare account ID
	AccountTag param.Field[string] `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections param.Field[ConnectionParam] `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt param.Field[time.Time] `json:"conns_active_at" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt param.Field[time.Time] `json:"conns_inactive_at" format:"date-time"`
	// Timestamp of when the tunnel was created.
	CreatedAt param.Field[time.Time] `json:"created_at" format:"date-time"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt param.Field[time.Time] `json:"deleted_at" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata param.Field[interface{}] `json:"metadata"`
	// A user-friendly name for the tunnel.
	Name param.Field[string] `json:"name"`
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	RemoteConfig param.Field[bool] `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status param.Field[string] `json:"status"`
	// The type of tunnel.
	TunType param.Field[TunnelTunType] `json:"tun_type"`
}

func (r TunnelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

func (r tunnelNewResponseJSON) RawJSON() string {
	return r.raw
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
	UUID string                          `json:"uuid"`
	JSON tunnelNewResponseConnectionJSON `json:"-"`
}

// tunnelNewResponseConnectionJSON contains the JSON metadata for the struct
// [TunnelNewResponseConnection]
type tunnelNewResponseConnectionJSON struct {
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	UUID               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TunnelNewResponseConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelNewResponseConnectionJSON) RawJSON() string {
	return r.raw
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type TunnelListResponse struct {
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections Connection `json:"connections"`
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
	// UUID of the tunnel.
	ID       string      `json:"id"`
	Metadata interface{} `json:"metadata,required"`
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
	TunType TunnelListResponseTunType `json:"tun_type"`
	JSON    tunnelListResponseJSON    `json:"-"`
	union   TunnelListResponseUnion
}

// tunnelListResponseJSON contains the JSON metadata for the struct
// [TunnelListResponse]
type tunnelListResponseJSON struct {
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	ID              apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	RemoteConfig    apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r tunnelListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *TunnelListResponse) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r TunnelListResponse) AsUnion() TunnelListResponseUnion {
	return r.union
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [zero_trust.Tunnel] or
// [zero_trust.TunnelListResponseTunnelWARPConnectorTunnel].
type TunnelListResponseUnion interface {
	implementsZeroTrustTunnelListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TunnelListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Tunnel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TunnelListResponseTunnelWARPConnectorTunnel{}),
		},
	)
}

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type TunnelListResponseTunnelWARPConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections Connection `json:"connections"`
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
	TunType TunnelListResponseTunnelWARPConnectorTunnelTunType `json:"tun_type"`
	JSON    tunnelListResponseTunnelWARPConnectorTunnelJSON    `json:"-"`
}

// tunnelListResponseTunnelWARPConnectorTunnelJSON contains the JSON metadata for
// the struct [TunnelListResponseTunnelWARPConnectorTunnel]
type tunnelListResponseTunnelWARPConnectorTunnelJSON struct {
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

func (r *TunnelListResponseTunnelWARPConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelListResponseTunnelWARPConnectorTunnelJSON) RawJSON() string {
	return r.raw
}

func (r TunnelListResponseTunnelWARPConnectorTunnel) implementsZeroTrustTunnelListResponse() {}

// The type of tunnel.
type TunnelListResponseTunnelWARPConnectorTunnelTunType string

const (
	TunnelListResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel     TunnelListResponseTunnelWARPConnectorTunnelTunType = "cfd_tunnel"
	TunnelListResponseTunnelWARPConnectorTunnelTunTypeWARPConnector TunnelListResponseTunnelWARPConnectorTunnelTunType = "warp_connector"
	TunnelListResponseTunnelWARPConnectorTunnelTunTypeIPSec         TunnelListResponseTunnelWARPConnectorTunnelTunType = "ip_sec"
	TunnelListResponseTunnelWARPConnectorTunnelTunTypeGRE           TunnelListResponseTunnelWARPConnectorTunnelTunType = "gre"
	TunnelListResponseTunnelWARPConnectorTunnelTunTypeCni           TunnelListResponseTunnelWARPConnectorTunnelTunType = "cni"
)

func (r TunnelListResponseTunnelWARPConnectorTunnelTunType) IsKnown() bool {
	switch r {
	case TunnelListResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel, TunnelListResponseTunnelWARPConnectorTunnelTunTypeWARPConnector, TunnelListResponseTunnelWARPConnectorTunnelTunTypeIPSec, TunnelListResponseTunnelWARPConnectorTunnelTunTypeGRE, TunnelListResponseTunnelWARPConnectorTunnelTunTypeCni:
		return true
	}
	return false
}

// The type of tunnel.
type TunnelListResponseTunType string

const (
	TunnelListResponseTunTypeCfdTunnel     TunnelListResponseTunType = "cfd_tunnel"
	TunnelListResponseTunTypeWARPConnector TunnelListResponseTunType = "warp_connector"
	TunnelListResponseTunTypeIPSec         TunnelListResponseTunType = "ip_sec"
	TunnelListResponseTunTypeGRE           TunnelListResponseTunType = "gre"
	TunnelListResponseTunTypeCni           TunnelListResponseTunType = "cni"
)

func (r TunnelListResponseTunType) IsKnown() bool {
	switch r {
	case TunnelListResponseTunTypeCfdTunnel, TunnelListResponseTunTypeWARPConnector, TunnelListResponseTunTypeIPSec, TunnelListResponseTunTypeGRE, TunnelListResponseTunTypeCni:
		return true
	}
	return false
}

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

func (r tunnelDeleteResponseJSON) RawJSON() string {
	return r.raw
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
	UUID string                             `json:"uuid"`
	JSON tunnelDeleteResponseConnectionJSON `json:"-"`
}

// tunnelDeleteResponseConnectionJSON contains the JSON metadata for the struct
// [TunnelDeleteResponseConnection]
type tunnelDeleteResponseConnectionJSON struct {
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	UUID               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TunnelDeleteResponseConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelDeleteResponseConnectionJSON) RawJSON() string {
	return r.raw
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type TunnelEditResponse struct {
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections Connection `json:"connections"`
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
	// UUID of the tunnel.
	ID       string      `json:"id"`
	Metadata interface{} `json:"metadata,required"`
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
	TunType TunnelEditResponseTunType `json:"tun_type"`
	JSON    tunnelEditResponseJSON    `json:"-"`
	union   TunnelEditResponseUnion
}

// tunnelEditResponseJSON contains the JSON metadata for the struct
// [TunnelEditResponse]
type tunnelEditResponseJSON struct {
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	ID              apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	RemoteConfig    apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r tunnelEditResponseJSON) RawJSON() string {
	return r.raw
}

func (r *TunnelEditResponse) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r TunnelEditResponse) AsUnion() TunnelEditResponseUnion {
	return r.union
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [zero_trust.Tunnel] or
// [zero_trust.TunnelEditResponseTunnelWARPConnectorTunnel].
type TunnelEditResponseUnion interface {
	implementsZeroTrustTunnelEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TunnelEditResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Tunnel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TunnelEditResponseTunnelWARPConnectorTunnel{}),
		},
	)
}

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type TunnelEditResponseTunnelWARPConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections Connection `json:"connections"`
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
	TunType TunnelEditResponseTunnelWARPConnectorTunnelTunType `json:"tun_type"`
	JSON    tunnelEditResponseTunnelWARPConnectorTunnelJSON    `json:"-"`
}

// tunnelEditResponseTunnelWARPConnectorTunnelJSON contains the JSON metadata for
// the struct [TunnelEditResponseTunnelWARPConnectorTunnel]
type tunnelEditResponseTunnelWARPConnectorTunnelJSON struct {
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

func (r *TunnelEditResponseTunnelWARPConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelEditResponseTunnelWARPConnectorTunnelJSON) RawJSON() string {
	return r.raw
}

func (r TunnelEditResponseTunnelWARPConnectorTunnel) implementsZeroTrustTunnelEditResponse() {}

// The type of tunnel.
type TunnelEditResponseTunnelWARPConnectorTunnelTunType string

const (
	TunnelEditResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel     TunnelEditResponseTunnelWARPConnectorTunnelTunType = "cfd_tunnel"
	TunnelEditResponseTunnelWARPConnectorTunnelTunTypeWARPConnector TunnelEditResponseTunnelWARPConnectorTunnelTunType = "warp_connector"
	TunnelEditResponseTunnelWARPConnectorTunnelTunTypeIPSec         TunnelEditResponseTunnelWARPConnectorTunnelTunType = "ip_sec"
	TunnelEditResponseTunnelWARPConnectorTunnelTunTypeGRE           TunnelEditResponseTunnelWARPConnectorTunnelTunType = "gre"
	TunnelEditResponseTunnelWARPConnectorTunnelTunTypeCni           TunnelEditResponseTunnelWARPConnectorTunnelTunType = "cni"
)

func (r TunnelEditResponseTunnelWARPConnectorTunnelTunType) IsKnown() bool {
	switch r {
	case TunnelEditResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel, TunnelEditResponseTunnelWARPConnectorTunnelTunTypeWARPConnector, TunnelEditResponseTunnelWARPConnectorTunnelTunTypeIPSec, TunnelEditResponseTunnelWARPConnectorTunnelTunTypeGRE, TunnelEditResponseTunnelWARPConnectorTunnelTunTypeCni:
		return true
	}
	return false
}

// The type of tunnel.
type TunnelEditResponseTunType string

const (
	TunnelEditResponseTunTypeCfdTunnel     TunnelEditResponseTunType = "cfd_tunnel"
	TunnelEditResponseTunTypeWARPConnector TunnelEditResponseTunType = "warp_connector"
	TunnelEditResponseTunTypeIPSec         TunnelEditResponseTunType = "ip_sec"
	TunnelEditResponseTunTypeGRE           TunnelEditResponseTunType = "gre"
	TunnelEditResponseTunTypeCni           TunnelEditResponseTunType = "cni"
)

func (r TunnelEditResponseTunType) IsKnown() bool {
	switch r {
	case TunnelEditResponseTunTypeCfdTunnel, TunnelEditResponseTunTypeWARPConnector, TunnelEditResponseTunTypeIPSec, TunnelEditResponseTunTypeGRE, TunnelEditResponseTunTypeCni:
		return true
	}
	return false
}

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

func (r tunnelGetResponseJSON) RawJSON() string {
	return r.raw
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
	UUID string                          `json:"uuid"`
	JSON tunnelGetResponseConnectionJSON `json:"-"`
}

// tunnelGetResponseConnectionJSON contains the JSON metadata for the struct
// [TunnelGetResponseConnection]
type tunnelGetResponseConnectionJSON struct {
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	UUID               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TunnelGetResponseConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelGetResponseConnectionJSON) RawJSON() string {
	return r.raw
}

type TunnelNewParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   TunnelNewResponse     `json:"result,required"`
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

func (r tunnelNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelNewResponseEnvelopeSuccess bool

const (
	TunnelNewResponseEnvelopeSuccessTrue TunnelNewResponseEnvelopeSuccess = true
)

func (r TunnelNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TunnelListParams struct {
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
	TunTypes param.Field[string] `query:"tun_types"`
	// UUID of the tunnel.
	UUID          param.Field[string]    `query:"uuid"`
	WasActiveAt   param.Field[time.Time] `query:"was_active_at" format:"date-time"`
	WasInactiveAt param.Field[time.Time] `query:"was_inactive_at" format:"date-time"`
}

// URLQuery serializes [TunnelListParams]'s query parameters as `url.Values`.
func (r TunnelListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TunnelDeleteParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	Body      interface{}         `json:"body,required"`
}

func (r TunnelDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type TunnelDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   TunnelDeleteResponse  `json:"result,required"`
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

func (r tunnelDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelDeleteResponseEnvelopeSuccess bool

const (
	TunnelDeleteResponseEnvelopeSuccessTrue TunnelDeleteResponseEnvelopeSuccess = true
)

func (r TunnelDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TunnelEditParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
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

func (r tunnelEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelEditResponseEnvelopeSuccess bool

const (
	TunnelEditResponseEnvelopeSuccessTrue TunnelEditResponseEnvelopeSuccess = true
)

func (r TunnelEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TunnelGetParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type TunnelGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   TunnelGetResponse     `json:"result,required"`
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

func (r tunnelGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelGetResponseEnvelopeSuccess bool

const (
	TunnelGetResponseEnvelopeSuccessTrue TunnelGetResponseEnvelopeSuccess = true
)

func (r TunnelGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
