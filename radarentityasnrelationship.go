// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarEntityASNRelationshipService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarEntityASNRelationshipService] method instead.
type RadarEntityASNRelationshipService struct {
	Options []option.RequestOption
}

// NewRadarEntityASNRelationshipService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEntityASNRelationshipService(opts ...option.RequestOption) (r *RadarEntityASNRelationshipService) {
	r = &RadarEntityASNRelationshipService{}
	r.Options = opts
	return
}

// Get AS-level relationship for given networks.
func (r *RadarEntityASNRelationshipService) List(ctx context.Context, asn int64, query RadarEntityASNRelationshipListParams, opts ...option.RequestOption) (res *RadarEntityASNRelationshipListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/entities/asns/%v/rel", asn)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEntityASNRelationshipListResponse struct {
	Result  RadarEntityASNRelationshipListResponseResult `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarEntityASNRelationshipListResponseJSON   `json:"-"`
}

// radarEntityASNRelationshipListResponseJSON contains the JSON metadata for the
// struct [RadarEntityASNRelationshipListResponse]
type radarEntityASNRelationshipListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityASNRelationshipListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityASNRelationshipListResponseResult struct {
	Meta RadarEntityASNRelationshipListResponseResultMeta  `json:"meta,required"`
	Rels []RadarEntityASNRelationshipListResponseResultRel `json:"rels,required"`
	JSON radarEntityASNRelationshipListResponseResultJSON  `json:"-"`
}

// radarEntityASNRelationshipListResponseResultJSON contains the JSON metadata for
// the struct [RadarEntityASNRelationshipListResponseResult]
type radarEntityASNRelationshipListResponseResultJSON struct {
	Meta        apijson.Field
	Rels        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityASNRelationshipListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityASNRelationshipListResponseResultMeta struct {
	DataTime   string                                               `json:"data_time,required"`
	QueryTime  string                                               `json:"query_time,required"`
	TotalPeers int64                                                `json:"total_peers,required"`
	JSON       radarEntityASNRelationshipListResponseResultMetaJSON `json:"-"`
}

// radarEntityASNRelationshipListResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarEntityASNRelationshipListResponseResultMeta]
type radarEntityASNRelationshipListResponseResultMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalPeers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityASNRelationshipListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityASNRelationshipListResponseResultRel struct {
	Asn1        int64                                               `json:"asn1,required"`
	Asn1Country string                                              `json:"asn1_country,required"`
	Asn1Name    string                                              `json:"asn1_name,required"`
	Asn2        int64                                               `json:"asn2,required"`
	Asn2Country string                                              `json:"asn2_country,required"`
	Asn2Name    string                                              `json:"asn2_name,required"`
	Rel         string                                              `json:"rel,required"`
	JSON        radarEntityASNRelationshipListResponseResultRelJSON `json:"-"`
}

// radarEntityASNRelationshipListResponseResultRelJSON contains the JSON metadata
// for the struct [RadarEntityASNRelationshipListResponseResultRel]
type radarEntityASNRelationshipListResponseResultRelJSON struct {
	Asn1        apijson.Field
	Asn1Country apijson.Field
	Asn1Name    apijson.Field
	Asn2        apijson.Field
	Asn2Country apijson.Field
	Asn2Name    apijson.Field
	Rel         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityASNRelationshipListResponseResultRel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityASNRelationshipListParams struct {
	// Get the AS relationship of ASN2 with respect to the given ASN
	Asn2 param.Field[int64] `query:"asn2"`
	// Format results are returned in.
	Format param.Field[RadarEntityASNRelationshipListParamsFormat] `query:"format"`
}

// URLQuery serializes [RadarEntityASNRelationshipListParams]'s query parameters as
// `url.Values`.
func (r RadarEntityASNRelationshipListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarEntityASNRelationshipListParamsFormat string

const (
	RadarEntityASNRelationshipListParamsFormatJson RadarEntityASNRelationshipListParamsFormat = "JSON"
	RadarEntityASNRelationshipListParamsFormatCsv  RadarEntityASNRelationshipListParamsFormat = "CSV"
)
