// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package warp_connector

import (
	"context"
	"errors"
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
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
)

// WARPConnectorService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWARPConnectorService] method instead.
type WARPConnectorService struct {
	Options []option.RequestOption
}

// NewWARPConnectorService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWARPConnectorService(opts ...option.RequestOption) (r *WARPConnectorService) {
	r = &WARPConnectorService{}
	r.Options = opts
	return
}

// Creates a new Warp Connector Tunnel in an account.
func (r *WARPConnectorService) New(ctx context.Context, params WARPConnectorNewParams, opts ...option.RequestOption) (res *WARPConnectorNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WARPConnectorNewResponseEnvelope
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/warp_connector", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists and filters Warp Connector Tunnels in an account.
func (r *WARPConnectorService) List(ctx context.Context, params WARPConnectorListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[WARPConnectorListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/warp_connector", params.AccountID)
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

// Lists and filters Warp Connector Tunnels in an account.
func (r *WARPConnectorService) ListAutoPaging(ctx context.Context, params WARPConnectorListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[WARPConnectorListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes a Warp Connector Tunnel from an account.
func (r *WARPConnectorService) Delete(ctx context.Context, tunnelID string, body WARPConnectorDeleteParams, opts ...option.RequestOption) (res *WARPConnectorDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WARPConnectorDeleteResponseEnvelope
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/warp_connector/%s", body.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing Warp Connector Tunnel.
func (r *WARPConnectorService) Edit(ctx context.Context, tunnelID string, params WARPConnectorEditParams, opts ...option.RequestOption) (res *WARPConnectorEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WARPConnectorEditResponseEnvelope
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/warp_connector/%s", params.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Warp Connector Tunnel.
func (r *WARPConnectorService) Get(ctx context.Context, tunnelID string, query WARPConnectorGetParams, opts ...option.RequestOption) (res *WARPConnectorGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WARPConnectorGetResponseEnvelope
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/warp_connector/%s", query.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets the token used to associate warp device with a specific Warp Connector
// tunnel.
func (r *WARPConnectorService) Token(ctx context.Context, tunnelID string, query WARPConnectorTokenParams, opts ...option.RequestOption) (res *WARPConnectorTokenResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env WARPConnectorTokenResponseEnvelope
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/warp_connector/%s/token", query.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type WARPConnectorNewResponse struct {
	// Cloudflare account ID
	AccountTag  string      `json:"account_tag"`
	Connections interface{} `json:"connections,required"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at" format:"date-time"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// UUID of the tunnel.
	ID       string      `json:"id" format:"uuid"`
	Metadata interface{} `json:"metadata,required"`
	// A user-friendly name for a tunnel.
	Name string `json:"name"`
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	RemoteConfig bool `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status WARPConnectorNewResponseStatus `json:"status"`
	// The type of tunnel.
	TunType WARPConnectorNewResponseTunType `json:"tun_type"`
	JSON    warpConnectorNewResponseJSON    `json:"-"`
	union   WARPConnectorNewResponseUnion
}

// warpConnectorNewResponseJSON contains the JSON metadata for the struct
// [WARPConnectorNewResponse]
type warpConnectorNewResponseJSON struct {
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

func (r warpConnectorNewResponseJSON) RawJSON() string {
	return r.raw
}

func (r *WARPConnectorNewResponse) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r WARPConnectorNewResponse) AsUnion() WARPConnectorNewResponseUnion {
	return r.union
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [shared.CloudflareTunnel] or
// [warp_connector.WARPConnectorNewResponseTunnelWARPConnectorTunnel].
type WARPConnectorNewResponseUnion interface {
	ImplementsWARPConnectorWARPConnectorNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WARPConnectorNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.CloudflareTunnel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WARPConnectorNewResponseTunnelWARPConnectorTunnel{}),
		},
	)
}

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type WARPConnectorNewResponseTunnelWARPConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id" format:"uuid"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WARPConnectorNewResponseTunnelWARPConnectorTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at" format:"date-time"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for a tunnel.
	Name string `json:"name"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status WARPConnectorNewResponseTunnelWARPConnectorTunnelStatus `json:"status"`
	// The type of tunnel.
	TunType WARPConnectorNewResponseTunnelWARPConnectorTunnelTunType `json:"tun_type"`
	JSON    warpConnectorNewResponseTunnelWARPConnectorTunnelJSON    `json:"-"`
}

// warpConnectorNewResponseTunnelWARPConnectorTunnelJSON contains the JSON metadata
// for the struct [WARPConnectorNewResponseTunnelWARPConnectorTunnel]
type warpConnectorNewResponseTunnelWARPConnectorTunnelJSON struct {
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

func (r *WARPConnectorNewResponseTunnelWARPConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r warpConnectorNewResponseTunnelWARPConnectorTunnelJSON) RawJSON() string {
	return r.raw
}

func (r WARPConnectorNewResponseTunnelWARPConnectorTunnel) ImplementsWARPConnectorWARPConnectorNewResponse() {
}

type WARPConnectorNewResponseTunnelWARPConnectorTunnelConnection struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id" format:"uuid"`
	// UUID of the Cloudflare Tunnel connector.
	ClientID string `json:"client_id" format:"uuid"`
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
	UUID string                                                          `json:"uuid" format:"uuid"`
	JSON warpConnectorNewResponseTunnelWARPConnectorTunnelConnectionJSON `json:"-"`
}

// warpConnectorNewResponseTunnelWARPConnectorTunnelConnectionJSON contains the
// JSON metadata for the struct
// [WARPConnectorNewResponseTunnelWARPConnectorTunnelConnection]
type warpConnectorNewResponseTunnelWARPConnectorTunnelConnectionJSON struct {
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

func (r *WARPConnectorNewResponseTunnelWARPConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r warpConnectorNewResponseTunnelWARPConnectorTunnelConnectionJSON) RawJSON() string {
	return r.raw
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type WARPConnectorNewResponseTunnelWARPConnectorTunnelStatus string

const (
	WARPConnectorNewResponseTunnelWARPConnectorTunnelStatusInactive WARPConnectorNewResponseTunnelWARPConnectorTunnelStatus = "inactive"
	WARPConnectorNewResponseTunnelWARPConnectorTunnelStatusDegraded WARPConnectorNewResponseTunnelWARPConnectorTunnelStatus = "degraded"
	WARPConnectorNewResponseTunnelWARPConnectorTunnelStatusHealthy  WARPConnectorNewResponseTunnelWARPConnectorTunnelStatus = "healthy"
	WARPConnectorNewResponseTunnelWARPConnectorTunnelStatusDown     WARPConnectorNewResponseTunnelWARPConnectorTunnelStatus = "down"
)

func (r WARPConnectorNewResponseTunnelWARPConnectorTunnelStatus) IsKnown() bool {
	switch r {
	case WARPConnectorNewResponseTunnelWARPConnectorTunnelStatusInactive, WARPConnectorNewResponseTunnelWARPConnectorTunnelStatusDegraded, WARPConnectorNewResponseTunnelWARPConnectorTunnelStatusHealthy, WARPConnectorNewResponseTunnelWARPConnectorTunnelStatusDown:
		return true
	}
	return false
}

// The type of tunnel.
type WARPConnectorNewResponseTunnelWARPConnectorTunnelTunType string

const (
	WARPConnectorNewResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel     WARPConnectorNewResponseTunnelWARPConnectorTunnelTunType = "cfd_tunnel"
	WARPConnectorNewResponseTunnelWARPConnectorTunnelTunTypeWARPConnector WARPConnectorNewResponseTunnelWARPConnectorTunnelTunType = "warp_connector"
	WARPConnectorNewResponseTunnelWARPConnectorTunnelTunTypeIPSec         WARPConnectorNewResponseTunnelWARPConnectorTunnelTunType = "ip_sec"
	WARPConnectorNewResponseTunnelWARPConnectorTunnelTunTypeGRE           WARPConnectorNewResponseTunnelWARPConnectorTunnelTunType = "gre"
	WARPConnectorNewResponseTunnelWARPConnectorTunnelTunTypeCNI           WARPConnectorNewResponseTunnelWARPConnectorTunnelTunType = "cni"
)

func (r WARPConnectorNewResponseTunnelWARPConnectorTunnelTunType) IsKnown() bool {
	switch r {
	case WARPConnectorNewResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel, WARPConnectorNewResponseTunnelWARPConnectorTunnelTunTypeWARPConnector, WARPConnectorNewResponseTunnelWARPConnectorTunnelTunTypeIPSec, WARPConnectorNewResponseTunnelWARPConnectorTunnelTunTypeGRE, WARPConnectorNewResponseTunnelWARPConnectorTunnelTunTypeCNI:
		return true
	}
	return false
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type WARPConnectorNewResponseStatus string

const (
	WARPConnectorNewResponseStatusInactive WARPConnectorNewResponseStatus = "inactive"
	WARPConnectorNewResponseStatusDegraded WARPConnectorNewResponseStatus = "degraded"
	WARPConnectorNewResponseStatusHealthy  WARPConnectorNewResponseStatus = "healthy"
	WARPConnectorNewResponseStatusDown     WARPConnectorNewResponseStatus = "down"
)

func (r WARPConnectorNewResponseStatus) IsKnown() bool {
	switch r {
	case WARPConnectorNewResponseStatusInactive, WARPConnectorNewResponseStatusDegraded, WARPConnectorNewResponseStatusHealthy, WARPConnectorNewResponseStatusDown:
		return true
	}
	return false
}

// The type of tunnel.
type WARPConnectorNewResponseTunType string

const (
	WARPConnectorNewResponseTunTypeCfdTunnel     WARPConnectorNewResponseTunType = "cfd_tunnel"
	WARPConnectorNewResponseTunTypeWARPConnector WARPConnectorNewResponseTunType = "warp_connector"
	WARPConnectorNewResponseTunTypeIPSec         WARPConnectorNewResponseTunType = "ip_sec"
	WARPConnectorNewResponseTunTypeGRE           WARPConnectorNewResponseTunType = "gre"
	WARPConnectorNewResponseTunTypeCNI           WARPConnectorNewResponseTunType = "cni"
)

func (r WARPConnectorNewResponseTunType) IsKnown() bool {
	switch r {
	case WARPConnectorNewResponseTunTypeCfdTunnel, WARPConnectorNewResponseTunTypeWARPConnector, WARPConnectorNewResponseTunTypeIPSec, WARPConnectorNewResponseTunTypeGRE, WARPConnectorNewResponseTunTypeCNI:
		return true
	}
	return false
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type WARPConnectorListResponse struct {
	// Cloudflare account ID
	AccountTag  string      `json:"account_tag"`
	Connections interface{} `json:"connections,required"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at" format:"date-time"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// UUID of the tunnel.
	ID       string      `json:"id" format:"uuid"`
	Metadata interface{} `json:"metadata,required"`
	// A user-friendly name for a tunnel.
	Name string `json:"name"`
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	RemoteConfig bool `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status WARPConnectorListResponseStatus `json:"status"`
	// The type of tunnel.
	TunType WARPConnectorListResponseTunType `json:"tun_type"`
	JSON    warpConnectorListResponseJSON    `json:"-"`
	union   WARPConnectorListResponseUnion
}

// warpConnectorListResponseJSON contains the JSON metadata for the struct
// [WARPConnectorListResponse]
type warpConnectorListResponseJSON struct {
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

func (r warpConnectorListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *WARPConnectorListResponse) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r WARPConnectorListResponse) AsUnion() WARPConnectorListResponseUnion {
	return r.union
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [shared.CloudflareTunnel] or
// [warp_connector.WARPConnectorListResponseTunnelWARPConnectorTunnel].
type WARPConnectorListResponseUnion interface {
	ImplementsWARPConnectorWARPConnectorListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WARPConnectorListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.CloudflareTunnel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WARPConnectorListResponseTunnelWARPConnectorTunnel{}),
		},
	)
}

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type WARPConnectorListResponseTunnelWARPConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id" format:"uuid"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WARPConnectorListResponseTunnelWARPConnectorTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at" format:"date-time"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for a tunnel.
	Name string `json:"name"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status WARPConnectorListResponseTunnelWARPConnectorTunnelStatus `json:"status"`
	// The type of tunnel.
	TunType WARPConnectorListResponseTunnelWARPConnectorTunnelTunType `json:"tun_type"`
	JSON    warpConnectorListResponseTunnelWARPConnectorTunnelJSON    `json:"-"`
}

// warpConnectorListResponseTunnelWARPConnectorTunnelJSON contains the JSON
// metadata for the struct [WARPConnectorListResponseTunnelWARPConnectorTunnel]
type warpConnectorListResponseTunnelWARPConnectorTunnelJSON struct {
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

func (r *WARPConnectorListResponseTunnelWARPConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r warpConnectorListResponseTunnelWARPConnectorTunnelJSON) RawJSON() string {
	return r.raw
}

func (r WARPConnectorListResponseTunnelWARPConnectorTunnel) ImplementsWARPConnectorWARPConnectorListResponse() {
}

type WARPConnectorListResponseTunnelWARPConnectorTunnelConnection struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id" format:"uuid"`
	// UUID of the Cloudflare Tunnel connector.
	ClientID string `json:"client_id" format:"uuid"`
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
	UUID string                                                           `json:"uuid" format:"uuid"`
	JSON warpConnectorListResponseTunnelWARPConnectorTunnelConnectionJSON `json:"-"`
}

// warpConnectorListResponseTunnelWARPConnectorTunnelConnectionJSON contains the
// JSON metadata for the struct
// [WARPConnectorListResponseTunnelWARPConnectorTunnelConnection]
type warpConnectorListResponseTunnelWARPConnectorTunnelConnectionJSON struct {
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

func (r *WARPConnectorListResponseTunnelWARPConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r warpConnectorListResponseTunnelWARPConnectorTunnelConnectionJSON) RawJSON() string {
	return r.raw
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type WARPConnectorListResponseTunnelWARPConnectorTunnelStatus string

const (
	WARPConnectorListResponseTunnelWARPConnectorTunnelStatusInactive WARPConnectorListResponseTunnelWARPConnectorTunnelStatus = "inactive"
	WARPConnectorListResponseTunnelWARPConnectorTunnelStatusDegraded WARPConnectorListResponseTunnelWARPConnectorTunnelStatus = "degraded"
	WARPConnectorListResponseTunnelWARPConnectorTunnelStatusHealthy  WARPConnectorListResponseTunnelWARPConnectorTunnelStatus = "healthy"
	WARPConnectorListResponseTunnelWARPConnectorTunnelStatusDown     WARPConnectorListResponseTunnelWARPConnectorTunnelStatus = "down"
)

func (r WARPConnectorListResponseTunnelWARPConnectorTunnelStatus) IsKnown() bool {
	switch r {
	case WARPConnectorListResponseTunnelWARPConnectorTunnelStatusInactive, WARPConnectorListResponseTunnelWARPConnectorTunnelStatusDegraded, WARPConnectorListResponseTunnelWARPConnectorTunnelStatusHealthy, WARPConnectorListResponseTunnelWARPConnectorTunnelStatusDown:
		return true
	}
	return false
}

// The type of tunnel.
type WARPConnectorListResponseTunnelWARPConnectorTunnelTunType string

const (
	WARPConnectorListResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel     WARPConnectorListResponseTunnelWARPConnectorTunnelTunType = "cfd_tunnel"
	WARPConnectorListResponseTunnelWARPConnectorTunnelTunTypeWARPConnector WARPConnectorListResponseTunnelWARPConnectorTunnelTunType = "warp_connector"
	WARPConnectorListResponseTunnelWARPConnectorTunnelTunTypeIPSec         WARPConnectorListResponseTunnelWARPConnectorTunnelTunType = "ip_sec"
	WARPConnectorListResponseTunnelWARPConnectorTunnelTunTypeGRE           WARPConnectorListResponseTunnelWARPConnectorTunnelTunType = "gre"
	WARPConnectorListResponseTunnelWARPConnectorTunnelTunTypeCNI           WARPConnectorListResponseTunnelWARPConnectorTunnelTunType = "cni"
)

func (r WARPConnectorListResponseTunnelWARPConnectorTunnelTunType) IsKnown() bool {
	switch r {
	case WARPConnectorListResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel, WARPConnectorListResponseTunnelWARPConnectorTunnelTunTypeWARPConnector, WARPConnectorListResponseTunnelWARPConnectorTunnelTunTypeIPSec, WARPConnectorListResponseTunnelWARPConnectorTunnelTunTypeGRE, WARPConnectorListResponseTunnelWARPConnectorTunnelTunTypeCNI:
		return true
	}
	return false
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type WARPConnectorListResponseStatus string

const (
	WARPConnectorListResponseStatusInactive WARPConnectorListResponseStatus = "inactive"
	WARPConnectorListResponseStatusDegraded WARPConnectorListResponseStatus = "degraded"
	WARPConnectorListResponseStatusHealthy  WARPConnectorListResponseStatus = "healthy"
	WARPConnectorListResponseStatusDown     WARPConnectorListResponseStatus = "down"
)

func (r WARPConnectorListResponseStatus) IsKnown() bool {
	switch r {
	case WARPConnectorListResponseStatusInactive, WARPConnectorListResponseStatusDegraded, WARPConnectorListResponseStatusHealthy, WARPConnectorListResponseStatusDown:
		return true
	}
	return false
}

// The type of tunnel.
type WARPConnectorListResponseTunType string

const (
	WARPConnectorListResponseTunTypeCfdTunnel     WARPConnectorListResponseTunType = "cfd_tunnel"
	WARPConnectorListResponseTunTypeWARPConnector WARPConnectorListResponseTunType = "warp_connector"
	WARPConnectorListResponseTunTypeIPSec         WARPConnectorListResponseTunType = "ip_sec"
	WARPConnectorListResponseTunTypeGRE           WARPConnectorListResponseTunType = "gre"
	WARPConnectorListResponseTunTypeCNI           WARPConnectorListResponseTunType = "cni"
)

func (r WARPConnectorListResponseTunType) IsKnown() bool {
	switch r {
	case WARPConnectorListResponseTunTypeCfdTunnel, WARPConnectorListResponseTunTypeWARPConnector, WARPConnectorListResponseTunTypeIPSec, WARPConnectorListResponseTunTypeGRE, WARPConnectorListResponseTunTypeCNI:
		return true
	}
	return false
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type WARPConnectorDeleteResponse struct {
	// Cloudflare account ID
	AccountTag  string      `json:"account_tag"`
	Connections interface{} `json:"connections,required"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at" format:"date-time"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// UUID of the tunnel.
	ID       string      `json:"id" format:"uuid"`
	Metadata interface{} `json:"metadata,required"`
	// A user-friendly name for a tunnel.
	Name string `json:"name"`
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	RemoteConfig bool `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status WARPConnectorDeleteResponseStatus `json:"status"`
	// The type of tunnel.
	TunType WARPConnectorDeleteResponseTunType `json:"tun_type"`
	JSON    warpConnectorDeleteResponseJSON    `json:"-"`
	union   WARPConnectorDeleteResponseUnion
}

// warpConnectorDeleteResponseJSON contains the JSON metadata for the struct
// [WARPConnectorDeleteResponse]
type warpConnectorDeleteResponseJSON struct {
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

func (r warpConnectorDeleteResponseJSON) RawJSON() string {
	return r.raw
}

func (r *WARPConnectorDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r WARPConnectorDeleteResponse) AsUnion() WARPConnectorDeleteResponseUnion {
	return r.union
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [shared.CloudflareTunnel] or
// [warp_connector.WARPConnectorDeleteResponseTunnelWARPConnectorTunnel].
type WARPConnectorDeleteResponseUnion interface {
	ImplementsWARPConnectorWARPConnectorDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WARPConnectorDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.CloudflareTunnel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WARPConnectorDeleteResponseTunnelWARPConnectorTunnel{}),
		},
	)
}

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type WARPConnectorDeleteResponseTunnelWARPConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id" format:"uuid"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WARPConnectorDeleteResponseTunnelWARPConnectorTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at" format:"date-time"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for a tunnel.
	Name string `json:"name"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status WARPConnectorDeleteResponseTunnelWARPConnectorTunnelStatus `json:"status"`
	// The type of tunnel.
	TunType WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunType `json:"tun_type"`
	JSON    warpConnectorDeleteResponseTunnelWARPConnectorTunnelJSON    `json:"-"`
}

// warpConnectorDeleteResponseTunnelWARPConnectorTunnelJSON contains the JSON
// metadata for the struct [WARPConnectorDeleteResponseTunnelWARPConnectorTunnel]
type warpConnectorDeleteResponseTunnelWARPConnectorTunnelJSON struct {
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

func (r *WARPConnectorDeleteResponseTunnelWARPConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r warpConnectorDeleteResponseTunnelWARPConnectorTunnelJSON) RawJSON() string {
	return r.raw
}

func (r WARPConnectorDeleteResponseTunnelWARPConnectorTunnel) ImplementsWARPConnectorWARPConnectorDeleteResponse() {
}

type WARPConnectorDeleteResponseTunnelWARPConnectorTunnelConnection struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id" format:"uuid"`
	// UUID of the Cloudflare Tunnel connector.
	ClientID string `json:"client_id" format:"uuid"`
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
	UUID string                                                             `json:"uuid" format:"uuid"`
	JSON warpConnectorDeleteResponseTunnelWARPConnectorTunnelConnectionJSON `json:"-"`
}

// warpConnectorDeleteResponseTunnelWARPConnectorTunnelConnectionJSON contains the
// JSON metadata for the struct
// [WARPConnectorDeleteResponseTunnelWARPConnectorTunnelConnection]
type warpConnectorDeleteResponseTunnelWARPConnectorTunnelConnectionJSON struct {
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

func (r *WARPConnectorDeleteResponseTunnelWARPConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r warpConnectorDeleteResponseTunnelWARPConnectorTunnelConnectionJSON) RawJSON() string {
	return r.raw
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type WARPConnectorDeleteResponseTunnelWARPConnectorTunnelStatus string

const (
	WARPConnectorDeleteResponseTunnelWARPConnectorTunnelStatusInactive WARPConnectorDeleteResponseTunnelWARPConnectorTunnelStatus = "inactive"
	WARPConnectorDeleteResponseTunnelWARPConnectorTunnelStatusDegraded WARPConnectorDeleteResponseTunnelWARPConnectorTunnelStatus = "degraded"
	WARPConnectorDeleteResponseTunnelWARPConnectorTunnelStatusHealthy  WARPConnectorDeleteResponseTunnelWARPConnectorTunnelStatus = "healthy"
	WARPConnectorDeleteResponseTunnelWARPConnectorTunnelStatusDown     WARPConnectorDeleteResponseTunnelWARPConnectorTunnelStatus = "down"
)

func (r WARPConnectorDeleteResponseTunnelWARPConnectorTunnelStatus) IsKnown() bool {
	switch r {
	case WARPConnectorDeleteResponseTunnelWARPConnectorTunnelStatusInactive, WARPConnectorDeleteResponseTunnelWARPConnectorTunnelStatusDegraded, WARPConnectorDeleteResponseTunnelWARPConnectorTunnelStatusHealthy, WARPConnectorDeleteResponseTunnelWARPConnectorTunnelStatusDown:
		return true
	}
	return false
}

// The type of tunnel.
type WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunType string

const (
	WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel     WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunType = "cfd_tunnel"
	WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunTypeWARPConnector WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunType = "warp_connector"
	WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunTypeIPSec         WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunType = "ip_sec"
	WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunTypeGRE           WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunType = "gre"
	WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunTypeCNI           WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunType = "cni"
)

func (r WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunType) IsKnown() bool {
	switch r {
	case WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel, WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunTypeWARPConnector, WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunTypeIPSec, WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunTypeGRE, WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunTypeCNI:
		return true
	}
	return false
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type WARPConnectorDeleteResponseStatus string

const (
	WARPConnectorDeleteResponseStatusInactive WARPConnectorDeleteResponseStatus = "inactive"
	WARPConnectorDeleteResponseStatusDegraded WARPConnectorDeleteResponseStatus = "degraded"
	WARPConnectorDeleteResponseStatusHealthy  WARPConnectorDeleteResponseStatus = "healthy"
	WARPConnectorDeleteResponseStatusDown     WARPConnectorDeleteResponseStatus = "down"
)

func (r WARPConnectorDeleteResponseStatus) IsKnown() bool {
	switch r {
	case WARPConnectorDeleteResponseStatusInactive, WARPConnectorDeleteResponseStatusDegraded, WARPConnectorDeleteResponseStatusHealthy, WARPConnectorDeleteResponseStatusDown:
		return true
	}
	return false
}

// The type of tunnel.
type WARPConnectorDeleteResponseTunType string

const (
	WARPConnectorDeleteResponseTunTypeCfdTunnel     WARPConnectorDeleteResponseTunType = "cfd_tunnel"
	WARPConnectorDeleteResponseTunTypeWARPConnector WARPConnectorDeleteResponseTunType = "warp_connector"
	WARPConnectorDeleteResponseTunTypeIPSec         WARPConnectorDeleteResponseTunType = "ip_sec"
	WARPConnectorDeleteResponseTunTypeGRE           WARPConnectorDeleteResponseTunType = "gre"
	WARPConnectorDeleteResponseTunTypeCNI           WARPConnectorDeleteResponseTunType = "cni"
)

func (r WARPConnectorDeleteResponseTunType) IsKnown() bool {
	switch r {
	case WARPConnectorDeleteResponseTunTypeCfdTunnel, WARPConnectorDeleteResponseTunTypeWARPConnector, WARPConnectorDeleteResponseTunTypeIPSec, WARPConnectorDeleteResponseTunTypeGRE, WARPConnectorDeleteResponseTunTypeCNI:
		return true
	}
	return false
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type WARPConnectorEditResponse struct {
	// Cloudflare account ID
	AccountTag  string      `json:"account_tag"`
	Connections interface{} `json:"connections,required"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at" format:"date-time"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// UUID of the tunnel.
	ID       string      `json:"id" format:"uuid"`
	Metadata interface{} `json:"metadata,required"`
	// A user-friendly name for a tunnel.
	Name string `json:"name"`
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	RemoteConfig bool `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status WARPConnectorEditResponseStatus `json:"status"`
	// The type of tunnel.
	TunType WARPConnectorEditResponseTunType `json:"tun_type"`
	JSON    warpConnectorEditResponseJSON    `json:"-"`
	union   WARPConnectorEditResponseUnion
}

// warpConnectorEditResponseJSON contains the JSON metadata for the struct
// [WARPConnectorEditResponse]
type warpConnectorEditResponseJSON struct {
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

func (r warpConnectorEditResponseJSON) RawJSON() string {
	return r.raw
}

func (r *WARPConnectorEditResponse) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r WARPConnectorEditResponse) AsUnion() WARPConnectorEditResponseUnion {
	return r.union
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [shared.CloudflareTunnel] or
// [warp_connector.WARPConnectorEditResponseTunnelWARPConnectorTunnel].
type WARPConnectorEditResponseUnion interface {
	ImplementsWARPConnectorWARPConnectorEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WARPConnectorEditResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.CloudflareTunnel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WARPConnectorEditResponseTunnelWARPConnectorTunnel{}),
		},
	)
}

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type WARPConnectorEditResponseTunnelWARPConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id" format:"uuid"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WARPConnectorEditResponseTunnelWARPConnectorTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at" format:"date-time"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for a tunnel.
	Name string `json:"name"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status WARPConnectorEditResponseTunnelWARPConnectorTunnelStatus `json:"status"`
	// The type of tunnel.
	TunType WARPConnectorEditResponseTunnelWARPConnectorTunnelTunType `json:"tun_type"`
	JSON    warpConnectorEditResponseTunnelWARPConnectorTunnelJSON    `json:"-"`
}

// warpConnectorEditResponseTunnelWARPConnectorTunnelJSON contains the JSON
// metadata for the struct [WARPConnectorEditResponseTunnelWARPConnectorTunnel]
type warpConnectorEditResponseTunnelWARPConnectorTunnelJSON struct {
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

func (r *WARPConnectorEditResponseTunnelWARPConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r warpConnectorEditResponseTunnelWARPConnectorTunnelJSON) RawJSON() string {
	return r.raw
}

func (r WARPConnectorEditResponseTunnelWARPConnectorTunnel) ImplementsWARPConnectorWARPConnectorEditResponse() {
}

type WARPConnectorEditResponseTunnelWARPConnectorTunnelConnection struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id" format:"uuid"`
	// UUID of the Cloudflare Tunnel connector.
	ClientID string `json:"client_id" format:"uuid"`
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
	UUID string                                                           `json:"uuid" format:"uuid"`
	JSON warpConnectorEditResponseTunnelWARPConnectorTunnelConnectionJSON `json:"-"`
}

// warpConnectorEditResponseTunnelWARPConnectorTunnelConnectionJSON contains the
// JSON metadata for the struct
// [WARPConnectorEditResponseTunnelWARPConnectorTunnelConnection]
type warpConnectorEditResponseTunnelWARPConnectorTunnelConnectionJSON struct {
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

func (r *WARPConnectorEditResponseTunnelWARPConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r warpConnectorEditResponseTunnelWARPConnectorTunnelConnectionJSON) RawJSON() string {
	return r.raw
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type WARPConnectorEditResponseTunnelWARPConnectorTunnelStatus string

const (
	WARPConnectorEditResponseTunnelWARPConnectorTunnelStatusInactive WARPConnectorEditResponseTunnelWARPConnectorTunnelStatus = "inactive"
	WARPConnectorEditResponseTunnelWARPConnectorTunnelStatusDegraded WARPConnectorEditResponseTunnelWARPConnectorTunnelStatus = "degraded"
	WARPConnectorEditResponseTunnelWARPConnectorTunnelStatusHealthy  WARPConnectorEditResponseTunnelWARPConnectorTunnelStatus = "healthy"
	WARPConnectorEditResponseTunnelWARPConnectorTunnelStatusDown     WARPConnectorEditResponseTunnelWARPConnectorTunnelStatus = "down"
)

func (r WARPConnectorEditResponseTunnelWARPConnectorTunnelStatus) IsKnown() bool {
	switch r {
	case WARPConnectorEditResponseTunnelWARPConnectorTunnelStatusInactive, WARPConnectorEditResponseTunnelWARPConnectorTunnelStatusDegraded, WARPConnectorEditResponseTunnelWARPConnectorTunnelStatusHealthy, WARPConnectorEditResponseTunnelWARPConnectorTunnelStatusDown:
		return true
	}
	return false
}

// The type of tunnel.
type WARPConnectorEditResponseTunnelWARPConnectorTunnelTunType string

const (
	WARPConnectorEditResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel     WARPConnectorEditResponseTunnelWARPConnectorTunnelTunType = "cfd_tunnel"
	WARPConnectorEditResponseTunnelWARPConnectorTunnelTunTypeWARPConnector WARPConnectorEditResponseTunnelWARPConnectorTunnelTunType = "warp_connector"
	WARPConnectorEditResponseTunnelWARPConnectorTunnelTunTypeIPSec         WARPConnectorEditResponseTunnelWARPConnectorTunnelTunType = "ip_sec"
	WARPConnectorEditResponseTunnelWARPConnectorTunnelTunTypeGRE           WARPConnectorEditResponseTunnelWARPConnectorTunnelTunType = "gre"
	WARPConnectorEditResponseTunnelWARPConnectorTunnelTunTypeCNI           WARPConnectorEditResponseTunnelWARPConnectorTunnelTunType = "cni"
)

func (r WARPConnectorEditResponseTunnelWARPConnectorTunnelTunType) IsKnown() bool {
	switch r {
	case WARPConnectorEditResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel, WARPConnectorEditResponseTunnelWARPConnectorTunnelTunTypeWARPConnector, WARPConnectorEditResponseTunnelWARPConnectorTunnelTunTypeIPSec, WARPConnectorEditResponseTunnelWARPConnectorTunnelTunTypeGRE, WARPConnectorEditResponseTunnelWARPConnectorTunnelTunTypeCNI:
		return true
	}
	return false
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type WARPConnectorEditResponseStatus string

const (
	WARPConnectorEditResponseStatusInactive WARPConnectorEditResponseStatus = "inactive"
	WARPConnectorEditResponseStatusDegraded WARPConnectorEditResponseStatus = "degraded"
	WARPConnectorEditResponseStatusHealthy  WARPConnectorEditResponseStatus = "healthy"
	WARPConnectorEditResponseStatusDown     WARPConnectorEditResponseStatus = "down"
)

func (r WARPConnectorEditResponseStatus) IsKnown() bool {
	switch r {
	case WARPConnectorEditResponseStatusInactive, WARPConnectorEditResponseStatusDegraded, WARPConnectorEditResponseStatusHealthy, WARPConnectorEditResponseStatusDown:
		return true
	}
	return false
}

// The type of tunnel.
type WARPConnectorEditResponseTunType string

const (
	WARPConnectorEditResponseTunTypeCfdTunnel     WARPConnectorEditResponseTunType = "cfd_tunnel"
	WARPConnectorEditResponseTunTypeWARPConnector WARPConnectorEditResponseTunType = "warp_connector"
	WARPConnectorEditResponseTunTypeIPSec         WARPConnectorEditResponseTunType = "ip_sec"
	WARPConnectorEditResponseTunTypeGRE           WARPConnectorEditResponseTunType = "gre"
	WARPConnectorEditResponseTunTypeCNI           WARPConnectorEditResponseTunType = "cni"
)

func (r WARPConnectorEditResponseTunType) IsKnown() bool {
	switch r {
	case WARPConnectorEditResponseTunTypeCfdTunnel, WARPConnectorEditResponseTunTypeWARPConnector, WARPConnectorEditResponseTunTypeIPSec, WARPConnectorEditResponseTunTypeGRE, WARPConnectorEditResponseTunTypeCNI:
		return true
	}
	return false
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type WARPConnectorGetResponse struct {
	// Cloudflare account ID
	AccountTag  string      `json:"account_tag"`
	Connections interface{} `json:"connections,required"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at" format:"date-time"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// UUID of the tunnel.
	ID       string      `json:"id" format:"uuid"`
	Metadata interface{} `json:"metadata,required"`
	// A user-friendly name for a tunnel.
	Name string `json:"name"`
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	RemoteConfig bool `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status WARPConnectorGetResponseStatus `json:"status"`
	// The type of tunnel.
	TunType WARPConnectorGetResponseTunType `json:"tun_type"`
	JSON    warpConnectorGetResponseJSON    `json:"-"`
	union   WARPConnectorGetResponseUnion
}

// warpConnectorGetResponseJSON contains the JSON metadata for the struct
// [WARPConnectorGetResponse]
type warpConnectorGetResponseJSON struct {
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

func (r warpConnectorGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *WARPConnectorGetResponse) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r WARPConnectorGetResponse) AsUnion() WARPConnectorGetResponseUnion {
	return r.union
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [shared.CloudflareTunnel] or
// [warp_connector.WARPConnectorGetResponseTunnelWARPConnectorTunnel].
type WARPConnectorGetResponseUnion interface {
	ImplementsWARPConnectorWARPConnectorGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WARPConnectorGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.CloudflareTunnel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WARPConnectorGetResponseTunnelWARPConnectorTunnel{}),
		},
	)
}

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type WARPConnectorGetResponseTunnelWARPConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id" format:"uuid"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WARPConnectorGetResponseTunnelWARPConnectorTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at" format:"date-time"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for a tunnel.
	Name string `json:"name"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status WARPConnectorGetResponseTunnelWARPConnectorTunnelStatus `json:"status"`
	// The type of tunnel.
	TunType WARPConnectorGetResponseTunnelWARPConnectorTunnelTunType `json:"tun_type"`
	JSON    warpConnectorGetResponseTunnelWARPConnectorTunnelJSON    `json:"-"`
}

// warpConnectorGetResponseTunnelWARPConnectorTunnelJSON contains the JSON metadata
// for the struct [WARPConnectorGetResponseTunnelWARPConnectorTunnel]
type warpConnectorGetResponseTunnelWARPConnectorTunnelJSON struct {
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

func (r *WARPConnectorGetResponseTunnelWARPConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r warpConnectorGetResponseTunnelWARPConnectorTunnelJSON) RawJSON() string {
	return r.raw
}

func (r WARPConnectorGetResponseTunnelWARPConnectorTunnel) ImplementsWARPConnectorWARPConnectorGetResponse() {
}

type WARPConnectorGetResponseTunnelWARPConnectorTunnelConnection struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id" format:"uuid"`
	// UUID of the Cloudflare Tunnel connector.
	ClientID string `json:"client_id" format:"uuid"`
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
	UUID string                                                          `json:"uuid" format:"uuid"`
	JSON warpConnectorGetResponseTunnelWARPConnectorTunnelConnectionJSON `json:"-"`
}

// warpConnectorGetResponseTunnelWARPConnectorTunnelConnectionJSON contains the
// JSON metadata for the struct
// [WARPConnectorGetResponseTunnelWARPConnectorTunnelConnection]
type warpConnectorGetResponseTunnelWARPConnectorTunnelConnectionJSON struct {
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

func (r *WARPConnectorGetResponseTunnelWARPConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r warpConnectorGetResponseTunnelWARPConnectorTunnelConnectionJSON) RawJSON() string {
	return r.raw
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type WARPConnectorGetResponseTunnelWARPConnectorTunnelStatus string

const (
	WARPConnectorGetResponseTunnelWARPConnectorTunnelStatusInactive WARPConnectorGetResponseTunnelWARPConnectorTunnelStatus = "inactive"
	WARPConnectorGetResponseTunnelWARPConnectorTunnelStatusDegraded WARPConnectorGetResponseTunnelWARPConnectorTunnelStatus = "degraded"
	WARPConnectorGetResponseTunnelWARPConnectorTunnelStatusHealthy  WARPConnectorGetResponseTunnelWARPConnectorTunnelStatus = "healthy"
	WARPConnectorGetResponseTunnelWARPConnectorTunnelStatusDown     WARPConnectorGetResponseTunnelWARPConnectorTunnelStatus = "down"
)

func (r WARPConnectorGetResponseTunnelWARPConnectorTunnelStatus) IsKnown() bool {
	switch r {
	case WARPConnectorGetResponseTunnelWARPConnectorTunnelStatusInactive, WARPConnectorGetResponseTunnelWARPConnectorTunnelStatusDegraded, WARPConnectorGetResponseTunnelWARPConnectorTunnelStatusHealthy, WARPConnectorGetResponseTunnelWARPConnectorTunnelStatusDown:
		return true
	}
	return false
}

// The type of tunnel.
type WARPConnectorGetResponseTunnelWARPConnectorTunnelTunType string

const (
	WARPConnectorGetResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel     WARPConnectorGetResponseTunnelWARPConnectorTunnelTunType = "cfd_tunnel"
	WARPConnectorGetResponseTunnelWARPConnectorTunnelTunTypeWARPConnector WARPConnectorGetResponseTunnelWARPConnectorTunnelTunType = "warp_connector"
	WARPConnectorGetResponseTunnelWARPConnectorTunnelTunTypeIPSec         WARPConnectorGetResponseTunnelWARPConnectorTunnelTunType = "ip_sec"
	WARPConnectorGetResponseTunnelWARPConnectorTunnelTunTypeGRE           WARPConnectorGetResponseTunnelWARPConnectorTunnelTunType = "gre"
	WARPConnectorGetResponseTunnelWARPConnectorTunnelTunTypeCNI           WARPConnectorGetResponseTunnelWARPConnectorTunnelTunType = "cni"
)

func (r WARPConnectorGetResponseTunnelWARPConnectorTunnelTunType) IsKnown() bool {
	switch r {
	case WARPConnectorGetResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel, WARPConnectorGetResponseTunnelWARPConnectorTunnelTunTypeWARPConnector, WARPConnectorGetResponseTunnelWARPConnectorTunnelTunTypeIPSec, WARPConnectorGetResponseTunnelWARPConnectorTunnelTunTypeGRE, WARPConnectorGetResponseTunnelWARPConnectorTunnelTunTypeCNI:
		return true
	}
	return false
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type WARPConnectorGetResponseStatus string

const (
	WARPConnectorGetResponseStatusInactive WARPConnectorGetResponseStatus = "inactive"
	WARPConnectorGetResponseStatusDegraded WARPConnectorGetResponseStatus = "degraded"
	WARPConnectorGetResponseStatusHealthy  WARPConnectorGetResponseStatus = "healthy"
	WARPConnectorGetResponseStatusDown     WARPConnectorGetResponseStatus = "down"
)

func (r WARPConnectorGetResponseStatus) IsKnown() bool {
	switch r {
	case WARPConnectorGetResponseStatusInactive, WARPConnectorGetResponseStatusDegraded, WARPConnectorGetResponseStatusHealthy, WARPConnectorGetResponseStatusDown:
		return true
	}
	return false
}

// The type of tunnel.
type WARPConnectorGetResponseTunType string

const (
	WARPConnectorGetResponseTunTypeCfdTunnel     WARPConnectorGetResponseTunType = "cfd_tunnel"
	WARPConnectorGetResponseTunTypeWARPConnector WARPConnectorGetResponseTunType = "warp_connector"
	WARPConnectorGetResponseTunTypeIPSec         WARPConnectorGetResponseTunType = "ip_sec"
	WARPConnectorGetResponseTunTypeGRE           WARPConnectorGetResponseTunType = "gre"
	WARPConnectorGetResponseTunTypeCNI           WARPConnectorGetResponseTunType = "cni"
)

func (r WARPConnectorGetResponseTunType) IsKnown() bool {
	switch r {
	case WARPConnectorGetResponseTunTypeCfdTunnel, WARPConnectorGetResponseTunTypeWARPConnector, WARPConnectorGetResponseTunTypeIPSec, WARPConnectorGetResponseTunTypeGRE, WARPConnectorGetResponseTunTypeCNI:
		return true
	}
	return false
}

// Union satisfied by [warp_connector.WARPConnectorTokenResponseUnknown],
// [warp_connector.WARPConnectorTokenResponseArray] or [shared.UnionString].
type WARPConnectorTokenResponseUnion interface {
	ImplementsWARPConnectorWARPConnectorTokenResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WARPConnectorTokenResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WARPConnectorTokenResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type WARPConnectorTokenResponseArray []interface{}

func (r WARPConnectorTokenResponseArray) ImplementsWARPConnectorWARPConnectorTokenResponseUnion() {}

type WARPConnectorNewParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// A user-friendly name for a tunnel.
	Name param.Field[string] `json:"name,required"`
}

func (r WARPConnectorNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WARPConnectorNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result WARPConnectorNewResponse `json:"result,required"`
	// Whether the API call was successful
	Success WARPConnectorNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    warpConnectorNewResponseEnvelopeJSON    `json:"-"`
}

// warpConnectorNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [WARPConnectorNewResponseEnvelope]
type warpConnectorNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WARPConnectorNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r warpConnectorNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WARPConnectorNewResponseEnvelopeSuccess bool

const (
	WARPConnectorNewResponseEnvelopeSuccessTrue WARPConnectorNewResponseEnvelopeSuccess = true
)

func (r WARPConnectorNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WARPConnectorNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type WARPConnectorListParams struct {
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
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status param.Field[WARPConnectorListParamsStatus] `query:"status"`
	// UUID of the tunnel.
	UUID          param.Field[string]    `query:"uuid" format:"uuid"`
	WasActiveAt   param.Field[time.Time] `query:"was_active_at" format:"date-time"`
	WasInactiveAt param.Field[time.Time] `query:"was_inactive_at" format:"date-time"`
}

// URLQuery serializes [WARPConnectorListParams]'s query parameters as
// `url.Values`.
func (r WARPConnectorListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type WARPConnectorListParamsStatus string

const (
	WARPConnectorListParamsStatusInactive WARPConnectorListParamsStatus = "inactive"
	WARPConnectorListParamsStatusDegraded WARPConnectorListParamsStatus = "degraded"
	WARPConnectorListParamsStatusHealthy  WARPConnectorListParamsStatus = "healthy"
	WARPConnectorListParamsStatusDown     WARPConnectorListParamsStatus = "down"
)

func (r WARPConnectorListParamsStatus) IsKnown() bool {
	switch r {
	case WARPConnectorListParamsStatusInactive, WARPConnectorListParamsStatusDegraded, WARPConnectorListParamsStatusHealthy, WARPConnectorListParamsStatusDown:
		return true
	}
	return false
}

type WARPConnectorDeleteParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type WARPConnectorDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result WARPConnectorDeleteResponse `json:"result,required"`
	// Whether the API call was successful
	Success WARPConnectorDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    warpConnectorDeleteResponseEnvelopeJSON    `json:"-"`
}

// warpConnectorDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [WARPConnectorDeleteResponseEnvelope]
type warpConnectorDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WARPConnectorDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r warpConnectorDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WARPConnectorDeleteResponseEnvelopeSuccess bool

const (
	WARPConnectorDeleteResponseEnvelopeSuccessTrue WARPConnectorDeleteResponseEnvelopeSuccess = true
)

func (r WARPConnectorDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WARPConnectorDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type WARPConnectorEditParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// A user-friendly name for a tunnel.
	Name param.Field[string] `json:"name"`
	// Sets the password required to run a locally-managed tunnel. Must be at least 32
	// bytes and encoded as a base64 string.
	TunnelSecret param.Field[string] `json:"tunnel_secret"`
}

func (r WARPConnectorEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WARPConnectorEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result WARPConnectorEditResponse `json:"result,required"`
	// Whether the API call was successful
	Success WARPConnectorEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    warpConnectorEditResponseEnvelopeJSON    `json:"-"`
}

// warpConnectorEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [WARPConnectorEditResponseEnvelope]
type warpConnectorEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WARPConnectorEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r warpConnectorEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WARPConnectorEditResponseEnvelopeSuccess bool

const (
	WARPConnectorEditResponseEnvelopeSuccessTrue WARPConnectorEditResponseEnvelopeSuccess = true
)

func (r WARPConnectorEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WARPConnectorEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type WARPConnectorGetParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type WARPConnectorGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result WARPConnectorGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success WARPConnectorGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    warpConnectorGetResponseEnvelopeJSON    `json:"-"`
}

// warpConnectorGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [WARPConnectorGetResponseEnvelope]
type warpConnectorGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WARPConnectorGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r warpConnectorGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WARPConnectorGetResponseEnvelopeSuccess bool

const (
	WARPConnectorGetResponseEnvelopeSuccessTrue WARPConnectorGetResponseEnvelopeSuccess = true
)

func (r WARPConnectorGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WARPConnectorGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type WARPConnectorTokenParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type WARPConnectorTokenResponseEnvelope struct {
	Errors   []shared.ResponseInfo           `json:"errors,required"`
	Messages []shared.ResponseInfo           `json:"messages,required"`
	Result   WARPConnectorTokenResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success WARPConnectorTokenResponseEnvelopeSuccess `json:"success,required"`
	JSON    warpConnectorTokenResponseEnvelopeJSON    `json:"-"`
}

// warpConnectorTokenResponseEnvelopeJSON contains the JSON metadata for the struct
// [WARPConnectorTokenResponseEnvelope]
type warpConnectorTokenResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WARPConnectorTokenResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r warpConnectorTokenResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WARPConnectorTokenResponseEnvelopeSuccess bool

const (
	WARPConnectorTokenResponseEnvelopeSuccessTrue WARPConnectorTokenResponseEnvelopeSuccess = true
)

func (r WARPConnectorTokenResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WARPConnectorTokenResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
