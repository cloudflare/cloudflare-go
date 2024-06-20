// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// IPSECTunnelService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIPSECTunnelService] method instead.
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
func (r *IPSECTunnelService) New(ctx context.Context, params IPSECTunnelNewParams, opts ...option.RequestOption) (res *IPSECTunnelNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IPSECTunnelNewResponseEnvelope
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists IPsec tunnels associated with an account.
func (r *IPSECTunnelService) List(ctx context.Context, query IPSECTunnelListParams, opts ...option.RequestOption) (res *IPSECTunnelListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IPSECTunnelListResponseEnvelope
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// The PSK metadata that includes when the PSK was generated.
type PSKMetadata struct {
	// The date and time the tunnel was last modified.
	LastGeneratedOn time.Time       `json:"last_generated_on" format:"date-time"`
	JSON            pskMetadataJSON `json:"-"`
}

// pskMetadataJSON contains the JSON metadata for the struct [PSKMetadata]
type pskMetadataJSON struct {
	LastGeneratedOn apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PSKMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pskMetadataJSON) RawJSON() string {
	return r.raw
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
	// The IP address assigned to the customer side of the IPsec tunnel. Not required,
	// but must be set for proactive traceroutes to work.
	CustomerEndpoint string `json:"customer_endpoint"`
	// An optional description forthe IPsec tunnel.
	Description string `json:"description"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The PSK metadata that includes when the PSK was generated.
	PSKMetadata PSKMetadata `json:"psk_metadata"`
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

type IPSECTunnelNewResponseIPSECTunnelsTunnelHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate HealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type HealthCheckType                                         `json:"type"`
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
	// The IP address assigned to the customer side of the IPsec tunnel. Not required,
	// but must be set for proactive traceroutes to work.
	CustomerEndpoint string `json:"customer_endpoint"`
	// An optional description forthe IPsec tunnel.
	Description string `json:"description"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The PSK metadata that includes when the PSK was generated.
	PSKMetadata PSKMetadata `json:"psk_metadata"`
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

type IPSECTunnelListResponseIPSECTunnelsTunnelHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate HealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type HealthCheckType                                          `json:"type"`
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

type IPSECTunnelNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The IP address assigned to the Cloudflare side of the IPsec tunnel.
	CloudflareEndpoint param.Field[string] `json:"cloudflare_endpoint,required"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress param.Field[string] `json:"interface_address,required"`
	// The name of the IPsec tunnel. The name cannot share a name with other tunnels.
	Name param.Field[string] `json:"name,required"`
	// The IP address assigned to the customer side of the IPsec tunnel. Not required,
	// but must be set for proactive traceroutes to work.
	CustomerEndpoint param.Field[string] `json:"customer_endpoint"`
	// An optional description forthe IPsec tunnel.
	Description param.Field[string]           `json:"description"`
	HealthCheck param.Field[HealthCheckParam] `json:"health_check"`
	// A randomly generated or provided string for use in the IPsec tunnel.
	PSK param.Field[string] `json:"psk"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection param.Field[bool] `json:"replay_protection"`
}

func (r IPSECTunnelNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IPSECTunnelNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo  `json:"errors,required"`
	Messages []shared.ResponseInfo  `json:"messages,required"`
	Result   IPSECTunnelNewResponse `json:"result,required"`
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

// Whether the API call was successful
type IPSECTunnelNewResponseEnvelopeSuccess bool

const (
	IPSECTunnelNewResponseEnvelopeSuccessTrue IPSECTunnelNewResponseEnvelopeSuccess = true
)

func (r IPSECTunnelNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IPSECTunnelNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IPSECTunnelListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type IPSECTunnelListResponseEnvelope struct {
	Errors   []shared.ResponseInfo   `json:"errors,required"`
	Messages []shared.ResponseInfo   `json:"messages,required"`
	Result   IPSECTunnelListResponse `json:"result,required"`
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

// Whether the API call was successful
type IPSECTunnelListResponseEnvelopeSuccess bool

const (
	IPSECTunnelListResponseEnvelopeSuccessTrue IPSECTunnelListResponseEnvelopeSuccess = true
)

func (r IPSECTunnelListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IPSECTunnelListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
