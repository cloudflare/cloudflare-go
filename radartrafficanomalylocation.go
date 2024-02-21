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
func (r *RadarTrafficAnomalyLocationService) Get(ctx context.Context, query RadarTrafficAnomalyLocationGetParams, opts ...option.RequestOption) (res *RadarTrafficAnomalyLocationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarTrafficAnomalyLocationGetResponseEnvelope
	path := "radar/traffic_anomalies/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarTrafficAnomalyLocationGetResponse struct {
	TrafficAnomalies []RadarTrafficAnomalyLocationGetResponseTrafficAnomaly `json:"trafficAnomalies,required"`
	JSON             radarTrafficAnomalyLocationGetResponseJSON             `json:"-"`
}

// radarTrafficAnomalyLocationGetResponseJSON contains the JSON metadata for the
// struct [RadarTrafficAnomalyLocationGetResponse]
type radarTrafficAnomalyLocationGetResponseJSON struct {
	TrafficAnomalies apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RadarTrafficAnomalyLocationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarTrafficAnomalyLocationGetResponseTrafficAnomaly struct {
	ClientCountryAlpha2 string                                                   `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                                   `json:"clientCountryName,required"`
	Value               string                                                   `json:"value,required"`
	JSON                radarTrafficAnomalyLocationGetResponseTrafficAnomalyJSON `json:"-"`
}

// radarTrafficAnomalyLocationGetResponseTrafficAnomalyJSON contains the JSON
// metadata for the struct [RadarTrafficAnomalyLocationGetResponseTrafficAnomaly]
type radarTrafficAnomalyLocationGetResponseTrafficAnomalyJSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarTrafficAnomalyLocationGetResponseTrafficAnomaly) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarTrafficAnomalyLocationGetParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Shorthand date ranges for the last X days - use when you don't need specific
	// start and end dates.
	DateRange param.Field[RadarTrafficAnomalyLocationGetParamsDateRange] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarTrafficAnomalyLocationGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit  param.Field[int64]                                      `query:"limit"`
	Status param.Field[RadarTrafficAnomalyLocationGetParamsStatus] `query:"status"`
}

// URLQuery serializes [RadarTrafficAnomalyLocationGetParams]'s query parameters as
// `url.Values`.
func (r RadarTrafficAnomalyLocationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Shorthand date ranges for the last X days - use when you don't need specific
// start and end dates.
type RadarTrafficAnomalyLocationGetParamsDateRange string

const (
	RadarTrafficAnomalyLocationGetParamsDateRange1d         RadarTrafficAnomalyLocationGetParamsDateRange = "1d"
	RadarTrafficAnomalyLocationGetParamsDateRange2d         RadarTrafficAnomalyLocationGetParamsDateRange = "2d"
	RadarTrafficAnomalyLocationGetParamsDateRange7d         RadarTrafficAnomalyLocationGetParamsDateRange = "7d"
	RadarTrafficAnomalyLocationGetParamsDateRange14d        RadarTrafficAnomalyLocationGetParamsDateRange = "14d"
	RadarTrafficAnomalyLocationGetParamsDateRange28d        RadarTrafficAnomalyLocationGetParamsDateRange = "28d"
	RadarTrafficAnomalyLocationGetParamsDateRange12w        RadarTrafficAnomalyLocationGetParamsDateRange = "12w"
	RadarTrafficAnomalyLocationGetParamsDateRange24w        RadarTrafficAnomalyLocationGetParamsDateRange = "24w"
	RadarTrafficAnomalyLocationGetParamsDateRange52w        RadarTrafficAnomalyLocationGetParamsDateRange = "52w"
	RadarTrafficAnomalyLocationGetParamsDateRange1dControl  RadarTrafficAnomalyLocationGetParamsDateRange = "1dControl"
	RadarTrafficAnomalyLocationGetParamsDateRange2dControl  RadarTrafficAnomalyLocationGetParamsDateRange = "2dControl"
	RadarTrafficAnomalyLocationGetParamsDateRange7dControl  RadarTrafficAnomalyLocationGetParamsDateRange = "7dControl"
	RadarTrafficAnomalyLocationGetParamsDateRange14dControl RadarTrafficAnomalyLocationGetParamsDateRange = "14dControl"
	RadarTrafficAnomalyLocationGetParamsDateRange28dControl RadarTrafficAnomalyLocationGetParamsDateRange = "28dControl"
	RadarTrafficAnomalyLocationGetParamsDateRange12wControl RadarTrafficAnomalyLocationGetParamsDateRange = "12wControl"
	RadarTrafficAnomalyLocationGetParamsDateRange24wControl RadarTrafficAnomalyLocationGetParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarTrafficAnomalyLocationGetParamsFormat string

const (
	RadarTrafficAnomalyLocationGetParamsFormatJson RadarTrafficAnomalyLocationGetParamsFormat = "JSON"
	RadarTrafficAnomalyLocationGetParamsFormatCsv  RadarTrafficAnomalyLocationGetParamsFormat = "CSV"
)

type RadarTrafficAnomalyLocationGetParamsStatus string

const (
	RadarTrafficAnomalyLocationGetParamsStatusVerified   RadarTrafficAnomalyLocationGetParamsStatus = "VERIFIED"
	RadarTrafficAnomalyLocationGetParamsStatusUnverified RadarTrafficAnomalyLocationGetParamsStatus = "UNVERIFIED"
)

type RadarTrafficAnomalyLocationGetResponseEnvelope struct {
	Result  RadarTrafficAnomalyLocationGetResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    radarTrafficAnomalyLocationGetResponseEnvelopeJSON `json:"-"`
}

// radarTrafficAnomalyLocationGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarTrafficAnomalyLocationGetResponseEnvelope]
type radarTrafficAnomalyLocationGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTrafficAnomalyLocationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
