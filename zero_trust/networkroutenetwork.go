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

// NetworkRouteNetworkService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewNetworkRouteNetworkService]
// method instead.
type NetworkRouteNetworkService struct {
	Options []option.RequestOption
}

// NewNetworkRouteNetworkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNetworkRouteNetworkService(opts ...option.RequestOption) (r *NetworkRouteNetworkService) {
	r = &NetworkRouteNetworkService{}
	r.Options = opts
	return
}

// Routes a private network through a Cloudflare Tunnel. The CIDR in
// `ip_network_encoded` must be written in URL-encoded format.
func (r *NetworkRouteNetworkService) New(ctx context.Context, ipNetworkEncoded string, params NetworkRouteNetworkNewParams, opts ...option.RequestOption) (res *NetworkRouteNetworkNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env NetworkRouteNetworkNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/routes/network/%s", params.AccountID, ipNetworkEncoded)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a private network route from an account. The CIDR in
// `ip_network_encoded` must be written in URL-encoded format. If no
// virtual_network_id is provided it will delete the route from the default vnet.
// If no tun_type is provided it will fetch the type from the tunnel_id or if that
// is missing it will assume Cloudflare Tunnel as default. If tunnel_id is provided
// it will delete the route from that tunnel, otherwise it will delete the route
// based on the vnet and tun_type.
func (r *NetworkRouteNetworkService) Delete(ctx context.Context, ipNetworkEncoded string, params NetworkRouteNetworkDeleteParams, opts ...option.RequestOption) (res *NetworkRouteNetworkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env NetworkRouteNetworkDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/routes/network/%s", params.AccountID, ipNetworkEncoded)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing private network route in an account. The CIDR in
// `ip_network_encoded` must be written in URL-encoded format.
func (r *NetworkRouteNetworkService) Edit(ctx context.Context, ipNetworkEncoded string, body NetworkRouteNetworkEditParams, opts ...option.RequestOption) (res *NetworkRouteNetworkEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env NetworkRouteNetworkEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/routes/network/%s", body.AccountID, ipNetworkEncoded)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type NetworkRouteNetworkNewResponse struct {
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
	VirtualNetworkID interface{}                        `json:"virtual_network_id"`
	JSON             networkRouteNetworkNewResponseJSON `json:"-"`
}

// networkRouteNetworkNewResponseJSON contains the JSON metadata for the struct
// [NetworkRouteNetworkNewResponse]
type networkRouteNetworkNewResponseJSON struct {
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

func (r *NetworkRouteNetworkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkRouteNetworkNewResponseJSON) RawJSON() string {
	return r.raw
}

type NetworkRouteNetworkDeleteResponse struct {
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
	VirtualNetworkID interface{}                           `json:"virtual_network_id"`
	JSON             networkRouteNetworkDeleteResponseJSON `json:"-"`
}

// networkRouteNetworkDeleteResponseJSON contains the JSON metadata for the struct
// [NetworkRouteNetworkDeleteResponse]
type networkRouteNetworkDeleteResponseJSON struct {
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

func (r *NetworkRouteNetworkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkRouteNetworkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type NetworkRouteNetworkEditResponse struct {
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
	VirtualNetworkID interface{}                         `json:"virtual_network_id"`
	JSON             networkRouteNetworkEditResponseJSON `json:"-"`
}

// networkRouteNetworkEditResponseJSON contains the JSON metadata for the struct
// [NetworkRouteNetworkEditResponse]
type networkRouteNetworkEditResponseJSON struct {
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

func (r *NetworkRouteNetworkEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkRouteNetworkEditResponseJSON) RawJSON() string {
	return r.raw
}

type NetworkRouteNetworkNewParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// Optional remark describing the route.
	Comment param.Field[string] `json:"comment"`
	// UUID of the Tunnel Virtual Network this route belongs to. If no virtual networks
	// are configured, the route is assigned to the default virtual network of the
	// account.
	VirtualNetworkID param.Field[interface{}] `json:"virtual_network_id"`
}

func (r NetworkRouteNetworkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NetworkRouteNetworkNewResponseEnvelope struct {
	Errors   []NetworkRouteNetworkNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []NetworkRouteNetworkNewResponseEnvelopeMessages `json:"messages,required"`
	Result   NetworkRouteNetworkNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success NetworkRouteNetworkNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    networkRouteNetworkNewResponseEnvelopeJSON    `json:"-"`
}

// networkRouteNetworkNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [NetworkRouteNetworkNewResponseEnvelope]
type networkRouteNetworkNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkRouteNetworkNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkRouteNetworkNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NetworkRouteNetworkNewResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    networkRouteNetworkNewResponseEnvelopeErrorsJSON `json:"-"`
}

// networkRouteNetworkNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [NetworkRouteNetworkNewResponseEnvelopeErrors]
type networkRouteNetworkNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkRouteNetworkNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkRouteNetworkNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type NetworkRouteNetworkNewResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    networkRouteNetworkNewResponseEnvelopeMessagesJSON `json:"-"`
}

// networkRouteNetworkNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [NetworkRouteNetworkNewResponseEnvelopeMessages]
type networkRouteNetworkNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkRouteNetworkNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkRouteNetworkNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NetworkRouteNetworkNewResponseEnvelopeSuccess bool

const (
	NetworkRouteNetworkNewResponseEnvelopeSuccessTrue NetworkRouteNetworkNewResponseEnvelopeSuccess = true
)

type NetworkRouteNetworkDeleteParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// The type of tunnel.
	TunType param.Field[NetworkRouteNetworkDeleteParamsTunType] `query:"tun_type"`
}

// URLQuery serializes [NetworkRouteNetworkDeleteParams]'s query parameters as
// `url.Values`.
func (r NetworkRouteNetworkDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The type of tunnel.
type NetworkRouteNetworkDeleteParamsTunType string

const (
	NetworkRouteNetworkDeleteParamsTunTypeCfdTunnel     NetworkRouteNetworkDeleteParamsTunType = "cfd_tunnel"
	NetworkRouteNetworkDeleteParamsTunTypeWARPConnector NetworkRouteNetworkDeleteParamsTunType = "warp_connector"
	NetworkRouteNetworkDeleteParamsTunTypeIPSec         NetworkRouteNetworkDeleteParamsTunType = "ip_sec"
	NetworkRouteNetworkDeleteParamsTunTypeGRE           NetworkRouteNetworkDeleteParamsTunType = "gre"
	NetworkRouteNetworkDeleteParamsTunTypeCni           NetworkRouteNetworkDeleteParamsTunType = "cni"
)

type NetworkRouteNetworkDeleteResponseEnvelope struct {
	Errors   []NetworkRouteNetworkDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []NetworkRouteNetworkDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   NetworkRouteNetworkDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success NetworkRouteNetworkDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    networkRouteNetworkDeleteResponseEnvelopeJSON    `json:"-"`
}

// networkRouteNetworkDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [NetworkRouteNetworkDeleteResponseEnvelope]
type networkRouteNetworkDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkRouteNetworkDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkRouteNetworkDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NetworkRouteNetworkDeleteResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    networkRouteNetworkDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// networkRouteNetworkDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [NetworkRouteNetworkDeleteResponseEnvelopeErrors]
type networkRouteNetworkDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkRouteNetworkDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkRouteNetworkDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type NetworkRouteNetworkDeleteResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    networkRouteNetworkDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// networkRouteNetworkDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [NetworkRouteNetworkDeleteResponseEnvelopeMessages]
type networkRouteNetworkDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkRouteNetworkDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkRouteNetworkDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NetworkRouteNetworkDeleteResponseEnvelopeSuccess bool

const (
	NetworkRouteNetworkDeleteResponseEnvelopeSuccessTrue NetworkRouteNetworkDeleteResponseEnvelopeSuccess = true
)

type NetworkRouteNetworkEditParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type NetworkRouteNetworkEditResponseEnvelope struct {
	Errors   []NetworkRouteNetworkEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []NetworkRouteNetworkEditResponseEnvelopeMessages `json:"messages,required"`
	Result   NetworkRouteNetworkEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success NetworkRouteNetworkEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    networkRouteNetworkEditResponseEnvelopeJSON    `json:"-"`
}

// networkRouteNetworkEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [NetworkRouteNetworkEditResponseEnvelope]
type networkRouteNetworkEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkRouteNetworkEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkRouteNetworkEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NetworkRouteNetworkEditResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    networkRouteNetworkEditResponseEnvelopeErrorsJSON `json:"-"`
}

// networkRouteNetworkEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [NetworkRouteNetworkEditResponseEnvelopeErrors]
type networkRouteNetworkEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkRouteNetworkEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkRouteNetworkEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type NetworkRouteNetworkEditResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    networkRouteNetworkEditResponseEnvelopeMessagesJSON `json:"-"`
}

// networkRouteNetworkEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [NetworkRouteNetworkEditResponseEnvelopeMessages]
type networkRouteNetworkEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkRouteNetworkEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkRouteNetworkEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NetworkRouteNetworkEditResponseEnvelopeSuccess bool

const (
	NetworkRouteNetworkEditResponseEnvelopeSuccessTrue NetworkRouteNetworkEditResponseEnvelopeSuccess = true
)
