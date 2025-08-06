// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// AIBotSummaryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIBotSummaryService] method instead.
type AIBotSummaryService struct {
	Options []option.RequestOption
}

// NewAIBotSummaryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIBotSummaryService(opts ...option.RequestOption) (r *AIBotSummaryService) {
	r = &AIBotSummaryService{}
	r.Options = opts
	return
}

// Retrieves an aggregated summary of AI bots HTTP requests grouped by the
// specified dimension.
func (r *AIBotSummaryService) Summary(ctx context.Context, dimension AIBotSummarySummaryParamsDimension, query AIBotSummarySummaryParams, opts ...option.RequestOption) (res *AIBotSummarySummaryResponse, err error) {
	var env AIBotSummarySummaryResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/ai/bots/summary/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves AI bots HTTP request volume over time.
func (r *AIBotSummaryService) Timeseries(ctx context.Context, query AIBotSummaryTimeseriesParams, opts ...option.RequestOption) (res *AIBotSummaryTimeseriesResponse, err error) {
	var env AIBotSummaryTimeseriesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/ai/bots/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of HTTP requests from AI bots, grouped by chosen the
// specified dimension over time.
func (r *AIBotSummaryService) TimeseriesGroups(ctx context.Context, dimension AIBotSummaryTimeseriesGroupsParamsDimension, query AIBotSummaryTimeseriesGroupsParams, opts ...option.RequestOption) (res *AIBotSummaryTimeseriesGroupsResponse, err error) {
	var env AIBotSummaryTimeseriesGroupsResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/ai/bots/timeseries_groups/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AIBotSummarySummaryResponse struct {
	// Metadata for the results.
	Meta     AIBotSummarySummaryResponseMeta `json:"meta,required"`
	Summary0 map[string]string               `json:"summary_0,required"`
	JSON     aiBotSummarySummaryResponseJSON `json:"-"`
}

// aiBotSummarySummaryResponseJSON contains the JSON metadata for the struct
// [AIBotSummarySummaryResponse]
type aiBotSummarySummaryResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotSummarySummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummarySummaryResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AIBotSummarySummaryResponseMeta struct {
	ConfidenceInfo AIBotSummarySummaryResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AIBotSummarySummaryResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AIBotSummarySummaryResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AIBotSummarySummaryResponseMetaUnit `json:"units,required"`
	JSON  aiBotSummarySummaryResponseMetaJSON   `json:"-"`
}

// aiBotSummarySummaryResponseMetaJSON contains the JSON metadata for the struct
// [AIBotSummarySummaryResponseMeta]
type aiBotSummarySummaryResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AIBotSummarySummaryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummarySummaryResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AIBotSummarySummaryResponseMetaConfidenceInfo struct {
	Annotations []AIBotSummarySummaryResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                             `json:"level,required"`
	JSON  aiBotSummarySummaryResponseMetaConfidenceInfoJSON `json:"-"`
}

// aiBotSummarySummaryResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [AIBotSummarySummaryResponseMetaConfidenceInfo]
type aiBotSummarySummaryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotSummarySummaryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummarySummaryResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AIBotSummarySummaryResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                        `json:"isInstantaneous,required"`
	LinkedURL       string                                                      `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                   `json:"startDate,required" format:"date-time"`
	JSON            aiBotSummarySummaryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// aiBotSummarySummaryResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [AIBotSummarySummaryResponseMetaConfidenceInfoAnnotation]
type aiBotSummarySummaryResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AIBotSummarySummaryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummarySummaryResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AIBotSummarySummaryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                    `json:"startTime,required" format:"date-time"`
	JSON      aiBotSummarySummaryResponseMetaDateRangeJSON `json:"-"`
}

// aiBotSummarySummaryResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [AIBotSummarySummaryResponseMetaDateRange]
type aiBotSummarySummaryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotSummarySummaryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummarySummaryResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AIBotSummarySummaryResponseMetaNormalization string

const (
	AIBotSummarySummaryResponseMetaNormalizationPercentage           AIBotSummarySummaryResponseMetaNormalization = "PERCENTAGE"
	AIBotSummarySummaryResponseMetaNormalizationMin0Max              AIBotSummarySummaryResponseMetaNormalization = "MIN0_MAX"
	AIBotSummarySummaryResponseMetaNormalizationMinMax               AIBotSummarySummaryResponseMetaNormalization = "MIN_MAX"
	AIBotSummarySummaryResponseMetaNormalizationRawValues            AIBotSummarySummaryResponseMetaNormalization = "RAW_VALUES"
	AIBotSummarySummaryResponseMetaNormalizationPercentageChange     AIBotSummarySummaryResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AIBotSummarySummaryResponseMetaNormalizationRollingAverage       AIBotSummarySummaryResponseMetaNormalization = "ROLLING_AVERAGE"
	AIBotSummarySummaryResponseMetaNormalizationOverlappedPercentage AIBotSummarySummaryResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AIBotSummarySummaryResponseMetaNormalizationRatio                AIBotSummarySummaryResponseMetaNormalization = "RATIO"
)

func (r AIBotSummarySummaryResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AIBotSummarySummaryResponseMetaNormalizationPercentage, AIBotSummarySummaryResponseMetaNormalizationMin0Max, AIBotSummarySummaryResponseMetaNormalizationMinMax, AIBotSummarySummaryResponseMetaNormalizationRawValues, AIBotSummarySummaryResponseMetaNormalizationPercentageChange, AIBotSummarySummaryResponseMetaNormalizationRollingAverage, AIBotSummarySummaryResponseMetaNormalizationOverlappedPercentage, AIBotSummarySummaryResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AIBotSummarySummaryResponseMetaUnit struct {
	Name  string                                  `json:"name,required"`
	Value string                                  `json:"value,required"`
	JSON  aiBotSummarySummaryResponseMetaUnitJSON `json:"-"`
}

// aiBotSummarySummaryResponseMetaUnitJSON contains the JSON metadata for the
// struct [AIBotSummarySummaryResponseMetaUnit]
type aiBotSummarySummaryResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotSummarySummaryResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummarySummaryResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AIBotSummaryTimeseriesResponse struct {
	// Metadata for the results.
	Meta        AIBotSummaryTimeseriesResponseMeta        `json:"meta,required"`
	ExtraFields map[string]AIBotSummaryTimeseriesResponse `json:"-,extras"`
	JSON        aiBotSummaryTimeseriesResponseJSON        `json:"-"`
}

// aiBotSummaryTimeseriesResponseJSON contains the JSON metadata for the struct
// [AIBotSummaryTimeseriesResponse]
type aiBotSummaryTimeseriesResponseJSON struct {
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotSummaryTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AIBotSummaryTimeseriesResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AIBotSummaryTimeseriesResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AIBotSummaryTimeseriesResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AIBotSummaryTimeseriesResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AIBotSummaryTimeseriesResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AIBotSummaryTimeseriesResponseMetaUnit `json:"units,required"`
	JSON  aiBotSummaryTimeseriesResponseMetaJSON   `json:"-"`
}

// aiBotSummaryTimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [AIBotSummaryTimeseriesResponseMeta]
type aiBotSummaryTimeseriesResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AIBotSummaryTimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryTimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AIBotSummaryTimeseriesResponseMetaAggInterval string

const (
	AIBotSummaryTimeseriesResponseMetaAggIntervalFifteenMinutes AIBotSummaryTimeseriesResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AIBotSummaryTimeseriesResponseMetaAggIntervalOneHour        AIBotSummaryTimeseriesResponseMetaAggInterval = "ONE_HOUR"
	AIBotSummaryTimeseriesResponseMetaAggIntervalOneDay         AIBotSummaryTimeseriesResponseMetaAggInterval = "ONE_DAY"
	AIBotSummaryTimeseriesResponseMetaAggIntervalOneWeek        AIBotSummaryTimeseriesResponseMetaAggInterval = "ONE_WEEK"
	AIBotSummaryTimeseriesResponseMetaAggIntervalOneMonth       AIBotSummaryTimeseriesResponseMetaAggInterval = "ONE_MONTH"
)

func (r AIBotSummaryTimeseriesResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AIBotSummaryTimeseriesResponseMetaAggIntervalFifteenMinutes, AIBotSummaryTimeseriesResponseMetaAggIntervalOneHour, AIBotSummaryTimeseriesResponseMetaAggIntervalOneDay, AIBotSummaryTimeseriesResponseMetaAggIntervalOneWeek, AIBotSummaryTimeseriesResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AIBotSummaryTimeseriesResponseMetaConfidenceInfo struct {
	Annotations []AIBotSummaryTimeseriesResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                `json:"level,required"`
	JSON  aiBotSummaryTimeseriesResponseMetaConfidenceInfoJSON `json:"-"`
}

// aiBotSummaryTimeseriesResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AIBotSummaryTimeseriesResponseMetaConfidenceInfo]
type aiBotSummaryTimeseriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotSummaryTimeseriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryTimeseriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AIBotSummaryTimeseriesResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                           `json:"isInstantaneous,required"`
	LinkedURL       string                                                         `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                      `json:"startDate,required" format:"date-time"`
	JSON            aiBotSummaryTimeseriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// aiBotSummaryTimeseriesResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [AIBotSummaryTimeseriesResponseMetaConfidenceInfoAnnotation]
type aiBotSummaryTimeseriesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AIBotSummaryTimeseriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryTimeseriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AIBotSummaryTimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      aiBotSummaryTimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// aiBotSummaryTimeseriesResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AIBotSummaryTimeseriesResponseMetaDateRange]
type aiBotSummaryTimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotSummaryTimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryTimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AIBotSummaryTimeseriesResponseMetaNormalization string

const (
	AIBotSummaryTimeseriesResponseMetaNormalizationPercentage           AIBotSummaryTimeseriesResponseMetaNormalization = "PERCENTAGE"
	AIBotSummaryTimeseriesResponseMetaNormalizationMin0Max              AIBotSummaryTimeseriesResponseMetaNormalization = "MIN0_MAX"
	AIBotSummaryTimeseriesResponseMetaNormalizationMinMax               AIBotSummaryTimeseriesResponseMetaNormalization = "MIN_MAX"
	AIBotSummaryTimeseriesResponseMetaNormalizationRawValues            AIBotSummaryTimeseriesResponseMetaNormalization = "RAW_VALUES"
	AIBotSummaryTimeseriesResponseMetaNormalizationPercentageChange     AIBotSummaryTimeseriesResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AIBotSummaryTimeseriesResponseMetaNormalizationRollingAverage       AIBotSummaryTimeseriesResponseMetaNormalization = "ROLLING_AVERAGE"
	AIBotSummaryTimeseriesResponseMetaNormalizationOverlappedPercentage AIBotSummaryTimeseriesResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AIBotSummaryTimeseriesResponseMetaNormalizationRatio                AIBotSummaryTimeseriesResponseMetaNormalization = "RATIO"
)

func (r AIBotSummaryTimeseriesResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AIBotSummaryTimeseriesResponseMetaNormalizationPercentage, AIBotSummaryTimeseriesResponseMetaNormalizationMin0Max, AIBotSummaryTimeseriesResponseMetaNormalizationMinMax, AIBotSummaryTimeseriesResponseMetaNormalizationRawValues, AIBotSummaryTimeseriesResponseMetaNormalizationPercentageChange, AIBotSummaryTimeseriesResponseMetaNormalizationRollingAverage, AIBotSummaryTimeseriesResponseMetaNormalizationOverlappedPercentage, AIBotSummaryTimeseriesResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AIBotSummaryTimeseriesResponseMetaUnit struct {
	Name  string                                     `json:"name,required"`
	Value string                                     `json:"value,required"`
	JSON  aiBotSummaryTimeseriesResponseMetaUnitJSON `json:"-"`
}

// aiBotSummaryTimeseriesResponseMetaUnitJSON contains the JSON metadata for the
// struct [AIBotSummaryTimeseriesResponseMetaUnit]
type aiBotSummaryTimeseriesResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotSummaryTimeseriesResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryTimeseriesResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AIBotSummaryTimeseriesGroupsResponse struct {
	// Metadata for the results.
	Meta   AIBotSummaryTimeseriesGroupsResponseMeta   `json:"meta,required"`
	Serie0 AIBotSummaryTimeseriesGroupsResponseSerie0 `json:"serie_0,required"`
	JSON   aiBotSummaryTimeseriesGroupsResponseJSON   `json:"-"`
}

// aiBotSummaryTimeseriesGroupsResponseJSON contains the JSON metadata for the
// struct [AIBotSummaryTimeseriesGroupsResponse]
type aiBotSummaryTimeseriesGroupsResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotSummaryTimeseriesGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryTimeseriesGroupsResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AIBotSummaryTimeseriesGroupsResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AIBotSummaryTimeseriesGroupsResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AIBotSummaryTimeseriesGroupsResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AIBotSummaryTimeseriesGroupsResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AIBotSummaryTimeseriesGroupsResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AIBotSummaryTimeseriesGroupsResponseMetaUnit `json:"units,required"`
	JSON  aiBotSummaryTimeseriesGroupsResponseMetaJSON   `json:"-"`
}

// aiBotSummaryTimeseriesGroupsResponseMetaJSON contains the JSON metadata for the
// struct [AIBotSummaryTimeseriesGroupsResponseMeta]
type aiBotSummaryTimeseriesGroupsResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AIBotSummaryTimeseriesGroupsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryTimeseriesGroupsResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AIBotSummaryTimeseriesGroupsResponseMetaAggInterval string

const (
	AIBotSummaryTimeseriesGroupsResponseMetaAggIntervalFifteenMinutes AIBotSummaryTimeseriesGroupsResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AIBotSummaryTimeseriesGroupsResponseMetaAggIntervalOneHour        AIBotSummaryTimeseriesGroupsResponseMetaAggInterval = "ONE_HOUR"
	AIBotSummaryTimeseriesGroupsResponseMetaAggIntervalOneDay         AIBotSummaryTimeseriesGroupsResponseMetaAggInterval = "ONE_DAY"
	AIBotSummaryTimeseriesGroupsResponseMetaAggIntervalOneWeek        AIBotSummaryTimeseriesGroupsResponseMetaAggInterval = "ONE_WEEK"
	AIBotSummaryTimeseriesGroupsResponseMetaAggIntervalOneMonth       AIBotSummaryTimeseriesGroupsResponseMetaAggInterval = "ONE_MONTH"
)

func (r AIBotSummaryTimeseriesGroupsResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AIBotSummaryTimeseriesGroupsResponseMetaAggIntervalFifteenMinutes, AIBotSummaryTimeseriesGroupsResponseMetaAggIntervalOneHour, AIBotSummaryTimeseriesGroupsResponseMetaAggIntervalOneDay, AIBotSummaryTimeseriesGroupsResponseMetaAggIntervalOneWeek, AIBotSummaryTimeseriesGroupsResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AIBotSummaryTimeseriesGroupsResponseMetaConfidenceInfo struct {
	Annotations []AIBotSummaryTimeseriesGroupsResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                      `json:"level,required"`
	JSON  aiBotSummaryTimeseriesGroupsResponseMetaConfidenceInfoJSON `json:"-"`
}

// aiBotSummaryTimeseriesGroupsResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [AIBotSummaryTimeseriesGroupsResponseMetaConfidenceInfo]
type aiBotSummaryTimeseriesGroupsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotSummaryTimeseriesGroupsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryTimeseriesGroupsResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AIBotSummaryTimeseriesGroupsResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                 `json:"isInstantaneous,required"`
	LinkedURL       string                                                               `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                            `json:"startDate,required" format:"date-time"`
	JSON            aiBotSummaryTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// aiBotSummaryTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [AIBotSummaryTimeseriesGroupsResponseMetaConfidenceInfoAnnotation]
type aiBotSummaryTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AIBotSummaryTimeseriesGroupsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AIBotSummaryTimeseriesGroupsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      aiBotSummaryTimeseriesGroupsResponseMetaDateRangeJSON `json:"-"`
}

// aiBotSummaryTimeseriesGroupsResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [AIBotSummaryTimeseriesGroupsResponseMetaDateRange]
type aiBotSummaryTimeseriesGroupsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotSummaryTimeseriesGroupsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryTimeseriesGroupsResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AIBotSummaryTimeseriesGroupsResponseMetaNormalization string

const (
	AIBotSummaryTimeseriesGroupsResponseMetaNormalizationPercentage           AIBotSummaryTimeseriesGroupsResponseMetaNormalization = "PERCENTAGE"
	AIBotSummaryTimeseriesGroupsResponseMetaNormalizationMin0Max              AIBotSummaryTimeseriesGroupsResponseMetaNormalization = "MIN0_MAX"
	AIBotSummaryTimeseriesGroupsResponseMetaNormalizationMinMax               AIBotSummaryTimeseriesGroupsResponseMetaNormalization = "MIN_MAX"
	AIBotSummaryTimeseriesGroupsResponseMetaNormalizationRawValues            AIBotSummaryTimeseriesGroupsResponseMetaNormalization = "RAW_VALUES"
	AIBotSummaryTimeseriesGroupsResponseMetaNormalizationPercentageChange     AIBotSummaryTimeseriesGroupsResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AIBotSummaryTimeseriesGroupsResponseMetaNormalizationRollingAverage       AIBotSummaryTimeseriesGroupsResponseMetaNormalization = "ROLLING_AVERAGE"
	AIBotSummaryTimeseriesGroupsResponseMetaNormalizationOverlappedPercentage AIBotSummaryTimeseriesGroupsResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AIBotSummaryTimeseriesGroupsResponseMetaNormalizationRatio                AIBotSummaryTimeseriesGroupsResponseMetaNormalization = "RATIO"
)

func (r AIBotSummaryTimeseriesGroupsResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AIBotSummaryTimeseriesGroupsResponseMetaNormalizationPercentage, AIBotSummaryTimeseriesGroupsResponseMetaNormalizationMin0Max, AIBotSummaryTimeseriesGroupsResponseMetaNormalizationMinMax, AIBotSummaryTimeseriesGroupsResponseMetaNormalizationRawValues, AIBotSummaryTimeseriesGroupsResponseMetaNormalizationPercentageChange, AIBotSummaryTimeseriesGroupsResponseMetaNormalizationRollingAverage, AIBotSummaryTimeseriesGroupsResponseMetaNormalizationOverlappedPercentage, AIBotSummaryTimeseriesGroupsResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AIBotSummaryTimeseriesGroupsResponseMetaUnit struct {
	Name  string                                           `json:"name,required"`
	Value string                                           `json:"value,required"`
	JSON  aiBotSummaryTimeseriesGroupsResponseMetaUnitJSON `json:"-"`
}

// aiBotSummaryTimeseriesGroupsResponseMetaUnitJSON contains the JSON metadata for
// the struct [AIBotSummaryTimeseriesGroupsResponseMetaUnit]
type aiBotSummaryTimeseriesGroupsResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotSummaryTimeseriesGroupsResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryTimeseriesGroupsResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AIBotSummaryTimeseriesGroupsResponseSerie0 struct {
	Timestamps  []time.Time                                    `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                            `json:"-,extras"`
	JSON        aiBotSummaryTimeseriesGroupsResponseSerie0JSON `json:"-"`
}

// aiBotSummaryTimeseriesGroupsResponseSerie0JSON contains the JSON metadata for
// the struct [AIBotSummaryTimeseriesGroupsResponseSerie0]
type aiBotSummaryTimeseriesGroupsResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotSummaryTimeseriesGroupsResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryTimeseriesGroupsResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AIBotSummarySummaryParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
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
	Format param.Field[AIBotSummarySummaryParamsFormat] `query:"format"`
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

// URLQuery serializes [AIBotSummarySummaryParams]'s query parameters as
// `url.Values`.
func (r AIBotSummarySummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the attribute by which to group the results.
type AIBotSummarySummaryParamsDimension string

const (
	AIBotSummarySummaryParamsDimensionUserAgent    AIBotSummarySummaryParamsDimension = "USER_AGENT"
	AIBotSummarySummaryParamsDimensionCrawlPurpose AIBotSummarySummaryParamsDimension = "CRAWL_PURPOSE"
)

func (r AIBotSummarySummaryParamsDimension) IsKnown() bool {
	switch r {
	case AIBotSummarySummaryParamsDimensionUserAgent, AIBotSummarySummaryParamsDimensionCrawlPurpose:
		return true
	}
	return false
}

// Format in which results will be returned.
type AIBotSummarySummaryParamsFormat string

const (
	AIBotSummarySummaryParamsFormatJson AIBotSummarySummaryParamsFormat = "JSON"
	AIBotSummarySummaryParamsFormatCsv  AIBotSummarySummaryParamsFormat = "CSV"
)

func (r AIBotSummarySummaryParamsFormat) IsKnown() bool {
	switch r {
	case AIBotSummarySummaryParamsFormatJson, AIBotSummarySummaryParamsFormatCsv:
		return true
	}
	return false
}

type AIBotSummarySummaryResponseEnvelope struct {
	Result  AIBotSummarySummaryResponse             `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    aiBotSummarySummaryResponseEnvelopeJSON `json:"-"`
}

// aiBotSummarySummaryResponseEnvelopeJSON contains the JSON metadata for the
// struct [AIBotSummarySummaryResponseEnvelope]
type aiBotSummarySummaryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotSummarySummaryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummarySummaryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AIBotSummaryTimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AIBotSummaryTimeseriesParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
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
	Format param.Field[AIBotSummaryTimeseriesParamsFormat] `query:"format"`
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
}

// URLQuery serializes [AIBotSummaryTimeseriesParams]'s query parameters as
// `url.Values`.
func (r AIBotSummaryTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AIBotSummaryTimeseriesParamsAggInterval string

const (
	AIBotSummaryTimeseriesParamsAggInterval15m AIBotSummaryTimeseriesParamsAggInterval = "15m"
	AIBotSummaryTimeseriesParamsAggInterval1h  AIBotSummaryTimeseriesParamsAggInterval = "1h"
	AIBotSummaryTimeseriesParamsAggInterval1d  AIBotSummaryTimeseriesParamsAggInterval = "1d"
	AIBotSummaryTimeseriesParamsAggInterval1w  AIBotSummaryTimeseriesParamsAggInterval = "1w"
)

func (r AIBotSummaryTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case AIBotSummaryTimeseriesParamsAggInterval15m, AIBotSummaryTimeseriesParamsAggInterval1h, AIBotSummaryTimeseriesParamsAggInterval1d, AIBotSummaryTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type AIBotSummaryTimeseriesParamsFormat string

const (
	AIBotSummaryTimeseriesParamsFormatJson AIBotSummaryTimeseriesParamsFormat = "JSON"
	AIBotSummaryTimeseriesParamsFormatCsv  AIBotSummaryTimeseriesParamsFormat = "CSV"
)

func (r AIBotSummaryTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case AIBotSummaryTimeseriesParamsFormatJson, AIBotSummaryTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type AIBotSummaryTimeseriesResponseEnvelope struct {
	Result  AIBotSummaryTimeseriesResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    aiBotSummaryTimeseriesResponseEnvelopeJSON `json:"-"`
}

// aiBotSummaryTimeseriesResponseEnvelopeJSON contains the JSON metadata for the
// struct [AIBotSummaryTimeseriesResponseEnvelope]
type aiBotSummaryTimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotSummaryTimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryTimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AIBotSummaryTimeseriesGroupsParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AIBotSummaryTimeseriesGroupsParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
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
	Format param.Field[AIBotSummaryTimeseriesGroupsParamsFormat] `query:"format"`
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
	Normalization param.Field[AIBotSummaryTimeseriesGroupsParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [AIBotSummaryTimeseriesGroupsParams]'s query parameters as
// `url.Values`.
func (r AIBotSummaryTimeseriesGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the attribute by which to group the results.
type AIBotSummaryTimeseriesGroupsParamsDimension string

const (
	AIBotSummaryTimeseriesGroupsParamsDimensionUserAgent    AIBotSummaryTimeseriesGroupsParamsDimension = "USER_AGENT"
	AIBotSummaryTimeseriesGroupsParamsDimensionCrawlPurpose AIBotSummaryTimeseriesGroupsParamsDimension = "CRAWL_PURPOSE"
)

func (r AIBotSummaryTimeseriesGroupsParamsDimension) IsKnown() bool {
	switch r {
	case AIBotSummaryTimeseriesGroupsParamsDimensionUserAgent, AIBotSummaryTimeseriesGroupsParamsDimensionCrawlPurpose:
		return true
	}
	return false
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AIBotSummaryTimeseriesGroupsParamsAggInterval string

const (
	AIBotSummaryTimeseriesGroupsParamsAggInterval15m AIBotSummaryTimeseriesGroupsParamsAggInterval = "15m"
	AIBotSummaryTimeseriesGroupsParamsAggInterval1h  AIBotSummaryTimeseriesGroupsParamsAggInterval = "1h"
	AIBotSummaryTimeseriesGroupsParamsAggInterval1d  AIBotSummaryTimeseriesGroupsParamsAggInterval = "1d"
	AIBotSummaryTimeseriesGroupsParamsAggInterval1w  AIBotSummaryTimeseriesGroupsParamsAggInterval = "1w"
)

func (r AIBotSummaryTimeseriesGroupsParamsAggInterval) IsKnown() bool {
	switch r {
	case AIBotSummaryTimeseriesGroupsParamsAggInterval15m, AIBotSummaryTimeseriesGroupsParamsAggInterval1h, AIBotSummaryTimeseriesGroupsParamsAggInterval1d, AIBotSummaryTimeseriesGroupsParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type AIBotSummaryTimeseriesGroupsParamsFormat string

const (
	AIBotSummaryTimeseriesGroupsParamsFormatJson AIBotSummaryTimeseriesGroupsParamsFormat = "JSON"
	AIBotSummaryTimeseriesGroupsParamsFormatCsv  AIBotSummaryTimeseriesGroupsParamsFormat = "CSV"
)

func (r AIBotSummaryTimeseriesGroupsParamsFormat) IsKnown() bool {
	switch r {
	case AIBotSummaryTimeseriesGroupsParamsFormatJson, AIBotSummaryTimeseriesGroupsParamsFormatCsv:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AIBotSummaryTimeseriesGroupsParamsNormalization string

const (
	AIBotSummaryTimeseriesGroupsParamsNormalizationPercentageChange AIBotSummaryTimeseriesGroupsParamsNormalization = "PERCENTAGE_CHANGE"
	AIBotSummaryTimeseriesGroupsParamsNormalizationMin0Max          AIBotSummaryTimeseriesGroupsParamsNormalization = "MIN0_MAX"
)

func (r AIBotSummaryTimeseriesGroupsParamsNormalization) IsKnown() bool {
	switch r {
	case AIBotSummaryTimeseriesGroupsParamsNormalizationPercentageChange, AIBotSummaryTimeseriesGroupsParamsNormalizationMin0Max:
		return true
	}
	return false
}

type AIBotSummaryTimeseriesGroupsResponseEnvelope struct {
	Result  AIBotSummaryTimeseriesGroupsResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    aiBotSummaryTimeseriesGroupsResponseEnvelopeJSON `json:"-"`
}

// aiBotSummaryTimeseriesGroupsResponseEnvelopeJSON contains the JSON metadata for
// the struct [AIBotSummaryTimeseriesGroupsResponseEnvelope]
type aiBotSummaryTimeseriesGroupsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotSummaryTimeseriesGroupsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryTimeseriesGroupsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
