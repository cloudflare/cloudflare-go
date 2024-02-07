// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
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
	Options []option.RequestOption
}

// NewTeamnetRouteService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTeamnetRouteService(opts ...option.RequestOption) (r *TeamnetRouteService) {
	r = &TeamnetRouteService{}
	r.Options = opts
	return
}

// Routes a private network through a Cloudflare Tunnel.
func (r *TeamnetRouteService) New(ctx context.Context, accountID string, body TeamnetRouteNewParams, opts ...option.RequestOption) (res *TeamnetRouteNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TeamnetRouteNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/routes", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing private network route in an account. The fields that are
// meant to be updated should be provided in the body of the request.
func (r *TeamnetRouteService) Update(ctx context.Context, accountID string, routeID string, body TeamnetRouteUpdateParams, opts ...option.RequestOption) (res *TeamnetRouteUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TeamnetRouteUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/routes/%s", accountID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a private network route from an account.
func (r *TeamnetRouteService) Delete(ctx context.Context, accountID string, routeID string, opts ...option.RequestOption) (res *TeamnetRouteDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TeamnetRouteDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/routes/%s", accountID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TeamnetRouteNewResponse struct {
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
	VirtualNetworkID interface{}                 `json:"virtual_network_id"`
	JSON             teamnetRouteNewResponseJSON `json:"-"`
}

// teamnetRouteNewResponseJSON contains the JSON metadata for the struct
// [TeamnetRouteNewResponse]
type teamnetRouteNewResponseJSON struct {
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

func (r *TeamnetRouteNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteUpdateResponse struct {
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
	VirtualNetworkID interface{}                    `json:"virtual_network_id"`
	JSON             teamnetRouteUpdateResponseJSON `json:"-"`
}

// teamnetRouteUpdateResponseJSON contains the JSON metadata for the struct
// [TeamnetRouteUpdateResponse]
type teamnetRouteUpdateResponseJSON struct {
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

func (r *TeamnetRouteUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteDeleteResponse struct {
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
	VirtualNetworkID interface{}                    `json:"virtual_network_id"`
	JSON             teamnetRouteDeleteResponseJSON `json:"-"`
}

// teamnetRouteDeleteResponseJSON contains the JSON metadata for the struct
// [TeamnetRouteDeleteResponse]
type teamnetRouteDeleteResponseJSON struct {
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

func (r *TeamnetRouteDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteNewParams struct {
	// The private IPv4 or IPv6 range connected by the route, in CIDR notation.
	IPNetwork param.Field[string] `json:"ip_network,required"`
	// Optional remark describing the route.
	Comment param.Field[string] `json:"comment"`
	// UUID of the Tunnel Virtual Network this route belongs to. If no virtual networks
	// are configured, the route is assigned to the default virtual network of the
	// account.
	VirtualNetworkID param.Field[interface{}] `json:"virtual_network_id"`
}

func (r TeamnetRouteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TeamnetRouteNewResponseEnvelope struct {
	Errors   []TeamnetRouteNewResponseEnvelopeErrors   `json:"errors"`
	Messages []TeamnetRouteNewResponseEnvelopeMessages `json:"messages"`
	Result   TeamnetRouteNewResponse                   `json:"result"`
	// Whether the API call was successful
	Success TeamnetRouteNewResponseEnvelopeSuccess `json:"success"`
	JSON    teamnetRouteNewResponseEnvelopeJSON    `json:"-"`
}

// teamnetRouteNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [TeamnetRouteNewResponseEnvelope]
type teamnetRouteNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteNewResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    teamnetRouteNewResponseEnvelopeErrorsJSON `json:"-"`
}

// teamnetRouteNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [TeamnetRouteNewResponseEnvelopeErrors]
type teamnetRouteNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteNewResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    teamnetRouteNewResponseEnvelopeMessagesJSON `json:"-"`
}

// teamnetRouteNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [TeamnetRouteNewResponseEnvelopeMessages]
type teamnetRouteNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TeamnetRouteNewResponseEnvelopeSuccess bool

const (
	TeamnetRouteNewResponseEnvelopeSuccessTrue TeamnetRouteNewResponseEnvelopeSuccess = true
)

type TeamnetRouteUpdateParams struct {
	// Optional remark describing the route.
	Comment param.Field[string] `json:"comment"`
	// The private IPv4 or IPv6 range connected by the route, in CIDR notation.
	Network param.Field[string] `json:"network"`
	// The type of tunnel.
	TunType param.Field[TeamnetRouteUpdateParamsTunType] `json:"tun_type"`
	// UUID of the Cloudflare Tunnel serving the route.
	TunnelID param.Field[interface{}] `json:"tunnel_id"`
	// UUID of the Tunnel Virtual Network this route belongs to. If no virtual networks
	// are configured, the route is assigned to the default virtual network of the
	// account.
	VirtualNetworkID param.Field[interface{}] `json:"virtual_network_id"`
}

func (r TeamnetRouteUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of tunnel.
type TeamnetRouteUpdateParamsTunType string

const (
	TeamnetRouteUpdateParamsTunTypeCfdTunnel     TeamnetRouteUpdateParamsTunType = "cfd_tunnel"
	TeamnetRouteUpdateParamsTunTypeWarpConnector TeamnetRouteUpdateParamsTunType = "warp_connector"
	TeamnetRouteUpdateParamsTunTypeIPSec         TeamnetRouteUpdateParamsTunType = "ip_sec"
	TeamnetRouteUpdateParamsTunTypeGre           TeamnetRouteUpdateParamsTunType = "gre"
	TeamnetRouteUpdateParamsTunTypeCni           TeamnetRouteUpdateParamsTunType = "cni"
)

type TeamnetRouteUpdateResponseEnvelope struct {
	Errors   []TeamnetRouteUpdateResponseEnvelopeErrors   `json:"errors"`
	Messages []TeamnetRouteUpdateResponseEnvelopeMessages `json:"messages"`
	Result   TeamnetRouteUpdateResponse                   `json:"result"`
	// Whether the API call was successful
	Success TeamnetRouteUpdateResponseEnvelopeSuccess `json:"success"`
	JSON    teamnetRouteUpdateResponseEnvelopeJSON    `json:"-"`
}

// teamnetRouteUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [TeamnetRouteUpdateResponseEnvelope]
type teamnetRouteUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteUpdateResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    teamnetRouteUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// teamnetRouteUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [TeamnetRouteUpdateResponseEnvelopeErrors]
type teamnetRouteUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteUpdateResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    teamnetRouteUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// teamnetRouteUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [TeamnetRouteUpdateResponseEnvelopeMessages]
type teamnetRouteUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TeamnetRouteUpdateResponseEnvelopeSuccess bool

const (
	TeamnetRouteUpdateResponseEnvelopeSuccessTrue TeamnetRouteUpdateResponseEnvelopeSuccess = true
)

type TeamnetRouteDeleteResponseEnvelope struct {
	Errors   []TeamnetRouteDeleteResponseEnvelopeErrors   `json:"errors"`
	Messages []TeamnetRouteDeleteResponseEnvelopeMessages `json:"messages"`
	Result   TeamnetRouteDeleteResponse                   `json:"result"`
	// Whether the API call was successful
	Success TeamnetRouteDeleteResponseEnvelopeSuccess `json:"success"`
	JSON    teamnetRouteDeleteResponseEnvelopeJSON    `json:"-"`
}

// teamnetRouteDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [TeamnetRouteDeleteResponseEnvelope]
type teamnetRouteDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteDeleteResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    teamnetRouteDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// teamnetRouteDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [TeamnetRouteDeleteResponseEnvelopeErrors]
type teamnetRouteDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetRouteDeleteResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    teamnetRouteDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// teamnetRouteDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [TeamnetRouteDeleteResponseEnvelopeMessages]
type teamnetRouteDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetRouteDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TeamnetRouteDeleteResponseEnvelopeSuccess bool

const (
	TeamnetRouteDeleteResponseEnvelopeSuccessTrue TeamnetRouteDeleteResponseEnvelopeSuccess = true
)
