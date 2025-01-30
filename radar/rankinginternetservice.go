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

// RankingInternetServiceService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRankingInternetServiceService] method instead.
type RankingInternetServiceService struct {
	Options []option.RequestOption
}

// NewRankingInternetServiceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRankingInternetServiceService(opts ...option.RequestOption) (r *RankingInternetServiceService) {
	r = &RankingInternetServiceService{}
	r.Options = opts
	return
}

// Gets Internet Services rank update changes over time. Raw values are returned.
func (r *RankingInternetServiceService) TimeseriesGroups(ctx context.Context, query RankingInternetServiceTimeseriesGroupsParams, opts ...option.RequestOption) (res *RankingInternetServiceTimeseriesGroupsResponse, err error) {
	var env RankingInternetServiceTimeseriesGroupsResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/ranking/internet_services/timeseries_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get top Internet services based on their rank.
func (r *RankingInternetServiceService) Top(ctx context.Context, query RankingInternetServiceTopParams, opts ...option.RequestOption) (res *RankingInternetServiceTopResponse, err error) {
	var env RankingInternetServiceTopResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/ranking/internet_services/top"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RankingInternetServiceTimeseriesGroupsResponse struct {
	Meta   RankingInternetServiceTimeseriesGroupsResponseMeta   `json:"meta,required"`
	Serie0 RankingInternetServiceTimeseriesGroupsResponseSerie0 `json:"serie_0,required"`
	JSON   rankingInternetServiceTimeseriesGroupsResponseJSON   `json:"-"`
}

// rankingInternetServiceTimeseriesGroupsResponseJSON contains the JSON metadata
// for the struct [RankingInternetServiceTimeseriesGroupsResponse]
type rankingInternetServiceTimeseriesGroupsResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingInternetServiceTimeseriesGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingInternetServiceTimeseriesGroupsResponseJSON) RawJSON() string {
	return r.raw
}

type RankingInternetServiceTimeseriesGroupsResponseMeta struct {
	DateRange []RankingInternetServiceTimeseriesGroupsResponseMetaDateRange `json:"dateRange,required"`
	JSON      rankingInternetServiceTimeseriesGroupsResponseMetaJSON        `json:"-"`
}

// rankingInternetServiceTimeseriesGroupsResponseMetaJSON contains the JSON
// metadata for the struct [RankingInternetServiceTimeseriesGroupsResponseMeta]
type rankingInternetServiceTimeseriesGroupsResponseMetaJSON struct {
	DateRange   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingInternetServiceTimeseriesGroupsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingInternetServiceTimeseriesGroupsResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RankingInternetServiceTimeseriesGroupsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                       `json:"startTime,required" format:"date-time"`
	JSON      rankingInternetServiceTimeseriesGroupsResponseMetaDateRangeJSON `json:"-"`
}

// rankingInternetServiceTimeseriesGroupsResponseMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RankingInternetServiceTimeseriesGroupsResponseMetaDateRange]
type rankingInternetServiceTimeseriesGroupsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingInternetServiceTimeseriesGroupsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingInternetServiceTimeseriesGroupsResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RankingInternetServiceTimeseriesGroupsResponseSerie0 struct {
	Timestamps  []string                                                               `json:"timestamps,required"`
	ExtraFields map[string][]RankingInternetServiceTimeseriesGroupsResponseSerie0Union `json:"-,extras"`
	JSON        rankingInternetServiceTimeseriesGroupsResponseSerie0JSON               `json:"-"`
}

// rankingInternetServiceTimeseriesGroupsResponseSerie0JSON contains the JSON
// metadata for the struct [RankingInternetServiceTimeseriesGroupsResponseSerie0]
type rankingInternetServiceTimeseriesGroupsResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingInternetServiceTimeseriesGroupsResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingInternetServiceTimeseriesGroupsResponseSerie0JSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type RankingInternetServiceTimeseriesGroupsResponseSerie0Union interface {
	ImplementsRankingInternetServiceTimeseriesGroupsResponseSerie0Union()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RankingInternetServiceTimeseriesGroupsResponseSerie0Union)(nil)).Elem(),
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

type RankingInternetServiceTopResponse struct {
	Meta RankingInternetServiceTopResponseMeta   `json:"meta,required"`
	Top0 []RankingInternetServiceTopResponseTop0 `json:"top_0,required"`
	JSON rankingInternetServiceTopResponseJSON   `json:"-"`
}

// rankingInternetServiceTopResponseJSON contains the JSON metadata for the struct
// [RankingInternetServiceTopResponse]
type rankingInternetServiceTopResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingInternetServiceTopResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingInternetServiceTopResponseJSON) RawJSON() string {
	return r.raw
}

type RankingInternetServiceTopResponseMeta struct {
	Top0 RankingInternetServiceTopResponseMetaTop0 `json:"top_0,required"`
	JSON rankingInternetServiceTopResponseMetaJSON `json:"-"`
}

// rankingInternetServiceTopResponseMetaJSON contains the JSON metadata for the
// struct [RankingInternetServiceTopResponseMeta]
type rankingInternetServiceTopResponseMetaJSON struct {
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingInternetServiceTopResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingInternetServiceTopResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RankingInternetServiceTopResponseMetaTop0 struct {
	Date            string                                        `json:"date,required"`
	ServiceCategory string                                        `json:"serviceCategory,required"`
	JSON            rankingInternetServiceTopResponseMetaTop0JSON `json:"-"`
}

// rankingInternetServiceTopResponseMetaTop0JSON contains the JSON metadata for the
// struct [RankingInternetServiceTopResponseMetaTop0]
type rankingInternetServiceTopResponseMetaTop0JSON struct {
	Date            apijson.Field
	ServiceCategory apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RankingInternetServiceTopResponseMetaTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingInternetServiceTopResponseMetaTop0JSON) RawJSON() string {
	return r.raw
}

type RankingInternetServiceTopResponseTop0 struct {
	Rank    int64                                     `json:"rank,required"`
	Service string                                    `json:"service,required"`
	JSON    rankingInternetServiceTopResponseTop0JSON `json:"-"`
}

// rankingInternetServiceTopResponseTop0JSON contains the JSON metadata for the
// struct [RankingInternetServiceTopResponseTop0]
type rankingInternetServiceTopResponseTop0JSON struct {
	Rank        apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingInternetServiceTopResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingInternetServiceTopResponseTop0JSON) RawJSON() string {
	return r.raw
}

type RankingInternetServiceTimeseriesGroupsParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RankingInternetServiceTimeseriesGroupsParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for services category.
	ServiceCategory param.Field[[]string] `query:"serviceCategory"`
}

// URLQuery serializes [RankingInternetServiceTimeseriesGroupsParams]'s query
// parameters as `url.Values`.
func (r RankingInternetServiceTimeseriesGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format results are returned in.
type RankingInternetServiceTimeseriesGroupsParamsFormat string

const (
	RankingInternetServiceTimeseriesGroupsParamsFormatJson RankingInternetServiceTimeseriesGroupsParamsFormat = "JSON"
	RankingInternetServiceTimeseriesGroupsParamsFormatCsv  RankingInternetServiceTimeseriesGroupsParamsFormat = "CSV"
)

func (r RankingInternetServiceTimeseriesGroupsParamsFormat) IsKnown() bool {
	switch r {
	case RankingInternetServiceTimeseriesGroupsParamsFormatJson, RankingInternetServiceTimeseriesGroupsParamsFormatCsv:
		return true
	}
	return false
}

type RankingInternetServiceTimeseriesGroupsResponseEnvelope struct {
	Result  RankingInternetServiceTimeseriesGroupsResponse             `json:"result,required"`
	Success bool                                                       `json:"success,required"`
	JSON    rankingInternetServiceTimeseriesGroupsResponseEnvelopeJSON `json:"-"`
}

// rankingInternetServiceTimeseriesGroupsResponseEnvelopeJSON contains the JSON
// metadata for the struct [RankingInternetServiceTimeseriesGroupsResponseEnvelope]
type rankingInternetServiceTimeseriesGroupsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingInternetServiceTimeseriesGroupsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingInternetServiceTimeseriesGroupsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RankingInternetServiceTopParams struct {
	// Array of dates to filter the ranking.
	Date param.Field[[]time.Time] `query:"date" format:"date"`
	// Format results are returned in.
	Format param.Field[RankingInternetServiceTopParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for services category.
	ServiceCategory param.Field[[]string] `query:"serviceCategory"`
}

// URLQuery serializes [RankingInternetServiceTopParams]'s query parameters as
// `url.Values`.
func (r RankingInternetServiceTopParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format results are returned in.
type RankingInternetServiceTopParamsFormat string

const (
	RankingInternetServiceTopParamsFormatJson RankingInternetServiceTopParamsFormat = "JSON"
	RankingInternetServiceTopParamsFormatCsv  RankingInternetServiceTopParamsFormat = "CSV"
)

func (r RankingInternetServiceTopParamsFormat) IsKnown() bool {
	switch r {
	case RankingInternetServiceTopParamsFormatJson, RankingInternetServiceTopParamsFormatCsv:
		return true
	}
	return false
}

type RankingInternetServiceTopResponseEnvelope struct {
	Result  RankingInternetServiceTopResponse             `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    rankingInternetServiceTopResponseEnvelopeJSON `json:"-"`
}

// rankingInternetServiceTopResponseEnvelopeJSON contains the JSON metadata for the
// struct [RankingInternetServiceTopResponseEnvelope]
type rankingInternetServiceTopResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingInternetServiceTopResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingInternetServiceTopResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
