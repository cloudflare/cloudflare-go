// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// InvestigateService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvestigateService] method instead.
type InvestigateService struct {
	Options    []option.RequestOption
	Detections *InvestigateDetectionService
	Preview    *InvestigatePreviewService
	Raw        *InvestigateRawService
	Trace      *InvestigateTraceService
	Move       *InvestigateMoveService
	Reclassify *InvestigateReclassifyService
	Release    *InvestigateReleaseService
}

// NewInvestigateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInvestigateService(opts ...option.RequestOption) (r *InvestigateService) {
	r = &InvestigateService{}
	r.Options = opts
	r.Detections = NewInvestigateDetectionService(opts...)
	r.Preview = NewInvestigatePreviewService(opts...)
	r.Raw = NewInvestigateRawService(opts...)
	r.Trace = NewInvestigateTraceService(opts...)
	r.Move = NewInvestigateMoveService(opts...)
	r.Reclassify = NewInvestigateReclassifyService(opts...)
	r.Release = NewInvestigateReleaseService(opts...)
	return
}

// Returns information for each email that matches the search parameter(s).
func (r *InvestigateService) List(ctx context.Context, params InvestigateListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[InvestigateListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate", params.AccountID)
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

// Returns information for each email that matches the search parameter(s).
func (r *InvestigateService) ListAutoPaging(ctx context.Context, params InvestigateListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[InvestigateListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Retrieves comprehensive details for a specific email message including headers,
// recipients, sender information, and current quarantine status. Use the
// investigate_id from search results to fetch detailed information.
func (r *InvestigateService) Get(ctx context.Context, investigateID string, params InvestigateGetParams, opts ...option.RequestOption) (res *InvestigateGetResponse, err error) {
	var env InvestigateGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if investigateID == "" {
		err = errors.New("missing required investigate_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/%s", params.AccountID, investigateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type InvestigateListResponse struct {
	// Unique identifier for a message retrieved from investigation
	ID string `json:"id" api:"required"`
	// Deprecated, use `GET /investigate/{investigate_id}/action_log` instead. End of
	// life: November 1, 2026.
	//
	// Deprecated: deprecated
	ActionLog         []InvestigateListResponseActionLog `json:"action_log" api:"required"`
	ClientRecipients  []string                           `json:"client_recipients" api:"required"`
	DetectionReasons  []string                           `json:"detection_reasons" api:"required"`
	IsPhishSubmission bool                               `json:"is_phish_submission" api:"required"`
	IsQuarantined     bool                               `json:"is_quarantined" api:"required"`
	// The identifier of the message
	PostfixID string `json:"postfix_id" api:"required"`
	// Message processing properties
	Properties InvestigateListResponseProperties `json:"properties" api:"required"`
	// Deprecated, use `scanned_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: deprecated
	Ts               string                                  `json:"ts" api:"required"`
	AlertID          string                                  `json:"alert_id" api:"nullable"`
	DeliveryMode     InvestigateListResponseDeliveryMode     `json:"delivery_mode"`
	DeliveryStatus   []InvestigateListResponseDeliveryStatus `json:"delivery_status" api:"nullable"`
	EdfHash          string                                  `json:"edf_hash" api:"nullable"`
	EnvelopeFrom     string                                  `json:"envelope_from" api:"nullable"`
	EnvelopeTo       []string                                `json:"envelope_to" api:"nullable"`
	FinalDisposition InvestigateListResponseFinalDisposition `json:"final_disposition"`
	// Deprecated, use the `findings` field from
	// `GET /investigate/{investigate_id}/detections` instead. End of life: November
	// 1, 2026. Detection findings for this message.
	//
	// Deprecated: deprecated
	Findings              []InvestigateListResponseFinding `json:"findings" api:"nullable"`
	From                  string                           `json:"from" api:"nullable"`
	FromName              string                           `json:"from_name" api:"nullable"`
	HtmltextStructureHash string                           `json:"htmltext_structure_hash" api:"nullable"`
	MessageID             string                           `json:"message_id" api:"nullable"`
	// Post-delivery operations performed on this message
	PostDeliveryOperations []InvestigateListResponsePostDeliveryOperation `json:"post_delivery_operations" api:"nullable"`
	PostfixIDOutbound      string                                         `json:"postfix_id_outbound" api:"nullable"`
	Replyto                string                                         `json:"replyto" api:"nullable"`
	// When the message was scanned (UTC)
	ScannedAt time.Time `json:"scanned_at" api:"nullable" format:"date-time"`
	// When the message was sent (UTC)
	SentAt           time.Time                         `json:"sent_at" api:"nullable" format:"date-time"`
	SentDate         string                            `json:"sent_date" api:"nullable"`
	Subject          string                            `json:"subject" api:"nullable"`
	ThreatCategories []string                          `json:"threat_categories" api:"nullable"`
	To               []string                          `json:"to" api:"nullable"`
	ToName           []string                          `json:"to_name" api:"nullable"`
	Validation       InvestigateListResponseValidation `json:"validation"`
	JSON             investigateListResponseJSON       `json:"-"`
}

// investigateListResponseJSON contains the JSON metadata for the struct
// [InvestigateListResponse]
type investigateListResponseJSON struct {
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
	Subject                apijson.Field
	ThreatCategories       apijson.Field
	To                     apijson.Field
	ToName                 apijson.Field
	Validation             apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *InvestigateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateListResponseJSON) RawJSON() string {
	return r.raw
}

type InvestigateListResponseActionLog struct {
	// Timestamp when action completed
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	// Type of action performed
	Operation InvestigateListResponseActionLogOperation `json:"operation" api:"required"`
	// Deprecated, use `completed_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: deprecated
	CompletedTimestamp string `json:"completed_timestamp"`
	// Additional properties for the action
	Properties InvestigateListResponseActionLogProperties `json:"properties"`
	// Status of the action
	Status string                               `json:"status" api:"nullable"`
	JSON   investigateListResponseActionLogJSON `json:"-"`
}

// investigateListResponseActionLogJSON contains the JSON metadata for the struct
// [InvestigateListResponseActionLog]
type investigateListResponseActionLogJSON struct {
	CompletedAt        apijson.Field
	Operation          apijson.Field
	CompletedTimestamp apijson.Field
	Properties         apijson.Field
	Status             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InvestigateListResponseActionLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateListResponseActionLogJSON) RawJSON() string {
	return r.raw
}

// Type of action performed
type InvestigateListResponseActionLogOperation string

const (
	InvestigateListResponseActionLogOperationMove              InvestigateListResponseActionLogOperation = "MOVE"
	InvestigateListResponseActionLogOperationRelease           InvestigateListResponseActionLogOperation = "RELEASE"
	InvestigateListResponseActionLogOperationReclassify        InvestigateListResponseActionLogOperation = "RECLASSIFY"
	InvestigateListResponseActionLogOperationSubmission        InvestigateListResponseActionLogOperation = "SUBMISSION"
	InvestigateListResponseActionLogOperationQuarantineRelease InvestigateListResponseActionLogOperation = "QUARANTINE_RELEASE"
	InvestigateListResponseActionLogOperationPreview           InvestigateListResponseActionLogOperation = "PREVIEW"
)

func (r InvestigateListResponseActionLogOperation) IsKnown() bool {
	switch r {
	case InvestigateListResponseActionLogOperationMove, InvestigateListResponseActionLogOperationRelease, InvestigateListResponseActionLogOperationReclassify, InvestigateListResponseActionLogOperationSubmission, InvestigateListResponseActionLogOperationQuarantineRelease, InvestigateListResponseActionLogOperationPreview:
		return true
	}
	return false
}

// Additional properties for the action
type InvestigateListResponseActionLogProperties struct {
	// Target folder for move operations
	Folder string `json:"folder"`
	// User who requested the action
	RequestedBy string                                         `json:"requested_by"`
	JSON        investigateListResponseActionLogPropertiesJSON `json:"-"`
}

// investigateListResponseActionLogPropertiesJSON contains the JSON metadata for
// the struct [InvestigateListResponseActionLogProperties]
type investigateListResponseActionLogPropertiesJSON struct {
	Folder      apijson.Field
	RequestedBy apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateListResponseActionLogProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateListResponseActionLogPropertiesJSON) RawJSON() string {
	return r.raw
}

// Message processing properties
type InvestigateListResponseProperties struct {
	// Pattern that allowlisted this message
	AllowlistedPattern string `json:"allowlisted_pattern" api:"nullable"`
	// Type of allowlist pattern
	AllowlistedPatternType InvestigateListResponsePropertiesAllowlistedPatternType `json:"allowlisted_pattern_type" api:"nullable"`
	// Whether message was blocklisted
	BlocklistedMessage bool `json:"blocklisted_message" api:"nullable"`
	// Pattern that blocklisted this message
	BlocklistedPattern string `json:"blocklisted_pattern" api:"nullable"`
	// Legacy field for allowlist pattern type
	WhitelistedPatternType InvestigateListResponsePropertiesWhitelistedPatternType `json:"whitelisted_pattern_type" api:"nullable"`
	JSON                   investigateListResponsePropertiesJSON                   `json:"-"`
}

// investigateListResponsePropertiesJSON contains the JSON metadata for the struct
// [InvestigateListResponseProperties]
type investigateListResponsePropertiesJSON struct {
	AllowlistedPattern     apijson.Field
	AllowlistedPatternType apijson.Field
	BlocklistedMessage     apijson.Field
	BlocklistedPattern     apijson.Field
	WhitelistedPatternType apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *InvestigateListResponseProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateListResponsePropertiesJSON) RawJSON() string {
	return r.raw
}

// Type of allowlist pattern
type InvestigateListResponsePropertiesAllowlistedPatternType string

const (
	InvestigateListResponsePropertiesAllowlistedPatternTypeQuarantineRelease       InvestigateListResponsePropertiesAllowlistedPatternType = "quarantine_release"
	InvestigateListResponsePropertiesAllowlistedPatternTypeAcceptableSender        InvestigateListResponsePropertiesAllowlistedPatternType = "acceptable_sender"
	InvestigateListResponsePropertiesAllowlistedPatternTypeAllowedSender           InvestigateListResponsePropertiesAllowlistedPatternType = "allowed_sender"
	InvestigateListResponsePropertiesAllowlistedPatternTypeAllowedRecipient        InvestigateListResponsePropertiesAllowlistedPatternType = "allowed_recipient"
	InvestigateListResponsePropertiesAllowlistedPatternTypeDomainSimilarity        InvestigateListResponsePropertiesAllowlistedPatternType = "domain_similarity"
	InvestigateListResponsePropertiesAllowlistedPatternTypeDomainRecency           InvestigateListResponsePropertiesAllowlistedPatternType = "domain_recency"
	InvestigateListResponsePropertiesAllowlistedPatternTypeManagedAcceptableSender InvestigateListResponsePropertiesAllowlistedPatternType = "managed_acceptable_sender"
	InvestigateListResponsePropertiesAllowlistedPatternTypeOutboundNdr             InvestigateListResponsePropertiesAllowlistedPatternType = "outbound_ndr"
)

func (r InvestigateListResponsePropertiesAllowlistedPatternType) IsKnown() bool {
	switch r {
	case InvestigateListResponsePropertiesAllowlistedPatternTypeQuarantineRelease, InvestigateListResponsePropertiesAllowlistedPatternTypeAcceptableSender, InvestigateListResponsePropertiesAllowlistedPatternTypeAllowedSender, InvestigateListResponsePropertiesAllowlistedPatternTypeAllowedRecipient, InvestigateListResponsePropertiesAllowlistedPatternTypeDomainSimilarity, InvestigateListResponsePropertiesAllowlistedPatternTypeDomainRecency, InvestigateListResponsePropertiesAllowlistedPatternTypeManagedAcceptableSender, InvestigateListResponsePropertiesAllowlistedPatternTypeOutboundNdr:
		return true
	}
	return false
}

// Legacy field for allowlist pattern type
type InvestigateListResponsePropertiesWhitelistedPatternType string

const (
	InvestigateListResponsePropertiesWhitelistedPatternTypeQuarantineRelease       InvestigateListResponsePropertiesWhitelistedPatternType = "quarantine_release"
	InvestigateListResponsePropertiesWhitelistedPatternTypeAcceptableSender        InvestigateListResponsePropertiesWhitelistedPatternType = "acceptable_sender"
	InvestigateListResponsePropertiesWhitelistedPatternTypeAllowedSender           InvestigateListResponsePropertiesWhitelistedPatternType = "allowed_sender"
	InvestigateListResponsePropertiesWhitelistedPatternTypeAllowedRecipient        InvestigateListResponsePropertiesWhitelistedPatternType = "allowed_recipient"
	InvestigateListResponsePropertiesWhitelistedPatternTypeDomainSimilarity        InvestigateListResponsePropertiesWhitelistedPatternType = "domain_similarity"
	InvestigateListResponsePropertiesWhitelistedPatternTypeDomainRecency           InvestigateListResponsePropertiesWhitelistedPatternType = "domain_recency"
	InvestigateListResponsePropertiesWhitelistedPatternTypeManagedAcceptableSender InvestigateListResponsePropertiesWhitelistedPatternType = "managed_acceptable_sender"
	InvestigateListResponsePropertiesWhitelistedPatternTypeOutboundNdr             InvestigateListResponsePropertiesWhitelistedPatternType = "outbound_ndr"
)

func (r InvestigateListResponsePropertiesWhitelistedPatternType) IsKnown() bool {
	switch r {
	case InvestigateListResponsePropertiesWhitelistedPatternTypeQuarantineRelease, InvestigateListResponsePropertiesWhitelistedPatternTypeAcceptableSender, InvestigateListResponsePropertiesWhitelistedPatternTypeAllowedSender, InvestigateListResponsePropertiesWhitelistedPatternTypeAllowedRecipient, InvestigateListResponsePropertiesWhitelistedPatternTypeDomainSimilarity, InvestigateListResponsePropertiesWhitelistedPatternTypeDomainRecency, InvestigateListResponsePropertiesWhitelistedPatternTypeManagedAcceptableSender, InvestigateListResponsePropertiesWhitelistedPatternTypeOutboundNdr:
		return true
	}
	return false
}

type InvestigateListResponseDeliveryMode string

const (
	InvestigateListResponseDeliveryModeDirect                InvestigateListResponseDeliveryMode = "DIRECT"
	InvestigateListResponseDeliveryModeBcc                   InvestigateListResponseDeliveryMode = "BCC"
	InvestigateListResponseDeliveryModeJournal               InvestigateListResponseDeliveryMode = "JOURNAL"
	InvestigateListResponseDeliveryModeReviewSubmission      InvestigateListResponseDeliveryMode = "REVIEW_SUBMISSION"
	InvestigateListResponseDeliveryModeDMARCUnverified       InvestigateListResponseDeliveryMode = "DMARC_UNVERIFIED"
	InvestigateListResponseDeliveryModeDMARCFailureReport    InvestigateListResponseDeliveryMode = "DMARC_FAILURE_REPORT"
	InvestigateListResponseDeliveryModeDMARCAggregateReport  InvestigateListResponseDeliveryMode = "DMARC_AGGREGATE_REPORT"
	InvestigateListResponseDeliveryModeThreatIntelSubmission InvestigateListResponseDeliveryMode = "THREAT_INTEL_SUBMISSION"
	InvestigateListResponseDeliveryModeSimulationSubmission  InvestigateListResponseDeliveryMode = "SIMULATION_SUBMISSION"
	InvestigateListResponseDeliveryModeAPI                   InvestigateListResponseDeliveryMode = "API"
	InvestigateListResponseDeliveryModeRetroScan             InvestigateListResponseDeliveryMode = "RETRO_SCAN"
)

func (r InvestigateListResponseDeliveryMode) IsKnown() bool {
	switch r {
	case InvestigateListResponseDeliveryModeDirect, InvestigateListResponseDeliveryModeBcc, InvestigateListResponseDeliveryModeJournal, InvestigateListResponseDeliveryModeReviewSubmission, InvestigateListResponseDeliveryModeDMARCUnverified, InvestigateListResponseDeliveryModeDMARCFailureReport, InvestigateListResponseDeliveryModeDMARCAggregateReport, InvestigateListResponseDeliveryModeThreatIntelSubmission, InvestigateListResponseDeliveryModeSimulationSubmission, InvestigateListResponseDeliveryModeAPI, InvestigateListResponseDeliveryModeRetroScan:
		return true
	}
	return false
}

// Delivery status of the message.
type InvestigateListResponseDeliveryStatus string

const (
	InvestigateListResponseDeliveryStatusDelivered   InvestigateListResponseDeliveryStatus = "delivered"
	InvestigateListResponseDeliveryStatusMoved       InvestigateListResponseDeliveryStatus = "moved"
	InvestigateListResponseDeliveryStatusQuarantined InvestigateListResponseDeliveryStatus = "quarantined"
	InvestigateListResponseDeliveryStatusRejected    InvestigateListResponseDeliveryStatus = "rejected"
	InvestigateListResponseDeliveryStatusDeferred    InvestigateListResponseDeliveryStatus = "deferred"
	InvestigateListResponseDeliveryStatusBounced     InvestigateListResponseDeliveryStatus = "bounced"
	InvestigateListResponseDeliveryStatusQueued      InvestigateListResponseDeliveryStatus = "queued"
)

func (r InvestigateListResponseDeliveryStatus) IsKnown() bool {
	switch r {
	case InvestigateListResponseDeliveryStatusDelivered, InvestigateListResponseDeliveryStatusMoved, InvestigateListResponseDeliveryStatusQuarantined, InvestigateListResponseDeliveryStatusRejected, InvestigateListResponseDeliveryStatusDeferred, InvestigateListResponseDeliveryStatusBounced, InvestigateListResponseDeliveryStatusQueued:
		return true
	}
	return false
}

type InvestigateListResponseFinalDisposition string

const (
	InvestigateListResponseFinalDispositionMalicious    InvestigateListResponseFinalDisposition = "MALICIOUS"
	InvestigateListResponseFinalDispositionMaliciousBec InvestigateListResponseFinalDisposition = "MALICIOUS-BEC"
	InvestigateListResponseFinalDispositionSuspicious   InvestigateListResponseFinalDisposition = "SUSPICIOUS"
	InvestigateListResponseFinalDispositionSpoof        InvestigateListResponseFinalDisposition = "SPOOF"
	InvestigateListResponseFinalDispositionSpam         InvestigateListResponseFinalDisposition = "SPAM"
	InvestigateListResponseFinalDispositionBulk         InvestigateListResponseFinalDisposition = "BULK"
	InvestigateListResponseFinalDispositionEncrypted    InvestigateListResponseFinalDisposition = "ENCRYPTED"
	InvestigateListResponseFinalDispositionExternal     InvestigateListResponseFinalDisposition = "EXTERNAL"
	InvestigateListResponseFinalDispositionUnknown      InvestigateListResponseFinalDisposition = "UNKNOWN"
	InvestigateListResponseFinalDispositionNone         InvestigateListResponseFinalDisposition = "NONE"
)

func (r InvestigateListResponseFinalDisposition) IsKnown() bool {
	switch r {
	case InvestigateListResponseFinalDispositionMalicious, InvestigateListResponseFinalDispositionMaliciousBec, InvestigateListResponseFinalDispositionSuspicious, InvestigateListResponseFinalDispositionSpoof, InvestigateListResponseFinalDispositionSpam, InvestigateListResponseFinalDispositionBulk, InvestigateListResponseFinalDispositionEncrypted, InvestigateListResponseFinalDispositionExternal, InvestigateListResponseFinalDispositionUnknown, InvestigateListResponseFinalDispositionNone:
		return true
	}
	return false
}

type InvestigateListResponseFinding struct {
	Attachment string                                   `json:"attachment" api:"nullable"`
	Detail     string                                   `json:"detail" api:"nullable"`
	Detection  InvestigateListResponseFindingsDetection `json:"detection"`
	Field      string                                   `json:"field" api:"nullable"`
	Name       string                                   `json:"name" api:"nullable"`
	Portion    string                                   `json:"portion" api:"nullable"`
	Reason     string                                   `json:"reason" api:"nullable"`
	Score      float64                                  `json:"score" api:"nullable"`
	Value      string                                   `json:"value" api:"nullable"`
	JSON       investigateListResponseFindingJSON       `json:"-"`
}

// investigateListResponseFindingJSON contains the JSON metadata for the struct
// [InvestigateListResponseFinding]
type investigateListResponseFindingJSON struct {
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

func (r *InvestigateListResponseFinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateListResponseFindingJSON) RawJSON() string {
	return r.raw
}

type InvestigateListResponseFindingsDetection string

const (
	InvestigateListResponseFindingsDetectionMalicious    InvestigateListResponseFindingsDetection = "MALICIOUS"
	InvestigateListResponseFindingsDetectionMaliciousBec InvestigateListResponseFindingsDetection = "MALICIOUS-BEC"
	InvestigateListResponseFindingsDetectionSuspicious   InvestigateListResponseFindingsDetection = "SUSPICIOUS"
	InvestigateListResponseFindingsDetectionSpoof        InvestigateListResponseFindingsDetection = "SPOOF"
	InvestigateListResponseFindingsDetectionSpam         InvestigateListResponseFindingsDetection = "SPAM"
	InvestigateListResponseFindingsDetectionBulk         InvestigateListResponseFindingsDetection = "BULK"
	InvestigateListResponseFindingsDetectionEncrypted    InvestigateListResponseFindingsDetection = "ENCRYPTED"
	InvestigateListResponseFindingsDetectionExternal     InvestigateListResponseFindingsDetection = "EXTERNAL"
	InvestigateListResponseFindingsDetectionUnknown      InvestigateListResponseFindingsDetection = "UNKNOWN"
	InvestigateListResponseFindingsDetectionNone         InvestigateListResponseFindingsDetection = "NONE"
)

func (r InvestigateListResponseFindingsDetection) IsKnown() bool {
	switch r {
	case InvestigateListResponseFindingsDetectionMalicious, InvestigateListResponseFindingsDetectionMaliciousBec, InvestigateListResponseFindingsDetectionSuspicious, InvestigateListResponseFindingsDetectionSpoof, InvestigateListResponseFindingsDetectionSpam, InvestigateListResponseFindingsDetectionBulk, InvestigateListResponseFindingsDetectionEncrypted, InvestigateListResponseFindingsDetectionExternal, InvestigateListResponseFindingsDetectionUnknown, InvestigateListResponseFindingsDetectionNone:
		return true
	}
	return false
}

type InvestigateListResponsePostDeliveryOperation string

const (
	InvestigateListResponsePostDeliveryOperationPreview           InvestigateListResponsePostDeliveryOperation = "PREVIEW"
	InvestigateListResponsePostDeliveryOperationQuarantineRelease InvestigateListResponsePostDeliveryOperation = "QUARANTINE_RELEASE"
	InvestigateListResponsePostDeliveryOperationSubmission        InvestigateListResponsePostDeliveryOperation = "SUBMISSION"
	InvestigateListResponsePostDeliveryOperationMove              InvestigateListResponsePostDeliveryOperation = "MOVE"
)

func (r InvestigateListResponsePostDeliveryOperation) IsKnown() bool {
	switch r {
	case InvestigateListResponsePostDeliveryOperationPreview, InvestigateListResponsePostDeliveryOperationQuarantineRelease, InvestigateListResponsePostDeliveryOperationSubmission, InvestigateListResponsePostDeliveryOperationMove:
		return true
	}
	return false
}

type InvestigateListResponseValidation struct {
	Comment string                                 `json:"comment" api:"nullable"`
	DKIM    InvestigateListResponseValidationDKIM  `json:"dkim"`
	DMARC   InvestigateListResponseValidationDMARC `json:"dmarc"`
	SPF     InvestigateListResponseValidationSPF   `json:"spf"`
	JSON    investigateListResponseValidationJSON  `json:"-"`
}

// investigateListResponseValidationJSON contains the JSON metadata for the struct
// [InvestigateListResponseValidation]
type investigateListResponseValidationJSON struct {
	Comment     apijson.Field
	DKIM        apijson.Field
	DMARC       apijson.Field
	SPF         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateListResponseValidation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateListResponseValidationJSON) RawJSON() string {
	return r.raw
}

type InvestigateListResponseValidationDKIM string

const (
	InvestigateListResponseValidationDKIMPass    InvestigateListResponseValidationDKIM = "pass"
	InvestigateListResponseValidationDKIMNeutral InvestigateListResponseValidationDKIM = "neutral"
	InvestigateListResponseValidationDKIMFail    InvestigateListResponseValidationDKIM = "fail"
	InvestigateListResponseValidationDKIMError   InvestigateListResponseValidationDKIM = "error"
	InvestigateListResponseValidationDKIMNone    InvestigateListResponseValidationDKIM = "none"
)

func (r InvestigateListResponseValidationDKIM) IsKnown() bool {
	switch r {
	case InvestigateListResponseValidationDKIMPass, InvestigateListResponseValidationDKIMNeutral, InvestigateListResponseValidationDKIMFail, InvestigateListResponseValidationDKIMError, InvestigateListResponseValidationDKIMNone:
		return true
	}
	return false
}

type InvestigateListResponseValidationDMARC string

const (
	InvestigateListResponseValidationDMARCPass    InvestigateListResponseValidationDMARC = "pass"
	InvestigateListResponseValidationDMARCNeutral InvestigateListResponseValidationDMARC = "neutral"
	InvestigateListResponseValidationDMARCFail    InvestigateListResponseValidationDMARC = "fail"
	InvestigateListResponseValidationDMARCError   InvestigateListResponseValidationDMARC = "error"
	InvestigateListResponseValidationDMARCNone    InvestigateListResponseValidationDMARC = "none"
)

func (r InvestigateListResponseValidationDMARC) IsKnown() bool {
	switch r {
	case InvestigateListResponseValidationDMARCPass, InvestigateListResponseValidationDMARCNeutral, InvestigateListResponseValidationDMARCFail, InvestigateListResponseValidationDMARCError, InvestigateListResponseValidationDMARCNone:
		return true
	}
	return false
}

type InvestigateListResponseValidationSPF string

const (
	InvestigateListResponseValidationSPFPass    InvestigateListResponseValidationSPF = "pass"
	InvestigateListResponseValidationSPFNeutral InvestigateListResponseValidationSPF = "neutral"
	InvestigateListResponseValidationSPFFail    InvestigateListResponseValidationSPF = "fail"
	InvestigateListResponseValidationSPFError   InvestigateListResponseValidationSPF = "error"
	InvestigateListResponseValidationSPFNone    InvestigateListResponseValidationSPF = "none"
)

func (r InvestigateListResponseValidationSPF) IsKnown() bool {
	switch r {
	case InvestigateListResponseValidationSPFPass, InvestigateListResponseValidationSPFNeutral, InvestigateListResponseValidationSPFFail, InvestigateListResponseValidationSPFError, InvestigateListResponseValidationSPFNone:
		return true
	}
	return false
}

type InvestigateGetResponse struct {
	// Unique identifier for a message retrieved from investigation
	ID string `json:"id" api:"required"`
	// Deprecated, use `GET /investigate/{investigate_id}/action_log` instead. End of
	// life: November 1, 2026.
	//
	// Deprecated: deprecated
	ActionLog         []InvestigateGetResponseActionLog `json:"action_log" api:"required"`
	ClientRecipients  []string                          `json:"client_recipients" api:"required"`
	DetectionReasons  []string                          `json:"detection_reasons" api:"required"`
	IsPhishSubmission bool                              `json:"is_phish_submission" api:"required"`
	IsQuarantined     bool                              `json:"is_quarantined" api:"required"`
	// The identifier of the message
	PostfixID string `json:"postfix_id" api:"required"`
	// Message processing properties
	Properties InvestigateGetResponseProperties `json:"properties" api:"required"`
	// Deprecated, use `scanned_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: deprecated
	Ts               string                                 `json:"ts" api:"required"`
	AlertID          string                                 `json:"alert_id" api:"nullable"`
	DeliveryMode     InvestigateGetResponseDeliveryMode     `json:"delivery_mode"`
	DeliveryStatus   []InvestigateGetResponseDeliveryStatus `json:"delivery_status" api:"nullable"`
	EdfHash          string                                 `json:"edf_hash" api:"nullable"`
	EnvelopeFrom     string                                 `json:"envelope_from" api:"nullable"`
	EnvelopeTo       []string                               `json:"envelope_to" api:"nullable"`
	FinalDisposition InvestigateGetResponseFinalDisposition `json:"final_disposition"`
	// Deprecated, use the `findings` field from
	// `GET /investigate/{investigate_id}/detections` instead. End of life: November
	// 1, 2026. Detection findings for this message.
	//
	// Deprecated: deprecated
	Findings              []InvestigateGetResponseFinding `json:"findings" api:"nullable"`
	From                  string                          `json:"from" api:"nullable"`
	FromName              string                          `json:"from_name" api:"nullable"`
	HtmltextStructureHash string                          `json:"htmltext_structure_hash" api:"nullable"`
	MessageID             string                          `json:"message_id" api:"nullable"`
	// Post-delivery operations performed on this message
	PostDeliveryOperations []InvestigateGetResponsePostDeliveryOperation `json:"post_delivery_operations" api:"nullable"`
	PostfixIDOutbound      string                                        `json:"postfix_id_outbound" api:"nullable"`
	Replyto                string                                        `json:"replyto" api:"nullable"`
	// When the message was scanned (UTC)
	ScannedAt time.Time `json:"scanned_at" api:"nullable" format:"date-time"`
	// When the message was sent (UTC)
	SentAt           time.Time                        `json:"sent_at" api:"nullable" format:"date-time"`
	SentDate         string                           `json:"sent_date" api:"nullable"`
	Subject          string                           `json:"subject" api:"nullable"`
	ThreatCategories []string                         `json:"threat_categories" api:"nullable"`
	To               []string                         `json:"to" api:"nullable"`
	ToName           []string                         `json:"to_name" api:"nullable"`
	Validation       InvestigateGetResponseValidation `json:"validation"`
	JSON             investigateGetResponseJSON       `json:"-"`
}

// investigateGetResponseJSON contains the JSON metadata for the struct
// [InvestigateGetResponse]
type investigateGetResponseJSON struct {
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
	Subject                apijson.Field
	ThreatCategories       apijson.Field
	To                     apijson.Field
	ToName                 apijson.Field
	Validation             apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *InvestigateGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateGetResponseJSON) RawJSON() string {
	return r.raw
}

type InvestigateGetResponseActionLog struct {
	// Timestamp when action completed
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	// Type of action performed
	Operation InvestigateGetResponseActionLogOperation `json:"operation" api:"required"`
	// Deprecated, use `completed_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: deprecated
	CompletedTimestamp string `json:"completed_timestamp"`
	// Additional properties for the action
	Properties InvestigateGetResponseActionLogProperties `json:"properties"`
	// Status of the action
	Status string                              `json:"status" api:"nullable"`
	JSON   investigateGetResponseActionLogJSON `json:"-"`
}

// investigateGetResponseActionLogJSON contains the JSON metadata for the struct
// [InvestigateGetResponseActionLog]
type investigateGetResponseActionLogJSON struct {
	CompletedAt        apijson.Field
	Operation          apijson.Field
	CompletedTimestamp apijson.Field
	Properties         apijson.Field
	Status             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InvestigateGetResponseActionLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateGetResponseActionLogJSON) RawJSON() string {
	return r.raw
}

// Type of action performed
type InvestigateGetResponseActionLogOperation string

const (
	InvestigateGetResponseActionLogOperationMove              InvestigateGetResponseActionLogOperation = "MOVE"
	InvestigateGetResponseActionLogOperationRelease           InvestigateGetResponseActionLogOperation = "RELEASE"
	InvestigateGetResponseActionLogOperationReclassify        InvestigateGetResponseActionLogOperation = "RECLASSIFY"
	InvestigateGetResponseActionLogOperationSubmission        InvestigateGetResponseActionLogOperation = "SUBMISSION"
	InvestigateGetResponseActionLogOperationQuarantineRelease InvestigateGetResponseActionLogOperation = "QUARANTINE_RELEASE"
	InvestigateGetResponseActionLogOperationPreview           InvestigateGetResponseActionLogOperation = "PREVIEW"
)

func (r InvestigateGetResponseActionLogOperation) IsKnown() bool {
	switch r {
	case InvestigateGetResponseActionLogOperationMove, InvestigateGetResponseActionLogOperationRelease, InvestigateGetResponseActionLogOperationReclassify, InvestigateGetResponseActionLogOperationSubmission, InvestigateGetResponseActionLogOperationQuarantineRelease, InvestigateGetResponseActionLogOperationPreview:
		return true
	}
	return false
}

// Additional properties for the action
type InvestigateGetResponseActionLogProperties struct {
	// Target folder for move operations
	Folder string `json:"folder"`
	// User who requested the action
	RequestedBy string                                        `json:"requested_by"`
	JSON        investigateGetResponseActionLogPropertiesJSON `json:"-"`
}

// investigateGetResponseActionLogPropertiesJSON contains the JSON metadata for the
// struct [InvestigateGetResponseActionLogProperties]
type investigateGetResponseActionLogPropertiesJSON struct {
	Folder      apijson.Field
	RequestedBy apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateGetResponseActionLogProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateGetResponseActionLogPropertiesJSON) RawJSON() string {
	return r.raw
}

// Message processing properties
type InvestigateGetResponseProperties struct {
	// Pattern that allowlisted this message
	AllowlistedPattern string `json:"allowlisted_pattern" api:"nullable"`
	// Type of allowlist pattern
	AllowlistedPatternType InvestigateGetResponsePropertiesAllowlistedPatternType `json:"allowlisted_pattern_type" api:"nullable"`
	// Whether message was blocklisted
	BlocklistedMessage bool `json:"blocklisted_message" api:"nullable"`
	// Pattern that blocklisted this message
	BlocklistedPattern string `json:"blocklisted_pattern" api:"nullable"`
	// Legacy field for allowlist pattern type
	WhitelistedPatternType InvestigateGetResponsePropertiesWhitelistedPatternType `json:"whitelisted_pattern_type" api:"nullable"`
	JSON                   investigateGetResponsePropertiesJSON                   `json:"-"`
}

// investigateGetResponsePropertiesJSON contains the JSON metadata for the struct
// [InvestigateGetResponseProperties]
type investigateGetResponsePropertiesJSON struct {
	AllowlistedPattern     apijson.Field
	AllowlistedPatternType apijson.Field
	BlocklistedMessage     apijson.Field
	BlocklistedPattern     apijson.Field
	WhitelistedPatternType apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *InvestigateGetResponseProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateGetResponsePropertiesJSON) RawJSON() string {
	return r.raw
}

// Type of allowlist pattern
type InvestigateGetResponsePropertiesAllowlistedPatternType string

const (
	InvestigateGetResponsePropertiesAllowlistedPatternTypeQuarantineRelease       InvestigateGetResponsePropertiesAllowlistedPatternType = "quarantine_release"
	InvestigateGetResponsePropertiesAllowlistedPatternTypeAcceptableSender        InvestigateGetResponsePropertiesAllowlistedPatternType = "acceptable_sender"
	InvestigateGetResponsePropertiesAllowlistedPatternTypeAllowedSender           InvestigateGetResponsePropertiesAllowlistedPatternType = "allowed_sender"
	InvestigateGetResponsePropertiesAllowlistedPatternTypeAllowedRecipient        InvestigateGetResponsePropertiesAllowlistedPatternType = "allowed_recipient"
	InvestigateGetResponsePropertiesAllowlistedPatternTypeDomainSimilarity        InvestigateGetResponsePropertiesAllowlistedPatternType = "domain_similarity"
	InvestigateGetResponsePropertiesAllowlistedPatternTypeDomainRecency           InvestigateGetResponsePropertiesAllowlistedPatternType = "domain_recency"
	InvestigateGetResponsePropertiesAllowlistedPatternTypeManagedAcceptableSender InvestigateGetResponsePropertiesAllowlistedPatternType = "managed_acceptable_sender"
	InvestigateGetResponsePropertiesAllowlistedPatternTypeOutboundNdr             InvestigateGetResponsePropertiesAllowlistedPatternType = "outbound_ndr"
)

func (r InvestigateGetResponsePropertiesAllowlistedPatternType) IsKnown() bool {
	switch r {
	case InvestigateGetResponsePropertiesAllowlistedPatternTypeQuarantineRelease, InvestigateGetResponsePropertiesAllowlistedPatternTypeAcceptableSender, InvestigateGetResponsePropertiesAllowlistedPatternTypeAllowedSender, InvestigateGetResponsePropertiesAllowlistedPatternTypeAllowedRecipient, InvestigateGetResponsePropertiesAllowlistedPatternTypeDomainSimilarity, InvestigateGetResponsePropertiesAllowlistedPatternTypeDomainRecency, InvestigateGetResponsePropertiesAllowlistedPatternTypeManagedAcceptableSender, InvestigateGetResponsePropertiesAllowlistedPatternTypeOutboundNdr:
		return true
	}
	return false
}

// Legacy field for allowlist pattern type
type InvestigateGetResponsePropertiesWhitelistedPatternType string

const (
	InvestigateGetResponsePropertiesWhitelistedPatternTypeQuarantineRelease       InvestigateGetResponsePropertiesWhitelistedPatternType = "quarantine_release"
	InvestigateGetResponsePropertiesWhitelistedPatternTypeAcceptableSender        InvestigateGetResponsePropertiesWhitelistedPatternType = "acceptable_sender"
	InvestigateGetResponsePropertiesWhitelistedPatternTypeAllowedSender           InvestigateGetResponsePropertiesWhitelistedPatternType = "allowed_sender"
	InvestigateGetResponsePropertiesWhitelistedPatternTypeAllowedRecipient        InvestigateGetResponsePropertiesWhitelistedPatternType = "allowed_recipient"
	InvestigateGetResponsePropertiesWhitelistedPatternTypeDomainSimilarity        InvestigateGetResponsePropertiesWhitelistedPatternType = "domain_similarity"
	InvestigateGetResponsePropertiesWhitelistedPatternTypeDomainRecency           InvestigateGetResponsePropertiesWhitelistedPatternType = "domain_recency"
	InvestigateGetResponsePropertiesWhitelistedPatternTypeManagedAcceptableSender InvestigateGetResponsePropertiesWhitelistedPatternType = "managed_acceptable_sender"
	InvestigateGetResponsePropertiesWhitelistedPatternTypeOutboundNdr             InvestigateGetResponsePropertiesWhitelistedPatternType = "outbound_ndr"
)

func (r InvestigateGetResponsePropertiesWhitelistedPatternType) IsKnown() bool {
	switch r {
	case InvestigateGetResponsePropertiesWhitelistedPatternTypeQuarantineRelease, InvestigateGetResponsePropertiesWhitelistedPatternTypeAcceptableSender, InvestigateGetResponsePropertiesWhitelistedPatternTypeAllowedSender, InvestigateGetResponsePropertiesWhitelistedPatternTypeAllowedRecipient, InvestigateGetResponsePropertiesWhitelistedPatternTypeDomainSimilarity, InvestigateGetResponsePropertiesWhitelistedPatternTypeDomainRecency, InvestigateGetResponsePropertiesWhitelistedPatternTypeManagedAcceptableSender, InvestigateGetResponsePropertiesWhitelistedPatternTypeOutboundNdr:
		return true
	}
	return false
}

type InvestigateGetResponseDeliveryMode string

const (
	InvestigateGetResponseDeliveryModeDirect                InvestigateGetResponseDeliveryMode = "DIRECT"
	InvestigateGetResponseDeliveryModeBcc                   InvestigateGetResponseDeliveryMode = "BCC"
	InvestigateGetResponseDeliveryModeJournal               InvestigateGetResponseDeliveryMode = "JOURNAL"
	InvestigateGetResponseDeliveryModeReviewSubmission      InvestigateGetResponseDeliveryMode = "REVIEW_SUBMISSION"
	InvestigateGetResponseDeliveryModeDMARCUnverified       InvestigateGetResponseDeliveryMode = "DMARC_UNVERIFIED"
	InvestigateGetResponseDeliveryModeDMARCFailureReport    InvestigateGetResponseDeliveryMode = "DMARC_FAILURE_REPORT"
	InvestigateGetResponseDeliveryModeDMARCAggregateReport  InvestigateGetResponseDeliveryMode = "DMARC_AGGREGATE_REPORT"
	InvestigateGetResponseDeliveryModeThreatIntelSubmission InvestigateGetResponseDeliveryMode = "THREAT_INTEL_SUBMISSION"
	InvestigateGetResponseDeliveryModeSimulationSubmission  InvestigateGetResponseDeliveryMode = "SIMULATION_SUBMISSION"
	InvestigateGetResponseDeliveryModeAPI                   InvestigateGetResponseDeliveryMode = "API"
	InvestigateGetResponseDeliveryModeRetroScan             InvestigateGetResponseDeliveryMode = "RETRO_SCAN"
)

func (r InvestigateGetResponseDeliveryMode) IsKnown() bool {
	switch r {
	case InvestigateGetResponseDeliveryModeDirect, InvestigateGetResponseDeliveryModeBcc, InvestigateGetResponseDeliveryModeJournal, InvestigateGetResponseDeliveryModeReviewSubmission, InvestigateGetResponseDeliveryModeDMARCUnverified, InvestigateGetResponseDeliveryModeDMARCFailureReport, InvestigateGetResponseDeliveryModeDMARCAggregateReport, InvestigateGetResponseDeliveryModeThreatIntelSubmission, InvestigateGetResponseDeliveryModeSimulationSubmission, InvestigateGetResponseDeliveryModeAPI, InvestigateGetResponseDeliveryModeRetroScan:
		return true
	}
	return false
}

// Delivery status of the message.
type InvestigateGetResponseDeliveryStatus string

const (
	InvestigateGetResponseDeliveryStatusDelivered   InvestigateGetResponseDeliveryStatus = "delivered"
	InvestigateGetResponseDeliveryStatusMoved       InvestigateGetResponseDeliveryStatus = "moved"
	InvestigateGetResponseDeliveryStatusQuarantined InvestigateGetResponseDeliveryStatus = "quarantined"
	InvestigateGetResponseDeliveryStatusRejected    InvestigateGetResponseDeliveryStatus = "rejected"
	InvestigateGetResponseDeliveryStatusDeferred    InvestigateGetResponseDeliveryStatus = "deferred"
	InvestigateGetResponseDeliveryStatusBounced     InvestigateGetResponseDeliveryStatus = "bounced"
	InvestigateGetResponseDeliveryStatusQueued      InvestigateGetResponseDeliveryStatus = "queued"
)

func (r InvestigateGetResponseDeliveryStatus) IsKnown() bool {
	switch r {
	case InvestigateGetResponseDeliveryStatusDelivered, InvestigateGetResponseDeliveryStatusMoved, InvestigateGetResponseDeliveryStatusQuarantined, InvestigateGetResponseDeliveryStatusRejected, InvestigateGetResponseDeliveryStatusDeferred, InvestigateGetResponseDeliveryStatusBounced, InvestigateGetResponseDeliveryStatusQueued:
		return true
	}
	return false
}

type InvestigateGetResponseFinalDisposition string

const (
	InvestigateGetResponseFinalDispositionMalicious    InvestigateGetResponseFinalDisposition = "MALICIOUS"
	InvestigateGetResponseFinalDispositionMaliciousBec InvestigateGetResponseFinalDisposition = "MALICIOUS-BEC"
	InvestigateGetResponseFinalDispositionSuspicious   InvestigateGetResponseFinalDisposition = "SUSPICIOUS"
	InvestigateGetResponseFinalDispositionSpoof        InvestigateGetResponseFinalDisposition = "SPOOF"
	InvestigateGetResponseFinalDispositionSpam         InvestigateGetResponseFinalDisposition = "SPAM"
	InvestigateGetResponseFinalDispositionBulk         InvestigateGetResponseFinalDisposition = "BULK"
	InvestigateGetResponseFinalDispositionEncrypted    InvestigateGetResponseFinalDisposition = "ENCRYPTED"
	InvestigateGetResponseFinalDispositionExternal     InvestigateGetResponseFinalDisposition = "EXTERNAL"
	InvestigateGetResponseFinalDispositionUnknown      InvestigateGetResponseFinalDisposition = "UNKNOWN"
	InvestigateGetResponseFinalDispositionNone         InvestigateGetResponseFinalDisposition = "NONE"
)

func (r InvestigateGetResponseFinalDisposition) IsKnown() bool {
	switch r {
	case InvestigateGetResponseFinalDispositionMalicious, InvestigateGetResponseFinalDispositionMaliciousBec, InvestigateGetResponseFinalDispositionSuspicious, InvestigateGetResponseFinalDispositionSpoof, InvestigateGetResponseFinalDispositionSpam, InvestigateGetResponseFinalDispositionBulk, InvestigateGetResponseFinalDispositionEncrypted, InvestigateGetResponseFinalDispositionExternal, InvestigateGetResponseFinalDispositionUnknown, InvestigateGetResponseFinalDispositionNone:
		return true
	}
	return false
}

type InvestigateGetResponseFinding struct {
	Attachment string                                  `json:"attachment" api:"nullable"`
	Detail     string                                  `json:"detail" api:"nullable"`
	Detection  InvestigateGetResponseFindingsDetection `json:"detection"`
	Field      string                                  `json:"field" api:"nullable"`
	Name       string                                  `json:"name" api:"nullable"`
	Portion    string                                  `json:"portion" api:"nullable"`
	Reason     string                                  `json:"reason" api:"nullable"`
	Score      float64                                 `json:"score" api:"nullable"`
	Value      string                                  `json:"value" api:"nullable"`
	JSON       investigateGetResponseFindingJSON       `json:"-"`
}

// investigateGetResponseFindingJSON contains the JSON metadata for the struct
// [InvestigateGetResponseFinding]
type investigateGetResponseFindingJSON struct {
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

func (r *InvestigateGetResponseFinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateGetResponseFindingJSON) RawJSON() string {
	return r.raw
}

type InvestigateGetResponseFindingsDetection string

const (
	InvestigateGetResponseFindingsDetectionMalicious    InvestigateGetResponseFindingsDetection = "MALICIOUS"
	InvestigateGetResponseFindingsDetectionMaliciousBec InvestigateGetResponseFindingsDetection = "MALICIOUS-BEC"
	InvestigateGetResponseFindingsDetectionSuspicious   InvestigateGetResponseFindingsDetection = "SUSPICIOUS"
	InvestigateGetResponseFindingsDetectionSpoof        InvestigateGetResponseFindingsDetection = "SPOOF"
	InvestigateGetResponseFindingsDetectionSpam         InvestigateGetResponseFindingsDetection = "SPAM"
	InvestigateGetResponseFindingsDetectionBulk         InvestigateGetResponseFindingsDetection = "BULK"
	InvestigateGetResponseFindingsDetectionEncrypted    InvestigateGetResponseFindingsDetection = "ENCRYPTED"
	InvestigateGetResponseFindingsDetectionExternal     InvestigateGetResponseFindingsDetection = "EXTERNAL"
	InvestigateGetResponseFindingsDetectionUnknown      InvestigateGetResponseFindingsDetection = "UNKNOWN"
	InvestigateGetResponseFindingsDetectionNone         InvestigateGetResponseFindingsDetection = "NONE"
)

func (r InvestigateGetResponseFindingsDetection) IsKnown() bool {
	switch r {
	case InvestigateGetResponseFindingsDetectionMalicious, InvestigateGetResponseFindingsDetectionMaliciousBec, InvestigateGetResponseFindingsDetectionSuspicious, InvestigateGetResponseFindingsDetectionSpoof, InvestigateGetResponseFindingsDetectionSpam, InvestigateGetResponseFindingsDetectionBulk, InvestigateGetResponseFindingsDetectionEncrypted, InvestigateGetResponseFindingsDetectionExternal, InvestigateGetResponseFindingsDetectionUnknown, InvestigateGetResponseFindingsDetectionNone:
		return true
	}
	return false
}

type InvestigateGetResponsePostDeliveryOperation string

const (
	InvestigateGetResponsePostDeliveryOperationPreview           InvestigateGetResponsePostDeliveryOperation = "PREVIEW"
	InvestigateGetResponsePostDeliveryOperationQuarantineRelease InvestigateGetResponsePostDeliveryOperation = "QUARANTINE_RELEASE"
	InvestigateGetResponsePostDeliveryOperationSubmission        InvestigateGetResponsePostDeliveryOperation = "SUBMISSION"
	InvestigateGetResponsePostDeliveryOperationMove              InvestigateGetResponsePostDeliveryOperation = "MOVE"
)

func (r InvestigateGetResponsePostDeliveryOperation) IsKnown() bool {
	switch r {
	case InvestigateGetResponsePostDeliveryOperationPreview, InvestigateGetResponsePostDeliveryOperationQuarantineRelease, InvestigateGetResponsePostDeliveryOperationSubmission, InvestigateGetResponsePostDeliveryOperationMove:
		return true
	}
	return false
}

type InvestigateGetResponseValidation struct {
	Comment string                                `json:"comment" api:"nullable"`
	DKIM    InvestigateGetResponseValidationDKIM  `json:"dkim"`
	DMARC   InvestigateGetResponseValidationDMARC `json:"dmarc"`
	SPF     InvestigateGetResponseValidationSPF   `json:"spf"`
	JSON    investigateGetResponseValidationJSON  `json:"-"`
}

// investigateGetResponseValidationJSON contains the JSON metadata for the struct
// [InvestigateGetResponseValidation]
type investigateGetResponseValidationJSON struct {
	Comment     apijson.Field
	DKIM        apijson.Field
	DMARC       apijson.Field
	SPF         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateGetResponseValidation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateGetResponseValidationJSON) RawJSON() string {
	return r.raw
}

type InvestigateGetResponseValidationDKIM string

const (
	InvestigateGetResponseValidationDKIMPass    InvestigateGetResponseValidationDKIM = "pass"
	InvestigateGetResponseValidationDKIMNeutral InvestigateGetResponseValidationDKIM = "neutral"
	InvestigateGetResponseValidationDKIMFail    InvestigateGetResponseValidationDKIM = "fail"
	InvestigateGetResponseValidationDKIMError   InvestigateGetResponseValidationDKIM = "error"
	InvestigateGetResponseValidationDKIMNone    InvestigateGetResponseValidationDKIM = "none"
)

func (r InvestigateGetResponseValidationDKIM) IsKnown() bool {
	switch r {
	case InvestigateGetResponseValidationDKIMPass, InvestigateGetResponseValidationDKIMNeutral, InvestigateGetResponseValidationDKIMFail, InvestigateGetResponseValidationDKIMError, InvestigateGetResponseValidationDKIMNone:
		return true
	}
	return false
}

type InvestigateGetResponseValidationDMARC string

const (
	InvestigateGetResponseValidationDMARCPass    InvestigateGetResponseValidationDMARC = "pass"
	InvestigateGetResponseValidationDMARCNeutral InvestigateGetResponseValidationDMARC = "neutral"
	InvestigateGetResponseValidationDMARCFail    InvestigateGetResponseValidationDMARC = "fail"
	InvestigateGetResponseValidationDMARCError   InvestigateGetResponseValidationDMARC = "error"
	InvestigateGetResponseValidationDMARCNone    InvestigateGetResponseValidationDMARC = "none"
)

func (r InvestigateGetResponseValidationDMARC) IsKnown() bool {
	switch r {
	case InvestigateGetResponseValidationDMARCPass, InvestigateGetResponseValidationDMARCNeutral, InvestigateGetResponseValidationDMARCFail, InvestigateGetResponseValidationDMARCError, InvestigateGetResponseValidationDMARCNone:
		return true
	}
	return false
}

type InvestigateGetResponseValidationSPF string

const (
	InvestigateGetResponseValidationSPFPass    InvestigateGetResponseValidationSPF = "pass"
	InvestigateGetResponseValidationSPFNeutral InvestigateGetResponseValidationSPF = "neutral"
	InvestigateGetResponseValidationSPFFail    InvestigateGetResponseValidationSPF = "fail"
	InvestigateGetResponseValidationSPFError   InvestigateGetResponseValidationSPF = "error"
	InvestigateGetResponseValidationSPFNone    InvestigateGetResponseValidationSPF = "none"
)

func (r InvestigateGetResponseValidationSPF) IsKnown() bool {
	switch r {
	case InvestigateGetResponseValidationSPFPass, InvestigateGetResponseValidationSPFNeutral, InvestigateGetResponseValidationSPFFail, InvestigateGetResponseValidationSPFError, InvestigateGetResponseValidationSPFNone:
		return true
	}
	return false
}

type InvestigateListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Whether to include the message action log in the response.
	ActionLog param.Field[bool]   `query:"action_log"`
	AlertID   param.Field[string] `query:"alert_id"`
	Cursor    param.Field[string] `query:"cursor"`
	// Whether to include only detections in search results.
	DetectionsOnly param.Field[bool] `query:"detections_only"`
	// Sender domains to filter by.
	Domain param.Field[string] `query:"domain"`
	// The end of the search date range. Defaults to `now`.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// Dispositions to filter by.
	FinalDisposition param.Field[InvestigateListParamsFinalDisposition] `query:"final_disposition"`
	// Message actions to filter by.
	MessageAction param.Field[InvestigateListParamsMessageAction] `query:"message_action"`
	MessageID     param.Field[string]                             `query:"message_id"`
	Metric        param.Field[string]                             `query:"metric"`
	// Deprecated: Use cursor pagination instead. End of life: November 1, 2026.
	Page param.Field[int64] `query:"page"`
	// The number of results per page. Maximum value is 1000.
	PerPage param.Field[int64] `query:"per_page"`
	// Space-delimited search term. Case-insensitive.
	Query     param.Field[string] `query:"query"`
	Recipient param.Field[string] `query:"recipient"`
	Sender    param.Field[string] `query:"sender"`
	// The beginning of the search date range. Defaults to `now - 30 days`.
	Start   param.Field[time.Time] `query:"start" format:"date-time"`
	Subject param.Field[string]    `query:"subject"`
}

// URLQuery serializes [InvestigateListParams]'s query parameters as `url.Values`.
func (r InvestigateListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Dispositions to filter by.
type InvestigateListParamsFinalDisposition string

const (
	InvestigateListParamsFinalDispositionMalicious  InvestigateListParamsFinalDisposition = "MALICIOUS"
	InvestigateListParamsFinalDispositionSuspicious InvestigateListParamsFinalDisposition = "SUSPICIOUS"
	InvestigateListParamsFinalDispositionSpoof      InvestigateListParamsFinalDisposition = "SPOOF"
	InvestigateListParamsFinalDispositionSpam       InvestigateListParamsFinalDisposition = "SPAM"
	InvestigateListParamsFinalDispositionBulk       InvestigateListParamsFinalDisposition = "BULK"
	InvestigateListParamsFinalDispositionNone       InvestigateListParamsFinalDisposition = "NONE"
)

func (r InvestigateListParamsFinalDisposition) IsKnown() bool {
	switch r {
	case InvestigateListParamsFinalDispositionMalicious, InvestigateListParamsFinalDispositionSuspicious, InvestigateListParamsFinalDispositionSpoof, InvestigateListParamsFinalDispositionSpam, InvestigateListParamsFinalDispositionBulk, InvestigateListParamsFinalDispositionNone:
		return true
	}
	return false
}

// Message actions to filter by.
type InvestigateListParamsMessageAction string

const (
	InvestigateListParamsMessageActionPreview            InvestigateListParamsMessageAction = "PREVIEW"
	InvestigateListParamsMessageActionQuarantineReleased InvestigateListParamsMessageAction = "QUARANTINE_RELEASED"
	InvestigateListParamsMessageActionMoved              InvestigateListParamsMessageAction = "MOVED"
)

func (r InvestigateListParamsMessageAction) IsKnown() bool {
	switch r {
	case InvestigateListParamsMessageActionPreview, InvestigateListParamsMessageActionQuarantineReleased, InvestigateListParamsMessageActionMoved:
		return true
	}
	return false
}

type InvestigateGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// When true, search the submissions datastore only. When false or omitted, search
	// the regular datastore only.
	Submission param.Field[bool] `query:"submission"`
}

// URLQuery serializes [InvestigateGetParams]'s query parameters as `url.Values`.
func (r InvestigateGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InvestigateGetResponseEnvelope struct {
	Errors   []InvestigateGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []InvestigateGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   InvestigateGetResponse                   `json:"result" api:"required"`
	// Whether the API call was successful.
	Success InvestigateGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    investigateGetResponseEnvelopeJSON    `json:"-"`
}

// investigateGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [InvestigateGetResponseEnvelope]
type investigateGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InvestigateGetResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code" api:"required"`
	Message          string                                     `json:"message" api:"required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           InvestigateGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             investigateGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// investigateGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [InvestigateGetResponseEnvelopeErrors]
type investigateGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvestigateGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type InvestigateGetResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    investigateGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// investigateGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [InvestigateGetResponseEnvelopeErrorsSource]
type investigateGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type InvestigateGetResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code" api:"required"`
	Message          string                                       `json:"message" api:"required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           InvestigateGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             investigateGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// investigateGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [InvestigateGetResponseEnvelopeMessages]
type investigateGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvestigateGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type InvestigateGetResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    investigateGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// investigateGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [InvestigateGetResponseEnvelopeMessagesSource]
type investigateGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type InvestigateGetResponseEnvelopeSuccess bool

const (
	InvestigateGetResponseEnvelopeSuccessTrue InvestigateGetResponseEnvelopeSuccess = true
)

func (r InvestigateGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case InvestigateGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
