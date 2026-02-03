// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
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

// EmailSecurityTopTLDService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailSecurityTopTLDService] method instead.
type EmailSecurityTopTLDService struct {
	Options   []option.RequestOption
	Malicious *EmailSecurityTopTLDMaliciousService
	Spam      *EmailSecurityTopTLDSpamService
	Spoof     *EmailSecurityTopTLDSpoofService
}

// NewEmailSecurityTopTLDService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmailSecurityTopTLDService(opts ...option.RequestOption) (r *EmailSecurityTopTLDService) {
	r = &EmailSecurityTopTLDService{}
	r.Options = opts
	r.Malicious = NewEmailSecurityTopTLDMaliciousService(opts...)
	r.Spam = NewEmailSecurityTopTLDSpamService(opts...)
	r.Spoof = NewEmailSecurityTopTLDSpoofService(opts...)
	return
}

// Retrieves the top TLDs by number of email messages.
func (r *EmailSecurityTopTLDService) Get(ctx context.Context, query EmailSecurityTopTLDGetParams, opts ...option.RequestOption) (res *EmailSecurityTopTLDGetResponse, err error) {
	var env EmailSecurityTopTLDGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/email/security/top/tlds"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailSecurityTopTLDGetResponse struct {
	// Metadata for the results.
	Meta EmailSecurityTopTLDGetResponseMeta   `json:"meta,required"`
	Top0 []EmailSecurityTopTLDGetResponseTop0 `json:"top_0,required"`
	JSON emailSecurityTopTLDGetResponseJSON   `json:"-"`
}

// emailSecurityTopTLDGetResponseJSON contains the JSON metadata for the struct
// [EmailSecurityTopTLDGetResponse]
type emailSecurityTopTLDGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTLDGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDGetResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailSecurityTopTLDGetResponseMeta struct {
	ConfidenceInfo EmailSecurityTopTLDGetResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []EmailSecurityTopTLDGetResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailSecurityTopTLDGetResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailSecurityTopTLDGetResponseMetaUnit `json:"units,required"`
	JSON  emailSecurityTopTLDGetResponseMetaJSON   `json:"-"`
}

// emailSecurityTopTLDGetResponseMetaJSON contains the JSON metadata for the struct
// [EmailSecurityTopTLDGetResponseMeta]
type emailSecurityTopTLDGetResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecurityTopTLDGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTLDGetResponseMetaConfidenceInfo struct {
	Annotations []EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                `json:"level,required"`
	JSON  emailSecurityTopTLDGetResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailSecurityTopTLDGetResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [EmailSecurityTopTLDGetResponseMetaConfidenceInfo]
type emailSecurityTopTLDGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTLDGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                `json:"description,required"`
	EndDate     time.Time                                                             `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                           `json:"isInstantaneous,required"`
	LinkedURL       string                                                         `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                      `json:"startDate,required" format:"date-time"`
	JSON            emailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotation]
type emailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceAll                EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceBGP                EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceBots               EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceCT                 EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceDNS                EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceDos                EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceFw                 EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceIQI                EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceNet                EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceAll, EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceBGP, EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceBots, EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceCT, EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceDNS, EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceDos, EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceFw, EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceIQI, EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceNet, EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsEventType string

const (
	EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsEventTypeEvent             EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsEventTypeOutage            EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsEventTypePipeline          EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsEventTypeEvent, EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsEventTypeOutage, EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsEventTypePipeline, EmailSecurityTopTLDGetResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type EmailSecurityTopTLDGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      emailSecurityTopTLDGetResponseMetaDateRangeJSON `json:"-"`
}

// emailSecurityTopTLDGetResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [EmailSecurityTopTLDGetResponseMetaDateRange]
type emailSecurityTopTLDGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTLDGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailSecurityTopTLDGetResponseMetaNormalization string

const (
	EmailSecurityTopTLDGetResponseMetaNormalizationPercentage           EmailSecurityTopTLDGetResponseMetaNormalization = "PERCENTAGE"
	EmailSecurityTopTLDGetResponseMetaNormalizationMin0Max              EmailSecurityTopTLDGetResponseMetaNormalization = "MIN0_MAX"
	EmailSecurityTopTLDGetResponseMetaNormalizationMinMax               EmailSecurityTopTLDGetResponseMetaNormalization = "MIN_MAX"
	EmailSecurityTopTLDGetResponseMetaNormalizationRawValues            EmailSecurityTopTLDGetResponseMetaNormalization = "RAW_VALUES"
	EmailSecurityTopTLDGetResponseMetaNormalizationPercentageChange     EmailSecurityTopTLDGetResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailSecurityTopTLDGetResponseMetaNormalizationRollingAverage       EmailSecurityTopTLDGetResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailSecurityTopTLDGetResponseMetaNormalizationOverlappedPercentage EmailSecurityTopTLDGetResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailSecurityTopTLDGetResponseMetaNormalizationRatio                EmailSecurityTopTLDGetResponseMetaNormalization = "RATIO"
)

func (r EmailSecurityTopTLDGetResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDGetResponseMetaNormalizationPercentage, EmailSecurityTopTLDGetResponseMetaNormalizationMin0Max, EmailSecurityTopTLDGetResponseMetaNormalizationMinMax, EmailSecurityTopTLDGetResponseMetaNormalizationRawValues, EmailSecurityTopTLDGetResponseMetaNormalizationPercentageChange, EmailSecurityTopTLDGetResponseMetaNormalizationRollingAverage, EmailSecurityTopTLDGetResponseMetaNormalizationOverlappedPercentage, EmailSecurityTopTLDGetResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailSecurityTopTLDGetResponseMetaUnit struct {
	Name  string                                     `json:"name,required"`
	Value string                                     `json:"value,required"`
	JSON  emailSecurityTopTLDGetResponseMetaUnitJSON `json:"-"`
}

// emailSecurityTopTLDGetResponseMetaUnitJSON contains the JSON metadata for the
// struct [EmailSecurityTopTLDGetResponseMetaUnit]
type emailSecurityTopTLDGetResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTLDGetResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDGetResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTLDGetResponseTop0 struct {
	Name string `json:"name,required"`
	// A numeric string.
	Value string                                 `json:"value,required"`
	JSON  emailSecurityTopTLDGetResponseTop0JSON `json:"-"`
}

// emailSecurityTopTLDGetResponseTop0JSON contains the JSON metadata for the struct
// [EmailSecurityTopTLDGetResponseTop0]
type emailSecurityTopTLDGetResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTLDGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTLDGetParams struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailSecurityTopTLDGetParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailSecurityTopTLDGetParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailSecurityTopTLDGetParamsDMARC] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[EmailSecurityTopTLDGetParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailSecurityTopTLDGetParamsSPF] `query:"spf"`
	// Filters results by TLD category.
	TLDCategory param.Field[EmailSecurityTopTLDGetParamsTLDCategory] `query:"tldCategory"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]EmailSecurityTopTLDGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecurityTopTLDGetParams]'s query parameters as
// `url.Values`.
func (r EmailSecurityTopTLDGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EmailSecurityTopTLDGetParamsARC string

const (
	EmailSecurityTopTLDGetParamsARCPass EmailSecurityTopTLDGetParamsARC = "PASS"
	EmailSecurityTopTLDGetParamsARCNone EmailSecurityTopTLDGetParamsARC = "NONE"
	EmailSecurityTopTLDGetParamsARCFail EmailSecurityTopTLDGetParamsARC = "FAIL"
)

func (r EmailSecurityTopTLDGetParamsARC) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDGetParamsARCPass, EmailSecurityTopTLDGetParamsARCNone, EmailSecurityTopTLDGetParamsARCFail:
		return true
	}
	return false
}

type EmailSecurityTopTLDGetParamsDKIM string

const (
	EmailSecurityTopTLDGetParamsDKIMPass EmailSecurityTopTLDGetParamsDKIM = "PASS"
	EmailSecurityTopTLDGetParamsDKIMNone EmailSecurityTopTLDGetParamsDKIM = "NONE"
	EmailSecurityTopTLDGetParamsDKIMFail EmailSecurityTopTLDGetParamsDKIM = "FAIL"
)

func (r EmailSecurityTopTLDGetParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDGetParamsDKIMPass, EmailSecurityTopTLDGetParamsDKIMNone, EmailSecurityTopTLDGetParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecurityTopTLDGetParamsDMARC string

const (
	EmailSecurityTopTLDGetParamsDMARCPass EmailSecurityTopTLDGetParamsDMARC = "PASS"
	EmailSecurityTopTLDGetParamsDMARCNone EmailSecurityTopTLDGetParamsDMARC = "NONE"
	EmailSecurityTopTLDGetParamsDMARCFail EmailSecurityTopTLDGetParamsDMARC = "FAIL"
)

func (r EmailSecurityTopTLDGetParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDGetParamsDMARCPass, EmailSecurityTopTLDGetParamsDMARCNone, EmailSecurityTopTLDGetParamsDMARCFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailSecurityTopTLDGetParamsFormat string

const (
	EmailSecurityTopTLDGetParamsFormatJson EmailSecurityTopTLDGetParamsFormat = "JSON"
	EmailSecurityTopTLDGetParamsFormatCsv  EmailSecurityTopTLDGetParamsFormat = "CSV"
)

func (r EmailSecurityTopTLDGetParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDGetParamsFormatJson, EmailSecurityTopTLDGetParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecurityTopTLDGetParamsSPF string

const (
	EmailSecurityTopTLDGetParamsSPFPass EmailSecurityTopTLDGetParamsSPF = "PASS"
	EmailSecurityTopTLDGetParamsSPFNone EmailSecurityTopTLDGetParamsSPF = "NONE"
	EmailSecurityTopTLDGetParamsSPFFail EmailSecurityTopTLDGetParamsSPF = "FAIL"
)

func (r EmailSecurityTopTLDGetParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDGetParamsSPFPass, EmailSecurityTopTLDGetParamsSPFNone, EmailSecurityTopTLDGetParamsSPFFail:
		return true
	}
	return false
}

// Filters results by TLD category.
type EmailSecurityTopTLDGetParamsTLDCategory string

const (
	EmailSecurityTopTLDGetParamsTLDCategoryClassic EmailSecurityTopTLDGetParamsTLDCategory = "CLASSIC"
	EmailSecurityTopTLDGetParamsTLDCategoryCountry EmailSecurityTopTLDGetParamsTLDCategory = "COUNTRY"
)

func (r EmailSecurityTopTLDGetParamsTLDCategory) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDGetParamsTLDCategoryClassic, EmailSecurityTopTLDGetParamsTLDCategoryCountry:
		return true
	}
	return false
}

type EmailSecurityTopTLDGetParamsTLSVersion string

const (
	EmailSecurityTopTLDGetParamsTLSVersionTlSv1_0 EmailSecurityTopTLDGetParamsTLSVersion = "TLSv1_0"
	EmailSecurityTopTLDGetParamsTLSVersionTlSv1_1 EmailSecurityTopTLDGetParamsTLSVersion = "TLSv1_1"
	EmailSecurityTopTLDGetParamsTLSVersionTlSv1_2 EmailSecurityTopTLDGetParamsTLSVersion = "TLSv1_2"
	EmailSecurityTopTLDGetParamsTLSVersionTlSv1_3 EmailSecurityTopTLDGetParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecurityTopTLDGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDGetParamsTLSVersionTlSv1_0, EmailSecurityTopTLDGetParamsTLSVersionTlSv1_1, EmailSecurityTopTLDGetParamsTLSVersionTlSv1_2, EmailSecurityTopTLDGetParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

type EmailSecurityTopTLDGetResponseEnvelope struct {
	Result  EmailSecurityTopTLDGetResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    emailSecurityTopTLDGetResponseEnvelopeJSON `json:"-"`
}

// emailSecurityTopTLDGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailSecurityTopTLDGetResponseEnvelope]
type emailSecurityTopTLDGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTLDGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
