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

// RadarDNSTopService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarDNSTopService] method
// instead.
type RadarDNSTopService struct {
	Options []option.RequestOption
}

// NewRadarDNSTopService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarDNSTopService(opts ...option.RequestOption) (r *RadarDNSTopService) {
	r = &RadarDNSTopService{}
	r.Options = opts
	return
}

// Get top autonomous systems by DNS queries made to Cloudflare's public DNS
// resolver.
func (r *RadarDNSTopService) Ases(ctx context.Context, query RadarDNSTopAsesParams, opts ...option.RequestOption) (res *RadarDNSTopAsesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarDNSTopAsesResponseEnvelope
	path := "radar/dns/top/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get top locations by DNS queries made to Cloudflare's public DNS resolver.
func (r *RadarDNSTopService) Locations(ctx context.Context, query RadarDNSTopLocationsParams, opts ...option.RequestOption) (res *RadarDNSTopLocationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarDNSTopLocationsResponseEnvelope
	path := "radar/dns/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarDNSTopAsesResponse struct {
	Meta RadarDNSTopAsesResponseMeta   `json:"meta,required"`
	Top0 []RadarDNSTopAsesResponseTop0 `json:"top_0,required"`
	JSON radarDNSTopAsesResponseJSON   `json:"-"`
}

// radarDNSTopAsesResponseJSON contains the JSON metadata for the struct
// [RadarDNSTopAsesResponse]
type radarDNSTopAsesResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTopAsesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTopAsesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTopAsesResponseMeta struct {
	DateRange      []RadarDNSTopAsesResponseMetaDateRange    `json:"dateRange,required"`
	ConfidenceInfo RadarDNSTopAsesResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarDNSTopAsesResponseMetaJSON           `json:"-"`
}

// radarDNSTopAsesResponseMetaJSON contains the JSON metadata for the struct
// [RadarDNSTopAsesResponseMeta]
type radarDNSTopAsesResponseMetaJSON struct {
	DateRange      apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarDNSTopAsesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTopAsesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTopAsesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                `json:"startTime,required" format:"date-time"`
	JSON      radarDNSTopAsesResponseMetaDateRangeJSON `json:"-"`
}

// radarDNSTopAsesResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [RadarDNSTopAsesResponseMetaDateRange]
type radarDNSTopAsesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTopAsesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTopAsesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTopAsesResponseMetaConfidenceInfo struct {
	Annotations []RadarDNSTopAsesResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                 `json:"level"`
	JSON        radarDNSTopAsesResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarDNSTopAsesResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [RadarDNSTopAsesResponseMetaConfidenceInfo]
type radarDNSTopAsesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTopAsesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTopAsesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTopAsesResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                  `json:"dataSource,required"`
	Description     string                                                  `json:"description,required"`
	EventType       string                                                  `json:"eventType,required"`
	IsInstantaneous interface{}                                             `json:"isInstantaneous,required"`
	EndTime         time.Time                                               `json:"endTime" format:"date-time"`
	LinkedURL       string                                                  `json:"linkedUrl"`
	StartTime       time.Time                                               `json:"startTime" format:"date-time"`
	JSON            radarDNSTopAsesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarDNSTopAsesResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [RadarDNSTopAsesResponseMetaConfidenceInfoAnnotation]
type radarDNSTopAsesResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	EndTime         apijson.Field
	LinkedURL       apijson.Field
	StartTime       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarDNSTopAsesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTopAsesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTopAsesResponseTop0 struct {
	ClientASN    int64                           `json:"clientASN,required"`
	ClientAsName string                          `json:"clientASName,required"`
	Value        string                          `json:"value,required"`
	JSON         radarDNSTopAsesResponseTop0JSON `json:"-"`
}

// radarDNSTopAsesResponseTop0JSON contains the JSON metadata for the struct
// [RadarDNSTopAsesResponseTop0]
type radarDNSTopAsesResponseTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarDNSTopAsesResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTopAsesResponseTop0JSON) RawJSON() string {
	return r.raw
}

type RadarDNSTopLocationsResponse struct {
	Meta RadarDNSTopLocationsResponseMeta   `json:"meta,required"`
	Top0 []RadarDNSTopLocationsResponseTop0 `json:"top_0,required"`
	JSON radarDNSTopLocationsResponseJSON   `json:"-"`
}

// radarDNSTopLocationsResponseJSON contains the JSON metadata for the struct
// [RadarDNSTopLocationsResponse]
type radarDNSTopLocationsResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTopLocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTopLocationsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTopLocationsResponseMeta struct {
	DateRange      []RadarDNSTopLocationsResponseMetaDateRange    `json:"dateRange,required"`
	ConfidenceInfo RadarDNSTopLocationsResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarDNSTopLocationsResponseMetaJSON           `json:"-"`
}

// radarDNSTopLocationsResponseMetaJSON contains the JSON metadata for the struct
// [RadarDNSTopLocationsResponseMeta]
type radarDNSTopLocationsResponseMetaJSON struct {
	DateRange      apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarDNSTopLocationsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTopLocationsResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTopLocationsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                     `json:"startTime,required" format:"date-time"`
	JSON      radarDNSTopLocationsResponseMetaDateRangeJSON `json:"-"`
}

// radarDNSTopLocationsResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [RadarDNSTopLocationsResponseMetaDateRange]
type radarDNSTopLocationsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTopLocationsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTopLocationsResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTopLocationsResponseMetaConfidenceInfo struct {
	Annotations []RadarDNSTopLocationsResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                      `json:"level"`
	JSON        radarDNSTopLocationsResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarDNSTopLocationsResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [RadarDNSTopLocationsResponseMetaConfidenceInfo]
type radarDNSTopLocationsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTopLocationsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTopLocationsResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTopLocationsResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                       `json:"dataSource,required"`
	Description     string                                                       `json:"description,required"`
	EventType       string                                                       `json:"eventType,required"`
	IsInstantaneous interface{}                                                  `json:"isInstantaneous,required"`
	EndTime         time.Time                                                    `json:"endTime" format:"date-time"`
	LinkedURL       string                                                       `json:"linkedUrl"`
	StartTime       time.Time                                                    `json:"startTime" format:"date-time"`
	JSON            radarDNSTopLocationsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarDNSTopLocationsResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [RadarDNSTopLocationsResponseMetaConfidenceInfoAnnotation]
type radarDNSTopLocationsResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	EndTime         apijson.Field
	LinkedURL       apijson.Field
	StartTime       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarDNSTopLocationsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTopLocationsResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTopLocationsResponseTop0 struct {
	ClientCountryAlpha2 string                               `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                               `json:"clientCountryName,required"`
	Value               string                               `json:"value,required"`
	JSON                radarDNSTopLocationsResponseTop0JSON `json:"-"`
}

// radarDNSTopLocationsResponseTop0JSON contains the JSON metadata for the struct
// [RadarDNSTopLocationsResponseTop0]
type radarDNSTopLocationsResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarDNSTopLocationsResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTopLocationsResponseTop0JSON) RawJSON() string {
	return r.raw
}

type RadarDNSTopAsesParams struct {
	// Array of domain names.
	Domain param.Field[[]string] `query:"domain,required"`
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
	DateRange param.Field[[]RadarDNSTopAsesParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarDNSTopAsesParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarDNSTopAsesParams]'s query parameters as `url.Values`.
func (r RadarDNSTopAsesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarDNSTopAsesParamsDateRange string

const (
	RadarDNSTopAsesParamsDateRange1d         RadarDNSTopAsesParamsDateRange = "1d"
	RadarDNSTopAsesParamsDateRange2d         RadarDNSTopAsesParamsDateRange = "2d"
	RadarDNSTopAsesParamsDateRange7d         RadarDNSTopAsesParamsDateRange = "7d"
	RadarDNSTopAsesParamsDateRange14d        RadarDNSTopAsesParamsDateRange = "14d"
	RadarDNSTopAsesParamsDateRange28d        RadarDNSTopAsesParamsDateRange = "28d"
	RadarDNSTopAsesParamsDateRange12w        RadarDNSTopAsesParamsDateRange = "12w"
	RadarDNSTopAsesParamsDateRange24w        RadarDNSTopAsesParamsDateRange = "24w"
	RadarDNSTopAsesParamsDateRange52w        RadarDNSTopAsesParamsDateRange = "52w"
	RadarDNSTopAsesParamsDateRange1dControl  RadarDNSTopAsesParamsDateRange = "1dControl"
	RadarDNSTopAsesParamsDateRange2dControl  RadarDNSTopAsesParamsDateRange = "2dControl"
	RadarDNSTopAsesParamsDateRange7dControl  RadarDNSTopAsesParamsDateRange = "7dControl"
	RadarDNSTopAsesParamsDateRange14dControl RadarDNSTopAsesParamsDateRange = "14dControl"
	RadarDNSTopAsesParamsDateRange28dControl RadarDNSTopAsesParamsDateRange = "28dControl"
	RadarDNSTopAsesParamsDateRange12wControl RadarDNSTopAsesParamsDateRange = "12wControl"
	RadarDNSTopAsesParamsDateRange24wControl RadarDNSTopAsesParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarDNSTopAsesParamsFormat string

const (
	RadarDNSTopAsesParamsFormatJson RadarDNSTopAsesParamsFormat = "JSON"
	RadarDNSTopAsesParamsFormatCsv  RadarDNSTopAsesParamsFormat = "CSV"
)

type RadarDNSTopAsesResponseEnvelope struct {
	Result  RadarDNSTopAsesResponse             `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    radarDNSTopAsesResponseEnvelopeJSON `json:"-"`
}

// radarDNSTopAsesResponseEnvelopeJSON contains the JSON metadata for the struct
// [RadarDNSTopAsesResponseEnvelope]
type radarDNSTopAsesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTopAsesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTopAsesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTopLocationsParams struct {
	// Array of domain names.
	Domain param.Field[[]string] `query:"domain,required"`
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
	DateRange param.Field[[]RadarDNSTopLocationsParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarDNSTopLocationsParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarDNSTopLocationsParams]'s query parameters as
// `url.Values`.
func (r RadarDNSTopLocationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarDNSTopLocationsParamsDateRange string

const (
	RadarDNSTopLocationsParamsDateRange1d         RadarDNSTopLocationsParamsDateRange = "1d"
	RadarDNSTopLocationsParamsDateRange2d         RadarDNSTopLocationsParamsDateRange = "2d"
	RadarDNSTopLocationsParamsDateRange7d         RadarDNSTopLocationsParamsDateRange = "7d"
	RadarDNSTopLocationsParamsDateRange14d        RadarDNSTopLocationsParamsDateRange = "14d"
	RadarDNSTopLocationsParamsDateRange28d        RadarDNSTopLocationsParamsDateRange = "28d"
	RadarDNSTopLocationsParamsDateRange12w        RadarDNSTopLocationsParamsDateRange = "12w"
	RadarDNSTopLocationsParamsDateRange24w        RadarDNSTopLocationsParamsDateRange = "24w"
	RadarDNSTopLocationsParamsDateRange52w        RadarDNSTopLocationsParamsDateRange = "52w"
	RadarDNSTopLocationsParamsDateRange1dControl  RadarDNSTopLocationsParamsDateRange = "1dControl"
	RadarDNSTopLocationsParamsDateRange2dControl  RadarDNSTopLocationsParamsDateRange = "2dControl"
	RadarDNSTopLocationsParamsDateRange7dControl  RadarDNSTopLocationsParamsDateRange = "7dControl"
	RadarDNSTopLocationsParamsDateRange14dControl RadarDNSTopLocationsParamsDateRange = "14dControl"
	RadarDNSTopLocationsParamsDateRange28dControl RadarDNSTopLocationsParamsDateRange = "28dControl"
	RadarDNSTopLocationsParamsDateRange12wControl RadarDNSTopLocationsParamsDateRange = "12wControl"
	RadarDNSTopLocationsParamsDateRange24wControl RadarDNSTopLocationsParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarDNSTopLocationsParamsFormat string

const (
	RadarDNSTopLocationsParamsFormatJson RadarDNSTopLocationsParamsFormat = "JSON"
	RadarDNSTopLocationsParamsFormatCsv  RadarDNSTopLocationsParamsFormat = "CSV"
)

type RadarDNSTopLocationsResponseEnvelope struct {
	Result  RadarDNSTopLocationsResponse             `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    radarDNSTopLocationsResponseEnvelopeJSON `json:"-"`
}

// radarDNSTopLocationsResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarDNSTopLocationsResponseEnvelope]
type radarDNSTopLocationsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTopLocationsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTopLocationsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
