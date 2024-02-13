// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// AnalyticsDashboardService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAnalyticsDashboardService] method
// instead.
type AnalyticsDashboardService struct {
	Options []option.RequestOption
}

// NewAnalyticsDashboardService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAnalyticsDashboardService(opts ...option.RequestOption) (r *AnalyticsDashboardService) {
	r = &AnalyticsDashboardService{}
	r.Options = opts
	return
}

// The dashboard view provides both totals and timeseries data for the given zone
// and time period across the entire Cloudflare network.
func (r *AnalyticsDashboardService) ZoneAnalyticsDeprecatedGetDashboard(ctx context.Context, zoneIdentifier string, query AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardParams, opts ...option.RequestOption) (res *AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelope
	path := fmt.Sprintf("zones/%s/analytics/dashboard", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Totals and timeseries data.
type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponse struct {
	// Time deltas containing metadata about each bucket of time. The number of buckets
	// (resolution) is determined by the amount of time between the since and until
	// parameters.
	Timeseries []AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimesery `json:"timeseries"`
	// Breakdown of totals by data type.
	Totals AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotals `json:"totals"`
	JSON   analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseJSON   `json:"-"`
}

// analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseJSON contains the
// JSON metadata for the struct
// [AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponse]
type analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseJSON struct {
	Timeseries  apijson.Field
	Totals      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimesery struct {
	// Breakdown of totals for bandwidth in the form of bytes.
	Bandwidth AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesBandwidth `json:"bandwidth"`
	// Breakdown of totals for pageviews.
	Pageviews AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesPageviews `json:"pageviews"`
	// Breakdown of totals for requests.
	Requests AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesRequests `json:"requests"`
	// The (inclusive) beginning of the requested time frame. This value can be a
	// negative integer representing the number of minutes in the past relative to time
	// the request is made, or can be an absolute timestamp that conforms to RFC 3339.
	// At this point in time, it cannot exceed a time in the past greater than one
	// year.
	//
	// Ranges that the Cloudflare web application provides will provide the following
	// period length for each point:
	//
	// - Last 60 minutes (from -59 to -1): 1 minute resolution
	// - Last 7 hours (from -419 to -60): 15 minutes resolution
	// - Last 15 hours (from -899 to -420): 30 minutes resolution
	// - Last 72 hours (from -4320 to -900): 1 hour resolution
	// - Older than 3 days (-525600 to -4320): 1 day resolution.
	Since AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesSince `json:"since"`
	// Breakdown of totals for threats.
	Threats AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesThreats `json:"threats"`
	Uniques AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesUniques `json:"uniques"`
	// The (exclusive) end of the requested time frame. This value can be a negative
	// integer representing the number of minutes in the past relative to time the
	// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
	// omitted, the time of the request is used.
	Until AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesUntil `json:"until"`
	JSON  analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseryJSON    `json:"-"`
}

// analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseryJSON
// contains the JSON metadata for the struct
// [AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimesery]
type analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseryJSON struct {
	Bandwidth   apijson.Field
	Pageviews   apijson.Field
	Requests    apijson.Field
	Since       apijson.Field
	Threats     apijson.Field
	Uniques     apijson.Field
	Until       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimesery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of totals for bandwidth in the form of bytes.
type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesBandwidth struct {
	// The total number of bytes served within the time frame.
	All int64 `json:"all"`
	// The number of bytes that were cached (and served) by Cloudflare.
	Cached int64 `json:"cached"`
	// A variable list of key/value pairs where the key represents the type of content
	// served, and the value is the number in bytes served.
	ContentType interface{} `json:"content_type"`
	// A variable list of key/value pairs where the key is a two-digit country code and
	// the value is the number of bytes served to that country.
	Country interface{} `json:"country"`
	// A break down of bytes served over HTTPS.
	SSL AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesBandwidthSSL `json:"ssl"`
	// A breakdown of requests by their SSL protocol.
	SSLProtocols AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesBandwidthSSLProtocols `json:"ssl_protocols"`
	// The number of bytes that were fetched and served from the origin server.
	Uncached int64                                                                                `json:"uncached"`
	JSON     analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesBandwidthJSON `json:"-"`
}

// analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesBandwidthJSON
// contains the JSON metadata for the struct
// [AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesBandwidth]
type analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesBandwidthJSON struct {
	All          apijson.Field
	Cached       apijson.Field
	ContentType  apijson.Field
	Country      apijson.Field
	SSL          apijson.Field
	SSLProtocols apijson.Field
	Uncached     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesBandwidth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A break down of bytes served over HTTPS.
type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesBandwidthSSL struct {
	// The number of bytes served over HTTPS.
	Encrypted int64 `json:"encrypted"`
	// The number of bytes served over HTTP.
	Unencrypted int64                                                                                   `json:"unencrypted"`
	JSON        analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesBandwidthSSLJSON `json:"-"`
}

// analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesBandwidthSSLJSON
// contains the JSON metadata for the struct
// [AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesBandwidthSSL]
type analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesBandwidthSSLJSON struct {
	Encrypted   apijson.Field
	Unencrypted apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesBandwidthSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A breakdown of requests by their SSL protocol.
type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesBandwidthSSLProtocols struct {
	// The number of requests served over HTTP.
	None int64 `json:"none"`
	// The number of requests served over TLS v1.0.
	TlSv1 int64 `json:"TLSv1"`
	// The number of requests served over TLS v1.1.
	TlSv1_1 int64 `json:"TLSv1.1"`
	// The number of requests served over TLS v1.2.
	TlSv1_2 int64 `json:"TLSv1.2"`
	// The number of requests served over TLS v1.3.
	TlSv1_3 int64                                                                                            `json:"TLSv1.3"`
	JSON    analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesBandwidthSSLProtocolsJSON `json:"-"`
}

// analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesBandwidthSSLProtocolsJSON
// contains the JSON metadata for the struct
// [AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesBandwidthSSLProtocols]
type analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesBandwidthSSLProtocolsJSON struct {
	None        apijson.Field
	TlSv1       apijson.Field
	TlSv1_1     apijson.Field
	TlSv1_2     apijson.Field
	TlSv1_3     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesBandwidthSSLProtocols) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of totals for pageviews.
type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesPageviews struct {
	// The total number of pageviews served within the time range.
	All int64 `json:"all"`
	// A variable list of key/value pairs representing the search engine and number of
	// hits.
	SearchEngine interface{}                                                                          `json:"search_engine"`
	JSON         analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesPageviewsJSON `json:"-"`
}

// analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesPageviewsJSON
// contains the JSON metadata for the struct
// [AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesPageviews]
type analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesPageviewsJSON struct {
	All          apijson.Field
	SearchEngine apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesPageviews) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of totals for requests.
type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesRequests struct {
	// Total number of requests served.
	All int64 `json:"all"`
	// Total number of cached requests served.
	Cached int64 `json:"cached"`
	// A variable list of key/value pairs where the key represents the type of content
	// served, and the value is the number of requests.
	ContentType interface{} `json:"content_type"`
	// A variable list of key/value pairs where the key is a two-digit country code and
	// the value is the number of requests served to that country.
	Country interface{} `json:"country"`
	// Key/value pairs where the key is a HTTP status code and the value is the number
	// of requests served with that code.
	HTTPStatus map[string]interface{} `json:"http_status"`
	// A break down of requests served over HTTPS.
	SSL AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesRequestsSSL `json:"ssl"`
	// A breakdown of requests by their SSL protocol.
	SSLProtocols AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesRequestsSSLProtocols `json:"ssl_protocols"`
	// Total number of requests served from the origin.
	Uncached int64                                                                               `json:"uncached"`
	JSON     analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesRequestsJSON `json:"-"`
}

// analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesRequestsJSON
// contains the JSON metadata for the struct
// [AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesRequests]
type analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesRequestsJSON struct {
	All          apijson.Field
	Cached       apijson.Field
	ContentType  apijson.Field
	Country      apijson.Field
	HTTPStatus   apijson.Field
	SSL          apijson.Field
	SSLProtocols apijson.Field
	Uncached     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesRequests) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A break down of requests served over HTTPS.
type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesRequestsSSL struct {
	// The number of requests served over HTTPS.
	Encrypted int64 `json:"encrypted"`
	// The number of requests served over HTTP.
	Unencrypted int64                                                                                  `json:"unencrypted"`
	JSON        analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesRequestsSSLJSON `json:"-"`
}

// analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesRequestsSSLJSON
// contains the JSON metadata for the struct
// [AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesRequestsSSL]
type analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesRequestsSSLJSON struct {
	Encrypted   apijson.Field
	Unencrypted apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesRequestsSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A breakdown of requests by their SSL protocol.
type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesRequestsSSLProtocols struct {
	// The number of requests served over HTTP.
	None int64 `json:"none"`
	// The number of requests served over TLS v1.0.
	TlSv1 int64 `json:"TLSv1"`
	// The number of requests served over TLS v1.1.
	TlSv1_1 int64 `json:"TLSv1.1"`
	// The number of requests served over TLS v1.2.
	TlSv1_2 int64 `json:"TLSv1.2"`
	// The number of requests served over TLS v1.3.
	TlSv1_3 int64                                                                                           `json:"TLSv1.3"`
	JSON    analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesRequestsSSLProtocolsJSON `json:"-"`
}

// analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesRequestsSSLProtocolsJSON
// contains the JSON metadata for the struct
// [AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesRequestsSSLProtocols]
type analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesRequestsSSLProtocolsJSON struct {
	None        apijson.Field
	TlSv1       apijson.Field
	TlSv1_1     apijson.Field
	TlSv1_2     apijson.Field
	TlSv1_3     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesRequestsSSLProtocols) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The (inclusive) beginning of the requested time frame. This value can be a
// negative integer representing the number of minutes in the past relative to time
// the request is made, or can be an absolute timestamp that conforms to RFC 3339.
// At this point in time, it cannot exceed a time in the past greater than one
// year.
//
// Ranges that the Cloudflare web application provides will provide the following
// period length for each point:
//
// - Last 60 minutes (from -59 to -1): 1 minute resolution
// - Last 7 hours (from -419 to -60): 15 minutes resolution
// - Last 15 hours (from -899 to -420): 30 minutes resolution
// - Last 72 hours (from -4320 to -900): 1 hour resolution
// - Older than 3 days (-525600 to -4320): 1 day resolution.
//
// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesSince interface {
	ImplementsAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesSince()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesSince)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}

// Breakdown of totals for threats.
type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesThreats struct {
	// The total number of identifiable threats received over the time frame.
	All int64 `json:"all"`
	// A list of key/value pairs where the key is a two-digit country code and the
	// value is the number of malicious requests received from that country.
	Country interface{} `json:"country"`
	// The list of key/value pairs where the key is a threat category and the value is
	// the number of requests.
	Type interface{}                                                                        `json:"type"`
	JSON analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesThreatsJSON `json:"-"`
}

// analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesThreatsJSON
// contains the JSON metadata for the struct
// [AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesThreats]
type analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesThreatsJSON struct {
	All         apijson.Field
	Country     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesThreats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesUniques struct {
	// Total number of unique IP addresses within the time range.
	All  int64                                                                              `json:"all"`
	JSON analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesUniquesJSON `json:"-"`
}

// analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesUniquesJSON
// contains the JSON metadata for the struct
// [AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesUniques]
type analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesUniquesJSON struct {
	All         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesUniques) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The (exclusive) end of the requested time frame. This value can be a negative
// integer representing the number of minutes in the past relative to time the
// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
// omitted, the time of the request is used.
//
// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesUntil interface {
	ImplementsAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesUntil()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTimeseriesUntil)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}

// Breakdown of totals by data type.
type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotals struct {
	// Breakdown of totals for bandwidth in the form of bytes.
	Bandwidth AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsBandwidth `json:"bandwidth"`
	// Breakdown of totals for pageviews.
	Pageviews AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsPageviews `json:"pageviews"`
	// Breakdown of totals for requests.
	Requests AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsRequests `json:"requests"`
	// The (inclusive) beginning of the requested time frame. This value can be a
	// negative integer representing the number of minutes in the past relative to time
	// the request is made, or can be an absolute timestamp that conforms to RFC 3339.
	// At this point in time, it cannot exceed a time in the past greater than one
	// year.
	//
	// Ranges that the Cloudflare web application provides will provide the following
	// period length for each point:
	//
	// - Last 60 minutes (from -59 to -1): 1 minute resolution
	// - Last 7 hours (from -419 to -60): 15 minutes resolution
	// - Last 15 hours (from -899 to -420): 30 minutes resolution
	// - Last 72 hours (from -4320 to -900): 1 hour resolution
	// - Older than 3 days (-525600 to -4320): 1 day resolution.
	Since AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsSince `json:"since"`
	// Breakdown of totals for threats.
	Threats AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsThreats `json:"threats"`
	Uniques AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsUniques `json:"uniques"`
	// The (exclusive) end of the requested time frame. This value can be a negative
	// integer representing the number of minutes in the past relative to time the
	// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
	// omitted, the time of the request is used.
	Until AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsUntil `json:"until"`
	JSON  analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsJSON  `json:"-"`
}

// analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsJSON contains
// the JSON metadata for the struct
// [AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotals]
type analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsJSON struct {
	Bandwidth   apijson.Field
	Pageviews   apijson.Field
	Requests    apijson.Field
	Since       apijson.Field
	Threats     apijson.Field
	Uniques     apijson.Field
	Until       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of totals for bandwidth in the form of bytes.
type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsBandwidth struct {
	// The total number of bytes served within the time frame.
	All int64 `json:"all"`
	// The number of bytes that were cached (and served) by Cloudflare.
	Cached int64 `json:"cached"`
	// A variable list of key/value pairs where the key represents the type of content
	// served, and the value is the number in bytes served.
	ContentType interface{} `json:"content_type"`
	// A variable list of key/value pairs where the key is a two-digit country code and
	// the value is the number of bytes served to that country.
	Country interface{} `json:"country"`
	// A break down of bytes served over HTTPS.
	SSL AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsBandwidthSSL `json:"ssl"`
	// A breakdown of requests by their SSL protocol.
	SSLProtocols AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsBandwidthSSLProtocols `json:"ssl_protocols"`
	// The number of bytes that were fetched and served from the origin server.
	Uncached int64                                                                            `json:"uncached"`
	JSON     analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsBandwidthJSON `json:"-"`
}

// analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsBandwidthJSON
// contains the JSON metadata for the struct
// [AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsBandwidth]
type analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsBandwidthJSON struct {
	All          apijson.Field
	Cached       apijson.Field
	ContentType  apijson.Field
	Country      apijson.Field
	SSL          apijson.Field
	SSLProtocols apijson.Field
	Uncached     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsBandwidth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A break down of bytes served over HTTPS.
type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsBandwidthSSL struct {
	// The number of bytes served over HTTPS.
	Encrypted int64 `json:"encrypted"`
	// The number of bytes served over HTTP.
	Unencrypted int64                                                                               `json:"unencrypted"`
	JSON        analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsBandwidthSSLJSON `json:"-"`
}

// analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsBandwidthSSLJSON
// contains the JSON metadata for the struct
// [AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsBandwidthSSL]
type analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsBandwidthSSLJSON struct {
	Encrypted   apijson.Field
	Unencrypted apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsBandwidthSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A breakdown of requests by their SSL protocol.
type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsBandwidthSSLProtocols struct {
	// The number of requests served over HTTP.
	None int64 `json:"none"`
	// The number of requests served over TLS v1.0.
	TlSv1 int64 `json:"TLSv1"`
	// The number of requests served over TLS v1.1.
	TlSv1_1 int64 `json:"TLSv1.1"`
	// The number of requests served over TLS v1.2.
	TlSv1_2 int64 `json:"TLSv1.2"`
	// The number of requests served over TLS v1.3.
	TlSv1_3 int64                                                                                        `json:"TLSv1.3"`
	JSON    analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsBandwidthSSLProtocolsJSON `json:"-"`
}

// analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsBandwidthSSLProtocolsJSON
// contains the JSON metadata for the struct
// [AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsBandwidthSSLProtocols]
type analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsBandwidthSSLProtocolsJSON struct {
	None        apijson.Field
	TlSv1       apijson.Field
	TlSv1_1     apijson.Field
	TlSv1_2     apijson.Field
	TlSv1_3     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsBandwidthSSLProtocols) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of totals for pageviews.
type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsPageviews struct {
	// The total number of pageviews served within the time range.
	All int64 `json:"all"`
	// A variable list of key/value pairs representing the search engine and number of
	// hits.
	SearchEngine interface{}                                                                      `json:"search_engine"`
	JSON         analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsPageviewsJSON `json:"-"`
}

// analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsPageviewsJSON
// contains the JSON metadata for the struct
// [AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsPageviews]
type analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsPageviewsJSON struct {
	All          apijson.Field
	SearchEngine apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsPageviews) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of totals for requests.
type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsRequests struct {
	// Total number of requests served.
	All int64 `json:"all"`
	// Total number of cached requests served.
	Cached int64 `json:"cached"`
	// A variable list of key/value pairs where the key represents the type of content
	// served, and the value is the number of requests.
	ContentType interface{} `json:"content_type"`
	// A variable list of key/value pairs where the key is a two-digit country code and
	// the value is the number of requests served to that country.
	Country interface{} `json:"country"`
	// Key/value pairs where the key is a HTTP status code and the value is the number
	// of requests served with that code.
	HTTPStatus map[string]interface{} `json:"http_status"`
	// A break down of requests served over HTTPS.
	SSL AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsRequestsSSL `json:"ssl"`
	// A breakdown of requests by their SSL protocol.
	SSLProtocols AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsRequestsSSLProtocols `json:"ssl_protocols"`
	// Total number of requests served from the origin.
	Uncached int64                                                                           `json:"uncached"`
	JSON     analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsRequestsJSON `json:"-"`
}

// analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsRequestsJSON
// contains the JSON metadata for the struct
// [AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsRequests]
type analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsRequestsJSON struct {
	All          apijson.Field
	Cached       apijson.Field
	ContentType  apijson.Field
	Country      apijson.Field
	HTTPStatus   apijson.Field
	SSL          apijson.Field
	SSLProtocols apijson.Field
	Uncached     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsRequests) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A break down of requests served over HTTPS.
type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsRequestsSSL struct {
	// The number of requests served over HTTPS.
	Encrypted int64 `json:"encrypted"`
	// The number of requests served over HTTP.
	Unencrypted int64                                                                              `json:"unencrypted"`
	JSON        analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsRequestsSSLJSON `json:"-"`
}

// analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsRequestsSSLJSON
// contains the JSON metadata for the struct
// [AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsRequestsSSL]
type analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsRequestsSSLJSON struct {
	Encrypted   apijson.Field
	Unencrypted apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsRequestsSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A breakdown of requests by their SSL protocol.
type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsRequestsSSLProtocols struct {
	// The number of requests served over HTTP.
	None int64 `json:"none"`
	// The number of requests served over TLS v1.0.
	TlSv1 int64 `json:"TLSv1"`
	// The number of requests served over TLS v1.1.
	TlSv1_1 int64 `json:"TLSv1.1"`
	// The number of requests served over TLS v1.2.
	TlSv1_2 int64 `json:"TLSv1.2"`
	// The number of requests served over TLS v1.3.
	TlSv1_3 int64                                                                                       `json:"TLSv1.3"`
	JSON    analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsRequestsSSLProtocolsJSON `json:"-"`
}

// analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsRequestsSSLProtocolsJSON
// contains the JSON metadata for the struct
// [AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsRequestsSSLProtocols]
type analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsRequestsSSLProtocolsJSON struct {
	None        apijson.Field
	TlSv1       apijson.Field
	TlSv1_1     apijson.Field
	TlSv1_2     apijson.Field
	TlSv1_3     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsRequestsSSLProtocols) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The (inclusive) beginning of the requested time frame. This value can be a
// negative integer representing the number of minutes in the past relative to time
// the request is made, or can be an absolute timestamp that conforms to RFC 3339.
// At this point in time, it cannot exceed a time in the past greater than one
// year.
//
// Ranges that the Cloudflare web application provides will provide the following
// period length for each point:
//
// - Last 60 minutes (from -59 to -1): 1 minute resolution
// - Last 7 hours (from -419 to -60): 15 minutes resolution
// - Last 15 hours (from -899 to -420): 30 minutes resolution
// - Last 72 hours (from -4320 to -900): 1 hour resolution
// - Older than 3 days (-525600 to -4320): 1 day resolution.
//
// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsSince interface {
	ImplementsAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsSince()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsSince)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}

// Breakdown of totals for threats.
type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsThreats struct {
	// The total number of identifiable threats received over the time frame.
	All int64 `json:"all"`
	// A list of key/value pairs where the key is a two-digit country code and the
	// value is the number of malicious requests received from that country.
	Country interface{} `json:"country"`
	// The list of key/value pairs where the key is a threat category and the value is
	// the number of requests.
	Type interface{}                                                                    `json:"type"`
	JSON analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsThreatsJSON `json:"-"`
}

// analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsThreatsJSON
// contains the JSON metadata for the struct
// [AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsThreats]
type analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsThreatsJSON struct {
	All         apijson.Field
	Country     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsThreats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsUniques struct {
	// Total number of unique IP addresses within the time range.
	All  int64                                                                          `json:"all"`
	JSON analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsUniquesJSON `json:"-"`
}

// analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsUniquesJSON
// contains the JSON metadata for the struct
// [AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsUniques]
type analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsUniquesJSON struct {
	All         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsUniques) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The (exclusive) end of the requested time frame. This value can be a negative
// integer representing the number of minutes in the past relative to time the
// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
// omitted, the time of the request is used.
//
// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsUntil interface {
	ImplementsAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsUntil()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseTotalsUntil)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}

type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardParams struct {
	// When set to true, the API will move the requested time window backward, until it
	// finds a region with completely aggregated data.
	//
	// The API response _may not represent the requested time window_.
	Continuous param.Field[bool] `query:"continuous"`
	// The (inclusive) beginning of the requested time frame. This value can be a
	// negative integer representing the number of minutes in the past relative to time
	// the request is made, or can be an absolute timestamp that conforms to RFC 3339.
	// At this point in time, it cannot exceed a time in the past greater than one
	// year.
	//
	// Ranges that the Cloudflare web application provides will provide the following
	// period length for each point:
	//
	// - Last 60 minutes (from -59 to -1): 1 minute resolution
	// - Last 7 hours (from -419 to -60): 15 minutes resolution
	// - Last 15 hours (from -899 to -420): 30 minutes resolution
	// - Last 72 hours (from -4320 to -900): 1 hour resolution
	// - Older than 3 days (-525600 to -4320): 1 day resolution.
	Since param.Field[AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardParamsSince] `query:"since"`
	// The (exclusive) end of the requested time frame. This value can be a negative
	// integer representing the number of minutes in the past relative to time the
	// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
	// omitted, the time of the request is used.
	Until param.Field[AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardParamsUntil] `query:"until"`
}

// URLQuery serializes
// [AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardParams]'s query parameters
// as `url.Values`.
func (r AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The (inclusive) beginning of the requested time frame. This value can be a
// negative integer representing the number of minutes in the past relative to time
// the request is made, or can be an absolute timestamp that conforms to RFC 3339.
// At this point in time, it cannot exceed a time in the past greater than one
// year.
//
// Ranges that the Cloudflare web application provides will provide the following
// period length for each point:
//
// - Last 60 minutes (from -59 to -1): 1 minute resolution
// - Last 7 hours (from -419 to -60): 15 minutes resolution
// - Last 15 hours (from -899 to -420): 30 minutes resolution
// - Last 72 hours (from -4320 to -900): 1 hour resolution
// - Older than 3 days (-525600 to -4320): 1 day resolution.
//
// Satisfied by [shared.UnionString], [shared.UnionInt].
type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardParamsSince interface {
	ImplementsAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardParamsSince()
}

// The (exclusive) end of the requested time frame. This value can be a negative
// integer representing the number of minutes in the past relative to time the
// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
// omitted, the time of the request is used.
//
// Satisfied by [shared.UnionString], [shared.UnionInt].
type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardParamsUntil interface {
	ImplementsAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardParamsUntil()
}

type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelope struct {
	Errors   []AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeMessages `json:"messages,required"`
	// Totals and timeseries data.
	Result AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponse `json:"result,required"`
	// Whether the API call was successful
	Success AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeSuccess `json:"success,required"`
	// The exact parameters/timestamps the analytics service used to return data.
	Query AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeQuery `json:"query"`
	JSON  analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeJSON  `json:"-"`
}

// analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelope]
type analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	Query       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeErrors struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeErrorsJSON `json:"-"`
}

// analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeErrors]
type analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeMessages struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeMessagesJSON `json:"-"`
}

// analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeMessages]
type analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeSuccess bool

const (
	AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeSuccessTrue AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeSuccess = true
)

// The exact parameters/timestamps the analytics service used to return data.
type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeQuery struct {
	// The (inclusive) beginning of the requested time frame. This value can be a
	// negative integer representing the number of minutes in the past relative to time
	// the request is made, or can be an absolute timestamp that conforms to RFC 3339.
	// At this point in time, it cannot exceed a time in the past greater than one
	// year.
	//
	// Ranges that the Cloudflare web application provides will provide the following
	// period length for each point:
	//
	// - Last 60 minutes (from -59 to -1): 1 minute resolution
	// - Last 7 hours (from -419 to -60): 15 minutes resolution
	// - Last 15 hours (from -899 to -420): 30 minutes resolution
	// - Last 72 hours (from -4320 to -900): 1 hour resolution
	// - Older than 3 days (-525600 to -4320): 1 day resolution.
	Since AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeQuerySince `json:"since"`
	// The amount of time (in minutes) that each data point in the timeseries
	// represents. The granularity of the time-series returned (e.g. each bucket in the
	// time series representing 1-minute vs 1-day) is calculated by the API based on
	// the time-range provided to the API.
	TimeDelta int64 `json:"time_delta"`
	// The (exclusive) end of the requested time frame. This value can be a negative
	// integer representing the number of minutes in the past relative to time the
	// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
	// omitted, the time of the request is used.
	Until AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeQueryUntil `json:"until"`
	JSON  analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeQueryJSON  `json:"-"`
}

// analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeQueryJSON
// contains the JSON metadata for the struct
// [AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeQuery]
type analyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeQueryJSON struct {
	Since       apijson.Field
	TimeDelta   apijson.Field
	Until       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The (inclusive) beginning of the requested time frame. This value can be a
// negative integer representing the number of minutes in the past relative to time
// the request is made, or can be an absolute timestamp that conforms to RFC 3339.
// At this point in time, it cannot exceed a time in the past greater than one
// year.
//
// Ranges that the Cloudflare web application provides will provide the following
// period length for each point:
//
// - Last 60 minutes (from -59 to -1): 1 minute resolution
// - Last 7 hours (from -419 to -60): 15 minutes resolution
// - Last 15 hours (from -899 to -420): 30 minutes resolution
// - Last 72 hours (from -4320 to -900): 1 hour resolution
// - Older than 3 days (-525600 to -4320): 1 day resolution.
//
// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeQuerySince interface {
	ImplementsAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeQuerySince()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeQuerySince)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}

// The (exclusive) end of the requested time frame. This value can be a negative
// integer representing the number of minutes in the past relative to time the
// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
// omitted, the time of the request is used.
//
// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeQueryUntil interface {
	ImplementsAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeQueryUntil()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseEnvelopeQueryUntil)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}
