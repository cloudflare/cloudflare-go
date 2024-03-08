// File generated from our OpenAPI spec by Stainless.

package cloudforce_one

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
)

// RequestService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewRequestService] method instead.
type RequestService struct {
	Options  []option.RequestOption
	Message  *RequestMessageService
	Priority *RequestPriorityService
}

// NewRequestService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRequestService(opts ...option.RequestOption) (r *RequestService) {
	r = &RequestService{}
	r.Options = opts
	r.Message = NewRequestMessageService(opts...)
	r.Priority = NewRequestPriorityService(opts...)
	return
}

// Creating a request adds the request into the Cloudforce One queue for analysis.
// In addition to the content, a short title, type, priority, and releasability
// should be provided. If one is not provided a default will be assigned.
func (r *RequestService) New(ctx context.Context, accountIdentifier string, body RequestNewParams, opts ...option.RequestOption) (res *RequestNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RequestNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/new", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updating a request alters the request in the Cloudforce One queue. This API may
// be used to update any attributes of the request after the initial submission.
// Only fields that you choose to update need to be add to the request body
func (r *RequestService) Update(ctx context.Context, accountIdentifier string, requestIdentifier string, body RequestUpdateParams, opts ...option.RequestOption) (res *RequestUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RequestUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s", accountIdentifier, requestIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Requests
func (r *RequestService) List(ctx context.Context, accountIdentifier string, body RequestListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[RequestListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests", accountIdentifier)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, body, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Requests
func (r *RequestService) ListAutoPaging(ctx context.Context, accountIdentifier string, body RequestListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[RequestListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, accountIdentifier, body, opts...))
}

// Delete a Request
func (r *RequestService) Delete(ctx context.Context, accountIdentifier string, requestIdentifier string, opts ...option.RequestOption) (res *RequestDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RequestDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s", accountIdentifier, requestIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Request Priority, Status, and TLP constants
func (r *RequestService) Constants(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *RequestConstantsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RequestConstantsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/constants", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a Request
func (r *RequestService) Get(ctx context.Context, accountIdentifier string, requestIdentifier string, opts ...option.RequestOption) (res *RequestGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RequestGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s", accountIdentifier, requestIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Request Quota
func (r *RequestService) Quota(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *RequestQuotaResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RequestQuotaResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/quota", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Request Types
func (r *RequestService) Types(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *[]string, err error) {
	opts = append(r.Options[:], opts...)
	var env RequestTypesResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/types", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RequestNewResponse struct {
	// UUID
	ID string `json:"id,required"`
	// Request content
	Content  string    `json:"content,required"`
	Created  time.Time `json:"created,required" format:"date-time"`
	Priority time.Time `json:"priority,required" format:"date-time"`
	// Requested information from request
	Request string `json:"request,required"`
	// Brief description of the request
	Summary string `json:"summary,required"`
	// The CISA defined Traffic Light Protocol (TLP)
	Tlp       RequestNewResponseTlp `json:"tlp,required"`
	Updated   time.Time             `json:"updated,required" format:"date-time"`
	Completed time.Time             `json:"completed" format:"date-time"`
	// Tokens for the request messages
	MessageTokens int64 `json:"message_tokens"`
	// Readable Request ID
	ReadableID string `json:"readable_id"`
	// Request Status
	Status RequestNewResponseStatus `json:"status"`
	// Tokens for the request
	Tokens int64                  `json:"tokens"`
	JSON   requestNewResponseJSON `json:"-"`
}

// requestNewResponseJSON contains the JSON metadata for the struct
// [RequestNewResponse]
type requestNewResponseJSON struct {
	ID            apijson.Field
	Content       apijson.Field
	Created       apijson.Field
	Priority      apijson.Field
	Request       apijson.Field
	Summary       apijson.Field
	Tlp           apijson.Field
	Updated       apijson.Field
	Completed     apijson.Field
	MessageTokens apijson.Field
	ReadableID    apijson.Field
	Status        apijson.Field
	Tokens        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RequestNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestNewResponseJSON) RawJSON() string {
	return r.raw
}

// The CISA defined Traffic Light Protocol (TLP)
type RequestNewResponseTlp string

const (
	RequestNewResponseTlpClear       RequestNewResponseTlp = "clear"
	RequestNewResponseTlpAmber       RequestNewResponseTlp = "amber"
	RequestNewResponseTlpAmberStrict RequestNewResponseTlp = "amber-strict"
	RequestNewResponseTlpGreen       RequestNewResponseTlp = "green"
	RequestNewResponseTlpRed         RequestNewResponseTlp = "red"
)

// Request Status
type RequestNewResponseStatus string

const (
	RequestNewResponseStatusOpen      RequestNewResponseStatus = "open"
	RequestNewResponseStatusAccepted  RequestNewResponseStatus = "accepted"
	RequestNewResponseStatusReported  RequestNewResponseStatus = "reported"
	RequestNewResponseStatusApproved  RequestNewResponseStatus = "approved"
	RequestNewResponseStatusCompleted RequestNewResponseStatus = "completed"
	RequestNewResponseStatusDeclined  RequestNewResponseStatus = "declined"
)

type RequestUpdateResponse struct {
	// UUID
	ID string `json:"id,required"`
	// Request content
	Content  string    `json:"content,required"`
	Created  time.Time `json:"created,required" format:"date-time"`
	Priority time.Time `json:"priority,required" format:"date-time"`
	// Requested information from request
	Request string `json:"request,required"`
	// Brief description of the request
	Summary string `json:"summary,required"`
	// The CISA defined Traffic Light Protocol (TLP)
	Tlp       RequestUpdateResponseTlp `json:"tlp,required"`
	Updated   time.Time                `json:"updated,required" format:"date-time"`
	Completed time.Time                `json:"completed" format:"date-time"`
	// Tokens for the request messages
	MessageTokens int64 `json:"message_tokens"`
	// Readable Request ID
	ReadableID string `json:"readable_id"`
	// Request Status
	Status RequestUpdateResponseStatus `json:"status"`
	// Tokens for the request
	Tokens int64                     `json:"tokens"`
	JSON   requestUpdateResponseJSON `json:"-"`
}

// requestUpdateResponseJSON contains the JSON metadata for the struct
// [RequestUpdateResponse]
type requestUpdateResponseJSON struct {
	ID            apijson.Field
	Content       apijson.Field
	Created       apijson.Field
	Priority      apijson.Field
	Request       apijson.Field
	Summary       apijson.Field
	Tlp           apijson.Field
	Updated       apijson.Field
	Completed     apijson.Field
	MessageTokens apijson.Field
	ReadableID    apijson.Field
	Status        apijson.Field
	Tokens        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RequestUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The CISA defined Traffic Light Protocol (TLP)
type RequestUpdateResponseTlp string

const (
	RequestUpdateResponseTlpClear       RequestUpdateResponseTlp = "clear"
	RequestUpdateResponseTlpAmber       RequestUpdateResponseTlp = "amber"
	RequestUpdateResponseTlpAmberStrict RequestUpdateResponseTlp = "amber-strict"
	RequestUpdateResponseTlpGreen       RequestUpdateResponseTlp = "green"
	RequestUpdateResponseTlpRed         RequestUpdateResponseTlp = "red"
)

// Request Status
type RequestUpdateResponseStatus string

const (
	RequestUpdateResponseStatusOpen      RequestUpdateResponseStatus = "open"
	RequestUpdateResponseStatusAccepted  RequestUpdateResponseStatus = "accepted"
	RequestUpdateResponseStatusReported  RequestUpdateResponseStatus = "reported"
	RequestUpdateResponseStatusApproved  RequestUpdateResponseStatus = "approved"
	RequestUpdateResponseStatusCompleted RequestUpdateResponseStatus = "completed"
	RequestUpdateResponseStatusDeclined  RequestUpdateResponseStatus = "declined"
)

type RequestListResponse struct {
	// UUID
	ID string `json:"id,required"`
	// Request creation time
	Created  time.Time                   `json:"created,required" format:"date-time"`
	Priority RequestListResponsePriority `json:"priority,required"`
	// Requested information from request
	Request string `json:"request,required"`
	// Brief description of the request
	Summary string `json:"summary,required"`
	// The CISA defined Traffic Light Protocol (TLP)
	Tlp RequestListResponseTlp `json:"tlp,required"`
	// Request last updated time
	Updated time.Time `json:"updated,required" format:"date-time"`
	// Request completion time
	Completed time.Time `json:"completed" format:"date-time"`
	// Tokens for the request messages
	MessageTokens int64 `json:"message_tokens"`
	// Readable Request ID
	ReadableID string `json:"readable_id"`
	// Request Status
	Status RequestListResponseStatus `json:"status"`
	// Tokens for the request
	Tokens int64                   `json:"tokens"`
	JSON   requestListResponseJSON `json:"-"`
}

// requestListResponseJSON contains the JSON metadata for the struct
// [RequestListResponse]
type requestListResponseJSON struct {
	ID            apijson.Field
	Created       apijson.Field
	Priority      apijson.Field
	Request       apijson.Field
	Summary       apijson.Field
	Tlp           apijson.Field
	Updated       apijson.Field
	Completed     apijson.Field
	MessageTokens apijson.Field
	ReadableID    apijson.Field
	Status        apijson.Field
	Tokens        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RequestListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestListResponseJSON) RawJSON() string {
	return r.raw
}

type RequestListResponsePriority string

const (
	RequestListResponsePriorityRoutine RequestListResponsePriority = "routine"
	RequestListResponsePriorityHigh    RequestListResponsePriority = "high"
	RequestListResponsePriorityUrgent  RequestListResponsePriority = "urgent"
)

// The CISA defined Traffic Light Protocol (TLP)
type RequestListResponseTlp string

const (
	RequestListResponseTlpClear       RequestListResponseTlp = "clear"
	RequestListResponseTlpAmber       RequestListResponseTlp = "amber"
	RequestListResponseTlpAmberStrict RequestListResponseTlp = "amber-strict"
	RequestListResponseTlpGreen       RequestListResponseTlp = "green"
	RequestListResponseTlpRed         RequestListResponseTlp = "red"
)

// Request Status
type RequestListResponseStatus string

const (
	RequestListResponseStatusOpen      RequestListResponseStatus = "open"
	RequestListResponseStatusAccepted  RequestListResponseStatus = "accepted"
	RequestListResponseStatusReported  RequestListResponseStatus = "reported"
	RequestListResponseStatusApproved  RequestListResponseStatus = "approved"
	RequestListResponseStatusCompleted RequestListResponseStatus = "completed"
	RequestListResponseStatusDeclined  RequestListResponseStatus = "declined"
)

// Union satisfied by [cloudforce_one.RequestDeleteResponseUnknown],
// [cloudforce_one.RequestDeleteResponseArray] or [shared.UnionString].
type RequestDeleteResponse interface {
	ImplementsCloudforceOneRequestDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RequestDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RequestDeleteResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type RequestDeleteResponseArray []interface{}

func (r RequestDeleteResponseArray) ImplementsCloudforceOneRequestDeleteResponse() {}

type RequestConstantsResponse struct {
	Priority []RequestConstantsResponsePriority `json:"priority"`
	Status   []RequestConstantsResponseStatus   `json:"status"`
	Tlp      []RequestConstantsResponseTlp      `json:"tlp"`
	JSON     requestConstantsResponseJSON       `json:"-"`
}

// requestConstantsResponseJSON contains the JSON metadata for the struct
// [RequestConstantsResponse]
type requestConstantsResponseJSON struct {
	Priority    apijson.Field
	Status      apijson.Field
	Tlp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestConstantsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestConstantsResponseJSON) RawJSON() string {
	return r.raw
}

type RequestConstantsResponsePriority string

const (
	RequestConstantsResponsePriorityRoutine RequestConstantsResponsePriority = "routine"
	RequestConstantsResponsePriorityHigh    RequestConstantsResponsePriority = "high"
	RequestConstantsResponsePriorityUrgent  RequestConstantsResponsePriority = "urgent"
)

// Request Status
type RequestConstantsResponseStatus string

const (
	RequestConstantsResponseStatusOpen      RequestConstantsResponseStatus = "open"
	RequestConstantsResponseStatusAccepted  RequestConstantsResponseStatus = "accepted"
	RequestConstantsResponseStatusReported  RequestConstantsResponseStatus = "reported"
	RequestConstantsResponseStatusApproved  RequestConstantsResponseStatus = "approved"
	RequestConstantsResponseStatusCompleted RequestConstantsResponseStatus = "completed"
	RequestConstantsResponseStatusDeclined  RequestConstantsResponseStatus = "declined"
)

// The CISA defined Traffic Light Protocol (TLP)
type RequestConstantsResponseTlp string

const (
	RequestConstantsResponseTlpClear       RequestConstantsResponseTlp = "clear"
	RequestConstantsResponseTlpAmber       RequestConstantsResponseTlp = "amber"
	RequestConstantsResponseTlpAmberStrict RequestConstantsResponseTlp = "amber-strict"
	RequestConstantsResponseTlpGreen       RequestConstantsResponseTlp = "green"
	RequestConstantsResponseTlpRed         RequestConstantsResponseTlp = "red"
)

type RequestGetResponse struct {
	// UUID
	ID string `json:"id,required"`
	// Request content
	Content  string    `json:"content,required"`
	Created  time.Time `json:"created,required" format:"date-time"`
	Priority time.Time `json:"priority,required" format:"date-time"`
	// Requested information from request
	Request string `json:"request,required"`
	// Brief description of the request
	Summary string `json:"summary,required"`
	// The CISA defined Traffic Light Protocol (TLP)
	Tlp       RequestGetResponseTlp `json:"tlp,required"`
	Updated   time.Time             `json:"updated,required" format:"date-time"`
	Completed time.Time             `json:"completed" format:"date-time"`
	// Tokens for the request messages
	MessageTokens int64 `json:"message_tokens"`
	// Readable Request ID
	ReadableID string `json:"readable_id"`
	// Request Status
	Status RequestGetResponseStatus `json:"status"`
	// Tokens for the request
	Tokens int64                  `json:"tokens"`
	JSON   requestGetResponseJSON `json:"-"`
}

// requestGetResponseJSON contains the JSON metadata for the struct
// [RequestGetResponse]
type requestGetResponseJSON struct {
	ID            apijson.Field
	Content       apijson.Field
	Created       apijson.Field
	Priority      apijson.Field
	Request       apijson.Field
	Summary       apijson.Field
	Tlp           apijson.Field
	Updated       apijson.Field
	Completed     apijson.Field
	MessageTokens apijson.Field
	ReadableID    apijson.Field
	Status        apijson.Field
	Tokens        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RequestGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestGetResponseJSON) RawJSON() string {
	return r.raw
}

// The CISA defined Traffic Light Protocol (TLP)
type RequestGetResponseTlp string

const (
	RequestGetResponseTlpClear       RequestGetResponseTlp = "clear"
	RequestGetResponseTlpAmber       RequestGetResponseTlp = "amber"
	RequestGetResponseTlpAmberStrict RequestGetResponseTlp = "amber-strict"
	RequestGetResponseTlpGreen       RequestGetResponseTlp = "green"
	RequestGetResponseTlpRed         RequestGetResponseTlp = "red"
)

// Request Status
type RequestGetResponseStatus string

const (
	RequestGetResponseStatusOpen      RequestGetResponseStatus = "open"
	RequestGetResponseStatusAccepted  RequestGetResponseStatus = "accepted"
	RequestGetResponseStatusReported  RequestGetResponseStatus = "reported"
	RequestGetResponseStatusApproved  RequestGetResponseStatus = "approved"
	RequestGetResponseStatusCompleted RequestGetResponseStatus = "completed"
	RequestGetResponseStatusDeclined  RequestGetResponseStatus = "declined"
)

type RequestQuotaResponse struct {
	// Anniversary date is when annual quota limit is refresh
	AnniversaryDate time.Time `json:"anniversary_date" format:"date-time"`
	// Quater anniversary date is when quota limit is refreshed each quarter
	QuarterAnniversaryDate time.Time `json:"quarter_anniversary_date" format:"date-time"`
	// Tokens for the quarter
	Quota int64 `json:"quota"`
	// Tokens remaining for the quarter
	Remaining int64                    `json:"remaining"`
	JSON      requestQuotaResponseJSON `json:"-"`
}

// requestQuotaResponseJSON contains the JSON metadata for the struct
// [RequestQuotaResponse]
type requestQuotaResponseJSON struct {
	AnniversaryDate        apijson.Field
	QuarterAnniversaryDate apijson.Field
	Quota                  apijson.Field
	Remaining              apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *RequestQuotaResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestQuotaResponseJSON) RawJSON() string {
	return r.raw
}

type RequestNewParams struct {
	// Request content
	Content param.Field[string] `json:"content"`
	// Priority for analyzing the request
	Priority param.Field[string] `json:"priority"`
	// Requested information from request
	RequestType param.Field[string] `json:"request_type"`
	// Brief description of the request
	Summary param.Field[string] `json:"summary"`
	// The CISA defined Traffic Light Protocol (TLP)
	Tlp param.Field[RequestNewParamsTlp] `json:"tlp"`
}

func (r RequestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The CISA defined Traffic Light Protocol (TLP)
type RequestNewParamsTlp string

const (
	RequestNewParamsTlpClear       RequestNewParamsTlp = "clear"
	RequestNewParamsTlpAmber       RequestNewParamsTlp = "amber"
	RequestNewParamsTlpAmberStrict RequestNewParamsTlp = "amber-strict"
	RequestNewParamsTlpGreen       RequestNewParamsTlp = "green"
	RequestNewParamsTlpRed         RequestNewParamsTlp = "red"
)

type RequestNewResponseEnvelope struct {
	Errors   []RequestNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RequestNewResponseEnvelopeMessages `json:"messages,required"`
	Result   RequestNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success RequestNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    requestNewResponseEnvelopeJSON    `json:"-"`
}

// requestNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestNewResponseEnvelope]
type requestNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RequestNewResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    requestNewResponseEnvelopeErrorsJSON `json:"-"`
}

// requestNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RequestNewResponseEnvelopeErrors]
type requestNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RequestNewResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    requestNewResponseEnvelopeMessagesJSON `json:"-"`
}

// requestNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RequestNewResponseEnvelopeMessages]
type requestNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestNewResponseEnvelopeSuccess bool

const (
	RequestNewResponseEnvelopeSuccessTrue RequestNewResponseEnvelopeSuccess = true
)

type RequestUpdateParams struct {
	// Request content
	Content param.Field[string] `json:"content"`
	// Priority for analyzing the request
	Priority param.Field[string] `json:"priority"`
	// Requested information from request
	RequestType param.Field[string] `json:"request_type"`
	// Brief description of the request
	Summary param.Field[string] `json:"summary"`
	// The CISA defined Traffic Light Protocol (TLP)
	Tlp param.Field[RequestUpdateParamsTlp] `json:"tlp"`
}

func (r RequestUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The CISA defined Traffic Light Protocol (TLP)
type RequestUpdateParamsTlp string

const (
	RequestUpdateParamsTlpClear       RequestUpdateParamsTlp = "clear"
	RequestUpdateParamsTlpAmber       RequestUpdateParamsTlp = "amber"
	RequestUpdateParamsTlpAmberStrict RequestUpdateParamsTlp = "amber-strict"
	RequestUpdateParamsTlpGreen       RequestUpdateParamsTlp = "green"
	RequestUpdateParamsTlpRed         RequestUpdateParamsTlp = "red"
)

type RequestUpdateResponseEnvelope struct {
	Errors   []RequestUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RequestUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   RequestUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success RequestUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    requestUpdateResponseEnvelopeJSON    `json:"-"`
}

// requestUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestUpdateResponseEnvelope]
type requestUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RequestUpdateResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    requestUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// requestUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RequestUpdateResponseEnvelopeErrors]
type requestUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RequestUpdateResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    requestUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// requestUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RequestUpdateResponseEnvelopeMessages]
type requestUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestUpdateResponseEnvelopeSuccess bool

const (
	RequestUpdateResponseEnvelopeSuccessTrue RequestUpdateResponseEnvelopeSuccess = true
)

type RequestListParams struct {
	// Page number of results
	Page param.Field[int64] `json:"page,required"`
	// Number of results per page
	PerPage param.Field[int64] `json:"per_page,required"`
	// Retrieve requests completed after this time
	CompletedAfter param.Field[time.Time] `json:"completed_after" format:"date-time"`
	// Retrieve requests completed before this time
	CompletedBefore param.Field[time.Time] `json:"completed_before" format:"date-time"`
	// Retrieve requests created after this time
	CreatedAfter param.Field[time.Time] `json:"created_after" format:"date-time"`
	// Retrieve requests created before this time
	CreatedBefore param.Field[time.Time] `json:"created_before" format:"date-time"`
	// Requested information from request
	RequestType param.Field[string] `json:"request_type"`
	// Field to sort results by
	SortBy param.Field[string] `json:"sort_by"`
	// Sort order (asc or desc)
	SortOrder param.Field[RequestListParamsSortOrder] `json:"sort_order"`
	// Request Status
	Status param.Field[RequestListParamsStatus] `json:"status"`
}

func (r RequestListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Sort order (asc or desc)
type RequestListParamsSortOrder string

const (
	RequestListParamsSortOrderAsc  RequestListParamsSortOrder = "asc"
	RequestListParamsSortOrderDesc RequestListParamsSortOrder = "desc"
)

// Request Status
type RequestListParamsStatus string

const (
	RequestListParamsStatusOpen      RequestListParamsStatus = "open"
	RequestListParamsStatusAccepted  RequestListParamsStatus = "accepted"
	RequestListParamsStatusReported  RequestListParamsStatus = "reported"
	RequestListParamsStatusApproved  RequestListParamsStatus = "approved"
	RequestListParamsStatusCompleted RequestListParamsStatus = "completed"
	RequestListParamsStatusDeclined  RequestListParamsStatus = "declined"
)

type RequestDeleteResponseEnvelope struct {
	Errors   []RequestDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RequestDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   RequestDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success RequestDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    requestDeleteResponseEnvelopeJSON    `json:"-"`
}

// requestDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestDeleteResponseEnvelope]
type requestDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RequestDeleteResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    requestDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// requestDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RequestDeleteResponseEnvelopeErrors]
type requestDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RequestDeleteResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    requestDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// requestDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RequestDeleteResponseEnvelopeMessages]
type requestDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestDeleteResponseEnvelopeSuccess bool

const (
	RequestDeleteResponseEnvelopeSuccessTrue RequestDeleteResponseEnvelopeSuccess = true
)

type RequestConstantsResponseEnvelope struct {
	Errors   []RequestConstantsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RequestConstantsResponseEnvelopeMessages `json:"messages,required"`
	Result   RequestConstantsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success RequestConstantsResponseEnvelopeSuccess `json:"success,required"`
	JSON    requestConstantsResponseEnvelopeJSON    `json:"-"`
}

// requestConstantsResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestConstantsResponseEnvelope]
type requestConstantsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestConstantsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestConstantsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RequestConstantsResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    requestConstantsResponseEnvelopeErrorsJSON `json:"-"`
}

// requestConstantsResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RequestConstantsResponseEnvelopeErrors]
type requestConstantsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestConstantsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestConstantsResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RequestConstantsResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    requestConstantsResponseEnvelopeMessagesJSON `json:"-"`
}

// requestConstantsResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RequestConstantsResponseEnvelopeMessages]
type requestConstantsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestConstantsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestConstantsResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestConstantsResponseEnvelopeSuccess bool

const (
	RequestConstantsResponseEnvelopeSuccessTrue RequestConstantsResponseEnvelopeSuccess = true
)

type RequestGetResponseEnvelope struct {
	Errors   []RequestGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RequestGetResponseEnvelopeMessages `json:"messages,required"`
	Result   RequestGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success RequestGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    requestGetResponseEnvelopeJSON    `json:"-"`
}

// requestGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestGetResponseEnvelope]
type requestGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RequestGetResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    requestGetResponseEnvelopeErrorsJSON `json:"-"`
}

// requestGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RequestGetResponseEnvelopeErrors]
type requestGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RequestGetResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    requestGetResponseEnvelopeMessagesJSON `json:"-"`
}

// requestGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RequestGetResponseEnvelopeMessages]
type requestGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestGetResponseEnvelopeSuccess bool

const (
	RequestGetResponseEnvelopeSuccessTrue RequestGetResponseEnvelopeSuccess = true
)

type RequestQuotaResponseEnvelope struct {
	Errors   []RequestQuotaResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RequestQuotaResponseEnvelopeMessages `json:"messages,required"`
	Result   RequestQuotaResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success RequestQuotaResponseEnvelopeSuccess `json:"success,required"`
	JSON    requestQuotaResponseEnvelopeJSON    `json:"-"`
}

// requestQuotaResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestQuotaResponseEnvelope]
type requestQuotaResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestQuotaResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestQuotaResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RequestQuotaResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    requestQuotaResponseEnvelopeErrorsJSON `json:"-"`
}

// requestQuotaResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RequestQuotaResponseEnvelopeErrors]
type requestQuotaResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestQuotaResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestQuotaResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RequestQuotaResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    requestQuotaResponseEnvelopeMessagesJSON `json:"-"`
}

// requestQuotaResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RequestQuotaResponseEnvelopeMessages]
type requestQuotaResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestQuotaResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestQuotaResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestQuotaResponseEnvelopeSuccess bool

const (
	RequestQuotaResponseEnvelopeSuccessTrue RequestQuotaResponseEnvelopeSuccess = true
)

type RequestTypesResponseEnvelope struct {
	Errors   []RequestTypesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RequestTypesResponseEnvelopeMessages `json:"messages,required"`
	Result   []string                               `json:"result,required"`
	// Whether the API call was successful
	Success RequestTypesResponseEnvelopeSuccess `json:"success,required"`
	JSON    requestTypesResponseEnvelopeJSON    `json:"-"`
}

// requestTypesResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestTypesResponseEnvelope]
type requestTypesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestTypesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestTypesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RequestTypesResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    requestTypesResponseEnvelopeErrorsJSON `json:"-"`
}

// requestTypesResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RequestTypesResponseEnvelopeErrors]
type requestTypesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestTypesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestTypesResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RequestTypesResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    requestTypesResponseEnvelopeMessagesJSON `json:"-"`
}

// requestTypesResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RequestTypesResponseEnvelopeMessages]
type requestTypesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestTypesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestTypesResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestTypesResponseEnvelopeSuccess bool

const (
	RequestTypesResponseEnvelopeSuccessTrue RequestTypesResponseEnvelopeSuccess = true
)
