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

// DEXTestUniqueDeviceService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDEXTestUniqueDeviceService]
// method instead.
type DEXTestUniqueDeviceService struct {
	Options []option.RequestOption
}

// NewDEXTestUniqueDeviceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDEXTestUniqueDeviceService(opts ...option.RequestOption) (r *DEXTestUniqueDeviceService) {
	r = &DEXTestUniqueDeviceService{}
	r.Options = opts
	return
}

// Returns unique count of devices that have run synthetic application monitoring
// tests in the past 7 days.
func (r *DEXTestUniqueDeviceService) List(ctx context.Context, params DEXTestUniqueDeviceListParams, opts ...option.RequestOption) (res *DEXTestUniqueDeviceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DEXTestUniqueDeviceListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/tests/unique-devices", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DEXTestUniqueDeviceListResponse struct {
	// total number of unique devices
	UniqueDevicesTotal int64                               `json:"uniqueDevicesTotal,required"`
	JSON               dexTestUniqueDeviceListResponseJSON `json:"-"`
}

// dexTestUniqueDeviceListResponseJSON contains the JSON metadata for the struct
// [DEXTestUniqueDeviceListResponse]
type dexTestUniqueDeviceListResponseJSON struct {
	UniqueDevicesTotal apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DEXTestUniqueDeviceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestUniqueDeviceListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Optionally filter result stats to a specific device(s). Cannot be used in
	// combination with colo param.
	DeviceID param.Field[[]string] `query:"deviceId"`
	// Optionally filter results by test name
	TestName param.Field[string] `query:"testName"`
}

// URLQuery serializes [DEXTestUniqueDeviceListParams]'s query parameters as
// `url.Values`.
func (r DEXTestUniqueDeviceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DEXTestUniqueDeviceListResponseEnvelope struct {
	Errors   []DEXTestUniqueDeviceListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DEXTestUniqueDeviceListResponseEnvelopeMessages `json:"messages,required"`
	Result   DEXTestUniqueDeviceListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DEXTestUniqueDeviceListResponseEnvelopeSuccess `json:"success,required"`
	JSON    dexTestUniqueDeviceListResponseEnvelopeJSON    `json:"-"`
}

// dexTestUniqueDeviceListResponseEnvelopeJSON contains the JSON metadata for the
// struct [DEXTestUniqueDeviceListResponseEnvelope]
type dexTestUniqueDeviceListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestUniqueDeviceListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestUniqueDeviceListResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    dexTestUniqueDeviceListResponseEnvelopeErrorsJSON `json:"-"`
}

// dexTestUniqueDeviceListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DEXTestUniqueDeviceListResponseEnvelopeErrors]
type dexTestUniqueDeviceListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestUniqueDeviceListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTestUniqueDeviceListResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    dexTestUniqueDeviceListResponseEnvelopeMessagesJSON `json:"-"`
}

// dexTestUniqueDeviceListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DEXTestUniqueDeviceListResponseEnvelopeMessages]
type dexTestUniqueDeviceListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestUniqueDeviceListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DEXTestUniqueDeviceListResponseEnvelopeSuccess bool

const (
	DEXTestUniqueDeviceListResponseEnvelopeSuccessTrue DEXTestUniqueDeviceListResponseEnvelopeSuccess = true
)
