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

// RadarEmailSecuritySpamService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEmailSecuritySpamService]
// method instead.
type RadarEmailSecuritySpamService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecuritySpamService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarEmailSecuritySpamService(opts ...option.RequestOption) (r *RadarEmailSecuritySpamService) {
	r = &RadarEmailSecuritySpamService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified as SPAM over time.
func (r *RadarEmailSecuritySpamService) List(ctx context.Context, query RadarEmailSecuritySpamListParams, opts ...option.RequestOption) (res *RadarEmailSecuritySpamListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/timeseries_groups/spam"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSecuritySpamListResponse struct {
	Result  RadarEmailSecuritySpamListResponseResult `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    radarEmailSecuritySpamListResponseJSON   `json:"-"`
}

// radarEmailSecuritySpamListResponseJSON contains the JSON metadata for the struct
// [RadarEmailSecuritySpamListResponse]
type radarEmailSecuritySpamListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySpamListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySpamListResponseResult struct {
	Meta   interface{}                                    `json:"meta,required"`
	Serie0 RadarEmailSecuritySpamListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecuritySpamListResponseResultJSON   `json:"-"`
}

// radarEmailSecuritySpamListResponseResultJSON contains the JSON metadata for the
// struct [RadarEmailSecuritySpamListResponseResult]
type radarEmailSecuritySpamListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySpamListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySpamListResponseResultSerie0 struct {
	NotSpam []string                                           `json:"NOT_SPAM,required"`
	Spam    []string                                           `json:"SPAM,required"`
	JSON    radarEmailSecuritySpamListResponseResultSerie0JSON `json:"-"`
}

// radarEmailSecuritySpamListResponseResultSerie0JSON contains the JSON metadata
// for the struct [RadarEmailSecuritySpamListResponseResultSerie0]
type radarEmailSecuritySpamListResponseResultSerie0JSON struct {
	NotSpam     apijson.Field
	Spam        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySpamListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySpamListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecuritySpamListParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecuritySpamListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecuritySpamListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSecuritySpamListParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecuritySpamListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecuritySpamListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailSecuritySpamListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecuritySpamListParams]'s query parameters as
// `url.Values`.
func (r RadarEmailSecuritySpamListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecuritySpamListParamsAggInterval string

const (
	RadarEmailSecuritySpamListParamsAggInterval15m RadarEmailSecuritySpamListParamsAggInterval = "15m"
	RadarEmailSecuritySpamListParamsAggInterval1h  RadarEmailSecuritySpamListParamsAggInterval = "1h"
	RadarEmailSecuritySpamListParamsAggInterval1d  RadarEmailSecuritySpamListParamsAggInterval = "1d"
	RadarEmailSecuritySpamListParamsAggInterval1w  RadarEmailSecuritySpamListParamsAggInterval = "1w"
)

type RadarEmailSecuritySpamListParamsArc string

const (
	RadarEmailSecuritySpamListParamsArcPass RadarEmailSecuritySpamListParamsArc = "PASS"
	RadarEmailSecuritySpamListParamsArcNone RadarEmailSecuritySpamListParamsArc = "NONE"
	RadarEmailSecuritySpamListParamsArcFail RadarEmailSecuritySpamListParamsArc = "FAIL"
)

type RadarEmailSecuritySpamListParamsDateRange string

const (
	RadarEmailSecuritySpamListParamsDateRange1d         RadarEmailSecuritySpamListParamsDateRange = "1d"
	RadarEmailSecuritySpamListParamsDateRange2d         RadarEmailSecuritySpamListParamsDateRange = "2d"
	RadarEmailSecuritySpamListParamsDateRange7d         RadarEmailSecuritySpamListParamsDateRange = "7d"
	RadarEmailSecuritySpamListParamsDateRange14d        RadarEmailSecuritySpamListParamsDateRange = "14d"
	RadarEmailSecuritySpamListParamsDateRange28d        RadarEmailSecuritySpamListParamsDateRange = "28d"
	RadarEmailSecuritySpamListParamsDateRange12w        RadarEmailSecuritySpamListParamsDateRange = "12w"
	RadarEmailSecuritySpamListParamsDateRange24w        RadarEmailSecuritySpamListParamsDateRange = "24w"
	RadarEmailSecuritySpamListParamsDateRange52w        RadarEmailSecuritySpamListParamsDateRange = "52w"
	RadarEmailSecuritySpamListParamsDateRange1dControl  RadarEmailSecuritySpamListParamsDateRange = "1dControl"
	RadarEmailSecuritySpamListParamsDateRange2dControl  RadarEmailSecuritySpamListParamsDateRange = "2dControl"
	RadarEmailSecuritySpamListParamsDateRange7dControl  RadarEmailSecuritySpamListParamsDateRange = "7dControl"
	RadarEmailSecuritySpamListParamsDateRange14dControl RadarEmailSecuritySpamListParamsDateRange = "14dControl"
	RadarEmailSecuritySpamListParamsDateRange28dControl RadarEmailSecuritySpamListParamsDateRange = "28dControl"
	RadarEmailSecuritySpamListParamsDateRange12wControl RadarEmailSecuritySpamListParamsDateRange = "12wControl"
	RadarEmailSecuritySpamListParamsDateRange24wControl RadarEmailSecuritySpamListParamsDateRange = "24wControl"
)

type RadarEmailSecuritySpamListParamsDkim string

const (
	RadarEmailSecuritySpamListParamsDkimPass RadarEmailSecuritySpamListParamsDkim = "PASS"
	RadarEmailSecuritySpamListParamsDkimNone RadarEmailSecuritySpamListParamsDkim = "NONE"
	RadarEmailSecuritySpamListParamsDkimFail RadarEmailSecuritySpamListParamsDkim = "FAIL"
)

type RadarEmailSecuritySpamListParamsDmarc string

const (
	RadarEmailSecuritySpamListParamsDmarcPass RadarEmailSecuritySpamListParamsDmarc = "PASS"
	RadarEmailSecuritySpamListParamsDmarcNone RadarEmailSecuritySpamListParamsDmarc = "NONE"
	RadarEmailSecuritySpamListParamsDmarcFail RadarEmailSecuritySpamListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecuritySpamListParamsFormat string

const (
	RadarEmailSecuritySpamListParamsFormatJson RadarEmailSecuritySpamListParamsFormat = "JSON"
	RadarEmailSecuritySpamListParamsFormatCsv  RadarEmailSecuritySpamListParamsFormat = "CSV"
)

type RadarEmailSecuritySpamListParamsSpf string

const (
	RadarEmailSecuritySpamListParamsSpfPass RadarEmailSecuritySpamListParamsSpf = "PASS"
	RadarEmailSecuritySpamListParamsSpfNone RadarEmailSecuritySpamListParamsSpf = "NONE"
	RadarEmailSecuritySpamListParamsSpfFail RadarEmailSecuritySpamListParamsSpf = "FAIL"
)
