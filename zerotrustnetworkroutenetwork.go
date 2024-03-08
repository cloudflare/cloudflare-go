// File generated from our OpenAPI spec by Stainless.

package cloudflare

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

// ZeroTrustNetworkRouteNetworkService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZeroTrustNetworkRouteNetworkService] method instead.
type ZeroTrustNetworkRouteNetworkService struct {
	Options []option.RequestOption
}

// NewZeroTrustNetworkRouteNetworkService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustNetworkRouteNetworkService(opts ...option.RequestOption) (r *ZeroTrustNetworkRouteNetworkService) {
	r = &ZeroTrustNetworkRouteNetworkService{}
	r.Options = opts
	return
}

// Routes a private network through a Cloudflare Tunnel. The CIDR in
// `ip_network_encoded` must be written in URL-encoded format.
func (r *ZeroTrustNetworkRouteNetworkService) New(ctx context.Context, ipNetworkEncoded string, params ZeroTrustNetworkRouteNetworkNewParams, opts ...option.RequestOption) (res *ZeroTrustNetworkRouteNetworkNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustNetworkRouteNetworkNewResponseEnvelope
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
func (r *ZeroTrustNetworkRouteNetworkService) Delete(ctx context.Context, ipNetworkEncoded string, params ZeroTrustNetworkRouteNetworkDeleteParams, opts ...option.RequestOption) (res *ZeroTrustNetworkRouteNetworkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustNetworkRouteNetworkDeleteResponseEnvelope
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
func (r *ZeroTrustNetworkRouteNetworkService) Edit(ctx context.Context, ipNetworkEncoded string, body ZeroTrustNetworkRouteNetworkEditParams, opts ...option.RequestOption) (res *ZeroTrustNetworkRouteNetworkEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustNetworkRouteNetworkEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/routes/network/%s", body.AccountID, ipNetworkEncoded)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustNetworkRouteNetworkNewResponse struct {
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
	VirtualNetworkID interface{}                                 `json:"virtual_network_id"`
	JSON             zeroTrustNetworkRouteNetworkNewResponseJSON `json:"-"`
}

// zeroTrustNetworkRouteNetworkNewResponseJSON contains the JSON metadata for the
// struct [ZeroTrustNetworkRouteNetworkNewResponse]
type zeroTrustNetworkRouteNetworkNewResponseJSON struct {
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

func (r *ZeroTrustNetworkRouteNetworkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustNetworkRouteNetworkNewResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustNetworkRouteNetworkDeleteResponse struct {
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
	VirtualNetworkID interface{}                                    `json:"virtual_network_id"`
	JSON             zeroTrustNetworkRouteNetworkDeleteResponseJSON `json:"-"`
}

// zeroTrustNetworkRouteNetworkDeleteResponseJSON contains the JSON metadata for
// the struct [ZeroTrustNetworkRouteNetworkDeleteResponse]
type zeroTrustNetworkRouteNetworkDeleteResponseJSON struct {
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

func (r *ZeroTrustNetworkRouteNetworkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustNetworkRouteNetworkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustNetworkRouteNetworkEditResponse struct {
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
	VirtualNetworkID interface{}                                  `json:"virtual_network_id"`
	JSON             zeroTrustNetworkRouteNetworkEditResponseJSON `json:"-"`
}

// zeroTrustNetworkRouteNetworkEditResponseJSON contains the JSON metadata for the
// struct [ZeroTrustNetworkRouteNetworkEditResponse]
type zeroTrustNetworkRouteNetworkEditResponseJSON struct {
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

func (r *ZeroTrustNetworkRouteNetworkEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustNetworkRouteNetworkEditResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustNetworkRouteNetworkNewParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// Optional remark describing the route.
	Comment param.Field[string] `json:"comment"`
	// UUID of the Tunnel Virtual Network this route belongs to. If no virtual networks
	// are configured, the route is assigned to the default virtual network of the
	// account.
	VirtualNetworkID param.Field[interface{}] `json:"virtual_network_id"`
}

func (r ZeroTrustNetworkRouteNetworkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustNetworkRouteNetworkNewResponseEnvelope struct {
	Errors   []ZeroTrustNetworkRouteNetworkNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustNetworkRouteNetworkNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustNetworkRouteNetworkNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustNetworkRouteNetworkNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustNetworkRouteNetworkNewResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustNetworkRouteNetworkNewResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustNetworkRouteNetworkNewResponseEnvelope]
type zeroTrustNetworkRouteNetworkNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkRouteNetworkNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustNetworkRouteNetworkNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustNetworkRouteNetworkNewResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zeroTrustNetworkRouteNetworkNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustNetworkRouteNetworkNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustNetworkRouteNetworkNewResponseEnvelopeErrors]
type zeroTrustNetworkRouteNetworkNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkRouteNetworkNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustNetworkRouteNetworkNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustNetworkRouteNetworkNewResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zeroTrustNetworkRouteNetworkNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustNetworkRouteNetworkNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustNetworkRouteNetworkNewResponseEnvelopeMessages]
type zeroTrustNetworkRouteNetworkNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkRouteNetworkNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustNetworkRouteNetworkNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustNetworkRouteNetworkNewResponseEnvelopeSuccess bool

const (
	ZeroTrustNetworkRouteNetworkNewResponseEnvelopeSuccessTrue ZeroTrustNetworkRouteNetworkNewResponseEnvelopeSuccess = true
)

type ZeroTrustNetworkRouteNetworkDeleteParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// The type of tunnel.
	TunType param.Field[ZeroTrustNetworkRouteNetworkDeleteParamsTunType] `query:"tun_type"`
}

// URLQuery serializes [ZeroTrustNetworkRouteNetworkDeleteParams]'s query
// parameters as `url.Values`.
func (r ZeroTrustNetworkRouteNetworkDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The type of tunnel.
type ZeroTrustNetworkRouteNetworkDeleteParamsTunType string

const (
	ZeroTrustNetworkRouteNetworkDeleteParamsTunTypeCfdTunnel     ZeroTrustNetworkRouteNetworkDeleteParamsTunType = "cfd_tunnel"
	ZeroTrustNetworkRouteNetworkDeleteParamsTunTypeWARPConnector ZeroTrustNetworkRouteNetworkDeleteParamsTunType = "warp_connector"
	ZeroTrustNetworkRouteNetworkDeleteParamsTunTypeIPSec         ZeroTrustNetworkRouteNetworkDeleteParamsTunType = "ip_sec"
	ZeroTrustNetworkRouteNetworkDeleteParamsTunTypeGRE           ZeroTrustNetworkRouteNetworkDeleteParamsTunType = "gre"
	ZeroTrustNetworkRouteNetworkDeleteParamsTunTypeCni           ZeroTrustNetworkRouteNetworkDeleteParamsTunType = "cni"
)

type ZeroTrustNetworkRouteNetworkDeleteResponseEnvelope struct {
	Errors   []ZeroTrustNetworkRouteNetworkDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustNetworkRouteNetworkDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustNetworkRouteNetworkDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustNetworkRouteNetworkDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustNetworkRouteNetworkDeleteResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustNetworkRouteNetworkDeleteResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustNetworkRouteNetworkDeleteResponseEnvelope]
type zeroTrustNetworkRouteNetworkDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkRouteNetworkDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustNetworkRouteNetworkDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustNetworkRouteNetworkDeleteResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zeroTrustNetworkRouteNetworkDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustNetworkRouteNetworkDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustNetworkRouteNetworkDeleteResponseEnvelopeErrors]
type zeroTrustNetworkRouteNetworkDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkRouteNetworkDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustNetworkRouteNetworkDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustNetworkRouteNetworkDeleteResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    zeroTrustNetworkRouteNetworkDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustNetworkRouteNetworkDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustNetworkRouteNetworkDeleteResponseEnvelopeMessages]
type zeroTrustNetworkRouteNetworkDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkRouteNetworkDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustNetworkRouteNetworkDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustNetworkRouteNetworkDeleteResponseEnvelopeSuccess bool

const (
	ZeroTrustNetworkRouteNetworkDeleteResponseEnvelopeSuccessTrue ZeroTrustNetworkRouteNetworkDeleteResponseEnvelopeSuccess = true
)

type ZeroTrustNetworkRouteNetworkEditParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZeroTrustNetworkRouteNetworkEditResponseEnvelope struct {
	Errors   []ZeroTrustNetworkRouteNetworkEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustNetworkRouteNetworkEditResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustNetworkRouteNetworkEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustNetworkRouteNetworkEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustNetworkRouteNetworkEditResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustNetworkRouteNetworkEditResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustNetworkRouteNetworkEditResponseEnvelope]
type zeroTrustNetworkRouteNetworkEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkRouteNetworkEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustNetworkRouteNetworkEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustNetworkRouteNetworkEditResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zeroTrustNetworkRouteNetworkEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustNetworkRouteNetworkEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustNetworkRouteNetworkEditResponseEnvelopeErrors]
type zeroTrustNetworkRouteNetworkEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkRouteNetworkEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustNetworkRouteNetworkEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustNetworkRouteNetworkEditResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zeroTrustNetworkRouteNetworkEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustNetworkRouteNetworkEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustNetworkRouteNetworkEditResponseEnvelopeMessages]
type zeroTrustNetworkRouteNetworkEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkRouteNetworkEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustNetworkRouteNetworkEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustNetworkRouteNetworkEditResponseEnvelopeSuccess bool

const (
	ZeroTrustNetworkRouteNetworkEditResponseEnvelopeSuccessTrue ZeroTrustNetworkRouteNetworkEditResponseEnvelopeSuccess = true
)
