// File generated from our OpenAPI spec by Stainless.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// NetworkRouteIPService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewNetworkRouteIPService] method
// instead.
type NetworkRouteIPService struct {
	Options []option.RequestOption
}

// NewNetworkRouteIPService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNetworkRouteIPService(opts ...option.RequestOption) (r *NetworkRouteIPService) {
	r = &NetworkRouteIPService{}
	r.Options = opts
	return
}

// Fetches routes that contain the given IP address.
func (r *NetworkRouteIPService) Get(ctx context.Context, ip string, params NetworkRouteIPGetParams, opts ...option.RequestOption) (res *NetworkRouteIPGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env NetworkRouteIPGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/routes/ip/%s", params.AccountID, ip)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type NetworkRouteIPGetResponse struct {
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
	TunType NetworkRouteIPGetResponseTunType `json:"tun_type"`
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
	JSON               networkRouteIPGetResponseJSON `json:"-"`
}

// networkRouteIPGetResponseJSON contains the JSON metadata for the struct
// [NetworkRouteIPGetResponse]
type networkRouteIPGetResponseJSON struct {
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

func (r *NetworkRouteIPGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkRouteIPGetResponseJSON) RawJSON() string {
	return r.raw
}

// The type of tunnel.
type NetworkRouteIPGetResponseTunType string

const (
	NetworkRouteIPGetResponseTunTypeCfdTunnel     NetworkRouteIPGetResponseTunType = "cfd_tunnel"
	NetworkRouteIPGetResponseTunTypeWARPConnector NetworkRouteIPGetResponseTunType = "warp_connector"
	NetworkRouteIPGetResponseTunTypeIPSec         NetworkRouteIPGetResponseTunType = "ip_sec"
	NetworkRouteIPGetResponseTunTypeGRE           NetworkRouteIPGetResponseTunType = "gre"
	NetworkRouteIPGetResponseTunTypeCni           NetworkRouteIPGetResponseTunType = "cni"
)

type NetworkRouteIPGetParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// UUID of the Tunnel Virtual Network this route belongs to. If no virtual networks
	// are configured, the route is assigned to the default virtual network of the
	// account.
	VirtualNetworkID param.Field[interface{}] `query:"virtual_network_id"`
}

// URLQuery serializes [NetworkRouteIPGetParams]'s query parameters as
// `url.Values`.
func (r NetworkRouteIPGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NetworkRouteIPGetResponseEnvelope struct {
	Errors   []NetworkRouteIPGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []NetworkRouteIPGetResponseEnvelopeMessages `json:"messages,required"`
	Result   NetworkRouteIPGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success NetworkRouteIPGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    networkRouteIPGetResponseEnvelopeJSON    `json:"-"`
}

// networkRouteIPGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [NetworkRouteIPGetResponseEnvelope]
type networkRouteIPGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkRouteIPGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkRouteIPGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NetworkRouteIPGetResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    networkRouteIPGetResponseEnvelopeErrorsJSON `json:"-"`
}

// networkRouteIPGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [NetworkRouteIPGetResponseEnvelopeErrors]
type networkRouteIPGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkRouteIPGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkRouteIPGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type NetworkRouteIPGetResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    networkRouteIPGetResponseEnvelopeMessagesJSON `json:"-"`
}

// networkRouteIPGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [NetworkRouteIPGetResponseEnvelopeMessages]
type networkRouteIPGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkRouteIPGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkRouteIPGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NetworkRouteIPGetResponseEnvelopeSuccess bool

const (
	NetworkRouteIPGetResponseEnvelopeSuccessTrue NetworkRouteIPGetResponseEnvelopeSuccess = true
)
