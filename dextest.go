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

// DEXTestService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewDEXTestService] method instead.
type DEXTestService struct {
	Options       []option.RequestOption
	UniqueDevices *DEXTestUniqueDeviceService
}

// NewDEXTestService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDEXTestService(opts ...option.RequestOption) (r *DEXTestService) {
	r = &DEXTestService{}
	r.Options = opts
	r.UniqueDevices = NewDEXTestUniqueDeviceService(opts...)
	return
}

// List DEX tests
func (r *DEXTestService) List(ctx context.Context, accountID string, query DEXTestListParams, opts ...option.RequestOption) (res *DEXTestListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DEXTestListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/tests", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DEXTestListResponse struct {
	OverviewMetrics DEXTestListResponseOverviewMetrics `json:"overviewMetrics,required"`
	// array of test results objects.
	Tests []DEXTestListResponseTest `json:"tests,required"`
	JSON  dexTestListResponseJSON   `json:"-"`
}

// dexTestListResponseJSON contains the JSON metadata for the struct
// [DEXTestListResponse]
type dexTestListResponseJSON struct {
	OverviewMetrics apijson.Field
	Tests           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DEXTestListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListResponseOverviewMetrics struct {
	// number of tests.
	TestsTotal int64 `json:"testsTotal,required"`
	// percentage availability for all traceroutes results in response
	AvgTracerouteAvailabilityPct float64                                `json:"avgTracerouteAvailabilityPct,nullable"`
	JSON                         dexTestListResponseOverviewMetricsJSON `json:"-"`
}

// dexTestListResponseOverviewMetricsJSON contains the JSON metadata for the struct
// [DEXTestListResponseOverviewMetrics]
type dexTestListResponseOverviewMetricsJSON struct {
	TestsTotal                   apijson.Field
	AvgTracerouteAvailabilityPct apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *DEXTestListResponseOverviewMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListResponseTest struct {
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
	Kind DEXTestListResponseTestsKind `json:"kind,required"`
	// name given to this test
	Name              string                                      `json:"name,required"`
	Updated           string                                      `json:"updated,required"`
	HTTPResults       DEXTestListResponseTestsHTTPResults         `json:"httpResults,nullable"`
	HTTPResultsByColo []DEXTestListResponseTestsHTTPResultsByColo `json:"httpResultsByColo"`
	// for HTTP, the method to use when running the test
	Method                  string                                            `json:"method"`
	TracerouteResults       DEXTestListResponseTestsTracerouteResults         `json:"tracerouteResults,nullable"`
	TracerouteResultsByColo []DEXTestListResponseTestsTracerouteResultsByColo `json:"tracerouteResultsByColo"`
	JSON                    dexTestListResponseTestJSON                       `json:"-"`
}

// dexTestListResponseTestJSON contains the JSON metadata for the struct
// [DEXTestListResponseTest]
type dexTestListResponseTestJSON struct {
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

func (r *DEXTestListResponseTest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// test type, http or traceroute
type DEXTestListResponseTestsKind string

const (
	DEXTestListResponseTestsKindHTTP       DEXTestListResponseTestsKind = "http"
	DEXTestListResponseTestsKindTraceroute DEXTestListResponseTestsKind = "traceroute"
)

type DEXTestListResponseTestsHTTPResults struct {
	ResourceFetchTime DEXTestListResponseTestsHTTPResultsResourceFetchTime `json:"resourceFetchTime,required"`
	JSON              dexTestListResponseTestsHTTPResultsJSON              `json:"-"`
}

// dexTestListResponseTestsHTTPResultsJSON contains the JSON metadata for the
// struct [DEXTestListResponseTestsHTTPResults]
type dexTestListResponseTestsHTTPResultsJSON struct {
	ResourceFetchTime apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DEXTestListResponseTestsHTTPResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListResponseTestsHTTPResultsResourceFetchTime struct {
	History  []DEXTestListResponseTestsHTTPResultsResourceFetchTimeHistory `json:"history,required"`
	AvgMs    int64                                                         `json:"avgMs,nullable"`
	OverTime DEXTestListResponseTestsHTTPResultsResourceFetchTimeOverTime  `json:"overTime,nullable"`
	JSON     dexTestListResponseTestsHTTPResultsResourceFetchTimeJSON      `json:"-"`
}

// dexTestListResponseTestsHTTPResultsResourceFetchTimeJSON contains the JSON
// metadata for the struct [DEXTestListResponseTestsHTTPResultsResourceFetchTime]
type dexTestListResponseTestsHTTPResultsResourceFetchTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseTestsHTTPResultsResourceFetchTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListResponseTestsHTTPResultsResourceFetchTimeHistory struct {
	TimePeriod DEXTestListResponseTestsHTTPResultsResourceFetchTimeHistoryTimePeriod `json:"timePeriod,required"`
	AvgMs      int64                                                                 `json:"avgMs,nullable"`
	DeltaPct   float64                                                               `json:"deltaPct,nullable"`
	JSON       dexTestListResponseTestsHTTPResultsResourceFetchTimeHistoryJSON       `json:"-"`
}

// dexTestListResponseTestsHTTPResultsResourceFetchTimeHistoryJSON contains the
// JSON metadata for the struct
// [DEXTestListResponseTestsHTTPResultsResourceFetchTimeHistory]
type dexTestListResponseTestsHTTPResultsResourceFetchTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseTestsHTTPResultsResourceFetchTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListResponseTestsHTTPResultsResourceFetchTimeHistoryTimePeriod struct {
	Units DEXTestListResponseTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits `json:"units,required"`
	Value int64                                                                      `json:"value,required"`
	JSON  dexTestListResponseTestsHTTPResultsResourceFetchTimeHistoryTimePeriodJSON  `json:"-"`
}

// dexTestListResponseTestsHTTPResultsResourceFetchTimeHistoryTimePeriodJSON
// contains the JSON metadata for the struct
// [DEXTestListResponseTestsHTTPResultsResourceFetchTimeHistoryTimePeriod]
type dexTestListResponseTestsHTTPResultsResourceFetchTimeHistoryTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseTestsHTTPResultsResourceFetchTimeHistoryTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListResponseTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits string

const (
	DEXTestListResponseTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnitsHours    DEXTestListResponseTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits = "hours"
	DEXTestListResponseTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnitsDays     DEXTestListResponseTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits = "days"
	DEXTestListResponseTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnitsTestRuns DEXTestListResponseTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits = "testRuns"
)

type DEXTestListResponseTestsHTTPResultsResourceFetchTimeOverTime struct {
	TimePeriod DEXTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeTimePeriod `json:"timePeriod,required"`
	Values     []DEXTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeValue    `json:"values,required"`
	JSON       dexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeJSON       `json:"-"`
}

// dexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeJSON contains the
// JSON metadata for the struct
// [DEXTestListResponseTestsHTTPResultsResourceFetchTimeOverTime]
type dexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseTestsHTTPResultsResourceFetchTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeTimePeriod struct {
	Units DEXTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits `json:"units,required"`
	Value int64                                                                       `json:"value,required"`
	JSON  dexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodJSON  `json:"-"`
}

// dexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodJSON
// contains the JSON metadata for the struct
// [DEXTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeTimePeriod]
type dexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits string

const (
	DEXTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnitsHours    DEXTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits = "hours"
	DEXTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnitsDays     DEXTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits = "days"
	DEXTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnitsTestRuns DEXTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits = "testRuns"
)

type DEXTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeValue struct {
	AvgMs     int64                                                                 `json:"avgMs,required"`
	Timestamp string                                                                `json:"timestamp,required"`
	JSON      dexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeValueJSON `json:"-"`
}

// dexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeValueJSON contains
// the JSON metadata for the struct
// [DEXTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeValue]
type dexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListResponseTestsHTTPResultsByColo struct {
	// Cloudflare colo
	Colo              string                                                     `json:"colo,required"`
	ResourceFetchTime DEXTestListResponseTestsHTTPResultsByColoResourceFetchTime `json:"resourceFetchTime,required"`
	JSON              dexTestListResponseTestsHTTPResultsByColoJSON              `json:"-"`
}

// dexTestListResponseTestsHTTPResultsByColoJSON contains the JSON metadata for the
// struct [DEXTestListResponseTestsHTTPResultsByColo]
type dexTestListResponseTestsHTTPResultsByColoJSON struct {
	Colo              apijson.Field
	ResourceFetchTime apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DEXTestListResponseTestsHTTPResultsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListResponseTestsHTTPResultsByColoResourceFetchTime struct {
	History  []DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistory `json:"history,required"`
	AvgMs    int64                                                               `json:"avgMs,nullable"`
	OverTime DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTime  `json:"overTime,nullable"`
	JSON     dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeJSON      `json:"-"`
}

// dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeJSON contains the JSON
// metadata for the struct
// [DEXTestListResponseTestsHTTPResultsByColoResourceFetchTime]
type dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseTestsHTTPResultsByColoResourceFetchTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistory struct {
	TimePeriod DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriod `json:"timePeriod,required"`
	AvgMs      int64                                                                       `json:"avgMs,nullable"`
	DeltaPct   float64                                                                     `json:"deltaPct,nullable"`
	JSON       dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryJSON       `json:"-"`
}

// dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryJSON contains
// the JSON metadata for the struct
// [DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistory]
type dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriod struct {
	Units DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits `json:"units,required"`
	Value int64                                                                            `json:"value,required"`
	JSON  dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodJSON  `json:"-"`
}

// dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodJSON
// contains the JSON metadata for the struct
// [DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriod]
type dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits string

const (
	DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnitsHours    DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits = "hours"
	DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnitsDays     DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits = "days"
	DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnitsTestRuns DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits = "testRuns"
)

type DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTime struct {
	TimePeriod DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriod `json:"timePeriod,required"`
	Values     []DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeValue    `json:"values,required"`
	JSON       dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeJSON       `json:"-"`
}

// dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeJSON contains
// the JSON metadata for the struct
// [DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTime]
type dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriod struct {
	Units DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits `json:"units,required"`
	Value int64                                                                             `json:"value,required"`
	JSON  dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodJSON  `json:"-"`
}

// dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodJSON
// contains the JSON metadata for the struct
// [DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriod]
type dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits string

const (
	DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnitsHours    DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits = "hours"
	DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnitsDays     DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits = "days"
	DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnitsTestRuns DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits = "testRuns"
)

type DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeValue struct {
	AvgMs     int64                                                                       `json:"avgMs,required"`
	Timestamp string                                                                      `json:"timestamp,required"`
	JSON      dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeValueJSON `json:"-"`
}

// dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeValueJSON
// contains the JSON metadata for the struct
// [DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeValue]
type dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListResponseTestsTracerouteResults struct {
	RoundTripTime DEXTestListResponseTestsTracerouteResultsRoundTripTime `json:"roundTripTime,required"`
	JSON          dexTestListResponseTestsTracerouteResultsJSON          `json:"-"`
}

// dexTestListResponseTestsTracerouteResultsJSON contains the JSON metadata for the
// struct [DEXTestListResponseTestsTracerouteResults]
type dexTestListResponseTestsTracerouteResultsJSON struct {
	RoundTripTime apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DEXTestListResponseTestsTracerouteResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListResponseTestsTracerouteResultsRoundTripTime struct {
	History  []DEXTestListResponseTestsTracerouteResultsRoundTripTimeHistory `json:"history,required"`
	AvgMs    int64                                                           `json:"avgMs,nullable"`
	OverTime DEXTestListResponseTestsTracerouteResultsRoundTripTimeOverTime  `json:"overTime,nullable"`
	JSON     dexTestListResponseTestsTracerouteResultsRoundTripTimeJSON      `json:"-"`
}

// dexTestListResponseTestsTracerouteResultsRoundTripTimeJSON contains the JSON
// metadata for the struct [DEXTestListResponseTestsTracerouteResultsRoundTripTime]
type dexTestListResponseTestsTracerouteResultsRoundTripTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseTestsTracerouteResultsRoundTripTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListResponseTestsTracerouteResultsRoundTripTimeHistory struct {
	TimePeriod DEXTestListResponseTestsTracerouteResultsRoundTripTimeHistoryTimePeriod `json:"timePeriod,required"`
	AvgMs      int64                                                                   `json:"avgMs,nullable"`
	DeltaPct   float64                                                                 `json:"deltaPct,nullable"`
	JSON       dexTestListResponseTestsTracerouteResultsRoundTripTimeHistoryJSON       `json:"-"`
}

// dexTestListResponseTestsTracerouteResultsRoundTripTimeHistoryJSON contains the
// JSON metadata for the struct
// [DEXTestListResponseTestsTracerouteResultsRoundTripTimeHistory]
type dexTestListResponseTestsTracerouteResultsRoundTripTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseTestsTracerouteResultsRoundTripTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListResponseTestsTracerouteResultsRoundTripTimeHistoryTimePeriod struct {
	Units DEXTestListResponseTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits `json:"units,required"`
	Value int64                                                                        `json:"value,required"`
	JSON  dexTestListResponseTestsTracerouteResultsRoundTripTimeHistoryTimePeriodJSON  `json:"-"`
}

// dexTestListResponseTestsTracerouteResultsRoundTripTimeHistoryTimePeriodJSON
// contains the JSON metadata for the struct
// [DEXTestListResponseTestsTracerouteResultsRoundTripTimeHistoryTimePeriod]
type dexTestListResponseTestsTracerouteResultsRoundTripTimeHistoryTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseTestsTracerouteResultsRoundTripTimeHistoryTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListResponseTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits string

const (
	DEXTestListResponseTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnitsHours    DEXTestListResponseTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits = "hours"
	DEXTestListResponseTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnitsDays     DEXTestListResponseTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits = "days"
	DEXTestListResponseTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnitsTestRuns DEXTestListResponseTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits = "testRuns"
)

type DEXTestListResponseTestsTracerouteResultsRoundTripTimeOverTime struct {
	TimePeriod DEXTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeTimePeriod `json:"timePeriod,required"`
	Values     []DEXTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeValue    `json:"values,required"`
	JSON       dexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeJSON       `json:"-"`
}

// dexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeJSON contains the
// JSON metadata for the struct
// [DEXTestListResponseTestsTracerouteResultsRoundTripTimeOverTime]
type dexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseTestsTracerouteResultsRoundTripTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeTimePeriod struct {
	Units DEXTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits `json:"units,required"`
	Value int64                                                                         `json:"value,required"`
	JSON  dexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodJSON  `json:"-"`
}

// dexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodJSON
// contains the JSON metadata for the struct
// [DEXTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeTimePeriod]
type dexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits string

const (
	DEXTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnitsHours    DEXTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits = "hours"
	DEXTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnitsDays     DEXTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits = "days"
	DEXTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnitsTestRuns DEXTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits = "testRuns"
)

type DEXTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeValue struct {
	AvgMs     int64                                                                   `json:"avgMs,required"`
	Timestamp string                                                                  `json:"timestamp,required"`
	JSON      dexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeValueJSON `json:"-"`
}

// dexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeValueJSON contains
// the JSON metadata for the struct
// [DEXTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeValue]
type dexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListResponseTestsTracerouteResultsByColo struct {
	// Cloudflare colo
	Colo          string                                                       `json:"colo,required"`
	RoundTripTime DEXTestListResponseTestsTracerouteResultsByColoRoundTripTime `json:"roundTripTime,required"`
	JSON          dexTestListResponseTestsTracerouteResultsByColoJSON          `json:"-"`
}

// dexTestListResponseTestsTracerouteResultsByColoJSON contains the JSON metadata
// for the struct [DEXTestListResponseTestsTracerouteResultsByColo]
type dexTestListResponseTestsTracerouteResultsByColoJSON struct {
	Colo          apijson.Field
	RoundTripTime apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DEXTestListResponseTestsTracerouteResultsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListResponseTestsTracerouteResultsByColoRoundTripTime struct {
	History  []DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistory `json:"history,required"`
	AvgMs    int64                                                                 `json:"avgMs,nullable"`
	OverTime DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTime  `json:"overTime,nullable"`
	JSON     dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeJSON      `json:"-"`
}

// dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeJSON contains the
// JSON metadata for the struct
// [DEXTestListResponseTestsTracerouteResultsByColoRoundTripTime]
type dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseTestsTracerouteResultsByColoRoundTripTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistory struct {
	TimePeriod DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriod `json:"timePeriod,required"`
	AvgMs      int64                                                                         `json:"avgMs,nullable"`
	DeltaPct   float64                                                                       `json:"deltaPct,nullable"`
	JSON       dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryJSON       `json:"-"`
}

// dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryJSON contains
// the JSON metadata for the struct
// [DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistory]
type dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriod struct {
	Units DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits `json:"units,required"`
	Value int64                                                                              `json:"value,required"`
	JSON  dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodJSON  `json:"-"`
}

// dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodJSON
// contains the JSON metadata for the struct
// [DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriod]
type dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits string

const (
	DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnitsHours    DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits = "hours"
	DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnitsDays     DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits = "days"
	DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnitsTestRuns DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits = "testRuns"
)

type DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTime struct {
	TimePeriod DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriod `json:"timePeriod,required"`
	Values     []DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeValue    `json:"values,required"`
	JSON       dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeJSON       `json:"-"`
}

// dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeJSON
// contains the JSON metadata for the struct
// [DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTime]
type dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriod struct {
	Units DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits `json:"units,required"`
	Value int64                                                                               `json:"value,required"`
	JSON  dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodJSON  `json:"-"`
}

// dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodJSON
// contains the JSON metadata for the struct
// [DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriod]
type dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits string

const (
	DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnitsHours    DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits = "hours"
	DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnitsDays     DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits = "days"
	DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnitsTestRuns DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits = "testRuns"
)

type DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeValue struct {
	AvgMs     int64                                                                         `json:"avgMs,required"`
	Timestamp string                                                                        `json:"timestamp,required"`
	JSON      dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeValueJSON `json:"-"`
}

// dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeValueJSON
// contains the JSON metadata for the struct
// [DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeValue]
type dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListParams struct {
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

// URLQuery serializes [DEXTestListParams]'s query parameters as `url.Values`.
func (r DEXTestListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DEXTestListResponseEnvelope struct {
	Errors   []DEXTestListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DEXTestListResponseEnvelopeMessages `json:"messages,required"`
	Result   DEXTestListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success    DEXTestListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DEXTestListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dexTestListResponseEnvelopeJSON       `json:"-"`
}

// dexTestListResponseEnvelopeJSON contains the JSON metadata for the struct
// [DEXTestListResponseEnvelope]
type dexTestListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    dexTestListResponseEnvelopeErrorsJSON `json:"-"`
}

// dexTestListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DEXTestListResponseEnvelopeErrors]
type dexTestListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestListResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    dexTestListResponseEnvelopeMessagesJSON `json:"-"`
}

// dexTestListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DEXTestListResponseEnvelopeMessages]
type dexTestListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DEXTestListResponseEnvelopeSuccess bool

const (
	DEXTestListResponseEnvelopeSuccessTrue DEXTestListResponseEnvelopeSuccess = true
)

type DEXTestListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                   `json:"total_count"`
	JSON       dexTestListResponseEnvelopeResultInfoJSON `json:"-"`
}

// dexTestListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [DEXTestListResponseEnvelopeResultInfo]
type dexTestListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
