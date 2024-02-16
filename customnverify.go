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

// CustomNVerifyService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCustomNVerifyService] method
// instead.
type CustomNVerifyService struct {
	Options []option.RequestOption
}

// NewCustomNVerifyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomNVerifyService(opts ...option.RequestOption) (r *CustomNVerifyService) {
	r = &CustomNVerifyService{}
	r.Options = opts
	return
}

// Verify Account Custom Nameserver Glue Records
func (r *CustomNVerifyService) Update(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]CustomNVerifyUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomNVerifyUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/custom_ns/verify", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A single account custom nameserver.
type CustomNVerifyUpdateResponse struct {
	// A and AAAA records associated with the nameserver.
	DNSRecords []CustomNVerifyUpdateResponseDNSRecord `json:"dns_records,required"`
	// The FQDN of the name server.
	NsName string `json:"ns_name,required" format:"hostname"`
	// Verification status of the nameserver.
	Status CustomNVerifyUpdateResponseStatus `json:"status,required"`
	// Identifier
	ZoneTag string `json:"zone_tag,required"`
	// The number of the set that this name server belongs to.
	NsSet float64                         `json:"ns_set"`
	JSON  customNVerifyUpdateResponseJSON `json:"-"`
}

// customNVerifyUpdateResponseJSON contains the JSON metadata for the struct
// [CustomNVerifyUpdateResponse]
type customNVerifyUpdateResponseJSON struct {
	DNSRecords  apijson.Field
	NsName      apijson.Field
	Status      apijson.Field
	ZoneTag     apijson.Field
	NsSet       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNVerifyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomNVerifyUpdateResponseDNSRecord struct {
	// DNS record type.
	Type CustomNVerifyUpdateResponseDNSRecordsType `json:"type"`
	// DNS record contents (an IPv4 or IPv6 address).
	Value string                                   `json:"value"`
	JSON  customNVerifyUpdateResponseDNSRecordJSON `json:"-"`
}

// customNVerifyUpdateResponseDNSRecordJSON contains the JSON metadata for the
// struct [CustomNVerifyUpdateResponseDNSRecord]
type customNVerifyUpdateResponseDNSRecordJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNVerifyUpdateResponseDNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// DNS record type.
type CustomNVerifyUpdateResponseDNSRecordsType string

const (
	CustomNVerifyUpdateResponseDNSRecordsTypeA    CustomNVerifyUpdateResponseDNSRecordsType = "A"
	CustomNVerifyUpdateResponseDNSRecordsTypeAaaa CustomNVerifyUpdateResponseDNSRecordsType = "AAAA"
)

// Verification status of the nameserver.
type CustomNVerifyUpdateResponseStatus string

const (
	CustomNVerifyUpdateResponseStatusMoved    CustomNVerifyUpdateResponseStatus = "moved"
	CustomNVerifyUpdateResponseStatusPending  CustomNVerifyUpdateResponseStatus = "pending"
	CustomNVerifyUpdateResponseStatusVerified CustomNVerifyUpdateResponseStatus = "verified"
)

type CustomNVerifyUpdateResponseEnvelope struct {
	Errors   []CustomNVerifyUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomNVerifyUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   []CustomNVerifyUpdateResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    CustomNVerifyUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo CustomNVerifyUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       customNVerifyUpdateResponseEnvelopeJSON       `json:"-"`
}

// customNVerifyUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [CustomNVerifyUpdateResponseEnvelope]
type customNVerifyUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNVerifyUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomNVerifyUpdateResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    customNVerifyUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// customNVerifyUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CustomNVerifyUpdateResponseEnvelopeErrors]
type customNVerifyUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNVerifyUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomNVerifyUpdateResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    customNVerifyUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// customNVerifyUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CustomNVerifyUpdateResponseEnvelopeMessages]
type customNVerifyUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNVerifyUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CustomNVerifyUpdateResponseEnvelopeSuccess bool

const (
	CustomNVerifyUpdateResponseEnvelopeSuccessTrue CustomNVerifyUpdateResponseEnvelopeSuccess = true
)

type CustomNVerifyUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                           `json:"total_count"`
	JSON       customNVerifyUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// customNVerifyUpdateResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [CustomNVerifyUpdateResponseEnvelopeResultInfo]
type customNVerifyUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNVerifyUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
