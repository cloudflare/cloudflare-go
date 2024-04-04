// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// AnnotationOutageService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAnnotationOutageService] method
// instead.
type AnnotationOutageService struct {
	Options []option.RequestOption
}

// NewAnnotationOutageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAnnotationOutageService(opts ...option.RequestOption) (r *AnnotationOutageService) {
	r = &AnnotationOutageService{}
	r.Options = opts
	return
}

// Get latest Internet outages and anomalies.
func (r *AnnotationOutageService) Get(ctx context.Context, query AnnotationOutageGetParams, opts ...option.RequestOption) (res *AnnotationOutageGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AnnotationOutageGetResponseEnvelope
	path := "radar/annotations/outages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the number of outages for locations.
func (r *AnnotationOutageService) Locations(ctx context.Context, query AnnotationOutageLocationsParams, opts ...option.RequestOption) (res *AnnotationOutageLocationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AnnotationOutageLocationsResponseEnvelope
	path := "radar/annotations/outages/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AnnotationOutageGetResponse struct {
	Annotations []AnnotationOutageGetResponseAnnotation `json:"annotations,required"`
	JSON        annotationOutageGetResponseJSON         `json:"-"`
}

// annotationOutageGetResponseJSON contains the JSON metadata for the struct
// [AnnotationOutageGetResponse]
type annotationOutageGetResponseJSON struct {
	Annotations apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationOutageGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationOutageGetResponseJSON) RawJSON() string {
	return r.raw
}

type AnnotationOutageGetResponseAnnotation struct {
	ID               string                                             `json:"id,required"`
	ASNs             []int64                                            `json:"asns,required"`
	ASNsDetails      []AnnotationOutageGetResponseAnnotationsASNsDetail `json:"asnsDetails,required"`
	DataSource       string                                             `json:"dataSource,required"`
	EventType        string                                             `json:"eventType,required"`
	Locations        []string                                           `json:"locations,required"`
	LocationsDetails []UnnamedSchemaRef16e559c45a31db5480e21fbe904b2e42 `json:"locationsDetails,required"`
	Outage           AnnotationOutageGetResponseAnnotationsOutage       `json:"outage,required"`
	StartDate        string                                             `json:"startDate,required"`
	Description      string                                             `json:"description"`
	EndDate          string                                             `json:"endDate"`
	LinkedURL        string                                             `json:"linkedUrl"`
	Scope            string                                             `json:"scope"`
	JSON             annotationOutageGetResponseAnnotationJSON          `json:"-"`
}

// annotationOutageGetResponseAnnotationJSON contains the JSON metadata for the
// struct [AnnotationOutageGetResponseAnnotation]
type annotationOutageGetResponseAnnotationJSON struct {
	ID               apijson.Field
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

func (r *AnnotationOutageGetResponseAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationOutageGetResponseAnnotationJSON) RawJSON() string {
	return r.raw
}

type AnnotationOutageGetResponseAnnotationsASNsDetail struct {
	ASN       string                                               `json:"asn,required"`
	Name      string                                               `json:"name,required"`
	Locations UnnamedSchemaRef16e559c45a31db5480e21fbe904b2e42     `json:"locations"`
	JSON      annotationOutageGetResponseAnnotationsASNsDetailJSON `json:"-"`
}

// annotationOutageGetResponseAnnotationsASNsDetailJSON contains the JSON metadata
// for the struct [AnnotationOutageGetResponseAnnotationsASNsDetail]
type annotationOutageGetResponseAnnotationsASNsDetailJSON struct {
	ASN         apijson.Field
	Name        apijson.Field
	Locations   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationOutageGetResponseAnnotationsASNsDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationOutageGetResponseAnnotationsASNsDetailJSON) RawJSON() string {
	return r.raw
}

type AnnotationOutageGetResponseAnnotationsOutage struct {
	OutageCause string                                           `json:"outageCause,required"`
	OutageType  string                                           `json:"outageType,required"`
	JSON        annotationOutageGetResponseAnnotationsOutageJSON `json:"-"`
}

// annotationOutageGetResponseAnnotationsOutageJSON contains the JSON metadata for
// the struct [AnnotationOutageGetResponseAnnotationsOutage]
type annotationOutageGetResponseAnnotationsOutageJSON struct {
	OutageCause apijson.Field
	OutageType  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationOutageGetResponseAnnotationsOutage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationOutageGetResponseAnnotationsOutageJSON) RawJSON() string {
	return r.raw
}

type AnnotationOutageLocationsResponse struct {
	Annotations []UnnamedSchemaRef83a14d589e799bc901b9ccc870251d09 `json:"annotations,required"`
	JSON        annotationOutageLocationsResponseJSON              `json:"-"`
}

// annotationOutageLocationsResponseJSON contains the JSON metadata for the struct
// [AnnotationOutageLocationsResponse]
type annotationOutageLocationsResponseJSON struct {
	Annotations apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationOutageLocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationOutageLocationsResponseJSON) RawJSON() string {
	return r.raw
}

type AnnotationOutageGetParams struct {
	// Single ASN as integer.
	ASN param.Field[int64] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Shorthand date ranges for the last X days - use when you don't need specific
	// start and end dates.
	DateRange param.Field[AnnotationOutageGetParamsDateRange] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AnnotationOutageGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Location Alpha2 code.
	Location param.Field[string] `query:"location"`
	// Number of objects to skip before grabbing results.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [AnnotationOutageGetParams]'s query parameters as
// `url.Values`.
func (r AnnotationOutageGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Shorthand date ranges for the last X days - use when you don't need specific
// start and end dates.
type AnnotationOutageGetParamsDateRange string

const (
	AnnotationOutageGetParamsDateRange1d         AnnotationOutageGetParamsDateRange = "1d"
	AnnotationOutageGetParamsDateRange2d         AnnotationOutageGetParamsDateRange = "2d"
	AnnotationOutageGetParamsDateRange7d         AnnotationOutageGetParamsDateRange = "7d"
	AnnotationOutageGetParamsDateRange14d        AnnotationOutageGetParamsDateRange = "14d"
	AnnotationOutageGetParamsDateRange28d        AnnotationOutageGetParamsDateRange = "28d"
	AnnotationOutageGetParamsDateRange12w        AnnotationOutageGetParamsDateRange = "12w"
	AnnotationOutageGetParamsDateRange24w        AnnotationOutageGetParamsDateRange = "24w"
	AnnotationOutageGetParamsDateRange52w        AnnotationOutageGetParamsDateRange = "52w"
	AnnotationOutageGetParamsDateRange1dControl  AnnotationOutageGetParamsDateRange = "1dControl"
	AnnotationOutageGetParamsDateRange2dControl  AnnotationOutageGetParamsDateRange = "2dControl"
	AnnotationOutageGetParamsDateRange7dControl  AnnotationOutageGetParamsDateRange = "7dControl"
	AnnotationOutageGetParamsDateRange14dControl AnnotationOutageGetParamsDateRange = "14dControl"
	AnnotationOutageGetParamsDateRange28dControl AnnotationOutageGetParamsDateRange = "28dControl"
	AnnotationOutageGetParamsDateRange12wControl AnnotationOutageGetParamsDateRange = "12wControl"
	AnnotationOutageGetParamsDateRange24wControl AnnotationOutageGetParamsDateRange = "24wControl"
)

func (r AnnotationOutageGetParamsDateRange) IsKnown() bool {
	switch r {
	case AnnotationOutageGetParamsDateRange1d, AnnotationOutageGetParamsDateRange2d, AnnotationOutageGetParamsDateRange7d, AnnotationOutageGetParamsDateRange14d, AnnotationOutageGetParamsDateRange28d, AnnotationOutageGetParamsDateRange12w, AnnotationOutageGetParamsDateRange24w, AnnotationOutageGetParamsDateRange52w, AnnotationOutageGetParamsDateRange1dControl, AnnotationOutageGetParamsDateRange2dControl, AnnotationOutageGetParamsDateRange7dControl, AnnotationOutageGetParamsDateRange14dControl, AnnotationOutageGetParamsDateRange28dControl, AnnotationOutageGetParamsDateRange12wControl, AnnotationOutageGetParamsDateRange24wControl:
		return true
	}
	return false
}

// Format results are returned in.
type AnnotationOutageGetParamsFormat string

const (
	AnnotationOutageGetParamsFormatJson AnnotationOutageGetParamsFormat = "JSON"
	AnnotationOutageGetParamsFormatCsv  AnnotationOutageGetParamsFormat = "CSV"
)

func (r AnnotationOutageGetParamsFormat) IsKnown() bool {
	switch r {
	case AnnotationOutageGetParamsFormatJson, AnnotationOutageGetParamsFormatCsv:
		return true
	}
	return false
}

type AnnotationOutageGetResponseEnvelope struct {
	Result  AnnotationOutageGetResponse             `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    annotationOutageGetResponseEnvelopeJSON `json:"-"`
}

// annotationOutageGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AnnotationOutageGetResponseEnvelope]
type annotationOutageGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationOutageGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationOutageGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AnnotationOutageLocationsParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Shorthand date ranges for the last X days - use when you don't need specific
	// start and end dates.
	DateRange param.Field[AnnotationOutageLocationsParamsDateRange] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AnnotationOutageLocationsParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [AnnotationOutageLocationsParams]'s query parameters as
// `url.Values`.
func (r AnnotationOutageLocationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Shorthand date ranges for the last X days - use when you don't need specific
// start and end dates.
type AnnotationOutageLocationsParamsDateRange string

const (
	AnnotationOutageLocationsParamsDateRange1d         AnnotationOutageLocationsParamsDateRange = "1d"
	AnnotationOutageLocationsParamsDateRange2d         AnnotationOutageLocationsParamsDateRange = "2d"
	AnnotationOutageLocationsParamsDateRange7d         AnnotationOutageLocationsParamsDateRange = "7d"
	AnnotationOutageLocationsParamsDateRange14d        AnnotationOutageLocationsParamsDateRange = "14d"
	AnnotationOutageLocationsParamsDateRange28d        AnnotationOutageLocationsParamsDateRange = "28d"
	AnnotationOutageLocationsParamsDateRange12w        AnnotationOutageLocationsParamsDateRange = "12w"
	AnnotationOutageLocationsParamsDateRange24w        AnnotationOutageLocationsParamsDateRange = "24w"
	AnnotationOutageLocationsParamsDateRange52w        AnnotationOutageLocationsParamsDateRange = "52w"
	AnnotationOutageLocationsParamsDateRange1dControl  AnnotationOutageLocationsParamsDateRange = "1dControl"
	AnnotationOutageLocationsParamsDateRange2dControl  AnnotationOutageLocationsParamsDateRange = "2dControl"
	AnnotationOutageLocationsParamsDateRange7dControl  AnnotationOutageLocationsParamsDateRange = "7dControl"
	AnnotationOutageLocationsParamsDateRange14dControl AnnotationOutageLocationsParamsDateRange = "14dControl"
	AnnotationOutageLocationsParamsDateRange28dControl AnnotationOutageLocationsParamsDateRange = "28dControl"
	AnnotationOutageLocationsParamsDateRange12wControl AnnotationOutageLocationsParamsDateRange = "12wControl"
	AnnotationOutageLocationsParamsDateRange24wControl AnnotationOutageLocationsParamsDateRange = "24wControl"
)

func (r AnnotationOutageLocationsParamsDateRange) IsKnown() bool {
	switch r {
	case AnnotationOutageLocationsParamsDateRange1d, AnnotationOutageLocationsParamsDateRange2d, AnnotationOutageLocationsParamsDateRange7d, AnnotationOutageLocationsParamsDateRange14d, AnnotationOutageLocationsParamsDateRange28d, AnnotationOutageLocationsParamsDateRange12w, AnnotationOutageLocationsParamsDateRange24w, AnnotationOutageLocationsParamsDateRange52w, AnnotationOutageLocationsParamsDateRange1dControl, AnnotationOutageLocationsParamsDateRange2dControl, AnnotationOutageLocationsParamsDateRange7dControl, AnnotationOutageLocationsParamsDateRange14dControl, AnnotationOutageLocationsParamsDateRange28dControl, AnnotationOutageLocationsParamsDateRange12wControl, AnnotationOutageLocationsParamsDateRange24wControl:
		return true
	}
	return false
}

// Format results are returned in.
type AnnotationOutageLocationsParamsFormat string

const (
	AnnotationOutageLocationsParamsFormatJson AnnotationOutageLocationsParamsFormat = "JSON"
	AnnotationOutageLocationsParamsFormatCsv  AnnotationOutageLocationsParamsFormat = "CSV"
)

func (r AnnotationOutageLocationsParamsFormat) IsKnown() bool {
	switch r {
	case AnnotationOutageLocationsParamsFormatJson, AnnotationOutageLocationsParamsFormatCsv:
		return true
	}
	return false
}

type AnnotationOutageLocationsResponseEnvelope struct {
	Result  AnnotationOutageLocationsResponse             `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    annotationOutageLocationsResponseEnvelopeJSON `json:"-"`
}

// annotationOutageLocationsResponseEnvelopeJSON contains the JSON metadata for the
// struct [AnnotationOutageLocationsResponseEnvelope]
type annotationOutageLocationsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationOutageLocationsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationOutageLocationsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
