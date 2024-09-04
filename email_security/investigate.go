// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// InvestigateService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvestigateService] method instead.
type InvestigateService struct {
	Options []option.RequestOption
}

// NewInvestigateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInvestigateService(opts ...option.RequestOption) (r *InvestigateService) {
	r = &InvestigateService{}
	r.Options = opts
	return
}

// This endpoint returns information for each email that matches the search
// parameter(s).
func (r *InvestigateService) List(ctx context.Context, params InvestigateListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[InvestigateListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
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

// This endpoint returns information for each email that matches the search
// parameter(s).
func (r *InvestigateService) ListAutoPaging(ctx context.Context, params InvestigateListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[InvestigateListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// For emails that have a detection, this endpoint returns detection details such
// as threat categories, sender information, and links.
func (r *InvestigateService) Detections(ctx context.Context, postfixID string, query InvestigateDetectionsParams, opts ...option.RequestOption) (res *InvestigateDetectionsResponse, err error) {
	var env InvestigateDetectionsResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if postfixID == "" {
		err = errors.New("missing required postfix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/%s/detections", query.AccountID, postfixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get message details
func (r *InvestigateService) Get(ctx context.Context, postfixID string, query InvestigateGetParams, opts ...option.RequestOption) (res *InvestigateGetResponse, err error) {
	var env InvestigateGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if postfixID == "" {
		err = errors.New("missing required postfix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/%s", query.AccountID, postfixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// For emails that have a detection, this endpoint returns a preview of the message
// body as a base64 encoded PNG image.
func (r *InvestigateService) Preview(ctx context.Context, postfixID string, query InvestigatePreviewParams, opts ...option.RequestOption) (res *InvestigatePreviewResponse, err error) {
	var env InvestigatePreviewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if postfixID == "" {
		err = errors.New("missing required postfix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/%s/preview", query.AccountID, postfixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// For emails that have a detection, this endpoint returns the raw email as an EML
// file.
func (r *InvestigateService) Raw(ctx context.Context, postfixID string, query InvestigateRawParams, opts ...option.RequestOption) (res *InvestigateRawResponse, err error) {
	var env InvestigateRawResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if postfixID == "" {
		err = errors.New("missing required postfix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/%s/raw", query.AccountID, postfixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get email trace
func (r *InvestigateService) Trace(ctx context.Context, postfixID string, query InvestigateTraceParams, opts ...option.RequestOption) (res *InvestigateTraceResponse, err error) {
	var env InvestigateTraceResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if postfixID == "" {
		err = errors.New("missing required postfix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/%s/trace", query.AccountID, postfixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type InvestigateListResponse struct {
	ID                string      `json:"id,required"`
	ActionLog         interface{} `json:"action_log,required"`
	ClientRecipients  []string    `json:"client_recipients,required"`
	DetectionReasons  []string    `json:"detection_reasons,required"`
	IsPhishSubmission bool        `json:"is_phish_submission,required"`
	IsQuarantined     bool        `json:"is_quarantined,required"`
	MessageID         string      `json:"message_id,required"`
	// Message identifier
	PostfixID        string                                  `json:"postfix_id,required"`
	Ts               string                                  `json:"ts,required"`
	AlertID          string                                  `json:"alert_id,nullable"`
	EdfHash          string                                  `json:"edf_hash,nullable"`
	FinalDisposition InvestigateListResponseFinalDisposition `json:"final_disposition,nullable"`
	From             string                                  `json:"from,nullable"`
	FromName         string                                  `json:"from_name,nullable"`
	SentDate         string                                  `json:"sent_date,nullable"`
	Subject          string                                  `json:"subject,nullable"`
	ThreatCategories []string                                `json:"threat_categories,nullable"`
	To               []string                                `json:"to,nullable"`
	ToName           []string                                `json:"to_name,nullable"`
	Validation       InvestigateListResponseValidation       `json:"validation,nullable"`
	JSON             investigateListResponseJSON             `json:"-"`
}

// investigateListResponseJSON contains the JSON metadata for the struct
// [InvestigateListResponse]
type investigateListResponseJSON struct {
	ID                apijson.Field
	ActionLog         apijson.Field
	ClientRecipients  apijson.Field
	DetectionReasons  apijson.Field
	IsPhishSubmission apijson.Field
	IsQuarantined     apijson.Field
	MessageID         apijson.Field
	PostfixID         apijson.Field
	Ts                apijson.Field
	AlertID           apijson.Field
	EdfHash           apijson.Field
	FinalDisposition  apijson.Field
	From              apijson.Field
	FromName          apijson.Field
	SentDate          apijson.Field
	Subject           apijson.Field
	ThreatCategories  apijson.Field
	To                apijson.Field
	ToName            apijson.Field
	Validation        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvestigateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateListResponseJSON) RawJSON() string {
	return r.raw
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

type InvestigateListResponseValidation struct {
	Comment string                                 `json:"comment,nullable"`
	DKIM    InvestigateListResponseValidationDKIM  `json:"dkim,nullable"`
	DMARC   InvestigateListResponseValidationDMARC `json:"dmarc,nullable"`
	SPF     InvestigateListResponseValidationSPF   `json:"spf,nullable"`
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

type InvestigateDetectionsResponse struct {
	Action           string                                        `json:"action,required"`
	Attachments      []InvestigateDetectionsResponseAttachment     `json:"attachments,required"`
	Headers          []InvestigateDetectionsResponseHeader         `json:"headers,required"`
	Links            []InvestigateDetectionsResponseLink           `json:"links,required"`
	SenderInfo       InvestigateDetectionsResponseSenderInfo       `json:"sender_info,required"`
	ThreatCategories []InvestigateDetectionsResponseThreatCategory `json:"threat_categories,required"`
	Validation       InvestigateDetectionsResponseValidation       `json:"validation,required"`
	FinalDisposition InvestigateDetectionsResponseFinalDisposition `json:"final_disposition,nullable"`
	JSON             investigateDetectionsResponseJSON             `json:"-"`
}

// investigateDetectionsResponseJSON contains the JSON metadata for the struct
// [InvestigateDetectionsResponse]
type investigateDetectionsResponseJSON struct {
	Action           apijson.Field
	Attachments      apijson.Field
	Headers          apijson.Field
	Links            apijson.Field
	SenderInfo       apijson.Field
	ThreatCategories apijson.Field
	Validation       apijson.Field
	FinalDisposition apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvestigateDetectionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateDetectionsResponseJSON) RawJSON() string {
	return r.raw
}

type InvestigateDetectionsResponseAttachment struct {
	Size        int64                                             `json:"size,required"`
	ContentType string                                            `json:"content_type,nullable"`
	Detection   InvestigateDetectionsResponseAttachmentsDetection `json:"detection,nullable"`
	Encrypted   bool                                              `json:"encrypted,nullable"`
	Name        string                                            `json:"name,nullable"`
	JSON        investigateDetectionsResponseAttachmentJSON       `json:"-"`
}

// investigateDetectionsResponseAttachmentJSON contains the JSON metadata for the
// struct [InvestigateDetectionsResponseAttachment]
type investigateDetectionsResponseAttachmentJSON struct {
	Size        apijson.Field
	ContentType apijson.Field
	Detection   apijson.Field
	Encrypted   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateDetectionsResponseAttachment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateDetectionsResponseAttachmentJSON) RawJSON() string {
	return r.raw
}

type InvestigateDetectionsResponseAttachmentsDetection string

const (
	InvestigateDetectionsResponseAttachmentsDetectionMalicious    InvestigateDetectionsResponseAttachmentsDetection = "MALICIOUS"
	InvestigateDetectionsResponseAttachmentsDetectionMaliciousBec InvestigateDetectionsResponseAttachmentsDetection = "MALICIOUS-BEC"
	InvestigateDetectionsResponseAttachmentsDetectionSuspicious   InvestigateDetectionsResponseAttachmentsDetection = "SUSPICIOUS"
	InvestigateDetectionsResponseAttachmentsDetectionSpoof        InvestigateDetectionsResponseAttachmentsDetection = "SPOOF"
	InvestigateDetectionsResponseAttachmentsDetectionSpam         InvestigateDetectionsResponseAttachmentsDetection = "SPAM"
	InvestigateDetectionsResponseAttachmentsDetectionBulk         InvestigateDetectionsResponseAttachmentsDetection = "BULK"
	InvestigateDetectionsResponseAttachmentsDetectionEncrypted    InvestigateDetectionsResponseAttachmentsDetection = "ENCRYPTED"
	InvestigateDetectionsResponseAttachmentsDetectionExternal     InvestigateDetectionsResponseAttachmentsDetection = "EXTERNAL"
	InvestigateDetectionsResponseAttachmentsDetectionUnknown      InvestigateDetectionsResponseAttachmentsDetection = "UNKNOWN"
	InvestigateDetectionsResponseAttachmentsDetectionNone         InvestigateDetectionsResponseAttachmentsDetection = "NONE"
)

func (r InvestigateDetectionsResponseAttachmentsDetection) IsKnown() bool {
	switch r {
	case InvestigateDetectionsResponseAttachmentsDetectionMalicious, InvestigateDetectionsResponseAttachmentsDetectionMaliciousBec, InvestigateDetectionsResponseAttachmentsDetectionSuspicious, InvestigateDetectionsResponseAttachmentsDetectionSpoof, InvestigateDetectionsResponseAttachmentsDetectionSpam, InvestigateDetectionsResponseAttachmentsDetectionBulk, InvestigateDetectionsResponseAttachmentsDetectionEncrypted, InvestigateDetectionsResponseAttachmentsDetectionExternal, InvestigateDetectionsResponseAttachmentsDetectionUnknown, InvestigateDetectionsResponseAttachmentsDetectionNone:
		return true
	}
	return false
}

type InvestigateDetectionsResponseHeader struct {
	Name  string                                  `json:"name,required"`
	Value string                                  `json:"value,required"`
	JSON  investigateDetectionsResponseHeaderJSON `json:"-"`
}

// investigateDetectionsResponseHeaderJSON contains the JSON metadata for the
// struct [InvestigateDetectionsResponseHeader]
type investigateDetectionsResponseHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateDetectionsResponseHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateDetectionsResponseHeaderJSON) RawJSON() string {
	return r.raw
}

type InvestigateDetectionsResponseLink struct {
	Href string                                `json:"href,required"`
	Text string                                `json:"text,nullable"`
	JSON investigateDetectionsResponseLinkJSON `json:"-"`
}

// investigateDetectionsResponseLinkJSON contains the JSON metadata for the struct
// [InvestigateDetectionsResponseLink]
type investigateDetectionsResponseLinkJSON struct {
	Href        apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateDetectionsResponseLink) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateDetectionsResponseLinkJSON) RawJSON() string {
	return r.raw
}

type InvestigateDetectionsResponseSenderInfo struct {
	// Name of the autonomous system
	AsName string `json:"as_name,nullable"`
	// Number of the autonomous system
	AsNumber int64                                       `json:"as_number,nullable"`
	Geo      string                                      `json:"geo,nullable"`
	IP       string                                      `json:"ip,nullable"`
	Pld      string                                      `json:"pld,nullable"`
	JSON     investigateDetectionsResponseSenderInfoJSON `json:"-"`
}

// investigateDetectionsResponseSenderInfoJSON contains the JSON metadata for the
// struct [InvestigateDetectionsResponseSenderInfo]
type investigateDetectionsResponseSenderInfoJSON struct {
	AsName      apijson.Field
	AsNumber    apijson.Field
	Geo         apijson.Field
	IP          apijson.Field
	Pld         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateDetectionsResponseSenderInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateDetectionsResponseSenderInfoJSON) RawJSON() string {
	return r.raw
}

type InvestigateDetectionsResponseThreatCategory struct {
	ID          int64                                           `json:"id,required"`
	Description string                                          `json:"description,nullable"`
	Name        string                                          `json:"name,nullable"`
	JSON        investigateDetectionsResponseThreatCategoryJSON `json:"-"`
}

// investigateDetectionsResponseThreatCategoryJSON contains the JSON metadata for
// the struct [InvestigateDetectionsResponseThreatCategory]
type investigateDetectionsResponseThreatCategoryJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateDetectionsResponseThreatCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateDetectionsResponseThreatCategoryJSON) RawJSON() string {
	return r.raw
}

type InvestigateDetectionsResponseValidation struct {
	Comment string                                       `json:"comment,nullable"`
	DKIM    InvestigateDetectionsResponseValidationDKIM  `json:"dkim,nullable"`
	DMARC   InvestigateDetectionsResponseValidationDMARC `json:"dmarc,nullable"`
	SPF     InvestigateDetectionsResponseValidationSPF   `json:"spf,nullable"`
	JSON    investigateDetectionsResponseValidationJSON  `json:"-"`
}

// investigateDetectionsResponseValidationJSON contains the JSON metadata for the
// struct [InvestigateDetectionsResponseValidation]
type investigateDetectionsResponseValidationJSON struct {
	Comment     apijson.Field
	DKIM        apijson.Field
	DMARC       apijson.Field
	SPF         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateDetectionsResponseValidation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateDetectionsResponseValidationJSON) RawJSON() string {
	return r.raw
}

type InvestigateDetectionsResponseValidationDKIM string

const (
	InvestigateDetectionsResponseValidationDKIMPass    InvestigateDetectionsResponseValidationDKIM = "pass"
	InvestigateDetectionsResponseValidationDKIMNeutral InvestigateDetectionsResponseValidationDKIM = "neutral"
	InvestigateDetectionsResponseValidationDKIMFail    InvestigateDetectionsResponseValidationDKIM = "fail"
	InvestigateDetectionsResponseValidationDKIMError   InvestigateDetectionsResponseValidationDKIM = "error"
	InvestigateDetectionsResponseValidationDKIMNone    InvestigateDetectionsResponseValidationDKIM = "none"
)

func (r InvestigateDetectionsResponseValidationDKIM) IsKnown() bool {
	switch r {
	case InvestigateDetectionsResponseValidationDKIMPass, InvestigateDetectionsResponseValidationDKIMNeutral, InvestigateDetectionsResponseValidationDKIMFail, InvestigateDetectionsResponseValidationDKIMError, InvestigateDetectionsResponseValidationDKIMNone:
		return true
	}
	return false
}

type InvestigateDetectionsResponseValidationDMARC string

const (
	InvestigateDetectionsResponseValidationDMARCPass    InvestigateDetectionsResponseValidationDMARC = "pass"
	InvestigateDetectionsResponseValidationDMARCNeutral InvestigateDetectionsResponseValidationDMARC = "neutral"
	InvestigateDetectionsResponseValidationDMARCFail    InvestigateDetectionsResponseValidationDMARC = "fail"
	InvestigateDetectionsResponseValidationDMARCError   InvestigateDetectionsResponseValidationDMARC = "error"
	InvestigateDetectionsResponseValidationDMARCNone    InvestigateDetectionsResponseValidationDMARC = "none"
)

func (r InvestigateDetectionsResponseValidationDMARC) IsKnown() bool {
	switch r {
	case InvestigateDetectionsResponseValidationDMARCPass, InvestigateDetectionsResponseValidationDMARCNeutral, InvestigateDetectionsResponseValidationDMARCFail, InvestigateDetectionsResponseValidationDMARCError, InvestigateDetectionsResponseValidationDMARCNone:
		return true
	}
	return false
}

type InvestigateDetectionsResponseValidationSPF string

const (
	InvestigateDetectionsResponseValidationSPFPass    InvestigateDetectionsResponseValidationSPF = "pass"
	InvestigateDetectionsResponseValidationSPFNeutral InvestigateDetectionsResponseValidationSPF = "neutral"
	InvestigateDetectionsResponseValidationSPFFail    InvestigateDetectionsResponseValidationSPF = "fail"
	InvestigateDetectionsResponseValidationSPFError   InvestigateDetectionsResponseValidationSPF = "error"
	InvestigateDetectionsResponseValidationSPFNone    InvestigateDetectionsResponseValidationSPF = "none"
)

func (r InvestigateDetectionsResponseValidationSPF) IsKnown() bool {
	switch r {
	case InvestigateDetectionsResponseValidationSPFPass, InvestigateDetectionsResponseValidationSPFNeutral, InvestigateDetectionsResponseValidationSPFFail, InvestigateDetectionsResponseValidationSPFError, InvestigateDetectionsResponseValidationSPFNone:
		return true
	}
	return false
}

type InvestigateDetectionsResponseFinalDisposition string

const (
	InvestigateDetectionsResponseFinalDispositionMalicious    InvestigateDetectionsResponseFinalDisposition = "MALICIOUS"
	InvestigateDetectionsResponseFinalDispositionMaliciousBec InvestigateDetectionsResponseFinalDisposition = "MALICIOUS-BEC"
	InvestigateDetectionsResponseFinalDispositionSuspicious   InvestigateDetectionsResponseFinalDisposition = "SUSPICIOUS"
	InvestigateDetectionsResponseFinalDispositionSpoof        InvestigateDetectionsResponseFinalDisposition = "SPOOF"
	InvestigateDetectionsResponseFinalDispositionSpam         InvestigateDetectionsResponseFinalDisposition = "SPAM"
	InvestigateDetectionsResponseFinalDispositionBulk         InvestigateDetectionsResponseFinalDisposition = "BULK"
	InvestigateDetectionsResponseFinalDispositionEncrypted    InvestigateDetectionsResponseFinalDisposition = "ENCRYPTED"
	InvestigateDetectionsResponseFinalDispositionExternal     InvestigateDetectionsResponseFinalDisposition = "EXTERNAL"
	InvestigateDetectionsResponseFinalDispositionUnknown      InvestigateDetectionsResponseFinalDisposition = "UNKNOWN"
	InvestigateDetectionsResponseFinalDispositionNone         InvestigateDetectionsResponseFinalDisposition = "NONE"
)

func (r InvestigateDetectionsResponseFinalDisposition) IsKnown() bool {
	switch r {
	case InvestigateDetectionsResponseFinalDispositionMalicious, InvestigateDetectionsResponseFinalDispositionMaliciousBec, InvestigateDetectionsResponseFinalDispositionSuspicious, InvestigateDetectionsResponseFinalDispositionSpoof, InvestigateDetectionsResponseFinalDispositionSpam, InvestigateDetectionsResponseFinalDispositionBulk, InvestigateDetectionsResponseFinalDispositionEncrypted, InvestigateDetectionsResponseFinalDispositionExternal, InvestigateDetectionsResponseFinalDispositionUnknown, InvestigateDetectionsResponseFinalDispositionNone:
		return true
	}
	return false
}

type InvestigateGetResponse struct {
	ID                string      `json:"id,required"`
	ActionLog         interface{} `json:"action_log,required"`
	ClientRecipients  []string    `json:"client_recipients,required"`
	DetectionReasons  []string    `json:"detection_reasons,required"`
	IsPhishSubmission bool        `json:"is_phish_submission,required"`
	IsQuarantined     bool        `json:"is_quarantined,required"`
	MessageID         string      `json:"message_id,required"`
	// Message identifier
	PostfixID        string                                 `json:"postfix_id,required"`
	Ts               string                                 `json:"ts,required"`
	AlertID          string                                 `json:"alert_id,nullable"`
	EdfHash          string                                 `json:"edf_hash,nullable"`
	FinalDisposition InvestigateGetResponseFinalDisposition `json:"final_disposition,nullable"`
	From             string                                 `json:"from,nullable"`
	FromName         string                                 `json:"from_name,nullable"`
	SentDate         string                                 `json:"sent_date,nullable"`
	Subject          string                                 `json:"subject,nullable"`
	ThreatCategories []string                               `json:"threat_categories,nullable"`
	To               []string                               `json:"to,nullable"`
	ToName           []string                               `json:"to_name,nullable"`
	Validation       InvestigateGetResponseValidation       `json:"validation,nullable"`
	JSON             investigateGetResponseJSON             `json:"-"`
}

// investigateGetResponseJSON contains the JSON metadata for the struct
// [InvestigateGetResponse]
type investigateGetResponseJSON struct {
	ID                apijson.Field
	ActionLog         apijson.Field
	ClientRecipients  apijson.Field
	DetectionReasons  apijson.Field
	IsPhishSubmission apijson.Field
	IsQuarantined     apijson.Field
	MessageID         apijson.Field
	PostfixID         apijson.Field
	Ts                apijson.Field
	AlertID           apijson.Field
	EdfHash           apijson.Field
	FinalDisposition  apijson.Field
	From              apijson.Field
	FromName          apijson.Field
	SentDate          apijson.Field
	Subject           apijson.Field
	ThreatCategories  apijson.Field
	To                apijson.Field
	ToName            apijson.Field
	Validation        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvestigateGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateGetResponseJSON) RawJSON() string {
	return r.raw
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

type InvestigateGetResponseValidation struct {
	Comment string                                `json:"comment,nullable"`
	DKIM    InvestigateGetResponseValidationDKIM  `json:"dkim,nullable"`
	DMARC   InvestigateGetResponseValidationDMARC `json:"dmarc,nullable"`
	SPF     InvestigateGetResponseValidationSPF   `json:"spf,nullable"`
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

type InvestigatePreviewResponse struct {
	// Base64 encoded PNG image
	Screenshot string                         `json:"screenshot,required"`
	JSON       investigatePreviewResponseJSON `json:"-"`
}

// investigatePreviewResponseJSON contains the JSON metadata for the struct
// [InvestigatePreviewResponse]
type investigatePreviewResponseJSON struct {
	Screenshot  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigatePreviewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigatePreviewResponseJSON) RawJSON() string {
	return r.raw
}

type InvestigateRawResponse struct {
	// UTF-8 encoded eml file
	Raw  string                     `json:"raw,required"`
	JSON investigateRawResponseJSON `json:"-"`
}

// investigateRawResponseJSON contains the JSON metadata for the struct
// [InvestigateRawResponse]
type investigateRawResponseJSON struct {
	Raw         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateRawResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateRawResponseJSON) RawJSON() string {
	return r.raw
}

type InvestigateTraceResponse struct {
	Inbound  InvestigateTraceResponseInbound  `json:"inbound,required"`
	Outbound InvestigateTraceResponseOutbound `json:"outbound,required"`
	JSON     investigateTraceResponseJSON     `json:"-"`
}

// investigateTraceResponseJSON contains the JSON metadata for the struct
// [InvestigateTraceResponse]
type investigateTraceResponseJSON struct {
	Inbound     apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateTraceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateTraceResponseJSON) RawJSON() string {
	return r.raw
}

type InvestigateTraceResponseInbound struct {
	Lines []InvestigateTraceResponseInboundLine `json:"lines,nullable"`
	JSON  investigateTraceResponseInboundJSON   `json:"-"`
}

// investigateTraceResponseInboundJSON contains the JSON metadata for the struct
// [InvestigateTraceResponseInbound]
type investigateTraceResponseInboundJSON struct {
	Lines       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateTraceResponseInbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateTraceResponseInboundJSON) RawJSON() string {
	return r.raw
}

type InvestigateTraceResponseInboundLine struct {
	Lineno  int64                                   `json:"lineno,required"`
	Message string                                  `json:"message,required"`
	Ts      time.Time                               `json:"ts,required" format:"date-time"`
	JSON    investigateTraceResponseInboundLineJSON `json:"-"`
}

// investigateTraceResponseInboundLineJSON contains the JSON metadata for the
// struct [InvestigateTraceResponseInboundLine]
type investigateTraceResponseInboundLineJSON struct {
	Lineno      apijson.Field
	Message     apijson.Field
	Ts          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateTraceResponseInboundLine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateTraceResponseInboundLineJSON) RawJSON() string {
	return r.raw
}

type InvestigateTraceResponseOutbound struct {
	Lines []InvestigateTraceResponseOutboundLine `json:"lines,nullable"`
	JSON  investigateTraceResponseOutboundJSON   `json:"-"`
}

// investigateTraceResponseOutboundJSON contains the JSON metadata for the struct
// [InvestigateTraceResponseOutbound]
type investigateTraceResponseOutboundJSON struct {
	Lines       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateTraceResponseOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateTraceResponseOutboundJSON) RawJSON() string {
	return r.raw
}

type InvestigateTraceResponseOutboundLine struct {
	Lineno  int64                                    `json:"lineno,required"`
	Message string                                   `json:"message,required"`
	Ts      time.Time                                `json:"ts,required" format:"date-time"`
	JSON    investigateTraceResponseOutboundLineJSON `json:"-"`
}

// investigateTraceResponseOutboundLineJSON contains the JSON metadata for the
// struct [InvestigateTraceResponseOutboundLine]
type investigateTraceResponseOutboundLineJSON struct {
	Lineno      apijson.Field
	Message     apijson.Field
	Ts          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateTraceResponseOutboundLine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateTraceResponseOutboundLineJSON) RawJSON() string {
	return r.raw
}

type InvestigateListParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Controls whether the message action log in included in the response.
	ActionLog param.Field[bool]   `query:"action_log"`
	AlertID   param.Field[string] `query:"alert_id"`
	// If `false`, the search includes non-detections.
	DetectionsOnly param.Field[bool] `query:"detections_only"`
	// Filter by the sender domain
	Domain param.Field[string] `query:"domain"`
	// The end of the search date range. Defaults to `now`.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// Filter messages by the provided disposition.
	FinalDisposition param.Field[InvestigateListParamsFinalDisposition] `query:"final_disposition"`
	// Filter messages by actions applied to them
	MessageAction param.Field[InvestigateListParamsMessageAction] `query:"message_action"`
	MessageID     param.Field[string]                             `query:"message_id"`
	Metric        param.Field[string]                             `query:"metric"`
	// Page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// Number of results to display.
	PerPage param.Field[int64] `query:"per_page"`
	// Space delimited query term(s). The search is case-insensitive.
	//
	// The content of the following email metadata fields are searched:
	//
	// - alert_id
	// - CC
	// - From (envelope_from)
	// - From Name
	// - final_disposition
	// - md5 hash (of any attachment)
	// - sha1 hash (of any attachment)
	// - sha256 hash (of any attachment)
	// - name (of any attachment)
	// - Reason
	// - Received DateTime (yyyy-mm-ddThh:mm:ss)
	// - Sent DateTime (yyyy-mm-ddThh:mm:ss)
	// - ReplyTo
	// - To (envelope_to)
	// - To Name
	// - Message-ID
	// - smtp_helo_server_ip
	// - smtp_previous_hop_ip
	// - x_originating_ip
	// - Subject
	Query     param.Field[string] `query:"query"`
	Recipient param.Field[string] `query:"recipient"`
	Sender    param.Field[string] `query:"sender"`
	// The beginning of the search date range. Defaults to `now - 30 days`.
	Start param.Field[time.Time] `query:"start" format:"date-time"`
}

// URLQuery serializes [InvestigateListParams]'s query parameters as `url.Values`.
func (r InvestigateListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter messages by the provided disposition.
type InvestigateListParamsFinalDisposition string

const (
	InvestigateListParamsFinalDispositionMalicious  InvestigateListParamsFinalDisposition = "MALICIOUS"
	InvestigateListParamsFinalDispositionSuspicious InvestigateListParamsFinalDisposition = "SUSPICIOUS"
	InvestigateListParamsFinalDispositionSpoof      InvestigateListParamsFinalDisposition = "SPOOF"
	InvestigateListParamsFinalDispositionSpam       InvestigateListParamsFinalDisposition = "SPAM"
	InvestigateListParamsFinalDispositionBulk       InvestigateListParamsFinalDisposition = "BULK"
)

func (r InvestigateListParamsFinalDisposition) IsKnown() bool {
	switch r {
	case InvestigateListParamsFinalDispositionMalicious, InvestigateListParamsFinalDispositionSuspicious, InvestigateListParamsFinalDispositionSpoof, InvestigateListParamsFinalDispositionSpam, InvestigateListParamsFinalDispositionBulk:
		return true
	}
	return false
}

// Filter messages by actions applied to them
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

type InvestigateDetectionsParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type InvestigateDetectionsResponseEnvelope struct {
	Errors   []shared.ResponseInfo                     `json:"errors,required"`
	Messages []shared.ResponseInfo                     `json:"messages,required"`
	Result   InvestigateDetectionsResponse             `json:"result,required"`
	Success  bool                                      `json:"success,required"`
	JSON     investigateDetectionsResponseEnvelopeJSON `json:"-"`
}

// investigateDetectionsResponseEnvelopeJSON contains the JSON metadata for the
// struct [InvestigateDetectionsResponseEnvelope]
type investigateDetectionsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateDetectionsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateDetectionsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InvestigateGetParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type InvestigateGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo              `json:"errors,required"`
	Messages []shared.ResponseInfo              `json:"messages,required"`
	Result   InvestigateGetResponse             `json:"result,required"`
	Success  bool                               `json:"success,required"`
	JSON     investigateGetResponseEnvelopeJSON `json:"-"`
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

type InvestigatePreviewParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type InvestigatePreviewResponseEnvelope struct {
	Errors   []shared.ResponseInfo                  `json:"errors,required"`
	Messages []shared.ResponseInfo                  `json:"messages,required"`
	Result   InvestigatePreviewResponse             `json:"result,required"`
	Success  bool                                   `json:"success,required"`
	JSON     investigatePreviewResponseEnvelopeJSON `json:"-"`
}

// investigatePreviewResponseEnvelopeJSON contains the JSON metadata for the struct
// [InvestigatePreviewResponseEnvelope]
type investigatePreviewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigatePreviewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigatePreviewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InvestigateRawParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type InvestigateRawResponseEnvelope struct {
	Errors   []shared.ResponseInfo              `json:"errors,required"`
	Messages []shared.ResponseInfo              `json:"messages,required"`
	Result   InvestigateRawResponse             `json:"result,required"`
	Success  bool                               `json:"success,required"`
	JSON     investigateRawResponseEnvelopeJSON `json:"-"`
}

// investigateRawResponseEnvelopeJSON contains the JSON metadata for the struct
// [InvestigateRawResponseEnvelope]
type investigateRawResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateRawResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateRawResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InvestigateTraceParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type InvestigateTraceResponseEnvelope struct {
	Errors   []shared.ResponseInfo                `json:"errors,required"`
	Messages []shared.ResponseInfo                `json:"messages,required"`
	Result   InvestigateTraceResponse             `json:"result,required"`
	Success  bool                                 `json:"success,required"`
	JSON     investigateTraceResponseEnvelopeJSON `json:"-"`
}

// investigateTraceResponseEnvelopeJSON contains the JSON metadata for the struct
// [InvestigateTraceResponseEnvelope]
type investigateTraceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateTraceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateTraceResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
