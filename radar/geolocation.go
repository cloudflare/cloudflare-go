// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// GeolocationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGeolocationService] method instead.
type GeolocationService struct {
	Options []option.RequestOption
}

// NewGeolocationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGeolocationService(opts ...option.RequestOption) (r *GeolocationService) {
	r = &GeolocationService{}
	r.Options = opts
	return
}

// Retrieves a list of geolocations.
func (r *GeolocationService) List(ctx context.Context, query GeolocationListParams, opts ...option.RequestOption) (res *GeolocationListResponse, err error) {
	var env GeolocationListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/geolocations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the requested Geolocation information.
func (r *GeolocationService) Get(ctx context.Context, geoID string, query GeolocationGetParams, opts ...option.RequestOption) (res *GeolocationGetResponse, err error) {
	var env GeolocationGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if geoID == "" {
		err = errors.New("missing required geo_id parameter")
		return
	}
	path := fmt.Sprintf("radar/geolocations/%s", geoID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type GeolocationListResponse struct {
	Geolocations []GeolocationListResponseGeolocation `json:"geolocations,required"`
	JSON         geolocationListResponseJSON          `json:"-"`
}

// geolocationListResponseJSON contains the JSON metadata for the struct
// [GeolocationListResponse]
type geolocationListResponseJSON struct {
	Geolocations apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *GeolocationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r geolocationListResponseJSON) RawJSON() string {
	return r.raw
}

type GeolocationListResponseGeolocation struct {
	GeoID string `json:"geoId,required"`
	// A numeric string.
	Latitude string `json:"latitude,required"`
	// A numeric string.
	Longitude string                                    `json:"longitude,required"`
	Name      string                                    `json:"name,required"`
	Parent    GeolocationListResponseGeolocationsParent `json:"parent,required"`
	// The type of the geolocation.
	Type GeolocationListResponseGeolocationsType `json:"type,required"`
	JSON geolocationListResponseGeolocationJSON  `json:"-"`
}

// geolocationListResponseGeolocationJSON contains the JSON metadata for the struct
// [GeolocationListResponseGeolocation]
type geolocationListResponseGeolocationJSON struct {
	GeoID       apijson.Field
	Latitude    apijson.Field
	Longitude   apijson.Field
	Name        apijson.Field
	Parent      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GeolocationListResponseGeolocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r geolocationListResponseGeolocationJSON) RawJSON() string {
	return r.raw
}

type GeolocationListResponseGeolocationsParent struct {
	GeoID string `json:"geoId,required"`
	// A numeric string.
	Latitude string `json:"latitude,required"`
	// A numeric string.
	Longitude string                                          `json:"longitude,required"`
	Name      string                                          `json:"name,required"`
	Parent    GeolocationListResponseGeolocationsParentParent `json:"parent,required"`
	// The type of the geolocation.
	Type GeolocationListResponseGeolocationsParentType `json:"type,required"`
	JSON geolocationListResponseGeolocationsParentJSON `json:"-"`
}

// geolocationListResponseGeolocationsParentJSON contains the JSON metadata for the
// struct [GeolocationListResponseGeolocationsParent]
type geolocationListResponseGeolocationsParentJSON struct {
	GeoID       apijson.Field
	Latitude    apijson.Field
	Longitude   apijson.Field
	Name        apijson.Field
	Parent      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GeolocationListResponseGeolocationsParent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r geolocationListResponseGeolocationsParentJSON) RawJSON() string {
	return r.raw
}

type GeolocationListResponseGeolocationsParentParent struct {
	GeoID string `json:"geoId,required"`
	// A numeric string.
	Latitude string `json:"latitude,required"`
	// A numeric string.
	Longitude string `json:"longitude,required"`
	Name      string `json:"name,required"`
	// The type of the geolocation.
	Type GeolocationListResponseGeolocationsParentParentType `json:"type,required"`
	JSON geolocationListResponseGeolocationsParentParentJSON `json:"-"`
}

// geolocationListResponseGeolocationsParentParentJSON contains the JSON metadata
// for the struct [GeolocationListResponseGeolocationsParentParent]
type geolocationListResponseGeolocationsParentParentJSON struct {
	GeoID       apijson.Field
	Latitude    apijson.Field
	Longitude   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GeolocationListResponseGeolocationsParentParent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r geolocationListResponseGeolocationsParentParentJSON) RawJSON() string {
	return r.raw
}

// The type of the geolocation.
type GeolocationListResponseGeolocationsParentParentType string

const (
	GeolocationListResponseGeolocationsParentParentTypeContinent GeolocationListResponseGeolocationsParentParentType = "CONTINENT"
	GeolocationListResponseGeolocationsParentParentTypeCountry   GeolocationListResponseGeolocationsParentParentType = "COUNTRY"
	GeolocationListResponseGeolocationsParentParentTypeAdm1      GeolocationListResponseGeolocationsParentParentType = "ADM1"
)

func (r GeolocationListResponseGeolocationsParentParentType) IsKnown() bool {
	switch r {
	case GeolocationListResponseGeolocationsParentParentTypeContinent, GeolocationListResponseGeolocationsParentParentTypeCountry, GeolocationListResponseGeolocationsParentParentTypeAdm1:
		return true
	}
	return false
}

// The type of the geolocation.
type GeolocationListResponseGeolocationsParentType string

const (
	GeolocationListResponseGeolocationsParentTypeContinent GeolocationListResponseGeolocationsParentType = "CONTINENT"
	GeolocationListResponseGeolocationsParentTypeCountry   GeolocationListResponseGeolocationsParentType = "COUNTRY"
	GeolocationListResponseGeolocationsParentTypeAdm1      GeolocationListResponseGeolocationsParentType = "ADM1"
)

func (r GeolocationListResponseGeolocationsParentType) IsKnown() bool {
	switch r {
	case GeolocationListResponseGeolocationsParentTypeContinent, GeolocationListResponseGeolocationsParentTypeCountry, GeolocationListResponseGeolocationsParentTypeAdm1:
		return true
	}
	return false
}

// The type of the geolocation.
type GeolocationListResponseGeolocationsType string

const (
	GeolocationListResponseGeolocationsTypeContinent GeolocationListResponseGeolocationsType = "CONTINENT"
	GeolocationListResponseGeolocationsTypeCountry   GeolocationListResponseGeolocationsType = "COUNTRY"
	GeolocationListResponseGeolocationsTypeAdm1      GeolocationListResponseGeolocationsType = "ADM1"
)

func (r GeolocationListResponseGeolocationsType) IsKnown() bool {
	switch r {
	case GeolocationListResponseGeolocationsTypeContinent, GeolocationListResponseGeolocationsTypeCountry, GeolocationListResponseGeolocationsTypeAdm1:
		return true
	}
	return false
}

type GeolocationGetResponse struct {
	Geolocation GeolocationGetResponseGeolocation `json:"geolocation,required"`
	JSON        geolocationGetResponseJSON        `json:"-"`
}

// geolocationGetResponseJSON contains the JSON metadata for the struct
// [GeolocationGetResponse]
type geolocationGetResponseJSON struct {
	Geolocation apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GeolocationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r geolocationGetResponseJSON) RawJSON() string {
	return r.raw
}

type GeolocationGetResponseGeolocation struct {
	GeoID string `json:"geoId,required"`
	// A numeric string.
	Latitude string `json:"latitude,required"`
	// A numeric string.
	Longitude string                                  `json:"longitude,required"`
	Name      string                                  `json:"name,required"`
	Parent    GeolocationGetResponseGeolocationParent `json:"parent,required"`
	// The type of the geolocation.
	Type GeolocationGetResponseGeolocationType `json:"type,required"`
	JSON geolocationGetResponseGeolocationJSON `json:"-"`
}

// geolocationGetResponseGeolocationJSON contains the JSON metadata for the struct
// [GeolocationGetResponseGeolocation]
type geolocationGetResponseGeolocationJSON struct {
	GeoID       apijson.Field
	Latitude    apijson.Field
	Longitude   apijson.Field
	Name        apijson.Field
	Parent      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GeolocationGetResponseGeolocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r geolocationGetResponseGeolocationJSON) RawJSON() string {
	return r.raw
}

type GeolocationGetResponseGeolocationParent struct {
	GeoID string `json:"geoId,required"`
	// A numeric string.
	Latitude string `json:"latitude,required"`
	// A numeric string.
	Longitude string                                        `json:"longitude,required"`
	Name      string                                        `json:"name,required"`
	Parent    GeolocationGetResponseGeolocationParentParent `json:"parent,required"`
	// The type of the geolocation.
	Type GeolocationGetResponseGeolocationParentType `json:"type,required"`
	JSON geolocationGetResponseGeolocationParentJSON `json:"-"`
}

// geolocationGetResponseGeolocationParentJSON contains the JSON metadata for the
// struct [GeolocationGetResponseGeolocationParent]
type geolocationGetResponseGeolocationParentJSON struct {
	GeoID       apijson.Field
	Latitude    apijson.Field
	Longitude   apijson.Field
	Name        apijson.Field
	Parent      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GeolocationGetResponseGeolocationParent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r geolocationGetResponseGeolocationParentJSON) RawJSON() string {
	return r.raw
}

type GeolocationGetResponseGeolocationParentParent struct {
	GeoID string `json:"geoId,required"`
	// A numeric string.
	Latitude string `json:"latitude,required"`
	// A numeric string.
	Longitude string `json:"longitude,required"`
	Name      string `json:"name,required"`
	// The type of the geolocation.
	Type GeolocationGetResponseGeolocationParentParentType `json:"type,required"`
	JSON geolocationGetResponseGeolocationParentParentJSON `json:"-"`
}

// geolocationGetResponseGeolocationParentParentJSON contains the JSON metadata for
// the struct [GeolocationGetResponseGeolocationParentParent]
type geolocationGetResponseGeolocationParentParentJSON struct {
	GeoID       apijson.Field
	Latitude    apijson.Field
	Longitude   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GeolocationGetResponseGeolocationParentParent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r geolocationGetResponseGeolocationParentParentJSON) RawJSON() string {
	return r.raw
}

// The type of the geolocation.
type GeolocationGetResponseGeolocationParentParentType string

const (
	GeolocationGetResponseGeolocationParentParentTypeContinent GeolocationGetResponseGeolocationParentParentType = "CONTINENT"
	GeolocationGetResponseGeolocationParentParentTypeCountry   GeolocationGetResponseGeolocationParentParentType = "COUNTRY"
	GeolocationGetResponseGeolocationParentParentTypeAdm1      GeolocationGetResponseGeolocationParentParentType = "ADM1"
)

func (r GeolocationGetResponseGeolocationParentParentType) IsKnown() bool {
	switch r {
	case GeolocationGetResponseGeolocationParentParentTypeContinent, GeolocationGetResponseGeolocationParentParentTypeCountry, GeolocationGetResponseGeolocationParentParentTypeAdm1:
		return true
	}
	return false
}

// The type of the geolocation.
type GeolocationGetResponseGeolocationParentType string

const (
	GeolocationGetResponseGeolocationParentTypeContinent GeolocationGetResponseGeolocationParentType = "CONTINENT"
	GeolocationGetResponseGeolocationParentTypeCountry   GeolocationGetResponseGeolocationParentType = "COUNTRY"
	GeolocationGetResponseGeolocationParentTypeAdm1      GeolocationGetResponseGeolocationParentType = "ADM1"
)

func (r GeolocationGetResponseGeolocationParentType) IsKnown() bool {
	switch r {
	case GeolocationGetResponseGeolocationParentTypeContinent, GeolocationGetResponseGeolocationParentTypeCountry, GeolocationGetResponseGeolocationParentTypeAdm1:
		return true
	}
	return false
}

// The type of the geolocation.
type GeolocationGetResponseGeolocationType string

const (
	GeolocationGetResponseGeolocationTypeContinent GeolocationGetResponseGeolocationType = "CONTINENT"
	GeolocationGetResponseGeolocationTypeCountry   GeolocationGetResponseGeolocationType = "COUNTRY"
	GeolocationGetResponseGeolocationTypeAdm1      GeolocationGetResponseGeolocationType = "ADM1"
)

func (r GeolocationGetResponseGeolocationType) IsKnown() bool {
	switch r {
	case GeolocationGetResponseGeolocationTypeContinent, GeolocationGetResponseGeolocationTypeCountry, GeolocationGetResponseGeolocationTypeAdm1:
		return true
	}
	return false
}

type GeolocationListParams struct {
	// Format in which results will be returned.
	Format param.Field[GeolocationListParamsFormat] `query:"format"`
	// Filters results by geolocation. Specify a comma-separated list of GeoNames IDs.
	GeoID param.Field[string] `query:"geoId"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 location
	// codes.
	Location param.Field[string] `query:"location"`
	// Skips the specified number of objects before fetching the results.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [GeolocationListParams]'s query parameters as `url.Values`.
func (r GeolocationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type GeolocationListParamsFormat string

const (
	GeolocationListParamsFormatJson GeolocationListParamsFormat = "JSON"
	GeolocationListParamsFormatCsv  GeolocationListParamsFormat = "CSV"
)

func (r GeolocationListParamsFormat) IsKnown() bool {
	switch r {
	case GeolocationListParamsFormatJson, GeolocationListParamsFormatCsv:
		return true
	}
	return false
}

type GeolocationListResponseEnvelope struct {
	Result  GeolocationListResponse             `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    geolocationListResponseEnvelopeJSON `json:"-"`
}

// geolocationListResponseEnvelopeJSON contains the JSON metadata for the struct
// [GeolocationListResponseEnvelope]
type geolocationListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GeolocationListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r geolocationListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type GeolocationGetParams struct {
	// Format in which results will be returned.
	Format param.Field[GeolocationGetParamsFormat] `query:"format"`
}

// URLQuery serializes [GeolocationGetParams]'s query parameters as `url.Values`.
func (r GeolocationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type GeolocationGetParamsFormat string

const (
	GeolocationGetParamsFormatJson GeolocationGetParamsFormat = "JSON"
	GeolocationGetParamsFormatCsv  GeolocationGetParamsFormat = "CSV"
)

func (r GeolocationGetParamsFormat) IsKnown() bool {
	switch r {
	case GeolocationGetParamsFormatJson, GeolocationGetParamsFormatCsv:
		return true
	}
	return false
}

type GeolocationGetResponseEnvelope struct {
	Result  GeolocationGetResponse             `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    geolocationGetResponseEnvelopeJSON `json:"-"`
}

// geolocationGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [GeolocationGetResponseEnvelope]
type geolocationGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GeolocationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r geolocationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
