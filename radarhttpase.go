// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarHTTPAseService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarHTTPAseService] method
// instead.
type RadarHTTPAseService struct {
	Options      []option.RequestOption
	BotClass     *RadarHTTPAseBotClassService
	DeviceType   *RadarHTTPAseDeviceTypeService
	HTTPProtocol *RadarHTTPAseHTTPProtocolService
	HTTPMethod   *RadarHTTPAseHTTPMethodService
	IPVersion    *RadarHTTPAseIPVersionService
	Os           *RadarHTTPAseOService
	TLSVersion   *RadarHTTPAseTLSVersionService
}

// NewRadarHTTPAseService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarHTTPAseService(opts ...option.RequestOption) (r *RadarHTTPAseService) {
	r = &RadarHTTPAseService{}
	r.Options = opts
	r.BotClass = NewRadarHTTPAseBotClassService(opts...)
	r.DeviceType = NewRadarHTTPAseDeviceTypeService(opts...)
	r.HTTPProtocol = NewRadarHTTPAseHTTPProtocolService(opts...)
	r.HTTPMethod = NewRadarHTTPAseHTTPMethodService(opts...)
	r.IPVersion = NewRadarHTTPAseIPVersionService(opts...)
	r.Os = NewRadarHTTPAseOService(opts...)
	r.TLSVersion = NewRadarHTTPAseTLSVersionService(opts...)
	return
}

// Get the top autonomous systems by HTTP traffic. Values are a percentage out of
// the total traffic.
func (r *RadarHTTPAseService) Get(ctx context.Context, query RadarHTTPAseGetParams, opts ...option.RequestOption) (res *RadarHTTPAseGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarHTTPAseGetResponseEnvelope
	path := "radar/http/top/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarHTTPAseGetResponse struct {
	Meta RadarHTTPAseGetResponseMeta   `json:"meta,required"`
	Top0 []RadarHTTPAseGetResponseTop0 `json:"top_0,required"`
	JSON radarHTTPAseGetResponseJSON   `json:"-"`
}

// radarHTTPAseGetResponseJSON contains the JSON metadata for the struct
// [RadarHTTPAseGetResponse]
type radarHTTPAseGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPAseGetResponseMeta struct {
	DateRange      []RadarHTTPAseGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                    `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPAseGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPAseGetResponseMetaJSON           `json:"-"`
}

// radarHTTPAseGetResponseMetaJSON contains the JSON metadata for the struct
// [RadarHTTPAseGetResponseMeta]
type radarHTTPAseGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPAseGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPAseGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPAseGetResponseMetaDateRangeJSON `json:"-"`
}

// radarHTTPAseGetResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [RadarHTTPAseGetResponseMetaDateRange]
type radarHTTPAseGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPAseGetResponseMetaConfidenceInfo struct {
	Annotations []RadarHTTPAseGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                 `json:"level"`
	JSON        radarHTTPAseGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPAseGetResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [RadarHTTPAseGetResponseMetaConfidenceInfo]
type radarHTTPAseGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPAseGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                  `json:"dataSource,required"`
	Description     string                                                  `json:"description,required"`
	EventType       string                                                  `json:"eventType,required"`
	IsInstantaneous interface{}                                             `json:"isInstantaneous,required"`
	EndTime         time.Time                                               `json:"endTime" format:"date-time"`
	LinkedURL       string                                                  `json:"linkedUrl"`
	StartTime       time.Time                                               `json:"startTime" format:"date-time"`
	JSON            radarHTTPAseGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPAseGetResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [RadarHTTPAseGetResponseMetaConfidenceInfoAnnotation]
type radarHTTPAseGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPAseGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPAseGetResponseTop0 struct {
	ClientAsn    int64                           `json:"clientASN,required"`
	ClientAsName string                          `json:"clientASName,required"`
	Value        string                          `json:"value,required"`
	JSON         radarHTTPAseGetResponseTop0JSON `json:"-"`
}

// radarHTTPAseGetResponseTop0JSON contains the JSON metadata for the struct
// [RadarHTTPAseGetResponseTop0]
type radarHTTPAseGetResponseTop0JSON struct {
	ClientAsn    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarHTTPAseGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPAseGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPAseGetParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPAseGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPAseGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPAseGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPAseGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPAseGetParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPAseGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPAseGetParamsO] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHTTPAseGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPAseGetParams]'s query parameters as `url.Values`.
func (r RadarHTTPAseGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarHTTPAseGetParamsBotClass string

const (
	RadarHTTPAseGetParamsBotClassLikelyAutomated RadarHTTPAseGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPAseGetParamsBotClassLikelyHuman     RadarHTTPAseGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPAseGetParamsDateRange string

const (
	RadarHTTPAseGetParamsDateRange1d         RadarHTTPAseGetParamsDateRange = "1d"
	RadarHTTPAseGetParamsDateRange2d         RadarHTTPAseGetParamsDateRange = "2d"
	RadarHTTPAseGetParamsDateRange7d         RadarHTTPAseGetParamsDateRange = "7d"
	RadarHTTPAseGetParamsDateRange14d        RadarHTTPAseGetParamsDateRange = "14d"
	RadarHTTPAseGetParamsDateRange28d        RadarHTTPAseGetParamsDateRange = "28d"
	RadarHTTPAseGetParamsDateRange12w        RadarHTTPAseGetParamsDateRange = "12w"
	RadarHTTPAseGetParamsDateRange24w        RadarHTTPAseGetParamsDateRange = "24w"
	RadarHTTPAseGetParamsDateRange52w        RadarHTTPAseGetParamsDateRange = "52w"
	RadarHTTPAseGetParamsDateRange1dControl  RadarHTTPAseGetParamsDateRange = "1dControl"
	RadarHTTPAseGetParamsDateRange2dControl  RadarHTTPAseGetParamsDateRange = "2dControl"
	RadarHTTPAseGetParamsDateRange7dControl  RadarHTTPAseGetParamsDateRange = "7dControl"
	RadarHTTPAseGetParamsDateRange14dControl RadarHTTPAseGetParamsDateRange = "14dControl"
	RadarHTTPAseGetParamsDateRange28dControl RadarHTTPAseGetParamsDateRange = "28dControl"
	RadarHTTPAseGetParamsDateRange12wControl RadarHTTPAseGetParamsDateRange = "12wControl"
	RadarHTTPAseGetParamsDateRange24wControl RadarHTTPAseGetParamsDateRange = "24wControl"
)

type RadarHTTPAseGetParamsDeviceType string

const (
	RadarHTTPAseGetParamsDeviceTypeDesktop RadarHTTPAseGetParamsDeviceType = "DESKTOP"
	RadarHTTPAseGetParamsDeviceTypeMobile  RadarHTTPAseGetParamsDeviceType = "MOBILE"
	RadarHTTPAseGetParamsDeviceTypeOther   RadarHTTPAseGetParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPAseGetParamsFormat string

const (
	RadarHTTPAseGetParamsFormatJson RadarHTTPAseGetParamsFormat = "JSON"
	RadarHTTPAseGetParamsFormatCsv  RadarHTTPAseGetParamsFormat = "CSV"
)

type RadarHTTPAseGetParamsHTTPProtocol string

const (
	RadarHTTPAseGetParamsHTTPProtocolHTTP  RadarHTTPAseGetParamsHTTPProtocol = "HTTP"
	RadarHTTPAseGetParamsHTTPProtocolHTTPS RadarHTTPAseGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPAseGetParamsHTTPVersion string

const (
	RadarHTTPAseGetParamsHTTPVersionHttPv1 RadarHTTPAseGetParamsHTTPVersion = "HTTPv1"
	RadarHTTPAseGetParamsHTTPVersionHttPv2 RadarHTTPAseGetParamsHTTPVersion = "HTTPv2"
	RadarHTTPAseGetParamsHTTPVersionHttPv3 RadarHTTPAseGetParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPAseGetParamsIPVersion string

const (
	RadarHTTPAseGetParamsIPVersionIPv4 RadarHTTPAseGetParamsIPVersion = "IPv4"
	RadarHTTPAseGetParamsIPVersionIPv6 RadarHTTPAseGetParamsIPVersion = "IPv6"
)

type RadarHTTPAseGetParamsO string

const (
	RadarHTTPAseGetParamsOWindows  RadarHTTPAseGetParamsO = "WINDOWS"
	RadarHTTPAseGetParamsOMacosx   RadarHTTPAseGetParamsO = "MACOSX"
	RadarHTTPAseGetParamsOIos      RadarHTTPAseGetParamsO = "IOS"
	RadarHTTPAseGetParamsOAndroid  RadarHTTPAseGetParamsO = "ANDROID"
	RadarHTTPAseGetParamsOChromeos RadarHTTPAseGetParamsO = "CHROMEOS"
	RadarHTTPAseGetParamsOLinux    RadarHTTPAseGetParamsO = "LINUX"
	RadarHTTPAseGetParamsOSmartTv  RadarHTTPAseGetParamsO = "SMART_TV"
)

type RadarHTTPAseGetParamsTLSVersion string

const (
	RadarHTTPAseGetParamsTLSVersionTlSv1_0  RadarHTTPAseGetParamsTLSVersion = "TLSv1_0"
	RadarHTTPAseGetParamsTLSVersionTlSv1_1  RadarHTTPAseGetParamsTLSVersion = "TLSv1_1"
	RadarHTTPAseGetParamsTLSVersionTlSv1_2  RadarHTTPAseGetParamsTLSVersion = "TLSv1_2"
	RadarHTTPAseGetParamsTLSVersionTlSv1_3  RadarHTTPAseGetParamsTLSVersion = "TLSv1_3"
	RadarHTTPAseGetParamsTLSVersionTlSvQuic RadarHTTPAseGetParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPAseGetResponseEnvelope struct {
	Result  RadarHTTPAseGetResponse             `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    radarHTTPAseGetResponseEnvelopeJSON `json:"-"`
}

// radarHTTPAseGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RadarHTTPAseGetResponseEnvelope]
type radarHTTPAseGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
