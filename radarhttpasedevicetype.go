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

// RadarHTTPAseDeviceTypeService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarHTTPAseDeviceTypeService]
// method instead.
type RadarHTTPAseDeviceTypeService struct {
	Options []option.RequestOption
}

// NewRadarHTTPAseDeviceTypeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarHTTPAseDeviceTypeService(opts ...option.RequestOption) (r *RadarHTTPAseDeviceTypeService) {
	r = &RadarHTTPAseDeviceTypeService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS), by HTTP traffic, of the requested device
// type. Values are a percentage out of the total traffic.
func (r *RadarHTTPAseDeviceTypeService) Get(ctx context.Context, deviceType RadarHTTPAseDeviceTypeGetParamsDeviceType, query RadarHTTPAseDeviceTypeGetParams, opts ...option.RequestOption) (res *RadarHTTPAseDeviceTypeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarHTTPAseDeviceTypeGetResponseEnvelope
	path := fmt.Sprintf("radar/http/top/ases/device_type/%v", deviceType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarHTTPAseDeviceTypeGetResponse struct {
	Meta RadarHTTPAseDeviceTypeGetResponseMeta   `json:"meta,required"`
	Top0 []RadarHTTPAseDeviceTypeGetResponseTop0 `json:"top_0,required"`
	JSON radarHTTPAseDeviceTypeGetResponseJSON   `json:"-"`
}

// radarHTTPAseDeviceTypeGetResponseJSON contains the JSON metadata for the struct
// [RadarHTTPAseDeviceTypeGetResponse]
type radarHTTPAseDeviceTypeGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseDeviceTypeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPAseDeviceTypeGetResponseMeta struct {
	DateRange      []RadarHTTPAseDeviceTypeGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                              `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPAseDeviceTypeGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPAseDeviceTypeGetResponseMetaJSON           `json:"-"`
}

// radarHTTPAseDeviceTypeGetResponseMetaJSON contains the JSON metadata for the
// struct [RadarHTTPAseDeviceTypeGetResponseMeta]
type radarHTTPAseDeviceTypeGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPAseDeviceTypeGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPAseDeviceTypeGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                          `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPAseDeviceTypeGetResponseMetaDateRangeJSON `json:"-"`
}

// radarHTTPAseDeviceTypeGetResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarHTTPAseDeviceTypeGetResponseMetaDateRange]
type radarHTTPAseDeviceTypeGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseDeviceTypeGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPAseDeviceTypeGetResponseMetaConfidenceInfo struct {
	Annotations []RadarHTTPAseDeviceTypeGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                           `json:"level"`
	JSON        radarHTTPAseDeviceTypeGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPAseDeviceTypeGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarHTTPAseDeviceTypeGetResponseMetaConfidenceInfo]
type radarHTTPAseDeviceTypeGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseDeviceTypeGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPAseDeviceTypeGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                            `json:"dataSource,required"`
	Description     string                                                            `json:"description,required"`
	EventType       string                                                            `json:"eventType,required"`
	IsInstantaneous interface{}                                                       `json:"isInstantaneous,required"`
	EndTime         time.Time                                                         `json:"endTime" format:"date-time"`
	LinkedURL       string                                                            `json:"linkedUrl"`
	StartTime       time.Time                                                         `json:"startTime" format:"date-time"`
	JSON            radarHTTPAseDeviceTypeGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPAseDeviceTypeGetResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarHTTPAseDeviceTypeGetResponseMetaConfidenceInfoAnnotation]
type radarHTTPAseDeviceTypeGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPAseDeviceTypeGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPAseDeviceTypeGetResponseTop0 struct {
	ClientASN    int64                                     `json:"clientASN,required"`
	ClientAsName string                                    `json:"clientASName,required"`
	Value        string                                    `json:"value,required"`
	JSON         radarHTTPAseDeviceTypeGetResponseTop0JSON `json:"-"`
}

// radarHTTPAseDeviceTypeGetResponseTop0JSON contains the JSON metadata for the
// struct [RadarHTTPAseDeviceTypeGetResponseTop0]
type radarHTTPAseDeviceTypeGetResponseTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarHTTPAseDeviceTypeGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPAseDeviceTypeGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPAseDeviceTypeGetParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPAseDeviceTypeGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarHTTPAseDeviceTypeGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPAseDeviceTypeGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPAseDeviceTypeGetParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPAseDeviceTypeGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]RadarHTTPAseDeviceTypeGetParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHTTPAseDeviceTypeGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPAseDeviceTypeGetParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPAseDeviceTypeGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Device type.
type RadarHTTPAseDeviceTypeGetParamsDeviceType string

const (
	RadarHTTPAseDeviceTypeGetParamsDeviceTypeDesktop RadarHTTPAseDeviceTypeGetParamsDeviceType = "DESKTOP"
	RadarHTTPAseDeviceTypeGetParamsDeviceTypeMobile  RadarHTTPAseDeviceTypeGetParamsDeviceType = "MOBILE"
	RadarHTTPAseDeviceTypeGetParamsDeviceTypeOther   RadarHTTPAseDeviceTypeGetParamsDeviceType = "OTHER"
)

type RadarHTTPAseDeviceTypeGetParamsBotClass string

const (
	RadarHTTPAseDeviceTypeGetParamsBotClassLikelyAutomated RadarHTTPAseDeviceTypeGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPAseDeviceTypeGetParamsBotClassLikelyHuman     RadarHTTPAseDeviceTypeGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPAseDeviceTypeGetParamsDateRange string

const (
	RadarHTTPAseDeviceTypeGetParamsDateRange1d         RadarHTTPAseDeviceTypeGetParamsDateRange = "1d"
	RadarHTTPAseDeviceTypeGetParamsDateRange2d         RadarHTTPAseDeviceTypeGetParamsDateRange = "2d"
	RadarHTTPAseDeviceTypeGetParamsDateRange7d         RadarHTTPAseDeviceTypeGetParamsDateRange = "7d"
	RadarHTTPAseDeviceTypeGetParamsDateRange14d        RadarHTTPAseDeviceTypeGetParamsDateRange = "14d"
	RadarHTTPAseDeviceTypeGetParamsDateRange28d        RadarHTTPAseDeviceTypeGetParamsDateRange = "28d"
	RadarHTTPAseDeviceTypeGetParamsDateRange12w        RadarHTTPAseDeviceTypeGetParamsDateRange = "12w"
	RadarHTTPAseDeviceTypeGetParamsDateRange24w        RadarHTTPAseDeviceTypeGetParamsDateRange = "24w"
	RadarHTTPAseDeviceTypeGetParamsDateRange52w        RadarHTTPAseDeviceTypeGetParamsDateRange = "52w"
	RadarHTTPAseDeviceTypeGetParamsDateRange1dControl  RadarHTTPAseDeviceTypeGetParamsDateRange = "1dControl"
	RadarHTTPAseDeviceTypeGetParamsDateRange2dControl  RadarHTTPAseDeviceTypeGetParamsDateRange = "2dControl"
	RadarHTTPAseDeviceTypeGetParamsDateRange7dControl  RadarHTTPAseDeviceTypeGetParamsDateRange = "7dControl"
	RadarHTTPAseDeviceTypeGetParamsDateRange14dControl RadarHTTPAseDeviceTypeGetParamsDateRange = "14dControl"
	RadarHTTPAseDeviceTypeGetParamsDateRange28dControl RadarHTTPAseDeviceTypeGetParamsDateRange = "28dControl"
	RadarHTTPAseDeviceTypeGetParamsDateRange12wControl RadarHTTPAseDeviceTypeGetParamsDateRange = "12wControl"
	RadarHTTPAseDeviceTypeGetParamsDateRange24wControl RadarHTTPAseDeviceTypeGetParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarHTTPAseDeviceTypeGetParamsFormat string

const (
	RadarHTTPAseDeviceTypeGetParamsFormatJson RadarHTTPAseDeviceTypeGetParamsFormat = "JSON"
	RadarHTTPAseDeviceTypeGetParamsFormatCsv  RadarHTTPAseDeviceTypeGetParamsFormat = "CSV"
)

type RadarHTTPAseDeviceTypeGetParamsHTTPProtocol string

const (
	RadarHTTPAseDeviceTypeGetParamsHTTPProtocolHTTP  RadarHTTPAseDeviceTypeGetParamsHTTPProtocol = "HTTP"
	RadarHTTPAseDeviceTypeGetParamsHTTPProtocolHTTPS RadarHTTPAseDeviceTypeGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPAseDeviceTypeGetParamsHTTPVersion string

const (
	RadarHTTPAseDeviceTypeGetParamsHTTPVersionHttPv1 RadarHTTPAseDeviceTypeGetParamsHTTPVersion = "HTTPv1"
	RadarHTTPAseDeviceTypeGetParamsHTTPVersionHttPv2 RadarHTTPAseDeviceTypeGetParamsHTTPVersion = "HTTPv2"
	RadarHTTPAseDeviceTypeGetParamsHTTPVersionHttPv3 RadarHTTPAseDeviceTypeGetParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPAseDeviceTypeGetParamsIPVersion string

const (
	RadarHTTPAseDeviceTypeGetParamsIPVersionIPv4 RadarHTTPAseDeviceTypeGetParamsIPVersion = "IPv4"
	RadarHTTPAseDeviceTypeGetParamsIPVersionIPv6 RadarHTTPAseDeviceTypeGetParamsIPVersion = "IPv6"
)

type RadarHTTPAseDeviceTypeGetParamsOS string

const (
	RadarHTTPAseDeviceTypeGetParamsOSWindows  RadarHTTPAseDeviceTypeGetParamsOS = "WINDOWS"
	RadarHTTPAseDeviceTypeGetParamsOSMacosx   RadarHTTPAseDeviceTypeGetParamsOS = "MACOSX"
	RadarHTTPAseDeviceTypeGetParamsOSIos      RadarHTTPAseDeviceTypeGetParamsOS = "IOS"
	RadarHTTPAseDeviceTypeGetParamsOSAndroid  RadarHTTPAseDeviceTypeGetParamsOS = "ANDROID"
	RadarHTTPAseDeviceTypeGetParamsOSChromeos RadarHTTPAseDeviceTypeGetParamsOS = "CHROMEOS"
	RadarHTTPAseDeviceTypeGetParamsOSLinux    RadarHTTPAseDeviceTypeGetParamsOS = "LINUX"
	RadarHTTPAseDeviceTypeGetParamsOSSmartTv  RadarHTTPAseDeviceTypeGetParamsOS = "SMART_TV"
)

type RadarHTTPAseDeviceTypeGetParamsTLSVersion string

const (
	RadarHTTPAseDeviceTypeGetParamsTLSVersionTlSv1_0  RadarHTTPAseDeviceTypeGetParamsTLSVersion = "TLSv1_0"
	RadarHTTPAseDeviceTypeGetParamsTLSVersionTlSv1_1  RadarHTTPAseDeviceTypeGetParamsTLSVersion = "TLSv1_1"
	RadarHTTPAseDeviceTypeGetParamsTLSVersionTlSv1_2  RadarHTTPAseDeviceTypeGetParamsTLSVersion = "TLSv1_2"
	RadarHTTPAseDeviceTypeGetParamsTLSVersionTlSv1_3  RadarHTTPAseDeviceTypeGetParamsTLSVersion = "TLSv1_3"
	RadarHTTPAseDeviceTypeGetParamsTLSVersionTlSvQuic RadarHTTPAseDeviceTypeGetParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPAseDeviceTypeGetResponseEnvelope struct {
	Result  RadarHTTPAseDeviceTypeGetResponse             `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    radarHTTPAseDeviceTypeGetResponseEnvelopeJSON `json:"-"`
}

// radarHTTPAseDeviceTypeGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarHTTPAseDeviceTypeGetResponseEnvelope]
type radarHTTPAseDeviceTypeGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseDeviceTypeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
