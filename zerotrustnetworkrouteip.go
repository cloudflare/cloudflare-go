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

// ZeroTrustNetworkRouteIPService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZeroTrustNetworkRouteIPService] method instead.
type ZeroTrustNetworkRouteIPService struct {
	Options []option.RequestOption
}

// NewZeroTrustNetworkRouteIPService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustNetworkRouteIPService(opts ...option.RequestOption) (r *ZeroTrustNetworkRouteIPService) {
	r = &ZeroTrustNetworkRouteIPService{}
	r.Options = opts
	return
}

// Fetches routes that contain the given IP address.
func (r *ZeroTrustNetworkRouteIPService) Get(ctx context.Context, ip string, params ZeroTrustNetworkRouteIPGetParams, opts ...option.RequestOption) (res *ZeroTrustNetworkRouteIPGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustNetworkRouteIPGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/routes/ip/%s", params.AccountID, ip)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustNetworkRouteIPGetResponse struct {
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
	TunType ZeroTrustNetworkRouteIPGetResponseTunType `json:"tun_type"`
	// UUID of the Cloudflare Tunnel serving the route.
	TunnelID interface{} `json:"tunnel_id"`
	// The user-friendly name of the Cloudflare Tunnel serving the route.
	TunnelName interface{} `json:"tunnel_name"`
	// UUID of the Tunnel Virtual Network this route belongs to. If no virtual networks
	// are configured, the route is assigned to the default virtual network of the
	// account.
	VirtualNetworkID interface{} `json:"virtual_network_id"`
	// A user-friendly name for the virtual network.
	VirtualNetworkName string                                 `json:"virtual_network_name"`
	JSON               zeroTrustNetworkRouteIPGetResponseJSON `json:"-"`
}

// zeroTrustNetworkRouteIPGetResponseJSON contains the JSON metadata for the struct
// [ZeroTrustNetworkRouteIPGetResponse]
type zeroTrustNetworkRouteIPGetResponseJSON struct {
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

func (r *ZeroTrustNetworkRouteIPGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type ZeroTrustNetworkRouteIPGetResponseTunType string

const (
	ZeroTrustNetworkRouteIPGetResponseTunTypeCfdTunnel     ZeroTrustNetworkRouteIPGetResponseTunType = "cfd_tunnel"
	ZeroTrustNetworkRouteIPGetResponseTunTypeWARPConnector ZeroTrustNetworkRouteIPGetResponseTunType = "warp_connector"
	ZeroTrustNetworkRouteIPGetResponseTunTypeIPSec         ZeroTrustNetworkRouteIPGetResponseTunType = "ip_sec"
	ZeroTrustNetworkRouteIPGetResponseTunTypeGRE           ZeroTrustNetworkRouteIPGetResponseTunType = "gre"
	ZeroTrustNetworkRouteIPGetResponseTunTypeCni           ZeroTrustNetworkRouteIPGetResponseTunType = "cni"
)

type ZeroTrustNetworkRouteIPGetParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// UUID of the Tunnel Virtual Network this route belongs to. If no virtual networks
	// are configured, the route is assigned to the default virtual network of the
	// account.
	VirtualNetworkID param.Field[interface{}] `query:"virtual_network_id"`
}

// URLQuery serializes [ZeroTrustNetworkRouteIPGetParams]'s query parameters as
// `url.Values`.
func (r ZeroTrustNetworkRouteIPGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZeroTrustNetworkRouteIPGetResponseEnvelope struct {
	Errors   []ZeroTrustNetworkRouteIPGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustNetworkRouteIPGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustNetworkRouteIPGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustNetworkRouteIPGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustNetworkRouteIPGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustNetworkRouteIPGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustNetworkRouteIPGetResponseEnvelope]
type zeroTrustNetworkRouteIPGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkRouteIPGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustNetworkRouteIPGetResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zeroTrustNetworkRouteIPGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustNetworkRouteIPGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustNetworkRouteIPGetResponseEnvelopeErrors]
type zeroTrustNetworkRouteIPGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkRouteIPGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustNetworkRouteIPGetResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustNetworkRouteIPGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustNetworkRouteIPGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustNetworkRouteIPGetResponseEnvelopeMessages]
type zeroTrustNetworkRouteIPGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkRouteIPGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustNetworkRouteIPGetResponseEnvelopeSuccess bool

const (
	ZeroTrustNetworkRouteIPGetResponseEnvelopeSuccessTrue ZeroTrustNetworkRouteIPGetResponseEnvelopeSuccess = true
)
