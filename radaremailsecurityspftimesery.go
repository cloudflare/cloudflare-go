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

// RadarEmailSecuritySpfTimeseryService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecuritySpfTimeseryService] method instead.
type RadarEmailSecuritySpfTimeseryService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecuritySpfTimeseryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecuritySpfTimeseryService(opts ...option.RequestOption) (r *RadarEmailSecuritySpfTimeseryService) {
	r = &RadarEmailSecuritySpfTimeseryService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per SPF validation over time.
func (r *RadarEmailSecuritySpfTimeseryService) List(ctx context.Context, query RadarEmailSecuritySpfTimeseryListParams, opts ...option.RequestOption) (res *RadarEmailSecuritySpfTimeseryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/timeseries_groups/spf"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSecuritySpfTimeseryListResponse struct {
	Result  RadarEmailSecuritySpfTimeseryListResponseResult `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarEmailSecuritySpfTimeseryListResponseJSON   `json:"-"`
}

// radarEmailSecuritySpfTimeseryListResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecuritySpfTimeseryListResponse]
type radarEmailSecuritySpfTimeseryListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySpfTimeseryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySpfTimeseryListResponseResult struct {
	Meta   interface{}                                           `json:"meta,required"`
	Serie0 RadarEmailSecuritySpfTimeseryListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecuritySpfTimeseryListResponseResultJSON   `json:"-"`
}

// radarEmailSecuritySpfTimeseryListResponseResultJSON contains the JSON metadata
// for the struct [RadarEmailSecuritySpfTimeseryListResponseResult]
type radarEmailSecuritySpfTimeseryListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySpfTimeseryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySpfTimeseryListResponseResultSerie0 struct {
	Fail []string                                                  `json:"FAIL,required"`
	None []string                                                  `json:"NONE,required"`
	Pass []string                                                  `json:"PASS,required"`
	JSON radarEmailSecuritySpfTimeseryListResponseResultSerie0JSON `json:"-"`
}

// radarEmailSecuritySpfTimeseryListResponseResultSerie0JSON contains the JSON
// metadata for the struct [RadarEmailSecuritySpfTimeseryListResponseResultSerie0]
type radarEmailSecuritySpfTimeseryListResponseResultSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySpfTimeseryListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySpfTimeseryListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecuritySpfTimeseryListParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecuritySpfTimeseryListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecuritySpfTimeseryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSecuritySpfTimeseryListParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecuritySpfTimeseryListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecuritySpfTimeseryListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarEmailSecuritySpfTimeseryListParams]'s query parameters
// as `url.Values`.
func (r RadarEmailSecuritySpfTimeseryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecuritySpfTimeseryListParamsAggInterval string

const (
	RadarEmailSecuritySpfTimeseryListParamsAggInterval15m RadarEmailSecuritySpfTimeseryListParamsAggInterval = "15m"
	RadarEmailSecuritySpfTimeseryListParamsAggInterval1h  RadarEmailSecuritySpfTimeseryListParamsAggInterval = "1h"
	RadarEmailSecuritySpfTimeseryListParamsAggInterval1d  RadarEmailSecuritySpfTimeseryListParamsAggInterval = "1d"
	RadarEmailSecuritySpfTimeseryListParamsAggInterval1w  RadarEmailSecuritySpfTimeseryListParamsAggInterval = "1w"
)

type RadarEmailSecuritySpfTimeseryListParamsArc string

const (
	RadarEmailSecuritySpfTimeseryListParamsArcPass RadarEmailSecuritySpfTimeseryListParamsArc = "PASS"
	RadarEmailSecuritySpfTimeseryListParamsArcNone RadarEmailSecuritySpfTimeseryListParamsArc = "NONE"
	RadarEmailSecuritySpfTimeseryListParamsArcFail RadarEmailSecuritySpfTimeseryListParamsArc = "FAIL"
)

type RadarEmailSecuritySpfTimeseryListParamsDateRange string

const (
	RadarEmailSecuritySpfTimeseryListParamsDateRange1d         RadarEmailSecuritySpfTimeseryListParamsDateRange = "1d"
	RadarEmailSecuritySpfTimeseryListParamsDateRange2d         RadarEmailSecuritySpfTimeseryListParamsDateRange = "2d"
	RadarEmailSecuritySpfTimeseryListParamsDateRange7d         RadarEmailSecuritySpfTimeseryListParamsDateRange = "7d"
	RadarEmailSecuritySpfTimeseryListParamsDateRange14d        RadarEmailSecuritySpfTimeseryListParamsDateRange = "14d"
	RadarEmailSecuritySpfTimeseryListParamsDateRange28d        RadarEmailSecuritySpfTimeseryListParamsDateRange = "28d"
	RadarEmailSecuritySpfTimeseryListParamsDateRange12w        RadarEmailSecuritySpfTimeseryListParamsDateRange = "12w"
	RadarEmailSecuritySpfTimeseryListParamsDateRange24w        RadarEmailSecuritySpfTimeseryListParamsDateRange = "24w"
	RadarEmailSecuritySpfTimeseryListParamsDateRange52w        RadarEmailSecuritySpfTimeseryListParamsDateRange = "52w"
	RadarEmailSecuritySpfTimeseryListParamsDateRange1dControl  RadarEmailSecuritySpfTimeseryListParamsDateRange = "1dControl"
	RadarEmailSecuritySpfTimeseryListParamsDateRange2dControl  RadarEmailSecuritySpfTimeseryListParamsDateRange = "2dControl"
	RadarEmailSecuritySpfTimeseryListParamsDateRange7dControl  RadarEmailSecuritySpfTimeseryListParamsDateRange = "7dControl"
	RadarEmailSecuritySpfTimeseryListParamsDateRange14dControl RadarEmailSecuritySpfTimeseryListParamsDateRange = "14dControl"
	RadarEmailSecuritySpfTimeseryListParamsDateRange28dControl RadarEmailSecuritySpfTimeseryListParamsDateRange = "28dControl"
	RadarEmailSecuritySpfTimeseryListParamsDateRange12wControl RadarEmailSecuritySpfTimeseryListParamsDateRange = "12wControl"
	RadarEmailSecuritySpfTimeseryListParamsDateRange24wControl RadarEmailSecuritySpfTimeseryListParamsDateRange = "24wControl"
)

type RadarEmailSecuritySpfTimeseryListParamsDkim string

const (
	RadarEmailSecuritySpfTimeseryListParamsDkimPass RadarEmailSecuritySpfTimeseryListParamsDkim = "PASS"
	RadarEmailSecuritySpfTimeseryListParamsDkimNone RadarEmailSecuritySpfTimeseryListParamsDkim = "NONE"
	RadarEmailSecuritySpfTimeseryListParamsDkimFail RadarEmailSecuritySpfTimeseryListParamsDkim = "FAIL"
)

type RadarEmailSecuritySpfTimeseryListParamsDmarc string

const (
	RadarEmailSecuritySpfTimeseryListParamsDmarcPass RadarEmailSecuritySpfTimeseryListParamsDmarc = "PASS"
	RadarEmailSecuritySpfTimeseryListParamsDmarcNone RadarEmailSecuritySpfTimeseryListParamsDmarc = "NONE"
	RadarEmailSecuritySpfTimeseryListParamsDmarcFail RadarEmailSecuritySpfTimeseryListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecuritySpfTimeseryListParamsFormat string

const (
	RadarEmailSecuritySpfTimeseryListParamsFormatJson RadarEmailSecuritySpfTimeseryListParamsFormat = "JSON"
	RadarEmailSecuritySpfTimeseryListParamsFormatCsv  RadarEmailSecuritySpfTimeseryListParamsFormat = "CSV"
)
