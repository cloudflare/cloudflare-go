// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// RadarBGPService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewRadarBGPService] method instead.
type RadarBGPService struct {
	Options []option.RequestOption
	Leaks   *RadarBGPLeakService
	Top     *RadarBGPTopService
	Hijacks *RadarBGPHijackService
	Routes  *RadarBGPRouteService
}

// NewRadarBGPService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarBGPService(opts ...option.RequestOption) (r *RadarBGPService) {
	r = &RadarBGPService{}
	r.Options = opts
	r.Leaks = NewRadarBGPLeakService(opts...)
	r.Top = NewRadarBGPTopService(opts...)
	r.Hijacks = NewRadarBGPHijackService(opts...)
	r.Routes = NewRadarBGPRouteService(opts...)
	return
}

// Gets BGP updates change over time. Raw values are returned. When requesting
// updates of an autonomous system (AS), only BGP updates of type announcement are
// returned.
func (r *RadarBGPService) Timeseries(ctx context.Context, query RadarBGPTimeseriesParams, opts ...option.RequestOption) (res *RadarBGPTimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarBGPTimeseriesResponseEnvelope
	path := "radar/bgp/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarBGPTimeseriesResponse struct {
	Meta   RadarBGPTimeseriesResponseMeta   `json:"meta,required"`
	Serie0 RadarBGPTimeseriesResponseSerie0 `json:"serie_0,required"`
	JSON   radarBGPTimeseriesResponseJSON   `json:"-"`
}

// radarBGPTimeseriesResponseJSON contains the JSON metadata for the struct
// [RadarBGPTimeseriesResponse]
type radarBGPTimeseriesResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBGPTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarBGPTimeseriesResponseMeta struct {
	AggInterval    string                                       `json:"aggInterval,required"`
	DateRange      []RadarBGPTimeseriesResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    time.Time                                    `json:"lastUpdated,required" format:"date-time"`
	ConfidenceInfo RadarBGPTimeseriesResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarBGPTimeseriesResponseMetaJSON           `json:"-"`
}

// radarBGPTimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [RadarBGPTimeseriesResponseMeta]
type radarBGPTimeseriesResponseMetaJSON struct {
	AggInterval    apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarBGPTimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBGPTimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarBGPTimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                   `json:"startTime,required" format:"date-time"`
	JSON      radarBGPTimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// radarBGPTimeseriesResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [RadarBGPTimeseriesResponseMetaDateRange]
type radarBGPTimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBGPTimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarBGPTimeseriesResponseMetaConfidenceInfo struct {
	Annotations []RadarBGPTimeseriesResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                    `json:"level"`
	JSON        radarBGPTimeseriesResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarBGPTimeseriesResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [RadarBGPTimeseriesResponseMetaConfidenceInfo]
type radarBGPTimeseriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTimeseriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBGPTimeseriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarBGPTimeseriesResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                     `json:"dataSource,required"`
	Description     string                                                     `json:"description,required"`
	EventType       string                                                     `json:"eventType,required"`
	IsInstantaneous interface{}                                                `json:"isInstantaneous,required"`
	EndTime         time.Time                                                  `json:"endTime" format:"date-time"`
	LinkedURL       string                                                     `json:"linkedUrl"`
	StartTime       time.Time                                                  `json:"startTime" format:"date-time"`
	JSON            radarBGPTimeseriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarBGPTimeseriesResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [RadarBGPTimeseriesResponseMetaConfidenceInfoAnnotation]
type radarBGPTimeseriesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarBGPTimeseriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBGPTimeseriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarBGPTimeseriesResponseSerie0 struct {
	Timestamps []time.Time                          `json:"timestamps,required" format:"date-time"`
	Values     []string                             `json:"values,required"`
	JSON       radarBGPTimeseriesResponseSerie0JSON `json:"-"`
}

// radarBGPTimeseriesResponseSerie0JSON contains the JSON metadata for the struct
// [RadarBGPTimeseriesResponseSerie0]
type radarBGPTimeseriesResponseSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTimeseriesResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBGPTimeseriesResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarBGPTimeseriesParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarBGPTimeseriesParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarBGPTimeseriesParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarBGPTimeseriesParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of BGP network prefixes.
	Prefix param.Field[[]string] `query:"prefix"`
	// Array of BGP update types.
	UpdateType param.Field[[]RadarBGPTimeseriesParamsUpdateType] `query:"updateType"`
}

// URLQuery serializes [RadarBGPTimeseriesParams]'s query parameters as
// `url.Values`.
func (r RadarBGPTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarBGPTimeseriesParamsAggInterval string

const (
	RadarBGPTimeseriesParamsAggInterval15m RadarBGPTimeseriesParamsAggInterval = "15m"
	RadarBGPTimeseriesParamsAggInterval1h  RadarBGPTimeseriesParamsAggInterval = "1h"
	RadarBGPTimeseriesParamsAggInterval1d  RadarBGPTimeseriesParamsAggInterval = "1d"
	RadarBGPTimeseriesParamsAggInterval1w  RadarBGPTimeseriesParamsAggInterval = "1w"
)

type RadarBGPTimeseriesParamsDateRange string

const (
	RadarBGPTimeseriesParamsDateRange1d         RadarBGPTimeseriesParamsDateRange = "1d"
	RadarBGPTimeseriesParamsDateRange2d         RadarBGPTimeseriesParamsDateRange = "2d"
	RadarBGPTimeseriesParamsDateRange7d         RadarBGPTimeseriesParamsDateRange = "7d"
	RadarBGPTimeseriesParamsDateRange14d        RadarBGPTimeseriesParamsDateRange = "14d"
	RadarBGPTimeseriesParamsDateRange28d        RadarBGPTimeseriesParamsDateRange = "28d"
	RadarBGPTimeseriesParamsDateRange12w        RadarBGPTimeseriesParamsDateRange = "12w"
	RadarBGPTimeseriesParamsDateRange24w        RadarBGPTimeseriesParamsDateRange = "24w"
	RadarBGPTimeseriesParamsDateRange52w        RadarBGPTimeseriesParamsDateRange = "52w"
	RadarBGPTimeseriesParamsDateRange1dControl  RadarBGPTimeseriesParamsDateRange = "1dControl"
	RadarBGPTimeseriesParamsDateRange2dControl  RadarBGPTimeseriesParamsDateRange = "2dControl"
	RadarBGPTimeseriesParamsDateRange7dControl  RadarBGPTimeseriesParamsDateRange = "7dControl"
	RadarBGPTimeseriesParamsDateRange14dControl RadarBGPTimeseriesParamsDateRange = "14dControl"
	RadarBGPTimeseriesParamsDateRange28dControl RadarBGPTimeseriesParamsDateRange = "28dControl"
	RadarBGPTimeseriesParamsDateRange12wControl RadarBGPTimeseriesParamsDateRange = "12wControl"
	RadarBGPTimeseriesParamsDateRange24wControl RadarBGPTimeseriesParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarBGPTimeseriesParamsFormat string

const (
	RadarBGPTimeseriesParamsFormatJson RadarBGPTimeseriesParamsFormat = "JSON"
	RadarBGPTimeseriesParamsFormatCsv  RadarBGPTimeseriesParamsFormat = "CSV"
)

type RadarBGPTimeseriesParamsUpdateType string

const (
	RadarBGPTimeseriesParamsUpdateTypeAnnouncement RadarBGPTimeseriesParamsUpdateType = "ANNOUNCEMENT"
	RadarBGPTimeseriesParamsUpdateTypeWithdrawal   RadarBGPTimeseriesParamsUpdateType = "WITHDRAWAL"
)

type RadarBGPTimeseriesResponseEnvelope struct {
	Result  RadarBGPTimeseriesResponse             `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    radarBGPTimeseriesResponseEnvelopeJSON `json:"-"`
}

// radarBGPTimeseriesResponseEnvelopeJSON contains the JSON metadata for the struct
// [RadarBGPTimeseriesResponseEnvelope]
type radarBGPTimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBGPTimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
