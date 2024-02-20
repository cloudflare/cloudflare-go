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

// RadarBGPRoutePfx2aService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarBGPRoutePfx2aService] method
// instead.
type RadarBGPRoutePfx2aService struct {
	Options []option.RequestOption
}

// NewRadarBGPRoutePfx2aService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarBGPRoutePfx2aService(opts ...option.RequestOption) (r *RadarBGPRoutePfx2aService) {
	r = &RadarBGPRoutePfx2aService{}
	r.Options = opts
	return
}

// Lookup prefix-to-origin mapping on global routing tables.
func (r *RadarBGPRoutePfx2aService) List(ctx context.Context, query RadarBGPRoutePfx2aListParams, opts ...option.RequestOption) (res *RadarBGPRoutePfx2aListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarBGPRoutePfx2aListResponseEnvelope
	path := "radar/bgp/routes/pfx2as"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarBGPRoutePfx2aListResponse struct {
	Meta          RadarBGPRoutePfx2aListResponseMeta           `json:"meta,required"`
	PrefixOrigins []RadarBGPRoutePfx2aListResponsePrefixOrigin `json:"prefix_origins,required"`
	JSON          radarBGPRoutePfx2aListResponseJSON           `json:"-"`
}

// radarBGPRoutePfx2aListResponseJSON contains the JSON metadata for the struct
// [RadarBGPRoutePfx2aListResponse]
type radarBGPRoutePfx2aListResponseJSON struct {
	Meta          apijson.Field
	PrefixOrigins apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RadarBGPRoutePfx2aListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPRoutePfx2aListResponseMeta struct {
	DataTime   string                                 `json:"data_time,required"`
	QueryTime  string                                 `json:"query_time,required"`
	TotalPeers int64                                  `json:"total_peers,required"`
	JSON       radarBGPRoutePfx2aListResponseMetaJSON `json:"-"`
}

// radarBGPRoutePfx2aListResponseMetaJSON contains the JSON metadata for the struct
// [RadarBGPRoutePfx2aListResponseMeta]
type radarBGPRoutePfx2aListResponseMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalPeers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPRoutePfx2aListResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPRoutePfx2aListResponsePrefixOrigin struct {
	Origin         int64                                          `json:"origin,required"`
	PeerCount      int64                                          `json:"peer_count,required"`
	Prefix         string                                         `json:"prefix,required"`
	RpkiValidation string                                         `json:"rpki_validation,required"`
	JSON           radarBGPRoutePfx2aListResponsePrefixOriginJSON `json:"-"`
}

// radarBGPRoutePfx2aListResponsePrefixOriginJSON contains the JSON metadata for
// the struct [RadarBGPRoutePfx2aListResponsePrefixOrigin]
type radarBGPRoutePfx2aListResponsePrefixOriginJSON struct {
	Origin         apijson.Field
	PeerCount      apijson.Field
	Prefix         apijson.Field
	RpkiValidation apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarBGPRoutePfx2aListResponsePrefixOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPRoutePfx2aListParams struct {
	// Format results are returned in.
	Format param.Field[RadarBGPRoutePfx2aListParamsFormat] `query:"format"`
	// Lookup prefixes originated by the given ASN
	Origin param.Field[int64] `query:"origin"`
	// Lookup origins of the given prefix
	Prefix param.Field[string] `query:"prefix"`
	// Return only results with matching rpki status: valid, invalid or unknown
	RpkiStatus param.Field[RadarBGPRoutePfx2aListParamsRpkiStatus] `query:"rpkiStatus"`
}

// URLQuery serializes [RadarBGPRoutePfx2aListParams]'s query parameters as
// `url.Values`.
func (r RadarBGPRoutePfx2aListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarBGPRoutePfx2aListParamsFormat string

const (
	RadarBGPRoutePfx2aListParamsFormatJson RadarBGPRoutePfx2aListParamsFormat = "JSON"
	RadarBGPRoutePfx2aListParamsFormatCsv  RadarBGPRoutePfx2aListParamsFormat = "CSV"
)

// Return only results with matching rpki status: valid, invalid or unknown
type RadarBGPRoutePfx2aListParamsRpkiStatus string

const (
	RadarBGPRoutePfx2aListParamsRpkiStatusValid   RadarBGPRoutePfx2aListParamsRpkiStatus = "VALID"
	RadarBGPRoutePfx2aListParamsRpkiStatusInvalid RadarBGPRoutePfx2aListParamsRpkiStatus = "INVALID"
	RadarBGPRoutePfx2aListParamsRpkiStatusUnknown RadarBGPRoutePfx2aListParamsRpkiStatus = "UNKNOWN"
)

type RadarBGPRoutePfx2aListResponseEnvelope struct {
	Result  RadarBGPRoutePfx2aListResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarBGPRoutePfx2aListResponseEnvelopeJSON `json:"-"`
}

// radarBGPRoutePfx2aListResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarBGPRoutePfx2aListResponseEnvelope]
type radarBGPRoutePfx2aListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPRoutePfx2aListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
