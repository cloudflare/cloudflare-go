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

// AccountDevicePolicyIncludeService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountDevicePolicyIncludeService] method instead.
type AccountDevicePolicyIncludeService struct {
	Options []option.RequestOption
}

// NewAccountDevicePolicyIncludeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountDevicePolicyIncludeService(opts ...option.RequestOption) (r *AccountDevicePolicyIncludeService) {
	r = &AccountDevicePolicyIncludeService{}
	r.Options = opts
	return
}

// Fetches the list of routes included in the WARP client's tunnel.
func (r *AccountDevicePolicyIncludeService) DevicesGetSplitTunnelIncludeList(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/include", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches the list of routes included in the WARP client's tunnel for a specific
// device settings profile.
func (r *AccountDevicePolicyIncludeService) DevicesGetSplitTunnelIncludeListForADeviceSettingsPolicy(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/include", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Sets the list of routes included in the WARP client's tunnel.
func (r *AccountDevicePolicyIncludeService) DevicesSetSplitTunnelIncludeList(ctx context.Context, identifier interface{}, body AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListParams, opts ...option.RequestOption) (res *AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/include", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Sets the list of routes included in the WARP client's tunnel for a specific
// device settings profile.
func (r *AccountDevicePolicyIncludeService) DevicesSetSplitTunnelIncludeListForADeviceSettingsPolicy(ctx context.Context, identifier interface{}, uuid string, body AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyParams, opts ...option.RequestOption) (res *AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/include", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponse struct {
	Errors     []AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseError    `json:"errors"`
	Messages   []AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseMessage  `json:"messages"`
	Result     []AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseResult   `json:"result"`
	ResultInfo AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful.
	Success AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseSuccess `json:"success"`
	JSON    accountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseJSON    `json:"-"`
}

// accountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseJSON contains
// the JSON metadata for the struct
// [AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponse]
type accountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseError struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    accountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseErrorJSON `json:"-"`
}

// accountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseError]
type accountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseMessage struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    accountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseMessageJSON `json:"-"`
}

// accountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseMessage]
type accountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseResult struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                                                                       `json:"host"`
	JSON accountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseResultJSON `json:"-"`
}

// accountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseResultJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseResult]
type accountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseResultJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                          `json:"total_count"`
	JSON       accountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseResultInfoJSON `json:"-"`
}

// accountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseResultInfo]
type accountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseSuccess bool

const (
	AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseSuccessTrue AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseSuccess = true
)

type AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponse struct {
	Errors     []AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseError    `json:"errors"`
	Messages   []AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseMessage  `json:"messages"`
	Result     []AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseResult   `json:"result"`
	ResultInfo AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseResultInfo `json:"result_info"`
	// Whether the API call was successful.
	Success AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseSuccess `json:"success"`
	JSON    accountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseJSON    `json:"-"`
}

// accountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponse]
type accountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseError struct {
	Code    int64                                                                                               `json:"code,required"`
	Message string                                                                                              `json:"message,required"`
	JSON    accountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseErrorJSON `json:"-"`
}

// accountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseError]
type accountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseMessage struct {
	Code    int64                                                                                                 `json:"code,required"`
	Message string                                                                                                `json:"message,required"`
	JSON    accountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseMessageJSON `json:"-"`
}

// accountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseMessage]
type accountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseResult struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                                                                                               `json:"host"`
	JSON accountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseResultJSON `json:"-"`
}

// accountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseResultJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseResult]
type accountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseResultJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                  `json:"total_count"`
	JSON       accountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseResultInfoJSON `json:"-"`
}

// accountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseResultInfo]
type accountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseSuccess bool

const (
	AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseSuccessTrue AccountDevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseSuccess = true
)

type AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponse struct {
	Errors     []AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseError    `json:"errors"`
	Messages   []AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseMessage  `json:"messages"`
	Result     []AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseResult   `json:"result"`
	ResultInfo AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful.
	Success AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseSuccess `json:"success"`
	JSON    accountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseJSON    `json:"-"`
}

// accountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseJSON contains
// the JSON metadata for the struct
// [AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponse]
type accountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseError struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    accountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseErrorJSON `json:"-"`
}

// accountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseError]
type accountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseMessage struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    accountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseMessageJSON `json:"-"`
}

// accountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseMessage]
type accountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseResult struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                                                                       `json:"host"`
	JSON accountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseResultJSON `json:"-"`
}

// accountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseResultJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseResult]
type accountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseResultJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                          `json:"total_count"`
	JSON       accountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseResultInfoJSON `json:"-"`
}

// accountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseResultInfo]
type accountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseSuccess bool

const (
	AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseSuccessTrue AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseSuccess = true
)

type AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponse struct {
	Errors     []AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseError    `json:"errors"`
	Messages   []AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseMessage  `json:"messages"`
	Result     []AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseResult   `json:"result"`
	ResultInfo AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseResultInfo `json:"result_info"`
	// Whether the API call was successful.
	Success AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseSuccess `json:"success"`
	JSON    accountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseJSON    `json:"-"`
}

// accountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponse]
type accountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseError struct {
	Code    int64                                                                                               `json:"code,required"`
	Message string                                                                                              `json:"message,required"`
	JSON    accountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseErrorJSON `json:"-"`
}

// accountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseError]
type accountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseMessage struct {
	Code    int64                                                                                                 `json:"code,required"`
	Message string                                                                                                `json:"message,required"`
	JSON    accountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseMessageJSON `json:"-"`
}

// accountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseMessage]
type accountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseResult struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                                                                                               `json:"host"`
	JSON accountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseResultJSON `json:"-"`
}

// accountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseResultJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseResult]
type accountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseResultJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                  `json:"total_count"`
	JSON       accountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseResultInfoJSON `json:"-"`
}

// accountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseResultInfo]
type accountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseSuccess bool

const (
	AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseSuccessTrue AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseSuccess = true
)

type AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListParams struct {
	Body param.Field[[]AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListParamsBody] `json:"body,required"`
}

func (r AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListParamsBody struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address param.Field[string] `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host param.Field[string] `json:"host"`
}

func (r AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyParams struct {
	Body param.Field[[]AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyParamsBody] `json:"body,required"`
}

func (r AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyParamsBody struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address param.Field[string] `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host param.Field[string] `json:"host"`
}

func (r AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
