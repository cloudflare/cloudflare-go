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

// CloudforceOneRequestService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCloudforceOneRequestService]
// method instead.
type CloudforceOneRequestService struct {
	Options  []option.RequestOption
	Message  *CloudforceOneRequestMessageService
	Priority *CloudforceOneRequestPriorityService
}

// NewCloudforceOneRequestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudforceOneRequestService(opts ...option.RequestOption) (r *CloudforceOneRequestService) {
	r = &CloudforceOneRequestService{}
	r.Options = opts
	r.Message = NewCloudforceOneRequestMessageService(opts...)
	r.Priority = NewCloudforceOneRequestPriorityService(opts...)
	return
}

// Creating a request adds the request into the Cloudforce One queue for analysis.
// In addition to the content, a short title, type, priority, and releasability
// should be provided. If one is not provided a default will be assigned.
func (r *CloudforceOneRequestService) New(ctx context.Context, accountIdentifier string, body CloudforceOneRequestNewParams, opts ...option.RequestOption) (res *CloudforceOneRequestItem, err error) {
	opts = append(r.Options[:], opts...)
	var env CloudforceOneRequestNewResponseEnvelope
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
func (r *CloudforceOneRequestService) Update(ctx context.Context, accountIdentifier string, requestIdentifier string, body CloudforceOneRequestUpdateParams, opts ...option.RequestOption) (res *CloudforceOneRequestItem, err error) {
	opts = append(r.Options[:], opts...)
	var env CloudforceOneRequestUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s", accountIdentifier, requestIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Requests
func (r *CloudforceOneRequestService) List(ctx context.Context, accountIdentifier string, body CloudforceOneRequestListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[CloudforceOneRequestListItem], err error) {
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
func (r *CloudforceOneRequestService) ListAutoPaging(ctx context.Context, accountIdentifier string, body CloudforceOneRequestListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[CloudforceOneRequestListItem] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, accountIdentifier, body, opts...))
}

// Delete a Request
func (r *CloudforceOneRequestService) Delete(ctx context.Context, accountIdentifier string, requestIdentifier string, opts ...option.RequestOption) (res *CloudforceOneRequestDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CloudforceOneRequestDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s", accountIdentifier, requestIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Request Priority, Status, and TLP constants
func (r *CloudforceOneRequestService) Constants(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *CloudforceOneRequestConstants, err error) {
	opts = append(r.Options[:], opts...)
	var env CloudforceOneRequestConstantsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/constants", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a Request
func (r *CloudforceOneRequestService) Get(ctx context.Context, accountIdentifier string, requestIdentifier string, opts ...option.RequestOption) (res *CloudforceOneRequestItem, err error) {
	opts = append(r.Options[:], opts...)
	var env CloudforceOneRequestGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s", accountIdentifier, requestIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Request Quota
func (r *CloudforceOneRequestService) Quota(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *CloudforceOneQuota, err error) {
	opts = append(r.Options[:], opts...)
	var env CloudforceOneRequestQuotaResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/quota", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Request Types
func (r *CloudforceOneRequestService) Types(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *CloudforceOneRequestTypes, err error) {
	opts = append(r.Options[:], opts...)
	var env CloudforceOneRequestTypesResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/types", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CloudforceOneQuota struct {
	// Anniversary date is when annual quota limit is refresh
	AnniversaryDate time.Time `json:"anniversary_date" format:"date-time"`
	// Quater anniversary date is when quota limit is refreshed each quarter
	QuarterAnniversaryDate time.Time `json:"quarter_anniversary_date" format:"date-time"`
	// Tokens for the quarter
	Quota int64 `json:"quota"`
	// Tokens remaining for the quarter
	Remaining int64                  `json:"remaining"`
	JSON      cloudforceOneQuotaJSON `json:"-"`
}

// cloudforceOneQuotaJSON contains the JSON metadata for the struct
// [CloudforceOneQuota]
type cloudforceOneQuotaJSON struct {
	AnniversaryDate        apijson.Field
	QuarterAnniversaryDate apijson.Field
	Quota                  apijson.Field
	Remaining              apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CloudforceOneQuota) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestConstants struct {
	Priority []CloudforceOneRequestConstantsPriority `json:"priority"`
	Status   []CloudforceOneRequestConstantsStatus   `json:"status"`
	Tlp      []CloudforceOneRequestConstantsTlp      `json:"tlp"`
	JSON     cloudforceOneRequestConstantsJSON       `json:"-"`
}

// cloudforceOneRequestConstantsJSON contains the JSON metadata for the struct
// [CloudforceOneRequestConstants]
type cloudforceOneRequestConstantsJSON struct {
	Priority    apijson.Field
	Status      apijson.Field
	Tlp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestConstants) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestConstantsPriority string

const (
	CloudforceOneRequestConstantsPriorityRoutine CloudforceOneRequestConstantsPriority = "routine"
	CloudforceOneRequestConstantsPriorityHigh    CloudforceOneRequestConstantsPriority = "high"
	CloudforceOneRequestConstantsPriorityUrgent  CloudforceOneRequestConstantsPriority = "urgent"
)

// Request Status
type CloudforceOneRequestConstantsStatus string

const (
	CloudforceOneRequestConstantsStatusOpen      CloudforceOneRequestConstantsStatus = "open"
	CloudforceOneRequestConstantsStatusAccepted  CloudforceOneRequestConstantsStatus = "accepted"
	CloudforceOneRequestConstantsStatusReported  CloudforceOneRequestConstantsStatus = "reported"
	CloudforceOneRequestConstantsStatusApproved  CloudforceOneRequestConstantsStatus = "approved"
	CloudforceOneRequestConstantsStatusCompleted CloudforceOneRequestConstantsStatus = "completed"
	CloudforceOneRequestConstantsStatusDeclined  CloudforceOneRequestConstantsStatus = "declined"
)

// The CISA defined Traffic Light Protocol (TLP)
type CloudforceOneRequestConstantsTlp string

const (
	CloudforceOneRequestConstantsTlpClear       CloudforceOneRequestConstantsTlp = "clear"
	CloudforceOneRequestConstantsTlpAmber       CloudforceOneRequestConstantsTlp = "amber"
	CloudforceOneRequestConstantsTlpAmberStrict CloudforceOneRequestConstantsTlp = "amber-strict"
	CloudforceOneRequestConstantsTlpGreen       CloudforceOneRequestConstantsTlp = "green"
	CloudforceOneRequestConstantsTlpRed         CloudforceOneRequestConstantsTlp = "red"
)

type CloudforceOneRequestItem struct {
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
	Tlp       CloudforceOneRequestItemTlp `json:"tlp,required"`
	Updated   time.Time                   `json:"updated,required" format:"date-time"`
	Completed time.Time                   `json:"completed" format:"date-time"`
	// Tokens for the request messages
	MessageTokens int64 `json:"message_tokens"`
	// Readable Request ID
	ReadableID string `json:"readable_id"`
	// Request Status
	Status CloudforceOneRequestItemStatus `json:"status"`
	// Tokens for the request
	Tokens int64                        `json:"tokens"`
	JSON   cloudforceOneRequestItemJSON `json:"-"`
}

// cloudforceOneRequestItemJSON contains the JSON metadata for the struct
// [CloudforceOneRequestItem]
type cloudforceOneRequestItemJSON struct {
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

func (r *CloudforceOneRequestItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The CISA defined Traffic Light Protocol (TLP)
type CloudforceOneRequestItemTlp string

const (
	CloudforceOneRequestItemTlpClear       CloudforceOneRequestItemTlp = "clear"
	CloudforceOneRequestItemTlpAmber       CloudforceOneRequestItemTlp = "amber"
	CloudforceOneRequestItemTlpAmberStrict CloudforceOneRequestItemTlp = "amber-strict"
	CloudforceOneRequestItemTlpGreen       CloudforceOneRequestItemTlp = "green"
	CloudforceOneRequestItemTlpRed         CloudforceOneRequestItemTlp = "red"
)

// Request Status
type CloudforceOneRequestItemStatus string

const (
	CloudforceOneRequestItemStatusOpen      CloudforceOneRequestItemStatus = "open"
	CloudforceOneRequestItemStatusAccepted  CloudforceOneRequestItemStatus = "accepted"
	CloudforceOneRequestItemStatusReported  CloudforceOneRequestItemStatus = "reported"
	CloudforceOneRequestItemStatusApproved  CloudforceOneRequestItemStatus = "approved"
	CloudforceOneRequestItemStatusCompleted CloudforceOneRequestItemStatus = "completed"
	CloudforceOneRequestItemStatusDeclined  CloudforceOneRequestItemStatus = "declined"
)

type CloudforceOneRequestListItem struct {
	// UUID
	ID string `json:"id,required"`
	// Request creation time
	Created  time.Time                            `json:"created,required" format:"date-time"`
	Priority CloudforceOneRequestListItemPriority `json:"priority,required"`
	// Requested information from request
	Request string `json:"request,required"`
	// Brief description of the request
	Summary string `json:"summary,required"`
	// The CISA defined Traffic Light Protocol (TLP)
	Tlp CloudforceOneRequestListItemTlp `json:"tlp,required"`
	// Request last updated time
	Updated time.Time `json:"updated,required" format:"date-time"`
	// Request completion time
	Completed time.Time `json:"completed" format:"date-time"`
	// Tokens for the request messages
	MessageTokens int64 `json:"message_tokens"`
	// Readable Request ID
	ReadableID string `json:"readable_id"`
	// Request Status
	Status CloudforceOneRequestListItemStatus `json:"status"`
	// Tokens for the request
	Tokens int64                            `json:"tokens"`
	JSON   cloudforceOneRequestListItemJSON `json:"-"`
}

// cloudforceOneRequestListItemJSON contains the JSON metadata for the struct
// [CloudforceOneRequestListItem]
type cloudforceOneRequestListItemJSON struct {
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

func (r *CloudforceOneRequestListItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestListItemPriority string

const (
	CloudforceOneRequestListItemPriorityRoutine CloudforceOneRequestListItemPriority = "routine"
	CloudforceOneRequestListItemPriorityHigh    CloudforceOneRequestListItemPriority = "high"
	CloudforceOneRequestListItemPriorityUrgent  CloudforceOneRequestListItemPriority = "urgent"
)

// The CISA defined Traffic Light Protocol (TLP)
type CloudforceOneRequestListItemTlp string

const (
	CloudforceOneRequestListItemTlpClear       CloudforceOneRequestListItemTlp = "clear"
	CloudforceOneRequestListItemTlpAmber       CloudforceOneRequestListItemTlp = "amber"
	CloudforceOneRequestListItemTlpAmberStrict CloudforceOneRequestListItemTlp = "amber-strict"
	CloudforceOneRequestListItemTlpGreen       CloudforceOneRequestListItemTlp = "green"
	CloudforceOneRequestListItemTlpRed         CloudforceOneRequestListItemTlp = "red"
)

// Request Status
type CloudforceOneRequestListItemStatus string

const (
	CloudforceOneRequestListItemStatusOpen      CloudforceOneRequestListItemStatus = "open"
	CloudforceOneRequestListItemStatusAccepted  CloudforceOneRequestListItemStatus = "accepted"
	CloudforceOneRequestListItemStatusReported  CloudforceOneRequestListItemStatus = "reported"
	CloudforceOneRequestListItemStatusApproved  CloudforceOneRequestListItemStatus = "approved"
	CloudforceOneRequestListItemStatusCompleted CloudforceOneRequestListItemStatus = "completed"
	CloudforceOneRequestListItemStatusDeclined  CloudforceOneRequestListItemStatus = "declined"
)

type CloudforceOneRequestTypes []string

// Union satisfied by [CloudforceOneRequestDeleteResponseUnknown],
// [CloudforceOneRequestDeleteResponseArray] or [shared.UnionString].
type CloudforceOneRequestDeleteResponse interface {
	ImplementsCloudforceOneRequestDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudforceOneRequestDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudforceOneRequestDeleteResponseArray []interface{}

func (r CloudforceOneRequestDeleteResponseArray) ImplementsCloudforceOneRequestDeleteResponse() {}

type CloudforceOneRequestNewParams struct {
	// Request content
	Content param.Field[string] `json:"content"`
	// Priority for analyzing the request
	Priority param.Field[string] `json:"priority"`
	// Requested information from request
	RequestType param.Field[string] `json:"request_type"`
	// Brief description of the request
	Summary param.Field[string] `json:"summary"`
	// The CISA defined Traffic Light Protocol (TLP)
	Tlp param.Field[CloudforceOneRequestNewParamsTlp] `json:"tlp"`
}

func (r CloudforceOneRequestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The CISA defined Traffic Light Protocol (TLP)
type CloudforceOneRequestNewParamsTlp string

const (
	CloudforceOneRequestNewParamsTlpClear       CloudforceOneRequestNewParamsTlp = "clear"
	CloudforceOneRequestNewParamsTlpAmber       CloudforceOneRequestNewParamsTlp = "amber"
	CloudforceOneRequestNewParamsTlpAmberStrict CloudforceOneRequestNewParamsTlp = "amber-strict"
	CloudforceOneRequestNewParamsTlpGreen       CloudforceOneRequestNewParamsTlp = "green"
	CloudforceOneRequestNewParamsTlpRed         CloudforceOneRequestNewParamsTlp = "red"
)

type CloudforceOneRequestNewResponseEnvelope struct {
	Errors   []CloudforceOneRequestNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CloudforceOneRequestNewResponseEnvelopeMessages `json:"messages,required"`
	Result   CloudforceOneRequestItem                          `json:"result,required"`
	// Whether the API call was successful
	Success CloudforceOneRequestNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    cloudforceOneRequestNewResponseEnvelopeJSON    `json:"-"`
}

// cloudforceOneRequestNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [CloudforceOneRequestNewResponseEnvelope]
type cloudforceOneRequestNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestNewResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    cloudforceOneRequestNewResponseEnvelopeErrorsJSON `json:"-"`
}

// cloudforceOneRequestNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CloudforceOneRequestNewResponseEnvelopeErrors]
type cloudforceOneRequestNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestNewResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    cloudforceOneRequestNewResponseEnvelopeMessagesJSON `json:"-"`
}

// cloudforceOneRequestNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CloudforceOneRequestNewResponseEnvelopeMessages]
type cloudforceOneRequestNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CloudforceOneRequestNewResponseEnvelopeSuccess bool

const (
	CloudforceOneRequestNewResponseEnvelopeSuccessTrue CloudforceOneRequestNewResponseEnvelopeSuccess = true
)

type CloudforceOneRequestUpdateParams struct {
	// Request content
	Content param.Field[string] `json:"content"`
	// Priority for analyzing the request
	Priority param.Field[string] `json:"priority"`
	// Requested information from request
	RequestType param.Field[string] `json:"request_type"`
	// Brief description of the request
	Summary param.Field[string] `json:"summary"`
	// The CISA defined Traffic Light Protocol (TLP)
	Tlp param.Field[CloudforceOneRequestUpdateParamsTlp] `json:"tlp"`
}

func (r CloudforceOneRequestUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The CISA defined Traffic Light Protocol (TLP)
type CloudforceOneRequestUpdateParamsTlp string

const (
	CloudforceOneRequestUpdateParamsTlpClear       CloudforceOneRequestUpdateParamsTlp = "clear"
	CloudforceOneRequestUpdateParamsTlpAmber       CloudforceOneRequestUpdateParamsTlp = "amber"
	CloudforceOneRequestUpdateParamsTlpAmberStrict CloudforceOneRequestUpdateParamsTlp = "amber-strict"
	CloudforceOneRequestUpdateParamsTlpGreen       CloudforceOneRequestUpdateParamsTlp = "green"
	CloudforceOneRequestUpdateParamsTlpRed         CloudforceOneRequestUpdateParamsTlp = "red"
)

type CloudforceOneRequestUpdateResponseEnvelope struct {
	Errors   []CloudforceOneRequestUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CloudforceOneRequestUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   CloudforceOneRequestItem                             `json:"result,required"`
	// Whether the API call was successful
	Success CloudforceOneRequestUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    cloudforceOneRequestUpdateResponseEnvelopeJSON    `json:"-"`
}

// cloudforceOneRequestUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [CloudforceOneRequestUpdateResponseEnvelope]
type cloudforceOneRequestUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestUpdateResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    cloudforceOneRequestUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// cloudforceOneRequestUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [CloudforceOneRequestUpdateResponseEnvelopeErrors]
type cloudforceOneRequestUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestUpdateResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    cloudforceOneRequestUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// cloudforceOneRequestUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [CloudforceOneRequestUpdateResponseEnvelopeMessages]
type cloudforceOneRequestUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CloudforceOneRequestUpdateResponseEnvelopeSuccess bool

const (
	CloudforceOneRequestUpdateResponseEnvelopeSuccessTrue CloudforceOneRequestUpdateResponseEnvelopeSuccess = true
)

type CloudforceOneRequestListParams struct {
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
	SortOrder param.Field[CloudforceOneRequestListParamsSortOrder] `json:"sort_order"`
	// Request Status
	Status param.Field[CloudforceOneRequestListParamsStatus] `json:"status"`
}

func (r CloudforceOneRequestListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Sort order (asc or desc)
type CloudforceOneRequestListParamsSortOrder string

const (
	CloudforceOneRequestListParamsSortOrderAsc  CloudforceOneRequestListParamsSortOrder = "asc"
	CloudforceOneRequestListParamsSortOrderDesc CloudforceOneRequestListParamsSortOrder = "desc"
)

// Request Status
type CloudforceOneRequestListParamsStatus string

const (
	CloudforceOneRequestListParamsStatusOpen      CloudforceOneRequestListParamsStatus = "open"
	CloudforceOneRequestListParamsStatusAccepted  CloudforceOneRequestListParamsStatus = "accepted"
	CloudforceOneRequestListParamsStatusReported  CloudforceOneRequestListParamsStatus = "reported"
	CloudforceOneRequestListParamsStatusApproved  CloudforceOneRequestListParamsStatus = "approved"
	CloudforceOneRequestListParamsStatusCompleted CloudforceOneRequestListParamsStatus = "completed"
	CloudforceOneRequestListParamsStatusDeclined  CloudforceOneRequestListParamsStatus = "declined"
)

type CloudforceOneRequestDeleteResponseEnvelope struct {
	Errors   []CloudforceOneRequestDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CloudforceOneRequestDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   CloudforceOneRequestDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CloudforceOneRequestDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    cloudforceOneRequestDeleteResponseEnvelopeJSON    `json:"-"`
}

// cloudforceOneRequestDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [CloudforceOneRequestDeleteResponseEnvelope]
type cloudforceOneRequestDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestDeleteResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    cloudforceOneRequestDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// cloudforceOneRequestDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [CloudforceOneRequestDeleteResponseEnvelopeErrors]
type cloudforceOneRequestDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestDeleteResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    cloudforceOneRequestDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// cloudforceOneRequestDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [CloudforceOneRequestDeleteResponseEnvelopeMessages]
type cloudforceOneRequestDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CloudforceOneRequestDeleteResponseEnvelopeSuccess bool

const (
	CloudforceOneRequestDeleteResponseEnvelopeSuccessTrue CloudforceOneRequestDeleteResponseEnvelopeSuccess = true
)

type CloudforceOneRequestConstantsResponseEnvelope struct {
	Errors   []CloudforceOneRequestConstantsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CloudforceOneRequestConstantsResponseEnvelopeMessages `json:"messages,required"`
	Result   CloudforceOneRequestConstants                           `json:"result,required"`
	// Whether the API call was successful
	Success CloudforceOneRequestConstantsResponseEnvelopeSuccess `json:"success,required"`
	JSON    cloudforceOneRequestConstantsResponseEnvelopeJSON    `json:"-"`
}

// cloudforceOneRequestConstantsResponseEnvelopeJSON contains the JSON metadata for
// the struct [CloudforceOneRequestConstantsResponseEnvelope]
type cloudforceOneRequestConstantsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestConstantsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestConstantsResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    cloudforceOneRequestConstantsResponseEnvelopeErrorsJSON `json:"-"`
}

// cloudforceOneRequestConstantsResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [CloudforceOneRequestConstantsResponseEnvelopeErrors]
type cloudforceOneRequestConstantsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestConstantsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestConstantsResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    cloudforceOneRequestConstantsResponseEnvelopeMessagesJSON `json:"-"`
}

// cloudforceOneRequestConstantsResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [CloudforceOneRequestConstantsResponseEnvelopeMessages]
type cloudforceOneRequestConstantsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestConstantsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CloudforceOneRequestConstantsResponseEnvelopeSuccess bool

const (
	CloudforceOneRequestConstantsResponseEnvelopeSuccessTrue CloudforceOneRequestConstantsResponseEnvelopeSuccess = true
)

type CloudforceOneRequestGetResponseEnvelope struct {
	Errors   []CloudforceOneRequestGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CloudforceOneRequestGetResponseEnvelopeMessages `json:"messages,required"`
	Result   CloudforceOneRequestItem                          `json:"result,required"`
	// Whether the API call was successful
	Success CloudforceOneRequestGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    cloudforceOneRequestGetResponseEnvelopeJSON    `json:"-"`
}

// cloudforceOneRequestGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [CloudforceOneRequestGetResponseEnvelope]
type cloudforceOneRequestGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestGetResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    cloudforceOneRequestGetResponseEnvelopeErrorsJSON `json:"-"`
}

// cloudforceOneRequestGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CloudforceOneRequestGetResponseEnvelopeErrors]
type cloudforceOneRequestGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestGetResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    cloudforceOneRequestGetResponseEnvelopeMessagesJSON `json:"-"`
}

// cloudforceOneRequestGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CloudforceOneRequestGetResponseEnvelopeMessages]
type cloudforceOneRequestGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CloudforceOneRequestGetResponseEnvelopeSuccess bool

const (
	CloudforceOneRequestGetResponseEnvelopeSuccessTrue CloudforceOneRequestGetResponseEnvelopeSuccess = true
)

type CloudforceOneRequestQuotaResponseEnvelope struct {
	Errors   []CloudforceOneRequestQuotaResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CloudforceOneRequestQuotaResponseEnvelopeMessages `json:"messages,required"`
	Result   CloudforceOneQuota                                  `json:"result,required"`
	// Whether the API call was successful
	Success CloudforceOneRequestQuotaResponseEnvelopeSuccess `json:"success,required"`
	JSON    cloudforceOneRequestQuotaResponseEnvelopeJSON    `json:"-"`
}

// cloudforceOneRequestQuotaResponseEnvelopeJSON contains the JSON metadata for the
// struct [CloudforceOneRequestQuotaResponseEnvelope]
type cloudforceOneRequestQuotaResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestQuotaResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestQuotaResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    cloudforceOneRequestQuotaResponseEnvelopeErrorsJSON `json:"-"`
}

// cloudforceOneRequestQuotaResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [CloudforceOneRequestQuotaResponseEnvelopeErrors]
type cloudforceOneRequestQuotaResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestQuotaResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestQuotaResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    cloudforceOneRequestQuotaResponseEnvelopeMessagesJSON `json:"-"`
}

// cloudforceOneRequestQuotaResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CloudforceOneRequestQuotaResponseEnvelopeMessages]
type cloudforceOneRequestQuotaResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestQuotaResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CloudforceOneRequestQuotaResponseEnvelopeSuccess bool

const (
	CloudforceOneRequestQuotaResponseEnvelopeSuccessTrue CloudforceOneRequestQuotaResponseEnvelopeSuccess = true
)

type CloudforceOneRequestTypesResponseEnvelope struct {
	Errors   []CloudforceOneRequestTypesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CloudforceOneRequestTypesResponseEnvelopeMessages `json:"messages,required"`
	Result   CloudforceOneRequestTypes                           `json:"result,required"`
	// Whether the API call was successful
	Success CloudforceOneRequestTypesResponseEnvelopeSuccess `json:"success,required"`
	JSON    cloudforceOneRequestTypesResponseEnvelopeJSON    `json:"-"`
}

// cloudforceOneRequestTypesResponseEnvelopeJSON contains the JSON metadata for the
// struct [CloudforceOneRequestTypesResponseEnvelope]
type cloudforceOneRequestTypesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestTypesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestTypesResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    cloudforceOneRequestTypesResponseEnvelopeErrorsJSON `json:"-"`
}

// cloudforceOneRequestTypesResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [CloudforceOneRequestTypesResponseEnvelopeErrors]
type cloudforceOneRequestTypesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestTypesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestTypesResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    cloudforceOneRequestTypesResponseEnvelopeMessagesJSON `json:"-"`
}

// cloudforceOneRequestTypesResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CloudforceOneRequestTypesResponseEnvelopeMessages]
type cloudforceOneRequestTypesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestTypesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CloudforceOneRequestTypesResponseEnvelopeSuccess bool

const (
	CloudforceOneRequestTypesResponseEnvelopeSuccessTrue CloudforceOneRequestTypesResponseEnvelopeSuccess = true
)
