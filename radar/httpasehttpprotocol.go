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

// HTTPAseHTTPProtocolService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTTPAseHTTPProtocolService] method instead.
type HTTPAseHTTPProtocolService struct {
	Options []option.RequestOption
}

// NewHTTPAseHTTPProtocolService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewHTTPAseHTTPProtocolService(opts ...option.RequestOption) (r *HTTPAseHTTPProtocolService) {
	r = &HTTPAseHTTPProtocolService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS), by HTTP traffic, of the requested HTTP
// protocol. Values are a percentage out of the total traffic.
func (r *HTTPAseHTTPProtocolService) Get(ctx context.Context, httpProtocol HTTPAseHTTPProtocolGetParamsHTTPProtocol, query HTTPAseHTTPProtocolGetParams, opts ...option.RequestOption) (res *HTTPAseHTTPProtocolGetResponse, err error) {
	var env HTTPAseHTTPProtocolGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/ases/http_protocol/%v", httpProtocol)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPAseHTTPProtocolGetResponse struct {
	Meta HTTPAseHTTPProtocolGetResponseMeta   `json:"meta,required"`
	Top0 []HTTPAseHTTPProtocolGetResponseTop0 `json:"top_0,required"`
	JSON httpAseHTTPProtocolGetResponseJSON   `json:"-"`
}

// httpAseHTTPProtocolGetResponseJSON contains the JSON metadata for the struct
// [HTTPAseHTTPProtocolGetResponse]
type httpAseHTTPProtocolGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseHTTPProtocolGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseHTTPProtocolGetResponseJSON) RawJSON() string {
	return r.raw
}

type HTTPAseHTTPProtocolGetResponseMeta struct {
	DateRange      []HTTPAseHTTPProtocolGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                           `json:"lastUpdated,required"`
	ConfidenceInfo HTTPAseHTTPProtocolGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           httpAseHTTPProtocolGetResponseMetaJSON           `json:"-"`
}

// httpAseHTTPProtocolGetResponseMetaJSON contains the JSON metadata for the struct
// [HTTPAseHTTPProtocolGetResponseMeta]
type httpAseHTTPProtocolGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPAseHTTPProtocolGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseHTTPProtocolGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPAseHTTPProtocolGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      httpAseHTTPProtocolGetResponseMetaDateRangeJSON `json:"-"`
}

// httpAseHTTPProtocolGetResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [HTTPAseHTTPProtocolGetResponseMetaDateRange]
type httpAseHTTPProtocolGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseHTTPProtocolGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseHTTPProtocolGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type HTTPAseHTTPProtocolGetResponseMetaConfidenceInfo struct {
	Annotations []HTTPAseHTTPProtocolGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                        `json:"level"`
	JSON        httpAseHTTPProtocolGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// httpAseHTTPProtocolGetResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [HTTPAseHTTPProtocolGetResponseMetaConfidenceInfo]
type httpAseHTTPProtocolGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseHTTPProtocolGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseHTTPProtocolGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type HTTPAseHTTPProtocolGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                         `json:"dataSource,required"`
	Description     string                                                         `json:"description,required"`
	EventType       string                                                         `json:"eventType,required"`
	IsInstantaneous bool                                                           `json:"isInstantaneous,required"`
	EndTime         time.Time                                                      `json:"endTime" format:"date-time"`
	LinkedURL       string                                                         `json:"linkedUrl"`
	StartTime       time.Time                                                      `json:"startTime" format:"date-time"`
	JSON            httpAseHTTPProtocolGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpAseHTTPProtocolGetResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [HTTPAseHTTPProtocolGetResponseMetaConfidenceInfoAnnotation]
type httpAseHTTPProtocolGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPAseHTTPProtocolGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseHTTPProtocolGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPAseHTTPProtocolGetResponseTop0 struct {
	ClientASN    int64                                  `json:"clientASN,required"`
	ClientAsName string                                 `json:"clientASName,required"`
	Value        string                                 `json:"value,required"`
	JSON         httpAseHTTPProtocolGetResponseTop0JSON `json:"-"`
}

// httpAseHTTPProtocolGetResponseTop0JSON contains the JSON metadata for the struct
// [HTTPAseHTTPProtocolGetResponseTop0]
type httpAseHTTPProtocolGetResponseTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *HTTPAseHTTPProtocolGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseHTTPProtocolGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPAseHTTPProtocolGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPAseHTTPProtocolGetParamsBotClass] `query:"botClass"`
	// Filter for browser family.
	BrowserFamily param.Field[[]HTTPAseHTTPProtocolGetParamsBrowserFamily] `query:"browserFamily"`
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
	DeviceType param.Field[[]HTTPAseHTTPProtocolGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[HTTPAseHTTPProtocolGetParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]HTTPAseHTTPProtocolGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]HTTPAseHTTPProtocolGetParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]HTTPAseHTTPProtocolGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPAseHTTPProtocolGetParams]'s query parameters as
// `url.Values`.
func (r HTTPAseHTTPProtocolGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// HTTP Protocol.
type HTTPAseHTTPProtocolGetParamsHTTPProtocol string

const (
	HTTPAseHTTPProtocolGetParamsHTTPProtocolHTTP  HTTPAseHTTPProtocolGetParamsHTTPProtocol = "HTTP"
	HTTPAseHTTPProtocolGetParamsHTTPProtocolHTTPS HTTPAseHTTPProtocolGetParamsHTTPProtocol = "HTTPS"
)

func (r HTTPAseHTTPProtocolGetParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPAseHTTPProtocolGetParamsHTTPProtocolHTTP, HTTPAseHTTPProtocolGetParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPAseHTTPProtocolGetParamsBotClass string

const (
	HTTPAseHTTPProtocolGetParamsBotClassLikelyAutomated HTTPAseHTTPProtocolGetParamsBotClass = "LIKELY_AUTOMATED"
	HTTPAseHTTPProtocolGetParamsBotClassLikelyHuman     HTTPAseHTTPProtocolGetParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPAseHTTPProtocolGetParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPAseHTTPProtocolGetParamsBotClassLikelyAutomated, HTTPAseHTTPProtocolGetParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPAseHTTPProtocolGetParamsBrowserFamily string

const (
	HTTPAseHTTPProtocolGetParamsBrowserFamilyChrome  HTTPAseHTTPProtocolGetParamsBrowserFamily = "CHROME"
	HTTPAseHTTPProtocolGetParamsBrowserFamilyEdge    HTTPAseHTTPProtocolGetParamsBrowserFamily = "EDGE"
	HTTPAseHTTPProtocolGetParamsBrowserFamilyFirefox HTTPAseHTTPProtocolGetParamsBrowserFamily = "FIREFOX"
	HTTPAseHTTPProtocolGetParamsBrowserFamilySafari  HTTPAseHTTPProtocolGetParamsBrowserFamily = "SAFARI"
)

func (r HTTPAseHTTPProtocolGetParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPAseHTTPProtocolGetParamsBrowserFamilyChrome, HTTPAseHTTPProtocolGetParamsBrowserFamilyEdge, HTTPAseHTTPProtocolGetParamsBrowserFamilyFirefox, HTTPAseHTTPProtocolGetParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPAseHTTPProtocolGetParamsDeviceType string

const (
	HTTPAseHTTPProtocolGetParamsDeviceTypeDesktop HTTPAseHTTPProtocolGetParamsDeviceType = "DESKTOP"
	HTTPAseHTTPProtocolGetParamsDeviceTypeMobile  HTTPAseHTTPProtocolGetParamsDeviceType = "MOBILE"
	HTTPAseHTTPProtocolGetParamsDeviceTypeOther   HTTPAseHTTPProtocolGetParamsDeviceType = "OTHER"
)

func (r HTTPAseHTTPProtocolGetParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPAseHTTPProtocolGetParamsDeviceTypeDesktop, HTTPAseHTTPProtocolGetParamsDeviceTypeMobile, HTTPAseHTTPProtocolGetParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format results are returned in.
type HTTPAseHTTPProtocolGetParamsFormat string

const (
	HTTPAseHTTPProtocolGetParamsFormatJson HTTPAseHTTPProtocolGetParamsFormat = "JSON"
	HTTPAseHTTPProtocolGetParamsFormatCsv  HTTPAseHTTPProtocolGetParamsFormat = "CSV"
)

func (r HTTPAseHTTPProtocolGetParamsFormat) IsKnown() bool {
	switch r {
	case HTTPAseHTTPProtocolGetParamsFormatJson, HTTPAseHTTPProtocolGetParamsFormatCsv:
		return true
	}
	return false
}

type HTTPAseHTTPProtocolGetParamsIPVersion string

const (
	HTTPAseHTTPProtocolGetParamsIPVersionIPv4 HTTPAseHTTPProtocolGetParamsIPVersion = "IPv4"
	HTTPAseHTTPProtocolGetParamsIPVersionIPv6 HTTPAseHTTPProtocolGetParamsIPVersion = "IPv6"
)

func (r HTTPAseHTTPProtocolGetParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPAseHTTPProtocolGetParamsIPVersionIPv4, HTTPAseHTTPProtocolGetParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPAseHTTPProtocolGetParamsOS string

const (
	HTTPAseHTTPProtocolGetParamsOSWindows  HTTPAseHTTPProtocolGetParamsOS = "WINDOWS"
	HTTPAseHTTPProtocolGetParamsOSMacosx   HTTPAseHTTPProtocolGetParamsOS = "MACOSX"
	HTTPAseHTTPProtocolGetParamsOSIos      HTTPAseHTTPProtocolGetParamsOS = "IOS"
	HTTPAseHTTPProtocolGetParamsOSAndroid  HTTPAseHTTPProtocolGetParamsOS = "ANDROID"
	HTTPAseHTTPProtocolGetParamsOSChromeos HTTPAseHTTPProtocolGetParamsOS = "CHROMEOS"
	HTTPAseHTTPProtocolGetParamsOSLinux    HTTPAseHTTPProtocolGetParamsOS = "LINUX"
	HTTPAseHTTPProtocolGetParamsOSSmartTv  HTTPAseHTTPProtocolGetParamsOS = "SMART_TV"
)

func (r HTTPAseHTTPProtocolGetParamsOS) IsKnown() bool {
	switch r {
	case HTTPAseHTTPProtocolGetParamsOSWindows, HTTPAseHTTPProtocolGetParamsOSMacosx, HTTPAseHTTPProtocolGetParamsOSIos, HTTPAseHTTPProtocolGetParamsOSAndroid, HTTPAseHTTPProtocolGetParamsOSChromeos, HTTPAseHTTPProtocolGetParamsOSLinux, HTTPAseHTTPProtocolGetParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPAseHTTPProtocolGetParamsTLSVersion string

const (
	HTTPAseHTTPProtocolGetParamsTLSVersionTlSv1_0  HTTPAseHTTPProtocolGetParamsTLSVersion = "TLSv1_0"
	HTTPAseHTTPProtocolGetParamsTLSVersionTlSv1_1  HTTPAseHTTPProtocolGetParamsTLSVersion = "TLSv1_1"
	HTTPAseHTTPProtocolGetParamsTLSVersionTlSv1_2  HTTPAseHTTPProtocolGetParamsTLSVersion = "TLSv1_2"
	HTTPAseHTTPProtocolGetParamsTLSVersionTlSv1_3  HTTPAseHTTPProtocolGetParamsTLSVersion = "TLSv1_3"
	HTTPAseHTTPProtocolGetParamsTLSVersionTlSvQuic HTTPAseHTTPProtocolGetParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPAseHTTPProtocolGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPAseHTTPProtocolGetParamsTLSVersionTlSv1_0, HTTPAseHTTPProtocolGetParamsTLSVersionTlSv1_1, HTTPAseHTTPProtocolGetParamsTLSVersionTlSv1_2, HTTPAseHTTPProtocolGetParamsTLSVersionTlSv1_3, HTTPAseHTTPProtocolGetParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPAseHTTPProtocolGetResponseEnvelope struct {
	Result  HTTPAseHTTPProtocolGetResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    httpAseHTTPProtocolGetResponseEnvelopeJSON `json:"-"`
}

// httpAseHTTPProtocolGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [HTTPAseHTTPProtocolGetResponseEnvelope]
type httpAseHTTPProtocolGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseHTTPProtocolGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseHTTPProtocolGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
