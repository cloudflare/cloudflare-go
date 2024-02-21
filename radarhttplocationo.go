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

// RadarHTTPLocationOService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarHTTPLocationOService] method
// instead.
type RadarHTTPLocationOService struct {
	Options []option.RequestOption
}

// NewRadarHTTPLocationOService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarHTTPLocationOService(opts ...option.RequestOption) (r *RadarHTTPLocationOService) {
	r = &RadarHTTPLocationOService{}
	r.Options = opts
	return
}

// Get the top locations, by HTTP traffic, of the requested operating systems.
// Values are a percentage out of the total traffic.
func (r *RadarHTTPLocationOService) Get(ctx context.Context, os RadarHTTPLocationOGetParamsOs, query RadarHTTPLocationOGetParams, opts ...option.RequestOption) (res *RadarHTTPLocationOGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarHTTPLocationOGetResponseEnvelope
	path := fmt.Sprintf("radar/http/top/locations/os/%v", os)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarHTTPLocationOGetResponse struct {
	Meta RadarHTTPLocationOGetResponseMeta   `json:"meta,required"`
	Top0 []RadarHTTPLocationOGetResponseTop0 `json:"top_0,required"`
	JSON radarHTTPLocationOGetResponseJSON   `json:"-"`
}

// radarHTTPLocationOGetResponseJSON contains the JSON metadata for the struct
// [RadarHTTPLocationOGetResponse]
type radarHTTPLocationOGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationOGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPLocationOGetResponseMeta struct {
	DateRange      []RadarHTTPLocationOGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                          `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPLocationOGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPLocationOGetResponseMetaJSON           `json:"-"`
}

// radarHTTPLocationOGetResponseMetaJSON contains the JSON metadata for the struct
// [RadarHTTPLocationOGetResponseMeta]
type radarHTTPLocationOGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPLocationOGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPLocationOGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                      `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPLocationOGetResponseMetaDateRangeJSON `json:"-"`
}

// radarHTTPLocationOGetResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [RadarHTTPLocationOGetResponseMetaDateRange]
type radarHTTPLocationOGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationOGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPLocationOGetResponseMetaConfidenceInfo struct {
	Annotations []RadarHTTPLocationOGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                       `json:"level"`
	JSON        radarHTTPLocationOGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPLocationOGetResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [RadarHTTPLocationOGetResponseMetaConfidenceInfo]
type radarHTTPLocationOGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationOGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPLocationOGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                        `json:"dataSource,required"`
	Description     string                                                        `json:"description,required"`
	EventType       string                                                        `json:"eventType,required"`
	IsInstantaneous interface{}                                                   `json:"isInstantaneous,required"`
	EndTime         time.Time                                                     `json:"endTime" format:"date-time"`
	LinkedURL       string                                                        `json:"linkedUrl"`
	StartTime       time.Time                                                     `json:"startTime" format:"date-time"`
	JSON            radarHTTPLocationOGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPLocationOGetResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [RadarHTTPLocationOGetResponseMetaConfidenceInfoAnnotation]
type radarHTTPLocationOGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPLocationOGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPLocationOGetResponseTop0 struct {
	ClientCountryAlpha2 string                                `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                `json:"clientCountryName,required"`
	Value               string                                `json:"value,required"`
	JSON                radarHTTPLocationOGetResponseTop0JSON `json:"-"`
}

// radarHTTPLocationOGetResponseTop0JSON contains the JSON metadata for the struct
// [RadarHTTPLocationOGetResponseTop0]
type radarHTTPLocationOGetResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarHTTPLocationOGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPLocationOGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPLocationOGetParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPLocationOGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPLocationOGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPLocationOGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPLocationOGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPLocationOGetParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPLocationOGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHTTPLocationOGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPLocationOGetParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPLocationOGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// IP version.
type RadarHTTPLocationOGetParamsOs string

const (
	RadarHTTPLocationOGetParamsOsWindows  RadarHTTPLocationOGetParamsOs = "WINDOWS"
	RadarHTTPLocationOGetParamsOsMacosx   RadarHTTPLocationOGetParamsOs = "MACOSX"
	RadarHTTPLocationOGetParamsOsIos      RadarHTTPLocationOGetParamsOs = "IOS"
	RadarHTTPLocationOGetParamsOsAndroid  RadarHTTPLocationOGetParamsOs = "ANDROID"
	RadarHTTPLocationOGetParamsOsChromeos RadarHTTPLocationOGetParamsOs = "CHROMEOS"
	RadarHTTPLocationOGetParamsOsLinux    RadarHTTPLocationOGetParamsOs = "LINUX"
	RadarHTTPLocationOGetParamsOsSmartTv  RadarHTTPLocationOGetParamsOs = "SMART_TV"
)

type RadarHTTPLocationOGetParamsBotClass string

const (
	RadarHTTPLocationOGetParamsBotClassLikelyAutomated RadarHTTPLocationOGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPLocationOGetParamsBotClassLikelyHuman     RadarHTTPLocationOGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPLocationOGetParamsDateRange string

const (
	RadarHTTPLocationOGetParamsDateRange1d         RadarHTTPLocationOGetParamsDateRange = "1d"
	RadarHTTPLocationOGetParamsDateRange2d         RadarHTTPLocationOGetParamsDateRange = "2d"
	RadarHTTPLocationOGetParamsDateRange7d         RadarHTTPLocationOGetParamsDateRange = "7d"
	RadarHTTPLocationOGetParamsDateRange14d        RadarHTTPLocationOGetParamsDateRange = "14d"
	RadarHTTPLocationOGetParamsDateRange28d        RadarHTTPLocationOGetParamsDateRange = "28d"
	RadarHTTPLocationOGetParamsDateRange12w        RadarHTTPLocationOGetParamsDateRange = "12w"
	RadarHTTPLocationOGetParamsDateRange24w        RadarHTTPLocationOGetParamsDateRange = "24w"
	RadarHTTPLocationOGetParamsDateRange52w        RadarHTTPLocationOGetParamsDateRange = "52w"
	RadarHTTPLocationOGetParamsDateRange1dControl  RadarHTTPLocationOGetParamsDateRange = "1dControl"
	RadarHTTPLocationOGetParamsDateRange2dControl  RadarHTTPLocationOGetParamsDateRange = "2dControl"
	RadarHTTPLocationOGetParamsDateRange7dControl  RadarHTTPLocationOGetParamsDateRange = "7dControl"
	RadarHTTPLocationOGetParamsDateRange14dControl RadarHTTPLocationOGetParamsDateRange = "14dControl"
	RadarHTTPLocationOGetParamsDateRange28dControl RadarHTTPLocationOGetParamsDateRange = "28dControl"
	RadarHTTPLocationOGetParamsDateRange12wControl RadarHTTPLocationOGetParamsDateRange = "12wControl"
	RadarHTTPLocationOGetParamsDateRange24wControl RadarHTTPLocationOGetParamsDateRange = "24wControl"
)

type RadarHTTPLocationOGetParamsDeviceType string

const (
	RadarHTTPLocationOGetParamsDeviceTypeDesktop RadarHTTPLocationOGetParamsDeviceType = "DESKTOP"
	RadarHTTPLocationOGetParamsDeviceTypeMobile  RadarHTTPLocationOGetParamsDeviceType = "MOBILE"
	RadarHTTPLocationOGetParamsDeviceTypeOther   RadarHTTPLocationOGetParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPLocationOGetParamsFormat string

const (
	RadarHTTPLocationOGetParamsFormatJson RadarHTTPLocationOGetParamsFormat = "JSON"
	RadarHTTPLocationOGetParamsFormatCsv  RadarHTTPLocationOGetParamsFormat = "CSV"
)

type RadarHTTPLocationOGetParamsHTTPProtocol string

const (
	RadarHTTPLocationOGetParamsHTTPProtocolHTTP  RadarHTTPLocationOGetParamsHTTPProtocol = "HTTP"
	RadarHTTPLocationOGetParamsHTTPProtocolHTTPS RadarHTTPLocationOGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPLocationOGetParamsHTTPVersion string

const (
	RadarHTTPLocationOGetParamsHTTPVersionHttPv1 RadarHTTPLocationOGetParamsHTTPVersion = "HTTPv1"
	RadarHTTPLocationOGetParamsHTTPVersionHttPv2 RadarHTTPLocationOGetParamsHTTPVersion = "HTTPv2"
	RadarHTTPLocationOGetParamsHTTPVersionHttPv3 RadarHTTPLocationOGetParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPLocationOGetParamsIPVersion string

const (
	RadarHTTPLocationOGetParamsIPVersionIPv4 RadarHTTPLocationOGetParamsIPVersion = "IPv4"
	RadarHTTPLocationOGetParamsIPVersionIPv6 RadarHTTPLocationOGetParamsIPVersion = "IPv6"
)

type RadarHTTPLocationOGetParamsTLSVersion string

const (
	RadarHTTPLocationOGetParamsTLSVersionTlSv1_0  RadarHTTPLocationOGetParamsTLSVersion = "TLSv1_0"
	RadarHTTPLocationOGetParamsTLSVersionTlSv1_1  RadarHTTPLocationOGetParamsTLSVersion = "TLSv1_1"
	RadarHTTPLocationOGetParamsTLSVersionTlSv1_2  RadarHTTPLocationOGetParamsTLSVersion = "TLSv1_2"
	RadarHTTPLocationOGetParamsTLSVersionTlSv1_3  RadarHTTPLocationOGetParamsTLSVersion = "TLSv1_3"
	RadarHTTPLocationOGetParamsTLSVersionTlSvQuic RadarHTTPLocationOGetParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPLocationOGetResponseEnvelope struct {
	Result  RadarHTTPLocationOGetResponse             `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    radarHTTPLocationOGetResponseEnvelopeJSON `json:"-"`
}

// radarHTTPLocationOGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarHTTPLocationOGetResponseEnvelope]
type radarHTTPLocationOGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationOGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
