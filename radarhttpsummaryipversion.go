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

// RadarHTTPSummaryIPVersionService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarHTTPSummaryIPVersionService] method instead.
type RadarHTTPSummaryIPVersionService struct {
	Options []option.RequestOption
}

// NewRadarHTTPSummaryIPVersionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPSummaryIPVersionService(opts ...option.RequestOption) (r *RadarHTTPSummaryIPVersionService) {
	r = &RadarHTTPSummaryIPVersionService{}
	r.Options = opts
	return
}

// Percentage distribution of traffic per IP protocol version.
func (r *RadarHTTPSummaryIPVersionService) List(ctx context.Context, query RadarHTTPSummaryIPVersionListParams, opts ...option.RequestOption) (res *RadarHTTPSummaryIPVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/summary/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPSummaryIPVersionListResponse struct {
	Result  RadarHTTPSummaryIPVersionListResponseResult `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    radarHTTPSummaryIPVersionListResponseJSON   `json:"-"`
}

// radarHTTPSummaryIPVersionListResponseJSON contains the JSON metadata for the
// struct [RadarHTTPSummaryIPVersionListResponse]
type radarHTTPSummaryIPVersionListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPSummaryIPVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPSummaryIPVersionListResponseResult struct {
	Meta     interface{}                                         `json:"meta,required"`
	Summary0 RadarHTTPSummaryIPVersionListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarHTTPSummaryIPVersionListResponseResultJSON     `json:"-"`
}

// radarHTTPSummaryIPVersionListResponseResultJSON contains the JSON metadata for
// the struct [RadarHTTPSummaryIPVersionListResponseResult]
type radarHTTPSummaryIPVersionListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPSummaryIPVersionListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPSummaryIPVersionListResponseResultSummary0 struct {
	IPv4 string                                                  `json:"IPv4,required"`
	IPv6 string                                                  `json:"IPv6,required"`
	JSON radarHTTPSummaryIPVersionListResponseResultSummary0JSON `json:"-"`
}

// radarHTTPSummaryIPVersionListResponseResultSummary0JSON contains the JSON
// metadata for the struct [RadarHTTPSummaryIPVersionListResponseResultSummary0]
type radarHTTPSummaryIPVersionListResponseResultSummary0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPSummaryIPVersionListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPSummaryIPVersionListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPSummaryIPVersionListParamsBotClass] `query:"botClass"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPSummaryIPVersionListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPSummaryIPVersionListParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPSummaryIPVersionListParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPSummaryIPVersionListParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPSummaryIPVersionListParamsHTTPVersion] `query:"httpVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPSummaryIPVersionListParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPSummaryIPVersionListParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPSummaryIPVersionListParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPSummaryIPVersionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarHTTPSummaryIPVersionListParamsBotClass string

const (
	RadarHTTPSummaryIPVersionListParamsBotClassLikelyAutomated RadarHTTPSummaryIPVersionListParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPSummaryIPVersionListParamsBotClassLikelyHuman     RadarHTTPSummaryIPVersionListParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPSummaryIPVersionListParamsDateRange string

const (
	RadarHTTPSummaryIPVersionListParamsDateRange1d         RadarHTTPSummaryIPVersionListParamsDateRange = "1d"
	RadarHTTPSummaryIPVersionListParamsDateRange7d         RadarHTTPSummaryIPVersionListParamsDateRange = "7d"
	RadarHTTPSummaryIPVersionListParamsDateRange14d        RadarHTTPSummaryIPVersionListParamsDateRange = "14d"
	RadarHTTPSummaryIPVersionListParamsDateRange28d        RadarHTTPSummaryIPVersionListParamsDateRange = "28d"
	RadarHTTPSummaryIPVersionListParamsDateRange12w        RadarHTTPSummaryIPVersionListParamsDateRange = "12w"
	RadarHTTPSummaryIPVersionListParamsDateRange24w        RadarHTTPSummaryIPVersionListParamsDateRange = "24w"
	RadarHTTPSummaryIPVersionListParamsDateRange52w        RadarHTTPSummaryIPVersionListParamsDateRange = "52w"
	RadarHTTPSummaryIPVersionListParamsDateRange1dControl  RadarHTTPSummaryIPVersionListParamsDateRange = "1dControl"
	RadarHTTPSummaryIPVersionListParamsDateRange7dControl  RadarHTTPSummaryIPVersionListParamsDateRange = "7dControl"
	RadarHTTPSummaryIPVersionListParamsDateRange14dControl RadarHTTPSummaryIPVersionListParamsDateRange = "14dControl"
	RadarHTTPSummaryIPVersionListParamsDateRange28dControl RadarHTTPSummaryIPVersionListParamsDateRange = "28dControl"
	RadarHTTPSummaryIPVersionListParamsDateRange12wControl RadarHTTPSummaryIPVersionListParamsDateRange = "12wControl"
	RadarHTTPSummaryIPVersionListParamsDateRange24wControl RadarHTTPSummaryIPVersionListParamsDateRange = "24wControl"
)

type RadarHTTPSummaryIPVersionListParamsDeviceType string

const (
	RadarHTTPSummaryIPVersionListParamsDeviceTypeDesktop RadarHTTPSummaryIPVersionListParamsDeviceType = "DESKTOP"
	RadarHTTPSummaryIPVersionListParamsDeviceTypeMobile  RadarHTTPSummaryIPVersionListParamsDeviceType = "MOBILE"
	RadarHTTPSummaryIPVersionListParamsDeviceTypeOther   RadarHTTPSummaryIPVersionListParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPSummaryIPVersionListParamsFormat string

const (
	RadarHTTPSummaryIPVersionListParamsFormatJson RadarHTTPSummaryIPVersionListParamsFormat = "JSON"
	RadarHTTPSummaryIPVersionListParamsFormatCsv  RadarHTTPSummaryIPVersionListParamsFormat = "CSV"
)

type RadarHTTPSummaryIPVersionListParamsHTTPProtocol string

const (
	RadarHTTPSummaryIPVersionListParamsHTTPProtocolHTTP  RadarHTTPSummaryIPVersionListParamsHTTPProtocol = "HTTP"
	RadarHTTPSummaryIPVersionListParamsHTTPProtocolHTTPs RadarHTTPSummaryIPVersionListParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPSummaryIPVersionListParamsHTTPVersion string

const (
	RadarHTTPSummaryIPVersionListParamsHTTPVersionHttPv1 RadarHTTPSummaryIPVersionListParamsHTTPVersion = "HTTPv1"
	RadarHTTPSummaryIPVersionListParamsHTTPVersionHttPv2 RadarHTTPSummaryIPVersionListParamsHTTPVersion = "HTTPv2"
	RadarHTTPSummaryIPVersionListParamsHTTPVersionHttPv3 RadarHTTPSummaryIPVersionListParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPSummaryIPVersionListParamsO string

const (
	RadarHTTPSummaryIPVersionListParamsOWindows  RadarHTTPSummaryIPVersionListParamsO = "WINDOWS"
	RadarHTTPSummaryIPVersionListParamsOMacosx   RadarHTTPSummaryIPVersionListParamsO = "MACOSX"
	RadarHTTPSummaryIPVersionListParamsOAndroid  RadarHTTPSummaryIPVersionListParamsO = "ANDROID"
	RadarHTTPSummaryIPVersionListParamsOChromeos RadarHTTPSummaryIPVersionListParamsO = "CHROMEOS"
	RadarHTTPSummaryIPVersionListParamsOLinux    RadarHTTPSummaryIPVersionListParamsO = "LINUX"
	RadarHTTPSummaryIPVersionListParamsOSmartTv  RadarHTTPSummaryIPVersionListParamsO = "SMART_TV"
)

type RadarHTTPSummaryIPVersionListParamsTlsVersion string

const (
	RadarHTTPSummaryIPVersionListParamsTlsVersionTlSv1_0  RadarHTTPSummaryIPVersionListParamsTlsVersion = "TLSv1_0"
	RadarHTTPSummaryIPVersionListParamsTlsVersionTlSv1_1  RadarHTTPSummaryIPVersionListParamsTlsVersion = "TLSv1_1"
	RadarHTTPSummaryIPVersionListParamsTlsVersionTlSv1_2  RadarHTTPSummaryIPVersionListParamsTlsVersion = "TLSv1_2"
	RadarHTTPSummaryIPVersionListParamsTlsVersionTlSv1_3  RadarHTTPSummaryIPVersionListParamsTlsVersion = "TLSv1_3"
	RadarHTTPSummaryIPVersionListParamsTlsVersionTlSvQuic RadarHTTPSummaryIPVersionListParamsTlsVersion = "TLSvQUIC"
)
