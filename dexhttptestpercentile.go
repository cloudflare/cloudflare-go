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
	path := fmt.Sprintf("accounts/%s/dex/http-tests/%s/percentiles", accountID, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DexHTTPTestPercentileListResponse struct {
	Errors   []DexHTTPTestPercentileListResponseError   `json:"errors"`
	Messages []DexHTTPTestPercentileListResponseMessage `json:"messages"`
	Result   DexHTTPTestPercentileListResponseResult    `json:"result"`
	// Whether the API call was successful
	Success DexHTTPTestPercentileListResponseSuccess `json:"success"`
	JSON    dexHTTPTestPercentileListResponseJSON    `json:"-"`
}

// dexHTTPTestPercentileListResponseJSON contains the JSON metadata for the struct
// [DexHTTPTestPercentileListResponse]
type dexHTTPTestPercentileListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestPercentileListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestPercentileListResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    dexHTTPTestPercentileListResponseErrorJSON `json:"-"`
}

// dexHTTPTestPercentileListResponseErrorJSON contains the JSON metadata for the
// struct [DexHTTPTestPercentileListResponseError]
type dexHTTPTestPercentileListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestPercentileListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestPercentileListResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    dexHTTPTestPercentileListResponseMessageJSON `json:"-"`
}

// dexHTTPTestPercentileListResponseMessageJSON contains the JSON metadata for the
// struct [DexHTTPTestPercentileListResponseMessage]
type dexHTTPTestPercentileListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestPercentileListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestPercentileListResponseResult struct {
	DNSResponseTimeMs    DexHTTPTestPercentileListResponseResultDNSResponseTimeMs    `json:"dnsResponseTimeMs"`
	ResourceFetchTimeMs  DexHTTPTestPercentileListResponseResultResourceFetchTimeMs  `json:"resourceFetchTimeMs"`
	ServerResponseTimeMs DexHTTPTestPercentileListResponseResultServerResponseTimeMs `json:"serverResponseTimeMs"`
	JSON                 dexHTTPTestPercentileListResponseResultJSON                 `json:"-"`
}

// dexHTTPTestPercentileListResponseResultJSON contains the JSON metadata for the
// struct [DexHTTPTestPercentileListResponseResult]
type dexHTTPTestPercentileListResponseResultJSON struct {
	DNSResponseTimeMs    apijson.Field
	ResourceFetchTimeMs  apijson.Field
	ServerResponseTimeMs apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DexHTTPTestPercentileListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestPercentileListResponseResultDNSResponseTimeMs struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                                      `json:"p99,nullable"`
	JSON dexHTTPTestPercentileListResponseResultDNSResponseTimeMsJSON `json:"-"`
}

// dexHTTPTestPercentileListResponseResultDNSResponseTimeMsJSON contains the JSON
// metadata for the struct
// [DexHTTPTestPercentileListResponseResultDNSResponseTimeMs]
type dexHTTPTestPercentileListResponseResultDNSResponseTimeMsJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestPercentileListResponseResultDNSResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestPercentileListResponseResultResourceFetchTimeMs struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                                        `json:"p99,nullable"`
	JSON dexHTTPTestPercentileListResponseResultResourceFetchTimeMsJSON `json:"-"`
}

// dexHTTPTestPercentileListResponseResultResourceFetchTimeMsJSON contains the JSON
// metadata for the struct
// [DexHTTPTestPercentileListResponseResultResourceFetchTimeMs]
type dexHTTPTestPercentileListResponseResultResourceFetchTimeMsJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestPercentileListResponseResultResourceFetchTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexHTTPTestPercentileListResponseResultServerResponseTimeMs struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                                         `json:"p99,nullable"`
	JSON dexHTTPTestPercentileListResponseResultServerResponseTimeMsJSON `json:"-"`
}

// dexHTTPTestPercentileListResponseResultServerResponseTimeMsJSON contains the
// JSON metadata for the struct
// [DexHTTPTestPercentileListResponseResultServerResponseTimeMs]
type dexHTTPTestPercentileListResponseResultServerResponseTimeMsJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexHTTPTestPercentileListResponseResultServerResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DexHTTPTestPercentileListResponseSuccess bool

const (
	DexHTTPTestPercentileListResponseSuccessTrue DexHTTPTestPercentileListResponseSuccess = true
)

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
