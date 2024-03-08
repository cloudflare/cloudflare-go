// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// RadarHTTPTimeseriesGroupService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarHTTPTimeseriesGroupService] method instead.
type RadarHTTPTimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTimeseriesGroupService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPTimeseriesGroupService(opts ...option.RequestOption) (r *RadarHTTPTimeseriesGroupService) {
	r = &RadarHTTPTimeseriesGroupService{}
	r.Options = opts
	return
}

// Get a time series of the percentage distribution of traffic classified as
// automated or human. Visit
// https://developers.cloudflare.com/radar/concepts/bot-classes/ for more
// information.
func (r *RadarHTTPTimeseriesGroupService) BotClass(ctx context.Context, query RadarHTTPTimeseriesGroupBotClassParams, opts ...option.RequestOption) (res *RadarHTTPTimeseriesGroupBotClassResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarHTTPTimeseriesGroupBotClassResponseEnvelope
	path := "radar/http/timeseries_groups/bot_class"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a time series of the percentage distribution of traffic of the top user
// agents.
func (r *RadarHTTPTimeseriesGroupService) Browser(ctx context.Context, query RadarHTTPTimeseriesGroupBrowserParams, opts ...option.RequestOption) (res *RadarHTTPTimeseriesGroupBrowserResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarHTTPTimeseriesGroupBrowserResponseEnvelope
	path := "radar/http/timeseries_groups/browser"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a time series of the percentage distribution of traffic of the top user
// agents aggregated in families.
func (r *RadarHTTPTimeseriesGroupService) BrowserFamily(ctx context.Context, query RadarHTTPTimeseriesGroupBrowserFamilyParams, opts ...option.RequestOption) (res *RadarHTTPTimeseriesGroupBrowserFamilyResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarHTTPTimeseriesGroupBrowserFamilyResponseEnvelope
	path := "radar/http/timeseries_groups/browser_family"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a time series of the percentage distribution of traffic per device type.
func (r *RadarHTTPTimeseriesGroupService) DeviceType(ctx context.Context, query RadarHTTPTimeseriesGroupDeviceTypeParams, opts ...option.RequestOption) (res *RadarHTTPTimeseriesGroupDeviceTypeResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarHTTPTimeseriesGroupDeviceTypeResponseEnvelope
	path := "radar/http/timeseries_groups/device_type"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a time series of the percentage distribution of traffic per HTTP protocol.
func (r *RadarHTTPTimeseriesGroupService) HTTPProtocol(ctx context.Context, query RadarHTTPTimeseriesGroupHTTPProtocolParams, opts ...option.RequestOption) (res *RadarHTTPTimeseriesGroupHTTPProtocolResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarHTTPTimeseriesGroupHTTPProtocolResponseEnvelope
	path := "radar/http/timeseries_groups/http_protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a time series of the percentage distribution of traffic per HTTP protocol
// version.
func (r *RadarHTTPTimeseriesGroupService) HTTPVersion(ctx context.Context, query RadarHTTPTimeseriesGroupHTTPVersionParams, opts ...option.RequestOption) (res *RadarHTTPTimeseriesGroupHTTPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarHTTPTimeseriesGroupHTTPVersionResponseEnvelope
	path := "radar/http/timeseries_groups/http_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a time series of the percentage distribution of traffic per IP protocol
// version.
func (r *RadarHTTPTimeseriesGroupService) IPVersion(ctx context.Context, query RadarHTTPTimeseriesGroupIPVersionParams, opts ...option.RequestOption) (res *RadarHTTPTimeseriesGroupIPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarHTTPTimeseriesGroupIPVersionResponseEnvelope
	path := "radar/http/timeseries_groups/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a time series of the percentage distribution of traffic of the top operating
// systems.
func (r *RadarHTTPTimeseriesGroupService) OS(ctx context.Context, query RadarHTTPTimeseriesGroupOSParams, opts ...option.RequestOption) (res *RadarHTTPTimeseriesGroupOSResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarHTTPTimeseriesGroupOSResponseEnvelope
	path := "radar/http/timeseries_groups/os"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a time series of the percentage distribution of traffic per TLS protocol
// version.
func (r *RadarHTTPTimeseriesGroupService) TLSVersion(ctx context.Context, query RadarHTTPTimeseriesGroupTLSVersionParams, opts ...option.RequestOption) (res *RadarHTTPTimeseriesGroupTLSVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarHTTPTimeseriesGroupTLSVersionResponseEnvelope
	path := "radar/http/timeseries_groups/tls_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarHTTPTimeseriesGroupBotClassResponse struct {
	Meta   interface{}                                    `json:"meta,required"`
	Serie0 RadarHTTPTimeseriesGroupBotClassResponseSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseriesGroupBotClassResponseJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupBotClassResponseJSON contains the JSON metadata for the
// struct [RadarHTTPTimeseriesGroupBotClassResponse]
type radarHTTPTimeseriesGroupBotClassResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupBotClassResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupBotClassResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupBotClassResponseSerie0 struct {
	Bot        []string                                           `json:"bot,required"`
	Human      []string                                           `json:"human,required"`
	Timestamps []string                                           `json:"timestamps,required"`
	JSON       radarHTTPTimeseriesGroupBotClassResponseSerie0JSON `json:"-"`
}

// radarHTTPTimeseriesGroupBotClassResponseSerie0JSON contains the JSON metadata
// for the struct [RadarHTTPTimeseriesGroupBotClassResponseSerie0]
type radarHTTPTimeseriesGroupBotClassResponseSerie0JSON struct {
	Bot         apijson.Field
	Human       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupBotClassResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupBotClassResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupBrowserResponse struct {
	Meta   interface{}                                   `json:"meta,required"`
	Serie0 RadarHTTPTimeseriesGroupBrowserResponseSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseriesGroupBrowserResponseJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupBrowserResponseJSON contains the JSON metadata for the
// struct [RadarHTTPTimeseriesGroupBrowserResponse]
type radarHTTPTimeseriesGroupBrowserResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupBrowserResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupBrowserResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupBrowserResponseSerie0 struct {
	Timestamps  []string                                          `json:"timestamps,required"`
	ExtraFields map[string][]string                               `json:"-,extras"`
	JSON        radarHTTPTimeseriesGroupBrowserResponseSerie0JSON `json:"-"`
}

// radarHTTPTimeseriesGroupBrowserResponseSerie0JSON contains the JSON metadata for
// the struct [RadarHTTPTimeseriesGroupBrowserResponseSerie0]
type radarHTTPTimeseriesGroupBrowserResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupBrowserResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupBrowserResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupBrowserFamilyResponse struct {
	Meta   interface{}                                         `json:"meta,required"`
	Serie0 RadarHTTPTimeseriesGroupBrowserFamilyResponseSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseriesGroupBrowserFamilyResponseJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupBrowserFamilyResponseJSON contains the JSON metadata for
// the struct [RadarHTTPTimeseriesGroupBrowserFamilyResponse]
type radarHTTPTimeseriesGroupBrowserFamilyResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupBrowserFamilyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupBrowserFamilyResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupBrowserFamilyResponseSerie0 struct {
	Timestamps  []string                                                `json:"timestamps,required"`
	ExtraFields map[string][]string                                     `json:"-,extras"`
	JSON        radarHTTPTimeseriesGroupBrowserFamilyResponseSerie0JSON `json:"-"`
}

// radarHTTPTimeseriesGroupBrowserFamilyResponseSerie0JSON contains the JSON
// metadata for the struct [RadarHTTPTimeseriesGroupBrowserFamilyResponseSerie0]
type radarHTTPTimeseriesGroupBrowserFamilyResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupBrowserFamilyResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupBrowserFamilyResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupDeviceTypeResponse struct {
	Meta   interface{}                                      `json:"meta,required"`
	Serie0 RadarHTTPTimeseriesGroupDeviceTypeResponseSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseriesGroupDeviceTypeResponseJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupDeviceTypeResponseJSON contains the JSON metadata for
// the struct [RadarHTTPTimeseriesGroupDeviceTypeResponse]
type radarHTTPTimeseriesGroupDeviceTypeResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupDeviceTypeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupDeviceTypeResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupDeviceTypeResponseSerie0 struct {
	Desktop    []string                                             `json:"desktop,required"`
	Mobile     []string                                             `json:"mobile,required"`
	Other      []string                                             `json:"other,required"`
	Timestamps []string                                             `json:"timestamps,required"`
	JSON       radarHTTPTimeseriesGroupDeviceTypeResponseSerie0JSON `json:"-"`
}

// radarHTTPTimeseriesGroupDeviceTypeResponseSerie0JSON contains the JSON metadata
// for the struct [RadarHTTPTimeseriesGroupDeviceTypeResponseSerie0]
type radarHTTPTimeseriesGroupDeviceTypeResponseSerie0JSON struct {
	Desktop     apijson.Field
	Mobile      apijson.Field
	Other       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupDeviceTypeResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupDeviceTypeResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupHTTPProtocolResponse struct {
	Meta   interface{}                                        `json:"meta,required"`
	Serie0 RadarHTTPTimeseriesGroupHTTPProtocolResponseSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseriesGroupHTTPProtocolResponseJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupHTTPProtocolResponseJSON contains the JSON metadata for
// the struct [RadarHTTPTimeseriesGroupHTTPProtocolResponse]
type radarHTTPTimeseriesGroupHTTPProtocolResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupHTTPProtocolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupHTTPProtocolResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupHTTPProtocolResponseSerie0 struct {
	HTTP       []string                                               `json:"http,required"`
	HTTPS      []string                                               `json:"https,required"`
	Timestamps []string                                               `json:"timestamps,required"`
	JSON       radarHTTPTimeseriesGroupHTTPProtocolResponseSerie0JSON `json:"-"`
}

// radarHTTPTimeseriesGroupHTTPProtocolResponseSerie0JSON contains the JSON
// metadata for the struct [RadarHTTPTimeseriesGroupHTTPProtocolResponseSerie0]
type radarHTTPTimeseriesGroupHTTPProtocolResponseSerie0JSON struct {
	HTTP        apijson.Field
	HTTPS       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupHTTPProtocolResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupHTTPProtocolResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupHTTPVersionResponse struct {
	Meta   interface{}                                       `json:"meta,required"`
	Serie0 RadarHTTPTimeseriesGroupHTTPVersionResponseSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseriesGroupHTTPVersionResponseJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupHTTPVersionResponseJSON contains the JSON metadata for
// the struct [RadarHTTPTimeseriesGroupHTTPVersionResponse]
type radarHTTPTimeseriesGroupHTTPVersionResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupHTTPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupHTTPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupHTTPVersionResponseSerie0 struct {
	HTTP1X     []string                                              `json:"HTTP/1.x,required"`
	HTTP2      []string                                              `json:"HTTP/2,required"`
	HTTP3      []string                                              `json:"HTTP/3,required"`
	Timestamps []string                                              `json:"timestamps,required"`
	JSON       radarHTTPTimeseriesGroupHTTPVersionResponseSerie0JSON `json:"-"`
}

// radarHTTPTimeseriesGroupHTTPVersionResponseSerie0JSON contains the JSON metadata
// for the struct [RadarHTTPTimeseriesGroupHTTPVersionResponseSerie0]
type radarHTTPTimeseriesGroupHTTPVersionResponseSerie0JSON struct {
	HTTP1X      apijson.Field
	HTTP2       apijson.Field
	HTTP3       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupHTTPVersionResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupHTTPVersionResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupIPVersionResponse struct {
	Meta   interface{}                                     `json:"meta,required"`
	Serie0 RadarHTTPTimeseriesGroupIPVersionResponseSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseriesGroupIPVersionResponseJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupIPVersionResponseJSON contains the JSON metadata for the
// struct [RadarHTTPTimeseriesGroupIPVersionResponse]
type radarHTTPTimeseriesGroupIPVersionResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupIPVersionResponseSerie0 struct {
	IPv4       []string                                            `json:"IPv4,required"`
	IPv6       []string                                            `json:"IPv6,required"`
	Timestamps []string                                            `json:"timestamps,required"`
	JSON       radarHTTPTimeseriesGroupIPVersionResponseSerie0JSON `json:"-"`
}

// radarHTTPTimeseriesGroupIPVersionResponseSerie0JSON contains the JSON metadata
// for the struct [RadarHTTPTimeseriesGroupIPVersionResponseSerie0]
type radarHTTPTimeseriesGroupIPVersionResponseSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupIPVersionResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupIPVersionResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupOSResponse struct {
	Meta   interface{}                              `json:"meta,required"`
	Serie0 RadarHTTPTimeseriesGroupOSResponseSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseriesGroupOSResponseJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupOSResponseJSON contains the JSON metadata for the struct
// [RadarHTTPTimeseriesGroupOSResponse]
type radarHTTPTimeseriesGroupOSResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupOSResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupOSResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupOSResponseSerie0 struct {
	Timestamps  []string                                     `json:"timestamps,required"`
	ExtraFields map[string][]string                          `json:"-,extras"`
	JSON        radarHTTPTimeseriesGroupOSResponseSerie0JSON `json:"-"`
}

// radarHTTPTimeseriesGroupOSResponseSerie0JSON contains the JSON metadata for the
// struct [RadarHTTPTimeseriesGroupOSResponseSerie0]
type radarHTTPTimeseriesGroupOSResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupOSResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupOSResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupTLSVersionResponse struct {
	Meta   interface{}                                      `json:"meta,required"`
	Serie0 RadarHTTPTimeseriesGroupTLSVersionResponseSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseriesGroupTLSVersionResponseJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupTLSVersionResponseJSON contains the JSON metadata for
// the struct [RadarHTTPTimeseriesGroupTLSVersionResponse]
type radarHTTPTimeseriesGroupTLSVersionResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupTLSVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupTLSVersionResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupTLSVersionResponseSerie0 struct {
	Timestamps []string                                             `json:"timestamps,required"`
	TLS1_0     []string                                             `json:"TLS 1.0,required"`
	TLS1_1     []string                                             `json:"TLS 1.1,required"`
	TLS1_2     []string                                             `json:"TLS 1.2,required"`
	TLS1_3     []string                                             `json:"TLS 1.3,required"`
	TLSQuic    []string                                             `json:"TLS QUIC,required"`
	JSON       radarHTTPTimeseriesGroupTLSVersionResponseSerie0JSON `json:"-"`
}

// radarHTTPTimeseriesGroupTLSVersionResponseSerie0JSON contains the JSON metadata
// for the struct [RadarHTTPTimeseriesGroupTLSVersionResponseSerie0]
type radarHTTPTimeseriesGroupTLSVersionResponseSerie0JSON struct {
	Timestamps  apijson.Field
	TLS1_0      apijson.Field
	TLS1_1      apijson.Field
	TLS1_2      apijson.Field
	TLS1_3      apijson.Field
	TLSQuic     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupTLSVersionResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupTLSVersionResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupBotClassParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseriesGroupBotClassParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTimeseriesGroupBotClassParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTimeseriesGroupBotClassParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTimeseriesGroupBotClassParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTimeseriesGroupBotClassParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTimeseriesGroupBotClassParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTimeseriesGroupBotClassParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]RadarHTTPTimeseriesGroupBotClassParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHTTPTimeseriesGroupBotClassParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseriesGroupBotClassParams]'s query parameters
// as `url.Values`.
func (r RadarHTTPTimeseriesGroupBotClassParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseriesGroupBotClassParamsAggInterval string

const (
	RadarHTTPTimeseriesGroupBotClassParamsAggInterval15m RadarHTTPTimeseriesGroupBotClassParamsAggInterval = "15m"
	RadarHTTPTimeseriesGroupBotClassParamsAggInterval1h  RadarHTTPTimeseriesGroupBotClassParamsAggInterval = "1h"
	RadarHTTPTimeseriesGroupBotClassParamsAggInterval1d  RadarHTTPTimeseriesGroupBotClassParamsAggInterval = "1d"
	RadarHTTPTimeseriesGroupBotClassParamsAggInterval1w  RadarHTTPTimeseriesGroupBotClassParamsAggInterval = "1w"
)

type RadarHTTPTimeseriesGroupBotClassParamsDateRange string

const (
	RadarHTTPTimeseriesGroupBotClassParamsDateRange1d         RadarHTTPTimeseriesGroupBotClassParamsDateRange = "1d"
	RadarHTTPTimeseriesGroupBotClassParamsDateRange2d         RadarHTTPTimeseriesGroupBotClassParamsDateRange = "2d"
	RadarHTTPTimeseriesGroupBotClassParamsDateRange7d         RadarHTTPTimeseriesGroupBotClassParamsDateRange = "7d"
	RadarHTTPTimeseriesGroupBotClassParamsDateRange14d        RadarHTTPTimeseriesGroupBotClassParamsDateRange = "14d"
	RadarHTTPTimeseriesGroupBotClassParamsDateRange28d        RadarHTTPTimeseriesGroupBotClassParamsDateRange = "28d"
	RadarHTTPTimeseriesGroupBotClassParamsDateRange12w        RadarHTTPTimeseriesGroupBotClassParamsDateRange = "12w"
	RadarHTTPTimeseriesGroupBotClassParamsDateRange24w        RadarHTTPTimeseriesGroupBotClassParamsDateRange = "24w"
	RadarHTTPTimeseriesGroupBotClassParamsDateRange52w        RadarHTTPTimeseriesGroupBotClassParamsDateRange = "52w"
	RadarHTTPTimeseriesGroupBotClassParamsDateRange1dControl  RadarHTTPTimeseriesGroupBotClassParamsDateRange = "1dControl"
	RadarHTTPTimeseriesGroupBotClassParamsDateRange2dControl  RadarHTTPTimeseriesGroupBotClassParamsDateRange = "2dControl"
	RadarHTTPTimeseriesGroupBotClassParamsDateRange7dControl  RadarHTTPTimeseriesGroupBotClassParamsDateRange = "7dControl"
	RadarHTTPTimeseriesGroupBotClassParamsDateRange14dControl RadarHTTPTimeseriesGroupBotClassParamsDateRange = "14dControl"
	RadarHTTPTimeseriesGroupBotClassParamsDateRange28dControl RadarHTTPTimeseriesGroupBotClassParamsDateRange = "28dControl"
	RadarHTTPTimeseriesGroupBotClassParamsDateRange12wControl RadarHTTPTimeseriesGroupBotClassParamsDateRange = "12wControl"
	RadarHTTPTimeseriesGroupBotClassParamsDateRange24wControl RadarHTTPTimeseriesGroupBotClassParamsDateRange = "24wControl"
)

type RadarHTTPTimeseriesGroupBotClassParamsDeviceType string

const (
	RadarHTTPTimeseriesGroupBotClassParamsDeviceTypeDesktop RadarHTTPTimeseriesGroupBotClassParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseriesGroupBotClassParamsDeviceTypeMobile  RadarHTTPTimeseriesGroupBotClassParamsDeviceType = "MOBILE"
	RadarHTTPTimeseriesGroupBotClassParamsDeviceTypeOther   RadarHTTPTimeseriesGroupBotClassParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTimeseriesGroupBotClassParamsFormat string

const (
	RadarHTTPTimeseriesGroupBotClassParamsFormatJson RadarHTTPTimeseriesGroupBotClassParamsFormat = "JSON"
	RadarHTTPTimeseriesGroupBotClassParamsFormatCsv  RadarHTTPTimeseriesGroupBotClassParamsFormat = "CSV"
)

type RadarHTTPTimeseriesGroupBotClassParamsHTTPProtocol string

const (
	RadarHTTPTimeseriesGroupBotClassParamsHTTPProtocolHTTP  RadarHTTPTimeseriesGroupBotClassParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseriesGroupBotClassParamsHTTPProtocolHTTPS RadarHTTPTimeseriesGroupBotClassParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTimeseriesGroupBotClassParamsHTTPVersion string

const (
	RadarHTTPTimeseriesGroupBotClassParamsHTTPVersionHttPv1 RadarHTTPTimeseriesGroupBotClassParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseriesGroupBotClassParamsHTTPVersionHttPv2 RadarHTTPTimeseriesGroupBotClassParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseriesGroupBotClassParamsHTTPVersionHttPv3 RadarHTTPTimeseriesGroupBotClassParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTimeseriesGroupBotClassParamsIPVersion string

const (
	RadarHTTPTimeseriesGroupBotClassParamsIPVersionIPv4 RadarHTTPTimeseriesGroupBotClassParamsIPVersion = "IPv4"
	RadarHTTPTimeseriesGroupBotClassParamsIPVersionIPv6 RadarHTTPTimeseriesGroupBotClassParamsIPVersion = "IPv6"
)

type RadarHTTPTimeseriesGroupBotClassParamsOS string

const (
	RadarHTTPTimeseriesGroupBotClassParamsOSWindows  RadarHTTPTimeseriesGroupBotClassParamsOS = "WINDOWS"
	RadarHTTPTimeseriesGroupBotClassParamsOSMacosx   RadarHTTPTimeseriesGroupBotClassParamsOS = "MACOSX"
	RadarHTTPTimeseriesGroupBotClassParamsOSIos      RadarHTTPTimeseriesGroupBotClassParamsOS = "IOS"
	RadarHTTPTimeseriesGroupBotClassParamsOSAndroid  RadarHTTPTimeseriesGroupBotClassParamsOS = "ANDROID"
	RadarHTTPTimeseriesGroupBotClassParamsOSChromeos RadarHTTPTimeseriesGroupBotClassParamsOS = "CHROMEOS"
	RadarHTTPTimeseriesGroupBotClassParamsOSLinux    RadarHTTPTimeseriesGroupBotClassParamsOS = "LINUX"
	RadarHTTPTimeseriesGroupBotClassParamsOSSmartTv  RadarHTTPTimeseriesGroupBotClassParamsOS = "SMART_TV"
)

type RadarHTTPTimeseriesGroupBotClassParamsTLSVersion string

const (
	RadarHTTPTimeseriesGroupBotClassParamsTLSVersionTlSv1_0  RadarHTTPTimeseriesGroupBotClassParamsTLSVersion = "TLSv1_0"
	RadarHTTPTimeseriesGroupBotClassParamsTLSVersionTlSv1_1  RadarHTTPTimeseriesGroupBotClassParamsTLSVersion = "TLSv1_1"
	RadarHTTPTimeseriesGroupBotClassParamsTLSVersionTlSv1_2  RadarHTTPTimeseriesGroupBotClassParamsTLSVersion = "TLSv1_2"
	RadarHTTPTimeseriesGroupBotClassParamsTLSVersionTlSv1_3  RadarHTTPTimeseriesGroupBotClassParamsTLSVersion = "TLSv1_3"
	RadarHTTPTimeseriesGroupBotClassParamsTLSVersionTlSvQuic RadarHTTPTimeseriesGroupBotClassParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPTimeseriesGroupBotClassResponseEnvelope struct {
	Result  RadarHTTPTimeseriesGroupBotClassResponse             `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarHTTPTimeseriesGroupBotClassResponseEnvelopeJSON `json:"-"`
}

// radarHTTPTimeseriesGroupBotClassResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarHTTPTimeseriesGroupBotClassResponseEnvelope]
type radarHTTPTimeseriesGroupBotClassResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupBotClassResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupBotClassResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupBrowserParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseriesGroupBrowserParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTimeseriesGroupBrowserParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTimeseriesGroupBrowserParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTimeseriesGroupBrowserParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTimeseriesGroupBrowserParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTimeseriesGroupBrowserParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTimeseriesGroupBrowserParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTimeseriesGroupBrowserParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects (eg browsers, verticals, etc) to the top items over
	// the time range.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]RadarHTTPTimeseriesGroupBrowserParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHTTPTimeseriesGroupBrowserParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseriesGroupBrowserParams]'s query parameters
// as `url.Values`.
func (r RadarHTTPTimeseriesGroupBrowserParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseriesGroupBrowserParamsAggInterval string

const (
	RadarHTTPTimeseriesGroupBrowserParamsAggInterval15m RadarHTTPTimeseriesGroupBrowserParamsAggInterval = "15m"
	RadarHTTPTimeseriesGroupBrowserParamsAggInterval1h  RadarHTTPTimeseriesGroupBrowserParamsAggInterval = "1h"
	RadarHTTPTimeseriesGroupBrowserParamsAggInterval1d  RadarHTTPTimeseriesGroupBrowserParamsAggInterval = "1d"
	RadarHTTPTimeseriesGroupBrowserParamsAggInterval1w  RadarHTTPTimeseriesGroupBrowserParamsAggInterval = "1w"
)

type RadarHTTPTimeseriesGroupBrowserParamsBotClass string

const (
	RadarHTTPTimeseriesGroupBrowserParamsBotClassLikelyAutomated RadarHTTPTimeseriesGroupBrowserParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTimeseriesGroupBrowserParamsBotClassLikelyHuman     RadarHTTPTimeseriesGroupBrowserParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTimeseriesGroupBrowserParamsDateRange string

const (
	RadarHTTPTimeseriesGroupBrowserParamsDateRange1d         RadarHTTPTimeseriesGroupBrowserParamsDateRange = "1d"
	RadarHTTPTimeseriesGroupBrowserParamsDateRange2d         RadarHTTPTimeseriesGroupBrowserParamsDateRange = "2d"
	RadarHTTPTimeseriesGroupBrowserParamsDateRange7d         RadarHTTPTimeseriesGroupBrowserParamsDateRange = "7d"
	RadarHTTPTimeseriesGroupBrowserParamsDateRange14d        RadarHTTPTimeseriesGroupBrowserParamsDateRange = "14d"
	RadarHTTPTimeseriesGroupBrowserParamsDateRange28d        RadarHTTPTimeseriesGroupBrowserParamsDateRange = "28d"
	RadarHTTPTimeseriesGroupBrowserParamsDateRange12w        RadarHTTPTimeseriesGroupBrowserParamsDateRange = "12w"
	RadarHTTPTimeseriesGroupBrowserParamsDateRange24w        RadarHTTPTimeseriesGroupBrowserParamsDateRange = "24w"
	RadarHTTPTimeseriesGroupBrowserParamsDateRange52w        RadarHTTPTimeseriesGroupBrowserParamsDateRange = "52w"
	RadarHTTPTimeseriesGroupBrowserParamsDateRange1dControl  RadarHTTPTimeseriesGroupBrowserParamsDateRange = "1dControl"
	RadarHTTPTimeseriesGroupBrowserParamsDateRange2dControl  RadarHTTPTimeseriesGroupBrowserParamsDateRange = "2dControl"
	RadarHTTPTimeseriesGroupBrowserParamsDateRange7dControl  RadarHTTPTimeseriesGroupBrowserParamsDateRange = "7dControl"
	RadarHTTPTimeseriesGroupBrowserParamsDateRange14dControl RadarHTTPTimeseriesGroupBrowserParamsDateRange = "14dControl"
	RadarHTTPTimeseriesGroupBrowserParamsDateRange28dControl RadarHTTPTimeseriesGroupBrowserParamsDateRange = "28dControl"
	RadarHTTPTimeseriesGroupBrowserParamsDateRange12wControl RadarHTTPTimeseriesGroupBrowserParamsDateRange = "12wControl"
	RadarHTTPTimeseriesGroupBrowserParamsDateRange24wControl RadarHTTPTimeseriesGroupBrowserParamsDateRange = "24wControl"
)

type RadarHTTPTimeseriesGroupBrowserParamsDeviceType string

const (
	RadarHTTPTimeseriesGroupBrowserParamsDeviceTypeDesktop RadarHTTPTimeseriesGroupBrowserParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseriesGroupBrowserParamsDeviceTypeMobile  RadarHTTPTimeseriesGroupBrowserParamsDeviceType = "MOBILE"
	RadarHTTPTimeseriesGroupBrowserParamsDeviceTypeOther   RadarHTTPTimeseriesGroupBrowserParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTimeseriesGroupBrowserParamsFormat string

const (
	RadarHTTPTimeseriesGroupBrowserParamsFormatJson RadarHTTPTimeseriesGroupBrowserParamsFormat = "JSON"
	RadarHTTPTimeseriesGroupBrowserParamsFormatCsv  RadarHTTPTimeseriesGroupBrowserParamsFormat = "CSV"
)

type RadarHTTPTimeseriesGroupBrowserParamsHTTPProtocol string

const (
	RadarHTTPTimeseriesGroupBrowserParamsHTTPProtocolHTTP  RadarHTTPTimeseriesGroupBrowserParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseriesGroupBrowserParamsHTTPProtocolHTTPS RadarHTTPTimeseriesGroupBrowserParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTimeseriesGroupBrowserParamsHTTPVersion string

const (
	RadarHTTPTimeseriesGroupBrowserParamsHTTPVersionHttPv1 RadarHTTPTimeseriesGroupBrowserParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseriesGroupBrowserParamsHTTPVersionHttPv2 RadarHTTPTimeseriesGroupBrowserParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseriesGroupBrowserParamsHTTPVersionHttPv3 RadarHTTPTimeseriesGroupBrowserParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTimeseriesGroupBrowserParamsIPVersion string

const (
	RadarHTTPTimeseriesGroupBrowserParamsIPVersionIPv4 RadarHTTPTimeseriesGroupBrowserParamsIPVersion = "IPv4"
	RadarHTTPTimeseriesGroupBrowserParamsIPVersionIPv6 RadarHTTPTimeseriesGroupBrowserParamsIPVersion = "IPv6"
)

type RadarHTTPTimeseriesGroupBrowserParamsOS string

const (
	RadarHTTPTimeseriesGroupBrowserParamsOSWindows  RadarHTTPTimeseriesGroupBrowserParamsOS = "WINDOWS"
	RadarHTTPTimeseriesGroupBrowserParamsOSMacosx   RadarHTTPTimeseriesGroupBrowserParamsOS = "MACOSX"
	RadarHTTPTimeseriesGroupBrowserParamsOSIos      RadarHTTPTimeseriesGroupBrowserParamsOS = "IOS"
	RadarHTTPTimeseriesGroupBrowserParamsOSAndroid  RadarHTTPTimeseriesGroupBrowserParamsOS = "ANDROID"
	RadarHTTPTimeseriesGroupBrowserParamsOSChromeos RadarHTTPTimeseriesGroupBrowserParamsOS = "CHROMEOS"
	RadarHTTPTimeseriesGroupBrowserParamsOSLinux    RadarHTTPTimeseriesGroupBrowserParamsOS = "LINUX"
	RadarHTTPTimeseriesGroupBrowserParamsOSSmartTv  RadarHTTPTimeseriesGroupBrowserParamsOS = "SMART_TV"
)

type RadarHTTPTimeseriesGroupBrowserParamsTLSVersion string

const (
	RadarHTTPTimeseriesGroupBrowserParamsTLSVersionTlSv1_0  RadarHTTPTimeseriesGroupBrowserParamsTLSVersion = "TLSv1_0"
	RadarHTTPTimeseriesGroupBrowserParamsTLSVersionTlSv1_1  RadarHTTPTimeseriesGroupBrowserParamsTLSVersion = "TLSv1_1"
	RadarHTTPTimeseriesGroupBrowserParamsTLSVersionTlSv1_2  RadarHTTPTimeseriesGroupBrowserParamsTLSVersion = "TLSv1_2"
	RadarHTTPTimeseriesGroupBrowserParamsTLSVersionTlSv1_3  RadarHTTPTimeseriesGroupBrowserParamsTLSVersion = "TLSv1_3"
	RadarHTTPTimeseriesGroupBrowserParamsTLSVersionTlSvQuic RadarHTTPTimeseriesGroupBrowserParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPTimeseriesGroupBrowserResponseEnvelope struct {
	Result  RadarHTTPTimeseriesGroupBrowserResponse             `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    radarHTTPTimeseriesGroupBrowserResponseEnvelopeJSON `json:"-"`
}

// radarHTTPTimeseriesGroupBrowserResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarHTTPTimeseriesGroupBrowserResponseEnvelope]
type radarHTTPTimeseriesGroupBrowserResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupBrowserResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupBrowserResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupBrowserFamilyParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseriesGroupBrowserFamilyParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTimeseriesGroupBrowserFamilyParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTimeseriesGroupBrowserFamilyParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTimeseriesGroupBrowserFamilyParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTimeseriesGroupBrowserFamilyParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTimeseriesGroupBrowserFamilyParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]RadarHTTPTimeseriesGroupBrowserFamilyParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHTTPTimeseriesGroupBrowserFamilyParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseriesGroupBrowserFamilyParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTimeseriesGroupBrowserFamilyParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseriesGroupBrowserFamilyParamsAggInterval string

const (
	RadarHTTPTimeseriesGroupBrowserFamilyParamsAggInterval15m RadarHTTPTimeseriesGroupBrowserFamilyParamsAggInterval = "15m"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsAggInterval1h  RadarHTTPTimeseriesGroupBrowserFamilyParamsAggInterval = "1h"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsAggInterval1d  RadarHTTPTimeseriesGroupBrowserFamilyParamsAggInterval = "1d"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsAggInterval1w  RadarHTTPTimeseriesGroupBrowserFamilyParamsAggInterval = "1w"
)

type RadarHTTPTimeseriesGroupBrowserFamilyParamsBotClass string

const (
	RadarHTTPTimeseriesGroupBrowserFamilyParamsBotClassLikelyAutomated RadarHTTPTimeseriesGroupBrowserFamilyParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsBotClassLikelyHuman     RadarHTTPTimeseriesGroupBrowserFamilyParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange string

const (
	RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange1d         RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange = "1d"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange2d         RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange = "2d"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange7d         RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange = "7d"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange14d        RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange = "14d"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange28d        RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange = "28d"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange12w        RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange = "12w"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange24w        RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange = "24w"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange52w        RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange = "52w"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange1dControl  RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange = "1dControl"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange2dControl  RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange = "2dControl"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange7dControl  RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange = "7dControl"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange14dControl RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange = "14dControl"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange28dControl RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange = "28dControl"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange12wControl RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange = "12wControl"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange24wControl RadarHTTPTimeseriesGroupBrowserFamilyParamsDateRange = "24wControl"
)

type RadarHTTPTimeseriesGroupBrowserFamilyParamsDeviceType string

const (
	RadarHTTPTimeseriesGroupBrowserFamilyParamsDeviceTypeDesktop RadarHTTPTimeseriesGroupBrowserFamilyParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsDeviceTypeMobile  RadarHTTPTimeseriesGroupBrowserFamilyParamsDeviceType = "MOBILE"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsDeviceTypeOther   RadarHTTPTimeseriesGroupBrowserFamilyParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTimeseriesGroupBrowserFamilyParamsFormat string

const (
	RadarHTTPTimeseriesGroupBrowserFamilyParamsFormatJson RadarHTTPTimeseriesGroupBrowserFamilyParamsFormat = "JSON"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsFormatCsv  RadarHTTPTimeseriesGroupBrowserFamilyParamsFormat = "CSV"
)

type RadarHTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocol string

const (
	RadarHTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocolHTTP  RadarHTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocolHTTPS RadarHTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTimeseriesGroupBrowserFamilyParamsHTTPVersion string

const (
	RadarHTTPTimeseriesGroupBrowserFamilyParamsHTTPVersionHttPv1 RadarHTTPTimeseriesGroupBrowserFamilyParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsHTTPVersionHttPv2 RadarHTTPTimeseriesGroupBrowserFamilyParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsHTTPVersionHttPv3 RadarHTTPTimeseriesGroupBrowserFamilyParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTimeseriesGroupBrowserFamilyParamsIPVersion string

const (
	RadarHTTPTimeseriesGroupBrowserFamilyParamsIPVersionIPv4 RadarHTTPTimeseriesGroupBrowserFamilyParamsIPVersion = "IPv4"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsIPVersionIPv6 RadarHTTPTimeseriesGroupBrowserFamilyParamsIPVersion = "IPv6"
)

type RadarHTTPTimeseriesGroupBrowserFamilyParamsOS string

const (
	RadarHTTPTimeseriesGroupBrowserFamilyParamsOSWindows  RadarHTTPTimeseriesGroupBrowserFamilyParamsOS = "WINDOWS"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsOSMacosx   RadarHTTPTimeseriesGroupBrowserFamilyParamsOS = "MACOSX"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsOSIos      RadarHTTPTimeseriesGroupBrowserFamilyParamsOS = "IOS"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsOSAndroid  RadarHTTPTimeseriesGroupBrowserFamilyParamsOS = "ANDROID"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsOSChromeos RadarHTTPTimeseriesGroupBrowserFamilyParamsOS = "CHROMEOS"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsOSLinux    RadarHTTPTimeseriesGroupBrowserFamilyParamsOS = "LINUX"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsOSSmartTv  RadarHTTPTimeseriesGroupBrowserFamilyParamsOS = "SMART_TV"
)

type RadarHTTPTimeseriesGroupBrowserFamilyParamsTLSVersion string

const (
	RadarHTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSv1_0  RadarHTTPTimeseriesGroupBrowserFamilyParamsTLSVersion = "TLSv1_0"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSv1_1  RadarHTTPTimeseriesGroupBrowserFamilyParamsTLSVersion = "TLSv1_1"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSv1_2  RadarHTTPTimeseriesGroupBrowserFamilyParamsTLSVersion = "TLSv1_2"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSv1_3  RadarHTTPTimeseriesGroupBrowserFamilyParamsTLSVersion = "TLSv1_3"
	RadarHTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSvQuic RadarHTTPTimeseriesGroupBrowserFamilyParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPTimeseriesGroupBrowserFamilyResponseEnvelope struct {
	Result  RadarHTTPTimeseriesGroupBrowserFamilyResponse             `json:"result,required"`
	Success bool                                                      `json:"success,required"`
	JSON    radarHTTPTimeseriesGroupBrowserFamilyResponseEnvelopeJSON `json:"-"`
}

// radarHTTPTimeseriesGroupBrowserFamilyResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarHTTPTimeseriesGroupBrowserFamilyResponseEnvelope]
type radarHTTPTimeseriesGroupBrowserFamilyResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupBrowserFamilyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupBrowserFamilyResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupDeviceTypeParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseriesGroupDeviceTypeParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTimeseriesGroupDeviceTypeParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTimeseriesGroupDeviceTypeParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTimeseriesGroupDeviceTypeParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTimeseriesGroupDeviceTypeParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTimeseriesGroupDeviceTypeParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]RadarHTTPTimeseriesGroupDeviceTypeParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHTTPTimeseriesGroupDeviceTypeParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseriesGroupDeviceTypeParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTimeseriesGroupDeviceTypeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseriesGroupDeviceTypeParamsAggInterval string

const (
	RadarHTTPTimeseriesGroupDeviceTypeParamsAggInterval15m RadarHTTPTimeseriesGroupDeviceTypeParamsAggInterval = "15m"
	RadarHTTPTimeseriesGroupDeviceTypeParamsAggInterval1h  RadarHTTPTimeseriesGroupDeviceTypeParamsAggInterval = "1h"
	RadarHTTPTimeseriesGroupDeviceTypeParamsAggInterval1d  RadarHTTPTimeseriesGroupDeviceTypeParamsAggInterval = "1d"
	RadarHTTPTimeseriesGroupDeviceTypeParamsAggInterval1w  RadarHTTPTimeseriesGroupDeviceTypeParamsAggInterval = "1w"
)

type RadarHTTPTimeseriesGroupDeviceTypeParamsBotClass string

const (
	RadarHTTPTimeseriesGroupDeviceTypeParamsBotClassLikelyAutomated RadarHTTPTimeseriesGroupDeviceTypeParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTimeseriesGroupDeviceTypeParamsBotClassLikelyHuman     RadarHTTPTimeseriesGroupDeviceTypeParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange string

const (
	RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange1d         RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange = "1d"
	RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange2d         RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange = "2d"
	RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange7d         RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange = "7d"
	RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange14d        RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange = "14d"
	RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange28d        RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange = "28d"
	RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange12w        RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange = "12w"
	RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange24w        RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange = "24w"
	RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange52w        RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange = "52w"
	RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange1dControl  RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange = "1dControl"
	RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange2dControl  RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange = "2dControl"
	RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange7dControl  RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange = "7dControl"
	RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange14dControl RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange = "14dControl"
	RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange28dControl RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange = "28dControl"
	RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange12wControl RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange = "12wControl"
	RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange24wControl RadarHTTPTimeseriesGroupDeviceTypeParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarHTTPTimeseriesGroupDeviceTypeParamsFormat string

const (
	RadarHTTPTimeseriesGroupDeviceTypeParamsFormatJson RadarHTTPTimeseriesGroupDeviceTypeParamsFormat = "JSON"
	RadarHTTPTimeseriesGroupDeviceTypeParamsFormatCsv  RadarHTTPTimeseriesGroupDeviceTypeParamsFormat = "CSV"
)

type RadarHTTPTimeseriesGroupDeviceTypeParamsHTTPProtocol string

const (
	RadarHTTPTimeseriesGroupDeviceTypeParamsHTTPProtocolHTTP  RadarHTTPTimeseriesGroupDeviceTypeParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseriesGroupDeviceTypeParamsHTTPProtocolHTTPS RadarHTTPTimeseriesGroupDeviceTypeParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTimeseriesGroupDeviceTypeParamsHTTPVersion string

const (
	RadarHTTPTimeseriesGroupDeviceTypeParamsHTTPVersionHttPv1 RadarHTTPTimeseriesGroupDeviceTypeParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseriesGroupDeviceTypeParamsHTTPVersionHttPv2 RadarHTTPTimeseriesGroupDeviceTypeParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseriesGroupDeviceTypeParamsHTTPVersionHttPv3 RadarHTTPTimeseriesGroupDeviceTypeParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTimeseriesGroupDeviceTypeParamsIPVersion string

const (
	RadarHTTPTimeseriesGroupDeviceTypeParamsIPVersionIPv4 RadarHTTPTimeseriesGroupDeviceTypeParamsIPVersion = "IPv4"
	RadarHTTPTimeseriesGroupDeviceTypeParamsIPVersionIPv6 RadarHTTPTimeseriesGroupDeviceTypeParamsIPVersion = "IPv6"
)

type RadarHTTPTimeseriesGroupDeviceTypeParamsOS string

const (
	RadarHTTPTimeseriesGroupDeviceTypeParamsOSWindows  RadarHTTPTimeseriesGroupDeviceTypeParamsOS = "WINDOWS"
	RadarHTTPTimeseriesGroupDeviceTypeParamsOSMacosx   RadarHTTPTimeseriesGroupDeviceTypeParamsOS = "MACOSX"
	RadarHTTPTimeseriesGroupDeviceTypeParamsOSIos      RadarHTTPTimeseriesGroupDeviceTypeParamsOS = "IOS"
	RadarHTTPTimeseriesGroupDeviceTypeParamsOSAndroid  RadarHTTPTimeseriesGroupDeviceTypeParamsOS = "ANDROID"
	RadarHTTPTimeseriesGroupDeviceTypeParamsOSChromeos RadarHTTPTimeseriesGroupDeviceTypeParamsOS = "CHROMEOS"
	RadarHTTPTimeseriesGroupDeviceTypeParamsOSLinux    RadarHTTPTimeseriesGroupDeviceTypeParamsOS = "LINUX"
	RadarHTTPTimeseriesGroupDeviceTypeParamsOSSmartTv  RadarHTTPTimeseriesGroupDeviceTypeParamsOS = "SMART_TV"
)

type RadarHTTPTimeseriesGroupDeviceTypeParamsTLSVersion string

const (
	RadarHTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSv1_0  RadarHTTPTimeseriesGroupDeviceTypeParamsTLSVersion = "TLSv1_0"
	RadarHTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSv1_1  RadarHTTPTimeseriesGroupDeviceTypeParamsTLSVersion = "TLSv1_1"
	RadarHTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSv1_2  RadarHTTPTimeseriesGroupDeviceTypeParamsTLSVersion = "TLSv1_2"
	RadarHTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSv1_3  RadarHTTPTimeseriesGroupDeviceTypeParamsTLSVersion = "TLSv1_3"
	RadarHTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSvQuic RadarHTTPTimeseriesGroupDeviceTypeParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPTimeseriesGroupDeviceTypeResponseEnvelope struct {
	Result  RadarHTTPTimeseriesGroupDeviceTypeResponse             `json:"result,required"`
	Success bool                                                   `json:"success,required"`
	JSON    radarHTTPTimeseriesGroupDeviceTypeResponseEnvelopeJSON `json:"-"`
}

// radarHTTPTimeseriesGroupDeviceTypeResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarHTTPTimeseriesGroupDeviceTypeResponseEnvelope]
type radarHTTPTimeseriesGroupDeviceTypeResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupDeviceTypeResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupDeviceTypeResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupHTTPProtocolParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseriesGroupHTTPProtocolParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTimeseriesGroupHTTPProtocolParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTimeseriesGroupHTTPProtocolParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTimeseriesGroupHTTPProtocolParamsFormat] `query:"format"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTimeseriesGroupHTTPProtocolParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTimeseriesGroupHTTPProtocolParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]RadarHTTPTimeseriesGroupHTTPProtocolParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHTTPTimeseriesGroupHTTPProtocolParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseriesGroupHTTPProtocolParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTimeseriesGroupHTTPProtocolParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseriesGroupHTTPProtocolParamsAggInterval string

const (
	RadarHTTPTimeseriesGroupHTTPProtocolParamsAggInterval15m RadarHTTPTimeseriesGroupHTTPProtocolParamsAggInterval = "15m"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsAggInterval1h  RadarHTTPTimeseriesGroupHTTPProtocolParamsAggInterval = "1h"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsAggInterval1d  RadarHTTPTimeseriesGroupHTTPProtocolParamsAggInterval = "1d"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsAggInterval1w  RadarHTTPTimeseriesGroupHTTPProtocolParamsAggInterval = "1w"
)

type RadarHTTPTimeseriesGroupHTTPProtocolParamsBotClass string

const (
	RadarHTTPTimeseriesGroupHTTPProtocolParamsBotClassLikelyAutomated RadarHTTPTimeseriesGroupHTTPProtocolParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsBotClassLikelyHuman     RadarHTTPTimeseriesGroupHTTPProtocolParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange string

const (
	RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange1d         RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange = "1d"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange2d         RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange = "2d"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange7d         RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange = "7d"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange14d        RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange = "14d"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange28d        RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange = "28d"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange12w        RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange = "12w"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange24w        RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange = "24w"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange52w        RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange = "52w"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange1dControl  RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange = "1dControl"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange2dControl  RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange = "2dControl"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange7dControl  RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange = "7dControl"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange14dControl RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange = "14dControl"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange28dControl RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange = "28dControl"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange12wControl RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange = "12wControl"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange24wControl RadarHTTPTimeseriesGroupHTTPProtocolParamsDateRange = "24wControl"
)

type RadarHTTPTimeseriesGroupHTTPProtocolParamsDeviceType string

const (
	RadarHTTPTimeseriesGroupHTTPProtocolParamsDeviceTypeDesktop RadarHTTPTimeseriesGroupHTTPProtocolParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsDeviceTypeMobile  RadarHTTPTimeseriesGroupHTTPProtocolParamsDeviceType = "MOBILE"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsDeviceTypeOther   RadarHTTPTimeseriesGroupHTTPProtocolParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTimeseriesGroupHTTPProtocolParamsFormat string

const (
	RadarHTTPTimeseriesGroupHTTPProtocolParamsFormatJson RadarHTTPTimeseriesGroupHTTPProtocolParamsFormat = "JSON"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsFormatCsv  RadarHTTPTimeseriesGroupHTTPProtocolParamsFormat = "CSV"
)

type RadarHTTPTimeseriesGroupHTTPProtocolParamsHTTPVersion string

const (
	RadarHTTPTimeseriesGroupHTTPProtocolParamsHTTPVersionHttPv1 RadarHTTPTimeseriesGroupHTTPProtocolParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsHTTPVersionHttPv2 RadarHTTPTimeseriesGroupHTTPProtocolParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsHTTPVersionHttPv3 RadarHTTPTimeseriesGroupHTTPProtocolParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTimeseriesGroupHTTPProtocolParamsIPVersion string

const (
	RadarHTTPTimeseriesGroupHTTPProtocolParamsIPVersionIPv4 RadarHTTPTimeseriesGroupHTTPProtocolParamsIPVersion = "IPv4"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsIPVersionIPv6 RadarHTTPTimeseriesGroupHTTPProtocolParamsIPVersion = "IPv6"
)

type RadarHTTPTimeseriesGroupHTTPProtocolParamsOS string

const (
	RadarHTTPTimeseriesGroupHTTPProtocolParamsOSWindows  RadarHTTPTimeseriesGroupHTTPProtocolParamsOS = "WINDOWS"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsOSMacosx   RadarHTTPTimeseriesGroupHTTPProtocolParamsOS = "MACOSX"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsOSIos      RadarHTTPTimeseriesGroupHTTPProtocolParamsOS = "IOS"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsOSAndroid  RadarHTTPTimeseriesGroupHTTPProtocolParamsOS = "ANDROID"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsOSChromeos RadarHTTPTimeseriesGroupHTTPProtocolParamsOS = "CHROMEOS"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsOSLinux    RadarHTTPTimeseriesGroupHTTPProtocolParamsOS = "LINUX"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsOSSmartTv  RadarHTTPTimeseriesGroupHTTPProtocolParamsOS = "SMART_TV"
)

type RadarHTTPTimeseriesGroupHTTPProtocolParamsTLSVersion string

const (
	RadarHTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSv1_0  RadarHTTPTimeseriesGroupHTTPProtocolParamsTLSVersion = "TLSv1_0"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSv1_1  RadarHTTPTimeseriesGroupHTTPProtocolParamsTLSVersion = "TLSv1_1"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSv1_2  RadarHTTPTimeseriesGroupHTTPProtocolParamsTLSVersion = "TLSv1_2"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSv1_3  RadarHTTPTimeseriesGroupHTTPProtocolParamsTLSVersion = "TLSv1_3"
	RadarHTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSvQuic RadarHTTPTimeseriesGroupHTTPProtocolParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPTimeseriesGroupHTTPProtocolResponseEnvelope struct {
	Result  RadarHTTPTimeseriesGroupHTTPProtocolResponse             `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    radarHTTPTimeseriesGroupHTTPProtocolResponseEnvelopeJSON `json:"-"`
}

// radarHTTPTimeseriesGroupHTTPProtocolResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarHTTPTimeseriesGroupHTTPProtocolResponseEnvelope]
type radarHTTPTimeseriesGroupHTTPProtocolResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupHTTPProtocolResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupHTTPProtocolResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupHTTPVersionParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseriesGroupHTTPVersionParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTimeseriesGroupHTTPVersionParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTimeseriesGroupHTTPVersionParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTimeseriesGroupHTTPVersionParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTimeseriesGroupHTTPVersionParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTimeseriesGroupHTTPVersionParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]RadarHTTPTimeseriesGroupHTTPVersionParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHTTPTimeseriesGroupHTTPVersionParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseriesGroupHTTPVersionParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTimeseriesGroupHTTPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseriesGroupHTTPVersionParamsAggInterval string

const (
	RadarHTTPTimeseriesGroupHTTPVersionParamsAggInterval15m RadarHTTPTimeseriesGroupHTTPVersionParamsAggInterval = "15m"
	RadarHTTPTimeseriesGroupHTTPVersionParamsAggInterval1h  RadarHTTPTimeseriesGroupHTTPVersionParamsAggInterval = "1h"
	RadarHTTPTimeseriesGroupHTTPVersionParamsAggInterval1d  RadarHTTPTimeseriesGroupHTTPVersionParamsAggInterval = "1d"
	RadarHTTPTimeseriesGroupHTTPVersionParamsAggInterval1w  RadarHTTPTimeseriesGroupHTTPVersionParamsAggInterval = "1w"
)

type RadarHTTPTimeseriesGroupHTTPVersionParamsBotClass string

const (
	RadarHTTPTimeseriesGroupHTTPVersionParamsBotClassLikelyAutomated RadarHTTPTimeseriesGroupHTTPVersionParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTimeseriesGroupHTTPVersionParamsBotClassLikelyHuman     RadarHTTPTimeseriesGroupHTTPVersionParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange string

const (
	RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange1d         RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange = "1d"
	RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange2d         RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange = "2d"
	RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange7d         RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange = "7d"
	RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange14d        RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange = "14d"
	RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange28d        RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange = "28d"
	RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange12w        RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange = "12w"
	RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange24w        RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange = "24w"
	RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange52w        RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange = "52w"
	RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange1dControl  RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange = "1dControl"
	RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange2dControl  RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange = "2dControl"
	RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange7dControl  RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange = "7dControl"
	RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange14dControl RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange = "14dControl"
	RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange28dControl RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange = "28dControl"
	RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange12wControl RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange = "12wControl"
	RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange24wControl RadarHTTPTimeseriesGroupHTTPVersionParamsDateRange = "24wControl"
)

type RadarHTTPTimeseriesGroupHTTPVersionParamsDeviceType string

const (
	RadarHTTPTimeseriesGroupHTTPVersionParamsDeviceTypeDesktop RadarHTTPTimeseriesGroupHTTPVersionParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseriesGroupHTTPVersionParamsDeviceTypeMobile  RadarHTTPTimeseriesGroupHTTPVersionParamsDeviceType = "MOBILE"
	RadarHTTPTimeseriesGroupHTTPVersionParamsDeviceTypeOther   RadarHTTPTimeseriesGroupHTTPVersionParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTimeseriesGroupHTTPVersionParamsFormat string

const (
	RadarHTTPTimeseriesGroupHTTPVersionParamsFormatJson RadarHTTPTimeseriesGroupHTTPVersionParamsFormat = "JSON"
	RadarHTTPTimeseriesGroupHTTPVersionParamsFormatCsv  RadarHTTPTimeseriesGroupHTTPVersionParamsFormat = "CSV"
)

type RadarHTTPTimeseriesGroupHTTPVersionParamsHTTPProtocol string

const (
	RadarHTTPTimeseriesGroupHTTPVersionParamsHTTPProtocolHTTP  RadarHTTPTimeseriesGroupHTTPVersionParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseriesGroupHTTPVersionParamsHTTPProtocolHTTPS RadarHTTPTimeseriesGroupHTTPVersionParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTimeseriesGroupHTTPVersionParamsIPVersion string

const (
	RadarHTTPTimeseriesGroupHTTPVersionParamsIPVersionIPv4 RadarHTTPTimeseriesGroupHTTPVersionParamsIPVersion = "IPv4"
	RadarHTTPTimeseriesGroupHTTPVersionParamsIPVersionIPv6 RadarHTTPTimeseriesGroupHTTPVersionParamsIPVersion = "IPv6"
)

type RadarHTTPTimeseriesGroupHTTPVersionParamsOS string

const (
	RadarHTTPTimeseriesGroupHTTPVersionParamsOSWindows  RadarHTTPTimeseriesGroupHTTPVersionParamsOS = "WINDOWS"
	RadarHTTPTimeseriesGroupHTTPVersionParamsOSMacosx   RadarHTTPTimeseriesGroupHTTPVersionParamsOS = "MACOSX"
	RadarHTTPTimeseriesGroupHTTPVersionParamsOSIos      RadarHTTPTimeseriesGroupHTTPVersionParamsOS = "IOS"
	RadarHTTPTimeseriesGroupHTTPVersionParamsOSAndroid  RadarHTTPTimeseriesGroupHTTPVersionParamsOS = "ANDROID"
	RadarHTTPTimeseriesGroupHTTPVersionParamsOSChromeos RadarHTTPTimeseriesGroupHTTPVersionParamsOS = "CHROMEOS"
	RadarHTTPTimeseriesGroupHTTPVersionParamsOSLinux    RadarHTTPTimeseriesGroupHTTPVersionParamsOS = "LINUX"
	RadarHTTPTimeseriesGroupHTTPVersionParamsOSSmartTv  RadarHTTPTimeseriesGroupHTTPVersionParamsOS = "SMART_TV"
)

type RadarHTTPTimeseriesGroupHTTPVersionParamsTLSVersion string

const (
	RadarHTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSv1_0  RadarHTTPTimeseriesGroupHTTPVersionParamsTLSVersion = "TLSv1_0"
	RadarHTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSv1_1  RadarHTTPTimeseriesGroupHTTPVersionParamsTLSVersion = "TLSv1_1"
	RadarHTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSv1_2  RadarHTTPTimeseriesGroupHTTPVersionParamsTLSVersion = "TLSv1_2"
	RadarHTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSv1_3  RadarHTTPTimeseriesGroupHTTPVersionParamsTLSVersion = "TLSv1_3"
	RadarHTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSvQuic RadarHTTPTimeseriesGroupHTTPVersionParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPTimeseriesGroupHTTPVersionResponseEnvelope struct {
	Result  RadarHTTPTimeseriesGroupHTTPVersionResponse             `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    radarHTTPTimeseriesGroupHTTPVersionResponseEnvelopeJSON `json:"-"`
}

// radarHTTPTimeseriesGroupHTTPVersionResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarHTTPTimeseriesGroupHTTPVersionResponseEnvelope]
type radarHTTPTimeseriesGroupHTTPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupHTTPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupHTTPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupIPVersionParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseriesGroupIPVersionParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTimeseriesGroupIPVersionParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTimeseriesGroupIPVersionParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTimeseriesGroupIPVersionParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTimeseriesGroupIPVersionParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTimeseriesGroupIPVersionParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTimeseriesGroupIPVersionParamsHTTPVersion] `query:"httpVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]RadarHTTPTimeseriesGroupIPVersionParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHTTPTimeseriesGroupIPVersionParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseriesGroupIPVersionParams]'s query parameters
// as `url.Values`.
func (r RadarHTTPTimeseriesGroupIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseriesGroupIPVersionParamsAggInterval string

const (
	RadarHTTPTimeseriesGroupIPVersionParamsAggInterval15m RadarHTTPTimeseriesGroupIPVersionParamsAggInterval = "15m"
	RadarHTTPTimeseriesGroupIPVersionParamsAggInterval1h  RadarHTTPTimeseriesGroupIPVersionParamsAggInterval = "1h"
	RadarHTTPTimeseriesGroupIPVersionParamsAggInterval1d  RadarHTTPTimeseriesGroupIPVersionParamsAggInterval = "1d"
	RadarHTTPTimeseriesGroupIPVersionParamsAggInterval1w  RadarHTTPTimeseriesGroupIPVersionParamsAggInterval = "1w"
)

type RadarHTTPTimeseriesGroupIPVersionParamsBotClass string

const (
	RadarHTTPTimeseriesGroupIPVersionParamsBotClassLikelyAutomated RadarHTTPTimeseriesGroupIPVersionParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTimeseriesGroupIPVersionParamsBotClassLikelyHuman     RadarHTTPTimeseriesGroupIPVersionParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTimeseriesGroupIPVersionParamsDateRange string

const (
	RadarHTTPTimeseriesGroupIPVersionParamsDateRange1d         RadarHTTPTimeseriesGroupIPVersionParamsDateRange = "1d"
	RadarHTTPTimeseriesGroupIPVersionParamsDateRange2d         RadarHTTPTimeseriesGroupIPVersionParamsDateRange = "2d"
	RadarHTTPTimeseriesGroupIPVersionParamsDateRange7d         RadarHTTPTimeseriesGroupIPVersionParamsDateRange = "7d"
	RadarHTTPTimeseriesGroupIPVersionParamsDateRange14d        RadarHTTPTimeseriesGroupIPVersionParamsDateRange = "14d"
	RadarHTTPTimeseriesGroupIPVersionParamsDateRange28d        RadarHTTPTimeseriesGroupIPVersionParamsDateRange = "28d"
	RadarHTTPTimeseriesGroupIPVersionParamsDateRange12w        RadarHTTPTimeseriesGroupIPVersionParamsDateRange = "12w"
	RadarHTTPTimeseriesGroupIPVersionParamsDateRange24w        RadarHTTPTimeseriesGroupIPVersionParamsDateRange = "24w"
	RadarHTTPTimeseriesGroupIPVersionParamsDateRange52w        RadarHTTPTimeseriesGroupIPVersionParamsDateRange = "52w"
	RadarHTTPTimeseriesGroupIPVersionParamsDateRange1dControl  RadarHTTPTimeseriesGroupIPVersionParamsDateRange = "1dControl"
	RadarHTTPTimeseriesGroupIPVersionParamsDateRange2dControl  RadarHTTPTimeseriesGroupIPVersionParamsDateRange = "2dControl"
	RadarHTTPTimeseriesGroupIPVersionParamsDateRange7dControl  RadarHTTPTimeseriesGroupIPVersionParamsDateRange = "7dControl"
	RadarHTTPTimeseriesGroupIPVersionParamsDateRange14dControl RadarHTTPTimeseriesGroupIPVersionParamsDateRange = "14dControl"
	RadarHTTPTimeseriesGroupIPVersionParamsDateRange28dControl RadarHTTPTimeseriesGroupIPVersionParamsDateRange = "28dControl"
	RadarHTTPTimeseriesGroupIPVersionParamsDateRange12wControl RadarHTTPTimeseriesGroupIPVersionParamsDateRange = "12wControl"
	RadarHTTPTimeseriesGroupIPVersionParamsDateRange24wControl RadarHTTPTimeseriesGroupIPVersionParamsDateRange = "24wControl"
)

type RadarHTTPTimeseriesGroupIPVersionParamsDeviceType string

const (
	RadarHTTPTimeseriesGroupIPVersionParamsDeviceTypeDesktop RadarHTTPTimeseriesGroupIPVersionParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseriesGroupIPVersionParamsDeviceTypeMobile  RadarHTTPTimeseriesGroupIPVersionParamsDeviceType = "MOBILE"
	RadarHTTPTimeseriesGroupIPVersionParamsDeviceTypeOther   RadarHTTPTimeseriesGroupIPVersionParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTimeseriesGroupIPVersionParamsFormat string

const (
	RadarHTTPTimeseriesGroupIPVersionParamsFormatJson RadarHTTPTimeseriesGroupIPVersionParamsFormat = "JSON"
	RadarHTTPTimeseriesGroupIPVersionParamsFormatCsv  RadarHTTPTimeseriesGroupIPVersionParamsFormat = "CSV"
)

type RadarHTTPTimeseriesGroupIPVersionParamsHTTPProtocol string

const (
	RadarHTTPTimeseriesGroupIPVersionParamsHTTPProtocolHTTP  RadarHTTPTimeseriesGroupIPVersionParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseriesGroupIPVersionParamsHTTPProtocolHTTPS RadarHTTPTimeseriesGroupIPVersionParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTimeseriesGroupIPVersionParamsHTTPVersion string

const (
	RadarHTTPTimeseriesGroupIPVersionParamsHTTPVersionHttPv1 RadarHTTPTimeseriesGroupIPVersionParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseriesGroupIPVersionParamsHTTPVersionHttPv2 RadarHTTPTimeseriesGroupIPVersionParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseriesGroupIPVersionParamsHTTPVersionHttPv3 RadarHTTPTimeseriesGroupIPVersionParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTimeseriesGroupIPVersionParamsOS string

const (
	RadarHTTPTimeseriesGroupIPVersionParamsOSWindows  RadarHTTPTimeseriesGroupIPVersionParamsOS = "WINDOWS"
	RadarHTTPTimeseriesGroupIPVersionParamsOSMacosx   RadarHTTPTimeseriesGroupIPVersionParamsOS = "MACOSX"
	RadarHTTPTimeseriesGroupIPVersionParamsOSIos      RadarHTTPTimeseriesGroupIPVersionParamsOS = "IOS"
	RadarHTTPTimeseriesGroupIPVersionParamsOSAndroid  RadarHTTPTimeseriesGroupIPVersionParamsOS = "ANDROID"
	RadarHTTPTimeseriesGroupIPVersionParamsOSChromeos RadarHTTPTimeseriesGroupIPVersionParamsOS = "CHROMEOS"
	RadarHTTPTimeseriesGroupIPVersionParamsOSLinux    RadarHTTPTimeseriesGroupIPVersionParamsOS = "LINUX"
	RadarHTTPTimeseriesGroupIPVersionParamsOSSmartTv  RadarHTTPTimeseriesGroupIPVersionParamsOS = "SMART_TV"
)

type RadarHTTPTimeseriesGroupIPVersionParamsTLSVersion string

const (
	RadarHTTPTimeseriesGroupIPVersionParamsTLSVersionTlSv1_0  RadarHTTPTimeseriesGroupIPVersionParamsTLSVersion = "TLSv1_0"
	RadarHTTPTimeseriesGroupIPVersionParamsTLSVersionTlSv1_1  RadarHTTPTimeseriesGroupIPVersionParamsTLSVersion = "TLSv1_1"
	RadarHTTPTimeseriesGroupIPVersionParamsTLSVersionTlSv1_2  RadarHTTPTimeseriesGroupIPVersionParamsTLSVersion = "TLSv1_2"
	RadarHTTPTimeseriesGroupIPVersionParamsTLSVersionTlSv1_3  RadarHTTPTimeseriesGroupIPVersionParamsTLSVersion = "TLSv1_3"
	RadarHTTPTimeseriesGroupIPVersionParamsTLSVersionTlSvQuic RadarHTTPTimeseriesGroupIPVersionParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPTimeseriesGroupIPVersionResponseEnvelope struct {
	Result  RadarHTTPTimeseriesGroupIPVersionResponse             `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    radarHTTPTimeseriesGroupIPVersionResponseEnvelopeJSON `json:"-"`
}

// radarHTTPTimeseriesGroupIPVersionResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarHTTPTimeseriesGroupIPVersionResponseEnvelope]
type radarHTTPTimeseriesGroupIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupIPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupOSParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseriesGroupOSParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTimeseriesGroupOSParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTimeseriesGroupOSParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTimeseriesGroupOSParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTimeseriesGroupOSParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTimeseriesGroupOSParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTimeseriesGroupOSParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTimeseriesGroupOSParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHTTPTimeseriesGroupOSParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseriesGroupOSParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPTimeseriesGroupOSParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseriesGroupOSParamsAggInterval string

const (
	RadarHTTPTimeseriesGroupOSParamsAggInterval15m RadarHTTPTimeseriesGroupOSParamsAggInterval = "15m"
	RadarHTTPTimeseriesGroupOSParamsAggInterval1h  RadarHTTPTimeseriesGroupOSParamsAggInterval = "1h"
	RadarHTTPTimeseriesGroupOSParamsAggInterval1d  RadarHTTPTimeseriesGroupOSParamsAggInterval = "1d"
	RadarHTTPTimeseriesGroupOSParamsAggInterval1w  RadarHTTPTimeseriesGroupOSParamsAggInterval = "1w"
)

type RadarHTTPTimeseriesGroupOSParamsBotClass string

const (
	RadarHTTPTimeseriesGroupOSParamsBotClassLikelyAutomated RadarHTTPTimeseriesGroupOSParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTimeseriesGroupOSParamsBotClassLikelyHuman     RadarHTTPTimeseriesGroupOSParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTimeseriesGroupOSParamsDateRange string

const (
	RadarHTTPTimeseriesGroupOSParamsDateRange1d         RadarHTTPTimeseriesGroupOSParamsDateRange = "1d"
	RadarHTTPTimeseriesGroupOSParamsDateRange2d         RadarHTTPTimeseriesGroupOSParamsDateRange = "2d"
	RadarHTTPTimeseriesGroupOSParamsDateRange7d         RadarHTTPTimeseriesGroupOSParamsDateRange = "7d"
	RadarHTTPTimeseriesGroupOSParamsDateRange14d        RadarHTTPTimeseriesGroupOSParamsDateRange = "14d"
	RadarHTTPTimeseriesGroupOSParamsDateRange28d        RadarHTTPTimeseriesGroupOSParamsDateRange = "28d"
	RadarHTTPTimeseriesGroupOSParamsDateRange12w        RadarHTTPTimeseriesGroupOSParamsDateRange = "12w"
	RadarHTTPTimeseriesGroupOSParamsDateRange24w        RadarHTTPTimeseriesGroupOSParamsDateRange = "24w"
	RadarHTTPTimeseriesGroupOSParamsDateRange52w        RadarHTTPTimeseriesGroupOSParamsDateRange = "52w"
	RadarHTTPTimeseriesGroupOSParamsDateRange1dControl  RadarHTTPTimeseriesGroupOSParamsDateRange = "1dControl"
	RadarHTTPTimeseriesGroupOSParamsDateRange2dControl  RadarHTTPTimeseriesGroupOSParamsDateRange = "2dControl"
	RadarHTTPTimeseriesGroupOSParamsDateRange7dControl  RadarHTTPTimeseriesGroupOSParamsDateRange = "7dControl"
	RadarHTTPTimeseriesGroupOSParamsDateRange14dControl RadarHTTPTimeseriesGroupOSParamsDateRange = "14dControl"
	RadarHTTPTimeseriesGroupOSParamsDateRange28dControl RadarHTTPTimeseriesGroupOSParamsDateRange = "28dControl"
	RadarHTTPTimeseriesGroupOSParamsDateRange12wControl RadarHTTPTimeseriesGroupOSParamsDateRange = "12wControl"
	RadarHTTPTimeseriesGroupOSParamsDateRange24wControl RadarHTTPTimeseriesGroupOSParamsDateRange = "24wControl"
)

type RadarHTTPTimeseriesGroupOSParamsDeviceType string

const (
	RadarHTTPTimeseriesGroupOSParamsDeviceTypeDesktop RadarHTTPTimeseriesGroupOSParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseriesGroupOSParamsDeviceTypeMobile  RadarHTTPTimeseriesGroupOSParamsDeviceType = "MOBILE"
	RadarHTTPTimeseriesGroupOSParamsDeviceTypeOther   RadarHTTPTimeseriesGroupOSParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTimeseriesGroupOSParamsFormat string

const (
	RadarHTTPTimeseriesGroupOSParamsFormatJson RadarHTTPTimeseriesGroupOSParamsFormat = "JSON"
	RadarHTTPTimeseriesGroupOSParamsFormatCsv  RadarHTTPTimeseriesGroupOSParamsFormat = "CSV"
)

type RadarHTTPTimeseriesGroupOSParamsHTTPProtocol string

const (
	RadarHTTPTimeseriesGroupOSParamsHTTPProtocolHTTP  RadarHTTPTimeseriesGroupOSParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseriesGroupOSParamsHTTPProtocolHTTPS RadarHTTPTimeseriesGroupOSParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTimeseriesGroupOSParamsHTTPVersion string

const (
	RadarHTTPTimeseriesGroupOSParamsHTTPVersionHttPv1 RadarHTTPTimeseriesGroupOSParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseriesGroupOSParamsHTTPVersionHttPv2 RadarHTTPTimeseriesGroupOSParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseriesGroupOSParamsHTTPVersionHttPv3 RadarHTTPTimeseriesGroupOSParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTimeseriesGroupOSParamsIPVersion string

const (
	RadarHTTPTimeseriesGroupOSParamsIPVersionIPv4 RadarHTTPTimeseriesGroupOSParamsIPVersion = "IPv4"
	RadarHTTPTimeseriesGroupOSParamsIPVersionIPv6 RadarHTTPTimeseriesGroupOSParamsIPVersion = "IPv6"
)

type RadarHTTPTimeseriesGroupOSParamsTLSVersion string

const (
	RadarHTTPTimeseriesGroupOSParamsTLSVersionTlSv1_0  RadarHTTPTimeseriesGroupOSParamsTLSVersion = "TLSv1_0"
	RadarHTTPTimeseriesGroupOSParamsTLSVersionTlSv1_1  RadarHTTPTimeseriesGroupOSParamsTLSVersion = "TLSv1_1"
	RadarHTTPTimeseriesGroupOSParamsTLSVersionTlSv1_2  RadarHTTPTimeseriesGroupOSParamsTLSVersion = "TLSv1_2"
	RadarHTTPTimeseriesGroupOSParamsTLSVersionTlSv1_3  RadarHTTPTimeseriesGroupOSParamsTLSVersion = "TLSv1_3"
	RadarHTTPTimeseriesGroupOSParamsTLSVersionTlSvQuic RadarHTTPTimeseriesGroupOSParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPTimeseriesGroupOSResponseEnvelope struct {
	Result  RadarHTTPTimeseriesGroupOSResponse             `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    radarHTTPTimeseriesGroupOSResponseEnvelopeJSON `json:"-"`
}

// radarHTTPTimeseriesGroupOSResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarHTTPTimeseriesGroupOSResponseEnvelope]
type radarHTTPTimeseriesGroupOSResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupOSResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupOSResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupTLSVersionParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseriesGroupTLSVersionParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTimeseriesGroupTLSVersionParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTimeseriesGroupTLSVersionParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTimeseriesGroupTLSVersionParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTimeseriesGroupTLSVersionParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTimeseriesGroupTLSVersionParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTimeseriesGroupTLSVersionParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTimeseriesGroupTLSVersionParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]RadarHTTPTimeseriesGroupTLSVersionParamsOS] `query:"os"`
}

// URLQuery serializes [RadarHTTPTimeseriesGroupTLSVersionParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTimeseriesGroupTLSVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseriesGroupTLSVersionParamsAggInterval string

const (
	RadarHTTPTimeseriesGroupTLSVersionParamsAggInterval15m RadarHTTPTimeseriesGroupTLSVersionParamsAggInterval = "15m"
	RadarHTTPTimeseriesGroupTLSVersionParamsAggInterval1h  RadarHTTPTimeseriesGroupTLSVersionParamsAggInterval = "1h"
	RadarHTTPTimeseriesGroupTLSVersionParamsAggInterval1d  RadarHTTPTimeseriesGroupTLSVersionParamsAggInterval = "1d"
	RadarHTTPTimeseriesGroupTLSVersionParamsAggInterval1w  RadarHTTPTimeseriesGroupTLSVersionParamsAggInterval = "1w"
)

type RadarHTTPTimeseriesGroupTLSVersionParamsBotClass string

const (
	RadarHTTPTimeseriesGroupTLSVersionParamsBotClassLikelyAutomated RadarHTTPTimeseriesGroupTLSVersionParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTimeseriesGroupTLSVersionParamsBotClassLikelyHuman     RadarHTTPTimeseriesGroupTLSVersionParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTimeseriesGroupTLSVersionParamsDateRange string

const (
	RadarHTTPTimeseriesGroupTLSVersionParamsDateRange1d         RadarHTTPTimeseriesGroupTLSVersionParamsDateRange = "1d"
	RadarHTTPTimeseriesGroupTLSVersionParamsDateRange2d         RadarHTTPTimeseriesGroupTLSVersionParamsDateRange = "2d"
	RadarHTTPTimeseriesGroupTLSVersionParamsDateRange7d         RadarHTTPTimeseriesGroupTLSVersionParamsDateRange = "7d"
	RadarHTTPTimeseriesGroupTLSVersionParamsDateRange14d        RadarHTTPTimeseriesGroupTLSVersionParamsDateRange = "14d"
	RadarHTTPTimeseriesGroupTLSVersionParamsDateRange28d        RadarHTTPTimeseriesGroupTLSVersionParamsDateRange = "28d"
	RadarHTTPTimeseriesGroupTLSVersionParamsDateRange12w        RadarHTTPTimeseriesGroupTLSVersionParamsDateRange = "12w"
	RadarHTTPTimeseriesGroupTLSVersionParamsDateRange24w        RadarHTTPTimeseriesGroupTLSVersionParamsDateRange = "24w"
	RadarHTTPTimeseriesGroupTLSVersionParamsDateRange52w        RadarHTTPTimeseriesGroupTLSVersionParamsDateRange = "52w"
	RadarHTTPTimeseriesGroupTLSVersionParamsDateRange1dControl  RadarHTTPTimeseriesGroupTLSVersionParamsDateRange = "1dControl"
	RadarHTTPTimeseriesGroupTLSVersionParamsDateRange2dControl  RadarHTTPTimeseriesGroupTLSVersionParamsDateRange = "2dControl"
	RadarHTTPTimeseriesGroupTLSVersionParamsDateRange7dControl  RadarHTTPTimeseriesGroupTLSVersionParamsDateRange = "7dControl"
	RadarHTTPTimeseriesGroupTLSVersionParamsDateRange14dControl RadarHTTPTimeseriesGroupTLSVersionParamsDateRange = "14dControl"
	RadarHTTPTimeseriesGroupTLSVersionParamsDateRange28dControl RadarHTTPTimeseriesGroupTLSVersionParamsDateRange = "28dControl"
	RadarHTTPTimeseriesGroupTLSVersionParamsDateRange12wControl RadarHTTPTimeseriesGroupTLSVersionParamsDateRange = "12wControl"
	RadarHTTPTimeseriesGroupTLSVersionParamsDateRange24wControl RadarHTTPTimeseriesGroupTLSVersionParamsDateRange = "24wControl"
)

type RadarHTTPTimeseriesGroupTLSVersionParamsDeviceType string

const (
	RadarHTTPTimeseriesGroupTLSVersionParamsDeviceTypeDesktop RadarHTTPTimeseriesGroupTLSVersionParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseriesGroupTLSVersionParamsDeviceTypeMobile  RadarHTTPTimeseriesGroupTLSVersionParamsDeviceType = "MOBILE"
	RadarHTTPTimeseriesGroupTLSVersionParamsDeviceTypeOther   RadarHTTPTimeseriesGroupTLSVersionParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTimeseriesGroupTLSVersionParamsFormat string

const (
	RadarHTTPTimeseriesGroupTLSVersionParamsFormatJson RadarHTTPTimeseriesGroupTLSVersionParamsFormat = "JSON"
	RadarHTTPTimeseriesGroupTLSVersionParamsFormatCsv  RadarHTTPTimeseriesGroupTLSVersionParamsFormat = "CSV"
)

type RadarHTTPTimeseriesGroupTLSVersionParamsHTTPProtocol string

const (
	RadarHTTPTimeseriesGroupTLSVersionParamsHTTPProtocolHTTP  RadarHTTPTimeseriesGroupTLSVersionParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseriesGroupTLSVersionParamsHTTPProtocolHTTPS RadarHTTPTimeseriesGroupTLSVersionParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTimeseriesGroupTLSVersionParamsHTTPVersion string

const (
	RadarHTTPTimeseriesGroupTLSVersionParamsHTTPVersionHttPv1 RadarHTTPTimeseriesGroupTLSVersionParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseriesGroupTLSVersionParamsHTTPVersionHttPv2 RadarHTTPTimeseriesGroupTLSVersionParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseriesGroupTLSVersionParamsHTTPVersionHttPv3 RadarHTTPTimeseriesGroupTLSVersionParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTimeseriesGroupTLSVersionParamsIPVersion string

const (
	RadarHTTPTimeseriesGroupTLSVersionParamsIPVersionIPv4 RadarHTTPTimeseriesGroupTLSVersionParamsIPVersion = "IPv4"
	RadarHTTPTimeseriesGroupTLSVersionParamsIPVersionIPv6 RadarHTTPTimeseriesGroupTLSVersionParamsIPVersion = "IPv6"
)

type RadarHTTPTimeseriesGroupTLSVersionParamsOS string

const (
	RadarHTTPTimeseriesGroupTLSVersionParamsOSWindows  RadarHTTPTimeseriesGroupTLSVersionParamsOS = "WINDOWS"
	RadarHTTPTimeseriesGroupTLSVersionParamsOSMacosx   RadarHTTPTimeseriesGroupTLSVersionParamsOS = "MACOSX"
	RadarHTTPTimeseriesGroupTLSVersionParamsOSIos      RadarHTTPTimeseriesGroupTLSVersionParamsOS = "IOS"
	RadarHTTPTimeseriesGroupTLSVersionParamsOSAndroid  RadarHTTPTimeseriesGroupTLSVersionParamsOS = "ANDROID"
	RadarHTTPTimeseriesGroupTLSVersionParamsOSChromeos RadarHTTPTimeseriesGroupTLSVersionParamsOS = "CHROMEOS"
	RadarHTTPTimeseriesGroupTLSVersionParamsOSLinux    RadarHTTPTimeseriesGroupTLSVersionParamsOS = "LINUX"
	RadarHTTPTimeseriesGroupTLSVersionParamsOSSmartTv  RadarHTTPTimeseriesGroupTLSVersionParamsOS = "SMART_TV"
)

type RadarHTTPTimeseriesGroupTLSVersionResponseEnvelope struct {
	Result  RadarHTTPTimeseriesGroupTLSVersionResponse             `json:"result,required"`
	Success bool                                                   `json:"success,required"`
	JSON    radarHTTPTimeseriesGroupTLSVersionResponseEnvelopeJSON `json:"-"`
}

// radarHTTPTimeseriesGroupTLSVersionResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarHTTPTimeseriesGroupTLSVersionResponseEnvelope]
type radarHTTPTimeseriesGroupTLSVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupTLSVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupTLSVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
