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

// ZeroTrustDEXHTTPTestService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustDEXHTTPTestService]
// method instead.
type ZeroTrustDEXHTTPTestService struct {
	Options     []option.RequestOption
	Percentiles *ZeroTrustDEXHTTPTestPercentileService
}

// NewZeroTrustDEXHTTPTestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustDEXHTTPTestService(opts ...option.RequestOption) (r *ZeroTrustDEXHTTPTestService) {
	r = &ZeroTrustDEXHTTPTestService{}
	r.Options = opts
	r.Percentiles = NewZeroTrustDEXHTTPTestPercentileService(opts...)
	return
}

// Get test details and aggregate performance metrics for an http test for a given
// time period between 1 hour and 7 days.
func (r *ZeroTrustDEXHTTPTestService) Get(ctx context.Context, testID string, params ZeroTrustDEXHTTPTestGetParams, opts ...option.RequestOption) (res *DigitalExperienceMonitoringHTTPDetails, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDexhttpTestGetResponseEnvelope
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
	Name string                                     `json:"name"`
	JSON digitalExperienceMonitoringHTTPDetailsJSON `json:"-"`
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
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringHTTPDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DigitalExperienceMonitoringHTTPDetailsHTTPStats struct {
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

type DigitalExperienceMonitoringHTTPDetailsHTTPStatsByColo struct {
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

type DigitalExperienceMonitoringHTTPDetailsKind string

const (
	DigitalExperienceMonitoringHTTPDetailsKindHTTP DigitalExperienceMonitoringHTTPDetailsKind = "http"
)

type ZeroTrustDEXHTTPTestGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Time interval for aggregate time slots.
	Interval param.Field[ZeroTrustDexhttpTestGetParamsInterval] `query:"interval,required"`
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

// URLQuery serializes [ZeroTrustDEXHTTPTestGetParams]'s query parameters as
// `url.Values`.
func (r ZeroTrustDEXHTTPTestGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Time interval for aggregate time slots.
type ZeroTrustDexhttpTestGetParamsInterval string

const (
	ZeroTrustDexhttpTestGetParamsIntervalMinute ZeroTrustDexhttpTestGetParamsInterval = "minute"
	ZeroTrustDexhttpTestGetParamsIntervalHour   ZeroTrustDexhttpTestGetParamsInterval = "hour"
)

type ZeroTrustDexhttpTestGetResponseEnvelope struct {
	Errors   []ZeroTrustDexhttpTestGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDexhttpTestGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DigitalExperienceMonitoringHTTPDetails            `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustDexhttpTestGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDexhttpTestGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDexhttpTestGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustDexhttpTestGetResponseEnvelope]
type zeroTrustDexhttpTestGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDexhttpTestGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDexhttpTestGetResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zeroTrustDexhttpTestGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDexhttpTestGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZeroTrustDexhttpTestGetResponseEnvelopeErrors]
type zeroTrustDexhttpTestGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDexhttpTestGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDexhttpTestGetResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zeroTrustDexhttpTestGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDexhttpTestGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustDexhttpTestGetResponseEnvelopeMessages]
type zeroTrustDexhttpTestGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDexhttpTestGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustDexhttpTestGetResponseEnvelopeSuccess bool

const (
	ZeroTrustDexhttpTestGetResponseEnvelopeSuccessTrue ZeroTrustDexhttpTestGetResponseEnvelopeSuccess = true
)
