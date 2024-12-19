// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package abuse_reports

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// AbuseReportService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAbuseReportService] method instead.
type AbuseReportService struct {
	Options []option.RequestOption
}

// NewAbuseReportService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAbuseReportService(opts ...option.RequestOption) (r *AbuseReportService) {
	r = &AbuseReportService{}
	r.Options = opts
	return
}

// Submit the Abuse Report of a particular type
func (r *AbuseReportService) New(ctx context.Context, reportType AbuseReportNewParamsReportType, params AbuseReportNewParams, opts ...option.RequestOption) (res *string, err error) {
	var env AbuseReportNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/v1/abuse-reports/%v", params.AccountID, reportType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AbuseReportNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// The abuse report type
	Act param.Field[AbuseReportNewParamsAct] `json:"act,required"`
	// A valid email of the abuse reporter
	Email param.Field[string] `json:"email,required"`
	// Should match the value provided in `email`
	Email2 param.Field[string] `json:"email2,required"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	HostNotification param.Field[AbuseReportNewParamsHostNotification] `json:"host_notification,required"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	NcmecNotification param.Field[AbuseReportNewParamsNcmecNotification] `json:"ncmec_notification,required"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	OwnerNotification param.Field[AbuseReportNewParamsOwnerNotification] `json:"owner_notification,required"`
	// A list of valid URLs separated by ‘ ’ (new line character). The list of the URLs
	// should not exceed 250 URLs. All URLs should have the same hostname. Each URL
	// should be unique
	URLs param.Field[string] `json:"urls,required"`
	// Text not exceeding 100 characters
	Address1 param.Field[string] `json:"address1"`
	// The name of the copyright holder. Text not exceeding 60 characters.
	AgentName param.Field[string] `json:"agent_name"`
	// Can be 0 or 1
	Agree param.Field[AbuseReportNewParamsAgree] `json:"agree"`
	// Text not exceeding 255 characters
	City param.Field[string] `json:"city"`
	// Any additional comments about the infringement not exceeding 2000 characters
	Comments param.Field[string] `json:"comments"`
	// Text not exceeding 100 characters
	Company param.Field[string] `json:"company"`
	// Text not exceeding 255 characters
	Country param.Field[string] `json:"country"`
	// A list of IP addresses separated by ‘ ’ (new line character). The list of
	// destination IPs should not exceed 30 IP addresses. Each one of the IP addresses
	// ought to be unique
	DestinationIPs param.Field[string] `json:"destination_ips"`
	// A detailed description of the infringement, including any necessary access
	// details and the exact steps needed to view the content, not exceeding 5000
	// characters
	Justification param.Field[string] `json:"justification"`
	// Text not exceeding 255 characters
	Name param.Field[string] `json:"name"`
	// If the submitter is the target of NCSEI in the URLs of the abuse report
	NcseiSubjectRepresentation param.Field[bool] `json:"ncsei_subject_representation"`
	// Text not exceeding 255 characters
	OriginalWork param.Field[string] `json:"original_work"`
	// A comma separated list of ports and protocols e.g. 80/TCP, 22/UDP. The total
	// size of the field should not exceed 2000 characters. Each individual
	// port/protocol should not exceed 100 characters. The list should not have more
	// than 30 unique ports and protocols.
	PortsProtocols param.Field[string] `json:"ports_protocols"`
	// Required for DMCA reports, should be same as Name. An affirmation that all
	// information in the report is true and accurate while agreeing to the policies of
	// Cloudflare's abuse reports
	Signature param.Field[string] `json:"signature"`
	// A list of IP addresses separated by ‘ ’ (new line character). The list of source
	// IPs should not exceed 30 IP addresses. Each one of the IP addresses ought to be
	// unique
	SourceIPs param.Field[string] `json:"source_ips"`
	// Text not exceeding 255 characters
	State param.Field[string] `json:"state"`
	// Text not exceeding 20 characters
	Tele param.Field[string] `json:"tele"`
	// Text not exceeding 255 characters
	Title param.Field[string] `json:"title"`
	// Text not exceeding 1000 characters
	TrademarkNumber param.Field[string] `json:"trademark_number"`
	// Text not exceeding 1000 characters
	TrademarkOffice param.Field[string] `json:"trademark_office"`
	// Text not exceeding 1000 characters
	TrademarkSymbol param.Field[string] `json:"trademark_symbol"`
}

func (r AbuseReportNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The abuse report type
type AbuseReportNewParamsReportType string

const (
	AbuseReportNewParamsReportTypeAbuseDmca           AbuseReportNewParamsReportType = "abuse_dmca"
	AbuseReportNewParamsReportTypeAbuseTrademark      AbuseReportNewParamsReportType = "abuse_trademark"
	AbuseReportNewParamsReportTypeAbuseGeneral        AbuseReportNewParamsReportType = "abuse_general"
	AbuseReportNewParamsReportTypeAbusePhishing       AbuseReportNewParamsReportType = "abuse_phishing"
	AbuseReportNewParamsReportTypeAbuseChildren       AbuseReportNewParamsReportType = "abuse_children"
	AbuseReportNewParamsReportTypeAbuseThreat         AbuseReportNewParamsReportType = "abuse_threat"
	AbuseReportNewParamsReportTypeAbuseRegistrarWhois AbuseReportNewParamsReportType = "abuse_registrar_whois"
	AbuseReportNewParamsReportTypeAbuseNcsei          AbuseReportNewParamsReportType = "abuse_ncsei"
)

func (r AbuseReportNewParamsReportType) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsReportTypeAbuseDmca, AbuseReportNewParamsReportTypeAbuseTrademark, AbuseReportNewParamsReportTypeAbuseGeneral, AbuseReportNewParamsReportTypeAbusePhishing, AbuseReportNewParamsReportTypeAbuseChildren, AbuseReportNewParamsReportTypeAbuseThreat, AbuseReportNewParamsReportTypeAbuseRegistrarWhois, AbuseReportNewParamsReportTypeAbuseNcsei:
		return true
	}
	return false
}

// The abuse report type
type AbuseReportNewParamsAct string

const (
	AbuseReportNewParamsActAbuseDmca           AbuseReportNewParamsAct = "abuse_dmca"
	AbuseReportNewParamsActAbuseTrademark      AbuseReportNewParamsAct = "abuse_trademark"
	AbuseReportNewParamsActAbuseGeneral        AbuseReportNewParamsAct = "abuse_general"
	AbuseReportNewParamsActAbusePhishing       AbuseReportNewParamsAct = "abuse_phishing"
	AbuseReportNewParamsActAbuseChildren       AbuseReportNewParamsAct = "abuse_children"
	AbuseReportNewParamsActAbuseThreat         AbuseReportNewParamsAct = "abuse_threat"
	AbuseReportNewParamsActAbuseRegistrarWhois AbuseReportNewParamsAct = "abuse_registrar_whois"
	AbuseReportNewParamsActAbuseNcsei          AbuseReportNewParamsAct = "abuse_ncsei"
)

func (r AbuseReportNewParamsAct) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsActAbuseDmca, AbuseReportNewParamsActAbuseTrademark, AbuseReportNewParamsActAbuseGeneral, AbuseReportNewParamsActAbusePhishing, AbuseReportNewParamsActAbuseChildren, AbuseReportNewParamsActAbuseThreat, AbuseReportNewParamsActAbuseRegistrarWhois, AbuseReportNewParamsActAbuseNcsei:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsHostNotification string

const (
	AbuseReportNewParamsHostNotificationSend     AbuseReportNewParamsHostNotification = "send"
	AbuseReportNewParamsHostNotificationSendAnon AbuseReportNewParamsHostNotification = "send-anon"
	AbuseReportNewParamsHostNotificationNone     AbuseReportNewParamsHostNotification = "none"
)

func (r AbuseReportNewParamsHostNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsHostNotificationSend, AbuseReportNewParamsHostNotificationSendAnon, AbuseReportNewParamsHostNotificationNone:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsNcmecNotification string

const (
	AbuseReportNewParamsNcmecNotificationSend     AbuseReportNewParamsNcmecNotification = "send"
	AbuseReportNewParamsNcmecNotificationSendAnon AbuseReportNewParamsNcmecNotification = "send-anon"
	AbuseReportNewParamsNcmecNotificationNone     AbuseReportNewParamsNcmecNotification = "none"
)

func (r AbuseReportNewParamsNcmecNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsNcmecNotificationSend, AbuseReportNewParamsNcmecNotificationSendAnon, AbuseReportNewParamsNcmecNotificationNone:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsOwnerNotification string

const (
	AbuseReportNewParamsOwnerNotificationSend     AbuseReportNewParamsOwnerNotification = "send"
	AbuseReportNewParamsOwnerNotificationSendAnon AbuseReportNewParamsOwnerNotification = "send-anon"
	AbuseReportNewParamsOwnerNotificationNone     AbuseReportNewParamsOwnerNotification = "none"
)

func (r AbuseReportNewParamsOwnerNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsOwnerNotificationSend, AbuseReportNewParamsOwnerNotificationSendAnon, AbuseReportNewParamsOwnerNotificationNone:
		return true
	}
	return false
}

// Can be 0 or 1
type AbuseReportNewParamsAgree int64

const (
	AbuseReportNewParamsAgree0 AbuseReportNewParamsAgree = 0
	AbuseReportNewParamsAgree1 AbuseReportNewParamsAgree = 1
)

func (r AbuseReportNewParamsAgree) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsAgree0, AbuseReportNewParamsAgree1:
		return true
	}
	return false
}

type AbuseReportNewResponseEnvelope struct {
	// The identifier for the submitted abuse report.
	AbuseRand string                                `json:"abuse_rand,required"`
	Request   AbuseReportNewResponseEnvelopeRequest `json:"request,required"`
	// The result should be 'success' for successful response
	Result string                             `json:"result,required"`
	JSON   abuseReportNewResponseEnvelopeJSON `json:"-"`
}

// abuseReportNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AbuseReportNewResponseEnvelope]
type abuseReportNewResponseEnvelopeJSON struct {
	AbuseRand   apijson.Field
	Request     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AbuseReportNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r abuseReportNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AbuseReportNewResponseEnvelopeRequest struct {
	// The abuse report type
	Act  AbuseReportNewResponseEnvelopeRequestAct  `json:"act,required"`
	JSON abuseReportNewResponseEnvelopeRequestJSON `json:"-"`
}

// abuseReportNewResponseEnvelopeRequestJSON contains the JSON metadata for the
// struct [AbuseReportNewResponseEnvelopeRequest]
type abuseReportNewResponseEnvelopeRequestJSON struct {
	Act         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AbuseReportNewResponseEnvelopeRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r abuseReportNewResponseEnvelopeRequestJSON) RawJSON() string {
	return r.raw
}

// The abuse report type
type AbuseReportNewResponseEnvelopeRequestAct string

const (
	AbuseReportNewResponseEnvelopeRequestActAbuseDmca           AbuseReportNewResponseEnvelopeRequestAct = "abuse_dmca"
	AbuseReportNewResponseEnvelopeRequestActAbuseTrademark      AbuseReportNewResponseEnvelopeRequestAct = "abuse_trademark"
	AbuseReportNewResponseEnvelopeRequestActAbuseGeneral        AbuseReportNewResponseEnvelopeRequestAct = "abuse_general"
	AbuseReportNewResponseEnvelopeRequestActAbusePhishing       AbuseReportNewResponseEnvelopeRequestAct = "abuse_phishing"
	AbuseReportNewResponseEnvelopeRequestActAbuseChildren       AbuseReportNewResponseEnvelopeRequestAct = "abuse_children"
	AbuseReportNewResponseEnvelopeRequestActAbuseThreat         AbuseReportNewResponseEnvelopeRequestAct = "abuse_threat"
	AbuseReportNewResponseEnvelopeRequestActAbuseRegistrarWhois AbuseReportNewResponseEnvelopeRequestAct = "abuse_registrar_whois"
	AbuseReportNewResponseEnvelopeRequestActAbuseNcsei          AbuseReportNewResponseEnvelopeRequestAct = "abuse_ncsei"
)

func (r AbuseReportNewResponseEnvelopeRequestAct) IsKnown() bool {
	switch r {
	case AbuseReportNewResponseEnvelopeRequestActAbuseDmca, AbuseReportNewResponseEnvelopeRequestActAbuseTrademark, AbuseReportNewResponseEnvelopeRequestActAbuseGeneral, AbuseReportNewResponseEnvelopeRequestActAbusePhishing, AbuseReportNewResponseEnvelopeRequestActAbuseChildren, AbuseReportNewResponseEnvelopeRequestActAbuseThreat, AbuseReportNewResponseEnvelopeRequestActAbuseRegistrarWhois, AbuseReportNewResponseEnvelopeRequestActAbuseNcsei:
		return true
	}
	return false
}
