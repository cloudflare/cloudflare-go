// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountCustomNVerifyService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountCustomNVerifyService]
// method instead.
type AccountCustomNVerifyService struct {
	Options []option.RequestOption
}

// NewAccountCustomNVerifyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountCustomNVerifyService(opts ...option.RequestOption) (r *AccountCustomNVerifyService) {
	r = &AccountCustomNVerifyService{}
	r.Options = opts
	return
}

// Verify Account Custom Nameserver Glue Records
func (r *AccountCustomNVerifyService) AccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecords(ctx context.Context, identifier string, opts ...option.RequestOption) (res *AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/custom_ns/verify", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponse struct {
	Errors     []AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseError    `json:"errors"`
	Messages   []AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseMessage  `json:"messages"`
	Result     []AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResult   `json:"result"`
	ResultInfo AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseSuccess `json:"success"`
	JSON    accountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseJSON    `json:"-"`
}

// accountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseJSON
// contains the JSON metadata for the struct
// [AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponse]
type accountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseError struct {
	Code    int64                                                                                                      `json:"code,required"`
	Message string                                                                                                     `json:"message,required"`
	JSON    accountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseErrorJSON `json:"-"`
}

// accountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseError]
type accountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseMessage struct {
	Code    int64                                                                                                        `json:"code,required"`
	Message string                                                                                                       `json:"message,required"`
	JSON    accountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseMessageJSON `json:"-"`
}

// accountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseMessage]
type accountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A single account custom nameserver.
type AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResult struct {
	// A and AAAA records associated with the nameserver.
	DNSRecords []AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResultDNSRecord `json:"dns_records,required"`
	// The FQDN of the name server.
	NsName string `json:"ns_name,required" format:"hostname"`
	// Verification status of the nameserver.
	Status AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResultStatus `json:"status,required"`
	// Identifier
	ZoneTag string `json:"zone_tag,required"`
	// The number of the set that this name server belongs to.
	NsSet float64                                                                                                     `json:"ns_set"`
	JSON  accountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResultJSON `json:"-"`
}

// accountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResult]
type accountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResultJSON struct {
	DNSRecords  apijson.Field
	NsName      apijson.Field
	Status      apijson.Field
	ZoneTag     apijson.Field
	NsSet       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResultDNSRecord struct {
	// DNS record type.
	Type AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResultDNSRecordsType `json:"type"`
	// DNS record contents (an IPv4 or IPv6 address).
	Value string                                                                                                               `json:"value"`
	JSON  accountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResultDNSRecordJSON `json:"-"`
}

// accountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResultDNSRecordJSON
// contains the JSON metadata for the struct
// [AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResultDNSRecord]
type accountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResultDNSRecordJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResultDNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// DNS record type.
type AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResultDNSRecordsType string

const (
	AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResultDNSRecordsTypeA    AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResultDNSRecordsType = "A"
	AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResultDNSRecordsTypeAaaa AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResultDNSRecordsType = "AAAA"
)

// Verification status of the nameserver.
type AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResultStatus string

const (
	AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResultStatusMoved    AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResultStatus = "moved"
	AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResultStatusPending  AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResultStatus = "pending"
	AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResultStatusVerified AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResultStatus = "verified"
)

type AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                         `json:"total_count"`
	JSON       accountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResultInfoJSON `json:"-"`
}

// accountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResultInfo]
type accountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseSuccess bool

const (
	AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseSuccessTrue AccountCustomNVerifyAccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecordsResponseSuccess = true
)
