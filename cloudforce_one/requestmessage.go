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

// RequestMessageService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRequestMessageService] method
// instead.
type RequestMessageService struct {
	Options []option.RequestOption
}

// NewRequestMessageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRequestMessageService(opts ...option.RequestOption) (r *RequestMessageService) {
	r = &RequestMessageService{}
	r.Options = opts
	return
}

// Creating a request adds the request into the Cloudforce One queue for analysis.
// In addition to the content, a short title, type, priority, and releasability
// should be provided. If one is not provided a default will be assigned.
func (r *RequestMessageService) New(ctx context.Context, accountIdentifier string, requestIdentifier string, body RequestMessageNewParams, opts ...option.RequestOption) (res *RequestMessageNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RequestMessageNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/message/new", accountIdentifier, requestIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a Request Message
func (r *RequestMessageService) Update(ctx context.Context, accountIdentifier string, requestIdentifier string, messageIdentifer int64, body RequestMessageUpdateParams, opts ...option.RequestOption) (res *RequestMessageUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RequestMessageUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/message/%v", accountIdentifier, requestIdentifier, messageIdentifer)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a Request Message
func (r *RequestMessageService) Delete(ctx context.Context, accountIdentifier string, requestIdentifier string, messageIdentifer int64, opts ...option.RequestOption) (res *RequestMessageDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RequestMessageDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/message/%v", accountIdentifier, requestIdentifier, messageIdentifer)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Request Messages
func (r *RequestMessageService) Get(ctx context.Context, accountIdentifier string, requestIdentifier string, body RequestMessageGetParams, opts ...option.RequestOption) (res *[]RequestMessageGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RequestMessageGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/message", accountIdentifier, requestIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RequestMessageNewResponse struct {
	// Message ID
	ID int64 `json:"id,required"`
	// Author of message
	Author string `json:"author,required"`
	// Content of message
	Content string `json:"content,required"`
	// Message is a follow-on request
	IsFollowOnRequest bool `json:"is_follow_on_request,required"`
	// Message last updated time
	Updated time.Time `json:"updated,required" format:"date-time"`
	// Message creation time
	Created time.Time                     `json:"created" format:"date-time"`
	JSON    requestMessageNewResponseJSON `json:"-"`
}

// requestMessageNewResponseJSON contains the JSON metadata for the struct
// [RequestMessageNewResponse]
type requestMessageNewResponseJSON struct {
	ID                apijson.Field
	Author            apijson.Field
	Content           apijson.Field
	IsFollowOnRequest apijson.Field
	Updated           apijson.Field
	Created           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RequestMessageNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageNewResponseJSON) RawJSON() string {
	return r.raw
}

type RequestMessageUpdateResponse struct {
	// Message ID
	ID int64 `json:"id,required"`
	// Author of message
	Author string `json:"author,required"`
	// Content of message
	Content string `json:"content,required"`
	// Message is a follow-on request
	IsFollowOnRequest bool `json:"is_follow_on_request,required"`
	// Message last updated time
	Updated time.Time `json:"updated,required" format:"date-time"`
	// Message creation time
	Created time.Time                        `json:"created" format:"date-time"`
	JSON    requestMessageUpdateResponseJSON `json:"-"`
}

// requestMessageUpdateResponseJSON contains the JSON metadata for the struct
// [RequestMessageUpdateResponse]
type requestMessageUpdateResponseJSON struct {
	ID                apijson.Field
	Author            apijson.Field
	Content           apijson.Field
	IsFollowOnRequest apijson.Field
	Updated           apijson.Field
	Created           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RequestMessageUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [cloudforce_one.RequestMessageDeleteResponseUnknown],
// [cloudforce_one.RequestMessageDeleteResponseArray] or [shared.UnionString].
type RequestMessageDeleteResponse interface {
	ImplementsCloudforceOneRequestMessageDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RequestMessageDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RequestMessageDeleteResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type RequestMessageDeleteResponseArray []interface{}

func (r RequestMessageDeleteResponseArray) ImplementsCloudforceOneRequestMessageDeleteResponse() {}

type RequestMessageGetResponse struct {
	// Message ID
	ID int64 `json:"id,required"`
	// Author of message
	Author string `json:"author,required"`
	// Content of message
	Content string `json:"content,required"`
	// Message is a follow-on request
	IsFollowOnRequest bool `json:"is_follow_on_request,required"`
	// Message last updated time
	Updated time.Time `json:"updated,required" format:"date-time"`
	// Message creation time
	Created time.Time                     `json:"created" format:"date-time"`
	JSON    requestMessageGetResponseJSON `json:"-"`
}

// requestMessageGetResponseJSON contains the JSON metadata for the struct
// [RequestMessageGetResponse]
type requestMessageGetResponseJSON struct {
	ID                apijson.Field
	Author            apijson.Field
	Content           apijson.Field
	IsFollowOnRequest apijson.Field
	Updated           apijson.Field
	Created           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RequestMessageGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageGetResponseJSON) RawJSON() string {
	return r.raw
}

type RequestMessageNewParams struct {
	// Content of message
	Content param.Field[string] `json:"content"`
}

func (r RequestMessageNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RequestMessageNewResponseEnvelope struct {
	Errors   []RequestMessageNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RequestMessageNewResponseEnvelopeMessages `json:"messages,required"`
	Result   RequestMessageNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success RequestMessageNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    requestMessageNewResponseEnvelopeJSON    `json:"-"`
}

// requestMessageNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestMessageNewResponseEnvelope]
type requestMessageNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestMessageNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RequestMessageNewResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    requestMessageNewResponseEnvelopeErrorsJSON `json:"-"`
}

// requestMessageNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RequestMessageNewResponseEnvelopeErrors]
type requestMessageNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestMessageNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RequestMessageNewResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    requestMessageNewResponseEnvelopeMessagesJSON `json:"-"`
}

// requestMessageNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RequestMessageNewResponseEnvelopeMessages]
type requestMessageNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestMessageNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestMessageNewResponseEnvelopeSuccess bool

const (
	RequestMessageNewResponseEnvelopeSuccessTrue RequestMessageNewResponseEnvelopeSuccess = true
)

type RequestMessageUpdateParams struct {
	// Request content
	Content param.Field[string] `json:"content"`
	// Priority for analyzing the request
	Priority param.Field[string] `json:"priority"`
	// Requested information from request
	RequestType param.Field[string] `json:"request_type"`
	// Brief description of the request
	Summary param.Field[string] `json:"summary"`
	// The CISA defined Traffic Light Protocol (TLP)
	Tlp param.Field[RequestMessageUpdateParamsTlp] `json:"tlp"`
}

func (r RequestMessageUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The CISA defined Traffic Light Protocol (TLP)
type RequestMessageUpdateParamsTlp string

const (
	RequestMessageUpdateParamsTlpClear       RequestMessageUpdateParamsTlp = "clear"
	RequestMessageUpdateParamsTlpAmber       RequestMessageUpdateParamsTlp = "amber"
	RequestMessageUpdateParamsTlpAmberStrict RequestMessageUpdateParamsTlp = "amber-strict"
	RequestMessageUpdateParamsTlpGreen       RequestMessageUpdateParamsTlp = "green"
	RequestMessageUpdateParamsTlpRed         RequestMessageUpdateParamsTlp = "red"
)

type RequestMessageUpdateResponseEnvelope struct {
	Errors   []RequestMessageUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RequestMessageUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   RequestMessageUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success RequestMessageUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    requestMessageUpdateResponseEnvelopeJSON    `json:"-"`
}

// requestMessageUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [RequestMessageUpdateResponseEnvelope]
type requestMessageUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestMessageUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RequestMessageUpdateResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    requestMessageUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// requestMessageUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [RequestMessageUpdateResponseEnvelopeErrors]
type requestMessageUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestMessageUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RequestMessageUpdateResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    requestMessageUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// requestMessageUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [RequestMessageUpdateResponseEnvelopeMessages]
type requestMessageUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestMessageUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestMessageUpdateResponseEnvelopeSuccess bool

const (
	RequestMessageUpdateResponseEnvelopeSuccessTrue RequestMessageUpdateResponseEnvelopeSuccess = true
)

type RequestMessageDeleteResponseEnvelope struct {
	Errors   []RequestMessageDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RequestMessageDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   RequestMessageDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success RequestMessageDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    requestMessageDeleteResponseEnvelopeJSON    `json:"-"`
}

// requestMessageDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [RequestMessageDeleteResponseEnvelope]
type requestMessageDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestMessageDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RequestMessageDeleteResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    requestMessageDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// requestMessageDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [RequestMessageDeleteResponseEnvelopeErrors]
type requestMessageDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestMessageDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RequestMessageDeleteResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    requestMessageDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// requestMessageDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [RequestMessageDeleteResponseEnvelopeMessages]
type requestMessageDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestMessageDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestMessageDeleteResponseEnvelopeSuccess bool

const (
	RequestMessageDeleteResponseEnvelopeSuccessTrue RequestMessageDeleteResponseEnvelopeSuccess = true
)

type RequestMessageGetParams struct {
	// Page number of results
	Page param.Field[int64] `json:"page,required"`
	// Number of results per page
	PerPage param.Field[int64] `json:"per_page,required"`
	// Retrieve messages created after this time
	After param.Field[time.Time] `json:"after" format:"date-time"`
	// Retrieve messages created before this time
	Before param.Field[time.Time] `json:"before" format:"date-time"`
	// Field to sort results by
	SortBy param.Field[string] `json:"sort_by"`
	// Sort order (asc or desc)
	SortOrder param.Field[RequestMessageGetParamsSortOrder] `json:"sort_order"`
}

func (r RequestMessageGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Sort order (asc or desc)
type RequestMessageGetParamsSortOrder string

const (
	RequestMessageGetParamsSortOrderAsc  RequestMessageGetParamsSortOrder = "asc"
	RequestMessageGetParamsSortOrderDesc RequestMessageGetParamsSortOrder = "desc"
)

type RequestMessageGetResponseEnvelope struct {
	Errors   []RequestMessageGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RequestMessageGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []RequestMessageGetResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success RequestMessageGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    requestMessageGetResponseEnvelopeJSON    `json:"-"`
}

// requestMessageGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestMessageGetResponseEnvelope]
type requestMessageGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestMessageGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RequestMessageGetResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    requestMessageGetResponseEnvelopeErrorsJSON `json:"-"`
}

// requestMessageGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RequestMessageGetResponseEnvelopeErrors]
type requestMessageGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestMessageGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RequestMessageGetResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    requestMessageGetResponseEnvelopeMessagesJSON `json:"-"`
}

// requestMessageGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RequestMessageGetResponseEnvelopeMessages]
type requestMessageGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestMessageGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestMessageGetResponseEnvelopeSuccess bool

const (
	RequestMessageGetResponseEnvelopeSuccessTrue RequestMessageGetResponseEnvelopeSuccess = true
)
