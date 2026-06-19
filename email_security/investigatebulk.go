// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/tidwall/gjson"
)

// InvestigateBulkService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvestigateBulkService] method instead.
type InvestigateBulkService struct {
	Options  []option.RequestOption
	Cancel   *InvestigateBulkCancelService
	Messages *InvestigateBulkMessageService
}

// NewInvestigateBulkService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInvestigateBulkService(opts ...option.RequestOption) (r *InvestigateBulkService) {
	r = &InvestigateBulkService{}
	r.Options = opts
	r.Cancel = NewInvestigateBulkCancelService(opts...)
	r.Messages = NewInvestigateBulkMessageService(opts...)
	return
}

// Create a bulk action job
func (r *InvestigateBulkService) New(ctx context.Context, params InvestigateBulkNewParams, opts ...option.RequestOption) (res *InvestigateBulkNewResponse, err error) {
	var env InvestigateBulkNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/bulk", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List bulk action jobs
func (r *InvestigateBulkService) List(ctx context.Context, params InvestigateBulkListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[InvestigateBulkListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/bulk", params.AccountID)
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

// List bulk action jobs
func (r *InvestigateBulkService) ListAutoPaging(ctx context.Context, params InvestigateBulkListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[InvestigateBulkListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Soft-deletes the job, hiding it from all list and detail endpoints. Only jobs in
// a terminal state (`COMPLETED`, `CANCELLED`, `FAILED`, or `SKIPPED`) can be
// deleted. To stop an in-progress job without removing it, use the cancel endpoint
// instead.
func (r *InvestigateBulkService) Delete(ctx context.Context, jobID string, body InvestigateBulkDeleteParams, opts ...option.RequestOption) (res *InvestigateBulkDeleteResponse, err error) {
	var env InvestigateBulkDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/bulk/%s", body.AccountID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Get bulk action job details
func (r *InvestigateBulkService) Get(ctx context.Context, jobID string, query InvestigateBulkGetParams, opts ...option.RequestOption) (res *InvestigateBulkGetResponse, err error) {
	var env InvestigateBulkGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/bulk/%s", query.AccountID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type InvestigateBulkNewResponse struct {
	ActionParams            InvestigateBulkNewResponseActionParams `json:"action_params" api:"required"`
	ActionType              InvestigateBulkNewResponseActionType   `json:"action_type" api:"required"`
	CreatedAt               time.Time                              `json:"created_at" api:"required" format:"date-time"`
	JobID                   string                                 `json:"job_id" api:"required" format:"uuid"`
	MessagesFailed          int64                                  `json:"messages_failed" api:"required"`
	MessagesPending         int64                                  `json:"messages_pending" api:"required"`
	MessagesSuccessful      int64                                  `json:"messages_successful" api:"required"`
	SearchParams            InvestigateBulkNewResponseSearchParams `json:"search_params" api:"required"`
	Status                  InvestigateBulkNewResponseStatus       `json:"status" api:"required"`
	TotalMessagesDiscovered int64                                  `json:"total_messages_discovered" api:"required"`
	Comment                 string                                 `json:"comment" api:"nullable"`
	CompletedAt             time.Time                              `json:"completed_at" api:"nullable" format:"date-time"`
	StartedAt               time.Time                              `json:"started_at" api:"nullable" format:"date-time"`
	StatusMessage           string                                 `json:"status_message" api:"nullable"`
	JSON                    investigateBulkNewResponseJSON         `json:"-"`
}

// investigateBulkNewResponseJSON contains the JSON metadata for the struct
// [InvestigateBulkNewResponse]
type investigateBulkNewResponseJSON struct {
	ActionParams            apijson.Field
	ActionType              apijson.Field
	CreatedAt               apijson.Field
	JobID                   apijson.Field
	MessagesFailed          apijson.Field
	MessagesPending         apijson.Field
	MessagesSuccessful      apijson.Field
	SearchParams            apijson.Field
	Status                  apijson.Field
	TotalMessagesDiscovered apijson.Field
	Comment                 apijson.Field
	CompletedAt             apijson.Field
	StartedAt               apijson.Field
	StatusMessage           apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *InvestigateBulkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkNewResponseJSON) RawJSON() string {
	return r.raw
}

type InvestigateBulkNewResponseActionParams struct {
	Type                InvestigateBulkNewResponseActionParamsType                `json:"type" api:"required"`
	Destination         InvestigateBulkNewResponseActionParamsDestination         `json:"destination"`
	ExpectedDisposition InvestigateBulkNewResponseActionParamsExpectedDisposition `json:"expected_disposition"`
	JSON                investigateBulkNewResponseActionParamsJSON                `json:"-"`
	union               InvestigateBulkNewResponseActionParamsUnion
}

// investigateBulkNewResponseActionParamsJSON contains the JSON metadata for the
// struct [InvestigateBulkNewResponseActionParams]
type investigateBulkNewResponseActionParamsJSON struct {
	Type                apijson.Field
	Destination         apijson.Field
	ExpectedDisposition apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r investigateBulkNewResponseActionParamsJSON) RawJSON() string {
	return r.raw
}

func (r *InvestigateBulkNewResponseActionParams) UnmarshalJSON(data []byte) (err error) {
	*r = InvestigateBulkNewResponseActionParams{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [InvestigateBulkNewResponseActionParamsUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [InvestigateBulkNewResponseActionParamsMove],
// [InvestigateBulkNewResponseActionParamsRelease].
func (r InvestigateBulkNewResponseActionParams) AsUnion() InvestigateBulkNewResponseActionParamsUnion {
	return r.union
}

// Union satisfied by [InvestigateBulkNewResponseActionParamsMove] or
// [InvestigateBulkNewResponseActionParamsRelease].
type InvestigateBulkNewResponseActionParamsUnion interface {
	implementsInvestigateBulkNewResponseActionParams()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InvestigateBulkNewResponseActionParamsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvestigateBulkNewResponseActionParamsMove{}),
			DiscriminatorValue: "MOVE",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvestigateBulkNewResponseActionParamsRelease{}),
			DiscriminatorValue: "RELEASE",
		},
	)
}

type InvestigateBulkNewResponseActionParamsMove struct {
	Destination         InvestigateBulkNewResponseActionParamsMoveDestination         `json:"destination" api:"required"`
	Type                InvestigateBulkNewResponseActionParamsMoveType                `json:"type" api:"required"`
	ExpectedDisposition InvestigateBulkNewResponseActionParamsMoveExpectedDisposition `json:"expected_disposition"`
	JSON                investigateBulkNewResponseActionParamsMoveJSON                `json:"-"`
}

// investigateBulkNewResponseActionParamsMoveJSON contains the JSON metadata for
// the struct [InvestigateBulkNewResponseActionParamsMove]
type investigateBulkNewResponseActionParamsMoveJSON struct {
	Destination         apijson.Field
	Type                apijson.Field
	ExpectedDisposition apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *InvestigateBulkNewResponseActionParamsMove) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkNewResponseActionParamsMoveJSON) RawJSON() string {
	return r.raw
}

func (r InvestigateBulkNewResponseActionParamsMove) implementsInvestigateBulkNewResponseActionParams() {
}

type InvestigateBulkNewResponseActionParamsMoveDestination string

const (
	InvestigateBulkNewResponseActionParamsMoveDestinationInbox                     InvestigateBulkNewResponseActionParamsMoveDestination = "Inbox"
	InvestigateBulkNewResponseActionParamsMoveDestinationJunkEmail                 InvestigateBulkNewResponseActionParamsMoveDestination = "JunkEmail"
	InvestigateBulkNewResponseActionParamsMoveDestinationDeletedItems              InvestigateBulkNewResponseActionParamsMoveDestination = "DeletedItems"
	InvestigateBulkNewResponseActionParamsMoveDestinationRecoverableItemsDeletions InvestigateBulkNewResponseActionParamsMoveDestination = "RecoverableItemsDeletions"
	InvestigateBulkNewResponseActionParamsMoveDestinationRecoverableItemsPurges    InvestigateBulkNewResponseActionParamsMoveDestination = "RecoverableItemsPurges"
)

func (r InvestigateBulkNewResponseActionParamsMoveDestination) IsKnown() bool {
	switch r {
	case InvestigateBulkNewResponseActionParamsMoveDestinationInbox, InvestigateBulkNewResponseActionParamsMoveDestinationJunkEmail, InvestigateBulkNewResponseActionParamsMoveDestinationDeletedItems, InvestigateBulkNewResponseActionParamsMoveDestinationRecoverableItemsDeletions, InvestigateBulkNewResponseActionParamsMoveDestinationRecoverableItemsPurges:
		return true
	}
	return false
}

type InvestigateBulkNewResponseActionParamsMoveType string

const (
	InvestigateBulkNewResponseActionParamsMoveTypeMove InvestigateBulkNewResponseActionParamsMoveType = "MOVE"
)

func (r InvestigateBulkNewResponseActionParamsMoveType) IsKnown() bool {
	switch r {
	case InvestigateBulkNewResponseActionParamsMoveTypeMove:
		return true
	}
	return false
}

type InvestigateBulkNewResponseActionParamsMoveExpectedDisposition string

const (
	InvestigateBulkNewResponseActionParamsMoveExpectedDispositionMalicious    InvestigateBulkNewResponseActionParamsMoveExpectedDisposition = "MALICIOUS"
	InvestigateBulkNewResponseActionParamsMoveExpectedDispositionMaliciousBec InvestigateBulkNewResponseActionParamsMoveExpectedDisposition = "MALICIOUS-BEC"
	InvestigateBulkNewResponseActionParamsMoveExpectedDispositionSuspicious   InvestigateBulkNewResponseActionParamsMoveExpectedDisposition = "SUSPICIOUS"
	InvestigateBulkNewResponseActionParamsMoveExpectedDispositionSpoof        InvestigateBulkNewResponseActionParamsMoveExpectedDisposition = "SPOOF"
	InvestigateBulkNewResponseActionParamsMoveExpectedDispositionSpam         InvestigateBulkNewResponseActionParamsMoveExpectedDisposition = "SPAM"
	InvestigateBulkNewResponseActionParamsMoveExpectedDispositionBulk         InvestigateBulkNewResponseActionParamsMoveExpectedDisposition = "BULK"
	InvestigateBulkNewResponseActionParamsMoveExpectedDispositionEncrypted    InvestigateBulkNewResponseActionParamsMoveExpectedDisposition = "ENCRYPTED"
	InvestigateBulkNewResponseActionParamsMoveExpectedDispositionExternal     InvestigateBulkNewResponseActionParamsMoveExpectedDisposition = "EXTERNAL"
	InvestigateBulkNewResponseActionParamsMoveExpectedDispositionUnknown      InvestigateBulkNewResponseActionParamsMoveExpectedDisposition = "UNKNOWN"
	InvestigateBulkNewResponseActionParamsMoveExpectedDispositionNone         InvestigateBulkNewResponseActionParamsMoveExpectedDisposition = "NONE"
)

func (r InvestigateBulkNewResponseActionParamsMoveExpectedDisposition) IsKnown() bool {
	switch r {
	case InvestigateBulkNewResponseActionParamsMoveExpectedDispositionMalicious, InvestigateBulkNewResponseActionParamsMoveExpectedDispositionMaliciousBec, InvestigateBulkNewResponseActionParamsMoveExpectedDispositionSuspicious, InvestigateBulkNewResponseActionParamsMoveExpectedDispositionSpoof, InvestigateBulkNewResponseActionParamsMoveExpectedDispositionSpam, InvestigateBulkNewResponseActionParamsMoveExpectedDispositionBulk, InvestigateBulkNewResponseActionParamsMoveExpectedDispositionEncrypted, InvestigateBulkNewResponseActionParamsMoveExpectedDispositionExternal, InvestigateBulkNewResponseActionParamsMoveExpectedDispositionUnknown, InvestigateBulkNewResponseActionParamsMoveExpectedDispositionNone:
		return true
	}
	return false
}

type InvestigateBulkNewResponseActionParamsRelease struct {
	Type InvestigateBulkNewResponseActionParamsReleaseType `json:"type" api:"required"`
	JSON investigateBulkNewResponseActionParamsReleaseJSON `json:"-"`
}

// investigateBulkNewResponseActionParamsReleaseJSON contains the JSON metadata for
// the struct [InvestigateBulkNewResponseActionParamsRelease]
type investigateBulkNewResponseActionParamsReleaseJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateBulkNewResponseActionParamsRelease) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkNewResponseActionParamsReleaseJSON) RawJSON() string {
	return r.raw
}

func (r InvestigateBulkNewResponseActionParamsRelease) implementsInvestigateBulkNewResponseActionParams() {
}

type InvestigateBulkNewResponseActionParamsReleaseType string

const (
	InvestigateBulkNewResponseActionParamsReleaseTypeRelease InvestigateBulkNewResponseActionParamsReleaseType = "RELEASE"
)

func (r InvestigateBulkNewResponseActionParamsReleaseType) IsKnown() bool {
	switch r {
	case InvestigateBulkNewResponseActionParamsReleaseTypeRelease:
		return true
	}
	return false
}

type InvestigateBulkNewResponseActionParamsType string

const (
	InvestigateBulkNewResponseActionParamsTypeMove    InvestigateBulkNewResponseActionParamsType = "MOVE"
	InvestigateBulkNewResponseActionParamsTypeRelease InvestigateBulkNewResponseActionParamsType = "RELEASE"
)

func (r InvestigateBulkNewResponseActionParamsType) IsKnown() bool {
	switch r {
	case InvestigateBulkNewResponseActionParamsTypeMove, InvestigateBulkNewResponseActionParamsTypeRelease:
		return true
	}
	return false
}

type InvestigateBulkNewResponseActionParamsDestination string

const (
	InvestigateBulkNewResponseActionParamsDestinationInbox                     InvestigateBulkNewResponseActionParamsDestination = "Inbox"
	InvestigateBulkNewResponseActionParamsDestinationJunkEmail                 InvestigateBulkNewResponseActionParamsDestination = "JunkEmail"
	InvestigateBulkNewResponseActionParamsDestinationDeletedItems              InvestigateBulkNewResponseActionParamsDestination = "DeletedItems"
	InvestigateBulkNewResponseActionParamsDestinationRecoverableItemsDeletions InvestigateBulkNewResponseActionParamsDestination = "RecoverableItemsDeletions"
	InvestigateBulkNewResponseActionParamsDestinationRecoverableItemsPurges    InvestigateBulkNewResponseActionParamsDestination = "RecoverableItemsPurges"
)

func (r InvestigateBulkNewResponseActionParamsDestination) IsKnown() bool {
	switch r {
	case InvestigateBulkNewResponseActionParamsDestinationInbox, InvestigateBulkNewResponseActionParamsDestinationJunkEmail, InvestigateBulkNewResponseActionParamsDestinationDeletedItems, InvestigateBulkNewResponseActionParamsDestinationRecoverableItemsDeletions, InvestigateBulkNewResponseActionParamsDestinationRecoverableItemsPurges:
		return true
	}
	return false
}

type InvestigateBulkNewResponseActionParamsExpectedDisposition string

const (
	InvestigateBulkNewResponseActionParamsExpectedDispositionMalicious    InvestigateBulkNewResponseActionParamsExpectedDisposition = "MALICIOUS"
	InvestigateBulkNewResponseActionParamsExpectedDispositionMaliciousBec InvestigateBulkNewResponseActionParamsExpectedDisposition = "MALICIOUS-BEC"
	InvestigateBulkNewResponseActionParamsExpectedDispositionSuspicious   InvestigateBulkNewResponseActionParamsExpectedDisposition = "SUSPICIOUS"
	InvestigateBulkNewResponseActionParamsExpectedDispositionSpoof        InvestigateBulkNewResponseActionParamsExpectedDisposition = "SPOOF"
	InvestigateBulkNewResponseActionParamsExpectedDispositionSpam         InvestigateBulkNewResponseActionParamsExpectedDisposition = "SPAM"
	InvestigateBulkNewResponseActionParamsExpectedDispositionBulk         InvestigateBulkNewResponseActionParamsExpectedDisposition = "BULK"
	InvestigateBulkNewResponseActionParamsExpectedDispositionEncrypted    InvestigateBulkNewResponseActionParamsExpectedDisposition = "ENCRYPTED"
	InvestigateBulkNewResponseActionParamsExpectedDispositionExternal     InvestigateBulkNewResponseActionParamsExpectedDisposition = "EXTERNAL"
	InvestigateBulkNewResponseActionParamsExpectedDispositionUnknown      InvestigateBulkNewResponseActionParamsExpectedDisposition = "UNKNOWN"
	InvestigateBulkNewResponseActionParamsExpectedDispositionNone         InvestigateBulkNewResponseActionParamsExpectedDisposition = "NONE"
)

func (r InvestigateBulkNewResponseActionParamsExpectedDisposition) IsKnown() bool {
	switch r {
	case InvestigateBulkNewResponseActionParamsExpectedDispositionMalicious, InvestigateBulkNewResponseActionParamsExpectedDispositionMaliciousBec, InvestigateBulkNewResponseActionParamsExpectedDispositionSuspicious, InvestigateBulkNewResponseActionParamsExpectedDispositionSpoof, InvestigateBulkNewResponseActionParamsExpectedDispositionSpam, InvestigateBulkNewResponseActionParamsExpectedDispositionBulk, InvestigateBulkNewResponseActionParamsExpectedDispositionEncrypted, InvestigateBulkNewResponseActionParamsExpectedDispositionExternal, InvestigateBulkNewResponseActionParamsExpectedDispositionUnknown, InvestigateBulkNewResponseActionParamsExpectedDispositionNone:
		return true
	}
	return false
}

type InvestigateBulkNewResponseActionType string

const (
	InvestigateBulkNewResponseActionTypeMove    InvestigateBulkNewResponseActionType = "MOVE"
	InvestigateBulkNewResponseActionTypeRelease InvestigateBulkNewResponseActionType = "RELEASE"
)

func (r InvestigateBulkNewResponseActionType) IsKnown() bool {
	switch r {
	case InvestigateBulkNewResponseActionTypeMove, InvestigateBulkNewResponseActionTypeRelease:
		return true
	}
	return false
}

type InvestigateBulkNewResponseSearchParams struct {
	// Deprecated, use `GET /investigate/{investigate_id}/action_log` instead. End of
	// life: November 1, 2026.
	//
	// Deprecated: deprecated
	ActionLog bool   `json:"action_log"`
	AlertID   string `json:"alert_id" api:"nullable"`
	// Delivery status of the message.
	DeliveryStatus InvestigateBulkNewResponseSearchParamsDeliveryStatus `json:"delivery_status"`
	DetectionsOnly bool                                                 `json:"detections_only"`
	Domain         string                                               `json:"domain" api:"nullable"`
	// End of search date range
	End              time.Time                                              `json:"end" format:"date-time"`
	ExactSubject     string                                                 `json:"exact_subject" api:"nullable"`
	FinalDisposition InvestigateBulkNewResponseSearchParamsFinalDisposition `json:"final_disposition"`
	MessageAction    InvestigateBulkNewResponseSearchParamsMessageAction    `json:"message_action" api:"nullable"`
	MessageID        string                                                 `json:"message_id" api:"nullable"`
	Metric           string                                                 `json:"metric" api:"nullable"`
	Query            string                                                 `json:"query" api:"nullable"`
	Recipient        string                                                 `json:"recipient" api:"nullable"`
	Sender           string                                                 `json:"sender" api:"nullable"`
	// Beginning of search date range
	Start       time.Time                                  `json:"start" format:"date-time"`
	Subject     string                                     `json:"subject" api:"nullable"`
	Submissions bool                                       `json:"submissions"`
	JSON        investigateBulkNewResponseSearchParamsJSON `json:"-"`
}

// investigateBulkNewResponseSearchParamsJSON contains the JSON metadata for the
// struct [InvestigateBulkNewResponseSearchParams]
type investigateBulkNewResponseSearchParamsJSON struct {
	ActionLog        apijson.Field
	AlertID          apijson.Field
	DeliveryStatus   apijson.Field
	DetectionsOnly   apijson.Field
	Domain           apijson.Field
	End              apijson.Field
	ExactSubject     apijson.Field
	FinalDisposition apijson.Field
	MessageAction    apijson.Field
	MessageID        apijson.Field
	Metric           apijson.Field
	Query            apijson.Field
	Recipient        apijson.Field
	Sender           apijson.Field
	Start            apijson.Field
	Subject          apijson.Field
	Submissions      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvestigateBulkNewResponseSearchParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkNewResponseSearchParamsJSON) RawJSON() string {
	return r.raw
}

// Delivery status of the message.
type InvestigateBulkNewResponseSearchParamsDeliveryStatus string

const (
	InvestigateBulkNewResponseSearchParamsDeliveryStatusDelivered   InvestigateBulkNewResponseSearchParamsDeliveryStatus = "delivered"
	InvestigateBulkNewResponseSearchParamsDeliveryStatusMoved       InvestigateBulkNewResponseSearchParamsDeliveryStatus = "moved"
	InvestigateBulkNewResponseSearchParamsDeliveryStatusQuarantined InvestigateBulkNewResponseSearchParamsDeliveryStatus = "quarantined"
	InvestigateBulkNewResponseSearchParamsDeliveryStatusRejected    InvestigateBulkNewResponseSearchParamsDeliveryStatus = "rejected"
	InvestigateBulkNewResponseSearchParamsDeliveryStatusDeferred    InvestigateBulkNewResponseSearchParamsDeliveryStatus = "deferred"
	InvestigateBulkNewResponseSearchParamsDeliveryStatusBounced     InvestigateBulkNewResponseSearchParamsDeliveryStatus = "bounced"
	InvestigateBulkNewResponseSearchParamsDeliveryStatusQueued      InvestigateBulkNewResponseSearchParamsDeliveryStatus = "queued"
)

func (r InvestigateBulkNewResponseSearchParamsDeliveryStatus) IsKnown() bool {
	switch r {
	case InvestigateBulkNewResponseSearchParamsDeliveryStatusDelivered, InvestigateBulkNewResponseSearchParamsDeliveryStatusMoved, InvestigateBulkNewResponseSearchParamsDeliveryStatusQuarantined, InvestigateBulkNewResponseSearchParamsDeliveryStatusRejected, InvestigateBulkNewResponseSearchParamsDeliveryStatusDeferred, InvestigateBulkNewResponseSearchParamsDeliveryStatusBounced, InvestigateBulkNewResponseSearchParamsDeliveryStatusQueued:
		return true
	}
	return false
}

type InvestigateBulkNewResponseSearchParamsFinalDisposition string

const (
	InvestigateBulkNewResponseSearchParamsFinalDispositionMalicious    InvestigateBulkNewResponseSearchParamsFinalDisposition = "MALICIOUS"
	InvestigateBulkNewResponseSearchParamsFinalDispositionMaliciousBec InvestigateBulkNewResponseSearchParamsFinalDisposition = "MALICIOUS-BEC"
	InvestigateBulkNewResponseSearchParamsFinalDispositionSuspicious   InvestigateBulkNewResponseSearchParamsFinalDisposition = "SUSPICIOUS"
	InvestigateBulkNewResponseSearchParamsFinalDispositionSpoof        InvestigateBulkNewResponseSearchParamsFinalDisposition = "SPOOF"
	InvestigateBulkNewResponseSearchParamsFinalDispositionSpam         InvestigateBulkNewResponseSearchParamsFinalDisposition = "SPAM"
	InvestigateBulkNewResponseSearchParamsFinalDispositionBulk         InvestigateBulkNewResponseSearchParamsFinalDisposition = "BULK"
	InvestigateBulkNewResponseSearchParamsFinalDispositionEncrypted    InvestigateBulkNewResponseSearchParamsFinalDisposition = "ENCRYPTED"
	InvestigateBulkNewResponseSearchParamsFinalDispositionExternal     InvestigateBulkNewResponseSearchParamsFinalDisposition = "EXTERNAL"
	InvestigateBulkNewResponseSearchParamsFinalDispositionUnknown      InvestigateBulkNewResponseSearchParamsFinalDisposition = "UNKNOWN"
	InvestigateBulkNewResponseSearchParamsFinalDispositionNone         InvestigateBulkNewResponseSearchParamsFinalDisposition = "NONE"
)

func (r InvestigateBulkNewResponseSearchParamsFinalDisposition) IsKnown() bool {
	switch r {
	case InvestigateBulkNewResponseSearchParamsFinalDispositionMalicious, InvestigateBulkNewResponseSearchParamsFinalDispositionMaliciousBec, InvestigateBulkNewResponseSearchParamsFinalDispositionSuspicious, InvestigateBulkNewResponseSearchParamsFinalDispositionSpoof, InvestigateBulkNewResponseSearchParamsFinalDispositionSpam, InvestigateBulkNewResponseSearchParamsFinalDispositionBulk, InvestigateBulkNewResponseSearchParamsFinalDispositionEncrypted, InvestigateBulkNewResponseSearchParamsFinalDispositionExternal, InvestigateBulkNewResponseSearchParamsFinalDispositionUnknown, InvestigateBulkNewResponseSearchParamsFinalDispositionNone:
		return true
	}
	return false
}

type InvestigateBulkNewResponseSearchParamsMessageAction string

const (
	InvestigateBulkNewResponseSearchParamsMessageActionPreview            InvestigateBulkNewResponseSearchParamsMessageAction = "PREVIEW"
	InvestigateBulkNewResponseSearchParamsMessageActionQuarantineReleased InvestigateBulkNewResponseSearchParamsMessageAction = "QUARANTINE_RELEASED"
	InvestigateBulkNewResponseSearchParamsMessageActionMoved              InvestigateBulkNewResponseSearchParamsMessageAction = "MOVED"
)

func (r InvestigateBulkNewResponseSearchParamsMessageAction) IsKnown() bool {
	switch r {
	case InvestigateBulkNewResponseSearchParamsMessageActionPreview, InvestigateBulkNewResponseSearchParamsMessageActionQuarantineReleased, InvestigateBulkNewResponseSearchParamsMessageActionMoved:
		return true
	}
	return false
}

type InvestigateBulkNewResponseStatus string

const (
	InvestigateBulkNewResponseStatusPending     InvestigateBulkNewResponseStatus = "PENDING"
	InvestigateBulkNewResponseStatusDiscovering InvestigateBulkNewResponseStatus = "DISCOVERING"
	InvestigateBulkNewResponseStatusProcessing  InvestigateBulkNewResponseStatus = "PROCESSING"
	InvestigateBulkNewResponseStatusCompleted   InvestigateBulkNewResponseStatus = "COMPLETED"
	InvestigateBulkNewResponseStatusFailed      InvestigateBulkNewResponseStatus = "FAILED"
	InvestigateBulkNewResponseStatusCancelled   InvestigateBulkNewResponseStatus = "CANCELLED"
	InvestigateBulkNewResponseStatusSkipped     InvestigateBulkNewResponseStatus = "SKIPPED"
)

func (r InvestigateBulkNewResponseStatus) IsKnown() bool {
	switch r {
	case InvestigateBulkNewResponseStatusPending, InvestigateBulkNewResponseStatusDiscovering, InvestigateBulkNewResponseStatusProcessing, InvestigateBulkNewResponseStatusCompleted, InvestigateBulkNewResponseStatusFailed, InvestigateBulkNewResponseStatusCancelled, InvestigateBulkNewResponseStatusSkipped:
		return true
	}
	return false
}

type InvestigateBulkListResponse struct {
	ActionParams            InvestigateBulkListResponseActionParams `json:"action_params" api:"required"`
	ActionType              InvestigateBulkListResponseActionType   `json:"action_type" api:"required"`
	CreatedAt               time.Time                               `json:"created_at" api:"required" format:"date-time"`
	JobID                   string                                  `json:"job_id" api:"required" format:"uuid"`
	MessagesFailed          int64                                   `json:"messages_failed" api:"required"`
	MessagesPending         int64                                   `json:"messages_pending" api:"required"`
	MessagesSuccessful      int64                                   `json:"messages_successful" api:"required"`
	SearchParams            InvestigateBulkListResponseSearchParams `json:"search_params" api:"required"`
	Status                  InvestigateBulkListResponseStatus       `json:"status" api:"required"`
	TotalMessagesDiscovered int64                                   `json:"total_messages_discovered" api:"required"`
	Comment                 string                                  `json:"comment" api:"nullable"`
	CompletedAt             time.Time                               `json:"completed_at" api:"nullable" format:"date-time"`
	StartedAt               time.Time                               `json:"started_at" api:"nullable" format:"date-time"`
	StatusMessage           string                                  `json:"status_message" api:"nullable"`
	JSON                    investigateBulkListResponseJSON         `json:"-"`
}

// investigateBulkListResponseJSON contains the JSON metadata for the struct
// [InvestigateBulkListResponse]
type investigateBulkListResponseJSON struct {
	ActionParams            apijson.Field
	ActionType              apijson.Field
	CreatedAt               apijson.Field
	JobID                   apijson.Field
	MessagesFailed          apijson.Field
	MessagesPending         apijson.Field
	MessagesSuccessful      apijson.Field
	SearchParams            apijson.Field
	Status                  apijson.Field
	TotalMessagesDiscovered apijson.Field
	Comment                 apijson.Field
	CompletedAt             apijson.Field
	StartedAt               apijson.Field
	StatusMessage           apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *InvestigateBulkListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkListResponseJSON) RawJSON() string {
	return r.raw
}

type InvestigateBulkListResponseActionParams struct {
	Type                InvestigateBulkListResponseActionParamsType                `json:"type" api:"required"`
	Destination         InvestigateBulkListResponseActionParamsDestination         `json:"destination"`
	ExpectedDisposition InvestigateBulkListResponseActionParamsExpectedDisposition `json:"expected_disposition"`
	JSON                investigateBulkListResponseActionParamsJSON                `json:"-"`
	union               InvestigateBulkListResponseActionParamsUnion
}

// investigateBulkListResponseActionParamsJSON contains the JSON metadata for the
// struct [InvestigateBulkListResponseActionParams]
type investigateBulkListResponseActionParamsJSON struct {
	Type                apijson.Field
	Destination         apijson.Field
	ExpectedDisposition apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r investigateBulkListResponseActionParamsJSON) RawJSON() string {
	return r.raw
}

func (r *InvestigateBulkListResponseActionParams) UnmarshalJSON(data []byte) (err error) {
	*r = InvestigateBulkListResponseActionParams{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [InvestigateBulkListResponseActionParamsUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [InvestigateBulkListResponseActionParamsMove],
// [InvestigateBulkListResponseActionParamsRelease].
func (r InvestigateBulkListResponseActionParams) AsUnion() InvestigateBulkListResponseActionParamsUnion {
	return r.union
}

// Union satisfied by [InvestigateBulkListResponseActionParamsMove] or
// [InvestigateBulkListResponseActionParamsRelease].
type InvestigateBulkListResponseActionParamsUnion interface {
	implementsInvestigateBulkListResponseActionParams()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InvestigateBulkListResponseActionParamsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvestigateBulkListResponseActionParamsMove{}),
			DiscriminatorValue: "MOVE",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvestigateBulkListResponseActionParamsRelease{}),
			DiscriminatorValue: "RELEASE",
		},
	)
}

type InvestigateBulkListResponseActionParamsMove struct {
	Destination         InvestigateBulkListResponseActionParamsMoveDestination         `json:"destination" api:"required"`
	Type                InvestigateBulkListResponseActionParamsMoveType                `json:"type" api:"required"`
	ExpectedDisposition InvestigateBulkListResponseActionParamsMoveExpectedDisposition `json:"expected_disposition"`
	JSON                investigateBulkListResponseActionParamsMoveJSON                `json:"-"`
}

// investigateBulkListResponseActionParamsMoveJSON contains the JSON metadata for
// the struct [InvestigateBulkListResponseActionParamsMove]
type investigateBulkListResponseActionParamsMoveJSON struct {
	Destination         apijson.Field
	Type                apijson.Field
	ExpectedDisposition apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *InvestigateBulkListResponseActionParamsMove) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkListResponseActionParamsMoveJSON) RawJSON() string {
	return r.raw
}

func (r InvestigateBulkListResponseActionParamsMove) implementsInvestigateBulkListResponseActionParams() {
}

type InvestigateBulkListResponseActionParamsMoveDestination string

const (
	InvestigateBulkListResponseActionParamsMoveDestinationInbox                     InvestigateBulkListResponseActionParamsMoveDestination = "Inbox"
	InvestigateBulkListResponseActionParamsMoveDestinationJunkEmail                 InvestigateBulkListResponseActionParamsMoveDestination = "JunkEmail"
	InvestigateBulkListResponseActionParamsMoveDestinationDeletedItems              InvestigateBulkListResponseActionParamsMoveDestination = "DeletedItems"
	InvestigateBulkListResponseActionParamsMoveDestinationRecoverableItemsDeletions InvestigateBulkListResponseActionParamsMoveDestination = "RecoverableItemsDeletions"
	InvestigateBulkListResponseActionParamsMoveDestinationRecoverableItemsPurges    InvestigateBulkListResponseActionParamsMoveDestination = "RecoverableItemsPurges"
)

func (r InvestigateBulkListResponseActionParamsMoveDestination) IsKnown() bool {
	switch r {
	case InvestigateBulkListResponseActionParamsMoveDestinationInbox, InvestigateBulkListResponseActionParamsMoveDestinationJunkEmail, InvestigateBulkListResponseActionParamsMoveDestinationDeletedItems, InvestigateBulkListResponseActionParamsMoveDestinationRecoverableItemsDeletions, InvestigateBulkListResponseActionParamsMoveDestinationRecoverableItemsPurges:
		return true
	}
	return false
}

type InvestigateBulkListResponseActionParamsMoveType string

const (
	InvestigateBulkListResponseActionParamsMoveTypeMove InvestigateBulkListResponseActionParamsMoveType = "MOVE"
)

func (r InvestigateBulkListResponseActionParamsMoveType) IsKnown() bool {
	switch r {
	case InvestigateBulkListResponseActionParamsMoveTypeMove:
		return true
	}
	return false
}

type InvestigateBulkListResponseActionParamsMoveExpectedDisposition string

const (
	InvestigateBulkListResponseActionParamsMoveExpectedDispositionMalicious    InvestigateBulkListResponseActionParamsMoveExpectedDisposition = "MALICIOUS"
	InvestigateBulkListResponseActionParamsMoveExpectedDispositionMaliciousBec InvestigateBulkListResponseActionParamsMoveExpectedDisposition = "MALICIOUS-BEC"
	InvestigateBulkListResponseActionParamsMoveExpectedDispositionSuspicious   InvestigateBulkListResponseActionParamsMoveExpectedDisposition = "SUSPICIOUS"
	InvestigateBulkListResponseActionParamsMoveExpectedDispositionSpoof        InvestigateBulkListResponseActionParamsMoveExpectedDisposition = "SPOOF"
	InvestigateBulkListResponseActionParamsMoveExpectedDispositionSpam         InvestigateBulkListResponseActionParamsMoveExpectedDisposition = "SPAM"
	InvestigateBulkListResponseActionParamsMoveExpectedDispositionBulk         InvestigateBulkListResponseActionParamsMoveExpectedDisposition = "BULK"
	InvestigateBulkListResponseActionParamsMoveExpectedDispositionEncrypted    InvestigateBulkListResponseActionParamsMoveExpectedDisposition = "ENCRYPTED"
	InvestigateBulkListResponseActionParamsMoveExpectedDispositionExternal     InvestigateBulkListResponseActionParamsMoveExpectedDisposition = "EXTERNAL"
	InvestigateBulkListResponseActionParamsMoveExpectedDispositionUnknown      InvestigateBulkListResponseActionParamsMoveExpectedDisposition = "UNKNOWN"
	InvestigateBulkListResponseActionParamsMoveExpectedDispositionNone         InvestigateBulkListResponseActionParamsMoveExpectedDisposition = "NONE"
)

func (r InvestigateBulkListResponseActionParamsMoveExpectedDisposition) IsKnown() bool {
	switch r {
	case InvestigateBulkListResponseActionParamsMoveExpectedDispositionMalicious, InvestigateBulkListResponseActionParamsMoveExpectedDispositionMaliciousBec, InvestigateBulkListResponseActionParamsMoveExpectedDispositionSuspicious, InvestigateBulkListResponseActionParamsMoveExpectedDispositionSpoof, InvestigateBulkListResponseActionParamsMoveExpectedDispositionSpam, InvestigateBulkListResponseActionParamsMoveExpectedDispositionBulk, InvestigateBulkListResponseActionParamsMoveExpectedDispositionEncrypted, InvestigateBulkListResponseActionParamsMoveExpectedDispositionExternal, InvestigateBulkListResponseActionParamsMoveExpectedDispositionUnknown, InvestigateBulkListResponseActionParamsMoveExpectedDispositionNone:
		return true
	}
	return false
}

type InvestigateBulkListResponseActionParamsRelease struct {
	Type InvestigateBulkListResponseActionParamsReleaseType `json:"type" api:"required"`
	JSON investigateBulkListResponseActionParamsReleaseJSON `json:"-"`
}

// investigateBulkListResponseActionParamsReleaseJSON contains the JSON metadata
// for the struct [InvestigateBulkListResponseActionParamsRelease]
type investigateBulkListResponseActionParamsReleaseJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateBulkListResponseActionParamsRelease) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkListResponseActionParamsReleaseJSON) RawJSON() string {
	return r.raw
}

func (r InvestigateBulkListResponseActionParamsRelease) implementsInvestigateBulkListResponseActionParams() {
}

type InvestigateBulkListResponseActionParamsReleaseType string

const (
	InvestigateBulkListResponseActionParamsReleaseTypeRelease InvestigateBulkListResponseActionParamsReleaseType = "RELEASE"
)

func (r InvestigateBulkListResponseActionParamsReleaseType) IsKnown() bool {
	switch r {
	case InvestigateBulkListResponseActionParamsReleaseTypeRelease:
		return true
	}
	return false
}

type InvestigateBulkListResponseActionParamsType string

const (
	InvestigateBulkListResponseActionParamsTypeMove    InvestigateBulkListResponseActionParamsType = "MOVE"
	InvestigateBulkListResponseActionParamsTypeRelease InvestigateBulkListResponseActionParamsType = "RELEASE"
)

func (r InvestigateBulkListResponseActionParamsType) IsKnown() bool {
	switch r {
	case InvestigateBulkListResponseActionParamsTypeMove, InvestigateBulkListResponseActionParamsTypeRelease:
		return true
	}
	return false
}

type InvestigateBulkListResponseActionParamsDestination string

const (
	InvestigateBulkListResponseActionParamsDestinationInbox                     InvestigateBulkListResponseActionParamsDestination = "Inbox"
	InvestigateBulkListResponseActionParamsDestinationJunkEmail                 InvestigateBulkListResponseActionParamsDestination = "JunkEmail"
	InvestigateBulkListResponseActionParamsDestinationDeletedItems              InvestigateBulkListResponseActionParamsDestination = "DeletedItems"
	InvestigateBulkListResponseActionParamsDestinationRecoverableItemsDeletions InvestigateBulkListResponseActionParamsDestination = "RecoverableItemsDeletions"
	InvestigateBulkListResponseActionParamsDestinationRecoverableItemsPurges    InvestigateBulkListResponseActionParamsDestination = "RecoverableItemsPurges"
)

func (r InvestigateBulkListResponseActionParamsDestination) IsKnown() bool {
	switch r {
	case InvestigateBulkListResponseActionParamsDestinationInbox, InvestigateBulkListResponseActionParamsDestinationJunkEmail, InvestigateBulkListResponseActionParamsDestinationDeletedItems, InvestigateBulkListResponseActionParamsDestinationRecoverableItemsDeletions, InvestigateBulkListResponseActionParamsDestinationRecoverableItemsPurges:
		return true
	}
	return false
}

type InvestigateBulkListResponseActionParamsExpectedDisposition string

const (
	InvestigateBulkListResponseActionParamsExpectedDispositionMalicious    InvestigateBulkListResponseActionParamsExpectedDisposition = "MALICIOUS"
	InvestigateBulkListResponseActionParamsExpectedDispositionMaliciousBec InvestigateBulkListResponseActionParamsExpectedDisposition = "MALICIOUS-BEC"
	InvestigateBulkListResponseActionParamsExpectedDispositionSuspicious   InvestigateBulkListResponseActionParamsExpectedDisposition = "SUSPICIOUS"
	InvestigateBulkListResponseActionParamsExpectedDispositionSpoof        InvestigateBulkListResponseActionParamsExpectedDisposition = "SPOOF"
	InvestigateBulkListResponseActionParamsExpectedDispositionSpam         InvestigateBulkListResponseActionParamsExpectedDisposition = "SPAM"
	InvestigateBulkListResponseActionParamsExpectedDispositionBulk         InvestigateBulkListResponseActionParamsExpectedDisposition = "BULK"
	InvestigateBulkListResponseActionParamsExpectedDispositionEncrypted    InvestigateBulkListResponseActionParamsExpectedDisposition = "ENCRYPTED"
	InvestigateBulkListResponseActionParamsExpectedDispositionExternal     InvestigateBulkListResponseActionParamsExpectedDisposition = "EXTERNAL"
	InvestigateBulkListResponseActionParamsExpectedDispositionUnknown      InvestigateBulkListResponseActionParamsExpectedDisposition = "UNKNOWN"
	InvestigateBulkListResponseActionParamsExpectedDispositionNone         InvestigateBulkListResponseActionParamsExpectedDisposition = "NONE"
)

func (r InvestigateBulkListResponseActionParamsExpectedDisposition) IsKnown() bool {
	switch r {
	case InvestigateBulkListResponseActionParamsExpectedDispositionMalicious, InvestigateBulkListResponseActionParamsExpectedDispositionMaliciousBec, InvestigateBulkListResponseActionParamsExpectedDispositionSuspicious, InvestigateBulkListResponseActionParamsExpectedDispositionSpoof, InvestigateBulkListResponseActionParamsExpectedDispositionSpam, InvestigateBulkListResponseActionParamsExpectedDispositionBulk, InvestigateBulkListResponseActionParamsExpectedDispositionEncrypted, InvestigateBulkListResponseActionParamsExpectedDispositionExternal, InvestigateBulkListResponseActionParamsExpectedDispositionUnknown, InvestigateBulkListResponseActionParamsExpectedDispositionNone:
		return true
	}
	return false
}

type InvestigateBulkListResponseActionType string

const (
	InvestigateBulkListResponseActionTypeMove    InvestigateBulkListResponseActionType = "MOVE"
	InvestigateBulkListResponseActionTypeRelease InvestigateBulkListResponseActionType = "RELEASE"
)

func (r InvestigateBulkListResponseActionType) IsKnown() bool {
	switch r {
	case InvestigateBulkListResponseActionTypeMove, InvestigateBulkListResponseActionTypeRelease:
		return true
	}
	return false
}

type InvestigateBulkListResponseSearchParams struct {
	// Deprecated, use `GET /investigate/{investigate_id}/action_log` instead. End of
	// life: November 1, 2026.
	//
	// Deprecated: deprecated
	ActionLog bool   `json:"action_log"`
	AlertID   string `json:"alert_id" api:"nullable"`
	// Delivery status of the message.
	DeliveryStatus InvestigateBulkListResponseSearchParamsDeliveryStatus `json:"delivery_status"`
	DetectionsOnly bool                                                  `json:"detections_only"`
	Domain         string                                                `json:"domain" api:"nullable"`
	// End of search date range
	End              time.Time                                               `json:"end" format:"date-time"`
	ExactSubject     string                                                  `json:"exact_subject" api:"nullable"`
	FinalDisposition InvestigateBulkListResponseSearchParamsFinalDisposition `json:"final_disposition"`
	MessageAction    InvestigateBulkListResponseSearchParamsMessageAction    `json:"message_action" api:"nullable"`
	MessageID        string                                                  `json:"message_id" api:"nullable"`
	Metric           string                                                  `json:"metric" api:"nullable"`
	Query            string                                                  `json:"query" api:"nullable"`
	Recipient        string                                                  `json:"recipient" api:"nullable"`
	Sender           string                                                  `json:"sender" api:"nullable"`
	// Beginning of search date range
	Start       time.Time                                   `json:"start" format:"date-time"`
	Subject     string                                      `json:"subject" api:"nullable"`
	Submissions bool                                        `json:"submissions"`
	JSON        investigateBulkListResponseSearchParamsJSON `json:"-"`
}

// investigateBulkListResponseSearchParamsJSON contains the JSON metadata for the
// struct [InvestigateBulkListResponseSearchParams]
type investigateBulkListResponseSearchParamsJSON struct {
	ActionLog        apijson.Field
	AlertID          apijson.Field
	DeliveryStatus   apijson.Field
	DetectionsOnly   apijson.Field
	Domain           apijson.Field
	End              apijson.Field
	ExactSubject     apijson.Field
	FinalDisposition apijson.Field
	MessageAction    apijson.Field
	MessageID        apijson.Field
	Metric           apijson.Field
	Query            apijson.Field
	Recipient        apijson.Field
	Sender           apijson.Field
	Start            apijson.Field
	Subject          apijson.Field
	Submissions      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvestigateBulkListResponseSearchParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkListResponseSearchParamsJSON) RawJSON() string {
	return r.raw
}

// Delivery status of the message.
type InvestigateBulkListResponseSearchParamsDeliveryStatus string

const (
	InvestigateBulkListResponseSearchParamsDeliveryStatusDelivered   InvestigateBulkListResponseSearchParamsDeliveryStatus = "delivered"
	InvestigateBulkListResponseSearchParamsDeliveryStatusMoved       InvestigateBulkListResponseSearchParamsDeliveryStatus = "moved"
	InvestigateBulkListResponseSearchParamsDeliveryStatusQuarantined InvestigateBulkListResponseSearchParamsDeliveryStatus = "quarantined"
	InvestigateBulkListResponseSearchParamsDeliveryStatusRejected    InvestigateBulkListResponseSearchParamsDeliveryStatus = "rejected"
	InvestigateBulkListResponseSearchParamsDeliveryStatusDeferred    InvestigateBulkListResponseSearchParamsDeliveryStatus = "deferred"
	InvestigateBulkListResponseSearchParamsDeliveryStatusBounced     InvestigateBulkListResponseSearchParamsDeliveryStatus = "bounced"
	InvestigateBulkListResponseSearchParamsDeliveryStatusQueued      InvestigateBulkListResponseSearchParamsDeliveryStatus = "queued"
)

func (r InvestigateBulkListResponseSearchParamsDeliveryStatus) IsKnown() bool {
	switch r {
	case InvestigateBulkListResponseSearchParamsDeliveryStatusDelivered, InvestigateBulkListResponseSearchParamsDeliveryStatusMoved, InvestigateBulkListResponseSearchParamsDeliveryStatusQuarantined, InvestigateBulkListResponseSearchParamsDeliveryStatusRejected, InvestigateBulkListResponseSearchParamsDeliveryStatusDeferred, InvestigateBulkListResponseSearchParamsDeliveryStatusBounced, InvestigateBulkListResponseSearchParamsDeliveryStatusQueued:
		return true
	}
	return false
}

type InvestigateBulkListResponseSearchParamsFinalDisposition string

const (
	InvestigateBulkListResponseSearchParamsFinalDispositionMalicious    InvestigateBulkListResponseSearchParamsFinalDisposition = "MALICIOUS"
	InvestigateBulkListResponseSearchParamsFinalDispositionMaliciousBec InvestigateBulkListResponseSearchParamsFinalDisposition = "MALICIOUS-BEC"
	InvestigateBulkListResponseSearchParamsFinalDispositionSuspicious   InvestigateBulkListResponseSearchParamsFinalDisposition = "SUSPICIOUS"
	InvestigateBulkListResponseSearchParamsFinalDispositionSpoof        InvestigateBulkListResponseSearchParamsFinalDisposition = "SPOOF"
	InvestigateBulkListResponseSearchParamsFinalDispositionSpam         InvestigateBulkListResponseSearchParamsFinalDisposition = "SPAM"
	InvestigateBulkListResponseSearchParamsFinalDispositionBulk         InvestigateBulkListResponseSearchParamsFinalDisposition = "BULK"
	InvestigateBulkListResponseSearchParamsFinalDispositionEncrypted    InvestigateBulkListResponseSearchParamsFinalDisposition = "ENCRYPTED"
	InvestigateBulkListResponseSearchParamsFinalDispositionExternal     InvestigateBulkListResponseSearchParamsFinalDisposition = "EXTERNAL"
	InvestigateBulkListResponseSearchParamsFinalDispositionUnknown      InvestigateBulkListResponseSearchParamsFinalDisposition = "UNKNOWN"
	InvestigateBulkListResponseSearchParamsFinalDispositionNone         InvestigateBulkListResponseSearchParamsFinalDisposition = "NONE"
)

func (r InvestigateBulkListResponseSearchParamsFinalDisposition) IsKnown() bool {
	switch r {
	case InvestigateBulkListResponseSearchParamsFinalDispositionMalicious, InvestigateBulkListResponseSearchParamsFinalDispositionMaliciousBec, InvestigateBulkListResponseSearchParamsFinalDispositionSuspicious, InvestigateBulkListResponseSearchParamsFinalDispositionSpoof, InvestigateBulkListResponseSearchParamsFinalDispositionSpam, InvestigateBulkListResponseSearchParamsFinalDispositionBulk, InvestigateBulkListResponseSearchParamsFinalDispositionEncrypted, InvestigateBulkListResponseSearchParamsFinalDispositionExternal, InvestigateBulkListResponseSearchParamsFinalDispositionUnknown, InvestigateBulkListResponseSearchParamsFinalDispositionNone:
		return true
	}
	return false
}

type InvestigateBulkListResponseSearchParamsMessageAction string

const (
	InvestigateBulkListResponseSearchParamsMessageActionPreview            InvestigateBulkListResponseSearchParamsMessageAction = "PREVIEW"
	InvestigateBulkListResponseSearchParamsMessageActionQuarantineReleased InvestigateBulkListResponseSearchParamsMessageAction = "QUARANTINE_RELEASED"
	InvestigateBulkListResponseSearchParamsMessageActionMoved              InvestigateBulkListResponseSearchParamsMessageAction = "MOVED"
)

func (r InvestigateBulkListResponseSearchParamsMessageAction) IsKnown() bool {
	switch r {
	case InvestigateBulkListResponseSearchParamsMessageActionPreview, InvestigateBulkListResponseSearchParamsMessageActionQuarantineReleased, InvestigateBulkListResponseSearchParamsMessageActionMoved:
		return true
	}
	return false
}

type InvestigateBulkListResponseStatus string

const (
	InvestigateBulkListResponseStatusPending     InvestigateBulkListResponseStatus = "PENDING"
	InvestigateBulkListResponseStatusDiscovering InvestigateBulkListResponseStatus = "DISCOVERING"
	InvestigateBulkListResponseStatusProcessing  InvestigateBulkListResponseStatus = "PROCESSING"
	InvestigateBulkListResponseStatusCompleted   InvestigateBulkListResponseStatus = "COMPLETED"
	InvestigateBulkListResponseStatusFailed      InvestigateBulkListResponseStatus = "FAILED"
	InvestigateBulkListResponseStatusCancelled   InvestigateBulkListResponseStatus = "CANCELLED"
	InvestigateBulkListResponseStatusSkipped     InvestigateBulkListResponseStatus = "SKIPPED"
)

func (r InvestigateBulkListResponseStatus) IsKnown() bool {
	switch r {
	case InvestigateBulkListResponseStatusPending, InvestigateBulkListResponseStatusDiscovering, InvestigateBulkListResponseStatusProcessing, InvestigateBulkListResponseStatusCompleted, InvestigateBulkListResponseStatusFailed, InvestigateBulkListResponseStatusCancelled, InvestigateBulkListResponseStatusSkipped:
		return true
	}
	return false
}

type InvestigateBulkDeleteResponse struct {
	ID   string                            `json:"id" api:"required" format:"uuid"`
	JSON investigateBulkDeleteResponseJSON `json:"-"`
}

// investigateBulkDeleteResponseJSON contains the JSON metadata for the struct
// [InvestigateBulkDeleteResponse]
type investigateBulkDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type InvestigateBulkGetResponse struct {
	ActionParams            InvestigateBulkGetResponseActionParams `json:"action_params" api:"required"`
	ActionType              InvestigateBulkGetResponseActionType   `json:"action_type" api:"required"`
	CreatedAt               time.Time                              `json:"created_at" api:"required" format:"date-time"`
	JobID                   string                                 `json:"job_id" api:"required" format:"uuid"`
	MessagesFailed          int64                                  `json:"messages_failed" api:"required"`
	MessagesPending         int64                                  `json:"messages_pending" api:"required"`
	MessagesSuccessful      int64                                  `json:"messages_successful" api:"required"`
	SearchParams            InvestigateBulkGetResponseSearchParams `json:"search_params" api:"required"`
	Status                  InvestigateBulkGetResponseStatus       `json:"status" api:"required"`
	TotalMessagesDiscovered int64                                  `json:"total_messages_discovered" api:"required"`
	Comment                 string                                 `json:"comment" api:"nullable"`
	CompletedAt             time.Time                              `json:"completed_at" api:"nullable" format:"date-time"`
	StartedAt               time.Time                              `json:"started_at" api:"nullable" format:"date-time"`
	StatusMessage           string                                 `json:"status_message" api:"nullable"`
	JSON                    investigateBulkGetResponseJSON         `json:"-"`
}

// investigateBulkGetResponseJSON contains the JSON metadata for the struct
// [InvestigateBulkGetResponse]
type investigateBulkGetResponseJSON struct {
	ActionParams            apijson.Field
	ActionType              apijson.Field
	CreatedAt               apijson.Field
	JobID                   apijson.Field
	MessagesFailed          apijson.Field
	MessagesPending         apijson.Field
	MessagesSuccessful      apijson.Field
	SearchParams            apijson.Field
	Status                  apijson.Field
	TotalMessagesDiscovered apijson.Field
	Comment                 apijson.Field
	CompletedAt             apijson.Field
	StartedAt               apijson.Field
	StatusMessage           apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *InvestigateBulkGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkGetResponseJSON) RawJSON() string {
	return r.raw
}

type InvestigateBulkGetResponseActionParams struct {
	Type                InvestigateBulkGetResponseActionParamsType                `json:"type" api:"required"`
	Destination         InvestigateBulkGetResponseActionParamsDestination         `json:"destination"`
	ExpectedDisposition InvestigateBulkGetResponseActionParamsExpectedDisposition `json:"expected_disposition"`
	JSON                investigateBulkGetResponseActionParamsJSON                `json:"-"`
	union               InvestigateBulkGetResponseActionParamsUnion
}

// investigateBulkGetResponseActionParamsJSON contains the JSON metadata for the
// struct [InvestigateBulkGetResponseActionParams]
type investigateBulkGetResponseActionParamsJSON struct {
	Type                apijson.Field
	Destination         apijson.Field
	ExpectedDisposition apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r investigateBulkGetResponseActionParamsJSON) RawJSON() string {
	return r.raw
}

func (r *InvestigateBulkGetResponseActionParams) UnmarshalJSON(data []byte) (err error) {
	*r = InvestigateBulkGetResponseActionParams{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [InvestigateBulkGetResponseActionParamsUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [InvestigateBulkGetResponseActionParamsMove],
// [InvestigateBulkGetResponseActionParamsRelease].
func (r InvestigateBulkGetResponseActionParams) AsUnion() InvestigateBulkGetResponseActionParamsUnion {
	return r.union
}

// Union satisfied by [InvestigateBulkGetResponseActionParamsMove] or
// [InvestigateBulkGetResponseActionParamsRelease].
type InvestigateBulkGetResponseActionParamsUnion interface {
	implementsInvestigateBulkGetResponseActionParams()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InvestigateBulkGetResponseActionParamsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvestigateBulkGetResponseActionParamsMove{}),
			DiscriminatorValue: "MOVE",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvestigateBulkGetResponseActionParamsRelease{}),
			DiscriminatorValue: "RELEASE",
		},
	)
}

type InvestigateBulkGetResponseActionParamsMove struct {
	Destination         InvestigateBulkGetResponseActionParamsMoveDestination         `json:"destination" api:"required"`
	Type                InvestigateBulkGetResponseActionParamsMoveType                `json:"type" api:"required"`
	ExpectedDisposition InvestigateBulkGetResponseActionParamsMoveExpectedDisposition `json:"expected_disposition"`
	JSON                investigateBulkGetResponseActionParamsMoveJSON                `json:"-"`
}

// investigateBulkGetResponseActionParamsMoveJSON contains the JSON metadata for
// the struct [InvestigateBulkGetResponseActionParamsMove]
type investigateBulkGetResponseActionParamsMoveJSON struct {
	Destination         apijson.Field
	Type                apijson.Field
	ExpectedDisposition apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *InvestigateBulkGetResponseActionParamsMove) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkGetResponseActionParamsMoveJSON) RawJSON() string {
	return r.raw
}

func (r InvestigateBulkGetResponseActionParamsMove) implementsInvestigateBulkGetResponseActionParams() {
}

type InvestigateBulkGetResponseActionParamsMoveDestination string

const (
	InvestigateBulkGetResponseActionParamsMoveDestinationInbox                     InvestigateBulkGetResponseActionParamsMoveDestination = "Inbox"
	InvestigateBulkGetResponseActionParamsMoveDestinationJunkEmail                 InvestigateBulkGetResponseActionParamsMoveDestination = "JunkEmail"
	InvestigateBulkGetResponseActionParamsMoveDestinationDeletedItems              InvestigateBulkGetResponseActionParamsMoveDestination = "DeletedItems"
	InvestigateBulkGetResponseActionParamsMoveDestinationRecoverableItemsDeletions InvestigateBulkGetResponseActionParamsMoveDestination = "RecoverableItemsDeletions"
	InvestigateBulkGetResponseActionParamsMoveDestinationRecoverableItemsPurges    InvestigateBulkGetResponseActionParamsMoveDestination = "RecoverableItemsPurges"
)

func (r InvestigateBulkGetResponseActionParamsMoveDestination) IsKnown() bool {
	switch r {
	case InvestigateBulkGetResponseActionParamsMoveDestinationInbox, InvestigateBulkGetResponseActionParamsMoveDestinationJunkEmail, InvestigateBulkGetResponseActionParamsMoveDestinationDeletedItems, InvestigateBulkGetResponseActionParamsMoveDestinationRecoverableItemsDeletions, InvestigateBulkGetResponseActionParamsMoveDestinationRecoverableItemsPurges:
		return true
	}
	return false
}

type InvestigateBulkGetResponseActionParamsMoveType string

const (
	InvestigateBulkGetResponseActionParamsMoveTypeMove InvestigateBulkGetResponseActionParamsMoveType = "MOVE"
)

func (r InvestigateBulkGetResponseActionParamsMoveType) IsKnown() bool {
	switch r {
	case InvestigateBulkGetResponseActionParamsMoveTypeMove:
		return true
	}
	return false
}

type InvestigateBulkGetResponseActionParamsMoveExpectedDisposition string

const (
	InvestigateBulkGetResponseActionParamsMoveExpectedDispositionMalicious    InvestigateBulkGetResponseActionParamsMoveExpectedDisposition = "MALICIOUS"
	InvestigateBulkGetResponseActionParamsMoveExpectedDispositionMaliciousBec InvestigateBulkGetResponseActionParamsMoveExpectedDisposition = "MALICIOUS-BEC"
	InvestigateBulkGetResponseActionParamsMoveExpectedDispositionSuspicious   InvestigateBulkGetResponseActionParamsMoveExpectedDisposition = "SUSPICIOUS"
	InvestigateBulkGetResponseActionParamsMoveExpectedDispositionSpoof        InvestigateBulkGetResponseActionParamsMoveExpectedDisposition = "SPOOF"
	InvestigateBulkGetResponseActionParamsMoveExpectedDispositionSpam         InvestigateBulkGetResponseActionParamsMoveExpectedDisposition = "SPAM"
	InvestigateBulkGetResponseActionParamsMoveExpectedDispositionBulk         InvestigateBulkGetResponseActionParamsMoveExpectedDisposition = "BULK"
	InvestigateBulkGetResponseActionParamsMoveExpectedDispositionEncrypted    InvestigateBulkGetResponseActionParamsMoveExpectedDisposition = "ENCRYPTED"
	InvestigateBulkGetResponseActionParamsMoveExpectedDispositionExternal     InvestigateBulkGetResponseActionParamsMoveExpectedDisposition = "EXTERNAL"
	InvestigateBulkGetResponseActionParamsMoveExpectedDispositionUnknown      InvestigateBulkGetResponseActionParamsMoveExpectedDisposition = "UNKNOWN"
	InvestigateBulkGetResponseActionParamsMoveExpectedDispositionNone         InvestigateBulkGetResponseActionParamsMoveExpectedDisposition = "NONE"
)

func (r InvestigateBulkGetResponseActionParamsMoveExpectedDisposition) IsKnown() bool {
	switch r {
	case InvestigateBulkGetResponseActionParamsMoveExpectedDispositionMalicious, InvestigateBulkGetResponseActionParamsMoveExpectedDispositionMaliciousBec, InvestigateBulkGetResponseActionParamsMoveExpectedDispositionSuspicious, InvestigateBulkGetResponseActionParamsMoveExpectedDispositionSpoof, InvestigateBulkGetResponseActionParamsMoveExpectedDispositionSpam, InvestigateBulkGetResponseActionParamsMoveExpectedDispositionBulk, InvestigateBulkGetResponseActionParamsMoveExpectedDispositionEncrypted, InvestigateBulkGetResponseActionParamsMoveExpectedDispositionExternal, InvestigateBulkGetResponseActionParamsMoveExpectedDispositionUnknown, InvestigateBulkGetResponseActionParamsMoveExpectedDispositionNone:
		return true
	}
	return false
}

type InvestigateBulkGetResponseActionParamsRelease struct {
	Type InvestigateBulkGetResponseActionParamsReleaseType `json:"type" api:"required"`
	JSON investigateBulkGetResponseActionParamsReleaseJSON `json:"-"`
}

// investigateBulkGetResponseActionParamsReleaseJSON contains the JSON metadata for
// the struct [InvestigateBulkGetResponseActionParamsRelease]
type investigateBulkGetResponseActionParamsReleaseJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateBulkGetResponseActionParamsRelease) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkGetResponseActionParamsReleaseJSON) RawJSON() string {
	return r.raw
}

func (r InvestigateBulkGetResponseActionParamsRelease) implementsInvestigateBulkGetResponseActionParams() {
}

type InvestigateBulkGetResponseActionParamsReleaseType string

const (
	InvestigateBulkGetResponseActionParamsReleaseTypeRelease InvestigateBulkGetResponseActionParamsReleaseType = "RELEASE"
)

func (r InvestigateBulkGetResponseActionParamsReleaseType) IsKnown() bool {
	switch r {
	case InvestigateBulkGetResponseActionParamsReleaseTypeRelease:
		return true
	}
	return false
}

type InvestigateBulkGetResponseActionParamsType string

const (
	InvestigateBulkGetResponseActionParamsTypeMove    InvestigateBulkGetResponseActionParamsType = "MOVE"
	InvestigateBulkGetResponseActionParamsTypeRelease InvestigateBulkGetResponseActionParamsType = "RELEASE"
)

func (r InvestigateBulkGetResponseActionParamsType) IsKnown() bool {
	switch r {
	case InvestigateBulkGetResponseActionParamsTypeMove, InvestigateBulkGetResponseActionParamsTypeRelease:
		return true
	}
	return false
}

type InvestigateBulkGetResponseActionParamsDestination string

const (
	InvestigateBulkGetResponseActionParamsDestinationInbox                     InvestigateBulkGetResponseActionParamsDestination = "Inbox"
	InvestigateBulkGetResponseActionParamsDestinationJunkEmail                 InvestigateBulkGetResponseActionParamsDestination = "JunkEmail"
	InvestigateBulkGetResponseActionParamsDestinationDeletedItems              InvestigateBulkGetResponseActionParamsDestination = "DeletedItems"
	InvestigateBulkGetResponseActionParamsDestinationRecoverableItemsDeletions InvestigateBulkGetResponseActionParamsDestination = "RecoverableItemsDeletions"
	InvestigateBulkGetResponseActionParamsDestinationRecoverableItemsPurges    InvestigateBulkGetResponseActionParamsDestination = "RecoverableItemsPurges"
)

func (r InvestigateBulkGetResponseActionParamsDestination) IsKnown() bool {
	switch r {
	case InvestigateBulkGetResponseActionParamsDestinationInbox, InvestigateBulkGetResponseActionParamsDestinationJunkEmail, InvestigateBulkGetResponseActionParamsDestinationDeletedItems, InvestigateBulkGetResponseActionParamsDestinationRecoverableItemsDeletions, InvestigateBulkGetResponseActionParamsDestinationRecoverableItemsPurges:
		return true
	}
	return false
}

type InvestigateBulkGetResponseActionParamsExpectedDisposition string

const (
	InvestigateBulkGetResponseActionParamsExpectedDispositionMalicious    InvestigateBulkGetResponseActionParamsExpectedDisposition = "MALICIOUS"
	InvestigateBulkGetResponseActionParamsExpectedDispositionMaliciousBec InvestigateBulkGetResponseActionParamsExpectedDisposition = "MALICIOUS-BEC"
	InvestigateBulkGetResponseActionParamsExpectedDispositionSuspicious   InvestigateBulkGetResponseActionParamsExpectedDisposition = "SUSPICIOUS"
	InvestigateBulkGetResponseActionParamsExpectedDispositionSpoof        InvestigateBulkGetResponseActionParamsExpectedDisposition = "SPOOF"
	InvestigateBulkGetResponseActionParamsExpectedDispositionSpam         InvestigateBulkGetResponseActionParamsExpectedDisposition = "SPAM"
	InvestigateBulkGetResponseActionParamsExpectedDispositionBulk         InvestigateBulkGetResponseActionParamsExpectedDisposition = "BULK"
	InvestigateBulkGetResponseActionParamsExpectedDispositionEncrypted    InvestigateBulkGetResponseActionParamsExpectedDisposition = "ENCRYPTED"
	InvestigateBulkGetResponseActionParamsExpectedDispositionExternal     InvestigateBulkGetResponseActionParamsExpectedDisposition = "EXTERNAL"
	InvestigateBulkGetResponseActionParamsExpectedDispositionUnknown      InvestigateBulkGetResponseActionParamsExpectedDisposition = "UNKNOWN"
	InvestigateBulkGetResponseActionParamsExpectedDispositionNone         InvestigateBulkGetResponseActionParamsExpectedDisposition = "NONE"
)

func (r InvestigateBulkGetResponseActionParamsExpectedDisposition) IsKnown() bool {
	switch r {
	case InvestigateBulkGetResponseActionParamsExpectedDispositionMalicious, InvestigateBulkGetResponseActionParamsExpectedDispositionMaliciousBec, InvestigateBulkGetResponseActionParamsExpectedDispositionSuspicious, InvestigateBulkGetResponseActionParamsExpectedDispositionSpoof, InvestigateBulkGetResponseActionParamsExpectedDispositionSpam, InvestigateBulkGetResponseActionParamsExpectedDispositionBulk, InvestigateBulkGetResponseActionParamsExpectedDispositionEncrypted, InvestigateBulkGetResponseActionParamsExpectedDispositionExternal, InvestigateBulkGetResponseActionParamsExpectedDispositionUnknown, InvestigateBulkGetResponseActionParamsExpectedDispositionNone:
		return true
	}
	return false
}

type InvestigateBulkGetResponseActionType string

const (
	InvestigateBulkGetResponseActionTypeMove    InvestigateBulkGetResponseActionType = "MOVE"
	InvestigateBulkGetResponseActionTypeRelease InvestigateBulkGetResponseActionType = "RELEASE"
)

func (r InvestigateBulkGetResponseActionType) IsKnown() bool {
	switch r {
	case InvestigateBulkGetResponseActionTypeMove, InvestigateBulkGetResponseActionTypeRelease:
		return true
	}
	return false
}

type InvestigateBulkGetResponseSearchParams struct {
	// Deprecated, use `GET /investigate/{investigate_id}/action_log` instead. End of
	// life: November 1, 2026.
	//
	// Deprecated: deprecated
	ActionLog bool   `json:"action_log"`
	AlertID   string `json:"alert_id" api:"nullable"`
	// Delivery status of the message.
	DeliveryStatus InvestigateBulkGetResponseSearchParamsDeliveryStatus `json:"delivery_status"`
	DetectionsOnly bool                                                 `json:"detections_only"`
	Domain         string                                               `json:"domain" api:"nullable"`
	// End of search date range
	End              time.Time                                              `json:"end" format:"date-time"`
	ExactSubject     string                                                 `json:"exact_subject" api:"nullable"`
	FinalDisposition InvestigateBulkGetResponseSearchParamsFinalDisposition `json:"final_disposition"`
	MessageAction    InvestigateBulkGetResponseSearchParamsMessageAction    `json:"message_action" api:"nullable"`
	MessageID        string                                                 `json:"message_id" api:"nullable"`
	Metric           string                                                 `json:"metric" api:"nullable"`
	Query            string                                                 `json:"query" api:"nullable"`
	Recipient        string                                                 `json:"recipient" api:"nullable"`
	Sender           string                                                 `json:"sender" api:"nullable"`
	// Beginning of search date range
	Start       time.Time                                  `json:"start" format:"date-time"`
	Subject     string                                     `json:"subject" api:"nullable"`
	Submissions bool                                       `json:"submissions"`
	JSON        investigateBulkGetResponseSearchParamsJSON `json:"-"`
}

// investigateBulkGetResponseSearchParamsJSON contains the JSON metadata for the
// struct [InvestigateBulkGetResponseSearchParams]
type investigateBulkGetResponseSearchParamsJSON struct {
	ActionLog        apijson.Field
	AlertID          apijson.Field
	DeliveryStatus   apijson.Field
	DetectionsOnly   apijson.Field
	Domain           apijson.Field
	End              apijson.Field
	ExactSubject     apijson.Field
	FinalDisposition apijson.Field
	MessageAction    apijson.Field
	MessageID        apijson.Field
	Metric           apijson.Field
	Query            apijson.Field
	Recipient        apijson.Field
	Sender           apijson.Field
	Start            apijson.Field
	Subject          apijson.Field
	Submissions      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvestigateBulkGetResponseSearchParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkGetResponseSearchParamsJSON) RawJSON() string {
	return r.raw
}

// Delivery status of the message.
type InvestigateBulkGetResponseSearchParamsDeliveryStatus string

const (
	InvestigateBulkGetResponseSearchParamsDeliveryStatusDelivered   InvestigateBulkGetResponseSearchParamsDeliveryStatus = "delivered"
	InvestigateBulkGetResponseSearchParamsDeliveryStatusMoved       InvestigateBulkGetResponseSearchParamsDeliveryStatus = "moved"
	InvestigateBulkGetResponseSearchParamsDeliveryStatusQuarantined InvestigateBulkGetResponseSearchParamsDeliveryStatus = "quarantined"
	InvestigateBulkGetResponseSearchParamsDeliveryStatusRejected    InvestigateBulkGetResponseSearchParamsDeliveryStatus = "rejected"
	InvestigateBulkGetResponseSearchParamsDeliveryStatusDeferred    InvestigateBulkGetResponseSearchParamsDeliveryStatus = "deferred"
	InvestigateBulkGetResponseSearchParamsDeliveryStatusBounced     InvestigateBulkGetResponseSearchParamsDeliveryStatus = "bounced"
	InvestigateBulkGetResponseSearchParamsDeliveryStatusQueued      InvestigateBulkGetResponseSearchParamsDeliveryStatus = "queued"
)

func (r InvestigateBulkGetResponseSearchParamsDeliveryStatus) IsKnown() bool {
	switch r {
	case InvestigateBulkGetResponseSearchParamsDeliveryStatusDelivered, InvestigateBulkGetResponseSearchParamsDeliveryStatusMoved, InvestigateBulkGetResponseSearchParamsDeliveryStatusQuarantined, InvestigateBulkGetResponseSearchParamsDeliveryStatusRejected, InvestigateBulkGetResponseSearchParamsDeliveryStatusDeferred, InvestigateBulkGetResponseSearchParamsDeliveryStatusBounced, InvestigateBulkGetResponseSearchParamsDeliveryStatusQueued:
		return true
	}
	return false
}

type InvestigateBulkGetResponseSearchParamsFinalDisposition string

const (
	InvestigateBulkGetResponseSearchParamsFinalDispositionMalicious    InvestigateBulkGetResponseSearchParamsFinalDisposition = "MALICIOUS"
	InvestigateBulkGetResponseSearchParamsFinalDispositionMaliciousBec InvestigateBulkGetResponseSearchParamsFinalDisposition = "MALICIOUS-BEC"
	InvestigateBulkGetResponseSearchParamsFinalDispositionSuspicious   InvestigateBulkGetResponseSearchParamsFinalDisposition = "SUSPICIOUS"
	InvestigateBulkGetResponseSearchParamsFinalDispositionSpoof        InvestigateBulkGetResponseSearchParamsFinalDisposition = "SPOOF"
	InvestigateBulkGetResponseSearchParamsFinalDispositionSpam         InvestigateBulkGetResponseSearchParamsFinalDisposition = "SPAM"
	InvestigateBulkGetResponseSearchParamsFinalDispositionBulk         InvestigateBulkGetResponseSearchParamsFinalDisposition = "BULK"
	InvestigateBulkGetResponseSearchParamsFinalDispositionEncrypted    InvestigateBulkGetResponseSearchParamsFinalDisposition = "ENCRYPTED"
	InvestigateBulkGetResponseSearchParamsFinalDispositionExternal     InvestigateBulkGetResponseSearchParamsFinalDisposition = "EXTERNAL"
	InvestigateBulkGetResponseSearchParamsFinalDispositionUnknown      InvestigateBulkGetResponseSearchParamsFinalDisposition = "UNKNOWN"
	InvestigateBulkGetResponseSearchParamsFinalDispositionNone         InvestigateBulkGetResponseSearchParamsFinalDisposition = "NONE"
)

func (r InvestigateBulkGetResponseSearchParamsFinalDisposition) IsKnown() bool {
	switch r {
	case InvestigateBulkGetResponseSearchParamsFinalDispositionMalicious, InvestigateBulkGetResponseSearchParamsFinalDispositionMaliciousBec, InvestigateBulkGetResponseSearchParamsFinalDispositionSuspicious, InvestigateBulkGetResponseSearchParamsFinalDispositionSpoof, InvestigateBulkGetResponseSearchParamsFinalDispositionSpam, InvestigateBulkGetResponseSearchParamsFinalDispositionBulk, InvestigateBulkGetResponseSearchParamsFinalDispositionEncrypted, InvestigateBulkGetResponseSearchParamsFinalDispositionExternal, InvestigateBulkGetResponseSearchParamsFinalDispositionUnknown, InvestigateBulkGetResponseSearchParamsFinalDispositionNone:
		return true
	}
	return false
}

type InvestigateBulkGetResponseSearchParamsMessageAction string

const (
	InvestigateBulkGetResponseSearchParamsMessageActionPreview            InvestigateBulkGetResponseSearchParamsMessageAction = "PREVIEW"
	InvestigateBulkGetResponseSearchParamsMessageActionQuarantineReleased InvestigateBulkGetResponseSearchParamsMessageAction = "QUARANTINE_RELEASED"
	InvestigateBulkGetResponseSearchParamsMessageActionMoved              InvestigateBulkGetResponseSearchParamsMessageAction = "MOVED"
)

func (r InvestigateBulkGetResponseSearchParamsMessageAction) IsKnown() bool {
	switch r {
	case InvestigateBulkGetResponseSearchParamsMessageActionPreview, InvestigateBulkGetResponseSearchParamsMessageActionQuarantineReleased, InvestigateBulkGetResponseSearchParamsMessageActionMoved:
		return true
	}
	return false
}

type InvestigateBulkGetResponseStatus string

const (
	InvestigateBulkGetResponseStatusPending     InvestigateBulkGetResponseStatus = "PENDING"
	InvestigateBulkGetResponseStatusDiscovering InvestigateBulkGetResponseStatus = "DISCOVERING"
	InvestigateBulkGetResponseStatusProcessing  InvestigateBulkGetResponseStatus = "PROCESSING"
	InvestigateBulkGetResponseStatusCompleted   InvestigateBulkGetResponseStatus = "COMPLETED"
	InvestigateBulkGetResponseStatusFailed      InvestigateBulkGetResponseStatus = "FAILED"
	InvestigateBulkGetResponseStatusCancelled   InvestigateBulkGetResponseStatus = "CANCELLED"
	InvestigateBulkGetResponseStatusSkipped     InvestigateBulkGetResponseStatus = "SKIPPED"
)

func (r InvestigateBulkGetResponseStatus) IsKnown() bool {
	switch r {
	case InvestigateBulkGetResponseStatusPending, InvestigateBulkGetResponseStatusDiscovering, InvestigateBulkGetResponseStatusProcessing, InvestigateBulkGetResponseStatusCompleted, InvestigateBulkGetResponseStatusFailed, InvestigateBulkGetResponseStatusCancelled, InvestigateBulkGetResponseStatusSkipped:
		return true
	}
	return false
}

type InvestigateBulkNewParams struct {
	// Identifier.
	AccountID           param.Field[string]                                      `path:"account_id" api:"required"`
	Action              param.Field[InvestigateBulkNewParamsAction]              `json:"action" api:"required"`
	SearchParams        param.Field[InvestigateBulkNewParamsSearchParams]        `json:"search_params" api:"required"`
	Comment             param.Field[string]                                      `json:"comment"`
	Destination         param.Field[InvestigateBulkNewParamsDestination]         `json:"destination"`
	ExpectedDisposition param.Field[InvestigateBulkNewParamsExpectedDisposition] `json:"expected_disposition"`
}

func (r InvestigateBulkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvestigateBulkNewParamsAction string

const (
	InvestigateBulkNewParamsActionMove    InvestigateBulkNewParamsAction = "MOVE"
	InvestigateBulkNewParamsActionRelease InvestigateBulkNewParamsAction = "RELEASE"
)

func (r InvestigateBulkNewParamsAction) IsKnown() bool {
	switch r {
	case InvestigateBulkNewParamsActionMove, InvestigateBulkNewParamsActionRelease:
		return true
	}
	return false
}

type InvestigateBulkNewParamsSearchParams struct {
	// Deprecated, use `GET /investigate/{investigate_id}/action_log` instead. End of
	// life: November 1, 2026.
	//
	// Deprecated: deprecated
	ActionLog param.Field[bool]   `json:"action_log"`
	AlertID   param.Field[string] `json:"alert_id"`
	// Delivery status of the message.
	DeliveryStatus param.Field[InvestigateBulkNewParamsSearchParamsDeliveryStatus] `json:"delivery_status"`
	DetectionsOnly param.Field[bool]                                               `json:"detections_only"`
	Domain         param.Field[string]                                             `json:"domain"`
	// End of search date range
	End              param.Field[time.Time]                                            `json:"end" format:"date-time"`
	ExactSubject     param.Field[string]                                               `json:"exact_subject"`
	FinalDisposition param.Field[InvestigateBulkNewParamsSearchParamsFinalDisposition] `json:"final_disposition"`
	MessageAction    param.Field[InvestigateBulkNewParamsSearchParamsMessageAction]    `json:"message_action"`
	MessageID        param.Field[string]                                               `json:"message_id"`
	Metric           param.Field[string]                                               `json:"metric"`
	Query            param.Field[string]                                               `json:"query"`
	Recipient        param.Field[string]                                               `json:"recipient"`
	Sender           param.Field[string]                                               `json:"sender"`
	// Beginning of search date range
	Start       param.Field[time.Time] `json:"start" format:"date-time"`
	Subject     param.Field[string]    `json:"subject"`
	Submissions param.Field[bool]      `json:"submissions"`
}

func (r InvestigateBulkNewParamsSearchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Delivery status of the message.
type InvestigateBulkNewParamsSearchParamsDeliveryStatus string

const (
	InvestigateBulkNewParamsSearchParamsDeliveryStatusDelivered   InvestigateBulkNewParamsSearchParamsDeliveryStatus = "delivered"
	InvestigateBulkNewParamsSearchParamsDeliveryStatusMoved       InvestigateBulkNewParamsSearchParamsDeliveryStatus = "moved"
	InvestigateBulkNewParamsSearchParamsDeliveryStatusQuarantined InvestigateBulkNewParamsSearchParamsDeliveryStatus = "quarantined"
	InvestigateBulkNewParamsSearchParamsDeliveryStatusRejected    InvestigateBulkNewParamsSearchParamsDeliveryStatus = "rejected"
	InvestigateBulkNewParamsSearchParamsDeliveryStatusDeferred    InvestigateBulkNewParamsSearchParamsDeliveryStatus = "deferred"
	InvestigateBulkNewParamsSearchParamsDeliveryStatusBounced     InvestigateBulkNewParamsSearchParamsDeliveryStatus = "bounced"
	InvestigateBulkNewParamsSearchParamsDeliveryStatusQueued      InvestigateBulkNewParamsSearchParamsDeliveryStatus = "queued"
)

func (r InvestigateBulkNewParamsSearchParamsDeliveryStatus) IsKnown() bool {
	switch r {
	case InvestigateBulkNewParamsSearchParamsDeliveryStatusDelivered, InvestigateBulkNewParamsSearchParamsDeliveryStatusMoved, InvestigateBulkNewParamsSearchParamsDeliveryStatusQuarantined, InvestigateBulkNewParamsSearchParamsDeliveryStatusRejected, InvestigateBulkNewParamsSearchParamsDeliveryStatusDeferred, InvestigateBulkNewParamsSearchParamsDeliveryStatusBounced, InvestigateBulkNewParamsSearchParamsDeliveryStatusQueued:
		return true
	}
	return false
}

type InvestigateBulkNewParamsSearchParamsFinalDisposition string

const (
	InvestigateBulkNewParamsSearchParamsFinalDispositionMalicious    InvestigateBulkNewParamsSearchParamsFinalDisposition = "MALICIOUS"
	InvestigateBulkNewParamsSearchParamsFinalDispositionMaliciousBec InvestigateBulkNewParamsSearchParamsFinalDisposition = "MALICIOUS-BEC"
	InvestigateBulkNewParamsSearchParamsFinalDispositionSuspicious   InvestigateBulkNewParamsSearchParamsFinalDisposition = "SUSPICIOUS"
	InvestigateBulkNewParamsSearchParamsFinalDispositionSpoof        InvestigateBulkNewParamsSearchParamsFinalDisposition = "SPOOF"
	InvestigateBulkNewParamsSearchParamsFinalDispositionSpam         InvestigateBulkNewParamsSearchParamsFinalDisposition = "SPAM"
	InvestigateBulkNewParamsSearchParamsFinalDispositionBulk         InvestigateBulkNewParamsSearchParamsFinalDisposition = "BULK"
	InvestigateBulkNewParamsSearchParamsFinalDispositionEncrypted    InvestigateBulkNewParamsSearchParamsFinalDisposition = "ENCRYPTED"
	InvestigateBulkNewParamsSearchParamsFinalDispositionExternal     InvestigateBulkNewParamsSearchParamsFinalDisposition = "EXTERNAL"
	InvestigateBulkNewParamsSearchParamsFinalDispositionUnknown      InvestigateBulkNewParamsSearchParamsFinalDisposition = "UNKNOWN"
	InvestigateBulkNewParamsSearchParamsFinalDispositionNone         InvestigateBulkNewParamsSearchParamsFinalDisposition = "NONE"
)

func (r InvestigateBulkNewParamsSearchParamsFinalDisposition) IsKnown() bool {
	switch r {
	case InvestigateBulkNewParamsSearchParamsFinalDispositionMalicious, InvestigateBulkNewParamsSearchParamsFinalDispositionMaliciousBec, InvestigateBulkNewParamsSearchParamsFinalDispositionSuspicious, InvestigateBulkNewParamsSearchParamsFinalDispositionSpoof, InvestigateBulkNewParamsSearchParamsFinalDispositionSpam, InvestigateBulkNewParamsSearchParamsFinalDispositionBulk, InvestigateBulkNewParamsSearchParamsFinalDispositionEncrypted, InvestigateBulkNewParamsSearchParamsFinalDispositionExternal, InvestigateBulkNewParamsSearchParamsFinalDispositionUnknown, InvestigateBulkNewParamsSearchParamsFinalDispositionNone:
		return true
	}
	return false
}

type InvestigateBulkNewParamsSearchParamsMessageAction string

const (
	InvestigateBulkNewParamsSearchParamsMessageActionPreview            InvestigateBulkNewParamsSearchParamsMessageAction = "PREVIEW"
	InvestigateBulkNewParamsSearchParamsMessageActionQuarantineReleased InvestigateBulkNewParamsSearchParamsMessageAction = "QUARANTINE_RELEASED"
	InvestigateBulkNewParamsSearchParamsMessageActionMoved              InvestigateBulkNewParamsSearchParamsMessageAction = "MOVED"
)

func (r InvestigateBulkNewParamsSearchParamsMessageAction) IsKnown() bool {
	switch r {
	case InvestigateBulkNewParamsSearchParamsMessageActionPreview, InvestigateBulkNewParamsSearchParamsMessageActionQuarantineReleased, InvestigateBulkNewParamsSearchParamsMessageActionMoved:
		return true
	}
	return false
}

type InvestigateBulkNewParamsDestination string

const (
	InvestigateBulkNewParamsDestinationInbox                     InvestigateBulkNewParamsDestination = "Inbox"
	InvestigateBulkNewParamsDestinationJunkEmail                 InvestigateBulkNewParamsDestination = "JunkEmail"
	InvestigateBulkNewParamsDestinationDeletedItems              InvestigateBulkNewParamsDestination = "DeletedItems"
	InvestigateBulkNewParamsDestinationRecoverableItemsDeletions InvestigateBulkNewParamsDestination = "RecoverableItemsDeletions"
	InvestigateBulkNewParamsDestinationRecoverableItemsPurges    InvestigateBulkNewParamsDestination = "RecoverableItemsPurges"
)

func (r InvestigateBulkNewParamsDestination) IsKnown() bool {
	switch r {
	case InvestigateBulkNewParamsDestinationInbox, InvestigateBulkNewParamsDestinationJunkEmail, InvestigateBulkNewParamsDestinationDeletedItems, InvestigateBulkNewParamsDestinationRecoverableItemsDeletions, InvestigateBulkNewParamsDestinationRecoverableItemsPurges:
		return true
	}
	return false
}

type InvestigateBulkNewParamsExpectedDisposition string

const (
	InvestigateBulkNewParamsExpectedDispositionMalicious    InvestigateBulkNewParamsExpectedDisposition = "MALICIOUS"
	InvestigateBulkNewParamsExpectedDispositionMaliciousBec InvestigateBulkNewParamsExpectedDisposition = "MALICIOUS-BEC"
	InvestigateBulkNewParamsExpectedDispositionSuspicious   InvestigateBulkNewParamsExpectedDisposition = "SUSPICIOUS"
	InvestigateBulkNewParamsExpectedDispositionSpoof        InvestigateBulkNewParamsExpectedDisposition = "SPOOF"
	InvestigateBulkNewParamsExpectedDispositionSpam         InvestigateBulkNewParamsExpectedDisposition = "SPAM"
	InvestigateBulkNewParamsExpectedDispositionBulk         InvestigateBulkNewParamsExpectedDisposition = "BULK"
	InvestigateBulkNewParamsExpectedDispositionEncrypted    InvestigateBulkNewParamsExpectedDisposition = "ENCRYPTED"
	InvestigateBulkNewParamsExpectedDispositionExternal     InvestigateBulkNewParamsExpectedDisposition = "EXTERNAL"
	InvestigateBulkNewParamsExpectedDispositionUnknown      InvestigateBulkNewParamsExpectedDisposition = "UNKNOWN"
	InvestigateBulkNewParamsExpectedDispositionNone         InvestigateBulkNewParamsExpectedDisposition = "NONE"
)

func (r InvestigateBulkNewParamsExpectedDisposition) IsKnown() bool {
	switch r {
	case InvestigateBulkNewParamsExpectedDispositionMalicious, InvestigateBulkNewParamsExpectedDispositionMaliciousBec, InvestigateBulkNewParamsExpectedDispositionSuspicious, InvestigateBulkNewParamsExpectedDispositionSpoof, InvestigateBulkNewParamsExpectedDispositionSpam, InvestigateBulkNewParamsExpectedDispositionBulk, InvestigateBulkNewParamsExpectedDispositionEncrypted, InvestigateBulkNewParamsExpectedDispositionExternal, InvestigateBulkNewParamsExpectedDispositionUnknown, InvestigateBulkNewParamsExpectedDispositionNone:
		return true
	}
	return false
}

type InvestigateBulkNewResponseEnvelope struct {
	Errors   []InvestigateBulkNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []InvestigateBulkNewResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   InvestigateBulkNewResponse                   `json:"result" api:"required"`
	// Whether the API call was successful.
	Success InvestigateBulkNewResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    investigateBulkNewResponseEnvelopeJSON    `json:"-"`
}

// investigateBulkNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [InvestigateBulkNewResponseEnvelope]
type investigateBulkNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateBulkNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InvestigateBulkNewResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code" api:"required"`
	Message          string                                         `json:"message" api:"required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           InvestigateBulkNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             investigateBulkNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// investigateBulkNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [InvestigateBulkNewResponseEnvelopeErrors]
type investigateBulkNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvestigateBulkNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type InvestigateBulkNewResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    investigateBulkNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// investigateBulkNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [InvestigateBulkNewResponseEnvelopeErrorsSource]
type investigateBulkNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateBulkNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type InvestigateBulkNewResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code" api:"required"`
	Message          string                                           `json:"message" api:"required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           InvestigateBulkNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             investigateBulkNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// investigateBulkNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [InvestigateBulkNewResponseEnvelopeMessages]
type investigateBulkNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvestigateBulkNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type InvestigateBulkNewResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    investigateBulkNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// investigateBulkNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [InvestigateBulkNewResponseEnvelopeMessagesSource]
type investigateBulkNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateBulkNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type InvestigateBulkNewResponseEnvelopeSuccess bool

const (
	InvestigateBulkNewResponseEnvelopeSuccessTrue InvestigateBulkNewResponseEnvelopeSuccess = true
)

func (r InvestigateBulkNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case InvestigateBulkNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type InvestigateBulkListParams struct {
	// Identifier.
	AccountID  param.Field[string]                              `path:"account_id" api:"required"`
	ActionType param.Field[InvestigateBulkListParamsActionType] `query:"action_type"`
	// Current page within paginated list of results.
	Page param.Field[int64] `query:"page"`
	// The number of results per page. Maximum value is 1000.
	PerPage param.Field[int64]                           `query:"per_page"`
	Status  param.Field[InvestigateBulkListParamsStatus] `query:"status"`
}

// URLQuery serializes [InvestigateBulkListParams]'s query parameters as
// `url.Values`.
func (r InvestigateBulkListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InvestigateBulkListParamsActionType string

const (
	InvestigateBulkListParamsActionTypeMove    InvestigateBulkListParamsActionType = "MOVE"
	InvestigateBulkListParamsActionTypeRelease InvestigateBulkListParamsActionType = "RELEASE"
)

func (r InvestigateBulkListParamsActionType) IsKnown() bool {
	switch r {
	case InvestigateBulkListParamsActionTypeMove, InvestigateBulkListParamsActionTypeRelease:
		return true
	}
	return false
}

type InvestigateBulkListParamsStatus string

const (
	InvestigateBulkListParamsStatusPending     InvestigateBulkListParamsStatus = "PENDING"
	InvestigateBulkListParamsStatusDiscovering InvestigateBulkListParamsStatus = "DISCOVERING"
	InvestigateBulkListParamsStatusProcessing  InvestigateBulkListParamsStatus = "PROCESSING"
	InvestigateBulkListParamsStatusCompleted   InvestigateBulkListParamsStatus = "COMPLETED"
	InvestigateBulkListParamsStatusFailed      InvestigateBulkListParamsStatus = "FAILED"
	InvestigateBulkListParamsStatusCancelled   InvestigateBulkListParamsStatus = "CANCELLED"
	InvestigateBulkListParamsStatusSkipped     InvestigateBulkListParamsStatus = "SKIPPED"
)

func (r InvestigateBulkListParamsStatus) IsKnown() bool {
	switch r {
	case InvestigateBulkListParamsStatusPending, InvestigateBulkListParamsStatusDiscovering, InvestigateBulkListParamsStatusProcessing, InvestigateBulkListParamsStatusCompleted, InvestigateBulkListParamsStatusFailed, InvestigateBulkListParamsStatusCancelled, InvestigateBulkListParamsStatusSkipped:
		return true
	}
	return false
}

type InvestigateBulkDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type InvestigateBulkDeleteResponseEnvelope struct {
	Errors   []InvestigateBulkDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []InvestigateBulkDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   InvestigateBulkDeleteResponse                   `json:"result" api:"required"`
	// Whether the API call was successful.
	Success InvestigateBulkDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    investigateBulkDeleteResponseEnvelopeJSON    `json:"-"`
}

// investigateBulkDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [InvestigateBulkDeleteResponseEnvelope]
type investigateBulkDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateBulkDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InvestigateBulkDeleteResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code" api:"required"`
	Message          string                                            `json:"message" api:"required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           InvestigateBulkDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             investigateBulkDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// investigateBulkDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [InvestigateBulkDeleteResponseEnvelopeErrors]
type investigateBulkDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvestigateBulkDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type InvestigateBulkDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    investigateBulkDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// investigateBulkDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [InvestigateBulkDeleteResponseEnvelopeErrorsSource]
type investigateBulkDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateBulkDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type InvestigateBulkDeleteResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code" api:"required"`
	Message          string                                              `json:"message" api:"required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           InvestigateBulkDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             investigateBulkDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// investigateBulkDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [InvestigateBulkDeleteResponseEnvelopeMessages]
type investigateBulkDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvestigateBulkDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type InvestigateBulkDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    investigateBulkDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// investigateBulkDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [InvestigateBulkDeleteResponseEnvelopeMessagesSource]
type investigateBulkDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateBulkDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type InvestigateBulkDeleteResponseEnvelopeSuccess bool

const (
	InvestigateBulkDeleteResponseEnvelopeSuccessTrue InvestigateBulkDeleteResponseEnvelopeSuccess = true
)

func (r InvestigateBulkDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case InvestigateBulkDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type InvestigateBulkGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type InvestigateBulkGetResponseEnvelope struct {
	Errors   []InvestigateBulkGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []InvestigateBulkGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   InvestigateBulkGetResponse                   `json:"result" api:"required"`
	// Whether the API call was successful.
	Success InvestigateBulkGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    investigateBulkGetResponseEnvelopeJSON    `json:"-"`
}

// investigateBulkGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [InvestigateBulkGetResponseEnvelope]
type investigateBulkGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateBulkGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InvestigateBulkGetResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code" api:"required"`
	Message          string                                         `json:"message" api:"required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           InvestigateBulkGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             investigateBulkGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// investigateBulkGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [InvestigateBulkGetResponseEnvelopeErrors]
type investigateBulkGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvestigateBulkGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type InvestigateBulkGetResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    investigateBulkGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// investigateBulkGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [InvestigateBulkGetResponseEnvelopeErrorsSource]
type investigateBulkGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateBulkGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type InvestigateBulkGetResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code" api:"required"`
	Message          string                                           `json:"message" api:"required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           InvestigateBulkGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             investigateBulkGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// investigateBulkGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [InvestigateBulkGetResponseEnvelopeMessages]
type investigateBulkGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvestigateBulkGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type InvestigateBulkGetResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    investigateBulkGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// investigateBulkGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [InvestigateBulkGetResponseEnvelopeMessagesSource]
type investigateBulkGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateBulkGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type InvestigateBulkGetResponseEnvelopeSuccess bool

const (
	InvestigateBulkGetResponseEnvelopeSuccessTrue InvestigateBulkGetResponseEnvelopeSuccess = true
)

func (r InvestigateBulkGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case InvestigateBulkGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
