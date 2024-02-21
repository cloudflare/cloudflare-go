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

// RadarHTTPAseBotClassService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarHTTPAseBotClassService]
// method instead.
type RadarHTTPAseBotClassService struct {
	Options []option.RequestOption
}

// NewRadarHTTPAseBotClassService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarHTTPAseBotClassService(opts ...option.RequestOption) (r *RadarHTTPAseBotClassService) {
	r = &RadarHTTPAseBotClassService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS), by HTTP traffic, of the requested bot
// class. These two categories use Cloudflare's bot score - refer to
// [Bot Scores](https://developers.cloudflare.com/bots/concepts/bot-score) for more
// information. Values are a percentage out of the total traffic.
func (r *RadarHTTPAseBotClassService) Get(ctx context.Context, botClass RadarHTTPAseBotClassGetParamsBotClass, query RadarHTTPAseBotClassGetParams, opts ...option.RequestOption) (res *RadarHTTPAseBotClassGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarHTTPAseBotClassGetResponseEnvelope
	path := fmt.Sprintf("radar/http/top/ases/bot_class/%v", botClass)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarHTTPAseBotClassGetResponse struct {
	Meta RadarHTTPAseBotClassGetResponseMeta   `json:"meta,required"`
	Top0 []RadarHTTPAseBotClassGetResponseTop0 `json:"top_0,required"`
	JSON radarHTTPAseBotClassGetResponseJSON   `json:"-"`
}

// radarHTTPAseBotClassGetResponseJSON contains the JSON metadata for the struct
// [RadarHTTPAseBotClassGetResponse]
type radarHTTPAseBotClassGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseBotClassGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPAseBotClassGetResponseMeta struct {
	DateRange      []RadarHTTPAseBotClassGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                            `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPAseBotClassGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPAseBotClassGetResponseMetaJSON           `json:"-"`
}

// radarHTTPAseBotClassGetResponseMetaJSON contains the JSON metadata for the
// struct [RadarHTTPAseBotClassGetResponseMeta]
type radarHTTPAseBotClassGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPAseBotClassGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPAseBotClassGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                        `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPAseBotClassGetResponseMetaDateRangeJSON `json:"-"`
}

// radarHTTPAseBotClassGetResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [RadarHTTPAseBotClassGetResponseMetaDateRange]
type radarHTTPAseBotClassGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseBotClassGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPAseBotClassGetResponseMetaConfidenceInfo struct {
	Annotations []RadarHTTPAseBotClassGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                         `json:"level"`
	JSON        radarHTTPAseBotClassGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPAseBotClassGetResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [RadarHTTPAseBotClassGetResponseMetaConfidenceInfo]
type radarHTTPAseBotClassGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseBotClassGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPAseBotClassGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                          `json:"dataSource,required"`
	Description     string                                                          `json:"description,required"`
	EventType       string                                                          `json:"eventType,required"`
	IsInstantaneous interface{}                                                     `json:"isInstantaneous,required"`
	EndTime         time.Time                                                       `json:"endTime" format:"date-time"`
	LinkedURL       string                                                          `json:"linkedUrl"`
	StartTime       time.Time                                                       `json:"startTime" format:"date-time"`
	JSON            radarHTTPAseBotClassGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPAseBotClassGetResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarHTTPAseBotClassGetResponseMetaConfidenceInfoAnnotation]
type radarHTTPAseBotClassGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPAseBotClassGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPAseBotClassGetResponseTop0 struct {
	ClientAsn    int64                                   `json:"clientASN,required"`
	ClientAsName string                                  `json:"clientASName,required"`
	Value        string                                  `json:"value,required"`
	JSON         radarHTTPAseBotClassGetResponseTop0JSON `json:"-"`
}

// radarHTTPAseBotClassGetResponseTop0JSON contains the JSON metadata for the
// struct [RadarHTTPAseBotClassGetResponseTop0]
type radarHTTPAseBotClassGetResponseTop0JSON struct {
	ClientAsn    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarHTTPAseBotClassGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPAseBotClassGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPAseBotClassGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPAseBotClassGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPAseBotClassGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPAseBotClassGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPAseBotClassGetParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPAseBotClassGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPAseBotClassGetParamsO] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHTTPAseBotClassGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPAseBotClassGetParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPAseBotClassGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Bot class.
type RadarHTTPAseBotClassGetParamsBotClass string

const (
	RadarHTTPAseBotClassGetParamsBotClassLikelyAutomated RadarHTTPAseBotClassGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPAseBotClassGetParamsBotClassLikelyHuman     RadarHTTPAseBotClassGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPAseBotClassGetParamsDateRange string

const (
	RadarHTTPAseBotClassGetParamsDateRange1d         RadarHTTPAseBotClassGetParamsDateRange = "1d"
	RadarHTTPAseBotClassGetParamsDateRange2d         RadarHTTPAseBotClassGetParamsDateRange = "2d"
	RadarHTTPAseBotClassGetParamsDateRange7d         RadarHTTPAseBotClassGetParamsDateRange = "7d"
	RadarHTTPAseBotClassGetParamsDateRange14d        RadarHTTPAseBotClassGetParamsDateRange = "14d"
	RadarHTTPAseBotClassGetParamsDateRange28d        RadarHTTPAseBotClassGetParamsDateRange = "28d"
	RadarHTTPAseBotClassGetParamsDateRange12w        RadarHTTPAseBotClassGetParamsDateRange = "12w"
	RadarHTTPAseBotClassGetParamsDateRange24w        RadarHTTPAseBotClassGetParamsDateRange = "24w"
	RadarHTTPAseBotClassGetParamsDateRange52w        RadarHTTPAseBotClassGetParamsDateRange = "52w"
	RadarHTTPAseBotClassGetParamsDateRange1dControl  RadarHTTPAseBotClassGetParamsDateRange = "1dControl"
	RadarHTTPAseBotClassGetParamsDateRange2dControl  RadarHTTPAseBotClassGetParamsDateRange = "2dControl"
	RadarHTTPAseBotClassGetParamsDateRange7dControl  RadarHTTPAseBotClassGetParamsDateRange = "7dControl"
	RadarHTTPAseBotClassGetParamsDateRange14dControl RadarHTTPAseBotClassGetParamsDateRange = "14dControl"
	RadarHTTPAseBotClassGetParamsDateRange28dControl RadarHTTPAseBotClassGetParamsDateRange = "28dControl"
	RadarHTTPAseBotClassGetParamsDateRange12wControl RadarHTTPAseBotClassGetParamsDateRange = "12wControl"
	RadarHTTPAseBotClassGetParamsDateRange24wControl RadarHTTPAseBotClassGetParamsDateRange = "24wControl"
)

type RadarHTTPAseBotClassGetParamsDeviceType string

const (
	RadarHTTPAseBotClassGetParamsDeviceTypeDesktop RadarHTTPAseBotClassGetParamsDeviceType = "DESKTOP"
	RadarHTTPAseBotClassGetParamsDeviceTypeMobile  RadarHTTPAseBotClassGetParamsDeviceType = "MOBILE"
	RadarHTTPAseBotClassGetParamsDeviceTypeOther   RadarHTTPAseBotClassGetParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPAseBotClassGetParamsFormat string

const (
	RadarHTTPAseBotClassGetParamsFormatJson RadarHTTPAseBotClassGetParamsFormat = "JSON"
	RadarHTTPAseBotClassGetParamsFormatCsv  RadarHTTPAseBotClassGetParamsFormat = "CSV"
)

type RadarHTTPAseBotClassGetParamsHTTPProtocol string

const (
	RadarHTTPAseBotClassGetParamsHTTPProtocolHTTP  RadarHTTPAseBotClassGetParamsHTTPProtocol = "HTTP"
	RadarHTTPAseBotClassGetParamsHTTPProtocolHTTPS RadarHTTPAseBotClassGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPAseBotClassGetParamsHTTPVersion string

const (
	RadarHTTPAseBotClassGetParamsHTTPVersionHttPv1 RadarHTTPAseBotClassGetParamsHTTPVersion = "HTTPv1"
	RadarHTTPAseBotClassGetParamsHTTPVersionHttPv2 RadarHTTPAseBotClassGetParamsHTTPVersion = "HTTPv2"
	RadarHTTPAseBotClassGetParamsHTTPVersionHttPv3 RadarHTTPAseBotClassGetParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPAseBotClassGetParamsIPVersion string

const (
	RadarHTTPAseBotClassGetParamsIPVersionIPv4 RadarHTTPAseBotClassGetParamsIPVersion = "IPv4"
	RadarHTTPAseBotClassGetParamsIPVersionIPv6 RadarHTTPAseBotClassGetParamsIPVersion = "IPv6"
)

type RadarHTTPAseBotClassGetParamsO string

const (
	RadarHTTPAseBotClassGetParamsOWindows  RadarHTTPAseBotClassGetParamsO = "WINDOWS"
	RadarHTTPAseBotClassGetParamsOMacosx   RadarHTTPAseBotClassGetParamsO = "MACOSX"
	RadarHTTPAseBotClassGetParamsOIos      RadarHTTPAseBotClassGetParamsO = "IOS"
	RadarHTTPAseBotClassGetParamsOAndroid  RadarHTTPAseBotClassGetParamsO = "ANDROID"
	RadarHTTPAseBotClassGetParamsOChromeos RadarHTTPAseBotClassGetParamsO = "CHROMEOS"
	RadarHTTPAseBotClassGetParamsOLinux    RadarHTTPAseBotClassGetParamsO = "LINUX"
	RadarHTTPAseBotClassGetParamsOSmartTv  RadarHTTPAseBotClassGetParamsO = "SMART_TV"
)

type RadarHTTPAseBotClassGetParamsTLSVersion string

const (
	RadarHTTPAseBotClassGetParamsTLSVersionTlSv1_0  RadarHTTPAseBotClassGetParamsTLSVersion = "TLSv1_0"
	RadarHTTPAseBotClassGetParamsTLSVersionTlSv1_1  RadarHTTPAseBotClassGetParamsTLSVersion = "TLSv1_1"
	RadarHTTPAseBotClassGetParamsTLSVersionTlSv1_2  RadarHTTPAseBotClassGetParamsTLSVersion = "TLSv1_2"
	RadarHTTPAseBotClassGetParamsTLSVersionTlSv1_3  RadarHTTPAseBotClassGetParamsTLSVersion = "TLSv1_3"
	RadarHTTPAseBotClassGetParamsTLSVersionTlSvQuic RadarHTTPAseBotClassGetParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPAseBotClassGetResponseEnvelope struct {
	Result  RadarHTTPAseBotClassGetResponse             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    radarHTTPAseBotClassGetResponseEnvelopeJSON `json:"-"`
}

// radarHTTPAseBotClassGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarHTTPAseBotClassGetResponseEnvelope]
type radarHTTPAseBotClassGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseBotClassGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
