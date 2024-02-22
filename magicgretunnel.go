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

// MagicGRETunnelService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewMagicGRETunnelService] method
// instead.
type MagicGRETunnelService struct {
	Options []option.RequestOption
}

// NewMagicGRETunnelService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMagicGRETunnelService(opts ...option.RequestOption) (r *MagicGRETunnelService) {
	r = &MagicGRETunnelService{}
	r.Options = opts
	return
}

// Creates new GRE tunnels. Use `?validate_only=true` as an optional query
// parameter to only run validation without persisting changes.
func (r *MagicGRETunnelService) New(ctx context.Context, accountIdentifier string, body MagicGRETunnelNewParams, opts ...option.RequestOption) (res *MagicGRETunnelNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicGRETunnelNewResponseEnvelope
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
func (r *MagicGRETunnelService) Update(ctx context.Context, accountIdentifier string, tunnelIdentifier string, body MagicGRETunnelUpdateParams, opts ...option.RequestOption) (res *MagicGRETunnelUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicGRETunnelUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists GRE tunnels associated with an account.
func (r *MagicGRETunnelService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *MagicGRETunnelListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicGRETunnelListResponseEnvelope
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
func (r *MagicGRETunnelService) Delete(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *MagicGRETunnelDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicGRETunnelDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists informtion for a specific GRE tunnel.
func (r *MagicGRETunnelService) Get(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *MagicGRETunnelGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicGRETunnelGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MagicGRETunnelNewResponse struct {
	GRETunnels []MagicGRETunnelNewResponseGRETunnel `json:"gre_tunnels"`
	JSON       magicGRETunnelNewResponseJSON        `json:"-"`
}

// magicGRETunnelNewResponseJSON contains the JSON metadata for the struct
// [MagicGRETunnelNewResponse]
type magicGRETunnelNewResponseJSON struct {
	GRETunnels  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGRETunnelNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGRETunnelNewResponseGRETunnel struct {
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
	Description string                                         `json:"description"`
	HealthCheck MagicGRETunnelNewResponseGRETunnelsHealthCheck `json:"health_check"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value
	// is 576.
	Mtu int64 `json:"mtu"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	TTL  int64                                  `json:"ttl"`
	JSON magicGRETunnelNewResponseGRETunnelJSON `json:"-"`
}

// magicGRETunnelNewResponseGRETunnelJSON contains the JSON metadata for the struct
// [MagicGRETunnelNewResponseGRETunnel]
type magicGRETunnelNewResponseGRETunnelJSON struct {
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

func (r *MagicGRETunnelNewResponseGRETunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGRETunnelNewResponseGRETunnelsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
	// in health_check is ignored as the interface_address is used to send traffic into
	// the tunnel.
	Direction MagicGRETunnelNewResponseGRETunnelsHealthCheckDirection `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate MagicGRETunnelNewResponseGRETunnelsHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type MagicGRETunnelNewResponseGRETunnelsHealthCheckType `json:"type"`
	JSON magicGRETunnelNewResponseGRETunnelsHealthCheckJSON `json:"-"`
}

// magicGRETunnelNewResponseGRETunnelsHealthCheckJSON contains the JSON metadata
// for the struct [MagicGRETunnelNewResponseGRETunnelsHealthCheck]
type magicGRETunnelNewResponseGRETunnelsHealthCheckJSON struct {
	Direction   apijson.Field
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGRETunnelNewResponseGRETunnelsHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
// in health_check is ignored as the interface_address is used to send traffic into
// the tunnel.
type MagicGRETunnelNewResponseGRETunnelsHealthCheckDirection string

const (
	MagicGRETunnelNewResponseGRETunnelsHealthCheckDirectionUnidirectional MagicGRETunnelNewResponseGRETunnelsHealthCheckDirection = "unidirectional"
	MagicGRETunnelNewResponseGRETunnelsHealthCheckDirectionBidirectional  MagicGRETunnelNewResponseGRETunnelsHealthCheckDirection = "bidirectional"
)

// How frequent the health check is run. The default value is `mid`.
type MagicGRETunnelNewResponseGRETunnelsHealthCheckRate string

const (
	MagicGRETunnelNewResponseGRETunnelsHealthCheckRateLow  MagicGRETunnelNewResponseGRETunnelsHealthCheckRate = "low"
	MagicGRETunnelNewResponseGRETunnelsHealthCheckRateMid  MagicGRETunnelNewResponseGRETunnelsHealthCheckRate = "mid"
	MagicGRETunnelNewResponseGRETunnelsHealthCheckRateHigh MagicGRETunnelNewResponseGRETunnelsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicGRETunnelNewResponseGRETunnelsHealthCheckType string

const (
	MagicGRETunnelNewResponseGRETunnelsHealthCheckTypeReply   MagicGRETunnelNewResponseGRETunnelsHealthCheckType = "reply"
	MagicGRETunnelNewResponseGRETunnelsHealthCheckTypeRequest MagicGRETunnelNewResponseGRETunnelsHealthCheckType = "request"
)

type MagicGRETunnelUpdateResponse struct {
	Modified          bool                             `json:"modified"`
	ModifiedGRETunnel interface{}                      `json:"modified_gre_tunnel"`
	JSON              magicGRETunnelUpdateResponseJSON `json:"-"`
}

// magicGRETunnelUpdateResponseJSON contains the JSON metadata for the struct
// [MagicGRETunnelUpdateResponse]
type magicGRETunnelUpdateResponseJSON struct {
	Modified          apijson.Field
	ModifiedGRETunnel apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *MagicGRETunnelUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGRETunnelListResponse struct {
	GRETunnels []MagicGRETunnelListResponseGRETunnel `json:"gre_tunnels"`
	JSON       magicGRETunnelListResponseJSON        `json:"-"`
}

// magicGRETunnelListResponseJSON contains the JSON metadata for the struct
// [MagicGRETunnelListResponse]
type magicGRETunnelListResponseJSON struct {
	GRETunnels  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGRETunnelListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGRETunnelListResponseGRETunnel struct {
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
	Description string                                          `json:"description"`
	HealthCheck MagicGRETunnelListResponseGRETunnelsHealthCheck `json:"health_check"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value
	// is 576.
	Mtu int64 `json:"mtu"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	TTL  int64                                   `json:"ttl"`
	JSON magicGRETunnelListResponseGRETunnelJSON `json:"-"`
}

// magicGRETunnelListResponseGRETunnelJSON contains the JSON metadata for the
// struct [MagicGRETunnelListResponseGRETunnel]
type magicGRETunnelListResponseGRETunnelJSON struct {
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

func (r *MagicGRETunnelListResponseGRETunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGRETunnelListResponseGRETunnelsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
	// in health_check is ignored as the interface_address is used to send traffic into
	// the tunnel.
	Direction MagicGRETunnelListResponseGRETunnelsHealthCheckDirection `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate MagicGRETunnelListResponseGRETunnelsHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type MagicGRETunnelListResponseGRETunnelsHealthCheckType `json:"type"`
	JSON magicGRETunnelListResponseGRETunnelsHealthCheckJSON `json:"-"`
}

// magicGRETunnelListResponseGRETunnelsHealthCheckJSON contains the JSON metadata
// for the struct [MagicGRETunnelListResponseGRETunnelsHealthCheck]
type magicGRETunnelListResponseGRETunnelsHealthCheckJSON struct {
	Direction   apijson.Field
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGRETunnelListResponseGRETunnelsHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
// in health_check is ignored as the interface_address is used to send traffic into
// the tunnel.
type MagicGRETunnelListResponseGRETunnelsHealthCheckDirection string

const (
	MagicGRETunnelListResponseGRETunnelsHealthCheckDirectionUnidirectional MagicGRETunnelListResponseGRETunnelsHealthCheckDirection = "unidirectional"
	MagicGRETunnelListResponseGRETunnelsHealthCheckDirectionBidirectional  MagicGRETunnelListResponseGRETunnelsHealthCheckDirection = "bidirectional"
)

// How frequent the health check is run. The default value is `mid`.
type MagicGRETunnelListResponseGRETunnelsHealthCheckRate string

const (
	MagicGRETunnelListResponseGRETunnelsHealthCheckRateLow  MagicGRETunnelListResponseGRETunnelsHealthCheckRate = "low"
	MagicGRETunnelListResponseGRETunnelsHealthCheckRateMid  MagicGRETunnelListResponseGRETunnelsHealthCheckRate = "mid"
	MagicGRETunnelListResponseGRETunnelsHealthCheckRateHigh MagicGRETunnelListResponseGRETunnelsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicGRETunnelListResponseGRETunnelsHealthCheckType string

const (
	MagicGRETunnelListResponseGRETunnelsHealthCheckTypeReply   MagicGRETunnelListResponseGRETunnelsHealthCheckType = "reply"
	MagicGRETunnelListResponseGRETunnelsHealthCheckTypeRequest MagicGRETunnelListResponseGRETunnelsHealthCheckType = "request"
)

type MagicGRETunnelDeleteResponse struct {
	Deleted          bool                             `json:"deleted"`
	DeletedGRETunnel interface{}                      `json:"deleted_gre_tunnel"`
	JSON             magicGRETunnelDeleteResponseJSON `json:"-"`
}

// magicGRETunnelDeleteResponseJSON contains the JSON metadata for the struct
// [MagicGRETunnelDeleteResponse]
type magicGRETunnelDeleteResponseJSON struct {
	Deleted          apijson.Field
	DeletedGRETunnel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MagicGRETunnelDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGRETunnelGetResponse struct {
	GRETunnel interface{}                   `json:"gre_tunnel"`
	JSON      magicGRETunnelGetResponseJSON `json:"-"`
}

// magicGRETunnelGetResponseJSON contains the JSON metadata for the struct
// [MagicGRETunnelGetResponse]
type magicGRETunnelGetResponseJSON struct {
	GRETunnel   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGRETunnelGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGRETunnelNewParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r MagicGRETunnelNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type MagicGRETunnelNewResponseEnvelope struct {
	Errors   []MagicGRETunnelNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicGRETunnelNewResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicGRETunnelNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicGRETunnelNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicGRETunnelNewResponseEnvelopeJSON    `json:"-"`
}

// magicGRETunnelNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [MagicGRETunnelNewResponseEnvelope]
type magicGRETunnelNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGRETunnelNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGRETunnelNewResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    magicGRETunnelNewResponseEnvelopeErrorsJSON `json:"-"`
}

// magicGRETunnelNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MagicGRETunnelNewResponseEnvelopeErrors]
type magicGRETunnelNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGRETunnelNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGRETunnelNewResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    magicGRETunnelNewResponseEnvelopeMessagesJSON `json:"-"`
}

// magicGRETunnelNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MagicGRETunnelNewResponseEnvelopeMessages]
type magicGRETunnelNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGRETunnelNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicGRETunnelNewResponseEnvelopeSuccess bool

const (
	MagicGRETunnelNewResponseEnvelopeSuccessTrue MagicGRETunnelNewResponseEnvelopeSuccess = true
)

type MagicGRETunnelUpdateParams struct {
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
	Description param.Field[string]                                `json:"description"`
	HealthCheck param.Field[MagicGRETunnelUpdateParamsHealthCheck] `json:"health_check"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value
	// is 576.
	Mtu param.Field[int64] `json:"mtu"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	TTL param.Field[int64] `json:"ttl"`
}

func (r MagicGRETunnelUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagicGRETunnelUpdateParamsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
	// in health_check is ignored as the interface_address is used to send traffic into
	// the tunnel.
	Direction param.Field[MagicGRETunnelUpdateParamsHealthCheckDirection] `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled param.Field[bool] `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate param.Field[MagicGRETunnelUpdateParamsHealthCheckRate] `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target.
	Target param.Field[string] `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type param.Field[MagicGRETunnelUpdateParamsHealthCheckType] `json:"type"`
}

func (r MagicGRETunnelUpdateParamsHealthCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
// in health_check is ignored as the interface_address is used to send traffic into
// the tunnel.
type MagicGRETunnelUpdateParamsHealthCheckDirection string

const (
	MagicGRETunnelUpdateParamsHealthCheckDirectionUnidirectional MagicGRETunnelUpdateParamsHealthCheckDirection = "unidirectional"
	MagicGRETunnelUpdateParamsHealthCheckDirectionBidirectional  MagicGRETunnelUpdateParamsHealthCheckDirection = "bidirectional"
)

// How frequent the health check is run. The default value is `mid`.
type MagicGRETunnelUpdateParamsHealthCheckRate string

const (
	MagicGRETunnelUpdateParamsHealthCheckRateLow  MagicGRETunnelUpdateParamsHealthCheckRate = "low"
	MagicGRETunnelUpdateParamsHealthCheckRateMid  MagicGRETunnelUpdateParamsHealthCheckRate = "mid"
	MagicGRETunnelUpdateParamsHealthCheckRateHigh MagicGRETunnelUpdateParamsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicGRETunnelUpdateParamsHealthCheckType string

const (
	MagicGRETunnelUpdateParamsHealthCheckTypeReply   MagicGRETunnelUpdateParamsHealthCheckType = "reply"
	MagicGRETunnelUpdateParamsHealthCheckTypeRequest MagicGRETunnelUpdateParamsHealthCheckType = "request"
)

type MagicGRETunnelUpdateResponseEnvelope struct {
	Errors   []MagicGRETunnelUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicGRETunnelUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicGRETunnelUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicGRETunnelUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicGRETunnelUpdateResponseEnvelopeJSON    `json:"-"`
}

// magicGRETunnelUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [MagicGRETunnelUpdateResponseEnvelope]
type magicGRETunnelUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGRETunnelUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGRETunnelUpdateResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    magicGRETunnelUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// magicGRETunnelUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [MagicGRETunnelUpdateResponseEnvelopeErrors]
type magicGRETunnelUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGRETunnelUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGRETunnelUpdateResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    magicGRETunnelUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// magicGRETunnelUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [MagicGRETunnelUpdateResponseEnvelopeMessages]
type magicGRETunnelUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGRETunnelUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicGRETunnelUpdateResponseEnvelopeSuccess bool

const (
	MagicGRETunnelUpdateResponseEnvelopeSuccessTrue MagicGRETunnelUpdateResponseEnvelopeSuccess = true
)

type MagicGRETunnelListResponseEnvelope struct {
	Errors   []MagicGRETunnelListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicGRETunnelListResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicGRETunnelListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicGRETunnelListResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicGRETunnelListResponseEnvelopeJSON    `json:"-"`
}

// magicGRETunnelListResponseEnvelopeJSON contains the JSON metadata for the struct
// [MagicGRETunnelListResponseEnvelope]
type magicGRETunnelListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGRETunnelListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGRETunnelListResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    magicGRETunnelListResponseEnvelopeErrorsJSON `json:"-"`
}

// magicGRETunnelListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MagicGRETunnelListResponseEnvelopeErrors]
type magicGRETunnelListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGRETunnelListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGRETunnelListResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    magicGRETunnelListResponseEnvelopeMessagesJSON `json:"-"`
}

// magicGRETunnelListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [MagicGRETunnelListResponseEnvelopeMessages]
type magicGRETunnelListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGRETunnelListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicGRETunnelListResponseEnvelopeSuccess bool

const (
	MagicGRETunnelListResponseEnvelopeSuccessTrue MagicGRETunnelListResponseEnvelopeSuccess = true
)

type MagicGRETunnelDeleteResponseEnvelope struct {
	Errors   []MagicGRETunnelDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicGRETunnelDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicGRETunnelDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicGRETunnelDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicGRETunnelDeleteResponseEnvelopeJSON    `json:"-"`
}

// magicGRETunnelDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [MagicGRETunnelDeleteResponseEnvelope]
type magicGRETunnelDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGRETunnelDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGRETunnelDeleteResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    magicGRETunnelDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// magicGRETunnelDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [MagicGRETunnelDeleteResponseEnvelopeErrors]
type magicGRETunnelDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGRETunnelDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGRETunnelDeleteResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    magicGRETunnelDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// magicGRETunnelDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [MagicGRETunnelDeleteResponseEnvelopeMessages]
type magicGRETunnelDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGRETunnelDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicGRETunnelDeleteResponseEnvelopeSuccess bool

const (
	MagicGRETunnelDeleteResponseEnvelopeSuccessTrue MagicGRETunnelDeleteResponseEnvelopeSuccess = true
)

type MagicGRETunnelGetResponseEnvelope struct {
	Errors   []MagicGRETunnelGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicGRETunnelGetResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicGRETunnelGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicGRETunnelGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicGRETunnelGetResponseEnvelopeJSON    `json:"-"`
}

// magicGRETunnelGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [MagicGRETunnelGetResponseEnvelope]
type magicGRETunnelGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGRETunnelGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGRETunnelGetResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    magicGRETunnelGetResponseEnvelopeErrorsJSON `json:"-"`
}

// magicGRETunnelGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MagicGRETunnelGetResponseEnvelopeErrors]
type magicGRETunnelGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGRETunnelGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGRETunnelGetResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    magicGRETunnelGetResponseEnvelopeMessagesJSON `json:"-"`
}

// magicGRETunnelGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MagicGRETunnelGetResponseEnvelopeMessages]
type magicGRETunnelGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGRETunnelGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicGRETunnelGetResponseEnvelopeSuccess bool

const (
	MagicGRETunnelGetResponseEnvelopeSuccessTrue MagicGRETunnelGetResponseEnvelopeSuccess = true
)
