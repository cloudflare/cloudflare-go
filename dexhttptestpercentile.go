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
func (r *DEXHTTPTestPercentileService) List(ctx context.Context, accountID string, testID string, query DEXHTTPTestPercentileListParams, opts ...option.RequestOption) (res *DexhttpTestPercentileListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DexhttpTestPercentileListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/http-tests/%s/percentiles", accountID, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DexhttpTestPercentileListResponse struct {
	DNSResponseTimeMs    DexhttpTestPercentileListResponseDNSResponseTimeMs    `json:"dnsResponseTimeMs"`
	ResourceFetchTimeMs  DexhttpTestPercentileListResponseResourceFetchTimeMs  `json:"resourceFetchTimeMs"`
	ServerResponseTimeMs DexhttpTestPercentileListResponseServerResponseTimeMs `json:"serverResponseTimeMs"`
	JSON                 dexhttpTestPercentileListResponseJSON                 `json:"-"`
}

// dexhttpTestPercentileListResponseJSON contains the JSON metadata for the struct
// [DexhttpTestPercentileListResponse]
type dexhttpTestPercentileListResponseJSON struct {
	DNSResponseTimeMs    apijson.Field
	ResourceFetchTimeMs  apijson.Field
	ServerResponseTimeMs apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DexhttpTestPercentileListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexhttpTestPercentileListResponseDNSResponseTimeMs struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                                `json:"p99,nullable"`
	JSON dexhttpTestPercentileListResponseDNSResponseTimeMsJSON `json:"-"`
}

// dexhttpTestPercentileListResponseDNSResponseTimeMsJSON contains the JSON
// metadata for the struct [DexhttpTestPercentileListResponseDNSResponseTimeMs]
type dexhttpTestPercentileListResponseDNSResponseTimeMsJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexhttpTestPercentileListResponseDNSResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexhttpTestPercentileListResponseResourceFetchTimeMs struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                                  `json:"p99,nullable"`
	JSON dexhttpTestPercentileListResponseResourceFetchTimeMsJSON `json:"-"`
}

// dexhttpTestPercentileListResponseResourceFetchTimeMsJSON contains the JSON
// metadata for the struct [DexhttpTestPercentileListResponseResourceFetchTimeMs]
type dexhttpTestPercentileListResponseResourceFetchTimeMsJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexhttpTestPercentileListResponseResourceFetchTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexhttpTestPercentileListResponseServerResponseTimeMs struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                                   `json:"p99,nullable"`
	JSON dexhttpTestPercentileListResponseServerResponseTimeMsJSON `json:"-"`
}

// dexhttpTestPercentileListResponseServerResponseTimeMsJSON contains the JSON
// metadata for the struct [DexhttpTestPercentileListResponseServerResponseTimeMs]
type dexhttpTestPercentileListResponseServerResponseTimeMsJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexhttpTestPercentileListResponseServerResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXHTTPTestPercentileListParams struct {
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
	Result   DexhttpTestPercentileListResponse                   `json:"result,required"`
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

// Whether the API call was successful
type DexhttpTestPercentileListResponseEnvelopeSuccess bool

const (
	DexhttpTestPercentileListResponseEnvelopeSuccessTrue DexhttpTestPercentileListResponseEnvelopeSuccess = true
)
