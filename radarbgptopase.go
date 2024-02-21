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
func (r *RadarBGPTopAseService) Get(ctx context.Context, query RadarBGPTopAseGetParams, opts ...option.RequestOption) (res *RadarBGPTopAseGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarBGPTopAseGetResponseEnvelope
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

type RadarBGPTopAseGetResponse struct {
	Meta RadarBGPTopAseGetResponseMeta   `json:"meta,required"`
	Top0 []RadarBGPTopAseGetResponseTop0 `json:"top_0,required"`
	JSON radarBGPTopAseGetResponseJSON   `json:"-"`
}

// radarBGPTopAseGetResponseJSON contains the JSON metadata for the struct
// [RadarBGPTopAseGetResponse]
type radarBGPTopAseGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopAseGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPTopAseGetResponseMeta struct {
	DateRange []RadarBGPTopAseGetResponseMetaDateRange `json:"dateRange,required"`
	JSON      radarBGPTopAseGetResponseMetaJSON        `json:"-"`
}

// radarBGPTopAseGetResponseMetaJSON contains the JSON metadata for the struct
// [RadarBGPTopAseGetResponseMeta]
type radarBGPTopAseGetResponseMetaJSON struct {
	DateRange   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopAseGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPTopAseGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                  `json:"startTime,required" format:"date-time"`
	JSON      radarBGPTopAseGetResponseMetaDateRangeJSON `json:"-"`
}

// radarBGPTopAseGetResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [RadarBGPTopAseGetResponseMetaDateRange]
type radarBGPTopAseGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopAseGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPTopAseGetResponseTop0 struct {
	Asn    int64  `json:"asn,required"`
	AsName string `json:"ASName,required"`
	// Percentage of updates by this AS out of the total updates by all autonomous
	// systems.
	Value string                            `json:"value,required"`
	JSON  radarBGPTopAseGetResponseTop0JSON `json:"-"`
}

// radarBGPTopAseGetResponseTop0JSON contains the JSON metadata for the struct
// [RadarBGPTopAseGetResponseTop0]
type radarBGPTopAseGetResponseTop0JSON struct {
	Asn         apijson.Field
	AsName      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopAseGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
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

type RadarBGPTopAseGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarBGPTopAseGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarBGPTopAseGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of BGP network prefixes.
	Prefix param.Field[[]string] `query:"prefix"`
	// Array of BGP update types.
	UpdateType param.Field[[]RadarBGPTopAseGetParamsUpdateType] `query:"updateType"`
}

// URLQuery serializes [RadarBGPTopAseGetParams]'s query parameters as
// `url.Values`.
func (r RadarBGPTopAseGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarBGPTopAseGetParamsDateRange string

const (
	RadarBGPTopAseGetParamsDateRange1d         RadarBGPTopAseGetParamsDateRange = "1d"
	RadarBGPTopAseGetParamsDateRange2d         RadarBGPTopAseGetParamsDateRange = "2d"
	RadarBGPTopAseGetParamsDateRange7d         RadarBGPTopAseGetParamsDateRange = "7d"
	RadarBGPTopAseGetParamsDateRange14d        RadarBGPTopAseGetParamsDateRange = "14d"
	RadarBGPTopAseGetParamsDateRange28d        RadarBGPTopAseGetParamsDateRange = "28d"
	RadarBGPTopAseGetParamsDateRange12w        RadarBGPTopAseGetParamsDateRange = "12w"
	RadarBGPTopAseGetParamsDateRange24w        RadarBGPTopAseGetParamsDateRange = "24w"
	RadarBGPTopAseGetParamsDateRange52w        RadarBGPTopAseGetParamsDateRange = "52w"
	RadarBGPTopAseGetParamsDateRange1dControl  RadarBGPTopAseGetParamsDateRange = "1dControl"
	RadarBGPTopAseGetParamsDateRange2dControl  RadarBGPTopAseGetParamsDateRange = "2dControl"
	RadarBGPTopAseGetParamsDateRange7dControl  RadarBGPTopAseGetParamsDateRange = "7dControl"
	RadarBGPTopAseGetParamsDateRange14dControl RadarBGPTopAseGetParamsDateRange = "14dControl"
	RadarBGPTopAseGetParamsDateRange28dControl RadarBGPTopAseGetParamsDateRange = "28dControl"
	RadarBGPTopAseGetParamsDateRange12wControl RadarBGPTopAseGetParamsDateRange = "12wControl"
	RadarBGPTopAseGetParamsDateRange24wControl RadarBGPTopAseGetParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarBGPTopAseGetParamsFormat string

const (
	RadarBGPTopAseGetParamsFormatJson RadarBGPTopAseGetParamsFormat = "JSON"
	RadarBGPTopAseGetParamsFormatCsv  RadarBGPTopAseGetParamsFormat = "CSV"
)

type RadarBGPTopAseGetParamsUpdateType string

const (
	RadarBGPTopAseGetParamsUpdateTypeAnnouncement RadarBGPTopAseGetParamsUpdateType = "ANNOUNCEMENT"
	RadarBGPTopAseGetParamsUpdateTypeWithdrawal   RadarBGPTopAseGetParamsUpdateType = "WITHDRAWAL"
)

type RadarBGPTopAseGetResponseEnvelope struct {
	Result  RadarBGPTopAseGetResponse             `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    radarBGPTopAseGetResponseEnvelopeJSON `json:"-"`
}

// radarBGPTopAseGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RadarBGPTopAseGetResponseEnvelope]
type radarBGPTopAseGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopAseGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
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
