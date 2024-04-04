// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// DEXHTTPTestService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDEXHTTPTestService] method
// instead.
type DEXHTTPTestService struct {
	Options     []option.RequestOption
	Percentiles *DEXHTTPTestPercentileService
}

// NewDEXHTTPTestService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDEXHTTPTestService(opts ...option.RequestOption) (r *DEXHTTPTestService) {
	r = &DEXHTTPTestService{}
	r.Options = opts
	r.Percentiles = NewDEXHTTPTestPercentileService(opts...)
	return
}

// Get test details and aggregate performance metrics for an http test for a given
// time period between 1 hour and 7 days.
func (r *DEXHTTPTestService) Get(ctx context.Context, testID string, params DEXHTTPTestGetParams, opts ...option.RequestOption) (res *DigitalExperienceMonitoringHTTPDetails, err error) {
	opts = append(r.Options[:], opts...)
	var env DexhttpTestGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/http-tests/%s", params.AccountID, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DigitalExperienceMonitoringHTTPDetails struct {
	// The url of the HTTP synthetic application test
	Host            string                                                  `json:"host"`
	HTTPStats       DigitalExperienceMonitoringHTTPDetailsHTTPStats         `json:"httpStats,nullable"`
	HTTPStatsByColo []DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColo `json:"httpStatsByColo"`
	// The interval at which the HTTP synthetic application test is set to run.
	Interval string                                     `json:"interval"`
	Kind     DigitalExperienceMonitoringHTTPDetailsKind `json:"kind"`
	// The HTTP method to use when running the test
	Method string `json:"method"`
	// The name of the HTTP synthetic application test
	Name           string                                               `json:"name"`
	TargetPolicies []DigitalExperienceMonitoringHTTPDetailsTargetPolicy `json:"target_policies"`
	Targeted       bool                                                 `json:"targeted"`
	JSON           digitalExperienceMonitoringHTTPDetailsJSON           `json:"-"`
}

// digitalExperienceMonitoringHTTPDetailsJSON contains the JSON metadata for the
// struct [DigitalExperienceMonitoringHTTPDetails]
type digitalExperienceMonitoringHTTPDetailsJSON struct {
	Host            apijson.Field
	HTTPStats       apijson.Field
	HTTPStatsByColo apijson.Field
	Interval        apijson.Field
	Kind            apijson.Field
	Method          apijson.Field
	Name            apijson.Field
	TargetPolicies  apijson.Field
	Targeted        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringHTTPDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringHTTPDetailsJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringHTTPDetailsHTTPStats struct {
	AvailabilityPct      DigitalExperienceMonitoringHTTPDetailsHTTPStatsAvailabilityPct      `json:"availabilityPct,required"`
	DNSResponseTimeMs    DigitalExperienceMonitoringHTTPDetailsHTTPStatsDNSResponseTimeMs    `json:"dnsResponseTimeMs,required"`
	HTTPStatusCode       []DigitalExperienceMonitoringHTTPDetailsHTTPStatsHTTPStatusCode     `json:"httpStatusCode,required"`
	ResourceFetchTimeMs  DigitalExperienceMonitoringHTTPDetailsHTTPStatsResourceFetchTimeMs  `json:"resourceFetchTimeMs,required"`
	ServerResponseTimeMs DigitalExperienceMonitoringHTTPDetailsHTTPStatsServerResponseTimeMs `json:"serverResponseTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                                               `json:"uniqueDevicesTotal,required"`
	JSON               digitalExperienceMonitoringHTTPDetailsHTTPStatsJSON `json:"-"`
}

// digitalExperienceMonitoringHTTPDetailsHTTPStatsJSON contains the JSON metadata
// for the struct [DigitalExperienceMonitoringHTTPDetailsHTTPStats]
type digitalExperienceMonitoringHTTPDetailsHTTPStatsJSON struct {
	AvailabilityPct      apijson.Field
	DNSResponseTimeMs    apijson.Field
	HTTPStatusCode       apijson.Field
	ResourceFetchTimeMs  apijson.Field
	ServerResponseTimeMs apijson.Field
	UniqueDevicesTotal   apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringHTTPDetailsHTTPStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringHTTPDetailsHTTPStatsJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringHTTPDetailsHTTPStatsAvailabilityPct struct {
	Slots []DigitalExperienceMonitoringHTTPDetailsHTTPStatsAvailabilityPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                                            `json:"min,nullable"`
	JSON digitalExperienceMonitoringHTTPDetailsHTTPStatsAvailabilityPctJSON `json:"-"`
}

// digitalExperienceMonitoringHTTPDetailsHTTPStatsAvailabilityPctJSON contains the
// JSON metadata for the struct
// [DigitalExperienceMonitoringHTTPDetailsHTTPStatsAvailabilityPct]
type digitalExperienceMonitoringHTTPDetailsHTTPStatsAvailabilityPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringHTTPDetailsHTTPStatsAvailabilityPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringHTTPDetailsHTTPStatsAvailabilityPctJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringHTTPDetailsHTTPStatsAvailabilityPctSlot struct {
	Timestamp string                                                                 `json:"timestamp,required"`
	Value     float64                                                                `json:"value,required"`
	JSON      digitalExperienceMonitoringHTTPDetailsHTTPStatsAvailabilityPctSlotJSON `json:"-"`
}

// digitalExperienceMonitoringHTTPDetailsHTTPStatsAvailabilityPctSlotJSON contains
// the JSON metadata for the struct
// [DigitalExperienceMonitoringHTTPDetailsHTTPStatsAvailabilityPctSlot]
type digitalExperienceMonitoringHTTPDetailsHTTPStatsAvailabilityPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringHTTPDetailsHTTPStatsAvailabilityPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringHTTPDetailsHTTPStatsAvailabilityPctSlotJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringHTTPDetailsHTTPStatsDNSResponseTimeMs struct {
	Slots []DigitalExperienceMonitoringHTTPDetailsHTTPStatsDNSResponseTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                                `json:"min,nullable"`
	JSON digitalExperienceMonitoringHTTPDetailsHTTPStatsDNSResponseTimeMsJSON `json:"-"`
}

// digitalExperienceMonitoringHTTPDetailsHTTPStatsDNSResponseTimeMsJSON contains
// the JSON metadata for the struct
// [DigitalExperienceMonitoringHTTPDetailsHTTPStatsDNSResponseTimeMs]
type digitalExperienceMonitoringHTTPDetailsHTTPStatsDNSResponseTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringHTTPDetailsHTTPStatsDNSResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringHTTPDetailsHTTPStatsDNSResponseTimeMsJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringHTTPDetailsHTTPStatsDNSResponseTimeMsSlot struct {
	Timestamp string                                                                   `json:"timestamp,required"`
	Value     int64                                                                    `json:"value,required"`
	JSON      digitalExperienceMonitoringHTTPDetailsHTTPStatsDNSResponseTimeMsSlotJSON `json:"-"`
}

// digitalExperienceMonitoringHTTPDetailsHTTPStatsDNSResponseTimeMsSlotJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringHTTPDetailsHTTPStatsDNSResponseTimeMsSlot]
type digitalExperienceMonitoringHTTPDetailsHTTPStatsDNSResponseTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringHTTPDetailsHTTPStatsDNSResponseTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringHTTPDetailsHTTPStatsDNSResponseTimeMsSlotJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringHTTPDetailsHTTPStatsHTTPStatusCode struct {
	Status200 int64                                                             `json:"status200,required"`
	Status300 int64                                                             `json:"status300,required"`
	Status400 int64                                                             `json:"status400,required"`
	Status500 int64                                                             `json:"status500,required"`
	Timestamp string                                                            `json:"timestamp,required"`
	JSON      digitalExperienceMonitoringHTTPDetailsHTTPStatsHTTPStatusCodeJSON `json:"-"`
}

// digitalExperienceMonitoringHTTPDetailsHTTPStatsHTTPStatusCodeJSON contains the
// JSON metadata for the struct
// [DigitalExperienceMonitoringHTTPDetailsHTTPStatsHTTPStatusCode]
type digitalExperienceMonitoringHTTPDetailsHTTPStatsHTTPStatusCodeJSON struct {
	Status200   apijson.Field
	Status300   apijson.Field
	Status400   apijson.Field
	Status500   apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringHTTPDetailsHTTPStatsHTTPStatusCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringHTTPDetailsHTTPStatsHTTPStatusCodeJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringHTTPDetailsHTTPStatsResourceFetchTimeMs struct {
	Slots []DigitalExperienceMonitoringHTTPDetailsHTTPStatsResourceFetchTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                                  `json:"min,nullable"`
	JSON digitalExperienceMonitoringHTTPDetailsHTTPStatsResourceFetchTimeMsJSON `json:"-"`
}

// digitalExperienceMonitoringHTTPDetailsHTTPStatsResourceFetchTimeMsJSON contains
// the JSON metadata for the struct
// [DigitalExperienceMonitoringHTTPDetailsHTTPStatsResourceFetchTimeMs]
type digitalExperienceMonitoringHTTPDetailsHTTPStatsResourceFetchTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringHTTPDetailsHTTPStatsResourceFetchTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringHTTPDetailsHTTPStatsResourceFetchTimeMsJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringHTTPDetailsHTTPStatsResourceFetchTimeMsSlot struct {
	Timestamp string                                                                     `json:"timestamp,required"`
	Value     int64                                                                      `json:"value,required"`
	JSON      digitalExperienceMonitoringHTTPDetailsHTTPStatsResourceFetchTimeMsSlotJSON `json:"-"`
}

// digitalExperienceMonitoringHTTPDetailsHTTPStatsResourceFetchTimeMsSlotJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringHTTPDetailsHTTPStatsResourceFetchTimeMsSlot]
type digitalExperienceMonitoringHTTPDetailsHTTPStatsResourceFetchTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringHTTPDetailsHTTPStatsResourceFetchTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringHTTPDetailsHTTPStatsResourceFetchTimeMsSlotJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringHTTPDetailsHTTPStatsServerResponseTimeMs struct {
	Slots []DigitalExperienceMonitoringHTTPDetailsHTTPStatsServerResponseTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                                   `json:"min,nullable"`
	JSON digitalExperienceMonitoringHTTPDetailsHTTPStatsServerResponseTimeMsJSON `json:"-"`
}

// digitalExperienceMonitoringHTTPDetailsHTTPStatsServerResponseTimeMsJSON contains
// the JSON metadata for the struct
// [DigitalExperienceMonitoringHTTPDetailsHTTPStatsServerResponseTimeMs]
type digitalExperienceMonitoringHTTPDetailsHTTPStatsServerResponseTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringHTTPDetailsHTTPStatsServerResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringHTTPDetailsHTTPStatsServerResponseTimeMsJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringHTTPDetailsHTTPStatsServerResponseTimeMsSlot struct {
	Timestamp string                                                                      `json:"timestamp,required"`
	Value     int64                                                                       `json:"value,required"`
	JSON      digitalExperienceMonitoringHTTPDetailsHTTPStatsServerResponseTimeMsSlotJSON `json:"-"`
}

// digitalExperienceMonitoringHTTPDetailsHTTPStatsServerResponseTimeMsSlotJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringHTTPDetailsHTTPStatsServerResponseTimeMsSlot]
type digitalExperienceMonitoringHTTPDetailsHTTPStatsServerResponseTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringHTTPDetailsHTTPStatsServerResponseTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringHTTPDetailsHTTPStatsServerResponseTimeMsSlotJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColo struct {
	AvailabilityPct      DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoAvailabilityPct      `json:"availabilityPct,required"`
	Colo                 string                                                                    `json:"colo,required"`
	DNSResponseTimeMs    DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoDNSResponseTimeMs    `json:"dnsResponseTimeMs,required"`
	HTTPStatusCode       []DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoHTTPStatusCode     `json:"httpStatusCode,required"`
	ResourceFetchTimeMs  DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoResourceFetchTimeMs  `json:"resourceFetchTimeMs,required"`
	ServerResponseTimeMs DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoServerResponseTimeMs `json:"serverResponseTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                                                     `json:"uniqueDevicesTotal,required"`
	JSON               digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoJSON `json:"-"`
}

// digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoJSON contains the JSON
// metadata for the struct [DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColo]
type digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoJSON struct {
	AvailabilityPct      apijson.Field
	Colo                 apijson.Field
	DNSResponseTimeMs    apijson.Field
	HTTPStatusCode       apijson.Field
	ResourceFetchTimeMs  apijson.Field
	ServerResponseTimeMs apijson.Field
	UniqueDevicesTotal   apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoAvailabilityPct struct {
	Slots []DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoAvailabilityPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                                                  `json:"min,nullable"`
	JSON digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoAvailabilityPctJSON `json:"-"`
}

// digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoAvailabilityPctJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoAvailabilityPct]
type digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoAvailabilityPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoAvailabilityPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoAvailabilityPctJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoAvailabilityPctSlot struct {
	Timestamp string                                                                       `json:"timestamp,required"`
	Value     float64                                                                      `json:"value,required"`
	JSON      digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoAvailabilityPctSlotJSON `json:"-"`
}

// digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoAvailabilityPctSlotJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoAvailabilityPctSlot]
type digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoAvailabilityPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoAvailabilityPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoAvailabilityPctSlotJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoDNSResponseTimeMs struct {
	Slots []DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoDNSResponseTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                                      `json:"min,nullable"`
	JSON digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoDNSResponseTimeMsJSON `json:"-"`
}

// digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoDNSResponseTimeMsJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoDNSResponseTimeMs]
type digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoDNSResponseTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoDNSResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoDNSResponseTimeMsJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoDNSResponseTimeMsSlot struct {
	Timestamp string                                                                         `json:"timestamp,required"`
	Value     int64                                                                          `json:"value,required"`
	JSON      digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoDNSResponseTimeMsSlotJSON `json:"-"`
}

// digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoDNSResponseTimeMsSlotJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoDNSResponseTimeMsSlot]
type digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoDNSResponseTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoDNSResponseTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoDNSResponseTimeMsSlotJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoHTTPStatusCode struct {
	Status200 int64                                                                   `json:"status200,required"`
	Status300 int64                                                                   `json:"status300,required"`
	Status400 int64                                                                   `json:"status400,required"`
	Status500 int64                                                                   `json:"status500,required"`
	Timestamp string                                                                  `json:"timestamp,required"`
	JSON      digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoHTTPStatusCodeJSON `json:"-"`
}

// digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoHTTPStatusCodeJSON contains
// the JSON metadata for the struct
// [DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoHTTPStatusCode]
type digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoHTTPStatusCodeJSON struct {
	Status200   apijson.Field
	Status300   apijson.Field
	Status400   apijson.Field
	Status500   apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoHTTPStatusCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoHTTPStatusCodeJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoResourceFetchTimeMs struct {
	Slots []DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoResourceFetchTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                                        `json:"min,nullable"`
	JSON digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoResourceFetchTimeMsJSON `json:"-"`
}

// digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoResourceFetchTimeMsJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoResourceFetchTimeMs]
type digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoResourceFetchTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoResourceFetchTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoResourceFetchTimeMsJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoResourceFetchTimeMsSlot struct {
	Timestamp string                                                                           `json:"timestamp,required"`
	Value     int64                                                                            `json:"value,required"`
	JSON      digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoResourceFetchTimeMsSlotJSON `json:"-"`
}

// digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoResourceFetchTimeMsSlotJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoResourceFetchTimeMsSlot]
type digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoResourceFetchTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoResourceFetchTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoResourceFetchTimeMsSlotJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoServerResponseTimeMs struct {
	Slots []DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoServerResponseTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                                         `json:"min,nullable"`
	JSON digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoServerResponseTimeMsJSON `json:"-"`
}

// digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoServerResponseTimeMsJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoServerResponseTimeMs]
type digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoServerResponseTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoServerResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoServerResponseTimeMsJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoServerResponseTimeMsSlot struct {
	Timestamp string                                                                            `json:"timestamp,required"`
	Value     int64                                                                             `json:"value,required"`
	JSON      digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoServerResponseTimeMsSlotJSON `json:"-"`
}

// digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoServerResponseTimeMsSlotJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoServerResponseTimeMsSlot]
type digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoServerResponseTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColoServerResponseTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringHTTPDetailsHTTPStatsByColoServerResponseTimeMsSlotJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringHTTPDetailsKind string

const (
	DigitalExperienceMonitoringHTTPDetailsKindHTTP DigitalExperienceMonitoringHTTPDetailsKind = "http"
)

func (r DigitalExperienceMonitoringHTTPDetailsKind) IsKnown() bool {
	switch r {
	case DigitalExperienceMonitoringHTTPDetailsKindHTTP:
		return true
	}
	return false
}

type DigitalExperienceMonitoringHTTPDetailsTargetPolicy struct {
	ID string `json:"id,required"`
	// Whether the policy is the default for the account
	Default bool                                                   `json:"default,required"`
	Name    string                                                 `json:"name,required"`
	JSON    digitalExperienceMonitoringHTTPDetailsTargetPolicyJSON `json:"-"`
}

// digitalExperienceMonitoringHTTPDetailsTargetPolicyJSON contains the JSON
// metadata for the struct [DigitalExperienceMonitoringHTTPDetailsTargetPolicy]
type digitalExperienceMonitoringHTTPDetailsTargetPolicyJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringHTTPDetailsTargetPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringHTTPDetailsTargetPolicyJSON) RawJSON() string {
	return r.raw
}

type DEXHTTPTestGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Time interval for aggregate time slots.
	Interval param.Field[DexhttpTestGetParamsInterval] `query:"interval,required"`
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

// URLQuery serializes [DEXHTTPTestGetParams]'s query parameters as `url.Values`.
func (r DEXHTTPTestGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Time interval for aggregate time slots.
type DexhttpTestGetParamsInterval string

const (
	DexhttpTestGetParamsIntervalMinute DexhttpTestGetParamsInterval = "minute"
	DexhttpTestGetParamsIntervalHour   DexhttpTestGetParamsInterval = "hour"
)

func (r DexhttpTestGetParamsInterval) IsKnown() bool {
	switch r {
	case DexhttpTestGetParamsIntervalMinute, DexhttpTestGetParamsIntervalHour:
		return true
	}
	return false
}

type DexhttpTestGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo                  `json:"errors,required"`
	Messages []shared.ResponseInfo                  `json:"messages,required"`
	Result   DigitalExperienceMonitoringHTTPDetails `json:"result,required"`
	// Whether the API call was successful
	Success DexhttpTestGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dexhttpTestGetResponseEnvelopeJSON    `json:"-"`
}

// dexhttpTestGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DexhttpTestGetResponseEnvelope]
type dexhttpTestGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexhttpTestGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexhttpTestGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DexhttpTestGetResponseEnvelopeSuccess bool

const (
	DexhttpTestGetResponseEnvelopeSuccessTrue DexhttpTestGetResponseEnvelopeSuccess = true
)

func (r DexhttpTestGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DexhttpTestGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
