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

// DexHTTPTestService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDexHTTPTestService] method
// instead.
type DexHTTPTestService struct {
	Options     []option.RequestOption
	Percentiles *DexHTTPTestPercentileService
}

// NewDexHTTPTestService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDexHTTPTestService(opts ...option.RequestOption) (r *DexHTTPTestService) {
	r = &DexHTTPTestService{}
	r.Options = opts
	r.Percentiles = NewDexHTTPTestPercentileService(opts...)
	return
}

// Get test details and aggregate performance metrics for an http test for a given
// time period between 1 hour and 7 days.
func (r *DexHTTPTestService) Get(ctx context.Context, accountID string, testID string, query DexHTTPTestGetParams, opts ...option.RequestOption) (res *DexHTTPTestGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DexHTTPTestGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/http-tests/%s", accountID, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DexHTTPTestGetResponse struct {
	// The url of the HTTP synthetic application test
	Host            string                                  `json:"host"`
	HTTPStats       DexHTTPTestGetResponseHTTPStats         `json:"httpStats,nullable"`
	HTTPStatsByColo []DexHTTPTestGetResponseHTTPStatsByColo `json:"httpStatsByColo"`
	// The interval at which the HTTP synthetic application test is set to run.
	Interval string                     `json:"interval"`
	Kind     DexHTTPTestGetResponseKind `json:"kind"`
	// The HTTP method to use when running the test
	Method string `json:"method"`
	// The name of the HTTP synthetic application test
	Name string                     `json:"name"`
	JSON dexHTTPTestGetResponseJSON `json:"-"`
}

// dexHTTPTestGetResponseJSON contains the JSON metadata for the struct
// [DexHTTPTestGetResponse]
type dexHTTPTestGetResponseJSON struct {
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

func (r *DexHTTPTestGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseHTTPStats struct {
	DNSResponseTimeMs    DexHTTPTestGetResponseHTTPStatsDNSResponseTimeMs    `json:"dnsResponseTimeMs,required"`
	HTTPStatusCode       []DexHTTPTestGetResponseHTTPStatsHTTPStatusCode     `json:"httpStatusCode,required"`
	ResourceFetchTimeMs  DexHTTPTestGetResponseHTTPStatsResourceFetchTimeMs  `json:"resourceFetchTimeMs,required"`
	ServerResponseTimeMs DexHTTPTestGetResponseHTTPStatsServerResponseTimeMs `json:"serverResponseTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                               `json:"uniqueDevicesTotal,required"`
	JSON               dexHTTPTestGetResponseHTTPStatsJSON `json:"-"`
}

// dexHTTPTestGetResponseHTTPStatsJSON contains the JSON metadata for the struct
// [DexHTTPTestGetResponseHTTPStats]
type dexHTTPTestGetResponseHTTPStatsJSON struct {
	DNSResponseTimeMs    apijson.Field
	HTTPStatusCode       apijson.Field
	ResourceFetchTimeMs  apijson.Field
	ServerResponseTimeMs apijson.Field
	UniqueDevicesTotal   apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseHTTPStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseHTTPStatsDNSResponseTimeMs struct {
	Slots []DexHTTPTestGetResponseHTTPStatsDNSResponseTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                `json:"min,nullable"`
	JSON dexHTTPTestGetResponseHTTPStatsDNSResponseTimeMsJSON `json:"-"`
}

// dexHTTPTestGetResponseHTTPStatsDNSResponseTimeMsJSON contains the JSON metadata
// for the struct [DexHTTPTestGetResponseHTTPStatsDNSResponseTimeMs]
type dexHTTPTestGetResponseHTTPStatsDNSResponseTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseHTTPStatsDNSResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseHTTPStatsDNSResponseTimeMsSlot struct {
	Timestamp string                                                   `json:"timestamp,required"`
	Value     int64                                                    `json:"value,required"`
	JSON      dexHTTPTestGetResponseHTTPStatsDNSResponseTimeMsSlotJSON `json:"-"`
}

// dexHTTPTestGetResponseHTTPStatsDNSResponseTimeMsSlotJSON contains the JSON
// metadata for the struct [DexHTTPTestGetResponseHTTPStatsDNSResponseTimeMsSlot]
type dexHTTPTestGetResponseHTTPStatsDNSResponseTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseHTTPStatsDNSResponseTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseHTTPStatsHTTPStatusCode struct {
	Status200 int64                                             `json:"status200,required"`
	Status300 int64                                             `json:"status300,required"`
	Status400 int64                                             `json:"status400,required"`
	Status500 int64                                             `json:"status500,required"`
	Timestamp string                                            `json:"timestamp,required"`
	JSON      dexHTTPTestGetResponseHTTPStatsHTTPStatusCodeJSON `json:"-"`
}

// dexHTTPTestGetResponseHTTPStatsHTTPStatusCodeJSON contains the JSON metadata for
// the struct [DexHTTPTestGetResponseHTTPStatsHTTPStatusCode]
type dexHTTPTestGetResponseHTTPStatsHTTPStatusCodeJSON struct {
	Status200   apijson.Field
	Status300   apijson.Field
	Status400   apijson.Field
	Status500   apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseHTTPStatsHTTPStatusCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseHTTPStatsResourceFetchTimeMs struct {
	Slots []DexHTTPTestGetResponseHTTPStatsResourceFetchTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                  `json:"min,nullable"`
	JSON dexHTTPTestGetResponseHTTPStatsResourceFetchTimeMsJSON `json:"-"`
}

// dexHTTPTestGetResponseHTTPStatsResourceFetchTimeMsJSON contains the JSON
// metadata for the struct [DexHTTPTestGetResponseHTTPStatsResourceFetchTimeMs]
type dexHTTPTestGetResponseHTTPStatsResourceFetchTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseHTTPStatsResourceFetchTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseHTTPStatsResourceFetchTimeMsSlot struct {
	Timestamp string                                                     `json:"timestamp,required"`
	Value     int64                                                      `json:"value,required"`
	JSON      dexHTTPTestGetResponseHTTPStatsResourceFetchTimeMsSlotJSON `json:"-"`
}

// dexHTTPTestGetResponseHTTPStatsResourceFetchTimeMsSlotJSON contains the JSON
// metadata for the struct [DexHTTPTestGetResponseHTTPStatsResourceFetchTimeMsSlot]
type dexHTTPTestGetResponseHTTPStatsResourceFetchTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseHTTPStatsResourceFetchTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseHTTPStatsServerResponseTimeMs struct {
	Slots []DexHTTPTestGetResponseHTTPStatsServerResponseTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                   `json:"min,nullable"`
	JSON dexHTTPTestGetResponseHTTPStatsServerResponseTimeMsJSON `json:"-"`
}

// dexHTTPTestGetResponseHTTPStatsServerResponseTimeMsJSON contains the JSON
// metadata for the struct [DexHTTPTestGetResponseHTTPStatsServerResponseTimeMs]
type dexHTTPTestGetResponseHTTPStatsServerResponseTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseHTTPStatsServerResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseHTTPStatsServerResponseTimeMsSlot struct {
	Timestamp string                                                      `json:"timestamp,required"`
	Value     int64                                                       `json:"value,required"`
	JSON      dexHTTPTestGetResponseHTTPStatsServerResponseTimeMsSlotJSON `json:"-"`
}

// dexHTTPTestGetResponseHTTPStatsServerResponseTimeMsSlotJSON contains the JSON
// metadata for the struct
// [DexHTTPTestGetResponseHTTPStatsServerResponseTimeMsSlot]
type dexHTTPTestGetResponseHTTPStatsServerResponseTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseHTTPStatsServerResponseTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseHTTPStatsByColo struct {
	Colo                 string                                                    `json:"colo,required"`
	DNSResponseTimeMs    DexHTTPTestGetResponseHTTPStatsByColoDNSResponseTimeMs    `json:"dnsResponseTimeMs,required"`
	HTTPStatusCode       []DexHTTPTestGetResponseHTTPStatsByColoHTTPStatusCode     `json:"httpStatusCode,required"`
	ResourceFetchTimeMs  DexHTTPTestGetResponseHTTPStatsByColoResourceFetchTimeMs  `json:"resourceFetchTimeMs,required"`
	ServerResponseTimeMs DexHTTPTestGetResponseHTTPStatsByColoServerResponseTimeMs `json:"serverResponseTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                                     `json:"uniqueDevicesTotal,required"`
	JSON               dexHTTPTestGetResponseHTTPStatsByColoJSON `json:"-"`
}

// dexHTTPTestGetResponseHTTPStatsByColoJSON contains the JSON metadata for the
// struct [DexHTTPTestGetResponseHTTPStatsByColo]
type dexHTTPTestGetResponseHTTPStatsByColoJSON struct {
	Colo                 apijson.Field
	DNSResponseTimeMs    apijson.Field
	HTTPStatusCode       apijson.Field
	ResourceFetchTimeMs  apijson.Field
	ServerResponseTimeMs apijson.Field
	UniqueDevicesTotal   apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseHTTPStatsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseHTTPStatsByColoDNSResponseTimeMs struct {
	Slots []DexHTTPTestGetResponseHTTPStatsByColoDNSResponseTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                      `json:"min,nullable"`
	JSON dexHTTPTestGetResponseHTTPStatsByColoDNSResponseTimeMsJSON `json:"-"`
}

// dexHTTPTestGetResponseHTTPStatsByColoDNSResponseTimeMsJSON contains the JSON
// metadata for the struct [DexHTTPTestGetResponseHTTPStatsByColoDNSResponseTimeMs]
type dexHTTPTestGetResponseHTTPStatsByColoDNSResponseTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseHTTPStatsByColoDNSResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseHTTPStatsByColoDNSResponseTimeMsSlot struct {
	Timestamp string                                                         `json:"timestamp,required"`
	Value     int64                                                          `json:"value,required"`
	JSON      dexHTTPTestGetResponseHTTPStatsByColoDNSResponseTimeMsSlotJSON `json:"-"`
}

// dexHTTPTestGetResponseHTTPStatsByColoDNSResponseTimeMsSlotJSON contains the JSON
// metadata for the struct
// [DexHTTPTestGetResponseHTTPStatsByColoDNSResponseTimeMsSlot]
type dexHTTPTestGetResponseHTTPStatsByColoDNSResponseTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseHTTPStatsByColoDNSResponseTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseHTTPStatsByColoHTTPStatusCode struct {
	Status200 int64                                                   `json:"status200,required"`
	Status300 int64                                                   `json:"status300,required"`
	Status400 int64                                                   `json:"status400,required"`
	Status500 int64                                                   `json:"status500,required"`
	Timestamp string                                                  `json:"timestamp,required"`
	JSON      dexHTTPTestGetResponseHTTPStatsByColoHTTPStatusCodeJSON `json:"-"`
}

// dexHTTPTestGetResponseHTTPStatsByColoHTTPStatusCodeJSON contains the JSON
// metadata for the struct [DexHTTPTestGetResponseHTTPStatsByColoHTTPStatusCode]
type dexHTTPTestGetResponseHTTPStatsByColoHTTPStatusCodeJSON struct {
	Status200   apijson.Field
	Status300   apijson.Field
	Status400   apijson.Field
	Status500   apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseHTTPStatsByColoHTTPStatusCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseHTTPStatsByColoResourceFetchTimeMs struct {
	Slots []DexHTTPTestGetResponseHTTPStatsByColoResourceFetchTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                        `json:"min,nullable"`
	JSON dexHTTPTestGetResponseHTTPStatsByColoResourceFetchTimeMsJSON `json:"-"`
}

// dexHTTPTestGetResponseHTTPStatsByColoResourceFetchTimeMsJSON contains the JSON
// metadata for the struct
// [DexHTTPTestGetResponseHTTPStatsByColoResourceFetchTimeMs]
type dexHTTPTestGetResponseHTTPStatsByColoResourceFetchTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseHTTPStatsByColoResourceFetchTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseHTTPStatsByColoResourceFetchTimeMsSlot struct {
	Timestamp string                                                           `json:"timestamp,required"`
	Value     int64                                                            `json:"value,required"`
	JSON      dexHTTPTestGetResponseHTTPStatsByColoResourceFetchTimeMsSlotJSON `json:"-"`
}

// dexHTTPTestGetResponseHTTPStatsByColoResourceFetchTimeMsSlotJSON contains the
// JSON metadata for the struct
// [DexHTTPTestGetResponseHTTPStatsByColoResourceFetchTimeMsSlot]
type dexHTTPTestGetResponseHTTPStatsByColoResourceFetchTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseHTTPStatsByColoResourceFetchTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseHTTPStatsByColoServerResponseTimeMs struct {
	Slots []DexHTTPTestGetResponseHTTPStatsByColoServerResponseTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                         `json:"min,nullable"`
	JSON dexHTTPTestGetResponseHTTPStatsByColoServerResponseTimeMsJSON `json:"-"`
}

// dexHTTPTestGetResponseHTTPStatsByColoServerResponseTimeMsJSON contains the JSON
// metadata for the struct
// [DexHTTPTestGetResponseHTTPStatsByColoServerResponseTimeMs]
type dexHTTPTestGetResponseHTTPStatsByColoServerResponseTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseHTTPStatsByColoServerResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseHTTPStatsByColoServerResponseTimeMsSlot struct {
	Timestamp string                                                            `json:"timestamp,required"`
	Value     int64                                                             `json:"value,required"`
	JSON      dexHTTPTestGetResponseHTTPStatsByColoServerResponseTimeMsSlotJSON `json:"-"`
}

// dexHTTPTestGetResponseHTTPStatsByColoServerResponseTimeMsSlotJSON contains the
// JSON metadata for the struct
// [DexHTTPTestGetResponseHTTPStatsByColoServerResponseTimeMsSlot]
type dexHTTPTestGetResponseHTTPStatsByColoServerResponseTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseHTTPStatsByColoServerResponseTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseKind string

const (
	DexHTTPTestGetResponseKindHTTP DexHTTPTestGetResponseKind = "http"
)

type DexHTTPTestGetParams struct {
	// Time interval for aggregate time slots.
	Interval param.Field[DexHTTPTestGetParamsInterval] `query:"interval,required"`
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

// URLQuery serializes [DexHTTPTestGetParams]'s query parameters as `url.Values`.
func (r DexHTTPTestGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Time interval for aggregate time slots.
type DexHTTPTestGetParamsInterval string

const (
	DexHTTPTestGetParamsIntervalMinute DexHTTPTestGetParamsInterval = "minute"
	DexHTTPTestGetParamsIntervalHour   DexHTTPTestGetParamsInterval = "hour"
)

type DexHTTPTestGetResponseEnvelope struct {
	Errors   []DexHTTPTestGetResponseEnvelopeErrors   `json:"errors"`
	Messages []DexHTTPTestGetResponseEnvelopeMessages `json:"messages"`
	Result   DexHTTPTestGetResponse                   `json:"result"`
	// Whether the API call was successful
	Success DexHTTPTestGetResponseEnvelopeSuccess `json:"success"`
	JSON    dexHTTPTestGetResponseEnvelopeJSON    `json:"-"`
}

// dexHTTPTestGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DexHTTPTestGetResponseEnvelope]
type dexHTTPTestGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    dexHTTPTestGetResponseEnvelopeErrorsJSON `json:"-"`
}

// dexHTTPTestGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DexHTTPTestGetResponseEnvelopeErrors]
type dexHTTPTestGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    dexHTTPTestGetResponseEnvelopeMessagesJSON `json:"-"`
}

// dexHTTPTestGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DexHTTPTestGetResponseEnvelopeMessages]
type dexHTTPTestGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DexHTTPTestGetResponseEnvelopeSuccess bool

const (
	DexHTTPTestGetResponseEnvelopeSuccessTrue DexHTTPTestGetResponseEnvelopeSuccess = true
)
