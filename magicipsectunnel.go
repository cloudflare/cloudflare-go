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

// Creates new IPsec tunnels associated with an account. Use `?validate_only=true`
// as an optional query parameter to only run validation without persisting
// changes.
func (r *MagicIpsecTunnelService) New(ctx context.Context, accountIdentifier string, body MagicIpsecTunnelNewParams, opts ...option.RequestOption) (res *MagicIpsecTunnelNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicIpsecTunnelNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists IPsec tunnels associated with an account.
func (r *MagicIpsecTunnelService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *MagicIpsecTunnelListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicIpsecTunnelListResponseEnvelope
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

// Updates a specific IPsec tunnel associated with an account. Use
// `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes.
func (r *MagicIpsecTunnelService) Replace(ctx context.Context, accountIdentifier string, tunnelIdentifier string, body MagicIpsecTunnelReplaceParams, opts ...option.RequestOption) (res *MagicIpsecTunnelReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicIpsecTunnelReplaceResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MagicIpsecTunnelNewResponse struct {
	IpsecTunnels []MagicIpsecTunnelNewResponseIpsecTunnel `json:"ipsec_tunnels"`
	JSON         magicIpsecTunnelNewResponseJSON          `json:"-"`
}

// magicIpsecTunnelNewResponseJSON contains the JSON metadata for the struct
// [MagicIpsecTunnelNewResponse]
type magicIpsecTunnelNewResponseJSON struct {
	IpsecTunnels apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MagicIpsecTunnelNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelNewResponseIpsecTunnel struct {
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
	PskMetadata MagicIpsecTunnelNewResponseIpsecTunnelsPskMetadata `json:"psk_metadata"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection  bool                                                     `json:"replay_protection"`
	TunnelHealthCheck MagicIpsecTunnelNewResponseIpsecTunnelsTunnelHealthCheck `json:"tunnel_health_check"`
	JSON              magicIpsecTunnelNewResponseIpsecTunnelJSON               `json:"-"`
}

// magicIpsecTunnelNewResponseIpsecTunnelJSON contains the JSON metadata for the
// struct [MagicIpsecTunnelNewResponseIpsecTunnel]
type magicIpsecTunnelNewResponseIpsecTunnelJSON struct {
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

func (r *MagicIpsecTunnelNewResponseIpsecTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The PSK metadata that includes when the PSK was generated.
type MagicIpsecTunnelNewResponseIpsecTunnelsPskMetadata struct {
	// The date and time the tunnel was last modified.
	LastGeneratedOn time.Time                                              `json:"last_generated_on" format:"date-time"`
	JSON            magicIpsecTunnelNewResponseIpsecTunnelsPskMetadataJSON `json:"-"`
}

// magicIpsecTunnelNewResponseIpsecTunnelsPskMetadataJSON contains the JSON
// metadata for the struct [MagicIpsecTunnelNewResponseIpsecTunnelsPskMetadata]
type magicIpsecTunnelNewResponseIpsecTunnelsPskMetadataJSON struct {
	LastGeneratedOn apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MagicIpsecTunnelNewResponseIpsecTunnelsPskMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelNewResponseIpsecTunnelsTunnelHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate MagicIpsecTunnelNewResponseIpsecTunnelsTunnelHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type MagicIpsecTunnelNewResponseIpsecTunnelsTunnelHealthCheckType `json:"type"`
	JSON magicIpsecTunnelNewResponseIpsecTunnelsTunnelHealthCheckJSON `json:"-"`
}

// magicIpsecTunnelNewResponseIpsecTunnelsTunnelHealthCheckJSON contains the JSON
// metadata for the struct
// [MagicIpsecTunnelNewResponseIpsecTunnelsTunnelHealthCheck]
type magicIpsecTunnelNewResponseIpsecTunnelsTunnelHealthCheckJSON struct {
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelNewResponseIpsecTunnelsTunnelHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// How frequent the health check is run. The default value is `mid`.
type MagicIpsecTunnelNewResponseIpsecTunnelsTunnelHealthCheckRate string

const (
	MagicIpsecTunnelNewResponseIpsecTunnelsTunnelHealthCheckRateLow  MagicIpsecTunnelNewResponseIpsecTunnelsTunnelHealthCheckRate = "low"
	MagicIpsecTunnelNewResponseIpsecTunnelsTunnelHealthCheckRateMid  MagicIpsecTunnelNewResponseIpsecTunnelsTunnelHealthCheckRate = "mid"
	MagicIpsecTunnelNewResponseIpsecTunnelsTunnelHealthCheckRateHigh MagicIpsecTunnelNewResponseIpsecTunnelsTunnelHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicIpsecTunnelNewResponseIpsecTunnelsTunnelHealthCheckType string

const (
	MagicIpsecTunnelNewResponseIpsecTunnelsTunnelHealthCheckTypeReply   MagicIpsecTunnelNewResponseIpsecTunnelsTunnelHealthCheckType = "reply"
	MagicIpsecTunnelNewResponseIpsecTunnelsTunnelHealthCheckTypeRequest MagicIpsecTunnelNewResponseIpsecTunnelsTunnelHealthCheckType = "request"
)

type MagicIpsecTunnelListResponse struct {
	IpsecTunnels []MagicIpsecTunnelListResponseIpsecTunnel `json:"ipsec_tunnels"`
	JSON         magicIpsecTunnelListResponseJSON          `json:"-"`
}

// magicIpsecTunnelListResponseJSON contains the JSON metadata for the struct
// [MagicIpsecTunnelListResponse]
type magicIpsecTunnelListResponseJSON struct {
	IpsecTunnels apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MagicIpsecTunnelListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelListResponseIpsecTunnel struct {
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
	PskMetadata MagicIpsecTunnelListResponseIpsecTunnelsPskMetadata `json:"psk_metadata"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection  bool                                                      `json:"replay_protection"`
	TunnelHealthCheck MagicIpsecTunnelListResponseIpsecTunnelsTunnelHealthCheck `json:"tunnel_health_check"`
	JSON              magicIpsecTunnelListResponseIpsecTunnelJSON               `json:"-"`
}

// magicIpsecTunnelListResponseIpsecTunnelJSON contains the JSON metadata for the
// struct [MagicIpsecTunnelListResponseIpsecTunnel]
type magicIpsecTunnelListResponseIpsecTunnelJSON struct {
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

func (r *MagicIpsecTunnelListResponseIpsecTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The PSK metadata that includes when the PSK was generated.
type MagicIpsecTunnelListResponseIpsecTunnelsPskMetadata struct {
	// The date and time the tunnel was last modified.
	LastGeneratedOn time.Time                                               `json:"last_generated_on" format:"date-time"`
	JSON            magicIpsecTunnelListResponseIpsecTunnelsPskMetadataJSON `json:"-"`
}

// magicIpsecTunnelListResponseIpsecTunnelsPskMetadataJSON contains the JSON
// metadata for the struct [MagicIpsecTunnelListResponseIpsecTunnelsPskMetadata]
type magicIpsecTunnelListResponseIpsecTunnelsPskMetadataJSON struct {
	LastGeneratedOn apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MagicIpsecTunnelListResponseIpsecTunnelsPskMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelListResponseIpsecTunnelsTunnelHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate MagicIpsecTunnelListResponseIpsecTunnelsTunnelHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type MagicIpsecTunnelListResponseIpsecTunnelsTunnelHealthCheckType `json:"type"`
	JSON magicIpsecTunnelListResponseIpsecTunnelsTunnelHealthCheckJSON `json:"-"`
}

// magicIpsecTunnelListResponseIpsecTunnelsTunnelHealthCheckJSON contains the JSON
// metadata for the struct
// [MagicIpsecTunnelListResponseIpsecTunnelsTunnelHealthCheck]
type magicIpsecTunnelListResponseIpsecTunnelsTunnelHealthCheckJSON struct {
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelListResponseIpsecTunnelsTunnelHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// How frequent the health check is run. The default value is `mid`.
type MagicIpsecTunnelListResponseIpsecTunnelsTunnelHealthCheckRate string

const (
	MagicIpsecTunnelListResponseIpsecTunnelsTunnelHealthCheckRateLow  MagicIpsecTunnelListResponseIpsecTunnelsTunnelHealthCheckRate = "low"
	MagicIpsecTunnelListResponseIpsecTunnelsTunnelHealthCheckRateMid  MagicIpsecTunnelListResponseIpsecTunnelsTunnelHealthCheckRate = "mid"
	MagicIpsecTunnelListResponseIpsecTunnelsTunnelHealthCheckRateHigh MagicIpsecTunnelListResponseIpsecTunnelsTunnelHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicIpsecTunnelListResponseIpsecTunnelsTunnelHealthCheckType string

const (
	MagicIpsecTunnelListResponseIpsecTunnelsTunnelHealthCheckTypeReply   MagicIpsecTunnelListResponseIpsecTunnelsTunnelHealthCheckType = "reply"
	MagicIpsecTunnelListResponseIpsecTunnelsTunnelHealthCheckTypeRequest MagicIpsecTunnelListResponseIpsecTunnelsTunnelHealthCheckType = "request"
)

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

type MagicIpsecTunnelReplaceResponse struct {
	Modified            bool                                `json:"modified"`
	ModifiedIpsecTunnel interface{}                         `json:"modified_ipsec_tunnel"`
	JSON                magicIpsecTunnelReplaceResponseJSON `json:"-"`
}

// magicIpsecTunnelReplaceResponseJSON contains the JSON metadata for the struct
// [MagicIpsecTunnelReplaceResponse]
type magicIpsecTunnelReplaceResponseJSON struct {
	Modified            apijson.Field
	ModifiedIpsecTunnel apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *MagicIpsecTunnelReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelNewParams struct {
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

func (r MagicIpsecTunnelNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagicIpsecTunnelNewResponseEnvelope struct {
	Errors   []MagicIpsecTunnelNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicIpsecTunnelNewResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicIpsecTunnelNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicIpsecTunnelNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicIpsecTunnelNewResponseEnvelopeJSON    `json:"-"`
}

// magicIpsecTunnelNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [MagicIpsecTunnelNewResponseEnvelope]
type magicIpsecTunnelNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelNewResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    magicIpsecTunnelNewResponseEnvelopeErrorsJSON `json:"-"`
}

// magicIpsecTunnelNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MagicIpsecTunnelNewResponseEnvelopeErrors]
type magicIpsecTunnelNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelNewResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    magicIpsecTunnelNewResponseEnvelopeMessagesJSON `json:"-"`
}

// magicIpsecTunnelNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [MagicIpsecTunnelNewResponseEnvelopeMessages]
type magicIpsecTunnelNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicIpsecTunnelNewResponseEnvelopeSuccess bool

const (
	MagicIpsecTunnelNewResponseEnvelopeSuccessTrue MagicIpsecTunnelNewResponseEnvelopeSuccess = true
)

type MagicIpsecTunnelListResponseEnvelope struct {
	Errors   []MagicIpsecTunnelListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicIpsecTunnelListResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicIpsecTunnelListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicIpsecTunnelListResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicIpsecTunnelListResponseEnvelopeJSON    `json:"-"`
}

// magicIpsecTunnelListResponseEnvelopeJSON contains the JSON metadata for the
// struct [MagicIpsecTunnelListResponseEnvelope]
type magicIpsecTunnelListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelListResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    magicIpsecTunnelListResponseEnvelopeErrorsJSON `json:"-"`
}

// magicIpsecTunnelListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [MagicIpsecTunnelListResponseEnvelopeErrors]
type magicIpsecTunnelListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelListResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    magicIpsecTunnelListResponseEnvelopeMessagesJSON `json:"-"`
}

// magicIpsecTunnelListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [MagicIpsecTunnelListResponseEnvelopeMessages]
type magicIpsecTunnelListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicIpsecTunnelListResponseEnvelopeSuccess bool

const (
	MagicIpsecTunnelListResponseEnvelopeSuccessTrue MagicIpsecTunnelListResponseEnvelopeSuccess = true
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

type MagicIpsecTunnelReplaceParams struct {
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

func (r MagicIpsecTunnelReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagicIpsecTunnelReplaceResponseEnvelope struct {
	Errors   []MagicIpsecTunnelReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicIpsecTunnelReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicIpsecTunnelReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicIpsecTunnelReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicIpsecTunnelReplaceResponseEnvelopeJSON    `json:"-"`
}

// magicIpsecTunnelReplaceResponseEnvelopeJSON contains the JSON metadata for the
// struct [MagicIpsecTunnelReplaceResponseEnvelope]
type magicIpsecTunnelReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelReplaceResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    magicIpsecTunnelReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// magicIpsecTunnelReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [MagicIpsecTunnelReplaceResponseEnvelopeErrors]
type magicIpsecTunnelReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelReplaceResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    magicIpsecTunnelReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// magicIpsecTunnelReplaceResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [MagicIpsecTunnelReplaceResponseEnvelopeMessages]
type magicIpsecTunnelReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicIpsecTunnelReplaceResponseEnvelopeSuccess bool

const (
	MagicIpsecTunnelReplaceResponseEnvelopeSuccessTrue MagicIpsecTunnelReplaceResponseEnvelopeSuccess = true
)
