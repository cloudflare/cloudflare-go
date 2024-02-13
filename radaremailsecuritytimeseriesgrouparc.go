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

// RadarEmailSecurityTimeseriesGroupArcService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTimeseriesGroupArcService] method instead.
type RadarEmailSecurityTimeseriesGroupArcService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTimeseriesGroupArcService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarEmailSecurityTimeseriesGroupArcService(opts ...option.RequestOption) (r *RadarEmailSecurityTimeseriesGroupArcService) {
	r = &RadarEmailSecurityTimeseriesGroupArcService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per Arc validation over time.
func (r *RadarEmailSecurityTimeseriesGroupArcService) List(ctx context.Context, query RadarEmailSecurityTimeseriesGroupArcListParams, opts ...option.RequestOption) (res *RadarEmailSecurityTimeseriesGroupArcListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTimeseriesGroupArcListResponseEnvelope
	path := "radar/email/security/timeseries_groups/arc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecurityTimeseriesGroupArcListResponse struct {
	Meta   interface{}                                            `json:"meta,required"`
	Serie0 RadarEmailSecurityTimeseriesGroupArcListResponseSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecurityTimeseriesGroupArcListResponseJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupArcListResponseJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTimeseriesGroupArcListResponse]
type radarEmailSecurityTimeseriesGroupArcListResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupArcListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTimeseriesGroupArcListResponseSerie0 struct {
	Fail []string                                                   `json:"FAIL,required"`
	None []string                                                   `json:"NONE,required"`
	Pass []string                                                   `json:"PASS,required"`
	JSON radarEmailSecurityTimeseriesGroupArcListResponseSerie0JSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupArcListResponseSerie0JSON contains the JSON
// metadata for the struct [RadarEmailSecurityTimeseriesGroupArcListResponseSerie0]
type radarEmailSecurityTimeseriesGroupArcListResponseSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupArcListResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTimeseriesGroupArcListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTimeseriesGroupArcListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTimeseriesGroupArcListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityTimeseriesGroupArcListParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecurityTimeseriesGroupArcListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTimeseriesGroupArcListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTimeseriesGroupArcListParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTimeseriesGroupArcListParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTimeseriesGroupArcListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupArcListParamsAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupArcListParamsAggInterval15m RadarEmailSecurityTimeseriesGroupArcListParamsAggInterval = "15m"
	RadarEmailSecurityTimeseriesGroupArcListParamsAggInterval1h  RadarEmailSecurityTimeseriesGroupArcListParamsAggInterval = "1h"
	RadarEmailSecurityTimeseriesGroupArcListParamsAggInterval1d  RadarEmailSecurityTimeseriesGroupArcListParamsAggInterval = "1d"
	RadarEmailSecurityTimeseriesGroupArcListParamsAggInterval1w  RadarEmailSecurityTimeseriesGroupArcListParamsAggInterval = "1w"
)

type RadarEmailSecurityTimeseriesGroupArcListParamsDateRange string

const (
	RadarEmailSecurityTimeseriesGroupArcListParamsDateRange1d         RadarEmailSecurityTimeseriesGroupArcListParamsDateRange = "1d"
	RadarEmailSecurityTimeseriesGroupArcListParamsDateRange2d         RadarEmailSecurityTimeseriesGroupArcListParamsDateRange = "2d"
	RadarEmailSecurityTimeseriesGroupArcListParamsDateRange7d         RadarEmailSecurityTimeseriesGroupArcListParamsDateRange = "7d"
	RadarEmailSecurityTimeseriesGroupArcListParamsDateRange14d        RadarEmailSecurityTimeseriesGroupArcListParamsDateRange = "14d"
	RadarEmailSecurityTimeseriesGroupArcListParamsDateRange28d        RadarEmailSecurityTimeseriesGroupArcListParamsDateRange = "28d"
	RadarEmailSecurityTimeseriesGroupArcListParamsDateRange12w        RadarEmailSecurityTimeseriesGroupArcListParamsDateRange = "12w"
	RadarEmailSecurityTimeseriesGroupArcListParamsDateRange24w        RadarEmailSecurityTimeseriesGroupArcListParamsDateRange = "24w"
	RadarEmailSecurityTimeseriesGroupArcListParamsDateRange52w        RadarEmailSecurityTimeseriesGroupArcListParamsDateRange = "52w"
	RadarEmailSecurityTimeseriesGroupArcListParamsDateRange1dControl  RadarEmailSecurityTimeseriesGroupArcListParamsDateRange = "1dControl"
	RadarEmailSecurityTimeseriesGroupArcListParamsDateRange2dControl  RadarEmailSecurityTimeseriesGroupArcListParamsDateRange = "2dControl"
	RadarEmailSecurityTimeseriesGroupArcListParamsDateRange7dControl  RadarEmailSecurityTimeseriesGroupArcListParamsDateRange = "7dControl"
	RadarEmailSecurityTimeseriesGroupArcListParamsDateRange14dControl RadarEmailSecurityTimeseriesGroupArcListParamsDateRange = "14dControl"
	RadarEmailSecurityTimeseriesGroupArcListParamsDateRange28dControl RadarEmailSecurityTimeseriesGroupArcListParamsDateRange = "28dControl"
	RadarEmailSecurityTimeseriesGroupArcListParamsDateRange12wControl RadarEmailSecurityTimeseriesGroupArcListParamsDateRange = "12wControl"
	RadarEmailSecurityTimeseriesGroupArcListParamsDateRange24wControl RadarEmailSecurityTimeseriesGroupArcListParamsDateRange = "24wControl"
)

type RadarEmailSecurityTimeseriesGroupArcListParamsDKIM string

const (
	RadarEmailSecurityTimeseriesGroupArcListParamsDKIMPass RadarEmailSecurityTimeseriesGroupArcListParamsDKIM = "PASS"
	RadarEmailSecurityTimeseriesGroupArcListParamsDKIMNone RadarEmailSecurityTimeseriesGroupArcListParamsDKIM = "NONE"
	RadarEmailSecurityTimeseriesGroupArcListParamsDKIMFail RadarEmailSecurityTimeseriesGroupArcListParamsDKIM = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupArcListParamsDmarc string

const (
	RadarEmailSecurityTimeseriesGroupArcListParamsDmarcPass RadarEmailSecurityTimeseriesGroupArcListParamsDmarc = "PASS"
	RadarEmailSecurityTimeseriesGroupArcListParamsDmarcNone RadarEmailSecurityTimeseriesGroupArcListParamsDmarc = "NONE"
	RadarEmailSecurityTimeseriesGroupArcListParamsDmarcFail RadarEmailSecurityTimeseriesGroupArcListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTimeseriesGroupArcListParamsFormat string

const (
	RadarEmailSecurityTimeseriesGroupArcListParamsFormatJson RadarEmailSecurityTimeseriesGroupArcListParamsFormat = "JSON"
	RadarEmailSecurityTimeseriesGroupArcListParamsFormatCsv  RadarEmailSecurityTimeseriesGroupArcListParamsFormat = "CSV"
)

type RadarEmailSecurityTimeseriesGroupArcListParamsSPF string

const (
	RadarEmailSecurityTimeseriesGroupArcListParamsSPFPass RadarEmailSecurityTimeseriesGroupArcListParamsSPF = "PASS"
	RadarEmailSecurityTimeseriesGroupArcListParamsSPFNone RadarEmailSecurityTimeseriesGroupArcListParamsSPF = "NONE"
	RadarEmailSecurityTimeseriesGroupArcListParamsSPFFail RadarEmailSecurityTimeseriesGroupArcListParamsSPF = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupArcListResponseEnvelope struct {
	Result  RadarEmailSecurityTimeseriesGroupArcListResponse             `json:"result,required"`
	Success bool                                                         `json:"success,required"`
	JSON    radarEmailSecurityTimeseriesGroupArcListResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupArcListResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTimeseriesGroupArcListResponseEnvelope]
type radarEmailSecurityTimeseriesGroupArcListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupArcListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
