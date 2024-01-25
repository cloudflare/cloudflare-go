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

// AccountSecondaryDNSACLService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountSecondaryDNSACLService]
// method instead.
type AccountSecondaryDNSACLService struct {
	Options []option.RequestOption
}

// NewAccountSecondaryDNSACLService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountSecondaryDNSACLService(opts ...option.RequestOption) (r *AccountSecondaryDNSACLService) {
	r = &AccountSecondaryDNSACLService{}
	r.Options = opts
	return
}

// Get ACL.
func (r *AccountSecondaryDNSACLService) Get(ctx context.Context, accountIdentifier interface{}, identifier interface{}, opts ...option.RequestOption) (res *AccountSecondaryDnsaclGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/secondary_dns/acls/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify ACL.
func (r *AccountSecondaryDNSACLService) Update(ctx context.Context, accountIdentifier interface{}, identifier interface{}, body AccountSecondaryDNSACLUpdateParams, opts ...option.RequestOption) (res *AccountSecondaryDnsaclUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/secondary_dns/acls/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete ACL.
func (r *AccountSecondaryDNSACLService) Delete(ctx context.Context, accountIdentifier interface{}, identifier interface{}, opts ...option.RequestOption) (res *AccountSecondaryDnsaclDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/secondary_dns/acls/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create ACL.
func (r *AccountSecondaryDNSACLService) SecondaryDNSACLNewACL(ctx context.Context, accountIdentifier interface{}, body AccountSecondaryDNSACLSecondaryDNSACLNewACLParams, opts ...option.RequestOption) (res *AccountSecondaryDnsaclSecondaryDnsaclNewACLResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/secondary_dns/acls", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List ACLs.
func (r *AccountSecondaryDNSACLService) SecondaryDNSACLListACLs(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *AccountSecondaryDnsaclSecondaryDnsaclListACLsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/secondary_dns/acls", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountSecondaryDnsaclGetResponse struct {
	Errors   []AccountSecondaryDnsaclGetResponseError   `json:"errors"`
	Messages []AccountSecondaryDnsaclGetResponseMessage `json:"messages"`
	Result   AccountSecondaryDnsaclGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountSecondaryDnsaclGetResponseSuccess `json:"success"`
	JSON    accountSecondaryDnsaclGetResponseJSON    `json:"-"`
}

// accountSecondaryDnsaclGetResponseJSON contains the JSON metadata for the struct
// [AccountSecondaryDnsaclGetResponse]
type accountSecondaryDnsaclGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDnsaclGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDnsaclGetResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountSecondaryDnsaclGetResponseErrorJSON `json:"-"`
}

// accountSecondaryDnsaclGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountSecondaryDnsaclGetResponseError]
type accountSecondaryDnsaclGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDnsaclGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDnsaclGetResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountSecondaryDnsaclGetResponseMessageJSON `json:"-"`
}

// accountSecondaryDnsaclGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountSecondaryDnsaclGetResponseMessage]
type accountSecondaryDnsaclGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDnsaclGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDnsaclGetResponseResult struct {
	ID interface{} `json:"id,required"`
	// Allowed IPv4/IPv6 address range of primary or secondary nameservers. This will
	// be applied for the entire account. The IP range is used to allow additional
	// NOTIFY IPs for secondary zones and IPs Cloudflare allows AXFR/IXFR requests from
	// for primary zones. CIDRs are limited to a maximum of /24 for IPv4 and /64 for
	// IPv6 respectively.
	IPRange string `json:"ip_range,required"`
	// The name of the acl.
	Name string                                      `json:"name,required"`
	JSON accountSecondaryDnsaclGetResponseResultJSON `json:"-"`
}

// accountSecondaryDnsaclGetResponseResultJSON contains the JSON metadata for the
// struct [AccountSecondaryDnsaclGetResponseResult]
type accountSecondaryDnsaclGetResponseResultJSON struct {
	ID          apijson.Field
	IPRange     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDnsaclGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountSecondaryDnsaclGetResponseSuccess bool

const (
	AccountSecondaryDnsaclGetResponseSuccessTrue AccountSecondaryDnsaclGetResponseSuccess = true
)

type AccountSecondaryDnsaclUpdateResponse struct {
	Errors   []AccountSecondaryDnsaclUpdateResponseError   `json:"errors"`
	Messages []AccountSecondaryDnsaclUpdateResponseMessage `json:"messages"`
	Result   AccountSecondaryDnsaclUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountSecondaryDnsaclUpdateResponseSuccess `json:"success"`
	JSON    accountSecondaryDnsaclUpdateResponseJSON    `json:"-"`
}

// accountSecondaryDnsaclUpdateResponseJSON contains the JSON metadata for the
// struct [AccountSecondaryDnsaclUpdateResponse]
type accountSecondaryDnsaclUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDnsaclUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDnsaclUpdateResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountSecondaryDnsaclUpdateResponseErrorJSON `json:"-"`
}

// accountSecondaryDnsaclUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccountSecondaryDnsaclUpdateResponseError]
type accountSecondaryDnsaclUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDnsaclUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDnsaclUpdateResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accountSecondaryDnsaclUpdateResponseMessageJSON `json:"-"`
}

// accountSecondaryDnsaclUpdateResponseMessageJSON contains the JSON metadata for
// the struct [AccountSecondaryDnsaclUpdateResponseMessage]
type accountSecondaryDnsaclUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDnsaclUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDnsaclUpdateResponseResult struct {
	ID interface{} `json:"id,required"`
	// Allowed IPv4/IPv6 address range of primary or secondary nameservers. This will
	// be applied for the entire account. The IP range is used to allow additional
	// NOTIFY IPs for secondary zones and IPs Cloudflare allows AXFR/IXFR requests from
	// for primary zones. CIDRs are limited to a maximum of /24 for IPv4 and /64 for
	// IPv6 respectively.
	IPRange string `json:"ip_range,required"`
	// The name of the acl.
	Name string                                         `json:"name,required"`
	JSON accountSecondaryDnsaclUpdateResponseResultJSON `json:"-"`
}

// accountSecondaryDnsaclUpdateResponseResultJSON contains the JSON metadata for
// the struct [AccountSecondaryDnsaclUpdateResponseResult]
type accountSecondaryDnsaclUpdateResponseResultJSON struct {
	ID          apijson.Field
	IPRange     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDnsaclUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountSecondaryDnsaclUpdateResponseSuccess bool

const (
	AccountSecondaryDnsaclUpdateResponseSuccessTrue AccountSecondaryDnsaclUpdateResponseSuccess = true
)

type AccountSecondaryDnsaclDeleteResponse struct {
	Errors   []AccountSecondaryDnsaclDeleteResponseError   `json:"errors"`
	Messages []AccountSecondaryDnsaclDeleteResponseMessage `json:"messages"`
	Result   AccountSecondaryDnsaclDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountSecondaryDnsaclDeleteResponseSuccess `json:"success"`
	JSON    accountSecondaryDnsaclDeleteResponseJSON    `json:"-"`
}

// accountSecondaryDnsaclDeleteResponseJSON contains the JSON metadata for the
// struct [AccountSecondaryDnsaclDeleteResponse]
type accountSecondaryDnsaclDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDnsaclDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDnsaclDeleteResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountSecondaryDnsaclDeleteResponseErrorJSON `json:"-"`
}

// accountSecondaryDnsaclDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountSecondaryDnsaclDeleteResponseError]
type accountSecondaryDnsaclDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDnsaclDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDnsaclDeleteResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accountSecondaryDnsaclDeleteResponseMessageJSON `json:"-"`
}

// accountSecondaryDnsaclDeleteResponseMessageJSON contains the JSON metadata for
// the struct [AccountSecondaryDnsaclDeleteResponseMessage]
type accountSecondaryDnsaclDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDnsaclDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDnsaclDeleteResponseResult struct {
	ID   interface{}                                    `json:"id"`
	JSON accountSecondaryDnsaclDeleteResponseResultJSON `json:"-"`
}

// accountSecondaryDnsaclDeleteResponseResultJSON contains the JSON metadata for
// the struct [AccountSecondaryDnsaclDeleteResponseResult]
type accountSecondaryDnsaclDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDnsaclDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountSecondaryDnsaclDeleteResponseSuccess bool

const (
	AccountSecondaryDnsaclDeleteResponseSuccessTrue AccountSecondaryDnsaclDeleteResponseSuccess = true
)

type AccountSecondaryDnsaclSecondaryDnsaclNewACLResponse struct {
	Errors   []AccountSecondaryDnsaclSecondaryDnsaclNewACLResponseError   `json:"errors"`
	Messages []AccountSecondaryDnsaclSecondaryDnsaclNewACLResponseMessage `json:"messages"`
	Result   AccountSecondaryDnsaclSecondaryDnsaclNewACLResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountSecondaryDnsaclSecondaryDnsaclNewACLResponseSuccess `json:"success"`
	JSON    accountSecondaryDnsaclSecondaryDnsaclNewACLResponseJSON    `json:"-"`
}

// accountSecondaryDnsaclSecondaryDnsaclNewACLResponseJSON contains the JSON
// metadata for the struct [AccountSecondaryDnsaclSecondaryDnsaclNewACLResponse]
type accountSecondaryDnsaclSecondaryDnsaclNewACLResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDnsaclSecondaryDnsaclNewACLResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDnsaclSecondaryDnsaclNewACLResponseError struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    accountSecondaryDnsaclSecondaryDnsaclNewACLResponseErrorJSON `json:"-"`
}

// accountSecondaryDnsaclSecondaryDnsaclNewACLResponseErrorJSON contains the JSON
// metadata for the struct
// [AccountSecondaryDnsaclSecondaryDnsaclNewACLResponseError]
type accountSecondaryDnsaclSecondaryDnsaclNewACLResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDnsaclSecondaryDnsaclNewACLResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDnsaclSecondaryDnsaclNewACLResponseMessage struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    accountSecondaryDnsaclSecondaryDnsaclNewACLResponseMessageJSON `json:"-"`
}

// accountSecondaryDnsaclSecondaryDnsaclNewACLResponseMessageJSON contains the JSON
// metadata for the struct
// [AccountSecondaryDnsaclSecondaryDnsaclNewACLResponseMessage]
type accountSecondaryDnsaclSecondaryDnsaclNewACLResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDnsaclSecondaryDnsaclNewACLResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDnsaclSecondaryDnsaclNewACLResponseResult struct {
	ID interface{} `json:"id,required"`
	// Allowed IPv4/IPv6 address range of primary or secondary nameservers. This will
	// be applied for the entire account. The IP range is used to allow additional
	// NOTIFY IPs for secondary zones and IPs Cloudflare allows AXFR/IXFR requests from
	// for primary zones. CIDRs are limited to a maximum of /24 for IPv4 and /64 for
	// IPv6 respectively.
	IPRange string `json:"ip_range,required"`
	// The name of the acl.
	Name string                                                        `json:"name,required"`
	JSON accountSecondaryDnsaclSecondaryDnsaclNewACLResponseResultJSON `json:"-"`
}

// accountSecondaryDnsaclSecondaryDnsaclNewACLResponseResultJSON contains the JSON
// metadata for the struct
// [AccountSecondaryDnsaclSecondaryDnsaclNewACLResponseResult]
type accountSecondaryDnsaclSecondaryDnsaclNewACLResponseResultJSON struct {
	ID          apijson.Field
	IPRange     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDnsaclSecondaryDnsaclNewACLResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountSecondaryDnsaclSecondaryDnsaclNewACLResponseSuccess bool

const (
	AccountSecondaryDnsaclSecondaryDnsaclNewACLResponseSuccessTrue AccountSecondaryDnsaclSecondaryDnsaclNewACLResponseSuccess = true
)

type AccountSecondaryDnsaclSecondaryDnsaclListACLsResponse struct {
	Errors     []AccountSecondaryDnsaclSecondaryDnsaclListACLsResponseError    `json:"errors"`
	Messages   []AccountSecondaryDnsaclSecondaryDnsaclListACLsResponseMessage  `json:"messages"`
	Result     []AccountSecondaryDnsaclSecondaryDnsaclListACLsResponseResult   `json:"result"`
	ResultInfo AccountSecondaryDnsaclSecondaryDnsaclListACLsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountSecondaryDnsaclSecondaryDnsaclListACLsResponseSuccess `json:"success"`
	JSON    accountSecondaryDnsaclSecondaryDnsaclListACLsResponseJSON    `json:"-"`
}

// accountSecondaryDnsaclSecondaryDnsaclListACLsResponseJSON contains the JSON
// metadata for the struct [AccountSecondaryDnsaclSecondaryDnsaclListACLsResponse]
type accountSecondaryDnsaclSecondaryDnsaclListACLsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDnsaclSecondaryDnsaclListACLsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDnsaclSecondaryDnsaclListACLsResponseError struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    accountSecondaryDnsaclSecondaryDnsaclListACLsResponseErrorJSON `json:"-"`
}

// accountSecondaryDnsaclSecondaryDnsaclListACLsResponseErrorJSON contains the JSON
// metadata for the struct
// [AccountSecondaryDnsaclSecondaryDnsaclListACLsResponseError]
type accountSecondaryDnsaclSecondaryDnsaclListACLsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDnsaclSecondaryDnsaclListACLsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDnsaclSecondaryDnsaclListACLsResponseMessage struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    accountSecondaryDnsaclSecondaryDnsaclListACLsResponseMessageJSON `json:"-"`
}

// accountSecondaryDnsaclSecondaryDnsaclListACLsResponseMessageJSON contains the
// JSON metadata for the struct
// [AccountSecondaryDnsaclSecondaryDnsaclListACLsResponseMessage]
type accountSecondaryDnsaclSecondaryDnsaclListACLsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDnsaclSecondaryDnsaclListACLsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDnsaclSecondaryDnsaclListACLsResponseResult struct {
	ID interface{} `json:"id,required"`
	// Allowed IPv4/IPv6 address range of primary or secondary nameservers. This will
	// be applied for the entire account. The IP range is used to allow additional
	// NOTIFY IPs for secondary zones and IPs Cloudflare allows AXFR/IXFR requests from
	// for primary zones. CIDRs are limited to a maximum of /24 for IPv4 and /64 for
	// IPv6 respectively.
	IPRange string `json:"ip_range,required"`
	// The name of the acl.
	Name string                                                          `json:"name,required"`
	JSON accountSecondaryDnsaclSecondaryDnsaclListACLsResponseResultJSON `json:"-"`
}

// accountSecondaryDnsaclSecondaryDnsaclListACLsResponseResultJSON contains the
// JSON metadata for the struct
// [AccountSecondaryDnsaclSecondaryDnsaclListACLsResponseResult]
type accountSecondaryDnsaclSecondaryDnsaclListACLsResponseResultJSON struct {
	ID          apijson.Field
	IPRange     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDnsaclSecondaryDnsaclListACLsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDnsaclSecondaryDnsaclListACLsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                             `json:"total_count"`
	JSON       accountSecondaryDnsaclSecondaryDnsaclListACLsResponseResultInfoJSON `json:"-"`
}

// accountSecondaryDnsaclSecondaryDnsaclListACLsResponseResultInfoJSON contains the
// JSON metadata for the struct
// [AccountSecondaryDnsaclSecondaryDnsaclListACLsResponseResultInfo]
type accountSecondaryDnsaclSecondaryDnsaclListACLsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDnsaclSecondaryDnsaclListACLsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountSecondaryDnsaclSecondaryDnsaclListACLsResponseSuccess bool

const (
	AccountSecondaryDnsaclSecondaryDnsaclListACLsResponseSuccessTrue AccountSecondaryDnsaclSecondaryDnsaclListACLsResponseSuccess = true
)

type AccountSecondaryDNSACLUpdateParams struct {
	// Allowed IPv4/IPv6 address range of primary or secondary nameservers. This will
	// be applied for the entire account. The IP range is used to allow additional
	// NOTIFY IPs for secondary zones and IPs Cloudflare allows AXFR/IXFR requests from
	// for primary zones. CIDRs are limited to a maximum of /24 for IPv4 and /64 for
	// IPv6 respectively.
	IPRange param.Field[string] `json:"ip_range,required"`
	// The name of the acl.
	Name param.Field[string] `json:"name,required"`
}

func (r AccountSecondaryDNSACLUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountSecondaryDNSACLSecondaryDNSACLNewACLParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountSecondaryDNSACLSecondaryDNSACLNewACLParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
