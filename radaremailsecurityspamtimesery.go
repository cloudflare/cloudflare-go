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

// RadarEmailSecuritySpamTimeseryService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecuritySpamTimeseryService] method instead.
type RadarEmailSecuritySpamTimeseryService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecuritySpamTimeseryService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecuritySpamTimeseryService(opts ...option.RequestOption) (r *RadarEmailSecuritySpamTimeseryService) {
	r = &RadarEmailSecuritySpamTimeseryService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified as SPAM over time.
func (r *RadarEmailSecuritySpamTimeseryService) List(ctx context.Context, query RadarEmailSecuritySpamTimeseryListParams, opts ...option.RequestOption) (res *RadarEmailSecuritySpamTimeseryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/timeseries_groups/spam"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSecuritySpamTimeseryListResponse struct {
	Result  RadarEmailSecuritySpamTimeseryListResponseResult `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarEmailSecuritySpamTimeseryListResponseJSON   `json:"-"`
}

// radarEmailSecuritySpamTimeseryListResponseJSON contains the JSON metadata for
// the struct [RadarEmailSecuritySpamTimeseryListResponse]
type radarEmailSecuritySpamTimeseryListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySpamTimeseryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySpamTimeseryListResponseResult struct {
	Meta   interface{}                                            `json:"meta,required"`
	Serie0 RadarEmailSecuritySpamTimeseryListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecuritySpamTimeseryListResponseResultJSON   `json:"-"`
}

// radarEmailSecuritySpamTimeseryListResponseResultJSON contains the JSON metadata
// for the struct [RadarEmailSecuritySpamTimeseryListResponseResult]
type radarEmailSecuritySpamTimeseryListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySpamTimeseryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySpamTimeseryListResponseResultSerie0 struct {
	NotSpam []string                                                   `json:"NOT_SPAM,required"`
	Spam    []string                                                   `json:"SPAM,required"`
	JSON    radarEmailSecuritySpamTimeseryListResponseResultSerie0JSON `json:"-"`
}

// radarEmailSecuritySpamTimeseryListResponseResultSerie0JSON contains the JSON
// metadata for the struct [RadarEmailSecuritySpamTimeseryListResponseResultSerie0]
type radarEmailSecuritySpamTimeseryListResponseResultSerie0JSON struct {
	NotSpam     apijson.Field
	Spam        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySpamTimeseryListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySpamTimeseryListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecuritySpamTimeseryListParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecuritySpamTimeseryListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecuritySpamTimeseryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSecuritySpamTimeseryListParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecuritySpamTimeseryListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecuritySpamTimeseryListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailSecuritySpamTimeseryListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecuritySpamTimeseryListParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecuritySpamTimeseryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecuritySpamTimeseryListParamsAggInterval string

const (
	RadarEmailSecuritySpamTimeseryListParamsAggInterval15m RadarEmailSecuritySpamTimeseryListParamsAggInterval = "15m"
	RadarEmailSecuritySpamTimeseryListParamsAggInterval1h  RadarEmailSecuritySpamTimeseryListParamsAggInterval = "1h"
	RadarEmailSecuritySpamTimeseryListParamsAggInterval1d  RadarEmailSecuritySpamTimeseryListParamsAggInterval = "1d"
	RadarEmailSecuritySpamTimeseryListParamsAggInterval1w  RadarEmailSecuritySpamTimeseryListParamsAggInterval = "1w"
)

type RadarEmailSecuritySpamTimeseryListParamsArc string

const (
	RadarEmailSecuritySpamTimeseryListParamsArcPass RadarEmailSecuritySpamTimeseryListParamsArc = "PASS"
	RadarEmailSecuritySpamTimeseryListParamsArcNone RadarEmailSecuritySpamTimeseryListParamsArc = "NONE"
	RadarEmailSecuritySpamTimeseryListParamsArcFail RadarEmailSecuritySpamTimeseryListParamsArc = "FAIL"
)

type RadarEmailSecuritySpamTimeseryListParamsDateRange string

const (
	RadarEmailSecuritySpamTimeseryListParamsDateRange1d         RadarEmailSecuritySpamTimeseryListParamsDateRange = "1d"
	RadarEmailSecuritySpamTimeseryListParamsDateRange2d         RadarEmailSecuritySpamTimeseryListParamsDateRange = "2d"
	RadarEmailSecuritySpamTimeseryListParamsDateRange7d         RadarEmailSecuritySpamTimeseryListParamsDateRange = "7d"
	RadarEmailSecuritySpamTimeseryListParamsDateRange14d        RadarEmailSecuritySpamTimeseryListParamsDateRange = "14d"
	RadarEmailSecuritySpamTimeseryListParamsDateRange28d        RadarEmailSecuritySpamTimeseryListParamsDateRange = "28d"
	RadarEmailSecuritySpamTimeseryListParamsDateRange12w        RadarEmailSecuritySpamTimeseryListParamsDateRange = "12w"
	RadarEmailSecuritySpamTimeseryListParamsDateRange24w        RadarEmailSecuritySpamTimeseryListParamsDateRange = "24w"
	RadarEmailSecuritySpamTimeseryListParamsDateRange52w        RadarEmailSecuritySpamTimeseryListParamsDateRange = "52w"
	RadarEmailSecuritySpamTimeseryListParamsDateRange1dControl  RadarEmailSecuritySpamTimeseryListParamsDateRange = "1dControl"
	RadarEmailSecuritySpamTimeseryListParamsDateRange2dControl  RadarEmailSecuritySpamTimeseryListParamsDateRange = "2dControl"
	RadarEmailSecuritySpamTimeseryListParamsDateRange7dControl  RadarEmailSecuritySpamTimeseryListParamsDateRange = "7dControl"
	RadarEmailSecuritySpamTimeseryListParamsDateRange14dControl RadarEmailSecuritySpamTimeseryListParamsDateRange = "14dControl"
	RadarEmailSecuritySpamTimeseryListParamsDateRange28dControl RadarEmailSecuritySpamTimeseryListParamsDateRange = "28dControl"
	RadarEmailSecuritySpamTimeseryListParamsDateRange12wControl RadarEmailSecuritySpamTimeseryListParamsDateRange = "12wControl"
	RadarEmailSecuritySpamTimeseryListParamsDateRange24wControl RadarEmailSecuritySpamTimeseryListParamsDateRange = "24wControl"
)

type RadarEmailSecuritySpamTimeseryListParamsDkim string

const (
	RadarEmailSecuritySpamTimeseryListParamsDkimPass RadarEmailSecuritySpamTimeseryListParamsDkim = "PASS"
	RadarEmailSecuritySpamTimeseryListParamsDkimNone RadarEmailSecuritySpamTimeseryListParamsDkim = "NONE"
	RadarEmailSecuritySpamTimeseryListParamsDkimFail RadarEmailSecuritySpamTimeseryListParamsDkim = "FAIL"
)

type RadarEmailSecuritySpamTimeseryListParamsDmarc string

const (
	RadarEmailSecuritySpamTimeseryListParamsDmarcPass RadarEmailSecuritySpamTimeseryListParamsDmarc = "PASS"
	RadarEmailSecuritySpamTimeseryListParamsDmarcNone RadarEmailSecuritySpamTimeseryListParamsDmarc = "NONE"
	RadarEmailSecuritySpamTimeseryListParamsDmarcFail RadarEmailSecuritySpamTimeseryListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecuritySpamTimeseryListParamsFormat string

const (
	RadarEmailSecuritySpamTimeseryListParamsFormatJson RadarEmailSecuritySpamTimeseryListParamsFormat = "JSON"
	RadarEmailSecuritySpamTimeseryListParamsFormatCsv  RadarEmailSecuritySpamTimeseryListParamsFormat = "CSV"
)

type RadarEmailSecuritySpamTimeseryListParamsSpf string

const (
	RadarEmailSecuritySpamTimeseryListParamsSpfPass RadarEmailSecuritySpamTimeseryListParamsSpf = "PASS"
	RadarEmailSecuritySpamTimeseryListParamsSpfNone RadarEmailSecuritySpamTimeseryListParamsSpf = "NONE"
	RadarEmailSecuritySpamTimeseryListParamsSpfFail RadarEmailSecuritySpamTimeseryListParamsSpf = "FAIL"
)
