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

// TeamnetRouteService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewTeamnetRouteService] method
// instead.
type TeamnetRouteService struct {
	Options  []option.RequestOption
	IPs      *TeamnetRouteIPService
	Networks *TeamnetRouteNetworkService
}

// NewTeamnetRouteService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTeamnetRouteService(opts ...option.RequestOption) (r *TeamnetRouteService) {
	r = &TeamnetRouteService{}
	r.Options = opts
	r.IPs = NewTeamnetRouteIPService(opts...)
	r.Networks = NewTeamnetRouteNetworkService(opts...)
	return
}

// Routes a private network through a Cloudflare Tunnel.
func (r *TeamnetRouteService) New(ctx context.Context, params TeamnetRouteNewParams, opts ...option.RequestOption) (res *TeamnetRouteNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TeamnetRouteNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/routes", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists and filters private network routes in an account.
func (r *TeamnetRouteService) List(ctx context.Context, params TeamnetRouteListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[TeamnetRouteListResponse], err error) {
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
func (r *TeamnetRouteService) ListAutoPaging(ctx context.Context, params TeamnetRouteListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[TeamnetRouteListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes a private network route from an account.
func (r *TeamnetRouteService) Delete(ctx context.Context, routeID string, body TeamnetRouteDeleteParams, opts ...option.RequestOption) (res *TeamnetRouteDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TeamnetRouteDeleteResponseEnvelope
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
func (r *TeamnetRouteService) Edit(ctx context.Context, routeID string, params TeamnetRouteEditParams, opts ...option.RequestOption) (res *TeamnetRouteEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TeamnetRouteEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/routes/%s", params.AccountID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TeamnetRouteNewResponse struct {
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
	VirtualNetworkID interface{}                 `json:"virtual_network_id"`
	JSON             teamnetRouteNewResponseJSON `json:"-"`
}

// teamnetRouteNewResponseJSON contains the JSON metadata for the struct
// [TeamnetRouteNewResponse]
type teamnetRouteNewResponseJSON struct {
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

func (r *TeamnetRouteNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteListResponse struct {
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
	TunType TeamnetRouteListResponseTunType `json:"tun_type"`
	// UUID of the Cloudflare Tunnel serving the route.
	TunnelID interface{} `json:"tunnel_id"`
	// The user-friendly name of the Cloudflare Tunnel serving the route.
	TunnelName interface{} `json:"tunnel_name"`
	// UUID of the Tunnel Virtual Network this route belongs to. If no virtual networks
	// are configured, the route is assigned to the default virtual network of the
	// account.
	VirtualNetworkID interface{} `json:"virtual_network_id"`
	// A user-friendly name for the virtual network.
	VirtualNetworkName string                       `json:"virtual_network_name"`
	JSON               teamnetRouteListResponseJSON `json:"-"`
}

// teamnetRouteListResponseJSON contains the JSON metadata for the struct
// [TeamnetRouteListResponse]
type teamnetRouteListResponseJSON struct {
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

func (r *TeamnetRouteListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type TeamnetRouteListResponseTunType string

const (
	TeamnetRouteListResponseTunTypeCfdTunnel     TeamnetRouteListResponseTunType = "cfd_tunnel"
	TeamnetRouteListResponseTunTypeWARPConnector TeamnetRouteListResponseTunType = "warp_connector"
	TeamnetRouteListResponseTunTypeIPSec         TeamnetRouteListResponseTunType = "ip_sec"
	TeamnetRouteListResponseTunTypeGRE           TeamnetRouteListResponseTunType = "gre"
	TeamnetRouteListResponseTunTypeCni           TeamnetRouteListResponseTunType = "cni"
)

type TeamnetRouteDeleteResponse struct {
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
	VirtualNetworkID interface{}                    `json:"virtual_network_id"`
	JSON             teamnetRouteDeleteResponseJSON `json:"-"`
}

// teamnetRouteDeleteResponseJSON contains the JSON metadata for the struct
// [TeamnetRouteDeleteResponse]
type teamnetRouteDeleteResponseJSON struct {
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

func (r *TeamnetRouteDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteEditResponse struct {
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
	VirtualNetworkID interface{}                  `json:"virtual_network_id"`
	JSON             teamnetRouteEditResponseJSON `json:"-"`
}

// teamnetRouteEditResponseJSON contains the JSON metadata for the struct
// [TeamnetRouteEditResponse]
type teamnetRouteEditResponseJSON struct {
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

func (r *TeamnetRouteEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteNewParams struct {
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

func (r TeamnetRouteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TeamnetRouteNewResponseEnvelope struct {
	Errors   []TeamnetRouteNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TeamnetRouteNewResponseEnvelopeMessages `json:"messages,required"`
	Result   TeamnetRouteNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success TeamnetRouteNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    teamnetRouteNewResponseEnvelopeJSON    `json:"-"`
}

// teamnetRouteNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [TeamnetRouteNewResponseEnvelope]
type teamnetRouteNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteNewResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    teamnetRouteNewResponseEnvelopeErrorsJSON `json:"-"`
}

// teamnetRouteNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [TeamnetRouteNewResponseEnvelopeErrors]
type teamnetRouteNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteNewResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    teamnetRouteNewResponseEnvelopeMessagesJSON `json:"-"`
}

// teamnetRouteNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [TeamnetRouteNewResponseEnvelopeMessages]
type teamnetRouteNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TeamnetRouteNewResponseEnvelopeSuccess bool

const (
	TeamnetRouteNewResponseEnvelopeSuccessTrue TeamnetRouteNewResponseEnvelopeSuccess = true
)

type TeamnetRouteListParams struct {
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

// URLQuery serializes [TeamnetRouteListParams]'s query parameters as `url.Values`.
func (r TeamnetRouteListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TeamnetRouteDeleteParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type TeamnetRouteDeleteResponseEnvelope struct {
	Errors   []TeamnetRouteDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TeamnetRouteDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   TeamnetRouteDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success TeamnetRouteDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    teamnetRouteDeleteResponseEnvelopeJSON    `json:"-"`
}

// teamnetRouteDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [TeamnetRouteDeleteResponseEnvelope]
type teamnetRouteDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteDeleteResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    teamnetRouteDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// teamnetRouteDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [TeamnetRouteDeleteResponseEnvelopeErrors]
type teamnetRouteDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteDeleteResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    teamnetRouteDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// teamnetRouteDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [TeamnetRouteDeleteResponseEnvelopeMessages]
type teamnetRouteDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TeamnetRouteDeleteResponseEnvelopeSuccess bool

const (
	TeamnetRouteDeleteResponseEnvelopeSuccessTrue TeamnetRouteDeleteResponseEnvelopeSuccess = true
)

type TeamnetRouteEditParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// Optional remark describing the route.
	Comment param.Field[string] `json:"comment"`
	// The private IPv4 or IPv6 range connected by the route, in CIDR notation.
	Network param.Field[string] `json:"network"`
	// The type of tunnel.
	TunType param.Field[TeamnetRouteEditParamsTunType] `json:"tun_type"`
	// UUID of the Cloudflare Tunnel serving the route.
	TunnelID param.Field[interface{}] `json:"tunnel_id"`
	// UUID of the Tunnel Virtual Network this route belongs to. If no virtual networks
	// are configured, the route is assigned to the default virtual network of the
	// account.
	VirtualNetworkID param.Field[interface{}] `json:"virtual_network_id"`
}

func (r TeamnetRouteEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of tunnel.
type TeamnetRouteEditParamsTunType string

const (
	TeamnetRouteEditParamsTunTypeCfdTunnel     TeamnetRouteEditParamsTunType = "cfd_tunnel"
	TeamnetRouteEditParamsTunTypeWARPConnector TeamnetRouteEditParamsTunType = "warp_connector"
	TeamnetRouteEditParamsTunTypeIPSec         TeamnetRouteEditParamsTunType = "ip_sec"
	TeamnetRouteEditParamsTunTypeGRE           TeamnetRouteEditParamsTunType = "gre"
	TeamnetRouteEditParamsTunTypeCni           TeamnetRouteEditParamsTunType = "cni"
)

type TeamnetRouteEditResponseEnvelope struct {
	Errors   []TeamnetRouteEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TeamnetRouteEditResponseEnvelopeMessages `json:"messages,required"`
	Result   TeamnetRouteEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success TeamnetRouteEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    teamnetRouteEditResponseEnvelopeJSON    `json:"-"`
}

// teamnetRouteEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [TeamnetRouteEditResponseEnvelope]
type teamnetRouteEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteEditResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    teamnetRouteEditResponseEnvelopeErrorsJSON `json:"-"`
}

// teamnetRouteEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [TeamnetRouteEditResponseEnvelopeErrors]
type teamnetRouteEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteEditResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    teamnetRouteEditResponseEnvelopeMessagesJSON `json:"-"`
}

// teamnetRouteEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [TeamnetRouteEditResponseEnvelopeMessages]
type teamnetRouteEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TeamnetRouteEditResponseEnvelopeSuccess bool

const (
	TeamnetRouteEditResponseEnvelopeSuccessTrue TeamnetRouteEditResponseEnvelopeSuccess = true
)
