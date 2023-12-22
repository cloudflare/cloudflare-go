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

// RadarAttackLayer7TopAttackService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarAttackLayer7TopAttackService] method instead.
type RadarAttackLayer7TopAttackService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer7TopAttackService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer7TopAttackService(opts ...option.RequestOption) (r *RadarAttackLayer7TopAttackService) {
	r = &RadarAttackLayer7TopAttackService{}
	r.Options = opts
	return
}

// Get the top attacks from origin to target location. Values are a percentage out
// of the total layer 7 attacks (with billing country). The attack magnitude can be
// defined by the number of mitigated requests or by the number of zones affected.
// You can optionally limit the number of attacks per origin/target location
// (useful if all the top attacks are from or to the same location).
func (r *RadarAttackLayer7TopAttackService) List(ctx context.Context, query RadarAttackLayer7TopAttackListParams, opts ...option.RequestOption) (res *RadarAttackLayer7TopAttackListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/top/attacks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer7TopAttackListResponse struct {
	Result  RadarAttackLayer7TopAttackListResponseResult `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarAttackLayer7TopAttackListResponseJSON   `json:"-"`
}

// radarAttackLayer7TopAttackListResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer7TopAttackListResponse]
type radarAttackLayer7TopAttackListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopAttackListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopAttackListResponseResult struct {
	Top0 []RadarAttackLayer7TopAttackListResponseResultTop0 `json:"top_0,required"`
	JSON radarAttackLayer7TopAttackListResponseResultJSON   `json:"-"`
}

// radarAttackLayer7TopAttackListResponseResultJSON contains the JSON metadata for
// the struct [RadarAttackLayer7TopAttackListResponseResult]
type radarAttackLayer7TopAttackListResponseResultJSON struct {
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopAttackListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopAttackListResponseResultTop0 struct {
	OriginCountryAlpha2 string                                               `json:"originCountryAlpha2,required"`
	OriginCountryName   string                                               `json:"originCountryName,required"`
	Rank                int64                                                `json:"rank,required"`
	TargetCountryAlpha2 string                                               `json:"targetCountryAlpha2,required"`
	TargetCountryName   string                                               `json:"targetCountryName,required"`
	Value               string                                               `json:"value,required"`
	JSON                radarAttackLayer7TopAttackListResponseResultTop0JSON `json:"-"`
}

// radarAttackLayer7TopAttackListResponseResultTop0JSON contains the JSON metadata
// for the struct [RadarAttackLayer7TopAttackListResponseResultTop0]
type radarAttackLayer7TopAttackListResponseResultTop0JSON struct {
	OriginCountryAlpha2 apijson.Field
	OriginCountryName   apijson.Field
	Rank                apijson.Field
	TargetCountryAlpha2 apijson.Field
	TargetCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAttackLayer7TopAttackListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopAttackListParams struct {
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7TopAttackListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7TopAttackListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of attack origin/target location attack limits. Together with
	// `limitPerLocation`, limits how many objects will be fetched per origin/target
	// location.
	LimitDirection param.Field[RadarAttackLayer7TopAttackListParamsLimitDirection] `query:"limitDirection"`
	// Limit the number of attacks per origin/target (refer to `limitDirection`
	// parameter) location.
	LimitPerLocation param.Field[int64] `query:"limitPerLocation"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Attack magnitude can be defined by total requests mitigated or by total zones
	// attacked.
	Magnitude param.Field[RadarAttackLayer7TopAttackListParamsMagnitude] `query:"magnitude"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7TopAttackListParams]'s query parameters as
// `url.Values`.
func (r RadarAttackLayer7TopAttackListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer7TopAttackListParamsDateRange string

const (
	RadarAttackLayer7TopAttackListParamsDateRange1d         RadarAttackLayer7TopAttackListParamsDateRange = "1d"
	RadarAttackLayer7TopAttackListParamsDateRange7d         RadarAttackLayer7TopAttackListParamsDateRange = "7d"
	RadarAttackLayer7TopAttackListParamsDateRange14d        RadarAttackLayer7TopAttackListParamsDateRange = "14d"
	RadarAttackLayer7TopAttackListParamsDateRange28d        RadarAttackLayer7TopAttackListParamsDateRange = "28d"
	RadarAttackLayer7TopAttackListParamsDateRange12w        RadarAttackLayer7TopAttackListParamsDateRange = "12w"
	RadarAttackLayer7TopAttackListParamsDateRange24w        RadarAttackLayer7TopAttackListParamsDateRange = "24w"
	RadarAttackLayer7TopAttackListParamsDateRange52w        RadarAttackLayer7TopAttackListParamsDateRange = "52w"
	RadarAttackLayer7TopAttackListParamsDateRange1dControl  RadarAttackLayer7TopAttackListParamsDateRange = "1dControl"
	RadarAttackLayer7TopAttackListParamsDateRange7dControl  RadarAttackLayer7TopAttackListParamsDateRange = "7dControl"
	RadarAttackLayer7TopAttackListParamsDateRange14dControl RadarAttackLayer7TopAttackListParamsDateRange = "14dControl"
	RadarAttackLayer7TopAttackListParamsDateRange28dControl RadarAttackLayer7TopAttackListParamsDateRange = "28dControl"
	RadarAttackLayer7TopAttackListParamsDateRange12wControl RadarAttackLayer7TopAttackListParamsDateRange = "12wControl"
	RadarAttackLayer7TopAttackListParamsDateRange24wControl RadarAttackLayer7TopAttackListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7TopAttackListParamsFormat string

const (
	RadarAttackLayer7TopAttackListParamsFormatJson RadarAttackLayer7TopAttackListParamsFormat = "JSON"
	RadarAttackLayer7TopAttackListParamsFormatCsv  RadarAttackLayer7TopAttackListParamsFormat = "CSV"
)

// Array of attack origin/target location attack limits. Together with
// `limitPerLocation`, limits how many objects will be fetched per origin/target
// location.
type RadarAttackLayer7TopAttackListParamsLimitDirection string

const (
	RadarAttackLayer7TopAttackListParamsLimitDirectionOrigin RadarAttackLayer7TopAttackListParamsLimitDirection = "ORIGIN"
	RadarAttackLayer7TopAttackListParamsLimitDirectionTarget RadarAttackLayer7TopAttackListParamsLimitDirection = "TARGET"
)

// Attack magnitude can be defined by total requests mitigated or by total zones
// attacked.
type RadarAttackLayer7TopAttackListParamsMagnitude string

const (
	RadarAttackLayer7TopAttackListParamsMagnitudeAffectedZones     RadarAttackLayer7TopAttackListParamsMagnitude = "AFFECTED_ZONES"
	RadarAttackLayer7TopAttackListParamsMagnitudeMitigatedRequests RadarAttackLayer7TopAttackListParamsMagnitude = "MITIGATED_REQUESTS"
)
