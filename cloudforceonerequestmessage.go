// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// CloudforceOneRequestMessageService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewCloudforceOneRequestMessageService] method instead.
type CloudforceOneRequestMessageService struct {
	Options []option.RequestOption
}

// NewCloudforceOneRequestMessageService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudforceOneRequestMessageService(opts ...option.RequestOption) (r *CloudforceOneRequestMessageService) {
	r = &CloudforceOneRequestMessageService{}
	r.Options = opts
	return
}

// Creating a request adds the request into the Cloudforce One queue for analysis.
// In addition to the content, a short title, type, priority, and releasability
// should be provided. If one is not provided a default will be assigned.
func (r *CloudforceOneRequestMessageService) New(ctx context.Context, accountIdentifier string, requestIdentifier string, body CloudforceOneRequestMessageNewParams, opts ...option.RequestOption) (res *CloudforceOneRequestMessageItem, err error) {
	opts = append(r.Options[:], opts...)
	var env CloudforceOneRequestMessageNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/message/new", accountIdentifier, requestIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a Request Message
func (r *CloudforceOneRequestMessageService) Update(ctx context.Context, accountIdentifier string, requestIdentifier string, messageIdentifer int64, body CloudforceOneRequestMessageUpdateParams, opts ...option.RequestOption) (res *CloudforceOneRequestMessageItem, err error) {
	opts = append(r.Options[:], opts...)
	var env CloudforceOneRequestMessageUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/message/%v", accountIdentifier, requestIdentifier, messageIdentifer)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a Request Message
func (r *CloudforceOneRequestMessageService) Delete(ctx context.Context, accountIdentifier string, requestIdentifier string, messageIdentifer int64, opts ...option.RequestOption) (res *CloudforceOneRequestMessageDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CloudforceOneRequestMessageDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/message/%v", accountIdentifier, requestIdentifier, messageIdentifer)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Request Messages
func (r *CloudforceOneRequestMessageService) Get(ctx context.Context, accountIdentifier string, requestIdentifier string, body CloudforceOneRequestMessageGetParams, opts ...option.RequestOption) (res *[]CloudforceOneRequestMessageItem, err error) {
	opts = append(r.Options[:], opts...)
	var env CloudforceOneRequestMessageGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/message", accountIdentifier, requestIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CloudforceOneRequestMessageItem struct {
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
	Created time.Time                           `json:"created" format:"date-time"`
	JSON    cloudforceOneRequestMessageItemJSON `json:"-"`
}

// cloudforceOneRequestMessageItemJSON contains the JSON metadata for the struct
// [CloudforceOneRequestMessageItem]
type cloudforceOneRequestMessageItemJSON struct {
	ID                apijson.Field
	Author            apijson.Field
	Content           apijson.Field
	IsFollowOnRequest apijson.Field
	Updated           apijson.Field
	Created           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudforceOneRequestMessageItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [CloudforceOneRequestMessageDeleteResponseUnknown],
// [CloudforceOneRequestMessageDeleteResponseArray] or [shared.UnionString].
type CloudforceOneRequestMessageDeleteResponse interface {
	ImplementsCloudforceOneRequestMessageDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudforceOneRequestMessageDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudforceOneRequestMessageDeleteResponseArray []interface{}

func (r CloudforceOneRequestMessageDeleteResponseArray) ImplementsCloudforceOneRequestMessageDeleteResponse() {
}

type CloudforceOneRequestMessageNewParams struct {
	// Content of message
	Content param.Field[string] `json:"content"`
}

func (r CloudforceOneRequestMessageNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudforceOneRequestMessageNewResponseEnvelope struct {
	Errors   []CloudforceOneRequestMessageNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CloudforceOneRequestMessageNewResponseEnvelopeMessages `json:"messages,required"`
	Result   CloudforceOneRequestMessageItem                          `json:"result,required"`
	// Whether the API call was successful
	Success CloudforceOneRequestMessageNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    cloudforceOneRequestMessageNewResponseEnvelopeJSON    `json:"-"`
}

// cloudforceOneRequestMessageNewResponseEnvelopeJSON contains the JSON metadata
// for the struct [CloudforceOneRequestMessageNewResponseEnvelope]
type cloudforceOneRequestMessageNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestMessageNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestMessageNewResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    cloudforceOneRequestMessageNewResponseEnvelopeErrorsJSON `json:"-"`
}

// cloudforceOneRequestMessageNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [CloudforceOneRequestMessageNewResponseEnvelopeErrors]
type cloudforceOneRequestMessageNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestMessageNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestMessageNewResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    cloudforceOneRequestMessageNewResponseEnvelopeMessagesJSON `json:"-"`
}

// cloudforceOneRequestMessageNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [CloudforceOneRequestMessageNewResponseEnvelopeMessages]
type cloudforceOneRequestMessageNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestMessageNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CloudforceOneRequestMessageNewResponseEnvelopeSuccess bool

const (
	CloudforceOneRequestMessageNewResponseEnvelopeSuccessTrue CloudforceOneRequestMessageNewResponseEnvelopeSuccess = true
)

type CloudforceOneRequestMessageUpdateParams struct {
	// Request content
	Content param.Field[string] `json:"content"`
	// Priority for analyzing the request
	Priority param.Field[string] `json:"priority"`
	// Requested information from request
	RequestType param.Field[string] `json:"request_type"`
	// Brief description of the request
	Summary param.Field[string] `json:"summary"`
	// The CISA defined Traffic Light Protocol (TLP)
	Tlp param.Field[CloudforceOneRequestMessageUpdateParamsTlp] `json:"tlp"`
}

func (r CloudforceOneRequestMessageUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The CISA defined Traffic Light Protocol (TLP)
type CloudforceOneRequestMessageUpdateParamsTlp string

const (
	CloudforceOneRequestMessageUpdateParamsTlpClear       CloudforceOneRequestMessageUpdateParamsTlp = "clear"
	CloudforceOneRequestMessageUpdateParamsTlpAmber       CloudforceOneRequestMessageUpdateParamsTlp = "amber"
	CloudforceOneRequestMessageUpdateParamsTlpAmberStrict CloudforceOneRequestMessageUpdateParamsTlp = "amber-strict"
	CloudforceOneRequestMessageUpdateParamsTlpGreen       CloudforceOneRequestMessageUpdateParamsTlp = "green"
	CloudforceOneRequestMessageUpdateParamsTlpRed         CloudforceOneRequestMessageUpdateParamsTlp = "red"
)

type CloudforceOneRequestMessageUpdateResponseEnvelope struct {
	Errors   []CloudforceOneRequestMessageUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CloudforceOneRequestMessageUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   CloudforceOneRequestMessageItem                             `json:"result,required"`
	// Whether the API call was successful
	Success CloudforceOneRequestMessageUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    cloudforceOneRequestMessageUpdateResponseEnvelopeJSON    `json:"-"`
}

// cloudforceOneRequestMessageUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [CloudforceOneRequestMessageUpdateResponseEnvelope]
type cloudforceOneRequestMessageUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestMessageUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestMessageUpdateResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    cloudforceOneRequestMessageUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// cloudforceOneRequestMessageUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [CloudforceOneRequestMessageUpdateResponseEnvelopeErrors]
type cloudforceOneRequestMessageUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestMessageUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestMessageUpdateResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    cloudforceOneRequestMessageUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// cloudforceOneRequestMessageUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [CloudforceOneRequestMessageUpdateResponseEnvelopeMessages]
type cloudforceOneRequestMessageUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestMessageUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CloudforceOneRequestMessageUpdateResponseEnvelopeSuccess bool

const (
	CloudforceOneRequestMessageUpdateResponseEnvelopeSuccessTrue CloudforceOneRequestMessageUpdateResponseEnvelopeSuccess = true
)

type CloudforceOneRequestMessageDeleteResponseEnvelope struct {
	Errors   []CloudforceOneRequestMessageDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CloudforceOneRequestMessageDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   CloudforceOneRequestMessageDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CloudforceOneRequestMessageDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    cloudforceOneRequestMessageDeleteResponseEnvelopeJSON    `json:"-"`
}

// cloudforceOneRequestMessageDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [CloudforceOneRequestMessageDeleteResponseEnvelope]
type cloudforceOneRequestMessageDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestMessageDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestMessageDeleteResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    cloudforceOneRequestMessageDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// cloudforceOneRequestMessageDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [CloudforceOneRequestMessageDeleteResponseEnvelopeErrors]
type cloudforceOneRequestMessageDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestMessageDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestMessageDeleteResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    cloudforceOneRequestMessageDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// cloudforceOneRequestMessageDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [CloudforceOneRequestMessageDeleteResponseEnvelopeMessages]
type cloudforceOneRequestMessageDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestMessageDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CloudforceOneRequestMessageDeleteResponseEnvelopeSuccess bool

const (
	CloudforceOneRequestMessageDeleteResponseEnvelopeSuccessTrue CloudforceOneRequestMessageDeleteResponseEnvelopeSuccess = true
)

type CloudforceOneRequestMessageGetParams struct {
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
	SortOrder param.Field[CloudforceOneRequestMessageGetParamsSortOrder] `json:"sort_order"`
}

func (r CloudforceOneRequestMessageGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Sort order (asc or desc)
type CloudforceOneRequestMessageGetParamsSortOrder string

const (
	CloudforceOneRequestMessageGetParamsSortOrderAsc  CloudforceOneRequestMessageGetParamsSortOrder = "asc"
	CloudforceOneRequestMessageGetParamsSortOrderDesc CloudforceOneRequestMessageGetParamsSortOrder = "desc"
)

type CloudforceOneRequestMessageGetResponseEnvelope struct {
	Errors   []CloudforceOneRequestMessageGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CloudforceOneRequestMessageGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []CloudforceOneRequestMessageItem                        `json:"result,required"`
	// Whether the API call was successful
	Success CloudforceOneRequestMessageGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    cloudforceOneRequestMessageGetResponseEnvelopeJSON    `json:"-"`
}

// cloudforceOneRequestMessageGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [CloudforceOneRequestMessageGetResponseEnvelope]
type cloudforceOneRequestMessageGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestMessageGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestMessageGetResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    cloudforceOneRequestMessageGetResponseEnvelopeErrorsJSON `json:"-"`
}

// cloudforceOneRequestMessageGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [CloudforceOneRequestMessageGetResponseEnvelopeErrors]
type cloudforceOneRequestMessageGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestMessageGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestMessageGetResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    cloudforceOneRequestMessageGetResponseEnvelopeMessagesJSON `json:"-"`
}

// cloudforceOneRequestMessageGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [CloudforceOneRequestMessageGetResponseEnvelopeMessages]
type cloudforceOneRequestMessageGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestMessageGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CloudforceOneRequestMessageGetResponseEnvelopeSuccess bool

const (
	CloudforceOneRequestMessageGetResponseEnvelopeSuccessTrue CloudforceOneRequestMessageGetResponseEnvelopeSuccess = true
)
