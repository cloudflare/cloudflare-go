// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountDevicePolicyFallbackDomainService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountDevicePolicyFallbackDomainService] method instead.
type AccountDevicePolicyFallbackDomainService struct {
	Options []option.RequestOption
}

// NewAccountDevicePolicyFallbackDomainService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountDevicePolicyFallbackDomainService(opts ...option.RequestOption) (r *AccountDevicePolicyFallbackDomainService) {
	r = &AccountDevicePolicyFallbackDomainService{}
	r.Options = opts
	return
}

// Fetches a list of domains to bypass Gateway DNS resolution. These domains will
// use the specified local DNS resolver instead.
func (r *AccountDevicePolicyFallbackDomainService) DevicesGetLocalDomainFallbackList(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/fallback_domains", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches the list of domains to bypass Gateway DNS resolution from a specified
// device settings profile. These domains will use the specified local DNS resolver
// instead.
func (r *AccountDevicePolicyFallbackDomainService) DevicesGetLocalDomainFallbackListForADeviceSettingsPolicy(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/fallback_domains", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Sets the list of domains to bypass Gateway DNS resolution. These domains will
// use the specified local DNS resolver instead.
func (r *AccountDevicePolicyFallbackDomainService) DevicesSetLocalDomainFallbackList(ctx context.Context, identifier interface{}, body AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListParams, opts ...option.RequestOption) (res *AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/fallback_domains", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Sets the list of domains to bypass Gateway DNS resolution. These domains will
// use the specified local DNS resolver instead. This will only apply to the
// specified device settings profile.
func (r *AccountDevicePolicyFallbackDomainService) DevicesSetLocalDomainFallbackListForADeviceSettingsPolicy(ctx context.Context, identifier interface{}, uuid string, body AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyParams, opts ...option.RequestOption) (res *AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/fallback_domains", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponse struct {
	Errors     []AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseError    `json:"errors"`
	Messages   []AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseMessage  `json:"messages"`
	Result     []AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseResult   `json:"result"`
	ResultInfo AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful.
	Success AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseSuccess `json:"success"`
	JSON    accountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseJSON    `json:"-"`
}

// accountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponse]
type accountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseError struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    accountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseErrorJSON `json:"-"`
}

// accountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseError]
type accountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseMessage struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    accountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseMessageJSON `json:"-"`
}

// accountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseMessage]
type accountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseResult struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                                                                        `json:"dns_server"`
	JSON      accountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseResultJSON `json:"-"`
}

// accountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseResultJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseResult]
type accountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseResultJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                  `json:"total_count"`
	JSON       accountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseResultInfoJSON `json:"-"`
}

// accountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseResultInfo]
type accountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseSuccess bool

const (
	AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseSuccessTrue AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseSuccess = true
)

type AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponse struct {
	Errors     []AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseError    `json:"errors"`
	Messages   []AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseMessage  `json:"messages"`
	Result     []AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseResult   `json:"result"`
	ResultInfo AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseResultInfo `json:"result_info"`
	// Whether the API call was successful.
	Success AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseSuccess `json:"success"`
	JSON    accountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseJSON    `json:"-"`
}

// accountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponse]
type accountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseError struct {
	Code    int64                                                                                                       `json:"code,required"`
	Message string                                                                                                      `json:"message,required"`
	JSON    accountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseErrorJSON `json:"-"`
}

// accountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseError]
type accountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseMessage struct {
	Code    int64                                                                                                         `json:"code,required"`
	Message string                                                                                                        `json:"message,required"`
	JSON    accountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseMessageJSON `json:"-"`
}

// accountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseMessage]
type accountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseResult struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                                                                                                `json:"dns_server"`
	JSON      accountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseResultJSON `json:"-"`
}

// accountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseResultJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseResult]
type accountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseResultJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                          `json:"total_count"`
	JSON       accountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseResultInfoJSON `json:"-"`
}

// accountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseResultInfo]
type accountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseSuccess bool

const (
	AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseSuccessTrue AccountDevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseSuccess = true
)

type AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponse struct {
	Errors     []AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseError    `json:"errors"`
	Messages   []AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseMessage  `json:"messages"`
	Result     []AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseResult   `json:"result"`
	ResultInfo AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful.
	Success AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseSuccess `json:"success"`
	JSON    accountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseJSON    `json:"-"`
}

// accountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponse]
type accountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseError struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    accountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseErrorJSON `json:"-"`
}

// accountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseError]
type accountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseMessage struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    accountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseMessageJSON `json:"-"`
}

// accountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseMessage]
type accountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseResult struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                                                                        `json:"dns_server"`
	JSON      accountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseResultJSON `json:"-"`
}

// accountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseResultJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseResult]
type accountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseResultJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                  `json:"total_count"`
	JSON       accountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseResultInfoJSON `json:"-"`
}

// accountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseResultInfo]
type accountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseSuccess bool

const (
	AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseSuccessTrue AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseSuccess = true
)

type AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponse struct {
	Errors     []AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseError    `json:"errors"`
	Messages   []AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseMessage  `json:"messages"`
	Result     []AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseResult   `json:"result"`
	ResultInfo AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseResultInfo `json:"result_info"`
	// Whether the API call was successful.
	Success AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseSuccess `json:"success"`
	JSON    accountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseJSON    `json:"-"`
}

// accountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponse]
type accountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseError struct {
	Code    int64                                                                                                       `json:"code,required"`
	Message string                                                                                                      `json:"message,required"`
	JSON    accountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseErrorJSON `json:"-"`
}

// accountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseError]
type accountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseMessage struct {
	Code    int64                                                                                                         `json:"code,required"`
	Message string                                                                                                        `json:"message,required"`
	JSON    accountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseMessageJSON `json:"-"`
}

// accountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseMessage]
type accountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseResult struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                                                                                                `json:"dns_server"`
	JSON      accountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseResultJSON `json:"-"`
}

// accountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseResultJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseResult]
type accountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseResultJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                          `json:"total_count"`
	JSON       accountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseResultInfoJSON `json:"-"`
}

// accountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseResultInfo]
type accountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseSuccess bool

const (
	AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseSuccessTrue AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseSuccess = true
)

type AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListParams struct {
	Body param.Field[[]AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListParamsBody] `json:"body,required"`
}

func (r AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListParamsBody struct {
	// The domain suffix to match when resolving locally.
	Suffix param.Field[string] `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description param.Field[string] `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer param.Field[[]interface{}] `json:"dns_server"`
}

func (r AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyParams struct {
	Body param.Field[[]AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyParamsBody] `json:"body,required"`
}

func (r AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyParamsBody struct {
	// The domain suffix to match when resolving locally.
	Suffix param.Field[string] `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description param.Field[string] `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer param.Field[[]interface{}] `json:"dns_server"`
}

func (r AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
