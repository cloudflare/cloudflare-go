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

// DevicePolicyExcludeService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDevicePolicyExcludeService]
// method instead.
type DevicePolicyExcludeService struct {
	Options []option.RequestOption
}

// NewDevicePolicyExcludeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDevicePolicyExcludeService(opts ...option.RequestOption) (r *DevicePolicyExcludeService) {
	r = &DevicePolicyExcludeService{}
	r.Options = opts
	return
}

// Fetches the list of routes excluded from the WARP client's tunnel.
func (r *DevicePolicyExcludeService) DevicesGetSplitTunnelExcludeList(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *[]DevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/exclude", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the list of routes excluded from the WARP client's tunnel for a specific
// device settings profile.
func (r *DevicePolicyExcludeService) DevicesGetSplitTunnelExcludeListForADeviceSettingsPolicy(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *[]DevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/exclude", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Sets the list of routes excluded from the WARP client's tunnel.
func (r *DevicePolicyExcludeService) DevicesSetSplitTunnelExcludeList(ctx context.Context, identifier interface{}, body DevicePolicyExcludeDevicesSetSplitTunnelExcludeListParams, opts ...option.RequestOption) (res *[]DevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/exclude", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Sets the list of routes excluded from the WARP client's tunnel for a specific
// device settings profile.
func (r *DevicePolicyExcludeService) DevicesSetSplitTunnelExcludeListForADeviceSettingsPolicy(ctx context.Context, identifier interface{}, uuid string, body DevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyParams, opts ...option.RequestOption) (res *[]DevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/exclude", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponse struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host string                                                          `json:"host"`
	JSON devicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseJSON `json:"-"`
}

// devicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseJSON contains the
// JSON metadata for the struct
// [DevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponse]
type devicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponse struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host string                                                                                  `json:"host"`
	JSON devicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseJSON `json:"-"`
}

// devicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseJSON
// contains the JSON metadata for the struct
// [DevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponse]
type devicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponse struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host string                                                          `json:"host"`
	JSON devicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseJSON `json:"-"`
}

// devicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseJSON contains the
// JSON metadata for the struct
// [DevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponse]
type devicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponse struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host string                                                                                  `json:"host"`
	JSON devicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseJSON `json:"-"`
}

// devicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseJSON
// contains the JSON metadata for the struct
// [DevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponse]
type devicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelope struct {
	Errors   []DevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [DevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelope]
type devicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelopeErrors struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    devicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [DevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelopeErrors]
type devicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelopeMessages struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    devicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [DevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelopeMessages]
type devicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelopeSuccess bool

const (
	DevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelopeSuccessTrue DevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelopeSuccess = true
)

type DevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                           `json:"total_count"`
	JSON       devicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [DevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelopeResultInfo]
type devicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeDevicesGetSplitTunnelExcludeListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelope struct {
	Errors   []DevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [DevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelope]
type devicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeErrors struct {
	Code    int64                                                                                                 `json:"code,required"`
	Message string                                                                                                `json:"message,required"`
	JSON    devicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [DevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeErrors]
type devicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeMessages struct {
	Code    int64                                                                                                   `json:"code,required"`
	Message string                                                                                                  `json:"message,required"`
	JSON    devicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [DevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeMessages]
type devicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeSuccess bool

const (
	DevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeSuccessTrue DevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeSuccess = true
)

type DevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                   `json:"total_count"`
	JSON       devicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [DevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeResultInfo]
type devicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeDevicesGetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeDevicesSetSplitTunnelExcludeListParams struct {
	Body param.Field[[]DevicePolicyExcludeDevicesSetSplitTunnelExcludeListParamsBody] `json:"body,required"`
}

func (r DevicePolicyExcludeDevicesSetSplitTunnelExcludeListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DevicePolicyExcludeDevicesSetSplitTunnelExcludeListParamsBody struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address param.Field[string] `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host param.Field[string] `json:"host"`
}

func (r DevicePolicyExcludeDevicesSetSplitTunnelExcludeListParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelope struct {
	Errors   []DevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [DevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelope]
type devicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelopeErrors struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    devicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [DevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelopeErrors]
type devicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelopeMessages struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    devicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [DevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelopeMessages]
type devicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelopeSuccess bool

const (
	DevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelopeSuccessTrue DevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelopeSuccess = true
)

type DevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                           `json:"total_count"`
	JSON       devicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [DevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelopeResultInfo]
type devicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeDevicesSetSplitTunnelExcludeListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyParams struct {
	Body param.Field[[]DevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyParamsBody] `json:"body,required"`
}

func (r DevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyParamsBody struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address param.Field[string] `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host param.Field[string] `json:"host"`
}

func (r DevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelope struct {
	Errors   []DevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [DevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelope]
type devicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeErrors struct {
	Code    int64                                                                                                 `json:"code,required"`
	Message string                                                                                                `json:"message,required"`
	JSON    devicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [DevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeErrors]
type devicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeMessages struct {
	Code    int64                                                                                                   `json:"code,required"`
	Message string                                                                                                  `json:"message,required"`
	JSON    devicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [DevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeMessages]
type devicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeSuccess bool

const (
	DevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeSuccessTrue DevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeSuccess = true
)

type DevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                   `json:"total_count"`
	JSON       devicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [DevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeResultInfo]
type devicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
