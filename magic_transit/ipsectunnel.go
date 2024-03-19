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

// IPSECTunnelService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewIPSECTunnelService] method
// instead.
type IPSECTunnelService struct {
	Options []option.RequestOption
}

// NewIPSECTunnelService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIPSECTunnelService(opts ...option.RequestOption) (r *IPSECTunnelService) {
	r = &IPSECTunnelService{}
	r.Options = opts
	return
}

// Creates new IPsec tunnels associated with an account. Use `?validate_only=true`
// as an optional query parameter to only run validation without persisting
// changes.
func (r *IPSECTunnelService) New(ctx context.Context, accountIdentifier string, body IPSECTunnelNewParams, opts ...option.RequestOption) (res *IPSECTunnelNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IPSECTunnelNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a specific IPsec tunnel associated with an account. Use
// `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes.
func (r *IPSECTunnelService) Update(ctx context.Context, accountIdentifier string, tunnelIdentifier string, body IPSECTunnelUpdateParams, opts ...option.RequestOption) (res *IPSECTunnelUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IPSECTunnelUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists IPsec tunnels associated with an account.
func (r *IPSECTunnelService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *IPSECTunnelListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IPSECTunnelListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Disables and removes a specific static IPsec Tunnel associated with an account.
// Use `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes.
func (r *IPSECTunnelService) Delete(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *IPSECTunnelDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IPSECTunnelDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists details for a specific IPsec tunnel.
func (r *IPSECTunnelService) Get(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *IPSECTunnelGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IPSECTunnelGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Generates a Pre Shared Key for a specific IPsec tunnel used in the IKE session.
// Use `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes. After a PSK is generated, the PSK is immediately
// persisted to Cloudflare's edge and cannot be retrieved later. Note the PSK in a
// safe place.
func (r *IPSECTunnelService) PSKGenerate(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *IPSECTunnelPSKGenerateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IPSECTunnelPSKGenerateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s/psk_generate", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IPSECTunnelNewResponse struct {
	IPSECTunnels []IPSECTunnelNewResponseIPSECTunnel `json:"ipsec_tunnels"`
	JSON         ipsecTunnelNewResponseJSON          `json:"-"`
}

// ipsecTunnelNewResponseJSON contains the JSON metadata for the struct
// [IPSECTunnelNewResponse]
type ipsecTunnelNewResponseJSON struct {
	IPSECTunnels apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IPSECTunnelNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelNewResponseJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelNewResponseIPSECTunnel struct {
	// The IP address assigned to the Cloudflare side of the IPsec tunnel.
	CloudflareEndpoint string `json:"cloudflare_endpoint,required"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress string `json:"interface_address,required"`
	// The name of the IPsec tunnel. The name cannot share a name with other tunnels.
	Name string `json:"name,required"`
	// Tunnel identifier tag.
	ID string `json:"id"`
	// When `true`, the tunnel can use a null-cipher (`ENCR_NULL`) in the ESP tunnel
	// (Phase 2).
	AllowNullCipher bool `json:"allow_null_cipher"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The IP address assigned to the customer side of the IPsec tunnel.
	CustomerEndpoint string `json:"customer_endpoint"`
	// An optional description forthe IPsec tunnel.
	Description string `json:"description"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The PSK metadata that includes when the PSK was generated.
	PSKMetadata IPSECTunnelNewResponseIPSECTunnelsPSKMetadata `json:"psk_metadata"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection  bool                                                `json:"replay_protection"`
	TunnelHealthCheck IPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheck `json:"tunnel_health_check"`
	JSON              ipsecTunnelNewResponseIPSECTunnelJSON               `json:"-"`
}

// ipsecTunnelNewResponseIPSECTunnelJSON contains the JSON metadata for the struct
// [IPSECTunnelNewResponseIPSECTunnel]
type ipsecTunnelNewResponseIPSECTunnelJSON struct {
	CloudflareEndpoint apijson.Field
	InterfaceAddress   apijson.Field
	Name               apijson.Field
	ID                 apijson.Field
	AllowNullCipher    apijson.Field
	CreatedOn          apijson.Field
	CustomerEndpoint   apijson.Field
	Description        apijson.Field
	ModifiedOn         apijson.Field
	PSKMetadata        apijson.Field
	ReplayProtection   apijson.Field
	TunnelHealthCheck  apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *IPSECTunnelNewResponseIPSECTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelNewResponseIPSECTunnelJSON) RawJSON() string {
	return r.raw
}

// The PSK metadata that includes when the PSK was generated.
type IPSECTunnelNewResponseIPSECTunnelsPSKMetadata struct {
	// The date and time the tunnel was last modified.
	LastGeneratedOn time.Time                                         `json:"last_generated_on" format:"date-time"`
	JSON            ipsecTunnelNewResponseIPSECTunnelsPSKMetadataJSON `json:"-"`
}

// ipsecTunnelNewResponseIPSECTunnelsPSKMetadataJSON contains the JSON metadata for
// the struct [IPSECTunnelNewResponseIPSECTunnelsPSKMetadata]
type ipsecTunnelNewResponseIPSECTunnelsPSKMetadataJSON struct {
	LastGeneratedOn apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *IPSECTunnelNewResponseIPSECTunnelsPSKMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelNewResponseIPSECTunnelsPSKMetadataJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate IPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type IPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckType `json:"type"`
	JSON ipsecTunnelNewResponseIPSECTunnelsTunnelHealthCheckJSON `json:"-"`
}

// ipsecTunnelNewResponseIPSECTunnelsTunnelHealthCheckJSON contains the JSON
// metadata for the struct [IPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheck]
type ipsecTunnelNewResponseIPSECTunnelsTunnelHealthCheckJSON struct {
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelNewResponseIPSECTunnelsTunnelHealthCheckJSON) RawJSON() string {
	return r.raw
}

// How frequent the health check is run. The default value is `mid`.
type IPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckRate string

const (
	IPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckRateLow  IPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckRate = "low"
	IPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckRateMid  IPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckRate = "mid"
	IPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckRateHigh IPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type IPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckType string

const (
	IPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckTypeReply   IPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckType = "reply"
	IPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckTypeRequest IPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckType = "request"
)

type IPSECTunnelUpdateResponse struct {
	Modified            bool                          `json:"modified"`
	ModifiedIPSECTunnel interface{}                   `json:"modified_ipsec_tunnel"`
	JSON                ipsecTunnelUpdateResponseJSON `json:"-"`
}

// ipsecTunnelUpdateResponseJSON contains the JSON metadata for the struct
// [IPSECTunnelUpdateResponse]
type ipsecTunnelUpdateResponseJSON struct {
	Modified            apijson.Field
	ModifiedIPSECTunnel apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *IPSECTunnelUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelListResponse struct {
	IPSECTunnels []IPSECTunnelListResponseIPSECTunnel `json:"ipsec_tunnels"`
	JSON         ipsecTunnelListResponseJSON          `json:"-"`
}

// ipsecTunnelListResponseJSON contains the JSON metadata for the struct
// [IPSECTunnelListResponse]
type ipsecTunnelListResponseJSON struct {
	IPSECTunnels apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IPSECTunnelListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelListResponseJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelListResponseIPSECTunnel struct {
	// The IP address assigned to the Cloudflare side of the IPsec tunnel.
	CloudflareEndpoint string `json:"cloudflare_endpoint,required"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress string `json:"interface_address,required"`
	// The name of the IPsec tunnel. The name cannot share a name with other tunnels.
	Name string `json:"name,required"`
	// Tunnel identifier tag.
	ID string `json:"id"`
	// When `true`, the tunnel can use a null-cipher (`ENCR_NULL`) in the ESP tunnel
	// (Phase 2).
	AllowNullCipher bool `json:"allow_null_cipher"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The IP address assigned to the customer side of the IPsec tunnel.
	CustomerEndpoint string `json:"customer_endpoint"`
	// An optional description forthe IPsec tunnel.
	Description string `json:"description"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The PSK metadata that includes when the PSK was generated.
	PSKMetadata IPSECTunnelListResponseIPSECTunnelsPSKMetadata `json:"psk_metadata"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection  bool                                                 `json:"replay_protection"`
	TunnelHealthCheck IPSECTunnelListResponseIPSECTunnelsTunnelHealthCheck `json:"tunnel_health_check"`
	JSON              ipsecTunnelListResponseIPSECTunnelJSON               `json:"-"`
}

// ipsecTunnelListResponseIPSECTunnelJSON contains the JSON metadata for the struct
// [IPSECTunnelListResponseIPSECTunnel]
type ipsecTunnelListResponseIPSECTunnelJSON struct {
	CloudflareEndpoint apijson.Field
	InterfaceAddress   apijson.Field
	Name               apijson.Field
	ID                 apijson.Field
	AllowNullCipher    apijson.Field
	CreatedOn          apijson.Field
	CustomerEndpoint   apijson.Field
	Description        apijson.Field
	ModifiedOn         apijson.Field
	PSKMetadata        apijson.Field
	ReplayProtection   apijson.Field
	TunnelHealthCheck  apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *IPSECTunnelListResponseIPSECTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelListResponseIPSECTunnelJSON) RawJSON() string {
	return r.raw
}

// The PSK metadata that includes when the PSK was generated.
type IPSECTunnelListResponseIPSECTunnelsPSKMetadata struct {
	// The date and time the tunnel was last modified.
	LastGeneratedOn time.Time                                          `json:"last_generated_on" format:"date-time"`
	JSON            ipsecTunnelListResponseIPSECTunnelsPSKMetadataJSON `json:"-"`
}

// ipsecTunnelListResponseIPSECTunnelsPSKMetadataJSON contains the JSON metadata
// for the struct [IPSECTunnelListResponseIPSECTunnelsPSKMetadata]
type ipsecTunnelListResponseIPSECTunnelsPSKMetadataJSON struct {
	LastGeneratedOn apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *IPSECTunnelListResponseIPSECTunnelsPSKMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelListResponseIPSECTunnelsPSKMetadataJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelListResponseIPSECTunnelsTunnelHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate IPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type IPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckType `json:"type"`
	JSON ipsecTunnelListResponseIPSECTunnelsTunnelHealthCheckJSON `json:"-"`
}

// ipsecTunnelListResponseIPSECTunnelsTunnelHealthCheckJSON contains the JSON
// metadata for the struct [IPSECTunnelListResponseIPSECTunnelsTunnelHealthCheck]
type ipsecTunnelListResponseIPSECTunnelsTunnelHealthCheckJSON struct {
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelListResponseIPSECTunnelsTunnelHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelListResponseIPSECTunnelsTunnelHealthCheckJSON) RawJSON() string {
	return r.raw
}

// How frequent the health check is run. The default value is `mid`.
type IPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckRate string

const (
	IPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckRateLow  IPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckRate = "low"
	IPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckRateMid  IPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckRate = "mid"
	IPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckRateHigh IPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type IPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckType string

const (
	IPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckTypeReply   IPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckType = "reply"
	IPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckTypeRequest IPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckType = "request"
)

type IPSECTunnelDeleteResponse struct {
	Deleted            bool                          `json:"deleted"`
	DeletedIPSECTunnel interface{}                   `json:"deleted_ipsec_tunnel"`
	JSON               ipsecTunnelDeleteResponseJSON `json:"-"`
}

// ipsecTunnelDeleteResponseJSON contains the JSON metadata for the struct
// [IPSECTunnelDeleteResponse]
type ipsecTunnelDeleteResponseJSON struct {
	Deleted            apijson.Field
	DeletedIPSECTunnel apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *IPSECTunnelDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelGetResponse struct {
	IPSECTunnel interface{}                `json:"ipsec_tunnel"`
	JSON        ipsecTunnelGetResponseJSON `json:"-"`
}

// ipsecTunnelGetResponseJSON contains the JSON metadata for the struct
// [IPSECTunnelGetResponse]
type ipsecTunnelGetResponseJSON struct {
	IPSECTunnel apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelGetResponseJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelPSKGenerateResponse struct {
	// Identifier
	IPSECTunnelID string `json:"ipsec_tunnel_id"`
	// A randomly generated or provided string for use in the IPsec tunnel.
	PSK string `json:"psk"`
	// The PSK metadata that includes when the PSK was generated.
	PSKMetadata IPSECTunnelPSKGenerateResponsePSKMetadata `json:"psk_metadata"`
	JSON        ipsecTunnelPSKGenerateResponseJSON        `json:"-"`
}

// ipsecTunnelPSKGenerateResponseJSON contains the JSON metadata for the struct
// [IPSECTunnelPSKGenerateResponse]
type ipsecTunnelPSKGenerateResponseJSON struct {
	IPSECTunnelID apijson.Field
	PSK           apijson.Field
	PSKMetadata   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *IPSECTunnelPSKGenerateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelPSKGenerateResponseJSON) RawJSON() string {
	return r.raw
}

// The PSK metadata that includes when the PSK was generated.
type IPSECTunnelPSKGenerateResponsePSKMetadata struct {
	// The date and time the tunnel was last modified.
	LastGeneratedOn time.Time                                     `json:"last_generated_on" format:"date-time"`
	JSON            ipsecTunnelPSKGenerateResponsePSKMetadataJSON `json:"-"`
}

// ipsecTunnelPSKGenerateResponsePSKMetadataJSON contains the JSON metadata for the
// struct [IPSECTunnelPSKGenerateResponsePSKMetadata]
type ipsecTunnelPSKGenerateResponsePSKMetadataJSON struct {
	LastGeneratedOn apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *IPSECTunnelPSKGenerateResponsePSKMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelPSKGenerateResponsePSKMetadataJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelNewParams struct {
	// The IP address assigned to the Cloudflare side of the IPsec tunnel.
	CloudflareEndpoint param.Field[string] `json:"cloudflare_endpoint,required"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress param.Field[string] `json:"interface_address,required"`
	// The name of the IPsec tunnel. The name cannot share a name with other tunnels.
	Name param.Field[string] `json:"name,required"`
	// The IP address assigned to the customer side of the IPsec tunnel.
	CustomerEndpoint param.Field[string] `json:"customer_endpoint"`
	// An optional description forthe IPsec tunnel.
	Description param.Field[string]                          `json:"description"`
	HealthCheck param.Field[IPSECTunnelNewParamsHealthCheck] `json:"health_check"`
	// A randomly generated or provided string for use in the IPsec tunnel.
	PSK param.Field[string] `json:"psk"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection param.Field[bool] `json:"replay_protection"`
}

func (r IPSECTunnelNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IPSECTunnelNewParamsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
	// in health_check is ignored as the interface_address is used to send traffic into
	// the tunnel.
	Direction param.Field[IPSECTunnelNewParamsHealthCheckDirection] `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled param.Field[bool] `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate param.Field[IPSECTunnelNewParamsHealthCheckRate] `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target.
	Target param.Field[string] `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type param.Field[IPSECTunnelNewParamsHealthCheckType] `json:"type"`
}

func (r IPSECTunnelNewParamsHealthCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
// in health_check is ignored as the interface_address is used to send traffic into
// the tunnel.
type IPSECTunnelNewParamsHealthCheckDirection string

const (
	IPSECTunnelNewParamsHealthCheckDirectionUnidirectional IPSECTunnelNewParamsHealthCheckDirection = "unidirectional"
	IPSECTunnelNewParamsHealthCheckDirectionBidirectional  IPSECTunnelNewParamsHealthCheckDirection = "bidirectional"
)

// How frequent the health check is run. The default value is `mid`.
type IPSECTunnelNewParamsHealthCheckRate string

const (
	IPSECTunnelNewParamsHealthCheckRateLow  IPSECTunnelNewParamsHealthCheckRate = "low"
	IPSECTunnelNewParamsHealthCheckRateMid  IPSECTunnelNewParamsHealthCheckRate = "mid"
	IPSECTunnelNewParamsHealthCheckRateHigh IPSECTunnelNewParamsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type IPSECTunnelNewParamsHealthCheckType string

const (
	IPSECTunnelNewParamsHealthCheckTypeReply   IPSECTunnelNewParamsHealthCheckType = "reply"
	IPSECTunnelNewParamsHealthCheckTypeRequest IPSECTunnelNewParamsHealthCheckType = "request"
)

type IPSECTunnelNewResponseEnvelope struct {
	Errors   []IPSECTunnelNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IPSECTunnelNewResponseEnvelopeMessages `json:"messages,required"`
	Result   IPSECTunnelNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success IPSECTunnelNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    ipsecTunnelNewResponseEnvelopeJSON    `json:"-"`
}

// ipsecTunnelNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [IPSECTunnelNewResponseEnvelope]
type ipsecTunnelNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelNewResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    ipsecTunnelNewResponseEnvelopeErrorsJSON `json:"-"`
}

// ipsecTunnelNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [IPSECTunnelNewResponseEnvelopeErrors]
type ipsecTunnelNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelNewResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    ipsecTunnelNewResponseEnvelopeMessagesJSON `json:"-"`
}

// ipsecTunnelNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [IPSECTunnelNewResponseEnvelopeMessages]
type ipsecTunnelNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IPSECTunnelNewResponseEnvelopeSuccess bool

const (
	IPSECTunnelNewResponseEnvelopeSuccessTrue IPSECTunnelNewResponseEnvelopeSuccess = true
)

type IPSECTunnelUpdateParams struct {
	// The IP address assigned to the Cloudflare side of the IPsec tunnel.
	CloudflareEndpoint param.Field[string] `json:"cloudflare_endpoint,required"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress param.Field[string] `json:"interface_address,required"`
	// The name of the IPsec tunnel. The name cannot share a name with other tunnels.
	Name param.Field[string] `json:"name,required"`
	// The IP address assigned to the customer side of the IPsec tunnel.
	CustomerEndpoint param.Field[string] `json:"customer_endpoint"`
	// An optional description forthe IPsec tunnel.
	Description param.Field[string]                             `json:"description"`
	HealthCheck param.Field[IPSECTunnelUpdateParamsHealthCheck] `json:"health_check"`
	// A randomly generated or provided string for use in the IPsec tunnel.
	PSK param.Field[string] `json:"psk"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection param.Field[bool] `json:"replay_protection"`
}

func (r IPSECTunnelUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IPSECTunnelUpdateParamsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
	// in health_check is ignored as the interface_address is used to send traffic into
	// the tunnel.
	Direction param.Field[IPSECTunnelUpdateParamsHealthCheckDirection] `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled param.Field[bool] `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate param.Field[IPSECTunnelUpdateParamsHealthCheckRate] `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target.
	Target param.Field[string] `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type param.Field[IPSECTunnelUpdateParamsHealthCheckType] `json:"type"`
}

func (r IPSECTunnelUpdateParamsHealthCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
// in health_check is ignored as the interface_address is used to send traffic into
// the tunnel.
type IPSECTunnelUpdateParamsHealthCheckDirection string

const (
	IPSECTunnelUpdateParamsHealthCheckDirectionUnidirectional IPSECTunnelUpdateParamsHealthCheckDirection = "unidirectional"
	IPSECTunnelUpdateParamsHealthCheckDirectionBidirectional  IPSECTunnelUpdateParamsHealthCheckDirection = "bidirectional"
)

// How frequent the health check is run. The default value is `mid`.
type IPSECTunnelUpdateParamsHealthCheckRate string

const (
	IPSECTunnelUpdateParamsHealthCheckRateLow  IPSECTunnelUpdateParamsHealthCheckRate = "low"
	IPSECTunnelUpdateParamsHealthCheckRateMid  IPSECTunnelUpdateParamsHealthCheckRate = "mid"
	IPSECTunnelUpdateParamsHealthCheckRateHigh IPSECTunnelUpdateParamsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type IPSECTunnelUpdateParamsHealthCheckType string

const (
	IPSECTunnelUpdateParamsHealthCheckTypeReply   IPSECTunnelUpdateParamsHealthCheckType = "reply"
	IPSECTunnelUpdateParamsHealthCheckTypeRequest IPSECTunnelUpdateParamsHealthCheckType = "request"
)

type IPSECTunnelUpdateResponseEnvelope struct {
	Errors   []IPSECTunnelUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IPSECTunnelUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   IPSECTunnelUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success IPSECTunnelUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    ipsecTunnelUpdateResponseEnvelopeJSON    `json:"-"`
}

// ipsecTunnelUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [IPSECTunnelUpdateResponseEnvelope]
type ipsecTunnelUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelUpdateResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    ipsecTunnelUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// ipsecTunnelUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [IPSECTunnelUpdateResponseEnvelopeErrors]
type ipsecTunnelUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelUpdateResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    ipsecTunnelUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// ipsecTunnelUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [IPSECTunnelUpdateResponseEnvelopeMessages]
type ipsecTunnelUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IPSECTunnelUpdateResponseEnvelopeSuccess bool

const (
	IPSECTunnelUpdateResponseEnvelopeSuccessTrue IPSECTunnelUpdateResponseEnvelopeSuccess = true
)

type IPSECTunnelListResponseEnvelope struct {
	Errors   []IPSECTunnelListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IPSECTunnelListResponseEnvelopeMessages `json:"messages,required"`
	Result   IPSECTunnelListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success IPSECTunnelListResponseEnvelopeSuccess `json:"success,required"`
	JSON    ipsecTunnelListResponseEnvelopeJSON    `json:"-"`
}

// ipsecTunnelListResponseEnvelopeJSON contains the JSON metadata for the struct
// [IPSECTunnelListResponseEnvelope]
type ipsecTunnelListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelListResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    ipsecTunnelListResponseEnvelopeErrorsJSON `json:"-"`
}

// ipsecTunnelListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [IPSECTunnelListResponseEnvelopeErrors]
type ipsecTunnelListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelListResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    ipsecTunnelListResponseEnvelopeMessagesJSON `json:"-"`
}

// ipsecTunnelListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [IPSECTunnelListResponseEnvelopeMessages]
type ipsecTunnelListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IPSECTunnelListResponseEnvelopeSuccess bool

const (
	IPSECTunnelListResponseEnvelopeSuccessTrue IPSECTunnelListResponseEnvelopeSuccess = true
)

type IPSECTunnelDeleteResponseEnvelope struct {
	Errors   []IPSECTunnelDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IPSECTunnelDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   IPSECTunnelDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success IPSECTunnelDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    ipsecTunnelDeleteResponseEnvelopeJSON    `json:"-"`
}

// ipsecTunnelDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [IPSECTunnelDeleteResponseEnvelope]
type ipsecTunnelDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelDeleteResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    ipsecTunnelDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// ipsecTunnelDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [IPSECTunnelDeleteResponseEnvelopeErrors]
type ipsecTunnelDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelDeleteResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    ipsecTunnelDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// ipsecTunnelDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [IPSECTunnelDeleteResponseEnvelopeMessages]
type ipsecTunnelDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IPSECTunnelDeleteResponseEnvelopeSuccess bool

const (
	IPSECTunnelDeleteResponseEnvelopeSuccessTrue IPSECTunnelDeleteResponseEnvelopeSuccess = true
)

type IPSECTunnelGetResponseEnvelope struct {
	Errors   []IPSECTunnelGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IPSECTunnelGetResponseEnvelopeMessages `json:"messages,required"`
	Result   IPSECTunnelGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success IPSECTunnelGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    ipsecTunnelGetResponseEnvelopeJSON    `json:"-"`
}

// ipsecTunnelGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [IPSECTunnelGetResponseEnvelope]
type ipsecTunnelGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    ipsecTunnelGetResponseEnvelopeErrorsJSON `json:"-"`
}

// ipsecTunnelGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [IPSECTunnelGetResponseEnvelopeErrors]
type ipsecTunnelGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    ipsecTunnelGetResponseEnvelopeMessagesJSON `json:"-"`
}

// ipsecTunnelGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [IPSECTunnelGetResponseEnvelopeMessages]
type ipsecTunnelGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IPSECTunnelGetResponseEnvelopeSuccess bool

const (
	IPSECTunnelGetResponseEnvelopeSuccessTrue IPSECTunnelGetResponseEnvelopeSuccess = true
)

type IPSECTunnelPSKGenerateResponseEnvelope struct {
	Errors   []IPSECTunnelPSKGenerateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IPSECTunnelPSKGenerateResponseEnvelopeMessages `json:"messages,required"`
	Result   IPSECTunnelPSKGenerateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success IPSECTunnelPSKGenerateResponseEnvelopeSuccess `json:"success,required"`
	JSON    ipsecTunnelPSKGenerateResponseEnvelopeJSON    `json:"-"`
}

// ipsecTunnelPSKGenerateResponseEnvelopeJSON contains the JSON metadata for the
// struct [IPSECTunnelPSKGenerateResponseEnvelope]
type ipsecTunnelPSKGenerateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelPSKGenerateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelPSKGenerateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelPSKGenerateResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    ipsecTunnelPSKGenerateResponseEnvelopeErrorsJSON `json:"-"`
}

// ipsecTunnelPSKGenerateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [IPSECTunnelPSKGenerateResponseEnvelopeErrors]
type ipsecTunnelPSKGenerateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelPSKGenerateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelPSKGenerateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelPSKGenerateResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    ipsecTunnelPSKGenerateResponseEnvelopeMessagesJSON `json:"-"`
}

// ipsecTunnelPSKGenerateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [IPSECTunnelPSKGenerateResponseEnvelopeMessages]
type ipsecTunnelPSKGenerateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelPSKGenerateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelPSKGenerateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IPSECTunnelPSKGenerateResponseEnvelopeSuccess bool

const (
	IPSECTunnelPSKGenerateResponseEnvelopeSuccessTrue IPSECTunnelPSKGenerateResponseEnvelopeSuccess = true
)
