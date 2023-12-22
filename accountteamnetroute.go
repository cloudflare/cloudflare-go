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

// Lists and filters private network routes in an account.
func (r *AccountTeamnetRouteService) TunnelRouteListTunnelRoutes(ctx context.Context, accountIdentifier string, query AccountTeamnetRouteTunnelRouteListTunnelRoutesParams, opts ...option.RequestOption) (res *AccountTeamnetRouteTunnelRouteListTunnelRoutesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/teamnet/routes", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountTeamnetRouteTunnelRouteListTunnelRoutesResponse struct {
	Errors     []AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseError    `json:"errors"`
	Messages   []AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseMessage  `json:"messages"`
	Result     []AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseResult   `json:"result"`
	ResultInfo AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseSuccess `json:"success"`
	JSON    accountTeamnetRouteTunnelRouteListTunnelRoutesResponseJSON    `json:"-"`
}

// accountTeamnetRouteTunnelRouteListTunnelRoutesResponseJSON contains the JSON
// metadata for the struct [AccountTeamnetRouteTunnelRouteListTunnelRoutesResponse]
type accountTeamnetRouteTunnelRouteListTunnelRoutesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteTunnelRouteListTunnelRoutesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseError struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    accountTeamnetRouteTunnelRouteListTunnelRoutesResponseErrorJSON `json:"-"`
}

// accountTeamnetRouteTunnelRouteListTunnelRoutesResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseError]
type accountTeamnetRouteTunnelRouteListTunnelRoutesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseMessage struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    accountTeamnetRouteTunnelRouteListTunnelRoutesResponseMessageJSON `json:"-"`
}

// accountTeamnetRouteTunnelRouteListTunnelRoutesResponseMessageJSON contains the
// JSON metadata for the struct
// [AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseMessage]
type accountTeamnetRouteTunnelRouteListTunnelRoutesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseResult struct {
	// Optional remark describing the route.
	Comment string `json:"comment,required"`
	// Timestamp of when the route was created.
	CreatedAt interface{} `json:"created_at,required"`
	// The private IPv4 or IPv6 range connected by the route, in CIDR notation.
	Network string `json:"network,required"`
	// UUID of the Cloudflare Tunnel serving the route.
	TunnelID interface{} `json:"tunnel_id,required"`
	// Timestamp of when the route was deleted. If `null`, the route has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// The user-friendly name of the Cloudflare Tunnel serving the route.
	TunnelName interface{} `json:"tunnel_name"`
	// UUID of the Tunnel Virtual Network this route belongs to. If no virtual networks
	// are configured, the route is assigned to the default virtual network of the
	// account.
	VirtualNetworkID interface{}                                                      `json:"virtual_network_id"`
	JSON             accountTeamnetRouteTunnelRouteListTunnelRoutesResponseResultJSON `json:"-"`
}

// accountTeamnetRouteTunnelRouteListTunnelRoutesResponseResultJSON contains the
// JSON metadata for the struct
// [AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseResult]
type accountTeamnetRouteTunnelRouteListTunnelRoutesResponseResultJSON struct {
	Comment          apijson.Field
	CreatedAt        apijson.Field
	Network          apijson.Field
	TunnelID         apijson.Field
	DeletedAt        apijson.Field
	TunnelName       apijson.Field
	VirtualNetworkID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                              `json:"total_count"`
	JSON       accountTeamnetRouteTunnelRouteListTunnelRoutesResponseResultInfoJSON `json:"-"`
}

// accountTeamnetRouteTunnelRouteListTunnelRoutesResponseResultInfoJSON contains
// the JSON metadata for the struct
// [AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseResultInfo]
type accountTeamnetRouteTunnelRouteListTunnelRoutesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseSuccess bool

const (
	AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseSuccessTrue AccountTeamnetRouteTunnelRouteListTunnelRoutesResponseSuccess = true
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
