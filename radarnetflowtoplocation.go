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

// RadarNetflowTopLocationService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarNetflowTopLocationService] method instead.
type RadarNetflowTopLocationService struct {
	Options []option.RequestOption
}

// NewRadarNetflowTopLocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarNetflowTopLocationService(opts ...option.RequestOption) (r *RadarNetflowTopLocationService) {
	r = &RadarNetflowTopLocationService{}
	r.Options = opts
	return
}

// Get the top locations by network traffic (NetFlows) over a given time period.
// Visit https://en.wikipedia.org/wiki/NetFlow for more information.
func (r *RadarNetflowTopLocationService) List(ctx context.Context, query RadarNetflowTopLocationListParams, opts ...option.RequestOption) (res *RadarNetflowTopLocationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/netflows/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarNetflowTopLocationListResponse struct {
	Result  RadarNetflowTopLocationListResponseResult `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    radarNetflowTopLocationListResponseJSON   `json:"-"`
}

// radarNetflowTopLocationListResponseJSON contains the JSON metadata for the
// struct [RadarNetflowTopLocationListResponse]
type radarNetflowTopLocationListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowTopLocationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarNetflowTopLocationListResponseResult struct {
	Top0 []RadarNetflowTopLocationListResponseResultTop0 `json:"top_0,required"`
	JSON radarNetflowTopLocationListResponseResultJSON   `json:"-"`
}

// radarNetflowTopLocationListResponseResultJSON contains the JSON metadata for the
// struct [RadarNetflowTopLocationListResponseResult]
type radarNetflowTopLocationListResponseResultJSON struct {
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowTopLocationListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarNetflowTopLocationListResponseResultTop0 struct {
	ClientCountryAlpha2 string                                            `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                            `json:"clientCountryName,required"`
	Value               string                                            `json:"value,required"`
	JSON                radarNetflowTopLocationListResponseResultTop0JSON `json:"-"`
}

// radarNetflowTopLocationListResponseResultTop0JSON contains the JSON metadata for
// the struct [RadarNetflowTopLocationListResponseResultTop0]
type radarNetflowTopLocationListResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarNetflowTopLocationListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarNetflowTopLocationListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarNetflowTopLocationListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarNetflowTopLocationListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarNetflowTopLocationListParams]'s query parameters as
// `url.Values`.
func (r RadarNetflowTopLocationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarNetflowTopLocationListParamsDateRange string

const (
	RadarNetflowTopLocationListParamsDateRange1d         RadarNetflowTopLocationListParamsDateRange = "1d"
	RadarNetflowTopLocationListParamsDateRange2d         RadarNetflowTopLocationListParamsDateRange = "2d"
	RadarNetflowTopLocationListParamsDateRange7d         RadarNetflowTopLocationListParamsDateRange = "7d"
	RadarNetflowTopLocationListParamsDateRange14d        RadarNetflowTopLocationListParamsDateRange = "14d"
	RadarNetflowTopLocationListParamsDateRange28d        RadarNetflowTopLocationListParamsDateRange = "28d"
	RadarNetflowTopLocationListParamsDateRange12w        RadarNetflowTopLocationListParamsDateRange = "12w"
	RadarNetflowTopLocationListParamsDateRange24w        RadarNetflowTopLocationListParamsDateRange = "24w"
	RadarNetflowTopLocationListParamsDateRange52w        RadarNetflowTopLocationListParamsDateRange = "52w"
	RadarNetflowTopLocationListParamsDateRange1dControl  RadarNetflowTopLocationListParamsDateRange = "1dControl"
	RadarNetflowTopLocationListParamsDateRange2dControl  RadarNetflowTopLocationListParamsDateRange = "2dControl"
	RadarNetflowTopLocationListParamsDateRange7dControl  RadarNetflowTopLocationListParamsDateRange = "7dControl"
	RadarNetflowTopLocationListParamsDateRange14dControl RadarNetflowTopLocationListParamsDateRange = "14dControl"
	RadarNetflowTopLocationListParamsDateRange28dControl RadarNetflowTopLocationListParamsDateRange = "28dControl"
	RadarNetflowTopLocationListParamsDateRange12wControl RadarNetflowTopLocationListParamsDateRange = "12wControl"
	RadarNetflowTopLocationListParamsDateRange24wControl RadarNetflowTopLocationListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarNetflowTopLocationListParamsFormat string

const (
	RadarNetflowTopLocationListParamsFormatJson RadarNetflowTopLocationListParamsFormat = "JSON"
	RadarNetflowTopLocationListParamsFormatCsv  RadarNetflowTopLocationListParamsFormat = "CSV"
)
