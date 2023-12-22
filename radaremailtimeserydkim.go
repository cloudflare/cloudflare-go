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

// RadarEmailTimeseryDkimService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEmailTimeseryDkimService]
// method instead.
type RadarEmailTimeseryDkimService struct {
	Options []option.RequestOption
}

// NewRadarEmailTimeseryDkimService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarEmailTimeseryDkimService(opts ...option.RequestOption) (r *RadarEmailTimeseryDkimService) {
	r = &RadarEmailTimeseryDkimService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per DKIM validation over time.
func (r *RadarEmailTimeseryDkimService) List(ctx context.Context, query RadarEmailTimeseryDkimListParams, opts ...option.RequestOption) (res *RadarEmailTimeseryDkimListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/timeseries/dkim"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailTimeseryDkimListResponse struct {
	Result  RadarEmailTimeseryDkimListResponseResult `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    radarEmailTimeseryDkimListResponseJSON   `json:"-"`
}

// radarEmailTimeseryDkimListResponseJSON contains the JSON metadata for the struct
// [RadarEmailTimeseryDkimListResponse]
type radarEmailTimeseryDkimListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTimeseryDkimListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTimeseryDkimListResponseResult struct {
	Meta   interface{}                                    `json:"meta,required"`
	Serie0 RadarEmailTimeseryDkimListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailTimeseryDkimListResponseResultJSON   `json:"-"`
}

// radarEmailTimeseryDkimListResponseResultJSON contains the JSON metadata for the
// struct [RadarEmailTimeseryDkimListResponseResult]
type radarEmailTimeseryDkimListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTimeseryDkimListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTimeseryDkimListResponseResultSerie0 struct {
	Fail []string                                           `json:"FAIL,required"`
	None []string                                           `json:"NONE,required"`
	Pass []string                                           `json:"PASS,required"`
	JSON radarEmailTimeseryDkimListResponseResultSerie0JSON `json:"-"`
}

// radarEmailTimeseryDkimListResponseResultSerie0JSON contains the JSON metadata
// for the struct [RadarEmailTimeseryDkimListResponseResultSerie0]
type radarEmailTimeseryDkimListResponseResultSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTimeseryDkimListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTimeseryDkimListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailTimeseryDkimListParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailTimeseryDkimListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailTimeseryDkimListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailTimeseryDkimListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailTimeseryDkimListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailTimeseryDkimListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailTimeseryDkimListParams]'s query parameters as
// `url.Values`.
func (r RadarEmailTimeseryDkimListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailTimeseryDkimListParamsAggInterval string

const (
	RadarEmailTimeseryDkimListParamsAggInterval15m RadarEmailTimeseryDkimListParamsAggInterval = "15m"
	RadarEmailTimeseryDkimListParamsAggInterval1h  RadarEmailTimeseryDkimListParamsAggInterval = "1h"
	RadarEmailTimeseryDkimListParamsAggInterval1d  RadarEmailTimeseryDkimListParamsAggInterval = "1d"
	RadarEmailTimeseryDkimListParamsAggInterval1w  RadarEmailTimeseryDkimListParamsAggInterval = "1w"
)

type RadarEmailTimeseryDkimListParamsArc string

const (
	RadarEmailTimeseryDkimListParamsArcPass RadarEmailTimeseryDkimListParamsArc = "PASS"
	RadarEmailTimeseryDkimListParamsArcNone RadarEmailTimeseryDkimListParamsArc = "NONE"
	RadarEmailTimeseryDkimListParamsArcFail RadarEmailTimeseryDkimListParamsArc = "FAIL"
)

type RadarEmailTimeseryDkimListParamsDateRange string

const (
	RadarEmailTimeseryDkimListParamsDateRange1d         RadarEmailTimeseryDkimListParamsDateRange = "1d"
	RadarEmailTimeseryDkimListParamsDateRange7d         RadarEmailTimeseryDkimListParamsDateRange = "7d"
	RadarEmailTimeseryDkimListParamsDateRange14d        RadarEmailTimeseryDkimListParamsDateRange = "14d"
	RadarEmailTimeseryDkimListParamsDateRange28d        RadarEmailTimeseryDkimListParamsDateRange = "28d"
	RadarEmailTimeseryDkimListParamsDateRange12w        RadarEmailTimeseryDkimListParamsDateRange = "12w"
	RadarEmailTimeseryDkimListParamsDateRange24w        RadarEmailTimeseryDkimListParamsDateRange = "24w"
	RadarEmailTimeseryDkimListParamsDateRange52w        RadarEmailTimeseryDkimListParamsDateRange = "52w"
	RadarEmailTimeseryDkimListParamsDateRange1dControl  RadarEmailTimeseryDkimListParamsDateRange = "1dControl"
	RadarEmailTimeseryDkimListParamsDateRange7dControl  RadarEmailTimeseryDkimListParamsDateRange = "7dControl"
	RadarEmailTimeseryDkimListParamsDateRange14dControl RadarEmailTimeseryDkimListParamsDateRange = "14dControl"
	RadarEmailTimeseryDkimListParamsDateRange28dControl RadarEmailTimeseryDkimListParamsDateRange = "28dControl"
	RadarEmailTimeseryDkimListParamsDateRange12wControl RadarEmailTimeseryDkimListParamsDateRange = "12wControl"
	RadarEmailTimeseryDkimListParamsDateRange24wControl RadarEmailTimeseryDkimListParamsDateRange = "24wControl"
)

type RadarEmailTimeseryDkimListParamsDmarc string

const (
	RadarEmailTimeseryDkimListParamsDmarcPass RadarEmailTimeseryDkimListParamsDmarc = "PASS"
	RadarEmailTimeseryDkimListParamsDmarcNone RadarEmailTimeseryDkimListParamsDmarc = "NONE"
	RadarEmailTimeseryDkimListParamsDmarcFail RadarEmailTimeseryDkimListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailTimeseryDkimListParamsFormat string

const (
	RadarEmailTimeseryDkimListParamsFormatJson RadarEmailTimeseryDkimListParamsFormat = "JSON"
	RadarEmailTimeseryDkimListParamsFormatCsv  RadarEmailTimeseryDkimListParamsFormat = "CSV"
)

type RadarEmailTimeseryDkimListParamsSpf string

const (
	RadarEmailTimeseryDkimListParamsSpfPass RadarEmailTimeseryDkimListParamsSpf = "PASS"
	RadarEmailTimeseryDkimListParamsSpfNone RadarEmailTimeseryDkimListParamsSpf = "NONE"
	RadarEmailTimeseryDkimListParamsSpfFail RadarEmailTimeseryDkimListParamsSpf = "FAIL"
)
