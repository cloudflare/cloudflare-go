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

// AnalyticsColoService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAnalyticsColoService] method
// instead.
type AnalyticsColoService struct {
	Options []option.RequestOption
}

// NewAnalyticsColoService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAnalyticsColoService(opts ...option.RequestOption) (r *AnalyticsColoService) {
	r = &AnalyticsColoService{}
	r.Options = opts
	return
}

// This view provides a breakdown of analytics data by datacenter. Note: This is
// available to Enterprise customers only.
func (r *AnalyticsColoService) ZoneAnalyticsDeprecatedGetAnalyticsByCoLocations(ctx context.Context, zoneIdentifier string, query AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParams, opts ...option.RequestOption) (res *[]AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelope
	path := fmt.Sprintf("zones/%s/analytics/colos", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponse struct {
	// The airport code identifer for the co-location.
	ColoID string `json:"colo_id"`
	// Time deltas containing metadata about each bucket of time. The number of buckets
	// (resolution) is determined by the amount of time between the since and until
	// parameters.
	Timeseries []AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimesery `json:"timeseries"`
	// Breakdown of totals by data type.
	Totals AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotals `json:"totals"`
	JSON   analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseJSON   `json:"-"`
}

// analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseJSON
// contains the JSON metadata for the struct
// [AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponse]
type analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseJSON struct {
	ColoID      apijson.Field
	Timeseries  apijson.Field
	Totals      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimesery struct {
	// Breakdown of totals for bandwidth in the form of bytes.
	Bandwidth AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseriesBandwidth `json:"bandwidth"`
	// Breakdown of totals for requests.
	Requests AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseriesRequests `json:"requests"`
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
	Since AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseriesSince `json:"since"`
	// Breakdown of totals for threats.
	Threats AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseriesThreats `json:"threats"`
	// The (exclusive) end of the requested time frame. This value can be a negative
	// integer representing the number of minutes in the past relative to time the
	// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
	// omitted, the time of the request is used.
	Until AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseriesUntil `json:"until"`
	JSON  analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseryJSON    `json:"-"`
}

// analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseryJSON
// contains the JSON metadata for the struct
// [AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimesery]
type analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseryJSON struct {
	Bandwidth   apijson.Field
	Requests    apijson.Field
	Since       apijson.Field
	Threats     apijson.Field
	Until       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimesery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of totals for bandwidth in the form of bytes.
type AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseriesBandwidth struct {
	// The total number of bytes served within the time frame.
	All int64 `json:"all"`
	// The number of bytes that were cached (and served) by Cloudflare.
	Cached int64 `json:"cached"`
	// The number of bytes that were fetched and served from the origin server.
	Uncached int64                                                                                        `json:"uncached"`
	JSON     analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseriesBandwidthJSON `json:"-"`
}

// analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseriesBandwidthJSON
// contains the JSON metadata for the struct
// [AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseriesBandwidth]
type analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseriesBandwidthJSON struct {
	All         apijson.Field
	Cached      apijson.Field
	Uncached    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseriesBandwidth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of totals for requests.
type AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseriesRequests struct {
	// Total number of requests served.
	All int64 `json:"all"`
	// Total number of cached requests served.
	Cached int64 `json:"cached"`
	// Key/value pairs where the key is a two-digit country code and the value is the
	// number of requests served to that country.
	Country map[string]interface{} `json:"country"`
	// A variable list of key/value pairs where the key is a HTTP status code and the
	// value is the number of requests with that code served.
	HTTPStatus interface{} `json:"http_status"`
	// Total number of requests served from the origin.
	Uncached int64                                                                                       `json:"uncached"`
	JSON     analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseriesRequestsJSON `json:"-"`
}

// analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseriesRequestsJSON
// contains the JSON metadata for the struct
// [AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseriesRequests]
type analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseriesRequestsJSON struct {
	All         apijson.Field
	Cached      apijson.Field
	Country     apijson.Field
	HTTPStatus  apijson.Field
	Uncached    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseriesRequests) UnmarshalJSON(data []byte) (err error) {
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
type AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseriesSince interface {
	ImplementsAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseriesSince()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseriesSince)(nil)).Elem(),
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
type AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseriesThreats struct {
	// The total number of identifiable threats received over the time frame.
	All int64 `json:"all"`
	// A list of key/value pairs where the key is a two-digit country code and the
	// value is the number of malicious requests received from that country.
	Country interface{} `json:"country"`
	// The list of key/value pairs where the key is a threat category and the value is
	// the number of requests.
	Type interface{}                                                                                `json:"type"`
	JSON analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseriesThreatsJSON `json:"-"`
}

// analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseriesThreatsJSON
// contains the JSON metadata for the struct
// [AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseriesThreats]
type analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseriesThreatsJSON struct {
	All         apijson.Field
	Country     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseriesThreats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The (exclusive) end of the requested time frame. This value can be a negative
// integer representing the number of minutes in the past relative to time the
// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
// omitted, the time of the request is used.
//
// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseriesUntil interface {
	ImplementsAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseriesUntil()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTimeseriesUntil)(nil)).Elem(),
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
type AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotals struct {
	// Breakdown of totals for bandwidth in the form of bytes.
	Bandwidth AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsBandwidth `json:"bandwidth"`
	// Breakdown of totals for requests.
	Requests AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsRequests `json:"requests"`
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
	Since AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsSince `json:"since"`
	// Breakdown of totals for threats.
	Threats AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsThreats `json:"threats"`
	// The (exclusive) end of the requested time frame. This value can be a negative
	// integer representing the number of minutes in the past relative to time the
	// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
	// omitted, the time of the request is used.
	Until AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsUntil `json:"until"`
	JSON  analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsJSON  `json:"-"`
}

// analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsJSON
// contains the JSON metadata for the struct
// [AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotals]
type analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsJSON struct {
	Bandwidth   apijson.Field
	Requests    apijson.Field
	Since       apijson.Field
	Threats     apijson.Field
	Until       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of totals for bandwidth in the form of bytes.
type AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsBandwidth struct {
	// The total number of bytes served within the time frame.
	All int64 `json:"all"`
	// The number of bytes that were cached (and served) by Cloudflare.
	Cached int64 `json:"cached"`
	// The number of bytes that were fetched and served from the origin server.
	Uncached int64                                                                                    `json:"uncached"`
	JSON     analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsBandwidthJSON `json:"-"`
}

// analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsBandwidthJSON
// contains the JSON metadata for the struct
// [AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsBandwidth]
type analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsBandwidthJSON struct {
	All         apijson.Field
	Cached      apijson.Field
	Uncached    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsBandwidth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of totals for requests.
type AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsRequests struct {
	// Total number of requests served.
	All int64 `json:"all"`
	// Total number of cached requests served.
	Cached int64 `json:"cached"`
	// Key/value pairs where the key is a two-digit country code and the value is the
	// number of requests served to that country.
	Country map[string]interface{} `json:"country"`
	// A variable list of key/value pairs where the key is a HTTP status code and the
	// value is the number of requests with that code served.
	HTTPStatus interface{} `json:"http_status"`
	// Total number of requests served from the origin.
	Uncached int64                                                                                   `json:"uncached"`
	JSON     analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsRequestsJSON `json:"-"`
}

// analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsRequestsJSON
// contains the JSON metadata for the struct
// [AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsRequests]
type analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsRequestsJSON struct {
	All         apijson.Field
	Cached      apijson.Field
	Country     apijson.Field
	HTTPStatus  apijson.Field
	Uncached    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsRequests) UnmarshalJSON(data []byte) (err error) {
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
type AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsSince interface {
	ImplementsAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsSince()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsSince)(nil)).Elem(),
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
type AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsThreats struct {
	// The total number of identifiable threats received over the time frame.
	All int64 `json:"all"`
	// A list of key/value pairs where the key is a two-digit country code and the
	// value is the number of malicious requests received from that country.
	Country interface{} `json:"country"`
	// The list of key/value pairs where the key is a threat category and the value is
	// the number of requests.
	Type interface{}                                                                            `json:"type"`
	JSON analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsThreatsJSON `json:"-"`
}

// analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsThreatsJSON
// contains the JSON metadata for the struct
// [AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsThreats]
type analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsThreatsJSON struct {
	All         apijson.Field
	Country     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsThreats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The (exclusive) end of the requested time frame. This value can be a negative
// integer representing the number of minutes in the past relative to time the
// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
// omitted, the time of the request is used.
//
// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsUntil interface {
	ImplementsAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsUntil()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseTotalsUntil)(nil)).Elem(),
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

type AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParams struct {
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
	Since param.Field[AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParamsSince] `query:"since"`
	// The (exclusive) end of the requested time frame. This value can be a negative
	// integer representing the number of minutes in the past relative to time the
	// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
	// omitted, the time of the request is used.
	Until param.Field[AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParamsUntil] `query:"until"`
}

// URLQuery serializes
// [AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParams]'s query
// parameters as `url.Values`.
func (r AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParams) URLQuery() (v url.Values) {
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
type AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParamsSince interface {
	ImplementsAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParamsSince()
}

// The (exclusive) end of the requested time frame. This value can be a negative
// integer representing the number of minutes in the past relative to time the
// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
// omitted, the time of the request is used.
//
// Satisfied by [shared.UnionString], [shared.UnionInt].
type AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParamsUntil interface {
	ImplementsAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParamsUntil()
}

type AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelope struct {
	Errors   []AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeMessages `json:"messages,required"`
	// A breakdown of all dashboard analytics data by co-locations. This is limited to
	// Enterprise zones only.
	Result []AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponse `json:"result,required"`
	// Whether the API call was successful
	Success AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeSuccess `json:"success,required"`
	// The exact parameters/timestamps the analytics service used to return data.
	Query AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeQuery `json:"query"`
	JSON  analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeJSON  `json:"-"`
}

// analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelope]
type analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	Query       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeErrors struct {
	Code    int64                                                                                   `json:"code,required"`
	Message string                                                                                  `json:"message,required"`
	JSON    analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeErrorsJSON `json:"-"`
}

// analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeErrors]
type analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeMessages struct {
	Code    int64                                                                                     `json:"code,required"`
	Message string                                                                                    `json:"message,required"`
	JSON    analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeMessagesJSON `json:"-"`
}

// analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeMessages]
type analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeSuccess bool

const (
	AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeSuccessTrue AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeSuccess = true
)

// The exact parameters/timestamps the analytics service used to return data.
type AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeQuery struct {
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
	Since AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeQuerySince `json:"since"`
	// The amount of time (in minutes) that each data point in the timeseries
	// represents. The granularity of the time-series returned (e.g. each bucket in the
	// time series representing 1-minute vs 1-day) is calculated by the API based on
	// the time-range provided to the API.
	TimeDelta int64 `json:"time_delta"`
	// The (exclusive) end of the requested time frame. This value can be a negative
	// integer representing the number of minutes in the past relative to time the
	// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
	// omitted, the time of the request is used.
	Until AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeQueryUntil `json:"until"`
	JSON  analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeQueryJSON  `json:"-"`
}

// analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeQueryJSON
// contains the JSON metadata for the struct
// [AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeQuery]
type analyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeQueryJSON struct {
	Since       apijson.Field
	TimeDelta   apijson.Field
	Until       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeQuery) UnmarshalJSON(data []byte) (err error) {
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
type AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeQuerySince interface {
	ImplementsAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeQuerySince()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeQuerySince)(nil)).Elem(),
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
type AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeQueryUntil interface {
	ImplementsAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeQueryUntil()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseEnvelopeQueryUntil)(nil)).Elem(),
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
