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

// AccountTeamnetRouteService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountTeamnetRouteService]
// method instead.
type AccountTeamnetRouteService struct {
	Options  []option.RequestOption
	IPs      *AccountTeamnetRouteIPService
	Networks *AccountTeamnetRouteNetworkService
}

// NewAccountTeamnetRouteService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountTeamnetRouteService(opts ...option.RequestOption) (r *AccountTeamnetRouteService) {
	r = &AccountTeamnetRouteService{}
	r.Options = opts
	r.IPs = NewAccountTeamnetRouteIPService(opts...)
	r.Networks = NewAccountTeamnetRouteNetworkService(opts...)
	return
}

// Routes a private network through a Cloudflare Tunnel.
func (r *AccountTeamnetRouteService) New(ctx context.Context, accountIdentifier string, body AccountTeamnetRouteNewParams, opts ...option.RequestOption) (res *AccountTeamnetRouteNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/teamnet/routes", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates an existing private network route in an account. The fields that are
// meant to be updated should be provided in the body of the request.
func (r *AccountTeamnetRouteService) Update(ctx context.Context, accountIdentifier string, routeID string, body AccountTeamnetRouteUpdateParams, opts ...option.RequestOption) (res *AccountTeamnetRouteUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/teamnet/routes/%s", accountIdentifier, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Deletes a private network route from an account.
func (r *AccountTeamnetRouteService) Delete(ctx context.Context, accountIdentifier string, routeID string, opts ...option.RequestOption) (res *AccountTeamnetRouteDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/teamnet/routes/%s", accountIdentifier, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Lists and filters private network routes in an account.
func (r *AccountTeamnetRouteService) TunnelRouteListTunnelRoutes(ctx context.Context, accountIdentifier string, query AccountTeamnetRouteTunnelRouteListTunnelRoutesParams, opts ...option.RequestOption) (res *shared.Page[AccountTeamnetRouteTunnelRouteListTunnelRoutesResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/teamnet/routes", accountIdentifier)
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

type AccountTeamnetRouteNewResponse struct {
	Errors   []AccountTeamnetRouteNewResponseError   `json:"errors"`
	Messages []AccountTeamnetRouteNewResponseMessage `json:"messages"`
	Result   AccountTeamnetRouteNewResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountTeamnetRouteNewResponseSuccess `json:"success"`
	JSON    accountTeamnetRouteNewResponseJSON    `json:"-"`
}

// accountTeamnetRouteNewResponseJSON contains the JSON metadata for the struct
// [AccountTeamnetRouteNewResponse]
type accountTeamnetRouteNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetRouteNewResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accountTeamnetRouteNewResponseErrorJSON `json:"-"`
}

// accountTeamnetRouteNewResponseErrorJSON contains the JSON metadata for the
// struct [AccountTeamnetRouteNewResponseError]
type accountTeamnetRouteNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetRouteNewResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountTeamnetRouteNewResponseMessageJSON `json:"-"`
}

// accountTeamnetRouteNewResponseMessageJSON contains the JSON metadata for the
// struct [AccountTeamnetRouteNewResponseMessage]
type accountTeamnetRouteNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetRouteNewResponseResult struct {
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
	VirtualNetworkID interface{}                              `json:"virtual_network_id"`
	JSON             accountTeamnetRouteNewResponseResultJSON `json:"-"`
}

// accountTeamnetRouteNewResponseResultJSON contains the JSON metadata for the
// struct [AccountTeamnetRouteNewResponseResult]
type accountTeamnetRouteNewResponseResultJSON struct {
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

func (r *AccountTeamnetRouteNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountTeamnetRouteNewResponseSuccess bool

const (
	AccountTeamnetRouteNewResponseSuccessTrue AccountTeamnetRouteNewResponseSuccess = true
)

type AccountTeamnetRouteUpdateResponse struct {
	Errors   []AccountTeamnetRouteUpdateResponseError   `json:"errors"`
	Messages []AccountTeamnetRouteUpdateResponseMessage `json:"messages"`
	Result   AccountTeamnetRouteUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountTeamnetRouteUpdateResponseSuccess `json:"success"`
	JSON    accountTeamnetRouteUpdateResponseJSON    `json:"-"`
}

// accountTeamnetRouteUpdateResponseJSON contains the JSON metadata for the struct
// [AccountTeamnetRouteUpdateResponse]
type accountTeamnetRouteUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetRouteUpdateResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountTeamnetRouteUpdateResponseErrorJSON `json:"-"`
}

// accountTeamnetRouteUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccountTeamnetRouteUpdateResponseError]
type accountTeamnetRouteUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetRouteUpdateResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountTeamnetRouteUpdateResponseMessageJSON `json:"-"`
}

// accountTeamnetRouteUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccountTeamnetRouteUpdateResponseMessage]
type accountTeamnetRouteUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetRouteUpdateResponseResult struct {
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
	VirtualNetworkID interface{}                                 `json:"virtual_network_id"`
	JSON             accountTeamnetRouteUpdateResponseResultJSON `json:"-"`
}

// accountTeamnetRouteUpdateResponseResultJSON contains the JSON metadata for the
// struct [AccountTeamnetRouteUpdateResponseResult]
type accountTeamnetRouteUpdateResponseResultJSON struct {
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

func (r *AccountTeamnetRouteUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountTeamnetRouteUpdateResponseSuccess bool

const (
	AccountTeamnetRouteUpdateResponseSuccessTrue AccountTeamnetRouteUpdateResponseSuccess = true
)

type AccountTeamnetRouteDeleteResponse struct {
	Errors   []AccountTeamnetRouteDeleteResponseError   `json:"errors"`
	Messages []AccountTeamnetRouteDeleteResponseMessage `json:"messages"`
	Result   AccountTeamnetRouteDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountTeamnetRouteDeleteResponseSuccess `json:"success"`
	JSON    accountTeamnetRouteDeleteResponseJSON    `json:"-"`
}

// accountTeamnetRouteDeleteResponseJSON contains the JSON metadata for the struct
// [AccountTeamnetRouteDeleteResponse]
type accountTeamnetRouteDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetRouteDeleteResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountTeamnetRouteDeleteResponseErrorJSON `json:"-"`
}

// accountTeamnetRouteDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountTeamnetRouteDeleteResponseError]
type accountTeamnetRouteDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetRouteDeleteResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountTeamnetRouteDeleteResponseMessageJSON `json:"-"`
}

// accountTeamnetRouteDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountTeamnetRouteDeleteResponseMessage]
type accountTeamnetRouteDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetRouteDeleteResponseResult struct {
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
	VirtualNetworkID interface{}                                 `json:"virtual_network_id"`
	JSON             accountTeamnetRouteDeleteResponseResultJSON `json:"-"`
}

// accountTeamnetRouteDeleteResponseResultJSON contains the JSON metadata for the
// struct [AccountTeamnetRouteDeleteResponseResult]
type accountTeamnetRouteDeleteResponseResultJSON struct {
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

func (r *AccountTeamnetRouteDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountTeamnetRouteDeleteResponseSuccess bool

const (
	AccountTeamnetRouteDeleteResponseSuccessTrue AccountTeamnetRouteDeleteResponseSuccess = true
)

type AccountTeamnetRouteTunnelRouteListTunnelRoutesResponse struct {
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
	TunType AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseTunType `json:"tun_type"`
	// UUID of the Cloudflare Tunnel serving the route.
	TunnelID interface{} `json:"tunnel_id"`
	// The user-friendly name of the Cloudflare Tunnel serving the route.
	TunnelName interface{} `json:"tunnel_name"`
	// UUID of the Tunnel Virtual Network this route belongs to. If no virtual networks
	// are configured, the route is assigned to the default virtual network of the
	// account.
	VirtualNetworkID interface{} `json:"virtual_network_id"`
	// A user-friendly name for the virtual network.
	VirtualNetworkName string                                                     `json:"virtual_network_name"`
	JSON               accountTeamnetRouteTunnelRouteListTunnelRoutesResponseJSON `json:"-"`
}

// accountTeamnetRouteTunnelRouteListTunnelRoutesResponseJSON contains the JSON
// metadata for the struct [AccountTeamnetRouteTunnelRouteListTunnelRoutesResponse]
type accountTeamnetRouteTunnelRouteListTunnelRoutesResponseJSON struct {
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

func (r *AccountTeamnetRouteTunnelRouteListTunnelRoutesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseTunType string

const (
	AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseTunTypeCfdTunnel     AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseTunType = "cfd_tunnel"
	AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseTunTypeWarpConnector AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseTunType = "warp_connector"
	AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseTunTypeIPSec         AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseTunType = "ip_sec"
	AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseTunTypeGre           AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseTunType = "gre"
	AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseTunTypeCni           AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseTunType = "cni"
)

type AccountTeamnetRouteNewParams struct {
	// The private IPv4 or IPv6 range connected by the route, in CIDR notation.
	IPNetwork param.Field[string] `json:"ip_network,required"`
	// Optional remark describing the route.
	Comment param.Field[string] `json:"comment"`
	// UUID of the Tunnel Virtual Network this route belongs to. If no virtual networks
	// are configured, the route is assigned to the default virtual network of the
	// account.
	VirtualNetworkID param.Field[interface{}] `json:"virtual_network_id"`
}

func (r AccountTeamnetRouteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountTeamnetRouteUpdateParams struct {
	// Optional remark describing the route.
	Comment param.Field[string] `json:"comment"`
	// The private IPv4 or IPv6 range connected by the route, in CIDR notation.
	Network param.Field[string] `json:"network"`
	// The type of tunnel.
	TunType param.Field[AccountTeamnetRouteUpdateParamsTunType] `json:"tun_type"`
	// UUID of the Cloudflare Tunnel serving the route.
	TunnelID param.Field[interface{}] `json:"tunnel_id"`
	// UUID of the Tunnel Virtual Network this route belongs to. If no virtual networks
	// are configured, the route is assigned to the default virtual network of the
	// account.
	VirtualNetworkID param.Field[interface{}] `json:"virtual_network_id"`
}

func (r AccountTeamnetRouteUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of tunnel.
type AccountTeamnetRouteUpdateParamsTunType string

const (
	AccountTeamnetRouteUpdateParamsTunTypeCfdTunnel     AccountTeamnetRouteUpdateParamsTunType = "cfd_tunnel"
	AccountTeamnetRouteUpdateParamsTunTypeWarpConnector AccountTeamnetRouteUpdateParamsTunType = "warp_connector"
	AccountTeamnetRouteUpdateParamsTunTypeIPSec         AccountTeamnetRouteUpdateParamsTunType = "ip_sec"
	AccountTeamnetRouteUpdateParamsTunTypeGre           AccountTeamnetRouteUpdateParamsTunType = "gre"
	AccountTeamnetRouteUpdateParamsTunTypeCni           AccountTeamnetRouteUpdateParamsTunType = "cni"
)

type AccountTeamnetRouteTunnelRouteListTunnelRoutesParams struct {
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

// URLQuery serializes [AccountTeamnetRouteTunnelRouteListTunnelRoutesParams]'s
// query parameters as `url.Values`.
func (r AccountTeamnetRouteTunnelRouteListTunnelRoutesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
