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

// ZeroTrustDevicePolicyFallbackDomainService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZeroTrustDevicePolicyFallbackDomainService] method instead.
type ZeroTrustDevicePolicyFallbackDomainService struct {
	Options []option.RequestOption
}

// NewZeroTrustDevicePolicyFallbackDomainService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewZeroTrustDevicePolicyFallbackDomainService(opts ...option.RequestOption) (r *ZeroTrustDevicePolicyFallbackDomainService) {
	r = &ZeroTrustDevicePolicyFallbackDomainService{}
	r.Options = opts
	return
}

// Sets the list of domains to bypass Gateway DNS resolution. These domains will
// use the specified local DNS resolver instead. This will only apply to the
// specified device settings profile.
func (r *ZeroTrustDevicePolicyFallbackDomainService) Update(ctx context.Context, policyID string, params ZeroTrustDevicePolicyFallbackDomainUpdateParams, opts ...option.RequestOption) (res *[]ZeroTrustDevicePolicyFallbackDomainUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/fallback_domains", params.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a list of domains to bypass Gateway DNS resolution. These domains will
// use the specified local DNS resolver instead.
func (r *ZeroTrustDevicePolicyFallbackDomainService) List(ctx context.Context, query ZeroTrustDevicePolicyFallbackDomainListParams, opts ...option.RequestOption) (res *[]ZeroTrustDevicePolicyFallbackDomainListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDevicePolicyFallbackDomainListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/fallback_domains", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the list of domains to bypass Gateway DNS resolution from a specified
// device settings profile. These domains will use the specified local DNS resolver
// instead.
func (r *ZeroTrustDevicePolicyFallbackDomainService) Get(ctx context.Context, policyID string, query ZeroTrustDevicePolicyFallbackDomainGetParams, opts ...option.RequestOption) (res *[]ZeroTrustDevicePolicyFallbackDomainGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDevicePolicyFallbackDomainGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/fallback_domains", query.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustDevicePolicyFallbackDomainUpdateResponse struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                                         `json:"dns_server"`
	JSON      zeroTrustDevicePolicyFallbackDomainUpdateResponseJSON `json:"-"`
}

// zeroTrustDevicePolicyFallbackDomainUpdateResponseJSON contains the JSON metadata
// for the struct [ZeroTrustDevicePolicyFallbackDomainUpdateResponse]
type zeroTrustDevicePolicyFallbackDomainUpdateResponseJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyFallbackDomainUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicePolicyFallbackDomainUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDevicePolicyFallbackDomainListResponse struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                                       `json:"dns_server"`
	JSON      zeroTrustDevicePolicyFallbackDomainListResponseJSON `json:"-"`
}

// zeroTrustDevicePolicyFallbackDomainListResponseJSON contains the JSON metadata
// for the struct [ZeroTrustDevicePolicyFallbackDomainListResponse]
type zeroTrustDevicePolicyFallbackDomainListResponseJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyFallbackDomainListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicePolicyFallbackDomainListResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDevicePolicyFallbackDomainGetResponse struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                                      `json:"dns_server"`
	JSON      zeroTrustDevicePolicyFallbackDomainGetResponseJSON `json:"-"`
}

// zeroTrustDevicePolicyFallbackDomainGetResponseJSON contains the JSON metadata
// for the struct [ZeroTrustDevicePolicyFallbackDomainGetResponse]
type zeroTrustDevicePolicyFallbackDomainGetResponseJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyFallbackDomainGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicePolicyFallbackDomainGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDevicePolicyFallbackDomainUpdateParams struct {
	AccountID param.Field[interface{}]                                           `path:"account_id,required"`
	Body      param.Field[[]ZeroTrustDevicePolicyFallbackDomainUpdateParamsBody] `json:"body,required"`
}

func (r ZeroTrustDevicePolicyFallbackDomainUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZeroTrustDevicePolicyFallbackDomainUpdateParamsBody struct {
	// The domain suffix to match when resolving locally.
	Suffix param.Field[string] `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description param.Field[string] `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer param.Field[[]interface{}] `json:"dns_server"`
}

func (r ZeroTrustDevicePolicyFallbackDomainUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelope struct {
	Errors   []ZeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustDevicePolicyFallbackDomainUpdateResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    ZeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [ZeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelope]
type zeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeErrors struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    zeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [ZeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeErrors]
type zeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeMessages struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    zeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [ZeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeMessages]
type zeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeSuccess bool

const (
	ZeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeSuccessTrue ZeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeSuccess = true
)

type ZeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                 `json:"total_count"`
	JSON       zeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeResultInfoJSON contains
// the JSON metadata for the struct
// [ZeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeResultInfo]
type zeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicePolicyFallbackDomainUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDevicePolicyFallbackDomainListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustDevicePolicyFallbackDomainListResponseEnvelope struct {
	Errors   []ZeroTrustDevicePolicyFallbackDomainListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDevicePolicyFallbackDomainListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustDevicePolicyFallbackDomainListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    ZeroTrustDevicePolicyFallbackDomainListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustDevicePolicyFallbackDomainListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustDevicePolicyFallbackDomainListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustDevicePolicyFallbackDomainListResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [ZeroTrustDevicePolicyFallbackDomainListResponseEnvelope]
type zeroTrustDevicePolicyFallbackDomainListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyFallbackDomainListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicePolicyFallbackDomainListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDevicePolicyFallbackDomainListResponseEnvelopeErrors struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    zeroTrustDevicePolicyFallbackDomainListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDevicePolicyFallbackDomainListResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [ZeroTrustDevicePolicyFallbackDomainListResponseEnvelopeErrors]
type zeroTrustDevicePolicyFallbackDomainListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyFallbackDomainListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicePolicyFallbackDomainListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDevicePolicyFallbackDomainListResponseEnvelopeMessages struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    zeroTrustDevicePolicyFallbackDomainListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDevicePolicyFallbackDomainListResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZeroTrustDevicePolicyFallbackDomainListResponseEnvelopeMessages]
type zeroTrustDevicePolicyFallbackDomainListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyFallbackDomainListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicePolicyFallbackDomainListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZeroTrustDevicePolicyFallbackDomainListResponseEnvelopeSuccess bool

const (
	ZeroTrustDevicePolicyFallbackDomainListResponseEnvelopeSuccessTrue ZeroTrustDevicePolicyFallbackDomainListResponseEnvelopeSuccess = true
)

type ZeroTrustDevicePolicyFallbackDomainListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                               `json:"total_count"`
	JSON       zeroTrustDevicePolicyFallbackDomainListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustDevicePolicyFallbackDomainListResponseEnvelopeResultInfoJSON contains
// the JSON metadata for the struct
// [ZeroTrustDevicePolicyFallbackDomainListResponseEnvelopeResultInfo]
type zeroTrustDevicePolicyFallbackDomainListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyFallbackDomainListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicePolicyFallbackDomainListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDevicePolicyFallbackDomainGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustDevicePolicyFallbackDomainGetResponseEnvelope struct {
	Errors   []ZeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustDevicePolicyFallbackDomainGetResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    ZeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustDevicePolicyFallbackDomainGetResponseEnvelope]
type zeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyFallbackDomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeErrors struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    zeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [ZeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeErrors]
type zeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeMessages struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    zeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeMessages]
type zeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeSuccess bool

const (
	ZeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeSuccessTrue ZeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeSuccess = true
)

type ZeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                              `json:"total_count"`
	JSON       zeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeResultInfoJSON contains
// the JSON metadata for the struct
// [ZeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeResultInfo]
type zeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicePolicyFallbackDomainGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
