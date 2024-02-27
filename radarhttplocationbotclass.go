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

// RadarHTTPLocationBotClassService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarHTTPLocationBotClassService] method instead.
type RadarHTTPLocationBotClassService struct {
	Options []option.RequestOption
}

// NewRadarHTTPLocationBotClassService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPLocationBotClassService(opts ...option.RequestOption) (r *RadarHTTPLocationBotClassService) {
	r = &RadarHTTPLocationBotClassService{}
	r.Options = opts
	return
}

// Get the top locations, by HTTP traffic, of the requested bot class. These two
// categories use Cloudflare's bot score - refer to [Bot
// scores])https://developers.cloudflare.com/bots/concepts/bot-score). Values are a
// percentage out of the total traffic.
func (r *RadarHTTPLocationBotClassService) Get(ctx context.Context, botClass RadarHTTPLocationBotClassGetParamsBotClass, query RadarHTTPLocationBotClassGetParams, opts ...option.RequestOption) (res *RadarHTTPLocationBotClassGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarHTTPLocationBotClassGetResponseEnvelope
	path := fmt.Sprintf("radar/http/top/locations/bot_class/%v", botClass)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarHTTPLocationBotClassGetResponse struct {
	Meta RadarHTTPLocationBotClassGetResponseMeta   `json:"meta,required"`
	Top0 []RadarHTTPLocationBotClassGetResponseTop0 `json:"top_0,required"`
	JSON radarHTTPLocationBotClassGetResponseJSON   `json:"-"`
}

// radarHTTPLocationBotClassGetResponseJSON contains the JSON metadata for the
// struct [RadarHTTPLocationBotClassGetResponse]
type radarHTTPLocationBotClassGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationBotClassGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPLocationBotClassGetResponseMeta struct {
	DateRange      []RadarHTTPLocationBotClassGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                 `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPLocationBotClassGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPLocationBotClassGetResponseMetaJSON           `json:"-"`
}

// radarHTTPLocationBotClassGetResponseMetaJSON contains the JSON metadata for the
// struct [RadarHTTPLocationBotClassGetResponseMeta]
type radarHTTPLocationBotClassGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPLocationBotClassGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPLocationBotClassGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPLocationBotClassGetResponseMetaDateRangeJSON `json:"-"`
}

// radarHTTPLocationBotClassGetResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarHTTPLocationBotClassGetResponseMetaDateRange]
type radarHTTPLocationBotClassGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationBotClassGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPLocationBotClassGetResponseMetaConfidenceInfo struct {
	Annotations []RadarHTTPLocationBotClassGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                              `json:"level"`
	JSON        radarHTTPLocationBotClassGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPLocationBotClassGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarHTTPLocationBotClassGetResponseMetaConfidenceInfo]
type radarHTTPLocationBotClassGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationBotClassGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPLocationBotClassGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                               `json:"dataSource,required"`
	Description     string                                                               `json:"description,required"`
	EventType       string                                                               `json:"eventType,required"`
	IsInstantaneous interface{}                                                          `json:"isInstantaneous,required"`
	EndTime         time.Time                                                            `json:"endTime" format:"date-time"`
	LinkedURL       string                                                               `json:"linkedUrl"`
	StartTime       time.Time                                                            `json:"startTime" format:"date-time"`
	JSON            radarHTTPLocationBotClassGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPLocationBotClassGetResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarHTTPLocationBotClassGetResponseMetaConfidenceInfoAnnotation]
type radarHTTPLocationBotClassGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPLocationBotClassGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPLocationBotClassGetResponseTop0 struct {
	ClientCountryAlpha2 string                                       `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                       `json:"clientCountryName,required"`
	Value               string                                       `json:"value,required"`
	JSON                radarHTTPLocationBotClassGetResponseTop0JSON `json:"-"`
}

// radarHTTPLocationBotClassGetResponseTop0JSON contains the JSON metadata for the
// struct [RadarHTTPLocationBotClassGetResponseTop0]
type radarHTTPLocationBotClassGetResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarHTTPLocationBotClassGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPLocationBotClassGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPLocationBotClassGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPLocationBotClassGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPLocationBotClassGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPLocationBotClassGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPLocationBotClassGetParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPLocationBotClassGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]RadarHTTPLocationBotClassGetParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHTTPLocationBotClassGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPLocationBotClassGetParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPLocationBotClassGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Bot class.
type RadarHTTPLocationBotClassGetParamsBotClass string

const (
	RadarHTTPLocationBotClassGetParamsBotClassLikelyAutomated RadarHTTPLocationBotClassGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPLocationBotClassGetParamsBotClassLikelyHuman     RadarHTTPLocationBotClassGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPLocationBotClassGetParamsDateRange string

const (
	RadarHTTPLocationBotClassGetParamsDateRange1d         RadarHTTPLocationBotClassGetParamsDateRange = "1d"
	RadarHTTPLocationBotClassGetParamsDateRange2d         RadarHTTPLocationBotClassGetParamsDateRange = "2d"
	RadarHTTPLocationBotClassGetParamsDateRange7d         RadarHTTPLocationBotClassGetParamsDateRange = "7d"
	RadarHTTPLocationBotClassGetParamsDateRange14d        RadarHTTPLocationBotClassGetParamsDateRange = "14d"
	RadarHTTPLocationBotClassGetParamsDateRange28d        RadarHTTPLocationBotClassGetParamsDateRange = "28d"
	RadarHTTPLocationBotClassGetParamsDateRange12w        RadarHTTPLocationBotClassGetParamsDateRange = "12w"
	RadarHTTPLocationBotClassGetParamsDateRange24w        RadarHTTPLocationBotClassGetParamsDateRange = "24w"
	RadarHTTPLocationBotClassGetParamsDateRange52w        RadarHTTPLocationBotClassGetParamsDateRange = "52w"
	RadarHTTPLocationBotClassGetParamsDateRange1dControl  RadarHTTPLocationBotClassGetParamsDateRange = "1dControl"
	RadarHTTPLocationBotClassGetParamsDateRange2dControl  RadarHTTPLocationBotClassGetParamsDateRange = "2dControl"
	RadarHTTPLocationBotClassGetParamsDateRange7dControl  RadarHTTPLocationBotClassGetParamsDateRange = "7dControl"
	RadarHTTPLocationBotClassGetParamsDateRange14dControl RadarHTTPLocationBotClassGetParamsDateRange = "14dControl"
	RadarHTTPLocationBotClassGetParamsDateRange28dControl RadarHTTPLocationBotClassGetParamsDateRange = "28dControl"
	RadarHTTPLocationBotClassGetParamsDateRange12wControl RadarHTTPLocationBotClassGetParamsDateRange = "12wControl"
	RadarHTTPLocationBotClassGetParamsDateRange24wControl RadarHTTPLocationBotClassGetParamsDateRange = "24wControl"
)

type RadarHTTPLocationBotClassGetParamsDeviceType string

const (
	RadarHTTPLocationBotClassGetParamsDeviceTypeDesktop RadarHTTPLocationBotClassGetParamsDeviceType = "DESKTOP"
	RadarHTTPLocationBotClassGetParamsDeviceTypeMobile  RadarHTTPLocationBotClassGetParamsDeviceType = "MOBILE"
	RadarHTTPLocationBotClassGetParamsDeviceTypeOther   RadarHTTPLocationBotClassGetParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPLocationBotClassGetParamsFormat string

const (
	RadarHTTPLocationBotClassGetParamsFormatJson RadarHTTPLocationBotClassGetParamsFormat = "JSON"
	RadarHTTPLocationBotClassGetParamsFormatCsv  RadarHTTPLocationBotClassGetParamsFormat = "CSV"
)

type RadarHTTPLocationBotClassGetParamsHTTPProtocol string

const (
	RadarHTTPLocationBotClassGetParamsHTTPProtocolHTTP  RadarHTTPLocationBotClassGetParamsHTTPProtocol = "HTTP"
	RadarHTTPLocationBotClassGetParamsHTTPProtocolHTTPS RadarHTTPLocationBotClassGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPLocationBotClassGetParamsHTTPVersion string

const (
	RadarHTTPLocationBotClassGetParamsHTTPVersionHttPv1 RadarHTTPLocationBotClassGetParamsHTTPVersion = "HTTPv1"
	RadarHTTPLocationBotClassGetParamsHTTPVersionHttPv2 RadarHTTPLocationBotClassGetParamsHTTPVersion = "HTTPv2"
	RadarHTTPLocationBotClassGetParamsHTTPVersionHttPv3 RadarHTTPLocationBotClassGetParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPLocationBotClassGetParamsIPVersion string

const (
	RadarHTTPLocationBotClassGetParamsIPVersionIPv4 RadarHTTPLocationBotClassGetParamsIPVersion = "IPv4"
	RadarHTTPLocationBotClassGetParamsIPVersionIPv6 RadarHTTPLocationBotClassGetParamsIPVersion = "IPv6"
)

type RadarHTTPLocationBotClassGetParamsOS string

const (
	RadarHTTPLocationBotClassGetParamsOSWindows  RadarHTTPLocationBotClassGetParamsOS = "WINDOWS"
	RadarHTTPLocationBotClassGetParamsOSMacosx   RadarHTTPLocationBotClassGetParamsOS = "MACOSX"
	RadarHTTPLocationBotClassGetParamsOSIos      RadarHTTPLocationBotClassGetParamsOS = "IOS"
	RadarHTTPLocationBotClassGetParamsOSAndroid  RadarHTTPLocationBotClassGetParamsOS = "ANDROID"
	RadarHTTPLocationBotClassGetParamsOSChromeos RadarHTTPLocationBotClassGetParamsOS = "CHROMEOS"
	RadarHTTPLocationBotClassGetParamsOSLinux    RadarHTTPLocationBotClassGetParamsOS = "LINUX"
	RadarHTTPLocationBotClassGetParamsOSSmartTv  RadarHTTPLocationBotClassGetParamsOS = "SMART_TV"
)

type RadarHTTPLocationBotClassGetParamsTLSVersion string

const (
	RadarHTTPLocationBotClassGetParamsTLSVersionTlSv1_0  RadarHTTPLocationBotClassGetParamsTLSVersion = "TLSv1_0"
	RadarHTTPLocationBotClassGetParamsTLSVersionTlSv1_1  RadarHTTPLocationBotClassGetParamsTLSVersion = "TLSv1_1"
	RadarHTTPLocationBotClassGetParamsTLSVersionTlSv1_2  RadarHTTPLocationBotClassGetParamsTLSVersion = "TLSv1_2"
	RadarHTTPLocationBotClassGetParamsTLSVersionTlSv1_3  RadarHTTPLocationBotClassGetParamsTLSVersion = "TLSv1_3"
	RadarHTTPLocationBotClassGetParamsTLSVersionTlSvQuic RadarHTTPLocationBotClassGetParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPLocationBotClassGetResponseEnvelope struct {
	Result  RadarHTTPLocationBotClassGetResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarHTTPLocationBotClassGetResponseEnvelopeJSON `json:"-"`
}

// radarHTTPLocationBotClassGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarHTTPLocationBotClassGetResponseEnvelope]
type radarHTTPLocationBotClassGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationBotClassGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
