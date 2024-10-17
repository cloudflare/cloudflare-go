// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// BGPRouteService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBGPRouteService] method instead.
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

// List all ASes in current global routing tables with routing statistics
func (r *BGPRouteService) Ases(ctx context.Context, query BGPRouteAsesParams, opts ...option.RequestOption) (res *BGPRouteAsesResponse, err error) {
	var env BGPRouteAsesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/bgp/routes/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all Multi-Origin AS (MOAS) prefixes on the global routing tables.
func (r *BGPRouteService) Moas(ctx context.Context, query BGPRouteMoasParams, opts ...option.RequestOption) (res *BGPRouteMoasResponse, err error) {
	var env BGPRouteMoasResponseEnvelope
	opts = append(r.Options[:], opts...)
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
	var env BGPRoutePfx2asResponseEnvelope
	opts = append(r.Options[:], opts...)
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
	var env BGPRouteStatsResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/bgp/routes/stats"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BGPRouteAsesResponse struct {
	ASNs []BGPRouteAsesResponseASN `json:"asns,required"`
	Meta BGPRouteAsesResponseMeta  `json:"meta,required"`
	JSON bgpRouteAsesResponseJSON  `json:"-"`
}

// bgpRouteAsesResponseJSON contains the JSON metadata for the struct
// [BGPRouteAsesResponse]
type bgpRouteAsesResponseJSON struct {
	ASNs        apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPRouteAsesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRouteAsesResponseJSON) RawJSON() string {
	return r.raw
}

type BGPRouteAsesResponseASN struct {
	ASN int64 `json:"asn,required"`
	// AS's customer cone size
	ConeSize int64 `json:"coneSize,required"`
	// 2-letter country code for the AS's registration country
	Country string `json:"country,required"`
	// number of IPv4 addresses originated by the AS
	IPV4Count int64 `json:"ipv4Count,required"`
	// number of IPv6 addresses originated by the AS
	IPV6Count string `json:"ipv6Count,required"`
	// name of the AS
	Name string `json:"name,required"`
	// number of total IP prefixes originated by the AS
	PfxsCount int64 `json:"pfxsCount,required"`
	// number of RPKI invalid prefixes originated by the AS
	RPKIInvalid int64 `json:"rpkiInvalid,required"`
	// number of RPKI unknown prefixes originated by the AS
	RPKIUnknown int64 `json:"rpkiUnknown,required"`
	// number of RPKI valid prefixes originated by the AS
	RPKIValid int64                       `json:"rpkiValid,required"`
	JSON      bgpRouteAsesResponseASNJSON `json:"-"`
}

// bgpRouteAsesResponseASNJSON contains the JSON metadata for the struct
// [BGPRouteAsesResponseASN]
type bgpRouteAsesResponseASNJSON struct {
	ASN         apijson.Field
	ConeSize    apijson.Field
	Country     apijson.Field
	IPV4Count   apijson.Field
	IPV6Count   apijson.Field
	Name        apijson.Field
	PfxsCount   apijson.Field
	RPKIInvalid apijson.Field
	RPKIUnknown apijson.Field
	RPKIValid   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPRouteAsesResponseASN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRouteAsesResponseASNJSON) RawJSON() string {
	return r.raw
}

type BGPRouteAsesResponseMeta struct {
	// the timestamp of when the data is generated
	DataTime string `json:"dataTime,required"`
	// the timestamp of the query
	QueryTime string `json:"queryTime,required"`
	// total number of route collector peers used to generate this data
	TotalPeers int64                        `json:"totalPeers,required"`
	JSON       bgpRouteAsesResponseMetaJSON `json:"-"`
}

// bgpRouteAsesResponseMetaJSON contains the JSON metadata for the struct
// [BGPRouteAsesResponseMeta]
type bgpRouteAsesResponseMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalPeers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPRouteAsesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRouteAsesResponseMetaJSON) RawJSON() string {
	return r.raw
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

type BGPRouteAsesParams struct {
	// Format results are returned in.
	Format param.Field[BGPRouteAsesParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Location Alpha2 code.
	Location param.Field[string] `query:"location"`
	// Return order results by given type
	SortBy param.Field[BGPRouteAsesParamsSortBy] `query:"sortBy"`
	// Sort by value ascending or descending
	SortOrder param.Field[BGPRouteAsesParamsSortOrder] `query:"sortOrder"`
}

// URLQuery serializes [BGPRouteAsesParams]'s query parameters as `url.Values`.
func (r BGPRouteAsesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format results are returned in.
type BGPRouteAsesParamsFormat string

const (
	BGPRouteAsesParamsFormatJson BGPRouteAsesParamsFormat = "JSON"
	BGPRouteAsesParamsFormatCsv  BGPRouteAsesParamsFormat = "CSV"
)

func (r BGPRouteAsesParamsFormat) IsKnown() bool {
	switch r {
	case BGPRouteAsesParamsFormatJson, BGPRouteAsesParamsFormatCsv:
		return true
	}
	return false
}

// Return order results by given type
type BGPRouteAsesParamsSortBy string

const (
	BGPRouteAsesParamsSortByCone        BGPRouteAsesParamsSortBy = "cone"
	BGPRouteAsesParamsSortByPfxs        BGPRouteAsesParamsSortBy = "pfxs"
	BGPRouteAsesParamsSortByIPV4        BGPRouteAsesParamsSortBy = "ipv4"
	BGPRouteAsesParamsSortByIPV6        BGPRouteAsesParamsSortBy = "ipv6"
	BGPRouteAsesParamsSortByRPKIValid   BGPRouteAsesParamsSortBy = "rpki_valid"
	BGPRouteAsesParamsSortByRPKIInvalid BGPRouteAsesParamsSortBy = "rpki_invalid"
	BGPRouteAsesParamsSortByRPKIUnknown BGPRouteAsesParamsSortBy = "rpki_unknown"
)

func (r BGPRouteAsesParamsSortBy) IsKnown() bool {
	switch r {
	case BGPRouteAsesParamsSortByCone, BGPRouteAsesParamsSortByPfxs, BGPRouteAsesParamsSortByIPV4, BGPRouteAsesParamsSortByIPV6, BGPRouteAsesParamsSortByRPKIValid, BGPRouteAsesParamsSortByRPKIInvalid, BGPRouteAsesParamsSortByRPKIUnknown:
		return true
	}
	return false
}

// Sort by value ascending or descending
type BGPRouteAsesParamsSortOrder string

const (
	BGPRouteAsesParamsSortOrderAsc  BGPRouteAsesParamsSortOrder = "asc"
	BGPRouteAsesParamsSortOrderDesc BGPRouteAsesParamsSortOrder = "desc"
)

func (r BGPRouteAsesParamsSortOrder) IsKnown() bool {
	switch r {
	case BGPRouteAsesParamsSortOrderAsc, BGPRouteAsesParamsSortOrderDesc:
		return true
	}
	return false
}

type BGPRouteAsesResponseEnvelope struct {
	Result  BGPRouteAsesResponse             `json:"result,required"`
	Success bool                             `json:"success,required"`
	JSON    bgpRouteAsesResponseEnvelopeJSON `json:"-"`
}

// bgpRouteAsesResponseEnvelopeJSON contains the JSON metadata for the struct
// [BGPRouteAsesResponseEnvelope]
type bgpRouteAsesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPRouteAsesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRouteAsesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BGPRouteMoasParams struct {
	// Format results are returned in.
	Format param.Field[BGPRouteMoasParamsFormat] `query:"format"`
	// Lookup only RPKI invalid MOASes
	InvalidOnly param.Field[bool] `query:"invalid_only"`
	// Lookup MOASes originated by the given ASN
	Origin param.Field[int64] `query:"origin"`
	// Network prefix, IPv4 or IPv6.
	Prefix param.Field[string] `query:"prefix"`
}

// URLQuery serializes [BGPRouteMoasParams]'s query parameters as `url.Values`.
func (r BGPRouteMoasParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
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
	// Network prefix, IPv4 or IPv6.
	Prefix param.Field[string] `query:"prefix"`
	// Return only results with matching rpki status: valid, invalid or unknown
	RPKIStatus param.Field[BGPRoutePfx2asParamsRPKIStatus] `query:"rpkiStatus"`
}

// URLQuery serializes [BGPRoutePfx2asParams]'s query parameters as `url.Values`.
func (r BGPRoutePfx2asParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
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
		NestedFormat: apiquery.NestedQueryFormatDots,
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
