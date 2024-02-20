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

// RadarBGPRouteStatService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarBGPRouteStatService] method
// instead.
type RadarBGPRouteStatService struct {
	Options []option.RequestOption
}

// NewRadarBGPRouteStatService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarBGPRouteStatService(opts ...option.RequestOption) (r *RadarBGPRouteStatService) {
	r = &RadarBGPRouteStatService{}
	r.Options = opts
	return
}

// Get the BGP routing table stats (Beta).
func (r *RadarBGPRouteStatService) List(ctx context.Context, query RadarBGPRouteStatListParams, opts ...option.RequestOption) (res *RadarBGPRouteStatListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarBGPRouteStatListResponseEnvelope
	path := "radar/bgp/routes/stats"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarBGPRouteStatListResponse struct {
	Meta  RadarBGPRouteStatListResponseMeta  `json:"meta,required"`
	Stats RadarBGPRouteStatListResponseStats `json:"stats,required"`
	JSON  radarBGPRouteStatListResponseJSON  `json:"-"`
}

// radarBGPRouteStatListResponseJSON contains the JSON metadata for the struct
// [RadarBGPRouteStatListResponse]
type radarBGPRouteStatListResponseJSON struct {
	Meta        apijson.Field
	Stats       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPRouteStatListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPRouteStatListResponseMeta struct {
	DataTime   string                                `json:"data_time,required"`
	QueryTime  string                                `json:"query_time,required"`
	TotalPeers int64                                 `json:"total_peers,required"`
	JSON       radarBGPRouteStatListResponseMetaJSON `json:"-"`
}

// radarBGPRouteStatListResponseMetaJSON contains the JSON metadata for the struct
// [RadarBGPRouteStatListResponseMeta]
type radarBGPRouteStatListResponseMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalPeers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPRouteStatListResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPRouteStatListResponseStats struct {
	DistinctOrigins      int64                                  `json:"distinct_origins,required"`
	DistinctOriginsIPV4  int64                                  `json:"distinct_origins_ipv4,required"`
	DistinctOriginsIPV6  int64                                  `json:"distinct_origins_ipv6,required"`
	DistinctPrefixes     int64                                  `json:"distinct_prefixes,required"`
	DistinctPrefixesIPV4 int64                                  `json:"distinct_prefixes_ipv4,required"`
	DistinctPrefixesIPV6 int64                                  `json:"distinct_prefixes_ipv6,required"`
	RoutesInvalid        int64                                  `json:"routes_invalid,required"`
	RoutesInvalidIPV4    int64                                  `json:"routes_invalid_ipv4,required"`
	RoutesInvalidIPV6    int64                                  `json:"routes_invalid_ipv6,required"`
	RoutesTotal          int64                                  `json:"routes_total,required"`
	RoutesTotalIPV4      int64                                  `json:"routes_total_ipv4,required"`
	RoutesTotalIPV6      int64                                  `json:"routes_total_ipv6,required"`
	RoutesUnknown        int64                                  `json:"routes_unknown,required"`
	RoutesUnknownIPV4    int64                                  `json:"routes_unknown_ipv4,required"`
	RoutesUnknownIPV6    int64                                  `json:"routes_unknown_ipv6,required"`
	RoutesValid          int64                                  `json:"routes_valid,required"`
	RoutesValidIPV4      int64                                  `json:"routes_valid_ipv4,required"`
	RoutesValidIPV6      int64                                  `json:"routes_valid_ipv6,required"`
	JSON                 radarBGPRouteStatListResponseStatsJSON `json:"-"`
}

// radarBGPRouteStatListResponseStatsJSON contains the JSON metadata for the struct
// [RadarBGPRouteStatListResponseStats]
type radarBGPRouteStatListResponseStatsJSON struct {
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

func (r *RadarBGPRouteStatListResponseStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPRouteStatListParams struct {
	// Single ASN as integer.
	Asn param.Field[int64] `query:"asn"`
	// Format results are returned in.
	Format param.Field[RadarBGPRouteStatListParamsFormat] `query:"format"`
	// Location Alpha2 code.
	Location param.Field[string] `query:"location"`
}

// URLQuery serializes [RadarBGPRouteStatListParams]'s query parameters as
// `url.Values`.
func (r RadarBGPRouteStatListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarBGPRouteStatListParamsFormat string

const (
	RadarBGPRouteStatListParamsFormatJson RadarBGPRouteStatListParamsFormat = "JSON"
	RadarBGPRouteStatListParamsFormatCsv  RadarBGPRouteStatListParamsFormat = "CSV"
)

type RadarBGPRouteStatListResponseEnvelope struct {
	Result  RadarBGPRouteStatListResponse             `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    radarBGPRouteStatListResponseEnvelopeJSON `json:"-"`
}

// radarBGPRouteStatListResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarBGPRouteStatListResponseEnvelope]
type radarBGPRouteStatListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPRouteStatListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
