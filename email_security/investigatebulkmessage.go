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

// InvestigateBulkMessageService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvestigateBulkMessageService] method instead.
type InvestigateBulkMessageService struct {
	Options []option.RequestOption
}

// NewInvestigateBulkMessageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInvestigateBulkMessageService(opts ...option.RequestOption) (r *InvestigateBulkMessageService) {
	r = &InvestigateBulkMessageService{}
	r.Options = opts
	return
}

// Returns the individual messages associated with a bulk action job, including
// their processing status.
func (r *InvestigateBulkMessageService) List(ctx context.Context, jobID string, params InvestigateBulkMessageListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[InvestigateBulkMessageListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/bulk/%s/messages", params.AccountID, jobID)
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

// Returns the individual messages associated with a bulk action job, including
// their processing status.
func (r *InvestigateBulkMessageService) ListAutoPaging(ctx context.Context, jobID string, params InvestigateBulkMessageListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[InvestigateBulkMessageListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, jobID, params, opts...))
}

type InvestigateBulkMessageListResponse struct {
	ActionParams   InvestigateBulkMessageListResponseActionParams `json:"action_params" api:"required"`
	ActionType     InvestigateBulkMessageListResponseActionType   `json:"action_type" api:"required"`
	CreatedAt      time.Time                                      `json:"created_at" api:"required" format:"date-time"`
	MessageID      string                                         `json:"message_id" api:"required" format:"uuid"`
	PostfixID      string                                         `json:"postfix_id" api:"required"`
	RetryCount     int64                                          `json:"retry_count" api:"required"`
	Status         InvestigateBulkMessageListResponseStatus       `json:"status" api:"required"`
	AlertID        string                                         `json:"alert_id" api:"nullable"`
	EmailMessageID string                                         `json:"email_message_id" api:"nullable"`
	Message        InvestigateBulkMessageListResponseMessage      `json:"message"`
	ProcessedAt    time.Time                                      `json:"processed_at" api:"nullable" format:"date-time"`
	// When to retry the action if it failed.
	RetryAfter    time.Time                              `json:"retry_after" api:"nullable" format:"date-time"`
	StatusMessage string                                 `json:"status_message" api:"nullable"`
	JSON          investigateBulkMessageListResponseJSON `json:"-"`
}

// investigateBulkMessageListResponseJSON contains the JSON metadata for the struct
// [InvestigateBulkMessageListResponse]
type investigateBulkMessageListResponseJSON struct {
	ActionParams   apijson.Field
	ActionType     apijson.Field
	CreatedAt      apijson.Field
	MessageID      apijson.Field
	PostfixID      apijson.Field
	RetryCount     apijson.Field
	Status         apijson.Field
	AlertID        apijson.Field
	EmailMessageID apijson.Field
	Message        apijson.Field
	ProcessedAt    apijson.Field
	RetryAfter     apijson.Field
	StatusMessage  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InvestigateBulkMessageListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkMessageListResponseJSON) RawJSON() string {
	return r.raw
}

type InvestigateBulkMessageListResponseActionParams struct {
	ClientRecipient     string                                                            `json:"client_recipient" api:"required"`
	Type                InvestigateBulkMessageListResponseActionParamsType                `json:"type" api:"required"`
	Destination         InvestigateBulkMessageListResponseActionParamsDestination         `json:"destination"`
	ExpectedDisposition InvestigateBulkMessageListResponseActionParamsExpectedDisposition `json:"expected_disposition"`
	JSON                investigateBulkMessageListResponseActionParamsJSON                `json:"-"`
	union               InvestigateBulkMessageListResponseActionParamsUnion
}

// investigateBulkMessageListResponseActionParamsJSON contains the JSON metadata
// for the struct [InvestigateBulkMessageListResponseActionParams]
type investigateBulkMessageListResponseActionParamsJSON struct {
	ClientRecipient     apijson.Field
	Type                apijson.Field
	Destination         apijson.Field
	ExpectedDisposition apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r investigateBulkMessageListResponseActionParamsJSON) RawJSON() string {
	return r.raw
}

func (r *InvestigateBulkMessageListResponseActionParams) UnmarshalJSON(data []byte) (err error) {
	*r = InvestigateBulkMessageListResponseActionParams{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [InvestigateBulkMessageListResponseActionParamsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [InvestigateBulkMessageListResponseActionParamsMove],
// [InvestigateBulkMessageListResponseActionParamsRelease].
func (r InvestigateBulkMessageListResponseActionParams) AsUnion() InvestigateBulkMessageListResponseActionParamsUnion {
	return r.union
}

// Union satisfied by [InvestigateBulkMessageListResponseActionParamsMove] or
// [InvestigateBulkMessageListResponseActionParamsRelease].
type InvestigateBulkMessageListResponseActionParamsUnion interface {
	implementsInvestigateBulkMessageListResponseActionParams()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InvestigateBulkMessageListResponseActionParamsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvestigateBulkMessageListResponseActionParamsMove{}),
			DiscriminatorValue: "MOVE",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvestigateBulkMessageListResponseActionParamsRelease{}),
			DiscriminatorValue: "RELEASE",
		},
	)
}

type InvestigateBulkMessageListResponseActionParamsMove struct {
	ClientRecipient     string                                                                `json:"client_recipient" api:"required"`
	Destination         InvestigateBulkMessageListResponseActionParamsMoveDestination         `json:"destination" api:"required"`
	Type                InvestigateBulkMessageListResponseActionParamsMoveType                `json:"type" api:"required"`
	ExpectedDisposition InvestigateBulkMessageListResponseActionParamsMoveExpectedDisposition `json:"expected_disposition"`
	JSON                investigateBulkMessageListResponseActionParamsMoveJSON                `json:"-"`
}

// investigateBulkMessageListResponseActionParamsMoveJSON contains the JSON
// metadata for the struct [InvestigateBulkMessageListResponseActionParamsMove]
type investigateBulkMessageListResponseActionParamsMoveJSON struct {
	ClientRecipient     apijson.Field
	Destination         apijson.Field
	Type                apijson.Field
	ExpectedDisposition apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *InvestigateBulkMessageListResponseActionParamsMove) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkMessageListResponseActionParamsMoveJSON) RawJSON() string {
	return r.raw
}

func (r InvestigateBulkMessageListResponseActionParamsMove) implementsInvestigateBulkMessageListResponseActionParams() {
}

type InvestigateBulkMessageListResponseActionParamsMoveDestination string

const (
	InvestigateBulkMessageListResponseActionParamsMoveDestinationInbox                     InvestigateBulkMessageListResponseActionParamsMoveDestination = "Inbox"
	InvestigateBulkMessageListResponseActionParamsMoveDestinationJunkEmail                 InvestigateBulkMessageListResponseActionParamsMoveDestination = "JunkEmail"
	InvestigateBulkMessageListResponseActionParamsMoveDestinationDeletedItems              InvestigateBulkMessageListResponseActionParamsMoveDestination = "DeletedItems"
	InvestigateBulkMessageListResponseActionParamsMoveDestinationRecoverableItemsDeletions InvestigateBulkMessageListResponseActionParamsMoveDestination = "RecoverableItemsDeletions"
	InvestigateBulkMessageListResponseActionParamsMoveDestinationRecoverableItemsPurges    InvestigateBulkMessageListResponseActionParamsMoveDestination = "RecoverableItemsPurges"
)

func (r InvestigateBulkMessageListResponseActionParamsMoveDestination) IsKnown() bool {
	switch r {
	case InvestigateBulkMessageListResponseActionParamsMoveDestinationInbox, InvestigateBulkMessageListResponseActionParamsMoveDestinationJunkEmail, InvestigateBulkMessageListResponseActionParamsMoveDestinationDeletedItems, InvestigateBulkMessageListResponseActionParamsMoveDestinationRecoverableItemsDeletions, InvestigateBulkMessageListResponseActionParamsMoveDestinationRecoverableItemsPurges:
		return true
	}
	return false
}

type InvestigateBulkMessageListResponseActionParamsMoveType string

const (
	InvestigateBulkMessageListResponseActionParamsMoveTypeMove InvestigateBulkMessageListResponseActionParamsMoveType = "MOVE"
)

func (r InvestigateBulkMessageListResponseActionParamsMoveType) IsKnown() bool {
	switch r {
	case InvestigateBulkMessageListResponseActionParamsMoveTypeMove:
		return true
	}
	return false
}

type InvestigateBulkMessageListResponseActionParamsMoveExpectedDisposition string

const (
	InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionMalicious    InvestigateBulkMessageListResponseActionParamsMoveExpectedDisposition = "MALICIOUS"
	InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionMaliciousBec InvestigateBulkMessageListResponseActionParamsMoveExpectedDisposition = "MALICIOUS-BEC"
	InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionSuspicious   InvestigateBulkMessageListResponseActionParamsMoveExpectedDisposition = "SUSPICIOUS"
	InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionSpoof        InvestigateBulkMessageListResponseActionParamsMoveExpectedDisposition = "SPOOF"
	InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionSpam         InvestigateBulkMessageListResponseActionParamsMoveExpectedDisposition = "SPAM"
	InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionBulk         InvestigateBulkMessageListResponseActionParamsMoveExpectedDisposition = "BULK"
	InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionEncrypted    InvestigateBulkMessageListResponseActionParamsMoveExpectedDisposition = "ENCRYPTED"
	InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionExternal     InvestigateBulkMessageListResponseActionParamsMoveExpectedDisposition = "EXTERNAL"
	InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionUnknown      InvestigateBulkMessageListResponseActionParamsMoveExpectedDisposition = "UNKNOWN"
	InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionNone         InvestigateBulkMessageListResponseActionParamsMoveExpectedDisposition = "NONE"
)

func (r InvestigateBulkMessageListResponseActionParamsMoveExpectedDisposition) IsKnown() bool {
	switch r {
	case InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionMalicious, InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionMaliciousBec, InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionSuspicious, InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionSpoof, InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionSpam, InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionBulk, InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionEncrypted, InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionExternal, InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionUnknown, InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionNone:
		return true
	}
	return false
}

type InvestigateBulkMessageListResponseActionParamsRelease struct {
	ClientRecipient string                                                    `json:"client_recipient" api:"required"`
	Type            InvestigateBulkMessageListResponseActionParamsReleaseType `json:"type" api:"required"`
	JSON            investigateBulkMessageListResponseActionParamsReleaseJSON `json:"-"`
}

// investigateBulkMessageListResponseActionParamsReleaseJSON contains the JSON
// metadata for the struct [InvestigateBulkMessageListResponseActionParamsRelease]
type investigateBulkMessageListResponseActionParamsReleaseJSON struct {
	ClientRecipient apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InvestigateBulkMessageListResponseActionParamsRelease) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkMessageListResponseActionParamsReleaseJSON) RawJSON() string {
	return r.raw
}

func (r InvestigateBulkMessageListResponseActionParamsRelease) implementsInvestigateBulkMessageListResponseActionParams() {
}

type InvestigateBulkMessageListResponseActionParamsReleaseType string

const (
	InvestigateBulkMessageListResponseActionParamsReleaseTypeRelease InvestigateBulkMessageListResponseActionParamsReleaseType = "RELEASE"
)

func (r InvestigateBulkMessageListResponseActionParamsReleaseType) IsKnown() bool {
	switch r {
	case InvestigateBulkMessageListResponseActionParamsReleaseTypeRelease:
		return true
	}
	return false
}

type InvestigateBulkMessageListResponseActionParamsType string

const (
	InvestigateBulkMessageListResponseActionParamsTypeMove    InvestigateBulkMessageListResponseActionParamsType = "MOVE"
	InvestigateBulkMessageListResponseActionParamsTypeRelease InvestigateBulkMessageListResponseActionParamsType = "RELEASE"
)

func (r InvestigateBulkMessageListResponseActionParamsType) IsKnown() bool {
	switch r {
	case InvestigateBulkMessageListResponseActionParamsTypeMove, InvestigateBulkMessageListResponseActionParamsTypeRelease:
		return true
	}
	return false
}

type InvestigateBulkMessageListResponseActionParamsDestination string

const (
	InvestigateBulkMessageListResponseActionParamsDestinationInbox                     InvestigateBulkMessageListResponseActionParamsDestination = "Inbox"
	InvestigateBulkMessageListResponseActionParamsDestinationJunkEmail                 InvestigateBulkMessageListResponseActionParamsDestination = "JunkEmail"
	InvestigateBulkMessageListResponseActionParamsDestinationDeletedItems              InvestigateBulkMessageListResponseActionParamsDestination = "DeletedItems"
	InvestigateBulkMessageListResponseActionParamsDestinationRecoverableItemsDeletions InvestigateBulkMessageListResponseActionParamsDestination = "RecoverableItemsDeletions"
	InvestigateBulkMessageListResponseActionParamsDestinationRecoverableItemsPurges    InvestigateBulkMessageListResponseActionParamsDestination = "RecoverableItemsPurges"
)

func (r InvestigateBulkMessageListResponseActionParamsDestination) IsKnown() bool {
	switch r {
	case InvestigateBulkMessageListResponseActionParamsDestinationInbox, InvestigateBulkMessageListResponseActionParamsDestinationJunkEmail, InvestigateBulkMessageListResponseActionParamsDestinationDeletedItems, InvestigateBulkMessageListResponseActionParamsDestinationRecoverableItemsDeletions, InvestigateBulkMessageListResponseActionParamsDestinationRecoverableItemsPurges:
		return true
	}
	return false
}

type InvestigateBulkMessageListResponseActionParamsExpectedDisposition string

const (
	InvestigateBulkMessageListResponseActionParamsExpectedDispositionMalicious    InvestigateBulkMessageListResponseActionParamsExpectedDisposition = "MALICIOUS"
	InvestigateBulkMessageListResponseActionParamsExpectedDispositionMaliciousBec InvestigateBulkMessageListResponseActionParamsExpectedDisposition = "MALICIOUS-BEC"
	InvestigateBulkMessageListResponseActionParamsExpectedDispositionSuspicious   InvestigateBulkMessageListResponseActionParamsExpectedDisposition = "SUSPICIOUS"
	InvestigateBulkMessageListResponseActionParamsExpectedDispositionSpoof        InvestigateBulkMessageListResponseActionParamsExpectedDisposition = "SPOOF"
	InvestigateBulkMessageListResponseActionParamsExpectedDispositionSpam         InvestigateBulkMessageListResponseActionParamsExpectedDisposition = "SPAM"
	InvestigateBulkMessageListResponseActionParamsExpectedDispositionBulk         InvestigateBulkMessageListResponseActionParamsExpectedDisposition = "BULK"
	InvestigateBulkMessageListResponseActionParamsExpectedDispositionEncrypted    InvestigateBulkMessageListResponseActionParamsExpectedDisposition = "ENCRYPTED"
	InvestigateBulkMessageListResponseActionParamsExpectedDispositionExternal     InvestigateBulkMessageListResponseActionParamsExpectedDisposition = "EXTERNAL"
	InvestigateBulkMessageListResponseActionParamsExpectedDispositionUnknown      InvestigateBulkMessageListResponseActionParamsExpectedDisposition = "UNKNOWN"
	InvestigateBulkMessageListResponseActionParamsExpectedDispositionNone         InvestigateBulkMessageListResponseActionParamsExpectedDisposition = "NONE"
)

func (r InvestigateBulkMessageListResponseActionParamsExpectedDisposition) IsKnown() bool {
	switch r {
	case InvestigateBulkMessageListResponseActionParamsExpectedDispositionMalicious, InvestigateBulkMessageListResponseActionParamsExpectedDispositionMaliciousBec, InvestigateBulkMessageListResponseActionParamsExpectedDispositionSuspicious, InvestigateBulkMessageListResponseActionParamsExpectedDispositionSpoof, InvestigateBulkMessageListResponseActionParamsExpectedDispositionSpam, InvestigateBulkMessageListResponseActionParamsExpectedDispositionBulk, InvestigateBulkMessageListResponseActionParamsExpectedDispositionEncrypted, InvestigateBulkMessageListResponseActionParamsExpectedDispositionExternal, InvestigateBulkMessageListResponseActionParamsExpectedDispositionUnknown, InvestigateBulkMessageListResponseActionParamsExpectedDispositionNone:
		return true
	}
	return false
}

type InvestigateBulkMessageListResponseActionType string

const (
	InvestigateBulkMessageListResponseActionTypeMove    InvestigateBulkMessageListResponseActionType = "MOVE"
	InvestigateBulkMessageListResponseActionTypeRelease InvestigateBulkMessageListResponseActionType = "RELEASE"
)

func (r InvestigateBulkMessageListResponseActionType) IsKnown() bool {
	switch r {
	case InvestigateBulkMessageListResponseActionTypeMove, InvestigateBulkMessageListResponseActionTypeRelease:
		return true
	}
	return false
}

type InvestigateBulkMessageListResponseStatus string

const (
	InvestigateBulkMessageListResponseStatusPending     InvestigateBulkMessageListResponseStatus = "PENDING"
	InvestigateBulkMessageListResponseStatusDiscovering InvestigateBulkMessageListResponseStatus = "DISCOVERING"
	InvestigateBulkMessageListResponseStatusProcessing  InvestigateBulkMessageListResponseStatus = "PROCESSING"
	InvestigateBulkMessageListResponseStatusCompleted   InvestigateBulkMessageListResponseStatus = "COMPLETED"
	InvestigateBulkMessageListResponseStatusFailed      InvestigateBulkMessageListResponseStatus = "FAILED"
	InvestigateBulkMessageListResponseStatusCancelled   InvestigateBulkMessageListResponseStatus = "CANCELLED"
	InvestigateBulkMessageListResponseStatusSkipped     InvestigateBulkMessageListResponseStatus = "SKIPPED"
)

func (r InvestigateBulkMessageListResponseStatus) IsKnown() bool {
	switch r {
	case InvestigateBulkMessageListResponseStatusPending, InvestigateBulkMessageListResponseStatusDiscovering, InvestigateBulkMessageListResponseStatusProcessing, InvestigateBulkMessageListResponseStatusCompleted, InvestigateBulkMessageListResponseStatusFailed, InvestigateBulkMessageListResponseStatusCancelled, InvestigateBulkMessageListResponseStatusSkipped:
		return true
	}
	return false
}

type InvestigateBulkMessageListResponseMessage struct {
	// Unique identifier for a message retrieved from investigation.
	ID string `json:"id" api:"required"`
	// Deprecated, use `GET /investigate/{investigate_id}/action_log` instead. End of
	// life: November 1, 2026.
	//
	// Deprecated: Use GET /investigate/{investigate_id}/action_log instead.
	ActionLog         []InvestigateBulkMessageListResponseMessageActionLog `json:"action_log" api:"required"`
	ClientRecipients  []string                                             `json:"client_recipients" api:"required"`
	DetectionReasons  []string                                             `json:"detection_reasons" api:"required"`
	IsPhishSubmission bool                                                 `json:"is_phish_submission" api:"required"`
	IsQuarantined     bool                                                 `json:"is_quarantined" api:"required"`
	// The identifier of the message.
	PostfixID string `json:"postfix_id" api:"required"`
	// Message processing properties.
	Properties InvestigateBulkMessageListResponseMessageProperties `json:"properties" api:"required"`
	// Deprecated, use `scanned_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: Use `scanned_at` instead.
	Ts               string                                                    `json:"ts" api:"required"`
	AlertID          string                                                    `json:"alert_id" api:"nullable"`
	DeliveryMode     InvestigateBulkMessageListResponseMessageDeliveryMode     `json:"delivery_mode"`
	DeliveryStatus   []InvestigateBulkMessageListResponseMessageDeliveryStatus `json:"delivery_status" api:"nullable"`
	EdfHash          string                                                    `json:"edf_hash" api:"nullable"`
	EnvelopeFrom     string                                                    `json:"envelope_from" api:"nullable"`
	EnvelopeTo       []string                                                  `json:"envelope_to" api:"nullable"`
	FinalDisposition InvestigateBulkMessageListResponseMessageFinalDisposition `json:"final_disposition"`
	// Deprecated, use the `findings` field from
	// `GET /investigate/{investigate_id}/detections` instead. End of life: November
	// 1, 2026. Detection findings for this message.
	//
	// Deprecated: Use the `findings` field from GET
	// /investigate/{investigate_id}/detections instead.
	Findings              []InvestigateBulkMessageListResponseMessageFinding `json:"findings" api:"nullable"`
	From                  string                                             `json:"from" api:"nullable"`
	FromName              string                                             `json:"from_name" api:"nullable"`
	HtmltextStructureHash string                                             `json:"htmltext_structure_hash" api:"nullable"`
	MessageID             string                                             `json:"message_id" api:"nullable"`
	// Post-delivery operations performed on this message.
	PostDeliveryOperations []InvestigateBulkMessageListResponseMessagePostDeliveryOperation `json:"post_delivery_operations" api:"nullable"`
	PostfixIDOutbound      string                                                           `json:"postfix_id_outbound" api:"nullable"`
	Replyto                string                                                           `json:"replyto" api:"nullable"`
	// When the message was scanned (UTC).
	ScannedAt time.Time `json:"scanned_at" api:"nullable" format:"date-time"`
	// When the message was sent (UTC).
	SentAt            time.Time                                           `json:"sent_at" api:"nullable" format:"date-time"`
	SentDate          string                                              `json:"sent_date" api:"nullable"`
	SmtpHeloServerIP  string                                              `json:"smtp_helo_server_ip" api:"nullable"`
	SmtpPreviousHopIP string                                              `json:"smtp_previous_hop_ip" api:"nullable"`
	Subject           string                                              `json:"subject" api:"nullable"`
	ThreatCategories  []string                                            `json:"threat_categories" api:"nullable"`
	To                []string                                            `json:"to" api:"nullable"`
	ToName            []string                                            `json:"to_name" api:"nullable"`
	Validation        InvestigateBulkMessageListResponseMessageValidation `json:"validation"`
	XOriginatingIP    string                                              `json:"x_originating_ip" api:"nullable"`
	JSON              investigateBulkMessageListResponseMessageJSON       `json:"-"`
}

// investigateBulkMessageListResponseMessageJSON contains the JSON metadata for the
// struct [InvestigateBulkMessageListResponseMessage]
type investigateBulkMessageListResponseMessageJSON struct {
	ID                     apijson.Field
	ActionLog              apijson.Field
	ClientRecipients       apijson.Field
	DetectionReasons       apijson.Field
	IsPhishSubmission      apijson.Field
	IsQuarantined          apijson.Field
	PostfixID              apijson.Field
	Properties             apijson.Field
	Ts                     apijson.Field
	AlertID                apijson.Field
	DeliveryMode           apijson.Field
	DeliveryStatus         apijson.Field
	EdfHash                apijson.Field
	EnvelopeFrom           apijson.Field
	EnvelopeTo             apijson.Field
	FinalDisposition       apijson.Field
	Findings               apijson.Field
	From                   apijson.Field
	FromName               apijson.Field
	HtmltextStructureHash  apijson.Field
	MessageID              apijson.Field
	PostDeliveryOperations apijson.Field
	PostfixIDOutbound      apijson.Field
	Replyto                apijson.Field
	ScannedAt              apijson.Field
	SentAt                 apijson.Field
	SentDate               apijson.Field
	SmtpHeloServerIP       apijson.Field
	SmtpPreviousHopIP      apijson.Field
	Subject                apijson.Field
	ThreatCategories       apijson.Field
	To                     apijson.Field
	ToName                 apijson.Field
	Validation             apijson.Field
	XOriginatingIP         apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *InvestigateBulkMessageListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkMessageListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type InvestigateBulkMessageListResponseMessageActionLog struct {
	// Timestamp when action completed.
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	// Type of action performed.
	Operation InvestigateBulkMessageListResponseMessageActionLogOperation `json:"operation" api:"required"`
	// Deprecated, use `completed_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: Use `completed_at` instead.
	CompletedTimestamp string `json:"completed_timestamp"`
	// Additional properties for the action.
	Properties InvestigateBulkMessageListResponseMessageActionLogProperties `json:"properties"`
	// Status of the action.
	Status string                                                 `json:"status" api:"nullable"`
	JSON   investigateBulkMessageListResponseMessageActionLogJSON `json:"-"`
}

// investigateBulkMessageListResponseMessageActionLogJSON contains the JSON
// metadata for the struct [InvestigateBulkMessageListResponseMessageActionLog]
type investigateBulkMessageListResponseMessageActionLogJSON struct {
	CompletedAt        apijson.Field
	Operation          apijson.Field
	CompletedTimestamp apijson.Field
	Properties         apijson.Field
	Status             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InvestigateBulkMessageListResponseMessageActionLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkMessageListResponseMessageActionLogJSON) RawJSON() string {
	return r.raw
}

// Type of action performed.
type InvestigateBulkMessageListResponseMessageActionLogOperation string

const (
	InvestigateBulkMessageListResponseMessageActionLogOperationMove              InvestigateBulkMessageListResponseMessageActionLogOperation = "MOVE"
	InvestigateBulkMessageListResponseMessageActionLogOperationRelease           InvestigateBulkMessageListResponseMessageActionLogOperation = "RELEASE"
	InvestigateBulkMessageListResponseMessageActionLogOperationReclassify        InvestigateBulkMessageListResponseMessageActionLogOperation = "RECLASSIFY"
	InvestigateBulkMessageListResponseMessageActionLogOperationSubmission        InvestigateBulkMessageListResponseMessageActionLogOperation = "SUBMISSION"
	InvestigateBulkMessageListResponseMessageActionLogOperationQuarantineRelease InvestigateBulkMessageListResponseMessageActionLogOperation = "QUARANTINE_RELEASE"
	InvestigateBulkMessageListResponseMessageActionLogOperationPreview           InvestigateBulkMessageListResponseMessageActionLogOperation = "PREVIEW"
)

func (r InvestigateBulkMessageListResponseMessageActionLogOperation) IsKnown() bool {
	switch r {
	case InvestigateBulkMessageListResponseMessageActionLogOperationMove, InvestigateBulkMessageListResponseMessageActionLogOperationRelease, InvestigateBulkMessageListResponseMessageActionLogOperationReclassify, InvestigateBulkMessageListResponseMessageActionLogOperationSubmission, InvestigateBulkMessageListResponseMessageActionLogOperationQuarantineRelease, InvestigateBulkMessageListResponseMessageActionLogOperationPreview:
		return true
	}
	return false
}

// Additional properties for the action.
type InvestigateBulkMessageListResponseMessageActionLogProperties struct {
	// Target folder for move operations.
	Folder string `json:"folder"`
	// User who requested the action.
	RequestedBy string                                                           `json:"requested_by"`
	JSON        investigateBulkMessageListResponseMessageActionLogPropertiesJSON `json:"-"`
}

// investigateBulkMessageListResponseMessageActionLogPropertiesJSON contains the
// JSON metadata for the struct
// [InvestigateBulkMessageListResponseMessageActionLogProperties]
type investigateBulkMessageListResponseMessageActionLogPropertiesJSON struct {
	Folder      apijson.Field
	RequestedBy apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateBulkMessageListResponseMessageActionLogProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkMessageListResponseMessageActionLogPropertiesJSON) RawJSON() string {
	return r.raw
}

// Message processing properties.
type InvestigateBulkMessageListResponseMessageProperties struct {
	// Pattern that allowlisted this message.
	AllowlistedPattern string `json:"allowlisted_pattern" api:"nullable"`
	// Type of allowlist pattern.
	AllowlistedPatternType InvestigateBulkMessageListResponseMessagePropertiesAllowlistedPatternType `json:"allowlisted_pattern_type" api:"nullable"`
	// Whether message was blocklisted.
	BlocklistedMessage bool `json:"blocklisted_message" api:"nullable"`
	// Pattern that blocklisted this message.
	BlocklistedPattern string `json:"blocklisted_pattern" api:"nullable"`
	// Legacy field for allowlist pattern type.
	WhitelistedPatternType InvestigateBulkMessageListResponseMessagePropertiesWhitelistedPatternType `json:"whitelisted_pattern_type" api:"nullable"`
	JSON                   investigateBulkMessageListResponseMessagePropertiesJSON                   `json:"-"`
}

// investigateBulkMessageListResponseMessagePropertiesJSON contains the JSON
// metadata for the struct [InvestigateBulkMessageListResponseMessageProperties]
type investigateBulkMessageListResponseMessagePropertiesJSON struct {
	AllowlistedPattern     apijson.Field
	AllowlistedPatternType apijson.Field
	BlocklistedMessage     apijson.Field
	BlocklistedPattern     apijson.Field
	WhitelistedPatternType apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *InvestigateBulkMessageListResponseMessageProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkMessageListResponseMessagePropertiesJSON) RawJSON() string {
	return r.raw
}

// Type of allowlist pattern.
type InvestigateBulkMessageListResponseMessagePropertiesAllowlistedPatternType string

const (
	InvestigateBulkMessageListResponseMessagePropertiesAllowlistedPatternTypeQuarantineRelease       InvestigateBulkMessageListResponseMessagePropertiesAllowlistedPatternType = "quarantine_release"
	InvestigateBulkMessageListResponseMessagePropertiesAllowlistedPatternTypeAcceptableSender        InvestigateBulkMessageListResponseMessagePropertiesAllowlistedPatternType = "acceptable_sender"
	InvestigateBulkMessageListResponseMessagePropertiesAllowlistedPatternTypeAllowedSender           InvestigateBulkMessageListResponseMessagePropertiesAllowlistedPatternType = "allowed_sender"
	InvestigateBulkMessageListResponseMessagePropertiesAllowlistedPatternTypeAllowedRecipient        InvestigateBulkMessageListResponseMessagePropertiesAllowlistedPatternType = "allowed_recipient"
	InvestigateBulkMessageListResponseMessagePropertiesAllowlistedPatternTypeDomainSimilarity        InvestigateBulkMessageListResponseMessagePropertiesAllowlistedPatternType = "domain_similarity"
	InvestigateBulkMessageListResponseMessagePropertiesAllowlistedPatternTypeDomainRecency           InvestigateBulkMessageListResponseMessagePropertiesAllowlistedPatternType = "domain_recency"
	InvestigateBulkMessageListResponseMessagePropertiesAllowlistedPatternTypeManagedAcceptableSender InvestigateBulkMessageListResponseMessagePropertiesAllowlistedPatternType = "managed_acceptable_sender"
	InvestigateBulkMessageListResponseMessagePropertiesAllowlistedPatternTypeOutboundNdr             InvestigateBulkMessageListResponseMessagePropertiesAllowlistedPatternType = "outbound_ndr"
)

func (r InvestigateBulkMessageListResponseMessagePropertiesAllowlistedPatternType) IsKnown() bool {
	switch r {
	case InvestigateBulkMessageListResponseMessagePropertiesAllowlistedPatternTypeQuarantineRelease, InvestigateBulkMessageListResponseMessagePropertiesAllowlistedPatternTypeAcceptableSender, InvestigateBulkMessageListResponseMessagePropertiesAllowlistedPatternTypeAllowedSender, InvestigateBulkMessageListResponseMessagePropertiesAllowlistedPatternTypeAllowedRecipient, InvestigateBulkMessageListResponseMessagePropertiesAllowlistedPatternTypeDomainSimilarity, InvestigateBulkMessageListResponseMessagePropertiesAllowlistedPatternTypeDomainRecency, InvestigateBulkMessageListResponseMessagePropertiesAllowlistedPatternTypeManagedAcceptableSender, InvestigateBulkMessageListResponseMessagePropertiesAllowlistedPatternTypeOutboundNdr:
		return true
	}
	return false
}

// Legacy field for allowlist pattern type.
type InvestigateBulkMessageListResponseMessagePropertiesWhitelistedPatternType string

const (
	InvestigateBulkMessageListResponseMessagePropertiesWhitelistedPatternTypeQuarantineRelease       InvestigateBulkMessageListResponseMessagePropertiesWhitelistedPatternType = "quarantine_release"
	InvestigateBulkMessageListResponseMessagePropertiesWhitelistedPatternTypeAcceptableSender        InvestigateBulkMessageListResponseMessagePropertiesWhitelistedPatternType = "acceptable_sender"
	InvestigateBulkMessageListResponseMessagePropertiesWhitelistedPatternTypeAllowedSender           InvestigateBulkMessageListResponseMessagePropertiesWhitelistedPatternType = "allowed_sender"
	InvestigateBulkMessageListResponseMessagePropertiesWhitelistedPatternTypeAllowedRecipient        InvestigateBulkMessageListResponseMessagePropertiesWhitelistedPatternType = "allowed_recipient"
	InvestigateBulkMessageListResponseMessagePropertiesWhitelistedPatternTypeDomainSimilarity        InvestigateBulkMessageListResponseMessagePropertiesWhitelistedPatternType = "domain_similarity"
	InvestigateBulkMessageListResponseMessagePropertiesWhitelistedPatternTypeDomainRecency           InvestigateBulkMessageListResponseMessagePropertiesWhitelistedPatternType = "domain_recency"
	InvestigateBulkMessageListResponseMessagePropertiesWhitelistedPatternTypeManagedAcceptableSender InvestigateBulkMessageListResponseMessagePropertiesWhitelistedPatternType = "managed_acceptable_sender"
	InvestigateBulkMessageListResponseMessagePropertiesWhitelistedPatternTypeOutboundNdr             InvestigateBulkMessageListResponseMessagePropertiesWhitelistedPatternType = "outbound_ndr"
)

func (r InvestigateBulkMessageListResponseMessagePropertiesWhitelistedPatternType) IsKnown() bool {
	switch r {
	case InvestigateBulkMessageListResponseMessagePropertiesWhitelistedPatternTypeQuarantineRelease, InvestigateBulkMessageListResponseMessagePropertiesWhitelistedPatternTypeAcceptableSender, InvestigateBulkMessageListResponseMessagePropertiesWhitelistedPatternTypeAllowedSender, InvestigateBulkMessageListResponseMessagePropertiesWhitelistedPatternTypeAllowedRecipient, InvestigateBulkMessageListResponseMessagePropertiesWhitelistedPatternTypeDomainSimilarity, InvestigateBulkMessageListResponseMessagePropertiesWhitelistedPatternTypeDomainRecency, InvestigateBulkMessageListResponseMessagePropertiesWhitelistedPatternTypeManagedAcceptableSender, InvestigateBulkMessageListResponseMessagePropertiesWhitelistedPatternTypeOutboundNdr:
		return true
	}
	return false
}

type InvestigateBulkMessageListResponseMessageDeliveryMode string

const (
	InvestigateBulkMessageListResponseMessageDeliveryModeDirect                InvestigateBulkMessageListResponseMessageDeliveryMode = "DIRECT"
	InvestigateBulkMessageListResponseMessageDeliveryModeBcc                   InvestigateBulkMessageListResponseMessageDeliveryMode = "BCC"
	InvestigateBulkMessageListResponseMessageDeliveryModeJournal               InvestigateBulkMessageListResponseMessageDeliveryMode = "JOURNAL"
	InvestigateBulkMessageListResponseMessageDeliveryModeReviewSubmission      InvestigateBulkMessageListResponseMessageDeliveryMode = "REVIEW_SUBMISSION"
	InvestigateBulkMessageListResponseMessageDeliveryModeDMARCUnverified       InvestigateBulkMessageListResponseMessageDeliveryMode = "DMARC_UNVERIFIED"
	InvestigateBulkMessageListResponseMessageDeliveryModeDMARCFailureReport    InvestigateBulkMessageListResponseMessageDeliveryMode = "DMARC_FAILURE_REPORT"
	InvestigateBulkMessageListResponseMessageDeliveryModeDMARCAggregateReport  InvestigateBulkMessageListResponseMessageDeliveryMode = "DMARC_AGGREGATE_REPORT"
	InvestigateBulkMessageListResponseMessageDeliveryModeThreatIntelSubmission InvestigateBulkMessageListResponseMessageDeliveryMode = "THREAT_INTEL_SUBMISSION"
	InvestigateBulkMessageListResponseMessageDeliveryModeSimulationSubmission  InvestigateBulkMessageListResponseMessageDeliveryMode = "SIMULATION_SUBMISSION"
	InvestigateBulkMessageListResponseMessageDeliveryModeAPI                   InvestigateBulkMessageListResponseMessageDeliveryMode = "API"
	InvestigateBulkMessageListResponseMessageDeliveryModeRetroScan             InvestigateBulkMessageListResponseMessageDeliveryMode = "RETRO_SCAN"
)

func (r InvestigateBulkMessageListResponseMessageDeliveryMode) IsKnown() bool {
	switch r {
	case InvestigateBulkMessageListResponseMessageDeliveryModeDirect, InvestigateBulkMessageListResponseMessageDeliveryModeBcc, InvestigateBulkMessageListResponseMessageDeliveryModeJournal, InvestigateBulkMessageListResponseMessageDeliveryModeReviewSubmission, InvestigateBulkMessageListResponseMessageDeliveryModeDMARCUnverified, InvestigateBulkMessageListResponseMessageDeliveryModeDMARCFailureReport, InvestigateBulkMessageListResponseMessageDeliveryModeDMARCAggregateReport, InvestigateBulkMessageListResponseMessageDeliveryModeThreatIntelSubmission, InvestigateBulkMessageListResponseMessageDeliveryModeSimulationSubmission, InvestigateBulkMessageListResponseMessageDeliveryModeAPI, InvestigateBulkMessageListResponseMessageDeliveryModeRetroScan:
		return true
	}
	return false
}

// Delivery status of the message.
type InvestigateBulkMessageListResponseMessageDeliveryStatus string

const (
	InvestigateBulkMessageListResponseMessageDeliveryStatusDelivered   InvestigateBulkMessageListResponseMessageDeliveryStatus = "delivered"
	InvestigateBulkMessageListResponseMessageDeliveryStatusMoved       InvestigateBulkMessageListResponseMessageDeliveryStatus = "moved"
	InvestigateBulkMessageListResponseMessageDeliveryStatusQuarantined InvestigateBulkMessageListResponseMessageDeliveryStatus = "quarantined"
	InvestigateBulkMessageListResponseMessageDeliveryStatusRejected    InvestigateBulkMessageListResponseMessageDeliveryStatus = "rejected"
	InvestigateBulkMessageListResponseMessageDeliveryStatusDeferred    InvestigateBulkMessageListResponseMessageDeliveryStatus = "deferred"
	InvestigateBulkMessageListResponseMessageDeliveryStatusBounced     InvestigateBulkMessageListResponseMessageDeliveryStatus = "bounced"
	InvestigateBulkMessageListResponseMessageDeliveryStatusQueued      InvestigateBulkMessageListResponseMessageDeliveryStatus = "queued"
	InvestigateBulkMessageListResponseMessageDeliveryStatusMoveFailed  InvestigateBulkMessageListResponseMessageDeliveryStatus = "move_failed"
)

func (r InvestigateBulkMessageListResponseMessageDeliveryStatus) IsKnown() bool {
	switch r {
	case InvestigateBulkMessageListResponseMessageDeliveryStatusDelivered, InvestigateBulkMessageListResponseMessageDeliveryStatusMoved, InvestigateBulkMessageListResponseMessageDeliveryStatusQuarantined, InvestigateBulkMessageListResponseMessageDeliveryStatusRejected, InvestigateBulkMessageListResponseMessageDeliveryStatusDeferred, InvestigateBulkMessageListResponseMessageDeliveryStatusBounced, InvestigateBulkMessageListResponseMessageDeliveryStatusQueued, InvestigateBulkMessageListResponseMessageDeliveryStatusMoveFailed:
		return true
	}
	return false
}

type InvestigateBulkMessageListResponseMessageFinalDisposition string

const (
	InvestigateBulkMessageListResponseMessageFinalDispositionMalicious    InvestigateBulkMessageListResponseMessageFinalDisposition = "MALICIOUS"
	InvestigateBulkMessageListResponseMessageFinalDispositionMaliciousBec InvestigateBulkMessageListResponseMessageFinalDisposition = "MALICIOUS-BEC"
	InvestigateBulkMessageListResponseMessageFinalDispositionSuspicious   InvestigateBulkMessageListResponseMessageFinalDisposition = "SUSPICIOUS"
	InvestigateBulkMessageListResponseMessageFinalDispositionSpoof        InvestigateBulkMessageListResponseMessageFinalDisposition = "SPOOF"
	InvestigateBulkMessageListResponseMessageFinalDispositionSpam         InvestigateBulkMessageListResponseMessageFinalDisposition = "SPAM"
	InvestigateBulkMessageListResponseMessageFinalDispositionBulk         InvestigateBulkMessageListResponseMessageFinalDisposition = "BULK"
	InvestigateBulkMessageListResponseMessageFinalDispositionEncrypted    InvestigateBulkMessageListResponseMessageFinalDisposition = "ENCRYPTED"
	InvestigateBulkMessageListResponseMessageFinalDispositionExternal     InvestigateBulkMessageListResponseMessageFinalDisposition = "EXTERNAL"
	InvestigateBulkMessageListResponseMessageFinalDispositionUnknown      InvestigateBulkMessageListResponseMessageFinalDisposition = "UNKNOWN"
	InvestigateBulkMessageListResponseMessageFinalDispositionNone         InvestigateBulkMessageListResponseMessageFinalDisposition = "NONE"
)

func (r InvestigateBulkMessageListResponseMessageFinalDisposition) IsKnown() bool {
	switch r {
	case InvestigateBulkMessageListResponseMessageFinalDispositionMalicious, InvestigateBulkMessageListResponseMessageFinalDispositionMaliciousBec, InvestigateBulkMessageListResponseMessageFinalDispositionSuspicious, InvestigateBulkMessageListResponseMessageFinalDispositionSpoof, InvestigateBulkMessageListResponseMessageFinalDispositionSpam, InvestigateBulkMessageListResponseMessageFinalDispositionBulk, InvestigateBulkMessageListResponseMessageFinalDispositionEncrypted, InvestigateBulkMessageListResponseMessageFinalDispositionExternal, InvestigateBulkMessageListResponseMessageFinalDispositionUnknown, InvestigateBulkMessageListResponseMessageFinalDispositionNone:
		return true
	}
	return false
}

type InvestigateBulkMessageListResponseMessageFinding struct {
	Attachment string                                                     `json:"attachment" api:"nullable"`
	Detail     string                                                     `json:"detail" api:"nullable"`
	Detection  InvestigateBulkMessageListResponseMessageFindingsDetection `json:"detection"`
	Field      string                                                     `json:"field" api:"nullable"`
	Name       string                                                     `json:"name" api:"nullable"`
	Portion    string                                                     `json:"portion" api:"nullable"`
	Reason     string                                                     `json:"reason" api:"nullable"`
	Score      float64                                                    `json:"score" api:"nullable"`
	Value      string                                                     `json:"value" api:"nullable"`
	JSON       investigateBulkMessageListResponseMessageFindingJSON       `json:"-"`
}

// investigateBulkMessageListResponseMessageFindingJSON contains the JSON metadata
// for the struct [InvestigateBulkMessageListResponseMessageFinding]
type investigateBulkMessageListResponseMessageFindingJSON struct {
	Attachment  apijson.Field
	Detail      apijson.Field
	Detection   apijson.Field
	Field       apijson.Field
	Name        apijson.Field
	Portion     apijson.Field
	Reason      apijson.Field
	Score       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateBulkMessageListResponseMessageFinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkMessageListResponseMessageFindingJSON) RawJSON() string {
	return r.raw
}

type InvestigateBulkMessageListResponseMessageFindingsDetection string

const (
	InvestigateBulkMessageListResponseMessageFindingsDetectionMalicious    InvestigateBulkMessageListResponseMessageFindingsDetection = "MALICIOUS"
	InvestigateBulkMessageListResponseMessageFindingsDetectionMaliciousBec InvestigateBulkMessageListResponseMessageFindingsDetection = "MALICIOUS-BEC"
	InvestigateBulkMessageListResponseMessageFindingsDetectionSuspicious   InvestigateBulkMessageListResponseMessageFindingsDetection = "SUSPICIOUS"
	InvestigateBulkMessageListResponseMessageFindingsDetectionSpoof        InvestigateBulkMessageListResponseMessageFindingsDetection = "SPOOF"
	InvestigateBulkMessageListResponseMessageFindingsDetectionSpam         InvestigateBulkMessageListResponseMessageFindingsDetection = "SPAM"
	InvestigateBulkMessageListResponseMessageFindingsDetectionBulk         InvestigateBulkMessageListResponseMessageFindingsDetection = "BULK"
	InvestigateBulkMessageListResponseMessageFindingsDetectionEncrypted    InvestigateBulkMessageListResponseMessageFindingsDetection = "ENCRYPTED"
	InvestigateBulkMessageListResponseMessageFindingsDetectionExternal     InvestigateBulkMessageListResponseMessageFindingsDetection = "EXTERNAL"
	InvestigateBulkMessageListResponseMessageFindingsDetectionUnknown      InvestigateBulkMessageListResponseMessageFindingsDetection = "UNKNOWN"
	InvestigateBulkMessageListResponseMessageFindingsDetectionNone         InvestigateBulkMessageListResponseMessageFindingsDetection = "NONE"
)

func (r InvestigateBulkMessageListResponseMessageFindingsDetection) IsKnown() bool {
	switch r {
	case InvestigateBulkMessageListResponseMessageFindingsDetectionMalicious, InvestigateBulkMessageListResponseMessageFindingsDetectionMaliciousBec, InvestigateBulkMessageListResponseMessageFindingsDetectionSuspicious, InvestigateBulkMessageListResponseMessageFindingsDetectionSpoof, InvestigateBulkMessageListResponseMessageFindingsDetectionSpam, InvestigateBulkMessageListResponseMessageFindingsDetectionBulk, InvestigateBulkMessageListResponseMessageFindingsDetectionEncrypted, InvestigateBulkMessageListResponseMessageFindingsDetectionExternal, InvestigateBulkMessageListResponseMessageFindingsDetectionUnknown, InvestigateBulkMessageListResponseMessageFindingsDetectionNone:
		return true
	}
	return false
}

type InvestigateBulkMessageListResponseMessagePostDeliveryOperation string

const (
	InvestigateBulkMessageListResponseMessagePostDeliveryOperationPreview           InvestigateBulkMessageListResponseMessagePostDeliveryOperation = "PREVIEW"
	InvestigateBulkMessageListResponseMessagePostDeliveryOperationQuarantineRelease InvestigateBulkMessageListResponseMessagePostDeliveryOperation = "QUARANTINE_RELEASE"
	InvestigateBulkMessageListResponseMessagePostDeliveryOperationSubmission        InvestigateBulkMessageListResponseMessagePostDeliveryOperation = "SUBMISSION"
	InvestigateBulkMessageListResponseMessagePostDeliveryOperationMove              InvestigateBulkMessageListResponseMessagePostDeliveryOperation = "MOVE"
)

func (r InvestigateBulkMessageListResponseMessagePostDeliveryOperation) IsKnown() bool {
	switch r {
	case InvestigateBulkMessageListResponseMessagePostDeliveryOperationPreview, InvestigateBulkMessageListResponseMessagePostDeliveryOperationQuarantineRelease, InvestigateBulkMessageListResponseMessagePostDeliveryOperationSubmission, InvestigateBulkMessageListResponseMessagePostDeliveryOperationMove:
		return true
	}
	return false
}

type InvestigateBulkMessageListResponseMessageValidation struct {
	Comment string                                                   `json:"comment" api:"nullable"`
	DKIM    InvestigateBulkMessageListResponseMessageValidationDKIM  `json:"dkim"`
	DMARC   InvestigateBulkMessageListResponseMessageValidationDMARC `json:"dmarc"`
	SPF     InvestigateBulkMessageListResponseMessageValidationSPF   `json:"spf"`
	JSON    investigateBulkMessageListResponseMessageValidationJSON  `json:"-"`
}

// investigateBulkMessageListResponseMessageValidationJSON contains the JSON
// metadata for the struct [InvestigateBulkMessageListResponseMessageValidation]
type investigateBulkMessageListResponseMessageValidationJSON struct {
	Comment     apijson.Field
	DKIM        apijson.Field
	DMARC       apijson.Field
	SPF         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateBulkMessageListResponseMessageValidation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkMessageListResponseMessageValidationJSON) RawJSON() string {
	return r.raw
}

type InvestigateBulkMessageListResponseMessageValidationDKIM string

const (
	InvestigateBulkMessageListResponseMessageValidationDKIMPass    InvestigateBulkMessageListResponseMessageValidationDKIM = "pass"
	InvestigateBulkMessageListResponseMessageValidationDKIMNeutral InvestigateBulkMessageListResponseMessageValidationDKIM = "neutral"
	InvestigateBulkMessageListResponseMessageValidationDKIMFail    InvestigateBulkMessageListResponseMessageValidationDKIM = "fail"
	InvestigateBulkMessageListResponseMessageValidationDKIMError   InvestigateBulkMessageListResponseMessageValidationDKIM = "error"
	InvestigateBulkMessageListResponseMessageValidationDKIMNone    InvestigateBulkMessageListResponseMessageValidationDKIM = "none"
)

func (r InvestigateBulkMessageListResponseMessageValidationDKIM) IsKnown() bool {
	switch r {
	case InvestigateBulkMessageListResponseMessageValidationDKIMPass, InvestigateBulkMessageListResponseMessageValidationDKIMNeutral, InvestigateBulkMessageListResponseMessageValidationDKIMFail, InvestigateBulkMessageListResponseMessageValidationDKIMError, InvestigateBulkMessageListResponseMessageValidationDKIMNone:
		return true
	}
	return false
}

type InvestigateBulkMessageListResponseMessageValidationDMARC string

const (
	InvestigateBulkMessageListResponseMessageValidationDMARCPass    InvestigateBulkMessageListResponseMessageValidationDMARC = "pass"
	InvestigateBulkMessageListResponseMessageValidationDMARCNeutral InvestigateBulkMessageListResponseMessageValidationDMARC = "neutral"
	InvestigateBulkMessageListResponseMessageValidationDMARCFail    InvestigateBulkMessageListResponseMessageValidationDMARC = "fail"
	InvestigateBulkMessageListResponseMessageValidationDMARCError   InvestigateBulkMessageListResponseMessageValidationDMARC = "error"
	InvestigateBulkMessageListResponseMessageValidationDMARCNone    InvestigateBulkMessageListResponseMessageValidationDMARC = "none"
)

func (r InvestigateBulkMessageListResponseMessageValidationDMARC) IsKnown() bool {
	switch r {
	case InvestigateBulkMessageListResponseMessageValidationDMARCPass, InvestigateBulkMessageListResponseMessageValidationDMARCNeutral, InvestigateBulkMessageListResponseMessageValidationDMARCFail, InvestigateBulkMessageListResponseMessageValidationDMARCError, InvestigateBulkMessageListResponseMessageValidationDMARCNone:
		return true
	}
	return false
}

type InvestigateBulkMessageListResponseMessageValidationSPF string

const (
	InvestigateBulkMessageListResponseMessageValidationSPFPass    InvestigateBulkMessageListResponseMessageValidationSPF = "pass"
	InvestigateBulkMessageListResponseMessageValidationSPFNeutral InvestigateBulkMessageListResponseMessageValidationSPF = "neutral"
	InvestigateBulkMessageListResponseMessageValidationSPFFail    InvestigateBulkMessageListResponseMessageValidationSPF = "fail"
	InvestigateBulkMessageListResponseMessageValidationSPFError   InvestigateBulkMessageListResponseMessageValidationSPF = "error"
	InvestigateBulkMessageListResponseMessageValidationSPFNone    InvestigateBulkMessageListResponseMessageValidationSPF = "none"
)

func (r InvestigateBulkMessageListResponseMessageValidationSPF) IsKnown() bool {
	switch r {
	case InvestigateBulkMessageListResponseMessageValidationSPFPass, InvestigateBulkMessageListResponseMessageValidationSPFNeutral, InvestigateBulkMessageListResponseMessageValidationSPFFail, InvestigateBulkMessageListResponseMessageValidationSPFError, InvestigateBulkMessageListResponseMessageValidationSPFNone:
		return true
	}
	return false
}

type InvestigateBulkMessageListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Current page within paginated list of results.
	Page param.Field[int64] `query:"page"`
	// The number of results per page. Maximum value is 1000.
	PerPage param.Field[int64]                                  `query:"per_page"`
	Status  param.Field[InvestigateBulkMessageListParamsStatus] `query:"status"`
}

// URLQuery serializes [InvestigateBulkMessageListParams]'s query parameters as
// `url.Values`.
func (r InvestigateBulkMessageListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InvestigateBulkMessageListParamsStatus string

const (
	InvestigateBulkMessageListParamsStatusPending     InvestigateBulkMessageListParamsStatus = "PENDING"
	InvestigateBulkMessageListParamsStatusDiscovering InvestigateBulkMessageListParamsStatus = "DISCOVERING"
	InvestigateBulkMessageListParamsStatusProcessing  InvestigateBulkMessageListParamsStatus = "PROCESSING"
	InvestigateBulkMessageListParamsStatusCompleted   InvestigateBulkMessageListParamsStatus = "COMPLETED"
	InvestigateBulkMessageListParamsStatusFailed      InvestigateBulkMessageListParamsStatus = "FAILED"
	InvestigateBulkMessageListParamsStatusCancelled   InvestigateBulkMessageListParamsStatus = "CANCELLED"
	InvestigateBulkMessageListParamsStatusSkipped     InvestigateBulkMessageListParamsStatus = "SKIPPED"
)

func (r InvestigateBulkMessageListParamsStatus) IsKnown() bool {
	switch r {
	case InvestigateBulkMessageListParamsStatusPending, InvestigateBulkMessageListParamsStatusDiscovering, InvestigateBulkMessageListParamsStatusProcessing, InvestigateBulkMessageListParamsStatusCompleted, InvestigateBulkMessageListParamsStatusFailed, InvestigateBulkMessageListParamsStatusCancelled, InvestigateBulkMessageListParamsStatusSkipped:
		return true
	}
	return false
}
