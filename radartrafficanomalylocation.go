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

// RadarTrafficAnomalyLocationService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarTrafficAnomalyLocationService] method instead.
type RadarTrafficAnomalyLocationService struct {
	Options []option.RequestOption
}

// NewRadarTrafficAnomalyLocationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarTrafficAnomalyLocationService(opts ...option.RequestOption) (r *RadarTrafficAnomalyLocationService) {
	r = &RadarTrafficAnomalyLocationService{}
	r.Options = opts
	return
}

// Internet traffic anomalies are signals that might point to an outage, These
// alerts are automatically detected by Radar and then manually verified by our
// team. This endpoint returns the sum of alerts grouped by location.
func (r *RadarTrafficAnomalyLocationService) List(ctx context.Context, query RadarTrafficAnomalyLocationListParams, opts ...option.RequestOption) (res *RadarTrafficAnomalyLocationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/traffic_anomalies/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarTrafficAnomalyLocationListResponse struct {
	Result  RadarTrafficAnomalyLocationListResponseResult `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    radarTrafficAnomalyLocationListResponseJSON   `json:"-"`
}

// radarTrafficAnomalyLocationListResponseJSON contains the JSON metadata for the
// struct [RadarTrafficAnomalyLocationListResponse]
type radarTrafficAnomalyLocationListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTrafficAnomalyLocationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarTrafficAnomalyLocationListResponseResult struct {
	TrafficAnomalies []RadarTrafficAnomalyLocationListResponseResultTrafficAnomaly `json:"trafficAnomalies,required"`
	JSON             radarTrafficAnomalyLocationListResponseResultJSON             `json:"-"`
}

// radarTrafficAnomalyLocationListResponseResultJSON contains the JSON metadata for
// the struct [RadarTrafficAnomalyLocationListResponseResult]
type radarTrafficAnomalyLocationListResponseResultJSON struct {
	TrafficAnomalies apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RadarTrafficAnomalyLocationListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarTrafficAnomalyLocationListResponseResultTrafficAnomaly struct {
	ClientCountryAlpha2 string                                                          `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                                          `json:"clientCountryName,required"`
	Value               string                                                          `json:"value,required"`
	JSON                radarTrafficAnomalyLocationListResponseResultTrafficAnomalyJSON `json:"-"`
}

// radarTrafficAnomalyLocationListResponseResultTrafficAnomalyJSON contains the
// JSON metadata for the struct
// [RadarTrafficAnomalyLocationListResponseResultTrafficAnomaly]
type radarTrafficAnomalyLocationListResponseResultTrafficAnomalyJSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarTrafficAnomalyLocationListResponseResultTrafficAnomaly) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarTrafficAnomalyLocationListParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Shorthand date ranges for the last X days - use when you don't need specific
	// start and end dates.
	DateRange param.Field[RadarTrafficAnomalyLocationListParamsDateRange] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarTrafficAnomalyLocationListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit  param.Field[int64]                                       `query:"limit"`
	Status param.Field[RadarTrafficAnomalyLocationListParamsStatus] `query:"status"`
}

// URLQuery serializes [RadarTrafficAnomalyLocationListParams]'s query parameters
// as `url.Values`.
func (r RadarTrafficAnomalyLocationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Shorthand date ranges for the last X days - use when you don't need specific
// start and end dates.
type RadarTrafficAnomalyLocationListParamsDateRange string

const (
	RadarTrafficAnomalyLocationListParamsDateRange1d         RadarTrafficAnomalyLocationListParamsDateRange = "1d"
	RadarTrafficAnomalyLocationListParamsDateRange2d         RadarTrafficAnomalyLocationListParamsDateRange = "2d"
	RadarTrafficAnomalyLocationListParamsDateRange7d         RadarTrafficAnomalyLocationListParamsDateRange = "7d"
	RadarTrafficAnomalyLocationListParamsDateRange14d        RadarTrafficAnomalyLocationListParamsDateRange = "14d"
	RadarTrafficAnomalyLocationListParamsDateRange28d        RadarTrafficAnomalyLocationListParamsDateRange = "28d"
	RadarTrafficAnomalyLocationListParamsDateRange12w        RadarTrafficAnomalyLocationListParamsDateRange = "12w"
	RadarTrafficAnomalyLocationListParamsDateRange24w        RadarTrafficAnomalyLocationListParamsDateRange = "24w"
	RadarTrafficAnomalyLocationListParamsDateRange52w        RadarTrafficAnomalyLocationListParamsDateRange = "52w"
	RadarTrafficAnomalyLocationListParamsDateRange1dControl  RadarTrafficAnomalyLocationListParamsDateRange = "1dControl"
	RadarTrafficAnomalyLocationListParamsDateRange2dControl  RadarTrafficAnomalyLocationListParamsDateRange = "2dControl"
	RadarTrafficAnomalyLocationListParamsDateRange7dControl  RadarTrafficAnomalyLocationListParamsDateRange = "7dControl"
	RadarTrafficAnomalyLocationListParamsDateRange14dControl RadarTrafficAnomalyLocationListParamsDateRange = "14dControl"
	RadarTrafficAnomalyLocationListParamsDateRange28dControl RadarTrafficAnomalyLocationListParamsDateRange = "28dControl"
	RadarTrafficAnomalyLocationListParamsDateRange12wControl RadarTrafficAnomalyLocationListParamsDateRange = "12wControl"
	RadarTrafficAnomalyLocationListParamsDateRange24wControl RadarTrafficAnomalyLocationListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarTrafficAnomalyLocationListParamsFormat string

const (
	RadarTrafficAnomalyLocationListParamsFormatJson RadarTrafficAnomalyLocationListParamsFormat = "JSON"
	RadarTrafficAnomalyLocationListParamsFormatCsv  RadarTrafficAnomalyLocationListParamsFormat = "CSV"
)

type RadarTrafficAnomalyLocationListParamsStatus string

const (
	RadarTrafficAnomalyLocationListParamsStatusVerified   RadarTrafficAnomalyLocationListParamsStatus = "VERIFIED"
	RadarTrafficAnomalyLocationListParamsStatusUnverified RadarTrafficAnomalyLocationListParamsStatus = "UNVERIFIED"
)
