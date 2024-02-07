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

// RadarEmailSecuritySpfService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEmailSecuritySpfService]
// method instead.
type RadarEmailSecuritySpfService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecuritySpfService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarEmailSecuritySpfService(opts ...option.RequestOption) (r *RadarEmailSecuritySpfService) {
	r = &RadarEmailSecuritySpfService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per SPF validation over time.
func (r *RadarEmailSecuritySpfService) List(ctx context.Context, query RadarEmailSecuritySpfListParams, opts ...option.RequestOption) (res *RadarEmailSecuritySpfListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/timeseries_groups/spf"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSecuritySpfListResponse struct {
	Result  RadarEmailSecuritySpfListResponseResult `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    radarEmailSecuritySpfListResponseJSON   `json:"-"`
}

// radarEmailSecuritySpfListResponseJSON contains the JSON metadata for the struct
// [RadarEmailSecuritySpfListResponse]
type radarEmailSecuritySpfListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySpfListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySpfListResponseResult struct {
	Meta   interface{}                                   `json:"meta,required"`
	Serie0 RadarEmailSecuritySpfListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecuritySpfListResponseResultJSON   `json:"-"`
}

// radarEmailSecuritySpfListResponseResultJSON contains the JSON metadata for the
// struct [RadarEmailSecuritySpfListResponseResult]
type radarEmailSecuritySpfListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySpfListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySpfListResponseResultSerie0 struct {
	Fail []string                                          `json:"FAIL,required"`
	None []string                                          `json:"NONE,required"`
	Pass []string                                          `json:"PASS,required"`
	JSON radarEmailSecuritySpfListResponseResultSerie0JSON `json:"-"`
}

// radarEmailSecuritySpfListResponseResultSerie0JSON contains the JSON metadata for
// the struct [RadarEmailSecuritySpfListResponseResultSerie0]
type radarEmailSecuritySpfListResponseResultSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySpfListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySpfListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecuritySpfListParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecuritySpfListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecuritySpfListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSecuritySpfListParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecuritySpfListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecuritySpfListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarEmailSecuritySpfListParams]'s query parameters as
// `url.Values`.
func (r RadarEmailSecuritySpfListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecuritySpfListParamsAggInterval string

const (
	RadarEmailSecuritySpfListParamsAggInterval15m RadarEmailSecuritySpfListParamsAggInterval = "15m"
	RadarEmailSecuritySpfListParamsAggInterval1h  RadarEmailSecuritySpfListParamsAggInterval = "1h"
	RadarEmailSecuritySpfListParamsAggInterval1d  RadarEmailSecuritySpfListParamsAggInterval = "1d"
	RadarEmailSecuritySpfListParamsAggInterval1w  RadarEmailSecuritySpfListParamsAggInterval = "1w"
)

type RadarEmailSecuritySpfListParamsArc string

const (
	RadarEmailSecuritySpfListParamsArcPass RadarEmailSecuritySpfListParamsArc = "PASS"
	RadarEmailSecuritySpfListParamsArcNone RadarEmailSecuritySpfListParamsArc = "NONE"
	RadarEmailSecuritySpfListParamsArcFail RadarEmailSecuritySpfListParamsArc = "FAIL"
)

type RadarEmailSecuritySpfListParamsDateRange string

const (
	RadarEmailSecuritySpfListParamsDateRange1d         RadarEmailSecuritySpfListParamsDateRange = "1d"
	RadarEmailSecuritySpfListParamsDateRange2d         RadarEmailSecuritySpfListParamsDateRange = "2d"
	RadarEmailSecuritySpfListParamsDateRange7d         RadarEmailSecuritySpfListParamsDateRange = "7d"
	RadarEmailSecuritySpfListParamsDateRange14d        RadarEmailSecuritySpfListParamsDateRange = "14d"
	RadarEmailSecuritySpfListParamsDateRange28d        RadarEmailSecuritySpfListParamsDateRange = "28d"
	RadarEmailSecuritySpfListParamsDateRange12w        RadarEmailSecuritySpfListParamsDateRange = "12w"
	RadarEmailSecuritySpfListParamsDateRange24w        RadarEmailSecuritySpfListParamsDateRange = "24w"
	RadarEmailSecuritySpfListParamsDateRange52w        RadarEmailSecuritySpfListParamsDateRange = "52w"
	RadarEmailSecuritySpfListParamsDateRange1dControl  RadarEmailSecuritySpfListParamsDateRange = "1dControl"
	RadarEmailSecuritySpfListParamsDateRange2dControl  RadarEmailSecuritySpfListParamsDateRange = "2dControl"
	RadarEmailSecuritySpfListParamsDateRange7dControl  RadarEmailSecuritySpfListParamsDateRange = "7dControl"
	RadarEmailSecuritySpfListParamsDateRange14dControl RadarEmailSecuritySpfListParamsDateRange = "14dControl"
	RadarEmailSecuritySpfListParamsDateRange28dControl RadarEmailSecuritySpfListParamsDateRange = "28dControl"
	RadarEmailSecuritySpfListParamsDateRange12wControl RadarEmailSecuritySpfListParamsDateRange = "12wControl"
	RadarEmailSecuritySpfListParamsDateRange24wControl RadarEmailSecuritySpfListParamsDateRange = "24wControl"
)

type RadarEmailSecuritySpfListParamsDkim string

const (
	RadarEmailSecuritySpfListParamsDkimPass RadarEmailSecuritySpfListParamsDkim = "PASS"
	RadarEmailSecuritySpfListParamsDkimNone RadarEmailSecuritySpfListParamsDkim = "NONE"
	RadarEmailSecuritySpfListParamsDkimFail RadarEmailSecuritySpfListParamsDkim = "FAIL"
)

type RadarEmailSecuritySpfListParamsDmarc string

const (
	RadarEmailSecuritySpfListParamsDmarcPass RadarEmailSecuritySpfListParamsDmarc = "PASS"
	RadarEmailSecuritySpfListParamsDmarcNone RadarEmailSecuritySpfListParamsDmarc = "NONE"
	RadarEmailSecuritySpfListParamsDmarcFail RadarEmailSecuritySpfListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecuritySpfListParamsFormat string

const (
	RadarEmailSecuritySpfListParamsFormatJson RadarEmailSecuritySpfListParamsFormat = "JSON"
	RadarEmailSecuritySpfListParamsFormatCsv  RadarEmailSecuritySpfListParamsFormat = "CSV"
)
