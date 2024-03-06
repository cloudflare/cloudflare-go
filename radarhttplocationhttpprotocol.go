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

// RadarHTTPLocationHTTPProtocolService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarHTTPLocationHTTPProtocolService] method instead.
type RadarHTTPLocationHTTPProtocolService struct {
	Options []option.RequestOption
}

// NewRadarHTTPLocationHTTPProtocolService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPLocationHTTPProtocolService(opts ...option.RequestOption) (r *RadarHTTPLocationHTTPProtocolService) {
	r = &RadarHTTPLocationHTTPProtocolService{}
	r.Options = opts
	return
}

// Get the top locations, by HTTP traffic, of the requested HTTP protocol. Values
// are a percentage out of the total traffic.
func (r *RadarHTTPLocationHTTPProtocolService) Get(ctx context.Context, httpProtocol RadarHTTPLocationHTTPProtocolGetParamsHTTPProtocol, query RadarHTTPLocationHTTPProtocolGetParams, opts ...option.RequestOption) (res *RadarHTTPLocationHTTPProtocolGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarHTTPLocationHTTPProtocolGetResponseEnvelope
	path := fmt.Sprintf("radar/http/top/locations/http_protocol/%v", httpProtocol)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarHTTPLocationHTTPProtocolGetResponse struct {
	Meta RadarHTTPLocationHTTPProtocolGetResponseMeta   `json:"meta,required"`
	Top0 []RadarHTTPLocationHTTPProtocolGetResponseTop0 `json:"top_0,required"`
	JSON radarHTTPLocationHTTPProtocolGetResponseJSON   `json:"-"`
}

// radarHTTPLocationHTTPProtocolGetResponseJSON contains the JSON metadata for the
// struct [RadarHTTPLocationHTTPProtocolGetResponse]
type radarHTTPLocationHTTPProtocolGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationHTTPProtocolGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPLocationHTTPProtocolGetResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPLocationHTTPProtocolGetResponseMeta struct {
	DateRange      []RadarHTTPLocationHTTPProtocolGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                     `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPLocationHTTPProtocolGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPLocationHTTPProtocolGetResponseMetaJSON           `json:"-"`
}

// radarHTTPLocationHTTPProtocolGetResponseMetaJSON contains the JSON metadata for
// the struct [RadarHTTPLocationHTTPProtocolGetResponseMeta]
type radarHTTPLocationHTTPProtocolGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPLocationHTTPProtocolGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPLocationHTTPProtocolGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPLocationHTTPProtocolGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                 `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPLocationHTTPProtocolGetResponseMetaDateRangeJSON `json:"-"`
}

// radarHTTPLocationHTTPProtocolGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarHTTPLocationHTTPProtocolGetResponseMetaDateRange]
type radarHTTPLocationHTTPProtocolGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationHTTPProtocolGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPLocationHTTPProtocolGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPLocationHTTPProtocolGetResponseMetaConfidenceInfo struct {
	Annotations []RadarHTTPLocationHTTPProtocolGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                  `json:"level"`
	JSON        radarHTTPLocationHTTPProtocolGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPLocationHTTPProtocolGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarHTTPLocationHTTPProtocolGetResponseMetaConfidenceInfo]
type radarHTTPLocationHTTPProtocolGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationHTTPProtocolGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPLocationHTTPProtocolGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPLocationHTTPProtocolGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                   `json:"dataSource,required"`
	Description     string                                                                   `json:"description,required"`
	EventType       string                                                                   `json:"eventType,required"`
	IsInstantaneous interface{}                                                              `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                   `json:"linkedUrl"`
	StartTime       time.Time                                                                `json:"startTime" format:"date-time"`
	JSON            radarHTTPLocationHTTPProtocolGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPLocationHTTPProtocolGetResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPLocationHTTPProtocolGetResponseMetaConfidenceInfoAnnotation]
type radarHTTPLocationHTTPProtocolGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPLocationHTTPProtocolGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPLocationHTTPProtocolGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPLocationHTTPProtocolGetResponseTop0 struct {
	ClientCountryAlpha2 string                                           `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                           `json:"clientCountryName,required"`
	Value               string                                           `json:"value,required"`
	JSON                radarHTTPLocationHTTPProtocolGetResponseTop0JSON `json:"-"`
}

// radarHTTPLocationHTTPProtocolGetResponseTop0JSON contains the JSON metadata for
// the struct [RadarHTTPLocationHTTPProtocolGetResponseTop0]
type radarHTTPLocationHTTPProtocolGetResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarHTTPLocationHTTPProtocolGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPLocationHTTPProtocolGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPLocationHTTPProtocolGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPLocationHTTPProtocolGetParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPLocationHTTPProtocolGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPLocationHTTPProtocolGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPLocationHTTPProtocolGetParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPLocationHTTPProtocolGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]RadarHTTPLocationHTTPProtocolGetParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHTTPLocationHTTPProtocolGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPLocationHTTPProtocolGetParams]'s query parameters
// as `url.Values`.
func (r RadarHTTPLocationHTTPProtocolGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// HTTP Protocol.
type RadarHTTPLocationHTTPProtocolGetParamsHTTPProtocol string

const (
	RadarHTTPLocationHTTPProtocolGetParamsHTTPProtocolHTTP  RadarHTTPLocationHTTPProtocolGetParamsHTTPProtocol = "HTTP"
	RadarHTTPLocationHTTPProtocolGetParamsHTTPProtocolHTTPS RadarHTTPLocationHTTPProtocolGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPLocationHTTPProtocolGetParamsBotClass string

const (
	RadarHTTPLocationHTTPProtocolGetParamsBotClassLikelyAutomated RadarHTTPLocationHTTPProtocolGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPLocationHTTPProtocolGetParamsBotClassLikelyHuman     RadarHTTPLocationHTTPProtocolGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPLocationHTTPProtocolGetParamsDateRange string

const (
	RadarHTTPLocationHTTPProtocolGetParamsDateRange1d         RadarHTTPLocationHTTPProtocolGetParamsDateRange = "1d"
	RadarHTTPLocationHTTPProtocolGetParamsDateRange2d         RadarHTTPLocationHTTPProtocolGetParamsDateRange = "2d"
	RadarHTTPLocationHTTPProtocolGetParamsDateRange7d         RadarHTTPLocationHTTPProtocolGetParamsDateRange = "7d"
	RadarHTTPLocationHTTPProtocolGetParamsDateRange14d        RadarHTTPLocationHTTPProtocolGetParamsDateRange = "14d"
	RadarHTTPLocationHTTPProtocolGetParamsDateRange28d        RadarHTTPLocationHTTPProtocolGetParamsDateRange = "28d"
	RadarHTTPLocationHTTPProtocolGetParamsDateRange12w        RadarHTTPLocationHTTPProtocolGetParamsDateRange = "12w"
	RadarHTTPLocationHTTPProtocolGetParamsDateRange24w        RadarHTTPLocationHTTPProtocolGetParamsDateRange = "24w"
	RadarHTTPLocationHTTPProtocolGetParamsDateRange52w        RadarHTTPLocationHTTPProtocolGetParamsDateRange = "52w"
	RadarHTTPLocationHTTPProtocolGetParamsDateRange1dControl  RadarHTTPLocationHTTPProtocolGetParamsDateRange = "1dControl"
	RadarHTTPLocationHTTPProtocolGetParamsDateRange2dControl  RadarHTTPLocationHTTPProtocolGetParamsDateRange = "2dControl"
	RadarHTTPLocationHTTPProtocolGetParamsDateRange7dControl  RadarHTTPLocationHTTPProtocolGetParamsDateRange = "7dControl"
	RadarHTTPLocationHTTPProtocolGetParamsDateRange14dControl RadarHTTPLocationHTTPProtocolGetParamsDateRange = "14dControl"
	RadarHTTPLocationHTTPProtocolGetParamsDateRange28dControl RadarHTTPLocationHTTPProtocolGetParamsDateRange = "28dControl"
	RadarHTTPLocationHTTPProtocolGetParamsDateRange12wControl RadarHTTPLocationHTTPProtocolGetParamsDateRange = "12wControl"
	RadarHTTPLocationHTTPProtocolGetParamsDateRange24wControl RadarHTTPLocationHTTPProtocolGetParamsDateRange = "24wControl"
)

type RadarHTTPLocationHTTPProtocolGetParamsDeviceType string

const (
	RadarHTTPLocationHTTPProtocolGetParamsDeviceTypeDesktop RadarHTTPLocationHTTPProtocolGetParamsDeviceType = "DESKTOP"
	RadarHTTPLocationHTTPProtocolGetParamsDeviceTypeMobile  RadarHTTPLocationHTTPProtocolGetParamsDeviceType = "MOBILE"
	RadarHTTPLocationHTTPProtocolGetParamsDeviceTypeOther   RadarHTTPLocationHTTPProtocolGetParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPLocationHTTPProtocolGetParamsFormat string

const (
	RadarHTTPLocationHTTPProtocolGetParamsFormatJson RadarHTTPLocationHTTPProtocolGetParamsFormat = "JSON"
	RadarHTTPLocationHTTPProtocolGetParamsFormatCsv  RadarHTTPLocationHTTPProtocolGetParamsFormat = "CSV"
)

type RadarHTTPLocationHTTPProtocolGetParamsIPVersion string

const (
	RadarHTTPLocationHTTPProtocolGetParamsIPVersionIPv4 RadarHTTPLocationHTTPProtocolGetParamsIPVersion = "IPv4"
	RadarHTTPLocationHTTPProtocolGetParamsIPVersionIPv6 RadarHTTPLocationHTTPProtocolGetParamsIPVersion = "IPv6"
)

type RadarHTTPLocationHTTPProtocolGetParamsOS string

const (
	RadarHTTPLocationHTTPProtocolGetParamsOSWindows  RadarHTTPLocationHTTPProtocolGetParamsOS = "WINDOWS"
	RadarHTTPLocationHTTPProtocolGetParamsOSMacosx   RadarHTTPLocationHTTPProtocolGetParamsOS = "MACOSX"
	RadarHTTPLocationHTTPProtocolGetParamsOSIos      RadarHTTPLocationHTTPProtocolGetParamsOS = "IOS"
	RadarHTTPLocationHTTPProtocolGetParamsOSAndroid  RadarHTTPLocationHTTPProtocolGetParamsOS = "ANDROID"
	RadarHTTPLocationHTTPProtocolGetParamsOSChromeos RadarHTTPLocationHTTPProtocolGetParamsOS = "CHROMEOS"
	RadarHTTPLocationHTTPProtocolGetParamsOSLinux    RadarHTTPLocationHTTPProtocolGetParamsOS = "LINUX"
	RadarHTTPLocationHTTPProtocolGetParamsOSSmartTv  RadarHTTPLocationHTTPProtocolGetParamsOS = "SMART_TV"
)

type RadarHTTPLocationHTTPProtocolGetParamsTLSVersion string

const (
	RadarHTTPLocationHTTPProtocolGetParamsTLSVersionTlSv1_0  RadarHTTPLocationHTTPProtocolGetParamsTLSVersion = "TLSv1_0"
	RadarHTTPLocationHTTPProtocolGetParamsTLSVersionTlSv1_1  RadarHTTPLocationHTTPProtocolGetParamsTLSVersion = "TLSv1_1"
	RadarHTTPLocationHTTPProtocolGetParamsTLSVersionTlSv1_2  RadarHTTPLocationHTTPProtocolGetParamsTLSVersion = "TLSv1_2"
	RadarHTTPLocationHTTPProtocolGetParamsTLSVersionTlSv1_3  RadarHTTPLocationHTTPProtocolGetParamsTLSVersion = "TLSv1_3"
	RadarHTTPLocationHTTPProtocolGetParamsTLSVersionTlSvQuic RadarHTTPLocationHTTPProtocolGetParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPLocationHTTPProtocolGetResponseEnvelope struct {
	Result  RadarHTTPLocationHTTPProtocolGetResponse             `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarHTTPLocationHTTPProtocolGetResponseEnvelopeJSON `json:"-"`
}

// radarHTTPLocationHTTPProtocolGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarHTTPLocationHTTPProtocolGetResponseEnvelope]
type radarHTTPLocationHTTPProtocolGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationHTTPProtocolGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPLocationHTTPProtocolGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
