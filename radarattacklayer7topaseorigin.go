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

// RadarAttackLayer7TopAseOriginService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer7TopAseOriginService] method instead.
type RadarAttackLayer7TopAseOriginService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer7TopAseOriginService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer7TopAseOriginService(opts ...option.RequestOption) (r *RadarAttackLayer7TopAseOriginService) {
	r = &RadarAttackLayer7TopAseOriginService{}
	r.Options = opts
	return
}

// Get the top origin ASes by layer 7 attacks. Values are a percentage out of the
// total layer 7 attacks. The origin location is determined by the client IP.
func (r *RadarAttackLayer7TopAseOriginService) List(ctx context.Context, query RadarAttackLayer7TopAseOriginListParams, opts ...option.RequestOption) (res *RadarAttackLayer7TopAseOriginListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/top/ases/origin"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer7TopAseOriginListResponse struct {
	Result  RadarAttackLayer7TopAseOriginListResponseResult `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarAttackLayer7TopAseOriginListResponseJSON   `json:"-"`
}

// radarAttackLayer7TopAseOriginListResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer7TopAseOriginListResponse]
type radarAttackLayer7TopAseOriginListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopAseOriginListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopAseOriginListResponseResult struct {
	Top0 []RadarAttackLayer7TopAseOriginListResponseResultTop0 `json:"top_0,required"`
	JSON radarAttackLayer7TopAseOriginListResponseResultJSON   `json:"-"`
}

// radarAttackLayer7TopAseOriginListResponseResultJSON contains the JSON metadata
// for the struct [RadarAttackLayer7TopAseOriginListResponseResult]
type radarAttackLayer7TopAseOriginListResponseResultJSON struct {
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopAseOriginListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopAseOriginListResponseResultTop0 struct {
	OriginASN     float64                                                 `json:"originAsn,required"`
	OriginASNName string                                                  `json:"originAsnName,required"`
	Rank          int64                                                   `json:"rank,required"`
	Value         string                                                  `json:"value,required"`
	JSON          radarAttackLayer7TopAseOriginListResponseResultTop0JSON `json:"-"`
}

// radarAttackLayer7TopAseOriginListResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarAttackLayer7TopAseOriginListResponseResultTop0]
type radarAttackLayer7TopAseOriginListResponseResultTop0JSON struct {
	OriginASN     apijson.Field
	OriginASNName apijson.Field
	Rank          apijson.Field
	Value         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RadarAttackLayer7TopAseOriginListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopAseOriginListParams struct {
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7TopAseOriginListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7TopAseOriginListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of locations (alpha-2 country codes).
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7TopAseOriginListParams]'s query parameters
// as `url.Values`.
func (r RadarAttackLayer7TopAseOriginListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer7TopAseOriginListParamsDateRange string

const (
	RadarAttackLayer7TopAseOriginListParamsDateRange1d         RadarAttackLayer7TopAseOriginListParamsDateRange = "1d"
	RadarAttackLayer7TopAseOriginListParamsDateRange7d         RadarAttackLayer7TopAseOriginListParamsDateRange = "7d"
	RadarAttackLayer7TopAseOriginListParamsDateRange14d        RadarAttackLayer7TopAseOriginListParamsDateRange = "14d"
	RadarAttackLayer7TopAseOriginListParamsDateRange28d        RadarAttackLayer7TopAseOriginListParamsDateRange = "28d"
	RadarAttackLayer7TopAseOriginListParamsDateRange12w        RadarAttackLayer7TopAseOriginListParamsDateRange = "12w"
	RadarAttackLayer7TopAseOriginListParamsDateRange24w        RadarAttackLayer7TopAseOriginListParamsDateRange = "24w"
	RadarAttackLayer7TopAseOriginListParamsDateRange52w        RadarAttackLayer7TopAseOriginListParamsDateRange = "52w"
	RadarAttackLayer7TopAseOriginListParamsDateRange1dControl  RadarAttackLayer7TopAseOriginListParamsDateRange = "1dControl"
	RadarAttackLayer7TopAseOriginListParamsDateRange7dControl  RadarAttackLayer7TopAseOriginListParamsDateRange = "7dControl"
	RadarAttackLayer7TopAseOriginListParamsDateRange14dControl RadarAttackLayer7TopAseOriginListParamsDateRange = "14dControl"
	RadarAttackLayer7TopAseOriginListParamsDateRange28dControl RadarAttackLayer7TopAseOriginListParamsDateRange = "28dControl"
	RadarAttackLayer7TopAseOriginListParamsDateRange12wControl RadarAttackLayer7TopAseOriginListParamsDateRange = "12wControl"
	RadarAttackLayer7TopAseOriginListParamsDateRange24wControl RadarAttackLayer7TopAseOriginListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7TopAseOriginListParamsFormat string

const (
	RadarAttackLayer7TopAseOriginListParamsFormatJson RadarAttackLayer7TopAseOriginListParamsFormat = "JSON"
	RadarAttackLayer7TopAseOriginListParamsFormatCsv  RadarAttackLayer7TopAseOriginListParamsFormat = "CSV"
)
