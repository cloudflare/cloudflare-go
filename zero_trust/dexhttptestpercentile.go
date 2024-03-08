// File generated from our OpenAPI spec by Stainless.

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
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// DEXHTTPTestPercentileService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDEXHTTPTestPercentileService]
// method instead.
type DEXHTTPTestPercentileService struct {
	Options []option.RequestOption
}

// NewDEXHTTPTestPercentileService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDEXHTTPTestPercentileService(opts ...option.RequestOption) (r *DEXHTTPTestPercentileService) {
	r = &DEXHTTPTestPercentileService{}
	r.Options = opts
	return
}

// Get percentiles for an http test for a given time period between 1 hour and 7
// days.
func (r *DEXHTTPTestPercentileService) List(ctx context.Context, testID string, params DEXHTTPTestPercentileListParams, opts ...option.RequestOption) (res *DigitalExperienceMonitoringHTTPDetailsPercentiles, err error) {
	opts = append(r.Options[:], opts...)
	var env DexhttpTestPercentileListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/http-tests/%s/percentiles", params.AccountID, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DigitalExperienceMonitoringHTTPDetailsPercentiles struct {
	DNSResponseTimeMs    DigitalExperienceMonitoringHTTPDetailsPercentilesDNSResponseTimeMs    `json:"dnsResponseTimeMs"`
	ResourceFetchTimeMs  DigitalExperienceMonitoringHTTPDetailsPercentilesResourceFetchTimeMs  `json:"resourceFetchTimeMs"`
	ServerResponseTimeMs DigitalExperienceMonitoringHTTPDetailsPercentilesServerResponseTimeMs `json:"serverResponseTimeMs"`
	JSON                 digitalExperienceMonitoringHTTPDetailsPercentilesJSON                 `json:"-"`
}

// digitalExperienceMonitoringHTTPDetailsPercentilesJSON contains the JSON metadata
// for the struct [DigitalExperienceMonitoringHTTPDetailsPercentiles]
type digitalExperienceMonitoringHTTPDetailsPercentilesJSON struct {
	DNSResponseTimeMs    apijson.Field
	ResourceFetchTimeMs  apijson.Field
	ServerResponseTimeMs apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringHTTPDetailsPercentiles) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringHTTPDetailsPercentilesJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringHTTPDetailsPercentilesDNSResponseTimeMs struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                                                `json:"p99,nullable"`
	JSON digitalExperienceMonitoringHTTPDetailsPercentilesDNSResponseTimeMsJSON `json:"-"`
}

// digitalExperienceMonitoringHTTPDetailsPercentilesDNSResponseTimeMsJSON contains
// the JSON metadata for the struct
// [DigitalExperienceMonitoringHTTPDetailsPercentilesDNSResponseTimeMs]
type digitalExperienceMonitoringHTTPDetailsPercentilesDNSResponseTimeMsJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringHTTPDetailsPercentilesDNSResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringHTTPDetailsPercentilesDNSResponseTimeMsJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringHTTPDetailsPercentilesResourceFetchTimeMs struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                                                  `json:"p99,nullable"`
	JSON digitalExperienceMonitoringHTTPDetailsPercentilesResourceFetchTimeMsJSON `json:"-"`
}

// digitalExperienceMonitoringHTTPDetailsPercentilesResourceFetchTimeMsJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringHTTPDetailsPercentilesResourceFetchTimeMs]
type digitalExperienceMonitoringHTTPDetailsPercentilesResourceFetchTimeMsJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringHTTPDetailsPercentilesResourceFetchTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringHTTPDetailsPercentilesResourceFetchTimeMsJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringHTTPDetailsPercentilesServerResponseTimeMs struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                                                   `json:"p99,nullable"`
	JSON digitalExperienceMonitoringHTTPDetailsPercentilesServerResponseTimeMsJSON `json:"-"`
}

// digitalExperienceMonitoringHTTPDetailsPercentilesServerResponseTimeMsJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringHTTPDetailsPercentilesServerResponseTimeMs]
type digitalExperienceMonitoringHTTPDetailsPercentilesServerResponseTimeMsJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringHTTPDetailsPercentilesServerResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringHTTPDetailsPercentilesServerResponseTimeMsJSON) RawJSON() string {
	return r.raw
}

type DEXHTTPTestPercentileListParams struct {
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

// URLQuery serializes [DEXHTTPTestPercentileListParams]'s query parameters as
// `url.Values`.
func (r DEXHTTPTestPercentileListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DexhttpTestPercentileListResponseEnvelope struct {
	Errors   []DexhttpTestPercentileListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DexhttpTestPercentileListResponseEnvelopeMessages `json:"messages,required"`
	Result   DigitalExperienceMonitoringHTTPDetailsPercentiles   `json:"result,required"`
	// Whether the API call was successful
	Success DexhttpTestPercentileListResponseEnvelopeSuccess `json:"success,required"`
	JSON    dexhttpTestPercentileListResponseEnvelopeJSON    `json:"-"`
}

// dexhttpTestPercentileListResponseEnvelopeJSON contains the JSON metadata for the
// struct [DexhttpTestPercentileListResponseEnvelope]
type dexhttpTestPercentileListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexhttpTestPercentileListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexhttpTestPercentileListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DexhttpTestPercentileListResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    dexhttpTestPercentileListResponseEnvelopeErrorsJSON `json:"-"`
}

// dexhttpTestPercentileListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DexhttpTestPercentileListResponseEnvelopeErrors]
type dexhttpTestPercentileListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexhttpTestPercentileListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexhttpTestPercentileListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DexhttpTestPercentileListResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    dexhttpTestPercentileListResponseEnvelopeMessagesJSON `json:"-"`
}

// dexhttpTestPercentileListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DexhttpTestPercentileListResponseEnvelopeMessages]
type dexhttpTestPercentileListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexhttpTestPercentileListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexhttpTestPercentileListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DexhttpTestPercentileListResponseEnvelopeSuccess bool

const (
	DexhttpTestPercentileListResponseEnvelopeSuccessTrue DexhttpTestPercentileListResponseEnvelopeSuccess = true
)
