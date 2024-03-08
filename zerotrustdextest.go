// File generated from our OpenAPI spec by Stainless.

package cloudflare

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

// ZeroTrustDEXTestService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustDEXTestService] method
// instead.
type ZeroTrustDEXTestService struct {
	Options       []option.RequestOption
	UniqueDevices *ZeroTrustDEXTestUniqueDeviceService
}

// NewZeroTrustDEXTestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustDEXTestService(opts ...option.RequestOption) (r *ZeroTrustDEXTestService) {
	r = &ZeroTrustDEXTestService{}
	r.Options = opts
	r.UniqueDevices = NewZeroTrustDEXTestUniqueDeviceService(opts...)
	return
}

// List DEX tests
func (r *ZeroTrustDEXTestService) List(ctx context.Context, params ZeroTrustDEXTestListParams, opts ...option.RequestOption) (res *shared.V4PagePagination[ZeroTrustDEXTestListResponse], err error) {
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
func (r *ZeroTrustDEXTestService) ListAutoPaging(ctx context.Context, params ZeroTrustDEXTestListParams, opts ...option.RequestOption) *shared.V4PagePaginationAutoPager[ZeroTrustDEXTestListResponse] {
	return shared.NewV4PagePaginationAutoPager(r.List(ctx, params, opts...))
}

type ZeroTrustDEXTestListResponse struct {
	Errors   []ZeroTrustDEXTestListResponseError   `json:"errors,required"`
	Messages []ZeroTrustDEXTestListResponseMessage `json:"messages,required"`
	Result   ZeroTrustDEXTestListResponseResult    `json:"result,required"`
	// Whether the API call was successful
	Success    ZeroTrustDEXTestListResponseSuccess    `json:"success,required"`
	ResultInfo ZeroTrustDEXTestListResponseResultInfo `json:"result_info"`
	JSON       zeroTrustDEXTestListResponseJSON       `json:"-"`
}

// zeroTrustDEXTestListResponseJSON contains the JSON metadata for the struct
// [ZeroTrustDEXTestListResponse]
type zeroTrustDEXTestListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zeroTrustDEXTestListResponseErrorJSON `json:"-"`
}

// zeroTrustDEXTestListResponseErrorJSON contains the JSON metadata for the struct
// [ZeroTrustDEXTestListResponseError]
type zeroTrustDEXTestListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zeroTrustDEXTestListResponseMessageJSON `json:"-"`
}

// zeroTrustDEXTestListResponseMessageJSON contains the JSON metadata for the
// struct [ZeroTrustDEXTestListResponseMessage]
type zeroTrustDEXTestListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseResult struct {
	OverviewMetrics ZeroTrustDEXTestListResponseResultOverviewMetrics `json:"overviewMetrics,required"`
	// array of test results objects.
	Tests []ZeroTrustDEXTestListResponseResultTest `json:"tests,required"`
	JSON  zeroTrustDEXTestListResponseResultJSON   `json:"-"`
}

// zeroTrustDEXTestListResponseResultJSON contains the JSON metadata for the struct
// [ZeroTrustDEXTestListResponseResult]
type zeroTrustDEXTestListResponseResultJSON struct {
	OverviewMetrics apijson.Field
	Tests           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseResultOverviewMetrics struct {
	// number of tests.
	TestsTotal int64 `json:"testsTotal,required"`
	// percentage availability for all traceroutes results in response
	AvgTracerouteAvailabilityPct float64                                               `json:"avgTracerouteAvailabilityPct,nullable"`
	JSON                         zeroTrustDEXTestListResponseResultOverviewMetricsJSON `json:"-"`
}

// zeroTrustDEXTestListResponseResultOverviewMetricsJSON contains the JSON metadata
// for the struct [ZeroTrustDEXTestListResponseResultOverviewMetrics]
type zeroTrustDEXTestListResponseResultOverviewMetricsJSON struct {
	TestsTotal                   apijson.Field
	AvgTracerouteAvailabilityPct apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseResultOverviewMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultOverviewMetricsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseResultTest struct {
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
	Kind ZeroTrustDEXTestListResponseResultTestsKind `json:"kind,required"`
	// name given to this test
	Name              string                                                     `json:"name,required"`
	Updated           string                                                     `json:"updated,required"`
	HTTPResults       ZeroTrustDEXTestListResponseResultTestsHTTPResults         `json:"httpResults,nullable"`
	HTTPResultsByColo []ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColo `json:"httpResultsByColo"`
	// for HTTP, the method to use when running the test
	Method                  string                                                           `json:"method"`
	TracerouteResults       ZeroTrustDEXTestListResponseResultTestsTracerouteResults         `json:"tracerouteResults,nullable"`
	TracerouteResultsByColo []ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColo `json:"tracerouteResultsByColo"`
	JSON                    zeroTrustDEXTestListResponseResultTestJSON                       `json:"-"`
}

// zeroTrustDEXTestListResponseResultTestJSON contains the JSON metadata for the
// struct [ZeroTrustDEXTestListResponseResultTest]
type zeroTrustDEXTestListResponseResultTestJSON struct {
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

func (r *ZeroTrustDEXTestListResponseResultTest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultTestJSON) RawJSON() string {
	return r.raw
}

// test type, http or traceroute
type ZeroTrustDEXTestListResponseResultTestsKind string

const (
	ZeroTrustDEXTestListResponseResultTestsKindHTTP       ZeroTrustDEXTestListResponseResultTestsKind = "http"
	ZeroTrustDEXTestListResponseResultTestsKindTraceroute ZeroTrustDEXTestListResponseResultTestsKind = "traceroute"
)

type ZeroTrustDEXTestListResponseResultTestsHTTPResults struct {
	ResourceFetchTime ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTime `json:"resourceFetchTime,required"`
	JSON              zeroTrustDEXTestListResponseResultTestsHTTPResultsJSON              `json:"-"`
}

// zeroTrustDEXTestListResponseResultTestsHTTPResultsJSON contains the JSON
// metadata for the struct [ZeroTrustDEXTestListResponseResultTestsHTTPResults]
type zeroTrustDEXTestListResponseResultTestsHTTPResultsJSON struct {
	ResourceFetchTime apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseResultTestsHTTPResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultTestsHTTPResultsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTime struct {
	History  []ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistory `json:"history,required"`
	AvgMs    int64                                                                        `json:"avgMs,nullable"`
	OverTime ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTime  `json:"overTime,nullable"`
	JSON     zeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeJSON      `json:"-"`
}

// zeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeJSON contains
// the JSON metadata for the struct
// [ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTime]
type zeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistory struct {
	TimePeriod ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriod `json:"timePeriod,required"`
	AvgMs      int64                                                                                `json:"avgMs,nullable"`
	DeltaPct   float64                                                                              `json:"deltaPct,nullable"`
	JSON       zeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryJSON       `json:"-"`
}

// zeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistory]
type zeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriod struct {
	Units ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits `json:"units,required"`
	Value int64                                                                                     `json:"value,required"`
	JSON  zeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodJSON  `json:"-"`
}

// zeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriod]
type zeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits string

const (
	ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnitsHours    ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits = "hours"
	ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnitsDays     ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits = "days"
	ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnitsTestRuns ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeHistoryTimePeriodUnits = "testRuns"
)

type ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTime struct {
	TimePeriod ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriod `json:"timePeriod,required"`
	Values     []ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeValue    `json:"values,required"`
	JSON       zeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeJSON       `json:"-"`
}

// zeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTime]
type zeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriod struct {
	Units ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits `json:"units,required"`
	Value int64                                                                                      `json:"value,required"`
	JSON  zeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodJSON  `json:"-"`
}

// zeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriod]
type zeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits string

const (
	ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnitsHours    ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits = "hours"
	ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnitsDays     ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits = "days"
	ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnitsTestRuns ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeTimePeriodUnits = "testRuns"
)

type ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeValue struct {
	AvgMs     int64                                                                                `json:"avgMs,required"`
	Timestamp string                                                                               `json:"timestamp,required"`
	JSON      zeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeValueJSON `json:"-"`
}

// zeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeValueJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeValue]
type zeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultTestsHTTPResultsResourceFetchTimeOverTimeValueJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColo struct {
	// Cloudflare colo
	Colo              string                                                                    `json:"colo,required"`
	ResourceFetchTime ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTime `json:"resourceFetchTime,required"`
	JSON              zeroTrustDEXTestListResponseResultTestsHTTPResultsByColoJSON              `json:"-"`
}

// zeroTrustDEXTestListResponseResultTestsHTTPResultsByColoJSON contains the JSON
// metadata for the struct
// [ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColo]
type zeroTrustDEXTestListResponseResultTestsHTTPResultsByColoJSON struct {
	Colo              apijson.Field
	ResourceFetchTime apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultTestsHTTPResultsByColoJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTime struct {
	History  []ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistory `json:"history,required"`
	AvgMs    int64                                                                              `json:"avgMs,nullable"`
	OverTime ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTime  `json:"overTime,nullable"`
	JSON     zeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeJSON      `json:"-"`
}

// zeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTime]
type zeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistory struct {
	TimePeriod ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriod `json:"timePeriod,required"`
	AvgMs      int64                                                                                      `json:"avgMs,nullable"`
	DeltaPct   float64                                                                                    `json:"deltaPct,nullable"`
	JSON       zeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryJSON       `json:"-"`
}

// zeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistory]
type zeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriod struct {
	Units ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits `json:"units,required"`
	Value int64                                                                                           `json:"value,required"`
	JSON  zeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodJSON  `json:"-"`
}

// zeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriod]
type zeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits string

const (
	ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnitsHours    ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits = "hours"
	ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnitsDays     ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits = "days"
	ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnitsTestRuns ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeHistoryTimePeriodUnits = "testRuns"
)

type ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTime struct {
	TimePeriod ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriod `json:"timePeriod,required"`
	Values     []ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeValue    `json:"values,required"`
	JSON       zeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeJSON       `json:"-"`
}

// zeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTime]
type zeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriod struct {
	Units ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits `json:"units,required"`
	Value int64                                                                                            `json:"value,required"`
	JSON  zeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodJSON  `json:"-"`
}

// zeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriod]
type zeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits string

const (
	ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnitsHours    ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits = "hours"
	ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnitsDays     ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits = "days"
	ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnitsTestRuns ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeTimePeriodUnits = "testRuns"
)

type ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeValue struct {
	AvgMs     int64                                                                                      `json:"avgMs,required"`
	Timestamp string                                                                                     `json:"timestamp,required"`
	JSON      zeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeValueJSON `json:"-"`
}

// zeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeValueJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeValue]
type zeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultTestsHTTPResultsByColoResourceFetchTimeOverTimeValueJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseResultTestsTracerouteResults struct {
	RoundTripTime ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTime `json:"roundTripTime,required"`
	JSON          zeroTrustDEXTestListResponseResultTestsTracerouteResultsJSON          `json:"-"`
}

// zeroTrustDEXTestListResponseResultTestsTracerouteResultsJSON contains the JSON
// metadata for the struct
// [ZeroTrustDEXTestListResponseResultTestsTracerouteResults]
type zeroTrustDEXTestListResponseResultTestsTracerouteResultsJSON struct {
	RoundTripTime apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseResultTestsTracerouteResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultTestsTracerouteResultsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTime struct {
	History  []ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistory `json:"history,required"`
	AvgMs    int64                                                                          `json:"avgMs,nullable"`
	OverTime ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTime  `json:"overTime,nullable"`
	JSON     zeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeJSON      `json:"-"`
}

// zeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTime]
type zeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistory struct {
	TimePeriod ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriod `json:"timePeriod,required"`
	AvgMs      int64                                                                                  `json:"avgMs,nullable"`
	DeltaPct   float64                                                                                `json:"deltaPct,nullable"`
	JSON       zeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryJSON       `json:"-"`
}

// zeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistory]
type zeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriod struct {
	Units ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits `json:"units,required"`
	Value int64                                                                                       `json:"value,required"`
	JSON  zeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodJSON  `json:"-"`
}

// zeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriod]
type zeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits string

const (
	ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnitsHours    ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits = "hours"
	ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnitsDays     ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits = "days"
	ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnitsTestRuns ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeHistoryTimePeriodUnits = "testRuns"
)

type ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTime struct {
	TimePeriod ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriod `json:"timePeriod,required"`
	Values     []ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeValue    `json:"values,required"`
	JSON       zeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeJSON       `json:"-"`
}

// zeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTime]
type zeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriod struct {
	Units ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits `json:"units,required"`
	Value int64                                                                                        `json:"value,required"`
	JSON  zeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodJSON  `json:"-"`
}

// zeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriod]
type zeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits string

const (
	ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnitsHours    ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits = "hours"
	ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnitsDays     ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits = "days"
	ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnitsTestRuns ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeTimePeriodUnits = "testRuns"
)

type ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeValue struct {
	AvgMs     int64                                                                                  `json:"avgMs,required"`
	Timestamp string                                                                                 `json:"timestamp,required"`
	JSON      zeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeValueJSON `json:"-"`
}

// zeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeValueJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeValue]
type zeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultTestsTracerouteResultsRoundTripTimeOverTimeValueJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColo struct {
	// Cloudflare colo
	Colo          string                                                                      `json:"colo,required"`
	RoundTripTime ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTime `json:"roundTripTime,required"`
	JSON          zeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoJSON          `json:"-"`
}

// zeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoJSON contains the
// JSON metadata for the struct
// [ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColo]
type zeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoJSON struct {
	Colo          apijson.Field
	RoundTripTime apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTime struct {
	History  []ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistory `json:"history,required"`
	AvgMs    int64                                                                                `json:"avgMs,nullable"`
	OverTime ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTime  `json:"overTime,nullable"`
	JSON     zeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeJSON      `json:"-"`
}

// zeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTime]
type zeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistory struct {
	TimePeriod ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriod `json:"timePeriod,required"`
	AvgMs      int64                                                                                        `json:"avgMs,nullable"`
	DeltaPct   float64                                                                                      `json:"deltaPct,nullable"`
	JSON       zeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryJSON       `json:"-"`
}

// zeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistory]
type zeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriod struct {
	Units ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits `json:"units,required"`
	Value int64                                                                                             `json:"value,required"`
	JSON  zeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodJSON  `json:"-"`
}

// zeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriod]
type zeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits string

const (
	ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnitsHours    ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits = "hours"
	ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnitsDays     ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits = "days"
	ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnitsTestRuns ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeHistoryTimePeriodUnits = "testRuns"
)

type ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTime struct {
	TimePeriod ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriod `json:"timePeriod,required"`
	Values     []ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeValue    `json:"values,required"`
	JSON       zeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeJSON       `json:"-"`
}

// zeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTime]
type zeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriod struct {
	Units ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits `json:"units,required"`
	Value int64                                                                                              `json:"value,required"`
	JSON  zeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodJSON  `json:"-"`
}

// zeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriod]
type zeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits string

const (
	ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnitsHours    ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits = "hours"
	ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnitsDays     ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits = "days"
	ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnitsTestRuns ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeTimePeriodUnits = "testRuns"
)

type ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeValue struct {
	AvgMs     int64                                                                                        `json:"avgMs,required"`
	Timestamp string                                                                                       `json:"timestamp,required"`
	JSON      zeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeValueJSON `json:"-"`
}

// zeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeValueJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeValue]
type zeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultTestsTracerouteResultsByColoRoundTripTimeOverTimeValueJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustDEXTestListResponseSuccess bool

const (
	ZeroTrustDEXTestListResponseSuccessTrue ZeroTrustDEXTestListResponseSuccess = true
)

type ZeroTrustDEXTestListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                    `json:"total_count"`
	JSON       zeroTrustDEXTestListResponseResultInfoJSON `json:"-"`
}

// zeroTrustDEXTestListResponseResultInfoJSON contains the JSON metadata for the
// struct [ZeroTrustDEXTestListResponseResultInfo]
type zeroTrustDEXTestListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTestListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestListParams struct {
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

// URLQuery serializes [ZeroTrustDEXTestListParams]'s query parameters as
// `url.Values`.
func (r ZeroTrustDEXTestListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
