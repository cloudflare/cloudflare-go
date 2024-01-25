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

// RadarBgpRoutePfx2aService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarBgpRoutePfx2aService] method
// instead.
type RadarBgpRoutePfx2aService struct {
	Options []option.RequestOption
}

// NewRadarBgpRoutePfx2aService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarBgpRoutePfx2aService(opts ...option.RequestOption) (r *RadarBgpRoutePfx2aService) {
	r = &RadarBgpRoutePfx2aService{}
	r.Options = opts
	return
}

// Lookup prefix-to-origin mapping on global routing tables.
func (r *RadarBgpRoutePfx2aService) List(ctx context.Context, query RadarBgpRoutePfx2aListParams, opts ...option.RequestOption) (res *RadarBgpRoutePfx2aListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/bgp/routes/pfx2as"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarBgpRoutePfx2aListResponse struct {
	Result  RadarBgpRoutePfx2aListResponseResult `json:"result,required"`
	Success bool                                 `json:"success,required"`
	JSON    radarBgpRoutePfx2aListResponseJSON   `json:"-"`
}

// radarBgpRoutePfx2aListResponseJSON contains the JSON metadata for the struct
// [RadarBgpRoutePfx2aListResponse]
type radarBgpRoutePfx2aListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpRoutePfx2aListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpRoutePfx2aListResponseResult struct {
	Meta          RadarBgpRoutePfx2aListResponseResultMeta           `json:"meta,required"`
	PrefixOrigins []RadarBgpRoutePfx2aListResponseResultPrefixOrigin `json:"prefix_origins,required"`
	JSON          radarBgpRoutePfx2aListResponseResultJSON           `json:"-"`
}

// radarBgpRoutePfx2aListResponseResultJSON contains the JSON metadata for the
// struct [RadarBgpRoutePfx2aListResponseResult]
type radarBgpRoutePfx2aListResponseResultJSON struct {
	Meta          apijson.Field
	PrefixOrigins apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RadarBgpRoutePfx2aListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpRoutePfx2aListResponseResultMeta struct {
	DataTime   string                                       `json:"data_time,required"`
	QueryTime  string                                       `json:"query_time,required"`
	TotalPeers int64                                        `json:"total_peers,required"`
	JSON       radarBgpRoutePfx2aListResponseResultMetaJSON `json:"-"`
}

// radarBgpRoutePfx2aListResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarBgpRoutePfx2aListResponseResultMeta]
type radarBgpRoutePfx2aListResponseResultMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalPeers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpRoutePfx2aListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpRoutePfx2aListResponseResultPrefixOrigin struct {
	Origin         int64                                                `json:"origin,required"`
	PeerCount      int64                                                `json:"peer_count,required"`
	Prefix         string                                               `json:"prefix,required"`
	RpkiValidation string                                               `json:"rpki_validation,required"`
	JSON           radarBgpRoutePfx2aListResponseResultPrefixOriginJSON `json:"-"`
}

// radarBgpRoutePfx2aListResponseResultPrefixOriginJSON contains the JSON metadata
// for the struct [RadarBgpRoutePfx2aListResponseResultPrefixOrigin]
type radarBgpRoutePfx2aListResponseResultPrefixOriginJSON struct {
	Origin         apijson.Field
	PeerCount      apijson.Field
	Prefix         apijson.Field
	RpkiValidation apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarBgpRoutePfx2aListResponseResultPrefixOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpRoutePfx2aListParams struct {
	// Format results are returned in.
	Format param.Field[RadarBgpRoutePfx2aListParamsFormat] `query:"format"`
	// Lookup prefixes originated by the given ASN
	Origin param.Field[int64] `query:"origin"`
	// Lookup origins of the given prefix
	Prefix param.Field[string] `query:"prefix"`
	// Return only results with matching rpki status: valid, invalid or unknown
	RpkiStatus param.Field[RadarBgpRoutePfx2aListParamsRpkiStatus] `query:"rpkiStatus"`
}

// URLQuery serializes [RadarBgpRoutePfx2aListParams]'s query parameters as
// `url.Values`.
func (r RadarBgpRoutePfx2aListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarBgpRoutePfx2aListParamsFormat string

const (
	RadarBgpRoutePfx2aListParamsFormatJson RadarBgpRoutePfx2aListParamsFormat = "JSON"
	RadarBgpRoutePfx2aListParamsFormatCsv  RadarBgpRoutePfx2aListParamsFormat = "CSV"
)

// Return only results with matching rpki status: valid, invalid or unknown
type RadarBgpRoutePfx2aListParamsRpkiStatus string

const (
	RadarBgpRoutePfx2aListParamsRpkiStatusValid   RadarBgpRoutePfx2aListParamsRpkiStatus = "VALID"
	RadarBgpRoutePfx2aListParamsRpkiStatusInvalid RadarBgpRoutePfx2aListParamsRpkiStatus = "INVALID"
	RadarBgpRoutePfx2aListParamsRpkiStatusUnknown RadarBgpRoutePfx2aListParamsRpkiStatus = "UNKNOWN"
)
