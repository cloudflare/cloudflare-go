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

// MagicTransitIPSECTunnelService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewMagicTransitIPSECTunnelService] method instead.
type MagicTransitIPSECTunnelService struct {
	Options []option.RequestOption
}

// NewMagicTransitIPSECTunnelService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMagicTransitIPSECTunnelService(opts ...option.RequestOption) (r *MagicTransitIPSECTunnelService) {
	r = &MagicTransitIPSECTunnelService{}
	r.Options = opts
	return
}

// Creates new IPsec tunnels associated with an account. Use `?validate_only=true`
// as an optional query parameter to only run validation without persisting
// changes.
func (r *MagicTransitIPSECTunnelService) New(ctx context.Context, accountIdentifier string, body MagicTransitIPSECTunnelNewParams, opts ...option.RequestOption) (res *MagicTransitIPSECTunnelNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicTransitIPSECTunnelNewResponseEnvelope
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
func (r *MagicTransitIPSECTunnelService) Update(ctx context.Context, accountIdentifier string, tunnelIdentifier string, body MagicTransitIPSECTunnelUpdateParams, opts ...option.RequestOption) (res *MagicTransitIPSECTunnelUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicTransitIPSECTunnelUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists IPsec tunnels associated with an account.
func (r *MagicTransitIPSECTunnelService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *MagicTransitIPSECTunnelListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicTransitIPSECTunnelListResponseEnvelope
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
func (r *MagicTransitIPSECTunnelService) Delete(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *MagicTransitIPSECTunnelDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicTransitIPSECTunnelDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists details for a specific IPsec tunnel.
func (r *MagicTransitIPSECTunnelService) Get(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *MagicTransitIPSECTunnelGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicTransitIPSECTunnelGetResponseEnvelope
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
func (r *MagicTransitIPSECTunnelService) PSKGenerate(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *MagicTransitIPSECTunnelPSKGenerateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicTransitIPSECTunnelPSKGenerateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s/psk_generate", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MagicTransitIPSECTunnelNewResponse struct {
	IPSECTunnels []MagicTransitIPSECTunnelNewResponseIPSECTunnel `json:"ipsec_tunnels"`
	JSON         magicTransitIPSECTunnelNewResponseJSON          `json:"-"`
}

// magicTransitIPSECTunnelNewResponseJSON contains the JSON metadata for the struct
// [MagicTransitIPSECTunnelNewResponse]
type magicTransitIPSECTunnelNewResponseJSON struct {
	IPSECTunnels apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitIPSECTunnelNewResponseJSON) RawJSON() string {
	return r.raw
}

type MagicTransitIPSECTunnelNewResponseIPSECTunnel struct {
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
	PSKMetadata MagicTransitIPSECTunnelNewResponseIPSECTunnelsPSKMetadata `json:"psk_metadata"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection  bool                                                            `json:"replay_protection"`
	TunnelHealthCheck MagicTransitIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheck `json:"tunnel_health_check"`
	JSON              magicTransitIPSECTunnelNewResponseIPSECTunnelJSON               `json:"-"`
}

// magicTransitIPSECTunnelNewResponseIPSECTunnelJSON contains the JSON metadata for
// the struct [MagicTransitIPSECTunnelNewResponseIPSECTunnel]
type magicTransitIPSECTunnelNewResponseIPSECTunnelJSON struct {
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

func (r *MagicTransitIPSECTunnelNewResponseIPSECTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitIPSECTunnelNewResponseIPSECTunnelJSON) RawJSON() string {
	return r.raw
}

// The PSK metadata that includes when the PSK was generated.
type MagicTransitIPSECTunnelNewResponseIPSECTunnelsPSKMetadata struct {
	// The date and time the tunnel was last modified.
	LastGeneratedOn time.Time                                                     `json:"last_generated_on" format:"date-time"`
	JSON            magicTransitIPSECTunnelNewResponseIPSECTunnelsPSKMetadataJSON `json:"-"`
}

// magicTransitIPSECTunnelNewResponseIPSECTunnelsPSKMetadataJSON contains the JSON
// metadata for the struct
// [MagicTransitIPSECTunnelNewResponseIPSECTunnelsPSKMetadata]
type magicTransitIPSECTunnelNewResponseIPSECTunnelsPSKMetadataJSON struct {
	LastGeneratedOn apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelNewResponseIPSECTunnelsPSKMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitIPSECTunnelNewResponseIPSECTunnelsPSKMetadataJSON) RawJSON() string {
	return r.raw
}

type MagicTransitIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate MagicTransitIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type MagicTransitIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckType `json:"type"`
	JSON magicTransitIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckJSON `json:"-"`
}

// magicTransitIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckJSON contains the
// JSON metadata for the struct
// [MagicTransitIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheck]
type magicTransitIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckJSON struct {
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckJSON) RawJSON() string {
	return r.raw
}

// How frequent the health check is run. The default value is `mid`.
type MagicTransitIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckRate string

const (
	MagicTransitIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckRateLow  MagicTransitIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckRate = "low"
	MagicTransitIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckRateMid  MagicTransitIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckRate = "mid"
	MagicTransitIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckRateHigh MagicTransitIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicTransitIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckType string

const (
	MagicTransitIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckTypeReply   MagicTransitIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckType = "reply"
	MagicTransitIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckTypeRequest MagicTransitIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckType = "request"
)

type MagicTransitIPSECTunnelUpdateResponse struct {
	Modified            bool                                      `json:"modified"`
	ModifiedIPSECTunnel interface{}                               `json:"modified_ipsec_tunnel"`
	JSON                magicTransitIPSECTunnelUpdateResponseJSON `json:"-"`
}

// magicTransitIPSECTunnelUpdateResponseJSON contains the JSON metadata for the
// struct [MagicTransitIPSECTunnelUpdateResponse]
type magicTransitIPSECTunnelUpdateResponseJSON struct {
	Modified            apijson.Field
	ModifiedIPSECTunnel apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitIPSECTunnelUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type MagicTransitIPSECTunnelListResponse struct {
	IPSECTunnels []MagicTransitIPSECTunnelListResponseIPSECTunnel `json:"ipsec_tunnels"`
	JSON         magicTransitIPSECTunnelListResponseJSON          `json:"-"`
}

// magicTransitIPSECTunnelListResponseJSON contains the JSON metadata for the
// struct [MagicTransitIPSECTunnelListResponse]
type magicTransitIPSECTunnelListResponseJSON struct {
	IPSECTunnels apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitIPSECTunnelListResponseJSON) RawJSON() string {
	return r.raw
}

type MagicTransitIPSECTunnelListResponseIPSECTunnel struct {
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
	PSKMetadata MagicTransitIPSECTunnelListResponseIPSECTunnelsPSKMetadata `json:"psk_metadata"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection  bool                                                             `json:"replay_protection"`
	TunnelHealthCheck MagicTransitIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheck `json:"tunnel_health_check"`
	JSON              magicTransitIPSECTunnelListResponseIPSECTunnelJSON               `json:"-"`
}

// magicTransitIPSECTunnelListResponseIPSECTunnelJSON contains the JSON metadata
// for the struct [MagicTransitIPSECTunnelListResponseIPSECTunnel]
type magicTransitIPSECTunnelListResponseIPSECTunnelJSON struct {
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

func (r *MagicTransitIPSECTunnelListResponseIPSECTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitIPSECTunnelListResponseIPSECTunnelJSON) RawJSON() string {
	return r.raw
}

// The PSK metadata that includes when the PSK was generated.
type MagicTransitIPSECTunnelListResponseIPSECTunnelsPSKMetadata struct {
	// The date and time the tunnel was last modified.
	LastGeneratedOn time.Time                                                      `json:"last_generated_on" format:"date-time"`
	JSON            magicTransitIPSECTunnelListResponseIPSECTunnelsPSKMetadataJSON `json:"-"`
}

// magicTransitIPSECTunnelListResponseIPSECTunnelsPSKMetadataJSON contains the JSON
// metadata for the struct
// [MagicTransitIPSECTunnelListResponseIPSECTunnelsPSKMetadata]
type magicTransitIPSECTunnelListResponseIPSECTunnelsPSKMetadataJSON struct {
	LastGeneratedOn apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelListResponseIPSECTunnelsPSKMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitIPSECTunnelListResponseIPSECTunnelsPSKMetadataJSON) RawJSON() string {
	return r.raw
}

type MagicTransitIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate MagicTransitIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type MagicTransitIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckType `json:"type"`
	JSON magicTransitIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckJSON `json:"-"`
}

// magicTransitIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckJSON contains
// the JSON metadata for the struct
// [MagicTransitIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheck]
type magicTransitIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckJSON struct {
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckJSON) RawJSON() string {
	return r.raw
}

// How frequent the health check is run. The default value is `mid`.
type MagicTransitIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckRate string

const (
	MagicTransitIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckRateLow  MagicTransitIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckRate = "low"
	MagicTransitIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckRateMid  MagicTransitIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckRate = "mid"
	MagicTransitIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckRateHigh MagicTransitIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicTransitIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckType string

const (
	MagicTransitIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckTypeReply   MagicTransitIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckType = "reply"
	MagicTransitIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckTypeRequest MagicTransitIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckType = "request"
)

type MagicTransitIPSECTunnelDeleteResponse struct {
	Deleted            bool                                      `json:"deleted"`
	DeletedIPSECTunnel interface{}                               `json:"deleted_ipsec_tunnel"`
	JSON               magicTransitIPSECTunnelDeleteResponseJSON `json:"-"`
}

// magicTransitIPSECTunnelDeleteResponseJSON contains the JSON metadata for the
// struct [MagicTransitIPSECTunnelDeleteResponse]
type magicTransitIPSECTunnelDeleteResponseJSON struct {
	Deleted            apijson.Field
	DeletedIPSECTunnel apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitIPSECTunnelDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type MagicTransitIPSECTunnelGetResponse struct {
	IPSECTunnel interface{}                            `json:"ipsec_tunnel"`
	JSON        magicTransitIPSECTunnelGetResponseJSON `json:"-"`
}

// magicTransitIPSECTunnelGetResponseJSON contains the JSON metadata for the struct
// [MagicTransitIPSECTunnelGetResponse]
type magicTransitIPSECTunnelGetResponseJSON struct {
	IPSECTunnel apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitIPSECTunnelGetResponseJSON) RawJSON() string {
	return r.raw
}

type MagicTransitIPSECTunnelPSKGenerateResponse struct {
	// Identifier
	IPSECTunnelID string `json:"ipsec_tunnel_id"`
	// A randomly generated or provided string for use in the IPsec tunnel.
	PSK string `json:"psk"`
	// The PSK metadata that includes when the PSK was generated.
	PSKMetadata MagicTransitIPSECTunnelPSKGenerateResponsePSKMetadata `json:"psk_metadata"`
	JSON        magicTransitIPSECTunnelPSKGenerateResponseJSON        `json:"-"`
}

// magicTransitIPSECTunnelPSKGenerateResponseJSON contains the JSON metadata for
// the struct [MagicTransitIPSECTunnelPSKGenerateResponse]
type magicTransitIPSECTunnelPSKGenerateResponseJSON struct {
	IPSECTunnelID apijson.Field
	PSK           apijson.Field
	PSKMetadata   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelPSKGenerateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitIPSECTunnelPSKGenerateResponseJSON) RawJSON() string {
	return r.raw
}

// The PSK metadata that includes when the PSK was generated.
type MagicTransitIPSECTunnelPSKGenerateResponsePSKMetadata struct {
	// The date and time the tunnel was last modified.
	LastGeneratedOn time.Time                                                 `json:"last_generated_on" format:"date-time"`
	JSON            magicTransitIPSECTunnelPSKGenerateResponsePSKMetadataJSON `json:"-"`
}

// magicTransitIPSECTunnelPSKGenerateResponsePSKMetadataJSON contains the JSON
// metadata for the struct [MagicTransitIPSECTunnelPSKGenerateResponsePSKMetadata]
type magicTransitIPSECTunnelPSKGenerateResponsePSKMetadataJSON struct {
	LastGeneratedOn apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelPSKGenerateResponsePSKMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitIPSECTunnelPSKGenerateResponsePSKMetadataJSON) RawJSON() string {
	return r.raw
}

type MagicTransitIPSECTunnelNewParams struct {
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
	Description param.Field[string]                                      `json:"description"`
	HealthCheck param.Field[MagicTransitIPSECTunnelNewParamsHealthCheck] `json:"health_check"`
	// A randomly generated or provided string for use in the IPsec tunnel.
	PSK param.Field[string] `json:"psk"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection param.Field[bool] `json:"replay_protection"`
}

func (r MagicTransitIPSECTunnelNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagicTransitIPSECTunnelNewParamsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
	// in health_check is ignored as the interface_address is used to send traffic into
	// the tunnel.
	Direction param.Field[MagicTransitIPSECTunnelNewParamsHealthCheckDirection] `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled param.Field[bool] `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate param.Field[MagicTransitIPSECTunnelNewParamsHealthCheckRate] `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target.
	Target param.Field[string] `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type param.Field[MagicTransitIPSECTunnelNewParamsHealthCheckType] `json:"type"`
}

func (r MagicTransitIPSECTunnelNewParamsHealthCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
// in health_check is ignored as the interface_address is used to send traffic into
// the tunnel.
type MagicTransitIPSECTunnelNewParamsHealthCheckDirection string

const (
	MagicTransitIPSECTunnelNewParamsHealthCheckDirectionUnidirectional MagicTransitIPSECTunnelNewParamsHealthCheckDirection = "unidirectional"
	MagicTransitIPSECTunnelNewParamsHealthCheckDirectionBidirectional  MagicTransitIPSECTunnelNewParamsHealthCheckDirection = "bidirectional"
)

// How frequent the health check is run. The default value is `mid`.
type MagicTransitIPSECTunnelNewParamsHealthCheckRate string

const (
	MagicTransitIPSECTunnelNewParamsHealthCheckRateLow  MagicTransitIPSECTunnelNewParamsHealthCheckRate = "low"
	MagicTransitIPSECTunnelNewParamsHealthCheckRateMid  MagicTransitIPSECTunnelNewParamsHealthCheckRate = "mid"
	MagicTransitIPSECTunnelNewParamsHealthCheckRateHigh MagicTransitIPSECTunnelNewParamsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicTransitIPSECTunnelNewParamsHealthCheckType string

const (
	MagicTransitIPSECTunnelNewParamsHealthCheckTypeReply   MagicTransitIPSECTunnelNewParamsHealthCheckType = "reply"
	MagicTransitIPSECTunnelNewParamsHealthCheckTypeRequest MagicTransitIPSECTunnelNewParamsHealthCheckType = "request"
)

type MagicTransitIPSECTunnelNewResponseEnvelope struct {
	Errors   []MagicTransitIPSECTunnelNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicTransitIPSECTunnelNewResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicTransitIPSECTunnelNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicTransitIPSECTunnelNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicTransitIPSECTunnelNewResponseEnvelopeJSON    `json:"-"`
}

// magicTransitIPSECTunnelNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [MagicTransitIPSECTunnelNewResponseEnvelope]
type magicTransitIPSECTunnelNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitIPSECTunnelNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MagicTransitIPSECTunnelNewResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    magicTransitIPSECTunnelNewResponseEnvelopeErrorsJSON `json:"-"`
}

// magicTransitIPSECTunnelNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [MagicTransitIPSECTunnelNewResponseEnvelopeErrors]
type magicTransitIPSECTunnelNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitIPSECTunnelNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MagicTransitIPSECTunnelNewResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    magicTransitIPSECTunnelNewResponseEnvelopeMessagesJSON `json:"-"`
}

// magicTransitIPSECTunnelNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [MagicTransitIPSECTunnelNewResponseEnvelopeMessages]
type magicTransitIPSECTunnelNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitIPSECTunnelNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicTransitIPSECTunnelNewResponseEnvelopeSuccess bool

const (
	MagicTransitIPSECTunnelNewResponseEnvelopeSuccessTrue MagicTransitIPSECTunnelNewResponseEnvelopeSuccess = true
)

type MagicTransitIPSECTunnelUpdateParams struct {
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
	Description param.Field[string]                                         `json:"description"`
	HealthCheck param.Field[MagicTransitIPSECTunnelUpdateParamsHealthCheck] `json:"health_check"`
	// A randomly generated or provided string for use in the IPsec tunnel.
	PSK param.Field[string] `json:"psk"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection param.Field[bool] `json:"replay_protection"`
}

func (r MagicTransitIPSECTunnelUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagicTransitIPSECTunnelUpdateParamsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
	// in health_check is ignored as the interface_address is used to send traffic into
	// the tunnel.
	Direction param.Field[MagicTransitIPSECTunnelUpdateParamsHealthCheckDirection] `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled param.Field[bool] `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate param.Field[MagicTransitIPSECTunnelUpdateParamsHealthCheckRate] `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target.
	Target param.Field[string] `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type param.Field[MagicTransitIPSECTunnelUpdateParamsHealthCheckType] `json:"type"`
}

func (r MagicTransitIPSECTunnelUpdateParamsHealthCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
// in health_check is ignored as the interface_address is used to send traffic into
// the tunnel.
type MagicTransitIPSECTunnelUpdateParamsHealthCheckDirection string

const (
	MagicTransitIPSECTunnelUpdateParamsHealthCheckDirectionUnidirectional MagicTransitIPSECTunnelUpdateParamsHealthCheckDirection = "unidirectional"
	MagicTransitIPSECTunnelUpdateParamsHealthCheckDirectionBidirectional  MagicTransitIPSECTunnelUpdateParamsHealthCheckDirection = "bidirectional"
)

// How frequent the health check is run. The default value is `mid`.
type MagicTransitIPSECTunnelUpdateParamsHealthCheckRate string

const (
	MagicTransitIPSECTunnelUpdateParamsHealthCheckRateLow  MagicTransitIPSECTunnelUpdateParamsHealthCheckRate = "low"
	MagicTransitIPSECTunnelUpdateParamsHealthCheckRateMid  MagicTransitIPSECTunnelUpdateParamsHealthCheckRate = "mid"
	MagicTransitIPSECTunnelUpdateParamsHealthCheckRateHigh MagicTransitIPSECTunnelUpdateParamsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicTransitIPSECTunnelUpdateParamsHealthCheckType string

const (
	MagicTransitIPSECTunnelUpdateParamsHealthCheckTypeReply   MagicTransitIPSECTunnelUpdateParamsHealthCheckType = "reply"
	MagicTransitIPSECTunnelUpdateParamsHealthCheckTypeRequest MagicTransitIPSECTunnelUpdateParamsHealthCheckType = "request"
)

type MagicTransitIPSECTunnelUpdateResponseEnvelope struct {
	Errors   []MagicTransitIPSECTunnelUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicTransitIPSECTunnelUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicTransitIPSECTunnelUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicTransitIPSECTunnelUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicTransitIPSECTunnelUpdateResponseEnvelopeJSON    `json:"-"`
}

// magicTransitIPSECTunnelUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [MagicTransitIPSECTunnelUpdateResponseEnvelope]
type magicTransitIPSECTunnelUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitIPSECTunnelUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MagicTransitIPSECTunnelUpdateResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    magicTransitIPSECTunnelUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// magicTransitIPSECTunnelUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [MagicTransitIPSECTunnelUpdateResponseEnvelopeErrors]
type magicTransitIPSECTunnelUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitIPSECTunnelUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MagicTransitIPSECTunnelUpdateResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    magicTransitIPSECTunnelUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// magicTransitIPSECTunnelUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [MagicTransitIPSECTunnelUpdateResponseEnvelopeMessages]
type magicTransitIPSECTunnelUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitIPSECTunnelUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicTransitIPSECTunnelUpdateResponseEnvelopeSuccess bool

const (
	MagicTransitIPSECTunnelUpdateResponseEnvelopeSuccessTrue MagicTransitIPSECTunnelUpdateResponseEnvelopeSuccess = true
)

type MagicTransitIPSECTunnelListResponseEnvelope struct {
	Errors   []MagicTransitIPSECTunnelListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicTransitIPSECTunnelListResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicTransitIPSECTunnelListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicTransitIPSECTunnelListResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicTransitIPSECTunnelListResponseEnvelopeJSON    `json:"-"`
}

// magicTransitIPSECTunnelListResponseEnvelopeJSON contains the JSON metadata for
// the struct [MagicTransitIPSECTunnelListResponseEnvelope]
type magicTransitIPSECTunnelListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitIPSECTunnelListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MagicTransitIPSECTunnelListResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    magicTransitIPSECTunnelListResponseEnvelopeErrorsJSON `json:"-"`
}

// magicTransitIPSECTunnelListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [MagicTransitIPSECTunnelListResponseEnvelopeErrors]
type magicTransitIPSECTunnelListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitIPSECTunnelListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MagicTransitIPSECTunnelListResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    magicTransitIPSECTunnelListResponseEnvelopeMessagesJSON `json:"-"`
}

// magicTransitIPSECTunnelListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [MagicTransitIPSECTunnelListResponseEnvelopeMessages]
type magicTransitIPSECTunnelListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitIPSECTunnelListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicTransitIPSECTunnelListResponseEnvelopeSuccess bool

const (
	MagicTransitIPSECTunnelListResponseEnvelopeSuccessTrue MagicTransitIPSECTunnelListResponseEnvelopeSuccess = true
)

type MagicTransitIPSECTunnelDeleteResponseEnvelope struct {
	Errors   []MagicTransitIPSECTunnelDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicTransitIPSECTunnelDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicTransitIPSECTunnelDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicTransitIPSECTunnelDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicTransitIPSECTunnelDeleteResponseEnvelopeJSON    `json:"-"`
}

// magicTransitIPSECTunnelDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [MagicTransitIPSECTunnelDeleteResponseEnvelope]
type magicTransitIPSECTunnelDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitIPSECTunnelDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MagicTransitIPSECTunnelDeleteResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    magicTransitIPSECTunnelDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// magicTransitIPSECTunnelDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [MagicTransitIPSECTunnelDeleteResponseEnvelopeErrors]
type magicTransitIPSECTunnelDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitIPSECTunnelDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MagicTransitIPSECTunnelDeleteResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    magicTransitIPSECTunnelDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// magicTransitIPSECTunnelDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [MagicTransitIPSECTunnelDeleteResponseEnvelopeMessages]
type magicTransitIPSECTunnelDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitIPSECTunnelDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicTransitIPSECTunnelDeleteResponseEnvelopeSuccess bool

const (
	MagicTransitIPSECTunnelDeleteResponseEnvelopeSuccessTrue MagicTransitIPSECTunnelDeleteResponseEnvelopeSuccess = true
)

type MagicTransitIPSECTunnelGetResponseEnvelope struct {
	Errors   []MagicTransitIPSECTunnelGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicTransitIPSECTunnelGetResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicTransitIPSECTunnelGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicTransitIPSECTunnelGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicTransitIPSECTunnelGetResponseEnvelopeJSON    `json:"-"`
}

// magicTransitIPSECTunnelGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [MagicTransitIPSECTunnelGetResponseEnvelope]
type magicTransitIPSECTunnelGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitIPSECTunnelGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MagicTransitIPSECTunnelGetResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    magicTransitIPSECTunnelGetResponseEnvelopeErrorsJSON `json:"-"`
}

// magicTransitIPSECTunnelGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [MagicTransitIPSECTunnelGetResponseEnvelopeErrors]
type magicTransitIPSECTunnelGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitIPSECTunnelGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MagicTransitIPSECTunnelGetResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    magicTransitIPSECTunnelGetResponseEnvelopeMessagesJSON `json:"-"`
}

// magicTransitIPSECTunnelGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [MagicTransitIPSECTunnelGetResponseEnvelopeMessages]
type magicTransitIPSECTunnelGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitIPSECTunnelGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicTransitIPSECTunnelGetResponseEnvelopeSuccess bool

const (
	MagicTransitIPSECTunnelGetResponseEnvelopeSuccessTrue MagicTransitIPSECTunnelGetResponseEnvelopeSuccess = true
)

type MagicTransitIPSECTunnelPSKGenerateResponseEnvelope struct {
	Errors   []MagicTransitIPSECTunnelPSKGenerateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicTransitIPSECTunnelPSKGenerateResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicTransitIPSECTunnelPSKGenerateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicTransitIPSECTunnelPSKGenerateResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicTransitIPSECTunnelPSKGenerateResponseEnvelopeJSON    `json:"-"`
}

// magicTransitIPSECTunnelPSKGenerateResponseEnvelopeJSON contains the JSON
// metadata for the struct [MagicTransitIPSECTunnelPSKGenerateResponseEnvelope]
type magicTransitIPSECTunnelPSKGenerateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelPSKGenerateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitIPSECTunnelPSKGenerateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MagicTransitIPSECTunnelPSKGenerateResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    magicTransitIPSECTunnelPSKGenerateResponseEnvelopeErrorsJSON `json:"-"`
}

// magicTransitIPSECTunnelPSKGenerateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [MagicTransitIPSECTunnelPSKGenerateResponseEnvelopeErrors]
type magicTransitIPSECTunnelPSKGenerateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelPSKGenerateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitIPSECTunnelPSKGenerateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MagicTransitIPSECTunnelPSKGenerateResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    magicTransitIPSECTunnelPSKGenerateResponseEnvelopeMessagesJSON `json:"-"`
}

// magicTransitIPSECTunnelPSKGenerateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [MagicTransitIPSECTunnelPSKGenerateResponseEnvelopeMessages]
type magicTransitIPSECTunnelPSKGenerateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelPSKGenerateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitIPSECTunnelPSKGenerateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicTransitIPSECTunnelPSKGenerateResponseEnvelopeSuccess bool

const (
	MagicTransitIPSECTunnelPSKGenerateResponseEnvelopeSuccessTrue MagicTransitIPSECTunnelPSKGenerateResponseEnvelopeSuccess = true
)
