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

// RadarEmailTimeseryMaliciousService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarEmailTimeseryMaliciousService] method instead.
type RadarEmailTimeseryMaliciousService struct {
	Options []option.RequestOption
}

// NewRadarEmailTimeseryMaliciousService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailTimeseryMaliciousService(opts ...option.RequestOption) (r *RadarEmailTimeseryMaliciousService) {
	r = &RadarEmailTimeseryMaliciousService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified as MALICIOUS over time.
func (r *RadarEmailTimeseryMaliciousService) List(ctx context.Context, query RadarEmailTimeseryMaliciousListParams, opts ...option.RequestOption) (res *RadarEmailTimeseryMaliciousListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/timeseries/malicious"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailTimeseryMaliciousListResponse struct {
	Result  RadarEmailTimeseryMaliciousListResponseResult `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    radarEmailTimeseryMaliciousListResponseJSON   `json:"-"`
}

// radarEmailTimeseryMaliciousListResponseJSON contains the JSON metadata for the
// struct [RadarEmailTimeseryMaliciousListResponse]
type radarEmailTimeseryMaliciousListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTimeseryMaliciousListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTimeseryMaliciousListResponseResult struct {
	Meta   interface{}                                         `json:"meta,required"`
	Serie0 RadarEmailTimeseryMaliciousListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailTimeseryMaliciousListResponseResultJSON   `json:"-"`
}

// radarEmailTimeseryMaliciousListResponseResultJSON contains the JSON metadata for
// the struct [RadarEmailTimeseryMaliciousListResponseResult]
type radarEmailTimeseryMaliciousListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTimeseryMaliciousListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTimeseryMaliciousListResponseResultSerie0 struct {
	Malicious    []string                                                `json:"MALICIOUS,required"`
	NotMalicious []string                                                `json:"NOT_MALICIOUS,required"`
	JSON         radarEmailTimeseryMaliciousListResponseResultSerie0JSON `json:"-"`
}

// radarEmailTimeseryMaliciousListResponseResultSerie0JSON contains the JSON
// metadata for the struct [RadarEmailTimeseryMaliciousListResponseResultSerie0]
type radarEmailTimeseryMaliciousListResponseResultSerie0JSON struct {
	Malicious    apijson.Field
	NotMalicious apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailTimeseryMaliciousListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTimeseryMaliciousListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailTimeseryMaliciousListParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailTimeseryMaliciousListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailTimeseryMaliciousListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailTimeseryMaliciousListParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailTimeseryMaliciousListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailTimeseryMaliciousListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailTimeseryMaliciousListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailTimeseryMaliciousListParams]'s query parameters
// as `url.Values`.
func (r RadarEmailTimeseryMaliciousListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailTimeseryMaliciousListParamsAggInterval string

const (
	RadarEmailTimeseryMaliciousListParamsAggInterval15m RadarEmailTimeseryMaliciousListParamsAggInterval = "15m"
	RadarEmailTimeseryMaliciousListParamsAggInterval1h  RadarEmailTimeseryMaliciousListParamsAggInterval = "1h"
	RadarEmailTimeseryMaliciousListParamsAggInterval1d  RadarEmailTimeseryMaliciousListParamsAggInterval = "1d"
	RadarEmailTimeseryMaliciousListParamsAggInterval1w  RadarEmailTimeseryMaliciousListParamsAggInterval = "1w"
)

type RadarEmailTimeseryMaliciousListParamsArc string

const (
	RadarEmailTimeseryMaliciousListParamsArcPass RadarEmailTimeseryMaliciousListParamsArc = "PASS"
	RadarEmailTimeseryMaliciousListParamsArcNone RadarEmailTimeseryMaliciousListParamsArc = "NONE"
	RadarEmailTimeseryMaliciousListParamsArcFail RadarEmailTimeseryMaliciousListParamsArc = "FAIL"
)

type RadarEmailTimeseryMaliciousListParamsDateRange string

const (
	RadarEmailTimeseryMaliciousListParamsDateRange1d         RadarEmailTimeseryMaliciousListParamsDateRange = "1d"
	RadarEmailTimeseryMaliciousListParamsDateRange7d         RadarEmailTimeseryMaliciousListParamsDateRange = "7d"
	RadarEmailTimeseryMaliciousListParamsDateRange14d        RadarEmailTimeseryMaliciousListParamsDateRange = "14d"
	RadarEmailTimeseryMaliciousListParamsDateRange28d        RadarEmailTimeseryMaliciousListParamsDateRange = "28d"
	RadarEmailTimeseryMaliciousListParamsDateRange12w        RadarEmailTimeseryMaliciousListParamsDateRange = "12w"
	RadarEmailTimeseryMaliciousListParamsDateRange24w        RadarEmailTimeseryMaliciousListParamsDateRange = "24w"
	RadarEmailTimeseryMaliciousListParamsDateRange52w        RadarEmailTimeseryMaliciousListParamsDateRange = "52w"
	RadarEmailTimeseryMaliciousListParamsDateRange1dControl  RadarEmailTimeseryMaliciousListParamsDateRange = "1dControl"
	RadarEmailTimeseryMaliciousListParamsDateRange7dControl  RadarEmailTimeseryMaliciousListParamsDateRange = "7dControl"
	RadarEmailTimeseryMaliciousListParamsDateRange14dControl RadarEmailTimeseryMaliciousListParamsDateRange = "14dControl"
	RadarEmailTimeseryMaliciousListParamsDateRange28dControl RadarEmailTimeseryMaliciousListParamsDateRange = "28dControl"
	RadarEmailTimeseryMaliciousListParamsDateRange12wControl RadarEmailTimeseryMaliciousListParamsDateRange = "12wControl"
	RadarEmailTimeseryMaliciousListParamsDateRange24wControl RadarEmailTimeseryMaliciousListParamsDateRange = "24wControl"
)

type RadarEmailTimeseryMaliciousListParamsDkim string

const (
	RadarEmailTimeseryMaliciousListParamsDkimPass RadarEmailTimeseryMaliciousListParamsDkim = "PASS"
	RadarEmailTimeseryMaliciousListParamsDkimNone RadarEmailTimeseryMaliciousListParamsDkim = "NONE"
	RadarEmailTimeseryMaliciousListParamsDkimFail RadarEmailTimeseryMaliciousListParamsDkim = "FAIL"
)

type RadarEmailTimeseryMaliciousListParamsDmarc string

const (
	RadarEmailTimeseryMaliciousListParamsDmarcPass RadarEmailTimeseryMaliciousListParamsDmarc = "PASS"
	RadarEmailTimeseryMaliciousListParamsDmarcNone RadarEmailTimeseryMaliciousListParamsDmarc = "NONE"
	RadarEmailTimeseryMaliciousListParamsDmarcFail RadarEmailTimeseryMaliciousListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailTimeseryMaliciousListParamsFormat string

const (
	RadarEmailTimeseryMaliciousListParamsFormatJson RadarEmailTimeseryMaliciousListParamsFormat = "JSON"
	RadarEmailTimeseryMaliciousListParamsFormatCsv  RadarEmailTimeseryMaliciousListParamsFormat = "CSV"
)

type RadarEmailTimeseryMaliciousListParamsSpf string

const (
	RadarEmailTimeseryMaliciousListParamsSpfPass RadarEmailTimeseryMaliciousListParamsSpf = "PASS"
	RadarEmailTimeseryMaliciousListParamsSpfNone RadarEmailTimeseryMaliciousListParamsSpf = "NONE"
	RadarEmailTimeseryMaliciousListParamsSpfFail RadarEmailTimeseryMaliciousListParamsSpf = "FAIL"
)
