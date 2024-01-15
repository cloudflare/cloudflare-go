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

// RadarEmailSecurityArcTimeseryService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityArcTimeseryService] method instead.
type RadarEmailSecurityArcTimeseryService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityArcTimeseryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityArcTimeseryService(opts ...option.RequestOption) (r *RadarEmailSecurityArcTimeseryService) {
	r = &RadarEmailSecurityArcTimeseryService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per Arc validation over time.
func (r *RadarEmailSecurityArcTimeseryService) List(ctx context.Context, query RadarEmailSecurityArcTimeseryListParams, opts ...option.RequestOption) (res *RadarEmailSecurityArcTimeseryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/timeseries_groups/arc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSecurityArcTimeseryListResponse struct {
	Result  RadarEmailSecurityArcTimeseryListResponseResult `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarEmailSecurityArcTimeseryListResponseJSON   `json:"-"`
}

// radarEmailSecurityArcTimeseryListResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecurityArcTimeseryListResponse]
type radarEmailSecurityArcTimeseryListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityArcTimeseryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityArcTimeseryListResponseResult struct {
	Meta   interface{}                                           `json:"meta,required"`
	Serie0 RadarEmailSecurityArcTimeseryListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecurityArcTimeseryListResponseResultJSON   `json:"-"`
}

// radarEmailSecurityArcTimeseryListResponseResultJSON contains the JSON metadata
// for the struct [RadarEmailSecurityArcTimeseryListResponseResult]
type radarEmailSecurityArcTimeseryListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityArcTimeseryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityArcTimeseryListResponseResultSerie0 struct {
	Fail []string                                                  `json:"FAIL,required"`
	None []string                                                  `json:"NONE,required"`
	Pass []string                                                  `json:"PASS,required"`
	JSON radarEmailSecurityArcTimeseryListResponseResultSerie0JSON `json:"-"`
}

// radarEmailSecurityArcTimeseryListResponseResultSerie0JSON contains the JSON
// metadata for the struct [RadarEmailSecurityArcTimeseryListResponseResultSerie0]
type radarEmailSecurityArcTimeseryListResponseResultSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityArcTimeseryListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityArcTimeseryListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityArcTimeseryListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityArcTimeseryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSecurityArcTimeseryListParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecurityArcTimeseryListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityArcTimeseryListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailSecurityArcTimeseryListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityArcTimeseryListParams]'s query parameters
// as `url.Values`.
func (r RadarEmailSecurityArcTimeseryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityArcTimeseryListParamsAggInterval string

const (
	RadarEmailSecurityArcTimeseryListParamsAggInterval15m RadarEmailSecurityArcTimeseryListParamsAggInterval = "15m"
	RadarEmailSecurityArcTimeseryListParamsAggInterval1h  RadarEmailSecurityArcTimeseryListParamsAggInterval = "1h"
	RadarEmailSecurityArcTimeseryListParamsAggInterval1d  RadarEmailSecurityArcTimeseryListParamsAggInterval = "1d"
	RadarEmailSecurityArcTimeseryListParamsAggInterval1w  RadarEmailSecurityArcTimeseryListParamsAggInterval = "1w"
)

type RadarEmailSecurityArcTimeseryListParamsDateRange string

const (
	RadarEmailSecurityArcTimeseryListParamsDateRange1d         RadarEmailSecurityArcTimeseryListParamsDateRange = "1d"
	RadarEmailSecurityArcTimeseryListParamsDateRange2d         RadarEmailSecurityArcTimeseryListParamsDateRange = "2d"
	RadarEmailSecurityArcTimeseryListParamsDateRange7d         RadarEmailSecurityArcTimeseryListParamsDateRange = "7d"
	RadarEmailSecurityArcTimeseryListParamsDateRange14d        RadarEmailSecurityArcTimeseryListParamsDateRange = "14d"
	RadarEmailSecurityArcTimeseryListParamsDateRange28d        RadarEmailSecurityArcTimeseryListParamsDateRange = "28d"
	RadarEmailSecurityArcTimeseryListParamsDateRange12w        RadarEmailSecurityArcTimeseryListParamsDateRange = "12w"
	RadarEmailSecurityArcTimeseryListParamsDateRange24w        RadarEmailSecurityArcTimeseryListParamsDateRange = "24w"
	RadarEmailSecurityArcTimeseryListParamsDateRange52w        RadarEmailSecurityArcTimeseryListParamsDateRange = "52w"
	RadarEmailSecurityArcTimeseryListParamsDateRange1dControl  RadarEmailSecurityArcTimeseryListParamsDateRange = "1dControl"
	RadarEmailSecurityArcTimeseryListParamsDateRange2dControl  RadarEmailSecurityArcTimeseryListParamsDateRange = "2dControl"
	RadarEmailSecurityArcTimeseryListParamsDateRange7dControl  RadarEmailSecurityArcTimeseryListParamsDateRange = "7dControl"
	RadarEmailSecurityArcTimeseryListParamsDateRange14dControl RadarEmailSecurityArcTimeseryListParamsDateRange = "14dControl"
	RadarEmailSecurityArcTimeseryListParamsDateRange28dControl RadarEmailSecurityArcTimeseryListParamsDateRange = "28dControl"
	RadarEmailSecurityArcTimeseryListParamsDateRange12wControl RadarEmailSecurityArcTimeseryListParamsDateRange = "12wControl"
	RadarEmailSecurityArcTimeseryListParamsDateRange24wControl RadarEmailSecurityArcTimeseryListParamsDateRange = "24wControl"
)

type RadarEmailSecurityArcTimeseryListParamsDkim string

const (
	RadarEmailSecurityArcTimeseryListParamsDkimPass RadarEmailSecurityArcTimeseryListParamsDkim = "PASS"
	RadarEmailSecurityArcTimeseryListParamsDkimNone RadarEmailSecurityArcTimeseryListParamsDkim = "NONE"
	RadarEmailSecurityArcTimeseryListParamsDkimFail RadarEmailSecurityArcTimeseryListParamsDkim = "FAIL"
)

type RadarEmailSecurityArcTimeseryListParamsDmarc string

const (
	RadarEmailSecurityArcTimeseryListParamsDmarcPass RadarEmailSecurityArcTimeseryListParamsDmarc = "PASS"
	RadarEmailSecurityArcTimeseryListParamsDmarcNone RadarEmailSecurityArcTimeseryListParamsDmarc = "NONE"
	RadarEmailSecurityArcTimeseryListParamsDmarcFail RadarEmailSecurityArcTimeseryListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityArcTimeseryListParamsFormat string

const (
	RadarEmailSecurityArcTimeseryListParamsFormatJson RadarEmailSecurityArcTimeseryListParamsFormat = "JSON"
	RadarEmailSecurityArcTimeseryListParamsFormatCsv  RadarEmailSecurityArcTimeseryListParamsFormat = "CSV"
)

type RadarEmailSecurityArcTimeseryListParamsSpf string

const (
	RadarEmailSecurityArcTimeseryListParamsSpfPass RadarEmailSecurityArcTimeseryListParamsSpf = "PASS"
	RadarEmailSecurityArcTimeseryListParamsSpfNone RadarEmailSecurityArcTimeseryListParamsSpf = "NONE"
	RadarEmailSecurityArcTimeseryListParamsSpfFail RadarEmailSecurityArcTimeseryListParamsSpf = "FAIL"
)
