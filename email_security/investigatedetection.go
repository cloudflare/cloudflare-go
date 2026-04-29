// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// InvestigateDetectionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvestigateDetectionService] method instead.
type InvestigateDetectionService struct {
	Options []option.RequestOption
}

// NewInvestigateDetectionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInvestigateDetectionService(opts ...option.RequestOption) (r *InvestigateDetectionService) {
	r = &InvestigateDetectionService{}
	r.Options = opts
	return
}

// Returns detection details such as threat categories and sender information for
// non-benign messages.
func (r *InvestigateDetectionService) Get(ctx context.Context, investigateID string, query InvestigateDetectionGetParams, opts ...option.RequestOption) (res *InvestigateDetectionGetResponse, err error) {
	var env InvestigateDetectionGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if investigateID == "" {
		err = errors.New("missing required investigate_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/%s/detections", query.AccountID, investigateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type InvestigateDetectionGetResponse struct {
	Action           string                                          `json:"action" api:"required"`
	Attachments      []InvestigateDetectionGetResponseAttachment     `json:"attachments" api:"required"`
	Findings         []InvestigateDetectionGetResponseFinding        `json:"findings" api:"required,nullable"`
	Headers          []InvestigateDetectionGetResponseHeader         `json:"headers" api:"required"`
	Links            []InvestigateDetectionGetResponseLink           `json:"links" api:"required"`
	SenderInfo       InvestigateDetectionGetResponseSenderInfo       `json:"sender_info" api:"required"`
	ThreatCategories []InvestigateDetectionGetResponseThreatCategory `json:"threat_categories" api:"required"`
	Validation       InvestigateDetectionGetResponseValidation       `json:"validation" api:"required"`
	FinalDisposition InvestigateDetectionGetResponseFinalDisposition `json:"final_disposition"`
	JSON             investigateDetectionGetResponseJSON             `json:"-"`
}

// investigateDetectionGetResponseJSON contains the JSON metadata for the struct
// [InvestigateDetectionGetResponse]
type investigateDetectionGetResponseJSON struct {
	Action           apijson.Field
	Attachments      apijson.Field
	Findings         apijson.Field
	Headers          apijson.Field
	Links            apijson.Field
	SenderInfo       apijson.Field
	ThreatCategories apijson.Field
	Validation       apijson.Field
	FinalDisposition apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvestigateDetectionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateDetectionGetResponseJSON) RawJSON() string {
	return r.raw
}

type InvestigateDetectionGetResponseAttachment struct {
	// Size of the attachment in bytes
	Size int64 `json:"size" api:"required"`
	// MIME type of the attachment
	ContentType string `json:"content_type" api:"nullable"`
	// Detection result for this attachment
	Detection InvestigateDetectionGetResponseAttachmentsDetection `json:"detection" api:"nullable"`
	// Whether the attachment is encrypted
	Encrypted bool `json:"encrypted" api:"nullable"`
	// Name of the attached file
	Filename string `json:"filename" api:"nullable"`
	// MD5 hash of the attachment
	Md5 string `json:"md5" api:"nullable"`
	// Attachment name (alternative to filename)
	Name string `json:"name" api:"nullable"`
	// SHA1 hash of the attachment
	Sha1 string `json:"sha1" api:"nullable"`
	// SHA256 hash of the attachment
	Sha256 string                                        `json:"sha256" api:"nullable"`
	JSON   investigateDetectionGetResponseAttachmentJSON `json:"-"`
}

// investigateDetectionGetResponseAttachmentJSON contains the JSON metadata for the
// struct [InvestigateDetectionGetResponseAttachment]
type investigateDetectionGetResponseAttachmentJSON struct {
	Size        apijson.Field
	ContentType apijson.Field
	Detection   apijson.Field
	Encrypted   apijson.Field
	Filename    apijson.Field
	Md5         apijson.Field
	Name        apijson.Field
	Sha1        apijson.Field
	Sha256      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateDetectionGetResponseAttachment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateDetectionGetResponseAttachmentJSON) RawJSON() string {
	return r.raw
}

// Detection result for this attachment
type InvestigateDetectionGetResponseAttachmentsDetection string

const (
	InvestigateDetectionGetResponseAttachmentsDetectionMalicious    InvestigateDetectionGetResponseAttachmentsDetection = "MALICIOUS"
	InvestigateDetectionGetResponseAttachmentsDetectionMaliciousBec InvestigateDetectionGetResponseAttachmentsDetection = "MALICIOUS-BEC"
	InvestigateDetectionGetResponseAttachmentsDetectionSuspicious   InvestigateDetectionGetResponseAttachmentsDetection = "SUSPICIOUS"
	InvestigateDetectionGetResponseAttachmentsDetectionSpoof        InvestigateDetectionGetResponseAttachmentsDetection = "SPOOF"
	InvestigateDetectionGetResponseAttachmentsDetectionSpam         InvestigateDetectionGetResponseAttachmentsDetection = "SPAM"
	InvestigateDetectionGetResponseAttachmentsDetectionBulk         InvestigateDetectionGetResponseAttachmentsDetection = "BULK"
	InvestigateDetectionGetResponseAttachmentsDetectionEncrypted    InvestigateDetectionGetResponseAttachmentsDetection = "ENCRYPTED"
	InvestigateDetectionGetResponseAttachmentsDetectionExternal     InvestigateDetectionGetResponseAttachmentsDetection = "EXTERNAL"
	InvestigateDetectionGetResponseAttachmentsDetectionUnknown      InvestigateDetectionGetResponseAttachmentsDetection = "UNKNOWN"
	InvestigateDetectionGetResponseAttachmentsDetectionNone         InvestigateDetectionGetResponseAttachmentsDetection = "NONE"
)

func (r InvestigateDetectionGetResponseAttachmentsDetection) IsKnown() bool {
	switch r {
	case InvestigateDetectionGetResponseAttachmentsDetectionMalicious, InvestigateDetectionGetResponseAttachmentsDetectionMaliciousBec, InvestigateDetectionGetResponseAttachmentsDetectionSuspicious, InvestigateDetectionGetResponseAttachmentsDetectionSpoof, InvestigateDetectionGetResponseAttachmentsDetectionSpam, InvestigateDetectionGetResponseAttachmentsDetectionBulk, InvestigateDetectionGetResponseAttachmentsDetectionEncrypted, InvestigateDetectionGetResponseAttachmentsDetectionExternal, InvestigateDetectionGetResponseAttachmentsDetectionUnknown, InvestigateDetectionGetResponseAttachmentsDetectionNone:
		return true
	}
	return false
}

type InvestigateDetectionGetResponseFinding struct {
	Attachment string                                           `json:"attachment" api:"nullable"`
	Detail     string                                           `json:"detail" api:"nullable"`
	Detection  InvestigateDetectionGetResponseFindingsDetection `json:"detection"`
	Field      string                                           `json:"field" api:"nullable"`
	Name       string                                           `json:"name" api:"nullable"`
	Portion    string                                           `json:"portion" api:"nullable"`
	Reason     string                                           `json:"reason" api:"nullable"`
	Score      float64                                          `json:"score" api:"nullable"`
	Value      string                                           `json:"value" api:"nullable"`
	JSON       investigateDetectionGetResponseFindingJSON       `json:"-"`
}

// investigateDetectionGetResponseFindingJSON contains the JSON metadata for the
// struct [InvestigateDetectionGetResponseFinding]
type investigateDetectionGetResponseFindingJSON struct {
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

func (r *InvestigateDetectionGetResponseFinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateDetectionGetResponseFindingJSON) RawJSON() string {
	return r.raw
}

type InvestigateDetectionGetResponseFindingsDetection string

const (
	InvestigateDetectionGetResponseFindingsDetectionMalicious    InvestigateDetectionGetResponseFindingsDetection = "MALICIOUS"
	InvestigateDetectionGetResponseFindingsDetectionMaliciousBec InvestigateDetectionGetResponseFindingsDetection = "MALICIOUS-BEC"
	InvestigateDetectionGetResponseFindingsDetectionSuspicious   InvestigateDetectionGetResponseFindingsDetection = "SUSPICIOUS"
	InvestigateDetectionGetResponseFindingsDetectionSpoof        InvestigateDetectionGetResponseFindingsDetection = "SPOOF"
	InvestigateDetectionGetResponseFindingsDetectionSpam         InvestigateDetectionGetResponseFindingsDetection = "SPAM"
	InvestigateDetectionGetResponseFindingsDetectionBulk         InvestigateDetectionGetResponseFindingsDetection = "BULK"
	InvestigateDetectionGetResponseFindingsDetectionEncrypted    InvestigateDetectionGetResponseFindingsDetection = "ENCRYPTED"
	InvestigateDetectionGetResponseFindingsDetectionExternal     InvestigateDetectionGetResponseFindingsDetection = "EXTERNAL"
	InvestigateDetectionGetResponseFindingsDetectionUnknown      InvestigateDetectionGetResponseFindingsDetection = "UNKNOWN"
	InvestigateDetectionGetResponseFindingsDetectionNone         InvestigateDetectionGetResponseFindingsDetection = "NONE"
)

func (r InvestigateDetectionGetResponseFindingsDetection) IsKnown() bool {
	switch r {
	case InvestigateDetectionGetResponseFindingsDetectionMalicious, InvestigateDetectionGetResponseFindingsDetectionMaliciousBec, InvestigateDetectionGetResponseFindingsDetectionSuspicious, InvestigateDetectionGetResponseFindingsDetectionSpoof, InvestigateDetectionGetResponseFindingsDetectionSpam, InvestigateDetectionGetResponseFindingsDetectionBulk, InvestigateDetectionGetResponseFindingsDetectionEncrypted, InvestigateDetectionGetResponseFindingsDetectionExternal, InvestigateDetectionGetResponseFindingsDetectionUnknown, InvestigateDetectionGetResponseFindingsDetectionNone:
		return true
	}
	return false
}

type InvestigateDetectionGetResponseHeader struct {
	Name  string                                    `json:"name" api:"required"`
	Value string                                    `json:"value" api:"required"`
	JSON  investigateDetectionGetResponseHeaderJSON `json:"-"`
}

// investigateDetectionGetResponseHeaderJSON contains the JSON metadata for the
// struct [InvestigateDetectionGetResponseHeader]
type investigateDetectionGetResponseHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateDetectionGetResponseHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateDetectionGetResponseHeaderJSON) RawJSON() string {
	return r.raw
}

type InvestigateDetectionGetResponseLink struct {
	Href string                                  `json:"href" api:"required"`
	Text string                                  `json:"text" api:"nullable"`
	JSON investigateDetectionGetResponseLinkJSON `json:"-"`
}

// investigateDetectionGetResponseLinkJSON contains the JSON metadata for the
// struct [InvestigateDetectionGetResponseLink]
type investigateDetectionGetResponseLinkJSON struct {
	Href        apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateDetectionGetResponseLink) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateDetectionGetResponseLinkJSON) RawJSON() string {
	return r.raw
}

type InvestigateDetectionGetResponseSenderInfo struct {
	// The name of the autonomous system.
	AsName string `json:"as_name" api:"nullable"`
	// The number of the autonomous system.
	AsNumber int64                                         `json:"as_number" api:"nullable"`
	Geo      string                                        `json:"geo" api:"nullable"`
	IP       string                                        `json:"ip" api:"nullable"`
	Pld      string                                        `json:"pld" api:"nullable"`
	JSON     investigateDetectionGetResponseSenderInfoJSON `json:"-"`
}

// investigateDetectionGetResponseSenderInfoJSON contains the JSON metadata for the
// struct [InvestigateDetectionGetResponseSenderInfo]
type investigateDetectionGetResponseSenderInfoJSON struct {
	AsName      apijson.Field
	AsNumber    apijson.Field
	Geo         apijson.Field
	IP          apijson.Field
	Pld         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateDetectionGetResponseSenderInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateDetectionGetResponseSenderInfoJSON) RawJSON() string {
	return r.raw
}

type InvestigateDetectionGetResponseThreatCategory struct {
	ID          int64                                             `json:"id"`
	Description string                                            `json:"description" api:"nullable"`
	Name        string                                            `json:"name" api:"nullable"`
	JSON        investigateDetectionGetResponseThreatCategoryJSON `json:"-"`
}

// investigateDetectionGetResponseThreatCategoryJSON contains the JSON metadata for
// the struct [InvestigateDetectionGetResponseThreatCategory]
type investigateDetectionGetResponseThreatCategoryJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateDetectionGetResponseThreatCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateDetectionGetResponseThreatCategoryJSON) RawJSON() string {
	return r.raw
}

type InvestigateDetectionGetResponseValidation struct {
	Comment string                                         `json:"comment" api:"nullable"`
	DKIM    InvestigateDetectionGetResponseValidationDKIM  `json:"dkim"`
	DMARC   InvestigateDetectionGetResponseValidationDMARC `json:"dmarc"`
	SPF     InvestigateDetectionGetResponseValidationSPF   `json:"spf"`
	JSON    investigateDetectionGetResponseValidationJSON  `json:"-"`
}

// investigateDetectionGetResponseValidationJSON contains the JSON metadata for the
// struct [InvestigateDetectionGetResponseValidation]
type investigateDetectionGetResponseValidationJSON struct {
	Comment     apijson.Field
	DKIM        apijson.Field
	DMARC       apijson.Field
	SPF         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateDetectionGetResponseValidation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateDetectionGetResponseValidationJSON) RawJSON() string {
	return r.raw
}

type InvestigateDetectionGetResponseValidationDKIM string

const (
	InvestigateDetectionGetResponseValidationDKIMPass    InvestigateDetectionGetResponseValidationDKIM = "pass"
	InvestigateDetectionGetResponseValidationDKIMNeutral InvestigateDetectionGetResponseValidationDKIM = "neutral"
	InvestigateDetectionGetResponseValidationDKIMFail    InvestigateDetectionGetResponseValidationDKIM = "fail"
	InvestigateDetectionGetResponseValidationDKIMError   InvestigateDetectionGetResponseValidationDKIM = "error"
	InvestigateDetectionGetResponseValidationDKIMNone    InvestigateDetectionGetResponseValidationDKIM = "none"
)

func (r InvestigateDetectionGetResponseValidationDKIM) IsKnown() bool {
	switch r {
	case InvestigateDetectionGetResponseValidationDKIMPass, InvestigateDetectionGetResponseValidationDKIMNeutral, InvestigateDetectionGetResponseValidationDKIMFail, InvestigateDetectionGetResponseValidationDKIMError, InvestigateDetectionGetResponseValidationDKIMNone:
		return true
	}
	return false
}

type InvestigateDetectionGetResponseValidationDMARC string

const (
	InvestigateDetectionGetResponseValidationDMARCPass    InvestigateDetectionGetResponseValidationDMARC = "pass"
	InvestigateDetectionGetResponseValidationDMARCNeutral InvestigateDetectionGetResponseValidationDMARC = "neutral"
	InvestigateDetectionGetResponseValidationDMARCFail    InvestigateDetectionGetResponseValidationDMARC = "fail"
	InvestigateDetectionGetResponseValidationDMARCError   InvestigateDetectionGetResponseValidationDMARC = "error"
	InvestigateDetectionGetResponseValidationDMARCNone    InvestigateDetectionGetResponseValidationDMARC = "none"
)

func (r InvestigateDetectionGetResponseValidationDMARC) IsKnown() bool {
	switch r {
	case InvestigateDetectionGetResponseValidationDMARCPass, InvestigateDetectionGetResponseValidationDMARCNeutral, InvestigateDetectionGetResponseValidationDMARCFail, InvestigateDetectionGetResponseValidationDMARCError, InvestigateDetectionGetResponseValidationDMARCNone:
		return true
	}
	return false
}

type InvestigateDetectionGetResponseValidationSPF string

const (
	InvestigateDetectionGetResponseValidationSPFPass    InvestigateDetectionGetResponseValidationSPF = "pass"
	InvestigateDetectionGetResponseValidationSPFNeutral InvestigateDetectionGetResponseValidationSPF = "neutral"
	InvestigateDetectionGetResponseValidationSPFFail    InvestigateDetectionGetResponseValidationSPF = "fail"
	InvestigateDetectionGetResponseValidationSPFError   InvestigateDetectionGetResponseValidationSPF = "error"
	InvestigateDetectionGetResponseValidationSPFNone    InvestigateDetectionGetResponseValidationSPF = "none"
)

func (r InvestigateDetectionGetResponseValidationSPF) IsKnown() bool {
	switch r {
	case InvestigateDetectionGetResponseValidationSPFPass, InvestigateDetectionGetResponseValidationSPFNeutral, InvestigateDetectionGetResponseValidationSPFFail, InvestigateDetectionGetResponseValidationSPFError, InvestigateDetectionGetResponseValidationSPFNone:
		return true
	}
	return false
}

type InvestigateDetectionGetResponseFinalDisposition string

const (
	InvestigateDetectionGetResponseFinalDispositionMalicious    InvestigateDetectionGetResponseFinalDisposition = "MALICIOUS"
	InvestigateDetectionGetResponseFinalDispositionMaliciousBec InvestigateDetectionGetResponseFinalDisposition = "MALICIOUS-BEC"
	InvestigateDetectionGetResponseFinalDispositionSuspicious   InvestigateDetectionGetResponseFinalDisposition = "SUSPICIOUS"
	InvestigateDetectionGetResponseFinalDispositionSpoof        InvestigateDetectionGetResponseFinalDisposition = "SPOOF"
	InvestigateDetectionGetResponseFinalDispositionSpam         InvestigateDetectionGetResponseFinalDisposition = "SPAM"
	InvestigateDetectionGetResponseFinalDispositionBulk         InvestigateDetectionGetResponseFinalDisposition = "BULK"
	InvestigateDetectionGetResponseFinalDispositionEncrypted    InvestigateDetectionGetResponseFinalDisposition = "ENCRYPTED"
	InvestigateDetectionGetResponseFinalDispositionExternal     InvestigateDetectionGetResponseFinalDisposition = "EXTERNAL"
	InvestigateDetectionGetResponseFinalDispositionUnknown      InvestigateDetectionGetResponseFinalDisposition = "UNKNOWN"
	InvestigateDetectionGetResponseFinalDispositionNone         InvestigateDetectionGetResponseFinalDisposition = "NONE"
)

func (r InvestigateDetectionGetResponseFinalDisposition) IsKnown() bool {
	switch r {
	case InvestigateDetectionGetResponseFinalDispositionMalicious, InvestigateDetectionGetResponseFinalDispositionMaliciousBec, InvestigateDetectionGetResponseFinalDispositionSuspicious, InvestigateDetectionGetResponseFinalDispositionSpoof, InvestigateDetectionGetResponseFinalDispositionSpam, InvestigateDetectionGetResponseFinalDispositionBulk, InvestigateDetectionGetResponseFinalDispositionEncrypted, InvestigateDetectionGetResponseFinalDispositionExternal, InvestigateDetectionGetResponseFinalDispositionUnknown, InvestigateDetectionGetResponseFinalDispositionNone:
		return true
	}
	return false
}

type InvestigateDetectionGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type InvestigateDetectionGetResponseEnvelope struct {
	Errors   []InvestigateDetectionGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []InvestigateDetectionGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   InvestigateDetectionGetResponse                   `json:"result" api:"required"`
	// Whether the API call was successful.
	Success InvestigateDetectionGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    investigateDetectionGetResponseEnvelopeJSON    `json:"-"`
}

// investigateDetectionGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [InvestigateDetectionGetResponseEnvelope]
type investigateDetectionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateDetectionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateDetectionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InvestigateDetectionGetResponseEnvelopeErrors struct {
	Code             int64                                               `json:"code" api:"required"`
	Message          string                                              `json:"message" api:"required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           InvestigateDetectionGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             investigateDetectionGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// investigateDetectionGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [InvestigateDetectionGetResponseEnvelopeErrors]
type investigateDetectionGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvestigateDetectionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateDetectionGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type InvestigateDetectionGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    investigateDetectionGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// investigateDetectionGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [InvestigateDetectionGetResponseEnvelopeErrorsSource]
type investigateDetectionGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateDetectionGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateDetectionGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type InvestigateDetectionGetResponseEnvelopeMessages struct {
	Code             int64                                                 `json:"code" api:"required"`
	Message          string                                                `json:"message" api:"required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           InvestigateDetectionGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             investigateDetectionGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// investigateDetectionGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [InvestigateDetectionGetResponseEnvelopeMessages]
type investigateDetectionGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvestigateDetectionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateDetectionGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type InvestigateDetectionGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    investigateDetectionGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// investigateDetectionGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [InvestigateDetectionGetResponseEnvelopeMessagesSource]
type investigateDetectionGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateDetectionGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateDetectionGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type InvestigateDetectionGetResponseEnvelopeSuccess bool

const (
	InvestigateDetectionGetResponseEnvelopeSuccessTrue InvestigateDetectionGetResponseEnvelopeSuccess = true
)

func (r InvestigateDetectionGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case InvestigateDetectionGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
