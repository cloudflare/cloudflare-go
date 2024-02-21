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

// Gets a list of autonomous systems (AS).
func (r *RadarEntityAsnService) List(ctx context.Context, query RadarEntityAsnListParams, opts ...option.RequestOption) (res *RadarEntityAsnListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEntityAsnListResponseEnvelope
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
func (r *RadarEntityAsnService) Get(ctx context.Context, asn int64, query RadarEntityAsnGetParams, opts ...option.RequestOption) (res *RadarEntityAsnGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEntityAsnGetResponseEnvelope
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
func (r *RadarEntityAsnService) IP(ctx context.Context, query RadarEntityAsnIPParams, opts ...option.RequestOption) (res *RadarEntityAsnIPResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEntityAsnIPResponseEnvelope
	path := "radar/entities/asns/ip"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
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

type RadarEntityAsnListResponse struct {
	Asns []RadarEntityAsnListResponseAsn `json:"asns,required"`
	JSON radarEntityAsnListResponseJSON  `json:"-"`
}

// radarEntityAsnListResponseJSON contains the JSON metadata for the struct
// [RadarEntityAsnListResponse]
type radarEntityAsnListResponseJSON struct {
	Asns        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityAsnListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityAsnListResponseAsn struct {
	Asn         int64  `json:"asn,required"`
	Country     string `json:"country,required"`
	CountryName string `json:"countryName,required"`
	Name        string `json:"name,required"`
	Aka         string `json:"aka"`
	// Deprecated field. Please use 'aka'.
	NameLong string                            `json:"nameLong"`
	OrgName  string                            `json:"orgName"`
	Website  string                            `json:"website"`
	JSON     radarEntityAsnListResponseAsnJSON `json:"-"`
}

// radarEntityAsnListResponseAsnJSON contains the JSON metadata for the struct
// [RadarEntityAsnListResponseAsn]
type radarEntityAsnListResponseAsnJSON struct {
	Asn         apijson.Field
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

func (r *RadarEntityAsnListResponseAsn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityAsnGetResponse struct {
	Asn  RadarEntityAsnGetResponseAsn  `json:"asn,required"`
	JSON radarEntityAsnGetResponseJSON `json:"-"`
}

// radarEntityAsnGetResponseJSON contains the JSON metadata for the struct
// [RadarEntityAsnGetResponse]
type radarEntityAsnGetResponseJSON struct {
	Asn         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityAsnGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityAsnGetResponseAsn struct {
	Asn             int64                                      `json:"asn,required"`
	ConfidenceLevel int64                                      `json:"confidenceLevel,required"`
	Country         string                                     `json:"country,required"`
	CountryName     string                                     `json:"countryName,required"`
	EstimatedUsers  RadarEntityAsnGetResponseAsnEstimatedUsers `json:"estimatedUsers,required"`
	Name            string                                     `json:"name,required"`
	OrgName         string                                     `json:"orgName,required"`
	Related         []RadarEntityAsnGetResponseAsnRelated      `json:"related,required"`
	// Regional Internet Registry
	Source  string `json:"source,required"`
	Website string `json:"website,required"`
	Aka     string `json:"aka"`
	// Deprecated field. Please use 'aka'.
	NameLong string                           `json:"nameLong"`
	JSON     radarEntityAsnGetResponseAsnJSON `json:"-"`
}

// radarEntityAsnGetResponseAsnJSON contains the JSON metadata for the struct
// [RadarEntityAsnGetResponseAsn]
type radarEntityAsnGetResponseAsnJSON struct {
	Asn             apijson.Field
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

func (r *RadarEntityAsnGetResponseAsn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityAsnGetResponseAsnEstimatedUsers struct {
	Locations []RadarEntityAsnGetResponseAsnEstimatedUsersLocation `json:"locations,required"`
	// Total estimated users
	EstimatedUsers int64                                          `json:"estimatedUsers"`
	JSON           radarEntityAsnGetResponseAsnEstimatedUsersJSON `json:"-"`
}

// radarEntityAsnGetResponseAsnEstimatedUsersJSON contains the JSON metadata for
// the struct [RadarEntityAsnGetResponseAsnEstimatedUsers]
type radarEntityAsnGetResponseAsnEstimatedUsersJSON struct {
	Locations      apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEntityAsnGetResponseAsnEstimatedUsers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityAsnGetResponseAsnEstimatedUsersLocation struct {
	LocationAlpha2 string `json:"locationAlpha2,required"`
	LocationName   string `json:"locationName,required"`
	// Estimated users per location
	EstimatedUsers int64                                                  `json:"estimatedUsers"`
	JSON           radarEntityAsnGetResponseAsnEstimatedUsersLocationJSON `json:"-"`
}

// radarEntityAsnGetResponseAsnEstimatedUsersLocationJSON contains the JSON
// metadata for the struct [RadarEntityAsnGetResponseAsnEstimatedUsersLocation]
type radarEntityAsnGetResponseAsnEstimatedUsersLocationJSON struct {
	LocationAlpha2 apijson.Field
	LocationName   apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEntityAsnGetResponseAsnEstimatedUsersLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityAsnGetResponseAsnRelated struct {
	Asn  int64  `json:"asn,required"`
	Name string `json:"name,required"`
	Aka  string `json:"aka"`
	// Total estimated users
	EstimatedUsers int64                                   `json:"estimatedUsers"`
	JSON           radarEntityAsnGetResponseAsnRelatedJSON `json:"-"`
}

// radarEntityAsnGetResponseAsnRelatedJSON contains the JSON metadata for the
// struct [RadarEntityAsnGetResponseAsnRelated]
type radarEntityAsnGetResponseAsnRelatedJSON struct {
	Asn            apijson.Field
	Name           apijson.Field
	Aka            apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEntityAsnGetResponseAsnRelated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityAsnIPResponse struct {
	Asn  RadarEntityAsnIPResponseAsn  `json:"asn,required"`
	JSON radarEntityAsnIPResponseJSON `json:"-"`
}

// radarEntityAsnIPResponseJSON contains the JSON metadata for the struct
// [RadarEntityAsnIPResponse]
type radarEntityAsnIPResponseJSON struct {
	Asn         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityAsnIPResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityAsnIPResponseAsn struct {
	Asn            int64                                     `json:"asn,required"`
	Country        string                                    `json:"country,required"`
	CountryName    string                                    `json:"countryName,required"`
	EstimatedUsers RadarEntityAsnIPResponseAsnEstimatedUsers `json:"estimatedUsers,required"`
	Name           string                                    `json:"name,required"`
	OrgName        string                                    `json:"orgName,required"`
	Related        []RadarEntityAsnIPResponseAsnRelated      `json:"related,required"`
	// Regional Internet Registry
	Source  string `json:"source,required"`
	Website string `json:"website,required"`
	Aka     string `json:"aka"`
	// Deprecated field. Please use 'aka'.
	NameLong string                          `json:"nameLong"`
	JSON     radarEntityAsnIPResponseAsnJSON `json:"-"`
}

// radarEntityAsnIPResponseAsnJSON contains the JSON metadata for the struct
// [RadarEntityAsnIPResponseAsn]
type radarEntityAsnIPResponseAsnJSON struct {
	Asn            apijson.Field
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

func (r *RadarEntityAsnIPResponseAsn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityAsnIPResponseAsnEstimatedUsers struct {
	Locations []RadarEntityAsnIPResponseAsnEstimatedUsersLocation `json:"locations,required"`
	// Total estimated users
	EstimatedUsers int64                                         `json:"estimatedUsers"`
	JSON           radarEntityAsnIPResponseAsnEstimatedUsersJSON `json:"-"`
}

// radarEntityAsnIPResponseAsnEstimatedUsersJSON contains the JSON metadata for the
// struct [RadarEntityAsnIPResponseAsnEstimatedUsers]
type radarEntityAsnIPResponseAsnEstimatedUsersJSON struct {
	Locations      apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEntityAsnIPResponseAsnEstimatedUsers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityAsnIPResponseAsnEstimatedUsersLocation struct {
	LocationAlpha2 string `json:"locationAlpha2,required"`
	LocationName   string `json:"locationName,required"`
	// Estimated users per location
	EstimatedUsers int64                                                 `json:"estimatedUsers"`
	JSON           radarEntityAsnIPResponseAsnEstimatedUsersLocationJSON `json:"-"`
}

// radarEntityAsnIPResponseAsnEstimatedUsersLocationJSON contains the JSON metadata
// for the struct [RadarEntityAsnIPResponseAsnEstimatedUsersLocation]
type radarEntityAsnIPResponseAsnEstimatedUsersLocationJSON struct {
	LocationAlpha2 apijson.Field
	LocationName   apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEntityAsnIPResponseAsnEstimatedUsersLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityAsnIPResponseAsnRelated struct {
	Asn  int64  `json:"asn,required"`
	Name string `json:"name,required"`
	Aka  string `json:"aka"`
	// Total estimated users
	EstimatedUsers int64                                  `json:"estimatedUsers"`
	JSON           radarEntityAsnIPResponseAsnRelatedJSON `json:"-"`
}

// radarEntityAsnIPResponseAsnRelatedJSON contains the JSON metadata for the struct
// [RadarEntityAsnIPResponseAsnRelated]
type radarEntityAsnIPResponseAsnRelatedJSON struct {
	Asn            apijson.Field
	Name           apijson.Field
	Aka            apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEntityAsnIPResponseAsnRelated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type RadarEntityAsnListParams struct {
	// Comma separated list of ASNs.
	Asn param.Field[string] `query:"asn"`
	// Format results are returned in.
	Format param.Field[RadarEntityAsnListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Location Alpha2 to filter results.
	Location param.Field[string] `query:"location"`
	// Number of objects to skip before grabbing results.
	Offset param.Field[int64] `query:"offset"`
	// Order asn list.
	OrderBy param.Field[RadarEntityAsnListParamsOrderBy] `query:"orderBy"`
}

// URLQuery serializes [RadarEntityAsnListParams]'s query parameters as
// `url.Values`.
func (r RadarEntityAsnListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarEntityAsnListParamsFormat string

const (
	RadarEntityAsnListParamsFormatJson RadarEntityAsnListParamsFormat = "JSON"
	RadarEntityAsnListParamsFormatCsv  RadarEntityAsnListParamsFormat = "CSV"
)

// Order asn list.
type RadarEntityAsnListParamsOrderBy string

const (
	RadarEntityAsnListParamsOrderByAsn        RadarEntityAsnListParamsOrderBy = "ASN"
	RadarEntityAsnListParamsOrderByPopulation RadarEntityAsnListParamsOrderBy = "POPULATION"
)

type RadarEntityAsnListResponseEnvelope struct {
	Result  RadarEntityAsnListResponse             `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    radarEntityAsnListResponseEnvelopeJSON `json:"-"`
}

// radarEntityAsnListResponseEnvelopeJSON contains the JSON metadata for the struct
// [RadarEntityAsnListResponseEnvelope]
type radarEntityAsnListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityAsnListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityAsnGetParams struct {
	// Format results are returned in.
	Format param.Field[RadarEntityAsnGetParamsFormat] `query:"format"`
}

// URLQuery serializes [RadarEntityAsnGetParams]'s query parameters as
// `url.Values`.
func (r RadarEntityAsnGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarEntityAsnGetParamsFormat string

const (
	RadarEntityAsnGetParamsFormatJson RadarEntityAsnGetParamsFormat = "JSON"
	RadarEntityAsnGetParamsFormatCsv  RadarEntityAsnGetParamsFormat = "CSV"
)

type RadarEntityAsnGetResponseEnvelope struct {
	Result  RadarEntityAsnGetResponse             `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    radarEntityAsnGetResponseEnvelopeJSON `json:"-"`
}

// radarEntityAsnGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RadarEntityAsnGetResponseEnvelope]
type radarEntityAsnGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityAsnGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityAsnIPParams struct {
	// IP address.
	IP param.Field[string] `query:"ip,required"`
	// Format results are returned in.
	Format param.Field[RadarEntityAsnIPParamsFormat] `query:"format"`
}

// URLQuery serializes [RadarEntityAsnIPParams]'s query parameters as `url.Values`.
func (r RadarEntityAsnIPParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarEntityAsnIPParamsFormat string

const (
	RadarEntityAsnIPParamsFormatJson RadarEntityAsnIPParamsFormat = "JSON"
	RadarEntityAsnIPParamsFormatCsv  RadarEntityAsnIPParamsFormat = "CSV"
)

type RadarEntityAsnIPResponseEnvelope struct {
	Result  RadarEntityAsnIPResponse             `json:"result,required"`
	Success bool                                 `json:"success,required"`
	JSON    radarEntityAsnIPResponseEnvelopeJSON `json:"-"`
}

// radarEntityAsnIPResponseEnvelopeJSON contains the JSON metadata for the struct
// [RadarEntityAsnIPResponseEnvelope]
type radarEntityAsnIPResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityAsnIPResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
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
