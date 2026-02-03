// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
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
)

// EmailSecurityTopTLDSpamService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailSecurityTopTLDSpamService] method instead.
type EmailSecurityTopTLDSpamService struct {
	Options []option.RequestOption
}

// NewEmailSecurityTopTLDSpamService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmailSecurityTopTLDSpamService(opts ...option.RequestOption) (r *EmailSecurityTopTLDSpamService) {
	r = &EmailSecurityTopTLDSpamService{}
	r.Options = opts
	return
}

// Retrieves the top TLDs by emails classified as spam or not.
func (r *EmailSecurityTopTLDSpamService) Get(ctx context.Context, spam EmailSecurityTopTLDSpamGetParamsSpam, query EmailSecurityTopTLDSpamGetParams, opts ...option.RequestOption) (res *EmailSecurityTopTLDSpamGetResponse, err error) {
	var env EmailSecurityTopTLDSpamGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/email/security/top/tlds/spam/%v", spam)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailSecurityTopTLDSpamGetResponse struct {
	// Metadata for the results.
	Meta EmailSecurityTopTLDSpamGetResponseMeta   `json:"meta,required"`
	Top0 []EmailSecurityTopTLDSpamGetResponseTop0 `json:"top_0,required"`
	JSON emailSecurityTopTLDSpamGetResponseJSON   `json:"-"`
}

// emailSecurityTopTLDSpamGetResponseJSON contains the JSON metadata for the struct
// [EmailSecurityTopTLDSpamGetResponse]
type emailSecurityTopTLDSpamGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTLDSpamGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDSpamGetResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailSecurityTopTLDSpamGetResponseMeta struct {
	ConfidenceInfo EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []EmailSecurityTopTLDSpamGetResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailSecurityTopTLDSpamGetResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailSecurityTopTLDSpamGetResponseMetaUnit `json:"units,required"`
	JSON  emailSecurityTopTLDSpamGetResponseMetaJSON   `json:"-"`
}

// emailSecurityTopTLDSpamGetResponseMetaJSON contains the JSON metadata for the
// struct [EmailSecurityTopTLDSpamGetResponseMeta]
type emailSecurityTopTLDSpamGetResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecurityTopTLDSpamGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDSpamGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfo struct {
	Annotations []EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                    `json:"level,required"`
	JSON  emailSecurityTopTLDSpamGetResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailSecurityTopTLDSpamGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfo]
type emailSecurityTopTLDSpamGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDSpamGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                    `json:"description,required"`
	EndDate     time.Time                                                                 `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                               `json:"isInstantaneous,required"`
	LinkedURL       string                                                             `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                          `json:"startDate,required" format:"date-time"`
	JSON            emailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotation]
type emailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceAll                EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceBGP                EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceBots               EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceCT                 EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceDNS                EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceDos                EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceFw                 EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceIQI                EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceNet                EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceAll, EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceBGP, EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceBots, EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceCT, EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceDNS, EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceDos, EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceFw, EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceIQI, EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceNet, EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsEventType string

const (
	EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsEventTypeEvent             EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsEventTypeOutage            EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsEventTypePipeline          EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsEventTypeEvent, EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsEventTypeOutage, EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsEventTypePipeline, EmailSecurityTopTLDSpamGetResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type EmailSecurityTopTLDSpamGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                           `json:"startTime,required" format:"date-time"`
	JSON      emailSecurityTopTLDSpamGetResponseMetaDateRangeJSON `json:"-"`
}

// emailSecurityTopTLDSpamGetResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [EmailSecurityTopTLDSpamGetResponseMetaDateRange]
type emailSecurityTopTLDSpamGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTLDSpamGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDSpamGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailSecurityTopTLDSpamGetResponseMetaNormalization string

const (
	EmailSecurityTopTLDSpamGetResponseMetaNormalizationPercentage           EmailSecurityTopTLDSpamGetResponseMetaNormalization = "PERCENTAGE"
	EmailSecurityTopTLDSpamGetResponseMetaNormalizationMin0Max              EmailSecurityTopTLDSpamGetResponseMetaNormalization = "MIN0_MAX"
	EmailSecurityTopTLDSpamGetResponseMetaNormalizationMinMax               EmailSecurityTopTLDSpamGetResponseMetaNormalization = "MIN_MAX"
	EmailSecurityTopTLDSpamGetResponseMetaNormalizationRawValues            EmailSecurityTopTLDSpamGetResponseMetaNormalization = "RAW_VALUES"
	EmailSecurityTopTLDSpamGetResponseMetaNormalizationPercentageChange     EmailSecurityTopTLDSpamGetResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailSecurityTopTLDSpamGetResponseMetaNormalizationRollingAverage       EmailSecurityTopTLDSpamGetResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailSecurityTopTLDSpamGetResponseMetaNormalizationOverlappedPercentage EmailSecurityTopTLDSpamGetResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailSecurityTopTLDSpamGetResponseMetaNormalizationRatio                EmailSecurityTopTLDSpamGetResponseMetaNormalization = "RATIO"
)

func (r EmailSecurityTopTLDSpamGetResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDSpamGetResponseMetaNormalizationPercentage, EmailSecurityTopTLDSpamGetResponseMetaNormalizationMin0Max, EmailSecurityTopTLDSpamGetResponseMetaNormalizationMinMax, EmailSecurityTopTLDSpamGetResponseMetaNormalizationRawValues, EmailSecurityTopTLDSpamGetResponseMetaNormalizationPercentageChange, EmailSecurityTopTLDSpamGetResponseMetaNormalizationRollingAverage, EmailSecurityTopTLDSpamGetResponseMetaNormalizationOverlappedPercentage, EmailSecurityTopTLDSpamGetResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailSecurityTopTLDSpamGetResponseMetaUnit struct {
	Name  string                                         `json:"name,required"`
	Value string                                         `json:"value,required"`
	JSON  emailSecurityTopTLDSpamGetResponseMetaUnitJSON `json:"-"`
}

// emailSecurityTopTLDSpamGetResponseMetaUnitJSON contains the JSON metadata for
// the struct [EmailSecurityTopTLDSpamGetResponseMetaUnit]
type emailSecurityTopTLDSpamGetResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTLDSpamGetResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDSpamGetResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTLDSpamGetResponseTop0 struct {
	Name string `json:"name,required"`
	// A numeric string.
	Value string                                     `json:"value,required"`
	JSON  emailSecurityTopTLDSpamGetResponseTop0JSON `json:"-"`
}

// emailSecurityTopTLDSpamGetResponseTop0JSON contains the JSON metadata for the
// struct [EmailSecurityTopTLDSpamGetResponseTop0]
type emailSecurityTopTLDSpamGetResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTLDSpamGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDSpamGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTLDSpamGetParams struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailSecurityTopTLDSpamGetParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailSecurityTopTLDSpamGetParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailSecurityTopTLDSpamGetParamsDMARC] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[EmailSecurityTopTLDSpamGetParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailSecurityTopTLDSpamGetParamsSPF] `query:"spf"`
	// Filters results by TLD category.
	TLDCategory param.Field[EmailSecurityTopTLDSpamGetParamsTLDCategory] `query:"tldCategory"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]EmailSecurityTopTLDSpamGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecurityTopTLDSpamGetParams]'s query parameters as
// `url.Values`.
func (r EmailSecurityTopTLDSpamGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Spam classification.
type EmailSecurityTopTLDSpamGetParamsSpam string

const (
	EmailSecurityTopTLDSpamGetParamsSpamSpam    EmailSecurityTopTLDSpamGetParamsSpam = "SPAM"
	EmailSecurityTopTLDSpamGetParamsSpamNotSpam EmailSecurityTopTLDSpamGetParamsSpam = "NOT_SPAM"
)

func (r EmailSecurityTopTLDSpamGetParamsSpam) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDSpamGetParamsSpamSpam, EmailSecurityTopTLDSpamGetParamsSpamNotSpam:
		return true
	}
	return false
}

type EmailSecurityTopTLDSpamGetParamsARC string

const (
	EmailSecurityTopTLDSpamGetParamsARCPass EmailSecurityTopTLDSpamGetParamsARC = "PASS"
	EmailSecurityTopTLDSpamGetParamsARCNone EmailSecurityTopTLDSpamGetParamsARC = "NONE"
	EmailSecurityTopTLDSpamGetParamsARCFail EmailSecurityTopTLDSpamGetParamsARC = "FAIL"
)

func (r EmailSecurityTopTLDSpamGetParamsARC) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDSpamGetParamsARCPass, EmailSecurityTopTLDSpamGetParamsARCNone, EmailSecurityTopTLDSpamGetParamsARCFail:
		return true
	}
	return false
}

type EmailSecurityTopTLDSpamGetParamsDKIM string

const (
	EmailSecurityTopTLDSpamGetParamsDKIMPass EmailSecurityTopTLDSpamGetParamsDKIM = "PASS"
	EmailSecurityTopTLDSpamGetParamsDKIMNone EmailSecurityTopTLDSpamGetParamsDKIM = "NONE"
	EmailSecurityTopTLDSpamGetParamsDKIMFail EmailSecurityTopTLDSpamGetParamsDKIM = "FAIL"
)

func (r EmailSecurityTopTLDSpamGetParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDSpamGetParamsDKIMPass, EmailSecurityTopTLDSpamGetParamsDKIMNone, EmailSecurityTopTLDSpamGetParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecurityTopTLDSpamGetParamsDMARC string

const (
	EmailSecurityTopTLDSpamGetParamsDMARCPass EmailSecurityTopTLDSpamGetParamsDMARC = "PASS"
	EmailSecurityTopTLDSpamGetParamsDMARCNone EmailSecurityTopTLDSpamGetParamsDMARC = "NONE"
	EmailSecurityTopTLDSpamGetParamsDMARCFail EmailSecurityTopTLDSpamGetParamsDMARC = "FAIL"
)

func (r EmailSecurityTopTLDSpamGetParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDSpamGetParamsDMARCPass, EmailSecurityTopTLDSpamGetParamsDMARCNone, EmailSecurityTopTLDSpamGetParamsDMARCFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailSecurityTopTLDSpamGetParamsFormat string

const (
	EmailSecurityTopTLDSpamGetParamsFormatJson EmailSecurityTopTLDSpamGetParamsFormat = "JSON"
	EmailSecurityTopTLDSpamGetParamsFormatCsv  EmailSecurityTopTLDSpamGetParamsFormat = "CSV"
)

func (r EmailSecurityTopTLDSpamGetParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDSpamGetParamsFormatJson, EmailSecurityTopTLDSpamGetParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecurityTopTLDSpamGetParamsSPF string

const (
	EmailSecurityTopTLDSpamGetParamsSPFPass EmailSecurityTopTLDSpamGetParamsSPF = "PASS"
	EmailSecurityTopTLDSpamGetParamsSPFNone EmailSecurityTopTLDSpamGetParamsSPF = "NONE"
	EmailSecurityTopTLDSpamGetParamsSPFFail EmailSecurityTopTLDSpamGetParamsSPF = "FAIL"
)

func (r EmailSecurityTopTLDSpamGetParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDSpamGetParamsSPFPass, EmailSecurityTopTLDSpamGetParamsSPFNone, EmailSecurityTopTLDSpamGetParamsSPFFail:
		return true
	}
	return false
}

// Filters results by TLD category.
type EmailSecurityTopTLDSpamGetParamsTLDCategory string

const (
	EmailSecurityTopTLDSpamGetParamsTLDCategoryClassic EmailSecurityTopTLDSpamGetParamsTLDCategory = "CLASSIC"
	EmailSecurityTopTLDSpamGetParamsTLDCategoryCountry EmailSecurityTopTLDSpamGetParamsTLDCategory = "COUNTRY"
)

func (r EmailSecurityTopTLDSpamGetParamsTLDCategory) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDSpamGetParamsTLDCategoryClassic, EmailSecurityTopTLDSpamGetParamsTLDCategoryCountry:
		return true
	}
	return false
}

type EmailSecurityTopTLDSpamGetParamsTLSVersion string

const (
	EmailSecurityTopTLDSpamGetParamsTLSVersionTlSv1_0 EmailSecurityTopTLDSpamGetParamsTLSVersion = "TLSv1_0"
	EmailSecurityTopTLDSpamGetParamsTLSVersionTlSv1_1 EmailSecurityTopTLDSpamGetParamsTLSVersion = "TLSv1_1"
	EmailSecurityTopTLDSpamGetParamsTLSVersionTlSv1_2 EmailSecurityTopTLDSpamGetParamsTLSVersion = "TLSv1_2"
	EmailSecurityTopTLDSpamGetParamsTLSVersionTlSv1_3 EmailSecurityTopTLDSpamGetParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecurityTopTLDSpamGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDSpamGetParamsTLSVersionTlSv1_0, EmailSecurityTopTLDSpamGetParamsTLSVersionTlSv1_1, EmailSecurityTopTLDSpamGetParamsTLSVersionTlSv1_2, EmailSecurityTopTLDSpamGetParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

type EmailSecurityTopTLDSpamGetResponseEnvelope struct {
	Result  EmailSecurityTopTLDSpamGetResponse             `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    emailSecurityTopTLDSpamGetResponseEnvelopeJSON `json:"-"`
}

// emailSecurityTopTLDSpamGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [EmailSecurityTopTLDSpamGetResponseEnvelope]
type emailSecurityTopTLDSpamGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTLDSpamGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDSpamGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
