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

// RadarEmailSecurityTimeseriesGroupService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTimeseriesGroupService] method instead.
type RadarEmailSecurityTimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTimeseriesGroupService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityTimeseriesGroupService(opts ...option.RequestOption) (r *RadarEmailSecurityTimeseriesGroupService) {
	r = &RadarEmailSecurityTimeseriesGroupService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per Arc validation over time.
func (r *RadarEmailSecurityTimeseriesGroupService) Arc(ctx context.Context, query RadarEmailSecurityTimeseriesGroupArcParams, opts ...option.RequestOption) (res *RadarEmailSecurityTimeseriesGroupArcResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTimeseriesGroupArcResponseEnvelope
	path := "radar/email/security/timeseries_groups/arc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified per DKIM validation over time.
func (r *RadarEmailSecurityTimeseriesGroupService) DKIM(ctx context.Context, query RadarEmailSecurityTimeseriesGroupDKIMParams, opts ...option.RequestOption) (res *RadarEmailSecurityTimeseriesGroupDKIMResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTimeseriesGroupDKIMResponseEnvelope
	path := "radar/email/security/timeseries_groups/dkim"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecurityTimeseriesGroupArcResponse struct {
	Meta   interface{}                                        `json:"meta,required"`
	Serie0 RadarEmailSecurityTimeseriesGroupArcResponseSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecurityTimeseriesGroupArcResponseJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupArcResponseJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTimeseriesGroupArcResponse]
type radarEmailSecurityTimeseriesGroupArcResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupArcResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTimeseriesGroupArcResponseSerie0 struct {
	Fail []string                                               `json:"FAIL,required"`
	None []string                                               `json:"NONE,required"`
	Pass []string                                               `json:"PASS,required"`
	JSON radarEmailSecurityTimeseriesGroupArcResponseSerie0JSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupArcResponseSerie0JSON contains the JSON
// metadata for the struct [RadarEmailSecurityTimeseriesGroupArcResponseSerie0]
type radarEmailSecurityTimeseriesGroupArcResponseSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupArcResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTimeseriesGroupDKIMResponse struct {
	Meta   interface{}                                         `json:"meta,required"`
	Serie0 RadarEmailSecurityTimeseriesGroupDKIMResponseSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecurityTimeseriesGroupDKIMResponseJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupDKIMResponseJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTimeseriesGroupDKIMResponse]
type radarEmailSecurityTimeseriesGroupDKIMResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupDKIMResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTimeseriesGroupDKIMResponseSerie0 struct {
	Fail []string                                                `json:"FAIL,required"`
	None []string                                                `json:"NONE,required"`
	Pass []string                                                `json:"PASS,required"`
	JSON radarEmailSecurityTimeseriesGroupDKIMResponseSerie0JSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupDKIMResponseSerie0JSON contains the JSON
// metadata for the struct [RadarEmailSecurityTimeseriesGroupDKIMResponseSerie0]
type radarEmailSecurityTimeseriesGroupDKIMResponseSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupDKIMResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTimeseriesGroupArcParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTimeseriesGroupArcParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTimeseriesGroupArcParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityTimeseriesGroupArcParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecurityTimeseriesGroupArcParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTimeseriesGroupArcParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTimeseriesGroupArcParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTimeseriesGroupArcParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTimeseriesGroupArcParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupArcParamsAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupArcParamsAggInterval15m RadarEmailSecurityTimeseriesGroupArcParamsAggInterval = "15m"
	RadarEmailSecurityTimeseriesGroupArcParamsAggInterval1h  RadarEmailSecurityTimeseriesGroupArcParamsAggInterval = "1h"
	RadarEmailSecurityTimeseriesGroupArcParamsAggInterval1d  RadarEmailSecurityTimeseriesGroupArcParamsAggInterval = "1d"
	RadarEmailSecurityTimeseriesGroupArcParamsAggInterval1w  RadarEmailSecurityTimeseriesGroupArcParamsAggInterval = "1w"
)

type RadarEmailSecurityTimeseriesGroupArcParamsDateRange string

const (
	RadarEmailSecurityTimeseriesGroupArcParamsDateRange1d         RadarEmailSecurityTimeseriesGroupArcParamsDateRange = "1d"
	RadarEmailSecurityTimeseriesGroupArcParamsDateRange2d         RadarEmailSecurityTimeseriesGroupArcParamsDateRange = "2d"
	RadarEmailSecurityTimeseriesGroupArcParamsDateRange7d         RadarEmailSecurityTimeseriesGroupArcParamsDateRange = "7d"
	RadarEmailSecurityTimeseriesGroupArcParamsDateRange14d        RadarEmailSecurityTimeseriesGroupArcParamsDateRange = "14d"
	RadarEmailSecurityTimeseriesGroupArcParamsDateRange28d        RadarEmailSecurityTimeseriesGroupArcParamsDateRange = "28d"
	RadarEmailSecurityTimeseriesGroupArcParamsDateRange12w        RadarEmailSecurityTimeseriesGroupArcParamsDateRange = "12w"
	RadarEmailSecurityTimeseriesGroupArcParamsDateRange24w        RadarEmailSecurityTimeseriesGroupArcParamsDateRange = "24w"
	RadarEmailSecurityTimeseriesGroupArcParamsDateRange52w        RadarEmailSecurityTimeseriesGroupArcParamsDateRange = "52w"
	RadarEmailSecurityTimeseriesGroupArcParamsDateRange1dControl  RadarEmailSecurityTimeseriesGroupArcParamsDateRange = "1dControl"
	RadarEmailSecurityTimeseriesGroupArcParamsDateRange2dControl  RadarEmailSecurityTimeseriesGroupArcParamsDateRange = "2dControl"
	RadarEmailSecurityTimeseriesGroupArcParamsDateRange7dControl  RadarEmailSecurityTimeseriesGroupArcParamsDateRange = "7dControl"
	RadarEmailSecurityTimeseriesGroupArcParamsDateRange14dControl RadarEmailSecurityTimeseriesGroupArcParamsDateRange = "14dControl"
	RadarEmailSecurityTimeseriesGroupArcParamsDateRange28dControl RadarEmailSecurityTimeseriesGroupArcParamsDateRange = "28dControl"
	RadarEmailSecurityTimeseriesGroupArcParamsDateRange12wControl RadarEmailSecurityTimeseriesGroupArcParamsDateRange = "12wControl"
	RadarEmailSecurityTimeseriesGroupArcParamsDateRange24wControl RadarEmailSecurityTimeseriesGroupArcParamsDateRange = "24wControl"
)

type RadarEmailSecurityTimeseriesGroupArcParamsDKIM string

const (
	RadarEmailSecurityTimeseriesGroupArcParamsDKIMPass RadarEmailSecurityTimeseriesGroupArcParamsDKIM = "PASS"
	RadarEmailSecurityTimeseriesGroupArcParamsDKIMNone RadarEmailSecurityTimeseriesGroupArcParamsDKIM = "NONE"
	RadarEmailSecurityTimeseriesGroupArcParamsDKIMFail RadarEmailSecurityTimeseriesGroupArcParamsDKIM = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupArcParamsDmarc string

const (
	RadarEmailSecurityTimeseriesGroupArcParamsDmarcPass RadarEmailSecurityTimeseriesGroupArcParamsDmarc = "PASS"
	RadarEmailSecurityTimeseriesGroupArcParamsDmarcNone RadarEmailSecurityTimeseriesGroupArcParamsDmarc = "NONE"
	RadarEmailSecurityTimeseriesGroupArcParamsDmarcFail RadarEmailSecurityTimeseriesGroupArcParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTimeseriesGroupArcParamsFormat string

const (
	RadarEmailSecurityTimeseriesGroupArcParamsFormatJson RadarEmailSecurityTimeseriesGroupArcParamsFormat = "JSON"
	RadarEmailSecurityTimeseriesGroupArcParamsFormatCsv  RadarEmailSecurityTimeseriesGroupArcParamsFormat = "CSV"
)

type RadarEmailSecurityTimeseriesGroupArcParamsSPF string

const (
	RadarEmailSecurityTimeseriesGroupArcParamsSPFPass RadarEmailSecurityTimeseriesGroupArcParamsSPF = "PASS"
	RadarEmailSecurityTimeseriesGroupArcParamsSPFNone RadarEmailSecurityTimeseriesGroupArcParamsSPF = "NONE"
	RadarEmailSecurityTimeseriesGroupArcParamsSPFFail RadarEmailSecurityTimeseriesGroupArcParamsSPF = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupArcResponseEnvelope struct {
	Result  RadarEmailSecurityTimeseriesGroupArcResponse             `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    radarEmailSecurityTimeseriesGroupArcResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupArcResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTimeseriesGroupArcResponseEnvelope]
type radarEmailSecurityTimeseriesGroupArcResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupArcResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTimeseriesGroupDKIMParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTimeseriesGroupDKIMParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityTimeseriesGroupDKIMParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecurityTimeseriesGroupDKIMParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTimeseriesGroupDKIMParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTimeseriesGroupDKIMParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTimeseriesGroupDKIMParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTimeseriesGroupDKIMParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupDKIMParamsAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupDKIMParamsAggInterval15m RadarEmailSecurityTimeseriesGroupDKIMParamsAggInterval = "15m"
	RadarEmailSecurityTimeseriesGroupDKIMParamsAggInterval1h  RadarEmailSecurityTimeseriesGroupDKIMParamsAggInterval = "1h"
	RadarEmailSecurityTimeseriesGroupDKIMParamsAggInterval1d  RadarEmailSecurityTimeseriesGroupDKIMParamsAggInterval = "1d"
	RadarEmailSecurityTimeseriesGroupDKIMParamsAggInterval1w  RadarEmailSecurityTimeseriesGroupDKIMParamsAggInterval = "1w"
)

type RadarEmailSecurityTimeseriesGroupDKIMParamsArc string

const (
	RadarEmailSecurityTimeseriesGroupDKIMParamsArcPass RadarEmailSecurityTimeseriesGroupDKIMParamsArc = "PASS"
	RadarEmailSecurityTimeseriesGroupDKIMParamsArcNone RadarEmailSecurityTimeseriesGroupDKIMParamsArc = "NONE"
	RadarEmailSecurityTimeseriesGroupDKIMParamsArcFail RadarEmailSecurityTimeseriesGroupDKIMParamsArc = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange string

const (
	RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange1d         RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange = "1d"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange2d         RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange = "2d"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange7d         RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange = "7d"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange14d        RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange = "14d"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange28d        RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange = "28d"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange12w        RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange = "12w"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange24w        RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange = "24w"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange52w        RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange = "52w"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange1dControl  RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange = "1dControl"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange2dControl  RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange = "2dControl"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange7dControl  RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange = "7dControl"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange14dControl RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange = "14dControl"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange28dControl RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange = "28dControl"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange12wControl RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange = "12wControl"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange24wControl RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange = "24wControl"
)

type RadarEmailSecurityTimeseriesGroupDKIMParamsDmarc string

const (
	RadarEmailSecurityTimeseriesGroupDKIMParamsDmarcPass RadarEmailSecurityTimeseriesGroupDKIMParamsDmarc = "PASS"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDmarcNone RadarEmailSecurityTimeseriesGroupDKIMParamsDmarc = "NONE"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDmarcFail RadarEmailSecurityTimeseriesGroupDKIMParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTimeseriesGroupDKIMParamsFormat string

const (
	RadarEmailSecurityTimeseriesGroupDKIMParamsFormatJson RadarEmailSecurityTimeseriesGroupDKIMParamsFormat = "JSON"
	RadarEmailSecurityTimeseriesGroupDKIMParamsFormatCsv  RadarEmailSecurityTimeseriesGroupDKIMParamsFormat = "CSV"
)

type RadarEmailSecurityTimeseriesGroupDKIMParamsSPF string

const (
	RadarEmailSecurityTimeseriesGroupDKIMParamsSPFPass RadarEmailSecurityTimeseriesGroupDKIMParamsSPF = "PASS"
	RadarEmailSecurityTimeseriesGroupDKIMParamsSPFNone RadarEmailSecurityTimeseriesGroupDKIMParamsSPF = "NONE"
	RadarEmailSecurityTimeseriesGroupDKIMParamsSPFFail RadarEmailSecurityTimeseriesGroupDKIMParamsSPF = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupDKIMResponseEnvelope struct {
	Result  RadarEmailSecurityTimeseriesGroupDKIMResponse             `json:"result,required"`
	Success bool                                                      `json:"success,required"`
	JSON    radarEmailSecurityTimeseriesGroupDKIMResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupDKIMResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTimeseriesGroupDKIMResponseEnvelope]
type radarEmailSecurityTimeseriesGroupDKIMResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupDKIMResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
