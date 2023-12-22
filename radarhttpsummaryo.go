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

// RadarHTTPSummaryOService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarHTTPSummaryOService] method
// instead.
type RadarHTTPSummaryOService struct {
	Options []option.RequestOption
}

// NewRadarHTTPSummaryOService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarHTTPSummaryOService(opts ...option.RequestOption) (r *RadarHTTPSummaryOService) {
	r = &RadarHTTPSummaryOService{}
	r.Options = opts
	return
}

// Percentage distribution of traffic per operating system.
func (r *RadarHTTPSummaryOService) List(ctx context.Context, query RadarHTTPSummaryOListParams, opts ...option.RequestOption) (res *RadarHTTPSummaryOListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/summary/os"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPSummaryOListResponse struct {
	Result  RadarHTTPSummaryOListResponseResult `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    radarHTTPSummaryOListResponseJSON   `json:"-"`
}

// radarHTTPSummaryOListResponseJSON contains the JSON metadata for the struct
// [RadarHTTPSummaryOListResponse]
type radarHTTPSummaryOListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPSummaryOListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPSummaryOListResponseResult struct {
	Meta     interface{}                                 `json:"meta,required"`
	Summary0 RadarHTTPSummaryOListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarHTTPSummaryOListResponseResultJSON     `json:"-"`
}

// radarHTTPSummaryOListResponseResultJSON contains the JSON metadata for the
// struct [RadarHTTPSummaryOListResponseResult]
type radarHTTPSummaryOListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPSummaryOListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPSummaryOListResponseResultSummary0 struct {
	Android string                                          `json:"ANDROID,required"`
	Ios     string                                          `json:"IOS,required"`
	JSON    radarHTTPSummaryOListResponseResultSummary0JSON `json:"-"`
}

// radarHTTPSummaryOListResponseResultSummary0JSON contains the JSON metadata for
// the struct [RadarHTTPSummaryOListResponseResultSummary0]
type radarHTTPSummaryOListResponseResultSummary0JSON struct {
	Android     apijson.Field
	Ios         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPSummaryOListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPSummaryOListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPSummaryOListParamsBotClass] `query:"botClass"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPSummaryOListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPSummaryOListParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPSummaryOListParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPSummaryOListParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPSummaryOListParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPSummaryOListParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPSummaryOListParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPSummaryOListParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPSummaryOListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarHTTPSummaryOListParamsBotClass string

const (
	RadarHTTPSummaryOListParamsBotClassLikelyAutomated RadarHTTPSummaryOListParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPSummaryOListParamsBotClassLikelyHuman     RadarHTTPSummaryOListParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPSummaryOListParamsDateRange string

const (
	RadarHTTPSummaryOListParamsDateRange1d         RadarHTTPSummaryOListParamsDateRange = "1d"
	RadarHTTPSummaryOListParamsDateRange7d         RadarHTTPSummaryOListParamsDateRange = "7d"
	RadarHTTPSummaryOListParamsDateRange14d        RadarHTTPSummaryOListParamsDateRange = "14d"
	RadarHTTPSummaryOListParamsDateRange28d        RadarHTTPSummaryOListParamsDateRange = "28d"
	RadarHTTPSummaryOListParamsDateRange12w        RadarHTTPSummaryOListParamsDateRange = "12w"
	RadarHTTPSummaryOListParamsDateRange24w        RadarHTTPSummaryOListParamsDateRange = "24w"
	RadarHTTPSummaryOListParamsDateRange52w        RadarHTTPSummaryOListParamsDateRange = "52w"
	RadarHTTPSummaryOListParamsDateRange1dControl  RadarHTTPSummaryOListParamsDateRange = "1dControl"
	RadarHTTPSummaryOListParamsDateRange7dControl  RadarHTTPSummaryOListParamsDateRange = "7dControl"
	RadarHTTPSummaryOListParamsDateRange14dControl RadarHTTPSummaryOListParamsDateRange = "14dControl"
	RadarHTTPSummaryOListParamsDateRange28dControl RadarHTTPSummaryOListParamsDateRange = "28dControl"
	RadarHTTPSummaryOListParamsDateRange12wControl RadarHTTPSummaryOListParamsDateRange = "12wControl"
	RadarHTTPSummaryOListParamsDateRange24wControl RadarHTTPSummaryOListParamsDateRange = "24wControl"
)

type RadarHTTPSummaryOListParamsDeviceType string

const (
	RadarHTTPSummaryOListParamsDeviceTypeDesktop RadarHTTPSummaryOListParamsDeviceType = "DESKTOP"
	RadarHTTPSummaryOListParamsDeviceTypeMobile  RadarHTTPSummaryOListParamsDeviceType = "MOBILE"
	RadarHTTPSummaryOListParamsDeviceTypeOther   RadarHTTPSummaryOListParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPSummaryOListParamsFormat string

const (
	RadarHTTPSummaryOListParamsFormatJson RadarHTTPSummaryOListParamsFormat = "JSON"
	RadarHTTPSummaryOListParamsFormatCsv  RadarHTTPSummaryOListParamsFormat = "CSV"
)

type RadarHTTPSummaryOListParamsHTTPProtocol string

const (
	RadarHTTPSummaryOListParamsHTTPProtocolHTTP  RadarHTTPSummaryOListParamsHTTPProtocol = "HTTP"
	RadarHTTPSummaryOListParamsHTTPProtocolHTTPs RadarHTTPSummaryOListParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPSummaryOListParamsHTTPVersion string

const (
	RadarHTTPSummaryOListParamsHTTPVersionHttPv1 RadarHTTPSummaryOListParamsHTTPVersion = "HTTPv1"
	RadarHTTPSummaryOListParamsHTTPVersionHttPv2 RadarHTTPSummaryOListParamsHTTPVersion = "HTTPv2"
	RadarHTTPSummaryOListParamsHTTPVersionHttPv3 RadarHTTPSummaryOListParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPSummaryOListParamsIPVersion string

const (
	RadarHTTPSummaryOListParamsIPVersionIPv4 RadarHTTPSummaryOListParamsIPVersion = "IPv4"
	RadarHTTPSummaryOListParamsIPVersionIPv6 RadarHTTPSummaryOListParamsIPVersion = "IPv6"
)

type RadarHTTPSummaryOListParamsTlsVersion string

const (
	RadarHTTPSummaryOListParamsTlsVersionTlSv1_0  RadarHTTPSummaryOListParamsTlsVersion = "TLSv1_0"
	RadarHTTPSummaryOListParamsTlsVersionTlSv1_1  RadarHTTPSummaryOListParamsTlsVersion = "TLSv1_1"
	RadarHTTPSummaryOListParamsTlsVersionTlSv1_2  RadarHTTPSummaryOListParamsTlsVersion = "TLSv1_2"
	RadarHTTPSummaryOListParamsTlsVersionTlSv1_3  RadarHTTPSummaryOListParamsTlsVersion = "TLSv1_3"
	RadarHTTPSummaryOListParamsTlsVersionTlSvQuic RadarHTTPSummaryOListParamsTlsVersion = "TLSvQUIC"
)
