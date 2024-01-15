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

// RadarBgpTopAseService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarBgpTopAseService] method
// instead.
type RadarBgpTopAseService struct {
	Options  []option.RequestOption
	Prefixes *RadarBgpTopAsePrefixService
}

// NewRadarBgpTopAseService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarBgpTopAseService(opts ...option.RequestOption) (r *RadarBgpTopAseService) {
	r = &RadarBgpTopAseService{}
	r.Options = opts
	r.Prefixes = NewRadarBgpTopAsePrefixService(opts...)
	return
}

// Get the top autonomous systems (AS) by BGP updates (announcements only). Values
// are a percentage out of the total updates.
func (r *RadarBgpTopAseService) List(ctx context.Context, query RadarBgpTopAseListParams, opts ...option.RequestOption) (res *RadarBgpTopAseListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/bgp/top/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarBgpTopAseListResponse struct {
	Result  RadarBgpTopAseListResponseResult `json:"result,required"`
	Success bool                             `json:"success,required"`
	JSON    radarBgpTopAseListResponseJSON   `json:"-"`
}

// radarBgpTopAseListResponseJSON contains the JSON metadata for the struct
// [RadarBgpTopAseListResponse]
type radarBgpTopAseListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTopAseListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpTopAseListResponseResult struct {
	Meta RadarBgpTopAseListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarBgpTopAseListResponseResultTop0 `json:"top_0,required"`
	JSON radarBgpTopAseListResponseResultJSON   `json:"-"`
}

// radarBgpTopAseListResponseResultJSON contains the JSON metadata for the struct
// [RadarBgpTopAseListResponseResult]
type radarBgpTopAseListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTopAseListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpTopAseListResponseResultMeta struct {
	DateRange []RadarBgpTopAseListResponseResultMetaDateRange `json:"dateRange,required"`
	JSON      radarBgpTopAseListResponseResultMetaJSON        `json:"-"`
}

// radarBgpTopAseListResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarBgpTopAseListResponseResultMeta]
type radarBgpTopAseListResponseResultMetaJSON struct {
	DateRange   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTopAseListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpTopAseListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                         `json:"startTime,required" format:"date-time"`
	JSON      radarBgpTopAseListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarBgpTopAseListResponseResultMetaDateRangeJSON contains the JSON metadata for
// the struct [RadarBgpTopAseListResponseResultMetaDateRange]
type radarBgpTopAseListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTopAseListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpTopAseListResponseResultTop0 struct {
	ASN    int64  `json:"asn,required"`
	AsName string `json:"ASName,required"`
	// Percentage of updates by this AS out of the total updates by all autonomous
	// systems.
	Value string                                   `json:"value,required"`
	JSON  radarBgpTopAseListResponseResultTop0JSON `json:"-"`
}

// radarBgpTopAseListResponseResultTop0JSON contains the JSON metadata for the
// struct [RadarBgpTopAseListResponseResultTop0]
type radarBgpTopAseListResponseResultTop0JSON struct {
	ASN         apijson.Field
	AsName      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTopAseListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpTopAseListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarBgpTopAseListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarBgpTopAseListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of BGP network prefixes.
	Prefix param.Field[[]string] `query:"prefix"`
	// Array of BGP update types.
	UpdateType param.Field[[]RadarBgpTopAseListParamsUpdateType] `query:"updateType"`
}

// URLQuery serializes [RadarBgpTopAseListParams]'s query parameters as
// `url.Values`.
func (r RadarBgpTopAseListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarBgpTopAseListParamsDateRange string

const (
	RadarBgpTopAseListParamsDateRange1d         RadarBgpTopAseListParamsDateRange = "1d"
	RadarBgpTopAseListParamsDateRange2d         RadarBgpTopAseListParamsDateRange = "2d"
	RadarBgpTopAseListParamsDateRange7d         RadarBgpTopAseListParamsDateRange = "7d"
	RadarBgpTopAseListParamsDateRange14d        RadarBgpTopAseListParamsDateRange = "14d"
	RadarBgpTopAseListParamsDateRange28d        RadarBgpTopAseListParamsDateRange = "28d"
	RadarBgpTopAseListParamsDateRange12w        RadarBgpTopAseListParamsDateRange = "12w"
	RadarBgpTopAseListParamsDateRange24w        RadarBgpTopAseListParamsDateRange = "24w"
	RadarBgpTopAseListParamsDateRange52w        RadarBgpTopAseListParamsDateRange = "52w"
	RadarBgpTopAseListParamsDateRange1dControl  RadarBgpTopAseListParamsDateRange = "1dControl"
	RadarBgpTopAseListParamsDateRange2dControl  RadarBgpTopAseListParamsDateRange = "2dControl"
	RadarBgpTopAseListParamsDateRange7dControl  RadarBgpTopAseListParamsDateRange = "7dControl"
	RadarBgpTopAseListParamsDateRange14dControl RadarBgpTopAseListParamsDateRange = "14dControl"
	RadarBgpTopAseListParamsDateRange28dControl RadarBgpTopAseListParamsDateRange = "28dControl"
	RadarBgpTopAseListParamsDateRange12wControl RadarBgpTopAseListParamsDateRange = "12wControl"
	RadarBgpTopAseListParamsDateRange24wControl RadarBgpTopAseListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarBgpTopAseListParamsFormat string

const (
	RadarBgpTopAseListParamsFormatJson RadarBgpTopAseListParamsFormat = "JSON"
	RadarBgpTopAseListParamsFormatCsv  RadarBgpTopAseListParamsFormat = "CSV"
)

type RadarBgpTopAseListParamsUpdateType string

const (
	RadarBgpTopAseListParamsUpdateTypeAnnouncement RadarBgpTopAseListParamsUpdateType = "ANNOUNCEMENT"
	RadarBgpTopAseListParamsUpdateTypeWithdrawal   RadarBgpTopAseListParamsUpdateType = "WITHDRAWAL"
)
