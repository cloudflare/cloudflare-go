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

// AccountDexTestService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountDexTestService] method
// instead.
type AccountDexTestService struct {
	Options       []option.RequestOption
	UniqueDevices *AccountDexTestUniqueDeviceService
}

// NewAccountDexTestService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountDexTestService(opts ...option.RequestOption) (r *AccountDexTestService) {
	r = &AccountDexTestService{}
	r.Options = opts
	r.UniqueDevices = NewAccountDexTestUniqueDeviceService(opts...)
	return
}

// List DEX tests
func (r *AccountDexTestService) List(ctx context.Context, accountIdentifier string, query AccountDexTestListParams, opts ...option.RequestOption) (res *AccountDexTestListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dex/tests", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountDexTestListResponse struct {
	Errors     []AccountDexTestListResponseError    `json:"errors"`
	Messages   []AccountDexTestListResponseMessage  `json:"messages"`
	Result     AccountDexTestListResponseResult     `json:"result"`
	ResultInfo AccountDexTestListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountDexTestListResponseSuccess `json:"success"`
	JSON    accountDexTestListResponseJSON    `json:"-"`
}

// accountDexTestListResponseJSON contains the JSON metadata for the struct
// [AccountDexTestListResponse]
type accountDexTestListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    accountDexTestListResponseErrorJSON `json:"-"`
}

// accountDexTestListResponseErrorJSON contains the JSON metadata for the struct
// [AccountDexTestListResponseError]
type accountDexTestListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    accountDexTestListResponseMessageJSON `json:"-"`
}

// accountDexTestListResponseMessageJSON contains the JSON metadata for the struct
// [AccountDexTestListResponseMessage]
type accountDexTestListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseResult struct {
	OverviewMetrics AccountDexTestListResponseResultOverviewMetrics `json:"overviewMetrics,required"`
	// array of test results objects.
	Tests []AccountDexTestListResponseResultTest `json:"tests,required"`
	JSON  accountDexTestListResponseResultJSON   `json:"-"`
}

// accountDexTestListResponseResultJSON contains the JSON metadata for the struct
// [AccountDexTestListResponseResult]
type accountDexTestListResponseResultJSON struct {
	OverviewMetrics apijson.Field
	Tests           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDexTestListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseResultOverviewMetrics struct {
	// number of tests.
	TestsTotal int64 `json:"testsTotal,required"`
	// percentage availability for all traceroutes results in response
	AvgTracerouteAvailabilityPct float64                                             `json:"avgTracerouteAvailabilityPct,nullable"`
	JSON                         accountDexTestListResponseResultOverviewMetricsJSON `json:"-"`
}

// accountDexTestListResponseResultOverviewMetricsJSON contains the JSON metadata
// for the struct [AccountDexTestListResponseResultOverviewMetrics]
type accountDexTestListResponseResultOverviewMetricsJSON struct {
	TestsTotal                   apijson.Field
	AvgTracerouteAvailabilityPct apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *AccountDexTestListResponseResultOverviewMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseResultTest struct {
	// API Resource UUID tag.
	ID string `json:"id,required"`
	// date the test was created.
	Created string `json:"created,required"`
	// the test description defined during configuration
	Description string `json:"description,required"`
	// if true, then the test will run on targeted devices. Else, the test will not
	// run.
	Enabled bool   `json:"enabled,required"`
	Host    string `json:"host,required"`
	// The interval at which the synthetic application test is set to run.
	Interval string `json:"interval,required"`
	// test type, http or traceroute
	Kind AccountDexTestListResponseResultTestsKind `json:"kind,required"`
	// name given to this test
	Name              string                                                   `json:"name,required"`
	Updated           string                                                   `json:"updated,required"`
	HTTPResults       AccountDexTestListResponseResultTestsHTTPResults         `json:"httpResults,nullable"`
	HTTPResultsByColo []AccountDexTestListResponseResultTestsHTTPResultsByColo `json:"httpResultsByColo"`
	// for HTTP, the method to use when running the test
	Method                  string                                                         `json:"method"`
	TracerouteResults       AccountDexTestListResponseResultTestsTracerouteResults         `json:"tracerouteResults,nullable"`
	TracerouteResultsByColo []AccountDexTestListResponseResultTestsTracerouteResultsByColo `json:"tracerouteResultsByColo"`
	JSON                    accountDexTestListResponseResultTestJSON                       `json:"-"`
}

// accountDexTestListResponseResultTestJSON contains the JSON metadata for the
// struct [AccountDexTestListResponseResultTest]
type accountDexTestListResponseResultTestJSON struct {
	ID                      apijson.Field
	Created                 apijson.Field
	Description             apijson.Field
	Enabled                 apijson.Field
	Host                    apijson.Field
	Interval                apijson.Field
	Kind                    apijson.Field
	Name                    apijson.Field
	Updated                 apijson.Field
	HTTPResults             apijson.Field
	HTTPResultsByColo       apijson.Field
	Method                  apijson.Field
	TracerouteResults       apijson.Field
	TracerouteResultsByColo apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccountDexTestListResponseResultTest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// test type, http or traceroute
type AccountDexTestListResponseResultTestsKind string

const (
	AccountDexTestListResponseResultTestsKindHTTP       AccountDexTestListResponseResultTestsKind = "http"
	AccountDexTestListResponseResultTestsKindTraceroute AccountDexTestListResponseResultTestsKind = "traceroute"
)

type AccountDexTestListResponseResultTestsHTTPResults struct {
	ResourceFetchTime AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTime `json:"resourceFetchTime,required"`
	JSON              accountDexTestListResponseResultTestsHTTPResultsJSON              `json:"-"`
}

// accountDexTestListResponseResultTestsHTTPResultsJSON contains the JSON metadata
// for the struct [AccountDexTestListResponseResultTestsHTTPResults]
type accountDexTestListResponseResultTestsHTTPResultsJSON struct {
	ResourceFetchTime apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountDexTestListResponseResultTestsHTTPResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTime struct {
	History  []AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistory `json:"history,required"`
	AvgMs    int64                                                                      `json:"avgMs,nullable"`
	OverTime AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTime  `json:"overTime,nullable"`
	JSON     accountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeJSON      `json:"-"`
}

// accountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeJSON contains
// the JSON metadata for the struct
// [AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTime]
type accountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistory struct {
	TimePeriod AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriod `json:"timePeriod,required"`
	AvgMs      int64                                                                              `json:"avgMs,nullable"`
	DeltaPct   float64                                                                            `json:"deltaPct,nullable"`
	JSON       accountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryJSON       `json:"-"`
}

// accountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryJSON
// contains the JSON metadata for the struct
// [AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistory]
type accountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriod struct {
	Units AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits `json:"units,required"`
	Value int64                                                                                   `json:"value,required"`
	JSON  accountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodJSON  `json:"-"`
}

// accountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodJSON
// contains the JSON metadata for the struct
// [AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriod]
type accountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits string

const (
	AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnitsHours    AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits = "hours"
	AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnitsDays     AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits = "days"
	AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnitsTestRuns AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits = "testRuns"
)

type AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTime struct {
	TimePeriod AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriod `json:"timePeriod,required"`
	Values     []AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeValue    `json:"values,required"`
	JSON       accountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeJSON       `json:"-"`
}

// accountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeJSON
// contains the JSON metadata for the struct
// [AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTime]
type accountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriod struct {
	Units AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits `json:"units,required"`
	Value int64                                                                                    `json:"value,required"`
	JSON  accountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodJSON  `json:"-"`
}

// accountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodJSON
// contains the JSON metadata for the struct
// [AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriod]
type accountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits string

const (
	AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnitsHours    AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits = "hours"
	AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnitsDays     AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits = "days"
	AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnitsTestRuns AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits = "testRuns"
)

type AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeValue struct {
	AvgMs     int64                                                                              `json:"avgMs,required"`
	Timestamp string                                                                             `json:"timestamp,required"`
	JSON      accountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeValueJSON `json:"-"`
}

// accountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeValueJSON
// contains the JSON metadata for the struct
// [AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeValue]
type accountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseResultTestsHTTPResultsByColo struct {
	// Cloudflare colo
	Colo              string                                                                  `json:"colo,required"`
	ResourceFetchTime AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTime `json:"resourceFetchTime,required"`
	JSON              accountDexTestListResponseResultTestsHTTPResultsByColoJSON              `json:"-"`
}

// accountDexTestListResponseResultTestsHTTPResultsByColoJSON contains the JSON
// metadata for the struct [AccountDexTestListResponseResultTestsHTTPResultsByColo]
type accountDexTestListResponseResultTestsHTTPResultsByColoJSON struct {
	Colo              apijson.Field
	ResourceFetchTime apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountDexTestListResponseResultTestsHTTPResultsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTime struct {
	History  []AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistory `json:"history,required"`
	AvgMs    int64                                                                            `json:"avgMs,nullable"`
	OverTime AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTime  `json:"overTime,nullable"`
	JSON     accountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeJSON      `json:"-"`
}

// accountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeJSON
// contains the JSON metadata for the struct
// [AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTime]
type accountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistory struct {
	TimePeriod AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriod `json:"timePeriod,required"`
	AvgMs      int64                                                                                    `json:"avgMs,nullable"`
	DeltaPct   float64                                                                                  `json:"deltaPct,nullable"`
	JSON       accountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryJSON       `json:"-"`
}

// accountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryJSON
// contains the JSON metadata for the struct
// [AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistory]
type accountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriod struct {
	Units AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits `json:"units,required"`
	Value int64                                                                                         `json:"value,required"`
	JSON  accountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodJSON  `json:"-"`
}

// accountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodJSON
// contains the JSON metadata for the struct
// [AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriod]
type accountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits string

const (
	AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnitsHours    AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits = "hours"
	AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnitsDays     AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits = "days"
	AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnitsTestRuns AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits = "testRuns"
)

type AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTime struct {
	TimePeriod AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriod `json:"timePeriod,required"`
	Values     []AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeValue    `json:"values,required"`
	JSON       accountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeJSON       `json:"-"`
}

// accountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeJSON
// contains the JSON metadata for the struct
// [AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTime]
type accountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriod struct {
	Units AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits `json:"units,required"`
	Value int64                                                                                          `json:"value,required"`
	JSON  accountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodJSON  `json:"-"`
}

// accountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodJSON
// contains the JSON metadata for the struct
// [AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriod]
type accountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits string

const (
	AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnitsHours    AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits = "hours"
	AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnitsDays     AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits = "days"
	AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnitsTestRuns AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits = "testRuns"
)

type AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeValue struct {
	AvgMs     int64                                                                                    `json:"avgMs,required"`
	Timestamp string                                                                                   `json:"timestamp,required"`
	JSON      accountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeValueJSON `json:"-"`
}

// accountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeValueJSON
// contains the JSON metadata for the struct
// [AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeValue]
type accountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseResultTestsTracerouteResults struct {
	RoundTripTime AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTime `json:"roundTripTime,required"`
	JSON          accountDexTestListResponseResultTestsTracerouteResultsJSON          `json:"-"`
}

// accountDexTestListResponseResultTestsTracerouteResultsJSON contains the JSON
// metadata for the struct [AccountDexTestListResponseResultTestsTracerouteResults]
type accountDexTestListResponseResultTestsTracerouteResultsJSON struct {
	RoundTripTime apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountDexTestListResponseResultTestsTracerouteResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTime struct {
	History  []AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistory `json:"history,required"`
	AvgMs    int64                                                                        `json:"avgMs,nullable"`
	OverTime AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTime  `json:"overTime,nullable"`
	JSON     accountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeJSON      `json:"-"`
}

// accountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeJSON contains
// the JSON metadata for the struct
// [AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTime]
type accountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistory struct {
	TimePeriod AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriod `json:"timePeriod,required"`
	AvgMs      int64                                                                                `json:"avgMs,nullable"`
	DeltaPct   float64                                                                              `json:"deltaPct,nullable"`
	JSON       accountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryJSON       `json:"-"`
}

// accountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryJSON
// contains the JSON metadata for the struct
// [AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistory]
type accountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriod struct {
	Units AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits `json:"units,required"`
	Value int64                                                                                     `json:"value,required"`
	JSON  accountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodJSON  `json:"-"`
}

// accountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodJSON
// contains the JSON metadata for the struct
// [AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriod]
type accountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits string

const (
	AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnitsHours    AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits = "hours"
	AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnitsDays     AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits = "days"
	AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnitsTestRuns AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits = "testRuns"
)

type AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTime struct {
	TimePeriod AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriod `json:"timePeriod,required"`
	Values     []AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeValue    `json:"values,required"`
	JSON       accountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeJSON       `json:"-"`
}

// accountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeJSON
// contains the JSON metadata for the struct
// [AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTime]
type accountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriod struct {
	Units AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits `json:"units,required"`
	Value int64                                                                                      `json:"value,required"`
	JSON  accountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodJSON  `json:"-"`
}

// accountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodJSON
// contains the JSON metadata for the struct
// [AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriod]
type accountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits string

const (
	AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnitsHours    AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits = "hours"
	AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnitsDays     AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits = "days"
	AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnitsTestRuns AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits = "testRuns"
)

type AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeValue struct {
	AvgMs     int64                                                                                `json:"avgMs,required"`
	Timestamp string                                                                               `json:"timestamp,required"`
	JSON      accountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeValueJSON `json:"-"`
}

// accountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeValueJSON
// contains the JSON metadata for the struct
// [AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeValue]
type accountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseResultTestsTracerouteResultsByColo struct {
	// Cloudflare colo
	Colo          string                                                                    `json:"colo,required"`
	RoundTripTime AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTime `json:"roundTripTime,required"`
	JSON          accountDexTestListResponseResultTestsTracerouteResultsByColoJSON          `json:"-"`
}

// accountDexTestListResponseResultTestsTracerouteResultsByColoJSON contains the
// JSON metadata for the struct
// [AccountDexTestListResponseResultTestsTracerouteResultsByColo]
type accountDexTestListResponseResultTestsTracerouteResultsByColoJSON struct {
	Colo          apijson.Field
	RoundTripTime apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountDexTestListResponseResultTestsTracerouteResultsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTime struct {
	History  []AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistory `json:"history,required"`
	AvgMs    int64                                                                              `json:"avgMs,nullable"`
	OverTime AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTime  `json:"overTime,nullable"`
	JSON     accountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeJSON      `json:"-"`
}

// accountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeJSON
// contains the JSON metadata for the struct
// [AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTime]
type accountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistory struct {
	TimePeriod AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriod `json:"timePeriod,required"`
	AvgMs      int64                                                                                      `json:"avgMs,nullable"`
	DeltaPct   float64                                                                                    `json:"deltaPct,nullable"`
	JSON       accountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryJSON       `json:"-"`
}

// accountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryJSON
// contains the JSON metadata for the struct
// [AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistory]
type accountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriod struct {
	Units AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits `json:"units,required"`
	Value int64                                                                                           `json:"value,required"`
	JSON  accountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodJSON  `json:"-"`
}

// accountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodJSON
// contains the JSON metadata for the struct
// [AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriod]
type accountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits string

const (
	AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnitsHours    AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits = "hours"
	AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnitsDays     AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits = "days"
	AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnitsTestRuns AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits = "testRuns"
)

type AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTime struct {
	TimePeriod AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriod `json:"timePeriod,required"`
	Values     []AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeValue    `json:"values,required"`
	JSON       accountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeJSON       `json:"-"`
}

// accountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeJSON
// contains the JSON metadata for the struct
// [AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTime]
type accountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriod struct {
	Units AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits `json:"units,required"`
	Value int64                                                                                            `json:"value,required"`
	JSON  accountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodJSON  `json:"-"`
}

// accountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodJSON
// contains the JSON metadata for the struct
// [AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriod]
type accountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits string

const (
	AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnitsHours    AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits = "hours"
	AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnitsDays     AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits = "days"
	AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnitsTestRuns AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits = "testRuns"
)

type AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeValue struct {
	AvgMs     int64                                                                                      `json:"avgMs,required"`
	Timestamp string                                                                                     `json:"timestamp,required"`
	JSON      accountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeValueJSON `json:"-"`
}

// accountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeValueJSON
// contains the JSON metadata for the struct
// [AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeValue]
type accountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       accountDexTestListResponseResultInfoJSON `json:"-"`
}

// accountDexTestListResponseResultInfoJSON contains the JSON metadata for the
// struct [AccountDexTestListResponseResultInfo]
type accountDexTestListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountDexTestListResponseSuccess bool

const (
	AccountDexTestListResponseSuccessTrue AccountDexTestListResponseSuccess = true
)

type AccountDexTestListParams struct {
	// Optionally filter result stats to a Cloudflare colo. Cannot be used in
	// combination with deviceId param.
	Colo param.Field[string] `query:"colo"`
	// Optionally filter result stats to a specific device(s). Cannot be used in
	// combination with colo param.
	DeviceID param.Field[[]string] `query:"deviceId"`
	// Page number of paginated results
	Page param.Field[float64] `query:"page"`
	// Number of items per page
	PerPage param.Field[float64] `query:"per_page"`
	// Optionally filter results by test name
	TestName param.Field[string] `query:"testName"`
}

// URLQuery serializes [AccountDexTestListParams]'s query parameters as
// `url.Values`.
func (r AccountDexTestListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
