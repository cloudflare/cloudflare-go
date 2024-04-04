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

// DEXTracerouteTestService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDEXTracerouteTestService] method
// instead.
type DEXTracerouteTestService struct {
	Options []option.RequestOption
}

// NewDEXTracerouteTestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDEXTracerouteTestService(opts ...option.RequestOption) (r *DEXTracerouteTestService) {
	r = &DEXTracerouteTestService{}
	r.Options = opts
	return
}

// Get test details and aggregate performance metrics for an traceroute test for a
// given time period between 1 hour and 7 days.
func (r *DEXTracerouteTestService) Get(ctx context.Context, testID string, params DEXTracerouteTestGetParams, opts ...option.RequestOption) (res *DigitalExperienceMonitoringTracerouteDetails, err error) {
	opts = append(r.Options[:], opts...)
	var env DEXTracerouteTestGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/traceroute-tests/%s", params.AccountID, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a breakdown of metrics by hop for individual traceroute test runs
func (r *DEXTracerouteTestService) NetworkPath(ctx context.Context, testID string, params DEXTracerouteTestNetworkPathParams, opts ...option.RequestOption) (res *DigitalExperienceMonitoringTracerouteTestNetworkPath, err error) {
	opts = append(r.Options[:], opts...)
	var env DEXTracerouteTestNetworkPathResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/traceroute-tests/%s/network-path", params.AccountID, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get percentiles for a traceroute test for a given time period between 1 hour and
// 7 days.
func (r *DEXTracerouteTestService) Percentiles(ctx context.Context, testID string, params DEXTracerouteTestPercentilesParams, opts ...option.RequestOption) (res *DigitalExperienceMonitoringTracerouteDetailsPercentiles, err error) {
	opts = append(r.Options[:], opts...)
	var env DEXTracerouteTestPercentilesResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/traceroute-tests/%s/percentiles", params.AccountID, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DigitalExperienceMonitoringTracerouteDetails struct {
	// The host of the Traceroute synthetic application test
	Host string `json:"host,required"`
	// The interval at which the Traceroute synthetic application test is set to run.
	Interval string                                           `json:"interval,required"`
	Kind     DigitalExperienceMonitoringTracerouteDetailsKind `json:"kind,required"`
	// The name of the Traceroute synthetic application test
	Name                  string                                                              `json:"name,required"`
	TargetPolicies        []UnnamedSchemaRefBf9e2abcf1b78a6cab8e6e29e2228a11                  `json:"target_policies"`
	Targeted              bool                                                                `json:"targeted"`
	TracerouteStats       DigitalExperienceMonitoringTracerouteDetailsTracerouteStats         `json:"tracerouteStats,nullable"`
	TracerouteStatsByColo []DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColo `json:"tracerouteStatsByColo"`
	JSON                  digitalExperienceMonitoringTracerouteDetailsJSON                    `json:"-"`
}

// digitalExperienceMonitoringTracerouteDetailsJSON contains the JSON metadata for
// the struct [DigitalExperienceMonitoringTracerouteDetails]
type digitalExperienceMonitoringTracerouteDetailsJSON struct {
	Host                  apijson.Field
	Interval              apijson.Field
	Kind                  apijson.Field
	Name                  apijson.Field
	TargetPolicies        apijson.Field
	Targeted              apijson.Field
	TracerouteStats       apijson.Field
	TracerouteStatsByColo apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTracerouteDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTracerouteDetailsJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTracerouteDetailsKind string

const (
	DigitalExperienceMonitoringTracerouteDetailsKindTraceroute DigitalExperienceMonitoringTracerouteDetailsKind = "traceroute"
)

func (r DigitalExperienceMonitoringTracerouteDetailsKind) IsKnown() bool {
	switch r {
	case DigitalExperienceMonitoringTracerouteDetailsKindTraceroute:
		return true
	}
	return false
}

type DigitalExperienceMonitoringTracerouteDetailsTracerouteStats struct {
	AvailabilityPct DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsAvailabilityPct `json:"availabilityPct,required"`
	HopsCount       DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsHopsCount       `json:"hopsCount,required"`
	PacketLossPct   DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsPacketLossPct   `json:"packetLossPct,required"`
	RoundTripTimeMs DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsRoundTripTimeMs `json:"roundTripTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                                                           `json:"uniqueDevicesTotal,required"`
	JSON               digitalExperienceMonitoringTracerouteDetailsTracerouteStatsJSON `json:"-"`
}

// digitalExperienceMonitoringTracerouteDetailsTracerouteStatsJSON contains the
// JSON metadata for the struct
// [DigitalExperienceMonitoringTracerouteDetailsTracerouteStats]
type digitalExperienceMonitoringTracerouteDetailsTracerouteStatsJSON struct {
	AvailabilityPct    apijson.Field
	HopsCount          apijson.Field
	PacketLossPct      apijson.Field
	RoundTripTimeMs    apijson.Field
	UniqueDevicesTotal apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTracerouteDetailsTracerouteStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTracerouteDetailsTracerouteStatsJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsAvailabilityPct struct {
	Slots []DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsAvailabilityPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                                                        `json:"min,nullable"`
	JSON digitalExperienceMonitoringTracerouteDetailsTracerouteStatsAvailabilityPctJSON `json:"-"`
}

// digitalExperienceMonitoringTracerouteDetailsTracerouteStatsAvailabilityPctJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsAvailabilityPct]
type digitalExperienceMonitoringTracerouteDetailsTracerouteStatsAvailabilityPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsAvailabilityPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTracerouteDetailsTracerouteStatsAvailabilityPctJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsAvailabilityPctSlot struct {
	Timestamp string                                                                             `json:"timestamp,required"`
	Value     float64                                                                            `json:"value,required"`
	JSON      digitalExperienceMonitoringTracerouteDetailsTracerouteStatsAvailabilityPctSlotJSON `json:"-"`
}

// digitalExperienceMonitoringTracerouteDetailsTracerouteStatsAvailabilityPctSlotJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsAvailabilityPctSlot]
type digitalExperienceMonitoringTracerouteDetailsTracerouteStatsAvailabilityPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsAvailabilityPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTracerouteDetailsTracerouteStatsAvailabilityPctSlotJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsHopsCount struct {
	Slots []DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsHopsCountSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                                    `json:"min,nullable"`
	JSON digitalExperienceMonitoringTracerouteDetailsTracerouteStatsHopsCountJSON `json:"-"`
}

// digitalExperienceMonitoringTracerouteDetailsTracerouteStatsHopsCountJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsHopsCount]
type digitalExperienceMonitoringTracerouteDetailsTracerouteStatsHopsCountJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsHopsCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTracerouteDetailsTracerouteStatsHopsCountJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsHopsCountSlot struct {
	Timestamp string                                                                       `json:"timestamp,required"`
	Value     int64                                                                        `json:"value,required"`
	JSON      digitalExperienceMonitoringTracerouteDetailsTracerouteStatsHopsCountSlotJSON `json:"-"`
}

// digitalExperienceMonitoringTracerouteDetailsTracerouteStatsHopsCountSlotJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsHopsCountSlot]
type digitalExperienceMonitoringTracerouteDetailsTracerouteStatsHopsCountSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsHopsCountSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTracerouteDetailsTracerouteStatsHopsCountSlotJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsPacketLossPct struct {
	Slots []DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsPacketLossPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                                                      `json:"min,nullable"`
	JSON digitalExperienceMonitoringTracerouteDetailsTracerouteStatsPacketLossPctJSON `json:"-"`
}

// digitalExperienceMonitoringTracerouteDetailsTracerouteStatsPacketLossPctJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsPacketLossPct]
type digitalExperienceMonitoringTracerouteDetailsTracerouteStatsPacketLossPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsPacketLossPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTracerouteDetailsTracerouteStatsPacketLossPctJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsPacketLossPctSlot struct {
	Timestamp string                                                                           `json:"timestamp,required"`
	Value     float64                                                                          `json:"value,required"`
	JSON      digitalExperienceMonitoringTracerouteDetailsTracerouteStatsPacketLossPctSlotJSON `json:"-"`
}

// digitalExperienceMonitoringTracerouteDetailsTracerouteStatsPacketLossPctSlotJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsPacketLossPctSlot]
type digitalExperienceMonitoringTracerouteDetailsTracerouteStatsPacketLossPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsPacketLossPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTracerouteDetailsTracerouteStatsPacketLossPctSlotJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsRoundTripTimeMs struct {
	Slots []DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsRoundTripTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                                          `json:"min,nullable"`
	JSON digitalExperienceMonitoringTracerouteDetailsTracerouteStatsRoundTripTimeMsJSON `json:"-"`
}

// digitalExperienceMonitoringTracerouteDetailsTracerouteStatsRoundTripTimeMsJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsRoundTripTimeMs]
type digitalExperienceMonitoringTracerouteDetailsTracerouteStatsRoundTripTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsRoundTripTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTracerouteDetailsTracerouteStatsRoundTripTimeMsJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsRoundTripTimeMsSlot struct {
	Timestamp string                                                                             `json:"timestamp,required"`
	Value     int64                                                                              `json:"value,required"`
	JSON      digitalExperienceMonitoringTracerouteDetailsTracerouteStatsRoundTripTimeMsSlotJSON `json:"-"`
}

// digitalExperienceMonitoringTracerouteDetailsTracerouteStatsRoundTripTimeMsSlotJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsRoundTripTimeMsSlot]
type digitalExperienceMonitoringTracerouteDetailsTracerouteStatsRoundTripTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsRoundTripTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTracerouteDetailsTracerouteStatsRoundTripTimeMsSlotJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColo struct {
	AvailabilityPct DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoAvailabilityPct `json:"availabilityPct,required"`
	Colo            string                                                                           `json:"colo,required"`
	HopsCount       DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoHopsCount       `json:"hopsCount,required"`
	PacketLossPct   DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoPacketLossPct   `json:"packetLossPct,required"`
	RoundTripTimeMs DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoRoundTripTimeMs `json:"roundTripTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                                                                 `json:"uniqueDevicesTotal,required"`
	JSON               digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoJSON `json:"-"`
}

// digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoJSON contains
// the JSON metadata for the struct
// [DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColo]
type digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoJSON struct {
	AvailabilityPct    apijson.Field
	Colo               apijson.Field
	HopsCount          apijson.Field
	PacketLossPct      apijson.Field
	RoundTripTimeMs    apijson.Field
	UniqueDevicesTotal apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoAvailabilityPct struct {
	Slots []DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoAvailabilityPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                                                              `json:"min,nullable"`
	JSON digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoAvailabilityPctJSON `json:"-"`
}

// digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoAvailabilityPctJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoAvailabilityPct]
type digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoAvailabilityPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoAvailabilityPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoAvailabilityPctJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoAvailabilityPctSlot struct {
	Timestamp string                                                                                   `json:"timestamp,required"`
	Value     float64                                                                                  `json:"value,required"`
	JSON      digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoAvailabilityPctSlotJSON `json:"-"`
}

// digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoAvailabilityPctSlotJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoAvailabilityPctSlot]
type digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoAvailabilityPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoAvailabilityPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoAvailabilityPctSlotJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoHopsCount struct {
	Slots []DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoHopsCountSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                                          `json:"min,nullable"`
	JSON digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoHopsCountJSON `json:"-"`
}

// digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoHopsCountJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoHopsCount]
type digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoHopsCountJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoHopsCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoHopsCountJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoHopsCountSlot struct {
	Timestamp string                                                                             `json:"timestamp,required"`
	Value     int64                                                                              `json:"value,required"`
	JSON      digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoHopsCountSlotJSON `json:"-"`
}

// digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoHopsCountSlotJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoHopsCountSlot]
type digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoHopsCountSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoHopsCountSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoHopsCountSlotJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoPacketLossPct struct {
	Slots []DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoPacketLossPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                                                            `json:"min,nullable"`
	JSON digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoPacketLossPctJSON `json:"-"`
}

// digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoPacketLossPctJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoPacketLossPct]
type digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoPacketLossPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoPacketLossPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoPacketLossPctJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoPacketLossPctSlot struct {
	Timestamp string                                                                                 `json:"timestamp,required"`
	Value     float64                                                                                `json:"value,required"`
	JSON      digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoPacketLossPctSlotJSON `json:"-"`
}

// digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoPacketLossPctSlotJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoPacketLossPctSlot]
type digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoPacketLossPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoPacketLossPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoPacketLossPctSlotJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoRoundTripTimeMs struct {
	Slots []DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoRoundTripTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                                                `json:"min,nullable"`
	JSON digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoRoundTripTimeMsJSON `json:"-"`
}

// digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoRoundTripTimeMsJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoRoundTripTimeMs]
type digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoRoundTripTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoRoundTripTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoRoundTripTimeMsJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoRoundTripTimeMsSlot struct {
	Timestamp string                                                                                   `json:"timestamp,required"`
	Value     int64                                                                                    `json:"value,required"`
	JSON      digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoRoundTripTimeMsSlotJSON `json:"-"`
}

// digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoRoundTripTimeMsSlotJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoRoundTripTimeMsSlot]
type digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoRoundTripTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoRoundTripTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTracerouteDetailsTracerouteStatsByColoRoundTripTimeMsSlotJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTracerouteDetailsPercentiles struct {
	HopsCount       DigitalExperienceMonitoringTracerouteDetailsPercentilesHopsCount       `json:"hopsCount"`
	PacketLossPct   DigitalExperienceMonitoringTracerouteDetailsPercentilesPacketLossPct   `json:"packetLossPct"`
	RoundTripTimeMs DigitalExperienceMonitoringTracerouteDetailsPercentilesRoundTripTimeMs `json:"roundTripTimeMs"`
	JSON            digitalExperienceMonitoringTracerouteDetailsPercentilesJSON            `json:"-"`
}

// digitalExperienceMonitoringTracerouteDetailsPercentilesJSON contains the JSON
// metadata for the struct
// [DigitalExperienceMonitoringTracerouteDetailsPercentiles]
type digitalExperienceMonitoringTracerouteDetailsPercentilesJSON struct {
	HopsCount       apijson.Field
	PacketLossPct   apijson.Field
	RoundTripTimeMs apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTracerouteDetailsPercentiles) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTracerouteDetailsPercentilesJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTracerouteDetailsPercentilesHopsCount struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                                              `json:"p99,nullable"`
	JSON digitalExperienceMonitoringTracerouteDetailsPercentilesHopsCountJSON `json:"-"`
}

// digitalExperienceMonitoringTracerouteDetailsPercentilesHopsCountJSON contains
// the JSON metadata for the struct
// [DigitalExperienceMonitoringTracerouteDetailsPercentilesHopsCount]
type digitalExperienceMonitoringTracerouteDetailsPercentilesHopsCountJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTracerouteDetailsPercentilesHopsCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTracerouteDetailsPercentilesHopsCountJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTracerouteDetailsPercentilesPacketLossPct struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                                                  `json:"p99,nullable"`
	JSON digitalExperienceMonitoringTracerouteDetailsPercentilesPacketLossPctJSON `json:"-"`
}

// digitalExperienceMonitoringTracerouteDetailsPercentilesPacketLossPctJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTracerouteDetailsPercentilesPacketLossPct]
type digitalExperienceMonitoringTracerouteDetailsPercentilesPacketLossPctJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTracerouteDetailsPercentilesPacketLossPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTracerouteDetailsPercentilesPacketLossPctJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTracerouteDetailsPercentilesRoundTripTimeMs struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                                                    `json:"p99,nullable"`
	JSON digitalExperienceMonitoringTracerouteDetailsPercentilesRoundTripTimeMsJSON `json:"-"`
}

// digitalExperienceMonitoringTracerouteDetailsPercentilesRoundTripTimeMsJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTracerouteDetailsPercentilesRoundTripTimeMs]
type digitalExperienceMonitoringTracerouteDetailsPercentilesRoundTripTimeMsJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTracerouteDetailsPercentilesRoundTripTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTracerouteDetailsPercentilesRoundTripTimeMsJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTracerouteTestNetworkPath struct {
	// API Resource UUID tag.
	ID         string `json:"id,required"`
	DeviceName string `json:"deviceName"`
	// The interval at which the Traceroute synthetic application test is set to run.
	Interval    string                                                          `json:"interval"`
	Kind        DigitalExperienceMonitoringTracerouteTestNetworkPathKind        `json:"kind"`
	Name        string                                                          `json:"name"`
	NetworkPath DigitalExperienceMonitoringTracerouteTestNetworkPathNetworkPath `json:"networkPath,nullable"`
	// The host of the Traceroute synthetic application test
	URL  string                                                   `json:"url"`
	JSON digitalExperienceMonitoringTracerouteTestNetworkPathJSON `json:"-"`
}

// digitalExperienceMonitoringTracerouteTestNetworkPathJSON contains the JSON
// metadata for the struct [DigitalExperienceMonitoringTracerouteTestNetworkPath]
type digitalExperienceMonitoringTracerouteTestNetworkPathJSON struct {
	ID          apijson.Field
	DeviceName  apijson.Field
	Interval    apijson.Field
	Kind        apijson.Field
	Name        apijson.Field
	NetworkPath apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTracerouteTestNetworkPath) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTracerouteTestNetworkPathJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTracerouteTestNetworkPathKind string

const (
	DigitalExperienceMonitoringTracerouteTestNetworkPathKindTraceroute DigitalExperienceMonitoringTracerouteTestNetworkPathKind = "traceroute"
)

func (r DigitalExperienceMonitoringTracerouteTestNetworkPathKind) IsKnown() bool {
	switch r {
	case DigitalExperienceMonitoringTracerouteTestNetworkPathKindTraceroute:
		return true
	}
	return false
}

type DigitalExperienceMonitoringTracerouteTestNetworkPathNetworkPath struct {
	Slots []DigitalExperienceMonitoringTracerouteTestNetworkPathNetworkPathSlot `json:"slots,required"`
	// Specifies the sampling applied, if any, to the slots response. When sampled,
	// results shown represent the first test run to the start of each sampling
	// interval.
	Sampling DigitalExperienceMonitoringTracerouteTestNetworkPathNetworkPathSampling `json:"sampling,nullable"`
	JSON     digitalExperienceMonitoringTracerouteTestNetworkPathNetworkPathJSON     `json:"-"`
}

// digitalExperienceMonitoringTracerouteTestNetworkPathNetworkPathJSON contains the
// JSON metadata for the struct
// [DigitalExperienceMonitoringTracerouteTestNetworkPathNetworkPath]
type digitalExperienceMonitoringTracerouteTestNetworkPathNetworkPathJSON struct {
	Slots       apijson.Field
	Sampling    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTracerouteTestNetworkPathNetworkPath) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTracerouteTestNetworkPathNetworkPathJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTracerouteTestNetworkPathNetworkPathSlot struct {
	// API Resource UUID tag.
	ID string `json:"id,required"`
	// Round trip time in ms of the client to app mile
	ClientToAppRTTMs int64 `json:"clientToAppRttMs,required,nullable"`
	// Round trip time in ms of the client to Cloudflare egress mile
	ClientToCfEgressRTTMs int64 `json:"clientToCfEgressRttMs,required,nullable"`
	// Round trip time in ms of the client to Cloudflare ingress mile
	ClientToCfIngressRTTMs int64  `json:"clientToCfIngressRttMs,required,nullable"`
	Timestamp              string `json:"timestamp,required"`
	// Round trip time in ms of the client to ISP mile
	ClientToIspRTTMs int64                                                                   `json:"clientToIspRttMs,nullable"`
	JSON             digitalExperienceMonitoringTracerouteTestNetworkPathNetworkPathSlotJSON `json:"-"`
}

// digitalExperienceMonitoringTracerouteTestNetworkPathNetworkPathSlotJSON contains
// the JSON metadata for the struct
// [DigitalExperienceMonitoringTracerouteTestNetworkPathNetworkPathSlot]
type digitalExperienceMonitoringTracerouteTestNetworkPathNetworkPathSlotJSON struct {
	ID                     apijson.Field
	ClientToAppRTTMs       apijson.Field
	ClientToCfEgressRTTMs  apijson.Field
	ClientToCfIngressRTTMs apijson.Field
	Timestamp              apijson.Field
	ClientToIspRTTMs       apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTracerouteTestNetworkPathNetworkPathSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTracerouteTestNetworkPathNetworkPathSlotJSON) RawJSON() string {
	return r.raw
}

// Specifies the sampling applied, if any, to the slots response. When sampled,
// results shown represent the first test run to the start of each sampling
// interval.
type DigitalExperienceMonitoringTracerouteTestNetworkPathNetworkPathSampling struct {
	Unit  DigitalExperienceMonitoringTracerouteTestNetworkPathNetworkPathSamplingUnit `json:"unit,required"`
	Value int64                                                                       `json:"value,required"`
	JSON  digitalExperienceMonitoringTracerouteTestNetworkPathNetworkPathSamplingJSON `json:"-"`
}

// digitalExperienceMonitoringTracerouteTestNetworkPathNetworkPathSamplingJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTracerouteTestNetworkPathNetworkPathSampling]
type digitalExperienceMonitoringTracerouteTestNetworkPathNetworkPathSamplingJSON struct {
	Unit        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTracerouteTestNetworkPathNetworkPathSampling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTracerouteTestNetworkPathNetworkPathSamplingJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTracerouteTestNetworkPathNetworkPathSamplingUnit string

const (
	DigitalExperienceMonitoringTracerouteTestNetworkPathNetworkPathSamplingUnitHours DigitalExperienceMonitoringTracerouteTestNetworkPathNetworkPathSamplingUnit = "hours"
)

func (r DigitalExperienceMonitoringTracerouteTestNetworkPathNetworkPathSamplingUnit) IsKnown() bool {
	switch r {
	case DigitalExperienceMonitoringTracerouteTestNetworkPathNetworkPathSamplingUnitHours:
		return true
	}
	return false
}

type DEXTracerouteTestGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Time interval for aggregate time slots.
	Interval param.Field[DEXTracerouteTestGetParamsInterval] `query:"interval,required"`
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

// URLQuery serializes [DEXTracerouteTestGetParams]'s query parameters as
// `url.Values`.
func (r DEXTracerouteTestGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Time interval for aggregate time slots.
type DEXTracerouteTestGetParamsInterval string

const (
	DEXTracerouteTestGetParamsIntervalMinute DEXTracerouteTestGetParamsInterval = "minute"
	DEXTracerouteTestGetParamsIntervalHour   DEXTracerouteTestGetParamsInterval = "hour"
)

func (r DEXTracerouteTestGetParamsInterval) IsKnown() bool {
	switch r {
	case DEXTracerouteTestGetParamsIntervalMinute, DEXTracerouteTestGetParamsIntervalHour:
		return true
	}
	return false
}

type DEXTracerouteTestGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   DigitalExperienceMonitoringTracerouteDetails              `json:"result,required"`
	// Whether the API call was successful
	Success DEXTracerouteTestGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dexTracerouteTestGetResponseEnvelopeJSON    `json:"-"`
}

// dexTracerouteTestGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DEXTracerouteTestGetResponseEnvelope]
type dexTracerouteTestGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DEXTracerouteTestGetResponseEnvelopeSuccess bool

const (
	DEXTracerouteTestGetResponseEnvelopeSuccessTrue DEXTracerouteTestGetResponseEnvelopeSuccess = true
)

func (r DEXTracerouteTestGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DEXTracerouteTestGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DEXTracerouteTestNetworkPathParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Device to filter tracroute result runs to
	DeviceID param.Field[string] `query:"deviceId,required"`
	// Time interval for aggregate time slots.
	Interval param.Field[DEXTracerouteTestNetworkPathParamsInterval] `query:"interval,required"`
	// End time for aggregate metrics in ISO ms
	TimeEnd param.Field[string] `query:"timeEnd,required"`
	// Start time for aggregate metrics in ISO ms
	TimeStart param.Field[string] `query:"timeStart,required"`
}

// URLQuery serializes [DEXTracerouteTestNetworkPathParams]'s query parameters as
// `url.Values`.
func (r DEXTracerouteTestNetworkPathParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Time interval for aggregate time slots.
type DEXTracerouteTestNetworkPathParamsInterval string

const (
	DEXTracerouteTestNetworkPathParamsIntervalMinute DEXTracerouteTestNetworkPathParamsInterval = "minute"
	DEXTracerouteTestNetworkPathParamsIntervalHour   DEXTracerouteTestNetworkPathParamsInterval = "hour"
)

func (r DEXTracerouteTestNetworkPathParamsInterval) IsKnown() bool {
	switch r {
	case DEXTracerouteTestNetworkPathParamsIntervalMinute, DEXTracerouteTestNetworkPathParamsIntervalHour:
		return true
	}
	return false
}

type DEXTracerouteTestNetworkPathResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   DigitalExperienceMonitoringTracerouteTestNetworkPath      `json:"result,required"`
	// Whether the API call was successful
	Success DEXTracerouteTestNetworkPathResponseEnvelopeSuccess `json:"success,required"`
	JSON    dexTracerouteTestNetworkPathResponseEnvelopeJSON    `json:"-"`
}

// dexTracerouteTestNetworkPathResponseEnvelopeJSON contains the JSON metadata for
// the struct [DEXTracerouteTestNetworkPathResponseEnvelope]
type dexTracerouteTestNetworkPathResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestNetworkPathResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestNetworkPathResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DEXTracerouteTestNetworkPathResponseEnvelopeSuccess bool

const (
	DEXTracerouteTestNetworkPathResponseEnvelopeSuccessTrue DEXTracerouteTestNetworkPathResponseEnvelopeSuccess = true
)

func (r DEXTracerouteTestNetworkPathResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DEXTracerouteTestNetworkPathResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DEXTracerouteTestPercentilesParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// End time for aggregate metrics in ISO format
	TimeEnd param.Field[string] `query:"timeEnd,required"`
	// Start time for aggregate metrics in ISO format
	TimeStart param.Field[string] `query:"timeStart,required"`
	// Optionally filter result stats to a Cloudflare colo. Cannot be used in
	// combination with deviceId param.
	Colo param.Field[string] `query:"colo"`
	// Optionally filter result stats to a specific device(s). Cannot be used in
	// combination with colo param.
	DeviceID param.Field[[]string] `query:"deviceId"`
}

// URLQuery serializes [DEXTracerouteTestPercentilesParams]'s query parameters as
// `url.Values`.
func (r DEXTracerouteTestPercentilesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DEXTracerouteTestPercentilesResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   DigitalExperienceMonitoringTracerouteDetailsPercentiles   `json:"result,required"`
	// Whether the API call was successful
	Success DEXTracerouteTestPercentilesResponseEnvelopeSuccess `json:"success,required"`
	JSON    dexTracerouteTestPercentilesResponseEnvelopeJSON    `json:"-"`
}

// dexTracerouteTestPercentilesResponseEnvelopeJSON contains the JSON metadata for
// the struct [DEXTracerouteTestPercentilesResponseEnvelope]
type dexTracerouteTestPercentilesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestPercentilesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestPercentilesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DEXTracerouteTestPercentilesResponseEnvelopeSuccess bool

const (
	DEXTracerouteTestPercentilesResponseEnvelopeSuccessTrue DEXTracerouteTestPercentilesResponseEnvelopeSuccess = true
)

func (r DEXTracerouteTestPercentilesResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DEXTracerouteTestPercentilesResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
