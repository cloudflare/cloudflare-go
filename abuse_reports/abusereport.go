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
	path := fmt.Sprintf("accounts/%s/abuse-reports/%v", params.AccountID, reportType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AbuseReportNewParams struct {
	AccountID param.Field[string]           `path:"account_id,required"`
	Body      AbuseReportNewParamsBodyUnion `json:"body"`
}

func (r AbuseReportNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
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

type AbuseReportNewParamsBody struct {
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	OwnerNotification param.Field[AbuseReportNewParamsBodyOwnerNotification] `json:"owner_notification,required"`
	// The abuse report type
	Act param.Field[AbuseReportNewParamsBodyAct] `json:"act"`
	// Text not exceeding 100 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Address1 param.Field[string] `json:"address1"`
	// The name of the copyright holder. Text not exceeding 60 characters. This field
	// may be released by Cloudflare to third parties such as the Lumen Database
	// (https://lumendatabase.org/).
	AgentName param.Field[string] `json:"agent_name"`
	// Can be `0` for false or `1` for true. Must be value: 1 for DMCA reports
	Agree param.Field[AbuseReportNewParamsBodyAgree] `json:"agree"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	City param.Field[string] `json:"city"`
	// Any additional comments about the infringement not exceeding 2000 characters
	Comments param.Field[string] `json:"comments"`
	// Text not exceeding 100 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Company param.Field[string] `json:"company"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Country param.Field[string] `json:"country"`
	// A list of IP addresses separated by ‘ ’ (new line character). The list of
	// destination IPs should not exceed 30 IP addresses. Each one of the IP addresses
	// ought to be unique
	DestinationIPs param.Field[string] `json:"destination_ips"`
	// A valid email of the abuse reporter. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Email param.Field[string] `json:"email"`
	// Should match the value provided in `email`
	Email2 param.Field[string] `json:"email2"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	HostNotification param.Field[AbuseReportNewParamsBodyHostNotification] `json:"host_notification"`
	// A detailed description of the infringement, including any necessary access
	// details and the exact steps needed to view the content, not exceeding 5000
	// characters
	Justification param.Field[string] `json:"justification"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Name param.Field[string] `json:"name"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	NcmecNotification param.Field[AbuseReportNewParamsBodyNcmecNotification] `json:"ncmec_notification"`
	// If the submitter is the target of NCSEI in the URLs of the abuse report.
	NcseiSubjectRepresentation param.Field[bool] `json:"ncsei_subject_representation"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
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
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	State param.Field[string] `json:"state"`
	// Text not exceeding 20 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Tele param.Field[string] `json:"tele"`
	// Text not exceeding 255 characters
	Title param.Field[string] `json:"title"`
	// Text not exceeding 1000 characters
	TrademarkNumber param.Field[string] `json:"trademark_number"`
	// Text not exceeding 1000 characters
	TrademarkOffice param.Field[string] `json:"trademark_office"`
	// Text not exceeding 1000 characters
	TrademarkSymbol param.Field[string] `json:"trademark_symbol"`
	// A list of valid URLs separated by ‘ ’ (new line character). The list of the URLs
	// should not exceed 250 URLs. All URLs should have the same hostname. Each URL
	// should be unique. This field may be released by Cloudflare to third parties such
	// as the Lumen Database (https://lumendatabase.org/).
	URLs param.Field[string] `json:"urls"`
}

func (r AbuseReportNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AbuseReportNewParamsBody) implementsAbuseReportNewParamsBodyUnion() {}

// Satisfied by [abuse_reports.AbuseReportNewParamsBodyAbuseReportsDmcaReport],
// [abuse_reports.AbuseReportNewParamsBodyAbuseReportsTrademarkReport],
// [abuse_reports.AbuseReportNewParamsBodyAbuseReportsGeneralReport],
// [abuse_reports.AbuseReportNewParamsBodyAbuseReportsPhishingReport],
// [abuse_reports.AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReport],
// [abuse_reports.AbuseReportNewParamsBodyAbuseReportsThreatReport],
// [abuse_reports.AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReport],
// [abuse_reports.AbuseReportNewParamsBodyAbuseReportsNcseiReport],
// [AbuseReportNewParamsBody].
type AbuseReportNewParamsBodyUnion interface {
	implementsAbuseReportNewParamsBodyUnion()
}

type AbuseReportNewParamsBodyAbuseReportsDmcaReport struct {
	// Text not exceeding 100 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Address1 param.Field[string] `json:"address1,required"`
	// The name of the copyright holder. Text not exceeding 60 characters. This field
	// may be released by Cloudflare to third parties such as the Lumen Database
	// (https://lumendatabase.org/).
	AgentName param.Field[string] `json:"agent_name,required"`
	// Can be `0` for false or `1` for true. Must be value: 1 for DMCA reports
	Agree param.Field[AbuseReportNewParamsBodyAbuseReportsDmcaReportAgree] `json:"agree,required"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	City param.Field[string] `json:"city,required"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Country param.Field[string] `json:"country,required"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	HostNotification param.Field[AbuseReportNewParamsBodyAbuseReportsDmcaReportHostNotification] `json:"host_notification,required"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	OriginalWork param.Field[string] `json:"original_work,required"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	OwnerNotification param.Field[AbuseReportNewParamsBodyAbuseReportsDmcaReportOwnerNotification] `json:"owner_notification,required"`
	// Required for DMCA reports, should be same as Name. An affirmation that all
	// information in the report is true and accurate while agreeing to the policies of
	// Cloudflare's abuse reports
	Signature param.Field[string] `json:"signature,required"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	State param.Field[string] `json:"state,required"`
	// The abuse report type
	Act param.Field[AbuseReportNewParamsBodyAbuseReportsDmcaReportAct] `json:"act"`
	// Any additional comments about the infringement not exceeding 2000 characters
	Comments param.Field[string] `json:"comments"`
	// Text not exceeding 100 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Company param.Field[string] `json:"company"`
	// A list of IP addresses separated by ‘ ’ (new line character). The list of
	// destination IPs should not exceed 30 IP addresses. Each one of the IP addresses
	// ought to be unique
	DestinationIPs param.Field[string] `json:"destination_ips"`
	// A valid email of the abuse reporter. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Email param.Field[string] `json:"email"`
	// Should match the value provided in `email`
	Email2 param.Field[string] `json:"email2"`
	// A detailed description of the infringement, including any necessary access
	// details and the exact steps needed to view the content, not exceeding 5000
	// characters
	Justification param.Field[string] `json:"justification"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Name param.Field[string] `json:"name"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	NcmecNotification param.Field[AbuseReportNewParamsBodyAbuseReportsDmcaReportNcmecNotification] `json:"ncmec_notification"`
	// If the submitter is the target of NCSEI in the URLs of the abuse report.
	NcseiSubjectRepresentation param.Field[bool] `json:"ncsei_subject_representation"`
	// A comma separated list of ports and protocols e.g. 80/TCP, 22/UDP. The total
	// size of the field should not exceed 2000 characters. Each individual
	// port/protocol should not exceed 100 characters. The list should not have more
	// than 30 unique ports and protocols.
	PortsProtocols param.Field[string] `json:"ports_protocols"`
	// A list of IP addresses separated by ‘ ’ (new line character). The list of source
	// IPs should not exceed 30 IP addresses. Each one of the IP addresses ought to be
	// unique
	SourceIPs param.Field[string] `json:"source_ips"`
	// Text not exceeding 20 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Tele param.Field[string] `json:"tele"`
	// Text not exceeding 255 characters
	Title param.Field[string] `json:"title"`
	// Text not exceeding 1000 characters
	TrademarkNumber param.Field[string] `json:"trademark_number"`
	// Text not exceeding 1000 characters
	TrademarkOffice param.Field[string] `json:"trademark_office"`
	// Text not exceeding 1000 characters
	TrademarkSymbol param.Field[string] `json:"trademark_symbol"`
	// A list of valid URLs separated by ‘ ’ (new line character). The list of the URLs
	// should not exceed 250 URLs. All URLs should have the same hostname. Each URL
	// should be unique. This field may be released by Cloudflare to third parties such
	// as the Lumen Database (https://lumendatabase.org/).
	URLs param.Field[string] `json:"urls"`
}

func (r AbuseReportNewParamsBodyAbuseReportsDmcaReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AbuseReportNewParamsBodyAbuseReportsDmcaReport) implementsAbuseReportNewParamsBodyUnion() {}

// Can be `0` for false or `1` for true. Must be value: 1 for DMCA reports
type AbuseReportNewParamsBodyAbuseReportsDmcaReportAgree int64

const (
	AbuseReportNewParamsBodyAbuseReportsDmcaReportAgree0 AbuseReportNewParamsBodyAbuseReportsDmcaReportAgree = 0
	AbuseReportNewParamsBodyAbuseReportsDmcaReportAgree1 AbuseReportNewParamsBodyAbuseReportsDmcaReportAgree = 1
)

func (r AbuseReportNewParamsBodyAbuseReportsDmcaReportAgree) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsDmcaReportAgree0, AbuseReportNewParamsBodyAbuseReportsDmcaReportAgree1:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsDmcaReportHostNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsDmcaReportHostNotificationSend     AbuseReportNewParamsBodyAbuseReportsDmcaReportHostNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsDmcaReportHostNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsDmcaReportHostNotification = "send-anon"
	AbuseReportNewParamsBodyAbuseReportsDmcaReportHostNotificationNone     AbuseReportNewParamsBodyAbuseReportsDmcaReportHostNotification = "none"
)

func (r AbuseReportNewParamsBodyAbuseReportsDmcaReportHostNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsDmcaReportHostNotificationSend, AbuseReportNewParamsBodyAbuseReportsDmcaReportHostNotificationSendAnon, AbuseReportNewParamsBodyAbuseReportsDmcaReportHostNotificationNone:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsDmcaReportOwnerNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsDmcaReportOwnerNotificationSend     AbuseReportNewParamsBodyAbuseReportsDmcaReportOwnerNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsDmcaReportOwnerNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsDmcaReportOwnerNotification = "send-anon"
	AbuseReportNewParamsBodyAbuseReportsDmcaReportOwnerNotificationNone     AbuseReportNewParamsBodyAbuseReportsDmcaReportOwnerNotification = "none"
)

func (r AbuseReportNewParamsBodyAbuseReportsDmcaReportOwnerNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsDmcaReportOwnerNotificationSend, AbuseReportNewParamsBodyAbuseReportsDmcaReportOwnerNotificationSendAnon, AbuseReportNewParamsBodyAbuseReportsDmcaReportOwnerNotificationNone:
		return true
	}
	return false
}

// The abuse report type
type AbuseReportNewParamsBodyAbuseReportsDmcaReportAct string

const (
	AbuseReportNewParamsBodyAbuseReportsDmcaReportActAbuseDmca           AbuseReportNewParamsBodyAbuseReportsDmcaReportAct = "abuse_dmca"
	AbuseReportNewParamsBodyAbuseReportsDmcaReportActAbuseTrademark      AbuseReportNewParamsBodyAbuseReportsDmcaReportAct = "abuse_trademark"
	AbuseReportNewParamsBodyAbuseReportsDmcaReportActAbuseGeneral        AbuseReportNewParamsBodyAbuseReportsDmcaReportAct = "abuse_general"
	AbuseReportNewParamsBodyAbuseReportsDmcaReportActAbusePhishing       AbuseReportNewParamsBodyAbuseReportsDmcaReportAct = "abuse_phishing"
	AbuseReportNewParamsBodyAbuseReportsDmcaReportActAbuseChildren       AbuseReportNewParamsBodyAbuseReportsDmcaReportAct = "abuse_children"
	AbuseReportNewParamsBodyAbuseReportsDmcaReportActAbuseThreat         AbuseReportNewParamsBodyAbuseReportsDmcaReportAct = "abuse_threat"
	AbuseReportNewParamsBodyAbuseReportsDmcaReportActAbuseRegistrarWhois AbuseReportNewParamsBodyAbuseReportsDmcaReportAct = "abuse_registrar_whois"
	AbuseReportNewParamsBodyAbuseReportsDmcaReportActAbuseNcsei          AbuseReportNewParamsBodyAbuseReportsDmcaReportAct = "abuse_ncsei"
)

func (r AbuseReportNewParamsBodyAbuseReportsDmcaReportAct) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsDmcaReportActAbuseDmca, AbuseReportNewParamsBodyAbuseReportsDmcaReportActAbuseTrademark, AbuseReportNewParamsBodyAbuseReportsDmcaReportActAbuseGeneral, AbuseReportNewParamsBodyAbuseReportsDmcaReportActAbusePhishing, AbuseReportNewParamsBodyAbuseReportsDmcaReportActAbuseChildren, AbuseReportNewParamsBodyAbuseReportsDmcaReportActAbuseThreat, AbuseReportNewParamsBodyAbuseReportsDmcaReportActAbuseRegistrarWhois, AbuseReportNewParamsBodyAbuseReportsDmcaReportActAbuseNcsei:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsDmcaReportNcmecNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsDmcaReportNcmecNotificationSend     AbuseReportNewParamsBodyAbuseReportsDmcaReportNcmecNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsDmcaReportNcmecNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsDmcaReportNcmecNotification = "send-anon"
	AbuseReportNewParamsBodyAbuseReportsDmcaReportNcmecNotificationNone     AbuseReportNewParamsBodyAbuseReportsDmcaReportNcmecNotification = "none"
)

func (r AbuseReportNewParamsBodyAbuseReportsDmcaReportNcmecNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsDmcaReportNcmecNotificationSend, AbuseReportNewParamsBodyAbuseReportsDmcaReportNcmecNotificationSendAnon, AbuseReportNewParamsBodyAbuseReportsDmcaReportNcmecNotificationNone:
		return true
	}
	return false
}

type AbuseReportNewParamsBodyAbuseReportsTrademarkReport struct {
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	HostNotification param.Field[AbuseReportNewParamsBodyAbuseReportsTrademarkReportHostNotification] `json:"host_notification,required"`
	// A detailed description of the infringement, including any necessary access
	// details and the exact steps needed to view the content, not exceeding 5000
	// characters
	Justification param.Field[string] `json:"justification,required"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	OwnerNotification param.Field[AbuseReportNewParamsBodyAbuseReportsTrademarkReportOwnerNotification] `json:"owner_notification,required"`
	// Text not exceeding 1000 characters
	TrademarkNumber param.Field[string] `json:"trademark_number,required"`
	// Text not exceeding 1000 characters
	TrademarkOffice param.Field[string] `json:"trademark_office,required"`
	// Text not exceeding 1000 characters
	TrademarkSymbol param.Field[string] `json:"trademark_symbol,required"`
	// The abuse report type
	Act param.Field[AbuseReportNewParamsBodyAbuseReportsTrademarkReportAct] `json:"act"`
	// Text not exceeding 100 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Address1 param.Field[string] `json:"address1"`
	// The name of the copyright holder. Text not exceeding 60 characters. This field
	// may be released by Cloudflare to third parties such as the Lumen Database
	// (https://lumendatabase.org/).
	AgentName param.Field[string] `json:"agent_name"`
	// Can be `0` for false or `1` for true. Must be value: 1 for DMCA reports
	Agree param.Field[AbuseReportNewParamsBodyAbuseReportsTrademarkReportAgree] `json:"agree"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	City param.Field[string] `json:"city"`
	// Any additional comments about the infringement not exceeding 2000 characters
	Comments param.Field[string] `json:"comments"`
	// Text not exceeding 100 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Company param.Field[string] `json:"company"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Country param.Field[string] `json:"country"`
	// A list of IP addresses separated by ‘ ’ (new line character). The list of
	// destination IPs should not exceed 30 IP addresses. Each one of the IP addresses
	// ought to be unique
	DestinationIPs param.Field[string] `json:"destination_ips"`
	// A valid email of the abuse reporter. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Email param.Field[string] `json:"email"`
	// Should match the value provided in `email`
	Email2 param.Field[string] `json:"email2"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Name param.Field[string] `json:"name"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	NcmecNotification param.Field[AbuseReportNewParamsBodyAbuseReportsTrademarkReportNcmecNotification] `json:"ncmec_notification"`
	// If the submitter is the target of NCSEI in the URLs of the abuse report.
	NcseiSubjectRepresentation param.Field[bool] `json:"ncsei_subject_representation"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
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
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	State param.Field[string] `json:"state"`
	// Text not exceeding 20 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Tele param.Field[string] `json:"tele"`
	// Text not exceeding 255 characters
	Title param.Field[string] `json:"title"`
	// A list of valid URLs separated by ‘ ’ (new line character). The list of the URLs
	// should not exceed 250 URLs. All URLs should have the same hostname. Each URL
	// should be unique. This field may be released by Cloudflare to third parties such
	// as the Lumen Database (https://lumendatabase.org/).
	URLs param.Field[string] `json:"urls"`
}

func (r AbuseReportNewParamsBodyAbuseReportsTrademarkReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AbuseReportNewParamsBodyAbuseReportsTrademarkReport) implementsAbuseReportNewParamsBodyUnion() {
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsTrademarkReportHostNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsTrademarkReportHostNotificationSend     AbuseReportNewParamsBodyAbuseReportsTrademarkReportHostNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsTrademarkReportHostNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsTrademarkReportHostNotification = "send-anon"
	AbuseReportNewParamsBodyAbuseReportsTrademarkReportHostNotificationNone     AbuseReportNewParamsBodyAbuseReportsTrademarkReportHostNotification = "none"
)

func (r AbuseReportNewParamsBodyAbuseReportsTrademarkReportHostNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsTrademarkReportHostNotificationSend, AbuseReportNewParamsBodyAbuseReportsTrademarkReportHostNotificationSendAnon, AbuseReportNewParamsBodyAbuseReportsTrademarkReportHostNotificationNone:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsTrademarkReportOwnerNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsTrademarkReportOwnerNotificationSend     AbuseReportNewParamsBodyAbuseReportsTrademarkReportOwnerNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsTrademarkReportOwnerNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsTrademarkReportOwnerNotification = "send-anon"
	AbuseReportNewParamsBodyAbuseReportsTrademarkReportOwnerNotificationNone     AbuseReportNewParamsBodyAbuseReportsTrademarkReportOwnerNotification = "none"
)

func (r AbuseReportNewParamsBodyAbuseReportsTrademarkReportOwnerNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsTrademarkReportOwnerNotificationSend, AbuseReportNewParamsBodyAbuseReportsTrademarkReportOwnerNotificationSendAnon, AbuseReportNewParamsBodyAbuseReportsTrademarkReportOwnerNotificationNone:
		return true
	}
	return false
}

// The abuse report type
type AbuseReportNewParamsBodyAbuseReportsTrademarkReportAct string

const (
	AbuseReportNewParamsBodyAbuseReportsTrademarkReportActAbuseDmca           AbuseReportNewParamsBodyAbuseReportsTrademarkReportAct = "abuse_dmca"
	AbuseReportNewParamsBodyAbuseReportsTrademarkReportActAbuseTrademark      AbuseReportNewParamsBodyAbuseReportsTrademarkReportAct = "abuse_trademark"
	AbuseReportNewParamsBodyAbuseReportsTrademarkReportActAbuseGeneral        AbuseReportNewParamsBodyAbuseReportsTrademarkReportAct = "abuse_general"
	AbuseReportNewParamsBodyAbuseReportsTrademarkReportActAbusePhishing       AbuseReportNewParamsBodyAbuseReportsTrademarkReportAct = "abuse_phishing"
	AbuseReportNewParamsBodyAbuseReportsTrademarkReportActAbuseChildren       AbuseReportNewParamsBodyAbuseReportsTrademarkReportAct = "abuse_children"
	AbuseReportNewParamsBodyAbuseReportsTrademarkReportActAbuseThreat         AbuseReportNewParamsBodyAbuseReportsTrademarkReportAct = "abuse_threat"
	AbuseReportNewParamsBodyAbuseReportsTrademarkReportActAbuseRegistrarWhois AbuseReportNewParamsBodyAbuseReportsTrademarkReportAct = "abuse_registrar_whois"
	AbuseReportNewParamsBodyAbuseReportsTrademarkReportActAbuseNcsei          AbuseReportNewParamsBodyAbuseReportsTrademarkReportAct = "abuse_ncsei"
)

func (r AbuseReportNewParamsBodyAbuseReportsTrademarkReportAct) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsTrademarkReportActAbuseDmca, AbuseReportNewParamsBodyAbuseReportsTrademarkReportActAbuseTrademark, AbuseReportNewParamsBodyAbuseReportsTrademarkReportActAbuseGeneral, AbuseReportNewParamsBodyAbuseReportsTrademarkReportActAbusePhishing, AbuseReportNewParamsBodyAbuseReportsTrademarkReportActAbuseChildren, AbuseReportNewParamsBodyAbuseReportsTrademarkReportActAbuseThreat, AbuseReportNewParamsBodyAbuseReportsTrademarkReportActAbuseRegistrarWhois, AbuseReportNewParamsBodyAbuseReportsTrademarkReportActAbuseNcsei:
		return true
	}
	return false
}

// Can be `0` for false or `1` for true. Must be value: 1 for DMCA reports
type AbuseReportNewParamsBodyAbuseReportsTrademarkReportAgree int64

const (
	AbuseReportNewParamsBodyAbuseReportsTrademarkReportAgree0 AbuseReportNewParamsBodyAbuseReportsTrademarkReportAgree = 0
	AbuseReportNewParamsBodyAbuseReportsTrademarkReportAgree1 AbuseReportNewParamsBodyAbuseReportsTrademarkReportAgree = 1
)

func (r AbuseReportNewParamsBodyAbuseReportsTrademarkReportAgree) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsTrademarkReportAgree0, AbuseReportNewParamsBodyAbuseReportsTrademarkReportAgree1:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsTrademarkReportNcmecNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsTrademarkReportNcmecNotificationSend     AbuseReportNewParamsBodyAbuseReportsTrademarkReportNcmecNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsTrademarkReportNcmecNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsTrademarkReportNcmecNotification = "send-anon"
	AbuseReportNewParamsBodyAbuseReportsTrademarkReportNcmecNotificationNone     AbuseReportNewParamsBodyAbuseReportsTrademarkReportNcmecNotification = "none"
)

func (r AbuseReportNewParamsBodyAbuseReportsTrademarkReportNcmecNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsTrademarkReportNcmecNotificationSend, AbuseReportNewParamsBodyAbuseReportsTrademarkReportNcmecNotificationSendAnon, AbuseReportNewParamsBodyAbuseReportsTrademarkReportNcmecNotificationNone:
		return true
	}
	return false
}

type AbuseReportNewParamsBodyAbuseReportsGeneralReport struct {
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	HostNotification param.Field[AbuseReportNewParamsBodyAbuseReportsGeneralReportHostNotification] `json:"host_notification,required"`
	// A detailed description of the infringement, including any necessary access
	// details and the exact steps needed to view the content, not exceeding 5000
	// characters
	Justification param.Field[string] `json:"justification,required"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	OwnerNotification param.Field[AbuseReportNewParamsBodyAbuseReportsGeneralReportOwnerNotification] `json:"owner_notification,required"`
	// The abuse report type
	Act param.Field[AbuseReportNewParamsBodyAbuseReportsGeneralReportAct] `json:"act"`
	// Text not exceeding 100 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Address1 param.Field[string] `json:"address1"`
	// The name of the copyright holder. Text not exceeding 60 characters. This field
	// may be released by Cloudflare to third parties such as the Lumen Database
	// (https://lumendatabase.org/).
	AgentName param.Field[string] `json:"agent_name"`
	// Can be `0` for false or `1` for true. Must be value: 1 for DMCA reports
	Agree param.Field[AbuseReportNewParamsBodyAbuseReportsGeneralReportAgree] `json:"agree"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	City param.Field[string] `json:"city"`
	// Any additional comments about the infringement not exceeding 2000 characters
	Comments param.Field[string] `json:"comments"`
	// Text not exceeding 100 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Company param.Field[string] `json:"company"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Country param.Field[string] `json:"country"`
	// A list of IP addresses separated by ‘ ’ (new line character). The list of
	// destination IPs should not exceed 30 IP addresses. Each one of the IP addresses
	// ought to be unique
	DestinationIPs param.Field[string] `json:"destination_ips"`
	// A valid email of the abuse reporter. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Email param.Field[string] `json:"email"`
	// Should match the value provided in `email`
	Email2 param.Field[string] `json:"email2"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Name param.Field[string] `json:"name"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	NcmecNotification param.Field[AbuseReportNewParamsBodyAbuseReportsGeneralReportNcmecNotification] `json:"ncmec_notification"`
	// If the submitter is the target of NCSEI in the URLs of the abuse report.
	NcseiSubjectRepresentation param.Field[bool] `json:"ncsei_subject_representation"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
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
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	State param.Field[string] `json:"state"`
	// Text not exceeding 20 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Tele param.Field[string] `json:"tele"`
	// Text not exceeding 255 characters
	Title param.Field[string] `json:"title"`
	// Text not exceeding 1000 characters
	TrademarkNumber param.Field[string] `json:"trademark_number"`
	// Text not exceeding 1000 characters
	TrademarkOffice param.Field[string] `json:"trademark_office"`
	// Text not exceeding 1000 characters
	TrademarkSymbol param.Field[string] `json:"trademark_symbol"`
	// A list of valid URLs separated by ‘ ’ (new line character). The list of the URLs
	// should not exceed 250 URLs. All URLs should have the same hostname. Each URL
	// should be unique. This field may be released by Cloudflare to third parties such
	// as the Lumen Database (https://lumendatabase.org/).
	URLs param.Field[string] `json:"urls"`
}

func (r AbuseReportNewParamsBodyAbuseReportsGeneralReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AbuseReportNewParamsBodyAbuseReportsGeneralReport) implementsAbuseReportNewParamsBodyUnion() {
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsGeneralReportHostNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsGeneralReportHostNotificationSend     AbuseReportNewParamsBodyAbuseReportsGeneralReportHostNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsGeneralReportHostNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsGeneralReportHostNotification = "send-anon"
	AbuseReportNewParamsBodyAbuseReportsGeneralReportHostNotificationNone     AbuseReportNewParamsBodyAbuseReportsGeneralReportHostNotification = "none"
)

func (r AbuseReportNewParamsBodyAbuseReportsGeneralReportHostNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsGeneralReportHostNotificationSend, AbuseReportNewParamsBodyAbuseReportsGeneralReportHostNotificationSendAnon, AbuseReportNewParamsBodyAbuseReportsGeneralReportHostNotificationNone:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsGeneralReportOwnerNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsGeneralReportOwnerNotificationSend     AbuseReportNewParamsBodyAbuseReportsGeneralReportOwnerNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsGeneralReportOwnerNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsGeneralReportOwnerNotification = "send-anon"
	AbuseReportNewParamsBodyAbuseReportsGeneralReportOwnerNotificationNone     AbuseReportNewParamsBodyAbuseReportsGeneralReportOwnerNotification = "none"
)

func (r AbuseReportNewParamsBodyAbuseReportsGeneralReportOwnerNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsGeneralReportOwnerNotificationSend, AbuseReportNewParamsBodyAbuseReportsGeneralReportOwnerNotificationSendAnon, AbuseReportNewParamsBodyAbuseReportsGeneralReportOwnerNotificationNone:
		return true
	}
	return false
}

// The abuse report type
type AbuseReportNewParamsBodyAbuseReportsGeneralReportAct string

const (
	AbuseReportNewParamsBodyAbuseReportsGeneralReportActAbuseDmca           AbuseReportNewParamsBodyAbuseReportsGeneralReportAct = "abuse_dmca"
	AbuseReportNewParamsBodyAbuseReportsGeneralReportActAbuseTrademark      AbuseReportNewParamsBodyAbuseReportsGeneralReportAct = "abuse_trademark"
	AbuseReportNewParamsBodyAbuseReportsGeneralReportActAbuseGeneral        AbuseReportNewParamsBodyAbuseReportsGeneralReportAct = "abuse_general"
	AbuseReportNewParamsBodyAbuseReportsGeneralReportActAbusePhishing       AbuseReportNewParamsBodyAbuseReportsGeneralReportAct = "abuse_phishing"
	AbuseReportNewParamsBodyAbuseReportsGeneralReportActAbuseChildren       AbuseReportNewParamsBodyAbuseReportsGeneralReportAct = "abuse_children"
	AbuseReportNewParamsBodyAbuseReportsGeneralReportActAbuseThreat         AbuseReportNewParamsBodyAbuseReportsGeneralReportAct = "abuse_threat"
	AbuseReportNewParamsBodyAbuseReportsGeneralReportActAbuseRegistrarWhois AbuseReportNewParamsBodyAbuseReportsGeneralReportAct = "abuse_registrar_whois"
	AbuseReportNewParamsBodyAbuseReportsGeneralReportActAbuseNcsei          AbuseReportNewParamsBodyAbuseReportsGeneralReportAct = "abuse_ncsei"
)

func (r AbuseReportNewParamsBodyAbuseReportsGeneralReportAct) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsGeneralReportActAbuseDmca, AbuseReportNewParamsBodyAbuseReportsGeneralReportActAbuseTrademark, AbuseReportNewParamsBodyAbuseReportsGeneralReportActAbuseGeneral, AbuseReportNewParamsBodyAbuseReportsGeneralReportActAbusePhishing, AbuseReportNewParamsBodyAbuseReportsGeneralReportActAbuseChildren, AbuseReportNewParamsBodyAbuseReportsGeneralReportActAbuseThreat, AbuseReportNewParamsBodyAbuseReportsGeneralReportActAbuseRegistrarWhois, AbuseReportNewParamsBodyAbuseReportsGeneralReportActAbuseNcsei:
		return true
	}
	return false
}

// Can be `0` for false or `1` for true. Must be value: 1 for DMCA reports
type AbuseReportNewParamsBodyAbuseReportsGeneralReportAgree int64

const (
	AbuseReportNewParamsBodyAbuseReportsGeneralReportAgree0 AbuseReportNewParamsBodyAbuseReportsGeneralReportAgree = 0
	AbuseReportNewParamsBodyAbuseReportsGeneralReportAgree1 AbuseReportNewParamsBodyAbuseReportsGeneralReportAgree = 1
)

func (r AbuseReportNewParamsBodyAbuseReportsGeneralReportAgree) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsGeneralReportAgree0, AbuseReportNewParamsBodyAbuseReportsGeneralReportAgree1:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsGeneralReportNcmecNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsGeneralReportNcmecNotificationSend     AbuseReportNewParamsBodyAbuseReportsGeneralReportNcmecNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsGeneralReportNcmecNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsGeneralReportNcmecNotification = "send-anon"
	AbuseReportNewParamsBodyAbuseReportsGeneralReportNcmecNotificationNone     AbuseReportNewParamsBodyAbuseReportsGeneralReportNcmecNotification = "none"
)

func (r AbuseReportNewParamsBodyAbuseReportsGeneralReportNcmecNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsGeneralReportNcmecNotificationSend, AbuseReportNewParamsBodyAbuseReportsGeneralReportNcmecNotificationSendAnon, AbuseReportNewParamsBodyAbuseReportsGeneralReportNcmecNotificationNone:
		return true
	}
	return false
}

type AbuseReportNewParamsBodyAbuseReportsPhishingReport struct {
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	HostNotification param.Field[AbuseReportNewParamsBodyAbuseReportsPhishingReportHostNotification] `json:"host_notification,required"`
	// A detailed description of the infringement, including any necessary access
	// details and the exact steps needed to view the content, not exceeding 5000
	// characters
	Justification param.Field[string] `json:"justification,required"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	OwnerNotification param.Field[AbuseReportNewParamsBodyAbuseReportsPhishingReportOwnerNotification] `json:"owner_notification,required"`
	// The abuse report type
	Act param.Field[AbuseReportNewParamsBodyAbuseReportsPhishingReportAct] `json:"act"`
	// Text not exceeding 100 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Address1 param.Field[string] `json:"address1"`
	// The name of the copyright holder. Text not exceeding 60 characters. This field
	// may be released by Cloudflare to third parties such as the Lumen Database
	// (https://lumendatabase.org/).
	AgentName param.Field[string] `json:"agent_name"`
	// Can be `0` for false or `1` for true. Must be value: 1 for DMCA reports
	Agree param.Field[AbuseReportNewParamsBodyAbuseReportsPhishingReportAgree] `json:"agree"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	City param.Field[string] `json:"city"`
	// Any additional comments about the infringement not exceeding 2000 characters
	Comments param.Field[string] `json:"comments"`
	// Text not exceeding 100 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Company param.Field[string] `json:"company"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Country param.Field[string] `json:"country"`
	// A list of IP addresses separated by ‘ ’ (new line character). The list of
	// destination IPs should not exceed 30 IP addresses. Each one of the IP addresses
	// ought to be unique
	DestinationIPs param.Field[string] `json:"destination_ips"`
	// A valid email of the abuse reporter. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Email param.Field[string] `json:"email"`
	// Should match the value provided in `email`
	Email2 param.Field[string] `json:"email2"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Name param.Field[string] `json:"name"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	NcmecNotification param.Field[AbuseReportNewParamsBodyAbuseReportsPhishingReportNcmecNotification] `json:"ncmec_notification"`
	// If the submitter is the target of NCSEI in the URLs of the abuse report.
	NcseiSubjectRepresentation param.Field[bool] `json:"ncsei_subject_representation"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
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
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	State param.Field[string] `json:"state"`
	// Text not exceeding 20 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Tele param.Field[string] `json:"tele"`
	// Text not exceeding 255 characters
	Title param.Field[string] `json:"title"`
	// Text not exceeding 1000 characters
	TrademarkNumber param.Field[string] `json:"trademark_number"`
	// Text not exceeding 1000 characters
	TrademarkOffice param.Field[string] `json:"trademark_office"`
	// Text not exceeding 1000 characters
	TrademarkSymbol param.Field[string] `json:"trademark_symbol"`
	// A list of valid URLs separated by ‘ ’ (new line character). The list of the URLs
	// should not exceed 250 URLs. All URLs should have the same hostname. Each URL
	// should be unique. This field may be released by Cloudflare to third parties such
	// as the Lumen Database (https://lumendatabase.org/).
	URLs param.Field[string] `json:"urls"`
}

func (r AbuseReportNewParamsBodyAbuseReportsPhishingReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AbuseReportNewParamsBodyAbuseReportsPhishingReport) implementsAbuseReportNewParamsBodyUnion() {
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsPhishingReportHostNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsPhishingReportHostNotificationSend     AbuseReportNewParamsBodyAbuseReportsPhishingReportHostNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsPhishingReportHostNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsPhishingReportHostNotification = "send-anon"
	AbuseReportNewParamsBodyAbuseReportsPhishingReportHostNotificationNone     AbuseReportNewParamsBodyAbuseReportsPhishingReportHostNotification = "none"
)

func (r AbuseReportNewParamsBodyAbuseReportsPhishingReportHostNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsPhishingReportHostNotificationSend, AbuseReportNewParamsBodyAbuseReportsPhishingReportHostNotificationSendAnon, AbuseReportNewParamsBodyAbuseReportsPhishingReportHostNotificationNone:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsPhishingReportOwnerNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsPhishingReportOwnerNotificationSend     AbuseReportNewParamsBodyAbuseReportsPhishingReportOwnerNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsPhishingReportOwnerNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsPhishingReportOwnerNotification = "send-anon"
	AbuseReportNewParamsBodyAbuseReportsPhishingReportOwnerNotificationNone     AbuseReportNewParamsBodyAbuseReportsPhishingReportOwnerNotification = "none"
)

func (r AbuseReportNewParamsBodyAbuseReportsPhishingReportOwnerNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsPhishingReportOwnerNotificationSend, AbuseReportNewParamsBodyAbuseReportsPhishingReportOwnerNotificationSendAnon, AbuseReportNewParamsBodyAbuseReportsPhishingReportOwnerNotificationNone:
		return true
	}
	return false
}

// The abuse report type
type AbuseReportNewParamsBodyAbuseReportsPhishingReportAct string

const (
	AbuseReportNewParamsBodyAbuseReportsPhishingReportActAbuseDmca           AbuseReportNewParamsBodyAbuseReportsPhishingReportAct = "abuse_dmca"
	AbuseReportNewParamsBodyAbuseReportsPhishingReportActAbuseTrademark      AbuseReportNewParamsBodyAbuseReportsPhishingReportAct = "abuse_trademark"
	AbuseReportNewParamsBodyAbuseReportsPhishingReportActAbuseGeneral        AbuseReportNewParamsBodyAbuseReportsPhishingReportAct = "abuse_general"
	AbuseReportNewParamsBodyAbuseReportsPhishingReportActAbusePhishing       AbuseReportNewParamsBodyAbuseReportsPhishingReportAct = "abuse_phishing"
	AbuseReportNewParamsBodyAbuseReportsPhishingReportActAbuseChildren       AbuseReportNewParamsBodyAbuseReportsPhishingReportAct = "abuse_children"
	AbuseReportNewParamsBodyAbuseReportsPhishingReportActAbuseThreat         AbuseReportNewParamsBodyAbuseReportsPhishingReportAct = "abuse_threat"
	AbuseReportNewParamsBodyAbuseReportsPhishingReportActAbuseRegistrarWhois AbuseReportNewParamsBodyAbuseReportsPhishingReportAct = "abuse_registrar_whois"
	AbuseReportNewParamsBodyAbuseReportsPhishingReportActAbuseNcsei          AbuseReportNewParamsBodyAbuseReportsPhishingReportAct = "abuse_ncsei"
)

func (r AbuseReportNewParamsBodyAbuseReportsPhishingReportAct) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsPhishingReportActAbuseDmca, AbuseReportNewParamsBodyAbuseReportsPhishingReportActAbuseTrademark, AbuseReportNewParamsBodyAbuseReportsPhishingReportActAbuseGeneral, AbuseReportNewParamsBodyAbuseReportsPhishingReportActAbusePhishing, AbuseReportNewParamsBodyAbuseReportsPhishingReportActAbuseChildren, AbuseReportNewParamsBodyAbuseReportsPhishingReportActAbuseThreat, AbuseReportNewParamsBodyAbuseReportsPhishingReportActAbuseRegistrarWhois, AbuseReportNewParamsBodyAbuseReportsPhishingReportActAbuseNcsei:
		return true
	}
	return false
}

// Can be `0` for false or `1` for true. Must be value: 1 for DMCA reports
type AbuseReportNewParamsBodyAbuseReportsPhishingReportAgree int64

const (
	AbuseReportNewParamsBodyAbuseReportsPhishingReportAgree0 AbuseReportNewParamsBodyAbuseReportsPhishingReportAgree = 0
	AbuseReportNewParamsBodyAbuseReportsPhishingReportAgree1 AbuseReportNewParamsBodyAbuseReportsPhishingReportAgree = 1
)

func (r AbuseReportNewParamsBodyAbuseReportsPhishingReportAgree) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsPhishingReportAgree0, AbuseReportNewParamsBodyAbuseReportsPhishingReportAgree1:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsPhishingReportNcmecNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsPhishingReportNcmecNotificationSend     AbuseReportNewParamsBodyAbuseReportsPhishingReportNcmecNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsPhishingReportNcmecNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsPhishingReportNcmecNotification = "send-anon"
	AbuseReportNewParamsBodyAbuseReportsPhishingReportNcmecNotificationNone     AbuseReportNewParamsBodyAbuseReportsPhishingReportNcmecNotification = "none"
)

func (r AbuseReportNewParamsBodyAbuseReportsPhishingReportNcmecNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsPhishingReportNcmecNotificationSend, AbuseReportNewParamsBodyAbuseReportsPhishingReportNcmecNotificationSendAnon, AbuseReportNewParamsBodyAbuseReportsPhishingReportNcmecNotificationNone:
		return true
	}
	return false
}

type AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReport struct {
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	HostNotification param.Field[AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportHostNotification] `json:"host_notification,required"`
	// A detailed description of the infringement, including any necessary access
	// details and the exact steps needed to view the content, not exceeding 5000
	// characters
	Justification param.Field[string] `json:"justification,required"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	NcmecNotification param.Field[AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportNcmecNotification] `json:"ncmec_notification,required"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	OwnerNotification param.Field[AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportOwnerNotification] `json:"owner_notification,required"`
	// The abuse report type
	Act param.Field[AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportAct] `json:"act"`
	// Text not exceeding 100 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Address1 param.Field[string] `json:"address1"`
	// The name of the copyright holder. Text not exceeding 60 characters. This field
	// may be released by Cloudflare to third parties such as the Lumen Database
	// (https://lumendatabase.org/).
	AgentName param.Field[string] `json:"agent_name"`
	// Can be `0` for false or `1` for true. Must be value: 1 for DMCA reports
	Agree param.Field[AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportAgree] `json:"agree"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	City param.Field[string] `json:"city"`
	// Any additional comments about the infringement not exceeding 2000 characters
	Comments param.Field[string] `json:"comments"`
	// Text not exceeding 100 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Company param.Field[string] `json:"company"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Country param.Field[string] `json:"country"`
	// A list of IP addresses separated by ‘ ’ (new line character). The list of
	// destination IPs should not exceed 30 IP addresses. Each one of the IP addresses
	// ought to be unique
	DestinationIPs param.Field[string] `json:"destination_ips"`
	// A valid email of the abuse reporter. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Email param.Field[string] `json:"email"`
	// Should match the value provided in `email`
	Email2 param.Field[string] `json:"email2"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Name param.Field[string] `json:"name"`
	// If the submitter is the target of NCSEI in the URLs of the abuse report.
	NcseiSubjectRepresentation param.Field[bool] `json:"ncsei_subject_representation"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
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
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	State param.Field[string] `json:"state"`
	// Text not exceeding 20 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Tele param.Field[string] `json:"tele"`
	// Text not exceeding 255 characters
	Title param.Field[string] `json:"title"`
	// Text not exceeding 1000 characters
	TrademarkNumber param.Field[string] `json:"trademark_number"`
	// Text not exceeding 1000 characters
	TrademarkOffice param.Field[string] `json:"trademark_office"`
	// Text not exceeding 1000 characters
	TrademarkSymbol param.Field[string] `json:"trademark_symbol"`
	// A list of valid URLs separated by ‘ ’ (new line character). The list of the URLs
	// should not exceed 250 URLs. All URLs should have the same hostname. Each URL
	// should be unique. This field may be released by Cloudflare to third parties such
	// as the Lumen Database (https://lumendatabase.org/).
	URLs param.Field[string] `json:"urls"`
}

func (r AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReport) implementsAbuseReportNewParamsBodyUnion() {
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportHostNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportHostNotificationSend     AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportHostNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportHostNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportHostNotification = "send-anon"
	AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportHostNotificationNone     AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportHostNotification = "none"
)

func (r AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportHostNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportHostNotificationSend, AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportHostNotificationSendAnon, AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportHostNotificationNone:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportNcmecNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportNcmecNotificationSend     AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportNcmecNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportNcmecNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportNcmecNotification = "send-anon"
	AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportNcmecNotificationNone     AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportNcmecNotification = "none"
)

func (r AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportNcmecNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportNcmecNotificationSend, AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportNcmecNotificationSendAnon, AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportNcmecNotificationNone:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportOwnerNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportOwnerNotificationSend     AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportOwnerNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportOwnerNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportOwnerNotification = "send-anon"
	AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportOwnerNotificationNone     AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportOwnerNotification = "none"
)

func (r AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportOwnerNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportOwnerNotificationSend, AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportOwnerNotificationSendAnon, AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportOwnerNotificationNone:
		return true
	}
	return false
}

// The abuse report type
type AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportAct string

const (
	AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportActAbuseDmca           AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportAct = "abuse_dmca"
	AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportActAbuseTrademark      AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportAct = "abuse_trademark"
	AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportActAbuseGeneral        AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportAct = "abuse_general"
	AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportActAbusePhishing       AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportAct = "abuse_phishing"
	AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportActAbuseChildren       AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportAct = "abuse_children"
	AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportActAbuseThreat         AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportAct = "abuse_threat"
	AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportActAbuseRegistrarWhois AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportAct = "abuse_registrar_whois"
	AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportActAbuseNcsei          AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportAct = "abuse_ncsei"
)

func (r AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportAct) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportActAbuseDmca, AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportActAbuseTrademark, AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportActAbuseGeneral, AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportActAbusePhishing, AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportActAbuseChildren, AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportActAbuseThreat, AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportActAbuseRegistrarWhois, AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportActAbuseNcsei:
		return true
	}
	return false
}

// Can be `0` for false or `1` for true. Must be value: 1 for DMCA reports
type AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportAgree int64

const (
	AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportAgree0 AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportAgree = 0
	AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportAgree1 AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportAgree = 1
)

func (r AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportAgree) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportAgree0, AbuseReportNewParamsBodyAbuseReportsChildrenAbuseReportAgree1:
		return true
	}
	return false
}

type AbuseReportNewParamsBodyAbuseReportsThreatReport struct {
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	HostNotification param.Field[AbuseReportNewParamsBodyAbuseReportsThreatReportHostNotification] `json:"host_notification,required"`
	// A detailed description of the infringement, including any necessary access
	// details and the exact steps needed to view the content, not exceeding 5000
	// characters
	Justification param.Field[string] `json:"justification,required"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	OwnerNotification param.Field[AbuseReportNewParamsBodyAbuseReportsThreatReportOwnerNotification] `json:"owner_notification,required"`
	// The abuse report type
	Act param.Field[AbuseReportNewParamsBodyAbuseReportsThreatReportAct] `json:"act"`
	// Text not exceeding 100 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Address1 param.Field[string] `json:"address1"`
	// The name of the copyright holder. Text not exceeding 60 characters. This field
	// may be released by Cloudflare to third parties such as the Lumen Database
	// (https://lumendatabase.org/).
	AgentName param.Field[string] `json:"agent_name"`
	// Can be `0` for false or `1` for true. Must be value: 1 for DMCA reports
	Agree param.Field[AbuseReportNewParamsBodyAbuseReportsThreatReportAgree] `json:"agree"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	City param.Field[string] `json:"city"`
	// Any additional comments about the infringement not exceeding 2000 characters
	Comments param.Field[string] `json:"comments"`
	// Text not exceeding 100 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Company param.Field[string] `json:"company"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Country param.Field[string] `json:"country"`
	// A list of IP addresses separated by ‘ ’ (new line character). The list of
	// destination IPs should not exceed 30 IP addresses. Each one of the IP addresses
	// ought to be unique
	DestinationIPs param.Field[string] `json:"destination_ips"`
	// A valid email of the abuse reporter. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Email param.Field[string] `json:"email"`
	// Should match the value provided in `email`
	Email2 param.Field[string] `json:"email2"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Name param.Field[string] `json:"name"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	NcmecNotification param.Field[AbuseReportNewParamsBodyAbuseReportsThreatReportNcmecNotification] `json:"ncmec_notification"`
	// If the submitter is the target of NCSEI in the URLs of the abuse report.
	NcseiSubjectRepresentation param.Field[bool] `json:"ncsei_subject_representation"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
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
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	State param.Field[string] `json:"state"`
	// Text not exceeding 20 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Tele param.Field[string] `json:"tele"`
	// Text not exceeding 255 characters
	Title param.Field[string] `json:"title"`
	// Text not exceeding 1000 characters
	TrademarkNumber param.Field[string] `json:"trademark_number"`
	// Text not exceeding 1000 characters
	TrademarkOffice param.Field[string] `json:"trademark_office"`
	// Text not exceeding 1000 characters
	TrademarkSymbol param.Field[string] `json:"trademark_symbol"`
	// A list of valid URLs separated by ‘ ’ (new line character). The list of the URLs
	// should not exceed 250 URLs. All URLs should have the same hostname. Each URL
	// should be unique. This field may be released by Cloudflare to third parties such
	// as the Lumen Database (https://lumendatabase.org/).
	URLs param.Field[string] `json:"urls"`
}

func (r AbuseReportNewParamsBodyAbuseReportsThreatReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AbuseReportNewParamsBodyAbuseReportsThreatReport) implementsAbuseReportNewParamsBodyUnion() {}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsThreatReportHostNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsThreatReportHostNotificationSend     AbuseReportNewParamsBodyAbuseReportsThreatReportHostNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsThreatReportHostNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsThreatReportHostNotification = "send-anon"
	AbuseReportNewParamsBodyAbuseReportsThreatReportHostNotificationNone     AbuseReportNewParamsBodyAbuseReportsThreatReportHostNotification = "none"
)

func (r AbuseReportNewParamsBodyAbuseReportsThreatReportHostNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsThreatReportHostNotificationSend, AbuseReportNewParamsBodyAbuseReportsThreatReportHostNotificationSendAnon, AbuseReportNewParamsBodyAbuseReportsThreatReportHostNotificationNone:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsThreatReportOwnerNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsThreatReportOwnerNotificationSend     AbuseReportNewParamsBodyAbuseReportsThreatReportOwnerNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsThreatReportOwnerNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsThreatReportOwnerNotification = "send-anon"
	AbuseReportNewParamsBodyAbuseReportsThreatReportOwnerNotificationNone     AbuseReportNewParamsBodyAbuseReportsThreatReportOwnerNotification = "none"
)

func (r AbuseReportNewParamsBodyAbuseReportsThreatReportOwnerNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsThreatReportOwnerNotificationSend, AbuseReportNewParamsBodyAbuseReportsThreatReportOwnerNotificationSendAnon, AbuseReportNewParamsBodyAbuseReportsThreatReportOwnerNotificationNone:
		return true
	}
	return false
}

// The abuse report type
type AbuseReportNewParamsBodyAbuseReportsThreatReportAct string

const (
	AbuseReportNewParamsBodyAbuseReportsThreatReportActAbuseDmca           AbuseReportNewParamsBodyAbuseReportsThreatReportAct = "abuse_dmca"
	AbuseReportNewParamsBodyAbuseReportsThreatReportActAbuseTrademark      AbuseReportNewParamsBodyAbuseReportsThreatReportAct = "abuse_trademark"
	AbuseReportNewParamsBodyAbuseReportsThreatReportActAbuseGeneral        AbuseReportNewParamsBodyAbuseReportsThreatReportAct = "abuse_general"
	AbuseReportNewParamsBodyAbuseReportsThreatReportActAbusePhishing       AbuseReportNewParamsBodyAbuseReportsThreatReportAct = "abuse_phishing"
	AbuseReportNewParamsBodyAbuseReportsThreatReportActAbuseChildren       AbuseReportNewParamsBodyAbuseReportsThreatReportAct = "abuse_children"
	AbuseReportNewParamsBodyAbuseReportsThreatReportActAbuseThreat         AbuseReportNewParamsBodyAbuseReportsThreatReportAct = "abuse_threat"
	AbuseReportNewParamsBodyAbuseReportsThreatReportActAbuseRegistrarWhois AbuseReportNewParamsBodyAbuseReportsThreatReportAct = "abuse_registrar_whois"
	AbuseReportNewParamsBodyAbuseReportsThreatReportActAbuseNcsei          AbuseReportNewParamsBodyAbuseReportsThreatReportAct = "abuse_ncsei"
)

func (r AbuseReportNewParamsBodyAbuseReportsThreatReportAct) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsThreatReportActAbuseDmca, AbuseReportNewParamsBodyAbuseReportsThreatReportActAbuseTrademark, AbuseReportNewParamsBodyAbuseReportsThreatReportActAbuseGeneral, AbuseReportNewParamsBodyAbuseReportsThreatReportActAbusePhishing, AbuseReportNewParamsBodyAbuseReportsThreatReportActAbuseChildren, AbuseReportNewParamsBodyAbuseReportsThreatReportActAbuseThreat, AbuseReportNewParamsBodyAbuseReportsThreatReportActAbuseRegistrarWhois, AbuseReportNewParamsBodyAbuseReportsThreatReportActAbuseNcsei:
		return true
	}
	return false
}

// Can be `0` for false or `1` for true. Must be value: 1 for DMCA reports
type AbuseReportNewParamsBodyAbuseReportsThreatReportAgree int64

const (
	AbuseReportNewParamsBodyAbuseReportsThreatReportAgree0 AbuseReportNewParamsBodyAbuseReportsThreatReportAgree = 0
	AbuseReportNewParamsBodyAbuseReportsThreatReportAgree1 AbuseReportNewParamsBodyAbuseReportsThreatReportAgree = 1
)

func (r AbuseReportNewParamsBodyAbuseReportsThreatReportAgree) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsThreatReportAgree0, AbuseReportNewParamsBodyAbuseReportsThreatReportAgree1:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsThreatReportNcmecNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsThreatReportNcmecNotificationSend     AbuseReportNewParamsBodyAbuseReportsThreatReportNcmecNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsThreatReportNcmecNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsThreatReportNcmecNotification = "send-anon"
	AbuseReportNewParamsBodyAbuseReportsThreatReportNcmecNotificationNone     AbuseReportNewParamsBodyAbuseReportsThreatReportNcmecNotification = "none"
)

func (r AbuseReportNewParamsBodyAbuseReportsThreatReportNcmecNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsThreatReportNcmecNotificationSend, AbuseReportNewParamsBodyAbuseReportsThreatReportNcmecNotificationSendAnon, AbuseReportNewParamsBodyAbuseReportsThreatReportNcmecNotificationNone:
		return true
	}
	return false
}

type AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReport struct {
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	OwnerNotification param.Field[AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportOwnerNotification] `json:"owner_notification,required"`
	// The abuse report type
	Act param.Field[AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportAct] `json:"act"`
	// Text not exceeding 100 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Address1 param.Field[string] `json:"address1"`
	// The name of the copyright holder. Text not exceeding 60 characters. This field
	// may be released by Cloudflare to third parties such as the Lumen Database
	// (https://lumendatabase.org/).
	AgentName param.Field[string] `json:"agent_name"`
	// Can be `0` for false or `1` for true. Must be value: 1 for DMCA reports
	Agree param.Field[AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportAgree] `json:"agree"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	City param.Field[string] `json:"city"`
	// Any additional comments about the infringement not exceeding 2000 characters
	Comments param.Field[string] `json:"comments"`
	// Text not exceeding 100 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Company param.Field[string] `json:"company"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Country param.Field[string] `json:"country"`
	// A list of IP addresses separated by ‘ ’ (new line character). The list of
	// destination IPs should not exceed 30 IP addresses. Each one of the IP addresses
	// ought to be unique
	DestinationIPs param.Field[string] `json:"destination_ips"`
	// A valid email of the abuse reporter. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Email param.Field[string] `json:"email"`
	// Should match the value provided in `email`
	Email2 param.Field[string] `json:"email2"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	HostNotification param.Field[AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportHostNotification] `json:"host_notification"`
	// A detailed description of the infringement, including any necessary access
	// details and the exact steps needed to view the content, not exceeding 5000
	// characters
	Justification param.Field[string] `json:"justification"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Name param.Field[string] `json:"name"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	NcmecNotification param.Field[AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportNcmecNotification] `json:"ncmec_notification"`
	// If the submitter is the target of NCSEI in the URLs of the abuse report.
	NcseiSubjectRepresentation param.Field[bool] `json:"ncsei_subject_representation"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
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
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	State param.Field[string] `json:"state"`
	// Text not exceeding 20 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Tele param.Field[string] `json:"tele"`
	// Text not exceeding 255 characters
	Title param.Field[string] `json:"title"`
	// Text not exceeding 1000 characters
	TrademarkNumber param.Field[string] `json:"trademark_number"`
	// Text not exceeding 1000 characters
	TrademarkOffice param.Field[string] `json:"trademark_office"`
	// Text not exceeding 1000 characters
	TrademarkSymbol param.Field[string] `json:"trademark_symbol"`
	// A list of valid URLs separated by ‘ ’ (new line character). The list of the URLs
	// should not exceed 250 URLs. All URLs should have the same hostname. Each URL
	// should be unique. This field may be released by Cloudflare to third parties such
	// as the Lumen Database (https://lumendatabase.org/).
	URLs param.Field[string] `json:"urls"`
}

func (r AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReport) implementsAbuseReportNewParamsBodyUnion() {
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportOwnerNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportOwnerNotificationSend     AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportOwnerNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportOwnerNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportOwnerNotification = "send-anon"
	AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportOwnerNotificationNone     AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportOwnerNotification = "none"
)

func (r AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportOwnerNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportOwnerNotificationSend, AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportOwnerNotificationSendAnon, AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportOwnerNotificationNone:
		return true
	}
	return false
}

// The abuse report type
type AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportAct string

const (
	AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportActAbuseDmca           AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportAct = "abuse_dmca"
	AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportActAbuseTrademark      AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportAct = "abuse_trademark"
	AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportActAbuseGeneral        AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportAct = "abuse_general"
	AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportActAbusePhishing       AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportAct = "abuse_phishing"
	AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportActAbuseChildren       AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportAct = "abuse_children"
	AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportActAbuseThreat         AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportAct = "abuse_threat"
	AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportActAbuseRegistrarWhois AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportAct = "abuse_registrar_whois"
	AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportActAbuseNcsei          AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportAct = "abuse_ncsei"
)

func (r AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportAct) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportActAbuseDmca, AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportActAbuseTrademark, AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportActAbuseGeneral, AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportActAbusePhishing, AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportActAbuseChildren, AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportActAbuseThreat, AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportActAbuseRegistrarWhois, AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportActAbuseNcsei:
		return true
	}
	return false
}

// Can be `0` for false or `1` for true. Must be value: 1 for DMCA reports
type AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportAgree int64

const (
	AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportAgree0 AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportAgree = 0
	AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportAgree1 AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportAgree = 1
)

func (r AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportAgree) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportAgree0, AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportAgree1:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportHostNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportHostNotificationSend     AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportHostNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportHostNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportHostNotification = "send-anon"
	AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportHostNotificationNone     AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportHostNotification = "none"
)

func (r AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportHostNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportHostNotificationSend, AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportHostNotificationSendAnon, AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportHostNotificationNone:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportNcmecNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportNcmecNotificationSend     AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportNcmecNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportNcmecNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportNcmecNotification = "send-anon"
	AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportNcmecNotificationNone     AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportNcmecNotification = "none"
)

func (r AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportNcmecNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportNcmecNotificationSend, AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportNcmecNotificationSendAnon, AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportNcmecNotificationNone:
		return true
	}
	return false
}

type AbuseReportNewParamsBodyAbuseReportsNcseiReport struct {
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	HostNotification param.Field[AbuseReportNewParamsBodyAbuseReportsNcseiReportHostNotification] `json:"host_notification,required"`
	// If the submitter is the target of NCSEI in the URLs of the abuse report.
	NcseiSubjectRepresentation param.Field[bool] `json:"ncsei_subject_representation,required"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	OwnerNotification param.Field[AbuseReportNewParamsBodyAbuseReportsNcseiReportOwnerNotification] `json:"owner_notification,required"`
	// The abuse report type
	Act param.Field[AbuseReportNewParamsBodyAbuseReportsNcseiReportAct] `json:"act"`
	// Text not exceeding 100 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Address1 param.Field[string] `json:"address1"`
	// The name of the copyright holder. Text not exceeding 60 characters. This field
	// may be released by Cloudflare to third parties such as the Lumen Database
	// (https://lumendatabase.org/).
	AgentName param.Field[string] `json:"agent_name"`
	// Can be `0` for false or `1` for true. Must be value: 1 for DMCA reports
	Agree param.Field[AbuseReportNewParamsBodyAbuseReportsNcseiReportAgree] `json:"agree"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	City param.Field[string] `json:"city"`
	// Any additional comments about the infringement not exceeding 2000 characters
	Comments param.Field[string] `json:"comments"`
	// Text not exceeding 100 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Company param.Field[string] `json:"company"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Country param.Field[string] `json:"country"`
	// A list of IP addresses separated by ‘ ’ (new line character). The list of
	// destination IPs should not exceed 30 IP addresses. Each one of the IP addresses
	// ought to be unique
	DestinationIPs param.Field[string] `json:"destination_ips"`
	// A valid email of the abuse reporter. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Email param.Field[string] `json:"email"`
	// Should match the value provided in `email`
	Email2 param.Field[string] `json:"email2"`
	// A detailed description of the infringement, including any necessary access
	// details and the exact steps needed to view the content, not exceeding 5000
	// characters
	Justification param.Field[string] `json:"justification"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Name param.Field[string] `json:"name"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	NcmecNotification param.Field[AbuseReportNewParamsBodyAbuseReportsNcseiReportNcmecNotification] `json:"ncmec_notification"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
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
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	State param.Field[string] `json:"state"`
	// Text not exceeding 20 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Tele param.Field[string] `json:"tele"`
	// Text not exceeding 255 characters
	Title param.Field[string] `json:"title"`
	// Text not exceeding 1000 characters
	TrademarkNumber param.Field[string] `json:"trademark_number"`
	// Text not exceeding 1000 characters
	TrademarkOffice param.Field[string] `json:"trademark_office"`
	// Text not exceeding 1000 characters
	TrademarkSymbol param.Field[string] `json:"trademark_symbol"`
	// A list of valid URLs separated by ‘ ’ (new line character). The list of the URLs
	// should not exceed 250 URLs. All URLs should have the same hostname. Each URL
	// should be unique. This field may be released by Cloudflare to third parties such
	// as the Lumen Database (https://lumendatabase.org/).
	URLs param.Field[string] `json:"urls"`
}

func (r AbuseReportNewParamsBodyAbuseReportsNcseiReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AbuseReportNewParamsBodyAbuseReportsNcseiReport) implementsAbuseReportNewParamsBodyUnion() {}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsNcseiReportHostNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsNcseiReportHostNotificationSend     AbuseReportNewParamsBodyAbuseReportsNcseiReportHostNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsNcseiReportHostNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsNcseiReportHostNotification = "send-anon"
	AbuseReportNewParamsBodyAbuseReportsNcseiReportHostNotificationNone     AbuseReportNewParamsBodyAbuseReportsNcseiReportHostNotification = "none"
)

func (r AbuseReportNewParamsBodyAbuseReportsNcseiReportHostNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsNcseiReportHostNotificationSend, AbuseReportNewParamsBodyAbuseReportsNcseiReportHostNotificationSendAnon, AbuseReportNewParamsBodyAbuseReportsNcseiReportHostNotificationNone:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsNcseiReportOwnerNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsNcseiReportOwnerNotificationSend     AbuseReportNewParamsBodyAbuseReportsNcseiReportOwnerNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsNcseiReportOwnerNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsNcseiReportOwnerNotification = "send-anon"
	AbuseReportNewParamsBodyAbuseReportsNcseiReportOwnerNotificationNone     AbuseReportNewParamsBodyAbuseReportsNcseiReportOwnerNotification = "none"
)

func (r AbuseReportNewParamsBodyAbuseReportsNcseiReportOwnerNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsNcseiReportOwnerNotificationSend, AbuseReportNewParamsBodyAbuseReportsNcseiReportOwnerNotificationSendAnon, AbuseReportNewParamsBodyAbuseReportsNcseiReportOwnerNotificationNone:
		return true
	}
	return false
}

// The abuse report type
type AbuseReportNewParamsBodyAbuseReportsNcseiReportAct string

const (
	AbuseReportNewParamsBodyAbuseReportsNcseiReportActAbuseDmca           AbuseReportNewParamsBodyAbuseReportsNcseiReportAct = "abuse_dmca"
	AbuseReportNewParamsBodyAbuseReportsNcseiReportActAbuseTrademark      AbuseReportNewParamsBodyAbuseReportsNcseiReportAct = "abuse_trademark"
	AbuseReportNewParamsBodyAbuseReportsNcseiReportActAbuseGeneral        AbuseReportNewParamsBodyAbuseReportsNcseiReportAct = "abuse_general"
	AbuseReportNewParamsBodyAbuseReportsNcseiReportActAbusePhishing       AbuseReportNewParamsBodyAbuseReportsNcseiReportAct = "abuse_phishing"
	AbuseReportNewParamsBodyAbuseReportsNcseiReportActAbuseChildren       AbuseReportNewParamsBodyAbuseReportsNcseiReportAct = "abuse_children"
	AbuseReportNewParamsBodyAbuseReportsNcseiReportActAbuseThreat         AbuseReportNewParamsBodyAbuseReportsNcseiReportAct = "abuse_threat"
	AbuseReportNewParamsBodyAbuseReportsNcseiReportActAbuseRegistrarWhois AbuseReportNewParamsBodyAbuseReportsNcseiReportAct = "abuse_registrar_whois"
	AbuseReportNewParamsBodyAbuseReportsNcseiReportActAbuseNcsei          AbuseReportNewParamsBodyAbuseReportsNcseiReportAct = "abuse_ncsei"
)

func (r AbuseReportNewParamsBodyAbuseReportsNcseiReportAct) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsNcseiReportActAbuseDmca, AbuseReportNewParamsBodyAbuseReportsNcseiReportActAbuseTrademark, AbuseReportNewParamsBodyAbuseReportsNcseiReportActAbuseGeneral, AbuseReportNewParamsBodyAbuseReportsNcseiReportActAbusePhishing, AbuseReportNewParamsBodyAbuseReportsNcseiReportActAbuseChildren, AbuseReportNewParamsBodyAbuseReportsNcseiReportActAbuseThreat, AbuseReportNewParamsBodyAbuseReportsNcseiReportActAbuseRegistrarWhois, AbuseReportNewParamsBodyAbuseReportsNcseiReportActAbuseNcsei:
		return true
	}
	return false
}

// Can be `0` for false or `1` for true. Must be value: 1 for DMCA reports
type AbuseReportNewParamsBodyAbuseReportsNcseiReportAgree int64

const (
	AbuseReportNewParamsBodyAbuseReportsNcseiReportAgree0 AbuseReportNewParamsBodyAbuseReportsNcseiReportAgree = 0
	AbuseReportNewParamsBodyAbuseReportsNcseiReportAgree1 AbuseReportNewParamsBodyAbuseReportsNcseiReportAgree = 1
)

func (r AbuseReportNewParamsBodyAbuseReportsNcseiReportAgree) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsNcseiReportAgree0, AbuseReportNewParamsBodyAbuseReportsNcseiReportAgree1:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsNcseiReportNcmecNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsNcseiReportNcmecNotificationSend     AbuseReportNewParamsBodyAbuseReportsNcseiReportNcmecNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsNcseiReportNcmecNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsNcseiReportNcmecNotification = "send-anon"
	AbuseReportNewParamsBodyAbuseReportsNcseiReportNcmecNotificationNone     AbuseReportNewParamsBodyAbuseReportsNcseiReportNcmecNotification = "none"
)

func (r AbuseReportNewParamsBodyAbuseReportsNcseiReportNcmecNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsNcseiReportNcmecNotificationSend, AbuseReportNewParamsBodyAbuseReportsNcseiReportNcmecNotificationSendAnon, AbuseReportNewParamsBodyAbuseReportsNcseiReportNcmecNotificationNone:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyOwnerNotification string

const (
	AbuseReportNewParamsBodyOwnerNotificationSend     AbuseReportNewParamsBodyOwnerNotification = "send"
	AbuseReportNewParamsBodyOwnerNotificationSendAnon AbuseReportNewParamsBodyOwnerNotification = "send-anon"
	AbuseReportNewParamsBodyOwnerNotificationNone     AbuseReportNewParamsBodyOwnerNotification = "none"
)

func (r AbuseReportNewParamsBodyOwnerNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyOwnerNotificationSend, AbuseReportNewParamsBodyOwnerNotificationSendAnon, AbuseReportNewParamsBodyOwnerNotificationNone:
		return true
	}
	return false
}

// The abuse report type
type AbuseReportNewParamsBodyAct string

const (
	AbuseReportNewParamsBodyActAbuseDmca           AbuseReportNewParamsBodyAct = "abuse_dmca"
	AbuseReportNewParamsBodyActAbuseTrademark      AbuseReportNewParamsBodyAct = "abuse_trademark"
	AbuseReportNewParamsBodyActAbuseGeneral        AbuseReportNewParamsBodyAct = "abuse_general"
	AbuseReportNewParamsBodyActAbusePhishing       AbuseReportNewParamsBodyAct = "abuse_phishing"
	AbuseReportNewParamsBodyActAbuseChildren       AbuseReportNewParamsBodyAct = "abuse_children"
	AbuseReportNewParamsBodyActAbuseThreat         AbuseReportNewParamsBodyAct = "abuse_threat"
	AbuseReportNewParamsBodyActAbuseRegistrarWhois AbuseReportNewParamsBodyAct = "abuse_registrar_whois"
	AbuseReportNewParamsBodyActAbuseNcsei          AbuseReportNewParamsBodyAct = "abuse_ncsei"
)

func (r AbuseReportNewParamsBodyAct) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyActAbuseDmca, AbuseReportNewParamsBodyActAbuseTrademark, AbuseReportNewParamsBodyActAbuseGeneral, AbuseReportNewParamsBodyActAbusePhishing, AbuseReportNewParamsBodyActAbuseChildren, AbuseReportNewParamsBodyActAbuseThreat, AbuseReportNewParamsBodyActAbuseRegistrarWhois, AbuseReportNewParamsBodyActAbuseNcsei:
		return true
	}
	return false
}

// Can be `0` for false or `1` for true. Must be value: 1 for DMCA reports
type AbuseReportNewParamsBodyAgree int64

const (
	AbuseReportNewParamsBodyAgree0 AbuseReportNewParamsBodyAgree = 0
	AbuseReportNewParamsBodyAgree1 AbuseReportNewParamsBodyAgree = 1
)

func (r AbuseReportNewParamsBodyAgree) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAgree0, AbuseReportNewParamsBodyAgree1:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyHostNotification string

const (
	AbuseReportNewParamsBodyHostNotificationSend     AbuseReportNewParamsBodyHostNotification = "send"
	AbuseReportNewParamsBodyHostNotificationSendAnon AbuseReportNewParamsBodyHostNotification = "send-anon"
	AbuseReportNewParamsBodyHostNotificationNone     AbuseReportNewParamsBodyHostNotification = "none"
)

func (r AbuseReportNewParamsBodyHostNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyHostNotificationSend, AbuseReportNewParamsBodyHostNotificationSendAnon, AbuseReportNewParamsBodyHostNotificationNone:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyNcmecNotification string

const (
	AbuseReportNewParamsBodyNcmecNotificationSend     AbuseReportNewParamsBodyNcmecNotification = "send"
	AbuseReportNewParamsBodyNcmecNotificationSendAnon AbuseReportNewParamsBodyNcmecNotification = "send-anon"
	AbuseReportNewParamsBodyNcmecNotificationNone     AbuseReportNewParamsBodyNcmecNotification = "none"
)

func (r AbuseReportNewParamsBodyNcmecNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyNcmecNotificationSend, AbuseReportNewParamsBodyNcmecNotificationSendAnon, AbuseReportNewParamsBodyNcmecNotificationNone:
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
