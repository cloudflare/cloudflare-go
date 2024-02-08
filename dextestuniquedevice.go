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

// DexTestUniqueDeviceService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDexTestUniqueDeviceService]
// method instead.
type DexTestUniqueDeviceService struct {
	Options []option.RequestOption
}

// NewDexTestUniqueDeviceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDexTestUniqueDeviceService(opts ...option.RequestOption) (r *DexTestUniqueDeviceService) {
	r = &DexTestUniqueDeviceService{}
	r.Options = opts
	return
}

// Returns unique count of devices that have run synthetic application monitoring
// tests in the past 7 days.
func (r *DexTestUniqueDeviceService) List(ctx context.Context, accountID string, query DexTestUniqueDeviceListParams, opts ...option.RequestOption) (res *DexTestUniqueDeviceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DexTestUniqueDeviceListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/tests/unique-devices", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DexTestUniqueDeviceListResponse struct {
	// total number of unique devices
	UniqueDevicesTotal int64                               `json:"uniqueDevicesTotal,required"`
	JSON               dexTestUniqueDeviceListResponseJSON `json:"-"`
}

// dexTestUniqueDeviceListResponseJSON contains the JSON metadata for the struct
// [DexTestUniqueDeviceListResponse]
type dexTestUniqueDeviceListResponseJSON struct {
	UniqueDevicesTotal apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DexTestUniqueDeviceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestUniqueDeviceListParams struct {
	// Optionally filter result stats to a specific device(s). Cannot be used in
	// combination with colo param.
	DeviceID param.Field[[]string] `query:"deviceId"`
	// Optionally filter results by test name
	TestName param.Field[string] `query:"testName"`
}

// URLQuery serializes [DexTestUniqueDeviceListParams]'s query parameters as
// `url.Values`.
func (r DexTestUniqueDeviceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DexTestUniqueDeviceListResponseEnvelope struct {
	Errors   []DexTestUniqueDeviceListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DexTestUniqueDeviceListResponseEnvelopeMessages `json:"messages,required"`
	Result   DexTestUniqueDeviceListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DexTestUniqueDeviceListResponseEnvelopeSuccess `json:"success,required"`
	JSON    dexTestUniqueDeviceListResponseEnvelopeJSON    `json:"-"`
}

// dexTestUniqueDeviceListResponseEnvelopeJSON contains the JSON metadata for the
// struct [DexTestUniqueDeviceListResponseEnvelope]
type dexTestUniqueDeviceListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestUniqueDeviceListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestUniqueDeviceListResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    dexTestUniqueDeviceListResponseEnvelopeErrorsJSON `json:"-"`
}

// dexTestUniqueDeviceListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DexTestUniqueDeviceListResponseEnvelopeErrors]
type dexTestUniqueDeviceListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestUniqueDeviceListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestUniqueDeviceListResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    dexTestUniqueDeviceListResponseEnvelopeMessagesJSON `json:"-"`
}

// dexTestUniqueDeviceListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DexTestUniqueDeviceListResponseEnvelopeMessages]
type dexTestUniqueDeviceListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestUniqueDeviceListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DexTestUniqueDeviceListResponseEnvelopeSuccess bool

const (
	DexTestUniqueDeviceListResponseEnvelopeSuccessTrue DexTestUniqueDeviceListResponseEnvelopeSuccess = true
)
