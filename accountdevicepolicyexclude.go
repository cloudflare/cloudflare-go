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

// AccountDevicePolicyExcludeService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountDevicePolicyExcludeService] method instead.
type AccountDevicePolicyExcludeService struct {
	Options []option.RequestOption
}

// NewAccountDevicePolicyExcludeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountDevicePolicyExcludeService(opts ...option.RequestOption) (r *AccountDevicePolicyExcludeService) {
	r = &AccountDevicePolicyExcludeService{}
	r.Options = opts
	return
}

// Fetches the list of routes excluded from the WARP client's tunnel.
func (r *AccountDevicePolicyExcludeService) DevicesGetSplitTunnelExcludeList(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/exclude", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches the list of routes excluded from the WARP client's tunnel for a specific
// device settings profile.
func (r *AccountDevicePolicyExcludeService) DevicesGetSplitTunnelExcludeListForADeviceSettingsPolicy(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/exclude", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Sets the list of routes excluded from the WARP client's tunnel.
func (r *AccountDevicePolicyExcludeService) DevicesSetSplitTunnelExcludeList(ctx context.Context, identifier interface{}, body AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListParams, opts ...option.RequestOption) (res *AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/exclude", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Sets the list of routes excluded from the WARP client's tunnel for a specific
// device settings profile.
func (r *AccountDevicePolicyExcludeService) DevicesSetSplitTunnelExcludeListForADeviceSettingsPolicy(ctx context.Context, identifier interface{}, uuid string, body AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyParams, opts ...option.RequestOption) (res *AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/exclude", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponse struct {
	Errors     []AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseError    `json:"errors"`
	Messages   []AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseMessage  `json:"messages"`
	Result     []AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseResult   `json:"result"`
	ResultInfo AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful.
	Success AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseSuccess `json:"success"`
	JSON    accountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseJSON    `json:"-"`
}

// accountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseJSON contains
// the JSON metadata for the struct
// [AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponse]
type accountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseError struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    accountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseErrorJSON `json:"-"`
}

// accountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseError]
type accountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseMessage struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    accountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseMessageJSON `json:"-"`
}

// accountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseMessage]
type accountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseResult struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host string                                                                       `json:"host"`
	JSON accountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseResultJSON `json:"-"`
}

// accountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseResultJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseResult]
type accountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseResultJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                          `json:"total_count"`
	JSON       accountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseResultInfoJSON `json:"-"`
}

// accountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseResultInfo]
type accountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseSuccess bool

const (
	AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseSuccessTrue AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseSuccess = true
)

type AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponse struct {
	Errors     []AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseError    `json:"errors"`
	Messages   []AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseMessage  `json:"messages"`
	Result     []AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseResult   `json:"result"`
	ResultInfo AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseResultInfo `json:"result_info"`
	// Whether the API call was successful.
	Success AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseSuccess `json:"success"`
	JSON    accountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseJSON    `json:"-"`
}

// accountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponse]
type accountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseError struct {
	Code    int64                                                                                               `json:"code,required"`
	Message string                                                                                              `json:"message,required"`
	JSON    accountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseErrorJSON `json:"-"`
}

// accountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseError]
type accountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseMessage struct {
	Code    int64                                                                                                 `json:"code,required"`
	Message string                                                                                                `json:"message,required"`
	JSON    accountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseMessageJSON `json:"-"`
}

// accountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseMessage]
type accountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseResult struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host string                                                                                               `json:"host"`
	JSON accountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseResultJSON `json:"-"`
}

// accountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseResultJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseResult]
type accountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseResultJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                  `json:"total_count"`
	JSON       accountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseResultInfoJSON `json:"-"`
}

// accountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseResultInfo]
type accountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseSuccess bool

const (
	AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseSuccessTrue AccountDevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseSuccess = true
)

type AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponse struct {
	Errors     []AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseError    `json:"errors"`
	Messages   []AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseMessage  `json:"messages"`
	Result     []AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseResult   `json:"result"`
	ResultInfo AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful.
	Success AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseSuccess `json:"success"`
	JSON    accountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseJSON    `json:"-"`
}

// accountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseJSON contains
// the JSON metadata for the struct
// [AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponse]
type accountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseError struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    accountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseErrorJSON `json:"-"`
}

// accountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseError]
type accountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseMessage struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    accountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseMessageJSON `json:"-"`
}

// accountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseMessage]
type accountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseResult struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host string                                                                       `json:"host"`
	JSON accountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseResultJSON `json:"-"`
}

// accountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseResultJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseResult]
type accountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseResultJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                          `json:"total_count"`
	JSON       accountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseResultInfoJSON `json:"-"`
}

// accountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseResultInfo]
type accountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseSuccess bool

const (
	AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseSuccessTrue AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseSuccess = true
)

type AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponse struct {
	Errors     []AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseError    `json:"errors"`
	Messages   []AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseMessage  `json:"messages"`
	Result     []AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseResult   `json:"result"`
	ResultInfo AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseResultInfo `json:"result_info"`
	// Whether the API call was successful.
	Success AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseSuccess `json:"success"`
	JSON    accountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseJSON    `json:"-"`
}

// accountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponse]
type accountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseError struct {
	Code    int64                                                                                               `json:"code,required"`
	Message string                                                                                              `json:"message,required"`
	JSON    accountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseErrorJSON `json:"-"`
}

// accountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseError]
type accountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseMessage struct {
	Code    int64                                                                                                 `json:"code,required"`
	Message string                                                                                                `json:"message,required"`
	JSON    accountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseMessageJSON `json:"-"`
}

// accountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseMessage]
type accountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseResult struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host string                                                                                               `json:"host"`
	JSON accountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseResultJSON `json:"-"`
}

// accountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseResultJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseResult]
type accountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseResultJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                  `json:"total_count"`
	JSON       accountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseResultInfoJSON `json:"-"`
}

// accountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseResultInfo]
type accountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseSuccess bool

const (
	AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseSuccessTrue AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseSuccess = true
)

type AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListParams struct {
	Body param.Field[[]AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListParamsBody] `json:"body,required"`
}

func (r AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListParamsBody struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address param.Field[string] `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host param.Field[string] `json:"host"`
}

func (r AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyParams struct {
	Body param.Field[[]AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyParamsBody] `json:"body,required"`
}

func (r AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyParamsBody struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address param.Field[string] `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host param.Field[string] `json:"host"`
}

func (r AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
