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

// RadarEmailTimeseryArcService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEmailTimeseryArcService]
// method instead.
type RadarEmailTimeseryArcService struct {
	Options []option.RequestOption
}

// NewRadarEmailTimeseryArcService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarEmailTimeseryArcService(opts ...option.RequestOption) (r *RadarEmailTimeseryArcService) {
	r = &RadarEmailTimeseryArcService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per Arc validation over time.
func (r *RadarEmailTimeseryArcService) List(ctx context.Context, query RadarEmailTimeseryArcListParams, opts ...option.RequestOption) (res *RadarEmailTimeseryArcListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/timeseries/arc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailTimeseryArcListResponse struct {
	Result  RadarEmailTimeseryArcListResponseResult `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    radarEmailTimeseryArcListResponseJSON   `json:"-"`
}

// radarEmailTimeseryArcListResponseJSON contains the JSON metadata for the struct
// [RadarEmailTimeseryArcListResponse]
type radarEmailTimeseryArcListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTimeseryArcListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTimeseryArcListResponseResult struct {
	Meta   interface{}                                   `json:"meta,required"`
	Serie0 RadarEmailTimeseryArcListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailTimeseryArcListResponseResultJSON   `json:"-"`
}

// radarEmailTimeseryArcListResponseResultJSON contains the JSON metadata for the
// struct [RadarEmailTimeseryArcListResponseResult]
type radarEmailTimeseryArcListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTimeseryArcListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTimeseryArcListResponseResultSerie0 struct {
	Fail []string                                          `json:"FAIL,required"`
	None []string                                          `json:"NONE,required"`
	Pass []string                                          `json:"PASS,required"`
	JSON radarEmailTimeseryArcListResponseResultSerie0JSON `json:"-"`
}

// radarEmailTimeseryArcListResponseResultSerie0JSON contains the JSON metadata for
// the struct [RadarEmailTimeseryArcListResponseResultSerie0]
type radarEmailTimeseryArcListResponseResultSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTimeseryArcListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTimeseryArcListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailTimeseryArcListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailTimeseryArcListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailTimeseryArcListParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailTimeseryArcListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailTimeseryArcListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailTimeseryArcListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailTimeseryArcListParams]'s query parameters as
// `url.Values`.
func (r RadarEmailTimeseryArcListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailTimeseryArcListParamsAggInterval string

const (
	RadarEmailTimeseryArcListParamsAggInterval15m RadarEmailTimeseryArcListParamsAggInterval = "15m"
	RadarEmailTimeseryArcListParamsAggInterval1h  RadarEmailTimeseryArcListParamsAggInterval = "1h"
	RadarEmailTimeseryArcListParamsAggInterval1d  RadarEmailTimeseryArcListParamsAggInterval = "1d"
	RadarEmailTimeseryArcListParamsAggInterval1w  RadarEmailTimeseryArcListParamsAggInterval = "1w"
)

type RadarEmailTimeseryArcListParamsDateRange string

const (
	RadarEmailTimeseryArcListParamsDateRange1d         RadarEmailTimeseryArcListParamsDateRange = "1d"
	RadarEmailTimeseryArcListParamsDateRange7d         RadarEmailTimeseryArcListParamsDateRange = "7d"
	RadarEmailTimeseryArcListParamsDateRange14d        RadarEmailTimeseryArcListParamsDateRange = "14d"
	RadarEmailTimeseryArcListParamsDateRange28d        RadarEmailTimeseryArcListParamsDateRange = "28d"
	RadarEmailTimeseryArcListParamsDateRange12w        RadarEmailTimeseryArcListParamsDateRange = "12w"
	RadarEmailTimeseryArcListParamsDateRange24w        RadarEmailTimeseryArcListParamsDateRange = "24w"
	RadarEmailTimeseryArcListParamsDateRange52w        RadarEmailTimeseryArcListParamsDateRange = "52w"
	RadarEmailTimeseryArcListParamsDateRange1dControl  RadarEmailTimeseryArcListParamsDateRange = "1dControl"
	RadarEmailTimeseryArcListParamsDateRange7dControl  RadarEmailTimeseryArcListParamsDateRange = "7dControl"
	RadarEmailTimeseryArcListParamsDateRange14dControl RadarEmailTimeseryArcListParamsDateRange = "14dControl"
	RadarEmailTimeseryArcListParamsDateRange28dControl RadarEmailTimeseryArcListParamsDateRange = "28dControl"
	RadarEmailTimeseryArcListParamsDateRange12wControl RadarEmailTimeseryArcListParamsDateRange = "12wControl"
	RadarEmailTimeseryArcListParamsDateRange24wControl RadarEmailTimeseryArcListParamsDateRange = "24wControl"
)

type RadarEmailTimeseryArcListParamsDkim string

const (
	RadarEmailTimeseryArcListParamsDkimPass RadarEmailTimeseryArcListParamsDkim = "PASS"
	RadarEmailTimeseryArcListParamsDkimNone RadarEmailTimeseryArcListParamsDkim = "NONE"
	RadarEmailTimeseryArcListParamsDkimFail RadarEmailTimeseryArcListParamsDkim = "FAIL"
)

type RadarEmailTimeseryArcListParamsDmarc string

const (
	RadarEmailTimeseryArcListParamsDmarcPass RadarEmailTimeseryArcListParamsDmarc = "PASS"
	RadarEmailTimeseryArcListParamsDmarcNone RadarEmailTimeseryArcListParamsDmarc = "NONE"
	RadarEmailTimeseryArcListParamsDmarcFail RadarEmailTimeseryArcListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailTimeseryArcListParamsFormat string

const (
	RadarEmailTimeseryArcListParamsFormatJson RadarEmailTimeseryArcListParamsFormat = "JSON"
	RadarEmailTimeseryArcListParamsFormatCsv  RadarEmailTimeseryArcListParamsFormat = "CSV"
)

type RadarEmailTimeseryArcListParamsSpf string

const (
	RadarEmailTimeseryArcListParamsSpfPass RadarEmailTimeseryArcListParamsSpf = "PASS"
	RadarEmailTimeseryArcListParamsSpfNone RadarEmailTimeseryArcListParamsSpf = "NONE"
	RadarEmailTimeseryArcListParamsSpfFail RadarEmailTimeseryArcListParamsSpf = "FAIL"
)
