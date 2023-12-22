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
func (r *AccountSecondaryDNSACLService) Get(ctx context.Context, accountIdentifier interface{}, identifier interface{}, opts ...option.RequestOption) (res *SingleResponseCboRkyA7, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/secondary_dns/acls/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify ACL.
func (r *AccountSecondaryDNSACLService) Update(ctx context.Context, accountIdentifier interface{}, identifier interface{}, body AccountSecondaryDNSACLUpdateParams, opts ...option.RequestOption) (res *SingleResponseCboRkyA7, err error) {
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
func (r *AccountSecondaryDNSACLService) SecondaryDNSACLNewACL(ctx context.Context, accountIdentifier interface{}, body AccountSecondaryDNSACLSecondaryDNSACLNewACLParams, opts ...option.RequestOption) (res *SingleResponseCboRkyA7, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/secondary_dns/acls", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List ACLs.
func (r *AccountSecondaryDNSACLService) SecondaryDNSACLListACLs(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *ResponseCollectionB1aY2eai, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/secondary_dns/acls", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ResponseCollectionB1aY2eai struct {
	Errors     []ResponseCollectionB1aY2eaiError    `json:"errors"`
	Messages   []ResponseCollectionB1aY2eaiMessage  `json:"messages"`
	Result     []ResponseCollectionB1aY2eaiResult   `json:"result"`
	ResultInfo ResponseCollectionB1aY2eaiResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ResponseCollectionB1aY2eaiSuccess `json:"success"`
	JSON    responseCollectionB1aY2eaiJSON    `json:"-"`
}

// responseCollectionB1aY2eaiJSON contains the JSON metadata for the struct
// [ResponseCollectionB1aY2eai]
type responseCollectionB1aY2eaiJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionB1aY2eai) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionB1aY2eaiError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    responseCollectionB1aY2eaiErrorJSON `json:"-"`
}

// responseCollectionB1aY2eaiErrorJSON contains the JSON metadata for the struct
// [ResponseCollectionB1aY2eaiError]
type responseCollectionB1aY2eaiErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionB1aY2eaiError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionB1aY2eaiMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    responseCollectionB1aY2eaiMessageJSON `json:"-"`
}

// responseCollectionB1aY2eaiMessageJSON contains the JSON metadata for the struct
// [ResponseCollectionB1aY2eaiMessage]
type responseCollectionB1aY2eaiMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionB1aY2eaiMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionB1aY2eaiResult struct {
	ID interface{} `json:"id,required"`
	// Allowed IPv4/IPv6 address range of primary or secondary nameservers. This will
	// be applied for the entire account. The IP range is used to allow additional
	// NOTIFY IPs for secondary zones and IPs Cloudflare allows AXFR/IXFR requests from
	// for primary zones. CIDRs are limited to a maximum of /24 for IPv4 and /64 for
	// IPv6 respectively.
	IPRange string `json:"ip_range,required"`
	// The name of the acl.
	Name string                               `json:"name,required"`
	JSON responseCollectionB1aY2eaiResultJSON `json:"-"`
}

// responseCollectionB1aY2eaiResultJSON contains the JSON metadata for the struct
// [ResponseCollectionB1aY2eaiResult]
type responseCollectionB1aY2eaiResultJSON struct {
	ID          apijson.Field
	IPRange     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionB1aY2eaiResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionB1aY2eaiResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       responseCollectionB1aY2eaiResultInfoJSON `json:"-"`
}

// responseCollectionB1aY2eaiResultInfoJSON contains the JSON metadata for the
// struct [ResponseCollectionB1aY2eaiResultInfo]
type responseCollectionB1aY2eaiResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionB1aY2eaiResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ResponseCollectionB1aY2eaiSuccess bool

const (
	ResponseCollectionB1aY2eaiSuccessTrue ResponseCollectionB1aY2eaiSuccess = true
)

type SingleResponseCboRkyA7 struct {
	Errors   []SingleResponseCboRkyA7Error   `json:"errors"`
	Messages []SingleResponseCboRkyA7Message `json:"messages"`
	Result   SingleResponseCboRkyA7Result    `json:"result"`
	// Whether the API call was successful
	Success SingleResponseCboRkyA7Success `json:"success"`
	JSON    singleResponseCboRkyA7JSON    `json:"-"`
}

// singleResponseCboRkyA7JSON contains the JSON metadata for the struct
// [SingleResponseCboRkyA7]
type singleResponseCboRkyA7JSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseCboRkyA7) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseCboRkyA7Error struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    singleResponseCboRkyA7ErrorJSON `json:"-"`
}

// singleResponseCboRkyA7ErrorJSON contains the JSON metadata for the struct
// [SingleResponseCboRkyA7Error]
type singleResponseCboRkyA7ErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseCboRkyA7Error) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseCboRkyA7Message struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    singleResponseCboRkyA7MessageJSON `json:"-"`
}

// singleResponseCboRkyA7MessageJSON contains the JSON metadata for the struct
// [SingleResponseCboRkyA7Message]
type singleResponseCboRkyA7MessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseCboRkyA7Message) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseCboRkyA7Result struct {
	ID interface{} `json:"id,required"`
	// Allowed IPv4/IPv6 address range of primary or secondary nameservers. This will
	// be applied for the entire account. The IP range is used to allow additional
	// NOTIFY IPs for secondary zones and IPs Cloudflare allows AXFR/IXFR requests from
	// for primary zones. CIDRs are limited to a maximum of /24 for IPv4 and /64 for
	// IPv6 respectively.
	IPRange string `json:"ip_range,required"`
	// The name of the acl.
	Name string                           `json:"name,required"`
	JSON singleResponseCboRkyA7ResultJSON `json:"-"`
}

// singleResponseCboRkyA7ResultJSON contains the JSON metadata for the struct
// [SingleResponseCboRkyA7Result]
type singleResponseCboRkyA7ResultJSON struct {
	ID          apijson.Field
	IPRange     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseCboRkyA7Result) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SingleResponseCboRkyA7Success bool

const (
	SingleResponseCboRkyA7SuccessTrue SingleResponseCboRkyA7Success = true
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
