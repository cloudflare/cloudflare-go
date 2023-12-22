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
func (r *AccountCustomNService) Delete(ctx context.Context, identifier string, nsName string, opts ...option.RequestOption) (res *EmptyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/custom_ns/%s", identifier, nsName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Add Account Custom Nameserver
func (r *AccountCustomNService) AccountLevelCustomNameserversAddAccountCustomNameserver(ctx context.Context, identifier string, body AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverParams, opts ...option.RequestOption) (res *AcnsResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/custom_ns", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List an account's custom nameservers.
func (r *AccountCustomNService) AccountLevelCustomNameserversListAccountCustomNameservers(ctx context.Context, identifier string, opts ...option.RequestOption) (res *AcnsResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/custom_ns", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AcnsResponseCollection struct {
	Errors     []AcnsResponseCollectionError    `json:"errors"`
	Messages   []AcnsResponseCollectionMessage  `json:"messages"`
	Result     []AcnsResponseCollectionResult   `json:"result"`
	ResultInfo AcnsResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AcnsResponseCollectionSuccess `json:"success"`
	JSON    acnsResponseCollectionJSON    `json:"-"`
}

// acnsResponseCollectionJSON contains the JSON metadata for the struct
// [AcnsResponseCollection]
type acnsResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AcnsResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AcnsResponseCollectionError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    acnsResponseCollectionErrorJSON `json:"-"`
}

// acnsResponseCollectionErrorJSON contains the JSON metadata for the struct
// [AcnsResponseCollectionError]
type acnsResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AcnsResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AcnsResponseCollectionMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    acnsResponseCollectionMessageJSON `json:"-"`
}

// acnsResponseCollectionMessageJSON contains the JSON metadata for the struct
// [AcnsResponseCollectionMessage]
type acnsResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AcnsResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A single account custom nameserver.
type AcnsResponseCollectionResult struct {
	// A and AAAA records associated with the nameserver.
	DNSRecords []AcnsResponseCollectionResultDNSRecord `json:"dns_records,required"`
	// The FQDN of the name server.
	NsName string `json:"ns_name,required" format:"hostname"`
	// Verification status of the nameserver.
	Status AcnsResponseCollectionResultStatus `json:"status,required"`
	// Identifier
	ZoneTag string                           `json:"zone_tag,required"`
	JSON    acnsResponseCollectionResultJSON `json:"-"`
}

// acnsResponseCollectionResultJSON contains the JSON metadata for the struct
// [AcnsResponseCollectionResult]
type acnsResponseCollectionResultJSON struct {
	DNSRecords  apijson.Field
	NsName      apijson.Field
	Status      apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AcnsResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AcnsResponseCollectionResultDNSRecord struct {
	// DNS record type.
	Type AcnsResponseCollectionResultDNSRecordsType `json:"type"`
	// DNS record contents (an IPv4 or IPv6 address).
	Value string                                    `json:"value"`
	JSON  acnsResponseCollectionResultDNSRecordJSON `json:"-"`
}

// acnsResponseCollectionResultDNSRecordJSON contains the JSON metadata for the
// struct [AcnsResponseCollectionResultDNSRecord]
type acnsResponseCollectionResultDNSRecordJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AcnsResponseCollectionResultDNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// DNS record type.
type AcnsResponseCollectionResultDNSRecordsType string

const (
	AcnsResponseCollectionResultDNSRecordsTypeA    AcnsResponseCollectionResultDNSRecordsType = "A"
	AcnsResponseCollectionResultDNSRecordsTypeAaaa AcnsResponseCollectionResultDNSRecordsType = "AAAA"
)

// Verification status of the nameserver.
type AcnsResponseCollectionResultStatus string

const (
	AcnsResponseCollectionResultStatusMoved    AcnsResponseCollectionResultStatus = "moved"
	AcnsResponseCollectionResultStatusPending  AcnsResponseCollectionResultStatus = "pending"
	AcnsResponseCollectionResultStatusVerified AcnsResponseCollectionResultStatus = "verified"
)

type AcnsResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                              `json:"total_count"`
	JSON       acnsResponseCollectionResultInfoJSON `json:"-"`
}

// acnsResponseCollectionResultInfoJSON contains the JSON metadata for the struct
// [AcnsResponseCollectionResultInfo]
type acnsResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AcnsResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AcnsResponseCollectionSuccess bool

const (
	AcnsResponseCollectionSuccessTrue AcnsResponseCollectionSuccess = true
)

type AcnsResponseSingle struct {
	Errors   []AcnsResponseSingleError   `json:"errors"`
	Messages []AcnsResponseSingleMessage `json:"messages"`
	// A single account custom nameserver.
	Result AcnsResponseSingleResult `json:"result"`
	// Whether the API call was successful
	Success AcnsResponseSingleSuccess `json:"success"`
	JSON    acnsResponseSingleJSON    `json:"-"`
}

// acnsResponseSingleJSON contains the JSON metadata for the struct
// [AcnsResponseSingle]
type acnsResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AcnsResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AcnsResponseSingleError struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    acnsResponseSingleErrorJSON `json:"-"`
}

// acnsResponseSingleErrorJSON contains the JSON metadata for the struct
// [AcnsResponseSingleError]
type acnsResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AcnsResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AcnsResponseSingleMessage struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    acnsResponseSingleMessageJSON `json:"-"`
}

// acnsResponseSingleMessageJSON contains the JSON metadata for the struct
// [AcnsResponseSingleMessage]
type acnsResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AcnsResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A single account custom nameserver.
type AcnsResponseSingleResult struct {
	// A and AAAA records associated with the nameserver.
	DNSRecords []AcnsResponseSingleResultDNSRecord `json:"dns_records,required"`
	// The FQDN of the name server.
	NsName string `json:"ns_name,required" format:"hostname"`
	// Verification status of the nameserver.
	Status AcnsResponseSingleResultStatus `json:"status,required"`
	// Identifier
	ZoneTag string                       `json:"zone_tag,required"`
	JSON    acnsResponseSingleResultJSON `json:"-"`
}

// acnsResponseSingleResultJSON contains the JSON metadata for the struct
// [AcnsResponseSingleResult]
type acnsResponseSingleResultJSON struct {
	DNSRecords  apijson.Field
	NsName      apijson.Field
	Status      apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AcnsResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AcnsResponseSingleResultDNSRecord struct {
	// DNS record type.
	Type AcnsResponseSingleResultDNSRecordsType `json:"type"`
	// DNS record contents (an IPv4 or IPv6 address).
	Value string                                `json:"value"`
	JSON  acnsResponseSingleResultDNSRecordJSON `json:"-"`
}

// acnsResponseSingleResultDNSRecordJSON contains the JSON metadata for the struct
// [AcnsResponseSingleResultDNSRecord]
type acnsResponseSingleResultDNSRecordJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AcnsResponseSingleResultDNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// DNS record type.
type AcnsResponseSingleResultDNSRecordsType string

const (
	AcnsResponseSingleResultDNSRecordsTypeA    AcnsResponseSingleResultDNSRecordsType = "A"
	AcnsResponseSingleResultDNSRecordsTypeAaaa AcnsResponseSingleResultDNSRecordsType = "AAAA"
)

// Verification status of the nameserver.
type AcnsResponseSingleResultStatus string

const (
	AcnsResponseSingleResultStatusMoved    AcnsResponseSingleResultStatus = "moved"
	AcnsResponseSingleResultStatusPending  AcnsResponseSingleResultStatus = "pending"
	AcnsResponseSingleResultStatusVerified AcnsResponseSingleResultStatus = "verified"
)

// Whether the API call was successful
type AcnsResponseSingleSuccess bool

const (
	AcnsResponseSingleSuccessTrue AcnsResponseSingleSuccess = true
)

type EmptyResponse struct {
	Errors     []EmptyResponseError    `json:"errors"`
	Messages   []EmptyResponseMessage  `json:"messages"`
	Result     []interface{}           `json:"result"`
	ResultInfo EmptyResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success EmptyResponseSuccess `json:"success"`
	JSON    emptyResponseJSON    `json:"-"`
}

// emptyResponseJSON contains the JSON metadata for the struct [EmptyResponse]
type emptyResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmptyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmptyResponseError struct {
	Code    int64                  `json:"code,required"`
	Message string                 `json:"message,required"`
	JSON    emptyResponseErrorJSON `json:"-"`
}

// emptyResponseErrorJSON contains the JSON metadata for the struct
// [EmptyResponseError]
type emptyResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmptyResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmptyResponseMessage struct {
	Code    int64                    `json:"code,required"`
	Message string                   `json:"message,required"`
	JSON    emptyResponseMessageJSON `json:"-"`
}

// emptyResponseMessageJSON contains the JSON metadata for the struct
// [EmptyResponseMessage]
type emptyResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmptyResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmptyResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                     `json:"total_count"`
	JSON       emptyResponseResultInfoJSON `json:"-"`
}

// emptyResponseResultInfoJSON contains the JSON metadata for the struct
// [EmptyResponseResultInfo]
type emptyResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmptyResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmptyResponseSuccess bool

const (
	EmptyResponseSuccessTrue EmptyResponseSuccess = true
)

type AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverParams struct {
	// The FQDN of the name server.
	NsName param.Field[string] `json:"ns_name,required" format:"hostname"`
}

func (r AccountCustomNAccountLevelCustomNameserversAddAccountCustomNameserverParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
