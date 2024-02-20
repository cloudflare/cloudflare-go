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

// Lists and filters private network routes in an account.
func (r *TeamnetRouteService) List(ctx context.Context, accountID string, query TeamnetRouteListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[TeamnetRouteListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/teamnet/routes", accountID)
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

// Lists and filters private network routes in an account.
func (r *TeamnetRouteService) ListAutoPaging(ctx context.Context, accountID string, query TeamnetRouteListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[TeamnetRouteListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, accountID, query, opts...))
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
	TeamnetRouteListResponseTunTypeWarpConnector TeamnetRouteListResponseTunType = "warp_connector"
	TeamnetRouteListResponseTunTypeIPSec         TeamnetRouteListResponseTunType = "ip_sec"
	TeamnetRouteListResponseTunTypeGre           TeamnetRouteListResponseTunType = "gre"
	TeamnetRouteListResponseTunTypeCni           TeamnetRouteListResponseTunType = "cni"
)

type TeamnetRouteListParams struct {
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
