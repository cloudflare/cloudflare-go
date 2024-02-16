// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// CustomNService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewCustomNService] method instead.
type CustomNService struct {
	Options        []option.RequestOption
	Availabilities *CustomNAvailabilityService
	Verifies       *CustomNVerifyService
}

// NewCustomNService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCustomNService(opts ...option.RequestOption) (r *CustomNService) {
	r = &CustomNService{}
	r.Options = opts
	r.Availabilities = NewCustomNAvailabilityService(opts...)
	r.Verifies = NewCustomNVerifyService(opts...)
	return
}

// Add Account Custom Nameserver
func (r *CustomNService) New(ctx context.Context, accountID string, body CustomNNewParams, opts ...option.RequestOption) (res *CustomNNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomNNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/custom_ns", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List an account's custom nameservers.
func (r *CustomNService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]CustomNListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomNListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/custom_ns", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete Account Custom Nameserver
func (r *CustomNService) Delete(ctx context.Context, accountID string, customNsID string, opts ...option.RequestOption) (res *CustomNDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomNDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/custom_ns/%s", accountID, customNsID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A single account custom nameserver.
type CustomNNewResponse struct {
	// A and AAAA records associated with the nameserver.
	DNSRecords []CustomNNewResponseDNSRecord `json:"dns_records,required"`
	// The FQDN of the name server.
	NsName string `json:"ns_name,required" format:"hostname"`
	// Verification status of the nameserver.
	Status CustomNNewResponseStatus `json:"status,required"`
	// Identifier
	ZoneTag string `json:"zone_tag,required"`
	// The number of the set that this name server belongs to.
	NsSet float64                `json:"ns_set"`
	JSON  customNNewResponseJSON `json:"-"`
}

// customNNewResponseJSON contains the JSON metadata for the struct
// [CustomNNewResponse]
type customNNewResponseJSON struct {
	DNSRecords  apijson.Field
	NsName      apijson.Field
	Status      apijson.Field
	ZoneTag     apijson.Field
	NsSet       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomNNewResponseDNSRecord struct {
	// DNS record type.
	Type CustomNNewResponseDNSRecordsType `json:"type"`
	// DNS record contents (an IPv4 or IPv6 address).
	Value string                          `json:"value"`
	JSON  customNNewResponseDNSRecordJSON `json:"-"`
}

// customNNewResponseDNSRecordJSON contains the JSON metadata for the struct
// [CustomNNewResponseDNSRecord]
type customNNewResponseDNSRecordJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNNewResponseDNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// DNS record type.
type CustomNNewResponseDNSRecordsType string

const (
	CustomNNewResponseDNSRecordsTypeA    CustomNNewResponseDNSRecordsType = "A"
	CustomNNewResponseDNSRecordsTypeAaaa CustomNNewResponseDNSRecordsType = "AAAA"
)

// Verification status of the nameserver.
type CustomNNewResponseStatus string

const (
	CustomNNewResponseStatusMoved    CustomNNewResponseStatus = "moved"
	CustomNNewResponseStatusPending  CustomNNewResponseStatus = "pending"
	CustomNNewResponseStatusVerified CustomNNewResponseStatus = "verified"
)

// A single account custom nameserver.
type CustomNListResponse struct {
	// A and AAAA records associated with the nameserver.
	DNSRecords []CustomNListResponseDNSRecord `json:"dns_records,required"`
	// The FQDN of the name server.
	NsName string `json:"ns_name,required" format:"hostname"`
	// Verification status of the nameserver.
	Status CustomNListResponseStatus `json:"status,required"`
	// Identifier
	ZoneTag string `json:"zone_tag,required"`
	// The number of the set that this name server belongs to.
	NsSet float64                 `json:"ns_set"`
	JSON  customNListResponseJSON `json:"-"`
}

// customNListResponseJSON contains the JSON metadata for the struct
// [CustomNListResponse]
type customNListResponseJSON struct {
	DNSRecords  apijson.Field
	NsName      apijson.Field
	Status      apijson.Field
	ZoneTag     apijson.Field
	NsSet       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomNListResponseDNSRecord struct {
	// DNS record type.
	Type CustomNListResponseDNSRecordsType `json:"type"`
	// DNS record contents (an IPv4 or IPv6 address).
	Value string                           `json:"value"`
	JSON  customNListResponseDNSRecordJSON `json:"-"`
}

// customNListResponseDNSRecordJSON contains the JSON metadata for the struct
// [CustomNListResponseDNSRecord]
type customNListResponseDNSRecordJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNListResponseDNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// DNS record type.
type CustomNListResponseDNSRecordsType string

const (
	CustomNListResponseDNSRecordsTypeA    CustomNListResponseDNSRecordsType = "A"
	CustomNListResponseDNSRecordsTypeAaaa CustomNListResponseDNSRecordsType = "AAAA"
)

// Verification status of the nameserver.
type CustomNListResponseStatus string

const (
	CustomNListResponseStatusMoved    CustomNListResponseStatus = "moved"
	CustomNListResponseStatusPending  CustomNListResponseStatus = "pending"
	CustomNListResponseStatusVerified CustomNListResponseStatus = "verified"
)

// Union satisfied by [CustomNDeleteResponseUnknown], [CustomNDeleteResponseArray]
// or [shared.UnionString].
type CustomNDeleteResponse interface {
	ImplementsCustomNDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomNDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CustomNDeleteResponseArray []interface{}

func (r CustomNDeleteResponseArray) ImplementsCustomNDeleteResponse() {}

type CustomNNewParams struct {
	// The FQDN of the name server.
	NsName param.Field[string] `json:"ns_name,required" format:"hostname"`
	// The number of the set that this name server belongs to.
	NsSet param.Field[float64] `json:"ns_set"`
}

func (r CustomNNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomNNewResponseEnvelope struct {
	Errors   []CustomNNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomNNewResponseEnvelopeMessages `json:"messages,required"`
	// A single account custom nameserver.
	Result CustomNNewResponse `json:"result,required"`
	// Whether the API call was successful
	Success CustomNNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    customNNewResponseEnvelopeJSON    `json:"-"`
}

// customNNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [CustomNNewResponseEnvelope]
type customNNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomNNewResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    customNNewResponseEnvelopeErrorsJSON `json:"-"`
}

// customNNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [CustomNNewResponseEnvelopeErrors]
type customNNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomNNewResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    customNNewResponseEnvelopeMessagesJSON `json:"-"`
}

// customNNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [CustomNNewResponseEnvelopeMessages]
type customNNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CustomNNewResponseEnvelopeSuccess bool

const (
	CustomNNewResponseEnvelopeSuccessTrue CustomNNewResponseEnvelopeSuccess = true
)

type CustomNListResponseEnvelope struct {
	Errors   []CustomNListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomNListResponseEnvelopeMessages `json:"messages,required"`
	Result   []CustomNListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    CustomNListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo CustomNListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       customNListResponseEnvelopeJSON       `json:"-"`
}

// customNListResponseEnvelopeJSON contains the JSON metadata for the struct
// [CustomNListResponseEnvelope]
type customNListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomNListResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    customNListResponseEnvelopeErrorsJSON `json:"-"`
}

// customNListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [CustomNListResponseEnvelopeErrors]
type customNListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomNListResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    customNListResponseEnvelopeMessagesJSON `json:"-"`
}

// customNListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CustomNListResponseEnvelopeMessages]
type customNListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CustomNListResponseEnvelopeSuccess bool

const (
	CustomNListResponseEnvelopeSuccessTrue CustomNListResponseEnvelopeSuccess = true
)

type CustomNListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                   `json:"total_count"`
	JSON       customNListResponseEnvelopeResultInfoJSON `json:"-"`
}

// customNListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [CustomNListResponseEnvelopeResultInfo]
type customNListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomNDeleteResponseEnvelope struct {
	Errors   []CustomNDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomNDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   CustomNDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    CustomNDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo CustomNDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       customNDeleteResponseEnvelopeJSON       `json:"-"`
}

// customNDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [CustomNDeleteResponseEnvelope]
type customNDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomNDeleteResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    customNDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// customNDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CustomNDeleteResponseEnvelopeErrors]
type customNDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomNDeleteResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    customNDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// customNDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CustomNDeleteResponseEnvelopeMessages]
type customNDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CustomNDeleteResponseEnvelopeSuccess bool

const (
	CustomNDeleteResponseEnvelopeSuccessTrue CustomNDeleteResponseEnvelopeSuccess = true
)

type CustomNDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                     `json:"total_count"`
	JSON       customNDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// customNDeleteResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [CustomNDeleteResponseEnvelopeResultInfo]
type customNDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
