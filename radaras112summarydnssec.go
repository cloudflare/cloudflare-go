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

// RadarAs112SummaryDnssecService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarAs112SummaryDnssecService] method instead.
type RadarAs112SummaryDnssecService struct {
	Options []option.RequestOption
}

// NewRadarAs112SummaryDnssecService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAs112SummaryDnssecService(opts ...option.RequestOption) (r *RadarAs112SummaryDnssecService) {
	r = &RadarAs112SummaryDnssecService{}
	r.Options = opts
	return
}

// Percentage distribution of dns requests classified per DNSSEC.
func (r *RadarAs112SummaryDnssecService) List(ctx context.Context, query RadarAs112SummaryDnssecListParams, opts ...option.RequestOption) (res *RadarAs112SummaryDnssecListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/summary/dnssec"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAs112SummaryDnssecListResponse struct {
	Result  RadarAs112SummaryDnssecListResponseResult `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    radarAs112SummaryDnssecListResponseJSON   `json:"-"`
}

// radarAs112SummaryDnssecListResponseJSON contains the JSON metadata for the
// struct [RadarAs112SummaryDnssecListResponse]
type radarAs112SummaryDnssecListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryDnssecListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryDnssecListResponseResult struct {
	Meta     interface{}                                       `json:"meta,required"`
	Summary0 RadarAs112SummaryDnssecListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarAs112SummaryDnssecListResponseResultJSON     `json:"-"`
}

// radarAs112SummaryDnssecListResponseResultJSON contains the JSON metadata for the
// struct [RadarAs112SummaryDnssecListResponseResult]
type radarAs112SummaryDnssecListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryDnssecListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryDnssecListResponseResultSummary0 struct {
	NotSupported string                                                `json:"NOT_SUPPORTED,required"`
	Supported    string                                                `json:"SUPPORTED,required"`
	JSON         radarAs112SummaryDnssecListResponseResultSummary0JSON `json:"-"`
}

// radarAs112SummaryDnssecListResponseResultSummary0JSON contains the JSON metadata
// for the struct [RadarAs112SummaryDnssecListResponseResultSummary0]
type radarAs112SummaryDnssecListResponseResultSummary0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarAs112SummaryDnssecListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryDnssecListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112SummaryDnssecListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112SummaryDnssecListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112SummaryDnssecListParams]'s query parameters as
// `url.Values`.
func (r RadarAs112SummaryDnssecListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAs112SummaryDnssecListParamsDateRange string

const (
	RadarAs112SummaryDnssecListParamsDateRange1d         RadarAs112SummaryDnssecListParamsDateRange = "1d"
	RadarAs112SummaryDnssecListParamsDateRange7d         RadarAs112SummaryDnssecListParamsDateRange = "7d"
	RadarAs112SummaryDnssecListParamsDateRange14d        RadarAs112SummaryDnssecListParamsDateRange = "14d"
	RadarAs112SummaryDnssecListParamsDateRange28d        RadarAs112SummaryDnssecListParamsDateRange = "28d"
	RadarAs112SummaryDnssecListParamsDateRange12w        RadarAs112SummaryDnssecListParamsDateRange = "12w"
	RadarAs112SummaryDnssecListParamsDateRange24w        RadarAs112SummaryDnssecListParamsDateRange = "24w"
	RadarAs112SummaryDnssecListParamsDateRange52w        RadarAs112SummaryDnssecListParamsDateRange = "52w"
	RadarAs112SummaryDnssecListParamsDateRange1dControl  RadarAs112SummaryDnssecListParamsDateRange = "1dControl"
	RadarAs112SummaryDnssecListParamsDateRange7dControl  RadarAs112SummaryDnssecListParamsDateRange = "7dControl"
	RadarAs112SummaryDnssecListParamsDateRange14dControl RadarAs112SummaryDnssecListParamsDateRange = "14dControl"
	RadarAs112SummaryDnssecListParamsDateRange28dControl RadarAs112SummaryDnssecListParamsDateRange = "28dControl"
	RadarAs112SummaryDnssecListParamsDateRange12wControl RadarAs112SummaryDnssecListParamsDateRange = "12wControl"
	RadarAs112SummaryDnssecListParamsDateRange24wControl RadarAs112SummaryDnssecListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112SummaryDnssecListParamsFormat string

const (
	RadarAs112SummaryDnssecListParamsFormatJson RadarAs112SummaryDnssecListParamsFormat = "JSON"
	RadarAs112SummaryDnssecListParamsFormatCsv  RadarAs112SummaryDnssecListParamsFormat = "CSV"
)
