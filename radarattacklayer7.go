// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarAttackLayer7Service contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarAttackLayer7Service] method
// instead.
type RadarAttackLayer7Service struct {
	Options          []option.RequestOption
	Summary          *RadarAttackLayer7SummaryService
	TimeseriesGroups *RadarAttackLayer7TimeseriesGroupService
	Top              *RadarAttackLayer7TopService
}

// NewRadarAttackLayer7Service generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAttackLayer7Service(opts ...option.RequestOption) (r *RadarAttackLayer7Service) {
	r = &RadarAttackLayer7Service{}
	r.Options = opts
	r.Summary = NewRadarAttackLayer7SummaryService(opts...)
	r.TimeseriesGroups = NewRadarAttackLayer7TimeseriesGroupService(opts...)
	r.Top = NewRadarAttackLayer7TopService(opts...)
	return
}

// Get a timeseries of Layer 7 attacks. Values represent HTTP requests and are
// normalized using min-max by default.
func (r *RadarAttackLayer7Service) Timeseries(ctx context.Context, query RadarAttackLayer7TimeseriesParams, opts ...option.RequestOption) (res *RadarAttackLayer7TimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7TimeseriesResponseEnvelope
	path := "radar/attacks/layer7/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAttackLayer7TimeseriesResponse struct {
	Meta   RadarAttackLayer7TimeseriesResponseMeta   `json:"meta,required"`
	Serie0 RadarAttackLayer7TimeseriesResponseSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7TimeseriesResponseJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer7TimeseriesResponse]
type radarAttackLayer7TimeseriesResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesResponseMeta struct {
	AggInterval    string                                                `json:"aggInterval,required"`
	DateRange      []RadarAttackLayer7TimeseriesResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    time.Time                                             `json:"lastUpdated,required" format:"date-time"`
	ConfidenceInfo RadarAttackLayer7TimeseriesResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7TimeseriesResponseMetaJSON           `json:"-"`
}

// radarAttackLayer7TimeseriesResponseMetaJSON contains the JSON metadata for the
// struct [RadarAttackLayer7TimeseriesResponseMeta]
type radarAttackLayer7TimeseriesResponseMetaJSON struct {
	AggInterval    apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                            `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7TimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7TimeseriesResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarAttackLayer7TimeseriesResponseMetaDateRange]
type radarAttackLayer7TimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                             `json:"level"`
	JSON        radarAttackLayer7TimeseriesResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7TimeseriesResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarAttackLayer7TimeseriesResponseMetaConfidenceInfo]
type radarAttackLayer7TimeseriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                              `json:"dataSource,required"`
	Description     string                                                              `json:"description,required"`
	EventType       string                                                              `json:"eventType,required"`
	IsInstantaneous interface{}                                                         `json:"isInstantaneous,required"`
	EndTime         time.Time                                                           `json:"endTime" format:"date-time"`
	LinkedURL       string                                                              `json:"linkedUrl"`
	StartTime       time.Time                                                           `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	EndTime         apijson.Field
	LinkedURL       apijson.Field
	StartTime       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesResponseSerie0 struct {
	Timestamps []time.Time                                   `json:"timestamps,required" format:"date-time"`
	Values     []string                                      `json:"values,required"`
	JSON       radarAttackLayer7TimeseriesResponseSerie0JSON `json:"-"`
}

// radarAttackLayer7TimeseriesResponseSerie0JSON contains the JSON metadata for the
// struct [RadarAttackLayer7TimeseriesResponseSerie0]
type radarAttackLayer7TimeseriesResponseSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7TimeseriesParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of L7 attack types.
	Attack param.Field[[]RadarAttackLayer7TimeseriesParamsAttack] `query:"attack"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7TimeseriesParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7TimeseriesParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer7TimeseriesParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [RadarAttackLayer7TimeseriesParams]'s query parameters as
// `url.Values`.
func (r RadarAttackLayer7TimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7TimeseriesParamsAggInterval string

const (
	RadarAttackLayer7TimeseriesParamsAggInterval15m RadarAttackLayer7TimeseriesParamsAggInterval = "15m"
	RadarAttackLayer7TimeseriesParamsAggInterval1h  RadarAttackLayer7TimeseriesParamsAggInterval = "1h"
	RadarAttackLayer7TimeseriesParamsAggInterval1d  RadarAttackLayer7TimeseriesParamsAggInterval = "1d"
	RadarAttackLayer7TimeseriesParamsAggInterval1w  RadarAttackLayer7TimeseriesParamsAggInterval = "1w"
)

type RadarAttackLayer7TimeseriesParamsAttack string

const (
	RadarAttackLayer7TimeseriesParamsAttackDDOS               RadarAttackLayer7TimeseriesParamsAttack = "DDOS"
	RadarAttackLayer7TimeseriesParamsAttackWAF                RadarAttackLayer7TimeseriesParamsAttack = "WAF"
	RadarAttackLayer7TimeseriesParamsAttackBotManagement      RadarAttackLayer7TimeseriesParamsAttack = "BOT_MANAGEMENT"
	RadarAttackLayer7TimeseriesParamsAttackAccessRules        RadarAttackLayer7TimeseriesParamsAttack = "ACCESS_RULES"
	RadarAttackLayer7TimeseriesParamsAttackIPReputation       RadarAttackLayer7TimeseriesParamsAttack = "IP_REPUTATION"
	RadarAttackLayer7TimeseriesParamsAttackAPIShield          RadarAttackLayer7TimeseriesParamsAttack = "API_SHIELD"
	RadarAttackLayer7TimeseriesParamsAttackDataLossPrevention RadarAttackLayer7TimeseriesParamsAttack = "DATA_LOSS_PREVENTION"
)

type RadarAttackLayer7TimeseriesParamsDateRange string

const (
	RadarAttackLayer7TimeseriesParamsDateRange1d         RadarAttackLayer7TimeseriesParamsDateRange = "1d"
	RadarAttackLayer7TimeseriesParamsDateRange2d         RadarAttackLayer7TimeseriesParamsDateRange = "2d"
	RadarAttackLayer7TimeseriesParamsDateRange7d         RadarAttackLayer7TimeseriesParamsDateRange = "7d"
	RadarAttackLayer7TimeseriesParamsDateRange14d        RadarAttackLayer7TimeseriesParamsDateRange = "14d"
	RadarAttackLayer7TimeseriesParamsDateRange28d        RadarAttackLayer7TimeseriesParamsDateRange = "28d"
	RadarAttackLayer7TimeseriesParamsDateRange12w        RadarAttackLayer7TimeseriesParamsDateRange = "12w"
	RadarAttackLayer7TimeseriesParamsDateRange24w        RadarAttackLayer7TimeseriesParamsDateRange = "24w"
	RadarAttackLayer7TimeseriesParamsDateRange52w        RadarAttackLayer7TimeseriesParamsDateRange = "52w"
	RadarAttackLayer7TimeseriesParamsDateRange1dControl  RadarAttackLayer7TimeseriesParamsDateRange = "1dControl"
	RadarAttackLayer7TimeseriesParamsDateRange2dControl  RadarAttackLayer7TimeseriesParamsDateRange = "2dControl"
	RadarAttackLayer7TimeseriesParamsDateRange7dControl  RadarAttackLayer7TimeseriesParamsDateRange = "7dControl"
	RadarAttackLayer7TimeseriesParamsDateRange14dControl RadarAttackLayer7TimeseriesParamsDateRange = "14dControl"
	RadarAttackLayer7TimeseriesParamsDateRange28dControl RadarAttackLayer7TimeseriesParamsDateRange = "28dControl"
	RadarAttackLayer7TimeseriesParamsDateRange12wControl RadarAttackLayer7TimeseriesParamsDateRange = "12wControl"
	RadarAttackLayer7TimeseriesParamsDateRange24wControl RadarAttackLayer7TimeseriesParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7TimeseriesParamsFormat string

const (
	RadarAttackLayer7TimeseriesParamsFormatJson RadarAttackLayer7TimeseriesParamsFormat = "JSON"
	RadarAttackLayer7TimeseriesParamsFormatCsv  RadarAttackLayer7TimeseriesParamsFormat = "CSV"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7TimeseriesParamsNormalization string

const (
	RadarAttackLayer7TimeseriesParamsNormalizationPercentageChange RadarAttackLayer7TimeseriesParamsNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer7TimeseriesParamsNormalizationMin0Max          RadarAttackLayer7TimeseriesParamsNormalization = "MIN0_MAX"
)

type RadarAttackLayer7TimeseriesResponseEnvelope struct {
	Result  RadarAttackLayer7TimeseriesResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarAttackLayer7TimeseriesResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer7TimeseriesResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarAttackLayer7TimeseriesResponseEnvelope]
type radarAttackLayer7TimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
