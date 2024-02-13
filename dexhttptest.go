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

// DEXHTTPTestService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDEXHTTPTestService] method
// instead.
type DEXHTTPTestService struct {
	Options     []option.RequestOption
	Percentiles *DEXHTTPTestPercentileService
}

// NewDEXHTTPTestService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDEXHTTPTestService(opts ...option.RequestOption) (r *DEXHTTPTestService) {
	r = &DEXHTTPTestService{}
	r.Options = opts
	r.Percentiles = NewDEXHTTPTestPercentileService(opts...)
	return
}

// Get test details and aggregate performance metrics for an http test for a given
// time period between 1 hour and 7 days.
func (r *DEXHTTPTestService) Get(ctx context.Context, accountID string, testID string, query DEXHTTPTestGetParams, opts ...option.RequestOption) (res *DexhttpTestGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DexhttpTestGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/http-tests/%s", accountID, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DexhttpTestGetResponse struct {
	// The url of the HTTP synthetic application test
	Host            string                                  `json:"host"`
	HTTPStats       DexhttpTestGetResponseHTTPStats         `json:"httpStats,nullable"`
	HTTPStatsByColo []DexhttpTestGetResponseHTTPStatsByColo `json:"httpStatsByColo"`
	// The interval at which the HTTP synthetic application test is set to run.
	Interval string                     `json:"interval"`
	Kind     DexhttpTestGetResponseKind `json:"kind"`
	// The HTTP method to use when running the test
	Method string `json:"method"`
	// The name of the HTTP synthetic application test
	Name string                     `json:"name"`
	JSON dexhttpTestGetResponseJSON `json:"-"`
}

// dexhttpTestGetResponseJSON contains the JSON metadata for the struct
// [DexhttpTestGetResponse]
type dexhttpTestGetResponseJSON struct {
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

func (r *DexhttpTestGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexhttpTestGetResponseHTTPStats struct {
	DNSResponseTimeMs    DexhttpTestGetResponseHTTPStatsDNSResponseTimeMs    `json:"dnsResponseTimeMs,required"`
	HTTPStatusCode       []DexhttpTestGetResponseHTTPStatsHTTPStatusCode     `json:"httpStatusCode,required"`
	ResourceFetchTimeMs  DexhttpTestGetResponseHTTPStatsResourceFetchTimeMs  `json:"resourceFetchTimeMs,required"`
	ServerResponseTimeMs DexhttpTestGetResponseHTTPStatsServerResponseTimeMs `json:"serverResponseTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                               `json:"uniqueDevicesTotal,required"`
	JSON               dexhttpTestGetResponseHTTPStatsJSON `json:"-"`
}

// dexhttpTestGetResponseHTTPStatsJSON contains the JSON metadata for the struct
// [DexhttpTestGetResponseHTTPStats]
type dexhttpTestGetResponseHTTPStatsJSON struct {
	DNSResponseTimeMs    apijson.Field
	HTTPStatusCode       apijson.Field
	ResourceFetchTimeMs  apijson.Field
	ServerResponseTimeMs apijson.Field
	UniqueDevicesTotal   apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DexhttpTestGetResponseHTTPStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexhttpTestGetResponseHTTPStatsDNSResponseTimeMs struct {
	Slots []DexhttpTestGetResponseHTTPStatsDNSResponseTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                `json:"min,nullable"`
	JSON dexhttpTestGetResponseHTTPStatsDNSResponseTimeMsJSON `json:"-"`
}

// dexhttpTestGetResponseHTTPStatsDNSResponseTimeMsJSON contains the JSON metadata
// for the struct [DexhttpTestGetResponseHTTPStatsDNSResponseTimeMs]
type dexhttpTestGetResponseHTTPStatsDNSResponseTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexhttpTestGetResponseHTTPStatsDNSResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexhttpTestGetResponseHTTPStatsDNSResponseTimeMsSlot struct {
	Timestamp string                                                   `json:"timestamp,required"`
	Value     int64                                                    `json:"value,required"`
	JSON      dexhttpTestGetResponseHTTPStatsDNSResponseTimeMsSlotJSON `json:"-"`
}

// dexhttpTestGetResponseHTTPStatsDNSResponseTimeMsSlotJSON contains the JSON
// metadata for the struct [DexhttpTestGetResponseHTTPStatsDNSResponseTimeMsSlot]
type dexhttpTestGetResponseHTTPStatsDNSResponseTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexhttpTestGetResponseHTTPStatsDNSResponseTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexhttpTestGetResponseHTTPStatsHTTPStatusCode struct {
	Status200 int64                                             `json:"status200,required"`
	Status300 int64                                             `json:"status300,required"`
	Status400 int64                                             `json:"status400,required"`
	Status500 int64                                             `json:"status500,required"`
	Timestamp string                                            `json:"timestamp,required"`
	JSON      dexhttpTestGetResponseHTTPStatsHTTPStatusCodeJSON `json:"-"`
}

// dexhttpTestGetResponseHTTPStatsHTTPStatusCodeJSON contains the JSON metadata for
// the struct [DexhttpTestGetResponseHTTPStatsHTTPStatusCode]
type dexhttpTestGetResponseHTTPStatsHTTPStatusCodeJSON struct {
	Status200   apijson.Field
	Status300   apijson.Field
	Status400   apijson.Field
	Status500   apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexhttpTestGetResponseHTTPStatsHTTPStatusCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexhttpTestGetResponseHTTPStatsResourceFetchTimeMs struct {
	Slots []DexhttpTestGetResponseHTTPStatsResourceFetchTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                  `json:"min,nullable"`
	JSON dexhttpTestGetResponseHTTPStatsResourceFetchTimeMsJSON `json:"-"`
}

// dexhttpTestGetResponseHTTPStatsResourceFetchTimeMsJSON contains the JSON
// metadata for the struct [DexhttpTestGetResponseHTTPStatsResourceFetchTimeMs]
type dexhttpTestGetResponseHTTPStatsResourceFetchTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexhttpTestGetResponseHTTPStatsResourceFetchTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexhttpTestGetResponseHTTPStatsResourceFetchTimeMsSlot struct {
	Timestamp string                                                     `json:"timestamp,required"`
	Value     int64                                                      `json:"value,required"`
	JSON      dexhttpTestGetResponseHTTPStatsResourceFetchTimeMsSlotJSON `json:"-"`
}

// dexhttpTestGetResponseHTTPStatsResourceFetchTimeMsSlotJSON contains the JSON
// metadata for the struct [DexhttpTestGetResponseHTTPStatsResourceFetchTimeMsSlot]
type dexhttpTestGetResponseHTTPStatsResourceFetchTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexhttpTestGetResponseHTTPStatsResourceFetchTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexhttpTestGetResponseHTTPStatsServerResponseTimeMs struct {
	Slots []DexhttpTestGetResponseHTTPStatsServerResponseTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                   `json:"min,nullable"`
	JSON dexhttpTestGetResponseHTTPStatsServerResponseTimeMsJSON `json:"-"`
}

// dexhttpTestGetResponseHTTPStatsServerResponseTimeMsJSON contains the JSON
// metadata for the struct [DexhttpTestGetResponseHTTPStatsServerResponseTimeMs]
type dexhttpTestGetResponseHTTPStatsServerResponseTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexhttpTestGetResponseHTTPStatsServerResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexhttpTestGetResponseHTTPStatsServerResponseTimeMsSlot struct {
	Timestamp string                                                      `json:"timestamp,required"`
	Value     int64                                                       `json:"value,required"`
	JSON      dexhttpTestGetResponseHTTPStatsServerResponseTimeMsSlotJSON `json:"-"`
}

// dexhttpTestGetResponseHTTPStatsServerResponseTimeMsSlotJSON contains the JSON
// metadata for the struct
// [DexhttpTestGetResponseHTTPStatsServerResponseTimeMsSlot]
type dexhttpTestGetResponseHTTPStatsServerResponseTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexhttpTestGetResponseHTTPStatsServerResponseTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexhttpTestGetResponseHTTPStatsByColo struct {
	Colo                 string                                                    `json:"colo,required"`
	DNSResponseTimeMs    DexhttpTestGetResponseHTTPStatsByColoDNSResponseTimeMs    `json:"dnsResponseTimeMs,required"`
	HTTPStatusCode       []DexhttpTestGetResponseHTTPStatsByColoHTTPStatusCode     `json:"httpStatusCode,required"`
	ResourceFetchTimeMs  DexhttpTestGetResponseHTTPStatsByColoResourceFetchTimeMs  `json:"resourceFetchTimeMs,required"`
	ServerResponseTimeMs DexhttpTestGetResponseHTTPStatsByColoServerResponseTimeMs `json:"serverResponseTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                                     `json:"uniqueDevicesTotal,required"`
	JSON               dexhttpTestGetResponseHTTPStatsByColoJSON `json:"-"`
}

// dexhttpTestGetResponseHTTPStatsByColoJSON contains the JSON metadata for the
// struct [DexhttpTestGetResponseHTTPStatsByColo]
type dexhttpTestGetResponseHTTPStatsByColoJSON struct {
	Colo                 apijson.Field
	DNSResponseTimeMs    apijson.Field
	HTTPStatusCode       apijson.Field
	ResourceFetchTimeMs  apijson.Field
	ServerResponseTimeMs apijson.Field
	UniqueDevicesTotal   apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DexhttpTestGetResponseHTTPStatsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexhttpTestGetResponseHTTPStatsByColoDNSResponseTimeMs struct {
	Slots []DexhttpTestGetResponseHTTPStatsByColoDNSResponseTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                      `json:"min,nullable"`
	JSON dexhttpTestGetResponseHTTPStatsByColoDNSResponseTimeMsJSON `json:"-"`
}

// dexhttpTestGetResponseHTTPStatsByColoDNSResponseTimeMsJSON contains the JSON
// metadata for the struct [DexhttpTestGetResponseHTTPStatsByColoDNSResponseTimeMs]
type dexhttpTestGetResponseHTTPStatsByColoDNSResponseTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexhttpTestGetResponseHTTPStatsByColoDNSResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexhttpTestGetResponseHTTPStatsByColoDNSResponseTimeMsSlot struct {
	Timestamp string                                                         `json:"timestamp,required"`
	Value     int64                                                          `json:"value,required"`
	JSON      dexhttpTestGetResponseHTTPStatsByColoDNSResponseTimeMsSlotJSON `json:"-"`
}

// dexhttpTestGetResponseHTTPStatsByColoDNSResponseTimeMsSlotJSON contains the JSON
// metadata for the struct
// [DexhttpTestGetResponseHTTPStatsByColoDNSResponseTimeMsSlot]
type dexhttpTestGetResponseHTTPStatsByColoDNSResponseTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexhttpTestGetResponseHTTPStatsByColoDNSResponseTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexhttpTestGetResponseHTTPStatsByColoHTTPStatusCode struct {
	Status200 int64                                                   `json:"status200,required"`
	Status300 int64                                                   `json:"status300,required"`
	Status400 int64                                                   `json:"status400,required"`
	Status500 int64                                                   `json:"status500,required"`
	Timestamp string                                                  `json:"timestamp,required"`
	JSON      dexhttpTestGetResponseHTTPStatsByColoHTTPStatusCodeJSON `json:"-"`
}

// dexhttpTestGetResponseHTTPStatsByColoHTTPStatusCodeJSON contains the JSON
// metadata for the struct [DexhttpTestGetResponseHTTPStatsByColoHTTPStatusCode]
type dexhttpTestGetResponseHTTPStatsByColoHTTPStatusCodeJSON struct {
	Status200   apijson.Field
	Status300   apijson.Field
	Status400   apijson.Field
	Status500   apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexhttpTestGetResponseHTTPStatsByColoHTTPStatusCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexhttpTestGetResponseHTTPStatsByColoResourceFetchTimeMs struct {
	Slots []DexhttpTestGetResponseHTTPStatsByColoResourceFetchTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                        `json:"min,nullable"`
	JSON dexhttpTestGetResponseHTTPStatsByColoResourceFetchTimeMsJSON `json:"-"`
}

// dexhttpTestGetResponseHTTPStatsByColoResourceFetchTimeMsJSON contains the JSON
// metadata for the struct
// [DexhttpTestGetResponseHTTPStatsByColoResourceFetchTimeMs]
type dexhttpTestGetResponseHTTPStatsByColoResourceFetchTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexhttpTestGetResponseHTTPStatsByColoResourceFetchTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexhttpTestGetResponseHTTPStatsByColoResourceFetchTimeMsSlot struct {
	Timestamp string                                                           `json:"timestamp,required"`
	Value     int64                                                            `json:"value,required"`
	JSON      dexhttpTestGetResponseHTTPStatsByColoResourceFetchTimeMsSlotJSON `json:"-"`
}

// dexhttpTestGetResponseHTTPStatsByColoResourceFetchTimeMsSlotJSON contains the
// JSON metadata for the struct
// [DexhttpTestGetResponseHTTPStatsByColoResourceFetchTimeMsSlot]
type dexhttpTestGetResponseHTTPStatsByColoResourceFetchTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexhttpTestGetResponseHTTPStatsByColoResourceFetchTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexhttpTestGetResponseHTTPStatsByColoServerResponseTimeMs struct {
	Slots []DexhttpTestGetResponseHTTPStatsByColoServerResponseTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                         `json:"min,nullable"`
	JSON dexhttpTestGetResponseHTTPStatsByColoServerResponseTimeMsJSON `json:"-"`
}

// dexhttpTestGetResponseHTTPStatsByColoServerResponseTimeMsJSON contains the JSON
// metadata for the struct
// [DexhttpTestGetResponseHTTPStatsByColoServerResponseTimeMs]
type dexhttpTestGetResponseHTTPStatsByColoServerResponseTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexhttpTestGetResponseHTTPStatsByColoServerResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexhttpTestGetResponseHTTPStatsByColoServerResponseTimeMsSlot struct {
	Timestamp string                                                            `json:"timestamp,required"`
	Value     int64                                                             `json:"value,required"`
	JSON      dexhttpTestGetResponseHTTPStatsByColoServerResponseTimeMsSlotJSON `json:"-"`
}

// dexhttpTestGetResponseHTTPStatsByColoServerResponseTimeMsSlotJSON contains the
// JSON metadata for the struct
// [DexhttpTestGetResponseHTTPStatsByColoServerResponseTimeMsSlot]
type dexhttpTestGetResponseHTTPStatsByColoServerResponseTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexhttpTestGetResponseHTTPStatsByColoServerResponseTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexhttpTestGetResponseKind string

const (
	DexhttpTestGetResponseKindHTTP DexhttpTestGetResponseKind = "http"
)

type DEXHTTPTestGetParams struct {
	// Time interval for aggregate time slots.
	Interval param.Field[DexhttpTestGetParamsInterval] `query:"interval,required"`
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

// URLQuery serializes [DEXHTTPTestGetParams]'s query parameters as `url.Values`.
func (r DEXHTTPTestGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Time interval for aggregate time slots.
type DexhttpTestGetParamsInterval string

const (
	DexhttpTestGetParamsIntervalMinute DexhttpTestGetParamsInterval = "minute"
	DexhttpTestGetParamsIntervalHour   DexhttpTestGetParamsInterval = "hour"
)

type DexhttpTestGetResponseEnvelope struct {
	Errors   []DexhttpTestGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DexhttpTestGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DexhttpTestGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DexhttpTestGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dexhttpTestGetResponseEnvelopeJSON    `json:"-"`
}

// dexhttpTestGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DexhttpTestGetResponseEnvelope]
type dexhttpTestGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexhttpTestGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexhttpTestGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    dexhttpTestGetResponseEnvelopeErrorsJSON `json:"-"`
}

// dexhttpTestGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DexhttpTestGetResponseEnvelopeErrors]
type dexhttpTestGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexhttpTestGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexhttpTestGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    dexhttpTestGetResponseEnvelopeMessagesJSON `json:"-"`
}

// dexhttpTestGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DexhttpTestGetResponseEnvelopeMessages]
type dexhttpTestGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexhttpTestGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DexhttpTestGetResponseEnvelopeSuccess bool

const (
	DexhttpTestGetResponseEnvelopeSuccessTrue DexhttpTestGetResponseEnvelopeSuccess = true
)
