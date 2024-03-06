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

// SecondaryDNSACLService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSecondaryDNSACLService] method
// instead.
type SecondaryDNSACLService struct {
	Options []option.RequestOption
}

// NewSecondaryDNSACLService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSecondaryDNSACLService(opts ...option.RequestOption) (r *SecondaryDNSACLService) {
	r = &SecondaryDNSACLService{}
	r.Options = opts
	return
}

// Create ACL.
func (r *SecondaryDNSACLService) New(ctx context.Context, params SecondaryDNSACLNewParams, opts ...option.RequestOption) (res *SecondaryDnsaclNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDnsaclNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/acls", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify ACL.
func (r *SecondaryDNSACLService) Update(ctx context.Context, aclID interface{}, params SecondaryDNSACLUpdateParams, opts ...option.RequestOption) (res *SecondaryDnsaclUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDnsaclUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/acls/%v", params.AccountID, aclID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List ACLs.
func (r *SecondaryDNSACLService) List(ctx context.Context, query SecondaryDNSACLListParams, opts ...option.RequestOption) (res *[]SecondaryDnsaclListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDnsaclListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/acls", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete ACL.
func (r *SecondaryDNSACLService) Delete(ctx context.Context, aclID interface{}, body SecondaryDNSACLDeleteParams, opts ...option.RequestOption) (res *SecondaryDnsaclDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDnsaclDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/acls/%v", body.AccountID, aclID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get ACL.
func (r *SecondaryDNSACLService) Get(ctx context.Context, aclID interface{}, query SecondaryDNSACLGetParams, opts ...option.RequestOption) (res *SecondaryDnsaclGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDnsaclGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/acls/%v", query.AccountID, aclID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SecondaryDnsaclNewResponse struct {
	ID interface{} `json:"id,required"`
	// Allowed IPv4/IPv6 address range of primary or secondary nameservers. This will
	// be applied for the entire account. The IP range is used to allow additional
	// NOTIFY IPs for secondary zones and IPs Cloudflare allows AXFR/IXFR requests from
	// for primary zones. CIDRs are limited to a maximum of /24 for IPv4 and /64 for
	// IPv6 respectively.
	IPRange string `json:"ip_range,required"`
	// The name of the acl.
	Name string                         `json:"name,required"`
	JSON secondaryDnsaclNewResponseJSON `json:"-"`
}

// secondaryDnsaclNewResponseJSON contains the JSON metadata for the struct
// [SecondaryDnsaclNewResponse]
type secondaryDnsaclNewResponseJSON struct {
	ID          apijson.Field
	IPRange     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnsaclNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secondaryDnsaclNewResponseJSON) RawJSON() string {
	return r.raw
}

type SecondaryDnsaclUpdateResponse struct {
	ID interface{} `json:"id,required"`
	// Allowed IPv4/IPv6 address range of primary or secondary nameservers. This will
	// be applied for the entire account. The IP range is used to allow additional
	// NOTIFY IPs for secondary zones and IPs Cloudflare allows AXFR/IXFR requests from
	// for primary zones. CIDRs are limited to a maximum of /24 for IPv4 and /64 for
	// IPv6 respectively.
	IPRange string `json:"ip_range,required"`
	// The name of the acl.
	Name string                            `json:"name,required"`
	JSON secondaryDnsaclUpdateResponseJSON `json:"-"`
}

// secondaryDnsaclUpdateResponseJSON contains the JSON metadata for the struct
// [SecondaryDnsaclUpdateResponse]
type secondaryDnsaclUpdateResponseJSON struct {
	ID          apijson.Field
	IPRange     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnsaclUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secondaryDnsaclUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type SecondaryDnsaclListResponse struct {
	ID interface{} `json:"id,required"`
	// Allowed IPv4/IPv6 address range of primary or secondary nameservers. This will
	// be applied for the entire account. The IP range is used to allow additional
	// NOTIFY IPs for secondary zones and IPs Cloudflare allows AXFR/IXFR requests from
	// for primary zones. CIDRs are limited to a maximum of /24 for IPv4 and /64 for
	// IPv6 respectively.
	IPRange string `json:"ip_range,required"`
	// The name of the acl.
	Name string                          `json:"name,required"`
	JSON secondaryDnsaclListResponseJSON `json:"-"`
}

// secondaryDnsaclListResponseJSON contains the JSON metadata for the struct
// [SecondaryDnsaclListResponse]
type secondaryDnsaclListResponseJSON struct {
	ID          apijson.Field
	IPRange     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnsaclListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secondaryDnsaclListResponseJSON) RawJSON() string {
	return r.raw
}

type SecondaryDnsaclDeleteResponse struct {
	ID   interface{}                       `json:"id"`
	JSON secondaryDnsaclDeleteResponseJSON `json:"-"`
}

// secondaryDnsaclDeleteResponseJSON contains the JSON metadata for the struct
// [SecondaryDnsaclDeleteResponse]
type secondaryDnsaclDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnsaclDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secondaryDnsaclDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SecondaryDnsaclGetResponse struct {
	ID interface{} `json:"id,required"`
	// Allowed IPv4/IPv6 address range of primary or secondary nameservers. This will
	// be applied for the entire account. The IP range is used to allow additional
	// NOTIFY IPs for secondary zones and IPs Cloudflare allows AXFR/IXFR requests from
	// for primary zones. CIDRs are limited to a maximum of /24 for IPv4 and /64 for
	// IPv6 respectively.
	IPRange string `json:"ip_range,required"`
	// The name of the acl.
	Name string                         `json:"name,required"`
	JSON secondaryDnsaclGetResponseJSON `json:"-"`
}

// secondaryDnsaclGetResponseJSON contains the JSON metadata for the struct
// [SecondaryDnsaclGetResponse]
type secondaryDnsaclGetResponseJSON struct {
	ID          apijson.Field
	IPRange     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnsaclGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secondaryDnsaclGetResponseJSON) RawJSON() string {
	return r.raw
}

type SecondaryDNSACLNewParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r SecondaryDNSACLNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type SecondaryDnsaclNewResponseEnvelope struct {
	Errors   []SecondaryDnsaclNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDnsaclNewResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDnsaclNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDnsaclNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDnsaclNewResponseEnvelopeJSON    `json:"-"`
}

// secondaryDnsaclNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SecondaryDnsaclNewResponseEnvelope]
type secondaryDnsaclNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnsaclNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secondaryDnsaclNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SecondaryDnsaclNewResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    secondaryDnsaclNewResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDnsaclNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SecondaryDnsaclNewResponseEnvelopeErrors]
type secondaryDnsaclNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnsaclNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secondaryDnsaclNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SecondaryDnsaclNewResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    secondaryDnsaclNewResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDnsaclNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SecondaryDnsaclNewResponseEnvelopeMessages]
type secondaryDnsaclNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnsaclNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secondaryDnsaclNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SecondaryDnsaclNewResponseEnvelopeSuccess bool

const (
	SecondaryDnsaclNewResponseEnvelopeSuccessTrue SecondaryDnsaclNewResponseEnvelopeSuccess = true
)

type SecondaryDNSACLUpdateParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// Allowed IPv4/IPv6 address range of primary or secondary nameservers. This will
	// be applied for the entire account. The IP range is used to allow additional
	// NOTIFY IPs for secondary zones and IPs Cloudflare allows AXFR/IXFR requests from
	// for primary zones. CIDRs are limited to a maximum of /24 for IPv4 and /64 for
	// IPv6 respectively.
	IPRange param.Field[string] `json:"ip_range,required"`
	// The name of the acl.
	Name param.Field[string] `json:"name,required"`
}

func (r SecondaryDNSACLUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SecondaryDnsaclUpdateResponseEnvelope struct {
	Errors   []SecondaryDnsaclUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDnsaclUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDnsaclUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDnsaclUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDnsaclUpdateResponseEnvelopeJSON    `json:"-"`
}

// secondaryDnsaclUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [SecondaryDnsaclUpdateResponseEnvelope]
type secondaryDnsaclUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnsaclUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secondaryDnsaclUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SecondaryDnsaclUpdateResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    secondaryDnsaclUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDnsaclUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SecondaryDnsaclUpdateResponseEnvelopeErrors]
type secondaryDnsaclUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnsaclUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secondaryDnsaclUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SecondaryDnsaclUpdateResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    secondaryDnsaclUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDnsaclUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SecondaryDnsaclUpdateResponseEnvelopeMessages]
type secondaryDnsaclUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnsaclUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secondaryDnsaclUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SecondaryDnsaclUpdateResponseEnvelopeSuccess bool

const (
	SecondaryDnsaclUpdateResponseEnvelopeSuccessTrue SecondaryDnsaclUpdateResponseEnvelopeSuccess = true
)

type SecondaryDNSACLListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type SecondaryDnsaclListResponseEnvelope struct {
	Errors   []SecondaryDnsaclListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDnsaclListResponseEnvelopeMessages `json:"messages,required"`
	Result   []SecondaryDnsaclListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    SecondaryDnsaclListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo SecondaryDnsaclListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       secondaryDnsaclListResponseEnvelopeJSON       `json:"-"`
}

// secondaryDnsaclListResponseEnvelopeJSON contains the JSON metadata for the
// struct [SecondaryDnsaclListResponseEnvelope]
type secondaryDnsaclListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnsaclListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secondaryDnsaclListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SecondaryDnsaclListResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    secondaryDnsaclListResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDnsaclListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SecondaryDnsaclListResponseEnvelopeErrors]
type secondaryDnsaclListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnsaclListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secondaryDnsaclListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SecondaryDnsaclListResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    secondaryDnsaclListResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDnsaclListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SecondaryDnsaclListResponseEnvelopeMessages]
type secondaryDnsaclListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnsaclListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secondaryDnsaclListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SecondaryDnsaclListResponseEnvelopeSuccess bool

const (
	SecondaryDnsaclListResponseEnvelopeSuccessTrue SecondaryDnsaclListResponseEnvelopeSuccess = true
)

type SecondaryDnsaclListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                           `json:"total_count"`
	JSON       secondaryDnsaclListResponseEnvelopeResultInfoJSON `json:"-"`
}

// secondaryDnsaclListResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [SecondaryDnsaclListResponseEnvelopeResultInfo]
type secondaryDnsaclListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnsaclListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secondaryDnsaclListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type SecondaryDNSACLDeleteParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type SecondaryDnsaclDeleteResponseEnvelope struct {
	Errors   []SecondaryDnsaclDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDnsaclDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDnsaclDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDnsaclDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDnsaclDeleteResponseEnvelopeJSON    `json:"-"`
}

// secondaryDnsaclDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [SecondaryDnsaclDeleteResponseEnvelope]
type secondaryDnsaclDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnsaclDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secondaryDnsaclDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SecondaryDnsaclDeleteResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    secondaryDnsaclDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDnsaclDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SecondaryDnsaclDeleteResponseEnvelopeErrors]
type secondaryDnsaclDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnsaclDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secondaryDnsaclDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SecondaryDnsaclDeleteResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    secondaryDnsaclDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDnsaclDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SecondaryDnsaclDeleteResponseEnvelopeMessages]
type secondaryDnsaclDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnsaclDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secondaryDnsaclDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SecondaryDnsaclDeleteResponseEnvelopeSuccess bool

const (
	SecondaryDnsaclDeleteResponseEnvelopeSuccessTrue SecondaryDnsaclDeleteResponseEnvelopeSuccess = true
)

type SecondaryDNSACLGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type SecondaryDnsaclGetResponseEnvelope struct {
	Errors   []SecondaryDnsaclGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDnsaclGetResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDnsaclGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDnsaclGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDnsaclGetResponseEnvelopeJSON    `json:"-"`
}

// secondaryDnsaclGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SecondaryDnsaclGetResponseEnvelope]
type secondaryDnsaclGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnsaclGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secondaryDnsaclGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SecondaryDnsaclGetResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    secondaryDnsaclGetResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDnsaclGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SecondaryDnsaclGetResponseEnvelopeErrors]
type secondaryDnsaclGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnsaclGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secondaryDnsaclGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SecondaryDnsaclGetResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    secondaryDnsaclGetResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDnsaclGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SecondaryDnsaclGetResponseEnvelopeMessages]
type secondaryDnsaclGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnsaclGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secondaryDnsaclGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SecondaryDnsaclGetResponseEnvelopeSuccess bool

const (
	SecondaryDnsaclGetResponseEnvelopeSuccessTrue SecondaryDnsaclGetResponseEnvelopeSuccess = true
)
