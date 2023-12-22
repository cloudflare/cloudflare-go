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

// RadarAnnotationOutageLocationService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAnnotationOutageLocationService] method instead.
type RadarAnnotationOutageLocationService struct {
	Options []option.RequestOption
}

// NewRadarAnnotationOutageLocationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAnnotationOutageLocationService(opts ...option.RequestOption) (r *RadarAnnotationOutageLocationService) {
	r = &RadarAnnotationOutageLocationService{}
	r.Options = opts
	return
}

// Get top outages annotations
func (r *RadarAnnotationOutageLocationService) List(ctx context.Context, query RadarAnnotationOutageLocationListParams, opts ...option.RequestOption) (res *RadarAnnotationOutageLocationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/annotations/outages/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAnnotationOutageLocationListResponse struct {
	Result  RadarAnnotationOutageLocationListResponseResult `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarAnnotationOutageLocationListResponseJSON   `json:"-"`
}

// radarAnnotationOutageLocationListResponseJSON contains the JSON metadata for the
// struct [RadarAnnotationOutageLocationListResponse]
type radarAnnotationOutageLocationListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAnnotationOutageLocationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAnnotationOutageLocationListResponseResult struct {
	Annotations []RadarAnnotationOutageLocationListResponseResultAnnotation `json:"annotations,required"`
	JSON        radarAnnotationOutageLocationListResponseResultJSON         `json:"-"`
}

// radarAnnotationOutageLocationListResponseResultJSON contains the JSON metadata
// for the struct [RadarAnnotationOutageLocationListResponseResult]
type radarAnnotationOutageLocationListResponseResultJSON struct {
	Annotations apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAnnotationOutageLocationListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAnnotationOutageLocationListResponseResultAnnotation struct {
	ClientCountryAlpha2 string                                                        `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                                        `json:"clientCountryName,required"`
	Value               string                                                        `json:"value,required"`
	JSON                radarAnnotationOutageLocationListResponseResultAnnotationJSON `json:"-"`
}

// radarAnnotationOutageLocationListResponseResultAnnotationJSON contains the JSON
// metadata for the struct
// [RadarAnnotationOutageLocationListResponseResultAnnotation]
type radarAnnotationOutageLocationListResponseResultAnnotationJSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAnnotationOutageLocationListResponseResultAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAnnotationOutageLocationListParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Shorthand date ranges for the last X days - use when you don't need specific
	// start and end dates.
	DateRange param.Field[RadarAnnotationOutageLocationListParamsDateRange] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAnnotationOutageLocationListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [RadarAnnotationOutageLocationListParams]'s query parameters
// as `url.Values`.
func (r RadarAnnotationOutageLocationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Shorthand date ranges for the last X days - use when you don't need specific
// start and end dates.
type RadarAnnotationOutageLocationListParamsDateRange string

const (
	RadarAnnotationOutageLocationListParamsDateRange1d         RadarAnnotationOutageLocationListParamsDateRange = "1d"
	RadarAnnotationOutageLocationListParamsDateRange7d         RadarAnnotationOutageLocationListParamsDateRange = "7d"
	RadarAnnotationOutageLocationListParamsDateRange14d        RadarAnnotationOutageLocationListParamsDateRange = "14d"
	RadarAnnotationOutageLocationListParamsDateRange28d        RadarAnnotationOutageLocationListParamsDateRange = "28d"
	RadarAnnotationOutageLocationListParamsDateRange12w        RadarAnnotationOutageLocationListParamsDateRange = "12w"
	RadarAnnotationOutageLocationListParamsDateRange24w        RadarAnnotationOutageLocationListParamsDateRange = "24w"
	RadarAnnotationOutageLocationListParamsDateRange52w        RadarAnnotationOutageLocationListParamsDateRange = "52w"
	RadarAnnotationOutageLocationListParamsDateRange1dControl  RadarAnnotationOutageLocationListParamsDateRange = "1dControl"
	RadarAnnotationOutageLocationListParamsDateRange7dControl  RadarAnnotationOutageLocationListParamsDateRange = "7dControl"
	RadarAnnotationOutageLocationListParamsDateRange14dControl RadarAnnotationOutageLocationListParamsDateRange = "14dControl"
	RadarAnnotationOutageLocationListParamsDateRange28dControl RadarAnnotationOutageLocationListParamsDateRange = "28dControl"
	RadarAnnotationOutageLocationListParamsDateRange12wControl RadarAnnotationOutageLocationListParamsDateRange = "12wControl"
	RadarAnnotationOutageLocationListParamsDateRange24wControl RadarAnnotationOutageLocationListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAnnotationOutageLocationListParamsFormat string

const (
	RadarAnnotationOutageLocationListParamsFormatJson RadarAnnotationOutageLocationListParamsFormat = "JSON"
	RadarAnnotationOutageLocationListParamsFormatCsv  RadarAnnotationOutageLocationListParamsFormat = "CSV"
)
