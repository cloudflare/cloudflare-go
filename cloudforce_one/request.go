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
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
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
func (r *RequestService) New(ctx context.Context, accountIdentifier string, body RequestNewParams, opts ...option.RequestOption) (res *Item, err error) {
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
func (r *RequestService) Update(ctx context.Context, accountIdentifier string, requestIdentifier string, body RequestUpdateParams, opts ...option.RequestOption) (res *Item, err error) {
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
func (r *RequestService) List(ctx context.Context, accountIdentifier string, body RequestListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[ListItem], err error) {
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
func (r *RequestService) ListAutoPaging(ctx context.Context, accountIdentifier string, body RequestListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[ListItem] {
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
func (r *RequestService) Constants(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *RequestConstants, err error) {
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
func (r *RequestService) Get(ctx context.Context, accountIdentifier string, requestIdentifier string, opts ...option.RequestOption) (res *Item, err error) {
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
func (r *RequestService) Quota(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *Quota, err error) {
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
func (r *RequestService) Types(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *RequestTypes, err error) {
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

type Item struct {
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
	TLP       ItemTLP   `json:"tlp,required"`
	Updated   time.Time `json:"updated,required" format:"date-time"`
	Completed time.Time `json:"completed" format:"date-time"`
	// Tokens for the request messages
	MessageTokens int64 `json:"message_tokens"`
	// Readable Request ID
	ReadableID string `json:"readable_id"`
	// Request Status
	Status ItemStatus `json:"status"`
	// Tokens for the request
	Tokens int64    `json:"tokens"`
	JSON   itemJSON `json:"-"`
}

// itemJSON contains the JSON metadata for the struct [Item]
type itemJSON struct {
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

func (r *Item) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r itemJSON) RawJSON() string {
	return r.raw
}

// The CISA defined Traffic Light Protocol (TLP)
type ItemTLP string

const (
	ItemTLPClear       ItemTLP = "clear"
	ItemTLPAmber       ItemTLP = "amber"
	ItemTLPAmberStrict ItemTLP = "amber-strict"
	ItemTLPGreen       ItemTLP = "green"
	ItemTLPRed         ItemTLP = "red"
)

func (r ItemTLP) IsKnown() bool {
	switch r {
	case ItemTLPClear, ItemTLPAmber, ItemTLPAmberStrict, ItemTLPGreen, ItemTLPRed:
		return true
	}
	return false
}

// Request Status
type ItemStatus string

const (
	ItemStatusOpen      ItemStatus = "open"
	ItemStatusAccepted  ItemStatus = "accepted"
	ItemStatusReported  ItemStatus = "reported"
	ItemStatusApproved  ItemStatus = "approved"
	ItemStatusCompleted ItemStatus = "completed"
	ItemStatusDeclined  ItemStatus = "declined"
)

func (r ItemStatus) IsKnown() bool {
	switch r {
	case ItemStatusOpen, ItemStatusAccepted, ItemStatusReported, ItemStatusApproved, ItemStatusCompleted, ItemStatusDeclined:
		return true
	}
	return false
}

type ListItem struct {
	// UUID
	ID string `json:"id,required"`
	// Request creation time
	Created  time.Time        `json:"created,required" format:"date-time"`
	Priority ListItemPriority `json:"priority,required"`
	// Requested information from request
	Request string `json:"request,required"`
	// Brief description of the request
	Summary string `json:"summary,required"`
	// The CISA defined Traffic Light Protocol (TLP)
	TLP ListItemTLP `json:"tlp,required"`
	// Request last updated time
	Updated time.Time `json:"updated,required" format:"date-time"`
	// Request completion time
	Completed time.Time `json:"completed" format:"date-time"`
	// Tokens for the request messages
	MessageTokens int64 `json:"message_tokens"`
	// Readable Request ID
	ReadableID string `json:"readable_id"`
	// Request Status
	Status ListItemStatus `json:"status"`
	// Tokens for the request
	Tokens int64        `json:"tokens"`
	JSON   listItemJSON `json:"-"`
}

// listItemJSON contains the JSON metadata for the struct [ListItem]
type listItemJSON struct {
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

func (r *ListItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemJSON) RawJSON() string {
	return r.raw
}

type ListItemPriority string

const (
	ListItemPriorityRoutine ListItemPriority = "routine"
	ListItemPriorityHigh    ListItemPriority = "high"
	ListItemPriorityUrgent  ListItemPriority = "urgent"
)

func (r ListItemPriority) IsKnown() bool {
	switch r {
	case ListItemPriorityRoutine, ListItemPriorityHigh, ListItemPriorityUrgent:
		return true
	}
	return false
}

// The CISA defined Traffic Light Protocol (TLP)
type ListItemTLP string

const (
	ListItemTLPClear       ListItemTLP = "clear"
	ListItemTLPAmber       ListItemTLP = "amber"
	ListItemTLPAmberStrict ListItemTLP = "amber-strict"
	ListItemTLPGreen       ListItemTLP = "green"
	ListItemTLPRed         ListItemTLP = "red"
)

func (r ListItemTLP) IsKnown() bool {
	switch r {
	case ListItemTLPClear, ListItemTLPAmber, ListItemTLPAmberStrict, ListItemTLPGreen, ListItemTLPRed:
		return true
	}
	return false
}

// Request Status
type ListItemStatus string

const (
	ListItemStatusOpen      ListItemStatus = "open"
	ListItemStatusAccepted  ListItemStatus = "accepted"
	ListItemStatusReported  ListItemStatus = "reported"
	ListItemStatusApproved  ListItemStatus = "approved"
	ListItemStatusCompleted ListItemStatus = "completed"
	ListItemStatusDeclined  ListItemStatus = "declined"
)

func (r ListItemStatus) IsKnown() bool {
	switch r {
	case ListItemStatusOpen, ListItemStatusAccepted, ListItemStatusReported, ListItemStatusApproved, ListItemStatusCompleted, ListItemStatusDeclined:
		return true
	}
	return false
}

type Quota struct {
	// Anniversary date is when annual quota limit is refresh
	AnniversaryDate time.Time `json:"anniversary_date" format:"date-time"`
	// Quater anniversary date is when quota limit is refreshed each quarter
	QuarterAnniversaryDate time.Time `json:"quarter_anniversary_date" format:"date-time"`
	// Tokens for the quarter
	Quota int64 `json:"quota"`
	// Tokens remaining for the quarter
	Remaining int64     `json:"remaining"`
	JSON      quotaJSON `json:"-"`
}

// quotaJSON contains the JSON metadata for the struct [Quota]
type quotaJSON struct {
	AnniversaryDate        apijson.Field
	QuarterAnniversaryDate apijson.Field
	Quota                  apijson.Field
	Remaining              apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *Quota) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r quotaJSON) RawJSON() string {
	return r.raw
}

type RequestConstants struct {
	Priority []RequestConstantsPriority `json:"priority"`
	Status   []RequestConstantsStatus   `json:"status"`
	TLP      []RequestConstantsTLP      `json:"tlp"`
	JSON     requestConstantsJSON       `json:"-"`
}

// requestConstantsJSON contains the JSON metadata for the struct
// [RequestConstants]
type requestConstantsJSON struct {
	Priority    apijson.Field
	Status      apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestConstants) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestConstantsJSON) RawJSON() string {
	return r.raw
}

type RequestConstantsPriority string

const (
	RequestConstantsPriorityRoutine RequestConstantsPriority = "routine"
	RequestConstantsPriorityHigh    RequestConstantsPriority = "high"
	RequestConstantsPriorityUrgent  RequestConstantsPriority = "urgent"
)

func (r RequestConstantsPriority) IsKnown() bool {
	switch r {
	case RequestConstantsPriorityRoutine, RequestConstantsPriorityHigh, RequestConstantsPriorityUrgent:
		return true
	}
	return false
}

// Request Status
type RequestConstantsStatus string

const (
	RequestConstantsStatusOpen      RequestConstantsStatus = "open"
	RequestConstantsStatusAccepted  RequestConstantsStatus = "accepted"
	RequestConstantsStatusReported  RequestConstantsStatus = "reported"
	RequestConstantsStatusApproved  RequestConstantsStatus = "approved"
	RequestConstantsStatusCompleted RequestConstantsStatus = "completed"
	RequestConstantsStatusDeclined  RequestConstantsStatus = "declined"
)

func (r RequestConstantsStatus) IsKnown() bool {
	switch r {
	case RequestConstantsStatusOpen, RequestConstantsStatusAccepted, RequestConstantsStatusReported, RequestConstantsStatusApproved, RequestConstantsStatusCompleted, RequestConstantsStatusDeclined:
		return true
	}
	return false
}

// The CISA defined Traffic Light Protocol (TLP)
type RequestConstantsTLP string

const (
	RequestConstantsTLPClear       RequestConstantsTLP = "clear"
	RequestConstantsTLPAmber       RequestConstantsTLP = "amber"
	RequestConstantsTLPAmberStrict RequestConstantsTLP = "amber-strict"
	RequestConstantsTLPGreen       RequestConstantsTLP = "green"
	RequestConstantsTLPRed         RequestConstantsTLP = "red"
)

func (r RequestConstantsTLP) IsKnown() bool {
	switch r {
	case RequestConstantsTLPClear, RequestConstantsTLPAmber, RequestConstantsTLPAmberStrict, RequestConstantsTLPGreen, RequestConstantsTLPRed:
		return true
	}
	return false
}

type RequestTypes []string

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
	Result  Item                              `json:"result"`
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
	Result  Item                                 `json:"result"`
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
	Result  RequestConstants                        `json:"result"`
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
	Result  Item                              `json:"result"`
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
	Result  Quota                               `json:"result"`
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
	Result  RequestTypes                        `json:"result"`
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
