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

// MagicIPSECTunnelService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewMagicIPSECTunnelService] method
// instead.
type MagicIPSECTunnelService struct {
	Options      []option.RequestOption
	PSKGenerates *MagicIPSECTunnelPSKGenerateService
}

// NewMagicIPSECTunnelService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMagicIPSECTunnelService(opts ...option.RequestOption) (r *MagicIPSECTunnelService) {
	r = &MagicIPSECTunnelService{}
	r.Options = opts
	r.PSKGenerates = NewMagicIPSECTunnelPSKGenerateService(opts...)
	return
}

// Creates new IPsec tunnels associated with an account. Use `?validate_only=true`
// as an optional query parameter to only run validation without persisting
// changes.
func (r *MagicIPSECTunnelService) New(ctx context.Context, accountIdentifier string, body MagicIPSECTunnelNewParams, opts ...option.RequestOption) (res *MagicIPSECTunnelNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicIPSECTunnelNewResponseEnvelope
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
func (r *MagicIPSECTunnelService) Update(ctx context.Context, accountIdentifier string, tunnelIdentifier string, body MagicIPSECTunnelUpdateParams, opts ...option.RequestOption) (res *MagicIPSECTunnelUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicIPSECTunnelUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists IPsec tunnels associated with an account.
func (r *MagicIPSECTunnelService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *MagicIPSECTunnelListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicIPSECTunnelListResponseEnvelope
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
func (r *MagicIPSECTunnelService) Delete(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *MagicIPSECTunnelDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicIPSECTunnelDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists details for a specific IPsec tunnel.
func (r *MagicIPSECTunnelService) Get(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *MagicIPSECTunnelGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicIPSECTunnelGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MagicIPSECTunnelNewResponse struct {
	IPSECTunnels []MagicIPSECTunnelNewResponseIPSECTunnel `json:"ipsec_tunnels"`
	JSON         magicIPSECTunnelNewResponseJSON          `json:"-"`
}

// magicIPSECTunnelNewResponseJSON contains the JSON metadata for the struct
// [MagicIPSECTunnelNewResponse]
type magicIPSECTunnelNewResponseJSON struct {
	IPSECTunnels apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MagicIPSECTunnelNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIPSECTunnelNewResponseIPSECTunnel struct {
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
	PSKMetadata MagicIPSECTunnelNewResponseIPSECTunnelsPSKMetadata `json:"psk_metadata"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection  bool                                                     `json:"replay_protection"`
	TunnelHealthCheck MagicIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheck `json:"tunnel_health_check"`
	JSON              magicIPSECTunnelNewResponseIPSECTunnelJSON               `json:"-"`
}

// magicIPSECTunnelNewResponseIPSECTunnelJSON contains the JSON metadata for the
// struct [MagicIPSECTunnelNewResponseIPSECTunnel]
type magicIPSECTunnelNewResponseIPSECTunnelJSON struct {
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

func (r *MagicIPSECTunnelNewResponseIPSECTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The PSK metadata that includes when the PSK was generated.
type MagicIPSECTunnelNewResponseIPSECTunnelsPSKMetadata struct {
	// The date and time the tunnel was last modified.
	LastGeneratedOn time.Time                                              `json:"last_generated_on" format:"date-time"`
	JSON            magicIPSECTunnelNewResponseIPSECTunnelsPSKMetadataJSON `json:"-"`
}

// magicIPSECTunnelNewResponseIPSECTunnelsPSKMetadataJSON contains the JSON
// metadata for the struct [MagicIPSECTunnelNewResponseIPSECTunnelsPSKMetadata]
type magicIPSECTunnelNewResponseIPSECTunnelsPSKMetadataJSON struct {
	LastGeneratedOn apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MagicIPSECTunnelNewResponseIPSECTunnelsPSKMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate MagicIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type MagicIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckType `json:"type"`
	JSON magicIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckJSON `json:"-"`
}

// magicIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckJSON contains the JSON
// metadata for the struct
// [MagicIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheck]
type magicIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckJSON struct {
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// How frequent the health check is run. The default value is `mid`.
type MagicIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckRate string

const (
	MagicIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckRateLow  MagicIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckRate = "low"
	MagicIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckRateMid  MagicIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckRate = "mid"
	MagicIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckRateHigh MagicIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckType string

const (
	MagicIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckTypeReply   MagicIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckType = "reply"
	MagicIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckTypeRequest MagicIPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheckType = "request"
)

type MagicIPSECTunnelUpdateResponse struct {
	Modified            bool                               `json:"modified"`
	ModifiedIPSECTunnel interface{}                        `json:"modified_ipsec_tunnel"`
	JSON                magicIPSECTunnelUpdateResponseJSON `json:"-"`
}

// magicIPSECTunnelUpdateResponseJSON contains the JSON metadata for the struct
// [MagicIPSECTunnelUpdateResponse]
type magicIPSECTunnelUpdateResponseJSON struct {
	Modified            apijson.Field
	ModifiedIPSECTunnel apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *MagicIPSECTunnelUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIPSECTunnelListResponse struct {
	IPSECTunnels []MagicIPSECTunnelListResponseIPSECTunnel `json:"ipsec_tunnels"`
	JSON         magicIPSECTunnelListResponseJSON          `json:"-"`
}

// magicIPSECTunnelListResponseJSON contains the JSON metadata for the struct
// [MagicIPSECTunnelListResponse]
type magicIPSECTunnelListResponseJSON struct {
	IPSECTunnels apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MagicIPSECTunnelListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIPSECTunnelListResponseIPSECTunnel struct {
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
	PSKMetadata MagicIPSECTunnelListResponseIPSECTunnelsPSKMetadata `json:"psk_metadata"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection  bool                                                      `json:"replay_protection"`
	TunnelHealthCheck MagicIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheck `json:"tunnel_health_check"`
	JSON              magicIPSECTunnelListResponseIPSECTunnelJSON               `json:"-"`
}

// magicIPSECTunnelListResponseIPSECTunnelJSON contains the JSON metadata for the
// struct [MagicIPSECTunnelListResponseIPSECTunnel]
type magicIPSECTunnelListResponseIPSECTunnelJSON struct {
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

func (r *MagicIPSECTunnelListResponseIPSECTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The PSK metadata that includes when the PSK was generated.
type MagicIPSECTunnelListResponseIPSECTunnelsPSKMetadata struct {
	// The date and time the tunnel was last modified.
	LastGeneratedOn time.Time                                               `json:"last_generated_on" format:"date-time"`
	JSON            magicIPSECTunnelListResponseIPSECTunnelsPSKMetadataJSON `json:"-"`
}

// magicIPSECTunnelListResponseIPSECTunnelsPSKMetadataJSON contains the JSON
// metadata for the struct [MagicIPSECTunnelListResponseIPSECTunnelsPSKMetadata]
type magicIPSECTunnelListResponseIPSECTunnelsPSKMetadataJSON struct {
	LastGeneratedOn apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MagicIPSECTunnelListResponseIPSECTunnelsPSKMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate MagicIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type MagicIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckType `json:"type"`
	JSON magicIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckJSON `json:"-"`
}

// magicIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckJSON contains the JSON
// metadata for the struct
// [MagicIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheck]
type magicIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckJSON struct {
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// How frequent the health check is run. The default value is `mid`.
type MagicIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckRate string

const (
	MagicIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckRateLow  MagicIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckRate = "low"
	MagicIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckRateMid  MagicIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckRate = "mid"
	MagicIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckRateHigh MagicIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckType string

const (
	MagicIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckTypeReply   MagicIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckType = "reply"
	MagicIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckTypeRequest MagicIPSECTunnelListResponseIPSECTunnelsTunnelHealthCheckType = "request"
)

type MagicIPSECTunnelDeleteResponse struct {
	Deleted            bool                               `json:"deleted"`
	DeletedIPSECTunnel interface{}                        `json:"deleted_ipsec_tunnel"`
	JSON               magicIPSECTunnelDeleteResponseJSON `json:"-"`
}

// magicIPSECTunnelDeleteResponseJSON contains the JSON metadata for the struct
// [MagicIPSECTunnelDeleteResponse]
type magicIPSECTunnelDeleteResponseJSON struct {
	Deleted            apijson.Field
	DeletedIPSECTunnel apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *MagicIPSECTunnelDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIPSECTunnelGetResponse struct {
	IPSECTunnel interface{}                     `json:"ipsec_tunnel"`
	JSON        magicIPSECTunnelGetResponseJSON `json:"-"`
}

// magicIPSECTunnelGetResponseJSON contains the JSON metadata for the struct
// [MagicIPSECTunnelGetResponse]
type magicIPSECTunnelGetResponseJSON struct {
	IPSECTunnel apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIPSECTunnelGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIPSECTunnelNewParams struct {
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
	Description param.Field[string]                               `json:"description"`
	HealthCheck param.Field[MagicIPSECTunnelNewParamsHealthCheck] `json:"health_check"`
	// A randomly generated or provided string for use in the IPsec tunnel.
	PSK param.Field[string] `json:"psk"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection param.Field[bool] `json:"replay_protection"`
}

func (r MagicIPSECTunnelNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagicIPSECTunnelNewParamsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
	// in health_check is ignored as the interface_address is used to send traffic into
	// the tunnel.
	Direction param.Field[MagicIPSECTunnelNewParamsHealthCheckDirection] `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled param.Field[bool] `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate param.Field[MagicIPSECTunnelNewParamsHealthCheckRate] `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target.
	Target param.Field[string] `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type param.Field[MagicIPSECTunnelNewParamsHealthCheckType] `json:"type"`
}

func (r MagicIPSECTunnelNewParamsHealthCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
// in health_check is ignored as the interface_address is used to send traffic into
// the tunnel.
type MagicIPSECTunnelNewParamsHealthCheckDirection string

const (
	MagicIPSECTunnelNewParamsHealthCheckDirectionUnidirectional MagicIPSECTunnelNewParamsHealthCheckDirection = "unidirectional"
	MagicIPSECTunnelNewParamsHealthCheckDirectionBidirectional  MagicIPSECTunnelNewParamsHealthCheckDirection = "bidirectional"
)

// How frequent the health check is run. The default value is `mid`.
type MagicIPSECTunnelNewParamsHealthCheckRate string

const (
	MagicIPSECTunnelNewParamsHealthCheckRateLow  MagicIPSECTunnelNewParamsHealthCheckRate = "low"
	MagicIPSECTunnelNewParamsHealthCheckRateMid  MagicIPSECTunnelNewParamsHealthCheckRate = "mid"
	MagicIPSECTunnelNewParamsHealthCheckRateHigh MagicIPSECTunnelNewParamsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicIPSECTunnelNewParamsHealthCheckType string

const (
	MagicIPSECTunnelNewParamsHealthCheckTypeReply   MagicIPSECTunnelNewParamsHealthCheckType = "reply"
	MagicIPSECTunnelNewParamsHealthCheckTypeRequest MagicIPSECTunnelNewParamsHealthCheckType = "request"
)

type MagicIPSECTunnelNewResponseEnvelope struct {
	Errors   []MagicIPSECTunnelNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicIPSECTunnelNewResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicIPSECTunnelNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicIPSECTunnelNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicIPSECTunnelNewResponseEnvelopeJSON    `json:"-"`
}

// magicIPSECTunnelNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [MagicIPSECTunnelNewResponseEnvelope]
type magicIPSECTunnelNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIPSECTunnelNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIPSECTunnelNewResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    magicIPSECTunnelNewResponseEnvelopeErrorsJSON `json:"-"`
}

// magicIPSECTunnelNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MagicIPSECTunnelNewResponseEnvelopeErrors]
type magicIPSECTunnelNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIPSECTunnelNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIPSECTunnelNewResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    magicIPSECTunnelNewResponseEnvelopeMessagesJSON `json:"-"`
}

// magicIPSECTunnelNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [MagicIPSECTunnelNewResponseEnvelopeMessages]
type magicIPSECTunnelNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIPSECTunnelNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicIPSECTunnelNewResponseEnvelopeSuccess bool

const (
	MagicIPSECTunnelNewResponseEnvelopeSuccessTrue MagicIPSECTunnelNewResponseEnvelopeSuccess = true
)

type MagicIPSECTunnelUpdateParams struct {
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
	Description param.Field[string]                                  `json:"description"`
	HealthCheck param.Field[MagicIPSECTunnelUpdateParamsHealthCheck] `json:"health_check"`
	// A randomly generated or provided string for use in the IPsec tunnel.
	PSK param.Field[string] `json:"psk"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection param.Field[bool] `json:"replay_protection"`
}

func (r MagicIPSECTunnelUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagicIPSECTunnelUpdateParamsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
	// in health_check is ignored as the interface_address is used to send traffic into
	// the tunnel.
	Direction param.Field[MagicIPSECTunnelUpdateParamsHealthCheckDirection] `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled param.Field[bool] `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate param.Field[MagicIPSECTunnelUpdateParamsHealthCheckRate] `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target.
	Target param.Field[string] `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type param.Field[MagicIPSECTunnelUpdateParamsHealthCheckType] `json:"type"`
}

func (r MagicIPSECTunnelUpdateParamsHealthCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
// in health_check is ignored as the interface_address is used to send traffic into
// the tunnel.
type MagicIPSECTunnelUpdateParamsHealthCheckDirection string

const (
	MagicIPSECTunnelUpdateParamsHealthCheckDirectionUnidirectional MagicIPSECTunnelUpdateParamsHealthCheckDirection = "unidirectional"
	MagicIPSECTunnelUpdateParamsHealthCheckDirectionBidirectional  MagicIPSECTunnelUpdateParamsHealthCheckDirection = "bidirectional"
)

// How frequent the health check is run. The default value is `mid`.
type MagicIPSECTunnelUpdateParamsHealthCheckRate string

const (
	MagicIPSECTunnelUpdateParamsHealthCheckRateLow  MagicIPSECTunnelUpdateParamsHealthCheckRate = "low"
	MagicIPSECTunnelUpdateParamsHealthCheckRateMid  MagicIPSECTunnelUpdateParamsHealthCheckRate = "mid"
	MagicIPSECTunnelUpdateParamsHealthCheckRateHigh MagicIPSECTunnelUpdateParamsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicIPSECTunnelUpdateParamsHealthCheckType string

const (
	MagicIPSECTunnelUpdateParamsHealthCheckTypeReply   MagicIPSECTunnelUpdateParamsHealthCheckType = "reply"
	MagicIPSECTunnelUpdateParamsHealthCheckTypeRequest MagicIPSECTunnelUpdateParamsHealthCheckType = "request"
)

type MagicIPSECTunnelUpdateResponseEnvelope struct {
	Errors   []MagicIPSECTunnelUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicIPSECTunnelUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicIPSECTunnelUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicIPSECTunnelUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicIPSECTunnelUpdateResponseEnvelopeJSON    `json:"-"`
}

// magicIPSECTunnelUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [MagicIPSECTunnelUpdateResponseEnvelope]
type magicIPSECTunnelUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIPSECTunnelUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIPSECTunnelUpdateResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    magicIPSECTunnelUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// magicIPSECTunnelUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [MagicIPSECTunnelUpdateResponseEnvelopeErrors]
type magicIPSECTunnelUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIPSECTunnelUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIPSECTunnelUpdateResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    magicIPSECTunnelUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// magicIPSECTunnelUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [MagicIPSECTunnelUpdateResponseEnvelopeMessages]
type magicIPSECTunnelUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIPSECTunnelUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicIPSECTunnelUpdateResponseEnvelopeSuccess bool

const (
	MagicIPSECTunnelUpdateResponseEnvelopeSuccessTrue MagicIPSECTunnelUpdateResponseEnvelopeSuccess = true
)

type MagicIPSECTunnelListResponseEnvelope struct {
	Errors   []MagicIPSECTunnelListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicIPSECTunnelListResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicIPSECTunnelListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicIPSECTunnelListResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicIPSECTunnelListResponseEnvelopeJSON    `json:"-"`
}

// magicIPSECTunnelListResponseEnvelopeJSON contains the JSON metadata for the
// struct [MagicIPSECTunnelListResponseEnvelope]
type magicIPSECTunnelListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIPSECTunnelListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIPSECTunnelListResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    magicIPSECTunnelListResponseEnvelopeErrorsJSON `json:"-"`
}

// magicIPSECTunnelListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [MagicIPSECTunnelListResponseEnvelopeErrors]
type magicIPSECTunnelListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIPSECTunnelListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIPSECTunnelListResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    magicIPSECTunnelListResponseEnvelopeMessagesJSON `json:"-"`
}

// magicIPSECTunnelListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [MagicIPSECTunnelListResponseEnvelopeMessages]
type magicIPSECTunnelListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIPSECTunnelListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicIPSECTunnelListResponseEnvelopeSuccess bool

const (
	MagicIPSECTunnelListResponseEnvelopeSuccessTrue MagicIPSECTunnelListResponseEnvelopeSuccess = true
)

type MagicIPSECTunnelDeleteResponseEnvelope struct {
	Errors   []MagicIPSECTunnelDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicIPSECTunnelDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicIPSECTunnelDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicIPSECTunnelDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicIPSECTunnelDeleteResponseEnvelopeJSON    `json:"-"`
}

// magicIPSECTunnelDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [MagicIPSECTunnelDeleteResponseEnvelope]
type magicIPSECTunnelDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIPSECTunnelDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIPSECTunnelDeleteResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    magicIPSECTunnelDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// magicIPSECTunnelDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [MagicIPSECTunnelDeleteResponseEnvelopeErrors]
type magicIPSECTunnelDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIPSECTunnelDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIPSECTunnelDeleteResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    magicIPSECTunnelDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// magicIPSECTunnelDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [MagicIPSECTunnelDeleteResponseEnvelopeMessages]
type magicIPSECTunnelDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIPSECTunnelDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicIPSECTunnelDeleteResponseEnvelopeSuccess bool

const (
	MagicIPSECTunnelDeleteResponseEnvelopeSuccessTrue MagicIPSECTunnelDeleteResponseEnvelopeSuccess = true
)

type MagicIPSECTunnelGetResponseEnvelope struct {
	Errors   []MagicIPSECTunnelGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicIPSECTunnelGetResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicIPSECTunnelGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicIPSECTunnelGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicIPSECTunnelGetResponseEnvelopeJSON    `json:"-"`
}

// magicIPSECTunnelGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [MagicIPSECTunnelGetResponseEnvelope]
type magicIPSECTunnelGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIPSECTunnelGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIPSECTunnelGetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    magicIPSECTunnelGetResponseEnvelopeErrorsJSON `json:"-"`
}

// magicIPSECTunnelGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MagicIPSECTunnelGetResponseEnvelopeErrors]
type magicIPSECTunnelGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIPSECTunnelGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIPSECTunnelGetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    magicIPSECTunnelGetResponseEnvelopeMessagesJSON `json:"-"`
}

// magicIPSECTunnelGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [MagicIPSECTunnelGetResponseEnvelopeMessages]
type magicIPSECTunnelGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIPSECTunnelGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicIPSECTunnelGetResponseEnvelopeSuccess bool

const (
	MagicIPSECTunnelGetResponseEnvelopeSuccessTrue MagicIPSECTunnelGetResponseEnvelopeSuccess = true
)
