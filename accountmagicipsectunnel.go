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

// AccountMagicIpsecTunnelService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountMagicIpsecTunnelService] method instead.
type AccountMagicIpsecTunnelService struct {
	Options      []option.RequestOption
	PskGenerates *AccountMagicIpsecTunnelPskGenerateService
}

// NewAccountMagicIpsecTunnelService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountMagicIpsecTunnelService(opts ...option.RequestOption) (r *AccountMagicIpsecTunnelService) {
	r = &AccountMagicIpsecTunnelService{}
	r.Options = opts
	r.PskGenerates = NewAccountMagicIpsecTunnelPskGenerateService(opts...)
	return
}

// Lists details for a specific IPsec tunnel.
func (r *AccountMagicIpsecTunnelService) Get(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *AccountMagicIpsecTunnelGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a specific IPsec tunnel associated with an account. Use
// `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes.
func (r *AccountMagicIpsecTunnelService) Update(ctx context.Context, accountIdentifier string, tunnelIdentifier string, body AccountMagicIpsecTunnelUpdateParams, opts ...option.RequestOption) (res *AccountMagicIpsecTunnelUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Disables and removes a specific static IPsec Tunnel associated with an account.
// Use `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes.
func (r *AccountMagicIpsecTunnelService) Delete(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *AccountMagicIpsecTunnelDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates new IPsec tunnels associated with an account. Use `?validate_only=true`
// as an optional query parameter to only run validation without persisting
// changes.
func (r *AccountMagicIpsecTunnelService) MagicIPsecTunnelsNewIPsecTunnels(ctx context.Context, accountIdentifier string, body AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsParams, opts ...option.RequestOption) (res *AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists IPsec tunnels associated with an account.
func (r *AccountMagicIpsecTunnelService) MagicIPsecTunnelsListIPsecTunnels(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update multiple IPsec tunnels associated with an account. Use
// `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes.
func (r *AccountMagicIpsecTunnelService) MagicIPsecTunnelsUpdateMultipleIPsecTunnels(ctx context.Context, accountIdentifier string, body AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsParams, opts ...option.RequestOption) (res *AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccountMagicIpsecTunnelGetResponse struct {
	Errors   []AccountMagicIpsecTunnelGetResponseError   `json:"errors"`
	Messages []AccountMagicIpsecTunnelGetResponseMessage `json:"messages"`
	Result   AccountMagicIpsecTunnelGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMagicIpsecTunnelGetResponseSuccess `json:"success"`
	JSON    accountMagicIpsecTunnelGetResponseJSON    `json:"-"`
}

// accountMagicIpsecTunnelGetResponseJSON contains the JSON metadata for the struct
// [AccountMagicIpsecTunnelGetResponse]
type accountMagicIpsecTunnelGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicIpsecTunnelGetResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountMagicIpsecTunnelGetResponseErrorJSON `json:"-"`
}

// accountMagicIpsecTunnelGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountMagicIpsecTunnelGetResponseError]
type accountMagicIpsecTunnelGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicIpsecTunnelGetResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountMagicIpsecTunnelGetResponseMessageJSON `json:"-"`
}

// accountMagicIpsecTunnelGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountMagicIpsecTunnelGetResponseMessage]
type accountMagicIpsecTunnelGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicIpsecTunnelGetResponseResult struct {
	IpsecTunnel interface{}                                  `json:"ipsec_tunnel"`
	JSON        accountMagicIpsecTunnelGetResponseResultJSON `json:"-"`
}

// accountMagicIpsecTunnelGetResponseResultJSON contains the JSON metadata for the
// struct [AccountMagicIpsecTunnelGetResponseResult]
type accountMagicIpsecTunnelGetResponseResultJSON struct {
	IpsecTunnel apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMagicIpsecTunnelGetResponseSuccess bool

const (
	AccountMagicIpsecTunnelGetResponseSuccessTrue AccountMagicIpsecTunnelGetResponseSuccess = true
)

type AccountMagicIpsecTunnelUpdateResponse struct {
	Errors   []AccountMagicIpsecTunnelUpdateResponseError   `json:"errors"`
	Messages []AccountMagicIpsecTunnelUpdateResponseMessage `json:"messages"`
	Result   AccountMagicIpsecTunnelUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMagicIpsecTunnelUpdateResponseSuccess `json:"success"`
	JSON    accountMagicIpsecTunnelUpdateResponseJSON    `json:"-"`
}

// accountMagicIpsecTunnelUpdateResponseJSON contains the JSON metadata for the
// struct [AccountMagicIpsecTunnelUpdateResponse]
type accountMagicIpsecTunnelUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicIpsecTunnelUpdateResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountMagicIpsecTunnelUpdateResponseErrorJSON `json:"-"`
}

// accountMagicIpsecTunnelUpdateResponseErrorJSON contains the JSON metadata for
// the struct [AccountMagicIpsecTunnelUpdateResponseError]
type accountMagicIpsecTunnelUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicIpsecTunnelUpdateResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accountMagicIpsecTunnelUpdateResponseMessageJSON `json:"-"`
}

// accountMagicIpsecTunnelUpdateResponseMessageJSON contains the JSON metadata for
// the struct [AccountMagicIpsecTunnelUpdateResponseMessage]
type accountMagicIpsecTunnelUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicIpsecTunnelUpdateResponseResult struct {
	Modified            bool                                            `json:"modified"`
	ModifiedIpsecTunnel interface{}                                     `json:"modified_ipsec_tunnel"`
	JSON                accountMagicIpsecTunnelUpdateResponseResultJSON `json:"-"`
}

// accountMagicIpsecTunnelUpdateResponseResultJSON contains the JSON metadata for
// the struct [AccountMagicIpsecTunnelUpdateResponseResult]
type accountMagicIpsecTunnelUpdateResponseResultJSON struct {
	Modified            apijson.Field
	ModifiedIpsecTunnel apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMagicIpsecTunnelUpdateResponseSuccess bool

const (
	AccountMagicIpsecTunnelUpdateResponseSuccessTrue AccountMagicIpsecTunnelUpdateResponseSuccess = true
)

type AccountMagicIpsecTunnelDeleteResponse struct {
	Errors   []AccountMagicIpsecTunnelDeleteResponseError   `json:"errors"`
	Messages []AccountMagicIpsecTunnelDeleteResponseMessage `json:"messages"`
	Result   AccountMagicIpsecTunnelDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMagicIpsecTunnelDeleteResponseSuccess `json:"success"`
	JSON    accountMagicIpsecTunnelDeleteResponseJSON    `json:"-"`
}

// accountMagicIpsecTunnelDeleteResponseJSON contains the JSON metadata for the
// struct [AccountMagicIpsecTunnelDeleteResponse]
type accountMagicIpsecTunnelDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicIpsecTunnelDeleteResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountMagicIpsecTunnelDeleteResponseErrorJSON `json:"-"`
}

// accountMagicIpsecTunnelDeleteResponseErrorJSON contains the JSON metadata for
// the struct [AccountMagicIpsecTunnelDeleteResponseError]
type accountMagicIpsecTunnelDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicIpsecTunnelDeleteResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accountMagicIpsecTunnelDeleteResponseMessageJSON `json:"-"`
}

// accountMagicIpsecTunnelDeleteResponseMessageJSON contains the JSON metadata for
// the struct [AccountMagicIpsecTunnelDeleteResponseMessage]
type accountMagicIpsecTunnelDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicIpsecTunnelDeleteResponseResult struct {
	Deleted            bool                                            `json:"deleted"`
	DeletedIpsecTunnel interface{}                                     `json:"deleted_ipsec_tunnel"`
	JSON               accountMagicIpsecTunnelDeleteResponseResultJSON `json:"-"`
}

// accountMagicIpsecTunnelDeleteResponseResultJSON contains the JSON metadata for
// the struct [AccountMagicIpsecTunnelDeleteResponseResult]
type accountMagicIpsecTunnelDeleteResponseResultJSON struct {
	Deleted            apijson.Field
	DeletedIpsecTunnel apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMagicIpsecTunnelDeleteResponseSuccess bool

const (
	AccountMagicIpsecTunnelDeleteResponseSuccessTrue AccountMagicIpsecTunnelDeleteResponseSuccess = true
)

type AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponse struct {
	Errors   []AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseError   `json:"errors"`
	Messages []AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseMessage `json:"messages"`
	Result   AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseSuccess `json:"success"`
	JSON    accountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseJSON    `json:"-"`
}

// accountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseJSON contains the
// JSON metadata for the struct
// [AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponse]
type accountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseError struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    accountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseErrorJSON `json:"-"`
}

// accountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseError]
type accountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseMessage struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    accountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseMessageJSON `json:"-"`
}

// accountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseMessage]
type accountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResult struct {
	IpsecTunnels []AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnel `json:"ipsec_tunnels"`
	JSON         accountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultJSON          `json:"-"`
}

// accountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResult]
type accountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultJSON struct {
	IpsecTunnels apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnel struct {
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
	PskMetadata AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnelsPskMetadata `json:"psk_metadata"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection  bool                                                                                               `json:"replay_protection"`
	TunnelHealthCheck AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheck `json:"tunnel_health_check"`
	JSON              accountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnelJSON               `json:"-"`
}

// accountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnelJSON
// contains the JSON metadata for the struct
// [AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnel]
type accountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnelJSON struct {
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

func (r *AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The PSK metadata that includes when the PSK was generated.
type AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnelsPskMetadata struct {
	// The date and time the tunnel was last modified.
	LastGeneratedOn time.Time                                                                                        `json:"last_generated_on" format:"date-time"`
	JSON            accountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnelsPskMetadataJSON `json:"-"`
}

// accountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnelsPskMetadataJSON
// contains the JSON metadata for the struct
// [AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnelsPskMetadata]
type accountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnelsPskMetadataJSON struct {
	LastGeneratedOn apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnelsPskMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckType `json:"type"`
	JSON accountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckJSON `json:"-"`
}

// accountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckJSON
// contains the JSON metadata for the struct
// [AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheck]
type accountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckJSON struct {
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// How frequent the health check is run. The default value is `mid`.
type AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckRate string

const (
	AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckRateLow  AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckRate = "low"
	AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckRateMid  AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckRate = "mid"
	AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckRateHigh AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckType string

const (
	AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckTypeReply   AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckType = "reply"
	AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckTypeRequest AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckType = "request"
)

// Whether the API call was successful
type AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseSuccess bool

const (
	AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseSuccessTrue AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsResponseSuccess = true
)

type AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponse struct {
	Errors   []AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseError   `json:"errors"`
	Messages []AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseMessage `json:"messages"`
	Result   AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseSuccess `json:"success"`
	JSON    accountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseJSON    `json:"-"`
}

// accountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseJSON contains
// the JSON metadata for the struct
// [AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponse]
type accountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseError struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    accountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseErrorJSON `json:"-"`
}

// accountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseError]
type accountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseMessage struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    accountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseMessageJSON `json:"-"`
}

// accountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseMessage]
type accountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResult struct {
	IpsecTunnels []AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnel `json:"ipsec_tunnels"`
	JSON         accountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultJSON          `json:"-"`
}

// accountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResult]
type accountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultJSON struct {
	IpsecTunnels apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnel struct {
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
	PskMetadata AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnelsPskMetadata `json:"psk_metadata"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection  bool                                                                                                `json:"replay_protection"`
	TunnelHealthCheck AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheck `json:"tunnel_health_check"`
	JSON              accountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnelJSON               `json:"-"`
}

// accountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnelJSON
// contains the JSON metadata for the struct
// [AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnel]
type accountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnelJSON struct {
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

func (r *AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The PSK metadata that includes when the PSK was generated.
type AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnelsPskMetadata struct {
	// The date and time the tunnel was last modified.
	LastGeneratedOn time.Time                                                                                         `json:"last_generated_on" format:"date-time"`
	JSON            accountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnelsPskMetadataJSON `json:"-"`
}

// accountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnelsPskMetadataJSON
// contains the JSON metadata for the struct
// [AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnelsPskMetadata]
type accountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnelsPskMetadataJSON struct {
	LastGeneratedOn apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnelsPskMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckType `json:"type"`
	JSON accountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckJSON `json:"-"`
}

// accountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckJSON
// contains the JSON metadata for the struct
// [AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheck]
type accountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckJSON struct {
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// How frequent the health check is run. The default value is `mid`.
type AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckRate string

const (
	AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckRateLow  AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckRate = "low"
	AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckRateMid  AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckRate = "mid"
	AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckRateHigh AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckType string

const (
	AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckTypeReply   AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckType = "reply"
	AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckTypeRequest AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseResultIpsecTunnelsTunnelHealthCheckType = "request"
)

// Whether the API call was successful
type AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseSuccess bool

const (
	AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseSuccessTrue AccountMagicIpsecTunnelMagicIPsecTunnelsListIPsecTunnelsResponseSuccess = true
)

type AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponse struct {
	Errors   []AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseError   `json:"errors"`
	Messages []AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseMessage `json:"messages"`
	Result   AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseSuccess `json:"success"`
	JSON    accountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseJSON    `json:"-"`
}

// accountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseJSON
// contains the JSON metadata for the struct
// [AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponse]
type accountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseError struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    accountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseErrorJSON `json:"-"`
}

// accountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseError]
type accountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseMessage struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    accountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseMessageJSON `json:"-"`
}

// accountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseMessage]
type accountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResult struct {
	Modified             bool                                                                                                  `json:"modified"`
	ModifiedIpsecTunnels []AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnel `json:"modified_ipsec_tunnels"`
	JSON                 accountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultJSON                  `json:"-"`
}

// accountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResult]
type accountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultJSON struct {
	Modified             apijson.Field
	ModifiedIpsecTunnels apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnel struct {
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
	PskMetadata AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnelsPskMetadata `json:"psk_metadata"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection  bool                                                                                                                  `json:"replay_protection"`
	TunnelHealthCheck AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnelsTunnelHealthCheck `json:"tunnel_health_check"`
	JSON              accountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnelJSON               `json:"-"`
}

// accountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnelJSON
// contains the JSON metadata for the struct
// [AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnel]
type accountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnelJSON struct {
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

func (r *AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The PSK metadata that includes when the PSK was generated.
type AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnelsPskMetadata struct {
	// The date and time the tunnel was last modified.
	LastGeneratedOn time.Time                                                                                                           `json:"last_generated_on" format:"date-time"`
	JSON            accountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnelsPskMetadataJSON `json:"-"`
}

// accountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnelsPskMetadataJSON
// contains the JSON metadata for the struct
// [AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnelsPskMetadata]
type accountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnelsPskMetadataJSON struct {
	LastGeneratedOn apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnelsPskMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnelsTunnelHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnelsTunnelHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnelsTunnelHealthCheckType `json:"type"`
	JSON accountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnelsTunnelHealthCheckJSON `json:"-"`
}

// accountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnelsTunnelHealthCheckJSON
// contains the JSON metadata for the struct
// [AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnelsTunnelHealthCheck]
type accountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnelsTunnelHealthCheckJSON struct {
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnelsTunnelHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// How frequent the health check is run. The default value is `mid`.
type AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnelsTunnelHealthCheckRate string

const (
	AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnelsTunnelHealthCheckRateLow  AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnelsTunnelHealthCheckRate = "low"
	AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnelsTunnelHealthCheckRateMid  AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnelsTunnelHealthCheckRate = "mid"
	AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnelsTunnelHealthCheckRateHigh AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnelsTunnelHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnelsTunnelHealthCheckType string

const (
	AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnelsTunnelHealthCheckTypeReply   AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnelsTunnelHealthCheckType = "reply"
	AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnelsTunnelHealthCheckTypeRequest AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseResultModifiedIpsecTunnelsTunnelHealthCheckType = "request"
)

// Whether the API call was successful
type AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseSuccess bool

const (
	AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseSuccessTrue AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsResponseSuccess = true
)

type AccountMagicIpsecTunnelUpdateParams struct {
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

func (r AccountMagicIpsecTunnelUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsParams struct {
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

func (r AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
