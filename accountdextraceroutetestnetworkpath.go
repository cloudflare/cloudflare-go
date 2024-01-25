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

// AccountDexTracerouteTestNetworkPathService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountDexTracerouteTestNetworkPathService] method instead.
type AccountDexTracerouteTestNetworkPathService struct {
	Options []option.RequestOption
}

// NewAccountDexTracerouteTestNetworkPathService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountDexTracerouteTestNetworkPathService(opts ...option.RequestOption) (r *AccountDexTracerouteTestNetworkPathService) {
	r = &AccountDexTracerouteTestNetworkPathService{}
	r.Options = opts
	return
}

// Get a breakdown of metrics by hop for individual traceroute test runs
func (r *AccountDexTracerouteTestNetworkPathService) List(ctx context.Context, accountIdentifier string, testID string, query AccountDexTracerouteTestNetworkPathListParams, opts ...option.RequestOption) (res *AccountDexTracerouteTestNetworkPathListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dex/traceroute-tests/%s/network-path", accountIdentifier, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountDexTracerouteTestNetworkPathListResponse struct {
	Errors   []AccountDexTracerouteTestNetworkPathListResponseError   `json:"errors"`
	Messages []AccountDexTracerouteTestNetworkPathListResponseMessage `json:"messages"`
	Result   AccountDexTracerouteTestNetworkPathListResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountDexTracerouteTestNetworkPathListResponseSuccess `json:"success"`
	JSON    accountDexTracerouteTestNetworkPathListResponseJSON    `json:"-"`
}

// accountDexTracerouteTestNetworkPathListResponseJSON contains the JSON metadata
// for the struct [AccountDexTracerouteTestNetworkPathListResponse]
type accountDexTracerouteTestNetworkPathListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestNetworkPathListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestNetworkPathListResponseError struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    accountDexTracerouteTestNetworkPathListResponseErrorJSON `json:"-"`
}

// accountDexTracerouteTestNetworkPathListResponseErrorJSON contains the JSON
// metadata for the struct [AccountDexTracerouteTestNetworkPathListResponseError]
type accountDexTracerouteTestNetworkPathListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestNetworkPathListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestNetworkPathListResponseMessage struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    accountDexTracerouteTestNetworkPathListResponseMessageJSON `json:"-"`
}

// accountDexTracerouteTestNetworkPathListResponseMessageJSON contains the JSON
// metadata for the struct [AccountDexTracerouteTestNetworkPathListResponseMessage]
type accountDexTracerouteTestNetworkPathListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestNetworkPathListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestNetworkPathListResponseResult struct {
	// API Resource UUID tag.
	ID         string `json:"id,required"`
	DeviceName string `json:"deviceName"`
	// The interval at which the Traceroute synthetic application test is set to run.
	Interval    string                                                           `json:"interval"`
	Kind        AccountDexTracerouteTestNetworkPathListResponseResultKind        `json:"kind"`
	Name        string                                                           `json:"name"`
	NetworkPath AccountDexTracerouteTestNetworkPathListResponseResultNetworkPath `json:"networkPath,nullable"`
	// The host of the Traceroute synthetic application test
	URL  string                                                    `json:"url"`
	JSON accountDexTracerouteTestNetworkPathListResponseResultJSON `json:"-"`
}

// accountDexTracerouteTestNetworkPathListResponseResultJSON contains the JSON
// metadata for the struct [AccountDexTracerouteTestNetworkPathListResponseResult]
type accountDexTracerouteTestNetworkPathListResponseResultJSON struct {
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

func (r *AccountDexTracerouteTestNetworkPathListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestNetworkPathListResponseResultKind string

const (
	AccountDexTracerouteTestNetworkPathListResponseResultKindTraceroute AccountDexTracerouteTestNetworkPathListResponseResultKind = "traceroute"
)

type AccountDexTracerouteTestNetworkPathListResponseResultNetworkPath struct {
	Slots []AccountDexTracerouteTestNetworkPathListResponseResultNetworkPathSlot `json:"slots,required"`
	// Specifies the sampling applied, if any, to the slots response. When sampled,
	// results shown represent the first test run to the start of each sampling
	// interval.
	Sampling AccountDexTracerouteTestNetworkPathListResponseResultNetworkPathSampling `json:"sampling,nullable"`
	JSON     accountDexTracerouteTestNetworkPathListResponseResultNetworkPathJSON     `json:"-"`
}

// accountDexTracerouteTestNetworkPathListResponseResultNetworkPathJSON contains
// the JSON metadata for the struct
// [AccountDexTracerouteTestNetworkPathListResponseResultNetworkPath]
type accountDexTracerouteTestNetworkPathListResponseResultNetworkPathJSON struct {
	Slots       apijson.Field
	Sampling    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestNetworkPathListResponseResultNetworkPath) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestNetworkPathListResponseResultNetworkPathSlot struct {
	// API Resource UUID tag.
	ID string `json:"id,required"`
	// Round trip time in ms of the client to app mile
	ClientToAppRttMs int64 `json:"clientToAppRttMs,required,nullable"`
	// Round trip time in ms of the client to Cloudflare egress mile
	ClientToCfEgressRttMs int64 `json:"clientToCfEgressRttMs,required,nullable"`
	// Round trip time in ms of the client to Cloudflare ingress mile
	ClientToCfIngressRttMs int64  `json:"clientToCfIngressRttMs,required,nullable"`
	Timestamp              string `json:"timestamp,required"`
	// Round trip time in ms of the client to ISP mile
	ClientToIspRttMs int64                                                                    `json:"clientToIspRttMs,nullable"`
	JSON             accountDexTracerouteTestNetworkPathListResponseResultNetworkPathSlotJSON `json:"-"`
}

// accountDexTracerouteTestNetworkPathListResponseResultNetworkPathSlotJSON
// contains the JSON metadata for the struct
// [AccountDexTracerouteTestNetworkPathListResponseResultNetworkPathSlot]
type accountDexTracerouteTestNetworkPathListResponseResultNetworkPathSlotJSON struct {
	ID                     apijson.Field
	ClientToAppRttMs       apijson.Field
	ClientToCfEgressRttMs  apijson.Field
	ClientToCfIngressRttMs apijson.Field
	Timestamp              apijson.Field
	ClientToIspRttMs       apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountDexTracerouteTestNetworkPathListResponseResultNetworkPathSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the sampling applied, if any, to the slots response. When sampled,
// results shown represent the first test run to the start of each sampling
// interval.
type AccountDexTracerouteTestNetworkPathListResponseResultNetworkPathSampling struct {
	Unit  AccountDexTracerouteTestNetworkPathListResponseResultNetworkPathSamplingUnit `json:"unit,required"`
	Value int64                                                                        `json:"value,required"`
	JSON  accountDexTracerouteTestNetworkPathListResponseResultNetworkPathSamplingJSON `json:"-"`
}

// accountDexTracerouteTestNetworkPathListResponseResultNetworkPathSamplingJSON
// contains the JSON metadata for the struct
// [AccountDexTracerouteTestNetworkPathListResponseResultNetworkPathSampling]
type accountDexTracerouteTestNetworkPathListResponseResultNetworkPathSamplingJSON struct {
	Unit        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestNetworkPathListResponseResultNetworkPathSampling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestNetworkPathListResponseResultNetworkPathSamplingUnit string

const (
	AccountDexTracerouteTestNetworkPathListResponseResultNetworkPathSamplingUnitHours AccountDexTracerouteTestNetworkPathListResponseResultNetworkPathSamplingUnit = "hours"
)

// Whether the API call was successful
type AccountDexTracerouteTestNetworkPathListResponseSuccess bool

const (
	AccountDexTracerouteTestNetworkPathListResponseSuccessTrue AccountDexTracerouteTestNetworkPathListResponseSuccess = true
)

type AccountDexTracerouteTestNetworkPathListParams struct {
	// Device to filter tracroute result runs to
	DeviceID param.Field[string] `query:"deviceId,required"`
	// Time interval for aggregate time slots.
	Interval param.Field[AccountDexTracerouteTestNetworkPathListParamsInterval] `query:"interval,required"`
	// End time for aggregate metrics in ISO ms
	TimeEnd param.Field[string] `query:"timeEnd,required"`
	// Start time for aggregate metrics in ISO ms
	TimeStart param.Field[string] `query:"timeStart,required"`
}

// URLQuery serializes [AccountDexTracerouteTestNetworkPathListParams]'s query
// parameters as `url.Values`.
func (r AccountDexTracerouteTestNetworkPathListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Time interval for aggregate time slots.
type AccountDexTracerouteTestNetworkPathListParamsInterval string

const (
	AccountDexTracerouteTestNetworkPathListParamsIntervalMinute AccountDexTracerouteTestNetworkPathListParamsInterval = "minute"
	AccountDexTracerouteTestNetworkPathListParamsIntervalHour   AccountDexTracerouteTestNetworkPathListParamsInterval = "hour"
)
