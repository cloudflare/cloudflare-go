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
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountTeamnetRouteNetworkService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountTeamnetRouteNetworkService] method instead.
type AccountTeamnetRouteNetworkService struct {
	Options []option.RequestOption
}

// NewAccountTeamnetRouteNetworkService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountTeamnetRouteNetworkService(opts ...option.RequestOption) (r *AccountTeamnetRouteNetworkService) {
	r = &AccountTeamnetRouteNetworkService{}
	r.Options = opts
	return
}

// Routes a private network through a Cloudflare Tunnel. The CIDR in
// `ip_network_encoded` must be written in URL-encoded format.
func (r *AccountTeamnetRouteNetworkService) Update(ctx context.Context, accountIdentifier string, ipNetworkEncoded string, body AccountTeamnetRouteNetworkUpdateParams, opts ...option.RequestOption) (res *AccountTeamnetRouteNetworkUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/teamnet/routes/network/%s", accountIdentifier, ipNetworkEncoded)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Deletes a private network route from an account. The CIDR in
// `ip_network_encoded` must be written in URL-encoded format. If no
// virtual_network_id is provided it will delete the route from the default vnet.
// If no tun_type is provided it will fetch the type from the tunnel_id or if that
// is missing it will assume Cloudflare Tunnel as default. If tunnel_id is provided
// it will delete the route from that tunnel, otherwise it will delete the route
// based on the vnet and tun_type.
func (r *AccountTeamnetRouteNetworkService) Delete(ctx context.Context, accountIdentifier string, ipNetworkEncoded string, body AccountTeamnetRouteNetworkDeleteParams, opts ...option.RequestOption) (res *AccountTeamnetRouteNetworkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/teamnet/routes/network/%s", accountIdentifier, ipNetworkEncoded)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type AccountTeamnetRouteNetworkUpdateResponse struct {
	Errors   []AccountTeamnetRouteNetworkUpdateResponseError   `json:"errors"`
	Messages []AccountTeamnetRouteNetworkUpdateResponseMessage `json:"messages"`
	Result   AccountTeamnetRouteNetworkUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountTeamnetRouteNetworkUpdateResponseSuccess `json:"success"`
	JSON    accountTeamnetRouteNetworkUpdateResponseJSON    `json:"-"`
}

// accountTeamnetRouteNetworkUpdateResponseJSON contains the JSON metadata for the
// struct [AccountTeamnetRouteNetworkUpdateResponse]
type accountTeamnetRouteNetworkUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteNetworkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetRouteNetworkUpdateResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountTeamnetRouteNetworkUpdateResponseErrorJSON `json:"-"`
}

// accountTeamnetRouteNetworkUpdateResponseErrorJSON contains the JSON metadata for
// the struct [AccountTeamnetRouteNetworkUpdateResponseError]
type accountTeamnetRouteNetworkUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteNetworkUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetRouteNetworkUpdateResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountTeamnetRouteNetworkUpdateResponseMessageJSON `json:"-"`
}

// accountTeamnetRouteNetworkUpdateResponseMessageJSON contains the JSON metadata
// for the struct [AccountTeamnetRouteNetworkUpdateResponseMessage]
type accountTeamnetRouteNetworkUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteNetworkUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetRouteNetworkUpdateResponseResult struct {
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
	VirtualNetworkID interface{}                                        `json:"virtual_network_id"`
	JSON             accountTeamnetRouteNetworkUpdateResponseResultJSON `json:"-"`
}

// accountTeamnetRouteNetworkUpdateResponseResultJSON contains the JSON metadata
// for the struct [AccountTeamnetRouteNetworkUpdateResponseResult]
type accountTeamnetRouteNetworkUpdateResponseResultJSON struct {
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

func (r *AccountTeamnetRouteNetworkUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountTeamnetRouteNetworkUpdateResponseSuccess bool

const (
	AccountTeamnetRouteNetworkUpdateResponseSuccessTrue AccountTeamnetRouteNetworkUpdateResponseSuccess = true
)

type AccountTeamnetRouteNetworkDeleteResponse struct {
	Errors   []AccountTeamnetRouteNetworkDeleteResponseError   `json:"errors"`
	Messages []AccountTeamnetRouteNetworkDeleteResponseMessage `json:"messages"`
	Result   AccountTeamnetRouteNetworkDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountTeamnetRouteNetworkDeleteResponseSuccess `json:"success"`
	JSON    accountTeamnetRouteNetworkDeleteResponseJSON    `json:"-"`
}

// accountTeamnetRouteNetworkDeleteResponseJSON contains the JSON metadata for the
// struct [AccountTeamnetRouteNetworkDeleteResponse]
type accountTeamnetRouteNetworkDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteNetworkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetRouteNetworkDeleteResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountTeamnetRouteNetworkDeleteResponseErrorJSON `json:"-"`
}

// accountTeamnetRouteNetworkDeleteResponseErrorJSON contains the JSON metadata for
// the struct [AccountTeamnetRouteNetworkDeleteResponseError]
type accountTeamnetRouteNetworkDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteNetworkDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetRouteNetworkDeleteResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountTeamnetRouteNetworkDeleteResponseMessageJSON `json:"-"`
}

// accountTeamnetRouteNetworkDeleteResponseMessageJSON contains the JSON metadata
// for the struct [AccountTeamnetRouteNetworkDeleteResponseMessage]
type accountTeamnetRouteNetworkDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteNetworkDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetRouteNetworkDeleteResponseResult struct {
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
	VirtualNetworkID interface{}                                        `json:"virtual_network_id"`
	JSON             accountTeamnetRouteNetworkDeleteResponseResultJSON `json:"-"`
}

// accountTeamnetRouteNetworkDeleteResponseResultJSON contains the JSON metadata
// for the struct [AccountTeamnetRouteNetworkDeleteResponseResult]
type accountTeamnetRouteNetworkDeleteResponseResultJSON struct {
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

func (r *AccountTeamnetRouteNetworkDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountTeamnetRouteNetworkDeleteResponseSuccess bool

const (
	AccountTeamnetRouteNetworkDeleteResponseSuccessTrue AccountTeamnetRouteNetworkDeleteResponseSuccess = true
)

type AccountTeamnetRouteNetworkUpdateParams struct {
	// Optional remark describing the route.
	Comment param.Field[string] `json:"comment"`
	// UUID of the Tunnel Virtual Network this route belongs to. If no virtual networks
	// are configured, the route is assigned to the default virtual network of the
	// account.
	VirtualNetworkID param.Field[interface{}] `json:"virtual_network_id"`
}

func (r AccountTeamnetRouteNetworkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountTeamnetRouteNetworkDeleteParams struct {
	// The type of tunnel.
	TunType param.Field[AccountTeamnetRouteNetworkDeleteParamsTunType] `query:"tun_type"`
}

// URLQuery serializes [AccountTeamnetRouteNetworkDeleteParams]'s query parameters
// as `url.Values`.
func (r AccountTeamnetRouteNetworkDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The type of tunnel.
type AccountTeamnetRouteNetworkDeleteParamsTunType string

const (
	AccountTeamnetRouteNetworkDeleteParamsTunTypeCfdTunnel     AccountTeamnetRouteNetworkDeleteParamsTunType = "cfd_tunnel"
	AccountTeamnetRouteNetworkDeleteParamsTunTypeWarpConnector AccountTeamnetRouteNetworkDeleteParamsTunType = "warp_connector"
	AccountTeamnetRouteNetworkDeleteParamsTunTypeIPSec         AccountTeamnetRouteNetworkDeleteParamsTunType = "ip_sec"
	AccountTeamnetRouteNetworkDeleteParamsTunTypeGre           AccountTeamnetRouteNetworkDeleteParamsTunType = "gre"
	AccountTeamnetRouteNetworkDeleteParamsTunTypeCni           AccountTeamnetRouteNetworkDeleteParamsTunType = "cni"
)
