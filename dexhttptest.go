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
	path := fmt.Sprintf("accounts/%s/dex/http-tests/%s", accountID, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DexHTTPTestGetResponse struct {
	Errors   []DexHTTPTestGetResponseError   `json:"errors"`
	Messages []DexHTTPTestGetResponseMessage `json:"messages"`
	Result   DexHTTPTestGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success DexHTTPTestGetResponseSuccess `json:"success"`
	JSON    dexHTTPTestGetResponseJSON    `json:"-"`
}

// dexHTTPTestGetResponseJSON contains the JSON metadata for the struct
// [DexHTTPTestGetResponse]
type dexHTTPTestGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    dexHTTPTestGetResponseErrorJSON `json:"-"`
}

// dexHTTPTestGetResponseErrorJSON contains the JSON metadata for the struct
// [DexHTTPTestGetResponseError]
type dexHTTPTestGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    dexHTTPTestGetResponseMessageJSON `json:"-"`
}

// dexHTTPTestGetResponseMessageJSON contains the JSON metadata for the struct
// [DexHTTPTestGetResponseMessage]
type dexHTTPTestGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseResult struct {
	// The url of the HTTP synthetic application test
	Host            string                                        `json:"host"`
	HTTPStats       DexHTTPTestGetResponseResultHTTPStats         `json:"httpStats,nullable"`
	HTTPStatsByColo []DexHTTPTestGetResponseResultHTTPStatsByColo `json:"httpStatsByColo"`
	// The interval at which the HTTP synthetic application test is set to run.
	Interval string                           `json:"interval"`
	Kind     DexHTTPTestGetResponseResultKind `json:"kind"`
	// The HTTP method to use when running the test
	Method string `json:"method"`
	// The name of the HTTP synthetic application test
	Name string                           `json:"name"`
	JSON dexHTTPTestGetResponseResultJSON `json:"-"`
}

// dexHTTPTestGetResponseResultJSON contains the JSON metadata for the struct
// [DexHTTPTestGetResponseResult]
type dexHTTPTestGetResponseResultJSON struct {
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

func (r *DexHTTPTestGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseResultHTTPStats struct {
	DNSResponseTimeMs    DexHTTPTestGetResponseResultHTTPStatsDNSResponseTimeMs    `json:"dnsResponseTimeMs,required"`
	HTTPStatusCode       []DexHTTPTestGetResponseResultHTTPStatsHTTPStatusCode     `json:"httpStatusCode,required"`
	ResourceFetchTimeMs  DexHTTPTestGetResponseResultHTTPStatsResourceFetchTimeMs  `json:"resourceFetchTimeMs,required"`
	ServerResponseTimeMs DexHTTPTestGetResponseResultHTTPStatsServerResponseTimeMs `json:"serverResponseTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                                     `json:"uniqueDevicesTotal,required"`
	JSON               dexHTTPTestGetResponseResultHTTPStatsJSON `json:"-"`
}

// dexHTTPTestGetResponseResultHTTPStatsJSON contains the JSON metadata for the
// struct [DexHTTPTestGetResponseResultHTTPStats]
type dexHTTPTestGetResponseResultHTTPStatsJSON struct {
	DNSResponseTimeMs    apijson.Field
	HTTPStatusCode       apijson.Field
	ResourceFetchTimeMs  apijson.Field
	ServerResponseTimeMs apijson.Field
	UniqueDevicesTotal   apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseResultHTTPStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseResultHTTPStatsDNSResponseTimeMs struct {
	Slots []DexHTTPTestGetResponseResultHTTPStatsDNSResponseTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                      `json:"min,nullable"`
	JSON dexHTTPTestGetResponseResultHTTPStatsDNSResponseTimeMsJSON `json:"-"`
}

// dexHTTPTestGetResponseResultHTTPStatsDNSResponseTimeMsJSON contains the JSON
// metadata for the struct [DexHTTPTestGetResponseResultHTTPStatsDNSResponseTimeMs]
type dexHTTPTestGetResponseResultHTTPStatsDNSResponseTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseResultHTTPStatsDNSResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseResultHTTPStatsDNSResponseTimeMsSlot struct {
	Timestamp string                                                         `json:"timestamp,required"`
	Value     int64                                                          `json:"value,required"`
	JSON      dexHTTPTestGetResponseResultHTTPStatsDNSResponseTimeMsSlotJSON `json:"-"`
}

// dexHTTPTestGetResponseResultHTTPStatsDNSResponseTimeMsSlotJSON contains the JSON
// metadata for the struct
// [DexHTTPTestGetResponseResultHTTPStatsDNSResponseTimeMsSlot]
type dexHTTPTestGetResponseResultHTTPStatsDNSResponseTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseResultHTTPStatsDNSResponseTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseResultHTTPStatsHTTPStatusCode struct {
	Status200 int64                                                   `json:"status200,required"`
	Status300 int64                                                   `json:"status300,required"`
	Status400 int64                                                   `json:"status400,required"`
	Status500 int64                                                   `json:"status500,required"`
	Timestamp string                                                  `json:"timestamp,required"`
	JSON      dexHTTPTestGetResponseResultHTTPStatsHTTPStatusCodeJSON `json:"-"`
}

// dexHTTPTestGetResponseResultHTTPStatsHTTPStatusCodeJSON contains the JSON
// metadata for the struct [DexHTTPTestGetResponseResultHTTPStatsHTTPStatusCode]
type dexHTTPTestGetResponseResultHTTPStatsHTTPStatusCodeJSON struct {
	Status200   apijson.Field
	Status300   apijson.Field
	Status400   apijson.Field
	Status500   apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseResultHTTPStatsHTTPStatusCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseResultHTTPStatsResourceFetchTimeMs struct {
	Slots []DexHTTPTestGetResponseResultHTTPStatsResourceFetchTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                        `json:"min,nullable"`
	JSON dexHTTPTestGetResponseResultHTTPStatsResourceFetchTimeMsJSON `json:"-"`
}

// dexHTTPTestGetResponseResultHTTPStatsResourceFetchTimeMsJSON contains the JSON
// metadata for the struct
// [DexHTTPTestGetResponseResultHTTPStatsResourceFetchTimeMs]
type dexHTTPTestGetResponseResultHTTPStatsResourceFetchTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseResultHTTPStatsResourceFetchTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseResultHTTPStatsResourceFetchTimeMsSlot struct {
	Timestamp string                                                           `json:"timestamp,required"`
	Value     int64                                                            `json:"value,required"`
	JSON      dexHTTPTestGetResponseResultHTTPStatsResourceFetchTimeMsSlotJSON `json:"-"`
}

// dexHTTPTestGetResponseResultHTTPStatsResourceFetchTimeMsSlotJSON contains the
// JSON metadata for the struct
// [DexHTTPTestGetResponseResultHTTPStatsResourceFetchTimeMsSlot]
type dexHTTPTestGetResponseResultHTTPStatsResourceFetchTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseResultHTTPStatsResourceFetchTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseResultHTTPStatsServerResponseTimeMs struct {
	Slots []DexHTTPTestGetResponseResultHTTPStatsServerResponseTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                         `json:"min,nullable"`
	JSON dexHTTPTestGetResponseResultHTTPStatsServerResponseTimeMsJSON `json:"-"`
}

// dexHTTPTestGetResponseResultHTTPStatsServerResponseTimeMsJSON contains the JSON
// metadata for the struct
// [DexHTTPTestGetResponseResultHTTPStatsServerResponseTimeMs]
type dexHTTPTestGetResponseResultHTTPStatsServerResponseTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseResultHTTPStatsServerResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseResultHTTPStatsServerResponseTimeMsSlot struct {
	Timestamp string                                                            `json:"timestamp,required"`
	Value     int64                                                             `json:"value,required"`
	JSON      dexHTTPTestGetResponseResultHTTPStatsServerResponseTimeMsSlotJSON `json:"-"`
}

// dexHTTPTestGetResponseResultHTTPStatsServerResponseTimeMsSlotJSON contains the
// JSON metadata for the struct
// [DexHTTPTestGetResponseResultHTTPStatsServerResponseTimeMsSlot]
type dexHTTPTestGetResponseResultHTTPStatsServerResponseTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseResultHTTPStatsServerResponseTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseResultHTTPStatsByColo struct {
	Colo                 string                                                          `json:"colo,required"`
	DNSResponseTimeMs    DexHTTPTestGetResponseResultHTTPStatsByColoDNSResponseTimeMs    `json:"dnsResponseTimeMs,required"`
	HTTPStatusCode       []DexHTTPTestGetResponseResultHTTPStatsByColoHTTPStatusCode     `json:"httpStatusCode,required"`
	ResourceFetchTimeMs  DexHTTPTestGetResponseResultHTTPStatsByColoResourceFetchTimeMs  `json:"resourceFetchTimeMs,required"`
	ServerResponseTimeMs DexHTTPTestGetResponseResultHTTPStatsByColoServerResponseTimeMs `json:"serverResponseTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                                           `json:"uniqueDevicesTotal,required"`
	JSON               dexHTTPTestGetResponseResultHTTPStatsByColoJSON `json:"-"`
}

// dexHTTPTestGetResponseResultHTTPStatsByColoJSON contains the JSON metadata for
// the struct [DexHTTPTestGetResponseResultHTTPStatsByColo]
type dexHTTPTestGetResponseResultHTTPStatsByColoJSON struct {
	Colo                 apijson.Field
	DNSResponseTimeMs    apijson.Field
	HTTPStatusCode       apijson.Field
	ResourceFetchTimeMs  apijson.Field
	ServerResponseTimeMs apijson.Field
	UniqueDevicesTotal   apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseResultHTTPStatsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseResultHTTPStatsByColoDNSResponseTimeMs struct {
	Slots []DexHTTPTestGetResponseResultHTTPStatsByColoDNSResponseTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                            `json:"min,nullable"`
	JSON dexHTTPTestGetResponseResultHTTPStatsByColoDNSResponseTimeMsJSON `json:"-"`
}

// dexHTTPTestGetResponseResultHTTPStatsByColoDNSResponseTimeMsJSON contains the
// JSON metadata for the struct
// [DexHTTPTestGetResponseResultHTTPStatsByColoDNSResponseTimeMs]
type dexHTTPTestGetResponseResultHTTPStatsByColoDNSResponseTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseResultHTTPStatsByColoDNSResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseResultHTTPStatsByColoDNSResponseTimeMsSlot struct {
	Timestamp string                                                               `json:"timestamp,required"`
	Value     int64                                                                `json:"value,required"`
	JSON      dexHTTPTestGetResponseResultHTTPStatsByColoDNSResponseTimeMsSlotJSON `json:"-"`
}

// dexHTTPTestGetResponseResultHTTPStatsByColoDNSResponseTimeMsSlotJSON contains
// the JSON metadata for the struct
// [DexHTTPTestGetResponseResultHTTPStatsByColoDNSResponseTimeMsSlot]
type dexHTTPTestGetResponseResultHTTPStatsByColoDNSResponseTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseResultHTTPStatsByColoDNSResponseTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseResultHTTPStatsByColoHTTPStatusCode struct {
	Status200 int64                                                         `json:"status200,required"`
	Status300 int64                                                         `json:"status300,required"`
	Status400 int64                                                         `json:"status400,required"`
	Status500 int64                                                         `json:"status500,required"`
	Timestamp string                                                        `json:"timestamp,required"`
	JSON      dexHTTPTestGetResponseResultHTTPStatsByColoHTTPStatusCodeJSON `json:"-"`
}

// dexHTTPTestGetResponseResultHTTPStatsByColoHTTPStatusCodeJSON contains the JSON
// metadata for the struct
// [DexHTTPTestGetResponseResultHTTPStatsByColoHTTPStatusCode]
type dexHTTPTestGetResponseResultHTTPStatsByColoHTTPStatusCodeJSON struct {
	Status200   apijson.Field
	Status300   apijson.Field
	Status400   apijson.Field
	Status500   apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseResultHTTPStatsByColoHTTPStatusCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseResultHTTPStatsByColoResourceFetchTimeMs struct {
	Slots []DexHTTPTestGetResponseResultHTTPStatsByColoResourceFetchTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                              `json:"min,nullable"`
	JSON dexHTTPTestGetResponseResultHTTPStatsByColoResourceFetchTimeMsJSON `json:"-"`
}

// dexHTTPTestGetResponseResultHTTPStatsByColoResourceFetchTimeMsJSON contains the
// JSON metadata for the struct
// [DexHTTPTestGetResponseResultHTTPStatsByColoResourceFetchTimeMs]
type dexHTTPTestGetResponseResultHTTPStatsByColoResourceFetchTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseResultHTTPStatsByColoResourceFetchTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseResultHTTPStatsByColoResourceFetchTimeMsSlot struct {
	Timestamp string                                                                 `json:"timestamp,required"`
	Value     int64                                                                  `json:"value,required"`
	JSON      dexHTTPTestGetResponseResultHTTPStatsByColoResourceFetchTimeMsSlotJSON `json:"-"`
}

// dexHTTPTestGetResponseResultHTTPStatsByColoResourceFetchTimeMsSlotJSON contains
// the JSON metadata for the struct
// [DexHTTPTestGetResponseResultHTTPStatsByColoResourceFetchTimeMsSlot]
type dexHTTPTestGetResponseResultHTTPStatsByColoResourceFetchTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseResultHTTPStatsByColoResourceFetchTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseResultHTTPStatsByColoServerResponseTimeMs struct {
	Slots []DexHTTPTestGetResponseResultHTTPStatsByColoServerResponseTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                               `json:"min,nullable"`
	JSON dexHTTPTestGetResponseResultHTTPStatsByColoServerResponseTimeMsJSON `json:"-"`
}

// dexHTTPTestGetResponseResultHTTPStatsByColoServerResponseTimeMsJSON contains the
// JSON metadata for the struct
// [DexHTTPTestGetResponseResultHTTPStatsByColoServerResponseTimeMs]
type dexHTTPTestGetResponseResultHTTPStatsByColoServerResponseTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseResultHTTPStatsByColoServerResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseResultHTTPStatsByColoServerResponseTimeMsSlot struct {
	Timestamp string                                                                  `json:"timestamp,required"`
	Value     int64                                                                   `json:"value,required"`
	JSON      dexHTTPTestGetResponseResultHTTPStatsByColoServerResponseTimeMsSlotJSON `json:"-"`
}

// dexHTTPTestGetResponseResultHTTPStatsByColoServerResponseTimeMsSlotJSON contains
// the JSON metadata for the struct
// [DexHTTPTestGetResponseResultHTTPStatsByColoServerResponseTimeMsSlot]
type dexHTTPTestGetResponseResultHTTPStatsByColoServerResponseTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestGetResponseResultHTTPStatsByColoServerResponseTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestGetResponseResultKind string

const (
	DexHTTPTestGetResponseResultKindHTTP DexHTTPTestGetResponseResultKind = "http"
)

// Whether the API call was successful
type DexHTTPTestGetResponseSuccess bool

const (
	DexHTTPTestGetResponseSuccessTrue DexHTTPTestGetResponseSuccess = true
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
