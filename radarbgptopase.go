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

// RadarBGPTopAseService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarBGPTopAseService] method
// instead.
type RadarBGPTopAseService struct {
	Options []option.RequestOption
}

// NewRadarBGPTopAseService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarBGPTopAseService(opts ...option.RequestOption) (r *RadarBGPTopAseService) {
	r = &RadarBGPTopAseService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS) by BGP updates (announcements only). Values
// are a percentage out of the total updates.
func (r *RadarBGPTopAseService) List(ctx context.Context, query RadarBGPTopAseListParams, opts ...option.RequestOption) (res *RadarBGPTopAseListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarBGPTopAseListResponseEnvelope
	path := "radar/bgp/top/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the full list of autonomous systems on the global routing table ordered by
// announced prefixes count. The data comes from public BGP MRT data archives and
// updates every 2 hours.
func (r *RadarBGPTopAseService) Prefixes(ctx context.Context, query RadarBGPTopAsePrefixesParams, opts ...option.RequestOption) (res *RadarBGPTopAsePrefixesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarBGPTopAsePrefixesResponseEnvelope
	path := "radar/bgp/top/ases/prefixes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarBGPTopAseListResponse struct {
	Meta RadarBGPTopAseListResponseMeta   `json:"meta,required"`
	Top0 []RadarBGPTopAseListResponseTop0 `json:"top_0,required"`
	JSON radarBGPTopAseListResponseJSON   `json:"-"`
}

// radarBGPTopAseListResponseJSON contains the JSON metadata for the struct
// [RadarBGPTopAseListResponse]
type radarBGPTopAseListResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopAseListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPTopAseListResponseMeta struct {
	DateRange []RadarBGPTopAseListResponseMetaDateRange `json:"dateRange,required"`
	JSON      radarBGPTopAseListResponseMetaJSON        `json:"-"`
}

// radarBGPTopAseListResponseMetaJSON contains the JSON metadata for the struct
// [RadarBGPTopAseListResponseMeta]
type radarBGPTopAseListResponseMetaJSON struct {
	DateRange   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopAseListResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPTopAseListResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                   `json:"startTime,required" format:"date-time"`
	JSON      radarBGPTopAseListResponseMetaDateRangeJSON `json:"-"`
}

// radarBGPTopAseListResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [RadarBGPTopAseListResponseMetaDateRange]
type radarBGPTopAseListResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopAseListResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPTopAseListResponseTop0 struct {
	Asn    int64  `json:"asn,required"`
	AsName string `json:"ASName,required"`
	// Percentage of updates by this AS out of the total updates by all autonomous
	// systems.
	Value string                             `json:"value,required"`
	JSON  radarBGPTopAseListResponseTop0JSON `json:"-"`
}

// radarBGPTopAseListResponseTop0JSON contains the JSON metadata for the struct
// [RadarBGPTopAseListResponseTop0]
type radarBGPTopAseListResponseTop0JSON struct {
	Asn         apijson.Field
	AsName      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopAseListResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPTopAsePrefixesResponse struct {
	Asns []RadarBGPTopAsePrefixesResponseAsn `json:"asns,required"`
	Meta RadarBGPTopAsePrefixesResponseMeta  `json:"meta,required"`
	JSON radarBGPTopAsePrefixesResponseJSON  `json:"-"`
}

// radarBGPTopAsePrefixesResponseJSON contains the JSON metadata for the struct
// [RadarBGPTopAsePrefixesResponse]
type radarBGPTopAsePrefixesResponseJSON struct {
	Asns        apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopAsePrefixesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPTopAsePrefixesResponseAsn struct {
	Asn       int64                                 `json:"asn,required"`
	Country   string                                `json:"country,required"`
	Name      string                                `json:"name,required"`
	PfxsCount int64                                 `json:"pfxs_count,required"`
	JSON      radarBGPTopAsePrefixesResponseAsnJSON `json:"-"`
}

// radarBGPTopAsePrefixesResponseAsnJSON contains the JSON metadata for the struct
// [RadarBGPTopAsePrefixesResponseAsn]
type radarBGPTopAsePrefixesResponseAsnJSON struct {
	Asn         apijson.Field
	Country     apijson.Field
	Name        apijson.Field
	PfxsCount   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopAsePrefixesResponseAsn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPTopAsePrefixesResponseMeta struct {
	DataTime   string                                 `json:"data_time,required"`
	QueryTime  string                                 `json:"query_time,required"`
	TotalPeers int64                                  `json:"total_peers,required"`
	JSON       radarBGPTopAsePrefixesResponseMetaJSON `json:"-"`
}

// radarBGPTopAsePrefixesResponseMetaJSON contains the JSON metadata for the struct
// [RadarBGPTopAsePrefixesResponseMeta]
type radarBGPTopAsePrefixesResponseMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalPeers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopAsePrefixesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPTopAseListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarBGPTopAseListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarBGPTopAseListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of BGP network prefixes.
	Prefix param.Field[[]string] `query:"prefix"`
	// Array of BGP update types.
	UpdateType param.Field[[]RadarBGPTopAseListParamsUpdateType] `query:"updateType"`
}

// URLQuery serializes [RadarBGPTopAseListParams]'s query parameters as
// `url.Values`.
func (r RadarBGPTopAseListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarBGPTopAseListParamsDateRange string

const (
	RadarBGPTopAseListParamsDateRange1d         RadarBGPTopAseListParamsDateRange = "1d"
	RadarBGPTopAseListParamsDateRange2d         RadarBGPTopAseListParamsDateRange = "2d"
	RadarBGPTopAseListParamsDateRange7d         RadarBGPTopAseListParamsDateRange = "7d"
	RadarBGPTopAseListParamsDateRange14d        RadarBGPTopAseListParamsDateRange = "14d"
	RadarBGPTopAseListParamsDateRange28d        RadarBGPTopAseListParamsDateRange = "28d"
	RadarBGPTopAseListParamsDateRange12w        RadarBGPTopAseListParamsDateRange = "12w"
	RadarBGPTopAseListParamsDateRange24w        RadarBGPTopAseListParamsDateRange = "24w"
	RadarBGPTopAseListParamsDateRange52w        RadarBGPTopAseListParamsDateRange = "52w"
	RadarBGPTopAseListParamsDateRange1dControl  RadarBGPTopAseListParamsDateRange = "1dControl"
	RadarBGPTopAseListParamsDateRange2dControl  RadarBGPTopAseListParamsDateRange = "2dControl"
	RadarBGPTopAseListParamsDateRange7dControl  RadarBGPTopAseListParamsDateRange = "7dControl"
	RadarBGPTopAseListParamsDateRange14dControl RadarBGPTopAseListParamsDateRange = "14dControl"
	RadarBGPTopAseListParamsDateRange28dControl RadarBGPTopAseListParamsDateRange = "28dControl"
	RadarBGPTopAseListParamsDateRange12wControl RadarBGPTopAseListParamsDateRange = "12wControl"
	RadarBGPTopAseListParamsDateRange24wControl RadarBGPTopAseListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarBGPTopAseListParamsFormat string

const (
	RadarBGPTopAseListParamsFormatJson RadarBGPTopAseListParamsFormat = "JSON"
	RadarBGPTopAseListParamsFormatCsv  RadarBGPTopAseListParamsFormat = "CSV"
)

type RadarBGPTopAseListParamsUpdateType string

const (
	RadarBGPTopAseListParamsUpdateTypeAnnouncement RadarBGPTopAseListParamsUpdateType = "ANNOUNCEMENT"
	RadarBGPTopAseListParamsUpdateTypeWithdrawal   RadarBGPTopAseListParamsUpdateType = "WITHDRAWAL"
)

type RadarBGPTopAseListResponseEnvelope struct {
	Result  RadarBGPTopAseListResponse             `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    radarBGPTopAseListResponseEnvelopeJSON `json:"-"`
}

// radarBGPTopAseListResponseEnvelopeJSON contains the JSON metadata for the struct
// [RadarBGPTopAseListResponseEnvelope]
type radarBGPTopAseListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopAseListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPTopAsePrefixesParams struct {
	// Alpha-2 country code.
	Country param.Field[string] `query:"country"`
	// Format results are returned in.
	Format param.Field[RadarBGPTopAsePrefixesParamsFormat] `query:"format"`
	// Maximum number of ASes to return
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [RadarBGPTopAsePrefixesParams]'s query parameters as
// `url.Values`.
func (r RadarBGPTopAsePrefixesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarBGPTopAsePrefixesParamsFormat string

const (
	RadarBGPTopAsePrefixesParamsFormatJson RadarBGPTopAsePrefixesParamsFormat = "JSON"
	RadarBGPTopAsePrefixesParamsFormatCsv  RadarBGPTopAsePrefixesParamsFormat = "CSV"
)

type RadarBGPTopAsePrefixesResponseEnvelope struct {
	Result  RadarBGPTopAsePrefixesResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarBGPTopAsePrefixesResponseEnvelopeJSON `json:"-"`
}

// radarBGPTopAsePrefixesResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarBGPTopAsePrefixesResponseEnvelope]
type radarBGPTopAsePrefixesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopAsePrefixesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
