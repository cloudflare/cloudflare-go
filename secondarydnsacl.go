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

// Modify ACL.
func (r *SecondaryDNSACLService) Update(ctx context.Context, accountIdentifier interface{}, identifier interface{}, body SecondaryDNSACLUpdateParams, opts ...option.RequestOption) (res *SecondaryDnsaclUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDnsaclUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/acls/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete ACL.
func (r *SecondaryDNSACLService) Delete(ctx context.Context, accountIdentifier interface{}, identifier interface{}, opts ...option.RequestOption) (res *SecondaryDnsaclDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDnsaclDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/acls/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get ACL.
func (r *SecondaryDNSACLService) Get(ctx context.Context, accountIdentifier interface{}, identifier interface{}, opts ...option.RequestOption) (res *SecondaryDnsaclGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDnsaclGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/acls/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Create ACL.
func (r *SecondaryDNSACLService) SecondaryDNSACLNewACL(ctx context.Context, accountIdentifier interface{}, body SecondaryDNSACLSecondaryDNSACLNewACLParams, opts ...option.RequestOption) (res *SecondaryDnsaclSecondaryDnsaclNewACLResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDnsaclSecondaryDnsaclNewACLResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/acls", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List ACLs.
func (r *SecondaryDNSACLService) SecondaryDNSACLListACLs(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *[]SecondaryDnsaclSecondaryDnsaclListACLsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDnsaclSecondaryDnsaclListACLsResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/acls", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
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

type SecondaryDnsaclSecondaryDnsaclNewACLResponse struct {
	ID interface{} `json:"id,required"`
	// Allowed IPv4/IPv6 address range of primary or secondary nameservers. This will
	// be applied for the entire account. The IP range is used to allow additional
	// NOTIFY IPs for secondary zones and IPs Cloudflare allows AXFR/IXFR requests from
	// for primary zones. CIDRs are limited to a maximum of /24 for IPv4 and /64 for
	// IPv6 respectively.
	IPRange string `json:"ip_range,required"`
	// The name of the acl.
	Name string                                           `json:"name,required"`
	JSON secondaryDnsaclSecondaryDnsaclNewACLResponseJSON `json:"-"`
}

// secondaryDnsaclSecondaryDnsaclNewACLResponseJSON contains the JSON metadata for
// the struct [SecondaryDnsaclSecondaryDnsaclNewACLResponse]
type secondaryDnsaclSecondaryDnsaclNewACLResponseJSON struct {
	ID          apijson.Field
	IPRange     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnsaclSecondaryDnsaclNewACLResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDnsaclSecondaryDnsaclListACLsResponse struct {
	ID interface{} `json:"id,required"`
	// Allowed IPv4/IPv6 address range of primary or secondary nameservers. This will
	// be applied for the entire account. The IP range is used to allow additional
	// NOTIFY IPs for secondary zones and IPs Cloudflare allows AXFR/IXFR requests from
	// for primary zones. CIDRs are limited to a maximum of /24 for IPv4 and /64 for
	// IPv6 respectively.
	IPRange string `json:"ip_range,required"`
	// The name of the acl.
	Name string                                             `json:"name,required"`
	JSON secondaryDnsaclSecondaryDnsaclListACLsResponseJSON `json:"-"`
}

// secondaryDnsaclSecondaryDnsaclListACLsResponseJSON contains the JSON metadata
// for the struct [SecondaryDnsaclSecondaryDnsaclListACLsResponse]
type secondaryDnsaclSecondaryDnsaclListACLsResponseJSON struct {
	ID          apijson.Field
	IPRange     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnsaclSecondaryDnsaclListACLsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSACLUpdateParams struct {
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

// Whether the API call was successful
type SecondaryDnsaclUpdateResponseEnvelopeSuccess bool

const (
	SecondaryDnsaclUpdateResponseEnvelopeSuccessTrue SecondaryDnsaclUpdateResponseEnvelopeSuccess = true
)

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

// Whether the API call was successful
type SecondaryDnsaclDeleteResponseEnvelopeSuccess bool

const (
	SecondaryDnsaclDeleteResponseEnvelopeSuccessTrue SecondaryDnsaclDeleteResponseEnvelopeSuccess = true
)

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

// Whether the API call was successful
type SecondaryDnsaclGetResponseEnvelopeSuccess bool

const (
	SecondaryDnsaclGetResponseEnvelopeSuccessTrue SecondaryDnsaclGetResponseEnvelopeSuccess = true
)

type SecondaryDNSACLSecondaryDNSACLNewACLParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r SecondaryDNSACLSecondaryDNSACLNewACLParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type SecondaryDnsaclSecondaryDnsaclNewACLResponseEnvelope struct {
	Errors   []SecondaryDnsaclSecondaryDnsaclNewACLResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDnsaclSecondaryDnsaclNewACLResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDnsaclSecondaryDnsaclNewACLResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDnsaclSecondaryDnsaclNewACLResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDnsaclSecondaryDnsaclNewACLResponseEnvelopeJSON    `json:"-"`
}

// secondaryDnsaclSecondaryDnsaclNewACLResponseEnvelopeJSON contains the JSON
// metadata for the struct [SecondaryDnsaclSecondaryDnsaclNewACLResponseEnvelope]
type secondaryDnsaclSecondaryDnsaclNewACLResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnsaclSecondaryDnsaclNewACLResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDnsaclSecondaryDnsaclNewACLResponseEnvelopeErrors struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    secondaryDnsaclSecondaryDnsaclNewACLResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDnsaclSecondaryDnsaclNewACLResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [SecondaryDnsaclSecondaryDnsaclNewACLResponseEnvelopeErrors]
type secondaryDnsaclSecondaryDnsaclNewACLResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnsaclSecondaryDnsaclNewACLResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDnsaclSecondaryDnsaclNewACLResponseEnvelopeMessages struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    secondaryDnsaclSecondaryDnsaclNewACLResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDnsaclSecondaryDnsaclNewACLResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [SecondaryDnsaclSecondaryDnsaclNewACLResponseEnvelopeMessages]
type secondaryDnsaclSecondaryDnsaclNewACLResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnsaclSecondaryDnsaclNewACLResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDnsaclSecondaryDnsaclNewACLResponseEnvelopeSuccess bool

const (
	SecondaryDnsaclSecondaryDnsaclNewACLResponseEnvelopeSuccessTrue SecondaryDnsaclSecondaryDnsaclNewACLResponseEnvelopeSuccess = true
)

type SecondaryDnsaclSecondaryDnsaclListACLsResponseEnvelope struct {
	Errors   []SecondaryDnsaclSecondaryDnsaclListACLsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDnsaclSecondaryDnsaclListACLsResponseEnvelopeMessages `json:"messages,required"`
	Result   []SecondaryDnsaclSecondaryDnsaclListACLsResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    SecondaryDnsaclSecondaryDnsaclListACLsResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo SecondaryDnsaclSecondaryDnsaclListACLsResponseEnvelopeResultInfo `json:"result_info"`
	JSON       secondaryDnsaclSecondaryDnsaclListACLsResponseEnvelopeJSON       `json:"-"`
}

// secondaryDnsaclSecondaryDnsaclListACLsResponseEnvelopeJSON contains the JSON
// metadata for the struct [SecondaryDnsaclSecondaryDnsaclListACLsResponseEnvelope]
type secondaryDnsaclSecondaryDnsaclListACLsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnsaclSecondaryDnsaclListACLsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDnsaclSecondaryDnsaclListACLsResponseEnvelopeErrors struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    secondaryDnsaclSecondaryDnsaclListACLsResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDnsaclSecondaryDnsaclListACLsResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [SecondaryDnsaclSecondaryDnsaclListACLsResponseEnvelopeErrors]
type secondaryDnsaclSecondaryDnsaclListACLsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnsaclSecondaryDnsaclListACLsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDnsaclSecondaryDnsaclListACLsResponseEnvelopeMessages struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    secondaryDnsaclSecondaryDnsaclListACLsResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDnsaclSecondaryDnsaclListACLsResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [SecondaryDnsaclSecondaryDnsaclListACLsResponseEnvelopeMessages]
type secondaryDnsaclSecondaryDnsaclListACLsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnsaclSecondaryDnsaclListACLsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDnsaclSecondaryDnsaclListACLsResponseEnvelopeSuccess bool

const (
	SecondaryDnsaclSecondaryDnsaclListACLsResponseEnvelopeSuccessTrue SecondaryDnsaclSecondaryDnsaclListACLsResponseEnvelopeSuccess = true
)

type SecondaryDnsaclSecondaryDnsaclListACLsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                              `json:"total_count"`
	JSON       secondaryDnsaclSecondaryDnsaclListACLsResponseEnvelopeResultInfoJSON `json:"-"`
}

// secondaryDnsaclSecondaryDnsaclListACLsResponseEnvelopeResultInfoJSON contains
// the JSON metadata for the struct
// [SecondaryDnsaclSecondaryDnsaclListACLsResponseEnvelopeResultInfo]
type secondaryDnsaclSecondaryDnsaclListACLsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnsaclSecondaryDnsaclListACLsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
