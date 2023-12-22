// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarRankingTopService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarRankingTopService] method
// instead.
type RadarRankingTopService struct {
	Options []option.RequestOption
}

// NewRadarRankingTopService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarRankingTopService(opts ...option.RequestOption) (r *RadarRankingTopService) {
	r = &RadarRankingTopService{}
	r.Options = opts
	return
}

// Gets Top Domains Rank globally or by country.
func (r *RadarRankingTopService) List(ctx context.Context, query RadarRankingTopListParams, opts ...option.RequestOption) (res *RadarRankingTopListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/ranking/top"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarRankingTopListResponse struct {
	Result  RadarRankingTopListResponseResult `json:"result,required"`
	Success bool                              `json:"success,required"`
	JSON    radarRankingTopListResponseJSON   `json:"-"`
}

// radarRankingTopListResponseJSON contains the JSON metadata for the struct
// [RadarRankingTopListResponse]
type radarRankingTopListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingTopListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingTopListResponseResult struct {
	Meta RadarRankingTopListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarRankingTopListResponseResultTop0 `json:"top_0,required"`
	JSON radarRankingTopListResponseResultJSON   `json:"-"`
}

// radarRankingTopListResponseResultJSON contains the JSON metadata for the struct
// [RadarRankingTopListResponseResult]
type radarRankingTopListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingTopListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingTopListResponseResultMeta struct {
	Top0 RadarRankingTopListResponseResultMetaTop0 `json:"top_0,required"`
	JSON radarRankingTopListResponseResultMetaJSON `json:"-"`
}

// radarRankingTopListResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarRankingTopListResponseResultMeta]
type radarRankingTopListResponseResultMetaJSON struct {
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingTopListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingTopListResponseResultMetaTop0 struct {
	Date string                                        `json:"date,required"`
	JSON radarRankingTopListResponseResultMetaTop0JSON `json:"-"`
}

// radarRankingTopListResponseResultMetaTop0JSON contains the JSON metadata for the
// struct [RadarRankingTopListResponseResultMetaTop0]
type radarRankingTopListResponseResultMetaTop0JSON struct {
	Date        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingTopListResponseResultMetaTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingTopListResponseResultTop0 struct {
	Categories []string                                  `json:"categories,required"`
	Domain     string                                    `json:"domain,required"`
	Rank       int64                                     `json:"rank,required"`
	JSON       radarRankingTopListResponseResultTop0JSON `json:"-"`
}

// radarRankingTopListResponseResultTop0JSON contains the JSON metadata for the
// struct [RadarRankingTopListResponseResultTop0]
type radarRankingTopListResponseResultTop0JSON struct {
	Categories  apijson.Field
	Domain      apijson.Field
	Rank        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingTopListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarRankingTopListParams struct {
	// Array of dates to filter the ranking.
	Date param.Field[[]time.Time] `query:"date" format:"date"`
	// Format results are returned in.
	Format param.Field[RadarRankingTopListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of locations (alpha-2 country codes).
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarRankingTopListParams]'s query parameters as
// `url.Values`.
func (r RadarRankingTopListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarRankingTopListParamsFormat string

const (
	RadarRankingTopListParamsFormatJson RadarRankingTopListParamsFormat = "JSON"
	RadarRankingTopListParamsFormatCsv  RadarRankingTopListParamsFormat = "CSV"
)
