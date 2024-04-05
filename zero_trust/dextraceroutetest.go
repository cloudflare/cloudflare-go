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
func (r *DEXTracerouteTestService) Get(ctx context.Context, testID string, params DEXTracerouteTestGetParams, opts ...option.RequestOption) (res *TracerouteDetails, err error) {
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
func (r *DEXTracerouteTestService) NetworkPath(ctx context.Context, testID string, params DEXTracerouteTestNetworkPathParams, opts ...option.RequestOption) (res *TracerouteTestNetworkPath, err error) {
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
func (r *DEXTracerouteTestService) Percentiles(ctx context.Context, testID string, params DEXTracerouteTestPercentilesParams, opts ...option.RequestOption) (res *TracerouteDetailsPercentiles, err error) {
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

type TracerouteDetails struct {
	// The host of the Traceroute synthetic application test
	Host string `json:"host,required"`
	// The interval at which the Traceroute synthetic application test is set to run.
	Interval string                `json:"interval,required"`
	Kind     TracerouteDetailsKind `json:"kind,required"`
	// The name of the Traceroute synthetic application test
	Name                  string                                             `json:"name,required"`
	TargetPolicies        []UnnamedSchemaRefBf9e2abcf1b78a6cab8e6e29e2228a11 `json:"target_policies"`
	Targeted              bool                                               `json:"targeted"`
	TracerouteStats       TracerouteDetailsTracerouteStats                   `json:"tracerouteStats,nullable"`
	TracerouteStatsByColo []TracerouteDetailsTracerouteStatsByColo           `json:"tracerouteStatsByColo"`
	JSON                  tracerouteDetailsJSON                              `json:"-"`
}

// tracerouteDetailsJSON contains the JSON metadata for the struct
// [TracerouteDetails]
type tracerouteDetailsJSON struct {
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

func (r *TracerouteDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteDetailsJSON) RawJSON() string {
	return r.raw
}

type TracerouteDetailsKind string

const (
	TracerouteDetailsKindTraceroute TracerouteDetailsKind = "traceroute"
)

func (r TracerouteDetailsKind) IsKnown() bool {
	switch r {
	case TracerouteDetailsKindTraceroute:
		return true
	}
	return false
}

type TracerouteDetailsTracerouteStats struct {
	AvailabilityPct TracerouteDetailsTracerouteStatsAvailabilityPct `json:"availabilityPct,required"`
	HopsCount       TestStatOverTime                                `json:"hopsCount,required"`
	PacketLossPct   TracerouteDetailsTracerouteStatsPacketLossPct   `json:"packetLossPct,required"`
	RoundTripTimeMs TestStatOverTime                                `json:"roundTripTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                                `json:"uniqueDevicesTotal,required"`
	JSON               tracerouteDetailsTracerouteStatsJSON `json:"-"`
}

// tracerouteDetailsTracerouteStatsJSON contains the JSON metadata for the struct
// [TracerouteDetailsTracerouteStats]
type tracerouteDetailsTracerouteStatsJSON struct {
	AvailabilityPct    apijson.Field
	HopsCount          apijson.Field
	PacketLossPct      apijson.Field
	RoundTripTimeMs    apijson.Field
	UniqueDevicesTotal apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TracerouteDetailsTracerouteStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteDetailsTracerouteStatsJSON) RawJSON() string {
	return r.raw
}

type TracerouteDetailsTracerouteStatsAvailabilityPct struct {
	Slots []TracerouteDetailsTracerouteStatsAvailabilityPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                             `json:"min,nullable"`
	JSON tracerouteDetailsTracerouteStatsAvailabilityPctJSON `json:"-"`
}

// tracerouteDetailsTracerouteStatsAvailabilityPctJSON contains the JSON metadata
// for the struct [TracerouteDetailsTracerouteStatsAvailabilityPct]
type tracerouteDetailsTracerouteStatsAvailabilityPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TracerouteDetailsTracerouteStatsAvailabilityPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteDetailsTracerouteStatsAvailabilityPctJSON) RawJSON() string {
	return r.raw
}

type TracerouteDetailsTracerouteStatsAvailabilityPctSlot struct {
	Timestamp string                                                  `json:"timestamp,required"`
	Value     float64                                                 `json:"value,required"`
	JSON      tracerouteDetailsTracerouteStatsAvailabilityPctSlotJSON `json:"-"`
}

// tracerouteDetailsTracerouteStatsAvailabilityPctSlotJSON contains the JSON
// metadata for the struct [TracerouteDetailsTracerouteStatsAvailabilityPctSlot]
type tracerouteDetailsTracerouteStatsAvailabilityPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TracerouteDetailsTracerouteStatsAvailabilityPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteDetailsTracerouteStatsAvailabilityPctSlotJSON) RawJSON() string {
	return r.raw
}

type TracerouteDetailsTracerouteStatsPacketLossPct struct {
	Slots []TracerouteDetailsTracerouteStatsPacketLossPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                           `json:"min,nullable"`
	JSON tracerouteDetailsTracerouteStatsPacketLossPctJSON `json:"-"`
}

// tracerouteDetailsTracerouteStatsPacketLossPctJSON contains the JSON metadata for
// the struct [TracerouteDetailsTracerouteStatsPacketLossPct]
type tracerouteDetailsTracerouteStatsPacketLossPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TracerouteDetailsTracerouteStatsPacketLossPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteDetailsTracerouteStatsPacketLossPctJSON) RawJSON() string {
	return r.raw
}

type TracerouteDetailsTracerouteStatsPacketLossPctSlot struct {
	Timestamp string                                                `json:"timestamp,required"`
	Value     float64                                               `json:"value,required"`
	JSON      tracerouteDetailsTracerouteStatsPacketLossPctSlotJSON `json:"-"`
}

// tracerouteDetailsTracerouteStatsPacketLossPctSlotJSON contains the JSON metadata
// for the struct [TracerouteDetailsTracerouteStatsPacketLossPctSlot]
type tracerouteDetailsTracerouteStatsPacketLossPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TracerouteDetailsTracerouteStatsPacketLossPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteDetailsTracerouteStatsPacketLossPctSlotJSON) RawJSON() string {
	return r.raw
}

type TracerouteDetailsTracerouteStatsByColo struct {
	AvailabilityPct TracerouteDetailsTracerouteStatsByColoAvailabilityPct `json:"availabilityPct,required"`
	Colo            string                                                `json:"colo,required"`
	HopsCount       TestStatOverTime                                      `json:"hopsCount,required"`
	PacketLossPct   TracerouteDetailsTracerouteStatsByColoPacketLossPct   `json:"packetLossPct,required"`
	RoundTripTimeMs TestStatOverTime                                      `json:"roundTripTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                                      `json:"uniqueDevicesTotal,required"`
	JSON               tracerouteDetailsTracerouteStatsByColoJSON `json:"-"`
}

// tracerouteDetailsTracerouteStatsByColoJSON contains the JSON metadata for the
// struct [TracerouteDetailsTracerouteStatsByColo]
type tracerouteDetailsTracerouteStatsByColoJSON struct {
	AvailabilityPct    apijson.Field
	Colo               apijson.Field
	HopsCount          apijson.Field
	PacketLossPct      apijson.Field
	RoundTripTimeMs    apijson.Field
	UniqueDevicesTotal apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TracerouteDetailsTracerouteStatsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteDetailsTracerouteStatsByColoJSON) RawJSON() string {
	return r.raw
}

type TracerouteDetailsTracerouteStatsByColoAvailabilityPct struct {
	Slots []TracerouteDetailsTracerouteStatsByColoAvailabilityPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                                   `json:"min,nullable"`
	JSON tracerouteDetailsTracerouteStatsByColoAvailabilityPctJSON `json:"-"`
}

// tracerouteDetailsTracerouteStatsByColoAvailabilityPctJSON contains the JSON
// metadata for the struct [TracerouteDetailsTracerouteStatsByColoAvailabilityPct]
type tracerouteDetailsTracerouteStatsByColoAvailabilityPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TracerouteDetailsTracerouteStatsByColoAvailabilityPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteDetailsTracerouteStatsByColoAvailabilityPctJSON) RawJSON() string {
	return r.raw
}

type TracerouteDetailsTracerouteStatsByColoAvailabilityPctSlot struct {
	Timestamp string                                                        `json:"timestamp,required"`
	Value     float64                                                       `json:"value,required"`
	JSON      tracerouteDetailsTracerouteStatsByColoAvailabilityPctSlotJSON `json:"-"`
}

// tracerouteDetailsTracerouteStatsByColoAvailabilityPctSlotJSON contains the JSON
// metadata for the struct
// [TracerouteDetailsTracerouteStatsByColoAvailabilityPctSlot]
type tracerouteDetailsTracerouteStatsByColoAvailabilityPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TracerouteDetailsTracerouteStatsByColoAvailabilityPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteDetailsTracerouteStatsByColoAvailabilityPctSlotJSON) RawJSON() string {
	return r.raw
}

type TracerouteDetailsTracerouteStatsByColoPacketLossPct struct {
	Slots []TracerouteDetailsTracerouteStatsByColoPacketLossPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                                 `json:"min,nullable"`
	JSON tracerouteDetailsTracerouteStatsByColoPacketLossPctJSON `json:"-"`
}

// tracerouteDetailsTracerouteStatsByColoPacketLossPctJSON contains the JSON
// metadata for the struct [TracerouteDetailsTracerouteStatsByColoPacketLossPct]
type tracerouteDetailsTracerouteStatsByColoPacketLossPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TracerouteDetailsTracerouteStatsByColoPacketLossPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteDetailsTracerouteStatsByColoPacketLossPctJSON) RawJSON() string {
	return r.raw
}

type TracerouteDetailsTracerouteStatsByColoPacketLossPctSlot struct {
	Timestamp string                                                      `json:"timestamp,required"`
	Value     float64                                                     `json:"value,required"`
	JSON      tracerouteDetailsTracerouteStatsByColoPacketLossPctSlotJSON `json:"-"`
}

// tracerouteDetailsTracerouteStatsByColoPacketLossPctSlotJSON contains the JSON
// metadata for the struct
// [TracerouteDetailsTracerouteStatsByColoPacketLossPctSlot]
type tracerouteDetailsTracerouteStatsByColoPacketLossPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TracerouteDetailsTracerouteStatsByColoPacketLossPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteDetailsTracerouteStatsByColoPacketLossPctSlotJSON) RawJSON() string {
	return r.raw
}

type TracerouteDetailsPercentiles struct {
	HopsCount       Percentiles                      `json:"hopsCount"`
	PacketLossPct   Percentiles                      `json:"packetLossPct"`
	RoundTripTimeMs Percentiles                      `json:"roundTripTimeMs"`
	JSON            tracerouteDetailsPercentilesJSON `json:"-"`
}

// tracerouteDetailsPercentilesJSON contains the JSON metadata for the struct
// [TracerouteDetailsPercentiles]
type tracerouteDetailsPercentilesJSON struct {
	HopsCount       apijson.Field
	PacketLossPct   apijson.Field
	RoundTripTimeMs apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TracerouteDetailsPercentiles) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteDetailsPercentilesJSON) RawJSON() string {
	return r.raw
}

type TracerouteTestNetworkPath struct {
	// API Resource UUID tag.
	ID         string `json:"id,required"`
	DeviceName string `json:"deviceName"`
	// The interval at which the Traceroute synthetic application test is set to run.
	Interval    string                               `json:"interval"`
	Kind        TracerouteTestNetworkPathKind        `json:"kind"`
	Name        string                               `json:"name"`
	NetworkPath TracerouteTestNetworkPathNetworkPath `json:"networkPath,nullable"`
	// The host of the Traceroute synthetic application test
	URL  string                        `json:"url"`
	JSON tracerouteTestNetworkPathJSON `json:"-"`
}

// tracerouteTestNetworkPathJSON contains the JSON metadata for the struct
// [TracerouteTestNetworkPath]
type tracerouteTestNetworkPathJSON struct {
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

func (r *TracerouteTestNetworkPath) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteTestNetworkPathJSON) RawJSON() string {
	return r.raw
}

type TracerouteTestNetworkPathKind string

const (
	TracerouteTestNetworkPathKindTraceroute TracerouteTestNetworkPathKind = "traceroute"
)

func (r TracerouteTestNetworkPathKind) IsKnown() bool {
	switch r {
	case TracerouteTestNetworkPathKindTraceroute:
		return true
	}
	return false
}

type TracerouteTestNetworkPathNetworkPath struct {
	Slots []TracerouteTestNetworkPathNetworkPathSlot `json:"slots,required"`
	// Specifies the sampling applied, if any, to the slots response. When sampled,
	// results shown represent the first test run to the start of each sampling
	// interval.
	Sampling TracerouteTestNetworkPathNetworkPathSampling `json:"sampling,nullable"`
	JSON     tracerouteTestNetworkPathNetworkPathJSON     `json:"-"`
}

// tracerouteTestNetworkPathNetworkPathJSON contains the JSON metadata for the
// struct [TracerouteTestNetworkPathNetworkPath]
type tracerouteTestNetworkPathNetworkPathJSON struct {
	Slots       apijson.Field
	Sampling    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TracerouteTestNetworkPathNetworkPath) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteTestNetworkPathNetworkPathJSON) RawJSON() string {
	return r.raw
}

type TracerouteTestNetworkPathNetworkPathSlot struct {
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
	ClientToIspRTTMs int64                                        `json:"clientToIspRttMs,nullable"`
	JSON             tracerouteTestNetworkPathNetworkPathSlotJSON `json:"-"`
}

// tracerouteTestNetworkPathNetworkPathSlotJSON contains the JSON metadata for the
// struct [TracerouteTestNetworkPathNetworkPathSlot]
type tracerouteTestNetworkPathNetworkPathSlotJSON struct {
	ID                     apijson.Field
	ClientToAppRTTMs       apijson.Field
	ClientToCfEgressRTTMs  apijson.Field
	ClientToCfIngressRTTMs apijson.Field
	Timestamp              apijson.Field
	ClientToIspRTTMs       apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *TracerouteTestNetworkPathNetworkPathSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteTestNetworkPathNetworkPathSlotJSON) RawJSON() string {
	return r.raw
}

// Specifies the sampling applied, if any, to the slots response. When sampled,
// results shown represent the first test run to the start of each sampling
// interval.
type TracerouteTestNetworkPathNetworkPathSampling struct {
	Unit  TracerouteTestNetworkPathNetworkPathSamplingUnit `json:"unit,required"`
	Value int64                                            `json:"value,required"`
	JSON  tracerouteTestNetworkPathNetworkPathSamplingJSON `json:"-"`
}

// tracerouteTestNetworkPathNetworkPathSamplingJSON contains the JSON metadata for
// the struct [TracerouteTestNetworkPathNetworkPathSampling]
type tracerouteTestNetworkPathNetworkPathSamplingJSON struct {
	Unit        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TracerouteTestNetworkPathNetworkPathSampling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteTestNetworkPathNetworkPathSamplingJSON) RawJSON() string {
	return r.raw
}

type TracerouteTestNetworkPathNetworkPathSamplingUnit string

const (
	TracerouteTestNetworkPathNetworkPathSamplingUnitHours TracerouteTestNetworkPathNetworkPathSamplingUnit = "hours"
)

func (r TracerouteTestNetworkPathNetworkPathSamplingUnit) IsKnown() bool {
	switch r {
	case TracerouteTestNetworkPathNetworkPathSamplingUnitHours:
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
	Result   TracerouteDetails                                         `json:"result,required"`
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
	Result   TracerouteTestNetworkPath                                 `json:"result,required"`
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
	Result   TracerouteDetailsPercentiles                              `json:"result,required"`
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
