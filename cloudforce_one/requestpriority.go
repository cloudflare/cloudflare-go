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
	"github.com/cloudflare/cloudflare-go/v3/shared"
)

// RequestPriorityService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRequestPriorityService] method instead.
type RequestPriorityService struct {
	Options []option.RequestOption
}

// NewRequestPriorityService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRequestPriorityService(opts ...option.RequestOption) (r *RequestPriorityService) {
	r = &RequestPriorityService{}
	r.Options = opts
	return
}

// Create a New Priority Intelligence Requirement
func (r *RequestPriorityService) New(ctx context.Context, accountIdentifier string, body RequestPriorityNewParams, opts ...option.RequestOption) (res *RequestPriorityNewResponse, err error) {
	var env RequestPriorityNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/new", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a Priority Intelligence Requirement
func (r *RequestPriorityService) Update(ctx context.Context, accountIdentifier string, priorityIdentifer string, body RequestPriorityUpdateParams, opts ...option.RequestOption) (res *RequestPriorityUpdateResponse, err error) {
	var env RequestPriorityUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if priorityIdentifer == "" {
		err = errors.New("missing required priority_identifer parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/%s", accountIdentifier, priorityIdentifer)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a Priority Intelligence Requirement
func (r *RequestPriorityService) Delete(ctx context.Context, accountIdentifier string, priorityIdentifer string, opts ...option.RequestOption) (res *RequestPriorityDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if priorityIdentifer == "" {
		err = errors.New("missing required priority_identifer parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/%s", accountIdentifier, priorityIdentifer)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get a Priority Intelligence Requirement
func (r *RequestPriorityService) Get(ctx context.Context, accountIdentifier string, priorityIdentifer string, opts ...option.RequestOption) (res *RequestPriorityGetResponse, err error) {
	var env RequestPriorityGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if priorityIdentifer == "" {
		err = errors.New("missing required priority_identifer parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/%s", accountIdentifier, priorityIdentifer)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Priority Intelligence Requirement Quota
func (r *RequestPriorityService) Quota(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *RequestPriorityQuotaResponse, err error) {
	var env RequestPriorityQuotaResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/quota", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RequestPriorityNewResponse struct {
	// UUID
	ID string `json:"id,required"`
	// Priority creation time
	Created time.Time `json:"created,required" format:"date-time"`
	// List of labels
	Labels []string `json:"labels,required"`
	// Priority
	Priority int64 `json:"priority,required"`
	// Requirement
	Requirement string `json:"requirement,required"`
	// The CISA defined Traffic Light Protocol (TLP)
	TLP RequestPriorityNewResponseTLP `json:"tlp,required"`
	// Priority last updated time
	Updated time.Time                      `json:"updated,required" format:"date-time"`
	JSON    requestPriorityNewResponseJSON `json:"-"`
}

// requestPriorityNewResponseJSON contains the JSON metadata for the struct
// [RequestPriorityNewResponse]
type requestPriorityNewResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Labels      apijson.Field
	Priority    apijson.Field
	Requirement apijson.Field
	TLP         apijson.Field
	Updated     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityNewResponseJSON) RawJSON() string {
	return r.raw
}

// The CISA defined Traffic Light Protocol (TLP)
type RequestPriorityNewResponseTLP string

const (
	RequestPriorityNewResponseTLPClear       RequestPriorityNewResponseTLP = "clear"
	RequestPriorityNewResponseTLPAmber       RequestPriorityNewResponseTLP = "amber"
	RequestPriorityNewResponseTLPAmberStrict RequestPriorityNewResponseTLP = "amber-strict"
	RequestPriorityNewResponseTLPGreen       RequestPriorityNewResponseTLP = "green"
	RequestPriorityNewResponseTLPRed         RequestPriorityNewResponseTLP = "red"
)

func (r RequestPriorityNewResponseTLP) IsKnown() bool {
	switch r {
	case RequestPriorityNewResponseTLPClear, RequestPriorityNewResponseTLPAmber, RequestPriorityNewResponseTLPAmberStrict, RequestPriorityNewResponseTLPGreen, RequestPriorityNewResponseTLPRed:
		return true
	}
	return false
}

type RequestPriorityUpdateResponse struct {
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
	TLP       RequestPriorityUpdateResponseTLP `json:"tlp,required"`
	Updated   time.Time                        `json:"updated,required" format:"date-time"`
	Completed time.Time                        `json:"completed" format:"date-time"`
	// Tokens for the request messages
	MessageTokens int64 `json:"message_tokens"`
	// Readable Request ID
	ReadableID string `json:"readable_id"`
	// Request Status
	Status RequestPriorityUpdateResponseStatus `json:"status"`
	// Tokens for the request
	Tokens int64                             `json:"tokens"`
	JSON   requestPriorityUpdateResponseJSON `json:"-"`
}

// requestPriorityUpdateResponseJSON contains the JSON metadata for the struct
// [RequestPriorityUpdateResponse]
type requestPriorityUpdateResponseJSON struct {
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

func (r *RequestPriorityUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The CISA defined Traffic Light Protocol (TLP)
type RequestPriorityUpdateResponseTLP string

const (
	RequestPriorityUpdateResponseTLPClear       RequestPriorityUpdateResponseTLP = "clear"
	RequestPriorityUpdateResponseTLPAmber       RequestPriorityUpdateResponseTLP = "amber"
	RequestPriorityUpdateResponseTLPAmberStrict RequestPriorityUpdateResponseTLP = "amber-strict"
	RequestPriorityUpdateResponseTLPGreen       RequestPriorityUpdateResponseTLP = "green"
	RequestPriorityUpdateResponseTLPRed         RequestPriorityUpdateResponseTLP = "red"
)

func (r RequestPriorityUpdateResponseTLP) IsKnown() bool {
	switch r {
	case RequestPriorityUpdateResponseTLPClear, RequestPriorityUpdateResponseTLPAmber, RequestPriorityUpdateResponseTLPAmberStrict, RequestPriorityUpdateResponseTLPGreen, RequestPriorityUpdateResponseTLPRed:
		return true
	}
	return false
}

// Request Status
type RequestPriorityUpdateResponseStatus string

const (
	RequestPriorityUpdateResponseStatusOpen      RequestPriorityUpdateResponseStatus = "open"
	RequestPriorityUpdateResponseStatusAccepted  RequestPriorityUpdateResponseStatus = "accepted"
	RequestPriorityUpdateResponseStatusReported  RequestPriorityUpdateResponseStatus = "reported"
	RequestPriorityUpdateResponseStatusApproved  RequestPriorityUpdateResponseStatus = "approved"
	RequestPriorityUpdateResponseStatusCompleted RequestPriorityUpdateResponseStatus = "completed"
	RequestPriorityUpdateResponseStatusDeclined  RequestPriorityUpdateResponseStatus = "declined"
)

func (r RequestPriorityUpdateResponseStatus) IsKnown() bool {
	switch r {
	case RequestPriorityUpdateResponseStatusOpen, RequestPriorityUpdateResponseStatusAccepted, RequestPriorityUpdateResponseStatusReported, RequestPriorityUpdateResponseStatusApproved, RequestPriorityUpdateResponseStatusCompleted, RequestPriorityUpdateResponseStatusDeclined:
		return true
	}
	return false
}

type RequestPriorityDeleteResponse struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RequestPriorityDeleteResponseSuccess `json:"success,required"`
	JSON    requestPriorityDeleteResponseJSON    `json:"-"`
}

// requestPriorityDeleteResponseJSON contains the JSON metadata for the struct
// [RequestPriorityDeleteResponse]
type requestPriorityDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestPriorityDeleteResponseSuccess bool

const (
	RequestPriorityDeleteResponseSuccessTrue RequestPriorityDeleteResponseSuccess = true
)

func (r RequestPriorityDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case RequestPriorityDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type RequestPriorityGetResponse struct {
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
	TLP       RequestPriorityGetResponseTLP `json:"tlp,required"`
	Updated   time.Time                     `json:"updated,required" format:"date-time"`
	Completed time.Time                     `json:"completed" format:"date-time"`
	// Tokens for the request messages
	MessageTokens int64 `json:"message_tokens"`
	// Readable Request ID
	ReadableID string `json:"readable_id"`
	// Request Status
	Status RequestPriorityGetResponseStatus `json:"status"`
	// Tokens for the request
	Tokens int64                          `json:"tokens"`
	JSON   requestPriorityGetResponseJSON `json:"-"`
}

// requestPriorityGetResponseJSON contains the JSON metadata for the struct
// [RequestPriorityGetResponse]
type requestPriorityGetResponseJSON struct {
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

func (r *RequestPriorityGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityGetResponseJSON) RawJSON() string {
	return r.raw
}

// The CISA defined Traffic Light Protocol (TLP)
type RequestPriorityGetResponseTLP string

const (
	RequestPriorityGetResponseTLPClear       RequestPriorityGetResponseTLP = "clear"
	RequestPriorityGetResponseTLPAmber       RequestPriorityGetResponseTLP = "amber"
	RequestPriorityGetResponseTLPAmberStrict RequestPriorityGetResponseTLP = "amber-strict"
	RequestPriorityGetResponseTLPGreen       RequestPriorityGetResponseTLP = "green"
	RequestPriorityGetResponseTLPRed         RequestPriorityGetResponseTLP = "red"
)

func (r RequestPriorityGetResponseTLP) IsKnown() bool {
	switch r {
	case RequestPriorityGetResponseTLPClear, RequestPriorityGetResponseTLPAmber, RequestPriorityGetResponseTLPAmberStrict, RequestPriorityGetResponseTLPGreen, RequestPriorityGetResponseTLPRed:
		return true
	}
	return false
}

// Request Status
type RequestPriorityGetResponseStatus string

const (
	RequestPriorityGetResponseStatusOpen      RequestPriorityGetResponseStatus = "open"
	RequestPriorityGetResponseStatusAccepted  RequestPriorityGetResponseStatus = "accepted"
	RequestPriorityGetResponseStatusReported  RequestPriorityGetResponseStatus = "reported"
	RequestPriorityGetResponseStatusApproved  RequestPriorityGetResponseStatus = "approved"
	RequestPriorityGetResponseStatusCompleted RequestPriorityGetResponseStatus = "completed"
	RequestPriorityGetResponseStatusDeclined  RequestPriorityGetResponseStatus = "declined"
)

func (r RequestPriorityGetResponseStatus) IsKnown() bool {
	switch r {
	case RequestPriorityGetResponseStatusOpen, RequestPriorityGetResponseStatusAccepted, RequestPriorityGetResponseStatusReported, RequestPriorityGetResponseStatusApproved, RequestPriorityGetResponseStatusCompleted, RequestPriorityGetResponseStatusDeclined:
		return true
	}
	return false
}

type RequestPriorityQuotaResponse struct {
	// Anniversary date is when annual quota limit is refresh
	AnniversaryDate time.Time `json:"anniversary_date" format:"date-time"`
	// Quater anniversary date is when quota limit is refreshed each quarter
	QuarterAnniversaryDate time.Time `json:"quarter_anniversary_date" format:"date-time"`
	// Tokens for the quarter
	Quota int64 `json:"quota"`
	// Tokens remaining for the quarter
	Remaining int64                            `json:"remaining"`
	JSON      requestPriorityQuotaResponseJSON `json:"-"`
}

// requestPriorityQuotaResponseJSON contains the JSON metadata for the struct
// [RequestPriorityQuotaResponse]
type requestPriorityQuotaResponseJSON struct {
	AnniversaryDate        apijson.Field
	QuarterAnniversaryDate apijson.Field
	Quota                  apijson.Field
	Remaining              apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *RequestPriorityQuotaResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityQuotaResponseJSON) RawJSON() string {
	return r.raw
}

type RequestPriorityNewParams struct {
	// List of labels
	Labels param.Field[[]string] `json:"labels,required"`
	// Priority
	Priority param.Field[int64] `json:"priority,required"`
	// Requirement
	Requirement param.Field[string] `json:"requirement,required"`
	// The CISA defined Traffic Light Protocol (TLP)
	TLP param.Field[RequestPriorityNewParamsTLP] `json:"tlp,required"`
}

func (r RequestPriorityNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The CISA defined Traffic Light Protocol (TLP)
type RequestPriorityNewParamsTLP string

const (
	RequestPriorityNewParamsTLPClear       RequestPriorityNewParamsTLP = "clear"
	RequestPriorityNewParamsTLPAmber       RequestPriorityNewParamsTLP = "amber"
	RequestPriorityNewParamsTLPAmberStrict RequestPriorityNewParamsTLP = "amber-strict"
	RequestPriorityNewParamsTLPGreen       RequestPriorityNewParamsTLP = "green"
	RequestPriorityNewParamsTLPRed         RequestPriorityNewParamsTLP = "red"
)

func (r RequestPriorityNewParamsTLP) IsKnown() bool {
	switch r {
	case RequestPriorityNewParamsTLPClear, RequestPriorityNewParamsTLPAmber, RequestPriorityNewParamsTLPAmberStrict, RequestPriorityNewParamsTLPGreen, RequestPriorityNewParamsTLPRed:
		return true
	}
	return false
}

type RequestPriorityNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RequestPriorityNewResponseEnvelopeSuccess `json:"success,required"`
	Result  RequestPriorityNewResponse                `json:"result"`
	JSON    requestPriorityNewResponseEnvelopeJSON    `json:"-"`
}

// requestPriorityNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestPriorityNewResponseEnvelope]
type requestPriorityNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestPriorityNewResponseEnvelopeSuccess bool

const (
	RequestPriorityNewResponseEnvelopeSuccessTrue RequestPriorityNewResponseEnvelopeSuccess = true
)

func (r RequestPriorityNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RequestPriorityNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RequestPriorityUpdateParams struct {
	// List of labels
	Labels param.Field[[]string] `json:"labels,required"`
	// Priority
	Priority param.Field[int64] `json:"priority,required"`
	// Requirement
	Requirement param.Field[string] `json:"requirement,required"`
	// The CISA defined Traffic Light Protocol (TLP)
	TLP param.Field[RequestPriorityUpdateParamsTLP] `json:"tlp,required"`
}

func (r RequestPriorityUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The CISA defined Traffic Light Protocol (TLP)
type RequestPriorityUpdateParamsTLP string

const (
	RequestPriorityUpdateParamsTLPClear       RequestPriorityUpdateParamsTLP = "clear"
	RequestPriorityUpdateParamsTLPAmber       RequestPriorityUpdateParamsTLP = "amber"
	RequestPriorityUpdateParamsTLPAmberStrict RequestPriorityUpdateParamsTLP = "amber-strict"
	RequestPriorityUpdateParamsTLPGreen       RequestPriorityUpdateParamsTLP = "green"
	RequestPriorityUpdateParamsTLPRed         RequestPriorityUpdateParamsTLP = "red"
)

func (r RequestPriorityUpdateParamsTLP) IsKnown() bool {
	switch r {
	case RequestPriorityUpdateParamsTLPClear, RequestPriorityUpdateParamsTLPAmber, RequestPriorityUpdateParamsTLPAmberStrict, RequestPriorityUpdateParamsTLPGreen, RequestPriorityUpdateParamsTLPRed:
		return true
	}
	return false
}

type RequestPriorityUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RequestPriorityUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  RequestPriorityUpdateResponse                `json:"result"`
	JSON    requestPriorityUpdateResponseEnvelopeJSON    `json:"-"`
}

// requestPriorityUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [RequestPriorityUpdateResponseEnvelope]
type requestPriorityUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestPriorityUpdateResponseEnvelopeSuccess bool

const (
	RequestPriorityUpdateResponseEnvelopeSuccessTrue RequestPriorityUpdateResponseEnvelopeSuccess = true
)

func (r RequestPriorityUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RequestPriorityUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RequestPriorityGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RequestPriorityGetResponseEnvelopeSuccess `json:"success,required"`
	Result  RequestPriorityGetResponse                `json:"result"`
	JSON    requestPriorityGetResponseEnvelopeJSON    `json:"-"`
}

// requestPriorityGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestPriorityGetResponseEnvelope]
type requestPriorityGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestPriorityGetResponseEnvelopeSuccess bool

const (
	RequestPriorityGetResponseEnvelopeSuccessTrue RequestPriorityGetResponseEnvelopeSuccess = true
)

func (r RequestPriorityGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RequestPriorityGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RequestPriorityQuotaResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RequestPriorityQuotaResponseEnvelopeSuccess `json:"success,required"`
	Result  RequestPriorityQuotaResponse                `json:"result"`
	JSON    requestPriorityQuotaResponseEnvelopeJSON    `json:"-"`
}

// requestPriorityQuotaResponseEnvelopeJSON contains the JSON metadata for the
// struct [RequestPriorityQuotaResponseEnvelope]
type requestPriorityQuotaResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityQuotaResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityQuotaResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestPriorityQuotaResponseEnvelopeSuccess bool

const (
	RequestPriorityQuotaResponseEnvelopeSuccessTrue RequestPriorityQuotaResponseEnvelopeSuccess = true
)

func (r RequestPriorityQuotaResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RequestPriorityQuotaResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
