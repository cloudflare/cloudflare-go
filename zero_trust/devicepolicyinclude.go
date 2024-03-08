// File generated from our OpenAPI spec by Stainless.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
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

// Sets the list of routes included in the WARP client's tunnel.
func (r *DevicePolicyIncludeService) Update(ctx context.Context, params DevicePolicyIncludeUpdateParams, opts ...option.RequestOption) (res *[]DevicePolicyIncludeUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyIncludeUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/include", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the list of routes included in the WARP client's tunnel.
func (r *DevicePolicyIncludeService) List(ctx context.Context, query DevicePolicyIncludeListParams, opts ...option.RequestOption) (res *[]DevicePolicyIncludeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyIncludeListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/include", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the list of routes included in the WARP client's tunnel for a specific
// device settings profile.
func (r *DevicePolicyIncludeService) Get(ctx context.Context, policyID string, query DevicePolicyIncludeGetParams, opts ...option.RequestOption) (res *[]DevicePolicyIncludeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyIncludeGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/include", query.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DevicePolicyIncludeUpdateResponse struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                                `json:"host"`
	JSON devicePolicyIncludeUpdateResponseJSON `json:"-"`
}

// devicePolicyIncludeUpdateResponseJSON contains the JSON metadata for the struct
// [DevicePolicyIncludeUpdateResponse]
type devicePolicyIncludeUpdateResponseJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyIncludeUpdateResponseJSON) RawJSON() string {
	return r.raw
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

func (r devicePolicyIncludeListResponseJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyIncludeGetResponse struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                             `json:"host"`
	JSON devicePolicyIncludeGetResponseJSON `json:"-"`
}

// devicePolicyIncludeGetResponseJSON contains the JSON metadata for the struct
// [DevicePolicyIncludeGetResponse]
type devicePolicyIncludeGetResponseJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyIncludeGetResponseJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyIncludeUpdateParams struct {
	AccountID param.Field[interface{}]                           `path:"account_id,required"`
	Body      param.Field[[]DevicePolicyIncludeUpdateParamsBody] `json:"body,required"`
}

func (r DevicePolicyIncludeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DevicePolicyIncludeUpdateParamsBody struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address param.Field[string] `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host param.Field[string] `json:"host"`
}

func (r DevicePolicyIncludeUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyIncludeUpdateResponseEnvelope struct {
	Errors   []DevicePolicyIncludeUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyIncludeUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyIncludeUpdateResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyIncludeUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyIncludeUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyIncludeUpdateResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyIncludeUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DevicePolicyIncludeUpdateResponseEnvelope]
type devicePolicyIncludeUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyIncludeUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyIncludeUpdateResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    devicePolicyIncludeUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyIncludeUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DevicePolicyIncludeUpdateResponseEnvelopeErrors]
type devicePolicyIncludeUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyIncludeUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyIncludeUpdateResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    devicePolicyIncludeUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyIncludeUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DevicePolicyIncludeUpdateResponseEnvelopeMessages]
type devicePolicyIncludeUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyIncludeUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyIncludeUpdateResponseEnvelopeSuccess bool

const (
	DevicePolicyIncludeUpdateResponseEnvelopeSuccessTrue DevicePolicyIncludeUpdateResponseEnvelopeSuccess = true
)

type DevicePolicyIncludeUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                 `json:"total_count"`
	JSON       devicePolicyIncludeUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyIncludeUpdateResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [DevicePolicyIncludeUpdateResponseEnvelopeResultInfo]
type devicePolicyIncludeUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyIncludeUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyIncludeListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
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

func (r devicePolicyIncludeListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r devicePolicyIncludeListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r devicePolicyIncludeListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r devicePolicyIncludeListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyIncludeGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type DevicePolicyIncludeGetResponseEnvelope struct {
	Errors   []DevicePolicyIncludeGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyIncludeGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyIncludeGetResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyIncludeGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyIncludeGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyIncludeGetResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyIncludeGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DevicePolicyIncludeGetResponseEnvelope]
type devicePolicyIncludeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyIncludeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyIncludeGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    devicePolicyIncludeGetResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyIncludeGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DevicePolicyIncludeGetResponseEnvelopeErrors]
type devicePolicyIncludeGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyIncludeGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyIncludeGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    devicePolicyIncludeGetResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyIncludeGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DevicePolicyIncludeGetResponseEnvelopeMessages]
type devicePolicyIncludeGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyIncludeGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyIncludeGetResponseEnvelopeSuccess bool

const (
	DevicePolicyIncludeGetResponseEnvelopeSuccessTrue DevicePolicyIncludeGetResponseEnvelopeSuccess = true
)

type DevicePolicyIncludeGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                              `json:"total_count"`
	JSON       devicePolicyIncludeGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyIncludeGetResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [DevicePolicyIncludeGetResponseEnvelopeResultInfo]
type devicePolicyIncludeGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyIncludeGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
