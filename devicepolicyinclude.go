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

// DevicePolicyIncludeService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDevicePolicyIncludeService]
// method instead.
type DevicePolicyIncludeService struct {
	Options []option.RequestOption
}

// NewDevicePolicyIncludeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDevicePolicyIncludeService(opts ...option.RequestOption) (r *DevicePolicyIncludeService) {
	r = &DevicePolicyIncludeService{}
	r.Options = opts
	return
}

// Fetches the list of routes included in the WARP client's tunnel.
func (r *DevicePolicyIncludeService) DevicesGetSplitTunnelIncludeList(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *[]DevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/include", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the list of routes included in the WARP client's tunnel for a specific
// device settings profile.
func (r *DevicePolicyIncludeService) DevicesGetSplitTunnelIncludeListForADeviceSettingsPolicy(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *[]DevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/include", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Sets the list of routes included in the WARP client's tunnel.
func (r *DevicePolicyIncludeService) DevicesSetSplitTunnelIncludeList(ctx context.Context, identifier interface{}, body DevicePolicyIncludeDevicesSetSplitTunnelIncludeListParams, opts ...option.RequestOption) (res *[]DevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/include", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Sets the list of routes included in the WARP client's tunnel for a specific
// device settings profile.
func (r *DevicePolicyIncludeService) DevicesSetSplitTunnelIncludeListForADeviceSettingsPolicy(ctx context.Context, identifier interface{}, uuid string, body DevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyParams, opts ...option.RequestOption) (res *[]DevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/include", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponse struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                                                          `json:"host"`
	JSON devicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseJSON `json:"-"`
}

// devicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseJSON contains the
// JSON metadata for the struct
// [DevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponse]
type devicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponse struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                                                                                  `json:"host"`
	JSON devicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseJSON `json:"-"`
}

// devicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseJSON
// contains the JSON metadata for the struct
// [DevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponse]
type devicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponse struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                                                          `json:"host"`
	JSON devicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseJSON `json:"-"`
}

// devicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseJSON contains the
// JSON metadata for the struct
// [DevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponse]
type devicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponse struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                                                                                  `json:"host"`
	JSON devicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseJSON `json:"-"`
}

// devicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseJSON
// contains the JSON metadata for the struct
// [DevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponse]
type devicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelope struct {
	Errors   []DevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [DevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelope]
type devicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelopeErrors struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    devicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [DevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelopeErrors]
type devicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelopeMessages struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    devicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [DevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelopeMessages]
type devicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelopeSuccess bool

const (
	DevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelopeSuccessTrue DevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelopeSuccess = true
)

type DevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                           `json:"total_count"`
	JSON       devicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [DevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelopeResultInfo]
type devicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeDevicesGetSplitTunnelIncludeListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelope struct {
	Errors   []DevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [DevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelope]
type devicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeErrors struct {
	Code    int64                                                                                                 `json:"code,required"`
	Message string                                                                                                `json:"message,required"`
	JSON    devicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [DevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeErrors]
type devicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeMessages struct {
	Code    int64                                                                                                   `json:"code,required"`
	Message string                                                                                                  `json:"message,required"`
	JSON    devicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [DevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeMessages]
type devicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeSuccess bool

const (
	DevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeSuccessTrue DevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeSuccess = true
)

type DevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                   `json:"total_count"`
	JSON       devicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [DevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeResultInfo]
type devicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeDevicesGetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyIncludeDevicesSetSplitTunnelIncludeListParams struct {
	Body param.Field[[]DevicePolicyIncludeDevicesSetSplitTunnelIncludeListParamsBody] `json:"body,required"`
}

func (r DevicePolicyIncludeDevicesSetSplitTunnelIncludeListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DevicePolicyIncludeDevicesSetSplitTunnelIncludeListParamsBody struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address param.Field[string] `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host param.Field[string] `json:"host"`
}

func (r DevicePolicyIncludeDevicesSetSplitTunnelIncludeListParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelope struct {
	Errors   []DevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [DevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelope]
type devicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelopeErrors struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    devicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [DevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelopeErrors]
type devicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelopeMessages struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    devicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [DevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelopeMessages]
type devicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelopeSuccess bool

const (
	DevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelopeSuccessTrue DevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelopeSuccess = true
)

type DevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                           `json:"total_count"`
	JSON       devicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [DevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelopeResultInfo]
type devicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeDevicesSetSplitTunnelIncludeListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyParams struct {
	Body param.Field[[]DevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyParamsBody] `json:"body,required"`
}

func (r DevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyParamsBody struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address param.Field[string] `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host param.Field[string] `json:"host"`
}

func (r DevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelope struct {
	Errors   []DevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [DevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelope]
type devicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeErrors struct {
	Code    int64                                                                                                 `json:"code,required"`
	Message string                                                                                                `json:"message,required"`
	JSON    devicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [DevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeErrors]
type devicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeMessages struct {
	Code    int64                                                                                                   `json:"code,required"`
	Message string                                                                                                  `json:"message,required"`
	JSON    devicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [DevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeMessages]
type devicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeSuccess bool

const (
	DevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeSuccessTrue DevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeSuccess = true
)

type DevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                   `json:"total_count"`
	JSON       devicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [DevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeResultInfo]
type devicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
