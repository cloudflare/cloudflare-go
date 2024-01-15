// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountMnmConfigService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountMnmConfigService] method
// instead.
type AccountMnmConfigService struct {
	Options []option.RequestOption
	Fulls   *AccountMnmConfigFullService
}

// NewAccountMnmConfigService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountMnmConfigService(opts ...option.RequestOption) (r *AccountMnmConfigService) {
	r = &AccountMnmConfigService{}
	r.Options = opts
	r.Fulls = NewAccountMnmConfigFullService(opts...)
	return
}

// Delete an existing network monitoring configuration.
func (r *AccountMnmConfigService) Delete(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *AccountMnmConfigDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/mnm/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a new network monitoring configuration.
func (r *AccountMnmConfigService) MagicNetworkMonitoringConfigurationNewAccountConfiguration(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *AccountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/mnm/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Lists default sampling and router IPs for account.
func (r *AccountMnmConfigService) MagicNetworkMonitoringConfigurationListAccountConfiguration(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *AccountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/mnm/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update fields in an existing network monitoring configuration.
func (r *AccountMnmConfigService) MagicNetworkMonitoringConfigurationUpdateAccountConfigurationFields(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/mnm/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

// Update an existing network monitoring configuration, requires the entire
// configuration to be updated at once.
func (r *AccountMnmConfigService) MagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfiguration(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/mnm/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

type AccountMnmConfigDeleteResponse struct {
	Errors   []AccountMnmConfigDeleteResponseError   `json:"errors"`
	Messages []AccountMnmConfigDeleteResponseMessage `json:"messages"`
	Result   AccountMnmConfigDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMnmConfigDeleteResponseSuccess `json:"success"`
	JSON    accountMnmConfigDeleteResponseJSON    `json:"-"`
}

// accountMnmConfigDeleteResponseJSON contains the JSON metadata for the struct
// [AccountMnmConfigDeleteResponse]
type accountMnmConfigDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmConfigDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmConfigDeleteResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accountMnmConfigDeleteResponseErrorJSON `json:"-"`
}

// accountMnmConfigDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountMnmConfigDeleteResponseError]
type accountMnmConfigDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmConfigDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmConfigDeleteResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountMnmConfigDeleteResponseMessageJSON `json:"-"`
}

// accountMnmConfigDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountMnmConfigDeleteResponseMessage]
type accountMnmConfigDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmConfigDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmConfigDeleteResponseResult struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                                   `json:"name,required"`
	RouterIPs []string                                 `json:"router_ips,required"`
	JSON      accountMnmConfigDeleteResponseResultJSON `json:"-"`
}

// accountMnmConfigDeleteResponseResultJSON contains the JSON metadata for the
// struct [AccountMnmConfigDeleteResponseResult]
type accountMnmConfigDeleteResponseResultJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountMnmConfigDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMnmConfigDeleteResponseSuccess bool

const (
	AccountMnmConfigDeleteResponseSuccessTrue AccountMnmConfigDeleteResponseSuccess = true
)

type AccountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponse struct {
	Errors   []AccountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseError   `json:"errors"`
	Messages []AccountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseMessage `json:"messages"`
	Result   AccountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseSuccess `json:"success"`
	JSON    accountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseJSON    `json:"-"`
}

// accountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseJSON
// contains the JSON metadata for the struct
// [AccountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponse]
type accountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseError struct {
	Code    int64                                                                                       `json:"code,required"`
	Message string                                                                                      `json:"message,required"`
	JSON    accountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseErrorJSON `json:"-"`
}

// accountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseError]
type accountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseMessage struct {
	Code    int64                                                                                         `json:"code,required"`
	Message string                                                                                        `json:"message,required"`
	JSON    accountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseMessageJSON `json:"-"`
}

// accountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseMessage]
type accountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseResult struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                                                                                       `json:"name,required"`
	RouterIPs []string                                                                                     `json:"router_ips,required"`
	JSON      accountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseResultJSON `json:"-"`
}

// accountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseResultJSON
// contains the JSON metadata for the struct
// [AccountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseResult]
type accountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseResultJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseSuccess bool

const (
	AccountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseSuccessTrue AccountMnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseSuccess = true
)

type AccountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponse struct {
	Errors   []AccountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseError   `json:"errors"`
	Messages []AccountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseMessage `json:"messages"`
	Result   AccountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseSuccess `json:"success"`
	JSON    accountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseJSON    `json:"-"`
}

// accountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseJSON
// contains the JSON metadata for the struct
// [AccountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponse]
type accountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseError struct {
	Code    int64                                                                                        `json:"code,required"`
	Message string                                                                                       `json:"message,required"`
	JSON    accountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseErrorJSON `json:"-"`
}

// accountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseError]
type accountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseMessage struct {
	Code    int64                                                                                          `json:"code,required"`
	Message string                                                                                         `json:"message,required"`
	JSON    accountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseMessageJSON `json:"-"`
}

// accountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseMessage]
type accountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseResult struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                                                                                        `json:"name,required"`
	RouterIPs []string                                                                                      `json:"router_ips,required"`
	JSON      accountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseResultJSON `json:"-"`
}

// accountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseResultJSON
// contains the JSON metadata for the struct
// [AccountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseResult]
type accountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseResultJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseSuccess bool

const (
	AccountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseSuccessTrue AccountMnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseSuccess = true
)

type AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponse struct {
	Errors   []AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseError   `json:"errors"`
	Messages []AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseMessage `json:"messages"`
	Result   AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseSuccess `json:"success"`
	JSON    accountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseJSON    `json:"-"`
}

// accountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseJSON
// contains the JSON metadata for the struct
// [AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponse]
type accountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseError struct {
	Code    int64                                                                                                `json:"code,required"`
	Message string                                                                                               `json:"message,required"`
	JSON    accountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseErrorJSON `json:"-"`
}

// accountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseError]
type accountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseMessage struct {
	Code    int64                                                                                                  `json:"code,required"`
	Message string                                                                                                 `json:"message,required"`
	JSON    accountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseMessageJSON `json:"-"`
}

// accountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseMessage]
type accountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseResult struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                                                                                                `json:"name,required"`
	RouterIPs []string                                                                                              `json:"router_ips,required"`
	JSON      accountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseResultJSON `json:"-"`
}

// accountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseResult]
type accountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseResultJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseSuccess bool

const (
	AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseSuccessTrue AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseSuccess = true
)

type AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponse struct {
	Errors   []AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseError   `json:"errors"`
	Messages []AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseMessage `json:"messages"`
	Result   AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseSuccess `json:"success"`
	JSON    accountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseJSON    `json:"-"`
}

// accountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseJSON
// contains the JSON metadata for the struct
// [AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponse]
type accountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseError struct {
	Code    int64                                                                                                  `json:"code,required"`
	Message string                                                                                                 `json:"message,required"`
	JSON    accountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseErrorJSON `json:"-"`
}

// accountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseError]
type accountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseMessage struct {
	Code    int64                                                                                                    `json:"code,required"`
	Message string                                                                                                   `json:"message,required"`
	JSON    accountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseMessageJSON `json:"-"`
}

// accountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseMessage]
type accountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseResult struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                                                                                                  `json:"name,required"`
	RouterIPs []string                                                                                                `json:"router_ips,required"`
	JSON      accountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseResultJSON `json:"-"`
}

// accountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseResultJSON
// contains the JSON metadata for the struct
// [AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseResult]
type accountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseResultJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseSuccess bool

const (
	AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseSuccessTrue AccountMnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseSuccess = true
)
