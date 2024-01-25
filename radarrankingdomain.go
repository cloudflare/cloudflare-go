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

// RadarRankingDomainService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarRankingDomainService] method
// instead.
type RadarRankingDomainService struct {
	Options []option.RequestOption
}

// NewRadarRankingDomainService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarRankingDomainService(opts ...option.RequestOption) (r *RadarRankingDomainService) {
	r = &RadarRankingDomainService{}
	r.Options = opts
	return
}

// Gets Domains Rank details. Cloudflare provides an ordered rank for the top 100
// domains, but for the remainder it only provides ranking buckets like top 200
// thousand, top one million, etc.. These are available through Radar datasets
// endpoints.
func (r *RadarRankingDomainService) Get(ctx context.Context, domain string, query RadarRankingDomainGetParams, opts ...option.RequestOption) (res *RadarRankingDomainGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/ranking/domain/%s", domain)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarRankingDomainGetResponse struct {
	Result  RadarRankingDomainGetResponseResult `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    radarRankingDomainGetResponseJSON   `json:"-"`
}

// radarRankingDomainGetResponseJSON contains the JSON metadata for the struct
// [RadarRankingDomainGetResponse]
type radarRankingDomainGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingDomainGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingDomainGetResponseResult struct {
	Details0 RadarRankingDomainGetResponseResultDetails0 `json:"details_0,required"`
	JSON     radarRankingDomainGetResponseResultJSON     `json:"-"`
}

// radarRankingDomainGetResponseResultJSON contains the JSON metadata for the
// struct [RadarRankingDomainGetResponseResult]
type radarRankingDomainGetResponseResultJSON struct {
	Details0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingDomainGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingDomainGetResponseResultDetails0 struct {
	Categories   []RadarRankingDomainGetResponseResultDetails0Category    `json:"categories,required"`
	TopLocations []RadarRankingDomainGetResponseResultDetails0TopLocation `json:"top_locations,required"`
	// Only available in POPULAR ranking for the most recent ranking.
	Bucket string                                          `json:"bucket"`
	Rank   int64                                           `json:"rank"`
	JSON   radarRankingDomainGetResponseResultDetails0JSON `json:"-"`
}

// radarRankingDomainGetResponseResultDetails0JSON contains the JSON metadata for
// the struct [RadarRankingDomainGetResponseResultDetails0]
type radarRankingDomainGetResponseResultDetails0JSON struct {
	Categories   apijson.Field
	TopLocations apijson.Field
	Bucket       apijson.Field
	Rank         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarRankingDomainGetResponseResultDetails0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingDomainGetResponseResultDetails0Category struct {
	ID              float64                                                 `json:"id,required"`
	Name            string                                                  `json:"name,required"`
	SuperCategoryID float64                                                 `json:"superCategoryId,required"`
	JSON            radarRankingDomainGetResponseResultDetails0CategoryJSON `json:"-"`
}

// radarRankingDomainGetResponseResultDetails0CategoryJSON contains the JSON
// metadata for the struct [RadarRankingDomainGetResponseResultDetails0Category]
type radarRankingDomainGetResponseResultDetails0CategoryJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarRankingDomainGetResponseResultDetails0Category) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingDomainGetResponseResultDetails0TopLocation struct {
	LocationCode string                                                     `json:"locationCode,required"`
	LocationName string                                                     `json:"locationName,required"`
	Rank         int64                                                      `json:"rank,required"`
	JSON         radarRankingDomainGetResponseResultDetails0TopLocationJSON `json:"-"`
}

// radarRankingDomainGetResponseResultDetails0TopLocationJSON contains the JSON
// metadata for the struct [RadarRankingDomainGetResponseResultDetails0TopLocation]
type radarRankingDomainGetResponseResultDetails0TopLocationJSON struct {
	LocationCode apijson.Field
	LocationName apijson.Field
	Rank         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarRankingDomainGetResponseResultDetails0TopLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingDomainGetParams struct {
	// Array of dates to filter the ranking.
	Date param.Field[[]string] `query:"date"`
	// Format results are returned in.
	Format param.Field[RadarRankingDomainGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// The ranking type.
	RankingType param.Field[RadarRankingDomainGetParamsRankingType] `query:"rankingType"`
}

// URLQuery serializes [RadarRankingDomainGetParams]'s query parameters as
// `url.Values`.
func (r RadarRankingDomainGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarRankingDomainGetParamsFormat string

const (
	RadarRankingDomainGetParamsFormatJson RadarRankingDomainGetParamsFormat = "JSON"
	RadarRankingDomainGetParamsFormatCsv  RadarRankingDomainGetParamsFormat = "CSV"
)

// The ranking type.
type RadarRankingDomainGetParamsRankingType string

const (
	RadarRankingDomainGetParamsRankingTypePopular        RadarRankingDomainGetParamsRankingType = "POPULAR"
	RadarRankingDomainGetParamsRankingTypeTrendingRise   RadarRankingDomainGetParamsRankingType = "TRENDING_RISE"
	RadarRankingDomainGetParamsRankingTypeTrendingSteady RadarRankingDomainGetParamsRankingType = "TRENDING_STEADY"
)
