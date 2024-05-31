// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// DEXTracerouteTestService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDEXTracerouteTestService] method instead.
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
func (r *DEXTracerouteTestService) Get(ctx context.Context, testID string, params DEXTracerouteTestGetParams, opts ...option.RequestOption) (res *Traceroute, err error) {
	opts = append(r.Options[:], opts...)
	var env DEXTracerouteTestGetResponseEnvelope
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if testID == "" {
		err = errors.New("missing required test_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/traceroute-tests/%s", params.AccountID, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a breakdown of metrics by hop for individual traceroute test runs
func (r *DEXTracerouteTestService) NetworkPath(ctx context.Context, testID string, params DEXTracerouteTestNetworkPathParams, opts ...option.RequestOption) (res *NetworkPathResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DEXTracerouteTestNetworkPathResponseEnvelope
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if testID == "" {
		err = errors.New("missing required test_id parameter")
		return
	}
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
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if testID == "" {
		err = errors.New("missing required test_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/traceroute-tests/%s/percentiles", params.AccountID, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Traceroute struct {
	// The host of the Traceroute synthetic application test
	Host string `json:"host,required"`
	// The interval at which the Traceroute synthetic application test is set to run.
	Interval string         `json:"interval,required"`
	Kind     TracerouteKind `json:"kind,required"`
	// The name of the Traceroute synthetic application test
	Name                  string                            `json:"name,required"`
	TargetPolicies        []DeviceExperienceMonitor         `json:"target_policies,nullable"`
	Targeted              bool                              `json:"targeted"`
	TracerouteStats       TracerouteTracerouteStats         `json:"tracerouteStats,nullable"`
	TracerouteStatsByColo []TracerouteTracerouteStatsByColo `json:"tracerouteStatsByColo"`
	JSON                  tracerouteJSON                    `json:"-"`
}

// tracerouteJSON contains the JSON metadata for the struct [Traceroute]
type tracerouteJSON struct {
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

func (r *Traceroute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteJSON) RawJSON() string {
	return r.raw
}

type TracerouteKind string

const (
	TracerouteKindTraceroute TracerouteKind = "traceroute"
)

func (r TracerouteKind) IsKnown() bool {
	switch r {
	case TracerouteKindTraceroute:
		return true
	}
	return false
}

type TracerouteTracerouteStats struct {
	AvailabilityPct TracerouteTracerouteStatsAvailabilityPct `json:"availabilityPct,required"`
	HopsCount       TestStatOverTime                         `json:"hopsCount,required"`
	PacketLossPct   TracerouteTracerouteStatsPacketLossPct   `json:"packetLossPct,required"`
	RoundTripTimeMs TestStatOverTime                         `json:"roundTripTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                         `json:"uniqueDevicesTotal,required"`
	JSON               tracerouteTracerouteStatsJSON `json:"-"`
}

// tracerouteTracerouteStatsJSON contains the JSON metadata for the struct
// [TracerouteTracerouteStats]
type tracerouteTracerouteStatsJSON struct {
	AvailabilityPct    apijson.Field
	HopsCount          apijson.Field
	PacketLossPct      apijson.Field
	RoundTripTimeMs    apijson.Field
	UniqueDevicesTotal apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TracerouteTracerouteStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteTracerouteStatsJSON) RawJSON() string {
	return r.raw
}

type TracerouteTracerouteStatsAvailabilityPct struct {
	Slots []TracerouteTracerouteStatsAvailabilityPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                      `json:"min,nullable"`
	JSON tracerouteTracerouteStatsAvailabilityPctJSON `json:"-"`
}

// tracerouteTracerouteStatsAvailabilityPctJSON contains the JSON metadata for the
// struct [TracerouteTracerouteStatsAvailabilityPct]
type tracerouteTracerouteStatsAvailabilityPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TracerouteTracerouteStatsAvailabilityPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteTracerouteStatsAvailabilityPctJSON) RawJSON() string {
	return r.raw
}

type TracerouteTracerouteStatsAvailabilityPctSlot struct {
	Timestamp string                                           `json:"timestamp,required"`
	Value     float64                                          `json:"value,required"`
	JSON      tracerouteTracerouteStatsAvailabilityPctSlotJSON `json:"-"`
}

// tracerouteTracerouteStatsAvailabilityPctSlotJSON contains the JSON metadata for
// the struct [TracerouteTracerouteStatsAvailabilityPctSlot]
type tracerouteTracerouteStatsAvailabilityPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TracerouteTracerouteStatsAvailabilityPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteTracerouteStatsAvailabilityPctSlotJSON) RawJSON() string {
	return r.raw
}

type TracerouteTracerouteStatsPacketLossPct struct {
	Slots []TracerouteTracerouteStatsPacketLossPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                    `json:"min,nullable"`
	JSON tracerouteTracerouteStatsPacketLossPctJSON `json:"-"`
}

// tracerouteTracerouteStatsPacketLossPctJSON contains the JSON metadata for the
// struct [TracerouteTracerouteStatsPacketLossPct]
type tracerouteTracerouteStatsPacketLossPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TracerouteTracerouteStatsPacketLossPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteTracerouteStatsPacketLossPctJSON) RawJSON() string {
	return r.raw
}

type TracerouteTracerouteStatsPacketLossPctSlot struct {
	Timestamp string                                         `json:"timestamp,required"`
	Value     float64                                        `json:"value,required"`
	JSON      tracerouteTracerouteStatsPacketLossPctSlotJSON `json:"-"`
}

// tracerouteTracerouteStatsPacketLossPctSlotJSON contains the JSON metadata for
// the struct [TracerouteTracerouteStatsPacketLossPctSlot]
type tracerouteTracerouteStatsPacketLossPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TracerouteTracerouteStatsPacketLossPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteTracerouteStatsPacketLossPctSlotJSON) RawJSON() string {
	return r.raw
}

type TracerouteTracerouteStatsByColo struct {
	AvailabilityPct TracerouteTracerouteStatsByColoAvailabilityPct `json:"availabilityPct,required"`
	Colo            string                                         `json:"colo,required"`
	HopsCount       TestStatOverTime                               `json:"hopsCount,required"`
	PacketLossPct   TracerouteTracerouteStatsByColoPacketLossPct   `json:"packetLossPct,required"`
	RoundTripTimeMs TestStatOverTime                               `json:"roundTripTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                               `json:"uniqueDevicesTotal,required"`
	JSON               tracerouteTracerouteStatsByColoJSON `json:"-"`
}

// tracerouteTracerouteStatsByColoJSON contains the JSON metadata for the struct
// [TracerouteTracerouteStatsByColo]
type tracerouteTracerouteStatsByColoJSON struct {
	AvailabilityPct    apijson.Field
	Colo               apijson.Field
	HopsCount          apijson.Field
	PacketLossPct      apijson.Field
	RoundTripTimeMs    apijson.Field
	UniqueDevicesTotal apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TracerouteTracerouteStatsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteTracerouteStatsByColoJSON) RawJSON() string {
	return r.raw
}

type TracerouteTracerouteStatsByColoAvailabilityPct struct {
	Slots []TracerouteTracerouteStatsByColoAvailabilityPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                            `json:"min,nullable"`
	JSON tracerouteTracerouteStatsByColoAvailabilityPctJSON `json:"-"`
}

// tracerouteTracerouteStatsByColoAvailabilityPctJSON contains the JSON metadata
// for the struct [TracerouteTracerouteStatsByColoAvailabilityPct]
type tracerouteTracerouteStatsByColoAvailabilityPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TracerouteTracerouteStatsByColoAvailabilityPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteTracerouteStatsByColoAvailabilityPctJSON) RawJSON() string {
	return r.raw
}

type TracerouteTracerouteStatsByColoAvailabilityPctSlot struct {
	Timestamp string                                                 `json:"timestamp,required"`
	Value     float64                                                `json:"value,required"`
	JSON      tracerouteTracerouteStatsByColoAvailabilityPctSlotJSON `json:"-"`
}

// tracerouteTracerouteStatsByColoAvailabilityPctSlotJSON contains the JSON
// metadata for the struct [TracerouteTracerouteStatsByColoAvailabilityPctSlot]
type tracerouteTracerouteStatsByColoAvailabilityPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TracerouteTracerouteStatsByColoAvailabilityPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteTracerouteStatsByColoAvailabilityPctSlotJSON) RawJSON() string {
	return r.raw
}

type TracerouteTracerouteStatsByColoPacketLossPct struct {
	Slots []TracerouteTracerouteStatsByColoPacketLossPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                          `json:"min,nullable"`
	JSON tracerouteTracerouteStatsByColoPacketLossPctJSON `json:"-"`
}

// tracerouteTracerouteStatsByColoPacketLossPctJSON contains the JSON metadata for
// the struct [TracerouteTracerouteStatsByColoPacketLossPct]
type tracerouteTracerouteStatsByColoPacketLossPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TracerouteTracerouteStatsByColoPacketLossPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteTracerouteStatsByColoPacketLossPctJSON) RawJSON() string {
	return r.raw
}

type TracerouteTracerouteStatsByColoPacketLossPctSlot struct {
	Timestamp string                                               `json:"timestamp,required"`
	Value     float64                                              `json:"value,required"`
	JSON      tracerouteTracerouteStatsByColoPacketLossPctSlotJSON `json:"-"`
}

// tracerouteTracerouteStatsByColoPacketLossPctSlotJSON contains the JSON metadata
// for the struct [TracerouteTracerouteStatsByColoPacketLossPctSlot]
type tracerouteTracerouteStatsByColoPacketLossPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TracerouteTracerouteStatsByColoPacketLossPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteTracerouteStatsByColoPacketLossPctSlotJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestPercentilesResponse struct {
	HopsCount       Percentiles                              `json:"hopsCount"`
	PacketLossPct   Percentiles                              `json:"packetLossPct"`
	RoundTripTimeMs Percentiles                              `json:"roundTripTimeMs"`
	JSON            dexTracerouteTestPercentilesResponseJSON `json:"-"`
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

type DEXTracerouteTestGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Start time for aggregate metrics in ISO ms
	From param.Field[string] `query:"from,required"`
	// Time interval for aggregate time slots.
	Interval param.Field[DEXTracerouteTestGetParamsInterval] `query:"interval,required"`
	// End time for aggregate metrics in ISO ms
	To param.Field[string] `query:"to,required"`
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DEXTracerouteTestGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Traceroute                                  `json:"result"`
	JSON    dexTracerouteTestGetResponseEnvelopeJSON    `json:"-"`
}

// dexTracerouteTestGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DEXTracerouteTestGetResponseEnvelope]
type dexTracerouteTestGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
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
	// Start time for aggregate metrics in ISO ms
	From param.Field[string] `query:"from,required"`
	// Time interval for aggregate time slots.
	Interval param.Field[DEXTracerouteTestNetworkPathParamsInterval] `query:"interval,required"`
	// End time for aggregate metrics in ISO ms
	To param.Field[string] `query:"to,required"`
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DEXTracerouteTestNetworkPathResponseEnvelopeSuccess `json:"success,required"`
	Result  NetworkPathResponse                                 `json:"result"`
	JSON    dexTracerouteTestNetworkPathResponseEnvelopeJSON    `json:"-"`
}

// dexTracerouteTestNetworkPathResponseEnvelopeJSON contains the JSON metadata for
// the struct [DEXTracerouteTestNetworkPathResponseEnvelope]
type dexTracerouteTestNetworkPathResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
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
	// Start time for aggregate metrics in ISO format
	From param.Field[string] `query:"from,required"`
	// End time for aggregate metrics in ISO format
	To param.Field[string] `query:"to,required"`
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DEXTracerouteTestPercentilesResponseEnvelopeSuccess `json:"success,required"`
	Result  DEXTracerouteTestPercentilesResponse                `json:"result"`
	JSON    dexTracerouteTestPercentilesResponseEnvelopeJSON    `json:"-"`
}

// dexTracerouteTestPercentilesResponseEnvelopeJSON contains the JSON metadata for
// the struct [DEXTracerouteTestPercentilesResponseEnvelope]
type dexTracerouteTestPercentilesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
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
