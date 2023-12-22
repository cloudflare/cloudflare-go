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

// RadarBgpTopPrefixService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarBgpTopPrefixService] method
// instead.
type RadarBgpTopPrefixService struct {
	Options []option.RequestOption
}

// NewRadarBgpTopPrefixService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarBgpTopPrefixService(opts ...option.RequestOption) (r *RadarBgpTopPrefixService) {
	r = &RadarBgpTopPrefixService{}
	r.Options = opts
	return
}

// Get the top network prefixes by BGP updates. Values are a percentage out of the
// total BGP updates.
func (r *RadarBgpTopPrefixService) List(ctx context.Context, query RadarBgpTopPrefixListParams, opts ...option.RequestOption) (res *RadarBgpTopPrefixListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/bgp/top/prefixes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarBgpTopPrefixListResponse struct {
	Result  RadarBgpTopPrefixListResponseResult `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    radarBgpTopPrefixListResponseJSON   `json:"-"`
}

// radarBgpTopPrefixListResponseJSON contains the JSON metadata for the struct
// [RadarBgpTopPrefixListResponse]
type radarBgpTopPrefixListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTopPrefixListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpTopPrefixListResponseResult struct {
	Meta RadarBgpTopPrefixListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarBgpTopPrefixListResponseResultTop0 `json:"top_0,required"`
	JSON radarBgpTopPrefixListResponseResultJSON   `json:"-"`
}

// radarBgpTopPrefixListResponseResultJSON contains the JSON metadata for the
// struct [RadarBgpTopPrefixListResponseResult]
type radarBgpTopPrefixListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTopPrefixListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpTopPrefixListResponseResultMeta struct {
	DateRange RadarBgpTopPrefixListResponseResultMetaDateRange `json:"dateRange,required"`
	JSON      radarBgpTopPrefixListResponseResultMetaJSON      `json:"-"`
}

// radarBgpTopPrefixListResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarBgpTopPrefixListResponseResultMeta]
type radarBgpTopPrefixListResponseResultMetaJSON struct {
	DateRange   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTopPrefixListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpTopPrefixListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                            `json:"startTime,required" format:"date-time"`
	JSON      radarBgpTopPrefixListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarBgpTopPrefixListResponseResultMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarBgpTopPrefixListResponseResultMetaDateRange]
type radarBgpTopPrefixListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTopPrefixListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpTopPrefixListResponseResultTop0 struct {
	Prefix string                                      `json:"prefix,required"`
	Value  string                                      `json:"value,required"`
	JSON   radarBgpTopPrefixListResponseResultTop0JSON `json:"-"`
}

// radarBgpTopPrefixListResponseResultTop0JSON contains the JSON metadata for the
// struct [RadarBgpTopPrefixListResponseResultTop0]
type radarBgpTopPrefixListResponseResultTop0JSON struct {
	Prefix      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTopPrefixListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpTopPrefixListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarBgpTopPrefixListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarBgpTopPrefixListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of BGP update types.
	UpdateType param.Field[[]RadarBgpTopPrefixListParamsUpdateType] `query:"updateType"`
}

// URLQuery serializes [RadarBgpTopPrefixListParams]'s query parameters as
// `url.Values`.
func (r RadarBgpTopPrefixListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarBgpTopPrefixListParamsDateRange string

const (
	RadarBgpTopPrefixListParamsDateRange1d         RadarBgpTopPrefixListParamsDateRange = "1d"
	RadarBgpTopPrefixListParamsDateRange7d         RadarBgpTopPrefixListParamsDateRange = "7d"
	RadarBgpTopPrefixListParamsDateRange14d        RadarBgpTopPrefixListParamsDateRange = "14d"
	RadarBgpTopPrefixListParamsDateRange28d        RadarBgpTopPrefixListParamsDateRange = "28d"
	RadarBgpTopPrefixListParamsDateRange12w        RadarBgpTopPrefixListParamsDateRange = "12w"
	RadarBgpTopPrefixListParamsDateRange24w        RadarBgpTopPrefixListParamsDateRange = "24w"
	RadarBgpTopPrefixListParamsDateRange52w        RadarBgpTopPrefixListParamsDateRange = "52w"
	RadarBgpTopPrefixListParamsDateRange1dControl  RadarBgpTopPrefixListParamsDateRange = "1dControl"
	RadarBgpTopPrefixListParamsDateRange7dControl  RadarBgpTopPrefixListParamsDateRange = "7dControl"
	RadarBgpTopPrefixListParamsDateRange14dControl RadarBgpTopPrefixListParamsDateRange = "14dControl"
	RadarBgpTopPrefixListParamsDateRange28dControl RadarBgpTopPrefixListParamsDateRange = "28dControl"
	RadarBgpTopPrefixListParamsDateRange12wControl RadarBgpTopPrefixListParamsDateRange = "12wControl"
	RadarBgpTopPrefixListParamsDateRange24wControl RadarBgpTopPrefixListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarBgpTopPrefixListParamsFormat string

const (
	RadarBgpTopPrefixListParamsFormatJson RadarBgpTopPrefixListParamsFormat = "JSON"
	RadarBgpTopPrefixListParamsFormatCsv  RadarBgpTopPrefixListParamsFormat = "CSV"
)

type RadarBgpTopPrefixListParamsUpdateType string

const (
	RadarBgpTopPrefixListParamsUpdateTypeAnnouncement RadarBgpTopPrefixListParamsUpdateType = "ANNOUNCEMENT"
	RadarBgpTopPrefixListParamsUpdateTypeWithdrawal   RadarBgpTopPrefixListParamsUpdateType = "WITHDRAWAL"
)
