// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZeroTrustNetworkRouteService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustNetworkRouteService]
// method instead.
type ZeroTrustNetworkRouteService struct {
	Options  []option.RequestOption
	IPs      *ZeroTrustNetworkRouteIPService
	Networks *ZeroTrustNetworkRouteNetworkService
}

// NewZeroTrustNetworkRouteService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustNetworkRouteService(opts ...option.RequestOption) (r *ZeroTrustNetworkRouteService) {
	r = &ZeroTrustNetworkRouteService{}
	r.Options = opts
	r.IPs = NewZeroTrustNetworkRouteIPService(opts...)
	r.Networks = NewZeroTrustNetworkRouteNetworkService(opts...)
	return
}

// Routes a private network through a Cloudflare Tunnel.
func (r *ZeroTrustNetworkRouteService) New(ctx context.Context, params ZeroTrustNetworkRouteNewParams, opts ...option.RequestOption) (res *TunnelRoute, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustNetworkRouteNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/routes", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists and filters private network routes in an account.
func (r *ZeroTrustNetworkRouteService) List(ctx context.Context, params ZeroTrustNetworkRouteListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[TunnelTeamnet], err error) {
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
func (r *ZeroTrustNetworkRouteService) ListAutoPaging(ctx context.Context, params ZeroTrustNetworkRouteListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[TunnelTeamnet] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes a private network route from an account.
func (r *ZeroTrustNetworkRouteService) Delete(ctx context.Context, routeID string, body ZeroTrustNetworkRouteDeleteParams, opts ...option.RequestOption) (res *TunnelRoute, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustNetworkRouteDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/routes/%s", body.AccountID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing private network route in an account. The fields that are
// meant to be updated should be provided in the body of the request.
func (r *ZeroTrustNetworkRouteService) Edit(ctx context.Context, routeID string, params ZeroTrustNetworkRouteEditParams, opts ...option.RequestOption) (res *TunnelRoute, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustNetworkRouteEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/routes/%s", params.AccountID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TunnelRoute struct {
	// UUID of the route.
	ID string `json:"id"`
	// Optional remark describing the route.
	Comment string `json:"comment"`
	// Timestamp of when the route was created.
	CreatedAt interface{} `json:"created_at"`
	// Timestamp of when the route was deleted. If `null`, the route has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// The private IPv4 or IPv6 range connected by the route, in CIDR notation.
	Network string `json:"network"`
	// UUID of the Cloudflare Tunnel serving the route.
	TunnelID interface{} `json:"tunnel_id"`
	// UUID of the Tunnel Virtual Network this route belongs to. If no virtual networks
	// are configured, the route is assigned to the default virtual network of the
	// account.
	VirtualNetworkID interface{}     `json:"virtual_network_id"`
	JSON             tunnelRouteJSON `json:"-"`
}

// tunnelRouteJSON contains the JSON metadata for the struct [TunnelRoute]
type tunnelRouteJSON struct {
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

func (r *TunnelRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelTeamnet struct {
	// UUID of the route.
	ID string `json:"id"`
	// Optional remark describing the route.
	Comment string `json:"comment"`
	// Timestamp of when the route was created.
	CreatedAt interface{} `json:"created_at"`
	// Timestamp of when the route was deleted. If `null`, the route has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// The private IPv4 or IPv6 range connected by the route, in CIDR notation.
	Network string `json:"network"`
	// The type of tunnel.
	TunType TunnelTeamnetTunType `json:"tun_type"`
	// UUID of the Cloudflare Tunnel serving the route.
	TunnelID interface{} `json:"tunnel_id"`
	// The user-friendly name of the Cloudflare Tunnel serving the route.
	TunnelName interface{} `json:"tunnel_name"`
	// UUID of the Tunnel Virtual Network this route belongs to. If no virtual networks
	// are configured, the route is assigned to the default virtual network of the
	// account.
	VirtualNetworkID interface{} `json:"virtual_network_id"`
	// A user-friendly name for the virtual network.
	VirtualNetworkName string            `json:"virtual_network_name"`
	JSON               tunnelTeamnetJSON `json:"-"`
}

// tunnelTeamnetJSON contains the JSON metadata for the struct [TunnelTeamnet]
type tunnelTeamnetJSON struct {
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

func (r *TunnelTeamnet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type TunnelTeamnetTunType string

const (
	TunnelTeamnetTunTypeCfdTunnel     TunnelTeamnetTunType = "cfd_tunnel"
	TunnelTeamnetTunTypeWARPConnector TunnelTeamnetTunType = "warp_connector"
	TunnelTeamnetTunTypeIPSec         TunnelTeamnetTunType = "ip_sec"
	TunnelTeamnetTunTypeGRE           TunnelTeamnetTunType = "gre"
	TunnelTeamnetTunTypeCni           TunnelTeamnetTunType = "cni"
)

type ZeroTrustNetworkRouteNewParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// The private IPv4 or IPv6 range connected by the route, in CIDR notation.
	IPNetwork param.Field[string] `json:"ip_network,required"`
	// Optional remark describing the route.
	Comment param.Field[string] `json:"comment"`
	// UUID of the Tunnel Virtual Network this route belongs to. If no virtual networks
	// are configured, the route is assigned to the default virtual network of the
	// account.
	VirtualNetworkID param.Field[interface{}] `json:"virtual_network_id"`
}

func (r ZeroTrustNetworkRouteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustNetworkRouteNewResponseEnvelope struct {
	Errors   []ZeroTrustNetworkRouteNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustNetworkRouteNewResponseEnvelopeMessages `json:"messages,required"`
	Result   TunnelRoute                                        `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustNetworkRouteNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustNetworkRouteNewResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustNetworkRouteNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustNetworkRouteNewResponseEnvelope]
type zeroTrustNetworkRouteNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkRouteNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustNetworkRouteNewResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zeroTrustNetworkRouteNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustNetworkRouteNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustNetworkRouteNewResponseEnvelopeErrors]
type zeroTrustNetworkRouteNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkRouteNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustNetworkRouteNewResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zeroTrustNetworkRouteNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustNetworkRouteNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustNetworkRouteNewResponseEnvelopeMessages]
type zeroTrustNetworkRouteNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkRouteNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustNetworkRouteNewResponseEnvelopeSuccess bool

const (
	ZeroTrustNetworkRouteNewResponseEnvelopeSuccessTrue ZeroTrustNetworkRouteNewResponseEnvelopeSuccess = true
)

type ZeroTrustNetworkRouteListParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// Optional remark describing the route.
	Comment param.Field[string] `query:"comment"`
	// If provided, include only routes that were created (and not deleted) before this
	// time.
	ExistedAt param.Field[interface{}] `query:"existed_at"`
	// If `true`, only include deleted routes. If `false`, exclude deleted routes. If
	// empty, all routes will be included.
	IsDeleted param.Field[interface{}] `query:"is_deleted"`
	// If set, only list routes that are contained within this IP range.
	NetworkSubset param.Field[interface{}] `query:"network_subset"`
	// If set, only list routes that contain this IP range.
	NetworkSuperset param.Field[interface{}] `query:"network_superset"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of results to display.
	PerPage param.Field[float64] `query:"per_page"`
	// The types of tunnels to filter separated by a comma.
	TunTypes param.Field[string] `query:"tun_types"`
	// UUID of the Cloudflare Tunnel serving the route.
	TunnelID param.Field[interface{}] `query:"tunnel_id"`
	// UUID of the Tunnel Virtual Network this route belongs to. If no virtual networks
	// are configured, the route is assigned to the default virtual network of the
	// account.
	VirtualNetworkID param.Field[interface{}] `query:"virtual_network_id"`
}

// URLQuery serializes [ZeroTrustNetworkRouteListParams]'s query parameters as
// `url.Values`.
func (r ZeroTrustNetworkRouteListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZeroTrustNetworkRouteDeleteParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZeroTrustNetworkRouteDeleteResponseEnvelope struct {
	Errors   []ZeroTrustNetworkRouteDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustNetworkRouteDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   TunnelRoute                                           `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustNetworkRouteDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustNetworkRouteDeleteResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustNetworkRouteDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustNetworkRouteDeleteResponseEnvelope]
type zeroTrustNetworkRouteDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkRouteDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustNetworkRouteDeleteResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zeroTrustNetworkRouteDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustNetworkRouteDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustNetworkRouteDeleteResponseEnvelopeErrors]
type zeroTrustNetworkRouteDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkRouteDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustNetworkRouteDeleteResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zeroTrustNetworkRouteDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustNetworkRouteDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustNetworkRouteDeleteResponseEnvelopeMessages]
type zeroTrustNetworkRouteDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkRouteDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustNetworkRouteDeleteResponseEnvelopeSuccess bool

const (
	ZeroTrustNetworkRouteDeleteResponseEnvelopeSuccessTrue ZeroTrustNetworkRouteDeleteResponseEnvelopeSuccess = true
)

type ZeroTrustNetworkRouteEditParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// Optional remark describing the route.
	Comment param.Field[string] `json:"comment"`
	// The private IPv4 or IPv6 range connected by the route, in CIDR notation.
	Network param.Field[string] `json:"network"`
	// The type of tunnel.
	TunType param.Field[ZeroTrustNetworkRouteEditParamsTunType] `json:"tun_type"`
	// UUID of the Cloudflare Tunnel serving the route.
	TunnelID param.Field[interface{}] `json:"tunnel_id"`
	// UUID of the Tunnel Virtual Network this route belongs to. If no virtual networks
	// are configured, the route is assigned to the default virtual network of the
	// account.
	VirtualNetworkID param.Field[interface{}] `json:"virtual_network_id"`
}

func (r ZeroTrustNetworkRouteEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of tunnel.
type ZeroTrustNetworkRouteEditParamsTunType string

const (
	ZeroTrustNetworkRouteEditParamsTunTypeCfdTunnel     ZeroTrustNetworkRouteEditParamsTunType = "cfd_tunnel"
	ZeroTrustNetworkRouteEditParamsTunTypeWARPConnector ZeroTrustNetworkRouteEditParamsTunType = "warp_connector"
	ZeroTrustNetworkRouteEditParamsTunTypeIPSec         ZeroTrustNetworkRouteEditParamsTunType = "ip_sec"
	ZeroTrustNetworkRouteEditParamsTunTypeGRE           ZeroTrustNetworkRouteEditParamsTunType = "gre"
	ZeroTrustNetworkRouteEditParamsTunTypeCni           ZeroTrustNetworkRouteEditParamsTunType = "cni"
)

type ZeroTrustNetworkRouteEditResponseEnvelope struct {
	Errors   []ZeroTrustNetworkRouteEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustNetworkRouteEditResponseEnvelopeMessages `json:"messages,required"`
	Result   TunnelRoute                                         `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustNetworkRouteEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustNetworkRouteEditResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustNetworkRouteEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustNetworkRouteEditResponseEnvelope]
type zeroTrustNetworkRouteEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkRouteEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustNetworkRouteEditResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zeroTrustNetworkRouteEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustNetworkRouteEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustNetworkRouteEditResponseEnvelopeErrors]
type zeroTrustNetworkRouteEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkRouteEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustNetworkRouteEditResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zeroTrustNetworkRouteEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustNetworkRouteEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustNetworkRouteEditResponseEnvelopeMessages]
type zeroTrustNetworkRouteEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkRouteEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustNetworkRouteEditResponseEnvelopeSuccess bool

const (
	ZeroTrustNetworkRouteEditResponseEnvelopeSuccessTrue ZeroTrustNetworkRouteEditResponseEnvelopeSuccess = true
)
