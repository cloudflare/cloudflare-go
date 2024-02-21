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

// RadarRankingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarRankingService] method
// instead.
type RadarRankingService struct {
	Options []option.RequestOption
	Ranking *RadarRankingRankingService
}

// NewRadarRankingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarRankingService(opts ...option.RequestOption) (r *RadarRankingService) {
	r = &RadarRankingService{}
	r.Options = opts
	r.Ranking = NewRadarRankingRankingService(opts...)
	return
}

// Gets Domains Rank updates change over time. Raw values are returned.
func (r *RadarRankingService) TimeseriesGroups(ctx context.Context, query RadarRankingTimeseriesGroupsParams, opts ...option.RequestOption) (res *RadarRankingTimeseriesGroupsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarRankingTimeseriesGroupsResponseEnvelope
	path := "radar/ranking/timeseries_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarRankingTimeseriesGroupsResponse struct {
	Meta   RadarRankingTimeseriesGroupsResponseMeta   `json:"meta,required"`
	Serie0 RadarRankingTimeseriesGroupsResponseSerie0 `json:"serie_0,required"`
	JSON   radarRankingTimeseriesGroupsResponseJSON   `json:"-"`
}

// radarRankingTimeseriesGroupsResponseJSON contains the JSON metadata for the
// struct [RadarRankingTimeseriesGroupsResponse]
type radarRankingTimeseriesGroupsResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingTimeseriesGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingTimeseriesGroupsResponseMeta struct {
	DateRange []RadarRankingTimeseriesGroupsResponseMetaDateRange `json:"dateRange,required"`
	JSON      radarRankingTimeseriesGroupsResponseMetaJSON        `json:"-"`
}

// radarRankingTimeseriesGroupsResponseMetaJSON contains the JSON metadata for the
// struct [RadarRankingTimeseriesGroupsResponseMeta]
type radarRankingTimeseriesGroupsResponseMetaJSON struct {
	DateRange   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingTimeseriesGroupsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingTimeseriesGroupsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      radarRankingTimeseriesGroupsResponseMetaDateRangeJSON `json:"-"`
}

// radarRankingTimeseriesGroupsResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarRankingTimeseriesGroupsResponseMetaDateRange]
type radarRankingTimeseriesGroupsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingTimeseriesGroupsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingTimeseriesGroupsResponseSerie0 struct {
	Timestamps  []string                                                `json:"timestamps,required"`
	ExtraFields map[string][]RadarRankingTimeseriesGroupsResponseSerie0 `json:"-,extras"`
	JSON        radarRankingTimeseriesGroupsResponseSerie0JSON          `json:"-"`
}

// radarRankingTimeseriesGroupsResponseSerie0JSON contains the JSON metadata for
// the struct [RadarRankingTimeseriesGroupsResponseSerie0]
type radarRankingTimeseriesGroupsResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingTimeseriesGroupsResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingTimeseriesGroupsParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarRankingTimeseriesGroupsParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Array of comma separated list of domains names.
	Domains param.Field[[]string] `query:"domains"`
	// Format results are returned in.
	Format param.Field[RadarRankingTimeseriesGroupsParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of locations (alpha-2 country codes).
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// The ranking type.
	RankingType param.Field[RadarRankingTimeseriesGroupsParamsRankingType] `query:"rankingType"`
}

// URLQuery serializes [RadarRankingTimeseriesGroupsParams]'s query parameters as
// `url.Values`.
func (r RadarRankingTimeseriesGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarRankingTimeseriesGroupsParamsDateRange string

const (
	RadarRankingTimeseriesGroupsParamsDateRange1d         RadarRankingTimeseriesGroupsParamsDateRange = "1d"
	RadarRankingTimeseriesGroupsParamsDateRange2d         RadarRankingTimeseriesGroupsParamsDateRange = "2d"
	RadarRankingTimeseriesGroupsParamsDateRange7d         RadarRankingTimeseriesGroupsParamsDateRange = "7d"
	RadarRankingTimeseriesGroupsParamsDateRange14d        RadarRankingTimeseriesGroupsParamsDateRange = "14d"
	RadarRankingTimeseriesGroupsParamsDateRange28d        RadarRankingTimeseriesGroupsParamsDateRange = "28d"
	RadarRankingTimeseriesGroupsParamsDateRange12w        RadarRankingTimeseriesGroupsParamsDateRange = "12w"
	RadarRankingTimeseriesGroupsParamsDateRange24w        RadarRankingTimeseriesGroupsParamsDateRange = "24w"
	RadarRankingTimeseriesGroupsParamsDateRange52w        RadarRankingTimeseriesGroupsParamsDateRange = "52w"
	RadarRankingTimeseriesGroupsParamsDateRange1dControl  RadarRankingTimeseriesGroupsParamsDateRange = "1dControl"
	RadarRankingTimeseriesGroupsParamsDateRange2dControl  RadarRankingTimeseriesGroupsParamsDateRange = "2dControl"
	RadarRankingTimeseriesGroupsParamsDateRange7dControl  RadarRankingTimeseriesGroupsParamsDateRange = "7dControl"
	RadarRankingTimeseriesGroupsParamsDateRange14dControl RadarRankingTimeseriesGroupsParamsDateRange = "14dControl"
	RadarRankingTimeseriesGroupsParamsDateRange28dControl RadarRankingTimeseriesGroupsParamsDateRange = "28dControl"
	RadarRankingTimeseriesGroupsParamsDateRange12wControl RadarRankingTimeseriesGroupsParamsDateRange = "12wControl"
	RadarRankingTimeseriesGroupsParamsDateRange24wControl RadarRankingTimeseriesGroupsParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarRankingTimeseriesGroupsParamsFormat string

const (
	RadarRankingTimeseriesGroupsParamsFormatJson RadarRankingTimeseriesGroupsParamsFormat = "JSON"
	RadarRankingTimeseriesGroupsParamsFormatCsv  RadarRankingTimeseriesGroupsParamsFormat = "CSV"
)

// The ranking type.
type RadarRankingTimeseriesGroupsParamsRankingType string

const (
	RadarRankingTimeseriesGroupsParamsRankingTypePopular        RadarRankingTimeseriesGroupsParamsRankingType = "POPULAR"
	RadarRankingTimeseriesGroupsParamsRankingTypeTrendingRise   RadarRankingTimeseriesGroupsParamsRankingType = "TRENDING_RISE"
	RadarRankingTimeseriesGroupsParamsRankingTypeTrendingSteady RadarRankingTimeseriesGroupsParamsRankingType = "TRENDING_STEADY"
)

type RadarRankingTimeseriesGroupsResponseEnvelope struct {
	Result  RadarRankingTimeseriesGroupsResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarRankingTimeseriesGroupsResponseEnvelopeJSON `json:"-"`
}

// radarRankingTimeseriesGroupsResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarRankingTimeseriesGroupsResponseEnvelope]
type radarRankingTimeseriesGroupsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingTimeseriesGroupsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
