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

// RadarBGPTopAsePrefixService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarBGPTopAsePrefixService]
// method instead.
type RadarBGPTopAsePrefixService struct {
	Options []option.RequestOption
}

// NewRadarBGPTopAsePrefixService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarBGPTopAsePrefixService(opts ...option.RequestOption) (r *RadarBGPTopAsePrefixService) {
	r = &RadarBGPTopAsePrefixService{}
	r.Options = opts
	return
}

// Get the full list of autonomous systems on the global routing table ordered by
// announced prefixes count. The data comes from public BGP MRT data archives and
// updates every 2 hours.
func (r *RadarBGPTopAsePrefixService) List(ctx context.Context, query RadarBGPTopAsePrefixListParams, opts ...option.RequestOption) (res *RadarBGPTopAsePrefixListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarBGPTopAsePrefixListResponseEnvelope
	path := "radar/bgp/top/ases/prefixes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarBGPTopAsePrefixListResponse struct {
	Asns []RadarBGPTopAsePrefixListResponseAsn `json:"asns,required"`
	Meta RadarBGPTopAsePrefixListResponseMeta  `json:"meta,required"`
	JSON radarBGPTopAsePrefixListResponseJSON  `json:"-"`
}

// radarBGPTopAsePrefixListResponseJSON contains the JSON metadata for the struct
// [RadarBGPTopAsePrefixListResponse]
type radarBGPTopAsePrefixListResponseJSON struct {
	Asns        apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopAsePrefixListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPTopAsePrefixListResponseAsn struct {
	Asn       int64                                   `json:"asn,required"`
	Country   string                                  `json:"country,required"`
	Name      string                                  `json:"name,required"`
	PfxsCount int64                                   `json:"pfxs_count,required"`
	JSON      radarBGPTopAsePrefixListResponseAsnJSON `json:"-"`
}

// radarBGPTopAsePrefixListResponseAsnJSON contains the JSON metadata for the
// struct [RadarBGPTopAsePrefixListResponseAsn]
type radarBGPTopAsePrefixListResponseAsnJSON struct {
	Asn         apijson.Field
	Country     apijson.Field
	Name        apijson.Field
	PfxsCount   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopAsePrefixListResponseAsn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPTopAsePrefixListResponseMeta struct {
	DataTime   string                                   `json:"data_time,required"`
	QueryTime  string                                   `json:"query_time,required"`
	TotalPeers int64                                    `json:"total_peers,required"`
	JSON       radarBGPTopAsePrefixListResponseMetaJSON `json:"-"`
}

// radarBGPTopAsePrefixListResponseMetaJSON contains the JSON metadata for the
// struct [RadarBGPTopAsePrefixListResponseMeta]
type radarBGPTopAsePrefixListResponseMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalPeers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopAsePrefixListResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPTopAsePrefixListParams struct {
	// Alpha-2 country code.
	Country param.Field[string] `query:"country"`
	// Format results are returned in.
	Format param.Field[RadarBGPTopAsePrefixListParamsFormat] `query:"format"`
	// Maximum number of ASes to return
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [RadarBGPTopAsePrefixListParams]'s query parameters as
// `url.Values`.
func (r RadarBGPTopAsePrefixListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarBGPTopAsePrefixListParamsFormat string

const (
	RadarBGPTopAsePrefixListParamsFormatJson RadarBGPTopAsePrefixListParamsFormat = "JSON"
	RadarBGPTopAsePrefixListParamsFormatCsv  RadarBGPTopAsePrefixListParamsFormat = "CSV"
)

type RadarBGPTopAsePrefixListResponseEnvelope struct {
	Result  RadarBGPTopAsePrefixListResponse             `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarBGPTopAsePrefixListResponseEnvelopeJSON `json:"-"`
}

// radarBGPTopAsePrefixListResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarBGPTopAsePrefixListResponseEnvelope]
type radarBGPTopAsePrefixListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopAsePrefixListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
