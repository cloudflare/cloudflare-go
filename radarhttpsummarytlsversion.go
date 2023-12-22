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

// RadarHTTPSummaryTlsVersionService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarHTTPSummaryTlsVersionService] method instead.
type RadarHTTPSummaryTlsVersionService struct {
	Options []option.RequestOption
}

// NewRadarHTTPSummaryTlsVersionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPSummaryTlsVersionService(opts ...option.RequestOption) (r *RadarHTTPSummaryTlsVersionService) {
	r = &RadarHTTPSummaryTlsVersionService{}
	r.Options = opts
	return
}

// Percentage distribution of traffic per TLS protocol version.
func (r *RadarHTTPSummaryTlsVersionService) List(ctx context.Context, query RadarHTTPSummaryTlsVersionListParams, opts ...option.RequestOption) (res *RadarHTTPSummaryTlsVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/summary/tls_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPSummaryTlsVersionListResponse struct {
	Result  RadarHTTPSummaryTlsVersionListResponseResult `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarHTTPSummaryTlsVersionListResponseJSON   `json:"-"`
}

// radarHTTPSummaryTlsVersionListResponseJSON contains the JSON metadata for the
// struct [RadarHTTPSummaryTlsVersionListResponse]
type radarHTTPSummaryTlsVersionListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPSummaryTlsVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPSummaryTlsVersionListResponseResult struct {
	Meta     interface{}                                          `json:"meta,required"`
	Summary0 RadarHTTPSummaryTlsVersionListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarHTTPSummaryTlsVersionListResponseResultJSON     `json:"-"`
}

// radarHTTPSummaryTlsVersionListResponseResultJSON contains the JSON metadata for
// the struct [RadarHTTPSummaryTlsVersionListResponseResult]
type radarHTTPSummaryTlsVersionListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPSummaryTlsVersionListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPSummaryTlsVersionListResponseResultSummary0 struct {
	Tls1_0  string                                                   `json:"TLS 1.0,required"`
	Tls1_1  string                                                   `json:"TLS 1.1,required"`
	Tls1_2  string                                                   `json:"TLS 1.2,required"`
	Tls1_3  string                                                   `json:"TLS 1.3,required"`
	TlsQuic string                                                   `json:"TLS QUIC,required"`
	JSON    radarHTTPSummaryTlsVersionListResponseResultSummary0JSON `json:"-"`
}

// radarHTTPSummaryTlsVersionListResponseResultSummary0JSON contains the JSON
// metadata for the struct [RadarHTTPSummaryTlsVersionListResponseResultSummary0]
type radarHTTPSummaryTlsVersionListResponseResultSummary0JSON struct {
	Tls1_0      apijson.Field
	Tls1_1      apijson.Field
	Tls1_2      apijson.Field
	Tls1_3      apijson.Field
	TlsQuic     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPSummaryTlsVersionListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPSummaryTlsVersionListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPSummaryTlsVersionListParamsBotClass] `query:"botClass"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPSummaryTlsVersionListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPSummaryTlsVersionListParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPSummaryTlsVersionListParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPSummaryTlsVersionListParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPSummaryTlsVersionListParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPSummaryTlsVersionListParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPSummaryTlsVersionListParamsO] `query:"os"`
}

// URLQuery serializes [RadarHTTPSummaryTlsVersionListParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPSummaryTlsVersionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarHTTPSummaryTlsVersionListParamsBotClass string

const (
	RadarHTTPSummaryTlsVersionListParamsBotClassLikelyAutomated RadarHTTPSummaryTlsVersionListParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPSummaryTlsVersionListParamsBotClassLikelyHuman     RadarHTTPSummaryTlsVersionListParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPSummaryTlsVersionListParamsDateRange string

const (
	RadarHTTPSummaryTlsVersionListParamsDateRange1d         RadarHTTPSummaryTlsVersionListParamsDateRange = "1d"
	RadarHTTPSummaryTlsVersionListParamsDateRange7d         RadarHTTPSummaryTlsVersionListParamsDateRange = "7d"
	RadarHTTPSummaryTlsVersionListParamsDateRange14d        RadarHTTPSummaryTlsVersionListParamsDateRange = "14d"
	RadarHTTPSummaryTlsVersionListParamsDateRange28d        RadarHTTPSummaryTlsVersionListParamsDateRange = "28d"
	RadarHTTPSummaryTlsVersionListParamsDateRange12w        RadarHTTPSummaryTlsVersionListParamsDateRange = "12w"
	RadarHTTPSummaryTlsVersionListParamsDateRange24w        RadarHTTPSummaryTlsVersionListParamsDateRange = "24w"
	RadarHTTPSummaryTlsVersionListParamsDateRange52w        RadarHTTPSummaryTlsVersionListParamsDateRange = "52w"
	RadarHTTPSummaryTlsVersionListParamsDateRange1dControl  RadarHTTPSummaryTlsVersionListParamsDateRange = "1dControl"
	RadarHTTPSummaryTlsVersionListParamsDateRange7dControl  RadarHTTPSummaryTlsVersionListParamsDateRange = "7dControl"
	RadarHTTPSummaryTlsVersionListParamsDateRange14dControl RadarHTTPSummaryTlsVersionListParamsDateRange = "14dControl"
	RadarHTTPSummaryTlsVersionListParamsDateRange28dControl RadarHTTPSummaryTlsVersionListParamsDateRange = "28dControl"
	RadarHTTPSummaryTlsVersionListParamsDateRange12wControl RadarHTTPSummaryTlsVersionListParamsDateRange = "12wControl"
	RadarHTTPSummaryTlsVersionListParamsDateRange24wControl RadarHTTPSummaryTlsVersionListParamsDateRange = "24wControl"
)

type RadarHTTPSummaryTlsVersionListParamsDeviceType string

const (
	RadarHTTPSummaryTlsVersionListParamsDeviceTypeDesktop RadarHTTPSummaryTlsVersionListParamsDeviceType = "DESKTOP"
	RadarHTTPSummaryTlsVersionListParamsDeviceTypeMobile  RadarHTTPSummaryTlsVersionListParamsDeviceType = "MOBILE"
	RadarHTTPSummaryTlsVersionListParamsDeviceTypeOther   RadarHTTPSummaryTlsVersionListParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPSummaryTlsVersionListParamsFormat string

const (
	RadarHTTPSummaryTlsVersionListParamsFormatJson RadarHTTPSummaryTlsVersionListParamsFormat = "JSON"
	RadarHTTPSummaryTlsVersionListParamsFormatCsv  RadarHTTPSummaryTlsVersionListParamsFormat = "CSV"
)

type RadarHTTPSummaryTlsVersionListParamsHTTPProtocol string

const (
	RadarHTTPSummaryTlsVersionListParamsHTTPProtocolHTTP  RadarHTTPSummaryTlsVersionListParamsHTTPProtocol = "HTTP"
	RadarHTTPSummaryTlsVersionListParamsHTTPProtocolHTTPs RadarHTTPSummaryTlsVersionListParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPSummaryTlsVersionListParamsHTTPVersion string

const (
	RadarHTTPSummaryTlsVersionListParamsHTTPVersionHttPv1 RadarHTTPSummaryTlsVersionListParamsHTTPVersion = "HTTPv1"
	RadarHTTPSummaryTlsVersionListParamsHTTPVersionHttPv2 RadarHTTPSummaryTlsVersionListParamsHTTPVersion = "HTTPv2"
	RadarHTTPSummaryTlsVersionListParamsHTTPVersionHttPv3 RadarHTTPSummaryTlsVersionListParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPSummaryTlsVersionListParamsIPVersion string

const (
	RadarHTTPSummaryTlsVersionListParamsIPVersionIPv4 RadarHTTPSummaryTlsVersionListParamsIPVersion = "IPv4"
	RadarHTTPSummaryTlsVersionListParamsIPVersionIPv6 RadarHTTPSummaryTlsVersionListParamsIPVersion = "IPv6"
)

type RadarHTTPSummaryTlsVersionListParamsO string

const (
	RadarHTTPSummaryTlsVersionListParamsOWindows  RadarHTTPSummaryTlsVersionListParamsO = "WINDOWS"
	RadarHTTPSummaryTlsVersionListParamsOMacosx   RadarHTTPSummaryTlsVersionListParamsO = "MACOSX"
	RadarHTTPSummaryTlsVersionListParamsOAndroid  RadarHTTPSummaryTlsVersionListParamsO = "ANDROID"
	RadarHTTPSummaryTlsVersionListParamsOChromeos RadarHTTPSummaryTlsVersionListParamsO = "CHROMEOS"
	RadarHTTPSummaryTlsVersionListParamsOLinux    RadarHTTPSummaryTlsVersionListParamsO = "LINUX"
	RadarHTTPSummaryTlsVersionListParamsOSmartTv  RadarHTTPSummaryTlsVersionListParamsO = "SMART_TV"
)
