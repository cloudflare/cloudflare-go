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
	path := fmt.Sprintf("accounts/%s/dex/traceroute-tests/%s", accountID, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get a breakdown of metrics by hop for individual traceroute test runs
func (r *DexTracerouteTestService) NetworkPath(ctx context.Context, accountID string, testID string, query DexTracerouteTestNetworkPathParams, opts ...option.RequestOption) (res *DexTracerouteTestNetworkPathResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dex/traceroute-tests/%s/network-path", accountID, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get percentiles for a traceroute test for a given time period between 1 hour and
// 7 days.
func (r *DexTracerouteTestService) Percentiles(ctx context.Context, accountID string, testID string, query DexTracerouteTestPercentilesParams, opts ...option.RequestOption) (res *DexTracerouteTestPercentilesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dex/traceroute-tests/%s/percentiles", accountID, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DexTracerouteTestGetResponse struct {
	Errors   []DexTracerouteTestGetResponseError   `json:"errors"`
	Messages []DexTracerouteTestGetResponseMessage `json:"messages"`
	Result   DexTracerouteTestGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success DexTracerouteTestGetResponseSuccess `json:"success"`
	JSON    dexTracerouteTestGetResponseJSON    `json:"-"`
}

// dexTracerouteTestGetResponseJSON contains the JSON metadata for the struct
// [DexTracerouteTestGetResponse]
type dexTracerouteTestGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    dexTracerouteTestGetResponseErrorJSON `json:"-"`
}

// dexTracerouteTestGetResponseErrorJSON contains the JSON metadata for the struct
// [DexTracerouteTestGetResponseError]
type dexTracerouteTestGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    dexTracerouteTestGetResponseMessageJSON `json:"-"`
}

// dexTracerouteTestGetResponseMessageJSON contains the JSON metadata for the
// struct [DexTracerouteTestGetResponseMessage]
type dexTracerouteTestGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseResult struct {
	// The host of the Traceroute synthetic application test
	Host string `json:"host,required"`
	// The interval at which the Traceroute synthetic application test is set to run.
	Interval string                                 `json:"interval,required"`
	Kind     DexTracerouteTestGetResponseResultKind `json:"kind,required"`
	// The name of the Traceroute synthetic application test
	Name                  string                                                    `json:"name,required"`
	TracerouteStats       DexTracerouteTestGetResponseResultTracerouteStats         `json:"tracerouteStats,nullable"`
	TracerouteStatsByColo []DexTracerouteTestGetResponseResultTracerouteStatsByColo `json:"tracerouteStatsByColo"`
	JSON                  dexTracerouteTestGetResponseResultJSON                    `json:"-"`
}

// dexTracerouteTestGetResponseResultJSON contains the JSON metadata for the struct
// [DexTracerouteTestGetResponseResult]
type dexTracerouteTestGetResponseResultJSON struct {
	Host                  apijson.Field
	Interval              apijson.Field
	Kind                  apijson.Field
	Name                  apijson.Field
	TracerouteStats       apijson.Field
	TracerouteStatsByColo apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseResultKind string

const (
	DexTracerouteTestGetResponseResultKindTraceroute DexTracerouteTestGetResponseResultKind = "traceroute"
)

type DexTracerouteTestGetResponseResultTracerouteStats struct {
	AvailabilityPct DexTracerouteTestGetResponseResultTracerouteStatsAvailabilityPct `json:"availabilityPct,required"`
	HopsCount       DexTracerouteTestGetResponseResultTracerouteStatsHopsCount       `json:"hopsCount,required"`
	PacketLossPct   DexTracerouteTestGetResponseResultTracerouteStatsPacketLossPct   `json:"packetLossPct,required"`
	RoundTripTimeMs DexTracerouteTestGetResponseResultTracerouteStatsRoundTripTimeMs `json:"roundTripTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                                                 `json:"uniqueDevicesTotal,required"`
	JSON               dexTracerouteTestGetResponseResultTracerouteStatsJSON `json:"-"`
}

// dexTracerouteTestGetResponseResultTracerouteStatsJSON contains the JSON metadata
// for the struct [DexTracerouteTestGetResponseResultTracerouteStats]
type dexTracerouteTestGetResponseResultTracerouteStatsJSON struct {
	AvailabilityPct    apijson.Field
	HopsCount          apijson.Field
	PacketLossPct      apijson.Field
	RoundTripTimeMs    apijson.Field
	UniqueDevicesTotal apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseResultTracerouteStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseResultTracerouteStatsAvailabilityPct struct {
	Slots []DexTracerouteTestGetResponseResultTracerouteStatsAvailabilityPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                                              `json:"min,nullable"`
	JSON dexTracerouteTestGetResponseResultTracerouteStatsAvailabilityPctJSON `json:"-"`
}

// dexTracerouteTestGetResponseResultTracerouteStatsAvailabilityPctJSON contains
// the JSON metadata for the struct
// [DexTracerouteTestGetResponseResultTracerouteStatsAvailabilityPct]
type dexTracerouteTestGetResponseResultTracerouteStatsAvailabilityPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseResultTracerouteStatsAvailabilityPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseResultTracerouteStatsAvailabilityPctSlot struct {
	Timestamp string                                                                   `json:"timestamp,required"`
	Value     float64                                                                  `json:"value,required"`
	JSON      dexTracerouteTestGetResponseResultTracerouteStatsAvailabilityPctSlotJSON `json:"-"`
}

// dexTracerouteTestGetResponseResultTracerouteStatsAvailabilityPctSlotJSON
// contains the JSON metadata for the struct
// [DexTracerouteTestGetResponseResultTracerouteStatsAvailabilityPctSlot]
type dexTracerouteTestGetResponseResultTracerouteStatsAvailabilityPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseResultTracerouteStatsAvailabilityPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseResultTracerouteStatsHopsCount struct {
	Slots []DexTracerouteTestGetResponseResultTracerouteStatsHopsCountSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                          `json:"min,nullable"`
	JSON dexTracerouteTestGetResponseResultTracerouteStatsHopsCountJSON `json:"-"`
}

// dexTracerouteTestGetResponseResultTracerouteStatsHopsCountJSON contains the JSON
// metadata for the struct
// [DexTracerouteTestGetResponseResultTracerouteStatsHopsCount]
type dexTracerouteTestGetResponseResultTracerouteStatsHopsCountJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseResultTracerouteStatsHopsCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseResultTracerouteStatsHopsCountSlot struct {
	Timestamp string                                                             `json:"timestamp,required"`
	Value     int64                                                              `json:"value,required"`
	JSON      dexTracerouteTestGetResponseResultTracerouteStatsHopsCountSlotJSON `json:"-"`
}

// dexTracerouteTestGetResponseResultTracerouteStatsHopsCountSlotJSON contains the
// JSON metadata for the struct
// [DexTracerouteTestGetResponseResultTracerouteStatsHopsCountSlot]
type dexTracerouteTestGetResponseResultTracerouteStatsHopsCountSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseResultTracerouteStatsHopsCountSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseResultTracerouteStatsPacketLossPct struct {
	Slots []DexTracerouteTestGetResponseResultTracerouteStatsPacketLossPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                                            `json:"min,nullable"`
	JSON dexTracerouteTestGetResponseResultTracerouteStatsPacketLossPctJSON `json:"-"`
}

// dexTracerouteTestGetResponseResultTracerouteStatsPacketLossPctJSON contains the
// JSON metadata for the struct
// [DexTracerouteTestGetResponseResultTracerouteStatsPacketLossPct]
type dexTracerouteTestGetResponseResultTracerouteStatsPacketLossPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseResultTracerouteStatsPacketLossPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseResultTracerouteStatsPacketLossPctSlot struct {
	Timestamp string                                                                 `json:"timestamp,required"`
	Value     float64                                                                `json:"value,required"`
	JSON      dexTracerouteTestGetResponseResultTracerouteStatsPacketLossPctSlotJSON `json:"-"`
}

// dexTracerouteTestGetResponseResultTracerouteStatsPacketLossPctSlotJSON contains
// the JSON metadata for the struct
// [DexTracerouteTestGetResponseResultTracerouteStatsPacketLossPctSlot]
type dexTracerouteTestGetResponseResultTracerouteStatsPacketLossPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseResultTracerouteStatsPacketLossPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseResultTracerouteStatsRoundTripTimeMs struct {
	Slots []DexTracerouteTestGetResponseResultTracerouteStatsRoundTripTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                                `json:"min,nullable"`
	JSON dexTracerouteTestGetResponseResultTracerouteStatsRoundTripTimeMsJSON `json:"-"`
}

// dexTracerouteTestGetResponseResultTracerouteStatsRoundTripTimeMsJSON contains
// the JSON metadata for the struct
// [DexTracerouteTestGetResponseResultTracerouteStatsRoundTripTimeMs]
type dexTracerouteTestGetResponseResultTracerouteStatsRoundTripTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseResultTracerouteStatsRoundTripTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseResultTracerouteStatsRoundTripTimeMsSlot struct {
	Timestamp string                                                                   `json:"timestamp,required"`
	Value     int64                                                                    `json:"value,required"`
	JSON      dexTracerouteTestGetResponseResultTracerouteStatsRoundTripTimeMsSlotJSON `json:"-"`
}

// dexTracerouteTestGetResponseResultTracerouteStatsRoundTripTimeMsSlotJSON
// contains the JSON metadata for the struct
// [DexTracerouteTestGetResponseResultTracerouteStatsRoundTripTimeMsSlot]
type dexTracerouteTestGetResponseResultTracerouteStatsRoundTripTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseResultTracerouteStatsRoundTripTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseResultTracerouteStatsByColo struct {
	AvailabilityPct DexTracerouteTestGetResponseResultTracerouteStatsByColoAvailabilityPct `json:"availabilityPct,required"`
	Colo            string                                                                 `json:"colo,required"`
	HopsCount       DexTracerouteTestGetResponseResultTracerouteStatsByColoHopsCount       `json:"hopsCount,required"`
	PacketLossPct   DexTracerouteTestGetResponseResultTracerouteStatsByColoPacketLossPct   `json:"packetLossPct,required"`
	RoundTripTimeMs DexTracerouteTestGetResponseResultTracerouteStatsByColoRoundTripTimeMs `json:"roundTripTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                                                       `json:"uniqueDevicesTotal,required"`
	JSON               dexTracerouteTestGetResponseResultTracerouteStatsByColoJSON `json:"-"`
}

// dexTracerouteTestGetResponseResultTracerouteStatsByColoJSON contains the JSON
// metadata for the struct
// [DexTracerouteTestGetResponseResultTracerouteStatsByColo]
type dexTracerouteTestGetResponseResultTracerouteStatsByColoJSON struct {
	AvailabilityPct    apijson.Field
	Colo               apijson.Field
	HopsCount          apijson.Field
	PacketLossPct      apijson.Field
	RoundTripTimeMs    apijson.Field
	UniqueDevicesTotal apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseResultTracerouteStatsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseResultTracerouteStatsByColoAvailabilityPct struct {
	Slots []DexTracerouteTestGetResponseResultTracerouteStatsByColoAvailabilityPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                                                    `json:"min,nullable"`
	JSON dexTracerouteTestGetResponseResultTracerouteStatsByColoAvailabilityPctJSON `json:"-"`
}

// dexTracerouteTestGetResponseResultTracerouteStatsByColoAvailabilityPctJSON
// contains the JSON metadata for the struct
// [DexTracerouteTestGetResponseResultTracerouteStatsByColoAvailabilityPct]
type dexTracerouteTestGetResponseResultTracerouteStatsByColoAvailabilityPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseResultTracerouteStatsByColoAvailabilityPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseResultTracerouteStatsByColoAvailabilityPctSlot struct {
	Timestamp string                                                                         `json:"timestamp,required"`
	Value     float64                                                                        `json:"value,required"`
	JSON      dexTracerouteTestGetResponseResultTracerouteStatsByColoAvailabilityPctSlotJSON `json:"-"`
}

// dexTracerouteTestGetResponseResultTracerouteStatsByColoAvailabilityPctSlotJSON
// contains the JSON metadata for the struct
// [DexTracerouteTestGetResponseResultTracerouteStatsByColoAvailabilityPctSlot]
type dexTracerouteTestGetResponseResultTracerouteStatsByColoAvailabilityPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseResultTracerouteStatsByColoAvailabilityPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseResultTracerouteStatsByColoHopsCount struct {
	Slots []DexTracerouteTestGetResponseResultTracerouteStatsByColoHopsCountSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                                `json:"min,nullable"`
	JSON dexTracerouteTestGetResponseResultTracerouteStatsByColoHopsCountJSON `json:"-"`
}

// dexTracerouteTestGetResponseResultTracerouteStatsByColoHopsCountJSON contains
// the JSON metadata for the struct
// [DexTracerouteTestGetResponseResultTracerouteStatsByColoHopsCount]
type dexTracerouteTestGetResponseResultTracerouteStatsByColoHopsCountJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseResultTracerouteStatsByColoHopsCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseResultTracerouteStatsByColoHopsCountSlot struct {
	Timestamp string                                                                   `json:"timestamp,required"`
	Value     int64                                                                    `json:"value,required"`
	JSON      dexTracerouteTestGetResponseResultTracerouteStatsByColoHopsCountSlotJSON `json:"-"`
}

// dexTracerouteTestGetResponseResultTracerouteStatsByColoHopsCountSlotJSON
// contains the JSON metadata for the struct
// [DexTracerouteTestGetResponseResultTracerouteStatsByColoHopsCountSlot]
type dexTracerouteTestGetResponseResultTracerouteStatsByColoHopsCountSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseResultTracerouteStatsByColoHopsCountSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseResultTracerouteStatsByColoPacketLossPct struct {
	Slots []DexTracerouteTestGetResponseResultTracerouteStatsByColoPacketLossPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                                                  `json:"min,nullable"`
	JSON dexTracerouteTestGetResponseResultTracerouteStatsByColoPacketLossPctJSON `json:"-"`
}

// dexTracerouteTestGetResponseResultTracerouteStatsByColoPacketLossPctJSON
// contains the JSON metadata for the struct
// [DexTracerouteTestGetResponseResultTracerouteStatsByColoPacketLossPct]
type dexTracerouteTestGetResponseResultTracerouteStatsByColoPacketLossPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseResultTracerouteStatsByColoPacketLossPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseResultTracerouteStatsByColoPacketLossPctSlot struct {
	Timestamp string                                                                       `json:"timestamp,required"`
	Value     float64                                                                      `json:"value,required"`
	JSON      dexTracerouteTestGetResponseResultTracerouteStatsByColoPacketLossPctSlotJSON `json:"-"`
}

// dexTracerouteTestGetResponseResultTracerouteStatsByColoPacketLossPctSlotJSON
// contains the JSON metadata for the struct
// [DexTracerouteTestGetResponseResultTracerouteStatsByColoPacketLossPctSlot]
type dexTracerouteTestGetResponseResultTracerouteStatsByColoPacketLossPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseResultTracerouteStatsByColoPacketLossPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseResultTracerouteStatsByColoRoundTripTimeMs struct {
	Slots []DexTracerouteTestGetResponseResultTracerouteStatsByColoRoundTripTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                                      `json:"min,nullable"`
	JSON dexTracerouteTestGetResponseResultTracerouteStatsByColoRoundTripTimeMsJSON `json:"-"`
}

// dexTracerouteTestGetResponseResultTracerouteStatsByColoRoundTripTimeMsJSON
// contains the JSON metadata for the struct
// [DexTracerouteTestGetResponseResultTracerouteStatsByColoRoundTripTimeMs]
type dexTracerouteTestGetResponseResultTracerouteStatsByColoRoundTripTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseResultTracerouteStatsByColoRoundTripTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestGetResponseResultTracerouteStatsByColoRoundTripTimeMsSlot struct {
	Timestamp string                                                                         `json:"timestamp,required"`
	Value     int64                                                                          `json:"value,required"`
	JSON      dexTracerouteTestGetResponseResultTracerouteStatsByColoRoundTripTimeMsSlotJSON `json:"-"`
}

// dexTracerouteTestGetResponseResultTracerouteStatsByColoRoundTripTimeMsSlotJSON
// contains the JSON metadata for the struct
// [DexTracerouteTestGetResponseResultTracerouteStatsByColoRoundTripTimeMsSlot]
type dexTracerouteTestGetResponseResultTracerouteStatsByColoRoundTripTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestGetResponseResultTracerouteStatsByColoRoundTripTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DexTracerouteTestGetResponseSuccess bool

const (
	DexTracerouteTestGetResponseSuccessTrue DexTracerouteTestGetResponseSuccess = true
)

type DexTracerouteTestNetworkPathResponse struct {
	Errors   []DexTracerouteTestNetworkPathResponseError   `json:"errors"`
	Messages []DexTracerouteTestNetworkPathResponseMessage `json:"messages"`
	Result   DexTracerouteTestNetworkPathResponseResult    `json:"result"`
	// Whether the API call was successful
	Success DexTracerouteTestNetworkPathResponseSuccess `json:"success"`
	JSON    dexTracerouteTestNetworkPathResponseJSON    `json:"-"`
}

// dexTracerouteTestNetworkPathResponseJSON contains the JSON metadata for the
// struct [DexTracerouteTestNetworkPathResponse]
type dexTracerouteTestNetworkPathResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestNetworkPathResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestNetworkPathResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    dexTracerouteTestNetworkPathResponseErrorJSON `json:"-"`
}

// dexTracerouteTestNetworkPathResponseErrorJSON contains the JSON metadata for the
// struct [DexTracerouteTestNetworkPathResponseError]
type dexTracerouteTestNetworkPathResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestNetworkPathResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestNetworkPathResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    dexTracerouteTestNetworkPathResponseMessageJSON `json:"-"`
}

// dexTracerouteTestNetworkPathResponseMessageJSON contains the JSON metadata for
// the struct [DexTracerouteTestNetworkPathResponseMessage]
type dexTracerouteTestNetworkPathResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestNetworkPathResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestNetworkPathResponseResult struct {
	// API Resource UUID tag.
	ID         string `json:"id,required"`
	DeviceName string `json:"deviceName"`
	// The interval at which the Traceroute synthetic application test is set to run.
	Interval    string                                                `json:"interval"`
	Kind        DexTracerouteTestNetworkPathResponseResultKind        `json:"kind"`
	Name        string                                                `json:"name"`
	NetworkPath DexTracerouteTestNetworkPathResponseResultNetworkPath `json:"networkPath,nullable"`
	// The host of the Traceroute synthetic application test
	URL  string                                         `json:"url"`
	JSON dexTracerouteTestNetworkPathResponseResultJSON `json:"-"`
}

// dexTracerouteTestNetworkPathResponseResultJSON contains the JSON metadata for
// the struct [DexTracerouteTestNetworkPathResponseResult]
type dexTracerouteTestNetworkPathResponseResultJSON struct {
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

func (r *DexTracerouteTestNetworkPathResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestNetworkPathResponseResultKind string

const (
	DexTracerouteTestNetworkPathResponseResultKindTraceroute DexTracerouteTestNetworkPathResponseResultKind = "traceroute"
)

type DexTracerouteTestNetworkPathResponseResultNetworkPath struct {
	Slots []DexTracerouteTestNetworkPathResponseResultNetworkPathSlot `json:"slots,required"`
	// Specifies the sampling applied, if any, to the slots response. When sampled,
	// results shown represent the first test run to the start of each sampling
	// interval.
	Sampling DexTracerouteTestNetworkPathResponseResultNetworkPathSampling `json:"sampling,nullable"`
	JSON     dexTracerouteTestNetworkPathResponseResultNetworkPathJSON     `json:"-"`
}

// dexTracerouteTestNetworkPathResponseResultNetworkPathJSON contains the JSON
// metadata for the struct [DexTracerouteTestNetworkPathResponseResultNetworkPath]
type dexTracerouteTestNetworkPathResponseResultNetworkPathJSON struct {
	Slots       apijson.Field
	Sampling    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestNetworkPathResponseResultNetworkPath) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestNetworkPathResponseResultNetworkPathSlot struct {
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
	ClientToIspRttMs int64                                                         `json:"clientToIspRttMs,nullable"`
	JSON             dexTracerouteTestNetworkPathResponseResultNetworkPathSlotJSON `json:"-"`
}

// dexTracerouteTestNetworkPathResponseResultNetworkPathSlotJSON contains the JSON
// metadata for the struct
// [DexTracerouteTestNetworkPathResponseResultNetworkPathSlot]
type dexTracerouteTestNetworkPathResponseResultNetworkPathSlotJSON struct {
	ID                     apijson.Field
	ClientToAppRttMs       apijson.Field
	ClientToCfEgressRttMs  apijson.Field
	ClientToCfIngressRttMs apijson.Field
	Timestamp              apijson.Field
	ClientToIspRttMs       apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *DexTracerouteTestNetworkPathResponseResultNetworkPathSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the sampling applied, if any, to the slots response. When sampled,
// results shown represent the first test run to the start of each sampling
// interval.
type DexTracerouteTestNetworkPathResponseResultNetworkPathSampling struct {
	Unit  DexTracerouteTestNetworkPathResponseResultNetworkPathSamplingUnit `json:"unit,required"`
	Value int64                                                             `json:"value,required"`
	JSON  dexTracerouteTestNetworkPathResponseResultNetworkPathSamplingJSON `json:"-"`
}

// dexTracerouteTestNetworkPathResponseResultNetworkPathSamplingJSON contains the
// JSON metadata for the struct
// [DexTracerouteTestNetworkPathResponseResultNetworkPathSampling]
type dexTracerouteTestNetworkPathResponseResultNetworkPathSamplingJSON struct {
	Unit        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestNetworkPathResponseResultNetworkPathSampling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestNetworkPathResponseResultNetworkPathSamplingUnit string

const (
	DexTracerouteTestNetworkPathResponseResultNetworkPathSamplingUnitHours DexTracerouteTestNetworkPathResponseResultNetworkPathSamplingUnit = "hours"
)

// Whether the API call was successful
type DexTracerouteTestNetworkPathResponseSuccess bool

const (
	DexTracerouteTestNetworkPathResponseSuccessTrue DexTracerouteTestNetworkPathResponseSuccess = true
)

type DexTracerouteTestPercentilesResponse struct {
	Errors   []DexTracerouteTestPercentilesResponseError   `json:"errors"`
	Messages []DexTracerouteTestPercentilesResponseMessage `json:"messages"`
	Result   DexTracerouteTestPercentilesResponseResult    `json:"result"`
	// Whether the API call was successful
	Success DexTracerouteTestPercentilesResponseSuccess `json:"success"`
	JSON    dexTracerouteTestPercentilesResponseJSON    `json:"-"`
}

// dexTracerouteTestPercentilesResponseJSON contains the JSON metadata for the
// struct [DexTracerouteTestPercentilesResponse]
type dexTracerouteTestPercentilesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestPercentilesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestPercentilesResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    dexTracerouteTestPercentilesResponseErrorJSON `json:"-"`
}

// dexTracerouteTestPercentilesResponseErrorJSON contains the JSON metadata for the
// struct [DexTracerouteTestPercentilesResponseError]
type dexTracerouteTestPercentilesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestPercentilesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestPercentilesResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    dexTracerouteTestPercentilesResponseMessageJSON `json:"-"`
}

// dexTracerouteTestPercentilesResponseMessageJSON contains the JSON metadata for
// the struct [DexTracerouteTestPercentilesResponseMessage]
type dexTracerouteTestPercentilesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestPercentilesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestPercentilesResponseResult struct {
	HopsCount       DexTracerouteTestPercentilesResponseResultHopsCount       `json:"hopsCount"`
	PacketLossPct   DexTracerouteTestPercentilesResponseResultPacketLossPct   `json:"packetLossPct"`
	RoundTripTimeMs DexTracerouteTestPercentilesResponseResultRoundTripTimeMs `json:"roundTripTimeMs"`
	JSON            dexTracerouteTestPercentilesResponseResultJSON            `json:"-"`
}

// dexTracerouteTestPercentilesResponseResultJSON contains the JSON metadata for
// the struct [DexTracerouteTestPercentilesResponseResult]
type dexTracerouteTestPercentilesResponseResultJSON struct {
	HopsCount       apijson.Field
	PacketLossPct   apijson.Field
	RoundTripTimeMs apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DexTracerouteTestPercentilesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestPercentilesResponseResultHopsCount struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                                 `json:"p99,nullable"`
	JSON dexTracerouteTestPercentilesResponseResultHopsCountJSON `json:"-"`
}

// dexTracerouteTestPercentilesResponseResultHopsCountJSON contains the JSON
// metadata for the struct [DexTracerouteTestPercentilesResponseResultHopsCount]
type dexTracerouteTestPercentilesResponseResultHopsCountJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestPercentilesResponseResultHopsCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestPercentilesResponseResultPacketLossPct struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                                     `json:"p99,nullable"`
	JSON dexTracerouteTestPercentilesResponseResultPacketLossPctJSON `json:"-"`
}

// dexTracerouteTestPercentilesResponseResultPacketLossPctJSON contains the JSON
// metadata for the struct
// [DexTracerouteTestPercentilesResponseResultPacketLossPct]
type dexTracerouteTestPercentilesResponseResultPacketLossPctJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestPercentilesResponseResultPacketLossPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestPercentilesResponseResultRoundTripTimeMs struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                                       `json:"p99,nullable"`
	JSON dexTracerouteTestPercentilesResponseResultRoundTripTimeMsJSON `json:"-"`
}

// dexTracerouteTestPercentilesResponseResultRoundTripTimeMsJSON contains the JSON
// metadata for the struct
// [DexTracerouteTestPercentilesResponseResultRoundTripTimeMs]
type dexTracerouteTestPercentilesResponseResultRoundTripTimeMsJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestPercentilesResponseResultRoundTripTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DexTracerouteTestPercentilesResponseSuccess bool

const (
	DexTracerouteTestPercentilesResponseSuccessTrue DexTracerouteTestPercentilesResponseSuccess = true
)

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
