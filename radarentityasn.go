// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// RadarEntityASNService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEntityASNService] method
// instead.
type RadarEntityASNService struct {
	Options []option.RequestOption
}

// NewRadarEntityASNService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarEntityASNService(opts ...option.RequestOption) (r *RadarEntityASNService) {
	r = &RadarEntityASNService{}
	r.Options = opts
	return
}

// Gets a list of autonomous systems (AS).
func (r *RadarEntityASNService) List(ctx context.Context, query RadarEntityASNListParams, opts ...option.RequestOption) (res *RadarEntityASNListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEntityASNListResponseEnvelope
	path := "radar/entities/asns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the requested autonomous system information. A confidence level below `5`
// indicates a low level of confidence in the traffic data - normally this happens
// because Cloudflare has a small amount of traffic from/to this AS). Population
// estimates come from APNIC (refer to https://labs.apnic.net/?p=526).
func (r *RadarEntityASNService) Get(ctx context.Context, asn int64, query RadarEntityASNGetParams, opts ...option.RequestOption) (res *RadarEntityASNGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEntityASNGetResponseEnvelope
	path := fmt.Sprintf("radar/entities/asns/%v", asn)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the requested autonomous system information based on IP address. Population
// estimates come from APNIC (refer to https://labs.apnic.net/?p=526).
func (r *RadarEntityASNService) IP(ctx context.Context, query RadarEntityASNIPParams, opts ...option.RequestOption) (res *RadarEntityAsnipResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEntityAsnipResponseEnvelope
	path := "radar/entities/asns/ip"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get AS-level relationship for given networks.
func (r *RadarEntityASNService) Rel(ctx context.Context, asn int64, query RadarEntityASNRelParams, opts ...option.RequestOption) (res *RadarEntityASNRelResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEntityASNRelResponseEnvelope
	path := fmt.Sprintf("radar/entities/asns/%v/rel", asn)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEntityASNListResponse struct {
	ASNs []RadarEntityASNListResponseASN `json:"asns,required"`
	JSON radarEntityASNListResponseJSON  `json:"-"`
}

// radarEntityASNListResponseJSON contains the JSON metadata for the struct
// [RadarEntityASNListResponse]
type radarEntityASNListResponseJSON struct {
	ASNs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityASNListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityASNListResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEntityASNListResponseASN struct {
	ASN         int64  `json:"asn,required"`
	Country     string `json:"country,required"`
	CountryName string `json:"countryName,required"`
	Name        string `json:"name,required"`
	Aka         string `json:"aka"`
	// Deprecated field. Please use 'aka'.
	NameLong string                            `json:"nameLong"`
	OrgName  string                            `json:"orgName"`
	Website  string                            `json:"website"`
	JSON     radarEntityASNListResponseASNJSON `json:"-"`
}

// radarEntityASNListResponseASNJSON contains the JSON metadata for the struct
// [RadarEntityASNListResponseASN]
type radarEntityASNListResponseASNJSON struct {
	ASN         apijson.Field
	Country     apijson.Field
	CountryName apijson.Field
	Name        apijson.Field
	Aka         apijson.Field
	NameLong    apijson.Field
	OrgName     apijson.Field
	Website     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityASNListResponseASN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityASNListResponseASNJSON) RawJSON() string {
	return r.raw
}

type RadarEntityASNGetResponse struct {
	ASN  RadarEntityASNGetResponseASN  `json:"asn,required"`
	JSON radarEntityASNGetResponseJSON `json:"-"`
}

// radarEntityASNGetResponseJSON contains the JSON metadata for the struct
// [RadarEntityASNGetResponse]
type radarEntityASNGetResponseJSON struct {
	ASN         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityASNGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityASNGetResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEntityASNGetResponseASN struct {
	ASN             int64                                      `json:"asn,required"`
	ConfidenceLevel int64                                      `json:"confidenceLevel,required"`
	Country         string                                     `json:"country,required"`
	CountryName     string                                     `json:"countryName,required"`
	EstimatedUsers  RadarEntityASNGetResponseASNEstimatedUsers `json:"estimatedUsers,required"`
	Name            string                                     `json:"name,required"`
	OrgName         string                                     `json:"orgName,required"`
	Related         []RadarEntityASNGetResponseASNRelated      `json:"related,required"`
	// Regional Internet Registry
	Source  string `json:"source,required"`
	Website string `json:"website,required"`
	Aka     string `json:"aka"`
	// Deprecated field. Please use 'aka'.
	NameLong string                           `json:"nameLong"`
	JSON     radarEntityASNGetResponseASNJSON `json:"-"`
}

// radarEntityASNGetResponseASNJSON contains the JSON metadata for the struct
// [RadarEntityASNGetResponseASN]
type radarEntityASNGetResponseASNJSON struct {
	ASN             apijson.Field
	ConfidenceLevel apijson.Field
	Country         apijson.Field
	CountryName     apijson.Field
	EstimatedUsers  apijson.Field
	Name            apijson.Field
	OrgName         apijson.Field
	Related         apijson.Field
	Source          apijson.Field
	Website         apijson.Field
	Aka             apijson.Field
	NameLong        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarEntityASNGetResponseASN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityASNGetResponseASNJSON) RawJSON() string {
	return r.raw
}

type RadarEntityASNGetResponseASNEstimatedUsers struct {
	Locations []RadarEntityASNGetResponseASNEstimatedUsersLocation `json:"locations,required"`
	// Total estimated users
	EstimatedUsers int64                                          `json:"estimatedUsers"`
	JSON           radarEntityASNGetResponseASNEstimatedUsersJSON `json:"-"`
}

// radarEntityASNGetResponseASNEstimatedUsersJSON contains the JSON metadata for
// the struct [RadarEntityASNGetResponseASNEstimatedUsers]
type radarEntityASNGetResponseASNEstimatedUsersJSON struct {
	Locations      apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEntityASNGetResponseASNEstimatedUsers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityASNGetResponseASNEstimatedUsersJSON) RawJSON() string {
	return r.raw
}

type RadarEntityASNGetResponseASNEstimatedUsersLocation struct {
	LocationAlpha2 string `json:"locationAlpha2,required"`
	LocationName   string `json:"locationName,required"`
	// Estimated users per location
	EstimatedUsers int64                                                  `json:"estimatedUsers"`
	JSON           radarEntityASNGetResponseASNEstimatedUsersLocationJSON `json:"-"`
}

// radarEntityASNGetResponseASNEstimatedUsersLocationJSON contains the JSON
// metadata for the struct [RadarEntityASNGetResponseASNEstimatedUsersLocation]
type radarEntityASNGetResponseASNEstimatedUsersLocationJSON struct {
	LocationAlpha2 apijson.Field
	LocationName   apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEntityASNGetResponseASNEstimatedUsersLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityASNGetResponseASNEstimatedUsersLocationJSON) RawJSON() string {
	return r.raw
}

type RadarEntityASNGetResponseASNRelated struct {
	ASN  int64  `json:"asn,required"`
	Name string `json:"name,required"`
	Aka  string `json:"aka"`
	// Total estimated users
	EstimatedUsers int64                                   `json:"estimatedUsers"`
	JSON           radarEntityASNGetResponseASNRelatedJSON `json:"-"`
}

// radarEntityASNGetResponseASNRelatedJSON contains the JSON metadata for the
// struct [RadarEntityASNGetResponseASNRelated]
type radarEntityASNGetResponseASNRelatedJSON struct {
	ASN            apijson.Field
	Name           apijson.Field
	Aka            apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEntityASNGetResponseASNRelated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityASNGetResponseASNRelatedJSON) RawJSON() string {
	return r.raw
}

type RadarEntityAsnipResponse struct {
	ASN  RadarEntityAsnipResponseASN  `json:"asn,required"`
	JSON radarEntityAsnipResponseJSON `json:"-"`
}

// radarEntityAsnipResponseJSON contains the JSON metadata for the struct
// [RadarEntityAsnipResponse]
type radarEntityAsnipResponseJSON struct {
	ASN         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityAsnipResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityAsnipResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEntityAsnipResponseASN struct {
	ASN            int64                                     `json:"asn,required"`
	Country        string                                    `json:"country,required"`
	CountryName    string                                    `json:"countryName,required"`
	EstimatedUsers RadarEntityAsnipResponseASNEstimatedUsers `json:"estimatedUsers,required"`
	Name           string                                    `json:"name,required"`
	OrgName        string                                    `json:"orgName,required"`
	Related        []RadarEntityAsnipResponseASNRelated      `json:"related,required"`
	// Regional Internet Registry
	Source  string `json:"source,required"`
	Website string `json:"website,required"`
	Aka     string `json:"aka"`
	// Deprecated field. Please use 'aka'.
	NameLong string                          `json:"nameLong"`
	JSON     radarEntityAsnipResponseASNJSON `json:"-"`
}

// radarEntityAsnipResponseASNJSON contains the JSON metadata for the struct
// [RadarEntityAsnipResponseASN]
type radarEntityAsnipResponseASNJSON struct {
	ASN            apijson.Field
	Country        apijson.Field
	CountryName    apijson.Field
	EstimatedUsers apijson.Field
	Name           apijson.Field
	OrgName        apijson.Field
	Related        apijson.Field
	Source         apijson.Field
	Website        apijson.Field
	Aka            apijson.Field
	NameLong       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEntityAsnipResponseASN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityAsnipResponseASNJSON) RawJSON() string {
	return r.raw
}

type RadarEntityAsnipResponseASNEstimatedUsers struct {
	Locations []RadarEntityAsnipResponseASNEstimatedUsersLocation `json:"locations,required"`
	// Total estimated users
	EstimatedUsers int64                                         `json:"estimatedUsers"`
	JSON           radarEntityAsnipResponseASNEstimatedUsersJSON `json:"-"`
}

// radarEntityAsnipResponseASNEstimatedUsersJSON contains the JSON metadata for the
// struct [RadarEntityAsnipResponseASNEstimatedUsers]
type radarEntityAsnipResponseASNEstimatedUsersJSON struct {
	Locations      apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEntityAsnipResponseASNEstimatedUsers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityAsnipResponseASNEstimatedUsersJSON) RawJSON() string {
	return r.raw
}

type RadarEntityAsnipResponseASNEstimatedUsersLocation struct {
	LocationAlpha2 string `json:"locationAlpha2,required"`
	LocationName   string `json:"locationName,required"`
	// Estimated users per location
	EstimatedUsers int64                                                 `json:"estimatedUsers"`
	JSON           radarEntityAsnipResponseASNEstimatedUsersLocationJSON `json:"-"`
}

// radarEntityAsnipResponseASNEstimatedUsersLocationJSON contains the JSON metadata
// for the struct [RadarEntityAsnipResponseASNEstimatedUsersLocation]
type radarEntityAsnipResponseASNEstimatedUsersLocationJSON struct {
	LocationAlpha2 apijson.Field
	LocationName   apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEntityAsnipResponseASNEstimatedUsersLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityAsnipResponseASNEstimatedUsersLocationJSON) RawJSON() string {
	return r.raw
}

type RadarEntityAsnipResponseASNRelated struct {
	ASN  int64  `json:"asn,required"`
	Name string `json:"name,required"`
	Aka  string `json:"aka"`
	// Total estimated users
	EstimatedUsers int64                                  `json:"estimatedUsers"`
	JSON           radarEntityAsnipResponseASNRelatedJSON `json:"-"`
}

// radarEntityAsnipResponseASNRelatedJSON contains the JSON metadata for the struct
// [RadarEntityAsnipResponseASNRelated]
type radarEntityAsnipResponseASNRelatedJSON struct {
	ASN            apijson.Field
	Name           apijson.Field
	Aka            apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEntityAsnipResponseASNRelated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityAsnipResponseASNRelatedJSON) RawJSON() string {
	return r.raw
}

type RadarEntityASNRelResponse struct {
	Meta RadarEntityASNRelResponseMeta  `json:"meta,required"`
	Rels []RadarEntityASNRelResponseRel `json:"rels,required"`
	JSON radarEntityASNRelResponseJSON  `json:"-"`
}

// radarEntityASNRelResponseJSON contains the JSON metadata for the struct
// [RadarEntityASNRelResponse]
type radarEntityASNRelResponseJSON struct {
	Meta        apijson.Field
	Rels        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityASNRelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityASNRelResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEntityASNRelResponseMeta struct {
	DataTime   string                            `json:"data_time,required"`
	QueryTime  string                            `json:"query_time,required"`
	TotalPeers int64                             `json:"total_peers,required"`
	JSON       radarEntityASNRelResponseMetaJSON `json:"-"`
}

// radarEntityASNRelResponseMetaJSON contains the JSON metadata for the struct
// [RadarEntityASNRelResponseMeta]
type radarEntityASNRelResponseMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalPeers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityASNRelResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityASNRelResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEntityASNRelResponseRel struct {
	Asn1        int64                            `json:"asn1,required"`
	Asn1Country string                           `json:"asn1_country,required"`
	Asn1Name    string                           `json:"asn1_name,required"`
	Asn2        int64                            `json:"asn2,required"`
	Asn2Country string                           `json:"asn2_country,required"`
	Asn2Name    string                           `json:"asn2_name,required"`
	Rel         string                           `json:"rel,required"`
	JSON        radarEntityASNRelResponseRelJSON `json:"-"`
}

// radarEntityASNRelResponseRelJSON contains the JSON metadata for the struct
// [RadarEntityASNRelResponseRel]
type radarEntityASNRelResponseRelJSON struct {
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

func (r *RadarEntityASNRelResponseRel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityASNRelResponseRelJSON) RawJSON() string {
	return r.raw
}

type RadarEntityASNListParams struct {
	// Comma separated list of ASNs.
	ASN param.Field[string] `query:"asn"`
	// Format results are returned in.
	Format param.Field[RadarEntityASNListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Location Alpha2 to filter results.
	Location param.Field[string] `query:"location"`
	// Number of objects to skip before grabbing results.
	Offset param.Field[int64] `query:"offset"`
	// Order asn list.
	OrderBy param.Field[RadarEntityASNListParamsOrderBy] `query:"orderBy"`
}

// URLQuery serializes [RadarEntityASNListParams]'s query parameters as
// `url.Values`.
func (r RadarEntityASNListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarEntityASNListParamsFormat string

const (
	RadarEntityASNListParamsFormatJson RadarEntityASNListParamsFormat = "JSON"
	RadarEntityASNListParamsFormatCsv  RadarEntityASNListParamsFormat = "CSV"
)

// Order asn list.
type RadarEntityASNListParamsOrderBy string

const (
	RadarEntityASNListParamsOrderByASN        RadarEntityASNListParamsOrderBy = "ASN"
	RadarEntityASNListParamsOrderByPopulation RadarEntityASNListParamsOrderBy = "POPULATION"
)

type RadarEntityASNListResponseEnvelope struct {
	Result  RadarEntityASNListResponse             `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    radarEntityASNListResponseEnvelopeJSON `json:"-"`
}

// radarEntityASNListResponseEnvelopeJSON contains the JSON metadata for the struct
// [RadarEntityASNListResponseEnvelope]
type radarEntityASNListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityASNListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityASNListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarEntityASNGetParams struct {
	// Format results are returned in.
	Format param.Field[RadarEntityASNGetParamsFormat] `query:"format"`
}

// URLQuery serializes [RadarEntityASNGetParams]'s query parameters as
// `url.Values`.
func (r RadarEntityASNGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarEntityASNGetParamsFormat string

const (
	RadarEntityASNGetParamsFormatJson RadarEntityASNGetParamsFormat = "JSON"
	RadarEntityASNGetParamsFormatCsv  RadarEntityASNGetParamsFormat = "CSV"
)

type RadarEntityASNGetResponseEnvelope struct {
	Result  RadarEntityASNGetResponse             `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    radarEntityASNGetResponseEnvelopeJSON `json:"-"`
}

// radarEntityASNGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RadarEntityASNGetResponseEnvelope]
type radarEntityASNGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityASNGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityASNGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarEntityASNIPParams struct {
	// IP address.
	IP param.Field[string] `query:"ip,required"`
	// Format results are returned in.
	Format param.Field[RadarEntityAsnipParamsFormat] `query:"format"`
}

// URLQuery serializes [RadarEntityASNIPParams]'s query parameters as `url.Values`.
func (r RadarEntityASNIPParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarEntityAsnipParamsFormat string

const (
	RadarEntityAsnipParamsFormatJson RadarEntityAsnipParamsFormat = "JSON"
	RadarEntityAsnipParamsFormatCsv  RadarEntityAsnipParamsFormat = "CSV"
)

type RadarEntityAsnipResponseEnvelope struct {
	Result  RadarEntityAsnipResponse             `json:"result,required"`
	Success bool                                 `json:"success,required"`
	JSON    radarEntityAsnipResponseEnvelopeJSON `json:"-"`
}

// radarEntityAsnipResponseEnvelopeJSON contains the JSON metadata for the struct
// [RadarEntityAsnipResponseEnvelope]
type radarEntityAsnipResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityAsnipResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityAsnipResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarEntityASNRelParams struct {
	// Get the AS relationship of ASN2 with respect to the given ASN
	Asn2 param.Field[int64] `query:"asn2"`
	// Format results are returned in.
	Format param.Field[RadarEntityASNRelParamsFormat] `query:"format"`
}

// URLQuery serializes [RadarEntityASNRelParams]'s query parameters as
// `url.Values`.
func (r RadarEntityASNRelParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarEntityASNRelParamsFormat string

const (
	RadarEntityASNRelParamsFormatJson RadarEntityASNRelParamsFormat = "JSON"
	RadarEntityASNRelParamsFormatCsv  RadarEntityASNRelParamsFormat = "CSV"
)

type RadarEntityASNRelResponseEnvelope struct {
	Result  RadarEntityASNRelResponse             `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    radarEntityASNRelResponseEnvelopeJSON `json:"-"`
}

// radarEntityASNRelResponseEnvelopeJSON contains the JSON metadata for the struct
// [RadarEntityASNRelResponseEnvelope]
type radarEntityASNRelResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityASNRelResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityASNRelResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
