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

// HTTPAseHTTPMethodService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewHTTPAseHTTPMethodService] method
// instead.
type HTTPAseHTTPMethodService struct {
	Options []option.RequestOption
}

// NewHTTPAseHTTPMethodService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewHTTPAseHTTPMethodService(opts ...option.RequestOption) (r *HTTPAseHTTPMethodService) {
	r = &HTTPAseHTTPMethodService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS), by HTTP traffic, of the requested HTTP
// protocol version. Values are a percentage out of the total traffic.
func (r *HTTPAseHTTPMethodService) Get(ctx context.Context, httpVersion HTTPAseHTTPMethodGetParamsHTTPVersion, query HTTPAseHTTPMethodGetParams, opts ...option.RequestOption) (res *HTTPAseHTTPMethodGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HTTPAseHTTPMethodGetResponseEnvelope
	path := fmt.Sprintf("radar/http/top/ases/http_version/%v", httpVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPAseHTTPMethodGetResponse struct {
	Meta HTTPAseHTTPMethodGetResponseMeta   `json:"meta,required"`
	Top0 []HTTPAseHTTPMethodGetResponseTop0 `json:"top_0,required"`
	JSON httpAseHTTPMethodGetResponseJSON   `json:"-"`
}

// httpAseHTTPMethodGetResponseJSON contains the JSON metadata for the struct
// [HTTPAseHTTPMethodGetResponse]
type httpAseHTTPMethodGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseHTTPMethodGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseHTTPMethodGetResponseJSON) RawJSON() string {
	return r.raw
}

type HTTPAseHTTPMethodGetResponseMeta struct {
	DateRange      []UnnamedSchemaRef175                          `json:"dateRange,required"`
	LastUpdated    string                                         `json:"lastUpdated,required"`
	ConfidenceInfo HTTPAseHTTPMethodGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           httpAseHTTPMethodGetResponseMetaJSON           `json:"-"`
}

// httpAseHTTPMethodGetResponseMetaJSON contains the JSON metadata for the struct
// [HTTPAseHTTPMethodGetResponseMeta]
type httpAseHTTPMethodGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPAseHTTPMethodGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseHTTPMethodGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPAseHTTPMethodGetResponseMetaConfidenceInfo struct {
	Annotations []UnnamedSchemaRef174                              `json:"annotations"`
	Level       int64                                              `json:"level"`
	JSON        httpAseHTTPMethodGetResponseMetaConfidenceInfoJSON `json:"-"`
}

// httpAseHTTPMethodGetResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [HTTPAseHTTPMethodGetResponseMetaConfidenceInfo]
type httpAseHTTPMethodGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseHTTPMethodGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseHTTPMethodGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type HTTPAseHTTPMethodGetResponseTop0 struct {
	ClientASN    int64                                `json:"clientASN,required"`
	ClientAsName string                               `json:"clientASName,required"`
	Value        string                               `json:"value,required"`
	JSON         httpAseHTTPMethodGetResponseTop0JSON `json:"-"`
}

// httpAseHTTPMethodGetResponseTop0JSON contains the JSON metadata for the struct
// [HTTPAseHTTPMethodGetResponseTop0]
type httpAseHTTPMethodGetResponseTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *HTTPAseHTTPMethodGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseHTTPMethodGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPAseHTTPMethodGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPAseHTTPMethodGetParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]HTTPAseHTTPMethodGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]HTTPAseHTTPMethodGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[HTTPAseHTTPMethodGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]HTTPAseHTTPMethodGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for ip version.
	IPVersion param.Field[[]HTTPAseHTTPMethodGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]HTTPAseHTTPMethodGetParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]HTTPAseHTTPMethodGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPAseHTTPMethodGetParams]'s query parameters as
// `url.Values`.
func (r HTTPAseHTTPMethodGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// HTTP version.
type HTTPAseHTTPMethodGetParamsHTTPVersion string

const (
	HTTPAseHTTPMethodGetParamsHTTPVersionHttPv1 HTTPAseHTTPMethodGetParamsHTTPVersion = "HTTPv1"
	HTTPAseHTTPMethodGetParamsHTTPVersionHttPv2 HTTPAseHTTPMethodGetParamsHTTPVersion = "HTTPv2"
	HTTPAseHTTPMethodGetParamsHTTPVersionHttPv3 HTTPAseHTTPMethodGetParamsHTTPVersion = "HTTPv3"
)

func (r HTTPAseHTTPMethodGetParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPAseHTTPMethodGetParamsHTTPVersionHttPv1, HTTPAseHTTPMethodGetParamsHTTPVersionHttPv2, HTTPAseHTTPMethodGetParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPAseHTTPMethodGetParamsBotClass string

const (
	HTTPAseHTTPMethodGetParamsBotClassLikelyAutomated HTTPAseHTTPMethodGetParamsBotClass = "LIKELY_AUTOMATED"
	HTTPAseHTTPMethodGetParamsBotClassLikelyHuman     HTTPAseHTTPMethodGetParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPAseHTTPMethodGetParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPAseHTTPMethodGetParamsBotClassLikelyAutomated, HTTPAseHTTPMethodGetParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPAseHTTPMethodGetParamsDateRange string

const (
	HTTPAseHTTPMethodGetParamsDateRange1d         HTTPAseHTTPMethodGetParamsDateRange = "1d"
	HTTPAseHTTPMethodGetParamsDateRange2d         HTTPAseHTTPMethodGetParamsDateRange = "2d"
	HTTPAseHTTPMethodGetParamsDateRange7d         HTTPAseHTTPMethodGetParamsDateRange = "7d"
	HTTPAseHTTPMethodGetParamsDateRange14d        HTTPAseHTTPMethodGetParamsDateRange = "14d"
	HTTPAseHTTPMethodGetParamsDateRange28d        HTTPAseHTTPMethodGetParamsDateRange = "28d"
	HTTPAseHTTPMethodGetParamsDateRange12w        HTTPAseHTTPMethodGetParamsDateRange = "12w"
	HTTPAseHTTPMethodGetParamsDateRange24w        HTTPAseHTTPMethodGetParamsDateRange = "24w"
	HTTPAseHTTPMethodGetParamsDateRange52w        HTTPAseHTTPMethodGetParamsDateRange = "52w"
	HTTPAseHTTPMethodGetParamsDateRange1dControl  HTTPAseHTTPMethodGetParamsDateRange = "1dControl"
	HTTPAseHTTPMethodGetParamsDateRange2dControl  HTTPAseHTTPMethodGetParamsDateRange = "2dControl"
	HTTPAseHTTPMethodGetParamsDateRange7dControl  HTTPAseHTTPMethodGetParamsDateRange = "7dControl"
	HTTPAseHTTPMethodGetParamsDateRange14dControl HTTPAseHTTPMethodGetParamsDateRange = "14dControl"
	HTTPAseHTTPMethodGetParamsDateRange28dControl HTTPAseHTTPMethodGetParamsDateRange = "28dControl"
	HTTPAseHTTPMethodGetParamsDateRange12wControl HTTPAseHTTPMethodGetParamsDateRange = "12wControl"
	HTTPAseHTTPMethodGetParamsDateRange24wControl HTTPAseHTTPMethodGetParamsDateRange = "24wControl"
)

func (r HTTPAseHTTPMethodGetParamsDateRange) IsKnown() bool {
	switch r {
	case HTTPAseHTTPMethodGetParamsDateRange1d, HTTPAseHTTPMethodGetParamsDateRange2d, HTTPAseHTTPMethodGetParamsDateRange7d, HTTPAseHTTPMethodGetParamsDateRange14d, HTTPAseHTTPMethodGetParamsDateRange28d, HTTPAseHTTPMethodGetParamsDateRange12w, HTTPAseHTTPMethodGetParamsDateRange24w, HTTPAseHTTPMethodGetParamsDateRange52w, HTTPAseHTTPMethodGetParamsDateRange1dControl, HTTPAseHTTPMethodGetParamsDateRange2dControl, HTTPAseHTTPMethodGetParamsDateRange7dControl, HTTPAseHTTPMethodGetParamsDateRange14dControl, HTTPAseHTTPMethodGetParamsDateRange28dControl, HTTPAseHTTPMethodGetParamsDateRange12wControl, HTTPAseHTTPMethodGetParamsDateRange24wControl:
		return true
	}
	return false
}

type HTTPAseHTTPMethodGetParamsDeviceType string

const (
	HTTPAseHTTPMethodGetParamsDeviceTypeDesktop HTTPAseHTTPMethodGetParamsDeviceType = "DESKTOP"
	HTTPAseHTTPMethodGetParamsDeviceTypeMobile  HTTPAseHTTPMethodGetParamsDeviceType = "MOBILE"
	HTTPAseHTTPMethodGetParamsDeviceTypeOther   HTTPAseHTTPMethodGetParamsDeviceType = "OTHER"
)

func (r HTTPAseHTTPMethodGetParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPAseHTTPMethodGetParamsDeviceTypeDesktop, HTTPAseHTTPMethodGetParamsDeviceTypeMobile, HTTPAseHTTPMethodGetParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format results are returned in.
type HTTPAseHTTPMethodGetParamsFormat string

const (
	HTTPAseHTTPMethodGetParamsFormatJson HTTPAseHTTPMethodGetParamsFormat = "JSON"
	HTTPAseHTTPMethodGetParamsFormatCsv  HTTPAseHTTPMethodGetParamsFormat = "CSV"
)

func (r HTTPAseHTTPMethodGetParamsFormat) IsKnown() bool {
	switch r {
	case HTTPAseHTTPMethodGetParamsFormatJson, HTTPAseHTTPMethodGetParamsFormatCsv:
		return true
	}
	return false
}

type HTTPAseHTTPMethodGetParamsHTTPProtocol string

const (
	HTTPAseHTTPMethodGetParamsHTTPProtocolHTTP  HTTPAseHTTPMethodGetParamsHTTPProtocol = "HTTP"
	HTTPAseHTTPMethodGetParamsHTTPProtocolHTTPS HTTPAseHTTPMethodGetParamsHTTPProtocol = "HTTPS"
)

func (r HTTPAseHTTPMethodGetParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPAseHTTPMethodGetParamsHTTPProtocolHTTP, HTTPAseHTTPMethodGetParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPAseHTTPMethodGetParamsIPVersion string

const (
	HTTPAseHTTPMethodGetParamsIPVersionIPv4 HTTPAseHTTPMethodGetParamsIPVersion = "IPv4"
	HTTPAseHTTPMethodGetParamsIPVersionIPv6 HTTPAseHTTPMethodGetParamsIPVersion = "IPv6"
)

func (r HTTPAseHTTPMethodGetParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPAseHTTPMethodGetParamsIPVersionIPv4, HTTPAseHTTPMethodGetParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPAseHTTPMethodGetParamsOS string

const (
	HTTPAseHTTPMethodGetParamsOSWindows  HTTPAseHTTPMethodGetParamsOS = "WINDOWS"
	HTTPAseHTTPMethodGetParamsOSMacosx   HTTPAseHTTPMethodGetParamsOS = "MACOSX"
	HTTPAseHTTPMethodGetParamsOSIos      HTTPAseHTTPMethodGetParamsOS = "IOS"
	HTTPAseHTTPMethodGetParamsOSAndroid  HTTPAseHTTPMethodGetParamsOS = "ANDROID"
	HTTPAseHTTPMethodGetParamsOSChromeos HTTPAseHTTPMethodGetParamsOS = "CHROMEOS"
	HTTPAseHTTPMethodGetParamsOSLinux    HTTPAseHTTPMethodGetParamsOS = "LINUX"
	HTTPAseHTTPMethodGetParamsOSSmartTv  HTTPAseHTTPMethodGetParamsOS = "SMART_TV"
)

func (r HTTPAseHTTPMethodGetParamsOS) IsKnown() bool {
	switch r {
	case HTTPAseHTTPMethodGetParamsOSWindows, HTTPAseHTTPMethodGetParamsOSMacosx, HTTPAseHTTPMethodGetParamsOSIos, HTTPAseHTTPMethodGetParamsOSAndroid, HTTPAseHTTPMethodGetParamsOSChromeos, HTTPAseHTTPMethodGetParamsOSLinux, HTTPAseHTTPMethodGetParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPAseHTTPMethodGetParamsTLSVersion string

const (
	HTTPAseHTTPMethodGetParamsTLSVersionTlSv1_0  HTTPAseHTTPMethodGetParamsTLSVersion = "TLSv1_0"
	HTTPAseHTTPMethodGetParamsTLSVersionTlSv1_1  HTTPAseHTTPMethodGetParamsTLSVersion = "TLSv1_1"
	HTTPAseHTTPMethodGetParamsTLSVersionTlSv1_2  HTTPAseHTTPMethodGetParamsTLSVersion = "TLSv1_2"
	HTTPAseHTTPMethodGetParamsTLSVersionTlSv1_3  HTTPAseHTTPMethodGetParamsTLSVersion = "TLSv1_3"
	HTTPAseHTTPMethodGetParamsTLSVersionTlSvQuic HTTPAseHTTPMethodGetParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPAseHTTPMethodGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPAseHTTPMethodGetParamsTLSVersionTlSv1_0, HTTPAseHTTPMethodGetParamsTLSVersionTlSv1_1, HTTPAseHTTPMethodGetParamsTLSVersionTlSv1_2, HTTPAseHTTPMethodGetParamsTLSVersionTlSv1_3, HTTPAseHTTPMethodGetParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPAseHTTPMethodGetResponseEnvelope struct {
	Result  HTTPAseHTTPMethodGetResponse             `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    httpAseHTTPMethodGetResponseEnvelopeJSON `json:"-"`
}

// httpAseHTTPMethodGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [HTTPAseHTTPMethodGetResponseEnvelope]
type httpAseHTTPMethodGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseHTTPMethodGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseHTTPMethodGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
