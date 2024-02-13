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

// RadarNetflowTopAseService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarNetflowTopAseService] method
// instead.
type RadarNetflowTopAseService struct {
	Options []option.RequestOption
}

// NewRadarNetflowTopAseService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarNetflowTopAseService(opts ...option.RequestOption) (r *RadarNetflowTopAseService) {
	r = &RadarNetflowTopAseService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS) by network traffic (NetFlows) over a given
// time period. Visit https://en.wikipedia.org/wiki/NetFlow for more information.
func (r *RadarNetflowTopAseService) List(ctx context.Context, query RadarNetflowTopAseListParams, opts ...option.RequestOption) (res *RadarNetflowTopAseListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarNetflowTopAseListResponseEnvelope
	path := "radar/netflows/top/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarNetflowTopAseListResponse struct {
	Top0 []RadarNetflowTopAseListResponseTop0 `json:"top_0,required"`
	JSON radarNetflowTopAseListResponseJSON   `json:"-"`
}

// radarNetflowTopAseListResponseJSON contains the JSON metadata for the struct
// [RadarNetflowTopAseListResponse]
type radarNetflowTopAseListResponseJSON struct {
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowTopAseListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarNetflowTopAseListResponseTop0 struct {
	ClientAsn    float64                                `json:"clientASN,required"`
	ClientAsName string                                 `json:"clientASName,required"`
	Value        string                                 `json:"value,required"`
	JSON         radarNetflowTopAseListResponseTop0JSON `json:"-"`
}

// radarNetflowTopAseListResponseTop0JSON contains the JSON metadata for the struct
// [RadarNetflowTopAseListResponseTop0]
type radarNetflowTopAseListResponseTop0JSON struct {
	ClientAsn    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarNetflowTopAseListResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarNetflowTopAseListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarNetflowTopAseListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarNetflowTopAseListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarNetflowTopAseListParams]'s query parameters as
// `url.Values`.
func (r RadarNetflowTopAseListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarNetflowTopAseListParamsDateRange string

const (
	RadarNetflowTopAseListParamsDateRange1d         RadarNetflowTopAseListParamsDateRange = "1d"
	RadarNetflowTopAseListParamsDateRange2d         RadarNetflowTopAseListParamsDateRange = "2d"
	RadarNetflowTopAseListParamsDateRange7d         RadarNetflowTopAseListParamsDateRange = "7d"
	RadarNetflowTopAseListParamsDateRange14d        RadarNetflowTopAseListParamsDateRange = "14d"
	RadarNetflowTopAseListParamsDateRange28d        RadarNetflowTopAseListParamsDateRange = "28d"
	RadarNetflowTopAseListParamsDateRange12w        RadarNetflowTopAseListParamsDateRange = "12w"
	RadarNetflowTopAseListParamsDateRange24w        RadarNetflowTopAseListParamsDateRange = "24w"
	RadarNetflowTopAseListParamsDateRange52w        RadarNetflowTopAseListParamsDateRange = "52w"
	RadarNetflowTopAseListParamsDateRange1dControl  RadarNetflowTopAseListParamsDateRange = "1dControl"
	RadarNetflowTopAseListParamsDateRange2dControl  RadarNetflowTopAseListParamsDateRange = "2dControl"
	RadarNetflowTopAseListParamsDateRange7dControl  RadarNetflowTopAseListParamsDateRange = "7dControl"
	RadarNetflowTopAseListParamsDateRange14dControl RadarNetflowTopAseListParamsDateRange = "14dControl"
	RadarNetflowTopAseListParamsDateRange28dControl RadarNetflowTopAseListParamsDateRange = "28dControl"
	RadarNetflowTopAseListParamsDateRange12wControl RadarNetflowTopAseListParamsDateRange = "12wControl"
	RadarNetflowTopAseListParamsDateRange24wControl RadarNetflowTopAseListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarNetflowTopAseListParamsFormat string

const (
	RadarNetflowTopAseListParamsFormatJson RadarNetflowTopAseListParamsFormat = "JSON"
	RadarNetflowTopAseListParamsFormatCsv  RadarNetflowTopAseListParamsFormat = "CSV"
)

type RadarNetflowTopAseListResponseEnvelope struct {
	Result  RadarNetflowTopAseListResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarNetflowTopAseListResponseEnvelopeJSON `json:"-"`
}

// radarNetflowTopAseListResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarNetflowTopAseListResponseEnvelope]
type radarNetflowTopAseListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowTopAseListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
