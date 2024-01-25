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

// AccountTeamnetRouteIPService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountTeamnetRouteIPService]
// method instead.
type AccountTeamnetRouteIPService struct {
	Options []option.RequestOption
}

// NewAccountTeamnetRouteIPService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountTeamnetRouteIPService(opts ...option.RequestOption) (r *AccountTeamnetRouteIPService) {
	r = &AccountTeamnetRouteIPService{}
	r.Options = opts
	return
}

// Fetches routes that contain the given IP address.
func (r *AccountTeamnetRouteIPService) Get(ctx context.Context, accountIdentifier string, ip string, query AccountTeamnetRouteIPGetParams, opts ...option.RequestOption) (res *AccountTeamnetRouteIPGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/teamnet/routes/ip/%s", accountIdentifier, ip)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountTeamnetRouteIPGetResponse struct {
	Errors   []AccountTeamnetRouteIPGetResponseError   `json:"errors"`
	Messages []AccountTeamnetRouteIPGetResponseMessage `json:"messages"`
	Result   AccountTeamnetRouteIPGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountTeamnetRouteIPGetResponseSuccess `json:"success"`
	JSON    accountTeamnetRouteIPGetResponseJSON    `json:"-"`
}

// accountTeamnetRouteIPGetResponseJSON contains the JSON metadata for the struct
// [AccountTeamnetRouteIPGetResponse]
type accountTeamnetRouteIPGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteIPGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetRouteIPGetResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountTeamnetRouteIPGetResponseErrorJSON `json:"-"`
}

// accountTeamnetRouteIPGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountTeamnetRouteIPGetResponseError]
type accountTeamnetRouteIPGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteIPGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetRouteIPGetResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountTeamnetRouteIPGetResponseMessageJSON `json:"-"`
}

// accountTeamnetRouteIPGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountTeamnetRouteIPGetResponseMessage]
type accountTeamnetRouteIPGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteIPGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetRouteIPGetResponseResult struct {
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
	TunType AccountTeamnetRouteIPGetResponseResultTunType `json:"tun_type"`
	// UUID of the Cloudflare Tunnel serving the route.
	TunnelID interface{} `json:"tunnel_id"`
	// The user-friendly name of the Cloudflare Tunnel serving the route.
	TunnelName interface{} `json:"tunnel_name"`
	// UUID of the Tunnel Virtual Network this route belongs to. If no virtual networks
	// are configured, the route is assigned to the default virtual network of the
	// account.
	VirtualNetworkID interface{} `json:"virtual_network_id"`
	// A user-friendly name for the virtual network.
	VirtualNetworkName string                                     `json:"virtual_network_name"`
	JSON               accountTeamnetRouteIPGetResponseResultJSON `json:"-"`
}

// accountTeamnetRouteIPGetResponseResultJSON contains the JSON metadata for the
// struct [AccountTeamnetRouteIPGetResponseResult]
type accountTeamnetRouteIPGetResponseResultJSON struct {
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

func (r *AccountTeamnetRouteIPGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type AccountTeamnetRouteIPGetResponseResultTunType string

const (
	AccountTeamnetRouteIPGetResponseResultTunTypeCfdTunnel     AccountTeamnetRouteIPGetResponseResultTunType = "cfd_tunnel"
	AccountTeamnetRouteIPGetResponseResultTunTypeWarpConnector AccountTeamnetRouteIPGetResponseResultTunType = "warp_connector"
	AccountTeamnetRouteIPGetResponseResultTunTypeIPSec         AccountTeamnetRouteIPGetResponseResultTunType = "ip_sec"
	AccountTeamnetRouteIPGetResponseResultTunTypeGre           AccountTeamnetRouteIPGetResponseResultTunType = "gre"
	AccountTeamnetRouteIPGetResponseResultTunTypeCni           AccountTeamnetRouteIPGetResponseResultTunType = "cni"
)

// Whether the API call was successful
type AccountTeamnetRouteIPGetResponseSuccess bool

const (
	AccountTeamnetRouteIPGetResponseSuccessTrue AccountTeamnetRouteIPGetResponseSuccess = true
)

type AccountTeamnetRouteIPGetParams struct {
	// UUID of the Tunnel Virtual Network this route belongs to. If no virtual networks
	// are configured, the route is assigned to the default virtual network of the
	// account.
	VirtualNetworkID param.Field[interface{}] `query:"virtual_network_id"`
}

// URLQuery serializes [AccountTeamnetRouteIPGetParams]'s query parameters as
// `url.Values`.
func (r AccountTeamnetRouteIPGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
