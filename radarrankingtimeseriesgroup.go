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

// RadarRankingTimeseriesGroupService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarRankingTimeseriesGroupService] method instead.
type RadarRankingTimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewRadarRankingTimeseriesGroupService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarRankingTimeseriesGroupService(opts ...option.RequestOption) (r *RadarRankingTimeseriesGroupService) {
	r = &RadarRankingTimeseriesGroupService{}
	r.Options = opts
	return
}

// Gets Domains Rank updates change over time. Raw values are returned.
func (r *RadarRankingTimeseriesGroupService) List(ctx context.Context, query RadarRankingTimeseriesGroupListParams, opts ...option.RequestOption) (res *RadarRankingTimeseriesGroupListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/ranking/timeseries_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarRankingTimeseriesGroupListResponse struct {
	Result  RadarRankingTimeseriesGroupListResponseResult `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    radarRankingTimeseriesGroupListResponseJSON   `json:"-"`
}

// radarRankingTimeseriesGroupListResponseJSON contains the JSON metadata for the
// struct [RadarRankingTimeseriesGroupListResponse]
type radarRankingTimeseriesGroupListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingTimeseriesGroupListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingTimeseriesGroupListResponseResult struct {
	Meta   RadarRankingTimeseriesGroupListResponseResultMeta   `json:"meta,required"`
	Serie0 RadarRankingTimeseriesGroupListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarRankingTimeseriesGroupListResponseResultJSON   `json:"-"`
}

// radarRankingTimeseriesGroupListResponseResultJSON contains the JSON metadata for
// the struct [RadarRankingTimeseriesGroupListResponseResult]
type radarRankingTimeseriesGroupListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingTimeseriesGroupListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingTimeseriesGroupListResponseResultMeta struct {
	DateRange []RadarRankingTimeseriesGroupListResponseResultMetaDateRange `json:"dateRange,required"`
	JSON      radarRankingTimeseriesGroupListResponseResultMetaJSON        `json:"-"`
}

// radarRankingTimeseriesGroupListResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarRankingTimeseriesGroupListResponseResultMeta]
type radarRankingTimeseriesGroupListResponseResultMetaJSON struct {
	DateRange   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingTimeseriesGroupListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingTimeseriesGroupListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                      `json:"startTime,required" format:"date-time"`
	JSON      radarRankingTimeseriesGroupListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarRankingTimeseriesGroupListResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarRankingTimeseriesGroupListResponseResultMetaDateRange]
type radarRankingTimeseriesGroupListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingTimeseriesGroupListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingTimeseriesGroupListResponseResultSerie0 struct {
	Timestamps []string                                                `json:"timestamps,required"`
	JSON       radarRankingTimeseriesGroupListResponseResultSerie0JSON `json:"-"`
}

// radarRankingTimeseriesGroupListResponseResultSerie0JSON contains the JSON
// metadata for the struct [RadarRankingTimeseriesGroupListResponseResultSerie0]
type radarRankingTimeseriesGroupListResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingTimeseriesGroupListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingTimeseriesGroupListParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarRankingTimeseriesGroupListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Array of comma separated list of domains names.
	Domains param.Field[[]string] `query:"domains"`
	// Format results are returned in.
	Format param.Field[RadarRankingTimeseriesGroupListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of locations (alpha-2 country codes).
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// The ranking type.
	RankingType param.Field[RadarRankingTimeseriesGroupListParamsRankingType] `query:"rankingType"`
}

// URLQuery serializes [RadarRankingTimeseriesGroupListParams]'s query parameters
// as `url.Values`.
func (r RadarRankingTimeseriesGroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarRankingTimeseriesGroupListParamsDateRange string

const (
	RadarRankingTimeseriesGroupListParamsDateRange1d         RadarRankingTimeseriesGroupListParamsDateRange = "1d"
	RadarRankingTimeseriesGroupListParamsDateRange2d         RadarRankingTimeseriesGroupListParamsDateRange = "2d"
	RadarRankingTimeseriesGroupListParamsDateRange7d         RadarRankingTimeseriesGroupListParamsDateRange = "7d"
	RadarRankingTimeseriesGroupListParamsDateRange14d        RadarRankingTimeseriesGroupListParamsDateRange = "14d"
	RadarRankingTimeseriesGroupListParamsDateRange28d        RadarRankingTimeseriesGroupListParamsDateRange = "28d"
	RadarRankingTimeseriesGroupListParamsDateRange12w        RadarRankingTimeseriesGroupListParamsDateRange = "12w"
	RadarRankingTimeseriesGroupListParamsDateRange24w        RadarRankingTimeseriesGroupListParamsDateRange = "24w"
	RadarRankingTimeseriesGroupListParamsDateRange52w        RadarRankingTimeseriesGroupListParamsDateRange = "52w"
	RadarRankingTimeseriesGroupListParamsDateRange1dControl  RadarRankingTimeseriesGroupListParamsDateRange = "1dControl"
	RadarRankingTimeseriesGroupListParamsDateRange2dControl  RadarRankingTimeseriesGroupListParamsDateRange = "2dControl"
	RadarRankingTimeseriesGroupListParamsDateRange7dControl  RadarRankingTimeseriesGroupListParamsDateRange = "7dControl"
	RadarRankingTimeseriesGroupListParamsDateRange14dControl RadarRankingTimeseriesGroupListParamsDateRange = "14dControl"
	RadarRankingTimeseriesGroupListParamsDateRange28dControl RadarRankingTimeseriesGroupListParamsDateRange = "28dControl"
	RadarRankingTimeseriesGroupListParamsDateRange12wControl RadarRankingTimeseriesGroupListParamsDateRange = "12wControl"
	RadarRankingTimeseriesGroupListParamsDateRange24wControl RadarRankingTimeseriesGroupListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarRankingTimeseriesGroupListParamsFormat string

const (
	RadarRankingTimeseriesGroupListParamsFormatJson RadarRankingTimeseriesGroupListParamsFormat = "JSON"
	RadarRankingTimeseriesGroupListParamsFormatCsv  RadarRankingTimeseriesGroupListParamsFormat = "CSV"
)

// The ranking type.
type RadarRankingTimeseriesGroupListParamsRankingType string

const (
	RadarRankingTimeseriesGroupListParamsRankingTypePopular        RadarRankingTimeseriesGroupListParamsRankingType = "POPULAR"
	RadarRankingTimeseriesGroupListParamsRankingTypeTrendingRise   RadarRankingTimeseriesGroupListParamsRankingType = "TRENDING_RISE"
	RadarRankingTimeseriesGroupListParamsRankingTypeTrendingSteady RadarRankingTimeseriesGroupListParamsRankingType = "TRENDING_STEADY"
)
