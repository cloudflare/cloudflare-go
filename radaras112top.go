// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// RadarAS112TopService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarAS112TopService] method
// instead.
type RadarAS112TopService struct {
	Options []option.RequestOption
}

// NewRadarAS112TopService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarAS112TopService(opts ...option.RequestOption) (r *RadarAS112TopService) {
	r = &RadarAS112TopService{}
	r.Options = opts
	return
}

// Get the top locations by DNS queries DNSSEC support to AS112.
func (r *RadarAS112TopService) DNSSEC(ctx context.Context, dnssec RadarAS112TopDNSSECParamsDNSSEC, query RadarAS112TopDNSSECParams, opts ...option.RequestOption) (res *RadarAS112TopDNSSECResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAS112TopDNSSECResponseEnvelope
	path := fmt.Sprintf("radar/as112/top/locations/dnssec/%v", dnssec)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the top locations, by DNS queries EDNS support to AS112.
func (r *RadarAS112TopService) Edns(ctx context.Context, edns RadarAS112TopEdnsParamsEdns, query RadarAS112TopEdnsParams, opts ...option.RequestOption) (res *RadarAS112TopEdnsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAS112TopEdnsResponseEnvelope
	path := fmt.Sprintf("radar/as112/top/locations/edns/%v", edns)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the top locations by DNS queries IP version to AS112.
func (r *RadarAS112TopService) IPVersion(ctx context.Context, ipVersion RadarAS112TopIPVersionParamsIPVersion, query RadarAS112TopIPVersionParams, opts ...option.RequestOption) (res *RadarAS112TopIPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAS112TopIPVersionResponseEnvelope
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
func (r *RadarAS112TopService) Locations(ctx context.Context, query RadarAS112TopLocationsParams, opts ...option.RequestOption) (res *RadarAS112TopLocationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAS112TopLocationsResponseEnvelope
	path := "radar/as112/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAS112TopDNSSECResponse struct {
	Meta RadarAS112TopDNSSECResponseMeta   `json:"meta,required"`
	Top0 []RadarAS112TopDNSSECResponseTop0 `json:"top_0,required"`
	JSON radarAS112TopDNSSECResponseJSON   `json:"-"`
}

// radarAS112TopDNSSECResponseJSON contains the JSON metadata for the struct
// [RadarAS112TopDNSSECResponse]
type radarAS112TopDNSSECResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TopDNSSECResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TopDNSSECResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAS112TopDNSSECResponseMeta struct {
	DateRange      []RadarAS112TopDNSSECResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                        `json:"lastUpdated,required"`
	ConfidenceInfo RadarAS112TopDNSSECResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAS112TopDNSSECResponseMetaJSON           `json:"-"`
}

// radarAS112TopDNSSECResponseMetaJSON contains the JSON metadata for the struct
// [RadarAS112TopDNSSECResponseMeta]
type radarAS112TopDNSSECResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAS112TopDNSSECResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TopDNSSECResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAS112TopDNSSECResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                    `json:"startTime,required" format:"date-time"`
	JSON      radarAS112TopDNSSECResponseMetaDateRangeJSON `json:"-"`
}

// radarAS112TopDNSSECResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [RadarAS112TopDNSSECResponseMetaDateRange]
type radarAS112TopDNSSECResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TopDNSSECResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TopDNSSECResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAS112TopDNSSECResponseMetaConfidenceInfo struct {
	Annotations []RadarAS112TopDNSSECResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                     `json:"level"`
	JSON        radarAS112TopDNSSECResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAS112TopDNSSECResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [RadarAS112TopDNSSECResponseMetaConfidenceInfo]
type radarAS112TopDNSSECResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TopDNSSECResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TopDNSSECResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAS112TopDNSSECResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                      `json:"dataSource,required"`
	Description     string                                                      `json:"description,required"`
	EventType       string                                                      `json:"eventType,required"`
	IsInstantaneous interface{}                                                 `json:"isInstantaneous,required"`
	EndTime         time.Time                                                   `json:"endTime" format:"date-time"`
	LinkedURL       string                                                      `json:"linkedUrl"`
	StartTime       time.Time                                                   `json:"startTime" format:"date-time"`
	JSON            radarAS112TopDNSSECResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAS112TopDNSSECResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [RadarAS112TopDNSSECResponseMetaConfidenceInfoAnnotation]
type radarAS112TopDNSSECResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAS112TopDNSSECResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TopDNSSECResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAS112TopDNSSECResponseTop0 struct {
	ClientCountryAlpha2 string                              `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                              `json:"clientCountryName,required"`
	Value               string                              `json:"value,required"`
	JSON                radarAS112TopDNSSECResponseTop0JSON `json:"-"`
}

// radarAS112TopDNSSECResponseTop0JSON contains the JSON metadata for the struct
// [RadarAS112TopDNSSECResponseTop0]
type radarAS112TopDNSSECResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAS112TopDNSSECResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TopDNSSECResponseTop0JSON) RawJSON() string {
	return r.raw
}

type RadarAS112TopEdnsResponse struct {
	Meta RadarAS112TopEdnsResponseMeta   `json:"meta,required"`
	Top0 []RadarAS112TopEdnsResponseTop0 `json:"top_0,required"`
	JSON radarAS112TopEdnsResponseJSON   `json:"-"`
}

// radarAS112TopEdnsResponseJSON contains the JSON metadata for the struct
// [RadarAS112TopEdnsResponse]
type radarAS112TopEdnsResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TopEdnsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TopEdnsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAS112TopEdnsResponseMeta struct {
	DateRange      []RadarAS112TopEdnsResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                      `json:"lastUpdated,required"`
	ConfidenceInfo RadarAS112TopEdnsResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAS112TopEdnsResponseMetaJSON           `json:"-"`
}

// radarAS112TopEdnsResponseMetaJSON contains the JSON metadata for the struct
// [RadarAS112TopEdnsResponseMeta]
type radarAS112TopEdnsResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAS112TopEdnsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TopEdnsResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAS112TopEdnsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                  `json:"startTime,required" format:"date-time"`
	JSON      radarAS112TopEdnsResponseMetaDateRangeJSON `json:"-"`
}

// radarAS112TopEdnsResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [RadarAS112TopEdnsResponseMetaDateRange]
type radarAS112TopEdnsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TopEdnsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TopEdnsResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAS112TopEdnsResponseMetaConfidenceInfo struct {
	Annotations []RadarAS112TopEdnsResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                   `json:"level"`
	JSON        radarAS112TopEdnsResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAS112TopEdnsResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [RadarAS112TopEdnsResponseMetaConfidenceInfo]
type radarAS112TopEdnsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TopEdnsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TopEdnsResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAS112TopEdnsResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                    `json:"dataSource,required"`
	Description     string                                                    `json:"description,required"`
	EventType       string                                                    `json:"eventType,required"`
	IsInstantaneous interface{}                                               `json:"isInstantaneous,required"`
	EndTime         time.Time                                                 `json:"endTime" format:"date-time"`
	LinkedURL       string                                                    `json:"linkedUrl"`
	StartTime       time.Time                                                 `json:"startTime" format:"date-time"`
	JSON            radarAS112TopEdnsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAS112TopEdnsResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [RadarAS112TopEdnsResponseMetaConfidenceInfoAnnotation]
type radarAS112TopEdnsResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAS112TopEdnsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TopEdnsResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAS112TopEdnsResponseTop0 struct {
	ClientCountryAlpha2 string                            `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                            `json:"clientCountryName,required"`
	Value               string                            `json:"value,required"`
	JSON                radarAS112TopEdnsResponseTop0JSON `json:"-"`
}

// radarAS112TopEdnsResponseTop0JSON contains the JSON metadata for the struct
// [RadarAS112TopEdnsResponseTop0]
type radarAS112TopEdnsResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAS112TopEdnsResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TopEdnsResponseTop0JSON) RawJSON() string {
	return r.raw
}

type RadarAS112TopIPVersionResponse struct {
	Meta RadarAS112TopIPVersionResponseMeta   `json:"meta,required"`
	Top0 []RadarAS112TopIPVersionResponseTop0 `json:"top_0,required"`
	JSON radarAS112TopIPVersionResponseJSON   `json:"-"`
}

// radarAS112TopIPVersionResponseJSON contains the JSON metadata for the struct
// [RadarAS112TopIPVersionResponse]
type radarAS112TopIPVersionResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TopIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TopIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAS112TopIPVersionResponseMeta struct {
	DateRange      []RadarAS112TopIPVersionResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                           `json:"lastUpdated,required"`
	ConfidenceInfo RadarAS112TopIPVersionResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAS112TopIPVersionResponseMetaJSON           `json:"-"`
}

// radarAS112TopIPVersionResponseMetaJSON contains the JSON metadata for the struct
// [RadarAS112TopIPVersionResponseMeta]
type radarAS112TopIPVersionResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAS112TopIPVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TopIPVersionResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAS112TopIPVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      radarAS112TopIPVersionResponseMetaDateRangeJSON `json:"-"`
}

// radarAS112TopIPVersionResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [RadarAS112TopIPVersionResponseMetaDateRange]
type radarAS112TopIPVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TopIPVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TopIPVersionResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAS112TopIPVersionResponseMetaConfidenceInfo struct {
	Annotations []RadarAS112TopIPVersionResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                        `json:"level"`
	JSON        radarAS112TopIPVersionResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAS112TopIPVersionResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [RadarAS112TopIPVersionResponseMetaConfidenceInfo]
type radarAS112TopIPVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TopIPVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TopIPVersionResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAS112TopIPVersionResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                         `json:"dataSource,required"`
	Description     string                                                         `json:"description,required"`
	EventType       string                                                         `json:"eventType,required"`
	IsInstantaneous interface{}                                                    `json:"isInstantaneous,required"`
	EndTime         time.Time                                                      `json:"endTime" format:"date-time"`
	LinkedURL       string                                                         `json:"linkedUrl"`
	StartTime       time.Time                                                      `json:"startTime" format:"date-time"`
	JSON            radarAS112TopIPVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAS112TopIPVersionResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [RadarAS112TopIPVersionResponseMetaConfidenceInfoAnnotation]
type radarAS112TopIPVersionResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAS112TopIPVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TopIPVersionResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAS112TopIPVersionResponseTop0 struct {
	ClientCountryAlpha2 string                                 `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                 `json:"clientCountryName,required"`
	Value               string                                 `json:"value,required"`
	JSON                radarAS112TopIPVersionResponseTop0JSON `json:"-"`
}

// radarAS112TopIPVersionResponseTop0JSON contains the JSON metadata for the struct
// [RadarAS112TopIPVersionResponseTop0]
type radarAS112TopIPVersionResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAS112TopIPVersionResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TopIPVersionResponseTop0JSON) RawJSON() string {
	return r.raw
}

type RadarAS112TopLocationsResponse struct {
	Meta RadarAS112TopLocationsResponseMeta   `json:"meta,required"`
	Top0 []RadarAS112TopLocationsResponseTop0 `json:"top_0,required"`
	JSON radarAS112TopLocationsResponseJSON   `json:"-"`
}

// radarAS112TopLocationsResponseJSON contains the JSON metadata for the struct
// [RadarAS112TopLocationsResponse]
type radarAS112TopLocationsResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TopLocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TopLocationsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAS112TopLocationsResponseMeta struct {
	DateRange      []RadarAS112TopLocationsResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                           `json:"lastUpdated,required"`
	ConfidenceInfo RadarAS112TopLocationsResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAS112TopLocationsResponseMetaJSON           `json:"-"`
}

// radarAS112TopLocationsResponseMetaJSON contains the JSON metadata for the struct
// [RadarAS112TopLocationsResponseMeta]
type radarAS112TopLocationsResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAS112TopLocationsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TopLocationsResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAS112TopLocationsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      radarAS112TopLocationsResponseMetaDateRangeJSON `json:"-"`
}

// radarAS112TopLocationsResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [RadarAS112TopLocationsResponseMetaDateRange]
type radarAS112TopLocationsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TopLocationsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TopLocationsResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAS112TopLocationsResponseMetaConfidenceInfo struct {
	Annotations []RadarAS112TopLocationsResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                        `json:"level"`
	JSON        radarAS112TopLocationsResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAS112TopLocationsResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [RadarAS112TopLocationsResponseMetaConfidenceInfo]
type radarAS112TopLocationsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TopLocationsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TopLocationsResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAS112TopLocationsResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                         `json:"dataSource,required"`
	Description     string                                                         `json:"description,required"`
	EventType       string                                                         `json:"eventType,required"`
	IsInstantaneous interface{}                                                    `json:"isInstantaneous,required"`
	EndTime         time.Time                                                      `json:"endTime" format:"date-time"`
	LinkedURL       string                                                         `json:"linkedUrl"`
	StartTime       time.Time                                                      `json:"startTime" format:"date-time"`
	JSON            radarAS112TopLocationsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAS112TopLocationsResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [RadarAS112TopLocationsResponseMetaConfidenceInfoAnnotation]
type radarAS112TopLocationsResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAS112TopLocationsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TopLocationsResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAS112TopLocationsResponseTop0 struct {
	ClientCountryAlpha2 string                                 `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                 `json:"clientCountryName,required"`
	Value               string                                 `json:"value,required"`
	JSON                radarAS112TopLocationsResponseTop0JSON `json:"-"`
}

// radarAS112TopLocationsResponseTop0JSON contains the JSON metadata for the struct
// [RadarAS112TopLocationsResponseTop0]
type radarAS112TopLocationsResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAS112TopLocationsResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TopLocationsResponseTop0JSON) RawJSON() string {
	return r.raw
}

type RadarAS112TopDNSSECParams struct {
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
	DateRange param.Field[[]RadarAS112TopDNSSECParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAS112TopDNSSECParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAS112TopDNSSECParams]'s query parameters as
// `url.Values`.
func (r RadarAS112TopDNSSECParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// DNSSEC.
type RadarAS112TopDNSSECParamsDNSSEC string

const (
	RadarAS112TopDNSSECParamsDNSSECSupported    RadarAS112TopDNSSECParamsDNSSEC = "SUPPORTED"
	RadarAS112TopDNSSECParamsDNSSECNotSupported RadarAS112TopDNSSECParamsDNSSEC = "NOT_SUPPORTED"
)

type RadarAS112TopDNSSECParamsDateRange string

const (
	RadarAS112TopDNSSECParamsDateRange1d         RadarAS112TopDNSSECParamsDateRange = "1d"
	RadarAS112TopDNSSECParamsDateRange2d         RadarAS112TopDNSSECParamsDateRange = "2d"
	RadarAS112TopDNSSECParamsDateRange7d         RadarAS112TopDNSSECParamsDateRange = "7d"
	RadarAS112TopDNSSECParamsDateRange14d        RadarAS112TopDNSSECParamsDateRange = "14d"
	RadarAS112TopDNSSECParamsDateRange28d        RadarAS112TopDNSSECParamsDateRange = "28d"
	RadarAS112TopDNSSECParamsDateRange12w        RadarAS112TopDNSSECParamsDateRange = "12w"
	RadarAS112TopDNSSECParamsDateRange24w        RadarAS112TopDNSSECParamsDateRange = "24w"
	RadarAS112TopDNSSECParamsDateRange52w        RadarAS112TopDNSSECParamsDateRange = "52w"
	RadarAS112TopDNSSECParamsDateRange1dControl  RadarAS112TopDNSSECParamsDateRange = "1dControl"
	RadarAS112TopDNSSECParamsDateRange2dControl  RadarAS112TopDNSSECParamsDateRange = "2dControl"
	RadarAS112TopDNSSECParamsDateRange7dControl  RadarAS112TopDNSSECParamsDateRange = "7dControl"
	RadarAS112TopDNSSECParamsDateRange14dControl RadarAS112TopDNSSECParamsDateRange = "14dControl"
	RadarAS112TopDNSSECParamsDateRange28dControl RadarAS112TopDNSSECParamsDateRange = "28dControl"
	RadarAS112TopDNSSECParamsDateRange12wControl RadarAS112TopDNSSECParamsDateRange = "12wControl"
	RadarAS112TopDNSSECParamsDateRange24wControl RadarAS112TopDNSSECParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAS112TopDNSSECParamsFormat string

const (
	RadarAS112TopDNSSECParamsFormatJson RadarAS112TopDNSSECParamsFormat = "JSON"
	RadarAS112TopDNSSECParamsFormatCsv  RadarAS112TopDNSSECParamsFormat = "CSV"
)

type RadarAS112TopDNSSECResponseEnvelope struct {
	Result  RadarAS112TopDNSSECResponse             `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    radarAS112TopDNSSECResponseEnvelopeJSON `json:"-"`
}

// radarAS112TopDNSSECResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarAS112TopDNSSECResponseEnvelope]
type radarAS112TopDNSSECResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TopDNSSECResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TopDNSSECResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarAS112TopEdnsParams struct {
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
	DateRange param.Field[[]RadarAS112TopEdnsParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAS112TopEdnsParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAS112TopEdnsParams]'s query parameters as
// `url.Values`.
func (r RadarAS112TopEdnsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// EDNS.
type RadarAS112TopEdnsParamsEdns string

const (
	RadarAS112TopEdnsParamsEdnsSupported    RadarAS112TopEdnsParamsEdns = "SUPPORTED"
	RadarAS112TopEdnsParamsEdnsNotSupported RadarAS112TopEdnsParamsEdns = "NOT_SUPPORTED"
)

type RadarAS112TopEdnsParamsDateRange string

const (
	RadarAS112TopEdnsParamsDateRange1d         RadarAS112TopEdnsParamsDateRange = "1d"
	RadarAS112TopEdnsParamsDateRange2d         RadarAS112TopEdnsParamsDateRange = "2d"
	RadarAS112TopEdnsParamsDateRange7d         RadarAS112TopEdnsParamsDateRange = "7d"
	RadarAS112TopEdnsParamsDateRange14d        RadarAS112TopEdnsParamsDateRange = "14d"
	RadarAS112TopEdnsParamsDateRange28d        RadarAS112TopEdnsParamsDateRange = "28d"
	RadarAS112TopEdnsParamsDateRange12w        RadarAS112TopEdnsParamsDateRange = "12w"
	RadarAS112TopEdnsParamsDateRange24w        RadarAS112TopEdnsParamsDateRange = "24w"
	RadarAS112TopEdnsParamsDateRange52w        RadarAS112TopEdnsParamsDateRange = "52w"
	RadarAS112TopEdnsParamsDateRange1dControl  RadarAS112TopEdnsParamsDateRange = "1dControl"
	RadarAS112TopEdnsParamsDateRange2dControl  RadarAS112TopEdnsParamsDateRange = "2dControl"
	RadarAS112TopEdnsParamsDateRange7dControl  RadarAS112TopEdnsParamsDateRange = "7dControl"
	RadarAS112TopEdnsParamsDateRange14dControl RadarAS112TopEdnsParamsDateRange = "14dControl"
	RadarAS112TopEdnsParamsDateRange28dControl RadarAS112TopEdnsParamsDateRange = "28dControl"
	RadarAS112TopEdnsParamsDateRange12wControl RadarAS112TopEdnsParamsDateRange = "12wControl"
	RadarAS112TopEdnsParamsDateRange24wControl RadarAS112TopEdnsParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAS112TopEdnsParamsFormat string

const (
	RadarAS112TopEdnsParamsFormatJson RadarAS112TopEdnsParamsFormat = "JSON"
	RadarAS112TopEdnsParamsFormatCsv  RadarAS112TopEdnsParamsFormat = "CSV"
)

type RadarAS112TopEdnsResponseEnvelope struct {
	Result  RadarAS112TopEdnsResponse             `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    radarAS112TopEdnsResponseEnvelopeJSON `json:"-"`
}

// radarAS112TopEdnsResponseEnvelopeJSON contains the JSON metadata for the struct
// [RadarAS112TopEdnsResponseEnvelope]
type radarAS112TopEdnsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TopEdnsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TopEdnsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarAS112TopIPVersionParams struct {
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
	DateRange param.Field[[]RadarAS112TopIPVersionParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAS112TopIPVersionParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAS112TopIPVersionParams]'s query parameters as
// `url.Values`.
func (r RadarAS112TopIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// IP Version.
type RadarAS112TopIPVersionParamsIPVersion string

const (
	RadarAS112TopIPVersionParamsIPVersionIPv4 RadarAS112TopIPVersionParamsIPVersion = "IPv4"
	RadarAS112TopIPVersionParamsIPVersionIPv6 RadarAS112TopIPVersionParamsIPVersion = "IPv6"
)

type RadarAS112TopIPVersionParamsDateRange string

const (
	RadarAS112TopIPVersionParamsDateRange1d         RadarAS112TopIPVersionParamsDateRange = "1d"
	RadarAS112TopIPVersionParamsDateRange2d         RadarAS112TopIPVersionParamsDateRange = "2d"
	RadarAS112TopIPVersionParamsDateRange7d         RadarAS112TopIPVersionParamsDateRange = "7d"
	RadarAS112TopIPVersionParamsDateRange14d        RadarAS112TopIPVersionParamsDateRange = "14d"
	RadarAS112TopIPVersionParamsDateRange28d        RadarAS112TopIPVersionParamsDateRange = "28d"
	RadarAS112TopIPVersionParamsDateRange12w        RadarAS112TopIPVersionParamsDateRange = "12w"
	RadarAS112TopIPVersionParamsDateRange24w        RadarAS112TopIPVersionParamsDateRange = "24w"
	RadarAS112TopIPVersionParamsDateRange52w        RadarAS112TopIPVersionParamsDateRange = "52w"
	RadarAS112TopIPVersionParamsDateRange1dControl  RadarAS112TopIPVersionParamsDateRange = "1dControl"
	RadarAS112TopIPVersionParamsDateRange2dControl  RadarAS112TopIPVersionParamsDateRange = "2dControl"
	RadarAS112TopIPVersionParamsDateRange7dControl  RadarAS112TopIPVersionParamsDateRange = "7dControl"
	RadarAS112TopIPVersionParamsDateRange14dControl RadarAS112TopIPVersionParamsDateRange = "14dControl"
	RadarAS112TopIPVersionParamsDateRange28dControl RadarAS112TopIPVersionParamsDateRange = "28dControl"
	RadarAS112TopIPVersionParamsDateRange12wControl RadarAS112TopIPVersionParamsDateRange = "12wControl"
	RadarAS112TopIPVersionParamsDateRange24wControl RadarAS112TopIPVersionParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAS112TopIPVersionParamsFormat string

const (
	RadarAS112TopIPVersionParamsFormatJson RadarAS112TopIPVersionParamsFormat = "JSON"
	RadarAS112TopIPVersionParamsFormatCsv  RadarAS112TopIPVersionParamsFormat = "CSV"
)

type RadarAS112TopIPVersionResponseEnvelope struct {
	Result  RadarAS112TopIPVersionResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarAS112TopIPVersionResponseEnvelopeJSON `json:"-"`
}

// radarAS112TopIPVersionResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarAS112TopIPVersionResponseEnvelope]
type radarAS112TopIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TopIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TopIPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarAS112TopLocationsParams struct {
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
	DateRange param.Field[[]RadarAS112TopLocationsParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAS112TopLocationsParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAS112TopLocationsParams]'s query parameters as
// `url.Values`.
func (r RadarAS112TopLocationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAS112TopLocationsParamsDateRange string

const (
	RadarAS112TopLocationsParamsDateRange1d         RadarAS112TopLocationsParamsDateRange = "1d"
	RadarAS112TopLocationsParamsDateRange2d         RadarAS112TopLocationsParamsDateRange = "2d"
	RadarAS112TopLocationsParamsDateRange7d         RadarAS112TopLocationsParamsDateRange = "7d"
	RadarAS112TopLocationsParamsDateRange14d        RadarAS112TopLocationsParamsDateRange = "14d"
	RadarAS112TopLocationsParamsDateRange28d        RadarAS112TopLocationsParamsDateRange = "28d"
	RadarAS112TopLocationsParamsDateRange12w        RadarAS112TopLocationsParamsDateRange = "12w"
	RadarAS112TopLocationsParamsDateRange24w        RadarAS112TopLocationsParamsDateRange = "24w"
	RadarAS112TopLocationsParamsDateRange52w        RadarAS112TopLocationsParamsDateRange = "52w"
	RadarAS112TopLocationsParamsDateRange1dControl  RadarAS112TopLocationsParamsDateRange = "1dControl"
	RadarAS112TopLocationsParamsDateRange2dControl  RadarAS112TopLocationsParamsDateRange = "2dControl"
	RadarAS112TopLocationsParamsDateRange7dControl  RadarAS112TopLocationsParamsDateRange = "7dControl"
	RadarAS112TopLocationsParamsDateRange14dControl RadarAS112TopLocationsParamsDateRange = "14dControl"
	RadarAS112TopLocationsParamsDateRange28dControl RadarAS112TopLocationsParamsDateRange = "28dControl"
	RadarAS112TopLocationsParamsDateRange12wControl RadarAS112TopLocationsParamsDateRange = "12wControl"
	RadarAS112TopLocationsParamsDateRange24wControl RadarAS112TopLocationsParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAS112TopLocationsParamsFormat string

const (
	RadarAS112TopLocationsParamsFormatJson RadarAS112TopLocationsParamsFormat = "JSON"
	RadarAS112TopLocationsParamsFormatCsv  RadarAS112TopLocationsParamsFormat = "CSV"
)

type RadarAS112TopLocationsResponseEnvelope struct {
	Result  RadarAS112TopLocationsResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarAS112TopLocationsResponseEnvelopeJSON `json:"-"`
}

// radarAS112TopLocationsResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarAS112TopLocationsResponseEnvelope]
type radarAS112TopLocationsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TopLocationsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TopLocationsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
