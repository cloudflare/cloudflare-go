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

// RadarAs112SummaryIPVersionService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarAs112SummaryIPVersionService] method instead.
type RadarAs112SummaryIPVersionService struct {
	Options []option.RequestOption
}

// NewRadarAs112SummaryIPVersionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAs112SummaryIPVersionService(opts ...option.RequestOption) (r *RadarAs112SummaryIPVersionService) {
	r = &RadarAs112SummaryIPVersionService{}
	r.Options = opts
	return
}

// Percentage distribution of dns requests classified per IP Version.
func (r *RadarAs112SummaryIPVersionService) List(ctx context.Context, query RadarAs112SummaryIPVersionListParams, opts ...option.RequestOption) (res *RadarAs112SummaryIPVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/summary/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAs112SummaryIPVersionListResponse struct {
	Result  RadarAs112SummaryIPVersionListResponseResult `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarAs112SummaryIPVersionListResponseJSON   `json:"-"`
}

// radarAs112SummaryIPVersionListResponseJSON contains the JSON metadata for the
// struct [RadarAs112SummaryIPVersionListResponse]
type radarAs112SummaryIPVersionListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryIPVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryIPVersionListResponseResult struct {
	Meta     interface{}                                          `json:"meta,required"`
	Summary0 RadarAs112SummaryIPVersionListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarAs112SummaryIPVersionListResponseResultJSON     `json:"-"`
}

// radarAs112SummaryIPVersionListResponseResultJSON contains the JSON metadata for
// the struct [RadarAs112SummaryIPVersionListResponseResult]
type radarAs112SummaryIPVersionListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryIPVersionListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryIPVersionListResponseResultSummary0 struct {
	IPv4 string                                                   `json:"IPv4,required"`
	IPv6 string                                                   `json:"IPv6,required"`
	JSON radarAs112SummaryIPVersionListResponseResultSummary0JSON `json:"-"`
}

// radarAs112SummaryIPVersionListResponseResultSummary0JSON contains the JSON
// metadata for the struct [RadarAs112SummaryIPVersionListResponseResultSummary0]
type radarAs112SummaryIPVersionListResponseResultSummary0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryIPVersionListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryIPVersionListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112SummaryIPVersionListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112SummaryIPVersionListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112SummaryIPVersionListParams]'s query parameters as
// `url.Values`.
func (r RadarAs112SummaryIPVersionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAs112SummaryIPVersionListParamsDateRange string

const (
	RadarAs112SummaryIPVersionListParamsDateRange1d         RadarAs112SummaryIPVersionListParamsDateRange = "1d"
	RadarAs112SummaryIPVersionListParamsDateRange7d         RadarAs112SummaryIPVersionListParamsDateRange = "7d"
	RadarAs112SummaryIPVersionListParamsDateRange14d        RadarAs112SummaryIPVersionListParamsDateRange = "14d"
	RadarAs112SummaryIPVersionListParamsDateRange28d        RadarAs112SummaryIPVersionListParamsDateRange = "28d"
	RadarAs112SummaryIPVersionListParamsDateRange12w        RadarAs112SummaryIPVersionListParamsDateRange = "12w"
	RadarAs112SummaryIPVersionListParamsDateRange24w        RadarAs112SummaryIPVersionListParamsDateRange = "24w"
	RadarAs112SummaryIPVersionListParamsDateRange52w        RadarAs112SummaryIPVersionListParamsDateRange = "52w"
	RadarAs112SummaryIPVersionListParamsDateRange1dControl  RadarAs112SummaryIPVersionListParamsDateRange = "1dControl"
	RadarAs112SummaryIPVersionListParamsDateRange7dControl  RadarAs112SummaryIPVersionListParamsDateRange = "7dControl"
	RadarAs112SummaryIPVersionListParamsDateRange14dControl RadarAs112SummaryIPVersionListParamsDateRange = "14dControl"
	RadarAs112SummaryIPVersionListParamsDateRange28dControl RadarAs112SummaryIPVersionListParamsDateRange = "28dControl"
	RadarAs112SummaryIPVersionListParamsDateRange12wControl RadarAs112SummaryIPVersionListParamsDateRange = "12wControl"
	RadarAs112SummaryIPVersionListParamsDateRange24wControl RadarAs112SummaryIPVersionListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112SummaryIPVersionListParamsFormat string

const (
	RadarAs112SummaryIPVersionListParamsFormatJson RadarAs112SummaryIPVersionListParamsFormat = "JSON"
	RadarAs112SummaryIPVersionListParamsFormatCsv  RadarAs112SummaryIPVersionListParamsFormat = "CSV"
)
