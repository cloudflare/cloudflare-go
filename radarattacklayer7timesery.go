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

// RadarAttackLayer7TimeseryService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarAttackLayer7TimeseryService] method instead.
type RadarAttackLayer7TimeseryService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer7TimeseryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer7TimeseryService(opts ...option.RequestOption) (r *RadarAttackLayer7TimeseryService) {
	r = &RadarAttackLayer7TimeseryService{}
	r.Options = opts
	return
}

// Get a timeseries of Layer 7 attacks. Values represent HTTP requests and are
// normalized using min-max by default.
func (r *RadarAttackLayer7TimeseryService) List(ctx context.Context, query RadarAttackLayer7TimeseryListParams, opts ...option.RequestOption) (res *RadarAttackLayer7TimeseryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer7TimeseryListResponse struct {
	Result  RadarAttackLayer7TimeseryListResponseResult `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    radarAttackLayer7TimeseryListResponseJSON   `json:"-"`
}

// radarAttackLayer7TimeseryListResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer7TimeseryListResponse]
type radarAttackLayer7TimeseryListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseryListResponseResult struct {
	Meta   RadarAttackLayer7TimeseryListResponseResultMeta   `json:"meta,required"`
	Serie0 RadarAttackLayer7TimeseryListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7TimeseryListResponseResultJSON   `json:"-"`
}

// radarAttackLayer7TimeseryListResponseResultJSON contains the JSON metadata for
// the struct [RadarAttackLayer7TimeseryListResponseResult]
type radarAttackLayer7TimeseryListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseryListResponseResultMeta struct {
	AggInterval    string                                                        `json:"aggInterval,required"`
	DateRange      []RadarAttackLayer7TimeseryListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    time.Time                                                     `json:"lastUpdated,required" format:"date-time"`
	ConfidenceInfo RadarAttackLayer7TimeseryListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7TimeseryListResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer7TimeseryListResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarAttackLayer7TimeseryListResponseResultMeta]
type radarAttackLayer7TimeseryListResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseryListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseryListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                    `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7TimeseryListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7TimeseryListResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7TimeseryListResponseResultMetaDateRange]
type radarAttackLayer7TimeseryListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseryListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseryListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7TimeseryListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                     `json:"level"`
	JSON        radarAttackLayer7TimeseryListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7TimeseryListResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7TimeseryListResponseResultMetaConfidenceInfo]
type radarAttackLayer7TimeseryListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseryListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseryListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                      `json:"dataSource,required"`
	Description     string                                                                      `json:"description,required"`
	EventType       string                                                                      `json:"eventType,required"`
	IsInstantaneous interface{}                                                                 `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                   `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                      `json:"linkedUrl"`
	StartTime       time.Time                                                                   `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7TimeseryListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7TimeseryListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseryListResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7TimeseryListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7TimeseryListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseryListResponseResultSerie0 struct {
	Timestamps []time.Time                                           `json:"timestamps,required" format:"date-time"`
	Values     []string                                              `json:"values,required"`
	JSON       radarAttackLayer7TimeseryListResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer7TimeseryListResponseResultSerie0JSON contains the JSON metadata
// for the struct [RadarAttackLayer7TimeseryListResponseResultSerie0]
type radarAttackLayer7TimeseryListResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseryListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseryListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7TimeseryListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of L7 attack types.
	Attack param.Field[[]RadarAttackLayer7TimeseryListParamsAttack] `query:"attack"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7TimeseryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7TimeseryListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer7TimeseryListParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [RadarAttackLayer7TimeseryListParams]'s query parameters as
// `url.Values`.
func (r RadarAttackLayer7TimeseryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7TimeseryListParamsAggInterval string

const (
	RadarAttackLayer7TimeseryListParamsAggInterval15m RadarAttackLayer7TimeseryListParamsAggInterval = "15m"
	RadarAttackLayer7TimeseryListParamsAggInterval1h  RadarAttackLayer7TimeseryListParamsAggInterval = "1h"
	RadarAttackLayer7TimeseryListParamsAggInterval1d  RadarAttackLayer7TimeseryListParamsAggInterval = "1d"
	RadarAttackLayer7TimeseryListParamsAggInterval1w  RadarAttackLayer7TimeseryListParamsAggInterval = "1w"
)

type RadarAttackLayer7TimeseryListParamsAttack string

const (
	RadarAttackLayer7TimeseryListParamsAttackDdos               RadarAttackLayer7TimeseryListParamsAttack = "DDOS"
	RadarAttackLayer7TimeseryListParamsAttackWaf                RadarAttackLayer7TimeseryListParamsAttack = "WAF"
	RadarAttackLayer7TimeseryListParamsAttackBotManagement      RadarAttackLayer7TimeseryListParamsAttack = "BOT_MANAGEMENT"
	RadarAttackLayer7TimeseryListParamsAttackAccessRules        RadarAttackLayer7TimeseryListParamsAttack = "ACCESS_RULES"
	RadarAttackLayer7TimeseryListParamsAttackIPReputation       RadarAttackLayer7TimeseryListParamsAttack = "IP_REPUTATION"
	RadarAttackLayer7TimeseryListParamsAttackAPIShield          RadarAttackLayer7TimeseryListParamsAttack = "API_SHIELD"
	RadarAttackLayer7TimeseryListParamsAttackDataLossPrevention RadarAttackLayer7TimeseryListParamsAttack = "DATA_LOSS_PREVENTION"
)

type RadarAttackLayer7TimeseryListParamsDateRange string

const (
	RadarAttackLayer7TimeseryListParamsDateRange1d         RadarAttackLayer7TimeseryListParamsDateRange = "1d"
	RadarAttackLayer7TimeseryListParamsDateRange2d         RadarAttackLayer7TimeseryListParamsDateRange = "2d"
	RadarAttackLayer7TimeseryListParamsDateRange7d         RadarAttackLayer7TimeseryListParamsDateRange = "7d"
	RadarAttackLayer7TimeseryListParamsDateRange14d        RadarAttackLayer7TimeseryListParamsDateRange = "14d"
	RadarAttackLayer7TimeseryListParamsDateRange28d        RadarAttackLayer7TimeseryListParamsDateRange = "28d"
	RadarAttackLayer7TimeseryListParamsDateRange12w        RadarAttackLayer7TimeseryListParamsDateRange = "12w"
	RadarAttackLayer7TimeseryListParamsDateRange24w        RadarAttackLayer7TimeseryListParamsDateRange = "24w"
	RadarAttackLayer7TimeseryListParamsDateRange52w        RadarAttackLayer7TimeseryListParamsDateRange = "52w"
	RadarAttackLayer7TimeseryListParamsDateRange1dControl  RadarAttackLayer7TimeseryListParamsDateRange = "1dControl"
	RadarAttackLayer7TimeseryListParamsDateRange2dControl  RadarAttackLayer7TimeseryListParamsDateRange = "2dControl"
	RadarAttackLayer7TimeseryListParamsDateRange7dControl  RadarAttackLayer7TimeseryListParamsDateRange = "7dControl"
	RadarAttackLayer7TimeseryListParamsDateRange14dControl RadarAttackLayer7TimeseryListParamsDateRange = "14dControl"
	RadarAttackLayer7TimeseryListParamsDateRange28dControl RadarAttackLayer7TimeseryListParamsDateRange = "28dControl"
	RadarAttackLayer7TimeseryListParamsDateRange12wControl RadarAttackLayer7TimeseryListParamsDateRange = "12wControl"
	RadarAttackLayer7TimeseryListParamsDateRange24wControl RadarAttackLayer7TimeseryListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7TimeseryListParamsFormat string

const (
	RadarAttackLayer7TimeseryListParamsFormatJson RadarAttackLayer7TimeseryListParamsFormat = "JSON"
	RadarAttackLayer7TimeseryListParamsFormatCsv  RadarAttackLayer7TimeseryListParamsFormat = "CSV"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7TimeseryListParamsNormalization string

const (
	RadarAttackLayer7TimeseryListParamsNormalizationPercentageChange RadarAttackLayer7TimeseryListParamsNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer7TimeseryListParamsNormalizationMin0Max          RadarAttackLayer7TimeseryListParamsNormalization = "MIN0_MAX"
)
