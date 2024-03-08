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

type DEXTestListResponse struct {
	Errors   []DEXTestListResponseError   `json:"errors,required"`
	Messages []DEXTestListResponseMessage `json:"messages,required"`
	Result   DEXTestListResponseResult    `json:"result,required"`
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

type DEXTestListResponseResult struct {
	OverviewMetrics DEXTestListResponseResultOverviewMetrics `json:"overviewMetrics,required"`
	// array of test results objects.
	Tests []DEXTestListResponseResultTest `json:"tests,required"`
	JSON  dexTestListResponseResultJSON   `json:"-"`
}

// dexTestListResponseResultJSON contains the JSON metadata for the struct
// [DEXTestListResponseResult]
type dexTestListResponseResultJSON struct {
	OverviewMetrics apijson.Field
	Tests           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DEXTestListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponseResultOverviewMetrics struct {
	// number of tests.
	TestsTotal int64 `json:"testsTotal,required"`
	// percentage availability for all traceroutes results in response
	AvgTracerouteAvailabilityPct float64                                      `json:"avgTracerouteAvailabilityPct,nullable"`
	JSON                         dexTestListResponseResultOverviewMetricsJSON `json:"-"`
}

// dexTestListResponseResultOverviewMetricsJSON contains the JSON metadata for the
// struct [DEXTestListResponseResultOverviewMetrics]
type dexTestListResponseResultOverviewMetricsJSON struct {
	TestsTotal                   apijson.Field
	AvgTracerouteAvailabilityPct apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *DEXTestListResponseResultOverviewMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultOverviewMetricsJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponseResultTest struct {
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
	Kind DEXTestListResponseResultTestsKind `json:"kind,required"`
	// name given to this test
	Name              string                                            `json:"name,required"`
	Updated           string                                            `json:"updated,required"`
	HTTPResults       DEXTestListResponseResultTestsHTTPResults         `json:"httpResults,nullable"`
	HTTPResultsByColo []DEXTestListResponseResultTestsHTTPResultsByColo `json:"httpResultsByColo"`
	// for HTTP, the method to use when running the test
	Method                  string                                                  `json:"method"`
	TracerouteResults       DEXTestListResponseResultTestsTracerouteResults         `json:"tracerouteResults,nullable"`
	TracerouteResultsByColo []DEXTestListResponseResultTestsTracerouteResultsByColo `json:"tracerouteResultsByColo"`
	JSON                    dexTestListResponseResultTestJSON                       `json:"-"`
}

// dexTestListResponseResultTestJSON contains the JSON metadata for the struct
// [DEXTestListResponseResultTest]
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

func (r *DEXTestListResponseResultTest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultTestJSON) RawJSON() string {
	return r.raw
}

// test type, http or traceroute
type DEXTestListResponseResultTestsKind string

const (
	DEXTestListResponseResultTestsKindHTTP       DEXTestListResponseResultTestsKind = "http"
	DEXTestListResponseResultTestsKindTraceroute DEXTestListResponseResultTestsKind = "traceroute"
)

type DEXTestListResponseResultTestsHTTPResults struct {
	ResourceFetchTime DEXTestListResponseResultTestsHTTPResultsResourceFetchTime `json:"resourceFetchTime,required"`
	JSON              dexTestListResponseResultTestsHTTPResultsJSON              `json:"-"`
}

// dexTestListResponseResultTestsHTTPResultsJSON contains the JSON metadata for the
// struct [DEXTestListResponseResultTestsHTTPResults]
type dexTestListResponseResultTestsHTTPResultsJSON struct {
	ResourceFetchTime apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DEXTestListResponseResultTestsHTTPResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultTestsHTTPResultsJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponseResultTestsHTTPResultsResourceFetchTime struct {
	History  []DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistory `json:"history,required"`
	AvgMs    int64                                                               `json:"avgMs,nullable"`
	OverTime DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTime  `json:"overTime,nullable"`
	JSON     dexTestListResponseResultTestsHTTPResultsResourceFetchTimeJSON      `json:"-"`
}

// dexTestListResponseResultTestsHTTPResultsResourceFetchTimeJSON contains the JSON
// metadata for the struct
// [DEXTestListResponseResultTestsHTTPResultsResourceFetchTime]
type dexTestListResponseResultTestsHTTPResultsResourceFetchTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseResultTestsHTTPResultsResourceFetchTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultTestsHTTPResultsResourceFetchTimeJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistory struct {
	TimePeriod DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriod `json:"timePeriod,required"`
	AvgMs      int64                                                                       `json:"avgMs,nullable"`
	DeltaPct   float64                                                                     `json:"deltaPct,nullable"`
	JSON       dexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryJSON       `json:"-"`
}

// dexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryJSON contains
// the JSON metadata for the struct
// [DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistory]
type dexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriod struct {
	Units DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits `json:"units,required"`
	Value int64                                                                            `json:"value,required"`
	JSON  dexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodJSON  `json:"-"`
}

// dexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodJSON
// contains the JSON metadata for the struct
// [DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriod]
type dexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits string

const (
	DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnitsHours    DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits = "hours"
	DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnitsDays     DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits = "days"
	DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnitsTestRuns DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits = "testRuns"
)

type DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTime struct {
	TimePeriod DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriod `json:"timePeriod,required"`
	Values     []DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeValue    `json:"values,required"`
	JSON       dexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeJSON       `json:"-"`
}

// dexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeJSON contains
// the JSON metadata for the struct
// [DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTime]
type dexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriod struct {
	Units DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits `json:"units,required"`
	Value int64                                                                             `json:"value,required"`
	JSON  dexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodJSON  `json:"-"`
}

// dexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodJSON
// contains the JSON metadata for the struct
// [DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriod]
type dexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits string

const (
	DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnitsHours    DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits = "hours"
	DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnitsDays     DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits = "days"
	DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnitsTestRuns DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits = "testRuns"
)

type DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeValue struct {
	AvgMs     int64                                                                       `json:"avgMs,required"`
	Timestamp string                                                                      `json:"timestamp,required"`
	JSON      dexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeValueJSON `json:"-"`
}

// dexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeValueJSON
// contains the JSON metadata for the struct
// [DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeValue]
type dexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeValueJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponseResultTestsHTTPResultsByColo struct {
	// Cloudflare colo
	Colo              string                                                           `json:"colo,required"`
	ResourceFetchTime DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTime `json:"resourceFetchTime,required"`
	JSON              dexTestListResponseResultTestsHTTPResultsByColoJSON              `json:"-"`
}

// dexTestListResponseResultTestsHTTPResultsByColoJSON contains the JSON metadata
// for the struct [DEXTestListResponseResultTestsHTTPResultsByColo]
type dexTestListResponseResultTestsHTTPResultsByColoJSON struct {
	Colo              apijson.Field
	ResourceFetchTime apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DEXTestListResponseResultTestsHTTPResultsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultTestsHTTPResultsByColoJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTime struct {
	History  []DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistory `json:"history,required"`
	AvgMs    int64                                                                     `json:"avgMs,nullable"`
	OverTime DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTime  `json:"overTime,nullable"`
	JSON     dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeJSON      `json:"-"`
}

// dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeJSON contains
// the JSON metadata for the struct
// [DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTime]
type dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistory struct {
	TimePeriod DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriod `json:"timePeriod,required"`
	AvgMs      int64                                                                             `json:"avgMs,nullable"`
	DeltaPct   float64                                                                           `json:"deltaPct,nullable"`
	JSON       dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryJSON       `json:"-"`
}

// dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryJSON
// contains the JSON metadata for the struct
// [DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistory]
type dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriod struct {
	Units DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits `json:"units,required"`
	Value int64                                                                                  `json:"value,required"`
	JSON  dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodJSON  `json:"-"`
}

// dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodJSON
// contains the JSON metadata for the struct
// [DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriod]
type dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits string

const (
	DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnitsHours    DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits = "hours"
	DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnitsDays     DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits = "days"
	DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnitsTestRuns DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits = "testRuns"
)

type DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTime struct {
	TimePeriod DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriod `json:"timePeriod,required"`
	Values     []DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeValue    `json:"values,required"`
	JSON       dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeJSON       `json:"-"`
}

// dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeJSON
// contains the JSON metadata for the struct
// [DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTime]
type dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriod struct {
	Units DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits `json:"units,required"`
	Value int64                                                                                   `json:"value,required"`
	JSON  dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodJSON  `json:"-"`
}

// dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodJSON
// contains the JSON metadata for the struct
// [DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriod]
type dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits string

const (
	DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnitsHours    DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits = "hours"
	DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnitsDays     DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits = "days"
	DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnitsTestRuns DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits = "testRuns"
)

type DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeValue struct {
	AvgMs     int64                                                                             `json:"avgMs,required"`
	Timestamp string                                                                            `json:"timestamp,required"`
	JSON      dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeValueJSON `json:"-"`
}

// dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeValueJSON
// contains the JSON metadata for the struct
// [DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeValue]
type dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeValueJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponseResultTestsTracerouteResults struct {
	RoundTripTime DEXTestListResponseResultTestsTracerouteResultsRoundTripTime `json:"roundTripTime,required"`
	JSON          dexTestListResponseResultTestsTracerouteResultsJSON          `json:"-"`
}

// dexTestListResponseResultTestsTracerouteResultsJSON contains the JSON metadata
// for the struct [DEXTestListResponseResultTestsTracerouteResults]
type dexTestListResponseResultTestsTracerouteResultsJSON struct {
	RoundTripTime apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DEXTestListResponseResultTestsTracerouteResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultTestsTracerouteResultsJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponseResultTestsTracerouteResultsRoundTripTime struct {
	History  []DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistory `json:"history,required"`
	AvgMs    int64                                                                 `json:"avgMs,nullable"`
	OverTime DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTime  `json:"overTime,nullable"`
	JSON     dexTestListResponseResultTestsTracerouteResultsRoundTripTimeJSON      `json:"-"`
}

// dexTestListResponseResultTestsTracerouteResultsRoundTripTimeJSON contains the
// JSON metadata for the struct
// [DEXTestListResponseResultTestsTracerouteResultsRoundTripTime]
type dexTestListResponseResultTestsTracerouteResultsRoundTripTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseResultTestsTracerouteResultsRoundTripTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultTestsTracerouteResultsRoundTripTimeJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistory struct {
	TimePeriod DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriod `json:"timePeriod,required"`
	AvgMs      int64                                                                         `json:"avgMs,nullable"`
	DeltaPct   float64                                                                       `json:"deltaPct,nullable"`
	JSON       dexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryJSON       `json:"-"`
}

// dexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryJSON contains
// the JSON metadata for the struct
// [DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistory]
type dexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriod struct {
	Units DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits `json:"units,required"`
	Value int64                                                                              `json:"value,required"`
	JSON  dexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodJSON  `json:"-"`
}

// dexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodJSON
// contains the JSON metadata for the struct
// [DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriod]
type dexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits string

const (
	DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnitsHours    DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits = "hours"
	DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnitsDays     DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits = "days"
	DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnitsTestRuns DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits = "testRuns"
)

type DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTime struct {
	TimePeriod DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriod `json:"timePeriod,required"`
	Values     []DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeValue    `json:"values,required"`
	JSON       dexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeJSON       `json:"-"`
}

// dexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeJSON
// contains the JSON metadata for the struct
// [DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTime]
type dexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriod struct {
	Units DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits `json:"units,required"`
	Value int64                                                                               `json:"value,required"`
	JSON  dexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodJSON  `json:"-"`
}

// dexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodJSON
// contains the JSON metadata for the struct
// [DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriod]
type dexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits string

const (
	DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnitsHours    DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits = "hours"
	DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnitsDays     DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits = "days"
	DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnitsTestRuns DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits = "testRuns"
)

type DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeValue struct {
	AvgMs     int64                                                                         `json:"avgMs,required"`
	Timestamp string                                                                        `json:"timestamp,required"`
	JSON      dexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeValueJSON `json:"-"`
}

// dexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeValueJSON
// contains the JSON metadata for the struct
// [DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeValue]
type dexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeValueJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponseResultTestsTracerouteResultsByColo struct {
	// Cloudflare colo
	Colo          string                                                             `json:"colo,required"`
	RoundTripTime DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTime `json:"roundTripTime,required"`
	JSON          dexTestListResponseResultTestsTracerouteResultsByColoJSON          `json:"-"`
}

// dexTestListResponseResultTestsTracerouteResultsByColoJSON contains the JSON
// metadata for the struct [DEXTestListResponseResultTestsTracerouteResultsByColo]
type dexTestListResponseResultTestsTracerouteResultsByColoJSON struct {
	Colo          apijson.Field
	RoundTripTime apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DEXTestListResponseResultTestsTracerouteResultsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultTestsTracerouteResultsByColoJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTime struct {
	History  []DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistory `json:"history,required"`
	AvgMs    int64                                                                       `json:"avgMs,nullable"`
	OverTime DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTime  `json:"overTime,nullable"`
	JSON     dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeJSON      `json:"-"`
}

// dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeJSON contains
// the JSON metadata for the struct
// [DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTime]
type dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistory struct {
	TimePeriod DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriod `json:"timePeriod,required"`
	AvgMs      int64                                                                               `json:"avgMs,nullable"`
	DeltaPct   float64                                                                             `json:"deltaPct,nullable"`
	JSON       dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryJSON       `json:"-"`
}

// dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryJSON
// contains the JSON metadata for the struct
// [DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistory]
type dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriod struct {
	Units DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits `json:"units,required"`
	Value int64                                                                                    `json:"value,required"`
	JSON  dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodJSON  `json:"-"`
}

// dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodJSON
// contains the JSON metadata for the struct
// [DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriod]
type dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits string

const (
	DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnitsHours    DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits = "hours"
	DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnitsDays     DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits = "days"
	DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnitsTestRuns DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits = "testRuns"
)

type DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTime struct {
	TimePeriod DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriod `json:"timePeriod,required"`
	Values     []DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeValue    `json:"values,required"`
	JSON       dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeJSON       `json:"-"`
}

// dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeJSON
// contains the JSON metadata for the struct
// [DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTime]
type dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriod struct {
	Units DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits `json:"units,required"`
	Value int64                                                                                     `json:"value,required"`
	JSON  dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodJSON  `json:"-"`
}

// dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodJSON
// contains the JSON metadata for the struct
// [DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriod]
type dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodJSON) RawJSON() string {
	return r.raw
}

type DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits string

const (
	DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnitsHours    DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits = "hours"
	DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnitsDays     DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits = "days"
	DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnitsTestRuns DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits = "testRuns"
)

type DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeValue struct {
	AvgMs     int64                                                                               `json:"avgMs,required"`
	Timestamp string                                                                              `json:"timestamp,required"`
	JSON      dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeValueJSON `json:"-"`
}

// dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeValueJSON
// contains the JSON metadata for the struct
// [DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeValue]
type dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeValueJSON) RawJSON() string {
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
