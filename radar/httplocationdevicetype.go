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

// HTTPLocationDeviceTypeService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTTPLocationDeviceTypeService] method instead.
type HTTPLocationDeviceTypeService struct {
	Options []option.RequestOption
}

// NewHTTPLocationDeviceTypeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewHTTPLocationDeviceTypeService(opts ...option.RequestOption) (r *HTTPLocationDeviceTypeService) {
	r = &HTTPLocationDeviceTypeService{}
	r.Options = opts
	return
}

// Get the top locations, by HTTP traffic, of the requested device type. Values are
// a percentage out of the total traffic.
func (r *HTTPLocationDeviceTypeService) Get(ctx context.Context, deviceType HTTPLocationDeviceTypeGetParamsDeviceType, query HTTPLocationDeviceTypeGetParams, opts ...option.RequestOption) (res *HTTPLocationDeviceTypeGetResponse, err error) {
	var env HTTPLocationDeviceTypeGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/locations/device_type/%v", deviceType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPLocationDeviceTypeGetResponse struct {
	Meta HTTPLocationDeviceTypeGetResponseMeta   `json:"meta,required"`
	Top0 []HTTPLocationDeviceTypeGetResponseTop0 `json:"top_0,required"`
	JSON httpLocationDeviceTypeGetResponseJSON   `json:"-"`
}

// httpLocationDeviceTypeGetResponseJSON contains the JSON metadata for the struct
// [HTTPLocationDeviceTypeGetResponse]
type httpLocationDeviceTypeGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationDeviceTypeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationDeviceTypeGetResponseJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationDeviceTypeGetResponseMeta struct {
	DateRange      []HTTPLocationDeviceTypeGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                              `json:"lastUpdated,required"`
	ConfidenceInfo HTTPLocationDeviceTypeGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           httpLocationDeviceTypeGetResponseMetaJSON           `json:"-"`
}

// httpLocationDeviceTypeGetResponseMetaJSON contains the JSON metadata for the
// struct [HTTPLocationDeviceTypeGetResponseMeta]
type httpLocationDeviceTypeGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPLocationDeviceTypeGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationDeviceTypeGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationDeviceTypeGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                          `json:"startTime,required" format:"date-time"`
	JSON      httpLocationDeviceTypeGetResponseMetaDateRangeJSON `json:"-"`
}

// httpLocationDeviceTypeGetResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [HTTPLocationDeviceTypeGetResponseMetaDateRange]
type httpLocationDeviceTypeGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationDeviceTypeGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationDeviceTypeGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationDeviceTypeGetResponseMetaConfidenceInfo struct {
	Annotations []HTTPLocationDeviceTypeGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                           `json:"level"`
	JSON        httpLocationDeviceTypeGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// httpLocationDeviceTypeGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [HTTPLocationDeviceTypeGetResponseMetaConfidenceInfo]
type httpLocationDeviceTypeGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationDeviceTypeGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationDeviceTypeGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationDeviceTypeGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                            `json:"dataSource,required"`
	Description     string                                                            `json:"description,required"`
	EventType       string                                                            `json:"eventType,required"`
	IsInstantaneous bool                                                              `json:"isInstantaneous,required"`
	EndTime         time.Time                                                         `json:"endTime" format:"date-time"`
	LinkedURL       string                                                            `json:"linkedUrl"`
	StartTime       time.Time                                                         `json:"startTime" format:"date-time"`
	JSON            httpLocationDeviceTypeGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpLocationDeviceTypeGetResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [HTTPLocationDeviceTypeGetResponseMetaConfidenceInfoAnnotation]
type httpLocationDeviceTypeGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPLocationDeviceTypeGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationDeviceTypeGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationDeviceTypeGetResponseTop0 struct {
	ClientCountryAlpha2 string                                    `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                    `json:"clientCountryName,required"`
	Value               string                                    `json:"value,required"`
	JSON                httpLocationDeviceTypeGetResponseTop0JSON `json:"-"`
}

// httpLocationDeviceTypeGetResponseTop0JSON contains the JSON metadata for the
// struct [HTTPLocationDeviceTypeGetResponseTop0]
type httpLocationDeviceTypeGetResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *HTTPLocationDeviceTypeGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationDeviceTypeGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPLocationDeviceTypeGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPLocationDeviceTypeGetParamsBotClass] `query:"botClass"`
	// Filter for browser family.
	BrowserFamily param.Field[[]HTTPLocationDeviceTypeGetParamsBrowserFamily] `query:"browserFamily"`
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
	// Format results are returned in.
	Format param.Field[HTTPLocationDeviceTypeGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]HTTPLocationDeviceTypeGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]HTTPLocationDeviceTypeGetParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]HTTPLocationDeviceTypeGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]HTTPLocationDeviceTypeGetParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]HTTPLocationDeviceTypeGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPLocationDeviceTypeGetParams]'s query parameters as
// `url.Values`.
func (r HTTPLocationDeviceTypeGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Device type.
type HTTPLocationDeviceTypeGetParamsDeviceType string

const (
	HTTPLocationDeviceTypeGetParamsDeviceTypeDesktop HTTPLocationDeviceTypeGetParamsDeviceType = "DESKTOP"
	HTTPLocationDeviceTypeGetParamsDeviceTypeMobile  HTTPLocationDeviceTypeGetParamsDeviceType = "MOBILE"
	HTTPLocationDeviceTypeGetParamsDeviceTypeOther   HTTPLocationDeviceTypeGetParamsDeviceType = "OTHER"
)

func (r HTTPLocationDeviceTypeGetParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPLocationDeviceTypeGetParamsDeviceTypeDesktop, HTTPLocationDeviceTypeGetParamsDeviceTypeMobile, HTTPLocationDeviceTypeGetParamsDeviceTypeOther:
		return true
	}
	return false
}

type HTTPLocationDeviceTypeGetParamsBotClass string

const (
	HTTPLocationDeviceTypeGetParamsBotClassLikelyAutomated HTTPLocationDeviceTypeGetParamsBotClass = "LIKELY_AUTOMATED"
	HTTPLocationDeviceTypeGetParamsBotClassLikelyHuman     HTTPLocationDeviceTypeGetParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPLocationDeviceTypeGetParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPLocationDeviceTypeGetParamsBotClassLikelyAutomated, HTTPLocationDeviceTypeGetParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPLocationDeviceTypeGetParamsBrowserFamily string

const (
	HTTPLocationDeviceTypeGetParamsBrowserFamilyChrome  HTTPLocationDeviceTypeGetParamsBrowserFamily = "CHROME"
	HTTPLocationDeviceTypeGetParamsBrowserFamilyEdge    HTTPLocationDeviceTypeGetParamsBrowserFamily = "EDGE"
	HTTPLocationDeviceTypeGetParamsBrowserFamilyFirefox HTTPLocationDeviceTypeGetParamsBrowserFamily = "FIREFOX"
	HTTPLocationDeviceTypeGetParamsBrowserFamilySafari  HTTPLocationDeviceTypeGetParamsBrowserFamily = "SAFARI"
)

func (r HTTPLocationDeviceTypeGetParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPLocationDeviceTypeGetParamsBrowserFamilyChrome, HTTPLocationDeviceTypeGetParamsBrowserFamilyEdge, HTTPLocationDeviceTypeGetParamsBrowserFamilyFirefox, HTTPLocationDeviceTypeGetParamsBrowserFamilySafari:
		return true
	}
	return false
}

// Format results are returned in.
type HTTPLocationDeviceTypeGetParamsFormat string

const (
	HTTPLocationDeviceTypeGetParamsFormatJson HTTPLocationDeviceTypeGetParamsFormat = "JSON"
	HTTPLocationDeviceTypeGetParamsFormatCsv  HTTPLocationDeviceTypeGetParamsFormat = "CSV"
)

func (r HTTPLocationDeviceTypeGetParamsFormat) IsKnown() bool {
	switch r {
	case HTTPLocationDeviceTypeGetParamsFormatJson, HTTPLocationDeviceTypeGetParamsFormatCsv:
		return true
	}
	return false
}

type HTTPLocationDeviceTypeGetParamsHTTPProtocol string

const (
	HTTPLocationDeviceTypeGetParamsHTTPProtocolHTTP  HTTPLocationDeviceTypeGetParamsHTTPProtocol = "HTTP"
	HTTPLocationDeviceTypeGetParamsHTTPProtocolHTTPS HTTPLocationDeviceTypeGetParamsHTTPProtocol = "HTTPS"
)

func (r HTTPLocationDeviceTypeGetParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPLocationDeviceTypeGetParamsHTTPProtocolHTTP, HTTPLocationDeviceTypeGetParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPLocationDeviceTypeGetParamsHTTPVersion string

const (
	HTTPLocationDeviceTypeGetParamsHTTPVersionHttPv1 HTTPLocationDeviceTypeGetParamsHTTPVersion = "HTTPv1"
	HTTPLocationDeviceTypeGetParamsHTTPVersionHttPv2 HTTPLocationDeviceTypeGetParamsHTTPVersion = "HTTPv2"
	HTTPLocationDeviceTypeGetParamsHTTPVersionHttPv3 HTTPLocationDeviceTypeGetParamsHTTPVersion = "HTTPv3"
)

func (r HTTPLocationDeviceTypeGetParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPLocationDeviceTypeGetParamsHTTPVersionHttPv1, HTTPLocationDeviceTypeGetParamsHTTPVersionHttPv2, HTTPLocationDeviceTypeGetParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPLocationDeviceTypeGetParamsIPVersion string

const (
	HTTPLocationDeviceTypeGetParamsIPVersionIPv4 HTTPLocationDeviceTypeGetParamsIPVersion = "IPv4"
	HTTPLocationDeviceTypeGetParamsIPVersionIPv6 HTTPLocationDeviceTypeGetParamsIPVersion = "IPv6"
)

func (r HTTPLocationDeviceTypeGetParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPLocationDeviceTypeGetParamsIPVersionIPv4, HTTPLocationDeviceTypeGetParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPLocationDeviceTypeGetParamsOS string

const (
	HTTPLocationDeviceTypeGetParamsOSWindows  HTTPLocationDeviceTypeGetParamsOS = "WINDOWS"
	HTTPLocationDeviceTypeGetParamsOSMacosx   HTTPLocationDeviceTypeGetParamsOS = "MACOSX"
	HTTPLocationDeviceTypeGetParamsOSIos      HTTPLocationDeviceTypeGetParamsOS = "IOS"
	HTTPLocationDeviceTypeGetParamsOSAndroid  HTTPLocationDeviceTypeGetParamsOS = "ANDROID"
	HTTPLocationDeviceTypeGetParamsOSChromeos HTTPLocationDeviceTypeGetParamsOS = "CHROMEOS"
	HTTPLocationDeviceTypeGetParamsOSLinux    HTTPLocationDeviceTypeGetParamsOS = "LINUX"
	HTTPLocationDeviceTypeGetParamsOSSmartTv  HTTPLocationDeviceTypeGetParamsOS = "SMART_TV"
)

func (r HTTPLocationDeviceTypeGetParamsOS) IsKnown() bool {
	switch r {
	case HTTPLocationDeviceTypeGetParamsOSWindows, HTTPLocationDeviceTypeGetParamsOSMacosx, HTTPLocationDeviceTypeGetParamsOSIos, HTTPLocationDeviceTypeGetParamsOSAndroid, HTTPLocationDeviceTypeGetParamsOSChromeos, HTTPLocationDeviceTypeGetParamsOSLinux, HTTPLocationDeviceTypeGetParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPLocationDeviceTypeGetParamsTLSVersion string

const (
	HTTPLocationDeviceTypeGetParamsTLSVersionTlSv1_0  HTTPLocationDeviceTypeGetParamsTLSVersion = "TLSv1_0"
	HTTPLocationDeviceTypeGetParamsTLSVersionTlSv1_1  HTTPLocationDeviceTypeGetParamsTLSVersion = "TLSv1_1"
	HTTPLocationDeviceTypeGetParamsTLSVersionTlSv1_2  HTTPLocationDeviceTypeGetParamsTLSVersion = "TLSv1_2"
	HTTPLocationDeviceTypeGetParamsTLSVersionTlSv1_3  HTTPLocationDeviceTypeGetParamsTLSVersion = "TLSv1_3"
	HTTPLocationDeviceTypeGetParamsTLSVersionTlSvQuic HTTPLocationDeviceTypeGetParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPLocationDeviceTypeGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPLocationDeviceTypeGetParamsTLSVersionTlSv1_0, HTTPLocationDeviceTypeGetParamsTLSVersionTlSv1_1, HTTPLocationDeviceTypeGetParamsTLSVersionTlSv1_2, HTTPLocationDeviceTypeGetParamsTLSVersionTlSv1_3, HTTPLocationDeviceTypeGetParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPLocationDeviceTypeGetResponseEnvelope struct {
	Result  HTTPLocationDeviceTypeGetResponse             `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    httpLocationDeviceTypeGetResponseEnvelopeJSON `json:"-"`
}

// httpLocationDeviceTypeGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [HTTPLocationDeviceTypeGetResponseEnvelope]
type httpLocationDeviceTypeGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationDeviceTypeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationDeviceTypeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
