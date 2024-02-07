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

// RadarHTTPTLSVersionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarHTTPTLSVersionService]
// method instead.
type RadarHTTPTLSVersionService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTLSVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarHTTPTLSVersionService(opts ...option.RequestOption) (r *RadarHTTPTLSVersionService) {
	r = &RadarHTTPTLSVersionService{}
	r.Options = opts
	return
}

// Get a time series of the percentage distribution of traffic per TLS protocol
// version.
func (r *RadarHTTPTLSVersionService) List(ctx context.Context, query RadarHTTPTLSVersionListParams, opts ...option.RequestOption) (res *RadarHttptlsVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/tls_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHttptlsVersionListResponse struct {
	Result  RadarHttptlsVersionListResponseResult `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    radarHttptlsVersionListResponseJSON   `json:"-"`
}

// radarHttptlsVersionListResponseJSON contains the JSON metadata for the struct
// [RadarHttptlsVersionListResponse]
type radarHttptlsVersionListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHttptlsVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHttptlsVersionListResponseResult struct {
	Meta   interface{}                                 `json:"meta,required"`
	Serie0 RadarHttptlsVersionListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHttptlsVersionListResponseResultJSON   `json:"-"`
}

// radarHttptlsVersionListResponseResultJSON contains the JSON metadata for the
// struct [RadarHttptlsVersionListResponseResult]
type radarHttptlsVersionListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHttptlsVersionListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHttptlsVersionListResponseResultSerie0 struct {
	Timestamps []string                                        `json:"timestamps,required"`
	TLS1_0     []string                                        `json:"TLS 1.0,required"`
	TLS1_1     []string                                        `json:"TLS 1.1,required"`
	TLS1_2     []string                                        `json:"TLS 1.2,required"`
	TLS1_3     []string                                        `json:"TLS 1.3,required"`
	TLSQuic    []string                                        `json:"TLS QUIC,required"`
	JSON       radarHttptlsVersionListResponseResultSerie0JSON `json:"-"`
}

// radarHttptlsVersionListResponseResultSerie0JSON contains the JSON metadata for
// the struct [RadarHttptlsVersionListResponseResultSerie0]
type radarHttptlsVersionListResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	TLS1_0      apijson.Field
	TLS1_1      apijson.Field
	TLS1_2      apijson.Field
	TLS1_3      apijson.Field
	TLSQuic     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHttptlsVersionListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTLSVersionListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHttptlsVersionListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHttptlsVersionListParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHttptlsVersionListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHttptlsVersionListParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHttptlsVersionListParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHttptlsVersionListParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHttptlsVersionListParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHttptlsVersionListParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHttptlsVersionListParamsO] `query:"os"`
}

// URLQuery serializes [RadarHTTPTLSVersionListParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPTLSVersionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHttptlsVersionListParamsAggInterval string

const (
	RadarHttptlsVersionListParamsAggInterval15m RadarHttptlsVersionListParamsAggInterval = "15m"
	RadarHttptlsVersionListParamsAggInterval1h  RadarHttptlsVersionListParamsAggInterval = "1h"
	RadarHttptlsVersionListParamsAggInterval1d  RadarHttptlsVersionListParamsAggInterval = "1d"
	RadarHttptlsVersionListParamsAggInterval1w  RadarHttptlsVersionListParamsAggInterval = "1w"
)

type RadarHttptlsVersionListParamsBotClass string

const (
	RadarHttptlsVersionListParamsBotClassLikelyAutomated RadarHttptlsVersionListParamsBotClass = "LIKELY_AUTOMATED"
	RadarHttptlsVersionListParamsBotClassLikelyHuman     RadarHttptlsVersionListParamsBotClass = "LIKELY_HUMAN"
)

type RadarHttptlsVersionListParamsDateRange string

const (
	RadarHttptlsVersionListParamsDateRange1d         RadarHttptlsVersionListParamsDateRange = "1d"
	RadarHttptlsVersionListParamsDateRange2d         RadarHttptlsVersionListParamsDateRange = "2d"
	RadarHttptlsVersionListParamsDateRange7d         RadarHttptlsVersionListParamsDateRange = "7d"
	RadarHttptlsVersionListParamsDateRange14d        RadarHttptlsVersionListParamsDateRange = "14d"
	RadarHttptlsVersionListParamsDateRange28d        RadarHttptlsVersionListParamsDateRange = "28d"
	RadarHttptlsVersionListParamsDateRange12w        RadarHttptlsVersionListParamsDateRange = "12w"
	RadarHttptlsVersionListParamsDateRange24w        RadarHttptlsVersionListParamsDateRange = "24w"
	RadarHttptlsVersionListParamsDateRange52w        RadarHttptlsVersionListParamsDateRange = "52w"
	RadarHttptlsVersionListParamsDateRange1dControl  RadarHttptlsVersionListParamsDateRange = "1dControl"
	RadarHttptlsVersionListParamsDateRange2dControl  RadarHttptlsVersionListParamsDateRange = "2dControl"
	RadarHttptlsVersionListParamsDateRange7dControl  RadarHttptlsVersionListParamsDateRange = "7dControl"
	RadarHttptlsVersionListParamsDateRange14dControl RadarHttptlsVersionListParamsDateRange = "14dControl"
	RadarHttptlsVersionListParamsDateRange28dControl RadarHttptlsVersionListParamsDateRange = "28dControl"
	RadarHttptlsVersionListParamsDateRange12wControl RadarHttptlsVersionListParamsDateRange = "12wControl"
	RadarHttptlsVersionListParamsDateRange24wControl RadarHttptlsVersionListParamsDateRange = "24wControl"
)

type RadarHttptlsVersionListParamsDeviceType string

const (
	RadarHttptlsVersionListParamsDeviceTypeDesktop RadarHttptlsVersionListParamsDeviceType = "DESKTOP"
	RadarHttptlsVersionListParamsDeviceTypeMobile  RadarHttptlsVersionListParamsDeviceType = "MOBILE"
	RadarHttptlsVersionListParamsDeviceTypeOther   RadarHttptlsVersionListParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHttptlsVersionListParamsFormat string

const (
	RadarHttptlsVersionListParamsFormatJson RadarHttptlsVersionListParamsFormat = "JSON"
	RadarHttptlsVersionListParamsFormatCsv  RadarHttptlsVersionListParamsFormat = "CSV"
)

type RadarHttptlsVersionListParamsHTTPProtocol string

const (
	RadarHttptlsVersionListParamsHTTPProtocolHTTP  RadarHttptlsVersionListParamsHTTPProtocol = "HTTP"
	RadarHttptlsVersionListParamsHTTPProtocolHTTPs RadarHttptlsVersionListParamsHTTPProtocol = "HTTPS"
)

type RadarHttptlsVersionListParamsHTTPVersion string

const (
	RadarHttptlsVersionListParamsHTTPVersionHttPv1 RadarHttptlsVersionListParamsHTTPVersion = "HTTPv1"
	RadarHttptlsVersionListParamsHTTPVersionHttPv2 RadarHttptlsVersionListParamsHTTPVersion = "HTTPv2"
	RadarHttptlsVersionListParamsHTTPVersionHttPv3 RadarHttptlsVersionListParamsHTTPVersion = "HTTPv3"
)

type RadarHttptlsVersionListParamsIPVersion string

const (
	RadarHttptlsVersionListParamsIPVersionIPv4 RadarHttptlsVersionListParamsIPVersion = "IPv4"
	RadarHttptlsVersionListParamsIPVersionIPv6 RadarHttptlsVersionListParamsIPVersion = "IPv6"
)

type RadarHttptlsVersionListParamsO string

const (
	RadarHttptlsVersionListParamsOWindows  RadarHttptlsVersionListParamsO = "WINDOWS"
	RadarHttptlsVersionListParamsOMacosx   RadarHttptlsVersionListParamsO = "MACOSX"
	RadarHttptlsVersionListParamsOIos      RadarHttptlsVersionListParamsO = "IOS"
	RadarHttptlsVersionListParamsOAndroid  RadarHttptlsVersionListParamsO = "ANDROID"
	RadarHttptlsVersionListParamsOChromeos RadarHttptlsVersionListParamsO = "CHROMEOS"
	RadarHttptlsVersionListParamsOLinux    RadarHttptlsVersionListParamsO = "LINUX"
	RadarHttptlsVersionListParamsOSmartTv  RadarHttptlsVersionListParamsO = "SMART_TV"
)
