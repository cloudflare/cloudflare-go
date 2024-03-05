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

// ZeroTrustDEXTracerouteTestService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZeroTrustDEXTracerouteTestService] method instead.
type ZeroTrustDEXTracerouteTestService struct {
	Options []option.RequestOption
}

// NewZeroTrustDEXTracerouteTestService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustDEXTracerouteTestService(opts ...option.RequestOption) (r *ZeroTrustDEXTracerouteTestService) {
	r = &ZeroTrustDEXTracerouteTestService{}
	r.Options = opts
	return
}

// Get test details and aggregate performance metrics for an traceroute test for a
// given time period between 1 hour and 7 days.
func (r *ZeroTrustDEXTracerouteTestService) Get(ctx context.Context, testID string, params ZeroTrustDEXTracerouteTestGetParams, opts ...option.RequestOption) (res *DigitalExperienceMonitoringTracerouteDetails, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDEXTracerouteTestGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/traceroute-tests/%s", params.AccountID, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a breakdown of metrics by hop for individual traceroute test runs
func (r *ZeroTrustDEXTracerouteTestService) NetworkPath(ctx context.Context, testID string, params ZeroTrustDEXTracerouteTestNetworkPathParams, opts ...option.RequestOption) (res *DigitalExperienceMonitoringTracerouteTestNetworkPath, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDEXTracerouteTestNetworkPathResponseEnvelope
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
func (r *ZeroTrustDEXTracerouteTestService) Percentiles(ctx context.Context, testID string, params ZeroTrustDEXTracerouteTestPercentilesParams, opts ...option.RequestOption) (res *DigitalExperienceMonitoringTracerouteDetailsPercentiles, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDEXTracerouteTestPercentilesResponseEnvelope
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
	TracerouteStats       apijson.Field
	TracerouteStatsByColo apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTracerouteDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DigitalExperienceMonitoringTracerouteDetailsKind string

const (
	DigitalExperienceMonitoringTracerouteDetailsKindTraceroute DigitalExperienceMonitoringTracerouteDetailsKind = "traceroute"
)

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

type DigitalExperienceMonitoringTracerouteTestNetworkPathKind string

const (
	DigitalExperienceMonitoringTracerouteTestNetworkPathKindTraceroute DigitalExperienceMonitoringTracerouteTestNetworkPathKind = "traceroute"
)

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

type DigitalExperienceMonitoringTracerouteTestNetworkPathNetworkPathSamplingUnit string

const (
	DigitalExperienceMonitoringTracerouteTestNetworkPathNetworkPathSamplingUnitHours DigitalExperienceMonitoringTracerouteTestNetworkPathNetworkPathSamplingUnit = "hours"
)

type ZeroTrustDEXTracerouteTestGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Time interval for aggregate time slots.
	Interval param.Field[ZeroTrustDEXTracerouteTestGetParamsInterval] `query:"interval,required"`
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

// URLQuery serializes [ZeroTrustDEXTracerouteTestGetParams]'s query parameters as
// `url.Values`.
func (r ZeroTrustDEXTracerouteTestGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Time interval for aggregate time slots.
type ZeroTrustDEXTracerouteTestGetParamsInterval string

const (
	ZeroTrustDEXTracerouteTestGetParamsIntervalMinute ZeroTrustDEXTracerouteTestGetParamsInterval = "minute"
	ZeroTrustDEXTracerouteTestGetParamsIntervalHour   ZeroTrustDEXTracerouteTestGetParamsInterval = "hour"
)

type ZeroTrustDEXTracerouteTestGetResponseEnvelope struct {
	Errors   []ZeroTrustDEXTracerouteTestGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDEXTracerouteTestGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DigitalExperienceMonitoringTracerouteDetails            `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustDEXTracerouteTestGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDEXTracerouteTestGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDEXTracerouteTestGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustDEXTracerouteTestGetResponseEnvelope]
type zeroTrustDEXTracerouteTestGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestGetResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zeroTrustDEXTracerouteTestGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustDEXTracerouteTestGetResponseEnvelopeErrors]
type zeroTrustDEXTracerouteTestGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestGetResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zeroTrustDEXTracerouteTestGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustDEXTracerouteTestGetResponseEnvelopeMessages]
type zeroTrustDEXTracerouteTestGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustDEXTracerouteTestGetResponseEnvelopeSuccess bool

const (
	ZeroTrustDEXTracerouteTestGetResponseEnvelopeSuccessTrue ZeroTrustDEXTracerouteTestGetResponseEnvelopeSuccess = true
)

type ZeroTrustDEXTracerouteTestNetworkPathParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Device to filter tracroute result runs to
	DeviceID param.Field[string] `query:"deviceId,required"`
	// Time interval for aggregate time slots.
	Interval param.Field[ZeroTrustDEXTracerouteTestNetworkPathParamsInterval] `query:"interval,required"`
	// End time for aggregate metrics in ISO ms
	TimeEnd param.Field[string] `query:"timeEnd,required"`
	// Start time for aggregate metrics in ISO ms
	TimeStart param.Field[string] `query:"timeStart,required"`
}

// URLQuery serializes [ZeroTrustDEXTracerouteTestNetworkPathParams]'s query
// parameters as `url.Values`.
func (r ZeroTrustDEXTracerouteTestNetworkPathParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Time interval for aggregate time slots.
type ZeroTrustDEXTracerouteTestNetworkPathParamsInterval string

const (
	ZeroTrustDEXTracerouteTestNetworkPathParamsIntervalMinute ZeroTrustDEXTracerouteTestNetworkPathParamsInterval = "minute"
	ZeroTrustDEXTracerouteTestNetworkPathParamsIntervalHour   ZeroTrustDEXTracerouteTestNetworkPathParamsInterval = "hour"
)

type ZeroTrustDEXTracerouteTestNetworkPathResponseEnvelope struct {
	Errors   []ZeroTrustDEXTracerouteTestNetworkPathResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDEXTracerouteTestNetworkPathResponseEnvelopeMessages `json:"messages,required"`
	Result   DigitalExperienceMonitoringTracerouteTestNetworkPath            `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustDEXTracerouteTestNetworkPathResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDEXTracerouteTestNetworkPathResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDEXTracerouteTestNetworkPathResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustDEXTracerouteTestNetworkPathResponseEnvelope]
type zeroTrustDEXTracerouteTestNetworkPathResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestNetworkPathResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestNetworkPathResponseEnvelopeErrors struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    zeroTrustDEXTracerouteTestNetworkPathResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestNetworkPathResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [ZeroTrustDEXTracerouteTestNetworkPathResponseEnvelopeErrors]
type zeroTrustDEXTracerouteTestNetworkPathResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestNetworkPathResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestNetworkPathResponseEnvelopeMessages struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    zeroTrustDEXTracerouteTestNetworkPathResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestNetworkPathResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZeroTrustDEXTracerouteTestNetworkPathResponseEnvelopeMessages]
type zeroTrustDEXTracerouteTestNetworkPathResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestNetworkPathResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustDEXTracerouteTestNetworkPathResponseEnvelopeSuccess bool

const (
	ZeroTrustDEXTracerouteTestNetworkPathResponseEnvelopeSuccessTrue ZeroTrustDEXTracerouteTestNetworkPathResponseEnvelopeSuccess = true
)

type ZeroTrustDEXTracerouteTestPercentilesParams struct {
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

// URLQuery serializes [ZeroTrustDEXTracerouteTestPercentilesParams]'s query
// parameters as `url.Values`.
func (r ZeroTrustDEXTracerouteTestPercentilesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZeroTrustDEXTracerouteTestPercentilesResponseEnvelope struct {
	Errors   []ZeroTrustDEXTracerouteTestPercentilesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDEXTracerouteTestPercentilesResponseEnvelopeMessages `json:"messages,required"`
	Result   DigitalExperienceMonitoringTracerouteDetailsPercentiles         `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustDEXTracerouteTestPercentilesResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDEXTracerouteTestPercentilesResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDEXTracerouteTestPercentilesResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustDEXTracerouteTestPercentilesResponseEnvelope]
type zeroTrustDEXTracerouteTestPercentilesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestPercentilesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestPercentilesResponseEnvelopeErrors struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    zeroTrustDEXTracerouteTestPercentilesResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestPercentilesResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [ZeroTrustDEXTracerouteTestPercentilesResponseEnvelopeErrors]
type zeroTrustDEXTracerouteTestPercentilesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestPercentilesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestPercentilesResponseEnvelopeMessages struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    zeroTrustDEXTracerouteTestPercentilesResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestPercentilesResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZeroTrustDEXTracerouteTestPercentilesResponseEnvelopeMessages]
type zeroTrustDEXTracerouteTestPercentilesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestPercentilesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustDEXTracerouteTestPercentilesResponseEnvelopeSuccess bool

const (
	ZeroTrustDEXTracerouteTestPercentilesResponseEnvelopeSuccessTrue ZeroTrustDEXTracerouteTestPercentilesResponseEnvelopeSuccess = true
)
