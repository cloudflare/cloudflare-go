// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// MagicTransitGRETunnelService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewMagicTransitGRETunnelService]
// method instead.
type MagicTransitGRETunnelService struct {
	Options []option.RequestOption
}

// NewMagicTransitGRETunnelService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMagicTransitGRETunnelService(opts ...option.RequestOption) (r *MagicTransitGRETunnelService) {
	r = &MagicTransitGRETunnelService{}
	r.Options = opts
	return
}

// Creates new GRE tunnels. Use `?validate_only=true` as an optional query
// parameter to only run validation without persisting changes.
func (r *MagicTransitGRETunnelService) New(ctx context.Context, accountIdentifier string, body MagicTransitGRETunnelNewParams, opts ...option.RequestOption) (res *MagicTransitGRETunnelNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicTransitGRETunnelNewResponseEnvelope
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
func (r *MagicTransitGRETunnelService) Update(ctx context.Context, accountIdentifier string, tunnelIdentifier string, body MagicTransitGRETunnelUpdateParams, opts ...option.RequestOption) (res *MagicTransitGRETunnelUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicTransitGRETunnelUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists GRE tunnels associated with an account.
func (r *MagicTransitGRETunnelService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *MagicTransitGRETunnelListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicTransitGRETunnelListResponseEnvelope
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
func (r *MagicTransitGRETunnelService) Delete(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *MagicTransitGRETunnelDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicTransitGRETunnelDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists informtion for a specific GRE tunnel.
func (r *MagicTransitGRETunnelService) Get(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *MagicTransitGRETunnelGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicTransitGRETunnelGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MagicTransitGRETunnelNewResponse struct {
	GRETunnels []MagicTransitGRETunnelNewResponseGRETunnel `json:"gre_tunnels"`
	JSON       magicTransitGRETunnelNewResponseJSON        `json:"-"`
}

// magicTransitGRETunnelNewResponseJSON contains the JSON metadata for the struct
// [MagicTransitGRETunnelNewResponse]
type magicTransitGRETunnelNewResponseJSON struct {
	GRETunnels  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitGRETunnelNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitGRETunnelNewResponseJSON) RawJSON() string {
	return r.raw
}

type MagicTransitGRETunnelNewResponseGRETunnel struct {
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
	Description string                                                `json:"description"`
	HealthCheck MagicTransitGRETunnelNewResponseGRETunnelsHealthCheck `json:"health_check"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value
	// is 576.
	Mtu int64 `json:"mtu"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	TTL  int64                                         `json:"ttl"`
	JSON magicTransitGRETunnelNewResponseGRETunnelJSON `json:"-"`
}

// magicTransitGRETunnelNewResponseGRETunnelJSON contains the JSON metadata for the
// struct [MagicTransitGRETunnelNewResponseGRETunnel]
type magicTransitGRETunnelNewResponseGRETunnelJSON struct {
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

func (r *MagicTransitGRETunnelNewResponseGRETunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitGRETunnelNewResponseGRETunnelJSON) RawJSON() string {
	return r.raw
}

type MagicTransitGRETunnelNewResponseGRETunnelsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
	// in health_check is ignored as the interface_address is used to send traffic into
	// the tunnel.
	Direction MagicTransitGRETunnelNewResponseGRETunnelsHealthCheckDirection `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate MagicTransitGRETunnelNewResponseGRETunnelsHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type MagicTransitGRETunnelNewResponseGRETunnelsHealthCheckType `json:"type"`
	JSON magicTransitGRETunnelNewResponseGRETunnelsHealthCheckJSON `json:"-"`
}

// magicTransitGRETunnelNewResponseGRETunnelsHealthCheckJSON contains the JSON
// metadata for the struct [MagicTransitGRETunnelNewResponseGRETunnelsHealthCheck]
type magicTransitGRETunnelNewResponseGRETunnelsHealthCheckJSON struct {
	Direction   apijson.Field
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitGRETunnelNewResponseGRETunnelsHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitGRETunnelNewResponseGRETunnelsHealthCheckJSON) RawJSON() string {
	return r.raw
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
// in health_check is ignored as the interface_address is used to send traffic into
// the tunnel.
type MagicTransitGRETunnelNewResponseGRETunnelsHealthCheckDirection string

const (
	MagicTransitGRETunnelNewResponseGRETunnelsHealthCheckDirectionUnidirectional MagicTransitGRETunnelNewResponseGRETunnelsHealthCheckDirection = "unidirectional"
	MagicTransitGRETunnelNewResponseGRETunnelsHealthCheckDirectionBidirectional  MagicTransitGRETunnelNewResponseGRETunnelsHealthCheckDirection = "bidirectional"
)

// How frequent the health check is run. The default value is `mid`.
type MagicTransitGRETunnelNewResponseGRETunnelsHealthCheckRate string

const (
	MagicTransitGRETunnelNewResponseGRETunnelsHealthCheckRateLow  MagicTransitGRETunnelNewResponseGRETunnelsHealthCheckRate = "low"
	MagicTransitGRETunnelNewResponseGRETunnelsHealthCheckRateMid  MagicTransitGRETunnelNewResponseGRETunnelsHealthCheckRate = "mid"
	MagicTransitGRETunnelNewResponseGRETunnelsHealthCheckRateHigh MagicTransitGRETunnelNewResponseGRETunnelsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicTransitGRETunnelNewResponseGRETunnelsHealthCheckType string

const (
	MagicTransitGRETunnelNewResponseGRETunnelsHealthCheckTypeReply   MagicTransitGRETunnelNewResponseGRETunnelsHealthCheckType = "reply"
	MagicTransitGRETunnelNewResponseGRETunnelsHealthCheckTypeRequest MagicTransitGRETunnelNewResponseGRETunnelsHealthCheckType = "request"
)

type MagicTransitGRETunnelUpdateResponse struct {
	Modified          bool                                    `json:"modified"`
	ModifiedGRETunnel interface{}                             `json:"modified_gre_tunnel"`
	JSON              magicTransitGRETunnelUpdateResponseJSON `json:"-"`
}

// magicTransitGRETunnelUpdateResponseJSON contains the JSON metadata for the
// struct [MagicTransitGRETunnelUpdateResponse]
type magicTransitGRETunnelUpdateResponseJSON struct {
	Modified          apijson.Field
	ModifiedGRETunnel apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *MagicTransitGRETunnelUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitGRETunnelUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type MagicTransitGRETunnelListResponse struct {
	GRETunnels []MagicTransitGRETunnelListResponseGRETunnel `json:"gre_tunnels"`
	JSON       magicTransitGRETunnelListResponseJSON        `json:"-"`
}

// magicTransitGRETunnelListResponseJSON contains the JSON metadata for the struct
// [MagicTransitGRETunnelListResponse]
type magicTransitGRETunnelListResponseJSON struct {
	GRETunnels  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitGRETunnelListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitGRETunnelListResponseJSON) RawJSON() string {
	return r.raw
}

type MagicTransitGRETunnelListResponseGRETunnel struct {
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
	Description string                                                 `json:"description"`
	HealthCheck MagicTransitGRETunnelListResponseGRETunnelsHealthCheck `json:"health_check"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value
	// is 576.
	Mtu int64 `json:"mtu"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	TTL  int64                                          `json:"ttl"`
	JSON magicTransitGRETunnelListResponseGRETunnelJSON `json:"-"`
}

// magicTransitGRETunnelListResponseGRETunnelJSON contains the JSON metadata for
// the struct [MagicTransitGRETunnelListResponseGRETunnel]
type magicTransitGRETunnelListResponseGRETunnelJSON struct {
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

func (r *MagicTransitGRETunnelListResponseGRETunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitGRETunnelListResponseGRETunnelJSON) RawJSON() string {
	return r.raw
}

type MagicTransitGRETunnelListResponseGRETunnelsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
	// in health_check is ignored as the interface_address is used to send traffic into
	// the tunnel.
	Direction MagicTransitGRETunnelListResponseGRETunnelsHealthCheckDirection `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate MagicTransitGRETunnelListResponseGRETunnelsHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type MagicTransitGRETunnelListResponseGRETunnelsHealthCheckType `json:"type"`
	JSON magicTransitGRETunnelListResponseGRETunnelsHealthCheckJSON `json:"-"`
}

// magicTransitGRETunnelListResponseGRETunnelsHealthCheckJSON contains the JSON
// metadata for the struct [MagicTransitGRETunnelListResponseGRETunnelsHealthCheck]
type magicTransitGRETunnelListResponseGRETunnelsHealthCheckJSON struct {
	Direction   apijson.Field
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitGRETunnelListResponseGRETunnelsHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitGRETunnelListResponseGRETunnelsHealthCheckJSON) RawJSON() string {
	return r.raw
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
// in health_check is ignored as the interface_address is used to send traffic into
// the tunnel.
type MagicTransitGRETunnelListResponseGRETunnelsHealthCheckDirection string

const (
	MagicTransitGRETunnelListResponseGRETunnelsHealthCheckDirectionUnidirectional MagicTransitGRETunnelListResponseGRETunnelsHealthCheckDirection = "unidirectional"
	MagicTransitGRETunnelListResponseGRETunnelsHealthCheckDirectionBidirectional  MagicTransitGRETunnelListResponseGRETunnelsHealthCheckDirection = "bidirectional"
)

// How frequent the health check is run. The default value is `mid`.
type MagicTransitGRETunnelListResponseGRETunnelsHealthCheckRate string

const (
	MagicTransitGRETunnelListResponseGRETunnelsHealthCheckRateLow  MagicTransitGRETunnelListResponseGRETunnelsHealthCheckRate = "low"
	MagicTransitGRETunnelListResponseGRETunnelsHealthCheckRateMid  MagicTransitGRETunnelListResponseGRETunnelsHealthCheckRate = "mid"
	MagicTransitGRETunnelListResponseGRETunnelsHealthCheckRateHigh MagicTransitGRETunnelListResponseGRETunnelsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicTransitGRETunnelListResponseGRETunnelsHealthCheckType string

const (
	MagicTransitGRETunnelListResponseGRETunnelsHealthCheckTypeReply   MagicTransitGRETunnelListResponseGRETunnelsHealthCheckType = "reply"
	MagicTransitGRETunnelListResponseGRETunnelsHealthCheckTypeRequest MagicTransitGRETunnelListResponseGRETunnelsHealthCheckType = "request"
)

type MagicTransitGRETunnelDeleteResponse struct {
	Deleted          bool                                    `json:"deleted"`
	DeletedGRETunnel interface{}                             `json:"deleted_gre_tunnel"`
	JSON             magicTransitGRETunnelDeleteResponseJSON `json:"-"`
}

// magicTransitGRETunnelDeleteResponseJSON contains the JSON metadata for the
// struct [MagicTransitGRETunnelDeleteResponse]
type magicTransitGRETunnelDeleteResponseJSON struct {
	Deleted          apijson.Field
	DeletedGRETunnel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MagicTransitGRETunnelDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitGRETunnelDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type MagicTransitGRETunnelGetResponse struct {
	GRETunnel interface{}                          `json:"gre_tunnel"`
	JSON      magicTransitGRETunnelGetResponseJSON `json:"-"`
}

// magicTransitGRETunnelGetResponseJSON contains the JSON metadata for the struct
// [MagicTransitGRETunnelGetResponse]
type magicTransitGRETunnelGetResponseJSON struct {
	GRETunnel   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitGRETunnelGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitGRETunnelGetResponseJSON) RawJSON() string {
	return r.raw
}

type MagicTransitGRETunnelNewParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r MagicTransitGRETunnelNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type MagicTransitGRETunnelNewResponseEnvelope struct {
	Errors   []MagicTransitGRETunnelNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicTransitGRETunnelNewResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicTransitGRETunnelNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicTransitGRETunnelNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicTransitGRETunnelNewResponseEnvelopeJSON    `json:"-"`
}

// magicTransitGRETunnelNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [MagicTransitGRETunnelNewResponseEnvelope]
type magicTransitGRETunnelNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitGRETunnelNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitGRETunnelNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MagicTransitGRETunnelNewResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    magicTransitGRETunnelNewResponseEnvelopeErrorsJSON `json:"-"`
}

// magicTransitGRETunnelNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [MagicTransitGRETunnelNewResponseEnvelopeErrors]
type magicTransitGRETunnelNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitGRETunnelNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitGRETunnelNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MagicTransitGRETunnelNewResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    magicTransitGRETunnelNewResponseEnvelopeMessagesJSON `json:"-"`
}

// magicTransitGRETunnelNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [MagicTransitGRETunnelNewResponseEnvelopeMessages]
type magicTransitGRETunnelNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitGRETunnelNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitGRETunnelNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicTransitGRETunnelNewResponseEnvelopeSuccess bool

const (
	MagicTransitGRETunnelNewResponseEnvelopeSuccessTrue MagicTransitGRETunnelNewResponseEnvelopeSuccess = true
)

type MagicTransitGRETunnelUpdateParams struct {
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
	Description param.Field[string]                                       `json:"description"`
	HealthCheck param.Field[MagicTransitGRETunnelUpdateParamsHealthCheck] `json:"health_check"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value
	// is 576.
	Mtu param.Field[int64] `json:"mtu"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	TTL param.Field[int64] `json:"ttl"`
}

func (r MagicTransitGRETunnelUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagicTransitGRETunnelUpdateParamsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
	// in health_check is ignored as the interface_address is used to send traffic into
	// the tunnel.
	Direction param.Field[MagicTransitGRETunnelUpdateParamsHealthCheckDirection] `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled param.Field[bool] `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate param.Field[MagicTransitGRETunnelUpdateParamsHealthCheckRate] `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target.
	Target param.Field[string] `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type param.Field[MagicTransitGRETunnelUpdateParamsHealthCheckType] `json:"type"`
}

func (r MagicTransitGRETunnelUpdateParamsHealthCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
// in health_check is ignored as the interface_address is used to send traffic into
// the tunnel.
type MagicTransitGRETunnelUpdateParamsHealthCheckDirection string

const (
	MagicTransitGRETunnelUpdateParamsHealthCheckDirectionUnidirectional MagicTransitGRETunnelUpdateParamsHealthCheckDirection = "unidirectional"
	MagicTransitGRETunnelUpdateParamsHealthCheckDirectionBidirectional  MagicTransitGRETunnelUpdateParamsHealthCheckDirection = "bidirectional"
)

// How frequent the health check is run. The default value is `mid`.
type MagicTransitGRETunnelUpdateParamsHealthCheckRate string

const (
	MagicTransitGRETunnelUpdateParamsHealthCheckRateLow  MagicTransitGRETunnelUpdateParamsHealthCheckRate = "low"
	MagicTransitGRETunnelUpdateParamsHealthCheckRateMid  MagicTransitGRETunnelUpdateParamsHealthCheckRate = "mid"
	MagicTransitGRETunnelUpdateParamsHealthCheckRateHigh MagicTransitGRETunnelUpdateParamsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicTransitGRETunnelUpdateParamsHealthCheckType string

const (
	MagicTransitGRETunnelUpdateParamsHealthCheckTypeReply   MagicTransitGRETunnelUpdateParamsHealthCheckType = "reply"
	MagicTransitGRETunnelUpdateParamsHealthCheckTypeRequest MagicTransitGRETunnelUpdateParamsHealthCheckType = "request"
)

type MagicTransitGRETunnelUpdateResponseEnvelope struct {
	Errors   []MagicTransitGRETunnelUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicTransitGRETunnelUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicTransitGRETunnelUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicTransitGRETunnelUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicTransitGRETunnelUpdateResponseEnvelopeJSON    `json:"-"`
}

// magicTransitGRETunnelUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [MagicTransitGRETunnelUpdateResponseEnvelope]
type magicTransitGRETunnelUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitGRETunnelUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitGRETunnelUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MagicTransitGRETunnelUpdateResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    magicTransitGRETunnelUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// magicTransitGRETunnelUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [MagicTransitGRETunnelUpdateResponseEnvelopeErrors]
type magicTransitGRETunnelUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitGRETunnelUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitGRETunnelUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MagicTransitGRETunnelUpdateResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    magicTransitGRETunnelUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// magicTransitGRETunnelUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [MagicTransitGRETunnelUpdateResponseEnvelopeMessages]
type magicTransitGRETunnelUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitGRETunnelUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitGRETunnelUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicTransitGRETunnelUpdateResponseEnvelopeSuccess bool

const (
	MagicTransitGRETunnelUpdateResponseEnvelopeSuccessTrue MagicTransitGRETunnelUpdateResponseEnvelopeSuccess = true
)

type MagicTransitGRETunnelListResponseEnvelope struct {
	Errors   []MagicTransitGRETunnelListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicTransitGRETunnelListResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicTransitGRETunnelListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicTransitGRETunnelListResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicTransitGRETunnelListResponseEnvelopeJSON    `json:"-"`
}

// magicTransitGRETunnelListResponseEnvelopeJSON contains the JSON metadata for the
// struct [MagicTransitGRETunnelListResponseEnvelope]
type magicTransitGRETunnelListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitGRETunnelListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitGRETunnelListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MagicTransitGRETunnelListResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    magicTransitGRETunnelListResponseEnvelopeErrorsJSON `json:"-"`
}

// magicTransitGRETunnelListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [MagicTransitGRETunnelListResponseEnvelopeErrors]
type magicTransitGRETunnelListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitGRETunnelListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitGRETunnelListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MagicTransitGRETunnelListResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    magicTransitGRETunnelListResponseEnvelopeMessagesJSON `json:"-"`
}

// magicTransitGRETunnelListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [MagicTransitGRETunnelListResponseEnvelopeMessages]
type magicTransitGRETunnelListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitGRETunnelListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitGRETunnelListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicTransitGRETunnelListResponseEnvelopeSuccess bool

const (
	MagicTransitGRETunnelListResponseEnvelopeSuccessTrue MagicTransitGRETunnelListResponseEnvelopeSuccess = true
)

type MagicTransitGRETunnelDeleteResponseEnvelope struct {
	Errors   []MagicTransitGRETunnelDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicTransitGRETunnelDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicTransitGRETunnelDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicTransitGRETunnelDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicTransitGRETunnelDeleteResponseEnvelopeJSON    `json:"-"`
}

// magicTransitGRETunnelDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [MagicTransitGRETunnelDeleteResponseEnvelope]
type magicTransitGRETunnelDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitGRETunnelDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitGRETunnelDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MagicTransitGRETunnelDeleteResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    magicTransitGRETunnelDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// magicTransitGRETunnelDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [MagicTransitGRETunnelDeleteResponseEnvelopeErrors]
type magicTransitGRETunnelDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitGRETunnelDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitGRETunnelDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MagicTransitGRETunnelDeleteResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    magicTransitGRETunnelDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// magicTransitGRETunnelDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [MagicTransitGRETunnelDeleteResponseEnvelopeMessages]
type magicTransitGRETunnelDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitGRETunnelDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitGRETunnelDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicTransitGRETunnelDeleteResponseEnvelopeSuccess bool

const (
	MagicTransitGRETunnelDeleteResponseEnvelopeSuccessTrue MagicTransitGRETunnelDeleteResponseEnvelopeSuccess = true
)

type MagicTransitGRETunnelGetResponseEnvelope struct {
	Errors   []MagicTransitGRETunnelGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicTransitGRETunnelGetResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicTransitGRETunnelGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicTransitGRETunnelGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicTransitGRETunnelGetResponseEnvelopeJSON    `json:"-"`
}

// magicTransitGRETunnelGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [MagicTransitGRETunnelGetResponseEnvelope]
type magicTransitGRETunnelGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitGRETunnelGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitGRETunnelGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MagicTransitGRETunnelGetResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    magicTransitGRETunnelGetResponseEnvelopeErrorsJSON `json:"-"`
}

// magicTransitGRETunnelGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [MagicTransitGRETunnelGetResponseEnvelopeErrors]
type magicTransitGRETunnelGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitGRETunnelGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitGRETunnelGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MagicTransitGRETunnelGetResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    magicTransitGRETunnelGetResponseEnvelopeMessagesJSON `json:"-"`
}

// magicTransitGRETunnelGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [MagicTransitGRETunnelGetResponseEnvelopeMessages]
type magicTransitGRETunnelGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitGRETunnelGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitGRETunnelGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicTransitGRETunnelGetResponseEnvelopeSuccess bool

const (
	MagicTransitGRETunnelGetResponseEnvelopeSuccessTrue MagicTransitGRETunnelGetResponseEnvelopeSuccess = true
)
