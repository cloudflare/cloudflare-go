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
func (r *RadarEmailSecurityTimeseriesGroupService) ARC(ctx context.Context, query RadarEmailSecurityTimeseriesGroupARCParams, opts ...option.RequestOption) (res *RadarEmailSecurityTimeseriesGroupARCResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTimeseriesGroupARCResponseEnvelope
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

type RadarEmailSecurityTimeseriesGroupARCResponse struct {
	Meta   interface{}                                        `json:"meta,required"`
	Serie0 RadarEmailSecurityTimeseriesGroupARCResponseSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecurityTimeseriesGroupARCResponseJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupARCResponseJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTimeseriesGroupARCResponse]
type radarEmailSecurityTimeseriesGroupARCResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupARCResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTimeseriesGroupARCResponseSerie0 struct {
	Fail []string                                               `json:"FAIL,required"`
	None []string                                               `json:"NONE,required"`
	Pass []string                                               `json:"PASS,required"`
	JSON radarEmailSecurityTimeseriesGroupARCResponseSerie0JSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupARCResponseSerie0JSON contains the JSON
// metadata for the struct [RadarEmailSecurityTimeseriesGroupARCResponseSerie0]
type radarEmailSecurityTimeseriesGroupARCResponseSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupARCResponseSerie0) UnmarshalJSON(data []byte) (err error) {
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

type RadarEmailSecurityTimeseriesGroupARCParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTimeseriesGroupARCParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTimeseriesGroupARCParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityTimeseriesGroupARCParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailSecurityTimeseriesGroupARCParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTimeseriesGroupARCParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTimeseriesGroupARCParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTimeseriesGroupARCParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTimeseriesGroupARCParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupARCParamsAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupARCParamsAggInterval15m RadarEmailSecurityTimeseriesGroupARCParamsAggInterval = "15m"
	RadarEmailSecurityTimeseriesGroupARCParamsAggInterval1h  RadarEmailSecurityTimeseriesGroupARCParamsAggInterval = "1h"
	RadarEmailSecurityTimeseriesGroupARCParamsAggInterval1d  RadarEmailSecurityTimeseriesGroupARCParamsAggInterval = "1d"
	RadarEmailSecurityTimeseriesGroupARCParamsAggInterval1w  RadarEmailSecurityTimeseriesGroupARCParamsAggInterval = "1w"
)

type RadarEmailSecurityTimeseriesGroupARCParamsDateRange string

const (
	RadarEmailSecurityTimeseriesGroupARCParamsDateRange1d         RadarEmailSecurityTimeseriesGroupARCParamsDateRange = "1d"
	RadarEmailSecurityTimeseriesGroupARCParamsDateRange2d         RadarEmailSecurityTimeseriesGroupARCParamsDateRange = "2d"
	RadarEmailSecurityTimeseriesGroupARCParamsDateRange7d         RadarEmailSecurityTimeseriesGroupARCParamsDateRange = "7d"
	RadarEmailSecurityTimeseriesGroupARCParamsDateRange14d        RadarEmailSecurityTimeseriesGroupARCParamsDateRange = "14d"
	RadarEmailSecurityTimeseriesGroupARCParamsDateRange28d        RadarEmailSecurityTimeseriesGroupARCParamsDateRange = "28d"
	RadarEmailSecurityTimeseriesGroupARCParamsDateRange12w        RadarEmailSecurityTimeseriesGroupARCParamsDateRange = "12w"
	RadarEmailSecurityTimeseriesGroupARCParamsDateRange24w        RadarEmailSecurityTimeseriesGroupARCParamsDateRange = "24w"
	RadarEmailSecurityTimeseriesGroupARCParamsDateRange52w        RadarEmailSecurityTimeseriesGroupARCParamsDateRange = "52w"
	RadarEmailSecurityTimeseriesGroupARCParamsDateRange1dControl  RadarEmailSecurityTimeseriesGroupARCParamsDateRange = "1dControl"
	RadarEmailSecurityTimeseriesGroupARCParamsDateRange2dControl  RadarEmailSecurityTimeseriesGroupARCParamsDateRange = "2dControl"
	RadarEmailSecurityTimeseriesGroupARCParamsDateRange7dControl  RadarEmailSecurityTimeseriesGroupARCParamsDateRange = "7dControl"
	RadarEmailSecurityTimeseriesGroupARCParamsDateRange14dControl RadarEmailSecurityTimeseriesGroupARCParamsDateRange = "14dControl"
	RadarEmailSecurityTimeseriesGroupARCParamsDateRange28dControl RadarEmailSecurityTimeseriesGroupARCParamsDateRange = "28dControl"
	RadarEmailSecurityTimeseriesGroupARCParamsDateRange12wControl RadarEmailSecurityTimeseriesGroupARCParamsDateRange = "12wControl"
	RadarEmailSecurityTimeseriesGroupARCParamsDateRange24wControl RadarEmailSecurityTimeseriesGroupARCParamsDateRange = "24wControl"
)

type RadarEmailSecurityTimeseriesGroupARCParamsDKIM string

const (
	RadarEmailSecurityTimeseriesGroupARCParamsDKIMPass RadarEmailSecurityTimeseriesGroupARCParamsDKIM = "PASS"
	RadarEmailSecurityTimeseriesGroupARCParamsDKIMNone RadarEmailSecurityTimeseriesGroupARCParamsDKIM = "NONE"
	RadarEmailSecurityTimeseriesGroupARCParamsDKIMFail RadarEmailSecurityTimeseriesGroupARCParamsDKIM = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupARCParamsDMARC string

const (
	RadarEmailSecurityTimeseriesGroupARCParamsDMARCPass RadarEmailSecurityTimeseriesGroupARCParamsDMARC = "PASS"
	RadarEmailSecurityTimeseriesGroupARCParamsDMARCNone RadarEmailSecurityTimeseriesGroupARCParamsDMARC = "NONE"
	RadarEmailSecurityTimeseriesGroupARCParamsDMARCFail RadarEmailSecurityTimeseriesGroupARCParamsDMARC = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTimeseriesGroupARCParamsFormat string

const (
	RadarEmailSecurityTimeseriesGroupARCParamsFormatJson RadarEmailSecurityTimeseriesGroupARCParamsFormat = "JSON"
	RadarEmailSecurityTimeseriesGroupARCParamsFormatCsv  RadarEmailSecurityTimeseriesGroupARCParamsFormat = "CSV"
)

type RadarEmailSecurityTimeseriesGroupARCParamsSPF string

const (
	RadarEmailSecurityTimeseriesGroupARCParamsSPFPass RadarEmailSecurityTimeseriesGroupARCParamsSPF = "PASS"
	RadarEmailSecurityTimeseriesGroupARCParamsSPFNone RadarEmailSecurityTimeseriesGroupARCParamsSPF = "NONE"
	RadarEmailSecurityTimeseriesGroupARCParamsSPFFail RadarEmailSecurityTimeseriesGroupARCParamsSPF = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupARCResponseEnvelope struct {
	Result  RadarEmailSecurityTimeseriesGroupARCResponse             `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    radarEmailSecurityTimeseriesGroupARCResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupARCResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTimeseriesGroupARCResponseEnvelope]
type radarEmailSecurityTimeseriesGroupARCResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupARCResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTimeseriesGroupDKIMParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTimeseriesGroupDKIMParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecurityTimeseriesGroupDKIMParamsARC] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailSecurityTimeseriesGroupDKIMParamsDMARC] `query:"dmarc"`
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

type RadarEmailSecurityTimeseriesGroupDKIMParamsARC string

const (
	RadarEmailSecurityTimeseriesGroupDKIMParamsARCPass RadarEmailSecurityTimeseriesGroupDKIMParamsARC = "PASS"
	RadarEmailSecurityTimeseriesGroupDKIMParamsARCNone RadarEmailSecurityTimeseriesGroupDKIMParamsARC = "NONE"
	RadarEmailSecurityTimeseriesGroupDKIMParamsARCFail RadarEmailSecurityTimeseriesGroupDKIMParamsARC = "FAIL"
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

type RadarEmailSecurityTimeseriesGroupDKIMParamsDMARC string

const (
	RadarEmailSecurityTimeseriesGroupDKIMParamsDMARCPass RadarEmailSecurityTimeseriesGroupDKIMParamsDMARC = "PASS"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDMARCNone RadarEmailSecurityTimeseriesGroupDKIMParamsDMARC = "NONE"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDMARCFail RadarEmailSecurityTimeseriesGroupDKIMParamsDMARC = "FAIL"
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
