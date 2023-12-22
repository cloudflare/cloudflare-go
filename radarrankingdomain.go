// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

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

// Gets Domains Rank details.
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
	Bucket       string                                                   `json:"bucket,required"`
	Categories   []string                                                 `json:"categories,required"`
	TopLocations []RadarRankingDomainGetResponseResultDetails0TopLocation `json:"top_locations,required"`
	Rank         int64                                                    `json:"rank"`
	JSON         radarRankingDomainGetResponseResultDetails0JSON          `json:"-"`
}

// radarRankingDomainGetResponseResultDetails0JSON contains the JSON metadata for
// the struct [RadarRankingDomainGetResponseResultDetails0]
type radarRankingDomainGetResponseResultDetails0JSON struct {
	Bucket       apijson.Field
	Categories   apijson.Field
	TopLocations apijson.Field
	Rank         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarRankingDomainGetResponseResultDetails0) UnmarshalJSON(data []byte) (err error) {
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
	Date param.Field[[]time.Time] `query:"date" format:"date"`
	// Format results are returned in.
	Format param.Field[RadarRankingDomainGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
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
