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

// ZeroTrustDevicePolicyExcludeService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZeroTrustDevicePolicyExcludeService] method instead.
type ZeroTrustDevicePolicyExcludeService struct {
	Options []option.RequestOption
}

// NewZeroTrustDevicePolicyExcludeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustDevicePolicyExcludeService(opts ...option.RequestOption) (r *ZeroTrustDevicePolicyExcludeService) {
	r = &ZeroTrustDevicePolicyExcludeService{}
	r.Options = opts
	return
}

// Sets the list of routes excluded from the WARP client's tunnel.
func (r *ZeroTrustDevicePolicyExcludeService) Update(ctx context.Context, params ZeroTrustDevicePolicyExcludeUpdateParams, opts ...option.RequestOption) (res *[]ZeroTrustDevicePolicyExcludeUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDevicePolicyExcludeUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/exclude", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the list of routes excluded from the WARP client's tunnel.
func (r *ZeroTrustDevicePolicyExcludeService) List(ctx context.Context, query ZeroTrustDevicePolicyExcludeListParams, opts ...option.RequestOption) (res *[]ZeroTrustDevicePolicyExcludeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDevicePolicyExcludeListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/exclude", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the list of routes excluded from the WARP client's tunnel for a specific
// device settings profile.
func (r *ZeroTrustDevicePolicyExcludeService) Get(ctx context.Context, policyID string, query ZeroTrustDevicePolicyExcludeGetParams, opts ...option.RequestOption) (res *[]ZeroTrustDevicePolicyExcludeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDevicePolicyExcludeGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/exclude", query.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustDevicePolicyExcludeUpdateResponse struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host string                                         `json:"host"`
	JSON zeroTrustDevicePolicyExcludeUpdateResponseJSON `json:"-"`
}

// zeroTrustDevicePolicyExcludeUpdateResponseJSON contains the JSON metadata for
// the struct [ZeroTrustDevicePolicyExcludeUpdateResponse]
type zeroTrustDevicePolicyExcludeUpdateResponseJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyExcludeUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyExcludeListResponse struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host string                                       `json:"host"`
	JSON zeroTrustDevicePolicyExcludeListResponseJSON `json:"-"`
}

// zeroTrustDevicePolicyExcludeListResponseJSON contains the JSON metadata for the
// struct [ZeroTrustDevicePolicyExcludeListResponse]
type zeroTrustDevicePolicyExcludeListResponseJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyExcludeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyExcludeGetResponse struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host string                                      `json:"host"`
	JSON zeroTrustDevicePolicyExcludeGetResponseJSON `json:"-"`
}

// zeroTrustDevicePolicyExcludeGetResponseJSON contains the JSON metadata for the
// struct [ZeroTrustDevicePolicyExcludeGetResponse]
type zeroTrustDevicePolicyExcludeGetResponseJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyExcludeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyExcludeUpdateParams struct {
	AccountID param.Field[interface{}]                                    `path:"account_id,required"`
	Body      param.Field[[]ZeroTrustDevicePolicyExcludeUpdateParamsBody] `json:"body,required"`
}

func (r ZeroTrustDevicePolicyExcludeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZeroTrustDevicePolicyExcludeUpdateParamsBody struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address param.Field[string] `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host param.Field[string] `json:"host"`
}

func (r ZeroTrustDevicePolicyExcludeUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustDevicePolicyExcludeUpdateResponseEnvelope struct {
	Errors   []ZeroTrustDevicePolicyExcludeUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDevicePolicyExcludeUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustDevicePolicyExcludeUpdateResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    ZeroTrustDevicePolicyExcludeUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustDevicePolicyExcludeUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustDevicePolicyExcludeUpdateResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustDevicePolicyExcludeUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustDevicePolicyExcludeUpdateResponseEnvelope]
type zeroTrustDevicePolicyExcludeUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyExcludeUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyExcludeUpdateResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zeroTrustDevicePolicyExcludeUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDevicePolicyExcludeUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustDevicePolicyExcludeUpdateResponseEnvelopeErrors]
type zeroTrustDevicePolicyExcludeUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyExcludeUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyExcludeUpdateResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    zeroTrustDevicePolicyExcludeUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDevicePolicyExcludeUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustDevicePolicyExcludeUpdateResponseEnvelopeMessages]
type zeroTrustDevicePolicyExcludeUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyExcludeUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type ZeroTrustDevicePolicyExcludeUpdateResponseEnvelopeSuccess bool

const (
	ZeroTrustDevicePolicyExcludeUpdateResponseEnvelopeSuccessTrue ZeroTrustDevicePolicyExcludeUpdateResponseEnvelopeSuccess = true
)

type ZeroTrustDevicePolicyExcludeUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                          `json:"total_count"`
	JSON       zeroTrustDevicePolicyExcludeUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustDevicePolicyExcludeUpdateResponseEnvelopeResultInfoJSON contains the
// JSON metadata for the struct
// [ZeroTrustDevicePolicyExcludeUpdateResponseEnvelopeResultInfo]
type zeroTrustDevicePolicyExcludeUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyExcludeUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyExcludeListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustDevicePolicyExcludeListResponseEnvelope struct {
	Errors   []ZeroTrustDevicePolicyExcludeListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDevicePolicyExcludeListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustDevicePolicyExcludeListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    ZeroTrustDevicePolicyExcludeListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustDevicePolicyExcludeListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustDevicePolicyExcludeListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustDevicePolicyExcludeListResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustDevicePolicyExcludeListResponseEnvelope]
type zeroTrustDevicePolicyExcludeListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyExcludeListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyExcludeListResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zeroTrustDevicePolicyExcludeListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDevicePolicyExcludeListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustDevicePolicyExcludeListResponseEnvelopeErrors]
type zeroTrustDevicePolicyExcludeListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyExcludeListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyExcludeListResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zeroTrustDevicePolicyExcludeListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDevicePolicyExcludeListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustDevicePolicyExcludeListResponseEnvelopeMessages]
type zeroTrustDevicePolicyExcludeListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyExcludeListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type ZeroTrustDevicePolicyExcludeListResponseEnvelopeSuccess bool

const (
	ZeroTrustDevicePolicyExcludeListResponseEnvelopeSuccessTrue ZeroTrustDevicePolicyExcludeListResponseEnvelopeSuccess = true
)

type ZeroTrustDevicePolicyExcludeListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                        `json:"total_count"`
	JSON       zeroTrustDevicePolicyExcludeListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustDevicePolicyExcludeListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [ZeroTrustDevicePolicyExcludeListResponseEnvelopeResultInfo]
type zeroTrustDevicePolicyExcludeListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyExcludeListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyExcludeGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustDevicePolicyExcludeGetResponseEnvelope struct {
	Errors   []ZeroTrustDevicePolicyExcludeGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDevicePolicyExcludeGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustDevicePolicyExcludeGetResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    ZeroTrustDevicePolicyExcludeGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustDevicePolicyExcludeGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustDevicePolicyExcludeGetResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustDevicePolicyExcludeGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustDevicePolicyExcludeGetResponseEnvelope]
type zeroTrustDevicePolicyExcludeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyExcludeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyExcludeGetResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zeroTrustDevicePolicyExcludeGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDevicePolicyExcludeGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustDevicePolicyExcludeGetResponseEnvelopeErrors]
type zeroTrustDevicePolicyExcludeGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyExcludeGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyExcludeGetResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zeroTrustDevicePolicyExcludeGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDevicePolicyExcludeGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustDevicePolicyExcludeGetResponseEnvelopeMessages]
type zeroTrustDevicePolicyExcludeGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyExcludeGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type ZeroTrustDevicePolicyExcludeGetResponseEnvelopeSuccess bool

const (
	ZeroTrustDevicePolicyExcludeGetResponseEnvelopeSuccessTrue ZeroTrustDevicePolicyExcludeGetResponseEnvelopeSuccess = true
)

type ZeroTrustDevicePolicyExcludeGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                       `json:"total_count"`
	JSON       zeroTrustDevicePolicyExcludeGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustDevicePolicyExcludeGetResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [ZeroTrustDevicePolicyExcludeGetResponseEnvelopeResultInfo]
type zeroTrustDevicePolicyExcludeGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyExcludeGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
