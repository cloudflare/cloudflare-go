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

// Fetches the list of domains to bypass Gateway DNS resolution from a specified
// device settings profile. These domains will use the specified local DNS resolver
// instead.
func (r *DevicePolicyFallbackDomainService) List(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *[]DevicePolicyFallbackDomainListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyFallbackDomainListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/fallback_domains", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Sets the list of domains to bypass Gateway DNS resolution. These domains will
// use the specified local DNS resolver instead. This will only apply to the
// specified device settings profile.
func (r *DevicePolicyFallbackDomainService) Replace(ctx context.Context, identifier interface{}, uuid string, body DevicePolicyFallbackDomainReplaceParams, opts ...option.RequestOption) (res *[]DevicePolicyFallbackDomainReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyFallbackDomainReplaceResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/fallback_domains", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DevicePolicyFallbackDomainListResponse struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                              `json:"dns_server"`
	JSON      devicePolicyFallbackDomainListResponseJSON `json:"-"`
}

// devicePolicyFallbackDomainListResponseJSON contains the JSON metadata for the
// struct [DevicePolicyFallbackDomainListResponse]
type devicePolicyFallbackDomainListResponseJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainReplaceResponse struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                                 `json:"dns_server"`
	JSON      devicePolicyFallbackDomainReplaceResponseJSON `json:"-"`
}

// devicePolicyFallbackDomainReplaceResponseJSON contains the JSON metadata for the
// struct [DevicePolicyFallbackDomainReplaceResponse]
type devicePolicyFallbackDomainReplaceResponseJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainListResponseEnvelope struct {
	Errors   []DevicePolicyFallbackDomainListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyFallbackDomainListResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyFallbackDomainListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyFallbackDomainListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyFallbackDomainListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyFallbackDomainListResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyFallbackDomainListResponseEnvelopeJSON contains the JSON metadata
// for the struct [DevicePolicyFallbackDomainListResponseEnvelope]
type devicePolicyFallbackDomainListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainListResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    devicePolicyFallbackDomainListResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyFallbackDomainListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DevicePolicyFallbackDomainListResponseEnvelopeErrors]
type devicePolicyFallbackDomainListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainListResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    devicePolicyFallbackDomainListResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyFallbackDomainListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DevicePolicyFallbackDomainListResponseEnvelopeMessages]
type devicePolicyFallbackDomainListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePolicyFallbackDomainListResponseEnvelopeSuccess bool

const (
	DevicePolicyFallbackDomainListResponseEnvelopeSuccessTrue DevicePolicyFallbackDomainListResponseEnvelopeSuccess = true
)

type DevicePolicyFallbackDomainListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                      `json:"total_count"`
	JSON       devicePolicyFallbackDomainListResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyFallbackDomainListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [DevicePolicyFallbackDomainListResponseEnvelopeResultInfo]
type devicePolicyFallbackDomainListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainReplaceParams struct {
	Body param.Field[[]DevicePolicyFallbackDomainReplaceParamsBody] `json:"body,required"`
}

func (r DevicePolicyFallbackDomainReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DevicePolicyFallbackDomainReplaceParamsBody struct {
	// The domain suffix to match when resolving locally.
	Suffix param.Field[string] `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description param.Field[string] `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer param.Field[[]interface{}] `json:"dns_server"`
}

func (r DevicePolicyFallbackDomainReplaceParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyFallbackDomainReplaceResponseEnvelope struct {
	Errors   []DevicePolicyFallbackDomainReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyFallbackDomainReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyFallbackDomainReplaceResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyFallbackDomainReplaceResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyFallbackDomainReplaceResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyFallbackDomainReplaceResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyFallbackDomainReplaceResponseEnvelopeJSON contains the JSON metadata
// for the struct [DevicePolicyFallbackDomainReplaceResponseEnvelope]
type devicePolicyFallbackDomainReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainReplaceResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    devicePolicyFallbackDomainReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyFallbackDomainReplaceResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [DevicePolicyFallbackDomainReplaceResponseEnvelopeErrors]
type devicePolicyFallbackDomainReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainReplaceResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    devicePolicyFallbackDomainReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyFallbackDomainReplaceResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [DevicePolicyFallbackDomainReplaceResponseEnvelopeMessages]
type devicePolicyFallbackDomainReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePolicyFallbackDomainReplaceResponseEnvelopeSuccess bool

const (
	DevicePolicyFallbackDomainReplaceResponseEnvelopeSuccessTrue DevicePolicyFallbackDomainReplaceResponseEnvelopeSuccess = true
)

type DevicePolicyFallbackDomainReplaceResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                         `json:"total_count"`
	JSON       devicePolicyFallbackDomainReplaceResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyFallbackDomainReplaceResponseEnvelopeResultInfoJSON contains the
// JSON metadata for the struct
// [DevicePolicyFallbackDomainReplaceResponseEnvelopeResultInfo]
type devicePolicyFallbackDomainReplaceResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainReplaceResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
