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

// RadarHTTPSummaryHTTPProtocolService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarHTTPSummaryHTTPProtocolService] method instead.
type RadarHTTPSummaryHTTPProtocolService struct {
	Options []option.RequestOption
}

// NewRadarHTTPSummaryHTTPProtocolService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPSummaryHTTPProtocolService(opts ...option.RequestOption) (r *RadarHTTPSummaryHTTPProtocolService) {
	r = &RadarHTTPSummaryHTTPProtocolService{}
	r.Options = opts
	return
}

// Percentage distribution of traffic per HTTP protocol.
func (r *RadarHTTPSummaryHTTPProtocolService) List(ctx context.Context, query RadarHTTPSummaryHTTPProtocolListParams, opts ...option.RequestOption) (res *RadarHTTPSummaryHTTPProtocolListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/summary/http_protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPSummaryHTTPProtocolListResponse struct {
	Result  RadarHTTPSummaryHTTPProtocolListResponseResult `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    radarHTTPSummaryHTTPProtocolListResponseJSON   `json:"-"`
}

// radarHTTPSummaryHTTPProtocolListResponseJSON contains the JSON metadata for the
// struct [RadarHTTPSummaryHTTPProtocolListResponse]
type radarHTTPSummaryHTTPProtocolListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPSummaryHTTPProtocolListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPSummaryHTTPProtocolListResponseResult struct {
	Meta     interface{}                                            `json:"meta,required"`
	Summary0 RadarHTTPSummaryHTTPProtocolListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarHTTPSummaryHTTPProtocolListResponseResultJSON     `json:"-"`
}

// radarHTTPSummaryHTTPProtocolListResponseResultJSON contains the JSON metadata
// for the struct [RadarHTTPSummaryHTTPProtocolListResponseResult]
type radarHTTPSummaryHTTPProtocolListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPSummaryHTTPProtocolListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPSummaryHTTPProtocolListResponseResultSummary0 struct {
	HTTP  string                                                     `json:"http,required"`
	HTTPs string                                                     `json:"https,required"`
	JSON  radarHTTPSummaryHTTPProtocolListResponseResultSummary0JSON `json:"-"`
}

// radarHTTPSummaryHTTPProtocolListResponseResultSummary0JSON contains the JSON
// metadata for the struct [RadarHTTPSummaryHTTPProtocolListResponseResultSummary0]
type radarHTTPSummaryHTTPProtocolListResponseResultSummary0JSON struct {
	HTTP        apijson.Field
	HTTPs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPSummaryHTTPProtocolListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPSummaryHTTPProtocolListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPSummaryHTTPProtocolListParamsBotClass] `query:"botClass"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPSummaryHTTPProtocolListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPSummaryHTTPProtocolListParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPSummaryHTTPProtocolListParamsFormat] `query:"format"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPSummaryHTTPProtocolListParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPSummaryHTTPProtocolListParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPSummaryHTTPProtocolListParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPSummaryHTTPProtocolListParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPSummaryHTTPProtocolListParams]'s query parameters
// as `url.Values`.
func (r RadarHTTPSummaryHTTPProtocolListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarHTTPSummaryHTTPProtocolListParamsBotClass string

const (
	RadarHTTPSummaryHTTPProtocolListParamsBotClassLikelyAutomated RadarHTTPSummaryHTTPProtocolListParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPSummaryHTTPProtocolListParamsBotClassLikelyHuman     RadarHTTPSummaryHTTPProtocolListParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPSummaryHTTPProtocolListParamsDateRange string

const (
	RadarHTTPSummaryHTTPProtocolListParamsDateRange1d         RadarHTTPSummaryHTTPProtocolListParamsDateRange = "1d"
	RadarHTTPSummaryHTTPProtocolListParamsDateRange7d         RadarHTTPSummaryHTTPProtocolListParamsDateRange = "7d"
	RadarHTTPSummaryHTTPProtocolListParamsDateRange14d        RadarHTTPSummaryHTTPProtocolListParamsDateRange = "14d"
	RadarHTTPSummaryHTTPProtocolListParamsDateRange28d        RadarHTTPSummaryHTTPProtocolListParamsDateRange = "28d"
	RadarHTTPSummaryHTTPProtocolListParamsDateRange12w        RadarHTTPSummaryHTTPProtocolListParamsDateRange = "12w"
	RadarHTTPSummaryHTTPProtocolListParamsDateRange24w        RadarHTTPSummaryHTTPProtocolListParamsDateRange = "24w"
	RadarHTTPSummaryHTTPProtocolListParamsDateRange52w        RadarHTTPSummaryHTTPProtocolListParamsDateRange = "52w"
	RadarHTTPSummaryHTTPProtocolListParamsDateRange1dControl  RadarHTTPSummaryHTTPProtocolListParamsDateRange = "1dControl"
	RadarHTTPSummaryHTTPProtocolListParamsDateRange7dControl  RadarHTTPSummaryHTTPProtocolListParamsDateRange = "7dControl"
	RadarHTTPSummaryHTTPProtocolListParamsDateRange14dControl RadarHTTPSummaryHTTPProtocolListParamsDateRange = "14dControl"
	RadarHTTPSummaryHTTPProtocolListParamsDateRange28dControl RadarHTTPSummaryHTTPProtocolListParamsDateRange = "28dControl"
	RadarHTTPSummaryHTTPProtocolListParamsDateRange12wControl RadarHTTPSummaryHTTPProtocolListParamsDateRange = "12wControl"
	RadarHTTPSummaryHTTPProtocolListParamsDateRange24wControl RadarHTTPSummaryHTTPProtocolListParamsDateRange = "24wControl"
)

type RadarHTTPSummaryHTTPProtocolListParamsDeviceType string

const (
	RadarHTTPSummaryHTTPProtocolListParamsDeviceTypeDesktop RadarHTTPSummaryHTTPProtocolListParamsDeviceType = "DESKTOP"
	RadarHTTPSummaryHTTPProtocolListParamsDeviceTypeMobile  RadarHTTPSummaryHTTPProtocolListParamsDeviceType = "MOBILE"
	RadarHTTPSummaryHTTPProtocolListParamsDeviceTypeOther   RadarHTTPSummaryHTTPProtocolListParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPSummaryHTTPProtocolListParamsFormat string

const (
	RadarHTTPSummaryHTTPProtocolListParamsFormatJson RadarHTTPSummaryHTTPProtocolListParamsFormat = "JSON"
	RadarHTTPSummaryHTTPProtocolListParamsFormatCsv  RadarHTTPSummaryHTTPProtocolListParamsFormat = "CSV"
)

type RadarHTTPSummaryHTTPProtocolListParamsHTTPVersion string

const (
	RadarHTTPSummaryHTTPProtocolListParamsHTTPVersionHttPv1 RadarHTTPSummaryHTTPProtocolListParamsHTTPVersion = "HTTPv1"
	RadarHTTPSummaryHTTPProtocolListParamsHTTPVersionHttPv2 RadarHTTPSummaryHTTPProtocolListParamsHTTPVersion = "HTTPv2"
	RadarHTTPSummaryHTTPProtocolListParamsHTTPVersionHttPv3 RadarHTTPSummaryHTTPProtocolListParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPSummaryHTTPProtocolListParamsIPVersion string

const (
	RadarHTTPSummaryHTTPProtocolListParamsIPVersionIPv4 RadarHTTPSummaryHTTPProtocolListParamsIPVersion = "IPv4"
	RadarHTTPSummaryHTTPProtocolListParamsIPVersionIPv6 RadarHTTPSummaryHTTPProtocolListParamsIPVersion = "IPv6"
)

type RadarHTTPSummaryHTTPProtocolListParamsO string

const (
	RadarHTTPSummaryHTTPProtocolListParamsOWindows  RadarHTTPSummaryHTTPProtocolListParamsO = "WINDOWS"
	RadarHTTPSummaryHTTPProtocolListParamsOMacosx   RadarHTTPSummaryHTTPProtocolListParamsO = "MACOSX"
	RadarHTTPSummaryHTTPProtocolListParamsOAndroid  RadarHTTPSummaryHTTPProtocolListParamsO = "ANDROID"
	RadarHTTPSummaryHTTPProtocolListParamsOChromeos RadarHTTPSummaryHTTPProtocolListParamsO = "CHROMEOS"
	RadarHTTPSummaryHTTPProtocolListParamsOLinux    RadarHTTPSummaryHTTPProtocolListParamsO = "LINUX"
	RadarHTTPSummaryHTTPProtocolListParamsOSmartTv  RadarHTTPSummaryHTTPProtocolListParamsO = "SMART_TV"
)

type RadarHTTPSummaryHTTPProtocolListParamsTlsVersion string

const (
	RadarHTTPSummaryHTTPProtocolListParamsTlsVersionTlSv1_0  RadarHTTPSummaryHTTPProtocolListParamsTlsVersion = "TLSv1_0"
	RadarHTTPSummaryHTTPProtocolListParamsTlsVersionTlSv1_1  RadarHTTPSummaryHTTPProtocolListParamsTlsVersion = "TLSv1_1"
	RadarHTTPSummaryHTTPProtocolListParamsTlsVersionTlSv1_2  RadarHTTPSummaryHTTPProtocolListParamsTlsVersion = "TLSv1_2"
	RadarHTTPSummaryHTTPProtocolListParamsTlsVersionTlSv1_3  RadarHTTPSummaryHTTPProtocolListParamsTlsVersion = "TLSv1_3"
	RadarHTTPSummaryHTTPProtocolListParamsTlsVersionTlSvQuic RadarHTTPSummaryHTTPProtocolListParamsTlsVersion = "TLSvQUIC"
)
