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

// ZeroTrustDevicePolicyIncludeService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZeroTrustDevicePolicyIncludeService] method instead.
type ZeroTrustDevicePolicyIncludeService struct {
	Options []option.RequestOption
}

// NewZeroTrustDevicePolicyIncludeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustDevicePolicyIncludeService(opts ...option.RequestOption) (r *ZeroTrustDevicePolicyIncludeService) {
	r = &ZeroTrustDevicePolicyIncludeService{}
	r.Options = opts
	return
}

// Sets the list of routes included in the WARP client's tunnel.
func (r *ZeroTrustDevicePolicyIncludeService) Update(ctx context.Context, params ZeroTrustDevicePolicyIncludeUpdateParams, opts ...option.RequestOption) (res *[]ZeroTrustDevicePolicyIncludeUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDevicePolicyIncludeUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/include", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the list of routes included in the WARP client's tunnel.
func (r *ZeroTrustDevicePolicyIncludeService) List(ctx context.Context, query ZeroTrustDevicePolicyIncludeListParams, opts ...option.RequestOption) (res *[]ZeroTrustDevicePolicyIncludeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDevicePolicyIncludeListResponseEnvelope
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
func (r *ZeroTrustDevicePolicyIncludeService) Get(ctx context.Context, policyID string, query ZeroTrustDevicePolicyIncludeGetParams, opts ...option.RequestOption) (res *[]ZeroTrustDevicePolicyIncludeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDevicePolicyIncludeGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/include", query.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustDevicePolicyIncludeUpdateResponse struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                                         `json:"host"`
	JSON zeroTrustDevicePolicyIncludeUpdateResponseJSON `json:"-"`
}

// zeroTrustDevicePolicyIncludeUpdateResponseJSON contains the JSON metadata for
// the struct [ZeroTrustDevicePolicyIncludeUpdateResponse]
type zeroTrustDevicePolicyIncludeUpdateResponseJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyIncludeUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicePolicyIncludeUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDevicePolicyIncludeListResponse struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                                       `json:"host"`
	JSON zeroTrustDevicePolicyIncludeListResponseJSON `json:"-"`
}

// zeroTrustDevicePolicyIncludeListResponseJSON contains the JSON metadata for the
// struct [ZeroTrustDevicePolicyIncludeListResponse]
type zeroTrustDevicePolicyIncludeListResponseJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyIncludeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicePolicyIncludeListResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDevicePolicyIncludeGetResponse struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                                      `json:"host"`
	JSON zeroTrustDevicePolicyIncludeGetResponseJSON `json:"-"`
}

// zeroTrustDevicePolicyIncludeGetResponseJSON contains the JSON metadata for the
// struct [ZeroTrustDevicePolicyIncludeGetResponse]
type zeroTrustDevicePolicyIncludeGetResponseJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyIncludeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicePolicyIncludeGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDevicePolicyIncludeUpdateParams struct {
	AccountID param.Field[interface{}]                                    `path:"account_id,required"`
	Body      param.Field[[]ZeroTrustDevicePolicyIncludeUpdateParamsBody] `json:"body,required"`
}

func (r ZeroTrustDevicePolicyIncludeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZeroTrustDevicePolicyIncludeUpdateParamsBody struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address param.Field[string] `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host param.Field[string] `json:"host"`
}

func (r ZeroTrustDevicePolicyIncludeUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustDevicePolicyIncludeUpdateResponseEnvelope struct {
	Errors   []ZeroTrustDevicePolicyIncludeUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDevicePolicyIncludeUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustDevicePolicyIncludeUpdateResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    ZeroTrustDevicePolicyIncludeUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustDevicePolicyIncludeUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustDevicePolicyIncludeUpdateResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustDevicePolicyIncludeUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustDevicePolicyIncludeUpdateResponseEnvelope]
type zeroTrustDevicePolicyIncludeUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyIncludeUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicePolicyIncludeUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDevicePolicyIncludeUpdateResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zeroTrustDevicePolicyIncludeUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDevicePolicyIncludeUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustDevicePolicyIncludeUpdateResponseEnvelopeErrors]
type zeroTrustDevicePolicyIncludeUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyIncludeUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicePolicyIncludeUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDevicePolicyIncludeUpdateResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    zeroTrustDevicePolicyIncludeUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDevicePolicyIncludeUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustDevicePolicyIncludeUpdateResponseEnvelopeMessages]
type zeroTrustDevicePolicyIncludeUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyIncludeUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicePolicyIncludeUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZeroTrustDevicePolicyIncludeUpdateResponseEnvelopeSuccess bool

const (
	ZeroTrustDevicePolicyIncludeUpdateResponseEnvelopeSuccessTrue ZeroTrustDevicePolicyIncludeUpdateResponseEnvelopeSuccess = true
)

type ZeroTrustDevicePolicyIncludeUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                          `json:"total_count"`
	JSON       zeroTrustDevicePolicyIncludeUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustDevicePolicyIncludeUpdateResponseEnvelopeResultInfoJSON contains the
// JSON metadata for the struct
// [ZeroTrustDevicePolicyIncludeUpdateResponseEnvelopeResultInfo]
type zeroTrustDevicePolicyIncludeUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyIncludeUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicePolicyIncludeUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDevicePolicyIncludeListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustDevicePolicyIncludeListResponseEnvelope struct {
	Errors   []ZeroTrustDevicePolicyIncludeListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDevicePolicyIncludeListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustDevicePolicyIncludeListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    ZeroTrustDevicePolicyIncludeListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustDevicePolicyIncludeListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustDevicePolicyIncludeListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustDevicePolicyIncludeListResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustDevicePolicyIncludeListResponseEnvelope]
type zeroTrustDevicePolicyIncludeListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyIncludeListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicePolicyIncludeListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDevicePolicyIncludeListResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zeroTrustDevicePolicyIncludeListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDevicePolicyIncludeListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustDevicePolicyIncludeListResponseEnvelopeErrors]
type zeroTrustDevicePolicyIncludeListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyIncludeListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicePolicyIncludeListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDevicePolicyIncludeListResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zeroTrustDevicePolicyIncludeListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDevicePolicyIncludeListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustDevicePolicyIncludeListResponseEnvelopeMessages]
type zeroTrustDevicePolicyIncludeListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyIncludeListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicePolicyIncludeListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZeroTrustDevicePolicyIncludeListResponseEnvelopeSuccess bool

const (
	ZeroTrustDevicePolicyIncludeListResponseEnvelopeSuccessTrue ZeroTrustDevicePolicyIncludeListResponseEnvelopeSuccess = true
)

type ZeroTrustDevicePolicyIncludeListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                        `json:"total_count"`
	JSON       zeroTrustDevicePolicyIncludeListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustDevicePolicyIncludeListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [ZeroTrustDevicePolicyIncludeListResponseEnvelopeResultInfo]
type zeroTrustDevicePolicyIncludeListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyIncludeListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicePolicyIncludeListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDevicePolicyIncludeGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustDevicePolicyIncludeGetResponseEnvelope struct {
	Errors   []ZeroTrustDevicePolicyIncludeGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDevicePolicyIncludeGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustDevicePolicyIncludeGetResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    ZeroTrustDevicePolicyIncludeGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustDevicePolicyIncludeGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustDevicePolicyIncludeGetResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustDevicePolicyIncludeGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustDevicePolicyIncludeGetResponseEnvelope]
type zeroTrustDevicePolicyIncludeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyIncludeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicePolicyIncludeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDevicePolicyIncludeGetResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zeroTrustDevicePolicyIncludeGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDevicePolicyIncludeGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustDevicePolicyIncludeGetResponseEnvelopeErrors]
type zeroTrustDevicePolicyIncludeGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyIncludeGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicePolicyIncludeGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDevicePolicyIncludeGetResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zeroTrustDevicePolicyIncludeGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDevicePolicyIncludeGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustDevicePolicyIncludeGetResponseEnvelopeMessages]
type zeroTrustDevicePolicyIncludeGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyIncludeGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicePolicyIncludeGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZeroTrustDevicePolicyIncludeGetResponseEnvelopeSuccess bool

const (
	ZeroTrustDevicePolicyIncludeGetResponseEnvelopeSuccessTrue ZeroTrustDevicePolicyIncludeGetResponseEnvelopeSuccess = true
)

type ZeroTrustDevicePolicyIncludeGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                       `json:"total_count"`
	JSON       zeroTrustDevicePolicyIncludeGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustDevicePolicyIncludeGetResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [ZeroTrustDevicePolicyIncludeGetResponseEnvelopeResultInfo]
type zeroTrustDevicePolicyIncludeGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyIncludeGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicePolicyIncludeGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
