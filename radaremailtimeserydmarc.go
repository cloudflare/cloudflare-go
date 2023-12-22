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

// RadarEmailTimeseryDmarcService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarEmailTimeseryDmarcService] method instead.
type RadarEmailTimeseryDmarcService struct {
	Options []option.RequestOption
}

// NewRadarEmailTimeseryDmarcService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarEmailTimeseryDmarcService(opts ...option.RequestOption) (r *RadarEmailTimeseryDmarcService) {
	r = &RadarEmailTimeseryDmarcService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per DMARC validation over time.
func (r *RadarEmailTimeseryDmarcService) List(ctx context.Context, query RadarEmailTimeseryDmarcListParams, opts ...option.RequestOption) (res *RadarEmailTimeseryDmarcListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/timeseries/dmarc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailTimeseryDmarcListResponse struct {
	Result  RadarEmailTimeseryDmarcListResponseResult `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    radarEmailTimeseryDmarcListResponseJSON   `json:"-"`
}

// radarEmailTimeseryDmarcListResponseJSON contains the JSON metadata for the
// struct [RadarEmailTimeseryDmarcListResponse]
type radarEmailTimeseryDmarcListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTimeseryDmarcListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTimeseryDmarcListResponseResult struct {
	Meta   interface{}                                     `json:"meta,required"`
	Serie0 RadarEmailTimeseryDmarcListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailTimeseryDmarcListResponseResultJSON   `json:"-"`
}

// radarEmailTimeseryDmarcListResponseResultJSON contains the JSON metadata for the
// struct [RadarEmailTimeseryDmarcListResponseResult]
type radarEmailTimeseryDmarcListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTimeseryDmarcListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTimeseryDmarcListResponseResultSerie0 struct {
	Fail []string                                            `json:"FAIL,required"`
	None []string                                            `json:"NONE,required"`
	Pass []string                                            `json:"PASS,required"`
	JSON radarEmailTimeseryDmarcListResponseResultSerie0JSON `json:"-"`
}

// radarEmailTimeseryDmarcListResponseResultSerie0JSON contains the JSON metadata
// for the struct [RadarEmailTimeseryDmarcListResponseResultSerie0]
type radarEmailTimeseryDmarcListResponseResultSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTimeseryDmarcListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTimeseryDmarcListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailTimeseryDmarcListParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailTimeseryDmarcListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailTimeseryDmarcListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailTimeseryDmarcListParamsDkim] `query:"dkim"`
	// Format results are returned in.
	Format param.Field[RadarEmailTimeseryDmarcListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailTimeseryDmarcListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailTimeseryDmarcListParams]'s query parameters as
// `url.Values`.
func (r RadarEmailTimeseryDmarcListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailTimeseryDmarcListParamsAggInterval string

const (
	RadarEmailTimeseryDmarcListParamsAggInterval15m RadarEmailTimeseryDmarcListParamsAggInterval = "15m"
	RadarEmailTimeseryDmarcListParamsAggInterval1h  RadarEmailTimeseryDmarcListParamsAggInterval = "1h"
	RadarEmailTimeseryDmarcListParamsAggInterval1d  RadarEmailTimeseryDmarcListParamsAggInterval = "1d"
	RadarEmailTimeseryDmarcListParamsAggInterval1w  RadarEmailTimeseryDmarcListParamsAggInterval = "1w"
)

type RadarEmailTimeseryDmarcListParamsArc string

const (
	RadarEmailTimeseryDmarcListParamsArcPass RadarEmailTimeseryDmarcListParamsArc = "PASS"
	RadarEmailTimeseryDmarcListParamsArcNone RadarEmailTimeseryDmarcListParamsArc = "NONE"
	RadarEmailTimeseryDmarcListParamsArcFail RadarEmailTimeseryDmarcListParamsArc = "FAIL"
)

type RadarEmailTimeseryDmarcListParamsDateRange string

const (
	RadarEmailTimeseryDmarcListParamsDateRange1d         RadarEmailTimeseryDmarcListParamsDateRange = "1d"
	RadarEmailTimeseryDmarcListParamsDateRange7d         RadarEmailTimeseryDmarcListParamsDateRange = "7d"
	RadarEmailTimeseryDmarcListParamsDateRange14d        RadarEmailTimeseryDmarcListParamsDateRange = "14d"
	RadarEmailTimeseryDmarcListParamsDateRange28d        RadarEmailTimeseryDmarcListParamsDateRange = "28d"
	RadarEmailTimeseryDmarcListParamsDateRange12w        RadarEmailTimeseryDmarcListParamsDateRange = "12w"
	RadarEmailTimeseryDmarcListParamsDateRange24w        RadarEmailTimeseryDmarcListParamsDateRange = "24w"
	RadarEmailTimeseryDmarcListParamsDateRange52w        RadarEmailTimeseryDmarcListParamsDateRange = "52w"
	RadarEmailTimeseryDmarcListParamsDateRange1dControl  RadarEmailTimeseryDmarcListParamsDateRange = "1dControl"
	RadarEmailTimeseryDmarcListParamsDateRange7dControl  RadarEmailTimeseryDmarcListParamsDateRange = "7dControl"
	RadarEmailTimeseryDmarcListParamsDateRange14dControl RadarEmailTimeseryDmarcListParamsDateRange = "14dControl"
	RadarEmailTimeseryDmarcListParamsDateRange28dControl RadarEmailTimeseryDmarcListParamsDateRange = "28dControl"
	RadarEmailTimeseryDmarcListParamsDateRange12wControl RadarEmailTimeseryDmarcListParamsDateRange = "12wControl"
	RadarEmailTimeseryDmarcListParamsDateRange24wControl RadarEmailTimeseryDmarcListParamsDateRange = "24wControl"
)

type RadarEmailTimeseryDmarcListParamsDkim string

const (
	RadarEmailTimeseryDmarcListParamsDkimPass RadarEmailTimeseryDmarcListParamsDkim = "PASS"
	RadarEmailTimeseryDmarcListParamsDkimNone RadarEmailTimeseryDmarcListParamsDkim = "NONE"
	RadarEmailTimeseryDmarcListParamsDkimFail RadarEmailTimeseryDmarcListParamsDkim = "FAIL"
)

// Format results are returned in.
type RadarEmailTimeseryDmarcListParamsFormat string

const (
	RadarEmailTimeseryDmarcListParamsFormatJson RadarEmailTimeseryDmarcListParamsFormat = "JSON"
	RadarEmailTimeseryDmarcListParamsFormatCsv  RadarEmailTimeseryDmarcListParamsFormat = "CSV"
)

type RadarEmailTimeseryDmarcListParamsSpf string

const (
	RadarEmailTimeseryDmarcListParamsSpfPass RadarEmailTimeseryDmarcListParamsSpf = "PASS"
	RadarEmailTimeseryDmarcListParamsSpfNone RadarEmailTimeseryDmarcListParamsSpf = "NONE"
	RadarEmailTimeseryDmarcListParamsSpfFail RadarEmailTimeseryDmarcListParamsSpf = "FAIL"
)
