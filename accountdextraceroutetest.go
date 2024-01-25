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

// AccountDexTracerouteTestService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountDexTracerouteTestService] method instead.
type AccountDexTracerouteTestService struct {
	Options     []option.RequestOption
	NetworkPath *AccountDexTracerouteTestNetworkPathService
}

// NewAccountDexTracerouteTestService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountDexTracerouteTestService(opts ...option.RequestOption) (r *AccountDexTracerouteTestService) {
	r = &AccountDexTracerouteTestService{}
	r.Options = opts
	r.NetworkPath = NewAccountDexTracerouteTestNetworkPathService(opts...)
	return
}

// Get test details and aggregate performance metrics for an traceroute test for a
// given time period between 1 hour and 7 days.
func (r *AccountDexTracerouteTestService) Get(ctx context.Context, accountIdentifier string, testID string, query AccountDexTracerouteTestGetParams, opts ...option.RequestOption) (res *AccountDexTracerouteTestGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dex/traceroute-tests/%s", accountIdentifier, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get percentiles for a traceroute test for a given time period between 1 hour and
// 7 days.
func (r *AccountDexTracerouteTestService) Percentiles(ctx context.Context, accountIdentifier string, testID string, query AccountDexTracerouteTestPercentilesParams, opts ...option.RequestOption) (res *AccountDexTracerouteTestPercentilesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dex/traceroute-tests/%s/percentiles", accountIdentifier, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountDexTracerouteTestGetResponse struct {
	Errors   []AccountDexTracerouteTestGetResponseError   `json:"errors"`
	Messages []AccountDexTracerouteTestGetResponseMessage `json:"messages"`
	Result   AccountDexTracerouteTestGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountDexTracerouteTestGetResponseSuccess `json:"success"`
	JSON    accountDexTracerouteTestGetResponseJSON    `json:"-"`
}

// accountDexTracerouteTestGetResponseJSON contains the JSON metadata for the
// struct [AccountDexTracerouteTestGetResponse]
type accountDexTracerouteTestGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestGetResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountDexTracerouteTestGetResponseErrorJSON `json:"-"`
}

// accountDexTracerouteTestGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountDexTracerouteTestGetResponseError]
type accountDexTracerouteTestGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestGetResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountDexTracerouteTestGetResponseMessageJSON `json:"-"`
}

// accountDexTracerouteTestGetResponseMessageJSON contains the JSON metadata for
// the struct [AccountDexTracerouteTestGetResponseMessage]
type accountDexTracerouteTestGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestGetResponseResult struct {
	// The host of the Traceroute synthetic application test
	Host string `json:"host,required"`
	// The interval at which the Traceroute synthetic application test is set to run.
	Interval string                                        `json:"interval,required"`
	Kind     AccountDexTracerouteTestGetResponseResultKind `json:"kind,required"`
	// The name of the Traceroute synthetic application test
	Name                  string                                                           `json:"name,required"`
	TracerouteStats       AccountDexTracerouteTestGetResponseResultTracerouteStats         `json:"tracerouteStats,nullable"`
	TracerouteStatsByColo []AccountDexTracerouteTestGetResponseResultTracerouteStatsByColo `json:"tracerouteStatsByColo"`
	JSON                  accountDexTracerouteTestGetResponseResultJSON                    `json:"-"`
}

// accountDexTracerouteTestGetResponseResultJSON contains the JSON metadata for the
// struct [AccountDexTracerouteTestGetResponseResult]
type accountDexTracerouteTestGetResponseResultJSON struct {
	Host                  apijson.Field
	Interval              apijson.Field
	Kind                  apijson.Field
	Name                  apijson.Field
	TracerouteStats       apijson.Field
	TracerouteStatsByColo apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestGetResponseResultKind string

const (
	AccountDexTracerouteTestGetResponseResultKindTraceroute AccountDexTracerouteTestGetResponseResultKind = "traceroute"
)

type AccountDexTracerouteTestGetResponseResultTracerouteStats struct {
	AvailabilityPct AccountDexTracerouteTestGetResponseResultTracerouteStatsAvailabilityPct `json:"availabilityPct,required"`
	HopsCount       AccountDexTracerouteTestGetResponseResultTracerouteStatsHopsCount       `json:"hopsCount,required"`
	PacketLossPct   AccountDexTracerouteTestGetResponseResultTracerouteStatsPacketLossPct   `json:"packetLossPct,required"`
	RoundTripTimeMs AccountDexTracerouteTestGetResponseResultTracerouteStatsRoundTripTimeMs `json:"roundTripTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                                                        `json:"uniqueDevicesTotal,required"`
	JSON               accountDexTracerouteTestGetResponseResultTracerouteStatsJSON `json:"-"`
}

// accountDexTracerouteTestGetResponseResultTracerouteStatsJSON contains the JSON
// metadata for the struct
// [AccountDexTracerouteTestGetResponseResultTracerouteStats]
type accountDexTracerouteTestGetResponseResultTracerouteStatsJSON struct {
	AvailabilityPct    apijson.Field
	HopsCount          apijson.Field
	PacketLossPct      apijson.Field
	RoundTripTimeMs    apijson.Field
	UniqueDevicesTotal apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetResponseResultTracerouteStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestGetResponseResultTracerouteStatsAvailabilityPct struct {
	Slots []AccountDexTracerouteTestGetResponseResultTracerouteStatsAvailabilityPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                                                     `json:"min,nullable"`
	JSON accountDexTracerouteTestGetResponseResultTracerouteStatsAvailabilityPctJSON `json:"-"`
}

// accountDexTracerouteTestGetResponseResultTracerouteStatsAvailabilityPctJSON
// contains the JSON metadata for the struct
// [AccountDexTracerouteTestGetResponseResultTracerouteStatsAvailabilityPct]
type accountDexTracerouteTestGetResponseResultTracerouteStatsAvailabilityPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetResponseResultTracerouteStatsAvailabilityPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestGetResponseResultTracerouteStatsAvailabilityPctSlot struct {
	Timestamp string                                                                          `json:"timestamp,required"`
	Value     float64                                                                         `json:"value,required"`
	JSON      accountDexTracerouteTestGetResponseResultTracerouteStatsAvailabilityPctSlotJSON `json:"-"`
}

// accountDexTracerouteTestGetResponseResultTracerouteStatsAvailabilityPctSlotJSON
// contains the JSON metadata for the struct
// [AccountDexTracerouteTestGetResponseResultTracerouteStatsAvailabilityPctSlot]
type accountDexTracerouteTestGetResponseResultTracerouteStatsAvailabilityPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetResponseResultTracerouteStatsAvailabilityPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestGetResponseResultTracerouteStatsHopsCount struct {
	Slots []AccountDexTracerouteTestGetResponseResultTracerouteStatsHopsCountSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                                 `json:"min,nullable"`
	JSON accountDexTracerouteTestGetResponseResultTracerouteStatsHopsCountJSON `json:"-"`
}

// accountDexTracerouteTestGetResponseResultTracerouteStatsHopsCountJSON contains
// the JSON metadata for the struct
// [AccountDexTracerouteTestGetResponseResultTracerouteStatsHopsCount]
type accountDexTracerouteTestGetResponseResultTracerouteStatsHopsCountJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetResponseResultTracerouteStatsHopsCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestGetResponseResultTracerouteStatsHopsCountSlot struct {
	Timestamp string                                                                    `json:"timestamp,required"`
	Value     int64                                                                     `json:"value,required"`
	JSON      accountDexTracerouteTestGetResponseResultTracerouteStatsHopsCountSlotJSON `json:"-"`
}

// accountDexTracerouteTestGetResponseResultTracerouteStatsHopsCountSlotJSON
// contains the JSON metadata for the struct
// [AccountDexTracerouteTestGetResponseResultTracerouteStatsHopsCountSlot]
type accountDexTracerouteTestGetResponseResultTracerouteStatsHopsCountSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetResponseResultTracerouteStatsHopsCountSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestGetResponseResultTracerouteStatsPacketLossPct struct {
	Slots []AccountDexTracerouteTestGetResponseResultTracerouteStatsPacketLossPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                                                   `json:"min,nullable"`
	JSON accountDexTracerouteTestGetResponseResultTracerouteStatsPacketLossPctJSON `json:"-"`
}

// accountDexTracerouteTestGetResponseResultTracerouteStatsPacketLossPctJSON
// contains the JSON metadata for the struct
// [AccountDexTracerouteTestGetResponseResultTracerouteStatsPacketLossPct]
type accountDexTracerouteTestGetResponseResultTracerouteStatsPacketLossPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetResponseResultTracerouteStatsPacketLossPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestGetResponseResultTracerouteStatsPacketLossPctSlot struct {
	Timestamp string                                                                        `json:"timestamp,required"`
	Value     float64                                                                       `json:"value,required"`
	JSON      accountDexTracerouteTestGetResponseResultTracerouteStatsPacketLossPctSlotJSON `json:"-"`
}

// accountDexTracerouteTestGetResponseResultTracerouteStatsPacketLossPctSlotJSON
// contains the JSON metadata for the struct
// [AccountDexTracerouteTestGetResponseResultTracerouteStatsPacketLossPctSlot]
type accountDexTracerouteTestGetResponseResultTracerouteStatsPacketLossPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetResponseResultTracerouteStatsPacketLossPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestGetResponseResultTracerouteStatsRoundTripTimeMs struct {
	Slots []AccountDexTracerouteTestGetResponseResultTracerouteStatsRoundTripTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                                       `json:"min,nullable"`
	JSON accountDexTracerouteTestGetResponseResultTracerouteStatsRoundTripTimeMsJSON `json:"-"`
}

// accountDexTracerouteTestGetResponseResultTracerouteStatsRoundTripTimeMsJSON
// contains the JSON metadata for the struct
// [AccountDexTracerouteTestGetResponseResultTracerouteStatsRoundTripTimeMs]
type accountDexTracerouteTestGetResponseResultTracerouteStatsRoundTripTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetResponseResultTracerouteStatsRoundTripTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestGetResponseResultTracerouteStatsRoundTripTimeMsSlot struct {
	Timestamp string                                                                          `json:"timestamp,required"`
	Value     int64                                                                           `json:"value,required"`
	JSON      accountDexTracerouteTestGetResponseResultTracerouteStatsRoundTripTimeMsSlotJSON `json:"-"`
}

// accountDexTracerouteTestGetResponseResultTracerouteStatsRoundTripTimeMsSlotJSON
// contains the JSON metadata for the struct
// [AccountDexTracerouteTestGetResponseResultTracerouteStatsRoundTripTimeMsSlot]
type accountDexTracerouteTestGetResponseResultTracerouteStatsRoundTripTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetResponseResultTracerouteStatsRoundTripTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestGetResponseResultTracerouteStatsByColo struct {
	AvailabilityPct AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoAvailabilityPct `json:"availabilityPct,required"`
	Colo            string                                                                        `json:"colo,required"`
	HopsCount       AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoHopsCount       `json:"hopsCount,required"`
	PacketLossPct   AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoPacketLossPct   `json:"packetLossPct,required"`
	RoundTripTimeMs AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoRoundTripTimeMs `json:"roundTripTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                                                              `json:"uniqueDevicesTotal,required"`
	JSON               accountDexTracerouteTestGetResponseResultTracerouteStatsByColoJSON `json:"-"`
}

// accountDexTracerouteTestGetResponseResultTracerouteStatsByColoJSON contains the
// JSON metadata for the struct
// [AccountDexTracerouteTestGetResponseResultTracerouteStatsByColo]
type accountDexTracerouteTestGetResponseResultTracerouteStatsByColoJSON struct {
	AvailabilityPct    apijson.Field
	Colo               apijson.Field
	HopsCount          apijson.Field
	PacketLossPct      apijson.Field
	RoundTripTimeMs    apijson.Field
	UniqueDevicesTotal apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetResponseResultTracerouteStatsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoAvailabilityPct struct {
	Slots []AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoAvailabilityPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                                                           `json:"min,nullable"`
	JSON accountDexTracerouteTestGetResponseResultTracerouteStatsByColoAvailabilityPctJSON `json:"-"`
}

// accountDexTracerouteTestGetResponseResultTracerouteStatsByColoAvailabilityPctJSON
// contains the JSON metadata for the struct
// [AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoAvailabilityPct]
type accountDexTracerouteTestGetResponseResultTracerouteStatsByColoAvailabilityPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoAvailabilityPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoAvailabilityPctSlot struct {
	Timestamp string                                                                                `json:"timestamp,required"`
	Value     float64                                                                               `json:"value,required"`
	JSON      accountDexTracerouteTestGetResponseResultTracerouteStatsByColoAvailabilityPctSlotJSON `json:"-"`
}

// accountDexTracerouteTestGetResponseResultTracerouteStatsByColoAvailabilityPctSlotJSON
// contains the JSON metadata for the struct
// [AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoAvailabilityPctSlot]
type accountDexTracerouteTestGetResponseResultTracerouteStatsByColoAvailabilityPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoAvailabilityPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoHopsCount struct {
	Slots []AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoHopsCountSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                                       `json:"min,nullable"`
	JSON accountDexTracerouteTestGetResponseResultTracerouteStatsByColoHopsCountJSON `json:"-"`
}

// accountDexTracerouteTestGetResponseResultTracerouteStatsByColoHopsCountJSON
// contains the JSON metadata for the struct
// [AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoHopsCount]
type accountDexTracerouteTestGetResponseResultTracerouteStatsByColoHopsCountJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoHopsCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoHopsCountSlot struct {
	Timestamp string                                                                          `json:"timestamp,required"`
	Value     int64                                                                           `json:"value,required"`
	JSON      accountDexTracerouteTestGetResponseResultTracerouteStatsByColoHopsCountSlotJSON `json:"-"`
}

// accountDexTracerouteTestGetResponseResultTracerouteStatsByColoHopsCountSlotJSON
// contains the JSON metadata for the struct
// [AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoHopsCountSlot]
type accountDexTracerouteTestGetResponseResultTracerouteStatsByColoHopsCountSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoHopsCountSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoPacketLossPct struct {
	Slots []AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoPacketLossPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                                                         `json:"min,nullable"`
	JSON accountDexTracerouteTestGetResponseResultTracerouteStatsByColoPacketLossPctJSON `json:"-"`
}

// accountDexTracerouteTestGetResponseResultTracerouteStatsByColoPacketLossPctJSON
// contains the JSON metadata for the struct
// [AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoPacketLossPct]
type accountDexTracerouteTestGetResponseResultTracerouteStatsByColoPacketLossPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoPacketLossPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoPacketLossPctSlot struct {
	Timestamp string                                                                              `json:"timestamp,required"`
	Value     float64                                                                             `json:"value,required"`
	JSON      accountDexTracerouteTestGetResponseResultTracerouteStatsByColoPacketLossPctSlotJSON `json:"-"`
}

// accountDexTracerouteTestGetResponseResultTracerouteStatsByColoPacketLossPctSlotJSON
// contains the JSON metadata for the struct
// [AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoPacketLossPctSlot]
type accountDexTracerouteTestGetResponseResultTracerouteStatsByColoPacketLossPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoPacketLossPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoRoundTripTimeMs struct {
	Slots []AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoRoundTripTimeMsSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                                                                             `json:"min,nullable"`
	JSON accountDexTracerouteTestGetResponseResultTracerouteStatsByColoRoundTripTimeMsJSON `json:"-"`
}

// accountDexTracerouteTestGetResponseResultTracerouteStatsByColoRoundTripTimeMsJSON
// contains the JSON metadata for the struct
// [AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoRoundTripTimeMs]
type accountDexTracerouteTestGetResponseResultTracerouteStatsByColoRoundTripTimeMsJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoRoundTripTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoRoundTripTimeMsSlot struct {
	Timestamp string                                                                                `json:"timestamp,required"`
	Value     int64                                                                                 `json:"value,required"`
	JSON      accountDexTracerouteTestGetResponseResultTracerouteStatsByColoRoundTripTimeMsSlotJSON `json:"-"`
}

// accountDexTracerouteTestGetResponseResultTracerouteStatsByColoRoundTripTimeMsSlotJSON
// contains the JSON metadata for the struct
// [AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoRoundTripTimeMsSlot]
type accountDexTracerouteTestGetResponseResultTracerouteStatsByColoRoundTripTimeMsSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetResponseResultTracerouteStatsByColoRoundTripTimeMsSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountDexTracerouteTestGetResponseSuccess bool

const (
	AccountDexTracerouteTestGetResponseSuccessTrue AccountDexTracerouteTestGetResponseSuccess = true
)

type AccountDexTracerouteTestPercentilesResponse struct {
	Errors   []AccountDexTracerouteTestPercentilesResponseError   `json:"errors"`
	Messages []AccountDexTracerouteTestPercentilesResponseMessage `json:"messages"`
	Result   AccountDexTracerouteTestPercentilesResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountDexTracerouteTestPercentilesResponseSuccess `json:"success"`
	JSON    accountDexTracerouteTestPercentilesResponseJSON    `json:"-"`
}

// accountDexTracerouteTestPercentilesResponseJSON contains the JSON metadata for
// the struct [AccountDexTracerouteTestPercentilesResponse]
type accountDexTracerouteTestPercentilesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestPercentilesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestPercentilesResponseError struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    accountDexTracerouteTestPercentilesResponseErrorJSON `json:"-"`
}

// accountDexTracerouteTestPercentilesResponseErrorJSON contains the JSON metadata
// for the struct [AccountDexTracerouteTestPercentilesResponseError]
type accountDexTracerouteTestPercentilesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestPercentilesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestPercentilesResponseMessage struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    accountDexTracerouteTestPercentilesResponseMessageJSON `json:"-"`
}

// accountDexTracerouteTestPercentilesResponseMessageJSON contains the JSON
// metadata for the struct [AccountDexTracerouteTestPercentilesResponseMessage]
type accountDexTracerouteTestPercentilesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestPercentilesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestPercentilesResponseResult struct {
	HopsCount       AccountDexTracerouteTestPercentilesResponseResultHopsCount       `json:"hopsCount"`
	PacketLossPct   AccountDexTracerouteTestPercentilesResponseResultPacketLossPct   `json:"packetLossPct"`
	RoundTripTimeMs AccountDexTracerouteTestPercentilesResponseResultRoundTripTimeMs `json:"roundTripTimeMs"`
	JSON            accountDexTracerouteTestPercentilesResponseResultJSON            `json:"-"`
}

// accountDexTracerouteTestPercentilesResponseResultJSON contains the JSON metadata
// for the struct [AccountDexTracerouteTestPercentilesResponseResult]
type accountDexTracerouteTestPercentilesResponseResultJSON struct {
	HopsCount       apijson.Field
	PacketLossPct   apijson.Field
	RoundTripTimeMs apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDexTracerouteTestPercentilesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestPercentilesResponseResultHopsCount struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                                        `json:"p99,nullable"`
	JSON accountDexTracerouteTestPercentilesResponseResultHopsCountJSON `json:"-"`
}

// accountDexTracerouteTestPercentilesResponseResultHopsCountJSON contains the JSON
// metadata for the struct
// [AccountDexTracerouteTestPercentilesResponseResultHopsCount]
type accountDexTracerouteTestPercentilesResponseResultHopsCountJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestPercentilesResponseResultHopsCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestPercentilesResponseResultPacketLossPct struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                                            `json:"p99,nullable"`
	JSON accountDexTracerouteTestPercentilesResponseResultPacketLossPctJSON `json:"-"`
}

// accountDexTracerouteTestPercentilesResponseResultPacketLossPctJSON contains the
// JSON metadata for the struct
// [AccountDexTracerouteTestPercentilesResponseResultPacketLossPct]
type accountDexTracerouteTestPercentilesResponseResultPacketLossPctJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestPercentilesResponseResultPacketLossPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestPercentilesResponseResultRoundTripTimeMs struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                                              `json:"p99,nullable"`
	JSON accountDexTracerouteTestPercentilesResponseResultRoundTripTimeMsJSON `json:"-"`
}

// accountDexTracerouteTestPercentilesResponseResultRoundTripTimeMsJSON contains
// the JSON metadata for the struct
// [AccountDexTracerouteTestPercentilesResponseResultRoundTripTimeMs]
type accountDexTracerouteTestPercentilesResponseResultRoundTripTimeMsJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestPercentilesResponseResultRoundTripTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountDexTracerouteTestPercentilesResponseSuccess bool

const (
	AccountDexTracerouteTestPercentilesResponseSuccessTrue AccountDexTracerouteTestPercentilesResponseSuccess = true
)

type AccountDexTracerouteTestGetParams struct {
	// Time interval for aggregate time slots.
	Interval param.Field[AccountDexTracerouteTestGetParamsInterval] `query:"interval,required"`
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

// URLQuery serializes [AccountDexTracerouteTestGetParams]'s query parameters as
// `url.Values`.
func (r AccountDexTracerouteTestGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Time interval for aggregate time slots.
type AccountDexTracerouteTestGetParamsInterval string

const (
	AccountDexTracerouteTestGetParamsIntervalMinute AccountDexTracerouteTestGetParamsInterval = "minute"
	AccountDexTracerouteTestGetParamsIntervalHour   AccountDexTracerouteTestGetParamsInterval = "hour"
)

type AccountDexTracerouteTestPercentilesParams struct {
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

// URLQuery serializes [AccountDexTracerouteTestPercentilesParams]'s query
// parameters as `url.Values`.
func (r AccountDexTracerouteTestPercentilesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
