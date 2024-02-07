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

// DexTestService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewDexTestService] method instead.
type DexTestService struct {
	Options       []option.RequestOption
	UniqueDevices *DexTestUniqueDeviceService
}

// NewDexTestService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDexTestService(opts ...option.RequestOption) (r *DexTestService) {
	r = &DexTestService{}
	r.Options = opts
	r.UniqueDevices = NewDexTestUniqueDeviceService(opts...)
	return
}

// List DEX tests
func (r *DexTestService) List(ctx context.Context, accountID string, query DexTestListParams, opts ...option.RequestOption) (res *DexTestListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dex/tests", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DexTestListResponse struct {
	Errors     []DexTestListResponseError    `json:"errors"`
	Messages   []DexTestListResponseMessage  `json:"messages"`
	Result     DexTestListResponseResult     `json:"result"`
	ResultInfo DexTestListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success DexTestListResponseSuccess `json:"success"`
	JSON    dexTestListResponseJSON    `json:"-"`
}

// dexTestListResponseJSON contains the JSON metadata for the struct
// [DexTestListResponse]
type dexTestListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseError struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    dexTestListResponseErrorJSON `json:"-"`
}

// dexTestListResponseErrorJSON contains the JSON metadata for the struct
// [DexTestListResponseError]
type dexTestListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseMessage struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    dexTestListResponseMessageJSON `json:"-"`
}

// dexTestListResponseMessageJSON contains the JSON metadata for the struct
// [DexTestListResponseMessage]
type dexTestListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseResult struct {
	OverviewMetrics DexTestListResponseResultOverviewMetrics `json:"overviewMetrics,required"`
	// array of test results objects.
	Tests []DexTestListResponseResultTest `json:"tests,required"`
	JSON  dexTestListResponseResultJSON   `json:"-"`
}

// dexTestListResponseResultJSON contains the JSON metadata for the struct
// [DexTestListResponseResult]
type dexTestListResponseResultJSON struct {
	OverviewMetrics apijson.Field
	Tests           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DexTestListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseResultOverviewMetrics struct {
	// number of tests.
	TestsTotal int64 `json:"testsTotal,required"`
	// percentage availability for all traceroutes results in response
	AvgTracerouteAvailabilityPct float64                                      `json:"avgTracerouteAvailabilityPct,nullable"`
	JSON                         dexTestListResponseResultOverviewMetricsJSON `json:"-"`
}

// dexTestListResponseResultOverviewMetricsJSON contains the JSON metadata for the
// struct [DexTestListResponseResultOverviewMetrics]
type dexTestListResponseResultOverviewMetricsJSON struct {
	TestsTotal                   apijson.Field
	AvgTracerouteAvailabilityPct apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *DexTestListResponseResultOverviewMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseResultTest struct {
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
	Kind DexTestListResponseResultTestsKind `json:"kind,required"`
	// name given to this test
	Name              string                                            `json:"name,required"`
	Updated           string                                            `json:"updated,required"`
	HTTPResults       DexTestListResponseResultTestsHTTPResults         `json:"httpResults,nullable"`
	HTTPResultsByColo []DexTestListResponseResultTestsHTTPResultsByColo `json:"httpResultsByColo"`
	// for HTTP, the method to use when running the test
	Method                  string                                                  `json:"method"`
	TracerouteResults       DexTestListResponseResultTestsTracerouteResults         `json:"tracerouteResults,nullable"`
	TracerouteResultsByColo []DexTestListResponseResultTestsTracerouteResultsByColo `json:"tracerouteResultsByColo"`
	JSON                    dexTestListResponseResultTestJSON                       `json:"-"`
}

// dexTestListResponseResultTestJSON contains the JSON metadata for the struct
// [DexTestListResponseResultTest]
type dexTestListResponseResultTestJSON struct {
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

func (r *DexTestListResponseResultTest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// test type, http or traceroute
type DexTestListResponseResultTestsKind string

const (
	DexTestListResponseResultTestsKindHTTP       DexTestListResponseResultTestsKind = "http"
	DexTestListResponseResultTestsKindTraceroute DexTestListResponseResultTestsKind = "traceroute"
)

type DexTestListResponseResultTestsHTTPResults struct {
	ResourceFetchTime DexTestListResponseResultTestsHTTPResultsResourceFetchTime `json:"resourceFetchTime,required"`
	JSON              dexTestListResponseResultTestsHTTPResultsJSON              `json:"-"`
}

// dexTestListResponseResultTestsHTTPResultsJSON contains the JSON metadata for the
// struct [DexTestListResponseResultTestsHTTPResults]
type dexTestListResponseResultTestsHTTPResultsJSON struct {
	ResourceFetchTime apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DexTestListResponseResultTestsHTTPResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseResultTestsHTTPResultsResourceFetchTime struct {
	History  []DexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistory `json:"history,required"`
	AvgMs    int64                                                               `json:"avgMs,nullable"`
	OverTime DexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTime  `json:"overTime,nullable"`
	JSON     dexTestListResponseResultTestsHTTPResultsResourceFetchTimeJSON      `json:"-"`
}

// dexTestListResponseResultTestsHTTPResultsResourceFetchTimeJSON contains the JSON
// metadata for the struct
// [DexTestListResponseResultTestsHTTPResultsResourceFetchTime]
type dexTestListResponseResultTestsHTTPResultsResourceFetchTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseResultTestsHTTPResultsResourceFetchTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistory struct {
	TimePeriod DexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriod `json:"timePeriod,required"`
	AvgMs      int64                                                                       `json:"avgMs,nullable"`
	DeltaPct   float64                                                                     `json:"deltaPct,nullable"`
	JSON       dexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryJSON       `json:"-"`
}

// dexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryJSON contains
// the JSON metadata for the struct
// [DexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistory]
type dexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriod struct {
	Units DexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits `json:"units,required"`
	Value int64                                                                            `json:"value,required"`
	JSON  dexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodJSON  `json:"-"`
}

// dexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodJSON
// contains the JSON metadata for the struct
// [DexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriod]
type dexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits string

const (
	DexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnitsHours    DexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits = "hours"
	DexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnitsDays     DexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits = "days"
	DexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnitsTestRuns DexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits = "testRuns"
)

type DexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTime struct {
	TimePeriod DexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriod `json:"timePeriod,required"`
	Values     []DexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeValue    `json:"values,required"`
	JSON       dexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeJSON       `json:"-"`
}

// dexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeJSON contains
// the JSON metadata for the struct
// [DexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTime]
type dexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriod struct {
	Units DexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits `json:"units,required"`
	Value int64                                                                             `json:"value,required"`
	JSON  dexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodJSON  `json:"-"`
}

// dexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodJSON
// contains the JSON metadata for the struct
// [DexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriod]
type dexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits string

const (
	DexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnitsHours    DexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits = "hours"
	DexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnitsDays     DexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits = "days"
	DexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnitsTestRuns DexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits = "testRuns"
)

type DexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeValue struct {
	AvgMs     int64                                                                       `json:"avgMs,required"`
	Timestamp string                                                                      `json:"timestamp,required"`
	JSON      dexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeValueJSON `json:"-"`
}

// dexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeValueJSON
// contains the JSON metadata for the struct
// [DexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeValue]
type dexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseResultTestsHTTPResultsByColo struct {
	// Cloudflare colo
	Colo              string                                                           `json:"colo,required"`
	ResourceFetchTime DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTime `json:"resourceFetchTime,required"`
	JSON              dexTestListResponseResultTestsHTTPResultsByColoJSON              `json:"-"`
}

// dexTestListResponseResultTestsHTTPResultsByColoJSON contains the JSON metadata
// for the struct [DexTestListResponseResultTestsHTTPResultsByColo]
type dexTestListResponseResultTestsHTTPResultsByColoJSON struct {
	Colo              apijson.Field
	ResourceFetchTime apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DexTestListResponseResultTestsHTTPResultsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTime struct {
	History  []DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistory `json:"history,required"`
	AvgMs    int64                                                                     `json:"avgMs,nullable"`
	OverTime DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTime  `json:"overTime,nullable"`
	JSON     dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeJSON      `json:"-"`
}

// dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeJSON contains
// the JSON metadata for the struct
// [DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTime]
type dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistory struct {
	TimePeriod DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriod `json:"timePeriod,required"`
	AvgMs      int64                                                                             `json:"avgMs,nullable"`
	DeltaPct   float64                                                                           `json:"deltaPct,nullable"`
	JSON       dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryJSON       `json:"-"`
}

// dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryJSON
// contains the JSON metadata for the struct
// [DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistory]
type dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriod struct {
	Units DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits `json:"units,required"`
	Value int64                                                                                  `json:"value,required"`
	JSON  dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodJSON  `json:"-"`
}

// dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodJSON
// contains the JSON metadata for the struct
// [DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriod]
type dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits string

const (
	DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnitsHours    DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits = "hours"
	DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnitsDays     DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits = "days"
	DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnitsTestRuns DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits = "testRuns"
)

type DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTime struct {
	TimePeriod DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriod `json:"timePeriod,required"`
	Values     []DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeValue    `json:"values,required"`
	JSON       dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeJSON       `json:"-"`
}

// dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeJSON
// contains the JSON metadata for the struct
// [DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTime]
type dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriod struct {
	Units DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits `json:"units,required"`
	Value int64                                                                                   `json:"value,required"`
	JSON  dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodJSON  `json:"-"`
}

// dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodJSON
// contains the JSON metadata for the struct
// [DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriod]
type dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits string

const (
	DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnitsHours    DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits = "hours"
	DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnitsDays     DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits = "days"
	DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnitsTestRuns DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits = "testRuns"
)

type DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeValue struct {
	AvgMs     int64                                                                             `json:"avgMs,required"`
	Timestamp string                                                                            `json:"timestamp,required"`
	JSON      dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeValueJSON `json:"-"`
}

// dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeValueJSON
// contains the JSON metadata for the struct
// [DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeValue]
type dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseResultTestsTracerouteResults struct {
	RoundTripTime DexTestListResponseResultTestsTracerouteResultsRoundTripTime `json:"roundTripTime,required"`
	JSON          dexTestListResponseResultTestsTracerouteResultsJSON          `json:"-"`
}

// dexTestListResponseResultTestsTracerouteResultsJSON contains the JSON metadata
// for the struct [DexTestListResponseResultTestsTracerouteResults]
type dexTestListResponseResultTestsTracerouteResultsJSON struct {
	RoundTripTime apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DexTestListResponseResultTestsTracerouteResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseResultTestsTracerouteResultsRoundTripTime struct {
	History  []DexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistory `json:"history,required"`
	AvgMs    int64                                                                 `json:"avgMs,nullable"`
	OverTime DexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTime  `json:"overTime,nullable"`
	JSON     dexTestListResponseResultTestsTracerouteResultsRoundTripTimeJSON      `json:"-"`
}

// dexTestListResponseResultTestsTracerouteResultsRoundTripTimeJSON contains the
// JSON metadata for the struct
// [DexTestListResponseResultTestsTracerouteResultsRoundTripTime]
type dexTestListResponseResultTestsTracerouteResultsRoundTripTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseResultTestsTracerouteResultsRoundTripTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistory struct {
	TimePeriod DexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriod `json:"timePeriod,required"`
	AvgMs      int64                                                                         `json:"avgMs,nullable"`
	DeltaPct   float64                                                                       `json:"deltaPct,nullable"`
	JSON       dexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryJSON       `json:"-"`
}

// dexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryJSON contains
// the JSON metadata for the struct
// [DexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistory]
type dexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriod struct {
	Units DexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits `json:"units,required"`
	Value int64                                                                              `json:"value,required"`
	JSON  dexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodJSON  `json:"-"`
}

// dexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodJSON
// contains the JSON metadata for the struct
// [DexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriod]
type dexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits string

const (
	DexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnitsHours    DexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits = "hours"
	DexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnitsDays     DexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits = "days"
	DexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnitsTestRuns DexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits = "testRuns"
)

type DexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTime struct {
	TimePeriod DexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriod `json:"timePeriod,required"`
	Values     []DexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeValue    `json:"values,required"`
	JSON       dexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeJSON       `json:"-"`
}

// dexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeJSON
// contains the JSON metadata for the struct
// [DexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTime]
type dexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriod struct {
	Units DexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits `json:"units,required"`
	Value int64                                                                               `json:"value,required"`
	JSON  dexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodJSON  `json:"-"`
}

// dexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodJSON
// contains the JSON metadata for the struct
// [DexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriod]
type dexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits string

const (
	DexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnitsHours    DexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits = "hours"
	DexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnitsDays     DexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits = "days"
	DexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnitsTestRuns DexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits = "testRuns"
)

type DexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeValue struct {
	AvgMs     int64                                                                         `json:"avgMs,required"`
	Timestamp string                                                                        `json:"timestamp,required"`
	JSON      dexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeValueJSON `json:"-"`
}

// dexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeValueJSON
// contains the JSON metadata for the struct
// [DexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeValue]
type dexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseResultTestsTracerouteResultsByColo struct {
	// Cloudflare colo
	Colo          string                                                             `json:"colo,required"`
	RoundTripTime DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTime `json:"roundTripTime,required"`
	JSON          dexTestListResponseResultTestsTracerouteResultsByColoJSON          `json:"-"`
}

// dexTestListResponseResultTestsTracerouteResultsByColoJSON contains the JSON
// metadata for the struct [DexTestListResponseResultTestsTracerouteResultsByColo]
type dexTestListResponseResultTestsTracerouteResultsByColoJSON struct {
	Colo          apijson.Field
	RoundTripTime apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DexTestListResponseResultTestsTracerouteResultsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTime struct {
	History  []DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistory `json:"history,required"`
	AvgMs    int64                                                                       `json:"avgMs,nullable"`
	OverTime DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTime  `json:"overTime,nullable"`
	JSON     dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeJSON      `json:"-"`
}

// dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeJSON contains
// the JSON metadata for the struct
// [DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTime]
type dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistory struct {
	TimePeriod DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriod `json:"timePeriod,required"`
	AvgMs      int64                                                                               `json:"avgMs,nullable"`
	DeltaPct   float64                                                                             `json:"deltaPct,nullable"`
	JSON       dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryJSON       `json:"-"`
}

// dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryJSON
// contains the JSON metadata for the struct
// [DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistory]
type dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriod struct {
	Units DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits `json:"units,required"`
	Value int64                                                                                    `json:"value,required"`
	JSON  dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodJSON  `json:"-"`
}

// dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodJSON
// contains the JSON metadata for the struct
// [DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriod]
type dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits string

const (
	DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnitsHours    DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits = "hours"
	DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnitsDays     DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits = "days"
	DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnitsTestRuns DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits = "testRuns"
)

type DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTime struct {
	TimePeriod DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriod `json:"timePeriod,required"`
	Values     []DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeValue    `json:"values,required"`
	JSON       dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeJSON       `json:"-"`
}

// dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeJSON
// contains the JSON metadata for the struct
// [DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTime]
type dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriod struct {
	Units DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits `json:"units,required"`
	Value int64                                                                                     `json:"value,required"`
	JSON  dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodJSON  `json:"-"`
}

// dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodJSON
// contains the JSON metadata for the struct
// [DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriod]
type dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits string

const (
	DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnitsHours    DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits = "hours"
	DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnitsDays     DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits = "days"
	DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnitsTestRuns DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits = "testRuns"
)

type DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeValue struct {
	AvgMs     int64                                                                               `json:"avgMs,required"`
	Timestamp string                                                                              `json:"timestamp,required"`
	JSON      dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeValueJSON `json:"-"`
}

// dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeValueJSON
// contains the JSON metadata for the struct
// [DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeValue]
type dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                           `json:"total_count"`
	JSON       dexTestListResponseResultInfoJSON `json:"-"`
}

// dexTestListResponseResultInfoJSON contains the JSON metadata for the struct
// [DexTestListResponseResultInfo]
type dexTestListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DexTestListResponseSuccess bool

const (
	DexTestListResponseSuccessTrue DexTestListResponseSuccess = true
)

type DexTestListParams struct {
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

// URLQuery serializes [DexTestListParams]'s query parameters as `url.Values`.
func (r DexTestListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
