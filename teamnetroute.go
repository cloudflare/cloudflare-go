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

// Lists and filters private network routes in an account.
func (r *TeamnetRouteService) TunnelRouteListTunnelRoutes(ctx context.Context, accountID string, query TeamnetRouteTunnelRouteListTunnelRoutesParams, opts ...option.RequestOption) (res *[]TeamnetRouteTunnelRouteListTunnelRoutesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TeamnetRouteTunnelRouteListTunnelRoutesResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/routes", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TeamnetRouteTunnelRouteListTunnelRoutesResponse struct {
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
	TunType TeamnetRouteTunnelRouteListTunnelRoutesResponseTunType `json:"tun_type"`
	// UUID of the Cloudflare Tunnel serving the route.
	TunnelID interface{} `json:"tunnel_id"`
	// The user-friendly name of the Cloudflare Tunnel serving the route.
	TunnelName interface{} `json:"tunnel_name"`
	// UUID of the Tunnel Virtual Network this route belongs to. If no virtual networks
	// are configured, the route is assigned to the default virtual network of the
	// account.
	VirtualNetworkID interface{} `json:"virtual_network_id"`
	// A user-friendly name for the virtual network.
	VirtualNetworkName string                                              `json:"virtual_network_name"`
	JSON               teamnetRouteTunnelRouteListTunnelRoutesResponseJSON `json:"-"`
}

// teamnetRouteTunnelRouteListTunnelRoutesResponseJSON contains the JSON metadata
// for the struct [TeamnetRouteTunnelRouteListTunnelRoutesResponse]
type teamnetRouteTunnelRouteListTunnelRoutesResponseJSON struct {
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

func (r *TeamnetRouteTunnelRouteListTunnelRoutesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type TeamnetRouteTunnelRouteListTunnelRoutesResponseTunType string

const (
	TeamnetRouteTunnelRouteListTunnelRoutesResponseTunTypeCfdTunnel     TeamnetRouteTunnelRouteListTunnelRoutesResponseTunType = "cfd_tunnel"
	TeamnetRouteTunnelRouteListTunnelRoutesResponseTunTypeWarpConnector TeamnetRouteTunnelRouteListTunnelRoutesResponseTunType = "warp_connector"
	TeamnetRouteTunnelRouteListTunnelRoutesResponseTunTypeIPSec         TeamnetRouteTunnelRouteListTunnelRoutesResponseTunType = "ip_sec"
	TeamnetRouteTunnelRouteListTunnelRoutesResponseTunTypeGre           TeamnetRouteTunnelRouteListTunnelRoutesResponseTunType = "gre"
	TeamnetRouteTunnelRouteListTunnelRoutesResponseTunTypeCni           TeamnetRouteTunnelRouteListTunnelRoutesResponseTunType = "cni"
)

type TeamnetRouteTunnelRouteListTunnelRoutesParams struct {
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

// URLQuery serializes [TeamnetRouteTunnelRouteListTunnelRoutesParams]'s query
// parameters as `url.Values`.
func (r TeamnetRouteTunnelRouteListTunnelRoutesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TeamnetRouteTunnelRouteListTunnelRoutesResponseEnvelope struct {
	Errors   []TeamnetRouteTunnelRouteListTunnelRoutesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TeamnetRouteTunnelRouteListTunnelRoutesResponseEnvelopeMessages `json:"messages,required"`
	Result   []TeamnetRouteTunnelRouteListTunnelRoutesResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    TeamnetRouteTunnelRouteListTunnelRoutesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo TeamnetRouteTunnelRouteListTunnelRoutesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       teamnetRouteTunnelRouteListTunnelRoutesResponseEnvelopeJSON       `json:"-"`
}

// teamnetRouteTunnelRouteListTunnelRoutesResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [TeamnetRouteTunnelRouteListTunnelRoutesResponseEnvelope]
type teamnetRouteTunnelRouteListTunnelRoutesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteTunnelRouteListTunnelRoutesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteTunnelRouteListTunnelRoutesResponseEnvelopeErrors struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    teamnetRouteTunnelRouteListTunnelRoutesResponseEnvelopeErrorsJSON `json:"-"`
}

// teamnetRouteTunnelRouteListTunnelRoutesResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [TeamnetRouteTunnelRouteListTunnelRoutesResponseEnvelopeErrors]
type teamnetRouteTunnelRouteListTunnelRoutesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteTunnelRouteListTunnelRoutesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteTunnelRouteListTunnelRoutesResponseEnvelopeMessages struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    teamnetRouteTunnelRouteListTunnelRoutesResponseEnvelopeMessagesJSON `json:"-"`
}

// teamnetRouteTunnelRouteListTunnelRoutesResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [TeamnetRouteTunnelRouteListTunnelRoutesResponseEnvelopeMessages]
type teamnetRouteTunnelRouteListTunnelRoutesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteTunnelRouteListTunnelRoutesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TeamnetRouteTunnelRouteListTunnelRoutesResponseEnvelopeSuccess bool

const (
	TeamnetRouteTunnelRouteListTunnelRoutesResponseEnvelopeSuccessTrue TeamnetRouteTunnelRouteListTunnelRoutesResponseEnvelopeSuccess = true
)

type TeamnetRouteTunnelRouteListTunnelRoutesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                               `json:"total_count"`
	JSON       teamnetRouteTunnelRouteListTunnelRoutesResponseEnvelopeResultInfoJSON `json:"-"`
}

// teamnetRouteTunnelRouteListTunnelRoutesResponseEnvelopeResultInfoJSON contains
// the JSON metadata for the struct
// [TeamnetRouteTunnelRouteListTunnelRoutesResponseEnvelopeResultInfo]
type teamnetRouteTunnelRouteListTunnelRoutesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteTunnelRouteListTunnelRoutesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
