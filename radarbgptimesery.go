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

// RadarBgpTimeseryService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarBgpTimeseryService] method
// instead.
type RadarBgpTimeseryService struct {
	Options []option.RequestOption
}

// NewRadarBgpTimeseryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarBgpTimeseryService(opts ...option.RequestOption) (r *RadarBgpTimeseryService) {
	r = &RadarBgpTimeseryService{}
	r.Options = opts
	return
}

// Gets BGP updates change over time. Raw values are returned. When requesting
// updates of an autonomous system (AS), only BGP updates of type announcement are
// returned.
func (r *RadarBgpTimeseryService) List(ctx context.Context, query RadarBgpTimeseryListParams, opts ...option.RequestOption) (res *RadarBgpTimeseryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/bgp/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarBgpTimeseryListResponse struct {
	Result  RadarBgpTimeseryListResponseResult `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    radarBgpTimeseryListResponseJSON   `json:"-"`
}

// radarBgpTimeseryListResponseJSON contains the JSON metadata for the struct
// [RadarBgpTimeseryListResponse]
type radarBgpTimeseryListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTimeseryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpTimeseryListResponseResult struct {
	Meta   RadarBgpTimeseryListResponseResultMeta   `json:"meta,required"`
	Serie0 RadarBgpTimeseryListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarBgpTimeseryListResponseResultJSON   `json:"-"`
}

// radarBgpTimeseryListResponseResultJSON contains the JSON metadata for the struct
// [RadarBgpTimeseryListResponseResult]
type radarBgpTimeseryListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTimeseryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpTimeseryListResponseResultMeta struct {
	AggInterval    string                                               `json:"aggInterval,required"`
	DateRange      []RadarBgpTimeseryListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    time.Time                                            `json:"lastUpdated,required" format:"date-time"`
	ConfidenceInfo RadarBgpTimeseryListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarBgpTimeseryListResponseResultMetaJSON           `json:"-"`
}

// radarBgpTimeseryListResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarBgpTimeseryListResponseResultMeta]
type radarBgpTimeseryListResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarBgpTimeseryListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpTimeseryListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                           `json:"startTime,required" format:"date-time"`
	JSON      radarBgpTimeseryListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarBgpTimeseryListResponseResultMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarBgpTimeseryListResponseResultMetaDateRange]
type radarBgpTimeseryListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTimeseryListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpTimeseryListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarBgpTimeseryListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                            `json:"level"`
	JSON        radarBgpTimeseryListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarBgpTimeseryListResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarBgpTimeseryListResponseResultMetaConfidenceInfo]
type radarBgpTimeseryListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTimeseryListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpTimeseryListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                             `json:"dataSource,required"`
	Description     string                                                             `json:"description,required"`
	EventType       string                                                             `json:"eventType,required"`
	IsInstantaneous interface{}                                                        `json:"isInstantaneous,required"`
	EndTime         time.Time                                                          `json:"endTime" format:"date-time"`
	LinkedURL       string                                                             `json:"linkedUrl"`
	StartTime       time.Time                                                          `json:"startTime" format:"date-time"`
	JSON            radarBgpTimeseryListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarBgpTimeseryListResponseResultMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarBgpTimeseryListResponseResultMetaConfidenceInfoAnnotation]
type radarBgpTimeseryListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarBgpTimeseryListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpTimeseryListResponseResultSerie0 struct {
	Timestamps []time.Time                                  `json:"timestamps,required" format:"date-time"`
	Values     []string                                     `json:"values,required"`
	JSON       radarBgpTimeseryListResponseResultSerie0JSON `json:"-"`
}

// radarBgpTimeseryListResponseResultSerie0JSON contains the JSON metadata for the
// struct [RadarBgpTimeseryListResponseResultSerie0]
type radarBgpTimeseryListResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTimeseryListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpTimeseryListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarBgpTimeseryListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarBgpTimeseryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarBgpTimeseryListParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of BGP network prefixes.
	Prefix param.Field[[]string] `query:"prefix"`
	// Array of BGP update types.
	UpdateType param.Field[[]RadarBgpTimeseryListParamsUpdateType] `query:"updateType"`
}

// URLQuery serializes [RadarBgpTimeseryListParams]'s query parameters as
// `url.Values`.
func (r RadarBgpTimeseryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarBgpTimeseryListParamsAggInterval string

const (
	RadarBgpTimeseryListParamsAggInterval15m RadarBgpTimeseryListParamsAggInterval = "15m"
	RadarBgpTimeseryListParamsAggInterval1h  RadarBgpTimeseryListParamsAggInterval = "1h"
	RadarBgpTimeseryListParamsAggInterval1d  RadarBgpTimeseryListParamsAggInterval = "1d"
	RadarBgpTimeseryListParamsAggInterval1w  RadarBgpTimeseryListParamsAggInterval = "1w"
)

type RadarBgpTimeseryListParamsDateRange string

const (
	RadarBgpTimeseryListParamsDateRange1d         RadarBgpTimeseryListParamsDateRange = "1d"
	RadarBgpTimeseryListParamsDateRange2d         RadarBgpTimeseryListParamsDateRange = "2d"
	RadarBgpTimeseryListParamsDateRange7d         RadarBgpTimeseryListParamsDateRange = "7d"
	RadarBgpTimeseryListParamsDateRange14d        RadarBgpTimeseryListParamsDateRange = "14d"
	RadarBgpTimeseryListParamsDateRange28d        RadarBgpTimeseryListParamsDateRange = "28d"
	RadarBgpTimeseryListParamsDateRange12w        RadarBgpTimeseryListParamsDateRange = "12w"
	RadarBgpTimeseryListParamsDateRange24w        RadarBgpTimeseryListParamsDateRange = "24w"
	RadarBgpTimeseryListParamsDateRange52w        RadarBgpTimeseryListParamsDateRange = "52w"
	RadarBgpTimeseryListParamsDateRange1dControl  RadarBgpTimeseryListParamsDateRange = "1dControl"
	RadarBgpTimeseryListParamsDateRange2dControl  RadarBgpTimeseryListParamsDateRange = "2dControl"
	RadarBgpTimeseryListParamsDateRange7dControl  RadarBgpTimeseryListParamsDateRange = "7dControl"
	RadarBgpTimeseryListParamsDateRange14dControl RadarBgpTimeseryListParamsDateRange = "14dControl"
	RadarBgpTimeseryListParamsDateRange28dControl RadarBgpTimeseryListParamsDateRange = "28dControl"
	RadarBgpTimeseryListParamsDateRange12wControl RadarBgpTimeseryListParamsDateRange = "12wControl"
	RadarBgpTimeseryListParamsDateRange24wControl RadarBgpTimeseryListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarBgpTimeseryListParamsFormat string

const (
	RadarBgpTimeseryListParamsFormatJson RadarBgpTimeseryListParamsFormat = "JSON"
	RadarBgpTimeseryListParamsFormatCsv  RadarBgpTimeseryListParamsFormat = "CSV"
)

type RadarBgpTimeseryListParamsUpdateType string

const (
	RadarBgpTimeseryListParamsUpdateTypeAnnouncement RadarBgpTimeseryListParamsUpdateType = "ANNOUNCEMENT"
	RadarBgpTimeseryListParamsUpdateTypeWithdrawal   RadarBgpTimeseryListParamsUpdateType = "WITHDRAWAL"
)
