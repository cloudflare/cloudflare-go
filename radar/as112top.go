// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// AS112TopService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAS112TopService] method instead.
type AS112TopService struct {
	Options []option.RequestOption
}

// NewAS112TopService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAS112TopService(opts ...option.RequestOption) (r *AS112TopService) {
	r = &AS112TopService{}
	r.Options = opts
	return
}

// Get the top locations by DNS queries DNSSEC support to AS112.
func (r *AS112TopService) DNSSEC(ctx context.Context, dnssec AS112TopDNSSECParamsDNSSEC, query AS112TopDNSSECParams, opts ...option.RequestOption) (res *AS112TopDNSSECResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AS112TopDNSSECResponseEnvelope
	path := fmt.Sprintf("radar/as112/top/locations/dnssec/%v", dnssec)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the top locations, by DNS queries EDNS support to AS112.
func (r *AS112TopService) Edns(ctx context.Context, edns AS112TopEdnsParamsEdns, query AS112TopEdnsParams, opts ...option.RequestOption) (res *AS112TopEdnsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AS112TopEdnsResponseEnvelope
	path := fmt.Sprintf("radar/as112/top/locations/edns/%v", edns)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the top locations by DNS queries IP version to AS112.
func (r *AS112TopService) IPVersion(ctx context.Context, ipVersion AS112TopIPVersionParamsIPVersion, query AS112TopIPVersionParams, opts ...option.RequestOption) (res *AS112TopIPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AS112TopIPVersionResponseEnvelope
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
func (r *AS112TopService) Locations(ctx context.Context, query AS112TopLocationsParams, opts ...option.RequestOption) (res *AS112TopLocationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AS112TopLocationsResponseEnvelope
	path := "radar/as112/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AS112TopDNSSECResponse struct {
	Meta AS112TopDNSSECResponseMeta   `json:"meta,required"`
	Top0 []AS112TopDNSSECResponseTop0 `json:"top_0,required"`
	JSON as112TopDNSSECResponseJSON   `json:"-"`
}

// as112TopDNSSECResponseJSON contains the JSON metadata for the struct
// [AS112TopDNSSECResponse]
type as112TopDNSSECResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopDNSSECResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopDNSSECResponseJSON) RawJSON() string {
	return r.raw
}

type AS112TopDNSSECResponseMeta struct {
	DateRange      []UnnamedSchemaRef175                    `json:"dateRange,required"`
	LastUpdated    string                                   `json:"lastUpdated,required"`
	ConfidenceInfo AS112TopDNSSECResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           as112TopDNSSECResponseMetaJSON           `json:"-"`
}

// as112TopDNSSECResponseMetaJSON contains the JSON metadata for the struct
// [AS112TopDNSSECResponseMeta]
type as112TopDNSSECResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AS112TopDNSSECResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopDNSSECResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AS112TopDNSSECResponseMetaConfidenceInfo struct {
	Annotations []UnnamedSchemaRef174                        `json:"annotations"`
	Level       int64                                        `json:"level"`
	JSON        as112TopDNSSECResponseMetaConfidenceInfoJSON `json:"-"`
}

// as112TopDNSSECResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [AS112TopDNSSECResponseMetaConfidenceInfo]
type as112TopDNSSECResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopDNSSECResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopDNSSECResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AS112TopDNSSECResponseTop0 struct {
	ClientCountryAlpha2 string                         `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                         `json:"clientCountryName,required"`
	Value               string                         `json:"value,required"`
	JSON                as112TopDNSSECResponseTop0JSON `json:"-"`
}

// as112TopDNSSECResponseTop0JSON contains the JSON metadata for the struct
// [AS112TopDNSSECResponseTop0]
type as112TopDNSSECResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AS112TopDNSSECResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopDNSSECResponseTop0JSON) RawJSON() string {
	return r.raw
}

type AS112TopEdnsResponse struct {
	Meta AS112TopEdnsResponseMeta   `json:"meta,required"`
	Top0 []AS112TopEdnsResponseTop0 `json:"top_0,required"`
	JSON as112TopEdnsResponseJSON   `json:"-"`
}

// as112TopEdnsResponseJSON contains the JSON metadata for the struct
// [AS112TopEdnsResponse]
type as112TopEdnsResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopEdnsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopEdnsResponseJSON) RawJSON() string {
	return r.raw
}

type AS112TopEdnsResponseMeta struct {
	DateRange      []UnnamedSchemaRef175                  `json:"dateRange,required"`
	LastUpdated    string                                 `json:"lastUpdated,required"`
	ConfidenceInfo AS112TopEdnsResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           as112TopEdnsResponseMetaJSON           `json:"-"`
}

// as112TopEdnsResponseMetaJSON contains the JSON metadata for the struct
// [AS112TopEdnsResponseMeta]
type as112TopEdnsResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AS112TopEdnsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopEdnsResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AS112TopEdnsResponseMetaConfidenceInfo struct {
	Annotations []UnnamedSchemaRef174                      `json:"annotations"`
	Level       int64                                      `json:"level"`
	JSON        as112TopEdnsResponseMetaConfidenceInfoJSON `json:"-"`
}

// as112TopEdnsResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [AS112TopEdnsResponseMetaConfidenceInfo]
type as112TopEdnsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopEdnsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopEdnsResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AS112TopEdnsResponseTop0 struct {
	ClientCountryAlpha2 string                       `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                       `json:"clientCountryName,required"`
	Value               string                       `json:"value,required"`
	JSON                as112TopEdnsResponseTop0JSON `json:"-"`
}

// as112TopEdnsResponseTop0JSON contains the JSON metadata for the struct
// [AS112TopEdnsResponseTop0]
type as112TopEdnsResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AS112TopEdnsResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopEdnsResponseTop0JSON) RawJSON() string {
	return r.raw
}

type AS112TopIPVersionResponse struct {
	Meta AS112TopIPVersionResponseMeta   `json:"meta,required"`
	Top0 []AS112TopIPVersionResponseTop0 `json:"top_0,required"`
	JSON as112TopIPVersionResponseJSON   `json:"-"`
}

// as112TopIPVersionResponseJSON contains the JSON metadata for the struct
// [AS112TopIPVersionResponse]
type as112TopIPVersionResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type AS112TopIPVersionResponseMeta struct {
	DateRange      []UnnamedSchemaRef175                       `json:"dateRange,required"`
	LastUpdated    string                                      `json:"lastUpdated,required"`
	ConfidenceInfo AS112TopIPVersionResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           as112TopIPVersionResponseMetaJSON           `json:"-"`
}

// as112TopIPVersionResponseMetaJSON contains the JSON metadata for the struct
// [AS112TopIPVersionResponseMeta]
type as112TopIPVersionResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AS112TopIPVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopIPVersionResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AS112TopIPVersionResponseMetaConfidenceInfo struct {
	Annotations []UnnamedSchemaRef174                           `json:"annotations"`
	Level       int64                                           `json:"level"`
	JSON        as112TopIPVersionResponseMetaConfidenceInfoJSON `json:"-"`
}

// as112TopIPVersionResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [AS112TopIPVersionResponseMetaConfidenceInfo]
type as112TopIPVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopIPVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopIPVersionResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AS112TopIPVersionResponseTop0 struct {
	ClientCountryAlpha2 string                            `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                            `json:"clientCountryName,required"`
	Value               string                            `json:"value,required"`
	JSON                as112TopIPVersionResponseTop0JSON `json:"-"`
}

// as112TopIPVersionResponseTop0JSON contains the JSON metadata for the struct
// [AS112TopIPVersionResponseTop0]
type as112TopIPVersionResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AS112TopIPVersionResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopIPVersionResponseTop0JSON) RawJSON() string {
	return r.raw
}

type AS112TopLocationsResponse struct {
	Meta AS112TopLocationsResponseMeta   `json:"meta,required"`
	Top0 []AS112TopLocationsResponseTop0 `json:"top_0,required"`
	JSON as112TopLocationsResponseJSON   `json:"-"`
}

// as112TopLocationsResponseJSON contains the JSON metadata for the struct
// [AS112TopLocationsResponse]
type as112TopLocationsResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopLocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopLocationsResponseJSON) RawJSON() string {
	return r.raw
}

type AS112TopLocationsResponseMeta struct {
	DateRange      []UnnamedSchemaRef175                       `json:"dateRange,required"`
	LastUpdated    string                                      `json:"lastUpdated,required"`
	ConfidenceInfo AS112TopLocationsResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           as112TopLocationsResponseMetaJSON           `json:"-"`
}

// as112TopLocationsResponseMetaJSON contains the JSON metadata for the struct
// [AS112TopLocationsResponseMeta]
type as112TopLocationsResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AS112TopLocationsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopLocationsResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AS112TopLocationsResponseMetaConfidenceInfo struct {
	Annotations []UnnamedSchemaRef174                           `json:"annotations"`
	Level       int64                                           `json:"level"`
	JSON        as112TopLocationsResponseMetaConfidenceInfoJSON `json:"-"`
}

// as112TopLocationsResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [AS112TopLocationsResponseMetaConfidenceInfo]
type as112TopLocationsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopLocationsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopLocationsResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AS112TopLocationsResponseTop0 struct {
	ClientCountryAlpha2 string                            `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                            `json:"clientCountryName,required"`
	Value               string                            `json:"value,required"`
	JSON                as112TopLocationsResponseTop0JSON `json:"-"`
}

// as112TopLocationsResponseTop0JSON contains the JSON metadata for the struct
// [AS112TopLocationsResponseTop0]
type as112TopLocationsResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AS112TopLocationsResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopLocationsResponseTop0JSON) RawJSON() string {
	return r.raw
}

type AS112TopDNSSECParams struct {
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
	DateRange param.Field[[]AS112TopDNSSECParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AS112TopDNSSECParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AS112TopDNSSECParams]'s query parameters as `url.Values`.
func (r AS112TopDNSSECParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// DNSSEC.
type AS112TopDNSSECParamsDNSSEC string

const (
	AS112TopDNSSECParamsDNSSECSupported    AS112TopDNSSECParamsDNSSEC = "SUPPORTED"
	AS112TopDNSSECParamsDNSSECNotSupported AS112TopDNSSECParamsDNSSEC = "NOT_SUPPORTED"
)

func (r AS112TopDNSSECParamsDNSSEC) IsKnown() bool {
	switch r {
	case AS112TopDNSSECParamsDNSSECSupported, AS112TopDNSSECParamsDNSSECNotSupported:
		return true
	}
	return false
}

type AS112TopDNSSECParamsDateRange string

const (
	AS112TopDNSSECParamsDateRange1d         AS112TopDNSSECParamsDateRange = "1d"
	AS112TopDNSSECParamsDateRange2d         AS112TopDNSSECParamsDateRange = "2d"
	AS112TopDNSSECParamsDateRange7d         AS112TopDNSSECParamsDateRange = "7d"
	AS112TopDNSSECParamsDateRange14d        AS112TopDNSSECParamsDateRange = "14d"
	AS112TopDNSSECParamsDateRange28d        AS112TopDNSSECParamsDateRange = "28d"
	AS112TopDNSSECParamsDateRange12w        AS112TopDNSSECParamsDateRange = "12w"
	AS112TopDNSSECParamsDateRange24w        AS112TopDNSSECParamsDateRange = "24w"
	AS112TopDNSSECParamsDateRange52w        AS112TopDNSSECParamsDateRange = "52w"
	AS112TopDNSSECParamsDateRange1dControl  AS112TopDNSSECParamsDateRange = "1dControl"
	AS112TopDNSSECParamsDateRange2dControl  AS112TopDNSSECParamsDateRange = "2dControl"
	AS112TopDNSSECParamsDateRange7dControl  AS112TopDNSSECParamsDateRange = "7dControl"
	AS112TopDNSSECParamsDateRange14dControl AS112TopDNSSECParamsDateRange = "14dControl"
	AS112TopDNSSECParamsDateRange28dControl AS112TopDNSSECParamsDateRange = "28dControl"
	AS112TopDNSSECParamsDateRange12wControl AS112TopDNSSECParamsDateRange = "12wControl"
	AS112TopDNSSECParamsDateRange24wControl AS112TopDNSSECParamsDateRange = "24wControl"
)

func (r AS112TopDNSSECParamsDateRange) IsKnown() bool {
	switch r {
	case AS112TopDNSSECParamsDateRange1d, AS112TopDNSSECParamsDateRange2d, AS112TopDNSSECParamsDateRange7d, AS112TopDNSSECParamsDateRange14d, AS112TopDNSSECParamsDateRange28d, AS112TopDNSSECParamsDateRange12w, AS112TopDNSSECParamsDateRange24w, AS112TopDNSSECParamsDateRange52w, AS112TopDNSSECParamsDateRange1dControl, AS112TopDNSSECParamsDateRange2dControl, AS112TopDNSSECParamsDateRange7dControl, AS112TopDNSSECParamsDateRange14dControl, AS112TopDNSSECParamsDateRange28dControl, AS112TopDNSSECParamsDateRange12wControl, AS112TopDNSSECParamsDateRange24wControl:
		return true
	}
	return false
}

// Format results are returned in.
type AS112TopDNSSECParamsFormat string

const (
	AS112TopDNSSECParamsFormatJson AS112TopDNSSECParamsFormat = "JSON"
	AS112TopDNSSECParamsFormatCsv  AS112TopDNSSECParamsFormat = "CSV"
)

func (r AS112TopDNSSECParamsFormat) IsKnown() bool {
	switch r {
	case AS112TopDNSSECParamsFormatJson, AS112TopDNSSECParamsFormatCsv:
		return true
	}
	return false
}

type AS112TopDNSSECResponseEnvelope struct {
	Result  AS112TopDNSSECResponse             `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    as112TopDNSSECResponseEnvelopeJSON `json:"-"`
}

// as112TopDNSSECResponseEnvelopeJSON contains the JSON metadata for the struct
// [AS112TopDNSSECResponseEnvelope]
type as112TopDNSSECResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopDNSSECResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopDNSSECResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AS112TopEdnsParams struct {
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
	DateRange param.Field[[]AS112TopEdnsParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AS112TopEdnsParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AS112TopEdnsParams]'s query parameters as `url.Values`.
func (r AS112TopEdnsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// EDNS.
type AS112TopEdnsParamsEdns string

const (
	AS112TopEdnsParamsEdnsSupported    AS112TopEdnsParamsEdns = "SUPPORTED"
	AS112TopEdnsParamsEdnsNotSupported AS112TopEdnsParamsEdns = "NOT_SUPPORTED"
)

func (r AS112TopEdnsParamsEdns) IsKnown() bool {
	switch r {
	case AS112TopEdnsParamsEdnsSupported, AS112TopEdnsParamsEdnsNotSupported:
		return true
	}
	return false
}

type AS112TopEdnsParamsDateRange string

const (
	AS112TopEdnsParamsDateRange1d         AS112TopEdnsParamsDateRange = "1d"
	AS112TopEdnsParamsDateRange2d         AS112TopEdnsParamsDateRange = "2d"
	AS112TopEdnsParamsDateRange7d         AS112TopEdnsParamsDateRange = "7d"
	AS112TopEdnsParamsDateRange14d        AS112TopEdnsParamsDateRange = "14d"
	AS112TopEdnsParamsDateRange28d        AS112TopEdnsParamsDateRange = "28d"
	AS112TopEdnsParamsDateRange12w        AS112TopEdnsParamsDateRange = "12w"
	AS112TopEdnsParamsDateRange24w        AS112TopEdnsParamsDateRange = "24w"
	AS112TopEdnsParamsDateRange52w        AS112TopEdnsParamsDateRange = "52w"
	AS112TopEdnsParamsDateRange1dControl  AS112TopEdnsParamsDateRange = "1dControl"
	AS112TopEdnsParamsDateRange2dControl  AS112TopEdnsParamsDateRange = "2dControl"
	AS112TopEdnsParamsDateRange7dControl  AS112TopEdnsParamsDateRange = "7dControl"
	AS112TopEdnsParamsDateRange14dControl AS112TopEdnsParamsDateRange = "14dControl"
	AS112TopEdnsParamsDateRange28dControl AS112TopEdnsParamsDateRange = "28dControl"
	AS112TopEdnsParamsDateRange12wControl AS112TopEdnsParamsDateRange = "12wControl"
	AS112TopEdnsParamsDateRange24wControl AS112TopEdnsParamsDateRange = "24wControl"
)

func (r AS112TopEdnsParamsDateRange) IsKnown() bool {
	switch r {
	case AS112TopEdnsParamsDateRange1d, AS112TopEdnsParamsDateRange2d, AS112TopEdnsParamsDateRange7d, AS112TopEdnsParamsDateRange14d, AS112TopEdnsParamsDateRange28d, AS112TopEdnsParamsDateRange12w, AS112TopEdnsParamsDateRange24w, AS112TopEdnsParamsDateRange52w, AS112TopEdnsParamsDateRange1dControl, AS112TopEdnsParamsDateRange2dControl, AS112TopEdnsParamsDateRange7dControl, AS112TopEdnsParamsDateRange14dControl, AS112TopEdnsParamsDateRange28dControl, AS112TopEdnsParamsDateRange12wControl, AS112TopEdnsParamsDateRange24wControl:
		return true
	}
	return false
}

// Format results are returned in.
type AS112TopEdnsParamsFormat string

const (
	AS112TopEdnsParamsFormatJson AS112TopEdnsParamsFormat = "JSON"
	AS112TopEdnsParamsFormatCsv  AS112TopEdnsParamsFormat = "CSV"
)

func (r AS112TopEdnsParamsFormat) IsKnown() bool {
	switch r {
	case AS112TopEdnsParamsFormatJson, AS112TopEdnsParamsFormatCsv:
		return true
	}
	return false
}

type AS112TopEdnsResponseEnvelope struct {
	Result  AS112TopEdnsResponse             `json:"result,required"`
	Success bool                             `json:"success,required"`
	JSON    as112TopEdnsResponseEnvelopeJSON `json:"-"`
}

// as112TopEdnsResponseEnvelopeJSON contains the JSON metadata for the struct
// [AS112TopEdnsResponseEnvelope]
type as112TopEdnsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopEdnsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopEdnsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AS112TopIPVersionParams struct {
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
	DateRange param.Field[[]AS112TopIPVersionParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AS112TopIPVersionParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AS112TopIPVersionParams]'s query parameters as
// `url.Values`.
func (r AS112TopIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// IP Version.
type AS112TopIPVersionParamsIPVersion string

const (
	AS112TopIPVersionParamsIPVersionIPv4 AS112TopIPVersionParamsIPVersion = "IPv4"
	AS112TopIPVersionParamsIPVersionIPv6 AS112TopIPVersionParamsIPVersion = "IPv6"
)

func (r AS112TopIPVersionParamsIPVersion) IsKnown() bool {
	switch r {
	case AS112TopIPVersionParamsIPVersionIPv4, AS112TopIPVersionParamsIPVersionIPv6:
		return true
	}
	return false
}

type AS112TopIPVersionParamsDateRange string

const (
	AS112TopIPVersionParamsDateRange1d         AS112TopIPVersionParamsDateRange = "1d"
	AS112TopIPVersionParamsDateRange2d         AS112TopIPVersionParamsDateRange = "2d"
	AS112TopIPVersionParamsDateRange7d         AS112TopIPVersionParamsDateRange = "7d"
	AS112TopIPVersionParamsDateRange14d        AS112TopIPVersionParamsDateRange = "14d"
	AS112TopIPVersionParamsDateRange28d        AS112TopIPVersionParamsDateRange = "28d"
	AS112TopIPVersionParamsDateRange12w        AS112TopIPVersionParamsDateRange = "12w"
	AS112TopIPVersionParamsDateRange24w        AS112TopIPVersionParamsDateRange = "24w"
	AS112TopIPVersionParamsDateRange52w        AS112TopIPVersionParamsDateRange = "52w"
	AS112TopIPVersionParamsDateRange1dControl  AS112TopIPVersionParamsDateRange = "1dControl"
	AS112TopIPVersionParamsDateRange2dControl  AS112TopIPVersionParamsDateRange = "2dControl"
	AS112TopIPVersionParamsDateRange7dControl  AS112TopIPVersionParamsDateRange = "7dControl"
	AS112TopIPVersionParamsDateRange14dControl AS112TopIPVersionParamsDateRange = "14dControl"
	AS112TopIPVersionParamsDateRange28dControl AS112TopIPVersionParamsDateRange = "28dControl"
	AS112TopIPVersionParamsDateRange12wControl AS112TopIPVersionParamsDateRange = "12wControl"
	AS112TopIPVersionParamsDateRange24wControl AS112TopIPVersionParamsDateRange = "24wControl"
)

func (r AS112TopIPVersionParamsDateRange) IsKnown() bool {
	switch r {
	case AS112TopIPVersionParamsDateRange1d, AS112TopIPVersionParamsDateRange2d, AS112TopIPVersionParamsDateRange7d, AS112TopIPVersionParamsDateRange14d, AS112TopIPVersionParamsDateRange28d, AS112TopIPVersionParamsDateRange12w, AS112TopIPVersionParamsDateRange24w, AS112TopIPVersionParamsDateRange52w, AS112TopIPVersionParamsDateRange1dControl, AS112TopIPVersionParamsDateRange2dControl, AS112TopIPVersionParamsDateRange7dControl, AS112TopIPVersionParamsDateRange14dControl, AS112TopIPVersionParamsDateRange28dControl, AS112TopIPVersionParamsDateRange12wControl, AS112TopIPVersionParamsDateRange24wControl:
		return true
	}
	return false
}

// Format results are returned in.
type AS112TopIPVersionParamsFormat string

const (
	AS112TopIPVersionParamsFormatJson AS112TopIPVersionParamsFormat = "JSON"
	AS112TopIPVersionParamsFormatCsv  AS112TopIPVersionParamsFormat = "CSV"
)

func (r AS112TopIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case AS112TopIPVersionParamsFormatJson, AS112TopIPVersionParamsFormatCsv:
		return true
	}
	return false
}

type AS112TopIPVersionResponseEnvelope struct {
	Result  AS112TopIPVersionResponse             `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    as112TopIPVersionResponseEnvelopeJSON `json:"-"`
}

// as112TopIPVersionResponseEnvelopeJSON contains the JSON metadata for the struct
// [AS112TopIPVersionResponseEnvelope]
type as112TopIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopIPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AS112TopLocationsParams struct {
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
	DateRange param.Field[[]AS112TopLocationsParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AS112TopLocationsParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AS112TopLocationsParams]'s query parameters as
// `url.Values`.
func (r AS112TopLocationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AS112TopLocationsParamsDateRange string

const (
	AS112TopLocationsParamsDateRange1d         AS112TopLocationsParamsDateRange = "1d"
	AS112TopLocationsParamsDateRange2d         AS112TopLocationsParamsDateRange = "2d"
	AS112TopLocationsParamsDateRange7d         AS112TopLocationsParamsDateRange = "7d"
	AS112TopLocationsParamsDateRange14d        AS112TopLocationsParamsDateRange = "14d"
	AS112TopLocationsParamsDateRange28d        AS112TopLocationsParamsDateRange = "28d"
	AS112TopLocationsParamsDateRange12w        AS112TopLocationsParamsDateRange = "12w"
	AS112TopLocationsParamsDateRange24w        AS112TopLocationsParamsDateRange = "24w"
	AS112TopLocationsParamsDateRange52w        AS112TopLocationsParamsDateRange = "52w"
	AS112TopLocationsParamsDateRange1dControl  AS112TopLocationsParamsDateRange = "1dControl"
	AS112TopLocationsParamsDateRange2dControl  AS112TopLocationsParamsDateRange = "2dControl"
	AS112TopLocationsParamsDateRange7dControl  AS112TopLocationsParamsDateRange = "7dControl"
	AS112TopLocationsParamsDateRange14dControl AS112TopLocationsParamsDateRange = "14dControl"
	AS112TopLocationsParamsDateRange28dControl AS112TopLocationsParamsDateRange = "28dControl"
	AS112TopLocationsParamsDateRange12wControl AS112TopLocationsParamsDateRange = "12wControl"
	AS112TopLocationsParamsDateRange24wControl AS112TopLocationsParamsDateRange = "24wControl"
)

func (r AS112TopLocationsParamsDateRange) IsKnown() bool {
	switch r {
	case AS112TopLocationsParamsDateRange1d, AS112TopLocationsParamsDateRange2d, AS112TopLocationsParamsDateRange7d, AS112TopLocationsParamsDateRange14d, AS112TopLocationsParamsDateRange28d, AS112TopLocationsParamsDateRange12w, AS112TopLocationsParamsDateRange24w, AS112TopLocationsParamsDateRange52w, AS112TopLocationsParamsDateRange1dControl, AS112TopLocationsParamsDateRange2dControl, AS112TopLocationsParamsDateRange7dControl, AS112TopLocationsParamsDateRange14dControl, AS112TopLocationsParamsDateRange28dControl, AS112TopLocationsParamsDateRange12wControl, AS112TopLocationsParamsDateRange24wControl:
		return true
	}
	return false
}

// Format results are returned in.
type AS112TopLocationsParamsFormat string

const (
	AS112TopLocationsParamsFormatJson AS112TopLocationsParamsFormat = "JSON"
	AS112TopLocationsParamsFormatCsv  AS112TopLocationsParamsFormat = "CSV"
)

func (r AS112TopLocationsParamsFormat) IsKnown() bool {
	switch r {
	case AS112TopLocationsParamsFormatJson, AS112TopLocationsParamsFormatCsv:
		return true
	}
	return false
}

type AS112TopLocationsResponseEnvelope struct {
	Result  AS112TopLocationsResponse             `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    as112TopLocationsResponseEnvelopeJSON `json:"-"`
}

// as112TopLocationsResponseEnvelopeJSON contains the JSON metadata for the struct
// [AS112TopLocationsResponseEnvelope]
type as112TopLocationsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopLocationsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopLocationsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
