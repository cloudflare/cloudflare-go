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
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *DEXTracerouteTestService) Get(ctx context.Context, testID string, params DEXTracerouteTestGetParams, opts ...option.RequestOption) (res *DEXTracerouteTestGetResponse, err error) {
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
func (r *DEXTracerouteTestService) NetworkPath(ctx context.Context, testID string, params DEXTracerouteTestNetworkPathParams, opts ...option.RequestOption) (res *DEXTracerouteTestNetworkPathResponse, err error) {
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
func (r *DEXTracerouteTestService) Percentiles(ctx context.Context, testID string, params DEXTracerouteTestPercentilesParams, opts ...option.RequestOption) (res *DEXTracerouteTestPercentilesResponse, err error) {
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

type DEXTracerouteTestGetResponse struct {
	// The host of the Traceroute synthetic application test
	Host string `json:"host,required"`
	// The interval at which the Traceroute synthetic application test is set to run.
	Interval string                           `json:"interval,required"`
	Kind     DEXTracerouteTestGetResponseKind `json:"kind,required"`
	// The name of the Traceroute synthetic application test
	Name                  string                                              `json:"name,required"`
	TracerouteStats       DEXTracerouteTestGetResponseTracerouteStats         `json:"tracerouteStats,nullable"`
	TracerouteStatsByColo []DEXTracerouteTestGetResponseTracerouteStatsByColo `json:"tracerouteStatsByColo"`
	JSON                  dexTracerouteTestGetResponseJSON                    `json:"-"`
}

// dexTracerouteTestGetResponseJSON contains the JSON metadata for the struct
// [DEXTracerouteTestGetResponse]
type dexTracerouteTestGetResponseJSON struct {
	Host                  apijson.Field
	Interval              apijson.Field
	Kind                  apijson.Field
	Name                  apijson.Field
	TracerouteStats       apijson.Field
	TracerouteStatsByColo apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *DEXTracerouteTestGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestGetResponseJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestGetResponseKind string

const (
	DEXTracerouteTestGetResponseKindTraceroute DEXTracerouteTestGetResponseKind = "traceroute"
)

type DEXTracerouteTestGetResponseTracerouteStats struct {
	AvailabilityPct DEXTracerouteTestGetResponseTracerouteStatsAvailabilityPct `json:"availabilityPct,required"`
	HopsCount       DEXTracerouteTestGetResponseTracerouteStatsHopsCount       `json:"hopsCount,required"`
	PacketLossPct   DEXTracerouteTestGetResponseTracerouteStatsPacketLossPct   `json:"packetLossPct,required"`
	RoundTripTimeMs DEXTracerouteTestGetResponseTracerouteStatsRoundTripTimeMs `json:"roundTripTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                                           `json:"uniqueDevicesTotal,required"`
	JSON               dexTracerouteTestGetResponseTracerouteStatsJSON `json:"-"`
}

// dexTracerouteTestGetResponseTracerouteStatsJSON contains the JSON metadata for
// the struct [DEXTracerouteTestGetResponseTracerouteStats]
type dexTracerouteTestGetResponseTracerouteStatsJSON struct {
	AvailabilityPct    apijson.Field
	HopsCount          apijson.Field
	PacketLossPct      apijson.Field
	RoundTripTimeMs    apijson.Field
	UniqueDevicesTotal apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DEXTracerouteTestGetResponseTracerouteStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestGetResponseTracerouteStatsJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestGetResponseTracerouteStatsAvailabilityPct struct {
	Slots []DEXTracerouteTestGetResponseTracerouteStatsAvailabilityPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                                        `json:"min,nullable"`
	JSON dexTracerouteTestGetResponseTracerouteStatsAvailabilityPctJSON `json:"-"`
}

// dexTracerouteTestGetResponseTracerouteStatsAvailabilityPctJSON contains the JSON
// metadata for the struct
// [DEXTracerouteTestGetResponseTracerouteStatsAvailabilityPct]
type dexTracerouteTestGetResponseTracerouteStatsAvailabilityPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestGetResponseTracerouteStatsAvailabilityPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestGetResponseTracerouteStatsAvailabilityPctJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestGetResponseTracerouteStatsAvailabilityPctSlot struct {
	Timestamp string                                                             `json:"timestamp,required"`
	Value     float64                                                            `json:"value,required"`
	JSON      dexTracerouteTestGetResponseTracerouteStatsAvailabilityPctSlotJSON `json:"-"`
}

// dexTracerouteTestGetResponseTracerouteStatsAvailabilityPctSlotJSON contains the
// JSON metadata for the struct
// [DEXTracerouteTestGetResponseTracerouteStatsAvailabilityPctSlot]
type dexTracerouteTestGetResponseTracerouteStatsAvailabilityPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestGetResponseTracerouteStatsAvailabilityPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestGetResponseTracerouteStatsAvailabilityPctSlotJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestGetResponseTracerouteStatsHopsCount struct {
	Slots []DEXTracerouteTestGetResponseTracerouteStatsHopsCountSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                    `json:"min,nullable"`
	JSON dexTracerouteTestGetResponseTracerouteStatsHopsCountJSON `json:"-"`
}

// dexTracerouteTestGetResponseTracerouteStatsHopsCountJSON contains the JSON
// metadata for the struct [DEXTracerouteTestGetResponseTracerouteStatsHopsCount]
type dexTracerouteTestGetResponseTracerouteStatsHopsCountJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestGetResponseTracerouteStatsHopsCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestGetResponseTracerouteStatsHopsCountJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestGetResponseTracerouteStatsHopsCountSlot struct {
	Timestamp string                                                       `json:"timestamp,required"`
	Value     int64                                                        `json:"value,required"`
	JSON      dexTracerouteTestGetResponseTracerouteStatsHopsCountSlotJSON `json:"-"`
}

// dexTracerouteTestGetResponseTracerouteStatsHopsCountSlotJSON contains the JSON
// metadata for the struct
// [DEXTracerouteTestGetResponseTracerouteStatsHopsCountSlot]
type dexTracerouteTestGetResponseTracerouteStatsHopsCountSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestGetResponseTracerouteStatsHopsCountSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestGetResponseTracerouteStatsHopsCountSlotJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestGetResponseTracerouteStatsPacketLossPct struct {
	Slots []DEXTracerouteTestGetResponseTracerouteStatsPacketLossPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                                      `json:"min,nullable"`
	JSON dexTracerouteTestGetResponseTracerouteStatsPacketLossPctJSON `json:"-"`
}

// dexTracerouteTestGetResponseTracerouteStatsPacketLossPctJSON contains the JSON
// metadata for the struct
// [DEXTracerouteTestGetResponseTracerouteStatsPacketLossPct]
type dexTracerouteTestGetResponseTracerouteStatsPacketLossPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestGetResponseTracerouteStatsPacketLossPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestGetResponseTracerouteStatsPacketLossPctJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestGetResponseTracerouteStatsPacketLossPctSlot struct {
	Timestamp string                                                           `json:"timestamp,required"`
	Value     float64                                                          `json:"value,required"`
	JSON      dexTracerouteTestGetResponseTracerouteStatsPacketLossPctSlotJSON `json:"-"`
}

// dexTracerouteTestGetResponseTracerouteStatsPacketLossPctSlotJSON contains the
// JSON metadata for the struct
// [DEXTracerouteTestGetResponseTracerouteStatsPacketLossPctSlot]
type dexTracerouteTestGetResponseTracerouteStatsPacketLossPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestGetResponseTracerouteStatsPacketLossPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestGetResponseTracerouteStatsPacketLossPctSlotJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestGetResponseTracerouteStatsRoundTripTimeMs struct {
	Slots []DEXTracerouteTestGetResponseTracerouteStatsRoundTripTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                          `json:"min,nullable"`
	JSON dexTracerouteTestGetResponseTracerouteStatsRoundTripTimeMsJSON `json:"-"`
}

// dexTracerouteTestGetResponseTracerouteStatsRoundTripTimeMsJSON contains the JSON
// metadata for the struct
// [DEXTracerouteTestGetResponseTracerouteStatsRoundTripTimeMs]
type dexTracerouteTestGetResponseTracerouteStatsRoundTripTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestGetResponseTracerouteStatsRoundTripTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestGetResponseTracerouteStatsRoundTripTimeMsJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestGetResponseTracerouteStatsRoundTripTimeMsSlot struct {
	Timestamp string                                                             `json:"timestamp,required"`
	Value     int64                                                              `json:"value,required"`
	JSON      dexTracerouteTestGetResponseTracerouteStatsRoundTripTimeMsSlotJSON `json:"-"`
}

// dexTracerouteTestGetResponseTracerouteStatsRoundTripTimeMsSlotJSON contains the
// JSON metadata for the struct
// [DEXTracerouteTestGetResponseTracerouteStatsRoundTripTimeMsSlot]
type dexTracerouteTestGetResponseTracerouteStatsRoundTripTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestGetResponseTracerouteStatsRoundTripTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestGetResponseTracerouteStatsRoundTripTimeMsSlotJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestGetResponseTracerouteStatsByColo struct {
	AvailabilityPct DEXTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPct `json:"availabilityPct,required"`
	Colo            string                                                           `json:"colo,required"`
	HopsCount       DEXTracerouteTestGetResponseTracerouteStatsByColoHopsCount       `json:"hopsCount,required"`
	PacketLossPct   DEXTracerouteTestGetResponseTracerouteStatsByColoPacketLossPct   `json:"packetLossPct,required"`
	RoundTripTimeMs DEXTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMs `json:"roundTripTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                                                 `json:"uniqueDevicesTotal,required"`
	JSON               dexTracerouteTestGetResponseTracerouteStatsByColoJSON `json:"-"`
}

// dexTracerouteTestGetResponseTracerouteStatsByColoJSON contains the JSON metadata
// for the struct [DEXTracerouteTestGetResponseTracerouteStatsByColo]
type dexTracerouteTestGetResponseTracerouteStatsByColoJSON struct {
	AvailabilityPct    apijson.Field
	Colo               apijson.Field
	HopsCount          apijson.Field
	PacketLossPct      apijson.Field
	RoundTripTimeMs    apijson.Field
	UniqueDevicesTotal apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DEXTracerouteTestGetResponseTracerouteStatsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestGetResponseTracerouteStatsByColoJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPct struct {
	Slots []DEXTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                                              `json:"min,nullable"`
	JSON dexTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPctJSON `json:"-"`
}

// dexTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPctJSON contains
// the JSON metadata for the struct
// [DEXTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPct]
type dexTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPctJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPctSlot struct {
	Timestamp string                                                                   `json:"timestamp,required"`
	Value     float64                                                                  `json:"value,required"`
	JSON      dexTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPctSlotJSON `json:"-"`
}

// dexTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPctSlotJSON
// contains the JSON metadata for the struct
// [DEXTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPctSlot]
type dexTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPctSlotJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestGetResponseTracerouteStatsByColoHopsCount struct {
	Slots []DEXTracerouteTestGetResponseTracerouteStatsByColoHopsCountSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                          `json:"min,nullable"`
	JSON dexTracerouteTestGetResponseTracerouteStatsByColoHopsCountJSON `json:"-"`
}

// dexTracerouteTestGetResponseTracerouteStatsByColoHopsCountJSON contains the JSON
// metadata for the struct
// [DEXTracerouteTestGetResponseTracerouteStatsByColoHopsCount]
type dexTracerouteTestGetResponseTracerouteStatsByColoHopsCountJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestGetResponseTracerouteStatsByColoHopsCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestGetResponseTracerouteStatsByColoHopsCountJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestGetResponseTracerouteStatsByColoHopsCountSlot struct {
	Timestamp string                                                             `json:"timestamp,required"`
	Value     int64                                                              `json:"value,required"`
	JSON      dexTracerouteTestGetResponseTracerouteStatsByColoHopsCountSlotJSON `json:"-"`
}

// dexTracerouteTestGetResponseTracerouteStatsByColoHopsCountSlotJSON contains the
// JSON metadata for the struct
// [DEXTracerouteTestGetResponseTracerouteStatsByColoHopsCountSlot]
type dexTracerouteTestGetResponseTracerouteStatsByColoHopsCountSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestGetResponseTracerouteStatsByColoHopsCountSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestGetResponseTracerouteStatsByColoHopsCountSlotJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestGetResponseTracerouteStatsByColoPacketLossPct struct {
	Slots []DEXTracerouteTestGetResponseTracerouteStatsByColoPacketLossPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                                            `json:"min,nullable"`
	JSON dexTracerouteTestGetResponseTracerouteStatsByColoPacketLossPctJSON `json:"-"`
}

// dexTracerouteTestGetResponseTracerouteStatsByColoPacketLossPctJSON contains the
// JSON metadata for the struct
// [DEXTracerouteTestGetResponseTracerouteStatsByColoPacketLossPct]
type dexTracerouteTestGetResponseTracerouteStatsByColoPacketLossPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestGetResponseTracerouteStatsByColoPacketLossPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestGetResponseTracerouteStatsByColoPacketLossPctJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestGetResponseTracerouteStatsByColoPacketLossPctSlot struct {
	Timestamp string                                                                 `json:"timestamp,required"`
	Value     float64                                                                `json:"value,required"`
	JSON      dexTracerouteTestGetResponseTracerouteStatsByColoPacketLossPctSlotJSON `json:"-"`
}

// dexTracerouteTestGetResponseTracerouteStatsByColoPacketLossPctSlotJSON contains
// the JSON metadata for the struct
// [DEXTracerouteTestGetResponseTracerouteStatsByColoPacketLossPctSlot]
type dexTracerouteTestGetResponseTracerouteStatsByColoPacketLossPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestGetResponseTracerouteStatsByColoPacketLossPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestGetResponseTracerouteStatsByColoPacketLossPctSlotJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMs struct {
	Slots []DEXTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                                `json:"min,nullable"`
	JSON dexTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMsJSON `json:"-"`
}

// dexTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMsJSON contains
// the JSON metadata for the struct
// [DEXTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMs]
type dexTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMsJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMsSlot struct {
	Timestamp string                                                                   `json:"timestamp,required"`
	Value     int64                                                                    `json:"value,required"`
	JSON      dexTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMsSlotJSON `json:"-"`
}

// dexTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMsSlotJSON
// contains the JSON metadata for the struct
// [DEXTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMsSlot]
type dexTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMsSlotJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestNetworkPathResponse struct {
	// API Resource UUID tag.
	ID         string `json:"id,required"`
	DeviceName string `json:"deviceName"`
	// The interval at which the Traceroute synthetic application test is set to run.
	Interval    string                                          `json:"interval"`
	Kind        DEXTracerouteTestNetworkPathResponseKind        `json:"kind"`
	Name        string                                          `json:"name"`
	NetworkPath DEXTracerouteTestNetworkPathResponseNetworkPath `json:"networkPath,nullable"`
	// The host of the Traceroute synthetic application test
	URL  string                                   `json:"url"`
	JSON dexTracerouteTestNetworkPathResponseJSON `json:"-"`
}

// dexTracerouteTestNetworkPathResponseJSON contains the JSON metadata for the
// struct [DEXTracerouteTestNetworkPathResponse]
type dexTracerouteTestNetworkPathResponseJSON struct {
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

func (r *DEXTracerouteTestNetworkPathResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestNetworkPathResponseJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestNetworkPathResponseKind string

const (
	DEXTracerouteTestNetworkPathResponseKindTraceroute DEXTracerouteTestNetworkPathResponseKind = "traceroute"
)

type DEXTracerouteTestNetworkPathResponseNetworkPath struct {
	Slots []DEXTracerouteTestNetworkPathResponseNetworkPathSlot `json:"slots,required"`
	// Specifies the sampling applied, if any, to the slots response. When sampled,
	// results shown represent the first test run to the start of each sampling
	// interval.
	Sampling DEXTracerouteTestNetworkPathResponseNetworkPathSampling `json:"sampling,nullable"`
	JSON     dexTracerouteTestNetworkPathResponseNetworkPathJSON     `json:"-"`
}

// dexTracerouteTestNetworkPathResponseNetworkPathJSON contains the JSON metadata
// for the struct [DEXTracerouteTestNetworkPathResponseNetworkPath]
type dexTracerouteTestNetworkPathResponseNetworkPathJSON struct {
	Slots       apijson.Field
	Sampling    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestNetworkPathResponseNetworkPath) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestNetworkPathResponseNetworkPathJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestNetworkPathResponseNetworkPathSlot struct {
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
	ClientToIspRTTMs int64                                                   `json:"clientToIspRttMs,nullable"`
	JSON             dexTracerouteTestNetworkPathResponseNetworkPathSlotJSON `json:"-"`
}

// dexTracerouteTestNetworkPathResponseNetworkPathSlotJSON contains the JSON
// metadata for the struct [DEXTracerouteTestNetworkPathResponseNetworkPathSlot]
type dexTracerouteTestNetworkPathResponseNetworkPathSlotJSON struct {
	ID                     apijson.Field
	ClientToAppRTTMs       apijson.Field
	ClientToCfEgressRTTMs  apijson.Field
	ClientToCfIngressRTTMs apijson.Field
	Timestamp              apijson.Field
	ClientToIspRTTMs       apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *DEXTracerouteTestNetworkPathResponseNetworkPathSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestNetworkPathResponseNetworkPathSlotJSON) RawJSON() string {
	return r.raw
}

// Specifies the sampling applied, if any, to the slots response. When sampled,
// results shown represent the first test run to the start of each sampling
// interval.
type DEXTracerouteTestNetworkPathResponseNetworkPathSampling struct {
	Unit  DEXTracerouteTestNetworkPathResponseNetworkPathSamplingUnit `json:"unit,required"`
	Value int64                                                       `json:"value,required"`
	JSON  dexTracerouteTestNetworkPathResponseNetworkPathSamplingJSON `json:"-"`
}

// dexTracerouteTestNetworkPathResponseNetworkPathSamplingJSON contains the JSON
// metadata for the struct
// [DEXTracerouteTestNetworkPathResponseNetworkPathSampling]
type dexTracerouteTestNetworkPathResponseNetworkPathSamplingJSON struct {
	Unit        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestNetworkPathResponseNetworkPathSampling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestNetworkPathResponseNetworkPathSamplingJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestNetworkPathResponseNetworkPathSamplingUnit string

const (
	DEXTracerouteTestNetworkPathResponseNetworkPathSamplingUnitHours DEXTracerouteTestNetworkPathResponseNetworkPathSamplingUnit = "hours"
)

type DEXTracerouteTestPercentilesResponse struct {
	HopsCount       DEXTracerouteTestPercentilesResponseHopsCount       `json:"hopsCount"`
	PacketLossPct   DEXTracerouteTestPercentilesResponsePacketLossPct   `json:"packetLossPct"`
	RoundTripTimeMs DEXTracerouteTestPercentilesResponseRoundTripTimeMs `json:"roundTripTimeMs"`
	JSON            dexTracerouteTestPercentilesResponseJSON            `json:"-"`
}

// dexTracerouteTestPercentilesResponseJSON contains the JSON metadata for the
// struct [DEXTracerouteTestPercentilesResponse]
type dexTracerouteTestPercentilesResponseJSON struct {
	HopsCount       apijson.Field
	PacketLossPct   apijson.Field
	RoundTripTimeMs apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DEXTracerouteTestPercentilesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestPercentilesResponseJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestPercentilesResponseHopsCount struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                           `json:"p99,nullable"`
	JSON dexTracerouteTestPercentilesResponseHopsCountJSON `json:"-"`
}

// dexTracerouteTestPercentilesResponseHopsCountJSON contains the JSON metadata for
// the struct [DEXTracerouteTestPercentilesResponseHopsCount]
type dexTracerouteTestPercentilesResponseHopsCountJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestPercentilesResponseHopsCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestPercentilesResponseHopsCountJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestPercentilesResponsePacketLossPct struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                               `json:"p99,nullable"`
	JSON dexTracerouteTestPercentilesResponsePacketLossPctJSON `json:"-"`
}

// dexTracerouteTestPercentilesResponsePacketLossPctJSON contains the JSON metadata
// for the struct [DEXTracerouteTestPercentilesResponsePacketLossPct]
type dexTracerouteTestPercentilesResponsePacketLossPctJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestPercentilesResponsePacketLossPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestPercentilesResponsePacketLossPctJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestPercentilesResponseRoundTripTimeMs struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                                 `json:"p99,nullable"`
	JSON dexTracerouteTestPercentilesResponseRoundTripTimeMsJSON `json:"-"`
}

// dexTracerouteTestPercentilesResponseRoundTripTimeMsJSON contains the JSON
// metadata for the struct [DEXTracerouteTestPercentilesResponseRoundTripTimeMs]
type dexTracerouteTestPercentilesResponseRoundTripTimeMsJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestPercentilesResponseRoundTripTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestPercentilesResponseRoundTripTimeMsJSON) RawJSON() string {
	return r.raw
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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Time interval for aggregate time slots.
type DEXTracerouteTestGetParamsInterval string

const (
	DEXTracerouteTestGetParamsIntervalMinute DEXTracerouteTestGetParamsInterval = "minute"
	DEXTracerouteTestGetParamsIntervalHour   DEXTracerouteTestGetParamsInterval = "hour"
)

type DEXTracerouteTestGetResponseEnvelope struct {
	Errors   []DEXTracerouteTestGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DEXTracerouteTestGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DEXTracerouteTestGetResponse                   `json:"result,required"`
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

type DEXTracerouteTestGetResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    dexTracerouteTestGetResponseEnvelopeErrorsJSON `json:"-"`
}

// dexTracerouteTestGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DEXTracerouteTestGetResponseEnvelopeErrors]
type dexTracerouteTestGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestGetResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    dexTracerouteTestGetResponseEnvelopeMessagesJSON `json:"-"`
}

// dexTracerouteTestGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DEXTracerouteTestGetResponseEnvelopeMessages]
type dexTracerouteTestGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DEXTracerouteTestGetResponseEnvelopeSuccess bool

const (
	DEXTracerouteTestGetResponseEnvelopeSuccessTrue DEXTracerouteTestGetResponseEnvelopeSuccess = true
)

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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Time interval for aggregate time slots.
type DEXTracerouteTestNetworkPathParamsInterval string

const (
	DEXTracerouteTestNetworkPathParamsIntervalMinute DEXTracerouteTestNetworkPathParamsInterval = "minute"
	DEXTracerouteTestNetworkPathParamsIntervalHour   DEXTracerouteTestNetworkPathParamsInterval = "hour"
)

type DEXTracerouteTestNetworkPathResponseEnvelope struct {
	Errors   []DEXTracerouteTestNetworkPathResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DEXTracerouteTestNetworkPathResponseEnvelopeMessages `json:"messages,required"`
	Result   DEXTracerouteTestNetworkPathResponse                   `json:"result,required"`
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

type DEXTracerouteTestNetworkPathResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    dexTracerouteTestNetworkPathResponseEnvelopeErrorsJSON `json:"-"`
}

// dexTracerouteTestNetworkPathResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DEXTracerouteTestNetworkPathResponseEnvelopeErrors]
type dexTracerouteTestNetworkPathResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestNetworkPathResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestNetworkPathResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestNetworkPathResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    dexTracerouteTestNetworkPathResponseEnvelopeMessagesJSON `json:"-"`
}

// dexTracerouteTestNetworkPathResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DEXTracerouteTestNetworkPathResponseEnvelopeMessages]
type dexTracerouteTestNetworkPathResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestNetworkPathResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestNetworkPathResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DEXTracerouteTestNetworkPathResponseEnvelopeSuccess bool

const (
	DEXTracerouteTestNetworkPathResponseEnvelopeSuccessTrue DEXTracerouteTestNetworkPathResponseEnvelopeSuccess = true
)

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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DEXTracerouteTestPercentilesResponseEnvelope struct {
	Errors   []DEXTracerouteTestPercentilesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DEXTracerouteTestPercentilesResponseEnvelopeMessages `json:"messages,required"`
	Result   DEXTracerouteTestPercentilesResponse                   `json:"result,required"`
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

type DEXTracerouteTestPercentilesResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    dexTracerouteTestPercentilesResponseEnvelopeErrorsJSON `json:"-"`
}

// dexTracerouteTestPercentilesResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DEXTracerouteTestPercentilesResponseEnvelopeErrors]
type dexTracerouteTestPercentilesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestPercentilesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestPercentilesResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestPercentilesResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    dexTracerouteTestPercentilesResponseEnvelopeMessagesJSON `json:"-"`
}

// dexTracerouteTestPercentilesResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DEXTracerouteTestPercentilesResponseEnvelopeMessages]
type dexTracerouteTestPercentilesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestPercentilesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestPercentilesResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DEXTracerouteTestPercentilesResponseEnvelopeSuccess bool

const (
	DEXTracerouteTestPercentilesResponseEnvelopeSuccessTrue DEXTracerouteTestPercentilesResponseEnvelopeSuccess = true
)
