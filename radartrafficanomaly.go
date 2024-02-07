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
	path := "radar/traffic_anomalies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarTrafficAnomalyListResponse struct {
	Result  RadarTrafficAnomalyListResponseResult `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    radarTrafficAnomalyListResponseJSON   `json:"-"`
}

// radarTrafficAnomalyListResponseJSON contains the JSON metadata for the struct
// [RadarTrafficAnomalyListResponse]
type radarTrafficAnomalyListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTrafficAnomalyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarTrafficAnomalyListResponseResult struct {
	TrafficAnomalies []RadarTrafficAnomalyListResponseResultTrafficAnomaly `json:"trafficAnomalies,required"`
	JSON             radarTrafficAnomalyListResponseResultJSON             `json:"-"`
}

// radarTrafficAnomalyListResponseResultJSON contains the JSON metadata for the
// struct [RadarTrafficAnomalyListResponseResult]
type radarTrafficAnomalyListResponseResultJSON struct {
	TrafficAnomalies apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RadarTrafficAnomalyListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarTrafficAnomalyListResponseResultTrafficAnomaly struct {
	StartDate            string                                                               `json:"startDate,required"`
	Status               string                                                               `json:"status,required"`
	Type                 string                                                               `json:"type,required"`
	Uuid                 string                                                               `json:"uuid,required"`
	AsnDetails           RadarTrafficAnomalyListResponseResultTrafficAnomaliesAsnDetails      `json:"asnDetails"`
	EndDate              string                                                               `json:"endDate"`
	LocationDetails      RadarTrafficAnomalyListResponseResultTrafficAnomaliesLocationDetails `json:"locationDetails"`
	VisibleInDataSources []string                                                             `json:"visibleInDataSources"`
	JSON                 radarTrafficAnomalyListResponseResultTrafficAnomalyJSON              `json:"-"`
}

// radarTrafficAnomalyListResponseResultTrafficAnomalyJSON contains the JSON
// metadata for the struct [RadarTrafficAnomalyListResponseResultTrafficAnomaly]
type radarTrafficAnomalyListResponseResultTrafficAnomalyJSON struct {
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

func (r *RadarTrafficAnomalyListResponseResultTrafficAnomaly) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarTrafficAnomalyListResponseResultTrafficAnomaliesAsnDetails struct {
	Asn       string                                                                   `json:"asn,required"`
	Name      string                                                                   `json:"name,required"`
	Locations RadarTrafficAnomalyListResponseResultTrafficAnomaliesAsnDetailsLocations `json:"locations"`
	JSON      radarTrafficAnomalyListResponseResultTrafficAnomaliesAsnDetailsJSON      `json:"-"`
}

// radarTrafficAnomalyListResponseResultTrafficAnomaliesAsnDetailsJSON contains the
// JSON metadata for the struct
// [RadarTrafficAnomalyListResponseResultTrafficAnomaliesAsnDetails]
type radarTrafficAnomalyListResponseResultTrafficAnomaliesAsnDetailsJSON struct {
	Asn         apijson.Field
	Name        apijson.Field
	Locations   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTrafficAnomalyListResponseResultTrafficAnomaliesAsnDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarTrafficAnomalyListResponseResultTrafficAnomaliesAsnDetailsLocations struct {
	Code string                                                                       `json:"code,required"`
	Name string                                                                       `json:"name,required"`
	JSON radarTrafficAnomalyListResponseResultTrafficAnomaliesAsnDetailsLocationsJSON `json:"-"`
}

// radarTrafficAnomalyListResponseResultTrafficAnomaliesAsnDetailsLocationsJSON
// contains the JSON metadata for the struct
// [RadarTrafficAnomalyListResponseResultTrafficAnomaliesAsnDetailsLocations]
type radarTrafficAnomalyListResponseResultTrafficAnomaliesAsnDetailsLocationsJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTrafficAnomalyListResponseResultTrafficAnomaliesAsnDetailsLocations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarTrafficAnomalyListResponseResultTrafficAnomaliesLocationDetails struct {
	Code string                                                                   `json:"code,required"`
	Name string                                                                   `json:"name,required"`
	JSON radarTrafficAnomalyListResponseResultTrafficAnomaliesLocationDetailsJSON `json:"-"`
}

// radarTrafficAnomalyListResponseResultTrafficAnomaliesLocationDetailsJSON
// contains the JSON metadata for the struct
// [RadarTrafficAnomalyListResponseResultTrafficAnomaliesLocationDetails]
type radarTrafficAnomalyListResponseResultTrafficAnomaliesLocationDetailsJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTrafficAnomalyListResponseResultTrafficAnomaliesLocationDetails) UnmarshalJSON(data []byte) (err error) {
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
