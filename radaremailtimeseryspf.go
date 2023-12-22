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

// RadarEmailTimeserySpfService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEmailTimeserySpfService]
// method instead.
type RadarEmailTimeserySpfService struct {
	Options []option.RequestOption
}

// NewRadarEmailTimeserySpfService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarEmailTimeserySpfService(opts ...option.RequestOption) (r *RadarEmailTimeserySpfService) {
	r = &RadarEmailTimeserySpfService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per SPF validation over time.
func (r *RadarEmailTimeserySpfService) List(ctx context.Context, query RadarEmailTimeserySpfListParams, opts ...option.RequestOption) (res *RadarEmailTimeserySpfListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/timeseries/spf"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailTimeserySpfListResponse struct {
	Result  RadarEmailTimeserySpfListResponseResult `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    radarEmailTimeserySpfListResponseJSON   `json:"-"`
}

// radarEmailTimeserySpfListResponseJSON contains the JSON metadata for the struct
// [RadarEmailTimeserySpfListResponse]
type radarEmailTimeserySpfListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTimeserySpfListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTimeserySpfListResponseResult struct {
	Meta   interface{}                                   `json:"meta,required"`
	Serie0 RadarEmailTimeserySpfListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailTimeserySpfListResponseResultJSON   `json:"-"`
}

// radarEmailTimeserySpfListResponseResultJSON contains the JSON metadata for the
// struct [RadarEmailTimeserySpfListResponseResult]
type radarEmailTimeserySpfListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTimeserySpfListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTimeserySpfListResponseResultSerie0 struct {
	Fail []string                                          `json:"FAIL,required"`
	None []string                                          `json:"NONE,required"`
	Pass []string                                          `json:"PASS,required"`
	JSON radarEmailTimeserySpfListResponseResultSerie0JSON `json:"-"`
}

// radarEmailTimeserySpfListResponseResultSerie0JSON contains the JSON metadata for
// the struct [RadarEmailTimeserySpfListResponseResultSerie0]
type radarEmailTimeserySpfListResponseResultSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTimeserySpfListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTimeserySpfListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailTimeserySpfListParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailTimeserySpfListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailTimeserySpfListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailTimeserySpfListParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailTimeserySpfListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailTimeserySpfListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarEmailTimeserySpfListParams]'s query parameters as
// `url.Values`.
func (r RadarEmailTimeserySpfListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailTimeserySpfListParamsAggInterval string

const (
	RadarEmailTimeserySpfListParamsAggInterval15m RadarEmailTimeserySpfListParamsAggInterval = "15m"
	RadarEmailTimeserySpfListParamsAggInterval1h  RadarEmailTimeserySpfListParamsAggInterval = "1h"
	RadarEmailTimeserySpfListParamsAggInterval1d  RadarEmailTimeserySpfListParamsAggInterval = "1d"
	RadarEmailTimeserySpfListParamsAggInterval1w  RadarEmailTimeserySpfListParamsAggInterval = "1w"
)

type RadarEmailTimeserySpfListParamsArc string

const (
	RadarEmailTimeserySpfListParamsArcPass RadarEmailTimeserySpfListParamsArc = "PASS"
	RadarEmailTimeserySpfListParamsArcNone RadarEmailTimeserySpfListParamsArc = "NONE"
	RadarEmailTimeserySpfListParamsArcFail RadarEmailTimeserySpfListParamsArc = "FAIL"
)

type RadarEmailTimeserySpfListParamsDateRange string

const (
	RadarEmailTimeserySpfListParamsDateRange1d         RadarEmailTimeserySpfListParamsDateRange = "1d"
	RadarEmailTimeserySpfListParamsDateRange7d         RadarEmailTimeserySpfListParamsDateRange = "7d"
	RadarEmailTimeserySpfListParamsDateRange14d        RadarEmailTimeserySpfListParamsDateRange = "14d"
	RadarEmailTimeserySpfListParamsDateRange28d        RadarEmailTimeserySpfListParamsDateRange = "28d"
	RadarEmailTimeserySpfListParamsDateRange12w        RadarEmailTimeserySpfListParamsDateRange = "12w"
	RadarEmailTimeserySpfListParamsDateRange24w        RadarEmailTimeserySpfListParamsDateRange = "24w"
	RadarEmailTimeserySpfListParamsDateRange52w        RadarEmailTimeserySpfListParamsDateRange = "52w"
	RadarEmailTimeserySpfListParamsDateRange1dControl  RadarEmailTimeserySpfListParamsDateRange = "1dControl"
	RadarEmailTimeserySpfListParamsDateRange7dControl  RadarEmailTimeserySpfListParamsDateRange = "7dControl"
	RadarEmailTimeserySpfListParamsDateRange14dControl RadarEmailTimeserySpfListParamsDateRange = "14dControl"
	RadarEmailTimeserySpfListParamsDateRange28dControl RadarEmailTimeserySpfListParamsDateRange = "28dControl"
	RadarEmailTimeserySpfListParamsDateRange12wControl RadarEmailTimeserySpfListParamsDateRange = "12wControl"
	RadarEmailTimeserySpfListParamsDateRange24wControl RadarEmailTimeserySpfListParamsDateRange = "24wControl"
)

type RadarEmailTimeserySpfListParamsDkim string

const (
	RadarEmailTimeserySpfListParamsDkimPass RadarEmailTimeserySpfListParamsDkim = "PASS"
	RadarEmailTimeserySpfListParamsDkimNone RadarEmailTimeserySpfListParamsDkim = "NONE"
	RadarEmailTimeserySpfListParamsDkimFail RadarEmailTimeserySpfListParamsDkim = "FAIL"
)

type RadarEmailTimeserySpfListParamsDmarc string

const (
	RadarEmailTimeserySpfListParamsDmarcPass RadarEmailTimeserySpfListParamsDmarc = "PASS"
	RadarEmailTimeserySpfListParamsDmarcNone RadarEmailTimeserySpfListParamsDmarc = "NONE"
	RadarEmailTimeserySpfListParamsDmarcFail RadarEmailTimeserySpfListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailTimeserySpfListParamsFormat string

const (
	RadarEmailTimeserySpfListParamsFormatJson RadarEmailTimeserySpfListParamsFormat = "JSON"
	RadarEmailTimeserySpfListParamsFormatCsv  RadarEmailTimeserySpfListParamsFormat = "CSV"
)
