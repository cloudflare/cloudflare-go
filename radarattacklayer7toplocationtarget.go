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

// RadarAttackLayer7TopLocationTargetService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer7TopLocationTargetService] method instead.
type RadarAttackLayer7TopLocationTargetService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer7TopLocationTargetService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarAttackLayer7TopLocationTargetService(opts ...option.RequestOption) (r *RadarAttackLayer7TopLocationTargetService) {
	r = &RadarAttackLayer7TopLocationTargetService{}
	r.Options = opts
	return
}

// Get the top target locations of and by layer 7 attacks. Values are a percentage
// out of the total layer 7 attacks. The target location is determined by the
// attacked zone's billing country, when available.
func (r *RadarAttackLayer7TopLocationTargetService) List(ctx context.Context, query RadarAttackLayer7TopLocationTargetListParams, opts ...option.RequestOption) (res *RadarAttackLayer7TopLocationTargetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/top/locations/target"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer7TopLocationTargetListResponse struct {
	Result  RadarAttackLayer7TopLocationTargetListResponseResult `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarAttackLayer7TopLocationTargetListResponseJSON   `json:"-"`
}

// radarAttackLayer7TopLocationTargetListResponseJSON contains the JSON metadata
// for the struct [RadarAttackLayer7TopLocationTargetListResponse]
type radarAttackLayer7TopLocationTargetListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationTargetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopLocationTargetListResponseResult struct {
	Top0 []RadarAttackLayer7TopLocationTargetListResponseResultTop0 `json:"top_0,required"`
	JSON radarAttackLayer7TopLocationTargetListResponseResultJSON   `json:"-"`
}

// radarAttackLayer7TopLocationTargetListResponseResultJSON contains the JSON
// metadata for the struct [RadarAttackLayer7TopLocationTargetListResponseResult]
type radarAttackLayer7TopLocationTargetListResponseResultJSON struct {
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationTargetListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopLocationTargetListResponseResultTop0 struct {
	Rank                int64                                                        `json:"rank,required"`
	TargetCountryAlpha2 string                                                       `json:"targetCountryAlpha2,required"`
	TargetCountryName   string                                                       `json:"targetCountryName,required"`
	Value               string                                                       `json:"value,required"`
	JSON                radarAttackLayer7TopLocationTargetListResponseResultTop0JSON `json:"-"`
}

// radarAttackLayer7TopLocationTargetListResponseResultTop0JSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7TopLocationTargetListResponseResultTop0]
type radarAttackLayer7TopLocationTargetListResponseResultTop0JSON struct {
	Rank                apijson.Field
	TargetCountryAlpha2 apijson.Field
	TargetCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationTargetListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopLocationTargetListParams struct {
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7TopLocationTargetListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7TopLocationTargetListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7TopLocationTargetListParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7TopLocationTargetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer7TopLocationTargetListParamsDateRange string

const (
	RadarAttackLayer7TopLocationTargetListParamsDateRange1d         RadarAttackLayer7TopLocationTargetListParamsDateRange = "1d"
	RadarAttackLayer7TopLocationTargetListParamsDateRange7d         RadarAttackLayer7TopLocationTargetListParamsDateRange = "7d"
	RadarAttackLayer7TopLocationTargetListParamsDateRange14d        RadarAttackLayer7TopLocationTargetListParamsDateRange = "14d"
	RadarAttackLayer7TopLocationTargetListParamsDateRange28d        RadarAttackLayer7TopLocationTargetListParamsDateRange = "28d"
	RadarAttackLayer7TopLocationTargetListParamsDateRange12w        RadarAttackLayer7TopLocationTargetListParamsDateRange = "12w"
	RadarAttackLayer7TopLocationTargetListParamsDateRange24w        RadarAttackLayer7TopLocationTargetListParamsDateRange = "24w"
	RadarAttackLayer7TopLocationTargetListParamsDateRange52w        RadarAttackLayer7TopLocationTargetListParamsDateRange = "52w"
	RadarAttackLayer7TopLocationTargetListParamsDateRange1dControl  RadarAttackLayer7TopLocationTargetListParamsDateRange = "1dControl"
	RadarAttackLayer7TopLocationTargetListParamsDateRange7dControl  RadarAttackLayer7TopLocationTargetListParamsDateRange = "7dControl"
	RadarAttackLayer7TopLocationTargetListParamsDateRange14dControl RadarAttackLayer7TopLocationTargetListParamsDateRange = "14dControl"
	RadarAttackLayer7TopLocationTargetListParamsDateRange28dControl RadarAttackLayer7TopLocationTargetListParamsDateRange = "28dControl"
	RadarAttackLayer7TopLocationTargetListParamsDateRange12wControl RadarAttackLayer7TopLocationTargetListParamsDateRange = "12wControl"
	RadarAttackLayer7TopLocationTargetListParamsDateRange24wControl RadarAttackLayer7TopLocationTargetListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7TopLocationTargetListParamsFormat string

const (
	RadarAttackLayer7TopLocationTargetListParamsFormatJson RadarAttackLayer7TopLocationTargetListParamsFormat = "JSON"
	RadarAttackLayer7TopLocationTargetListParamsFormatCsv  RadarAttackLayer7TopLocationTargetListParamsFormat = "CSV"
)
