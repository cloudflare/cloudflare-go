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

// AccountCustomNService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountCustomNService] method
// instead.
type AccountCustomNService struct {
	Options        []option.RequestOption
	Availabilities *AccountCustomNAvailabilityService
	Verifies       *AccountCustomNVerifyService
}

// NewAccountCustomNService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountCustomNService(opts ...option.RequestOption) (r *AccountCustomNService) {
	r = &AccountCustomNService{}
	r.Options = opts
	r.Availabilities = NewAccountCustomNAvailabilityService(opts...)
	r.Verifies = NewAccountCustomNVerifyService(opts...)
	return
}

// Delete Account Custom Nameserver
func (r *AccountCustomNService) Delete(ctx context.Context, identifier string, nsName string, opts ...option.RequestOption) (res *AccountCustomNDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/custom_ns/%s", identifier, nsName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Add Account Custom Nameserver
func (r *AccountCustomNService) AccountLevelCustomNameserversAddAccountCustomNameserver(ctx context.Context, identifier string, body AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverParams, opts ...option.RequestOption) (res *AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/custom_ns", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List an account's custom nameservers.
func (r *AccountCustomNService) AccountLevelCustomNameserversListAccountCustomNameservers(ctx context.Context, identifier string, opts ...option.RequestOption) (res *AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/custom_ns", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountCustomNDeleteResponse struct {
	Errors     []AccountCustomNDeleteResponseError    `json:"errors"`
	Messages   []AccountCustomNDeleteResponseMessage  `json:"messages"`
	Result     []interface{}                          `json:"result"`
	ResultInfo AccountCustomNDeleteResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountCustomNDeleteResponseSuccess `json:"success"`
	JSON    accountCustomNDeleteResponseJSON    `json:"-"`
}

// accountCustomNDeleteResponseJSON contains the JSON metadata for the struct
// [AccountCustomNDeleteResponse]
type accountCustomNDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCustomNDeleteResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    accountCustomNDeleteResponseErrorJSON `json:"-"`
}

// accountCustomNDeleteResponseErrorJSON contains the JSON metadata for the struct
// [AccountCustomNDeleteResponseError]
type accountCustomNDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCustomNDeleteResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accountCustomNDeleteResponseMessageJSON `json:"-"`
}

// accountCustomNDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountCustomNDeleteResponseMessage]
type accountCustomNDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCustomNDeleteResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                    `json:"total_count"`
	JSON       accountCustomNDeleteResponseResultInfoJSON `json:"-"`
}

// accountCustomNDeleteResponseResultInfoJSON contains the JSON metadata for the
// struct [AccountCustomNDeleteResponseResultInfo]
type accountCustomNDeleteResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNDeleteResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountCustomNDeleteResponseSuccess bool

const (
	AccountCustomNDeleteResponseSuccessTrue AccountCustomNDeleteResponseSuccess = true
)

type AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponse struct {
	Errors   []AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseError   `json:"errors"`
	Messages []AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseMessage `json:"messages"`
	// A single account custom nameserver.
	Result AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseResult `json:"result"`
	// Whether the API call was successful
	Success AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseSuccess `json:"success"`
	JSON    accountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseJSON    `json:"-"`
}

// accountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseJSON
// contains the JSON metadata for the struct
// [AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponse]
type accountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseError struct {
	Code    int64                                                                                  `json:"code,required"`
	Message string                                                                                 `json:"message,required"`
	JSON    accountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseErrorJSON `json:"-"`
}

// accountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseError]
type accountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseMessage struct {
	Code    int64                                                                                    `json:"code,required"`
	Message string                                                                                   `json:"message,required"`
	JSON    accountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseMessageJSON `json:"-"`
}

// accountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseMessage]
type accountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A single account custom nameserver.
type AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseResult struct {
	// A and AAAA records associated with the nameserver.
	DNSRecords []AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseResultDNSRecord `json:"dns_records,required"`
	// The FQDN of the name server.
	NsName string `json:"ns_name,required" format:"hostname"`
	// Verification status of the nameserver.
	Status AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseResultStatus `json:"status,required"`
	// Identifier
	ZoneTag string `json:"zone_tag,required"`
	// The number of the set that this name server belongs to.
	NsSet float64                                                                                 `json:"ns_set"`
	JSON  accountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseResultJSON `json:"-"`
}

// accountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseResultJSON
// contains the JSON metadata for the struct
// [AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseResult]
type accountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseResultJSON struct {
	DNSRecords  apijson.Field
	NsName      apijson.Field
	Status      apijson.Field
	ZoneTag     apijson.Field
	NsSet       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseResultDNSRecord struct {
	// DNS record type.
	Type AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseResultDNSRecordsType `json:"type"`
	// DNS record contents (an IPv4 or IPv6 address).
	Value string                                                                                           `json:"value"`
	JSON  accountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseResultDNSRecordJSON `json:"-"`
}

// accountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseResultDNSRecordJSON
// contains the JSON metadata for the struct
// [AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseResultDNSRecord]
type accountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseResultDNSRecordJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseResultDNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// DNS record type.
type AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseResultDNSRecordsType string

const (
	AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseResultDNSRecordsTypeA    AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseResultDNSRecordsType = "A"
	AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseResultDNSRecordsTypeAaaa AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseResultDNSRecordsType = "AAAA"
)

// Verification status of the nameserver.
type AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseResultStatus string

const (
	AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseResultStatusMoved    AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseResultStatus = "moved"
	AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseResultStatusPending  AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseResultStatus = "pending"
	AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseResultStatusVerified AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseResultStatus = "verified"
)

// Whether the API call was successful
type AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseSuccess bool

const (
	AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseSuccessTrue AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverResponseSuccess = true
)

type AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponse struct {
	Errors     []AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseError    `json:"errors"`
	Messages   []AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseMessage  `json:"messages"`
	Result     []AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResult   `json:"result"`
	ResultInfo AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseSuccess `json:"success"`
	JSON    accountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseJSON    `json:"-"`
}

// accountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseJSON
// contains the JSON metadata for the struct
// [AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponse]
type accountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseError struct {
	Code    int64                                                                                    `json:"code,required"`
	Message string                                                                                   `json:"message,required"`
	JSON    accountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseErrorJSON `json:"-"`
}

// accountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseError]
type accountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseMessage struct {
	Code    int64                                                                                      `json:"code,required"`
	Message string                                                                                     `json:"message,required"`
	JSON    accountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseMessageJSON `json:"-"`
}

// accountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseMessage]
type accountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A single account custom nameserver.
type AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResult struct {
	// A and AAAA records associated with the nameserver.
	DNSRecords []AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResultDNSRecord `json:"dns_records,required"`
	// The FQDN of the name server.
	NsName string `json:"ns_name,required" format:"hostname"`
	// Verification status of the nameserver.
	Status AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResultStatus `json:"status,required"`
	// Identifier
	ZoneTag string `json:"zone_tag,required"`
	// The number of the set that this name server belongs to.
	NsSet float64                                                                                   `json:"ns_set"`
	JSON  accountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResultJSON `json:"-"`
}

// accountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResultJSON
// contains the JSON metadata for the struct
// [AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResult]
type accountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResultJSON struct {
	DNSRecords  apijson.Field
	NsName      apijson.Field
	Status      apijson.Field
	ZoneTag     apijson.Field
	NsSet       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResultDNSRecord struct {
	// DNS record type.
	Type AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResultDNSRecordsType `json:"type"`
	// DNS record contents (an IPv4 or IPv6 address).
	Value string                                                                                             `json:"value"`
	JSON  accountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResultDNSRecordJSON `json:"-"`
}

// accountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResultDNSRecordJSON
// contains the JSON metadata for the struct
// [AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResultDNSRecord]
type accountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResultDNSRecordJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResultDNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// DNS record type.
type AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResultDNSRecordsType string

const (
	AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResultDNSRecordsTypeA    AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResultDNSRecordsType = "A"
	AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResultDNSRecordsTypeAaaa AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResultDNSRecordsType = "AAAA"
)

// Verification status of the nameserver.
type AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResultStatus string

const (
	AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResultStatusMoved    AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResultStatus = "moved"
	AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResultStatusPending  AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResultStatus = "pending"
	AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResultStatusVerified AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResultStatus = "verified"
)

type AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                       `json:"total_count"`
	JSON       accountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResultInfoJSON `json:"-"`
}

// accountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResultInfo]
type accountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseSuccess bool

const (
	AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseSuccessTrue AccountCustomNAccountLevelCustomNameserversListAccountCustomNameserversResponseSuccess = true
)

type AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverParams struct {
	// The FQDN of the name server.
	NsName param.Field[string] `json:"ns_name,required" format:"hostname"`
	// The number of the set that this name server belongs to.
	NsSet param.Field[float64] `json:"ns_set"`
}

func (r AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
