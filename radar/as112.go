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

// AS112Service contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAS112Service] method instead.
type AS112Service struct {
	Options          []option.RequestOption
	Summary          *AS112SummaryService
	TimeseriesGroups *AS112TimeseriesGroupService
	Top              *AS112TopService
}

// NewAS112Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAS112Service(opts ...option.RequestOption) (r *AS112Service) {
	r = &AS112Service{}
	r.Options = opts
	r.Summary = NewAS112SummaryService(opts...)
	r.TimeseriesGroups = NewAS112TimeseriesGroupService(opts...)
	r.Top = NewAS112TopService(opts...)
	return
}

// Get AS112 queries change over time.
func (r *AS112Service) Timeseries(ctx context.Context, query AS112TimeseriesParams, opts ...option.RequestOption) (res *AS112TimeseriesResponse, err error) {
	var env AS112TimeseriesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/as112/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AS112TimeseriesResponse struct {
	Meta   AS112TimeseriesResponseMeta   `json:"meta,required"`
	Serie0 AS112TimeseriesResponseSerie0 `json:"serie_0,required"`
	JSON   as112TimeseriesResponseJSON   `json:"-"`
}

// as112TimeseriesResponseJSON contains the JSON metadata for the struct
// [AS112TimeseriesResponse]
type as112TimeseriesResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesResponseMeta struct {
	AggInterval    string                                    `json:"aggInterval,required"`
	DateRange      []AS112TimeseriesResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    time.Time                                 `json:"lastUpdated,required" format:"date-time"`
	ConfidenceInfo AS112TimeseriesResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           as112TimeseriesResponseMetaJSON           `json:"-"`
}

// as112TimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [AS112TimeseriesResponseMeta]
type as112TimeseriesResponseMetaJSON struct {
	AggInterval    apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AS112TimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                `json:"startTime,required" format:"date-time"`
	JSON      as112TimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// as112TimeseriesResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [AS112TimeseriesResponseMetaDateRange]
type as112TimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesResponseMetaConfidenceInfo struct {
	Annotations []AS112TimeseriesResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                 `json:"level"`
	JSON        as112TimeseriesResponseMetaConfidenceInfoJSON         `json:"-"`
}

// as112TimeseriesResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [AS112TimeseriesResponseMetaConfidenceInfo]
type as112TimeseriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                  `json:"dataSource,required"`
	Description     string                                                  `json:"description,required"`
	EventType       string                                                  `json:"eventType,required"`
	IsInstantaneous bool                                                    `json:"isInstantaneous,required"`
	EndTime         time.Time                                               `json:"endTime" format:"date-time"`
	LinkedURL       string                                                  `json:"linkedUrl"`
	StartTime       time.Time                                               `json:"startTime" format:"date-time"`
	JSON            as112TimeseriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// as112TimeseriesResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [AS112TimeseriesResponseMetaConfidenceInfoAnnotation]
type as112TimeseriesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AS112TimeseriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesResponseSerie0 struct {
	Timestamps []time.Time                       `json:"timestamps,required" format:"date-time"`
	Values     []string                          `json:"values,required"`
	JSON       as112TimeseriesResponseSerie0JSON `json:"-"`
}

// as112TimeseriesResponseSerie0JSON contains the JSON metadata for the struct
// [AS112TimeseriesResponseSerie0]
type as112TimeseriesResponseSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AS112TimeseriesParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AS112TimeseriesParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AS112TimeseriesParams]'s query parameters as `url.Values`.
func (r AS112TimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AS112TimeseriesParamsAggInterval string

const (
	AS112TimeseriesParamsAggInterval15m AS112TimeseriesParamsAggInterval = "15m"
	AS112TimeseriesParamsAggInterval1h  AS112TimeseriesParamsAggInterval = "1h"
	AS112TimeseriesParamsAggInterval1d  AS112TimeseriesParamsAggInterval = "1d"
	AS112TimeseriesParamsAggInterval1w  AS112TimeseriesParamsAggInterval = "1w"
)

func (r AS112TimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case AS112TimeseriesParamsAggInterval15m, AS112TimeseriesParamsAggInterval1h, AS112TimeseriesParamsAggInterval1d, AS112TimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Format results are returned in.
type AS112TimeseriesParamsFormat string

const (
	AS112TimeseriesParamsFormatJson AS112TimeseriesParamsFormat = "JSON"
	AS112TimeseriesParamsFormatCsv  AS112TimeseriesParamsFormat = "CSV"
)

func (r AS112TimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case AS112TimeseriesParamsFormatJson, AS112TimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type AS112TimeseriesResponseEnvelope struct {
	Result  AS112TimeseriesResponse             `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    as112TimeseriesResponseEnvelopeJSON `json:"-"`
}

// as112TimeseriesResponseEnvelopeJSON contains the JSON metadata for the struct
// [AS112TimeseriesResponseEnvelope]
type as112TimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
