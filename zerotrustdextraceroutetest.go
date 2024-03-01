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
func (r *ZeroTrustDEXTracerouteTestService) Get(ctx context.Context, testID string, params ZeroTrustDEXTracerouteTestGetParams, opts ...option.RequestOption) (res *ZeroTrustDEXTracerouteTestGetResponse, err error) {
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
func (r *ZeroTrustDEXTracerouteTestService) NetworkPath(ctx context.Context, testID string, params ZeroTrustDEXTracerouteTestNetworkPathParams, opts ...option.RequestOption) (res *ZeroTrustDEXTracerouteTestNetworkPathResponse, err error) {
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
func (r *ZeroTrustDEXTracerouteTestService) Percentiles(ctx context.Context, testID string, params ZeroTrustDEXTracerouteTestPercentilesParams, opts ...option.RequestOption) (res *ZeroTrustDEXTracerouteTestPercentilesResponse, err error) {
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

type ZeroTrustDEXTracerouteTestGetResponse struct {
	// The host of the Traceroute synthetic application test
	Host string `json:"host,required"`
	// The interval at which the Traceroute synthetic application test is set to run.
	Interval string                                    `json:"interval,required"`
	Kind     ZeroTrustDEXTracerouteTestGetResponseKind `json:"kind,required"`
	// The name of the Traceroute synthetic application test
	Name                  string                                                       `json:"name,required"`
	TracerouteStats       ZeroTrustDEXTracerouteTestGetResponseTracerouteStats         `json:"tracerouteStats,nullable"`
	TracerouteStatsByColo []ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColo `json:"tracerouteStatsByColo"`
	JSON                  zeroTrustDEXTracerouteTestGetResponseJSON                    `json:"-"`
}

// zeroTrustDEXTracerouteTestGetResponseJSON contains the JSON metadata for the
// struct [ZeroTrustDEXTracerouteTestGetResponse]
type zeroTrustDEXTracerouteTestGetResponseJSON struct {
	Host                  apijson.Field
	Interval              apijson.Field
	Kind                  apijson.Field
	Name                  apijson.Field
	TracerouteStats       apijson.Field
	TracerouteStatsByColo apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestGetResponseKind string

const (
	ZeroTrustDEXTracerouteTestGetResponseKindTraceroute ZeroTrustDEXTracerouteTestGetResponseKind = "traceroute"
)

type ZeroTrustDEXTracerouteTestGetResponseTracerouteStats struct {
	AvailabilityPct ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsAvailabilityPct `json:"availabilityPct,required"`
	HopsCount       ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsHopsCount       `json:"hopsCount,required"`
	PacketLossPct   ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsPacketLossPct   `json:"packetLossPct,required"`
	RoundTripTimeMs ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsRoundTripTimeMs `json:"roundTripTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                                                    `json:"uniqueDevicesTotal,required"`
	JSON               zeroTrustDEXTracerouteTestGetResponseTracerouteStatsJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestGetResponseTracerouteStatsJSON contains the JSON
// metadata for the struct [ZeroTrustDEXTracerouteTestGetResponseTracerouteStats]
type zeroTrustDEXTracerouteTestGetResponseTracerouteStatsJSON struct {
	AvailabilityPct    apijson.Field
	HopsCount          apijson.Field
	PacketLossPct      apijson.Field
	RoundTripTimeMs    apijson.Field
	UniqueDevicesTotal apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestGetResponseTracerouteStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsAvailabilityPct struct {
	Slots []ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsAvailabilityPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                                                 `json:"min,nullable"`
	JSON zeroTrustDEXTracerouteTestGetResponseTracerouteStatsAvailabilityPctJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestGetResponseTracerouteStatsAvailabilityPctJSON contains
// the JSON metadata for the struct
// [ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsAvailabilityPct]
type zeroTrustDEXTracerouteTestGetResponseTracerouteStatsAvailabilityPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsAvailabilityPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsAvailabilityPctSlot struct {
	Timestamp string                                                                      `json:"timestamp,required"`
	Value     float64                                                                     `json:"value,required"`
	JSON      zeroTrustDEXTracerouteTestGetResponseTracerouteStatsAvailabilityPctSlotJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestGetResponseTracerouteStatsAvailabilityPctSlotJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsAvailabilityPctSlot]
type zeroTrustDEXTracerouteTestGetResponseTracerouteStatsAvailabilityPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsAvailabilityPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsHopsCount struct {
	Slots []ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsHopsCountSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                             `json:"min,nullable"`
	JSON zeroTrustDEXTracerouteTestGetResponseTracerouteStatsHopsCountJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestGetResponseTracerouteStatsHopsCountJSON contains the
// JSON metadata for the struct
// [ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsHopsCount]
type zeroTrustDEXTracerouteTestGetResponseTracerouteStatsHopsCountJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsHopsCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsHopsCountSlot struct {
	Timestamp string                                                                `json:"timestamp,required"`
	Value     int64                                                                 `json:"value,required"`
	JSON      zeroTrustDEXTracerouteTestGetResponseTracerouteStatsHopsCountSlotJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestGetResponseTracerouteStatsHopsCountSlotJSON contains
// the JSON metadata for the struct
// [ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsHopsCountSlot]
type zeroTrustDEXTracerouteTestGetResponseTracerouteStatsHopsCountSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsHopsCountSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsPacketLossPct struct {
	Slots []ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsPacketLossPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                                               `json:"min,nullable"`
	JSON zeroTrustDEXTracerouteTestGetResponseTracerouteStatsPacketLossPctJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestGetResponseTracerouteStatsPacketLossPctJSON contains
// the JSON metadata for the struct
// [ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsPacketLossPct]
type zeroTrustDEXTracerouteTestGetResponseTracerouteStatsPacketLossPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsPacketLossPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsPacketLossPctSlot struct {
	Timestamp string                                                                    `json:"timestamp,required"`
	Value     float64                                                                   `json:"value,required"`
	JSON      zeroTrustDEXTracerouteTestGetResponseTracerouteStatsPacketLossPctSlotJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestGetResponseTracerouteStatsPacketLossPctSlotJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsPacketLossPctSlot]
type zeroTrustDEXTracerouteTestGetResponseTracerouteStatsPacketLossPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsPacketLossPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsRoundTripTimeMs struct {
	Slots []ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsRoundTripTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                                   `json:"min,nullable"`
	JSON zeroTrustDEXTracerouteTestGetResponseTracerouteStatsRoundTripTimeMsJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestGetResponseTracerouteStatsRoundTripTimeMsJSON contains
// the JSON metadata for the struct
// [ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsRoundTripTimeMs]
type zeroTrustDEXTracerouteTestGetResponseTracerouteStatsRoundTripTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsRoundTripTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsRoundTripTimeMsSlot struct {
	Timestamp string                                                                      `json:"timestamp,required"`
	Value     int64                                                                       `json:"value,required"`
	JSON      zeroTrustDEXTracerouteTestGetResponseTracerouteStatsRoundTripTimeMsSlotJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestGetResponseTracerouteStatsRoundTripTimeMsSlotJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsRoundTripTimeMsSlot]
type zeroTrustDEXTracerouteTestGetResponseTracerouteStatsRoundTripTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsRoundTripTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColo struct {
	AvailabilityPct ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPct `json:"availabilityPct,required"`
	Colo            string                                                                    `json:"colo,required"`
	HopsCount       ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoHopsCount       `json:"hopsCount,required"`
	PacketLossPct   ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoPacketLossPct   `json:"packetLossPct,required"`
	RoundTripTimeMs ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMs `json:"roundTripTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                                                          `json:"uniqueDevicesTotal,required"`
	JSON               zeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoJSON contains the JSON
// metadata for the struct
// [ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColo]
type zeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoJSON struct {
	AvailabilityPct    apijson.Field
	Colo               apijson.Field
	HopsCount          apijson.Field
	PacketLossPct      apijson.Field
	RoundTripTimeMs    apijson.Field
	UniqueDevicesTotal apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPct struct {
	Slots []ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                                                       `json:"min,nullable"`
	JSON zeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPctJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPctJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPct]
type zeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPctSlot struct {
	Timestamp string                                                                            `json:"timestamp,required"`
	Value     float64                                                                           `json:"value,required"`
	JSON      zeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPctSlotJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPctSlotJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPctSlot]
type zeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoAvailabilityPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoHopsCount struct {
	Slots []ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoHopsCountSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                                   `json:"min,nullable"`
	JSON zeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoHopsCountJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoHopsCountJSON contains
// the JSON metadata for the struct
// [ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoHopsCount]
type zeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoHopsCountJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoHopsCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoHopsCountSlot struct {
	Timestamp string                                                                      `json:"timestamp,required"`
	Value     int64                                                                       `json:"value,required"`
	JSON      zeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoHopsCountSlotJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoHopsCountSlotJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoHopsCountSlot]
type zeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoHopsCountSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoHopsCountSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoPacketLossPct struct {
	Slots []ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoPacketLossPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                                                     `json:"min,nullable"`
	JSON zeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoPacketLossPctJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoPacketLossPctJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoPacketLossPct]
type zeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoPacketLossPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoPacketLossPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoPacketLossPctSlot struct {
	Timestamp string                                                                          `json:"timestamp,required"`
	Value     float64                                                                         `json:"value,required"`
	JSON      zeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoPacketLossPctSlotJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoPacketLossPctSlotJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoPacketLossPctSlot]
type zeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoPacketLossPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoPacketLossPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMs struct {
	Slots []ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                                         `json:"min,nullable"`
	JSON zeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMsJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMsJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMs]
type zeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMsSlot struct {
	Timestamp string                                                                            `json:"timestamp,required"`
	Value     int64                                                                             `json:"value,required"`
	JSON      zeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMsSlotJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMsSlotJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMsSlot]
type zeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestGetResponseTracerouteStatsByColoRoundTripTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestNetworkPathResponse struct {
	// API Resource UUID tag.
	ID         string `json:"id,required"`
	DeviceName string `json:"deviceName"`
	// The interval at which the Traceroute synthetic application test is set to run.
	Interval    string                                                   `json:"interval"`
	Kind        ZeroTrustDEXTracerouteTestNetworkPathResponseKind        `json:"kind"`
	Name        string                                                   `json:"name"`
	NetworkPath ZeroTrustDEXTracerouteTestNetworkPathResponseNetworkPath `json:"networkPath,nullable"`
	// The host of the Traceroute synthetic application test
	URL  string                                            `json:"url"`
	JSON zeroTrustDEXTracerouteTestNetworkPathResponseJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestNetworkPathResponseJSON contains the JSON metadata for
// the struct [ZeroTrustDEXTracerouteTestNetworkPathResponse]
type zeroTrustDEXTracerouteTestNetworkPathResponseJSON struct {
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

func (r *ZeroTrustDEXTracerouteTestNetworkPathResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestNetworkPathResponseKind string

const (
	ZeroTrustDEXTracerouteTestNetworkPathResponseKindTraceroute ZeroTrustDEXTracerouteTestNetworkPathResponseKind = "traceroute"
)

type ZeroTrustDEXTracerouteTestNetworkPathResponseNetworkPath struct {
	Slots []ZeroTrustDEXTracerouteTestNetworkPathResponseNetworkPathSlot `json:"slots,required"`
	// Specifies the sampling applied, if any, to the slots response. When sampled,
	// results shown represent the first test run to the start of each sampling
	// interval.
	Sampling ZeroTrustDEXTracerouteTestNetworkPathResponseNetworkPathSampling `json:"sampling,nullable"`
	JSON     zeroTrustDEXTracerouteTestNetworkPathResponseNetworkPathJSON     `json:"-"`
}

// zeroTrustDEXTracerouteTestNetworkPathResponseNetworkPathJSON contains the JSON
// metadata for the struct
// [ZeroTrustDEXTracerouteTestNetworkPathResponseNetworkPath]
type zeroTrustDEXTracerouteTestNetworkPathResponseNetworkPathJSON struct {
	Slots       apijson.Field
	Sampling    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestNetworkPathResponseNetworkPath) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestNetworkPathResponseNetworkPathSlot struct {
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
	ClientToIspRTTMs int64                                                            `json:"clientToIspRttMs,nullable"`
	JSON             zeroTrustDEXTracerouteTestNetworkPathResponseNetworkPathSlotJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestNetworkPathResponseNetworkPathSlotJSON contains the
// JSON metadata for the struct
// [ZeroTrustDEXTracerouteTestNetworkPathResponseNetworkPathSlot]
type zeroTrustDEXTracerouteTestNetworkPathResponseNetworkPathSlotJSON struct {
	ID                     apijson.Field
	ClientToAppRTTMs       apijson.Field
	ClientToCfEgressRTTMs  apijson.Field
	ClientToCfIngressRTTMs apijson.Field
	Timestamp              apijson.Field
	ClientToIspRTTMs       apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestNetworkPathResponseNetworkPathSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the sampling applied, if any, to the slots response. When sampled,
// results shown represent the first test run to the start of each sampling
// interval.
type ZeroTrustDEXTracerouteTestNetworkPathResponseNetworkPathSampling struct {
	Unit  ZeroTrustDEXTracerouteTestNetworkPathResponseNetworkPathSamplingUnit `json:"unit,required"`
	Value int64                                                                `json:"value,required"`
	JSON  zeroTrustDEXTracerouteTestNetworkPathResponseNetworkPathSamplingJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestNetworkPathResponseNetworkPathSamplingJSON contains
// the JSON metadata for the struct
// [ZeroTrustDEXTracerouteTestNetworkPathResponseNetworkPathSampling]
type zeroTrustDEXTracerouteTestNetworkPathResponseNetworkPathSamplingJSON struct {
	Unit        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestNetworkPathResponseNetworkPathSampling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestNetworkPathResponseNetworkPathSamplingUnit string

const (
	ZeroTrustDEXTracerouteTestNetworkPathResponseNetworkPathSamplingUnitHours ZeroTrustDEXTracerouteTestNetworkPathResponseNetworkPathSamplingUnit = "hours"
)

type ZeroTrustDEXTracerouteTestPercentilesResponse struct {
	HopsCount       ZeroTrustDEXTracerouteTestPercentilesResponseHopsCount       `json:"hopsCount"`
	PacketLossPct   ZeroTrustDEXTracerouteTestPercentilesResponsePacketLossPct   `json:"packetLossPct"`
	RoundTripTimeMs ZeroTrustDEXTracerouteTestPercentilesResponseRoundTripTimeMs `json:"roundTripTimeMs"`
	JSON            zeroTrustDEXTracerouteTestPercentilesResponseJSON            `json:"-"`
}

// zeroTrustDEXTracerouteTestPercentilesResponseJSON contains the JSON metadata for
// the struct [ZeroTrustDEXTracerouteTestPercentilesResponse]
type zeroTrustDEXTracerouteTestPercentilesResponseJSON struct {
	HopsCount       apijson.Field
	PacketLossPct   apijson.Field
	RoundTripTimeMs apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestPercentilesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestPercentilesResponseHopsCount struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                                    `json:"p99,nullable"`
	JSON zeroTrustDEXTracerouteTestPercentilesResponseHopsCountJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestPercentilesResponseHopsCountJSON contains the JSON
// metadata for the struct [ZeroTrustDEXTracerouteTestPercentilesResponseHopsCount]
type zeroTrustDEXTracerouteTestPercentilesResponseHopsCountJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestPercentilesResponseHopsCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestPercentilesResponsePacketLossPct struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                                        `json:"p99,nullable"`
	JSON zeroTrustDEXTracerouteTestPercentilesResponsePacketLossPctJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestPercentilesResponsePacketLossPctJSON contains the JSON
// metadata for the struct
// [ZeroTrustDEXTracerouteTestPercentilesResponsePacketLossPct]
type zeroTrustDEXTracerouteTestPercentilesResponsePacketLossPctJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestPercentilesResponsePacketLossPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestPercentilesResponseRoundTripTimeMs struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                                          `json:"p99,nullable"`
	JSON zeroTrustDEXTracerouteTestPercentilesResponseRoundTripTimeMsJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestPercentilesResponseRoundTripTimeMsJSON contains the
// JSON metadata for the struct
// [ZeroTrustDEXTracerouteTestPercentilesResponseRoundTripTimeMs]
type zeroTrustDEXTracerouteTestPercentilesResponseRoundTripTimeMsJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestPercentilesResponseRoundTripTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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
	Result   ZeroTrustDEXTracerouteTestGetResponse                   `json:"result,required"`
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
	Result   ZeroTrustDEXTracerouteTestNetworkPathResponse                   `json:"result,required"`
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
	Result   ZeroTrustDEXTracerouteTestPercentilesResponse                   `json:"result,required"`
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
