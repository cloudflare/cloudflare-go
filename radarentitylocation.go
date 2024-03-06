// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarEntityLocationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEntityLocationService]
// method instead.
type RadarEntityLocationService struct {
	Options []option.RequestOption
}

// NewRadarEntityLocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarEntityLocationService(opts ...option.RequestOption) (r *RadarEntityLocationService) {
	r = &RadarEntityLocationService{}
	r.Options = opts
	return
}

// Get a list of locations.
func (r *RadarEntityLocationService) List(ctx context.Context, query RadarEntityLocationListParams, opts ...option.RequestOption) (res *RadarEntityLocationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEntityLocationListResponseEnvelope
	path := "radar/entities/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the requested location information. A confidence level below `5` indicates a
// low level of confidence in the traffic data - normally this happens because
// Cloudflare has a small amount of traffic from/to this location).
func (r *RadarEntityLocationService) Get(ctx context.Context, location string, query RadarEntityLocationGetParams, opts ...option.RequestOption) (res *RadarEntityLocationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEntityLocationGetResponseEnvelope
	path := fmt.Sprintf("radar/entities/locations/%s", location)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEntityLocationListResponse struct {
	Locations []RadarEntityLocationListResponseLocation `json:"locations,required"`
	JSON      radarEntityLocationListResponseJSON       `json:"-"`
}

// radarEntityLocationListResponseJSON contains the JSON metadata for the struct
// [RadarEntityLocationListResponse]
type radarEntityLocationListResponseJSON struct {
	Locations   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityLocationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityLocationListResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEntityLocationListResponseLocation struct {
	Alpha2    string                                      `json:"alpha2,required"`
	Latitude  string                                      `json:"latitude,required"`
	Longitude string                                      `json:"longitude,required"`
	Name      string                                      `json:"name,required"`
	JSON      radarEntityLocationListResponseLocationJSON `json:"-"`
}

// radarEntityLocationListResponseLocationJSON contains the JSON metadata for the
// struct [RadarEntityLocationListResponseLocation]
type radarEntityLocationListResponseLocationJSON struct {
	Alpha2      apijson.Field
	Latitude    apijson.Field
	Longitude   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityLocationListResponseLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityLocationListResponseLocationJSON) RawJSON() string {
	return r.raw
}

type RadarEntityLocationGetResponse struct {
	Location RadarEntityLocationGetResponseLocation `json:"location,required"`
	JSON     radarEntityLocationGetResponseJSON     `json:"-"`
}

// radarEntityLocationGetResponseJSON contains the JSON metadata for the struct
// [RadarEntityLocationGetResponse]
type radarEntityLocationGetResponseJSON struct {
	Location    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityLocationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityLocationGetResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEntityLocationGetResponseLocation struct {
	Alpha2          string                                     `json:"alpha2,required"`
	ConfidenceLevel int64                                      `json:"confidenceLevel,required"`
	Latitude        string                                     `json:"latitude,required"`
	Longitude       string                                     `json:"longitude,required"`
	Name            string                                     `json:"name,required"`
	Region          string                                     `json:"region,required"`
	Subregion       string                                     `json:"subregion,required"`
	JSON            radarEntityLocationGetResponseLocationJSON `json:"-"`
}

// radarEntityLocationGetResponseLocationJSON contains the JSON metadata for the
// struct [RadarEntityLocationGetResponseLocation]
type radarEntityLocationGetResponseLocationJSON struct {
	Alpha2          apijson.Field
	ConfidenceLevel apijson.Field
	Latitude        apijson.Field
	Longitude       apijson.Field
	Name            apijson.Field
	Region          apijson.Field
	Subregion       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarEntityLocationGetResponseLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityLocationGetResponseLocationJSON) RawJSON() string {
	return r.raw
}

type RadarEntityLocationListParams struct {
	// Format results are returned in.
	Format param.Field[RadarEntityLocationListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma separated list of locations.
	Location param.Field[string] `query:"location"`
	// Number of objects to skip before grabbing results.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [RadarEntityLocationListParams]'s query parameters as
// `url.Values`.
func (r RadarEntityLocationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarEntityLocationListParamsFormat string

const (
	RadarEntityLocationListParamsFormatJson RadarEntityLocationListParamsFormat = "JSON"
	RadarEntityLocationListParamsFormatCsv  RadarEntityLocationListParamsFormat = "CSV"
)

type RadarEntityLocationListResponseEnvelope struct {
	Result  RadarEntityLocationListResponse             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    radarEntityLocationListResponseEnvelopeJSON `json:"-"`
}

// radarEntityLocationListResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarEntityLocationListResponseEnvelope]
type radarEntityLocationListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityLocationListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityLocationListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarEntityLocationGetParams struct {
	// Format results are returned in.
	Format param.Field[RadarEntityLocationGetParamsFormat] `query:"format"`
}

// URLQuery serializes [RadarEntityLocationGetParams]'s query parameters as
// `url.Values`.
func (r RadarEntityLocationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarEntityLocationGetParamsFormat string

const (
	RadarEntityLocationGetParamsFormatJson RadarEntityLocationGetParamsFormat = "JSON"
	RadarEntityLocationGetParamsFormatCsv  RadarEntityLocationGetParamsFormat = "CSV"
)

type RadarEntityLocationGetResponseEnvelope struct {
	Result  RadarEntityLocationGetResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarEntityLocationGetResponseEnvelopeJSON `json:"-"`
}

// radarEntityLocationGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarEntityLocationGetResponseEnvelope]
type radarEntityLocationGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityLocationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityLocationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
