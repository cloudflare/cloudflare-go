// File generated from our OpenAPI spec by Stainless.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// TrafficAnomalyLocationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewTrafficAnomalyLocationService]
// method instead.
type TrafficAnomalyLocationService struct {
	Options []option.RequestOption
}

// NewTrafficAnomalyLocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTrafficAnomalyLocationService(opts ...option.RequestOption) (r *TrafficAnomalyLocationService) {
	r = &TrafficAnomalyLocationService{}
	r.Options = opts
	return
}

// Internet traffic anomalies are signals that might point to an outage, These
// alerts are automatically detected by Radar and then manually verified by our
// team. This endpoint returns the sum of alerts grouped by location.
func (r *TrafficAnomalyLocationService) Get(ctx context.Context, query TrafficAnomalyLocationGetParams, opts ...option.RequestOption) (res *TrafficAnomalyLocationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TrafficAnomalyLocationGetResponseEnvelope
	path := "radar/traffic_anomalies/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TrafficAnomalyLocationGetResponse struct {
	TrafficAnomalies []TrafficAnomalyLocationGetResponseTrafficAnomaly `json:"trafficAnomalies,required"`
	JSON             trafficAnomalyLocationGetResponseJSON             `json:"-"`
}

// trafficAnomalyLocationGetResponseJSON contains the JSON metadata for the struct
// [TrafficAnomalyLocationGetResponse]
type trafficAnomalyLocationGetResponseJSON struct {
	TrafficAnomalies apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TrafficAnomalyLocationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r trafficAnomalyLocationGetResponseJSON) RawJSON() string {
	return r.raw
}

type TrafficAnomalyLocationGetResponseTrafficAnomaly struct {
	ClientCountryAlpha2 string                                              `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                              `json:"clientCountryName,required"`
	Value               string                                              `json:"value,required"`
	JSON                trafficAnomalyLocationGetResponseTrafficAnomalyJSON `json:"-"`
}

// trafficAnomalyLocationGetResponseTrafficAnomalyJSON contains the JSON metadata
// for the struct [TrafficAnomalyLocationGetResponseTrafficAnomaly]
type trafficAnomalyLocationGetResponseTrafficAnomalyJSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *TrafficAnomalyLocationGetResponseTrafficAnomaly) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r trafficAnomalyLocationGetResponseTrafficAnomalyJSON) RawJSON() string {
	return r.raw
}

type TrafficAnomalyLocationGetParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Shorthand date ranges for the last X days - use when you don't need specific
	// start and end dates.
	DateRange param.Field[TrafficAnomalyLocationGetParamsDateRange] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[TrafficAnomalyLocationGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit  param.Field[int64]                                 `query:"limit"`
	Status param.Field[TrafficAnomalyLocationGetParamsStatus] `query:"status"`
}

// URLQuery serializes [TrafficAnomalyLocationGetParams]'s query parameters as
// `url.Values`.
func (r TrafficAnomalyLocationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Shorthand date ranges for the last X days - use when you don't need specific
// start and end dates.
type TrafficAnomalyLocationGetParamsDateRange string

const (
	TrafficAnomalyLocationGetParamsDateRange1d         TrafficAnomalyLocationGetParamsDateRange = "1d"
	TrafficAnomalyLocationGetParamsDateRange2d         TrafficAnomalyLocationGetParamsDateRange = "2d"
	TrafficAnomalyLocationGetParamsDateRange7d         TrafficAnomalyLocationGetParamsDateRange = "7d"
	TrafficAnomalyLocationGetParamsDateRange14d        TrafficAnomalyLocationGetParamsDateRange = "14d"
	TrafficAnomalyLocationGetParamsDateRange28d        TrafficAnomalyLocationGetParamsDateRange = "28d"
	TrafficAnomalyLocationGetParamsDateRange12w        TrafficAnomalyLocationGetParamsDateRange = "12w"
	TrafficAnomalyLocationGetParamsDateRange24w        TrafficAnomalyLocationGetParamsDateRange = "24w"
	TrafficAnomalyLocationGetParamsDateRange52w        TrafficAnomalyLocationGetParamsDateRange = "52w"
	TrafficAnomalyLocationGetParamsDateRange1dControl  TrafficAnomalyLocationGetParamsDateRange = "1dControl"
	TrafficAnomalyLocationGetParamsDateRange2dControl  TrafficAnomalyLocationGetParamsDateRange = "2dControl"
	TrafficAnomalyLocationGetParamsDateRange7dControl  TrafficAnomalyLocationGetParamsDateRange = "7dControl"
	TrafficAnomalyLocationGetParamsDateRange14dControl TrafficAnomalyLocationGetParamsDateRange = "14dControl"
	TrafficAnomalyLocationGetParamsDateRange28dControl TrafficAnomalyLocationGetParamsDateRange = "28dControl"
	TrafficAnomalyLocationGetParamsDateRange12wControl TrafficAnomalyLocationGetParamsDateRange = "12wControl"
	TrafficAnomalyLocationGetParamsDateRange24wControl TrafficAnomalyLocationGetParamsDateRange = "24wControl"
)

// Format results are returned in.
type TrafficAnomalyLocationGetParamsFormat string

const (
	TrafficAnomalyLocationGetParamsFormatJson TrafficAnomalyLocationGetParamsFormat = "JSON"
	TrafficAnomalyLocationGetParamsFormatCsv  TrafficAnomalyLocationGetParamsFormat = "CSV"
)

type TrafficAnomalyLocationGetParamsStatus string

const (
	TrafficAnomalyLocationGetParamsStatusVerified   TrafficAnomalyLocationGetParamsStatus = "VERIFIED"
	TrafficAnomalyLocationGetParamsStatusUnverified TrafficAnomalyLocationGetParamsStatus = "UNVERIFIED"
)

type TrafficAnomalyLocationGetResponseEnvelope struct {
	Result  TrafficAnomalyLocationGetResponse             `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    trafficAnomalyLocationGetResponseEnvelopeJSON `json:"-"`
}

// trafficAnomalyLocationGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [TrafficAnomalyLocationGetResponseEnvelope]
type trafficAnomalyLocationGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TrafficAnomalyLocationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r trafficAnomalyLocationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
