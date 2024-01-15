// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarBgpRouteStatService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarBgpRouteStatService] method
// instead.
type RadarBgpRouteStatService struct {
	Options []option.RequestOption
}

// NewRadarBgpRouteStatService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarBgpRouteStatService(opts ...option.RequestOption) (r *RadarBgpRouteStatService) {
	r = &RadarBgpRouteStatService{}
	r.Options = opts
	return
}

// Get the BGP routing table stats (Beta).
func (r *RadarBgpRouteStatService) List(ctx context.Context, query RadarBgpRouteStatListParams, opts ...option.RequestOption) (res *RadarBgpRouteStatListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/bgp/routes/stats"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarBgpRouteStatListResponse struct {
	Result  RadarBgpRouteStatListResponseResult `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    radarBgpRouteStatListResponseJSON   `json:"-"`
}

// radarBgpRouteStatListResponseJSON contains the JSON metadata for the struct
// [RadarBgpRouteStatListResponse]
type radarBgpRouteStatListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpRouteStatListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpRouteStatListResponseResult struct {
	Meta  RadarBgpRouteStatListResponseResultMeta  `json:"meta,required"`
	Stats RadarBgpRouteStatListResponseResultStats `json:"stats,required"`
	JSON  radarBgpRouteStatListResponseResultJSON  `json:"-"`
}

// radarBgpRouteStatListResponseResultJSON contains the JSON metadata for the
// struct [RadarBgpRouteStatListResponseResult]
type radarBgpRouteStatListResponseResultJSON struct {
	Meta        apijson.Field
	Stats       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpRouteStatListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpRouteStatListResponseResultMeta struct {
	DataTime   string                                      `json:"data_time,required"`
	QueryTime  string                                      `json:"query_time,required"`
	TotalPeers int64                                       `json:"total_peers,required"`
	JSON       radarBgpRouteStatListResponseResultMetaJSON `json:"-"`
}

// radarBgpRouteStatListResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarBgpRouteStatListResponseResultMeta]
type radarBgpRouteStatListResponseResultMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalPeers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpRouteStatListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpRouteStatListResponseResultStats struct {
	DistinctOrigins      int64                                        `json:"distinct_origins,required"`
	DistinctOriginsIpv4  int64                                        `json:"distinct_origins_ipv4,required"`
	DistinctOriginsIpv6  int64                                        `json:"distinct_origins_ipv6,required"`
	DistinctPrefixes     int64                                        `json:"distinct_prefixes,required"`
	DistinctPrefixesIpv4 int64                                        `json:"distinct_prefixes_ipv4,required"`
	DistinctPrefixesIpv6 int64                                        `json:"distinct_prefixes_ipv6,required"`
	RoutesInvalid        int64                                        `json:"routes_invalid,required"`
	RoutesInvalidIpv4    int64                                        `json:"routes_invalid_ipv4,required"`
	RoutesInvalidIpv6    int64                                        `json:"routes_invalid_ipv6,required"`
	RoutesTotal          int64                                        `json:"routes_total,required"`
	RoutesTotalIpv4      int64                                        `json:"routes_total_ipv4,required"`
	RoutesTotalIpv6      int64                                        `json:"routes_total_ipv6,required"`
	RoutesUnknown        int64                                        `json:"routes_unknown,required"`
	RoutesUnknownIpv4    int64                                        `json:"routes_unknown_ipv4,required"`
	RoutesUnknownIpv6    int64                                        `json:"routes_unknown_ipv6,required"`
	RoutesValid          int64                                        `json:"routes_valid,required"`
	RoutesValidIpv4      int64                                        `json:"routes_valid_ipv4,required"`
	RoutesValidIpv6      int64                                        `json:"routes_valid_ipv6,required"`
	JSON                 radarBgpRouteStatListResponseResultStatsJSON `json:"-"`
}

// radarBgpRouteStatListResponseResultStatsJSON contains the JSON metadata for the
// struct [RadarBgpRouteStatListResponseResultStats]
type radarBgpRouteStatListResponseResultStatsJSON struct {
	DistinctOrigins      apijson.Field
	DistinctOriginsIpv4  apijson.Field
	DistinctOriginsIpv6  apijson.Field
	DistinctPrefixes     apijson.Field
	DistinctPrefixesIpv4 apijson.Field
	DistinctPrefixesIpv6 apijson.Field
	RoutesInvalid        apijson.Field
	RoutesInvalidIpv4    apijson.Field
	RoutesInvalidIpv6    apijson.Field
	RoutesTotal          apijson.Field
	RoutesTotalIpv4      apijson.Field
	RoutesTotalIpv6      apijson.Field
	RoutesUnknown        apijson.Field
	RoutesUnknownIpv4    apijson.Field
	RoutesUnknownIpv6    apijson.Field
	RoutesValid          apijson.Field
	RoutesValidIpv4      apijson.Field
	RoutesValidIpv6      apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *RadarBgpRouteStatListResponseResultStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpRouteStatListParams struct {
	// Single ASN as integer.
	ASN param.Field[int64] `query:"asn"`
	// Format results are returned in.
	Format param.Field[RadarBgpRouteStatListParamsFormat] `query:"format"`
	// Location Alpha2 code.
	Location param.Field[string] `query:"location"`
}

// URLQuery serializes [RadarBgpRouteStatListParams]'s query parameters as
// `url.Values`.
func (r RadarBgpRouteStatListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarBgpRouteStatListParamsFormat string

const (
	RadarBgpRouteStatListParamsFormatJson RadarBgpRouteStatListParamsFormat = "JSON"
	RadarBgpRouteStatListParamsFormatCsv  RadarBgpRouteStatListParamsFormat = "CSV"
)
