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

// RadarHTTPTimeseryBrowserService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarHTTPTimeseryBrowserService] method instead.
type RadarHTTPTimeseryBrowserService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTimeseryBrowserService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPTimeseryBrowserService(opts ...option.RequestOption) (r *RadarHTTPTimeseryBrowserService) {
	r = &RadarHTTPTimeseryBrowserService{}
	r.Options = opts
	return
}

// Percentage distribution of traffic of the top user agents in the selected time
// range, over time.
func (r *RadarHTTPTimeseryBrowserService) List(ctx context.Context, query RadarHTTPTimeseryBrowserListParams, opts ...option.RequestOption) (res *RadarHTTPTimeseryBrowserListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries/browser"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTimeseryBrowserListResponse struct {
	Result  RadarHTTPTimeseryBrowserListResponseResult `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarHTTPTimeseryBrowserListResponseJSON   `json:"-"`
}

// radarHTTPTimeseryBrowserListResponseJSON contains the JSON metadata for the
// struct [RadarHTTPTimeseryBrowserListResponse]
type radarHTTPTimeseryBrowserListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseryBrowserListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseryBrowserListResponseResult struct {
	Meta   interface{}                                      `json:"meta,required"`
	Serie0 RadarHTTPTimeseryBrowserListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseryBrowserListResponseResultJSON   `json:"-"`
}

// radarHTTPTimeseryBrowserListResponseResultJSON contains the JSON metadata for
// the struct [RadarHTTPTimeseryBrowserListResponseResult]
type radarHTTPTimeseryBrowserListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseryBrowserListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseryBrowserListResponseResultSerie0 struct {
	BrowserName []string                                             `json:"<browser name>,required"`
	Timestamps  []string                                             `json:"timestamps,required"`
	JSON        radarHTTPTimeseryBrowserListResponseResultSerie0JSON `json:"-"`
}

// radarHTTPTimeseryBrowserListResponseResultSerie0JSON contains the JSON metadata
// for the struct [RadarHTTPTimeseryBrowserListResponseResultSerie0]
type radarHTTPTimeseryBrowserListResponseResultSerie0JSON struct {
	BrowserName apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseryBrowserListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseryBrowserListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseryBrowserListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTimeseryBrowserListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTimeseryBrowserListParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTimeseryBrowserListParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTimeseryBrowserListParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTimeseryBrowserListParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTimeseryBrowserListParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects (eg browsers) to the top items over the time range.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTimeseryBrowserListParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTimeseryBrowserListParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseryBrowserListParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPTimeseryBrowserListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseryBrowserListParamsAggInterval string

const (
	RadarHTTPTimeseryBrowserListParamsAggInterval15m RadarHTTPTimeseryBrowserListParamsAggInterval = "15m"
	RadarHTTPTimeseryBrowserListParamsAggInterval1h  RadarHTTPTimeseryBrowserListParamsAggInterval = "1h"
	RadarHTTPTimeseryBrowserListParamsAggInterval1d  RadarHTTPTimeseryBrowserListParamsAggInterval = "1d"
	RadarHTTPTimeseryBrowserListParamsAggInterval1w  RadarHTTPTimeseryBrowserListParamsAggInterval = "1w"
)

type RadarHTTPTimeseryBrowserListParamsDateRange string

const (
	RadarHTTPTimeseryBrowserListParamsDateRange1d         RadarHTTPTimeseryBrowserListParamsDateRange = "1d"
	RadarHTTPTimeseryBrowserListParamsDateRange7d         RadarHTTPTimeseryBrowserListParamsDateRange = "7d"
	RadarHTTPTimeseryBrowserListParamsDateRange14d        RadarHTTPTimeseryBrowserListParamsDateRange = "14d"
	RadarHTTPTimeseryBrowserListParamsDateRange28d        RadarHTTPTimeseryBrowserListParamsDateRange = "28d"
	RadarHTTPTimeseryBrowserListParamsDateRange12w        RadarHTTPTimeseryBrowserListParamsDateRange = "12w"
	RadarHTTPTimeseryBrowserListParamsDateRange24w        RadarHTTPTimeseryBrowserListParamsDateRange = "24w"
	RadarHTTPTimeseryBrowserListParamsDateRange52w        RadarHTTPTimeseryBrowserListParamsDateRange = "52w"
	RadarHTTPTimeseryBrowserListParamsDateRange1dControl  RadarHTTPTimeseryBrowserListParamsDateRange = "1dControl"
	RadarHTTPTimeseryBrowserListParamsDateRange7dControl  RadarHTTPTimeseryBrowserListParamsDateRange = "7dControl"
	RadarHTTPTimeseryBrowserListParamsDateRange14dControl RadarHTTPTimeseryBrowserListParamsDateRange = "14dControl"
	RadarHTTPTimeseryBrowserListParamsDateRange28dControl RadarHTTPTimeseryBrowserListParamsDateRange = "28dControl"
	RadarHTTPTimeseryBrowserListParamsDateRange12wControl RadarHTTPTimeseryBrowserListParamsDateRange = "12wControl"
	RadarHTTPTimeseryBrowserListParamsDateRange24wControl RadarHTTPTimeseryBrowserListParamsDateRange = "24wControl"
)

type RadarHTTPTimeseryBrowserListParamsDeviceType string

const (
	RadarHTTPTimeseryBrowserListParamsDeviceTypeDesktop RadarHTTPTimeseryBrowserListParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseryBrowserListParamsDeviceTypeMobile  RadarHTTPTimeseryBrowserListParamsDeviceType = "MOBILE"
	RadarHTTPTimeseryBrowserListParamsDeviceTypeOther   RadarHTTPTimeseryBrowserListParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTimeseryBrowserListParamsFormat string

const (
	RadarHTTPTimeseryBrowserListParamsFormatJson RadarHTTPTimeseryBrowserListParamsFormat = "JSON"
	RadarHTTPTimeseryBrowserListParamsFormatCsv  RadarHTTPTimeseryBrowserListParamsFormat = "CSV"
)

type RadarHTTPTimeseryBrowserListParamsHTTPProtocol string

const (
	RadarHTTPTimeseryBrowserListParamsHTTPProtocolHTTP  RadarHTTPTimeseryBrowserListParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseryBrowserListParamsHTTPProtocolHTTPs RadarHTTPTimeseryBrowserListParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTimeseryBrowserListParamsHTTPVersion string

const (
	RadarHTTPTimeseryBrowserListParamsHTTPVersionHttPv1 RadarHTTPTimeseryBrowserListParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseryBrowserListParamsHTTPVersionHttPv2 RadarHTTPTimeseryBrowserListParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseryBrowserListParamsHTTPVersionHttPv3 RadarHTTPTimeseryBrowserListParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTimeseryBrowserListParamsIPVersion string

const (
	RadarHTTPTimeseryBrowserListParamsIPVersionIPv4 RadarHTTPTimeseryBrowserListParamsIPVersion = "IPv4"
	RadarHTTPTimeseryBrowserListParamsIPVersionIPv6 RadarHTTPTimeseryBrowserListParamsIPVersion = "IPv6"
)

type RadarHTTPTimeseryBrowserListParamsO string

const (
	RadarHTTPTimeseryBrowserListParamsOWindows  RadarHTTPTimeseryBrowserListParamsO = "WINDOWS"
	RadarHTTPTimeseryBrowserListParamsOMacosx   RadarHTTPTimeseryBrowserListParamsO = "MACOSX"
	RadarHTTPTimeseryBrowserListParamsOAndroid  RadarHTTPTimeseryBrowserListParamsO = "ANDROID"
	RadarHTTPTimeseryBrowserListParamsOChromeos RadarHTTPTimeseryBrowserListParamsO = "CHROMEOS"
	RadarHTTPTimeseryBrowserListParamsOLinux    RadarHTTPTimeseryBrowserListParamsO = "LINUX"
	RadarHTTPTimeseryBrowserListParamsOSmartTv  RadarHTTPTimeseryBrowserListParamsO = "SMART_TV"
)

type RadarHTTPTimeseryBrowserListParamsTlsVersion string

const (
	RadarHTTPTimeseryBrowserListParamsTlsVersionTlSv1_0  RadarHTTPTimeseryBrowserListParamsTlsVersion = "TLSv1_0"
	RadarHTTPTimeseryBrowserListParamsTlsVersionTlSv1_1  RadarHTTPTimeseryBrowserListParamsTlsVersion = "TLSv1_1"
	RadarHTTPTimeseryBrowserListParamsTlsVersionTlSv1_2  RadarHTTPTimeseryBrowserListParamsTlsVersion = "TLSv1_2"
	RadarHTTPTimeseryBrowserListParamsTlsVersionTlSv1_3  RadarHTTPTimeseryBrowserListParamsTlsVersion = "TLSv1_3"
	RadarHTTPTimeseryBrowserListParamsTlsVersionTlSvQuic RadarHTTPTimeseryBrowserListParamsTlsVersion = "TLSvQUIC"
)
