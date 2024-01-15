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

// AccountDexHTTPTestPercentileService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountDexHTTPTestPercentileService] method instead.
type AccountDexHTTPTestPercentileService struct {
	Options []option.RequestOption
}

// NewAccountDexHTTPTestPercentileService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountDexHTTPTestPercentileService(opts ...option.RequestOption) (r *AccountDexHTTPTestPercentileService) {
	r = &AccountDexHTTPTestPercentileService{}
	r.Options = opts
	return
}

// Get percentiles for an http test for a given time period between 1 hour and 7
// days.
func (r *AccountDexHTTPTestPercentileService) List(ctx context.Context, accountIdentifier string, testID string, query AccountDexHTTPTestPercentileListParams, opts ...option.RequestOption) (res *AccountDexHTTPTestPercentileListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dex/http-tests/%s/percentiles", accountIdentifier, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountDexHTTPTestPercentileListResponse struct {
	Errors   []AccountDexHTTPTestPercentileListResponseError   `json:"errors"`
	Messages []AccountDexHTTPTestPercentileListResponseMessage `json:"messages"`
	Result   AccountDexHTTPTestPercentileListResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountDexHTTPTestPercentileListResponseSuccess `json:"success"`
	JSON    accountDexHTTPTestPercentileListResponseJSON    `json:"-"`
}

// accountDexHTTPTestPercentileListResponseJSON contains the JSON metadata for the
// struct [AccountDexHTTPTestPercentileListResponse]
type accountDexHTTPTestPercentileListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexHTTPTestPercentileListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexHTTPTestPercentileListResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountDexHTTPTestPercentileListResponseErrorJSON `json:"-"`
}

// accountDexHTTPTestPercentileListResponseErrorJSON contains the JSON metadata for
// the struct [AccountDexHTTPTestPercentileListResponseError]
type accountDexHTTPTestPercentileListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexHTTPTestPercentileListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexHTTPTestPercentileListResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountDexHTTPTestPercentileListResponseMessageJSON `json:"-"`
}

// accountDexHTTPTestPercentileListResponseMessageJSON contains the JSON metadata
// for the struct [AccountDexHTTPTestPercentileListResponseMessage]
type accountDexHTTPTestPercentileListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexHTTPTestPercentileListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexHTTPTestPercentileListResponseResult struct {
	DNSResponseTimeMs    AccountDexHTTPTestPercentileListResponseResultDNSResponseTimeMs    `json:"dnsResponseTimeMs"`
	ResourceFetchTimeMs  AccountDexHTTPTestPercentileListResponseResultResourceFetchTimeMs  `json:"resourceFetchTimeMs"`
	ServerResponseTimeMs AccountDexHTTPTestPercentileListResponseResultServerResponseTimeMs `json:"serverResponseTimeMs"`
	JSON                 accountDexHTTPTestPercentileListResponseResultJSON                 `json:"-"`
}

// accountDexHTTPTestPercentileListResponseResultJSON contains the JSON metadata
// for the struct [AccountDexHTTPTestPercentileListResponseResult]
type accountDexHTTPTestPercentileListResponseResultJSON struct {
	DNSResponseTimeMs    apijson.Field
	ResourceFetchTimeMs  apijson.Field
	ServerResponseTimeMs apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountDexHTTPTestPercentileListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexHTTPTestPercentileListResponseResultDNSResponseTimeMs struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                                             `json:"p99,nullable"`
	JSON accountDexHTTPTestPercentileListResponseResultDNSResponseTimeMsJSON `json:"-"`
}

// accountDexHTTPTestPercentileListResponseResultDNSResponseTimeMsJSON contains the
// JSON metadata for the struct
// [AccountDexHTTPTestPercentileListResponseResultDNSResponseTimeMs]
type accountDexHTTPTestPercentileListResponseResultDNSResponseTimeMsJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexHTTPTestPercentileListResponseResultDNSResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexHTTPTestPercentileListResponseResultResourceFetchTimeMs struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                                               `json:"p99,nullable"`
	JSON accountDexHTTPTestPercentileListResponseResultResourceFetchTimeMsJSON `json:"-"`
}

// accountDexHTTPTestPercentileListResponseResultResourceFetchTimeMsJSON contains
// the JSON metadata for the struct
// [AccountDexHTTPTestPercentileListResponseResultResourceFetchTimeMs]
type accountDexHTTPTestPercentileListResponseResultResourceFetchTimeMsJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexHTTPTestPercentileListResponseResultResourceFetchTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexHTTPTestPercentileListResponseResultServerResponseTimeMs struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64                                                                `json:"p99,nullable"`
	JSON accountDexHTTPTestPercentileListResponseResultServerResponseTimeMsJSON `json:"-"`
}

// accountDexHTTPTestPercentileListResponseResultServerResponseTimeMsJSON contains
// the JSON metadata for the struct
// [AccountDexHTTPTestPercentileListResponseResultServerResponseTimeMs]
type accountDexHTTPTestPercentileListResponseResultServerResponseTimeMsJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexHTTPTestPercentileListResponseResultServerResponseTimeMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountDexHTTPTestPercentileListResponseSuccess bool

const (
	AccountDexHTTPTestPercentileListResponseSuccessTrue AccountDexHTTPTestPercentileListResponseSuccess = true
)

type AccountDexHTTPTestPercentileListParams struct {
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

// URLQuery serializes [AccountDexHTTPTestPercentileListParams]'s query parameters
// as `url.Values`.
func (r AccountDexHTTPTestPercentileListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
