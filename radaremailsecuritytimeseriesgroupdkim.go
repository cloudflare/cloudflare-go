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

// RadarEmailSecurityTimeseriesGroupDKIMService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTimeseriesGroupDKIMService] method instead.
type RadarEmailSecurityTimeseriesGroupDKIMService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTimeseriesGroupDKIMService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarEmailSecurityTimeseriesGroupDKIMService(opts ...option.RequestOption) (r *RadarEmailSecurityTimeseriesGroupDKIMService) {
	r = &RadarEmailSecurityTimeseriesGroupDKIMService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per DKIM validation over time.
func (r *RadarEmailSecurityTimeseriesGroupDKIMService) List(ctx context.Context, query RadarEmailSecurityTimeseriesGroupDKIMListParams, opts ...option.RequestOption) (res *RadarEmailSecurityTimeseriesGroupDKIMListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTimeseriesGroupDKIMListResponseEnvelope
	path := "radar/email/security/timeseries_groups/dkim"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecurityTimeseriesGroupDKIMListResponse struct {
	Meta   interface{}                                             `json:"meta,required"`
	Serie0 RadarEmailSecurityTimeseriesGroupDKIMListResponseSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecurityTimeseriesGroupDKIMListResponseJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupDKIMListResponseJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTimeseriesGroupDKIMListResponse]
type radarEmailSecurityTimeseriesGroupDKIMListResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupDKIMListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTimeseriesGroupDKIMListResponseSerie0 struct {
	Fail []string                                                    `json:"FAIL,required"`
	None []string                                                    `json:"NONE,required"`
	Pass []string                                                    `json:"PASS,required"`
	JSON radarEmailSecurityTimeseriesGroupDKIMListResponseSerie0JSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupDKIMListResponseSerie0JSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTimeseriesGroupDKIMListResponseSerie0]
type radarEmailSecurityTimeseriesGroupDKIMListResponseSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupDKIMListResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTimeseriesGroupDKIMListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTimeseriesGroupDKIMListParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityTimeseriesGroupDKIMListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecurityTimeseriesGroupDKIMListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTimeseriesGroupDKIMListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTimeseriesGroupDKIMListParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTimeseriesGroupDKIMListParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTimeseriesGroupDKIMListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupDKIMListParamsAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupDKIMListParamsAggInterval15m RadarEmailSecurityTimeseriesGroupDKIMListParamsAggInterval = "15m"
	RadarEmailSecurityTimeseriesGroupDKIMListParamsAggInterval1h  RadarEmailSecurityTimeseriesGroupDKIMListParamsAggInterval = "1h"
	RadarEmailSecurityTimeseriesGroupDKIMListParamsAggInterval1d  RadarEmailSecurityTimeseriesGroupDKIMListParamsAggInterval = "1d"
	RadarEmailSecurityTimeseriesGroupDKIMListParamsAggInterval1w  RadarEmailSecurityTimeseriesGroupDKIMListParamsAggInterval = "1w"
)

type RadarEmailSecurityTimeseriesGroupDKIMListParamsArc string

const (
	RadarEmailSecurityTimeseriesGroupDKIMListParamsArcPass RadarEmailSecurityTimeseriesGroupDKIMListParamsArc = "PASS"
	RadarEmailSecurityTimeseriesGroupDKIMListParamsArcNone RadarEmailSecurityTimeseriesGroupDKIMListParamsArc = "NONE"
	RadarEmailSecurityTimeseriesGroupDKIMListParamsArcFail RadarEmailSecurityTimeseriesGroupDKIMListParamsArc = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange string

const (
	RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange1d         RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange = "1d"
	RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange2d         RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange = "2d"
	RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange7d         RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange = "7d"
	RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange14d        RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange = "14d"
	RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange28d        RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange = "28d"
	RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange12w        RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange = "12w"
	RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange24w        RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange = "24w"
	RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange52w        RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange = "52w"
	RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange1dControl  RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange = "1dControl"
	RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange2dControl  RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange = "2dControl"
	RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange7dControl  RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange = "7dControl"
	RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange14dControl RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange = "14dControl"
	RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange28dControl RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange = "28dControl"
	RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange12wControl RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange = "12wControl"
	RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange24wControl RadarEmailSecurityTimeseriesGroupDKIMListParamsDateRange = "24wControl"
)

type RadarEmailSecurityTimeseriesGroupDKIMListParamsDmarc string

const (
	RadarEmailSecurityTimeseriesGroupDKIMListParamsDmarcPass RadarEmailSecurityTimeseriesGroupDKIMListParamsDmarc = "PASS"
	RadarEmailSecurityTimeseriesGroupDKIMListParamsDmarcNone RadarEmailSecurityTimeseriesGroupDKIMListParamsDmarc = "NONE"
	RadarEmailSecurityTimeseriesGroupDKIMListParamsDmarcFail RadarEmailSecurityTimeseriesGroupDKIMListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTimeseriesGroupDKIMListParamsFormat string

const (
	RadarEmailSecurityTimeseriesGroupDKIMListParamsFormatJson RadarEmailSecurityTimeseriesGroupDKIMListParamsFormat = "JSON"
	RadarEmailSecurityTimeseriesGroupDKIMListParamsFormatCsv  RadarEmailSecurityTimeseriesGroupDKIMListParamsFormat = "CSV"
)

type RadarEmailSecurityTimeseriesGroupDKIMListParamsSPF string

const (
	RadarEmailSecurityTimeseriesGroupDKIMListParamsSPFPass RadarEmailSecurityTimeseriesGroupDKIMListParamsSPF = "PASS"
	RadarEmailSecurityTimeseriesGroupDKIMListParamsSPFNone RadarEmailSecurityTimeseriesGroupDKIMListParamsSPF = "NONE"
	RadarEmailSecurityTimeseriesGroupDKIMListParamsSPFFail RadarEmailSecurityTimeseriesGroupDKIMListParamsSPF = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupDKIMListResponseEnvelope struct {
	Result  RadarEmailSecurityTimeseriesGroupDKIMListResponse             `json:"result,required"`
	Success bool                                                          `json:"success,required"`
	JSON    radarEmailSecurityTimeseriesGroupDKIMListResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupDKIMListResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTimeseriesGroupDKIMListResponseEnvelope]
type radarEmailSecurityTimeseriesGroupDKIMListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupDKIMListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
