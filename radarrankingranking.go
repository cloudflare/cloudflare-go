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

// RadarRankingRankingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarRankingRankingService]
// method instead.
type RadarRankingRankingService struct {
	Options []option.RequestOption
}

// NewRadarRankingRankingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarRankingRankingService(opts ...option.RequestOption) (r *RadarRankingRankingService) {
	r = &RadarRankingRankingService{}
	r.Options = opts
	return
}

// Gets Domains Rank details. Cloudflare provides an ordered rank for the top 100
// domains, but for the remainder it only provides ranking buckets like top 200
// thousand, top one million, etc.. These are available through Radar datasets
// endpoints.
func (r *RadarRankingRankingService) Get(ctx context.Context, domain string, query RadarRankingRankingGetParams, opts ...option.RequestOption) (res *RadarRankingRankingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarRankingRankingGetResponseEnvelope
	path := fmt.Sprintf("radar/ranking/domain/%s", domain)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get top or trending domains based on their rank. Popular domains are domains of
// broad appeal based on how people use the Internet. Trending domains are domains
// that are generating a surge in interest. For more information on top domains,
// see https://blog.cloudflare.com/radar-domain-rankings/.
func (r *RadarRankingRankingService) Top(ctx context.Context, query RadarRankingRankingTopParams, opts ...option.RequestOption) (res *RadarRankingRankingTopResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarRankingRankingTopResponseEnvelope
	path := "radar/ranking/top"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarRankingRankingGetResponse struct {
	Details0 RadarRankingRankingGetResponseDetails0 `json:"details_0,required"`
	JSON     radarRankingRankingGetResponseJSON     `json:"-"`
}

// radarRankingRankingGetResponseJSON contains the JSON metadata for the struct
// [RadarRankingRankingGetResponse]
type radarRankingRankingGetResponseJSON struct {
	Details0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingRankingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingRankingGetResponseDetails0 struct {
	Categories   []RadarRankingRankingGetResponseDetails0Category    `json:"categories,required"`
	TopLocations []RadarRankingRankingGetResponseDetails0TopLocation `json:"top_locations,required"`
	// Only available in POPULAR ranking for the most recent ranking.
	Bucket string                                     `json:"bucket"`
	Rank   int64                                      `json:"rank"`
	JSON   radarRankingRankingGetResponseDetails0JSON `json:"-"`
}

// radarRankingRankingGetResponseDetails0JSON contains the JSON metadata for the
// struct [RadarRankingRankingGetResponseDetails0]
type radarRankingRankingGetResponseDetails0JSON struct {
	Categories   apijson.Field
	TopLocations apijson.Field
	Bucket       apijson.Field
	Rank         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarRankingRankingGetResponseDetails0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingRankingGetResponseDetails0Category struct {
	ID              float64                                            `json:"id,required"`
	Name            string                                             `json:"name,required"`
	SuperCategoryID float64                                            `json:"superCategoryId,required"`
	JSON            radarRankingRankingGetResponseDetails0CategoryJSON `json:"-"`
}

// radarRankingRankingGetResponseDetails0CategoryJSON contains the JSON metadata
// for the struct [RadarRankingRankingGetResponseDetails0Category]
type radarRankingRankingGetResponseDetails0CategoryJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarRankingRankingGetResponseDetails0Category) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingRankingGetResponseDetails0TopLocation struct {
	LocationCode string                                                `json:"locationCode,required"`
	LocationName string                                                `json:"locationName,required"`
	Rank         int64                                                 `json:"rank,required"`
	JSON         radarRankingRankingGetResponseDetails0TopLocationJSON `json:"-"`
}

// radarRankingRankingGetResponseDetails0TopLocationJSON contains the JSON metadata
// for the struct [RadarRankingRankingGetResponseDetails0TopLocation]
type radarRankingRankingGetResponseDetails0TopLocationJSON struct {
	LocationCode apijson.Field
	LocationName apijson.Field
	Rank         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarRankingRankingGetResponseDetails0TopLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingRankingTopResponse struct {
	Meta RadarRankingRankingTopResponseMeta   `json:"meta,required"`
	Top0 []RadarRankingRankingTopResponseTop0 `json:"top_0,required"`
	JSON radarRankingRankingTopResponseJSON   `json:"-"`
}

// radarRankingRankingTopResponseJSON contains the JSON metadata for the struct
// [RadarRankingRankingTopResponse]
type radarRankingRankingTopResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingRankingTopResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingRankingTopResponseMeta struct {
	Top0 RadarRankingRankingTopResponseMetaTop0 `json:"top_0,required"`
	JSON radarRankingRankingTopResponseMetaJSON `json:"-"`
}

// radarRankingRankingTopResponseMetaJSON contains the JSON metadata for the struct
// [RadarRankingRankingTopResponseMeta]
type radarRankingRankingTopResponseMetaJSON struct {
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingRankingTopResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingRankingTopResponseMetaTop0 struct {
	Date string                                     `json:"date,required"`
	JSON radarRankingRankingTopResponseMetaTop0JSON `json:"-"`
}

// radarRankingRankingTopResponseMetaTop0JSON contains the JSON metadata for the
// struct [RadarRankingRankingTopResponseMetaTop0]
type radarRankingRankingTopResponseMetaTop0JSON struct {
	Date        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingRankingTopResponseMetaTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingRankingTopResponseTop0 struct {
	Categories []RadarRankingRankingTopResponseTop0Category `json:"categories,required"`
	Domain     string                                       `json:"domain,required"`
	Rank       int64                                        `json:"rank,required"`
	// Only available in TRENDING rankings.
	PctRankChange float64                                `json:"pctRankChange"`
	JSON          radarRankingRankingTopResponseTop0JSON `json:"-"`
}

// radarRankingRankingTopResponseTop0JSON contains the JSON metadata for the struct
// [RadarRankingRankingTopResponseTop0]
type radarRankingRankingTopResponseTop0JSON struct {
	Categories    apijson.Field
	Domain        apijson.Field
	Rank          apijson.Field
	PctRankChange apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RadarRankingRankingTopResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingRankingTopResponseTop0Category struct {
	ID              float64                                        `json:"id,required"`
	Name            string                                         `json:"name,required"`
	SuperCategoryID float64                                        `json:"superCategoryId,required"`
	JSON            radarRankingRankingTopResponseTop0CategoryJSON `json:"-"`
}

// radarRankingRankingTopResponseTop0CategoryJSON contains the JSON metadata for
// the struct [RadarRankingRankingTopResponseTop0Category]
type radarRankingRankingTopResponseTop0CategoryJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarRankingRankingTopResponseTop0Category) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingRankingGetParams struct {
	// Array of dates to filter the ranking.
	Date param.Field[[]string] `query:"date"`
	// Format results are returned in.
	Format param.Field[RadarRankingRankingGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// The ranking type.
	RankingType param.Field[RadarRankingRankingGetParamsRankingType] `query:"rankingType"`
}

// URLQuery serializes [RadarRankingRankingGetParams]'s query parameters as
// `url.Values`.
func (r RadarRankingRankingGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarRankingRankingGetParamsFormat string

const (
	RadarRankingRankingGetParamsFormatJson RadarRankingRankingGetParamsFormat = "JSON"
	RadarRankingRankingGetParamsFormatCsv  RadarRankingRankingGetParamsFormat = "CSV"
)

// The ranking type.
type RadarRankingRankingGetParamsRankingType string

const (
	RadarRankingRankingGetParamsRankingTypePopular        RadarRankingRankingGetParamsRankingType = "POPULAR"
	RadarRankingRankingGetParamsRankingTypeTrendingRise   RadarRankingRankingGetParamsRankingType = "TRENDING_RISE"
	RadarRankingRankingGetParamsRankingTypeTrendingSteady RadarRankingRankingGetParamsRankingType = "TRENDING_STEADY"
)

type RadarRankingRankingGetResponseEnvelope struct {
	Result  RadarRankingRankingGetResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarRankingRankingGetResponseEnvelopeJSON `json:"-"`
}

// radarRankingRankingGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarRankingRankingGetResponseEnvelope]
type radarRankingRankingGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingRankingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingRankingTopParams struct {
	// Array of dates to filter the ranking.
	Date param.Field[[]string] `query:"date"`
	// Format results are returned in.
	Format param.Field[RadarRankingRankingTopParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of locations (alpha-2 country codes).
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// The ranking type.
	RankingType param.Field[RadarRankingRankingTopParamsRankingType] `query:"rankingType"`
}

// URLQuery serializes [RadarRankingRankingTopParams]'s query parameters as
// `url.Values`.
func (r RadarRankingRankingTopParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarRankingRankingTopParamsFormat string

const (
	RadarRankingRankingTopParamsFormatJson RadarRankingRankingTopParamsFormat = "JSON"
	RadarRankingRankingTopParamsFormatCsv  RadarRankingRankingTopParamsFormat = "CSV"
)

// The ranking type.
type RadarRankingRankingTopParamsRankingType string

const (
	RadarRankingRankingTopParamsRankingTypePopular        RadarRankingRankingTopParamsRankingType = "POPULAR"
	RadarRankingRankingTopParamsRankingTypeTrendingRise   RadarRankingRankingTopParamsRankingType = "TRENDING_RISE"
	RadarRankingRankingTopParamsRankingTypeTrendingSteady RadarRankingRankingTopParamsRankingType = "TRENDING_STEADY"
)

type RadarRankingRankingTopResponseEnvelope struct {
	Result  RadarRankingRankingTopResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarRankingRankingTopResponseEnvelopeJSON `json:"-"`
}

// radarRankingRankingTopResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarRankingRankingTopResponseEnvelope]
type radarRankingRankingTopResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingRankingTopResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
