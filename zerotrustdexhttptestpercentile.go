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

// ZeroTrustDEXHTTPTestPercentileService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZeroTrustDEXHTTPTestPercentileService] method instead.
type ZeroTrustDEXHTTPTestPercentileService struct {
	Options []option.RequestOption
}

// NewZeroTrustDEXHTTPTestPercentileService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustDEXHTTPTestPercentileService(opts ...option.RequestOption) (r *ZeroTrustDEXHTTPTestPercentileService) {
	r = &ZeroTrustDEXHTTPTestPercentileService{}
	r.Options = opts
	return
}

// Get percentiles for an http test for a given time period between 1 hour and 7
// days.
func (r *ZeroTrustDEXHTTPTestPercentileService) List(ctx context.Context, testID string, params ZeroTrustDEXHTTPTestPercentileListParams, opts ...option.RequestOption) (res *ZeroTrustDexhttpTestPercentileListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDexhttpTestPercentileListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/http-tests/%s/percentiles", params.AccountID, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustDexhttpTestPercentileListResponse struct {
	DNSResponseTimeMs    ZeroTrustDexhttpTestPercentileListResponseDNSResponseTimeMs    `json:"dnsResponseTimeMs"`
	ResourceFetchTimeMs  ZeroTrustDexhttpTestPercentileListResponseResourceFetchTimeMs  `json:"resourceFetchTimeMs"`
	ServerResponseTimeMs ZeroTrustDexhttpTestPercentileListResponseServerResponseTimeMs `json:"serverResponseTimeMs"`
	JSON                 zeroTrustDexhttpTestPercentileListResponseJSON                 `json:"-"`
}

// zeroTrustDexhttpTestPercentileListResponseJSON contains the JSON metadata for
// the struct [ZeroTrustDexhttpTestPercentileListResponse]
type zeroTrustDexhttpTestPercentileListResponseJSON struct {
	DNSResponseTimeMs    apijson.Field
	ResourceFetchTimeMs  apijson.Field
	ServerResponseTimeMs apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustDexhttpTestPercentileListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDexhttpTestPercentileListResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDexhttpTestPercentileListResponseDNSResponseTimeMs struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                                         `json:"p99,nullable"`
	JSON zeroTrustDexhttpTestPercentileListResponseDNSResponseTimeMsJSON `json:"-"`
}

// zeroTrustDexhttpTestPercentileListResponseDNSResponseTimeMsJSON contains the
// JSON metadata for the struct
// [ZeroTrustDexhttpTestPercentileListResponseDNSResponseTimeMs]
type zeroTrustDexhttpTestPercentileListResponseDNSResponseTimeMsJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDexhttpTestPercentileListResponseDNSResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDexhttpTestPercentileListResponseDNSResponseTimeMsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDexhttpTestPercentileListResponseResourceFetchTimeMs struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                                           `json:"p99,nullable"`
	JSON zeroTrustDexhttpTestPercentileListResponseResourceFetchTimeMsJSON `json:"-"`
}

// zeroTrustDexhttpTestPercentileListResponseResourceFetchTimeMsJSON contains the
// JSON metadata for the struct
// [ZeroTrustDexhttpTestPercentileListResponseResourceFetchTimeMs]
type zeroTrustDexhttpTestPercentileListResponseResourceFetchTimeMsJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDexhttpTestPercentileListResponseResourceFetchTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDexhttpTestPercentileListResponseResourceFetchTimeMsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDexhttpTestPercentileListResponseServerResponseTimeMs struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                                            `json:"p99,nullable"`
	JSON zeroTrustDexhttpTestPercentileListResponseServerResponseTimeMsJSON `json:"-"`
}

// zeroTrustDexhttpTestPercentileListResponseServerResponseTimeMsJSON contains the
// JSON metadata for the struct
// [ZeroTrustDexhttpTestPercentileListResponseServerResponseTimeMs]
type zeroTrustDexhttpTestPercentileListResponseServerResponseTimeMsJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDexhttpTestPercentileListResponseServerResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDexhttpTestPercentileListResponseServerResponseTimeMsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXHTTPTestPercentileListParams struct {
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

// URLQuery serializes [ZeroTrustDEXHTTPTestPercentileListParams]'s query
// parameters as `url.Values`.
func (r ZeroTrustDEXHTTPTestPercentileListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZeroTrustDexhttpTestPercentileListResponseEnvelope struct {
	Errors   []ZeroTrustDexhttpTestPercentileListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDexhttpTestPercentileListResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustDexhttpTestPercentileListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustDexhttpTestPercentileListResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDexhttpTestPercentileListResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDexhttpTestPercentileListResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustDexhttpTestPercentileListResponseEnvelope]
type zeroTrustDexhttpTestPercentileListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDexhttpTestPercentileListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDexhttpTestPercentileListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDexhttpTestPercentileListResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zeroTrustDexhttpTestPercentileListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDexhttpTestPercentileListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustDexhttpTestPercentileListResponseEnvelopeErrors]
type zeroTrustDexhttpTestPercentileListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDexhttpTestPercentileListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDexhttpTestPercentileListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDexhttpTestPercentileListResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    zeroTrustDexhttpTestPercentileListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDexhttpTestPercentileListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustDexhttpTestPercentileListResponseEnvelopeMessages]
type zeroTrustDexhttpTestPercentileListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDexhttpTestPercentileListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDexhttpTestPercentileListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustDexhttpTestPercentileListResponseEnvelopeSuccess bool

const (
	ZeroTrustDexhttpTestPercentileListResponseEnvelopeSuccessTrue ZeroTrustDexhttpTestPercentileListResponseEnvelopeSuccess = true
)
