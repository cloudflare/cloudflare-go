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

// TeamnetRouteNetworkService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewTeamnetRouteNetworkService]
// method instead.
type TeamnetRouteNetworkService struct {
	Options []option.RequestOption
}

// NewTeamnetRouteNetworkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTeamnetRouteNetworkService(opts ...option.RequestOption) (r *TeamnetRouteNetworkService) {
	r = &TeamnetRouteNetworkService{}
	r.Options = opts
	return
}

// Routes a private network through a Cloudflare Tunnel. The CIDR in
// `ip_network_encoded` must be written in URL-encoded format.
func (r *TeamnetRouteNetworkService) New(ctx context.Context, ipNetworkEncoded string, params TeamnetRouteNetworkNewParams, opts ...option.RequestOption) (res *TeamnetRouteNetworkNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TeamnetRouteNetworkNewResponseEnvelope
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
func (r *TeamnetRouteNetworkService) Delete(ctx context.Context, ipNetworkEncoded string, params TeamnetRouteNetworkDeleteParams, opts ...option.RequestOption) (res *TeamnetRouteNetworkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TeamnetRouteNetworkDeleteResponseEnvelope
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
func (r *TeamnetRouteNetworkService) Edit(ctx context.Context, ipNetworkEncoded string, body TeamnetRouteNetworkEditParams, opts ...option.RequestOption) (res *TeamnetRouteNetworkEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TeamnetRouteNetworkEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/routes/network/%s", body.AccountID, ipNetworkEncoded)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TeamnetRouteNetworkNewResponse struct {
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
	JSON             teamnetRouteNetworkNewResponseJSON `json:"-"`
}

// teamnetRouteNetworkNewResponseJSON contains the JSON metadata for the struct
// [TeamnetRouteNetworkNewResponse]
type teamnetRouteNetworkNewResponseJSON struct {
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

func (r *TeamnetRouteNetworkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteNetworkDeleteResponse struct {
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
	JSON             teamnetRouteNetworkDeleteResponseJSON `json:"-"`
}

// teamnetRouteNetworkDeleteResponseJSON contains the JSON metadata for the struct
// [TeamnetRouteNetworkDeleteResponse]
type teamnetRouteNetworkDeleteResponseJSON struct {
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

func (r *TeamnetRouteNetworkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteNetworkEditResponse struct {
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
	JSON             teamnetRouteNetworkEditResponseJSON `json:"-"`
}

// teamnetRouteNetworkEditResponseJSON contains the JSON metadata for the struct
// [TeamnetRouteNetworkEditResponse]
type teamnetRouteNetworkEditResponseJSON struct {
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

func (r *TeamnetRouteNetworkEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteNetworkNewParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// Optional remark describing the route.
	Comment param.Field[string] `json:"comment"`
	// UUID of the Tunnel Virtual Network this route belongs to. If no virtual networks
	// are configured, the route is assigned to the default virtual network of the
	// account.
	VirtualNetworkID param.Field[interface{}] `json:"virtual_network_id"`
}

func (r TeamnetRouteNetworkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TeamnetRouteNetworkNewResponseEnvelope struct {
	Errors   []TeamnetRouteNetworkNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TeamnetRouteNetworkNewResponseEnvelopeMessages `json:"messages,required"`
	Result   TeamnetRouteNetworkNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success TeamnetRouteNetworkNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    teamnetRouteNetworkNewResponseEnvelopeJSON    `json:"-"`
}

// teamnetRouteNetworkNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [TeamnetRouteNetworkNewResponseEnvelope]
type teamnetRouteNetworkNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteNetworkNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteNetworkNewResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    teamnetRouteNetworkNewResponseEnvelopeErrorsJSON `json:"-"`
}

// teamnetRouteNetworkNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [TeamnetRouteNetworkNewResponseEnvelopeErrors]
type teamnetRouteNetworkNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteNetworkNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteNetworkNewResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    teamnetRouteNetworkNewResponseEnvelopeMessagesJSON `json:"-"`
}

// teamnetRouteNetworkNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [TeamnetRouteNetworkNewResponseEnvelopeMessages]
type teamnetRouteNetworkNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteNetworkNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TeamnetRouteNetworkNewResponseEnvelopeSuccess bool

const (
	TeamnetRouteNetworkNewResponseEnvelopeSuccessTrue TeamnetRouteNetworkNewResponseEnvelopeSuccess = true
)

type TeamnetRouteNetworkDeleteParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// The type of tunnel.
	TunType param.Field[TeamnetRouteNetworkDeleteParamsTunType] `query:"tun_type"`
}

// URLQuery serializes [TeamnetRouteNetworkDeleteParams]'s query parameters as
// `url.Values`.
func (r TeamnetRouteNetworkDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The type of tunnel.
type TeamnetRouteNetworkDeleteParamsTunType string

const (
	TeamnetRouteNetworkDeleteParamsTunTypeCfdTunnel     TeamnetRouteNetworkDeleteParamsTunType = "cfd_tunnel"
	TeamnetRouteNetworkDeleteParamsTunTypeWARPConnector TeamnetRouteNetworkDeleteParamsTunType = "warp_connector"
	TeamnetRouteNetworkDeleteParamsTunTypeIPSec         TeamnetRouteNetworkDeleteParamsTunType = "ip_sec"
	TeamnetRouteNetworkDeleteParamsTunTypeGRE           TeamnetRouteNetworkDeleteParamsTunType = "gre"
	TeamnetRouteNetworkDeleteParamsTunTypeCni           TeamnetRouteNetworkDeleteParamsTunType = "cni"
)

type TeamnetRouteNetworkDeleteResponseEnvelope struct {
	Errors   []TeamnetRouteNetworkDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TeamnetRouteNetworkDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   TeamnetRouteNetworkDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success TeamnetRouteNetworkDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    teamnetRouteNetworkDeleteResponseEnvelopeJSON    `json:"-"`
}

// teamnetRouteNetworkDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [TeamnetRouteNetworkDeleteResponseEnvelope]
type teamnetRouteNetworkDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteNetworkDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteNetworkDeleteResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    teamnetRouteNetworkDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// teamnetRouteNetworkDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [TeamnetRouteNetworkDeleteResponseEnvelopeErrors]
type teamnetRouteNetworkDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteNetworkDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteNetworkDeleteResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    teamnetRouteNetworkDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// teamnetRouteNetworkDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [TeamnetRouteNetworkDeleteResponseEnvelopeMessages]
type teamnetRouteNetworkDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteNetworkDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TeamnetRouteNetworkDeleteResponseEnvelopeSuccess bool

const (
	TeamnetRouteNetworkDeleteResponseEnvelopeSuccessTrue TeamnetRouteNetworkDeleteResponseEnvelopeSuccess = true
)

type TeamnetRouteNetworkEditParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type TeamnetRouteNetworkEditResponseEnvelope struct {
	Errors   []TeamnetRouteNetworkEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TeamnetRouteNetworkEditResponseEnvelopeMessages `json:"messages,required"`
	Result   TeamnetRouteNetworkEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success TeamnetRouteNetworkEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    teamnetRouteNetworkEditResponseEnvelopeJSON    `json:"-"`
}

// teamnetRouteNetworkEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [TeamnetRouteNetworkEditResponseEnvelope]
type teamnetRouteNetworkEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteNetworkEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteNetworkEditResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    teamnetRouteNetworkEditResponseEnvelopeErrorsJSON `json:"-"`
}

// teamnetRouteNetworkEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [TeamnetRouteNetworkEditResponseEnvelopeErrors]
type teamnetRouteNetworkEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteNetworkEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteNetworkEditResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    teamnetRouteNetworkEditResponseEnvelopeMessagesJSON `json:"-"`
}

// teamnetRouteNetworkEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [TeamnetRouteNetworkEditResponseEnvelopeMessages]
type teamnetRouteNetworkEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteNetworkEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TeamnetRouteNetworkEditResponseEnvelopeSuccess bool

const (
	TeamnetRouteNetworkEditResponseEnvelopeSuccessTrue TeamnetRouteNetworkEditResponseEnvelopeSuccess = true
)
