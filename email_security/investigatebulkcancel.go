// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/tidwall/gjson"
)

// InvestigateBulkCancelService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvestigateBulkCancelService] method instead.
type InvestigateBulkCancelService struct {
	Options []option.RequestOption
}

// NewInvestigateBulkCancelService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInvestigateBulkCancelService(opts ...option.RequestOption) (r *InvestigateBulkCancelService) {
	r = &InvestigateBulkCancelService{}
	r.Options = opts
	return
}

// Marks the job as cancelled and stops any pending message processing. The job
// record remains visible in list and detail endpoints.
func (r *InvestigateBulkCancelService) New(ctx context.Context, jobID string, body InvestigateBulkCancelNewParams, opts ...option.RequestOption) (res *InvestigateBulkCancelNewResponse, err error) {
	var env InvestigateBulkCancelNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/bulk/%s/cancel", body.AccountID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type InvestigateBulkCancelNewResponse struct {
	ActionParams            InvestigateBulkCancelNewResponseActionParams `json:"action_params" api:"required"`
	ActionType              InvestigateBulkCancelNewResponseActionType   `json:"action_type" api:"required"`
	CreatedAt               time.Time                                    `json:"created_at" api:"required" format:"date-time"`
	JobID                   string                                       `json:"job_id" api:"required" format:"uuid"`
	MessagesFailed          int64                                        `json:"messages_failed" api:"required"`
	MessagesPending         int64                                        `json:"messages_pending" api:"required"`
	MessagesSuccessful      int64                                        `json:"messages_successful" api:"required"`
	SearchParams            InvestigateBulkCancelNewResponseSearchParams `json:"search_params" api:"required"`
	Status                  InvestigateBulkCancelNewResponseStatus       `json:"status" api:"required"`
	TotalMessagesDiscovered int64                                        `json:"total_messages_discovered" api:"required"`
	Comment                 string                                       `json:"comment" api:"nullable"`
	CompletedAt             time.Time                                    `json:"completed_at" api:"nullable" format:"date-time"`
	StartedAt               time.Time                                    `json:"started_at" api:"nullable" format:"date-time"`
	StatusMessage           string                                       `json:"status_message" api:"nullable"`
	JSON                    investigateBulkCancelNewResponseJSON         `json:"-"`
}

// investigateBulkCancelNewResponseJSON contains the JSON metadata for the struct
// [InvestigateBulkCancelNewResponse]
type investigateBulkCancelNewResponseJSON struct {
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

func (r *InvestigateBulkCancelNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkCancelNewResponseJSON) RawJSON() string {
	return r.raw
}

type InvestigateBulkCancelNewResponseActionParams struct {
	Type                InvestigateBulkCancelNewResponseActionParamsType                `json:"type" api:"required"`
	Destination         InvestigateBulkCancelNewResponseActionParamsDestination         `json:"destination"`
	ExpectedDisposition InvestigateBulkCancelNewResponseActionParamsExpectedDisposition `json:"expected_disposition"`
	JSON                investigateBulkCancelNewResponseActionParamsJSON                `json:"-"`
	union               InvestigateBulkCancelNewResponseActionParamsUnion
}

// investigateBulkCancelNewResponseActionParamsJSON contains the JSON metadata for
// the struct [InvestigateBulkCancelNewResponseActionParams]
type investigateBulkCancelNewResponseActionParamsJSON struct {
	Type                apijson.Field
	Destination         apijson.Field
	ExpectedDisposition apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r investigateBulkCancelNewResponseActionParamsJSON) RawJSON() string {
	return r.raw
}

func (r *InvestigateBulkCancelNewResponseActionParams) UnmarshalJSON(data []byte) (err error) {
	*r = InvestigateBulkCancelNewResponseActionParams{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [InvestigateBulkCancelNewResponseActionParamsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [InvestigateBulkCancelNewResponseActionParamsMove],
// [InvestigateBulkCancelNewResponseActionParamsRelease].
func (r InvestigateBulkCancelNewResponseActionParams) AsUnion() InvestigateBulkCancelNewResponseActionParamsUnion {
	return r.union
}

// Union satisfied by [InvestigateBulkCancelNewResponseActionParamsMove] or
// [InvestigateBulkCancelNewResponseActionParamsRelease].
type InvestigateBulkCancelNewResponseActionParamsUnion interface {
	implementsInvestigateBulkCancelNewResponseActionParams()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InvestigateBulkCancelNewResponseActionParamsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvestigateBulkCancelNewResponseActionParamsMove{}),
			DiscriminatorValue: "MOVE",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvestigateBulkCancelNewResponseActionParamsRelease{}),
			DiscriminatorValue: "RELEASE",
		},
	)
}

type InvestigateBulkCancelNewResponseActionParamsMove struct {
	Destination         InvestigateBulkCancelNewResponseActionParamsMoveDestination         `json:"destination" api:"required"`
	Type                InvestigateBulkCancelNewResponseActionParamsMoveType                `json:"type" api:"required"`
	ExpectedDisposition InvestigateBulkCancelNewResponseActionParamsMoveExpectedDisposition `json:"expected_disposition"`
	JSON                investigateBulkCancelNewResponseActionParamsMoveJSON                `json:"-"`
}

// investigateBulkCancelNewResponseActionParamsMoveJSON contains the JSON metadata
// for the struct [InvestigateBulkCancelNewResponseActionParamsMove]
type investigateBulkCancelNewResponseActionParamsMoveJSON struct {
	Destination         apijson.Field
	Type                apijson.Field
	ExpectedDisposition apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *InvestigateBulkCancelNewResponseActionParamsMove) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkCancelNewResponseActionParamsMoveJSON) RawJSON() string {
	return r.raw
}

func (r InvestigateBulkCancelNewResponseActionParamsMove) implementsInvestigateBulkCancelNewResponseActionParams() {
}

type InvestigateBulkCancelNewResponseActionParamsMoveDestination string

const (
	InvestigateBulkCancelNewResponseActionParamsMoveDestinationInbox                     InvestigateBulkCancelNewResponseActionParamsMoveDestination = "Inbox"
	InvestigateBulkCancelNewResponseActionParamsMoveDestinationJunkEmail                 InvestigateBulkCancelNewResponseActionParamsMoveDestination = "JunkEmail"
	InvestigateBulkCancelNewResponseActionParamsMoveDestinationDeletedItems              InvestigateBulkCancelNewResponseActionParamsMoveDestination = "DeletedItems"
	InvestigateBulkCancelNewResponseActionParamsMoveDestinationRecoverableItemsDeletions InvestigateBulkCancelNewResponseActionParamsMoveDestination = "RecoverableItemsDeletions"
	InvestigateBulkCancelNewResponseActionParamsMoveDestinationRecoverableItemsPurges    InvestigateBulkCancelNewResponseActionParamsMoveDestination = "RecoverableItemsPurges"
)

func (r InvestigateBulkCancelNewResponseActionParamsMoveDestination) IsKnown() bool {
	switch r {
	case InvestigateBulkCancelNewResponseActionParamsMoveDestinationInbox, InvestigateBulkCancelNewResponseActionParamsMoveDestinationJunkEmail, InvestigateBulkCancelNewResponseActionParamsMoveDestinationDeletedItems, InvestigateBulkCancelNewResponseActionParamsMoveDestinationRecoverableItemsDeletions, InvestigateBulkCancelNewResponseActionParamsMoveDestinationRecoverableItemsPurges:
		return true
	}
	return false
}

type InvestigateBulkCancelNewResponseActionParamsMoveType string

const (
	InvestigateBulkCancelNewResponseActionParamsMoveTypeMove InvestigateBulkCancelNewResponseActionParamsMoveType = "MOVE"
)

func (r InvestigateBulkCancelNewResponseActionParamsMoveType) IsKnown() bool {
	switch r {
	case InvestigateBulkCancelNewResponseActionParamsMoveTypeMove:
		return true
	}
	return false
}

type InvestigateBulkCancelNewResponseActionParamsMoveExpectedDisposition string

const (
	InvestigateBulkCancelNewResponseActionParamsMoveExpectedDispositionMalicious    InvestigateBulkCancelNewResponseActionParamsMoveExpectedDisposition = "MALICIOUS"
	InvestigateBulkCancelNewResponseActionParamsMoveExpectedDispositionMaliciousBec InvestigateBulkCancelNewResponseActionParamsMoveExpectedDisposition = "MALICIOUS-BEC"
	InvestigateBulkCancelNewResponseActionParamsMoveExpectedDispositionSuspicious   InvestigateBulkCancelNewResponseActionParamsMoveExpectedDisposition = "SUSPICIOUS"
	InvestigateBulkCancelNewResponseActionParamsMoveExpectedDispositionSpoof        InvestigateBulkCancelNewResponseActionParamsMoveExpectedDisposition = "SPOOF"
	InvestigateBulkCancelNewResponseActionParamsMoveExpectedDispositionSpam         InvestigateBulkCancelNewResponseActionParamsMoveExpectedDisposition = "SPAM"
	InvestigateBulkCancelNewResponseActionParamsMoveExpectedDispositionBulk         InvestigateBulkCancelNewResponseActionParamsMoveExpectedDisposition = "BULK"
	InvestigateBulkCancelNewResponseActionParamsMoveExpectedDispositionEncrypted    InvestigateBulkCancelNewResponseActionParamsMoveExpectedDisposition = "ENCRYPTED"
	InvestigateBulkCancelNewResponseActionParamsMoveExpectedDispositionExternal     InvestigateBulkCancelNewResponseActionParamsMoveExpectedDisposition = "EXTERNAL"
	InvestigateBulkCancelNewResponseActionParamsMoveExpectedDispositionUnknown      InvestigateBulkCancelNewResponseActionParamsMoveExpectedDisposition = "UNKNOWN"
	InvestigateBulkCancelNewResponseActionParamsMoveExpectedDispositionNone         InvestigateBulkCancelNewResponseActionParamsMoveExpectedDisposition = "NONE"
)

func (r InvestigateBulkCancelNewResponseActionParamsMoveExpectedDisposition) IsKnown() bool {
	switch r {
	case InvestigateBulkCancelNewResponseActionParamsMoveExpectedDispositionMalicious, InvestigateBulkCancelNewResponseActionParamsMoveExpectedDispositionMaliciousBec, InvestigateBulkCancelNewResponseActionParamsMoveExpectedDispositionSuspicious, InvestigateBulkCancelNewResponseActionParamsMoveExpectedDispositionSpoof, InvestigateBulkCancelNewResponseActionParamsMoveExpectedDispositionSpam, InvestigateBulkCancelNewResponseActionParamsMoveExpectedDispositionBulk, InvestigateBulkCancelNewResponseActionParamsMoveExpectedDispositionEncrypted, InvestigateBulkCancelNewResponseActionParamsMoveExpectedDispositionExternal, InvestigateBulkCancelNewResponseActionParamsMoveExpectedDispositionUnknown, InvestigateBulkCancelNewResponseActionParamsMoveExpectedDispositionNone:
		return true
	}
	return false
}

type InvestigateBulkCancelNewResponseActionParamsRelease struct {
	Type InvestigateBulkCancelNewResponseActionParamsReleaseType `json:"type" api:"required"`
	JSON investigateBulkCancelNewResponseActionParamsReleaseJSON `json:"-"`
}

// investigateBulkCancelNewResponseActionParamsReleaseJSON contains the JSON
// metadata for the struct [InvestigateBulkCancelNewResponseActionParamsRelease]
type investigateBulkCancelNewResponseActionParamsReleaseJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateBulkCancelNewResponseActionParamsRelease) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkCancelNewResponseActionParamsReleaseJSON) RawJSON() string {
	return r.raw
}

func (r InvestigateBulkCancelNewResponseActionParamsRelease) implementsInvestigateBulkCancelNewResponseActionParams() {
}

type InvestigateBulkCancelNewResponseActionParamsReleaseType string

const (
	InvestigateBulkCancelNewResponseActionParamsReleaseTypeRelease InvestigateBulkCancelNewResponseActionParamsReleaseType = "RELEASE"
)

func (r InvestigateBulkCancelNewResponseActionParamsReleaseType) IsKnown() bool {
	switch r {
	case InvestigateBulkCancelNewResponseActionParamsReleaseTypeRelease:
		return true
	}
	return false
}

type InvestigateBulkCancelNewResponseActionParamsType string

const (
	InvestigateBulkCancelNewResponseActionParamsTypeMove    InvestigateBulkCancelNewResponseActionParamsType = "MOVE"
	InvestigateBulkCancelNewResponseActionParamsTypeRelease InvestigateBulkCancelNewResponseActionParamsType = "RELEASE"
)

func (r InvestigateBulkCancelNewResponseActionParamsType) IsKnown() bool {
	switch r {
	case InvestigateBulkCancelNewResponseActionParamsTypeMove, InvestigateBulkCancelNewResponseActionParamsTypeRelease:
		return true
	}
	return false
}

type InvestigateBulkCancelNewResponseActionParamsDestination string

const (
	InvestigateBulkCancelNewResponseActionParamsDestinationInbox                     InvestigateBulkCancelNewResponseActionParamsDestination = "Inbox"
	InvestigateBulkCancelNewResponseActionParamsDestinationJunkEmail                 InvestigateBulkCancelNewResponseActionParamsDestination = "JunkEmail"
	InvestigateBulkCancelNewResponseActionParamsDestinationDeletedItems              InvestigateBulkCancelNewResponseActionParamsDestination = "DeletedItems"
	InvestigateBulkCancelNewResponseActionParamsDestinationRecoverableItemsDeletions InvestigateBulkCancelNewResponseActionParamsDestination = "RecoverableItemsDeletions"
	InvestigateBulkCancelNewResponseActionParamsDestinationRecoverableItemsPurges    InvestigateBulkCancelNewResponseActionParamsDestination = "RecoverableItemsPurges"
)

func (r InvestigateBulkCancelNewResponseActionParamsDestination) IsKnown() bool {
	switch r {
	case InvestigateBulkCancelNewResponseActionParamsDestinationInbox, InvestigateBulkCancelNewResponseActionParamsDestinationJunkEmail, InvestigateBulkCancelNewResponseActionParamsDestinationDeletedItems, InvestigateBulkCancelNewResponseActionParamsDestinationRecoverableItemsDeletions, InvestigateBulkCancelNewResponseActionParamsDestinationRecoverableItemsPurges:
		return true
	}
	return false
}

type InvestigateBulkCancelNewResponseActionParamsExpectedDisposition string

const (
	InvestigateBulkCancelNewResponseActionParamsExpectedDispositionMalicious    InvestigateBulkCancelNewResponseActionParamsExpectedDisposition = "MALICIOUS"
	InvestigateBulkCancelNewResponseActionParamsExpectedDispositionMaliciousBec InvestigateBulkCancelNewResponseActionParamsExpectedDisposition = "MALICIOUS-BEC"
	InvestigateBulkCancelNewResponseActionParamsExpectedDispositionSuspicious   InvestigateBulkCancelNewResponseActionParamsExpectedDisposition = "SUSPICIOUS"
	InvestigateBulkCancelNewResponseActionParamsExpectedDispositionSpoof        InvestigateBulkCancelNewResponseActionParamsExpectedDisposition = "SPOOF"
	InvestigateBulkCancelNewResponseActionParamsExpectedDispositionSpam         InvestigateBulkCancelNewResponseActionParamsExpectedDisposition = "SPAM"
	InvestigateBulkCancelNewResponseActionParamsExpectedDispositionBulk         InvestigateBulkCancelNewResponseActionParamsExpectedDisposition = "BULK"
	InvestigateBulkCancelNewResponseActionParamsExpectedDispositionEncrypted    InvestigateBulkCancelNewResponseActionParamsExpectedDisposition = "ENCRYPTED"
	InvestigateBulkCancelNewResponseActionParamsExpectedDispositionExternal     InvestigateBulkCancelNewResponseActionParamsExpectedDisposition = "EXTERNAL"
	InvestigateBulkCancelNewResponseActionParamsExpectedDispositionUnknown      InvestigateBulkCancelNewResponseActionParamsExpectedDisposition = "UNKNOWN"
	InvestigateBulkCancelNewResponseActionParamsExpectedDispositionNone         InvestigateBulkCancelNewResponseActionParamsExpectedDisposition = "NONE"
)

func (r InvestigateBulkCancelNewResponseActionParamsExpectedDisposition) IsKnown() bool {
	switch r {
	case InvestigateBulkCancelNewResponseActionParamsExpectedDispositionMalicious, InvestigateBulkCancelNewResponseActionParamsExpectedDispositionMaliciousBec, InvestigateBulkCancelNewResponseActionParamsExpectedDispositionSuspicious, InvestigateBulkCancelNewResponseActionParamsExpectedDispositionSpoof, InvestigateBulkCancelNewResponseActionParamsExpectedDispositionSpam, InvestigateBulkCancelNewResponseActionParamsExpectedDispositionBulk, InvestigateBulkCancelNewResponseActionParamsExpectedDispositionEncrypted, InvestigateBulkCancelNewResponseActionParamsExpectedDispositionExternal, InvestigateBulkCancelNewResponseActionParamsExpectedDispositionUnknown, InvestigateBulkCancelNewResponseActionParamsExpectedDispositionNone:
		return true
	}
	return false
}

type InvestigateBulkCancelNewResponseActionType string

const (
	InvestigateBulkCancelNewResponseActionTypeMove    InvestigateBulkCancelNewResponseActionType = "MOVE"
	InvestigateBulkCancelNewResponseActionTypeRelease InvestigateBulkCancelNewResponseActionType = "RELEASE"
)

func (r InvestigateBulkCancelNewResponseActionType) IsKnown() bool {
	switch r {
	case InvestigateBulkCancelNewResponseActionTypeMove, InvestigateBulkCancelNewResponseActionTypeRelease:
		return true
	}
	return false
}

type InvestigateBulkCancelNewResponseSearchParams struct {
	// Deprecated, use `GET /investigate/{investigate_id}/action_log` instead. End of
	// life: November 1, 2026.
	//
	// Deprecated: deprecated
	ActionLog bool   `json:"action_log"`
	AlertID   string `json:"alert_id" api:"nullable"`
	// Delivery status of the message.
	DeliveryStatus InvestigateBulkCancelNewResponseSearchParamsDeliveryStatus `json:"delivery_status"`
	DetectionsOnly bool                                                       `json:"detections_only"`
	Domain         string                                                     `json:"domain" api:"nullable"`
	// End of search date range
	End              time.Time                                                    `json:"end" format:"date-time"`
	ExactSubject     string                                                       `json:"exact_subject" api:"nullable"`
	FinalDisposition InvestigateBulkCancelNewResponseSearchParamsFinalDisposition `json:"final_disposition"`
	MessageAction    InvestigateBulkCancelNewResponseSearchParamsMessageAction    `json:"message_action" api:"nullable"`
	MessageID        string                                                       `json:"message_id" api:"nullable"`
	Metric           string                                                       `json:"metric" api:"nullable"`
	Query            string                                                       `json:"query" api:"nullable"`
	Recipient        string                                                       `json:"recipient" api:"nullable"`
	Sender           string                                                       `json:"sender" api:"nullable"`
	// Beginning of search date range
	Start       time.Time                                        `json:"start" format:"date-time"`
	Subject     string                                           `json:"subject" api:"nullable"`
	Submissions bool                                             `json:"submissions"`
	JSON        investigateBulkCancelNewResponseSearchParamsJSON `json:"-"`
}

// investigateBulkCancelNewResponseSearchParamsJSON contains the JSON metadata for
// the struct [InvestigateBulkCancelNewResponseSearchParams]
type investigateBulkCancelNewResponseSearchParamsJSON struct {
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

func (r *InvestigateBulkCancelNewResponseSearchParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkCancelNewResponseSearchParamsJSON) RawJSON() string {
	return r.raw
}

// Delivery status of the message.
type InvestigateBulkCancelNewResponseSearchParamsDeliveryStatus string

const (
	InvestigateBulkCancelNewResponseSearchParamsDeliveryStatusDelivered   InvestigateBulkCancelNewResponseSearchParamsDeliveryStatus = "delivered"
	InvestigateBulkCancelNewResponseSearchParamsDeliveryStatusMoved       InvestigateBulkCancelNewResponseSearchParamsDeliveryStatus = "moved"
	InvestigateBulkCancelNewResponseSearchParamsDeliveryStatusQuarantined InvestigateBulkCancelNewResponseSearchParamsDeliveryStatus = "quarantined"
	InvestigateBulkCancelNewResponseSearchParamsDeliveryStatusRejected    InvestigateBulkCancelNewResponseSearchParamsDeliveryStatus = "rejected"
	InvestigateBulkCancelNewResponseSearchParamsDeliveryStatusDeferred    InvestigateBulkCancelNewResponseSearchParamsDeliveryStatus = "deferred"
	InvestigateBulkCancelNewResponseSearchParamsDeliveryStatusBounced     InvestigateBulkCancelNewResponseSearchParamsDeliveryStatus = "bounced"
	InvestigateBulkCancelNewResponseSearchParamsDeliveryStatusQueued      InvestigateBulkCancelNewResponseSearchParamsDeliveryStatus = "queued"
)

func (r InvestigateBulkCancelNewResponseSearchParamsDeliveryStatus) IsKnown() bool {
	switch r {
	case InvestigateBulkCancelNewResponseSearchParamsDeliveryStatusDelivered, InvestigateBulkCancelNewResponseSearchParamsDeliveryStatusMoved, InvestigateBulkCancelNewResponseSearchParamsDeliveryStatusQuarantined, InvestigateBulkCancelNewResponseSearchParamsDeliveryStatusRejected, InvestigateBulkCancelNewResponseSearchParamsDeliveryStatusDeferred, InvestigateBulkCancelNewResponseSearchParamsDeliveryStatusBounced, InvestigateBulkCancelNewResponseSearchParamsDeliveryStatusQueued:
		return true
	}
	return false
}

type InvestigateBulkCancelNewResponseSearchParamsFinalDisposition string

const (
	InvestigateBulkCancelNewResponseSearchParamsFinalDispositionMalicious    InvestigateBulkCancelNewResponseSearchParamsFinalDisposition = "MALICIOUS"
	InvestigateBulkCancelNewResponseSearchParamsFinalDispositionMaliciousBec InvestigateBulkCancelNewResponseSearchParamsFinalDisposition = "MALICIOUS-BEC"
	InvestigateBulkCancelNewResponseSearchParamsFinalDispositionSuspicious   InvestigateBulkCancelNewResponseSearchParamsFinalDisposition = "SUSPICIOUS"
	InvestigateBulkCancelNewResponseSearchParamsFinalDispositionSpoof        InvestigateBulkCancelNewResponseSearchParamsFinalDisposition = "SPOOF"
	InvestigateBulkCancelNewResponseSearchParamsFinalDispositionSpam         InvestigateBulkCancelNewResponseSearchParamsFinalDisposition = "SPAM"
	InvestigateBulkCancelNewResponseSearchParamsFinalDispositionBulk         InvestigateBulkCancelNewResponseSearchParamsFinalDisposition = "BULK"
	InvestigateBulkCancelNewResponseSearchParamsFinalDispositionEncrypted    InvestigateBulkCancelNewResponseSearchParamsFinalDisposition = "ENCRYPTED"
	InvestigateBulkCancelNewResponseSearchParamsFinalDispositionExternal     InvestigateBulkCancelNewResponseSearchParamsFinalDisposition = "EXTERNAL"
	InvestigateBulkCancelNewResponseSearchParamsFinalDispositionUnknown      InvestigateBulkCancelNewResponseSearchParamsFinalDisposition = "UNKNOWN"
	InvestigateBulkCancelNewResponseSearchParamsFinalDispositionNone         InvestigateBulkCancelNewResponseSearchParamsFinalDisposition = "NONE"
)

func (r InvestigateBulkCancelNewResponseSearchParamsFinalDisposition) IsKnown() bool {
	switch r {
	case InvestigateBulkCancelNewResponseSearchParamsFinalDispositionMalicious, InvestigateBulkCancelNewResponseSearchParamsFinalDispositionMaliciousBec, InvestigateBulkCancelNewResponseSearchParamsFinalDispositionSuspicious, InvestigateBulkCancelNewResponseSearchParamsFinalDispositionSpoof, InvestigateBulkCancelNewResponseSearchParamsFinalDispositionSpam, InvestigateBulkCancelNewResponseSearchParamsFinalDispositionBulk, InvestigateBulkCancelNewResponseSearchParamsFinalDispositionEncrypted, InvestigateBulkCancelNewResponseSearchParamsFinalDispositionExternal, InvestigateBulkCancelNewResponseSearchParamsFinalDispositionUnknown, InvestigateBulkCancelNewResponseSearchParamsFinalDispositionNone:
		return true
	}
	return false
}

type InvestigateBulkCancelNewResponseSearchParamsMessageAction string

const (
	InvestigateBulkCancelNewResponseSearchParamsMessageActionPreview            InvestigateBulkCancelNewResponseSearchParamsMessageAction = "PREVIEW"
	InvestigateBulkCancelNewResponseSearchParamsMessageActionQuarantineReleased InvestigateBulkCancelNewResponseSearchParamsMessageAction = "QUARANTINE_RELEASED"
	InvestigateBulkCancelNewResponseSearchParamsMessageActionMoved              InvestigateBulkCancelNewResponseSearchParamsMessageAction = "MOVED"
)

func (r InvestigateBulkCancelNewResponseSearchParamsMessageAction) IsKnown() bool {
	switch r {
	case InvestigateBulkCancelNewResponseSearchParamsMessageActionPreview, InvestigateBulkCancelNewResponseSearchParamsMessageActionQuarantineReleased, InvestigateBulkCancelNewResponseSearchParamsMessageActionMoved:
		return true
	}
	return false
}

type InvestigateBulkCancelNewResponseStatus string

const (
	InvestigateBulkCancelNewResponseStatusPending     InvestigateBulkCancelNewResponseStatus = "PENDING"
	InvestigateBulkCancelNewResponseStatusDiscovering InvestigateBulkCancelNewResponseStatus = "DISCOVERING"
	InvestigateBulkCancelNewResponseStatusProcessing  InvestigateBulkCancelNewResponseStatus = "PROCESSING"
	InvestigateBulkCancelNewResponseStatusCompleted   InvestigateBulkCancelNewResponseStatus = "COMPLETED"
	InvestigateBulkCancelNewResponseStatusFailed      InvestigateBulkCancelNewResponseStatus = "FAILED"
	InvestigateBulkCancelNewResponseStatusCancelled   InvestigateBulkCancelNewResponseStatus = "CANCELLED"
	InvestigateBulkCancelNewResponseStatusSkipped     InvestigateBulkCancelNewResponseStatus = "SKIPPED"
)

func (r InvestigateBulkCancelNewResponseStatus) IsKnown() bool {
	switch r {
	case InvestigateBulkCancelNewResponseStatusPending, InvestigateBulkCancelNewResponseStatusDiscovering, InvestigateBulkCancelNewResponseStatusProcessing, InvestigateBulkCancelNewResponseStatusCompleted, InvestigateBulkCancelNewResponseStatusFailed, InvestigateBulkCancelNewResponseStatusCancelled, InvestigateBulkCancelNewResponseStatusSkipped:
		return true
	}
	return false
}

type InvestigateBulkCancelNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type InvestigateBulkCancelNewResponseEnvelope struct {
	Errors   []InvestigateBulkCancelNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []InvestigateBulkCancelNewResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   InvestigateBulkCancelNewResponse                   `json:"result" api:"required"`
	// Whether the API call was successful.
	Success InvestigateBulkCancelNewResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    investigateBulkCancelNewResponseEnvelopeJSON    `json:"-"`
}

// investigateBulkCancelNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [InvestigateBulkCancelNewResponseEnvelope]
type investigateBulkCancelNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateBulkCancelNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkCancelNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InvestigateBulkCancelNewResponseEnvelopeErrors struct {
	Code             int64                                                `json:"code" api:"required"`
	Message          string                                               `json:"message" api:"required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           InvestigateBulkCancelNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             investigateBulkCancelNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// investigateBulkCancelNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [InvestigateBulkCancelNewResponseEnvelopeErrors]
type investigateBulkCancelNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvestigateBulkCancelNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkCancelNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type InvestigateBulkCancelNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    investigateBulkCancelNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// investigateBulkCancelNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [InvestigateBulkCancelNewResponseEnvelopeErrorsSource]
type investigateBulkCancelNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateBulkCancelNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkCancelNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type InvestigateBulkCancelNewResponseEnvelopeMessages struct {
	Code             int64                                                  `json:"code" api:"required"`
	Message          string                                                 `json:"message" api:"required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           InvestigateBulkCancelNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             investigateBulkCancelNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// investigateBulkCancelNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [InvestigateBulkCancelNewResponseEnvelopeMessages]
type investigateBulkCancelNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvestigateBulkCancelNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkCancelNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type InvestigateBulkCancelNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    investigateBulkCancelNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// investigateBulkCancelNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [InvestigateBulkCancelNewResponseEnvelopeMessagesSource]
type investigateBulkCancelNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateBulkCancelNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkCancelNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type InvestigateBulkCancelNewResponseEnvelopeSuccess bool

const (
	InvestigateBulkCancelNewResponseEnvelopeSuccessTrue InvestigateBulkCancelNewResponseEnvelopeSuccess = true
)

func (r InvestigateBulkCancelNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case InvestigateBulkCancelNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
