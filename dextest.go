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
	var env DexTestListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/tests", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DexTestListResponse struct {
	OverviewMetrics DexTestListResponseOverviewMetrics `json:"overviewMetrics,required"`
	// array of test results objects.
	Tests []DexTestListResponseTest `json:"tests,required"`
	JSON  dexTestListResponseJSON   `json:"-"`
}

// dexTestListResponseJSON contains the JSON metadata for the struct
// [DexTestListResponse]
type dexTestListResponseJSON struct {
	OverviewMetrics apijson.Field
	Tests           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DexTestListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseOverviewMetrics struct {
	// number of tests.
	TestsTotal int64 `json:"testsTotal,required"`
	// percentage availability for all traceroutes results in response
	AvgTracerouteAvailabilityPct float64                                `json:"avgTracerouteAvailabilityPct,nullable"`
	JSON                         dexTestListResponseOverviewMetricsJSON `json:"-"`
}

// dexTestListResponseOverviewMetricsJSON contains the JSON metadata for the struct
// [DexTestListResponseOverviewMetrics]
type dexTestListResponseOverviewMetricsJSON struct {
	TestsTotal                   apijson.Field
	AvgTracerouteAvailabilityPct apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *DexTestListResponseOverviewMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseTest struct {
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
	Kind DexTestListResponseTestsKind `json:"kind,required"`
	// name given to this test
	Name              string                                      `json:"name,required"`
	Updated           string                                      `json:"updated,required"`
	HTTPResults       DexTestListResponseTestsHTTPResults         `json:"httpResults,nullable"`
	HTTPResultsByColo []DexTestListResponseTestsHTTPResultsByColo `json:"httpResultsByColo"`
	// for HTTP, the method to use when running the test
	Method                  string                                            `json:"method"`
	TracerouteResults       DexTestListResponseTestsTracerouteResults         `json:"tracerouteResults,nullable"`
	TracerouteResultsByColo []DexTestListResponseTestsTracerouteResultsByColo `json:"tracerouteResultsByColo"`
	JSON                    dexTestListResponseTestJSON                       `json:"-"`
}

// dexTestListResponseTestJSON contains the JSON metadata for the struct
// [DexTestListResponseTest]
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

func (r *DexTestListResponseTest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// test type, http or traceroute
type DexTestListResponseTestsKind string

const (
	DexTestListResponseTestsKindHTTP       DexTestListResponseTestsKind = "http"
	DexTestListResponseTestsKindTraceroute DexTestListResponseTestsKind = "traceroute"
)

type DexTestListResponseTestsHTTPResults struct {
	ResourceFetchTime DexTestListResponseTestsHTTPResultsResourceFetchTime `json:"resourceFetchTime,required"`
	JSON              dexTestListResponseTestsHTTPResultsJSON              `json:"-"`
}

// dexTestListResponseTestsHTTPResultsJSON contains the JSON metadata for the
// struct [DexTestListResponseTestsHTTPResults]
type dexTestListResponseTestsHTTPResultsJSON struct {
	ResourceFetchTime apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DexTestListResponseTestsHTTPResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseTestsHTTPResultsResourceFetchTime struct {
	History  []DexTestListResponseTestsHTTPResultsResourceFetchTimeHistory `json:"history,required"`
	AvgMs    int64                                                         `json:"avgMs,nullable"`
	OverTime DexTestListResponseTestsHTTPResultsResourceFetchTimeOverTime  `json:"overTime,nullable"`
	JSON     dexTestListResponseTestsHTTPResultsResourceFetchTimeJSON      `json:"-"`
}

// dexTestListResponseTestsHTTPResultsResourceFetchTimeJSON contains the JSON
// metadata for the struct [DexTestListResponseTestsHTTPResultsResourceFetchTime]
type dexTestListResponseTestsHTTPResultsResourceFetchTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseTestsHTTPResultsResourceFetchTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseTestsHTTPResultsResourceFetchTimeHistory struct {
	TimePeriod DexTestListResponseTestsHTTPResultsResourceFetchTimeHistoryTimePeriod `json:"timePeriod,required"`
	AvgMs      int64                                                                 `json:"avgMs,nullable"`
	DeltaPct   float64                                                               `json:"deltaPct,nullable"`
	JSON       dexTestListResponseTestsHTTPResultsResourceFetchTimeHistoryJSON       `json:"-"`
}

// dexTestListResponseTestsHTTPResultsResourceFetchTimeHistoryJSON contains the
// JSON metadata for the struct
// [DexTestListResponseTestsHTTPResultsResourceFetchTimeHistory]
type dexTestListResponseTestsHTTPResultsResourceFetchTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseTestsHTTPResultsResourceFetchTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseTestsHTTPResultsResourceFetchTimeHistoryTimePeriod struct {
	Units DexTestListResponseTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits `json:"units,required"`
	Value int64                                                                      `json:"value,required"`
	JSON  dexTestListResponseTestsHTTPResultsResourceFetchTimeHistoryTimePeriodJSON  `json:"-"`
}

// dexTestListResponseTestsHTTPResultsResourceFetchTimeHistoryTimePeriodJSON
// contains the JSON metadata for the struct
// [DexTestListResponseTestsHTTPResultsResourceFetchTimeHistoryTimePeriod]
type dexTestListResponseTestsHTTPResultsResourceFetchTimeHistoryTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseTestsHTTPResultsResourceFetchTimeHistoryTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits string

const (
	DexTestListResponseTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnitsHours    DexTestListResponseTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits = "hours"
	DexTestListResponseTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnitsDays     DexTestListResponseTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits = "days"
	DexTestListResponseTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnitsTestRuns DexTestListResponseTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits = "testRuns"
)

type DexTestListResponseTestsHTTPResultsResourceFetchTimeOverTime struct {
	TimePeriod DexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeTimePeriod `json:"timePeriod,required"`
	Values     []DexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeValue    `json:"values,required"`
	JSON       dexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeJSON       `json:"-"`
}

// dexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeJSON contains the
// JSON metadata for the struct
// [DexTestListResponseTestsHTTPResultsResourceFetchTimeOverTime]
type dexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseTestsHTTPResultsResourceFetchTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeTimePeriod struct {
	Units DexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits `json:"units,required"`
	Value int64                                                                       `json:"value,required"`
	JSON  dexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodJSON  `json:"-"`
}

// dexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodJSON
// contains the JSON metadata for the struct
// [DexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeTimePeriod]
type dexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits string

const (
	DexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnitsHours    DexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits = "hours"
	DexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnitsDays     DexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits = "days"
	DexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnitsTestRuns DexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits = "testRuns"
)

type DexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeValue struct {
	AvgMs     int64                                                                 `json:"avgMs,required"`
	Timestamp string                                                                `json:"timestamp,required"`
	JSON      dexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeValueJSON `json:"-"`
}

// dexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeValueJSON contains
// the JSON metadata for the struct
// [DexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeValue]
type dexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseTestsHTTPResultsResourceFetchTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseTestsHTTPResultsByColo struct {
	// Cloudflare colo
	Colo              string                                                     `json:"colo,required"`
	ResourceFetchTime DexTestListResponseTestsHTTPResultsByColoResourceFetchTime `json:"resourceFetchTime,required"`
	JSON              dexTestListResponseTestsHTTPResultsByColoJSON              `json:"-"`
}

// dexTestListResponseTestsHTTPResultsByColoJSON contains the JSON metadata for the
// struct [DexTestListResponseTestsHTTPResultsByColo]
type dexTestListResponseTestsHTTPResultsByColoJSON struct {
	Colo              apijson.Field
	ResourceFetchTime apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DexTestListResponseTestsHTTPResultsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseTestsHTTPResultsByColoResourceFetchTime struct {
	History  []DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistory `json:"history,required"`
	AvgMs    int64                                                               `json:"avgMs,nullable"`
	OverTime DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTime  `json:"overTime,nullable"`
	JSON     dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeJSON      `json:"-"`
}

// dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeJSON contains the JSON
// metadata for the struct
// [DexTestListResponseTestsHTTPResultsByColoResourceFetchTime]
type dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseTestsHTTPResultsByColoResourceFetchTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistory struct {
	TimePeriod DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriod `json:"timePeriod,required"`
	AvgMs      int64                                                                       `json:"avgMs,nullable"`
	DeltaPct   float64                                                                     `json:"deltaPct,nullable"`
	JSON       dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryJSON       `json:"-"`
}

// dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryJSON contains
// the JSON metadata for the struct
// [DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistory]
type dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriod struct {
	Units DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits `json:"units,required"`
	Value int64                                                                            `json:"value,required"`
	JSON  dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodJSON  `json:"-"`
}

// dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodJSON
// contains the JSON metadata for the struct
// [DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriod]
type dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits string

const (
	DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnitsHours    DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits = "hours"
	DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnitsDays     DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits = "days"
	DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnitsTestRuns DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits = "testRuns"
)

type DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTime struct {
	TimePeriod DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriod `json:"timePeriod,required"`
	Values     []DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeValue    `json:"values,required"`
	JSON       dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeJSON       `json:"-"`
}

// dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeJSON contains
// the JSON metadata for the struct
// [DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTime]
type dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriod struct {
	Units DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits `json:"units,required"`
	Value int64                                                                             `json:"value,required"`
	JSON  dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodJSON  `json:"-"`
}

// dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodJSON
// contains the JSON metadata for the struct
// [DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriod]
type dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits string

const (
	DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnitsHours    DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits = "hours"
	DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnitsDays     DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits = "days"
	DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnitsTestRuns DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits = "testRuns"
)

type DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeValue struct {
	AvgMs     int64                                                                       `json:"avgMs,required"`
	Timestamp string                                                                      `json:"timestamp,required"`
	JSON      dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeValueJSON `json:"-"`
}

// dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeValueJSON
// contains the JSON metadata for the struct
// [DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeValue]
type dexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseTestsHTTPResultsByColoResourceFetchTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseTestsTracerouteResults struct {
	RoundTripTime DexTestListResponseTestsTracerouteResultsRoundTripTime `json:"roundTripTime,required"`
	JSON          dexTestListResponseTestsTracerouteResultsJSON          `json:"-"`
}

// dexTestListResponseTestsTracerouteResultsJSON contains the JSON metadata for the
// struct [DexTestListResponseTestsTracerouteResults]
type dexTestListResponseTestsTracerouteResultsJSON struct {
	RoundTripTime apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DexTestListResponseTestsTracerouteResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseTestsTracerouteResultsRoundTripTime struct {
	History  []DexTestListResponseTestsTracerouteResultsRoundTripTimeHistory `json:"history,required"`
	AvgMs    int64                                                           `json:"avgMs,nullable"`
	OverTime DexTestListResponseTestsTracerouteResultsRoundTripTimeOverTime  `json:"overTime,nullable"`
	JSON     dexTestListResponseTestsTracerouteResultsRoundTripTimeJSON      `json:"-"`
}

// dexTestListResponseTestsTracerouteResultsRoundTripTimeJSON contains the JSON
// metadata for the struct [DexTestListResponseTestsTracerouteResultsRoundTripTime]
type dexTestListResponseTestsTracerouteResultsRoundTripTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseTestsTracerouteResultsRoundTripTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseTestsTracerouteResultsRoundTripTimeHistory struct {
	TimePeriod DexTestListResponseTestsTracerouteResultsRoundTripTimeHistoryTimePeriod `json:"timePeriod,required"`
	AvgMs      int64                                                                   `json:"avgMs,nullable"`
	DeltaPct   float64                                                                 `json:"deltaPct,nullable"`
	JSON       dexTestListResponseTestsTracerouteResultsRoundTripTimeHistoryJSON       `json:"-"`
}

// dexTestListResponseTestsTracerouteResultsRoundTripTimeHistoryJSON contains the
// JSON metadata for the struct
// [DexTestListResponseTestsTracerouteResultsRoundTripTimeHistory]
type dexTestListResponseTestsTracerouteResultsRoundTripTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseTestsTracerouteResultsRoundTripTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseTestsTracerouteResultsRoundTripTimeHistoryTimePeriod struct {
	Units DexTestListResponseTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits `json:"units,required"`
	Value int64                                                                        `json:"value,required"`
	JSON  dexTestListResponseTestsTracerouteResultsRoundTripTimeHistoryTimePeriodJSON  `json:"-"`
}

// dexTestListResponseTestsTracerouteResultsRoundTripTimeHistoryTimePeriodJSON
// contains the JSON metadata for the struct
// [DexTestListResponseTestsTracerouteResultsRoundTripTimeHistoryTimePeriod]
type dexTestListResponseTestsTracerouteResultsRoundTripTimeHistoryTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseTestsTracerouteResultsRoundTripTimeHistoryTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits string

const (
	DexTestListResponseTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnitsHours    DexTestListResponseTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits = "hours"
	DexTestListResponseTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnitsDays     DexTestListResponseTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits = "days"
	DexTestListResponseTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnitsTestRuns DexTestListResponseTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits = "testRuns"
)

type DexTestListResponseTestsTracerouteResultsRoundTripTimeOverTime struct {
	TimePeriod DexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeTimePeriod `json:"timePeriod,required"`
	Values     []DexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeValue    `json:"values,required"`
	JSON       dexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeJSON       `json:"-"`
}

// dexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeJSON contains the
// JSON metadata for the struct
// [DexTestListResponseTestsTracerouteResultsRoundTripTimeOverTime]
type dexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseTestsTracerouteResultsRoundTripTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeTimePeriod struct {
	Units DexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits `json:"units,required"`
	Value int64                                                                         `json:"value,required"`
	JSON  dexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodJSON  `json:"-"`
}

// dexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodJSON
// contains the JSON metadata for the struct
// [DexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeTimePeriod]
type dexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits string

const (
	DexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnitsHours    DexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits = "hours"
	DexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnitsDays     DexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits = "days"
	DexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnitsTestRuns DexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits = "testRuns"
)

type DexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeValue struct {
	AvgMs     int64                                                                   `json:"avgMs,required"`
	Timestamp string                                                                  `json:"timestamp,required"`
	JSON      dexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeValueJSON `json:"-"`
}

// dexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeValueJSON contains
// the JSON metadata for the struct
// [DexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeValue]
type dexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseTestsTracerouteResultsRoundTripTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseTestsTracerouteResultsByColo struct {
	// Cloudflare colo
	Colo          string                                                       `json:"colo,required"`
	RoundTripTime DexTestListResponseTestsTracerouteResultsByColoRoundTripTime `json:"roundTripTime,required"`
	JSON          dexTestListResponseTestsTracerouteResultsByColoJSON          `json:"-"`
}

// dexTestListResponseTestsTracerouteResultsByColoJSON contains the JSON metadata
// for the struct [DexTestListResponseTestsTracerouteResultsByColo]
type dexTestListResponseTestsTracerouteResultsByColoJSON struct {
	Colo          apijson.Field
	RoundTripTime apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DexTestListResponseTestsTracerouteResultsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseTestsTracerouteResultsByColoRoundTripTime struct {
	History  []DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistory `json:"history,required"`
	AvgMs    int64                                                                 `json:"avgMs,nullable"`
	OverTime DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTime  `json:"overTime,nullable"`
	JSON     dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeJSON      `json:"-"`
}

// dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeJSON contains the
// JSON metadata for the struct
// [DexTestListResponseTestsTracerouteResultsByColoRoundTripTime]
type dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseTestsTracerouteResultsByColoRoundTripTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistory struct {
	TimePeriod DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriod `json:"timePeriod,required"`
	AvgMs      int64                                                                         `json:"avgMs,nullable"`
	DeltaPct   float64                                                                       `json:"deltaPct,nullable"`
	JSON       dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryJSON       `json:"-"`
}

// dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryJSON contains
// the JSON metadata for the struct
// [DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistory]
type dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriod struct {
	Units DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits `json:"units,required"`
	Value int64                                                                              `json:"value,required"`
	JSON  dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodJSON  `json:"-"`
}

// dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodJSON
// contains the JSON metadata for the struct
// [DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriod]
type dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits string

const (
	DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnitsHours    DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits = "hours"
	DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnitsDays     DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits = "days"
	DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnitsTestRuns DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits = "testRuns"
)

type DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTime struct {
	TimePeriod DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriod `json:"timePeriod,required"`
	Values     []DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeValue    `json:"values,required"`
	JSON       dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeJSON       `json:"-"`
}

// dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeJSON
// contains the JSON metadata for the struct
// [DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTime]
type dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriod struct {
	Units DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits `json:"units,required"`
	Value int64                                                                               `json:"value,required"`
	JSON  dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodJSON  `json:"-"`
}

// dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodJSON
// contains the JSON metadata for the struct
// [DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriod]
type dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits string

const (
	DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnitsHours    DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits = "hours"
	DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnitsDays     DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits = "days"
	DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnitsTestRuns DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits = "testRuns"
)

type DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeValue struct {
	AvgMs     int64                                                                         `json:"avgMs,required"`
	Timestamp string                                                                        `json:"timestamp,required"`
	JSON      dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeValueJSON `json:"-"`
}

// dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeValueJSON
// contains the JSON metadata for the struct
// [DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeValue]
type dexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseTestsTracerouteResultsByColoRoundTripTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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

type DexTestListResponseEnvelope struct {
	Errors   []DexTestListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DexTestListResponseEnvelopeMessages `json:"messages,required"`
	Result   DexTestListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success    DexTestListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DexTestListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dexTestListResponseEnvelopeJSON       `json:"-"`
}

// dexTestListResponseEnvelopeJSON contains the JSON metadata for the struct
// [DexTestListResponseEnvelope]
type dexTestListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    dexTestListResponseEnvelopeErrorsJSON `json:"-"`
}

// dexTestListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DexTestListResponseEnvelopeErrors]
type dexTestListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestListResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    dexTestListResponseEnvelopeMessagesJSON `json:"-"`
}

// dexTestListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DexTestListResponseEnvelopeMessages]
type dexTestListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DexTestListResponseEnvelopeSuccess bool

const (
	DexTestListResponseEnvelopeSuccessTrue DexTestListResponseEnvelopeSuccess = true
)

type DexTestListResponseEnvelopeResultInfo struct {
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
// struct [DexTestListResponseEnvelopeResultInfo]
type dexTestListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
