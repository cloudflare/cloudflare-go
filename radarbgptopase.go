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

// RadarBGPTopAseService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarBGPTopAseService] method
// instead.
type RadarBGPTopAseService struct {
	Options []option.RequestOption
}

// NewRadarBGPTopAseService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarBGPTopAseService(opts ...option.RequestOption) (r *RadarBGPTopAseService) {
	r = &RadarBGPTopAseService{}
	r.Options = opts
	return
}

// Get the full list of autonomous systems on the global routing table ordered by
// announced prefixes count. The data comes from public BGP MRT data archives and
// updates every 2 hours.
func (r *RadarBGPTopAseService) Prefixes(ctx context.Context, query RadarBGPTopAsePrefixesParams, opts ...option.RequestOption) (res *RadarBGPTopAsePrefixesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarBGPTopAsePrefixesResponseEnvelope
	path := "radar/bgp/top/ases/prefixes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarBGPTopAsePrefixesResponse struct {
	Asns []RadarBGPTopAsePrefixesResponseAsn `json:"asns,required"`
	Meta RadarBGPTopAsePrefixesResponseMeta  `json:"meta,required"`
	JSON radarBGPTopAsePrefixesResponseJSON  `json:"-"`
}

// radarBGPTopAsePrefixesResponseJSON contains the JSON metadata for the struct
// [RadarBGPTopAsePrefixesResponse]
type radarBGPTopAsePrefixesResponseJSON struct {
	Asns        apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopAsePrefixesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPTopAsePrefixesResponseAsn struct {
	Asn       int64                                 `json:"asn,required"`
	Country   string                                `json:"country,required"`
	Name      string                                `json:"name,required"`
	PfxsCount int64                                 `json:"pfxs_count,required"`
	JSON      radarBGPTopAsePrefixesResponseAsnJSON `json:"-"`
}

// radarBGPTopAsePrefixesResponseAsnJSON contains the JSON metadata for the struct
// [RadarBGPTopAsePrefixesResponseAsn]
type radarBGPTopAsePrefixesResponseAsnJSON struct {
	Asn         apijson.Field
	Country     apijson.Field
	Name        apijson.Field
	PfxsCount   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopAsePrefixesResponseAsn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPTopAsePrefixesResponseMeta struct {
	DataTime   string                                 `json:"data_time,required"`
	QueryTime  string                                 `json:"query_time,required"`
	TotalPeers int64                                  `json:"total_peers,required"`
	JSON       radarBGPTopAsePrefixesResponseMetaJSON `json:"-"`
}

// radarBGPTopAsePrefixesResponseMetaJSON contains the JSON metadata for the struct
// [RadarBGPTopAsePrefixesResponseMeta]
type radarBGPTopAsePrefixesResponseMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalPeers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopAsePrefixesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPTopAsePrefixesParams struct {
	// Alpha-2 country code.
	Country param.Field[string] `query:"country"`
	// Format results are returned in.
	Format param.Field[RadarBGPTopAsePrefixesParamsFormat] `query:"format"`
	// Maximum number of ASes to return
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [RadarBGPTopAsePrefixesParams]'s query parameters as
// `url.Values`.
func (r RadarBGPTopAsePrefixesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarBGPTopAsePrefixesParamsFormat string

const (
	RadarBGPTopAsePrefixesParamsFormatJson RadarBGPTopAsePrefixesParamsFormat = "JSON"
	RadarBGPTopAsePrefixesParamsFormatCsv  RadarBGPTopAsePrefixesParamsFormat = "CSV"
)

type RadarBGPTopAsePrefixesResponseEnvelope struct {
	Result  RadarBGPTopAsePrefixesResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarBGPTopAsePrefixesResponseEnvelopeJSON `json:"-"`
}

// radarBGPTopAsePrefixesResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarBGPTopAsePrefixesResponseEnvelope]
type radarBGPTopAsePrefixesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopAsePrefixesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
