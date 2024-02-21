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

// RadarTrafficAnomalyService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarTrafficAnomalyService]
// method instead.
type RadarTrafficAnomalyService struct {
	Options   []option.RequestOption
	Locations *RadarTrafficAnomalyLocationService
}

// NewRadarTrafficAnomalyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarTrafficAnomalyService(opts ...option.RequestOption) (r *RadarTrafficAnomalyService) {
	r = &RadarTrafficAnomalyService{}
	r.Options = opts
	r.Locations = NewRadarTrafficAnomalyLocationService(opts...)
	return
}

// Internet traffic anomalies are signals that might point to an outage, These
// alerts are automatically detected by Radar and then manually verified by our
// team. This endpoint returns the latest alerts.
func (r *RadarTrafficAnomalyService) Get(ctx context.Context, query RadarTrafficAnomalyGetParams, opts ...option.RequestOption) (res *RadarTrafficAnomalyGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarTrafficAnomalyGetResponseEnvelope
	path := "radar/traffic_anomalies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarTrafficAnomalyGetResponse struct {
	TrafficAnomalies []RadarTrafficAnomalyGetResponseTrafficAnomaly `json:"trafficAnomalies,required"`
	JSON             radarTrafficAnomalyGetResponseJSON             `json:"-"`
}

// radarTrafficAnomalyGetResponseJSON contains the JSON metadata for the struct
// [RadarTrafficAnomalyGetResponse]
type radarTrafficAnomalyGetResponseJSON struct {
	TrafficAnomalies apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RadarTrafficAnomalyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarTrafficAnomalyGetResponseTrafficAnomaly struct {
	StartDate            string                                                        `json:"startDate,required"`
	Status               string                                                        `json:"status,required"`
	Type                 string                                                        `json:"type,required"`
	Uuid                 string                                                        `json:"uuid,required"`
	AsnDetails           RadarTrafficAnomalyGetResponseTrafficAnomaliesAsnDetails      `json:"asnDetails"`
	EndDate              string                                                        `json:"endDate"`
	LocationDetails      RadarTrafficAnomalyGetResponseTrafficAnomaliesLocationDetails `json:"locationDetails"`
	VisibleInDataSources []string                                                      `json:"visibleInDataSources"`
	JSON                 radarTrafficAnomalyGetResponseTrafficAnomalyJSON              `json:"-"`
}

// radarTrafficAnomalyGetResponseTrafficAnomalyJSON contains the JSON metadata for
// the struct [RadarTrafficAnomalyGetResponseTrafficAnomaly]
type radarTrafficAnomalyGetResponseTrafficAnomalyJSON struct {
	StartDate            apijson.Field
	Status               apijson.Field
	Type                 apijson.Field
	Uuid                 apijson.Field
	AsnDetails           apijson.Field
	EndDate              apijson.Field
	LocationDetails      apijson.Field
	VisibleInDataSources apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *RadarTrafficAnomalyGetResponseTrafficAnomaly) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarTrafficAnomalyGetResponseTrafficAnomaliesAsnDetails struct {
	Asn       string                                                            `json:"asn,required"`
	Name      string                                                            `json:"name,required"`
	Locations RadarTrafficAnomalyGetResponseTrafficAnomaliesAsnDetailsLocations `json:"locations"`
	JSON      radarTrafficAnomalyGetResponseTrafficAnomaliesAsnDetailsJSON      `json:"-"`
}

// radarTrafficAnomalyGetResponseTrafficAnomaliesAsnDetailsJSON contains the JSON
// metadata for the struct
// [RadarTrafficAnomalyGetResponseTrafficAnomaliesAsnDetails]
type radarTrafficAnomalyGetResponseTrafficAnomaliesAsnDetailsJSON struct {
	Asn         apijson.Field
	Name        apijson.Field
	Locations   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTrafficAnomalyGetResponseTrafficAnomaliesAsnDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarTrafficAnomalyGetResponseTrafficAnomaliesAsnDetailsLocations struct {
	Code string                                                                `json:"code,required"`
	Name string                                                                `json:"name,required"`
	JSON radarTrafficAnomalyGetResponseTrafficAnomaliesAsnDetailsLocationsJSON `json:"-"`
}

// radarTrafficAnomalyGetResponseTrafficAnomaliesAsnDetailsLocationsJSON contains
// the JSON metadata for the struct
// [RadarTrafficAnomalyGetResponseTrafficAnomaliesAsnDetailsLocations]
type radarTrafficAnomalyGetResponseTrafficAnomaliesAsnDetailsLocationsJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTrafficAnomalyGetResponseTrafficAnomaliesAsnDetailsLocations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarTrafficAnomalyGetResponseTrafficAnomaliesLocationDetails struct {
	Code string                                                            `json:"code,required"`
	Name string                                                            `json:"name,required"`
	JSON radarTrafficAnomalyGetResponseTrafficAnomaliesLocationDetailsJSON `json:"-"`
}

// radarTrafficAnomalyGetResponseTrafficAnomaliesLocationDetailsJSON contains the
// JSON metadata for the struct
// [RadarTrafficAnomalyGetResponseTrafficAnomaliesLocationDetails]
type radarTrafficAnomalyGetResponseTrafficAnomaliesLocationDetailsJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTrafficAnomalyGetResponseTrafficAnomaliesLocationDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarTrafficAnomalyGetParams struct {
	// Single ASN as integer.
	Asn param.Field[int64] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Shorthand date ranges for the last X days - use when you don't need specific
	// start and end dates.
	DateRange param.Field[RadarTrafficAnomalyGetParamsDateRange] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarTrafficAnomalyGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Location Alpha2 code.
	Location param.Field[string] `query:"location"`
	// Number of objects to skip before grabbing results.
	Offset param.Field[int64]                              `query:"offset"`
	Status param.Field[RadarTrafficAnomalyGetParamsStatus] `query:"status"`
}

// URLQuery serializes [RadarTrafficAnomalyGetParams]'s query parameters as
// `url.Values`.
func (r RadarTrafficAnomalyGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Shorthand date ranges for the last X days - use when you don't need specific
// start and end dates.
type RadarTrafficAnomalyGetParamsDateRange string

const (
	RadarTrafficAnomalyGetParamsDateRange1d         RadarTrafficAnomalyGetParamsDateRange = "1d"
	RadarTrafficAnomalyGetParamsDateRange2d         RadarTrafficAnomalyGetParamsDateRange = "2d"
	RadarTrafficAnomalyGetParamsDateRange7d         RadarTrafficAnomalyGetParamsDateRange = "7d"
	RadarTrafficAnomalyGetParamsDateRange14d        RadarTrafficAnomalyGetParamsDateRange = "14d"
	RadarTrafficAnomalyGetParamsDateRange28d        RadarTrafficAnomalyGetParamsDateRange = "28d"
	RadarTrafficAnomalyGetParamsDateRange12w        RadarTrafficAnomalyGetParamsDateRange = "12w"
	RadarTrafficAnomalyGetParamsDateRange24w        RadarTrafficAnomalyGetParamsDateRange = "24w"
	RadarTrafficAnomalyGetParamsDateRange52w        RadarTrafficAnomalyGetParamsDateRange = "52w"
	RadarTrafficAnomalyGetParamsDateRange1dControl  RadarTrafficAnomalyGetParamsDateRange = "1dControl"
	RadarTrafficAnomalyGetParamsDateRange2dControl  RadarTrafficAnomalyGetParamsDateRange = "2dControl"
	RadarTrafficAnomalyGetParamsDateRange7dControl  RadarTrafficAnomalyGetParamsDateRange = "7dControl"
	RadarTrafficAnomalyGetParamsDateRange14dControl RadarTrafficAnomalyGetParamsDateRange = "14dControl"
	RadarTrafficAnomalyGetParamsDateRange28dControl RadarTrafficAnomalyGetParamsDateRange = "28dControl"
	RadarTrafficAnomalyGetParamsDateRange12wControl RadarTrafficAnomalyGetParamsDateRange = "12wControl"
	RadarTrafficAnomalyGetParamsDateRange24wControl RadarTrafficAnomalyGetParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarTrafficAnomalyGetParamsFormat string

const (
	RadarTrafficAnomalyGetParamsFormatJson RadarTrafficAnomalyGetParamsFormat = "JSON"
	RadarTrafficAnomalyGetParamsFormatCsv  RadarTrafficAnomalyGetParamsFormat = "CSV"
)

type RadarTrafficAnomalyGetParamsStatus string

const (
	RadarTrafficAnomalyGetParamsStatusVerified   RadarTrafficAnomalyGetParamsStatus = "VERIFIED"
	RadarTrafficAnomalyGetParamsStatusUnverified RadarTrafficAnomalyGetParamsStatus = "UNVERIFIED"
)

type RadarTrafficAnomalyGetResponseEnvelope struct {
	Result  RadarTrafficAnomalyGetResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarTrafficAnomalyGetResponseEnvelopeJSON `json:"-"`
}

// radarTrafficAnomalyGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarTrafficAnomalyGetResponseEnvelope]
type radarTrafficAnomalyGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTrafficAnomalyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
