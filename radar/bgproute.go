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

// BGPRouteService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewBGPRouteService] method instead.
type BGPRouteService struct {
	Options []option.RequestOption
}

// NewBGPRouteService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBGPRouteService(opts ...option.RequestOption) (r *BGPRouteService) {
	r = &BGPRouteService{}
	r.Options = opts
	return
}

// List all Multi-origin AS (MOAS) prefixes on the global routing tables.
func (r *BGPRouteService) Moas(ctx context.Context, query BGPRouteMoasParams, opts ...option.RequestOption) (res *BGPRouteMoasResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env BGPRouteMoasResponseEnvelope
	path := "radar/bgp/routes/moas"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lookup prefix-to-ASN mapping on global routing tables.
func (r *BGPRouteService) Pfx2as(ctx context.Context, query BGPRoutePfx2asParams, opts ...option.RequestOption) (res *BGPRoutePfx2asResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env BGPRoutePfx2asResponseEnvelope
	path := "radar/bgp/routes/pfx2as"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the BGP routing table stats (Beta).
func (r *BGPRouteService) Stats(ctx context.Context, query BGPRouteStatsParams, opts ...option.RequestOption) (res *BGPRouteStatsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env BGPRouteStatsResponseEnvelope
	path := "radar/bgp/routes/stats"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets time-series data for the announced IP space count, represented as the
// number of IPv4 /24s and IPv6 /48s, for a given ASN.
func (r *BGPRouteService) Timeseries(ctx context.Context, query BGPRouteTimeseriesParams, opts ...option.RequestOption) (res *BGPRouteTimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env BGPRouteTimeseriesResponseEnvelope
	path := "radar/bgp/routes/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BGPRouteMoasResponse struct {
	Meta BGPRouteMoasResponseMeta  `json:"meta,required"`
	Moas []BGPRouteMoasResponseMoa `json:"moas,required"`
	JSON bgpRouteMoasResponseJSON  `json:"-"`
}

// bgpRouteMoasResponseJSON contains the JSON metadata for the struct
// [BGPRouteMoasResponse]
type bgpRouteMoasResponseJSON struct {
	Meta        apijson.Field
	Moas        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPRouteMoasResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRouteMoasResponseJSON) RawJSON() string {
	return r.raw
}

type BGPRouteMoasResponseMeta struct {
	DataTime   string                       `json:"data_time,required"`
	QueryTime  string                       `json:"query_time,required"`
	TotalPeers int64                        `json:"total_peers,required"`
	JSON       bgpRouteMoasResponseMetaJSON `json:"-"`
}

// bgpRouteMoasResponseMetaJSON contains the JSON metadata for the struct
// [BGPRouteMoasResponseMeta]
type bgpRouteMoasResponseMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalPeers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPRouteMoasResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRouteMoasResponseMetaJSON) RawJSON() string {
	return r.raw
}

type BGPRouteMoasResponseMoa struct {
	Origins []BGPRouteMoasResponseMoasOrigin `json:"origins,required"`
	Prefix  string                           `json:"prefix,required"`
	JSON    bgpRouteMoasResponseMoaJSON      `json:"-"`
}

// bgpRouteMoasResponseMoaJSON contains the JSON metadata for the struct
// [BGPRouteMoasResponseMoa]
type bgpRouteMoasResponseMoaJSON struct {
	Origins     apijson.Field
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPRouteMoasResponseMoa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRouteMoasResponseMoaJSON) RawJSON() string {
	return r.raw
}

type BGPRouteMoasResponseMoasOrigin struct {
	Origin         int64                              `json:"origin,required"`
	PeerCount      int64                              `json:"peer_count,required"`
	RPKIValidation string                             `json:"rpki_validation,required"`
	JSON           bgpRouteMoasResponseMoasOriginJSON `json:"-"`
}

// bgpRouteMoasResponseMoasOriginJSON contains the JSON metadata for the struct
// [BGPRouteMoasResponseMoasOrigin]
type bgpRouteMoasResponseMoasOriginJSON struct {
	Origin         apijson.Field
	PeerCount      apijson.Field
	RPKIValidation apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BGPRouteMoasResponseMoasOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRouteMoasResponseMoasOriginJSON) RawJSON() string {
	return r.raw
}

type BGPRoutePfx2asResponse struct {
	Meta          BGPRoutePfx2asResponseMeta           `json:"meta,required"`
	PrefixOrigins []BGPRoutePfx2asResponsePrefixOrigin `json:"prefix_origins,required"`
	JSON          bgpRoutePfx2asResponseJSON           `json:"-"`
}

// bgpRoutePfx2asResponseJSON contains the JSON metadata for the struct
// [BGPRoutePfx2asResponse]
type bgpRoutePfx2asResponseJSON struct {
	Meta          apijson.Field
	PrefixOrigins apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *BGPRoutePfx2asResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRoutePfx2asResponseJSON) RawJSON() string {
	return r.raw
}

type BGPRoutePfx2asResponseMeta struct {
	DataTime   string                         `json:"data_time,required"`
	QueryTime  string                         `json:"query_time,required"`
	TotalPeers int64                          `json:"total_peers,required"`
	JSON       bgpRoutePfx2asResponseMetaJSON `json:"-"`
}

// bgpRoutePfx2asResponseMetaJSON contains the JSON metadata for the struct
// [BGPRoutePfx2asResponseMeta]
type bgpRoutePfx2asResponseMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalPeers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPRoutePfx2asResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRoutePfx2asResponseMetaJSON) RawJSON() string {
	return r.raw
}

type BGPRoutePfx2asResponsePrefixOrigin struct {
	Origin         int64                                  `json:"origin,required"`
	PeerCount      int64                                  `json:"peer_count,required"`
	Prefix         string                                 `json:"prefix,required"`
	RPKIValidation string                                 `json:"rpki_validation,required"`
	JSON           bgpRoutePfx2asResponsePrefixOriginJSON `json:"-"`
}

// bgpRoutePfx2asResponsePrefixOriginJSON contains the JSON metadata for the struct
// [BGPRoutePfx2asResponsePrefixOrigin]
type bgpRoutePfx2asResponsePrefixOriginJSON struct {
	Origin         apijson.Field
	PeerCount      apijson.Field
	Prefix         apijson.Field
	RPKIValidation apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BGPRoutePfx2asResponsePrefixOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRoutePfx2asResponsePrefixOriginJSON) RawJSON() string {
	return r.raw
}

type BGPRouteStatsResponse struct {
	Meta  BGPRouteStatsResponseMeta  `json:"meta,required"`
	Stats BGPRouteStatsResponseStats `json:"stats,required"`
	JSON  bgpRouteStatsResponseJSON  `json:"-"`
}

// bgpRouteStatsResponseJSON contains the JSON metadata for the struct
// [BGPRouteStatsResponse]
type bgpRouteStatsResponseJSON struct {
	Meta        apijson.Field
	Stats       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPRouteStatsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRouteStatsResponseJSON) RawJSON() string {
	return r.raw
}

type BGPRouteStatsResponseMeta struct {
	DataTime   string                        `json:"data_time,required"`
	QueryTime  string                        `json:"query_time,required"`
	TotalPeers int64                         `json:"total_peers,required"`
	JSON       bgpRouteStatsResponseMetaJSON `json:"-"`
}

// bgpRouteStatsResponseMetaJSON contains the JSON metadata for the struct
// [BGPRouteStatsResponseMeta]
type bgpRouteStatsResponseMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalPeers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPRouteStatsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRouteStatsResponseMetaJSON) RawJSON() string {
	return r.raw
}

type BGPRouteStatsResponseStats struct {
	DistinctOrigins      int64                          `json:"distinct_origins,required"`
	DistinctOriginsIPV4  int64                          `json:"distinct_origins_ipv4,required"`
	DistinctOriginsIPV6  int64                          `json:"distinct_origins_ipv6,required"`
	DistinctPrefixes     int64                          `json:"distinct_prefixes,required"`
	DistinctPrefixesIPV4 int64                          `json:"distinct_prefixes_ipv4,required"`
	DistinctPrefixesIPV6 int64                          `json:"distinct_prefixes_ipv6,required"`
	RoutesInvalid        int64                          `json:"routes_invalid,required"`
	RoutesInvalidIPV4    int64                          `json:"routes_invalid_ipv4,required"`
	RoutesInvalidIPV6    int64                          `json:"routes_invalid_ipv6,required"`
	RoutesTotal          int64                          `json:"routes_total,required"`
	RoutesTotalIPV4      int64                          `json:"routes_total_ipv4,required"`
	RoutesTotalIPV6      int64                          `json:"routes_total_ipv6,required"`
	RoutesUnknown        int64                          `json:"routes_unknown,required"`
	RoutesUnknownIPV4    int64                          `json:"routes_unknown_ipv4,required"`
	RoutesUnknownIPV6    int64                          `json:"routes_unknown_ipv6,required"`
	RoutesValid          int64                          `json:"routes_valid,required"`
	RoutesValidIPV4      int64                          `json:"routes_valid_ipv4,required"`
	RoutesValidIPV6      int64                          `json:"routes_valid_ipv6,required"`
	JSON                 bgpRouteStatsResponseStatsJSON `json:"-"`
}

// bgpRouteStatsResponseStatsJSON contains the JSON metadata for the struct
// [BGPRouteStatsResponseStats]
type bgpRouteStatsResponseStatsJSON struct {
	DistinctOrigins      apijson.Field
	DistinctOriginsIPV4  apijson.Field
	DistinctOriginsIPV6  apijson.Field
	DistinctPrefixes     apijson.Field
	DistinctPrefixesIPV4 apijson.Field
	DistinctPrefixesIPV6 apijson.Field
	RoutesInvalid        apijson.Field
	RoutesInvalidIPV4    apijson.Field
	RoutesInvalidIPV6    apijson.Field
	RoutesTotal          apijson.Field
	RoutesTotalIPV4      apijson.Field
	RoutesTotalIPV6      apijson.Field
	RoutesUnknown        apijson.Field
	RoutesUnknownIPV4    apijson.Field
	RoutesUnknownIPV6    apijson.Field
	RoutesValid          apijson.Field
	RoutesValidIPV4      apijson.Field
	RoutesValidIPV6      apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *BGPRouteStatsResponseStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRouteStatsResponseStatsJSON) RawJSON() string {
	return r.raw
}

type BGPRouteTimeseriesResponse struct {
	Meta          BGPRouteTimeseriesResponseMeta          `json:"meta,required"`
	SerieIPV4_24s BGPRouteTimeseriesResponseSerieIPV4_24s `json:"serie_ipv4_24s,required"`
	SerieIPV6_48s BGPRouteTimeseriesResponseSerieIPV6_48s `json:"serie_ipv6_48s,required"`
	JSON          bgpRouteTimeseriesResponseJSON          `json:"-"`
}

// bgpRouteTimeseriesResponseJSON contains the JSON metadata for the struct
// [BGPRouteTimeseriesResponse]
type bgpRouteTimeseriesResponseJSON struct {
	Meta          apijson.Field
	SerieIPV4_24s apijson.Field
	SerieIPV6_48s apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *BGPRouteTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRouteTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type BGPRouteTimeseriesResponseMeta struct {
	DateRange []BGPRouteTimeseriesResponseMetaDateRange `json:"dateRange,required"`
	JSON      bgpRouteTimeseriesResponseMetaJSON        `json:"-"`
}

// bgpRouteTimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [BGPRouteTimeseriesResponseMeta]
type bgpRouteTimeseriesResponseMetaJSON struct {
	DateRange   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPRouteTimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRouteTimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type BGPRouteTimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                   `json:"startTime,required" format:"date-time"`
	JSON      bgpRouteTimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// bgpRouteTimeseriesResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [BGPRouteTimeseriesResponseMetaDateRange]
type bgpRouteTimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPRouteTimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRouteTimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type BGPRouteTimeseriesResponseSerieIPV4_24s struct {
	Timestamps []time.Time                                 `json:"timestamps,required" format:"date-time"`
	Values     []int64                                     `json:"values,required"`
	JSON       bgpRouteTimeseriesResponseSerieIPV4_24sJSON `json:"-"`
}

// bgpRouteTimeseriesResponseSerieIPV4_24sJSON contains the JSON metadata for the
// struct [BGPRouteTimeseriesResponseSerieIPV4_24s]
type bgpRouteTimeseriesResponseSerieIPV4_24sJSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPRouteTimeseriesResponseSerieIPV4_24s) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRouteTimeseriesResponseSerieIPV4_24sJSON) RawJSON() string {
	return r.raw
}

type BGPRouteTimeseriesResponseSerieIPV6_48s struct {
	Timestamps []time.Time                                 `json:"timestamps,required" format:"date-time"`
	Values     []int64                                     `json:"values,required"`
	JSON       bgpRouteTimeseriesResponseSerieIPV6_48sJSON `json:"-"`
}

// bgpRouteTimeseriesResponseSerieIPV6_48sJSON contains the JSON metadata for the
// struct [BGPRouteTimeseriesResponseSerieIPV6_48s]
type bgpRouteTimeseriesResponseSerieIPV6_48sJSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPRouteTimeseriesResponseSerieIPV6_48s) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRouteTimeseriesResponseSerieIPV6_48sJSON) RawJSON() string {
	return r.raw
}

type BGPRouteMoasParams struct {
	// Format results are returned in.
	Format param.Field[BGPRouteMoasParamsFormat] `query:"format"`
	// Lookup only RPKI invalid MOASes
	InvalidOnly param.Field[bool] `query:"invalid_only"`
	// Lookup MOASes originated by the given ASN
	Origin param.Field[int64] `query:"origin"`
	// Lookup MOASes by prefix
	Prefix param.Field[string] `query:"prefix"`
}

// URLQuery serializes [BGPRouteMoasParams]'s query parameters as `url.Values`.
func (r BGPRouteMoasParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type BGPRouteMoasParamsFormat string

const (
	BGPRouteMoasParamsFormatJson BGPRouteMoasParamsFormat = "JSON"
	BGPRouteMoasParamsFormatCsv  BGPRouteMoasParamsFormat = "CSV"
)

func (r BGPRouteMoasParamsFormat) IsKnown() bool {
	switch r {
	case BGPRouteMoasParamsFormatJson, BGPRouteMoasParamsFormatCsv:
		return true
	}
	return false
}

type BGPRouteMoasResponseEnvelope struct {
	Result  BGPRouteMoasResponse             `json:"result,required"`
	Success bool                             `json:"success,required"`
	JSON    bgpRouteMoasResponseEnvelopeJSON `json:"-"`
}

// bgpRouteMoasResponseEnvelopeJSON contains the JSON metadata for the struct
// [BGPRouteMoasResponseEnvelope]
type bgpRouteMoasResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPRouteMoasResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRouteMoasResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BGPRoutePfx2asParams struct {
	// Format results are returned in.
	Format param.Field[BGPRoutePfx2asParamsFormat] `query:"format"`
	// Return only results with the longest prefix match for the given prefix. For
	// example, specify a /32 prefix to lookup the origin ASN for an IPv4 address.
	LongestPrefixMatch param.Field[bool] `query:"longestPrefixMatch"`
	// Lookup prefixes originated by the given ASN
	Origin param.Field[int64] `query:"origin"`
	// Lookup origin ASNs of the given prefix
	Prefix param.Field[string] `query:"prefix"`
	// Return only results with matching rpki status: valid, invalid or unknown
	RPKIStatus param.Field[BGPRoutePfx2asParamsRPKIStatus] `query:"rpkiStatus"`
}

// URLQuery serializes [BGPRoutePfx2asParams]'s query parameters as `url.Values`.
func (r BGPRoutePfx2asParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type BGPRoutePfx2asParamsFormat string

const (
	BGPRoutePfx2asParamsFormatJson BGPRoutePfx2asParamsFormat = "JSON"
	BGPRoutePfx2asParamsFormatCsv  BGPRoutePfx2asParamsFormat = "CSV"
)

func (r BGPRoutePfx2asParamsFormat) IsKnown() bool {
	switch r {
	case BGPRoutePfx2asParamsFormatJson, BGPRoutePfx2asParamsFormatCsv:
		return true
	}
	return false
}

// Return only results with matching rpki status: valid, invalid or unknown
type BGPRoutePfx2asParamsRPKIStatus string

const (
	BGPRoutePfx2asParamsRPKIStatusValid   BGPRoutePfx2asParamsRPKIStatus = "VALID"
	BGPRoutePfx2asParamsRPKIStatusInvalid BGPRoutePfx2asParamsRPKIStatus = "INVALID"
	BGPRoutePfx2asParamsRPKIStatusUnknown BGPRoutePfx2asParamsRPKIStatus = "UNKNOWN"
)

func (r BGPRoutePfx2asParamsRPKIStatus) IsKnown() bool {
	switch r {
	case BGPRoutePfx2asParamsRPKIStatusValid, BGPRoutePfx2asParamsRPKIStatusInvalid, BGPRoutePfx2asParamsRPKIStatusUnknown:
		return true
	}
	return false
}

type BGPRoutePfx2asResponseEnvelope struct {
	Result  BGPRoutePfx2asResponse             `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    bgpRoutePfx2asResponseEnvelopeJSON `json:"-"`
}

// bgpRoutePfx2asResponseEnvelopeJSON contains the JSON metadata for the struct
// [BGPRoutePfx2asResponseEnvelope]
type bgpRoutePfx2asResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPRoutePfx2asResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRoutePfx2asResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BGPRouteStatsParams struct {
	// Single ASN as integer.
	ASN param.Field[int64] `query:"asn"`
	// Format results are returned in.
	Format param.Field[BGPRouteStatsParamsFormat] `query:"format"`
	// Location Alpha2 code.
	Location param.Field[string] `query:"location"`
}

// URLQuery serializes [BGPRouteStatsParams]'s query parameters as `url.Values`.
func (r BGPRouteStatsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type BGPRouteStatsParamsFormat string

const (
	BGPRouteStatsParamsFormatJson BGPRouteStatsParamsFormat = "JSON"
	BGPRouteStatsParamsFormatCsv  BGPRouteStatsParamsFormat = "CSV"
)

func (r BGPRouteStatsParamsFormat) IsKnown() bool {
	switch r {
	case BGPRouteStatsParamsFormatJson, BGPRouteStatsParamsFormatCsv:
		return true
	}
	return false
}

type BGPRouteStatsResponseEnvelope struct {
	Result  BGPRouteStatsResponse             `json:"result,required"`
	Success bool                              `json:"success,required"`
	JSON    bgpRouteStatsResponseEnvelopeJSON `json:"-"`
}

// bgpRouteStatsResponseEnvelopeJSON contains the JSON metadata for the struct
// [BGPRouteStatsResponseEnvelope]
type bgpRouteStatsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPRouteStatsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRouteStatsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BGPRouteTimeseriesParams struct {
	// Single ASN as integer.
	ASN param.Field[int64] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Shorthand date ranges for the last X days - use when you don't need specific
	// start and end dates.
	DateRange param.Field[BGPRouteTimeseriesParamsDateRange] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[BGPRouteTimeseriesParamsFormat] `query:"format"`
	// Include data delay meta information
	IncludeDelay param.Field[bool] `query:"includeDelay"`
	// Location Alpha2 code.
	Location param.Field[string] `query:"location"`
}

// URLQuery serializes [BGPRouteTimeseriesParams]'s query parameters as
// `url.Values`.
func (r BGPRouteTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Shorthand date ranges for the last X days - use when you don't need specific
// start and end dates.
type BGPRouteTimeseriesParamsDateRange string

const (
	BGPRouteTimeseriesParamsDateRange1d         BGPRouteTimeseriesParamsDateRange = "1d"
	BGPRouteTimeseriesParamsDateRange2d         BGPRouteTimeseriesParamsDateRange = "2d"
	BGPRouteTimeseriesParamsDateRange7d         BGPRouteTimeseriesParamsDateRange = "7d"
	BGPRouteTimeseriesParamsDateRange14d        BGPRouteTimeseriesParamsDateRange = "14d"
	BGPRouteTimeseriesParamsDateRange28d        BGPRouteTimeseriesParamsDateRange = "28d"
	BGPRouteTimeseriesParamsDateRange12w        BGPRouteTimeseriesParamsDateRange = "12w"
	BGPRouteTimeseriesParamsDateRange24w        BGPRouteTimeseriesParamsDateRange = "24w"
	BGPRouteTimeseriesParamsDateRange52w        BGPRouteTimeseriesParamsDateRange = "52w"
	BGPRouteTimeseriesParamsDateRange1dControl  BGPRouteTimeseriesParamsDateRange = "1dControl"
	BGPRouteTimeseriesParamsDateRange2dControl  BGPRouteTimeseriesParamsDateRange = "2dControl"
	BGPRouteTimeseriesParamsDateRange7dControl  BGPRouteTimeseriesParamsDateRange = "7dControl"
	BGPRouteTimeseriesParamsDateRange14dControl BGPRouteTimeseriesParamsDateRange = "14dControl"
	BGPRouteTimeseriesParamsDateRange28dControl BGPRouteTimeseriesParamsDateRange = "28dControl"
	BGPRouteTimeseriesParamsDateRange12wControl BGPRouteTimeseriesParamsDateRange = "12wControl"
	BGPRouteTimeseriesParamsDateRange24wControl BGPRouteTimeseriesParamsDateRange = "24wControl"
)

func (r BGPRouteTimeseriesParamsDateRange) IsKnown() bool {
	switch r {
	case BGPRouteTimeseriesParamsDateRange1d, BGPRouteTimeseriesParamsDateRange2d, BGPRouteTimeseriesParamsDateRange7d, BGPRouteTimeseriesParamsDateRange14d, BGPRouteTimeseriesParamsDateRange28d, BGPRouteTimeseriesParamsDateRange12w, BGPRouteTimeseriesParamsDateRange24w, BGPRouteTimeseriesParamsDateRange52w, BGPRouteTimeseriesParamsDateRange1dControl, BGPRouteTimeseriesParamsDateRange2dControl, BGPRouteTimeseriesParamsDateRange7dControl, BGPRouteTimeseriesParamsDateRange14dControl, BGPRouteTimeseriesParamsDateRange28dControl, BGPRouteTimeseriesParamsDateRange12wControl, BGPRouteTimeseriesParamsDateRange24wControl:
		return true
	}
	return false
}

// Format results are returned in.
type BGPRouteTimeseriesParamsFormat string

const (
	BGPRouteTimeseriesParamsFormatJson BGPRouteTimeseriesParamsFormat = "JSON"
	BGPRouteTimeseriesParamsFormatCsv  BGPRouteTimeseriesParamsFormat = "CSV"
)

func (r BGPRouteTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case BGPRouteTimeseriesParamsFormatJson, BGPRouteTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type BGPRouteTimeseriesResponseEnvelope struct {
	Result  BGPRouteTimeseriesResponse             `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    bgpRouteTimeseriesResponseEnvelopeJSON `json:"-"`
}

// bgpRouteTimeseriesResponseEnvelopeJSON contains the JSON metadata for the struct
// [BGPRouteTimeseriesResponseEnvelope]
type bgpRouteTimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPRouteTimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRouteTimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
