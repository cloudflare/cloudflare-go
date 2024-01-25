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

// RadarEmailSecurityMaliciousTimeseryService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityMaliciousTimeseryService] method instead.
type RadarEmailSecurityMaliciousTimeseryService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityMaliciousTimeseryService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarEmailSecurityMaliciousTimeseryService(opts ...option.RequestOption) (r *RadarEmailSecurityMaliciousTimeseryService) {
	r = &RadarEmailSecurityMaliciousTimeseryService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified as MALICIOUS over time.
func (r *RadarEmailSecurityMaliciousTimeseryService) List(ctx context.Context, query RadarEmailSecurityMaliciousTimeseryListParams, opts ...option.RequestOption) (res *RadarEmailSecurityMaliciousTimeseryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/timeseries_groups/malicious"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSecurityMaliciousTimeseryListResponse struct {
	Result  RadarEmailSecurityMaliciousTimeseryListResponseResult `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    radarEmailSecurityMaliciousTimeseryListResponseJSON   `json:"-"`
}

// radarEmailSecurityMaliciousTimeseryListResponseJSON contains the JSON metadata
// for the struct [RadarEmailSecurityMaliciousTimeseryListResponse]
type radarEmailSecurityMaliciousTimeseryListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityMaliciousTimeseryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityMaliciousTimeseryListResponseResult struct {
	Meta   interface{}                                                 `json:"meta,required"`
	Serie0 RadarEmailSecurityMaliciousTimeseryListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecurityMaliciousTimeseryListResponseResultJSON   `json:"-"`
}

// radarEmailSecurityMaliciousTimeseryListResponseResultJSON contains the JSON
// metadata for the struct [RadarEmailSecurityMaliciousTimeseryListResponseResult]
type radarEmailSecurityMaliciousTimeseryListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityMaliciousTimeseryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityMaliciousTimeseryListResponseResultSerie0 struct {
	Malicious    []string                                                        `json:"MALICIOUS,required"`
	NotMalicious []string                                                        `json:"NOT_MALICIOUS,required"`
	JSON         radarEmailSecurityMaliciousTimeseryListResponseResultSerie0JSON `json:"-"`
}

// radarEmailSecurityMaliciousTimeseryListResponseResultSerie0JSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityMaliciousTimeseryListResponseResultSerie0]
type radarEmailSecurityMaliciousTimeseryListResponseResultSerie0JSON struct {
	Malicious    apijson.Field
	NotMalicious apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailSecurityMaliciousTimeseryListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityMaliciousTimeseryListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityMaliciousTimeseryListParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityMaliciousTimeseryListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityMaliciousTimeseryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSecurityMaliciousTimeseryListParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecurityMaliciousTimeseryListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityMaliciousTimeseryListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailSecurityMaliciousTimeseryListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityMaliciousTimeseryListParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityMaliciousTimeseryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityMaliciousTimeseryListParamsAggInterval string

const (
	RadarEmailSecurityMaliciousTimeseryListParamsAggInterval15m RadarEmailSecurityMaliciousTimeseryListParamsAggInterval = "15m"
	RadarEmailSecurityMaliciousTimeseryListParamsAggInterval1h  RadarEmailSecurityMaliciousTimeseryListParamsAggInterval = "1h"
	RadarEmailSecurityMaliciousTimeseryListParamsAggInterval1d  RadarEmailSecurityMaliciousTimeseryListParamsAggInterval = "1d"
	RadarEmailSecurityMaliciousTimeseryListParamsAggInterval1w  RadarEmailSecurityMaliciousTimeseryListParamsAggInterval = "1w"
)

type RadarEmailSecurityMaliciousTimeseryListParamsArc string

const (
	RadarEmailSecurityMaliciousTimeseryListParamsArcPass RadarEmailSecurityMaliciousTimeseryListParamsArc = "PASS"
	RadarEmailSecurityMaliciousTimeseryListParamsArcNone RadarEmailSecurityMaliciousTimeseryListParamsArc = "NONE"
	RadarEmailSecurityMaliciousTimeseryListParamsArcFail RadarEmailSecurityMaliciousTimeseryListParamsArc = "FAIL"
)

type RadarEmailSecurityMaliciousTimeseryListParamsDateRange string

const (
	RadarEmailSecurityMaliciousTimeseryListParamsDateRange1d         RadarEmailSecurityMaliciousTimeseryListParamsDateRange = "1d"
	RadarEmailSecurityMaliciousTimeseryListParamsDateRange2d         RadarEmailSecurityMaliciousTimeseryListParamsDateRange = "2d"
	RadarEmailSecurityMaliciousTimeseryListParamsDateRange7d         RadarEmailSecurityMaliciousTimeseryListParamsDateRange = "7d"
	RadarEmailSecurityMaliciousTimeseryListParamsDateRange14d        RadarEmailSecurityMaliciousTimeseryListParamsDateRange = "14d"
	RadarEmailSecurityMaliciousTimeseryListParamsDateRange28d        RadarEmailSecurityMaliciousTimeseryListParamsDateRange = "28d"
	RadarEmailSecurityMaliciousTimeseryListParamsDateRange12w        RadarEmailSecurityMaliciousTimeseryListParamsDateRange = "12w"
	RadarEmailSecurityMaliciousTimeseryListParamsDateRange24w        RadarEmailSecurityMaliciousTimeseryListParamsDateRange = "24w"
	RadarEmailSecurityMaliciousTimeseryListParamsDateRange52w        RadarEmailSecurityMaliciousTimeseryListParamsDateRange = "52w"
	RadarEmailSecurityMaliciousTimeseryListParamsDateRange1dControl  RadarEmailSecurityMaliciousTimeseryListParamsDateRange = "1dControl"
	RadarEmailSecurityMaliciousTimeseryListParamsDateRange2dControl  RadarEmailSecurityMaliciousTimeseryListParamsDateRange = "2dControl"
	RadarEmailSecurityMaliciousTimeseryListParamsDateRange7dControl  RadarEmailSecurityMaliciousTimeseryListParamsDateRange = "7dControl"
	RadarEmailSecurityMaliciousTimeseryListParamsDateRange14dControl RadarEmailSecurityMaliciousTimeseryListParamsDateRange = "14dControl"
	RadarEmailSecurityMaliciousTimeseryListParamsDateRange28dControl RadarEmailSecurityMaliciousTimeseryListParamsDateRange = "28dControl"
	RadarEmailSecurityMaliciousTimeseryListParamsDateRange12wControl RadarEmailSecurityMaliciousTimeseryListParamsDateRange = "12wControl"
	RadarEmailSecurityMaliciousTimeseryListParamsDateRange24wControl RadarEmailSecurityMaliciousTimeseryListParamsDateRange = "24wControl"
)

type RadarEmailSecurityMaliciousTimeseryListParamsDkim string

const (
	RadarEmailSecurityMaliciousTimeseryListParamsDkimPass RadarEmailSecurityMaliciousTimeseryListParamsDkim = "PASS"
	RadarEmailSecurityMaliciousTimeseryListParamsDkimNone RadarEmailSecurityMaliciousTimeseryListParamsDkim = "NONE"
	RadarEmailSecurityMaliciousTimeseryListParamsDkimFail RadarEmailSecurityMaliciousTimeseryListParamsDkim = "FAIL"
)

type RadarEmailSecurityMaliciousTimeseryListParamsDmarc string

const (
	RadarEmailSecurityMaliciousTimeseryListParamsDmarcPass RadarEmailSecurityMaliciousTimeseryListParamsDmarc = "PASS"
	RadarEmailSecurityMaliciousTimeseryListParamsDmarcNone RadarEmailSecurityMaliciousTimeseryListParamsDmarc = "NONE"
	RadarEmailSecurityMaliciousTimeseryListParamsDmarcFail RadarEmailSecurityMaliciousTimeseryListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityMaliciousTimeseryListParamsFormat string

const (
	RadarEmailSecurityMaliciousTimeseryListParamsFormatJson RadarEmailSecurityMaliciousTimeseryListParamsFormat = "JSON"
	RadarEmailSecurityMaliciousTimeseryListParamsFormatCsv  RadarEmailSecurityMaliciousTimeseryListParamsFormat = "CSV"
)

type RadarEmailSecurityMaliciousTimeseryListParamsSpf string

const (
	RadarEmailSecurityMaliciousTimeseryListParamsSpfPass RadarEmailSecurityMaliciousTimeseryListParamsSpf = "PASS"
	RadarEmailSecurityMaliciousTimeseryListParamsSpfNone RadarEmailSecurityMaliciousTimeseryListParamsSpf = "NONE"
	RadarEmailSecurityMaliciousTimeseryListParamsSpfFail RadarEmailSecurityMaliciousTimeseryListParamsSpf = "FAIL"
)
