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
func (r *AccountMagicIpsecTunnelService) Get(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *SchemasTunnelSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a specific IPsec tunnel associated with an account. Use
// `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes.
func (r *AccountMagicIpsecTunnelService) Update(ctx context.Context, accountIdentifier string, tunnelIdentifier string, body AccountMagicIpsecTunnelUpdateParams, opts ...option.RequestOption) (res *SchemasTunnelModifiedResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Disables and removes a specific static IPsec Tunnel associated with an account.
// Use `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes.
func (r *AccountMagicIpsecTunnelService) Delete(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *SchemasTunnelDeletedResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates new IPsec tunnels associated with an account. Use `?validate_only=true`
// as an optional query parameter to only run validation without persisting
// changes.
func (r *AccountMagicIpsecTunnelService) MagicIPsecTunnelsNewIPsecTunnels(ctx context.Context, accountIdentifier string, body AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsParams, opts ...option.RequestOption) (res *SchemasTunnelsCollectionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists IPsec tunnels associated with an account.
func (r *AccountMagicIpsecTunnelService) MagicIPsecTunnelsListIPsecTunnels(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *SchemasTunnelsCollectionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update multiple IPsec tunnels associated with an account. Use
// `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes.
func (r *AccountMagicIpsecTunnelService) MagicIPsecTunnelsUpdateMultipleIPsecTunnels(ctx context.Context, accountIdentifier string, body AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsParams, opts ...option.RequestOption) (res *SchemasModifiedTunnelsCollectionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type SchemasModifiedTunnelsCollectionResponse struct {
	Errors   []SchemasModifiedTunnelsCollectionResponseError   `json:"errors"`
	Messages []SchemasModifiedTunnelsCollectionResponseMessage `json:"messages"`
	Result   SchemasModifiedTunnelsCollectionResponseResult    `json:"result"`
	// Whether the API call was successful
	Success SchemasModifiedTunnelsCollectionResponseSuccess `json:"success"`
	JSON    schemasModifiedTunnelsCollectionResponseJSON    `json:"-"`
}

// schemasModifiedTunnelsCollectionResponseJSON contains the JSON metadata for the
// struct [SchemasModifiedTunnelsCollectionResponse]
type schemasModifiedTunnelsCollectionResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasModifiedTunnelsCollectionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasModifiedTunnelsCollectionResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    schemasModifiedTunnelsCollectionResponseErrorJSON `json:"-"`
}

// schemasModifiedTunnelsCollectionResponseErrorJSON contains the JSON metadata for
// the struct [SchemasModifiedTunnelsCollectionResponseError]
type schemasModifiedTunnelsCollectionResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasModifiedTunnelsCollectionResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasModifiedTunnelsCollectionResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    schemasModifiedTunnelsCollectionResponseMessageJSON `json:"-"`
}

// schemasModifiedTunnelsCollectionResponseMessageJSON contains the JSON metadata
// for the struct [SchemasModifiedTunnelsCollectionResponseMessage]
type schemasModifiedTunnelsCollectionResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasModifiedTunnelsCollectionResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasModifiedTunnelsCollectionResponseResult struct {
	Modified             bool                                                                `json:"modified"`
	ModifiedIpsecTunnels []SchemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnel `json:"modified_ipsec_tunnels"`
	JSON                 schemasModifiedTunnelsCollectionResponseResultJSON                  `json:"-"`
}

// schemasModifiedTunnelsCollectionResponseResultJSON contains the JSON metadata
// for the struct [SchemasModifiedTunnelsCollectionResponseResult]
type schemasModifiedTunnelsCollectionResponseResultJSON struct {
	Modified             apijson.Field
	ModifiedIpsecTunnels apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *SchemasModifiedTunnelsCollectionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnel struct {
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
	PskMetadata       SchemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnelsPskMetadata       `json:"psk_metadata"`
	TunnelHealthCheck SchemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnelsTunnelHealthCheck `json:"tunnel_health_check"`
	JSON              schemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnelJSON               `json:"-"`
}

// schemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnelJSON contains
// the JSON metadata for the struct
// [SchemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnel]
type schemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnelJSON struct {
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
	TunnelHealthCheck  apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SchemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The PSK metadata that includes when the PSK was generated.
type SchemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnelsPskMetadata struct {
	// The date and time the tunnel was last modified.
	LastGeneratedOn time.Time                                                                         `json:"last_generated_on" format:"date-time"`
	JSON            schemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnelsPskMetadataJSON `json:"-"`
}

// schemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnelsPskMetadataJSON
// contains the JSON metadata for the struct
// [SchemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnelsPskMetadata]
type schemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnelsPskMetadataJSON struct {
	LastGeneratedOn apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SchemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnelsPskMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnelsTunnelHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate SchemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnelsTunnelHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type SchemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnelsTunnelHealthCheckType `json:"type"`
	JSON schemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnelsTunnelHealthCheckJSON `json:"-"`
}

// schemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnelsTunnelHealthCheckJSON
// contains the JSON metadata for the struct
// [SchemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnelsTunnelHealthCheck]
type schemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnelsTunnelHealthCheckJSON struct {
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnelsTunnelHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// How frequent the health check is run. The default value is `mid`.
type SchemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnelsTunnelHealthCheckRate string

const (
	SchemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnelsTunnelHealthCheckRateLow  SchemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnelsTunnelHealthCheckRate = "low"
	SchemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnelsTunnelHealthCheckRateMid  SchemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnelsTunnelHealthCheckRate = "mid"
	SchemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnelsTunnelHealthCheckRateHigh SchemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnelsTunnelHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type SchemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnelsTunnelHealthCheckType string

const (
	SchemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnelsTunnelHealthCheckTypeReply   SchemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnelsTunnelHealthCheckType = "reply"
	SchemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnelsTunnelHealthCheckTypeRequest SchemasModifiedTunnelsCollectionResponseResultModifiedIpsecTunnelsTunnelHealthCheckType = "request"
)

// Whether the API call was successful
type SchemasModifiedTunnelsCollectionResponseSuccess bool

const (
	SchemasModifiedTunnelsCollectionResponseSuccessTrue SchemasModifiedTunnelsCollectionResponseSuccess = true
)

type SchemasTunnelDeletedResponse struct {
	Errors   []SchemasTunnelDeletedResponseError   `json:"errors"`
	Messages []SchemasTunnelDeletedResponseMessage `json:"messages"`
	Result   SchemasTunnelDeletedResponseResult    `json:"result"`
	// Whether the API call was successful
	Success SchemasTunnelDeletedResponseSuccess `json:"success"`
	JSON    schemasTunnelDeletedResponseJSON    `json:"-"`
}

// schemasTunnelDeletedResponseJSON contains the JSON metadata for the struct
// [SchemasTunnelDeletedResponse]
type schemasTunnelDeletedResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasTunnelDeletedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasTunnelDeletedResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    schemasTunnelDeletedResponseErrorJSON `json:"-"`
}

// schemasTunnelDeletedResponseErrorJSON contains the JSON metadata for the struct
// [SchemasTunnelDeletedResponseError]
type schemasTunnelDeletedResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasTunnelDeletedResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasTunnelDeletedResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    schemasTunnelDeletedResponseMessageJSON `json:"-"`
}

// schemasTunnelDeletedResponseMessageJSON contains the JSON metadata for the
// struct [SchemasTunnelDeletedResponseMessage]
type schemasTunnelDeletedResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasTunnelDeletedResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasTunnelDeletedResponseResult struct {
	Deleted            bool                                   `json:"deleted"`
	DeletedIpsecTunnel interface{}                            `json:"deleted_ipsec_tunnel"`
	JSON               schemasTunnelDeletedResponseResultJSON `json:"-"`
}

// schemasTunnelDeletedResponseResultJSON contains the JSON metadata for the struct
// [SchemasTunnelDeletedResponseResult]
type schemasTunnelDeletedResponseResultJSON struct {
	Deleted            apijson.Field
	DeletedIpsecTunnel apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SchemasTunnelDeletedResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SchemasTunnelDeletedResponseSuccess bool

const (
	SchemasTunnelDeletedResponseSuccessTrue SchemasTunnelDeletedResponseSuccess = true
)

type SchemasTunnelModifiedResponse struct {
	Errors   []SchemasTunnelModifiedResponseError   `json:"errors"`
	Messages []SchemasTunnelModifiedResponseMessage `json:"messages"`
	Result   SchemasTunnelModifiedResponseResult    `json:"result"`
	// Whether the API call was successful
	Success SchemasTunnelModifiedResponseSuccess `json:"success"`
	JSON    schemasTunnelModifiedResponseJSON    `json:"-"`
}

// schemasTunnelModifiedResponseJSON contains the JSON metadata for the struct
// [SchemasTunnelModifiedResponse]
type schemasTunnelModifiedResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasTunnelModifiedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasTunnelModifiedResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    schemasTunnelModifiedResponseErrorJSON `json:"-"`
}

// schemasTunnelModifiedResponseErrorJSON contains the JSON metadata for the struct
// [SchemasTunnelModifiedResponseError]
type schemasTunnelModifiedResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasTunnelModifiedResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasTunnelModifiedResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    schemasTunnelModifiedResponseMessageJSON `json:"-"`
}

// schemasTunnelModifiedResponseMessageJSON contains the JSON metadata for the
// struct [SchemasTunnelModifiedResponseMessage]
type schemasTunnelModifiedResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasTunnelModifiedResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasTunnelModifiedResponseResult struct {
	Modified            bool                                    `json:"modified"`
	ModifiedIpsecTunnel interface{}                             `json:"modified_ipsec_tunnel"`
	JSON                schemasTunnelModifiedResponseResultJSON `json:"-"`
}

// schemasTunnelModifiedResponseResultJSON contains the JSON metadata for the
// struct [SchemasTunnelModifiedResponseResult]
type schemasTunnelModifiedResponseResultJSON struct {
	Modified            apijson.Field
	ModifiedIpsecTunnel apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *SchemasTunnelModifiedResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SchemasTunnelModifiedResponseSuccess bool

const (
	SchemasTunnelModifiedResponseSuccessTrue SchemasTunnelModifiedResponseSuccess = true
)

type SchemasTunnelSingleResponse struct {
	Errors   []SchemasTunnelSingleResponseError   `json:"errors"`
	Messages []SchemasTunnelSingleResponseMessage `json:"messages"`
	Result   SchemasTunnelSingleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success SchemasTunnelSingleResponseSuccess `json:"success"`
	JSON    schemasTunnelSingleResponseJSON    `json:"-"`
}

// schemasTunnelSingleResponseJSON contains the JSON metadata for the struct
// [SchemasTunnelSingleResponse]
type schemasTunnelSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasTunnelSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasTunnelSingleResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    schemasTunnelSingleResponseErrorJSON `json:"-"`
}

// schemasTunnelSingleResponseErrorJSON contains the JSON metadata for the struct
// [SchemasTunnelSingleResponseError]
type schemasTunnelSingleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasTunnelSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasTunnelSingleResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    schemasTunnelSingleResponseMessageJSON `json:"-"`
}

// schemasTunnelSingleResponseMessageJSON contains the JSON metadata for the struct
// [SchemasTunnelSingleResponseMessage]
type schemasTunnelSingleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasTunnelSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasTunnelSingleResponseResult struct {
	IpsecTunnel interface{}                           `json:"ipsec_tunnel"`
	JSON        schemasTunnelSingleResponseResultJSON `json:"-"`
}

// schemasTunnelSingleResponseResultJSON contains the JSON metadata for the struct
// [SchemasTunnelSingleResponseResult]
type schemasTunnelSingleResponseResultJSON struct {
	IpsecTunnel apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasTunnelSingleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SchemasTunnelSingleResponseSuccess bool

const (
	SchemasTunnelSingleResponseSuccessTrue SchemasTunnelSingleResponseSuccess = true
)

type SchemasTunnelsCollectionResponse struct {
	Errors   []SchemasTunnelsCollectionResponseError   `json:"errors"`
	Messages []SchemasTunnelsCollectionResponseMessage `json:"messages"`
	Result   SchemasTunnelsCollectionResponseResult    `json:"result"`
	// Whether the API call was successful
	Success SchemasTunnelsCollectionResponseSuccess `json:"success"`
	JSON    schemasTunnelsCollectionResponseJSON    `json:"-"`
}

// schemasTunnelsCollectionResponseJSON contains the JSON metadata for the struct
// [SchemasTunnelsCollectionResponse]
type schemasTunnelsCollectionResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasTunnelsCollectionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasTunnelsCollectionResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    schemasTunnelsCollectionResponseErrorJSON `json:"-"`
}

// schemasTunnelsCollectionResponseErrorJSON contains the JSON metadata for the
// struct [SchemasTunnelsCollectionResponseError]
type schemasTunnelsCollectionResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasTunnelsCollectionResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasTunnelsCollectionResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    schemasTunnelsCollectionResponseMessageJSON `json:"-"`
}

// schemasTunnelsCollectionResponseMessageJSON contains the JSON metadata for the
// struct [SchemasTunnelsCollectionResponseMessage]
type schemasTunnelsCollectionResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasTunnelsCollectionResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasTunnelsCollectionResponseResult struct {
	IpsecTunnels []SchemasTunnelsCollectionResponseResultIpsecTunnel `json:"ipsec_tunnels"`
	JSON         schemasTunnelsCollectionResponseResultJSON          `json:"-"`
}

// schemasTunnelsCollectionResponseResultJSON contains the JSON metadata for the
// struct [SchemasTunnelsCollectionResponseResult]
type schemasTunnelsCollectionResponseResultJSON struct {
	IpsecTunnels apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SchemasTunnelsCollectionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasTunnelsCollectionResponseResultIpsecTunnel struct {
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
	PskMetadata       SchemasTunnelsCollectionResponseResultIpsecTunnelsPskMetadata       `json:"psk_metadata"`
	TunnelHealthCheck SchemasTunnelsCollectionResponseResultIpsecTunnelsTunnelHealthCheck `json:"tunnel_health_check"`
	JSON              schemasTunnelsCollectionResponseResultIpsecTunnelJSON               `json:"-"`
}

// schemasTunnelsCollectionResponseResultIpsecTunnelJSON contains the JSON metadata
// for the struct [SchemasTunnelsCollectionResponseResultIpsecTunnel]
type schemasTunnelsCollectionResponseResultIpsecTunnelJSON struct {
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
	TunnelHealthCheck  apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SchemasTunnelsCollectionResponseResultIpsecTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The PSK metadata that includes when the PSK was generated.
type SchemasTunnelsCollectionResponseResultIpsecTunnelsPskMetadata struct {
	// The date and time the tunnel was last modified.
	LastGeneratedOn time.Time                                                         `json:"last_generated_on" format:"date-time"`
	JSON            schemasTunnelsCollectionResponseResultIpsecTunnelsPskMetadataJSON `json:"-"`
}

// schemasTunnelsCollectionResponseResultIpsecTunnelsPskMetadataJSON contains the
// JSON metadata for the struct
// [SchemasTunnelsCollectionResponseResultIpsecTunnelsPskMetadata]
type schemasTunnelsCollectionResponseResultIpsecTunnelsPskMetadataJSON struct {
	LastGeneratedOn apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SchemasTunnelsCollectionResponseResultIpsecTunnelsPskMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasTunnelsCollectionResponseResultIpsecTunnelsTunnelHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate SchemasTunnelsCollectionResponseResultIpsecTunnelsTunnelHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type SchemasTunnelsCollectionResponseResultIpsecTunnelsTunnelHealthCheckType `json:"type"`
	JSON schemasTunnelsCollectionResponseResultIpsecTunnelsTunnelHealthCheckJSON `json:"-"`
}

// schemasTunnelsCollectionResponseResultIpsecTunnelsTunnelHealthCheckJSON contains
// the JSON metadata for the struct
// [SchemasTunnelsCollectionResponseResultIpsecTunnelsTunnelHealthCheck]
type schemasTunnelsCollectionResponseResultIpsecTunnelsTunnelHealthCheckJSON struct {
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasTunnelsCollectionResponseResultIpsecTunnelsTunnelHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// How frequent the health check is run. The default value is `mid`.
type SchemasTunnelsCollectionResponseResultIpsecTunnelsTunnelHealthCheckRate string

const (
	SchemasTunnelsCollectionResponseResultIpsecTunnelsTunnelHealthCheckRateLow  SchemasTunnelsCollectionResponseResultIpsecTunnelsTunnelHealthCheckRate = "low"
	SchemasTunnelsCollectionResponseResultIpsecTunnelsTunnelHealthCheckRateMid  SchemasTunnelsCollectionResponseResultIpsecTunnelsTunnelHealthCheckRate = "mid"
	SchemasTunnelsCollectionResponseResultIpsecTunnelsTunnelHealthCheckRateHigh SchemasTunnelsCollectionResponseResultIpsecTunnelsTunnelHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type SchemasTunnelsCollectionResponseResultIpsecTunnelsTunnelHealthCheckType string

const (
	SchemasTunnelsCollectionResponseResultIpsecTunnelsTunnelHealthCheckTypeReply   SchemasTunnelsCollectionResponseResultIpsecTunnelsTunnelHealthCheckType = "reply"
	SchemasTunnelsCollectionResponseResultIpsecTunnelsTunnelHealthCheckTypeRequest SchemasTunnelsCollectionResponseResultIpsecTunnelsTunnelHealthCheckType = "request"
)

// Whether the API call was successful
type SchemasTunnelsCollectionResponseSuccess bool

const (
	SchemasTunnelsCollectionResponseSuccessTrue SchemasTunnelsCollectionResponseSuccess = true
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
}

func (r AccountMagicIpsecTunnelUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountMagicIpsecTunnelMagicIPsecTunnelsNewIPsecTunnelsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountMagicIpsecTunnelMagicIPsecTunnelsUpdateMultipleIPsecTunnelsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
