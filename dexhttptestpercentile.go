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

// DexHTTPTestPercentileService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDexHTTPTestPercentileService]
// method instead.
type DexHTTPTestPercentileService struct {
	Options []option.RequestOption
}

// NewDexHTTPTestPercentileService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDexHTTPTestPercentileService(opts ...option.RequestOption) (r *DexHTTPTestPercentileService) {
	r = &DexHTTPTestPercentileService{}
	r.Options = opts
	return
}

// Get percentiles for an http test for a given time period between 1 hour and 7
// days.
func (r *DexHTTPTestPercentileService) List(ctx context.Context, accountID string, testID string, query DexHTTPTestPercentileListParams, opts ...option.RequestOption) (res *DexHTTPTestPercentileListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DexHTTPTestPercentileListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/http-tests/%s/percentiles", accountID, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DexHTTPTestPercentileListResponse struct {
	DNSResponseTimeMs    DexHTTPTestPercentileListResponseDNSResponseTimeMs    `json:"dnsResponseTimeMs"`
	ResourceFetchTimeMs  DexHTTPTestPercentileListResponseResourceFetchTimeMs  `json:"resourceFetchTimeMs"`
	ServerResponseTimeMs DexHTTPTestPercentileListResponseServerResponseTimeMs `json:"serverResponseTimeMs"`
	JSON                 dexHTTPTestPercentileListResponseJSON                 `json:"-"`
}

// dexHTTPTestPercentileListResponseJSON contains the JSON metadata for the struct
// [DexHTTPTestPercentileListResponse]
type dexHTTPTestPercentileListResponseJSON struct {
	DNSResponseTimeMs    apijson.Field
	ResourceFetchTimeMs  apijson.Field
	ServerResponseTimeMs apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DexHTTPTestPercentileListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestPercentileListResponseDNSResponseTimeMs struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                                `json:"p99,nullable"`
	JSON dexHTTPTestPercentileListResponseDNSResponseTimeMsJSON `json:"-"`
}

// dexHTTPTestPercentileListResponseDNSResponseTimeMsJSON contains the JSON
// metadata for the struct [DexHTTPTestPercentileListResponseDNSResponseTimeMs]
type dexHTTPTestPercentileListResponseDNSResponseTimeMsJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestPercentileListResponseDNSResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestPercentileListResponseResourceFetchTimeMs struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                                  `json:"p99,nullable"`
	JSON dexHTTPTestPercentileListResponseResourceFetchTimeMsJSON `json:"-"`
}

// dexHTTPTestPercentileListResponseResourceFetchTimeMsJSON contains the JSON
// metadata for the struct [DexHTTPTestPercentileListResponseResourceFetchTimeMs]
type dexHTTPTestPercentileListResponseResourceFetchTimeMsJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestPercentileListResponseResourceFetchTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestPercentileListResponseServerResponseTimeMs struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                                   `json:"p99,nullable"`
	JSON dexHTTPTestPercentileListResponseServerResponseTimeMsJSON `json:"-"`
}

// dexHTTPTestPercentileListResponseServerResponseTimeMsJSON contains the JSON
// metadata for the struct [DexHTTPTestPercentileListResponseServerResponseTimeMs]
type dexHTTPTestPercentileListResponseServerResponseTimeMsJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestPercentileListResponseServerResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestPercentileListParams struct {
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

// URLQuery serializes [DexHTTPTestPercentileListParams]'s query parameters as
// `url.Values`.
func (r DexHTTPTestPercentileListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DexHTTPTestPercentileListResponseEnvelope struct {
	Errors   []DexHTTPTestPercentileListResponseEnvelopeErrors   `json:"errors"`
	Messages []DexHTTPTestPercentileListResponseEnvelopeMessages `json:"messages"`
	Result   DexHTTPTestPercentileListResponse                   `json:"result"`
	// Whether the API call was successful
	Success DexHTTPTestPercentileListResponseEnvelopeSuccess `json:"success"`
	JSON    dexHTTPTestPercentileListResponseEnvelopeJSON    `json:"-"`
}

// dexHTTPTestPercentileListResponseEnvelopeJSON contains the JSON metadata for the
// struct [DexHTTPTestPercentileListResponseEnvelope]
type dexHTTPTestPercentileListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestPercentileListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestPercentileListResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    dexHTTPTestPercentileListResponseEnvelopeErrorsJSON `json:"-"`
}

// dexHTTPTestPercentileListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DexHTTPTestPercentileListResponseEnvelopeErrors]
type dexHTTPTestPercentileListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestPercentileListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestPercentileListResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    dexHTTPTestPercentileListResponseEnvelopeMessagesJSON `json:"-"`
}

// dexHTTPTestPercentileListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DexHTTPTestPercentileListResponseEnvelopeMessages]
type dexHTTPTestPercentileListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestPercentileListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DexHTTPTestPercentileListResponseEnvelopeSuccess bool

const (
	DexHTTPTestPercentileListResponseEnvelopeSuccessTrue DexHTTPTestPercentileListResponseEnvelopeSuccess = true
)
