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

// RadarBgpTopAsePrefixService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarBgpTopAsePrefixService]
// method instead.
type RadarBgpTopAsePrefixService struct {
	Options []option.RequestOption
}

// NewRadarBgpTopAsePrefixService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarBgpTopAsePrefixService(opts ...option.RequestOption) (r *RadarBgpTopAsePrefixService) {
	r = &RadarBgpTopAsePrefixService{}
	r.Options = opts
	return
}

// Get the full list of autonomous systems on the global routing table ordered by
// announced prefixes count. The data comes from public BGP MRT data archives and
// updates every 2 hours.
func (r *RadarBgpTopAsePrefixService) List(ctx context.Context, query RadarBgpTopAsePrefixListParams, opts ...option.RequestOption) (res *RadarBgpTopAsePrefixListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/bgp/top/ases/prefixes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarBgpTopAsePrefixListResponse struct {
	Result  RadarBgpTopAsePrefixListResponseResult `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    radarBgpTopAsePrefixListResponseJSON   `json:"-"`
}

// radarBgpTopAsePrefixListResponseJSON contains the JSON metadata for the struct
// [RadarBgpTopAsePrefixListResponse]
type radarBgpTopAsePrefixListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTopAsePrefixListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpTopAsePrefixListResponseResult struct {
	ASNs []RadarBgpTopAsePrefixListResponseResultASN `json:"asns,required"`
	Meta RadarBgpTopAsePrefixListResponseResultMeta  `json:"meta,required"`
	JSON radarBgpTopAsePrefixListResponseResultJSON  `json:"-"`
}

// radarBgpTopAsePrefixListResponseResultJSON contains the JSON metadata for the
// struct [RadarBgpTopAsePrefixListResponseResult]
type radarBgpTopAsePrefixListResponseResultJSON struct {
	ASNs        apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTopAsePrefixListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpTopAsePrefixListResponseResultASN struct {
	ASN       int64                                         `json:"asn,required"`
	Country   string                                        `json:"country,required"`
	Name      string                                        `json:"name,required"`
	PfxsCount int64                                         `json:"pfxs_count,required"`
	JSON      radarBgpTopAsePrefixListResponseResultASNJSON `json:"-"`
}

// radarBgpTopAsePrefixListResponseResultASNJSON contains the JSON metadata for the
// struct [RadarBgpTopAsePrefixListResponseResultASN]
type radarBgpTopAsePrefixListResponseResultASNJSON struct {
	ASN         apijson.Field
	Country     apijson.Field
	Name        apijson.Field
	PfxsCount   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTopAsePrefixListResponseResultASN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpTopAsePrefixListResponseResultMeta struct {
	DataTime   string                                         `json:"data_time,required"`
	QueryTime  string                                         `json:"query_time,required"`
	TotalPeers int64                                          `json:"total_peers,required"`
	JSON       radarBgpTopAsePrefixListResponseResultMetaJSON `json:"-"`
}

// radarBgpTopAsePrefixListResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarBgpTopAsePrefixListResponseResultMeta]
type radarBgpTopAsePrefixListResponseResultMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalPeers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTopAsePrefixListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpTopAsePrefixListParams struct {
	// Alpha-2 country code.
	Country param.Field[string] `query:"country"`
	// Format results are returned in.
	Format param.Field[RadarBgpTopAsePrefixListParamsFormat] `query:"format"`
	// Maximum number of ASes to return
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [RadarBgpTopAsePrefixListParams]'s query parameters as
// `url.Values`.
func (r RadarBgpTopAsePrefixListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarBgpTopAsePrefixListParamsFormat string

const (
	RadarBgpTopAsePrefixListParamsFormatJson RadarBgpTopAsePrefixListParamsFormat = "JSON"
	RadarBgpTopAsePrefixListParamsFormatCsv  RadarBgpTopAsePrefixListParamsFormat = "CSV"
)
