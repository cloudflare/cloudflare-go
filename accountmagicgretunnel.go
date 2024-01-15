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
func (r *AccountMagicGreTunnelService) Get(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *AccountMagicGreTunnelGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a specific GRE tunnel. Use `?validate_only=true` as an optional query
// parameter to only run validation without persisting changes.
func (r *AccountMagicGreTunnelService) Update(ctx context.Context, accountIdentifier string, tunnelIdentifier string, body AccountMagicGreTunnelUpdateParams, opts ...option.RequestOption) (res *AccountMagicGreTunnelUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Disables and removes a specific static GRE tunnel. Use `?validate_only=true` as
// an optional query parameter to only run validation without persisting changes.
func (r *AccountMagicGreTunnelService) Delete(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *AccountMagicGreTunnelDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates new GRE tunnels. Use `?validate_only=true` as an optional query
// parameter to only run validation without persisting changes.
func (r *AccountMagicGreTunnelService) MagicGreTunnelsNewGreTunnels(ctx context.Context, accountIdentifier string, body AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsParams, opts ...option.RequestOption) (res *AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists GRE tunnels associated with an account.
func (r *AccountMagicGreTunnelService) MagicGreTunnelsListGreTunnels(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates multiple GRE tunnels. Use `?validate_only=true` as an optional query
// parameter to only run validation without persisting changes.
func (r *AccountMagicGreTunnelService) MagicGreTunnelsUpdateMultipleGreTunnels(ctx context.Context, accountIdentifier string, body AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsParams, opts ...option.RequestOption) (res *AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccountMagicGreTunnelGetResponse struct {
	Errors   []AccountMagicGreTunnelGetResponseError   `json:"errors"`
	Messages []AccountMagicGreTunnelGetResponseMessage `json:"messages"`
	Result   AccountMagicGreTunnelGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMagicGreTunnelGetResponseSuccess `json:"success"`
	JSON    accountMagicGreTunnelGetResponseJSON    `json:"-"`
}

// accountMagicGreTunnelGetResponseJSON contains the JSON metadata for the struct
// [AccountMagicGreTunnelGetResponse]
type accountMagicGreTunnelGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicGreTunnelGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicGreTunnelGetResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountMagicGreTunnelGetResponseErrorJSON `json:"-"`
}

// accountMagicGreTunnelGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountMagicGreTunnelGetResponseError]
type accountMagicGreTunnelGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicGreTunnelGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicGreTunnelGetResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountMagicGreTunnelGetResponseMessageJSON `json:"-"`
}

// accountMagicGreTunnelGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountMagicGreTunnelGetResponseMessage]
type accountMagicGreTunnelGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicGreTunnelGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicGreTunnelGetResponseResult struct {
	GreTunnel interface{}                                `json:"gre_tunnel"`
	JSON      accountMagicGreTunnelGetResponseResultJSON `json:"-"`
}

// accountMagicGreTunnelGetResponseResultJSON contains the JSON metadata for the
// struct [AccountMagicGreTunnelGetResponseResult]
type accountMagicGreTunnelGetResponseResultJSON struct {
	GreTunnel   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicGreTunnelGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMagicGreTunnelGetResponseSuccess bool

const (
	AccountMagicGreTunnelGetResponseSuccessTrue AccountMagicGreTunnelGetResponseSuccess = true
)

type AccountMagicGreTunnelUpdateResponse struct {
	Errors   []AccountMagicGreTunnelUpdateResponseError   `json:"errors"`
	Messages []AccountMagicGreTunnelUpdateResponseMessage `json:"messages"`
	Result   AccountMagicGreTunnelUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMagicGreTunnelUpdateResponseSuccess `json:"success"`
	JSON    accountMagicGreTunnelUpdateResponseJSON    `json:"-"`
}

// accountMagicGreTunnelUpdateResponseJSON contains the JSON metadata for the
// struct [AccountMagicGreTunnelUpdateResponse]
type accountMagicGreTunnelUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicGreTunnelUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicGreTunnelUpdateResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountMagicGreTunnelUpdateResponseErrorJSON `json:"-"`
}

// accountMagicGreTunnelUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccountMagicGreTunnelUpdateResponseError]
type accountMagicGreTunnelUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicGreTunnelUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicGreTunnelUpdateResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountMagicGreTunnelUpdateResponseMessageJSON `json:"-"`
}

// accountMagicGreTunnelUpdateResponseMessageJSON contains the JSON metadata for
// the struct [AccountMagicGreTunnelUpdateResponseMessage]
type accountMagicGreTunnelUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicGreTunnelUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicGreTunnelUpdateResponseResult struct {
	Modified          bool                                          `json:"modified"`
	ModifiedGreTunnel interface{}                                   `json:"modified_gre_tunnel"`
	JSON              accountMagicGreTunnelUpdateResponseResultJSON `json:"-"`
}

// accountMagicGreTunnelUpdateResponseResultJSON contains the JSON metadata for the
// struct [AccountMagicGreTunnelUpdateResponseResult]
type accountMagicGreTunnelUpdateResponseResultJSON struct {
	Modified          apijson.Field
	ModifiedGreTunnel apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountMagicGreTunnelUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMagicGreTunnelUpdateResponseSuccess bool

const (
	AccountMagicGreTunnelUpdateResponseSuccessTrue AccountMagicGreTunnelUpdateResponseSuccess = true
)

type AccountMagicGreTunnelDeleteResponse struct {
	Errors   []AccountMagicGreTunnelDeleteResponseError   `json:"errors"`
	Messages []AccountMagicGreTunnelDeleteResponseMessage `json:"messages"`
	Result   AccountMagicGreTunnelDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMagicGreTunnelDeleteResponseSuccess `json:"success"`
	JSON    accountMagicGreTunnelDeleteResponseJSON    `json:"-"`
}

// accountMagicGreTunnelDeleteResponseJSON contains the JSON metadata for the
// struct [AccountMagicGreTunnelDeleteResponse]
type accountMagicGreTunnelDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicGreTunnelDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicGreTunnelDeleteResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountMagicGreTunnelDeleteResponseErrorJSON `json:"-"`
}

// accountMagicGreTunnelDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountMagicGreTunnelDeleteResponseError]
type accountMagicGreTunnelDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicGreTunnelDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicGreTunnelDeleteResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountMagicGreTunnelDeleteResponseMessageJSON `json:"-"`
}

// accountMagicGreTunnelDeleteResponseMessageJSON contains the JSON metadata for
// the struct [AccountMagicGreTunnelDeleteResponseMessage]
type accountMagicGreTunnelDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicGreTunnelDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicGreTunnelDeleteResponseResult struct {
	Deleted          bool                                          `json:"deleted"`
	DeletedGreTunnel interface{}                                   `json:"deleted_gre_tunnel"`
	JSON             accountMagicGreTunnelDeleteResponseResultJSON `json:"-"`
}

// accountMagicGreTunnelDeleteResponseResultJSON contains the JSON metadata for the
// struct [AccountMagicGreTunnelDeleteResponseResult]
type accountMagicGreTunnelDeleteResponseResultJSON struct {
	Deleted          apijson.Field
	DeletedGreTunnel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountMagicGreTunnelDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMagicGreTunnelDeleteResponseSuccess bool

const (
	AccountMagicGreTunnelDeleteResponseSuccessTrue AccountMagicGreTunnelDeleteResponseSuccess = true
)

type AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponse struct {
	Errors   []AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseError   `json:"errors"`
	Messages []AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseMessage `json:"messages"`
	Result   AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseSuccess `json:"success"`
	JSON    accountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseJSON    `json:"-"`
}

// accountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseJSON contains the JSON
// metadata for the struct
// [AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponse]
type accountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseError struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    accountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseErrorJSON `json:"-"`
}

// accountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseError]
type accountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseMessage struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    accountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseMessageJSON `json:"-"`
}

// accountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseMessageJSON contains
// the JSON metadata for the struct
// [AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseMessage]
type accountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResult struct {
	GreTunnels []AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnel `json:"gre_tunnels"`
	JSON       accountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultJSON        `json:"-"`
}

// accountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultJSON contains the
// JSON metadata for the struct
// [AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResult]
type accountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultJSON struct {
	GreTunnels  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnel struct {
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
	Description string                                                                               `json:"description"`
	HealthCheck AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnelsHealthCheck `json:"health_check"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value
	// is 576.
	Mtu int64 `json:"mtu"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	Ttl  int64                                                                        `json:"ttl"`
	JSON accountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnelJSON `json:"-"`
}

// accountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnelJSON
// contains the JSON metadata for the struct
// [AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnel]
type accountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnelJSON struct {
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

func (r *AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnelsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel.
	Direction AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnelsHealthCheckDirection `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnelsHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnelsHealthCheckType `json:"type"`
	JSON accountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnelsHealthCheckJSON `json:"-"`
}

// accountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnelsHealthCheckJSON
// contains the JSON metadata for the struct
// [AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnelsHealthCheck]
type accountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnelsHealthCheckJSON struct {
	Direction   apijson.Field
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnelsHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel.
type AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnelsHealthCheckDirection string

const (
	AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnelsHealthCheckDirectionUnidirectional AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnelsHealthCheckDirection = "unidirectional"
	AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnelsHealthCheckDirectionBidirectional  AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnelsHealthCheckDirection = "bidirectional"
)

// How frequent the health check is run. The default value is `mid`.
type AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnelsHealthCheckRate string

const (
	AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnelsHealthCheckRateLow  AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnelsHealthCheckRate = "low"
	AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnelsHealthCheckRateMid  AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnelsHealthCheckRate = "mid"
	AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnelsHealthCheckRateHigh AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnelsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnelsHealthCheckType string

const (
	AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnelsHealthCheckTypeReply   AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnelsHealthCheckType = "reply"
	AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnelsHealthCheckTypeRequest AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseResultGreTunnelsHealthCheckType = "request"
)

// Whether the API call was successful
type AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseSuccess bool

const (
	AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseSuccessTrue AccountMagicGreTunnelMagicGreTunnelsNewGreTunnelsResponseSuccess = true
)

type AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponse struct {
	Errors   []AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseError   `json:"errors"`
	Messages []AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseMessage `json:"messages"`
	Result   AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseSuccess `json:"success"`
	JSON    accountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseJSON    `json:"-"`
}

// accountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseJSON contains the JSON
// metadata for the struct
// [AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponse]
type accountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseError struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    accountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseErrorJSON `json:"-"`
}

// accountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseError]
type accountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseMessage struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    accountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseMessageJSON `json:"-"`
}

// accountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseMessageJSON contains
// the JSON metadata for the struct
// [AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseMessage]
type accountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResult struct {
	GreTunnels []AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnel `json:"gre_tunnels"`
	JSON       accountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultJSON        `json:"-"`
}

// accountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultJSON contains
// the JSON metadata for the struct
// [AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResult]
type accountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultJSON struct {
	GreTunnels  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnel struct {
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
	Description string                                                                                `json:"description"`
	HealthCheck AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnelsHealthCheck `json:"health_check"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value
	// is 576.
	Mtu int64 `json:"mtu"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	Ttl  int64                                                                         `json:"ttl"`
	JSON accountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnelJSON `json:"-"`
}

// accountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnelJSON
// contains the JSON metadata for the struct
// [AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnel]
type accountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnelJSON struct {
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

func (r *AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnelsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel.
	Direction AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnelsHealthCheckDirection `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnelsHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnelsHealthCheckType `json:"type"`
	JSON accountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnelsHealthCheckJSON `json:"-"`
}

// accountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnelsHealthCheckJSON
// contains the JSON metadata for the struct
// [AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnelsHealthCheck]
type accountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnelsHealthCheckJSON struct {
	Direction   apijson.Field
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnelsHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel.
type AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnelsHealthCheckDirection string

const (
	AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnelsHealthCheckDirectionUnidirectional AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnelsHealthCheckDirection = "unidirectional"
	AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnelsHealthCheckDirectionBidirectional  AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnelsHealthCheckDirection = "bidirectional"
)

// How frequent the health check is run. The default value is `mid`.
type AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnelsHealthCheckRate string

const (
	AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnelsHealthCheckRateLow  AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnelsHealthCheckRate = "low"
	AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnelsHealthCheckRateMid  AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnelsHealthCheckRate = "mid"
	AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnelsHealthCheckRateHigh AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnelsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnelsHealthCheckType string

const (
	AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnelsHealthCheckTypeReply   AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnelsHealthCheckType = "reply"
	AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnelsHealthCheckTypeRequest AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseResultGreTunnelsHealthCheckType = "request"
)

// Whether the API call was successful
type AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseSuccess bool

const (
	AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseSuccessTrue AccountMagicGreTunnelMagicGreTunnelsListGreTunnelsResponseSuccess = true
)

type AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponse struct {
	Errors   []AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseError   `json:"errors"`
	Messages []AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseMessage `json:"messages"`
	Result   AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseSuccess `json:"success"`
	JSON    accountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseJSON    `json:"-"`
}

// accountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseJSON
// contains the JSON metadata for the struct
// [AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponse]
type accountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseError struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    accountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseErrorJSON `json:"-"`
}

// accountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseError]
type accountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseMessage struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    accountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseMessageJSON `json:"-"`
}

// accountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseMessage]
type accountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResult struct {
	Modified           bool                                                                                          `json:"modified"`
	ModifiedGreTunnels []AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnel `json:"modified_gre_tunnels"`
	JSON               accountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultJSON                `json:"-"`
}

// accountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResult]
type accountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultJSON struct {
	Modified           apijson.Field
	ModifiedGreTunnels apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnel struct {
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
	Description string                                                                                                  `json:"description"`
	HealthCheck AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnelsHealthCheck `json:"health_check"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value
	// is 576.
	Mtu int64 `json:"mtu"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	Ttl  int64                                                                                           `json:"ttl"`
	JSON accountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnelJSON `json:"-"`
}

// accountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnelJSON
// contains the JSON metadata for the struct
// [AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnel]
type accountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnelJSON struct {
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

func (r *AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnelsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel.
	Direction AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnelsHealthCheckDirection `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnelsHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnelsHealthCheckType `json:"type"`
	JSON accountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnelsHealthCheckJSON `json:"-"`
}

// accountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnelsHealthCheckJSON
// contains the JSON metadata for the struct
// [AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnelsHealthCheck]
type accountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnelsHealthCheckJSON struct {
	Direction   apijson.Field
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnelsHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel.
type AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnelsHealthCheckDirection string

const (
	AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnelsHealthCheckDirectionUnidirectional AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnelsHealthCheckDirection = "unidirectional"
	AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnelsHealthCheckDirectionBidirectional  AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnelsHealthCheckDirection = "bidirectional"
)

// How frequent the health check is run. The default value is `mid`.
type AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnelsHealthCheckRate string

const (
	AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnelsHealthCheckRateLow  AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnelsHealthCheckRate = "low"
	AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnelsHealthCheckRateMid  AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnelsHealthCheckRate = "mid"
	AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnelsHealthCheckRateHigh AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnelsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnelsHealthCheckType string

const (
	AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnelsHealthCheckTypeReply   AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnelsHealthCheckType = "reply"
	AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnelsHealthCheckTypeRequest AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseResultModifiedGreTunnelsHealthCheckType = "request"
)

// Whether the API call was successful
type AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseSuccess bool

const (
	AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseSuccessTrue AccountMagicGreTunnelMagicGreTunnelsUpdateMultipleGreTunnelsResponseSuccess = true
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
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel.
	Direction param.Field[AccountMagicGreTunnelUpdateParamsHealthCheckDirection] `json:"direction"`
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

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel.
type AccountMagicGreTunnelUpdateParamsHealthCheckDirection string

const (
	AccountMagicGreTunnelUpdateParamsHealthCheckDirectionUnidirectional AccountMagicGreTunnelUpdateParamsHealthCheckDirection = "unidirectional"
	AccountMagicGreTunnelUpdateParamsHealthCheckDirectionBidirectional  AccountMagicGreTunnelUpdateParamsHealthCheckDirection = "bidirectional"
)

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
