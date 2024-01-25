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

// RadarSearchGlobalService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarSearchGlobalService] method
// instead.
type RadarSearchGlobalService struct {
	Options []option.RequestOption
}

// NewRadarSearchGlobalService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarSearchGlobalService(opts ...option.RequestOption) (r *RadarSearchGlobalService) {
	r = &RadarSearchGlobalService{}
	r.Options = opts
	return
}

// Lets you search for locations, autonomous systems (AS) and reports.
func (r *RadarSearchGlobalService) List(ctx context.Context, query RadarSearchGlobalListParams, opts ...option.RequestOption) (res *RadarSearchGlobalListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/search/global"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarSearchGlobalListResponse struct {
	Result  RadarSearchGlobalListResponseResult `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    radarSearchGlobalListResponseJSON   `json:"-"`
}

// radarSearchGlobalListResponseJSON contains the JSON metadata for the struct
// [RadarSearchGlobalListResponse]
type radarSearchGlobalListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarSearchGlobalListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarSearchGlobalListResponseResult struct {
	Search []RadarSearchGlobalListResponseResultSearch `json:"search,required"`
	JSON   radarSearchGlobalListResponseResultJSON     `json:"-"`
}

// radarSearchGlobalListResponseResultJSON contains the JSON metadata for the
// struct [RadarSearchGlobalListResponseResult]
type radarSearchGlobalListResponseResultJSON struct {
	Search      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarSearchGlobalListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarSearchGlobalListResponseResultSearch struct {
	Code string                                        `json:"code,required"`
	Name string                                        `json:"name,required"`
	Type string                                        `json:"type,required"`
	JSON radarSearchGlobalListResponseResultSearchJSON `json:"-"`
}

// radarSearchGlobalListResponseResultSearchJSON contains the JSON metadata for the
// struct [RadarSearchGlobalListResponseResultSearch]
type radarSearchGlobalListResponseResultSearchJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarSearchGlobalListResponseResultSearch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarSearchGlobalListParams struct {
	// Search for locations, AS and reports.
	Query param.Field[string] `query:"query,required"`
	// Search types to be excluded from results.
	Exclude param.Field[[]RadarSearchGlobalListParamsExclude] `query:"exclude"`
	// Format results are returned in.
	Format param.Field[RadarSearchGlobalListParamsFormat] `query:"format"`
	// Search types to be included in results.
	Include param.Field[[]RadarSearchGlobalListParamsInclude] `query:"include"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Limit the number of objects per search category.
	LimitPerGroup param.Field[float64] `query:"limitPerGroup"`
}

// URLQuery serializes [RadarSearchGlobalListParams]'s query parameters as
// `url.Values`.
func (r RadarSearchGlobalListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarSearchGlobalListParamsExclude string

const (
	RadarSearchGlobalListParamsExcludeSpecialEvents RadarSearchGlobalListParamsExclude = "SPECIAL_EVENTS"
	RadarSearchGlobalListParamsExcludeNotebooks     RadarSearchGlobalListParamsExclude = "NOTEBOOKS"
	RadarSearchGlobalListParamsExcludeLocations     RadarSearchGlobalListParamsExclude = "LOCATIONS"
	RadarSearchGlobalListParamsExcludeASNs          RadarSearchGlobalListParamsExclude = "ASNS"
)

// Format results are returned in.
type RadarSearchGlobalListParamsFormat string

const (
	RadarSearchGlobalListParamsFormatJson RadarSearchGlobalListParamsFormat = "JSON"
	RadarSearchGlobalListParamsFormatCsv  RadarSearchGlobalListParamsFormat = "CSV"
)

type RadarSearchGlobalListParamsInclude string

const (
	RadarSearchGlobalListParamsIncludeSpecialEvents RadarSearchGlobalListParamsInclude = "SPECIAL_EVENTS"
	RadarSearchGlobalListParamsIncludeNotebooks     RadarSearchGlobalListParamsInclude = "NOTEBOOKS"
	RadarSearchGlobalListParamsIncludeLocations     RadarSearchGlobalListParamsInclude = "LOCATIONS"
	RadarSearchGlobalListParamsIncludeASNs          RadarSearchGlobalListParamsInclude = "ASNS"
)
