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

// RadarSpecialeventService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarSpecialeventService] method
// instead.
type RadarSpecialeventService struct {
	Options []option.RequestOption
}

// NewRadarSpecialeventService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarSpecialeventService(opts ...option.RequestOption) (r *RadarSpecialeventService) {
	r = &RadarSpecialeventService{}
	r.Options = opts
	return
}

// Get a single Special Event.
func (r *RadarSpecialeventService) Get(ctx context.Context, eventAlias string, query RadarSpecialeventGetParams, opts ...option.RequestOption) (res *RadarSpecialeventGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/specialevents/%s", eventAlias)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get Special Events time series
func (r *RadarSpecialeventService) List(ctx context.Context, eventAlias string, categoryAlias string, query RadarSpecialeventListParams, opts ...option.RequestOption) (res *RadarSpecialeventListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/specialevents/%s/%s", eventAlias, categoryAlias)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarSpecialeventGetResponse struct {
	Result  RadarSpecialeventGetResponseResult `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    radarSpecialeventGetResponseJSON   `json:"-"`
}

// radarSpecialeventGetResponseJSON contains the JSON metadata for the struct
// [RadarSpecialeventGetResponse]
type radarSpecialeventGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarSpecialeventGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarSpecialeventGetResponseResult struct {
	SpecialEvent RadarSpecialeventGetResponseResultSpecialEvent `json:"specialEvent,required"`
	JSON         radarSpecialeventGetResponseResultJSON         `json:"-"`
}

// radarSpecialeventGetResponseResultJSON contains the JSON metadata for the struct
// [RadarSpecialeventGetResponseResult]
type radarSpecialeventGetResponseResultJSON struct {
	SpecialEvent apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarSpecialeventGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarSpecialeventGetResponseResultSpecialEvent struct {
	ID          int64                                                    `json:"id,required"`
	Alias       string                                                   `json:"alias,required"`
	Categories  []RadarSpecialeventGetResponseResultSpecialEventCategory `json:"categories,required"`
	Title       string                                                   `json:"title,required"`
	Description string                                                   `json:"description"`
	JSON        radarSpecialeventGetResponseResultSpecialEventJSON       `json:"-"`
}

// radarSpecialeventGetResponseResultSpecialEventJSON contains the JSON metadata
// for the struct [RadarSpecialeventGetResponseResultSpecialEvent]
type radarSpecialeventGetResponseResultSpecialEventJSON struct {
	ID          apijson.Field
	Alias       apijson.Field
	Categories  apijson.Field
	Title       apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarSpecialeventGetResponseResultSpecialEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarSpecialeventGetResponseResultSpecialEventCategory struct {
	CategoryAlias string                                                     `json:"categoryAlias,required"`
	Location      string                                                     `json:"location,required"`
	JSON          radarSpecialeventGetResponseResultSpecialEventCategoryJSON `json:"-"`
}

// radarSpecialeventGetResponseResultSpecialEventCategoryJSON contains the JSON
// metadata for the struct [RadarSpecialeventGetResponseResultSpecialEventCategory]
type radarSpecialeventGetResponseResultSpecialEventCategoryJSON struct {
	CategoryAlias apijson.Field
	Location      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RadarSpecialeventGetResponseResultSpecialEventCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarSpecialeventListResponse struct {
	Result  RadarSpecialeventListResponseResult `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    radarSpecialeventListResponseJSON   `json:"-"`
}

// radarSpecialeventListResponseJSON contains the JSON metadata for the struct
// [RadarSpecialeventListResponse]
type radarSpecialeventListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarSpecialeventListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarSpecialeventListResponseResult struct {
	Serie0 RadarSpecialeventListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarSpecialeventListResponseResultJSON   `json:"-"`
}

// radarSpecialeventListResponseResultJSON contains the JSON metadata for the
// struct [RadarSpecialeventListResponseResult]
type radarSpecialeventListResponseResultJSON struct {
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarSpecialeventListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarSpecialeventListResponseResultSerie0 struct {
	Timestamps []time.Time                                   `json:"timestamps,required" format:"date-time"`
	Values     []string                                      `json:"values,required"`
	JSON       radarSpecialeventListResponseResultSerie0JSON `json:"-"`
}

// radarSpecialeventListResponseResultSerie0JSON contains the JSON metadata for the
// struct [RadarSpecialeventListResponseResultSerie0]
type radarSpecialeventListResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarSpecialeventListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarSpecialeventGetParams struct {
	// Format results are returned in.
	Format param.Field[RadarSpecialeventGetParamsFormat] `query:"format"`
}

// URLQuery serializes [RadarSpecialeventGetParams]'s query parameters as
// `url.Values`.
func (r RadarSpecialeventGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarSpecialeventGetParamsFormat string

const (
	RadarSpecialeventGetParamsFormatJson RadarSpecialeventGetParamsFormat = "JSON"
	RadarSpecialeventGetParamsFormatCsv  RadarSpecialeventGetParamsFormat = "CSV"
)

type RadarSpecialeventListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarSpecialeventListParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarSpecialeventListParamsFormat] `query:"format"`
	// Location Alpha2 code.
	Location param.Field[string] `query:"location"`
	// Name that will be used to name the series in response.
	Name param.Field[string] `query:"name"`
}

// URLQuery serializes [RadarSpecialeventListParams]'s query parameters as
// `url.Values`.
func (r RadarSpecialeventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarSpecialeventListParamsAggInterval string

const (
	RadarSpecialeventListParamsAggInterval15m RadarSpecialeventListParamsAggInterval = "15m"
	RadarSpecialeventListParamsAggInterval1h  RadarSpecialeventListParamsAggInterval = "1h"
	RadarSpecialeventListParamsAggInterval1d  RadarSpecialeventListParamsAggInterval = "1d"
	RadarSpecialeventListParamsAggInterval1w  RadarSpecialeventListParamsAggInterval = "1w"
)

// Format results are returned in.
type RadarSpecialeventListParamsFormat string

const (
	RadarSpecialeventListParamsFormatJson RadarSpecialeventListParamsFormat = "JSON"
	RadarSpecialeventListParamsFormatCsv  RadarSpecialeventListParamsFormat = "CSV"
)
