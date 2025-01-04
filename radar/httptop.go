// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// HTTPTopService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTTPTopService] method instead.
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

// Get the top user agents by HTTP traffic. Values are a percentage out of the
// total traffic.
func (r *HTTPTopService) Browser(ctx context.Context, query HTTPTopBrowserParams, opts ...option.RequestOption) (res *HTTPTopBrowserResponse, err error) {
	var env HTTPTopBrowserResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/http/top/browser"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the top user agents aggregated in families by HTTP traffic. Values are a
// percentage out of the total traffic.
func (r *HTTPTopService) BrowserFamily(ctx context.Context, query HTTPTopBrowserFamilyParams, opts ...option.RequestOption) (res *HTTPTopBrowserFamilyResponse, err error) {
	var env HTTPTopBrowserFamilyResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/http/top/browser_family"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPTopBrowserResponse struct {
	Meta HTTPTopBrowserResponseMeta   `json:"meta,required"`
	Top0 []HTTPTopBrowserResponseTop0 `json:"top_0,required"`
	JSON httpTopBrowserResponseJSON   `json:"-"`
}

// httpTopBrowserResponseJSON contains the JSON metadata for the struct
// [HTTPTopBrowserResponse]
type httpTopBrowserResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowserResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserResponseJSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowserResponseMeta struct {
	DateRange      []HTTPTopBrowserResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                   `json:"lastUpdated,required"`
	ConfidenceInfo HTTPTopBrowserResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           httpTopBrowserResponseMetaJSON           `json:"-"`
}

// httpTopBrowserResponseMetaJSON contains the JSON metadata for the struct
// [HTTPTopBrowserResponseMeta]
type httpTopBrowserResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPTopBrowserResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowserResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                               `json:"startTime,required" format:"date-time"`
	JSON      httpTopBrowserResponseMetaDateRangeJSON `json:"-"`
}

// httpTopBrowserResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [HTTPTopBrowserResponseMetaDateRange]
type httpTopBrowserResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowserResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowserResponseMetaConfidenceInfo struct {
	Annotations []HTTPTopBrowserResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                `json:"level"`
	JSON        httpTopBrowserResponseMetaConfidenceInfoJSON         `json:"-"`
}

// httpTopBrowserResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [HTTPTopBrowserResponseMetaConfidenceInfo]
type httpTopBrowserResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowserResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowserResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                 `json:"dataSource,required"`
	Description     string                                                 `json:"description,required"`
	EventType       string                                                 `json:"eventType,required"`
	IsInstantaneous bool                                                   `json:"isInstantaneous,required"`
	EndTime         time.Time                                              `json:"endTime" format:"date-time"`
	LinkedURL       string                                                 `json:"linkedUrl"`
	StartTime       time.Time                                              `json:"startTime" format:"date-time"`
	JSON            httpTopBrowserResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpTopBrowserResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [HTTPTopBrowserResponseMetaConfidenceInfoAnnotation]
type httpTopBrowserResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPTopBrowserResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowserResponseTop0 struct {
	Name  string                         `json:"name,required"`
	Value string                         `json:"value,required"`
	JSON  httpTopBrowserResponseTop0JSON `json:"-"`
}

// httpTopBrowserResponseTop0JSON contains the JSON metadata for the struct
// [HTTPTopBrowserResponseTop0]
type httpTopBrowserResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowserResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowserFamilyResponse struct {
	Meta HTTPTopBrowserFamilyResponseMeta   `json:"meta,required"`
	Top0 []HTTPTopBrowserFamilyResponseTop0 `json:"top_0,required"`
	JSON httpTopBrowserFamilyResponseJSON   `json:"-"`
}

// httpTopBrowserFamilyResponseJSON contains the JSON metadata for the struct
// [HTTPTopBrowserFamilyResponse]
type httpTopBrowserFamilyResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowserFamilyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserFamilyResponseJSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowserFamilyResponseMeta struct {
	DateRange      []HTTPTopBrowserFamilyResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                         `json:"lastUpdated,required"`
	ConfidenceInfo HTTPTopBrowserFamilyResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           httpTopBrowserFamilyResponseMetaJSON           `json:"-"`
}

// httpTopBrowserFamilyResponseMetaJSON contains the JSON metadata for the struct
// [HTTPTopBrowserFamilyResponseMeta]
type httpTopBrowserFamilyResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPTopBrowserFamilyResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserFamilyResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowserFamilyResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                     `json:"startTime,required" format:"date-time"`
	JSON      httpTopBrowserFamilyResponseMetaDateRangeJSON `json:"-"`
}

// httpTopBrowserFamilyResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [HTTPTopBrowserFamilyResponseMetaDateRange]
type httpTopBrowserFamilyResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowserFamilyResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserFamilyResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowserFamilyResponseMetaConfidenceInfo struct {
	Annotations []HTTPTopBrowserFamilyResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                      `json:"level"`
	JSON        httpTopBrowserFamilyResponseMetaConfidenceInfoJSON         `json:"-"`
}

// httpTopBrowserFamilyResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [HTTPTopBrowserFamilyResponseMetaConfidenceInfo]
type httpTopBrowserFamilyResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowserFamilyResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserFamilyResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowserFamilyResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                       `json:"dataSource,required"`
	Description     string                                                       `json:"description,required"`
	EventType       string                                                       `json:"eventType,required"`
	IsInstantaneous bool                                                         `json:"isInstantaneous,required"`
	EndTime         time.Time                                                    `json:"endTime" format:"date-time"`
	LinkedURL       string                                                       `json:"linkedUrl"`
	StartTime       time.Time                                                    `json:"startTime" format:"date-time"`
	JSON            httpTopBrowserFamilyResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpTopBrowserFamilyResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [HTTPTopBrowserFamilyResponseMetaConfidenceInfoAnnotation]
type httpTopBrowserFamilyResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPTopBrowserFamilyResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserFamilyResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowserFamilyResponseTop0 struct {
	Name  string                               `json:"name,required"`
	Value string                               `json:"value,required"`
	JSON  httpTopBrowserFamilyResponseTop0JSON `json:"-"`
}

// httpTopBrowserFamilyResponseTop0JSON contains the JSON metadata for the struct
// [HTTPTopBrowserFamilyResponseTop0]
type httpTopBrowserFamilyResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowserFamilyResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserFamilyResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowserParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPTopBrowserParamsBotClass] `query:"botClass"`
	// Filter for browser family.
	BrowserFamily param.Field[[]HTTPTopBrowserParamsBrowserFamily] `query:"browserFamily"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]HTTPTopBrowserParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[HTTPTopBrowserParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]HTTPTopBrowserParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]HTTPTopBrowserParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]HTTPTopBrowserParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]HTTPTopBrowserParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]HTTPTopBrowserParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPTopBrowserParams]'s query parameters as `url.Values`.
func (r HTTPTopBrowserParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type HTTPTopBrowserParamsBotClass string

const (
	HTTPTopBrowserParamsBotClassLikelyAutomated HTTPTopBrowserParamsBotClass = "LIKELY_AUTOMATED"
	HTTPTopBrowserParamsBotClassLikelyHuman     HTTPTopBrowserParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPTopBrowserParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPTopBrowserParamsBotClassLikelyAutomated, HTTPTopBrowserParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPTopBrowserParamsBrowserFamily string

const (
	HTTPTopBrowserParamsBrowserFamilyChrome  HTTPTopBrowserParamsBrowserFamily = "CHROME"
	HTTPTopBrowserParamsBrowserFamilyEdge    HTTPTopBrowserParamsBrowserFamily = "EDGE"
	HTTPTopBrowserParamsBrowserFamilyFirefox HTTPTopBrowserParamsBrowserFamily = "FIREFOX"
	HTTPTopBrowserParamsBrowserFamilySafari  HTTPTopBrowserParamsBrowserFamily = "SAFARI"
)

func (r HTTPTopBrowserParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPTopBrowserParamsBrowserFamilyChrome, HTTPTopBrowserParamsBrowserFamilyEdge, HTTPTopBrowserParamsBrowserFamilyFirefox, HTTPTopBrowserParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPTopBrowserParamsDeviceType string

const (
	HTTPTopBrowserParamsDeviceTypeDesktop HTTPTopBrowserParamsDeviceType = "DESKTOP"
	HTTPTopBrowserParamsDeviceTypeMobile  HTTPTopBrowserParamsDeviceType = "MOBILE"
	HTTPTopBrowserParamsDeviceTypeOther   HTTPTopBrowserParamsDeviceType = "OTHER"
)

func (r HTTPTopBrowserParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPTopBrowserParamsDeviceTypeDesktop, HTTPTopBrowserParamsDeviceTypeMobile, HTTPTopBrowserParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format results are returned in.
type HTTPTopBrowserParamsFormat string

const (
	HTTPTopBrowserParamsFormatJson HTTPTopBrowserParamsFormat = "JSON"
	HTTPTopBrowserParamsFormatCsv  HTTPTopBrowserParamsFormat = "CSV"
)

func (r HTTPTopBrowserParamsFormat) IsKnown() bool {
	switch r {
	case HTTPTopBrowserParamsFormatJson, HTTPTopBrowserParamsFormatCsv:
		return true
	}
	return false
}

type HTTPTopBrowserParamsHTTPProtocol string

const (
	HTTPTopBrowserParamsHTTPProtocolHTTP  HTTPTopBrowserParamsHTTPProtocol = "HTTP"
	HTTPTopBrowserParamsHTTPProtocolHTTPS HTTPTopBrowserParamsHTTPProtocol = "HTTPS"
)

func (r HTTPTopBrowserParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPTopBrowserParamsHTTPProtocolHTTP, HTTPTopBrowserParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPTopBrowserParamsHTTPVersion string

const (
	HTTPTopBrowserParamsHTTPVersionHttPv1 HTTPTopBrowserParamsHTTPVersion = "HTTPv1"
	HTTPTopBrowserParamsHTTPVersionHttPv2 HTTPTopBrowserParamsHTTPVersion = "HTTPv2"
	HTTPTopBrowserParamsHTTPVersionHttPv3 HTTPTopBrowserParamsHTTPVersion = "HTTPv3"
)

func (r HTTPTopBrowserParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPTopBrowserParamsHTTPVersionHttPv1, HTTPTopBrowserParamsHTTPVersionHttPv2, HTTPTopBrowserParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPTopBrowserParamsIPVersion string

const (
	HTTPTopBrowserParamsIPVersionIPv4 HTTPTopBrowserParamsIPVersion = "IPv4"
	HTTPTopBrowserParamsIPVersionIPv6 HTTPTopBrowserParamsIPVersion = "IPv6"
)

func (r HTTPTopBrowserParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPTopBrowserParamsIPVersionIPv4, HTTPTopBrowserParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPTopBrowserParamsOS string

const (
	HTTPTopBrowserParamsOSWindows  HTTPTopBrowserParamsOS = "WINDOWS"
	HTTPTopBrowserParamsOSMacosx   HTTPTopBrowserParamsOS = "MACOSX"
	HTTPTopBrowserParamsOSIos      HTTPTopBrowserParamsOS = "IOS"
	HTTPTopBrowserParamsOSAndroid  HTTPTopBrowserParamsOS = "ANDROID"
	HTTPTopBrowserParamsOSChromeos HTTPTopBrowserParamsOS = "CHROMEOS"
	HTTPTopBrowserParamsOSLinux    HTTPTopBrowserParamsOS = "LINUX"
	HTTPTopBrowserParamsOSSmartTv  HTTPTopBrowserParamsOS = "SMART_TV"
)

func (r HTTPTopBrowserParamsOS) IsKnown() bool {
	switch r {
	case HTTPTopBrowserParamsOSWindows, HTTPTopBrowserParamsOSMacosx, HTTPTopBrowserParamsOSIos, HTTPTopBrowserParamsOSAndroid, HTTPTopBrowserParamsOSChromeos, HTTPTopBrowserParamsOSLinux, HTTPTopBrowserParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPTopBrowserParamsTLSVersion string

const (
	HTTPTopBrowserParamsTLSVersionTlSv1_0  HTTPTopBrowserParamsTLSVersion = "TLSv1_0"
	HTTPTopBrowserParamsTLSVersionTlSv1_1  HTTPTopBrowserParamsTLSVersion = "TLSv1_1"
	HTTPTopBrowserParamsTLSVersionTlSv1_2  HTTPTopBrowserParamsTLSVersion = "TLSv1_2"
	HTTPTopBrowserParamsTLSVersionTlSv1_3  HTTPTopBrowserParamsTLSVersion = "TLSv1_3"
	HTTPTopBrowserParamsTLSVersionTlSvQuic HTTPTopBrowserParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPTopBrowserParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPTopBrowserParamsTLSVersionTlSv1_0, HTTPTopBrowserParamsTLSVersionTlSv1_1, HTTPTopBrowserParamsTLSVersionTlSv1_2, HTTPTopBrowserParamsTLSVersionTlSv1_3, HTTPTopBrowserParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPTopBrowserResponseEnvelope struct {
	Result  HTTPTopBrowserResponse             `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    httpTopBrowserResponseEnvelopeJSON `json:"-"`
}

// httpTopBrowserResponseEnvelopeJSON contains the JSON metadata for the struct
// [HTTPTopBrowserResponseEnvelope]
type httpTopBrowserResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowserResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowserFamilyParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPTopBrowserFamilyParamsBotClass] `query:"botClass"`
	// Filter for browser family.
	BrowserFamily param.Field[[]HTTPTopBrowserFamilyParamsBrowserFamily] `query:"browserFamily"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]HTTPTopBrowserFamilyParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[HTTPTopBrowserFamilyParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]HTTPTopBrowserFamilyParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]HTTPTopBrowserFamilyParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]HTTPTopBrowserFamilyParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]HTTPTopBrowserFamilyParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]HTTPTopBrowserFamilyParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPTopBrowserFamilyParams]'s query parameters as
// `url.Values`.
func (r HTTPTopBrowserFamilyParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type HTTPTopBrowserFamilyParamsBotClass string

const (
	HTTPTopBrowserFamilyParamsBotClassLikelyAutomated HTTPTopBrowserFamilyParamsBotClass = "LIKELY_AUTOMATED"
	HTTPTopBrowserFamilyParamsBotClassLikelyHuman     HTTPTopBrowserFamilyParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPTopBrowserFamilyParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPTopBrowserFamilyParamsBotClassLikelyAutomated, HTTPTopBrowserFamilyParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPTopBrowserFamilyParamsBrowserFamily string

const (
	HTTPTopBrowserFamilyParamsBrowserFamilyChrome  HTTPTopBrowserFamilyParamsBrowserFamily = "CHROME"
	HTTPTopBrowserFamilyParamsBrowserFamilyEdge    HTTPTopBrowserFamilyParamsBrowserFamily = "EDGE"
	HTTPTopBrowserFamilyParamsBrowserFamilyFirefox HTTPTopBrowserFamilyParamsBrowserFamily = "FIREFOX"
	HTTPTopBrowserFamilyParamsBrowserFamilySafari  HTTPTopBrowserFamilyParamsBrowserFamily = "SAFARI"
)

func (r HTTPTopBrowserFamilyParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPTopBrowserFamilyParamsBrowserFamilyChrome, HTTPTopBrowserFamilyParamsBrowserFamilyEdge, HTTPTopBrowserFamilyParamsBrowserFamilyFirefox, HTTPTopBrowserFamilyParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPTopBrowserFamilyParamsDeviceType string

const (
	HTTPTopBrowserFamilyParamsDeviceTypeDesktop HTTPTopBrowserFamilyParamsDeviceType = "DESKTOP"
	HTTPTopBrowserFamilyParamsDeviceTypeMobile  HTTPTopBrowserFamilyParamsDeviceType = "MOBILE"
	HTTPTopBrowserFamilyParamsDeviceTypeOther   HTTPTopBrowserFamilyParamsDeviceType = "OTHER"
)

func (r HTTPTopBrowserFamilyParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPTopBrowserFamilyParamsDeviceTypeDesktop, HTTPTopBrowserFamilyParamsDeviceTypeMobile, HTTPTopBrowserFamilyParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format results are returned in.
type HTTPTopBrowserFamilyParamsFormat string

const (
	HTTPTopBrowserFamilyParamsFormatJson HTTPTopBrowserFamilyParamsFormat = "JSON"
	HTTPTopBrowserFamilyParamsFormatCsv  HTTPTopBrowserFamilyParamsFormat = "CSV"
)

func (r HTTPTopBrowserFamilyParamsFormat) IsKnown() bool {
	switch r {
	case HTTPTopBrowserFamilyParamsFormatJson, HTTPTopBrowserFamilyParamsFormatCsv:
		return true
	}
	return false
}

type HTTPTopBrowserFamilyParamsHTTPProtocol string

const (
	HTTPTopBrowserFamilyParamsHTTPProtocolHTTP  HTTPTopBrowserFamilyParamsHTTPProtocol = "HTTP"
	HTTPTopBrowserFamilyParamsHTTPProtocolHTTPS HTTPTopBrowserFamilyParamsHTTPProtocol = "HTTPS"
)

func (r HTTPTopBrowserFamilyParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPTopBrowserFamilyParamsHTTPProtocolHTTP, HTTPTopBrowserFamilyParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPTopBrowserFamilyParamsHTTPVersion string

const (
	HTTPTopBrowserFamilyParamsHTTPVersionHttPv1 HTTPTopBrowserFamilyParamsHTTPVersion = "HTTPv1"
	HTTPTopBrowserFamilyParamsHTTPVersionHttPv2 HTTPTopBrowserFamilyParamsHTTPVersion = "HTTPv2"
	HTTPTopBrowserFamilyParamsHTTPVersionHttPv3 HTTPTopBrowserFamilyParamsHTTPVersion = "HTTPv3"
)

func (r HTTPTopBrowserFamilyParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPTopBrowserFamilyParamsHTTPVersionHttPv1, HTTPTopBrowserFamilyParamsHTTPVersionHttPv2, HTTPTopBrowserFamilyParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPTopBrowserFamilyParamsIPVersion string

const (
	HTTPTopBrowserFamilyParamsIPVersionIPv4 HTTPTopBrowserFamilyParamsIPVersion = "IPv4"
	HTTPTopBrowserFamilyParamsIPVersionIPv6 HTTPTopBrowserFamilyParamsIPVersion = "IPv6"
)

func (r HTTPTopBrowserFamilyParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPTopBrowserFamilyParamsIPVersionIPv4, HTTPTopBrowserFamilyParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPTopBrowserFamilyParamsOS string

const (
	HTTPTopBrowserFamilyParamsOSWindows  HTTPTopBrowserFamilyParamsOS = "WINDOWS"
	HTTPTopBrowserFamilyParamsOSMacosx   HTTPTopBrowserFamilyParamsOS = "MACOSX"
	HTTPTopBrowserFamilyParamsOSIos      HTTPTopBrowserFamilyParamsOS = "IOS"
	HTTPTopBrowserFamilyParamsOSAndroid  HTTPTopBrowserFamilyParamsOS = "ANDROID"
	HTTPTopBrowserFamilyParamsOSChromeos HTTPTopBrowserFamilyParamsOS = "CHROMEOS"
	HTTPTopBrowserFamilyParamsOSLinux    HTTPTopBrowserFamilyParamsOS = "LINUX"
	HTTPTopBrowserFamilyParamsOSSmartTv  HTTPTopBrowserFamilyParamsOS = "SMART_TV"
)

func (r HTTPTopBrowserFamilyParamsOS) IsKnown() bool {
	switch r {
	case HTTPTopBrowserFamilyParamsOSWindows, HTTPTopBrowserFamilyParamsOSMacosx, HTTPTopBrowserFamilyParamsOSIos, HTTPTopBrowserFamilyParamsOSAndroid, HTTPTopBrowserFamilyParamsOSChromeos, HTTPTopBrowserFamilyParamsOSLinux, HTTPTopBrowserFamilyParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPTopBrowserFamilyParamsTLSVersion string

const (
	HTTPTopBrowserFamilyParamsTLSVersionTlSv1_0  HTTPTopBrowserFamilyParamsTLSVersion = "TLSv1_0"
	HTTPTopBrowserFamilyParamsTLSVersionTlSv1_1  HTTPTopBrowserFamilyParamsTLSVersion = "TLSv1_1"
	HTTPTopBrowserFamilyParamsTLSVersionTlSv1_2  HTTPTopBrowserFamilyParamsTLSVersion = "TLSv1_2"
	HTTPTopBrowserFamilyParamsTLSVersionTlSv1_3  HTTPTopBrowserFamilyParamsTLSVersion = "TLSv1_3"
	HTTPTopBrowserFamilyParamsTLSVersionTlSvQuic HTTPTopBrowserFamilyParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPTopBrowserFamilyParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPTopBrowserFamilyParamsTLSVersionTlSv1_0, HTTPTopBrowserFamilyParamsTLSVersionTlSv1_1, HTTPTopBrowserFamilyParamsTLSVersionTlSv1_2, HTTPTopBrowserFamilyParamsTLSVersionTlSv1_3, HTTPTopBrowserFamilyParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPTopBrowserFamilyResponseEnvelope struct {
	Result  HTTPTopBrowserFamilyResponse             `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    httpTopBrowserFamilyResponseEnvelopeJSON `json:"-"`
}

// httpTopBrowserFamilyResponseEnvelopeJSON contains the JSON metadata for the
// struct [HTTPTopBrowserFamilyResponseEnvelope]
type httpTopBrowserFamilyResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowserFamilyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserFamilyResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
