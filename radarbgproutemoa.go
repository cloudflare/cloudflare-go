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

// RadarBgpRouteMoaService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarBgpRouteMoaService] method
// instead.
type RadarBgpRouteMoaService struct {
	Options []option.RequestOption
}

// NewRadarBgpRouteMoaService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarBgpRouteMoaService(opts ...option.RequestOption) (r *RadarBgpRouteMoaService) {
	r = &RadarBgpRouteMoaService{}
	r.Options = opts
	return
}

// List all Multi-origin AS (MOAS) prefixes on the global routing tables.
func (r *RadarBgpRouteMoaService) List(ctx context.Context, query RadarBgpRouteMoaListParams, opts ...option.RequestOption) (res *RadarBgpRouteMoaListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/bgp/routes/moas"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarBgpRouteMoaListResponse struct {
	Result  RadarBgpRouteMoaListResponseResult `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    radarBgpRouteMoaListResponseJSON   `json:"-"`
}

// radarBgpRouteMoaListResponseJSON contains the JSON metadata for the struct
// [RadarBgpRouteMoaListResponse]
type radarBgpRouteMoaListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpRouteMoaListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpRouteMoaListResponseResult struct {
	Meta RadarBgpRouteMoaListResponseResultMeta  `json:"meta,required"`
	Moas []RadarBgpRouteMoaListResponseResultMoa `json:"moas,required"`
	JSON radarBgpRouteMoaListResponseResultJSON  `json:"-"`
}

// radarBgpRouteMoaListResponseResultJSON contains the JSON metadata for the struct
// [RadarBgpRouteMoaListResponseResult]
type radarBgpRouteMoaListResponseResultJSON struct {
	Meta        apijson.Field
	Moas        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpRouteMoaListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpRouteMoaListResponseResultMeta struct {
	DataTime   string                                     `json:"data_time,required"`
	QueryTime  string                                     `json:"query_time,required"`
	TotalPeers int64                                      `json:"total_peers,required"`
	JSON       radarBgpRouteMoaListResponseResultMetaJSON `json:"-"`
}

// radarBgpRouteMoaListResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarBgpRouteMoaListResponseResultMeta]
type radarBgpRouteMoaListResponseResultMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalPeers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpRouteMoaListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpRouteMoaListResponseResultMoa struct {
	Origins []RadarBgpRouteMoaListResponseResultMoasOrigin `json:"origins,required"`
	Prefix  string                                         `json:"prefix,required"`
	JSON    radarBgpRouteMoaListResponseResultMoaJSON      `json:"-"`
}

// radarBgpRouteMoaListResponseResultMoaJSON contains the JSON metadata for the
// struct [RadarBgpRouteMoaListResponseResultMoa]
type radarBgpRouteMoaListResponseResultMoaJSON struct {
	Origins     apijson.Field
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpRouteMoaListResponseResultMoa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpRouteMoaListResponseResultMoasOrigin struct {
	Origin         int64                                            `json:"origin,required"`
	PeerCount      int64                                            `json:"peer_count,required"`
	RpkiValidation string                                           `json:"rpki_validation,required"`
	JSON           radarBgpRouteMoaListResponseResultMoasOriginJSON `json:"-"`
}

// radarBgpRouteMoaListResponseResultMoasOriginJSON contains the JSON metadata for
// the struct [RadarBgpRouteMoaListResponseResultMoasOrigin]
type radarBgpRouteMoaListResponseResultMoasOriginJSON struct {
	Origin         apijson.Field
	PeerCount      apijson.Field
	RpkiValidation apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarBgpRouteMoaListResponseResultMoasOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpRouteMoaListParams struct {
	// Format results are returned in.
	Format param.Field[RadarBgpRouteMoaListParamsFormat] `query:"format"`
	// Lookup only RPKI invalid MOASes
	InvalidOnly param.Field[bool] `query:"invalid_only"`
	// Lookup MOASes originated by the given ASN
	Origin param.Field[int64] `query:"origin"`
	// Lookup MOASes by prefix
	Prefix param.Field[string] `query:"prefix"`
}

// URLQuery serializes [RadarBgpRouteMoaListParams]'s query parameters as
// `url.Values`.
func (r RadarBgpRouteMoaListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarBgpRouteMoaListParamsFormat string

const (
	RadarBgpRouteMoaListParamsFormatJson RadarBgpRouteMoaListParamsFormat = "JSON"
	RadarBgpRouteMoaListParamsFormatCsv  RadarBgpRouteMoaListParamsFormat = "CSV"
)
