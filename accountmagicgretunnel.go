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

// AccountMagicGreTunnelService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountMagicGreTunnelService]
// method instead.
type AccountMagicGreTunnelService struct {
	Options []option.RequestOption
}

// NewAccountMagicGreTunnelService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountMagicGreTunnelService(opts ...option.RequestOption) (r *AccountMagicGreTunnelService) {
	r = &AccountMagicGreTunnelService{}
	r.Options = opts
	return
}

// Lists informtion for a specific GRE tunnel.
func (r *AccountMagicGreTunnelService) Get(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *TunnelSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a specific GRE tunnel. Use `?validate_only=true` as an optional query
// parameter to only run validation without persisting changes.
func (r *AccountMagicGreTunnelService) Update(ctx context.Context, accountIdentifier string, tunnelIdentifier string, body AccountMagicGreTunnelUpdateParams, opts ...option.RequestOption) (res *TunnelModifiedResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Disables and removes a specific static GRE tunnel. Use `?validate_only=true` as
// an optional query parameter to only run validation without persisting changes.
func (r *AccountMagicGreTunnelService) Delete(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *TunnelDeletedResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates new GRE tunnels. Use `?validate_only=true` as an optional query
// parameter to only run validation without persisting changes.
func (r *AccountMagicGreTunnelService) MagicGreTunnelsNewGreTunnels(ctx context.Context, accountIdentifier string, body AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsParams, opts ...option.RequestOption) (res *TunnelsCollectionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists GRE tunnels associated with an account.
func (r *AccountMagicGreTunnelService) MagicGreTunnelsListGreTunnels(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *TunnelsCollectionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates multiple GRE tunnels. Use `?validate_only=true` as an optional query
// parameter to only run validation without persisting changes.
func (r *AccountMagicGreTunnelService) MagicGreTunnelsUpdateMultipleGreTunnels(ctx context.Context, accountIdentifier string, body AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsParams, opts ...option.RequestOption) (res *ModifiedTunnelsCollectionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ModifiedTunnelsCollectionResponse struct {
	Errors   []ModifiedTunnelsCollectionResponseError   `json:"errors"`
	Messages []ModifiedTunnelsCollectionResponseMessage `json:"messages"`
	Result   ModifiedTunnelsCollectionResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ModifiedTunnelsCollectionResponseSuccess `json:"success"`
	JSON    modifiedTunnelsCollectionResponseJSON    `json:"-"`
}

// modifiedTunnelsCollectionResponseJSON contains the JSON metadata for the struct
// [ModifiedTunnelsCollectionResponse]
type modifiedTunnelsCollectionResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModifiedTunnelsCollectionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ModifiedTunnelsCollectionResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    modifiedTunnelsCollectionResponseErrorJSON `json:"-"`
}

// modifiedTunnelsCollectionResponseErrorJSON contains the JSON metadata for the
// struct [ModifiedTunnelsCollectionResponseError]
type modifiedTunnelsCollectionResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModifiedTunnelsCollectionResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ModifiedTunnelsCollectionResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    modifiedTunnelsCollectionResponseMessageJSON `json:"-"`
}

// modifiedTunnelsCollectionResponseMessageJSON contains the JSON metadata for the
// struct [ModifiedTunnelsCollectionResponseMessage]
type modifiedTunnelsCollectionResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModifiedTunnelsCollectionResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ModifiedTunnelsCollectionResponseResult struct {
	Modified           bool                                                       `json:"modified"`
	ModifiedGreTunnels []ModifiedTunnelsCollectionResponseResultModifiedGreTunnel `json:"modified_gre_tunnels"`
	JSON               modifiedTunnelsCollectionResponseResultJSON                `json:"-"`
}

// modifiedTunnelsCollectionResponseResultJSON contains the JSON metadata for the
// struct [ModifiedTunnelsCollectionResponseResult]
type modifiedTunnelsCollectionResponseResultJSON struct {
	Modified           apijson.Field
	ModifiedGreTunnels apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ModifiedTunnelsCollectionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ModifiedTunnelsCollectionResponseResultModifiedGreTunnel struct {
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
	Description string                                                               `json:"description"`
	HealthCheck ModifiedTunnelsCollectionResponseResultModifiedGreTunnelsHealthCheck `json:"health_check"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value
	// is 576.
	Mtu int64 `json:"mtu"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	Ttl  int64                                                        `json:"ttl"`
	JSON modifiedTunnelsCollectionResponseResultModifiedGreTunnelJSON `json:"-"`
}

// modifiedTunnelsCollectionResponseResultModifiedGreTunnelJSON contains the JSON
// metadata for the struct
// [ModifiedTunnelsCollectionResponseResultModifiedGreTunnel]
type modifiedTunnelsCollectionResponseResultModifiedGreTunnelJSON struct {
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
	Ttl                   apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ModifiedTunnelsCollectionResponseResultModifiedGreTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ModifiedTunnelsCollectionResponseResultModifiedGreTunnelsHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate ModifiedTunnelsCollectionResponseResultModifiedGreTunnelsHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type ModifiedTunnelsCollectionResponseResultModifiedGreTunnelsHealthCheckType `json:"type"`
	JSON modifiedTunnelsCollectionResponseResultModifiedGreTunnelsHealthCheckJSON `json:"-"`
}

// modifiedTunnelsCollectionResponseResultModifiedGreTunnelsHealthCheckJSON
// contains the JSON metadata for the struct
// [ModifiedTunnelsCollectionResponseResultModifiedGreTunnelsHealthCheck]
type modifiedTunnelsCollectionResponseResultModifiedGreTunnelsHealthCheckJSON struct {
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModifiedTunnelsCollectionResponseResultModifiedGreTunnelsHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// How frequent the health check is run. The default value is `mid`.
type ModifiedTunnelsCollectionResponseResultModifiedGreTunnelsHealthCheckRate string

const (
	ModifiedTunnelsCollectionResponseResultModifiedGreTunnelsHealthCheckRateLow  ModifiedTunnelsCollectionResponseResultModifiedGreTunnelsHealthCheckRate = "low"
	ModifiedTunnelsCollectionResponseResultModifiedGreTunnelsHealthCheckRateMid  ModifiedTunnelsCollectionResponseResultModifiedGreTunnelsHealthCheckRate = "mid"
	ModifiedTunnelsCollectionResponseResultModifiedGreTunnelsHealthCheckRateHigh ModifiedTunnelsCollectionResponseResultModifiedGreTunnelsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type ModifiedTunnelsCollectionResponseResultModifiedGreTunnelsHealthCheckType string

const (
	ModifiedTunnelsCollectionResponseResultModifiedGreTunnelsHealthCheckTypeReply   ModifiedTunnelsCollectionResponseResultModifiedGreTunnelsHealthCheckType = "reply"
	ModifiedTunnelsCollectionResponseResultModifiedGreTunnelsHealthCheckTypeRequest ModifiedTunnelsCollectionResponseResultModifiedGreTunnelsHealthCheckType = "request"
)

// Whether the API call was successful
type ModifiedTunnelsCollectionResponseSuccess bool

const (
	ModifiedTunnelsCollectionResponseSuccessTrue ModifiedTunnelsCollectionResponseSuccess = true
)

type TunnelDeletedResponse struct {
	Errors   []TunnelDeletedResponseError   `json:"errors"`
	Messages []TunnelDeletedResponseMessage `json:"messages"`
	Result   TunnelDeletedResponseResult    `json:"result"`
	// Whether the API call was successful
	Success TunnelDeletedResponseSuccess `json:"success"`
	JSON    tunnelDeletedResponseJSON    `json:"-"`
}

// tunnelDeletedResponseJSON contains the JSON metadata for the struct
// [TunnelDeletedResponse]
type tunnelDeletedResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelDeletedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelDeletedResponseError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    tunnelDeletedResponseErrorJSON `json:"-"`
}

// tunnelDeletedResponseErrorJSON contains the JSON metadata for the struct
// [TunnelDeletedResponseError]
type tunnelDeletedResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelDeletedResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelDeletedResponseMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    tunnelDeletedResponseMessageJSON `json:"-"`
}

// tunnelDeletedResponseMessageJSON contains the JSON metadata for the struct
// [TunnelDeletedResponseMessage]
type tunnelDeletedResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelDeletedResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelDeletedResponseResult struct {
	Deleted          bool                            `json:"deleted"`
	DeletedGreTunnel interface{}                     `json:"deleted_gre_tunnel"`
	JSON             tunnelDeletedResponseResultJSON `json:"-"`
}

// tunnelDeletedResponseResultJSON contains the JSON metadata for the struct
// [TunnelDeletedResponseResult]
type tunnelDeletedResponseResultJSON struct {
	Deleted          apijson.Field
	DeletedGreTunnel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TunnelDeletedResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TunnelDeletedResponseSuccess bool

const (
	TunnelDeletedResponseSuccessTrue TunnelDeletedResponseSuccess = true
)

type TunnelModifiedResponse struct {
	Errors   []TunnelModifiedResponseError   `json:"errors"`
	Messages []TunnelModifiedResponseMessage `json:"messages"`
	Result   TunnelModifiedResponseResult    `json:"result"`
	// Whether the API call was successful
	Success TunnelModifiedResponseSuccess `json:"success"`
	JSON    tunnelModifiedResponseJSON    `json:"-"`
}

// tunnelModifiedResponseJSON contains the JSON metadata for the struct
// [TunnelModifiedResponse]
type tunnelModifiedResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelModifiedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelModifiedResponseError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    tunnelModifiedResponseErrorJSON `json:"-"`
}

// tunnelModifiedResponseErrorJSON contains the JSON metadata for the struct
// [TunnelModifiedResponseError]
type tunnelModifiedResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelModifiedResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelModifiedResponseMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    tunnelModifiedResponseMessageJSON `json:"-"`
}

// tunnelModifiedResponseMessageJSON contains the JSON metadata for the struct
// [TunnelModifiedResponseMessage]
type tunnelModifiedResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelModifiedResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelModifiedResponseResult struct {
	Modified          bool                             `json:"modified"`
	ModifiedGreTunnel interface{}                      `json:"modified_gre_tunnel"`
	JSON              tunnelModifiedResponseResultJSON `json:"-"`
}

// tunnelModifiedResponseResultJSON contains the JSON metadata for the struct
// [TunnelModifiedResponseResult]
type tunnelModifiedResponseResultJSON struct {
	Modified          apijson.Field
	ModifiedGreTunnel apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TunnelModifiedResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TunnelModifiedResponseSuccess bool

const (
	TunnelModifiedResponseSuccessTrue TunnelModifiedResponseSuccess = true
)

type TunnelSingleResponse struct {
	Errors   []TunnelSingleResponseError   `json:"errors"`
	Messages []TunnelSingleResponseMessage `json:"messages"`
	Result   TunnelSingleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success TunnelSingleResponseSuccess `json:"success"`
	JSON    tunnelSingleResponseJSON    `json:"-"`
}

// tunnelSingleResponseJSON contains the JSON metadata for the struct
// [TunnelSingleResponse]
type tunnelSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelSingleResponseError struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    tunnelSingleResponseErrorJSON `json:"-"`
}

// tunnelSingleResponseErrorJSON contains the JSON metadata for the struct
// [TunnelSingleResponseError]
type tunnelSingleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelSingleResponseMessage struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    tunnelSingleResponseMessageJSON `json:"-"`
}

// tunnelSingleResponseMessageJSON contains the JSON metadata for the struct
// [TunnelSingleResponseMessage]
type tunnelSingleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelSingleResponseResult struct {
	GreTunnel interface{}                    `json:"gre_tunnel"`
	JSON      tunnelSingleResponseResultJSON `json:"-"`
}

// tunnelSingleResponseResultJSON contains the JSON metadata for the struct
// [TunnelSingleResponseResult]
type tunnelSingleResponseResultJSON struct {
	GreTunnel   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelSingleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TunnelSingleResponseSuccess bool

const (
	TunnelSingleResponseSuccessTrue TunnelSingleResponseSuccess = true
)

type TunnelsCollectionResponse struct {
	Errors   []TunnelsCollectionResponseError   `json:"errors"`
	Messages []TunnelsCollectionResponseMessage `json:"messages"`
	Result   TunnelsCollectionResponseResult    `json:"result"`
	// Whether the API call was successful
	Success TunnelsCollectionResponseSuccess `json:"success"`
	JSON    tunnelsCollectionResponseJSON    `json:"-"`
}

// tunnelsCollectionResponseJSON contains the JSON metadata for the struct
// [TunnelsCollectionResponse]
type tunnelsCollectionResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelsCollectionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelsCollectionResponseError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    tunnelsCollectionResponseErrorJSON `json:"-"`
}

// tunnelsCollectionResponseErrorJSON contains the JSON metadata for the struct
// [TunnelsCollectionResponseError]
type tunnelsCollectionResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelsCollectionResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelsCollectionResponseMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    tunnelsCollectionResponseMessageJSON `json:"-"`
}

// tunnelsCollectionResponseMessageJSON contains the JSON metadata for the struct
// [TunnelsCollectionResponseMessage]
type tunnelsCollectionResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelsCollectionResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelsCollectionResponseResult struct {
	GreTunnels []TunnelsCollectionResponseResultGreTunnel `json:"gre_tunnels"`
	JSON       tunnelsCollectionResponseResultJSON        `json:"-"`
}

// tunnelsCollectionResponseResultJSON contains the JSON metadata for the struct
// [TunnelsCollectionResponseResult]
type tunnelsCollectionResponseResultJSON struct {
	GreTunnels  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelsCollectionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelsCollectionResponseResultGreTunnel struct {
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
	Description string                                               `json:"description"`
	HealthCheck TunnelsCollectionResponseResultGreTunnelsHealthCheck `json:"health_check"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value
	// is 576.
	Mtu int64 `json:"mtu"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	Ttl  int64                                        `json:"ttl"`
	JSON tunnelsCollectionResponseResultGreTunnelJSON `json:"-"`
}

// tunnelsCollectionResponseResultGreTunnelJSON contains the JSON metadata for the
// struct [TunnelsCollectionResponseResultGreTunnel]
type tunnelsCollectionResponseResultGreTunnelJSON struct {
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
	Ttl                   apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *TunnelsCollectionResponseResultGreTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelsCollectionResponseResultGreTunnelsHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate TunnelsCollectionResponseResultGreTunnelsHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type TunnelsCollectionResponseResultGreTunnelsHealthCheckType `json:"type"`
	JSON tunnelsCollectionResponseResultGreTunnelsHealthCheckJSON `json:"-"`
}

// tunnelsCollectionResponseResultGreTunnelsHealthCheckJSON contains the JSON
// metadata for the struct [TunnelsCollectionResponseResultGreTunnelsHealthCheck]
type tunnelsCollectionResponseResultGreTunnelsHealthCheckJSON struct {
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelsCollectionResponseResultGreTunnelsHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// How frequent the health check is run. The default value is `mid`.
type TunnelsCollectionResponseResultGreTunnelsHealthCheckRate string

const (
	TunnelsCollectionResponseResultGreTunnelsHealthCheckRateLow  TunnelsCollectionResponseResultGreTunnelsHealthCheckRate = "low"
	TunnelsCollectionResponseResultGreTunnelsHealthCheckRateMid  TunnelsCollectionResponseResultGreTunnelsHealthCheckRate = "mid"
	TunnelsCollectionResponseResultGreTunnelsHealthCheckRateHigh TunnelsCollectionResponseResultGreTunnelsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type TunnelsCollectionResponseResultGreTunnelsHealthCheckType string

const (
	TunnelsCollectionResponseResultGreTunnelsHealthCheckTypeReply   TunnelsCollectionResponseResultGreTunnelsHealthCheckType = "reply"
	TunnelsCollectionResponseResultGreTunnelsHealthCheckTypeRequest TunnelsCollectionResponseResultGreTunnelsHealthCheckType = "request"
)

// Whether the API call was successful
type TunnelsCollectionResponseSuccess bool

const (
	TunnelsCollectionResponseSuccessTrue TunnelsCollectionResponseSuccess = true
)

type AccountMagicGreTunnelUpdateParams struct {
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
	Description param.Field[string]                                       `json:"description"`
	HealthCheck param.Field[AccountMagicGreTunnelUpdateParamsHealthCheck] `json:"health_check"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value
	// is 576.
	Mtu param.Field[int64] `json:"mtu"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	Ttl param.Field[int64] `json:"ttl"`
}

func (r AccountMagicGreTunnelUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMagicGreTunnelUpdateParamsHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled param.Field[bool] `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate param.Field[AccountMagicGreTunnelUpdateParamsHealthCheckRate] `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target param.Field[string] `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type param.Field[AccountMagicGreTunnelUpdateParamsHealthCheckType] `json:"type"`
}

func (r AccountMagicGreTunnelUpdateParamsHealthCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How frequent the health check is run. The default value is `mid`.
type AccountMagicGreTunnelUpdateParamsHealthCheckRate string

const (
	AccountMagicGreTunnelUpdateParamsHealthCheckRateLow  AccountMagicGreTunnelUpdateParamsHealthCheckRate = "low"
	AccountMagicGreTunnelUpdateParamsHealthCheckRateMid  AccountMagicGreTunnelUpdateParamsHealthCheckRate = "mid"
	AccountMagicGreTunnelUpdateParamsHealthCheckRateHigh AccountMagicGreTunnelUpdateParamsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type AccountMagicGreTunnelUpdateParamsHealthCheckType string

const (
	AccountMagicGreTunnelUpdateParamsHealthCheckTypeReply   AccountMagicGreTunnelUpdateParamsHealthCheckType = "reply"
	AccountMagicGreTunnelUpdateParamsHealthCheckTypeRequest AccountMagicGreTunnelUpdateParamsHealthCheckType = "request"
)

type AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
