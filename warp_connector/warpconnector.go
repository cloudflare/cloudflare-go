// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package warp_connector

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
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/tidwall/gjson"
)

// WARPConnectorService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWARPConnectorService] method
// instead.
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
func (r *WARPConnectorService) New(ctx context.Context, params WARPConnectorNewParams, opts ...option.RequestOption) (res *WARPConnector, err error) {
	opts = append(r.Options[:], opts...)
	var env WARPConnectorNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/warp_connector", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists and filters Warp Connector Tunnels in an account.
func (r *WARPConnectorService) List(ctx context.Context, params WARPConnectorListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[WARPConnector], err error) {
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
func (r *WARPConnectorService) ListAutoPaging(ctx context.Context, params WARPConnectorListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[WARPConnector] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes a Warp Connector Tunnel from an account.
func (r *WARPConnectorService) Delete(ctx context.Context, tunnelID string, params WARPConnectorDeleteParams, opts ...option.RequestOption) (res *WARPConnector, err error) {
	opts = append(r.Options[:], opts...)
	var env WARPConnectorDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/warp_connector/%s", params.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing Warp Connector Tunnel.
func (r *WARPConnectorService) Edit(ctx context.Context, tunnelID string, params WARPConnectorEditParams, opts ...option.RequestOption) (res *WARPConnector, err error) {
	opts = append(r.Options[:], opts...)
	var env WARPConnectorEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/warp_connector/%s", params.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Warp Connector Tunnel.
func (r *WARPConnectorService) Get(ctx context.Context, tunnelID string, query WARPConnectorGetParams, opts ...option.RequestOption) (res *WARPConnector, err error) {
	opts = append(r.Options[:], opts...)
	var env WARPConnectorGetResponseEnvelope
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
	path := fmt.Sprintf("accounts/%s/warp_connector/%s/token", query.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type WARPConnector struct {
	// Cloudflare account ID
	AccountTag  string      `json:"account_tag"`
	Connections interface{} `json:"connections,required"`
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
	TunType WARPConnectorTunType `json:"tun_type"`
	JSON    warpConnectorJSON    `json:"-"`
	union   WARPConnectorUnion
}

// warpConnectorJSON contains the JSON metadata for the struct [WARPConnector]
type warpConnectorJSON struct {
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

func (r warpConnectorJSON) RawJSON() string {
	return r.raw
}

func (r *WARPConnector) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r WARPConnector) AsUnion() WARPConnectorUnion {
	return r.union
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [warp_connector.WARPConnectorTunnelCfdTunnel] or
// [warp_connector.WARPConnectorTunnelWARPConnectorTunnel].
type WARPConnectorUnion interface {
	implementsWARPConnectorWARPConnector()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WARPConnectorUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WARPConnectorTunnelCfdTunnel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WARPConnectorTunnelWARPConnectorTunnel{}),
		},
	)
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type WARPConnectorTunnelCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []zero_trust.TunnelConnection `json:"connections"`
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
	TunType WARPConnectorTunnelCfdTunnelTunType `json:"tun_type"`
	JSON    warpConnectorTunnelCfdTunnelJSON    `json:"-"`
}

// warpConnectorTunnelCfdTunnelJSON contains the JSON metadata for the struct
// [WARPConnectorTunnelCfdTunnel]
type warpConnectorTunnelCfdTunnelJSON struct {
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

func (r *WARPConnectorTunnelCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r warpConnectorTunnelCfdTunnelJSON) RawJSON() string {
	return r.raw
}

func (r WARPConnectorTunnelCfdTunnel) implementsWARPConnectorWARPConnector() {}

// The type of tunnel.
type WARPConnectorTunnelCfdTunnelTunType string

const (
	WARPConnectorTunnelCfdTunnelTunTypeCfdTunnel     WARPConnectorTunnelCfdTunnelTunType = "cfd_tunnel"
	WARPConnectorTunnelCfdTunnelTunTypeWARPConnector WARPConnectorTunnelCfdTunnelTunType = "warp_connector"
	WARPConnectorTunnelCfdTunnelTunTypeIPSec         WARPConnectorTunnelCfdTunnelTunType = "ip_sec"
	WARPConnectorTunnelCfdTunnelTunTypeGRE           WARPConnectorTunnelCfdTunnelTunType = "gre"
	WARPConnectorTunnelCfdTunnelTunTypeCni           WARPConnectorTunnelCfdTunnelTunType = "cni"
)

func (r WARPConnectorTunnelCfdTunnelTunType) IsKnown() bool {
	switch r {
	case WARPConnectorTunnelCfdTunnelTunTypeCfdTunnel, WARPConnectorTunnelCfdTunnelTunTypeWARPConnector, WARPConnectorTunnelCfdTunnelTunTypeIPSec, WARPConnectorTunnelCfdTunnelTunTypeGRE, WARPConnectorTunnelCfdTunnelTunTypeCni:
		return true
	}
	return false
}

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type WARPConnectorTunnelWARPConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []zero_trust.TunnelConnection `json:"connections"`
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
	TunType WARPConnectorTunnelWARPConnectorTunnelTunType `json:"tun_type"`
	JSON    warpConnectorTunnelWARPConnectorTunnelJSON    `json:"-"`
}

// warpConnectorTunnelWARPConnectorTunnelJSON contains the JSON metadata for the
// struct [WARPConnectorTunnelWARPConnectorTunnel]
type warpConnectorTunnelWARPConnectorTunnelJSON struct {
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

func (r *WARPConnectorTunnelWARPConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r warpConnectorTunnelWARPConnectorTunnelJSON) RawJSON() string {
	return r.raw
}

func (r WARPConnectorTunnelWARPConnectorTunnel) implementsWARPConnectorWARPConnector() {}

// The type of tunnel.
type WARPConnectorTunnelWARPConnectorTunnelTunType string

const (
	WARPConnectorTunnelWARPConnectorTunnelTunTypeCfdTunnel     WARPConnectorTunnelWARPConnectorTunnelTunType = "cfd_tunnel"
	WARPConnectorTunnelWARPConnectorTunnelTunTypeWARPConnector WARPConnectorTunnelWARPConnectorTunnelTunType = "warp_connector"
	WARPConnectorTunnelWARPConnectorTunnelTunTypeIPSec         WARPConnectorTunnelWARPConnectorTunnelTunType = "ip_sec"
	WARPConnectorTunnelWARPConnectorTunnelTunTypeGRE           WARPConnectorTunnelWARPConnectorTunnelTunType = "gre"
	WARPConnectorTunnelWARPConnectorTunnelTunTypeCni           WARPConnectorTunnelWARPConnectorTunnelTunType = "cni"
)

func (r WARPConnectorTunnelWARPConnectorTunnelTunType) IsKnown() bool {
	switch r {
	case WARPConnectorTunnelWARPConnectorTunnelTunTypeCfdTunnel, WARPConnectorTunnelWARPConnectorTunnelTunTypeWARPConnector, WARPConnectorTunnelWARPConnectorTunnelTunTypeIPSec, WARPConnectorTunnelWARPConnectorTunnelTunTypeGRE, WARPConnectorTunnelWARPConnectorTunnelTunTypeCni:
		return true
	}
	return false
}

// The type of tunnel.
type WARPConnectorTunType string

const (
	WARPConnectorTunTypeCfdTunnel     WARPConnectorTunType = "cfd_tunnel"
	WARPConnectorTunTypeWARPConnector WARPConnectorTunType = "warp_connector"
	WARPConnectorTunTypeIPSec         WARPConnectorTunType = "ip_sec"
	WARPConnectorTunTypeGRE           WARPConnectorTunType = "gre"
	WARPConnectorTunTypeCni           WARPConnectorTunType = "cni"
)

func (r WARPConnectorTunType) IsKnown() bool {
	switch r {
	case WARPConnectorTunTypeCfdTunnel, WARPConnectorTunTypeWARPConnector, WARPConnectorTunTypeIPSec, WARPConnectorTunTypeGRE, WARPConnectorTunTypeCni:
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
	// A user-friendly name for the tunnel.
	Name param.Field[string] `json:"name,required"`
}

func (r WARPConnectorNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WARPConnectorNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result WARPConnector `json:"result,required"`
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
	// UUID of the tunnel.
	UUID          param.Field[string]    `query:"uuid"`
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

type WARPConnectorDeleteParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	Body      interface{}         `json:"body,required"`
}

func (r WARPConnectorDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type WARPConnectorDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result WARPConnector `json:"result,required"`
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
	// A user-friendly name for the tunnel.
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
	Result WARPConnector `json:"result,required"`
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
	Result WARPConnector `json:"result,required"`
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
