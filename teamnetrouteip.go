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

// TeamnetRouteIPService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewTeamnetRouteIPService] method
// instead.
type TeamnetRouteIPService struct {
	Options []option.RequestOption
}

// NewTeamnetRouteIPService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTeamnetRouteIPService(opts ...option.RequestOption) (r *TeamnetRouteIPService) {
	r = &TeamnetRouteIPService{}
	r.Options = opts
	return
}

// Fetches routes that contain the given IP address.
func (r *TeamnetRouteIPService) Get(ctx context.Context, accountID string, ip string, query TeamnetRouteIPGetParams, opts ...option.RequestOption) (res *TeamnetRouteIPGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TeamnetRouteIPGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/routes/ip/%s", accountID, ip)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TeamnetRouteIPGetResponse struct {
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
	TunType TeamnetRouteIPGetResponseTunType `json:"tun_type"`
	// UUID of the Cloudflare Tunnel serving the route.
	TunnelID interface{} `json:"tunnel_id"`
	// The user-friendly name of the Cloudflare Tunnel serving the route.
	TunnelName interface{} `json:"tunnel_name"`
	// UUID of the Tunnel Virtual Network this route belongs to. If no virtual networks
	// are configured, the route is assigned to the default virtual network of the
	// account.
	VirtualNetworkID interface{} `json:"virtual_network_id"`
	// A user-friendly name for the virtual network.
	VirtualNetworkName string                        `json:"virtual_network_name"`
	JSON               teamnetRouteIPGetResponseJSON `json:"-"`
}

// teamnetRouteIPGetResponseJSON contains the JSON metadata for the struct
// [TeamnetRouteIPGetResponse]
type teamnetRouteIPGetResponseJSON struct {
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

func (r *TeamnetRouteIPGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type TeamnetRouteIPGetResponseTunType string

const (
	TeamnetRouteIPGetResponseTunTypeCfdTunnel     TeamnetRouteIPGetResponseTunType = "cfd_tunnel"
	TeamnetRouteIPGetResponseTunTypeWARPConnector TeamnetRouteIPGetResponseTunType = "warp_connector"
	TeamnetRouteIPGetResponseTunTypeIPSec         TeamnetRouteIPGetResponseTunType = "ip_sec"
	TeamnetRouteIPGetResponseTunTypeGRE           TeamnetRouteIPGetResponseTunType = "gre"
	TeamnetRouteIPGetResponseTunTypeCni           TeamnetRouteIPGetResponseTunType = "cni"
)

type TeamnetRouteIPGetParams struct {
	// UUID of the Tunnel Virtual Network this route belongs to. If no virtual networks
	// are configured, the route is assigned to the default virtual network of the
	// account.
	VirtualNetworkID param.Field[interface{}] `query:"virtual_network_id"`
}

// URLQuery serializes [TeamnetRouteIPGetParams]'s query parameters as
// `url.Values`.
func (r TeamnetRouteIPGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TeamnetRouteIPGetResponseEnvelope struct {
	Errors   []TeamnetRouteIPGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TeamnetRouteIPGetResponseEnvelopeMessages `json:"messages,required"`
	Result   TeamnetRouteIPGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success TeamnetRouteIPGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    teamnetRouteIPGetResponseEnvelopeJSON    `json:"-"`
}

// teamnetRouteIPGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [TeamnetRouteIPGetResponseEnvelope]
type teamnetRouteIPGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteIPGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteIPGetResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    teamnetRouteIPGetResponseEnvelopeErrorsJSON `json:"-"`
}

// teamnetRouteIPGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [TeamnetRouteIPGetResponseEnvelopeErrors]
type teamnetRouteIPGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteIPGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteIPGetResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    teamnetRouteIPGetResponseEnvelopeMessagesJSON `json:"-"`
}

// teamnetRouteIPGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [TeamnetRouteIPGetResponseEnvelopeMessages]
type teamnetRouteIPGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteIPGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TeamnetRouteIPGetResponseEnvelopeSuccess bool

const (
	TeamnetRouteIPGetResponseEnvelopeSuccessTrue TeamnetRouteIPGetResponseEnvelopeSuccess = true
)
