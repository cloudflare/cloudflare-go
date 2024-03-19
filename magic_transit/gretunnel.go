// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// GRETunnelService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewGRETunnelService] method instead.
type GRETunnelService struct {
	Options []option.RequestOption
}

// NewGRETunnelService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGRETunnelService(opts ...option.RequestOption) (r *GRETunnelService) {
	r = &GRETunnelService{}
	r.Options = opts
	return
}

// Creates new GRE tunnels. Use `?validate_only=true` as an optional query
// parameter to only run validation without persisting changes.
func (r *GRETunnelService) New(ctx context.Context, accountIdentifier string, body GRETunnelNewParams, opts ...option.RequestOption) (res *GRETunnelNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GRETunnelNewResponseEnvelope
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
func (r *GRETunnelService) Update(ctx context.Context, accountIdentifier string, tunnelIdentifier string, body GRETunnelUpdateParams, opts ...option.RequestOption) (res *GRETunnelUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GRETunnelUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists GRE tunnels associated with an account.
func (r *GRETunnelService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *GRETunnelListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GRETunnelListResponseEnvelope
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
func (r *GRETunnelService) Delete(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *GRETunnelDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GRETunnelDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists informtion for a specific GRE tunnel.
func (r *GRETunnelService) Get(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *GRETunnelGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GRETunnelGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type GRETunnelNewResponse struct {
	GRETunnels []GRETunnelNewResponseGRETunnel `json:"gre_tunnels"`
	JSON       greTunnelNewResponseJSON        `json:"-"`
}

// greTunnelNewResponseJSON contains the JSON metadata for the struct
// [GRETunnelNewResponse]
type greTunnelNewResponseJSON struct {
	GRETunnels  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelNewResponseJSON) RawJSON() string {
	return r.raw
}

type GRETunnelNewResponseGRETunnel struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel.
	CloudflareGREEndpoint string `json:"cloudflare_gre_endpoint,required"`
	// The IP address assigned to the customer side of the GRE tunnel.
	CustomerGREEndpoint string `json:"customer_gre_endpoint,required"`
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
	Description string                                    `json:"description"`
	HealthCheck GRETunnelNewResponseGRETunnelsHealthCheck `json:"health_check"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value
	// is 576.
	Mtu int64 `json:"mtu"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	TTL  int64                             `json:"ttl"`
	JSON greTunnelNewResponseGRETunnelJSON `json:"-"`
}

// greTunnelNewResponseGRETunnelJSON contains the JSON metadata for the struct
// [GRETunnelNewResponseGRETunnel]
type greTunnelNewResponseGRETunnelJSON struct {
	CloudflareGREEndpoint apijson.Field
	CustomerGREEndpoint   apijson.Field
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

func (r *GRETunnelNewResponseGRETunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelNewResponseGRETunnelJSON) RawJSON() string {
	return r.raw
}

type GRETunnelNewResponseGRETunnelsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
	// in health_check is ignored as the interface_address is used to send traffic into
	// the tunnel.
	Direction GRETunnelNewResponseGRETunnelsHealthCheckDirection `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate GRETunnelNewResponseGRETunnelsHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type GRETunnelNewResponseGRETunnelsHealthCheckType `json:"type"`
	JSON greTunnelNewResponseGRETunnelsHealthCheckJSON `json:"-"`
}

// greTunnelNewResponseGRETunnelsHealthCheckJSON contains the JSON metadata for the
// struct [GRETunnelNewResponseGRETunnelsHealthCheck]
type greTunnelNewResponseGRETunnelsHealthCheckJSON struct {
	Direction   apijson.Field
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelNewResponseGRETunnelsHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelNewResponseGRETunnelsHealthCheckJSON) RawJSON() string {
	return r.raw
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
// in health_check is ignored as the interface_address is used to send traffic into
// the tunnel.
type GRETunnelNewResponseGRETunnelsHealthCheckDirection string

const (
	GRETunnelNewResponseGRETunnelsHealthCheckDirectionUnidirectional GRETunnelNewResponseGRETunnelsHealthCheckDirection = "unidirectional"
	GRETunnelNewResponseGRETunnelsHealthCheckDirectionBidirectional  GRETunnelNewResponseGRETunnelsHealthCheckDirection = "bidirectional"
)

// How frequent the health check is run. The default value is `mid`.
type GRETunnelNewResponseGRETunnelsHealthCheckRate string

const (
	GRETunnelNewResponseGRETunnelsHealthCheckRateLow  GRETunnelNewResponseGRETunnelsHealthCheckRate = "low"
	GRETunnelNewResponseGRETunnelsHealthCheckRateMid  GRETunnelNewResponseGRETunnelsHealthCheckRate = "mid"
	GRETunnelNewResponseGRETunnelsHealthCheckRateHigh GRETunnelNewResponseGRETunnelsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type GRETunnelNewResponseGRETunnelsHealthCheckType string

const (
	GRETunnelNewResponseGRETunnelsHealthCheckTypeReply   GRETunnelNewResponseGRETunnelsHealthCheckType = "reply"
	GRETunnelNewResponseGRETunnelsHealthCheckTypeRequest GRETunnelNewResponseGRETunnelsHealthCheckType = "request"
)

type GRETunnelUpdateResponse struct {
	Modified          bool                        `json:"modified"`
	ModifiedGRETunnel interface{}                 `json:"modified_gre_tunnel"`
	JSON              greTunnelUpdateResponseJSON `json:"-"`
}

// greTunnelUpdateResponseJSON contains the JSON metadata for the struct
// [GRETunnelUpdateResponse]
type greTunnelUpdateResponseJSON struct {
	Modified          apijson.Field
	ModifiedGRETunnel apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *GRETunnelUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type GRETunnelListResponse struct {
	GRETunnels []GRETunnelListResponseGRETunnel `json:"gre_tunnels"`
	JSON       greTunnelListResponseJSON        `json:"-"`
}

// greTunnelListResponseJSON contains the JSON metadata for the struct
// [GRETunnelListResponse]
type greTunnelListResponseJSON struct {
	GRETunnels  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelListResponseJSON) RawJSON() string {
	return r.raw
}

type GRETunnelListResponseGRETunnel struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel.
	CloudflareGREEndpoint string `json:"cloudflare_gre_endpoint,required"`
	// The IP address assigned to the customer side of the GRE tunnel.
	CustomerGREEndpoint string `json:"customer_gre_endpoint,required"`
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
	Description string                                     `json:"description"`
	HealthCheck GRETunnelListResponseGRETunnelsHealthCheck `json:"health_check"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value
	// is 576.
	Mtu int64 `json:"mtu"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	TTL  int64                              `json:"ttl"`
	JSON greTunnelListResponseGRETunnelJSON `json:"-"`
}

// greTunnelListResponseGRETunnelJSON contains the JSON metadata for the struct
// [GRETunnelListResponseGRETunnel]
type greTunnelListResponseGRETunnelJSON struct {
	CloudflareGREEndpoint apijson.Field
	CustomerGREEndpoint   apijson.Field
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

func (r *GRETunnelListResponseGRETunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelListResponseGRETunnelJSON) RawJSON() string {
	return r.raw
}

type GRETunnelListResponseGRETunnelsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
	// in health_check is ignored as the interface_address is used to send traffic into
	// the tunnel.
	Direction GRETunnelListResponseGRETunnelsHealthCheckDirection `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate GRETunnelListResponseGRETunnelsHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type GRETunnelListResponseGRETunnelsHealthCheckType `json:"type"`
	JSON greTunnelListResponseGRETunnelsHealthCheckJSON `json:"-"`
}

// greTunnelListResponseGRETunnelsHealthCheckJSON contains the JSON metadata for
// the struct [GRETunnelListResponseGRETunnelsHealthCheck]
type greTunnelListResponseGRETunnelsHealthCheckJSON struct {
	Direction   apijson.Field
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelListResponseGRETunnelsHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelListResponseGRETunnelsHealthCheckJSON) RawJSON() string {
	return r.raw
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
// in health_check is ignored as the interface_address is used to send traffic into
// the tunnel.
type GRETunnelListResponseGRETunnelsHealthCheckDirection string

const (
	GRETunnelListResponseGRETunnelsHealthCheckDirectionUnidirectional GRETunnelListResponseGRETunnelsHealthCheckDirection = "unidirectional"
	GRETunnelListResponseGRETunnelsHealthCheckDirectionBidirectional  GRETunnelListResponseGRETunnelsHealthCheckDirection = "bidirectional"
)

// How frequent the health check is run. The default value is `mid`.
type GRETunnelListResponseGRETunnelsHealthCheckRate string

const (
	GRETunnelListResponseGRETunnelsHealthCheckRateLow  GRETunnelListResponseGRETunnelsHealthCheckRate = "low"
	GRETunnelListResponseGRETunnelsHealthCheckRateMid  GRETunnelListResponseGRETunnelsHealthCheckRate = "mid"
	GRETunnelListResponseGRETunnelsHealthCheckRateHigh GRETunnelListResponseGRETunnelsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type GRETunnelListResponseGRETunnelsHealthCheckType string

const (
	GRETunnelListResponseGRETunnelsHealthCheckTypeReply   GRETunnelListResponseGRETunnelsHealthCheckType = "reply"
	GRETunnelListResponseGRETunnelsHealthCheckTypeRequest GRETunnelListResponseGRETunnelsHealthCheckType = "request"
)

type GRETunnelDeleteResponse struct {
	Deleted          bool                        `json:"deleted"`
	DeletedGRETunnel interface{}                 `json:"deleted_gre_tunnel"`
	JSON             greTunnelDeleteResponseJSON `json:"-"`
}

// greTunnelDeleteResponseJSON contains the JSON metadata for the struct
// [GRETunnelDeleteResponse]
type greTunnelDeleteResponseJSON struct {
	Deleted          apijson.Field
	DeletedGRETunnel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *GRETunnelDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type GRETunnelGetResponse struct {
	GRETunnel interface{}              `json:"gre_tunnel"`
	JSON      greTunnelGetResponseJSON `json:"-"`
}

// greTunnelGetResponseJSON contains the JSON metadata for the struct
// [GRETunnelGetResponse]
type greTunnelGetResponseJSON struct {
	GRETunnel   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelGetResponseJSON) RawJSON() string {
	return r.raw
}

type GRETunnelNewParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r GRETunnelNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type GRETunnelNewResponseEnvelope struct {
	Errors   []GRETunnelNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GRETunnelNewResponseEnvelopeMessages `json:"messages,required"`
	Result   GRETunnelNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GRETunnelNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    greTunnelNewResponseEnvelopeJSON    `json:"-"`
}

// greTunnelNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [GRETunnelNewResponseEnvelope]
type greTunnelNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type GRETunnelNewResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    greTunnelNewResponseEnvelopeErrorsJSON `json:"-"`
}

// greTunnelNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [GRETunnelNewResponseEnvelopeErrors]
type greTunnelNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type GRETunnelNewResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    greTunnelNewResponseEnvelopeMessagesJSON `json:"-"`
}

// greTunnelNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [GRETunnelNewResponseEnvelopeMessages]
type greTunnelNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GRETunnelNewResponseEnvelopeSuccess bool

const (
	GRETunnelNewResponseEnvelopeSuccessTrue GRETunnelNewResponseEnvelopeSuccess = true
)

type GRETunnelUpdateParams struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel.
	CloudflareGREEndpoint param.Field[string] `json:"cloudflare_gre_endpoint,required"`
	// The IP address assigned to the customer side of the GRE tunnel.
	CustomerGREEndpoint param.Field[string] `json:"customer_gre_endpoint,required"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress param.Field[string] `json:"interface_address,required"`
	// The name of the tunnel. The name cannot contain spaces or special characters,
	// must be 15 characters or less, and cannot share a name with another GRE tunnel.
	Name param.Field[string] `json:"name,required"`
	// An optional description of the GRE tunnel.
	Description param.Field[string]                           `json:"description"`
	HealthCheck param.Field[GRETunnelUpdateParamsHealthCheck] `json:"health_check"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value
	// is 576.
	Mtu param.Field[int64] `json:"mtu"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	TTL param.Field[int64] `json:"ttl"`
}

func (r GRETunnelUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GRETunnelUpdateParamsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
	// in health_check is ignored as the interface_address is used to send traffic into
	// the tunnel.
	Direction param.Field[GRETunnelUpdateParamsHealthCheckDirection] `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled param.Field[bool] `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate param.Field[GRETunnelUpdateParamsHealthCheckRate] `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target.
	Target param.Field[string] `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type param.Field[GRETunnelUpdateParamsHealthCheckType] `json:"type"`
}

func (r GRETunnelUpdateParamsHealthCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
// in health_check is ignored as the interface_address is used to send traffic into
// the tunnel.
type GRETunnelUpdateParamsHealthCheckDirection string

const (
	GRETunnelUpdateParamsHealthCheckDirectionUnidirectional GRETunnelUpdateParamsHealthCheckDirection = "unidirectional"
	GRETunnelUpdateParamsHealthCheckDirectionBidirectional  GRETunnelUpdateParamsHealthCheckDirection = "bidirectional"
)

// How frequent the health check is run. The default value is `mid`.
type GRETunnelUpdateParamsHealthCheckRate string

const (
	GRETunnelUpdateParamsHealthCheckRateLow  GRETunnelUpdateParamsHealthCheckRate = "low"
	GRETunnelUpdateParamsHealthCheckRateMid  GRETunnelUpdateParamsHealthCheckRate = "mid"
	GRETunnelUpdateParamsHealthCheckRateHigh GRETunnelUpdateParamsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type GRETunnelUpdateParamsHealthCheckType string

const (
	GRETunnelUpdateParamsHealthCheckTypeReply   GRETunnelUpdateParamsHealthCheckType = "reply"
	GRETunnelUpdateParamsHealthCheckTypeRequest GRETunnelUpdateParamsHealthCheckType = "request"
)

type GRETunnelUpdateResponseEnvelope struct {
	Errors   []GRETunnelUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GRETunnelUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   GRETunnelUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GRETunnelUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    greTunnelUpdateResponseEnvelopeJSON    `json:"-"`
}

// greTunnelUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [GRETunnelUpdateResponseEnvelope]
type greTunnelUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type GRETunnelUpdateResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    greTunnelUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// greTunnelUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [GRETunnelUpdateResponseEnvelopeErrors]
type greTunnelUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type GRETunnelUpdateResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    greTunnelUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// greTunnelUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [GRETunnelUpdateResponseEnvelopeMessages]
type greTunnelUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GRETunnelUpdateResponseEnvelopeSuccess bool

const (
	GRETunnelUpdateResponseEnvelopeSuccessTrue GRETunnelUpdateResponseEnvelopeSuccess = true
)

type GRETunnelListResponseEnvelope struct {
	Errors   []GRETunnelListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GRETunnelListResponseEnvelopeMessages `json:"messages,required"`
	Result   GRETunnelListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GRETunnelListResponseEnvelopeSuccess `json:"success,required"`
	JSON    greTunnelListResponseEnvelopeJSON    `json:"-"`
}

// greTunnelListResponseEnvelopeJSON contains the JSON metadata for the struct
// [GRETunnelListResponseEnvelope]
type greTunnelListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type GRETunnelListResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    greTunnelListResponseEnvelopeErrorsJSON `json:"-"`
}

// greTunnelListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [GRETunnelListResponseEnvelopeErrors]
type greTunnelListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type GRETunnelListResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    greTunnelListResponseEnvelopeMessagesJSON `json:"-"`
}

// greTunnelListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [GRETunnelListResponseEnvelopeMessages]
type greTunnelListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GRETunnelListResponseEnvelopeSuccess bool

const (
	GRETunnelListResponseEnvelopeSuccessTrue GRETunnelListResponseEnvelopeSuccess = true
)

type GRETunnelDeleteResponseEnvelope struct {
	Errors   []GRETunnelDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GRETunnelDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   GRETunnelDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GRETunnelDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    greTunnelDeleteResponseEnvelopeJSON    `json:"-"`
}

// greTunnelDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [GRETunnelDeleteResponseEnvelope]
type greTunnelDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type GRETunnelDeleteResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    greTunnelDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// greTunnelDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [GRETunnelDeleteResponseEnvelopeErrors]
type greTunnelDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type GRETunnelDeleteResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    greTunnelDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// greTunnelDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [GRETunnelDeleteResponseEnvelopeMessages]
type greTunnelDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GRETunnelDeleteResponseEnvelopeSuccess bool

const (
	GRETunnelDeleteResponseEnvelopeSuccessTrue GRETunnelDeleteResponseEnvelopeSuccess = true
)

type GRETunnelGetResponseEnvelope struct {
	Errors   []GRETunnelGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GRETunnelGetResponseEnvelopeMessages `json:"messages,required"`
	Result   GRETunnelGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GRETunnelGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    greTunnelGetResponseEnvelopeJSON    `json:"-"`
}

// greTunnelGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [GRETunnelGetResponseEnvelope]
type greTunnelGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type GRETunnelGetResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    greTunnelGetResponseEnvelopeErrorsJSON `json:"-"`
}

// greTunnelGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [GRETunnelGetResponseEnvelopeErrors]
type greTunnelGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type GRETunnelGetResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    greTunnelGetResponseEnvelopeMessagesJSON `json:"-"`
}

// greTunnelGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [GRETunnelGetResponseEnvelopeMessages]
type greTunnelGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GRETunnelGetResponseEnvelopeSuccess bool

const (
	GRETunnelGetResponseEnvelopeSuccessTrue GRETunnelGetResponseEnvelopeSuccess = true
)
