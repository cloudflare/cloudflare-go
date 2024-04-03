// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_nameservers

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// CustomNameserverService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCustomNameserverService] method
// instead.
type CustomNameserverService struct {
	Options []option.RequestOption
}

// NewCustomNameserverService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCustomNameserverService(opts ...option.RequestOption) (r *CustomNameserverService) {
	r = &CustomNameserverService{}
	r.Options = opts
	return
}

// Add Account Custom Nameserver
func (r *CustomNameserverService) New(ctx context.Context, params CustomNameserverNewParams, opts ...option.RequestOption) (res *CustomNameserver, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomNameserverNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/custom_ns", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete Account Custom Nameserver
func (r *CustomNameserverService) Delete(ctx context.Context, customNSID string, params CustomNameserverDeleteParams, opts ...option.RequestOption) (res *CustomNameserverDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomNameserverDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/custom_ns/%s", params.AccountID, customNSID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Eligible Zones for Account Custom Nameservers
func (r *CustomNameserverService) Availabilty(ctx context.Context, query CustomNameserverAvailabiltyParams, opts ...option.RequestOption) (res *[]string, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomNameserverAvailabiltyResponseEnvelope
	path := fmt.Sprintf("accounts/%s/custom_ns/availability", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List an account's custom nameservers.
func (r *CustomNameserverService) Get(ctx context.Context, query CustomNameserverGetParams, opts ...option.RequestOption) (res *[]CustomNameserver, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomNameserverGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/custom_ns", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Verify Account Custom Nameserver Glue Records
func (r *CustomNameserverService) Verify(ctx context.Context, params CustomNameserverVerifyParams, opts ...option.RequestOption) (res *[]CustomNameserver, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomNameserverVerifyResponseEnvelope
	path := fmt.Sprintf("accounts/%s/custom_ns/verify", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A single account custom nameserver.
type CustomNameserver struct {
	// A and AAAA records associated with the nameserver.
	DNSRecords []CustomNameserverDNSRecord `json:"dns_records,required"`
	// The FQDN of the name server.
	NSName string `json:"ns_name,required" format:"hostname"`
	// Verification status of the nameserver.
	Status CustomNameserverStatus `json:"status,required"`
	// Identifier
	ZoneTag string `json:"zone_tag,required"`
	// The number of the set that this name server belongs to.
	NSSet float64              `json:"ns_set"`
	JSON  customNameserverJSON `json:"-"`
}

// customNameserverJSON contains the JSON metadata for the struct
// [CustomNameserver]
type customNameserverJSON struct {
	DNSRecords  apijson.Field
	NSName      apijson.Field
	Status      apijson.Field
	ZoneTag     apijson.Field
	NSSet       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNameserver) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverJSON) RawJSON() string {
	return r.raw
}

type CustomNameserverDNSRecord struct {
	// DNS record type.
	Type CustomNameserverDNSRecordsType `json:"type"`
	// DNS record contents (an IPv4 or IPv6 address).
	Value string                        `json:"value"`
	JSON  customNameserverDNSRecordJSON `json:"-"`
}

// customNameserverDNSRecordJSON contains the JSON metadata for the struct
// [CustomNameserverDNSRecord]
type customNameserverDNSRecordJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNameserverDNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverDNSRecordJSON) RawJSON() string {
	return r.raw
}

// DNS record type.
type CustomNameserverDNSRecordsType string

const (
	CustomNameserverDNSRecordsTypeA    CustomNameserverDNSRecordsType = "A"
	CustomNameserverDNSRecordsTypeAAAA CustomNameserverDNSRecordsType = "AAAA"
)

func (r CustomNameserverDNSRecordsType) IsKnown() bool {
	switch r {
	case CustomNameserverDNSRecordsTypeA, CustomNameserverDNSRecordsTypeAAAA:
		return true
	}
	return false
}

// Verification status of the nameserver.
type CustomNameserverStatus string

const (
	CustomNameserverStatusMoved    CustomNameserverStatus = "moved"
	CustomNameserverStatusPending  CustomNameserverStatus = "pending"
	CustomNameserverStatusVerified CustomNameserverStatus = "verified"
)

func (r CustomNameserverStatus) IsKnown() bool {
	switch r {
	case CustomNameserverStatusMoved, CustomNameserverStatusPending, CustomNameserverStatusVerified:
		return true
	}
	return false
}

// Union satisfied by [custom_nameservers.CustomNameserverDeleteResponseUnknown],
// [custom_nameservers.CustomNameserverDeleteResponseArray] or
// [shared.UnionString].
type CustomNameserverDeleteResponse interface {
	ImplementsCustomNameserversCustomNameserverDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomNameserverDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CustomNameserverDeleteResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CustomNameserverDeleteResponseArray []interface{}

func (r CustomNameserverDeleteResponseArray) ImplementsCustomNameserversCustomNameserverDeleteResponse() {
}

type CustomNameserverNewParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// The FQDN of the name server.
	NSName param.Field[string] `json:"ns_name,required" format:"hostname"`
	// The number of the set that this name server belongs to.
	NSSet param.Field[float64] `json:"ns_set"`
}

func (r CustomNameserverNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomNameserverNewResponseEnvelope struct {
	Errors   []CustomNameserverNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomNameserverNewResponseEnvelopeMessages `json:"messages,required"`
	// A single account custom nameserver.
	Result CustomNameserver `json:"result,required"`
	// Whether the API call was successful
	Success CustomNameserverNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    customNameserverNewResponseEnvelopeJSON    `json:"-"`
}

// customNameserverNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [CustomNameserverNewResponseEnvelope]
type customNameserverNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNameserverNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CustomNameserverNewResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    customNameserverNewResponseEnvelopeErrorsJSON `json:"-"`
}

// customNameserverNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CustomNameserverNewResponseEnvelopeErrors]
type customNameserverNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNameserverNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CustomNameserverNewResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    customNameserverNewResponseEnvelopeMessagesJSON `json:"-"`
}

// customNameserverNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CustomNameserverNewResponseEnvelopeMessages]
type customNameserverNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNameserverNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CustomNameserverNewResponseEnvelopeSuccess bool

const (
	CustomNameserverNewResponseEnvelopeSuccessTrue CustomNameserverNewResponseEnvelopeSuccess = true
)

func (r CustomNameserverNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CustomNameserverNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CustomNameserverDeleteParams struct {
	// Account identifier tag.
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r CustomNameserverDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type CustomNameserverDeleteResponseEnvelope struct {
	Errors   []CustomNameserverDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomNameserverDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   CustomNameserverDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    CustomNameserverDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo CustomNameserverDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       customNameserverDeleteResponseEnvelopeJSON       `json:"-"`
}

// customNameserverDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [CustomNameserverDeleteResponseEnvelope]
type customNameserverDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNameserverDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CustomNameserverDeleteResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    customNameserverDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// customNameserverDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CustomNameserverDeleteResponseEnvelopeErrors]
type customNameserverDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNameserverDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CustomNameserverDeleteResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    customNameserverDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// customNameserverDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CustomNameserverDeleteResponseEnvelopeMessages]
type customNameserverDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNameserverDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CustomNameserverDeleteResponseEnvelopeSuccess bool

const (
	CustomNameserverDeleteResponseEnvelopeSuccessTrue CustomNameserverDeleteResponseEnvelopeSuccess = true
)

func (r CustomNameserverDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CustomNameserverDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CustomNameserverDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                              `json:"total_count"`
	JSON       customNameserverDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// customNameserverDeleteResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [CustomNameserverDeleteResponseEnvelopeResultInfo]
type customNameserverDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNameserverDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type CustomNameserverAvailabiltyParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type CustomNameserverAvailabiltyResponseEnvelope struct {
	Errors   []CustomNameserverAvailabiltyResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomNameserverAvailabiltyResponseEnvelopeMessages `json:"messages,required"`
	Result   []string                                              `json:"result,required,nullable" format:"hostname"`
	// Whether the API call was successful
	Success    CustomNameserverAvailabiltyResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo CustomNameserverAvailabiltyResponseEnvelopeResultInfo `json:"result_info"`
	JSON       customNameserverAvailabiltyResponseEnvelopeJSON       `json:"-"`
}

// customNameserverAvailabiltyResponseEnvelopeJSON contains the JSON metadata for
// the struct [CustomNameserverAvailabiltyResponseEnvelope]
type customNameserverAvailabiltyResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNameserverAvailabiltyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverAvailabiltyResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CustomNameserverAvailabiltyResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    customNameserverAvailabiltyResponseEnvelopeErrorsJSON `json:"-"`
}

// customNameserverAvailabiltyResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [CustomNameserverAvailabiltyResponseEnvelopeErrors]
type customNameserverAvailabiltyResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNameserverAvailabiltyResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverAvailabiltyResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CustomNameserverAvailabiltyResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    customNameserverAvailabiltyResponseEnvelopeMessagesJSON `json:"-"`
}

// customNameserverAvailabiltyResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [CustomNameserverAvailabiltyResponseEnvelopeMessages]
type customNameserverAvailabiltyResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNameserverAvailabiltyResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverAvailabiltyResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CustomNameserverAvailabiltyResponseEnvelopeSuccess bool

const (
	CustomNameserverAvailabiltyResponseEnvelopeSuccessTrue CustomNameserverAvailabiltyResponseEnvelopeSuccess = true
)

func (r CustomNameserverAvailabiltyResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CustomNameserverAvailabiltyResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CustomNameserverAvailabiltyResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                   `json:"total_count"`
	JSON       customNameserverAvailabiltyResponseEnvelopeResultInfoJSON `json:"-"`
}

// customNameserverAvailabiltyResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [CustomNameserverAvailabiltyResponseEnvelopeResultInfo]
type customNameserverAvailabiltyResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNameserverAvailabiltyResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverAvailabiltyResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type CustomNameserverGetParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type CustomNameserverGetResponseEnvelope struct {
	Errors   []CustomNameserverGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomNameserverGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []CustomNameserver                            `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    CustomNameserverGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo CustomNameserverGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       customNameserverGetResponseEnvelopeJSON       `json:"-"`
}

// customNameserverGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [CustomNameserverGetResponseEnvelope]
type customNameserverGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNameserverGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CustomNameserverGetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    customNameserverGetResponseEnvelopeErrorsJSON `json:"-"`
}

// customNameserverGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CustomNameserverGetResponseEnvelopeErrors]
type customNameserverGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNameserverGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CustomNameserverGetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    customNameserverGetResponseEnvelopeMessagesJSON `json:"-"`
}

// customNameserverGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CustomNameserverGetResponseEnvelopeMessages]
type customNameserverGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNameserverGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CustomNameserverGetResponseEnvelopeSuccess bool

const (
	CustomNameserverGetResponseEnvelopeSuccessTrue CustomNameserverGetResponseEnvelopeSuccess = true
)

func (r CustomNameserverGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CustomNameserverGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CustomNameserverGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                           `json:"total_count"`
	JSON       customNameserverGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// customNameserverGetResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [CustomNameserverGetResponseEnvelopeResultInfo]
type customNameserverGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNameserverGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type CustomNameserverVerifyParams struct {
	// Account identifier tag.
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r CustomNameserverVerifyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type CustomNameserverVerifyResponseEnvelope struct {
	Errors   []CustomNameserverVerifyResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomNameserverVerifyResponseEnvelopeMessages `json:"messages,required"`
	Result   []CustomNameserver                               `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    CustomNameserverVerifyResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo CustomNameserverVerifyResponseEnvelopeResultInfo `json:"result_info"`
	JSON       customNameserverVerifyResponseEnvelopeJSON       `json:"-"`
}

// customNameserverVerifyResponseEnvelopeJSON contains the JSON metadata for the
// struct [CustomNameserverVerifyResponseEnvelope]
type customNameserverVerifyResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNameserverVerifyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverVerifyResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CustomNameserverVerifyResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    customNameserverVerifyResponseEnvelopeErrorsJSON `json:"-"`
}

// customNameserverVerifyResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CustomNameserverVerifyResponseEnvelopeErrors]
type customNameserverVerifyResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNameserverVerifyResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverVerifyResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CustomNameserverVerifyResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    customNameserverVerifyResponseEnvelopeMessagesJSON `json:"-"`
}

// customNameserverVerifyResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CustomNameserverVerifyResponseEnvelopeMessages]
type customNameserverVerifyResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNameserverVerifyResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverVerifyResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CustomNameserverVerifyResponseEnvelopeSuccess bool

const (
	CustomNameserverVerifyResponseEnvelopeSuccessTrue CustomNameserverVerifyResponseEnvelopeSuccess = true
)

func (r CustomNameserverVerifyResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CustomNameserverVerifyResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CustomNameserverVerifyResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                              `json:"total_count"`
	JSON       customNameserverVerifyResponseEnvelopeResultInfoJSON `json:"-"`
}

// customNameserverVerifyResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [CustomNameserverVerifyResponseEnvelopeResultInfo]
type customNameserverVerifyResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNameserverVerifyResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverVerifyResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
