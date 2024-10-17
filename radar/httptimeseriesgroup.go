// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// HTTPTimeseriesGroupService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTTPTimeseriesGroupService] method instead.
type HTTPTimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewHTTPTimeseriesGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewHTTPTimeseriesGroupService(opts ...option.RequestOption) (r *HTTPTimeseriesGroupService) {
	r = &HTTPTimeseriesGroupService{}
	r.Options = opts
	return
}

// Get a time series of the percentage distribution of traffic classified as
// automated or human. Visit
// https://developers.cloudflare.com/radar/concepts/bot-classes/ for more
// information.
func (r *HTTPTimeseriesGroupService) BotClass(ctx context.Context, query HTTPTimeseriesGroupBotClassParams, opts ...option.RequestOption) (res *HTTPTimeseriesGroupBotClassResponse, err error) {
	var env HTTPTimeseriesGroupBotClassResponseEnvelope
	opts = append(r.Options[:], opts...)
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
func (r *HTTPTimeseriesGroupService) Browser(ctx context.Context, query HTTPTimeseriesGroupBrowserParams, opts ...option.RequestOption) (res *HTTPTimeseriesGroupBrowserResponse, err error) {
	var env HTTPTimeseriesGroupBrowserResponseEnvelope
	opts = append(r.Options[:], opts...)
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
func (r *HTTPTimeseriesGroupService) BrowserFamily(ctx context.Context, query HTTPTimeseriesGroupBrowserFamilyParams, opts ...option.RequestOption) (res *HTTPTimeseriesGroupBrowserFamilyResponse, err error) {
	var env HTTPTimeseriesGroupBrowserFamilyResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/browser_family"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a time series of the percentage distribution of traffic by device type.
func (r *HTTPTimeseriesGroupService) DeviceType(ctx context.Context, query HTTPTimeseriesGroupDeviceTypeParams, opts ...option.RequestOption) (res *HTTPTimeseriesGroupDeviceTypeResponse, err error) {
	var env HTTPTimeseriesGroupDeviceTypeResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/device_type"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a time series of the percentage distribution of traffic by HTTP protocol.
func (r *HTTPTimeseriesGroupService) HTTPProtocol(ctx context.Context, query HTTPTimeseriesGroupHTTPProtocolParams, opts ...option.RequestOption) (res *HTTPTimeseriesGroupHTTPProtocolResponse, err error) {
	var env HTTPTimeseriesGroupHTTPProtocolResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/http_protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a time series of the percentage distribution of traffic by HTTP version.
func (r *HTTPTimeseriesGroupService) HTTPVersion(ctx context.Context, query HTTPTimeseriesGroupHTTPVersionParams, opts ...option.RequestOption) (res *HTTPTimeseriesGroupHTTPVersionResponse, err error) {
	var env HTTPTimeseriesGroupHTTPVersionResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/http_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a time series of the percentage distribution of traffic by IP version.
func (r *HTTPTimeseriesGroupService) IPVersion(ctx context.Context, query HTTPTimeseriesGroupIPVersionParams, opts ...option.RequestOption) (res *HTTPTimeseriesGroupIPVersionResponse, err error) {
	var env HTTPTimeseriesGroupIPVersionResponseEnvelope
	opts = append(r.Options[:], opts...)
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
func (r *HTTPTimeseriesGroupService) OS(ctx context.Context, query HTTPTimeseriesGroupOSParams, opts ...option.RequestOption) (res *HTTPTimeseriesGroupOSResponse, err error) {
	var env HTTPTimeseriesGroupOSResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/os"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a time series of the percentage distribution of traffic by post-quantum
// suport.
func (r *HTTPTimeseriesGroupService) PostQuantum(ctx context.Context, query HTTPTimeseriesGroupPostQuantumParams, opts ...option.RequestOption) (res *HTTPTimeseriesGroupPostQuantumResponse, err error) {
	var env HTTPTimeseriesGroupPostQuantumResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/post_quantum"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a time series of the percentage distribution of traffic by TLS protocol
// version.
func (r *HTTPTimeseriesGroupService) TLSVersion(ctx context.Context, query HTTPTimeseriesGroupTLSVersionParams, opts ...option.RequestOption) (res *HTTPTimeseriesGroupTLSVersionResponse, err error) {
	var env HTTPTimeseriesGroupTLSVersionResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/tls_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPTimeseriesGroupBotClassResponse struct {
	Meta   interface{}                               `json:"meta,required"`
	Serie0 HTTPTimeseriesGroupBotClassResponseSerie0 `json:"serie_0,required"`
	JSON   httpTimeseriesGroupBotClassResponseJSON   `json:"-"`
}

// httpTimeseriesGroupBotClassResponseJSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupBotClassResponse]
type httpTimeseriesGroupBotClassResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupBotClassResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBotClassResponseJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupBotClassResponseSerie0 struct {
	Bot        []string                                      `json:"bot,required"`
	Human      []string                                      `json:"human,required"`
	Timestamps []string                                      `json:"timestamps,required"`
	JSON       httpTimeseriesGroupBotClassResponseSerie0JSON `json:"-"`
}

// httpTimeseriesGroupBotClassResponseSerie0JSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupBotClassResponseSerie0]
type httpTimeseriesGroupBotClassResponseSerie0JSON struct {
	Bot         apijson.Field
	Human       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupBotClassResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBotClassResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupBrowserResponse struct {
	Meta   interface{}                              `json:"meta,required"`
	Serie0 HTTPTimeseriesGroupBrowserResponseSerie0 `json:"serie_0,required"`
	JSON   httpTimeseriesGroupBrowserResponseJSON   `json:"-"`
}

// httpTimeseriesGroupBrowserResponseJSON contains the JSON metadata for the struct
// [HTTPTimeseriesGroupBrowserResponse]
type httpTimeseriesGroupBrowserResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupBrowserResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBrowserResponseJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupBrowserResponseSerie0 struct {
	Timestamps  []string                                     `json:"timestamps,required"`
	ExtraFields map[string][]string                          `json:"-,extras"`
	JSON        httpTimeseriesGroupBrowserResponseSerie0JSON `json:"-"`
}

// httpTimeseriesGroupBrowserResponseSerie0JSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupBrowserResponseSerie0]
type httpTimeseriesGroupBrowserResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupBrowserResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBrowserResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupBrowserFamilyResponse struct {
	Meta   interface{}                                    `json:"meta,required"`
	Serie0 HTTPTimeseriesGroupBrowserFamilyResponseSerie0 `json:"serie_0,required"`
	JSON   httpTimeseriesGroupBrowserFamilyResponseJSON   `json:"-"`
}

// httpTimeseriesGroupBrowserFamilyResponseJSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupBrowserFamilyResponse]
type httpTimeseriesGroupBrowserFamilyResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupBrowserFamilyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBrowserFamilyResponseJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupBrowserFamilyResponseSerie0 struct {
	Timestamps  []string                                           `json:"timestamps,required"`
	ExtraFields map[string][]string                                `json:"-,extras"`
	JSON        httpTimeseriesGroupBrowserFamilyResponseSerie0JSON `json:"-"`
}

// httpTimeseriesGroupBrowserFamilyResponseSerie0JSON contains the JSON metadata
// for the struct [HTTPTimeseriesGroupBrowserFamilyResponseSerie0]
type httpTimeseriesGroupBrowserFamilyResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupBrowserFamilyResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBrowserFamilyResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupDeviceTypeResponse struct {
	Meta   interface{}                                 `json:"meta,required"`
	Serie0 HTTPTimeseriesGroupDeviceTypeResponseSerie0 `json:"serie_0,required"`
	JSON   httpTimeseriesGroupDeviceTypeResponseJSON   `json:"-"`
}

// httpTimeseriesGroupDeviceTypeResponseJSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupDeviceTypeResponse]
type httpTimeseriesGroupDeviceTypeResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupDeviceTypeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupDeviceTypeResponseJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupDeviceTypeResponseSerie0 struct {
	Desktop    []string                                        `json:"desktop,required"`
	Mobile     []string                                        `json:"mobile,required"`
	Other      []string                                        `json:"other,required"`
	Timestamps []string                                        `json:"timestamps,required"`
	JSON       httpTimeseriesGroupDeviceTypeResponseSerie0JSON `json:"-"`
}

// httpTimeseriesGroupDeviceTypeResponseSerie0JSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupDeviceTypeResponseSerie0]
type httpTimeseriesGroupDeviceTypeResponseSerie0JSON struct {
	Desktop     apijson.Field
	Mobile      apijson.Field
	Other       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupDeviceTypeResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupDeviceTypeResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupHTTPProtocolResponse struct {
	Meta   interface{}                                   `json:"meta,required"`
	Serie0 HTTPTimeseriesGroupHTTPProtocolResponseSerie0 `json:"serie_0,required"`
	JSON   httpTimeseriesGroupHTTPProtocolResponseJSON   `json:"-"`
}

// httpTimeseriesGroupHTTPProtocolResponseJSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupHTTPProtocolResponse]
type httpTimeseriesGroupHTTPProtocolResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupHTTPProtocolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupHTTPProtocolResponseJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupHTTPProtocolResponseSerie0 struct {
	HTTP       []string                                          `json:"http,required"`
	HTTPS      []string                                          `json:"https,required"`
	Timestamps []string                                          `json:"timestamps,required"`
	JSON       httpTimeseriesGroupHTTPProtocolResponseSerie0JSON `json:"-"`
}

// httpTimeseriesGroupHTTPProtocolResponseSerie0JSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupHTTPProtocolResponseSerie0]
type httpTimeseriesGroupHTTPProtocolResponseSerie0JSON struct {
	HTTP        apijson.Field
	HTTPS       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupHTTPProtocolResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupHTTPProtocolResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupHTTPVersionResponse struct {
	Meta   interface{}                                  `json:"meta,required"`
	Serie0 HTTPTimeseriesGroupHTTPVersionResponseSerie0 `json:"serie_0,required"`
	JSON   httpTimeseriesGroupHTTPVersionResponseJSON   `json:"-"`
}

// httpTimeseriesGroupHTTPVersionResponseJSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupHTTPVersionResponse]
type httpTimeseriesGroupHTTPVersionResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupHTTPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupHTTPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupHTTPVersionResponseSerie0 struct {
	HTTP1X     []string                                         `json:"HTTP/1.x,required"`
	HTTP2      []string                                         `json:"HTTP/2,required"`
	HTTP3      []string                                         `json:"HTTP/3,required"`
	Timestamps []string                                         `json:"timestamps,required"`
	JSON       httpTimeseriesGroupHTTPVersionResponseSerie0JSON `json:"-"`
}

// httpTimeseriesGroupHTTPVersionResponseSerie0JSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupHTTPVersionResponseSerie0]
type httpTimeseriesGroupHTTPVersionResponseSerie0JSON struct {
	HTTP1X      apijson.Field
	HTTP2       apijson.Field
	HTTP3       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupHTTPVersionResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupHTTPVersionResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupIPVersionResponse struct {
	Meta   interface{}                                `json:"meta,required"`
	Serie0 HTTPTimeseriesGroupIPVersionResponseSerie0 `json:"serie_0,required"`
	JSON   httpTimeseriesGroupIPVersionResponseJSON   `json:"-"`
}

// httpTimeseriesGroupIPVersionResponseJSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupIPVersionResponse]
type httpTimeseriesGroupIPVersionResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupIPVersionResponseSerie0 struct {
	IPv4       []string                                       `json:"IPv4,required"`
	IPv6       []string                                       `json:"IPv6,required"`
	Timestamps []string                                       `json:"timestamps,required"`
	JSON       httpTimeseriesGroupIPVersionResponseSerie0JSON `json:"-"`
}

// httpTimeseriesGroupIPVersionResponseSerie0JSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupIPVersionResponseSerie0]
type httpTimeseriesGroupIPVersionResponseSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupIPVersionResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupIPVersionResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupOSResponse struct {
	Meta   interface{}                         `json:"meta,required"`
	Serie0 HTTPTimeseriesGroupOSResponseSerie0 `json:"serie_0,required"`
	JSON   httpTimeseriesGroupOSResponseJSON   `json:"-"`
}

// httpTimeseriesGroupOSResponseJSON contains the JSON metadata for the struct
// [HTTPTimeseriesGroupOSResponse]
type httpTimeseriesGroupOSResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupOSResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupOSResponseJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupOSResponseSerie0 struct {
	Timestamps  []string                                `json:"timestamps,required"`
	ExtraFields map[string][]string                     `json:"-,extras"`
	JSON        httpTimeseriesGroupOSResponseSerie0JSON `json:"-"`
}

// httpTimeseriesGroupOSResponseSerie0JSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupOSResponseSerie0]
type httpTimeseriesGroupOSResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupOSResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupOSResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupPostQuantumResponse struct {
	Meta   interface{}                                  `json:"meta,required"`
	Serie0 HTTPTimeseriesGroupPostQuantumResponseSerie0 `json:"serie_0,required"`
	JSON   httpTimeseriesGroupPostQuantumResponseJSON   `json:"-"`
}

// httpTimeseriesGroupPostQuantumResponseJSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupPostQuantumResponse]
type httpTimeseriesGroupPostQuantumResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupPostQuantumResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupPostQuantumResponseJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupPostQuantumResponseSerie0 struct {
	NotSupported []string                                         `json:"NOT_SUPPORTED,required"`
	Supported    []string                                         `json:"SUPPORTED,required"`
	Timestamps   []string                                         `json:"timestamps,required"`
	JSON         httpTimeseriesGroupPostQuantumResponseSerie0JSON `json:"-"`
}

// httpTimeseriesGroupPostQuantumResponseSerie0JSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupPostQuantumResponseSerie0]
type httpTimeseriesGroupPostQuantumResponseSerie0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	Timestamps   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupPostQuantumResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupPostQuantumResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupTLSVersionResponse struct {
	Meta   interface{}                                 `json:"meta,required"`
	Serie0 HTTPTimeseriesGroupTLSVersionResponseSerie0 `json:"serie_0,required"`
	JSON   httpTimeseriesGroupTLSVersionResponseJSON   `json:"-"`
}

// httpTimeseriesGroupTLSVersionResponseJSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupTLSVersionResponse]
type httpTimeseriesGroupTLSVersionResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupTLSVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupTLSVersionResponseJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupTLSVersionResponseSerie0 struct {
	Timestamps []string                                        `json:"timestamps,required"`
	TLS1_0     []string                                        `json:"TLS 1.0,required"`
	TLS1_1     []string                                        `json:"TLS 1.1,required"`
	TLS1_2     []string                                        `json:"TLS 1.2,required"`
	TLS1_3     []string                                        `json:"TLS 1.3,required"`
	TLSQuic    []string                                        `json:"TLS QUIC,required"`
	JSON       httpTimeseriesGroupTLSVersionResponseSerie0JSON `json:"-"`
}

// httpTimeseriesGroupTLSVersionResponseSerie0JSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupTLSVersionResponseSerie0]
type httpTimeseriesGroupTLSVersionResponseSerie0JSON struct {
	Timestamps  apijson.Field
	TLS1_0      apijson.Field
	TLS1_1      apijson.Field
	TLS1_2      apijson.Field
	TLS1_3      apijson.Field
	TLSQuic     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupTLSVersionResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupTLSVersionResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupBotClassParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[HTTPTimeseriesGroupBotClassParamsAggInterval] `query:"aggInterval"`
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
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]HTTPTimeseriesGroupBotClassParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[HTTPTimeseriesGroupBotClassParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]HTTPTimeseriesGroupBotClassParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]HTTPTimeseriesGroupBotClassParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]HTTPTimeseriesGroupBotClassParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]HTTPTimeseriesGroupBotClassParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]HTTPTimeseriesGroupBotClassParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPTimeseriesGroupBotClassParams]'s query parameters as
// `url.Values`.
func (r HTTPTimeseriesGroupBotClassParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupBotClassParamsAggInterval string

const (
	HTTPTimeseriesGroupBotClassParamsAggInterval15m HTTPTimeseriesGroupBotClassParamsAggInterval = "15m"
	HTTPTimeseriesGroupBotClassParamsAggInterval1h  HTTPTimeseriesGroupBotClassParamsAggInterval = "1h"
	HTTPTimeseriesGroupBotClassParamsAggInterval1d  HTTPTimeseriesGroupBotClassParamsAggInterval = "1d"
	HTTPTimeseriesGroupBotClassParamsAggInterval1w  HTTPTimeseriesGroupBotClassParamsAggInterval = "1w"
)

func (r HTTPTimeseriesGroupBotClassParamsAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBotClassParamsAggInterval15m, HTTPTimeseriesGroupBotClassParamsAggInterval1h, HTTPTimeseriesGroupBotClassParamsAggInterval1d, HTTPTimeseriesGroupBotClassParamsAggInterval1w:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBotClassParamsDeviceType string

const (
	HTTPTimeseriesGroupBotClassParamsDeviceTypeDesktop HTTPTimeseriesGroupBotClassParamsDeviceType = "DESKTOP"
	HTTPTimeseriesGroupBotClassParamsDeviceTypeMobile  HTTPTimeseriesGroupBotClassParamsDeviceType = "MOBILE"
	HTTPTimeseriesGroupBotClassParamsDeviceTypeOther   HTTPTimeseriesGroupBotClassParamsDeviceType = "OTHER"
)

func (r HTTPTimeseriesGroupBotClassParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBotClassParamsDeviceTypeDesktop, HTTPTimeseriesGroupBotClassParamsDeviceTypeMobile, HTTPTimeseriesGroupBotClassParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format results are returned in.
type HTTPTimeseriesGroupBotClassParamsFormat string

const (
	HTTPTimeseriesGroupBotClassParamsFormatJson HTTPTimeseriesGroupBotClassParamsFormat = "JSON"
	HTTPTimeseriesGroupBotClassParamsFormatCsv  HTTPTimeseriesGroupBotClassParamsFormat = "CSV"
)

func (r HTTPTimeseriesGroupBotClassParamsFormat) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBotClassParamsFormatJson, HTTPTimeseriesGroupBotClassParamsFormatCsv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBotClassParamsHTTPProtocol string

const (
	HTTPTimeseriesGroupBotClassParamsHTTPProtocolHTTP  HTTPTimeseriesGroupBotClassParamsHTTPProtocol = "HTTP"
	HTTPTimeseriesGroupBotClassParamsHTTPProtocolHTTPS HTTPTimeseriesGroupBotClassParamsHTTPProtocol = "HTTPS"
)

func (r HTTPTimeseriesGroupBotClassParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBotClassParamsHTTPProtocolHTTP, HTTPTimeseriesGroupBotClassParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBotClassParamsHTTPVersion string

const (
	HTTPTimeseriesGroupBotClassParamsHTTPVersionHttPv1 HTTPTimeseriesGroupBotClassParamsHTTPVersion = "HTTPv1"
	HTTPTimeseriesGroupBotClassParamsHTTPVersionHttPv2 HTTPTimeseriesGroupBotClassParamsHTTPVersion = "HTTPv2"
	HTTPTimeseriesGroupBotClassParamsHTTPVersionHttPv3 HTTPTimeseriesGroupBotClassParamsHTTPVersion = "HTTPv3"
)

func (r HTTPTimeseriesGroupBotClassParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBotClassParamsHTTPVersionHttPv1, HTTPTimeseriesGroupBotClassParamsHTTPVersionHttPv2, HTTPTimeseriesGroupBotClassParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBotClassParamsIPVersion string

const (
	HTTPTimeseriesGroupBotClassParamsIPVersionIPv4 HTTPTimeseriesGroupBotClassParamsIPVersion = "IPv4"
	HTTPTimeseriesGroupBotClassParamsIPVersionIPv6 HTTPTimeseriesGroupBotClassParamsIPVersion = "IPv6"
)

func (r HTTPTimeseriesGroupBotClassParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBotClassParamsIPVersionIPv4, HTTPTimeseriesGroupBotClassParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBotClassParamsOS string

const (
	HTTPTimeseriesGroupBotClassParamsOSWindows  HTTPTimeseriesGroupBotClassParamsOS = "WINDOWS"
	HTTPTimeseriesGroupBotClassParamsOSMacosx   HTTPTimeseriesGroupBotClassParamsOS = "MACOSX"
	HTTPTimeseriesGroupBotClassParamsOSIos      HTTPTimeseriesGroupBotClassParamsOS = "IOS"
	HTTPTimeseriesGroupBotClassParamsOSAndroid  HTTPTimeseriesGroupBotClassParamsOS = "ANDROID"
	HTTPTimeseriesGroupBotClassParamsOSChromeos HTTPTimeseriesGroupBotClassParamsOS = "CHROMEOS"
	HTTPTimeseriesGroupBotClassParamsOSLinux    HTTPTimeseriesGroupBotClassParamsOS = "LINUX"
	HTTPTimeseriesGroupBotClassParamsOSSmartTv  HTTPTimeseriesGroupBotClassParamsOS = "SMART_TV"
)

func (r HTTPTimeseriesGroupBotClassParamsOS) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBotClassParamsOSWindows, HTTPTimeseriesGroupBotClassParamsOSMacosx, HTTPTimeseriesGroupBotClassParamsOSIos, HTTPTimeseriesGroupBotClassParamsOSAndroid, HTTPTimeseriesGroupBotClassParamsOSChromeos, HTTPTimeseriesGroupBotClassParamsOSLinux, HTTPTimeseriesGroupBotClassParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBotClassParamsTLSVersion string

const (
	HTTPTimeseriesGroupBotClassParamsTLSVersionTlSv1_0  HTTPTimeseriesGroupBotClassParamsTLSVersion = "TLSv1_0"
	HTTPTimeseriesGroupBotClassParamsTLSVersionTlSv1_1  HTTPTimeseriesGroupBotClassParamsTLSVersion = "TLSv1_1"
	HTTPTimeseriesGroupBotClassParamsTLSVersionTlSv1_2  HTTPTimeseriesGroupBotClassParamsTLSVersion = "TLSv1_2"
	HTTPTimeseriesGroupBotClassParamsTLSVersionTlSv1_3  HTTPTimeseriesGroupBotClassParamsTLSVersion = "TLSv1_3"
	HTTPTimeseriesGroupBotClassParamsTLSVersionTlSvQuic HTTPTimeseriesGroupBotClassParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPTimeseriesGroupBotClassParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBotClassParamsTLSVersionTlSv1_0, HTTPTimeseriesGroupBotClassParamsTLSVersionTlSv1_1, HTTPTimeseriesGroupBotClassParamsTLSVersionTlSv1_2, HTTPTimeseriesGroupBotClassParamsTLSVersionTlSv1_3, HTTPTimeseriesGroupBotClassParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBotClassResponseEnvelope struct {
	Result  HTTPTimeseriesGroupBotClassResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    httpTimeseriesGroupBotClassResponseEnvelopeJSON `json:"-"`
}

// httpTimeseriesGroupBotClassResponseEnvelopeJSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupBotClassResponseEnvelope]
type httpTimeseriesGroupBotClassResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupBotClassResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBotClassResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupBrowserParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[HTTPTimeseriesGroupBrowserParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPTimeseriesGroupBrowserParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]HTTPTimeseriesGroupBrowserParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[HTTPTimeseriesGroupBrowserParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]HTTPTimeseriesGroupBrowserParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]HTTPTimeseriesGroupBrowserParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]HTTPTimeseriesGroupBrowserParamsIPVersion] `query:"ipVersion"`
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
	OS param.Field[[]HTTPTimeseriesGroupBrowserParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]HTTPTimeseriesGroupBrowserParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPTimeseriesGroupBrowserParams]'s query parameters as
// `url.Values`.
func (r HTTPTimeseriesGroupBrowserParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupBrowserParamsAggInterval string

const (
	HTTPTimeseriesGroupBrowserParamsAggInterval15m HTTPTimeseriesGroupBrowserParamsAggInterval = "15m"
	HTTPTimeseriesGroupBrowserParamsAggInterval1h  HTTPTimeseriesGroupBrowserParamsAggInterval = "1h"
	HTTPTimeseriesGroupBrowserParamsAggInterval1d  HTTPTimeseriesGroupBrowserParamsAggInterval = "1d"
	HTTPTimeseriesGroupBrowserParamsAggInterval1w  HTTPTimeseriesGroupBrowserParamsAggInterval = "1w"
)

func (r HTTPTimeseriesGroupBrowserParamsAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserParamsAggInterval15m, HTTPTimeseriesGroupBrowserParamsAggInterval1h, HTTPTimeseriesGroupBrowserParamsAggInterval1d, HTTPTimeseriesGroupBrowserParamsAggInterval1w:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserParamsBotClass string

const (
	HTTPTimeseriesGroupBrowserParamsBotClassLikelyAutomated HTTPTimeseriesGroupBrowserParamsBotClass = "LIKELY_AUTOMATED"
	HTTPTimeseriesGroupBrowserParamsBotClassLikelyHuman     HTTPTimeseriesGroupBrowserParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPTimeseriesGroupBrowserParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserParamsBotClassLikelyAutomated, HTTPTimeseriesGroupBrowserParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserParamsDeviceType string

const (
	HTTPTimeseriesGroupBrowserParamsDeviceTypeDesktop HTTPTimeseriesGroupBrowserParamsDeviceType = "DESKTOP"
	HTTPTimeseriesGroupBrowserParamsDeviceTypeMobile  HTTPTimeseriesGroupBrowserParamsDeviceType = "MOBILE"
	HTTPTimeseriesGroupBrowserParamsDeviceTypeOther   HTTPTimeseriesGroupBrowserParamsDeviceType = "OTHER"
)

func (r HTTPTimeseriesGroupBrowserParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserParamsDeviceTypeDesktop, HTTPTimeseriesGroupBrowserParamsDeviceTypeMobile, HTTPTimeseriesGroupBrowserParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format results are returned in.
type HTTPTimeseriesGroupBrowserParamsFormat string

const (
	HTTPTimeseriesGroupBrowserParamsFormatJson HTTPTimeseriesGroupBrowserParamsFormat = "JSON"
	HTTPTimeseriesGroupBrowserParamsFormatCsv  HTTPTimeseriesGroupBrowserParamsFormat = "CSV"
)

func (r HTTPTimeseriesGroupBrowserParamsFormat) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserParamsFormatJson, HTTPTimeseriesGroupBrowserParamsFormatCsv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserParamsHTTPProtocol string

const (
	HTTPTimeseriesGroupBrowserParamsHTTPProtocolHTTP  HTTPTimeseriesGroupBrowserParamsHTTPProtocol = "HTTP"
	HTTPTimeseriesGroupBrowserParamsHTTPProtocolHTTPS HTTPTimeseriesGroupBrowserParamsHTTPProtocol = "HTTPS"
)

func (r HTTPTimeseriesGroupBrowserParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserParamsHTTPProtocolHTTP, HTTPTimeseriesGroupBrowserParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserParamsHTTPVersion string

const (
	HTTPTimeseriesGroupBrowserParamsHTTPVersionHttPv1 HTTPTimeseriesGroupBrowserParamsHTTPVersion = "HTTPv1"
	HTTPTimeseriesGroupBrowserParamsHTTPVersionHttPv2 HTTPTimeseriesGroupBrowserParamsHTTPVersion = "HTTPv2"
	HTTPTimeseriesGroupBrowserParamsHTTPVersionHttPv3 HTTPTimeseriesGroupBrowserParamsHTTPVersion = "HTTPv3"
)

func (r HTTPTimeseriesGroupBrowserParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserParamsHTTPVersionHttPv1, HTTPTimeseriesGroupBrowserParamsHTTPVersionHttPv2, HTTPTimeseriesGroupBrowserParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserParamsIPVersion string

const (
	HTTPTimeseriesGroupBrowserParamsIPVersionIPv4 HTTPTimeseriesGroupBrowserParamsIPVersion = "IPv4"
	HTTPTimeseriesGroupBrowserParamsIPVersionIPv6 HTTPTimeseriesGroupBrowserParamsIPVersion = "IPv6"
)

func (r HTTPTimeseriesGroupBrowserParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserParamsIPVersionIPv4, HTTPTimeseriesGroupBrowserParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserParamsOS string

const (
	HTTPTimeseriesGroupBrowserParamsOSWindows  HTTPTimeseriesGroupBrowserParamsOS = "WINDOWS"
	HTTPTimeseriesGroupBrowserParamsOSMacosx   HTTPTimeseriesGroupBrowserParamsOS = "MACOSX"
	HTTPTimeseriesGroupBrowserParamsOSIos      HTTPTimeseriesGroupBrowserParamsOS = "IOS"
	HTTPTimeseriesGroupBrowserParamsOSAndroid  HTTPTimeseriesGroupBrowserParamsOS = "ANDROID"
	HTTPTimeseriesGroupBrowserParamsOSChromeos HTTPTimeseriesGroupBrowserParamsOS = "CHROMEOS"
	HTTPTimeseriesGroupBrowserParamsOSLinux    HTTPTimeseriesGroupBrowserParamsOS = "LINUX"
	HTTPTimeseriesGroupBrowserParamsOSSmartTv  HTTPTimeseriesGroupBrowserParamsOS = "SMART_TV"
)

func (r HTTPTimeseriesGroupBrowserParamsOS) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserParamsOSWindows, HTTPTimeseriesGroupBrowserParamsOSMacosx, HTTPTimeseriesGroupBrowserParamsOSIos, HTTPTimeseriesGroupBrowserParamsOSAndroid, HTTPTimeseriesGroupBrowserParamsOSChromeos, HTTPTimeseriesGroupBrowserParamsOSLinux, HTTPTimeseriesGroupBrowserParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserParamsTLSVersion string

const (
	HTTPTimeseriesGroupBrowserParamsTLSVersionTlSv1_0  HTTPTimeseriesGroupBrowserParamsTLSVersion = "TLSv1_0"
	HTTPTimeseriesGroupBrowserParamsTLSVersionTlSv1_1  HTTPTimeseriesGroupBrowserParamsTLSVersion = "TLSv1_1"
	HTTPTimeseriesGroupBrowserParamsTLSVersionTlSv1_2  HTTPTimeseriesGroupBrowserParamsTLSVersion = "TLSv1_2"
	HTTPTimeseriesGroupBrowserParamsTLSVersionTlSv1_3  HTTPTimeseriesGroupBrowserParamsTLSVersion = "TLSv1_3"
	HTTPTimeseriesGroupBrowserParamsTLSVersionTlSvQuic HTTPTimeseriesGroupBrowserParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPTimeseriesGroupBrowserParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserParamsTLSVersionTlSv1_0, HTTPTimeseriesGroupBrowserParamsTLSVersionTlSv1_1, HTTPTimeseriesGroupBrowserParamsTLSVersionTlSv1_2, HTTPTimeseriesGroupBrowserParamsTLSVersionTlSv1_3, HTTPTimeseriesGroupBrowserParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserResponseEnvelope struct {
	Result  HTTPTimeseriesGroupBrowserResponse             `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    httpTimeseriesGroupBrowserResponseEnvelopeJSON `json:"-"`
}

// httpTimeseriesGroupBrowserResponseEnvelopeJSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupBrowserResponseEnvelope]
type httpTimeseriesGroupBrowserResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupBrowserResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBrowserResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupBrowserFamilyParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[HTTPTimeseriesGroupBrowserFamilyParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPTimeseriesGroupBrowserFamilyParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]HTTPTimeseriesGroupBrowserFamilyParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[HTTPTimeseriesGroupBrowserFamilyParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]HTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]HTTPTimeseriesGroupBrowserFamilyParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]HTTPTimeseriesGroupBrowserFamilyParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]HTTPTimeseriesGroupBrowserFamilyParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPTimeseriesGroupBrowserFamilyParams]'s query parameters
// as `url.Values`.
func (r HTTPTimeseriesGroupBrowserFamilyParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupBrowserFamilyParamsAggInterval string

const (
	HTTPTimeseriesGroupBrowserFamilyParamsAggInterval15m HTTPTimeseriesGroupBrowserFamilyParamsAggInterval = "15m"
	HTTPTimeseriesGroupBrowserFamilyParamsAggInterval1h  HTTPTimeseriesGroupBrowserFamilyParamsAggInterval = "1h"
	HTTPTimeseriesGroupBrowserFamilyParamsAggInterval1d  HTTPTimeseriesGroupBrowserFamilyParamsAggInterval = "1d"
	HTTPTimeseriesGroupBrowserFamilyParamsAggInterval1w  HTTPTimeseriesGroupBrowserFamilyParamsAggInterval = "1w"
)

func (r HTTPTimeseriesGroupBrowserFamilyParamsAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserFamilyParamsAggInterval15m, HTTPTimeseriesGroupBrowserFamilyParamsAggInterval1h, HTTPTimeseriesGroupBrowserFamilyParamsAggInterval1d, HTTPTimeseriesGroupBrowserFamilyParamsAggInterval1w:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserFamilyParamsBotClass string

const (
	HTTPTimeseriesGroupBrowserFamilyParamsBotClassLikelyAutomated HTTPTimeseriesGroupBrowserFamilyParamsBotClass = "LIKELY_AUTOMATED"
	HTTPTimeseriesGroupBrowserFamilyParamsBotClassLikelyHuman     HTTPTimeseriesGroupBrowserFamilyParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPTimeseriesGroupBrowserFamilyParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserFamilyParamsBotClassLikelyAutomated, HTTPTimeseriesGroupBrowserFamilyParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserFamilyParamsDeviceType string

const (
	HTTPTimeseriesGroupBrowserFamilyParamsDeviceTypeDesktop HTTPTimeseriesGroupBrowserFamilyParamsDeviceType = "DESKTOP"
	HTTPTimeseriesGroupBrowserFamilyParamsDeviceTypeMobile  HTTPTimeseriesGroupBrowserFamilyParamsDeviceType = "MOBILE"
	HTTPTimeseriesGroupBrowserFamilyParamsDeviceTypeOther   HTTPTimeseriesGroupBrowserFamilyParamsDeviceType = "OTHER"
)

func (r HTTPTimeseriesGroupBrowserFamilyParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserFamilyParamsDeviceTypeDesktop, HTTPTimeseriesGroupBrowserFamilyParamsDeviceTypeMobile, HTTPTimeseriesGroupBrowserFamilyParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format results are returned in.
type HTTPTimeseriesGroupBrowserFamilyParamsFormat string

const (
	HTTPTimeseriesGroupBrowserFamilyParamsFormatJson HTTPTimeseriesGroupBrowserFamilyParamsFormat = "JSON"
	HTTPTimeseriesGroupBrowserFamilyParamsFormatCsv  HTTPTimeseriesGroupBrowserFamilyParamsFormat = "CSV"
)

func (r HTTPTimeseriesGroupBrowserFamilyParamsFormat) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserFamilyParamsFormatJson, HTTPTimeseriesGroupBrowserFamilyParamsFormatCsv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocol string

const (
	HTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocolHTTP  HTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocol = "HTTP"
	HTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocolHTTPS HTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocol = "HTTPS"
)

func (r HTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocolHTTP, HTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersion string

const (
	HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersionHttPv1 HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersion = "HTTPv1"
	HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersionHttPv2 HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersion = "HTTPv2"
	HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersionHttPv3 HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersion = "HTTPv3"
)

func (r HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersionHttPv1, HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersionHttPv2, HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserFamilyParamsIPVersion string

const (
	HTTPTimeseriesGroupBrowserFamilyParamsIPVersionIPv4 HTTPTimeseriesGroupBrowserFamilyParamsIPVersion = "IPv4"
	HTTPTimeseriesGroupBrowserFamilyParamsIPVersionIPv6 HTTPTimeseriesGroupBrowserFamilyParamsIPVersion = "IPv6"
)

func (r HTTPTimeseriesGroupBrowserFamilyParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserFamilyParamsIPVersionIPv4, HTTPTimeseriesGroupBrowserFamilyParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserFamilyParamsOS string

const (
	HTTPTimeseriesGroupBrowserFamilyParamsOSWindows  HTTPTimeseriesGroupBrowserFamilyParamsOS = "WINDOWS"
	HTTPTimeseriesGroupBrowserFamilyParamsOSMacosx   HTTPTimeseriesGroupBrowserFamilyParamsOS = "MACOSX"
	HTTPTimeseriesGroupBrowserFamilyParamsOSIos      HTTPTimeseriesGroupBrowserFamilyParamsOS = "IOS"
	HTTPTimeseriesGroupBrowserFamilyParamsOSAndroid  HTTPTimeseriesGroupBrowserFamilyParamsOS = "ANDROID"
	HTTPTimeseriesGroupBrowserFamilyParamsOSChromeos HTTPTimeseriesGroupBrowserFamilyParamsOS = "CHROMEOS"
	HTTPTimeseriesGroupBrowserFamilyParamsOSLinux    HTTPTimeseriesGroupBrowserFamilyParamsOS = "LINUX"
	HTTPTimeseriesGroupBrowserFamilyParamsOSSmartTv  HTTPTimeseriesGroupBrowserFamilyParamsOS = "SMART_TV"
)

func (r HTTPTimeseriesGroupBrowserFamilyParamsOS) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserFamilyParamsOSWindows, HTTPTimeseriesGroupBrowserFamilyParamsOSMacosx, HTTPTimeseriesGroupBrowserFamilyParamsOSIos, HTTPTimeseriesGroupBrowserFamilyParamsOSAndroid, HTTPTimeseriesGroupBrowserFamilyParamsOSChromeos, HTTPTimeseriesGroupBrowserFamilyParamsOSLinux, HTTPTimeseriesGroupBrowserFamilyParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserFamilyParamsTLSVersion string

const (
	HTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSv1_0  HTTPTimeseriesGroupBrowserFamilyParamsTLSVersion = "TLSv1_0"
	HTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSv1_1  HTTPTimeseriesGroupBrowserFamilyParamsTLSVersion = "TLSv1_1"
	HTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSv1_2  HTTPTimeseriesGroupBrowserFamilyParamsTLSVersion = "TLSv1_2"
	HTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSv1_3  HTTPTimeseriesGroupBrowserFamilyParamsTLSVersion = "TLSv1_3"
	HTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSvQuic HTTPTimeseriesGroupBrowserFamilyParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPTimeseriesGroupBrowserFamilyParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSv1_0, HTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSv1_1, HTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSv1_2, HTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSv1_3, HTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserFamilyResponseEnvelope struct {
	Result  HTTPTimeseriesGroupBrowserFamilyResponse             `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    httpTimeseriesGroupBrowserFamilyResponseEnvelopeJSON `json:"-"`
}

// httpTimeseriesGroupBrowserFamilyResponseEnvelopeJSON contains the JSON metadata
// for the struct [HTTPTimeseriesGroupBrowserFamilyResponseEnvelope]
type httpTimeseriesGroupBrowserFamilyResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupBrowserFamilyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBrowserFamilyResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupDeviceTypeParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[HTTPTimeseriesGroupDeviceTypeParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPTimeseriesGroupDeviceTypeParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[HTTPTimeseriesGroupDeviceTypeParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]HTTPTimeseriesGroupDeviceTypeParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]HTTPTimeseriesGroupDeviceTypeParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]HTTPTimeseriesGroupDeviceTypeParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]HTTPTimeseriesGroupDeviceTypeParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]HTTPTimeseriesGroupDeviceTypeParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPTimeseriesGroupDeviceTypeParams]'s query parameters as
// `url.Values`.
func (r HTTPTimeseriesGroupDeviceTypeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupDeviceTypeParamsAggInterval string

const (
	HTTPTimeseriesGroupDeviceTypeParamsAggInterval15m HTTPTimeseriesGroupDeviceTypeParamsAggInterval = "15m"
	HTTPTimeseriesGroupDeviceTypeParamsAggInterval1h  HTTPTimeseriesGroupDeviceTypeParamsAggInterval = "1h"
	HTTPTimeseriesGroupDeviceTypeParamsAggInterval1d  HTTPTimeseriesGroupDeviceTypeParamsAggInterval = "1d"
	HTTPTimeseriesGroupDeviceTypeParamsAggInterval1w  HTTPTimeseriesGroupDeviceTypeParamsAggInterval = "1w"
)

func (r HTTPTimeseriesGroupDeviceTypeParamsAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupDeviceTypeParamsAggInterval15m, HTTPTimeseriesGroupDeviceTypeParamsAggInterval1h, HTTPTimeseriesGroupDeviceTypeParamsAggInterval1d, HTTPTimeseriesGroupDeviceTypeParamsAggInterval1w:
		return true
	}
	return false
}

type HTTPTimeseriesGroupDeviceTypeParamsBotClass string

const (
	HTTPTimeseriesGroupDeviceTypeParamsBotClassLikelyAutomated HTTPTimeseriesGroupDeviceTypeParamsBotClass = "LIKELY_AUTOMATED"
	HTTPTimeseriesGroupDeviceTypeParamsBotClassLikelyHuman     HTTPTimeseriesGroupDeviceTypeParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPTimeseriesGroupDeviceTypeParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupDeviceTypeParamsBotClassLikelyAutomated, HTTPTimeseriesGroupDeviceTypeParamsBotClassLikelyHuman:
		return true
	}
	return false
}

// Format results are returned in.
type HTTPTimeseriesGroupDeviceTypeParamsFormat string

const (
	HTTPTimeseriesGroupDeviceTypeParamsFormatJson HTTPTimeseriesGroupDeviceTypeParamsFormat = "JSON"
	HTTPTimeseriesGroupDeviceTypeParamsFormatCsv  HTTPTimeseriesGroupDeviceTypeParamsFormat = "CSV"
)

func (r HTTPTimeseriesGroupDeviceTypeParamsFormat) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupDeviceTypeParamsFormatJson, HTTPTimeseriesGroupDeviceTypeParamsFormatCsv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupDeviceTypeParamsHTTPProtocol string

const (
	HTTPTimeseriesGroupDeviceTypeParamsHTTPProtocolHTTP  HTTPTimeseriesGroupDeviceTypeParamsHTTPProtocol = "HTTP"
	HTTPTimeseriesGroupDeviceTypeParamsHTTPProtocolHTTPS HTTPTimeseriesGroupDeviceTypeParamsHTTPProtocol = "HTTPS"
)

func (r HTTPTimeseriesGroupDeviceTypeParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupDeviceTypeParamsHTTPProtocolHTTP, HTTPTimeseriesGroupDeviceTypeParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPTimeseriesGroupDeviceTypeParamsHTTPVersion string

const (
	HTTPTimeseriesGroupDeviceTypeParamsHTTPVersionHttPv1 HTTPTimeseriesGroupDeviceTypeParamsHTTPVersion = "HTTPv1"
	HTTPTimeseriesGroupDeviceTypeParamsHTTPVersionHttPv2 HTTPTimeseriesGroupDeviceTypeParamsHTTPVersion = "HTTPv2"
	HTTPTimeseriesGroupDeviceTypeParamsHTTPVersionHttPv3 HTTPTimeseriesGroupDeviceTypeParamsHTTPVersion = "HTTPv3"
)

func (r HTTPTimeseriesGroupDeviceTypeParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupDeviceTypeParamsHTTPVersionHttPv1, HTTPTimeseriesGroupDeviceTypeParamsHTTPVersionHttPv2, HTTPTimeseriesGroupDeviceTypeParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPTimeseriesGroupDeviceTypeParamsIPVersion string

const (
	HTTPTimeseriesGroupDeviceTypeParamsIPVersionIPv4 HTTPTimeseriesGroupDeviceTypeParamsIPVersion = "IPv4"
	HTTPTimeseriesGroupDeviceTypeParamsIPVersionIPv6 HTTPTimeseriesGroupDeviceTypeParamsIPVersion = "IPv6"
)

func (r HTTPTimeseriesGroupDeviceTypeParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupDeviceTypeParamsIPVersionIPv4, HTTPTimeseriesGroupDeviceTypeParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPTimeseriesGroupDeviceTypeParamsOS string

const (
	HTTPTimeseriesGroupDeviceTypeParamsOSWindows  HTTPTimeseriesGroupDeviceTypeParamsOS = "WINDOWS"
	HTTPTimeseriesGroupDeviceTypeParamsOSMacosx   HTTPTimeseriesGroupDeviceTypeParamsOS = "MACOSX"
	HTTPTimeseriesGroupDeviceTypeParamsOSIos      HTTPTimeseriesGroupDeviceTypeParamsOS = "IOS"
	HTTPTimeseriesGroupDeviceTypeParamsOSAndroid  HTTPTimeseriesGroupDeviceTypeParamsOS = "ANDROID"
	HTTPTimeseriesGroupDeviceTypeParamsOSChromeos HTTPTimeseriesGroupDeviceTypeParamsOS = "CHROMEOS"
	HTTPTimeseriesGroupDeviceTypeParamsOSLinux    HTTPTimeseriesGroupDeviceTypeParamsOS = "LINUX"
	HTTPTimeseriesGroupDeviceTypeParamsOSSmartTv  HTTPTimeseriesGroupDeviceTypeParamsOS = "SMART_TV"
)

func (r HTTPTimeseriesGroupDeviceTypeParamsOS) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupDeviceTypeParamsOSWindows, HTTPTimeseriesGroupDeviceTypeParamsOSMacosx, HTTPTimeseriesGroupDeviceTypeParamsOSIos, HTTPTimeseriesGroupDeviceTypeParamsOSAndroid, HTTPTimeseriesGroupDeviceTypeParamsOSChromeos, HTTPTimeseriesGroupDeviceTypeParamsOSLinux, HTTPTimeseriesGroupDeviceTypeParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupDeviceTypeParamsTLSVersion string

const (
	HTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSv1_0  HTTPTimeseriesGroupDeviceTypeParamsTLSVersion = "TLSv1_0"
	HTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSv1_1  HTTPTimeseriesGroupDeviceTypeParamsTLSVersion = "TLSv1_1"
	HTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSv1_2  HTTPTimeseriesGroupDeviceTypeParamsTLSVersion = "TLSv1_2"
	HTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSv1_3  HTTPTimeseriesGroupDeviceTypeParamsTLSVersion = "TLSv1_3"
	HTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSvQuic HTTPTimeseriesGroupDeviceTypeParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPTimeseriesGroupDeviceTypeParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSv1_0, HTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSv1_1, HTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSv1_2, HTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSv1_3, HTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPTimeseriesGroupDeviceTypeResponseEnvelope struct {
	Result  HTTPTimeseriesGroupDeviceTypeResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    httpTimeseriesGroupDeviceTypeResponseEnvelopeJSON `json:"-"`
}

// httpTimeseriesGroupDeviceTypeResponseEnvelopeJSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupDeviceTypeResponseEnvelope]
type httpTimeseriesGroupDeviceTypeResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupDeviceTypeResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupDeviceTypeResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupHTTPProtocolParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[HTTPTimeseriesGroupHTTPProtocolParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPTimeseriesGroupHTTPProtocolParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]HTTPTimeseriesGroupHTTPProtocolParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[HTTPTimeseriesGroupHTTPProtocolParamsFormat] `query:"format"`
	// Filter for http version.
	HTTPVersion param.Field[[]HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]HTTPTimeseriesGroupHTTPProtocolParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]HTTPTimeseriesGroupHTTPProtocolParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]HTTPTimeseriesGroupHTTPProtocolParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPTimeseriesGroupHTTPProtocolParams]'s query parameters
// as `url.Values`.
func (r HTTPTimeseriesGroupHTTPProtocolParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupHTTPProtocolParamsAggInterval string

const (
	HTTPTimeseriesGroupHTTPProtocolParamsAggInterval15m HTTPTimeseriesGroupHTTPProtocolParamsAggInterval = "15m"
	HTTPTimeseriesGroupHTTPProtocolParamsAggInterval1h  HTTPTimeseriesGroupHTTPProtocolParamsAggInterval = "1h"
	HTTPTimeseriesGroupHTTPProtocolParamsAggInterval1d  HTTPTimeseriesGroupHTTPProtocolParamsAggInterval = "1d"
	HTTPTimeseriesGroupHTTPProtocolParamsAggInterval1w  HTTPTimeseriesGroupHTTPProtocolParamsAggInterval = "1w"
)

func (r HTTPTimeseriesGroupHTTPProtocolParamsAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPProtocolParamsAggInterval15m, HTTPTimeseriesGroupHTTPProtocolParamsAggInterval1h, HTTPTimeseriesGroupHTTPProtocolParamsAggInterval1d, HTTPTimeseriesGroupHTTPProtocolParamsAggInterval1w:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPProtocolParamsBotClass string

const (
	HTTPTimeseriesGroupHTTPProtocolParamsBotClassLikelyAutomated HTTPTimeseriesGroupHTTPProtocolParamsBotClass = "LIKELY_AUTOMATED"
	HTTPTimeseriesGroupHTTPProtocolParamsBotClassLikelyHuman     HTTPTimeseriesGroupHTTPProtocolParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPTimeseriesGroupHTTPProtocolParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPProtocolParamsBotClassLikelyAutomated, HTTPTimeseriesGroupHTTPProtocolParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPProtocolParamsDeviceType string

const (
	HTTPTimeseriesGroupHTTPProtocolParamsDeviceTypeDesktop HTTPTimeseriesGroupHTTPProtocolParamsDeviceType = "DESKTOP"
	HTTPTimeseriesGroupHTTPProtocolParamsDeviceTypeMobile  HTTPTimeseriesGroupHTTPProtocolParamsDeviceType = "MOBILE"
	HTTPTimeseriesGroupHTTPProtocolParamsDeviceTypeOther   HTTPTimeseriesGroupHTTPProtocolParamsDeviceType = "OTHER"
)

func (r HTTPTimeseriesGroupHTTPProtocolParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPProtocolParamsDeviceTypeDesktop, HTTPTimeseriesGroupHTTPProtocolParamsDeviceTypeMobile, HTTPTimeseriesGroupHTTPProtocolParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format results are returned in.
type HTTPTimeseriesGroupHTTPProtocolParamsFormat string

const (
	HTTPTimeseriesGroupHTTPProtocolParamsFormatJson HTTPTimeseriesGroupHTTPProtocolParamsFormat = "JSON"
	HTTPTimeseriesGroupHTTPProtocolParamsFormatCsv  HTTPTimeseriesGroupHTTPProtocolParamsFormat = "CSV"
)

func (r HTTPTimeseriesGroupHTTPProtocolParamsFormat) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPProtocolParamsFormatJson, HTTPTimeseriesGroupHTTPProtocolParamsFormatCsv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersion string

const (
	HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersionHttPv1 HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersion = "HTTPv1"
	HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersionHttPv2 HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersion = "HTTPv2"
	HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersionHttPv3 HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersion = "HTTPv3"
)

func (r HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersionHttPv1, HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersionHttPv2, HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPProtocolParamsIPVersion string

const (
	HTTPTimeseriesGroupHTTPProtocolParamsIPVersionIPv4 HTTPTimeseriesGroupHTTPProtocolParamsIPVersion = "IPv4"
	HTTPTimeseriesGroupHTTPProtocolParamsIPVersionIPv6 HTTPTimeseriesGroupHTTPProtocolParamsIPVersion = "IPv6"
)

func (r HTTPTimeseriesGroupHTTPProtocolParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPProtocolParamsIPVersionIPv4, HTTPTimeseriesGroupHTTPProtocolParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPProtocolParamsOS string

const (
	HTTPTimeseriesGroupHTTPProtocolParamsOSWindows  HTTPTimeseriesGroupHTTPProtocolParamsOS = "WINDOWS"
	HTTPTimeseriesGroupHTTPProtocolParamsOSMacosx   HTTPTimeseriesGroupHTTPProtocolParamsOS = "MACOSX"
	HTTPTimeseriesGroupHTTPProtocolParamsOSIos      HTTPTimeseriesGroupHTTPProtocolParamsOS = "IOS"
	HTTPTimeseriesGroupHTTPProtocolParamsOSAndroid  HTTPTimeseriesGroupHTTPProtocolParamsOS = "ANDROID"
	HTTPTimeseriesGroupHTTPProtocolParamsOSChromeos HTTPTimeseriesGroupHTTPProtocolParamsOS = "CHROMEOS"
	HTTPTimeseriesGroupHTTPProtocolParamsOSLinux    HTTPTimeseriesGroupHTTPProtocolParamsOS = "LINUX"
	HTTPTimeseriesGroupHTTPProtocolParamsOSSmartTv  HTTPTimeseriesGroupHTTPProtocolParamsOS = "SMART_TV"
)

func (r HTTPTimeseriesGroupHTTPProtocolParamsOS) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPProtocolParamsOSWindows, HTTPTimeseriesGroupHTTPProtocolParamsOSMacosx, HTTPTimeseriesGroupHTTPProtocolParamsOSIos, HTTPTimeseriesGroupHTTPProtocolParamsOSAndroid, HTTPTimeseriesGroupHTTPProtocolParamsOSChromeos, HTTPTimeseriesGroupHTTPProtocolParamsOSLinux, HTTPTimeseriesGroupHTTPProtocolParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPProtocolParamsTLSVersion string

const (
	HTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSv1_0  HTTPTimeseriesGroupHTTPProtocolParamsTLSVersion = "TLSv1_0"
	HTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSv1_1  HTTPTimeseriesGroupHTTPProtocolParamsTLSVersion = "TLSv1_1"
	HTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSv1_2  HTTPTimeseriesGroupHTTPProtocolParamsTLSVersion = "TLSv1_2"
	HTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSv1_3  HTTPTimeseriesGroupHTTPProtocolParamsTLSVersion = "TLSv1_3"
	HTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSvQuic HTTPTimeseriesGroupHTTPProtocolParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPTimeseriesGroupHTTPProtocolParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSv1_0, HTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSv1_1, HTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSv1_2, HTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSv1_3, HTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPProtocolResponseEnvelope struct {
	Result  HTTPTimeseriesGroupHTTPProtocolResponse             `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    httpTimeseriesGroupHTTPProtocolResponseEnvelopeJSON `json:"-"`
}

// httpTimeseriesGroupHTTPProtocolResponseEnvelopeJSON contains the JSON metadata
// for the struct [HTTPTimeseriesGroupHTTPProtocolResponseEnvelope]
type httpTimeseriesGroupHTTPProtocolResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupHTTPProtocolResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupHTTPProtocolResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupHTTPVersionParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[HTTPTimeseriesGroupHTTPVersionParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPTimeseriesGroupHTTPVersionParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]HTTPTimeseriesGroupHTTPVersionParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[HTTPTimeseriesGroupHTTPVersionParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]HTTPTimeseriesGroupHTTPVersionParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for ip version.
	IPVersion param.Field[[]HTTPTimeseriesGroupHTTPVersionParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]HTTPTimeseriesGroupHTTPVersionParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]HTTPTimeseriesGroupHTTPVersionParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPTimeseriesGroupHTTPVersionParams]'s query parameters as
// `url.Values`.
func (r HTTPTimeseriesGroupHTTPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupHTTPVersionParamsAggInterval string

const (
	HTTPTimeseriesGroupHTTPVersionParamsAggInterval15m HTTPTimeseriesGroupHTTPVersionParamsAggInterval = "15m"
	HTTPTimeseriesGroupHTTPVersionParamsAggInterval1h  HTTPTimeseriesGroupHTTPVersionParamsAggInterval = "1h"
	HTTPTimeseriesGroupHTTPVersionParamsAggInterval1d  HTTPTimeseriesGroupHTTPVersionParamsAggInterval = "1d"
	HTTPTimeseriesGroupHTTPVersionParamsAggInterval1w  HTTPTimeseriesGroupHTTPVersionParamsAggInterval = "1w"
)

func (r HTTPTimeseriesGroupHTTPVersionParamsAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPVersionParamsAggInterval15m, HTTPTimeseriesGroupHTTPVersionParamsAggInterval1h, HTTPTimeseriesGroupHTTPVersionParamsAggInterval1d, HTTPTimeseriesGroupHTTPVersionParamsAggInterval1w:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPVersionParamsBotClass string

const (
	HTTPTimeseriesGroupHTTPVersionParamsBotClassLikelyAutomated HTTPTimeseriesGroupHTTPVersionParamsBotClass = "LIKELY_AUTOMATED"
	HTTPTimeseriesGroupHTTPVersionParamsBotClassLikelyHuman     HTTPTimeseriesGroupHTTPVersionParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPTimeseriesGroupHTTPVersionParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPVersionParamsBotClassLikelyAutomated, HTTPTimeseriesGroupHTTPVersionParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPVersionParamsDeviceType string

const (
	HTTPTimeseriesGroupHTTPVersionParamsDeviceTypeDesktop HTTPTimeseriesGroupHTTPVersionParamsDeviceType = "DESKTOP"
	HTTPTimeseriesGroupHTTPVersionParamsDeviceTypeMobile  HTTPTimeseriesGroupHTTPVersionParamsDeviceType = "MOBILE"
	HTTPTimeseriesGroupHTTPVersionParamsDeviceTypeOther   HTTPTimeseriesGroupHTTPVersionParamsDeviceType = "OTHER"
)

func (r HTTPTimeseriesGroupHTTPVersionParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPVersionParamsDeviceTypeDesktop, HTTPTimeseriesGroupHTTPVersionParamsDeviceTypeMobile, HTTPTimeseriesGroupHTTPVersionParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format results are returned in.
type HTTPTimeseriesGroupHTTPVersionParamsFormat string

const (
	HTTPTimeseriesGroupHTTPVersionParamsFormatJson HTTPTimeseriesGroupHTTPVersionParamsFormat = "JSON"
	HTTPTimeseriesGroupHTTPVersionParamsFormatCsv  HTTPTimeseriesGroupHTTPVersionParamsFormat = "CSV"
)

func (r HTTPTimeseriesGroupHTTPVersionParamsFormat) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPVersionParamsFormatJson, HTTPTimeseriesGroupHTTPVersionParamsFormatCsv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPVersionParamsHTTPProtocol string

const (
	HTTPTimeseriesGroupHTTPVersionParamsHTTPProtocolHTTP  HTTPTimeseriesGroupHTTPVersionParamsHTTPProtocol = "HTTP"
	HTTPTimeseriesGroupHTTPVersionParamsHTTPProtocolHTTPS HTTPTimeseriesGroupHTTPVersionParamsHTTPProtocol = "HTTPS"
)

func (r HTTPTimeseriesGroupHTTPVersionParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPVersionParamsHTTPProtocolHTTP, HTTPTimeseriesGroupHTTPVersionParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPVersionParamsIPVersion string

const (
	HTTPTimeseriesGroupHTTPVersionParamsIPVersionIPv4 HTTPTimeseriesGroupHTTPVersionParamsIPVersion = "IPv4"
	HTTPTimeseriesGroupHTTPVersionParamsIPVersionIPv6 HTTPTimeseriesGroupHTTPVersionParamsIPVersion = "IPv6"
)

func (r HTTPTimeseriesGroupHTTPVersionParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPVersionParamsIPVersionIPv4, HTTPTimeseriesGroupHTTPVersionParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPVersionParamsOS string

const (
	HTTPTimeseriesGroupHTTPVersionParamsOSWindows  HTTPTimeseriesGroupHTTPVersionParamsOS = "WINDOWS"
	HTTPTimeseriesGroupHTTPVersionParamsOSMacosx   HTTPTimeseriesGroupHTTPVersionParamsOS = "MACOSX"
	HTTPTimeseriesGroupHTTPVersionParamsOSIos      HTTPTimeseriesGroupHTTPVersionParamsOS = "IOS"
	HTTPTimeseriesGroupHTTPVersionParamsOSAndroid  HTTPTimeseriesGroupHTTPVersionParamsOS = "ANDROID"
	HTTPTimeseriesGroupHTTPVersionParamsOSChromeos HTTPTimeseriesGroupHTTPVersionParamsOS = "CHROMEOS"
	HTTPTimeseriesGroupHTTPVersionParamsOSLinux    HTTPTimeseriesGroupHTTPVersionParamsOS = "LINUX"
	HTTPTimeseriesGroupHTTPVersionParamsOSSmartTv  HTTPTimeseriesGroupHTTPVersionParamsOS = "SMART_TV"
)

func (r HTTPTimeseriesGroupHTTPVersionParamsOS) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPVersionParamsOSWindows, HTTPTimeseriesGroupHTTPVersionParamsOSMacosx, HTTPTimeseriesGroupHTTPVersionParamsOSIos, HTTPTimeseriesGroupHTTPVersionParamsOSAndroid, HTTPTimeseriesGroupHTTPVersionParamsOSChromeos, HTTPTimeseriesGroupHTTPVersionParamsOSLinux, HTTPTimeseriesGroupHTTPVersionParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPVersionParamsTLSVersion string

const (
	HTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSv1_0  HTTPTimeseriesGroupHTTPVersionParamsTLSVersion = "TLSv1_0"
	HTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSv1_1  HTTPTimeseriesGroupHTTPVersionParamsTLSVersion = "TLSv1_1"
	HTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSv1_2  HTTPTimeseriesGroupHTTPVersionParamsTLSVersion = "TLSv1_2"
	HTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSv1_3  HTTPTimeseriesGroupHTTPVersionParamsTLSVersion = "TLSv1_3"
	HTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSvQuic HTTPTimeseriesGroupHTTPVersionParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPTimeseriesGroupHTTPVersionParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSv1_0, HTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSv1_1, HTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSv1_2, HTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSv1_3, HTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPVersionResponseEnvelope struct {
	Result  HTTPTimeseriesGroupHTTPVersionResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    httpTimeseriesGroupHTTPVersionResponseEnvelopeJSON `json:"-"`
}

// httpTimeseriesGroupHTTPVersionResponseEnvelopeJSON contains the JSON metadata
// for the struct [HTTPTimeseriesGroupHTTPVersionResponseEnvelope]
type httpTimeseriesGroupHTTPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupHTTPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupHTTPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupIPVersionParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[HTTPTimeseriesGroupIPVersionParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPTimeseriesGroupIPVersionParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]HTTPTimeseriesGroupIPVersionParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[HTTPTimeseriesGroupIPVersionParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]HTTPTimeseriesGroupIPVersionParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]HTTPTimeseriesGroupIPVersionParamsHTTPVersion] `query:"httpVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]HTTPTimeseriesGroupIPVersionParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]HTTPTimeseriesGroupIPVersionParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPTimeseriesGroupIPVersionParams]'s query parameters as
// `url.Values`.
func (r HTTPTimeseriesGroupIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupIPVersionParamsAggInterval string

const (
	HTTPTimeseriesGroupIPVersionParamsAggInterval15m HTTPTimeseriesGroupIPVersionParamsAggInterval = "15m"
	HTTPTimeseriesGroupIPVersionParamsAggInterval1h  HTTPTimeseriesGroupIPVersionParamsAggInterval = "1h"
	HTTPTimeseriesGroupIPVersionParamsAggInterval1d  HTTPTimeseriesGroupIPVersionParamsAggInterval = "1d"
	HTTPTimeseriesGroupIPVersionParamsAggInterval1w  HTTPTimeseriesGroupIPVersionParamsAggInterval = "1w"
)

func (r HTTPTimeseriesGroupIPVersionParamsAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupIPVersionParamsAggInterval15m, HTTPTimeseriesGroupIPVersionParamsAggInterval1h, HTTPTimeseriesGroupIPVersionParamsAggInterval1d, HTTPTimeseriesGroupIPVersionParamsAggInterval1w:
		return true
	}
	return false
}

type HTTPTimeseriesGroupIPVersionParamsBotClass string

const (
	HTTPTimeseriesGroupIPVersionParamsBotClassLikelyAutomated HTTPTimeseriesGroupIPVersionParamsBotClass = "LIKELY_AUTOMATED"
	HTTPTimeseriesGroupIPVersionParamsBotClassLikelyHuman     HTTPTimeseriesGroupIPVersionParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPTimeseriesGroupIPVersionParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupIPVersionParamsBotClassLikelyAutomated, HTTPTimeseriesGroupIPVersionParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPTimeseriesGroupIPVersionParamsDeviceType string

const (
	HTTPTimeseriesGroupIPVersionParamsDeviceTypeDesktop HTTPTimeseriesGroupIPVersionParamsDeviceType = "DESKTOP"
	HTTPTimeseriesGroupIPVersionParamsDeviceTypeMobile  HTTPTimeseriesGroupIPVersionParamsDeviceType = "MOBILE"
	HTTPTimeseriesGroupIPVersionParamsDeviceTypeOther   HTTPTimeseriesGroupIPVersionParamsDeviceType = "OTHER"
)

func (r HTTPTimeseriesGroupIPVersionParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupIPVersionParamsDeviceTypeDesktop, HTTPTimeseriesGroupIPVersionParamsDeviceTypeMobile, HTTPTimeseriesGroupIPVersionParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format results are returned in.
type HTTPTimeseriesGroupIPVersionParamsFormat string

const (
	HTTPTimeseriesGroupIPVersionParamsFormatJson HTTPTimeseriesGroupIPVersionParamsFormat = "JSON"
	HTTPTimeseriesGroupIPVersionParamsFormatCsv  HTTPTimeseriesGroupIPVersionParamsFormat = "CSV"
)

func (r HTTPTimeseriesGroupIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupIPVersionParamsFormatJson, HTTPTimeseriesGroupIPVersionParamsFormatCsv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupIPVersionParamsHTTPProtocol string

const (
	HTTPTimeseriesGroupIPVersionParamsHTTPProtocolHTTP  HTTPTimeseriesGroupIPVersionParamsHTTPProtocol = "HTTP"
	HTTPTimeseriesGroupIPVersionParamsHTTPProtocolHTTPS HTTPTimeseriesGroupIPVersionParamsHTTPProtocol = "HTTPS"
)

func (r HTTPTimeseriesGroupIPVersionParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupIPVersionParamsHTTPProtocolHTTP, HTTPTimeseriesGroupIPVersionParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPTimeseriesGroupIPVersionParamsHTTPVersion string

const (
	HTTPTimeseriesGroupIPVersionParamsHTTPVersionHttPv1 HTTPTimeseriesGroupIPVersionParamsHTTPVersion = "HTTPv1"
	HTTPTimeseriesGroupIPVersionParamsHTTPVersionHttPv2 HTTPTimeseriesGroupIPVersionParamsHTTPVersion = "HTTPv2"
	HTTPTimeseriesGroupIPVersionParamsHTTPVersionHttPv3 HTTPTimeseriesGroupIPVersionParamsHTTPVersion = "HTTPv3"
)

func (r HTTPTimeseriesGroupIPVersionParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupIPVersionParamsHTTPVersionHttPv1, HTTPTimeseriesGroupIPVersionParamsHTTPVersionHttPv2, HTTPTimeseriesGroupIPVersionParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPTimeseriesGroupIPVersionParamsOS string

const (
	HTTPTimeseriesGroupIPVersionParamsOSWindows  HTTPTimeseriesGroupIPVersionParamsOS = "WINDOWS"
	HTTPTimeseriesGroupIPVersionParamsOSMacosx   HTTPTimeseriesGroupIPVersionParamsOS = "MACOSX"
	HTTPTimeseriesGroupIPVersionParamsOSIos      HTTPTimeseriesGroupIPVersionParamsOS = "IOS"
	HTTPTimeseriesGroupIPVersionParamsOSAndroid  HTTPTimeseriesGroupIPVersionParamsOS = "ANDROID"
	HTTPTimeseriesGroupIPVersionParamsOSChromeos HTTPTimeseriesGroupIPVersionParamsOS = "CHROMEOS"
	HTTPTimeseriesGroupIPVersionParamsOSLinux    HTTPTimeseriesGroupIPVersionParamsOS = "LINUX"
	HTTPTimeseriesGroupIPVersionParamsOSSmartTv  HTTPTimeseriesGroupIPVersionParamsOS = "SMART_TV"
)

func (r HTTPTimeseriesGroupIPVersionParamsOS) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupIPVersionParamsOSWindows, HTTPTimeseriesGroupIPVersionParamsOSMacosx, HTTPTimeseriesGroupIPVersionParamsOSIos, HTTPTimeseriesGroupIPVersionParamsOSAndroid, HTTPTimeseriesGroupIPVersionParamsOSChromeos, HTTPTimeseriesGroupIPVersionParamsOSLinux, HTTPTimeseriesGroupIPVersionParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupIPVersionParamsTLSVersion string

const (
	HTTPTimeseriesGroupIPVersionParamsTLSVersionTlSv1_0  HTTPTimeseriesGroupIPVersionParamsTLSVersion = "TLSv1_0"
	HTTPTimeseriesGroupIPVersionParamsTLSVersionTlSv1_1  HTTPTimeseriesGroupIPVersionParamsTLSVersion = "TLSv1_1"
	HTTPTimeseriesGroupIPVersionParamsTLSVersionTlSv1_2  HTTPTimeseriesGroupIPVersionParamsTLSVersion = "TLSv1_2"
	HTTPTimeseriesGroupIPVersionParamsTLSVersionTlSv1_3  HTTPTimeseriesGroupIPVersionParamsTLSVersion = "TLSv1_3"
	HTTPTimeseriesGroupIPVersionParamsTLSVersionTlSvQuic HTTPTimeseriesGroupIPVersionParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPTimeseriesGroupIPVersionParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupIPVersionParamsTLSVersionTlSv1_0, HTTPTimeseriesGroupIPVersionParamsTLSVersionTlSv1_1, HTTPTimeseriesGroupIPVersionParamsTLSVersionTlSv1_2, HTTPTimeseriesGroupIPVersionParamsTLSVersionTlSv1_3, HTTPTimeseriesGroupIPVersionParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPTimeseriesGroupIPVersionResponseEnvelope struct {
	Result  HTTPTimeseriesGroupIPVersionResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    httpTimeseriesGroupIPVersionResponseEnvelopeJSON `json:"-"`
}

// httpTimeseriesGroupIPVersionResponseEnvelopeJSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupIPVersionResponseEnvelope]
type httpTimeseriesGroupIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupIPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupOSParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[HTTPTimeseriesGroupOSParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPTimeseriesGroupOSParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]HTTPTimeseriesGroupOSParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[HTTPTimeseriesGroupOSParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]HTTPTimeseriesGroupOSParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]HTTPTimeseriesGroupOSParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]HTTPTimeseriesGroupOSParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for tls version.
	TLSVersion param.Field[[]HTTPTimeseriesGroupOSParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPTimeseriesGroupOSParams]'s query parameters as
// `url.Values`.
func (r HTTPTimeseriesGroupOSParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupOSParamsAggInterval string

const (
	HTTPTimeseriesGroupOSParamsAggInterval15m HTTPTimeseriesGroupOSParamsAggInterval = "15m"
	HTTPTimeseriesGroupOSParamsAggInterval1h  HTTPTimeseriesGroupOSParamsAggInterval = "1h"
	HTTPTimeseriesGroupOSParamsAggInterval1d  HTTPTimeseriesGroupOSParamsAggInterval = "1d"
	HTTPTimeseriesGroupOSParamsAggInterval1w  HTTPTimeseriesGroupOSParamsAggInterval = "1w"
)

func (r HTTPTimeseriesGroupOSParamsAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupOSParamsAggInterval15m, HTTPTimeseriesGroupOSParamsAggInterval1h, HTTPTimeseriesGroupOSParamsAggInterval1d, HTTPTimeseriesGroupOSParamsAggInterval1w:
		return true
	}
	return false
}

type HTTPTimeseriesGroupOSParamsBotClass string

const (
	HTTPTimeseriesGroupOSParamsBotClassLikelyAutomated HTTPTimeseriesGroupOSParamsBotClass = "LIKELY_AUTOMATED"
	HTTPTimeseriesGroupOSParamsBotClassLikelyHuman     HTTPTimeseriesGroupOSParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPTimeseriesGroupOSParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupOSParamsBotClassLikelyAutomated, HTTPTimeseriesGroupOSParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPTimeseriesGroupOSParamsDeviceType string

const (
	HTTPTimeseriesGroupOSParamsDeviceTypeDesktop HTTPTimeseriesGroupOSParamsDeviceType = "DESKTOP"
	HTTPTimeseriesGroupOSParamsDeviceTypeMobile  HTTPTimeseriesGroupOSParamsDeviceType = "MOBILE"
	HTTPTimeseriesGroupOSParamsDeviceTypeOther   HTTPTimeseriesGroupOSParamsDeviceType = "OTHER"
)

func (r HTTPTimeseriesGroupOSParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupOSParamsDeviceTypeDesktop, HTTPTimeseriesGroupOSParamsDeviceTypeMobile, HTTPTimeseriesGroupOSParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format results are returned in.
type HTTPTimeseriesGroupOSParamsFormat string

const (
	HTTPTimeseriesGroupOSParamsFormatJson HTTPTimeseriesGroupOSParamsFormat = "JSON"
	HTTPTimeseriesGroupOSParamsFormatCsv  HTTPTimeseriesGroupOSParamsFormat = "CSV"
)

func (r HTTPTimeseriesGroupOSParamsFormat) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupOSParamsFormatJson, HTTPTimeseriesGroupOSParamsFormatCsv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupOSParamsHTTPProtocol string

const (
	HTTPTimeseriesGroupOSParamsHTTPProtocolHTTP  HTTPTimeseriesGroupOSParamsHTTPProtocol = "HTTP"
	HTTPTimeseriesGroupOSParamsHTTPProtocolHTTPS HTTPTimeseriesGroupOSParamsHTTPProtocol = "HTTPS"
)

func (r HTTPTimeseriesGroupOSParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupOSParamsHTTPProtocolHTTP, HTTPTimeseriesGroupOSParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPTimeseriesGroupOSParamsHTTPVersion string

const (
	HTTPTimeseriesGroupOSParamsHTTPVersionHttPv1 HTTPTimeseriesGroupOSParamsHTTPVersion = "HTTPv1"
	HTTPTimeseriesGroupOSParamsHTTPVersionHttPv2 HTTPTimeseriesGroupOSParamsHTTPVersion = "HTTPv2"
	HTTPTimeseriesGroupOSParamsHTTPVersionHttPv3 HTTPTimeseriesGroupOSParamsHTTPVersion = "HTTPv3"
)

func (r HTTPTimeseriesGroupOSParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupOSParamsHTTPVersionHttPv1, HTTPTimeseriesGroupOSParamsHTTPVersionHttPv2, HTTPTimeseriesGroupOSParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPTimeseriesGroupOSParamsIPVersion string

const (
	HTTPTimeseriesGroupOSParamsIPVersionIPv4 HTTPTimeseriesGroupOSParamsIPVersion = "IPv4"
	HTTPTimeseriesGroupOSParamsIPVersionIPv6 HTTPTimeseriesGroupOSParamsIPVersion = "IPv6"
)

func (r HTTPTimeseriesGroupOSParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupOSParamsIPVersionIPv4, HTTPTimeseriesGroupOSParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPTimeseriesGroupOSParamsTLSVersion string

const (
	HTTPTimeseriesGroupOSParamsTLSVersionTlSv1_0  HTTPTimeseriesGroupOSParamsTLSVersion = "TLSv1_0"
	HTTPTimeseriesGroupOSParamsTLSVersionTlSv1_1  HTTPTimeseriesGroupOSParamsTLSVersion = "TLSv1_1"
	HTTPTimeseriesGroupOSParamsTLSVersionTlSv1_2  HTTPTimeseriesGroupOSParamsTLSVersion = "TLSv1_2"
	HTTPTimeseriesGroupOSParamsTLSVersionTlSv1_3  HTTPTimeseriesGroupOSParamsTLSVersion = "TLSv1_3"
	HTTPTimeseriesGroupOSParamsTLSVersionTlSvQuic HTTPTimeseriesGroupOSParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPTimeseriesGroupOSParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupOSParamsTLSVersionTlSv1_0, HTTPTimeseriesGroupOSParamsTLSVersionTlSv1_1, HTTPTimeseriesGroupOSParamsTLSVersionTlSv1_2, HTTPTimeseriesGroupOSParamsTLSVersionTlSv1_3, HTTPTimeseriesGroupOSParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPTimeseriesGroupOSResponseEnvelope struct {
	Result  HTTPTimeseriesGroupOSResponse             `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    httpTimeseriesGroupOSResponseEnvelopeJSON `json:"-"`
}

// httpTimeseriesGroupOSResponseEnvelopeJSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupOSResponseEnvelope]
type httpTimeseriesGroupOSResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupOSResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupOSResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupPostQuantumParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[HTTPTimeseriesGroupPostQuantumParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPTimeseriesGroupPostQuantumParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]HTTPTimeseriesGroupPostQuantumParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[HTTPTimeseriesGroupPostQuantumParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]HTTPTimeseriesGroupPostQuantumParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]HTTPTimeseriesGroupPostQuantumParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]HTTPTimeseriesGroupPostQuantumParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]HTTPTimeseriesGroupPostQuantumParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]HTTPTimeseriesGroupPostQuantumParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPTimeseriesGroupPostQuantumParams]'s query parameters as
// `url.Values`.
func (r HTTPTimeseriesGroupPostQuantumParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupPostQuantumParamsAggInterval string

const (
	HTTPTimeseriesGroupPostQuantumParamsAggInterval15m HTTPTimeseriesGroupPostQuantumParamsAggInterval = "15m"
	HTTPTimeseriesGroupPostQuantumParamsAggInterval1h  HTTPTimeseriesGroupPostQuantumParamsAggInterval = "1h"
	HTTPTimeseriesGroupPostQuantumParamsAggInterval1d  HTTPTimeseriesGroupPostQuantumParamsAggInterval = "1d"
	HTTPTimeseriesGroupPostQuantumParamsAggInterval1w  HTTPTimeseriesGroupPostQuantumParamsAggInterval = "1w"
)

func (r HTTPTimeseriesGroupPostQuantumParamsAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupPostQuantumParamsAggInterval15m, HTTPTimeseriesGroupPostQuantumParamsAggInterval1h, HTTPTimeseriesGroupPostQuantumParamsAggInterval1d, HTTPTimeseriesGroupPostQuantumParamsAggInterval1w:
		return true
	}
	return false
}

type HTTPTimeseriesGroupPostQuantumParamsBotClass string

const (
	HTTPTimeseriesGroupPostQuantumParamsBotClassLikelyAutomated HTTPTimeseriesGroupPostQuantumParamsBotClass = "LIKELY_AUTOMATED"
	HTTPTimeseriesGroupPostQuantumParamsBotClassLikelyHuman     HTTPTimeseriesGroupPostQuantumParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPTimeseriesGroupPostQuantumParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupPostQuantumParamsBotClassLikelyAutomated, HTTPTimeseriesGroupPostQuantumParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPTimeseriesGroupPostQuantumParamsDeviceType string

const (
	HTTPTimeseriesGroupPostQuantumParamsDeviceTypeDesktop HTTPTimeseriesGroupPostQuantumParamsDeviceType = "DESKTOP"
	HTTPTimeseriesGroupPostQuantumParamsDeviceTypeMobile  HTTPTimeseriesGroupPostQuantumParamsDeviceType = "MOBILE"
	HTTPTimeseriesGroupPostQuantumParamsDeviceTypeOther   HTTPTimeseriesGroupPostQuantumParamsDeviceType = "OTHER"
)

func (r HTTPTimeseriesGroupPostQuantumParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupPostQuantumParamsDeviceTypeDesktop, HTTPTimeseriesGroupPostQuantumParamsDeviceTypeMobile, HTTPTimeseriesGroupPostQuantumParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format results are returned in.
type HTTPTimeseriesGroupPostQuantumParamsFormat string

const (
	HTTPTimeseriesGroupPostQuantumParamsFormatJson HTTPTimeseriesGroupPostQuantumParamsFormat = "JSON"
	HTTPTimeseriesGroupPostQuantumParamsFormatCsv  HTTPTimeseriesGroupPostQuantumParamsFormat = "CSV"
)

func (r HTTPTimeseriesGroupPostQuantumParamsFormat) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupPostQuantumParamsFormatJson, HTTPTimeseriesGroupPostQuantumParamsFormatCsv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupPostQuantumParamsHTTPProtocol string

const (
	HTTPTimeseriesGroupPostQuantumParamsHTTPProtocolHTTP  HTTPTimeseriesGroupPostQuantumParamsHTTPProtocol = "HTTP"
	HTTPTimeseriesGroupPostQuantumParamsHTTPProtocolHTTPS HTTPTimeseriesGroupPostQuantumParamsHTTPProtocol = "HTTPS"
)

func (r HTTPTimeseriesGroupPostQuantumParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupPostQuantumParamsHTTPProtocolHTTP, HTTPTimeseriesGroupPostQuantumParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPTimeseriesGroupPostQuantumParamsHTTPVersion string

const (
	HTTPTimeseriesGroupPostQuantumParamsHTTPVersionHttPv1 HTTPTimeseriesGroupPostQuantumParamsHTTPVersion = "HTTPv1"
	HTTPTimeseriesGroupPostQuantumParamsHTTPVersionHttPv2 HTTPTimeseriesGroupPostQuantumParamsHTTPVersion = "HTTPv2"
	HTTPTimeseriesGroupPostQuantumParamsHTTPVersionHttPv3 HTTPTimeseriesGroupPostQuantumParamsHTTPVersion = "HTTPv3"
)

func (r HTTPTimeseriesGroupPostQuantumParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupPostQuantumParamsHTTPVersionHttPv1, HTTPTimeseriesGroupPostQuantumParamsHTTPVersionHttPv2, HTTPTimeseriesGroupPostQuantumParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPTimeseriesGroupPostQuantumParamsIPVersion string

const (
	HTTPTimeseriesGroupPostQuantumParamsIPVersionIPv4 HTTPTimeseriesGroupPostQuantumParamsIPVersion = "IPv4"
	HTTPTimeseriesGroupPostQuantumParamsIPVersionIPv6 HTTPTimeseriesGroupPostQuantumParamsIPVersion = "IPv6"
)

func (r HTTPTimeseriesGroupPostQuantumParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupPostQuantumParamsIPVersionIPv4, HTTPTimeseriesGroupPostQuantumParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPTimeseriesGroupPostQuantumParamsOS string

const (
	HTTPTimeseriesGroupPostQuantumParamsOSWindows  HTTPTimeseriesGroupPostQuantumParamsOS = "WINDOWS"
	HTTPTimeseriesGroupPostQuantumParamsOSMacosx   HTTPTimeseriesGroupPostQuantumParamsOS = "MACOSX"
	HTTPTimeseriesGroupPostQuantumParamsOSIos      HTTPTimeseriesGroupPostQuantumParamsOS = "IOS"
	HTTPTimeseriesGroupPostQuantumParamsOSAndroid  HTTPTimeseriesGroupPostQuantumParamsOS = "ANDROID"
	HTTPTimeseriesGroupPostQuantumParamsOSChromeos HTTPTimeseriesGroupPostQuantumParamsOS = "CHROMEOS"
	HTTPTimeseriesGroupPostQuantumParamsOSLinux    HTTPTimeseriesGroupPostQuantumParamsOS = "LINUX"
	HTTPTimeseriesGroupPostQuantumParamsOSSmartTv  HTTPTimeseriesGroupPostQuantumParamsOS = "SMART_TV"
)

func (r HTTPTimeseriesGroupPostQuantumParamsOS) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupPostQuantumParamsOSWindows, HTTPTimeseriesGroupPostQuantumParamsOSMacosx, HTTPTimeseriesGroupPostQuantumParamsOSIos, HTTPTimeseriesGroupPostQuantumParamsOSAndroid, HTTPTimeseriesGroupPostQuantumParamsOSChromeos, HTTPTimeseriesGroupPostQuantumParamsOSLinux, HTTPTimeseriesGroupPostQuantumParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupPostQuantumParamsTLSVersion string

const (
	HTTPTimeseriesGroupPostQuantumParamsTLSVersionTlSv1_0  HTTPTimeseriesGroupPostQuantumParamsTLSVersion = "TLSv1_0"
	HTTPTimeseriesGroupPostQuantumParamsTLSVersionTlSv1_1  HTTPTimeseriesGroupPostQuantumParamsTLSVersion = "TLSv1_1"
	HTTPTimeseriesGroupPostQuantumParamsTLSVersionTlSv1_2  HTTPTimeseriesGroupPostQuantumParamsTLSVersion = "TLSv1_2"
	HTTPTimeseriesGroupPostQuantumParamsTLSVersionTlSv1_3  HTTPTimeseriesGroupPostQuantumParamsTLSVersion = "TLSv1_3"
	HTTPTimeseriesGroupPostQuantumParamsTLSVersionTlSvQuic HTTPTimeseriesGroupPostQuantumParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPTimeseriesGroupPostQuantumParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupPostQuantumParamsTLSVersionTlSv1_0, HTTPTimeseriesGroupPostQuantumParamsTLSVersionTlSv1_1, HTTPTimeseriesGroupPostQuantumParamsTLSVersionTlSv1_2, HTTPTimeseriesGroupPostQuantumParamsTLSVersionTlSv1_3, HTTPTimeseriesGroupPostQuantumParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPTimeseriesGroupPostQuantumResponseEnvelope struct {
	Result  HTTPTimeseriesGroupPostQuantumResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    httpTimeseriesGroupPostQuantumResponseEnvelopeJSON `json:"-"`
}

// httpTimeseriesGroupPostQuantumResponseEnvelopeJSON contains the JSON metadata
// for the struct [HTTPTimeseriesGroupPostQuantumResponseEnvelope]
type httpTimeseriesGroupPostQuantumResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupPostQuantumResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupPostQuantumResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupTLSVersionParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[HTTPTimeseriesGroupTLSVersionParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPTimeseriesGroupTLSVersionParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]HTTPTimeseriesGroupTLSVersionParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[HTTPTimeseriesGroupTLSVersionParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]HTTPTimeseriesGroupTLSVersionParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]HTTPTimeseriesGroupTLSVersionParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]HTTPTimeseriesGroupTLSVersionParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]HTTPTimeseriesGroupTLSVersionParamsOS] `query:"os"`
}

// URLQuery serializes [HTTPTimeseriesGroupTLSVersionParams]'s query parameters as
// `url.Values`.
func (r HTTPTimeseriesGroupTLSVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupTLSVersionParamsAggInterval string

const (
	HTTPTimeseriesGroupTLSVersionParamsAggInterval15m HTTPTimeseriesGroupTLSVersionParamsAggInterval = "15m"
	HTTPTimeseriesGroupTLSVersionParamsAggInterval1h  HTTPTimeseriesGroupTLSVersionParamsAggInterval = "1h"
	HTTPTimeseriesGroupTLSVersionParamsAggInterval1d  HTTPTimeseriesGroupTLSVersionParamsAggInterval = "1d"
	HTTPTimeseriesGroupTLSVersionParamsAggInterval1w  HTTPTimeseriesGroupTLSVersionParamsAggInterval = "1w"
)

func (r HTTPTimeseriesGroupTLSVersionParamsAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupTLSVersionParamsAggInterval15m, HTTPTimeseriesGroupTLSVersionParamsAggInterval1h, HTTPTimeseriesGroupTLSVersionParamsAggInterval1d, HTTPTimeseriesGroupTLSVersionParamsAggInterval1w:
		return true
	}
	return false
}

type HTTPTimeseriesGroupTLSVersionParamsBotClass string

const (
	HTTPTimeseriesGroupTLSVersionParamsBotClassLikelyAutomated HTTPTimeseriesGroupTLSVersionParamsBotClass = "LIKELY_AUTOMATED"
	HTTPTimeseriesGroupTLSVersionParamsBotClassLikelyHuman     HTTPTimeseriesGroupTLSVersionParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPTimeseriesGroupTLSVersionParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupTLSVersionParamsBotClassLikelyAutomated, HTTPTimeseriesGroupTLSVersionParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPTimeseriesGroupTLSVersionParamsDeviceType string

const (
	HTTPTimeseriesGroupTLSVersionParamsDeviceTypeDesktop HTTPTimeseriesGroupTLSVersionParamsDeviceType = "DESKTOP"
	HTTPTimeseriesGroupTLSVersionParamsDeviceTypeMobile  HTTPTimeseriesGroupTLSVersionParamsDeviceType = "MOBILE"
	HTTPTimeseriesGroupTLSVersionParamsDeviceTypeOther   HTTPTimeseriesGroupTLSVersionParamsDeviceType = "OTHER"
)

func (r HTTPTimeseriesGroupTLSVersionParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupTLSVersionParamsDeviceTypeDesktop, HTTPTimeseriesGroupTLSVersionParamsDeviceTypeMobile, HTTPTimeseriesGroupTLSVersionParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format results are returned in.
type HTTPTimeseriesGroupTLSVersionParamsFormat string

const (
	HTTPTimeseriesGroupTLSVersionParamsFormatJson HTTPTimeseriesGroupTLSVersionParamsFormat = "JSON"
	HTTPTimeseriesGroupTLSVersionParamsFormatCsv  HTTPTimeseriesGroupTLSVersionParamsFormat = "CSV"
)

func (r HTTPTimeseriesGroupTLSVersionParamsFormat) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupTLSVersionParamsFormatJson, HTTPTimeseriesGroupTLSVersionParamsFormatCsv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupTLSVersionParamsHTTPProtocol string

const (
	HTTPTimeseriesGroupTLSVersionParamsHTTPProtocolHTTP  HTTPTimeseriesGroupTLSVersionParamsHTTPProtocol = "HTTP"
	HTTPTimeseriesGroupTLSVersionParamsHTTPProtocolHTTPS HTTPTimeseriesGroupTLSVersionParamsHTTPProtocol = "HTTPS"
)

func (r HTTPTimeseriesGroupTLSVersionParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupTLSVersionParamsHTTPProtocolHTTP, HTTPTimeseriesGroupTLSVersionParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPTimeseriesGroupTLSVersionParamsHTTPVersion string

const (
	HTTPTimeseriesGroupTLSVersionParamsHTTPVersionHttPv1 HTTPTimeseriesGroupTLSVersionParamsHTTPVersion = "HTTPv1"
	HTTPTimeseriesGroupTLSVersionParamsHTTPVersionHttPv2 HTTPTimeseriesGroupTLSVersionParamsHTTPVersion = "HTTPv2"
	HTTPTimeseriesGroupTLSVersionParamsHTTPVersionHttPv3 HTTPTimeseriesGroupTLSVersionParamsHTTPVersion = "HTTPv3"
)

func (r HTTPTimeseriesGroupTLSVersionParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupTLSVersionParamsHTTPVersionHttPv1, HTTPTimeseriesGroupTLSVersionParamsHTTPVersionHttPv2, HTTPTimeseriesGroupTLSVersionParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPTimeseriesGroupTLSVersionParamsIPVersion string

const (
	HTTPTimeseriesGroupTLSVersionParamsIPVersionIPv4 HTTPTimeseriesGroupTLSVersionParamsIPVersion = "IPv4"
	HTTPTimeseriesGroupTLSVersionParamsIPVersionIPv6 HTTPTimeseriesGroupTLSVersionParamsIPVersion = "IPv6"
)

func (r HTTPTimeseriesGroupTLSVersionParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupTLSVersionParamsIPVersionIPv4, HTTPTimeseriesGroupTLSVersionParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPTimeseriesGroupTLSVersionParamsOS string

const (
	HTTPTimeseriesGroupTLSVersionParamsOSWindows  HTTPTimeseriesGroupTLSVersionParamsOS = "WINDOWS"
	HTTPTimeseriesGroupTLSVersionParamsOSMacosx   HTTPTimeseriesGroupTLSVersionParamsOS = "MACOSX"
	HTTPTimeseriesGroupTLSVersionParamsOSIos      HTTPTimeseriesGroupTLSVersionParamsOS = "IOS"
	HTTPTimeseriesGroupTLSVersionParamsOSAndroid  HTTPTimeseriesGroupTLSVersionParamsOS = "ANDROID"
	HTTPTimeseriesGroupTLSVersionParamsOSChromeos HTTPTimeseriesGroupTLSVersionParamsOS = "CHROMEOS"
	HTTPTimeseriesGroupTLSVersionParamsOSLinux    HTTPTimeseriesGroupTLSVersionParamsOS = "LINUX"
	HTTPTimeseriesGroupTLSVersionParamsOSSmartTv  HTTPTimeseriesGroupTLSVersionParamsOS = "SMART_TV"
)

func (r HTTPTimeseriesGroupTLSVersionParamsOS) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupTLSVersionParamsOSWindows, HTTPTimeseriesGroupTLSVersionParamsOSMacosx, HTTPTimeseriesGroupTLSVersionParamsOSIos, HTTPTimeseriesGroupTLSVersionParamsOSAndroid, HTTPTimeseriesGroupTLSVersionParamsOSChromeos, HTTPTimeseriesGroupTLSVersionParamsOSLinux, HTTPTimeseriesGroupTLSVersionParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupTLSVersionResponseEnvelope struct {
	Result  HTTPTimeseriesGroupTLSVersionResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    httpTimeseriesGroupTLSVersionResponseEnvelopeJSON `json:"-"`
}

// httpTimeseriesGroupTLSVersionResponseEnvelopeJSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupTLSVersionResponseEnvelope]
type httpTimeseriesGroupTLSVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupTLSVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupTLSVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
