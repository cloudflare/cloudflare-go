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

// RequestPriorityService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRequestPriorityService] method
// instead.
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

// Create a New Priority Requirement
func (r *RequestPriorityService) New(ctx context.Context, accountIdentifier string, body RequestPriorityNewParams, opts ...option.RequestOption) (res *CloudforceOnePriorityItem, err error) {
	opts = append(r.Options[:], opts...)
	var env RequestPriorityNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/new", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a Priority Intelligence Requirement
func (r *RequestPriorityService) Update(ctx context.Context, accountIdentifier string, priorityIdentifer string, body RequestPriorityUpdateParams, opts ...option.RequestOption) (res *CloudforceOneRequestItem, err error) {
	opts = append(r.Options[:], opts...)
	var env RequestPriorityUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/%s", accountIdentifier, priorityIdentifer)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a Priority Intelligence Report
func (r *RequestPriorityService) Delete(ctx context.Context, accountIdentifier string, priorityIdentifer string, opts ...option.RequestOption) (res *RequestPriorityDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RequestPriorityDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/%s", accountIdentifier, priorityIdentifer)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a Priority Intelligence Requirement
func (r *RequestPriorityService) Get(ctx context.Context, accountIdentifier string, priorityIdentifer string, opts ...option.RequestOption) (res *CloudforceOneRequestItem, err error) {
	opts = append(r.Options[:], opts...)
	var env RequestPriorityGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/%s", accountIdentifier, priorityIdentifer)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Priority Intelligence Requirement Quota
func (r *RequestPriorityService) Quota(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *CloudforceOneQuota, err error) {
	opts = append(r.Options[:], opts...)
	var env RequestPriorityQuotaResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/quota", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CloudforceOnePriorityItem struct {
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
	Tlp CloudforceOnePriorityItemTlp `json:"tlp,required"`
	// Priority last updated time
	Updated time.Time                     `json:"updated,required" format:"date-time"`
	JSON    cloudforceOnePriorityItemJSON `json:"-"`
}

// cloudforceOnePriorityItemJSON contains the JSON metadata for the struct
// [CloudforceOnePriorityItem]
type cloudforceOnePriorityItemJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Labels      apijson.Field
	Priority    apijson.Field
	Requirement apijson.Field
	Tlp         apijson.Field
	Updated     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOnePriorityItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudforceOnePriorityItemJSON) RawJSON() string {
	return r.raw
}

// The CISA defined Traffic Light Protocol (TLP)
type CloudforceOnePriorityItemTlp string

const (
	CloudforceOnePriorityItemTlpClear       CloudforceOnePriorityItemTlp = "clear"
	CloudforceOnePriorityItemTlpAmber       CloudforceOnePriorityItemTlp = "amber"
	CloudforceOnePriorityItemTlpAmberStrict CloudforceOnePriorityItemTlp = "amber-strict"
	CloudforceOnePriorityItemTlpGreen       CloudforceOnePriorityItemTlp = "green"
	CloudforceOnePriorityItemTlpRed         CloudforceOnePriorityItemTlp = "red"
)

// Union satisfied by [cloudforce_one.RequestPriorityDeleteResponseUnknown],
// [cloudforce_one.RequestPriorityDeleteResponseArray] or [shared.UnionString].
type RequestPriorityDeleteResponse interface {
	ImplementsCloudforceOneRequestPriorityDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RequestPriorityDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RequestPriorityDeleteResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type RequestPriorityDeleteResponseArray []interface{}

func (r RequestPriorityDeleteResponseArray) ImplementsCloudforceOneRequestPriorityDeleteResponse() {}

type RequestPriorityNewParams struct {
	// List of labels
	Labels param.Field[[]string] `json:"labels,required"`
	// Priority
	Priority param.Field[int64] `json:"priority,required"`
	// Requirement
	Requirement param.Field[string] `json:"requirement,required"`
	// The CISA defined Traffic Light Protocol (TLP)
	Tlp param.Field[RequestPriorityNewParamsTlp] `json:"tlp,required"`
}

func (r RequestPriorityNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The CISA defined Traffic Light Protocol (TLP)
type RequestPriorityNewParamsTlp string

const (
	RequestPriorityNewParamsTlpClear       RequestPriorityNewParamsTlp = "clear"
	RequestPriorityNewParamsTlpAmber       RequestPriorityNewParamsTlp = "amber"
	RequestPriorityNewParamsTlpAmberStrict RequestPriorityNewParamsTlp = "amber-strict"
	RequestPriorityNewParamsTlpGreen       RequestPriorityNewParamsTlp = "green"
	RequestPriorityNewParamsTlpRed         RequestPriorityNewParamsTlp = "red"
)

type RequestPriorityNewResponseEnvelope struct {
	Errors   []RequestPriorityNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RequestPriorityNewResponseEnvelopeMessages `json:"messages,required"`
	Result   CloudforceOnePriorityItem                    `json:"result,required"`
	// Whether the API call was successful
	Success RequestPriorityNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    requestPriorityNewResponseEnvelopeJSON    `json:"-"`
}

// requestPriorityNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestPriorityNewResponseEnvelope]
type requestPriorityNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RequestPriorityNewResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    requestPriorityNewResponseEnvelopeErrorsJSON `json:"-"`
}

// requestPriorityNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RequestPriorityNewResponseEnvelopeErrors]
type requestPriorityNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RequestPriorityNewResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    requestPriorityNewResponseEnvelopeMessagesJSON `json:"-"`
}

// requestPriorityNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [RequestPriorityNewResponseEnvelopeMessages]
type requestPriorityNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestPriorityNewResponseEnvelopeSuccess bool

const (
	RequestPriorityNewResponseEnvelopeSuccessTrue RequestPriorityNewResponseEnvelopeSuccess = true
)

type RequestPriorityUpdateParams struct {
	// List of labels
	Labels param.Field[[]string] `json:"labels,required"`
	// Priority
	Priority param.Field[int64] `json:"priority,required"`
	// Requirement
	Requirement param.Field[string] `json:"requirement,required"`
	// The CISA defined Traffic Light Protocol (TLP)
	Tlp param.Field[RequestPriorityUpdateParamsTlp] `json:"tlp,required"`
}

func (r RequestPriorityUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The CISA defined Traffic Light Protocol (TLP)
type RequestPriorityUpdateParamsTlp string

const (
	RequestPriorityUpdateParamsTlpClear       RequestPriorityUpdateParamsTlp = "clear"
	RequestPriorityUpdateParamsTlpAmber       RequestPriorityUpdateParamsTlp = "amber"
	RequestPriorityUpdateParamsTlpAmberStrict RequestPriorityUpdateParamsTlp = "amber-strict"
	RequestPriorityUpdateParamsTlpGreen       RequestPriorityUpdateParamsTlp = "green"
	RequestPriorityUpdateParamsTlpRed         RequestPriorityUpdateParamsTlp = "red"
)

type RequestPriorityUpdateResponseEnvelope struct {
	Errors   []RequestPriorityUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RequestPriorityUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   CloudforceOneRequestItem                        `json:"result,required"`
	// Whether the API call was successful
	Success RequestPriorityUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    requestPriorityUpdateResponseEnvelopeJSON    `json:"-"`
}

// requestPriorityUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [RequestPriorityUpdateResponseEnvelope]
type requestPriorityUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RequestPriorityUpdateResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    requestPriorityUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// requestPriorityUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [RequestPriorityUpdateResponseEnvelopeErrors]
type requestPriorityUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RequestPriorityUpdateResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    requestPriorityUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// requestPriorityUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [RequestPriorityUpdateResponseEnvelopeMessages]
type requestPriorityUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestPriorityUpdateResponseEnvelopeSuccess bool

const (
	RequestPriorityUpdateResponseEnvelopeSuccessTrue RequestPriorityUpdateResponseEnvelopeSuccess = true
)

type RequestPriorityDeleteResponseEnvelope struct {
	Errors   []RequestPriorityDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RequestPriorityDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   RequestPriorityDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success RequestPriorityDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    requestPriorityDeleteResponseEnvelopeJSON    `json:"-"`
}

// requestPriorityDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [RequestPriorityDeleteResponseEnvelope]
type requestPriorityDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RequestPriorityDeleteResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    requestPriorityDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// requestPriorityDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [RequestPriorityDeleteResponseEnvelopeErrors]
type requestPriorityDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RequestPriorityDeleteResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    requestPriorityDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// requestPriorityDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [RequestPriorityDeleteResponseEnvelopeMessages]
type requestPriorityDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestPriorityDeleteResponseEnvelopeSuccess bool

const (
	RequestPriorityDeleteResponseEnvelopeSuccessTrue RequestPriorityDeleteResponseEnvelopeSuccess = true
)

type RequestPriorityGetResponseEnvelope struct {
	Errors   []RequestPriorityGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RequestPriorityGetResponseEnvelopeMessages `json:"messages,required"`
	Result   CloudforceOneRequestItem                     `json:"result,required"`
	// Whether the API call was successful
	Success RequestPriorityGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    requestPriorityGetResponseEnvelopeJSON    `json:"-"`
}

// requestPriorityGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestPriorityGetResponseEnvelope]
type requestPriorityGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RequestPriorityGetResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    requestPriorityGetResponseEnvelopeErrorsJSON `json:"-"`
}

// requestPriorityGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RequestPriorityGetResponseEnvelopeErrors]
type requestPriorityGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RequestPriorityGetResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    requestPriorityGetResponseEnvelopeMessagesJSON `json:"-"`
}

// requestPriorityGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [RequestPriorityGetResponseEnvelopeMessages]
type requestPriorityGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestPriorityGetResponseEnvelopeSuccess bool

const (
	RequestPriorityGetResponseEnvelopeSuccessTrue RequestPriorityGetResponseEnvelopeSuccess = true
)

type RequestPriorityQuotaResponseEnvelope struct {
	Errors   []RequestPriorityQuotaResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RequestPriorityQuotaResponseEnvelopeMessages `json:"messages,required"`
	Result   CloudforceOneQuota                             `json:"result,required"`
	// Whether the API call was successful
	Success RequestPriorityQuotaResponseEnvelopeSuccess `json:"success,required"`
	JSON    requestPriorityQuotaResponseEnvelopeJSON    `json:"-"`
}

// requestPriorityQuotaResponseEnvelopeJSON contains the JSON metadata for the
// struct [RequestPriorityQuotaResponseEnvelope]
type requestPriorityQuotaResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityQuotaResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityQuotaResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RequestPriorityQuotaResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    requestPriorityQuotaResponseEnvelopeErrorsJSON `json:"-"`
}

// requestPriorityQuotaResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [RequestPriorityQuotaResponseEnvelopeErrors]
type requestPriorityQuotaResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityQuotaResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityQuotaResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RequestPriorityQuotaResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    requestPriorityQuotaResponseEnvelopeMessagesJSON `json:"-"`
}

// requestPriorityQuotaResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [RequestPriorityQuotaResponseEnvelopeMessages]
type requestPriorityQuotaResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityQuotaResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityQuotaResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestPriorityQuotaResponseEnvelopeSuccess bool

const (
	RequestPriorityQuotaResponseEnvelopeSuccessTrue RequestPriorityQuotaResponseEnvelopeSuccess = true
)
