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

// HTTPTopService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewHTTPTopService] method instead.
type HTTPTopService struct {
	Options []option.RequestOption
}

// NewHTTPTopService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHTTPTopService(opts ...option.RequestOption) (r *HTTPTopService) {
	r = &HTTPTopService{}
	r.Options = opts
	return
}

// Get the top user agents aggregated in families by HTTP traffic. Values are a
// percentage out of the total traffic.
func (r *HTTPTopService) BrowserFamilies(ctx context.Context, query HTTPTopBrowserFamiliesParams, opts ...option.RequestOption) (res *HTTPTopBrowserFamiliesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HTTPTopBrowserFamiliesResponseEnvelope
	path := "radar/http/top/browser_families"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the top user agents by HTTP traffic. Values are a percentage out of the
// total traffic.
func (r *HTTPTopService) Browsers(ctx context.Context, query HTTPTopBrowsersParams, opts ...option.RequestOption) (res *HTTPTopBrowsersResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HTTPTopBrowsersResponseEnvelope
	path := "radar/http/top/browsers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPTopBrowserFamiliesResponse struct {
	Meta HTTPTopBrowserFamiliesResponseMeta   `json:"meta,required"`
	Top0 []HTTPTopBrowserFamiliesResponseTop0 `json:"top_0,required"`
	JSON httpTopBrowserFamiliesResponseJSON   `json:"-"`
}

// httpTopBrowserFamiliesResponseJSON contains the JSON metadata for the struct
// [HTTPTopBrowserFamiliesResponse]
type httpTopBrowserFamiliesResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowserFamiliesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserFamiliesResponseJSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowserFamiliesResponseMeta struct {
	DateRange      []HTTPTopBrowserFamiliesResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                           `json:"lastUpdated,required"`
	ConfidenceInfo HTTPTopBrowserFamiliesResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           httpTopBrowserFamiliesResponseMetaJSON           `json:"-"`
}

// httpTopBrowserFamiliesResponseMetaJSON contains the JSON metadata for the struct
// [HTTPTopBrowserFamiliesResponseMeta]
type httpTopBrowserFamiliesResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPTopBrowserFamiliesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserFamiliesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowserFamiliesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      httpTopBrowserFamiliesResponseMetaDateRangeJSON `json:"-"`
}

// httpTopBrowserFamiliesResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [HTTPTopBrowserFamiliesResponseMetaDateRange]
type httpTopBrowserFamiliesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowserFamiliesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserFamiliesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowserFamiliesResponseMetaConfidenceInfo struct {
	Annotations []HTTPTopBrowserFamiliesResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                        `json:"level"`
	JSON        httpTopBrowserFamiliesResponseMetaConfidenceInfoJSON         `json:"-"`
}

// httpTopBrowserFamiliesResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [HTTPTopBrowserFamiliesResponseMetaConfidenceInfo]
type httpTopBrowserFamiliesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowserFamiliesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserFamiliesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowserFamiliesResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                         `json:"dataSource,required"`
	Description     string                                                         `json:"description,required"`
	EventType       string                                                         `json:"eventType,required"`
	IsInstantaneous interface{}                                                    `json:"isInstantaneous,required"`
	EndTime         time.Time                                                      `json:"endTime" format:"date-time"`
	LinkedURL       string                                                         `json:"linkedUrl"`
	StartTime       time.Time                                                      `json:"startTime" format:"date-time"`
	JSON            httpTopBrowserFamiliesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpTopBrowserFamiliesResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [HTTPTopBrowserFamiliesResponseMetaConfidenceInfoAnnotation]
type httpTopBrowserFamiliesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPTopBrowserFamiliesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserFamiliesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowserFamiliesResponseTop0 struct {
	Name  string                                 `json:"name,required"`
	Value string                                 `json:"value,required"`
	JSON  httpTopBrowserFamiliesResponseTop0JSON `json:"-"`
}

// httpTopBrowserFamiliesResponseTop0JSON contains the JSON metadata for the struct
// [HTTPTopBrowserFamiliesResponseTop0]
type httpTopBrowserFamiliesResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowserFamiliesResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserFamiliesResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowsersResponse struct {
	Meta HTTPTopBrowsersResponseMeta   `json:"meta,required"`
	Top0 []HTTPTopBrowsersResponseTop0 `json:"top_0,required"`
	JSON httpTopBrowsersResponseJSON   `json:"-"`
}

// httpTopBrowsersResponseJSON contains the JSON metadata for the struct
// [HTTPTopBrowsersResponse]
type httpTopBrowsersResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowsersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowsersResponseJSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowsersResponseMeta struct {
	DateRange      []HTTPTopBrowsersResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                    `json:"lastUpdated,required"`
	ConfidenceInfo HTTPTopBrowsersResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           httpTopBrowsersResponseMetaJSON           `json:"-"`
}

// httpTopBrowsersResponseMetaJSON contains the JSON metadata for the struct
// [HTTPTopBrowsersResponseMeta]
type httpTopBrowsersResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPTopBrowsersResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowsersResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowsersResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                `json:"startTime,required" format:"date-time"`
	JSON      httpTopBrowsersResponseMetaDateRangeJSON `json:"-"`
}

// httpTopBrowsersResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [HTTPTopBrowsersResponseMetaDateRange]
type httpTopBrowsersResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowsersResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowsersResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowsersResponseMetaConfidenceInfo struct {
	Annotations []HTTPTopBrowsersResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                 `json:"level"`
	JSON        httpTopBrowsersResponseMetaConfidenceInfoJSON         `json:"-"`
}

// httpTopBrowsersResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [HTTPTopBrowsersResponseMetaConfidenceInfo]
type httpTopBrowsersResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowsersResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowsersResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowsersResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                  `json:"dataSource,required"`
	Description     string                                                  `json:"description,required"`
	EventType       string                                                  `json:"eventType,required"`
	IsInstantaneous interface{}                                             `json:"isInstantaneous,required"`
	EndTime         time.Time                                               `json:"endTime" format:"date-time"`
	LinkedURL       string                                                  `json:"linkedUrl"`
	StartTime       time.Time                                               `json:"startTime" format:"date-time"`
	JSON            httpTopBrowsersResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpTopBrowsersResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [HTTPTopBrowsersResponseMetaConfidenceInfoAnnotation]
type httpTopBrowsersResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPTopBrowsersResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowsersResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowsersResponseTop0 struct {
	Name  string                          `json:"name,required"`
	Value string                          `json:"value,required"`
	JSON  httpTopBrowsersResponseTop0JSON `json:"-"`
}

// httpTopBrowsersResponseTop0JSON contains the JSON metadata for the struct
// [HTTPTopBrowsersResponseTop0]
type httpTopBrowsersResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowsersResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowsersResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowserFamiliesParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPTopBrowserFamiliesParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]HTTPTopBrowserFamiliesParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]HTTPTopBrowserFamiliesParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[HTTPTopBrowserFamiliesParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]HTTPTopBrowserFamiliesParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]HTTPTopBrowserFamiliesParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]HTTPTopBrowserFamiliesParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]HTTPTopBrowserFamiliesParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]HTTPTopBrowserFamiliesParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPTopBrowserFamiliesParams]'s query parameters as
// `url.Values`.
func (r HTTPTopBrowserFamiliesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HTTPTopBrowserFamiliesParamsBotClass string

const (
	HTTPTopBrowserFamiliesParamsBotClassLikelyAutomated HTTPTopBrowserFamiliesParamsBotClass = "LIKELY_AUTOMATED"
	HTTPTopBrowserFamiliesParamsBotClassLikelyHuman     HTTPTopBrowserFamiliesParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPTopBrowserFamiliesParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPTopBrowserFamiliesParamsBotClassLikelyAutomated, HTTPTopBrowserFamiliesParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPTopBrowserFamiliesParamsDateRange string

const (
	HTTPTopBrowserFamiliesParamsDateRange1d         HTTPTopBrowserFamiliesParamsDateRange = "1d"
	HTTPTopBrowserFamiliesParamsDateRange2d         HTTPTopBrowserFamiliesParamsDateRange = "2d"
	HTTPTopBrowserFamiliesParamsDateRange7d         HTTPTopBrowserFamiliesParamsDateRange = "7d"
	HTTPTopBrowserFamiliesParamsDateRange14d        HTTPTopBrowserFamiliesParamsDateRange = "14d"
	HTTPTopBrowserFamiliesParamsDateRange28d        HTTPTopBrowserFamiliesParamsDateRange = "28d"
	HTTPTopBrowserFamiliesParamsDateRange12w        HTTPTopBrowserFamiliesParamsDateRange = "12w"
	HTTPTopBrowserFamiliesParamsDateRange24w        HTTPTopBrowserFamiliesParamsDateRange = "24w"
	HTTPTopBrowserFamiliesParamsDateRange52w        HTTPTopBrowserFamiliesParamsDateRange = "52w"
	HTTPTopBrowserFamiliesParamsDateRange1dControl  HTTPTopBrowserFamiliesParamsDateRange = "1dControl"
	HTTPTopBrowserFamiliesParamsDateRange2dControl  HTTPTopBrowserFamiliesParamsDateRange = "2dControl"
	HTTPTopBrowserFamiliesParamsDateRange7dControl  HTTPTopBrowserFamiliesParamsDateRange = "7dControl"
	HTTPTopBrowserFamiliesParamsDateRange14dControl HTTPTopBrowserFamiliesParamsDateRange = "14dControl"
	HTTPTopBrowserFamiliesParamsDateRange28dControl HTTPTopBrowserFamiliesParamsDateRange = "28dControl"
	HTTPTopBrowserFamiliesParamsDateRange12wControl HTTPTopBrowserFamiliesParamsDateRange = "12wControl"
	HTTPTopBrowserFamiliesParamsDateRange24wControl HTTPTopBrowserFamiliesParamsDateRange = "24wControl"
)

func (r HTTPTopBrowserFamiliesParamsDateRange) IsKnown() bool {
	switch r {
	case HTTPTopBrowserFamiliesParamsDateRange1d, HTTPTopBrowserFamiliesParamsDateRange2d, HTTPTopBrowserFamiliesParamsDateRange7d, HTTPTopBrowserFamiliesParamsDateRange14d, HTTPTopBrowserFamiliesParamsDateRange28d, HTTPTopBrowserFamiliesParamsDateRange12w, HTTPTopBrowserFamiliesParamsDateRange24w, HTTPTopBrowserFamiliesParamsDateRange52w, HTTPTopBrowserFamiliesParamsDateRange1dControl, HTTPTopBrowserFamiliesParamsDateRange2dControl, HTTPTopBrowserFamiliesParamsDateRange7dControl, HTTPTopBrowserFamiliesParamsDateRange14dControl, HTTPTopBrowserFamiliesParamsDateRange28dControl, HTTPTopBrowserFamiliesParamsDateRange12wControl, HTTPTopBrowserFamiliesParamsDateRange24wControl:
		return true
	}
	return false
}

type HTTPTopBrowserFamiliesParamsDeviceType string

const (
	HTTPTopBrowserFamiliesParamsDeviceTypeDesktop HTTPTopBrowserFamiliesParamsDeviceType = "DESKTOP"
	HTTPTopBrowserFamiliesParamsDeviceTypeMobile  HTTPTopBrowserFamiliesParamsDeviceType = "MOBILE"
	HTTPTopBrowserFamiliesParamsDeviceTypeOther   HTTPTopBrowserFamiliesParamsDeviceType = "OTHER"
)

func (r HTTPTopBrowserFamiliesParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPTopBrowserFamiliesParamsDeviceTypeDesktop, HTTPTopBrowserFamiliesParamsDeviceTypeMobile, HTTPTopBrowserFamiliesParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format results are returned in.
type HTTPTopBrowserFamiliesParamsFormat string

const (
	HTTPTopBrowserFamiliesParamsFormatJson HTTPTopBrowserFamiliesParamsFormat = "JSON"
	HTTPTopBrowserFamiliesParamsFormatCsv  HTTPTopBrowserFamiliesParamsFormat = "CSV"
)

func (r HTTPTopBrowserFamiliesParamsFormat) IsKnown() bool {
	switch r {
	case HTTPTopBrowserFamiliesParamsFormatJson, HTTPTopBrowserFamiliesParamsFormatCsv:
		return true
	}
	return false
}

type HTTPTopBrowserFamiliesParamsHTTPProtocol string

const (
	HTTPTopBrowserFamiliesParamsHTTPProtocolHTTP  HTTPTopBrowserFamiliesParamsHTTPProtocol = "HTTP"
	HTTPTopBrowserFamiliesParamsHTTPProtocolHTTPS HTTPTopBrowserFamiliesParamsHTTPProtocol = "HTTPS"
)

func (r HTTPTopBrowserFamiliesParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPTopBrowserFamiliesParamsHTTPProtocolHTTP, HTTPTopBrowserFamiliesParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPTopBrowserFamiliesParamsHTTPVersion string

const (
	HTTPTopBrowserFamiliesParamsHTTPVersionHttPv1 HTTPTopBrowserFamiliesParamsHTTPVersion = "HTTPv1"
	HTTPTopBrowserFamiliesParamsHTTPVersionHttPv2 HTTPTopBrowserFamiliesParamsHTTPVersion = "HTTPv2"
	HTTPTopBrowserFamiliesParamsHTTPVersionHttPv3 HTTPTopBrowserFamiliesParamsHTTPVersion = "HTTPv3"
)

func (r HTTPTopBrowserFamiliesParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPTopBrowserFamiliesParamsHTTPVersionHttPv1, HTTPTopBrowserFamiliesParamsHTTPVersionHttPv2, HTTPTopBrowserFamiliesParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPTopBrowserFamiliesParamsIPVersion string

const (
	HTTPTopBrowserFamiliesParamsIPVersionIPv4 HTTPTopBrowserFamiliesParamsIPVersion = "IPv4"
	HTTPTopBrowserFamiliesParamsIPVersionIPv6 HTTPTopBrowserFamiliesParamsIPVersion = "IPv6"
)

func (r HTTPTopBrowserFamiliesParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPTopBrowserFamiliesParamsIPVersionIPv4, HTTPTopBrowserFamiliesParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPTopBrowserFamiliesParamsOS string

const (
	HTTPTopBrowserFamiliesParamsOSWindows  HTTPTopBrowserFamiliesParamsOS = "WINDOWS"
	HTTPTopBrowserFamiliesParamsOSMacosx   HTTPTopBrowserFamiliesParamsOS = "MACOSX"
	HTTPTopBrowserFamiliesParamsOSIos      HTTPTopBrowserFamiliesParamsOS = "IOS"
	HTTPTopBrowserFamiliesParamsOSAndroid  HTTPTopBrowserFamiliesParamsOS = "ANDROID"
	HTTPTopBrowserFamiliesParamsOSChromeos HTTPTopBrowserFamiliesParamsOS = "CHROMEOS"
	HTTPTopBrowserFamiliesParamsOSLinux    HTTPTopBrowserFamiliesParamsOS = "LINUX"
	HTTPTopBrowserFamiliesParamsOSSmartTv  HTTPTopBrowserFamiliesParamsOS = "SMART_TV"
)

func (r HTTPTopBrowserFamiliesParamsOS) IsKnown() bool {
	switch r {
	case HTTPTopBrowserFamiliesParamsOSWindows, HTTPTopBrowserFamiliesParamsOSMacosx, HTTPTopBrowserFamiliesParamsOSIos, HTTPTopBrowserFamiliesParamsOSAndroid, HTTPTopBrowserFamiliesParamsOSChromeos, HTTPTopBrowserFamiliesParamsOSLinux, HTTPTopBrowserFamiliesParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPTopBrowserFamiliesParamsTLSVersion string

const (
	HTTPTopBrowserFamiliesParamsTLSVersionTlSv1_0  HTTPTopBrowserFamiliesParamsTLSVersion = "TLSv1_0"
	HTTPTopBrowserFamiliesParamsTLSVersionTlSv1_1  HTTPTopBrowserFamiliesParamsTLSVersion = "TLSv1_1"
	HTTPTopBrowserFamiliesParamsTLSVersionTlSv1_2  HTTPTopBrowserFamiliesParamsTLSVersion = "TLSv1_2"
	HTTPTopBrowserFamiliesParamsTLSVersionTlSv1_3  HTTPTopBrowserFamiliesParamsTLSVersion = "TLSv1_3"
	HTTPTopBrowserFamiliesParamsTLSVersionTlSvQuic HTTPTopBrowserFamiliesParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPTopBrowserFamiliesParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPTopBrowserFamiliesParamsTLSVersionTlSv1_0, HTTPTopBrowserFamiliesParamsTLSVersionTlSv1_1, HTTPTopBrowserFamiliesParamsTLSVersionTlSv1_2, HTTPTopBrowserFamiliesParamsTLSVersionTlSv1_3, HTTPTopBrowserFamiliesParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPTopBrowserFamiliesResponseEnvelope struct {
	Result  HTTPTopBrowserFamiliesResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    httpTopBrowserFamiliesResponseEnvelopeJSON `json:"-"`
}

// httpTopBrowserFamiliesResponseEnvelopeJSON contains the JSON metadata for the
// struct [HTTPTopBrowserFamiliesResponseEnvelope]
type httpTopBrowserFamiliesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowserFamiliesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserFamiliesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowsersParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPTopBrowsersParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]HTTPTopBrowsersParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]HTTPTopBrowsersParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[HTTPTopBrowsersParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]HTTPTopBrowsersParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]HTTPTopBrowsersParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]HTTPTopBrowsersParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]HTTPTopBrowsersParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]HTTPTopBrowsersParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPTopBrowsersParams]'s query parameters as `url.Values`.
func (r HTTPTopBrowsersParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HTTPTopBrowsersParamsBotClass string

const (
	HTTPTopBrowsersParamsBotClassLikelyAutomated HTTPTopBrowsersParamsBotClass = "LIKELY_AUTOMATED"
	HTTPTopBrowsersParamsBotClassLikelyHuman     HTTPTopBrowsersParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPTopBrowsersParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPTopBrowsersParamsBotClassLikelyAutomated, HTTPTopBrowsersParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPTopBrowsersParamsDateRange string

const (
	HTTPTopBrowsersParamsDateRange1d         HTTPTopBrowsersParamsDateRange = "1d"
	HTTPTopBrowsersParamsDateRange2d         HTTPTopBrowsersParamsDateRange = "2d"
	HTTPTopBrowsersParamsDateRange7d         HTTPTopBrowsersParamsDateRange = "7d"
	HTTPTopBrowsersParamsDateRange14d        HTTPTopBrowsersParamsDateRange = "14d"
	HTTPTopBrowsersParamsDateRange28d        HTTPTopBrowsersParamsDateRange = "28d"
	HTTPTopBrowsersParamsDateRange12w        HTTPTopBrowsersParamsDateRange = "12w"
	HTTPTopBrowsersParamsDateRange24w        HTTPTopBrowsersParamsDateRange = "24w"
	HTTPTopBrowsersParamsDateRange52w        HTTPTopBrowsersParamsDateRange = "52w"
	HTTPTopBrowsersParamsDateRange1dControl  HTTPTopBrowsersParamsDateRange = "1dControl"
	HTTPTopBrowsersParamsDateRange2dControl  HTTPTopBrowsersParamsDateRange = "2dControl"
	HTTPTopBrowsersParamsDateRange7dControl  HTTPTopBrowsersParamsDateRange = "7dControl"
	HTTPTopBrowsersParamsDateRange14dControl HTTPTopBrowsersParamsDateRange = "14dControl"
	HTTPTopBrowsersParamsDateRange28dControl HTTPTopBrowsersParamsDateRange = "28dControl"
	HTTPTopBrowsersParamsDateRange12wControl HTTPTopBrowsersParamsDateRange = "12wControl"
	HTTPTopBrowsersParamsDateRange24wControl HTTPTopBrowsersParamsDateRange = "24wControl"
)

func (r HTTPTopBrowsersParamsDateRange) IsKnown() bool {
	switch r {
	case HTTPTopBrowsersParamsDateRange1d, HTTPTopBrowsersParamsDateRange2d, HTTPTopBrowsersParamsDateRange7d, HTTPTopBrowsersParamsDateRange14d, HTTPTopBrowsersParamsDateRange28d, HTTPTopBrowsersParamsDateRange12w, HTTPTopBrowsersParamsDateRange24w, HTTPTopBrowsersParamsDateRange52w, HTTPTopBrowsersParamsDateRange1dControl, HTTPTopBrowsersParamsDateRange2dControl, HTTPTopBrowsersParamsDateRange7dControl, HTTPTopBrowsersParamsDateRange14dControl, HTTPTopBrowsersParamsDateRange28dControl, HTTPTopBrowsersParamsDateRange12wControl, HTTPTopBrowsersParamsDateRange24wControl:
		return true
	}
	return false
}

type HTTPTopBrowsersParamsDeviceType string

const (
	HTTPTopBrowsersParamsDeviceTypeDesktop HTTPTopBrowsersParamsDeviceType = "DESKTOP"
	HTTPTopBrowsersParamsDeviceTypeMobile  HTTPTopBrowsersParamsDeviceType = "MOBILE"
	HTTPTopBrowsersParamsDeviceTypeOther   HTTPTopBrowsersParamsDeviceType = "OTHER"
)

func (r HTTPTopBrowsersParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPTopBrowsersParamsDeviceTypeDesktop, HTTPTopBrowsersParamsDeviceTypeMobile, HTTPTopBrowsersParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format results are returned in.
type HTTPTopBrowsersParamsFormat string

const (
	HTTPTopBrowsersParamsFormatJson HTTPTopBrowsersParamsFormat = "JSON"
	HTTPTopBrowsersParamsFormatCsv  HTTPTopBrowsersParamsFormat = "CSV"
)

func (r HTTPTopBrowsersParamsFormat) IsKnown() bool {
	switch r {
	case HTTPTopBrowsersParamsFormatJson, HTTPTopBrowsersParamsFormatCsv:
		return true
	}
	return false
}

type HTTPTopBrowsersParamsHTTPProtocol string

const (
	HTTPTopBrowsersParamsHTTPProtocolHTTP  HTTPTopBrowsersParamsHTTPProtocol = "HTTP"
	HTTPTopBrowsersParamsHTTPProtocolHTTPS HTTPTopBrowsersParamsHTTPProtocol = "HTTPS"
)

func (r HTTPTopBrowsersParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPTopBrowsersParamsHTTPProtocolHTTP, HTTPTopBrowsersParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPTopBrowsersParamsHTTPVersion string

const (
	HTTPTopBrowsersParamsHTTPVersionHttPv1 HTTPTopBrowsersParamsHTTPVersion = "HTTPv1"
	HTTPTopBrowsersParamsHTTPVersionHttPv2 HTTPTopBrowsersParamsHTTPVersion = "HTTPv2"
	HTTPTopBrowsersParamsHTTPVersionHttPv3 HTTPTopBrowsersParamsHTTPVersion = "HTTPv3"
)

func (r HTTPTopBrowsersParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPTopBrowsersParamsHTTPVersionHttPv1, HTTPTopBrowsersParamsHTTPVersionHttPv2, HTTPTopBrowsersParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPTopBrowsersParamsIPVersion string

const (
	HTTPTopBrowsersParamsIPVersionIPv4 HTTPTopBrowsersParamsIPVersion = "IPv4"
	HTTPTopBrowsersParamsIPVersionIPv6 HTTPTopBrowsersParamsIPVersion = "IPv6"
)

func (r HTTPTopBrowsersParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPTopBrowsersParamsIPVersionIPv4, HTTPTopBrowsersParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPTopBrowsersParamsOS string

const (
	HTTPTopBrowsersParamsOSWindows  HTTPTopBrowsersParamsOS = "WINDOWS"
	HTTPTopBrowsersParamsOSMacosx   HTTPTopBrowsersParamsOS = "MACOSX"
	HTTPTopBrowsersParamsOSIos      HTTPTopBrowsersParamsOS = "IOS"
	HTTPTopBrowsersParamsOSAndroid  HTTPTopBrowsersParamsOS = "ANDROID"
	HTTPTopBrowsersParamsOSChromeos HTTPTopBrowsersParamsOS = "CHROMEOS"
	HTTPTopBrowsersParamsOSLinux    HTTPTopBrowsersParamsOS = "LINUX"
	HTTPTopBrowsersParamsOSSmartTv  HTTPTopBrowsersParamsOS = "SMART_TV"
)

func (r HTTPTopBrowsersParamsOS) IsKnown() bool {
	switch r {
	case HTTPTopBrowsersParamsOSWindows, HTTPTopBrowsersParamsOSMacosx, HTTPTopBrowsersParamsOSIos, HTTPTopBrowsersParamsOSAndroid, HTTPTopBrowsersParamsOSChromeos, HTTPTopBrowsersParamsOSLinux, HTTPTopBrowsersParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPTopBrowsersParamsTLSVersion string

const (
	HTTPTopBrowsersParamsTLSVersionTlSv1_0  HTTPTopBrowsersParamsTLSVersion = "TLSv1_0"
	HTTPTopBrowsersParamsTLSVersionTlSv1_1  HTTPTopBrowsersParamsTLSVersion = "TLSv1_1"
	HTTPTopBrowsersParamsTLSVersionTlSv1_2  HTTPTopBrowsersParamsTLSVersion = "TLSv1_2"
	HTTPTopBrowsersParamsTLSVersionTlSv1_3  HTTPTopBrowsersParamsTLSVersion = "TLSv1_3"
	HTTPTopBrowsersParamsTLSVersionTlSvQuic HTTPTopBrowsersParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPTopBrowsersParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPTopBrowsersParamsTLSVersionTlSv1_0, HTTPTopBrowsersParamsTLSVersionTlSv1_1, HTTPTopBrowsersParamsTLSVersionTlSv1_2, HTTPTopBrowsersParamsTLSVersionTlSv1_3, HTTPTopBrowsersParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPTopBrowsersResponseEnvelope struct {
	Result  HTTPTopBrowsersResponse             `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    httpTopBrowsersResponseEnvelopeJSON `json:"-"`
}

// httpTopBrowsersResponseEnvelopeJSON contains the JSON metadata for the struct
// [HTTPTopBrowsersResponseEnvelope]
type httpTopBrowsersResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowsersResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowsersResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
