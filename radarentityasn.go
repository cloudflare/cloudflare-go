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

// RadarEntityASNService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEntityASNService] method
// instead.
type RadarEntityASNService struct {
	Options []option.RequestOption
	IPs     *RadarEntityASNIPService
}

// NewRadarEntityASNService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarEntityASNService(opts ...option.RequestOption) (r *RadarEntityASNService) {
	r = &RadarEntityASNService{}
	r.Options = opts
	r.IPs = NewRadarEntityASNIPService(opts...)
	return
}

// Get the requested autonomous system information. A confidence level below `5`
// indicates a low level of confidence in the traffic data - normally this happens
// because Cloudflare has a small amount of traffic from/to this AS). Population
// estimates come from APNIC (refer to https://labs.apnic.net/?p=526).
func (r *RadarEntityASNService) Get(ctx context.Context, asn int64, query RadarEntityASNGetParams, opts ...option.RequestOption) (res *RadarEntityASNGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/entities/asns/%v", asn)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Gets a list of autonomous systems (AS).
func (r *RadarEntityASNService) List(ctx context.Context, query RadarEntityASNListParams, opts ...option.RequestOption) (res *RadarEntityASNListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/entities/asns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEntityASNGetResponse struct {
	Result  RadarEntityASNGetResponseResult `json:"result,required"`
	Success bool                            `json:"success,required"`
	JSON    radarEntityASNGetResponseJSON   `json:"-"`
}

// radarEntityASNGetResponseJSON contains the JSON metadata for the struct
// [RadarEntityASNGetResponse]
type radarEntityASNGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityASNGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityASNGetResponseResult struct {
	ASN  RadarEntityASNGetResponseResultASN  `json:"asn,required"`
	JSON radarEntityASNGetResponseResultJSON `json:"-"`
}

// radarEntityASNGetResponseResultJSON contains the JSON metadata for the struct
// [RadarEntityASNGetResponseResult]
type radarEntityASNGetResponseResultJSON struct {
	ASN         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityASNGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityASNGetResponseResultASN struct {
	Aka             string                                           `json:"aka,required"`
	ASN             int64                                            `json:"asn,required"`
	ConfidenceLevel int64                                            `json:"confidenceLevel,required"`
	Country         string                                           `json:"country,required"`
	CountryName     string                                           `json:"countryName,required"`
	EstimatedUsers  RadarEntityASNGetResponseResultASNEstimatedUsers `json:"estimatedUsers,required"`
	Name            string                                           `json:"name,required"`
	NameLong        string                                           `json:"nameLong,required"`
	OrgName         string                                           `json:"orgName,required"`
	Related         []RadarEntityASNGetResponseResultASNRelated      `json:"related,required"`
	Website         string                                           `json:"website,required"`
	JSON            radarEntityASNGetResponseResultASNJSON           `json:"-"`
}

// radarEntityASNGetResponseResultASNJSON contains the JSON metadata for the struct
// [RadarEntityASNGetResponseResultASN]
type radarEntityASNGetResponseResultASNJSON struct {
	Aka             apijson.Field
	ASN             apijson.Field
	ConfidenceLevel apijson.Field
	Country         apijson.Field
	CountryName     apijson.Field
	EstimatedUsers  apijson.Field
	Name            apijson.Field
	NameLong        apijson.Field
	OrgName         apijson.Field
	Related         apijson.Field
	Website         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarEntityASNGetResponseResultASN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityASNGetResponseResultASNEstimatedUsers struct {
	Locations []RadarEntityASNGetResponseResultASNEstimatedUsersLocation `json:"locations,required"`
	// Total estimated users
	EstimatedUsers int64                                                `json:"estimatedUsers"`
	JSON           radarEntityASNGetResponseResultASNEstimatedUsersJSON `json:"-"`
}

// radarEntityASNGetResponseResultASNEstimatedUsersJSON contains the JSON metadata
// for the struct [RadarEntityASNGetResponseResultASNEstimatedUsers]
type radarEntityASNGetResponseResultASNEstimatedUsersJSON struct {
	Locations      apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEntityASNGetResponseResultASNEstimatedUsers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityASNGetResponseResultASNEstimatedUsersLocation struct {
	// Estimated users per location
	EstimatedUsers int64                                                        `json:"estimatedUsers,required"`
	LocationAlpha2 string                                                       `json:"locationAlpha2,required"`
	LocationName   string                                                       `json:"locationName,required"`
	JSON           radarEntityASNGetResponseResultASNEstimatedUsersLocationJSON `json:"-"`
}

// radarEntityASNGetResponseResultASNEstimatedUsersLocationJSON contains the JSON
// metadata for the struct
// [RadarEntityASNGetResponseResultASNEstimatedUsersLocation]
type radarEntityASNGetResponseResultASNEstimatedUsersLocationJSON struct {
	EstimatedUsers apijson.Field
	LocationAlpha2 apijson.Field
	LocationName   apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEntityASNGetResponseResultASNEstimatedUsersLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityASNGetResponseResultASNRelated struct {
	Aka  string `json:"aka,required"`
	ASN  int64  `json:"asn,required"`
	Name string `json:"name,required"`
	// Total estimated users
	EstimatedUsers int64                                         `json:"estimatedUsers"`
	JSON           radarEntityASNGetResponseResultASNRelatedJSON `json:"-"`
}

// radarEntityASNGetResponseResultASNRelatedJSON contains the JSON metadata for the
// struct [RadarEntityASNGetResponseResultASNRelated]
type radarEntityASNGetResponseResultASNRelatedJSON struct {
	Aka            apijson.Field
	ASN            apijson.Field
	Name           apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEntityASNGetResponseResultASNRelated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityASNListResponse struct {
	Result  RadarEntityASNListResponseResult `json:"result,required"`
	Success bool                             `json:"success,required"`
	JSON    radarEntityASNListResponseJSON   `json:"-"`
}

// radarEntityASNListResponseJSON contains the JSON metadata for the struct
// [RadarEntityASNListResponse]
type radarEntityASNListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityASNListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityASNListResponseResult struct {
	ASNs []RadarEntityASNListResponseResultASN `json:"asns,required"`
	JSON radarEntityASNListResponseResultJSON  `json:"-"`
}

// radarEntityASNListResponseResultJSON contains the JSON metadata for the struct
// [RadarEntityASNListResponseResult]
type radarEntityASNListResponseResultJSON struct {
	ASNs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityASNListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityASNListResponseResultASN struct {
	Aka         string                                  `json:"aka,required"`
	ASN         int64                                   `json:"asn,required"`
	Country     string                                  `json:"country,required"`
	CountryName string                                  `json:"countryName,required"`
	Name        string                                  `json:"name,required"`
	NameLong    string                                  `json:"nameLong,required"`
	OrgName     string                                  `json:"orgName,required"`
	Website     string                                  `json:"website,required"`
	JSON        radarEntityASNListResponseResultASNJSON `json:"-"`
}

// radarEntityASNListResponseResultASNJSON contains the JSON metadata for the
// struct [RadarEntityASNListResponseResultASN]
type radarEntityASNListResponseResultASNJSON struct {
	Aka         apijson.Field
	ASN         apijson.Field
	Country     apijson.Field
	CountryName apijson.Field
	Name        apijson.Field
	NameLong    apijson.Field
	OrgName     apijson.Field
	Website     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityASNListResponseResultASN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type RadarEntityASNListParams struct {
	// Comma separated list of ASNs.
	ASN param.Field[string] `query:"asn"`
	// Format results are returned in.
	Format param.Field[RadarEntityASNListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Number of objects to skip before grabbing results.
	Offset param.Field[int64] `query:"offset"`
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
