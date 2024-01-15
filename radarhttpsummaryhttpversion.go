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

// RadarHTTPSummaryHTTPVersionService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarHTTPSummaryHTTPVersionService] method instead.
type RadarHTTPSummaryHTTPVersionService struct {
	Options []option.RequestOption
}

// NewRadarHTTPSummaryHTTPVersionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPSummaryHTTPVersionService(opts ...option.RequestOption) (r *RadarHTTPSummaryHTTPVersionService) {
	r = &RadarHTTPSummaryHTTPVersionService{}
	r.Options = opts
	return
}

// Percentage distribution of traffic per HTTP protocol version over a given time
// period.
func (r *RadarHTTPSummaryHTTPVersionService) List(ctx context.Context, query RadarHTTPSummaryHTTPVersionListParams, opts ...option.RequestOption) (res *RadarHTTPSummaryHTTPVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/summary/http_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPSummaryHTTPVersionListResponse struct {
	Result  RadarHTTPSummaryHTTPVersionListResponseResult `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    radarHTTPSummaryHTTPVersionListResponseJSON   `json:"-"`
}

// radarHTTPSummaryHTTPVersionListResponseJSON contains the JSON metadata for the
// struct [RadarHTTPSummaryHTTPVersionListResponse]
type radarHTTPSummaryHTTPVersionListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPSummaryHTTPVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPSummaryHTTPVersionListResponseResult struct {
	Meta     RadarHTTPSummaryHTTPVersionListResponseResultMeta     `json:"meta,required"`
	Summary0 RadarHTTPSummaryHTTPVersionListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarHTTPSummaryHTTPVersionListResponseResultJSON     `json:"-"`
}

// radarHTTPSummaryHTTPVersionListResponseResultJSON contains the JSON metadata for
// the struct [RadarHTTPSummaryHTTPVersionListResponseResult]
type radarHTTPSummaryHTTPVersionListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPSummaryHTTPVersionListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPSummaryHTTPVersionListResponseResultMeta struct {
	DateRange      []RadarHTTPSummaryHTTPVersionListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                          `json:"lastUpdated,required"`
	Normalization  string                                                          `json:"normalization,required"`
	ConfidenceInfo RadarHTTPSummaryHTTPVersionListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPSummaryHTTPVersionListResponseResultMetaJSON           `json:"-"`
}

// radarHTTPSummaryHTTPVersionListResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarHTTPSummaryHTTPVersionListResponseResultMeta]
type radarHTTPSummaryHTTPVersionListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPSummaryHTTPVersionListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPSummaryHTTPVersionListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                      `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPSummaryHTTPVersionListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPSummaryHTTPVersionListResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarHTTPSummaryHTTPVersionListResponseResultMetaDateRange]
type radarHTTPSummaryHTTPVersionListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPSummaryHTTPVersionListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPSummaryHTTPVersionListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPSummaryHTTPVersionListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                       `json:"level"`
	JSON        radarHTTPSummaryHTTPVersionListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPSummaryHTTPVersionListResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarHTTPSummaryHTTPVersionListResponseResultMetaConfidenceInfo]
type radarHTTPSummaryHTTPVersionListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPSummaryHTTPVersionListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPSummaryHTTPVersionListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                        `json:"dataSource,required"`
	Description     string                                                                        `json:"description,required"`
	EventType       string                                                                        `json:"eventType,required"`
	IsInstantaneous interface{}                                                                   `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                     `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                        `json:"linkedUrl"`
	StartTime       time.Time                                                                     `json:"startTime" format:"date-time"`
	JSON            radarHTTPSummaryHTTPVersionListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPSummaryHTTPVersionListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPSummaryHTTPVersionListResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPSummaryHTTPVersionListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPSummaryHTTPVersionListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPSummaryHTTPVersionListResponseResultSummary0 struct {
	HTTP1X string                                                    `json:"HTTP/1.x,required"`
	HTTP2  string                                                    `json:"HTTP/2,required"`
	HTTP3  string                                                    `json:"HTTP/3,required"`
	JSON   radarHTTPSummaryHTTPVersionListResponseResultSummary0JSON `json:"-"`
}

// radarHTTPSummaryHTTPVersionListResponseResultSummary0JSON contains the JSON
// metadata for the struct [RadarHTTPSummaryHTTPVersionListResponseResultSummary0]
type radarHTTPSummaryHTTPVersionListResponseResultSummary0JSON struct {
	HTTP1X      apijson.Field
	HTTP2       apijson.Field
	HTTP3       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPSummaryHTTPVersionListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPSummaryHTTPVersionListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPSummaryHTTPVersionListParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPSummaryHTTPVersionListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPSummaryHTTPVersionListParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPSummaryHTTPVersionListParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPSummaryHTTPVersionListParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPSummaryHTTPVersionListParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPSummaryHTTPVersionListParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPSummaryHTTPVersionListParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPSummaryHTTPVersionListParams]'s query parameters
// as `url.Values`.
func (r RadarHTTPSummaryHTTPVersionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarHTTPSummaryHTTPVersionListParamsBotClass string

const (
	RadarHTTPSummaryHTTPVersionListParamsBotClassLikelyAutomated RadarHTTPSummaryHTTPVersionListParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPSummaryHTTPVersionListParamsBotClassLikelyHuman     RadarHTTPSummaryHTTPVersionListParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPSummaryHTTPVersionListParamsDateRange string

const (
	RadarHTTPSummaryHTTPVersionListParamsDateRange1d         RadarHTTPSummaryHTTPVersionListParamsDateRange = "1d"
	RadarHTTPSummaryHTTPVersionListParamsDateRange2d         RadarHTTPSummaryHTTPVersionListParamsDateRange = "2d"
	RadarHTTPSummaryHTTPVersionListParamsDateRange7d         RadarHTTPSummaryHTTPVersionListParamsDateRange = "7d"
	RadarHTTPSummaryHTTPVersionListParamsDateRange14d        RadarHTTPSummaryHTTPVersionListParamsDateRange = "14d"
	RadarHTTPSummaryHTTPVersionListParamsDateRange28d        RadarHTTPSummaryHTTPVersionListParamsDateRange = "28d"
	RadarHTTPSummaryHTTPVersionListParamsDateRange12w        RadarHTTPSummaryHTTPVersionListParamsDateRange = "12w"
	RadarHTTPSummaryHTTPVersionListParamsDateRange24w        RadarHTTPSummaryHTTPVersionListParamsDateRange = "24w"
	RadarHTTPSummaryHTTPVersionListParamsDateRange52w        RadarHTTPSummaryHTTPVersionListParamsDateRange = "52w"
	RadarHTTPSummaryHTTPVersionListParamsDateRange1dControl  RadarHTTPSummaryHTTPVersionListParamsDateRange = "1dControl"
	RadarHTTPSummaryHTTPVersionListParamsDateRange2dControl  RadarHTTPSummaryHTTPVersionListParamsDateRange = "2dControl"
	RadarHTTPSummaryHTTPVersionListParamsDateRange7dControl  RadarHTTPSummaryHTTPVersionListParamsDateRange = "7dControl"
	RadarHTTPSummaryHTTPVersionListParamsDateRange14dControl RadarHTTPSummaryHTTPVersionListParamsDateRange = "14dControl"
	RadarHTTPSummaryHTTPVersionListParamsDateRange28dControl RadarHTTPSummaryHTTPVersionListParamsDateRange = "28dControl"
	RadarHTTPSummaryHTTPVersionListParamsDateRange12wControl RadarHTTPSummaryHTTPVersionListParamsDateRange = "12wControl"
	RadarHTTPSummaryHTTPVersionListParamsDateRange24wControl RadarHTTPSummaryHTTPVersionListParamsDateRange = "24wControl"
)

type RadarHTTPSummaryHTTPVersionListParamsDeviceType string

const (
	RadarHTTPSummaryHTTPVersionListParamsDeviceTypeDesktop RadarHTTPSummaryHTTPVersionListParamsDeviceType = "DESKTOP"
	RadarHTTPSummaryHTTPVersionListParamsDeviceTypeMobile  RadarHTTPSummaryHTTPVersionListParamsDeviceType = "MOBILE"
	RadarHTTPSummaryHTTPVersionListParamsDeviceTypeOther   RadarHTTPSummaryHTTPVersionListParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPSummaryHTTPVersionListParamsFormat string

const (
	RadarHTTPSummaryHTTPVersionListParamsFormatJson RadarHTTPSummaryHTTPVersionListParamsFormat = "JSON"
	RadarHTTPSummaryHTTPVersionListParamsFormatCsv  RadarHTTPSummaryHTTPVersionListParamsFormat = "CSV"
)

type RadarHTTPSummaryHTTPVersionListParamsHTTPProtocol string

const (
	RadarHTTPSummaryHTTPVersionListParamsHTTPProtocolHTTP  RadarHTTPSummaryHTTPVersionListParamsHTTPProtocol = "HTTP"
	RadarHTTPSummaryHTTPVersionListParamsHTTPProtocolHTTPs RadarHTTPSummaryHTTPVersionListParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPSummaryHTTPVersionListParamsIPVersion string

const (
	RadarHTTPSummaryHTTPVersionListParamsIPVersionIPv4 RadarHTTPSummaryHTTPVersionListParamsIPVersion = "IPv4"
	RadarHTTPSummaryHTTPVersionListParamsIPVersionIPv6 RadarHTTPSummaryHTTPVersionListParamsIPVersion = "IPv6"
)

type RadarHTTPSummaryHTTPVersionListParamsO string

const (
	RadarHTTPSummaryHTTPVersionListParamsOWindows  RadarHTTPSummaryHTTPVersionListParamsO = "WINDOWS"
	RadarHTTPSummaryHTTPVersionListParamsOMacosx   RadarHTTPSummaryHTTPVersionListParamsO = "MACOSX"
	RadarHTTPSummaryHTTPVersionListParamsOIos      RadarHTTPSummaryHTTPVersionListParamsO = "IOS"
	RadarHTTPSummaryHTTPVersionListParamsOAndroid  RadarHTTPSummaryHTTPVersionListParamsO = "ANDROID"
	RadarHTTPSummaryHTTPVersionListParamsOChromeos RadarHTTPSummaryHTTPVersionListParamsO = "CHROMEOS"
	RadarHTTPSummaryHTTPVersionListParamsOLinux    RadarHTTPSummaryHTTPVersionListParamsO = "LINUX"
	RadarHTTPSummaryHTTPVersionListParamsOSmartTv  RadarHTTPSummaryHTTPVersionListParamsO = "SMART_TV"
)

type RadarHTTPSummaryHTTPVersionListParamsTlsVersion string

const (
	RadarHTTPSummaryHTTPVersionListParamsTlsVersionTlSv1_0  RadarHTTPSummaryHTTPVersionListParamsTlsVersion = "TLSv1_0"
	RadarHTTPSummaryHTTPVersionListParamsTlsVersionTlSv1_1  RadarHTTPSummaryHTTPVersionListParamsTlsVersion = "TLSv1_1"
	RadarHTTPSummaryHTTPVersionListParamsTlsVersionTlSv1_2  RadarHTTPSummaryHTTPVersionListParamsTlsVersion = "TLSv1_2"
	RadarHTTPSummaryHTTPVersionListParamsTlsVersionTlSv1_3  RadarHTTPSummaryHTTPVersionListParamsTlsVersion = "TLSv1_3"
	RadarHTTPSummaryHTTPVersionListParamsTlsVersionTlSvQuic RadarHTTPSummaryHTTPVersionListParamsTlsVersion = "TLSvQUIC"
)
