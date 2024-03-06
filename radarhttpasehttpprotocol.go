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

// RadarHTTPAseHTTPProtocolService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarHTTPAseHTTPProtocolService] method instead.
type RadarHTTPAseHTTPProtocolService struct {
	Options []option.RequestOption
}

// NewRadarHTTPAseHTTPProtocolService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPAseHTTPProtocolService(opts ...option.RequestOption) (r *RadarHTTPAseHTTPProtocolService) {
	r = &RadarHTTPAseHTTPProtocolService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS), by HTTP traffic, of the requested HTTP
// protocol. Values are a percentage out of the total traffic.
func (r *RadarHTTPAseHTTPProtocolService) Get(ctx context.Context, httpProtocol RadarHTTPAseHTTPProtocolGetParamsHTTPProtocol, query RadarHTTPAseHTTPProtocolGetParams, opts ...option.RequestOption) (res *RadarHTTPAseHTTPProtocolGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarHTTPAseHTTPProtocolGetResponseEnvelope
	path := fmt.Sprintf("radar/http/top/ases/http_protocol/%v", httpProtocol)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarHTTPAseHTTPProtocolGetResponse struct {
	Meta RadarHTTPAseHTTPProtocolGetResponseMeta   `json:"meta,required"`
	Top0 []RadarHTTPAseHTTPProtocolGetResponseTop0 `json:"top_0,required"`
	JSON radarHTTPAseHTTPProtocolGetResponseJSON   `json:"-"`
}

// radarHTTPAseHTTPProtocolGetResponseJSON contains the JSON metadata for the
// struct [RadarHTTPAseHTTPProtocolGetResponse]
type radarHTTPAseHTTPProtocolGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseHTTPProtocolGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPAseHTTPProtocolGetResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPAseHTTPProtocolGetResponseMeta struct {
	DateRange      []RadarHTTPAseHTTPProtocolGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPAseHTTPProtocolGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPAseHTTPProtocolGetResponseMetaJSON           `json:"-"`
}

// radarHTTPAseHTTPProtocolGetResponseMetaJSON contains the JSON metadata for the
// struct [RadarHTTPAseHTTPProtocolGetResponseMeta]
type radarHTTPAseHTTPProtocolGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPAseHTTPProtocolGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPAseHTTPProtocolGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPAseHTTPProtocolGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                            `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPAseHTTPProtocolGetResponseMetaDateRangeJSON `json:"-"`
}

// radarHTTPAseHTTPProtocolGetResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarHTTPAseHTTPProtocolGetResponseMetaDateRange]
type radarHTTPAseHTTPProtocolGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseHTTPProtocolGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPAseHTTPProtocolGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPAseHTTPProtocolGetResponseMetaConfidenceInfo struct {
	Annotations []RadarHTTPAseHTTPProtocolGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                             `json:"level"`
	JSON        radarHTTPAseHTTPProtocolGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPAseHTTPProtocolGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarHTTPAseHTTPProtocolGetResponseMetaConfidenceInfo]
type radarHTTPAseHTTPProtocolGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseHTTPProtocolGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPAseHTTPProtocolGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPAseHTTPProtocolGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                              `json:"dataSource,required"`
	Description     string                                                              `json:"description,required"`
	EventType       string                                                              `json:"eventType,required"`
	IsInstantaneous interface{}                                                         `json:"isInstantaneous,required"`
	EndTime         time.Time                                                           `json:"endTime" format:"date-time"`
	LinkedURL       string                                                              `json:"linkedUrl"`
	StartTime       time.Time                                                           `json:"startTime" format:"date-time"`
	JSON            radarHTTPAseHTTPProtocolGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPAseHTTPProtocolGetResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarHTTPAseHTTPProtocolGetResponseMetaConfidenceInfoAnnotation]
type radarHTTPAseHTTPProtocolGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPAseHTTPProtocolGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPAseHTTPProtocolGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPAseHTTPProtocolGetResponseTop0 struct {
	ClientASN    int64                                       `json:"clientASN,required"`
	ClientAsName string                                      `json:"clientASName,required"`
	Value        string                                      `json:"value,required"`
	JSON         radarHTTPAseHTTPProtocolGetResponseTop0JSON `json:"-"`
}

// radarHTTPAseHTTPProtocolGetResponseTop0JSON contains the JSON metadata for the
// struct [RadarHTTPAseHTTPProtocolGetResponseTop0]
type radarHTTPAseHTTPProtocolGetResponseTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarHTTPAseHTTPProtocolGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPAseHTTPProtocolGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPAseHTTPProtocolGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPAseHTTPProtocolGetParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPAseHTTPProtocolGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPAseHTTPProtocolGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPAseHTTPProtocolGetParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPAseHTTPProtocolGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]RadarHTTPAseHTTPProtocolGetParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHTTPAseHTTPProtocolGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPAseHTTPProtocolGetParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPAseHTTPProtocolGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// HTTP Protocol.
type RadarHTTPAseHTTPProtocolGetParamsHTTPProtocol string

const (
	RadarHTTPAseHTTPProtocolGetParamsHTTPProtocolHTTP  RadarHTTPAseHTTPProtocolGetParamsHTTPProtocol = "HTTP"
	RadarHTTPAseHTTPProtocolGetParamsHTTPProtocolHTTPS RadarHTTPAseHTTPProtocolGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPAseHTTPProtocolGetParamsBotClass string

const (
	RadarHTTPAseHTTPProtocolGetParamsBotClassLikelyAutomated RadarHTTPAseHTTPProtocolGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPAseHTTPProtocolGetParamsBotClassLikelyHuman     RadarHTTPAseHTTPProtocolGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPAseHTTPProtocolGetParamsDateRange string

const (
	RadarHTTPAseHTTPProtocolGetParamsDateRange1d         RadarHTTPAseHTTPProtocolGetParamsDateRange = "1d"
	RadarHTTPAseHTTPProtocolGetParamsDateRange2d         RadarHTTPAseHTTPProtocolGetParamsDateRange = "2d"
	RadarHTTPAseHTTPProtocolGetParamsDateRange7d         RadarHTTPAseHTTPProtocolGetParamsDateRange = "7d"
	RadarHTTPAseHTTPProtocolGetParamsDateRange14d        RadarHTTPAseHTTPProtocolGetParamsDateRange = "14d"
	RadarHTTPAseHTTPProtocolGetParamsDateRange28d        RadarHTTPAseHTTPProtocolGetParamsDateRange = "28d"
	RadarHTTPAseHTTPProtocolGetParamsDateRange12w        RadarHTTPAseHTTPProtocolGetParamsDateRange = "12w"
	RadarHTTPAseHTTPProtocolGetParamsDateRange24w        RadarHTTPAseHTTPProtocolGetParamsDateRange = "24w"
	RadarHTTPAseHTTPProtocolGetParamsDateRange52w        RadarHTTPAseHTTPProtocolGetParamsDateRange = "52w"
	RadarHTTPAseHTTPProtocolGetParamsDateRange1dControl  RadarHTTPAseHTTPProtocolGetParamsDateRange = "1dControl"
	RadarHTTPAseHTTPProtocolGetParamsDateRange2dControl  RadarHTTPAseHTTPProtocolGetParamsDateRange = "2dControl"
	RadarHTTPAseHTTPProtocolGetParamsDateRange7dControl  RadarHTTPAseHTTPProtocolGetParamsDateRange = "7dControl"
	RadarHTTPAseHTTPProtocolGetParamsDateRange14dControl RadarHTTPAseHTTPProtocolGetParamsDateRange = "14dControl"
	RadarHTTPAseHTTPProtocolGetParamsDateRange28dControl RadarHTTPAseHTTPProtocolGetParamsDateRange = "28dControl"
	RadarHTTPAseHTTPProtocolGetParamsDateRange12wControl RadarHTTPAseHTTPProtocolGetParamsDateRange = "12wControl"
	RadarHTTPAseHTTPProtocolGetParamsDateRange24wControl RadarHTTPAseHTTPProtocolGetParamsDateRange = "24wControl"
)

type RadarHTTPAseHTTPProtocolGetParamsDeviceType string

const (
	RadarHTTPAseHTTPProtocolGetParamsDeviceTypeDesktop RadarHTTPAseHTTPProtocolGetParamsDeviceType = "DESKTOP"
	RadarHTTPAseHTTPProtocolGetParamsDeviceTypeMobile  RadarHTTPAseHTTPProtocolGetParamsDeviceType = "MOBILE"
	RadarHTTPAseHTTPProtocolGetParamsDeviceTypeOther   RadarHTTPAseHTTPProtocolGetParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPAseHTTPProtocolGetParamsFormat string

const (
	RadarHTTPAseHTTPProtocolGetParamsFormatJson RadarHTTPAseHTTPProtocolGetParamsFormat = "JSON"
	RadarHTTPAseHTTPProtocolGetParamsFormatCsv  RadarHTTPAseHTTPProtocolGetParamsFormat = "CSV"
)

type RadarHTTPAseHTTPProtocolGetParamsIPVersion string

const (
	RadarHTTPAseHTTPProtocolGetParamsIPVersionIPv4 RadarHTTPAseHTTPProtocolGetParamsIPVersion = "IPv4"
	RadarHTTPAseHTTPProtocolGetParamsIPVersionIPv6 RadarHTTPAseHTTPProtocolGetParamsIPVersion = "IPv6"
)

type RadarHTTPAseHTTPProtocolGetParamsOS string

const (
	RadarHTTPAseHTTPProtocolGetParamsOSWindows  RadarHTTPAseHTTPProtocolGetParamsOS = "WINDOWS"
	RadarHTTPAseHTTPProtocolGetParamsOSMacosx   RadarHTTPAseHTTPProtocolGetParamsOS = "MACOSX"
	RadarHTTPAseHTTPProtocolGetParamsOSIos      RadarHTTPAseHTTPProtocolGetParamsOS = "IOS"
	RadarHTTPAseHTTPProtocolGetParamsOSAndroid  RadarHTTPAseHTTPProtocolGetParamsOS = "ANDROID"
	RadarHTTPAseHTTPProtocolGetParamsOSChromeos RadarHTTPAseHTTPProtocolGetParamsOS = "CHROMEOS"
	RadarHTTPAseHTTPProtocolGetParamsOSLinux    RadarHTTPAseHTTPProtocolGetParamsOS = "LINUX"
	RadarHTTPAseHTTPProtocolGetParamsOSSmartTv  RadarHTTPAseHTTPProtocolGetParamsOS = "SMART_TV"
)

type RadarHTTPAseHTTPProtocolGetParamsTLSVersion string

const (
	RadarHTTPAseHTTPProtocolGetParamsTLSVersionTlSv1_0  RadarHTTPAseHTTPProtocolGetParamsTLSVersion = "TLSv1_0"
	RadarHTTPAseHTTPProtocolGetParamsTLSVersionTlSv1_1  RadarHTTPAseHTTPProtocolGetParamsTLSVersion = "TLSv1_1"
	RadarHTTPAseHTTPProtocolGetParamsTLSVersionTlSv1_2  RadarHTTPAseHTTPProtocolGetParamsTLSVersion = "TLSv1_2"
	RadarHTTPAseHTTPProtocolGetParamsTLSVersionTlSv1_3  RadarHTTPAseHTTPProtocolGetParamsTLSVersion = "TLSv1_3"
	RadarHTTPAseHTTPProtocolGetParamsTLSVersionTlSvQuic RadarHTTPAseHTTPProtocolGetParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPAseHTTPProtocolGetResponseEnvelope struct {
	Result  RadarHTTPAseHTTPProtocolGetResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarHTTPAseHTTPProtocolGetResponseEnvelopeJSON `json:"-"`
}

// radarHTTPAseHTTPProtocolGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarHTTPAseHTTPProtocolGetResponseEnvelope]
type radarHTTPAseHTTPProtocolGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseHTTPProtocolGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPAseHTTPProtocolGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
