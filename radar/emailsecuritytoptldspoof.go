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

// EmailSecurityTopTLDSpoofService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailSecurityTopTLDSpoofService] method instead.
type EmailSecurityTopTLDSpoofService struct {
	Options []option.RequestOption
}

// NewEmailSecurityTopTLDSpoofService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEmailSecurityTopTLDSpoofService(opts ...option.RequestOption) (r *EmailSecurityTopTLDSpoofService) {
	r = &EmailSecurityTopTLDSpoofService{}
	r.Options = opts
	return
}

// Retrieves the top TLDs by emails classified as spoof or not.
func (r *EmailSecurityTopTLDSpoofService) Get(ctx context.Context, spoof EmailSecurityTopTLDSpoofGetParamsSpoof, query EmailSecurityTopTLDSpoofGetParams, opts ...option.RequestOption) (res *EmailSecurityTopTLDSpoofGetResponse, err error) {
	var env EmailSecurityTopTLDSpoofGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/email/security/top/tlds/spoof/%v", spoof)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailSecurityTopTLDSpoofGetResponse struct {
	// Metadata for the results.
	Meta EmailSecurityTopTLDSpoofGetResponseMeta   `json:"meta,required"`
	Top0 []EmailSecurityTopTLDSpoofGetResponseTop0 `json:"top_0,required"`
	JSON emailSecurityTopTLDSpoofGetResponseJSON   `json:"-"`
}

// emailSecurityTopTLDSpoofGetResponseJSON contains the JSON metadata for the
// struct [EmailSecurityTopTLDSpoofGetResponse]
type emailSecurityTopTLDSpoofGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTLDSpoofGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDSpoofGetResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailSecurityTopTLDSpoofGetResponseMeta struct {
	ConfidenceInfo EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []EmailSecurityTopTLDSpoofGetResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailSecurityTopTLDSpoofGetResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailSecurityTopTLDSpoofGetResponseMetaUnit `json:"units,required"`
	JSON  emailSecurityTopTLDSpoofGetResponseMetaJSON   `json:"-"`
}

// emailSecurityTopTLDSpoofGetResponseMetaJSON contains the JSON metadata for the
// struct [EmailSecurityTopTLDSpoofGetResponseMeta]
type emailSecurityTopTLDSpoofGetResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecurityTopTLDSpoofGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDSpoofGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfo struct {
	Annotations []EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                     `json:"level,required"`
	JSON  emailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfo]
type emailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                     `json:"description,required"`
	EndDate     time.Time                                                                  `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                `json:"isInstantaneous,required"`
	LinkedURL       string                                                              `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                           `json:"startDate,required" format:"date-time"`
	JSON            emailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotation]
type emailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceAll                EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceBGP                EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceBots               EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceCT                 EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceDNS                EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceDos                EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceFw                 EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceIQI                EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceNet                EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceAll, EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceBGP, EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceBots, EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceCT, EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceDNS, EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceDos, EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceFw, EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceIQI, EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceNet, EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsEventType string

const (
	EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsEventTypeEvent             EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsEventTypeOutage            EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsEventTypePipeline          EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsEventTypeEvent, EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsEventTypeOutage, EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsEventTypePipeline, EmailSecurityTopTLDSpoofGetResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type EmailSecurityTopTLDSpoofGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                            `json:"startTime,required" format:"date-time"`
	JSON      emailSecurityTopTLDSpoofGetResponseMetaDateRangeJSON `json:"-"`
}

// emailSecurityTopTLDSpoofGetResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [EmailSecurityTopTLDSpoofGetResponseMetaDateRange]
type emailSecurityTopTLDSpoofGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTLDSpoofGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDSpoofGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailSecurityTopTLDSpoofGetResponseMetaNormalization string

const (
	EmailSecurityTopTLDSpoofGetResponseMetaNormalizationPercentage           EmailSecurityTopTLDSpoofGetResponseMetaNormalization = "PERCENTAGE"
	EmailSecurityTopTLDSpoofGetResponseMetaNormalizationMin0Max              EmailSecurityTopTLDSpoofGetResponseMetaNormalization = "MIN0_MAX"
	EmailSecurityTopTLDSpoofGetResponseMetaNormalizationMinMax               EmailSecurityTopTLDSpoofGetResponseMetaNormalization = "MIN_MAX"
	EmailSecurityTopTLDSpoofGetResponseMetaNormalizationRawValues            EmailSecurityTopTLDSpoofGetResponseMetaNormalization = "RAW_VALUES"
	EmailSecurityTopTLDSpoofGetResponseMetaNormalizationPercentageChange     EmailSecurityTopTLDSpoofGetResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailSecurityTopTLDSpoofGetResponseMetaNormalizationRollingAverage       EmailSecurityTopTLDSpoofGetResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailSecurityTopTLDSpoofGetResponseMetaNormalizationOverlappedPercentage EmailSecurityTopTLDSpoofGetResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailSecurityTopTLDSpoofGetResponseMetaNormalizationRatio                EmailSecurityTopTLDSpoofGetResponseMetaNormalization = "RATIO"
)

func (r EmailSecurityTopTLDSpoofGetResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDSpoofGetResponseMetaNormalizationPercentage, EmailSecurityTopTLDSpoofGetResponseMetaNormalizationMin0Max, EmailSecurityTopTLDSpoofGetResponseMetaNormalizationMinMax, EmailSecurityTopTLDSpoofGetResponseMetaNormalizationRawValues, EmailSecurityTopTLDSpoofGetResponseMetaNormalizationPercentageChange, EmailSecurityTopTLDSpoofGetResponseMetaNormalizationRollingAverage, EmailSecurityTopTLDSpoofGetResponseMetaNormalizationOverlappedPercentage, EmailSecurityTopTLDSpoofGetResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailSecurityTopTLDSpoofGetResponseMetaUnit struct {
	Name  string                                          `json:"name,required"`
	Value string                                          `json:"value,required"`
	JSON  emailSecurityTopTLDSpoofGetResponseMetaUnitJSON `json:"-"`
}

// emailSecurityTopTLDSpoofGetResponseMetaUnitJSON contains the JSON metadata for
// the struct [EmailSecurityTopTLDSpoofGetResponseMetaUnit]
type emailSecurityTopTLDSpoofGetResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTLDSpoofGetResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDSpoofGetResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTLDSpoofGetResponseTop0 struct {
	Name string `json:"name,required"`
	// A numeric string.
	Value string                                      `json:"value,required"`
	JSON  emailSecurityTopTLDSpoofGetResponseTop0JSON `json:"-"`
}

// emailSecurityTopTLDSpoofGetResponseTop0JSON contains the JSON metadata for the
// struct [EmailSecurityTopTLDSpoofGetResponseTop0]
type emailSecurityTopTLDSpoofGetResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTLDSpoofGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDSpoofGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTLDSpoofGetParams struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailSecurityTopTLDSpoofGetParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailSecurityTopTLDSpoofGetParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailSecurityTopTLDSpoofGetParamsDMARC] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[EmailSecurityTopTLDSpoofGetParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailSecurityTopTLDSpoofGetParamsSPF] `query:"spf"`
	// Filters results by TLD category.
	TLDCategory param.Field[EmailSecurityTopTLDSpoofGetParamsTLDCategory] `query:"tldCategory"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]EmailSecurityTopTLDSpoofGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecurityTopTLDSpoofGetParams]'s query parameters as
// `url.Values`.
func (r EmailSecurityTopTLDSpoofGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Spoof classification.
type EmailSecurityTopTLDSpoofGetParamsSpoof string

const (
	EmailSecurityTopTLDSpoofGetParamsSpoofSpoof    EmailSecurityTopTLDSpoofGetParamsSpoof = "SPOOF"
	EmailSecurityTopTLDSpoofGetParamsSpoofNotSpoof EmailSecurityTopTLDSpoofGetParamsSpoof = "NOT_SPOOF"
)

func (r EmailSecurityTopTLDSpoofGetParamsSpoof) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDSpoofGetParamsSpoofSpoof, EmailSecurityTopTLDSpoofGetParamsSpoofNotSpoof:
		return true
	}
	return false
}

type EmailSecurityTopTLDSpoofGetParamsARC string

const (
	EmailSecurityTopTLDSpoofGetParamsARCPass EmailSecurityTopTLDSpoofGetParamsARC = "PASS"
	EmailSecurityTopTLDSpoofGetParamsARCNone EmailSecurityTopTLDSpoofGetParamsARC = "NONE"
	EmailSecurityTopTLDSpoofGetParamsARCFail EmailSecurityTopTLDSpoofGetParamsARC = "FAIL"
)

func (r EmailSecurityTopTLDSpoofGetParamsARC) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDSpoofGetParamsARCPass, EmailSecurityTopTLDSpoofGetParamsARCNone, EmailSecurityTopTLDSpoofGetParamsARCFail:
		return true
	}
	return false
}

type EmailSecurityTopTLDSpoofGetParamsDKIM string

const (
	EmailSecurityTopTLDSpoofGetParamsDKIMPass EmailSecurityTopTLDSpoofGetParamsDKIM = "PASS"
	EmailSecurityTopTLDSpoofGetParamsDKIMNone EmailSecurityTopTLDSpoofGetParamsDKIM = "NONE"
	EmailSecurityTopTLDSpoofGetParamsDKIMFail EmailSecurityTopTLDSpoofGetParamsDKIM = "FAIL"
)

func (r EmailSecurityTopTLDSpoofGetParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDSpoofGetParamsDKIMPass, EmailSecurityTopTLDSpoofGetParamsDKIMNone, EmailSecurityTopTLDSpoofGetParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecurityTopTLDSpoofGetParamsDMARC string

const (
	EmailSecurityTopTLDSpoofGetParamsDMARCPass EmailSecurityTopTLDSpoofGetParamsDMARC = "PASS"
	EmailSecurityTopTLDSpoofGetParamsDMARCNone EmailSecurityTopTLDSpoofGetParamsDMARC = "NONE"
	EmailSecurityTopTLDSpoofGetParamsDMARCFail EmailSecurityTopTLDSpoofGetParamsDMARC = "FAIL"
)

func (r EmailSecurityTopTLDSpoofGetParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDSpoofGetParamsDMARCPass, EmailSecurityTopTLDSpoofGetParamsDMARCNone, EmailSecurityTopTLDSpoofGetParamsDMARCFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailSecurityTopTLDSpoofGetParamsFormat string

const (
	EmailSecurityTopTLDSpoofGetParamsFormatJson EmailSecurityTopTLDSpoofGetParamsFormat = "JSON"
	EmailSecurityTopTLDSpoofGetParamsFormatCsv  EmailSecurityTopTLDSpoofGetParamsFormat = "CSV"
)

func (r EmailSecurityTopTLDSpoofGetParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDSpoofGetParamsFormatJson, EmailSecurityTopTLDSpoofGetParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecurityTopTLDSpoofGetParamsSPF string

const (
	EmailSecurityTopTLDSpoofGetParamsSPFPass EmailSecurityTopTLDSpoofGetParamsSPF = "PASS"
	EmailSecurityTopTLDSpoofGetParamsSPFNone EmailSecurityTopTLDSpoofGetParamsSPF = "NONE"
	EmailSecurityTopTLDSpoofGetParamsSPFFail EmailSecurityTopTLDSpoofGetParamsSPF = "FAIL"
)

func (r EmailSecurityTopTLDSpoofGetParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDSpoofGetParamsSPFPass, EmailSecurityTopTLDSpoofGetParamsSPFNone, EmailSecurityTopTLDSpoofGetParamsSPFFail:
		return true
	}
	return false
}

// Filters results by TLD category.
type EmailSecurityTopTLDSpoofGetParamsTLDCategory string

const (
	EmailSecurityTopTLDSpoofGetParamsTLDCategoryClassic EmailSecurityTopTLDSpoofGetParamsTLDCategory = "CLASSIC"
	EmailSecurityTopTLDSpoofGetParamsTLDCategoryCountry EmailSecurityTopTLDSpoofGetParamsTLDCategory = "COUNTRY"
)

func (r EmailSecurityTopTLDSpoofGetParamsTLDCategory) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDSpoofGetParamsTLDCategoryClassic, EmailSecurityTopTLDSpoofGetParamsTLDCategoryCountry:
		return true
	}
	return false
}

type EmailSecurityTopTLDSpoofGetParamsTLSVersion string

const (
	EmailSecurityTopTLDSpoofGetParamsTLSVersionTlSv1_0 EmailSecurityTopTLDSpoofGetParamsTLSVersion = "TLSv1_0"
	EmailSecurityTopTLDSpoofGetParamsTLSVersionTlSv1_1 EmailSecurityTopTLDSpoofGetParamsTLSVersion = "TLSv1_1"
	EmailSecurityTopTLDSpoofGetParamsTLSVersionTlSv1_2 EmailSecurityTopTLDSpoofGetParamsTLSVersion = "TLSv1_2"
	EmailSecurityTopTLDSpoofGetParamsTLSVersionTlSv1_3 EmailSecurityTopTLDSpoofGetParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecurityTopTLDSpoofGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecurityTopTLDSpoofGetParamsTLSVersionTlSv1_0, EmailSecurityTopTLDSpoofGetParamsTLSVersionTlSv1_1, EmailSecurityTopTLDSpoofGetParamsTLSVersionTlSv1_2, EmailSecurityTopTLDSpoofGetParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

type EmailSecurityTopTLDSpoofGetResponseEnvelope struct {
	Result  EmailSecurityTopTLDSpoofGetResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    emailSecurityTopTLDSpoofGetResponseEnvelopeJSON `json:"-"`
}

// emailSecurityTopTLDSpoofGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [EmailSecurityTopTLDSpoofGetResponseEnvelope]
type emailSecurityTopTLDSpoofGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTLDSpoofGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTLDSpoofGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
