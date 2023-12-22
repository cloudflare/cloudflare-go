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

// RadarAs112SummaryQueryTypeService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarAs112SummaryQueryTypeService] method instead.
type RadarAs112SummaryQueryTypeService struct {
	Options []option.RequestOption
}

// NewRadarAs112SummaryQueryTypeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAs112SummaryQueryTypeService(opts ...option.RequestOption) (r *RadarAs112SummaryQueryTypeService) {
	r = &RadarAs112SummaryQueryTypeService{}
	r.Options = opts
	return
}

// Percentage distribution of dns requests classified per Query Type.
func (r *RadarAs112SummaryQueryTypeService) List(ctx context.Context, query RadarAs112SummaryQueryTypeListParams, opts ...option.RequestOption) (res *RadarAs112SummaryQueryTypeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/summary/query_type"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAs112SummaryQueryTypeListResponse struct {
	Result  RadarAs112SummaryQueryTypeListResponseResult `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarAs112SummaryQueryTypeListResponseJSON   `json:"-"`
}

// radarAs112SummaryQueryTypeListResponseJSON contains the JSON metadata for the
// struct [RadarAs112SummaryQueryTypeListResponse]
type radarAs112SummaryQueryTypeListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryQueryTypeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryQueryTypeListResponseResult struct {
	Meta     interface{}                                          `json:"meta,required"`
	Summary0 RadarAs112SummaryQueryTypeListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarAs112SummaryQueryTypeListResponseResultJSON     `json:"-"`
}

// radarAs112SummaryQueryTypeListResponseResultJSON contains the JSON metadata for
// the struct [RadarAs112SummaryQueryTypeListResponseResult]
type radarAs112SummaryQueryTypeListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryQueryTypeListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryQueryTypeListResponseResultSummary0 struct {
	A    string                                                   `json:"A,required"`
	Aaaa string                                                   `json:"AAAA,required"`
	Ptr  string                                                   `json:"PTR,required"`
	Soa  string                                                   `json:"SOA,required"`
	Srv  string                                                   `json:"SRV,required"`
	JSON radarAs112SummaryQueryTypeListResponseResultSummary0JSON `json:"-"`
}

// radarAs112SummaryQueryTypeListResponseResultSummary0JSON contains the JSON
// metadata for the struct [RadarAs112SummaryQueryTypeListResponseResultSummary0]
type radarAs112SummaryQueryTypeListResponseResultSummary0JSON struct {
	A           apijson.Field
	Aaaa        apijson.Field
	Ptr         apijson.Field
	Soa         apijson.Field
	Srv         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryQueryTypeListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryQueryTypeListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112SummaryQueryTypeListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112SummaryQueryTypeListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112SummaryQueryTypeListParams]'s query parameters as
// `url.Values`.
func (r RadarAs112SummaryQueryTypeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAs112SummaryQueryTypeListParamsDateRange string

const (
	RadarAs112SummaryQueryTypeListParamsDateRange1d         RadarAs112SummaryQueryTypeListParamsDateRange = "1d"
	RadarAs112SummaryQueryTypeListParamsDateRange7d         RadarAs112SummaryQueryTypeListParamsDateRange = "7d"
	RadarAs112SummaryQueryTypeListParamsDateRange14d        RadarAs112SummaryQueryTypeListParamsDateRange = "14d"
	RadarAs112SummaryQueryTypeListParamsDateRange28d        RadarAs112SummaryQueryTypeListParamsDateRange = "28d"
	RadarAs112SummaryQueryTypeListParamsDateRange12w        RadarAs112SummaryQueryTypeListParamsDateRange = "12w"
	RadarAs112SummaryQueryTypeListParamsDateRange24w        RadarAs112SummaryQueryTypeListParamsDateRange = "24w"
	RadarAs112SummaryQueryTypeListParamsDateRange52w        RadarAs112SummaryQueryTypeListParamsDateRange = "52w"
	RadarAs112SummaryQueryTypeListParamsDateRange1dControl  RadarAs112SummaryQueryTypeListParamsDateRange = "1dControl"
	RadarAs112SummaryQueryTypeListParamsDateRange7dControl  RadarAs112SummaryQueryTypeListParamsDateRange = "7dControl"
	RadarAs112SummaryQueryTypeListParamsDateRange14dControl RadarAs112SummaryQueryTypeListParamsDateRange = "14dControl"
	RadarAs112SummaryQueryTypeListParamsDateRange28dControl RadarAs112SummaryQueryTypeListParamsDateRange = "28dControl"
	RadarAs112SummaryQueryTypeListParamsDateRange12wControl RadarAs112SummaryQueryTypeListParamsDateRange = "12wControl"
	RadarAs112SummaryQueryTypeListParamsDateRange24wControl RadarAs112SummaryQueryTypeListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112SummaryQueryTypeListParamsFormat string

const (
	RadarAs112SummaryQueryTypeListParamsFormatJson RadarAs112SummaryQueryTypeListParamsFormat = "JSON"
	RadarAs112SummaryQueryTypeListParamsFormatCsv  RadarAs112SummaryQueryTypeListParamsFormat = "CSV"
)
