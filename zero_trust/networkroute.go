// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// NetworkRouteService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewNetworkRouteService] method
// instead.
type NetworkRouteService struct {
	Options  []option.RequestOption
	IPs      *NetworkRouteIPService
	Networks *NetworkRouteNetworkService
}

// NewNetworkRouteService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNetworkRouteService(opts ...option.RequestOption) (r *NetworkRouteService) {
	r = &NetworkRouteService{}
	r.Options = opts
	r.IPs = NewNetworkRouteIPService(opts...)
	r.Networks = NewNetworkRouteNetworkService(opts...)
	return
}

// Routes a private network through a Cloudflare Tunnel.
func (r *NetworkRouteService) New(ctx context.Context, params NetworkRouteNewParams, opts ...option.RequestOption) (res *Route, err error) {
	opts = append(r.Options[:], opts...)
	var env NetworkRouteNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/routes", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists and filters private network routes in an account.
func (r *NetworkRouteService) List(ctx context.Context, params NetworkRouteListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[Teamnet], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/teamnet/routes", params.AccountID)
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

// Lists and filters private network routes in an account.
func (r *NetworkRouteService) ListAutoPaging(ctx context.Context, params NetworkRouteListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[Teamnet] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes a private network route from an account.
func (r *NetworkRouteService) Delete(ctx context.Context, routeID string, body NetworkRouteDeleteParams, opts ...option.RequestOption) (res *Route, err error) {
	opts = append(r.Options[:], opts...)
	var env NetworkRouteDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/routes/%s", body.AccountID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing private network route in an account. The fields that are
// meant to be updated should be provided in the body of the request.
func (r *NetworkRouteService) Edit(ctx context.Context, routeID string, params NetworkRouteEditParams, opts ...option.RequestOption) (res *Route, err error) {
	opts = append(r.Options[:], opts...)
	var env NetworkRouteEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/routes/%s", params.AccountID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Route struct {
	// UUID of the route.
	ID string `json:"id"`
	// Optional remark describing the route.
	Comment string `json:"comment"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// The private IPv4 or IPv6 range connected by the route, in CIDR notation.
	Network string `json:"network"`
	// UUID of the tunnel.
	TunnelID string `json:"tunnel_id" format:"uuid"`
	// UUID of the virtual network.
	VirtualNetworkID string    `json:"virtual_network_id" format:"uuid"`
	JSON             routeJSON `json:"-"`
}

// routeJSON contains the JSON metadata for the struct [Route]
type routeJSON struct {
	ID               apijson.Field
	Comment          apijson.Field
	CreatedAt        apijson.Field
	DeletedAt        apijson.Field
	Network          apijson.Field
	TunnelID         apijson.Field
	VirtualNetworkID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Route) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeJSON) RawJSON() string {
	return r.raw
}

type Teamnet struct {
	// UUID of the route.
	ID string `json:"id"`
	// Optional remark describing the route.
	Comment string `json:"comment"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// The private IPv4 or IPv6 range connected by the route, in CIDR notation.
	Network string `json:"network"`
	// The type of tunnel.
	TunType TeamnetTunType `json:"tun_type"`
	// UUID of the tunnel.
	TunnelID string `json:"tunnel_id" format:"uuid"`
	// A user-friendly name for a tunnel.
	TunnelName string `json:"tunnel_name"`
	// UUID of the virtual network.
	VirtualNetworkID string `json:"virtual_network_id" format:"uuid"`
	// A user-friendly name for the virtual network.
	VirtualNetworkName string      `json:"virtual_network_name"`
	JSON               teamnetJSON `json:"-"`
}

// teamnetJSON contains the JSON metadata for the struct [Teamnet]
type teamnetJSON struct {
	ID                 apijson.Field
	Comment            apijson.Field
	CreatedAt          apijson.Field
	DeletedAt          apijson.Field
	Network            apijson.Field
	TunType            apijson.Field
	TunnelID           apijson.Field
	TunnelName         apijson.Field
	VirtualNetworkID   apijson.Field
	VirtualNetworkName apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Teamnet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamnetJSON) RawJSON() string {
	return r.raw
}

// The type of tunnel.
type TeamnetTunType string

const (
	TeamnetTunTypeCfdTunnel     TeamnetTunType = "cfd_tunnel"
	TeamnetTunTypeWARPConnector TeamnetTunType = "warp_connector"
	TeamnetTunTypeIPSec         TeamnetTunType = "ip_sec"
	TeamnetTunTypeGRE           TeamnetTunType = "gre"
	TeamnetTunTypeCNI           TeamnetTunType = "cni"
)

func (r TeamnetTunType) IsKnown() bool {
	switch r {
	case TeamnetTunTypeCfdTunnel, TeamnetTunTypeWARPConnector, TeamnetTunTypeIPSec, TeamnetTunTypeGRE, TeamnetTunTypeCNI:
		return true
	}
	return false
}

type NetworkRouteNewParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// The private IPv4 or IPv6 range connected by the route, in CIDR notation.
	Network param.Field[string] `json:"network,required"`
	// Optional remark describing the route.
	Comment param.Field[string] `json:"comment"`
	// UUID of the virtual network.
	VirtualNetworkID param.Field[string] `json:"virtual_network_id" format:"uuid"`
}

func (r NetworkRouteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NetworkRouteNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Route                 `json:"result,required"`
	// Whether the API call was successful
	Success NetworkRouteNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    networkRouteNewResponseEnvelopeJSON    `json:"-"`
}

// networkRouteNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [NetworkRouteNewResponseEnvelope]
type networkRouteNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkRouteNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkRouteNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NetworkRouteNewResponseEnvelopeSuccess bool

const (
	NetworkRouteNewResponseEnvelopeSuccessTrue NetworkRouteNewResponseEnvelopeSuccess = true
)

func (r NetworkRouteNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NetworkRouteNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NetworkRouteListParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// Optional remark describing the route.
	Comment param.Field[string] `query:"comment"`
	// If provided, include only tunnels that were created (and not deleted) before
	// this time.
	ExistedAt param.Field[time.Time] `query:"existed_at" format:"date-time"`
	// If `true`, only include deleted routes. If `false`, exclude deleted routes. If
	// empty, all routes will be included.
	IsDeleted param.Field[bool] `query:"is_deleted"`
	// If set, only list routes that are contained within this IP range.
	NetworkSubset param.Field[string] `query:"network_subset"`
	// If set, only list routes that contain this IP range.
	NetworkSuperset param.Field[string] `query:"network_superset"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of results to display.
	PerPage param.Field[float64] `query:"per_page"`
	// UUID of the route.
	RouteID param.Field[string] `query:"route_id"`
	// The types of tunnels to filter separated by a comma.
	TunTypes param.Field[string] `query:"tun_types"`
	// UUID of the tunnel.
	TunnelID param.Field[string] `query:"tunnel_id" format:"uuid"`
	// UUID of the virtual network.
	VirtualNetworkID param.Field[string] `query:"virtual_network_id" format:"uuid"`
}

// URLQuery serializes [NetworkRouteListParams]'s query parameters as `url.Values`.
func (r NetworkRouteListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NetworkRouteDeleteParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type NetworkRouteDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Route                 `json:"result,required"`
	// Whether the API call was successful
	Success NetworkRouteDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    networkRouteDeleteResponseEnvelopeJSON    `json:"-"`
}

// networkRouteDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [NetworkRouteDeleteResponseEnvelope]
type networkRouteDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkRouteDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkRouteDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NetworkRouteDeleteResponseEnvelopeSuccess bool

const (
	NetworkRouteDeleteResponseEnvelopeSuccessTrue NetworkRouteDeleteResponseEnvelopeSuccess = true
)

func (r NetworkRouteDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NetworkRouteDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NetworkRouteEditParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// Optional remark describing the route.
	Comment param.Field[string] `json:"comment"`
	// The private IPv4 or IPv6 range connected by the route, in CIDR notation.
	Network param.Field[string] `json:"network"`
	// UUID of the virtual network.
	VirtualNetworkID param.Field[string] `json:"virtual_network_id" format:"uuid"`
}

func (r NetworkRouteEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NetworkRouteEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Route                 `json:"result,required"`
	// Whether the API call was successful
	Success NetworkRouteEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    networkRouteEditResponseEnvelopeJSON    `json:"-"`
}

// networkRouteEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [NetworkRouteEditResponseEnvelope]
type networkRouteEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkRouteEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkRouteEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NetworkRouteEditResponseEnvelopeSuccess bool

const (
	NetworkRouteEditResponseEnvelopeSuccessTrue NetworkRouteEditResponseEnvelopeSuccess = true
)

func (r NetworkRouteEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NetworkRouteEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
