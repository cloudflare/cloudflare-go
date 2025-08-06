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

// AIBotService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIBotService] method instead.
type AIBotService struct {
	Options []option.RequestOption
}

// NewAIBotService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAIBotService(opts ...option.RequestOption) (r *AIBotService) {
	r = &AIBotService{}
	r.Options = opts
	return
}

// Retrieves an aggregated summary of AI bots HTTP requests grouped by the
// specified dimension.
func (r *AIBotService) Summary(ctx context.Context, dimension AIBotSummaryParamsDimension, query AIBotSummaryParams, opts ...option.RequestOption) (res *AIBotSummaryResponse, err error) {
	var env AIBotSummaryResponseEnvelope
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
func (r *AIBotService) Timeseries(ctx context.Context, query AIBotTimeseriesParams, opts ...option.RequestOption) (res *AIBotTimeseriesResponse, err error) {
	var env AIBotTimeseriesResponseEnvelope
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
func (r *AIBotService) TimeseriesGroups(ctx context.Context, dimension AIBotTimeseriesGroupsParamsDimension, query AIBotTimeseriesGroupsParams, opts ...option.RequestOption) (res *AIBotTimeseriesGroupsResponse, err error) {
	var env AIBotTimeseriesGroupsResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/ai/bots/timeseries_groups/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AIBotSummaryResponse struct {
	// Metadata for the results.
	Meta     AIBotSummaryResponseMeta `json:"meta,required"`
	Summary0 map[string]string        `json:"summary_0,required"`
	JSON     aiBotSummaryResponseJSON `json:"-"`
}

// aiBotSummaryResponseJSON contains the JSON metadata for the struct
// [AIBotSummaryResponse]
type aiBotSummaryResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AIBotSummaryResponseMeta struct {
	ConfidenceInfo AIBotSummaryResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AIBotSummaryResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AIBotSummaryResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AIBotSummaryResponseMetaUnit `json:"units,required"`
	JSON  aiBotSummaryResponseMetaJSON   `json:"-"`
}

// aiBotSummaryResponseMetaJSON contains the JSON metadata for the struct
// [AIBotSummaryResponseMeta]
type aiBotSummaryResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AIBotSummaryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AIBotSummaryResponseMetaConfidenceInfo struct {
	Annotations []AIBotSummaryResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                      `json:"level,required"`
	JSON  aiBotSummaryResponseMetaConfidenceInfoJSON `json:"-"`
}

// aiBotSummaryResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [AIBotSummaryResponseMetaConfidenceInfo]
type aiBotSummaryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotSummaryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AIBotSummaryResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                 `json:"isInstantaneous,required"`
	LinkedURL       string                                               `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                            `json:"startDate,required" format:"date-time"`
	JSON            aiBotSummaryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// aiBotSummaryResponseMetaConfidenceInfoAnnotationJSON contains the JSON metadata
// for the struct [AIBotSummaryResponseMetaConfidenceInfoAnnotation]
type aiBotSummaryResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AIBotSummaryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AIBotSummaryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                             `json:"startTime,required" format:"date-time"`
	JSON      aiBotSummaryResponseMetaDateRangeJSON `json:"-"`
}

// aiBotSummaryResponseMetaDateRangeJSON contains the JSON metadata for the struct
// [AIBotSummaryResponseMetaDateRange]
type aiBotSummaryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotSummaryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AIBotSummaryResponseMetaNormalization string

const (
	AIBotSummaryResponseMetaNormalizationPercentage           AIBotSummaryResponseMetaNormalization = "PERCENTAGE"
	AIBotSummaryResponseMetaNormalizationMin0Max              AIBotSummaryResponseMetaNormalization = "MIN0_MAX"
	AIBotSummaryResponseMetaNormalizationMinMax               AIBotSummaryResponseMetaNormalization = "MIN_MAX"
	AIBotSummaryResponseMetaNormalizationRawValues            AIBotSummaryResponseMetaNormalization = "RAW_VALUES"
	AIBotSummaryResponseMetaNormalizationPercentageChange     AIBotSummaryResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AIBotSummaryResponseMetaNormalizationRollingAverage       AIBotSummaryResponseMetaNormalization = "ROLLING_AVERAGE"
	AIBotSummaryResponseMetaNormalizationOverlappedPercentage AIBotSummaryResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AIBotSummaryResponseMetaNormalizationRatio                AIBotSummaryResponseMetaNormalization = "RATIO"
)

func (r AIBotSummaryResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AIBotSummaryResponseMetaNormalizationPercentage, AIBotSummaryResponseMetaNormalizationMin0Max, AIBotSummaryResponseMetaNormalizationMinMax, AIBotSummaryResponseMetaNormalizationRawValues, AIBotSummaryResponseMetaNormalizationPercentageChange, AIBotSummaryResponseMetaNormalizationRollingAverage, AIBotSummaryResponseMetaNormalizationOverlappedPercentage, AIBotSummaryResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AIBotSummaryResponseMetaUnit struct {
	Name  string                           `json:"name,required"`
	Value string                           `json:"value,required"`
	JSON  aiBotSummaryResponseMetaUnitJSON `json:"-"`
}

// aiBotSummaryResponseMetaUnitJSON contains the JSON metadata for the struct
// [AIBotSummaryResponseMetaUnit]
type aiBotSummaryResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotSummaryResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryResponseMetaUnitJSON) RawJSON() string {
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
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
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
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
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

type AIBotSummaryParams struct {
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
	Format param.Field[AIBotSummaryParamsFormat] `query:"format"`
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

// URLQuery serializes [AIBotSummaryParams]'s query parameters as `url.Values`.
func (r AIBotSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the attribute by which to group the results.
type AIBotSummaryParamsDimension string

const (
	AIBotSummaryParamsDimensionUserAgent    AIBotSummaryParamsDimension = "USER_AGENT"
	AIBotSummaryParamsDimensionCrawlPurpose AIBotSummaryParamsDimension = "CRAWL_PURPOSE"
)

func (r AIBotSummaryParamsDimension) IsKnown() bool {
	switch r {
	case AIBotSummaryParamsDimensionUserAgent, AIBotSummaryParamsDimensionCrawlPurpose:
		return true
	}
	return false
}

// Format in which results will be returned.
type AIBotSummaryParamsFormat string

const (
	AIBotSummaryParamsFormatJson AIBotSummaryParamsFormat = "JSON"
	AIBotSummaryParamsFormatCsv  AIBotSummaryParamsFormat = "CSV"
)

func (r AIBotSummaryParamsFormat) IsKnown() bool {
	switch r {
	case AIBotSummaryParamsFormatJson, AIBotSummaryParamsFormatCsv:
		return true
	}
	return false
}

type AIBotSummaryResponseEnvelope struct {
	Result  AIBotSummaryResponse             `json:"result,required"`
	Success bool                             `json:"success,required"`
	JSON    aiBotSummaryResponseEnvelopeJSON `json:"-"`
}

// aiBotSummaryResponseEnvelopeJSON contains the JSON metadata for the struct
// [AIBotSummaryResponseEnvelope]
type aiBotSummaryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotSummaryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryResponseEnvelopeJSON) RawJSON() string {
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
)

func (r AIBotTimeseriesGroupsParamsDimension) IsKnown() bool {
	switch r {
	case AIBotTimeseriesGroupsParamsDimensionUserAgent, AIBotTimeseriesGroupsParamsDimensionCrawlPurpose:
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
	AIBotTimeseriesGroupsParamsNormalizationPercentageChange AIBotTimeseriesGroupsParamsNormalization = "PERCENTAGE_CHANGE"
	AIBotTimeseriesGroupsParamsNormalizationMin0Max          AIBotTimeseriesGroupsParamsNormalization = "MIN0_MAX"
)

func (r AIBotTimeseriesGroupsParamsNormalization) IsKnown() bool {
	switch r {
	case AIBotTimeseriesGroupsParamsNormalizationPercentageChange, AIBotTimeseriesGroupsParamsNormalizationMin0Max:
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
