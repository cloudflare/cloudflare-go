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

// EmailSecurityTopTLDMaliciousService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailSecurityTopTLDMaliciousService] method instead.
type EmailSecurityTopTLDMaliciousService struct {
	Options []option.RequestOption
}

// NewEmailSecurityTopTLDMaliciousService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEmailSecurityTopTLDMaliciousService(opts ...option.RequestOption) (r *EmailSecurityTopTLDMaliciousService) {
	r = &EmailSecurityTopTLDMaliciousService{}
	r.Options = opts
	return
}

// Retrieves the top TLDs by emails classified as malicious or not.
func (r *EmailSecurityTopTLDMaliciousService) Get(ctx context.Context, malicious EmailSecurityTopTLDMaliciousGetParamsMalicious, query EmailSecurityTopTLDMaliciousGetParams, opts ...option.RequestOption) (res *EmailSecurityTopTLDMaliciousGetResponse, err error) {
	var env EmailSecurityTopTLDMaliciousGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/email/security/top/tlds/malicious/%v", malicious)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailSecurityTopTLDMaliciousGetResponse struct {
	// Metadata for the results.
	Meta EmailSecurityTopTLDMaliciousGetResponseMeta   `json:"meta,required"`
	Top0 []EmailSecurityTopTLDMaliciousGetResponseTop0 `json:"top_0,required"`
	JSON emailSecurityTopTLDMaliciousGetResponseJSON   `json:"-"`
}

// emailSecurityTopTLDMaliciousGetResponseJSON contains the JSON metadata for the
// struct [EmailSecurityTopTLDMaliciousGetResponse]
type emailSecurityTopTLDMaliciousGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTLDMaliciousGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDMaliciousGetResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailSecurityTopTLDMaliciousGetResponseMeta struct {
	ConfidenceInfo EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []EmailSecurityTopTLDMaliciousGetResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailSecurityTopTLDMaliciousGetResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailSecurityTopTLDMaliciousGetResponseMetaUnit `json:"units,required"`
	JSON  emailSecurityTopTLDMaliciousGetResponseMetaJSON   `json:"-"`
}

// emailSecurityTopTLDMaliciousGetResponseMetaJSON contains the JSON metadata for
// the struct [EmailSecurityTopTLDMaliciousGetResponseMeta]
type emailSecurityTopTLDMaliciousGetResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecurityTopTLDMaliciousGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDMaliciousGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfo struct {
	Annotations []EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                         `json:"level,required"`
	JSON  emailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfo]
type emailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                         `json:"description,required"`
	EndDate     time.Time                                                                      `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                    `json:"isInstantaneous,required"`
	LinkedURL       string                                                                  `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                               `json:"startDate,required" format:"date-time"`
	JSON            emailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotation]
type emailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceAll                EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceBGP                EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceBots               EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceCT                 EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceDNS                EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceDos                EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceFw                 EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceIQI                EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceNet                EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceAll, EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceBGP, EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceBots, EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceCT, EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceDNS, EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceDos, EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceFw, EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceIQI, EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceNet, EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsEventType string

const (
	EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsEventTypeEvent             EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsEventTypeOutage            EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsEventTypePipeline          EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsEventTypeEvent, EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsEventTypeOutage, EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsEventTypePipeline, EmailSecurityTopTLDMaliciousGetResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type EmailSecurityTopTLDMaliciousGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                `json:"startTime,required" format:"date-time"`
	JSON      emailSecurityTopTLDMaliciousGetResponseMetaDateRangeJSON `json:"-"`
}

// emailSecurityTopTLDMaliciousGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [EmailSecurityTopTLDMaliciousGetResponseMetaDateRange]
type emailSecurityTopTLDMaliciousGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTLDMaliciousGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDMaliciousGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailSecurityTopTLDMaliciousGetResponseMetaNormalization string

const (
	EmailSecurityTopTLDMaliciousGetResponseMetaNormalizationPercentage           EmailSecurityTopTLDMaliciousGetResponseMetaNormalization = "PERCENTAGE"
	EmailSecurityTopTLDMaliciousGetResponseMetaNormalizationMin0Max              EmailSecurityTopTLDMaliciousGetResponseMetaNormalization = "MIN0_MAX"
	EmailSecurityTopTLDMaliciousGetResponseMetaNormalizationMinMax               EmailSecurityTopTLDMaliciousGetResponseMetaNormalization = "MIN_MAX"
	EmailSecurityTopTLDMaliciousGetResponseMetaNormalizationRawValues            EmailSecurityTopTLDMaliciousGetResponseMetaNormalization = "RAW_VALUES"
	EmailSecurityTopTLDMaliciousGetResponseMetaNormalizationPercentageChange     EmailSecurityTopTLDMaliciousGetResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailSecurityTopTLDMaliciousGetResponseMetaNormalizationRollingAverage       EmailSecurityTopTLDMaliciousGetResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailSecurityTopTLDMaliciousGetResponseMetaNormalizationOverlappedPercentage EmailSecurityTopTLDMaliciousGetResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailSecurityTopTLDMaliciousGetResponseMetaNormalizationRatio                EmailSecurityTopTLDMaliciousGetResponseMetaNormalization = "RATIO"
)

func (r EmailSecurityTopTLDMaliciousGetResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDMaliciousGetResponseMetaNormalizationPercentage, EmailSecurityTopTLDMaliciousGetResponseMetaNormalizationMin0Max, EmailSecurityTopTLDMaliciousGetResponseMetaNormalizationMinMax, EmailSecurityTopTLDMaliciousGetResponseMetaNormalizationRawValues, EmailSecurityTopTLDMaliciousGetResponseMetaNormalizationPercentageChange, EmailSecurityTopTLDMaliciousGetResponseMetaNormalizationRollingAverage, EmailSecurityTopTLDMaliciousGetResponseMetaNormalizationOverlappedPercentage, EmailSecurityTopTLDMaliciousGetResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailSecurityTopTLDMaliciousGetResponseMetaUnit struct {
	Name  string                                              `json:"name,required"`
	Value string                                              `json:"value,required"`
	JSON  emailSecurityTopTLDMaliciousGetResponseMetaUnitJSON `json:"-"`
}

// emailSecurityTopTLDMaliciousGetResponseMetaUnitJSON contains the JSON metadata
// for the struct [EmailSecurityTopTLDMaliciousGetResponseMetaUnit]
type emailSecurityTopTLDMaliciousGetResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTLDMaliciousGetResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDMaliciousGetResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTLDMaliciousGetResponseTop0 struct {
	Name string `json:"name,required"`
	// A numeric string.
	Value string                                          `json:"value,required"`
	JSON  emailSecurityTopTLDMaliciousGetResponseTop0JSON `json:"-"`
}

// emailSecurityTopTLDMaliciousGetResponseTop0JSON contains the JSON metadata for
// the struct [EmailSecurityTopTLDMaliciousGetResponseTop0]
type emailSecurityTopTLDMaliciousGetResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTLDMaliciousGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDMaliciousGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTLDMaliciousGetParams struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailSecurityTopTLDMaliciousGetParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailSecurityTopTLDMaliciousGetParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailSecurityTopTLDMaliciousGetParamsDMARC] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[EmailSecurityTopTLDMaliciousGetParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailSecurityTopTLDMaliciousGetParamsSPF] `query:"spf"`
	// Filters results by TLD category.
	TLDCategory param.Field[EmailSecurityTopTLDMaliciousGetParamsTLDCategory] `query:"tldCategory"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]EmailSecurityTopTLDMaliciousGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecurityTopTLDMaliciousGetParams]'s query parameters
// as `url.Values`.
func (r EmailSecurityTopTLDMaliciousGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Malicious classification.
type EmailSecurityTopTLDMaliciousGetParamsMalicious string

const (
	EmailSecurityTopTLDMaliciousGetParamsMaliciousMalicious    EmailSecurityTopTLDMaliciousGetParamsMalicious = "MALICIOUS"
	EmailSecurityTopTLDMaliciousGetParamsMaliciousNotMalicious EmailSecurityTopTLDMaliciousGetParamsMalicious = "NOT_MALICIOUS"
)

func (r EmailSecurityTopTLDMaliciousGetParamsMalicious) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDMaliciousGetParamsMaliciousMalicious, EmailSecurityTopTLDMaliciousGetParamsMaliciousNotMalicious:
		return true
	}
	return false
}

type EmailSecurityTopTLDMaliciousGetParamsARC string

const (
	EmailSecurityTopTLDMaliciousGetParamsARCPass EmailSecurityTopTLDMaliciousGetParamsARC = "PASS"
	EmailSecurityTopTLDMaliciousGetParamsARCNone EmailSecurityTopTLDMaliciousGetParamsARC = "NONE"
	EmailSecurityTopTLDMaliciousGetParamsARCFail EmailSecurityTopTLDMaliciousGetParamsARC = "FAIL"
)

func (r EmailSecurityTopTLDMaliciousGetParamsARC) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDMaliciousGetParamsARCPass, EmailSecurityTopTLDMaliciousGetParamsARCNone, EmailSecurityTopTLDMaliciousGetParamsARCFail:
		return true
	}
	return false
}

type EmailSecurityTopTLDMaliciousGetParamsDKIM string

const (
	EmailSecurityTopTLDMaliciousGetParamsDKIMPass EmailSecurityTopTLDMaliciousGetParamsDKIM = "PASS"
	EmailSecurityTopTLDMaliciousGetParamsDKIMNone EmailSecurityTopTLDMaliciousGetParamsDKIM = "NONE"
	EmailSecurityTopTLDMaliciousGetParamsDKIMFail EmailSecurityTopTLDMaliciousGetParamsDKIM = "FAIL"
)

func (r EmailSecurityTopTLDMaliciousGetParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDMaliciousGetParamsDKIMPass, EmailSecurityTopTLDMaliciousGetParamsDKIMNone, EmailSecurityTopTLDMaliciousGetParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecurityTopTLDMaliciousGetParamsDMARC string

const (
	EmailSecurityTopTLDMaliciousGetParamsDMARCPass EmailSecurityTopTLDMaliciousGetParamsDMARC = "PASS"
	EmailSecurityTopTLDMaliciousGetParamsDMARCNone EmailSecurityTopTLDMaliciousGetParamsDMARC = "NONE"
	EmailSecurityTopTLDMaliciousGetParamsDMARCFail EmailSecurityTopTLDMaliciousGetParamsDMARC = "FAIL"
)

func (r EmailSecurityTopTLDMaliciousGetParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDMaliciousGetParamsDMARCPass, EmailSecurityTopTLDMaliciousGetParamsDMARCNone, EmailSecurityTopTLDMaliciousGetParamsDMARCFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailSecurityTopTLDMaliciousGetParamsFormat string

const (
	EmailSecurityTopTLDMaliciousGetParamsFormatJson EmailSecurityTopTLDMaliciousGetParamsFormat = "JSON"
	EmailSecurityTopTLDMaliciousGetParamsFormatCsv  EmailSecurityTopTLDMaliciousGetParamsFormat = "CSV"
)

func (r EmailSecurityTopTLDMaliciousGetParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDMaliciousGetParamsFormatJson, EmailSecurityTopTLDMaliciousGetParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecurityTopTLDMaliciousGetParamsSPF string

const (
	EmailSecurityTopTLDMaliciousGetParamsSPFPass EmailSecurityTopTLDMaliciousGetParamsSPF = "PASS"
	EmailSecurityTopTLDMaliciousGetParamsSPFNone EmailSecurityTopTLDMaliciousGetParamsSPF = "NONE"
	EmailSecurityTopTLDMaliciousGetParamsSPFFail EmailSecurityTopTLDMaliciousGetParamsSPF = "FAIL"
)

func (r EmailSecurityTopTLDMaliciousGetParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDMaliciousGetParamsSPFPass, EmailSecurityTopTLDMaliciousGetParamsSPFNone, EmailSecurityTopTLDMaliciousGetParamsSPFFail:
		return true
	}
	return false
}

// Filters results by TLD category.
type EmailSecurityTopTLDMaliciousGetParamsTLDCategory string

const (
	EmailSecurityTopTLDMaliciousGetParamsTLDCategoryClassic EmailSecurityTopTLDMaliciousGetParamsTLDCategory = "CLASSIC"
	EmailSecurityTopTLDMaliciousGetParamsTLDCategoryCountry EmailSecurityTopTLDMaliciousGetParamsTLDCategory = "COUNTRY"
)

func (r EmailSecurityTopTLDMaliciousGetParamsTLDCategory) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDMaliciousGetParamsTLDCategoryClassic, EmailSecurityTopTLDMaliciousGetParamsTLDCategoryCountry:
		return true
	}
	return false
}

type EmailSecurityTopTLDMaliciousGetParamsTLSVersion string

const (
	EmailSecurityTopTLDMaliciousGetParamsTLSVersionTlSv1_0 EmailSecurityTopTLDMaliciousGetParamsTLSVersion = "TLSv1_0"
	EmailSecurityTopTLDMaliciousGetParamsTLSVersionTlSv1_1 EmailSecurityTopTLDMaliciousGetParamsTLSVersion = "TLSv1_1"
	EmailSecurityTopTLDMaliciousGetParamsTLSVersionTlSv1_2 EmailSecurityTopTLDMaliciousGetParamsTLSVersion = "TLSv1_2"
	EmailSecurityTopTLDMaliciousGetParamsTLSVersionTlSv1_3 EmailSecurityTopTLDMaliciousGetParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecurityTopTLDMaliciousGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDMaliciousGetParamsTLSVersionTlSv1_0, EmailSecurityTopTLDMaliciousGetParamsTLSVersionTlSv1_1, EmailSecurityTopTLDMaliciousGetParamsTLSVersionTlSv1_2, EmailSecurityTopTLDMaliciousGetParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

type EmailSecurityTopTLDMaliciousGetResponseEnvelope struct {
	Result  EmailSecurityTopTLDMaliciousGetResponse             `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    emailSecurityTopTLDMaliciousGetResponseEnvelopeJSON `json:"-"`
}

// emailSecurityTopTLDMaliciousGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailSecurityTopTLDMaliciousGetResponseEnvelope]
type emailSecurityTopTLDMaliciousGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTLDMaliciousGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDMaliciousGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
