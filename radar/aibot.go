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

// AIBotService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIBotService] method instead.
type AIBotService struct {
	Options []option.RequestOption
	Summary *AIBotSummaryService
}

// NewAIBotService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAIBotService(opts ...option.RequestOption) (r *AIBotService) {
	r = &AIBotService{}
	r.Options = opts
	r.Summary = NewAIBotSummaryService(opts...)
	return
}

// Retrieves an aggregated summary of AI bots HTTP requests grouped by the
// specified dimension.
func (r *AIBotService) SummaryV2(ctx context.Context, dimension AIBotSummaryV2ParamsDimension, query AIBotSummaryV2Params, opts ...option.RequestOption) (res *AIBotSummaryV2Response, err error) {
	var env AIBotSummaryV2ResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/ai/bots/summary/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves AI bots HTTP request volume over time.
func (r *AIBotService) Timeseries(ctx context.Context, query AIBotTimeseriesParams, opts ...option.RequestOption) (res *AIBotTimeseriesResponse, err error) {
	var env AIBotTimeseriesResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/ai/bots/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of HTTP requests from AI bots, grouped by the
// specified dimension over time.
func (r *AIBotService) TimeseriesGroups(ctx context.Context, dimension AIBotTimeseriesGroupsParamsDimension, query AIBotTimeseriesGroupsParams, opts ...option.RequestOption) (res *AIBotTimeseriesGroupsResponse, err error) {
	var env AIBotTimeseriesGroupsResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/ai/bots/timeseries_groups/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AIBotSummaryV2Response struct {
	// Metadata for the results.
	Meta     AIBotSummaryV2ResponseMeta `json:"meta,required"`
	Summary0 map[string]string          `json:"summary_0,required"`
	JSON     aiBotSummaryV2ResponseJSON `json:"-"`
}

// aiBotSummaryV2ResponseJSON contains the JSON metadata for the struct
// [AIBotSummaryV2Response]
type aiBotSummaryV2ResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotSummaryV2Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryV2ResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AIBotSummaryV2ResponseMeta struct {
	ConfidenceInfo AIBotSummaryV2ResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AIBotSummaryV2ResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AIBotSummaryV2ResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AIBotSummaryV2ResponseMetaUnit `json:"units,required"`
	JSON  aiBotSummaryV2ResponseMetaJSON   `json:"-"`
}

// aiBotSummaryV2ResponseMetaJSON contains the JSON metadata for the struct
// [AIBotSummaryV2ResponseMeta]
type aiBotSummaryV2ResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AIBotSummaryV2ResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryV2ResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AIBotSummaryV2ResponseMetaConfidenceInfo struct {
	Annotations []AIBotSummaryV2ResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                        `json:"level,required"`
	JSON  aiBotSummaryV2ResponseMetaConfidenceInfoJSON `json:"-"`
}

// aiBotSummaryV2ResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [AIBotSummaryV2ResponseMetaConfidenceInfo]
type aiBotSummaryV2ResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotSummaryV2ResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryV2ResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AIBotSummaryV2ResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                        `json:"description,required"`
	EndDate     time.Time                                                     `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                   `json:"isInstantaneous,required"`
	LinkedURL       string                                                 `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                              `json:"startDate,required" format:"date-time"`
	JSON            aiBotSummaryV2ResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// aiBotSummaryV2ResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [AIBotSummaryV2ResponseMetaConfidenceInfoAnnotation]
type aiBotSummaryV2ResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AIBotSummaryV2ResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryV2ResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAll                AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBGP                AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBots               AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceCT                 AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNS                AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDos                AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFw                 AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceIQI                AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceNet                AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAll, AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBGP, AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBots, AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceCT, AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNS, AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDos, AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFw, AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceIQI, AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceNet, AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType string

const (
	AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeEvent             AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeOutage            AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePipeline          AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeEvent, AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeOutage, AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePipeline, AIBotSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type AIBotSummaryV2ResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                               `json:"startTime,required" format:"date-time"`
	JSON      aiBotSummaryV2ResponseMetaDateRangeJSON `json:"-"`
}

// aiBotSummaryV2ResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [AIBotSummaryV2ResponseMetaDateRange]
type aiBotSummaryV2ResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotSummaryV2ResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryV2ResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AIBotSummaryV2ResponseMetaNormalization string

const (
	AIBotSummaryV2ResponseMetaNormalizationPercentage           AIBotSummaryV2ResponseMetaNormalization = "PERCENTAGE"
	AIBotSummaryV2ResponseMetaNormalizationMin0Max              AIBotSummaryV2ResponseMetaNormalization = "MIN0_MAX"
	AIBotSummaryV2ResponseMetaNormalizationMinMax               AIBotSummaryV2ResponseMetaNormalization = "MIN_MAX"
	AIBotSummaryV2ResponseMetaNormalizationRawValues            AIBotSummaryV2ResponseMetaNormalization = "RAW_VALUES"
	AIBotSummaryV2ResponseMetaNormalizationPercentageChange     AIBotSummaryV2ResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AIBotSummaryV2ResponseMetaNormalizationRollingAverage       AIBotSummaryV2ResponseMetaNormalization = "ROLLING_AVERAGE"
	AIBotSummaryV2ResponseMetaNormalizationOverlappedPercentage AIBotSummaryV2ResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AIBotSummaryV2ResponseMetaNormalizationRatio                AIBotSummaryV2ResponseMetaNormalization = "RATIO"
)

func (r AIBotSummaryV2ResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AIBotSummaryV2ResponseMetaNormalizationPercentage, AIBotSummaryV2ResponseMetaNormalizationMin0Max, AIBotSummaryV2ResponseMetaNormalizationMinMax, AIBotSummaryV2ResponseMetaNormalizationRawValues, AIBotSummaryV2ResponseMetaNormalizationPercentageChange, AIBotSummaryV2ResponseMetaNormalizationRollingAverage, AIBotSummaryV2ResponseMetaNormalizationOverlappedPercentage, AIBotSummaryV2ResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AIBotSummaryV2ResponseMetaUnit struct {
	Name  string                             `json:"name,required"`
	Value string                             `json:"value,required"`
	JSON  aiBotSummaryV2ResponseMetaUnitJSON `json:"-"`
}

// aiBotSummaryV2ResponseMetaUnitJSON contains the JSON metadata for the struct
// [AIBotSummaryV2ResponseMetaUnit]
type aiBotSummaryV2ResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotSummaryV2ResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryV2ResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AIBotTimeseriesResponse struct {
	// Metadata for the results.
	Meta        AIBotTimeseriesResponseMeta        `json:"meta,required"`
	ExtraFields map[string]AIBotTimeseriesResponse `json:"-,extras"`
	JSON        aiBotTimeseriesResponseJSON        `json:"-"`
}

// aiBotTimeseriesResponseJSON contains the JSON metadata for the struct
// [AIBotTimeseriesResponse]
type aiBotTimeseriesResponseJSON struct {
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AIBotTimeseriesResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AIBotTimeseriesResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AIBotTimeseriesResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AIBotTimeseriesResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AIBotTimeseriesResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AIBotTimeseriesResponseMetaUnit `json:"units,required"`
	JSON  aiBotTimeseriesResponseMetaJSON   `json:"-"`
}

// aiBotTimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [AIBotTimeseriesResponseMeta]
type aiBotTimeseriesResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AIBotTimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotTimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AIBotTimeseriesResponseMetaAggInterval string

const (
	AIBotTimeseriesResponseMetaAggIntervalFifteenMinutes AIBotTimeseriesResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AIBotTimeseriesResponseMetaAggIntervalOneHour        AIBotTimeseriesResponseMetaAggInterval = "ONE_HOUR"
	AIBotTimeseriesResponseMetaAggIntervalOneDay         AIBotTimeseriesResponseMetaAggInterval = "ONE_DAY"
	AIBotTimeseriesResponseMetaAggIntervalOneWeek        AIBotTimeseriesResponseMetaAggInterval = "ONE_WEEK"
	AIBotTimeseriesResponseMetaAggIntervalOneMonth       AIBotTimeseriesResponseMetaAggInterval = "ONE_MONTH"
)

func (r AIBotTimeseriesResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AIBotTimeseriesResponseMetaAggIntervalFifteenMinutes, AIBotTimeseriesResponseMetaAggIntervalOneHour, AIBotTimeseriesResponseMetaAggIntervalOneDay, AIBotTimeseriesResponseMetaAggIntervalOneWeek, AIBotTimeseriesResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AIBotTimeseriesResponseMetaConfidenceInfo struct {
	Annotations []AIBotTimeseriesResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                         `json:"level,required"`
	JSON  aiBotTimeseriesResponseMetaConfidenceInfoJSON `json:"-"`
}

// aiBotTimeseriesResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [AIBotTimeseriesResponseMetaConfidenceInfo]
type aiBotTimeseriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotTimeseriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotTimeseriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AIBotTimeseriesResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                         `json:"description,required"`
	EndDate     time.Time                                                      `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                    `json:"isInstantaneous,required"`
	LinkedURL       string                                                  `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                               `json:"startDate,required" format:"date-time"`
	JSON            aiBotTimeseriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// aiBotTimeseriesResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [AIBotTimeseriesResponseMetaConfidenceInfoAnnotation]
type aiBotTimeseriesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AIBotTimeseriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotTimeseriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAll                AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBGP                AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBots               AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceCT                 AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNS                AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDos                AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFw                 AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceIQI                AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceNet                AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAll, AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBGP, AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBots, AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceCT, AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNS, AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDos, AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFw, AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceIQI, AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceNet, AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsEventType string

const (
	AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeEvent             AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeOutage            AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePipeline          AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeEvent, AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeOutage, AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePipeline, AIBotTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type AIBotTimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                `json:"startTime,required" format:"date-time"`
	JSON      aiBotTimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// aiBotTimeseriesResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [AIBotTimeseriesResponseMetaDateRange]
type aiBotTimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotTimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotTimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AIBotTimeseriesResponseMetaNormalization string

const (
	AIBotTimeseriesResponseMetaNormalizationPercentage           AIBotTimeseriesResponseMetaNormalization = "PERCENTAGE"
	AIBotTimeseriesResponseMetaNormalizationMin0Max              AIBotTimeseriesResponseMetaNormalization = "MIN0_MAX"
	AIBotTimeseriesResponseMetaNormalizationMinMax               AIBotTimeseriesResponseMetaNormalization = "MIN_MAX"
	AIBotTimeseriesResponseMetaNormalizationRawValues            AIBotTimeseriesResponseMetaNormalization = "RAW_VALUES"
	AIBotTimeseriesResponseMetaNormalizationPercentageChange     AIBotTimeseriesResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AIBotTimeseriesResponseMetaNormalizationRollingAverage       AIBotTimeseriesResponseMetaNormalization = "ROLLING_AVERAGE"
	AIBotTimeseriesResponseMetaNormalizationOverlappedPercentage AIBotTimeseriesResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AIBotTimeseriesResponseMetaNormalizationRatio                AIBotTimeseriesResponseMetaNormalization = "RATIO"
)

func (r AIBotTimeseriesResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AIBotTimeseriesResponseMetaNormalizationPercentage, AIBotTimeseriesResponseMetaNormalizationMin0Max, AIBotTimeseriesResponseMetaNormalizationMinMax, AIBotTimeseriesResponseMetaNormalizationRawValues, AIBotTimeseriesResponseMetaNormalizationPercentageChange, AIBotTimeseriesResponseMetaNormalizationRollingAverage, AIBotTimeseriesResponseMetaNormalizationOverlappedPercentage, AIBotTimeseriesResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AIBotTimeseriesResponseMetaUnit struct {
	Name  string                              `json:"name,required"`
	Value string                              `json:"value,required"`
	JSON  aiBotTimeseriesResponseMetaUnitJSON `json:"-"`
}

// aiBotTimeseriesResponseMetaUnitJSON contains the JSON metadata for the struct
// [AIBotTimeseriesResponseMetaUnit]
type aiBotTimeseriesResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotTimeseriesResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotTimeseriesResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AIBotTimeseriesGroupsResponse struct {
	// Metadata for the results.
	Meta   AIBotTimeseriesGroupsResponseMeta   `json:"meta,required"`
	Serie0 AIBotTimeseriesGroupsResponseSerie0 `json:"serie_0,required"`
	JSON   aiBotTimeseriesGroupsResponseJSON   `json:"-"`
}

// aiBotTimeseriesGroupsResponseJSON contains the JSON metadata for the struct
// [AIBotTimeseriesGroupsResponse]
type aiBotTimeseriesGroupsResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotTimeseriesGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotTimeseriesGroupsResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AIBotTimeseriesGroupsResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AIBotTimeseriesGroupsResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AIBotTimeseriesGroupsResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AIBotTimeseriesGroupsResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AIBotTimeseriesGroupsResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AIBotTimeseriesGroupsResponseMetaUnit `json:"units,required"`
	JSON  aiBotTimeseriesGroupsResponseMetaJSON   `json:"-"`
}

// aiBotTimeseriesGroupsResponseMetaJSON contains the JSON metadata for the struct
// [AIBotTimeseriesGroupsResponseMeta]
type aiBotTimeseriesGroupsResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AIBotTimeseriesGroupsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotTimeseriesGroupsResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AIBotTimeseriesGroupsResponseMetaAggInterval string

const (
	AIBotTimeseriesGroupsResponseMetaAggIntervalFifteenMinutes AIBotTimeseriesGroupsResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AIBotTimeseriesGroupsResponseMetaAggIntervalOneHour        AIBotTimeseriesGroupsResponseMetaAggInterval = "ONE_HOUR"
	AIBotTimeseriesGroupsResponseMetaAggIntervalOneDay         AIBotTimeseriesGroupsResponseMetaAggInterval = "ONE_DAY"
	AIBotTimeseriesGroupsResponseMetaAggIntervalOneWeek        AIBotTimeseriesGroupsResponseMetaAggInterval = "ONE_WEEK"
	AIBotTimeseriesGroupsResponseMetaAggIntervalOneMonth       AIBotTimeseriesGroupsResponseMetaAggInterval = "ONE_MONTH"
)

func (r AIBotTimeseriesGroupsResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AIBotTimeseriesGroupsResponseMetaAggIntervalFifteenMinutes, AIBotTimeseriesGroupsResponseMetaAggIntervalOneHour, AIBotTimeseriesGroupsResponseMetaAggIntervalOneDay, AIBotTimeseriesGroupsResponseMetaAggIntervalOneWeek, AIBotTimeseriesGroupsResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AIBotTimeseriesGroupsResponseMetaConfidenceInfo struct {
	Annotations []AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                               `json:"level,required"`
	JSON  aiBotTimeseriesGroupsResponseMetaConfidenceInfoJSON `json:"-"`
}

// aiBotTimeseriesGroupsResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AIBotTimeseriesGroupsResponseMetaConfidenceInfo]
type aiBotTimeseriesGroupsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotTimeseriesGroupsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotTimeseriesGroupsResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                               `json:"description,required"`
	EndDate     time.Time                                                            `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                          `json:"isInstantaneous,required"`
	LinkedURL       string                                                        `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                     `json:"startDate,required" format:"date-time"`
	JSON            aiBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// aiBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotation]
type aiBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceAll                AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceBGP                AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceBots               AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceCT                 AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDNS                AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDos                AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceFw                 AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceIQI                AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceNet                AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceAll, AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceBGP, AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceBots, AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceCT, AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDNS, AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDos, AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceFw, AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceIQI, AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceNet, AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType string

const (
	AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeEvent             AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeOutage            AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypePipeline          AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeEvent, AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeOutage, AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypePipeline, AIBotTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type AIBotTimeseriesGroupsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                      `json:"startTime,required" format:"date-time"`
	JSON      aiBotTimeseriesGroupsResponseMetaDateRangeJSON `json:"-"`
}

// aiBotTimeseriesGroupsResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AIBotTimeseriesGroupsResponseMetaDateRange]
type aiBotTimeseriesGroupsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotTimeseriesGroupsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotTimeseriesGroupsResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AIBotTimeseriesGroupsResponseMetaNormalization string

const (
	AIBotTimeseriesGroupsResponseMetaNormalizationPercentage           AIBotTimeseriesGroupsResponseMetaNormalization = "PERCENTAGE"
	AIBotTimeseriesGroupsResponseMetaNormalizationMin0Max              AIBotTimeseriesGroupsResponseMetaNormalization = "MIN0_MAX"
	AIBotTimeseriesGroupsResponseMetaNormalizationMinMax               AIBotTimeseriesGroupsResponseMetaNormalization = "MIN_MAX"
	AIBotTimeseriesGroupsResponseMetaNormalizationRawValues            AIBotTimeseriesGroupsResponseMetaNormalization = "RAW_VALUES"
	AIBotTimeseriesGroupsResponseMetaNormalizationPercentageChange     AIBotTimeseriesGroupsResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AIBotTimeseriesGroupsResponseMetaNormalizationRollingAverage       AIBotTimeseriesGroupsResponseMetaNormalization = "ROLLING_AVERAGE"
	AIBotTimeseriesGroupsResponseMetaNormalizationOverlappedPercentage AIBotTimeseriesGroupsResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AIBotTimeseriesGroupsResponseMetaNormalizationRatio                AIBotTimeseriesGroupsResponseMetaNormalization = "RATIO"
)

func (r AIBotTimeseriesGroupsResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AIBotTimeseriesGroupsResponseMetaNormalizationPercentage, AIBotTimeseriesGroupsResponseMetaNormalizationMin0Max, AIBotTimeseriesGroupsResponseMetaNormalizationMinMax, AIBotTimeseriesGroupsResponseMetaNormalizationRawValues, AIBotTimeseriesGroupsResponseMetaNormalizationPercentageChange, AIBotTimeseriesGroupsResponseMetaNormalizationRollingAverage, AIBotTimeseriesGroupsResponseMetaNormalizationOverlappedPercentage, AIBotTimeseriesGroupsResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AIBotTimeseriesGroupsResponseMetaUnit struct {
	Name  string                                    `json:"name,required"`
	Value string                                    `json:"value,required"`
	JSON  aiBotTimeseriesGroupsResponseMetaUnitJSON `json:"-"`
}

// aiBotTimeseriesGroupsResponseMetaUnitJSON contains the JSON metadata for the
// struct [AIBotTimeseriesGroupsResponseMetaUnit]
type aiBotTimeseriesGroupsResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotTimeseriesGroupsResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotTimeseriesGroupsResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AIBotTimeseriesGroupsResponseSerie0 struct {
	Timestamps  []time.Time                             `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                     `json:"-,extras"`
	JSON        aiBotTimeseriesGroupsResponseSerie0JSON `json:"-"`
}

// aiBotTimeseriesGroupsResponseSerie0JSON contains the JSON metadata for the
// struct [AIBotTimeseriesGroupsResponseSerie0]
type aiBotTimeseriesGroupsResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotTimeseriesGroupsResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotTimeseriesGroupsResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AIBotSummaryV2Params struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by content type category.
	ContentType param.Field[[]AIBotSummaryV2ParamsContentType] `query:"contentType"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// Filters results by bot crawl purpose.
	CrawlPurpose param.Field[[]string] `query:"crawlPurpose"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[AIBotSummaryV2ParamsFormat] `query:"format"`
	// Filters results by industry.
	Industry param.Field[[]string] `query:"industry"`
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
	// Filters results by user agent.
	UserAgent param.Field[[]string] `query:"userAgent"`
	// Filters results by vertical.
	Vertical param.Field[[]string] `query:"vertical"`
}

// URLQuery serializes [AIBotSummaryV2Params]'s query parameters as `url.Values`.
func (r AIBotSummaryV2Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the attribute by which to group the results.
type AIBotSummaryV2ParamsDimension string

const (
	AIBotSummaryV2ParamsDimensionUserAgent    AIBotSummaryV2ParamsDimension = "USER_AGENT"
	AIBotSummaryV2ParamsDimensionCrawlPurpose AIBotSummaryV2ParamsDimension = "CRAWL_PURPOSE"
	AIBotSummaryV2ParamsDimensionIndustry     AIBotSummaryV2ParamsDimension = "INDUSTRY"
	AIBotSummaryV2ParamsDimensionVertical     AIBotSummaryV2ParamsDimension = "VERTICAL"
	AIBotSummaryV2ParamsDimensionContentType  AIBotSummaryV2ParamsDimension = "CONTENT_TYPE"
)

func (r AIBotSummaryV2ParamsDimension) IsKnown() bool {
	switch r {
	case AIBotSummaryV2ParamsDimensionUserAgent, AIBotSummaryV2ParamsDimensionCrawlPurpose, AIBotSummaryV2ParamsDimensionIndustry, AIBotSummaryV2ParamsDimensionVertical, AIBotSummaryV2ParamsDimensionContentType:
		return true
	}
	return false
}

type AIBotSummaryV2ParamsContentType string

const (
	AIBotSummaryV2ParamsContentTypeHTML          AIBotSummaryV2ParamsContentType = "HTML"
	AIBotSummaryV2ParamsContentTypeImages        AIBotSummaryV2ParamsContentType = "IMAGES"
	AIBotSummaryV2ParamsContentTypeJson          AIBotSummaryV2ParamsContentType = "JSON"
	AIBotSummaryV2ParamsContentTypeJavascript    AIBotSummaryV2ParamsContentType = "JAVASCRIPT"
	AIBotSummaryV2ParamsContentTypeCSS           AIBotSummaryV2ParamsContentType = "CSS"
	AIBotSummaryV2ParamsContentTypePlainText     AIBotSummaryV2ParamsContentType = "PLAIN_TEXT"
	AIBotSummaryV2ParamsContentTypeFonts         AIBotSummaryV2ParamsContentType = "FONTS"
	AIBotSummaryV2ParamsContentTypeXml           AIBotSummaryV2ParamsContentType = "XML"
	AIBotSummaryV2ParamsContentTypeYaml          AIBotSummaryV2ParamsContentType = "YAML"
	AIBotSummaryV2ParamsContentTypeVideo         AIBotSummaryV2ParamsContentType = "VIDEO"
	AIBotSummaryV2ParamsContentTypeAudio         AIBotSummaryV2ParamsContentType = "AUDIO"
	AIBotSummaryV2ParamsContentTypeMarkdown      AIBotSummaryV2ParamsContentType = "MARKDOWN"
	AIBotSummaryV2ParamsContentTypeDocuments     AIBotSummaryV2ParamsContentType = "DOCUMENTS"
	AIBotSummaryV2ParamsContentTypeBinary        AIBotSummaryV2ParamsContentType = "BINARY"
	AIBotSummaryV2ParamsContentTypeSerialization AIBotSummaryV2ParamsContentType = "SERIALIZATION"
	AIBotSummaryV2ParamsContentTypeOther         AIBotSummaryV2ParamsContentType = "OTHER"
)

func (r AIBotSummaryV2ParamsContentType) IsKnown() bool {
	switch r {
	case AIBotSummaryV2ParamsContentTypeHTML, AIBotSummaryV2ParamsContentTypeImages, AIBotSummaryV2ParamsContentTypeJson, AIBotSummaryV2ParamsContentTypeJavascript, AIBotSummaryV2ParamsContentTypeCSS, AIBotSummaryV2ParamsContentTypePlainText, AIBotSummaryV2ParamsContentTypeFonts, AIBotSummaryV2ParamsContentTypeXml, AIBotSummaryV2ParamsContentTypeYaml, AIBotSummaryV2ParamsContentTypeVideo, AIBotSummaryV2ParamsContentTypeAudio, AIBotSummaryV2ParamsContentTypeMarkdown, AIBotSummaryV2ParamsContentTypeDocuments, AIBotSummaryV2ParamsContentTypeBinary, AIBotSummaryV2ParamsContentTypeSerialization, AIBotSummaryV2ParamsContentTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type AIBotSummaryV2ParamsFormat string

const (
	AIBotSummaryV2ParamsFormatJson AIBotSummaryV2ParamsFormat = "JSON"
	AIBotSummaryV2ParamsFormatCsv  AIBotSummaryV2ParamsFormat = "CSV"
)

func (r AIBotSummaryV2ParamsFormat) IsKnown() bool {
	switch r {
	case AIBotSummaryV2ParamsFormatJson, AIBotSummaryV2ParamsFormatCsv:
		return true
	}
	return false
}

type AIBotSummaryV2ResponseEnvelope struct {
	Result  AIBotSummaryV2Response             `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    aiBotSummaryV2ResponseEnvelopeJSON `json:"-"`
}

// aiBotSummaryV2ResponseEnvelopeJSON contains the JSON metadata for the struct
// [AIBotSummaryV2ResponseEnvelope]
type aiBotSummaryV2ResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotSummaryV2ResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryV2ResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AIBotTimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AIBotTimeseriesParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by content type category.
	ContentType param.Field[[]AIBotTimeseriesParamsContentType] `query:"contentType"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// Filters results by bot crawl purpose.
	CrawlPurpose param.Field[[]string] `query:"crawlPurpose"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[AIBotTimeseriesParamsFormat] `query:"format"`
	// Filters results by industry.
	Industry param.Field[[]string] `query:"industry"`
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
	// Filters results by user agent.
	UserAgent param.Field[[]string] `query:"userAgent"`
	// Filters results by vertical.
	Vertical param.Field[[]string] `query:"vertical"`
}

// URLQuery serializes [AIBotTimeseriesParams]'s query parameters as `url.Values`.
func (r AIBotTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AIBotTimeseriesParamsAggInterval string

const (
	AIBotTimeseriesParamsAggInterval15m AIBotTimeseriesParamsAggInterval = "15m"
	AIBotTimeseriesParamsAggInterval1h  AIBotTimeseriesParamsAggInterval = "1h"
	AIBotTimeseriesParamsAggInterval1d  AIBotTimeseriesParamsAggInterval = "1d"
	AIBotTimeseriesParamsAggInterval1w  AIBotTimeseriesParamsAggInterval = "1w"
)

func (r AIBotTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case AIBotTimeseriesParamsAggInterval15m, AIBotTimeseriesParamsAggInterval1h, AIBotTimeseriesParamsAggInterval1d, AIBotTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

type AIBotTimeseriesParamsContentType string

const (
	AIBotTimeseriesParamsContentTypeHTML          AIBotTimeseriesParamsContentType = "HTML"
	AIBotTimeseriesParamsContentTypeImages        AIBotTimeseriesParamsContentType = "IMAGES"
	AIBotTimeseriesParamsContentTypeJson          AIBotTimeseriesParamsContentType = "JSON"
	AIBotTimeseriesParamsContentTypeJavascript    AIBotTimeseriesParamsContentType = "JAVASCRIPT"
	AIBotTimeseriesParamsContentTypeCSS           AIBotTimeseriesParamsContentType = "CSS"
	AIBotTimeseriesParamsContentTypePlainText     AIBotTimeseriesParamsContentType = "PLAIN_TEXT"
	AIBotTimeseriesParamsContentTypeFonts         AIBotTimeseriesParamsContentType = "FONTS"
	AIBotTimeseriesParamsContentTypeXml           AIBotTimeseriesParamsContentType = "XML"
	AIBotTimeseriesParamsContentTypeYaml          AIBotTimeseriesParamsContentType = "YAML"
	AIBotTimeseriesParamsContentTypeVideo         AIBotTimeseriesParamsContentType = "VIDEO"
	AIBotTimeseriesParamsContentTypeAudio         AIBotTimeseriesParamsContentType = "AUDIO"
	AIBotTimeseriesParamsContentTypeMarkdown      AIBotTimeseriesParamsContentType = "MARKDOWN"
	AIBotTimeseriesParamsContentTypeDocuments     AIBotTimeseriesParamsContentType = "DOCUMENTS"
	AIBotTimeseriesParamsContentTypeBinary        AIBotTimeseriesParamsContentType = "BINARY"
	AIBotTimeseriesParamsContentTypeSerialization AIBotTimeseriesParamsContentType = "SERIALIZATION"
	AIBotTimeseriesParamsContentTypeOther         AIBotTimeseriesParamsContentType = "OTHER"
)

func (r AIBotTimeseriesParamsContentType) IsKnown() bool {
	switch r {
	case AIBotTimeseriesParamsContentTypeHTML, AIBotTimeseriesParamsContentTypeImages, AIBotTimeseriesParamsContentTypeJson, AIBotTimeseriesParamsContentTypeJavascript, AIBotTimeseriesParamsContentTypeCSS, AIBotTimeseriesParamsContentTypePlainText, AIBotTimeseriesParamsContentTypeFonts, AIBotTimeseriesParamsContentTypeXml, AIBotTimeseriesParamsContentTypeYaml, AIBotTimeseriesParamsContentTypeVideo, AIBotTimeseriesParamsContentTypeAudio, AIBotTimeseriesParamsContentTypeMarkdown, AIBotTimeseriesParamsContentTypeDocuments, AIBotTimeseriesParamsContentTypeBinary, AIBotTimeseriesParamsContentTypeSerialization, AIBotTimeseriesParamsContentTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type AIBotTimeseriesParamsFormat string

const (
	AIBotTimeseriesParamsFormatJson AIBotTimeseriesParamsFormat = "JSON"
	AIBotTimeseriesParamsFormatCsv  AIBotTimeseriesParamsFormat = "CSV"
)

func (r AIBotTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case AIBotTimeseriesParamsFormatJson, AIBotTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type AIBotTimeseriesResponseEnvelope struct {
	Result  AIBotTimeseriesResponse             `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    aiBotTimeseriesResponseEnvelopeJSON `json:"-"`
}

// aiBotTimeseriesResponseEnvelopeJSON contains the JSON metadata for the struct
// [AIBotTimeseriesResponseEnvelope]
type aiBotTimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotTimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotTimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AIBotTimeseriesGroupsParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AIBotTimeseriesGroupsParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by content type category.
	ContentType param.Field[[]AIBotTimeseriesGroupsParamsContentType] `query:"contentType"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// Filters results by bot crawl purpose.
	CrawlPurpose param.Field[[]string] `query:"crawlPurpose"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[AIBotTimeseriesGroupsParamsFormat] `query:"format"`
	// Filters results by industry.
	Industry param.Field[[]string] `query:"industry"`
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
	Normalization param.Field[AIBotTimeseriesGroupsParamsNormalization] `query:"normalization"`
	// Filters results by user agent.
	UserAgent param.Field[[]string] `query:"userAgent"`
	// Filters results by vertical.
	Vertical param.Field[[]string] `query:"vertical"`
}

// URLQuery serializes [AIBotTimeseriesGroupsParams]'s query parameters as
// `url.Values`.
func (r AIBotTimeseriesGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the attribute by which to group the results.
type AIBotTimeseriesGroupsParamsDimension string

const (
	AIBotTimeseriesGroupsParamsDimensionUserAgent    AIBotTimeseriesGroupsParamsDimension = "USER_AGENT"
	AIBotTimeseriesGroupsParamsDimensionCrawlPurpose AIBotTimeseriesGroupsParamsDimension = "CRAWL_PURPOSE"
	AIBotTimeseriesGroupsParamsDimensionIndustry     AIBotTimeseriesGroupsParamsDimension = "INDUSTRY"
	AIBotTimeseriesGroupsParamsDimensionVertical     AIBotTimeseriesGroupsParamsDimension = "VERTICAL"
	AIBotTimeseriesGroupsParamsDimensionContentType  AIBotTimeseriesGroupsParamsDimension = "CONTENT_TYPE"
)

func (r AIBotTimeseriesGroupsParamsDimension) IsKnown() bool {
	switch r {
	case AIBotTimeseriesGroupsParamsDimensionUserAgent, AIBotTimeseriesGroupsParamsDimensionCrawlPurpose, AIBotTimeseriesGroupsParamsDimensionIndustry, AIBotTimeseriesGroupsParamsDimensionVertical, AIBotTimeseriesGroupsParamsDimensionContentType:
		return true
	}
	return false
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AIBotTimeseriesGroupsParamsAggInterval string

const (
	AIBotTimeseriesGroupsParamsAggInterval15m AIBotTimeseriesGroupsParamsAggInterval = "15m"
	AIBotTimeseriesGroupsParamsAggInterval1h  AIBotTimeseriesGroupsParamsAggInterval = "1h"
	AIBotTimeseriesGroupsParamsAggInterval1d  AIBotTimeseriesGroupsParamsAggInterval = "1d"
	AIBotTimeseriesGroupsParamsAggInterval1w  AIBotTimeseriesGroupsParamsAggInterval = "1w"
)

func (r AIBotTimeseriesGroupsParamsAggInterval) IsKnown() bool {
	switch r {
	case AIBotTimeseriesGroupsParamsAggInterval15m, AIBotTimeseriesGroupsParamsAggInterval1h, AIBotTimeseriesGroupsParamsAggInterval1d, AIBotTimeseriesGroupsParamsAggInterval1w:
		return true
	}
	return false
}

type AIBotTimeseriesGroupsParamsContentType string

const (
	AIBotTimeseriesGroupsParamsContentTypeHTML          AIBotTimeseriesGroupsParamsContentType = "HTML"
	AIBotTimeseriesGroupsParamsContentTypeImages        AIBotTimeseriesGroupsParamsContentType = "IMAGES"
	AIBotTimeseriesGroupsParamsContentTypeJson          AIBotTimeseriesGroupsParamsContentType = "JSON"
	AIBotTimeseriesGroupsParamsContentTypeJavascript    AIBotTimeseriesGroupsParamsContentType = "JAVASCRIPT"
	AIBotTimeseriesGroupsParamsContentTypeCSS           AIBotTimeseriesGroupsParamsContentType = "CSS"
	AIBotTimeseriesGroupsParamsContentTypePlainText     AIBotTimeseriesGroupsParamsContentType = "PLAIN_TEXT"
	AIBotTimeseriesGroupsParamsContentTypeFonts         AIBotTimeseriesGroupsParamsContentType = "FONTS"
	AIBotTimeseriesGroupsParamsContentTypeXml           AIBotTimeseriesGroupsParamsContentType = "XML"
	AIBotTimeseriesGroupsParamsContentTypeYaml          AIBotTimeseriesGroupsParamsContentType = "YAML"
	AIBotTimeseriesGroupsParamsContentTypeVideo         AIBotTimeseriesGroupsParamsContentType = "VIDEO"
	AIBotTimeseriesGroupsParamsContentTypeAudio         AIBotTimeseriesGroupsParamsContentType = "AUDIO"
	AIBotTimeseriesGroupsParamsContentTypeMarkdown      AIBotTimeseriesGroupsParamsContentType = "MARKDOWN"
	AIBotTimeseriesGroupsParamsContentTypeDocuments     AIBotTimeseriesGroupsParamsContentType = "DOCUMENTS"
	AIBotTimeseriesGroupsParamsContentTypeBinary        AIBotTimeseriesGroupsParamsContentType = "BINARY"
	AIBotTimeseriesGroupsParamsContentTypeSerialization AIBotTimeseriesGroupsParamsContentType = "SERIALIZATION"
	AIBotTimeseriesGroupsParamsContentTypeOther         AIBotTimeseriesGroupsParamsContentType = "OTHER"
)

func (r AIBotTimeseriesGroupsParamsContentType) IsKnown() bool {
	switch r {
	case AIBotTimeseriesGroupsParamsContentTypeHTML, AIBotTimeseriesGroupsParamsContentTypeImages, AIBotTimeseriesGroupsParamsContentTypeJson, AIBotTimeseriesGroupsParamsContentTypeJavascript, AIBotTimeseriesGroupsParamsContentTypeCSS, AIBotTimeseriesGroupsParamsContentTypePlainText, AIBotTimeseriesGroupsParamsContentTypeFonts, AIBotTimeseriesGroupsParamsContentTypeXml, AIBotTimeseriesGroupsParamsContentTypeYaml, AIBotTimeseriesGroupsParamsContentTypeVideo, AIBotTimeseriesGroupsParamsContentTypeAudio, AIBotTimeseriesGroupsParamsContentTypeMarkdown, AIBotTimeseriesGroupsParamsContentTypeDocuments, AIBotTimeseriesGroupsParamsContentTypeBinary, AIBotTimeseriesGroupsParamsContentTypeSerialization, AIBotTimeseriesGroupsParamsContentTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type AIBotTimeseriesGroupsParamsFormat string

const (
	AIBotTimeseriesGroupsParamsFormatJson AIBotTimeseriesGroupsParamsFormat = "JSON"
	AIBotTimeseriesGroupsParamsFormatCsv  AIBotTimeseriesGroupsParamsFormat = "CSV"
)

func (r AIBotTimeseriesGroupsParamsFormat) IsKnown() bool {
	switch r {
	case AIBotTimeseriesGroupsParamsFormatJson, AIBotTimeseriesGroupsParamsFormatCsv:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AIBotTimeseriesGroupsParamsNormalization string

const (
	AIBotTimeseriesGroupsParamsNormalizationPercentage AIBotTimeseriesGroupsParamsNormalization = "PERCENTAGE"
	AIBotTimeseriesGroupsParamsNormalizationMin0Max    AIBotTimeseriesGroupsParamsNormalization = "MIN0_MAX"
)

func (r AIBotTimeseriesGroupsParamsNormalization) IsKnown() bool {
	switch r {
	case AIBotTimeseriesGroupsParamsNormalizationPercentage, AIBotTimeseriesGroupsParamsNormalizationMin0Max:
		return true
	}
	return false
}

type AIBotTimeseriesGroupsResponseEnvelope struct {
	Result  AIBotTimeseriesGroupsResponse             `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    aiBotTimeseriesGroupsResponseEnvelopeJSON `json:"-"`
}

// aiBotTimeseriesGroupsResponseEnvelopeJSON contains the JSON metadata for the
// struct [AIBotTimeseriesGroupsResponseEnvelope]
type aiBotTimeseriesGroupsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotTimeseriesGroupsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotTimeseriesGroupsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
