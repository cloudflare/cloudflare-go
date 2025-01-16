// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
	"github.com/tidwall/gjson"
)

// RankingService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRankingService] method instead.
type RankingService struct {
	Options []option.RequestOption
	Domain  *RankingDomainService
}

// NewRankingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRankingService(opts ...option.RequestOption) (r *RankingService) {
	r = &RankingService{}
	r.Options = opts
	r.Domain = NewRankingDomainService(opts...)
	return
}

// Gets Domains Rank updates change over time. Raw values are returned.
func (r *RankingService) TimeseriesGroups(ctx context.Context, query RankingTimeseriesGroupsParams, opts ...option.RequestOption) (res *RankingTimeseriesGroupsResponse, err error) {
	var env RankingTimeseriesGroupsResponseEnvelope
	opts = append(r.Options[:], opts...)
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
func (r *RankingService) Top(ctx context.Context, query RankingTopParams, opts ...option.RequestOption) (res *RankingTopResponse, err error) {
	var env RankingTopResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/ranking/top"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RankingTimeseriesGroupsResponse struct {
	Meta   RankingTimeseriesGroupsResponseMeta   `json:"meta,required"`
	Serie0 RankingTimeseriesGroupsResponseSerie0 `json:"serie_0,required"`
	JSON   rankingTimeseriesGroupsResponseJSON   `json:"-"`
}

// rankingTimeseriesGroupsResponseJSON contains the JSON metadata for the struct
// [RankingTimeseriesGroupsResponse]
type rankingTimeseriesGroupsResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingTimeseriesGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingTimeseriesGroupsResponseJSON) RawJSON() string {
	return r.raw
}

type RankingTimeseriesGroupsResponseMeta struct {
	DateRange []RankingTimeseriesGroupsResponseMetaDateRange `json:"dateRange,required"`
	JSON      rankingTimeseriesGroupsResponseMetaJSON        `json:"-"`
}

// rankingTimeseriesGroupsResponseMetaJSON contains the JSON metadata for the
// struct [RankingTimeseriesGroupsResponseMeta]
type rankingTimeseriesGroupsResponseMetaJSON struct {
	DateRange   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingTimeseriesGroupsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingTimeseriesGroupsResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RankingTimeseriesGroupsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                        `json:"startTime,required" format:"date-time"`
	JSON      rankingTimeseriesGroupsResponseMetaDateRangeJSON `json:"-"`
}

// rankingTimeseriesGroupsResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [RankingTimeseriesGroupsResponseMetaDateRange]
type rankingTimeseriesGroupsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingTimeseriesGroupsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingTimeseriesGroupsResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RankingTimeseriesGroupsResponseSerie0 struct {
	Timestamps  []string                                                `json:"timestamps,required"`
	ExtraFields map[string][]RankingTimeseriesGroupsResponseSerie0Union `json:"-,extras"`
	JSON        rankingTimeseriesGroupsResponseSerie0JSON               `json:"-"`
}

// rankingTimeseriesGroupsResponseSerie0JSON contains the JSON metadata for the
// struct [RankingTimeseriesGroupsResponseSerie0]
type rankingTimeseriesGroupsResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingTimeseriesGroupsResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingTimeseriesGroupsResponseSerie0JSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type RankingTimeseriesGroupsResponseSerie0Union interface {
	ImplementsRadarRankingTimeseriesGroupsResponseSerie0Union()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RankingTimeseriesGroupsResponseSerie0Union)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type RankingTopResponse struct {
	Meta RankingTopResponseMeta   `json:"meta,required"`
	Top0 []RankingTopResponseTop0 `json:"top_0,required"`
	JSON rankingTopResponseJSON   `json:"-"`
}

// rankingTopResponseJSON contains the JSON metadata for the struct
// [RankingTopResponse]
type rankingTopResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingTopResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingTopResponseJSON) RawJSON() string {
	return r.raw
}

type RankingTopResponseMeta struct {
	Top0 RankingTopResponseMetaTop0 `json:"top_0,required"`
	JSON rankingTopResponseMetaJSON `json:"-"`
}

// rankingTopResponseMetaJSON contains the JSON metadata for the struct
// [RankingTopResponseMeta]
type rankingTopResponseMetaJSON struct {
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingTopResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingTopResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RankingTopResponseMetaTop0 struct {
	Date string                         `json:"date,required"`
	JSON rankingTopResponseMetaTop0JSON `json:"-"`
}

// rankingTopResponseMetaTop0JSON contains the JSON metadata for the struct
// [RankingTopResponseMetaTop0]
type rankingTopResponseMetaTop0JSON struct {
	Date        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingTopResponseMetaTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingTopResponseMetaTop0JSON) RawJSON() string {
	return r.raw
}

type RankingTopResponseTop0 struct {
	Categories []RankingTopResponseTop0Category `json:"categories,required"`
	Domain     string                           `json:"domain,required"`
	Rank       int64                            `json:"rank,required"`
	// Only available in TRENDING rankings.
	PctRankChange float64                    `json:"pctRankChange"`
	JSON          rankingTopResponseTop0JSON `json:"-"`
}

// rankingTopResponseTop0JSON contains the JSON metadata for the struct
// [RankingTopResponseTop0]
type rankingTopResponseTop0JSON struct {
	Categories    apijson.Field
	Domain        apijson.Field
	Rank          apijson.Field
	PctRankChange apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RankingTopResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingTopResponseTop0JSON) RawJSON() string {
	return r.raw
}

type RankingTopResponseTop0Category struct {
	ID              float64                            `json:"id,required"`
	Name            string                             `json:"name,required"`
	SuperCategoryID float64                            `json:"superCategoryId,required"`
	JSON            rankingTopResponseTop0CategoryJSON `json:"-"`
}

// rankingTopResponseTop0CategoryJSON contains the JSON metadata for the struct
// [RankingTopResponseTop0Category]
type rankingTopResponseTop0CategoryJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RankingTopResponseTop0Category) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingTopResponseTop0CategoryJSON) RawJSON() string {
	return r.raw
}

type RankingTimeseriesGroupsParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Array of comma separated list of domains names.
	Domains param.Field[[]string] `query:"domains"`
	// Format results are returned in.
	Format param.Field[RankingTimeseriesGroupsParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of locations (alpha-2 country codes).
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// The ranking type.
	RankingType param.Field[RankingTimeseriesGroupsParamsRankingType] `query:"rankingType"`
}

// URLQuery serializes [RankingTimeseriesGroupsParams]'s query parameters as
// `url.Values`.
func (r RankingTimeseriesGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format results are returned in.
type RankingTimeseriesGroupsParamsFormat string

const (
	RankingTimeseriesGroupsParamsFormatJson RankingTimeseriesGroupsParamsFormat = "JSON"
	RankingTimeseriesGroupsParamsFormatCsv  RankingTimeseriesGroupsParamsFormat = "CSV"
)

func (r RankingTimeseriesGroupsParamsFormat) IsKnown() bool {
	switch r {
	case RankingTimeseriesGroupsParamsFormatJson, RankingTimeseriesGroupsParamsFormatCsv:
		return true
	}
	return false
}

// The ranking type.
type RankingTimeseriesGroupsParamsRankingType string

const (
	RankingTimeseriesGroupsParamsRankingTypePopular        RankingTimeseriesGroupsParamsRankingType = "POPULAR"
	RankingTimeseriesGroupsParamsRankingTypeTrendingRise   RankingTimeseriesGroupsParamsRankingType = "TRENDING_RISE"
	RankingTimeseriesGroupsParamsRankingTypeTrendingSteady RankingTimeseriesGroupsParamsRankingType = "TRENDING_STEADY"
)

func (r RankingTimeseriesGroupsParamsRankingType) IsKnown() bool {
	switch r {
	case RankingTimeseriesGroupsParamsRankingTypePopular, RankingTimeseriesGroupsParamsRankingTypeTrendingRise, RankingTimeseriesGroupsParamsRankingTypeTrendingSteady:
		return true
	}
	return false
}

type RankingTimeseriesGroupsResponseEnvelope struct {
	Result  RankingTimeseriesGroupsResponse             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    rankingTimeseriesGroupsResponseEnvelopeJSON `json:"-"`
}

// rankingTimeseriesGroupsResponseEnvelopeJSON contains the JSON metadata for the
// struct [RankingTimeseriesGroupsResponseEnvelope]
type rankingTimeseriesGroupsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingTimeseriesGroupsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingTimeseriesGroupsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RankingTopParams struct {
	// Array of dates to filter the ranking.
	Date param.Field[[]string] `query:"date"`
	// Format results are returned in.
	Format param.Field[RankingTopParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of locations (alpha-2 country codes).
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// The ranking type.
	RankingType param.Field[RankingTopParamsRankingType] `query:"rankingType"`
}

// URLQuery serializes [RankingTopParams]'s query parameters as `url.Values`.
func (r RankingTopParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format results are returned in.
type RankingTopParamsFormat string

const (
	RankingTopParamsFormatJson RankingTopParamsFormat = "JSON"
	RankingTopParamsFormatCsv  RankingTopParamsFormat = "CSV"
)

func (r RankingTopParamsFormat) IsKnown() bool {
	switch r {
	case RankingTopParamsFormatJson, RankingTopParamsFormatCsv:
		return true
	}
	return false
}

// The ranking type.
type RankingTopParamsRankingType string

const (
	RankingTopParamsRankingTypePopular        RankingTopParamsRankingType = "POPULAR"
	RankingTopParamsRankingTypeTrendingRise   RankingTopParamsRankingType = "TRENDING_RISE"
	RankingTopParamsRankingTypeTrendingSteady RankingTopParamsRankingType = "TRENDING_STEADY"
)

func (r RankingTopParamsRankingType) IsKnown() bool {
	switch r {
	case RankingTopParamsRankingTypePopular, RankingTopParamsRankingTypeTrendingRise, RankingTopParamsRankingTypeTrendingSteady:
		return true
	}
	return false
}

type RankingTopResponseEnvelope struct {
	Result  RankingTopResponse             `json:"result,required"`
	Success bool                           `json:"success,required"`
	JSON    rankingTopResponseEnvelopeJSON `json:"-"`
}

// rankingTopResponseEnvelopeJSON contains the JSON metadata for the struct
// [RankingTopResponseEnvelope]
type rankingTopResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingTopResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingTopResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
