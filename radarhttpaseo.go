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

// RadarHTTPAseOService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarHTTPAseOService] method
// instead.
type RadarHTTPAseOService struct {
	Options []option.RequestOption
}

// NewRadarHTTPAseOService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarHTTPAseOService(opts ...option.RequestOption) (r *RadarHTTPAseOService) {
	r = &RadarHTTPAseOService{}
	r.Options = opts
	return
}

// Get the top autonomous systems, by HTTP traffic, of the requested operating
// systems. Values are a percentage out of the total traffic.
func (r *RadarHTTPAseOService) Get(ctx context.Context, os RadarHTTPAseOGetParamsOs, query RadarHTTPAseOGetParams, opts ...option.RequestOption) (res *RadarHTTPAseOGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarHTTPAseOGetResponseEnvelope
	path := fmt.Sprintf("radar/http/top/ases/os/%v", os)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarHTTPAseOGetResponse struct {
	Meta RadarHTTPAseOGetResponseMeta   `json:"meta,required"`
	Top0 []RadarHTTPAseOGetResponseTop0 `json:"top_0,required"`
	JSON radarHTTPAseOGetResponseJSON   `json:"-"`
}

// radarHTTPAseOGetResponseJSON contains the JSON metadata for the struct
// [RadarHTTPAseOGetResponse]
type radarHTTPAseOGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseOGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPAseOGetResponseMeta struct {
	DateRange      []RadarHTTPAseOGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                     `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPAseOGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPAseOGetResponseMetaJSON           `json:"-"`
}

// radarHTTPAseOGetResponseMetaJSON contains the JSON metadata for the struct
// [RadarHTTPAseOGetResponseMeta]
type radarHTTPAseOGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPAseOGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPAseOGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                 `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPAseOGetResponseMetaDateRangeJSON `json:"-"`
}

// radarHTTPAseOGetResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [RadarHTTPAseOGetResponseMetaDateRange]
type radarHTTPAseOGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseOGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPAseOGetResponseMetaConfidenceInfo struct {
	Annotations []RadarHTTPAseOGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                  `json:"level"`
	JSON        radarHTTPAseOGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPAseOGetResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [RadarHTTPAseOGetResponseMetaConfidenceInfo]
type radarHTTPAseOGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseOGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPAseOGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                   `json:"dataSource,required"`
	Description     string                                                   `json:"description,required"`
	EventType       string                                                   `json:"eventType,required"`
	IsInstantaneous interface{}                                              `json:"isInstantaneous,required"`
	EndTime         time.Time                                                `json:"endTime" format:"date-time"`
	LinkedURL       string                                                   `json:"linkedUrl"`
	StartTime       time.Time                                                `json:"startTime" format:"date-time"`
	JSON            radarHTTPAseOGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPAseOGetResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [RadarHTTPAseOGetResponseMetaConfidenceInfoAnnotation]
type radarHTTPAseOGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPAseOGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPAseOGetResponseTop0 struct {
	ClientAsn    int64                            `json:"clientASN,required"`
	ClientAsName string                           `json:"clientASName,required"`
	Value        string                           `json:"value,required"`
	JSON         radarHTTPAseOGetResponseTop0JSON `json:"-"`
}

// radarHTTPAseOGetResponseTop0JSON contains the JSON metadata for the struct
// [RadarHTTPAseOGetResponseTop0]
type radarHTTPAseOGetResponseTop0JSON struct {
	ClientAsn    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarHTTPAseOGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPAseOGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPAseOGetParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPAseOGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPAseOGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPAseOGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPAseOGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPAseOGetParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPAseOGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHTTPAseOGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPAseOGetParams]'s query parameters as `url.Values`.
func (r RadarHTTPAseOGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// IP version.
type RadarHTTPAseOGetParamsOs string

const (
	RadarHTTPAseOGetParamsOsWindows  RadarHTTPAseOGetParamsOs = "WINDOWS"
	RadarHTTPAseOGetParamsOsMacosx   RadarHTTPAseOGetParamsOs = "MACOSX"
	RadarHTTPAseOGetParamsOsIos      RadarHTTPAseOGetParamsOs = "IOS"
	RadarHTTPAseOGetParamsOsAndroid  RadarHTTPAseOGetParamsOs = "ANDROID"
	RadarHTTPAseOGetParamsOsChromeos RadarHTTPAseOGetParamsOs = "CHROMEOS"
	RadarHTTPAseOGetParamsOsLinux    RadarHTTPAseOGetParamsOs = "LINUX"
	RadarHTTPAseOGetParamsOsSmartTv  RadarHTTPAseOGetParamsOs = "SMART_TV"
)

type RadarHTTPAseOGetParamsBotClass string

const (
	RadarHTTPAseOGetParamsBotClassLikelyAutomated RadarHTTPAseOGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPAseOGetParamsBotClassLikelyHuman     RadarHTTPAseOGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPAseOGetParamsDateRange string

const (
	RadarHTTPAseOGetParamsDateRange1d         RadarHTTPAseOGetParamsDateRange = "1d"
	RadarHTTPAseOGetParamsDateRange2d         RadarHTTPAseOGetParamsDateRange = "2d"
	RadarHTTPAseOGetParamsDateRange7d         RadarHTTPAseOGetParamsDateRange = "7d"
	RadarHTTPAseOGetParamsDateRange14d        RadarHTTPAseOGetParamsDateRange = "14d"
	RadarHTTPAseOGetParamsDateRange28d        RadarHTTPAseOGetParamsDateRange = "28d"
	RadarHTTPAseOGetParamsDateRange12w        RadarHTTPAseOGetParamsDateRange = "12w"
	RadarHTTPAseOGetParamsDateRange24w        RadarHTTPAseOGetParamsDateRange = "24w"
	RadarHTTPAseOGetParamsDateRange52w        RadarHTTPAseOGetParamsDateRange = "52w"
	RadarHTTPAseOGetParamsDateRange1dControl  RadarHTTPAseOGetParamsDateRange = "1dControl"
	RadarHTTPAseOGetParamsDateRange2dControl  RadarHTTPAseOGetParamsDateRange = "2dControl"
	RadarHTTPAseOGetParamsDateRange7dControl  RadarHTTPAseOGetParamsDateRange = "7dControl"
	RadarHTTPAseOGetParamsDateRange14dControl RadarHTTPAseOGetParamsDateRange = "14dControl"
	RadarHTTPAseOGetParamsDateRange28dControl RadarHTTPAseOGetParamsDateRange = "28dControl"
	RadarHTTPAseOGetParamsDateRange12wControl RadarHTTPAseOGetParamsDateRange = "12wControl"
	RadarHTTPAseOGetParamsDateRange24wControl RadarHTTPAseOGetParamsDateRange = "24wControl"
)

type RadarHTTPAseOGetParamsDeviceType string

const (
	RadarHTTPAseOGetParamsDeviceTypeDesktop RadarHTTPAseOGetParamsDeviceType = "DESKTOP"
	RadarHTTPAseOGetParamsDeviceTypeMobile  RadarHTTPAseOGetParamsDeviceType = "MOBILE"
	RadarHTTPAseOGetParamsDeviceTypeOther   RadarHTTPAseOGetParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPAseOGetParamsFormat string

const (
	RadarHTTPAseOGetParamsFormatJson RadarHTTPAseOGetParamsFormat = "JSON"
	RadarHTTPAseOGetParamsFormatCsv  RadarHTTPAseOGetParamsFormat = "CSV"
)

type RadarHTTPAseOGetParamsHTTPProtocol string

const (
	RadarHTTPAseOGetParamsHTTPProtocolHTTP  RadarHTTPAseOGetParamsHTTPProtocol = "HTTP"
	RadarHTTPAseOGetParamsHTTPProtocolHTTPS RadarHTTPAseOGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPAseOGetParamsHTTPVersion string

const (
	RadarHTTPAseOGetParamsHTTPVersionHttPv1 RadarHTTPAseOGetParamsHTTPVersion = "HTTPv1"
	RadarHTTPAseOGetParamsHTTPVersionHttPv2 RadarHTTPAseOGetParamsHTTPVersion = "HTTPv2"
	RadarHTTPAseOGetParamsHTTPVersionHttPv3 RadarHTTPAseOGetParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPAseOGetParamsIPVersion string

const (
	RadarHTTPAseOGetParamsIPVersionIPv4 RadarHTTPAseOGetParamsIPVersion = "IPv4"
	RadarHTTPAseOGetParamsIPVersionIPv6 RadarHTTPAseOGetParamsIPVersion = "IPv6"
)

type RadarHTTPAseOGetParamsTLSVersion string

const (
	RadarHTTPAseOGetParamsTLSVersionTlSv1_0  RadarHTTPAseOGetParamsTLSVersion = "TLSv1_0"
	RadarHTTPAseOGetParamsTLSVersionTlSv1_1  RadarHTTPAseOGetParamsTLSVersion = "TLSv1_1"
	RadarHTTPAseOGetParamsTLSVersionTlSv1_2  RadarHTTPAseOGetParamsTLSVersion = "TLSv1_2"
	RadarHTTPAseOGetParamsTLSVersionTlSv1_3  RadarHTTPAseOGetParamsTLSVersion = "TLSv1_3"
	RadarHTTPAseOGetParamsTLSVersionTlSvQuic RadarHTTPAseOGetParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPAseOGetResponseEnvelope struct {
	Result  RadarHTTPAseOGetResponse             `json:"result,required"`
	Success bool                                 `json:"success,required"`
	JSON    radarHTTPAseOGetResponseEnvelopeJSON `json:"-"`
}

// radarHTTPAseOGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RadarHTTPAseOGetResponseEnvelope]
type radarHTTPAseOGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseOGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
