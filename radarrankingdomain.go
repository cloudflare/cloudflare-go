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
	var env RadarRankingDomainGetResponseEnvelope
	path := fmt.Sprintf("radar/ranking/domain/%s", domain)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarRankingDomainGetResponse struct {
	Details0 RadarRankingDomainGetResponseDetails0 `json:"details_0,required"`
	JSON     radarRankingDomainGetResponseJSON     `json:"-"`
}

// radarRankingDomainGetResponseJSON contains the JSON metadata for the struct
// [RadarRankingDomainGetResponse]
type radarRankingDomainGetResponseJSON struct {
	Details0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingDomainGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingDomainGetResponseJSON) RawJSON() string {
	return r.raw
}

type RadarRankingDomainGetResponseDetails0 struct {
	Categories   []RadarRankingDomainGetResponseDetails0Category    `json:"categories,required"`
	TopLocations []RadarRankingDomainGetResponseDetails0TopLocation `json:"top_locations,required"`
	// Only available in POPULAR ranking for the most recent ranking.
	Bucket string                                    `json:"bucket"`
	Rank   int64                                     `json:"rank"`
	JSON   radarRankingDomainGetResponseDetails0JSON `json:"-"`
}

// radarRankingDomainGetResponseDetails0JSON contains the JSON metadata for the
// struct [RadarRankingDomainGetResponseDetails0]
type radarRankingDomainGetResponseDetails0JSON struct {
	Categories   apijson.Field
	TopLocations apijson.Field
	Bucket       apijson.Field
	Rank         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarRankingDomainGetResponseDetails0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingDomainGetResponseDetails0JSON) RawJSON() string {
	return r.raw
}

type RadarRankingDomainGetResponseDetails0Category struct {
	ID              float64                                           `json:"id,required"`
	Name            string                                            `json:"name,required"`
	SuperCategoryID float64                                           `json:"superCategoryId,required"`
	JSON            radarRankingDomainGetResponseDetails0CategoryJSON `json:"-"`
}

// radarRankingDomainGetResponseDetails0CategoryJSON contains the JSON metadata for
// the struct [RadarRankingDomainGetResponseDetails0Category]
type radarRankingDomainGetResponseDetails0CategoryJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarRankingDomainGetResponseDetails0Category) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingDomainGetResponseDetails0CategoryJSON) RawJSON() string {
	return r.raw
}

type RadarRankingDomainGetResponseDetails0TopLocation struct {
	LocationCode string                                               `json:"locationCode,required"`
	LocationName string                                               `json:"locationName,required"`
	Rank         int64                                                `json:"rank,required"`
	JSON         radarRankingDomainGetResponseDetails0TopLocationJSON `json:"-"`
}

// radarRankingDomainGetResponseDetails0TopLocationJSON contains the JSON metadata
// for the struct [RadarRankingDomainGetResponseDetails0TopLocation]
type radarRankingDomainGetResponseDetails0TopLocationJSON struct {
	LocationCode apijson.Field
	LocationName apijson.Field
	Rank         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarRankingDomainGetResponseDetails0TopLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingDomainGetResponseDetails0TopLocationJSON) RawJSON() string {
	return r.raw
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

type RadarRankingDomainGetResponseEnvelope struct {
	Result  RadarRankingDomainGetResponse             `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    radarRankingDomainGetResponseEnvelopeJSON `json:"-"`
}

// radarRankingDomainGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarRankingDomainGetResponseEnvelope]
type radarRankingDomainGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingDomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingDomainGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
