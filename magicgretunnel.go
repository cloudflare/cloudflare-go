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

// MagicGreTunnelService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewMagicGreTunnelService] method
// instead.
type MagicGreTunnelService struct {
	Options []option.RequestOption
}

// NewMagicGreTunnelService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMagicGreTunnelService(opts ...option.RequestOption) (r *MagicGreTunnelService) {
	r = &MagicGreTunnelService{}
	r.Options = opts
	return
}

// Creates new GRE tunnels. Use `?validate_only=true` as an optional query
// parameter to only run validation without persisting changes.
func (r *MagicGreTunnelService) New(ctx context.Context, accountIdentifier string, body MagicGreTunnelNewParams, opts ...option.RequestOption) (res *MagicGreTunnelNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicGreTunnelNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a specific GRE tunnel. Use `?validate_only=true` as an optional query
// parameter to only run validation without persisting changes.
func (r *MagicGreTunnelService) Update(ctx context.Context, accountIdentifier string, tunnelIdentifier string, body MagicGreTunnelUpdateParams, opts ...option.RequestOption) (res *MagicGreTunnelUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicGreTunnelUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists GRE tunnels associated with an account.
func (r *MagicGreTunnelService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *MagicGreTunnelListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicGreTunnelListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Disables and removes a specific static GRE tunnel. Use `?validate_only=true` as
// an optional query parameter to only run validation without persisting changes.
func (r *MagicGreTunnelService) Delete(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *MagicGreTunnelDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicGreTunnelDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists informtion for a specific GRE tunnel.
func (r *MagicGreTunnelService) Get(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *MagicGreTunnelGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicGreTunnelGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MagicGreTunnelNewResponse struct {
	GreTunnels []MagicGreTunnelNewResponseGreTunnel `json:"gre_tunnels"`
	JSON       magicGreTunnelNewResponseJSON        `json:"-"`
}

// magicGreTunnelNewResponseJSON contains the JSON metadata for the struct
// [MagicGreTunnelNewResponse]
type magicGreTunnelNewResponseJSON struct {
	GreTunnels  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGreTunnelNewResponseGreTunnel struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel.
	CloudflareGreEndpoint string `json:"cloudflare_gre_endpoint,required"`
	// The IP address assigned to the customer side of the GRE tunnel.
	CustomerGreEndpoint string `json:"customer_gre_endpoint,required"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress string `json:"interface_address,required"`
	// The name of the tunnel. The name cannot contain spaces or special characters,
	// must be 15 characters or less, and cannot share a name with another GRE tunnel.
	Name string `json:"name,required"`
	// Tunnel identifier tag.
	ID string `json:"id"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the GRE tunnel.
	Description string                                         `json:"description"`
	HealthCheck MagicGreTunnelNewResponseGreTunnelsHealthCheck `json:"health_check"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value
	// is 576.
	Mtu int64 `json:"mtu"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	TTL  int64                                  `json:"ttl"`
	JSON magicGreTunnelNewResponseGreTunnelJSON `json:"-"`
}

// magicGreTunnelNewResponseGreTunnelJSON contains the JSON metadata for the struct
// [MagicGreTunnelNewResponseGreTunnel]
type magicGreTunnelNewResponseGreTunnelJSON struct {
	CloudflareGreEndpoint apijson.Field
	CustomerGreEndpoint   apijson.Field
	InterfaceAddress      apijson.Field
	Name                  apijson.Field
	ID                    apijson.Field
	CreatedOn             apijson.Field
	Description           apijson.Field
	HealthCheck           apijson.Field
	ModifiedOn            apijson.Field
	Mtu                   apijson.Field
	TTL                   apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *MagicGreTunnelNewResponseGreTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGreTunnelNewResponseGreTunnelsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel.
	Direction MagicGreTunnelNewResponseGreTunnelsHealthCheckDirection `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate MagicGreTunnelNewResponseGreTunnelsHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type MagicGreTunnelNewResponseGreTunnelsHealthCheckType `json:"type"`
	JSON magicGreTunnelNewResponseGreTunnelsHealthCheckJSON `json:"-"`
}

// magicGreTunnelNewResponseGreTunnelsHealthCheckJSON contains the JSON metadata
// for the struct [MagicGreTunnelNewResponseGreTunnelsHealthCheck]
type magicGreTunnelNewResponseGreTunnelsHealthCheckJSON struct {
	Direction   apijson.Field
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelNewResponseGreTunnelsHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel.
type MagicGreTunnelNewResponseGreTunnelsHealthCheckDirection string

const (
	MagicGreTunnelNewResponseGreTunnelsHealthCheckDirectionUnidirectional MagicGreTunnelNewResponseGreTunnelsHealthCheckDirection = "unidirectional"
	MagicGreTunnelNewResponseGreTunnelsHealthCheckDirectionBidirectional  MagicGreTunnelNewResponseGreTunnelsHealthCheckDirection = "bidirectional"
)

// How frequent the health check is run. The default value is `mid`.
type MagicGreTunnelNewResponseGreTunnelsHealthCheckRate string

const (
	MagicGreTunnelNewResponseGreTunnelsHealthCheckRateLow  MagicGreTunnelNewResponseGreTunnelsHealthCheckRate = "low"
	MagicGreTunnelNewResponseGreTunnelsHealthCheckRateMid  MagicGreTunnelNewResponseGreTunnelsHealthCheckRate = "mid"
	MagicGreTunnelNewResponseGreTunnelsHealthCheckRateHigh MagicGreTunnelNewResponseGreTunnelsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicGreTunnelNewResponseGreTunnelsHealthCheckType string

const (
	MagicGreTunnelNewResponseGreTunnelsHealthCheckTypeReply   MagicGreTunnelNewResponseGreTunnelsHealthCheckType = "reply"
	MagicGreTunnelNewResponseGreTunnelsHealthCheckTypeRequest MagicGreTunnelNewResponseGreTunnelsHealthCheckType = "request"
)

type MagicGreTunnelUpdateResponse struct {
	Modified          bool                             `json:"modified"`
	ModifiedGreTunnel interface{}                      `json:"modified_gre_tunnel"`
	JSON              magicGreTunnelUpdateResponseJSON `json:"-"`
}

// magicGreTunnelUpdateResponseJSON contains the JSON metadata for the struct
// [MagicGreTunnelUpdateResponse]
type magicGreTunnelUpdateResponseJSON struct {
	Modified          apijson.Field
	ModifiedGreTunnel apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *MagicGreTunnelUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGreTunnelListResponse struct {
	GreTunnels []MagicGreTunnelListResponseGreTunnel `json:"gre_tunnels"`
	JSON       magicGreTunnelListResponseJSON        `json:"-"`
}

// magicGreTunnelListResponseJSON contains the JSON metadata for the struct
// [MagicGreTunnelListResponse]
type magicGreTunnelListResponseJSON struct {
	GreTunnels  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGreTunnelListResponseGreTunnel struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel.
	CloudflareGreEndpoint string `json:"cloudflare_gre_endpoint,required"`
	// The IP address assigned to the customer side of the GRE tunnel.
	CustomerGreEndpoint string `json:"customer_gre_endpoint,required"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress string `json:"interface_address,required"`
	// The name of the tunnel. The name cannot contain spaces or special characters,
	// must be 15 characters or less, and cannot share a name with another GRE tunnel.
	Name string `json:"name,required"`
	// Tunnel identifier tag.
	ID string `json:"id"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the GRE tunnel.
	Description string                                          `json:"description"`
	HealthCheck MagicGreTunnelListResponseGreTunnelsHealthCheck `json:"health_check"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value
	// is 576.
	Mtu int64 `json:"mtu"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	TTL  int64                                   `json:"ttl"`
	JSON magicGreTunnelListResponseGreTunnelJSON `json:"-"`
}

// magicGreTunnelListResponseGreTunnelJSON contains the JSON metadata for the
// struct [MagicGreTunnelListResponseGreTunnel]
type magicGreTunnelListResponseGreTunnelJSON struct {
	CloudflareGreEndpoint apijson.Field
	CustomerGreEndpoint   apijson.Field
	InterfaceAddress      apijson.Field
	Name                  apijson.Field
	ID                    apijson.Field
	CreatedOn             apijson.Field
	Description           apijson.Field
	HealthCheck           apijson.Field
	ModifiedOn            apijson.Field
	Mtu                   apijson.Field
	TTL                   apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *MagicGreTunnelListResponseGreTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGreTunnelListResponseGreTunnelsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel.
	Direction MagicGreTunnelListResponseGreTunnelsHealthCheckDirection `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate MagicGreTunnelListResponseGreTunnelsHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type MagicGreTunnelListResponseGreTunnelsHealthCheckType `json:"type"`
	JSON magicGreTunnelListResponseGreTunnelsHealthCheckJSON `json:"-"`
}

// magicGreTunnelListResponseGreTunnelsHealthCheckJSON contains the JSON metadata
// for the struct [MagicGreTunnelListResponseGreTunnelsHealthCheck]
type magicGreTunnelListResponseGreTunnelsHealthCheckJSON struct {
	Direction   apijson.Field
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelListResponseGreTunnelsHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel.
type MagicGreTunnelListResponseGreTunnelsHealthCheckDirection string

const (
	MagicGreTunnelListResponseGreTunnelsHealthCheckDirectionUnidirectional MagicGreTunnelListResponseGreTunnelsHealthCheckDirection = "unidirectional"
	MagicGreTunnelListResponseGreTunnelsHealthCheckDirectionBidirectional  MagicGreTunnelListResponseGreTunnelsHealthCheckDirection = "bidirectional"
)

// How frequent the health check is run. The default value is `mid`.
type MagicGreTunnelListResponseGreTunnelsHealthCheckRate string

const (
	MagicGreTunnelListResponseGreTunnelsHealthCheckRateLow  MagicGreTunnelListResponseGreTunnelsHealthCheckRate = "low"
	MagicGreTunnelListResponseGreTunnelsHealthCheckRateMid  MagicGreTunnelListResponseGreTunnelsHealthCheckRate = "mid"
	MagicGreTunnelListResponseGreTunnelsHealthCheckRateHigh MagicGreTunnelListResponseGreTunnelsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicGreTunnelListResponseGreTunnelsHealthCheckType string

const (
	MagicGreTunnelListResponseGreTunnelsHealthCheckTypeReply   MagicGreTunnelListResponseGreTunnelsHealthCheckType = "reply"
	MagicGreTunnelListResponseGreTunnelsHealthCheckTypeRequest MagicGreTunnelListResponseGreTunnelsHealthCheckType = "request"
)

type MagicGreTunnelDeleteResponse struct {
	Deleted          bool                             `json:"deleted"`
	DeletedGreTunnel interface{}                      `json:"deleted_gre_tunnel"`
	JSON             magicGreTunnelDeleteResponseJSON `json:"-"`
}

// magicGreTunnelDeleteResponseJSON contains the JSON metadata for the struct
// [MagicGreTunnelDeleteResponse]
type magicGreTunnelDeleteResponseJSON struct {
	Deleted          apijson.Field
	DeletedGreTunnel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MagicGreTunnelDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGreTunnelGetResponse struct {
	GreTunnel interface{}                   `json:"gre_tunnel"`
	JSON      magicGreTunnelGetResponseJSON `json:"-"`
}

// magicGreTunnelGetResponseJSON contains the JSON metadata for the struct
// [MagicGreTunnelGetResponse]
type magicGreTunnelGetResponseJSON struct {
	GreTunnel   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGreTunnelNewParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r MagicGreTunnelNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type MagicGreTunnelNewResponseEnvelope struct {
	Errors   []MagicGreTunnelNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicGreTunnelNewResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicGreTunnelNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicGreTunnelNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicGreTunnelNewResponseEnvelopeJSON    `json:"-"`
}

// magicGreTunnelNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [MagicGreTunnelNewResponseEnvelope]
type magicGreTunnelNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGreTunnelNewResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    magicGreTunnelNewResponseEnvelopeErrorsJSON `json:"-"`
}

// magicGreTunnelNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MagicGreTunnelNewResponseEnvelopeErrors]
type magicGreTunnelNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGreTunnelNewResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    magicGreTunnelNewResponseEnvelopeMessagesJSON `json:"-"`
}

// magicGreTunnelNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MagicGreTunnelNewResponseEnvelopeMessages]
type magicGreTunnelNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicGreTunnelNewResponseEnvelopeSuccess bool

const (
	MagicGreTunnelNewResponseEnvelopeSuccessTrue MagicGreTunnelNewResponseEnvelopeSuccess = true
)

type MagicGreTunnelUpdateParams struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel.
	CloudflareGreEndpoint param.Field[string] `json:"cloudflare_gre_endpoint,required"`
	// The IP address assigned to the customer side of the GRE tunnel.
	CustomerGreEndpoint param.Field[string] `json:"customer_gre_endpoint,required"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress param.Field[string] `json:"interface_address,required"`
	// The name of the tunnel. The name cannot contain spaces or special characters,
	// must be 15 characters or less, and cannot share a name with another GRE tunnel.
	Name param.Field[string] `json:"name,required"`
	// An optional description of the GRE tunnel.
	Description param.Field[string]                                `json:"description"`
	HealthCheck param.Field[MagicGreTunnelUpdateParamsHealthCheck] `json:"health_check"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value
	// is 576.
	Mtu param.Field[int64] `json:"mtu"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	TTL param.Field[int64] `json:"ttl"`
}

func (r MagicGreTunnelUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagicGreTunnelUpdateParamsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel.
	Direction param.Field[MagicGreTunnelUpdateParamsHealthCheckDirection] `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled param.Field[bool] `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate param.Field[MagicGreTunnelUpdateParamsHealthCheckRate] `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target param.Field[string] `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type param.Field[MagicGreTunnelUpdateParamsHealthCheckType] `json:"type"`
}

func (r MagicGreTunnelUpdateParamsHealthCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel.
type MagicGreTunnelUpdateParamsHealthCheckDirection string

const (
	MagicGreTunnelUpdateParamsHealthCheckDirectionUnidirectional MagicGreTunnelUpdateParamsHealthCheckDirection = "unidirectional"
	MagicGreTunnelUpdateParamsHealthCheckDirectionBidirectional  MagicGreTunnelUpdateParamsHealthCheckDirection = "bidirectional"
)

// How frequent the health check is run. The default value is `mid`.
type MagicGreTunnelUpdateParamsHealthCheckRate string

const (
	MagicGreTunnelUpdateParamsHealthCheckRateLow  MagicGreTunnelUpdateParamsHealthCheckRate = "low"
	MagicGreTunnelUpdateParamsHealthCheckRateMid  MagicGreTunnelUpdateParamsHealthCheckRate = "mid"
	MagicGreTunnelUpdateParamsHealthCheckRateHigh MagicGreTunnelUpdateParamsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicGreTunnelUpdateParamsHealthCheckType string

const (
	MagicGreTunnelUpdateParamsHealthCheckTypeReply   MagicGreTunnelUpdateParamsHealthCheckType = "reply"
	MagicGreTunnelUpdateParamsHealthCheckTypeRequest MagicGreTunnelUpdateParamsHealthCheckType = "request"
)

type MagicGreTunnelUpdateResponseEnvelope struct {
	Errors   []MagicGreTunnelUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicGreTunnelUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicGreTunnelUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicGreTunnelUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicGreTunnelUpdateResponseEnvelopeJSON    `json:"-"`
}

// magicGreTunnelUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [MagicGreTunnelUpdateResponseEnvelope]
type magicGreTunnelUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGreTunnelUpdateResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    magicGreTunnelUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// magicGreTunnelUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [MagicGreTunnelUpdateResponseEnvelopeErrors]
type magicGreTunnelUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGreTunnelUpdateResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    magicGreTunnelUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// magicGreTunnelUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [MagicGreTunnelUpdateResponseEnvelopeMessages]
type magicGreTunnelUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicGreTunnelUpdateResponseEnvelopeSuccess bool

const (
	MagicGreTunnelUpdateResponseEnvelopeSuccessTrue MagicGreTunnelUpdateResponseEnvelopeSuccess = true
)

type MagicGreTunnelListResponseEnvelope struct {
	Errors   []MagicGreTunnelListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicGreTunnelListResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicGreTunnelListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicGreTunnelListResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicGreTunnelListResponseEnvelopeJSON    `json:"-"`
}

// magicGreTunnelListResponseEnvelopeJSON contains the JSON metadata for the struct
// [MagicGreTunnelListResponseEnvelope]
type magicGreTunnelListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGreTunnelListResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    magicGreTunnelListResponseEnvelopeErrorsJSON `json:"-"`
}

// magicGreTunnelListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MagicGreTunnelListResponseEnvelopeErrors]
type magicGreTunnelListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGreTunnelListResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    magicGreTunnelListResponseEnvelopeMessagesJSON `json:"-"`
}

// magicGreTunnelListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [MagicGreTunnelListResponseEnvelopeMessages]
type magicGreTunnelListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicGreTunnelListResponseEnvelopeSuccess bool

const (
	MagicGreTunnelListResponseEnvelopeSuccessTrue MagicGreTunnelListResponseEnvelopeSuccess = true
)

type MagicGreTunnelDeleteResponseEnvelope struct {
	Errors   []MagicGreTunnelDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicGreTunnelDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicGreTunnelDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicGreTunnelDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicGreTunnelDeleteResponseEnvelopeJSON    `json:"-"`
}

// magicGreTunnelDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [MagicGreTunnelDeleteResponseEnvelope]
type magicGreTunnelDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGreTunnelDeleteResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    magicGreTunnelDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// magicGreTunnelDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [MagicGreTunnelDeleteResponseEnvelopeErrors]
type magicGreTunnelDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGreTunnelDeleteResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    magicGreTunnelDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// magicGreTunnelDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [MagicGreTunnelDeleteResponseEnvelopeMessages]
type magicGreTunnelDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicGreTunnelDeleteResponseEnvelopeSuccess bool

const (
	MagicGreTunnelDeleteResponseEnvelopeSuccessTrue MagicGreTunnelDeleteResponseEnvelopeSuccess = true
)

type MagicGreTunnelGetResponseEnvelope struct {
	Errors   []MagicGreTunnelGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicGreTunnelGetResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicGreTunnelGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicGreTunnelGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicGreTunnelGetResponseEnvelopeJSON    `json:"-"`
}

// magicGreTunnelGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [MagicGreTunnelGetResponseEnvelope]
type magicGreTunnelGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGreTunnelGetResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    magicGreTunnelGetResponseEnvelopeErrorsJSON `json:"-"`
}

// magicGreTunnelGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MagicGreTunnelGetResponseEnvelopeErrors]
type magicGreTunnelGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGreTunnelGetResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    magicGreTunnelGetResponseEnvelopeMessagesJSON `json:"-"`
}

// magicGreTunnelGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MagicGreTunnelGetResponseEnvelopeMessages]
type magicGreTunnelGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicGreTunnelGetResponseEnvelopeSuccess bool

const (
	MagicGreTunnelGetResponseEnvelopeSuccessTrue MagicGreTunnelGetResponseEnvelopeSuccess = true
)
