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

// RadarSearchService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarSearchService] method
// instead.
type RadarSearchService struct {
	Options []option.RequestOption
}

// NewRadarSearchService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarSearchService(opts ...option.RequestOption) (r *RadarSearchService) {
	r = &RadarSearchService{}
	r.Options = opts
	return
}

// Lets you search for locations, autonomous systems (AS) and reports.
func (r *RadarSearchService) Global(ctx context.Context, query RadarSearchGlobalParams, opts ...option.RequestOption) (res *RadarSearchGlobalResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarSearchGlobalResponseEnvelope
	path := "radar/search/global"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarSearchGlobalResponse struct {
	Search []RadarSearchGlobalResponseSearch `json:"search,required"`
	JSON   radarSearchGlobalResponseJSON     `json:"-"`
}

// radarSearchGlobalResponseJSON contains the JSON metadata for the struct
// [RadarSearchGlobalResponse]
type radarSearchGlobalResponseJSON struct {
	Search      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarSearchGlobalResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarSearchGlobalResponseSearch struct {
	Code string                              `json:"code,required"`
	Name string                              `json:"name,required"`
	Type string                              `json:"type,required"`
	JSON radarSearchGlobalResponseSearchJSON `json:"-"`
}

// radarSearchGlobalResponseSearchJSON contains the JSON metadata for the struct
// [RadarSearchGlobalResponseSearch]
type radarSearchGlobalResponseSearchJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarSearchGlobalResponseSearch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarSearchGlobalParams struct {
	// Search for locations, AS and reports.
	Query param.Field[string] `query:"query,required"`
	// Search types to be excluded from results.
	Exclude param.Field[[]RadarSearchGlobalParamsExclude] `query:"exclude"`
	// Format results are returned in.
	Format param.Field[RadarSearchGlobalParamsFormat] `query:"format"`
	// Search types to be included in results.
	Include param.Field[[]RadarSearchGlobalParamsInclude] `query:"include"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Limit the number of objects per search category.
	LimitPerGroup param.Field[float64] `query:"limitPerGroup"`
}

// URLQuery serializes [RadarSearchGlobalParams]'s query parameters as
// `url.Values`.
func (r RadarSearchGlobalParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarSearchGlobalParamsExclude string

const (
	RadarSearchGlobalParamsExcludeSpecialEvents RadarSearchGlobalParamsExclude = "SPECIAL_EVENTS"
	RadarSearchGlobalParamsExcludeNotebooks     RadarSearchGlobalParamsExclude = "NOTEBOOKS"
	RadarSearchGlobalParamsExcludeLocations     RadarSearchGlobalParamsExclude = "LOCATIONS"
	RadarSearchGlobalParamsExcludeAsns          RadarSearchGlobalParamsExclude = "ASNS"
)

// Format results are returned in.
type RadarSearchGlobalParamsFormat string

const (
	RadarSearchGlobalParamsFormatJson RadarSearchGlobalParamsFormat = "JSON"
	RadarSearchGlobalParamsFormatCsv  RadarSearchGlobalParamsFormat = "CSV"
)

type RadarSearchGlobalParamsInclude string

const (
	RadarSearchGlobalParamsIncludeSpecialEvents RadarSearchGlobalParamsInclude = "SPECIAL_EVENTS"
	RadarSearchGlobalParamsIncludeNotebooks     RadarSearchGlobalParamsInclude = "NOTEBOOKS"
	RadarSearchGlobalParamsIncludeLocations     RadarSearchGlobalParamsInclude = "LOCATIONS"
	RadarSearchGlobalParamsIncludeAsns          RadarSearchGlobalParamsInclude = "ASNS"
)

type RadarSearchGlobalResponseEnvelope struct {
	Result  RadarSearchGlobalResponse             `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    radarSearchGlobalResponseEnvelopeJSON `json:"-"`
}

// radarSearchGlobalResponseEnvelopeJSON contains the JSON metadata for the struct
// [RadarSearchGlobalResponseEnvelope]
type radarSearchGlobalResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarSearchGlobalResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
