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

// DevicePolicyFallbackDomainService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewDevicePolicyFallbackDomainService] method instead.
type DevicePolicyFallbackDomainService struct {
	Options []option.RequestOption
}

// NewDevicePolicyFallbackDomainService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDevicePolicyFallbackDomainService(opts ...option.RequestOption) (r *DevicePolicyFallbackDomainService) {
	r = &DevicePolicyFallbackDomainService{}
	r.Options = opts
	return
}

// Fetches a list of domains to bypass Gateway DNS resolution. These domains will
// use the specified local DNS resolver instead.
func (r *DevicePolicyFallbackDomainService) DevicesGetLocalDomainFallbackList(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *[]DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/fallback_domains", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the list of domains to bypass Gateway DNS resolution from a specified
// device settings profile. These domains will use the specified local DNS resolver
// instead.
func (r *DevicePolicyFallbackDomainService) DevicesGetLocalDomainFallbackListForADeviceSettingsPolicy(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *[]DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/fallback_domains", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Sets the list of domains to bypass Gateway DNS resolution. These domains will
// use the specified local DNS resolver instead.
func (r *DevicePolicyFallbackDomainService) DevicesSetLocalDomainFallbackList(ctx context.Context, identifier interface{}, body DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListParams, opts ...option.RequestOption) (res *[]DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/fallback_domains", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Sets the list of domains to bypass Gateway DNS resolution. These domains will
// use the specified local DNS resolver instead. This will only apply to the
// specified device settings profile.
func (r *DevicePolicyFallbackDomainService) DevicesSetLocalDomainFallbackListForADeviceSettingsPolicy(ctx context.Context, identifier interface{}, uuid string, body DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyParams, opts ...option.RequestOption) (res *[]DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/fallback_domains", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponse struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                                                           `json:"dns_server"`
	JSON      devicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseJSON `json:"-"`
}

// devicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseJSON contains
// the JSON metadata for the struct
// [DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponse]
type devicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponse struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                                                                                   `json:"dns_server"`
	JSON      devicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseJSON `json:"-"`
}

// devicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseJSON
// contains the JSON metadata for the struct
// [DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponse]
type devicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponse struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                                                           `json:"dns_server"`
	JSON      devicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseJSON `json:"-"`
}

// devicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseJSON contains
// the JSON metadata for the struct
// [DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponse]
type devicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponse struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                                                                                   `json:"dns_server"`
	JSON      devicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseJSON `json:"-"`
}

// devicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseJSON
// contains the JSON metadata for the struct
// [DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponse]
type devicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelope struct {
	Errors   []DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelope]
type devicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelopeErrors struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    devicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelopeErrors]
type devicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelopeMessages struct {
	Code    int64                                                                                   `json:"code,required"`
	Message string                                                                                  `json:"message,required"`
	JSON    devicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelopeMessages]
type devicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelopeSuccess bool

const (
	DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelopeSuccessTrue DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelopeSuccess = true
)

type DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                   `json:"total_count"`
	JSON       devicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelopeResultInfo]
type devicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelope struct {
	Errors   []DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelope]
type devicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeErrors struct {
	Code    int64                                                                                                         `json:"code,required"`
	Message string                                                                                                        `json:"message,required"`
	JSON    devicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeErrors]
type devicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeMessages struct {
	Code    int64                                                                                                           `json:"code,required"`
	Message string                                                                                                          `json:"message,required"`
	JSON    devicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeMessages]
type devicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeSuccess bool

const (
	DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeSuccessTrue DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeSuccess = true
)

type DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                           `json:"total_count"`
	JSON       devicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeResultInfo]
type devicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainDevicesGetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListParams struct {
	Body param.Field[[]DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListParamsBody] `json:"body,required"`
}

func (r DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListParamsBody struct {
	// The domain suffix to match when resolving locally.
	Suffix param.Field[string] `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description param.Field[string] `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer param.Field[[]interface{}] `json:"dns_server"`
}

func (r DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelope struct {
	Errors   []DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelope]
type devicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelopeErrors struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    devicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelopeErrors]
type devicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelopeMessages struct {
	Code    int64                                                                                   `json:"code,required"`
	Message string                                                                                  `json:"message,required"`
	JSON    devicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelopeMessages]
type devicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelopeSuccess bool

const (
	DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelopeSuccessTrue DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelopeSuccess = true
)

type DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                   `json:"total_count"`
	JSON       devicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelopeResultInfo]
type devicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyParams struct {
	Body param.Field[[]DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyParamsBody] `json:"body,required"`
}

func (r DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyParamsBody struct {
	// The domain suffix to match when resolving locally.
	Suffix param.Field[string] `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description param.Field[string] `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer param.Field[[]interface{}] `json:"dns_server"`
}

func (r DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelope struct {
	Errors   []DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelope]
type devicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeErrors struct {
	Code    int64                                                                                                         `json:"code,required"`
	Message string                                                                                                        `json:"message,required"`
	JSON    devicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeErrors]
type devicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeMessages struct {
	Code    int64                                                                                                           `json:"code,required"`
	Message string                                                                                                          `json:"message,required"`
	JSON    devicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeMessages]
type devicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeSuccess bool

const (
	DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeSuccessTrue DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeSuccess = true
)

type DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                           `json:"total_count"`
	JSON       devicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeResultInfo]
type devicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
