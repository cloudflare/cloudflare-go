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

// ZoneAnalyticsColoService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneAnalyticsColoService] method
// instead.
type ZoneAnalyticsColoService struct {
	Options []option.RequestOption
}

// NewZoneAnalyticsColoService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneAnalyticsColoService(opts ...option.RequestOption) (r *ZoneAnalyticsColoService) {
	r = &ZoneAnalyticsColoService{}
	r.Options = opts
	return
}

// This view provides a breakdown of analytics data by datacenter. Note: This is
// available to Enterprise customers only.
func (r *ZoneAnalyticsColoService) ZoneAnalyticsDeprecatedGetAnalyticsByCoLocations(ctx context.Context, zoneIdentifier string, query ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParams, opts ...option.RequestOption) (res *ColoResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/analytics/colos", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ColoResponse struct {
	Errors   []ColoResponseError   `json:"errors"`
	Messages []ColoResponseMessage `json:"messages"`
	// The exact parameters/timestamps the analytics service used to return data.
	Query ColoResponseQuery `json:"query"`
	// A breakdown of all dashboard analytics data by co-locations. This is limited to
	// Enterprise zones only.
	Result []ColoResponseResult `json:"result"`
	// Whether the API call was successful
	Success ColoResponseSuccess `json:"success"`
	JSON    coloResponseJSON    `json:"-"`
}

// coloResponseJSON contains the JSON metadata for the struct [ColoResponse]
type coloResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Query       apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ColoResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ColoResponseError struct {
	Code    int64                 `json:"code,required"`
	Message string                `json:"message,required"`
	JSON    coloResponseErrorJSON `json:"-"`
}

// coloResponseErrorJSON contains the JSON metadata for the struct
// [ColoResponseError]
type coloResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ColoResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ColoResponseMessage struct {
	Code    int64                   `json:"code,required"`
	Message string                  `json:"message,required"`
	JSON    coloResponseMessageJSON `json:"-"`
}

// coloResponseMessageJSON contains the JSON metadata for the struct
// [ColoResponseMessage]
type coloResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ColoResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The exact parameters/timestamps the analytics service used to return data.
type ColoResponseQuery struct {
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
	Since ColoResponseQuerySince `json:"since"`
	// The amount of time (in minutes) that each data point in the timeseries
	// represents. The granularity of the time-series returned (e.g. each bucket in the
	// time series representing 1-minute vs 1-day) is calculated by the API based on
	// the time-range provided to the API.
	TimeDelta int64 `json:"time_delta"`
	// The (exclusive) end of the requested time frame. This value can be a negative
	// integer representing the number of minutes in the past relative to time the
	// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
	// omitted, the time of the request is used.
	Until ColoResponseQueryUntil `json:"until"`
	JSON  coloResponseQueryJSON  `json:"-"`
}

// coloResponseQueryJSON contains the JSON metadata for the struct
// [ColoResponseQuery]
type coloResponseQueryJSON struct {
	Since       apijson.Field
	TimeDelta   apijson.Field
	Until       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ColoResponseQuery) UnmarshalJSON(data []byte) (err error) {
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
type ColoResponseQuerySince interface {
	ImplementsColoResponseQuerySince()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ColoResponseQuerySince)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}

// The (exclusive) end of the requested time frame. This value can be a negative
// integer representing the number of minutes in the past relative to time the
// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
// omitted, the time of the request is used.
//
// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type ColoResponseQueryUntil interface {
	ImplementsColoResponseQueryUntil()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ColoResponseQueryUntil)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}

type ColoResponseResult struct {
	// The airport code identifer for the co-location.
	ColoID string `json:"colo_id"`
	// Time deltas containing metadata about each bucket of time. The number of buckets
	// (resolution) is determined by the amount of time between the since and until
	// parameters.
	Timeseries []ColoResponseResultTimesery `json:"timeseries"`
	// Breakdown of totals by data type.
	Totals ColoResponseResultTotals `json:"totals"`
	JSON   coloResponseResultJSON   `json:"-"`
}

// coloResponseResultJSON contains the JSON metadata for the struct
// [ColoResponseResult]
type coloResponseResultJSON struct {
	ColoID      apijson.Field
	Timeseries  apijson.Field
	Totals      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ColoResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ColoResponseResultTimesery struct {
	// Breakdown of totals for bandwidth in the form of bytes.
	Bandwidth ColoResponseResultTimeseriesBandwidth `json:"bandwidth"`
	// Breakdown of totals for requests.
	Requests ColoResponseResultTimeseriesRequests `json:"requests"`
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
	Since ColoResponseResultTimeseriesSince `json:"since"`
	// Breakdown of totals for threats.
	Threats ColoResponseResultTimeseriesThreats `json:"threats"`
	// The (exclusive) end of the requested time frame. This value can be a negative
	// integer representing the number of minutes in the past relative to time the
	// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
	// omitted, the time of the request is used.
	Until ColoResponseResultTimeseriesUntil `json:"until"`
	JSON  coloResponseResultTimeseryJSON    `json:"-"`
}

// coloResponseResultTimeseryJSON contains the JSON metadata for the struct
// [ColoResponseResultTimesery]
type coloResponseResultTimeseryJSON struct {
	Bandwidth   apijson.Field
	Requests    apijson.Field
	Since       apijson.Field
	Threats     apijson.Field
	Until       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ColoResponseResultTimesery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of totals for bandwidth in the form of bytes.
type ColoResponseResultTimeseriesBandwidth struct {
	// The total number of bytes served within the time frame.
	All int64 `json:"all"`
	// The number of bytes that were cached (and served) by Cloudflare.
	Cached int64 `json:"cached"`
	// The number of bytes that were fetched and served from the origin server.
	Uncached int64                                     `json:"uncached"`
	JSON     coloResponseResultTimeseriesBandwidthJSON `json:"-"`
}

// coloResponseResultTimeseriesBandwidthJSON contains the JSON metadata for the
// struct [ColoResponseResultTimeseriesBandwidth]
type coloResponseResultTimeseriesBandwidthJSON struct {
	All         apijson.Field
	Cached      apijson.Field
	Uncached    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ColoResponseResultTimeseriesBandwidth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of totals for requests.
type ColoResponseResultTimeseriesRequests struct {
	// Total number of requests served.
	All int64 `json:"all"`
	// Total number of cached requests served.
	Cached int64 `json:"cached"`
	// Key/value pairs where the key is a two-digit country code and the value is the
	// number of requests served to that country.
	Country interface{} `json:"country"`
	// A variable list of key/value pairs where the key is a HTTP status code and the
	// value is the number of requests with that code served.
	HTTPStatus interface{} `json:"http_status"`
	// Total number of requests served from the origin.
	Uncached int64                                    `json:"uncached"`
	JSON     coloResponseResultTimeseriesRequestsJSON `json:"-"`
}

// coloResponseResultTimeseriesRequestsJSON contains the JSON metadata for the
// struct [ColoResponseResultTimeseriesRequests]
type coloResponseResultTimeseriesRequestsJSON struct {
	All         apijson.Field
	Cached      apijson.Field
	Country     apijson.Field
	HTTPStatus  apijson.Field
	Uncached    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ColoResponseResultTimeseriesRequests) UnmarshalJSON(data []byte) (err error) {
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
type ColoResponseResultTimeseriesSince interface {
	ImplementsColoResponseResultTimeseriesSince()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ColoResponseResultTimeseriesSince)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}

// Breakdown of totals for threats.
type ColoResponseResultTimeseriesThreats struct {
	// The total number of identifiable threats received over the time frame.
	All int64 `json:"all"`
	// A list of key/value pairs where the key is a two-digit country code and the
	// value is the number of malicious requests received from that country.
	Country interface{} `json:"country"`
	// The list of key/value pairs where the key is a threat category and the value is
	// the number of requests.
	Type interface{}                             `json:"type"`
	JSON coloResponseResultTimeseriesThreatsJSON `json:"-"`
}

// coloResponseResultTimeseriesThreatsJSON contains the JSON metadata for the
// struct [ColoResponseResultTimeseriesThreats]
type coloResponseResultTimeseriesThreatsJSON struct {
	All         apijson.Field
	Country     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ColoResponseResultTimeseriesThreats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The (exclusive) end of the requested time frame. This value can be a negative
// integer representing the number of minutes in the past relative to time the
// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
// omitted, the time of the request is used.
//
// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type ColoResponseResultTimeseriesUntil interface {
	ImplementsColoResponseResultTimeseriesUntil()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ColoResponseResultTimeseriesUntil)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}

// Breakdown of totals by data type.
type ColoResponseResultTotals struct {
	// Breakdown of totals for bandwidth in the form of bytes.
	Bandwidth ColoResponseResultTotalsBandwidth `json:"bandwidth"`
	// Breakdown of totals for requests.
	Requests ColoResponseResultTotalsRequests `json:"requests"`
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
	Since ColoResponseResultTotalsSince `json:"since"`
	// Breakdown of totals for threats.
	Threats ColoResponseResultTotalsThreats `json:"threats"`
	// The (exclusive) end of the requested time frame. This value can be a negative
	// integer representing the number of minutes in the past relative to time the
	// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
	// omitted, the time of the request is used.
	Until ColoResponseResultTotalsUntil `json:"until"`
	JSON  coloResponseResultTotalsJSON  `json:"-"`
}

// coloResponseResultTotalsJSON contains the JSON metadata for the struct
// [ColoResponseResultTotals]
type coloResponseResultTotalsJSON struct {
	Bandwidth   apijson.Field
	Requests    apijson.Field
	Since       apijson.Field
	Threats     apijson.Field
	Until       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ColoResponseResultTotals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of totals for bandwidth in the form of bytes.
type ColoResponseResultTotalsBandwidth struct {
	// The total number of bytes served within the time frame.
	All int64 `json:"all"`
	// The number of bytes that were cached (and served) by Cloudflare.
	Cached int64 `json:"cached"`
	// The number of bytes that were fetched and served from the origin server.
	Uncached int64                                 `json:"uncached"`
	JSON     coloResponseResultTotalsBandwidthJSON `json:"-"`
}

// coloResponseResultTotalsBandwidthJSON contains the JSON metadata for the struct
// [ColoResponseResultTotalsBandwidth]
type coloResponseResultTotalsBandwidthJSON struct {
	All         apijson.Field
	Cached      apijson.Field
	Uncached    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ColoResponseResultTotalsBandwidth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of totals for requests.
type ColoResponseResultTotalsRequests struct {
	// Total number of requests served.
	All int64 `json:"all"`
	// Total number of cached requests served.
	Cached int64 `json:"cached"`
	// Key/value pairs where the key is a two-digit country code and the value is the
	// number of requests served to that country.
	Country interface{} `json:"country"`
	// A variable list of key/value pairs where the key is a HTTP status code and the
	// value is the number of requests with that code served.
	HTTPStatus interface{} `json:"http_status"`
	// Total number of requests served from the origin.
	Uncached int64                                `json:"uncached"`
	JSON     coloResponseResultTotalsRequestsJSON `json:"-"`
}

// coloResponseResultTotalsRequestsJSON contains the JSON metadata for the struct
// [ColoResponseResultTotalsRequests]
type coloResponseResultTotalsRequestsJSON struct {
	All         apijson.Field
	Cached      apijson.Field
	Country     apijson.Field
	HTTPStatus  apijson.Field
	Uncached    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ColoResponseResultTotalsRequests) UnmarshalJSON(data []byte) (err error) {
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
type ColoResponseResultTotalsSince interface {
	ImplementsColoResponseResultTotalsSince()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ColoResponseResultTotalsSince)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}

// Breakdown of totals for threats.
type ColoResponseResultTotalsThreats struct {
	// The total number of identifiable threats received over the time frame.
	All int64 `json:"all"`
	// A list of key/value pairs where the key is a two-digit country code and the
	// value is the number of malicious requests received from that country.
	Country interface{} `json:"country"`
	// The list of key/value pairs where the key is a threat category and the value is
	// the number of requests.
	Type interface{}                         `json:"type"`
	JSON coloResponseResultTotalsThreatsJSON `json:"-"`
}

// coloResponseResultTotalsThreatsJSON contains the JSON metadata for the struct
// [ColoResponseResultTotalsThreats]
type coloResponseResultTotalsThreatsJSON struct {
	All         apijson.Field
	Country     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ColoResponseResultTotalsThreats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The (exclusive) end of the requested time frame. This value can be a negative
// integer representing the number of minutes in the past relative to time the
// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
// omitted, the time of the request is used.
//
// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type ColoResponseResultTotalsUntil interface {
	ImplementsColoResponseResultTotalsUntil()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ColoResponseResultTotalsUntil)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}

// Whether the API call was successful
type ColoResponseSuccess bool

const (
	ColoResponseSuccessTrue ColoResponseSuccess = true
)

type ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParams struct {
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
	Since param.Field[ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParamsSince] `query:"since"`
	// The (exclusive) end of the requested time frame. This value can be a negative
	// integer representing the number of minutes in the past relative to time the
	// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
	// omitted, the time of the request is used.
	Until param.Field[ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParamsUntil] `query:"until"`
}

// URLQuery serializes
// [ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParams]'s
// query parameters as `url.Values`.
func (r ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParams) URLQuery() (v url.Values) {
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
type ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParamsSince interface {
	ImplementsZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParamsSince()
}

// The (exclusive) end of the requested time frame. This value can be a negative
// integer representing the number of minutes in the past relative to time the
// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
// omitted, the time of the request is used.
//
// Satisfied by [shared.UnionString], [shared.UnionInt].
type ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParamsUntil interface {
	ImplementsZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParamsUntil()
}
