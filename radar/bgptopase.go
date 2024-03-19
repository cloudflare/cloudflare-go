// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// BGPTopAseService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewBGPTopAseService] method instead.
type BGPTopAseService struct {
	Options []option.RequestOption
}

// NewBGPTopAseService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBGPTopAseService(opts ...option.RequestOption) (r *BGPTopAseService) {
	r = &BGPTopAseService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS) by BGP updates (announcements only). Values
// are a percentage out of the total updates.
func (r *BGPTopAseService) Get(ctx context.Context, query BGPTopAseGetParams, opts ...option.RequestOption) (res *BGPTopAseGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env BGPTopAseGetResponseEnvelope
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
func (r *BGPTopAseService) Prefixes(ctx context.Context, query BGPTopAsePrefixesParams, opts ...option.RequestOption) (res *BGPTopAsePrefixesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env BGPTopAsePrefixesResponseEnvelope
	path := "radar/bgp/top/ases/prefixes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BGPTopAseGetResponse struct {
	Meta BGPTopAseGetResponseMeta   `json:"meta,required"`
	Top0 []BGPTopAseGetResponseTop0 `json:"top_0,required"`
	JSON bgpTopAseGetResponseJSON   `json:"-"`
}

// bgpTopAseGetResponseJSON contains the JSON metadata for the struct
// [BGPTopAseGetResponse]
type bgpTopAseGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPTopAseGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpTopAseGetResponseJSON) RawJSON() string {
	return r.raw
}

type BGPTopAseGetResponseMeta struct {
	DateRange []BGPTopAseGetResponseMetaDateRange `json:"dateRange,required"`
	JSON      bgpTopAseGetResponseMetaJSON        `json:"-"`
}

// bgpTopAseGetResponseMetaJSON contains the JSON metadata for the struct
// [BGPTopAseGetResponseMeta]
type bgpTopAseGetResponseMetaJSON struct {
	DateRange   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPTopAseGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpTopAseGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type BGPTopAseGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                             `json:"startTime,required" format:"date-time"`
	JSON      bgpTopAseGetResponseMetaDateRangeJSON `json:"-"`
}

// bgpTopAseGetResponseMetaDateRangeJSON contains the JSON metadata for the struct
// [BGPTopAseGetResponseMetaDateRange]
type bgpTopAseGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPTopAseGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpTopAseGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type BGPTopAseGetResponseTop0 struct {
	ASN    int64  `json:"asn,required"`
	AsName string `json:"ASName,required"`
	// Percentage of updates by this AS out of the total updates by all autonomous
	// systems.
	Value string                       `json:"value,required"`
	JSON  bgpTopAseGetResponseTop0JSON `json:"-"`
}

// bgpTopAseGetResponseTop0JSON contains the JSON metadata for the struct
// [BGPTopAseGetResponseTop0]
type bgpTopAseGetResponseTop0JSON struct {
	ASN         apijson.Field
	AsName      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPTopAseGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpTopAseGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type BGPTopAsePrefixesResponse struct {
	ASNs []BGPTopAsePrefixesResponseASN `json:"asns,required"`
	Meta BGPTopAsePrefixesResponseMeta  `json:"meta,required"`
	JSON bgpTopAsePrefixesResponseJSON  `json:"-"`
}

// bgpTopAsePrefixesResponseJSON contains the JSON metadata for the struct
// [BGPTopAsePrefixesResponse]
type bgpTopAsePrefixesResponseJSON struct {
	ASNs        apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPTopAsePrefixesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpTopAsePrefixesResponseJSON) RawJSON() string {
	return r.raw
}

type BGPTopAsePrefixesResponseASN struct {
	ASN       int64                            `json:"asn,required"`
	Country   string                           `json:"country,required"`
	Name      string                           `json:"name,required"`
	PfxsCount int64                            `json:"pfxs_count,required"`
	JSON      bgpTopAsePrefixesResponseASNJSON `json:"-"`
}

// bgpTopAsePrefixesResponseASNJSON contains the JSON metadata for the struct
// [BGPTopAsePrefixesResponseASN]
type bgpTopAsePrefixesResponseASNJSON struct {
	ASN         apijson.Field
	Country     apijson.Field
	Name        apijson.Field
	PfxsCount   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPTopAsePrefixesResponseASN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpTopAsePrefixesResponseASNJSON) RawJSON() string {
	return r.raw
}

type BGPTopAsePrefixesResponseMeta struct {
	DataTime   string                            `json:"data_time,required"`
	QueryTime  string                            `json:"query_time,required"`
	TotalPeers int64                             `json:"total_peers,required"`
	JSON       bgpTopAsePrefixesResponseMetaJSON `json:"-"`
}

// bgpTopAsePrefixesResponseMetaJSON contains the JSON metadata for the struct
// [BGPTopAsePrefixesResponseMeta]
type bgpTopAsePrefixesResponseMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalPeers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPTopAsePrefixesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpTopAsePrefixesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type BGPTopAseGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]BGPTopAseGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[BGPTopAseGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of BGP network prefixes.
	Prefix param.Field[[]string] `query:"prefix"`
	// Array of BGP update types.
	UpdateType param.Field[[]BGPTopAseGetParamsUpdateType] `query:"updateType"`
}

// URLQuery serializes [BGPTopAseGetParams]'s query parameters as `url.Values`.
func (r BGPTopAseGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BGPTopAseGetParamsDateRange string

const (
	BGPTopAseGetParamsDateRange1d         BGPTopAseGetParamsDateRange = "1d"
	BGPTopAseGetParamsDateRange2d         BGPTopAseGetParamsDateRange = "2d"
	BGPTopAseGetParamsDateRange7d         BGPTopAseGetParamsDateRange = "7d"
	BGPTopAseGetParamsDateRange14d        BGPTopAseGetParamsDateRange = "14d"
	BGPTopAseGetParamsDateRange28d        BGPTopAseGetParamsDateRange = "28d"
	BGPTopAseGetParamsDateRange12w        BGPTopAseGetParamsDateRange = "12w"
	BGPTopAseGetParamsDateRange24w        BGPTopAseGetParamsDateRange = "24w"
	BGPTopAseGetParamsDateRange52w        BGPTopAseGetParamsDateRange = "52w"
	BGPTopAseGetParamsDateRange1dControl  BGPTopAseGetParamsDateRange = "1dControl"
	BGPTopAseGetParamsDateRange2dControl  BGPTopAseGetParamsDateRange = "2dControl"
	BGPTopAseGetParamsDateRange7dControl  BGPTopAseGetParamsDateRange = "7dControl"
	BGPTopAseGetParamsDateRange14dControl BGPTopAseGetParamsDateRange = "14dControl"
	BGPTopAseGetParamsDateRange28dControl BGPTopAseGetParamsDateRange = "28dControl"
	BGPTopAseGetParamsDateRange12wControl BGPTopAseGetParamsDateRange = "12wControl"
	BGPTopAseGetParamsDateRange24wControl BGPTopAseGetParamsDateRange = "24wControl"
)

func (r BGPTopAseGetParamsDateRange) IsKnown() bool {
	switch r {
	case BGPTopAseGetParamsDateRange1d, BGPTopAseGetParamsDateRange2d, BGPTopAseGetParamsDateRange7d, BGPTopAseGetParamsDateRange14d, BGPTopAseGetParamsDateRange28d, BGPTopAseGetParamsDateRange12w, BGPTopAseGetParamsDateRange24w, BGPTopAseGetParamsDateRange52w, BGPTopAseGetParamsDateRange1dControl, BGPTopAseGetParamsDateRange2dControl, BGPTopAseGetParamsDateRange7dControl, BGPTopAseGetParamsDateRange14dControl, BGPTopAseGetParamsDateRange28dControl, BGPTopAseGetParamsDateRange12wControl, BGPTopAseGetParamsDateRange24wControl:
		return true
	}
	return false
}

// Format results are returned in.
type BGPTopAseGetParamsFormat string

const (
	BGPTopAseGetParamsFormatJson BGPTopAseGetParamsFormat = "JSON"
	BGPTopAseGetParamsFormatCsv  BGPTopAseGetParamsFormat = "CSV"
)

func (r BGPTopAseGetParamsFormat) IsKnown() bool {
	switch r {
	case BGPTopAseGetParamsFormatJson, BGPTopAseGetParamsFormatCsv:
		return true
	}
	return false
}

type BGPTopAseGetParamsUpdateType string

const (
	BGPTopAseGetParamsUpdateTypeAnnouncement BGPTopAseGetParamsUpdateType = "ANNOUNCEMENT"
	BGPTopAseGetParamsUpdateTypeWithdrawal   BGPTopAseGetParamsUpdateType = "WITHDRAWAL"
)

func (r BGPTopAseGetParamsUpdateType) IsKnown() bool {
	switch r {
	case BGPTopAseGetParamsUpdateTypeAnnouncement, BGPTopAseGetParamsUpdateTypeWithdrawal:
		return true
	}
	return false
}

type BGPTopAseGetResponseEnvelope struct {
	Result  BGPTopAseGetResponse             `json:"result,required"`
	Success bool                             `json:"success,required"`
	JSON    bgpTopAseGetResponseEnvelopeJSON `json:"-"`
}

// bgpTopAseGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [BGPTopAseGetResponseEnvelope]
type bgpTopAseGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPTopAseGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpTopAseGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BGPTopAsePrefixesParams struct {
	// Alpha-2 country code.
	Country param.Field[string] `query:"country"`
	// Format results are returned in.
	Format param.Field[BGPTopAsePrefixesParamsFormat] `query:"format"`
	// Maximum number of ASes to return
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [BGPTopAsePrefixesParams]'s query parameters as
// `url.Values`.
func (r BGPTopAsePrefixesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type BGPTopAsePrefixesParamsFormat string

const (
	BGPTopAsePrefixesParamsFormatJson BGPTopAsePrefixesParamsFormat = "JSON"
	BGPTopAsePrefixesParamsFormatCsv  BGPTopAsePrefixesParamsFormat = "CSV"
)

func (r BGPTopAsePrefixesParamsFormat) IsKnown() bool {
	switch r {
	case BGPTopAsePrefixesParamsFormatJson, BGPTopAsePrefixesParamsFormatCsv:
		return true
	}
	return false
}

type BGPTopAsePrefixesResponseEnvelope struct {
	Result  BGPTopAsePrefixesResponse             `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    bgpTopAsePrefixesResponseEnvelopeJSON `json:"-"`
}

// bgpTopAsePrefixesResponseEnvelopeJSON contains the JSON metadata for the struct
// [BGPTopAsePrefixesResponseEnvelope]
type bgpTopAsePrefixesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPTopAsePrefixesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpTopAsePrefixesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
