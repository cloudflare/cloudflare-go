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

// EmailSecurityService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailSecurityService] method instead.
type EmailSecurityService struct {
	Options          []option.RequestOption
	Top              *EmailSecurityTopService
	Summary          *EmailSecuritySummaryService
	TimeseriesGroups *EmailSecurityTimeseriesGroupService
}

// NewEmailSecurityService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmailSecurityService(opts ...option.RequestOption) (r *EmailSecurityService) {
	r = &EmailSecurityService{}
	r.Options = opts
	r.Top = NewEmailSecurityTopService(opts...)
	r.Summary = NewEmailSecuritySummaryService(opts...)
	r.TimeseriesGroups = NewEmailSecurityTimeseriesGroupService(opts...)
	return
}

// Retrieves the distribution of email security metrics by the specified dimension.
func (r *EmailSecurityService) SummaryV2(ctx context.Context, dimension EmailSecuritySummaryV2ParamsDimension, query EmailSecuritySummaryV2Params, opts ...option.RequestOption) (res *EmailSecuritySummaryV2Response, err error) {
	var env EmailSecuritySummaryV2ResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/email/security/summary/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of email security metrics grouped by dimension over
// time.
func (r *EmailSecurityService) TimeseriesGroupsV2(ctx context.Context, dimension EmailSecurityTimeseriesGroupsV2ParamsDimension, query EmailSecurityTimeseriesGroupsV2Params, opts ...option.RequestOption) (res *EmailSecurityTimeseriesGroupsV2Response, err error) {
	var env EmailSecurityTimeseriesGroupsV2ResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/email/security/timeseries_groups/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailSecuritySummaryV2Response struct {
	// Metadata for the results.
	Meta     EmailSecuritySummaryV2ResponseMeta `json:"meta,required"`
	Summary0 map[string]string                  `json:"summary_0,required"`
	JSON     emailSecuritySummaryV2ResponseJSON `json:"-"`
}

// emailSecuritySummaryV2ResponseJSON contains the JSON metadata for the struct
// [EmailSecuritySummaryV2Response]
type emailSecuritySummaryV2ResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryV2Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryV2ResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailSecuritySummaryV2ResponseMeta struct {
	ConfidenceInfo EmailSecuritySummaryV2ResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailSecuritySummaryV2ResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailSecuritySummaryV2ResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailSecuritySummaryV2ResponseMetaUnit `json:"units,required"`
	JSON  emailSecuritySummaryV2ResponseMetaJSON   `json:"-"`
}

// emailSecuritySummaryV2ResponseMetaJSON contains the JSON metadata for the struct
// [EmailSecuritySummaryV2ResponseMeta]
type emailSecuritySummaryV2ResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecuritySummaryV2ResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryV2ResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryV2ResponseMetaConfidenceInfo struct {
	Annotations []EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                `json:"level,required"`
	JSON  emailSecuritySummaryV2ResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailSecuritySummaryV2ResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [EmailSecuritySummaryV2ResponseMetaConfidenceInfo]
type emailSecuritySummaryV2ResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryV2ResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryV2ResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                `json:"description,required"`
	EndDate     time.Time                                                             `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                           `json:"isInstantaneous,required"`
	LinkedURL       string                                                         `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                      `json:"startDate,required" format:"date-time"`
	JSON            emailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotation]
type emailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAll                EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBGP                EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBots               EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceCT                 EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNS                EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDos                EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFw                 EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceIQI                EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceNet                EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAll, EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBGP, EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBots, EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceCT, EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNS, EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDos, EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFw, EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceIQI, EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceNet, EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsEventType string

const (
	EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeEvent             EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeOutage            EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePipeline          EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeEvent, EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeOutage, EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePipeline, EmailSecuritySummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type EmailSecuritySummaryV2ResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      emailSecuritySummaryV2ResponseMetaDateRangeJSON `json:"-"`
}

// emailSecuritySummaryV2ResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [EmailSecuritySummaryV2ResponseMetaDateRange]
type emailSecuritySummaryV2ResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryV2ResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryV2ResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailSecuritySummaryV2ResponseMetaNormalization string

const (
	EmailSecuritySummaryV2ResponseMetaNormalizationPercentage           EmailSecuritySummaryV2ResponseMetaNormalization = "PERCENTAGE"
	EmailSecuritySummaryV2ResponseMetaNormalizationMin0Max              EmailSecuritySummaryV2ResponseMetaNormalization = "MIN0_MAX"
	EmailSecuritySummaryV2ResponseMetaNormalizationMinMax               EmailSecuritySummaryV2ResponseMetaNormalization = "MIN_MAX"
	EmailSecuritySummaryV2ResponseMetaNormalizationRawValues            EmailSecuritySummaryV2ResponseMetaNormalization = "RAW_VALUES"
	EmailSecuritySummaryV2ResponseMetaNormalizationPercentageChange     EmailSecuritySummaryV2ResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailSecuritySummaryV2ResponseMetaNormalizationRollingAverage       EmailSecuritySummaryV2ResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailSecuritySummaryV2ResponseMetaNormalizationOverlappedPercentage EmailSecuritySummaryV2ResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailSecuritySummaryV2ResponseMetaNormalizationRatio                EmailSecuritySummaryV2ResponseMetaNormalization = "RATIO"
)

func (r EmailSecuritySummaryV2ResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryV2ResponseMetaNormalizationPercentage, EmailSecuritySummaryV2ResponseMetaNormalizationMin0Max, EmailSecuritySummaryV2ResponseMetaNormalizationMinMax, EmailSecuritySummaryV2ResponseMetaNormalizationRawValues, EmailSecuritySummaryV2ResponseMetaNormalizationPercentageChange, EmailSecuritySummaryV2ResponseMetaNormalizationRollingAverage, EmailSecuritySummaryV2ResponseMetaNormalizationOverlappedPercentage, EmailSecuritySummaryV2ResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailSecuritySummaryV2ResponseMetaUnit struct {
	Name  string                                     `json:"name,required"`
	Value string                                     `json:"value,required"`
	JSON  emailSecuritySummaryV2ResponseMetaUnitJSON `json:"-"`
}

// emailSecuritySummaryV2ResponseMetaUnitJSON contains the JSON metadata for the
// struct [EmailSecuritySummaryV2ResponseMetaUnit]
type emailSecuritySummaryV2ResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryV2ResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryV2ResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupsV2Response struct {
	// Metadata for the results.
	Meta   EmailSecurityTimeseriesGroupsV2ResponseMeta   `json:"meta,required"`
	Serie0 EmailSecurityTimeseriesGroupsV2ResponseSerie0 `json:"serie_0,required"`
	JSON   emailSecurityTimeseriesGroupsV2ResponseJSON   `json:"-"`
}

// emailSecurityTimeseriesGroupsV2ResponseJSON contains the JSON metadata for the
// struct [EmailSecurityTimeseriesGroupsV2Response]
type emailSecurityTimeseriesGroupsV2ResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupsV2Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupsV2ResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailSecurityTimeseriesGroupsV2ResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    EmailSecurityTimeseriesGroupsV2ResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailSecurityTimeseriesGroupsV2ResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailSecurityTimeseriesGroupsV2ResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailSecurityTimeseriesGroupsV2ResponseMetaUnit `json:"units,required"`
	JSON  emailSecurityTimeseriesGroupsV2ResponseMetaJSON   `json:"-"`
}

// emailSecurityTimeseriesGroupsV2ResponseMetaJSON contains the JSON metadata for
// the struct [EmailSecurityTimeseriesGroupsV2ResponseMeta]
type emailSecurityTimeseriesGroupsV2ResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupsV2ResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupsV2ResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailSecurityTimeseriesGroupsV2ResponseMetaAggInterval string

const (
	EmailSecurityTimeseriesGroupsV2ResponseMetaAggIntervalFifteenMinutes EmailSecurityTimeseriesGroupsV2ResponseMetaAggInterval = "FIFTEEN_MINUTES"
	EmailSecurityTimeseriesGroupsV2ResponseMetaAggIntervalOneHour        EmailSecurityTimeseriesGroupsV2ResponseMetaAggInterval = "ONE_HOUR"
	EmailSecurityTimeseriesGroupsV2ResponseMetaAggIntervalOneDay         EmailSecurityTimeseriesGroupsV2ResponseMetaAggInterval = "ONE_DAY"
	EmailSecurityTimeseriesGroupsV2ResponseMetaAggIntervalOneWeek        EmailSecurityTimeseriesGroupsV2ResponseMetaAggInterval = "ONE_WEEK"
	EmailSecurityTimeseriesGroupsV2ResponseMetaAggIntervalOneMonth       EmailSecurityTimeseriesGroupsV2ResponseMetaAggInterval = "ONE_MONTH"
)

func (r EmailSecurityTimeseriesGroupsV2ResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupsV2ResponseMetaAggIntervalFifteenMinutes, EmailSecurityTimeseriesGroupsV2ResponseMetaAggIntervalOneHour, EmailSecurityTimeseriesGroupsV2ResponseMetaAggIntervalOneDay, EmailSecurityTimeseriesGroupsV2ResponseMetaAggIntervalOneWeek, EmailSecurityTimeseriesGroupsV2ResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfo struct {
	Annotations []EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                         `json:"level,required"`
	JSON  emailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfo]
type emailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                         `json:"description,required"`
	EndDate     time.Time                                                                      `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                    `json:"isInstantaneous,required"`
	LinkedURL       string                                                                  `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                               `json:"startDate,required" format:"date-time"`
	JSON            emailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation]
type emailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAll                EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceBGP                EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceBots               EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceCT                 EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNS                EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDos                EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceFw                 EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceIQI                EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceNet                EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAll, EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceBGP, EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceBots, EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceCT, EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNS, EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDos, EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceFw, EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceIQI, EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceNet, EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType string

const (
	EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeEvent             EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeOutage            EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypePipeline          EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeEvent, EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeOutage, EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypePipeline, EmailSecurityTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupsV2ResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                `json:"startTime,required" format:"date-time"`
	JSON      emailSecurityTimeseriesGroupsV2ResponseMetaDateRangeJSON `json:"-"`
}

// emailSecurityTimeseriesGroupsV2ResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [EmailSecurityTimeseriesGroupsV2ResponseMetaDateRange]
type emailSecurityTimeseriesGroupsV2ResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupsV2ResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupsV2ResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailSecurityTimeseriesGroupsV2ResponseMetaNormalization string

const (
	EmailSecurityTimeseriesGroupsV2ResponseMetaNormalizationPercentage           EmailSecurityTimeseriesGroupsV2ResponseMetaNormalization = "PERCENTAGE"
	EmailSecurityTimeseriesGroupsV2ResponseMetaNormalizationMin0Max              EmailSecurityTimeseriesGroupsV2ResponseMetaNormalization = "MIN0_MAX"
	EmailSecurityTimeseriesGroupsV2ResponseMetaNormalizationMinMax               EmailSecurityTimeseriesGroupsV2ResponseMetaNormalization = "MIN_MAX"
	EmailSecurityTimeseriesGroupsV2ResponseMetaNormalizationRawValues            EmailSecurityTimeseriesGroupsV2ResponseMetaNormalization = "RAW_VALUES"
	EmailSecurityTimeseriesGroupsV2ResponseMetaNormalizationPercentageChange     EmailSecurityTimeseriesGroupsV2ResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailSecurityTimeseriesGroupsV2ResponseMetaNormalizationRollingAverage       EmailSecurityTimeseriesGroupsV2ResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailSecurityTimeseriesGroupsV2ResponseMetaNormalizationOverlappedPercentage EmailSecurityTimeseriesGroupsV2ResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailSecurityTimeseriesGroupsV2ResponseMetaNormalizationRatio                EmailSecurityTimeseriesGroupsV2ResponseMetaNormalization = "RATIO"
)

func (r EmailSecurityTimeseriesGroupsV2ResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupsV2ResponseMetaNormalizationPercentage, EmailSecurityTimeseriesGroupsV2ResponseMetaNormalizationMin0Max, EmailSecurityTimeseriesGroupsV2ResponseMetaNormalizationMinMax, EmailSecurityTimeseriesGroupsV2ResponseMetaNormalizationRawValues, EmailSecurityTimeseriesGroupsV2ResponseMetaNormalizationPercentageChange, EmailSecurityTimeseriesGroupsV2ResponseMetaNormalizationRollingAverage, EmailSecurityTimeseriesGroupsV2ResponseMetaNormalizationOverlappedPercentage, EmailSecurityTimeseriesGroupsV2ResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupsV2ResponseMetaUnit struct {
	Name  string                                              `json:"name,required"`
	Value string                                              `json:"value,required"`
	JSON  emailSecurityTimeseriesGroupsV2ResponseMetaUnitJSON `json:"-"`
}

// emailSecurityTimeseriesGroupsV2ResponseMetaUnitJSON contains the JSON metadata
// for the struct [EmailSecurityTimeseriesGroupsV2ResponseMetaUnit]
type emailSecurityTimeseriesGroupsV2ResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupsV2ResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupsV2ResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupsV2ResponseSerie0 struct {
	Timestamps  []time.Time                                       `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                               `json:"-,extras"`
	JSON        emailSecurityTimeseriesGroupsV2ResponseSerie0JSON `json:"-"`
}

// emailSecurityTimeseriesGroupsV2ResponseSerie0JSON contains the JSON metadata for
// the struct [EmailSecurityTimeseriesGroupsV2ResponseSerie0]
type emailSecurityTimeseriesGroupsV2ResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupsV2ResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupsV2ResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryV2Params struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailSecuritySummaryV2ParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailSecuritySummaryV2ParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailSecuritySummaryV2ParamsDMARC] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[EmailSecuritySummaryV2ParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailSecuritySummaryV2ParamsSPF] `query:"spf"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]EmailSecuritySummaryV2ParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecuritySummaryV2Params]'s query parameters as
// `url.Values`.
func (r EmailSecuritySummaryV2Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the attribute by which to group the results.
type EmailSecuritySummaryV2ParamsDimension string

const (
	EmailSecuritySummaryV2ParamsDimensionSpam           EmailSecuritySummaryV2ParamsDimension = "SPAM"
	EmailSecuritySummaryV2ParamsDimensionMalicious      EmailSecuritySummaryV2ParamsDimension = "MALICIOUS"
	EmailSecuritySummaryV2ParamsDimensionSpoof          EmailSecuritySummaryV2ParamsDimension = "SPOOF"
	EmailSecuritySummaryV2ParamsDimensionThreatCategory EmailSecuritySummaryV2ParamsDimension = "THREAT_CATEGORY"
	EmailSecuritySummaryV2ParamsDimensionARC            EmailSecuritySummaryV2ParamsDimension = "ARC"
	EmailSecuritySummaryV2ParamsDimensionDKIM           EmailSecuritySummaryV2ParamsDimension = "DKIM"
	EmailSecuritySummaryV2ParamsDimensionDMARC          EmailSecuritySummaryV2ParamsDimension = "DMARC"
	EmailSecuritySummaryV2ParamsDimensionSPF            EmailSecuritySummaryV2ParamsDimension = "SPF"
	EmailSecuritySummaryV2ParamsDimensionTLSVersion     EmailSecuritySummaryV2ParamsDimension = "TLS_VERSION"
)

func (r EmailSecuritySummaryV2ParamsDimension) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryV2ParamsDimensionSpam, EmailSecuritySummaryV2ParamsDimensionMalicious, EmailSecuritySummaryV2ParamsDimensionSpoof, EmailSecuritySummaryV2ParamsDimensionThreatCategory, EmailSecuritySummaryV2ParamsDimensionARC, EmailSecuritySummaryV2ParamsDimensionDKIM, EmailSecuritySummaryV2ParamsDimensionDMARC, EmailSecuritySummaryV2ParamsDimensionSPF, EmailSecuritySummaryV2ParamsDimensionTLSVersion:
		return true
	}
	return false
}

type EmailSecuritySummaryV2ParamsARC string

const (
	EmailSecuritySummaryV2ParamsARCPass EmailSecuritySummaryV2ParamsARC = "PASS"
	EmailSecuritySummaryV2ParamsARCNone EmailSecuritySummaryV2ParamsARC = "NONE"
	EmailSecuritySummaryV2ParamsARCFail EmailSecuritySummaryV2ParamsARC = "FAIL"
)

func (r EmailSecuritySummaryV2ParamsARC) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryV2ParamsARCPass, EmailSecuritySummaryV2ParamsARCNone, EmailSecuritySummaryV2ParamsARCFail:
		return true
	}
	return false
}

type EmailSecuritySummaryV2ParamsDKIM string

const (
	EmailSecuritySummaryV2ParamsDKIMPass EmailSecuritySummaryV2ParamsDKIM = "PASS"
	EmailSecuritySummaryV2ParamsDKIMNone EmailSecuritySummaryV2ParamsDKIM = "NONE"
	EmailSecuritySummaryV2ParamsDKIMFail EmailSecuritySummaryV2ParamsDKIM = "FAIL"
)

func (r EmailSecuritySummaryV2ParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryV2ParamsDKIMPass, EmailSecuritySummaryV2ParamsDKIMNone, EmailSecuritySummaryV2ParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecuritySummaryV2ParamsDMARC string

const (
	EmailSecuritySummaryV2ParamsDMARCPass EmailSecuritySummaryV2ParamsDMARC = "PASS"
	EmailSecuritySummaryV2ParamsDMARCNone EmailSecuritySummaryV2ParamsDMARC = "NONE"
	EmailSecuritySummaryV2ParamsDMARCFail EmailSecuritySummaryV2ParamsDMARC = "FAIL"
)

func (r EmailSecuritySummaryV2ParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryV2ParamsDMARCPass, EmailSecuritySummaryV2ParamsDMARCNone, EmailSecuritySummaryV2ParamsDMARCFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailSecuritySummaryV2ParamsFormat string

const (
	EmailSecuritySummaryV2ParamsFormatJson EmailSecuritySummaryV2ParamsFormat = "JSON"
	EmailSecuritySummaryV2ParamsFormatCsv  EmailSecuritySummaryV2ParamsFormat = "CSV"
)

func (r EmailSecuritySummaryV2ParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryV2ParamsFormatJson, EmailSecuritySummaryV2ParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecuritySummaryV2ParamsSPF string

const (
	EmailSecuritySummaryV2ParamsSPFPass EmailSecuritySummaryV2ParamsSPF = "PASS"
	EmailSecuritySummaryV2ParamsSPFNone EmailSecuritySummaryV2ParamsSPF = "NONE"
	EmailSecuritySummaryV2ParamsSPFFail EmailSecuritySummaryV2ParamsSPF = "FAIL"
)

func (r EmailSecuritySummaryV2ParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryV2ParamsSPFPass, EmailSecuritySummaryV2ParamsSPFNone, EmailSecuritySummaryV2ParamsSPFFail:
		return true
	}
	return false
}

type EmailSecuritySummaryV2ParamsTLSVersion string

const (
	EmailSecuritySummaryV2ParamsTLSVersionTlSv1_0 EmailSecuritySummaryV2ParamsTLSVersion = "TLSv1_0"
	EmailSecuritySummaryV2ParamsTLSVersionTlSv1_1 EmailSecuritySummaryV2ParamsTLSVersion = "TLSv1_1"
	EmailSecuritySummaryV2ParamsTLSVersionTlSv1_2 EmailSecuritySummaryV2ParamsTLSVersion = "TLSv1_2"
	EmailSecuritySummaryV2ParamsTLSVersionTlSv1_3 EmailSecuritySummaryV2ParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecuritySummaryV2ParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryV2ParamsTLSVersionTlSv1_0, EmailSecuritySummaryV2ParamsTLSVersionTlSv1_1, EmailSecuritySummaryV2ParamsTLSVersionTlSv1_2, EmailSecuritySummaryV2ParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

type EmailSecuritySummaryV2ResponseEnvelope struct {
	Result  EmailSecuritySummaryV2Response             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    emailSecuritySummaryV2ResponseEnvelopeJSON `json:"-"`
}

// emailSecuritySummaryV2ResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailSecuritySummaryV2ResponseEnvelope]
type emailSecuritySummaryV2ResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryV2ResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryV2ResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupsV2Params struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailSecurityTimeseriesGroupsV2ParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailSecurityTimeseriesGroupsV2ParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailSecurityTimeseriesGroupsV2ParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailSecurityTimeseriesGroupsV2ParamsDMARC] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[EmailSecurityTimeseriesGroupsV2ParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailSecurityTimeseriesGroupsV2ParamsSPF] `query:"spf"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]EmailSecurityTimeseriesGroupsV2ParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecurityTimeseriesGroupsV2Params]'s query parameters
// as `url.Values`.
func (r EmailSecurityTimeseriesGroupsV2Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the attribute by which to group the results.
type EmailSecurityTimeseriesGroupsV2ParamsDimension string

const (
	EmailSecurityTimeseriesGroupsV2ParamsDimensionSpam           EmailSecurityTimeseriesGroupsV2ParamsDimension = "SPAM"
	EmailSecurityTimeseriesGroupsV2ParamsDimensionMalicious      EmailSecurityTimeseriesGroupsV2ParamsDimension = "MALICIOUS"
	EmailSecurityTimeseriesGroupsV2ParamsDimensionSpoof          EmailSecurityTimeseriesGroupsV2ParamsDimension = "SPOOF"
	EmailSecurityTimeseriesGroupsV2ParamsDimensionThreatCategory EmailSecurityTimeseriesGroupsV2ParamsDimension = "THREAT_CATEGORY"
	EmailSecurityTimeseriesGroupsV2ParamsDimensionARC            EmailSecurityTimeseriesGroupsV2ParamsDimension = "ARC"
	EmailSecurityTimeseriesGroupsV2ParamsDimensionDKIM           EmailSecurityTimeseriesGroupsV2ParamsDimension = "DKIM"
	EmailSecurityTimeseriesGroupsV2ParamsDimensionDMARC          EmailSecurityTimeseriesGroupsV2ParamsDimension = "DMARC"
	EmailSecurityTimeseriesGroupsV2ParamsDimensionSPF            EmailSecurityTimeseriesGroupsV2ParamsDimension = "SPF"
	EmailSecurityTimeseriesGroupsV2ParamsDimensionTLSVersion     EmailSecurityTimeseriesGroupsV2ParamsDimension = "TLS_VERSION"
)

func (r EmailSecurityTimeseriesGroupsV2ParamsDimension) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupsV2ParamsDimensionSpam, EmailSecurityTimeseriesGroupsV2ParamsDimensionMalicious, EmailSecurityTimeseriesGroupsV2ParamsDimensionSpoof, EmailSecurityTimeseriesGroupsV2ParamsDimensionThreatCategory, EmailSecurityTimeseriesGroupsV2ParamsDimensionARC, EmailSecurityTimeseriesGroupsV2ParamsDimensionDKIM, EmailSecurityTimeseriesGroupsV2ParamsDimensionDMARC, EmailSecurityTimeseriesGroupsV2ParamsDimensionSPF, EmailSecurityTimeseriesGroupsV2ParamsDimensionTLSVersion:
		return true
	}
	return false
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailSecurityTimeseriesGroupsV2ParamsAggInterval string

const (
	EmailSecurityTimeseriesGroupsV2ParamsAggInterval15m EmailSecurityTimeseriesGroupsV2ParamsAggInterval = "15m"
	EmailSecurityTimeseriesGroupsV2ParamsAggInterval1h  EmailSecurityTimeseriesGroupsV2ParamsAggInterval = "1h"
	EmailSecurityTimeseriesGroupsV2ParamsAggInterval1d  EmailSecurityTimeseriesGroupsV2ParamsAggInterval = "1d"
	EmailSecurityTimeseriesGroupsV2ParamsAggInterval1w  EmailSecurityTimeseriesGroupsV2ParamsAggInterval = "1w"
)

func (r EmailSecurityTimeseriesGroupsV2ParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupsV2ParamsAggInterval15m, EmailSecurityTimeseriesGroupsV2ParamsAggInterval1h, EmailSecurityTimeseriesGroupsV2ParamsAggInterval1d, EmailSecurityTimeseriesGroupsV2ParamsAggInterval1w:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupsV2ParamsARC string

const (
	EmailSecurityTimeseriesGroupsV2ParamsARCPass EmailSecurityTimeseriesGroupsV2ParamsARC = "PASS"
	EmailSecurityTimeseriesGroupsV2ParamsARCNone EmailSecurityTimeseriesGroupsV2ParamsARC = "NONE"
	EmailSecurityTimeseriesGroupsV2ParamsARCFail EmailSecurityTimeseriesGroupsV2ParamsARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupsV2ParamsARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupsV2ParamsARCPass, EmailSecurityTimeseriesGroupsV2ParamsARCNone, EmailSecurityTimeseriesGroupsV2ParamsARCFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupsV2ParamsDKIM string

const (
	EmailSecurityTimeseriesGroupsV2ParamsDKIMPass EmailSecurityTimeseriesGroupsV2ParamsDKIM = "PASS"
	EmailSecurityTimeseriesGroupsV2ParamsDKIMNone EmailSecurityTimeseriesGroupsV2ParamsDKIM = "NONE"
	EmailSecurityTimeseriesGroupsV2ParamsDKIMFail EmailSecurityTimeseriesGroupsV2ParamsDKIM = "FAIL"
)

func (r EmailSecurityTimeseriesGroupsV2ParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupsV2ParamsDKIMPass, EmailSecurityTimeseriesGroupsV2ParamsDKIMNone, EmailSecurityTimeseriesGroupsV2ParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupsV2ParamsDMARC string

const (
	EmailSecurityTimeseriesGroupsV2ParamsDMARCPass EmailSecurityTimeseriesGroupsV2ParamsDMARC = "PASS"
	EmailSecurityTimeseriesGroupsV2ParamsDMARCNone EmailSecurityTimeseriesGroupsV2ParamsDMARC = "NONE"
	EmailSecurityTimeseriesGroupsV2ParamsDMARCFail EmailSecurityTimeseriesGroupsV2ParamsDMARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupsV2ParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupsV2ParamsDMARCPass, EmailSecurityTimeseriesGroupsV2ParamsDMARCNone, EmailSecurityTimeseriesGroupsV2ParamsDMARCFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailSecurityTimeseriesGroupsV2ParamsFormat string

const (
	EmailSecurityTimeseriesGroupsV2ParamsFormatJson EmailSecurityTimeseriesGroupsV2ParamsFormat = "JSON"
	EmailSecurityTimeseriesGroupsV2ParamsFormatCsv  EmailSecurityTimeseriesGroupsV2ParamsFormat = "CSV"
)

func (r EmailSecurityTimeseriesGroupsV2ParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupsV2ParamsFormatJson, EmailSecurityTimeseriesGroupsV2ParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupsV2ParamsSPF string

const (
	EmailSecurityTimeseriesGroupsV2ParamsSPFPass EmailSecurityTimeseriesGroupsV2ParamsSPF = "PASS"
	EmailSecurityTimeseriesGroupsV2ParamsSPFNone EmailSecurityTimeseriesGroupsV2ParamsSPF = "NONE"
	EmailSecurityTimeseriesGroupsV2ParamsSPFFail EmailSecurityTimeseriesGroupsV2ParamsSPF = "FAIL"
)

func (r EmailSecurityTimeseriesGroupsV2ParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupsV2ParamsSPFPass, EmailSecurityTimeseriesGroupsV2ParamsSPFNone, EmailSecurityTimeseriesGroupsV2ParamsSPFFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupsV2ParamsTLSVersion string

const (
	EmailSecurityTimeseriesGroupsV2ParamsTLSVersionTlSv1_0 EmailSecurityTimeseriesGroupsV2ParamsTLSVersion = "TLSv1_0"
	EmailSecurityTimeseriesGroupsV2ParamsTLSVersionTlSv1_1 EmailSecurityTimeseriesGroupsV2ParamsTLSVersion = "TLSv1_1"
	EmailSecurityTimeseriesGroupsV2ParamsTLSVersionTlSv1_2 EmailSecurityTimeseriesGroupsV2ParamsTLSVersion = "TLSv1_2"
	EmailSecurityTimeseriesGroupsV2ParamsTLSVersionTlSv1_3 EmailSecurityTimeseriesGroupsV2ParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecurityTimeseriesGroupsV2ParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupsV2ParamsTLSVersionTlSv1_0, EmailSecurityTimeseriesGroupsV2ParamsTLSVersionTlSv1_1, EmailSecurityTimeseriesGroupsV2ParamsTLSVersionTlSv1_2, EmailSecurityTimeseriesGroupsV2ParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupsV2ResponseEnvelope struct {
	Result  EmailSecurityTimeseriesGroupsV2Response             `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    emailSecurityTimeseriesGroupsV2ResponseEnvelopeJSON `json:"-"`
}

// emailSecurityTimeseriesGroupsV2ResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailSecurityTimeseriesGroupsV2ResponseEnvelope]
type emailSecurityTimeseriesGroupsV2ResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupsV2ResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupsV2ResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
