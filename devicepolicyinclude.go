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
func (r *DevicePolicyIncludeService) List(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *[]DevicePolicyIncludeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyIncludeListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/include", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Sets the list of routes included in the WARP client's tunnel.
func (r *DevicePolicyIncludeService) Replace(ctx context.Context, identifier interface{}, body DevicePolicyIncludeReplaceParams, opts ...option.RequestOption) (res *[]DevicePolicyIncludeReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyIncludeReplaceResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/include", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DevicePolicyIncludeListResponse struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                              `json:"host"`
	JSON devicePolicyIncludeListResponseJSON `json:"-"`
}

// devicePolicyIncludeListResponseJSON contains the JSON metadata for the struct
// [DevicePolicyIncludeListResponse]
type devicePolicyIncludeListResponseJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyIncludeReplaceResponse struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                                 `json:"host"`
	JSON devicePolicyIncludeReplaceResponseJSON `json:"-"`
}

// devicePolicyIncludeReplaceResponseJSON contains the JSON metadata for the struct
// [DevicePolicyIncludeReplaceResponse]
type devicePolicyIncludeReplaceResponseJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyIncludeListResponseEnvelope struct {
	Errors   []DevicePolicyIncludeListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyIncludeListResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyIncludeListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyIncludeListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyIncludeListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyIncludeListResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyIncludeListResponseEnvelopeJSON contains the JSON metadata for the
// struct [DevicePolicyIncludeListResponseEnvelope]
type devicePolicyIncludeListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyIncludeListResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    devicePolicyIncludeListResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyIncludeListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DevicePolicyIncludeListResponseEnvelopeErrors]
type devicePolicyIncludeListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyIncludeListResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    devicePolicyIncludeListResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyIncludeListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DevicePolicyIncludeListResponseEnvelopeMessages]
type devicePolicyIncludeListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePolicyIncludeListResponseEnvelopeSuccess bool

const (
	DevicePolicyIncludeListResponseEnvelopeSuccessTrue DevicePolicyIncludeListResponseEnvelopeSuccess = true
)

type DevicePolicyIncludeListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                               `json:"total_count"`
	JSON       devicePolicyIncludeListResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyIncludeListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [DevicePolicyIncludeListResponseEnvelopeResultInfo]
type devicePolicyIncludeListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyIncludeReplaceParams struct {
	Body param.Field[[]DevicePolicyIncludeReplaceParamsBody] `json:"body,required"`
}

func (r DevicePolicyIncludeReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DevicePolicyIncludeReplaceParamsBody struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address param.Field[string] `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host param.Field[string] `json:"host"`
}

func (r DevicePolicyIncludeReplaceParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyIncludeReplaceResponseEnvelope struct {
	Errors   []DevicePolicyIncludeReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyIncludeReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyIncludeReplaceResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyIncludeReplaceResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyIncludeReplaceResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyIncludeReplaceResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyIncludeReplaceResponseEnvelopeJSON contains the JSON metadata for
// the struct [DevicePolicyIncludeReplaceResponseEnvelope]
type devicePolicyIncludeReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyIncludeReplaceResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    devicePolicyIncludeReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyIncludeReplaceResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DevicePolicyIncludeReplaceResponseEnvelopeErrors]
type devicePolicyIncludeReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyIncludeReplaceResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    devicePolicyIncludeReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyIncludeReplaceResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DevicePolicyIncludeReplaceResponseEnvelopeMessages]
type devicePolicyIncludeReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePolicyIncludeReplaceResponseEnvelopeSuccess bool

const (
	DevicePolicyIncludeReplaceResponseEnvelopeSuccessTrue DevicePolicyIncludeReplaceResponseEnvelopeSuccess = true
)

type DevicePolicyIncludeReplaceResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                  `json:"total_count"`
	JSON       devicePolicyIncludeReplaceResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyIncludeReplaceResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [DevicePolicyIncludeReplaceResponseEnvelopeResultInfo]
type devicePolicyIncludeReplaceResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeReplaceResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
