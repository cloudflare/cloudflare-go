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

// CloudforceOneRequestPriorityService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewCloudforceOneRequestPriorityService] method instead.
type CloudforceOneRequestPriorityService struct {
	Options []option.RequestOption
}

// NewCloudforceOneRequestPriorityService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudforceOneRequestPriorityService(opts ...option.RequestOption) (r *CloudforceOneRequestPriorityService) {
	r = &CloudforceOneRequestPriorityService{}
	r.Options = opts
	return
}

// Create a New Priority Requirement
func (r *CloudforceOneRequestPriorityService) New(ctx context.Context, accountIdentifier string, body CloudforceOneRequestPriorityNewParams, opts ...option.RequestOption) (res *CloudforceOneRequestPriorityNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CloudforceOneRequestPriorityNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/new", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a Priority Intelligence Requirement
func (r *CloudforceOneRequestPriorityService) Update(ctx context.Context, accountIdentifier string, priorityIdentifer string, body CloudforceOneRequestPriorityUpdateParams, opts ...option.RequestOption) (res *CloudforceOneRequestPriorityUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CloudforceOneRequestPriorityUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/%s", accountIdentifier, priorityIdentifer)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a Priority Intelligence Report
func (r *CloudforceOneRequestPriorityService) Delete(ctx context.Context, accountIdentifier string, priorityIdentifer string, opts ...option.RequestOption) (res *CloudforceOneRequestPriorityDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CloudforceOneRequestPriorityDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/%s", accountIdentifier, priorityIdentifer)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Priority Intelligence Requirements
func (r *CloudforceOneRequestPriorityService) DoSomethingUnknown(ctx context.Context, accountIdentifier string, body CloudforceOneRequestPriorityDoSomethingUnknownParams, opts ...option.RequestOption) (res *[]CloudforceOneRequestPriorityDoSomethingUnknownResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CloudforceOneRequestPriorityDoSomethingUnknownResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a Priority Intelligence Requirement
func (r *CloudforceOneRequestPriorityService) Get(ctx context.Context, accountIdentifier string, priorityIdentifer string, opts ...option.RequestOption) (res *CloudforceOneRequestPriorityGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CloudforceOneRequestPriorityGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/%s", accountIdentifier, priorityIdentifer)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Priority Intelligence Requirement Quota
func (r *CloudforceOneRequestPriorityService) Quota(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *CloudforceOneRequestPriorityQuotaResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CloudforceOneRequestPriorityQuotaResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/quota", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CloudforceOneRequestPriorityNewResponse struct {
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
	Tlp CloudforceOneRequestPriorityNewResponseTlp `json:"tlp,required"`
	// Priority last updated time
	Updated time.Time                                   `json:"updated,required" format:"date-time"`
	JSON    cloudforceOneRequestPriorityNewResponseJSON `json:"-"`
}

// cloudforceOneRequestPriorityNewResponseJSON contains the JSON metadata for the
// struct [CloudforceOneRequestPriorityNewResponse]
type cloudforceOneRequestPriorityNewResponseJSON struct {
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

func (r *CloudforceOneRequestPriorityNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The CISA defined Traffic Light Protocol (TLP)
type CloudforceOneRequestPriorityNewResponseTlp string

const (
	CloudforceOneRequestPriorityNewResponseTlpClear       CloudforceOneRequestPriorityNewResponseTlp = "clear"
	CloudforceOneRequestPriorityNewResponseTlpAmber       CloudforceOneRequestPriorityNewResponseTlp = "amber"
	CloudforceOneRequestPriorityNewResponseTlpAmberStrict CloudforceOneRequestPriorityNewResponseTlp = "amber-strict"
	CloudforceOneRequestPriorityNewResponseTlpGreen       CloudforceOneRequestPriorityNewResponseTlp = "green"
	CloudforceOneRequestPriorityNewResponseTlpRed         CloudforceOneRequestPriorityNewResponseTlp = "red"
)

type CloudforceOneRequestPriorityUpdateResponse struct {
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
	Tlp       CloudforceOneRequestPriorityUpdateResponseTlp `json:"tlp,required"`
	Updated   time.Time                                     `json:"updated,required" format:"date-time"`
	Completed time.Time                                     `json:"completed" format:"date-time"`
	// Tokens for the request messages
	MessageTokens int64 `json:"message_tokens"`
	// Readable Request ID
	ReadableID string `json:"readable_id"`
	// Request Status
	Status CloudforceOneRequestPriorityUpdateResponseStatus `json:"status"`
	// Tokens for the request
	Tokens int64                                          `json:"tokens"`
	JSON   cloudforceOneRequestPriorityUpdateResponseJSON `json:"-"`
}

// cloudforceOneRequestPriorityUpdateResponseJSON contains the JSON metadata for
// the struct [CloudforceOneRequestPriorityUpdateResponse]
type cloudforceOneRequestPriorityUpdateResponseJSON struct {
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

func (r *CloudforceOneRequestPriorityUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The CISA defined Traffic Light Protocol (TLP)
type CloudforceOneRequestPriorityUpdateResponseTlp string

const (
	CloudforceOneRequestPriorityUpdateResponseTlpClear       CloudforceOneRequestPriorityUpdateResponseTlp = "clear"
	CloudforceOneRequestPriorityUpdateResponseTlpAmber       CloudforceOneRequestPriorityUpdateResponseTlp = "amber"
	CloudforceOneRequestPriorityUpdateResponseTlpAmberStrict CloudforceOneRequestPriorityUpdateResponseTlp = "amber-strict"
	CloudforceOneRequestPriorityUpdateResponseTlpGreen       CloudforceOneRequestPriorityUpdateResponseTlp = "green"
	CloudforceOneRequestPriorityUpdateResponseTlpRed         CloudforceOneRequestPriorityUpdateResponseTlp = "red"
)

// Request Status
type CloudforceOneRequestPriorityUpdateResponseStatus string

const (
	CloudforceOneRequestPriorityUpdateResponseStatusOpen      CloudforceOneRequestPriorityUpdateResponseStatus = "open"
	CloudforceOneRequestPriorityUpdateResponseStatusAccepted  CloudforceOneRequestPriorityUpdateResponseStatus = "accepted"
	CloudforceOneRequestPriorityUpdateResponseStatusReported  CloudforceOneRequestPriorityUpdateResponseStatus = "reported"
	CloudforceOneRequestPriorityUpdateResponseStatusApproved  CloudforceOneRequestPriorityUpdateResponseStatus = "approved"
	CloudforceOneRequestPriorityUpdateResponseStatusCompleted CloudforceOneRequestPriorityUpdateResponseStatus = "completed"
	CloudforceOneRequestPriorityUpdateResponseStatusDeclined  CloudforceOneRequestPriorityUpdateResponseStatus = "declined"
)

// Union satisfied by [CloudforceOneRequestPriorityDeleteResponseUnknown],
// [CloudforceOneRequestPriorityDeleteResponseArray] or [shared.UnionString].
type CloudforceOneRequestPriorityDeleteResponse interface {
	ImplementsCloudforceOneRequestPriorityDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudforceOneRequestPriorityDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudforceOneRequestPriorityDeleteResponseArray []interface{}

func (r CloudforceOneRequestPriorityDeleteResponseArray) ImplementsCloudforceOneRequestPriorityDeleteResponse() {
}

type CloudforceOneRequestPriorityDoSomethingUnknownResponse struct {
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
	Tlp CloudforceOneRequestPriorityDoSomethingUnknownResponseTlp `json:"tlp,required"`
	// Priority last updated time
	Updated time.Time                                                  `json:"updated,required" format:"date-time"`
	JSON    cloudforceOneRequestPriorityDoSomethingUnknownResponseJSON `json:"-"`
}

// cloudforceOneRequestPriorityDoSomethingUnknownResponseJSON contains the JSON
// metadata for the struct [CloudforceOneRequestPriorityDoSomethingUnknownResponse]
type cloudforceOneRequestPriorityDoSomethingUnknownResponseJSON struct {
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

func (r *CloudforceOneRequestPriorityDoSomethingUnknownResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The CISA defined Traffic Light Protocol (TLP)
type CloudforceOneRequestPriorityDoSomethingUnknownResponseTlp string

const (
	CloudforceOneRequestPriorityDoSomethingUnknownResponseTlpClear       CloudforceOneRequestPriorityDoSomethingUnknownResponseTlp = "clear"
	CloudforceOneRequestPriorityDoSomethingUnknownResponseTlpAmber       CloudforceOneRequestPriorityDoSomethingUnknownResponseTlp = "amber"
	CloudforceOneRequestPriorityDoSomethingUnknownResponseTlpAmberStrict CloudforceOneRequestPriorityDoSomethingUnknownResponseTlp = "amber-strict"
	CloudforceOneRequestPriorityDoSomethingUnknownResponseTlpGreen       CloudforceOneRequestPriorityDoSomethingUnknownResponseTlp = "green"
	CloudforceOneRequestPriorityDoSomethingUnknownResponseTlpRed         CloudforceOneRequestPriorityDoSomethingUnknownResponseTlp = "red"
)

type CloudforceOneRequestPriorityGetResponse struct {
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
	Tlp       CloudforceOneRequestPriorityGetResponseTlp `json:"tlp,required"`
	Updated   time.Time                                  `json:"updated,required" format:"date-time"`
	Completed time.Time                                  `json:"completed" format:"date-time"`
	// Tokens for the request messages
	MessageTokens int64 `json:"message_tokens"`
	// Readable Request ID
	ReadableID string `json:"readable_id"`
	// Request Status
	Status CloudforceOneRequestPriorityGetResponseStatus `json:"status"`
	// Tokens for the request
	Tokens int64                                       `json:"tokens"`
	JSON   cloudforceOneRequestPriorityGetResponseJSON `json:"-"`
}

// cloudforceOneRequestPriorityGetResponseJSON contains the JSON metadata for the
// struct [CloudforceOneRequestPriorityGetResponse]
type cloudforceOneRequestPriorityGetResponseJSON struct {
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

func (r *CloudforceOneRequestPriorityGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The CISA defined Traffic Light Protocol (TLP)
type CloudforceOneRequestPriorityGetResponseTlp string

const (
	CloudforceOneRequestPriorityGetResponseTlpClear       CloudforceOneRequestPriorityGetResponseTlp = "clear"
	CloudforceOneRequestPriorityGetResponseTlpAmber       CloudforceOneRequestPriorityGetResponseTlp = "amber"
	CloudforceOneRequestPriorityGetResponseTlpAmberStrict CloudforceOneRequestPriorityGetResponseTlp = "amber-strict"
	CloudforceOneRequestPriorityGetResponseTlpGreen       CloudforceOneRequestPriorityGetResponseTlp = "green"
	CloudforceOneRequestPriorityGetResponseTlpRed         CloudforceOneRequestPriorityGetResponseTlp = "red"
)

// Request Status
type CloudforceOneRequestPriorityGetResponseStatus string

const (
	CloudforceOneRequestPriorityGetResponseStatusOpen      CloudforceOneRequestPriorityGetResponseStatus = "open"
	CloudforceOneRequestPriorityGetResponseStatusAccepted  CloudforceOneRequestPriorityGetResponseStatus = "accepted"
	CloudforceOneRequestPriorityGetResponseStatusReported  CloudforceOneRequestPriorityGetResponseStatus = "reported"
	CloudforceOneRequestPriorityGetResponseStatusApproved  CloudforceOneRequestPriorityGetResponseStatus = "approved"
	CloudforceOneRequestPriorityGetResponseStatusCompleted CloudforceOneRequestPriorityGetResponseStatus = "completed"
	CloudforceOneRequestPriorityGetResponseStatusDeclined  CloudforceOneRequestPriorityGetResponseStatus = "declined"
)

type CloudforceOneRequestPriorityQuotaResponse struct {
	// Anniversary date is when annual quota limit is refresh
	AnniversaryDate time.Time `json:"anniversary_date" format:"date-time"`
	// Quater anniversary date is when quota limit is refreshed each quarter
	QuarterAnniversaryDate time.Time `json:"quarter_anniversary_date" format:"date-time"`
	// Tokens for the quarter
	Quota int64 `json:"quota"`
	// Tokens remaining for the quarter
	Remaining int64                                         `json:"remaining"`
	JSON      cloudforceOneRequestPriorityQuotaResponseJSON `json:"-"`
}

// cloudforceOneRequestPriorityQuotaResponseJSON contains the JSON metadata for the
// struct [CloudforceOneRequestPriorityQuotaResponse]
type cloudforceOneRequestPriorityQuotaResponseJSON struct {
	AnniversaryDate        apijson.Field
	QuarterAnniversaryDate apijson.Field
	Quota                  apijson.Field
	Remaining              apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityQuotaResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestPriorityNewParams struct {
	// List of labels
	Labels param.Field[[]string] `json:"labels,required"`
	// Priority
	Priority param.Field[int64] `json:"priority,required"`
	// Requirement
	Requirement param.Field[string] `json:"requirement,required"`
	// The CISA defined Traffic Light Protocol (TLP)
	Tlp param.Field[CloudforceOneRequestPriorityNewParamsTlp] `json:"tlp,required"`
}

func (r CloudforceOneRequestPriorityNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The CISA defined Traffic Light Protocol (TLP)
type CloudforceOneRequestPriorityNewParamsTlp string

const (
	CloudforceOneRequestPriorityNewParamsTlpClear       CloudforceOneRequestPriorityNewParamsTlp = "clear"
	CloudforceOneRequestPriorityNewParamsTlpAmber       CloudforceOneRequestPriorityNewParamsTlp = "amber"
	CloudforceOneRequestPriorityNewParamsTlpAmberStrict CloudforceOneRequestPriorityNewParamsTlp = "amber-strict"
	CloudforceOneRequestPriorityNewParamsTlpGreen       CloudforceOneRequestPriorityNewParamsTlp = "green"
	CloudforceOneRequestPriorityNewParamsTlpRed         CloudforceOneRequestPriorityNewParamsTlp = "red"
)

type CloudforceOneRequestPriorityNewResponseEnvelope struct {
	Errors   []CloudforceOneRequestPriorityNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CloudforceOneRequestPriorityNewResponseEnvelopeMessages `json:"messages,required"`
	Result   CloudforceOneRequestPriorityNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CloudforceOneRequestPriorityNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    cloudforceOneRequestPriorityNewResponseEnvelopeJSON    `json:"-"`
}

// cloudforceOneRequestPriorityNewResponseEnvelopeJSON contains the JSON metadata
// for the struct [CloudforceOneRequestPriorityNewResponseEnvelope]
type cloudforceOneRequestPriorityNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestPriorityNewResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    cloudforceOneRequestPriorityNewResponseEnvelopeErrorsJSON `json:"-"`
}

// cloudforceOneRequestPriorityNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [CloudforceOneRequestPriorityNewResponseEnvelopeErrors]
type cloudforceOneRequestPriorityNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestPriorityNewResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    cloudforceOneRequestPriorityNewResponseEnvelopeMessagesJSON `json:"-"`
}

// cloudforceOneRequestPriorityNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [CloudforceOneRequestPriorityNewResponseEnvelopeMessages]
type cloudforceOneRequestPriorityNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CloudforceOneRequestPriorityNewResponseEnvelopeSuccess bool

const (
	CloudforceOneRequestPriorityNewResponseEnvelopeSuccessTrue CloudforceOneRequestPriorityNewResponseEnvelopeSuccess = true
)

type CloudforceOneRequestPriorityUpdateParams struct {
	// List of labels
	Labels param.Field[[]string] `json:"labels,required"`
	// Priority
	Priority param.Field[int64] `json:"priority,required"`
	// Requirement
	Requirement param.Field[string] `json:"requirement,required"`
	// The CISA defined Traffic Light Protocol (TLP)
	Tlp param.Field[CloudforceOneRequestPriorityUpdateParamsTlp] `json:"tlp,required"`
}

func (r CloudforceOneRequestPriorityUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The CISA defined Traffic Light Protocol (TLP)
type CloudforceOneRequestPriorityUpdateParamsTlp string

const (
	CloudforceOneRequestPriorityUpdateParamsTlpClear       CloudforceOneRequestPriorityUpdateParamsTlp = "clear"
	CloudforceOneRequestPriorityUpdateParamsTlpAmber       CloudforceOneRequestPriorityUpdateParamsTlp = "amber"
	CloudforceOneRequestPriorityUpdateParamsTlpAmberStrict CloudforceOneRequestPriorityUpdateParamsTlp = "amber-strict"
	CloudforceOneRequestPriorityUpdateParamsTlpGreen       CloudforceOneRequestPriorityUpdateParamsTlp = "green"
	CloudforceOneRequestPriorityUpdateParamsTlpRed         CloudforceOneRequestPriorityUpdateParamsTlp = "red"
)

type CloudforceOneRequestPriorityUpdateResponseEnvelope struct {
	Errors   []CloudforceOneRequestPriorityUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CloudforceOneRequestPriorityUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   CloudforceOneRequestPriorityUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CloudforceOneRequestPriorityUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    cloudforceOneRequestPriorityUpdateResponseEnvelopeJSON    `json:"-"`
}

// cloudforceOneRequestPriorityUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [CloudforceOneRequestPriorityUpdateResponseEnvelope]
type cloudforceOneRequestPriorityUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestPriorityUpdateResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    cloudforceOneRequestPriorityUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// cloudforceOneRequestPriorityUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [CloudforceOneRequestPriorityUpdateResponseEnvelopeErrors]
type cloudforceOneRequestPriorityUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestPriorityUpdateResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    cloudforceOneRequestPriorityUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// cloudforceOneRequestPriorityUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [CloudforceOneRequestPriorityUpdateResponseEnvelopeMessages]
type cloudforceOneRequestPriorityUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CloudforceOneRequestPriorityUpdateResponseEnvelopeSuccess bool

const (
	CloudforceOneRequestPriorityUpdateResponseEnvelopeSuccessTrue CloudforceOneRequestPriorityUpdateResponseEnvelopeSuccess = true
)

type CloudforceOneRequestPriorityDeleteResponseEnvelope struct {
	Errors   []CloudforceOneRequestPriorityDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CloudforceOneRequestPriorityDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   CloudforceOneRequestPriorityDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CloudforceOneRequestPriorityDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    cloudforceOneRequestPriorityDeleteResponseEnvelopeJSON    `json:"-"`
}

// cloudforceOneRequestPriorityDeleteResponseEnvelopeJSON contains the JSON
// metadata for the struct [CloudforceOneRequestPriorityDeleteResponseEnvelope]
type cloudforceOneRequestPriorityDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestPriorityDeleteResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    cloudforceOneRequestPriorityDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// cloudforceOneRequestPriorityDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [CloudforceOneRequestPriorityDeleteResponseEnvelopeErrors]
type cloudforceOneRequestPriorityDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestPriorityDeleteResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    cloudforceOneRequestPriorityDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// cloudforceOneRequestPriorityDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [CloudforceOneRequestPriorityDeleteResponseEnvelopeMessages]
type cloudforceOneRequestPriorityDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CloudforceOneRequestPriorityDeleteResponseEnvelopeSuccess bool

const (
	CloudforceOneRequestPriorityDeleteResponseEnvelopeSuccessTrue CloudforceOneRequestPriorityDeleteResponseEnvelopeSuccess = true
)

type CloudforceOneRequestPriorityDoSomethingUnknownParams struct {
	// Page number of results
	Page param.Field[int64] `json:"page,required"`
	// Number of results per page
	PerPage param.Field[int64] `json:"per_page,required"`
}

func (r CloudforceOneRequestPriorityDoSomethingUnknownParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudforceOneRequestPriorityDoSomethingUnknownResponseEnvelope struct {
	Errors   []CloudforceOneRequestPriorityDoSomethingUnknownResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CloudforceOneRequestPriorityDoSomethingUnknownResponseEnvelopeMessages `json:"messages,required"`
	Result   []CloudforceOneRequestPriorityDoSomethingUnknownResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success CloudforceOneRequestPriorityDoSomethingUnknownResponseEnvelopeSuccess `json:"success,required"`
	JSON    cloudforceOneRequestPriorityDoSomethingUnknownResponseEnvelopeJSON    `json:"-"`
}

// cloudforceOneRequestPriorityDoSomethingUnknownResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [CloudforceOneRequestPriorityDoSomethingUnknownResponseEnvelope]
type cloudforceOneRequestPriorityDoSomethingUnknownResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityDoSomethingUnknownResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestPriorityDoSomethingUnknownResponseEnvelopeErrors struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    cloudforceOneRequestPriorityDoSomethingUnknownResponseEnvelopeErrorsJSON `json:"-"`
}

// cloudforceOneRequestPriorityDoSomethingUnknownResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [CloudforceOneRequestPriorityDoSomethingUnknownResponseEnvelopeErrors]
type cloudforceOneRequestPriorityDoSomethingUnknownResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityDoSomethingUnknownResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestPriorityDoSomethingUnknownResponseEnvelopeMessages struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    cloudforceOneRequestPriorityDoSomethingUnknownResponseEnvelopeMessagesJSON `json:"-"`
}

// cloudforceOneRequestPriorityDoSomethingUnknownResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [CloudforceOneRequestPriorityDoSomethingUnknownResponseEnvelopeMessages]
type cloudforceOneRequestPriorityDoSomethingUnknownResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityDoSomethingUnknownResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CloudforceOneRequestPriorityDoSomethingUnknownResponseEnvelopeSuccess bool

const (
	CloudforceOneRequestPriorityDoSomethingUnknownResponseEnvelopeSuccessTrue CloudforceOneRequestPriorityDoSomethingUnknownResponseEnvelopeSuccess = true
)

type CloudforceOneRequestPriorityGetResponseEnvelope struct {
	Errors   []CloudforceOneRequestPriorityGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CloudforceOneRequestPriorityGetResponseEnvelopeMessages `json:"messages,required"`
	Result   CloudforceOneRequestPriorityGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CloudforceOneRequestPriorityGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    cloudforceOneRequestPriorityGetResponseEnvelopeJSON    `json:"-"`
}

// cloudforceOneRequestPriorityGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [CloudforceOneRequestPriorityGetResponseEnvelope]
type cloudforceOneRequestPriorityGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestPriorityGetResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    cloudforceOneRequestPriorityGetResponseEnvelopeErrorsJSON `json:"-"`
}

// cloudforceOneRequestPriorityGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [CloudforceOneRequestPriorityGetResponseEnvelopeErrors]
type cloudforceOneRequestPriorityGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestPriorityGetResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    cloudforceOneRequestPriorityGetResponseEnvelopeMessagesJSON `json:"-"`
}

// cloudforceOneRequestPriorityGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [CloudforceOneRequestPriorityGetResponseEnvelopeMessages]
type cloudforceOneRequestPriorityGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CloudforceOneRequestPriorityGetResponseEnvelopeSuccess bool

const (
	CloudforceOneRequestPriorityGetResponseEnvelopeSuccessTrue CloudforceOneRequestPriorityGetResponseEnvelopeSuccess = true
)

type CloudforceOneRequestPriorityQuotaResponseEnvelope struct {
	Errors   []CloudforceOneRequestPriorityQuotaResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CloudforceOneRequestPriorityQuotaResponseEnvelopeMessages `json:"messages,required"`
	Result   CloudforceOneRequestPriorityQuotaResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CloudforceOneRequestPriorityQuotaResponseEnvelopeSuccess `json:"success,required"`
	JSON    cloudforceOneRequestPriorityQuotaResponseEnvelopeJSON    `json:"-"`
}

// cloudforceOneRequestPriorityQuotaResponseEnvelopeJSON contains the JSON metadata
// for the struct [CloudforceOneRequestPriorityQuotaResponseEnvelope]
type cloudforceOneRequestPriorityQuotaResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityQuotaResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestPriorityQuotaResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    cloudforceOneRequestPriorityQuotaResponseEnvelopeErrorsJSON `json:"-"`
}

// cloudforceOneRequestPriorityQuotaResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [CloudforceOneRequestPriorityQuotaResponseEnvelopeErrors]
type cloudforceOneRequestPriorityQuotaResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityQuotaResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestPriorityQuotaResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    cloudforceOneRequestPriorityQuotaResponseEnvelopeMessagesJSON `json:"-"`
}

// cloudforceOneRequestPriorityQuotaResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [CloudforceOneRequestPriorityQuotaResponseEnvelopeMessages]
type cloudforceOneRequestPriorityQuotaResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityQuotaResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CloudforceOneRequestPriorityQuotaResponseEnvelopeSuccess bool

const (
	CloudforceOneRequestPriorityQuotaResponseEnvelopeSuccessTrue CloudforceOneRequestPriorityQuotaResponseEnvelopeSuccess = true
)
