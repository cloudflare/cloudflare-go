// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// RadarBGPRouteService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarBGPRouteService] method
// instead.
type RadarBGPRouteService struct {
	Options []option.RequestOption
}

// NewRadarBGPRouteService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarBGPRouteService(opts ...option.RequestOption) (r *RadarBGPRouteService) {
	r = &RadarBGPRouteService{}
	r.Options = opts
	return
}

// List all Multi-origin AS (MOAS) prefixes on the global routing tables.
func (r *RadarBGPRouteService) Moas(ctx context.Context, query RadarBGPRouteMoasParams, opts ...option.RequestOption) (res *RadarBGPRouteMoasResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarBGPRouteMoasResponseEnvelope
	path := "radar/bgp/routes/moas"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lookup prefix-to-origin mapping on global routing tables.
func (r *RadarBGPRouteService) Pfx2as(ctx context.Context, query RadarBGPRoutePfx2asParams, opts ...option.RequestOption) (res *RadarBGPRoutePfx2asResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarBGPRoutePfx2asResponseEnvelope
	path := "radar/bgp/routes/pfx2as"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the BGP routing table stats (Beta).
func (r *RadarBGPRouteService) Stats(ctx context.Context, query RadarBGPRouteStatsParams, opts ...option.RequestOption) (res *RadarBGPRouteStatsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarBGPRouteStatsResponseEnvelope
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
func (r *RadarBGPRouteService) Timeseries(ctx context.Context, query RadarBGPRouteTimeseriesParams, opts ...option.RequestOption) (res *RadarBGPRouteTimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarBGPRouteTimeseriesResponseEnvelope
	path := "radar/bgp/routes/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarBGPRouteMoasResponse struct {
	Meta RadarBGPRouteMoasResponseMeta  `json:"meta,required"`
	Moas []RadarBGPRouteMoasResponseMoa `json:"moas,required"`
	JSON radarBGPRouteMoasResponseJSON  `json:"-"`
}

// radarBGPRouteMoasResponseJSON contains the JSON metadata for the struct
// [RadarBGPRouteMoasResponse]
type radarBGPRouteMoasResponseJSON struct {
	Meta        apijson.Field
	Moas        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPRouteMoasResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBGPRouteMoasResponseJSON) RawJSON() string {
	return r.raw
}

type RadarBGPRouteMoasResponseMeta struct {
	DataTime   string                            `json:"data_time,required"`
	QueryTime  string                            `json:"query_time,required"`
	TotalPeers int64                             `json:"total_peers,required"`
	JSON       radarBGPRouteMoasResponseMetaJSON `json:"-"`
}

// radarBGPRouteMoasResponseMetaJSON contains the JSON metadata for the struct
// [RadarBGPRouteMoasResponseMeta]
type radarBGPRouteMoasResponseMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalPeers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPRouteMoasResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBGPRouteMoasResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarBGPRouteMoasResponseMoa struct {
	Origins []RadarBGPRouteMoasResponseMoasOrigin `json:"origins,required"`
	Prefix  string                                `json:"prefix,required"`
	JSON    radarBGPRouteMoasResponseMoaJSON      `json:"-"`
}

// radarBGPRouteMoasResponseMoaJSON contains the JSON metadata for the struct
// [RadarBGPRouteMoasResponseMoa]
type radarBGPRouteMoasResponseMoaJSON struct {
	Origins     apijson.Field
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPRouteMoasResponseMoa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBGPRouteMoasResponseMoaJSON) RawJSON() string {
	return r.raw
}

type RadarBGPRouteMoasResponseMoasOrigin struct {
	Origin         int64                                   `json:"origin,required"`
	PeerCount      int64                                   `json:"peer_count,required"`
	RpkiValidation string                                  `json:"rpki_validation,required"`
	JSON           radarBGPRouteMoasResponseMoasOriginJSON `json:"-"`
}

// radarBGPRouteMoasResponseMoasOriginJSON contains the JSON metadata for the
// struct [RadarBGPRouteMoasResponseMoasOrigin]
type radarBGPRouteMoasResponseMoasOriginJSON struct {
	Origin         apijson.Field
	PeerCount      apijson.Field
	RpkiValidation apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarBGPRouteMoasResponseMoasOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBGPRouteMoasResponseMoasOriginJSON) RawJSON() string {
	return r.raw
}

type RadarBGPRoutePfx2asResponse struct {
	Meta          RadarBGPRoutePfx2asResponseMeta           `json:"meta,required"`
	PrefixOrigins []RadarBGPRoutePfx2asResponsePrefixOrigin `json:"prefix_origins,required"`
	JSON          radarBGPRoutePfx2asResponseJSON           `json:"-"`
}

// radarBGPRoutePfx2asResponseJSON contains the JSON metadata for the struct
// [RadarBGPRoutePfx2asResponse]
type radarBGPRoutePfx2asResponseJSON struct {
	Meta          apijson.Field
	PrefixOrigins apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RadarBGPRoutePfx2asResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBGPRoutePfx2asResponseJSON) RawJSON() string {
	return r.raw
}

type RadarBGPRoutePfx2asResponseMeta struct {
	DataTime   string                              `json:"data_time,required"`
	QueryTime  string                              `json:"query_time,required"`
	TotalPeers int64                               `json:"total_peers,required"`
	JSON       radarBGPRoutePfx2asResponseMetaJSON `json:"-"`
}

// radarBGPRoutePfx2asResponseMetaJSON contains the JSON metadata for the struct
// [RadarBGPRoutePfx2asResponseMeta]
type radarBGPRoutePfx2asResponseMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalPeers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPRoutePfx2asResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBGPRoutePfx2asResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarBGPRoutePfx2asResponsePrefixOrigin struct {
	Origin         int64                                       `json:"origin,required"`
	PeerCount      int64                                       `json:"peer_count,required"`
	Prefix         string                                      `json:"prefix,required"`
	RpkiValidation string                                      `json:"rpki_validation,required"`
	JSON           radarBGPRoutePfx2asResponsePrefixOriginJSON `json:"-"`
}

// radarBGPRoutePfx2asResponsePrefixOriginJSON contains the JSON metadata for the
// struct [RadarBGPRoutePfx2asResponsePrefixOrigin]
type radarBGPRoutePfx2asResponsePrefixOriginJSON struct {
	Origin         apijson.Field
	PeerCount      apijson.Field
	Prefix         apijson.Field
	RpkiValidation apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarBGPRoutePfx2asResponsePrefixOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBGPRoutePfx2asResponsePrefixOriginJSON) RawJSON() string {
	return r.raw
}

type RadarBGPRouteStatsResponse struct {
	Meta  RadarBGPRouteStatsResponseMeta  `json:"meta,required"`
	Stats RadarBGPRouteStatsResponseStats `json:"stats,required"`
	JSON  radarBGPRouteStatsResponseJSON  `json:"-"`
}

// radarBGPRouteStatsResponseJSON contains the JSON metadata for the struct
// [RadarBGPRouteStatsResponse]
type radarBGPRouteStatsResponseJSON struct {
	Meta        apijson.Field
	Stats       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPRouteStatsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBGPRouteStatsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarBGPRouteStatsResponseMeta struct {
	DataTime   string                             `json:"data_time,required"`
	QueryTime  string                             `json:"query_time,required"`
	TotalPeers int64                              `json:"total_peers,required"`
	JSON       radarBGPRouteStatsResponseMetaJSON `json:"-"`
}

// radarBGPRouteStatsResponseMetaJSON contains the JSON metadata for the struct
// [RadarBGPRouteStatsResponseMeta]
type radarBGPRouteStatsResponseMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalPeers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPRouteStatsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBGPRouteStatsResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarBGPRouteStatsResponseStats struct {
	DistinctOrigins      int64                               `json:"distinct_origins,required"`
	DistinctOriginsIPV4  int64                               `json:"distinct_origins_ipv4,required"`
	DistinctOriginsIPV6  int64                               `json:"distinct_origins_ipv6,required"`
	DistinctPrefixes     int64                               `json:"distinct_prefixes,required"`
	DistinctPrefixesIPV4 int64                               `json:"distinct_prefixes_ipv4,required"`
	DistinctPrefixesIPV6 int64                               `json:"distinct_prefixes_ipv6,required"`
	RoutesInvalid        int64                               `json:"routes_invalid,required"`
	RoutesInvalidIPV4    int64                               `json:"routes_invalid_ipv4,required"`
	RoutesInvalidIPV6    int64                               `json:"routes_invalid_ipv6,required"`
	RoutesTotal          int64                               `json:"routes_total,required"`
	RoutesTotalIPV4      int64                               `json:"routes_total_ipv4,required"`
	RoutesTotalIPV6      int64                               `json:"routes_total_ipv6,required"`
	RoutesUnknown        int64                               `json:"routes_unknown,required"`
	RoutesUnknownIPV4    int64                               `json:"routes_unknown_ipv4,required"`
	RoutesUnknownIPV6    int64                               `json:"routes_unknown_ipv6,required"`
	RoutesValid          int64                               `json:"routes_valid,required"`
	RoutesValidIPV4      int64                               `json:"routes_valid_ipv4,required"`
	RoutesValidIPV6      int64                               `json:"routes_valid_ipv6,required"`
	JSON                 radarBGPRouteStatsResponseStatsJSON `json:"-"`
}

// radarBGPRouteStatsResponseStatsJSON contains the JSON metadata for the struct
// [RadarBGPRouteStatsResponseStats]
type radarBGPRouteStatsResponseStatsJSON struct {
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

func (r *RadarBGPRouteStatsResponseStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBGPRouteStatsResponseStatsJSON) RawJSON() string {
	return r.raw
}

type RadarBGPRouteTimeseriesResponse struct {
	Meta          RadarBGPRouteTimeseriesResponseMeta          `json:"meta,required"`
	SerieIPV4_24s RadarBGPRouteTimeseriesResponseSerieIPV4_24s `json:"serie_ipv4_24s,required"`
	SerieIPV6_48s RadarBGPRouteTimeseriesResponseSerieIPV6_48s `json:"serie_ipv6_48s,required"`
	JSON          radarBGPRouteTimeseriesResponseJSON          `json:"-"`
}

// radarBGPRouteTimeseriesResponseJSON contains the JSON metadata for the struct
// [RadarBGPRouteTimeseriesResponse]
type radarBGPRouteTimeseriesResponseJSON struct {
	Meta          apijson.Field
	SerieIPV4_24s apijson.Field
	SerieIPV6_48s apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RadarBGPRouteTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBGPRouteTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarBGPRouteTimeseriesResponseMeta struct {
	DateRange []RadarBGPRouteTimeseriesResponseMetaDateRange `json:"dateRange,required"`
	JSON      radarBGPRouteTimeseriesResponseMetaJSON        `json:"-"`
}

// radarBGPRouteTimeseriesResponseMetaJSON contains the JSON metadata for the
// struct [RadarBGPRouteTimeseriesResponseMeta]
type radarBGPRouteTimeseriesResponseMetaJSON struct {
	DateRange   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPRouteTimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBGPRouteTimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarBGPRouteTimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                        `json:"startTime,required" format:"date-time"`
	JSON      radarBGPRouteTimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// radarBGPRouteTimeseriesResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [RadarBGPRouteTimeseriesResponseMetaDateRange]
type radarBGPRouteTimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPRouteTimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBGPRouteTimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarBGPRouteTimeseriesResponseSerieIPV4_24s struct {
	Timestamps []time.Time                                      `json:"timestamps,required" format:"date-time"`
	Values     []int64                                          `json:"values,required"`
	JSON       radarBGPRouteTimeseriesResponseSerieIPV4_24sJSON `json:"-"`
}

// radarBGPRouteTimeseriesResponseSerieIPV4_24sJSON contains the JSON metadata for
// the struct [RadarBGPRouteTimeseriesResponseSerieIPV4_24s]
type radarBGPRouteTimeseriesResponseSerieIPV4_24sJSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPRouteTimeseriesResponseSerieIPV4_24s) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBGPRouteTimeseriesResponseSerieIPV4_24sJSON) RawJSON() string {
	return r.raw
}

type RadarBGPRouteTimeseriesResponseSerieIPV6_48s struct {
	Timestamps []time.Time                                      `json:"timestamps,required" format:"date-time"`
	Values     []int64                                          `json:"values,required"`
	JSON       radarBGPRouteTimeseriesResponseSerieIPV6_48sJSON `json:"-"`
}

// radarBGPRouteTimeseriesResponseSerieIPV6_48sJSON contains the JSON metadata for
// the struct [RadarBGPRouteTimeseriesResponseSerieIPV6_48s]
type radarBGPRouteTimeseriesResponseSerieIPV6_48sJSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPRouteTimeseriesResponseSerieIPV6_48s) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBGPRouteTimeseriesResponseSerieIPV6_48sJSON) RawJSON() string {
	return r.raw
}

type RadarBGPRouteMoasParams struct {
	// Format results are returned in.
	Format param.Field[RadarBGPRouteMoasParamsFormat] `query:"format"`
	// Lookup only RPKI invalid MOASes
	InvalidOnly param.Field[bool] `query:"invalid_only"`
	// Lookup MOASes originated by the given ASN
	Origin param.Field[int64] `query:"origin"`
	// Lookup MOASes by prefix
	Prefix param.Field[string] `query:"prefix"`
}

// URLQuery serializes [RadarBGPRouteMoasParams]'s query parameters as
// `url.Values`.
func (r RadarBGPRouteMoasParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarBGPRouteMoasParamsFormat string

const (
	RadarBGPRouteMoasParamsFormatJson RadarBGPRouteMoasParamsFormat = "JSON"
	RadarBGPRouteMoasParamsFormatCsv  RadarBGPRouteMoasParamsFormat = "CSV"
)

type RadarBGPRouteMoasResponseEnvelope struct {
	Result  RadarBGPRouteMoasResponse             `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    radarBGPRouteMoasResponseEnvelopeJSON `json:"-"`
}

// radarBGPRouteMoasResponseEnvelopeJSON contains the JSON metadata for the struct
// [RadarBGPRouteMoasResponseEnvelope]
type radarBGPRouteMoasResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPRouteMoasResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBGPRouteMoasResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarBGPRoutePfx2asParams struct {
	// Format results are returned in.
	Format param.Field[RadarBGPRoutePfx2asParamsFormat] `query:"format"`
	// Lookup prefixes originated by the given ASN
	Origin param.Field[int64] `query:"origin"`
	// Lookup origins of the given prefix
	Prefix param.Field[string] `query:"prefix"`
	// Return only results with matching rpki status: valid, invalid or unknown
	RpkiStatus param.Field[RadarBGPRoutePfx2asParamsRpkiStatus] `query:"rpkiStatus"`
}

// URLQuery serializes [RadarBGPRoutePfx2asParams]'s query parameters as
// `url.Values`.
func (r RadarBGPRoutePfx2asParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarBGPRoutePfx2asParamsFormat string

const (
	RadarBGPRoutePfx2asParamsFormatJson RadarBGPRoutePfx2asParamsFormat = "JSON"
	RadarBGPRoutePfx2asParamsFormatCsv  RadarBGPRoutePfx2asParamsFormat = "CSV"
)

// Return only results with matching rpki status: valid, invalid or unknown
type RadarBGPRoutePfx2asParamsRpkiStatus string

const (
	RadarBGPRoutePfx2asParamsRpkiStatusValid   RadarBGPRoutePfx2asParamsRpkiStatus = "VALID"
	RadarBGPRoutePfx2asParamsRpkiStatusInvalid RadarBGPRoutePfx2asParamsRpkiStatus = "INVALID"
	RadarBGPRoutePfx2asParamsRpkiStatusUnknown RadarBGPRoutePfx2asParamsRpkiStatus = "UNKNOWN"
)

type RadarBGPRoutePfx2asResponseEnvelope struct {
	Result  RadarBGPRoutePfx2asResponse             `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    radarBGPRoutePfx2asResponseEnvelopeJSON `json:"-"`
}

// radarBGPRoutePfx2asResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarBGPRoutePfx2asResponseEnvelope]
type radarBGPRoutePfx2asResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPRoutePfx2asResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBGPRoutePfx2asResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarBGPRouteStatsParams struct {
	// Single ASN as integer.
	ASN param.Field[int64] `query:"asn"`
	// Format results are returned in.
	Format param.Field[RadarBGPRouteStatsParamsFormat] `query:"format"`
	// Location Alpha2 code.
	Location param.Field[string] `query:"location"`
}

// URLQuery serializes [RadarBGPRouteStatsParams]'s query parameters as
// `url.Values`.
func (r RadarBGPRouteStatsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarBGPRouteStatsParamsFormat string

const (
	RadarBGPRouteStatsParamsFormatJson RadarBGPRouteStatsParamsFormat = "JSON"
	RadarBGPRouteStatsParamsFormatCsv  RadarBGPRouteStatsParamsFormat = "CSV"
)

type RadarBGPRouteStatsResponseEnvelope struct {
	Result  RadarBGPRouteStatsResponse             `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    radarBGPRouteStatsResponseEnvelopeJSON `json:"-"`
}

// radarBGPRouteStatsResponseEnvelopeJSON contains the JSON metadata for the struct
// [RadarBGPRouteStatsResponseEnvelope]
type radarBGPRouteStatsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPRouteStatsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBGPRouteStatsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarBGPRouteTimeseriesParams struct {
	// Single ASN as integer.
	ASN param.Field[int64] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Shorthand date ranges for the last X days - use when you don't need specific
	// start and end dates.
	DateRange param.Field[RadarBGPRouteTimeseriesParamsDateRange] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarBGPRouteTimeseriesParamsFormat] `query:"format"`
}

// URLQuery serializes [RadarBGPRouteTimeseriesParams]'s query parameters as
// `url.Values`.
func (r RadarBGPRouteTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Shorthand date ranges for the last X days - use when you don't need specific
// start and end dates.
type RadarBGPRouteTimeseriesParamsDateRange string

const (
	RadarBGPRouteTimeseriesParamsDateRange1d         RadarBGPRouteTimeseriesParamsDateRange = "1d"
	RadarBGPRouteTimeseriesParamsDateRange2d         RadarBGPRouteTimeseriesParamsDateRange = "2d"
	RadarBGPRouteTimeseriesParamsDateRange7d         RadarBGPRouteTimeseriesParamsDateRange = "7d"
	RadarBGPRouteTimeseriesParamsDateRange14d        RadarBGPRouteTimeseriesParamsDateRange = "14d"
	RadarBGPRouteTimeseriesParamsDateRange28d        RadarBGPRouteTimeseriesParamsDateRange = "28d"
	RadarBGPRouteTimeseriesParamsDateRange12w        RadarBGPRouteTimeseriesParamsDateRange = "12w"
	RadarBGPRouteTimeseriesParamsDateRange24w        RadarBGPRouteTimeseriesParamsDateRange = "24w"
	RadarBGPRouteTimeseriesParamsDateRange52w        RadarBGPRouteTimeseriesParamsDateRange = "52w"
	RadarBGPRouteTimeseriesParamsDateRange1dControl  RadarBGPRouteTimeseriesParamsDateRange = "1dControl"
	RadarBGPRouteTimeseriesParamsDateRange2dControl  RadarBGPRouteTimeseriesParamsDateRange = "2dControl"
	RadarBGPRouteTimeseriesParamsDateRange7dControl  RadarBGPRouteTimeseriesParamsDateRange = "7dControl"
	RadarBGPRouteTimeseriesParamsDateRange14dControl RadarBGPRouteTimeseriesParamsDateRange = "14dControl"
	RadarBGPRouteTimeseriesParamsDateRange28dControl RadarBGPRouteTimeseriesParamsDateRange = "28dControl"
	RadarBGPRouteTimeseriesParamsDateRange12wControl RadarBGPRouteTimeseriesParamsDateRange = "12wControl"
	RadarBGPRouteTimeseriesParamsDateRange24wControl RadarBGPRouteTimeseriesParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarBGPRouteTimeseriesParamsFormat string

const (
	RadarBGPRouteTimeseriesParamsFormatJson RadarBGPRouteTimeseriesParamsFormat = "JSON"
	RadarBGPRouteTimeseriesParamsFormatCsv  RadarBGPRouteTimeseriesParamsFormat = "CSV"
)

type RadarBGPRouteTimeseriesResponseEnvelope struct {
	Result  RadarBGPRouteTimeseriesResponse             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    radarBGPRouteTimeseriesResponseEnvelopeJSON `json:"-"`
}

// radarBGPRouteTimeseriesResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarBGPRouteTimeseriesResponseEnvelope]
type radarBGPRouteTimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPRouteTimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBGPRouteTimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
