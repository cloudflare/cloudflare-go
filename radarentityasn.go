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

// RadarEntityAsnService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEntityAsnService] method
// instead.
type RadarEntityAsnService struct {
	Options []option.RequestOption
}

// NewRadarEntityAsnService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarEntityAsnService(opts ...option.RequestOption) (r *RadarEntityAsnService) {
	r = &RadarEntityAsnService{}
	r.Options = opts
	return
}

// Get AS-level relationship for given networks.
func (r *RadarEntityAsnService) Rel(ctx context.Context, asn int64, query RadarEntityAsnRelParams, opts ...option.RequestOption) (res *RadarEntityAsnRelResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEntityAsnRelResponseEnvelope
	path := fmt.Sprintf("radar/entities/asns/%v/rel", asn)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEntityAsnRelResponse struct {
	Meta RadarEntityAsnRelResponseMeta  `json:"meta,required"`
	Rels []RadarEntityAsnRelResponseRel `json:"rels,required"`
	JSON radarEntityAsnRelResponseJSON  `json:"-"`
}

// radarEntityAsnRelResponseJSON contains the JSON metadata for the struct
// [RadarEntityAsnRelResponse]
type radarEntityAsnRelResponseJSON struct {
	Meta        apijson.Field
	Rels        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityAsnRelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityAsnRelResponseMeta struct {
	DataTime   string                            `json:"data_time,required"`
	QueryTime  string                            `json:"query_time,required"`
	TotalPeers int64                             `json:"total_peers,required"`
	JSON       radarEntityAsnRelResponseMetaJSON `json:"-"`
}

// radarEntityAsnRelResponseMetaJSON contains the JSON metadata for the struct
// [RadarEntityAsnRelResponseMeta]
type radarEntityAsnRelResponseMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalPeers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityAsnRelResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityAsnRelResponseRel struct {
	Asn1        int64                            `json:"asn1,required"`
	Asn1Country string                           `json:"asn1_country,required"`
	Asn1Name    string                           `json:"asn1_name,required"`
	Asn2        int64                            `json:"asn2,required"`
	Asn2Country string                           `json:"asn2_country,required"`
	Asn2Name    string                           `json:"asn2_name,required"`
	Rel         string                           `json:"rel,required"`
	JSON        radarEntityAsnRelResponseRelJSON `json:"-"`
}

// radarEntityAsnRelResponseRelJSON contains the JSON metadata for the struct
// [RadarEntityAsnRelResponseRel]
type radarEntityAsnRelResponseRelJSON struct {
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

func (r *RadarEntityAsnRelResponseRel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityAsnRelParams struct {
	// Get the AS relationship of ASN2 with respect to the given ASN
	Asn2 param.Field[int64] `query:"asn2"`
	// Format results are returned in.
	Format param.Field[RadarEntityAsnRelParamsFormat] `query:"format"`
}

// URLQuery serializes [RadarEntityAsnRelParams]'s query parameters as
// `url.Values`.
func (r RadarEntityAsnRelParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarEntityAsnRelParamsFormat string

const (
	RadarEntityAsnRelParamsFormatJson RadarEntityAsnRelParamsFormat = "JSON"
	RadarEntityAsnRelParamsFormatCsv  RadarEntityAsnRelParamsFormat = "CSV"
)

type RadarEntityAsnRelResponseEnvelope struct {
	Result  RadarEntityAsnRelResponse             `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    radarEntityAsnRelResponseEnvelopeJSON `json:"-"`
}

// radarEntityAsnRelResponseEnvelopeJSON contains the JSON metadata for the struct
// [RadarEntityAsnRelResponseEnvelope]
type radarEntityAsnRelResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityAsnRelResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
