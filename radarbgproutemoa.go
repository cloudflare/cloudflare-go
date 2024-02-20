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

// RadarBGPRouteMoaService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarBGPRouteMoaService] method
// instead.
type RadarBGPRouteMoaService struct {
	Options []option.RequestOption
}

// NewRadarBGPRouteMoaService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarBGPRouteMoaService(opts ...option.RequestOption) (r *RadarBGPRouteMoaService) {
	r = &RadarBGPRouteMoaService{}
	r.Options = opts
	return
}

// List all Multi-origin AS (MOAS) prefixes on the global routing tables.
func (r *RadarBGPRouteMoaService) List(ctx context.Context, query RadarBGPRouteMoaListParams, opts ...option.RequestOption) (res *RadarBGPRouteMoaListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarBGPRouteMoaListResponseEnvelope
	path := "radar/bgp/routes/moas"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarBGPRouteMoaListResponse struct {
	Meta RadarBGPRouteMoaListResponseMeta  `json:"meta,required"`
	Moas []RadarBGPRouteMoaListResponseMoa `json:"moas,required"`
	JSON radarBGPRouteMoaListResponseJSON  `json:"-"`
}

// radarBGPRouteMoaListResponseJSON contains the JSON metadata for the struct
// [RadarBGPRouteMoaListResponse]
type radarBGPRouteMoaListResponseJSON struct {
	Meta        apijson.Field
	Moas        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPRouteMoaListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPRouteMoaListResponseMeta struct {
	DataTime   string                               `json:"data_time,required"`
	QueryTime  string                               `json:"query_time,required"`
	TotalPeers int64                                `json:"total_peers,required"`
	JSON       radarBGPRouteMoaListResponseMetaJSON `json:"-"`
}

// radarBGPRouteMoaListResponseMetaJSON contains the JSON metadata for the struct
// [RadarBGPRouteMoaListResponseMeta]
type radarBGPRouteMoaListResponseMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalPeers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPRouteMoaListResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPRouteMoaListResponseMoa struct {
	Origins []RadarBGPRouteMoaListResponseMoasOrigin `json:"origins,required"`
	Prefix  string                                   `json:"prefix,required"`
	JSON    radarBGPRouteMoaListResponseMoaJSON      `json:"-"`
}

// radarBGPRouteMoaListResponseMoaJSON contains the JSON metadata for the struct
// [RadarBGPRouteMoaListResponseMoa]
type radarBGPRouteMoaListResponseMoaJSON struct {
	Origins     apijson.Field
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPRouteMoaListResponseMoa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPRouteMoaListResponseMoasOrigin struct {
	Origin         int64                                      `json:"origin,required"`
	PeerCount      int64                                      `json:"peer_count,required"`
	RpkiValidation string                                     `json:"rpki_validation,required"`
	JSON           radarBGPRouteMoaListResponseMoasOriginJSON `json:"-"`
}

// radarBGPRouteMoaListResponseMoasOriginJSON contains the JSON metadata for the
// struct [RadarBGPRouteMoaListResponseMoasOrigin]
type radarBGPRouteMoaListResponseMoasOriginJSON struct {
	Origin         apijson.Field
	PeerCount      apijson.Field
	RpkiValidation apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarBGPRouteMoaListResponseMoasOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPRouteMoaListParams struct {
	// Format results are returned in.
	Format param.Field[RadarBGPRouteMoaListParamsFormat] `query:"format"`
	// Lookup only RPKI invalid MOASes
	InvalidOnly param.Field[bool] `query:"invalid_only"`
	// Lookup MOASes originated by the given ASN
	Origin param.Field[int64] `query:"origin"`
	// Lookup MOASes by prefix
	Prefix param.Field[string] `query:"prefix"`
}

// URLQuery serializes [RadarBGPRouteMoaListParams]'s query parameters as
// `url.Values`.
func (r RadarBGPRouteMoaListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarBGPRouteMoaListParamsFormat string

const (
	RadarBGPRouteMoaListParamsFormatJson RadarBGPRouteMoaListParamsFormat = "JSON"
	RadarBGPRouteMoaListParamsFormatCsv  RadarBGPRouteMoaListParamsFormat = "CSV"
)

type RadarBGPRouteMoaListResponseEnvelope struct {
	Result  RadarBGPRouteMoaListResponse             `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    radarBGPRouteMoaListResponseEnvelopeJSON `json:"-"`
}

// radarBGPRouteMoaListResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarBGPRouteMoaListResponseEnvelope]
type radarBGPRouteMoaListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPRouteMoaListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
