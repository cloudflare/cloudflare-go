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

// AIMarkdownForAgentService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIMarkdownForAgentService] method instead.
type AIMarkdownForAgentService struct {
	Options []option.RequestOption
}

// NewAIMarkdownForAgentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIMarkdownForAgentService(opts ...option.RequestOption) (r *AIMarkdownForAgentService) {
	r = &AIMarkdownForAgentService{}
	r.Options = opts
	return
}

// Retrieves the overall median HTML-to-markdown reduction ratio for AI agent
// requests over the given date range.
func (r *AIMarkdownForAgentService) Summary(ctx context.Context, query AIMarkdownForAgentSummaryParams, opts ...option.RequestOption) (res *AIMarkdownForAgentSummaryResponse, err error) {
	var env AIMarkdownForAgentSummaryResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/ai/markdown_for_agents/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieves the median HTML-to-markdown reduction ratio over time for AI agent
// requests.
func (r *AIMarkdownForAgentService) Timeseries(ctx context.Context, query AIMarkdownForAgentTimeseriesParams, opts ...option.RequestOption) (res *AIMarkdownForAgentTimeseriesResponse, err error) {
	var env AIMarkdownForAgentTimeseriesResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/ai/markdown_for_agents/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type AIMarkdownForAgentSummaryResponse struct {
	// Metadata for the results.
	Meta     AIMarkdownForAgentSummaryResponseMeta     `json:"meta" api:"required"`
	Summary0 AIMarkdownForAgentSummaryResponseSummary0 `json:"summary_0" api:"required"`
	JSON     aiMarkdownForAgentSummaryResponseJSON     `json:"-"`
}

// aiMarkdownForAgentSummaryResponseJSON contains the JSON metadata for the struct
// [AIMarkdownForAgentSummaryResponse]
type aiMarkdownForAgentSummaryResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIMarkdownForAgentSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiMarkdownForAgentSummaryResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AIMarkdownForAgentSummaryResponseMeta struct {
	ConfidenceInfo AIMarkdownForAgentSummaryResponseMetaConfidenceInfo `json:"confidenceInfo" api:"required"`
	DateRange      []AIMarkdownForAgentSummaryResponseMetaDateRange    `json:"dateRange" api:"required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated" api:"required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AIMarkdownForAgentSummaryResponseMetaNormalization `json:"normalization" api:"required"`
	// Measurement units for the results.
	Units []AIMarkdownForAgentSummaryResponseMetaUnit `json:"units" api:"required"`
	JSON  aiMarkdownForAgentSummaryResponseMetaJSON   `json:"-"`
}

// aiMarkdownForAgentSummaryResponseMetaJSON contains the JSON metadata for the
// struct [AIMarkdownForAgentSummaryResponseMeta]
type aiMarkdownForAgentSummaryResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AIMarkdownForAgentSummaryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiMarkdownForAgentSummaryResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AIMarkdownForAgentSummaryResponseMetaConfidenceInfo struct {
	Annotations []AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotation `json:"annotations" api:"required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                   `json:"level" api:"required"`
	JSON  aiMarkdownForAgentSummaryResponseMetaConfidenceInfoJSON `json:"-"`
}

// aiMarkdownForAgentSummaryResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [AIMarkdownForAgentSummaryResponseMetaConfidenceInfo]
type aiMarkdownForAgentSummaryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIMarkdownForAgentSummaryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiMarkdownForAgentSummaryResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource" api:"required"`
	Description string                                                                   `json:"description" api:"required"`
	EndDate     time.Time                                                                `json:"endDate" api:"required" format:"date-time"`
	// Event type for annotations.
	EventType AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType" api:"required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                              `json:"isInstantaneous" api:"required"`
	LinkedURL       string                                                            `json:"linkedUrl" api:"required" format:"uri"`
	StartDate       time.Time                                                         `json:"startDate" api:"required" format:"date-time"`
	JSON            aiMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// aiMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotation]
type aiMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceAll                AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceBGP                AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceBots               AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceCT                 AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDNS                AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDos                AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceFw                 AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceIQI                AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceNet                AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceAll, AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceBGP, AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceBots, AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceCT, AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDNS, AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDos, AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceFw, AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceIQI, AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceNet, AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsEventType string

const (
	AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsEventTypeEvent             AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsEventTypeOutage            AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsEventTypePipeline          AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsEventTypeEvent, AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsEventTypeOutage, AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsEventTypePipeline, AIMarkdownForAgentSummaryResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type AIMarkdownForAgentSummaryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime" api:"required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                          `json:"startTime" api:"required" format:"date-time"`
	JSON      aiMarkdownForAgentSummaryResponseMetaDateRangeJSON `json:"-"`
}

// aiMarkdownForAgentSummaryResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [AIMarkdownForAgentSummaryResponseMetaDateRange]
type aiMarkdownForAgentSummaryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIMarkdownForAgentSummaryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiMarkdownForAgentSummaryResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AIMarkdownForAgentSummaryResponseMetaNormalization string

const (
	AIMarkdownForAgentSummaryResponseMetaNormalizationPercentage           AIMarkdownForAgentSummaryResponseMetaNormalization = "PERCENTAGE"
	AIMarkdownForAgentSummaryResponseMetaNormalizationMin0Max              AIMarkdownForAgentSummaryResponseMetaNormalization = "MIN0_MAX"
	AIMarkdownForAgentSummaryResponseMetaNormalizationMinMax               AIMarkdownForAgentSummaryResponseMetaNormalization = "MIN_MAX"
	AIMarkdownForAgentSummaryResponseMetaNormalizationRawValues            AIMarkdownForAgentSummaryResponseMetaNormalization = "RAW_VALUES"
	AIMarkdownForAgentSummaryResponseMetaNormalizationPercentageChange     AIMarkdownForAgentSummaryResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AIMarkdownForAgentSummaryResponseMetaNormalizationRollingAverage       AIMarkdownForAgentSummaryResponseMetaNormalization = "ROLLING_AVERAGE"
	AIMarkdownForAgentSummaryResponseMetaNormalizationOverlappedPercentage AIMarkdownForAgentSummaryResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AIMarkdownForAgentSummaryResponseMetaNormalizationRatio                AIMarkdownForAgentSummaryResponseMetaNormalization = "RATIO"
)

func (r AIMarkdownForAgentSummaryResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AIMarkdownForAgentSummaryResponseMetaNormalizationPercentage, AIMarkdownForAgentSummaryResponseMetaNormalizationMin0Max, AIMarkdownForAgentSummaryResponseMetaNormalizationMinMax, AIMarkdownForAgentSummaryResponseMetaNormalizationRawValues, AIMarkdownForAgentSummaryResponseMetaNormalizationPercentageChange, AIMarkdownForAgentSummaryResponseMetaNormalizationRollingAverage, AIMarkdownForAgentSummaryResponseMetaNormalizationOverlappedPercentage, AIMarkdownForAgentSummaryResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AIMarkdownForAgentSummaryResponseMetaUnit struct {
	Name  string                                        `json:"name" api:"required"`
	Value string                                        `json:"value" api:"required"`
	JSON  aiMarkdownForAgentSummaryResponseMetaUnitJSON `json:"-"`
}

// aiMarkdownForAgentSummaryResponseMetaUnitJSON contains the JSON metadata for the
// struct [AIMarkdownForAgentSummaryResponseMetaUnit]
type aiMarkdownForAgentSummaryResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIMarkdownForAgentSummaryResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiMarkdownForAgentSummaryResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AIMarkdownForAgentSummaryResponseSummary0 struct {
	// A numeric string that can include decimals and infinity values.
	Value string                                        `json:"value" api:"required"`
	JSON  aiMarkdownForAgentSummaryResponseSummary0JSON `json:"-"`
}

// aiMarkdownForAgentSummaryResponseSummary0JSON contains the JSON metadata for the
// struct [AIMarkdownForAgentSummaryResponseSummary0]
type aiMarkdownForAgentSummaryResponseSummary0JSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIMarkdownForAgentSummaryResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiMarkdownForAgentSummaryResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type AIMarkdownForAgentTimeseriesResponse struct {
	// Metadata for the results.
	Meta        AIMarkdownForAgentTimeseriesResponseMeta        `json:"meta" api:"required"`
	ExtraFields map[string]AIMarkdownForAgentTimeseriesResponse `json:"-" api:"extrafields"`
	JSON        aiMarkdownForAgentTimeseriesResponseJSON        `json:"-"`
}

// aiMarkdownForAgentTimeseriesResponseJSON contains the JSON metadata for the
// struct [AIMarkdownForAgentTimeseriesResponse]
type aiMarkdownForAgentTimeseriesResponseJSON struct {
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIMarkdownForAgentTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiMarkdownForAgentTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AIMarkdownForAgentTimeseriesResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AIMarkdownForAgentTimeseriesResponseMetaAggInterval    `json:"aggInterval" api:"required"`
	ConfidenceInfo AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfo `json:"confidenceInfo" api:"required"`
	DateRange      []AIMarkdownForAgentTimeseriesResponseMetaDateRange    `json:"dateRange" api:"required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated" api:"required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AIMarkdownForAgentTimeseriesResponseMetaNormalization `json:"normalization" api:"required"`
	// Measurement units for the results.
	Units []AIMarkdownForAgentTimeseriesResponseMetaUnit `json:"units" api:"required"`
	JSON  aiMarkdownForAgentTimeseriesResponseMetaJSON   `json:"-"`
}

// aiMarkdownForAgentTimeseriesResponseMetaJSON contains the JSON metadata for the
// struct [AIMarkdownForAgentTimeseriesResponseMeta]
type aiMarkdownForAgentTimeseriesResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AIMarkdownForAgentTimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiMarkdownForAgentTimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AIMarkdownForAgentTimeseriesResponseMetaAggInterval string

const (
	AIMarkdownForAgentTimeseriesResponseMetaAggIntervalFifteenMinutes AIMarkdownForAgentTimeseriesResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AIMarkdownForAgentTimeseriesResponseMetaAggIntervalOneHour        AIMarkdownForAgentTimeseriesResponseMetaAggInterval = "ONE_HOUR"
	AIMarkdownForAgentTimeseriesResponseMetaAggIntervalOneDay         AIMarkdownForAgentTimeseriesResponseMetaAggInterval = "ONE_DAY"
	AIMarkdownForAgentTimeseriesResponseMetaAggIntervalOneWeek        AIMarkdownForAgentTimeseriesResponseMetaAggInterval = "ONE_WEEK"
	AIMarkdownForAgentTimeseriesResponseMetaAggIntervalOneMonth       AIMarkdownForAgentTimeseriesResponseMetaAggInterval = "ONE_MONTH"
)

func (r AIMarkdownForAgentTimeseriesResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AIMarkdownForAgentTimeseriesResponseMetaAggIntervalFifteenMinutes, AIMarkdownForAgentTimeseriesResponseMetaAggIntervalOneHour, AIMarkdownForAgentTimeseriesResponseMetaAggIntervalOneDay, AIMarkdownForAgentTimeseriesResponseMetaAggIntervalOneWeek, AIMarkdownForAgentTimeseriesResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfo struct {
	Annotations []AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotation `json:"annotations" api:"required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                      `json:"level" api:"required"`
	JSON  aiMarkdownForAgentTimeseriesResponseMetaConfidenceInfoJSON `json:"-"`
}

// aiMarkdownForAgentTimeseriesResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfo]
type aiMarkdownForAgentTimeseriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiMarkdownForAgentTimeseriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource" api:"required"`
	Description string                                                                      `json:"description" api:"required"`
	EndDate     time.Time                                                                   `json:"endDate" api:"required" format:"date-time"`
	// Event type for annotations.
	EventType AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType" api:"required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                 `json:"isInstantaneous" api:"required"`
	LinkedURL       string                                                               `json:"linkedUrl" api:"required" format:"uri"`
	StartDate       time.Time                                                            `json:"startDate" api:"required" format:"date-time"`
	JSON            aiMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// aiMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotation]
type aiMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAll                AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBGP                AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBots               AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceCT                 AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNS                AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDos                AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFw                 AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceIQI                AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceNet                AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAll, AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBGP, AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBots, AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceCT, AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNS, AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDos, AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFw, AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceIQI, AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceNet, AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsEventType string

const (
	AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeEvent             AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeOutage            AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePipeline          AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeEvent, AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeOutage, AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePipeline, AIMarkdownForAgentTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type AIMarkdownForAgentTimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime" api:"required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime" api:"required" format:"date-time"`
	JSON      aiMarkdownForAgentTimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// aiMarkdownForAgentTimeseriesResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [AIMarkdownForAgentTimeseriesResponseMetaDateRange]
type aiMarkdownForAgentTimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIMarkdownForAgentTimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiMarkdownForAgentTimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AIMarkdownForAgentTimeseriesResponseMetaNormalization string

const (
	AIMarkdownForAgentTimeseriesResponseMetaNormalizationPercentage           AIMarkdownForAgentTimeseriesResponseMetaNormalization = "PERCENTAGE"
	AIMarkdownForAgentTimeseriesResponseMetaNormalizationMin0Max              AIMarkdownForAgentTimeseriesResponseMetaNormalization = "MIN0_MAX"
	AIMarkdownForAgentTimeseriesResponseMetaNormalizationMinMax               AIMarkdownForAgentTimeseriesResponseMetaNormalization = "MIN_MAX"
	AIMarkdownForAgentTimeseriesResponseMetaNormalizationRawValues            AIMarkdownForAgentTimeseriesResponseMetaNormalization = "RAW_VALUES"
	AIMarkdownForAgentTimeseriesResponseMetaNormalizationPercentageChange     AIMarkdownForAgentTimeseriesResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AIMarkdownForAgentTimeseriesResponseMetaNormalizationRollingAverage       AIMarkdownForAgentTimeseriesResponseMetaNormalization = "ROLLING_AVERAGE"
	AIMarkdownForAgentTimeseriesResponseMetaNormalizationOverlappedPercentage AIMarkdownForAgentTimeseriesResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AIMarkdownForAgentTimeseriesResponseMetaNormalizationRatio                AIMarkdownForAgentTimeseriesResponseMetaNormalization = "RATIO"
)

func (r AIMarkdownForAgentTimeseriesResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AIMarkdownForAgentTimeseriesResponseMetaNormalizationPercentage, AIMarkdownForAgentTimeseriesResponseMetaNormalizationMin0Max, AIMarkdownForAgentTimeseriesResponseMetaNormalizationMinMax, AIMarkdownForAgentTimeseriesResponseMetaNormalizationRawValues, AIMarkdownForAgentTimeseriesResponseMetaNormalizationPercentageChange, AIMarkdownForAgentTimeseriesResponseMetaNormalizationRollingAverage, AIMarkdownForAgentTimeseriesResponseMetaNormalizationOverlappedPercentage, AIMarkdownForAgentTimeseriesResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AIMarkdownForAgentTimeseriesResponseMetaUnit struct {
	Name  string                                           `json:"name" api:"required"`
	Value string                                           `json:"value" api:"required"`
	JSON  aiMarkdownForAgentTimeseriesResponseMetaUnitJSON `json:"-"`
}

// aiMarkdownForAgentTimeseriesResponseMetaUnitJSON contains the JSON metadata for
// the struct [AIMarkdownForAgentTimeseriesResponseMetaUnit]
type aiMarkdownForAgentTimeseriesResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIMarkdownForAgentTimeseriesResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiMarkdownForAgentTimeseriesResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AIMarkdownForAgentSummaryParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[AIMarkdownForAgentSummaryParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AIMarkdownForAgentSummaryParams]'s query parameters as
// `url.Values`.
func (r AIMarkdownForAgentSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AIMarkdownForAgentSummaryParamsFormat string

const (
	AIMarkdownForAgentSummaryParamsFormatJson AIMarkdownForAgentSummaryParamsFormat = "JSON"
	AIMarkdownForAgentSummaryParamsFormatCsv  AIMarkdownForAgentSummaryParamsFormat = "CSV"
)

func (r AIMarkdownForAgentSummaryParamsFormat) IsKnown() bool {
	switch r {
	case AIMarkdownForAgentSummaryParamsFormatJson, AIMarkdownForAgentSummaryParamsFormatCsv:
		return true
	}
	return false
}

type AIMarkdownForAgentSummaryResponseEnvelope struct {
	Result  AIMarkdownForAgentSummaryResponse             `json:"result" api:"required"`
	Success bool                                          `json:"success" api:"required"`
	JSON    aiMarkdownForAgentSummaryResponseEnvelopeJSON `json:"-"`
}

// aiMarkdownForAgentSummaryResponseEnvelopeJSON contains the JSON metadata for the
// struct [AIMarkdownForAgentSummaryResponseEnvelope]
type aiMarkdownForAgentSummaryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIMarkdownForAgentSummaryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiMarkdownForAgentSummaryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AIMarkdownForAgentTimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AIMarkdownForAgentTimeseriesParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[AIMarkdownForAgentTimeseriesParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AIMarkdownForAgentTimeseriesParams]'s query parameters as
// `url.Values`.
func (r AIMarkdownForAgentTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AIMarkdownForAgentTimeseriesParamsAggInterval string

const (
	AIMarkdownForAgentTimeseriesParamsAggInterval15m AIMarkdownForAgentTimeseriesParamsAggInterval = "15m"
	AIMarkdownForAgentTimeseriesParamsAggInterval1h  AIMarkdownForAgentTimeseriesParamsAggInterval = "1h"
	AIMarkdownForAgentTimeseriesParamsAggInterval1d  AIMarkdownForAgentTimeseriesParamsAggInterval = "1d"
	AIMarkdownForAgentTimeseriesParamsAggInterval1w  AIMarkdownForAgentTimeseriesParamsAggInterval = "1w"
)

func (r AIMarkdownForAgentTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case AIMarkdownForAgentTimeseriesParamsAggInterval15m, AIMarkdownForAgentTimeseriesParamsAggInterval1h, AIMarkdownForAgentTimeseriesParamsAggInterval1d, AIMarkdownForAgentTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type AIMarkdownForAgentTimeseriesParamsFormat string

const (
	AIMarkdownForAgentTimeseriesParamsFormatJson AIMarkdownForAgentTimeseriesParamsFormat = "JSON"
	AIMarkdownForAgentTimeseriesParamsFormatCsv  AIMarkdownForAgentTimeseriesParamsFormat = "CSV"
)

func (r AIMarkdownForAgentTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case AIMarkdownForAgentTimeseriesParamsFormatJson, AIMarkdownForAgentTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type AIMarkdownForAgentTimeseriesResponseEnvelope struct {
	Result  AIMarkdownForAgentTimeseriesResponse             `json:"result" api:"required"`
	Success bool                                             `json:"success" api:"required"`
	JSON    aiMarkdownForAgentTimeseriesResponseEnvelopeJSON `json:"-"`
}

// aiMarkdownForAgentTimeseriesResponseEnvelopeJSON contains the JSON metadata for
// the struct [AIMarkdownForAgentTimeseriesResponseEnvelope]
type aiMarkdownForAgentTimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIMarkdownForAgentTimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiMarkdownForAgentTimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
