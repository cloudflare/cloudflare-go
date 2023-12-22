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

// RadarAs112SummaryResponseCodeService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAs112SummaryResponseCodeService] method instead.
type RadarAs112SummaryResponseCodeService struct {
	Options []option.RequestOption
}

// NewRadarAs112SummaryResponseCodeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAs112SummaryResponseCodeService(opts ...option.RequestOption) (r *RadarAs112SummaryResponseCodeService) {
	r = &RadarAs112SummaryResponseCodeService{}
	r.Options = opts
	return
}

// Percentage distribution of dns requests classified per Response Codes.
func (r *RadarAs112SummaryResponseCodeService) List(ctx context.Context, query RadarAs112SummaryResponseCodeListParams, opts ...option.RequestOption) (res *RadarAs112SummaryResponseCodeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/summary/response_codes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAs112SummaryResponseCodeListResponse struct {
	Result  RadarAs112SummaryResponseCodeListResponseResult `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarAs112SummaryResponseCodeListResponseJSON   `json:"-"`
}

// radarAs112SummaryResponseCodeListResponseJSON contains the JSON metadata for the
// struct [RadarAs112SummaryResponseCodeListResponse]
type radarAs112SummaryResponseCodeListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryResponseCodeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryResponseCodeListResponseResult struct {
	Meta     interface{}                                             `json:"meta,required"`
	Summary0 RadarAs112SummaryResponseCodeListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarAs112SummaryResponseCodeListResponseResultJSON     `json:"-"`
}

// radarAs112SummaryResponseCodeListResponseResultJSON contains the JSON metadata
// for the struct [RadarAs112SummaryResponseCodeListResponseResult]
type radarAs112SummaryResponseCodeListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryResponseCodeListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryResponseCodeListResponseResultSummary0 struct {
	Noerror  string                                                      `json:"NOERROR,required"`
	Nxdomain string                                                      `json:"NXDOMAIN,required"`
	JSON     radarAs112SummaryResponseCodeListResponseResultSummary0JSON `json:"-"`
}

// radarAs112SummaryResponseCodeListResponseResultSummary0JSON contains the JSON
// metadata for the struct
// [RadarAs112SummaryResponseCodeListResponseResultSummary0]
type radarAs112SummaryResponseCodeListResponseResultSummary0JSON struct {
	Noerror     apijson.Field
	Nxdomain    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryResponseCodeListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryResponseCodeListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112SummaryResponseCodeListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112SummaryResponseCodeListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112SummaryResponseCodeListParams]'s query parameters
// as `url.Values`.
func (r RadarAs112SummaryResponseCodeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAs112SummaryResponseCodeListParamsDateRange string

const (
	RadarAs112SummaryResponseCodeListParamsDateRange1d         RadarAs112SummaryResponseCodeListParamsDateRange = "1d"
	RadarAs112SummaryResponseCodeListParamsDateRange7d         RadarAs112SummaryResponseCodeListParamsDateRange = "7d"
	RadarAs112SummaryResponseCodeListParamsDateRange14d        RadarAs112SummaryResponseCodeListParamsDateRange = "14d"
	RadarAs112SummaryResponseCodeListParamsDateRange28d        RadarAs112SummaryResponseCodeListParamsDateRange = "28d"
	RadarAs112SummaryResponseCodeListParamsDateRange12w        RadarAs112SummaryResponseCodeListParamsDateRange = "12w"
	RadarAs112SummaryResponseCodeListParamsDateRange24w        RadarAs112SummaryResponseCodeListParamsDateRange = "24w"
	RadarAs112SummaryResponseCodeListParamsDateRange52w        RadarAs112SummaryResponseCodeListParamsDateRange = "52w"
	RadarAs112SummaryResponseCodeListParamsDateRange1dControl  RadarAs112SummaryResponseCodeListParamsDateRange = "1dControl"
	RadarAs112SummaryResponseCodeListParamsDateRange7dControl  RadarAs112SummaryResponseCodeListParamsDateRange = "7dControl"
	RadarAs112SummaryResponseCodeListParamsDateRange14dControl RadarAs112SummaryResponseCodeListParamsDateRange = "14dControl"
	RadarAs112SummaryResponseCodeListParamsDateRange28dControl RadarAs112SummaryResponseCodeListParamsDateRange = "28dControl"
	RadarAs112SummaryResponseCodeListParamsDateRange12wControl RadarAs112SummaryResponseCodeListParamsDateRange = "12wControl"
	RadarAs112SummaryResponseCodeListParamsDateRange24wControl RadarAs112SummaryResponseCodeListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112SummaryResponseCodeListParamsFormat string

const (
	RadarAs112SummaryResponseCodeListParamsFormatJson RadarAs112SummaryResponseCodeListParamsFormat = "JSON"
	RadarAs112SummaryResponseCodeListParamsFormatCsv  RadarAs112SummaryResponseCodeListParamsFormat = "CSV"
)
