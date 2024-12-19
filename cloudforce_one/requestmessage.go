// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// RequestMessageService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRequestMessageService] method instead.
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

// Create a New Request Message
func (r *RequestMessageService) New(ctx context.Context, accountIdentifier string, requestIdentifier string, body RequestMessageNewParams, opts ...option.RequestOption) (res *Message, err error) {
	var env RequestMessageNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if requestIdentifier == "" {
		err = errors.New("missing required request_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/message/new", accountIdentifier, requestIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a Request Message
func (r *RequestMessageService) Update(ctx context.Context, accountIdentifier string, requestIdentifier string, messageIdentifer int64, body RequestMessageUpdateParams, opts ...option.RequestOption) (res *Message, err error) {
	var env RequestMessageUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if requestIdentifier == "" {
		err = errors.New("missing required request_identifier parameter")
		return
	}
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
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if requestIdentifier == "" {
		err = errors.New("missing required request_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/message/%v", accountIdentifier, requestIdentifier, messageIdentifer)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// List Request Messages
func (r *RequestMessageService) Get(ctx context.Context, accountIdentifier string, requestIdentifier string, body RequestMessageGetParams, opts ...option.RequestOption) (res *[]Message, err error) {
	var env RequestMessageGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if requestIdentifier == "" {
		err = errors.New("missing required request_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/message", accountIdentifier, requestIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Message struct {
	// Message ID
	ID int64 `json:"id,required"`
	// Author of message
	Author string `json:"author,required"`
	// Content of message
	Content string `json:"content,required"`
	// Whether the message is a follow-on request
	IsFollowOnRequest bool `json:"is_follow_on_request,required"`
	// Message last updated time
	Updated time.Time `json:"updated,required" format:"date-time"`
	// Message creation time
	Created time.Time   `json:"created" format:"date-time"`
	JSON    messageJSON `json:"-"`
}

// messageJSON contains the JSON metadata for the struct [Message]
type messageJSON struct {
	ID                apijson.Field
	Author            apijson.Field
	Content           apijson.Field
	IsFollowOnRequest apijson.Field
	Updated           apijson.Field
	Created           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Message) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageJSON) RawJSON() string {
	return r.raw
}

type RequestMessageDeleteResponse struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RequestMessageDeleteResponseSuccess `json:"success,required"`
	JSON    requestMessageDeleteResponseJSON    `json:"-"`
}

// requestMessageDeleteResponseJSON contains the JSON metadata for the struct
// [RequestMessageDeleteResponse]
type requestMessageDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestMessageDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestMessageDeleteResponseSuccess bool

const (
	RequestMessageDeleteResponseSuccessTrue RequestMessageDeleteResponseSuccess = true
)

func (r RequestMessageDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case RequestMessageDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type RequestMessageNewParams struct {
	// Content of message
	Content param.Field[string] `json:"content"`
}

func (r RequestMessageNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RequestMessageNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RequestMessageNewResponseEnvelopeSuccess `json:"success,required"`
	Result  Message                                  `json:"result"`
	JSON    requestMessageNewResponseEnvelopeJSON    `json:"-"`
}

// requestMessageNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestMessageNewResponseEnvelope]
type requestMessageNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestMessageNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestMessageNewResponseEnvelopeSuccess bool

const (
	RequestMessageNewResponseEnvelopeSuccessTrue RequestMessageNewResponseEnvelopeSuccess = true
)

func (r RequestMessageNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RequestMessageNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RequestMessageUpdateParams struct {
	// Content of message
	Content param.Field[string] `json:"content"`
}

func (r RequestMessageUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RequestMessageUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RequestMessageUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  Message                                     `json:"result"`
	JSON    requestMessageUpdateResponseEnvelopeJSON    `json:"-"`
}

// requestMessageUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [RequestMessageUpdateResponseEnvelope]
type requestMessageUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestMessageUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestMessageUpdateResponseEnvelopeSuccess bool

const (
	RequestMessageUpdateResponseEnvelopeSuccessTrue RequestMessageUpdateResponseEnvelopeSuccess = true
)

func (r RequestMessageUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RequestMessageUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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

func (r RequestMessageGetParamsSortOrder) IsKnown() bool {
	switch r {
	case RequestMessageGetParamsSortOrderAsc, RequestMessageGetParamsSortOrderDesc:
		return true
	}
	return false
}

type RequestMessageGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RequestMessageGetResponseEnvelopeSuccess `json:"success,required"`
	Result  []Message                                `json:"result"`
	JSON    requestMessageGetResponseEnvelopeJSON    `json:"-"`
}

// requestMessageGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestMessageGetResponseEnvelope]
type requestMessageGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestMessageGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestMessageGetResponseEnvelopeSuccess bool

const (
	RequestMessageGetResponseEnvelopeSuccessTrue RequestMessageGetResponseEnvelopeSuccess = true
)

func (r RequestMessageGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RequestMessageGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
