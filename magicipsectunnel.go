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

// MagicIpsecTunnelService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewMagicIpsecTunnelService] method
// instead.
type MagicIpsecTunnelService struct {
	Options      []option.RequestOption
	PskGenerates *MagicIpsecTunnelPskGenerateService
}

// NewMagicIpsecTunnelService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMagicIpsecTunnelService(opts ...option.RequestOption) (r *MagicIpsecTunnelService) {
	r = &MagicIpsecTunnelService{}
	r.Options = opts
	r.PskGenerates = NewMagicIpsecTunnelPskGenerateService(opts...)
	return
}

// Updates a specific IPsec tunnel associated with an account. Use
// `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes.
func (r *MagicIpsecTunnelService) Update(ctx context.Context, accountIdentifier string, tunnelIdentifier string, body MagicIpsecTunnelUpdateParams, opts ...option.RequestOption) (res *MagicIpsecTunnelUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicIpsecTunnelUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Disables and removes a specific static IPsec Tunnel associated with an account.
// Use `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes.
func (r *MagicIpsecTunnelService) Delete(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *MagicIpsecTunnelDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicIpsecTunnelDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists details for a specific IPsec tunnel.
func (r *MagicIpsecTunnelService) Get(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *MagicIpsecTunnelGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicIpsecTunnelGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Creates new IPsec tunnels associated with an account. Use `?validate_only=true`
// as an optional query parameter to only run validation without persisting
// changes.
func (r *MagicIpsecTunnelService) MagicIPsecTunnelsNewIPsecTunnels(ctx context.Context, accountIdentifier string, body MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsParams, opts ...option.RequestOption) (res *MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists IPsec tunnels associated with an account.
func (r *MagicIpsecTunnelService) MagicIPsecTunnelsListIPsecTunnels(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update multiple IPsec tunnels associated with an account. Use
// `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes.
func (r *MagicIpsecTunnelService) MagicIPsecTunnelsUpdateMultipleIPsecTunnels(ctx context.Context, accountIdentifier string, body MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsParams, opts ...option.RequestOption) (res *MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MagicIpsecTunnelUpdateResponse struct {
	Modified            bool                               `json:"modified"`
	ModifiedIpsecTunnel interface{}                        `json:"modified_ipsec_tunnel"`
	JSON                magicIpsecTunnelUpdateResponseJSON `json:"-"`
}

// magicIpsecTunnelUpdateResponseJSON contains the JSON metadata for the struct
// [MagicIpsecTunnelUpdateResponse]
type magicIpsecTunnelUpdateResponseJSON struct {
	Modified            apijson.Field
	ModifiedIpsecTunnel apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *MagicIpsecTunnelUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelDeleteResponse struct {
	Deleted            bool                               `json:"deleted"`
	DeletedIpsecTunnel interface{}                        `json:"deleted_ipsec_tunnel"`
	JSON               magicIpsecTunnelDeleteResponseJSON `json:"-"`
}

// magicIpsecTunnelDeleteResponseJSON contains the JSON metadata for the struct
// [MagicIpsecTunnelDeleteResponse]
type magicIpsecTunnelDeleteResponseJSON struct {
	Deleted            apijson.Field
	DeletedIpsecTunnel apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *MagicIpsecTunnelDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelGetResponse struct {
	IpsecTunnel interface{}                     `json:"ipsec_tunnel"`
	JSON        magicIpsecTunnelGetResponseJSON `json:"-"`
}

// magicIpsecTunnelGetResponseJSON contains the JSON metadata for the struct
// [MagicIpsecTunnelGetResponse]
type magicIpsecTunnelGetResponseJSON struct {
	IpsecTunnel apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponse struct {
	IpsecTunnels []MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnel `json:"ipsec_tunnels"`
	JSON         magicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseJSON          `json:"-"`
}

// magicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseJSON contains the JSON
// metadata for the struct
// [MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponse]
type magicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseJSON struct {
	IpsecTunnels apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnel struct {
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
	PskMetadata MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnelsPskMetadata `json:"psk_metadata"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection  bool                                                                                  `json:"replay_protection"`
	TunnelHealthCheck MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheck `json:"tunnel_health_check"`
	JSON              magicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnelJSON               `json:"-"`
}

// magicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnelJSON contains
// the JSON metadata for the struct
// [MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnel]
type magicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnelJSON struct {
	CloudflareEndpoint apijson.Field
	InterfaceAddress   apijson.Field
	Name               apijson.Field
	ID                 apijson.Field
	AllowNullCipher    apijson.Field
	CreatedOn          apijson.Field
	CustomerEndpoint   apijson.Field
	Description        apijson.Field
	ModifiedOn         apijson.Field
	PskMetadata        apijson.Field
	ReplayProtection   apijson.Field
	TunnelHealthCheck  apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The PSK metadata that includes when the PSK was generated.
type MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnelsPskMetadata struct {
	// The date and time the tunnel was last modified.
	LastGeneratedOn time.Time                                                                           `json:"last_generated_on" format:"date-time"`
	JSON            magicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnelsPskMetadataJSON `json:"-"`
}

// magicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnelsPskMetadataJSON
// contains the JSON metadata for the struct
// [MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnelsPskMetadata]
type magicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnelsPskMetadataJSON struct {
	LastGeneratedOn apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnelsPskMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckType `json:"type"`
	JSON magicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckJSON `json:"-"`
}

// magicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckJSON
// contains the JSON metadata for the struct
// [MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheck]
type magicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckJSON struct {
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// How frequent the health check is run. The default value is `mid`.
type MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckRate string

const (
	MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckRateLow  MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckRate = "low"
	MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckRateMid  MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckRate = "mid"
	MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckRateHigh MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckType string

const (
	MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckTypeReply   MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckType = "reply"
	MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckTypeRequest MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckType = "request"
)

type MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponse struct {
	IpsecTunnels []MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnel `json:"ipsec_tunnels"`
	JSON         magicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseJSON          `json:"-"`
}

// magicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseJSON contains the JSON
// metadata for the struct
// [MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponse]
type magicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseJSON struct {
	IpsecTunnels apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnel struct {
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
	PskMetadata MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnelsPskMetadata `json:"psk_metadata"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection  bool                                                                                   `json:"replay_protection"`
	TunnelHealthCheck MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheck `json:"tunnel_health_check"`
	JSON              magicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnelJSON               `json:"-"`
}

// magicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnelJSON
// contains the JSON metadata for the struct
// [MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnel]
type magicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnelJSON struct {
	CloudflareEndpoint apijson.Field
	InterfaceAddress   apijson.Field
	Name               apijson.Field
	ID                 apijson.Field
	AllowNullCipher    apijson.Field
	CreatedOn          apijson.Field
	CustomerEndpoint   apijson.Field
	Description        apijson.Field
	ModifiedOn         apijson.Field
	PskMetadata        apijson.Field
	ReplayProtection   apijson.Field
	TunnelHealthCheck  apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The PSK metadata that includes when the PSK was generated.
type MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnelsPskMetadata struct {
	// The date and time the tunnel was last modified.
	LastGeneratedOn time.Time                                                                            `json:"last_generated_on" format:"date-time"`
	JSON            magicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnelsPskMetadataJSON `json:"-"`
}

// magicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnelsPskMetadataJSON
// contains the JSON metadata for the struct
// [MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnelsPskMetadata]
type magicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnelsPskMetadataJSON struct {
	LastGeneratedOn apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnelsPskMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckType `json:"type"`
	JSON magicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckJSON `json:"-"`
}

// magicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckJSON
// contains the JSON metadata for the struct
// [MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheck]
type magicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckJSON struct {
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// How frequent the health check is run. The default value is `mid`.
type MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckRate string

const (
	MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckRateLow  MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckRate = "low"
	MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckRateMid  MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckRate = "mid"
	MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckRateHigh MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckType string

const (
	MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckTypeReply   MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckType = "reply"
	MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckTypeRequest MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseIpsecTunnelsTunnelHealthCheckType = "request"
)

type MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponse struct {
	Modified             bool                                                                                     `json:"modified"`
	ModifiedIpsecTunnels []MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnel `json:"modified_ipsec_tunnels"`
	JSON                 magicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseJSON                  `json:"-"`
}

// magicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseJSON contains
// the JSON metadata for the struct
// [MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponse]
type magicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseJSON struct {
	Modified             apijson.Field
	ModifiedIpsecTunnels apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnel struct {
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
	PskMetadata MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnelsPskMetadata `json:"psk_metadata"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection  bool                                                                                                     `json:"replay_protection"`
	TunnelHealthCheck MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnelsTunnelHealthCheck `json:"tunnel_health_check"`
	JSON              magicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnelJSON               `json:"-"`
}

// magicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnelJSON
// contains the JSON metadata for the struct
// [MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnel]
type magicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnelJSON struct {
	CloudflareEndpoint apijson.Field
	InterfaceAddress   apijson.Field
	Name               apijson.Field
	ID                 apijson.Field
	AllowNullCipher    apijson.Field
	CreatedOn          apijson.Field
	CustomerEndpoint   apijson.Field
	Description        apijson.Field
	ModifiedOn         apijson.Field
	PskMetadata        apijson.Field
	ReplayProtection   apijson.Field
	TunnelHealthCheck  apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The PSK metadata that includes when the PSK was generated.
type MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnelsPskMetadata struct {
	// The date and time the tunnel was last modified.
	LastGeneratedOn time.Time                                                                                              `json:"last_generated_on" format:"date-time"`
	JSON            magicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnelsPskMetadataJSON `json:"-"`
}

// magicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnelsPskMetadataJSON
// contains the JSON metadata for the struct
// [MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnelsPskMetadata]
type magicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnelsPskMetadataJSON struct {
	LastGeneratedOn apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnelsPskMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnelsTunnelHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnelsTunnelHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnelsTunnelHealthCheckType `json:"type"`
	JSON magicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnelsTunnelHealthCheckJSON `json:"-"`
}

// magicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnelsTunnelHealthCheckJSON
// contains the JSON metadata for the struct
// [MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnelsTunnelHealthCheck]
type magicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnelsTunnelHealthCheckJSON struct {
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnelsTunnelHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// How frequent the health check is run. The default value is `mid`.
type MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnelsTunnelHealthCheckRate string

const (
	MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnelsTunnelHealthCheckRateLow  MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnelsTunnelHealthCheckRate = "low"
	MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnelsTunnelHealthCheckRateMid  MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnelsTunnelHealthCheckRate = "mid"
	MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnelsTunnelHealthCheckRateHigh MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnelsTunnelHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnelsTunnelHealthCheckType string

const (
	MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnelsTunnelHealthCheckTypeReply   MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnelsTunnelHealthCheckType = "reply"
	MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnelsTunnelHealthCheckTypeRequest MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseModifiedIpsecTunnelsTunnelHealthCheckType = "request"
)

type MagicIpsecTunnelUpdateParams struct {
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
	Description param.Field[string] `json:"description"`
	// A randomly generated or provided string for use in the IPsec tunnel.
	Psk param.Field[string] `json:"psk"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection param.Field[bool] `json:"replay_protection"`
}

func (r MagicIpsecTunnelUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagicIpsecTunnelUpdateResponseEnvelope struct {
	Errors   []MagicIpsecTunnelUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicIpsecTunnelUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicIpsecTunnelUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicIpsecTunnelUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicIpsecTunnelUpdateResponseEnvelopeJSON    `json:"-"`
}

// magicIpsecTunnelUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [MagicIpsecTunnelUpdateResponseEnvelope]
type magicIpsecTunnelUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelUpdateResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    magicIpsecTunnelUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// magicIpsecTunnelUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [MagicIpsecTunnelUpdateResponseEnvelopeErrors]
type magicIpsecTunnelUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelUpdateResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    magicIpsecTunnelUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// magicIpsecTunnelUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [MagicIpsecTunnelUpdateResponseEnvelopeMessages]
type magicIpsecTunnelUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicIpsecTunnelUpdateResponseEnvelopeSuccess bool

const (
	MagicIpsecTunnelUpdateResponseEnvelopeSuccessTrue MagicIpsecTunnelUpdateResponseEnvelopeSuccess = true
)

type MagicIpsecTunnelDeleteResponseEnvelope struct {
	Errors   []MagicIpsecTunnelDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicIpsecTunnelDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicIpsecTunnelDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicIpsecTunnelDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicIpsecTunnelDeleteResponseEnvelopeJSON    `json:"-"`
}

// magicIpsecTunnelDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [MagicIpsecTunnelDeleteResponseEnvelope]
type magicIpsecTunnelDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelDeleteResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    magicIpsecTunnelDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// magicIpsecTunnelDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [MagicIpsecTunnelDeleteResponseEnvelopeErrors]
type magicIpsecTunnelDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelDeleteResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    magicIpsecTunnelDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// magicIpsecTunnelDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [MagicIpsecTunnelDeleteResponseEnvelopeMessages]
type magicIpsecTunnelDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicIpsecTunnelDeleteResponseEnvelopeSuccess bool

const (
	MagicIpsecTunnelDeleteResponseEnvelopeSuccessTrue MagicIpsecTunnelDeleteResponseEnvelopeSuccess = true
)

type MagicIpsecTunnelGetResponseEnvelope struct {
	Errors   []MagicIpsecTunnelGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicIpsecTunnelGetResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicIpsecTunnelGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicIpsecTunnelGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicIpsecTunnelGetResponseEnvelopeJSON    `json:"-"`
}

// magicIpsecTunnelGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [MagicIpsecTunnelGetResponseEnvelope]
type magicIpsecTunnelGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelGetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    magicIpsecTunnelGetResponseEnvelopeErrorsJSON `json:"-"`
}

// magicIpsecTunnelGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MagicIpsecTunnelGetResponseEnvelopeErrors]
type magicIpsecTunnelGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelGetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    magicIpsecTunnelGetResponseEnvelopeMessagesJSON `json:"-"`
}

// magicIpsecTunnelGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [MagicIpsecTunnelGetResponseEnvelopeMessages]
type magicIpsecTunnelGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicIpsecTunnelGetResponseEnvelopeSuccess bool

const (
	MagicIpsecTunnelGetResponseEnvelopeSuccessTrue MagicIpsecTunnelGetResponseEnvelopeSuccess = true
)

type MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsParams struct {
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
	Description param.Field[string] `json:"description"`
	// A randomly generated or provided string for use in the IPsec tunnel.
	Psk param.Field[string] `json:"psk"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection param.Field[bool] `json:"replay_protection"`
}

func (r MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseEnvelope struct {
	Errors   []MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseEnvelopeJSON    `json:"-"`
}

// magicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseEnvelope]
type magicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseEnvelopeErrors struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    magicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseEnvelopeErrorsJSON `json:"-"`
}

// magicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseEnvelopeErrors]
type magicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseEnvelopeMessages struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    magicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseEnvelopeMessagesJSON `json:"-"`
}

// magicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseEnvelopeMessages]
type magicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseEnvelopeSuccess bool

const (
	MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseEnvelopeSuccessTrue MagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseEnvelopeSuccess = true
)

type MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseEnvelope struct {
	Errors   []MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseEnvelopeJSON    `json:"-"`
}

// magicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseEnvelope]
type magicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseEnvelopeErrors struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    magicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseEnvelopeErrorsJSON `json:"-"`
}

// magicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseEnvelopeErrors]
type magicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseEnvelopeMessages struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    magicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseEnvelopeMessagesJSON `json:"-"`
}

// magicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseEnvelopeMessages]
type magicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseEnvelopeSuccess bool

const (
	MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseEnvelopeSuccessTrue MagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseEnvelopeSuccess = true
)

type MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseEnvelope struct {
	Errors   []MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseEnvelopeJSON    `json:"-"`
}

// magicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseEnvelope]
type magicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseEnvelopeErrors struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    magicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseEnvelopeErrorsJSON `json:"-"`
}

// magicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseEnvelopeErrors]
type magicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseEnvelopeMessages struct {
	Code    int64                                                                                   `json:"code,required"`
	Message string                                                                                  `json:"message,required"`
	JSON    magicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseEnvelopeMessagesJSON `json:"-"`
}

// magicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseEnvelopeMessages]
type magicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseEnvelopeSuccess bool

const (
	MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseEnvelopeSuccessTrue MagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseEnvelopeSuccess = true
)
