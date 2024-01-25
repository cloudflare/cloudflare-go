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

// AccountDexHTTPTestService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountDexHTTPTestService] method
// instead.
type AccountDexHTTPTestService struct {
	Options     []option.RequestOption
	Percentiles *AccountDexHTTPTestPercentileService
}

// NewAccountDexHTTPTestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDexHTTPTestService(opts ...option.RequestOption) (r *AccountDexHTTPTestService) {
	r = &AccountDexHTTPTestService{}
	r.Options = opts
	r.Percentiles = NewAccountDexHTTPTestPercentileService(opts...)
	return
}

// Get test details and aggregate performance metrics for an http test for a given
// time period between 1 hour and 7 days.
func (r *AccountDexHTTPTestService) Get(ctx context.Context, accountIdentifier string, testID string, query AccountDexHTTPTestGetParams, opts ...option.RequestOption) (res *AccountDexHTTPTestGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dex/http-tests/%s", accountIdentifier, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountDexHTTPTestGetResponse struct {
	Errors   []AccountDexHTTPTestGetResponseError   `json:"errors"`
	Messages []AccountDexHTTPTestGetResponseMessage `json:"messages"`
	Result   AccountDexHTTPTestGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountDexHTTPTestGetResponseSuccess `json:"success"`
	JSON    accountDexHTTPTestGetResponseJSON    `json:"-"`
}

// accountDexHTTPTestGetResponseJSON contains the JSON metadata for the struct
// [AccountDexHTTPTestGetResponse]
type accountDexHTTPTestGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexHTTPTestGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexHTTPTestGetResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    accountDexHTTPTestGetResponseErrorJSON `json:"-"`
}

// accountDexHTTPTestGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountDexHTTPTestGetResponseError]
type accountDexHTTPTestGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexHTTPTestGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexHTTPTestGetResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accountDexHTTPTestGetResponseMessageJSON `json:"-"`
}

// accountDexHTTPTestGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountDexHTTPTestGetResponseMessage]
type accountDexHTTPTestGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexHTTPTestGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexHTTPTestGetResponseResult struct {
	// The url of the HTTP synthetic application test
	Host            string                                               `json:"host"`
	HTTPStats       AccountDexHTTPTestGetResponseResultHTTPStats         `json:"httpStats,nullable"`
	HTTPStatsByColo []AccountDexHTTPTestGetResponseResultHTTPStatsByColo `json:"httpStatsByColo"`
	// The interval at which the HTTP synthetic application test is set to run.
	Interval string                                  `json:"interval"`
	Kind     AccountDexHTTPTestGetResponseResultKind `json:"kind"`
	// The HTTP method to use when running the test
	Method string `json:"method"`
	// The name of the HTTP synthetic application test
	Name string                                  `json:"name"`
	JSON accountDexHTTPTestGetResponseResultJSON `json:"-"`
}

// accountDexHTTPTestGetResponseResultJSON contains the JSON metadata for the
// struct [AccountDexHTTPTestGetResponseResult]
type accountDexHTTPTestGetResponseResultJSON struct {
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

func (r *AccountDexHTTPTestGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexHTTPTestGetResponseResultHTTPStats struct {
	DNSResponseTimeMs    AccountDexHTTPTestGetResponseResultHTTPStatsDNSResponseTimeMs    `json:"dnsResponseTimeMs,required"`
	HTTPStatusCode       []AccountDexHTTPTestGetResponseResultHTTPStatsHTTPStatusCode     `json:"httpStatusCode,required"`
	ResourceFetchTimeMs  AccountDexHTTPTestGetResponseResultHTTPStatsResourceFetchTimeMs  `json:"resourceFetchTimeMs,required"`
	ServerResponseTimeMs AccountDexHTTPTestGetResponseResultHTTPStatsServerResponseTimeMs `json:"serverResponseTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                                            `json:"uniqueDevicesTotal,required"`
	JSON               accountDexHTTPTestGetResponseResultHTTPStatsJSON `json:"-"`
}

// accountDexHTTPTestGetResponseResultHTTPStatsJSON contains the JSON metadata for
// the struct [AccountDexHTTPTestGetResponseResultHTTPStats]
type accountDexHTTPTestGetResponseResultHTTPStatsJSON struct {
	DNSResponseTimeMs    apijson.Field
	HTTPStatusCode       apijson.Field
	ResourceFetchTimeMs  apijson.Field
	ServerResponseTimeMs apijson.Field
	UniqueDevicesTotal   apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountDexHTTPTestGetResponseResultHTTPStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexHTTPTestGetResponseResultHTTPStatsDNSResponseTimeMs struct {
	Slots []AccountDexHTTPTestGetResponseResultHTTPStatsDNSResponseTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                             `json:"min,nullable"`
	JSON accountDexHTTPTestGetResponseResultHTTPStatsDNSResponseTimeMsJSON `json:"-"`
}

// accountDexHTTPTestGetResponseResultHTTPStatsDNSResponseTimeMsJSON contains the
// JSON metadata for the struct
// [AccountDexHTTPTestGetResponseResultHTTPStatsDNSResponseTimeMs]
type accountDexHTTPTestGetResponseResultHTTPStatsDNSResponseTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexHTTPTestGetResponseResultHTTPStatsDNSResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexHTTPTestGetResponseResultHTTPStatsDNSResponseTimeMsSlot struct {
	Timestamp string                                                                `json:"timestamp,required"`
	Value     int64                                                                 `json:"value,required"`
	JSON      accountDexHTTPTestGetResponseResultHTTPStatsDNSResponseTimeMsSlotJSON `json:"-"`
}

// accountDexHTTPTestGetResponseResultHTTPStatsDNSResponseTimeMsSlotJSON contains
// the JSON metadata for the struct
// [AccountDexHTTPTestGetResponseResultHTTPStatsDNSResponseTimeMsSlot]
type accountDexHTTPTestGetResponseResultHTTPStatsDNSResponseTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexHTTPTestGetResponseResultHTTPStatsDNSResponseTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexHTTPTestGetResponseResultHTTPStatsHTTPStatusCode struct {
	Status200 int64                                                          `json:"status200,required"`
	Status300 int64                                                          `json:"status300,required"`
	Status400 int64                                                          `json:"status400,required"`
	Status500 int64                                                          `json:"status500,required"`
	Timestamp string                                                         `json:"timestamp,required"`
	JSON      accountDexHTTPTestGetResponseResultHTTPStatsHTTPStatusCodeJSON `json:"-"`
}

// accountDexHTTPTestGetResponseResultHTTPStatsHTTPStatusCodeJSON contains the JSON
// metadata for the struct
// [AccountDexHTTPTestGetResponseResultHTTPStatsHTTPStatusCode]
type accountDexHTTPTestGetResponseResultHTTPStatsHTTPStatusCodeJSON struct {
	Status200   apijson.Field
	Status300   apijson.Field
	Status400   apijson.Field
	Status500   apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexHTTPTestGetResponseResultHTTPStatsHTTPStatusCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexHTTPTestGetResponseResultHTTPStatsResourceFetchTimeMs struct {
	Slots []AccountDexHTTPTestGetResponseResultHTTPStatsResourceFetchTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                               `json:"min,nullable"`
	JSON accountDexHTTPTestGetResponseResultHTTPStatsResourceFetchTimeMsJSON `json:"-"`
}

// accountDexHTTPTestGetResponseResultHTTPStatsResourceFetchTimeMsJSON contains the
// JSON metadata for the struct
// [AccountDexHTTPTestGetResponseResultHTTPStatsResourceFetchTimeMs]
type accountDexHTTPTestGetResponseResultHTTPStatsResourceFetchTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexHTTPTestGetResponseResultHTTPStatsResourceFetchTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexHTTPTestGetResponseResultHTTPStatsResourceFetchTimeMsSlot struct {
	Timestamp string                                                                  `json:"timestamp,required"`
	Value     int64                                                                   `json:"value,required"`
	JSON      accountDexHTTPTestGetResponseResultHTTPStatsResourceFetchTimeMsSlotJSON `json:"-"`
}

// accountDexHTTPTestGetResponseResultHTTPStatsResourceFetchTimeMsSlotJSON contains
// the JSON metadata for the struct
// [AccountDexHTTPTestGetResponseResultHTTPStatsResourceFetchTimeMsSlot]
type accountDexHTTPTestGetResponseResultHTTPStatsResourceFetchTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexHTTPTestGetResponseResultHTTPStatsResourceFetchTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexHTTPTestGetResponseResultHTTPStatsServerResponseTimeMs struct {
	Slots []AccountDexHTTPTestGetResponseResultHTTPStatsServerResponseTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                                `json:"min,nullable"`
	JSON accountDexHTTPTestGetResponseResultHTTPStatsServerResponseTimeMsJSON `json:"-"`
}

// accountDexHTTPTestGetResponseResultHTTPStatsServerResponseTimeMsJSON contains
// the JSON metadata for the struct
// [AccountDexHTTPTestGetResponseResultHTTPStatsServerResponseTimeMs]
type accountDexHTTPTestGetResponseResultHTTPStatsServerResponseTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexHTTPTestGetResponseResultHTTPStatsServerResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexHTTPTestGetResponseResultHTTPStatsServerResponseTimeMsSlot struct {
	Timestamp string                                                                   `json:"timestamp,required"`
	Value     int64                                                                    `json:"value,required"`
	JSON      accountDexHTTPTestGetResponseResultHTTPStatsServerResponseTimeMsSlotJSON `json:"-"`
}

// accountDexHTTPTestGetResponseResultHTTPStatsServerResponseTimeMsSlotJSON
// contains the JSON metadata for the struct
// [AccountDexHTTPTestGetResponseResultHTTPStatsServerResponseTimeMsSlot]
type accountDexHTTPTestGetResponseResultHTTPStatsServerResponseTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexHTTPTestGetResponseResultHTTPStatsServerResponseTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexHTTPTestGetResponseResultHTTPStatsByColo struct {
	Colo                 string                                                                 `json:"colo,required"`
	DNSResponseTimeMs    AccountDexHTTPTestGetResponseResultHTTPStatsByColoDNSResponseTimeMs    `json:"dnsResponseTimeMs,required"`
	HTTPStatusCode       []AccountDexHTTPTestGetResponseResultHTTPStatsByColoHTTPStatusCode     `json:"httpStatusCode,required"`
	ResourceFetchTimeMs  AccountDexHTTPTestGetResponseResultHTTPStatsByColoResourceFetchTimeMs  `json:"resourceFetchTimeMs,required"`
	ServerResponseTimeMs AccountDexHTTPTestGetResponseResultHTTPStatsByColoServerResponseTimeMs `json:"serverResponseTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                                                  `json:"uniqueDevicesTotal,required"`
	JSON               accountDexHTTPTestGetResponseResultHTTPStatsByColoJSON `json:"-"`
}

// accountDexHTTPTestGetResponseResultHTTPStatsByColoJSON contains the JSON
// metadata for the struct [AccountDexHTTPTestGetResponseResultHTTPStatsByColo]
type accountDexHTTPTestGetResponseResultHTTPStatsByColoJSON struct {
	Colo                 apijson.Field
	DNSResponseTimeMs    apijson.Field
	HTTPStatusCode       apijson.Field
	ResourceFetchTimeMs  apijson.Field
	ServerResponseTimeMs apijson.Field
	UniqueDevicesTotal   apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountDexHTTPTestGetResponseResultHTTPStatsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexHTTPTestGetResponseResultHTTPStatsByColoDNSResponseTimeMs struct {
	Slots []AccountDexHTTPTestGetResponseResultHTTPStatsByColoDNSResponseTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                                   `json:"min,nullable"`
	JSON accountDexHTTPTestGetResponseResultHTTPStatsByColoDNSResponseTimeMsJSON `json:"-"`
}

// accountDexHTTPTestGetResponseResultHTTPStatsByColoDNSResponseTimeMsJSON contains
// the JSON metadata for the struct
// [AccountDexHTTPTestGetResponseResultHTTPStatsByColoDNSResponseTimeMs]
type accountDexHTTPTestGetResponseResultHTTPStatsByColoDNSResponseTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexHTTPTestGetResponseResultHTTPStatsByColoDNSResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexHTTPTestGetResponseResultHTTPStatsByColoDNSResponseTimeMsSlot struct {
	Timestamp string                                                                      `json:"timestamp,required"`
	Value     int64                                                                       `json:"value,required"`
	JSON      accountDexHTTPTestGetResponseResultHTTPStatsByColoDNSResponseTimeMsSlotJSON `json:"-"`
}

// accountDexHTTPTestGetResponseResultHTTPStatsByColoDNSResponseTimeMsSlotJSON
// contains the JSON metadata for the struct
// [AccountDexHTTPTestGetResponseResultHTTPStatsByColoDNSResponseTimeMsSlot]
type accountDexHTTPTestGetResponseResultHTTPStatsByColoDNSResponseTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexHTTPTestGetResponseResultHTTPStatsByColoDNSResponseTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexHTTPTestGetResponseResultHTTPStatsByColoHTTPStatusCode struct {
	Status200 int64                                                                `json:"status200,required"`
	Status300 int64                                                                `json:"status300,required"`
	Status400 int64                                                                `json:"status400,required"`
	Status500 int64                                                                `json:"status500,required"`
	Timestamp string                                                               `json:"timestamp,required"`
	JSON      accountDexHTTPTestGetResponseResultHTTPStatsByColoHTTPStatusCodeJSON `json:"-"`
}

// accountDexHTTPTestGetResponseResultHTTPStatsByColoHTTPStatusCodeJSON contains
// the JSON metadata for the struct
// [AccountDexHTTPTestGetResponseResultHTTPStatsByColoHTTPStatusCode]
type accountDexHTTPTestGetResponseResultHTTPStatsByColoHTTPStatusCodeJSON struct {
	Status200   apijson.Field
	Status300   apijson.Field
	Status400   apijson.Field
	Status500   apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexHTTPTestGetResponseResultHTTPStatsByColoHTTPStatusCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexHTTPTestGetResponseResultHTTPStatsByColoResourceFetchTimeMs struct {
	Slots []AccountDexHTTPTestGetResponseResultHTTPStatsByColoResourceFetchTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                                     `json:"min,nullable"`
	JSON accountDexHTTPTestGetResponseResultHTTPStatsByColoResourceFetchTimeMsJSON `json:"-"`
}

// accountDexHTTPTestGetResponseResultHTTPStatsByColoResourceFetchTimeMsJSON
// contains the JSON metadata for the struct
// [AccountDexHTTPTestGetResponseResultHTTPStatsByColoResourceFetchTimeMs]
type accountDexHTTPTestGetResponseResultHTTPStatsByColoResourceFetchTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexHTTPTestGetResponseResultHTTPStatsByColoResourceFetchTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexHTTPTestGetResponseResultHTTPStatsByColoResourceFetchTimeMsSlot struct {
	Timestamp string                                                                        `json:"timestamp,required"`
	Value     int64                                                                         `json:"value,required"`
	JSON      accountDexHTTPTestGetResponseResultHTTPStatsByColoResourceFetchTimeMsSlotJSON `json:"-"`
}

// accountDexHTTPTestGetResponseResultHTTPStatsByColoResourceFetchTimeMsSlotJSON
// contains the JSON metadata for the struct
// [AccountDexHTTPTestGetResponseResultHTTPStatsByColoResourceFetchTimeMsSlot]
type accountDexHTTPTestGetResponseResultHTTPStatsByColoResourceFetchTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexHTTPTestGetResponseResultHTTPStatsByColoResourceFetchTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexHTTPTestGetResponseResultHTTPStatsByColoServerResponseTimeMs struct {
	Slots []AccountDexHTTPTestGetResponseResultHTTPStatsByColoServerResponseTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                                      `json:"min,nullable"`
	JSON accountDexHTTPTestGetResponseResultHTTPStatsByColoServerResponseTimeMsJSON `json:"-"`
}

// accountDexHTTPTestGetResponseResultHTTPStatsByColoServerResponseTimeMsJSON
// contains the JSON metadata for the struct
// [AccountDexHTTPTestGetResponseResultHTTPStatsByColoServerResponseTimeMs]
type accountDexHTTPTestGetResponseResultHTTPStatsByColoServerResponseTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexHTTPTestGetResponseResultHTTPStatsByColoServerResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexHTTPTestGetResponseResultHTTPStatsByColoServerResponseTimeMsSlot struct {
	Timestamp string                                                                         `json:"timestamp,required"`
	Value     int64                                                                          `json:"value,required"`
	JSON      accountDexHTTPTestGetResponseResultHTTPStatsByColoServerResponseTimeMsSlotJSON `json:"-"`
}

// accountDexHTTPTestGetResponseResultHTTPStatsByColoServerResponseTimeMsSlotJSON
// contains the JSON metadata for the struct
// [AccountDexHTTPTestGetResponseResultHTTPStatsByColoServerResponseTimeMsSlot]
type accountDexHTTPTestGetResponseResultHTTPStatsByColoServerResponseTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexHTTPTestGetResponseResultHTTPStatsByColoServerResponseTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexHTTPTestGetResponseResultKind string

const (
	AccountDexHTTPTestGetResponseResultKindHTTP AccountDexHTTPTestGetResponseResultKind = "http"
)

// Whether the API call was successful
type AccountDexHTTPTestGetResponseSuccess bool

const (
	AccountDexHTTPTestGetResponseSuccessTrue AccountDexHTTPTestGetResponseSuccess = true
)

type AccountDexHTTPTestGetParams struct {
	// Time interval for aggregate time slots.
	Interval param.Field[AccountDexHTTPTestGetParamsInterval] `query:"interval,required"`
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

// URLQuery serializes [AccountDexHTTPTestGetParams]'s query parameters as
// `url.Values`.
func (r AccountDexHTTPTestGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Time interval for aggregate time slots.
type AccountDexHTTPTestGetParamsInterval string

const (
	AccountDexHTTPTestGetParamsIntervalMinute AccountDexHTTPTestGetParamsInterval = "minute"
	AccountDexHTTPTestGetParamsIntervalHour   AccountDexHTTPTestGetParamsInterval = "hour"
)
