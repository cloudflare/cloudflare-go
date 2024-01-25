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

// AccountMagicCfInterconnectService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountMagicCfInterconnectService] method instead.
type AccountMagicCfInterconnectService struct {
	Options []option.RequestOption
}

// NewAccountMagicCfInterconnectService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountMagicCfInterconnectService(opts ...option.RequestOption) (r *AccountMagicCfInterconnectService) {
	r = &AccountMagicCfInterconnectService{}
	r.Options = opts
	return
}

// Lists details for a specific interconnect.
func (r *AccountMagicCfInterconnectService) Get(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *AccountMagicCfInterconnectGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/cf_interconnects/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a specific interconnect associated with an account. Use
// `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes.
func (r *AccountMagicCfInterconnectService) Update(ctx context.Context, accountIdentifier string, tunnelIdentifier string, body AccountMagicCfInterconnectUpdateParams, opts ...option.RequestOption) (res *AccountMagicCfInterconnectUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/cf_interconnects/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists interconnects associated with an account.
func (r *AccountMagicCfInterconnectService) MagicInterconnectsListInterconnects(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/cf_interconnects", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates multiple interconnects associated with an account. Use
// `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes.
func (r *AccountMagicCfInterconnectService) MagicInterconnectsUpdateMultipleInterconnects(ctx context.Context, accountIdentifier string, body AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsParams, opts ...option.RequestOption) (res *AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/cf_interconnects", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccountMagicCfInterconnectGetResponse struct {
	Errors   []AccountMagicCfInterconnectGetResponseError   `json:"errors"`
	Messages []AccountMagicCfInterconnectGetResponseMessage `json:"messages"`
	Result   AccountMagicCfInterconnectGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMagicCfInterconnectGetResponseSuccess `json:"success"`
	JSON    accountMagicCfInterconnectGetResponseJSON    `json:"-"`
}

// accountMagicCfInterconnectGetResponseJSON contains the JSON metadata for the
// struct [AccountMagicCfInterconnectGetResponse]
type accountMagicCfInterconnectGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCfInterconnectGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicCfInterconnectGetResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountMagicCfInterconnectGetResponseErrorJSON `json:"-"`
}

// accountMagicCfInterconnectGetResponseErrorJSON contains the JSON metadata for
// the struct [AccountMagicCfInterconnectGetResponseError]
type accountMagicCfInterconnectGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCfInterconnectGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicCfInterconnectGetResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accountMagicCfInterconnectGetResponseMessageJSON `json:"-"`
}

// accountMagicCfInterconnectGetResponseMessageJSON contains the JSON metadata for
// the struct [AccountMagicCfInterconnectGetResponseMessage]
type accountMagicCfInterconnectGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCfInterconnectGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicCfInterconnectGetResponseResult struct {
	Interconnect interface{}                                     `json:"interconnect"`
	JSON         accountMagicCfInterconnectGetResponseResultJSON `json:"-"`
}

// accountMagicCfInterconnectGetResponseResultJSON contains the JSON metadata for
// the struct [AccountMagicCfInterconnectGetResponseResult]
type accountMagicCfInterconnectGetResponseResultJSON struct {
	Interconnect apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountMagicCfInterconnectGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMagicCfInterconnectGetResponseSuccess bool

const (
	AccountMagicCfInterconnectGetResponseSuccessTrue AccountMagicCfInterconnectGetResponseSuccess = true
)

type AccountMagicCfInterconnectUpdateResponse struct {
	Errors   []AccountMagicCfInterconnectUpdateResponseError   `json:"errors"`
	Messages []AccountMagicCfInterconnectUpdateResponseMessage `json:"messages"`
	Result   AccountMagicCfInterconnectUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMagicCfInterconnectUpdateResponseSuccess `json:"success"`
	JSON    accountMagicCfInterconnectUpdateResponseJSON    `json:"-"`
}

// accountMagicCfInterconnectUpdateResponseJSON contains the JSON metadata for the
// struct [AccountMagicCfInterconnectUpdateResponse]
type accountMagicCfInterconnectUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCfInterconnectUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicCfInterconnectUpdateResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountMagicCfInterconnectUpdateResponseErrorJSON `json:"-"`
}

// accountMagicCfInterconnectUpdateResponseErrorJSON contains the JSON metadata for
// the struct [AccountMagicCfInterconnectUpdateResponseError]
type accountMagicCfInterconnectUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCfInterconnectUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicCfInterconnectUpdateResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountMagicCfInterconnectUpdateResponseMessageJSON `json:"-"`
}

// accountMagicCfInterconnectUpdateResponseMessageJSON contains the JSON metadata
// for the struct [AccountMagicCfInterconnectUpdateResponseMessage]
type accountMagicCfInterconnectUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCfInterconnectUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicCfInterconnectUpdateResponseResult struct {
	Modified             bool                                               `json:"modified"`
	ModifiedInterconnect interface{}                                        `json:"modified_interconnect"`
	JSON                 accountMagicCfInterconnectUpdateResponseResultJSON `json:"-"`
}

// accountMagicCfInterconnectUpdateResponseResultJSON contains the JSON metadata
// for the struct [AccountMagicCfInterconnectUpdateResponseResult]
type accountMagicCfInterconnectUpdateResponseResultJSON struct {
	Modified             apijson.Field
	ModifiedInterconnect apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountMagicCfInterconnectUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMagicCfInterconnectUpdateResponseSuccess bool

const (
	AccountMagicCfInterconnectUpdateResponseSuccessTrue AccountMagicCfInterconnectUpdateResponseSuccess = true
)

type AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponse struct {
	Errors   []AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseError   `json:"errors"`
	Messages []AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseMessage `json:"messages"`
	Result   AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseSuccess `json:"success"`
	JSON    accountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseJSON    `json:"-"`
}

// accountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseJSON
// contains the JSON metadata for the struct
// [AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponse]
type accountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseError struct {
	Code    int64                                                                          `json:"code,required"`
	Message string                                                                         `json:"message,required"`
	JSON    accountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseErrorJSON `json:"-"`
}

// accountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseError]
type accountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseMessage struct {
	Code    int64                                                                            `json:"code,required"`
	Message string                                                                           `json:"message,required"`
	JSON    accountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseMessageJSON `json:"-"`
}

// accountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseMessage]
type accountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResult struct {
	Interconnects []AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnect `json:"interconnects"`
	JSON          accountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultJSON           `json:"-"`
}

// accountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResult]
type accountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultJSON struct {
	Interconnects apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnect struct {
	// Tunnel identifier tag.
	ID string `json:"id"`
	// The name of the interconnect. The name cannot share a name with other tunnels.
	ColoName string `json:"colo_name"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the interconnect.
	Description string `json:"description"`
	// The configuration specific to GRE interconnects.
	Gre         AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnectsGre         `json:"gre"`
	HealthCheck AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnectsHealthCheck `json:"health_check"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress string `json:"interface_address"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The Maximum Transmission Unit (MTU) in bytes for the interconnect. The minimum
	// value is 576.
	Mtu int64 `json:"mtu"`
	// The name of the interconnect. The name cannot share a name with other tunnels.
	Name string                                                                                      `json:"name"`
	JSON accountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnectJSON `json:"-"`
}

// accountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnectJSON
// contains the JSON metadata for the struct
// [AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnect]
type accountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnectJSON struct {
	ID               apijson.Field
	ColoName         apijson.Field
	CreatedOn        apijson.Field
	Description      apijson.Field
	Gre              apijson.Field
	HealthCheck      apijson.Field
	InterfaceAddress apijson.Field
	ModifiedOn       apijson.Field
	Mtu              apijson.Field
	Name             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration specific to GRE interconnects.
type AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnectsGre struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel created as part
	// of the Interconnect.
	CloudflareEndpoint string                                                                                          `json:"cloudflare_endpoint"`
	JSON               accountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnectsGreJSON `json:"-"`
}

// accountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnectsGreJSON
// contains the JSON metadata for the struct
// [AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnectsGre]
type accountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnectsGreJSON struct {
	CloudflareEndpoint apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnectsGre) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnectsHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnectsHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnectsHealthCheckType `json:"type"`
	JSON accountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnectsHealthCheckJSON `json:"-"`
}

// accountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnectsHealthCheckJSON
// contains the JSON metadata for the struct
// [AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnectsHealthCheck]
type accountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnectsHealthCheckJSON struct {
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnectsHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// How frequent the health check is run. The default value is `mid`.
type AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnectsHealthCheckRate string

const (
	AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnectsHealthCheckRateLow  AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnectsHealthCheckRate = "low"
	AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnectsHealthCheckRateMid  AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnectsHealthCheckRate = "mid"
	AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnectsHealthCheckRateHigh AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnectsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnectsHealthCheckType string

const (
	AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnectsHealthCheckTypeReply   AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnectsHealthCheckType = "reply"
	AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnectsHealthCheckTypeRequest AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseResultInterconnectsHealthCheckType = "request"
)

// Whether the API call was successful
type AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseSuccess bool

const (
	AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseSuccessTrue AccountMagicCfInterconnectMagicInterconnectsListInterconnectsResponseSuccess = true
)

type AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponse struct {
	Errors   []AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseError   `json:"errors"`
	Messages []AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseMessage `json:"messages"`
	Result   AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseSuccess `json:"success"`
	JSON    accountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseJSON    `json:"-"`
}

// accountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseJSON
// contains the JSON metadata for the struct
// [AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponse]
type accountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseError struct {
	Code    int64                                                                                    `json:"code,required"`
	Message string                                                                                   `json:"message,required"`
	JSON    accountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseErrorJSON `json:"-"`
}

// accountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseError]
type accountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseMessage struct {
	Code    int64                                                                                      `json:"code,required"`
	Message string                                                                                     `json:"message,required"`
	JSON    accountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseMessageJSON `json:"-"`
}

// accountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseMessage]
type accountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResult struct {
	Modified              bool                                                                                                        `json:"modified"`
	ModifiedInterconnects []AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnect `json:"modified_interconnects"`
	JSON                  accountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultJSON                   `json:"-"`
}

// accountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResult]
type accountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultJSON struct {
	Modified              apijson.Field
	ModifiedInterconnects apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnect struct {
	// Tunnel identifier tag.
	ID string `json:"id"`
	// The name of the interconnect. The name cannot share a name with other tunnels.
	ColoName string `json:"colo_name"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the interconnect.
	Description string `json:"description"`
	// The configuration specific to GRE interconnects.
	Gre         AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnectsGre         `json:"gre"`
	HealthCheck AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnectsHealthCheck `json:"health_check"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress string `json:"interface_address"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The Maximum Transmission Unit (MTU) in bytes for the interconnect. The minimum
	// value is 576.
	Mtu int64 `json:"mtu"`
	// The name of the interconnect. The name cannot share a name with other tunnels.
	Name string                                                                                                        `json:"name"`
	JSON accountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnectJSON `json:"-"`
}

// accountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnectJSON
// contains the JSON metadata for the struct
// [AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnect]
type accountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnectJSON struct {
	ID               apijson.Field
	ColoName         apijson.Field
	CreatedOn        apijson.Field
	Description      apijson.Field
	Gre              apijson.Field
	HealthCheck      apijson.Field
	InterfaceAddress apijson.Field
	ModifiedOn       apijson.Field
	Mtu              apijson.Field
	Name             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration specific to GRE interconnects.
type AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnectsGre struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel created as part
	// of the Interconnect.
	CloudflareEndpoint string                                                                                                            `json:"cloudflare_endpoint"`
	JSON               accountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnectsGreJSON `json:"-"`
}

// accountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnectsGreJSON
// contains the JSON metadata for the struct
// [AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnectsGre]
type accountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnectsGreJSON struct {
	CloudflareEndpoint apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnectsGre) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnectsHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnectsHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnectsHealthCheckType `json:"type"`
	JSON accountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnectsHealthCheckJSON `json:"-"`
}

// accountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnectsHealthCheckJSON
// contains the JSON metadata for the struct
// [AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnectsHealthCheck]
type accountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnectsHealthCheckJSON struct {
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnectsHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// How frequent the health check is run. The default value is `mid`.
type AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnectsHealthCheckRate string

const (
	AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnectsHealthCheckRateLow  AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnectsHealthCheckRate = "low"
	AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnectsHealthCheckRateMid  AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnectsHealthCheckRate = "mid"
	AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnectsHealthCheckRateHigh AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnectsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnectsHealthCheckType string

const (
	AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnectsHealthCheckTypeReply   AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnectsHealthCheckType = "reply"
	AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnectsHealthCheckTypeRequest AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseResultModifiedInterconnectsHealthCheckType = "request"
)

// Whether the API call was successful
type AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseSuccess bool

const (
	AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseSuccessTrue AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseSuccess = true
)

type AccountMagicCfInterconnectUpdateParams struct {
	// An optional description of the interconnect.
	Description param.Field[string] `json:"description"`
	// The configuration specific to GRE interconnects.
	Gre         param.Field[AccountMagicCfInterconnectUpdateParamsGre]         `json:"gre"`
	HealthCheck param.Field[AccountMagicCfInterconnectUpdateParamsHealthCheck] `json:"health_check"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress param.Field[string] `json:"interface_address"`
	// The Maximum Transmission Unit (MTU) in bytes for the interconnect. The minimum
	// value is 576.
	Mtu param.Field[int64] `json:"mtu"`
}

func (r AccountMagicCfInterconnectUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration specific to GRE interconnects.
type AccountMagicCfInterconnectUpdateParamsGre struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel created as part
	// of the Interconnect.
	CloudflareEndpoint param.Field[string] `json:"cloudflare_endpoint"`
}

func (r AccountMagicCfInterconnectUpdateParamsGre) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMagicCfInterconnectUpdateParamsHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled param.Field[bool] `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate param.Field[AccountMagicCfInterconnectUpdateParamsHealthCheckRate] `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target param.Field[string] `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type param.Field[AccountMagicCfInterconnectUpdateParamsHealthCheckType] `json:"type"`
}

func (r AccountMagicCfInterconnectUpdateParamsHealthCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How frequent the health check is run. The default value is `mid`.
type AccountMagicCfInterconnectUpdateParamsHealthCheckRate string

const (
	AccountMagicCfInterconnectUpdateParamsHealthCheckRateLow  AccountMagicCfInterconnectUpdateParamsHealthCheckRate = "low"
	AccountMagicCfInterconnectUpdateParamsHealthCheckRateMid  AccountMagicCfInterconnectUpdateParamsHealthCheckRate = "mid"
	AccountMagicCfInterconnectUpdateParamsHealthCheckRateHigh AccountMagicCfInterconnectUpdateParamsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type AccountMagicCfInterconnectUpdateParamsHealthCheckType string

const (
	AccountMagicCfInterconnectUpdateParamsHealthCheckTypeReply   AccountMagicCfInterconnectUpdateParamsHealthCheckType = "reply"
	AccountMagicCfInterconnectUpdateParamsHealthCheckTypeRequest AccountMagicCfInterconnectUpdateParamsHealthCheckType = "request"
)

type AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountMagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
