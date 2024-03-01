// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZeroTrustDEXHTTPTestService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustDEXHTTPTestService]
// method instead.
type ZeroTrustDEXHTTPTestService struct {
	Options     []option.RequestOption
	Percentiles *ZeroTrustDEXHTTPTestPercentileService
}

// NewZeroTrustDEXHTTPTestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustDEXHTTPTestService(opts ...option.RequestOption) (r *ZeroTrustDEXHTTPTestService) {
	r = &ZeroTrustDEXHTTPTestService{}
	r.Options = opts
	r.Percentiles = NewZeroTrustDEXHTTPTestPercentileService(opts...)
	return
}

// Get test details and aggregate performance metrics for an http test for a given
// time period between 1 hour and 7 days.
func (r *ZeroTrustDEXHTTPTestService) Get(ctx context.Context, testID string, params ZeroTrustDEXHTTPTestGetParams, opts ...option.RequestOption) (res *ZeroTrustDexhttpTestGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDexhttpTestGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/http-tests/%s", params.AccountID, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustDexhttpTestGetResponse struct {
	// The url of the HTTP synthetic application test
	Host            string                                           `json:"host"`
	HTTPStats       ZeroTrustDexhttpTestGetResponseHTTPStats         `json:"httpStats,nullable"`
	HTTPStatsByColo []ZeroTrustDexhttpTestGetResponseHTTPStatsByColo `json:"httpStatsByColo"`
	// The interval at which the HTTP synthetic application test is set to run.
	Interval string                              `json:"interval"`
	Kind     ZeroTrustDexhttpTestGetResponseKind `json:"kind"`
	// The HTTP method to use when running the test
	Method string `json:"method"`
	// The name of the HTTP synthetic application test
	Name string                              `json:"name"`
	JSON zeroTrustDexhttpTestGetResponseJSON `json:"-"`
}

// zeroTrustDexhttpTestGetResponseJSON contains the JSON metadata for the struct
// [ZeroTrustDexhttpTestGetResponse]
type zeroTrustDexhttpTestGetResponseJSON struct {
	Host            apijson.Field
	HTTPStats       apijson.Field
	HTTPStatsByColo apijson.Field
	Interval        apijson.Field
	Kind            apijson.Field
	Method          apijson.Field
	Name            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustDexhttpTestGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDexhttpTestGetResponseHTTPStats struct {
	DNSResponseTimeMs    ZeroTrustDexhttpTestGetResponseHTTPStatsDNSResponseTimeMs    `json:"dnsResponseTimeMs,required"`
	HTTPStatusCode       []ZeroTrustDexhttpTestGetResponseHTTPStatsHTTPStatusCode     `json:"httpStatusCode,required"`
	ResourceFetchTimeMs  ZeroTrustDexhttpTestGetResponseHTTPStatsResourceFetchTimeMs  `json:"resourceFetchTimeMs,required"`
	ServerResponseTimeMs ZeroTrustDexhttpTestGetResponseHTTPStatsServerResponseTimeMs `json:"serverResponseTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                                        `json:"uniqueDevicesTotal,required"`
	JSON               zeroTrustDexhttpTestGetResponseHTTPStatsJSON `json:"-"`
}

// zeroTrustDexhttpTestGetResponseHTTPStatsJSON contains the JSON metadata for the
// struct [ZeroTrustDexhttpTestGetResponseHTTPStats]
type zeroTrustDexhttpTestGetResponseHTTPStatsJSON struct {
	DNSResponseTimeMs    apijson.Field
	HTTPStatusCode       apijson.Field
	ResourceFetchTimeMs  apijson.Field
	ServerResponseTimeMs apijson.Field
	UniqueDevicesTotal   apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustDexhttpTestGetResponseHTTPStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDexhttpTestGetResponseHTTPStatsDNSResponseTimeMs struct {
	Slots []ZeroTrustDexhttpTestGetResponseHTTPStatsDNSResponseTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                         `json:"min,nullable"`
	JSON zeroTrustDexhttpTestGetResponseHTTPStatsDNSResponseTimeMsJSON `json:"-"`
}

// zeroTrustDexhttpTestGetResponseHTTPStatsDNSResponseTimeMsJSON contains the JSON
// metadata for the struct
// [ZeroTrustDexhttpTestGetResponseHTTPStatsDNSResponseTimeMs]
type zeroTrustDexhttpTestGetResponseHTTPStatsDNSResponseTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDexhttpTestGetResponseHTTPStatsDNSResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDexhttpTestGetResponseHTTPStatsDNSResponseTimeMsSlot struct {
	Timestamp string                                                            `json:"timestamp,required"`
	Value     int64                                                             `json:"value,required"`
	JSON      zeroTrustDexhttpTestGetResponseHTTPStatsDNSResponseTimeMsSlotJSON `json:"-"`
}

// zeroTrustDexhttpTestGetResponseHTTPStatsDNSResponseTimeMsSlotJSON contains the
// JSON metadata for the struct
// [ZeroTrustDexhttpTestGetResponseHTTPStatsDNSResponseTimeMsSlot]
type zeroTrustDexhttpTestGetResponseHTTPStatsDNSResponseTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDexhttpTestGetResponseHTTPStatsDNSResponseTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDexhttpTestGetResponseHTTPStatsHTTPStatusCode struct {
	Status200 int64                                                      `json:"status200,required"`
	Status300 int64                                                      `json:"status300,required"`
	Status400 int64                                                      `json:"status400,required"`
	Status500 int64                                                      `json:"status500,required"`
	Timestamp string                                                     `json:"timestamp,required"`
	JSON      zeroTrustDexhttpTestGetResponseHTTPStatsHTTPStatusCodeJSON `json:"-"`
}

// zeroTrustDexhttpTestGetResponseHTTPStatsHTTPStatusCodeJSON contains the JSON
// metadata for the struct [ZeroTrustDexhttpTestGetResponseHTTPStatsHTTPStatusCode]
type zeroTrustDexhttpTestGetResponseHTTPStatsHTTPStatusCodeJSON struct {
	Status200   apijson.Field
	Status300   apijson.Field
	Status400   apijson.Field
	Status500   apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDexhttpTestGetResponseHTTPStatsHTTPStatusCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDexhttpTestGetResponseHTTPStatsResourceFetchTimeMs struct {
	Slots []ZeroTrustDexhttpTestGetResponseHTTPStatsResourceFetchTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                           `json:"min,nullable"`
	JSON zeroTrustDexhttpTestGetResponseHTTPStatsResourceFetchTimeMsJSON `json:"-"`
}

// zeroTrustDexhttpTestGetResponseHTTPStatsResourceFetchTimeMsJSON contains the
// JSON metadata for the struct
// [ZeroTrustDexhttpTestGetResponseHTTPStatsResourceFetchTimeMs]
type zeroTrustDexhttpTestGetResponseHTTPStatsResourceFetchTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDexhttpTestGetResponseHTTPStatsResourceFetchTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDexhttpTestGetResponseHTTPStatsResourceFetchTimeMsSlot struct {
	Timestamp string                                                              `json:"timestamp,required"`
	Value     int64                                                               `json:"value,required"`
	JSON      zeroTrustDexhttpTestGetResponseHTTPStatsResourceFetchTimeMsSlotJSON `json:"-"`
}

// zeroTrustDexhttpTestGetResponseHTTPStatsResourceFetchTimeMsSlotJSON contains the
// JSON metadata for the struct
// [ZeroTrustDexhttpTestGetResponseHTTPStatsResourceFetchTimeMsSlot]
type zeroTrustDexhttpTestGetResponseHTTPStatsResourceFetchTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDexhttpTestGetResponseHTTPStatsResourceFetchTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDexhttpTestGetResponseHTTPStatsServerResponseTimeMs struct {
	Slots []ZeroTrustDexhttpTestGetResponseHTTPStatsServerResponseTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                            `json:"min,nullable"`
	JSON zeroTrustDexhttpTestGetResponseHTTPStatsServerResponseTimeMsJSON `json:"-"`
}

// zeroTrustDexhttpTestGetResponseHTTPStatsServerResponseTimeMsJSON contains the
// JSON metadata for the struct
// [ZeroTrustDexhttpTestGetResponseHTTPStatsServerResponseTimeMs]
type zeroTrustDexhttpTestGetResponseHTTPStatsServerResponseTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDexhttpTestGetResponseHTTPStatsServerResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDexhttpTestGetResponseHTTPStatsServerResponseTimeMsSlot struct {
	Timestamp string                                                               `json:"timestamp,required"`
	Value     int64                                                                `json:"value,required"`
	JSON      zeroTrustDexhttpTestGetResponseHTTPStatsServerResponseTimeMsSlotJSON `json:"-"`
}

// zeroTrustDexhttpTestGetResponseHTTPStatsServerResponseTimeMsSlotJSON contains
// the JSON metadata for the struct
// [ZeroTrustDexhttpTestGetResponseHTTPStatsServerResponseTimeMsSlot]
type zeroTrustDexhttpTestGetResponseHTTPStatsServerResponseTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDexhttpTestGetResponseHTTPStatsServerResponseTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDexhttpTestGetResponseHTTPStatsByColo struct {
	Colo                 string                                                             `json:"colo,required"`
	DNSResponseTimeMs    ZeroTrustDexhttpTestGetResponseHTTPStatsByColoDNSResponseTimeMs    `json:"dnsResponseTimeMs,required"`
	HTTPStatusCode       []ZeroTrustDexhttpTestGetResponseHTTPStatsByColoHTTPStatusCode     `json:"httpStatusCode,required"`
	ResourceFetchTimeMs  ZeroTrustDexhttpTestGetResponseHTTPStatsByColoResourceFetchTimeMs  `json:"resourceFetchTimeMs,required"`
	ServerResponseTimeMs ZeroTrustDexhttpTestGetResponseHTTPStatsByColoServerResponseTimeMs `json:"serverResponseTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                                              `json:"uniqueDevicesTotal,required"`
	JSON               zeroTrustDexhttpTestGetResponseHTTPStatsByColoJSON `json:"-"`
}

// zeroTrustDexhttpTestGetResponseHTTPStatsByColoJSON contains the JSON metadata
// for the struct [ZeroTrustDexhttpTestGetResponseHTTPStatsByColo]
type zeroTrustDexhttpTestGetResponseHTTPStatsByColoJSON struct {
	Colo                 apijson.Field
	DNSResponseTimeMs    apijson.Field
	HTTPStatusCode       apijson.Field
	ResourceFetchTimeMs  apijson.Field
	ServerResponseTimeMs apijson.Field
	UniqueDevicesTotal   apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustDexhttpTestGetResponseHTTPStatsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDexhttpTestGetResponseHTTPStatsByColoDNSResponseTimeMs struct {
	Slots []ZeroTrustDexhttpTestGetResponseHTTPStatsByColoDNSResponseTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                               `json:"min,nullable"`
	JSON zeroTrustDexhttpTestGetResponseHTTPStatsByColoDNSResponseTimeMsJSON `json:"-"`
}

// zeroTrustDexhttpTestGetResponseHTTPStatsByColoDNSResponseTimeMsJSON contains the
// JSON metadata for the struct
// [ZeroTrustDexhttpTestGetResponseHTTPStatsByColoDNSResponseTimeMs]
type zeroTrustDexhttpTestGetResponseHTTPStatsByColoDNSResponseTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDexhttpTestGetResponseHTTPStatsByColoDNSResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDexhttpTestGetResponseHTTPStatsByColoDNSResponseTimeMsSlot struct {
	Timestamp string                                                                  `json:"timestamp,required"`
	Value     int64                                                                   `json:"value,required"`
	JSON      zeroTrustDexhttpTestGetResponseHTTPStatsByColoDNSResponseTimeMsSlotJSON `json:"-"`
}

// zeroTrustDexhttpTestGetResponseHTTPStatsByColoDNSResponseTimeMsSlotJSON contains
// the JSON metadata for the struct
// [ZeroTrustDexhttpTestGetResponseHTTPStatsByColoDNSResponseTimeMsSlot]
type zeroTrustDexhttpTestGetResponseHTTPStatsByColoDNSResponseTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDexhttpTestGetResponseHTTPStatsByColoDNSResponseTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDexhttpTestGetResponseHTTPStatsByColoHTTPStatusCode struct {
	Status200 int64                                                            `json:"status200,required"`
	Status300 int64                                                            `json:"status300,required"`
	Status400 int64                                                            `json:"status400,required"`
	Status500 int64                                                            `json:"status500,required"`
	Timestamp string                                                           `json:"timestamp,required"`
	JSON      zeroTrustDexhttpTestGetResponseHTTPStatsByColoHTTPStatusCodeJSON `json:"-"`
}

// zeroTrustDexhttpTestGetResponseHTTPStatsByColoHTTPStatusCodeJSON contains the
// JSON metadata for the struct
// [ZeroTrustDexhttpTestGetResponseHTTPStatsByColoHTTPStatusCode]
type zeroTrustDexhttpTestGetResponseHTTPStatsByColoHTTPStatusCodeJSON struct {
	Status200   apijson.Field
	Status300   apijson.Field
	Status400   apijson.Field
	Status500   apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDexhttpTestGetResponseHTTPStatsByColoHTTPStatusCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDexhttpTestGetResponseHTTPStatsByColoResourceFetchTimeMs struct {
	Slots []ZeroTrustDexhttpTestGetResponseHTTPStatsByColoResourceFetchTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                                 `json:"min,nullable"`
	JSON zeroTrustDexhttpTestGetResponseHTTPStatsByColoResourceFetchTimeMsJSON `json:"-"`
}

// zeroTrustDexhttpTestGetResponseHTTPStatsByColoResourceFetchTimeMsJSON contains
// the JSON metadata for the struct
// [ZeroTrustDexhttpTestGetResponseHTTPStatsByColoResourceFetchTimeMs]
type zeroTrustDexhttpTestGetResponseHTTPStatsByColoResourceFetchTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDexhttpTestGetResponseHTTPStatsByColoResourceFetchTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDexhttpTestGetResponseHTTPStatsByColoResourceFetchTimeMsSlot struct {
	Timestamp string                                                                    `json:"timestamp,required"`
	Value     int64                                                                     `json:"value,required"`
	JSON      zeroTrustDexhttpTestGetResponseHTTPStatsByColoResourceFetchTimeMsSlotJSON `json:"-"`
}

// zeroTrustDexhttpTestGetResponseHTTPStatsByColoResourceFetchTimeMsSlotJSON
// contains the JSON metadata for the struct
// [ZeroTrustDexhttpTestGetResponseHTTPStatsByColoResourceFetchTimeMsSlot]
type zeroTrustDexhttpTestGetResponseHTTPStatsByColoResourceFetchTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDexhttpTestGetResponseHTTPStatsByColoResourceFetchTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDexhttpTestGetResponseHTTPStatsByColoServerResponseTimeMs struct {
	Slots []ZeroTrustDexhttpTestGetResponseHTTPStatsByColoServerResponseTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                                  `json:"min,nullable"`
	JSON zeroTrustDexhttpTestGetResponseHTTPStatsByColoServerResponseTimeMsJSON `json:"-"`
}

// zeroTrustDexhttpTestGetResponseHTTPStatsByColoServerResponseTimeMsJSON contains
// the JSON metadata for the struct
// [ZeroTrustDexhttpTestGetResponseHTTPStatsByColoServerResponseTimeMs]
type zeroTrustDexhttpTestGetResponseHTTPStatsByColoServerResponseTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDexhttpTestGetResponseHTTPStatsByColoServerResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDexhttpTestGetResponseHTTPStatsByColoServerResponseTimeMsSlot struct {
	Timestamp string                                                                     `json:"timestamp,required"`
	Value     int64                                                                      `json:"value,required"`
	JSON      zeroTrustDexhttpTestGetResponseHTTPStatsByColoServerResponseTimeMsSlotJSON `json:"-"`
}

// zeroTrustDexhttpTestGetResponseHTTPStatsByColoServerResponseTimeMsSlotJSON
// contains the JSON metadata for the struct
// [ZeroTrustDexhttpTestGetResponseHTTPStatsByColoServerResponseTimeMsSlot]
type zeroTrustDexhttpTestGetResponseHTTPStatsByColoServerResponseTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDexhttpTestGetResponseHTTPStatsByColoServerResponseTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDexhttpTestGetResponseKind string

const (
	ZeroTrustDexhttpTestGetResponseKindHTTP ZeroTrustDexhttpTestGetResponseKind = "http"
)

type ZeroTrustDEXHTTPTestGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Time interval for aggregate time slots.
	Interval param.Field[ZeroTrustDexhttpTestGetParamsInterval] `query:"interval,required"`
	// End time for aggregate metrics in ISO ms
	TimeEnd param.Field[string] `query:"timeEnd,required"`
	// Start time for aggregate metrics in ISO ms
	TimeStart param.Field[string] `query:"timeStart,required"`
	// Optionally filter result stats to a Cloudflare colo. Cannot be used in
	// combination with deviceId param.
	Colo param.Field[string] `query:"colo"`
	// Optionally filter result stats to a specific device(s). Cannot be used in
	// combination with colo param.
	DeviceID param.Field[[]string] `query:"deviceId"`
}

// URLQuery serializes [ZeroTrustDEXHTTPTestGetParams]'s query parameters as
// `url.Values`.
func (r ZeroTrustDEXHTTPTestGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Time interval for aggregate time slots.
type ZeroTrustDexhttpTestGetParamsInterval string

const (
	ZeroTrustDexhttpTestGetParamsIntervalMinute ZeroTrustDexhttpTestGetParamsInterval = "minute"
	ZeroTrustDexhttpTestGetParamsIntervalHour   ZeroTrustDexhttpTestGetParamsInterval = "hour"
)

type ZeroTrustDexhttpTestGetResponseEnvelope struct {
	Errors   []ZeroTrustDexhttpTestGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDexhttpTestGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustDexhttpTestGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustDexhttpTestGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDexhttpTestGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDexhttpTestGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustDexhttpTestGetResponseEnvelope]
type zeroTrustDexhttpTestGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDexhttpTestGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDexhttpTestGetResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zeroTrustDexhttpTestGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDexhttpTestGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZeroTrustDexhttpTestGetResponseEnvelopeErrors]
type zeroTrustDexhttpTestGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDexhttpTestGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDexhttpTestGetResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zeroTrustDexhttpTestGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDexhttpTestGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustDexhttpTestGetResponseEnvelopeMessages]
type zeroTrustDexhttpTestGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDexhttpTestGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustDexhttpTestGetResponseEnvelopeSuccess bool

const (
	ZeroTrustDexhttpTestGetResponseEnvelopeSuccessTrue ZeroTrustDexhttpTestGetResponseEnvelopeSuccess = true
)
