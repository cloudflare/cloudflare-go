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

// RadarHTTPTopLocationHTTPVersionService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarHTTPTopLocationHTTPVersionService] method instead.
type RadarHTTPTopLocationHTTPVersionService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTopLocationHTTPVersionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPTopLocationHTTPVersionService(opts ...option.RequestOption) (r *RadarHTTPTopLocationHTTPVersionService) {
	r = &RadarHTTPTopLocationHTTPVersionService{}
	r.Options = opts
	return
}

// Get the top locations, by HTTP traffic, of the requested HTTP protocol. Values
// are a percentage out of the total traffic.
func (r *RadarHTTPTopLocationHTTPVersionService) Get(ctx context.Context, httpVersion RadarHTTPTopLocationHTTPVersionGetParamsHTTPVersion, query RadarHTTPTopLocationHTTPVersionGetParams, opts ...option.RequestOption) (res *RadarHTTPTopLocationHTTPVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/locations/http_version/%v", httpVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTopLocationHTTPVersionGetResponse struct {
	Result  RadarHTTPTopLocationHTTPVersionGetResponseResult `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarHTTPTopLocationHTTPVersionGetResponseJSON   `json:"-"`
}

// radarHTTPTopLocationHTTPVersionGetResponseJSON contains the JSON metadata for
// the struct [RadarHTTPTopLocationHTTPVersionGetResponse]
type radarHTTPTopLocationHTTPVersionGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationHTTPVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationHTTPVersionGetResponseResult struct {
	Meta RadarHTTPTopLocationHTTPVersionGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopLocationHTTPVersionGetResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopLocationHTTPVersionGetResponseResultJSON   `json:"-"`
}

// radarHTTPTopLocationHTTPVersionGetResponseResultJSON contains the JSON metadata
// for the struct [RadarHTTPTopLocationHTTPVersionGetResponseResult]
type radarHTTPTopLocationHTTPVersionGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationHTTPVersionGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationHTTPVersionGetResponseResultMeta struct {
	DateRange      []RadarHTTPTopLocationHTTPVersionGetResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                             `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPTopLocationHTTPVersionGetResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPTopLocationHTTPVersionGetResponseResultMetaJSON           `json:"-"`
}

// radarHTTPTopLocationHTTPVersionGetResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarHTTPTopLocationHTTPVersionGetResponseResultMeta]
type radarHTTPTopLocationHTTPVersionGetResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopLocationHTTPVersionGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationHTTPVersionGetResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                         `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopLocationHTTPVersionGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopLocationHTTPVersionGetResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopLocationHTTPVersionGetResponseResultMetaDateRange]
type radarHTTPTopLocationHTTPVersionGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationHTTPVersionGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationHTTPVersionGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopLocationHTTPVersionGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                          `json:"level"`
	JSON        radarHTTPTopLocationHTTPVersionGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPTopLocationHTTPVersionGetResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarHTTPTopLocationHTTPVersionGetResponseResultMetaConfidenceInfo]
type radarHTTPTopLocationHTTPVersionGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationHTTPVersionGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationHTTPVersionGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                           `json:"dataSource,required"`
	Description     string                                                                           `json:"description,required"`
	EventType       string                                                                           `json:"eventType,required"`
	IsInstantaneous interface{}                                                                      `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                        `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                           `json:"linkedUrl"`
	StartTime       time.Time                                                                        `json:"startTime" format:"date-time"`
	JSON            radarHTTPTopLocationHTTPVersionGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopLocationHTTPVersionGetResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopLocationHTTPVersionGetResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopLocationHTTPVersionGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopLocationHTTPVersionGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationHTTPVersionGetResponseResultTop0 struct {
	ClientCountryAlpha2 string                                                   `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                                   `json:"clientCountryName,required"`
	Value               string                                                   `json:"value,required"`
	JSON                radarHTTPTopLocationHTTPVersionGetResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopLocationHTTPVersionGetResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarHTTPTopLocationHTTPVersionGetResponseResultTop0]
type radarHTTPTopLocationHTTPVersionGetResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarHTTPTopLocationHTTPVersionGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationHTTPVersionGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopLocationHTTPVersionGetParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTopLocationHTTPVersionGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTopLocationHTTPVersionGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTopLocationHTTPVersionGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTopLocationHTTPVersionGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTopLocationHTTPVersionGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTopLocationHTTPVersionGetParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTopLocationHTTPVersionGetParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopLocationHTTPVersionGetParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTopLocationHTTPVersionGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// HTTP version.
type RadarHTTPTopLocationHTTPVersionGetParamsHTTPVersion string

const (
	RadarHTTPTopLocationHTTPVersionGetParamsHTTPVersionHttPv1 RadarHTTPTopLocationHTTPVersionGetParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopLocationHTTPVersionGetParamsHTTPVersionHttPv2 RadarHTTPTopLocationHTTPVersionGetParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopLocationHTTPVersionGetParamsHTTPVersionHttPv3 RadarHTTPTopLocationHTTPVersionGetParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTopLocationHTTPVersionGetParamsBotClass string

const (
	RadarHTTPTopLocationHTTPVersionGetParamsBotClassLikelyAutomated RadarHTTPTopLocationHTTPVersionGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopLocationHTTPVersionGetParamsBotClassLikelyHuman     RadarHTTPTopLocationHTTPVersionGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTopLocationHTTPVersionGetParamsDateRange string

const (
	RadarHTTPTopLocationHTTPVersionGetParamsDateRange1d         RadarHTTPTopLocationHTTPVersionGetParamsDateRange = "1d"
	RadarHTTPTopLocationHTTPVersionGetParamsDateRange2d         RadarHTTPTopLocationHTTPVersionGetParamsDateRange = "2d"
	RadarHTTPTopLocationHTTPVersionGetParamsDateRange7d         RadarHTTPTopLocationHTTPVersionGetParamsDateRange = "7d"
	RadarHTTPTopLocationHTTPVersionGetParamsDateRange14d        RadarHTTPTopLocationHTTPVersionGetParamsDateRange = "14d"
	RadarHTTPTopLocationHTTPVersionGetParamsDateRange28d        RadarHTTPTopLocationHTTPVersionGetParamsDateRange = "28d"
	RadarHTTPTopLocationHTTPVersionGetParamsDateRange12w        RadarHTTPTopLocationHTTPVersionGetParamsDateRange = "12w"
	RadarHTTPTopLocationHTTPVersionGetParamsDateRange24w        RadarHTTPTopLocationHTTPVersionGetParamsDateRange = "24w"
	RadarHTTPTopLocationHTTPVersionGetParamsDateRange52w        RadarHTTPTopLocationHTTPVersionGetParamsDateRange = "52w"
	RadarHTTPTopLocationHTTPVersionGetParamsDateRange1dControl  RadarHTTPTopLocationHTTPVersionGetParamsDateRange = "1dControl"
	RadarHTTPTopLocationHTTPVersionGetParamsDateRange2dControl  RadarHTTPTopLocationHTTPVersionGetParamsDateRange = "2dControl"
	RadarHTTPTopLocationHTTPVersionGetParamsDateRange7dControl  RadarHTTPTopLocationHTTPVersionGetParamsDateRange = "7dControl"
	RadarHTTPTopLocationHTTPVersionGetParamsDateRange14dControl RadarHTTPTopLocationHTTPVersionGetParamsDateRange = "14dControl"
	RadarHTTPTopLocationHTTPVersionGetParamsDateRange28dControl RadarHTTPTopLocationHTTPVersionGetParamsDateRange = "28dControl"
	RadarHTTPTopLocationHTTPVersionGetParamsDateRange12wControl RadarHTTPTopLocationHTTPVersionGetParamsDateRange = "12wControl"
	RadarHTTPTopLocationHTTPVersionGetParamsDateRange24wControl RadarHTTPTopLocationHTTPVersionGetParamsDateRange = "24wControl"
)

type RadarHTTPTopLocationHTTPVersionGetParamsDeviceType string

const (
	RadarHTTPTopLocationHTTPVersionGetParamsDeviceTypeDesktop RadarHTTPTopLocationHTTPVersionGetParamsDeviceType = "DESKTOP"
	RadarHTTPTopLocationHTTPVersionGetParamsDeviceTypeMobile  RadarHTTPTopLocationHTTPVersionGetParamsDeviceType = "MOBILE"
	RadarHTTPTopLocationHTTPVersionGetParamsDeviceTypeOther   RadarHTTPTopLocationHTTPVersionGetParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTopLocationHTTPVersionGetParamsFormat string

const (
	RadarHTTPTopLocationHTTPVersionGetParamsFormatJson RadarHTTPTopLocationHTTPVersionGetParamsFormat = "JSON"
	RadarHTTPTopLocationHTTPVersionGetParamsFormatCsv  RadarHTTPTopLocationHTTPVersionGetParamsFormat = "CSV"
)

type RadarHTTPTopLocationHTTPVersionGetParamsHTTPProtocol string

const (
	RadarHTTPTopLocationHTTPVersionGetParamsHTTPProtocolHTTP  RadarHTTPTopLocationHTTPVersionGetParamsHTTPProtocol = "HTTP"
	RadarHTTPTopLocationHTTPVersionGetParamsHTTPProtocolHTTPs RadarHTTPTopLocationHTTPVersionGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTopLocationHTTPVersionGetParamsIPVersion string

const (
	RadarHTTPTopLocationHTTPVersionGetParamsIPVersionIPv4 RadarHTTPTopLocationHTTPVersionGetParamsIPVersion = "IPv4"
	RadarHTTPTopLocationHTTPVersionGetParamsIPVersionIPv6 RadarHTTPTopLocationHTTPVersionGetParamsIPVersion = "IPv6"
)

type RadarHTTPTopLocationHTTPVersionGetParamsO string

const (
	RadarHTTPTopLocationHTTPVersionGetParamsOWindows  RadarHTTPTopLocationHTTPVersionGetParamsO = "WINDOWS"
	RadarHTTPTopLocationHTTPVersionGetParamsOMacosx   RadarHTTPTopLocationHTTPVersionGetParamsO = "MACOSX"
	RadarHTTPTopLocationHTTPVersionGetParamsOIos      RadarHTTPTopLocationHTTPVersionGetParamsO = "IOS"
	RadarHTTPTopLocationHTTPVersionGetParamsOAndroid  RadarHTTPTopLocationHTTPVersionGetParamsO = "ANDROID"
	RadarHTTPTopLocationHTTPVersionGetParamsOChromeos RadarHTTPTopLocationHTTPVersionGetParamsO = "CHROMEOS"
	RadarHTTPTopLocationHTTPVersionGetParamsOLinux    RadarHTTPTopLocationHTTPVersionGetParamsO = "LINUX"
	RadarHTTPTopLocationHTTPVersionGetParamsOSmartTv  RadarHTTPTopLocationHTTPVersionGetParamsO = "SMART_TV"
)

type RadarHTTPTopLocationHTTPVersionGetParamsTlsVersion string

const (
	RadarHTTPTopLocationHTTPVersionGetParamsTlsVersionTlSv1_0  RadarHTTPTopLocationHTTPVersionGetParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopLocationHTTPVersionGetParamsTlsVersionTlSv1_1  RadarHTTPTopLocationHTTPVersionGetParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopLocationHTTPVersionGetParamsTlsVersionTlSv1_2  RadarHTTPTopLocationHTTPVersionGetParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopLocationHTTPVersionGetParamsTlsVersionTlSv1_3  RadarHTTPTopLocationHTTPVersionGetParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopLocationHTTPVersionGetParamsTlsVersionTlSvQuic RadarHTTPTopLocationHTTPVersionGetParamsTlsVersion = "TLSvQUIC"
)
