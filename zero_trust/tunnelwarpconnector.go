// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v6/shared"
)

// TunnelWARPConnectorService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTunnelWARPConnectorService] method instead.
type TunnelWARPConnectorService struct {
	Options []option.RequestOption
	Token   *TunnelWARPConnectorTokenService
}

// NewTunnelWARPConnectorService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTunnelWARPConnectorService(opts ...option.RequestOption) (r *TunnelWARPConnectorService) {
	r = &TunnelWARPConnectorService{}
	r.Options = opts
	r.Token = NewTunnelWARPConnectorTokenService(opts...)
	return
}

// Creates a new Warp Connector Tunnel in an account.
func (r *TunnelWARPConnectorService) New(ctx context.Context, params TunnelWARPConnectorNewParams, opts ...option.RequestOption) (res *TunnelWARPConnectorNewResponse, err error) {
	var env TunnelWARPConnectorNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
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
func (r *TunnelWARPConnectorService) List(ctx context.Context, params TunnelWARPConnectorListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[TunnelWARPConnectorListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
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
func (r *TunnelWARPConnectorService) ListAutoPaging(ctx context.Context, params TunnelWARPConnectorListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[TunnelWARPConnectorListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes a Warp Connector Tunnel from an account.
func (r *TunnelWARPConnectorService) Delete(ctx context.Context, tunnelID string, body TunnelWARPConnectorDeleteParams, opts ...option.RequestOption) (res *TunnelWARPConnectorDeleteResponse, err error) {
	var env TunnelWARPConnectorDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
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
func (r *TunnelWARPConnectorService) Edit(ctx context.Context, tunnelID string, params TunnelWARPConnectorEditParams, opts ...option.RequestOption) (res *TunnelWARPConnectorEditResponse, err error) {
	var env TunnelWARPConnectorEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
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
func (r *TunnelWARPConnectorService) Get(ctx context.Context, tunnelID string, query TunnelWARPConnectorGetParams, opts ...option.RequestOption) (res *TunnelWARPConnectorGetResponse, err error) {
	var env TunnelWARPConnectorGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
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

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type TunnelWARPConnectorNewResponse struct {
	// UUID of the tunnel.
	ID string `json:"id" format:"uuid"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	//
	// Deprecated: This field will start returning an empty array. To fetch the
	// connections of a given tunnel, please use the dedicated endpoint
	// `/accounts/{account_id}/{tunnel_type}/{tunnel_id}/connections`
	Connections []TunnelWARPConnectorNewResponseConnection `json:"connections"`
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
	Status TunnelWARPConnectorNewResponseStatus `json:"status"`
	// The type of tunnel.
	TunType TunnelWARPConnectorNewResponseTunType `json:"tun_type"`
	JSON    tunnelWARPConnectorNewResponseJSON    `json:"-"`
}

// tunnelWARPConnectorNewResponseJSON contains the JSON metadata for the struct
// [TunnelWARPConnectorNewResponse]
type tunnelWARPConnectorNewResponseJSON struct {
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

func (r *TunnelWARPConnectorNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorNewResponseJSON) RawJSON() string {
	return r.raw
}

type TunnelWARPConnectorNewResponseConnection struct {
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
	UUID string                                       `json:"uuid" format:"uuid"`
	JSON tunnelWARPConnectorNewResponseConnectionJSON `json:"-"`
}

// tunnelWARPConnectorNewResponseConnectionJSON contains the JSON metadata for the
// struct [TunnelWARPConnectorNewResponseConnection]
type tunnelWARPConnectorNewResponseConnectionJSON struct {
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

func (r *TunnelWARPConnectorNewResponseConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorNewResponseConnectionJSON) RawJSON() string {
	return r.raw
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type TunnelWARPConnectorNewResponseStatus string

const (
	TunnelWARPConnectorNewResponseStatusInactive TunnelWARPConnectorNewResponseStatus = "inactive"
	TunnelWARPConnectorNewResponseStatusDegraded TunnelWARPConnectorNewResponseStatus = "degraded"
	TunnelWARPConnectorNewResponseStatusHealthy  TunnelWARPConnectorNewResponseStatus = "healthy"
	TunnelWARPConnectorNewResponseStatusDown     TunnelWARPConnectorNewResponseStatus = "down"
)

func (r TunnelWARPConnectorNewResponseStatus) IsKnown() bool {
	switch r {
	case TunnelWARPConnectorNewResponseStatusInactive, TunnelWARPConnectorNewResponseStatusDegraded, TunnelWARPConnectorNewResponseStatusHealthy, TunnelWARPConnectorNewResponseStatusDown:
		return true
	}
	return false
}

// The type of tunnel.
type TunnelWARPConnectorNewResponseTunType string

const (
	TunnelWARPConnectorNewResponseTunTypeCfdTunnel     TunnelWARPConnectorNewResponseTunType = "cfd_tunnel"
	TunnelWARPConnectorNewResponseTunTypeWARPConnector TunnelWARPConnectorNewResponseTunType = "warp_connector"
	TunnelWARPConnectorNewResponseTunTypeWARP          TunnelWARPConnectorNewResponseTunType = "warp"
	TunnelWARPConnectorNewResponseTunTypeMagic         TunnelWARPConnectorNewResponseTunType = "magic"
	TunnelWARPConnectorNewResponseTunTypeIPSec         TunnelWARPConnectorNewResponseTunType = "ip_sec"
	TunnelWARPConnectorNewResponseTunTypeGRE           TunnelWARPConnectorNewResponseTunType = "gre"
	TunnelWARPConnectorNewResponseTunTypeCNI           TunnelWARPConnectorNewResponseTunType = "cni"
)

func (r TunnelWARPConnectorNewResponseTunType) IsKnown() bool {
	switch r {
	case TunnelWARPConnectorNewResponseTunTypeCfdTunnel, TunnelWARPConnectorNewResponseTunTypeWARPConnector, TunnelWARPConnectorNewResponseTunTypeWARP, TunnelWARPConnectorNewResponseTunTypeMagic, TunnelWARPConnectorNewResponseTunTypeIPSec, TunnelWARPConnectorNewResponseTunTypeGRE, TunnelWARPConnectorNewResponseTunTypeCNI:
		return true
	}
	return false
}

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type TunnelWARPConnectorListResponse struct {
	// UUID of the tunnel.
	ID string `json:"id" format:"uuid"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	//
	// Deprecated: This field will start returning an empty array. To fetch the
	// connections of a given tunnel, please use the dedicated endpoint
	// `/accounts/{account_id}/{tunnel_type}/{tunnel_id}/connections`
	Connections []TunnelWARPConnectorListResponseConnection `json:"connections"`
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
	Status TunnelWARPConnectorListResponseStatus `json:"status"`
	// The type of tunnel.
	TunType TunnelWARPConnectorListResponseTunType `json:"tun_type"`
	JSON    tunnelWARPConnectorListResponseJSON    `json:"-"`
}

// tunnelWARPConnectorListResponseJSON contains the JSON metadata for the struct
// [TunnelWARPConnectorListResponse]
type tunnelWARPConnectorListResponseJSON struct {
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

func (r *TunnelWARPConnectorListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorListResponseJSON) RawJSON() string {
	return r.raw
}

type TunnelWARPConnectorListResponseConnection struct {
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
	UUID string                                        `json:"uuid" format:"uuid"`
	JSON tunnelWARPConnectorListResponseConnectionJSON `json:"-"`
}

// tunnelWARPConnectorListResponseConnectionJSON contains the JSON metadata for the
// struct [TunnelWARPConnectorListResponseConnection]
type tunnelWARPConnectorListResponseConnectionJSON struct {
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

func (r *TunnelWARPConnectorListResponseConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorListResponseConnectionJSON) RawJSON() string {
	return r.raw
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type TunnelWARPConnectorListResponseStatus string

const (
	TunnelWARPConnectorListResponseStatusInactive TunnelWARPConnectorListResponseStatus = "inactive"
	TunnelWARPConnectorListResponseStatusDegraded TunnelWARPConnectorListResponseStatus = "degraded"
	TunnelWARPConnectorListResponseStatusHealthy  TunnelWARPConnectorListResponseStatus = "healthy"
	TunnelWARPConnectorListResponseStatusDown     TunnelWARPConnectorListResponseStatus = "down"
)

func (r TunnelWARPConnectorListResponseStatus) IsKnown() bool {
	switch r {
	case TunnelWARPConnectorListResponseStatusInactive, TunnelWARPConnectorListResponseStatusDegraded, TunnelWARPConnectorListResponseStatusHealthy, TunnelWARPConnectorListResponseStatusDown:
		return true
	}
	return false
}

// The type of tunnel.
type TunnelWARPConnectorListResponseTunType string

const (
	TunnelWARPConnectorListResponseTunTypeCfdTunnel     TunnelWARPConnectorListResponseTunType = "cfd_tunnel"
	TunnelWARPConnectorListResponseTunTypeWARPConnector TunnelWARPConnectorListResponseTunType = "warp_connector"
	TunnelWARPConnectorListResponseTunTypeWARP          TunnelWARPConnectorListResponseTunType = "warp"
	TunnelWARPConnectorListResponseTunTypeMagic         TunnelWARPConnectorListResponseTunType = "magic"
	TunnelWARPConnectorListResponseTunTypeIPSec         TunnelWARPConnectorListResponseTunType = "ip_sec"
	TunnelWARPConnectorListResponseTunTypeGRE           TunnelWARPConnectorListResponseTunType = "gre"
	TunnelWARPConnectorListResponseTunTypeCNI           TunnelWARPConnectorListResponseTunType = "cni"
)

func (r TunnelWARPConnectorListResponseTunType) IsKnown() bool {
	switch r {
	case TunnelWARPConnectorListResponseTunTypeCfdTunnel, TunnelWARPConnectorListResponseTunTypeWARPConnector, TunnelWARPConnectorListResponseTunTypeWARP, TunnelWARPConnectorListResponseTunTypeMagic, TunnelWARPConnectorListResponseTunTypeIPSec, TunnelWARPConnectorListResponseTunTypeGRE, TunnelWARPConnectorListResponseTunTypeCNI:
		return true
	}
	return false
}

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type TunnelWARPConnectorDeleteResponse struct {
	// UUID of the tunnel.
	ID string `json:"id" format:"uuid"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	//
	// Deprecated: This field will start returning an empty array. To fetch the
	// connections of a given tunnel, please use the dedicated endpoint
	// `/accounts/{account_id}/{tunnel_type}/{tunnel_id}/connections`
	Connections []TunnelWARPConnectorDeleteResponseConnection `json:"connections"`
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
	Status TunnelWARPConnectorDeleteResponseStatus `json:"status"`
	// The type of tunnel.
	TunType TunnelWARPConnectorDeleteResponseTunType `json:"tun_type"`
	JSON    tunnelWARPConnectorDeleteResponseJSON    `json:"-"`
}

// tunnelWARPConnectorDeleteResponseJSON contains the JSON metadata for the struct
// [TunnelWARPConnectorDeleteResponse]
type tunnelWARPConnectorDeleteResponseJSON struct {
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

func (r *TunnelWARPConnectorDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type TunnelWARPConnectorDeleteResponseConnection struct {
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
	UUID string                                          `json:"uuid" format:"uuid"`
	JSON tunnelWARPConnectorDeleteResponseConnectionJSON `json:"-"`
}

// tunnelWARPConnectorDeleteResponseConnectionJSON contains the JSON metadata for
// the struct [TunnelWARPConnectorDeleteResponseConnection]
type tunnelWARPConnectorDeleteResponseConnectionJSON struct {
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

func (r *TunnelWARPConnectorDeleteResponseConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorDeleteResponseConnectionJSON) RawJSON() string {
	return r.raw
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type TunnelWARPConnectorDeleteResponseStatus string

const (
	TunnelWARPConnectorDeleteResponseStatusInactive TunnelWARPConnectorDeleteResponseStatus = "inactive"
	TunnelWARPConnectorDeleteResponseStatusDegraded TunnelWARPConnectorDeleteResponseStatus = "degraded"
	TunnelWARPConnectorDeleteResponseStatusHealthy  TunnelWARPConnectorDeleteResponseStatus = "healthy"
	TunnelWARPConnectorDeleteResponseStatusDown     TunnelWARPConnectorDeleteResponseStatus = "down"
)

func (r TunnelWARPConnectorDeleteResponseStatus) IsKnown() bool {
	switch r {
	case TunnelWARPConnectorDeleteResponseStatusInactive, TunnelWARPConnectorDeleteResponseStatusDegraded, TunnelWARPConnectorDeleteResponseStatusHealthy, TunnelWARPConnectorDeleteResponseStatusDown:
		return true
	}
	return false
}

// The type of tunnel.
type TunnelWARPConnectorDeleteResponseTunType string

const (
	TunnelWARPConnectorDeleteResponseTunTypeCfdTunnel     TunnelWARPConnectorDeleteResponseTunType = "cfd_tunnel"
	TunnelWARPConnectorDeleteResponseTunTypeWARPConnector TunnelWARPConnectorDeleteResponseTunType = "warp_connector"
	TunnelWARPConnectorDeleteResponseTunTypeWARP          TunnelWARPConnectorDeleteResponseTunType = "warp"
	TunnelWARPConnectorDeleteResponseTunTypeMagic         TunnelWARPConnectorDeleteResponseTunType = "magic"
	TunnelWARPConnectorDeleteResponseTunTypeIPSec         TunnelWARPConnectorDeleteResponseTunType = "ip_sec"
	TunnelWARPConnectorDeleteResponseTunTypeGRE           TunnelWARPConnectorDeleteResponseTunType = "gre"
	TunnelWARPConnectorDeleteResponseTunTypeCNI           TunnelWARPConnectorDeleteResponseTunType = "cni"
)

func (r TunnelWARPConnectorDeleteResponseTunType) IsKnown() bool {
	switch r {
	case TunnelWARPConnectorDeleteResponseTunTypeCfdTunnel, TunnelWARPConnectorDeleteResponseTunTypeWARPConnector, TunnelWARPConnectorDeleteResponseTunTypeWARP, TunnelWARPConnectorDeleteResponseTunTypeMagic, TunnelWARPConnectorDeleteResponseTunTypeIPSec, TunnelWARPConnectorDeleteResponseTunTypeGRE, TunnelWARPConnectorDeleteResponseTunTypeCNI:
		return true
	}
	return false
}

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type TunnelWARPConnectorEditResponse struct {
	// UUID of the tunnel.
	ID string `json:"id" format:"uuid"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	//
	// Deprecated: This field will start returning an empty array. To fetch the
	// connections of a given tunnel, please use the dedicated endpoint
	// `/accounts/{account_id}/{tunnel_type}/{tunnel_id}/connections`
	Connections []TunnelWARPConnectorEditResponseConnection `json:"connections"`
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
	Status TunnelWARPConnectorEditResponseStatus `json:"status"`
	// The type of tunnel.
	TunType TunnelWARPConnectorEditResponseTunType `json:"tun_type"`
	JSON    tunnelWARPConnectorEditResponseJSON    `json:"-"`
}

// tunnelWARPConnectorEditResponseJSON contains the JSON metadata for the struct
// [TunnelWARPConnectorEditResponse]
type tunnelWARPConnectorEditResponseJSON struct {
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

func (r *TunnelWARPConnectorEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorEditResponseJSON) RawJSON() string {
	return r.raw
}

type TunnelWARPConnectorEditResponseConnection struct {
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
	UUID string                                        `json:"uuid" format:"uuid"`
	JSON tunnelWARPConnectorEditResponseConnectionJSON `json:"-"`
}

// tunnelWARPConnectorEditResponseConnectionJSON contains the JSON metadata for the
// struct [TunnelWARPConnectorEditResponseConnection]
type tunnelWARPConnectorEditResponseConnectionJSON struct {
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

func (r *TunnelWARPConnectorEditResponseConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorEditResponseConnectionJSON) RawJSON() string {
	return r.raw
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type TunnelWARPConnectorEditResponseStatus string

const (
	TunnelWARPConnectorEditResponseStatusInactive TunnelWARPConnectorEditResponseStatus = "inactive"
	TunnelWARPConnectorEditResponseStatusDegraded TunnelWARPConnectorEditResponseStatus = "degraded"
	TunnelWARPConnectorEditResponseStatusHealthy  TunnelWARPConnectorEditResponseStatus = "healthy"
	TunnelWARPConnectorEditResponseStatusDown     TunnelWARPConnectorEditResponseStatus = "down"
)

func (r TunnelWARPConnectorEditResponseStatus) IsKnown() bool {
	switch r {
	case TunnelWARPConnectorEditResponseStatusInactive, TunnelWARPConnectorEditResponseStatusDegraded, TunnelWARPConnectorEditResponseStatusHealthy, TunnelWARPConnectorEditResponseStatusDown:
		return true
	}
	return false
}

// The type of tunnel.
type TunnelWARPConnectorEditResponseTunType string

const (
	TunnelWARPConnectorEditResponseTunTypeCfdTunnel     TunnelWARPConnectorEditResponseTunType = "cfd_tunnel"
	TunnelWARPConnectorEditResponseTunTypeWARPConnector TunnelWARPConnectorEditResponseTunType = "warp_connector"
	TunnelWARPConnectorEditResponseTunTypeWARP          TunnelWARPConnectorEditResponseTunType = "warp"
	TunnelWARPConnectorEditResponseTunTypeMagic         TunnelWARPConnectorEditResponseTunType = "magic"
	TunnelWARPConnectorEditResponseTunTypeIPSec         TunnelWARPConnectorEditResponseTunType = "ip_sec"
	TunnelWARPConnectorEditResponseTunTypeGRE           TunnelWARPConnectorEditResponseTunType = "gre"
	TunnelWARPConnectorEditResponseTunTypeCNI           TunnelWARPConnectorEditResponseTunType = "cni"
)

func (r TunnelWARPConnectorEditResponseTunType) IsKnown() bool {
	switch r {
	case TunnelWARPConnectorEditResponseTunTypeCfdTunnel, TunnelWARPConnectorEditResponseTunTypeWARPConnector, TunnelWARPConnectorEditResponseTunTypeWARP, TunnelWARPConnectorEditResponseTunTypeMagic, TunnelWARPConnectorEditResponseTunTypeIPSec, TunnelWARPConnectorEditResponseTunTypeGRE, TunnelWARPConnectorEditResponseTunTypeCNI:
		return true
	}
	return false
}

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type TunnelWARPConnectorGetResponse struct {
	// UUID of the tunnel.
	ID string `json:"id" format:"uuid"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	//
	// Deprecated: This field will start returning an empty array. To fetch the
	// connections of a given tunnel, please use the dedicated endpoint
	// `/accounts/{account_id}/{tunnel_type}/{tunnel_id}/connections`
	Connections []TunnelWARPConnectorGetResponseConnection `json:"connections"`
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
	Status TunnelWARPConnectorGetResponseStatus `json:"status"`
	// The type of tunnel.
	TunType TunnelWARPConnectorGetResponseTunType `json:"tun_type"`
	JSON    tunnelWARPConnectorGetResponseJSON    `json:"-"`
}

// tunnelWARPConnectorGetResponseJSON contains the JSON metadata for the struct
// [TunnelWARPConnectorGetResponse]
type tunnelWARPConnectorGetResponseJSON struct {
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

func (r *TunnelWARPConnectorGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorGetResponseJSON) RawJSON() string {
	return r.raw
}

type TunnelWARPConnectorGetResponseConnection struct {
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
	UUID string                                       `json:"uuid" format:"uuid"`
	JSON tunnelWARPConnectorGetResponseConnectionJSON `json:"-"`
}

// tunnelWARPConnectorGetResponseConnectionJSON contains the JSON metadata for the
// struct [TunnelWARPConnectorGetResponseConnection]
type tunnelWARPConnectorGetResponseConnectionJSON struct {
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

func (r *TunnelWARPConnectorGetResponseConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorGetResponseConnectionJSON) RawJSON() string {
	return r.raw
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type TunnelWARPConnectorGetResponseStatus string

const (
	TunnelWARPConnectorGetResponseStatusInactive TunnelWARPConnectorGetResponseStatus = "inactive"
	TunnelWARPConnectorGetResponseStatusDegraded TunnelWARPConnectorGetResponseStatus = "degraded"
	TunnelWARPConnectorGetResponseStatusHealthy  TunnelWARPConnectorGetResponseStatus = "healthy"
	TunnelWARPConnectorGetResponseStatusDown     TunnelWARPConnectorGetResponseStatus = "down"
)

func (r TunnelWARPConnectorGetResponseStatus) IsKnown() bool {
	switch r {
	case TunnelWARPConnectorGetResponseStatusInactive, TunnelWARPConnectorGetResponseStatusDegraded, TunnelWARPConnectorGetResponseStatusHealthy, TunnelWARPConnectorGetResponseStatusDown:
		return true
	}
	return false
}

// The type of tunnel.
type TunnelWARPConnectorGetResponseTunType string

const (
	TunnelWARPConnectorGetResponseTunTypeCfdTunnel     TunnelWARPConnectorGetResponseTunType = "cfd_tunnel"
	TunnelWARPConnectorGetResponseTunTypeWARPConnector TunnelWARPConnectorGetResponseTunType = "warp_connector"
	TunnelWARPConnectorGetResponseTunTypeWARP          TunnelWARPConnectorGetResponseTunType = "warp"
	TunnelWARPConnectorGetResponseTunTypeMagic         TunnelWARPConnectorGetResponseTunType = "magic"
	TunnelWARPConnectorGetResponseTunTypeIPSec         TunnelWARPConnectorGetResponseTunType = "ip_sec"
	TunnelWARPConnectorGetResponseTunTypeGRE           TunnelWARPConnectorGetResponseTunType = "gre"
	TunnelWARPConnectorGetResponseTunTypeCNI           TunnelWARPConnectorGetResponseTunType = "cni"
)

func (r TunnelWARPConnectorGetResponseTunType) IsKnown() bool {
	switch r {
	case TunnelWARPConnectorGetResponseTunTypeCfdTunnel, TunnelWARPConnectorGetResponseTunTypeWARPConnector, TunnelWARPConnectorGetResponseTunTypeWARP, TunnelWARPConnectorGetResponseTunTypeMagic, TunnelWARPConnectorGetResponseTunTypeIPSec, TunnelWARPConnectorGetResponseTunTypeGRE, TunnelWARPConnectorGetResponseTunTypeCNI:
		return true
	}
	return false
}

type TunnelWARPConnectorNewParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// A user-friendly name for a tunnel.
	Name param.Field[string] `json:"name,required"`
}

func (r TunnelWARPConnectorNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TunnelWARPConnectorNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
	Result TunnelWARPConnectorNewResponse `json:"result,required"`
	// Whether the API call was successful
	Success TunnelWARPConnectorNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelWARPConnectorNewResponseEnvelopeJSON    `json:"-"`
}

// tunnelWARPConnectorNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [TunnelWARPConnectorNewResponseEnvelope]
type tunnelWARPConnectorNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelWARPConnectorNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelWARPConnectorNewResponseEnvelopeSuccess bool

const (
	TunnelWARPConnectorNewResponseEnvelopeSuccessTrue TunnelWARPConnectorNewResponseEnvelopeSuccess = true
)

func (r TunnelWARPConnectorNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelWARPConnectorNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TunnelWARPConnectorListParams struct {
	// Cloudflare account ID
	AccountID     param.Field[string] `path:"account_id,required"`
	ExcludePrefix param.Field[string] `query:"exclude_prefix"`
	// If provided, include only resources that were created (and not deleted) before
	// this time. URL encoded.
	ExistedAt     param.Field[string] `query:"existed_at" format:"url-encoded-date-time"`
	IncludePrefix param.Field[string] `query:"include_prefix"`
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
	Status param.Field[TunnelWARPConnectorListParamsStatus] `query:"status"`
	// UUID of the tunnel.
	UUID          param.Field[string]    `query:"uuid" format:"uuid"`
	WasActiveAt   param.Field[time.Time] `query:"was_active_at" format:"date-time"`
	WasInactiveAt param.Field[time.Time] `query:"was_inactive_at" format:"date-time"`
}

// URLQuery serializes [TunnelWARPConnectorListParams]'s query parameters as
// `url.Values`.
func (r TunnelWARPConnectorListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type TunnelWARPConnectorListParamsStatus string

const (
	TunnelWARPConnectorListParamsStatusInactive TunnelWARPConnectorListParamsStatus = "inactive"
	TunnelWARPConnectorListParamsStatusDegraded TunnelWARPConnectorListParamsStatus = "degraded"
	TunnelWARPConnectorListParamsStatusHealthy  TunnelWARPConnectorListParamsStatus = "healthy"
	TunnelWARPConnectorListParamsStatusDown     TunnelWARPConnectorListParamsStatus = "down"
)

func (r TunnelWARPConnectorListParamsStatus) IsKnown() bool {
	switch r {
	case TunnelWARPConnectorListParamsStatusInactive, TunnelWARPConnectorListParamsStatusDegraded, TunnelWARPConnectorListParamsStatusHealthy, TunnelWARPConnectorListParamsStatusDown:
		return true
	}
	return false
}

type TunnelWARPConnectorDeleteParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type TunnelWARPConnectorDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
	Result TunnelWARPConnectorDeleteResponse `json:"result,required"`
	// Whether the API call was successful
	Success TunnelWARPConnectorDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelWARPConnectorDeleteResponseEnvelopeJSON    `json:"-"`
}

// tunnelWARPConnectorDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [TunnelWARPConnectorDeleteResponseEnvelope]
type tunnelWARPConnectorDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelWARPConnectorDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelWARPConnectorDeleteResponseEnvelopeSuccess bool

const (
	TunnelWARPConnectorDeleteResponseEnvelopeSuccessTrue TunnelWARPConnectorDeleteResponseEnvelopeSuccess = true
)

func (r TunnelWARPConnectorDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelWARPConnectorDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TunnelWARPConnectorEditParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// A user-friendly name for a tunnel.
	Name param.Field[string] `json:"name"`
	// Sets the password required to run a locally-managed tunnel. Must be at least 32
	// bytes and encoded as a base64 string.
	TunnelSecret param.Field[string] `json:"tunnel_secret"`
}

func (r TunnelWARPConnectorEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TunnelWARPConnectorEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
	Result TunnelWARPConnectorEditResponse `json:"result,required"`
	// Whether the API call was successful
	Success TunnelWARPConnectorEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelWARPConnectorEditResponseEnvelopeJSON    `json:"-"`
}

// tunnelWARPConnectorEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [TunnelWARPConnectorEditResponseEnvelope]
type tunnelWARPConnectorEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelWARPConnectorEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelWARPConnectorEditResponseEnvelopeSuccess bool

const (
	TunnelWARPConnectorEditResponseEnvelopeSuccessTrue TunnelWARPConnectorEditResponseEnvelopeSuccess = true
)

func (r TunnelWARPConnectorEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelWARPConnectorEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TunnelWARPConnectorGetParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type TunnelWARPConnectorGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
	Result TunnelWARPConnectorGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success TunnelWARPConnectorGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelWARPConnectorGetResponseEnvelopeJSON    `json:"-"`
}

// tunnelWARPConnectorGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [TunnelWARPConnectorGetResponseEnvelope]
type tunnelWARPConnectorGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelWARPConnectorGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelWARPConnectorGetResponseEnvelopeSuccess bool

const (
	TunnelWARPConnectorGetResponseEnvelopeSuccessTrue TunnelWARPConnectorGetResponseEnvelopeSuccess = true
)

func (r TunnelWARPConnectorGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelWARPConnectorGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
