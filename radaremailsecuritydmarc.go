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

// RadarEmailSecurityDmarcService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarEmailSecurityDmarcService] method instead.
type RadarEmailSecurityDmarcService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityDmarcService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityDmarcService(opts ...option.RequestOption) (r *RadarEmailSecurityDmarcService) {
	r = &RadarEmailSecurityDmarcService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per DMARC validation over time.
func (r *RadarEmailSecurityDmarcService) List(ctx context.Context, query RadarEmailSecurityDmarcListParams, opts ...option.RequestOption) (res *RadarEmailSecurityDmarcListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityDmarcListResponseEnvelope
	path := "radar/email/security/timeseries_groups/dmarc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecurityDmarcListResponse struct {
	Meta   interface{}                               `json:"meta,required"`
	Serie0 RadarEmailSecurityDmarcListResponseSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecurityDmarcListResponseJSON   `json:"-"`
}

// radarEmailSecurityDmarcListResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecurityDmarcListResponse]
type radarEmailSecurityDmarcListResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityDmarcListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityDmarcListResponseSerie0 struct {
	Fail []string                                      `json:"FAIL,required"`
	None []string                                      `json:"NONE,required"`
	Pass []string                                      `json:"PASS,required"`
	JSON radarEmailSecurityDmarcListResponseSerie0JSON `json:"-"`
}

// radarEmailSecurityDmarcListResponseSerie0JSON contains the JSON metadata for the
// struct [RadarEmailSecurityDmarcListResponseSerie0]
type radarEmailSecurityDmarcListResponseSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityDmarcListResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityDmarcListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityDmarcListParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityDmarcListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityDmarcListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSecurityDmarcListParamsDkim] `query:"dkim"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityDmarcListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailSecurityDmarcListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityDmarcListParams]'s query parameters as
// `url.Values`.
func (r RadarEmailSecurityDmarcListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityDmarcListParamsAggInterval string

const (
	RadarEmailSecurityDmarcListParamsAggInterval15m RadarEmailSecurityDmarcListParamsAggInterval = "15m"
	RadarEmailSecurityDmarcListParamsAggInterval1h  RadarEmailSecurityDmarcListParamsAggInterval = "1h"
	RadarEmailSecurityDmarcListParamsAggInterval1d  RadarEmailSecurityDmarcListParamsAggInterval = "1d"
	RadarEmailSecurityDmarcListParamsAggInterval1w  RadarEmailSecurityDmarcListParamsAggInterval = "1w"
)

type RadarEmailSecurityDmarcListParamsArc string

const (
	RadarEmailSecurityDmarcListParamsArcPass RadarEmailSecurityDmarcListParamsArc = "PASS"
	RadarEmailSecurityDmarcListParamsArcNone RadarEmailSecurityDmarcListParamsArc = "NONE"
	RadarEmailSecurityDmarcListParamsArcFail RadarEmailSecurityDmarcListParamsArc = "FAIL"
)

type RadarEmailSecurityDmarcListParamsDateRange string

const (
	RadarEmailSecurityDmarcListParamsDateRange1d         RadarEmailSecurityDmarcListParamsDateRange = "1d"
	RadarEmailSecurityDmarcListParamsDateRange2d         RadarEmailSecurityDmarcListParamsDateRange = "2d"
	RadarEmailSecurityDmarcListParamsDateRange7d         RadarEmailSecurityDmarcListParamsDateRange = "7d"
	RadarEmailSecurityDmarcListParamsDateRange14d        RadarEmailSecurityDmarcListParamsDateRange = "14d"
	RadarEmailSecurityDmarcListParamsDateRange28d        RadarEmailSecurityDmarcListParamsDateRange = "28d"
	RadarEmailSecurityDmarcListParamsDateRange12w        RadarEmailSecurityDmarcListParamsDateRange = "12w"
	RadarEmailSecurityDmarcListParamsDateRange24w        RadarEmailSecurityDmarcListParamsDateRange = "24w"
	RadarEmailSecurityDmarcListParamsDateRange52w        RadarEmailSecurityDmarcListParamsDateRange = "52w"
	RadarEmailSecurityDmarcListParamsDateRange1dControl  RadarEmailSecurityDmarcListParamsDateRange = "1dControl"
	RadarEmailSecurityDmarcListParamsDateRange2dControl  RadarEmailSecurityDmarcListParamsDateRange = "2dControl"
	RadarEmailSecurityDmarcListParamsDateRange7dControl  RadarEmailSecurityDmarcListParamsDateRange = "7dControl"
	RadarEmailSecurityDmarcListParamsDateRange14dControl RadarEmailSecurityDmarcListParamsDateRange = "14dControl"
	RadarEmailSecurityDmarcListParamsDateRange28dControl RadarEmailSecurityDmarcListParamsDateRange = "28dControl"
	RadarEmailSecurityDmarcListParamsDateRange12wControl RadarEmailSecurityDmarcListParamsDateRange = "12wControl"
	RadarEmailSecurityDmarcListParamsDateRange24wControl RadarEmailSecurityDmarcListParamsDateRange = "24wControl"
)

type RadarEmailSecurityDmarcListParamsDkim string

const (
	RadarEmailSecurityDmarcListParamsDkimPass RadarEmailSecurityDmarcListParamsDkim = "PASS"
	RadarEmailSecurityDmarcListParamsDkimNone RadarEmailSecurityDmarcListParamsDkim = "NONE"
	RadarEmailSecurityDmarcListParamsDkimFail RadarEmailSecurityDmarcListParamsDkim = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityDmarcListParamsFormat string

const (
	RadarEmailSecurityDmarcListParamsFormatJson RadarEmailSecurityDmarcListParamsFormat = "JSON"
	RadarEmailSecurityDmarcListParamsFormatCsv  RadarEmailSecurityDmarcListParamsFormat = "CSV"
)

type RadarEmailSecurityDmarcListParamsSpf string

const (
	RadarEmailSecurityDmarcListParamsSpfPass RadarEmailSecurityDmarcListParamsSpf = "PASS"
	RadarEmailSecurityDmarcListParamsSpfNone RadarEmailSecurityDmarcListParamsSpf = "NONE"
	RadarEmailSecurityDmarcListParamsSpfFail RadarEmailSecurityDmarcListParamsSpf = "FAIL"
)

type RadarEmailSecurityDmarcListResponseEnvelope struct {
	Result  RadarEmailSecurityDmarcListResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarEmailSecurityDmarcListResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityDmarcListResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarEmailSecurityDmarcListResponseEnvelope]
type radarEmailSecurityDmarcListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityDmarcListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
