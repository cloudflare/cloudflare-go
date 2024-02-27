// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarAs112TopService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarAs112TopService] method
// instead.
type RadarAs112TopService struct {
	Options []option.RequestOption
}

// NewRadarAs112TopService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarAs112TopService(opts ...option.RequestOption) (r *RadarAs112TopService) {
	r = &RadarAs112TopService{}
	r.Options = opts
	return
}

// Get the top locations by DNS queries DNSSEC support to AS112.
func (r *RadarAs112TopService) DNSSEC(ctx context.Context, dnssec RadarAs112TopDNSSECParamsDNSSEC, query RadarAs112TopDNSSECParams, opts ...option.RequestOption) (res *RadarAs112TopDNSSECResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAs112TopDNSSECResponseEnvelope
	path := fmt.Sprintf("radar/as112/top/locations/dnssec/%v", dnssec)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the top locations, by DNS queries EDNS support to AS112.
func (r *RadarAs112TopService) Edns(ctx context.Context, edns RadarAs112TopEdnsParamsEdns, query RadarAs112TopEdnsParams, opts ...option.RequestOption) (res *RadarAs112TopEdnsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAs112TopEdnsResponseEnvelope
	path := fmt.Sprintf("radar/as112/top/locations/edns/%v", edns)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the top locations by DNS queries IP version to AS112.
func (r *RadarAs112TopService) IPVersion(ctx context.Context, ipVersion RadarAs112TopIPVersionParamsIPVersion, query RadarAs112TopIPVersionParams, opts ...option.RequestOption) (res *RadarAs112TopIPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAs112TopIPVersionResponseEnvelope
	path := fmt.Sprintf("radar/as112/top/locations/ip_version/%v", ipVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the top locations by AS112 DNS queries. Values are a percentage out of the
// total queries.
func (r *RadarAs112TopService) Locations(ctx context.Context, query RadarAs112TopLocationsParams, opts ...option.RequestOption) (res *RadarAs112TopLocationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAs112TopLocationsResponseEnvelope
	path := "radar/as112/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAs112TopDNSSECResponse struct {
	Meta RadarAs112TopDNSSECResponseMeta   `json:"meta,required"`
	Top0 []RadarAs112TopDNSSECResponseTop0 `json:"top_0,required"`
	JSON radarAs112TopDNSSECResponseJSON   `json:"-"`
}

// radarAs112TopDNSSECResponseJSON contains the JSON metadata for the struct
// [RadarAs112TopDNSSECResponse]
type radarAs112TopDNSSECResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopDNSSECResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopDNSSECResponseMeta struct {
	DateRange      []RadarAs112TopDNSSECResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                        `json:"lastUpdated,required"`
	ConfidenceInfo RadarAs112TopDNSSECResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAs112TopDNSSECResponseMetaJSON           `json:"-"`
}

// radarAs112TopDNSSECResponseMetaJSON contains the JSON metadata for the struct
// [RadarAs112TopDNSSECResponseMeta]
type radarAs112TopDNSSECResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112TopDNSSECResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopDNSSECResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                    `json:"startTime,required" format:"date-time"`
	JSON      radarAs112TopDNSSECResponseMetaDateRangeJSON `json:"-"`
}

// radarAs112TopDNSSECResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [RadarAs112TopDNSSECResponseMetaDateRange]
type radarAs112TopDNSSECResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopDNSSECResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopDNSSECResponseMetaConfidenceInfo struct {
	Annotations []RadarAs112TopDNSSECResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                     `json:"level"`
	JSON        radarAs112TopDNSSECResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAs112TopDNSSECResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [RadarAs112TopDNSSECResponseMetaConfidenceInfo]
type radarAs112TopDNSSECResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopDNSSECResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopDNSSECResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                      `json:"dataSource,required"`
	Description     string                                                      `json:"description,required"`
	EventType       string                                                      `json:"eventType,required"`
	IsInstantaneous interface{}                                                 `json:"isInstantaneous,required"`
	EndTime         time.Time                                                   `json:"endTime" format:"date-time"`
	LinkedURL       string                                                      `json:"linkedUrl"`
	StartTime       time.Time                                                   `json:"startTime" format:"date-time"`
	JSON            radarAs112TopDNSSECResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112TopDNSSECResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [RadarAs112TopDNSSECResponseMetaConfidenceInfoAnnotation]
type radarAs112TopDNSSECResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAs112TopDNSSECResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopDNSSECResponseTop0 struct {
	ClientCountryAlpha2 string                              `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                              `json:"clientCountryName,required"`
	Value               string                              `json:"value,required"`
	JSON                radarAs112TopDNSSECResponseTop0JSON `json:"-"`
}

// radarAs112TopDNSSECResponseTop0JSON contains the JSON metadata for the struct
// [RadarAs112TopDNSSECResponseTop0]
type radarAs112TopDNSSECResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAs112TopDNSSECResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopEdnsResponse struct {
	Meta RadarAs112TopEdnsResponseMeta   `json:"meta,required"`
	Top0 []RadarAs112TopEdnsResponseTop0 `json:"top_0,required"`
	JSON radarAs112TopEdnsResponseJSON   `json:"-"`
}

// radarAs112TopEdnsResponseJSON contains the JSON metadata for the struct
// [RadarAs112TopEdnsResponse]
type radarAs112TopEdnsResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopEdnsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopEdnsResponseMeta struct {
	DateRange      []RadarAs112TopEdnsResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                      `json:"lastUpdated,required"`
	ConfidenceInfo RadarAs112TopEdnsResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAs112TopEdnsResponseMetaJSON           `json:"-"`
}

// radarAs112TopEdnsResponseMetaJSON contains the JSON metadata for the struct
// [RadarAs112TopEdnsResponseMeta]
type radarAs112TopEdnsResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112TopEdnsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopEdnsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                  `json:"startTime,required" format:"date-time"`
	JSON      radarAs112TopEdnsResponseMetaDateRangeJSON `json:"-"`
}

// radarAs112TopEdnsResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [RadarAs112TopEdnsResponseMetaDateRange]
type radarAs112TopEdnsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopEdnsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopEdnsResponseMetaConfidenceInfo struct {
	Annotations []RadarAs112TopEdnsResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                   `json:"level"`
	JSON        radarAs112TopEdnsResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAs112TopEdnsResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [RadarAs112TopEdnsResponseMetaConfidenceInfo]
type radarAs112TopEdnsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopEdnsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopEdnsResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                    `json:"dataSource,required"`
	Description     string                                                    `json:"description,required"`
	EventType       string                                                    `json:"eventType,required"`
	IsInstantaneous interface{}                                               `json:"isInstantaneous,required"`
	EndTime         time.Time                                                 `json:"endTime" format:"date-time"`
	LinkedURL       string                                                    `json:"linkedUrl"`
	StartTime       time.Time                                                 `json:"startTime" format:"date-time"`
	JSON            radarAs112TopEdnsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112TopEdnsResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [RadarAs112TopEdnsResponseMetaConfidenceInfoAnnotation]
type radarAs112TopEdnsResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAs112TopEdnsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopEdnsResponseTop0 struct {
	ClientCountryAlpha2 string                            `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                            `json:"clientCountryName,required"`
	Value               string                            `json:"value,required"`
	JSON                radarAs112TopEdnsResponseTop0JSON `json:"-"`
}

// radarAs112TopEdnsResponseTop0JSON contains the JSON metadata for the struct
// [RadarAs112TopEdnsResponseTop0]
type radarAs112TopEdnsResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAs112TopEdnsResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopIPVersionResponse struct {
	Meta RadarAs112TopIPVersionResponseMeta   `json:"meta,required"`
	Top0 []RadarAs112TopIPVersionResponseTop0 `json:"top_0,required"`
	JSON radarAs112TopIPVersionResponseJSON   `json:"-"`
}

// radarAs112TopIPVersionResponseJSON contains the JSON metadata for the struct
// [RadarAs112TopIPVersionResponse]
type radarAs112TopIPVersionResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopIPVersionResponseMeta struct {
	DateRange      []RadarAs112TopIPVersionResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                           `json:"lastUpdated,required"`
	ConfidenceInfo RadarAs112TopIPVersionResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAs112TopIPVersionResponseMetaJSON           `json:"-"`
}

// radarAs112TopIPVersionResponseMetaJSON contains the JSON metadata for the struct
// [RadarAs112TopIPVersionResponseMeta]
type radarAs112TopIPVersionResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112TopIPVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopIPVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      radarAs112TopIPVersionResponseMetaDateRangeJSON `json:"-"`
}

// radarAs112TopIPVersionResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [RadarAs112TopIPVersionResponseMetaDateRange]
type radarAs112TopIPVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopIPVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopIPVersionResponseMetaConfidenceInfo struct {
	Annotations []RadarAs112TopIPVersionResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                        `json:"level"`
	JSON        radarAs112TopIPVersionResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAs112TopIPVersionResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [RadarAs112TopIPVersionResponseMetaConfidenceInfo]
type radarAs112TopIPVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopIPVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopIPVersionResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                         `json:"dataSource,required"`
	Description     string                                                         `json:"description,required"`
	EventType       string                                                         `json:"eventType,required"`
	IsInstantaneous interface{}                                                    `json:"isInstantaneous,required"`
	EndTime         time.Time                                                      `json:"endTime" format:"date-time"`
	LinkedURL       string                                                         `json:"linkedUrl"`
	StartTime       time.Time                                                      `json:"startTime" format:"date-time"`
	JSON            radarAs112TopIPVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112TopIPVersionResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [RadarAs112TopIPVersionResponseMetaConfidenceInfoAnnotation]
type radarAs112TopIPVersionResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAs112TopIPVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopIPVersionResponseTop0 struct {
	ClientCountryAlpha2 string                                 `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                 `json:"clientCountryName,required"`
	Value               string                                 `json:"value,required"`
	JSON                radarAs112TopIPVersionResponseTop0JSON `json:"-"`
}

// radarAs112TopIPVersionResponseTop0JSON contains the JSON metadata for the struct
// [RadarAs112TopIPVersionResponseTop0]
type radarAs112TopIPVersionResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAs112TopIPVersionResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationsResponse struct {
	Meta RadarAs112TopLocationsResponseMeta   `json:"meta,required"`
	Top0 []RadarAs112TopLocationsResponseTop0 `json:"top_0,required"`
	JSON radarAs112TopLocationsResponseJSON   `json:"-"`
}

// radarAs112TopLocationsResponseJSON contains the JSON metadata for the struct
// [RadarAs112TopLocationsResponse]
type radarAs112TopLocationsResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationsResponseMeta struct {
	DateRange      []RadarAs112TopLocationsResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                           `json:"lastUpdated,required"`
	ConfidenceInfo RadarAs112TopLocationsResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAs112TopLocationsResponseMetaJSON           `json:"-"`
}

// radarAs112TopLocationsResponseMetaJSON contains the JSON metadata for the struct
// [RadarAs112TopLocationsResponseMeta]
type radarAs112TopLocationsResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112TopLocationsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      radarAs112TopLocationsResponseMetaDateRangeJSON `json:"-"`
}

// radarAs112TopLocationsResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [RadarAs112TopLocationsResponseMetaDateRange]
type radarAs112TopLocationsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationsResponseMetaConfidenceInfo struct {
	Annotations []RadarAs112TopLocationsResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                        `json:"level"`
	JSON        radarAs112TopLocationsResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAs112TopLocationsResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [RadarAs112TopLocationsResponseMetaConfidenceInfo]
type radarAs112TopLocationsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationsResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                         `json:"dataSource,required"`
	Description     string                                                         `json:"description,required"`
	EventType       string                                                         `json:"eventType,required"`
	IsInstantaneous interface{}                                                    `json:"isInstantaneous,required"`
	EndTime         time.Time                                                      `json:"endTime" format:"date-time"`
	LinkedURL       string                                                         `json:"linkedUrl"`
	StartTime       time.Time                                                      `json:"startTime" format:"date-time"`
	JSON            radarAs112TopLocationsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112TopLocationsResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [RadarAs112TopLocationsResponseMetaConfidenceInfoAnnotation]
type radarAs112TopLocationsResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAs112TopLocationsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationsResponseTop0 struct {
	ClientCountryAlpha2 string                                 `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                 `json:"clientCountryName,required"`
	Value               string                                 `json:"value,required"`
	JSON                radarAs112TopLocationsResponseTop0JSON `json:"-"`
}

// radarAs112TopLocationsResponseTop0JSON contains the JSON metadata for the struct
// [RadarAs112TopLocationsResponseTop0]
type radarAs112TopLocationsResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAs112TopLocationsResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopDNSSECParams struct {
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
	DateRange param.Field[[]RadarAs112TopDNSSECParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112TopDNSSECParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TopDNSSECParams]'s query parameters as
// `url.Values`.
func (r RadarAs112TopDNSSECParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// DNSSEC.
type RadarAs112TopDNSSECParamsDNSSEC string

const (
	RadarAs112TopDNSSECParamsDNSSECSupported    RadarAs112TopDNSSECParamsDNSSEC = "SUPPORTED"
	RadarAs112TopDNSSECParamsDNSSECNotSupported RadarAs112TopDNSSECParamsDNSSEC = "NOT_SUPPORTED"
)

type RadarAs112TopDNSSECParamsDateRange string

const (
	RadarAs112TopDNSSECParamsDateRange1d         RadarAs112TopDNSSECParamsDateRange = "1d"
	RadarAs112TopDNSSECParamsDateRange2d         RadarAs112TopDNSSECParamsDateRange = "2d"
	RadarAs112TopDNSSECParamsDateRange7d         RadarAs112TopDNSSECParamsDateRange = "7d"
	RadarAs112TopDNSSECParamsDateRange14d        RadarAs112TopDNSSECParamsDateRange = "14d"
	RadarAs112TopDNSSECParamsDateRange28d        RadarAs112TopDNSSECParamsDateRange = "28d"
	RadarAs112TopDNSSECParamsDateRange12w        RadarAs112TopDNSSECParamsDateRange = "12w"
	RadarAs112TopDNSSECParamsDateRange24w        RadarAs112TopDNSSECParamsDateRange = "24w"
	RadarAs112TopDNSSECParamsDateRange52w        RadarAs112TopDNSSECParamsDateRange = "52w"
	RadarAs112TopDNSSECParamsDateRange1dControl  RadarAs112TopDNSSECParamsDateRange = "1dControl"
	RadarAs112TopDNSSECParamsDateRange2dControl  RadarAs112TopDNSSECParamsDateRange = "2dControl"
	RadarAs112TopDNSSECParamsDateRange7dControl  RadarAs112TopDNSSECParamsDateRange = "7dControl"
	RadarAs112TopDNSSECParamsDateRange14dControl RadarAs112TopDNSSECParamsDateRange = "14dControl"
	RadarAs112TopDNSSECParamsDateRange28dControl RadarAs112TopDNSSECParamsDateRange = "28dControl"
	RadarAs112TopDNSSECParamsDateRange12wControl RadarAs112TopDNSSECParamsDateRange = "12wControl"
	RadarAs112TopDNSSECParamsDateRange24wControl RadarAs112TopDNSSECParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112TopDNSSECParamsFormat string

const (
	RadarAs112TopDNSSECParamsFormatJson RadarAs112TopDNSSECParamsFormat = "JSON"
	RadarAs112TopDNSSECParamsFormatCsv  RadarAs112TopDNSSECParamsFormat = "CSV"
)

type RadarAs112TopDNSSECResponseEnvelope struct {
	Result  RadarAs112TopDNSSECResponse             `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    radarAs112TopDNSSECResponseEnvelopeJSON `json:"-"`
}

// radarAs112TopDNSSECResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarAs112TopDNSSECResponseEnvelope]
type radarAs112TopDNSSECResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopDNSSECResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopEdnsParams struct {
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
	DateRange param.Field[[]RadarAs112TopEdnsParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112TopEdnsParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TopEdnsParams]'s query parameters as
// `url.Values`.
func (r RadarAs112TopEdnsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// EDNS.
type RadarAs112TopEdnsParamsEdns string

const (
	RadarAs112TopEdnsParamsEdnsSupported    RadarAs112TopEdnsParamsEdns = "SUPPORTED"
	RadarAs112TopEdnsParamsEdnsNotSupported RadarAs112TopEdnsParamsEdns = "NOT_SUPPORTED"
)

type RadarAs112TopEdnsParamsDateRange string

const (
	RadarAs112TopEdnsParamsDateRange1d         RadarAs112TopEdnsParamsDateRange = "1d"
	RadarAs112TopEdnsParamsDateRange2d         RadarAs112TopEdnsParamsDateRange = "2d"
	RadarAs112TopEdnsParamsDateRange7d         RadarAs112TopEdnsParamsDateRange = "7d"
	RadarAs112TopEdnsParamsDateRange14d        RadarAs112TopEdnsParamsDateRange = "14d"
	RadarAs112TopEdnsParamsDateRange28d        RadarAs112TopEdnsParamsDateRange = "28d"
	RadarAs112TopEdnsParamsDateRange12w        RadarAs112TopEdnsParamsDateRange = "12w"
	RadarAs112TopEdnsParamsDateRange24w        RadarAs112TopEdnsParamsDateRange = "24w"
	RadarAs112TopEdnsParamsDateRange52w        RadarAs112TopEdnsParamsDateRange = "52w"
	RadarAs112TopEdnsParamsDateRange1dControl  RadarAs112TopEdnsParamsDateRange = "1dControl"
	RadarAs112TopEdnsParamsDateRange2dControl  RadarAs112TopEdnsParamsDateRange = "2dControl"
	RadarAs112TopEdnsParamsDateRange7dControl  RadarAs112TopEdnsParamsDateRange = "7dControl"
	RadarAs112TopEdnsParamsDateRange14dControl RadarAs112TopEdnsParamsDateRange = "14dControl"
	RadarAs112TopEdnsParamsDateRange28dControl RadarAs112TopEdnsParamsDateRange = "28dControl"
	RadarAs112TopEdnsParamsDateRange12wControl RadarAs112TopEdnsParamsDateRange = "12wControl"
	RadarAs112TopEdnsParamsDateRange24wControl RadarAs112TopEdnsParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112TopEdnsParamsFormat string

const (
	RadarAs112TopEdnsParamsFormatJson RadarAs112TopEdnsParamsFormat = "JSON"
	RadarAs112TopEdnsParamsFormatCsv  RadarAs112TopEdnsParamsFormat = "CSV"
)

type RadarAs112TopEdnsResponseEnvelope struct {
	Result  RadarAs112TopEdnsResponse             `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    radarAs112TopEdnsResponseEnvelopeJSON `json:"-"`
}

// radarAs112TopEdnsResponseEnvelopeJSON contains the JSON metadata for the struct
// [RadarAs112TopEdnsResponseEnvelope]
type radarAs112TopEdnsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopEdnsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopIPVersionParams struct {
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
	DateRange param.Field[[]RadarAs112TopIPVersionParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112TopIPVersionParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TopIPVersionParams]'s query parameters as
// `url.Values`.
func (r RadarAs112TopIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// IP Version.
type RadarAs112TopIPVersionParamsIPVersion string

const (
	RadarAs112TopIPVersionParamsIPVersionIPv4 RadarAs112TopIPVersionParamsIPVersion = "IPv4"
	RadarAs112TopIPVersionParamsIPVersionIPv6 RadarAs112TopIPVersionParamsIPVersion = "IPv6"
)

type RadarAs112TopIPVersionParamsDateRange string

const (
	RadarAs112TopIPVersionParamsDateRange1d         RadarAs112TopIPVersionParamsDateRange = "1d"
	RadarAs112TopIPVersionParamsDateRange2d         RadarAs112TopIPVersionParamsDateRange = "2d"
	RadarAs112TopIPVersionParamsDateRange7d         RadarAs112TopIPVersionParamsDateRange = "7d"
	RadarAs112TopIPVersionParamsDateRange14d        RadarAs112TopIPVersionParamsDateRange = "14d"
	RadarAs112TopIPVersionParamsDateRange28d        RadarAs112TopIPVersionParamsDateRange = "28d"
	RadarAs112TopIPVersionParamsDateRange12w        RadarAs112TopIPVersionParamsDateRange = "12w"
	RadarAs112TopIPVersionParamsDateRange24w        RadarAs112TopIPVersionParamsDateRange = "24w"
	RadarAs112TopIPVersionParamsDateRange52w        RadarAs112TopIPVersionParamsDateRange = "52w"
	RadarAs112TopIPVersionParamsDateRange1dControl  RadarAs112TopIPVersionParamsDateRange = "1dControl"
	RadarAs112TopIPVersionParamsDateRange2dControl  RadarAs112TopIPVersionParamsDateRange = "2dControl"
	RadarAs112TopIPVersionParamsDateRange7dControl  RadarAs112TopIPVersionParamsDateRange = "7dControl"
	RadarAs112TopIPVersionParamsDateRange14dControl RadarAs112TopIPVersionParamsDateRange = "14dControl"
	RadarAs112TopIPVersionParamsDateRange28dControl RadarAs112TopIPVersionParamsDateRange = "28dControl"
	RadarAs112TopIPVersionParamsDateRange12wControl RadarAs112TopIPVersionParamsDateRange = "12wControl"
	RadarAs112TopIPVersionParamsDateRange24wControl RadarAs112TopIPVersionParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112TopIPVersionParamsFormat string

const (
	RadarAs112TopIPVersionParamsFormatJson RadarAs112TopIPVersionParamsFormat = "JSON"
	RadarAs112TopIPVersionParamsFormatCsv  RadarAs112TopIPVersionParamsFormat = "CSV"
)

type RadarAs112TopIPVersionResponseEnvelope struct {
	Result  RadarAs112TopIPVersionResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarAs112TopIPVersionResponseEnvelopeJSON `json:"-"`
}

// radarAs112TopIPVersionResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarAs112TopIPVersionResponseEnvelope]
type radarAs112TopIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationsParams struct {
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
	DateRange param.Field[[]RadarAs112TopLocationsParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112TopLocationsParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TopLocationsParams]'s query parameters as
// `url.Values`.
func (r RadarAs112TopLocationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAs112TopLocationsParamsDateRange string

const (
	RadarAs112TopLocationsParamsDateRange1d         RadarAs112TopLocationsParamsDateRange = "1d"
	RadarAs112TopLocationsParamsDateRange2d         RadarAs112TopLocationsParamsDateRange = "2d"
	RadarAs112TopLocationsParamsDateRange7d         RadarAs112TopLocationsParamsDateRange = "7d"
	RadarAs112TopLocationsParamsDateRange14d        RadarAs112TopLocationsParamsDateRange = "14d"
	RadarAs112TopLocationsParamsDateRange28d        RadarAs112TopLocationsParamsDateRange = "28d"
	RadarAs112TopLocationsParamsDateRange12w        RadarAs112TopLocationsParamsDateRange = "12w"
	RadarAs112TopLocationsParamsDateRange24w        RadarAs112TopLocationsParamsDateRange = "24w"
	RadarAs112TopLocationsParamsDateRange52w        RadarAs112TopLocationsParamsDateRange = "52w"
	RadarAs112TopLocationsParamsDateRange1dControl  RadarAs112TopLocationsParamsDateRange = "1dControl"
	RadarAs112TopLocationsParamsDateRange2dControl  RadarAs112TopLocationsParamsDateRange = "2dControl"
	RadarAs112TopLocationsParamsDateRange7dControl  RadarAs112TopLocationsParamsDateRange = "7dControl"
	RadarAs112TopLocationsParamsDateRange14dControl RadarAs112TopLocationsParamsDateRange = "14dControl"
	RadarAs112TopLocationsParamsDateRange28dControl RadarAs112TopLocationsParamsDateRange = "28dControl"
	RadarAs112TopLocationsParamsDateRange12wControl RadarAs112TopLocationsParamsDateRange = "12wControl"
	RadarAs112TopLocationsParamsDateRange24wControl RadarAs112TopLocationsParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112TopLocationsParamsFormat string

const (
	RadarAs112TopLocationsParamsFormatJson RadarAs112TopLocationsParamsFormat = "JSON"
	RadarAs112TopLocationsParamsFormatCsv  RadarAs112TopLocationsParamsFormat = "CSV"
)

type RadarAs112TopLocationsResponseEnvelope struct {
	Result  RadarAs112TopLocationsResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarAs112TopLocationsResponseEnvelopeJSON `json:"-"`
}

// radarAs112TopLocationsResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarAs112TopLocationsResponseEnvelope]
type radarAs112TopLocationsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
