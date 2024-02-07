// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// TeamnetRouteService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewTeamnetRouteService] method
// instead.
type TeamnetRouteService struct {
	Options []option.RequestOption
}

// NewTeamnetRouteService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTeamnetRouteService(opts ...option.RequestOption) (r *TeamnetRouteService) {
	r = &TeamnetRouteService{}
	r.Options = opts
	return
}

// Routes a private network through a Cloudflare Tunnel.
func (r *TeamnetRouteService) New(ctx context.Context, accountID string, body TeamnetRouteNewParams, opts ...option.RequestOption) (res *TeamnetRouteNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/teamnet/routes", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates an existing private network route in an account. The fields that are
// meant to be updated should be provided in the body of the request.
func (r *TeamnetRouteService) Update(ctx context.Context, accountID string, routeID string, body TeamnetRouteUpdateParams, opts ...option.RequestOption) (res *TeamnetRouteUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/teamnet/routes/%s", accountID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Deletes a private network route from an account.
func (r *TeamnetRouteService) Delete(ctx context.Context, accountID string, routeID string, opts ...option.RequestOption) (res *TeamnetRouteDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/teamnet/routes/%s", accountID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type TeamnetRouteNewResponse struct {
	Errors   []TeamnetRouteNewResponseError   `json:"errors"`
	Messages []TeamnetRouteNewResponseMessage `json:"messages"`
	Result   TeamnetRouteNewResponseResult    `json:"result"`
	// Whether the API call was successful
	Success TeamnetRouteNewResponseSuccess `json:"success"`
	JSON    teamnetRouteNewResponseJSON    `json:"-"`
}

// teamnetRouteNewResponseJSON contains the JSON metadata for the struct
// [TeamnetRouteNewResponse]
type teamnetRouteNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteNewResponseError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    teamnetRouteNewResponseErrorJSON `json:"-"`
}

// teamnetRouteNewResponseErrorJSON contains the JSON metadata for the struct
// [TeamnetRouteNewResponseError]
type teamnetRouteNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteNewResponseMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    teamnetRouteNewResponseMessageJSON `json:"-"`
}

// teamnetRouteNewResponseMessageJSON contains the JSON metadata for the struct
// [TeamnetRouteNewResponseMessage]
type teamnetRouteNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteNewResponseResult struct {
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
	VirtualNetworkID interface{}                       `json:"virtual_network_id"`
	JSON             teamnetRouteNewResponseResultJSON `json:"-"`
}

// teamnetRouteNewResponseResultJSON contains the JSON metadata for the struct
// [TeamnetRouteNewResponseResult]
type teamnetRouteNewResponseResultJSON struct {
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

func (r *TeamnetRouteNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TeamnetRouteNewResponseSuccess bool

const (
	TeamnetRouteNewResponseSuccessTrue TeamnetRouteNewResponseSuccess = true
)

type TeamnetRouteUpdateResponse struct {
	Errors   []TeamnetRouteUpdateResponseError   `json:"errors"`
	Messages []TeamnetRouteUpdateResponseMessage `json:"messages"`
	Result   TeamnetRouteUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success TeamnetRouteUpdateResponseSuccess `json:"success"`
	JSON    teamnetRouteUpdateResponseJSON    `json:"-"`
}

// teamnetRouteUpdateResponseJSON contains the JSON metadata for the struct
// [TeamnetRouteUpdateResponse]
type teamnetRouteUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteUpdateResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    teamnetRouteUpdateResponseErrorJSON `json:"-"`
}

// teamnetRouteUpdateResponseErrorJSON contains the JSON metadata for the struct
// [TeamnetRouteUpdateResponseError]
type teamnetRouteUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteUpdateResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    teamnetRouteUpdateResponseMessageJSON `json:"-"`
}

// teamnetRouteUpdateResponseMessageJSON contains the JSON metadata for the struct
// [TeamnetRouteUpdateResponseMessage]
type teamnetRouteUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteUpdateResponseResult struct {
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
	VirtualNetworkID interface{}                          `json:"virtual_network_id"`
	JSON             teamnetRouteUpdateResponseResultJSON `json:"-"`
}

// teamnetRouteUpdateResponseResultJSON contains the JSON metadata for the struct
// [TeamnetRouteUpdateResponseResult]
type teamnetRouteUpdateResponseResultJSON struct {
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

func (r *TeamnetRouteUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TeamnetRouteUpdateResponseSuccess bool

const (
	TeamnetRouteUpdateResponseSuccessTrue TeamnetRouteUpdateResponseSuccess = true
)

type TeamnetRouteDeleteResponse struct {
	Errors   []TeamnetRouteDeleteResponseError   `json:"errors"`
	Messages []TeamnetRouteDeleteResponseMessage `json:"messages"`
	Result   TeamnetRouteDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success TeamnetRouteDeleteResponseSuccess `json:"success"`
	JSON    teamnetRouteDeleteResponseJSON    `json:"-"`
}

// teamnetRouteDeleteResponseJSON contains the JSON metadata for the struct
// [TeamnetRouteDeleteResponse]
type teamnetRouteDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteDeleteResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    teamnetRouteDeleteResponseErrorJSON `json:"-"`
}

// teamnetRouteDeleteResponseErrorJSON contains the JSON metadata for the struct
// [TeamnetRouteDeleteResponseError]
type teamnetRouteDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteDeleteResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    teamnetRouteDeleteResponseMessageJSON `json:"-"`
}

// teamnetRouteDeleteResponseMessageJSON contains the JSON metadata for the struct
// [TeamnetRouteDeleteResponseMessage]
type teamnetRouteDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteDeleteResponseResult struct {
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
	VirtualNetworkID interface{}                          `json:"virtual_network_id"`
	JSON             teamnetRouteDeleteResponseResultJSON `json:"-"`
}

// teamnetRouteDeleteResponseResultJSON contains the JSON metadata for the struct
// [TeamnetRouteDeleteResponseResult]
type teamnetRouteDeleteResponseResultJSON struct {
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

func (r *TeamnetRouteDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TeamnetRouteDeleteResponseSuccess bool

const (
	TeamnetRouteDeleteResponseSuccessTrue TeamnetRouteDeleteResponseSuccess = true
)

type TeamnetRouteNewParams struct {
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

type TeamnetRouteUpdateParams struct {
	// Optional remark describing the route.
	Comment param.Field[string] `json:"comment"`
	// The private IPv4 or IPv6 range connected by the route, in CIDR notation.
	Network param.Field[string] `json:"network"`
	// The type of tunnel.
	TunType param.Field[TeamnetRouteUpdateParamsTunType] `json:"tun_type"`
	// UUID of the Cloudflare Tunnel serving the route.
	TunnelID param.Field[interface{}] `json:"tunnel_id"`
	// UUID of the Tunnel Virtual Network this route belongs to. If no virtual networks
	// are configured, the route is assigned to the default virtual network of the
	// account.
	VirtualNetworkID param.Field[interface{}] `json:"virtual_network_id"`
}

func (r TeamnetRouteUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of tunnel.
type TeamnetRouteUpdateParamsTunType string

const (
	TeamnetRouteUpdateParamsTunTypeCfdTunnel     TeamnetRouteUpdateParamsTunType = "cfd_tunnel"
	TeamnetRouteUpdateParamsTunTypeWarpConnector TeamnetRouteUpdateParamsTunType = "warp_connector"
	TeamnetRouteUpdateParamsTunTypeIPSec         TeamnetRouteUpdateParamsTunType = "ip_sec"
	TeamnetRouteUpdateParamsTunTypeGre           TeamnetRouteUpdateParamsTunType = "gre"
	TeamnetRouteUpdateParamsTunTypeCni           TeamnetRouteUpdateParamsTunType = "cni"
)
