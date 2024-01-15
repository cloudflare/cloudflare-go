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

// RadarHTTPSummaryBotClassService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarHTTPSummaryBotClassService] method instead.
type RadarHTTPSummaryBotClassService struct {
	Options []option.RequestOption
}

// NewRadarHTTPSummaryBotClassService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPSummaryBotClassService(opts ...option.RequestOption) (r *RadarHTTPSummaryBotClassService) {
	r = &RadarHTTPSummaryBotClassService{}
	r.Options = opts
	return
}

// Percentage distribution of bot-generated traffic to genuine human traffic, as
// classified by Cloudflare. Visit
// https://developers.cloudflare.com/radar/concepts/bot-classes/ for more
// information.
func (r *RadarHTTPSummaryBotClassService) List(ctx context.Context, query RadarHTTPSummaryBotClassListParams, opts ...option.RequestOption) (res *RadarHTTPSummaryBotClassListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/summary/bot_class"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPSummaryBotClassListResponse struct {
	Result  RadarHTTPSummaryBotClassListResponseResult `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarHTTPSummaryBotClassListResponseJSON   `json:"-"`
}

// radarHTTPSummaryBotClassListResponseJSON contains the JSON metadata for the
// struct [RadarHTTPSummaryBotClassListResponse]
type radarHTTPSummaryBotClassListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPSummaryBotClassListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPSummaryBotClassListResponseResult struct {
	Meta     RadarHTTPSummaryBotClassListResponseResultMeta     `json:"meta,required"`
	Summary0 RadarHTTPSummaryBotClassListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarHTTPSummaryBotClassListResponseResultJSON     `json:"-"`
}

// radarHTTPSummaryBotClassListResponseResultJSON contains the JSON metadata for
// the struct [RadarHTTPSummaryBotClassListResponseResult]
type radarHTTPSummaryBotClassListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPSummaryBotClassListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPSummaryBotClassListResponseResultMeta struct {
	DateRange      []RadarHTTPSummaryBotClassListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                       `json:"lastUpdated,required"`
	Normalization  string                                                       `json:"normalization,required"`
	ConfidenceInfo RadarHTTPSummaryBotClassListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPSummaryBotClassListResponseResultMetaJSON           `json:"-"`
}

// radarHTTPSummaryBotClassListResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarHTTPSummaryBotClassListResponseResultMeta]
type radarHTTPSummaryBotClassListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPSummaryBotClassListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPSummaryBotClassListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                   `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPSummaryBotClassListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPSummaryBotClassListResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarHTTPSummaryBotClassListResponseResultMetaDateRange]
type radarHTTPSummaryBotClassListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPSummaryBotClassListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPSummaryBotClassListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPSummaryBotClassListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                    `json:"level"`
	JSON        radarHTTPSummaryBotClassListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPSummaryBotClassListResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarHTTPSummaryBotClassListResponseResultMetaConfidenceInfo]
type radarHTTPSummaryBotClassListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPSummaryBotClassListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPSummaryBotClassListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                     `json:"dataSource,required"`
	Description     string                                                                     `json:"description,required"`
	EventType       string                                                                     `json:"eventType,required"`
	IsInstantaneous interface{}                                                                `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                  `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                     `json:"linkedUrl"`
	StartTime       time.Time                                                                  `json:"startTime" format:"date-time"`
	JSON            radarHTTPSummaryBotClassListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPSummaryBotClassListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPSummaryBotClassListResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPSummaryBotClassListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPSummaryBotClassListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPSummaryBotClassListResponseResultSummary0 struct {
	Bot   string                                                 `json:"bot,required"`
	Human string                                                 `json:"human,required"`
	JSON  radarHTTPSummaryBotClassListResponseResultSummary0JSON `json:"-"`
}

// radarHTTPSummaryBotClassListResponseResultSummary0JSON contains the JSON
// metadata for the struct [RadarHTTPSummaryBotClassListResponseResultSummary0]
type radarHTTPSummaryBotClassListResponseResultSummary0JSON struct {
	Bot         apijson.Field
	Human       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPSummaryBotClassListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPSummaryBotClassListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPSummaryBotClassListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPSummaryBotClassListParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPSummaryBotClassListParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPSummaryBotClassListParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPSummaryBotClassListParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPSummaryBotClassListParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPSummaryBotClassListParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPSummaryBotClassListParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPSummaryBotClassListParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPSummaryBotClassListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarHTTPSummaryBotClassListParamsDateRange string

const (
	RadarHTTPSummaryBotClassListParamsDateRange1d         RadarHTTPSummaryBotClassListParamsDateRange = "1d"
	RadarHTTPSummaryBotClassListParamsDateRange2d         RadarHTTPSummaryBotClassListParamsDateRange = "2d"
	RadarHTTPSummaryBotClassListParamsDateRange7d         RadarHTTPSummaryBotClassListParamsDateRange = "7d"
	RadarHTTPSummaryBotClassListParamsDateRange14d        RadarHTTPSummaryBotClassListParamsDateRange = "14d"
	RadarHTTPSummaryBotClassListParamsDateRange28d        RadarHTTPSummaryBotClassListParamsDateRange = "28d"
	RadarHTTPSummaryBotClassListParamsDateRange12w        RadarHTTPSummaryBotClassListParamsDateRange = "12w"
	RadarHTTPSummaryBotClassListParamsDateRange24w        RadarHTTPSummaryBotClassListParamsDateRange = "24w"
	RadarHTTPSummaryBotClassListParamsDateRange52w        RadarHTTPSummaryBotClassListParamsDateRange = "52w"
	RadarHTTPSummaryBotClassListParamsDateRange1dControl  RadarHTTPSummaryBotClassListParamsDateRange = "1dControl"
	RadarHTTPSummaryBotClassListParamsDateRange2dControl  RadarHTTPSummaryBotClassListParamsDateRange = "2dControl"
	RadarHTTPSummaryBotClassListParamsDateRange7dControl  RadarHTTPSummaryBotClassListParamsDateRange = "7dControl"
	RadarHTTPSummaryBotClassListParamsDateRange14dControl RadarHTTPSummaryBotClassListParamsDateRange = "14dControl"
	RadarHTTPSummaryBotClassListParamsDateRange28dControl RadarHTTPSummaryBotClassListParamsDateRange = "28dControl"
	RadarHTTPSummaryBotClassListParamsDateRange12wControl RadarHTTPSummaryBotClassListParamsDateRange = "12wControl"
	RadarHTTPSummaryBotClassListParamsDateRange24wControl RadarHTTPSummaryBotClassListParamsDateRange = "24wControl"
)

type RadarHTTPSummaryBotClassListParamsDeviceType string

const (
	RadarHTTPSummaryBotClassListParamsDeviceTypeDesktop RadarHTTPSummaryBotClassListParamsDeviceType = "DESKTOP"
	RadarHTTPSummaryBotClassListParamsDeviceTypeMobile  RadarHTTPSummaryBotClassListParamsDeviceType = "MOBILE"
	RadarHTTPSummaryBotClassListParamsDeviceTypeOther   RadarHTTPSummaryBotClassListParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPSummaryBotClassListParamsFormat string

const (
	RadarHTTPSummaryBotClassListParamsFormatJson RadarHTTPSummaryBotClassListParamsFormat = "JSON"
	RadarHTTPSummaryBotClassListParamsFormatCsv  RadarHTTPSummaryBotClassListParamsFormat = "CSV"
)

type RadarHTTPSummaryBotClassListParamsHTTPProtocol string

const (
	RadarHTTPSummaryBotClassListParamsHTTPProtocolHTTP  RadarHTTPSummaryBotClassListParamsHTTPProtocol = "HTTP"
	RadarHTTPSummaryBotClassListParamsHTTPProtocolHTTPs RadarHTTPSummaryBotClassListParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPSummaryBotClassListParamsHTTPVersion string

const (
	RadarHTTPSummaryBotClassListParamsHTTPVersionHttPv1 RadarHTTPSummaryBotClassListParamsHTTPVersion = "HTTPv1"
	RadarHTTPSummaryBotClassListParamsHTTPVersionHttPv2 RadarHTTPSummaryBotClassListParamsHTTPVersion = "HTTPv2"
	RadarHTTPSummaryBotClassListParamsHTTPVersionHttPv3 RadarHTTPSummaryBotClassListParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPSummaryBotClassListParamsIPVersion string

const (
	RadarHTTPSummaryBotClassListParamsIPVersionIPv4 RadarHTTPSummaryBotClassListParamsIPVersion = "IPv4"
	RadarHTTPSummaryBotClassListParamsIPVersionIPv6 RadarHTTPSummaryBotClassListParamsIPVersion = "IPv6"
)

type RadarHTTPSummaryBotClassListParamsO string

const (
	RadarHTTPSummaryBotClassListParamsOWindows  RadarHTTPSummaryBotClassListParamsO = "WINDOWS"
	RadarHTTPSummaryBotClassListParamsOMacosx   RadarHTTPSummaryBotClassListParamsO = "MACOSX"
	RadarHTTPSummaryBotClassListParamsOIos      RadarHTTPSummaryBotClassListParamsO = "IOS"
	RadarHTTPSummaryBotClassListParamsOAndroid  RadarHTTPSummaryBotClassListParamsO = "ANDROID"
	RadarHTTPSummaryBotClassListParamsOChromeos RadarHTTPSummaryBotClassListParamsO = "CHROMEOS"
	RadarHTTPSummaryBotClassListParamsOLinux    RadarHTTPSummaryBotClassListParamsO = "LINUX"
	RadarHTTPSummaryBotClassListParamsOSmartTv  RadarHTTPSummaryBotClassListParamsO = "SMART_TV"
)

type RadarHTTPSummaryBotClassListParamsTlsVersion string

const (
	RadarHTTPSummaryBotClassListParamsTlsVersionTlSv1_0  RadarHTTPSummaryBotClassListParamsTlsVersion = "TLSv1_0"
	RadarHTTPSummaryBotClassListParamsTlsVersionTlSv1_1  RadarHTTPSummaryBotClassListParamsTlsVersion = "TLSv1_1"
	RadarHTTPSummaryBotClassListParamsTlsVersionTlSv1_2  RadarHTTPSummaryBotClassListParamsTlsVersion = "TLSv1_2"
	RadarHTTPSummaryBotClassListParamsTlsVersionTlSv1_3  RadarHTTPSummaryBotClassListParamsTlsVersion = "TLSv1_3"
	RadarHTTPSummaryBotClassListParamsTlsVersionTlSvQuic RadarHTTPSummaryBotClassListParamsTlsVersion = "TLSvQUIC"
)
