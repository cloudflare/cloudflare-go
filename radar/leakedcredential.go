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

// LeakedCredentialService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLeakedCredentialService] method instead.
type LeakedCredentialService struct {
	Options          []option.RequestOption
	Summary          *LeakedCredentialSummaryService
	TimeseriesGroups *LeakedCredentialTimeseriesGroupService
}

// NewLeakedCredentialService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLeakedCredentialService(opts ...option.RequestOption) (r *LeakedCredentialService) {
	r = &LeakedCredentialService{}
	r.Options = opts
	r.Summary = NewLeakedCredentialSummaryService(opts...)
	r.TimeseriesGroups = NewLeakedCredentialTimeseriesGroupService(opts...)
	return
}

// Retrieves an aggregated summary of HTTP authentication requests grouped by the
// specified dimension.
func (r *LeakedCredentialService) SummaryV2(ctx context.Context, dimension LeakedCredentialSummaryV2ParamsDimension, query LeakedCredentialSummaryV2Params, opts ...option.RequestOption) (res *LeakedCredentialSummaryV2Response, err error) {
	var env LeakedCredentialSummaryV2ResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/leaked_credential_checks/summary/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of HTTP authentication requests, grouped by the
// specified dimension over time.
func (r *LeakedCredentialService) TimeseriesGroupsV2(ctx context.Context, dimension LeakedCredentialTimeseriesGroupsV2ParamsDimension, query LeakedCredentialTimeseriesGroupsV2Params, opts ...option.RequestOption) (res *LeakedCredentialTimeseriesGroupsV2Response, err error) {
	var env LeakedCredentialTimeseriesGroupsV2ResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/leaked_credential_checks/timeseries_groups/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LeakedCredentialSummaryV2Response struct {
	// Metadata for the results.
	Meta     LeakedCredentialSummaryV2ResponseMeta `json:"meta,required"`
	Summary0 map[string]string                     `json:"summary_0,required"`
	JSON     leakedCredentialSummaryV2ResponseJSON `json:"-"`
}

// leakedCredentialSummaryV2ResponseJSON contains the JSON metadata for the struct
// [LeakedCredentialSummaryV2Response]
type leakedCredentialSummaryV2ResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialSummaryV2Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialSummaryV2ResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type LeakedCredentialSummaryV2ResponseMeta struct {
	ConfidenceInfo LeakedCredentialSummaryV2ResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []LeakedCredentialSummaryV2ResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization LeakedCredentialSummaryV2ResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []LeakedCredentialSummaryV2ResponseMetaUnit `json:"units,required"`
	JSON  leakedCredentialSummaryV2ResponseMetaJSON   `json:"-"`
}

// leakedCredentialSummaryV2ResponseMetaJSON contains the JSON metadata for the
// struct [LeakedCredentialSummaryV2ResponseMeta]
type leakedCredentialSummaryV2ResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LeakedCredentialSummaryV2ResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialSummaryV2ResponseMetaJSON) RawJSON() string {
	return r.raw
}

type LeakedCredentialSummaryV2ResponseMetaConfidenceInfo struct {
	Annotations []LeakedCredentialSummaryV2ResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                   `json:"level,required"`
	JSON  leakedCredentialSummaryV2ResponseMetaConfidenceInfoJSON `json:"-"`
}

// leakedCredentialSummaryV2ResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [LeakedCredentialSummaryV2ResponseMetaConfidenceInfo]
type leakedCredentialSummaryV2ResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialSummaryV2ResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialSummaryV2ResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type LeakedCredentialSummaryV2ResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                              `json:"isInstantaneous,required"`
	LinkedURL       string                                                            `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                         `json:"startDate,required" format:"date-time"`
	JSON            leakedCredentialSummaryV2ResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// leakedCredentialSummaryV2ResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [LeakedCredentialSummaryV2ResponseMetaConfidenceInfoAnnotation]
type leakedCredentialSummaryV2ResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *LeakedCredentialSummaryV2ResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialSummaryV2ResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type LeakedCredentialSummaryV2ResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                          `json:"startTime,required" format:"date-time"`
	JSON      leakedCredentialSummaryV2ResponseMetaDateRangeJSON `json:"-"`
}

// leakedCredentialSummaryV2ResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [LeakedCredentialSummaryV2ResponseMetaDateRange]
type leakedCredentialSummaryV2ResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialSummaryV2ResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialSummaryV2ResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type LeakedCredentialSummaryV2ResponseMetaNormalization string

const (
	LeakedCredentialSummaryV2ResponseMetaNormalizationPercentage           LeakedCredentialSummaryV2ResponseMetaNormalization = "PERCENTAGE"
	LeakedCredentialSummaryV2ResponseMetaNormalizationMin0Max              LeakedCredentialSummaryV2ResponseMetaNormalization = "MIN0_MAX"
	LeakedCredentialSummaryV2ResponseMetaNormalizationMinMax               LeakedCredentialSummaryV2ResponseMetaNormalization = "MIN_MAX"
	LeakedCredentialSummaryV2ResponseMetaNormalizationRawValues            LeakedCredentialSummaryV2ResponseMetaNormalization = "RAW_VALUES"
	LeakedCredentialSummaryV2ResponseMetaNormalizationPercentageChange     LeakedCredentialSummaryV2ResponseMetaNormalization = "PERCENTAGE_CHANGE"
	LeakedCredentialSummaryV2ResponseMetaNormalizationRollingAverage       LeakedCredentialSummaryV2ResponseMetaNormalization = "ROLLING_AVERAGE"
	LeakedCredentialSummaryV2ResponseMetaNormalizationOverlappedPercentage LeakedCredentialSummaryV2ResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	LeakedCredentialSummaryV2ResponseMetaNormalizationRatio                LeakedCredentialSummaryV2ResponseMetaNormalization = "RATIO"
)

func (r LeakedCredentialSummaryV2ResponseMetaNormalization) IsKnown() bool {
	switch r {
	case LeakedCredentialSummaryV2ResponseMetaNormalizationPercentage, LeakedCredentialSummaryV2ResponseMetaNormalizationMin0Max, LeakedCredentialSummaryV2ResponseMetaNormalizationMinMax, LeakedCredentialSummaryV2ResponseMetaNormalizationRawValues, LeakedCredentialSummaryV2ResponseMetaNormalizationPercentageChange, LeakedCredentialSummaryV2ResponseMetaNormalizationRollingAverage, LeakedCredentialSummaryV2ResponseMetaNormalizationOverlappedPercentage, LeakedCredentialSummaryV2ResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type LeakedCredentialSummaryV2ResponseMetaUnit struct {
	Name  string                                        `json:"name,required"`
	Value string                                        `json:"value,required"`
	JSON  leakedCredentialSummaryV2ResponseMetaUnitJSON `json:"-"`
}

// leakedCredentialSummaryV2ResponseMetaUnitJSON contains the JSON metadata for the
// struct [LeakedCredentialSummaryV2ResponseMetaUnit]
type leakedCredentialSummaryV2ResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialSummaryV2ResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialSummaryV2ResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type LeakedCredentialTimeseriesGroupsV2Response struct {
	// Metadata for the results.
	Meta   LeakedCredentialTimeseriesGroupsV2ResponseMeta   `json:"meta,required"`
	Serie0 LeakedCredentialTimeseriesGroupsV2ResponseSerie0 `json:"serie_0,required"`
	JSON   leakedCredentialTimeseriesGroupsV2ResponseJSON   `json:"-"`
}

// leakedCredentialTimeseriesGroupsV2ResponseJSON contains the JSON metadata for
// the struct [LeakedCredentialTimeseriesGroupsV2Response]
type leakedCredentialTimeseriesGroupsV2ResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialTimeseriesGroupsV2Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialTimeseriesGroupsV2ResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type LeakedCredentialTimeseriesGroupsV2ResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    LeakedCredentialTimeseriesGroupsV2ResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo LeakedCredentialTimeseriesGroupsV2ResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []LeakedCredentialTimeseriesGroupsV2ResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization LeakedCredentialTimeseriesGroupsV2ResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []LeakedCredentialTimeseriesGroupsV2ResponseMetaUnit `json:"units,required"`
	JSON  leakedCredentialTimeseriesGroupsV2ResponseMetaJSON   `json:"-"`
}

// leakedCredentialTimeseriesGroupsV2ResponseMetaJSON contains the JSON metadata
// for the struct [LeakedCredentialTimeseriesGroupsV2ResponseMeta]
type leakedCredentialTimeseriesGroupsV2ResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LeakedCredentialTimeseriesGroupsV2ResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialTimeseriesGroupsV2ResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type LeakedCredentialTimeseriesGroupsV2ResponseMetaAggInterval string

const (
	LeakedCredentialTimeseriesGroupsV2ResponseMetaAggIntervalFifteenMinutes LeakedCredentialTimeseriesGroupsV2ResponseMetaAggInterval = "FIFTEEN_MINUTES"
	LeakedCredentialTimeseriesGroupsV2ResponseMetaAggIntervalOneHour        LeakedCredentialTimeseriesGroupsV2ResponseMetaAggInterval = "ONE_HOUR"
	LeakedCredentialTimeseriesGroupsV2ResponseMetaAggIntervalOneDay         LeakedCredentialTimeseriesGroupsV2ResponseMetaAggInterval = "ONE_DAY"
	LeakedCredentialTimeseriesGroupsV2ResponseMetaAggIntervalOneWeek        LeakedCredentialTimeseriesGroupsV2ResponseMetaAggInterval = "ONE_WEEK"
	LeakedCredentialTimeseriesGroupsV2ResponseMetaAggIntervalOneMonth       LeakedCredentialTimeseriesGroupsV2ResponseMetaAggInterval = "ONE_MONTH"
)

func (r LeakedCredentialTimeseriesGroupsV2ResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case LeakedCredentialTimeseriesGroupsV2ResponseMetaAggIntervalFifteenMinutes, LeakedCredentialTimeseriesGroupsV2ResponseMetaAggIntervalOneHour, LeakedCredentialTimeseriesGroupsV2ResponseMetaAggIntervalOneDay, LeakedCredentialTimeseriesGroupsV2ResponseMetaAggIntervalOneWeek, LeakedCredentialTimeseriesGroupsV2ResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type LeakedCredentialTimeseriesGroupsV2ResponseMetaConfidenceInfo struct {
	Annotations []LeakedCredentialTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                            `json:"level,required"`
	JSON  leakedCredentialTimeseriesGroupsV2ResponseMetaConfidenceInfoJSON `json:"-"`
}

// leakedCredentialTimeseriesGroupsV2ResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [LeakedCredentialTimeseriesGroupsV2ResponseMetaConfidenceInfo]
type leakedCredentialTimeseriesGroupsV2ResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialTimeseriesGroupsV2ResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialTimeseriesGroupsV2ResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type LeakedCredentialTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                       `json:"isInstantaneous,required"`
	LinkedURL       string                                                                     `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                  `json:"startDate,required" format:"date-time"`
	JSON            leakedCredentialTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// leakedCredentialTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [LeakedCredentialTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation]
type leakedCredentialTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *LeakedCredentialTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type LeakedCredentialTimeseriesGroupsV2ResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                   `json:"startTime,required" format:"date-time"`
	JSON      leakedCredentialTimeseriesGroupsV2ResponseMetaDateRangeJSON `json:"-"`
}

// leakedCredentialTimeseriesGroupsV2ResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [LeakedCredentialTimeseriesGroupsV2ResponseMetaDateRange]
type leakedCredentialTimeseriesGroupsV2ResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialTimeseriesGroupsV2ResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialTimeseriesGroupsV2ResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type LeakedCredentialTimeseriesGroupsV2ResponseMetaNormalization string

const (
	LeakedCredentialTimeseriesGroupsV2ResponseMetaNormalizationPercentage           LeakedCredentialTimeseriesGroupsV2ResponseMetaNormalization = "PERCENTAGE"
	LeakedCredentialTimeseriesGroupsV2ResponseMetaNormalizationMin0Max              LeakedCredentialTimeseriesGroupsV2ResponseMetaNormalization = "MIN0_MAX"
	LeakedCredentialTimeseriesGroupsV2ResponseMetaNormalizationMinMax               LeakedCredentialTimeseriesGroupsV2ResponseMetaNormalization = "MIN_MAX"
	LeakedCredentialTimeseriesGroupsV2ResponseMetaNormalizationRawValues            LeakedCredentialTimeseriesGroupsV2ResponseMetaNormalization = "RAW_VALUES"
	LeakedCredentialTimeseriesGroupsV2ResponseMetaNormalizationPercentageChange     LeakedCredentialTimeseriesGroupsV2ResponseMetaNormalization = "PERCENTAGE_CHANGE"
	LeakedCredentialTimeseriesGroupsV2ResponseMetaNormalizationRollingAverage       LeakedCredentialTimeseriesGroupsV2ResponseMetaNormalization = "ROLLING_AVERAGE"
	LeakedCredentialTimeseriesGroupsV2ResponseMetaNormalizationOverlappedPercentage LeakedCredentialTimeseriesGroupsV2ResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	LeakedCredentialTimeseriesGroupsV2ResponseMetaNormalizationRatio                LeakedCredentialTimeseriesGroupsV2ResponseMetaNormalization = "RATIO"
)

func (r LeakedCredentialTimeseriesGroupsV2ResponseMetaNormalization) IsKnown() bool {
	switch r {
	case LeakedCredentialTimeseriesGroupsV2ResponseMetaNormalizationPercentage, LeakedCredentialTimeseriesGroupsV2ResponseMetaNormalizationMin0Max, LeakedCredentialTimeseriesGroupsV2ResponseMetaNormalizationMinMax, LeakedCredentialTimeseriesGroupsV2ResponseMetaNormalizationRawValues, LeakedCredentialTimeseriesGroupsV2ResponseMetaNormalizationPercentageChange, LeakedCredentialTimeseriesGroupsV2ResponseMetaNormalizationRollingAverage, LeakedCredentialTimeseriesGroupsV2ResponseMetaNormalizationOverlappedPercentage, LeakedCredentialTimeseriesGroupsV2ResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type LeakedCredentialTimeseriesGroupsV2ResponseMetaUnit struct {
	Name  string                                                 `json:"name,required"`
	Value string                                                 `json:"value,required"`
	JSON  leakedCredentialTimeseriesGroupsV2ResponseMetaUnitJSON `json:"-"`
}

// leakedCredentialTimeseriesGroupsV2ResponseMetaUnitJSON contains the JSON
// metadata for the struct [LeakedCredentialTimeseriesGroupsV2ResponseMetaUnit]
type leakedCredentialTimeseriesGroupsV2ResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialTimeseriesGroupsV2ResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialTimeseriesGroupsV2ResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type LeakedCredentialTimeseriesGroupsV2ResponseSerie0 struct {
	Timestamps  []time.Time                                          `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                                  `json:"-,extras"`
	JSON        leakedCredentialTimeseriesGroupsV2ResponseSerie0JSON `json:"-"`
}

// leakedCredentialTimeseriesGroupsV2ResponseSerie0JSON contains the JSON metadata
// for the struct [LeakedCredentialTimeseriesGroupsV2ResponseSerie0]
type leakedCredentialTimeseriesGroupsV2ResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialTimeseriesGroupsV2ResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialTimeseriesGroupsV2ResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type LeakedCredentialSummaryV2Params struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]LeakedCredentialSummaryV2ParamsBotClass] `query:"botClass"`
	// Filters results by compromised credential status (clean vs. compromised).
	Compromised param.Field[[]LeakedCredentialSummaryV2ParamsCompromised] `query:"compromised"`
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
	Format param.Field[LeakedCredentialSummaryV2ParamsFormat] `query:"format"`
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

// URLQuery serializes [LeakedCredentialSummaryV2Params]'s query parameters as
// `url.Values`.
func (r LeakedCredentialSummaryV2Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the attribute by which to group the results.
type LeakedCredentialSummaryV2ParamsDimension string

const (
	LeakedCredentialSummaryV2ParamsDimensionCompromised LeakedCredentialSummaryV2ParamsDimension = "COMPROMISED"
	LeakedCredentialSummaryV2ParamsDimensionBotClass    LeakedCredentialSummaryV2ParamsDimension = "BOT_CLASS"
)

func (r LeakedCredentialSummaryV2ParamsDimension) IsKnown() bool {
	switch r {
	case LeakedCredentialSummaryV2ParamsDimensionCompromised, LeakedCredentialSummaryV2ParamsDimensionBotClass:
		return true
	}
	return false
}

type LeakedCredentialSummaryV2ParamsBotClass string

const (
	LeakedCredentialSummaryV2ParamsBotClassLikelyAutomated LeakedCredentialSummaryV2ParamsBotClass = "LIKELY_AUTOMATED"
	LeakedCredentialSummaryV2ParamsBotClassLikelyHuman     LeakedCredentialSummaryV2ParamsBotClass = "LIKELY_HUMAN"
)

func (r LeakedCredentialSummaryV2ParamsBotClass) IsKnown() bool {
	switch r {
	case LeakedCredentialSummaryV2ParamsBotClassLikelyAutomated, LeakedCredentialSummaryV2ParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type LeakedCredentialSummaryV2ParamsCompromised string

const (
	LeakedCredentialSummaryV2ParamsCompromisedClean       LeakedCredentialSummaryV2ParamsCompromised = "CLEAN"
	LeakedCredentialSummaryV2ParamsCompromisedCompromised LeakedCredentialSummaryV2ParamsCompromised = "COMPROMISED"
)

func (r LeakedCredentialSummaryV2ParamsCompromised) IsKnown() bool {
	switch r {
	case LeakedCredentialSummaryV2ParamsCompromisedClean, LeakedCredentialSummaryV2ParamsCompromisedCompromised:
		return true
	}
	return false
}

// Format in which results will be returned.
type LeakedCredentialSummaryV2ParamsFormat string

const (
	LeakedCredentialSummaryV2ParamsFormatJson LeakedCredentialSummaryV2ParamsFormat = "JSON"
	LeakedCredentialSummaryV2ParamsFormatCsv  LeakedCredentialSummaryV2ParamsFormat = "CSV"
)

func (r LeakedCredentialSummaryV2ParamsFormat) IsKnown() bool {
	switch r {
	case LeakedCredentialSummaryV2ParamsFormatJson, LeakedCredentialSummaryV2ParamsFormatCsv:
		return true
	}
	return false
}

type LeakedCredentialSummaryV2ResponseEnvelope struct {
	Result  LeakedCredentialSummaryV2Response             `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    leakedCredentialSummaryV2ResponseEnvelopeJSON `json:"-"`
}

// leakedCredentialSummaryV2ResponseEnvelopeJSON contains the JSON metadata for the
// struct [LeakedCredentialSummaryV2ResponseEnvelope]
type leakedCredentialSummaryV2ResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialSummaryV2ResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialSummaryV2ResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LeakedCredentialTimeseriesGroupsV2Params struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[LeakedCredentialTimeseriesGroupsV2ParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]LeakedCredentialTimeseriesGroupsV2ParamsBotClass] `query:"botClass"`
	// Filters results by leaked credential check result.
	CheckResult param.Field[[]LeakedCredentialTimeseriesGroupsV2ParamsCheckResult] `query:"checkResult"`
	// Filters results by compromised credential status (clean vs. compromised).
	Compromised param.Field[[]LeakedCredentialTimeseriesGroupsV2ParamsCompromised] `query:"compromised"`
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
	Format param.Field[LeakedCredentialTimeseriesGroupsV2ParamsFormat] `query:"format"`
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
	Normalization param.Field[LeakedCredentialTimeseriesGroupsV2ParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [LeakedCredentialTimeseriesGroupsV2Params]'s query
// parameters as `url.Values`.
func (r LeakedCredentialTimeseriesGroupsV2Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the attribute by which to group the results.
type LeakedCredentialTimeseriesGroupsV2ParamsDimension string

const (
	LeakedCredentialTimeseriesGroupsV2ParamsDimensionCompromised LeakedCredentialTimeseriesGroupsV2ParamsDimension = "COMPROMISED"
	LeakedCredentialTimeseriesGroupsV2ParamsDimensionBotClass    LeakedCredentialTimeseriesGroupsV2ParamsDimension = "BOT_CLASS"
)

func (r LeakedCredentialTimeseriesGroupsV2ParamsDimension) IsKnown() bool {
	switch r {
	case LeakedCredentialTimeseriesGroupsV2ParamsDimensionCompromised, LeakedCredentialTimeseriesGroupsV2ParamsDimensionBotClass:
		return true
	}
	return false
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type LeakedCredentialTimeseriesGroupsV2ParamsAggInterval string

const (
	LeakedCredentialTimeseriesGroupsV2ParamsAggInterval15m LeakedCredentialTimeseriesGroupsV2ParamsAggInterval = "15m"
	LeakedCredentialTimeseriesGroupsV2ParamsAggInterval1h  LeakedCredentialTimeseriesGroupsV2ParamsAggInterval = "1h"
	LeakedCredentialTimeseriesGroupsV2ParamsAggInterval1d  LeakedCredentialTimeseriesGroupsV2ParamsAggInterval = "1d"
	LeakedCredentialTimeseriesGroupsV2ParamsAggInterval1w  LeakedCredentialTimeseriesGroupsV2ParamsAggInterval = "1w"
)

func (r LeakedCredentialTimeseriesGroupsV2ParamsAggInterval) IsKnown() bool {
	switch r {
	case LeakedCredentialTimeseriesGroupsV2ParamsAggInterval15m, LeakedCredentialTimeseriesGroupsV2ParamsAggInterval1h, LeakedCredentialTimeseriesGroupsV2ParamsAggInterval1d, LeakedCredentialTimeseriesGroupsV2ParamsAggInterval1w:
		return true
	}
	return false
}

type LeakedCredentialTimeseriesGroupsV2ParamsBotClass string

const (
	LeakedCredentialTimeseriesGroupsV2ParamsBotClassLikelyAutomated LeakedCredentialTimeseriesGroupsV2ParamsBotClass = "LIKELY_AUTOMATED"
	LeakedCredentialTimeseriesGroupsV2ParamsBotClassLikelyHuman     LeakedCredentialTimeseriesGroupsV2ParamsBotClass = "LIKELY_HUMAN"
)

func (r LeakedCredentialTimeseriesGroupsV2ParamsBotClass) IsKnown() bool {
	switch r {
	case LeakedCredentialTimeseriesGroupsV2ParamsBotClassLikelyAutomated, LeakedCredentialTimeseriesGroupsV2ParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type LeakedCredentialTimeseriesGroupsV2ParamsCheckResult string

const (
	LeakedCredentialTimeseriesGroupsV2ParamsCheckResultClean                     LeakedCredentialTimeseriesGroupsV2ParamsCheckResult = "CLEAN"
	LeakedCredentialTimeseriesGroupsV2ParamsCheckResultUsernameLeaked            LeakedCredentialTimeseriesGroupsV2ParamsCheckResult = "USERNAME_LEAKED"
	LeakedCredentialTimeseriesGroupsV2ParamsCheckResultUsernamePasswordSimilar   LeakedCredentialTimeseriesGroupsV2ParamsCheckResult = "USERNAME_PASSWORD_SIMILAR"
	LeakedCredentialTimeseriesGroupsV2ParamsCheckResultUsernameAndPasswordLeaked LeakedCredentialTimeseriesGroupsV2ParamsCheckResult = "USERNAME_AND_PASSWORD_LEAKED"
	LeakedCredentialTimeseriesGroupsV2ParamsCheckResultPasswordLeaked            LeakedCredentialTimeseriesGroupsV2ParamsCheckResult = "PASSWORD_LEAKED"
)

func (r LeakedCredentialTimeseriesGroupsV2ParamsCheckResult) IsKnown() bool {
	switch r {
	case LeakedCredentialTimeseriesGroupsV2ParamsCheckResultClean, LeakedCredentialTimeseriesGroupsV2ParamsCheckResultUsernameLeaked, LeakedCredentialTimeseriesGroupsV2ParamsCheckResultUsernamePasswordSimilar, LeakedCredentialTimeseriesGroupsV2ParamsCheckResultUsernameAndPasswordLeaked, LeakedCredentialTimeseriesGroupsV2ParamsCheckResultPasswordLeaked:
		return true
	}
	return false
}

type LeakedCredentialTimeseriesGroupsV2ParamsCompromised string

const (
	LeakedCredentialTimeseriesGroupsV2ParamsCompromisedClean       LeakedCredentialTimeseriesGroupsV2ParamsCompromised = "CLEAN"
	LeakedCredentialTimeseriesGroupsV2ParamsCompromisedCompromised LeakedCredentialTimeseriesGroupsV2ParamsCompromised = "COMPROMISED"
)

func (r LeakedCredentialTimeseriesGroupsV2ParamsCompromised) IsKnown() bool {
	switch r {
	case LeakedCredentialTimeseriesGroupsV2ParamsCompromisedClean, LeakedCredentialTimeseriesGroupsV2ParamsCompromisedCompromised:
		return true
	}
	return false
}

// Format in which results will be returned.
type LeakedCredentialTimeseriesGroupsV2ParamsFormat string

const (
	LeakedCredentialTimeseriesGroupsV2ParamsFormatJson LeakedCredentialTimeseriesGroupsV2ParamsFormat = "JSON"
	LeakedCredentialTimeseriesGroupsV2ParamsFormatCsv  LeakedCredentialTimeseriesGroupsV2ParamsFormat = "CSV"
)

func (r LeakedCredentialTimeseriesGroupsV2ParamsFormat) IsKnown() bool {
	switch r {
	case LeakedCredentialTimeseriesGroupsV2ParamsFormatJson, LeakedCredentialTimeseriesGroupsV2ParamsFormatCsv:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type LeakedCredentialTimeseriesGroupsV2ParamsNormalization string

const (
	LeakedCredentialTimeseriesGroupsV2ParamsNormalizationPercentageChange LeakedCredentialTimeseriesGroupsV2ParamsNormalization = "PERCENTAGE_CHANGE"
	LeakedCredentialTimeseriesGroupsV2ParamsNormalizationMin0Max          LeakedCredentialTimeseriesGroupsV2ParamsNormalization = "MIN0_MAX"
)

func (r LeakedCredentialTimeseriesGroupsV2ParamsNormalization) IsKnown() bool {
	switch r {
	case LeakedCredentialTimeseriesGroupsV2ParamsNormalizationPercentageChange, LeakedCredentialTimeseriesGroupsV2ParamsNormalizationMin0Max:
		return true
	}
	return false
}

type LeakedCredentialTimeseriesGroupsV2ResponseEnvelope struct {
	Result  LeakedCredentialTimeseriesGroupsV2Response             `json:"result,required"`
	Success bool                                                   `json:"success,required"`
	JSON    leakedCredentialTimeseriesGroupsV2ResponseEnvelopeJSON `json:"-"`
}

// leakedCredentialTimeseriesGroupsV2ResponseEnvelopeJSON contains the JSON
// metadata for the struct [LeakedCredentialTimeseriesGroupsV2ResponseEnvelope]
type leakedCredentialTimeseriesGroupsV2ResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialTimeseriesGroupsV2ResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialTimeseriesGroupsV2ResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
