// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// HTTPLocationTLSVersionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTTPLocationTLSVersionService] method instead.
type HTTPLocationTLSVersionService struct {
	Options []option.RequestOption
}

// NewHTTPLocationTLSVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewHTTPLocationTLSVersionService(opts ...option.RequestOption) (r *HTTPLocationTLSVersionService) {
	r = &HTTPLocationTLSVersionService{}
	r.Options = opts
	return
}

// Get the top locations, by HTTP traffic, of the requested TLS protocol version.
// Values are a percentage out of the total traffic.
func (r *HTTPLocationTLSVersionService) Get(ctx context.Context, tlsVersion HTTPLocationTLSVersionGetParamsTLSVersion, query HTTPLocationTLSVersionGetParams, opts ...option.RequestOption) (res *HTTPLocationTLSVersionGetResponse, err error) {
	var env HTTPLocationTLSVersionGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/locations/tls_version/%v", tlsVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPLocationTLSVersionGetResponse struct {
	Meta HTTPLocationTLSVersionGetResponseMeta   `json:"meta,required"`
	Top0 []HTTPLocationTLSVersionGetResponseTop0 `json:"top_0,required"`
	JSON httpLocationTLSVersionGetResponseJSON   `json:"-"`
}

// httpLocationTLSVersionGetResponseJSON contains the JSON metadata for the struct
// [HTTPLocationTLSVersionGetResponse]
type httpLocationTLSVersionGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationTLSVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationTLSVersionGetResponseJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationTLSVersionGetResponseMeta struct {
	DateRange      []HTTPLocationTLSVersionGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                              `json:"lastUpdated,required"`
	ConfidenceInfo HTTPLocationTLSVersionGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           httpLocationTLSVersionGetResponseMetaJSON           `json:"-"`
}

// httpLocationTLSVersionGetResponseMetaJSON contains the JSON metadata for the
// struct [HTTPLocationTLSVersionGetResponseMeta]
type httpLocationTLSVersionGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPLocationTLSVersionGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationTLSVersionGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationTLSVersionGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                          `json:"startTime,required" format:"date-time"`
	JSON      httpLocationTLSVersionGetResponseMetaDateRangeJSON `json:"-"`
}

// httpLocationTLSVersionGetResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [HTTPLocationTLSVersionGetResponseMetaDateRange]
type httpLocationTLSVersionGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationTLSVersionGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationTLSVersionGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationTLSVersionGetResponseMetaConfidenceInfo struct {
	Annotations []HTTPLocationTLSVersionGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                           `json:"level"`
	JSON        httpLocationTLSVersionGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// httpLocationTLSVersionGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [HTTPLocationTLSVersionGetResponseMetaConfidenceInfo]
type httpLocationTLSVersionGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationTLSVersionGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationTLSVersionGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationTLSVersionGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                            `json:"dataSource,required"`
	Description     string                                                            `json:"description,required"`
	EventType       string                                                            `json:"eventType,required"`
	IsInstantaneous bool                                                              `json:"isInstantaneous,required"`
	EndTime         time.Time                                                         `json:"endTime" format:"date-time"`
	LinkedURL       string                                                            `json:"linkedUrl"`
	StartTime       time.Time                                                         `json:"startTime" format:"date-time"`
	JSON            httpLocationTLSVersionGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpLocationTLSVersionGetResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [HTTPLocationTLSVersionGetResponseMetaConfidenceInfoAnnotation]
type httpLocationTLSVersionGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPLocationTLSVersionGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationTLSVersionGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationTLSVersionGetResponseTop0 struct {
	ClientCountryAlpha2 string                                    `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                    `json:"clientCountryName,required"`
	Value               string                                    `json:"value,required"`
	JSON                httpLocationTLSVersionGetResponseTop0JSON `json:"-"`
}

// httpLocationTLSVersionGetResponseTop0JSON contains the JSON metadata for the
// struct [HTTPLocationTLSVersionGetResponseTop0]
type httpLocationTLSVersionGetResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *HTTPLocationTLSVersionGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationTLSVersionGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPLocationTLSVersionGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPLocationTLSVersionGetParamsBotClass] `query:"botClass"`
	// Filter for browser family.
	BrowserFamily param.Field[[]HTTPLocationTLSVersionGetParamsBrowserFamily] `query:"browserFamily"`
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
	DeviceType param.Field[[]HTTPLocationTLSVersionGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[HTTPLocationTLSVersionGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]HTTPLocationTLSVersionGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]HTTPLocationTLSVersionGetParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]HTTPLocationTLSVersionGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]HTTPLocationTLSVersionGetParamsOS] `query:"os"`
}

// URLQuery serializes [HTTPLocationTLSVersionGetParams]'s query parameters as
// `url.Values`.
func (r HTTPLocationTLSVersionGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// TLS version.
type HTTPLocationTLSVersionGetParamsTLSVersion string

const (
	HTTPLocationTLSVersionGetParamsTLSVersionTlSv1_0  HTTPLocationTLSVersionGetParamsTLSVersion = "TLSv1_0"
	HTTPLocationTLSVersionGetParamsTLSVersionTlSv1_1  HTTPLocationTLSVersionGetParamsTLSVersion = "TLSv1_1"
	HTTPLocationTLSVersionGetParamsTLSVersionTlSv1_2  HTTPLocationTLSVersionGetParamsTLSVersion = "TLSv1_2"
	HTTPLocationTLSVersionGetParamsTLSVersionTlSv1_3  HTTPLocationTLSVersionGetParamsTLSVersion = "TLSv1_3"
	HTTPLocationTLSVersionGetParamsTLSVersionTlSvQuic HTTPLocationTLSVersionGetParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPLocationTLSVersionGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPLocationTLSVersionGetParamsTLSVersionTlSv1_0, HTTPLocationTLSVersionGetParamsTLSVersionTlSv1_1, HTTPLocationTLSVersionGetParamsTLSVersionTlSv1_2, HTTPLocationTLSVersionGetParamsTLSVersionTlSv1_3, HTTPLocationTLSVersionGetParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPLocationTLSVersionGetParamsBotClass string

const (
	HTTPLocationTLSVersionGetParamsBotClassLikelyAutomated HTTPLocationTLSVersionGetParamsBotClass = "LIKELY_AUTOMATED"
	HTTPLocationTLSVersionGetParamsBotClassLikelyHuman     HTTPLocationTLSVersionGetParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPLocationTLSVersionGetParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPLocationTLSVersionGetParamsBotClassLikelyAutomated, HTTPLocationTLSVersionGetParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPLocationTLSVersionGetParamsBrowserFamily string

const (
	HTTPLocationTLSVersionGetParamsBrowserFamilyChrome  HTTPLocationTLSVersionGetParamsBrowserFamily = "CHROME"
	HTTPLocationTLSVersionGetParamsBrowserFamilyEdge    HTTPLocationTLSVersionGetParamsBrowserFamily = "EDGE"
	HTTPLocationTLSVersionGetParamsBrowserFamilyFirefox HTTPLocationTLSVersionGetParamsBrowserFamily = "FIREFOX"
	HTTPLocationTLSVersionGetParamsBrowserFamilySafari  HTTPLocationTLSVersionGetParamsBrowserFamily = "SAFARI"
)

func (r HTTPLocationTLSVersionGetParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPLocationTLSVersionGetParamsBrowserFamilyChrome, HTTPLocationTLSVersionGetParamsBrowserFamilyEdge, HTTPLocationTLSVersionGetParamsBrowserFamilyFirefox, HTTPLocationTLSVersionGetParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPLocationTLSVersionGetParamsDeviceType string

const (
	HTTPLocationTLSVersionGetParamsDeviceTypeDesktop HTTPLocationTLSVersionGetParamsDeviceType = "DESKTOP"
	HTTPLocationTLSVersionGetParamsDeviceTypeMobile  HTTPLocationTLSVersionGetParamsDeviceType = "MOBILE"
	HTTPLocationTLSVersionGetParamsDeviceTypeOther   HTTPLocationTLSVersionGetParamsDeviceType = "OTHER"
)

func (r HTTPLocationTLSVersionGetParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPLocationTLSVersionGetParamsDeviceTypeDesktop, HTTPLocationTLSVersionGetParamsDeviceTypeMobile, HTTPLocationTLSVersionGetParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format results are returned in.
type HTTPLocationTLSVersionGetParamsFormat string

const (
	HTTPLocationTLSVersionGetParamsFormatJson HTTPLocationTLSVersionGetParamsFormat = "JSON"
	HTTPLocationTLSVersionGetParamsFormatCsv  HTTPLocationTLSVersionGetParamsFormat = "CSV"
)

func (r HTTPLocationTLSVersionGetParamsFormat) IsKnown() bool {
	switch r {
	case HTTPLocationTLSVersionGetParamsFormatJson, HTTPLocationTLSVersionGetParamsFormatCsv:
		return true
	}
	return false
}

type HTTPLocationTLSVersionGetParamsHTTPProtocol string

const (
	HTTPLocationTLSVersionGetParamsHTTPProtocolHTTP  HTTPLocationTLSVersionGetParamsHTTPProtocol = "HTTP"
	HTTPLocationTLSVersionGetParamsHTTPProtocolHTTPS HTTPLocationTLSVersionGetParamsHTTPProtocol = "HTTPS"
)

func (r HTTPLocationTLSVersionGetParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPLocationTLSVersionGetParamsHTTPProtocolHTTP, HTTPLocationTLSVersionGetParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPLocationTLSVersionGetParamsHTTPVersion string

const (
	HTTPLocationTLSVersionGetParamsHTTPVersionHttPv1 HTTPLocationTLSVersionGetParamsHTTPVersion = "HTTPv1"
	HTTPLocationTLSVersionGetParamsHTTPVersionHttPv2 HTTPLocationTLSVersionGetParamsHTTPVersion = "HTTPv2"
	HTTPLocationTLSVersionGetParamsHTTPVersionHttPv3 HTTPLocationTLSVersionGetParamsHTTPVersion = "HTTPv3"
)

func (r HTTPLocationTLSVersionGetParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPLocationTLSVersionGetParamsHTTPVersionHttPv1, HTTPLocationTLSVersionGetParamsHTTPVersionHttPv2, HTTPLocationTLSVersionGetParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPLocationTLSVersionGetParamsIPVersion string

const (
	HTTPLocationTLSVersionGetParamsIPVersionIPv4 HTTPLocationTLSVersionGetParamsIPVersion = "IPv4"
	HTTPLocationTLSVersionGetParamsIPVersionIPv6 HTTPLocationTLSVersionGetParamsIPVersion = "IPv6"
)

func (r HTTPLocationTLSVersionGetParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPLocationTLSVersionGetParamsIPVersionIPv4, HTTPLocationTLSVersionGetParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPLocationTLSVersionGetParamsOS string

const (
	HTTPLocationTLSVersionGetParamsOSWindows  HTTPLocationTLSVersionGetParamsOS = "WINDOWS"
	HTTPLocationTLSVersionGetParamsOSMacosx   HTTPLocationTLSVersionGetParamsOS = "MACOSX"
	HTTPLocationTLSVersionGetParamsOSIos      HTTPLocationTLSVersionGetParamsOS = "IOS"
	HTTPLocationTLSVersionGetParamsOSAndroid  HTTPLocationTLSVersionGetParamsOS = "ANDROID"
	HTTPLocationTLSVersionGetParamsOSChromeos HTTPLocationTLSVersionGetParamsOS = "CHROMEOS"
	HTTPLocationTLSVersionGetParamsOSLinux    HTTPLocationTLSVersionGetParamsOS = "LINUX"
	HTTPLocationTLSVersionGetParamsOSSmartTv  HTTPLocationTLSVersionGetParamsOS = "SMART_TV"
)

func (r HTTPLocationTLSVersionGetParamsOS) IsKnown() bool {
	switch r {
	case HTTPLocationTLSVersionGetParamsOSWindows, HTTPLocationTLSVersionGetParamsOSMacosx, HTTPLocationTLSVersionGetParamsOSIos, HTTPLocationTLSVersionGetParamsOSAndroid, HTTPLocationTLSVersionGetParamsOSChromeos, HTTPLocationTLSVersionGetParamsOSLinux, HTTPLocationTLSVersionGetParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPLocationTLSVersionGetResponseEnvelope struct {
	Result  HTTPLocationTLSVersionGetResponse             `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    httpLocationTLSVersionGetResponseEnvelopeJSON `json:"-"`
}

// httpLocationTLSVersionGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [HTTPLocationTLSVersionGetResponseEnvelope]
type httpLocationTLSVersionGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationTLSVersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationTLSVersionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
