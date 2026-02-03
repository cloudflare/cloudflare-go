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

// AIInferenceService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIInferenceService] method instead.
type AIInferenceService struct {
	Options          []option.RequestOption
	Summary          *AIInferenceSummaryService
	TimeseriesGroups *AIInferenceTimeseriesGroupService
}

// NewAIInferenceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIInferenceService(opts ...option.RequestOption) (r *AIInferenceService) {
	r = &AIInferenceService{}
	r.Options = opts
	r.Summary = NewAIInferenceSummaryService(opts...)
	r.TimeseriesGroups = NewAIInferenceTimeseriesGroupService(opts...)
	return
}

// Retrieves an aggregated summary of unique accounts using Workers AI inference
// grouped by the specified dimension.
func (r *AIInferenceService) SummaryV2(ctx context.Context, dimension AIInferenceSummaryV2ParamsDimension, query AIInferenceSummaryV2Params, opts ...option.RequestOption) (res *AIInferenceSummaryV2Response, err error) {
	var env AIInferenceSummaryV2ResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/ai/inference/summary/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of unique accounts using Workers AI inference,
// grouped by the specified dimension over time.
func (r *AIInferenceService) TimeseriesGroupsV2(ctx context.Context, dimension AIInferenceTimeseriesGroupsV2ParamsDimension, query AIInferenceTimeseriesGroupsV2Params, opts ...option.RequestOption) (res *AIInferenceTimeseriesGroupsV2Response, err error) {
	var env AIInferenceTimeseriesGroupsV2ResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/ai/inference/timeseries_groups/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AIInferenceSummaryV2Response struct {
	// Metadata for the results.
	Meta     AIInferenceSummaryV2ResponseMeta `json:"meta,required"`
	Summary0 map[string]string                `json:"summary_0,required"`
	JSON     aiInferenceSummaryV2ResponseJSON `json:"-"`
}

// aiInferenceSummaryV2ResponseJSON contains the JSON metadata for the struct
// [AIInferenceSummaryV2Response]
type aiInferenceSummaryV2ResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceSummaryV2Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryV2ResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AIInferenceSummaryV2ResponseMeta struct {
	ConfidenceInfo AIInferenceSummaryV2ResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AIInferenceSummaryV2ResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AIInferenceSummaryV2ResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AIInferenceSummaryV2ResponseMetaUnit `json:"units,required"`
	JSON  aiInferenceSummaryV2ResponseMetaJSON   `json:"-"`
}

// aiInferenceSummaryV2ResponseMetaJSON contains the JSON metadata for the struct
// [AIInferenceSummaryV2ResponseMeta]
type aiInferenceSummaryV2ResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AIInferenceSummaryV2ResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryV2ResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AIInferenceSummaryV2ResponseMetaConfidenceInfo struct {
	Annotations []AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                              `json:"level,required"`
	JSON  aiInferenceSummaryV2ResponseMetaConfidenceInfoJSON `json:"-"`
}

// aiInferenceSummaryV2ResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AIInferenceSummaryV2ResponseMetaConfidenceInfo]
type aiInferenceSummaryV2ResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceSummaryV2ResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryV2ResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                              `json:"description,required"`
	EndDate     time.Time                                                           `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                         `json:"isInstantaneous,required"`
	LinkedURL       string                                                       `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                    `json:"startDate,required" format:"date-time"`
	JSON            aiInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// aiInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotation]
type aiInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAll                AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBGP                AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBots               AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceCT                 AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNS                AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDos                AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFw                 AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceIQI                AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceNet                AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAll, AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBGP, AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBots, AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceCT, AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNS, AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDos, AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFw, AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceIQI, AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceNet, AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType string

const (
	AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeEvent             AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeOutage            AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePipeline          AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeEvent, AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeOutage, AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePipeline, AIInferenceSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type AIInferenceSummaryV2ResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                     `json:"startTime,required" format:"date-time"`
	JSON      aiInferenceSummaryV2ResponseMetaDateRangeJSON `json:"-"`
}

// aiInferenceSummaryV2ResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [AIInferenceSummaryV2ResponseMetaDateRange]
type aiInferenceSummaryV2ResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceSummaryV2ResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryV2ResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AIInferenceSummaryV2ResponseMetaNormalization string

const (
	AIInferenceSummaryV2ResponseMetaNormalizationPercentage           AIInferenceSummaryV2ResponseMetaNormalization = "PERCENTAGE"
	AIInferenceSummaryV2ResponseMetaNormalizationMin0Max              AIInferenceSummaryV2ResponseMetaNormalization = "MIN0_MAX"
	AIInferenceSummaryV2ResponseMetaNormalizationMinMax               AIInferenceSummaryV2ResponseMetaNormalization = "MIN_MAX"
	AIInferenceSummaryV2ResponseMetaNormalizationRawValues            AIInferenceSummaryV2ResponseMetaNormalization = "RAW_VALUES"
	AIInferenceSummaryV2ResponseMetaNormalizationPercentageChange     AIInferenceSummaryV2ResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AIInferenceSummaryV2ResponseMetaNormalizationRollingAverage       AIInferenceSummaryV2ResponseMetaNormalization = "ROLLING_AVERAGE"
	AIInferenceSummaryV2ResponseMetaNormalizationOverlappedPercentage AIInferenceSummaryV2ResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AIInferenceSummaryV2ResponseMetaNormalizationRatio                AIInferenceSummaryV2ResponseMetaNormalization = "RATIO"
)

func (r AIInferenceSummaryV2ResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AIInferenceSummaryV2ResponseMetaNormalizationPercentage, AIInferenceSummaryV2ResponseMetaNormalizationMin0Max, AIInferenceSummaryV2ResponseMetaNormalizationMinMax, AIInferenceSummaryV2ResponseMetaNormalizationRawValues, AIInferenceSummaryV2ResponseMetaNormalizationPercentageChange, AIInferenceSummaryV2ResponseMetaNormalizationRollingAverage, AIInferenceSummaryV2ResponseMetaNormalizationOverlappedPercentage, AIInferenceSummaryV2ResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AIInferenceSummaryV2ResponseMetaUnit struct {
	Name  string                                   `json:"name,required"`
	Value string                                   `json:"value,required"`
	JSON  aiInferenceSummaryV2ResponseMetaUnitJSON `json:"-"`
}

// aiInferenceSummaryV2ResponseMetaUnitJSON contains the JSON metadata for the
// struct [AIInferenceSummaryV2ResponseMetaUnit]
type aiInferenceSummaryV2ResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceSummaryV2ResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryV2ResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AIInferenceTimeseriesGroupsV2Response struct {
	// Metadata for the results.
	Meta   AIInferenceTimeseriesGroupsV2ResponseMeta   `json:"meta,required"`
	Serie0 AIInferenceTimeseriesGroupsV2ResponseSerie0 `json:"serie_0,required"`
	JSON   aiInferenceTimeseriesGroupsV2ResponseJSON   `json:"-"`
}

// aiInferenceTimeseriesGroupsV2ResponseJSON contains the JSON metadata for the
// struct [AIInferenceTimeseriesGroupsV2Response]
type aiInferenceTimeseriesGroupsV2ResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceTimeseriesGroupsV2Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceTimeseriesGroupsV2ResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AIInferenceTimeseriesGroupsV2ResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AIInferenceTimeseriesGroupsV2ResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AIInferenceTimeseriesGroupsV2ResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AIInferenceTimeseriesGroupsV2ResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AIInferenceTimeseriesGroupsV2ResponseMetaUnit `json:"units,required"`
	JSON  aiInferenceTimeseriesGroupsV2ResponseMetaJSON   `json:"-"`
}

// aiInferenceTimeseriesGroupsV2ResponseMetaJSON contains the JSON metadata for the
// struct [AIInferenceTimeseriesGroupsV2ResponseMeta]
type aiInferenceTimeseriesGroupsV2ResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AIInferenceTimeseriesGroupsV2ResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceTimeseriesGroupsV2ResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AIInferenceTimeseriesGroupsV2ResponseMetaAggInterval string

const (
	AIInferenceTimeseriesGroupsV2ResponseMetaAggIntervalFifteenMinutes AIInferenceTimeseriesGroupsV2ResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AIInferenceTimeseriesGroupsV2ResponseMetaAggIntervalOneHour        AIInferenceTimeseriesGroupsV2ResponseMetaAggInterval = "ONE_HOUR"
	AIInferenceTimeseriesGroupsV2ResponseMetaAggIntervalOneDay         AIInferenceTimeseriesGroupsV2ResponseMetaAggInterval = "ONE_DAY"
	AIInferenceTimeseriesGroupsV2ResponseMetaAggIntervalOneWeek        AIInferenceTimeseriesGroupsV2ResponseMetaAggInterval = "ONE_WEEK"
	AIInferenceTimeseriesGroupsV2ResponseMetaAggIntervalOneMonth       AIInferenceTimeseriesGroupsV2ResponseMetaAggInterval = "ONE_MONTH"
)

func (r AIInferenceTimeseriesGroupsV2ResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AIInferenceTimeseriesGroupsV2ResponseMetaAggIntervalFifteenMinutes, AIInferenceTimeseriesGroupsV2ResponseMetaAggIntervalOneHour, AIInferenceTimeseriesGroupsV2ResponseMetaAggIntervalOneDay, AIInferenceTimeseriesGroupsV2ResponseMetaAggIntervalOneWeek, AIInferenceTimeseriesGroupsV2ResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfo struct {
	Annotations []AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                       `json:"level,required"`
	JSON  aiInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoJSON `json:"-"`
}

// aiInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfo]
type aiInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                       `json:"description,required"`
	EndDate     time.Time                                                                    `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                  `json:"isInstantaneous,required"`
	LinkedURL       string                                                                `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                             `json:"startDate,required" format:"date-time"`
	JSON            aiInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// aiInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation]
type aiInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAll                AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceBGP                AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceBots               AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceCT                 AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNS                AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDos                AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceFw                 AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceIQI                AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceNet                AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAll, AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceBGP, AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceBots, AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceCT, AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNS, AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDos, AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceFw, AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceIQI, AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceNet, AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType string

const (
	AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeEvent             AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeOutage            AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypePipeline          AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeEvent, AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeOutage, AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypePipeline, AIInferenceTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type AIInferenceTimeseriesGroupsV2ResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                              `json:"startTime,required" format:"date-time"`
	JSON      aiInferenceTimeseriesGroupsV2ResponseMetaDateRangeJSON `json:"-"`
}

// aiInferenceTimeseriesGroupsV2ResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [AIInferenceTimeseriesGroupsV2ResponseMetaDateRange]
type aiInferenceTimeseriesGroupsV2ResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceTimeseriesGroupsV2ResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceTimeseriesGroupsV2ResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AIInferenceTimeseriesGroupsV2ResponseMetaNormalization string

const (
	AIInferenceTimeseriesGroupsV2ResponseMetaNormalizationPercentage           AIInferenceTimeseriesGroupsV2ResponseMetaNormalization = "PERCENTAGE"
	AIInferenceTimeseriesGroupsV2ResponseMetaNormalizationMin0Max              AIInferenceTimeseriesGroupsV2ResponseMetaNormalization = "MIN0_MAX"
	AIInferenceTimeseriesGroupsV2ResponseMetaNormalizationMinMax               AIInferenceTimeseriesGroupsV2ResponseMetaNormalization = "MIN_MAX"
	AIInferenceTimeseriesGroupsV2ResponseMetaNormalizationRawValues            AIInferenceTimeseriesGroupsV2ResponseMetaNormalization = "RAW_VALUES"
	AIInferenceTimeseriesGroupsV2ResponseMetaNormalizationPercentageChange     AIInferenceTimeseriesGroupsV2ResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AIInferenceTimeseriesGroupsV2ResponseMetaNormalizationRollingAverage       AIInferenceTimeseriesGroupsV2ResponseMetaNormalization = "ROLLING_AVERAGE"
	AIInferenceTimeseriesGroupsV2ResponseMetaNormalizationOverlappedPercentage AIInferenceTimeseriesGroupsV2ResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AIInferenceTimeseriesGroupsV2ResponseMetaNormalizationRatio                AIInferenceTimeseriesGroupsV2ResponseMetaNormalization = "RATIO"
)

func (r AIInferenceTimeseriesGroupsV2ResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AIInferenceTimeseriesGroupsV2ResponseMetaNormalizationPercentage, AIInferenceTimeseriesGroupsV2ResponseMetaNormalizationMin0Max, AIInferenceTimeseriesGroupsV2ResponseMetaNormalizationMinMax, AIInferenceTimeseriesGroupsV2ResponseMetaNormalizationRawValues, AIInferenceTimeseriesGroupsV2ResponseMetaNormalizationPercentageChange, AIInferenceTimeseriesGroupsV2ResponseMetaNormalizationRollingAverage, AIInferenceTimeseriesGroupsV2ResponseMetaNormalizationOverlappedPercentage, AIInferenceTimeseriesGroupsV2ResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AIInferenceTimeseriesGroupsV2ResponseMetaUnit struct {
	Name  string                                            `json:"name,required"`
	Value string                                            `json:"value,required"`
	JSON  aiInferenceTimeseriesGroupsV2ResponseMetaUnitJSON `json:"-"`
}

// aiInferenceTimeseriesGroupsV2ResponseMetaUnitJSON contains the JSON metadata for
// the struct [AIInferenceTimeseriesGroupsV2ResponseMetaUnit]
type aiInferenceTimeseriesGroupsV2ResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceTimeseriesGroupsV2ResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceTimeseriesGroupsV2ResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AIInferenceTimeseriesGroupsV2ResponseSerie0 struct {
	Timestamps  []time.Time                                     `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                             `json:"-,extras"`
	JSON        aiInferenceTimeseriesGroupsV2ResponseSerie0JSON `json:"-"`
}

// aiInferenceTimeseriesGroupsV2ResponseSerie0JSON contains the JSON metadata for
// the struct [AIInferenceTimeseriesGroupsV2ResponseSerie0]
type aiInferenceTimeseriesGroupsV2ResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceTimeseriesGroupsV2ResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceTimeseriesGroupsV2ResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AIInferenceSummaryV2Params struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[AIInferenceSummaryV2ParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AIInferenceSummaryV2Params]'s query parameters as
// `url.Values`.
func (r AIInferenceSummaryV2Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the attribute by which to group the results.
type AIInferenceSummaryV2ParamsDimension string

const (
	AIInferenceSummaryV2ParamsDimensionModel AIInferenceSummaryV2ParamsDimension = "MODEL"
	AIInferenceSummaryV2ParamsDimensionTask  AIInferenceSummaryV2ParamsDimension = "TASK"
)

func (r AIInferenceSummaryV2ParamsDimension) IsKnown() bool {
	switch r {
	case AIInferenceSummaryV2ParamsDimensionModel, AIInferenceSummaryV2ParamsDimensionTask:
		return true
	}
	return false
}

// Format in which results will be returned.
type AIInferenceSummaryV2ParamsFormat string

const (
	AIInferenceSummaryV2ParamsFormatJson AIInferenceSummaryV2ParamsFormat = "JSON"
	AIInferenceSummaryV2ParamsFormatCsv  AIInferenceSummaryV2ParamsFormat = "CSV"
)

func (r AIInferenceSummaryV2ParamsFormat) IsKnown() bool {
	switch r {
	case AIInferenceSummaryV2ParamsFormatJson, AIInferenceSummaryV2ParamsFormatCsv:
		return true
	}
	return false
}

type AIInferenceSummaryV2ResponseEnvelope struct {
	Result  AIInferenceSummaryV2Response             `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    aiInferenceSummaryV2ResponseEnvelopeJSON `json:"-"`
}

// aiInferenceSummaryV2ResponseEnvelopeJSON contains the JSON metadata for the
// struct [AIInferenceSummaryV2ResponseEnvelope]
type aiInferenceSummaryV2ResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceSummaryV2ResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryV2ResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AIInferenceTimeseriesGroupsV2Params struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AIInferenceTimeseriesGroupsV2ParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[AIInferenceTimeseriesGroupsV2ParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AIInferenceTimeseriesGroupsV2ParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [AIInferenceTimeseriesGroupsV2Params]'s query parameters as
// `url.Values`.
func (r AIInferenceTimeseriesGroupsV2Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the attribute by which to group the results.
type AIInferenceTimeseriesGroupsV2ParamsDimension string

const (
	AIInferenceTimeseriesGroupsV2ParamsDimensionModel AIInferenceTimeseriesGroupsV2ParamsDimension = "MODEL"
	AIInferenceTimeseriesGroupsV2ParamsDimensionTask  AIInferenceTimeseriesGroupsV2ParamsDimension = "TASK"
)

func (r AIInferenceTimeseriesGroupsV2ParamsDimension) IsKnown() bool {
	switch r {
	case AIInferenceTimeseriesGroupsV2ParamsDimensionModel, AIInferenceTimeseriesGroupsV2ParamsDimensionTask:
		return true
	}
	return false
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AIInferenceTimeseriesGroupsV2ParamsAggInterval string

const (
	AIInferenceTimeseriesGroupsV2ParamsAggInterval15m AIInferenceTimeseriesGroupsV2ParamsAggInterval = "15m"
	AIInferenceTimeseriesGroupsV2ParamsAggInterval1h  AIInferenceTimeseriesGroupsV2ParamsAggInterval = "1h"
	AIInferenceTimeseriesGroupsV2ParamsAggInterval1d  AIInferenceTimeseriesGroupsV2ParamsAggInterval = "1d"
	AIInferenceTimeseriesGroupsV2ParamsAggInterval1w  AIInferenceTimeseriesGroupsV2ParamsAggInterval = "1w"
)

func (r AIInferenceTimeseriesGroupsV2ParamsAggInterval) IsKnown() bool {
	switch r {
	case AIInferenceTimeseriesGroupsV2ParamsAggInterval15m, AIInferenceTimeseriesGroupsV2ParamsAggInterval1h, AIInferenceTimeseriesGroupsV2ParamsAggInterval1d, AIInferenceTimeseriesGroupsV2ParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type AIInferenceTimeseriesGroupsV2ParamsFormat string

const (
	AIInferenceTimeseriesGroupsV2ParamsFormatJson AIInferenceTimeseriesGroupsV2ParamsFormat = "JSON"
	AIInferenceTimeseriesGroupsV2ParamsFormatCsv  AIInferenceTimeseriesGroupsV2ParamsFormat = "CSV"
)

func (r AIInferenceTimeseriesGroupsV2ParamsFormat) IsKnown() bool {
	switch r {
	case AIInferenceTimeseriesGroupsV2ParamsFormatJson, AIInferenceTimeseriesGroupsV2ParamsFormatCsv:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AIInferenceTimeseriesGroupsV2ParamsNormalization string

const (
	AIInferenceTimeseriesGroupsV2ParamsNormalizationPercentageChange AIInferenceTimeseriesGroupsV2ParamsNormalization = "PERCENTAGE_CHANGE"
	AIInferenceTimeseriesGroupsV2ParamsNormalizationMin0Max          AIInferenceTimeseriesGroupsV2ParamsNormalization = "MIN0_MAX"
)

func (r AIInferenceTimeseriesGroupsV2ParamsNormalization) IsKnown() bool {
	switch r {
	case AIInferenceTimeseriesGroupsV2ParamsNormalizationPercentageChange, AIInferenceTimeseriesGroupsV2ParamsNormalizationMin0Max:
		return true
	}
	return false
}

type AIInferenceTimeseriesGroupsV2ResponseEnvelope struct {
	Result  AIInferenceTimeseriesGroupsV2Response             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    aiInferenceTimeseriesGroupsV2ResponseEnvelopeJSON `json:"-"`
}

// aiInferenceTimeseriesGroupsV2ResponseEnvelopeJSON contains the JSON metadata for
// the struct [AIInferenceTimeseriesGroupsV2ResponseEnvelope]
type aiInferenceTimeseriesGroupsV2ResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceTimeseriesGroupsV2ResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceTimeseriesGroupsV2ResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
