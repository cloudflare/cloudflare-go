// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// ZeroTrustDEXTestUniqueDeviceService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZeroTrustDEXTestUniqueDeviceService] method instead.
type ZeroTrustDEXTestUniqueDeviceService struct {
	Options []option.RequestOption
}

// NewZeroTrustDEXTestUniqueDeviceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustDEXTestUniqueDeviceService(opts ...option.RequestOption) (r *ZeroTrustDEXTestUniqueDeviceService) {
	r = &ZeroTrustDEXTestUniqueDeviceService{}
	r.Options = opts
	return
}

// Returns unique count of devices that have run synthetic application monitoring
// tests in the past 7 days.
func (r *ZeroTrustDEXTestUniqueDeviceService) List(ctx context.Context, params ZeroTrustDEXTestUniqueDeviceListParams, opts ...option.RequestOption) (res *ZeroTrustDEXTestUniqueDeviceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDEXTestUniqueDeviceListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/tests/unique-devices", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustDEXTestUniqueDeviceListResponse struct {
	// total number of unique devices
	UniqueDevicesTotal int64                                        `json:"uniqueDevicesTotal,required"`
	JSON               zeroTrustDEXTestUniqueDeviceListResponseJSON `json:"-"`
}

// zeroTrustDEXTestUniqueDeviceListResponseJSON contains the JSON metadata for the
// struct [ZeroTrustDEXTestUniqueDeviceListResponse]
type zeroTrustDEXTestUniqueDeviceListResponseJSON struct {
	UniqueDevicesTotal apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustDEXTestUniqueDeviceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestUniqueDeviceListResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestUniqueDeviceListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Optionally filter result stats to a specific device(s). Cannot be used in
	// combination with colo param.
	DeviceID param.Field[[]string] `query:"deviceId"`
	// Optionally filter results by test name
	TestName param.Field[string] `query:"testName"`
}

// URLQuery serializes [ZeroTrustDEXTestUniqueDeviceListParams]'s query parameters
// as `url.Values`.
func (r ZeroTrustDEXTestUniqueDeviceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZeroTrustDEXTestUniqueDeviceListResponseEnvelope struct {
	Errors   []ZeroTrustDEXTestUniqueDeviceListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDEXTestUniqueDeviceListResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustDEXTestUniqueDeviceListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustDEXTestUniqueDeviceListResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDEXTestUniqueDeviceListResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDEXTestUniqueDeviceListResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustDEXTestUniqueDeviceListResponseEnvelope]
type zeroTrustDEXTestUniqueDeviceListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTestUniqueDeviceListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestUniqueDeviceListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestUniqueDeviceListResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zeroTrustDEXTestUniqueDeviceListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDEXTestUniqueDeviceListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustDEXTestUniqueDeviceListResponseEnvelopeErrors]
type zeroTrustDEXTestUniqueDeviceListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTestUniqueDeviceListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestUniqueDeviceListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXTestUniqueDeviceListResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zeroTrustDEXTestUniqueDeviceListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDEXTestUniqueDeviceListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustDEXTestUniqueDeviceListResponseEnvelopeMessages]
type zeroTrustDEXTestUniqueDeviceListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTestUniqueDeviceListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXTestUniqueDeviceListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustDEXTestUniqueDeviceListResponseEnvelopeSuccess bool

const (
	ZeroTrustDEXTestUniqueDeviceListResponseEnvelopeSuccessTrue ZeroTrustDEXTestUniqueDeviceListResponseEnvelopeSuccess = true
)
