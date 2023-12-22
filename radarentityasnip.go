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

// RadarEntityASNIPService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEntityASNIPService] method
// instead.
type RadarEntityASNIPService struct {
	Options []option.RequestOption
}

// NewRadarEntityASNIPService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarEntityASNIPService(opts ...option.RequestOption) (r *RadarEntityASNIPService) {
	r = &RadarEntityASNIPService{}
	r.Options = opts
	return
}

// Get the requested autonomous system information based on IP address. Population
// estimates come from APNIC (refer to https://labs.apnic.net/?p=526).
func (r *RadarEntityASNIPService) List(ctx context.Context, query RadarEntityASNIPListParams, opts ...option.RequestOption) (res *RadarEntityAsnipListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/entities/asns/ip"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEntityAsnipListResponse struct {
	Result  RadarEntityAsnipListResponseResult `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    radarEntityAsnipListResponseJSON   `json:"-"`
}

// radarEntityAsnipListResponseJSON contains the JSON metadata for the struct
// [RadarEntityAsnipListResponse]
type radarEntityAsnipListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityAsnipListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityAsnipListResponseResult struct {
	ASN  RadarEntityAsnipListResponseResultASN  `json:"asn,required"`
	JSON radarEntityAsnipListResponseResultJSON `json:"-"`
}

// radarEntityAsnipListResponseResultJSON contains the JSON metadata for the struct
// [RadarEntityAsnipListResponseResult]
type radarEntityAsnipListResponseResultJSON struct {
	ASN         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityAsnipListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityAsnipListResponseResultASN struct {
	Aka            string                                              `json:"aka,required"`
	ASN            int64                                               `json:"asn,required"`
	Country        string                                              `json:"country,required"`
	CountryName    string                                              `json:"countryName,required"`
	EstimatedUsers RadarEntityAsnipListResponseResultASNEstimatedUsers `json:"estimatedUsers,required"`
	Name           string                                              `json:"name,required"`
	NameLong       string                                              `json:"nameLong,required"`
	OrgName        string                                              `json:"orgName,required"`
	Related        []RadarEntityAsnipListResponseResultASNRelated      `json:"related,required"`
	Website        string                                              `json:"website,required"`
	JSON           radarEntityAsnipListResponseResultASNJSON           `json:"-"`
}

// radarEntityAsnipListResponseResultASNJSON contains the JSON metadata for the
// struct [RadarEntityAsnipListResponseResultASN]
type radarEntityAsnipListResponseResultASNJSON struct {
	Aka            apijson.Field
	ASN            apijson.Field
	Country        apijson.Field
	CountryName    apijson.Field
	EstimatedUsers apijson.Field
	Name           apijson.Field
	NameLong       apijson.Field
	OrgName        apijson.Field
	Related        apijson.Field
	Website        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEntityAsnipListResponseResultASN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityAsnipListResponseResultASNEstimatedUsers struct {
	Locations []RadarEntityAsnipListResponseResultASNEstimatedUsersLocation `json:"locations,required"`
	// Total estimated users
	EstimatedUsers int64                                                   `json:"estimatedUsers"`
	JSON           radarEntityAsnipListResponseResultASNEstimatedUsersJSON `json:"-"`
}

// radarEntityAsnipListResponseResultASNEstimatedUsersJSON contains the JSON
// metadata for the struct [RadarEntityAsnipListResponseResultASNEstimatedUsers]
type radarEntityAsnipListResponseResultASNEstimatedUsersJSON struct {
	Locations      apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEntityAsnipListResponseResultASNEstimatedUsers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityAsnipListResponseResultASNEstimatedUsersLocation struct {
	// Estimated users per location
	EstimatedUsers int64                                                           `json:"estimatedUsers,required"`
	LocationAlpha2 string                                                          `json:"locationAlpha2,required"`
	LocationName   string                                                          `json:"locationName,required"`
	JSON           radarEntityAsnipListResponseResultASNEstimatedUsersLocationJSON `json:"-"`
}

// radarEntityAsnipListResponseResultASNEstimatedUsersLocationJSON contains the
// JSON metadata for the struct
// [RadarEntityAsnipListResponseResultASNEstimatedUsersLocation]
type radarEntityAsnipListResponseResultASNEstimatedUsersLocationJSON struct {
	EstimatedUsers apijson.Field
	LocationAlpha2 apijson.Field
	LocationName   apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEntityAsnipListResponseResultASNEstimatedUsersLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityAsnipListResponseResultASNRelated struct {
	Aka  string `json:"aka,required"`
	ASN  int64  `json:"asn,required"`
	Name string `json:"name,required"`
	// Total estimated users
	EstimatedUsers int64                                            `json:"estimatedUsers"`
	JSON           radarEntityAsnipListResponseResultASNRelatedJSON `json:"-"`
}

// radarEntityAsnipListResponseResultASNRelatedJSON contains the JSON metadata for
// the struct [RadarEntityAsnipListResponseResultASNRelated]
type radarEntityAsnipListResponseResultASNRelatedJSON struct {
	Aka            apijson.Field
	ASN            apijson.Field
	Name           apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEntityAsnipListResponseResultASNRelated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityASNIPListParams struct {
	// IP address.
	IP param.Field[string] `query:"ip,required"`
	// Format results are returned in.
	Format param.Field[RadarEntityAsnipListParamsFormat] `query:"format"`
}

// URLQuery serializes [RadarEntityASNIPListParams]'s query parameters as
// `url.Values`.
func (r RadarEntityASNIPListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarEntityAsnipListParamsFormat string

const (
	RadarEntityAsnipListParamsFormatJson RadarEntityAsnipListParamsFormat = "JSON"
	RadarEntityAsnipListParamsFormatCsv  RadarEntityAsnipListParamsFormat = "CSV"
)
