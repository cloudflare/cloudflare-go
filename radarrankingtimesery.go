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

// RadarRankingTimeseryService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarRankingTimeseryService]
// method instead.
type RadarRankingTimeseryService struct {
	Options []option.RequestOption
}

// NewRadarRankingTimeseryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarRankingTimeseryService(opts ...option.RequestOption) (r *RadarRankingTimeseryService) {
	r = &RadarRankingTimeseryService{}
	r.Options = opts
	return
}

// Gets Domains Rank updates change over time. Raw values are returned.
func (r *RadarRankingTimeseryService) List(ctx context.Context, query RadarRankingTimeseryListParams, opts ...option.RequestOption) (res *RadarRankingTimeseryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/ranking/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarRankingTimeseryListResponse struct {
	Result  RadarRankingTimeseryListResponseResult `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    radarRankingTimeseryListResponseJSON   `json:"-"`
}

// radarRankingTimeseryListResponseJSON contains the JSON metadata for the struct
// [RadarRankingTimeseryListResponse]
type radarRankingTimeseryListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingTimeseryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingTimeseryListResponseResult struct {
	Meta   RadarRankingTimeseryListResponseResultMeta   `json:"meta,required"`
	Serie0 RadarRankingTimeseryListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarRankingTimeseryListResponseResultJSON   `json:"-"`
}

// radarRankingTimeseryListResponseResultJSON contains the JSON metadata for the
// struct [RadarRankingTimeseryListResponseResult]
type radarRankingTimeseryListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingTimeseryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingTimeseryListResponseResultMeta struct {
	DateRange RadarRankingTimeseryListResponseResultMetaDateRange `json:"dateRange,required"`
	JSON      radarRankingTimeseryListResponseResultMetaJSON      `json:"-"`
}

// radarRankingTimeseryListResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarRankingTimeseryListResponseResultMeta]
type radarRankingTimeseryListResponseResultMetaJSON struct {
	DateRange   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingTimeseryListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingTimeseryListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      radarRankingTimeseryListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarRankingTimeseryListResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarRankingTimeseryListResponseResultMetaDateRange]
type radarRankingTimeseryListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingTimeseryListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingTimeseryListResponseResultSerie0 struct {
	DomainName []int64                                          `json:"<domain name>,required"`
	Timestamps []string                                         `json:"timestamps,required"`
	JSON       radarRankingTimeseryListResponseResultSerie0JSON `json:"-"`
}

// radarRankingTimeseryListResponseResultSerie0JSON contains the JSON metadata for
// the struct [RadarRankingTimeseryListResponseResultSerie0]
type radarRankingTimeseryListResponseResultSerie0JSON struct {
	DomainName  apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingTimeseryListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingTimeseryListParams struct {
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarRankingTimeseryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Array of comma separated list of domains names.
	Domains param.Field[[]string] `query:"domains"`
	// Format results are returned in.
	Format param.Field[RadarRankingTimeseryListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of locations (alpha-2 country codes).
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarRankingTimeseryListParams]'s query parameters as
// `url.Values`.
func (r RadarRankingTimeseryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarRankingTimeseryListParamsDateRange string

const (
	RadarRankingTimeseryListParamsDateRange1d         RadarRankingTimeseryListParamsDateRange = "1d"
	RadarRankingTimeseryListParamsDateRange7d         RadarRankingTimeseryListParamsDateRange = "7d"
	RadarRankingTimeseryListParamsDateRange14d        RadarRankingTimeseryListParamsDateRange = "14d"
	RadarRankingTimeseryListParamsDateRange28d        RadarRankingTimeseryListParamsDateRange = "28d"
	RadarRankingTimeseryListParamsDateRange12w        RadarRankingTimeseryListParamsDateRange = "12w"
	RadarRankingTimeseryListParamsDateRange24w        RadarRankingTimeseryListParamsDateRange = "24w"
	RadarRankingTimeseryListParamsDateRange52w        RadarRankingTimeseryListParamsDateRange = "52w"
	RadarRankingTimeseryListParamsDateRange1dControl  RadarRankingTimeseryListParamsDateRange = "1dControl"
	RadarRankingTimeseryListParamsDateRange7dControl  RadarRankingTimeseryListParamsDateRange = "7dControl"
	RadarRankingTimeseryListParamsDateRange14dControl RadarRankingTimeseryListParamsDateRange = "14dControl"
	RadarRankingTimeseryListParamsDateRange28dControl RadarRankingTimeseryListParamsDateRange = "28dControl"
	RadarRankingTimeseryListParamsDateRange12wControl RadarRankingTimeseryListParamsDateRange = "12wControl"
	RadarRankingTimeseryListParamsDateRange24wControl RadarRankingTimeseryListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarRankingTimeseryListParamsFormat string

const (
	RadarRankingTimeseryListParamsFormatJson RadarRankingTimeseryListParamsFormat = "JSON"
	RadarRankingTimeseryListParamsFormatCsv  RadarRankingTimeseryListParamsFormat = "CSV"
)
