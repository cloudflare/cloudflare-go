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

// RadarEmailTimeserySpamService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEmailTimeserySpamService]
// method instead.
type RadarEmailTimeserySpamService struct {
	Options []option.RequestOption
}

// NewRadarEmailTimeserySpamService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarEmailTimeserySpamService(opts ...option.RequestOption) (r *RadarEmailTimeserySpamService) {
	r = &RadarEmailTimeserySpamService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified as SPAM over time.
func (r *RadarEmailTimeserySpamService) List(ctx context.Context, query RadarEmailTimeserySpamListParams, opts ...option.RequestOption) (res *RadarEmailTimeserySpamListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/timeseries/spam"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailTimeserySpamListResponse struct {
	Result  RadarEmailTimeserySpamListResponseResult `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    radarEmailTimeserySpamListResponseJSON   `json:"-"`
}

// radarEmailTimeserySpamListResponseJSON contains the JSON metadata for the struct
// [RadarEmailTimeserySpamListResponse]
type radarEmailTimeserySpamListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTimeserySpamListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTimeserySpamListResponseResult struct {
	Meta   interface{}                                    `json:"meta,required"`
	Serie0 RadarEmailTimeserySpamListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailTimeserySpamListResponseResultJSON   `json:"-"`
}

// radarEmailTimeserySpamListResponseResultJSON contains the JSON metadata for the
// struct [RadarEmailTimeserySpamListResponseResult]
type radarEmailTimeserySpamListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTimeserySpamListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTimeserySpamListResponseResultSerie0 struct {
	NotSpam []string                                           `json:"NOT_SPAM,required"`
	Spam    []string                                           `json:"SPAM,required"`
	JSON    radarEmailTimeserySpamListResponseResultSerie0JSON `json:"-"`
}

// radarEmailTimeserySpamListResponseResultSerie0JSON contains the JSON metadata
// for the struct [RadarEmailTimeserySpamListResponseResultSerie0]
type radarEmailTimeserySpamListResponseResultSerie0JSON struct {
	NotSpam     apijson.Field
	Spam        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTimeserySpamListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTimeserySpamListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailTimeserySpamListParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailTimeserySpamListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailTimeserySpamListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailTimeserySpamListParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailTimeserySpamListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailTimeserySpamListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailTimeserySpamListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailTimeserySpamListParams]'s query parameters as
// `url.Values`.
func (r RadarEmailTimeserySpamListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailTimeserySpamListParamsAggInterval string

const (
	RadarEmailTimeserySpamListParamsAggInterval15m RadarEmailTimeserySpamListParamsAggInterval = "15m"
	RadarEmailTimeserySpamListParamsAggInterval1h  RadarEmailTimeserySpamListParamsAggInterval = "1h"
	RadarEmailTimeserySpamListParamsAggInterval1d  RadarEmailTimeserySpamListParamsAggInterval = "1d"
	RadarEmailTimeserySpamListParamsAggInterval1w  RadarEmailTimeserySpamListParamsAggInterval = "1w"
)

type RadarEmailTimeserySpamListParamsArc string

const (
	RadarEmailTimeserySpamListParamsArcPass RadarEmailTimeserySpamListParamsArc = "PASS"
	RadarEmailTimeserySpamListParamsArcNone RadarEmailTimeserySpamListParamsArc = "NONE"
	RadarEmailTimeserySpamListParamsArcFail RadarEmailTimeserySpamListParamsArc = "FAIL"
)

type RadarEmailTimeserySpamListParamsDateRange string

const (
	RadarEmailTimeserySpamListParamsDateRange1d         RadarEmailTimeserySpamListParamsDateRange = "1d"
	RadarEmailTimeserySpamListParamsDateRange7d         RadarEmailTimeserySpamListParamsDateRange = "7d"
	RadarEmailTimeserySpamListParamsDateRange14d        RadarEmailTimeserySpamListParamsDateRange = "14d"
	RadarEmailTimeserySpamListParamsDateRange28d        RadarEmailTimeserySpamListParamsDateRange = "28d"
	RadarEmailTimeserySpamListParamsDateRange12w        RadarEmailTimeserySpamListParamsDateRange = "12w"
	RadarEmailTimeserySpamListParamsDateRange24w        RadarEmailTimeserySpamListParamsDateRange = "24w"
	RadarEmailTimeserySpamListParamsDateRange52w        RadarEmailTimeserySpamListParamsDateRange = "52w"
	RadarEmailTimeserySpamListParamsDateRange1dControl  RadarEmailTimeserySpamListParamsDateRange = "1dControl"
	RadarEmailTimeserySpamListParamsDateRange7dControl  RadarEmailTimeserySpamListParamsDateRange = "7dControl"
	RadarEmailTimeserySpamListParamsDateRange14dControl RadarEmailTimeserySpamListParamsDateRange = "14dControl"
	RadarEmailTimeserySpamListParamsDateRange28dControl RadarEmailTimeserySpamListParamsDateRange = "28dControl"
	RadarEmailTimeserySpamListParamsDateRange12wControl RadarEmailTimeserySpamListParamsDateRange = "12wControl"
	RadarEmailTimeserySpamListParamsDateRange24wControl RadarEmailTimeserySpamListParamsDateRange = "24wControl"
)

type RadarEmailTimeserySpamListParamsDkim string

const (
	RadarEmailTimeserySpamListParamsDkimPass RadarEmailTimeserySpamListParamsDkim = "PASS"
	RadarEmailTimeserySpamListParamsDkimNone RadarEmailTimeserySpamListParamsDkim = "NONE"
	RadarEmailTimeserySpamListParamsDkimFail RadarEmailTimeserySpamListParamsDkim = "FAIL"
)

type RadarEmailTimeserySpamListParamsDmarc string

const (
	RadarEmailTimeserySpamListParamsDmarcPass RadarEmailTimeserySpamListParamsDmarc = "PASS"
	RadarEmailTimeserySpamListParamsDmarcNone RadarEmailTimeserySpamListParamsDmarc = "NONE"
	RadarEmailTimeserySpamListParamsDmarcFail RadarEmailTimeserySpamListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailTimeserySpamListParamsFormat string

const (
	RadarEmailTimeserySpamListParamsFormatJson RadarEmailTimeserySpamListParamsFormat = "JSON"
	RadarEmailTimeserySpamListParamsFormatCsv  RadarEmailTimeserySpamListParamsFormat = "CSV"
)

type RadarEmailTimeserySpamListParamsSpf string

const (
	RadarEmailTimeserySpamListParamsSpfPass RadarEmailTimeserySpamListParamsSpf = "PASS"
	RadarEmailTimeserySpamListParamsSpfNone RadarEmailTimeserySpamListParamsSpf = "NONE"
	RadarEmailTimeserySpamListParamsSpfFail RadarEmailTimeserySpamListParamsSpf = "FAIL"
)
