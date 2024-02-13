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

// RadarEmailSecurityMaliciousService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarEmailSecurityMaliciousService] method instead.
type RadarEmailSecurityMaliciousService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityMaliciousService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityMaliciousService(opts ...option.RequestOption) (r *RadarEmailSecurityMaliciousService) {
	r = &RadarEmailSecurityMaliciousService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified as MALICIOUS over time.
func (r *RadarEmailSecurityMaliciousService) List(ctx context.Context, query RadarEmailSecurityMaliciousListParams, opts ...option.RequestOption) (res *RadarEmailSecurityMaliciousListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityMaliciousListResponseEnvelope
	path := "radar/email/security/timeseries_groups/malicious"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecurityMaliciousListResponse struct {
	Meta   interface{}                                   `json:"meta,required"`
	Serie0 RadarEmailSecurityMaliciousListResponseSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecurityMaliciousListResponseJSON   `json:"-"`
}

// radarEmailSecurityMaliciousListResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecurityMaliciousListResponse]
type radarEmailSecurityMaliciousListResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityMaliciousListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityMaliciousListResponseSerie0 struct {
	Malicious    []string                                          `json:"MALICIOUS,required"`
	NotMalicious []string                                          `json:"NOT_MALICIOUS,required"`
	JSON         radarEmailSecurityMaliciousListResponseSerie0JSON `json:"-"`
}

// radarEmailSecurityMaliciousListResponseSerie0JSON contains the JSON metadata for
// the struct [RadarEmailSecurityMaliciousListResponseSerie0]
type radarEmailSecurityMaliciousListResponseSerie0JSON struct {
	Malicious    apijson.Field
	NotMalicious apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailSecurityMaliciousListResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityMaliciousListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityMaliciousListParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityMaliciousListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityMaliciousListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityMaliciousListParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecurityMaliciousListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityMaliciousListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityMaliciousListParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityMaliciousListParams]'s query parameters
// as `url.Values`.
func (r RadarEmailSecurityMaliciousListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityMaliciousListParamsAggInterval string

const (
	RadarEmailSecurityMaliciousListParamsAggInterval15m RadarEmailSecurityMaliciousListParamsAggInterval = "15m"
	RadarEmailSecurityMaliciousListParamsAggInterval1h  RadarEmailSecurityMaliciousListParamsAggInterval = "1h"
	RadarEmailSecurityMaliciousListParamsAggInterval1d  RadarEmailSecurityMaliciousListParamsAggInterval = "1d"
	RadarEmailSecurityMaliciousListParamsAggInterval1w  RadarEmailSecurityMaliciousListParamsAggInterval = "1w"
)

type RadarEmailSecurityMaliciousListParamsArc string

const (
	RadarEmailSecurityMaliciousListParamsArcPass RadarEmailSecurityMaliciousListParamsArc = "PASS"
	RadarEmailSecurityMaliciousListParamsArcNone RadarEmailSecurityMaliciousListParamsArc = "NONE"
	RadarEmailSecurityMaliciousListParamsArcFail RadarEmailSecurityMaliciousListParamsArc = "FAIL"
)

type RadarEmailSecurityMaliciousListParamsDateRange string

const (
	RadarEmailSecurityMaliciousListParamsDateRange1d         RadarEmailSecurityMaliciousListParamsDateRange = "1d"
	RadarEmailSecurityMaliciousListParamsDateRange2d         RadarEmailSecurityMaliciousListParamsDateRange = "2d"
	RadarEmailSecurityMaliciousListParamsDateRange7d         RadarEmailSecurityMaliciousListParamsDateRange = "7d"
	RadarEmailSecurityMaliciousListParamsDateRange14d        RadarEmailSecurityMaliciousListParamsDateRange = "14d"
	RadarEmailSecurityMaliciousListParamsDateRange28d        RadarEmailSecurityMaliciousListParamsDateRange = "28d"
	RadarEmailSecurityMaliciousListParamsDateRange12w        RadarEmailSecurityMaliciousListParamsDateRange = "12w"
	RadarEmailSecurityMaliciousListParamsDateRange24w        RadarEmailSecurityMaliciousListParamsDateRange = "24w"
	RadarEmailSecurityMaliciousListParamsDateRange52w        RadarEmailSecurityMaliciousListParamsDateRange = "52w"
	RadarEmailSecurityMaliciousListParamsDateRange1dControl  RadarEmailSecurityMaliciousListParamsDateRange = "1dControl"
	RadarEmailSecurityMaliciousListParamsDateRange2dControl  RadarEmailSecurityMaliciousListParamsDateRange = "2dControl"
	RadarEmailSecurityMaliciousListParamsDateRange7dControl  RadarEmailSecurityMaliciousListParamsDateRange = "7dControl"
	RadarEmailSecurityMaliciousListParamsDateRange14dControl RadarEmailSecurityMaliciousListParamsDateRange = "14dControl"
	RadarEmailSecurityMaliciousListParamsDateRange28dControl RadarEmailSecurityMaliciousListParamsDateRange = "28dControl"
	RadarEmailSecurityMaliciousListParamsDateRange12wControl RadarEmailSecurityMaliciousListParamsDateRange = "12wControl"
	RadarEmailSecurityMaliciousListParamsDateRange24wControl RadarEmailSecurityMaliciousListParamsDateRange = "24wControl"
)

type RadarEmailSecurityMaliciousListParamsDKIM string

const (
	RadarEmailSecurityMaliciousListParamsDKIMPass RadarEmailSecurityMaliciousListParamsDKIM = "PASS"
	RadarEmailSecurityMaliciousListParamsDKIMNone RadarEmailSecurityMaliciousListParamsDKIM = "NONE"
	RadarEmailSecurityMaliciousListParamsDKIMFail RadarEmailSecurityMaliciousListParamsDKIM = "FAIL"
)

type RadarEmailSecurityMaliciousListParamsDmarc string

const (
	RadarEmailSecurityMaliciousListParamsDmarcPass RadarEmailSecurityMaliciousListParamsDmarc = "PASS"
	RadarEmailSecurityMaliciousListParamsDmarcNone RadarEmailSecurityMaliciousListParamsDmarc = "NONE"
	RadarEmailSecurityMaliciousListParamsDmarcFail RadarEmailSecurityMaliciousListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityMaliciousListParamsFormat string

const (
	RadarEmailSecurityMaliciousListParamsFormatJson RadarEmailSecurityMaliciousListParamsFormat = "JSON"
	RadarEmailSecurityMaliciousListParamsFormatCsv  RadarEmailSecurityMaliciousListParamsFormat = "CSV"
)

type RadarEmailSecurityMaliciousListParamsSPF string

const (
	RadarEmailSecurityMaliciousListParamsSPFPass RadarEmailSecurityMaliciousListParamsSPF = "PASS"
	RadarEmailSecurityMaliciousListParamsSPFNone RadarEmailSecurityMaliciousListParamsSPF = "NONE"
	RadarEmailSecurityMaliciousListParamsSPFFail RadarEmailSecurityMaliciousListParamsSPF = "FAIL"
)

type RadarEmailSecurityMaliciousListResponseEnvelope struct {
	Result  RadarEmailSecurityMaliciousListResponse             `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    radarEmailSecurityMaliciousListResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityMaliciousListResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarEmailSecurityMaliciousListResponseEnvelope]
type radarEmailSecurityMaliciousListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityMaliciousListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
