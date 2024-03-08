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

// RadarRankingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarRankingService] method
// instead.
type RadarRankingService struct {
	Options []option.RequestOption
	Domain  *RadarRankingDomainService
}

// NewRadarRankingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarRankingService(opts ...option.RequestOption) (r *RadarRankingService) {
	r = &RadarRankingService{}
	r.Options = opts
	r.Domain = NewRadarRankingDomainService(opts...)
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

// Get top or trending domains based on their rank. Popular domains are domains of
// broad appeal based on how people use the Internet. Trending domains are domains
// that are generating a surge in interest. For more information on top domains,
// see https://blog.cloudflare.com/radar-domain-rankings/.
func (r *RadarRankingService) Top(ctx context.Context, query RadarRankingTopParams, opts ...option.RequestOption) (res *RadarRankingTopResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarRankingTopResponseEnvelope
	path := "radar/ranking/top"
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

func (r radarRankingTimeseriesGroupsResponseJSON) RawJSON() string {
	return r.raw
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

func (r radarRankingTimeseriesGroupsResponseMetaJSON) RawJSON() string {
	return r.raw
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

func (r radarRankingTimeseriesGroupsResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
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

func (r radarRankingTimeseriesGroupsResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarRankingTopResponse struct {
	Meta RadarRankingTopResponseMeta   `json:"meta,required"`
	Top0 []RadarRankingTopResponseTop0 `json:"top_0,required"`
	JSON radarRankingTopResponseJSON   `json:"-"`
}

// radarRankingTopResponseJSON contains the JSON metadata for the struct
// [RadarRankingTopResponse]
type radarRankingTopResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingTopResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingTopResponseJSON) RawJSON() string {
	return r.raw
}

type RadarRankingTopResponseMeta struct {
	Top0 RadarRankingTopResponseMetaTop0 `json:"top_0,required"`
	JSON radarRankingTopResponseMetaJSON `json:"-"`
}

// radarRankingTopResponseMetaJSON contains the JSON metadata for the struct
// [RadarRankingTopResponseMeta]
type radarRankingTopResponseMetaJSON struct {
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingTopResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingTopResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarRankingTopResponseMetaTop0 struct {
	Date string                              `json:"date,required"`
	JSON radarRankingTopResponseMetaTop0JSON `json:"-"`
}

// radarRankingTopResponseMetaTop0JSON contains the JSON metadata for the struct
// [RadarRankingTopResponseMetaTop0]
type radarRankingTopResponseMetaTop0JSON struct {
	Date        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingTopResponseMetaTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingTopResponseMetaTop0JSON) RawJSON() string {
	return r.raw
}

type RadarRankingTopResponseTop0 struct {
	Categories []RadarRankingTopResponseTop0Category `json:"categories,required"`
	Domain     string                                `json:"domain,required"`
	Rank       int64                                 `json:"rank,required"`
	// Only available in TRENDING rankings.
	PctRankChange float64                         `json:"pctRankChange"`
	JSON          radarRankingTopResponseTop0JSON `json:"-"`
}

// radarRankingTopResponseTop0JSON contains the JSON metadata for the struct
// [RadarRankingTopResponseTop0]
type radarRankingTopResponseTop0JSON struct {
	Categories    apijson.Field
	Domain        apijson.Field
	Rank          apijson.Field
	PctRankChange apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RadarRankingTopResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingTopResponseTop0JSON) RawJSON() string {
	return r.raw
}

type RadarRankingTopResponseTop0Category struct {
	ID              float64                                 `json:"id,required"`
	Name            string                                  `json:"name,required"`
	SuperCategoryID float64                                 `json:"superCategoryId,required"`
	JSON            radarRankingTopResponseTop0CategoryJSON `json:"-"`
}

// radarRankingTopResponseTop0CategoryJSON contains the JSON metadata for the
// struct [RadarRankingTopResponseTop0Category]
type radarRankingTopResponseTop0CategoryJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarRankingTopResponseTop0Category) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingTopResponseTop0CategoryJSON) RawJSON() string {
	return r.raw
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

func (r radarRankingTimeseriesGroupsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarRankingTopParams struct {
	// Array of dates to filter the ranking.
	Date param.Field[[]string] `query:"date"`
	// Format results are returned in.
	Format param.Field[RadarRankingTopParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of locations (alpha-2 country codes).
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// The ranking type.
	RankingType param.Field[RadarRankingTopParamsRankingType] `query:"rankingType"`
}

// URLQuery serializes [RadarRankingTopParams]'s query parameters as `url.Values`.
func (r RadarRankingTopParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarRankingTopParamsFormat string

const (
	RadarRankingTopParamsFormatJson RadarRankingTopParamsFormat = "JSON"
	RadarRankingTopParamsFormatCsv  RadarRankingTopParamsFormat = "CSV"
)

// The ranking type.
type RadarRankingTopParamsRankingType string

const (
	RadarRankingTopParamsRankingTypePopular        RadarRankingTopParamsRankingType = "POPULAR"
	RadarRankingTopParamsRankingTypeTrendingRise   RadarRankingTopParamsRankingType = "TRENDING_RISE"
	RadarRankingTopParamsRankingTypeTrendingSteady RadarRankingTopParamsRankingType = "TRENDING_STEADY"
)

type RadarRankingTopResponseEnvelope struct {
	Result  RadarRankingTopResponse             `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    radarRankingTopResponseEnvelopeJSON `json:"-"`
}

// radarRankingTopResponseEnvelopeJSON contains the JSON metadata for the struct
// [RadarRankingTopResponseEnvelope]
type radarRankingTopResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingTopResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingTopResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
