// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// AITimeseriesGroupService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAITimeseriesGroupService] method instead.
type AITimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewAITimeseriesGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAITimeseriesGroupService(opts ...option.RequestOption) (r *AITimeseriesGroupService) {
	r = &AITimeseriesGroupService{}
	r.Options = opts
	return
}

// Retrieves the distribution of traffic by AI user agent over time.
func (r *AITimeseriesGroupService) UserAgent(ctx context.Context, query AITimeseriesGroupUserAgentParams, opts ...option.RequestOption) (res *AITimeseriesGroupUserAgentResponse, err error) {
	var env AITimeseriesGroupUserAgentResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/ai/bots/timeseries_groups/user_agent"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AITimeseriesGroupUserAgentResponse struct {
	// Metadata for the results.
	Meta   AITimeseriesGroupUserAgentResponseMeta   `json:"meta,required"`
	Serie0 AITimeseriesGroupUserAgentResponseSerie0 `json:"serie_0,required"`
	JSON   aiTimeseriesGroupUserAgentResponseJSON   `json:"-"`
}

// aiTimeseriesGroupUserAgentResponseJSON contains the JSON metadata for the struct
// [AITimeseriesGroupUserAgentResponse]
type aiTimeseriesGroupUserAgentResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITimeseriesGroupUserAgentResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupUserAgentResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AITimeseriesGroupUserAgentResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AITimeseriesGroupUserAgentResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AITimeseriesGroupUserAgentResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AITimeseriesGroupUserAgentResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AITimeseriesGroupUserAgentResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AITimeseriesGroupUserAgentResponseMetaUnit `json:"units,required"`
	JSON  aiTimeseriesGroupUserAgentResponseMetaJSON   `json:"-"`
}

// aiTimeseriesGroupUserAgentResponseMetaJSON contains the JSON metadata for the
// struct [AITimeseriesGroupUserAgentResponseMeta]
type aiTimeseriesGroupUserAgentResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AITimeseriesGroupUserAgentResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupUserAgentResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AITimeseriesGroupUserAgentResponseMetaAggInterval string

const (
	AITimeseriesGroupUserAgentResponseMetaAggIntervalFifteenMinutes AITimeseriesGroupUserAgentResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AITimeseriesGroupUserAgentResponseMetaAggIntervalOneHour        AITimeseriesGroupUserAgentResponseMetaAggInterval = "ONE_HOUR"
	AITimeseriesGroupUserAgentResponseMetaAggIntervalOneDay         AITimeseriesGroupUserAgentResponseMetaAggInterval = "ONE_DAY"
	AITimeseriesGroupUserAgentResponseMetaAggIntervalOneWeek        AITimeseriesGroupUserAgentResponseMetaAggInterval = "ONE_WEEK"
	AITimeseriesGroupUserAgentResponseMetaAggIntervalOneMonth       AITimeseriesGroupUserAgentResponseMetaAggInterval = "ONE_MONTH"
)

func (r AITimeseriesGroupUserAgentResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AITimeseriesGroupUserAgentResponseMetaAggIntervalFifteenMinutes, AITimeseriesGroupUserAgentResponseMetaAggIntervalOneHour, AITimeseriesGroupUserAgentResponseMetaAggIntervalOneDay, AITimeseriesGroupUserAgentResponseMetaAggIntervalOneWeek, AITimeseriesGroupUserAgentResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AITimeseriesGroupUserAgentResponseMetaConfidenceInfo struct {
	Annotations []AITimeseriesGroupUserAgentResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                    `json:"level,required"`
	JSON  aiTimeseriesGroupUserAgentResponseMetaConfidenceInfoJSON `json:"-"`
}

// aiTimeseriesGroupUserAgentResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [AITimeseriesGroupUserAgentResponseMetaConfidenceInfo]
type aiTimeseriesGroupUserAgentResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITimeseriesGroupUserAgentResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupUserAgentResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AITimeseriesGroupUserAgentResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                               `json:"isInstantaneous,required"`
	LinkedURL       string                                                             `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                          `json:"startDate,required" format:"date-time"`
	JSON            aiTimeseriesGroupUserAgentResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// aiTimeseriesGroupUserAgentResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AITimeseriesGroupUserAgentResponseMetaConfidenceInfoAnnotation]
type aiTimeseriesGroupUserAgentResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AITimeseriesGroupUserAgentResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupUserAgentResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AITimeseriesGroupUserAgentResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                           `json:"startTime,required" format:"date-time"`
	JSON      aiTimeseriesGroupUserAgentResponseMetaDateRangeJSON `json:"-"`
}

// aiTimeseriesGroupUserAgentResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [AITimeseriesGroupUserAgentResponseMetaDateRange]
type aiTimeseriesGroupUserAgentResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITimeseriesGroupUserAgentResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupUserAgentResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AITimeseriesGroupUserAgentResponseMetaNormalization string

const (
	AITimeseriesGroupUserAgentResponseMetaNormalizationPercentage           AITimeseriesGroupUserAgentResponseMetaNormalization = "PERCENTAGE"
	AITimeseriesGroupUserAgentResponseMetaNormalizationMin0Max              AITimeseriesGroupUserAgentResponseMetaNormalization = "MIN0_MAX"
	AITimeseriesGroupUserAgentResponseMetaNormalizationMinMax               AITimeseriesGroupUserAgentResponseMetaNormalization = "MIN_MAX"
	AITimeseriesGroupUserAgentResponseMetaNormalizationRawValues            AITimeseriesGroupUserAgentResponseMetaNormalization = "RAW_VALUES"
	AITimeseriesGroupUserAgentResponseMetaNormalizationPercentageChange     AITimeseriesGroupUserAgentResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AITimeseriesGroupUserAgentResponseMetaNormalizationRollingAverage       AITimeseriesGroupUserAgentResponseMetaNormalization = "ROLLING_AVERAGE"
	AITimeseriesGroupUserAgentResponseMetaNormalizationOverlappedPercentage AITimeseriesGroupUserAgentResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AITimeseriesGroupUserAgentResponseMetaNormalizationRatio                AITimeseriesGroupUserAgentResponseMetaNormalization = "RATIO"
)

func (r AITimeseriesGroupUserAgentResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AITimeseriesGroupUserAgentResponseMetaNormalizationPercentage, AITimeseriesGroupUserAgentResponseMetaNormalizationMin0Max, AITimeseriesGroupUserAgentResponseMetaNormalizationMinMax, AITimeseriesGroupUserAgentResponseMetaNormalizationRawValues, AITimeseriesGroupUserAgentResponseMetaNormalizationPercentageChange, AITimeseriesGroupUserAgentResponseMetaNormalizationRollingAverage, AITimeseriesGroupUserAgentResponseMetaNormalizationOverlappedPercentage, AITimeseriesGroupUserAgentResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AITimeseriesGroupUserAgentResponseMetaUnit struct {
	Name  string                                         `json:"name,required"`
	Value string                                         `json:"value,required"`
	JSON  aiTimeseriesGroupUserAgentResponseMetaUnitJSON `json:"-"`
}

// aiTimeseriesGroupUserAgentResponseMetaUnitJSON contains the JSON metadata for
// the struct [AITimeseriesGroupUserAgentResponseMetaUnit]
type aiTimeseriesGroupUserAgentResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITimeseriesGroupUserAgentResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupUserAgentResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AITimeseriesGroupUserAgentResponseSerie0 struct {
	Timestamps  []time.Time                                  `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                          `json:"-,extras"`
	JSON        aiTimeseriesGroupUserAgentResponseSerie0JSON `json:"-"`
}

// aiTimeseriesGroupUserAgentResponseSerie0JSON contains the JSON metadata for the
// struct [AITimeseriesGroupUserAgentResponseSerie0]
type aiTimeseriesGroupUserAgentResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITimeseriesGroupUserAgentResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupUserAgentResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AITimeseriesGroupUserAgentParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AITimeseriesGroupUserAgentParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[AITimeseriesGroupUserAgentParamsFormat] `query:"format"`
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

// URLQuery serializes [AITimeseriesGroupUserAgentParams]'s query parameters as
// `url.Values`.
func (r AITimeseriesGroupUserAgentParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AITimeseriesGroupUserAgentParamsAggInterval string

const (
	AITimeseriesGroupUserAgentParamsAggInterval15m AITimeseriesGroupUserAgentParamsAggInterval = "15m"
	AITimeseriesGroupUserAgentParamsAggInterval1h  AITimeseriesGroupUserAgentParamsAggInterval = "1h"
	AITimeseriesGroupUserAgentParamsAggInterval1d  AITimeseriesGroupUserAgentParamsAggInterval = "1d"
	AITimeseriesGroupUserAgentParamsAggInterval1w  AITimeseriesGroupUserAgentParamsAggInterval = "1w"
)

func (r AITimeseriesGroupUserAgentParamsAggInterval) IsKnown() bool {
	switch r {
	case AITimeseriesGroupUserAgentParamsAggInterval15m, AITimeseriesGroupUserAgentParamsAggInterval1h, AITimeseriesGroupUserAgentParamsAggInterval1d, AITimeseriesGroupUserAgentParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type AITimeseriesGroupUserAgentParamsFormat string

const (
	AITimeseriesGroupUserAgentParamsFormatJson AITimeseriesGroupUserAgentParamsFormat = "JSON"
	AITimeseriesGroupUserAgentParamsFormatCsv  AITimeseriesGroupUserAgentParamsFormat = "CSV"
)

func (r AITimeseriesGroupUserAgentParamsFormat) IsKnown() bool {
	switch r {
	case AITimeseriesGroupUserAgentParamsFormatJson, AITimeseriesGroupUserAgentParamsFormatCsv:
		return true
	}
	return false
}

type AITimeseriesGroupUserAgentResponseEnvelope struct {
	Result  AITimeseriesGroupUserAgentResponse             `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    aiTimeseriesGroupUserAgentResponseEnvelopeJSON `json:"-"`
}

// aiTimeseriesGroupUserAgentResponseEnvelopeJSON contains the JSON metadata for
// the struct [AITimeseriesGroupUserAgentResponseEnvelope]
type aiTimeseriesGroupUserAgentResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITimeseriesGroupUserAgentResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupUserAgentResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
