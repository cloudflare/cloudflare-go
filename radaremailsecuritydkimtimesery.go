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

// RadarEmailSecurityDkimTimeseryService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityDkimTimeseryService] method instead.
type RadarEmailSecurityDkimTimeseryService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityDkimTimeseryService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityDkimTimeseryService(opts ...option.RequestOption) (r *RadarEmailSecurityDkimTimeseryService) {
	r = &RadarEmailSecurityDkimTimeseryService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per DKIM validation over time.
func (r *RadarEmailSecurityDkimTimeseryService) List(ctx context.Context, query RadarEmailSecurityDkimTimeseryListParams, opts ...option.RequestOption) (res *RadarEmailSecurityDkimTimeseryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/timeseries_groups/dkim"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSecurityDkimTimeseryListResponse struct {
	Result  RadarEmailSecurityDkimTimeseryListResponseResult `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarEmailSecurityDkimTimeseryListResponseJSON   `json:"-"`
}

// radarEmailSecurityDkimTimeseryListResponseJSON contains the JSON metadata for
// the struct [RadarEmailSecurityDkimTimeseryListResponse]
type radarEmailSecurityDkimTimeseryListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityDkimTimeseryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityDkimTimeseryListResponseResult struct {
	Meta   interface{}                                            `json:"meta,required"`
	Serie0 RadarEmailSecurityDkimTimeseryListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecurityDkimTimeseryListResponseResultJSON   `json:"-"`
}

// radarEmailSecurityDkimTimeseryListResponseResultJSON contains the JSON metadata
// for the struct [RadarEmailSecurityDkimTimeseryListResponseResult]
type radarEmailSecurityDkimTimeseryListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityDkimTimeseryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityDkimTimeseryListResponseResultSerie0 struct {
	Fail []string                                                   `json:"FAIL,required"`
	None []string                                                   `json:"NONE,required"`
	Pass []string                                                   `json:"PASS,required"`
	JSON radarEmailSecurityDkimTimeseryListResponseResultSerie0JSON `json:"-"`
}

// radarEmailSecurityDkimTimeseryListResponseResultSerie0JSON contains the JSON
// metadata for the struct [RadarEmailSecurityDkimTimeseryListResponseResultSerie0]
type radarEmailSecurityDkimTimeseryListResponseResultSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityDkimTimeseryListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityDkimTimeseryListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityDkimTimeseryListParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityDkimTimeseryListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityDkimTimeseryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecurityDkimTimeseryListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityDkimTimeseryListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailSecurityDkimTimeseryListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityDkimTimeseryListParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityDkimTimeseryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityDkimTimeseryListParamsAggInterval string

const (
	RadarEmailSecurityDkimTimeseryListParamsAggInterval15m RadarEmailSecurityDkimTimeseryListParamsAggInterval = "15m"
	RadarEmailSecurityDkimTimeseryListParamsAggInterval1h  RadarEmailSecurityDkimTimeseryListParamsAggInterval = "1h"
	RadarEmailSecurityDkimTimeseryListParamsAggInterval1d  RadarEmailSecurityDkimTimeseryListParamsAggInterval = "1d"
	RadarEmailSecurityDkimTimeseryListParamsAggInterval1w  RadarEmailSecurityDkimTimeseryListParamsAggInterval = "1w"
)

type RadarEmailSecurityDkimTimeseryListParamsArc string

const (
	RadarEmailSecurityDkimTimeseryListParamsArcPass RadarEmailSecurityDkimTimeseryListParamsArc = "PASS"
	RadarEmailSecurityDkimTimeseryListParamsArcNone RadarEmailSecurityDkimTimeseryListParamsArc = "NONE"
	RadarEmailSecurityDkimTimeseryListParamsArcFail RadarEmailSecurityDkimTimeseryListParamsArc = "FAIL"
)

type RadarEmailSecurityDkimTimeseryListParamsDateRange string

const (
	RadarEmailSecurityDkimTimeseryListParamsDateRange1d         RadarEmailSecurityDkimTimeseryListParamsDateRange = "1d"
	RadarEmailSecurityDkimTimeseryListParamsDateRange2d         RadarEmailSecurityDkimTimeseryListParamsDateRange = "2d"
	RadarEmailSecurityDkimTimeseryListParamsDateRange7d         RadarEmailSecurityDkimTimeseryListParamsDateRange = "7d"
	RadarEmailSecurityDkimTimeseryListParamsDateRange14d        RadarEmailSecurityDkimTimeseryListParamsDateRange = "14d"
	RadarEmailSecurityDkimTimeseryListParamsDateRange28d        RadarEmailSecurityDkimTimeseryListParamsDateRange = "28d"
	RadarEmailSecurityDkimTimeseryListParamsDateRange12w        RadarEmailSecurityDkimTimeseryListParamsDateRange = "12w"
	RadarEmailSecurityDkimTimeseryListParamsDateRange24w        RadarEmailSecurityDkimTimeseryListParamsDateRange = "24w"
	RadarEmailSecurityDkimTimeseryListParamsDateRange52w        RadarEmailSecurityDkimTimeseryListParamsDateRange = "52w"
	RadarEmailSecurityDkimTimeseryListParamsDateRange1dControl  RadarEmailSecurityDkimTimeseryListParamsDateRange = "1dControl"
	RadarEmailSecurityDkimTimeseryListParamsDateRange2dControl  RadarEmailSecurityDkimTimeseryListParamsDateRange = "2dControl"
	RadarEmailSecurityDkimTimeseryListParamsDateRange7dControl  RadarEmailSecurityDkimTimeseryListParamsDateRange = "7dControl"
	RadarEmailSecurityDkimTimeseryListParamsDateRange14dControl RadarEmailSecurityDkimTimeseryListParamsDateRange = "14dControl"
	RadarEmailSecurityDkimTimeseryListParamsDateRange28dControl RadarEmailSecurityDkimTimeseryListParamsDateRange = "28dControl"
	RadarEmailSecurityDkimTimeseryListParamsDateRange12wControl RadarEmailSecurityDkimTimeseryListParamsDateRange = "12wControl"
	RadarEmailSecurityDkimTimeseryListParamsDateRange24wControl RadarEmailSecurityDkimTimeseryListParamsDateRange = "24wControl"
)

type RadarEmailSecurityDkimTimeseryListParamsDmarc string

const (
	RadarEmailSecurityDkimTimeseryListParamsDmarcPass RadarEmailSecurityDkimTimeseryListParamsDmarc = "PASS"
	RadarEmailSecurityDkimTimeseryListParamsDmarcNone RadarEmailSecurityDkimTimeseryListParamsDmarc = "NONE"
	RadarEmailSecurityDkimTimeseryListParamsDmarcFail RadarEmailSecurityDkimTimeseryListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityDkimTimeseryListParamsFormat string

const (
	RadarEmailSecurityDkimTimeseryListParamsFormatJson RadarEmailSecurityDkimTimeseryListParamsFormat = "JSON"
	RadarEmailSecurityDkimTimeseryListParamsFormatCsv  RadarEmailSecurityDkimTimeseryListParamsFormat = "CSV"
)

type RadarEmailSecurityDkimTimeseryListParamsSpf string

const (
	RadarEmailSecurityDkimTimeseryListParamsSpfPass RadarEmailSecurityDkimTimeseryListParamsSpf = "PASS"
	RadarEmailSecurityDkimTimeseryListParamsSpfNone RadarEmailSecurityDkimTimeseryListParamsSpf = "NONE"
	RadarEmailSecurityDkimTimeseryListParamsSpfFail RadarEmailSecurityDkimTimeseryListParamsSpf = "FAIL"
)
