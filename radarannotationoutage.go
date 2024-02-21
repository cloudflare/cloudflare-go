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

// RadarAnnotationOutageService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarAnnotationOutageService]
// method instead.
type RadarAnnotationOutageService struct {
	Options []option.RequestOption
}

// NewRadarAnnotationOutageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAnnotationOutageService(opts ...option.RequestOption) (r *RadarAnnotationOutageService) {
	r = &RadarAnnotationOutageService{}
	r.Options = opts
	return
}

// Get the number of outages for locations.
func (r *RadarAnnotationOutageService) Locations(ctx context.Context, query RadarAnnotationOutageLocationsParams, opts ...option.RequestOption) (res *RadarAnnotationOutageLocationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAnnotationOutageLocationsResponseEnvelope
	path := "radar/annotations/outages/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAnnotationOutageLocationsResponse struct {
	Annotations []RadarAnnotationOutageLocationsResponseAnnotation `json:"annotations,required"`
	JSON        radarAnnotationOutageLocationsResponseJSON         `json:"-"`
}

// radarAnnotationOutageLocationsResponseJSON contains the JSON metadata for the
// struct [RadarAnnotationOutageLocationsResponse]
type radarAnnotationOutageLocationsResponseJSON struct {
	Annotations apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAnnotationOutageLocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAnnotationOutageLocationsResponseAnnotation struct {
	ClientCountryAlpha2 string                                               `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                               `json:"clientCountryName,required"`
	Value               string                                               `json:"value,required"`
	JSON                radarAnnotationOutageLocationsResponseAnnotationJSON `json:"-"`
}

// radarAnnotationOutageLocationsResponseAnnotationJSON contains the JSON metadata
// for the struct [RadarAnnotationOutageLocationsResponseAnnotation]
type radarAnnotationOutageLocationsResponseAnnotationJSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAnnotationOutageLocationsResponseAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAnnotationOutageLocationsParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Shorthand date ranges for the last X days - use when you don't need specific
	// start and end dates.
	DateRange param.Field[RadarAnnotationOutageLocationsParamsDateRange] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAnnotationOutageLocationsParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [RadarAnnotationOutageLocationsParams]'s query parameters as
// `url.Values`.
func (r RadarAnnotationOutageLocationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Shorthand date ranges for the last X days - use when you don't need specific
// start and end dates.
type RadarAnnotationOutageLocationsParamsDateRange string

const (
	RadarAnnotationOutageLocationsParamsDateRange1d         RadarAnnotationOutageLocationsParamsDateRange = "1d"
	RadarAnnotationOutageLocationsParamsDateRange2d         RadarAnnotationOutageLocationsParamsDateRange = "2d"
	RadarAnnotationOutageLocationsParamsDateRange7d         RadarAnnotationOutageLocationsParamsDateRange = "7d"
	RadarAnnotationOutageLocationsParamsDateRange14d        RadarAnnotationOutageLocationsParamsDateRange = "14d"
	RadarAnnotationOutageLocationsParamsDateRange28d        RadarAnnotationOutageLocationsParamsDateRange = "28d"
	RadarAnnotationOutageLocationsParamsDateRange12w        RadarAnnotationOutageLocationsParamsDateRange = "12w"
	RadarAnnotationOutageLocationsParamsDateRange24w        RadarAnnotationOutageLocationsParamsDateRange = "24w"
	RadarAnnotationOutageLocationsParamsDateRange52w        RadarAnnotationOutageLocationsParamsDateRange = "52w"
	RadarAnnotationOutageLocationsParamsDateRange1dControl  RadarAnnotationOutageLocationsParamsDateRange = "1dControl"
	RadarAnnotationOutageLocationsParamsDateRange2dControl  RadarAnnotationOutageLocationsParamsDateRange = "2dControl"
	RadarAnnotationOutageLocationsParamsDateRange7dControl  RadarAnnotationOutageLocationsParamsDateRange = "7dControl"
	RadarAnnotationOutageLocationsParamsDateRange14dControl RadarAnnotationOutageLocationsParamsDateRange = "14dControl"
	RadarAnnotationOutageLocationsParamsDateRange28dControl RadarAnnotationOutageLocationsParamsDateRange = "28dControl"
	RadarAnnotationOutageLocationsParamsDateRange12wControl RadarAnnotationOutageLocationsParamsDateRange = "12wControl"
	RadarAnnotationOutageLocationsParamsDateRange24wControl RadarAnnotationOutageLocationsParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAnnotationOutageLocationsParamsFormat string

const (
	RadarAnnotationOutageLocationsParamsFormatJson RadarAnnotationOutageLocationsParamsFormat = "JSON"
	RadarAnnotationOutageLocationsParamsFormatCsv  RadarAnnotationOutageLocationsParamsFormat = "CSV"
)

type RadarAnnotationOutageLocationsResponseEnvelope struct {
	Result  RadarAnnotationOutageLocationsResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    radarAnnotationOutageLocationsResponseEnvelopeJSON `json:"-"`
}

// radarAnnotationOutageLocationsResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarAnnotationOutageLocationsResponseEnvelope]
type radarAnnotationOutageLocationsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAnnotationOutageLocationsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
