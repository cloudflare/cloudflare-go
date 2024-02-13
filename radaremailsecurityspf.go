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

// RadarEmailSecuritySPFService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEmailSecuritySPFService]
// method instead.
type RadarEmailSecuritySPFService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecuritySPFService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarEmailSecuritySPFService(opts ...option.RequestOption) (r *RadarEmailSecuritySPFService) {
	r = &RadarEmailSecuritySPFService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per SPF validation over time.
func (r *RadarEmailSecuritySPFService) List(ctx context.Context, query RadarEmailSecuritySPFListParams, opts ...option.RequestOption) (res *RadarEmailSecuritySPFListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecuritySPFListResponseEnvelope
	path := "radar/email/security/timeseries_groups/spf"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecuritySPFListResponse struct {
	Meta   interface{}                             `json:"meta,required"`
	Serie0 RadarEmailSecuritySPFListResponseSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecuritySPFListResponseJSON   `json:"-"`
}

// radarEmailSecuritySPFListResponseJSON contains the JSON metadata for the struct
// [RadarEmailSecuritySPFListResponse]
type radarEmailSecuritySPFListResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySPFListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySPFListResponseSerie0 struct {
	Fail []string                                    `json:"FAIL,required"`
	None []string                                    `json:"NONE,required"`
	Pass []string                                    `json:"PASS,required"`
	JSON radarEmailSecuritySPFListResponseSerie0JSON `json:"-"`
}

// radarEmailSecuritySPFListResponseSerie0JSON contains the JSON metadata for the
// struct [RadarEmailSecuritySPFListResponseSerie0]
type radarEmailSecuritySPFListResponseSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySPFListResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySPFListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecuritySPFListParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecuritySPFListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecuritySPFListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecuritySPFListParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecuritySPFListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecuritySPFListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarEmailSecuritySPFListParams]'s query parameters as
// `url.Values`.
func (r RadarEmailSecuritySPFListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecuritySPFListParamsAggInterval string

const (
	RadarEmailSecuritySPFListParamsAggInterval15m RadarEmailSecuritySPFListParamsAggInterval = "15m"
	RadarEmailSecuritySPFListParamsAggInterval1h  RadarEmailSecuritySPFListParamsAggInterval = "1h"
	RadarEmailSecuritySPFListParamsAggInterval1d  RadarEmailSecuritySPFListParamsAggInterval = "1d"
	RadarEmailSecuritySPFListParamsAggInterval1w  RadarEmailSecuritySPFListParamsAggInterval = "1w"
)

type RadarEmailSecuritySPFListParamsArc string

const (
	RadarEmailSecuritySPFListParamsArcPass RadarEmailSecuritySPFListParamsArc = "PASS"
	RadarEmailSecuritySPFListParamsArcNone RadarEmailSecuritySPFListParamsArc = "NONE"
	RadarEmailSecuritySPFListParamsArcFail RadarEmailSecuritySPFListParamsArc = "FAIL"
)

type RadarEmailSecuritySPFListParamsDateRange string

const (
	RadarEmailSecuritySPFListParamsDateRange1d         RadarEmailSecuritySPFListParamsDateRange = "1d"
	RadarEmailSecuritySPFListParamsDateRange2d         RadarEmailSecuritySPFListParamsDateRange = "2d"
	RadarEmailSecuritySPFListParamsDateRange7d         RadarEmailSecuritySPFListParamsDateRange = "7d"
	RadarEmailSecuritySPFListParamsDateRange14d        RadarEmailSecuritySPFListParamsDateRange = "14d"
	RadarEmailSecuritySPFListParamsDateRange28d        RadarEmailSecuritySPFListParamsDateRange = "28d"
	RadarEmailSecuritySPFListParamsDateRange12w        RadarEmailSecuritySPFListParamsDateRange = "12w"
	RadarEmailSecuritySPFListParamsDateRange24w        RadarEmailSecuritySPFListParamsDateRange = "24w"
	RadarEmailSecuritySPFListParamsDateRange52w        RadarEmailSecuritySPFListParamsDateRange = "52w"
	RadarEmailSecuritySPFListParamsDateRange1dControl  RadarEmailSecuritySPFListParamsDateRange = "1dControl"
	RadarEmailSecuritySPFListParamsDateRange2dControl  RadarEmailSecuritySPFListParamsDateRange = "2dControl"
	RadarEmailSecuritySPFListParamsDateRange7dControl  RadarEmailSecuritySPFListParamsDateRange = "7dControl"
	RadarEmailSecuritySPFListParamsDateRange14dControl RadarEmailSecuritySPFListParamsDateRange = "14dControl"
	RadarEmailSecuritySPFListParamsDateRange28dControl RadarEmailSecuritySPFListParamsDateRange = "28dControl"
	RadarEmailSecuritySPFListParamsDateRange12wControl RadarEmailSecuritySPFListParamsDateRange = "12wControl"
	RadarEmailSecuritySPFListParamsDateRange24wControl RadarEmailSecuritySPFListParamsDateRange = "24wControl"
)

type RadarEmailSecuritySPFListParamsDKIM string

const (
	RadarEmailSecuritySPFListParamsDKIMPass RadarEmailSecuritySPFListParamsDKIM = "PASS"
	RadarEmailSecuritySPFListParamsDKIMNone RadarEmailSecuritySPFListParamsDKIM = "NONE"
	RadarEmailSecuritySPFListParamsDKIMFail RadarEmailSecuritySPFListParamsDKIM = "FAIL"
)

type RadarEmailSecuritySPFListParamsDmarc string

const (
	RadarEmailSecuritySPFListParamsDmarcPass RadarEmailSecuritySPFListParamsDmarc = "PASS"
	RadarEmailSecuritySPFListParamsDmarcNone RadarEmailSecuritySPFListParamsDmarc = "NONE"
	RadarEmailSecuritySPFListParamsDmarcFail RadarEmailSecuritySPFListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecuritySPFListParamsFormat string

const (
	RadarEmailSecuritySPFListParamsFormatJson RadarEmailSecuritySPFListParamsFormat = "JSON"
	RadarEmailSecuritySPFListParamsFormatCsv  RadarEmailSecuritySPFListParamsFormat = "CSV"
)

type RadarEmailSecuritySPFListResponseEnvelope struct {
	Result  RadarEmailSecuritySPFListResponse             `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    radarEmailSecuritySPFListResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecuritySPFListResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarEmailSecuritySPFListResponseEnvelope]
type radarEmailSecuritySPFListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySPFListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
