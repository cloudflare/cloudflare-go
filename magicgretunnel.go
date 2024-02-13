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

// Creates new GRE tunnels. Use `?validate_only=true` as an optional query
// parameter to only run validation without persisting changes.
func (r *MagicGreTunnelService) MagicGreTunnelsNewGreTunnels(ctx context.Context, accountIdentifier string, body MagicGreTunnelMagicGreTunnelsNewGreTunnelsParams, opts ...option.RequestOption) (res *MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists GRE tunnels associated with an account.
func (r *MagicGreTunnelService) MagicGreTunnelsListGreTunnels(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *MagicGreTunnelMagicGreTunnelsListGreTunnelsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates multiple GRE tunnels. Use `?validate_only=true` as an optional query
// parameter to only run validation without persisting changes.
func (r *MagicGreTunnelService) MagicGreTunnelsUpdateMultipleGreTunnels(ctx context.Context, accountIdentifier string, body MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsParams, opts ...option.RequestOption) (res *MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

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

type MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponse struct {
	GreTunnels []MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnel `json:"gre_tunnels"`
	JSON       magicGreTunnelMagicGreTunnelsNewGreTunnelsResponseJSON        `json:"-"`
}

// magicGreTunnelMagicGreTunnelsNewGreTunnelsResponseJSON contains the JSON
// metadata for the struct [MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponse]
type magicGreTunnelMagicGreTunnelsNewGreTunnelsResponseJSON struct {
	GreTunnels  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnel struct {
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
	Description string                                                                  `json:"description"`
	HealthCheck MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnelsHealthCheck `json:"health_check"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value
	// is 576.
	Mtu int64 `json:"mtu"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	TTL  int64                                                           `json:"ttl"`
	JSON magicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnelJSON `json:"-"`
}

// magicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnelJSON contains the
// JSON metadata for the struct
// [MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnel]
type magicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnelJSON struct {
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

func (r *MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnelsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel.
	Direction MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnelsHealthCheckDirection `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnelsHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnelsHealthCheckType `json:"type"`
	JSON magicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnelsHealthCheckJSON `json:"-"`
}

// magicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnelsHealthCheckJSON
// contains the JSON metadata for the struct
// [MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnelsHealthCheck]
type magicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnelsHealthCheckJSON struct {
	Direction   apijson.Field
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnelsHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel.
type MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnelsHealthCheckDirection string

const (
	MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnelsHealthCheckDirectionUnidirectional MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnelsHealthCheckDirection = "unidirectional"
	MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnelsHealthCheckDirectionBidirectional  MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnelsHealthCheckDirection = "bidirectional"
)

// How frequent the health check is run. The default value is `mid`.
type MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnelsHealthCheckRate string

const (
	MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnelsHealthCheckRateLow  MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnelsHealthCheckRate = "low"
	MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnelsHealthCheckRateMid  MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnelsHealthCheckRate = "mid"
	MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnelsHealthCheckRateHigh MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnelsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnelsHealthCheckType string

const (
	MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnelsHealthCheckTypeReply   MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnelsHealthCheckType = "reply"
	MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnelsHealthCheckTypeRequest MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseGreTunnelsHealthCheckType = "request"
)

type MagicGreTunnelMagicGreTunnelsListGreTunnelsResponse struct {
	GreTunnels []MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnel `json:"gre_tunnels"`
	JSON       magicGreTunnelMagicGreTunnelsListGreTunnelsResponseJSON        `json:"-"`
}

// magicGreTunnelMagicGreTunnelsListGreTunnelsResponseJSON contains the JSON
// metadata for the struct [MagicGreTunnelMagicGreTunnelsListGreTunnelsResponse]
type magicGreTunnelMagicGreTunnelsListGreTunnelsResponseJSON struct {
	GreTunnels  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelMagicGreTunnelsListGreTunnelsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnel struct {
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
	Description string                                                                   `json:"description"`
	HealthCheck MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnelsHealthCheck `json:"health_check"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value
	// is 576.
	Mtu int64 `json:"mtu"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	TTL  int64                                                            `json:"ttl"`
	JSON magicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnelJSON `json:"-"`
}

// magicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnelJSON contains the
// JSON metadata for the struct
// [MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnel]
type magicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnelJSON struct {
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

func (r *MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnelsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel.
	Direction MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnelsHealthCheckDirection `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnelsHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnelsHealthCheckType `json:"type"`
	JSON magicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnelsHealthCheckJSON `json:"-"`
}

// magicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnelsHealthCheckJSON
// contains the JSON metadata for the struct
// [MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnelsHealthCheck]
type magicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnelsHealthCheckJSON struct {
	Direction   apijson.Field
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnelsHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel.
type MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnelsHealthCheckDirection string

const (
	MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnelsHealthCheckDirectionUnidirectional MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnelsHealthCheckDirection = "unidirectional"
	MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnelsHealthCheckDirectionBidirectional  MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnelsHealthCheckDirection = "bidirectional"
)

// How frequent the health check is run. The default value is `mid`.
type MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnelsHealthCheckRate string

const (
	MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnelsHealthCheckRateLow  MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnelsHealthCheckRate = "low"
	MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnelsHealthCheckRateMid  MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnelsHealthCheckRate = "mid"
	MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnelsHealthCheckRateHigh MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnelsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnelsHealthCheckType string

const (
	MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnelsHealthCheckTypeReply   MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnelsHealthCheckType = "reply"
	MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnelsHealthCheckTypeRequest MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseGreTunnelsHealthCheckType = "request"
)

type MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponse struct {
	Modified           bool                                                                             `json:"modified"`
	ModifiedGreTunnels []MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnel `json:"modified_gre_tunnels"`
	JSON               magicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseJSON                `json:"-"`
}

// magicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseJSON contains the
// JSON metadata for the struct
// [MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponse]
type magicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseJSON struct {
	Modified           apijson.Field
	ModifiedGreTunnels apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnel struct {
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
	Description string                                                                                     `json:"description"`
	HealthCheck MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnelsHealthCheck `json:"health_check"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value
	// is 576.
	Mtu int64 `json:"mtu"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	TTL  int64                                                                              `json:"ttl"`
	JSON magicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnelJSON `json:"-"`
}

// magicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnelJSON
// contains the JSON metadata for the struct
// [MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnel]
type magicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnelJSON struct {
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

func (r *MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnelsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel.
	Direction MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnelsHealthCheckDirection `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnelsHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnelsHealthCheckType `json:"type"`
	JSON magicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnelsHealthCheckJSON `json:"-"`
}

// magicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnelsHealthCheckJSON
// contains the JSON metadata for the struct
// [MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnelsHealthCheck]
type magicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnelsHealthCheckJSON struct {
	Direction   apijson.Field
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnelsHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel.
type MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnelsHealthCheckDirection string

const (
	MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnelsHealthCheckDirectionUnidirectional MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnelsHealthCheckDirection = "unidirectional"
	MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnelsHealthCheckDirectionBidirectional  MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnelsHealthCheckDirection = "bidirectional"
)

// How frequent the health check is run. The default value is `mid`.
type MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnelsHealthCheckRate string

const (
	MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnelsHealthCheckRateLow  MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnelsHealthCheckRate = "low"
	MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnelsHealthCheckRateMid  MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnelsHealthCheckRate = "mid"
	MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnelsHealthCheckRateHigh MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnelsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnelsHealthCheckType string

const (
	MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnelsHealthCheckTypeReply   MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnelsHealthCheckType = "reply"
	MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnelsHealthCheckTypeRequest MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseModifiedGreTunnelsHealthCheckType = "request"
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

type MagicGreTunnelMagicGreTunnelsNewGreTunnelsParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r MagicGreTunnelMagicGreTunnelsNewGreTunnelsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseEnvelope struct {
	Errors   []MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicGreTunnelMagicGreTunnelsNewGreTunnelsResponseEnvelopeJSON    `json:"-"`
}

// magicGreTunnelMagicGreTunnelsNewGreTunnelsResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseEnvelope]
type magicGreTunnelMagicGreTunnelsNewGreTunnelsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseEnvelopeErrors struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    magicGreTunnelMagicGreTunnelsNewGreTunnelsResponseEnvelopeErrorsJSON `json:"-"`
}

// magicGreTunnelMagicGreTunnelsNewGreTunnelsResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseEnvelopeErrors]
type magicGreTunnelMagicGreTunnelsNewGreTunnelsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseEnvelopeMessages struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    magicGreTunnelMagicGreTunnelsNewGreTunnelsResponseEnvelopeMessagesJSON `json:"-"`
}

// magicGreTunnelMagicGreTunnelsNewGreTunnelsResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseEnvelopeMessages]
type magicGreTunnelMagicGreTunnelsNewGreTunnelsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseEnvelopeSuccess bool

const (
	MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseEnvelopeSuccessTrue MagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseEnvelopeSuccess = true
)

type MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseEnvelope struct {
	Errors   []MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicGreTunnelMagicGreTunnelsListGreTunnelsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicGreTunnelMagicGreTunnelsListGreTunnelsResponseEnvelopeJSON    `json:"-"`
}

// magicGreTunnelMagicGreTunnelsListGreTunnelsResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseEnvelope]
type magicGreTunnelMagicGreTunnelsListGreTunnelsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseEnvelopeErrors struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    magicGreTunnelMagicGreTunnelsListGreTunnelsResponseEnvelopeErrorsJSON `json:"-"`
}

// magicGreTunnelMagicGreTunnelsListGreTunnelsResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseEnvelopeErrors]
type magicGreTunnelMagicGreTunnelsListGreTunnelsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseEnvelopeMessages struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    magicGreTunnelMagicGreTunnelsListGreTunnelsResponseEnvelopeMessagesJSON `json:"-"`
}

// magicGreTunnelMagicGreTunnelsListGreTunnelsResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseEnvelopeMessages]
type magicGreTunnelMagicGreTunnelsListGreTunnelsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseEnvelopeSuccess bool

const (
	MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseEnvelopeSuccessTrue MagicGreTunnelMagicGreTunnelsListGreTunnelsResponseEnvelopeSuccess = true
)

type MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseEnvelope struct {
	Errors   []MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseEnvelopeJSON    `json:"-"`
}

// magicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseEnvelope]
type magicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseEnvelopeErrors struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    magicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseEnvelopeErrorsJSON `json:"-"`
}

// magicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseEnvelopeErrors]
type magicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseEnvelopeMessages struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    magicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseEnvelopeMessagesJSON `json:"-"`
}

// magicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseEnvelopeMessages]
type magicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseEnvelopeSuccess bool

const (
	MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseEnvelopeSuccessTrue MagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseEnvelopeSuccess = true
)
