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

// DexTracerouteTestService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDexTracerouteTestService] method
// instead.
type DexTracerouteTestService struct {
	Options []option.RequestOption
}

// NewDexTracerouteTestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDexTracerouteTestService(opts ...option.RequestOption) (r *DexTracerouteTestService) {
	r = &DexTracerouteTestService{}
	r.Options = opts
	return
}

// Get test details and aggregate performance metrics for an traceroute test for a
// given time period between 1 hour and 7 days.
func (r *DexTracerouteTestService) Get(ctx context.Context, accountID string, testID string, query DexTracerouteTestGetParams, opts ...option.RequestOption) (res *DexTracerouteTestGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DexTracerouteTestGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/traceroute-tests/%s", accountID, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a breakdown of metrics by hop for individual traceroute test runs
func (r *DexTracerouteTestService) NetworkPath(ctx context.Context, accountID string, testID string, query DexTracerouteTestNetworkPathParams, opts ...option.RequestOption) (res *DexTracerouteTestNetworkPathResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DexTracerouteTestNetworkPathResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/traceroute-tests/%s/network-path", accountID, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get percentiles for a traceroute test for a given time period between 1 hour and
// 7 days.
func (r *DexTracerouteTestService) Percentiles(ctx context.Context, accountID string, testID string, query DexTracerouteTestPercentilesParams, opts ...option.RequestOption) (res *DexTracerouteTestPercentilesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DexTracerouteTestPercentilesResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/traceroute-tests/%s/percentiles", accountID, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DexTracerouteTestGetResponse struct {
	// The host of the Traceroute synthetic application test
	Host string `json:"host,required"`
	// The interval at which the Traceroute synthetic application test is set to run.
	Interval string                           `json:"interval,required"`
	Kind     DexTracerouteTestGetResponseKind `json:"kind,required"`
	// The name of the Traceroute synthetic application test
	Name                  string                                              `json:"name,required"`
	TracerouteStats       DexTracerouteTestGetResponseTracerouteStats         `json:"tracerouteStats,nullable"`
	TracerouteStatsByColo []DexTracerouteTestGetResponseTracerouteStatsByColo `json:"tracerouteStatsByColo"`
	JSON                  dexTracerouteTestGetResponseJSON                    `json:"-"`
}

// dexTracerouteTestGetResponseJSON contains the JSON metadata for the struct
// [DexTracerouteTestGetResponse]
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

func (r *DexTracerouteTestGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseKind string

const (
	DexTracerouteTestGetResponseKindTraceroute DexTracerouteTestGetResponseKind = "traceroute"
)

type DexTracerouteTestGetResponseTracerouteStats struct {
	AvailabilityPct DexTracerouteTestGetResponseTracerouteStatsAvailabilityPct `json:"availabilityPct,required"`
	HopsCount       DexTracerouteTestGetResponseTracerouteStatsHopsCount       `json:"hopsCount,required"`
	PacketLossPct   DexTracerouteTestGetResponseTracerouteStatsPacketLossPct   `json:"packetLossPct,required"`
	RoundTripTimeMs DexTracerouteTestGetResponseTracerouteStatsRoundTripTimeMs `json:"roundTripTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                                           `json:"uniqueDevicesTotal,required"`
	JSON               dexTracerouteTestGetResponseTracerouteStatsJSON `json:"-"`
}

// dexTracerouteTestGetResponseTracerouteStatsJSON contains the JSON metadata for
// the struct [DexTracerouteTestGetResponseTracerouteStats]
type dexTracerouteTestGetResponseTracerouteStatsJSON struct {
	AvailabilityPct    apijson.Field
	HopsCount          apijson.Field
	PacketLossPct      apijson.Field
	RoundTripTimeMs    apijson.Field
	UniqueDevicesTotal apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseTracerouteStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseTracerouteStatsAvailabilityPct struct {
	Slots []DexTracerouteTestGetResponseTracerouteStatsAvailabilityPctSlot `json:"slots,required"`
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
// [DexTracerouteTestGetResponseTracerouteStatsAvailabilityPct]
type dexTracerouteTestGetResponseTracerouteStatsAvailabilityPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseTracerouteStatsAvailabilityPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseTracerouteStatsAvailabilityPctSlot struct {
	Timestamp string                                                             `json:"timestamp,required"`
	Value     float64                                                            `json:"value,required"`
	JSON      dexTracerouteTestGetResponseTracerouteStatsAvailabilityPctSlotJSON `json:"-"`
}

// dexTracerouteTestGetResponseTracerouteStatsAvailabilityPctSlotJSON contains the
// JSON metadata for the struct
// [DexTracerouteTestGetResponseTracerouteStatsAvailabilityPctSlot]
type dexTracerouteTestGetResponseTracerouteStatsAvailabilityPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseTracerouteStatsAvailabilityPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseTracerouteStatsHopsCount struct {
	Slots []DexTracerouteTestGetResponseTracerouteStatsHopsCountSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                    `json:"min,nullable"`
	JSON dexTracerouteTestGetResponseTracerouteStatsHopsCountJSON `json:"-"`
}

// dexTracerouteTestGetResponseTracerouteStatsHopsCountJSON contains the JSON
// metadata for the struct [DexTracerouteTestGetResponseTracerouteStatsHopsCount]
type dexTracerouteTestGetResponseTracerouteStatsHopsCountJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseTracerouteStatsHopsCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseTracerouteStatsHopsCountSlot struct {
	Timestamp string                                                       `json:"timestamp,required"`
	Value     int64                                                        `json:"value,required"`
	JSON      dexTracerouteTestGetResponseTracerouteStatsHopsCountSlotJSON `json:"-"`
}

// dexTracerouteTestGetResponseTracerouteStatsHopsCountSlotJSON contains the JSON
// metadata for the struct
// [DexTracerouteTestGetResponseTracerouteStatsHopsCountSlot]
type dexTracerouteTestGetResponseTracerouteStatsHopsCountSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseTracerouteStatsHopsCountSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseTracerouteStatsPacketLossPct struct {
	Slots []DexTracerouteTestGetResponseTracerouteStatsPacketLossPctSlot `json:"slots,required"`
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
// [DexTracerouteTestGetResponseTracerouteStatsPacketLossPct]
type dexTracerouteTestGetResponseTracerouteStatsPacketLossPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseTracerouteStatsPacketLossPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseTracerouteStatsPacketLossPctSlot struct {
	Timestamp string                                                           `json:"timestamp,required"`
	Value     float64                                                          `json:"value,required"`
	JSON      dexTracerouteTestGetResponseTracerouteStatsPacketLossPctSlotJSON `json:"-"`
}

// dexTracerouteTestGetResponseTracerouteStatsPacketLossPctSlotJSON contains the
// JSON metadata for the struct
// [DexTracerouteTestGetResponseTracerouteStatsPacketLossPctSlot]
type dexTracerouteTestGetResponseTracerouteStatsPacketLossPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseTracerouteStatsPacketLossPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseTracerouteStatsRoundTripTimeMs struct {
	Slots []DexTracerouteTestGetResponseTracerouteStatsRoundTripTimeMsSlot `json:"slots,required"`
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
// [DexTracerouteTestGetResponseTracerouteStatsRoundTripTimeMs]
type dexTracerouteTestGetResponseTracerouteStatsRoundTripTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseTracerouteStatsRoundTripTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseTracerouteStatsRoundTripTimeMsSlot struct {
	Timestamp string                                                             `json:"timestamp,required"`
	Value     int64                                                              `json:"value,required"`
	JSON      dexTracerouteTestGetResponseTracerouteStatsRoundTripTimeMsSlotJSON `json:"-"`
}

// dexTracerouteTestGetResponseTracerouteStatsRoundTripTimeMsSlotJSON contains the
// JSON metadata for the struct
// [DexTracerouteTestGetResponseTracerouteStatsRoundTripTimeMsSlot]
type dexTracerouteTestGetResponseTracerouteStatsRoundTripTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseTracerouteStatsRoundTripTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseTracerouteStatsByColo struct {
	AvailabilityPct DexTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPct `json:"availabilityPct,required"`
	Colo            string                                                           `json:"colo,required"`
	HopsCount       DexTracerouteTestGetResponseTracerouteStatsByColoHopsCount       `json:"hopsCount,required"`
	PacketLossPct   DexTracerouteTestGetResponseTracerouteStatsByColoPacketLossPct   `json:"packetLossPct,required"`
	RoundTripTimeMs DexTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMs `json:"roundTripTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                                                 `json:"uniqueDevicesTotal,required"`
	JSON               dexTracerouteTestGetResponseTracerouteStatsByColoJSON `json:"-"`
}

// dexTracerouteTestGetResponseTracerouteStatsByColoJSON contains the JSON metadata
// for the struct [DexTracerouteTestGetResponseTracerouteStatsByColo]
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

func (r *DexTracerouteTestGetResponseTracerouteStatsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPct struct {
	Slots []DexTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPctSlot `json:"slots,required"`
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
// [DexTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPct]
type dexTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPctSlot struct {
	Timestamp string                                                                   `json:"timestamp,required"`
	Value     float64                                                                  `json:"value,required"`
	JSON      dexTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPctSlotJSON `json:"-"`
}

// dexTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPctSlotJSON
// contains the JSON metadata for the struct
// [DexTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPctSlot]
type dexTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseTracerouteStatsByColoHopsCount struct {
	Slots []DexTracerouteTestGetResponseTracerouteStatsByColoHopsCountSlot `json:"slots,required"`
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
// [DexTracerouteTestGetResponseTracerouteStatsByColoHopsCount]
type dexTracerouteTestGetResponseTracerouteStatsByColoHopsCountJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseTracerouteStatsByColoHopsCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseTracerouteStatsByColoHopsCountSlot struct {
	Timestamp string                                                             `json:"timestamp,required"`
	Value     int64                                                              `json:"value,required"`
	JSON      dexTracerouteTestGetResponseTracerouteStatsByColoHopsCountSlotJSON `json:"-"`
}

// dexTracerouteTestGetResponseTracerouteStatsByColoHopsCountSlotJSON contains the
// JSON metadata for the struct
// [DexTracerouteTestGetResponseTracerouteStatsByColoHopsCountSlot]
type dexTracerouteTestGetResponseTracerouteStatsByColoHopsCountSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseTracerouteStatsByColoHopsCountSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseTracerouteStatsByColoPacketLossPct struct {
	Slots []DexTracerouteTestGetResponseTracerouteStatsByColoPacketLossPctSlot `json:"slots,required"`
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
// [DexTracerouteTestGetResponseTracerouteStatsByColoPacketLossPct]
type dexTracerouteTestGetResponseTracerouteStatsByColoPacketLossPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseTracerouteStatsByColoPacketLossPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseTracerouteStatsByColoPacketLossPctSlot struct {
	Timestamp string                                                                 `json:"timestamp,required"`
	Value     float64                                                                `json:"value,required"`
	JSON      dexTracerouteTestGetResponseTracerouteStatsByColoPacketLossPctSlotJSON `json:"-"`
}

// dexTracerouteTestGetResponseTracerouteStatsByColoPacketLossPctSlotJSON contains
// the JSON metadata for the struct
// [DexTracerouteTestGetResponseTracerouteStatsByColoPacketLossPctSlot]
type dexTracerouteTestGetResponseTracerouteStatsByColoPacketLossPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseTracerouteStatsByColoPacketLossPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMs struct {
	Slots []DexTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMsSlot `json:"slots,required"`
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
// [DexTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMs]
type dexTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMsSlot struct {
	Timestamp string                                                                   `json:"timestamp,required"`
	Value     int64                                                                    `json:"value,required"`
	JSON      dexTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMsSlotJSON `json:"-"`
}

// dexTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMsSlotJSON
// contains the JSON metadata for the struct
// [DexTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMsSlot]
type dexTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestNetworkPathResponse struct {
	// API Resource UUID tag.
	ID         string `json:"id,required"`
	DeviceName string `json:"deviceName"`
	// The interval at which the Traceroute synthetic application test is set to run.
	Interval    string                                          `json:"interval"`
	Kind        DexTracerouteTestNetworkPathResponseKind        `json:"kind"`
	Name        string                                          `json:"name"`
	NetworkPath DexTracerouteTestNetworkPathResponseNetworkPath `json:"networkPath,nullable"`
	// The host of the Traceroute synthetic application test
	URL  string                                   `json:"url"`
	JSON dexTracerouteTestNetworkPathResponseJSON `json:"-"`
}

// dexTracerouteTestNetworkPathResponseJSON contains the JSON metadata for the
// struct [DexTracerouteTestNetworkPathResponse]
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

func (r *DexTracerouteTestNetworkPathResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestNetworkPathResponseKind string

const (
	DexTracerouteTestNetworkPathResponseKindTraceroute DexTracerouteTestNetworkPathResponseKind = "traceroute"
)

type DexTracerouteTestNetworkPathResponseNetworkPath struct {
	Slots []DexTracerouteTestNetworkPathResponseNetworkPathSlot `json:"slots,required"`
	// Specifies the sampling applied, if any, to the slots response. When sampled,
	// results shown represent the first test run to the start of each sampling
	// interval.
	Sampling DexTracerouteTestNetworkPathResponseNetworkPathSampling `json:"sampling,nullable"`
	JSON     dexTracerouteTestNetworkPathResponseNetworkPathJSON     `json:"-"`
}

// dexTracerouteTestNetworkPathResponseNetworkPathJSON contains the JSON metadata
// for the struct [DexTracerouteTestNetworkPathResponseNetworkPath]
type dexTracerouteTestNetworkPathResponseNetworkPathJSON struct {
	Slots       apijson.Field
	Sampling    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestNetworkPathResponseNetworkPath) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestNetworkPathResponseNetworkPathSlot struct {
	// API Resource UUID tag.
	ID string `json:"id,required"`
	// Round trip time in ms of the client to app mile
	ClientToAppRttMs int64 `json:"clientToAppRttMs,required,nullable"`
	// Round trip time in ms of the client to Cloudflare egress mile
	ClientToCfEgressRttMs int64 `json:"clientToCfEgressRttMs,required,nullable"`
	// Round trip time in ms of the client to Cloudflare ingress mile
	ClientToCfIngressRttMs int64  `json:"clientToCfIngressRttMs,required,nullable"`
	Timestamp              string `json:"timestamp,required"`
	// Round trip time in ms of the client to ISP mile
	ClientToIspRttMs int64                                                   `json:"clientToIspRttMs,nullable"`
	JSON             dexTracerouteTestNetworkPathResponseNetworkPathSlotJSON `json:"-"`
}

// dexTracerouteTestNetworkPathResponseNetworkPathSlotJSON contains the JSON
// metadata for the struct [DexTracerouteTestNetworkPathResponseNetworkPathSlot]
type dexTracerouteTestNetworkPathResponseNetworkPathSlotJSON struct {
	ID                     apijson.Field
	ClientToAppRttMs       apijson.Field
	ClientToCfEgressRttMs  apijson.Field
	ClientToCfIngressRttMs apijson.Field
	Timestamp              apijson.Field
	ClientToIspRttMs       apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *DexTracerouteTestNetworkPathResponseNetworkPathSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the sampling applied, if any, to the slots response. When sampled,
// results shown represent the first test run to the start of each sampling
// interval.
type DexTracerouteTestNetworkPathResponseNetworkPathSampling struct {
	Unit  DexTracerouteTestNetworkPathResponseNetworkPathSamplingUnit `json:"unit,required"`
	Value int64                                                       `json:"value,required"`
	JSON  dexTracerouteTestNetworkPathResponseNetworkPathSamplingJSON `json:"-"`
}

// dexTracerouteTestNetworkPathResponseNetworkPathSamplingJSON contains the JSON
// metadata for the struct
// [DexTracerouteTestNetworkPathResponseNetworkPathSampling]
type dexTracerouteTestNetworkPathResponseNetworkPathSamplingJSON struct {
	Unit        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestNetworkPathResponseNetworkPathSampling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestNetworkPathResponseNetworkPathSamplingUnit string

const (
	DexTracerouteTestNetworkPathResponseNetworkPathSamplingUnitHours DexTracerouteTestNetworkPathResponseNetworkPathSamplingUnit = "hours"
)

type DexTracerouteTestPercentilesResponse struct {
	HopsCount       DexTracerouteTestPercentilesResponseHopsCount       `json:"hopsCount"`
	PacketLossPct   DexTracerouteTestPercentilesResponsePacketLossPct   `json:"packetLossPct"`
	RoundTripTimeMs DexTracerouteTestPercentilesResponseRoundTripTimeMs `json:"roundTripTimeMs"`
	JSON            dexTracerouteTestPercentilesResponseJSON            `json:"-"`
}

// dexTracerouteTestPercentilesResponseJSON contains the JSON metadata for the
// struct [DexTracerouteTestPercentilesResponse]
type dexTracerouteTestPercentilesResponseJSON struct {
	HopsCount       apijson.Field
	PacketLossPct   apijson.Field
	RoundTripTimeMs apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DexTracerouteTestPercentilesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestPercentilesResponseHopsCount struct {
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
// the struct [DexTracerouteTestPercentilesResponseHopsCount]
type dexTracerouteTestPercentilesResponseHopsCountJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestPercentilesResponseHopsCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestPercentilesResponsePacketLossPct struct {
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
// for the struct [DexTracerouteTestPercentilesResponsePacketLossPct]
type dexTracerouteTestPercentilesResponsePacketLossPctJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestPercentilesResponsePacketLossPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestPercentilesResponseRoundTripTimeMs struct {
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
// metadata for the struct [DexTracerouteTestPercentilesResponseRoundTripTimeMs]
type dexTracerouteTestPercentilesResponseRoundTripTimeMsJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestPercentilesResponseRoundTripTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetParams struct {
	// Time interval for aggregate time slots.
	Interval param.Field[DexTracerouteTestGetParamsInterval] `query:"interval,required"`
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

// URLQuery serializes [DexTracerouteTestGetParams]'s query parameters as
// `url.Values`.
func (r DexTracerouteTestGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Time interval for aggregate time slots.
type DexTracerouteTestGetParamsInterval string

const (
	DexTracerouteTestGetParamsIntervalMinute DexTracerouteTestGetParamsInterval = "minute"
	DexTracerouteTestGetParamsIntervalHour   DexTracerouteTestGetParamsInterval = "hour"
)

type DexTracerouteTestGetResponseEnvelope struct {
	Errors   []DexTracerouteTestGetResponseEnvelopeErrors   `json:"errors"`
	Messages []DexTracerouteTestGetResponseEnvelopeMessages `json:"messages"`
	Result   DexTracerouteTestGetResponse                   `json:"result"`
	// Whether the API call was successful
	Success DexTracerouteTestGetResponseEnvelopeSuccess `json:"success"`
	JSON    dexTracerouteTestGetResponseEnvelopeJSON    `json:"-"`
}

// dexTracerouteTestGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DexTracerouteTestGetResponseEnvelope]
type dexTracerouteTestGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    dexTracerouteTestGetResponseEnvelopeErrorsJSON `json:"-"`
}

// dexTracerouteTestGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DexTracerouteTestGetResponseEnvelopeErrors]
type dexTracerouteTestGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    dexTracerouteTestGetResponseEnvelopeMessagesJSON `json:"-"`
}

// dexTracerouteTestGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DexTracerouteTestGetResponseEnvelopeMessages]
type dexTracerouteTestGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DexTracerouteTestGetResponseEnvelopeSuccess bool

const (
	DexTracerouteTestGetResponseEnvelopeSuccessTrue DexTracerouteTestGetResponseEnvelopeSuccess = true
)

type DexTracerouteTestNetworkPathParams struct {
	// Device to filter tracroute result runs to
	DeviceID param.Field[string] `query:"deviceId,required"`
	// Time interval for aggregate time slots.
	Interval param.Field[DexTracerouteTestNetworkPathParamsInterval] `query:"interval,required"`
	// End time for aggregate metrics in ISO ms
	TimeEnd param.Field[string] `query:"timeEnd,required"`
	// Start time for aggregate metrics in ISO ms
	TimeStart param.Field[string] `query:"timeStart,required"`
}

// URLQuery serializes [DexTracerouteTestNetworkPathParams]'s query parameters as
// `url.Values`.
func (r DexTracerouteTestNetworkPathParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Time interval for aggregate time slots.
type DexTracerouteTestNetworkPathParamsInterval string

const (
	DexTracerouteTestNetworkPathParamsIntervalMinute DexTracerouteTestNetworkPathParamsInterval = "minute"
	DexTracerouteTestNetworkPathParamsIntervalHour   DexTracerouteTestNetworkPathParamsInterval = "hour"
)

type DexTracerouteTestNetworkPathResponseEnvelope struct {
	Errors   []DexTracerouteTestNetworkPathResponseEnvelopeErrors   `json:"errors"`
	Messages []DexTracerouteTestNetworkPathResponseEnvelopeMessages `json:"messages"`
	Result   DexTracerouteTestNetworkPathResponse                   `json:"result"`
	// Whether the API call was successful
	Success DexTracerouteTestNetworkPathResponseEnvelopeSuccess `json:"success"`
	JSON    dexTracerouteTestNetworkPathResponseEnvelopeJSON    `json:"-"`
}

// dexTracerouteTestNetworkPathResponseEnvelopeJSON contains the JSON metadata for
// the struct [DexTracerouteTestNetworkPathResponseEnvelope]
type dexTracerouteTestNetworkPathResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestNetworkPathResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestNetworkPathResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    dexTracerouteTestNetworkPathResponseEnvelopeErrorsJSON `json:"-"`
}

// dexTracerouteTestNetworkPathResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DexTracerouteTestNetworkPathResponseEnvelopeErrors]
type dexTracerouteTestNetworkPathResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestNetworkPathResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestNetworkPathResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    dexTracerouteTestNetworkPathResponseEnvelopeMessagesJSON `json:"-"`
}

// dexTracerouteTestNetworkPathResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DexTracerouteTestNetworkPathResponseEnvelopeMessages]
type dexTracerouteTestNetworkPathResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestNetworkPathResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DexTracerouteTestNetworkPathResponseEnvelopeSuccess bool

const (
	DexTracerouteTestNetworkPathResponseEnvelopeSuccessTrue DexTracerouteTestNetworkPathResponseEnvelopeSuccess = true
)

type DexTracerouteTestPercentilesParams struct {
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

// URLQuery serializes [DexTracerouteTestPercentilesParams]'s query parameters as
// `url.Values`.
func (r DexTracerouteTestPercentilesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DexTracerouteTestPercentilesResponseEnvelope struct {
	Errors   []DexTracerouteTestPercentilesResponseEnvelopeErrors   `json:"errors"`
	Messages []DexTracerouteTestPercentilesResponseEnvelopeMessages `json:"messages"`
	Result   DexTracerouteTestPercentilesResponse                   `json:"result"`
	// Whether the API call was successful
	Success DexTracerouteTestPercentilesResponseEnvelopeSuccess `json:"success"`
	JSON    dexTracerouteTestPercentilesResponseEnvelopeJSON    `json:"-"`
}

// dexTracerouteTestPercentilesResponseEnvelopeJSON contains the JSON metadata for
// the struct [DexTracerouteTestPercentilesResponseEnvelope]
type dexTracerouteTestPercentilesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestPercentilesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestPercentilesResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    dexTracerouteTestPercentilesResponseEnvelopeErrorsJSON `json:"-"`
}

// dexTracerouteTestPercentilesResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DexTracerouteTestPercentilesResponseEnvelopeErrors]
type dexTracerouteTestPercentilesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestPercentilesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestPercentilesResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    dexTracerouteTestPercentilesResponseEnvelopeMessagesJSON `json:"-"`
}

// dexTracerouteTestPercentilesResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DexTracerouteTestPercentilesResponseEnvelopeMessages]
type dexTracerouteTestPercentilesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestPercentilesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DexTracerouteTestPercentilesResponseEnvelopeSuccess bool

const (
	DexTracerouteTestPercentilesResponseEnvelopeSuccessTrue DexTracerouteTestPercentilesResponseEnvelopeSuccess = true
)
