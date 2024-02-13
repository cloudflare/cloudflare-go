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
	Options   []option.RequestOption
	Locations *RadarAnnotationOutageLocationService
}

// NewRadarAnnotationOutageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAnnotationOutageService(opts ...option.RequestOption) (r *RadarAnnotationOutageService) {
	r = &RadarAnnotationOutageService{}
	r.Options = opts
	r.Locations = NewRadarAnnotationOutageLocationService(opts...)
	return
}

// Get latest Internet outages and anomalies.
func (r *RadarAnnotationOutageService) List(ctx context.Context, query RadarAnnotationOutageListParams, opts ...option.RequestOption) (res *RadarAnnotationOutageListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAnnotationOutageListResponseEnvelope
	path := "radar/annotations/outages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAnnotationOutageListResponse struct {
	Annotations []RadarAnnotationOutageListResponseAnnotation `json:"annotations,required"`
	JSON        radarAnnotationOutageListResponseJSON         `json:"-"`
}

// radarAnnotationOutageListResponseJSON contains the JSON metadata for the struct
// [RadarAnnotationOutageListResponse]
type radarAnnotationOutageListResponseJSON struct {
	Annotations apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAnnotationOutageListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAnnotationOutageListResponseAnnotation struct {
	ID               string                                                        `json:"id,required"`
	Asns             []int64                                                       `json:"asns,required"`
	AsnsDetails      []RadarAnnotationOutageListResponseAnnotationsAsnsDetail      `json:"asnsDetails,required"`
	DataSource       string                                                        `json:"dataSource,required"`
	EventType        string                                                        `json:"eventType,required"`
	Locations        []string                                                      `json:"locations,required"`
	LocationsDetails []RadarAnnotationOutageListResponseAnnotationsLocationsDetail `json:"locationsDetails,required"`
	Outage           RadarAnnotationOutageListResponseAnnotationsOutage            `json:"outage,required"`
	StartDate        string                                                        `json:"startDate,required"`
	Description      string                                                        `json:"description"`
	EndDate          string                                                        `json:"endDate"`
	LinkedURL        string                                                        `json:"linkedUrl"`
	Scope            string                                                        `json:"scope"`
	JSON             radarAnnotationOutageListResponseAnnotationJSON               `json:"-"`
}

// radarAnnotationOutageListResponseAnnotationJSON contains the JSON metadata for
// the struct [RadarAnnotationOutageListResponseAnnotation]
type radarAnnotationOutageListResponseAnnotationJSON struct {
	ID               apijson.Field
	Asns             apijson.Field
	AsnsDetails      apijson.Field
	DataSource       apijson.Field
	EventType        apijson.Field
	Locations        apijson.Field
	LocationsDetails apijson.Field
	Outage           apijson.Field
	StartDate        apijson.Field
	Description      apijson.Field
	EndDate          apijson.Field
	LinkedURL        apijson.Field
	Scope            apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RadarAnnotationOutageListResponseAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAnnotationOutageListResponseAnnotationsAsnsDetail struct {
	Asn       string                                                           `json:"asn,required"`
	Name      string                                                           `json:"name,required"`
	Locations RadarAnnotationOutageListResponseAnnotationsAsnsDetailsLocations `json:"locations"`
	JSON      radarAnnotationOutageListResponseAnnotationsAsnsDetailJSON       `json:"-"`
}

// radarAnnotationOutageListResponseAnnotationsAsnsDetailJSON contains the JSON
// metadata for the struct [RadarAnnotationOutageListResponseAnnotationsAsnsDetail]
type radarAnnotationOutageListResponseAnnotationsAsnsDetailJSON struct {
	Asn         apijson.Field
	Name        apijson.Field
	Locations   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAnnotationOutageListResponseAnnotationsAsnsDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAnnotationOutageListResponseAnnotationsAsnsDetailsLocations struct {
	Code string                                                               `json:"code,required"`
	Name string                                                               `json:"name,required"`
	JSON radarAnnotationOutageListResponseAnnotationsAsnsDetailsLocationsJSON `json:"-"`
}

// radarAnnotationOutageListResponseAnnotationsAsnsDetailsLocationsJSON contains
// the JSON metadata for the struct
// [RadarAnnotationOutageListResponseAnnotationsAsnsDetailsLocations]
type radarAnnotationOutageListResponseAnnotationsAsnsDetailsLocationsJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAnnotationOutageListResponseAnnotationsAsnsDetailsLocations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAnnotationOutageListResponseAnnotationsLocationsDetail struct {
	Code string                                                          `json:"code,required"`
	Name string                                                          `json:"name,required"`
	JSON radarAnnotationOutageListResponseAnnotationsLocationsDetailJSON `json:"-"`
}

// radarAnnotationOutageListResponseAnnotationsLocationsDetailJSON contains the
// JSON metadata for the struct
// [RadarAnnotationOutageListResponseAnnotationsLocationsDetail]
type radarAnnotationOutageListResponseAnnotationsLocationsDetailJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAnnotationOutageListResponseAnnotationsLocationsDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAnnotationOutageListResponseAnnotationsOutage struct {
	OutageCause string                                                 `json:"outageCause,required"`
	OutageType  string                                                 `json:"outageType,required"`
	JSON        radarAnnotationOutageListResponseAnnotationsOutageJSON `json:"-"`
}

// radarAnnotationOutageListResponseAnnotationsOutageJSON contains the JSON
// metadata for the struct [RadarAnnotationOutageListResponseAnnotationsOutage]
type radarAnnotationOutageListResponseAnnotationsOutageJSON struct {
	OutageCause apijson.Field
	OutageType  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAnnotationOutageListResponseAnnotationsOutage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAnnotationOutageListParams struct {
	// Single ASN as integer.
	Asn param.Field[int64] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Shorthand date ranges for the last X days - use when you don't need specific
	// start and end dates.
	DateRange param.Field[RadarAnnotationOutageListParamsDateRange] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAnnotationOutageListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Location Alpha2 code.
	Location param.Field[string] `query:"location"`
	// Number of objects to skip before grabbing results.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [RadarAnnotationOutageListParams]'s query parameters as
// `url.Values`.
func (r RadarAnnotationOutageListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Shorthand date ranges for the last X days - use when you don't need specific
// start and end dates.
type RadarAnnotationOutageListParamsDateRange string

const (
	RadarAnnotationOutageListParamsDateRange1d         RadarAnnotationOutageListParamsDateRange = "1d"
	RadarAnnotationOutageListParamsDateRange2d         RadarAnnotationOutageListParamsDateRange = "2d"
	RadarAnnotationOutageListParamsDateRange7d         RadarAnnotationOutageListParamsDateRange = "7d"
	RadarAnnotationOutageListParamsDateRange14d        RadarAnnotationOutageListParamsDateRange = "14d"
	RadarAnnotationOutageListParamsDateRange28d        RadarAnnotationOutageListParamsDateRange = "28d"
	RadarAnnotationOutageListParamsDateRange12w        RadarAnnotationOutageListParamsDateRange = "12w"
	RadarAnnotationOutageListParamsDateRange24w        RadarAnnotationOutageListParamsDateRange = "24w"
	RadarAnnotationOutageListParamsDateRange52w        RadarAnnotationOutageListParamsDateRange = "52w"
	RadarAnnotationOutageListParamsDateRange1dControl  RadarAnnotationOutageListParamsDateRange = "1dControl"
	RadarAnnotationOutageListParamsDateRange2dControl  RadarAnnotationOutageListParamsDateRange = "2dControl"
	RadarAnnotationOutageListParamsDateRange7dControl  RadarAnnotationOutageListParamsDateRange = "7dControl"
	RadarAnnotationOutageListParamsDateRange14dControl RadarAnnotationOutageListParamsDateRange = "14dControl"
	RadarAnnotationOutageListParamsDateRange28dControl RadarAnnotationOutageListParamsDateRange = "28dControl"
	RadarAnnotationOutageListParamsDateRange12wControl RadarAnnotationOutageListParamsDateRange = "12wControl"
	RadarAnnotationOutageListParamsDateRange24wControl RadarAnnotationOutageListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAnnotationOutageListParamsFormat string

const (
	RadarAnnotationOutageListParamsFormatJson RadarAnnotationOutageListParamsFormat = "JSON"
	RadarAnnotationOutageListParamsFormatCsv  RadarAnnotationOutageListParamsFormat = "CSV"
)

type RadarAnnotationOutageListResponseEnvelope struct {
	Result  RadarAnnotationOutageListResponse             `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    radarAnnotationOutageListResponseEnvelopeJSON `json:"-"`
}

// radarAnnotationOutageListResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarAnnotationOutageListResponseEnvelope]
type radarAnnotationOutageListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAnnotationOutageListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
