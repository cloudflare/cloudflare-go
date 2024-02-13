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

// RadarBGPTimeseryService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarBGPTimeseryService] method
// instead.
type RadarBGPTimeseryService struct {
	Options []option.RequestOption
}

// NewRadarBGPTimeseryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarBGPTimeseryService(opts ...option.RequestOption) (r *RadarBGPTimeseryService) {
	r = &RadarBGPTimeseryService{}
	r.Options = opts
	return
}

// Gets BGP updates change over time. Raw values are returned. When requesting
// updates of an autonomous system (AS), only BGP updates of type announcement are
// returned.
func (r *RadarBGPTimeseryService) List(ctx context.Context, query RadarBGPTimeseryListParams, opts ...option.RequestOption) (res *RadarBGPTimeseryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarBGPTimeseryListResponseEnvelope
	path := "radar/bgp/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarBGPTimeseryListResponse struct {
	Meta   RadarBGPTimeseryListResponseMeta   `json:"meta,required"`
	Serie0 RadarBGPTimeseryListResponseSerie0 `json:"serie_0,required"`
	JSON   radarBGPTimeseryListResponseJSON   `json:"-"`
}

// radarBGPTimeseryListResponseJSON contains the JSON metadata for the struct
// [RadarBGPTimeseryListResponse]
type radarBGPTimeseryListResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTimeseryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPTimeseryListResponseMeta struct {
	AggInterval    string                                         `json:"aggInterval,required"`
	DateRange      []RadarBGPTimeseryListResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    time.Time                                      `json:"lastUpdated,required" format:"date-time"`
	ConfidenceInfo RadarBGPTimeseryListResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarBGPTimeseryListResponseMetaJSON           `json:"-"`
}

// radarBGPTimeseryListResponseMetaJSON contains the JSON metadata for the struct
// [RadarBGPTimeseryListResponseMeta]
type radarBGPTimeseryListResponseMetaJSON struct {
	AggInterval    apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarBGPTimeseryListResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPTimeseryListResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                     `json:"startTime,required" format:"date-time"`
	JSON      radarBGPTimeseryListResponseMetaDateRangeJSON `json:"-"`
}

// radarBGPTimeseryListResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [RadarBGPTimeseryListResponseMetaDateRange]
type radarBGPTimeseryListResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTimeseryListResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPTimeseryListResponseMetaConfidenceInfo struct {
	Annotations []RadarBGPTimeseryListResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                      `json:"level"`
	JSON        radarBGPTimeseryListResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarBGPTimeseryListResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [RadarBGPTimeseryListResponseMetaConfidenceInfo]
type radarBGPTimeseryListResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTimeseryListResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPTimeseryListResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                       `json:"dataSource,required"`
	Description     string                                                       `json:"description,required"`
	EventType       string                                                       `json:"eventType,required"`
	IsInstantaneous interface{}                                                  `json:"isInstantaneous,required"`
	EndTime         time.Time                                                    `json:"endTime" format:"date-time"`
	LinkedURL       string                                                       `json:"linkedUrl"`
	StartTime       time.Time                                                    `json:"startTime" format:"date-time"`
	JSON            radarBGPTimeseryListResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarBGPTimeseryListResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [RadarBGPTimeseryListResponseMetaConfidenceInfoAnnotation]
type radarBGPTimeseryListResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarBGPTimeseryListResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPTimeseryListResponseSerie0 struct {
	Timestamps []time.Time                            `json:"timestamps,required" format:"date-time"`
	Values     []string                               `json:"values,required"`
	JSON       radarBGPTimeseryListResponseSerie0JSON `json:"-"`
}

// radarBGPTimeseryListResponseSerie0JSON contains the JSON metadata for the struct
// [RadarBGPTimeseryListResponseSerie0]
type radarBGPTimeseryListResponseSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTimeseryListResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPTimeseryListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarBGPTimeseryListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarBGPTimeseryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarBGPTimeseryListParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of BGP network prefixes.
	Prefix param.Field[[]string] `query:"prefix"`
	// Array of BGP update types.
	UpdateType param.Field[[]RadarBGPTimeseryListParamsUpdateType] `query:"updateType"`
}

// URLQuery serializes [RadarBGPTimeseryListParams]'s query parameters as
// `url.Values`.
func (r RadarBGPTimeseryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarBGPTimeseryListParamsAggInterval string

const (
	RadarBGPTimeseryListParamsAggInterval15m RadarBGPTimeseryListParamsAggInterval = "15m"
	RadarBGPTimeseryListParamsAggInterval1h  RadarBGPTimeseryListParamsAggInterval = "1h"
	RadarBGPTimeseryListParamsAggInterval1d  RadarBGPTimeseryListParamsAggInterval = "1d"
	RadarBGPTimeseryListParamsAggInterval1w  RadarBGPTimeseryListParamsAggInterval = "1w"
)

type RadarBGPTimeseryListParamsDateRange string

const (
	RadarBGPTimeseryListParamsDateRange1d         RadarBGPTimeseryListParamsDateRange = "1d"
	RadarBGPTimeseryListParamsDateRange2d         RadarBGPTimeseryListParamsDateRange = "2d"
	RadarBGPTimeseryListParamsDateRange7d         RadarBGPTimeseryListParamsDateRange = "7d"
	RadarBGPTimeseryListParamsDateRange14d        RadarBGPTimeseryListParamsDateRange = "14d"
	RadarBGPTimeseryListParamsDateRange28d        RadarBGPTimeseryListParamsDateRange = "28d"
	RadarBGPTimeseryListParamsDateRange12w        RadarBGPTimeseryListParamsDateRange = "12w"
	RadarBGPTimeseryListParamsDateRange24w        RadarBGPTimeseryListParamsDateRange = "24w"
	RadarBGPTimeseryListParamsDateRange52w        RadarBGPTimeseryListParamsDateRange = "52w"
	RadarBGPTimeseryListParamsDateRange1dControl  RadarBGPTimeseryListParamsDateRange = "1dControl"
	RadarBGPTimeseryListParamsDateRange2dControl  RadarBGPTimeseryListParamsDateRange = "2dControl"
	RadarBGPTimeseryListParamsDateRange7dControl  RadarBGPTimeseryListParamsDateRange = "7dControl"
	RadarBGPTimeseryListParamsDateRange14dControl RadarBGPTimeseryListParamsDateRange = "14dControl"
	RadarBGPTimeseryListParamsDateRange28dControl RadarBGPTimeseryListParamsDateRange = "28dControl"
	RadarBGPTimeseryListParamsDateRange12wControl RadarBGPTimeseryListParamsDateRange = "12wControl"
	RadarBGPTimeseryListParamsDateRange24wControl RadarBGPTimeseryListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarBGPTimeseryListParamsFormat string

const (
	RadarBGPTimeseryListParamsFormatJson RadarBGPTimeseryListParamsFormat = "JSON"
	RadarBGPTimeseryListParamsFormatCsv  RadarBGPTimeseryListParamsFormat = "CSV"
)

type RadarBGPTimeseryListParamsUpdateType string

const (
	RadarBGPTimeseryListParamsUpdateTypeAnnouncement RadarBGPTimeseryListParamsUpdateType = "ANNOUNCEMENT"
	RadarBGPTimeseryListParamsUpdateTypeWithdrawal   RadarBGPTimeseryListParamsUpdateType = "WITHDRAWAL"
)

type RadarBGPTimeseryListResponseEnvelope struct {
	Result  RadarBGPTimeseryListResponse             `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    radarBGPTimeseryListResponseEnvelopeJSON `json:"-"`
}

// radarBGPTimeseryListResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarBGPTimeseryListResponseEnvelope]
type radarBGPTimeseryListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTimeseryListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
