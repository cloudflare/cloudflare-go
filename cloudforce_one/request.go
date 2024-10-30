// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v3/shared"
)

// RequestService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRequestService] method instead.
type RequestService struct {
	Options  []option.RequestOption
	Message  *RequestMessageService
	Priority *RequestPriorityService
	Assets   *RequestAssetService
}

// NewRequestService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRequestService(opts ...option.RequestOption) (r *RequestService) {
	r = &RequestService{}
	r.Options = opts
	r.Message = NewRequestMessageService(opts...)
	r.Priority = NewRequestPriorityService(opts...)
	r.Assets = NewRequestAssetService(opts...)
	return
}

// Creating a request adds the request into the Cloudforce One queue for analysis.
// In addition to the content, a short title, type, priority, and releasability
// should be provided. If one is not provided, a default will be assigned.
func (r *RequestService) New(ctx context.Context, accountIdentifier string, body RequestNewParams, opts ...option.RequestOption) (res *RequestNewResponse, err error) {
	var env RequestNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
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
// Only fields that you choose to update need to be add to the request body.
func (r *RequestService) Update(ctx context.Context, accountIdentifier string, requestIdentifier string, body RequestUpdateParams, opts ...option.RequestOption) (res *RequestUpdateResponse, err error) {
	var env RequestUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if requestIdentifier == "" {
		err = errors.New("missing required request_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s", accountIdentifier, requestIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Requests
func (r *RequestService) List(ctx context.Context, accountIdentifier string, body RequestListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[RequestListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
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
func (r *RequestService) ListAutoPaging(ctx context.Context, accountIdentifier string, body RequestListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[RequestListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, accountIdentifier, body, opts...))
}

// Delete a Request
func (r *RequestService) Delete(ctx context.Context, accountIdentifier string, requestIdentifier string, opts ...option.RequestOption) (res *RequestDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if requestIdentifier == "" {
		err = errors.New("missing required request_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s", accountIdentifier, requestIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get Request Priority, Status, and TLP constants
func (r *RequestService) Constants(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *RequestConstantsResponse, err error) {
	var env RequestConstantsResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
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
	var env RequestGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if requestIdentifier == "" {
		err = errors.New("missing required request_identifier parameter")
		return
	}
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
	var env RequestQuotaResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
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
	var env RequestTypesResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
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
	TLP       RequestNewResponseTLP `json:"tlp,required"`
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
	TLP           apijson.Field
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
type RequestNewResponseTLP string

const (
	RequestNewResponseTLPClear       RequestNewResponseTLP = "clear"
	RequestNewResponseTLPAmber       RequestNewResponseTLP = "amber"
	RequestNewResponseTLPAmberStrict RequestNewResponseTLP = "amber-strict"
	RequestNewResponseTLPGreen       RequestNewResponseTLP = "green"
	RequestNewResponseTLPRed         RequestNewResponseTLP = "red"
)

func (r RequestNewResponseTLP) IsKnown() bool {
	switch r {
	case RequestNewResponseTLPClear, RequestNewResponseTLPAmber, RequestNewResponseTLPAmberStrict, RequestNewResponseTLPGreen, RequestNewResponseTLPRed:
		return true
	}
	return false
}

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

func (r RequestNewResponseStatus) IsKnown() bool {
	switch r {
	case RequestNewResponseStatusOpen, RequestNewResponseStatusAccepted, RequestNewResponseStatusReported, RequestNewResponseStatusApproved, RequestNewResponseStatusCompleted, RequestNewResponseStatusDeclined:
		return true
	}
	return false
}

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
	TLP       RequestUpdateResponseTLP `json:"tlp,required"`
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
	TLP           apijson.Field
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
type RequestUpdateResponseTLP string

const (
	RequestUpdateResponseTLPClear       RequestUpdateResponseTLP = "clear"
	RequestUpdateResponseTLPAmber       RequestUpdateResponseTLP = "amber"
	RequestUpdateResponseTLPAmberStrict RequestUpdateResponseTLP = "amber-strict"
	RequestUpdateResponseTLPGreen       RequestUpdateResponseTLP = "green"
	RequestUpdateResponseTLPRed         RequestUpdateResponseTLP = "red"
)

func (r RequestUpdateResponseTLP) IsKnown() bool {
	switch r {
	case RequestUpdateResponseTLPClear, RequestUpdateResponseTLPAmber, RequestUpdateResponseTLPAmberStrict, RequestUpdateResponseTLPGreen, RequestUpdateResponseTLPRed:
		return true
	}
	return false
}

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

func (r RequestUpdateResponseStatus) IsKnown() bool {
	switch r {
	case RequestUpdateResponseStatusOpen, RequestUpdateResponseStatusAccepted, RequestUpdateResponseStatusReported, RequestUpdateResponseStatusApproved, RequestUpdateResponseStatusCompleted, RequestUpdateResponseStatusDeclined:
		return true
	}
	return false
}

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
	TLP RequestListResponseTLP `json:"tlp,required"`
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
	TLP           apijson.Field
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

func (r RequestListResponsePriority) IsKnown() bool {
	switch r {
	case RequestListResponsePriorityRoutine, RequestListResponsePriorityHigh, RequestListResponsePriorityUrgent:
		return true
	}
	return false
}

// The CISA defined Traffic Light Protocol (TLP)
type RequestListResponseTLP string

const (
	RequestListResponseTLPClear       RequestListResponseTLP = "clear"
	RequestListResponseTLPAmber       RequestListResponseTLP = "amber"
	RequestListResponseTLPAmberStrict RequestListResponseTLP = "amber-strict"
	RequestListResponseTLPGreen       RequestListResponseTLP = "green"
	RequestListResponseTLPRed         RequestListResponseTLP = "red"
)

func (r RequestListResponseTLP) IsKnown() bool {
	switch r {
	case RequestListResponseTLPClear, RequestListResponseTLPAmber, RequestListResponseTLPAmberStrict, RequestListResponseTLPGreen, RequestListResponseTLPRed:
		return true
	}
	return false
}

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

func (r RequestListResponseStatus) IsKnown() bool {
	switch r {
	case RequestListResponseStatusOpen, RequestListResponseStatusAccepted, RequestListResponseStatusReported, RequestListResponseStatusApproved, RequestListResponseStatusCompleted, RequestListResponseStatusDeclined:
		return true
	}
	return false
}

type RequestDeleteResponse struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RequestDeleteResponseSuccess `json:"success,required"`
	JSON    requestDeleteResponseJSON    `json:"-"`
}

// requestDeleteResponseJSON contains the JSON metadata for the struct
// [RequestDeleteResponse]
type requestDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestDeleteResponseSuccess bool

const (
	RequestDeleteResponseSuccessTrue RequestDeleteResponseSuccess = true
)

func (r RequestDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case RequestDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type RequestConstantsResponse struct {
	Priority []RequestConstantsResponsePriority `json:"priority"`
	Status   []RequestConstantsResponseStatus   `json:"status"`
	TLP      []RequestConstantsResponseTLP      `json:"tlp"`
	JSON     requestConstantsResponseJSON       `json:"-"`
}

// requestConstantsResponseJSON contains the JSON metadata for the struct
// [RequestConstantsResponse]
type requestConstantsResponseJSON struct {
	Priority    apijson.Field
	Status      apijson.Field
	TLP         apijson.Field
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

func (r RequestConstantsResponsePriority) IsKnown() bool {
	switch r {
	case RequestConstantsResponsePriorityRoutine, RequestConstantsResponsePriorityHigh, RequestConstantsResponsePriorityUrgent:
		return true
	}
	return false
}

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

func (r RequestConstantsResponseStatus) IsKnown() bool {
	switch r {
	case RequestConstantsResponseStatusOpen, RequestConstantsResponseStatusAccepted, RequestConstantsResponseStatusReported, RequestConstantsResponseStatusApproved, RequestConstantsResponseStatusCompleted, RequestConstantsResponseStatusDeclined:
		return true
	}
	return false
}

// The CISA defined Traffic Light Protocol (TLP)
type RequestConstantsResponseTLP string

const (
	RequestConstantsResponseTLPClear       RequestConstantsResponseTLP = "clear"
	RequestConstantsResponseTLPAmber       RequestConstantsResponseTLP = "amber"
	RequestConstantsResponseTLPAmberStrict RequestConstantsResponseTLP = "amber-strict"
	RequestConstantsResponseTLPGreen       RequestConstantsResponseTLP = "green"
	RequestConstantsResponseTLPRed         RequestConstantsResponseTLP = "red"
)

func (r RequestConstantsResponseTLP) IsKnown() bool {
	switch r {
	case RequestConstantsResponseTLPClear, RequestConstantsResponseTLPAmber, RequestConstantsResponseTLPAmberStrict, RequestConstantsResponseTLPGreen, RequestConstantsResponseTLPRed:
		return true
	}
	return false
}

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
	TLP       RequestGetResponseTLP `json:"tlp,required"`
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
	TLP           apijson.Field
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
type RequestGetResponseTLP string

const (
	RequestGetResponseTLPClear       RequestGetResponseTLP = "clear"
	RequestGetResponseTLPAmber       RequestGetResponseTLP = "amber"
	RequestGetResponseTLPAmberStrict RequestGetResponseTLP = "amber-strict"
	RequestGetResponseTLPGreen       RequestGetResponseTLP = "green"
	RequestGetResponseTLPRed         RequestGetResponseTLP = "red"
)

func (r RequestGetResponseTLP) IsKnown() bool {
	switch r {
	case RequestGetResponseTLPClear, RequestGetResponseTLPAmber, RequestGetResponseTLPAmberStrict, RequestGetResponseTLPGreen, RequestGetResponseTLPRed:
		return true
	}
	return false
}

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

func (r RequestGetResponseStatus) IsKnown() bool {
	switch r {
	case RequestGetResponseStatusOpen, RequestGetResponseStatusAccepted, RequestGetResponseStatusReported, RequestGetResponseStatusApproved, RequestGetResponseStatusCompleted, RequestGetResponseStatusDeclined:
		return true
	}
	return false
}

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
	TLP param.Field[RequestNewParamsTLP] `json:"tlp"`
}

func (r RequestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The CISA defined Traffic Light Protocol (TLP)
type RequestNewParamsTLP string

const (
	RequestNewParamsTLPClear       RequestNewParamsTLP = "clear"
	RequestNewParamsTLPAmber       RequestNewParamsTLP = "amber"
	RequestNewParamsTLPAmberStrict RequestNewParamsTLP = "amber-strict"
	RequestNewParamsTLPGreen       RequestNewParamsTLP = "green"
	RequestNewParamsTLPRed         RequestNewParamsTLP = "red"
)

func (r RequestNewParamsTLP) IsKnown() bool {
	switch r {
	case RequestNewParamsTLPClear, RequestNewParamsTLPAmber, RequestNewParamsTLPAmberStrict, RequestNewParamsTLPGreen, RequestNewParamsTLPRed:
		return true
	}
	return false
}

type RequestNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RequestNewResponseEnvelopeSuccess `json:"success,required"`
	Result  RequestNewResponse                `json:"result"`
	JSON    requestNewResponseEnvelopeJSON    `json:"-"`
}

// requestNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestNewResponseEnvelope]
type requestNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestNewResponseEnvelopeSuccess bool

const (
	RequestNewResponseEnvelopeSuccessTrue RequestNewResponseEnvelopeSuccess = true
)

func (r RequestNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RequestNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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
	TLP param.Field[RequestUpdateParamsTLP] `json:"tlp"`
}

func (r RequestUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The CISA defined Traffic Light Protocol (TLP)
type RequestUpdateParamsTLP string

const (
	RequestUpdateParamsTLPClear       RequestUpdateParamsTLP = "clear"
	RequestUpdateParamsTLPAmber       RequestUpdateParamsTLP = "amber"
	RequestUpdateParamsTLPAmberStrict RequestUpdateParamsTLP = "amber-strict"
	RequestUpdateParamsTLPGreen       RequestUpdateParamsTLP = "green"
	RequestUpdateParamsTLPRed         RequestUpdateParamsTLP = "red"
)

func (r RequestUpdateParamsTLP) IsKnown() bool {
	switch r {
	case RequestUpdateParamsTLPClear, RequestUpdateParamsTLPAmber, RequestUpdateParamsTLPAmberStrict, RequestUpdateParamsTLPGreen, RequestUpdateParamsTLPRed:
		return true
	}
	return false
}

type RequestUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RequestUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  RequestUpdateResponse                `json:"result"`
	JSON    requestUpdateResponseEnvelopeJSON    `json:"-"`
}

// requestUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestUpdateResponseEnvelope]
type requestUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestUpdateResponseEnvelopeSuccess bool

const (
	RequestUpdateResponseEnvelopeSuccessTrue RequestUpdateResponseEnvelopeSuccess = true
)

func (r RequestUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RequestUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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

func (r RequestListParamsSortOrder) IsKnown() bool {
	switch r {
	case RequestListParamsSortOrderAsc, RequestListParamsSortOrderDesc:
		return true
	}
	return false
}

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

func (r RequestListParamsStatus) IsKnown() bool {
	switch r {
	case RequestListParamsStatusOpen, RequestListParamsStatusAccepted, RequestListParamsStatusReported, RequestListParamsStatusApproved, RequestListParamsStatusCompleted, RequestListParamsStatusDeclined:
		return true
	}
	return false
}

type RequestConstantsResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RequestConstantsResponseEnvelopeSuccess `json:"success,required"`
	Result  RequestConstantsResponse                `json:"result"`
	JSON    requestConstantsResponseEnvelopeJSON    `json:"-"`
}

// requestConstantsResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestConstantsResponseEnvelope]
type requestConstantsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestConstantsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestConstantsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestConstantsResponseEnvelopeSuccess bool

const (
	RequestConstantsResponseEnvelopeSuccessTrue RequestConstantsResponseEnvelopeSuccess = true
)

func (r RequestConstantsResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RequestConstantsResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RequestGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RequestGetResponseEnvelopeSuccess `json:"success,required"`
	Result  RequestGetResponse                `json:"result"`
	JSON    requestGetResponseEnvelopeJSON    `json:"-"`
}

// requestGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestGetResponseEnvelope]
type requestGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestGetResponseEnvelopeSuccess bool

const (
	RequestGetResponseEnvelopeSuccessTrue RequestGetResponseEnvelopeSuccess = true
)

func (r RequestGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RequestGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RequestQuotaResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RequestQuotaResponseEnvelopeSuccess `json:"success,required"`
	Result  RequestQuotaResponse                `json:"result"`
	JSON    requestQuotaResponseEnvelopeJSON    `json:"-"`
}

// requestQuotaResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestQuotaResponseEnvelope]
type requestQuotaResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestQuotaResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestQuotaResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestQuotaResponseEnvelopeSuccess bool

const (
	RequestQuotaResponseEnvelopeSuccessTrue RequestQuotaResponseEnvelopeSuccess = true
)

func (r RequestQuotaResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RequestQuotaResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RequestTypesResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RequestTypesResponseEnvelopeSuccess `json:"success,required"`
	Result  []string                            `json:"result"`
	JSON    requestTypesResponseEnvelopeJSON    `json:"-"`
}

// requestTypesResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestTypesResponseEnvelope]
type requestTypesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestTypesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestTypesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestTypesResponseEnvelopeSuccess bool

const (
	RequestTypesResponseEnvelopeSuccessTrue RequestTypesResponseEnvelopeSuccess = true
)

func (r RequestTypesResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RequestTypesResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
