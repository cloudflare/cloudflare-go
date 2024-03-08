// File generated from our OpenAPI spec by Stainless.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *DEXTestService) List(ctx context.Context, params DEXTestListParams, opts ...option.RequestOption) (res *shared.V4PagePagination[DEXTestListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/dex/tests", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List DEX tests
func (r *DEXTestService) ListAutoPaging(ctx context.Context, params DEXTestListParams, opts ...option.RequestOption) *shared.V4PagePaginationAutoPager[DEXTestListResponse] {
	return shared.NewV4PagePaginationAutoPager(r.List(ctx, params, opts...))
}

type DigitalExperienceMonitoringTests struct {
	OverviewMetrics DigitalExperienceMonitoringTestsOverviewMetrics `json:"overviewMetrics,required"`
	// array of test results objects.
	Tests []DigitalExperienceMonitoringTestsTest `json:"tests,required"`
	JSON  digitalExperienceMonitoringTestsJSON   `json:"-"`
}

// digitalExperienceMonitoringTestsJSON contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTests]
type digitalExperienceMonitoringTestsJSON struct {
	OverviewMetrics apijson.Field
	Tests           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTests) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTestsJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTestsOverviewMetrics struct {
	// number of tests.
	TestsTotal int64 `json:"testsTotal,required"`
	// percentage availability for all traceroutes results in response
	AvgTracerouteAvailabilityPct float64                                             `json:"avgTracerouteAvailabilityPct,nullable"`
	JSON                         digitalExperienceMonitoringTestsOverviewMetricsJSON `json:"-"`
}

// digitalExperienceMonitoringTestsOverviewMetricsJSON contains the JSON metadata
// for the struct [DigitalExperienceMonitoringTestsOverviewMetrics]
type digitalExperienceMonitoringTestsOverviewMetricsJSON struct {
	TestsTotal                   apijson.Field
	AvgTracerouteAvailabilityPct apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTestsOverviewMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTestsOverviewMetricsJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTestsTest struct {
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
	Kind DigitalExperienceMonitoringTestsTestsKind `json:"kind,required"`
	// name given to this test
	Name              string                                                   `json:"name,required"`
	Updated           string                                                   `json:"updated,required"`
	HTTPResults       DigitalExperienceMonitoringTestsTestsHTTPResults         `json:"httpResults,nullable"`
	HTTPResultsByColo []DigitalExperienceMonitoringTestsTestsHTTPResultsByColo `json:"httpResultsByColo"`
	// for HTTP, the method to use when running the test
	Method                  string                                                         `json:"method"`
	TracerouteResults       DigitalExperienceMonitoringTestsTestsTracerouteResults         `json:"tracerouteResults,nullable"`
	TracerouteResultsByColo []DigitalExperienceMonitoringTestsTestsTracerouteResultsByColo `json:"tracerouteResultsByColo"`
	JSON                    digitalExperienceMonitoringTestsTestJSON                       `json:"-"`
}

// digitalExperienceMonitoringTestsTestJSON contains the JSON metadata for the
// struct [DigitalExperienceMonitoringTestsTest]
type digitalExperienceMonitoringTestsTestJSON struct {
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

func (r *DigitalExperienceMonitoringTestsTest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTestsTestJSON) RawJSON() string {
	return r.raw
}

// test type, http or traceroute
type DigitalExperienceMonitoringTestsTestsKind string

const (
	DigitalExperienceMonitoringTestsTestsKindHTTP       DigitalExperienceMonitoringTestsTestsKind = "http"
	DigitalExperienceMonitoringTestsTestsKindTraceroute DigitalExperienceMonitoringTestsTestsKind = "traceroute"
)

type DigitalExperienceMonitoringTestsTestsHTTPResults struct {
	ResourceFetchTime DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTime `json:"resourceFetchTime,required"`
	JSON              digitalExperienceMonitoringTestsTestsHTTPResultsJSON              `json:"-"`
}

// digitalExperienceMonitoringTestsTestsHTTPResultsJSON contains the JSON metadata
// for the struct [DigitalExperienceMonitoringTestsTestsHTTPResults]
type digitalExperienceMonitoringTestsTestsHTTPResultsJSON struct {
	ResourceFetchTime apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTestsTestsHTTPResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTestsTestsHTTPResultsJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTime struct {
	History  []DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeHistory `json:"history,required"`
	AvgMs    int64                                                                      `json:"avgMs,nullable"`
	OverTime DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTime  `json:"overTime,nullable"`
	JSON     digitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeJSON      `json:"-"`
}

// digitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeJSON contains
// the JSON metadata for the struct
// [DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTime]
type digitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeHistory struct {
	TimePeriod DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeHistoryTimePeriod `json:"timePeriod,required"`
	AvgMs      int64                                                                              `json:"avgMs,nullable"`
	DeltaPct   float64                                                                            `json:"deltaPct,nullable"`
	JSON       digitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeHistoryJSON       `json:"-"`
}

// digitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeHistoryJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeHistory]
type digitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeHistoryJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeHistoryTimePeriod struct {
	Units DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits `json:"units,required"`
	Value int64                                                                                   `json:"value,required"`
	JSON  digitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeHistoryTimePeriodJSON  `json:"-"`
}

// digitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeHistoryTimePeriodJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeHistoryTimePeriod]
type digitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeHistoryTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeHistoryTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeHistoryTimePeriodJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits string

const (
	DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnitsHours    DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits = "hours"
	DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnitsDays     DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits = "days"
	DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnitsTestRuns DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits = "testRuns"
)

type DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTime struct {
	TimePeriod DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTimeTimePeriod `json:"timePeriod,required"`
	Values     []DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTimeValue    `json:"values,required"`
	JSON       digitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTimeJSON       `json:"-"`
}

// digitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTimeJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTime]
type digitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTimeJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTimeTimePeriod struct {
	Units DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits `json:"units,required"`
	Value int64                                                                                    `json:"value,required"`
	JSON  digitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodJSON  `json:"-"`
}

// digitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTimeTimePeriod]
type digitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTimeTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits string

const (
	DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnitsHours    DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits = "hours"
	DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnitsDays     DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits = "days"
	DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnitsTestRuns DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits = "testRuns"
)

type DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTimeValue struct {
	AvgMs     int64                                                                              `json:"avgMs,required"`
	Timestamp string                                                                             `json:"timestamp,required"`
	JSON      digitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTimeValueJSON `json:"-"`
}

// digitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTimeValueJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTimeValue]
type digitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTestsTestsHTTPResultsResourceFetchTimeOverTimeValueJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTestsTestsHTTPResultsByColo struct {
	// Cloudflare colo
	Colo              string                                                                  `json:"colo,required"`
	ResourceFetchTime DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTime `json:"resourceFetchTime,required"`
	JSON              digitalExperienceMonitoringTestsTestsHTTPResultsByColoJSON              `json:"-"`
}

// digitalExperienceMonitoringTestsTestsHTTPResultsByColoJSON contains the JSON
// metadata for the struct [DigitalExperienceMonitoringTestsTestsHTTPResultsByColo]
type digitalExperienceMonitoringTestsTestsHTTPResultsByColoJSON struct {
	Colo              apijson.Field
	ResourceFetchTime apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTestsTestsHTTPResultsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTestsTestsHTTPResultsByColoJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTime struct {
	History  []DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeHistory `json:"history,required"`
	AvgMs    int64                                                                            `json:"avgMs,nullable"`
	OverTime DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTime  `json:"overTime,nullable"`
	JSON     digitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeJSON      `json:"-"`
}

// digitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTime]
type digitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeHistory struct {
	TimePeriod DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriod `json:"timePeriod,required"`
	AvgMs      int64                                                                                    `json:"avgMs,nullable"`
	DeltaPct   float64                                                                                  `json:"deltaPct,nullable"`
	JSON       digitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeHistoryJSON       `json:"-"`
}

// digitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeHistoryJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeHistory]
type digitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeHistoryJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriod struct {
	Units DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits `json:"units,required"`
	Value int64                                                                                         `json:"value,required"`
	JSON  digitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodJSON  `json:"-"`
}

// digitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriod]
type digitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits string

const (
	DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnitsHours    DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits = "hours"
	DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnitsDays     DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits = "days"
	DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnitsTestRuns DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits = "testRuns"
)

type DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTime struct {
	TimePeriod DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriod `json:"timePeriod,required"`
	Values     []DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTimeValue    `json:"values,required"`
	JSON       digitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTimeJSON       `json:"-"`
}

// digitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTimeJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTime]
type digitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTimeJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriod struct {
	Units DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits `json:"units,required"`
	Value int64                                                                                          `json:"value,required"`
	JSON  digitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodJSON  `json:"-"`
}

// digitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriod]
type digitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits string

const (
	DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnitsHours    DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits = "hours"
	DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnitsDays     DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits = "days"
	DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnitsTestRuns DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits = "testRuns"
)

type DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTimeValue struct {
	AvgMs     int64                                                                                    `json:"avgMs,required"`
	Timestamp string                                                                                   `json:"timestamp,required"`
	JSON      digitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTimeValueJSON `json:"-"`
}

// digitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTimeValueJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTimeValue]
type digitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTestsTestsHTTPResultsByColoResourceFetchTimeOverTimeValueJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTestsTestsTracerouteResults struct {
	RoundTripTime DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTime `json:"roundTripTime,required"`
	JSON          digitalExperienceMonitoringTestsTestsTracerouteResultsJSON          `json:"-"`
}

// digitalExperienceMonitoringTestsTestsTracerouteResultsJSON contains the JSON
// metadata for the struct [DigitalExperienceMonitoringTestsTestsTracerouteResults]
type digitalExperienceMonitoringTestsTestsTracerouteResultsJSON struct {
	RoundTripTime apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTestsTestsTracerouteResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTestsTestsTracerouteResultsJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTime struct {
	History  []DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeHistory `json:"history,required"`
	AvgMs    int64                                                                        `json:"avgMs,nullable"`
	OverTime DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTime  `json:"overTime,nullable"`
	JSON     digitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeJSON      `json:"-"`
}

// digitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeJSON contains
// the JSON metadata for the struct
// [DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTime]
type digitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeHistory struct {
	TimePeriod DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeHistoryTimePeriod `json:"timePeriod,required"`
	AvgMs      int64                                                                                `json:"avgMs,nullable"`
	DeltaPct   float64                                                                              `json:"deltaPct,nullable"`
	JSON       digitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeHistoryJSON       `json:"-"`
}

// digitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeHistoryJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeHistory]
type digitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeHistoryJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeHistoryTimePeriod struct {
	Units DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits `json:"units,required"`
	Value int64                                                                                     `json:"value,required"`
	JSON  digitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeHistoryTimePeriodJSON  `json:"-"`
}

// digitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeHistoryTimePeriodJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeHistoryTimePeriod]
type digitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeHistoryTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeHistoryTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeHistoryTimePeriodJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits string

const (
	DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnitsHours    DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits = "hours"
	DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnitsDays     DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits = "days"
	DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnitsTestRuns DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits = "testRuns"
)

type DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTime struct {
	TimePeriod DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTimeTimePeriod `json:"timePeriod,required"`
	Values     []DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTimeValue    `json:"values,required"`
	JSON       digitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTimeJSON       `json:"-"`
}

// digitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTimeJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTime]
type digitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTimeJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTimeTimePeriod struct {
	Units DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits `json:"units,required"`
	Value int64                                                                                      `json:"value,required"`
	JSON  digitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodJSON  `json:"-"`
}

// digitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTimeTimePeriod]
type digitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTimeTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits string

const (
	DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnitsHours    DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits = "hours"
	DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnitsDays     DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits = "days"
	DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnitsTestRuns DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits = "testRuns"
)

type DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTimeValue struct {
	AvgMs     int64                                                                                `json:"avgMs,required"`
	Timestamp string                                                                               `json:"timestamp,required"`
	JSON      digitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTimeValueJSON `json:"-"`
}

// digitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTimeValueJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTimeValue]
type digitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTestsTestsTracerouteResultsRoundTripTimeOverTimeValueJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTestsTestsTracerouteResultsByColo struct {
	// Cloudflare colo
	Colo          string                                                                    `json:"colo,required"`
	RoundTripTime DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTime `json:"roundTripTime,required"`
	JSON          digitalExperienceMonitoringTestsTestsTracerouteResultsByColoJSON          `json:"-"`
}

// digitalExperienceMonitoringTestsTestsTracerouteResultsByColoJSON contains the
// JSON metadata for the struct
// [DigitalExperienceMonitoringTestsTestsTracerouteResultsByColo]
type digitalExperienceMonitoringTestsTestsTracerouteResultsByColoJSON struct {
	Colo          apijson.Field
	RoundTripTime apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTestsTestsTracerouteResultsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTestsTestsTracerouteResultsByColoJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTime struct {
	History  []DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeHistory `json:"history,required"`
	AvgMs    int64                                                                              `json:"avgMs,nullable"`
	OverTime DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTime  `json:"overTime,nullable"`
	JSON     digitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeJSON      `json:"-"`
}

// digitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTime]
type digitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeHistory struct {
	TimePeriod DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriod `json:"timePeriod,required"`
	AvgMs      int64                                                                                      `json:"avgMs,nullable"`
	DeltaPct   float64                                                                                    `json:"deltaPct,nullable"`
	JSON       digitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeHistoryJSON       `json:"-"`
}

// digitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeHistoryJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeHistory]
type digitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeHistoryJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriod struct {
	Units DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits `json:"units,required"`
	Value int64                                                                                           `json:"value,required"`
	JSON  digitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodJSON  `json:"-"`
}

// digitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriod]
type digitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits string

const (
	DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnitsHours    DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits = "hours"
	DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnitsDays     DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits = "days"
	DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnitsTestRuns DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits = "testRuns"
)

type DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTime struct {
	TimePeriod DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriod `json:"timePeriod,required"`
	Values     []DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTimeValue    `json:"values,required"`
	JSON       digitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTimeJSON       `json:"-"`
}

// digitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTimeJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTime]
type digitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTimeJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriod struct {
	Units DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits `json:"units,required"`
	Value int64                                                                                            `json:"value,required"`
	JSON  digitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodJSON  `json:"-"`
}

// digitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriod]
type digitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits string

const (
	DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnitsHours    DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits = "hours"
	DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnitsDays     DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits = "days"
	DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnitsTestRuns DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits = "testRuns"
)

type DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTimeValue struct {
	AvgMs     int64                                                                                      `json:"avgMs,required"`
	Timestamp string                                                                                     `json:"timestamp,required"`
	JSON      digitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTimeValueJSON `json:"-"`
}

// digitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTimeValueJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTimeValue]
type digitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTestsTestsTracerouteResultsByColoRoundTripTimeOverTimeValueJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponse struct {
	Errors   []DEXTestListResponseError       `json:"errors,required"`
	Messages []DEXTestListResponseMessage     `json:"messages,required"`
	Result   DigitalExperienceMonitoringTests `json:"result,required"`
	// Whether the API call was successful
	Success    DEXTestListResponseSuccess    `json:"success,required"`
	ResultInfo DEXTestListResponseResultInfo `json:"result_info"`
	JSON       dexTestListResponseJSON       `json:"-"`
}

// dexTestListResponseJSON contains the JSON metadata for the struct
// [DEXTestListResponse]
type dexTestListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponseError struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    dexTestListResponseErrorJSON `json:"-"`
}

// dexTestListResponseErrorJSON contains the JSON metadata for the struct
// [DEXTestListResponseError]
type dexTestListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponseMessage struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    dexTestListResponseMessageJSON `json:"-"`
}

// dexTestListResponseMessageJSON contains the JSON metadata for the struct
// [DEXTestListResponseMessage]
type dexTestListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseMessageJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DEXTestListResponseSuccess bool

const (
	DEXTestListResponseSuccessTrue DEXTestListResponseSuccess = true
)

type DEXTestListResponseResultInfo struct {
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
// [DEXTestListResponseResultInfo]
type dexTestListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type DEXTestListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
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
