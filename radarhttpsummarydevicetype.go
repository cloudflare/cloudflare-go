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

// RadarHTTPSummaryDeviceTypeService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarHTTPSummaryDeviceTypeService] method instead.
type RadarHTTPSummaryDeviceTypeService struct {
	Options []option.RequestOption
}

// NewRadarHTTPSummaryDeviceTypeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPSummaryDeviceTypeService(opts ...option.RequestOption) (r *RadarHTTPSummaryDeviceTypeService) {
	r = &RadarHTTPSummaryDeviceTypeService{}
	r.Options = opts
	return
}

// Percentage distribution of traffic per device type.
func (r *RadarHTTPSummaryDeviceTypeService) List(ctx context.Context, query RadarHTTPSummaryDeviceTypeListParams, opts ...option.RequestOption) (res *RadarHTTPSummaryDeviceTypeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/summary/device_type"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPSummaryDeviceTypeListResponse struct {
	Result  RadarHTTPSummaryDeviceTypeListResponseResult `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarHTTPSummaryDeviceTypeListResponseJSON   `json:"-"`
}

// radarHTTPSummaryDeviceTypeListResponseJSON contains the JSON metadata for the
// struct [RadarHTTPSummaryDeviceTypeListResponse]
type radarHTTPSummaryDeviceTypeListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPSummaryDeviceTypeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPSummaryDeviceTypeListResponseResult struct {
	Meta     interface{}                                          `json:"meta,required"`
	Summary0 RadarHTTPSummaryDeviceTypeListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarHTTPSummaryDeviceTypeListResponseResultJSON     `json:"-"`
}

// radarHTTPSummaryDeviceTypeListResponseResultJSON contains the JSON metadata for
// the struct [RadarHTTPSummaryDeviceTypeListResponseResult]
type radarHTTPSummaryDeviceTypeListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPSummaryDeviceTypeListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPSummaryDeviceTypeListResponseResultSummary0 struct {
	Desktop string                                                   `json:"desktop,required"`
	Mobile  string                                                   `json:"mobile,required"`
	Other   string                                                   `json:"other,required"`
	JSON    radarHTTPSummaryDeviceTypeListResponseResultSummary0JSON `json:"-"`
}

// radarHTTPSummaryDeviceTypeListResponseResultSummary0JSON contains the JSON
// metadata for the struct [RadarHTTPSummaryDeviceTypeListResponseResultSummary0]
type radarHTTPSummaryDeviceTypeListResponseResultSummary0JSON struct {
	Desktop     apijson.Field
	Mobile      apijson.Field
	Other       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPSummaryDeviceTypeListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPSummaryDeviceTypeListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPSummaryDeviceTypeListParamsBotClass] `query:"botClass"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPSummaryDeviceTypeListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarHTTPSummaryDeviceTypeListParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPSummaryDeviceTypeListParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPSummaryDeviceTypeListParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPSummaryDeviceTypeListParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPSummaryDeviceTypeListParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPSummaryDeviceTypeListParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPSummaryDeviceTypeListParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPSummaryDeviceTypeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarHTTPSummaryDeviceTypeListParamsBotClass string

const (
	RadarHTTPSummaryDeviceTypeListParamsBotClassLikelyAutomated RadarHTTPSummaryDeviceTypeListParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPSummaryDeviceTypeListParamsBotClassLikelyHuman     RadarHTTPSummaryDeviceTypeListParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPSummaryDeviceTypeListParamsDateRange string

const (
	RadarHTTPSummaryDeviceTypeListParamsDateRange1d         RadarHTTPSummaryDeviceTypeListParamsDateRange = "1d"
	RadarHTTPSummaryDeviceTypeListParamsDateRange7d         RadarHTTPSummaryDeviceTypeListParamsDateRange = "7d"
	RadarHTTPSummaryDeviceTypeListParamsDateRange14d        RadarHTTPSummaryDeviceTypeListParamsDateRange = "14d"
	RadarHTTPSummaryDeviceTypeListParamsDateRange28d        RadarHTTPSummaryDeviceTypeListParamsDateRange = "28d"
	RadarHTTPSummaryDeviceTypeListParamsDateRange12w        RadarHTTPSummaryDeviceTypeListParamsDateRange = "12w"
	RadarHTTPSummaryDeviceTypeListParamsDateRange24w        RadarHTTPSummaryDeviceTypeListParamsDateRange = "24w"
	RadarHTTPSummaryDeviceTypeListParamsDateRange52w        RadarHTTPSummaryDeviceTypeListParamsDateRange = "52w"
	RadarHTTPSummaryDeviceTypeListParamsDateRange1dControl  RadarHTTPSummaryDeviceTypeListParamsDateRange = "1dControl"
	RadarHTTPSummaryDeviceTypeListParamsDateRange7dControl  RadarHTTPSummaryDeviceTypeListParamsDateRange = "7dControl"
	RadarHTTPSummaryDeviceTypeListParamsDateRange14dControl RadarHTTPSummaryDeviceTypeListParamsDateRange = "14dControl"
	RadarHTTPSummaryDeviceTypeListParamsDateRange28dControl RadarHTTPSummaryDeviceTypeListParamsDateRange = "28dControl"
	RadarHTTPSummaryDeviceTypeListParamsDateRange12wControl RadarHTTPSummaryDeviceTypeListParamsDateRange = "12wControl"
	RadarHTTPSummaryDeviceTypeListParamsDateRange24wControl RadarHTTPSummaryDeviceTypeListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarHTTPSummaryDeviceTypeListParamsFormat string

const (
	RadarHTTPSummaryDeviceTypeListParamsFormatJson RadarHTTPSummaryDeviceTypeListParamsFormat = "JSON"
	RadarHTTPSummaryDeviceTypeListParamsFormatCsv  RadarHTTPSummaryDeviceTypeListParamsFormat = "CSV"
)

type RadarHTTPSummaryDeviceTypeListParamsHTTPProtocol string

const (
	RadarHTTPSummaryDeviceTypeListParamsHTTPProtocolHTTP  RadarHTTPSummaryDeviceTypeListParamsHTTPProtocol = "HTTP"
	RadarHTTPSummaryDeviceTypeListParamsHTTPProtocolHTTPs RadarHTTPSummaryDeviceTypeListParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPSummaryDeviceTypeListParamsHTTPVersion string

const (
	RadarHTTPSummaryDeviceTypeListParamsHTTPVersionHttPv1 RadarHTTPSummaryDeviceTypeListParamsHTTPVersion = "HTTPv1"
	RadarHTTPSummaryDeviceTypeListParamsHTTPVersionHttPv2 RadarHTTPSummaryDeviceTypeListParamsHTTPVersion = "HTTPv2"
	RadarHTTPSummaryDeviceTypeListParamsHTTPVersionHttPv3 RadarHTTPSummaryDeviceTypeListParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPSummaryDeviceTypeListParamsIPVersion string

const (
	RadarHTTPSummaryDeviceTypeListParamsIPVersionIPv4 RadarHTTPSummaryDeviceTypeListParamsIPVersion = "IPv4"
	RadarHTTPSummaryDeviceTypeListParamsIPVersionIPv6 RadarHTTPSummaryDeviceTypeListParamsIPVersion = "IPv6"
)

type RadarHTTPSummaryDeviceTypeListParamsO string

const (
	RadarHTTPSummaryDeviceTypeListParamsOWindows  RadarHTTPSummaryDeviceTypeListParamsO = "WINDOWS"
	RadarHTTPSummaryDeviceTypeListParamsOMacosx   RadarHTTPSummaryDeviceTypeListParamsO = "MACOSX"
	RadarHTTPSummaryDeviceTypeListParamsOAndroid  RadarHTTPSummaryDeviceTypeListParamsO = "ANDROID"
	RadarHTTPSummaryDeviceTypeListParamsOChromeos RadarHTTPSummaryDeviceTypeListParamsO = "CHROMEOS"
	RadarHTTPSummaryDeviceTypeListParamsOLinux    RadarHTTPSummaryDeviceTypeListParamsO = "LINUX"
	RadarHTTPSummaryDeviceTypeListParamsOSmartTv  RadarHTTPSummaryDeviceTypeListParamsO = "SMART_TV"
)

type RadarHTTPSummaryDeviceTypeListParamsTlsVersion string

const (
	RadarHTTPSummaryDeviceTypeListParamsTlsVersionTlSv1_0  RadarHTTPSummaryDeviceTypeListParamsTlsVersion = "TLSv1_0"
	RadarHTTPSummaryDeviceTypeListParamsTlsVersionTlSv1_1  RadarHTTPSummaryDeviceTypeListParamsTlsVersion = "TLSv1_1"
	RadarHTTPSummaryDeviceTypeListParamsTlsVersionTlSv1_2  RadarHTTPSummaryDeviceTypeListParamsTlsVersion = "TLSv1_2"
	RadarHTTPSummaryDeviceTypeListParamsTlsVersionTlSv1_3  RadarHTTPSummaryDeviceTypeListParamsTlsVersion = "TLSv1_3"
	RadarHTTPSummaryDeviceTypeListParamsTlsVersionTlSvQuic RadarHTTPSummaryDeviceTypeListParamsTlsVersion = "TLSvQUIC"
)
