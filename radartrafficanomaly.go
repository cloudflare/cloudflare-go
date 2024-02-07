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
func (r *RadarTrafficAnomalyService) List(ctx context.Context, query RadarTrafficAnomalyListParams, opts ...option.RequestOption) (res *RadarTrafficAnomalyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarTrafficAnomalyListResponseEnvelope
	path := "radar/traffic_anomalies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarTrafficAnomalyListResponse struct {
	TrafficAnomalies []RadarTrafficAnomalyListResponseTrafficAnomaly `json:"trafficAnomalies,required"`
	JSON             radarTrafficAnomalyListResponseJSON             `json:"-"`
}

// radarTrafficAnomalyListResponseJSON contains the JSON metadata for the struct
// [RadarTrafficAnomalyListResponse]
type radarTrafficAnomalyListResponseJSON struct {
	TrafficAnomalies apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RadarTrafficAnomalyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarTrafficAnomalyListResponseTrafficAnomaly struct {
	StartDate            string                                                         `json:"startDate,required"`
	Status               string                                                         `json:"status,required"`
	Type                 string                                                         `json:"type,required"`
	Uuid                 string                                                         `json:"uuid,required"`
	AsnDetails           RadarTrafficAnomalyListResponseTrafficAnomaliesAsnDetails      `json:"asnDetails"`
	EndDate              string                                                         `json:"endDate"`
	LocationDetails      RadarTrafficAnomalyListResponseTrafficAnomaliesLocationDetails `json:"locationDetails"`
	VisibleInDataSources []string                                                       `json:"visibleInDataSources"`
	JSON                 radarTrafficAnomalyListResponseTrafficAnomalyJSON              `json:"-"`
}

// radarTrafficAnomalyListResponseTrafficAnomalyJSON contains the JSON metadata for
// the struct [RadarTrafficAnomalyListResponseTrafficAnomaly]
type radarTrafficAnomalyListResponseTrafficAnomalyJSON struct {
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

func (r *RadarTrafficAnomalyListResponseTrafficAnomaly) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarTrafficAnomalyListResponseTrafficAnomaliesAsnDetails struct {
	Asn       string                                                             `json:"asn,required"`
	Name      string                                                             `json:"name,required"`
	Locations RadarTrafficAnomalyListResponseTrafficAnomaliesAsnDetailsLocations `json:"locations"`
	JSON      radarTrafficAnomalyListResponseTrafficAnomaliesAsnDetailsJSON      `json:"-"`
}

// radarTrafficAnomalyListResponseTrafficAnomaliesAsnDetailsJSON contains the JSON
// metadata for the struct
// [RadarTrafficAnomalyListResponseTrafficAnomaliesAsnDetails]
type radarTrafficAnomalyListResponseTrafficAnomaliesAsnDetailsJSON struct {
	Asn         apijson.Field
	Name        apijson.Field
	Locations   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTrafficAnomalyListResponseTrafficAnomaliesAsnDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarTrafficAnomalyListResponseTrafficAnomaliesAsnDetailsLocations struct {
	Code string                                                                 `json:"code,required"`
	Name string                                                                 `json:"name,required"`
	JSON radarTrafficAnomalyListResponseTrafficAnomaliesAsnDetailsLocationsJSON `json:"-"`
}

// radarTrafficAnomalyListResponseTrafficAnomaliesAsnDetailsLocationsJSON contains
// the JSON metadata for the struct
// [RadarTrafficAnomalyListResponseTrafficAnomaliesAsnDetailsLocations]
type radarTrafficAnomalyListResponseTrafficAnomaliesAsnDetailsLocationsJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTrafficAnomalyListResponseTrafficAnomaliesAsnDetailsLocations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarTrafficAnomalyListResponseTrafficAnomaliesLocationDetails struct {
	Code string                                                             `json:"code,required"`
	Name string                                                             `json:"name,required"`
	JSON radarTrafficAnomalyListResponseTrafficAnomaliesLocationDetailsJSON `json:"-"`
}

// radarTrafficAnomalyListResponseTrafficAnomaliesLocationDetailsJSON contains the
// JSON metadata for the struct
// [RadarTrafficAnomalyListResponseTrafficAnomaliesLocationDetails]
type radarTrafficAnomalyListResponseTrafficAnomaliesLocationDetailsJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTrafficAnomalyListResponseTrafficAnomaliesLocationDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarTrafficAnomalyListParams struct {
	// Single ASN as integer.
	Asn param.Field[int64] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Shorthand date ranges for the last X days - use when you don't need specific
	// start and end dates.
	DateRange param.Field[RadarTrafficAnomalyListParamsDateRange] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarTrafficAnomalyListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Location Alpha2 code.
	Location param.Field[string] `query:"location"`
	// Number of objects to skip before grabbing results.
	Offset param.Field[int64]                               `query:"offset"`
	Status param.Field[RadarTrafficAnomalyListParamsStatus] `query:"status"`
}

// URLQuery serializes [RadarTrafficAnomalyListParams]'s query parameters as
// `url.Values`.
func (r RadarTrafficAnomalyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Shorthand date ranges for the last X days - use when you don't need specific
// start and end dates.
type RadarTrafficAnomalyListParamsDateRange string

const (
	RadarTrafficAnomalyListParamsDateRange1d         RadarTrafficAnomalyListParamsDateRange = "1d"
	RadarTrafficAnomalyListParamsDateRange2d         RadarTrafficAnomalyListParamsDateRange = "2d"
	RadarTrafficAnomalyListParamsDateRange7d         RadarTrafficAnomalyListParamsDateRange = "7d"
	RadarTrafficAnomalyListParamsDateRange14d        RadarTrafficAnomalyListParamsDateRange = "14d"
	RadarTrafficAnomalyListParamsDateRange28d        RadarTrafficAnomalyListParamsDateRange = "28d"
	RadarTrafficAnomalyListParamsDateRange12w        RadarTrafficAnomalyListParamsDateRange = "12w"
	RadarTrafficAnomalyListParamsDateRange24w        RadarTrafficAnomalyListParamsDateRange = "24w"
	RadarTrafficAnomalyListParamsDateRange52w        RadarTrafficAnomalyListParamsDateRange = "52w"
	RadarTrafficAnomalyListParamsDateRange1dControl  RadarTrafficAnomalyListParamsDateRange = "1dControl"
	RadarTrafficAnomalyListParamsDateRange2dControl  RadarTrafficAnomalyListParamsDateRange = "2dControl"
	RadarTrafficAnomalyListParamsDateRange7dControl  RadarTrafficAnomalyListParamsDateRange = "7dControl"
	RadarTrafficAnomalyListParamsDateRange14dControl RadarTrafficAnomalyListParamsDateRange = "14dControl"
	RadarTrafficAnomalyListParamsDateRange28dControl RadarTrafficAnomalyListParamsDateRange = "28dControl"
	RadarTrafficAnomalyListParamsDateRange12wControl RadarTrafficAnomalyListParamsDateRange = "12wControl"
	RadarTrafficAnomalyListParamsDateRange24wControl RadarTrafficAnomalyListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarTrafficAnomalyListParamsFormat string

const (
	RadarTrafficAnomalyListParamsFormatJson RadarTrafficAnomalyListParamsFormat = "JSON"
	RadarTrafficAnomalyListParamsFormatCsv  RadarTrafficAnomalyListParamsFormat = "CSV"
)

type RadarTrafficAnomalyListParamsStatus string

const (
	RadarTrafficAnomalyListParamsStatusVerified   RadarTrafficAnomalyListParamsStatus = "VERIFIED"
	RadarTrafficAnomalyListParamsStatusUnverified RadarTrafficAnomalyListParamsStatus = "UNVERIFIED"
)

type RadarTrafficAnomalyListResponseEnvelope struct {
	Result  RadarTrafficAnomalyListResponse             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    radarTrafficAnomalyListResponseEnvelopeJSON `json:"-"`
}

// radarTrafficAnomalyListResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarTrafficAnomalyListResponseEnvelope]
type radarTrafficAnomalyListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTrafficAnomalyListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
