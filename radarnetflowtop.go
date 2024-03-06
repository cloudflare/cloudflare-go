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

// RadarNetflowTopService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarNetflowTopService] method
// instead.
type RadarNetflowTopService struct {
	Options []option.RequestOption
}

// NewRadarNetflowTopService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarNetflowTopService(opts ...option.RequestOption) (r *RadarNetflowTopService) {
	r = &RadarNetflowTopService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS) by network traffic (NetFlows) over a given
// time period. Visit https://en.wikipedia.org/wiki/NetFlow for more information.
func (r *RadarNetflowTopService) Ases(ctx context.Context, query RadarNetflowTopAsesParams, opts ...option.RequestOption) (res *RadarNetflowTopAsesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarNetflowTopAsesResponseEnvelope
	path := "radar/netflows/top/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the top locations by network traffic (NetFlows) over a given time period.
// Visit https://en.wikipedia.org/wiki/NetFlow for more information.
func (r *RadarNetflowTopService) Locations(ctx context.Context, query RadarNetflowTopLocationsParams, opts ...option.RequestOption) (res *RadarNetflowTopLocationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarNetflowTopLocationsResponseEnvelope
	path := "radar/netflows/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarNetflowTopAsesResponse struct {
	Top0 []RadarNetflowTopAsesResponseTop0 `json:"top_0,required"`
	JSON radarNetflowTopAsesResponseJSON   `json:"-"`
}

// radarNetflowTopAsesResponseJSON contains the JSON metadata for the struct
// [RadarNetflowTopAsesResponse]
type radarNetflowTopAsesResponseJSON struct {
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowTopAsesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowTopAsesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowTopAsesResponseTop0 struct {
	ClientASN    float64                             `json:"clientASN,required"`
	ClientAsName string                              `json:"clientASName,required"`
	Value        string                              `json:"value,required"`
	JSON         radarNetflowTopAsesResponseTop0JSON `json:"-"`
}

// radarNetflowTopAsesResponseTop0JSON contains the JSON metadata for the struct
// [RadarNetflowTopAsesResponseTop0]
type radarNetflowTopAsesResponseTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarNetflowTopAsesResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowTopAsesResponseTop0JSON) RawJSON() string {
	return r.raw
}

type RadarNetflowTopLocationsResponse struct {
	Top0 []RadarNetflowTopLocationsResponseTop0 `json:"top_0,required"`
	JSON radarNetflowTopLocationsResponseJSON   `json:"-"`
}

// radarNetflowTopLocationsResponseJSON contains the JSON metadata for the struct
// [RadarNetflowTopLocationsResponse]
type radarNetflowTopLocationsResponseJSON struct {
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowTopLocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowTopLocationsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowTopLocationsResponseTop0 struct {
	ClientCountryAlpha2 string                                   `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                   `json:"clientCountryName,required"`
	Value               string                                   `json:"value,required"`
	JSON                radarNetflowTopLocationsResponseTop0JSON `json:"-"`
}

// radarNetflowTopLocationsResponseTop0JSON contains the JSON metadata for the
// struct [RadarNetflowTopLocationsResponseTop0]
type radarNetflowTopLocationsResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarNetflowTopLocationsResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowTopLocationsResponseTop0JSON) RawJSON() string {
	return r.raw
}

type RadarNetflowTopAsesParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarNetflowTopAsesParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarNetflowTopAsesParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarNetflowTopAsesParams]'s query parameters as
// `url.Values`.
func (r RadarNetflowTopAsesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarNetflowTopAsesParamsDateRange string

const (
	RadarNetflowTopAsesParamsDateRange1d         RadarNetflowTopAsesParamsDateRange = "1d"
	RadarNetflowTopAsesParamsDateRange2d         RadarNetflowTopAsesParamsDateRange = "2d"
	RadarNetflowTopAsesParamsDateRange7d         RadarNetflowTopAsesParamsDateRange = "7d"
	RadarNetflowTopAsesParamsDateRange14d        RadarNetflowTopAsesParamsDateRange = "14d"
	RadarNetflowTopAsesParamsDateRange28d        RadarNetflowTopAsesParamsDateRange = "28d"
	RadarNetflowTopAsesParamsDateRange12w        RadarNetflowTopAsesParamsDateRange = "12w"
	RadarNetflowTopAsesParamsDateRange24w        RadarNetflowTopAsesParamsDateRange = "24w"
	RadarNetflowTopAsesParamsDateRange52w        RadarNetflowTopAsesParamsDateRange = "52w"
	RadarNetflowTopAsesParamsDateRange1dControl  RadarNetflowTopAsesParamsDateRange = "1dControl"
	RadarNetflowTopAsesParamsDateRange2dControl  RadarNetflowTopAsesParamsDateRange = "2dControl"
	RadarNetflowTopAsesParamsDateRange7dControl  RadarNetflowTopAsesParamsDateRange = "7dControl"
	RadarNetflowTopAsesParamsDateRange14dControl RadarNetflowTopAsesParamsDateRange = "14dControl"
	RadarNetflowTopAsesParamsDateRange28dControl RadarNetflowTopAsesParamsDateRange = "28dControl"
	RadarNetflowTopAsesParamsDateRange12wControl RadarNetflowTopAsesParamsDateRange = "12wControl"
	RadarNetflowTopAsesParamsDateRange24wControl RadarNetflowTopAsesParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarNetflowTopAsesParamsFormat string

const (
	RadarNetflowTopAsesParamsFormatJson RadarNetflowTopAsesParamsFormat = "JSON"
	RadarNetflowTopAsesParamsFormatCsv  RadarNetflowTopAsesParamsFormat = "CSV"
)

type RadarNetflowTopAsesResponseEnvelope struct {
	Result  RadarNetflowTopAsesResponse             `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    radarNetflowTopAsesResponseEnvelopeJSON `json:"-"`
}

// radarNetflowTopAsesResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarNetflowTopAsesResponseEnvelope]
type radarNetflowTopAsesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowTopAsesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowTopAsesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowTopLocationsParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarNetflowTopLocationsParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarNetflowTopLocationsParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarNetflowTopLocationsParams]'s query parameters as
// `url.Values`.
func (r RadarNetflowTopLocationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarNetflowTopLocationsParamsDateRange string

const (
	RadarNetflowTopLocationsParamsDateRange1d         RadarNetflowTopLocationsParamsDateRange = "1d"
	RadarNetflowTopLocationsParamsDateRange2d         RadarNetflowTopLocationsParamsDateRange = "2d"
	RadarNetflowTopLocationsParamsDateRange7d         RadarNetflowTopLocationsParamsDateRange = "7d"
	RadarNetflowTopLocationsParamsDateRange14d        RadarNetflowTopLocationsParamsDateRange = "14d"
	RadarNetflowTopLocationsParamsDateRange28d        RadarNetflowTopLocationsParamsDateRange = "28d"
	RadarNetflowTopLocationsParamsDateRange12w        RadarNetflowTopLocationsParamsDateRange = "12w"
	RadarNetflowTopLocationsParamsDateRange24w        RadarNetflowTopLocationsParamsDateRange = "24w"
	RadarNetflowTopLocationsParamsDateRange52w        RadarNetflowTopLocationsParamsDateRange = "52w"
	RadarNetflowTopLocationsParamsDateRange1dControl  RadarNetflowTopLocationsParamsDateRange = "1dControl"
	RadarNetflowTopLocationsParamsDateRange2dControl  RadarNetflowTopLocationsParamsDateRange = "2dControl"
	RadarNetflowTopLocationsParamsDateRange7dControl  RadarNetflowTopLocationsParamsDateRange = "7dControl"
	RadarNetflowTopLocationsParamsDateRange14dControl RadarNetflowTopLocationsParamsDateRange = "14dControl"
	RadarNetflowTopLocationsParamsDateRange28dControl RadarNetflowTopLocationsParamsDateRange = "28dControl"
	RadarNetflowTopLocationsParamsDateRange12wControl RadarNetflowTopLocationsParamsDateRange = "12wControl"
	RadarNetflowTopLocationsParamsDateRange24wControl RadarNetflowTopLocationsParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarNetflowTopLocationsParamsFormat string

const (
	RadarNetflowTopLocationsParamsFormatJson RadarNetflowTopLocationsParamsFormat = "JSON"
	RadarNetflowTopLocationsParamsFormatCsv  RadarNetflowTopLocationsParamsFormat = "CSV"
)

type RadarNetflowTopLocationsResponseEnvelope struct {
	Result  RadarNetflowTopLocationsResponse             `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarNetflowTopLocationsResponseEnvelopeJSON `json:"-"`
}

// radarNetflowTopLocationsResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarNetflowTopLocationsResponseEnvelope]
type radarNetflowTopLocationsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowTopLocationsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowTopLocationsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
