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
func (r *DevicePolicyExcludeService) List(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *[]DevicePolicyExcludeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyExcludeListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/exclude", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Sets the list of routes excluded from the WARP client's tunnel.
func (r *DevicePolicyExcludeService) Replace(ctx context.Context, identifier interface{}, body DevicePolicyExcludeReplaceParams, opts ...option.RequestOption) (res *[]DevicePolicyExcludeReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyExcludeReplaceResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/exclude", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DevicePolicyExcludeListResponse struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host string                              `json:"host"`
	JSON devicePolicyExcludeListResponseJSON `json:"-"`
}

// devicePolicyExcludeListResponseJSON contains the JSON metadata for the struct
// [DevicePolicyExcludeListResponse]
type devicePolicyExcludeListResponseJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeReplaceResponse struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host string                                 `json:"host"`
	JSON devicePolicyExcludeReplaceResponseJSON `json:"-"`
}

// devicePolicyExcludeReplaceResponseJSON contains the JSON metadata for the struct
// [DevicePolicyExcludeReplaceResponse]
type devicePolicyExcludeReplaceResponseJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeListResponseEnvelope struct {
	Errors   []DevicePolicyExcludeListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyExcludeListResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyExcludeListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyExcludeListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyExcludeListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyExcludeListResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyExcludeListResponseEnvelopeJSON contains the JSON metadata for the
// struct [DevicePolicyExcludeListResponseEnvelope]
type devicePolicyExcludeListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeListResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    devicePolicyExcludeListResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyExcludeListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DevicePolicyExcludeListResponseEnvelopeErrors]
type devicePolicyExcludeListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeListResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    devicePolicyExcludeListResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyExcludeListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DevicePolicyExcludeListResponseEnvelopeMessages]
type devicePolicyExcludeListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePolicyExcludeListResponseEnvelopeSuccess bool

const (
	DevicePolicyExcludeListResponseEnvelopeSuccessTrue DevicePolicyExcludeListResponseEnvelopeSuccess = true
)

type DevicePolicyExcludeListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                               `json:"total_count"`
	JSON       devicePolicyExcludeListResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyExcludeListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [DevicePolicyExcludeListResponseEnvelopeResultInfo]
type devicePolicyExcludeListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeReplaceParams struct {
	Body param.Field[[]DevicePolicyExcludeReplaceParamsBody] `json:"body,required"`
}

func (r DevicePolicyExcludeReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DevicePolicyExcludeReplaceParamsBody struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address param.Field[string] `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host param.Field[string] `json:"host"`
}

func (r DevicePolicyExcludeReplaceParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyExcludeReplaceResponseEnvelope struct {
	Errors   []DevicePolicyExcludeReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyExcludeReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyExcludeReplaceResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyExcludeReplaceResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyExcludeReplaceResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyExcludeReplaceResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyExcludeReplaceResponseEnvelopeJSON contains the JSON metadata for
// the struct [DevicePolicyExcludeReplaceResponseEnvelope]
type devicePolicyExcludeReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeReplaceResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    devicePolicyExcludeReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyExcludeReplaceResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DevicePolicyExcludeReplaceResponseEnvelopeErrors]
type devicePolicyExcludeReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeReplaceResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    devicePolicyExcludeReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyExcludeReplaceResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DevicePolicyExcludeReplaceResponseEnvelopeMessages]
type devicePolicyExcludeReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePolicyExcludeReplaceResponseEnvelopeSuccess bool

const (
	DevicePolicyExcludeReplaceResponseEnvelopeSuccessTrue DevicePolicyExcludeReplaceResponseEnvelopeSuccess = true
)

type DevicePolicyExcludeReplaceResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                  `json:"total_count"`
	JSON       devicePolicyExcludeReplaceResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyExcludeReplaceResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [DevicePolicyExcludeReplaceResponseEnvelopeResultInfo]
type devicePolicyExcludeReplaceResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeReplaceResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
