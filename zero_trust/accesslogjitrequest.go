// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// AccessLogJitRequestService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessLogJitRequestService] method instead.
type AccessLogJitRequestService struct {
	Options []option.RequestOption
}

// NewAccessLogJitRequestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessLogJitRequestService(opts ...option.RequestOption) (r *AccessLogJitRequestService) {
	r = &AccessLogJitRequestService{}
	r.Options = opts
	return
}

// Lists account-wide Access JIT request logs reconstructed from request lifecycle
// events.
func (r *AccessLogJitRequestService) List(ctx context.Context, params AccessLogJitRequestListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AccessLogJitRequestListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/logs/jit_requests", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Lists account-wide Access JIT request logs reconstructed from request lifecycle
// events.
func (r *AccessLogJitRequestService) ListAutoPaging(ctx context.Context, params AccessLogJitRequestListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AccessLogJitRequestListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Gets an account-scoped Access JIT request summary and its lifecycle events in
// chronological order.
func (r *AccessLogJitRequestService) Get(ctx context.Context, knockRequestID string, query AccessLogJitRequestGetParams, opts ...option.RequestOption) (res *AccessLogJitRequestGetResponse, err error) {
	var env AccessLogJitRequestGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if knockRequestID == "" {
		err = errors.New("missing required knock_request_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/logs/jit_requests/%s", query.AccountID, knockRequestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type AccessLogJitRequestListResponse struct {
	// Audience of the Access application.
	AppAUD      string `json:"app_aud"`
	AppHostname string `json:"app_hostname"`
	// Number of unique actors who approved the request.
	ApprovalsReceived int64 `json:"approvals_received"`
	// Unique actors who approved the request.
	ApproverEmails []string  `json:"approver_emails" format:"email"`
	CreatedAt      time.Time `json:"created_at" format:"date-time"`
	ExpiresAt      time.Time `json:"expires_at" format:"date-time"`
	// Unique identifier for the JIT request.
	KnockRequestID       string `json:"knock_request_id"`
	PurposeJustification string `json:"purpose_justification"`
	RequesterEmail       string `json:"requester_email" format:"email"`
	// JIT request status. `SPENT` is deprecated and interpreted as `APPROVED`.
	Status AccessLogJitRequestListResponseStatus `json:"status"`
	JSON   accessLogJitRequestListResponseJSON   `json:"-"`
}

// accessLogJitRequestListResponseJSON contains the JSON metadata for the struct
// [AccessLogJitRequestListResponse]
type accessLogJitRequestListResponseJSON struct {
	AppAUD               apijson.Field
	AppHostname          apijson.Field
	ApprovalsReceived    apijson.Field
	ApproverEmails       apijson.Field
	CreatedAt            apijson.Field
	ExpiresAt            apijson.Field
	KnockRequestID       apijson.Field
	PurposeJustification apijson.Field
	RequesterEmail       apijson.Field
	Status               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessLogJitRequestListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessLogJitRequestListResponseJSON) RawJSON() string {
	return r.raw
}

// JIT request status. `SPENT` is deprecated and interpreted as `APPROVED`.
type AccessLogJitRequestListResponseStatus string

const (
	AccessLogJitRequestListResponseStatusPending  AccessLogJitRequestListResponseStatus = "PENDING"
	AccessLogJitRequestListResponseStatusApproved AccessLogJitRequestListResponseStatus = "APPROVED"
	AccessLogJitRequestListResponseStatusDenied   AccessLogJitRequestListResponseStatus = "DENIED"
	AccessLogJitRequestListResponseStatusCanceled AccessLogJitRequestListResponseStatus = "CANCELED"
	AccessLogJitRequestListResponseStatusSpent    AccessLogJitRequestListResponseStatus = "SPENT"
)

func (r AccessLogJitRequestListResponseStatus) IsKnown() bool {
	switch r {
	case AccessLogJitRequestListResponseStatusPending, AccessLogJitRequestListResponseStatusApproved, AccessLogJitRequestListResponseStatusDenied, AccessLogJitRequestListResponseStatusCanceled, AccessLogJitRequestListResponseStatusSpent:
		return true
	}
	return false
}

type AccessLogJitRequestGetResponse struct {
	// Audience of the Access application.
	AppAUD      string `json:"app_aud"`
	AppHostname string `json:"app_hostname"`
	// Number of unique actors who approved the request.
	ApprovalsReceived int64 `json:"approvals_received"`
	// Unique actors who approved the request.
	ApproverEmails []string                              `json:"approver_emails" format:"email"`
	CreatedAt      time.Time                             `json:"created_at" format:"date-time"`
	Events         []AccessLogJitRequestGetResponseEvent `json:"events"`
	ExpiresAt      time.Time                             `json:"expires_at" format:"date-time"`
	// Unique identifier for the JIT request.
	KnockRequestID       string `json:"knock_request_id"`
	PurposeJustification string `json:"purpose_justification"`
	RequesterEmail       string `json:"requester_email" format:"email"`
	// JIT request status. `SPENT` is deprecated and interpreted as `APPROVED`.
	Status AccessLogJitRequestGetResponseStatus `json:"status"`
	JSON   accessLogJitRequestGetResponseJSON   `json:"-"`
}

// accessLogJitRequestGetResponseJSON contains the JSON metadata for the struct
// [AccessLogJitRequestGetResponse]
type accessLogJitRequestGetResponseJSON struct {
	AppAUD               apijson.Field
	AppHostname          apijson.Field
	ApprovalsReceived    apijson.Field
	ApproverEmails       apijson.Field
	CreatedAt            apijson.Field
	Events               apijson.Field
	ExpiresAt            apijson.Field
	KnockRequestID       apijson.Field
	PurposeJustification apijson.Field
	RequesterEmail       apijson.Field
	Status               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessLogJitRequestGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessLogJitRequestGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccessLogJitRequestGetResponseEvent struct {
	ActorEmail             string                                   `json:"actor_email" format:"email"`
	ActorIdP               string                                   `json:"actor_idp"`
	ActorUUID              string                                   `json:"actor_uuid" format:"uuid"`
	DecisionOutcome        string                                   `json:"decision_outcome"`
	RayID                  string                                   `json:"ray_id"`
	SessionDurationSeconds int64                                    `json:"session_duration_seconds"`
	Success                bool                                     `json:"success"`
	Timestamp              time.Time                                `json:"timestamp" format:"date-time"`
	Type                   AccessLogJitRequestGetResponseEventsType `json:"type"`
	JSON                   accessLogJitRequestGetResponseEventJSON  `json:"-"`
}

// accessLogJitRequestGetResponseEventJSON contains the JSON metadata for the
// struct [AccessLogJitRequestGetResponseEvent]
type accessLogJitRequestGetResponseEventJSON struct {
	ActorEmail             apijson.Field
	ActorIdP               apijson.Field
	ActorUUID              apijson.Field
	DecisionOutcome        apijson.Field
	RayID                  apijson.Field
	SessionDurationSeconds apijson.Field
	Success                apijson.Field
	Timestamp              apijson.Field
	Type                   apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessLogJitRequestGetResponseEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessLogJitRequestGetResponseEventJSON) RawJSON() string {
	return r.raw
}

type AccessLogJitRequestGetResponseEventsType string

const (
	AccessLogJitRequestGetResponseEventsTypeCreated          AccessLogJitRequestGetResponseEventsType = "CREATED"
	AccessLogJitRequestGetResponseEventsTypeApproverDecision AccessLogJitRequestGetResponseEventsType = "APPROVER_DECISION"
	AccessLogJitRequestGetResponseEventsTypeCommitted        AccessLogJitRequestGetResponseEventsType = "COMMITTED"
	AccessLogJitRequestGetResponseEventsTypeCanceled         AccessLogJitRequestGetResponseEventsType = "CANCELED"
	AccessLogJitRequestGetResponseEventsTypeLoginSuccess     AccessLogJitRequestGetResponseEventsType = "LOGIN_SUCCESS"
	AccessLogJitRequestGetResponseEventsTypeRedirectToActive AccessLogJitRequestGetResponseEventsType = "REDIRECT_TO_ACTIVE"
)

func (r AccessLogJitRequestGetResponseEventsType) IsKnown() bool {
	switch r {
	case AccessLogJitRequestGetResponseEventsTypeCreated, AccessLogJitRequestGetResponseEventsTypeApproverDecision, AccessLogJitRequestGetResponseEventsTypeCommitted, AccessLogJitRequestGetResponseEventsTypeCanceled, AccessLogJitRequestGetResponseEventsTypeLoginSuccess, AccessLogJitRequestGetResponseEventsTypeRedirectToActive:
		return true
	}
	return false
}

// JIT request status. `SPENT` is deprecated and interpreted as `APPROVED`.
type AccessLogJitRequestGetResponseStatus string

const (
	AccessLogJitRequestGetResponseStatusPending  AccessLogJitRequestGetResponseStatus = "PENDING"
	AccessLogJitRequestGetResponseStatusApproved AccessLogJitRequestGetResponseStatus = "APPROVED"
	AccessLogJitRequestGetResponseStatusDenied   AccessLogJitRequestGetResponseStatus = "DENIED"
	AccessLogJitRequestGetResponseStatusCanceled AccessLogJitRequestGetResponseStatus = "CANCELED"
	AccessLogJitRequestGetResponseStatusSpent    AccessLogJitRequestGetResponseStatus = "SPENT"
)

func (r AccessLogJitRequestGetResponseStatus) IsKnown() bool {
	switch r {
	case AccessLogJitRequestGetResponseStatusPending, AccessLogJitRequestGetResponseStatusApproved, AccessLogJitRequestGetResponseStatusDenied, AccessLogJitRequestGetResponseStatusCanceled, AccessLogJitRequestGetResponseStatusSpent:
		return true
	}
	return false
}

type AccessLogJitRequestListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Page number of results.
	Page param.Field[int64] `query:"page"`
	// Number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Case-insensitive search over request ID, requester email, application audience,
	// and application hostname.
	Search param.Field[string] `query:"search"`
	// The earliest request timestamp to query. Defaults to 366 days before the current
	// time.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// JIT request status. `SPENT` is deprecated and interpreted as `APPROVED`.
	Status param.Field[AccessLogJitRequestListParamsStatus] `query:"status"`
	// The latest request timestamp to query. Defaults to the current time.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [AccessLogJitRequestListParams]'s query parameters as
// `url.Values`.
func (r AccessLogJitRequestListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// JIT request status. `SPENT` is deprecated and interpreted as `APPROVED`.
type AccessLogJitRequestListParamsStatus string

const (
	AccessLogJitRequestListParamsStatusPending  AccessLogJitRequestListParamsStatus = "PENDING"
	AccessLogJitRequestListParamsStatusApproved AccessLogJitRequestListParamsStatus = "APPROVED"
	AccessLogJitRequestListParamsStatusDenied   AccessLogJitRequestListParamsStatus = "DENIED"
	AccessLogJitRequestListParamsStatusCanceled AccessLogJitRequestListParamsStatus = "CANCELED"
	AccessLogJitRequestListParamsStatusSpent    AccessLogJitRequestListParamsStatus = "SPENT"
)

func (r AccessLogJitRequestListParamsStatus) IsKnown() bool {
	switch r {
	case AccessLogJitRequestListParamsStatusPending, AccessLogJitRequestListParamsStatusApproved, AccessLogJitRequestListParamsStatusDenied, AccessLogJitRequestListParamsStatusCanceled, AccessLogJitRequestListParamsStatusSpent:
		return true
	}
	return false
}

type AccessLogJitRequestGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AccessLogJitRequestGetResponseEnvelope struct {
	Errors   []AccessLogJitRequestGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AccessLogJitRequestGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AccessLogJitRequestGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AccessLogJitRequestGetResponse                `json:"result"`
	JSON    accessLogJitRequestGetResponseEnvelopeJSON    `json:"-"`
}

// accessLogJitRequestGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessLogJitRequestGetResponseEnvelope]
type accessLogJitRequestGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessLogJitRequestGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessLogJitRequestGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessLogJitRequestGetResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code" api:"required"`
	Message          string                                             `json:"message" api:"required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           AccessLogJitRequestGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessLogJitRequestGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessLogJitRequestGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessLogJitRequestGetResponseEnvelopeErrors]
type accessLogJitRequestGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessLogJitRequestGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessLogJitRequestGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessLogJitRequestGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    accessLogJitRequestGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessLogJitRequestGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [AccessLogJitRequestGetResponseEnvelopeErrorsSource]
type accessLogJitRequestGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessLogJitRequestGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessLogJitRequestGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessLogJitRequestGetResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code" api:"required"`
	Message          string                                               `json:"message" api:"required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           AccessLogJitRequestGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessLogJitRequestGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessLogJitRequestGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessLogJitRequestGetResponseEnvelopeMessages]
type accessLogJitRequestGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessLogJitRequestGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessLogJitRequestGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessLogJitRequestGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    accessLogJitRequestGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessLogJitRequestGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [AccessLogJitRequestGetResponseEnvelopeMessagesSource]
type accessLogJitRequestGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessLogJitRequestGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessLogJitRequestGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessLogJitRequestGetResponseEnvelopeSuccess bool

const (
	AccessLogJitRequestGetResponseEnvelopeSuccessTrue AccessLogJitRequestGetResponseEnvelopeSuccess = true
)

func (r AccessLogJitRequestGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessLogJitRequestGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
