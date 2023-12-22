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

// RadarAttackLayer7TimeseriesGroupService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer7TimeseriesGroupService] method instead.
type RadarAttackLayer7TimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer7TimeseriesGroupService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer7TimeseriesGroupService(opts ...option.RequestOption) (r *RadarAttackLayer7TimeseriesGroupService) {
	r = &RadarAttackLayer7TimeseriesGroupService{}
	r.Options = opts
	return
}

// Get percentage of what type of mitigation techniques are used to block layer 7
// attacks, over time.
func (r *RadarAttackLayer7TimeseriesGroupService) List(ctx context.Context, query RadarAttackLayer7TimeseriesGroupListParams, opts ...option.RequestOption) (res *RadarAttackLayer7TimeseriesGroupListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer7TimeseriesGroupListResponse struct {
	Result  RadarAttackLayer7TimeseriesGroupListResponseResult `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    radarAttackLayer7TimeseriesGroupListResponseJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupListResponseJSON contains the JSON metadata for
// the struct [RadarAttackLayer7TimeseriesGroupListResponse]
type radarAttackLayer7TimeseriesGroupListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupListResponseResult struct {
	Meta   RadarAttackLayer7TimeseriesGroupListResponseResultMeta   `json:"meta,required"`
	Serie0 RadarAttackLayer7TimeseriesGroupListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7TimeseriesGroupListResponseResultJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupListResponseResultJSON contains the JSON
// metadata for the struct [RadarAttackLayer7TimeseriesGroupListResponseResult]
type radarAttackLayer7TimeseriesGroupListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupListResponseResultMeta struct {
	AggInterval    string                                                               `json:"aggInterval,required"`
	ConfidenceInfo RadarAttackLayer7TimeseriesGroupListResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      RadarAttackLayer7TimeseriesGroupListResponseResultMetaDateRange      `json:"dateRange,required"`
	LastUpdated    time.Time                                                            `json:"lastUpdated,required" format:"date-time"`
	JSON           radarAttackLayer7TimeseriesGroupListResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer7TimeseriesGroupListResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarAttackLayer7TimeseriesGroupListResponseResultMeta]
type radarAttackLayer7TimeseriesGroupListResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7TimeseriesGroupListResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	Level       int64                                                                            `json:"level,required"`
	JSON        radarAttackLayer7TimeseriesGroupListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7TimeseriesGroupListResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupListResponseResultMetaConfidenceInfo]
type radarAttackLayer7TimeseriesGroupListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string                                                                             `json:"dataSource,required"`
	Description string                                                                             `json:"description,required"`
	EndTime     time.Time                                                                          `json:"endTime,required" format:"date-time"`
	EventType   string                                                                             `json:"eventType,required"`
	LinkedURL   string                                                                             `json:"linkedUrl,required"`
	StartTime   time.Time                                                                          `json:"startTime,required" format:"date-time"`
	JSON        radarAttackLayer7TimeseriesGroupListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupListResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7TimeseriesGroupListResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource  apijson.Field
	Description apijson.Field
	EndTime     apijson.Field
	EventType   apijson.Field
	LinkedURL   apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                           `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7TimeseriesGroupListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupListResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupListResponseResultMetaDateRange]
type radarAttackLayer7TimeseriesGroupListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupListResponseResultSerie0 struct {
	AccessRules        []string                                                     `json:"access_rules,required"`
	APIShield          []string                                                     `json:"api_shield,required"`
	BotManagement      []string                                                     `json:"bot_management,required"`
	DataLossPrevention []string                                                     `json:"data_loss_prevention,required"`
	Ddos               []string                                                     `json:"ddos,required"`
	IPReputation       []string                                                     `json:"ip_reputation,required"`
	Timestamps         []time.Time                                                  `json:"timestamps,required" format:"date-time"`
	Waf                []string                                                     `json:"waf,required"`
	JSON               radarAttackLayer7TimeseriesGroupListResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupListResponseResultSerie0JSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7TimeseriesGroupListResponseResultSerie0]
type radarAttackLayer7TimeseriesGroupListResponseResultSerie0JSON struct {
	AccessRules        apijson.Field
	APIShield          apijson.Field
	BotManagement      apijson.Field
	DataLossPrevention apijson.Field
	Ddos               apijson.Field
	IPReputation       apijson.Field
	Timestamps         apijson.Field
	Waf                apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7TimeseriesGroupListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7TimeseriesGroupListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7TimeseriesGroupListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7TimeseriesGroupListParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7TimeseriesGroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7TimeseriesGroupListParamsAggInterval string

const (
	RadarAttackLayer7TimeseriesGroupListParamsAggInterval15m RadarAttackLayer7TimeseriesGroupListParamsAggInterval = "15m"
	RadarAttackLayer7TimeseriesGroupListParamsAggInterval1h  RadarAttackLayer7TimeseriesGroupListParamsAggInterval = "1h"
	RadarAttackLayer7TimeseriesGroupListParamsAggInterval1d  RadarAttackLayer7TimeseriesGroupListParamsAggInterval = "1d"
	RadarAttackLayer7TimeseriesGroupListParamsAggInterval1w  RadarAttackLayer7TimeseriesGroupListParamsAggInterval = "1w"
)

type RadarAttackLayer7TimeseriesGroupListParamsDateRange string

const (
	RadarAttackLayer7TimeseriesGroupListParamsDateRange1d         RadarAttackLayer7TimeseriesGroupListParamsDateRange = "1d"
	RadarAttackLayer7TimeseriesGroupListParamsDateRange7d         RadarAttackLayer7TimeseriesGroupListParamsDateRange = "7d"
	RadarAttackLayer7TimeseriesGroupListParamsDateRange14d        RadarAttackLayer7TimeseriesGroupListParamsDateRange = "14d"
	RadarAttackLayer7TimeseriesGroupListParamsDateRange28d        RadarAttackLayer7TimeseriesGroupListParamsDateRange = "28d"
	RadarAttackLayer7TimeseriesGroupListParamsDateRange12w        RadarAttackLayer7TimeseriesGroupListParamsDateRange = "12w"
	RadarAttackLayer7TimeseriesGroupListParamsDateRange24w        RadarAttackLayer7TimeseriesGroupListParamsDateRange = "24w"
	RadarAttackLayer7TimeseriesGroupListParamsDateRange52w        RadarAttackLayer7TimeseriesGroupListParamsDateRange = "52w"
	RadarAttackLayer7TimeseriesGroupListParamsDateRange1dControl  RadarAttackLayer7TimeseriesGroupListParamsDateRange = "1dControl"
	RadarAttackLayer7TimeseriesGroupListParamsDateRange7dControl  RadarAttackLayer7TimeseriesGroupListParamsDateRange = "7dControl"
	RadarAttackLayer7TimeseriesGroupListParamsDateRange14dControl RadarAttackLayer7TimeseriesGroupListParamsDateRange = "14dControl"
	RadarAttackLayer7TimeseriesGroupListParamsDateRange28dControl RadarAttackLayer7TimeseriesGroupListParamsDateRange = "28dControl"
	RadarAttackLayer7TimeseriesGroupListParamsDateRange12wControl RadarAttackLayer7TimeseriesGroupListParamsDateRange = "12wControl"
	RadarAttackLayer7TimeseriesGroupListParamsDateRange24wControl RadarAttackLayer7TimeseriesGroupListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7TimeseriesGroupListParamsFormat string

const (
	RadarAttackLayer7TimeseriesGroupListParamsFormatJson RadarAttackLayer7TimeseriesGroupListParamsFormat = "JSON"
	RadarAttackLayer7TimeseriesGroupListParamsFormatCsv  RadarAttackLayer7TimeseriesGroupListParamsFormat = "CSV"
)
