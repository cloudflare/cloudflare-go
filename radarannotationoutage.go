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

// Get outages annotations
func (r *RadarAnnotationOutageService) List(ctx context.Context, query RadarAnnotationOutageListParams, opts ...option.RequestOption) (res *RadarAnnotationOutageListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/annotations/outages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAnnotationOutageListResponse struct {
	Result  RadarAnnotationOutageListResponseResult `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    radarAnnotationOutageListResponseJSON   `json:"-"`
}

// radarAnnotationOutageListResponseJSON contains the JSON metadata for the struct
// [RadarAnnotationOutageListResponse]
type radarAnnotationOutageListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAnnotationOutageListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAnnotationOutageListResponseResult struct {
	Annotations []RadarAnnotationOutageListResponseResultAnnotation `json:"annotations,required"`
	JSON        radarAnnotationOutageListResponseResultJSON         `json:"-"`
}

// radarAnnotationOutageListResponseResultJSON contains the JSON metadata for the
// struct [RadarAnnotationOutageListResponseResult]
type radarAnnotationOutageListResponseResultJSON struct {
	Annotations apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAnnotationOutageListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAnnotationOutageListResponseResultAnnotation struct {
	ASNs             []int64                                                             `json:"asns,required"`
	ASNsDetails      []RadarAnnotationOutageListResponseResultAnnotationsASNsDetail      `json:"asnsDetails,required"`
	DataSource       string                                                              `json:"dataSource,required"`
	EventType        string                                                              `json:"eventType,required"`
	Locations        []string                                                            `json:"locations,required"`
	LocationsDetails []RadarAnnotationOutageListResponseResultAnnotationsLocationsDetail `json:"locationsDetails,required"`
	Outage           RadarAnnotationOutageListResponseResultAnnotationsOutage            `json:"outage,required"`
	StartDate        string                                                              `json:"startDate,required"`
	Description      string                                                              `json:"description"`
	EndDate          string                                                              `json:"endDate"`
	LinkedURL        string                                                              `json:"linkedUrl"`
	Scope            string                                                              `json:"scope"`
	JSON             radarAnnotationOutageListResponseResultAnnotationJSON               `json:"-"`
}

// radarAnnotationOutageListResponseResultAnnotationJSON contains the JSON metadata
// for the struct [RadarAnnotationOutageListResponseResultAnnotation]
type radarAnnotationOutageListResponseResultAnnotationJSON struct {
	ASNs             apijson.Field
	ASNsDetails      apijson.Field
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

func (r *RadarAnnotationOutageListResponseResultAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAnnotationOutageListResponseResultAnnotationsASNsDetail struct {
	ASN       int64                                                                  `json:"asn,required"`
	Locations RadarAnnotationOutageListResponseResultAnnotationsASNsDetailsLocations `json:"locations,required"`
	Name      string                                                                 `json:"name,required"`
	JSON      radarAnnotationOutageListResponseResultAnnotationsASNsDetailJSON       `json:"-"`
}

// radarAnnotationOutageListResponseResultAnnotationsASNsDetailJSON contains the
// JSON metadata for the struct
// [RadarAnnotationOutageListResponseResultAnnotationsASNsDetail]
type radarAnnotationOutageListResponseResultAnnotationsASNsDetailJSON struct {
	ASN         apijson.Field
	Locations   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAnnotationOutageListResponseResultAnnotationsASNsDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAnnotationOutageListResponseResultAnnotationsASNsDetailsLocations struct {
	ASN  string                                                                     `json:"asn,required"`
	Name string                                                                     `json:"name,required"`
	JSON radarAnnotationOutageListResponseResultAnnotationsASNsDetailsLocationsJSON `json:"-"`
}

// radarAnnotationOutageListResponseResultAnnotationsASNsDetailsLocationsJSON
// contains the JSON metadata for the struct
// [RadarAnnotationOutageListResponseResultAnnotationsASNsDetailsLocations]
type radarAnnotationOutageListResponseResultAnnotationsASNsDetailsLocationsJSON struct {
	ASN         apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAnnotationOutageListResponseResultAnnotationsASNsDetailsLocations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAnnotationOutageListResponseResultAnnotationsLocationsDetail struct {
	Code string                                                                `json:"code,required"`
	Name string                                                                `json:"name,required"`
	JSON radarAnnotationOutageListResponseResultAnnotationsLocationsDetailJSON `json:"-"`
}

// radarAnnotationOutageListResponseResultAnnotationsLocationsDetailJSON contains
// the JSON metadata for the struct
// [RadarAnnotationOutageListResponseResultAnnotationsLocationsDetail]
type radarAnnotationOutageListResponseResultAnnotationsLocationsDetailJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAnnotationOutageListResponseResultAnnotationsLocationsDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAnnotationOutageListResponseResultAnnotationsOutage struct {
	OutageCause string                                                       `json:"outageCause,required"`
	OutageType  string                                                       `json:"outageType,required"`
	JSON        radarAnnotationOutageListResponseResultAnnotationsOutageJSON `json:"-"`
}

// radarAnnotationOutageListResponseResultAnnotationsOutageJSON contains the JSON
// metadata for the struct
// [RadarAnnotationOutageListResponseResultAnnotationsOutage]
type radarAnnotationOutageListResponseResultAnnotationsOutageJSON struct {
	OutageCause apijson.Field
	OutageType  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAnnotationOutageListResponseResultAnnotationsOutage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAnnotationOutageListParams struct {
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
	RadarAnnotationOutageListParamsDateRange7d         RadarAnnotationOutageListParamsDateRange = "7d"
	RadarAnnotationOutageListParamsDateRange14d        RadarAnnotationOutageListParamsDateRange = "14d"
	RadarAnnotationOutageListParamsDateRange28d        RadarAnnotationOutageListParamsDateRange = "28d"
	RadarAnnotationOutageListParamsDateRange12w        RadarAnnotationOutageListParamsDateRange = "12w"
	RadarAnnotationOutageListParamsDateRange24w        RadarAnnotationOutageListParamsDateRange = "24w"
	RadarAnnotationOutageListParamsDateRange52w        RadarAnnotationOutageListParamsDateRange = "52w"
	RadarAnnotationOutageListParamsDateRange1dControl  RadarAnnotationOutageListParamsDateRange = "1dControl"
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
