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
func (r *CustomNVerifyService) New(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]CustomNVerifyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomNVerifyNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/custom_ns/verify", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A single account custom nameserver.
type CustomNVerifyNewResponse struct {
	// A and AAAA records associated with the nameserver.
	DNSRecords []CustomNVerifyNewResponseDNSRecord `json:"dns_records,required"`
	// The FQDN of the name server.
	NsName string `json:"ns_name,required" format:"hostname"`
	// Verification status of the nameserver.
	Status CustomNVerifyNewResponseStatus `json:"status,required"`
	// Identifier
	ZoneTag string `json:"zone_tag,required"`
	// The number of the set that this name server belongs to.
	NsSet float64                      `json:"ns_set"`
	JSON  customNVerifyNewResponseJSON `json:"-"`
}

// customNVerifyNewResponseJSON contains the JSON metadata for the struct
// [CustomNVerifyNewResponse]
type customNVerifyNewResponseJSON struct {
	DNSRecords  apijson.Field
	NsName      apijson.Field
	Status      apijson.Field
	ZoneTag     apijson.Field
	NsSet       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNVerifyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomNVerifyNewResponseDNSRecord struct {
	// DNS record type.
	Type CustomNVerifyNewResponseDNSRecordsType `json:"type"`
	// DNS record contents (an IPv4 or IPv6 address).
	Value string                                `json:"value"`
	JSON  customNVerifyNewResponseDNSRecordJSON `json:"-"`
}

// customNVerifyNewResponseDNSRecordJSON contains the JSON metadata for the struct
// [CustomNVerifyNewResponseDNSRecord]
type customNVerifyNewResponseDNSRecordJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNVerifyNewResponseDNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// DNS record type.
type CustomNVerifyNewResponseDNSRecordsType string

const (
	CustomNVerifyNewResponseDNSRecordsTypeA    CustomNVerifyNewResponseDNSRecordsType = "A"
	CustomNVerifyNewResponseDNSRecordsTypeAaaa CustomNVerifyNewResponseDNSRecordsType = "AAAA"
)

// Verification status of the nameserver.
type CustomNVerifyNewResponseStatus string

const (
	CustomNVerifyNewResponseStatusMoved    CustomNVerifyNewResponseStatus = "moved"
	CustomNVerifyNewResponseStatusPending  CustomNVerifyNewResponseStatus = "pending"
	CustomNVerifyNewResponseStatusVerified CustomNVerifyNewResponseStatus = "verified"
)

type CustomNVerifyNewResponseEnvelope struct {
	Errors   []CustomNVerifyNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomNVerifyNewResponseEnvelopeMessages `json:"messages,required"`
	Result   []CustomNVerifyNewResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    CustomNVerifyNewResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo CustomNVerifyNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       customNVerifyNewResponseEnvelopeJSON       `json:"-"`
}

// customNVerifyNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [CustomNVerifyNewResponseEnvelope]
type customNVerifyNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNVerifyNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomNVerifyNewResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    customNVerifyNewResponseEnvelopeErrorsJSON `json:"-"`
}

// customNVerifyNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CustomNVerifyNewResponseEnvelopeErrors]
type customNVerifyNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNVerifyNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomNVerifyNewResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    customNVerifyNewResponseEnvelopeMessagesJSON `json:"-"`
}

// customNVerifyNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CustomNVerifyNewResponseEnvelopeMessages]
type customNVerifyNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNVerifyNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CustomNVerifyNewResponseEnvelopeSuccess bool

const (
	CustomNVerifyNewResponseEnvelopeSuccessTrue CustomNVerifyNewResponseEnvelopeSuccess = true
)

type CustomNVerifyNewResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                        `json:"total_count"`
	JSON       customNVerifyNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// customNVerifyNewResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [CustomNVerifyNewResponseEnvelopeResultInfo]
type customNVerifyNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNVerifyNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
