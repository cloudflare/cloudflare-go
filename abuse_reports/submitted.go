// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package abuse_reports

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
	"github.com/cloudflare/cloudflare-go/v7/shared"
	"github.com/tidwall/gjson"
)

// SubmittedService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubmittedService] method instead.
type SubmittedService struct {
	Options []option.RequestOption
	Emails  *SubmittedEmailService
}

// NewSubmittedService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSubmittedService(opts ...option.RequestOption) (r *SubmittedService) {
	r = &SubmittedService{}
	r.Options = opts
	r.Emails = NewSubmittedEmailService(opts...)
	return
}

// List abuse reports submitted by the account.
func (r *SubmittedService) List(ctx context.Context, params SubmittedListParams, opts ...option.RequestOption) (res *pagination.V4PagePagination[SubmittedListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/abuse-reports/submitted", params.AccountID)
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

// List abuse reports submitted by the account.
func (r *SubmittedService) ListAutoPaging(ctx context.Context, params SubmittedListParams, opts ...option.RequestOption) *pagination.V4PagePaginationAutoPager[SubmittedListResponse] {
	return pagination.NewV4PagePaginationAutoPager(r.List(ctx, params, opts...))
}

// Retrieve a report submitted by the account.
func (r *SubmittedService) Get(ctx context.Context, reportID string, query SubmittedGetParams, opts ...option.RequestOption) (res *SubmittedGetResponse, err error) {
	var env SubmittedGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if reportID == "" {
		err = errors.New("missing required report_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/abuse-reports/submitted/%s", query.AccountID, reportID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type SubmittedListResponse struct {
	Reports []SubmittedListResponseReport `json:"reports" api:"required"`
	JSON    submittedListResponseJSON     `json:"-"`
}

// submittedListResponseJSON contains the JSON metadata for the struct
// [SubmittedListResponse]
type submittedListResponseJSON struct {
	Reports     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubmittedListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r submittedListResponseJSON) RawJSON() string {
	return r.raw
}

type SubmittedListResponseReport struct {
	// Public report code.
	ID string `json:"id" api:"required"`
	// Time the report was submitted.
	Cdate time.Time `json:"cdate" api:"required" format:"date-time"`
	// Submitter-safe reason for a denied report. Null when unavailable.
	DenialReason SubmittedListResponseReportsDenialReason `json:"denial_reason" api:"required,nullable"`
	// Domain identified in the report.
	Domain string `json:"domain" api:"required"`
	// Status visible to the account that submitted the report.
	Status SubmittedListResponseReportsStatus `json:"status" api:"required"`
	// The abuse report type
	Type SubmittedListResponseReportsType `json:"type" api:"required"`
	// Information about the submitter of the report.
	Submitter SubmittedListResponseReportsSubmitter `json:"submitter"`
	JSON      submittedListResponseReportJSON       `json:"-"`
}

// submittedListResponseReportJSON contains the JSON metadata for the struct
// [SubmittedListResponseReport]
type submittedListResponseReportJSON struct {
	ID           apijson.Field
	Cdate        apijson.Field
	DenialReason apijson.Field
	Domain       apijson.Field
	Status       apijson.Field
	Type         apijson.Field
	Submitter    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SubmittedListResponseReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r submittedListResponseReportJSON) RawJSON() string {
	return r.raw
}

// Submitter-safe reason for a denied report. Null when unavailable.
type SubmittedListResponseReportsDenialReason string

const (
	SubmittedListResponseReportsDenialReasonUnableToConfirm                SubmittedListResponseReportsDenialReason = "unable_to_confirm"
	SubmittedListResponseReportsDenialReasonIncompleteReport               SubmittedListResponseReportsDenialReason = "incomplete_report"
	SubmittedListResponseReportsDenialReasonNotOnCloudflare                SubmittedListResponseReportsDenialReason = "not_on_cloudflare"
	SubmittedListResponseReportsDenialReasonDuplicateReport                SubmittedListResponseReportsDenialReason = "duplicate_report"
	SubmittedListResponseReportsDenialReasonContentRemoved                 SubmittedListResponseReportsDenialReason = "content_removed"
	SubmittedListResponseReportsDenialReasonReportDetailsMismatch          SubmittedListResponseReportsDenialReason = "report_details_mismatch"
	SubmittedListResponseReportsDenialReasonNoAbuseFound                   SubmittedListResponseReportsDenialReason = "no_abuse_found"
	SubmittedListResponseReportsDenialReasonMissingOriginalWork            SubmittedListResponseReportsDenialReason = "missing_original_work"
	SubmittedListResponseReportsDenialReasonDirectURLRequired              SubmittedListResponseReportsDenialReason = "direct_url_required"
	SubmittedListResponseReportsDenialReasonWrongReportCategory            SubmittedListResponseReportsDenialReason = "wrong_report_category"
	SubmittedListResponseReportsDenialReasonContentUnavailable             SubmittedListResponseReportsDenialReason = "content_unavailable"
	SubmittedListResponseReportsDenialReasonLawEnforcementReferralRequired SubmittedListResponseReportsDenialReason = "law_enforcement_referral_required"
	SubmittedListResponseReportsDenialReasonDomainDisputeProcessRequired   SubmittedListResponseReportsDenialReason = "domain_dispute_process_required"
)

func (r SubmittedListResponseReportsDenialReason) IsKnown() bool {
	switch r {
	case SubmittedListResponseReportsDenialReasonUnableToConfirm, SubmittedListResponseReportsDenialReasonIncompleteReport, SubmittedListResponseReportsDenialReasonNotOnCloudflare, SubmittedListResponseReportsDenialReasonDuplicateReport, SubmittedListResponseReportsDenialReasonContentRemoved, SubmittedListResponseReportsDenialReasonReportDetailsMismatch, SubmittedListResponseReportsDenialReasonNoAbuseFound, SubmittedListResponseReportsDenialReasonMissingOriginalWork, SubmittedListResponseReportsDenialReasonDirectURLRequired, SubmittedListResponseReportsDenialReasonWrongReportCategory, SubmittedListResponseReportsDenialReasonContentUnavailable, SubmittedListResponseReportsDenialReasonLawEnforcementReferralRequired, SubmittedListResponseReportsDenialReasonDomainDisputeProcessRequired:
		return true
	}
	return false
}

// Status visible to the account that submitted the report.
type SubmittedListResponseReportsStatus string

const (
	SubmittedListResponseReportsStatusSubmitted SubmittedListResponseReportsStatus = "submitted"
	SubmittedListResponseReportsStatusAccepted  SubmittedListResponseReportsStatus = "accepted"
	SubmittedListResponseReportsStatusDenied    SubmittedListResponseReportsStatus = "denied"
)

func (r SubmittedListResponseReportsStatus) IsKnown() bool {
	switch r {
	case SubmittedListResponseReportsStatusSubmitted, SubmittedListResponseReportsStatusAccepted, SubmittedListResponseReportsStatusDenied:
		return true
	}
	return false
}

// The abuse report type
type SubmittedListResponseReportsType string

const (
	SubmittedListResponseReportsTypePhish   SubmittedListResponseReportsType = "PHISH"
	SubmittedListResponseReportsTypeGen     SubmittedListResponseReportsType = "GEN"
	SubmittedListResponseReportsTypeThreat  SubmittedListResponseReportsType = "THREAT"
	SubmittedListResponseReportsTypeDmca    SubmittedListResponseReportsType = "DMCA"
	SubmittedListResponseReportsTypeEmer    SubmittedListResponseReportsType = "EMER"
	SubmittedListResponseReportsTypeTm      SubmittedListResponseReportsType = "TM"
	SubmittedListResponseReportsTypeRegWho  SubmittedListResponseReportsType = "REG_WHO"
	SubmittedListResponseReportsTypeNcsei   SubmittedListResponseReportsType = "NCSEI"
	SubmittedListResponseReportsTypeNetwork SubmittedListResponseReportsType = "NETWORK"
)

func (r SubmittedListResponseReportsType) IsKnown() bool {
	switch r {
	case SubmittedListResponseReportsTypePhish, SubmittedListResponseReportsTypeGen, SubmittedListResponseReportsTypeThreat, SubmittedListResponseReportsTypeDmca, SubmittedListResponseReportsTypeEmer, SubmittedListResponseReportsTypeTm, SubmittedListResponseReportsTypeRegWho, SubmittedListResponseReportsTypeNcsei, SubmittedListResponseReportsTypeNetwork:
		return true
	}
	return false
}

// Information about the submitter of the report.
type SubmittedListResponseReportsSubmitter struct {
	Company   string                                    `json:"company"`
	Email     string                                    `json:"email"`
	Name      string                                    `json:"name"`
	Telephone string                                    `json:"telephone"`
	JSON      submittedListResponseReportsSubmitterJSON `json:"-"`
}

// submittedListResponseReportsSubmitterJSON contains the JSON metadata for the
// struct [SubmittedListResponseReportsSubmitter]
type submittedListResponseReportsSubmitterJSON struct {
	Company     apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	Telephone   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubmittedListResponseReportsSubmitter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r submittedListResponseReportsSubmitterJSON) RawJSON() string {
	return r.raw
}

type SubmittedGetResponse struct {
	// Public report code.
	ID string `json:"id" api:"required"`
	// Time the report was submitted.
	Cdate time.Time `json:"cdate" api:"required" format:"date-time"`
	// Submitter-safe reason for a denied report. Null when unavailable.
	DenialReason SubmittedGetResponseDenialReason `json:"denial_reason" api:"required,nullable"`
	// Domain identified in the report.
	Domain string `json:"domain" api:"required"`
	// Whether the submitter provided the Digital Services Act attestation.
	DsaAttestation bool `json:"dsa_attestation" api:"required"`
	// Status visible to the account that submitted the report.
	Status SubmittedGetResponseStatus `json:"status" api:"required"`
	// The abuse report type
	Type SubmittedGetResponseType `json:"type" api:"required"`
	// URLs supplied with the report.
	URLs []string `json:"urls" api:"required"`
	// Authorized agent name supplied with the report.
	AgentName string `json:"agent_name"`
	// Additional comments supplied with the report.
	Comments string `json:"comments"`
	// The string "on" when a court proceeding applies to the report; otherwise
	// omitted.
	Court string `json:"court"`
	// Destination IP addresses supplied with a network abuse report.
	DestinationIPs []string `json:"destination_ips"`
	// Submitter preference for notifying the hosting provider.
	HostNotification string `json:"host_notification"`
	// Evidence supplied with the report.
	Justification string `json:"justification"`
	// Submitter preference for notifying NCMEC.
	NcmecNotification string `json:"ncmec_notification"`
	// Representation supplied for an NCSEI report.
	NcseiSubjectRepresentation bool `json:"ncsei_subject_representation"`
	// Original work or targeted brand supplied with the report.
	OriginalWork string `json:"original_work"`
	// Submitter preference for notifying the content owner.
	OwnerNotification string `json:"owner_notification"`
	// Ports and protocols supplied with a network abuse report.
	PortsProtocols []string `json:"ports_protocols"`
	// RDP-mandated fields for registrar WHOIS data disclosure requests.
	RegWhoRequest SubmittedGetResponseRegWhoRequest `json:"reg_who_request"`
	// Country associated with the reported activity.
	ReportedCountry string `json:"reported_country"`
	// User agent associated with the reported activity.
	ReportedUserAgent string `json:"reported_user_agent"`
	// Source IP addresses supplied with a network abuse report.
	SourceIPs []string `json:"source_ips"`
	// Information about the submitter of the report.
	Submitter SubmittedGetResponseSubmitter `json:"submitter"`
	// Additional abuse classifications supplied with the report.
	Subtypes []string `json:"subtypes"`
	// Title supplied with the report.
	Title string `json:"title"`
	// The string "on" when a UDRP proceeding applies to the report; otherwise omitted.
	Udrp string `json:"udrp"`
	// The string "on" when a URS proceeding applies to the report; otherwise omitted.
	Urs  string                   `json:"urs"`
	JSON submittedGetResponseJSON `json:"-"`
}

// submittedGetResponseJSON contains the JSON metadata for the struct
// [SubmittedGetResponse]
type submittedGetResponseJSON struct {
	ID                         apijson.Field
	Cdate                      apijson.Field
	DenialReason               apijson.Field
	Domain                     apijson.Field
	DsaAttestation             apijson.Field
	Status                     apijson.Field
	Type                       apijson.Field
	URLs                       apijson.Field
	AgentName                  apijson.Field
	Comments                   apijson.Field
	Court                      apijson.Field
	DestinationIPs             apijson.Field
	HostNotification           apijson.Field
	Justification              apijson.Field
	NcmecNotification          apijson.Field
	NcseiSubjectRepresentation apijson.Field
	OriginalWork               apijson.Field
	OwnerNotification          apijson.Field
	PortsProtocols             apijson.Field
	RegWhoRequest              apijson.Field
	ReportedCountry            apijson.Field
	ReportedUserAgent          apijson.Field
	SourceIPs                  apijson.Field
	Submitter                  apijson.Field
	Subtypes                   apijson.Field
	Title                      apijson.Field
	Udrp                       apijson.Field
	Urs                        apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *SubmittedGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r submittedGetResponseJSON) RawJSON() string {
	return r.raw
}

// Submitter-safe reason for a denied report. Null when unavailable.
type SubmittedGetResponseDenialReason string

const (
	SubmittedGetResponseDenialReasonUnableToConfirm                SubmittedGetResponseDenialReason = "unable_to_confirm"
	SubmittedGetResponseDenialReasonIncompleteReport               SubmittedGetResponseDenialReason = "incomplete_report"
	SubmittedGetResponseDenialReasonNotOnCloudflare                SubmittedGetResponseDenialReason = "not_on_cloudflare"
	SubmittedGetResponseDenialReasonDuplicateReport                SubmittedGetResponseDenialReason = "duplicate_report"
	SubmittedGetResponseDenialReasonContentRemoved                 SubmittedGetResponseDenialReason = "content_removed"
	SubmittedGetResponseDenialReasonReportDetailsMismatch          SubmittedGetResponseDenialReason = "report_details_mismatch"
	SubmittedGetResponseDenialReasonNoAbuseFound                   SubmittedGetResponseDenialReason = "no_abuse_found"
	SubmittedGetResponseDenialReasonMissingOriginalWork            SubmittedGetResponseDenialReason = "missing_original_work"
	SubmittedGetResponseDenialReasonDirectURLRequired              SubmittedGetResponseDenialReason = "direct_url_required"
	SubmittedGetResponseDenialReasonWrongReportCategory            SubmittedGetResponseDenialReason = "wrong_report_category"
	SubmittedGetResponseDenialReasonContentUnavailable             SubmittedGetResponseDenialReason = "content_unavailable"
	SubmittedGetResponseDenialReasonLawEnforcementReferralRequired SubmittedGetResponseDenialReason = "law_enforcement_referral_required"
	SubmittedGetResponseDenialReasonDomainDisputeProcessRequired   SubmittedGetResponseDenialReason = "domain_dispute_process_required"
)

func (r SubmittedGetResponseDenialReason) IsKnown() bool {
	switch r {
	case SubmittedGetResponseDenialReasonUnableToConfirm, SubmittedGetResponseDenialReasonIncompleteReport, SubmittedGetResponseDenialReasonNotOnCloudflare, SubmittedGetResponseDenialReasonDuplicateReport, SubmittedGetResponseDenialReasonContentRemoved, SubmittedGetResponseDenialReasonReportDetailsMismatch, SubmittedGetResponseDenialReasonNoAbuseFound, SubmittedGetResponseDenialReasonMissingOriginalWork, SubmittedGetResponseDenialReasonDirectURLRequired, SubmittedGetResponseDenialReasonWrongReportCategory, SubmittedGetResponseDenialReasonContentUnavailable, SubmittedGetResponseDenialReasonLawEnforcementReferralRequired, SubmittedGetResponseDenialReasonDomainDisputeProcessRequired:
		return true
	}
	return false
}

// Status visible to the account that submitted the report.
type SubmittedGetResponseStatus string

const (
	SubmittedGetResponseStatusSubmitted SubmittedGetResponseStatus = "submitted"
	SubmittedGetResponseStatusAccepted  SubmittedGetResponseStatus = "accepted"
	SubmittedGetResponseStatusDenied    SubmittedGetResponseStatus = "denied"
)

func (r SubmittedGetResponseStatus) IsKnown() bool {
	switch r {
	case SubmittedGetResponseStatusSubmitted, SubmittedGetResponseStatusAccepted, SubmittedGetResponseStatusDenied:
		return true
	}
	return false
}

// The abuse report type
type SubmittedGetResponseType string

const (
	SubmittedGetResponseTypePhish   SubmittedGetResponseType = "PHISH"
	SubmittedGetResponseTypeGen     SubmittedGetResponseType = "GEN"
	SubmittedGetResponseTypeThreat  SubmittedGetResponseType = "THREAT"
	SubmittedGetResponseTypeDmca    SubmittedGetResponseType = "DMCA"
	SubmittedGetResponseTypeEmer    SubmittedGetResponseType = "EMER"
	SubmittedGetResponseTypeTm      SubmittedGetResponseType = "TM"
	SubmittedGetResponseTypeRegWho  SubmittedGetResponseType = "REG_WHO"
	SubmittedGetResponseTypeNcsei   SubmittedGetResponseType = "NCSEI"
	SubmittedGetResponseTypeNetwork SubmittedGetResponseType = "NETWORK"
)

func (r SubmittedGetResponseType) IsKnown() bool {
	switch r {
	case SubmittedGetResponseTypePhish, SubmittedGetResponseTypeGen, SubmittedGetResponseTypeThreat, SubmittedGetResponseTypeDmca, SubmittedGetResponseTypeEmer, SubmittedGetResponseTypeTm, SubmittedGetResponseTypeRegWho, SubmittedGetResponseTypeNcsei, SubmittedGetResponseTypeNetwork:
		return true
	}
	return false
}

// RDP-mandated fields for registrar WHOIS data disclosure requests.
type SubmittedGetResponseRegWhoRequest struct {
	// Affirmation that the request is made in good faith per RDP 10.2.4. Must be true.
	RegWhoGoodFaithAffirmation bool `json:"reg_who_good_faith_affirmation" api:"required"`
	// Agreement to process data lawfully per RDP 10.2.5. Must be true.
	RegWhoLawfulProcessingAgreement bool `json:"reg_who_lawful_processing_agreement" api:"required"`
	// Legal rights and rationale for the request per RDP 10.2.3. Required for all
	// WHOIS requests.
	RegWhoLegalBasis string `json:"reg_who_legal_basis" api:"required"`
	// The type of WHOIS data request per RDP procedure.
	RegWhoRequestType SubmittedGetResponseRegWhoRequestRegWhoRequestType `json:"reg_who_request_type" api:"required"`
	// The specific WHOIS data elements being requested per RDP 10.2.2. Required for
	// all WHOIS requests.
	RegWhoRequestedDataElements []SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElement `json:"reg_who_requested_data_elements" api:"required"`
	// Optional authorization statement or power of attorney per RDP 10.2.1.3.
	RegWhoAuthorizationStatement string `json:"reg_who_authorization_statement"`
	// The nature of the requestor per RDP 10.2.1.2.
	RegWhoRequestorType SubmittedGetResponseRegWhoRequestRegWhoRequestorType `json:"reg_who_requestor_type"`
	JSON                submittedGetResponseRegWhoRequestJSON                `json:"-"`
}

// submittedGetResponseRegWhoRequestJSON contains the JSON metadata for the struct
// [SubmittedGetResponseRegWhoRequest]
type submittedGetResponseRegWhoRequestJSON struct {
	RegWhoGoodFaithAffirmation      apijson.Field
	RegWhoLawfulProcessingAgreement apijson.Field
	RegWhoLegalBasis                apijson.Field
	RegWhoRequestType               apijson.Field
	RegWhoRequestedDataElements     apijson.Field
	RegWhoAuthorizationStatement    apijson.Field
	RegWhoRequestorType             apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *SubmittedGetResponseRegWhoRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r submittedGetResponseRegWhoRequestJSON) RawJSON() string {
	return r.raw
}

// The type of WHOIS data request per RDP procedure.
type SubmittedGetResponseRegWhoRequestRegWhoRequestType string

const (
	SubmittedGetResponseRegWhoRequestRegWhoRequestTypeDisclosure   SubmittedGetResponseRegWhoRequestRegWhoRequestType = "disclosure"
	SubmittedGetResponseRegWhoRequestRegWhoRequestTypeInvalidWhois SubmittedGetResponseRegWhoRequestRegWhoRequestType = "invalid_whois"
)

func (r SubmittedGetResponseRegWhoRequestRegWhoRequestType) IsKnown() bool {
	switch r {
	case SubmittedGetResponseRegWhoRequestRegWhoRequestTypeDisclosure, SubmittedGetResponseRegWhoRequestRegWhoRequestTypeInvalidWhois:
		return true
	}
	return false
}

type SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElement string

const (
	SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementRegistrantName              SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElement = "registrant_name"
	SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementRegistrantOrganization      SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElement = "registrant_organization"
	SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementRegistrantEmail             SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElement = "registrant_email"
	SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementRegistrantPhone             SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElement = "registrant_phone"
	SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementRegistrantAddress           SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElement = "registrant_address"
	SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementRegistrantAddressCountry    SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElement = "registrant_address_country"
	SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementRegistrantAddressPostalCode SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElement = "registrant_address_postal_code"
	SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementAdminName                   SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElement = "admin_name"
	SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementAdminOrganization           SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElement = "admin_organization"
	SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementAdminEmail                  SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElement = "admin_email"
	SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementAdminPhone                  SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElement = "admin_phone"
	SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementAdminAddress                SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElement = "admin_address"
	SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementTechName                    SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElement = "tech_name"
	SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementTechOrganization            SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElement = "tech_organization"
	SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementTechEmail                   SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElement = "tech_email"
	SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementTechPhone                   SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElement = "tech_phone"
	SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementTechAddress                 SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElement = "tech_address"
)

func (r SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElement) IsKnown() bool {
	switch r {
	case SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementRegistrantName, SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementRegistrantOrganization, SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementRegistrantEmail, SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementRegistrantPhone, SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementRegistrantAddress, SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementRegistrantAddressCountry, SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementRegistrantAddressPostalCode, SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementAdminName, SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementAdminOrganization, SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementAdminEmail, SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementAdminPhone, SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementAdminAddress, SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementTechName, SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementTechOrganization, SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementTechEmail, SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementTechPhone, SubmittedGetResponseRegWhoRequestRegWhoRequestedDataElementTechAddress:
		return true
	}
	return false
}

// The nature of the requestor per RDP 10.2.1.2.
type SubmittedGetResponseRegWhoRequestRegWhoRequestorType string

const (
	SubmittedGetResponseRegWhoRequestRegWhoRequestorTypeGovernment  SubmittedGetResponseRegWhoRequestRegWhoRequestorType = "government"
	SubmittedGetResponseRegWhoRequestRegWhoRequestorTypeCorporation SubmittedGetResponseRegWhoRequestRegWhoRequestorType = "corporation"
	SubmittedGetResponseRegWhoRequestRegWhoRequestorTypeIndividual  SubmittedGetResponseRegWhoRequestRegWhoRequestorType = "individual"
)

func (r SubmittedGetResponseRegWhoRequestRegWhoRequestorType) IsKnown() bool {
	switch r {
	case SubmittedGetResponseRegWhoRequestRegWhoRequestorTypeGovernment, SubmittedGetResponseRegWhoRequestRegWhoRequestorTypeCorporation, SubmittedGetResponseRegWhoRequestRegWhoRequestorTypeIndividual:
		return true
	}
	return false
}

// Information about the submitter of the report.
type SubmittedGetResponseSubmitter struct {
	Company   string                            `json:"company"`
	Email     string                            `json:"email"`
	Name      string                            `json:"name"`
	Telephone string                            `json:"telephone"`
	JSON      submittedGetResponseSubmitterJSON `json:"-"`
}

// submittedGetResponseSubmitterJSON contains the JSON metadata for the struct
// [SubmittedGetResponseSubmitter]
type submittedGetResponseSubmitterJSON struct {
	Company     apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	Telephone   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubmittedGetResponseSubmitter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r submittedGetResponseSubmitterJSON) RawJSON() string {
	return r.raw
}

type SubmittedListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filter by report code.
	ID param.Field[string] `query:"id"`
	// Return reports submitted after this time.
	CreatedAfter param.Field[time.Time] `query:"created_after" format:"date-time"`
	// Return reports submitted before this time.
	CreatedBefore param.Field[time.Time] `query:"created_before" format:"date-time"`
	// Filter by reported domain. This parameter can be specified multiple times.
	Domain param.Field[[]string] `query:"domain"`
	// Page of submitted reports to return.
	Page param.Field[int64] `query:"page"`
	// Number of submitted reports per page.
	PerPage param.Field[int64] `query:"per_page"`
	// A property and direction to sort by (id, cdate, domain, type, status).
	Sort param.Field[string] `query:"sort"`
	// Filter by submitter-facing status. This parameter can be specified multiple
	// times.
	Status param.Field[[]SubmittedListParamsStatus] `query:"status"`
	// Filter by report type. This parameter can be specified multiple times.
	Type param.Field[[]SubmittedListParamsType] `query:"type"`
}

// URLQuery serializes [SubmittedListParams]'s query parameters as `url.Values`.
func (r SubmittedListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Status visible to the account that submitted the report.
type SubmittedListParamsStatus string

const (
	SubmittedListParamsStatusSubmitted SubmittedListParamsStatus = "submitted"
	SubmittedListParamsStatusAccepted  SubmittedListParamsStatus = "accepted"
	SubmittedListParamsStatusDenied    SubmittedListParamsStatus = "denied"
)

func (r SubmittedListParamsStatus) IsKnown() bool {
	switch r {
	case SubmittedListParamsStatusSubmitted, SubmittedListParamsStatusAccepted, SubmittedListParamsStatusDenied:
		return true
	}
	return false
}

// The abuse report type
type SubmittedListParamsType string

const (
	SubmittedListParamsTypePhish   SubmittedListParamsType = "PHISH"
	SubmittedListParamsTypeGen     SubmittedListParamsType = "GEN"
	SubmittedListParamsTypeThreat  SubmittedListParamsType = "THREAT"
	SubmittedListParamsTypeDmca    SubmittedListParamsType = "DMCA"
	SubmittedListParamsTypeEmer    SubmittedListParamsType = "EMER"
	SubmittedListParamsTypeTm      SubmittedListParamsType = "TM"
	SubmittedListParamsTypeRegWho  SubmittedListParamsType = "REG_WHO"
	SubmittedListParamsTypeNcsei   SubmittedListParamsType = "NCSEI"
	SubmittedListParamsTypeNetwork SubmittedListParamsType = "NETWORK"
)

func (r SubmittedListParamsType) IsKnown() bool {
	switch r {
	case SubmittedListParamsTypePhish, SubmittedListParamsTypeGen, SubmittedListParamsTypeThreat, SubmittedListParamsTypeDmca, SubmittedListParamsTypeEmer, SubmittedListParamsTypeTm, SubmittedListParamsTypeRegWho, SubmittedListParamsTypeNcsei, SubmittedListParamsTypeNetwork:
		return true
	}
	return false
}

type SubmittedGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SubmittedGetResponseEnvelope struct {
	Result   SubmittedGetResponse                   `json:"result" api:"required"`
	Success  bool                                   `json:"success" api:"required"`
	Errors   []SubmittedGetResponseEnvelopeErrors   `json:"errors"`
	Messages []SubmittedGetResponseEnvelopeMessages `json:"messages"`
	JSON     submittedGetResponseEnvelopeJSON       `json:"-"`
}

// submittedGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SubmittedGetResponseEnvelope]
type submittedGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubmittedGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r submittedGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SubmittedGetResponseEnvelopeErrors struct {
	Message string                                 `json:"message" api:"required"`
	Code    SubmittedGetResponseEnvelopeErrorsCode `json:"code"`
	JSON    submittedGetResponseEnvelopeErrorsJSON `json:"-"`
}

// submittedGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SubmittedGetResponseEnvelopeErrors]
type submittedGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubmittedGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r submittedGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type SubmittedGetResponseEnvelopeErrorsCode interface {
	ImplementsSubmittedGetResponseEnvelopeErrorsCode()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubmittedGetResponseEnvelopeErrorsCode)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type SubmittedGetResponseEnvelopeMessages struct {
	Message string                                   `json:"message" api:"required"`
	JSON    submittedGetResponseEnvelopeMessagesJSON `json:"-"`
}

// submittedGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SubmittedGetResponseEnvelopeMessages]
type submittedGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubmittedGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r submittedGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
