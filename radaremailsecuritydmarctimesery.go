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

// RadarEmailSecurityDmarcTimeseryService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityDmarcTimeseryService] method instead.
type RadarEmailSecurityDmarcTimeseryService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityDmarcTimeseryService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityDmarcTimeseryService(opts ...option.RequestOption) (r *RadarEmailSecurityDmarcTimeseryService) {
	r = &RadarEmailSecurityDmarcTimeseryService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per DMARC validation over time.
func (r *RadarEmailSecurityDmarcTimeseryService) List(ctx context.Context, query RadarEmailSecurityDmarcTimeseryListParams, opts ...option.RequestOption) (res *RadarEmailSecurityDmarcTimeseryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/timeseries_groups/dmarc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSecurityDmarcTimeseryListResponse struct {
	Result  RadarEmailSecurityDmarcTimeseryListResponseResult `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    radarEmailSecurityDmarcTimeseryListResponseJSON   `json:"-"`
}

// radarEmailSecurityDmarcTimeseryListResponseJSON contains the JSON metadata for
// the struct [RadarEmailSecurityDmarcTimeseryListResponse]
type radarEmailSecurityDmarcTimeseryListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityDmarcTimeseryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityDmarcTimeseryListResponseResult struct {
	Meta   interface{}                                             `json:"meta,required"`
	Serie0 RadarEmailSecurityDmarcTimeseryListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecurityDmarcTimeseryListResponseResultJSON   `json:"-"`
}

// radarEmailSecurityDmarcTimeseryListResponseResultJSON contains the JSON metadata
// for the struct [RadarEmailSecurityDmarcTimeseryListResponseResult]
type radarEmailSecurityDmarcTimeseryListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityDmarcTimeseryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityDmarcTimeseryListResponseResultSerie0 struct {
	Fail []string                                                    `json:"FAIL,required"`
	None []string                                                    `json:"NONE,required"`
	Pass []string                                                    `json:"PASS,required"`
	JSON radarEmailSecurityDmarcTimeseryListResponseResultSerie0JSON `json:"-"`
}

// radarEmailSecurityDmarcTimeseryListResponseResultSerie0JSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityDmarcTimeseryListResponseResultSerie0]
type radarEmailSecurityDmarcTimeseryListResponseResultSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityDmarcTimeseryListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityDmarcTimeseryListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityDmarcTimeseryListParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityDmarcTimeseryListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityDmarcTimeseryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSecurityDmarcTimeseryListParamsDkim] `query:"dkim"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityDmarcTimeseryListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailSecurityDmarcTimeseryListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityDmarcTimeseryListParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityDmarcTimeseryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityDmarcTimeseryListParamsAggInterval string

const (
	RadarEmailSecurityDmarcTimeseryListParamsAggInterval15m RadarEmailSecurityDmarcTimeseryListParamsAggInterval = "15m"
	RadarEmailSecurityDmarcTimeseryListParamsAggInterval1h  RadarEmailSecurityDmarcTimeseryListParamsAggInterval = "1h"
	RadarEmailSecurityDmarcTimeseryListParamsAggInterval1d  RadarEmailSecurityDmarcTimeseryListParamsAggInterval = "1d"
	RadarEmailSecurityDmarcTimeseryListParamsAggInterval1w  RadarEmailSecurityDmarcTimeseryListParamsAggInterval = "1w"
)

type RadarEmailSecurityDmarcTimeseryListParamsArc string

const (
	RadarEmailSecurityDmarcTimeseryListParamsArcPass RadarEmailSecurityDmarcTimeseryListParamsArc = "PASS"
	RadarEmailSecurityDmarcTimeseryListParamsArcNone RadarEmailSecurityDmarcTimeseryListParamsArc = "NONE"
	RadarEmailSecurityDmarcTimeseryListParamsArcFail RadarEmailSecurityDmarcTimeseryListParamsArc = "FAIL"
)

type RadarEmailSecurityDmarcTimeseryListParamsDateRange string

const (
	RadarEmailSecurityDmarcTimeseryListParamsDateRange1d         RadarEmailSecurityDmarcTimeseryListParamsDateRange = "1d"
	RadarEmailSecurityDmarcTimeseryListParamsDateRange2d         RadarEmailSecurityDmarcTimeseryListParamsDateRange = "2d"
	RadarEmailSecurityDmarcTimeseryListParamsDateRange7d         RadarEmailSecurityDmarcTimeseryListParamsDateRange = "7d"
	RadarEmailSecurityDmarcTimeseryListParamsDateRange14d        RadarEmailSecurityDmarcTimeseryListParamsDateRange = "14d"
	RadarEmailSecurityDmarcTimeseryListParamsDateRange28d        RadarEmailSecurityDmarcTimeseryListParamsDateRange = "28d"
	RadarEmailSecurityDmarcTimeseryListParamsDateRange12w        RadarEmailSecurityDmarcTimeseryListParamsDateRange = "12w"
	RadarEmailSecurityDmarcTimeseryListParamsDateRange24w        RadarEmailSecurityDmarcTimeseryListParamsDateRange = "24w"
	RadarEmailSecurityDmarcTimeseryListParamsDateRange52w        RadarEmailSecurityDmarcTimeseryListParamsDateRange = "52w"
	RadarEmailSecurityDmarcTimeseryListParamsDateRange1dControl  RadarEmailSecurityDmarcTimeseryListParamsDateRange = "1dControl"
	RadarEmailSecurityDmarcTimeseryListParamsDateRange2dControl  RadarEmailSecurityDmarcTimeseryListParamsDateRange = "2dControl"
	RadarEmailSecurityDmarcTimeseryListParamsDateRange7dControl  RadarEmailSecurityDmarcTimeseryListParamsDateRange = "7dControl"
	RadarEmailSecurityDmarcTimeseryListParamsDateRange14dControl RadarEmailSecurityDmarcTimeseryListParamsDateRange = "14dControl"
	RadarEmailSecurityDmarcTimeseryListParamsDateRange28dControl RadarEmailSecurityDmarcTimeseryListParamsDateRange = "28dControl"
	RadarEmailSecurityDmarcTimeseryListParamsDateRange12wControl RadarEmailSecurityDmarcTimeseryListParamsDateRange = "12wControl"
	RadarEmailSecurityDmarcTimeseryListParamsDateRange24wControl RadarEmailSecurityDmarcTimeseryListParamsDateRange = "24wControl"
)

type RadarEmailSecurityDmarcTimeseryListParamsDkim string

const (
	RadarEmailSecurityDmarcTimeseryListParamsDkimPass RadarEmailSecurityDmarcTimeseryListParamsDkim = "PASS"
	RadarEmailSecurityDmarcTimeseryListParamsDkimNone RadarEmailSecurityDmarcTimeseryListParamsDkim = "NONE"
	RadarEmailSecurityDmarcTimeseryListParamsDkimFail RadarEmailSecurityDmarcTimeseryListParamsDkim = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityDmarcTimeseryListParamsFormat string

const (
	RadarEmailSecurityDmarcTimeseryListParamsFormatJson RadarEmailSecurityDmarcTimeseryListParamsFormat = "JSON"
	RadarEmailSecurityDmarcTimeseryListParamsFormatCsv  RadarEmailSecurityDmarcTimeseryListParamsFormat = "CSV"
)

type RadarEmailSecurityDmarcTimeseryListParamsSpf string

const (
	RadarEmailSecurityDmarcTimeseryListParamsSpfPass RadarEmailSecurityDmarcTimeseryListParamsSpf = "PASS"
	RadarEmailSecurityDmarcTimeseryListParamsSpfNone RadarEmailSecurityDmarcTimeseryListParamsSpf = "NONE"
	RadarEmailSecurityDmarcTimeseryListParamsSpfFail RadarEmailSecurityDmarcTimeseryListParamsSpf = "FAIL"
)
